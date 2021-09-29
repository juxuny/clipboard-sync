package main

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SyncerGetDataResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Data string `json:"data"`
		Time string `json:"time"`
	} `json:"result"`
}
