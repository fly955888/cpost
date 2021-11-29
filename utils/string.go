package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	s := fmt.Sprintf("%x", md5.Sum(data))
	return s
}
