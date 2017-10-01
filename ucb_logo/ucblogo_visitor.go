// Generated from UCBLogo.g4 by ANTLR 4.7.

package ucb_logo // UCBLogo
import "github.com/antlr/antlr4/runtime/Go/antlr"


  import java.util.Map;
  import java.util.HashMap;


// A complete Visitor for a parse tree produced by UCBLogoParser.
type UCBLogoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by UCBLogoParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedureDefInstruction.
	VisitProcedureDefInstruction(ctx *ProcedureDefInstructionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#macroDefInstruction.
	VisitMacroDefInstruction(ctx *MacroDefInstructionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedureCallExtraInputInstruction.
	VisitProcedureCallExtraInputInstruction(ctx *ProcedureCallExtraInputInstructionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedureCallInstruction.
	VisitProcedureCallInstruction(ctx *ProcedureCallInstructionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedure_def.
	VisitProcedure_def(ctx *Procedure_defContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#macro_def.
	VisitMacro_def(ctx *Macro_defContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#variables.
	VisitVariables(ctx *VariablesContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#body_def.
	VisitBody_def(ctx *Body_defContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#body_instruction.
	VisitBody_instruction(ctx *Body_instructionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedure_call_extra_input.
	VisitProcedure_call_extra_input(ctx *Procedure_call_extra_inputContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedure_call.
	VisitProcedure_call(ctx *Procedure_callContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#notEqualsExpressionExpression.
	VisitNotEqualsExpressionExpression(ctx *NotEqualsExpressionExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#arrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#additionExpression.
	VisitAdditionExpression(ctx *AdditionExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#wordExpression.
	VisitWordExpression(ctx *WordExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#numberExpression.
	VisitNumberExpression(ctx *NumberExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#parensExpression.
	VisitParensExpression(ctx *ParensExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#multiplyExpression.
	VisitMultiplyExpression(ctx *MultiplyExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#greaterThanExpression.
	VisitGreaterThanExpression(ctx *GreaterThanExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#variableExpression.
	VisitVariableExpression(ctx *VariableExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#divideExpression.
	VisitDivideExpression(ctx *DivideExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#lessThanEqualsExpression.
	VisitLessThanEqualsExpression(ctx *LessThanEqualsExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#greaterThanEqualsExpression.
	VisitGreaterThanEqualsExpression(ctx *GreaterThanEqualsExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#unaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#quotedWordExpression.
	VisitQuotedWordExpression(ctx *QuotedWordExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#equalsExpression.
	VisitEqualsExpression(ctx *EqualsExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#subtractionExpression.
	VisitSubtractionExpression(ctx *SubtractionExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedureCallExpression.
	VisitProcedureCallExpression(ctx *ProcedureCallExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#lessThanExpression.
	VisitLessThanExpression(ctx *LessThanExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#procedureCallExtraInput.
	VisitProcedureCallExtraInput(ctx *ProcedureCallExtraInputContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#listExpression.
	VisitListExpression(ctx *ListExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#nameExpression.
	VisitNameExpression(ctx *NameExpressionContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by UCBLogoParser#list.
	VisitList(ctx *ListContext) interface{}

}