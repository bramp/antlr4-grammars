// Generated from HTMLParser.g4 by ANTLR 4.7.

package html // HTMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseHTMLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHTMLParserVisitor) VisitHtmlDocument(ctx *HtmlDocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlElements(ctx *HtmlElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlElement(ctx *HtmlElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlContent(ctx *HtmlContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlAttribute(ctx *HtmlAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlAttributeName(ctx *HtmlAttributeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlAttributeValue(ctx *HtmlAttributeValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlTagName(ctx *HtmlTagNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlChardata(ctx *HtmlChardataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlMisc(ctx *HtmlMiscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitHtmlComment(ctx *HtmlCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitXhtmlCDATA(ctx *XhtmlCDATAContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitDtd(ctx *DtdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitXml(ctx *XmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitScriptlet(ctx *ScriptletContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHTMLParserVisitor) VisitStyle(ctx *StyleContext) interface{} {
	return v.VisitChildren(ctx)
}
