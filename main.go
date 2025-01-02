package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	in := os.Getenv("INPUT_IN")
	out := os.Getenv("INPUT_OUT")

	renderer := newPdfRenderer()
	if err := renderer.Run(in, out); err != nil {
		log.Fatal(err)
	}
}

func Fn(name string) {
	fmt.Printf("Hello %s\n", name)
}
