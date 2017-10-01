// Generated from LessParser.g4 by ANTLR 4.7.

package less // LessParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLessParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLessParserVisitor) VisitStylesheet(ctx *StylesheetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitCommandStatement(ctx *CommandStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMathCharacter(ctx *MathCharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMathStatement(ctx *MathStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitConditions(ctx *ConditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitConditionStatement(ctx *ConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitReferenceUrl(ctx *ReferenceUrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMediaTypes(ctx *MediaTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitRuleset(ctx *RulesetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMixinDefinition(ctx *MixinDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMixinGuard(ctx *MixinGuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMixinDefinitionParam(ctx *MixinDefinitionParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMixinReference(ctx *MixinReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitSelectors(ctx *SelectorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitAttrib(ctx *AttribContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitPseudo(ctx *PseudoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitSelectorPrefix(ctx *SelectorPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitAttribRelate(ctx *AttribRelateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitIdentifierPart(ctx *IdentifierPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitIdentifierVariableName(ctx *IdentifierVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitValues(ctx *ValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitUrl(ctx *UrlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLessParserVisitor) VisitMeasurement(ctx *MeasurementContext) interface{} {
	return v.VisitChildren(ctx)
}
