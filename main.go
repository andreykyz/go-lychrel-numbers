package main

import (
	"fmt"
	"time"
)

func main() {
	//candidate := "196"
	candidate := "89"
	size := int64(len(candidate))
	iter:=int64(1)
	start := time.Now().UnixNano()
	timeCounter :=int64(0)
	for size < 5000000000 {
		candidate=add(candidate,size)
		size = int64(len(candidate))
//		if testPalindrom(candidate,size,iter) { break}
		go testPalindrom(candidate,size,iter)
		iter++
		timeCounter++
		if timeCounter == 1000 {
			next := time.Now().UnixNano()
			fmt.Println("rate is ", 1000000.0/float64((next-start)/1000000.) , " len is ", size)
			timeCounter = 0 
			start=next
		}
	}
}
func testPalindrom(str string, size int64, iter int64) bool{
	j:=0
	for i:=size-1 ; i > (size>>1); i--{
		if str[i] != str[j] {
			return false
		}
		j++
	}
	fmt.Println("got it len is ", size, " iteration ", iter, "str ", str)
	return true
}

func add(str string, size int64) string {
	result:=""
	j:=0
	carry := byte(0)
	for i:=size-1; i >= 0; i--{
		var res [1]byte
		res[0] = (str[i]) + (str[j]-48) + carry
		if res[0] > (9 + 48) {
			res[0] = res[0]-10
			carry = 1
		} else {
			carry = 0
		}
		result += string(res[:])
		j++
	}
	if carry == 1 {
		result += "1"
	}
	return result
}
