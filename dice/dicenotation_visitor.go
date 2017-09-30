// Generated from DiceNotation.g4 by ANTLR 4.7.

package dice // DiceNotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DiceNotationParser.
type DiceNotationVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DiceNotationParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by DiceNotationParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by DiceNotationParser#binaryOp.
	VisitBinaryOp(ctx *BinaryOpContext) interface{}

	// Visit a parse tree produced by DiceNotationParser#dice.
	VisitDice(ctx *DiceContext) interface{}
}
