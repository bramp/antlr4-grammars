// Generated from postalcode.g4 by ANTLR 4.7.

package postalcode // postalcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepostalcodeListener is a complete listener for a parse tree produced by postalcodeParser.
type BasepostalcodeListener struct{}

var _ postalcodeListener = &BasepostalcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepostalcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepostalcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepostalcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepostalcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPostalcode is called when production postalcode is entered.
func (s *BasepostalcodeListener) EnterPostalcode(ctx *PostalcodeContext) {}

// ExitPostalcode is called when production postalcode is exited.
func (s *BasepostalcodeListener) ExitPostalcode(ctx *PostalcodeContext) {}
