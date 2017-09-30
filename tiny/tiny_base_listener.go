// Generated from tiny.g4 by ANTLR 4.7.

package tiny // tiny
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetinyListener is a complete listener for a parse tree produced by tinyParser.
type BasetinyListener struct{}

var _ tinyListener = &BasetinyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetinyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetinyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetinyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetinyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetinyListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetinyListener) ExitProgram(ctx *ProgramContext) {}

// EnterStmt_list is called when production stmt_list is entered.
func (s *BasetinyListener) EnterStmt_list(ctx *Stmt_listContext) {}

// ExitStmt_list is called when production stmt_list is exited.
func (s *BasetinyListener) ExitStmt_list(ctx *Stmt_listContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasetinyListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasetinyListener) ExitStmt(ctx *StmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BasetinyListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BasetinyListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterRead_stmt is called when production read_stmt is entered.
func (s *BasetinyListener) EnterRead_stmt(ctx *Read_stmtContext) {}

// ExitRead_stmt is called when production read_stmt is exited.
func (s *BasetinyListener) ExitRead_stmt(ctx *Read_stmtContext) {}

// EnterWrite_stmt is called when production write_stmt is entered.
func (s *BasetinyListener) EnterWrite_stmt(ctx *Write_stmtContext) {}

// ExitWrite_stmt is called when production write_stmt is exited.
func (s *BasetinyListener) ExitWrite_stmt(ctx *Write_stmtContext) {}

// EnterId_list is called when production id_list is entered.
func (s *BasetinyListener) EnterId_list(ctx *Id_listContext) {}

// ExitId_list is called when production id_list is exited.
func (s *BasetinyListener) ExitId_list(ctx *Id_listContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BasetinyListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BasetinyListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasetinyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasetinyListener) ExitExpr(ctx *ExprContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasetinyListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasetinyListener) ExitFactor(ctx *FactorContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasetinyListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasetinyListener) ExitInteger(ctx *IntegerContext) {}

// EnterOp is called when production op is entered.
func (s *BasetinyListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BasetinyListener) ExitOp(ctx *OpContext) {}

// EnterIdent is called when production ident is entered.
func (s *BasetinyListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BasetinyListener) ExitIdent(ctx *IdentContext) {}
