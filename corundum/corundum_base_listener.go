// Code generated from Corundum.g4 by ANTLR 4.9.3. DO NOT EDIT.

package corundum // Corundum
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCorundumListener is a complete listener for a parse tree produced by CorundumParser.
type BaseCorundumListener struct{}

var _ CorundumListener = &BaseCorundumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCorundumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCorundumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCorundumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCorundumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseCorundumListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseCorundumListener) ExitProg(ctx *ProgContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseCorundumListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseCorundumListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCorundumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCorundumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGlobal_get is called when production global_get is entered.
func (s *BaseCorundumListener) EnterGlobal_get(ctx *Global_getContext) {}

// ExitGlobal_get is called when production global_get is exited.
func (s *BaseCorundumListener) ExitGlobal_get(ctx *Global_getContext) {}

// EnterGlobal_set is called when production global_set is entered.
func (s *BaseCorundumListener) EnterGlobal_set(ctx *Global_setContext) {}

// ExitGlobal_set is called when production global_set is exited.
func (s *BaseCorundumListener) ExitGlobal_set(ctx *Global_setContext) {}

// EnterGlobal_result is called when production global_result is entered.
func (s *BaseCorundumListener) EnterGlobal_result(ctx *Global_resultContext) {}

// ExitGlobal_result is called when production global_result is exited.
func (s *BaseCorundumListener) ExitGlobal_result(ctx *Global_resultContext) {}

// EnterFunction_inline_call is called when production function_inline_call is entered.
func (s *BaseCorundumListener) EnterFunction_inline_call(ctx *Function_inline_callContext) {}

// ExitFunction_inline_call is called when production function_inline_call is exited.
func (s *BaseCorundumListener) ExitFunction_inline_call(ctx *Function_inline_callContext) {}

// EnterRequire_block is called when production require_block is entered.
func (s *BaseCorundumListener) EnterRequire_block(ctx *Require_blockContext) {}

// ExitRequire_block is called when production require_block is exited.
func (s *BaseCorundumListener) ExitRequire_block(ctx *Require_blockContext) {}

// EnterPir_inline is called when production pir_inline is entered.
func (s *BaseCorundumListener) EnterPir_inline(ctx *Pir_inlineContext) {}

// ExitPir_inline is called when production pir_inline is exited.
func (s *BaseCorundumListener) ExitPir_inline(ctx *Pir_inlineContext) {}

// EnterPir_expression_list is called when production pir_expression_list is entered.
func (s *BaseCorundumListener) EnterPir_expression_list(ctx *Pir_expression_listContext) {}

// ExitPir_expression_list is called when production pir_expression_list is exited.
func (s *BaseCorundumListener) ExitPir_expression_list(ctx *Pir_expression_listContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseCorundumListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseCorundumListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterFunction_definition_body is called when production function_definition_body is entered.
func (s *BaseCorundumListener) EnterFunction_definition_body(ctx *Function_definition_bodyContext) {}

// ExitFunction_definition_body is called when production function_definition_body is exited.
func (s *BaseCorundumListener) ExitFunction_definition_body(ctx *Function_definition_bodyContext) {}

// EnterFunction_definition_header is called when production function_definition_header is entered.
func (s *BaseCorundumListener) EnterFunction_definition_header(ctx *Function_definition_headerContext) {
}

// ExitFunction_definition_header is called when production function_definition_header is exited.
func (s *BaseCorundumListener) ExitFunction_definition_header(ctx *Function_definition_headerContext) {
}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseCorundumListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseCorundumListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_definition_params is called when production function_definition_params is entered.
func (s *BaseCorundumListener) EnterFunction_definition_params(ctx *Function_definition_paramsContext) {
}

// ExitFunction_definition_params is called when production function_definition_params is exited.
func (s *BaseCorundumListener) ExitFunction_definition_params(ctx *Function_definition_paramsContext) {
}

// EnterFunction_definition_params_list is called when production function_definition_params_list is entered.
func (s *BaseCorundumListener) EnterFunction_definition_params_list(ctx *Function_definition_params_listContext) {
}

// ExitFunction_definition_params_list is called when production function_definition_params_list is exited.
func (s *BaseCorundumListener) ExitFunction_definition_params_list(ctx *Function_definition_params_listContext) {
}

// EnterFunction_definition_param_id is called when production function_definition_param_id is entered.
func (s *BaseCorundumListener) EnterFunction_definition_param_id(ctx *Function_definition_param_idContext) {
}

// ExitFunction_definition_param_id is called when production function_definition_param_id is exited.
func (s *BaseCorundumListener) ExitFunction_definition_param_id(ctx *Function_definition_param_idContext) {
}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseCorundumListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseCorundumListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseCorundumListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseCorundumListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterFunction_call_param_list is called when production function_call_param_list is entered.
func (s *BaseCorundumListener) EnterFunction_call_param_list(ctx *Function_call_param_listContext) {}

// ExitFunction_call_param_list is called when production function_call_param_list is exited.
func (s *BaseCorundumListener) ExitFunction_call_param_list(ctx *Function_call_param_listContext) {}

// EnterFunction_call_params is called when production function_call_params is entered.
func (s *BaseCorundumListener) EnterFunction_call_params(ctx *Function_call_paramsContext) {}

// ExitFunction_call_params is called when production function_call_params is exited.
func (s *BaseCorundumListener) ExitFunction_call_params(ctx *Function_call_paramsContext) {}

// EnterFunction_param is called when production function_param is entered.
func (s *BaseCorundumListener) EnterFunction_param(ctx *Function_paramContext) {}

// ExitFunction_param is called when production function_param is exited.
func (s *BaseCorundumListener) ExitFunction_param(ctx *Function_paramContext) {}

// EnterFunction_unnamed_param is called when production function_unnamed_param is entered.
func (s *BaseCorundumListener) EnterFunction_unnamed_param(ctx *Function_unnamed_paramContext) {}

// ExitFunction_unnamed_param is called when production function_unnamed_param is exited.
func (s *BaseCorundumListener) ExitFunction_unnamed_param(ctx *Function_unnamed_paramContext) {}

// EnterFunction_named_param is called when production function_named_param is entered.
func (s *BaseCorundumListener) EnterFunction_named_param(ctx *Function_named_paramContext) {}

// ExitFunction_named_param is called when production function_named_param is exited.
func (s *BaseCorundumListener) ExitFunction_named_param(ctx *Function_named_paramContext) {}

// EnterFunction_call_assignment is called when production function_call_assignment is entered.
func (s *BaseCorundumListener) EnterFunction_call_assignment(ctx *Function_call_assignmentContext) {}

// ExitFunction_call_assignment is called when production function_call_assignment is exited.
func (s *BaseCorundumListener) ExitFunction_call_assignment(ctx *Function_call_assignmentContext) {}

// EnterAll_result is called when production all_result is entered.
func (s *BaseCorundumListener) EnterAll_result(ctx *All_resultContext) {}

// ExitAll_result is called when production all_result is exited.
func (s *BaseCorundumListener) ExitAll_result(ctx *All_resultContext) {}

// EnterElsif_statement is called when production elsif_statement is entered.
func (s *BaseCorundumListener) EnterElsif_statement(ctx *Elsif_statementContext) {}

// ExitElsif_statement is called when production elsif_statement is exited.
func (s *BaseCorundumListener) ExitElsif_statement(ctx *Elsif_statementContext) {}

// EnterIf_elsif_statement is called when production if_elsif_statement is entered.
func (s *BaseCorundumListener) EnterIf_elsif_statement(ctx *If_elsif_statementContext) {}

// ExitIf_elsif_statement is called when production if_elsif_statement is exited.
func (s *BaseCorundumListener) ExitIf_elsif_statement(ctx *If_elsif_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseCorundumListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseCorundumListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterUnless_statement is called when production unless_statement is entered.
func (s *BaseCorundumListener) EnterUnless_statement(ctx *Unless_statementContext) {}

// ExitUnless_statement is called when production unless_statement is exited.
func (s *BaseCorundumListener) ExitUnless_statement(ctx *Unless_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseCorundumListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseCorundumListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseCorundumListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseCorundumListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterInit_expression is called when production init_expression is entered.
func (s *BaseCorundumListener) EnterInit_expression(ctx *Init_expressionContext) {}

// ExitInit_expression is called when production init_expression is exited.
func (s *BaseCorundumListener) ExitInit_expression(ctx *Init_expressionContext) {}

// EnterAll_assignment is called when production all_assignment is entered.
func (s *BaseCorundumListener) EnterAll_assignment(ctx *All_assignmentContext) {}

// ExitAll_assignment is called when production all_assignment is exited.
func (s *BaseCorundumListener) ExitAll_assignment(ctx *All_assignmentContext) {}

// EnterFor_init_list is called when production for_init_list is entered.
func (s *BaseCorundumListener) EnterFor_init_list(ctx *For_init_listContext) {}

// ExitFor_init_list is called when production for_init_list is exited.
func (s *BaseCorundumListener) ExitFor_init_list(ctx *For_init_listContext) {}

// EnterCond_expression is called when production cond_expression is entered.
func (s *BaseCorundumListener) EnterCond_expression(ctx *Cond_expressionContext) {}

// ExitCond_expression is called when production cond_expression is exited.
func (s *BaseCorundumListener) ExitCond_expression(ctx *Cond_expressionContext) {}

// EnterLoop_expression is called when production loop_expression is entered.
func (s *BaseCorundumListener) EnterLoop_expression(ctx *Loop_expressionContext) {}

// ExitLoop_expression is called when production loop_expression is exited.
func (s *BaseCorundumListener) ExitLoop_expression(ctx *Loop_expressionContext) {}

// EnterFor_loop_list is called when production for_loop_list is entered.
func (s *BaseCorundumListener) EnterFor_loop_list(ctx *For_loop_listContext) {}

// ExitFor_loop_list is called when production for_loop_list is exited.
func (s *BaseCorundumListener) ExitFor_loop_list(ctx *For_loop_listContext) {}

// EnterStatement_body is called when production statement_body is entered.
func (s *BaseCorundumListener) EnterStatement_body(ctx *Statement_bodyContext) {}

// ExitStatement_body is called when production statement_body is exited.
func (s *BaseCorundumListener) ExitStatement_body(ctx *Statement_bodyContext) {}

// EnterStatement_expression_list is called when production statement_expression_list is entered.
func (s *BaseCorundumListener) EnterStatement_expression_list(ctx *Statement_expression_listContext) {
}

// ExitStatement_expression_list is called when production statement_expression_list is exited.
func (s *BaseCorundumListener) ExitStatement_expression_list(ctx *Statement_expression_listContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCorundumListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCorundumListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterDynamic_assignment is called when production dynamic_assignment is entered.
func (s *BaseCorundumListener) EnterDynamic_assignment(ctx *Dynamic_assignmentContext) {}

// ExitDynamic_assignment is called when production dynamic_assignment is exited.
func (s *BaseCorundumListener) ExitDynamic_assignment(ctx *Dynamic_assignmentContext) {}

// EnterInt_assignment is called when production int_assignment is entered.
func (s *BaseCorundumListener) EnterInt_assignment(ctx *Int_assignmentContext) {}

// ExitInt_assignment is called when production int_assignment is exited.
func (s *BaseCorundumListener) ExitInt_assignment(ctx *Int_assignmentContext) {}

// EnterFloat_assignment is called when production float_assignment is entered.
func (s *BaseCorundumListener) EnterFloat_assignment(ctx *Float_assignmentContext) {}

// ExitFloat_assignment is called when production float_assignment is exited.
func (s *BaseCorundumListener) ExitFloat_assignment(ctx *Float_assignmentContext) {}

// EnterString_assignment is called when production string_assignment is entered.
func (s *BaseCorundumListener) EnterString_assignment(ctx *String_assignmentContext) {}

// ExitString_assignment is called when production string_assignment is exited.
func (s *BaseCorundumListener) ExitString_assignment(ctx *String_assignmentContext) {}

// EnterInitial_array_assignment is called when production initial_array_assignment is entered.
func (s *BaseCorundumListener) EnterInitial_array_assignment(ctx *Initial_array_assignmentContext) {}

// ExitInitial_array_assignment is called when production initial_array_assignment is exited.
func (s *BaseCorundumListener) ExitInitial_array_assignment(ctx *Initial_array_assignmentContext) {}

// EnterArray_assignment is called when production array_assignment is entered.
func (s *BaseCorundumListener) EnterArray_assignment(ctx *Array_assignmentContext) {}

// ExitArray_assignment is called when production array_assignment is exited.
func (s *BaseCorundumListener) ExitArray_assignment(ctx *Array_assignmentContext) {}

// EnterArray_definition is called when production array_definition is entered.
func (s *BaseCorundumListener) EnterArray_definition(ctx *Array_definitionContext) {}

// ExitArray_definition is called when production array_definition is exited.
func (s *BaseCorundumListener) ExitArray_definition(ctx *Array_definitionContext) {}

// EnterArray_definition_elements is called when production array_definition_elements is entered.
func (s *BaseCorundumListener) EnterArray_definition_elements(ctx *Array_definition_elementsContext) {
}

// ExitArray_definition_elements is called when production array_definition_elements is exited.
func (s *BaseCorundumListener) ExitArray_definition_elements(ctx *Array_definition_elementsContext) {}

// EnterArray_selector is called when production array_selector is entered.
func (s *BaseCorundumListener) EnterArray_selector(ctx *Array_selectorContext) {}

// ExitArray_selector is called when production array_selector is exited.
func (s *BaseCorundumListener) ExitArray_selector(ctx *Array_selectorContext) {}

// EnterDynamic_result is called when production dynamic_result is entered.
func (s *BaseCorundumListener) EnterDynamic_result(ctx *Dynamic_resultContext) {}

// ExitDynamic_result is called when production dynamic_result is exited.
func (s *BaseCorundumListener) ExitDynamic_result(ctx *Dynamic_resultContext) {}

// EnterDynamic_ is called when production dynamic_ is entered.
func (s *BaseCorundumListener) EnterDynamic_(ctx *Dynamic_Context) {}

// ExitDynamic_ is called when production dynamic_ is exited.
func (s *BaseCorundumListener) ExitDynamic_(ctx *Dynamic_Context) {}

// EnterInt_result is called when production int_result is entered.
func (s *BaseCorundumListener) EnterInt_result(ctx *Int_resultContext) {}

// ExitInt_result is called when production int_result is exited.
func (s *BaseCorundumListener) ExitInt_result(ctx *Int_resultContext) {}

// EnterFloat_result is called when production float_result is entered.
func (s *BaseCorundumListener) EnterFloat_result(ctx *Float_resultContext) {}

// ExitFloat_result is called when production float_result is exited.
func (s *BaseCorundumListener) ExitFloat_result(ctx *Float_resultContext) {}

// EnterString_result is called when production string_result is entered.
func (s *BaseCorundumListener) EnterString_result(ctx *String_resultContext) {}

// ExitString_result is called when production string_result is exited.
func (s *BaseCorundumListener) ExitString_result(ctx *String_resultContext) {}

// EnterComparison_list is called when production comparison_list is entered.
func (s *BaseCorundumListener) EnterComparison_list(ctx *Comparison_listContext) {}

// ExitComparison_list is called when production comparison_list is exited.
func (s *BaseCorundumListener) ExitComparison_list(ctx *Comparison_listContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseCorundumListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseCorundumListener) ExitComparison(ctx *ComparisonContext) {}

// EnterComp_var is called when production comp_var is entered.
func (s *BaseCorundumListener) EnterComp_var(ctx *Comp_varContext) {}

// ExitComp_var is called when production comp_var is exited.
func (s *BaseCorundumListener) ExitComp_var(ctx *Comp_varContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseCorundumListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseCorundumListener) ExitLvalue(ctx *LvalueContext) {}

// EnterRvalue is called when production rvalue is entered.
func (s *BaseCorundumListener) EnterRvalue(ctx *RvalueContext) {}

// ExitRvalue is called when production rvalue is exited.
func (s *BaseCorundumListener) ExitRvalue(ctx *RvalueContext) {}

// EnterBreak_expression is called when production break_expression is entered.
func (s *BaseCorundumListener) EnterBreak_expression(ctx *Break_expressionContext) {}

// ExitBreak_expression is called when production break_expression is exited.
func (s *BaseCorundumListener) ExitBreak_expression(ctx *Break_expressionContext) {}

// EnterLiteral_t is called when production literal_t is entered.
func (s *BaseCorundumListener) EnterLiteral_t(ctx *Literal_tContext) {}

// ExitLiteral_t is called when production literal_t is exited.
func (s *BaseCorundumListener) ExitLiteral_t(ctx *Literal_tContext) {}

// EnterFloat_t is called when production float_t is entered.
func (s *BaseCorundumListener) EnterFloat_t(ctx *Float_tContext) {}

// ExitFloat_t is called when production float_t is exited.
func (s *BaseCorundumListener) ExitFloat_t(ctx *Float_tContext) {}

// EnterInt_t is called when production int_t is entered.
func (s *BaseCorundumListener) EnterInt_t(ctx *Int_tContext) {}

// ExitInt_t is called when production int_t is exited.
func (s *BaseCorundumListener) ExitInt_t(ctx *Int_tContext) {}

// EnterBool_t is called when production bool_t is entered.
func (s *BaseCorundumListener) EnterBool_t(ctx *Bool_tContext) {}

// ExitBool_t is called when production bool_t is exited.
func (s *BaseCorundumListener) ExitBool_t(ctx *Bool_tContext) {}

// EnterNil_t is called when production nil_t is entered.
func (s *BaseCorundumListener) EnterNil_t(ctx *Nil_tContext) {}

// ExitNil_t is called when production nil_t is exited.
func (s *BaseCorundumListener) ExitNil_t(ctx *Nil_tContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseCorundumListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseCorundumListener) ExitId_(ctx *Id_Context) {}

// EnterId_global is called when production id_global is entered.
func (s *BaseCorundumListener) EnterId_global(ctx *Id_globalContext) {}

// ExitId_global is called when production id_global is exited.
func (s *BaseCorundumListener) ExitId_global(ctx *Id_globalContext) {}

// EnterId_function is called when production id_function is entered.
func (s *BaseCorundumListener) EnterId_function(ctx *Id_functionContext) {}

// ExitId_function is called when production id_function is exited.
func (s *BaseCorundumListener) ExitId_function(ctx *Id_functionContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseCorundumListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseCorundumListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterElse_token is called when production else_token is entered.
func (s *BaseCorundumListener) EnterElse_token(ctx *Else_tokenContext) {}

// ExitElse_token is called when production else_token is exited.
func (s *BaseCorundumListener) ExitElse_token(ctx *Else_tokenContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BaseCorundumListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BaseCorundumListener) ExitCrlf(ctx *CrlfContext) {}
