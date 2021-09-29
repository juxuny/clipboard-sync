package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 产生随机字符串
// length : 字符串长度
// source : 包含字符
func RandomString(lens ...interface{}) string {
	l := 8
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(lens) > 0 {
		l = lens[0].(int)
	}
	if len(lens) > 1 {
		str = lens[1].(string)
	}

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 产生随机数字字符串
func RandomNumberString(l int) string {
	return RandomString(l, "0123456789")
}

// 字符串中出现最多次数的字符的个数
func StringMaxCharCount(s string) int {
	hash := map[string]int{}
	bytes := []byte(s)
	max := 0
	for _, v := range bytes {
		hash[string(v)] += 1
		if hash[string(v)] > max {
			max = hash[string(v)]
		}
	}
	return max
}

// 字符串中出现的字符的数量
func StringCharCount(s string) int {
	hash := map[string]string{}
	bytes := []byte(s)
	for _, v := range bytes {
		hash[string(v)] = string(v)
	}
	return len(hash)
}

// 判断字符串是否是国内手机号
func IsChineseMobileNumber(mobile string) bool {
	ok, _ := regexp.MatchString("1(3|4|6|5|7|8|9)[0-9]{9}", mobile)
	return ok
}

func StringStar(str string, start, end int) string {
	arr := strings.Split(str, "")
	for i, _ := range arr {
		if i >= start && i < end {
			arr[i] = "*"
		}
	}
	return strings.Join(arr, "")
}

// 手机号打星星
func PhoneStar(phone string) string {
	return StringStar(phone, 3, len(phone)-4)
}

// 银行卡号打星星
func BankAccountStar(account string) string {
	return StringStar(account, 3, len(account)-4)
}

func Uin64sToStrings(arr []uint64) []string {
	list := []string{}
	for _, v := range arr {
		list = append(list, fmt.Sprintf("%v", v))
	}
	return list
}

// 字符过长处理
func LenStr(s string) string {
	var userNick string
	str := []rune(s)
	if len(str) > 7 {
		userNick = string(str[:6]) + "..."
	} else {
		userNick = s
	}
	return userNick
}

const IndexNotFound = -1 // 找不到索引
// 搜索次匹配到字符串的index
func Index(strs []string, s string) int {
	for i, v := range strs {
		if v == s {
			return i
		}
	}
	return IndexNotFound
}

func StringToTime(toBeCharge string) (theTime time.Time, err error) {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	var loc *time.Location
	if loc, err = time.LoadLocation("Local"); err != nil { //重要：获取时区
		return theTime, err
	}
	if theTime, err = time.ParseInLocation(timeLayout, toBeCharge, loc); err != nil {
		return theTime, err
	}
	return theTime, nil
}

// 按逗号分割ID列表
func SplitID(s string) (ret []ID, err error) {
	l := strings.Split(s, ",")
	for _, item := range l {
		if v, err := strconv.ParseInt(item, 10, 64); err != nil {
			return ret, errors.Wrapf(err, "无效ID: %v", item)
		} else {
			ret = append(ret, ID(v))
		}
	}
	return
}
