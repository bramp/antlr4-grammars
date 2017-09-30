// Generated from ScssParser.g4 by ANTLR 4.7.

package scss // ScssParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseScssParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseScssParserVisitor) VisitStylesheet(ctx *StylesheetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitParamOptionalValue(ctx *ParamOptionalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitMixinDeclaration(ctx *MixinDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIncludeDeclaration(ctx *IncludeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFunctionReturn(ctx *FunctionReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitCommandStatement(ctx *CommandStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitMathCharacter(ctx *MathCharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitMathStatement(ctx *MathStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIfDeclaration(ctx *IfDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitElseIfStatement(ctx *ElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitConditions(ctx *ConditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitForDeclaration(ctx *ForDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFromNumber(ctx *FromNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitThroughNumber(ctx *ThroughNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitWhileDeclaration(ctx *WhileDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitEachDeclaration(ctx *EachDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitEachValueList(ctx *EachValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIdentifierListOrMap(ctx *IdentifierListOrMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIdentifierValue(ctx *IdentifierValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitReferenceUrl(ctx *ReferenceUrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitMediaTypes(ctx *MediaTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitNested(ctx *NestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitNest(ctx *NestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitRuleset(ctx *RulesetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitSelectors(ctx *SelectorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitSelectorPrefix(ctx *SelectorPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitPseudo(ctx *PseudoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitAttrib(ctx *AttribContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitAttribRelate(ctx *AttribRelateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIdentifierPart(ctx *IdentifierPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitIdentifierVariableName(ctx *IdentifierVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitUrl(ctx *UrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitMeasurement(ctx *MeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseScssParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}
