package middle

import (
	"fmt"
	"github.com/juxuny/clipboard-sync/cmd/api/param"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/lib/router"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
	"strings"
	"time"
)

func Base(context *routing.Context) error {
	var method = http.MethodGet
	var code int
	var path = context.Request.RequestURI()
	var args string
	if context.IsPost() {
		args = string(context.PostBody())
		method = http.MethodPost
	} else {
		args = string(context.QueryArgs().QueryString())
	}

	ip := parseRealIp(context)

	startTime := time.Now()
	context.Response.Header.Set("Content-Type", "application/json;charset=utf8")
	defer func() {
		code = context.Response.StatusCode()
		duration := time.Now().Sub(startTime)
		log.Info(fmt.Sprintf("[%s] %s | %d | %s | %s | %v", ip, method, code, path, args, duration))
	}()
	if err := context.Next(); err != nil {
		log.Error(err)
		respServerError(context, err.Error())
		return nil
	}
	return nil
}

func parseRealIp(context *routing.Context) string {
	ip := string(context.Request.Header.Peek("X-Appengine-Remote-Addr"))
	if ip == "" {
		ip = string(context.Request.Header.Peek("X-Forwarded-Proto"))
	}
	if ip == "" {
		ip = string(context.Request.Header.Peek("X-Real-Ip"))
	}
	if ip == "" {
		s := context.Conn().RemoteAddr().String()
		l := strings.Split(s, ":")
		if len(l) > 1 {
			ip = l[0]
		}
	}
	context.Set(router.Key.RealIp, ip)
	return ip
}

func respServerError(context *routing.Context, msg string) {
	context.Response.SetStatusCode(http.StatusBadRequest)
	_ = context.WriteData(param.FailedJson(msg))
	context.Abort()
}

func respFailed(context *routing.Context, msg string, code ...int) {
	c := http.StatusBadRequest
	if len(code) > 0 {
		c = code[0]
	}
	context.Response.SetStatusCode(c)
	failed := param.Failed(msg)
	failed.Code = c
	_, _ = context.WriteString(lib.ToJSON(failed))
	context.Abort()
}
