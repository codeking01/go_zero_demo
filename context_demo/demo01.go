package main

import "fmt"

func main() {
	var test = make(chan bool, 1)
	var test1 = make(chan bool, 1)
	fmt.Printf("test:%T\n", test)
	fmt.Printf("test1:%T\n", test1)
	fmt.Printf("test1:%v\n", test)
	test <- true
	temp := <-test
	fmt.Printf("temp:%v\n", temp)
}
