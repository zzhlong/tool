package db

import (
	"fmt"
	"strings"
)

//数据库 in ()使用
func ArrayToString(data []string) string {
	result := strings.Builder{}
	for _, v := range data {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write([]byte(fmt.Sprintf(" '%s'", v)))
	}
	return result.String()
}

func ArrayInt64ToString(data []int64) string {
	result := strings.Builder{}
	for _, v := range data {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write([]byte(fmt.Sprintf(" '%d'", v)))
	}
	return result.String()
}

func ArrayIntToString(data []int) string {
	result := strings.Builder{}
	for _, v := range data {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write([]byte(fmt.Sprintf(" '%d'", v)))
	}
	return result.String()
}
