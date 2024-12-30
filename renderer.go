package main

import (
	"io"

	"github.com/go-pdf/fpdf"
	"github.com/gomarkdown/markdown/ast"
)

type pdfRenderer struct {
	pdf *fpdf.Fpdf
}

// RenderFooter implements markdown.Renderer.
func (p *pdfRenderer) RenderFooter(w io.Writer, ast ast.Node) {
	// panic("unimplemented")
}

// RenderHeader implements markdown.Renderer.
func (p *pdfRenderer) RenderHeader(w io.Writer, ast ast.Node) {
	// panic("unimplemented")
}

func (p *pdfRenderer) document(node *ast.Document, entering bool) {

}

// RenderNode implements markdown.Renderer.
func (p *pdfRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Document:
		p.document(node, entering)
	case *ast.Heading:
	case *ast.Text:
	case *ast.Paragraph:
	case *ast.Strong:
	case *ast.Hardbreak:
	case *ast.Link:
	case *ast.Emph:
	case *ast.List:
	case *ast.ListItem:
	}
	return ast.GoToNext
}

func newPdfRenderer() *pdfRenderer {
	r := new(pdfRenderer)
	r.pdf = fpdf.New("P", "mm", "A4", "")
	return r
}