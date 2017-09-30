// Generated from Corundum.g4 by ANTLR 4.7.

package ruby // Corundum
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCorundumVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCorundumVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitGlobal_get(ctx *Global_getContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitGlobal_set(ctx *Global_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitGlobal_result(ctx *Global_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_inline_call(ctx *Function_inline_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitRequire_block(ctx *Require_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitPir_inline(ctx *Pir_inlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitPir_expression_list(ctx *Pir_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition(ctx *Function_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition_body(ctx *Function_definition_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition_header(ctx *Function_definition_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition_params(ctx *Function_definition_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition_params_list(ctx *Function_definition_params_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_definition_param_id(ctx *Function_definition_param_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_call_param_list(ctx *Function_call_param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_call_params(ctx *Function_call_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_param(ctx *Function_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_unnamed_param(ctx *Function_unnamed_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_named_param(ctx *Function_named_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFunction_call_assignment(ctx *Function_call_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitAll_result(ctx *All_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitElsif_statement(ctx *Elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitIf_elsif_statement(ctx *If_elsif_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitUnless_statement(ctx *Unless_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitInit_expression(ctx *Init_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitAll_assignment(ctx *All_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFor_init_list(ctx *For_init_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitCond_expression(ctx *Cond_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitLoop_expression(ctx *Loop_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFor_loop_list(ctx *For_loop_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitStatement_body(ctx *Statement_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitStatement_expression_list(ctx *Statement_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitDynamic_assignment(ctx *Dynamic_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitInt_assignment(ctx *Int_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFloat_assignment(ctx *Float_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitString_assignment(ctx *String_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitInitial_array_assignment(ctx *Initial_array_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitArray_assignment(ctx *Array_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitArray_definition(ctx *Array_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitArray_definition_elements(ctx *Array_definition_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitArray_selector(ctx *Array_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitDynamic_result(ctx *Dynamic_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitDynamic(ctx *DynamicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitInt_result(ctx *Int_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFloat_result(ctx *Float_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitString_result(ctx *String_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitComparison_list(ctx *Comparison_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitComp_var(ctx *Comp_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitLvalue(ctx *LvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitRvalue(ctx *RvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitBreak_expression(ctx *Break_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitLiteral_t(ctx *Literal_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitFloat_t(ctx *Float_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitInt_t(ctx *Int_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitBool_t(ctx *Bool_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitNil_t(ctx *Nil_tContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitId_global(ctx *Id_globalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitId_function(ctx *Id_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitTerminator(ctx *TerminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitElse_token(ctx *Else_tokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCorundumVisitor) VisitCrlf(ctx *CrlfContext) interface{} {
	return v.VisitChildren(ctx)
}
