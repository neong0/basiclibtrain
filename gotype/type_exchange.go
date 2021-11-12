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
