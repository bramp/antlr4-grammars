// Generated from ScssParser.g4 by ANTLR 4.7.

package scss // ScssParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScssParserListener is a complete listener for a parse tree produced by ScssParser.
type ScssParserListener interface {
	antlr.ParseTreeListener

	// EnterStylesheet is called when entering the stylesheet production.
	EnterStylesheet(c *StylesheetContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterParamOptionalValue is called when entering the paramOptionalValue production.
	EnterParamOptionalValue(c *ParamOptionalValueContext)

	// EnterMixinDeclaration is called when entering the mixinDeclaration production.
	EnterMixinDeclaration(c *MixinDeclarationContext)

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

	// EnterIdentifierListOrMap is called when entering the identifierListOrMap production.
	EnterIdentifierListOrMap(c *IdentifierListOrMapContext)

	// EnterIdentifierValue is called when entering the identifierValue production.
	EnterIdentifierValue(c *IdentifierValueContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterReferenceUrl is called when entering the referenceUrl production.
	EnterReferenceUrl(c *ReferenceUrlContext)

	// EnterMediaTypes is called when entering the mediaTypes production.
	EnterMediaTypes(c *MediaTypesContext)

	// EnterNested is called when entering the nested production.
	EnterNested(c *NestedContext)

	// EnterNest is called when entering the nest production.
	EnterNest(c *NestContext)

	// EnterRuleset is called when entering the ruleset production.
	EnterRuleset(c *RulesetContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSelectors is called when entering the selectors production.
	EnterSelectors(c *SelectorsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterSelectorPrefix is called when entering the selectorPrefix production.
	EnterSelectorPrefix(c *SelectorPrefixContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterPseudo is called when entering the pseudo production.
	EnterPseudo(c *PseudoContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

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

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// ExitStylesheet is called when exiting the stylesheet production.
	ExitStylesheet(c *StylesheetContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitParamOptionalValue is called when exiting the paramOptionalValue production.
	ExitParamOptionalValue(c *ParamOptionalValueContext)

	// ExitMixinDeclaration is called when exiting the mixinDeclaration production.
	ExitMixinDeclaration(c *MixinDeclarationContext)

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

	// ExitIdentifierListOrMap is called when exiting the identifierListOrMap production.
	ExitIdentifierListOrMap(c *IdentifierListOrMapContext)

	// ExitIdentifierValue is called when exiting the identifierValue production.
	ExitIdentifierValue(c *IdentifierValueContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitReferenceUrl is called when exiting the referenceUrl production.
	ExitReferenceUrl(c *ReferenceUrlContext)

	// ExitMediaTypes is called when exiting the mediaTypes production.
	ExitMediaTypes(c *MediaTypesContext)

	// ExitNested is called when exiting the nested production.
	ExitNested(c *NestedContext)

	// ExitNest is called when exiting the nest production.
	ExitNest(c *NestContext)

	// ExitRuleset is called when exiting the ruleset production.
	ExitRuleset(c *RulesetContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSelectors is called when exiting the selectors production.
	ExitSelectors(c *SelectorsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitSelectorPrefix is called when exiting the selectorPrefix production.
	ExitSelectorPrefix(c *SelectorPrefixContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitPseudo is called when exiting the pseudo production.
	ExitPseudo(c *PseudoContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

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

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)
}
