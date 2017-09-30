// Generated from powerbuilder.g4 by ANTLR 4.7.

package powerbuilder // powerbuilder
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by powerbuilderParser.
type powerbuilderVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by powerbuilderParser#start_rule.
	VisitStart_rule(ctx *Start_ruleContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#header_rule.
	VisitHeader_rule(ctx *Header_ruleContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#body_rule.
	VisitBody_rule(ctx *Body_ruleContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#export_header.
	VisitExport_header(ctx *Export_headerContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#release_information.
	VisitRelease_information(ctx *Release_informationContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#window_property_line.
	VisitWindow_property_line(ctx *Window_property_lineContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#window_property.
	VisitWindow_property(ctx *Window_propertyContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#window_property_attributes_sub.
	VisitWindow_property_attributes_sub(ctx *Window_property_attributes_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#window_property_attribute_sub.
	VisitWindow_property_attribute_sub(ctx *Window_property_attribute_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#attribute_name.
	VisitAttribute_name(ctx *Attribute_nameContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#attribute_value.
	VisitAttribute_value(ctx *Attribute_valueContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#forward_decl.
	VisitForward_decl(ctx *Forward_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#datatype_decl.
	VisitDatatype_decl(ctx *Datatype_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#type_variables_decl.
	VisitType_variables_decl(ctx *Type_variables_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#global_variables_decl.
	VisitGlobal_variables_decl(ctx *Global_variables_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#variable_decl_sub.
	VisitVariable_decl_sub(ctx *Variable_decl_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#variable_decl.
	VisitVariable_decl(ctx *Variable_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#decimal_decl_sub.
	VisitDecimal_decl_sub(ctx *Decimal_decl_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#array_decl_sub.
	VisitArray_decl_sub(ctx *Array_decl_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#constant_decl_sub.
	VisitConstant_decl_sub(ctx *Constant_decl_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#constant_decl.
	VisitConstant_decl(ctx *Constant_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#function_forward_decl.
	VisitFunction_forward_decl(ctx *Function_forward_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#parameter_sub.
	VisitParameter_sub(ctx *Parameter_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#parameters_list_sub.
	VisitParameters_list_sub(ctx *Parameters_list_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#functions_forward_decl.
	VisitFunctions_forward_decl(ctx *Functions_forward_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#function_body.
	VisitFunction_body(ctx *Function_bodyContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#on_body.
	VisitOn_body(ctx *On_bodyContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#event_forward_decl_sub.
	VisitEvent_forward_decl_sub(ctx *Event_forward_decl_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#event_forward_decl.
	VisitEvent_forward_decl(ctx *Event_forward_declContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#event_body.
	VisitEvent_body(ctx *Event_bodyContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#access_modif.
	VisitAccess_modif(ctx *Access_modifContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#access_modif_part.
	VisitAccess_modif_part(ctx *Access_modif_partContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#scope_modif.
	VisitScope_modif(ctx *Scope_modifContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#boolean_expression.
	VisitBoolean_expression(ctx *Boolean_expressionContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#condition_or.
	VisitCondition_or(ctx *Condition_orContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#condition_and.
	VisitCondition_and(ctx *Condition_andContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#condition_not.
	VisitCondition_not(ctx *Condition_notContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#condition_comparison.
	VisitCondition_comparison(ctx *Condition_comparisonContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#add_expr.
	VisitAdd_expr(ctx *Add_exprContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#mul_expr.
	VisitMul_expr(ctx *Mul_exprContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#unary_sign_expr.
	VisitUnary_sign_expr(ctx *Unary_sign_exprContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#statement_sub.
	VisitStatement_sub(ctx *Statement_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#assignment_sub.
	VisitAssignment_sub(ctx *Assignment_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#lvalue_sub.
	VisitLvalue_sub(ctx *Lvalue_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#return_sub.
	VisitReturn_sub(ctx *Return_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#function_call_expression_sub.
	VisitFunction_call_expression_sub(ctx *Function_call_expression_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#function_virtual_call_expression_sub.
	VisitFunction_virtual_call_expression_sub(ctx *Function_virtual_call_expression_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#open_call_sub.
	VisitOpen_call_sub(ctx *Open_call_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#close_call_sub.
	VisitClose_call_sub(ctx *Close_call_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#function_call_statement.
	VisitFunction_call_statement(ctx *Function_call_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#super_call_sub.
	VisitSuper_call_sub(ctx *Super_call_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#super_call_statement.
	VisitSuper_call_statement(ctx *Super_call_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#event_call_statement_sub.
	VisitEvent_call_statement_sub(ctx *Event_call_statement_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#event_call_statement.
	VisitEvent_call_statement(ctx *Event_call_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#create_call_sub.
	VisitCreate_call_sub(ctx *Create_call_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#create_call_statement.
	VisitCreate_call_statement(ctx *Create_call_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#destroy_call_sub.
	VisitDestroy_call_sub(ctx *Destroy_call_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#destroy_call_statement.
	VisitDestroy_call_statement(ctx *Destroy_call_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#for_loop_statement.
	VisitFor_loop_statement(ctx *For_loop_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#do_while_loop_statement.
	VisitDo_while_loop_statement(ctx *Do_while_loop_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#do_loop_while_statement.
	VisitDo_loop_while_statement(ctx *Do_loop_while_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#if_simple_statement.
	VisitIf_simple_statement(ctx *If_simple_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#continue_sub.
	VisitContinue_sub(ctx *Continue_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#post_event_sub.
	VisitPost_event_sub(ctx *Post_event_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#post_event.
	VisitPost_event(ctx *Post_eventContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#exit_statement_sub.
	VisitExit_statement_sub(ctx *Exit_statement_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#exit_statement.
	VisitExit_statement(ctx *Exit_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#choose_statement.
	VisitChoose_statement(ctx *Choose_statementContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#choose_case_value_sub.
	VisitChoose_case_value_sub(ctx *Choose_case_value_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#choose_case_cond_sub.
	VisitChoose_case_cond_sub(ctx *Choose_case_cond_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#choose_case_range_sub.
	VisitChoose_case_range_sub(ctx *Choose_case_range_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#choose_case_else_sub.
	VisitChoose_case_else_sub(ctx *Choose_case_else_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#goto_stat_sub.
	VisitGoto_stat_sub(ctx *Goto_stat_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#goto_stat.
	VisitGoto_stat(ctx *Goto_statContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#label_stat.
	VisitLabel_stat(ctx *Label_statContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#try_catch_block.
	VisitTry_catch_block(ctx *Try_catch_blockContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#throw_stat_sub.
	VisitThrow_stat_sub(ctx *Throw_stat_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#throw_stat.
	VisitThrow_stat(ctx *Throw_statContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#identifier_name.
	VisitIdentifier_name(ctx *Identifier_nameContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#identifier_name_ex.
	VisitIdentifier_name_ex(ctx *Identifier_name_exContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom_sub.
	VisitAtom_sub(ctx *Atom_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom_sub_call1.
	VisitAtom_sub_call1(ctx *Atom_sub_call1Context) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom_sub_array1.
	VisitAtom_sub_array1(ctx *Atom_sub_array1Context) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom_sub_ref1.
	VisitAtom_sub_ref1(ctx *Atom_sub_ref1Context) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom_sub_member1.
	VisitAtom_sub_member1(ctx *Atom_sub_member1Context) interface{}

	// Visit a parse tree produced by powerbuilderParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#swallow_to_semi.
	VisitSwallow_to_semi(ctx *Swallow_to_semiContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#swallow_to_newline.
	VisitSwallow_to_newline(ctx *Swallow_to_newlineContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#array_access_atom.
	VisitArray_access_atom(ctx *Array_access_atomContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#numeric_atom.
	VisitNumeric_atom(ctx *Numeric_atomContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#boolean_atom.
	VisitBoolean_atom(ctx *Boolean_atomContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#cast_expression.
	VisitCast_expression(ctx *Cast_expressionContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#data_type_sub.
	VisitData_type_sub(ctx *Data_type_subContext) interface{}

	// Visit a parse tree produced by powerbuilderParser#data_type_name.
	VisitData_type_name(ctx *Data_type_nameContext) interface{}
}
