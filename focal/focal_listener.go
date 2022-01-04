// Code generated from focal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package focal // focal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// focalListener is a complete listener for a parse tree produced by focalParser.
type focalListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterGrpnum is called when entering the grpnum production.
	EnterGrpnum(c *GrpnumContext)

	// EnterLinenum is called when entering the linenum production.
	EnterLinenum(c *LinenumContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterAsk is called when entering the ask production.
	EnterAsk(c *AskContext)

	// EnterAskpair is called when entering the askpair production.
	EnterAskpair(c *AskpairContext)

	// EnterDo_ is called when entering the do_ production.
	EnterDo_(c *Do_Context)

	// EnterFor_ is called when entering the for_ production.
	EnterFor_(c *For_Context)

	// EnterQuit is called when entering the quit production.
	EnterQuit(c *QuitContext)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// EnterGoto_ is called when entering the goto_ production.
	EnterGoto_(c *Goto_Context)

	// EnterIf_ is called when entering the if_ production.
	EnterIf_(c *If_Context)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeexpression is called when entering the typeexpression production.
	EnterTypeexpression(c *TypeexpressionContext)

	// EnterReturn_ is called when entering the return_ production.
	EnterReturn_(c *Return_Context)

	// EnterWrite_ is called when entering the write_ production.
	EnterWrite_(c *Write_Context)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterMantissa is called when entering the mantissa production.
	EnterMantissa(c *MantissaContext)

	// EnterSigned_ is called when entering the signed_ production.
	EnterSigned_(c *Signed_Context)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitGrpnum is called when exiting the grpnum production.
	ExitGrpnum(c *GrpnumContext)

	// ExitLinenum is called when exiting the linenum production.
	ExitLinenum(c *LinenumContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitAsk is called when exiting the ask production.
	ExitAsk(c *AskContext)

	// ExitAskpair is called when exiting the askpair production.
	ExitAskpair(c *AskpairContext)

	// ExitDo_ is called when exiting the do_ production.
	ExitDo_(c *Do_Context)

	// ExitFor_ is called when exiting the for_ production.
	ExitFor_(c *For_Context)

	// ExitQuit is called when exiting the quit production.
	ExitQuit(c *QuitContext)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)

	// ExitGoto_ is called when exiting the goto_ production.
	ExitGoto_(c *Goto_Context)

	// ExitIf_ is called when exiting the if_ production.
	ExitIf_(c *If_Context)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeexpression is called when exiting the typeexpression production.
	ExitTypeexpression(c *TypeexpressionContext)

	// ExitReturn_ is called when exiting the return_ production.
	ExitReturn_(c *Return_Context)

	// ExitWrite_ is called when exiting the write_ production.
	ExitWrite_(c *Write_Context)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitMantissa is called when exiting the mantissa production.
	ExitMantissa(c *MantissaContext)

	// ExitSigned_ is called when exiting the signed_ production.
	ExitSigned_(c *Signed_Context)
}
