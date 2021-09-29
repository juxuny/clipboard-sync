package route

import (
	"github.com/juxuny/clipboard-sync/cmd/api/param"
	"github.com/juxuny/clipboard-sync/lib/router"
	"github.com/juxuny/clipboard-sync/service/cache"
)

func SyncerInterfaceCreator() (interface{}, []interface{}) {
	return &Syncer{}, []interface{}{}
}

type Syncer struct{}

func (*Syncer) SetData(c *router.Context, req *param.SyncerSetDataReq) error {
	_ = cache.HSet(param.SyncerGetDataResp{
		Data: req.Data,
		Time: req.Time,
	}, cache.Key.ClipBoard, req.Token)
	c.MESSAGE("OK")
	return nil
}

func (*Syncer) GetData(c *router.Context, req *param.SyncerGetDataReq) error {
	var out param.SyncerGetDataResp
	_ = cache.HGet(&out, cache.Key.ClipBoard, req.Token)
	c.DATA(out)
	return nil
}
