// Code generated from pl0.g4 by ANTLR 4.7.2. DO NOT EDIT.

package pl0 // pl0
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basepl0Listener is a complete listener for a parse tree produced by pl0Parser.
type Basepl0Listener struct{}

var _ pl0Listener = &Basepl0Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basepl0Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basepl0Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basepl0Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basepl0Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Basepl0Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Basepl0Listener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *Basepl0Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *Basepl0Listener) ExitBlock(ctx *BlockContext) {}

// EnterConsts is called when production consts is entered.
func (s *Basepl0Listener) EnterConsts(ctx *ConstsContext) {}

// ExitConsts is called when production consts is exited.
func (s *Basepl0Listener) ExitConsts(ctx *ConstsContext) {}

// EnterVars is called when production vars is entered.
func (s *Basepl0Listener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *Basepl0Listener) ExitVars(ctx *VarsContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *Basepl0Listener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *Basepl0Listener) ExitProcedure(ctx *ProcedureContext) {}

// EnterStatement is called when production statement is entered.
func (s *Basepl0Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *Basepl0Listener) ExitStatement(ctx *StatementContext) {}

// EnterAssignstmt is called when production assignstmt is entered.
func (s *Basepl0Listener) EnterAssignstmt(ctx *AssignstmtContext) {}

// ExitAssignstmt is called when production assignstmt is exited.
func (s *Basepl0Listener) ExitAssignstmt(ctx *AssignstmtContext) {}

// EnterCallstmt is called when production callstmt is entered.
func (s *Basepl0Listener) EnterCallstmt(ctx *CallstmtContext) {}

// ExitCallstmt is called when production callstmt is exited.
func (s *Basepl0Listener) ExitCallstmt(ctx *CallstmtContext) {}

// EnterWritestmt is called when production writestmt is entered.
func (s *Basepl0Listener) EnterWritestmt(ctx *WritestmtContext) {}

// ExitWritestmt is called when production writestmt is exited.
func (s *Basepl0Listener) ExitWritestmt(ctx *WritestmtContext) {}

// EnterQstmt is called when production qstmt is entered.
func (s *Basepl0Listener) EnterQstmt(ctx *QstmtContext) {}

// ExitQstmt is called when production qstmt is exited.
func (s *Basepl0Listener) ExitQstmt(ctx *QstmtContext) {}

// EnterBangstmt is called when production bangstmt is entered.
func (s *Basepl0Listener) EnterBangstmt(ctx *BangstmtContext) {}

// ExitBangstmt is called when production bangstmt is exited.
func (s *Basepl0Listener) ExitBangstmt(ctx *BangstmtContext) {}

// EnterBeginstmt is called when production beginstmt is entered.
func (s *Basepl0Listener) EnterBeginstmt(ctx *BeginstmtContext) {}

// ExitBeginstmt is called when production beginstmt is exited.
func (s *Basepl0Listener) ExitBeginstmt(ctx *BeginstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *Basepl0Listener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *Basepl0Listener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *Basepl0Listener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *Basepl0Listener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterCondition is called when production condition is entered.
func (s *Basepl0Listener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *Basepl0Listener) ExitCondition(ctx *ConditionContext) {}

// EnterExpression is called when production expression is entered.
func (s *Basepl0Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basepl0Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *Basepl0Listener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *Basepl0Listener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *Basepl0Listener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *Basepl0Listener) ExitFactor(ctx *FactorContext) {}

// EnterIdent is called when production ident is entered.
func (s *Basepl0Listener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *Basepl0Listener) ExitIdent(ctx *IdentContext) {}

// EnterNumber is called when production number is entered.
func (s *Basepl0Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Basepl0Listener) ExitNumber(ctx *NumberContext) {}
