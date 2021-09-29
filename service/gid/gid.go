package gid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/env"
	"github.com/juxuny/clipboard-sync/lib/log"
)

var node *snowflake.Node

var gidLogger = log.NewPrefix("[gid-service]")

func init() {
	var err error
	node, err = snowflake.NewNode(env.GetInt64(env.Key.WorkerId, 1))
	if err != nil {
		gidLogger.Error(err)
		panic(err)
	}
}

func GID() lib.ID {
	id := node.Generate()
	return lib.ID(id)
}

func OrderNo(prefix ...string) string {
	p := "o"
	if len(prefix) > 0 {
		p = prefix[0]
	}
	return GID().String(p)
}
