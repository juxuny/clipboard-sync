package param

import (
	"encoding/json"
	"net/http"
	"strings"
)

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataResp struct {
	BaseResp
	Result interface{} `json:"result"`
}

func Success(message ...string) BaseResp {
	m := "ok"
	if len(message) > 0 {
		m = strings.Join(message, ",")
	}
	return BaseResp{
		Code: http.StatusOK,
		Msg:  m,
	}
}

func SuccessJson(message ...string) []byte {
	data, _ := json.Marshal(Success(message...))
	return data
}

func Failed(message string) BaseResp {
	return BaseResp{
		Code: http.StatusBadRequest,
		Msg:  message,
	}
}

func FailedJson(message string) []byte {
	data, _ := json.Marshal(Failed(message))
	return data
}

func Data(v interface{}) DataResp {
	return DataResp{
		BaseResp: BaseResp{
			Code: http.StatusOK,
		},
		Result: v,
	}
}

func DataJson(v interface{}) []byte {
	data, _ := json.Marshal(Data(v))
	return data
}
