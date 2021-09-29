package main

import (
	"fmt"
	"github.com/juxuny/clipboard-sync/cmd/api/route"
	"github.com/juxuny/clipboard-sync/lib/env"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/lib/router"
	"github.com/juxuny/clipboard-sync/middle"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {

	r := routing.New()
	builder := router.NewBuilder("/api")
	// add middleware handler
	builder.Use(middle.Trace, middle.Base, middle.Health, middle.Auth)
	// register api group
	builder.Register(route.SyncerInterfaceCreator)

	r = builder.Build(r)
	port := env.GetInt(env.Key.Port, 8080)
	log.Info("listen port:", port)
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), r.HandleRequest); err != nil {
		panic(err)
	}
}
