// Generated from ZOperatorParser.g4 by ANTLR 4.7.

package z // ZOperatorParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZOperatorParserListener is a complete listener for a parse tree produced by ZOperatorParser.
type ZOperatorParserListener interface {
	antlr.ParseTreeListener

	// EnterSpecification is called when entering the specification production.
	EnterSpecification(c *SpecificationContext)

	// EnterInheritingSection is called when entering the InheritingSection production.
	EnterInheritingSection(c *InheritingSectionContext)

	// EnterBaseSection is called when entering the BaseSection production.
	EnterBaseSection(c *BaseSectionContext)

	// EnterOperatorTemplateParagraph is called when entering the OperatorTemplateParagraph production.
	EnterOperatorTemplateParagraph(c *OperatorTemplateParagraphContext)

	// EnterAxiomaticDescriptionParagraph is called when entering the AxiomaticDescriptionParagraph production.
	EnterAxiomaticDescriptionParagraph(c *AxiomaticDescriptionParagraphContext)

	// EnterSchemaDefinitionParagraph is called when entering the SchemaDefinitionParagraph production.
	EnterSchemaDefinitionParagraph(c *SchemaDefinitionParagraphContext)

	// EnterNONOperatorTemplateParagraph is called when entering the NONOperatorTemplateParagraph production.
	EnterNONOperatorTemplateParagraph(c *NONOperatorTemplateParagraphContext)

	// EnterFormals is called when entering the formals production.
	EnterFormals(c *FormalsContext)

	// EnterRelationOperatorTemplate is called when entering the RelationOperatorTemplate production.
	EnterRelationOperatorTemplate(c *RelationOperatorTemplateContext)

	// EnterFunctionOperatorTemplate is called when entering the FunctionOperatorTemplate production.
	EnterFunctionOperatorTemplate(c *FunctionOperatorTemplateContext)

	// EnterGenericOperatorTemplate is called when entering the GenericOperatorTemplate production.
	EnterGenericOperatorTemplate(c *GenericOperatorTemplateContext)

	// EnterCategoryTemplate is called when entering the categoryTemplate production.
	EnterCategoryTemplate(c *CategoryTemplateContext)

	// EnterPrec is called when entering the prec production.
	EnterPrec(c *PrecContext)

	// EnterAssoc is called when entering the assoc production.
	EnterAssoc(c *AssocContext)

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterPrefixTemplate is called when entering the prefixTemplate production.
	EnterPrefixTemplate(c *PrefixTemplateContext)

	// EnterPostfixTemplate is called when entering the postfixTemplate production.
	EnterPostfixTemplate(c *PostfixTemplateContext)

	// EnterInfixTemplate is called when entering the infixTemplate production.
	EnterInfixTemplate(c *InfixTemplateContext)

	// EnterNofixTemplate is called when entering the nofixTemplate production.
	EnterNofixTemplate(c *NofixTemplateContext)

	// EnterOptArgName is called when entering the optArgName production.
	EnterOptArgName(c *OptArgNameContext)

	// EnterOptListName is called when entering the optListName production.
	EnterOptListName(c *OptListNameContext)

	// EnterArgName is called when entering the argName production.
	EnterArgName(c *ArgNameContext)

	// EnterListName is called when entering the listName production.
	EnterListName(c *ListNameContext)

	// EnterPrefixName is called when entering the prefixName production.
	EnterPrefixName(c *PrefixNameContext)

	// EnterPostfixName is called when entering the postfixName production.
	EnterPostfixName(c *PostfixNameContext)

	// EnterInfixName is called when entering the infixName production.
	EnterInfixName(c *InfixNameContext)

	// EnterNofixName is called when entering the nofixName production.
	EnterNofixName(c *NofixNameContext)

	// ExitSpecification is called when exiting the specification production.
	ExitSpecification(c *SpecificationContext)

	// ExitInheritingSection is called when exiting the InheritingSection production.
	ExitInheritingSection(c *InheritingSectionContext)

	// ExitBaseSection is called when exiting the BaseSection production.
	ExitBaseSection(c *BaseSectionContext)

	// ExitOperatorTemplateParagraph is called when exiting the OperatorTemplateParagraph production.
	ExitOperatorTemplateParagraph(c *OperatorTemplateParagraphContext)

	// ExitAxiomaticDescriptionParagraph is called when exiting the AxiomaticDescriptionParagraph production.
	ExitAxiomaticDescriptionParagraph(c *AxiomaticDescriptionParagraphContext)

	// ExitSchemaDefinitionParagraph is called when exiting the SchemaDefinitionParagraph production.
	ExitSchemaDefinitionParagraph(c *SchemaDefinitionParagraphContext)

	// ExitNONOperatorTemplateParagraph is called when exiting the NONOperatorTemplateParagraph production.
	ExitNONOperatorTemplateParagraph(c *NONOperatorTemplateParagraphContext)

	// ExitFormals is called when exiting the formals production.
	ExitFormals(c *FormalsContext)

	// ExitRelationOperatorTemplate is called when exiting the RelationOperatorTemplate production.
	ExitRelationOperatorTemplate(c *RelationOperatorTemplateContext)

	// ExitFunctionOperatorTemplate is called when exiting the FunctionOperatorTemplate production.
	ExitFunctionOperatorTemplate(c *FunctionOperatorTemplateContext)

	// ExitGenericOperatorTemplate is called when exiting the GenericOperatorTemplate production.
	ExitGenericOperatorTemplate(c *GenericOperatorTemplateContext)

	// ExitCategoryTemplate is called when exiting the categoryTemplate production.
	ExitCategoryTemplate(c *CategoryTemplateContext)

	// ExitPrec is called when exiting the prec production.
	ExitPrec(c *PrecContext)

	// ExitAssoc is called when exiting the assoc production.
	ExitAssoc(c *AssocContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitPrefixTemplate is called when exiting the prefixTemplate production.
	ExitPrefixTemplate(c *PrefixTemplateContext)

	// ExitPostfixTemplate is called when exiting the postfixTemplate production.
	ExitPostfixTemplate(c *PostfixTemplateContext)

	// ExitInfixTemplate is called when exiting the infixTemplate production.
	ExitInfixTemplate(c *InfixTemplateContext)

	// ExitNofixTemplate is called when exiting the nofixTemplate production.
	ExitNofixTemplate(c *NofixTemplateContext)

	// ExitOptArgName is called when exiting the optArgName production.
	ExitOptArgName(c *OptArgNameContext)

	// ExitOptListName is called when exiting the optListName production.
	ExitOptListName(c *OptListNameContext)

	// ExitArgName is called when exiting the argName production.
	ExitArgName(c *ArgNameContext)

	// ExitListName is called when exiting the listName production.
	ExitListName(c *ListNameContext)

	// ExitPrefixName is called when exiting the prefixName production.
	ExitPrefixName(c *PrefixNameContext)

	// ExitPostfixName is called when exiting the postfixName production.
	ExitPostfixName(c *PostfixNameContext)

	// ExitInfixName is called when exiting the infixName production.
	ExitInfixName(c *InfixNameContext)

	// ExitNofixName is called when exiting the nofixName production.
	ExitNofixName(c *NofixNameContext)
}
