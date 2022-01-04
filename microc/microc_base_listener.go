// Code generated from microc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package microc // microc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemicrocListener is a complete listener for a parse tree produced by microcParser.
type BasemicrocListener struct{}

var _ microcListener = &BasemicrocListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemicrocListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemicrocListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemicrocListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemicrocListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasemicrocListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasemicrocListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasemicrocListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasemicrocListener) ExitStatement(ctx *StatementContext) {}

// EnterIfstatement is called when production ifstatement is entered.
func (s *BasemicrocListener) EnterIfstatement(ctx *IfstatementContext) {}

// ExitIfstatement is called when production ifstatement is exited.
func (s *BasemicrocListener) ExitIfstatement(ctx *IfstatementContext) {}

// EnterWhilestatement is called when production whilestatement is entered.
func (s *BasemicrocListener) EnterWhilestatement(ctx *WhilestatementContext) {}

// ExitWhilestatement is called when production whilestatement is exited.
func (s *BasemicrocListener) ExitWhilestatement(ctx *WhilestatementContext) {}

// EnterBlockstatement is called when production blockstatement is entered.
func (s *BasemicrocListener) EnterBlockstatement(ctx *BlockstatementContext) {}

// ExitBlockstatement is called when production blockstatement is exited.
func (s *BasemicrocListener) ExitBlockstatement(ctx *BlockstatementContext) {}

// EnterExprstatement is called when production exprstatement is entered.
func (s *BasemicrocListener) EnterExprstatement(ctx *ExprstatementContext) {}

// ExitExprstatement is called when production exprstatement is exited.
func (s *BasemicrocListener) ExitExprstatement(ctx *ExprstatementContext) {}

// EnterParen_expr is called when production paren_expr is entered.
func (s *BasemicrocListener) EnterParen_expr(ctx *Paren_exprContext) {}

// ExitParen_expr is called when production paren_expr is exited.
func (s *BasemicrocListener) ExitParen_expr(ctx *Paren_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasemicrocListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasemicrocListener) ExitExpr(ctx *ExprContext) {}

// EnterTest is called when production test is entered.
func (s *BasemicrocListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasemicrocListener) ExitTest(ctx *TestContext) {}

// EnterSum_ is called when production sum_ is entered.
func (s *BasemicrocListener) EnterSum_(ctx *Sum_Context) {}

// ExitSum_ is called when production sum_ is exited.
func (s *BasemicrocListener) ExitSum_(ctx *Sum_Context) {}

// EnterTerm is called when production term is entered.
func (s *BasemicrocListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasemicrocListener) ExitTerm(ctx *TermContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BasemicrocListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BasemicrocListener) ExitId_(ctx *Id_Context) {}

// EnterInteger is called when production integer is entered.
func (s *BasemicrocListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasemicrocListener) ExitInteger(ctx *IntegerContext) {}
