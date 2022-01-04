// Code generated from PowerBuilderParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerbuilder // PowerBuilderParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePowerBuilderParserListener is a complete listener for a parse tree produced by PowerBuilderParser.
type BasePowerBuilderParserListener struct{}

var _ PowerBuilderParserListener = &BasePowerBuilderParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePowerBuilderParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePowerBuilderParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePowerBuilderParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePowerBuilderParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart_rule is called when production start_rule is entered.
func (s *BasePowerBuilderParserListener) EnterStart_rule(ctx *Start_ruleContext) {}

// ExitStart_rule is called when production start_rule is exited.
func (s *BasePowerBuilderParserListener) ExitStart_rule(ctx *Start_ruleContext) {}

// EnterBody_rule is called when production body_rule is entered.
func (s *BasePowerBuilderParserListener) EnterBody_rule(ctx *Body_ruleContext) {}

// ExitBody_rule is called when production body_rule is exited.
func (s *BasePowerBuilderParserListener) ExitBody_rule(ctx *Body_ruleContext) {}

// EnterForward_decl is called when production forward_decl is entered.
func (s *BasePowerBuilderParserListener) EnterForward_decl(ctx *Forward_declContext) {}

// ExitForward_decl is called when production forward_decl is exited.
func (s *BasePowerBuilderParserListener) ExitForward_decl(ctx *Forward_declContext) {}

// EnterDatatype_decl is called when production datatype_decl is entered.
func (s *BasePowerBuilderParserListener) EnterDatatype_decl(ctx *Datatype_declContext) {}

// ExitDatatype_decl is called when production datatype_decl is exited.
func (s *BasePowerBuilderParserListener) ExitDatatype_decl(ctx *Datatype_declContext) {}

// EnterType_variables_decl is called when production type_variables_decl is entered.
func (s *BasePowerBuilderParserListener) EnterType_variables_decl(ctx *Type_variables_declContext) {}

// ExitType_variables_decl is called when production type_variables_decl is exited.
func (s *BasePowerBuilderParserListener) ExitType_variables_decl(ctx *Type_variables_declContext) {}

// EnterGlobal_variables_decl is called when production global_variables_decl is entered.
func (s *BasePowerBuilderParserListener) EnterGlobal_variables_decl(ctx *Global_variables_declContext) {
}

// ExitGlobal_variables_decl is called when production global_variables_decl is exited.
func (s *BasePowerBuilderParserListener) ExitGlobal_variables_decl(ctx *Global_variables_declContext) {
}

// EnterVariable_decl is called when production variable_decl is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl(ctx *Variable_declContext) {}

// ExitVariable_decl is called when production variable_decl is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl(ctx *Variable_declContext) {}

// EnterVariable_decl_sub is called when production variable_decl_sub is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl_sub(ctx *Variable_decl_subContext) {}

// ExitVariable_decl_sub is called when production variable_decl_sub is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl_sub(ctx *Variable_decl_subContext) {}

// EnterVariable_decl_sub0 is called when production variable_decl_sub0 is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl_sub0(ctx *Variable_decl_sub0Context) {}

// ExitVariable_decl_sub0 is called when production variable_decl_sub0 is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl_sub0(ctx *Variable_decl_sub0Context) {}

// EnterVariable_decl_sub1 is called when production variable_decl_sub1 is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl_sub1(ctx *Variable_decl_sub1Context) {}

// ExitVariable_decl_sub1 is called when production variable_decl_sub1 is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl_sub1(ctx *Variable_decl_sub1Context) {}

// EnterVariable_decl_sub2 is called when production variable_decl_sub2 is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl_sub2(ctx *Variable_decl_sub2Context) {}

// ExitVariable_decl_sub2 is called when production variable_decl_sub2 is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl_sub2(ctx *Variable_decl_sub2Context) {}

// EnterVariable_decl_event is called when production variable_decl_event is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_decl_event(ctx *Variable_decl_eventContext) {}

// ExitVariable_decl_event is called when production variable_decl_event is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_decl_event(ctx *Variable_decl_eventContext) {}

// EnterDecimal_decl_sub is called when production decimal_decl_sub is entered.
func (s *BasePowerBuilderParserListener) EnterDecimal_decl_sub(ctx *Decimal_decl_subContext) {}

// ExitDecimal_decl_sub is called when production decimal_decl_sub is exited.
func (s *BasePowerBuilderParserListener) ExitDecimal_decl_sub(ctx *Decimal_decl_subContext) {}

// EnterArray_decl_sub is called when production array_decl_sub is entered.
func (s *BasePowerBuilderParserListener) EnterArray_decl_sub(ctx *Array_decl_subContext) {}

// ExitArray_decl_sub is called when production array_decl_sub is exited.
func (s *BasePowerBuilderParserListener) ExitArray_decl_sub(ctx *Array_decl_subContext) {}

// EnterConstant_decl_sub is called when production constant_decl_sub is entered.
func (s *BasePowerBuilderParserListener) EnterConstant_decl_sub(ctx *Constant_decl_subContext) {}

// ExitConstant_decl_sub is called when production constant_decl_sub is exited.
func (s *BasePowerBuilderParserListener) ExitConstant_decl_sub(ctx *Constant_decl_subContext) {}

// EnterConstant_decl is called when production constant_decl is entered.
func (s *BasePowerBuilderParserListener) EnterConstant_decl(ctx *Constant_declContext) {}

// ExitConstant_decl is called when production constant_decl is exited.
func (s *BasePowerBuilderParserListener) ExitConstant_decl(ctx *Constant_declContext) {}

// EnterFunction_forward_decl is called when production function_forward_decl is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_forward_decl(ctx *Function_forward_declContext) {
}

// ExitFunction_forward_decl is called when production function_forward_decl is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_forward_decl(ctx *Function_forward_declContext) {
}

// EnterFunction_forward_decl_alias is called when production function_forward_decl_alias is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_forward_decl_alias(ctx *Function_forward_decl_aliasContext) {
}

// ExitFunction_forward_decl_alias is called when production function_forward_decl_alias is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_forward_decl_alias(ctx *Function_forward_decl_aliasContext) {
}

// EnterParameter_sub is called when production parameter_sub is entered.
func (s *BasePowerBuilderParserListener) EnterParameter_sub(ctx *Parameter_subContext) {}

// ExitParameter_sub is called when production parameter_sub is exited.
func (s *BasePowerBuilderParserListener) ExitParameter_sub(ctx *Parameter_subContext) {}

// EnterParameters_list_sub is called when production parameters_list_sub is entered.
func (s *BasePowerBuilderParserListener) EnterParameters_list_sub(ctx *Parameters_list_subContext) {}

// ExitParameters_list_sub is called when production parameters_list_sub is exited.
func (s *BasePowerBuilderParserListener) ExitParameters_list_sub(ctx *Parameters_list_subContext) {}

// EnterFunctions_forward_decl is called when production functions_forward_decl is entered.
func (s *BasePowerBuilderParserListener) EnterFunctions_forward_decl(ctx *Functions_forward_declContext) {
}

// ExitFunctions_forward_decl is called when production functions_forward_decl is exited.
func (s *BasePowerBuilderParserListener) ExitFunctions_forward_decl(ctx *Functions_forward_declContext) {
}

// EnterFunction_body is called when production function_body is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterOn_body is called when production on_body is entered.
func (s *BasePowerBuilderParserListener) EnterOn_body(ctx *On_bodyContext) {}

// ExitOn_body is called when production on_body is exited.
func (s *BasePowerBuilderParserListener) ExitOn_body(ctx *On_bodyContext) {}

// EnterEvent_forward_decl is called when production event_forward_decl is entered.
func (s *BasePowerBuilderParserListener) EnterEvent_forward_decl(ctx *Event_forward_declContext) {}

// ExitEvent_forward_decl is called when production event_forward_decl is exited.
func (s *BasePowerBuilderParserListener) ExitEvent_forward_decl(ctx *Event_forward_declContext) {}

// EnterEvent_body is called when production event_body is entered.
func (s *BasePowerBuilderParserListener) EnterEvent_body(ctx *Event_bodyContext) {}

// ExitEvent_body is called when production event_body is exited.
func (s *BasePowerBuilderParserListener) ExitEvent_body(ctx *Event_bodyContext) {}

// EnterAccess_type is called when production access_type is entered.
func (s *BasePowerBuilderParserListener) EnterAccess_type(ctx *Access_typeContext) {}

// ExitAccess_type is called when production access_type is exited.
func (s *BasePowerBuilderParserListener) ExitAccess_type(ctx *Access_typeContext) {}

// EnterAccess_modif is called when production access_modif is entered.
func (s *BasePowerBuilderParserListener) EnterAccess_modif(ctx *Access_modifContext) {}

// ExitAccess_modif is called when production access_modif is exited.
func (s *BasePowerBuilderParserListener) ExitAccess_modif(ctx *Access_modifContext) {}

// EnterAccess_modif_part is called when production access_modif_part is entered.
func (s *BasePowerBuilderParserListener) EnterAccess_modif_part(ctx *Access_modif_partContext) {}

// ExitAccess_modif_part is called when production access_modif_part is exited.
func (s *BasePowerBuilderParserListener) ExitAccess_modif_part(ctx *Access_modif_partContext) {}

// EnterScope_modif is called when production scope_modif is entered.
func (s *BasePowerBuilderParserListener) EnterScope_modif(ctx *Scope_modifContext) {}

// ExitScope_modif is called when production scope_modif is exited.
func (s *BasePowerBuilderParserListener) ExitScope_modif(ctx *Scope_modifContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePowerBuilderParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePowerBuilderParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterValue is called when production value is entered.
func (s *BasePowerBuilderParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePowerBuilderParserListener) ExitValue(ctx *ValueContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasePowerBuilderParserListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasePowerBuilderParserListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterBoolean_expression is called when production boolean_expression is entered.
func (s *BasePowerBuilderParserListener) EnterBoolean_expression(ctx *Boolean_expressionContext) {}

// ExitBoolean_expression is called when production boolean_expression is exited.
func (s *BasePowerBuilderParserListener) ExitBoolean_expression(ctx *Boolean_expressionContext) {}

// EnterCondition_or is called when production condition_or is entered.
func (s *BasePowerBuilderParserListener) EnterCondition_or(ctx *Condition_orContext) {}

// ExitCondition_or is called when production condition_or is exited.
func (s *BasePowerBuilderParserListener) ExitCondition_or(ctx *Condition_orContext) {}

// EnterCondition_and is called when production condition_and is entered.
func (s *BasePowerBuilderParserListener) EnterCondition_and(ctx *Condition_andContext) {}

// ExitCondition_and is called when production condition_and is exited.
func (s *BasePowerBuilderParserListener) ExitCondition_and(ctx *Condition_andContext) {}

// EnterCondition_not is called when production condition_not is entered.
func (s *BasePowerBuilderParserListener) EnterCondition_not(ctx *Condition_notContext) {}

// ExitCondition_not is called when production condition_not is exited.
func (s *BasePowerBuilderParserListener) ExitCondition_not(ctx *Condition_notContext) {}

// EnterCondition_comparison is called when production condition_comparison is entered.
func (s *BasePowerBuilderParserListener) EnterCondition_comparison(ctx *Condition_comparisonContext) {
}

// ExitCondition_comparison is called when production condition_comparison is exited.
func (s *BasePowerBuilderParserListener) ExitCondition_comparison(ctx *Condition_comparisonContext) {}

// EnterAdd_expr is called when production add_expr is entered.
func (s *BasePowerBuilderParserListener) EnterAdd_expr(ctx *Add_exprContext) {}

// ExitAdd_expr is called when production add_expr is exited.
func (s *BasePowerBuilderParserListener) ExitAdd_expr(ctx *Add_exprContext) {}

// EnterMul_expr is called when production mul_expr is entered.
func (s *BasePowerBuilderParserListener) EnterMul_expr(ctx *Mul_exprContext) {}

// ExitMul_expr is called when production mul_expr is exited.
func (s *BasePowerBuilderParserListener) ExitMul_expr(ctx *Mul_exprContext) {}

// EnterUnary_sign_expr is called when production unary_sign_expr is entered.
func (s *BasePowerBuilderParserListener) EnterUnary_sign_expr(ctx *Unary_sign_exprContext) {}

// ExitUnary_sign_expr is called when production unary_sign_expr is exited.
func (s *BasePowerBuilderParserListener) ExitUnary_sign_expr(ctx *Unary_sign_exprContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePowerBuilderParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePowerBuilderParserListener) ExitStatement(ctx *StatementContext) {}

// EnterPublic_statement is called when production public_statement is entered.
func (s *BasePowerBuilderParserListener) EnterPublic_statement(ctx *Public_statementContext) {}

// ExitPublic_statement is called when production public_statement is exited.
func (s *BasePowerBuilderParserListener) ExitPublic_statement(ctx *Public_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BasePowerBuilderParserListener) EnterThrow_statement(ctx *Throw_statementContext) {}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BasePowerBuilderParserListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *BasePowerBuilderParserListener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *BasePowerBuilderParserListener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterStatement_sub is called when production statement_sub is entered.
func (s *BasePowerBuilderParserListener) EnterStatement_sub(ctx *Statement_subContext) {}

// ExitStatement_sub is called when production statement_sub is exited.
func (s *BasePowerBuilderParserListener) ExitStatement_sub(ctx *Statement_subContext) {}

// EnterTry_catch_statement is called when production try_catch_statement is entered.
func (s *BasePowerBuilderParserListener) EnterTry_catch_statement(ctx *Try_catch_statementContext) {}

// ExitTry_catch_statement is called when production try_catch_statement is exited.
func (s *BasePowerBuilderParserListener) ExitTry_catch_statement(ctx *Try_catch_statementContext) {}

// EnterSql_statement is called when production sql_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_statement(ctx *Sql_statementContext) {}

// ExitSql_statement is called when production sql_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_statement(ctx *Sql_statementContext) {}

// EnterSql_insert_statement is called when production sql_insert_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_insert_statement(ctx *Sql_insert_statementContext) {
}

// ExitSql_insert_statement is called when production sql_insert_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_insert_statement(ctx *Sql_insert_statementContext) {}

// EnterSql_values is called when production sql_values is entered.
func (s *BasePowerBuilderParserListener) EnterSql_values(ctx *Sql_valuesContext) {}

// ExitSql_values is called when production sql_values is exited.
func (s *BasePowerBuilderParserListener) ExitSql_values(ctx *Sql_valuesContext) {}

// EnterSql_delete_statement is called when production sql_delete_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_delete_statement(ctx *Sql_delete_statementContext) {
}

// ExitSql_delete_statement is called when production sql_delete_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_delete_statement(ctx *Sql_delete_statementContext) {}

// EnterSql_select_statement is called when production sql_select_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_select_statement(ctx *Sql_select_statementContext) {
}

// ExitSql_select_statement is called when production sql_select_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_select_statement(ctx *Sql_select_statementContext) {}

// EnterSql_update_statement is called when production sql_update_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_update_statement(ctx *Sql_update_statementContext) {
}

// ExitSql_update_statement is called when production sql_update_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_update_statement(ctx *Sql_update_statementContext) {}

// EnterSql_connect_statement is called when production sql_connect_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_connect_statement(ctx *Sql_connect_statementContext) {
}

// ExitSql_connect_statement is called when production sql_connect_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_connect_statement(ctx *Sql_connect_statementContext) {
}

// EnterSet_value is called when production set_value is entered.
func (s *BasePowerBuilderParserListener) EnterSet_value(ctx *Set_valueContext) {}

// ExitSet_value is called when production set_value is exited.
func (s *BasePowerBuilderParserListener) ExitSet_value(ctx *Set_valueContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BasePowerBuilderParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BasePowerBuilderParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BasePowerBuilderParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BasePowerBuilderParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSql_commit_statement is called when production sql_commit_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSql_commit_statement(ctx *Sql_commit_statementContext) {
}

// ExitSql_commit_statement is called when production sql_commit_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSql_commit_statement(ctx *Sql_commit_statementContext) {}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *BasePowerBuilderParserListener) EnterExecute_statement(ctx *Execute_statementContext) {}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *BasePowerBuilderParserListener) ExitExecute_statement(ctx *Execute_statementContext) {}

// EnterClose_sql_statement is called when production close_sql_statement is entered.
func (s *BasePowerBuilderParserListener) EnterClose_sql_statement(ctx *Close_sql_statementContext) {}

// ExitClose_sql_statement is called when production close_sql_statement is exited.
func (s *BasePowerBuilderParserListener) ExitClose_sql_statement(ctx *Close_sql_statementContext) {}

// EnterDeclare_procedure_statement is called when production declare_procedure_statement is entered.
func (s *BasePowerBuilderParserListener) EnterDeclare_procedure_statement(ctx *Declare_procedure_statementContext) {
}

// ExitDeclare_procedure_statement is called when production declare_procedure_statement is exited.
func (s *BasePowerBuilderParserListener) ExitDeclare_procedure_statement(ctx *Declare_procedure_statementContext) {
}

// EnterDeclare_cursor_statement is called when production declare_cursor_statement is entered.
func (s *BasePowerBuilderParserListener) EnterDeclare_cursor_statement(ctx *Declare_cursor_statementContext) {
}

// ExitDeclare_cursor_statement is called when production declare_cursor_statement is exited.
func (s *BasePowerBuilderParserListener) ExitDeclare_cursor_statement(ctx *Declare_cursor_statementContext) {
}

// EnterOpen_cursor_statement is called when production open_cursor_statement is entered.
func (s *BasePowerBuilderParserListener) EnterOpen_cursor_statement(ctx *Open_cursor_statementContext) {
}

// ExitOpen_cursor_statement is called when production open_cursor_statement is exited.
func (s *BasePowerBuilderParserListener) ExitOpen_cursor_statement(ctx *Open_cursor_statementContext) {
}

// EnterClose_cursor_statement is called when production close_cursor_statement is entered.
func (s *BasePowerBuilderParserListener) EnterClose_cursor_statement(ctx *Close_cursor_statementContext) {
}

// ExitClose_cursor_statement is called when production close_cursor_statement is exited.
func (s *BasePowerBuilderParserListener) ExitClose_cursor_statement(ctx *Close_cursor_statementContext) {
}

// EnterFetch_into_statement is called when production fetch_into_statement is entered.
func (s *BasePowerBuilderParserListener) EnterFetch_into_statement(ctx *Fetch_into_statementContext) {
}

// ExitFetch_into_statement is called when production fetch_into_statement is exited.
func (s *BasePowerBuilderParserListener) ExitFetch_into_statement(ctx *Fetch_into_statementContext) {}

// EnterPrepare_sql_stateent is called when production prepare_sql_stateent is entered.
func (s *BasePowerBuilderParserListener) EnterPrepare_sql_stateent(ctx *Prepare_sql_stateentContext) {
}

// ExitPrepare_sql_stateent is called when production prepare_sql_stateent is exited.
func (s *BasePowerBuilderParserListener) ExitPrepare_sql_stateent(ctx *Prepare_sql_stateentContext) {}

// EnterIncrement_decrement_statement is called when production increment_decrement_statement is entered.
func (s *BasePowerBuilderParserListener) EnterIncrement_decrement_statement(ctx *Increment_decrement_statementContext) {
}

// ExitIncrement_decrement_statement is called when production increment_decrement_statement is exited.
func (s *BasePowerBuilderParserListener) ExitIncrement_decrement_statement(ctx *Increment_decrement_statementContext) {
}

// EnterAssignment_rhs is called when production assignment_rhs is entered.
func (s *BasePowerBuilderParserListener) EnterAssignment_rhs(ctx *Assignment_rhsContext) {}

// ExitAssignment_rhs is called when production assignment_rhs is exited.
func (s *BasePowerBuilderParserListener) ExitAssignment_rhs(ctx *Assignment_rhsContext) {}

// EnterDescribe_function_call is called when production describe_function_call is entered.
func (s *BasePowerBuilderParserListener) EnterDescribe_function_call(ctx *Describe_function_callContext) {
}

// ExitDescribe_function_call is called when production describe_function_call is exited.
func (s *BasePowerBuilderParserListener) ExitDescribe_function_call(ctx *Describe_function_callContext) {
}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BasePowerBuilderParserListener) EnterAssignment_statement(ctx *Assignment_statementContext) {
}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BasePowerBuilderParserListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BasePowerBuilderParserListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BasePowerBuilderParserListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BasePowerBuilderParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BasePowerBuilderParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterFunction_call_expression_sub is called when production function_call_expression_sub is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_call_expression_sub(ctx *Function_call_expression_subContext) {
}

// ExitFunction_call_expression_sub is called when production function_call_expression_sub is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_call_expression_sub(ctx *Function_call_expression_subContext) {
}

// EnterFunction_name is called when production function_name is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_event_call is called when production function_event_call is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_event_call(ctx *Function_event_callContext) {}

// ExitFunction_event_call is called when production function_event_call is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_event_call(ctx *Function_event_callContext) {}

// EnterFunction_virtual_call_expression_sub is called when production function_virtual_call_expression_sub is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_virtual_call_expression_sub(ctx *Function_virtual_call_expression_subContext) {
}

// ExitFunction_virtual_call_expression_sub is called when production function_virtual_call_expression_sub is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_virtual_call_expression_sub(ctx *Function_virtual_call_expression_subContext) {
}

// EnterOpen_call_sub is called when production open_call_sub is entered.
func (s *BasePowerBuilderParserListener) EnterOpen_call_sub(ctx *Open_call_subContext) {}

// ExitOpen_call_sub is called when production open_call_sub is exited.
func (s *BasePowerBuilderParserListener) ExitOpen_call_sub(ctx *Open_call_subContext) {}

// EnterClose_call_sub is called when production close_call_sub is entered.
func (s *BasePowerBuilderParserListener) EnterClose_call_sub(ctx *Close_call_subContext) {}

// ExitClose_call_sub is called when production close_call_sub is exited.
func (s *BasePowerBuilderParserListener) ExitClose_call_sub(ctx *Close_call_subContext) {}

// EnterFunction_call_statement is called when production function_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterFunction_call_statement(ctx *Function_call_statementContext) {
}

// ExitFunction_call_statement is called when production function_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitFunction_call_statement(ctx *Function_call_statementContext) {
}

// EnterAncestor_function_call is called when production ancestor_function_call is entered.
func (s *BasePowerBuilderParserListener) EnterAncestor_function_call(ctx *Ancestor_function_callContext) {
}

// ExitAncestor_function_call is called when production ancestor_function_call is exited.
func (s *BasePowerBuilderParserListener) ExitAncestor_function_call(ctx *Ancestor_function_callContext) {
}

// EnterCall_statement is called when production call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterCall_statement(ctx *Call_statementContext) {}

// ExitCall_statement is called when production call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitCall_statement(ctx *Call_statementContext) {}

// EnterSuper_call_statement is called when production super_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterSuper_call_statement(ctx *Super_call_statementContext) {
}

// ExitSuper_call_statement is called when production super_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitSuper_call_statement(ctx *Super_call_statementContext) {}

// EnterAncestor_event_call_statement is called when production ancestor_event_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterAncestor_event_call_statement(ctx *Ancestor_event_call_statementContext) {
}

// ExitAncestor_event_call_statement is called when production ancestor_event_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitAncestor_event_call_statement(ctx *Ancestor_event_call_statementContext) {
}

// EnterEvent_call_statement_sub is called when production event_call_statement_sub is entered.
func (s *BasePowerBuilderParserListener) EnterEvent_call_statement_sub(ctx *Event_call_statement_subContext) {
}

// ExitEvent_call_statement_sub is called when production event_call_statement_sub is exited.
func (s *BasePowerBuilderParserListener) ExitEvent_call_statement_sub(ctx *Event_call_statement_subContext) {
}

// EnterEvent_call_statement is called when production event_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterEvent_call_statement(ctx *Event_call_statementContext) {
}

// ExitEvent_call_statement is called when production event_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitEvent_call_statement(ctx *Event_call_statementContext) {}

// EnterCreate_call_sub is called when production create_call_sub is entered.
func (s *BasePowerBuilderParserListener) EnterCreate_call_sub(ctx *Create_call_subContext) {}

// ExitCreate_call_sub is called when production create_call_sub is exited.
func (s *BasePowerBuilderParserListener) ExitCreate_call_sub(ctx *Create_call_subContext) {}

// EnterCreate_call_statement is called when production create_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterCreate_call_statement(ctx *Create_call_statementContext) {
}

// ExitCreate_call_statement is called when production create_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitCreate_call_statement(ctx *Create_call_statementContext) {
}

// EnterDestroy_call_sub is called when production destroy_call_sub is entered.
func (s *BasePowerBuilderParserListener) EnterDestroy_call_sub(ctx *Destroy_call_subContext) {}

// ExitDestroy_call_sub is called when production destroy_call_sub is exited.
func (s *BasePowerBuilderParserListener) ExitDestroy_call_sub(ctx *Destroy_call_subContext) {}

// EnterDestroy_call_statement is called when production destroy_call_statement is entered.
func (s *BasePowerBuilderParserListener) EnterDestroy_call_statement(ctx *Destroy_call_statementContext) {
}

// ExitDestroy_call_statement is called when production destroy_call_statement is exited.
func (s *BasePowerBuilderParserListener) ExitDestroy_call_statement(ctx *Destroy_call_statementContext) {
}

// EnterFor_loop_statement is called when production for_loop_statement is entered.
func (s *BasePowerBuilderParserListener) EnterFor_loop_statement(ctx *For_loop_statementContext) {}

// ExitFor_loop_statement is called when production for_loop_statement is exited.
func (s *BasePowerBuilderParserListener) ExitFor_loop_statement(ctx *For_loop_statementContext) {}

// EnterDo_while_loop_statement is called when production do_while_loop_statement is entered.
func (s *BasePowerBuilderParserListener) EnterDo_while_loop_statement(ctx *Do_while_loop_statementContext) {
}

// ExitDo_while_loop_statement is called when production do_while_loop_statement is exited.
func (s *BasePowerBuilderParserListener) ExitDo_while_loop_statement(ctx *Do_while_loop_statementContext) {
}

// EnterDo_loop_while_statement is called when production do_loop_while_statement is entered.
func (s *BasePowerBuilderParserListener) EnterDo_loop_while_statement(ctx *Do_loop_while_statementContext) {
}

// ExitDo_loop_while_statement is called when production do_loop_while_statement is exited.
func (s *BasePowerBuilderParserListener) ExitDo_loop_while_statement(ctx *Do_loop_while_statementContext) {
}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasePowerBuilderParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasePowerBuilderParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElseif_statement is called when production elseif_statement is entered.
func (s *BasePowerBuilderParserListener) EnterElseif_statement(ctx *Elseif_statementContext) {}

// ExitElseif_statement is called when production elseif_statement is exited.
func (s *BasePowerBuilderParserListener) ExitElseif_statement(ctx *Elseif_statementContext) {}

// EnterElse_statement is called when production else_statement is entered.
func (s *BasePowerBuilderParserListener) EnterElse_statement(ctx *Else_statementContext) {}

// ExitElse_statement is called when production else_statement is exited.
func (s *BasePowerBuilderParserListener) ExitElse_statement(ctx *Else_statementContext) {}

// EnterIf_simple_statement is called when production if_simple_statement is entered.
func (s *BasePowerBuilderParserListener) EnterIf_simple_statement(ctx *If_simple_statementContext) {}

// ExitIf_simple_statement is called when production if_simple_statement is exited.
func (s *BasePowerBuilderParserListener) ExitIf_simple_statement(ctx *If_simple_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BasePowerBuilderParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BasePowerBuilderParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterContinue_sub is called when production continue_sub is entered.
func (s *BasePowerBuilderParserListener) EnterContinue_sub(ctx *Continue_subContext) {}

// ExitContinue_sub is called when production continue_sub is exited.
func (s *BasePowerBuilderParserListener) ExitContinue_sub(ctx *Continue_subContext) {}

// EnterPost_event is called when production post_event is entered.
func (s *BasePowerBuilderParserListener) EnterPost_event(ctx *Post_eventContext) {}

// ExitPost_event is called when production post_event is exited.
func (s *BasePowerBuilderParserListener) ExitPost_event(ctx *Post_eventContext) {}

// EnterExit_statement is called when production exit_statement is entered.
func (s *BasePowerBuilderParserListener) EnterExit_statement(ctx *Exit_statementContext) {}

// ExitExit_statement is called when production exit_statement is exited.
func (s *BasePowerBuilderParserListener) ExitExit_statement(ctx *Exit_statementContext) {}

// EnterChoose_statement is called when production choose_statement is entered.
func (s *BasePowerBuilderParserListener) EnterChoose_statement(ctx *Choose_statementContext) {}

// ExitChoose_statement is called when production choose_statement is exited.
func (s *BasePowerBuilderParserListener) ExitChoose_statement(ctx *Choose_statementContext) {}

// EnterChoose_case_value_sub is called when production choose_case_value_sub is entered.
func (s *BasePowerBuilderParserListener) EnterChoose_case_value_sub(ctx *Choose_case_value_subContext) {
}

// ExitChoose_case_value_sub is called when production choose_case_value_sub is exited.
func (s *BasePowerBuilderParserListener) ExitChoose_case_value_sub(ctx *Choose_case_value_subContext) {
}

// EnterChoose_case_cond_sub is called when production choose_case_cond_sub is entered.
func (s *BasePowerBuilderParserListener) EnterChoose_case_cond_sub(ctx *Choose_case_cond_subContext) {
}

// ExitChoose_case_cond_sub is called when production choose_case_cond_sub is exited.
func (s *BasePowerBuilderParserListener) ExitChoose_case_cond_sub(ctx *Choose_case_cond_subContext) {}

// EnterChoose_case_else_sub is called when production choose_case_else_sub is entered.
func (s *BasePowerBuilderParserListener) EnterChoose_case_else_sub(ctx *Choose_case_else_subContext) {
}

// ExitChoose_case_else_sub is called when production choose_case_else_sub is exited.
func (s *BasePowerBuilderParserListener) ExitChoose_case_else_sub(ctx *Choose_case_else_subContext) {}

// EnterLabel_stat is called when production label_stat is entered.
func (s *BasePowerBuilderParserListener) EnterLabel_stat(ctx *Label_statContext) {}

// ExitLabel_stat is called when production label_stat is exited.
func (s *BasePowerBuilderParserListener) ExitLabel_stat(ctx *Label_statContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePowerBuilderParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePowerBuilderParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BasePowerBuilderParserListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BasePowerBuilderParserListener) ExitString_literal(ctx *String_literalContext) {}

// EnterIdentifier_array is called when production identifier_array is entered.
func (s *BasePowerBuilderParserListener) EnterIdentifier_array(ctx *Identifier_arrayContext) {}

// ExitIdentifier_array is called when production identifier_array is exited.
func (s *BasePowerBuilderParserListener) ExitIdentifier_array(ctx *Identifier_arrayContext) {}

// EnterOperator is called when production operator is entered.
func (s *BasePowerBuilderParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BasePowerBuilderParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterIdentifier_name_ex is called when production identifier_name_ex is entered.
func (s *BasePowerBuilderParserListener) EnterIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// ExitIdentifier_name_ex is called when production identifier_name_ex is exited.
func (s *BasePowerBuilderParserListener) ExitIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// EnterIdentifier_name is called when production identifier_name is entered.
func (s *BasePowerBuilderParserListener) EnterIdentifier_name(ctx *Identifier_nameContext) {}

// ExitIdentifier_name is called when production identifier_name is exited.
func (s *BasePowerBuilderParserListener) ExitIdentifier_name(ctx *Identifier_nameContext) {}

// EnterBind_param is called when production bind_param is entered.
func (s *BasePowerBuilderParserListener) EnterBind_param(ctx *Bind_paramContext) {}

// ExitBind_param is called when production bind_param is exited.
func (s *BasePowerBuilderParserListener) ExitBind_param(ctx *Bind_paramContext) {}

// EnterAtom_sub is called when production atom_sub is entered.
func (s *BasePowerBuilderParserListener) EnterAtom_sub(ctx *Atom_subContext) {}

// ExitAtom_sub is called when production atom_sub is exited.
func (s *BasePowerBuilderParserListener) ExitAtom_sub(ctx *Atom_subContext) {}

// EnterAtom_sub_call1 is called when production atom_sub_call1 is entered.
func (s *BasePowerBuilderParserListener) EnterAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// ExitAtom_sub_call1 is called when production atom_sub_call1 is exited.
func (s *BasePowerBuilderParserListener) ExitAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// EnterAtom_sub_member1 is called when production atom_sub_member1 is entered.
func (s *BasePowerBuilderParserListener) EnterAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// ExitAtom_sub_member1 is called when production atom_sub_member1 is exited.
func (s *BasePowerBuilderParserListener) ExitAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// EnterArray_access_atom is called when production array_access_atom is entered.
func (s *BasePowerBuilderParserListener) EnterArray_access_atom(ctx *Array_access_atomContext) {}

// ExitArray_access_atom is called when production array_access_atom is exited.
func (s *BasePowerBuilderParserListener) ExitArray_access_atom(ctx *Array_access_atomContext) {}

// EnterData_type_name is called when production data_type_name is entered.
func (s *BasePowerBuilderParserListener) EnterData_type_name(ctx *Data_type_nameContext) {}

// ExitData_type_name is called when production data_type_name is exited.
func (s *BasePowerBuilderParserListener) ExitData_type_name(ctx *Data_type_nameContext) {}

// EnterDataTypeSub is called when production dataTypeSub is entered.
func (s *BasePowerBuilderParserListener) EnterDataTypeSub(ctx *DataTypeSubContext) {}

// ExitDataTypeSub is called when production dataTypeSub is exited.
func (s *BasePowerBuilderParserListener) ExitDataTypeSub(ctx *DataTypeSubContext) {}
