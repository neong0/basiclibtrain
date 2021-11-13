package gotype

import (
	"fmt"
	"strconv"
)

func Int2StringByfmt(i int) string {
	return fmt.Sprintf("%d", i)
}

func Int2StringByStrconv(i int) string {
	return strconv.Itoa(i)
}

func Int642StringByfmt(i int64) string {
	return fmt.Sprintf("%d", i)
}

func Int642StringByStrconv(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Rune2String(s rune) string {
	return string(s)
}

func String2Rune(s string) []rune {
	res := make([]rune, len(s))
	for _, r := range s {
		res = append(res, r)
	}
	return res
}

func ChangeStringByByte(s string) string {
	temp := []byte(s)
	temp[0] = 'A'
	return string(temp)
}
func ChangeStringByRune(s string) string {
	temp := []rune(s)
	temp[0] = '#'
	return string(temp)
}
