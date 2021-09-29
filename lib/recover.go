package lib

import (
	"github.com/juxuny/clipboard-sync/lib/log"
	"runtime/debug"
)

func CollectRecover() {
	if err := recover(); err != nil {
		log.Error(err)
		log.Error(string(debug.Stack()))
	}
}
