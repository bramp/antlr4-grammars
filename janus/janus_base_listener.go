// Code generated from janus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package janus // janus
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasejanusListener is a complete listener for a parse tree produced by janusParser.
type BasejanusListener struct{}

var _ janusListener = &BasejanusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejanusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejanusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejanusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejanusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasejanusListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasejanusListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BasejanusListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BasejanusListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasejanusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasejanusListener) ExitStatement(ctx *StatementContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BasejanusListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BasejanusListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterDostmt is called when production dostmt is entered.
func (s *BasejanusListener) EnterDostmt(ctx *DostmtContext) {}

// ExitDostmt is called when production dostmt is exited.
func (s *BasejanusListener) ExitDostmt(ctx *DostmtContext) {}

// EnterCallstmt is called when production callstmt is entered.
func (s *BasejanusListener) EnterCallstmt(ctx *CallstmtContext) {}

// ExitCallstmt is called when production callstmt is exited.
func (s *BasejanusListener) ExitCallstmt(ctx *CallstmtContext) {}

// EnterReadstmt is called when production readstmt is entered.
func (s *BasejanusListener) EnterReadstmt(ctx *ReadstmtContext) {}

// ExitReadstmt is called when production readstmt is exited.
func (s *BasejanusListener) ExitReadstmt(ctx *ReadstmtContext) {}

// EnterWritestmt is called when production writestmt is entered.
func (s *BasejanusListener) EnterWritestmt(ctx *WritestmtContext) {}

// ExitWritestmt is called when production writestmt is exited.
func (s *BasejanusListener) ExitWritestmt(ctx *WritestmtContext) {}

// EnterLvalstmt is called when production lvalstmt is entered.
func (s *BasejanusListener) EnterLvalstmt(ctx *LvalstmtContext) {}

// ExitLvalstmt is called when production lvalstmt is exited.
func (s *BasejanusListener) ExitLvalstmt(ctx *LvalstmtContext) {}

// EnterModstmt is called when production modstmt is entered.
func (s *BasejanusListener) EnterModstmt(ctx *ModstmtContext) {}

// ExitModstmt is called when production modstmt is exited.
func (s *BasejanusListener) ExitModstmt(ctx *ModstmtContext) {}

// EnterSwapstmt is called when production swapstmt is entered.
func (s *BasejanusListener) EnterSwapstmt(ctx *SwapstmtContext) {}

// ExitSwapstmt is called when production swapstmt is exited.
func (s *BasejanusListener) ExitSwapstmt(ctx *SwapstmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasejanusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasejanusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMinexp is called when production minexp is entered.
func (s *BasejanusListener) EnterMinexp(ctx *MinexpContext) {}

// ExitMinexp is called when production minexp is exited.
func (s *BasejanusListener) ExitMinexp(ctx *MinexpContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BasejanusListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BasejanusListener) ExitLvalue(ctx *LvalueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasejanusListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasejanusListener) ExitConstant(ctx *ConstantContext) {}
