// Generated from ZOperatorParser.g4 by ANTLR 4.7.

package z // ZOperatorParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ZOperatorParser.
type ZOperatorParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZOperatorParser#specification.
	VisitSpecification(ctx *SpecificationContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#InheritingSection.
	VisitInheritingSection(ctx *InheritingSectionContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#BaseSection.
	VisitBaseSection(ctx *BaseSectionContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#OperatorTemplateParagraph.
	VisitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#AxiomaticDescriptionParagraph.
	VisitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#SchemaDefinitionParagraph.
	VisitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#NONOperatorTemplateParagraph.
	VisitNONOperatorTemplateParagraph(ctx *NONOperatorTemplateParagraphContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#formals.
	VisitFormals(ctx *FormalsContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#RelationOperatorTemplate.
	VisitRelationOperatorTemplate(ctx *RelationOperatorTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#FunctionOperatorTemplate.
	VisitFunctionOperatorTemplate(ctx *FunctionOperatorTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#GenericOperatorTemplate.
	VisitGenericOperatorTemplate(ctx *GenericOperatorTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#categoryTemplate.
	VisitCategoryTemplate(ctx *CategoryTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#prec.
	VisitPrec(ctx *PrecContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#assoc.
	VisitAssoc(ctx *AssocContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#prefixTemplate.
	VisitPrefixTemplate(ctx *PrefixTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#postfixTemplate.
	VisitPostfixTemplate(ctx *PostfixTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#infixTemplate.
	VisitInfixTemplate(ctx *InfixTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#nofixTemplate.
	VisitNofixTemplate(ctx *NofixTemplateContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#optArgName.
	VisitOptArgName(ctx *OptArgNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#optListName.
	VisitOptListName(ctx *OptListNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#argName.
	VisitArgName(ctx *ArgNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#listName.
	VisitListName(ctx *ListNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#prefixName.
	VisitPrefixName(ctx *PrefixNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#postfixName.
	VisitPostfixName(ctx *PostfixNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#infixName.
	VisitInfixName(ctx *InfixNameContext) interface{}

	// Visit a parse tree produced by ZOperatorParser#nofixName.
	VisitNofixName(ctx *NofixNameContext) interface{}
}
