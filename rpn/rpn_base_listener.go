// Generated from rpn.g4 by ANTLR 4.7.

package rpn // rpn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaserpnListener is a complete listener for a parse tree produced by rpnParser.
type BaserpnListener struct{}

var _ rpnListener = &BaserpnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaserpnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaserpnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaserpnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaserpnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaserpnListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaserpnListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaserpnListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaserpnListener) ExitTerm(ctx *TermContext) {}

// EnterOper is called when production oper is entered.
func (s *BaserpnListener) EnterOper(ctx *OperContext) {}

// ExitOper is called when production oper is exited.
func (s *BaserpnListener) ExitOper(ctx *OperContext) {}

// EnterSignedAtom is called when production signedAtom is entered.
func (s *BaserpnListener) EnterSignedAtom(ctx *SignedAtomContext) {}

// ExitSignedAtom is called when production signedAtom is exited.
func (s *BaserpnListener) ExitSignedAtom(ctx *SignedAtomContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaserpnListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaserpnListener) ExitVariable(ctx *VariableContext) {}

// EnterScientific is called when production scientific is entered.
func (s *BaserpnListener) EnterScientific(ctx *ScientificContext) {}

// ExitScientific is called when production scientific is exited.
func (s *BaserpnListener) ExitScientific(ctx *ScientificContext) {}
