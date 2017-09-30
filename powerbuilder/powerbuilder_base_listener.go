// Generated from powerbuilder.g4 by ANTLR 4.7.

package powerbuilder // powerbuilder
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepowerbuilderListener is a complete listener for a parse tree produced by powerbuilderParser.
type BasepowerbuilderListener struct{}

var _ powerbuilderListener = &BasepowerbuilderListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepowerbuilderListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepowerbuilderListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepowerbuilderListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepowerbuilderListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart_rule is called when production start_rule is entered.
func (s *BasepowerbuilderListener) EnterStart_rule(ctx *Start_ruleContext) {}

// ExitStart_rule is called when production start_rule is exited.
func (s *BasepowerbuilderListener) ExitStart_rule(ctx *Start_ruleContext) {}

// EnterHeader_rule is called when production header_rule is entered.
func (s *BasepowerbuilderListener) EnterHeader_rule(ctx *Header_ruleContext) {}

// ExitHeader_rule is called when production header_rule is exited.
func (s *BasepowerbuilderListener) ExitHeader_rule(ctx *Header_ruleContext) {}

// EnterBody_rule is called when production body_rule is entered.
func (s *BasepowerbuilderListener) EnterBody_rule(ctx *Body_ruleContext) {}

// ExitBody_rule is called when production body_rule is exited.
func (s *BasepowerbuilderListener) ExitBody_rule(ctx *Body_ruleContext) {}

// EnterExport_header is called when production export_header is entered.
func (s *BasepowerbuilderListener) EnterExport_header(ctx *Export_headerContext) {}

// ExitExport_header is called when production export_header is exited.
func (s *BasepowerbuilderListener) ExitExport_header(ctx *Export_headerContext) {}

// EnterRelease_information is called when production release_information is entered.
func (s *BasepowerbuilderListener) EnterRelease_information(ctx *Release_informationContext) {}

// ExitRelease_information is called when production release_information is exited.
func (s *BasepowerbuilderListener) ExitRelease_information(ctx *Release_informationContext) {}

// EnterWindow_property_line is called when production window_property_line is entered.
func (s *BasepowerbuilderListener) EnterWindow_property_line(ctx *Window_property_lineContext) {}

// ExitWindow_property_line is called when production window_property_line is exited.
func (s *BasepowerbuilderListener) ExitWindow_property_line(ctx *Window_property_lineContext) {}

// EnterWindow_property is called when production window_property is entered.
func (s *BasepowerbuilderListener) EnterWindow_property(ctx *Window_propertyContext) {}

// ExitWindow_property is called when production window_property is exited.
func (s *BasepowerbuilderListener) ExitWindow_property(ctx *Window_propertyContext) {}

// EnterWindow_property_attributes_sub is called when production window_property_attributes_sub is entered.
func (s *BasepowerbuilderListener) EnterWindow_property_attributes_sub(ctx *Window_property_attributes_subContext) {
}

// ExitWindow_property_attributes_sub is called when production window_property_attributes_sub is exited.
func (s *BasepowerbuilderListener) ExitWindow_property_attributes_sub(ctx *Window_property_attributes_subContext) {
}

// EnterWindow_property_attribute_sub is called when production window_property_attribute_sub is entered.
func (s *BasepowerbuilderListener) EnterWindow_property_attribute_sub(ctx *Window_property_attribute_subContext) {
}

// ExitWindow_property_attribute_sub is called when production window_property_attribute_sub is exited.
func (s *BasepowerbuilderListener) ExitWindow_property_attribute_sub(ctx *Window_property_attribute_subContext) {
}

// EnterAttribute_name is called when production attribute_name is entered.
func (s *BasepowerbuilderListener) EnterAttribute_name(ctx *Attribute_nameContext) {}

// ExitAttribute_name is called when production attribute_name is exited.
func (s *BasepowerbuilderListener) ExitAttribute_name(ctx *Attribute_nameContext) {}

// EnterAttribute_value is called when production attribute_value is entered.
func (s *BasepowerbuilderListener) EnterAttribute_value(ctx *Attribute_valueContext) {}

// ExitAttribute_value is called when production attribute_value is exited.
func (s *BasepowerbuilderListener) ExitAttribute_value(ctx *Attribute_valueContext) {}

// EnterForward_decl is called when production forward_decl is entered.
func (s *BasepowerbuilderListener) EnterForward_decl(ctx *Forward_declContext) {}

// ExitForward_decl is called when production forward_decl is exited.
func (s *BasepowerbuilderListener) ExitForward_decl(ctx *Forward_declContext) {}

// EnterDatatype_decl is called when production datatype_decl is entered.
func (s *BasepowerbuilderListener) EnterDatatype_decl(ctx *Datatype_declContext) {}

// ExitDatatype_decl is called when production datatype_decl is exited.
func (s *BasepowerbuilderListener) ExitDatatype_decl(ctx *Datatype_declContext) {}

// EnterType_variables_decl is called when production type_variables_decl is entered.
func (s *BasepowerbuilderListener) EnterType_variables_decl(ctx *Type_variables_declContext) {}

// ExitType_variables_decl is called when production type_variables_decl is exited.
func (s *BasepowerbuilderListener) ExitType_variables_decl(ctx *Type_variables_declContext) {}

// EnterGlobal_variables_decl is called when production global_variables_decl is entered.
func (s *BasepowerbuilderListener) EnterGlobal_variables_decl(ctx *Global_variables_declContext) {}

// ExitGlobal_variables_decl is called when production global_variables_decl is exited.
func (s *BasepowerbuilderListener) ExitGlobal_variables_decl(ctx *Global_variables_declContext) {}

// EnterVariable_decl_sub is called when production variable_decl_sub is entered.
func (s *BasepowerbuilderListener) EnterVariable_decl_sub(ctx *Variable_decl_subContext) {}

// ExitVariable_decl_sub is called when production variable_decl_sub is exited.
func (s *BasepowerbuilderListener) ExitVariable_decl_sub(ctx *Variable_decl_subContext) {}

// EnterVariable_decl is called when production variable_decl is entered.
func (s *BasepowerbuilderListener) EnterVariable_decl(ctx *Variable_declContext) {}

// ExitVariable_decl is called when production variable_decl is exited.
func (s *BasepowerbuilderListener) ExitVariable_decl(ctx *Variable_declContext) {}

// EnterDecimal_decl_sub is called when production decimal_decl_sub is entered.
func (s *BasepowerbuilderListener) EnterDecimal_decl_sub(ctx *Decimal_decl_subContext) {}

// ExitDecimal_decl_sub is called when production decimal_decl_sub is exited.
func (s *BasepowerbuilderListener) ExitDecimal_decl_sub(ctx *Decimal_decl_subContext) {}

// EnterArray_decl_sub is called when production array_decl_sub is entered.
func (s *BasepowerbuilderListener) EnterArray_decl_sub(ctx *Array_decl_subContext) {}

// ExitArray_decl_sub is called when production array_decl_sub is exited.
func (s *BasepowerbuilderListener) ExitArray_decl_sub(ctx *Array_decl_subContext) {}

// EnterConstant_decl_sub is called when production constant_decl_sub is entered.
func (s *BasepowerbuilderListener) EnterConstant_decl_sub(ctx *Constant_decl_subContext) {}

// ExitConstant_decl_sub is called when production constant_decl_sub is exited.
func (s *BasepowerbuilderListener) ExitConstant_decl_sub(ctx *Constant_decl_subContext) {}

// EnterConstant_decl is called when production constant_decl is entered.
func (s *BasepowerbuilderListener) EnterConstant_decl(ctx *Constant_declContext) {}

// ExitConstant_decl is called when production constant_decl is exited.
func (s *BasepowerbuilderListener) ExitConstant_decl(ctx *Constant_declContext) {}

// EnterFunction_forward_decl is called when production function_forward_decl is entered.
func (s *BasepowerbuilderListener) EnterFunction_forward_decl(ctx *Function_forward_declContext) {}

// ExitFunction_forward_decl is called when production function_forward_decl is exited.
func (s *BasepowerbuilderListener) ExitFunction_forward_decl(ctx *Function_forward_declContext) {}

// EnterParameter_sub is called when production parameter_sub is entered.
func (s *BasepowerbuilderListener) EnterParameter_sub(ctx *Parameter_subContext) {}

// ExitParameter_sub is called when production parameter_sub is exited.
func (s *BasepowerbuilderListener) ExitParameter_sub(ctx *Parameter_subContext) {}

// EnterParameters_list_sub is called when production parameters_list_sub is entered.
func (s *BasepowerbuilderListener) EnterParameters_list_sub(ctx *Parameters_list_subContext) {}

// ExitParameters_list_sub is called when production parameters_list_sub is exited.
func (s *BasepowerbuilderListener) ExitParameters_list_sub(ctx *Parameters_list_subContext) {}

// EnterFunctions_forward_decl is called when production functions_forward_decl is entered.
func (s *BasepowerbuilderListener) EnterFunctions_forward_decl(ctx *Functions_forward_declContext) {}

// ExitFunctions_forward_decl is called when production functions_forward_decl is exited.
func (s *BasepowerbuilderListener) ExitFunctions_forward_decl(ctx *Functions_forward_declContext) {}

// EnterFunction_body is called when production function_body is entered.
func (s *BasepowerbuilderListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BasepowerbuilderListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterOn_body is called when production on_body is entered.
func (s *BasepowerbuilderListener) EnterOn_body(ctx *On_bodyContext) {}

// ExitOn_body is called when production on_body is exited.
func (s *BasepowerbuilderListener) ExitOn_body(ctx *On_bodyContext) {}

// EnterEvent_forward_decl_sub is called when production event_forward_decl_sub is entered.
func (s *BasepowerbuilderListener) EnterEvent_forward_decl_sub(ctx *Event_forward_decl_subContext) {}

// ExitEvent_forward_decl_sub is called when production event_forward_decl_sub is exited.
func (s *BasepowerbuilderListener) ExitEvent_forward_decl_sub(ctx *Event_forward_decl_subContext) {}

// EnterEvent_forward_decl is called when production event_forward_decl is entered.
func (s *BasepowerbuilderListener) EnterEvent_forward_decl(ctx *Event_forward_declContext) {}

// ExitEvent_forward_decl is called when production event_forward_decl is exited.
func (s *BasepowerbuilderListener) ExitEvent_forward_decl(ctx *Event_forward_declContext) {}

// EnterEvent_body is called when production event_body is entered.
func (s *BasepowerbuilderListener) EnterEvent_body(ctx *Event_bodyContext) {}

// ExitEvent_body is called when production event_body is exited.
func (s *BasepowerbuilderListener) ExitEvent_body(ctx *Event_bodyContext) {}

// EnterAccess_modif is called when production access_modif is entered.
func (s *BasepowerbuilderListener) EnterAccess_modif(ctx *Access_modifContext) {}

// ExitAccess_modif is called when production access_modif is exited.
func (s *BasepowerbuilderListener) ExitAccess_modif(ctx *Access_modifContext) {}

// EnterAccess_modif_part is called when production access_modif_part is entered.
func (s *BasepowerbuilderListener) EnterAccess_modif_part(ctx *Access_modif_partContext) {}

// ExitAccess_modif_part is called when production access_modif_part is exited.
func (s *BasepowerbuilderListener) ExitAccess_modif_part(ctx *Access_modif_partContext) {}

// EnterScope_modif is called when production scope_modif is entered.
func (s *BasepowerbuilderListener) EnterScope_modif(ctx *Scope_modifContext) {}

// ExitScope_modif is called when production scope_modif is exited.
func (s *BasepowerbuilderListener) ExitScope_modif(ctx *Scope_modifContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasepowerbuilderListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasepowerbuilderListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasepowerbuilderListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasepowerbuilderListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterBoolean_expression is called when production boolean_expression is entered.
func (s *BasepowerbuilderListener) EnterBoolean_expression(ctx *Boolean_expressionContext) {}

// ExitBoolean_expression is called when production boolean_expression is exited.
func (s *BasepowerbuilderListener) ExitBoolean_expression(ctx *Boolean_expressionContext) {}

// EnterCondition_or is called when production condition_or is entered.
func (s *BasepowerbuilderListener) EnterCondition_or(ctx *Condition_orContext) {}

// ExitCondition_or is called when production condition_or is exited.
func (s *BasepowerbuilderListener) ExitCondition_or(ctx *Condition_orContext) {}

// EnterCondition_and is called when production condition_and is entered.
func (s *BasepowerbuilderListener) EnterCondition_and(ctx *Condition_andContext) {}

// ExitCondition_and is called when production condition_and is exited.
func (s *BasepowerbuilderListener) ExitCondition_and(ctx *Condition_andContext) {}

// EnterCondition_not is called when production condition_not is entered.
func (s *BasepowerbuilderListener) EnterCondition_not(ctx *Condition_notContext) {}

// ExitCondition_not is called when production condition_not is exited.
func (s *BasepowerbuilderListener) ExitCondition_not(ctx *Condition_notContext) {}

// EnterCondition_comparison is called when production condition_comparison is entered.
func (s *BasepowerbuilderListener) EnterCondition_comparison(ctx *Condition_comparisonContext) {}

// ExitCondition_comparison is called when production condition_comparison is exited.
func (s *BasepowerbuilderListener) ExitCondition_comparison(ctx *Condition_comparisonContext) {}

// EnterAdd_expr is called when production add_expr is entered.
func (s *BasepowerbuilderListener) EnterAdd_expr(ctx *Add_exprContext) {}

// ExitAdd_expr is called when production add_expr is exited.
func (s *BasepowerbuilderListener) ExitAdd_expr(ctx *Add_exprContext) {}

// EnterMul_expr is called when production mul_expr is entered.
func (s *BasepowerbuilderListener) EnterMul_expr(ctx *Mul_exprContext) {}

// ExitMul_expr is called when production mul_expr is exited.
func (s *BasepowerbuilderListener) ExitMul_expr(ctx *Mul_exprContext) {}

// EnterUnary_sign_expr is called when production unary_sign_expr is entered.
func (s *BasepowerbuilderListener) EnterUnary_sign_expr(ctx *Unary_sign_exprContext) {}

// ExitUnary_sign_expr is called when production unary_sign_expr is exited.
func (s *BasepowerbuilderListener) ExitUnary_sign_expr(ctx *Unary_sign_exprContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasepowerbuilderListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasepowerbuilderListener) ExitStatement(ctx *StatementContext) {}

// EnterStatement_sub is called when production statement_sub is entered.
func (s *BasepowerbuilderListener) EnterStatement_sub(ctx *Statement_subContext) {}

// ExitStatement_sub is called when production statement_sub is exited.
func (s *BasepowerbuilderListener) ExitStatement_sub(ctx *Statement_subContext) {}

// EnterAssignment_sub is called when production assignment_sub is entered.
func (s *BasepowerbuilderListener) EnterAssignment_sub(ctx *Assignment_subContext) {}

// ExitAssignment_sub is called when production assignment_sub is exited.
func (s *BasepowerbuilderListener) ExitAssignment_sub(ctx *Assignment_subContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BasepowerbuilderListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BasepowerbuilderListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterLvalue_sub is called when production lvalue_sub is entered.
func (s *BasepowerbuilderListener) EnterLvalue_sub(ctx *Lvalue_subContext) {}

// ExitLvalue_sub is called when production lvalue_sub is exited.
func (s *BasepowerbuilderListener) ExitLvalue_sub(ctx *Lvalue_subContext) {}

// EnterReturn_sub is called when production return_sub is entered.
func (s *BasepowerbuilderListener) EnterReturn_sub(ctx *Return_subContext) {}

// ExitReturn_sub is called when production return_sub is exited.
func (s *BasepowerbuilderListener) ExitReturn_sub(ctx *Return_subContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BasepowerbuilderListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BasepowerbuilderListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterFunction_call_expression_sub is called when production function_call_expression_sub is entered.
func (s *BasepowerbuilderListener) EnterFunction_call_expression_sub(ctx *Function_call_expression_subContext) {
}

// ExitFunction_call_expression_sub is called when production function_call_expression_sub is exited.
func (s *BasepowerbuilderListener) ExitFunction_call_expression_sub(ctx *Function_call_expression_subContext) {
}

// EnterFunction_virtual_call_expression_sub is called when production function_virtual_call_expression_sub is entered.
func (s *BasepowerbuilderListener) EnterFunction_virtual_call_expression_sub(ctx *Function_virtual_call_expression_subContext) {
}

// ExitFunction_virtual_call_expression_sub is called when production function_virtual_call_expression_sub is exited.
func (s *BasepowerbuilderListener) ExitFunction_virtual_call_expression_sub(ctx *Function_virtual_call_expression_subContext) {
}

// EnterOpen_call_sub is called when production open_call_sub is entered.
func (s *BasepowerbuilderListener) EnterOpen_call_sub(ctx *Open_call_subContext) {}

// ExitOpen_call_sub is called when production open_call_sub is exited.
func (s *BasepowerbuilderListener) ExitOpen_call_sub(ctx *Open_call_subContext) {}

// EnterClose_call_sub is called when production close_call_sub is entered.
func (s *BasepowerbuilderListener) EnterClose_call_sub(ctx *Close_call_subContext) {}

// ExitClose_call_sub is called when production close_call_sub is exited.
func (s *BasepowerbuilderListener) ExitClose_call_sub(ctx *Close_call_subContext) {}

// EnterFunction_call_statement is called when production function_call_statement is entered.
func (s *BasepowerbuilderListener) EnterFunction_call_statement(ctx *Function_call_statementContext) {}

// ExitFunction_call_statement is called when production function_call_statement is exited.
func (s *BasepowerbuilderListener) ExitFunction_call_statement(ctx *Function_call_statementContext) {}

// EnterSuper_call_sub is called when production super_call_sub is entered.
func (s *BasepowerbuilderListener) EnterSuper_call_sub(ctx *Super_call_subContext) {}

// ExitSuper_call_sub is called when production super_call_sub is exited.
func (s *BasepowerbuilderListener) ExitSuper_call_sub(ctx *Super_call_subContext) {}

// EnterSuper_call_statement is called when production super_call_statement is entered.
func (s *BasepowerbuilderListener) EnterSuper_call_statement(ctx *Super_call_statementContext) {}

// ExitSuper_call_statement is called when production super_call_statement is exited.
func (s *BasepowerbuilderListener) ExitSuper_call_statement(ctx *Super_call_statementContext) {}

// EnterEvent_call_statement_sub is called when production event_call_statement_sub is entered.
func (s *BasepowerbuilderListener) EnterEvent_call_statement_sub(ctx *Event_call_statement_subContext) {
}

// ExitEvent_call_statement_sub is called when production event_call_statement_sub is exited.
func (s *BasepowerbuilderListener) ExitEvent_call_statement_sub(ctx *Event_call_statement_subContext) {
}

// EnterEvent_call_statement is called when production event_call_statement is entered.
func (s *BasepowerbuilderListener) EnterEvent_call_statement(ctx *Event_call_statementContext) {}

// ExitEvent_call_statement is called when production event_call_statement is exited.
func (s *BasepowerbuilderListener) ExitEvent_call_statement(ctx *Event_call_statementContext) {}

// EnterCreate_call_sub is called when production create_call_sub is entered.
func (s *BasepowerbuilderListener) EnterCreate_call_sub(ctx *Create_call_subContext) {}

// ExitCreate_call_sub is called when production create_call_sub is exited.
func (s *BasepowerbuilderListener) ExitCreate_call_sub(ctx *Create_call_subContext) {}

// EnterCreate_call_statement is called when production create_call_statement is entered.
func (s *BasepowerbuilderListener) EnterCreate_call_statement(ctx *Create_call_statementContext) {}

// ExitCreate_call_statement is called when production create_call_statement is exited.
func (s *BasepowerbuilderListener) ExitCreate_call_statement(ctx *Create_call_statementContext) {}

// EnterDestroy_call_sub is called when production destroy_call_sub is entered.
func (s *BasepowerbuilderListener) EnterDestroy_call_sub(ctx *Destroy_call_subContext) {}

// ExitDestroy_call_sub is called when production destroy_call_sub is exited.
func (s *BasepowerbuilderListener) ExitDestroy_call_sub(ctx *Destroy_call_subContext) {}

// EnterDestroy_call_statement is called when production destroy_call_statement is entered.
func (s *BasepowerbuilderListener) EnterDestroy_call_statement(ctx *Destroy_call_statementContext) {}

// ExitDestroy_call_statement is called when production destroy_call_statement is exited.
func (s *BasepowerbuilderListener) ExitDestroy_call_statement(ctx *Destroy_call_statementContext) {}

// EnterFor_loop_statement is called when production for_loop_statement is entered.
func (s *BasepowerbuilderListener) EnterFor_loop_statement(ctx *For_loop_statementContext) {}

// ExitFor_loop_statement is called when production for_loop_statement is exited.
func (s *BasepowerbuilderListener) ExitFor_loop_statement(ctx *For_loop_statementContext) {}

// EnterDo_while_loop_statement is called when production do_while_loop_statement is entered.
func (s *BasepowerbuilderListener) EnterDo_while_loop_statement(ctx *Do_while_loop_statementContext) {}

// ExitDo_while_loop_statement is called when production do_while_loop_statement is exited.
func (s *BasepowerbuilderListener) ExitDo_while_loop_statement(ctx *Do_while_loop_statementContext) {}

// EnterDo_loop_while_statement is called when production do_loop_while_statement is entered.
func (s *BasepowerbuilderListener) EnterDo_loop_while_statement(ctx *Do_loop_while_statementContext) {}

// ExitDo_loop_while_statement is called when production do_loop_while_statement is exited.
func (s *BasepowerbuilderListener) ExitDo_loop_while_statement(ctx *Do_loop_while_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasepowerbuilderListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasepowerbuilderListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterIf_simple_statement is called when production if_simple_statement is entered.
func (s *BasepowerbuilderListener) EnterIf_simple_statement(ctx *If_simple_statementContext) {}

// ExitIf_simple_statement is called when production if_simple_statement is exited.
func (s *BasepowerbuilderListener) ExitIf_simple_statement(ctx *If_simple_statementContext) {}

// EnterContinue_sub is called when production continue_sub is entered.
func (s *BasepowerbuilderListener) EnterContinue_sub(ctx *Continue_subContext) {}

// ExitContinue_sub is called when production continue_sub is exited.
func (s *BasepowerbuilderListener) ExitContinue_sub(ctx *Continue_subContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BasepowerbuilderListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BasepowerbuilderListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterPost_event_sub is called when production post_event_sub is entered.
func (s *BasepowerbuilderListener) EnterPost_event_sub(ctx *Post_event_subContext) {}

// ExitPost_event_sub is called when production post_event_sub is exited.
func (s *BasepowerbuilderListener) ExitPost_event_sub(ctx *Post_event_subContext) {}

// EnterPost_event is called when production post_event is entered.
func (s *BasepowerbuilderListener) EnterPost_event(ctx *Post_eventContext) {}

// ExitPost_event is called when production post_event is exited.
func (s *BasepowerbuilderListener) ExitPost_event(ctx *Post_eventContext) {}

// EnterExit_statement_sub is called when production exit_statement_sub is entered.
func (s *BasepowerbuilderListener) EnterExit_statement_sub(ctx *Exit_statement_subContext) {}

// ExitExit_statement_sub is called when production exit_statement_sub is exited.
func (s *BasepowerbuilderListener) ExitExit_statement_sub(ctx *Exit_statement_subContext) {}

// EnterExit_statement is called when production exit_statement is entered.
func (s *BasepowerbuilderListener) EnterExit_statement(ctx *Exit_statementContext) {}

// ExitExit_statement is called when production exit_statement is exited.
func (s *BasepowerbuilderListener) ExitExit_statement(ctx *Exit_statementContext) {}

// EnterChoose_statement is called when production choose_statement is entered.
func (s *BasepowerbuilderListener) EnterChoose_statement(ctx *Choose_statementContext) {}

// ExitChoose_statement is called when production choose_statement is exited.
func (s *BasepowerbuilderListener) ExitChoose_statement(ctx *Choose_statementContext) {}

// EnterChoose_case_value_sub is called when production choose_case_value_sub is entered.
func (s *BasepowerbuilderListener) EnterChoose_case_value_sub(ctx *Choose_case_value_subContext) {}

// ExitChoose_case_value_sub is called when production choose_case_value_sub is exited.
func (s *BasepowerbuilderListener) ExitChoose_case_value_sub(ctx *Choose_case_value_subContext) {}

// EnterChoose_case_cond_sub is called when production choose_case_cond_sub is entered.
func (s *BasepowerbuilderListener) EnterChoose_case_cond_sub(ctx *Choose_case_cond_subContext) {}

// ExitChoose_case_cond_sub is called when production choose_case_cond_sub is exited.
func (s *BasepowerbuilderListener) ExitChoose_case_cond_sub(ctx *Choose_case_cond_subContext) {}

// EnterChoose_case_range_sub is called when production choose_case_range_sub is entered.
func (s *BasepowerbuilderListener) EnterChoose_case_range_sub(ctx *Choose_case_range_subContext) {}

// ExitChoose_case_range_sub is called when production choose_case_range_sub is exited.
func (s *BasepowerbuilderListener) ExitChoose_case_range_sub(ctx *Choose_case_range_subContext) {}

// EnterChoose_case_else_sub is called when production choose_case_else_sub is entered.
func (s *BasepowerbuilderListener) EnterChoose_case_else_sub(ctx *Choose_case_else_subContext) {}

// ExitChoose_case_else_sub is called when production choose_case_else_sub is exited.
func (s *BasepowerbuilderListener) ExitChoose_case_else_sub(ctx *Choose_case_else_subContext) {}

// EnterGoto_stat_sub is called when production goto_stat_sub is entered.
func (s *BasepowerbuilderListener) EnterGoto_stat_sub(ctx *Goto_stat_subContext) {}

// ExitGoto_stat_sub is called when production goto_stat_sub is exited.
func (s *BasepowerbuilderListener) ExitGoto_stat_sub(ctx *Goto_stat_subContext) {}

// EnterGoto_stat is called when production goto_stat is entered.
func (s *BasepowerbuilderListener) EnterGoto_stat(ctx *Goto_statContext) {}

// ExitGoto_stat is called when production goto_stat is exited.
func (s *BasepowerbuilderListener) ExitGoto_stat(ctx *Goto_statContext) {}

// EnterLabel_stat is called when production label_stat is entered.
func (s *BasepowerbuilderListener) EnterLabel_stat(ctx *Label_statContext) {}

// ExitLabel_stat is called when production label_stat is exited.
func (s *BasepowerbuilderListener) ExitLabel_stat(ctx *Label_statContext) {}

// EnterTry_catch_block is called when production try_catch_block is entered.
func (s *BasepowerbuilderListener) EnterTry_catch_block(ctx *Try_catch_blockContext) {}

// ExitTry_catch_block is called when production try_catch_block is exited.
func (s *BasepowerbuilderListener) ExitTry_catch_block(ctx *Try_catch_blockContext) {}

// EnterThrow_stat_sub is called when production throw_stat_sub is entered.
func (s *BasepowerbuilderListener) EnterThrow_stat_sub(ctx *Throw_stat_subContext) {}

// ExitThrow_stat_sub is called when production throw_stat_sub is exited.
func (s *BasepowerbuilderListener) ExitThrow_stat_sub(ctx *Throw_stat_subContext) {}

// EnterThrow_stat is called when production throw_stat is entered.
func (s *BasepowerbuilderListener) EnterThrow_stat(ctx *Throw_statContext) {}

// ExitThrow_stat is called when production throw_stat is exited.
func (s *BasepowerbuilderListener) ExitThrow_stat(ctx *Throw_statContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasepowerbuilderListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasepowerbuilderListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifier_name is called when production identifier_name is entered.
func (s *BasepowerbuilderListener) EnterIdentifier_name(ctx *Identifier_nameContext) {}

// ExitIdentifier_name is called when production identifier_name is exited.
func (s *BasepowerbuilderListener) ExitIdentifier_name(ctx *Identifier_nameContext) {}

// EnterIdentifier_name_ex is called when production identifier_name_ex is entered.
func (s *BasepowerbuilderListener) EnterIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// ExitIdentifier_name_ex is called when production identifier_name_ex is exited.
func (s *BasepowerbuilderListener) ExitIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// EnterAtom_sub is called when production atom_sub is entered.
func (s *BasepowerbuilderListener) EnterAtom_sub(ctx *Atom_subContext) {}

// ExitAtom_sub is called when production atom_sub is exited.
func (s *BasepowerbuilderListener) ExitAtom_sub(ctx *Atom_subContext) {}

// EnterAtom_sub_call1 is called when production atom_sub_call1 is entered.
func (s *BasepowerbuilderListener) EnterAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// ExitAtom_sub_call1 is called when production atom_sub_call1 is exited.
func (s *BasepowerbuilderListener) ExitAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// EnterAtom_sub_array1 is called when production atom_sub_array1 is entered.
func (s *BasepowerbuilderListener) EnterAtom_sub_array1(ctx *Atom_sub_array1Context) {}

// ExitAtom_sub_array1 is called when production atom_sub_array1 is exited.
func (s *BasepowerbuilderListener) ExitAtom_sub_array1(ctx *Atom_sub_array1Context) {}

// EnterAtom_sub_ref1 is called when production atom_sub_ref1 is entered.
func (s *BasepowerbuilderListener) EnterAtom_sub_ref1(ctx *Atom_sub_ref1Context) {}

// ExitAtom_sub_ref1 is called when production atom_sub_ref1 is exited.
func (s *BasepowerbuilderListener) ExitAtom_sub_ref1(ctx *Atom_sub_ref1Context) {}

// EnterAtom_sub_member1 is called when production atom_sub_member1 is entered.
func (s *BasepowerbuilderListener) EnterAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// ExitAtom_sub_member1 is called when production atom_sub_member1 is exited.
func (s *BasepowerbuilderListener) ExitAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// EnterAtom is called when production atom is entered.
func (s *BasepowerbuilderListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasepowerbuilderListener) ExitAtom(ctx *AtomContext) {}

// EnterSwallow_to_semi is called when production swallow_to_semi is entered.
func (s *BasepowerbuilderListener) EnterSwallow_to_semi(ctx *Swallow_to_semiContext) {}

// ExitSwallow_to_semi is called when production swallow_to_semi is exited.
func (s *BasepowerbuilderListener) ExitSwallow_to_semi(ctx *Swallow_to_semiContext) {}

// EnterSwallow_to_newline is called when production swallow_to_newline is entered.
func (s *BasepowerbuilderListener) EnterSwallow_to_newline(ctx *Swallow_to_newlineContext) {}

// ExitSwallow_to_newline is called when production swallow_to_newline is exited.
func (s *BasepowerbuilderListener) ExitSwallow_to_newline(ctx *Swallow_to_newlineContext) {}

// EnterArray_access_atom is called when production array_access_atom is entered.
func (s *BasepowerbuilderListener) EnterArray_access_atom(ctx *Array_access_atomContext) {}

// ExitArray_access_atom is called when production array_access_atom is exited.
func (s *BasepowerbuilderListener) ExitArray_access_atom(ctx *Array_access_atomContext) {}

// EnterNumeric_atom is called when production numeric_atom is entered.
func (s *BasepowerbuilderListener) EnterNumeric_atom(ctx *Numeric_atomContext) {}

// ExitNumeric_atom is called when production numeric_atom is exited.
func (s *BasepowerbuilderListener) ExitNumeric_atom(ctx *Numeric_atomContext) {}

// EnterBoolean_atom is called when production boolean_atom is entered.
func (s *BasepowerbuilderListener) EnterBoolean_atom(ctx *Boolean_atomContext) {}

// ExitBoolean_atom is called when production boolean_atom is exited.
func (s *BasepowerbuilderListener) ExitBoolean_atom(ctx *Boolean_atomContext) {}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BasepowerbuilderListener) EnterCast_expression(ctx *Cast_expressionContext) {}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BasepowerbuilderListener) ExitCast_expression(ctx *Cast_expressionContext) {}

// EnterData_type_sub is called when production data_type_sub is entered.
func (s *BasepowerbuilderListener) EnterData_type_sub(ctx *Data_type_subContext) {}

// ExitData_type_sub is called when production data_type_sub is exited.
func (s *BasepowerbuilderListener) ExitData_type_sub(ctx *Data_type_subContext) {}

// EnterData_type_name is called when production data_type_name is entered.
func (s *BasepowerbuilderListener) EnterData_type_name(ctx *Data_type_nameContext) {}

// ExitData_type_name is called when production data_type_name is exited.
func (s *BasepowerbuilderListener) ExitData_type_name(ctx *Data_type_nameContext) {}
