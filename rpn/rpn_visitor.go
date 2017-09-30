// Generated from rpn.g4 by ANTLR 4.7.

package rpn // rpn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by rpnParser.
type rpnVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by rpnParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by rpnParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by rpnParser#oper.
	VisitOper(ctx *OperContext) interface{}

	// Visit a parse tree produced by rpnParser#signedAtom.
	VisitSignedAtom(ctx *SignedAtomContext) interface{}

	// Visit a parse tree produced by rpnParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by rpnParser#scientific.
	VisitScientific(ctx *ScientificContext) interface{}
}
