// Generated from tnt.g4 by ANTLR 4.7.

package tnt // tnt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetntListener is a complete listener for a parse tree produced by tntParser.
type BasetntListener struct{}

var _ tntListener = &BasetntListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetntListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetntListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetntListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetntListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEquation is called when production equation is entered.
func (s *BasetntListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BasetntListener) ExitEquation(ctx *EquationContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasetntListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasetntListener) ExitAtom(ctx *AtomContext) {}

// EnterNumber is called when production number is entered.
func (s *BasetntListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasetntListener) ExitNumber(ctx *NumberContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasetntListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasetntListener) ExitVariable(ctx *VariableContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasetntListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasetntListener) ExitExpression(ctx *ExpressionContext) {}

// EnterForevery is called when production forevery is entered.
func (s *BasetntListener) EnterForevery(ctx *ForeveryContext) {}

// ExitForevery is called when production forevery is exited.
func (s *BasetntListener) ExitForevery(ctx *ForeveryContext) {}

// EnterExists is called when production exists is entered.
func (s *BasetntListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BasetntListener) ExitExists(ctx *ExistsContext) {}
