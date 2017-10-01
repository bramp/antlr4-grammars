// Generated from STGParser.g4 by ANTLR 4.7.

package stringtemplate // STGParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSTGParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSTGParserVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitDelimiters(ctx *DelimitersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitImports(ctx *ImportsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitFormalArgs(ctx *FormalArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitFormalArg(ctx *FormalArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitDict(ctx *DictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitDictPairs(ctx *DictPairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitKeyValuePair(ctx *KeyValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitDefaultValuePair(ctx *DefaultValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTGParserVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}
