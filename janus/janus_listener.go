// Code generated from janus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package janus // janus
import "github.com/antlr/antlr4/runtime/Go/antlr"

// janusListener is a complete listener for a parse tree produced by janusParser.
type janusListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterDostmt is called when entering the dostmt production.
	EnterDostmt(c *DostmtContext)

	// EnterCallstmt is called when entering the callstmt production.
	EnterCallstmt(c *CallstmtContext)

	// EnterReadstmt is called when entering the readstmt production.
	EnterReadstmt(c *ReadstmtContext)

	// EnterWritestmt is called when entering the writestmt production.
	EnterWritestmt(c *WritestmtContext)

	// EnterLvalstmt is called when entering the lvalstmt production.
	EnterLvalstmt(c *LvalstmtContext)

	// EnterModstmt is called when entering the modstmt production.
	EnterModstmt(c *ModstmtContext)

	// EnterSwapstmt is called when entering the swapstmt production.
	EnterSwapstmt(c *SwapstmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMinexp is called when entering the minexp production.
	EnterMinexp(c *MinexpContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitDostmt is called when exiting the dostmt production.
	ExitDostmt(c *DostmtContext)

	// ExitCallstmt is called when exiting the callstmt production.
	ExitCallstmt(c *CallstmtContext)

	// ExitReadstmt is called when exiting the readstmt production.
	ExitReadstmt(c *ReadstmtContext)

	// ExitWritestmt is called when exiting the writestmt production.
	ExitWritestmt(c *WritestmtContext)

	// ExitLvalstmt is called when exiting the lvalstmt production.
	ExitLvalstmt(c *LvalstmtContext)

	// ExitModstmt is called when exiting the modstmt production.
	ExitModstmt(c *ModstmtContext)

	// ExitSwapstmt is called when exiting the swapstmt production.
	ExitSwapstmt(c *SwapstmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMinexp is called when exiting the minexp production.
	ExitMinexp(c *MinexpContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}
