// Generated from SUOKIF.g4 by ANTLR 4.7.

package suokif // SUOKIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SUOKIFParser.
type SUOKIFVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SUOKIFParser#top_level.
	VisitTop_level(ctx *Top_levelContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#funterm.
	VisitFunterm(ctx *FuntermContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#relsent.
	VisitRelsent(ctx *RelsentContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#logsent.
	VisitLogsent(ctx *LogsentContext) interface{}

	// Visit a parse tree produced by SUOKIFParser#quantsent.
	VisitQuantsent(ctx *QuantsentContext) interface{}
}
