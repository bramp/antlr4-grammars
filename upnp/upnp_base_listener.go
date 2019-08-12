// Code generated from Upnp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package upnp // Upnp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUpnpListener is a complete listener for a parse tree produced by UpnpParser.
type BaseUpnpListener struct{}

var _ UpnpListener = &BaseUpnpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUpnpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUpnpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUpnpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUpnpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSearchCrit is called when production searchCrit is entered.
func (s *BaseUpnpListener) EnterSearchCrit(ctx *SearchCritContext) {}

// ExitSearchCrit is called when production searchCrit is exited.
func (s *BaseUpnpListener) ExitSearchCrit(ctx *SearchCritContext) {}

// EnterSearchExp is called when production searchExp is entered.
func (s *BaseUpnpListener) EnterSearchExp(ctx *SearchExpContext) {}

// ExitSearchExp is called when production searchExp is exited.
func (s *BaseUpnpListener) ExitSearchExp(ctx *SearchExpContext) {}

// EnterRelExp is called when production relExp is entered.
func (s *BaseUpnpListener) EnterRelExp(ctx *RelExpContext) {}

// ExitRelExp is called when production relExp is exited.
func (s *BaseUpnpListener) ExitRelExp(ctx *RelExpContext) {}

// EnterQuotedVal is called when production quotedVal is entered.
func (s *BaseUpnpListener) EnterQuotedVal(ctx *QuotedValContext) {}

// ExitQuotedVal is called when production quotedVal is exited.
func (s *BaseUpnpListener) ExitQuotedVal(ctx *QuotedValContext) {}

// EnterEscapedQuote is called when production escapedQuote is entered.
func (s *BaseUpnpListener) EnterEscapedQuote(ctx *EscapedQuoteContext) {}

// ExitEscapedQuote is called when production escapedQuote is exited.
func (s *BaseUpnpListener) ExitEscapedQuote(ctx *EscapedQuoteContext) {}
