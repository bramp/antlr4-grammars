// Generated from ZOperatorParser.g4 by ANTLR 4.7.

package z // ZOperatorParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseZOperatorParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZOperatorParserVisitor) VisitSpecification(ctx *SpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitInheritingSection(ctx *InheritingSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitBaseSection(ctx *BaseSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitNONOperatorTemplateParagraph(ctx *NONOperatorTemplateParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitFormals(ctx *FormalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitRelationOperatorTemplate(ctx *RelationOperatorTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitFunctionOperatorTemplate(ctx *FunctionOperatorTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitGenericOperatorTemplate(ctx *GenericOperatorTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitCategoryTemplate(ctx *CategoryTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitPrec(ctx *PrecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitAssoc(ctx *AssocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitPrefixTemplate(ctx *PrefixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitPostfixTemplate(ctx *PostfixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitInfixTemplate(ctx *InfixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitNofixTemplate(ctx *NofixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitOptArgName(ctx *OptArgNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitOptListName(ctx *OptListNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitArgName(ctx *ArgNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitListName(ctx *ListNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitPrefixName(ctx *PrefixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitPostfixName(ctx *PostfixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitInfixName(ctx *InfixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZOperatorParserVisitor) VisitNofixName(ctx *NofixNameContext) interface{} {
	return v.VisitChildren(ctx)
}
