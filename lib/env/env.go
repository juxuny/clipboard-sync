package env

import (
	"github.com/juxuny/env"
	"strings"
)

var (
	GetIntList    = env.GetInt
	GetStringList = env.GetStringList
	GetString     = env.GetString
	GetBool       = env.GetBool
	GetInt        = env.GetInt
	GetInt64      = env.GetInt64
)

func GetStringWithDefault(k string, defaultValue string) string {
	v := env.GetString(k)
	if v == "" {
		return defaultValue
	}
	return v
}

func Debug() bool {
	m := GetString(Key.Mode, "debug")
	return strings.ToLower(m) == "debug"
}

func Prod() bool {
	return strings.ToLower(GetString(Key.Mode, "debug")) == "prod"
}
