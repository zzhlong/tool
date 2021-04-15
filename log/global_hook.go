package log

import (
	"github.com/sirupsen/logrus"
)

type globalHook struct {
	service string
	machine string
	levels  []uint32
}

func (h *globalHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *globalHook) Fire(e *logrus.Entry) error {
	e.Data[Service] = h.service
	e.Data[Machine] = h.machine
	return nil
}
