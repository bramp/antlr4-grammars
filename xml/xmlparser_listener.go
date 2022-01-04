// Code generated from XMLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package xml // XMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// XMLParserListener is a complete listener for a parse tree produced by XMLParser.
type XMLParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterProlog is called when entering the prolog production.
	EnterProlog(c *PrologContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterChardata is called when entering the chardata production.
	EnterChardata(c *ChardataContext)

	// EnterMisc is called when entering the misc production.
	EnterMisc(c *MiscContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitProlog is called when exiting the prolog production.
	ExitProlog(c *PrologContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitChardata is called when exiting the chardata production.
	ExitChardata(c *ChardataContext)

	// ExitMisc is called when exiting the misc production.
	ExitMisc(c *MiscContext)
}
