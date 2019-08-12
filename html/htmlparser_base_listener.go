// Code generated from HTMLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package html // HTMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHTMLParserListener is a complete listener for a parse tree produced by HTMLParser.
type BaseHTMLParserListener struct{}

var _ HTMLParserListener = &BaseHTMLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHTMLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHTMLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHTMLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHTMLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHtmlDocument is called when production htmlDocument is entered.
func (s *BaseHTMLParserListener) EnterHtmlDocument(ctx *HtmlDocumentContext) {}

// ExitHtmlDocument is called when production htmlDocument is exited.
func (s *BaseHTMLParserListener) ExitHtmlDocument(ctx *HtmlDocumentContext) {}

// EnterHtmlElements is called when production htmlElements is entered.
func (s *BaseHTMLParserListener) EnterHtmlElements(ctx *HtmlElementsContext) {}

// ExitHtmlElements is called when production htmlElements is exited.
func (s *BaseHTMLParserListener) ExitHtmlElements(ctx *HtmlElementsContext) {}

// EnterHtmlElement is called when production htmlElement is entered.
func (s *BaseHTMLParserListener) EnterHtmlElement(ctx *HtmlElementContext) {}

// ExitHtmlElement is called when production htmlElement is exited.
func (s *BaseHTMLParserListener) ExitHtmlElement(ctx *HtmlElementContext) {}

// EnterHtmlContent is called when production htmlContent is entered.
func (s *BaseHTMLParserListener) EnterHtmlContent(ctx *HtmlContentContext) {}

// ExitHtmlContent is called when production htmlContent is exited.
func (s *BaseHTMLParserListener) ExitHtmlContent(ctx *HtmlContentContext) {}

// EnterHtmlAttribute is called when production htmlAttribute is entered.
func (s *BaseHTMLParserListener) EnterHtmlAttribute(ctx *HtmlAttributeContext) {}

// ExitHtmlAttribute is called when production htmlAttribute is exited.
func (s *BaseHTMLParserListener) ExitHtmlAttribute(ctx *HtmlAttributeContext) {}

// EnterHtmlAttributeName is called when production htmlAttributeName is entered.
func (s *BaseHTMLParserListener) EnterHtmlAttributeName(ctx *HtmlAttributeNameContext) {}

// ExitHtmlAttributeName is called when production htmlAttributeName is exited.
func (s *BaseHTMLParserListener) ExitHtmlAttributeName(ctx *HtmlAttributeNameContext) {}

// EnterHtmlAttributeValue is called when production htmlAttributeValue is entered.
func (s *BaseHTMLParserListener) EnterHtmlAttributeValue(ctx *HtmlAttributeValueContext) {}

// ExitHtmlAttributeValue is called when production htmlAttributeValue is exited.
func (s *BaseHTMLParserListener) ExitHtmlAttributeValue(ctx *HtmlAttributeValueContext) {}

// EnterHtmlTagName is called when production htmlTagName is entered.
func (s *BaseHTMLParserListener) EnterHtmlTagName(ctx *HtmlTagNameContext) {}

// ExitHtmlTagName is called when production htmlTagName is exited.
func (s *BaseHTMLParserListener) ExitHtmlTagName(ctx *HtmlTagNameContext) {}

// EnterHtmlChardata is called when production htmlChardata is entered.
func (s *BaseHTMLParserListener) EnterHtmlChardata(ctx *HtmlChardataContext) {}

// ExitHtmlChardata is called when production htmlChardata is exited.
func (s *BaseHTMLParserListener) ExitHtmlChardata(ctx *HtmlChardataContext) {}

// EnterHtmlMisc is called when production htmlMisc is entered.
func (s *BaseHTMLParserListener) EnterHtmlMisc(ctx *HtmlMiscContext) {}

// ExitHtmlMisc is called when production htmlMisc is exited.
func (s *BaseHTMLParserListener) ExitHtmlMisc(ctx *HtmlMiscContext) {}

// EnterHtmlComment is called when production htmlComment is entered.
func (s *BaseHTMLParserListener) EnterHtmlComment(ctx *HtmlCommentContext) {}

// ExitHtmlComment is called when production htmlComment is exited.
func (s *BaseHTMLParserListener) ExitHtmlComment(ctx *HtmlCommentContext) {}

// EnterXhtmlCDATA is called when production xhtmlCDATA is entered.
func (s *BaseHTMLParserListener) EnterXhtmlCDATA(ctx *XhtmlCDATAContext) {}

// ExitXhtmlCDATA is called when production xhtmlCDATA is exited.
func (s *BaseHTMLParserListener) ExitXhtmlCDATA(ctx *XhtmlCDATAContext) {}

// EnterDtd is called when production dtd is entered.
func (s *BaseHTMLParserListener) EnterDtd(ctx *DtdContext) {}

// ExitDtd is called when production dtd is exited.
func (s *BaseHTMLParserListener) ExitDtd(ctx *DtdContext) {}

// EnterXml is called when production xml is entered.
func (s *BaseHTMLParserListener) EnterXml(ctx *XmlContext) {}

// ExitXml is called when production xml is exited.
func (s *BaseHTMLParserListener) ExitXml(ctx *XmlContext) {}

// EnterScriptlet is called when production scriptlet is entered.
func (s *BaseHTMLParserListener) EnterScriptlet(ctx *ScriptletContext) {}

// ExitScriptlet is called when production scriptlet is exited.
func (s *BaseHTMLParserListener) ExitScriptlet(ctx *ScriptletContext) {}

// EnterScript is called when production script is entered.
func (s *BaseHTMLParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseHTMLParserListener) ExitScript(ctx *ScriptContext) {}

// EnterStyle is called when production style is entered.
func (s *BaseHTMLParserListener) EnterStyle(ctx *StyleContext) {}

// ExitStyle is called when production style is exited.
func (s *BaseHTMLParserListener) ExitStyle(ctx *StyleContext) {}
