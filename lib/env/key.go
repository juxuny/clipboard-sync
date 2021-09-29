package env

import "github.com/juxuny/env/ks"

// 定义系统变量的Key，每个变量一个字段， init函数初始化之后，自动给 Key 里面所有变量赋值 LuckyGiftRule = "LUCKY_GIFT_RULE"
var Key struct {
	Mode             string
	Port             string
	WorkerId         string // snowflake 生成唯一ID需要用一个workerId
	JwtSecret        string
	EncryptSecretKey string // 接口加密用的Key
	CheckSign        string
	SignSecret       string // 请求参数签名secret
	LogAppName       string // 输出日志时带上的标识
	LogServerPrefix  string
	LogServerHost    string

	AllowToken string // 允许的token
}

func init() {
	ks.InitKeyName(&Key, true)
}
