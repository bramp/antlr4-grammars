// Code generated from pii.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pii // pii
import "github.com/antlr/antlr4/runtime/Go/antlr"

// piiListener is a complete listener for a parse tree produced by piiParser.
type piiListener interface {
	antlr.ParseTreeListener

	// EnterPii is called when entering the pii production.
	EnterPii(c *PiiContext)

	// EnterIssn is called when entering the issn production.
	EnterIssn(c *IssnContext)

	// EnterIsbn is called when entering the isbn production.
	EnterIsbn(c *IsbnContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterCheck is called when entering the check production.
	EnterCheck(c *CheckContext)

	// ExitPii is called when exiting the pii production.
	ExitPii(c *PiiContext)

	// ExitIssn is called when exiting the issn production.
	ExitIssn(c *IssnContext)

	// ExitIsbn is called when exiting the isbn production.
	ExitIsbn(c *IsbnContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitCheck is called when exiting the check production.
	ExitCheck(c *CheckContext)
}
