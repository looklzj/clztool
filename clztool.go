package clztool

import (
	"crypto/rand"
	"fmt"
	"reflect"
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

func Includes(arr interface{}, item interface{}) bool {
	var temp []interface{}
	v := reflect.ValueOf(arr)
	switch v.Type().String() {
	case "[]string":
		for _, v2 := range arr.([]string) {
			temp = append(temp, v2)
		}
	case "[]int":
		for _, v2 := range arr.([]int) {
			temp = append(temp, v2)
		}
	case "[]uint":
		for _, v2 := range arr.([]uint) {
			temp = append(temp, v2)
		}
	case "[]int64":
		for _, v2 := range arr.([]int64) {
			temp = append(temp, v2)
		}
	case "[]uint64":
		for _, v2 := range arr.([]uint64) {
			temp = append(temp, v2)
		}
	case "[]interface{}":
		for _, v2 := range arr.([]interface{}) {
			temp = append(temp, v2)
		}
	}
	itemt := fmt.Sprintf("%v", item)
	for _, v := range temp {
		if fmt.Sprintf("%v", v) == itemt {
			return true
		}
	}
	return false
}

func Remove(arr interface{}, item interface{}) interface{} {

	itemt := fmt.Sprintf("%v", item)
	typ := reflect.ValueOf(arr)
	switch typ.Type().String() {
	case "[]string":
		for k, v := range arr.([]string) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]string)
	case "[]int":
		for k, v := range arr.([]int) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]int)
	case "[]uint":
		for k, v := range arr.([]uint) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]uint)
	case "[]int64":
		for k, v := range arr.([]int64) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]int64)
	case "[]uint64":
		for k, v := range arr.([]uint64) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]uint64)
	case "[]interface{}":
		for k, v := range arr.([]interface{}) {
			if fmt.Sprintf("%v", v) == itemt {
				return append(arr.([]string)[:k], arr.([]string)[k+1:]...)
			}
		}
		return arr.([]interface{})
	}

	return nil
}
