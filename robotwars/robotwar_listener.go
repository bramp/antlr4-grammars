// Generated from robotwar.g4 by ANTLR 4.7.

package robotwars // robotwar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// robotwarListener is a complete listener for a parse tree produced by robotwarParser.
type robotwarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAccumstatement is called when entering the accumstatement production.
	EnterAccumstatement(c *AccumstatementContext)

	// EnterAccumexpression is called when entering the accumexpression production.
	EnterAccumexpression(c *AccumexpressionContext)

	// EnterGosubstatement is called when entering the gosubstatement production.
	EnterGosubstatement(c *GosubstatementContext)

	// EnterGotostatement is called when entering the gotostatement production.
	EnterGotostatement(c *GotostatementContext)

	// EnterTostatement is called when entering the tostatement production.
	EnterTostatement(c *TostatementContext)

	// EnterEndsubstatement is called when entering the endsubstatement production.
	EnterEndsubstatement(c *EndsubstatementContext)

	// EnterIfstatement is called when entering the ifstatement production.
	EnterIfstatement(c *IfstatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAccumstatement is called when exiting the accumstatement production.
	ExitAccumstatement(c *AccumstatementContext)

	// ExitAccumexpression is called when exiting the accumexpression production.
	ExitAccumexpression(c *AccumexpressionContext)

	// ExitGosubstatement is called when exiting the gosubstatement production.
	ExitGosubstatement(c *GosubstatementContext)

	// ExitGotostatement is called when exiting the gotostatement production.
	ExitGotostatement(c *GotostatementContext)

	// ExitTostatement is called when exiting the tostatement production.
	ExitTostatement(c *TostatementContext)

	// ExitEndsubstatement is called when exiting the endsubstatement production.
	ExitEndsubstatement(c *EndsubstatementContext)

	// ExitIfstatement is called when exiting the ifstatement production.
	ExitIfstatement(c *IfstatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
