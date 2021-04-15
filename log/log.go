package log

import (
	"context"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	tools "github.com/zzhlong/tool/utils"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type logEntity struct {
	l       *logrus.Logger
	request string
}

//Log ...
var Log *logEntity

const (
	//Service ...
	Service = "service"
	//Machine ...
	Machine = "machine"
	//Request ...
	Request = "request"
)

//Init ...
func Init(service string, path string, level string, isConsole bool) {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "UnKnown"
	}
	l := logrus.New()
	l.AddHook(&globalHook{
		machine: hostName,
		service: service,
	})
	l.SetFormatter(&formatter{isConsole: isConsole})
	l.SetLevel(getLevel(level))
	l.SetOutput(getWriter(path))
	Log = &logEntity{l: l}
}

func getLevel(level string) logrus.Level {
	level = strings.ToUpper(level)
	switch level {
	case "PANIC":
		return logrus.PanicLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.WarnLevel

	}
}

func getWriter(p string) io.Writer {
	if len(p) == 0 {
		p, e := tools.GetCurrentPath()
		if e != nil {
			panic(e)
		}
		p = path.Join(p, "log")
	}
	writer, err := rotatelogs.New(
		path.Join(p, "%Y%m%d.log"),
		rotatelogs.WithLinkName(p),
		rotatelogs.WithMaxAge(2*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return writer
}

//NewLog ...
func NewLog(ctx context.Context) *logEntity {
	r := ctx.Value(Request)
	if r == nil {
		r = "-"
	}
	l := &logEntity{l: Log.l, request: fmt.Sprintf("%v", r)}
	return l
}
func (l *logEntity) base() *logrus.Entry {
	if len(l.request) == 0 {
		l.request = "-"
	}
	return l.l.WithFields(logrus.Fields{
		Request: l.request,
	})
}
func (l *logEntity) Panic(args ...interface{}) {
	l.base().Panic(args...)
}
func (l *logEntity) Fatal(args ...interface{}) {
	l.base().Fatal(args...)
}
func (l *logEntity) Error(args ...interface{}) {
	l.base().Error(args...)
}
func (l *logEntity) Warn(args ...interface{}) {
	l.base().Warn(args...)
}
func (l *logEntity) Info(args ...interface{}) {
	l.base().Info(args...)
}
func (l *logEntity) Debug(args ...interface{}) {
	l.base().Debug(args...)
}
func (l *logEntity) Trace(args ...interface{}) {
	l.base().Trace(args...)
}
