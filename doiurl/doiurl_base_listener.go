// Code generated from doiurl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package doiurl // doiurl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedoiurlListener is a complete listener for a parse tree produced by doiurlParser.
type BasedoiurlListener struct{}

var _ doiurlListener = &BasedoiurlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedoiurlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedoiurlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedoiurlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedoiurlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDoiuri is called when production doiuri is entered.
func (s *BasedoiurlListener) EnterDoiuri(ctx *DoiuriContext) {}

// ExitDoiuri is called when production doiuri is exited.
func (s *BasedoiurlListener) ExitDoiuri(ctx *DoiuriContext) {}

// EnterScheme is called when production scheme is entered.
func (s *BasedoiurlListener) EnterScheme(ctx *SchemeContext) {}

// ExitScheme is called when production scheme is exited.
func (s *BasedoiurlListener) ExitScheme(ctx *SchemeContext) {}

// EnterEncodeddoi is called when production encodeddoi is entered.
func (s *BasedoiurlListener) EnterEncodeddoi(ctx *EncodeddoiContext) {}

// ExitEncodeddoi is called when production encodeddoi is exited.
func (s *BasedoiurlListener) ExitEncodeddoi(ctx *EncodeddoiContext) {}

// EnterPrefix_ is called when production prefix_ is entered.
func (s *BasedoiurlListener) EnterPrefix_(ctx *Prefix_Context) {}

// ExitPrefix_ is called when production prefix_ is exited.
func (s *BasedoiurlListener) ExitPrefix_(ctx *Prefix_Context) {}

// EnterSuffix is called when production suffix is entered.
func (s *BasedoiurlListener) EnterSuffix(ctx *SuffixContext) {}

// ExitSuffix is called when production suffix is exited.
func (s *BasedoiurlListener) ExitSuffix(ctx *SuffixContext) {}

// EnterSegment is called when production segment is entered.
func (s *BasedoiurlListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BasedoiurlListener) ExitSegment(ctx *SegmentContext) {}

// EnterQuery is called when production query is entered.
func (s *BasedoiurlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasedoiurlListener) ExitQuery(ctx *QueryContext) {}

// EnterFragment_ is called when production fragment_ is entered.
func (s *BasedoiurlListener) EnterFragment_(ctx *Fragment_Context) {}

// ExitFragment_ is called when production fragment_ is exited.
func (s *BasedoiurlListener) ExitFragment_(ctx *Fragment_Context) {}
