package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ID int64

func (id ID) String(prefix ...string) string {
	if len(prefix) > 0 {
		return prefix[0] + fmt.Sprintf("%v", int64(id))
	}
	return fmt.Sprintf("%v", int64(id))
}

func (id *ID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*id = ID(v)
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", id)), nil
}

func (id ID) Int64() int64 {
	return int64(id)
}

func ParseID(value interface{}) (ID, error) {
	switch v := value.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		tmp, _ := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		return ID(tmp), nil
	case string:
		tmp, err := strconv.ParseInt(strings.Trim(v, "\"' "), 10, 64)
		return ID(tmp), err
	case ID:
		return v, nil
	}
	return 0, fmt.Errorf("无效ID:%v", value)
}

type IDList []ID

func (t IDList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t IDList) Len() int {
	return len(t)
}

func (t IDList) Less(i, j int) bool {
	return t[i] < t[j]
}

// 无效的ID
const InvalidID = ID(0)

// 启用状态
type EnableStatus int

const (
	EnableStatusDisable = EnableStatus(0) // 启用状态：禁用
	EnableStatusEnable  = EnableStatus(1) // 启用状态：启用
)

// app平台
type AppPlatform int

const (
	AppPlatformDefault = AppPlatform(0) // app平台：默认
	AppPlatformIOS     = AppPlatform(1) // app平台：iOS
	AppPlatformAndroid = AppPlatform(2) // app平台：android
)

// app下载源
type AppInstallSource int

const (
	AppInstallSourceDefault = AppInstallSource(0) // app下载源：默认
	// iOS相关
	AppInstallSourceAppStore  = AppInstallSource(1) // app下载源：苹果商店
	AppInstallSourceDandelion = AppInstallSource(2) // app下载源：蒲公英

	// 官方渠道，iOS和Android都有
	AppInstallSourceOfficial = AppInstallSource(50)

	// Android相关
	AppInstallSourceXiaomi       = AppInstallSource(51)
	AppInstallSourceQihu360      = AppInstallSource(52)
	AppInstallSourceAlibaba      = AppInstallSource(53)
	AppInstallSourceHuawei       = AppInstallSource(54)
	AppInstallSourceVivo         = AppInstallSource(55)
	AppInstallSourceVivoguoshen  = AppInstallSource(56)
	AppInstallSourceOppo         = AppInstallSource(57)
	AppInstallSourceYingyongbao  = AppInstallSource(58)
	AppInstallSourceMeizu        = AppInstallSource(59)
	AppInstallSourceBaiduHelper  = AppInstallSource(60)
	AppInstallSourceBaiduMarket  = AppInstallSource(61)
	AppInstallSourceBaiduNineone = AppInstallSource(62)
	AppInstallSourceShenma1      = AppInstallSource(63)
	AppInstallSourceShenma2      = AppInstallSource(64)
	AppInstallSourceShenma3      = AppInstallSource(65)
	AppInstallSourceShenma4      = AppInstallSource(66)
	AppInstallSourceYidongmm     = AppInstallSource(67)
)

type AppOS string

const (
	AppOSIOS     = AppOS("ios")
	AppOSAndroid = AppOS("android")
)

// app版本号
type AppVersion string

var ErrNotFound = errors.New("not found")

/**
 * 解析为日期
 */
type Date time.Time

func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := ParseDayTime(s)
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(str), nil
}

func (j Date) GetTime() time.Time {
	return time.Time(j)
}

// Maybe a Format function for printing your date
func (j Date) Format(s string) string {
	return j.GetTime().Format(s)
}

/**
 * 解析时间戳
 */
type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%v", time.Time(t).UnixNano()/1e6)
	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := string(b)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(n/1000, 0)
	return nil
}

func (t Timestamp) GetTime() time.Time {
	return time.Time(t)
}

type Bool bool

func (b *Bool) UnmarshalJSON(in []byte) error {
	ok := false
	switch string(in) {
	case "true", "1":
		ok = true
	case "false", "0":
		ok = false
	}
	*b = Bool(ok)
	return nil
}
func (b *Bool) MarshalJSON() ([]byte, error) {
	if *b {
		return []byte("true"), nil
	} else {
		return []byte("false"), nil
	}
}
func (b *Bool) GetBool() bool {
	return bool(*b)
}

type GsonDateTime time.Time

const gsonFormat = "Jan 2, 2006 3:04:05 PM"

// 实现它的json序列化方法
func (this GsonDateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", this.GetTime().Format(gsonFormat))
	return []byte(stamp), nil
}
func (j *GsonDateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation(gsonFormat, s, time.Local)
	if err != nil {
		return err
	}
	*j = GsonDateTime(t)
	return nil
}
func (this *GsonDateTime) GetTime() time.Time {
	return time.Time(*this)
}

type JsonTime time.Time

func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", this.GetTime().Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (this JsonTime) GetTime() time.Time {
	return time.Time(this)
}

func (this *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		return err
	}
	*this = JsonTime(t)
	return nil
}

// string转数组格式
type StringToArray string

func (this StringToArray) MarshalJSON() (by []byte, err error) {
	var arrayString []string
	if arrayString, err = this.GetArray(); err != nil {
		return []byte("[]"), nil
	}
	by, err = json.Marshal(arrayString)
	return
}

func (this StringToArray) GetArray() (arrayString []string, err error) {
	arrayString = make([]string, 0)
	if err = json.Unmarshal([]byte(this), &arrayString); err != nil {
		return
	}
	return
}

// ID集合处理
type IDSet struct {
	m map[ID]bool
}

func NewIDSet(id ...ID) *IDSet {
	ret := &IDSet{
		m: make(map[ID]bool),
	}
	for _, item := range id {
		ret.m[item] = true
	}
	return ret
}

func (t *IDSet) Size() int {
	return len(t.m)
}

func (t *IDSet) Add(id ...ID) *IDSet {
	for _, item := range id {
		t.m[item] = true
	}
	return t
}

func (t *IDSet) Remove(id ...ID) *IDSet {
	for _, item := range id {
		delete(t.m, item)
	}
	return t
}

func (t *IDSet) Range() []ID {
	ret := make([]ID, 0)
	for id := range t.m {
		ret = append(ret, id)
	}
	return ret
}

func (t *IDSet) Clone() *IDSet {
	return NewIDSet(t.Range()...)
}

func (t *IDSet) Contain(id ID) bool {
	return t.m[id]
}

// 去掉已经包含在s集合里的ID
func (t *IDSet) Sub(s *IDSet) *IDSet {
	return t.Remove(s.Range()...)
}

func (t *IDSet) Merge(s *IDSet) *IDSet {
	return t.Add(s.Range()...)
}

// 两个集合的交集
func (t *IDSet) Interact(s *IDSet) *IDSet {
	ret := NewIDSet()
	for _, id := range t.Range() {
		if s.Contain(id) {
			ret.Add(id)
		}
	}
	return ret
}
