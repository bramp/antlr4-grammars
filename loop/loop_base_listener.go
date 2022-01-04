// Code generated from loop.g4 by ANTLR 4.9.3. DO NOT EDIT.

package loop // loop
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseloopListener is a complete listener for a parse tree produced by loopParser.
type BaseloopListener struct{}

var _ loopListener = &BaseloopListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseloopListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseloopListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseloopListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseloopListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseloopListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseloopListener) ExitProg(ctx *ProgContext) {}

// EnterStatementlist is called when production statementlist is entered.
func (s *BaseloopListener) EnterStatementlist(ctx *StatementlistContext) {}

// ExitStatementlist is called when production statementlist is exited.
func (s *BaseloopListener) ExitStatementlist(ctx *StatementlistContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseloopListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseloopListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignstmt is called when production assignstmt is entered.
func (s *BaseloopListener) EnterAssignstmt(ctx *AssignstmtContext) {}

// ExitAssignstmt is called when production assignstmt is exited.
func (s *BaseloopListener) ExitAssignstmt(ctx *AssignstmtContext) {}

// EnterIncrementstmt is called when production incrementstmt is entered.
func (s *BaseloopListener) EnterIncrementstmt(ctx *IncrementstmtContext) {}

// ExitIncrementstmt is called when production incrementstmt is exited.
func (s *BaseloopListener) ExitIncrementstmt(ctx *IncrementstmtContext) {}

// EnterLoopstmt is called when production loopstmt is entered.
func (s *BaseloopListener) EnterLoopstmt(ctx *LoopstmtContext) {}

// ExitLoopstmt is called when production loopstmt is exited.
func (s *BaseloopListener) ExitLoopstmt(ctx *LoopstmtContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseloopListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseloopListener) ExitVar_(ctx *Var_Context) {}

// EnterNumber is called when production number is entered.
func (s *BaseloopListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseloopListener) ExitNumber(ctx *NumberContext) {}
