package main

import "fmt"

const englishHelloPrefix = "Hello "
const koreaHelloPrefix = "Hello "
const english = "English"
const korea = "Korea"

//Hello is return hello
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case korea:
		prefix = koreaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
