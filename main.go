package main

import "fmt"

func main() {

	s := "Hello"
	ptr := &s
	fmt.Println(s)
	fmt.Println(ptr)
	fmt.Println(&s)
	fmt.Println(*ptr)
}
