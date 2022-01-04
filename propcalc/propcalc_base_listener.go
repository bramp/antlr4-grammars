// Code generated from propcalc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package propcalc // propcalc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepropcalcListener is a complete listener for a parse tree produced by propcalcParser.
type BasepropcalcListener struct{}

var _ propcalcListener = &BasepropcalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepropcalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepropcalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepropcalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepropcalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProposition is called when production proposition is entered.
func (s *BasepropcalcListener) EnterProposition(ctx *PropositionContext) {}

// ExitProposition is called when production proposition is exited.
func (s *BasepropcalcListener) ExitProposition(ctx *PropositionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasepropcalcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasepropcalcListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelExpression is called when production relExpression is entered.
func (s *BasepropcalcListener) EnterRelExpression(ctx *RelExpressionContext) {}

// ExitRelExpression is called when production relExpression is exited.
func (s *BasepropcalcListener) ExitRelExpression(ctx *RelExpressionContext) {}

// EnterAtoms is called when production atoms is entered.
func (s *BasepropcalcListener) EnterAtoms(ctx *AtomsContext) {}

// ExitAtoms is called when production atoms is exited.
func (s *BasepropcalcListener) ExitAtoms(ctx *AtomsContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasepropcalcListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasepropcalcListener) ExitAtom(ctx *AtomContext) {}

// EnterEquiv is called when production equiv is entered.
func (s *BasepropcalcListener) EnterEquiv(ctx *EquivContext) {}

// ExitEquiv is called when production equiv is exited.
func (s *BasepropcalcListener) ExitEquiv(ctx *EquivContext) {}

// EnterImplies is called when production implies is entered.
func (s *BasepropcalcListener) EnterImplies(ctx *ImpliesContext) {}

// ExitImplies is called when production implies is exited.
func (s *BasepropcalcListener) ExitImplies(ctx *ImpliesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasepropcalcListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasepropcalcListener) ExitVariable(ctx *VariableContext) {}
