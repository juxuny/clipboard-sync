package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

func ToJSON(v interface{}) string {
	jsonData, _ := json.Marshal(v)
	return string(jsonData)
}

//密码强度必须为字⺟⼤⼩写+数字+符号，6位以上
func CheckPasswordLever(ps string) error {
	if len(ps) < 6 {
		return fmt.Errorf("密码长度 < 6")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	//A_Z := `[A-Z]{1}`
	//symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return errors.New("密码需要包含数字")
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return errors.New("密码需要包含小写字母")
	}
	//if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
	//	return errors.New("密码需要包含大写字母")
	//}
	//if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
	//	return errors.New("密码需要包含特殊符号")
	//}
	return nil
}

func CheckPaymentPassword(ps string) error {
	if len(ps) != 6 {
		return fmt.Errorf("密码只能是6位数字")
	}
	for _, i := range ps {
		if i < '0' || i > '9' {
			return fmt.Errorf("密码只能是数字")
		}
	}
	return nil
}
