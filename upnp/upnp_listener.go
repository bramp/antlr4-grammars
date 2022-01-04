// Code generated from Upnp.g4 by ANTLR 4.9.3. DO NOT EDIT.

package upnp // Upnp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// UpnpListener is a complete listener for a parse tree produced by UpnpParser.
type UpnpListener interface {
	antlr.ParseTreeListener

	// EnterSearchCrit is called when entering the searchCrit production.
	EnterSearchCrit(c *SearchCritContext)

	// EnterSearchExp is called when entering the searchExp production.
	EnterSearchExp(c *SearchExpContext)

	// EnterRelExp is called when entering the relExp production.
	EnterRelExp(c *RelExpContext)

	// EnterQuotedVal is called when entering the quotedVal production.
	EnterQuotedVal(c *QuotedValContext)

	// EnterEscapedQuote is called when entering the escapedQuote production.
	EnterEscapedQuote(c *EscapedQuoteContext)

	// ExitSearchCrit is called when exiting the searchCrit production.
	ExitSearchCrit(c *SearchCritContext)

	// ExitSearchExp is called when exiting the searchExp production.
	ExitSearchExp(c *SearchExpContext)

	// ExitRelExp is called when exiting the relExp production.
	ExitRelExp(c *RelExpContext)

	// ExitQuotedVal is called when exiting the quotedVal production.
	ExitQuotedVal(c *QuotedValContext)

	// ExitEscapedQuote is called when exiting the escapedQuote production.
	ExitEscapedQuote(c *EscapedQuoteContext)
}
