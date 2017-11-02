// Generated from classify.g4 by ANTLR 4.7.

package classify // classify
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseclassifyListener is a complete listener for a parse tree produced by classifyParser.
type BaseclassifyListener struct{}

var _ classifyListener = &BaseclassifyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseclassifyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseclassifyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseclassifyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseclassifyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCodepoint is called when production codepoint is entered.
func (s *BaseclassifyListener) EnterCodepoint(ctx *CodepointContext) {}

// ExitCodepoint is called when production codepoint is exited.
func (s *BaseclassifyListener) ExitCodepoint(ctx *CodepointContext) {}
