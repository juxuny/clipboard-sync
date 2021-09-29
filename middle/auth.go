package middle

import (
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/env"
	"github.com/juxuny/clipboard-sync/lib/log"
	"github.com/juxuny/clipboard-sync/lib/router"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
)

func Auth(context *routing.Context) error {
	path := string(context.Request.RequestURI())
	if router.IsNoAuthPath(path) {
		return nil
	}
	token := string(context.QueryArgs().Peek("token"))
	if token == "" && context.IsPost() {
		token = string(context.Request.PostArgs().Peek("token"))
	}
	if token == "" {
		token = string(context.Request.Header.Peek("token"))
	}
	if token == "" {
		token = string(context.Request.Header.Peek("AccessToken"))
	}
	if token == "" {
		respFailed(context, "invalid token", http.StatusBadRequest)
		return nil
	}
	if err := checkAllowToken(token); err != nil {
		log.Error("not allow token: ", token)
		respFailed(context, "invalid token")
		return nil
	}
	needCheckSign := env.GetBool(env.Key.CheckSign, false)
	if needCheckSign {
		secret := env.GetString(env.Key.SignSecret)
		var query string
		if context.IsPost() {
			query = string(context.PostBody())
		} else {
			query = context.QueryArgs().String()
		}
		computeSign := lib.MD5(query + secret)
		receiveSign := string(context.Request.Header.Peek("sign"))
		log.Debug("sign: ", computeSign)
		if receiveSign != computeSign {
			log.Error("invalid sign")
			respFailed(context, "invalid request")
			return nil
		}
	}
	//claims, err := jwt.Parse(token)
	//if err != nil {
	//	log.Error(err)
	//	respFailed(context, "请登录", http.StatusUnauthorized)
	//	return nil
	//}
	//context.Set(router.Key.UserId, claims.UserId)
	return nil
}

func checkAllowToken(token string) error {
	list := env.GetStringList(env.Key.AllowToken, ",", []string{})
	for _, item := range list {
		if token == item {
			return nil
		}
	}
	return lib.ErrNotFound
}
