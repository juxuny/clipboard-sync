package log

type gormLogger struct {
	*Logger
}

func NewGormLogger() *gormLogger {
	l := gormLogger{
		Logger: NewPrefix("gorm").SetReportCaller(false),
	}
	return &l
}

// Print format & print log
func (l gormLogger) Print(values ...interface{}) {
	l.Logger.Print(gormLogFormatter(values...)...)
}
