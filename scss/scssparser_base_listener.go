// Code generated from ScssParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package scss // ScssParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseScssParserListener is a complete listener for a parse tree produced by ScssParser.
type BaseScssParserListener struct{}

var _ ScssParserListener = &BaseScssParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseScssParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseScssParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseScssParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseScssParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStylesheet is called when production stylesheet is entered.
func (s *BaseScssParserListener) EnterStylesheet(ctx *StylesheetContext) {}

// ExitStylesheet is called when production stylesheet is exited.
func (s *BaseScssParserListener) ExitStylesheet(ctx *StylesheetContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseScssParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseScssParserListener) ExitStatement(ctx *StatementContext) {}

// EnterParams is called when production params is entered.
func (s *BaseScssParserListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseScssParserListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseScssParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseScssParserListener) ExitParam(ctx *ParamContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseScssParserListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseScssParserListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterParamOptionalValue is called when production paramOptionalValue is entered.
func (s *BaseScssParserListener) EnterParamOptionalValue(ctx *ParamOptionalValueContext) {}

// ExitParamOptionalValue is called when production paramOptionalValue is exited.
func (s *BaseScssParserListener) ExitParamOptionalValue(ctx *ParamOptionalValueContext) {}

// EnterMixinDeclaration is called when production mixinDeclaration is entered.
func (s *BaseScssParserListener) EnterMixinDeclaration(ctx *MixinDeclarationContext) {}

// ExitMixinDeclaration is called when production mixinDeclaration is exited.
func (s *BaseScssParserListener) ExitMixinDeclaration(ctx *MixinDeclarationContext) {}

// EnterIncludeDeclaration is called when production includeDeclaration is entered.
func (s *BaseScssParserListener) EnterIncludeDeclaration(ctx *IncludeDeclarationContext) {}

// ExitIncludeDeclaration is called when production includeDeclaration is exited.
func (s *BaseScssParserListener) ExitIncludeDeclaration(ctx *IncludeDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseScssParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseScssParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseScssParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseScssParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterFunctionReturn is called when production functionReturn is entered.
func (s *BaseScssParserListener) EnterFunctionReturn(ctx *FunctionReturnContext) {}

// ExitFunctionReturn is called when production functionReturn is exited.
func (s *BaseScssParserListener) ExitFunctionReturn(ctx *FunctionReturnContext) {}

// EnterFunctionStatement is called when production functionStatement is entered.
func (s *BaseScssParserListener) EnterFunctionStatement(ctx *FunctionStatementContext) {}

// ExitFunctionStatement is called when production functionStatement is exited.
func (s *BaseScssParserListener) ExitFunctionStatement(ctx *FunctionStatementContext) {}

// EnterCommandStatement is called when production commandStatement is entered.
func (s *BaseScssParserListener) EnterCommandStatement(ctx *CommandStatementContext) {}

// ExitCommandStatement is called when production commandStatement is exited.
func (s *BaseScssParserListener) ExitCommandStatement(ctx *CommandStatementContext) {}

// EnterMathCharacter is called when production mathCharacter is entered.
func (s *BaseScssParserListener) EnterMathCharacter(ctx *MathCharacterContext) {}

// ExitMathCharacter is called when production mathCharacter is exited.
func (s *BaseScssParserListener) ExitMathCharacter(ctx *MathCharacterContext) {}

// EnterMathStatement is called when production mathStatement is entered.
func (s *BaseScssParserListener) EnterMathStatement(ctx *MathStatementContext) {}

// ExitMathStatement is called when production mathStatement is exited.
func (s *BaseScssParserListener) ExitMathStatement(ctx *MathStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseScssParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseScssParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIfDeclaration is called when production ifDeclaration is entered.
func (s *BaseScssParserListener) EnterIfDeclaration(ctx *IfDeclarationContext) {}

// ExitIfDeclaration is called when production ifDeclaration is exited.
func (s *BaseScssParserListener) ExitIfDeclaration(ctx *IfDeclarationContext) {}

// EnterElseIfStatement is called when production elseIfStatement is entered.
func (s *BaseScssParserListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production elseIfStatement is exited.
func (s *BaseScssParserListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseScssParserListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseScssParserListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaseScssParserListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaseScssParserListener) ExitConditions(ctx *ConditionsContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseScssParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseScssParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseScssParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseScssParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *BaseScssParserListener) EnterForDeclaration(ctx *ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *BaseScssParserListener) ExitForDeclaration(ctx *ForDeclarationContext) {}

// EnterFromNumber is called when production fromNumber is entered.
func (s *BaseScssParserListener) EnterFromNumber(ctx *FromNumberContext) {}

// ExitFromNumber is called when production fromNumber is exited.
func (s *BaseScssParserListener) ExitFromNumber(ctx *FromNumberContext) {}

// EnterThroughNumber is called when production throughNumber is entered.
func (s *BaseScssParserListener) EnterThroughNumber(ctx *ThroughNumberContext) {}

// ExitThroughNumber is called when production throughNumber is exited.
func (s *BaseScssParserListener) ExitThroughNumber(ctx *ThroughNumberContext) {}

// EnterWhileDeclaration is called when production whileDeclaration is entered.
func (s *BaseScssParserListener) EnterWhileDeclaration(ctx *WhileDeclarationContext) {}

// ExitWhileDeclaration is called when production whileDeclaration is exited.
func (s *BaseScssParserListener) ExitWhileDeclaration(ctx *WhileDeclarationContext) {}

// EnterEachDeclaration is called when production eachDeclaration is entered.
func (s *BaseScssParserListener) EnterEachDeclaration(ctx *EachDeclarationContext) {}

// ExitEachDeclaration is called when production eachDeclaration is exited.
func (s *BaseScssParserListener) ExitEachDeclaration(ctx *EachDeclarationContext) {}

// EnterEachValueList is called when production eachValueList is entered.
func (s *BaseScssParserListener) EnterEachValueList(ctx *EachValueListContext) {}

// ExitEachValueList is called when production eachValueList is exited.
func (s *BaseScssParserListener) ExitEachValueList(ctx *EachValueListContext) {}

// EnterIdentifierListOrMap is called when production identifierListOrMap is entered.
func (s *BaseScssParserListener) EnterIdentifierListOrMap(ctx *IdentifierListOrMapContext) {}

// ExitIdentifierListOrMap is called when production identifierListOrMap is exited.
func (s *BaseScssParserListener) ExitIdentifierListOrMap(ctx *IdentifierListOrMapContext) {}

// EnterIdentifierValue is called when production identifierValue is entered.
func (s *BaseScssParserListener) EnterIdentifierValue(ctx *IdentifierValueContext) {}

// ExitIdentifierValue is called when production identifierValue is exited.
func (s *BaseScssParserListener) ExitIdentifierValue(ctx *IdentifierValueContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseScssParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseScssParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterReferenceUrl is called when production referenceUrl is entered.
func (s *BaseScssParserListener) EnterReferenceUrl(ctx *ReferenceUrlContext) {}

// ExitReferenceUrl is called when production referenceUrl is exited.
func (s *BaseScssParserListener) ExitReferenceUrl(ctx *ReferenceUrlContext) {}

// EnterMediaTypes is called when production mediaTypes is entered.
func (s *BaseScssParserListener) EnterMediaTypes(ctx *MediaTypesContext) {}

// ExitMediaTypes is called when production mediaTypes is exited.
func (s *BaseScssParserListener) ExitMediaTypes(ctx *MediaTypesContext) {}

// EnterNested is called when production nested is entered.
func (s *BaseScssParserListener) EnterNested(ctx *NestedContext) {}

// ExitNested is called when production nested is exited.
func (s *BaseScssParserListener) ExitNested(ctx *NestedContext) {}

// EnterNest is called when production nest is entered.
func (s *BaseScssParserListener) EnterNest(ctx *NestContext) {}

// ExitNest is called when production nest is exited.
func (s *BaseScssParserListener) ExitNest(ctx *NestContext) {}

// EnterRuleset is called when production ruleset is entered.
func (s *BaseScssParserListener) EnterRuleset(ctx *RulesetContext) {}

// ExitRuleset is called when production ruleset is exited.
func (s *BaseScssParserListener) ExitRuleset(ctx *RulesetContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseScssParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseScssParserListener) ExitBlock(ctx *BlockContext) {}

// EnterSelectors is called when production selectors is entered.
func (s *BaseScssParserListener) EnterSelectors(ctx *SelectorsContext) {}

// ExitSelectors is called when production selectors is exited.
func (s *BaseScssParserListener) ExitSelectors(ctx *SelectorsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseScssParserListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseScssParserListener) ExitSelector(ctx *SelectorContext) {}

// EnterSelectorPrefix is called when production selectorPrefix is entered.
func (s *BaseScssParserListener) EnterSelectorPrefix(ctx *SelectorPrefixContext) {}

// ExitSelectorPrefix is called when production selectorPrefix is exited.
func (s *BaseScssParserListener) ExitSelectorPrefix(ctx *SelectorPrefixContext) {}

// EnterElement is called when production element is entered.
func (s *BaseScssParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseScssParserListener) ExitElement(ctx *ElementContext) {}

// EnterPseudo is called when production pseudo is entered.
func (s *BaseScssParserListener) EnterPseudo(ctx *PseudoContext) {}

// ExitPseudo is called when production pseudo is exited.
func (s *BaseScssParserListener) ExitPseudo(ctx *PseudoContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseScssParserListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseScssParserListener) ExitAttrib(ctx *AttribContext) {}

// EnterAttribRelate is called when production attribRelate is entered.
func (s *BaseScssParserListener) EnterAttribRelate(ctx *AttribRelateContext) {}

// ExitAttribRelate is called when production attribRelate is exited.
func (s *BaseScssParserListener) ExitAttribRelate(ctx *AttribRelateContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseScssParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseScssParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierPart is called when production identifierPart is entered.
func (s *BaseScssParserListener) EnterIdentifierPart(ctx *IdentifierPartContext) {}

// ExitIdentifierPart is called when production identifierPart is exited.
func (s *BaseScssParserListener) ExitIdentifierPart(ctx *IdentifierPartContext) {}

// EnterIdentifierVariableName is called when production identifierVariableName is entered.
func (s *BaseScssParserListener) EnterIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// ExitIdentifierVariableName is called when production identifierVariableName is exited.
func (s *BaseScssParserListener) ExitIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseScssParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseScssParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterValues is called when production values is entered.
func (s *BaseScssParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseScssParserListener) ExitValues(ctx *ValuesContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseScssParserListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseScssParserListener) ExitUrl(ctx *UrlContext) {}

// EnterMeasurement is called when production measurement is entered.
func (s *BaseScssParserListener) EnterMeasurement(ctx *MeasurementContext) {}

// ExitMeasurement is called when production measurement is exited.
func (s *BaseScssParserListener) ExitMeasurement(ctx *MeasurementContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseScssParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseScssParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}
