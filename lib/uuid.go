package lib

import (
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func UUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

// 生成账单ID
// 共25位
// YYYYMMDD(8) + hhmmss(6) + ms(3) + randomNumber(11)
func CreateBillID() string {
	n := time.Now()
	return time.Now().Format("20060102150405") + strings.Trim(n.Format(".999"), ".") + RandomNumberString(8)
}
