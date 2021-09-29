package cache

import "github.com/juxuny/env/ks"

// key 与redis保持一致
var Key = struct {
	AllowToken string
	ClipBoard  string
}{}

//var Key = struct {
//	ListeningKey string
//	SysConf      string
//}{}
//
func init() {
	ks.InitKeyName(&Key, false)
}
