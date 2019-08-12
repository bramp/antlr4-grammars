// Code generated from LessParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package less // LessParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLessParserListener is a complete listener for a parse tree produced by LessParser.
type BaseLessParserListener struct{}

var _ LessParserListener = &BaseLessParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLessParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLessParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLessParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLessParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStylesheet is called when production stylesheet is entered.
func (s *BaseLessParserListener) EnterStylesheet(ctx *StylesheetContext) {}

// ExitStylesheet is called when production stylesheet is exited.
func (s *BaseLessParserListener) ExitStylesheet(ctx *StylesheetContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLessParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLessParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseLessParserListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseLessParserListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterCommandStatement is called when production commandStatement is entered.
func (s *BaseLessParserListener) EnterCommandStatement(ctx *CommandStatementContext) {}

// ExitCommandStatement is called when production commandStatement is exited.
func (s *BaseLessParserListener) ExitCommandStatement(ctx *CommandStatementContext) {}

// EnterMathCharacter is called when production mathCharacter is entered.
func (s *BaseLessParserListener) EnterMathCharacter(ctx *MathCharacterContext) {}

// ExitMathCharacter is called when production mathCharacter is exited.
func (s *BaseLessParserListener) ExitMathCharacter(ctx *MathCharacterContext) {}

// EnterMathStatement is called when production mathStatement is entered.
func (s *BaseLessParserListener) EnterMathStatement(ctx *MathStatementContext) {}

// ExitMathStatement is called when production mathStatement is exited.
func (s *BaseLessParserListener) ExitMathStatement(ctx *MathStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLessParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLessParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseLessParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseLessParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaseLessParserListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaseLessParserListener) ExitConditions(ctx *ConditionsContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseLessParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseLessParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterConditionStatement is called when production conditionStatement is entered.
func (s *BaseLessParserListener) EnterConditionStatement(ctx *ConditionStatementContext) {}

// ExitConditionStatement is called when production conditionStatement is exited.
func (s *BaseLessParserListener) ExitConditionStatement(ctx *ConditionStatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseLessParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseLessParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseLessParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseLessParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterImportOption is called when production importOption is entered.
func (s *BaseLessParserListener) EnterImportOption(ctx *ImportOptionContext) {}

// ExitImportOption is called when production importOption is exited.
func (s *BaseLessParserListener) ExitImportOption(ctx *ImportOptionContext) {}

// EnterReferenceUrl is called when production referenceUrl is entered.
func (s *BaseLessParserListener) EnterReferenceUrl(ctx *ReferenceUrlContext) {}

// ExitReferenceUrl is called when production referenceUrl is exited.
func (s *BaseLessParserListener) ExitReferenceUrl(ctx *ReferenceUrlContext) {}

// EnterMediaTypes is called when production mediaTypes is entered.
func (s *BaseLessParserListener) EnterMediaTypes(ctx *MediaTypesContext) {}

// ExitMediaTypes is called when production mediaTypes is exited.
func (s *BaseLessParserListener) ExitMediaTypes(ctx *MediaTypesContext) {}

// EnterRuleset is called when production ruleset is entered.
func (s *BaseLessParserListener) EnterRuleset(ctx *RulesetContext) {}

// ExitRuleset is called when production ruleset is exited.
func (s *BaseLessParserListener) ExitRuleset(ctx *RulesetContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLessParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLessParserListener) ExitBlock(ctx *BlockContext) {}

// EnterMixinDefinition is called when production mixinDefinition is entered.
func (s *BaseLessParserListener) EnterMixinDefinition(ctx *MixinDefinitionContext) {}

// ExitMixinDefinition is called when production mixinDefinition is exited.
func (s *BaseLessParserListener) ExitMixinDefinition(ctx *MixinDefinitionContext) {}

// EnterMixinGuard is called when production mixinGuard is entered.
func (s *BaseLessParserListener) EnterMixinGuard(ctx *MixinGuardContext) {}

// ExitMixinGuard is called when production mixinGuard is exited.
func (s *BaseLessParserListener) ExitMixinGuard(ctx *MixinGuardContext) {}

// EnterMixinDefinitionParam is called when production mixinDefinitionParam is entered.
func (s *BaseLessParserListener) EnterMixinDefinitionParam(ctx *MixinDefinitionParamContext) {}

// ExitMixinDefinitionParam is called when production mixinDefinitionParam is exited.
func (s *BaseLessParserListener) ExitMixinDefinitionParam(ctx *MixinDefinitionParamContext) {}

// EnterMixinReference is called when production mixinReference is entered.
func (s *BaseLessParserListener) EnterMixinReference(ctx *MixinReferenceContext) {}

// ExitMixinReference is called when production mixinReference is exited.
func (s *BaseLessParserListener) ExitMixinReference(ctx *MixinReferenceContext) {}

// EnterSelectors is called when production selectors is entered.
func (s *BaseLessParserListener) EnterSelectors(ctx *SelectorsContext) {}

// ExitSelectors is called when production selectors is exited.
func (s *BaseLessParserListener) ExitSelectors(ctx *SelectorsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseLessParserListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseLessParserListener) ExitSelector(ctx *SelectorContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseLessParserListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseLessParserListener) ExitAttrib(ctx *AttribContext) {}

// EnterNegation is called when production negation is entered.
func (s *BaseLessParserListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production negation is exited.
func (s *BaseLessParserListener) ExitNegation(ctx *NegationContext) {}

// EnterPseudo is called when production pseudo is entered.
func (s *BaseLessParserListener) EnterPseudo(ctx *PseudoContext) {}

// ExitPseudo is called when production pseudo is exited.
func (s *BaseLessParserListener) ExitPseudo(ctx *PseudoContext) {}

// EnterElement is called when production element is entered.
func (s *BaseLessParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseLessParserListener) ExitElement(ctx *ElementContext) {}

// EnterSelectorPrefix is called when production selectorPrefix is entered.
func (s *BaseLessParserListener) EnterSelectorPrefix(ctx *SelectorPrefixContext) {}

// ExitSelectorPrefix is called when production selectorPrefix is exited.
func (s *BaseLessParserListener) ExitSelectorPrefix(ctx *SelectorPrefixContext) {}

// EnterAttribRelate is called when production attribRelate is entered.
func (s *BaseLessParserListener) EnterAttribRelate(ctx *AttribRelateContext) {}

// ExitAttribRelate is called when production attribRelate is exited.
func (s *BaseLessParserListener) ExitAttribRelate(ctx *AttribRelateContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseLessParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseLessParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierPart is called when production identifierPart is entered.
func (s *BaseLessParserListener) EnterIdentifierPart(ctx *IdentifierPartContext) {}

// ExitIdentifierPart is called when production identifierPart is exited.
func (s *BaseLessParserListener) ExitIdentifierPart(ctx *IdentifierPartContext) {}

// EnterIdentifierVariableName is called when production identifierVariableName is entered.
func (s *BaseLessParserListener) EnterIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// ExitIdentifierVariableName is called when production identifierVariableName is exited.
func (s *BaseLessParserListener) ExitIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseLessParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseLessParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterValues is called when production values is entered.
func (s *BaseLessParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseLessParserListener) ExitValues(ctx *ValuesContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseLessParserListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseLessParserListener) ExitUrl(ctx *UrlContext) {}

// EnterMeasurement is called when production measurement is entered.
func (s *BaseLessParserListener) EnterMeasurement(ctx *MeasurementContext) {}

// ExitMeasurement is called when production measurement is exited.
func (s *BaseLessParserListener) ExitMeasurement(ctx *MeasurementContext) {}
