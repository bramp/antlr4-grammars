// Code generated from fol.g4 by ANTLR 4.7.2. DO NOT EDIT.

package fol // fol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefolListener is a complete listener for a parse tree produced by folParser.
type BasefolListener struct{}

var _ folListener = &BasefolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasefolListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasefolListener) ExitCondition(ctx *ConditionContext) {}

// EnterFormula is called when production formula is entered.
func (s *BasefolListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BasefolListener) ExitFormula(ctx *FormulaContext) {}

// EnterTerm is called when production term is entered.
func (s *BasefolListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasefolListener) ExitTerm(ctx *TermContext) {}

// EnterBin_connective is called when production bin_connective is entered.
func (s *BasefolListener) EnterBin_connective(ctx *Bin_connectiveContext) {}

// ExitBin_connective is called when production bin_connective is exited.
func (s *BasefolListener) ExitBin_connective(ctx *Bin_connectiveContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasefolListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasefolListener) ExitVariable(ctx *VariableContext) {}

// EnterPred_constant is called when production pred_constant is entered.
func (s *BasefolListener) EnterPred_constant(ctx *Pred_constantContext) {}

// ExitPred_constant is called when production pred_constant is exited.
func (s *BasefolListener) ExitPred_constant(ctx *Pred_constantContext) {}

// EnterInd_constant is called when production ind_constant is entered.
func (s *BasefolListener) EnterInd_constant(ctx *Ind_constantContext) {}

// ExitInd_constant is called when production ind_constant is exited.
func (s *BasefolListener) ExitInd_constant(ctx *Ind_constantContext) {}

// EnterFunc_constant is called when production func_constant is entered.
func (s *BasefolListener) EnterFunc_constant(ctx *Func_constantContext) {}

// ExitFunc_constant is called when production func_constant is exited.
func (s *BasefolListener) ExitFunc_constant(ctx *Func_constantContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BasefolListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BasefolListener) ExitSeparator(ctx *SeparatorContext) {}
