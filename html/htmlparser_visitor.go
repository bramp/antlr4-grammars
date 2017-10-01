// Generated from HTMLParser.g4 by ANTLR 4.7.

package html // HTMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by HTMLParser.
type HTMLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HTMLParser#htmlDocument.
	VisitHtmlDocument(ctx *HtmlDocumentContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlElements.
	VisitHtmlElements(ctx *HtmlElementsContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlElement.
	VisitHtmlElement(ctx *HtmlElementContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlContent.
	VisitHtmlContent(ctx *HtmlContentContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlAttribute.
	VisitHtmlAttribute(ctx *HtmlAttributeContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlAttributeName.
	VisitHtmlAttributeName(ctx *HtmlAttributeNameContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlAttributeValue.
	VisitHtmlAttributeValue(ctx *HtmlAttributeValueContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlTagName.
	VisitHtmlTagName(ctx *HtmlTagNameContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlChardata.
	VisitHtmlChardata(ctx *HtmlChardataContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlMisc.
	VisitHtmlMisc(ctx *HtmlMiscContext) interface{}

	// Visit a parse tree produced by HTMLParser#htmlComment.
	VisitHtmlComment(ctx *HtmlCommentContext) interface{}

	// Visit a parse tree produced by HTMLParser#xhtmlCDATA.
	VisitXhtmlCDATA(ctx *XhtmlCDATAContext) interface{}

	// Visit a parse tree produced by HTMLParser#dtd.
	VisitDtd(ctx *DtdContext) interface{}

	// Visit a parse tree produced by HTMLParser#xml.
	VisitXml(ctx *XmlContext) interface{}

	// Visit a parse tree produced by HTMLParser#scriptlet.
	VisitScriptlet(ctx *ScriptletContext) interface{}

	// Visit a parse tree produced by HTMLParser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by HTMLParser#style.
	VisitStyle(ctx *StyleContext) interface{}
}
