// Code generated from STL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stl // STL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// STLListener is a complete listener for a parse tree produced by STLParser.
type STLListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterTriangle is called when entering the triangle production.
	EnterTriangle(c *TriangleContext)

	// EnterVertex is called when entering the vertex production.
	EnterVertex(c *VertexContext)

	// EnterTriple is called when entering the triple production.
	EnterTriple(c *TripleContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterFooter is called when entering the footer production.
	EnterFooter(c *FooterContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitTriangle is called when exiting the triangle production.
	ExitTriangle(c *TriangleContext)

	// ExitVertex is called when exiting the vertex production.
	ExitVertex(c *VertexContext)

	// ExitTriple is called when exiting the triple production.
	ExitTriple(c *TripleContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitFooter is called when exiting the footer production.
	ExitFooter(c *FooterContext)
}
