package lib

import "fmt"

// 把ID转字符，并用splitter拼接
// 保持 id 原来的顺序
func IdJoin(a []ID, splitter string) string {
	ret := ""
	for i := 0; i < len(a); i++ {
		if ret != "" {
			ret += splitter
		}
		ret += fmt.Sprintf("%v", a[i])
	}
	return ret
}

func IdSlice2StringSlice(ids []ID) []string {
	ret := make([]string, len(ids))
	for i := range ids {
		ret[i] = fmt.Sprintf("%d", ids[i])
	}
	return ret
}
