// Code generated from sickbay.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sickbay // sickbay
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesickbayListener is a complete listener for a parse tree produced by sickbayParser.
type BasesickbayListener struct{}

var _ sickbayListener = &BasesickbayListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesickbayListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesickbayListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesickbayListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesickbayListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSickbay is called when production sickbay is entered.
func (s *BasesickbayListener) EnterSickbay(ctx *SickbayContext) {}

// ExitSickbay is called when production sickbay is exited.
func (s *BasesickbayListener) ExitSickbay(ctx *SickbayContext) {}

// EnterLine is called when production line is entered.
func (s *BasesickbayListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasesickbayListener) ExitLine(ctx *LineContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasesickbayListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasesickbayListener) ExitStmt(ctx *StmtContext) {}

// EnterIntExpr is called when production intExpr is entered.
func (s *BasesickbayListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production intExpr is exited.
func (s *BasesickbayListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterIntVar is called when production intVar is entered.
func (s *BasesickbayListener) EnterIntVar(ctx *IntVarContext) {}

// ExitIntVar is called when production intVar is exited.
func (s *BasesickbayListener) ExitIntVar(ctx *IntVarContext) {}
