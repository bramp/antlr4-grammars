// Generated from ZOperatorParser.g4 by ANTLR 4.7.

package z // ZOperatorParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZOperatorParserListener is a complete listener for a parse tree produced by ZOperatorParser.
type BaseZOperatorParserListener struct{}

var _ ZOperatorParserListener = &BaseZOperatorParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZOperatorParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZOperatorParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZOperatorParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZOperatorParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSpecification is called when production specification is entered.
func (s *BaseZOperatorParserListener) EnterSpecification(ctx *SpecificationContext) {}

// ExitSpecification is called when production specification is exited.
func (s *BaseZOperatorParserListener) ExitSpecification(ctx *SpecificationContext) {}

// EnterInheritingSection is called when production InheritingSection is entered.
func (s *BaseZOperatorParserListener) EnterInheritingSection(ctx *InheritingSectionContext) {}

// ExitInheritingSection is called when production InheritingSection is exited.
func (s *BaseZOperatorParserListener) ExitInheritingSection(ctx *InheritingSectionContext) {}

// EnterBaseSection is called when production BaseSection is entered.
func (s *BaseZOperatorParserListener) EnterBaseSection(ctx *BaseSectionContext) {}

// ExitBaseSection is called when production BaseSection is exited.
func (s *BaseZOperatorParserListener) ExitBaseSection(ctx *BaseSectionContext) {}

// EnterOperatorTemplateParagraph is called when production OperatorTemplateParagraph is entered.
func (s *BaseZOperatorParserListener) EnterOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) {
}

// ExitOperatorTemplateParagraph is called when production OperatorTemplateParagraph is exited.
func (s *BaseZOperatorParserListener) ExitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) {
}

// EnterAxiomaticDescriptionParagraph is called when production AxiomaticDescriptionParagraph is entered.
func (s *BaseZOperatorParserListener) EnterAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) {
}

// ExitAxiomaticDescriptionParagraph is called when production AxiomaticDescriptionParagraph is exited.
func (s *BaseZOperatorParserListener) ExitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) {
}

// EnterSchemaDefinitionParagraph is called when production SchemaDefinitionParagraph is entered.
func (s *BaseZOperatorParserListener) EnterSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) {
}

// ExitSchemaDefinitionParagraph is called when production SchemaDefinitionParagraph is exited.
func (s *BaseZOperatorParserListener) ExitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) {
}

// EnterNONOperatorTemplateParagraph is called when production NONOperatorTemplateParagraph is entered.
func (s *BaseZOperatorParserListener) EnterNONOperatorTemplateParagraph(ctx *NONOperatorTemplateParagraphContext) {
}

// ExitNONOperatorTemplateParagraph is called when production NONOperatorTemplateParagraph is exited.
func (s *BaseZOperatorParserListener) ExitNONOperatorTemplateParagraph(ctx *NONOperatorTemplateParagraphContext) {
}

// EnterFormals is called when production formals is entered.
func (s *BaseZOperatorParserListener) EnterFormals(ctx *FormalsContext) {}

// ExitFormals is called when production formals is exited.
func (s *BaseZOperatorParserListener) ExitFormals(ctx *FormalsContext) {}

// EnterRelationOperatorTemplate is called when production RelationOperatorTemplate is entered.
func (s *BaseZOperatorParserListener) EnterRelationOperatorTemplate(ctx *RelationOperatorTemplateContext) {
}

// ExitRelationOperatorTemplate is called when production RelationOperatorTemplate is exited.
func (s *BaseZOperatorParserListener) ExitRelationOperatorTemplate(ctx *RelationOperatorTemplateContext) {
}

// EnterFunctionOperatorTemplate is called when production FunctionOperatorTemplate is entered.
func (s *BaseZOperatorParserListener) EnterFunctionOperatorTemplate(ctx *FunctionOperatorTemplateContext) {
}

// ExitFunctionOperatorTemplate is called when production FunctionOperatorTemplate is exited.
func (s *BaseZOperatorParserListener) ExitFunctionOperatorTemplate(ctx *FunctionOperatorTemplateContext) {
}

// EnterGenericOperatorTemplate is called when production GenericOperatorTemplate is entered.
func (s *BaseZOperatorParserListener) EnterGenericOperatorTemplate(ctx *GenericOperatorTemplateContext) {
}

// ExitGenericOperatorTemplate is called when production GenericOperatorTemplate is exited.
func (s *BaseZOperatorParserListener) ExitGenericOperatorTemplate(ctx *GenericOperatorTemplateContext) {
}

// EnterCategoryTemplate is called when production categoryTemplate is entered.
func (s *BaseZOperatorParserListener) EnterCategoryTemplate(ctx *CategoryTemplateContext) {}

// ExitCategoryTemplate is called when production categoryTemplate is exited.
func (s *BaseZOperatorParserListener) ExitCategoryTemplate(ctx *CategoryTemplateContext) {}

// EnterPrec is called when production prec is entered.
func (s *BaseZOperatorParserListener) EnterPrec(ctx *PrecContext) {}

// ExitPrec is called when production prec is exited.
func (s *BaseZOperatorParserListener) ExitPrec(ctx *PrecContext) {}

// EnterAssoc is called when production assoc is entered.
func (s *BaseZOperatorParserListener) EnterAssoc(ctx *AssocContext) {}

// ExitAssoc is called when production assoc is exited.
func (s *BaseZOperatorParserListener) ExitAssoc(ctx *AssocContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseZOperatorParserListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseZOperatorParserListener) ExitTemplate(ctx *TemplateContext) {}

// EnterPrefixTemplate is called when production prefixTemplate is entered.
func (s *BaseZOperatorParserListener) EnterPrefixTemplate(ctx *PrefixTemplateContext) {}

// ExitPrefixTemplate is called when production prefixTemplate is exited.
func (s *BaseZOperatorParserListener) ExitPrefixTemplate(ctx *PrefixTemplateContext) {}

// EnterPostfixTemplate is called when production postfixTemplate is entered.
func (s *BaseZOperatorParserListener) EnterPostfixTemplate(ctx *PostfixTemplateContext) {}

// ExitPostfixTemplate is called when production postfixTemplate is exited.
func (s *BaseZOperatorParserListener) ExitPostfixTemplate(ctx *PostfixTemplateContext) {}

// EnterInfixTemplate is called when production infixTemplate is entered.
func (s *BaseZOperatorParserListener) EnterInfixTemplate(ctx *InfixTemplateContext) {}

// ExitInfixTemplate is called when production infixTemplate is exited.
func (s *BaseZOperatorParserListener) ExitInfixTemplate(ctx *InfixTemplateContext) {}

// EnterNofixTemplate is called when production nofixTemplate is entered.
func (s *BaseZOperatorParserListener) EnterNofixTemplate(ctx *NofixTemplateContext) {}

// ExitNofixTemplate is called when production nofixTemplate is exited.
func (s *BaseZOperatorParserListener) ExitNofixTemplate(ctx *NofixTemplateContext) {}

// EnterOptArgName is called when production optArgName is entered.
func (s *BaseZOperatorParserListener) EnterOptArgName(ctx *OptArgNameContext) {}

// ExitOptArgName is called when production optArgName is exited.
func (s *BaseZOperatorParserListener) ExitOptArgName(ctx *OptArgNameContext) {}

// EnterOptListName is called when production optListName is entered.
func (s *BaseZOperatorParserListener) EnterOptListName(ctx *OptListNameContext) {}

// ExitOptListName is called when production optListName is exited.
func (s *BaseZOperatorParserListener) ExitOptListName(ctx *OptListNameContext) {}

// EnterArgName is called when production argName is entered.
func (s *BaseZOperatorParserListener) EnterArgName(ctx *ArgNameContext) {}

// ExitArgName is called when production argName is exited.
func (s *BaseZOperatorParserListener) ExitArgName(ctx *ArgNameContext) {}

// EnterListName is called when production listName is entered.
func (s *BaseZOperatorParserListener) EnterListName(ctx *ListNameContext) {}

// ExitListName is called when production listName is exited.
func (s *BaseZOperatorParserListener) ExitListName(ctx *ListNameContext) {}

// EnterPrefixName is called when production prefixName is entered.
func (s *BaseZOperatorParserListener) EnterPrefixName(ctx *PrefixNameContext) {}

// ExitPrefixName is called when production prefixName is exited.
func (s *BaseZOperatorParserListener) ExitPrefixName(ctx *PrefixNameContext) {}

// EnterPostfixName is called when production postfixName is entered.
func (s *BaseZOperatorParserListener) EnterPostfixName(ctx *PostfixNameContext) {}

// ExitPostfixName is called when production postfixName is exited.
func (s *BaseZOperatorParserListener) ExitPostfixName(ctx *PostfixNameContext) {}

// EnterInfixName is called when production infixName is entered.
func (s *BaseZOperatorParserListener) EnterInfixName(ctx *InfixNameContext) {}

// ExitInfixName is called when production infixName is exited.
func (s *BaseZOperatorParserListener) ExitInfixName(ctx *InfixNameContext) {}

// EnterNofixName is called when production nofixName is entered.
func (s *BaseZOperatorParserListener) EnterNofixName(ctx *NofixNameContext) {}

// ExitNofixName is called when production nofixName is exited.
func (s *BaseZOperatorParserListener) ExitNofixName(ctx *NofixNameContext) {}
