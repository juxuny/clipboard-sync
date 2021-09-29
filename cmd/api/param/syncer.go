package param

import (
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"time"
)

type SyncerSetDataReq struct {
	Token      string `json:"token" schema:"token"`
	Data       string `json:"data" schema:"data"`
	Time       string `json:"time" schema:"time"`
	parsedTime time.Time
}

type SyncerGetDataReq struct {
	Token string `json:"token" schema:"token"`
}

type SyncerGetDataResp struct {
	Data string `json:"data"`
	Time string `json:"time"`
}

func (t *SyncerSetDataReq) Validate() error {
	var err error
	t.parsedTime, err = lib.ParseDateTime(t.Time)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
