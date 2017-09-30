// Generated from Abnf.g4 by ANTLR 4.7.

package abnf // Abnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AbnfParser.
type AbnfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AbnfParser#rulelist.
	VisitRulelist(ctx *RulelistContext) interface{}

	// Visit a parse tree produced by AbnfParser#rule_.
	VisitRule_(ctx *Rule_Context) interface{}

	// Visit a parse tree produced by AbnfParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by AbnfParser#alternation.
	VisitAlternation(ctx *AlternationContext) interface{}

	// Visit a parse tree produced by AbnfParser#concatenation.
	VisitConcatenation(ctx *ConcatenationContext) interface{}

	// Visit a parse tree produced by AbnfParser#repetition.
	VisitRepetition(ctx *RepetitionContext) interface{}

	// Visit a parse tree produced by AbnfParser#repeat.
	VisitRepeat(ctx *RepeatContext) interface{}

	// Visit a parse tree produced by AbnfParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by AbnfParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by AbnfParser#option.
	VisitOption(ctx *OptionContext) interface{}
}
