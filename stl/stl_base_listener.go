// Code generated from STL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stl // STL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSTLListener is a complete listener for a parse tree produced by STLParser.
type BaseSTLListener struct{}

var _ STLListener = &BaseSTLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSTLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSTLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSTLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSTLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile_ is called when production file_ is entered.
func (s *BaseSTLListener) EnterFile_(ctx *File_Context) {}

// ExitFile_ is called when production file_ is exited.
func (s *BaseSTLListener) ExitFile_(ctx *File_Context) {}

// EnterTriangle is called when production triangle is entered.
func (s *BaseSTLListener) EnterTriangle(ctx *TriangleContext) {}

// ExitTriangle is called when production triangle is exited.
func (s *BaseSTLListener) ExitTriangle(ctx *TriangleContext) {}

// EnterVertex is called when production vertex is entered.
func (s *BaseSTLListener) EnterVertex(ctx *VertexContext) {}

// ExitVertex is called when production vertex is exited.
func (s *BaseSTLListener) ExitVertex(ctx *VertexContext) {}

// EnterTriple is called when production triple is entered.
func (s *BaseSTLListener) EnterTriple(ctx *TripleContext) {}

// ExitTriple is called when production triple is exited.
func (s *BaseSTLListener) ExitTriple(ctx *TripleContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseSTLListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseSTLListener) ExitHeader(ctx *HeaderContext) {}

// EnterFooter is called when production footer is entered.
func (s *BaseSTLListener) EnterFooter(ctx *FooterContext) {}

// ExitFooter is called when production footer is exited.
func (s *BaseSTLListener) ExitFooter(ctx *FooterContext) {}
