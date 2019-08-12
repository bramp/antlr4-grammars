// Code generated from arithmetic.g4 by ANTLR 4.7.2. DO NOT EDIT.

package arithmetic // arithmetic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasearithmeticListener is a complete listener for a parse tree produced by arithmeticParser.
type BasearithmeticListener struct{}

var _ arithmeticListener = &BasearithmeticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasearithmeticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasearithmeticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasearithmeticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasearithmeticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BasearithmeticListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasearithmeticListener) ExitFile(ctx *FileContext) {}

// EnterEquation is called when production equation is entered.
func (s *BasearithmeticListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BasearithmeticListener) ExitEquation(ctx *EquationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasearithmeticListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasearithmeticListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasearithmeticListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasearithmeticListener) ExitAtom(ctx *AtomContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BasearithmeticListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BasearithmeticListener) ExitScientific(ctx *ScientificContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasearithmeticListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasearithmeticListener) ExitVariable(ctx *VariableContext) {}

// EnterRelop is called when production relop is entered.
func (s *BasearithmeticListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BasearithmeticListener) ExitRelop(ctx *RelopContext) {}
