package main

import "fmt"

const englishHelloPrefi = "Hello "

//Hello is return hello
func Hello(name string) string {
	return englishHelloPrefi + name
}

func main() {
	fmt.Println(Hello("world"))
}
