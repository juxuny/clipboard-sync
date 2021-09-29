package daemon

import (
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/lib/trace"
	"time"
)

var jobMap = map[*Job]struct{}{}

func Register(j *Job) {
	jobMap[j] = struct{}{}
}

func Start() {
	log.Info("start daemon")
	for job := range jobMap {
		trace.GoRun(func() {
			ticker := time.NewTicker(job.Duration)
			for range ticker.C {
				func() {
					defer lib.CollectRecover()
					job.Func()
				}()
			}
		})
	}
}
