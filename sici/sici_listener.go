// Code generated from sici.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sici // sici
import "github.com/antlr/antlr4/runtime/Go/antlr"

// siciListener is a complete listener for a parse tree produced by siciParser.
type siciListener interface {
	antlr.ParseTreeListener

	// EnterSici is called when entering the sici production.
	EnterSici(c *SiciContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterIssn is called when entering the issn production.
	EnterIssn(c *IssnContext)

	// EnterChronology is called when entering the chronology production.
	EnterChronology(c *ChronologyContext)

	// EnterEnumeration is called when entering the enumeration production.
	EnterEnumeration(c *EnumerationContext)

	// EnterContribution is called when entering the contribution production.
	EnterContribution(c *ContributionContext)

	// EnterLocation is called when entering the location production.
	EnterLocation(c *LocationContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterControl is called when entering the control production.
	EnterControl(c *ControlContext)

	// EnterCsi is called when entering the csi production.
	EnterCsi(c *CsiContext)

	// EnterDpi is called when entering the dpi production.
	EnterDpi(c *DpiContext)

	// EnterMfi is called when entering the mfi production.
	EnterMfi(c *MfiContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterCheck is called when entering the check production.
	EnterCheck(c *CheckContext)

	// ExitSici is called when exiting the sici production.
	ExitSici(c *SiciContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitIssn is called when exiting the issn production.
	ExitIssn(c *IssnContext)

	// ExitChronology is called when exiting the chronology production.
	ExitChronology(c *ChronologyContext)

	// ExitEnumeration is called when exiting the enumeration production.
	ExitEnumeration(c *EnumerationContext)

	// ExitContribution is called when exiting the contribution production.
	ExitContribution(c *ContributionContext)

	// ExitLocation is called when exiting the location production.
	ExitLocation(c *LocationContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitControl is called when exiting the control production.
	ExitControl(c *ControlContext)

	// ExitCsi is called when exiting the csi production.
	ExitCsi(c *CsiContext)

	// ExitDpi is called when exiting the dpi production.
	ExitDpi(c *DpiContext)

	// ExitMfi is called when exiting the mfi production.
	ExitMfi(c *MfiContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitCheck is called when exiting the check production.
	ExitCheck(c *CheckContext)
}
