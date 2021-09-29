package middle

import routing "github.com/qiangxue/fasthttp-routing"

func Health(context *routing.Context) error {
	if string(context.Request.RequestURI()) == "/healthz" {
		_, _ = context.WriteString("ok")
		context.Abort()
		return nil
	}
	return nil
}
