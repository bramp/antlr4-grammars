// Code generated from pii.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pii // pii
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepiiListener is a complete listener for a parse tree produced by piiParser.
type BasepiiListener struct{}

var _ piiListener = &BasepiiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepiiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepiiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepiiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepiiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPii is called when production pii is entered.
func (s *BasepiiListener) EnterPii(ctx *PiiContext) {}

// ExitPii is called when production pii is exited.
func (s *BasepiiListener) ExitPii(ctx *PiiContext) {}

// EnterIssn is called when production issn is entered.
func (s *BasepiiListener) EnterIssn(ctx *IssnContext) {}

// ExitIssn is called when production issn is exited.
func (s *BasepiiListener) ExitIssn(ctx *IssnContext) {}

// EnterIsbn is called when production isbn is entered.
func (s *BasepiiListener) EnterIsbn(ctx *IsbnContext) {}

// ExitIsbn is called when production isbn is exited.
func (s *BasepiiListener) ExitIsbn(ctx *IsbnContext) {}

// EnterYear is called when production year is entered.
func (s *BasepiiListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BasepiiListener) ExitYear(ctx *YearContext) {}

// EnterItem is called when production item is entered.
func (s *BasepiiListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasepiiListener) ExitItem(ctx *ItemContext) {}

// EnterCheck is called when production check is entered.
func (s *BasepiiListener) EnterCheck(ctx *CheckContext) {}

// ExitCheck is called when production check is exited.
func (s *BasepiiListener) ExitCheck(ctx *CheckContext) {}
