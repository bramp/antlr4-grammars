// Generated from HTMLParser.g4 by ANTLR 4.7.

package html // HTMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// HTMLParserListener is a complete listener for a parse tree produced by HTMLParser.
type HTMLParserListener interface {
	antlr.ParseTreeListener

	// EnterHtmlDocument is called when entering the htmlDocument production.
	EnterHtmlDocument(c *HtmlDocumentContext)

	// EnterHtmlElements is called when entering the htmlElements production.
	EnterHtmlElements(c *HtmlElementsContext)

	// EnterHtmlElement is called when entering the htmlElement production.
	EnterHtmlElement(c *HtmlElementContext)

	// EnterHtmlContent is called when entering the htmlContent production.
	EnterHtmlContent(c *HtmlContentContext)

	// EnterHtmlAttribute is called when entering the htmlAttribute production.
	EnterHtmlAttribute(c *HtmlAttributeContext)

	// EnterHtmlAttributeName is called when entering the htmlAttributeName production.
	EnterHtmlAttributeName(c *HtmlAttributeNameContext)

	// EnterHtmlAttributeValue is called when entering the htmlAttributeValue production.
	EnterHtmlAttributeValue(c *HtmlAttributeValueContext)

	// EnterHtmlTagName is called when entering the htmlTagName production.
	EnterHtmlTagName(c *HtmlTagNameContext)

	// EnterHtmlChardata is called when entering the htmlChardata production.
	EnterHtmlChardata(c *HtmlChardataContext)

	// EnterHtmlMisc is called when entering the htmlMisc production.
	EnterHtmlMisc(c *HtmlMiscContext)

	// EnterHtmlComment is called when entering the htmlComment production.
	EnterHtmlComment(c *HtmlCommentContext)

	// EnterXhtmlCDATA is called when entering the xhtmlCDATA production.
	EnterXhtmlCDATA(c *XhtmlCDATAContext)

	// EnterDtd is called when entering the dtd production.
	EnterDtd(c *DtdContext)

	// EnterXml is called when entering the xml production.
	EnterXml(c *XmlContext)

	// EnterScriptlet is called when entering the scriptlet production.
	EnterScriptlet(c *ScriptletContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterStyle is called when entering the style production.
	EnterStyle(c *StyleContext)

	// ExitHtmlDocument is called when exiting the htmlDocument production.
	ExitHtmlDocument(c *HtmlDocumentContext)

	// ExitHtmlElements is called when exiting the htmlElements production.
	ExitHtmlElements(c *HtmlElementsContext)

	// ExitHtmlElement is called when exiting the htmlElement production.
	ExitHtmlElement(c *HtmlElementContext)

	// ExitHtmlContent is called when exiting the htmlContent production.
	ExitHtmlContent(c *HtmlContentContext)

	// ExitHtmlAttribute is called when exiting the htmlAttribute production.
	ExitHtmlAttribute(c *HtmlAttributeContext)

	// ExitHtmlAttributeName is called when exiting the htmlAttributeName production.
	ExitHtmlAttributeName(c *HtmlAttributeNameContext)

	// ExitHtmlAttributeValue is called when exiting the htmlAttributeValue production.
	ExitHtmlAttributeValue(c *HtmlAttributeValueContext)

	// ExitHtmlTagName is called when exiting the htmlTagName production.
	ExitHtmlTagName(c *HtmlTagNameContext)

	// ExitHtmlChardata is called when exiting the htmlChardata production.
	ExitHtmlChardata(c *HtmlChardataContext)

	// ExitHtmlMisc is called when exiting the htmlMisc production.
	ExitHtmlMisc(c *HtmlMiscContext)

	// ExitHtmlComment is called when exiting the htmlComment production.
	ExitHtmlComment(c *HtmlCommentContext)

	// ExitXhtmlCDATA is called when exiting the xhtmlCDATA production.
	ExitXhtmlCDATA(c *XhtmlCDATAContext)

	// ExitDtd is called when exiting the dtd production.
	ExitDtd(c *DtdContext)

	// ExitXml is called when exiting the xml production.
	ExitXml(c *XmlContext)

	// ExitScriptlet is called when exiting the scriptlet production.
	ExitScriptlet(c *ScriptletContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitStyle is called when exiting the style production.
	ExitStyle(c *StyleContext)
}
