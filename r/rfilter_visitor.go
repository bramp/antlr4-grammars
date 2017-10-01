// Generated from RFilter.g4 by ANTLR 4.7.

package r // RFilter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RFilter.
type RFilterVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RFilter#stream.
	VisitStream(ctx *StreamContext) interface{}

	// Visit a parse tree produced by RFilter#eat.
	VisitEat(ctx *EatContext) interface{}

	// Visit a parse tree produced by RFilter#elem.
	VisitElem(ctx *ElemContext) interface{}

	// Visit a parse tree produced by RFilter#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by RFilter#op.
	VisitOp(ctx *OpContext) interface{}
}
