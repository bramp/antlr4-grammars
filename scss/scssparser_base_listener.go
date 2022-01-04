// Code generated from ScssParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

// EnterDeclaredParams is called when production declaredParams is entered.
func (s *BaseScssParserListener) EnterDeclaredParams(ctx *DeclaredParamsContext) {}

// ExitDeclaredParams is called when production declaredParams is exited.
func (s *BaseScssParserListener) ExitDeclaredParams(ctx *DeclaredParamsContext) {}

// EnterDeclaredParam is called when production declaredParam is entered.
func (s *BaseScssParserListener) EnterDeclaredParam(ctx *DeclaredParamContext) {}

// ExitDeclaredParam is called when production declaredParam is exited.
func (s *BaseScssParserListener) ExitDeclaredParam(ctx *DeclaredParamContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseScssParserListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseScssParserListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterParamOptionalValue is called when production paramOptionalValue is entered.
func (s *BaseScssParserListener) EnterParamOptionalValue(ctx *ParamOptionalValueContext) {}

// ExitParamOptionalValue is called when production paramOptionalValue is exited.
func (s *BaseScssParserListener) ExitParamOptionalValue(ctx *ParamOptionalValueContext) {}

// EnterPassedParams is called when production passedParams is entered.
func (s *BaseScssParserListener) EnterPassedParams(ctx *PassedParamsContext) {}

// ExitPassedParams is called when production passedParams is exited.
func (s *BaseScssParserListener) ExitPassedParams(ctx *PassedParamsContext) {}

// EnterPassedParam is called when production passedParam is entered.
func (s *BaseScssParserListener) EnterPassedParam(ctx *PassedParamContext) {}

// ExitPassedParam is called when production passedParam is exited.
func (s *BaseScssParserListener) ExitPassedParam(ctx *PassedParamContext) {}

// EnterMixinDeclaration is called when production mixinDeclaration is entered.
func (s *BaseScssParserListener) EnterMixinDeclaration(ctx *MixinDeclarationContext) {}

// ExitMixinDeclaration is called when production mixinDeclaration is exited.
func (s *BaseScssParserListener) ExitMixinDeclaration(ctx *MixinDeclarationContext) {}

// EnterContentDeclaration is called when production contentDeclaration is entered.
func (s *BaseScssParserListener) EnterContentDeclaration(ctx *ContentDeclarationContext) {}

// ExitContentDeclaration is called when production contentDeclaration is exited.
func (s *BaseScssParserListener) ExitContentDeclaration(ctx *ContentDeclarationContext) {}

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

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseScssParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseScssParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterReferenceUrl is called when production referenceUrl is entered.
func (s *BaseScssParserListener) EnterReferenceUrl(ctx *ReferenceUrlContext) {}

// ExitReferenceUrl is called when production referenceUrl is exited.
func (s *BaseScssParserListener) ExitReferenceUrl(ctx *ReferenceUrlContext) {}

// EnterAsClause is called when production asClause is entered.
func (s *BaseScssParserListener) EnterAsClause(ctx *AsClauseContext) {}

// ExitAsClause is called when production asClause is exited.
func (s *BaseScssParserListener) ExitAsClause(ctx *AsClauseContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseScssParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseScssParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterKeywordArgument is called when production keywordArgument is entered.
func (s *BaseScssParserListener) EnterKeywordArgument(ctx *KeywordArgumentContext) {}

// ExitKeywordArgument is called when production keywordArgument is exited.
func (s *BaseScssParserListener) ExitKeywordArgument(ctx *KeywordArgumentContext) {}

// EnterMediaDeclaration is called when production mediaDeclaration is entered.
func (s *BaseScssParserListener) EnterMediaDeclaration(ctx *MediaDeclarationContext) {}

// ExitMediaDeclaration is called when production mediaDeclaration is exited.
func (s *BaseScssParserListener) ExitMediaDeclaration(ctx *MediaDeclarationContext) {}

// EnterMediaQueryList is called when production mediaQueryList is entered.
func (s *BaseScssParserListener) EnterMediaQueryList(ctx *MediaQueryListContext) {}

// ExitMediaQueryList is called when production mediaQueryList is exited.
func (s *BaseScssParserListener) ExitMediaQueryList(ctx *MediaQueryListContext) {}

// EnterMediaQuery is called when production mediaQuery is entered.
func (s *BaseScssParserListener) EnterMediaQuery(ctx *MediaQueryContext) {}

// ExitMediaQuery is called when production mediaQuery is exited.
func (s *BaseScssParserListener) ExitMediaQuery(ctx *MediaQueryContext) {}

// EnterMediaType is called when production mediaType is entered.
func (s *BaseScssParserListener) EnterMediaType(ctx *MediaTypeContext) {}

// ExitMediaType is called when production mediaType is exited.
func (s *BaseScssParserListener) ExitMediaType(ctx *MediaTypeContext) {}

// EnterMediaExpression is called when production mediaExpression is entered.
func (s *BaseScssParserListener) EnterMediaExpression(ctx *MediaExpressionContext) {}

// ExitMediaExpression is called when production mediaExpression is exited.
func (s *BaseScssParserListener) ExitMediaExpression(ctx *MediaExpressionContext) {}

// EnterMediaFeature is called when production mediaFeature is entered.
func (s *BaseScssParserListener) EnterMediaFeature(ctx *MediaFeatureContext) {}

// ExitMediaFeature is called when production mediaFeature is exited.
func (s *BaseScssParserListener) ExitMediaFeature(ctx *MediaFeatureContext) {}

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

// EnterElement is called when production element is entered.
func (s *BaseScssParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseScssParserListener) ExitElement(ctx *ElementContext) {}

// EnterCombinator is called when production combinator is entered.
func (s *BaseScssParserListener) EnterCombinator(ctx *CombinatorContext) {}

// ExitCombinator is called when production combinator is exited.
func (s *BaseScssParserListener) ExitCombinator(ctx *CombinatorContext) {}

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

// EnterPseudoIdentifier is called when production pseudoIdentifier is entered.
func (s *BaseScssParserListener) EnterPseudoIdentifier(ctx *PseudoIdentifierContext) {}

// ExitPseudoIdentifier is called when production pseudoIdentifier is exited.
func (s *BaseScssParserListener) ExitPseudoIdentifier(ctx *PseudoIdentifierContext) {}

// EnterIdentifierPart is called when production identifierPart is entered.
func (s *BaseScssParserListener) EnterIdentifierPart(ctx *IdentifierPartContext) {}

// ExitIdentifierPart is called when production identifierPart is exited.
func (s *BaseScssParserListener) ExitIdentifierPart(ctx *IdentifierPartContext) {}

// EnterIdentifierVariableName is called when production identifierVariableName is entered.
func (s *BaseScssParserListener) EnterIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// ExitIdentifierVariableName is called when production identifierVariableName is exited.
func (s *BaseScssParserListener) ExitIdentifierVariableName(ctx *IdentifierVariableNameContext) {}

// EnterProperty_ is called when production property_ is entered.
func (s *BaseScssParserListener) EnterProperty_(ctx *Property_Context) {}

// ExitProperty_ is called when production property_ is exited.
func (s *BaseScssParserListener) ExitProperty_(ctx *Property_Context) {}

// EnterLastProperty is called when production lastProperty is entered.
func (s *BaseScssParserListener) EnterLastProperty(ctx *LastPropertyContext) {}

// ExitLastProperty is called when production lastProperty is exited.
func (s *BaseScssParserListener) ExitLastProperty(ctx *LastPropertyContext) {}

// EnterPropertyValue is called when production propertyValue is entered.
func (s *BaseScssParserListener) EnterPropertyValue(ctx *PropertyValueContext) {}

// ExitPropertyValue is called when production propertyValue is exited.
func (s *BaseScssParserListener) ExitPropertyValue(ctx *PropertyValueContext) {}

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

// EnterNamespace is called when production namespace is entered.
func (s *BaseScssParserListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseScssParserListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseScssParserListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseScssParserListener) ExitList_(ctx *List_Context) {}

// EnterListCommaSeparated is called when production listCommaSeparated is entered.
func (s *BaseScssParserListener) EnterListCommaSeparated(ctx *ListCommaSeparatedContext) {}

// ExitListCommaSeparated is called when production listCommaSeparated is exited.
func (s *BaseScssParserListener) ExitListCommaSeparated(ctx *ListCommaSeparatedContext) {}

// EnterListSpaceSeparated is called when production listSpaceSeparated is entered.
func (s *BaseScssParserListener) EnterListSpaceSeparated(ctx *ListSpaceSeparatedContext) {}

// ExitListSpaceSeparated is called when production listSpaceSeparated is exited.
func (s *BaseScssParserListener) ExitListSpaceSeparated(ctx *ListSpaceSeparatedContext) {}

// EnterListBracketed is called when production listBracketed is entered.
func (s *BaseScssParserListener) EnterListBracketed(ctx *ListBracketedContext) {}

// ExitListBracketed is called when production listBracketed is exited.
func (s *BaseScssParserListener) ExitListBracketed(ctx *ListBracketedContext) {}

// EnterListElement is called when production listElement is entered.
func (s *BaseScssParserListener) EnterListElement(ctx *ListElementContext) {}

// ExitListElement is called when production listElement is exited.
func (s *BaseScssParserListener) ExitListElement(ctx *ListElementContext) {}

// EnterMap_ is called when production map_ is entered.
func (s *BaseScssParserListener) EnterMap_(ctx *Map_Context) {}

// ExitMap_ is called when production map_ is exited.
func (s *BaseScssParserListener) ExitMap_(ctx *Map_Context) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseScssParserListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseScssParserListener) ExitMapEntry(ctx *MapEntryContext) {}

// EnterMapKey is called when production mapKey is entered.
func (s *BaseScssParserListener) EnterMapKey(ctx *MapKeyContext) {}

// ExitMapKey is called when production mapKey is exited.
func (s *BaseScssParserListener) ExitMapKey(ctx *MapKeyContext) {}

// EnterMapValue is called when production mapValue is entered.
func (s *BaseScssParserListener) EnterMapValue(ctx *MapValueContext) {}

// ExitMapValue is called when production mapValue is exited.
func (s *BaseScssParserListener) ExitMapValue(ctx *MapValueContext) {}
