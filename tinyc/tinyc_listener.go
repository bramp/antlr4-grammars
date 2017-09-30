// Generated from tinyc.g4 by ANTLR 4.7.

package tinyc // tinyc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tinycListener is a complete listener for a parse tree produced by tinycParser.
type tinycListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterParen_expr is called when entering the paren_expr production.
	EnterParen_expr(c *Paren_exprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitParen_expr is called when exiting the paren_expr production.
	ExitParen_expr(c *Paren_exprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)
}
