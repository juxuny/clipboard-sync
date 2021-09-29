package log

import (
	"fmt"
	"github.com/juxuny/clipboard-sync/lib/env"
	"runtime"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

type loggerFormatter struct {
	fields Fields // 输出固定参数
}

// 时间格式
const timeFormat = "2006-01-02 15:04:05"

func (f *loggerFormatter) Format(entry *log.Entry) ([]byte, error) {

	// 将默认的格式化数据添加到日志中
	for k, v := range f.fields {
		entry.Data[k] = v
	}

	var formatter log.Formatter
	if env.Debug() {
		// debug 模式输出单行日志
		gray := color.New(color.FgHiBlack).SprintFunc()
		formatter = &log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
			CallerPrettyfier: func(caller *runtime.Frame) (function string, file string) {
				function = caller.Function + "\n"
				file = fmt.Sprintf(" %s:%d", caller.File, caller.Line)
				return gray(function), gray(file)
			},
		}
	} else {
		// debug 模式输出单行日志
		gray := color.New(color.FgHiBlack).SprintFunc()
		formatter = &log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
			CallerPrettyfier: func(caller *runtime.Frame) (function string, file string) {
				function = caller.Function + " "
				file = fmt.Sprintf(" %s:%d", caller.File, caller.Line)
				return gray(function), gray(file)
			},
		}
		// 线上模式输出json日志
		//formatter = &log.JSONFormatter{
		//	TimestampFormat: timeFormat,
		//}
	}

	return formatter.Format(entry)
}
