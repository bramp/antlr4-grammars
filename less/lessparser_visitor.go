// Generated from LessParser.g4 by ANTLR 4.7.

package less // LessParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LessParser.
type LessParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LessParser#stylesheet.
	VisitStylesheet(ctx *StylesheetContext) interface{}

	// Visit a parse tree produced by LessParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LessParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by LessParser#commandStatement.
	VisitCommandStatement(ctx *CommandStatementContext) interface{}

	// Visit a parse tree produced by LessParser#mathCharacter.
	VisitMathCharacter(ctx *MathCharacterContext) interface{}

	// Visit a parse tree produced by LessParser#mathStatement.
	VisitMathStatement(ctx *MathStatementContext) interface{}

	// Visit a parse tree produced by LessParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by LessParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by LessParser#conditions.
	VisitConditions(ctx *ConditionsContext) interface{}

	// Visit a parse tree produced by LessParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by LessParser#conditionStatement.
	VisitConditionStatement(ctx *ConditionStatementContext) interface{}

	// Visit a parse tree produced by LessParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by LessParser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by LessParser#referenceUrl.
	VisitReferenceUrl(ctx *ReferenceUrlContext) interface{}

	// Visit a parse tree produced by LessParser#mediaTypes.
	VisitMediaTypes(ctx *MediaTypesContext) interface{}

	// Visit a parse tree produced by LessParser#ruleset.
	VisitRuleset(ctx *RulesetContext) interface{}

	// Visit a parse tree produced by LessParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LessParser#mixinDefinition.
	VisitMixinDefinition(ctx *MixinDefinitionContext) interface{}

	// Visit a parse tree produced by LessParser#mixinGuard.
	VisitMixinGuard(ctx *MixinGuardContext) interface{}

	// Visit a parse tree produced by LessParser#mixinDefinitionParam.
	VisitMixinDefinitionParam(ctx *MixinDefinitionParamContext) interface{}

	// Visit a parse tree produced by LessParser#mixinReference.
	VisitMixinReference(ctx *MixinReferenceContext) interface{}

	// Visit a parse tree produced by LessParser#selectors.
	VisitSelectors(ctx *SelectorsContext) interface{}

	// Visit a parse tree produced by LessParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by LessParser#attrib.
	VisitAttrib(ctx *AttribContext) interface{}

	// Visit a parse tree produced by LessParser#pseudo.
	VisitPseudo(ctx *PseudoContext) interface{}

	// Visit a parse tree produced by LessParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by LessParser#selectorPrefix.
	VisitSelectorPrefix(ctx *SelectorPrefixContext) interface{}

	// Visit a parse tree produced by LessParser#attribRelate.
	VisitAttribRelate(ctx *AttribRelateContext) interface{}

	// Visit a parse tree produced by LessParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by LessParser#identifierPart.
	VisitIdentifierPart(ctx *IdentifierPartContext) interface{}

	// Visit a parse tree produced by LessParser#identifierVariableName.
	VisitIdentifierVariableName(ctx *IdentifierVariableNameContext) interface{}

	// Visit a parse tree produced by LessParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by LessParser#values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by LessParser#url.
	VisitUrl(ctx *UrlContext) interface{}

	// Visit a parse tree produced by LessParser#measurement.
	VisitMeasurement(ctx *MeasurementContext) interface{}
}
