// Code generated from calculator.g4 by ANTLR 4.9.3. DO NOT EDIT.

package calculator // calculator
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecalculatorListener is a complete listener for a parse tree produced by calculatorParser.
type BasecalculatorListener struct{}

var _ calculatorListener = &BasecalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEquation is called when production equation is entered.
func (s *BasecalculatorListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BasecalculatorListener) ExitEquation(ctx *EquationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasecalculatorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasecalculatorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BasecalculatorListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BasecalculatorListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterPowExpression is called when production powExpression is entered.
func (s *BasecalculatorListener) EnterPowExpression(ctx *PowExpressionContext) {}

// ExitPowExpression is called when production powExpression is exited.
func (s *BasecalculatorListener) ExitPowExpression(ctx *PowExpressionContext) {}

// EnterSignedAtom is called when production signedAtom is entered.
func (s *BasecalculatorListener) EnterSignedAtom(ctx *SignedAtomContext) {}

// ExitSignedAtom is called when production signedAtom is exited.
func (s *BasecalculatorListener) ExitSignedAtom(ctx *SignedAtomContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasecalculatorListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasecalculatorListener) ExitAtom(ctx *AtomContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BasecalculatorListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BasecalculatorListener) ExitScientific(ctx *ScientificContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasecalculatorListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasecalculatorListener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasecalculatorListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasecalculatorListener) ExitVariable(ctx *VariableContext) {}

// EnterFunc_ is called when production func_ is entered.
func (s *BasecalculatorListener) EnterFunc_(ctx *Func_Context) {}

// ExitFunc_ is called when production func_ is exited.
func (s *BasecalculatorListener) ExitFunc_(ctx *Func_Context) {}

// EnterFuncname is called when production funcname is entered.
func (s *BasecalculatorListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BasecalculatorListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterRelop is called when production relop is entered.
func (s *BasecalculatorListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BasecalculatorListener) ExitRelop(ctx *RelopContext) {}
