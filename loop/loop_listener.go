// Code generated from loop.g4 by ANTLR 4.9.3. DO NOT EDIT.

package loop // loop
import "github.com/antlr/antlr4/runtime/Go/antlr"

// loopListener is a complete listener for a parse tree produced by loopParser.
type loopListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatementlist is called when entering the statementlist production.
	EnterStatementlist(c *StatementlistContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignstmt is called when entering the assignstmt production.
	EnterAssignstmt(c *AssignstmtContext)

	// EnterIncrementstmt is called when entering the incrementstmt production.
	EnterIncrementstmt(c *IncrementstmtContext)

	// EnterLoopstmt is called when entering the loopstmt production.
	EnterLoopstmt(c *LoopstmtContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatementlist is called when exiting the statementlist production.
	ExitStatementlist(c *StatementlistContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignstmt is called when exiting the assignstmt production.
	ExitAssignstmt(c *AssignstmtContext)

	// ExitIncrementstmt is called when exiting the incrementstmt production.
	ExitIncrementstmt(c *IncrementstmtContext)

	// ExitLoopstmt is called when exiting the loopstmt production.
	ExitLoopstmt(c *LoopstmtContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
