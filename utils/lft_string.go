// @Title tools
// @Description // TODO
// @Author chenhaoran
// @Datetime  2021/2/27 11:33 上午
package tools

import (
	"bytes"
	"crypto/md5"
	r "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}
func BytesSplit(s []byte, n int) [][]byte {
	if n <= 0 {
		n = len(s)
	}
	a := make([][]byte, n)
	var size int
	na := 0
	for len(s) > 0 {
		if na+1 >= n {
			a[na] = s
			na++
			break
		}
		_, size = utf8.DecodeRune(s)
		a[na] = s[0:size:size]
		s = s[size:]
		na++
	}
	return a[0:na]
}

//MD5算法
func Md5(str string) (result string) {
	m := md5.New()
	m.Write([]byte(str))
	result = hex.EncodeToString(m.Sum(nil))
	return
}

func Sha256(str string) (result string) {
	if len(str) == 0 {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(str))
	result = hex.EncodeToString(h.Sum(nil))
	return
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func GUID() string {
	b := make([]byte, 48)
	io.ReadFull(r.Reader, b)
	data := []byte(base64.URLEncoding.EncodeToString(b))
	has := md5.Sum(data)
	return fmt.Sprintf("%X", has)
}

func RandomDigit(digit int) string {
	if digit < 1 || digit > 10 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	var target string
	switch digit {
	case 1:
		d := rand.Intn(10)
		target = fmt.Sprintf("%d", d)
	case 2:
		d := rand.Intn(100)
		target = fmt.Sprintf("%02d", d)
	case 3:
		d := rand.Intn(1000)
		target = fmt.Sprintf("%03d", d)
	case 4:
		d := rand.Intn(10000)
		target = fmt.Sprintf("%04d", d)
	case 5:
		d := rand.Intn(100000)
		target = fmt.Sprintf("%05d", d)
	case 6:
		d := rand.Intn(1000000)
		target = fmt.Sprintf("%06d", d)
	case 7:
		d := rand.Intn(10000000)
		target = fmt.Sprintf("%07d", d)
	case 8:
		d := rand.Intn(100000000)
		target = fmt.Sprintf("%08d", d)
	case 9:
		d := rand.Intn(1000000000)
		target = fmt.Sprintf("%09d", d)
	default:

	}
	return target
}
