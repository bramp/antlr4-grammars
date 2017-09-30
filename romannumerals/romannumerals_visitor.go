// Generated from romannumerals.g4 by ANTLR 4.7.

package romannumerals // romannumerals
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by romannumeralsParser.
type romannumeralsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by romannumeralsParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#thousands.
	VisitThousands(ctx *ThousandsContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#thous_part.
	VisitThous_part(ctx *Thous_partContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#hundreds.
	VisitHundreds(ctx *HundredsContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#hun_part.
	VisitHun_part(ctx *Hun_partContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#hun_rep.
	VisitHun_rep(ctx *Hun_repContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#tens.
	VisitTens(ctx *TensContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#tens_part.
	VisitTens_part(ctx *Tens_partContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#tens_rep.
	VisitTens_rep(ctx *Tens_repContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#ones.
	VisitOnes(ctx *OnesContext) interface{}

	// Visit a parse tree produced by romannumeralsParser#ones_rep.
	VisitOnes_rep(ctx *Ones_repContext) interface{}
}
