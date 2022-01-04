// Code generated from HTMLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package html // HTMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// HTMLParserListener is a complete listener for a parse tree produced by HTMLParser.
type HTMLParserListener interface {
	antlr.ParseTreeListener

	// EnterHtmlDocument is called when entering the htmlDocument production.
	EnterHtmlDocument(c *HtmlDocumentContext)

	// EnterScriptletOrSeaWs is called when entering the scriptletOrSeaWs production.
	EnterScriptletOrSeaWs(c *ScriptletOrSeaWsContext)

	// EnterHtmlElements is called when entering the htmlElements production.
	EnterHtmlElements(c *HtmlElementsContext)

	// EnterHtmlElement is called when entering the htmlElement production.
	EnterHtmlElement(c *HtmlElementContext)

	// EnterHtmlContent is called when entering the htmlContent production.
	EnterHtmlContent(c *HtmlContentContext)

	// EnterHtmlAttribute is called when entering the htmlAttribute production.
	EnterHtmlAttribute(c *HtmlAttributeContext)

	// EnterHtmlChardata is called when entering the htmlChardata production.
	EnterHtmlChardata(c *HtmlChardataContext)

	// EnterHtmlMisc is called when entering the htmlMisc production.
	EnterHtmlMisc(c *HtmlMiscContext)

	// EnterHtmlComment is called when entering the htmlComment production.
	EnterHtmlComment(c *HtmlCommentContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterStyle is called when entering the style production.
	EnterStyle(c *StyleContext)

	// ExitHtmlDocument is called when exiting the htmlDocument production.
	ExitHtmlDocument(c *HtmlDocumentContext)

	// ExitScriptletOrSeaWs is called when exiting the scriptletOrSeaWs production.
	ExitScriptletOrSeaWs(c *ScriptletOrSeaWsContext)

	// ExitHtmlElements is called when exiting the htmlElements production.
	ExitHtmlElements(c *HtmlElementsContext)

	// ExitHtmlElement is called when exiting the htmlElement production.
	ExitHtmlElement(c *HtmlElementContext)

	// ExitHtmlContent is called when exiting the htmlContent production.
	ExitHtmlContent(c *HtmlContentContext)

	// ExitHtmlAttribute is called when exiting the htmlAttribute production.
	ExitHtmlAttribute(c *HtmlAttributeContext)

	// ExitHtmlChardata is called when exiting the htmlChardata production.
	ExitHtmlChardata(c *HtmlChardataContext)

	// ExitHtmlMisc is called when exiting the htmlMisc production.
	ExitHtmlMisc(c *HtmlMiscContext)

	// ExitHtmlComment is called when exiting the htmlComment production.
	ExitHtmlComment(c *HtmlCommentContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitStyle is called when exiting the style production.
	ExitStyle(c *StyleContext)
}
