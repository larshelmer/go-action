package main

import "fmt"

func main() {
	Fn("world")
}

func Fn(name string) {
	fmt.Printf("Hello %s\n", name)
}
