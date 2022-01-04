// Code generated from microc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package microc // microc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// microcListener is a complete listener for a parse tree produced by microcParser.
type microcListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfstatement is called when entering the ifstatement production.
	EnterIfstatement(c *IfstatementContext)

	// EnterWhilestatement is called when entering the whilestatement production.
	EnterWhilestatement(c *WhilestatementContext)

	// EnterBlockstatement is called when entering the blockstatement production.
	EnterBlockstatement(c *BlockstatementContext)

	// EnterExprstatement is called when entering the exprstatement production.
	EnterExprstatement(c *ExprstatementContext)

	// EnterParen_expr is called when entering the paren_expr production.
	EnterParen_expr(c *Paren_exprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterSum_ is called when entering the sum_ production.
	EnterSum_(c *Sum_Context)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfstatement is called when exiting the ifstatement production.
	ExitIfstatement(c *IfstatementContext)

	// ExitWhilestatement is called when exiting the whilestatement production.
	ExitWhilestatement(c *WhilestatementContext)

	// ExitBlockstatement is called when exiting the blockstatement production.
	ExitBlockstatement(c *BlockstatementContext)

	// ExitExprstatement is called when exiting the exprstatement production.
	ExitExprstatement(c *ExprstatementContext)

	// ExitParen_expr is called when exiting the paren_expr production.
	ExitParen_expr(c *Paren_exprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitSum_ is called when exiting the sum_ production.
	ExitSum_(c *Sum_Context)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)
}
