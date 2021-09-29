package main

import (
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/curl"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/service/clipboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
	"time"
)

type runCmd struct {
	Flag struct {
		globalFlag
		Host     string
		Secret   string
		Token    string
		Duration int
	}
	localSign        string // local clipboard data md5 value
	localUpdateTime  time.Time
	remoteSign       string // remote clipboard data md5 value
	remoteUpdateTime time.Time
}

func (t *runCmd) initFlag(cmd *cobra.Command) {
	initGlobalFlag(cmd, &t.Flag.globalFlag)
	cmd.PersistentFlags().StringVar(&t.Flag.Host, "host", "", "host address")
	cmd.PersistentFlags().StringVarP(&t.Flag.Secret, "secret", "s", "", "sign secret")
	cmd.PersistentFlags().StringVarP(&t.Flag.Token, "token", "t", "", "access token")
	cmd.PersistentFlags().IntVarP(&t.Flag.Duration, "duration", "d", 6, "sleep seconds")
}

func (t *runCmd) setRemoteData(data []byte) error {
	values := url.Values{}
	values.Set("token", t.Flag.Token)
	values.Set("data", string(data))
	values.Set("time", time.Now().Format(lib.DateTimeLayout))
	sign := lib.MD5(values.Encode() + t.Flag.Secret)
	var resp BaseResp
	if _, err := curl.PostForm(t.Flag.Host+"/api/syncer/set/data", values, &resp, lib.H{
		"sign": sign,
	}); err != nil {
		return err
	}
	if resp.Code != http.StatusOK {
		return errors.Errorf("push clipboard data failed, code: %v", resp.Msg)
	}
	return nil
}

func (t *runCmd) getRemoteData() (data []byte, err error) {
	values := url.Values{}
	values.Set("token", t.Flag.Token)
	sign := lib.MD5(values.Encode() + t.Flag.Secret)
	var resp SyncerGetDataResp
	if _, err := curl.PostForm(t.Flag.Host+"/api/syncer/get/data", values, &resp, lib.H{
		"sign": sign,
	}); err != nil {
		log.Error(err)
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.Errorf("sync clipboard data failed: %v", resp.Msg)
	}
	return
}

func (t *runCmd) run() {
	lib.CollectRecover()
	localData, err := clipboard.GetLocal()
	if err != nil {
		log.Error(err)
		return
	}
	localSign := lib.MD5FromBytes(localData)
	if localSign != t.localSign {
		t.localUpdateTime = time.Now()
		log.Debug("local data update")
	}
	remoteData, err := t.getRemoteData()
	if err != nil {
		log.Error(err)
		return
	}
	remoteSign := lib.MD5FromBytes(remoteData)
	if remoteSign != t.remoteSign {
		t.remoteUpdateTime = time.Now()
		log.Debug("remote data update")
	}
	if t.localSign == t.remoteSign {
		return // the clipboard data is same
	}
	if t.localUpdateTime.After(t.remoteUpdateTime) && len(localData) > 0 {
		log.Debug("set remote clipboard")
		if err := t.setRemoteData(localData); err != nil {
			log.Error(err)
		}
		t.remoteSign = localSign
	}
	if t.localUpdateTime.Before(t.remoteUpdateTime) && len(t.remoteSign) > 0 {
		log.Debug("set local clipboard")
		if err := clipboard.SetLocal(remoteData); err != nil {
			log.Error(err)
		}
		t.localSign = remoteSign
	}
}

func (t *runCmd) Build() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			if t.Flag.Secret == "" {
				log.Error("missing argument: secret")
				return
			}
			if t.Flag.Token == "" {
				log.Error("missing argument: token")
				return
			}
			if t.Flag.Host == "" {
				log.Error("missing argument: host")
				return
			}
			ticker := time.NewTicker(time.Second * time.Duration(t.Flag.Duration))
			for range ticker.C {
				t.run()
			}
		},
	}
	t.initFlag(cmd)
	return cmd
}

func init() {
	rootCmd.AddCommand((&runCmd{}).Build())
}
