// Generated from ScssParser.g4 by ANTLR 4.7.

package scss // ScssParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ScssParser.
type ScssParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ScssParser#stylesheet.
	VisitStylesheet(ctx *StylesheetContext) interface{}

	// Visit a parse tree produced by ScssParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ScssParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by ScssParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by ScssParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by ScssParser#paramOptionalValue.
	VisitParamOptionalValue(ctx *ParamOptionalValueContext) interface{}

	// Visit a parse tree produced by ScssParser#mixinDeclaration.
	VisitMixinDeclaration(ctx *MixinDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#includeDeclaration.
	VisitIncludeDeclaration(ctx *IncludeDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ScssParser#functionReturn.
	VisitFunctionReturn(ctx *FunctionReturnContext) interface{}

	// Visit a parse tree produced by ScssParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by ScssParser#commandStatement.
	VisitCommandStatement(ctx *CommandStatementContext) interface{}

	// Visit a parse tree produced by ScssParser#mathCharacter.
	VisitMathCharacter(ctx *MathCharacterContext) interface{}

	// Visit a parse tree produced by ScssParser#mathStatement.
	VisitMathStatement(ctx *MathStatementContext) interface{}

	// Visit a parse tree produced by ScssParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ScssParser#ifDeclaration.
	VisitIfDeclaration(ctx *IfDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#elseIfStatement.
	VisitElseIfStatement(ctx *ElseIfStatementContext) interface{}

	// Visit a parse tree produced by ScssParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by ScssParser#conditions.
	VisitConditions(ctx *ConditionsContext) interface{}

	// Visit a parse tree produced by ScssParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by ScssParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#forDeclaration.
	VisitForDeclaration(ctx *ForDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#fromNumber.
	VisitFromNumber(ctx *FromNumberContext) interface{}

	// Visit a parse tree produced by ScssParser#throughNumber.
	VisitThroughNumber(ctx *ThroughNumberContext) interface{}

	// Visit a parse tree produced by ScssParser#whileDeclaration.
	VisitWhileDeclaration(ctx *WhileDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#eachDeclaration.
	VisitEachDeclaration(ctx *EachDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#eachValueList.
	VisitEachValueList(ctx *EachValueListContext) interface{}

	// Visit a parse tree produced by ScssParser#identifierListOrMap.
	VisitIdentifierListOrMap(ctx *IdentifierListOrMapContext) interface{}

	// Visit a parse tree produced by ScssParser#identifierValue.
	VisitIdentifierValue(ctx *IdentifierValueContext) interface{}

	// Visit a parse tree produced by ScssParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by ScssParser#referenceUrl.
	VisitReferenceUrl(ctx *ReferenceUrlContext) interface{}

	// Visit a parse tree produced by ScssParser#mediaTypes.
	VisitMediaTypes(ctx *MediaTypesContext) interface{}

	// Visit a parse tree produced by ScssParser#nested.
	VisitNested(ctx *NestedContext) interface{}

	// Visit a parse tree produced by ScssParser#nest.
	VisitNest(ctx *NestContext) interface{}

	// Visit a parse tree produced by ScssParser#ruleset.
	VisitRuleset(ctx *RulesetContext) interface{}

	// Visit a parse tree produced by ScssParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ScssParser#selectors.
	VisitSelectors(ctx *SelectorsContext) interface{}

	// Visit a parse tree produced by ScssParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by ScssParser#selectorPrefix.
	VisitSelectorPrefix(ctx *SelectorPrefixContext) interface{}

	// Visit a parse tree produced by ScssParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by ScssParser#pseudo.
	VisitPseudo(ctx *PseudoContext) interface{}

	// Visit a parse tree produced by ScssParser#attrib.
	VisitAttrib(ctx *AttribContext) interface{}

	// Visit a parse tree produced by ScssParser#attribRelate.
	VisitAttribRelate(ctx *AttribRelateContext) interface{}

	// Visit a parse tree produced by ScssParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by ScssParser#identifierPart.
	VisitIdentifierPart(ctx *IdentifierPartContext) interface{}

	// Visit a parse tree produced by ScssParser#identifierVariableName.
	VisitIdentifierVariableName(ctx *IdentifierVariableNameContext) interface{}

	// Visit a parse tree produced by ScssParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by ScssParser#values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by ScssParser#url.
	VisitUrl(ctx *UrlContext) interface{}

	// Visit a parse tree produced by ScssParser#measurement.
	VisitMeasurement(ctx *MeasurementContext) interface{}

	// Visit a parse tree produced by ScssParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}
}
