package intser

import (
	"fmt"
	"time"
)

type Config struct {
}

func (c *Config) P() {
	var timestamp = time.Now().Local().Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间:%s \n", timestamp)
}
