package middle

import (
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/lib/trace"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Trace(context *routing.Context) error {
	trace.InitReqId()
	defer func() {
		trace.CleanReqId()
	}()
	if err := context.Next(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
