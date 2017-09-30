// Generated from DiceNotation.g4 by ANTLR 4.7.

package dice // DiceNotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiceNotationListener is a complete listener for a parse tree produced by DiceNotationParser.
type BaseDiceNotationListener struct{}

var _ DiceNotationListener = &BaseDiceNotationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiceNotationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiceNotationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiceNotationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiceNotationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseDiceNotationListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseDiceNotationListener) ExitParse(ctx *ParseContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseDiceNotationListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseDiceNotationListener) ExitFunction(ctx *FunctionContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BaseDiceNotationListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BaseDiceNotationListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterDice is called when production dice is entered.
func (s *BaseDiceNotationListener) EnterDice(ctx *DiceContext) {}

// ExitDice is called when production dice is exited.
func (s *BaseDiceNotationListener) ExitDice(ctx *DiceContext) {}
