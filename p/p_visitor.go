// Generated from p.g4 by ANTLR 4.7.

package p // p
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by pParser.
type pVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by pParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by pParser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by pParser#iterate.
	VisitIterate(ctx *IterateContext) interface{}

	// Visit a parse tree produced by pParser#atom.
	VisitAtom(ctx *AtomContext) interface{}
}
