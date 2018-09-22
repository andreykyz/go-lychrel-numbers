package main

import (
	"fmt"
	"time"
)

func main() {
	//candidate := []byte{1,9,6}
	candidate := []byte{1,1}
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
func testPalindrom(str []byte, size int64, iter int64) bool{
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

func add(str []byte, size int64) []byte {
	result:=[]byte{}
	j:=0
	carry := byte(0)
	for i:=size-1; i >= 0; i--{
		res := str[i] + str[j] + carry
		if res > 9 {
			res = res-10
			carry = 1
		} else {
			carry = 0
		}
		result = append(result, res)
		j++
	}
	if carry == 1 {
		result = append(result, 1)
	}
	return result
}
