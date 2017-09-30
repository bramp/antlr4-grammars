// Generated from UCBLogo.g4 by ANTLR 4.7.

package ucb-logo // UCBLogo
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseUCBLogoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseUCBLogoVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedureDefInstruction(ctx *ProcedureDefInstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitMacroDefInstruction(ctx *MacroDefInstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedureCallExtraInputInstruction(ctx *ProcedureCallExtraInputInstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedureCallInstruction(ctx *ProcedureCallInstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedure_def(ctx *Procedure_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitMacro_def(ctx *Macro_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitVariables(ctx *VariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitBody_def(ctx *Body_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitBody_instruction(ctx *Body_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedure_call_extra_input(ctx *Procedure_call_extra_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedure_call(ctx *Procedure_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitNotEqualsExpressionExpression(ctx *NotEqualsExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitArrayExpression(ctx *ArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitAdditionExpression(ctx *AdditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitWordExpression(ctx *WordExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitNumberExpression(ctx *NumberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitParensExpression(ctx *ParensExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitMultiplyExpression(ctx *MultiplyExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitGreaterThanExpression(ctx *GreaterThanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitDivideExpression(ctx *DivideExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitLessThanEqualsExpression(ctx *LessThanEqualsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitGreaterThanEqualsExpression(ctx *GreaterThanEqualsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitQuotedWordExpression(ctx *QuotedWordExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitEqualsExpression(ctx *EqualsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitSubtractionExpression(ctx *SubtractionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedureCallExpression(ctx *ProcedureCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitLessThanExpression(ctx *LessThanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitProcedureCallExtraInput(ctx *ProcedureCallExtraInputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitListExpression(ctx *ListExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitNameExpression(ctx *NameExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseUCBLogoVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}
