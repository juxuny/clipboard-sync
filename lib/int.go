package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		tmp, _ := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		return tmp, nil
	case string:
		tmp, err := strconv.ParseInt(strings.Trim(v, "\"' "), 10, 64)
		return tmp, err
	default:
		tmp, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		return tmp, err
	}
}

func ParseInt(value interface{}) (int, error) {
	v, err := ParseInt64(value)
	return int(v), err
}

func Int64Pointer(v int64) *int64 {
	return &v
}
