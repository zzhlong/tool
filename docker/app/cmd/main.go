package main

import (
	"github.com/zzhlong/tool/docker/app/intser"
	"time"
)

func main() {
	c := intser.Config{}
	for true {
		time.Sleep(time.Second)
		c.P()
	}
}
