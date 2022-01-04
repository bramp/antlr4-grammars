// Code generated from sici.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sici // sici
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesiciListener is a complete listener for a parse tree produced by siciParser.
type BasesiciListener struct{}

var _ siciListener = &BasesiciListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesiciListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesiciListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesiciListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesiciListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSici is called when production sici is entered.
func (s *BasesiciListener) EnterSici(ctx *SiciContext) {}

// ExitSici is called when production sici is exited.
func (s *BasesiciListener) ExitSici(ctx *SiciContext) {}

// EnterItem is called when production item is entered.
func (s *BasesiciListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasesiciListener) ExitItem(ctx *ItemContext) {}

// EnterIssn is called when production issn is entered.
func (s *BasesiciListener) EnterIssn(ctx *IssnContext) {}

// ExitIssn is called when production issn is exited.
func (s *BasesiciListener) ExitIssn(ctx *IssnContext) {}

// EnterChronology is called when production chronology is entered.
func (s *BasesiciListener) EnterChronology(ctx *ChronologyContext) {}

// ExitChronology is called when production chronology is exited.
func (s *BasesiciListener) ExitChronology(ctx *ChronologyContext) {}

// EnterEnumeration is called when production enumeration is entered.
func (s *BasesiciListener) EnterEnumeration(ctx *EnumerationContext) {}

// ExitEnumeration is called when production enumeration is exited.
func (s *BasesiciListener) ExitEnumeration(ctx *EnumerationContext) {}

// EnterContribution is called when production contribution is entered.
func (s *BasesiciListener) EnterContribution(ctx *ContributionContext) {}

// ExitContribution is called when production contribution is exited.
func (s *BasesiciListener) ExitContribution(ctx *ContributionContext) {}

// EnterLocation is called when production location is entered.
func (s *BasesiciListener) EnterLocation(ctx *LocationContext) {}

// ExitLocation is called when production location is exited.
func (s *BasesiciListener) ExitLocation(ctx *LocationContext) {}

// EnterTitle is called when production title is entered.
func (s *BasesiciListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BasesiciListener) ExitTitle(ctx *TitleContext) {}

// EnterControl is called when production control is entered.
func (s *BasesiciListener) EnterControl(ctx *ControlContext) {}

// ExitControl is called when production control is exited.
func (s *BasesiciListener) ExitControl(ctx *ControlContext) {}

// EnterCsi is called when production csi is entered.
func (s *BasesiciListener) EnterCsi(ctx *CsiContext) {}

// ExitCsi is called when production csi is exited.
func (s *BasesiciListener) ExitCsi(ctx *CsiContext) {}

// EnterDpi is called when production dpi is entered.
func (s *BasesiciListener) EnterDpi(ctx *DpiContext) {}

// ExitDpi is called when production dpi is exited.
func (s *BasesiciListener) ExitDpi(ctx *DpiContext) {}

// EnterMfi is called when production mfi is entered.
func (s *BasesiciListener) EnterMfi(ctx *MfiContext) {}

// ExitMfi is called when production mfi is exited.
func (s *BasesiciListener) ExitMfi(ctx *MfiContext) {}

// EnterVersion is called when production version is entered.
func (s *BasesiciListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BasesiciListener) ExitVersion(ctx *VersionContext) {}

// EnterCheck is called when production check is entered.
func (s *BasesiciListener) EnterCheck(ctx *CheckContext) {}

// ExitCheck is called when production check is exited.
func (s *BasesiciListener) ExitCheck(ctx *CheckContext) {}
