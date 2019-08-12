// Code generated from tinybasic.g4 by ANTLR 4.7.2. DO NOT EDIT.

package tinybasic // tinybasic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetinybasicListener is a complete listener for a parse tree produced by tinybasicParser.
type BasetinybasicListener struct{}

var _ tinybasicListener = &BasetinybasicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetinybasicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetinybasicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetinybasicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetinybasicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetinybasicListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetinybasicListener) ExitProgram(ctx *ProgramContext) {}

// EnterLine is called when production line is entered.
func (s *BasetinybasicListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasetinybasicListener) ExitLine(ctx *LineContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasetinybasicListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasetinybasicListener) ExitStatement(ctx *StatementContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BasetinybasicListener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BasetinybasicListener) ExitExprlist(ctx *ExprlistContext) {}

// EnterVarlist is called when production varlist is entered.
func (s *BasetinybasicListener) EnterVarlist(ctx *VarlistContext) {}

// ExitVarlist is called when production varlist is exited.
func (s *BasetinybasicListener) ExitVarlist(ctx *VarlistContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasetinybasicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasetinybasicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BasetinybasicListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasetinybasicListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasetinybasicListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasetinybasicListener) ExitFactor(ctx *FactorContext) {}

// EnterVara is called when production vara is entered.
func (s *BasetinybasicListener) EnterVara(ctx *VaraContext) {}

// ExitVara is called when production vara is exited.
func (s *BasetinybasicListener) ExitVara(ctx *VaraContext) {}

// EnterNumber is called when production number is entered.
func (s *BasetinybasicListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasetinybasicListener) ExitNumber(ctx *NumberContext) {}

// EnterRelop is called when production relop is entered.
func (s *BasetinybasicListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BasetinybasicListener) ExitRelop(ctx *RelopContext) {}
