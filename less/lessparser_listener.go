// Code generated from LessParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package less // LessParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// LessParserListener is a complete listener for a parse tree produced by LessParser.
type LessParserListener interface {
	antlr.ParseTreeListener

	// EnterStylesheet is called when entering the stylesheet production.
	EnterStylesheet(c *StylesheetContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterCommandStatement is called when entering the commandStatement production.
	EnterCommandStatement(c *CommandStatementContext)

	// EnterMathCharacter is called when entering the mathCharacter production.
	EnterMathCharacter(c *MathCharacterContext)

	// EnterMathStatement is called when entering the mathStatement production.
	EnterMathStatement(c *MathStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterConditionStatement is called when entering the conditionStatement production.
	EnterConditionStatement(c *ConditionStatementContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterImportOption is called when entering the importOption production.
	EnterImportOption(c *ImportOptionContext)

	// EnterReferenceUrl is called when entering the referenceUrl production.
	EnterReferenceUrl(c *ReferenceUrlContext)

	// EnterMediaTypes is called when entering the mediaTypes production.
	EnterMediaTypes(c *MediaTypesContext)

	// EnterRuleset is called when entering the ruleset production.
	EnterRuleset(c *RulesetContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterMixinDefinition is called when entering the mixinDefinition production.
	EnterMixinDefinition(c *MixinDefinitionContext)

	// EnterMixinGuard is called when entering the mixinGuard production.
	EnterMixinGuard(c *MixinGuardContext)

	// EnterMixinDefinitionParam is called when entering the mixinDefinitionParam production.
	EnterMixinDefinitionParam(c *MixinDefinitionParamContext)

	// EnterMixinReference is called when entering the mixinReference production.
	EnterMixinReference(c *MixinReferenceContext)

	// EnterSelectors is called when entering the selectors production.
	EnterSelectors(c *SelectorsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

	// EnterPseudo is called when entering the pseudo production.
	EnterPseudo(c *PseudoContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterSelectorPrefix is called when entering the selectorPrefix production.
	EnterSelectorPrefix(c *SelectorPrefixContext)

	// EnterAttribRelate is called when entering the attribRelate production.
	EnterAttribRelate(c *AttribRelateContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierPart is called when entering the identifierPart production.
	EnterIdentifierPart(c *IdentifierPartContext)

	// EnterIdentifierVariableName is called when entering the identifierVariableName production.
	EnterIdentifierVariableName(c *IdentifierVariableNameContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterMeasurement is called when entering the measurement production.
	EnterMeasurement(c *MeasurementContext)

	// ExitStylesheet is called when exiting the stylesheet production.
	ExitStylesheet(c *StylesheetContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitCommandStatement is called when exiting the commandStatement production.
	ExitCommandStatement(c *CommandStatementContext)

	// ExitMathCharacter is called when exiting the mathCharacter production.
	ExitMathCharacter(c *MathCharacterContext)

	// ExitMathStatement is called when exiting the mathStatement production.
	ExitMathStatement(c *MathStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitConditionStatement is called when exiting the conditionStatement production.
	ExitConditionStatement(c *ConditionStatementContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitImportOption is called when exiting the importOption production.
	ExitImportOption(c *ImportOptionContext)

	// ExitReferenceUrl is called when exiting the referenceUrl production.
	ExitReferenceUrl(c *ReferenceUrlContext)

	// ExitMediaTypes is called when exiting the mediaTypes production.
	ExitMediaTypes(c *MediaTypesContext)

	// ExitRuleset is called when exiting the ruleset production.
	ExitRuleset(c *RulesetContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitMixinDefinition is called when exiting the mixinDefinition production.
	ExitMixinDefinition(c *MixinDefinitionContext)

	// ExitMixinGuard is called when exiting the mixinGuard production.
	ExitMixinGuard(c *MixinGuardContext)

	// ExitMixinDefinitionParam is called when exiting the mixinDefinitionParam production.
	ExitMixinDefinitionParam(c *MixinDefinitionParamContext)

	// ExitMixinReference is called when exiting the mixinReference production.
	ExitMixinReference(c *MixinReferenceContext)

	// ExitSelectors is called when exiting the selectors production.
	ExitSelectors(c *SelectorsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)

	// ExitPseudo is called when exiting the pseudo production.
	ExitPseudo(c *PseudoContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitSelectorPrefix is called when exiting the selectorPrefix production.
	ExitSelectorPrefix(c *SelectorPrefixContext)

	// ExitAttribRelate is called when exiting the attribRelate production.
	ExitAttribRelate(c *AttribRelateContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierPart is called when exiting the identifierPart production.
	ExitIdentifierPart(c *IdentifierPartContext)

	// ExitIdentifierVariableName is called when exiting the identifierVariableName production.
	ExitIdentifierVariableName(c *IdentifierVariableNameContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitMeasurement is called when exiting the measurement production.
	ExitMeasurement(c *MeasurementContext)
}
