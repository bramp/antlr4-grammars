// Code generated from pl0.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pl0 // pl0
import "github.com/antlr/antlr4/runtime/Go/antlr"

// pl0Listener is a complete listener for a parse tree produced by pl0Parser.
type pl0Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterConsts is called when entering the consts production.
	EnterConsts(c *ConstsContext)

	// EnterVars_ is called when entering the vars_ production.
	EnterVars_(c *Vars_Context)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignstmt is called when entering the assignstmt production.
	EnterAssignstmt(c *AssignstmtContext)

	// EnterCallstmt is called when entering the callstmt production.
	EnterCallstmt(c *CallstmtContext)

	// EnterWritestmt is called when entering the writestmt production.
	EnterWritestmt(c *WritestmtContext)

	// EnterQstmt is called when entering the qstmt production.
	EnterQstmt(c *QstmtContext)

	// EnterBangstmt is called when entering the bangstmt production.
	EnterBangstmt(c *BangstmtContext)

	// EnterBeginstmt is called when entering the beginstmt production.
	EnterBeginstmt(c *BeginstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitConsts is called when exiting the consts production.
	ExitConsts(c *ConstsContext)

	// ExitVars_ is called when exiting the vars_ production.
	ExitVars_(c *Vars_Context)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignstmt is called when exiting the assignstmt production.
	ExitAssignstmt(c *AssignstmtContext)

	// ExitCallstmt is called when exiting the callstmt production.
	ExitCallstmt(c *CallstmtContext)

	// ExitWritestmt is called when exiting the writestmt production.
	ExitWritestmt(c *WritestmtContext)

	// ExitQstmt is called when exiting the qstmt production.
	ExitQstmt(c *QstmtContext)

	// ExitBangstmt is called when exiting the bangstmt production.
	ExitBangstmt(c *BangstmtContext)

	// ExitBeginstmt is called when exiting the beginstmt production.
	ExitBeginstmt(c *BeginstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
