package router

import (
	"fmt"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/env/ks"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
)

var logger = log.NewPrefix("router").SetCallStackDepth(1)

// context store key
var Key = struct {
	UserId string
	RealIp string
}{}

func init() {
	ks.InitKeyName(&Key, false)
}

type Context struct {
	*routing.Context
	Sent bool
}

func (t *Context) GetRealIp() string {
	v := t.Get(Key.RealIp)
	if v != nil {
		ip, _ := v.(string)
		return ip
	}
	return ""
}

func (t *Context) GetUserId() lib.ID {
	v, _ := t.Get(Key.UserId).(lib.ID)
	return v
}

func (t *Context) Json(v interface{}) {
	t.Context.Response.Header.Set("Content-Type", "application/json;charset=utf8")
	_, err := t.Context.WriteString(lib.ToJSON(v))
	if err != nil {
		log.Error(err)
		t.Abort()
	}
}

func (t *Context) DATA(v interface{}, code ...int) {
	c := http.StatusOK
	if len(code) > 0 {
		c = code[0]
	}
	t.Json(map[string]interface{}{
		"code":   c,
		"result": v,
	})
}

func (t *Context) ERROR(err interface{}, code ...int) {
	c := http.StatusBadRequest
	if len(code) > 0 {
		c = code[0]
	}
	logger.Error(err)
	t.Json(map[string]interface{}{
		"code": c,
		"msg":  fmt.Sprintf("%v", err),
	})
}

func (t *Context) MESSAGE(msg string, code ...int) {
	c := http.StatusOK
	if len(code) > 0 {
		c = code[0]
	}
	t.Json(map[string]interface{}{
		"code": c,
		"msg":  msg,
	})
}
