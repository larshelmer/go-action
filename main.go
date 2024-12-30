package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	in := os.Getenv("INPUT_IN")
	out := os.Getenv("INPUT_OUT")
	md, err := os.ReadFile(in)
	if err != nil {
		log.Fatalf("os.ReadFile: %v\n", err)
	}

	p := parser.New()
	d := p.Parse(md)

	renderer := newPdfRenderer()

	pdf := markdown.Render(d, renderer)
	os.WriteFile(out, pdf, 0666)
}

func Fn(name string) {
	fmt.Printf("Hello %s\n", name)
}
