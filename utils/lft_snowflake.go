package tools

import (
	"github.com/bwmarrin/snowflake"
	"hash/fnv"
	"os"
)

var snowNode *snowflake.Node

func init() {
	name, _ := os.Hostname()
	h := fnv.New32()
	h.Write([]byte(name))
	sID := h.Sum32()
	ID := sID % 1024
	snowNode, _ = snowflake.NewNode(int64(ID))
}

//雪花ID
func NextSnowflakeId() int64 {
	return int64(snowNode.Generate())
}
