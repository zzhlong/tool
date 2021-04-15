package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type formatter struct {
	isConsole bool
}

func (s *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	//时间 级别 服务 机器 请求ID 消息
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("%s %s %s %s %s %s \n", timestamp, strings.ToUpper(entry.Level.String()), entry.Data[Service], entry.Data[Machine], entry.Data[Request], strings.Replace(entry.Message, "\n", "\\n", -1))
	if s.isConsole {
		fmt.Print(msg)
	}
	return []byte(msg), nil
}
