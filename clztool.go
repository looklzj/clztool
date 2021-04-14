package clztool

import (
	"crypto/rand"
	"fmt"
)

func RandomString(n int, mode int) string {
	var pool string
	switch mode {
	case 0:
		pool = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 1:
		pool = "0123456789abcdefghijklmnopqrstuvwxyz"
	case 2:
		pool = "0123456789"
	case 3:
		pool = "abcdefghijklmnopqrstuvwxyz"
	case 4:
		pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = pool[b%byte(len(pool))]
	}
	return string(bytes)
}

func Includes(arr []interface{}, item interface{}) bool {
	itemt := fmt.Sprintf("%v", item)
	for _, v := range arr {
		if fmt.Sprintf("%v", v) == itemt {
			return true
		}
	}
	return false
}
