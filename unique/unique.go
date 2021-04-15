package unique

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/bwmarrin/snowflake"
	"io"
)

func UUID32() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))
	return hex.EncodeToString(h.Sum(nil))
}

func NextSnowflakeId() uint64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return uint64(node.Generate())
}
