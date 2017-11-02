// Generated from Corundum.g4 by ANTLR 4.7.

package corundum // Corundum
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CorundumListener is a complete listener for a parse tree produced by CorundumParser.
type CorundumListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGlobal_get is called when entering the global_get production.
	EnterGlobal_get(c *Global_getContext)

	// EnterGlobal_set is called when entering the global_set production.
	EnterGlobal_set(c *Global_setContext)

	// EnterGlobal_result is called when entering the global_result production.
	EnterGlobal_result(c *Global_resultContext)

	// EnterFunction_inline_call is called when entering the function_inline_call production.
	EnterFunction_inline_call(c *Function_inline_callContext)

	// EnterRequire_block is called when entering the require_block production.
	EnterRequire_block(c *Require_blockContext)

	// EnterPir_inline is called when entering the pir_inline production.
	EnterPir_inline(c *Pir_inlineContext)

	// EnterPir_expression_list is called when entering the pir_expression_list production.
	EnterPir_expression_list(c *Pir_expression_listContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterFunction_definition_body is called when entering the function_definition_body production.
	EnterFunction_definition_body(c *Function_definition_bodyContext)

	// EnterFunction_definition_header is called when entering the function_definition_header production.
	EnterFunction_definition_header(c *Function_definition_headerContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_definition_params is called when entering the function_definition_params production.
	EnterFunction_definition_params(c *Function_definition_paramsContext)

	// EnterFunction_definition_params_list is called when entering the function_definition_params_list production.
	EnterFunction_definition_params_list(c *Function_definition_params_listContext)

	// EnterFunction_definition_param_id is called when entering the function_definition_param_id production.
	EnterFunction_definition_param_id(c *Function_definition_param_idContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterFunction_call_param_list is called when entering the function_call_param_list production.
	EnterFunction_call_param_list(c *Function_call_param_listContext)

	// EnterFunction_call_params is called when entering the function_call_params production.
	EnterFunction_call_params(c *Function_call_paramsContext)

	// EnterFunction_param is called when entering the function_param production.
	EnterFunction_param(c *Function_paramContext)

	// EnterFunction_unnamed_param is called when entering the function_unnamed_param production.
	EnterFunction_unnamed_param(c *Function_unnamed_paramContext)

	// EnterFunction_named_param is called when entering the function_named_param production.
	EnterFunction_named_param(c *Function_named_paramContext)

	// EnterFunction_call_assignment is called when entering the function_call_assignment production.
	EnterFunction_call_assignment(c *Function_call_assignmentContext)

	// EnterAll_result is called when entering the all_result production.
	EnterAll_result(c *All_resultContext)

	// EnterElsif_statement is called when entering the elsif_statement production.
	EnterElsif_statement(c *Elsif_statementContext)

	// EnterIf_elsif_statement is called when entering the if_elsif_statement production.
	EnterIf_elsif_statement(c *If_elsif_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterUnless_statement is called when entering the unless_statement production.
	EnterUnless_statement(c *Unless_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterInit_expression is called when entering the init_expression production.
	EnterInit_expression(c *Init_expressionContext)

	// EnterAll_assignment is called when entering the all_assignment production.
	EnterAll_assignment(c *All_assignmentContext)

	// EnterFor_init_list is called when entering the for_init_list production.
	EnterFor_init_list(c *For_init_listContext)

	// EnterCond_expression is called when entering the cond_expression production.
	EnterCond_expression(c *Cond_expressionContext)

	// EnterLoop_expression is called when entering the loop_expression production.
	EnterLoop_expression(c *Loop_expressionContext)

	// EnterFor_loop_list is called when entering the for_loop_list production.
	EnterFor_loop_list(c *For_loop_listContext)

	// EnterStatement_body is called when entering the statement_body production.
	EnterStatement_body(c *Statement_bodyContext)

	// EnterStatement_expression_list is called when entering the statement_expression_list production.
	EnterStatement_expression_list(c *Statement_expression_listContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterDynamic_assignment is called when entering the dynamic_assignment production.
	EnterDynamic_assignment(c *Dynamic_assignmentContext)

	// EnterInt_assignment is called when entering the int_assignment production.
	EnterInt_assignment(c *Int_assignmentContext)

	// EnterFloat_assignment is called when entering the float_assignment production.
	EnterFloat_assignment(c *Float_assignmentContext)

	// EnterString_assignment is called when entering the string_assignment production.
	EnterString_assignment(c *String_assignmentContext)

	// EnterInitial_array_assignment is called when entering the initial_array_assignment production.
	EnterInitial_array_assignment(c *Initial_array_assignmentContext)

	// EnterArray_assignment is called when entering the array_assignment production.
	EnterArray_assignment(c *Array_assignmentContext)

	// EnterArray_definition is called when entering the array_definition production.
	EnterArray_definition(c *Array_definitionContext)

	// EnterArray_definition_elements is called when entering the array_definition_elements production.
	EnterArray_definition_elements(c *Array_definition_elementsContext)

	// EnterArray_selector is called when entering the array_selector production.
	EnterArray_selector(c *Array_selectorContext)

	// EnterDynamic_result is called when entering the dynamic_result production.
	EnterDynamic_result(c *Dynamic_resultContext)

	// EnterDynamic is called when entering the dynamic production.
	EnterDynamic(c *DynamicContext)

	// EnterInt_result is called when entering the int_result production.
	EnterInt_result(c *Int_resultContext)

	// EnterFloat_result is called when entering the float_result production.
	EnterFloat_result(c *Float_resultContext)

	// EnterString_result is called when entering the string_result production.
	EnterString_result(c *String_resultContext)

	// EnterComparison_list is called when entering the comparison_list production.
	EnterComparison_list(c *Comparison_listContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterComp_var is called when entering the comp_var production.
	EnterComp_var(c *Comp_varContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterRvalue is called when entering the rvalue production.
	EnterRvalue(c *RvalueContext)

	// EnterBreak_expression is called when entering the break_expression production.
	EnterBreak_expression(c *Break_expressionContext)

	// EnterLiteral_t is called when entering the literal_t production.
	EnterLiteral_t(c *Literal_tContext)

	// EnterFloat_t is called when entering the float_t production.
	EnterFloat_t(c *Float_tContext)

	// EnterInt_t is called when entering the int_t production.
	EnterInt_t(c *Int_tContext)

	// EnterBool_t is called when entering the bool_t production.
	EnterBool_t(c *Bool_tContext)

	// EnterNil_t is called when entering the nil_t production.
	EnterNil_t(c *Nil_tContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterId_global is called when entering the id_global production.
	EnterId_global(c *Id_globalContext)

	// EnterId_function is called when entering the id_function production.
	EnterId_function(c *Id_functionContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// EnterElse_token is called when entering the else_token production.
	EnterElse_token(c *Else_tokenContext)

	// EnterCrlf is called when entering the crlf production.
	EnterCrlf(c *CrlfContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGlobal_get is called when exiting the global_get production.
	ExitGlobal_get(c *Global_getContext)

	// ExitGlobal_set is called when exiting the global_set production.
	ExitGlobal_set(c *Global_setContext)

	// ExitGlobal_result is called when exiting the global_result production.
	ExitGlobal_result(c *Global_resultContext)

	// ExitFunction_inline_call is called when exiting the function_inline_call production.
	ExitFunction_inline_call(c *Function_inline_callContext)

	// ExitRequire_block is called when exiting the require_block production.
	ExitRequire_block(c *Require_blockContext)

	// ExitPir_inline is called when exiting the pir_inline production.
	ExitPir_inline(c *Pir_inlineContext)

	// ExitPir_expression_list is called when exiting the pir_expression_list production.
	ExitPir_expression_list(c *Pir_expression_listContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitFunction_definition_body is called when exiting the function_definition_body production.
	ExitFunction_definition_body(c *Function_definition_bodyContext)

	// ExitFunction_definition_header is called when exiting the function_definition_header production.
	ExitFunction_definition_header(c *Function_definition_headerContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_definition_params is called when exiting the function_definition_params production.
	ExitFunction_definition_params(c *Function_definition_paramsContext)

	// ExitFunction_definition_params_list is called when exiting the function_definition_params_list production.
	ExitFunction_definition_params_list(c *Function_definition_params_listContext)

	// ExitFunction_definition_param_id is called when exiting the function_definition_param_id production.
	ExitFunction_definition_param_id(c *Function_definition_param_idContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitFunction_call_param_list is called when exiting the function_call_param_list production.
	ExitFunction_call_param_list(c *Function_call_param_listContext)

	// ExitFunction_call_params is called when exiting the function_call_params production.
	ExitFunction_call_params(c *Function_call_paramsContext)

	// ExitFunction_param is called when exiting the function_param production.
	ExitFunction_param(c *Function_paramContext)

	// ExitFunction_unnamed_param is called when exiting the function_unnamed_param production.
	ExitFunction_unnamed_param(c *Function_unnamed_paramContext)

	// ExitFunction_named_param is called when exiting the function_named_param production.
	ExitFunction_named_param(c *Function_named_paramContext)

	// ExitFunction_call_assignment is called when exiting the function_call_assignment production.
	ExitFunction_call_assignment(c *Function_call_assignmentContext)

	// ExitAll_result is called when exiting the all_result production.
	ExitAll_result(c *All_resultContext)

	// ExitElsif_statement is called when exiting the elsif_statement production.
	ExitElsif_statement(c *Elsif_statementContext)

	// ExitIf_elsif_statement is called when exiting the if_elsif_statement production.
	ExitIf_elsif_statement(c *If_elsif_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitUnless_statement is called when exiting the unless_statement production.
	ExitUnless_statement(c *Unless_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitInit_expression is called when exiting the init_expression production.
	ExitInit_expression(c *Init_expressionContext)

	// ExitAll_assignment is called when exiting the all_assignment production.
	ExitAll_assignment(c *All_assignmentContext)

	// ExitFor_init_list is called when exiting the for_init_list production.
	ExitFor_init_list(c *For_init_listContext)

	// ExitCond_expression is called when exiting the cond_expression production.
	ExitCond_expression(c *Cond_expressionContext)

	// ExitLoop_expression is called when exiting the loop_expression production.
	ExitLoop_expression(c *Loop_expressionContext)

	// ExitFor_loop_list is called when exiting the for_loop_list production.
	ExitFor_loop_list(c *For_loop_listContext)

	// ExitStatement_body is called when exiting the statement_body production.
	ExitStatement_body(c *Statement_bodyContext)

	// ExitStatement_expression_list is called when exiting the statement_expression_list production.
	ExitStatement_expression_list(c *Statement_expression_listContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitDynamic_assignment is called when exiting the dynamic_assignment production.
	ExitDynamic_assignment(c *Dynamic_assignmentContext)

	// ExitInt_assignment is called when exiting the int_assignment production.
	ExitInt_assignment(c *Int_assignmentContext)

	// ExitFloat_assignment is called when exiting the float_assignment production.
	ExitFloat_assignment(c *Float_assignmentContext)

	// ExitString_assignment is called when exiting the string_assignment production.
	ExitString_assignment(c *String_assignmentContext)

	// ExitInitial_array_assignment is called when exiting the initial_array_assignment production.
	ExitInitial_array_assignment(c *Initial_array_assignmentContext)

	// ExitArray_assignment is called when exiting the array_assignment production.
	ExitArray_assignment(c *Array_assignmentContext)

	// ExitArray_definition is called when exiting the array_definition production.
	ExitArray_definition(c *Array_definitionContext)

	// ExitArray_definition_elements is called when exiting the array_definition_elements production.
	ExitArray_definition_elements(c *Array_definition_elementsContext)

	// ExitArray_selector is called when exiting the array_selector production.
	ExitArray_selector(c *Array_selectorContext)

	// ExitDynamic_result is called when exiting the dynamic_result production.
	ExitDynamic_result(c *Dynamic_resultContext)

	// ExitDynamic is called when exiting the dynamic production.
	ExitDynamic(c *DynamicContext)

	// ExitInt_result is called when exiting the int_result production.
	ExitInt_result(c *Int_resultContext)

	// ExitFloat_result is called when exiting the float_result production.
	ExitFloat_result(c *Float_resultContext)

	// ExitString_result is called when exiting the string_result production.
	ExitString_result(c *String_resultContext)

	// ExitComparison_list is called when exiting the comparison_list production.
	ExitComparison_list(c *Comparison_listContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitComp_var is called when exiting the comp_var production.
	ExitComp_var(c *Comp_varContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitRvalue is called when exiting the rvalue production.
	ExitRvalue(c *RvalueContext)

	// ExitBreak_expression is called when exiting the break_expression production.
	ExitBreak_expression(c *Break_expressionContext)

	// ExitLiteral_t is called when exiting the literal_t production.
	ExitLiteral_t(c *Literal_tContext)

	// ExitFloat_t is called when exiting the float_t production.
	ExitFloat_t(c *Float_tContext)

	// ExitInt_t is called when exiting the int_t production.
	ExitInt_t(c *Int_tContext)

	// ExitBool_t is called when exiting the bool_t production.
	ExitBool_t(c *Bool_tContext)

	// ExitNil_t is called when exiting the nil_t production.
	ExitNil_t(c *Nil_tContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitId_global is called when exiting the id_global production.
	ExitId_global(c *Id_globalContext)

	// ExitId_function is called when exiting the id_function production.
	ExitId_function(c *Id_functionContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitElse_token is called when exiting the else_token production.
	ExitElse_token(c *Else_tokenContext)

	// ExitCrlf is called when exiting the crlf production.
	ExitCrlf(c *CrlfContext)
}
