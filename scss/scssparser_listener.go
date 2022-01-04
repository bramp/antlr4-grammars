// Code generated from ScssParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scss // ScssParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScssParserListener is a complete listener for a parse tree produced by ScssParser.
type ScssParserListener interface {
	antlr.ParseTreeListener

	// EnterStylesheet is called when entering the stylesheet production.
	EnterStylesheet(c *StylesheetContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaredParams is called when entering the declaredParams production.
	EnterDeclaredParams(c *DeclaredParamsContext)

	// EnterDeclaredParam is called when entering the declaredParam production.
	EnterDeclaredParam(c *DeclaredParamContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterParamOptionalValue is called when entering the paramOptionalValue production.
	EnterParamOptionalValue(c *ParamOptionalValueContext)

	// EnterPassedParams is called when entering the passedParams production.
	EnterPassedParams(c *PassedParamsContext)

	// EnterPassedParam is called when entering the passedParam production.
	EnterPassedParam(c *PassedParamContext)

	// EnterMixinDeclaration is called when entering the mixinDeclaration production.
	EnterMixinDeclaration(c *MixinDeclarationContext)

	// EnterContentDeclaration is called when entering the contentDeclaration production.
	EnterContentDeclaration(c *ContentDeclarationContext)

	// EnterIncludeDeclaration is called when entering the includeDeclaration production.
	EnterIncludeDeclaration(c *IncludeDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterFunctionReturn is called when entering the functionReturn production.
	EnterFunctionReturn(c *FunctionReturnContext)

	// EnterFunctionStatement is called when entering the functionStatement production.
	EnterFunctionStatement(c *FunctionStatementContext)

	// EnterCommandStatement is called when entering the commandStatement production.
	EnterCommandStatement(c *CommandStatementContext)

	// EnterMathCharacter is called when entering the mathCharacter production.
	EnterMathCharacter(c *MathCharacterContext)

	// EnterMathStatement is called when entering the mathStatement production.
	EnterMathStatement(c *MathStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIfDeclaration is called when entering the ifDeclaration production.
	EnterIfDeclaration(c *IfDeclarationContext)

	// EnterElseIfStatement is called when entering the elseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterForDeclaration is called when entering the forDeclaration production.
	EnterForDeclaration(c *ForDeclarationContext)

	// EnterFromNumber is called when entering the fromNumber production.
	EnterFromNumber(c *FromNumberContext)

	// EnterThroughNumber is called when entering the throughNumber production.
	EnterThroughNumber(c *ThroughNumberContext)

	// EnterWhileDeclaration is called when entering the whileDeclaration production.
	EnterWhileDeclaration(c *WhileDeclarationContext)

	// EnterEachDeclaration is called when entering the eachDeclaration production.
	EnterEachDeclaration(c *EachDeclarationContext)

	// EnterEachValueList is called when entering the eachValueList production.
	EnterEachValueList(c *EachValueListContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterReferenceUrl is called when entering the referenceUrl production.
	EnterReferenceUrl(c *ReferenceUrlContext)

	// EnterAsClause is called when entering the asClause production.
	EnterAsClause(c *AsClauseContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterKeywordArgument is called when entering the keywordArgument production.
	EnterKeywordArgument(c *KeywordArgumentContext)

	// EnterMediaDeclaration is called when entering the mediaDeclaration production.
	EnterMediaDeclaration(c *MediaDeclarationContext)

	// EnterMediaQueryList is called when entering the mediaQueryList production.
	EnterMediaQueryList(c *MediaQueryListContext)

	// EnterMediaQuery is called when entering the mediaQuery production.
	EnterMediaQuery(c *MediaQueryContext)

	// EnterMediaType is called when entering the mediaType production.
	EnterMediaType(c *MediaTypeContext)

	// EnterMediaExpression is called when entering the mediaExpression production.
	EnterMediaExpression(c *MediaExpressionContext)

	// EnterMediaFeature is called when entering the mediaFeature production.
	EnterMediaFeature(c *MediaFeatureContext)

	// EnterRuleset is called when entering the ruleset production.
	EnterRuleset(c *RulesetContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSelectors is called when entering the selectors production.
	EnterSelectors(c *SelectorsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterCombinator is called when entering the combinator production.
	EnterCombinator(c *CombinatorContext)

	// EnterPseudo is called when entering the pseudo production.
	EnterPseudo(c *PseudoContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterAttribRelate is called when entering the attribRelate production.
	EnterAttribRelate(c *AttribRelateContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterPseudoIdentifier is called when entering the pseudoIdentifier production.
	EnterPseudoIdentifier(c *PseudoIdentifierContext)

	// EnterIdentifierPart is called when entering the identifierPart production.
	EnterIdentifierPart(c *IdentifierPartContext)

	// EnterIdentifierVariableName is called when entering the identifierVariableName production.
	EnterIdentifierVariableName(c *IdentifierVariableNameContext)

	// EnterProperty_ is called when entering the property_ production.
	EnterProperty_(c *Property_Context)

	// EnterLastProperty is called when entering the lastProperty production.
	EnterLastProperty(c *LastPropertyContext)

	// EnterPropertyValue is called when entering the propertyValue production.
	EnterPropertyValue(c *PropertyValueContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterMeasurement is called when entering the measurement production.
	EnterMeasurement(c *MeasurementContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterListCommaSeparated is called when entering the listCommaSeparated production.
	EnterListCommaSeparated(c *ListCommaSeparatedContext)

	// EnterListSpaceSeparated is called when entering the listSpaceSeparated production.
	EnterListSpaceSeparated(c *ListSpaceSeparatedContext)

	// EnterListBracketed is called when entering the listBracketed production.
	EnterListBracketed(c *ListBracketedContext)

	// EnterListElement is called when entering the listElement production.
	EnterListElement(c *ListElementContext)

	// EnterMap_ is called when entering the map_ production.
	EnterMap_(c *Map_Context)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterMapKey is called when entering the mapKey production.
	EnterMapKey(c *MapKeyContext)

	// EnterMapValue is called when entering the mapValue production.
	EnterMapValue(c *MapValueContext)

	// ExitStylesheet is called when exiting the stylesheet production.
	ExitStylesheet(c *StylesheetContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaredParams is called when exiting the declaredParams production.
	ExitDeclaredParams(c *DeclaredParamsContext)

	// ExitDeclaredParam is called when exiting the declaredParam production.
	ExitDeclaredParam(c *DeclaredParamContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitParamOptionalValue is called when exiting the paramOptionalValue production.
	ExitParamOptionalValue(c *ParamOptionalValueContext)

	// ExitPassedParams is called when exiting the passedParams production.
	ExitPassedParams(c *PassedParamsContext)

	// ExitPassedParam is called when exiting the passedParam production.
	ExitPassedParam(c *PassedParamContext)

	// ExitMixinDeclaration is called when exiting the mixinDeclaration production.
	ExitMixinDeclaration(c *MixinDeclarationContext)

	// ExitContentDeclaration is called when exiting the contentDeclaration production.
	ExitContentDeclaration(c *ContentDeclarationContext)

	// ExitIncludeDeclaration is called when exiting the includeDeclaration production.
	ExitIncludeDeclaration(c *IncludeDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitFunctionReturn is called when exiting the functionReturn production.
	ExitFunctionReturn(c *FunctionReturnContext)

	// ExitFunctionStatement is called when exiting the functionStatement production.
	ExitFunctionStatement(c *FunctionStatementContext)

	// ExitCommandStatement is called when exiting the commandStatement production.
	ExitCommandStatement(c *CommandStatementContext)

	// ExitMathCharacter is called when exiting the mathCharacter production.
	ExitMathCharacter(c *MathCharacterContext)

	// ExitMathStatement is called when exiting the mathStatement production.
	ExitMathStatement(c *MathStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIfDeclaration is called when exiting the ifDeclaration production.
	ExitIfDeclaration(c *IfDeclarationContext)

	// ExitElseIfStatement is called when exiting the elseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitForDeclaration is called when exiting the forDeclaration production.
	ExitForDeclaration(c *ForDeclarationContext)

	// ExitFromNumber is called when exiting the fromNumber production.
	ExitFromNumber(c *FromNumberContext)

	// ExitThroughNumber is called when exiting the throughNumber production.
	ExitThroughNumber(c *ThroughNumberContext)

	// ExitWhileDeclaration is called when exiting the whileDeclaration production.
	ExitWhileDeclaration(c *WhileDeclarationContext)

	// ExitEachDeclaration is called when exiting the eachDeclaration production.
	ExitEachDeclaration(c *EachDeclarationContext)

	// ExitEachValueList is called when exiting the eachValueList production.
	ExitEachValueList(c *EachValueListContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitReferenceUrl is called when exiting the referenceUrl production.
	ExitReferenceUrl(c *ReferenceUrlContext)

	// ExitAsClause is called when exiting the asClause production.
	ExitAsClause(c *AsClauseContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitKeywordArgument is called when exiting the keywordArgument production.
	ExitKeywordArgument(c *KeywordArgumentContext)

	// ExitMediaDeclaration is called when exiting the mediaDeclaration production.
	ExitMediaDeclaration(c *MediaDeclarationContext)

	// ExitMediaQueryList is called when exiting the mediaQueryList production.
	ExitMediaQueryList(c *MediaQueryListContext)

	// ExitMediaQuery is called when exiting the mediaQuery production.
	ExitMediaQuery(c *MediaQueryContext)

	// ExitMediaType is called when exiting the mediaType production.
	ExitMediaType(c *MediaTypeContext)

	// ExitMediaExpression is called when exiting the mediaExpression production.
	ExitMediaExpression(c *MediaExpressionContext)

	// ExitMediaFeature is called when exiting the mediaFeature production.
	ExitMediaFeature(c *MediaFeatureContext)

	// ExitRuleset is called when exiting the ruleset production.
	ExitRuleset(c *RulesetContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSelectors is called when exiting the selectors production.
	ExitSelectors(c *SelectorsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitCombinator is called when exiting the combinator production.
	ExitCombinator(c *CombinatorContext)

	// ExitPseudo is called when exiting the pseudo production.
	ExitPseudo(c *PseudoContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitAttribRelate is called when exiting the attribRelate production.
	ExitAttribRelate(c *AttribRelateContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitPseudoIdentifier is called when exiting the pseudoIdentifier production.
	ExitPseudoIdentifier(c *PseudoIdentifierContext)

	// ExitIdentifierPart is called when exiting the identifierPart production.
	ExitIdentifierPart(c *IdentifierPartContext)

	// ExitIdentifierVariableName is called when exiting the identifierVariableName production.
	ExitIdentifierVariableName(c *IdentifierVariableNameContext)

	// ExitProperty_ is called when exiting the property_ production.
	ExitProperty_(c *Property_Context)

	// ExitLastProperty is called when exiting the lastProperty production.
	ExitLastProperty(c *LastPropertyContext)

	// ExitPropertyValue is called when exiting the propertyValue production.
	ExitPropertyValue(c *PropertyValueContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitMeasurement is called when exiting the measurement production.
	ExitMeasurement(c *MeasurementContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitListCommaSeparated is called when exiting the listCommaSeparated production.
	ExitListCommaSeparated(c *ListCommaSeparatedContext)

	// ExitListSpaceSeparated is called when exiting the listSpaceSeparated production.
	ExitListSpaceSeparated(c *ListSpaceSeparatedContext)

	// ExitListBracketed is called when exiting the listBracketed production.
	ExitListBracketed(c *ListBracketedContext)

	// ExitListElement is called when exiting the listElement production.
	ExitListElement(c *ListElementContext)

	// ExitMap_ is called when exiting the map_ production.
	ExitMap_(c *Map_Context)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitMapKey is called when exiting the mapKey production.
	ExitMapKey(c *MapKeyContext)

	// ExitMapValue is called when exiting the mapValue production.
	ExitMapValue(c *MapValueContext)
}
