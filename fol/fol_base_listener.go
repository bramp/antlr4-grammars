// Generated from fol.g4 by ANTLR 4.7.

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

// EnterDisjunction is called when production disjunction is entered.
func (s *BasefolListener) EnterDisjunction(ctx *DisjunctionContext) {}

// ExitDisjunction is called when production disjunction is exited.
func (s *BasefolListener) ExitDisjunction(ctx *DisjunctionContext) {}

// EnterConjunction is called when production conjunction is entered.
func (s *BasefolListener) EnterConjunction(ctx *ConjunctionContext) {}

// ExitConjunction is called when production conjunction is exited.
func (s *BasefolListener) ExitConjunction(ctx *ConjunctionContext) {}

// EnterNegation is called when production negation is entered.
func (s *BasefolListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production negation is exited.
func (s *BasefolListener) ExitNegation(ctx *NegationContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasefolListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasefolListener) ExitPredicate(ctx *PredicateContext) {}

// EnterPredicateTuple is called when production predicateTuple is entered.
func (s *BasefolListener) EnterPredicateTuple(ctx *PredicateTupleContext) {}

// ExitPredicateTuple is called when production predicateTuple is exited.
func (s *BasefolListener) ExitPredicateTuple(ctx *PredicateTupleContext) {}

// EnterTerm is called when production term is entered.
func (s *BasefolListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasefolListener) ExitTerm(ctx *TermContext) {}

// EnterFunction is called when production function is entered.
func (s *BasefolListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasefolListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionTuple is called when production functionTuple is entered.
func (s *BasefolListener) EnterFunctionTuple(ctx *FunctionTupleContext) {}

// ExitFunctionTuple is called when production functionTuple is exited.
func (s *BasefolListener) ExitFunctionTuple(ctx *FunctionTupleContext) {}
