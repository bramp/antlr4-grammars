// Generated from XMLParser.g4 by ANTLR 4.7.

package xml // XMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseXMLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseXMLParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitProlog(ctx *PrologContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitReference(ctx *ReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitChardata(ctx *ChardataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitMisc(ctx *MiscContext) interface{} {
	return v.VisitChildren(ctx)
}
