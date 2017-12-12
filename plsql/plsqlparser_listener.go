// Code generated from PlSqlParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package plsql // PlSqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PlSqlParserListener is a complete listener for a parse tree produced by PlSqlParser.
type PlSqlParserListener interface {
	antlr.ParseTreeListener

	// EnterSql_script is called when entering the sql_script production.
	EnterSql_script(c *Sql_scriptContext)

	// EnterUnit_statement is called when entering the unit_statement production.
	EnterUnit_statement(c *Unit_statementContext)

	// EnterDrop_function is called when entering the drop_function production.
	EnterDrop_function(c *Drop_functionContext)

	// EnterAlter_function is called when entering the alter_function production.
	EnterAlter_function(c *Alter_functionContext)

	// EnterCreate_function_body is called when entering the create_function_body production.
	EnterCreate_function_body(c *Create_function_bodyContext)

	// EnterParallel_enable_clause is called when entering the parallel_enable_clause production.
	EnterParallel_enable_clause(c *Parallel_enable_clauseContext)

	// EnterPartition_by_clause is called when entering the partition_by_clause production.
	EnterPartition_by_clause(c *Partition_by_clauseContext)

	// EnterResult_cache_clause is called when entering the result_cache_clause production.
	EnterResult_cache_clause(c *Result_cache_clauseContext)

	// EnterRelies_on_part is called when entering the relies_on_part production.
	EnterRelies_on_part(c *Relies_on_partContext)

	// EnterStreaming_clause is called when entering the streaming_clause production.
	EnterStreaming_clause(c *Streaming_clauseContext)

	// EnterDrop_package is called when entering the drop_package production.
	EnterDrop_package(c *Drop_packageContext)

	// EnterAlter_package is called when entering the alter_package production.
	EnterAlter_package(c *Alter_packageContext)

	// EnterCreate_package is called when entering the create_package production.
	EnterCreate_package(c *Create_packageContext)

	// EnterCreate_package_body is called when entering the create_package_body production.
	EnterCreate_package_body(c *Create_package_bodyContext)

	// EnterPackage_obj_spec is called when entering the package_obj_spec production.
	EnterPackage_obj_spec(c *Package_obj_specContext)

	// EnterProcedure_spec is called when entering the procedure_spec production.
	EnterProcedure_spec(c *Procedure_specContext)

	// EnterFunction_spec is called when entering the function_spec production.
	EnterFunction_spec(c *Function_specContext)

	// EnterPackage_obj_body is called when entering the package_obj_body production.
	EnterPackage_obj_body(c *Package_obj_bodyContext)

	// EnterDrop_procedure is called when entering the drop_procedure production.
	EnterDrop_procedure(c *Drop_procedureContext)

	// EnterAlter_procedure is called when entering the alter_procedure production.
	EnterAlter_procedure(c *Alter_procedureContext)

	// EnterFunction_body is called when entering the function_body production.
	EnterFunction_body(c *Function_bodyContext)

	// EnterProcedure_body is called when entering the procedure_body production.
	EnterProcedure_body(c *Procedure_bodyContext)

	// EnterCreate_procedure_body is called when entering the create_procedure_body production.
	EnterCreate_procedure_body(c *Create_procedure_bodyContext)

	// EnterDrop_trigger is called when entering the drop_trigger production.
	EnterDrop_trigger(c *Drop_triggerContext)

	// EnterAlter_trigger is called when entering the alter_trigger production.
	EnterAlter_trigger(c *Alter_triggerContext)

	// EnterCreate_trigger is called when entering the create_trigger production.
	EnterCreate_trigger(c *Create_triggerContext)

	// EnterTrigger_follows_clause is called when entering the trigger_follows_clause production.
	EnterTrigger_follows_clause(c *Trigger_follows_clauseContext)

	// EnterTrigger_when_clause is called when entering the trigger_when_clause production.
	EnterTrigger_when_clause(c *Trigger_when_clauseContext)

	// EnterSimple_dml_trigger is called when entering the simple_dml_trigger production.
	EnterSimple_dml_trigger(c *Simple_dml_triggerContext)

	// EnterFor_each_row is called when entering the for_each_row production.
	EnterFor_each_row(c *For_each_rowContext)

	// EnterCompound_dml_trigger is called when entering the compound_dml_trigger production.
	EnterCompound_dml_trigger(c *Compound_dml_triggerContext)

	// EnterNon_dml_trigger is called when entering the non_dml_trigger production.
	EnterNon_dml_trigger(c *Non_dml_triggerContext)

	// EnterTrigger_body is called when entering the trigger_body production.
	EnterTrigger_body(c *Trigger_bodyContext)

	// EnterRoutine_clause is called when entering the routine_clause production.
	EnterRoutine_clause(c *Routine_clauseContext)

	// EnterCompound_trigger_block is called when entering the compound_trigger_block production.
	EnterCompound_trigger_block(c *Compound_trigger_blockContext)

	// EnterTiming_point_section is called when entering the timing_point_section production.
	EnterTiming_point_section(c *Timing_point_sectionContext)

	// EnterNon_dml_event is called when entering the non_dml_event production.
	EnterNon_dml_event(c *Non_dml_eventContext)

	// EnterDml_event_clause is called when entering the dml_event_clause production.
	EnterDml_event_clause(c *Dml_event_clauseContext)

	// EnterDml_event_element is called when entering the dml_event_element production.
	EnterDml_event_element(c *Dml_event_elementContext)

	// EnterDml_event_nested_clause is called when entering the dml_event_nested_clause production.
	EnterDml_event_nested_clause(c *Dml_event_nested_clauseContext)

	// EnterReferencing_clause is called when entering the referencing_clause production.
	EnterReferencing_clause(c *Referencing_clauseContext)

	// EnterReferencing_element is called when entering the referencing_element production.
	EnterReferencing_element(c *Referencing_elementContext)

	// EnterDrop_type is called when entering the drop_type production.
	EnterDrop_type(c *Drop_typeContext)

	// EnterAlter_type is called when entering the alter_type production.
	EnterAlter_type(c *Alter_typeContext)

	// EnterCompile_type_clause is called when entering the compile_type_clause production.
	EnterCompile_type_clause(c *Compile_type_clauseContext)

	// EnterReplace_type_clause is called when entering the replace_type_clause production.
	EnterReplace_type_clause(c *Replace_type_clauseContext)

	// EnterAlter_method_spec is called when entering the alter_method_spec production.
	EnterAlter_method_spec(c *Alter_method_specContext)

	// EnterAlter_method_element is called when entering the alter_method_element production.
	EnterAlter_method_element(c *Alter_method_elementContext)

	// EnterAlter_attribute_definition is called when entering the alter_attribute_definition production.
	EnterAlter_attribute_definition(c *Alter_attribute_definitionContext)

	// EnterAttribute_definition is called when entering the attribute_definition production.
	EnterAttribute_definition(c *Attribute_definitionContext)

	// EnterAlter_collection_clauses is called when entering the alter_collection_clauses production.
	EnterAlter_collection_clauses(c *Alter_collection_clausesContext)

	// EnterDependent_handling_clause is called when entering the dependent_handling_clause production.
	EnterDependent_handling_clause(c *Dependent_handling_clauseContext)

	// EnterDependent_exceptions_part is called when entering the dependent_exceptions_part production.
	EnterDependent_exceptions_part(c *Dependent_exceptions_partContext)

	// EnterCreate_type is called when entering the create_type production.
	EnterCreate_type(c *Create_typeContext)

	// EnterType_definition is called when entering the type_definition production.
	EnterType_definition(c *Type_definitionContext)

	// EnterObject_type_def is called when entering the object_type_def production.
	EnterObject_type_def(c *Object_type_defContext)

	// EnterObject_as_part is called when entering the object_as_part production.
	EnterObject_as_part(c *Object_as_partContext)

	// EnterObject_under_part is called when entering the object_under_part production.
	EnterObject_under_part(c *Object_under_partContext)

	// EnterNested_table_type_def is called when entering the nested_table_type_def production.
	EnterNested_table_type_def(c *Nested_table_type_defContext)

	// EnterSqlj_object_type is called when entering the sqlj_object_type production.
	EnterSqlj_object_type(c *Sqlj_object_typeContext)

	// EnterType_body is called when entering the type_body production.
	EnterType_body(c *Type_bodyContext)

	// EnterType_body_elements is called when entering the type_body_elements production.
	EnterType_body_elements(c *Type_body_elementsContext)

	// EnterMap_order_func_declaration is called when entering the map_order_func_declaration production.
	EnterMap_order_func_declaration(c *Map_order_func_declarationContext)

	// EnterSubprog_decl_in_type is called when entering the subprog_decl_in_type production.
	EnterSubprog_decl_in_type(c *Subprog_decl_in_typeContext)

	// EnterProc_decl_in_type is called when entering the proc_decl_in_type production.
	EnterProc_decl_in_type(c *Proc_decl_in_typeContext)

	// EnterFunc_decl_in_type is called when entering the func_decl_in_type production.
	EnterFunc_decl_in_type(c *Func_decl_in_typeContext)

	// EnterConstructor_declaration is called when entering the constructor_declaration production.
	EnterConstructor_declaration(c *Constructor_declarationContext)

	// EnterModifier_clause is called when entering the modifier_clause production.
	EnterModifier_clause(c *Modifier_clauseContext)

	// EnterObject_member_spec is called when entering the object_member_spec production.
	EnterObject_member_spec(c *Object_member_specContext)

	// EnterSqlj_object_type_attr is called when entering the sqlj_object_type_attr production.
	EnterSqlj_object_type_attr(c *Sqlj_object_type_attrContext)

	// EnterElement_spec is called when entering the element_spec production.
	EnterElement_spec(c *Element_specContext)

	// EnterElement_spec_options is called when entering the element_spec_options production.
	EnterElement_spec_options(c *Element_spec_optionsContext)

	// EnterSubprogram_spec is called when entering the subprogram_spec production.
	EnterSubprogram_spec(c *Subprogram_specContext)

	// EnterType_procedure_spec is called when entering the type_procedure_spec production.
	EnterType_procedure_spec(c *Type_procedure_specContext)

	// EnterType_function_spec is called when entering the type_function_spec production.
	EnterType_function_spec(c *Type_function_specContext)

	// EnterConstructor_spec is called when entering the constructor_spec production.
	EnterConstructor_spec(c *Constructor_specContext)

	// EnterMap_order_function_spec is called when entering the map_order_function_spec production.
	EnterMap_order_function_spec(c *Map_order_function_specContext)

	// EnterPragma_clause is called when entering the pragma_clause production.
	EnterPragma_clause(c *Pragma_clauseContext)

	// EnterPragma_elements is called when entering the pragma_elements production.
	EnterPragma_elements(c *Pragma_elementsContext)

	// EnterType_elements_parameter is called when entering the type_elements_parameter production.
	EnterType_elements_parameter(c *Type_elements_parameterContext)

	// EnterDrop_sequence is called when entering the drop_sequence production.
	EnterDrop_sequence(c *Drop_sequenceContext)

	// EnterAlter_sequence is called when entering the alter_sequence production.
	EnterAlter_sequence(c *Alter_sequenceContext)

	// EnterCreate_sequence is called when entering the create_sequence production.
	EnterCreate_sequence(c *Create_sequenceContext)

	// EnterSequence_spec is called when entering the sequence_spec production.
	EnterSequence_spec(c *Sequence_specContext)

	// EnterSequence_start_clause is called when entering the sequence_start_clause production.
	EnterSequence_start_clause(c *Sequence_start_clauseContext)

	// EnterCreate_index is called when entering the create_index production.
	EnterCreate_index(c *Create_indexContext)

	// EnterAlter_index is called when entering the alter_index production.
	EnterAlter_index(c *Alter_indexContext)

	// EnterDrop_index is called when entering the drop_index production.
	EnterDrop_index(c *Drop_indexContext)

	// EnterCreate_table is called when entering the create_table production.
	EnterCreate_table(c *Create_tableContext)

	// EnterSize_clause is called when entering the size_clause production.
	EnterSize_clause(c *Size_clauseContext)

	// EnterDrop_table is called when entering the drop_table production.
	EnterDrop_table(c *Drop_tableContext)

	// EnterComment_on_column is called when entering the comment_on_column production.
	EnterComment_on_column(c *Comment_on_columnContext)

	// EnterCreate_synonym is called when entering the create_synonym production.
	EnterCreate_synonym(c *Create_synonymContext)

	// EnterComment_on_table is called when entering the comment_on_table production.
	EnterComment_on_table(c *Comment_on_tableContext)

	// EnterAlter_table is called when entering the alter_table production.
	EnterAlter_table(c *Alter_tableContext)

	// EnterAdd_constraint is called when entering the add_constraint production.
	EnterAdd_constraint(c *Add_constraintContext)

	// EnterCheck_constraint is called when entering the check_constraint production.
	EnterCheck_constraint(c *Check_constraintContext)

	// EnterDrop_constraint is called when entering the drop_constraint production.
	EnterDrop_constraint(c *Drop_constraintContext)

	// EnterEnable_constraint is called when entering the enable_constraint production.
	EnterEnable_constraint(c *Enable_constraintContext)

	// EnterDisable_constraint is called when entering the disable_constraint production.
	EnterDisable_constraint(c *Disable_constraintContext)

	// EnterForeign_key_clause is called when entering the foreign_key_clause production.
	EnterForeign_key_clause(c *Foreign_key_clauseContext)

	// EnterReferences_clause is called when entering the references_clause production.
	EnterReferences_clause(c *References_clauseContext)

	// EnterOn_delete_clause is called when entering the on_delete_clause production.
	EnterOn_delete_clause(c *On_delete_clauseContext)

	// EnterUnique_key_clause is called when entering the unique_key_clause production.
	EnterUnique_key_clause(c *Unique_key_clauseContext)

	// EnterPrimary_key_clause is called when entering the primary_key_clause production.
	EnterPrimary_key_clause(c *Primary_key_clauseContext)

	// EnterAnonymous_block is called when entering the anonymous_block production.
	EnterAnonymous_block(c *Anonymous_blockContext)

	// EnterInvoker_rights_clause is called when entering the invoker_rights_clause production.
	EnterInvoker_rights_clause(c *Invoker_rights_clauseContext)

	// EnterCompiler_parameters_clause is called when entering the compiler_parameters_clause production.
	EnterCompiler_parameters_clause(c *Compiler_parameters_clauseContext)

	// EnterCall_spec is called when entering the call_spec production.
	EnterCall_spec(c *Call_specContext)

	// EnterJava_spec is called when entering the java_spec production.
	EnterJava_spec(c *Java_specContext)

	// EnterC_spec is called when entering the c_spec production.
	EnterC_spec(c *C_specContext)

	// EnterC_agent_in_clause is called when entering the c_agent_in_clause production.
	EnterC_agent_in_clause(c *C_agent_in_clauseContext)

	// EnterC_parameters_clause is called when entering the c_parameters_clause production.
	EnterC_parameters_clause(c *C_parameters_clauseContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterDefault_value_part is called when entering the default_value_part production.
	EnterDefault_value_part(c *Default_value_partContext)

	// EnterDeclare_spec is called when entering the declare_spec production.
	EnterDeclare_spec(c *Declare_specContext)

	// EnterVariable_declaration is called when entering the variable_declaration production.
	EnterVariable_declaration(c *Variable_declarationContext)

	// EnterSubtype_declaration is called when entering the subtype_declaration production.
	EnterSubtype_declaration(c *Subtype_declarationContext)

	// EnterCursor_declaration is called when entering the cursor_declaration production.
	EnterCursor_declaration(c *Cursor_declarationContext)

	// EnterParameter_spec is called when entering the parameter_spec production.
	EnterParameter_spec(c *Parameter_specContext)

	// EnterException_declaration is called when entering the exception_declaration production.
	EnterException_declaration(c *Exception_declarationContext)

	// EnterPragma_declaration is called when entering the pragma_declaration production.
	EnterPragma_declaration(c *Pragma_declarationContext)

	// EnterRecord_type_def is called when entering the record_type_def production.
	EnterRecord_type_def(c *Record_type_defContext)

	// EnterField_spec is called when entering the field_spec production.
	EnterField_spec(c *Field_specContext)

	// EnterRef_cursor_type_def is called when entering the ref_cursor_type_def production.
	EnterRef_cursor_type_def(c *Ref_cursor_type_defContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterTable_type_def is called when entering the table_type_def production.
	EnterTable_type_def(c *Table_type_defContext)

	// EnterTable_indexed_by_part is called when entering the table_indexed_by_part production.
	EnterTable_indexed_by_part(c *Table_indexed_by_partContext)

	// EnterVarray_type_def is called when entering the varray_type_def production.
	EnterVarray_type_def(c *Varray_type_defContext)

	// EnterSeq_of_statements is called when entering the seq_of_statements production.
	EnterSeq_of_statements(c *Seq_of_statementsContext)

	// EnterLabel_declaration is called when entering the label_declaration production.
	EnterLabel_declaration(c *Label_declarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSwallow_to_semi is called when entering the swallow_to_semi production.
	EnterSwallow_to_semi(c *Swallow_to_semiContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterExit_statement is called when entering the exit_statement production.
	EnterExit_statement(c *Exit_statementContext)

	// EnterGoto_statement is called when entering the goto_statement production.
	EnterGoto_statement(c *Goto_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterElsif_part is called when entering the elsif_part production.
	EnterElsif_part(c *Elsif_partContext)

	// EnterElse_part is called when entering the else_part production.
	EnterElse_part(c *Else_partContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterCursor_loop_param is called when entering the cursor_loop_param production.
	EnterCursor_loop_param(c *Cursor_loop_paramContext)

	// EnterForall_statement is called when entering the forall_statement production.
	EnterForall_statement(c *Forall_statementContext)

	// EnterBounds_clause is called when entering the bounds_clause production.
	EnterBounds_clause(c *Bounds_clauseContext)

	// EnterBetween_bound is called when entering the between_bound production.
	EnterBetween_bound(c *Between_boundContext)

	// EnterLower_bound is called when entering the lower_bound production.
	EnterLower_bound(c *Lower_boundContext)

	// EnterUpper_bound is called when entering the upper_bound production.
	EnterUpper_bound(c *Upper_boundContext)

	// EnterNull_statement is called when entering the null_statement production.
	EnterNull_statement(c *Null_statementContext)

	// EnterRaise_statement is called when entering the raise_statement production.
	EnterRaise_statement(c *Raise_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterPipe_row_statement is called when entering the pipe_row_statement production.
	EnterPipe_row_statement(c *Pipe_row_statementContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterException_handler is called when entering the exception_handler production.
	EnterException_handler(c *Exception_handlerContext)

	// EnterTrigger_block is called when entering the trigger_block production.
	EnterTrigger_block(c *Trigger_blockContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSql_statement is called when entering the sql_statement production.
	EnterSql_statement(c *Sql_statementContext)

	// EnterExecute_immediate is called when entering the execute_immediate production.
	EnterExecute_immediate(c *Execute_immediateContext)

	// EnterDynamic_returning_clause is called when entering the dynamic_returning_clause production.
	EnterDynamic_returning_clause(c *Dynamic_returning_clauseContext)

	// EnterData_manipulation_language_statements is called when entering the data_manipulation_language_statements production.
	EnterData_manipulation_language_statements(c *Data_manipulation_language_statementsContext)

	// EnterCursor_manipulation_statements is called when entering the cursor_manipulation_statements production.
	EnterCursor_manipulation_statements(c *Cursor_manipulation_statementsContext)

	// EnterClose_statement is called when entering the close_statement production.
	EnterClose_statement(c *Close_statementContext)

	// EnterOpen_statement is called when entering the open_statement production.
	EnterOpen_statement(c *Open_statementContext)

	// EnterFetch_statement is called when entering the fetch_statement production.
	EnterFetch_statement(c *Fetch_statementContext)

	// EnterOpen_for_statement is called when entering the open_for_statement production.
	EnterOpen_for_statement(c *Open_for_statementContext)

	// EnterTransaction_control_statements is called when entering the transaction_control_statements production.
	EnterTransaction_control_statements(c *Transaction_control_statementsContext)

	// EnterSet_transaction_command is called when entering the set_transaction_command production.
	EnterSet_transaction_command(c *Set_transaction_commandContext)

	// EnterSet_constraint_command is called when entering the set_constraint_command production.
	EnterSet_constraint_command(c *Set_constraint_commandContext)

	// EnterCommit_statement is called when entering the commit_statement production.
	EnterCommit_statement(c *Commit_statementContext)

	// EnterWrite_clause is called when entering the write_clause production.
	EnterWrite_clause(c *Write_clauseContext)

	// EnterRollback_statement is called when entering the rollback_statement production.
	EnterRollback_statement(c *Rollback_statementContext)

	// EnterSavepoint_statement is called when entering the savepoint_statement production.
	EnterSavepoint_statement(c *Savepoint_statementContext)

	// EnterExplain_statement is called when entering the explain_statement production.
	EnterExplain_statement(c *Explain_statementContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterSubquery_factoring_clause is called when entering the subquery_factoring_clause production.
	EnterSubquery_factoring_clause(c *Subquery_factoring_clauseContext)

	// EnterFactoring_element is called when entering the factoring_element production.
	EnterFactoring_element(c *Factoring_elementContext)

	// EnterSearch_clause is called when entering the search_clause production.
	EnterSearch_clause(c *Search_clauseContext)

	// EnterCycle_clause is called when entering the cycle_clause production.
	EnterCycle_clause(c *Cycle_clauseContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterSubquery_basic_elements is called when entering the subquery_basic_elements production.
	EnterSubquery_basic_elements(c *Subquery_basic_elementsContext)

	// EnterSubquery_operation_part is called when entering the subquery_operation_part production.
	EnterSubquery_operation_part(c *Subquery_operation_partContext)

	// EnterQuery_block is called when entering the query_block production.
	EnterQuery_block(c *Query_blockContext)

	// EnterSelected_element is called when entering the selected_element production.
	EnterSelected_element(c *Selected_elementContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterSelect_list_elements is called when entering the select_list_elements production.
	EnterSelect_list_elements(c *Select_list_elementsContext)

	// EnterTable_ref_list is called when entering the table_ref_list production.
	EnterTable_ref_list(c *Table_ref_listContext)

	// EnterTable_ref is called when entering the table_ref production.
	EnterTable_ref(c *Table_refContext)

	// EnterTable_ref_aux is called when entering the table_ref_aux production.
	EnterTable_ref_aux(c *Table_ref_auxContext)

	// EnterTable_ref_aux_internal_one is called when entering the table_ref_aux_internal_one production.
	EnterTable_ref_aux_internal_one(c *Table_ref_aux_internal_oneContext)

	// EnterTable_ref_aux_internal_two is called when entering the table_ref_aux_internal_two production.
	EnterTable_ref_aux_internal_two(c *Table_ref_aux_internal_twoContext)

	// EnterTable_ref_aux_internal_three is called when entering the table_ref_aux_internal_three production.
	EnterTable_ref_aux_internal_three(c *Table_ref_aux_internal_threeContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterJoin_on_part is called when entering the join_on_part production.
	EnterJoin_on_part(c *Join_on_partContext)

	// EnterJoin_using_part is called when entering the join_using_part production.
	EnterJoin_using_part(c *Join_using_partContext)

	// EnterOuter_join_type is called when entering the outer_join_type production.
	EnterOuter_join_type(c *Outer_join_typeContext)

	// EnterQuery_partition_clause is called when entering the query_partition_clause production.
	EnterQuery_partition_clause(c *Query_partition_clauseContext)

	// EnterFlashback_query_clause is called when entering the flashback_query_clause production.
	EnterFlashback_query_clause(c *Flashback_query_clauseContext)

	// EnterPivot_clause is called when entering the pivot_clause production.
	EnterPivot_clause(c *Pivot_clauseContext)

	// EnterPivot_element is called when entering the pivot_element production.
	EnterPivot_element(c *Pivot_elementContext)

	// EnterPivot_for_clause is called when entering the pivot_for_clause production.
	EnterPivot_for_clause(c *Pivot_for_clauseContext)

	// EnterPivot_in_clause is called when entering the pivot_in_clause production.
	EnterPivot_in_clause(c *Pivot_in_clauseContext)

	// EnterPivot_in_clause_element is called when entering the pivot_in_clause_element production.
	EnterPivot_in_clause_element(c *Pivot_in_clause_elementContext)

	// EnterPivot_in_clause_elements is called when entering the pivot_in_clause_elements production.
	EnterPivot_in_clause_elements(c *Pivot_in_clause_elementsContext)

	// EnterUnpivot_clause is called when entering the unpivot_clause production.
	EnterUnpivot_clause(c *Unpivot_clauseContext)

	// EnterUnpivot_in_clause is called when entering the unpivot_in_clause production.
	EnterUnpivot_in_clause(c *Unpivot_in_clauseContext)

	// EnterUnpivot_in_elements is called when entering the unpivot_in_elements production.
	EnterUnpivot_in_elements(c *Unpivot_in_elementsContext)

	// EnterHierarchical_query_clause is called when entering the hierarchical_query_clause production.
	EnterHierarchical_query_clause(c *Hierarchical_query_clauseContext)

	// EnterStart_part is called when entering the start_part production.
	EnterStart_part(c *Start_partContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterGroup_by_elements is called when entering the group_by_elements production.
	EnterGroup_by_elements(c *Group_by_elementsContext)

	// EnterRollup_cube_clause is called when entering the rollup_cube_clause production.
	EnterRollup_cube_clause(c *Rollup_cube_clauseContext)

	// EnterGrouping_sets_clause is called when entering the grouping_sets_clause production.
	EnterGrouping_sets_clause(c *Grouping_sets_clauseContext)

	// EnterGrouping_sets_elements is called when entering the grouping_sets_elements production.
	EnterGrouping_sets_elements(c *Grouping_sets_elementsContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterModel_clause is called when entering the model_clause production.
	EnterModel_clause(c *Model_clauseContext)

	// EnterCell_reference_options is called when entering the cell_reference_options production.
	EnterCell_reference_options(c *Cell_reference_optionsContext)

	// EnterReturn_rows_clause is called when entering the return_rows_clause production.
	EnterReturn_rows_clause(c *Return_rows_clauseContext)

	// EnterReference_model is called when entering the reference_model production.
	EnterReference_model(c *Reference_modelContext)

	// EnterMain_model is called when entering the main_model production.
	EnterMain_model(c *Main_modelContext)

	// EnterModel_column_clauses is called when entering the model_column_clauses production.
	EnterModel_column_clauses(c *Model_column_clausesContext)

	// EnterModel_column_partition_part is called when entering the model_column_partition_part production.
	EnterModel_column_partition_part(c *Model_column_partition_partContext)

	// EnterModel_column_list is called when entering the model_column_list production.
	EnterModel_column_list(c *Model_column_listContext)

	// EnterModel_column is called when entering the model_column production.
	EnterModel_column(c *Model_columnContext)

	// EnterModel_rules_clause is called when entering the model_rules_clause production.
	EnterModel_rules_clause(c *Model_rules_clauseContext)

	// EnterModel_rules_part is called when entering the model_rules_part production.
	EnterModel_rules_part(c *Model_rules_partContext)

	// EnterModel_rules_element is called when entering the model_rules_element production.
	EnterModel_rules_element(c *Model_rules_elementContext)

	// EnterCell_assignment is called when entering the cell_assignment production.
	EnterCell_assignment(c *Cell_assignmentContext)

	// EnterModel_iterate_clause is called when entering the model_iterate_clause production.
	EnterModel_iterate_clause(c *Model_iterate_clauseContext)

	// EnterUntil_part is called when entering the until_part production.
	EnterUntil_part(c *Until_partContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterOrder_by_elements is called when entering the order_by_elements production.
	EnterOrder_by_elements(c *Order_by_elementsContext)

	// EnterFor_update_clause is called when entering the for_update_clause production.
	EnterFor_update_clause(c *For_update_clauseContext)

	// EnterFor_update_of_part is called when entering the for_update_of_part production.
	EnterFor_update_of_part(c *For_update_of_partContext)

	// EnterFor_update_options is called when entering the for_update_options production.
	EnterFor_update_options(c *For_update_optionsContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterUpdate_set_clause is called when entering the update_set_clause production.
	EnterUpdate_set_clause(c *Update_set_clauseContext)

	// EnterColumn_based_update_set_clause is called when entering the column_based_update_set_clause production.
	EnterColumn_based_update_set_clause(c *Column_based_update_set_clauseContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterInsert_statement is called when entering the insert_statement production.
	EnterInsert_statement(c *Insert_statementContext)

	// EnterSingle_table_insert is called when entering the single_table_insert production.
	EnterSingle_table_insert(c *Single_table_insertContext)

	// EnterMulti_table_insert is called when entering the multi_table_insert production.
	EnterMulti_table_insert(c *Multi_table_insertContext)

	// EnterMulti_table_element is called when entering the multi_table_element production.
	EnterMulti_table_element(c *Multi_table_elementContext)

	// EnterConditional_insert_clause is called when entering the conditional_insert_clause production.
	EnterConditional_insert_clause(c *Conditional_insert_clauseContext)

	// EnterConditional_insert_when_part is called when entering the conditional_insert_when_part production.
	EnterConditional_insert_when_part(c *Conditional_insert_when_partContext)

	// EnterConditional_insert_else_part is called when entering the conditional_insert_else_part production.
	EnterConditional_insert_else_part(c *Conditional_insert_else_partContext)

	// EnterInsert_into_clause is called when entering the insert_into_clause production.
	EnterInsert_into_clause(c *Insert_into_clauseContext)

	// EnterValues_clause is called when entering the values_clause production.
	EnterValues_clause(c *Values_clauseContext)

	// EnterMerge_statement is called when entering the merge_statement production.
	EnterMerge_statement(c *Merge_statementContext)

	// EnterMerge_update_clause is called when entering the merge_update_clause production.
	EnterMerge_update_clause(c *Merge_update_clauseContext)

	// EnterMerge_element is called when entering the merge_element production.
	EnterMerge_element(c *Merge_elementContext)

	// EnterMerge_update_delete_part is called when entering the merge_update_delete_part production.
	EnterMerge_update_delete_part(c *Merge_update_delete_partContext)

	// EnterMerge_insert_clause is called when entering the merge_insert_clause production.
	EnterMerge_insert_clause(c *Merge_insert_clauseContext)

	// EnterSelected_tableview is called when entering the selected_tableview production.
	EnterSelected_tableview(c *Selected_tableviewContext)

	// EnterLock_table_statement is called when entering the lock_table_statement production.
	EnterLock_table_statement(c *Lock_table_statementContext)

	// EnterWait_nowait_part is called when entering the wait_nowait_part production.
	EnterWait_nowait_part(c *Wait_nowait_partContext)

	// EnterLock_table_element is called when entering the lock_table_element production.
	EnterLock_table_element(c *Lock_table_elementContext)

	// EnterLock_mode is called when entering the lock_mode production.
	EnterLock_mode(c *Lock_modeContext)

	// EnterGeneral_table_ref is called when entering the general_table_ref production.
	EnterGeneral_table_ref(c *General_table_refContext)

	// EnterStatic_returning_clause is called when entering the static_returning_clause production.
	EnterStatic_returning_clause(c *Static_returning_clauseContext)

	// EnterError_logging_clause is called when entering the error_logging_clause production.
	EnterError_logging_clause(c *Error_logging_clauseContext)

	// EnterError_logging_into_part is called when entering the error_logging_into_part production.
	EnterError_logging_into_part(c *Error_logging_into_partContext)

	// EnterError_logging_reject_part is called when entering the error_logging_reject_part production.
	EnterError_logging_reject_part(c *Error_logging_reject_partContext)

	// EnterDml_table_expression_clause is called when entering the dml_table_expression_clause production.
	EnterDml_table_expression_clause(c *Dml_table_expression_clauseContext)

	// EnterTable_collection_expression is called when entering the table_collection_expression production.
	EnterTable_collection_expression(c *Table_collection_expressionContext)

	// EnterSubquery_restriction_clause is called when entering the subquery_restriction_clause production.
	EnterSubquery_restriction_clause(c *Subquery_restriction_clauseContext)

	// EnterSample_clause is called when entering the sample_clause production.
	EnterSample_clause(c *Sample_clauseContext)

	// EnterSeed_part is called when entering the seed_part production.
	EnterSeed_part(c *Seed_partContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCursor_expression is called when entering the cursor_expression production.
	EnterCursor_expression(c *Cursor_expressionContext)

	// EnterLogical_expression is called when entering the logical_expression production.
	EnterLogical_expression(c *Logical_expressionContext)

	// EnterMultiset_expression is called when entering the multiset_expression production.
	EnterMultiset_expression(c *Multiset_expressionContext)

	// EnterRelational_expression is called when entering the relational_expression production.
	EnterRelational_expression(c *Relational_expressionContext)

	// EnterCompound_expression is called when entering the compound_expression production.
	EnterCompound_expression(c *Compound_expressionContext)

	// EnterRelational_operator is called when entering the relational_operator production.
	EnterRelational_operator(c *Relational_operatorContext)

	// EnterIn_elements is called when entering the in_elements production.
	EnterIn_elements(c *In_elementsContext)

	// EnterBetween_elements is called when entering the between_elements production.
	EnterBetween_elements(c *Between_elementsContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterInterval_expression is called when entering the interval_expression production.
	EnterInterval_expression(c *Interval_expressionContext)

	// EnterModel_expression is called when entering the model_expression production.
	EnterModel_expression(c *Model_expressionContext)

	// EnterModel_expression_element is called when entering the model_expression_element production.
	EnterModel_expression_element(c *Model_expression_elementContext)

	// EnterSingle_column_for_loop is called when entering the single_column_for_loop production.
	EnterSingle_column_for_loop(c *Single_column_for_loopContext)

	// EnterMulti_column_for_loop is called when entering the multi_column_for_loop production.
	EnterMulti_column_for_loop(c *Multi_column_for_loopContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterSimple_case_statement is called when entering the simple_case_statement production.
	EnterSimple_case_statement(c *Simple_case_statementContext)

	// EnterSimple_case_when_part is called when entering the simple_case_when_part production.
	EnterSimple_case_when_part(c *Simple_case_when_partContext)

	// EnterSearched_case_statement is called when entering the searched_case_statement production.
	EnterSearched_case_statement(c *Searched_case_statementContext)

	// EnterSearched_case_when_part is called when entering the searched_case_when_part production.
	EnterSearched_case_when_part(c *Searched_case_when_partContext)

	// EnterCase_else_part is called when entering the case_else_part production.
	EnterCase_else_part(c *Case_else_partContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterQuantified_expression is called when entering the quantified_expression production.
	EnterQuantified_expression(c *Quantified_expressionContext)

	// EnterString_function is called when entering the string_function production.
	EnterString_function(c *String_functionContext)

	// EnterStandard_function is called when entering the standard_function production.
	EnterStandard_function(c *Standard_functionContext)

	// EnterNumeric_function_wrapper is called when entering the numeric_function_wrapper production.
	EnterNumeric_function_wrapper(c *Numeric_function_wrapperContext)

	// EnterNumeric_function is called when entering the numeric_function production.
	EnterNumeric_function(c *Numeric_functionContext)

	// EnterOther_function is called when entering the other_function production.
	EnterOther_function(c *Other_functionContext)

	// EnterOver_clause_keyword is called when entering the over_clause_keyword production.
	EnterOver_clause_keyword(c *Over_clause_keywordContext)

	// EnterWithin_or_over_clause_keyword is called when entering the within_or_over_clause_keyword production.
	EnterWithin_or_over_clause_keyword(c *Within_or_over_clause_keywordContext)

	// EnterStandard_prediction_function_keyword is called when entering the standard_prediction_function_keyword production.
	EnterStandard_prediction_function_keyword(c *Standard_prediction_function_keywordContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterWindowing_clause is called when entering the windowing_clause production.
	EnterWindowing_clause(c *Windowing_clauseContext)

	// EnterWindowing_type is called when entering the windowing_type production.
	EnterWindowing_type(c *Windowing_typeContext)

	// EnterWindowing_elements is called when entering the windowing_elements production.
	EnterWindowing_elements(c *Windowing_elementsContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterUsing_element is called when entering the using_element production.
	EnterUsing_element(c *Using_elementContext)

	// EnterCollect_order_by_part is called when entering the collect_order_by_part production.
	EnterCollect_order_by_part(c *Collect_order_by_partContext)

	// EnterWithin_or_over_part is called when entering the within_or_over_part production.
	EnterWithin_or_over_part(c *Within_or_over_partContext)

	// EnterCost_matrix_clause is called when entering the cost_matrix_clause production.
	EnterCost_matrix_clause(c *Cost_matrix_clauseContext)

	// EnterXml_passing_clause is called when entering the xml_passing_clause production.
	EnterXml_passing_clause(c *Xml_passing_clauseContext)

	// EnterXml_attributes_clause is called when entering the xml_attributes_clause production.
	EnterXml_attributes_clause(c *Xml_attributes_clauseContext)

	// EnterXml_namespaces_clause is called when entering the xml_namespaces_clause production.
	EnterXml_namespaces_clause(c *Xml_namespaces_clauseContext)

	// EnterXml_table_column is called when entering the xml_table_column production.
	EnterXml_table_column(c *Xml_table_columnContext)

	// EnterXml_general_default_part is called when entering the xml_general_default_part production.
	EnterXml_general_default_part(c *Xml_general_default_partContext)

	// EnterXml_multiuse_expression_element is called when entering the xml_multiuse_expression_element production.
	EnterXml_multiuse_expression_element(c *Xml_multiuse_expression_elementContext)

	// EnterXmlroot_param_version_part is called when entering the xmlroot_param_version_part production.
	EnterXmlroot_param_version_part(c *Xmlroot_param_version_partContext)

	// EnterXmlroot_param_standalone_part is called when entering the xmlroot_param_standalone_part production.
	EnterXmlroot_param_standalone_part(c *Xmlroot_param_standalone_partContext)

	// EnterXmlserialize_param_enconding_part is called when entering the xmlserialize_param_enconding_part production.
	EnterXmlserialize_param_enconding_part(c *Xmlserialize_param_enconding_partContext)

	// EnterXmlserialize_param_version_part is called when entering the xmlserialize_param_version_part production.
	EnterXmlserialize_param_version_part(c *Xmlserialize_param_version_partContext)

	// EnterXmlserialize_param_ident_part is called when entering the xmlserialize_param_ident_part production.
	EnterXmlserialize_param_ident_part(c *Xmlserialize_param_ident_partContext)

	// EnterSql_plus_command is called when entering the sql_plus_command production.
	EnterSql_plus_command(c *Sql_plus_commandContext)

	// EnterWhenever_command is called when entering the whenever_command production.
	EnterWhenever_command(c *Whenever_commandContext)

	// EnterSet_command is called when entering the set_command production.
	EnterSet_command(c *Set_commandContext)

	// EnterPartition_extension_clause is called when entering the partition_extension_clause production.
	EnterPartition_extension_clause(c *Partition_extension_clauseContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterInto_clause is called when entering the into_clause production.
	EnterInto_clause(c *Into_clauseContext)

	// EnterXml_column_name is called when entering the xml_column_name production.
	EnterXml_column_name(c *Xml_column_nameContext)

	// EnterCost_class_name is called when entering the cost_class_name production.
	EnterCost_class_name(c *Cost_class_nameContext)

	// EnterAttribute_name is called when entering the attribute_name production.
	EnterAttribute_name(c *Attribute_nameContext)

	// EnterSavepoint_name is called when entering the savepoint_name production.
	EnterSavepoint_name(c *Savepoint_nameContext)

	// EnterRollback_segment_name is called when entering the rollback_segment_name production.
	EnterRollback_segment_name(c *Rollback_segment_nameContext)

	// EnterTable_var_name is called when entering the table_var_name production.
	EnterTable_var_name(c *Table_var_nameContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterRoutine_name is called when entering the routine_name production.
	EnterRoutine_name(c *Routine_nameContext)

	// EnterPackage_name is called when entering the package_name production.
	EnterPackage_name(c *Package_nameContext)

	// EnterImplementation_type_name is called when entering the implementation_type_name production.
	EnterImplementation_type_name(c *Implementation_type_nameContext)

	// EnterParameter_name is called when entering the parameter_name production.
	EnterParameter_name(c *Parameter_nameContext)

	// EnterReference_model_name is called when entering the reference_model_name production.
	EnterReference_model_name(c *Reference_model_nameContext)

	// EnterMain_model_name is called when entering the main_model_name production.
	EnterMain_model_name(c *Main_model_nameContext)

	// EnterAggregate_function_name is called when entering the aggregate_function_name production.
	EnterAggregate_function_name(c *Aggregate_function_nameContext)

	// EnterQuery_name is called when entering the query_name production.
	EnterQuery_name(c *Query_nameContext)

	// EnterConstraint_name is called when entering the constraint_name production.
	EnterConstraint_name(c *Constraint_nameContext)

	// EnterLabel_name is called when entering the label_name production.
	EnterLabel_name(c *Label_nameContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterSequence_name is called when entering the sequence_name production.
	EnterSequence_name(c *Sequence_nameContext)

	// EnterException_name is called when entering the exception_name production.
	EnterException_name(c *Exception_nameContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterProcedure_name is called when entering the procedure_name production.
	EnterProcedure_name(c *Procedure_nameContext)

	// EnterTrigger_name is called when entering the trigger_name production.
	EnterTrigger_name(c *Trigger_nameContext)

	// EnterVariable_name is called when entering the variable_name production.
	EnterVariable_name(c *Variable_nameContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterCursor_name is called when entering the cursor_name production.
	EnterCursor_name(c *Cursor_nameContext)

	// EnterRecord_name is called when entering the record_name production.
	EnterRecord_name(c *Record_nameContext)

	// EnterCollection_name is called when entering the collection_name production.
	EnterCollection_name(c *Collection_nameContext)

	// EnterLink_name is called when entering the link_name production.
	EnterLink_name(c *Link_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterTableview_name is called when entering the tableview_name production.
	EnterTableview_name(c *Tableview_nameContext)

	// EnterChar_set_name is called when entering the char_set_name production.
	EnterChar_set_name(c *Char_set_nameContext)

	// EnterSynonym_name is called when entering the synonym_name production.
	EnterSynonym_name(c *Synonym_nameContext)

	// EnterSchema_object_name is called when entering the schema_object_name production.
	EnterSchema_object_name(c *Schema_object_nameContext)

	// EnterKeep_clause is called when entering the keep_clause production.
	EnterKeep_clause(c *Keep_clauseContext)

	// EnterFunction_argument is called when entering the function_argument production.
	EnterFunction_argument(c *Function_argumentContext)

	// EnterFunction_argument_analytic is called when entering the function_argument_analytic production.
	EnterFunction_argument_analytic(c *Function_argument_analyticContext)

	// EnterFunction_argument_modeling is called when entering the function_argument_modeling production.
	EnterFunction_argument_modeling(c *Function_argument_modelingContext)

	// EnterRespect_or_ignore_nulls is called when entering the respect_or_ignore_nulls production.
	EnterRespect_or_ignore_nulls(c *Respect_or_ignore_nullsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterType_spec is called when entering the type_spec production.
	EnterType_spec(c *Type_specContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterPrecision_part is called when entering the precision_part production.
	EnterPrecision_part(c *Precision_partContext)

	// EnterNative_datatype_element is called when entering the native_datatype_element production.
	EnterNative_datatype_element(c *Native_datatype_elementContext)

	// EnterBind_variable is called when entering the bind_variable production.
	EnterBind_variable(c *Bind_variableContext)

	// EnterGeneral_element is called when entering the general_element production.
	EnterGeneral_element(c *General_elementContext)

	// EnterGeneral_element_part is called when entering the general_element_part production.
	EnterGeneral_element_part(c *General_element_partContext)

	// EnterTable_element is called when entering the table_element production.
	EnterTable_element(c *Table_elementContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumeric is called when entering the numeric production.
	EnterNumeric(c *NumericContext)

	// EnterNumeric_negative is called when entering the numeric_negative production.
	EnterNumeric_negative(c *Numeric_negativeContext)

	// EnterQuoted_string is called when entering the quoted_string production.
	EnterQuoted_string(c *Quoted_stringContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterId_expression is called when entering the id_expression production.
	EnterId_expression(c *Id_expressionContext)

	// EnterOuter_join_sign is called when entering the outer_join_sign production.
	EnterOuter_join_sign(c *Outer_join_signContext)

	// EnterRegular_id is called when entering the regular_id production.
	EnterRegular_id(c *Regular_idContext)

	// EnterString_function_name is called when entering the string_function_name production.
	EnterString_function_name(c *String_function_nameContext)

	// EnterNumeric_function_name is called when entering the numeric_function_name production.
	EnterNumeric_function_name(c *Numeric_function_nameContext)

	// ExitSql_script is called when exiting the sql_script production.
	ExitSql_script(c *Sql_scriptContext)

	// ExitUnit_statement is called when exiting the unit_statement production.
	ExitUnit_statement(c *Unit_statementContext)

	// ExitDrop_function is called when exiting the drop_function production.
	ExitDrop_function(c *Drop_functionContext)

	// ExitAlter_function is called when exiting the alter_function production.
	ExitAlter_function(c *Alter_functionContext)

	// ExitCreate_function_body is called when exiting the create_function_body production.
	ExitCreate_function_body(c *Create_function_bodyContext)

	// ExitParallel_enable_clause is called when exiting the parallel_enable_clause production.
	ExitParallel_enable_clause(c *Parallel_enable_clauseContext)

	// ExitPartition_by_clause is called when exiting the partition_by_clause production.
	ExitPartition_by_clause(c *Partition_by_clauseContext)

	// ExitResult_cache_clause is called when exiting the result_cache_clause production.
	ExitResult_cache_clause(c *Result_cache_clauseContext)

	// ExitRelies_on_part is called when exiting the relies_on_part production.
	ExitRelies_on_part(c *Relies_on_partContext)

	// ExitStreaming_clause is called when exiting the streaming_clause production.
	ExitStreaming_clause(c *Streaming_clauseContext)

	// ExitDrop_package is called when exiting the drop_package production.
	ExitDrop_package(c *Drop_packageContext)

	// ExitAlter_package is called when exiting the alter_package production.
	ExitAlter_package(c *Alter_packageContext)

	// ExitCreate_package is called when exiting the create_package production.
	ExitCreate_package(c *Create_packageContext)

	// ExitCreate_package_body is called when exiting the create_package_body production.
	ExitCreate_package_body(c *Create_package_bodyContext)

	// ExitPackage_obj_spec is called when exiting the package_obj_spec production.
	ExitPackage_obj_spec(c *Package_obj_specContext)

	// ExitProcedure_spec is called when exiting the procedure_spec production.
	ExitProcedure_spec(c *Procedure_specContext)

	// ExitFunction_spec is called when exiting the function_spec production.
	ExitFunction_spec(c *Function_specContext)

	// ExitPackage_obj_body is called when exiting the package_obj_body production.
	ExitPackage_obj_body(c *Package_obj_bodyContext)

	// ExitDrop_procedure is called when exiting the drop_procedure production.
	ExitDrop_procedure(c *Drop_procedureContext)

	// ExitAlter_procedure is called when exiting the alter_procedure production.
	ExitAlter_procedure(c *Alter_procedureContext)

	// ExitFunction_body is called when exiting the function_body production.
	ExitFunction_body(c *Function_bodyContext)

	// ExitProcedure_body is called when exiting the procedure_body production.
	ExitProcedure_body(c *Procedure_bodyContext)

	// ExitCreate_procedure_body is called when exiting the create_procedure_body production.
	ExitCreate_procedure_body(c *Create_procedure_bodyContext)

	// ExitDrop_trigger is called when exiting the drop_trigger production.
	ExitDrop_trigger(c *Drop_triggerContext)

	// ExitAlter_trigger is called when exiting the alter_trigger production.
	ExitAlter_trigger(c *Alter_triggerContext)

	// ExitCreate_trigger is called when exiting the create_trigger production.
	ExitCreate_trigger(c *Create_triggerContext)

	// ExitTrigger_follows_clause is called when exiting the trigger_follows_clause production.
	ExitTrigger_follows_clause(c *Trigger_follows_clauseContext)

	// ExitTrigger_when_clause is called when exiting the trigger_when_clause production.
	ExitTrigger_when_clause(c *Trigger_when_clauseContext)

	// ExitSimple_dml_trigger is called when exiting the simple_dml_trigger production.
	ExitSimple_dml_trigger(c *Simple_dml_triggerContext)

	// ExitFor_each_row is called when exiting the for_each_row production.
	ExitFor_each_row(c *For_each_rowContext)

	// ExitCompound_dml_trigger is called when exiting the compound_dml_trigger production.
	ExitCompound_dml_trigger(c *Compound_dml_triggerContext)

	// ExitNon_dml_trigger is called when exiting the non_dml_trigger production.
	ExitNon_dml_trigger(c *Non_dml_triggerContext)

	// ExitTrigger_body is called when exiting the trigger_body production.
	ExitTrigger_body(c *Trigger_bodyContext)

	// ExitRoutine_clause is called when exiting the routine_clause production.
	ExitRoutine_clause(c *Routine_clauseContext)

	// ExitCompound_trigger_block is called when exiting the compound_trigger_block production.
	ExitCompound_trigger_block(c *Compound_trigger_blockContext)

	// ExitTiming_point_section is called when exiting the timing_point_section production.
	ExitTiming_point_section(c *Timing_point_sectionContext)

	// ExitNon_dml_event is called when exiting the non_dml_event production.
	ExitNon_dml_event(c *Non_dml_eventContext)

	// ExitDml_event_clause is called when exiting the dml_event_clause production.
	ExitDml_event_clause(c *Dml_event_clauseContext)

	// ExitDml_event_element is called when exiting the dml_event_element production.
	ExitDml_event_element(c *Dml_event_elementContext)

	// ExitDml_event_nested_clause is called when exiting the dml_event_nested_clause production.
	ExitDml_event_nested_clause(c *Dml_event_nested_clauseContext)

	// ExitReferencing_clause is called when exiting the referencing_clause production.
	ExitReferencing_clause(c *Referencing_clauseContext)

	// ExitReferencing_element is called when exiting the referencing_element production.
	ExitReferencing_element(c *Referencing_elementContext)

	// ExitDrop_type is called when exiting the drop_type production.
	ExitDrop_type(c *Drop_typeContext)

	// ExitAlter_type is called when exiting the alter_type production.
	ExitAlter_type(c *Alter_typeContext)

	// ExitCompile_type_clause is called when exiting the compile_type_clause production.
	ExitCompile_type_clause(c *Compile_type_clauseContext)

	// ExitReplace_type_clause is called when exiting the replace_type_clause production.
	ExitReplace_type_clause(c *Replace_type_clauseContext)

	// ExitAlter_method_spec is called when exiting the alter_method_spec production.
	ExitAlter_method_spec(c *Alter_method_specContext)

	// ExitAlter_method_element is called when exiting the alter_method_element production.
	ExitAlter_method_element(c *Alter_method_elementContext)

	// ExitAlter_attribute_definition is called when exiting the alter_attribute_definition production.
	ExitAlter_attribute_definition(c *Alter_attribute_definitionContext)

	// ExitAttribute_definition is called when exiting the attribute_definition production.
	ExitAttribute_definition(c *Attribute_definitionContext)

	// ExitAlter_collection_clauses is called when exiting the alter_collection_clauses production.
	ExitAlter_collection_clauses(c *Alter_collection_clausesContext)

	// ExitDependent_handling_clause is called when exiting the dependent_handling_clause production.
	ExitDependent_handling_clause(c *Dependent_handling_clauseContext)

	// ExitDependent_exceptions_part is called when exiting the dependent_exceptions_part production.
	ExitDependent_exceptions_part(c *Dependent_exceptions_partContext)

	// ExitCreate_type is called when exiting the create_type production.
	ExitCreate_type(c *Create_typeContext)

	// ExitType_definition is called when exiting the type_definition production.
	ExitType_definition(c *Type_definitionContext)

	// ExitObject_type_def is called when exiting the object_type_def production.
	ExitObject_type_def(c *Object_type_defContext)

	// ExitObject_as_part is called when exiting the object_as_part production.
	ExitObject_as_part(c *Object_as_partContext)

	// ExitObject_under_part is called when exiting the object_under_part production.
	ExitObject_under_part(c *Object_under_partContext)

	// ExitNested_table_type_def is called when exiting the nested_table_type_def production.
	ExitNested_table_type_def(c *Nested_table_type_defContext)

	// ExitSqlj_object_type is called when exiting the sqlj_object_type production.
	ExitSqlj_object_type(c *Sqlj_object_typeContext)

	// ExitType_body is called when exiting the type_body production.
	ExitType_body(c *Type_bodyContext)

	// ExitType_body_elements is called when exiting the type_body_elements production.
	ExitType_body_elements(c *Type_body_elementsContext)

	// ExitMap_order_func_declaration is called when exiting the map_order_func_declaration production.
	ExitMap_order_func_declaration(c *Map_order_func_declarationContext)

	// ExitSubprog_decl_in_type is called when exiting the subprog_decl_in_type production.
	ExitSubprog_decl_in_type(c *Subprog_decl_in_typeContext)

	// ExitProc_decl_in_type is called when exiting the proc_decl_in_type production.
	ExitProc_decl_in_type(c *Proc_decl_in_typeContext)

	// ExitFunc_decl_in_type is called when exiting the func_decl_in_type production.
	ExitFunc_decl_in_type(c *Func_decl_in_typeContext)

	// ExitConstructor_declaration is called when exiting the constructor_declaration production.
	ExitConstructor_declaration(c *Constructor_declarationContext)

	// ExitModifier_clause is called when exiting the modifier_clause production.
	ExitModifier_clause(c *Modifier_clauseContext)

	// ExitObject_member_spec is called when exiting the object_member_spec production.
	ExitObject_member_spec(c *Object_member_specContext)

	// ExitSqlj_object_type_attr is called when exiting the sqlj_object_type_attr production.
	ExitSqlj_object_type_attr(c *Sqlj_object_type_attrContext)

	// ExitElement_spec is called when exiting the element_spec production.
	ExitElement_spec(c *Element_specContext)

	// ExitElement_spec_options is called when exiting the element_spec_options production.
	ExitElement_spec_options(c *Element_spec_optionsContext)

	// ExitSubprogram_spec is called when exiting the subprogram_spec production.
	ExitSubprogram_spec(c *Subprogram_specContext)

	// ExitType_procedure_spec is called when exiting the type_procedure_spec production.
	ExitType_procedure_spec(c *Type_procedure_specContext)

	// ExitType_function_spec is called when exiting the type_function_spec production.
	ExitType_function_spec(c *Type_function_specContext)

	// ExitConstructor_spec is called when exiting the constructor_spec production.
	ExitConstructor_spec(c *Constructor_specContext)

	// ExitMap_order_function_spec is called when exiting the map_order_function_spec production.
	ExitMap_order_function_spec(c *Map_order_function_specContext)

	// ExitPragma_clause is called when exiting the pragma_clause production.
	ExitPragma_clause(c *Pragma_clauseContext)

	// ExitPragma_elements is called when exiting the pragma_elements production.
	ExitPragma_elements(c *Pragma_elementsContext)

	// ExitType_elements_parameter is called when exiting the type_elements_parameter production.
	ExitType_elements_parameter(c *Type_elements_parameterContext)

	// ExitDrop_sequence is called when exiting the drop_sequence production.
	ExitDrop_sequence(c *Drop_sequenceContext)

	// ExitAlter_sequence is called when exiting the alter_sequence production.
	ExitAlter_sequence(c *Alter_sequenceContext)

	// ExitCreate_sequence is called when exiting the create_sequence production.
	ExitCreate_sequence(c *Create_sequenceContext)

	// ExitSequence_spec is called when exiting the sequence_spec production.
	ExitSequence_spec(c *Sequence_specContext)

	// ExitSequence_start_clause is called when exiting the sequence_start_clause production.
	ExitSequence_start_clause(c *Sequence_start_clauseContext)

	// ExitCreate_index is called when exiting the create_index production.
	ExitCreate_index(c *Create_indexContext)

	// ExitAlter_index is called when exiting the alter_index production.
	ExitAlter_index(c *Alter_indexContext)

	// ExitDrop_index is called when exiting the drop_index production.
	ExitDrop_index(c *Drop_indexContext)

	// ExitCreate_table is called when exiting the create_table production.
	ExitCreate_table(c *Create_tableContext)

	// ExitSize_clause is called when exiting the size_clause production.
	ExitSize_clause(c *Size_clauseContext)

	// ExitDrop_table is called when exiting the drop_table production.
	ExitDrop_table(c *Drop_tableContext)

	// ExitComment_on_column is called when exiting the comment_on_column production.
	ExitComment_on_column(c *Comment_on_columnContext)

	// ExitCreate_synonym is called when exiting the create_synonym production.
	ExitCreate_synonym(c *Create_synonymContext)

	// ExitComment_on_table is called when exiting the comment_on_table production.
	ExitComment_on_table(c *Comment_on_tableContext)

	// ExitAlter_table is called when exiting the alter_table production.
	ExitAlter_table(c *Alter_tableContext)

	// ExitAdd_constraint is called when exiting the add_constraint production.
	ExitAdd_constraint(c *Add_constraintContext)

	// ExitCheck_constraint is called when exiting the check_constraint production.
	ExitCheck_constraint(c *Check_constraintContext)

	// ExitDrop_constraint is called when exiting the drop_constraint production.
	ExitDrop_constraint(c *Drop_constraintContext)

	// ExitEnable_constraint is called when exiting the enable_constraint production.
	ExitEnable_constraint(c *Enable_constraintContext)

	// ExitDisable_constraint is called when exiting the disable_constraint production.
	ExitDisable_constraint(c *Disable_constraintContext)

	// ExitForeign_key_clause is called when exiting the foreign_key_clause production.
	ExitForeign_key_clause(c *Foreign_key_clauseContext)

	// ExitReferences_clause is called when exiting the references_clause production.
	ExitReferences_clause(c *References_clauseContext)

	// ExitOn_delete_clause is called when exiting the on_delete_clause production.
	ExitOn_delete_clause(c *On_delete_clauseContext)

	// ExitUnique_key_clause is called when exiting the unique_key_clause production.
	ExitUnique_key_clause(c *Unique_key_clauseContext)

	// ExitPrimary_key_clause is called when exiting the primary_key_clause production.
	ExitPrimary_key_clause(c *Primary_key_clauseContext)

	// ExitAnonymous_block is called when exiting the anonymous_block production.
	ExitAnonymous_block(c *Anonymous_blockContext)

	// ExitInvoker_rights_clause is called when exiting the invoker_rights_clause production.
	ExitInvoker_rights_clause(c *Invoker_rights_clauseContext)

	// ExitCompiler_parameters_clause is called when exiting the compiler_parameters_clause production.
	ExitCompiler_parameters_clause(c *Compiler_parameters_clauseContext)

	// ExitCall_spec is called when exiting the call_spec production.
	ExitCall_spec(c *Call_specContext)

	// ExitJava_spec is called when exiting the java_spec production.
	ExitJava_spec(c *Java_specContext)

	// ExitC_spec is called when exiting the c_spec production.
	ExitC_spec(c *C_specContext)

	// ExitC_agent_in_clause is called when exiting the c_agent_in_clause production.
	ExitC_agent_in_clause(c *C_agent_in_clauseContext)

	// ExitC_parameters_clause is called when exiting the c_parameters_clause production.
	ExitC_parameters_clause(c *C_parameters_clauseContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitDefault_value_part is called when exiting the default_value_part production.
	ExitDefault_value_part(c *Default_value_partContext)

	// ExitDeclare_spec is called when exiting the declare_spec production.
	ExitDeclare_spec(c *Declare_specContext)

	// ExitVariable_declaration is called when exiting the variable_declaration production.
	ExitVariable_declaration(c *Variable_declarationContext)

	// ExitSubtype_declaration is called when exiting the subtype_declaration production.
	ExitSubtype_declaration(c *Subtype_declarationContext)

	// ExitCursor_declaration is called when exiting the cursor_declaration production.
	ExitCursor_declaration(c *Cursor_declarationContext)

	// ExitParameter_spec is called when exiting the parameter_spec production.
	ExitParameter_spec(c *Parameter_specContext)

	// ExitException_declaration is called when exiting the exception_declaration production.
	ExitException_declaration(c *Exception_declarationContext)

	// ExitPragma_declaration is called when exiting the pragma_declaration production.
	ExitPragma_declaration(c *Pragma_declarationContext)

	// ExitRecord_type_def is called when exiting the record_type_def production.
	ExitRecord_type_def(c *Record_type_defContext)

	// ExitField_spec is called when exiting the field_spec production.
	ExitField_spec(c *Field_specContext)

	// ExitRef_cursor_type_def is called when exiting the ref_cursor_type_def production.
	ExitRef_cursor_type_def(c *Ref_cursor_type_defContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitTable_type_def is called when exiting the table_type_def production.
	ExitTable_type_def(c *Table_type_defContext)

	// ExitTable_indexed_by_part is called when exiting the table_indexed_by_part production.
	ExitTable_indexed_by_part(c *Table_indexed_by_partContext)

	// ExitVarray_type_def is called when exiting the varray_type_def production.
	ExitVarray_type_def(c *Varray_type_defContext)

	// ExitSeq_of_statements is called when exiting the seq_of_statements production.
	ExitSeq_of_statements(c *Seq_of_statementsContext)

	// ExitLabel_declaration is called when exiting the label_declaration production.
	ExitLabel_declaration(c *Label_declarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSwallow_to_semi is called when exiting the swallow_to_semi production.
	ExitSwallow_to_semi(c *Swallow_to_semiContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitExit_statement is called when exiting the exit_statement production.
	ExitExit_statement(c *Exit_statementContext)

	// ExitGoto_statement is called when exiting the goto_statement production.
	ExitGoto_statement(c *Goto_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitElsif_part is called when exiting the elsif_part production.
	ExitElsif_part(c *Elsif_partContext)

	// ExitElse_part is called when exiting the else_part production.
	ExitElse_part(c *Else_partContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitCursor_loop_param is called when exiting the cursor_loop_param production.
	ExitCursor_loop_param(c *Cursor_loop_paramContext)

	// ExitForall_statement is called when exiting the forall_statement production.
	ExitForall_statement(c *Forall_statementContext)

	// ExitBounds_clause is called when exiting the bounds_clause production.
	ExitBounds_clause(c *Bounds_clauseContext)

	// ExitBetween_bound is called when exiting the between_bound production.
	ExitBetween_bound(c *Between_boundContext)

	// ExitLower_bound is called when exiting the lower_bound production.
	ExitLower_bound(c *Lower_boundContext)

	// ExitUpper_bound is called when exiting the upper_bound production.
	ExitUpper_bound(c *Upper_boundContext)

	// ExitNull_statement is called when exiting the null_statement production.
	ExitNull_statement(c *Null_statementContext)

	// ExitRaise_statement is called when exiting the raise_statement production.
	ExitRaise_statement(c *Raise_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitPipe_row_statement is called when exiting the pipe_row_statement production.
	ExitPipe_row_statement(c *Pipe_row_statementContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitException_handler is called when exiting the exception_handler production.
	ExitException_handler(c *Exception_handlerContext)

	// ExitTrigger_block is called when exiting the trigger_block production.
	ExitTrigger_block(c *Trigger_blockContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSql_statement is called when exiting the sql_statement production.
	ExitSql_statement(c *Sql_statementContext)

	// ExitExecute_immediate is called when exiting the execute_immediate production.
	ExitExecute_immediate(c *Execute_immediateContext)

	// ExitDynamic_returning_clause is called when exiting the dynamic_returning_clause production.
	ExitDynamic_returning_clause(c *Dynamic_returning_clauseContext)

	// ExitData_manipulation_language_statements is called when exiting the data_manipulation_language_statements production.
	ExitData_manipulation_language_statements(c *Data_manipulation_language_statementsContext)

	// ExitCursor_manipulation_statements is called when exiting the cursor_manipulation_statements production.
	ExitCursor_manipulation_statements(c *Cursor_manipulation_statementsContext)

	// ExitClose_statement is called when exiting the close_statement production.
	ExitClose_statement(c *Close_statementContext)

	// ExitOpen_statement is called when exiting the open_statement production.
	ExitOpen_statement(c *Open_statementContext)

	// ExitFetch_statement is called when exiting the fetch_statement production.
	ExitFetch_statement(c *Fetch_statementContext)

	// ExitOpen_for_statement is called when exiting the open_for_statement production.
	ExitOpen_for_statement(c *Open_for_statementContext)

	// ExitTransaction_control_statements is called when exiting the transaction_control_statements production.
	ExitTransaction_control_statements(c *Transaction_control_statementsContext)

	// ExitSet_transaction_command is called when exiting the set_transaction_command production.
	ExitSet_transaction_command(c *Set_transaction_commandContext)

	// ExitSet_constraint_command is called when exiting the set_constraint_command production.
	ExitSet_constraint_command(c *Set_constraint_commandContext)

	// ExitCommit_statement is called when exiting the commit_statement production.
	ExitCommit_statement(c *Commit_statementContext)

	// ExitWrite_clause is called when exiting the write_clause production.
	ExitWrite_clause(c *Write_clauseContext)

	// ExitRollback_statement is called when exiting the rollback_statement production.
	ExitRollback_statement(c *Rollback_statementContext)

	// ExitSavepoint_statement is called when exiting the savepoint_statement production.
	ExitSavepoint_statement(c *Savepoint_statementContext)

	// ExitExplain_statement is called when exiting the explain_statement production.
	ExitExplain_statement(c *Explain_statementContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitSubquery_factoring_clause is called when exiting the subquery_factoring_clause production.
	ExitSubquery_factoring_clause(c *Subquery_factoring_clauseContext)

	// ExitFactoring_element is called when exiting the factoring_element production.
	ExitFactoring_element(c *Factoring_elementContext)

	// ExitSearch_clause is called when exiting the search_clause production.
	ExitSearch_clause(c *Search_clauseContext)

	// ExitCycle_clause is called when exiting the cycle_clause production.
	ExitCycle_clause(c *Cycle_clauseContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitSubquery_basic_elements is called when exiting the subquery_basic_elements production.
	ExitSubquery_basic_elements(c *Subquery_basic_elementsContext)

	// ExitSubquery_operation_part is called when exiting the subquery_operation_part production.
	ExitSubquery_operation_part(c *Subquery_operation_partContext)

	// ExitQuery_block is called when exiting the query_block production.
	ExitQuery_block(c *Query_blockContext)

	// ExitSelected_element is called when exiting the selected_element production.
	ExitSelected_element(c *Selected_elementContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitSelect_list_elements is called when exiting the select_list_elements production.
	ExitSelect_list_elements(c *Select_list_elementsContext)

	// ExitTable_ref_list is called when exiting the table_ref_list production.
	ExitTable_ref_list(c *Table_ref_listContext)

	// ExitTable_ref is called when exiting the table_ref production.
	ExitTable_ref(c *Table_refContext)

	// ExitTable_ref_aux is called when exiting the table_ref_aux production.
	ExitTable_ref_aux(c *Table_ref_auxContext)

	// ExitTable_ref_aux_internal_one is called when exiting the table_ref_aux_internal_one production.
	ExitTable_ref_aux_internal_one(c *Table_ref_aux_internal_oneContext)

	// ExitTable_ref_aux_internal_two is called when exiting the table_ref_aux_internal_two production.
	ExitTable_ref_aux_internal_two(c *Table_ref_aux_internal_twoContext)

	// ExitTable_ref_aux_internal_three is called when exiting the table_ref_aux_internal_three production.
	ExitTable_ref_aux_internal_three(c *Table_ref_aux_internal_threeContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitJoin_on_part is called when exiting the join_on_part production.
	ExitJoin_on_part(c *Join_on_partContext)

	// ExitJoin_using_part is called when exiting the join_using_part production.
	ExitJoin_using_part(c *Join_using_partContext)

	// ExitOuter_join_type is called when exiting the outer_join_type production.
	ExitOuter_join_type(c *Outer_join_typeContext)

	// ExitQuery_partition_clause is called when exiting the query_partition_clause production.
	ExitQuery_partition_clause(c *Query_partition_clauseContext)

	// ExitFlashback_query_clause is called when exiting the flashback_query_clause production.
	ExitFlashback_query_clause(c *Flashback_query_clauseContext)

	// ExitPivot_clause is called when exiting the pivot_clause production.
	ExitPivot_clause(c *Pivot_clauseContext)

	// ExitPivot_element is called when exiting the pivot_element production.
	ExitPivot_element(c *Pivot_elementContext)

	// ExitPivot_for_clause is called when exiting the pivot_for_clause production.
	ExitPivot_for_clause(c *Pivot_for_clauseContext)

	// ExitPivot_in_clause is called when exiting the pivot_in_clause production.
	ExitPivot_in_clause(c *Pivot_in_clauseContext)

	// ExitPivot_in_clause_element is called when exiting the pivot_in_clause_element production.
	ExitPivot_in_clause_element(c *Pivot_in_clause_elementContext)

	// ExitPivot_in_clause_elements is called when exiting the pivot_in_clause_elements production.
	ExitPivot_in_clause_elements(c *Pivot_in_clause_elementsContext)

	// ExitUnpivot_clause is called when exiting the unpivot_clause production.
	ExitUnpivot_clause(c *Unpivot_clauseContext)

	// ExitUnpivot_in_clause is called when exiting the unpivot_in_clause production.
	ExitUnpivot_in_clause(c *Unpivot_in_clauseContext)

	// ExitUnpivot_in_elements is called when exiting the unpivot_in_elements production.
	ExitUnpivot_in_elements(c *Unpivot_in_elementsContext)

	// ExitHierarchical_query_clause is called when exiting the hierarchical_query_clause production.
	ExitHierarchical_query_clause(c *Hierarchical_query_clauseContext)

	// ExitStart_part is called when exiting the start_part production.
	ExitStart_part(c *Start_partContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitGroup_by_elements is called when exiting the group_by_elements production.
	ExitGroup_by_elements(c *Group_by_elementsContext)

	// ExitRollup_cube_clause is called when exiting the rollup_cube_clause production.
	ExitRollup_cube_clause(c *Rollup_cube_clauseContext)

	// ExitGrouping_sets_clause is called when exiting the grouping_sets_clause production.
	ExitGrouping_sets_clause(c *Grouping_sets_clauseContext)

	// ExitGrouping_sets_elements is called when exiting the grouping_sets_elements production.
	ExitGrouping_sets_elements(c *Grouping_sets_elementsContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitModel_clause is called when exiting the model_clause production.
	ExitModel_clause(c *Model_clauseContext)

	// ExitCell_reference_options is called when exiting the cell_reference_options production.
	ExitCell_reference_options(c *Cell_reference_optionsContext)

	// ExitReturn_rows_clause is called when exiting the return_rows_clause production.
	ExitReturn_rows_clause(c *Return_rows_clauseContext)

	// ExitReference_model is called when exiting the reference_model production.
	ExitReference_model(c *Reference_modelContext)

	// ExitMain_model is called when exiting the main_model production.
	ExitMain_model(c *Main_modelContext)

	// ExitModel_column_clauses is called when exiting the model_column_clauses production.
	ExitModel_column_clauses(c *Model_column_clausesContext)

	// ExitModel_column_partition_part is called when exiting the model_column_partition_part production.
	ExitModel_column_partition_part(c *Model_column_partition_partContext)

	// ExitModel_column_list is called when exiting the model_column_list production.
	ExitModel_column_list(c *Model_column_listContext)

	// ExitModel_column is called when exiting the model_column production.
	ExitModel_column(c *Model_columnContext)

	// ExitModel_rules_clause is called when exiting the model_rules_clause production.
	ExitModel_rules_clause(c *Model_rules_clauseContext)

	// ExitModel_rules_part is called when exiting the model_rules_part production.
	ExitModel_rules_part(c *Model_rules_partContext)

	// ExitModel_rules_element is called when exiting the model_rules_element production.
	ExitModel_rules_element(c *Model_rules_elementContext)

	// ExitCell_assignment is called when exiting the cell_assignment production.
	ExitCell_assignment(c *Cell_assignmentContext)

	// ExitModel_iterate_clause is called when exiting the model_iterate_clause production.
	ExitModel_iterate_clause(c *Model_iterate_clauseContext)

	// ExitUntil_part is called when exiting the until_part production.
	ExitUntil_part(c *Until_partContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitOrder_by_elements is called when exiting the order_by_elements production.
	ExitOrder_by_elements(c *Order_by_elementsContext)

	// ExitFor_update_clause is called when exiting the for_update_clause production.
	ExitFor_update_clause(c *For_update_clauseContext)

	// ExitFor_update_of_part is called when exiting the for_update_of_part production.
	ExitFor_update_of_part(c *For_update_of_partContext)

	// ExitFor_update_options is called when exiting the for_update_options production.
	ExitFor_update_options(c *For_update_optionsContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitUpdate_set_clause is called when exiting the update_set_clause production.
	ExitUpdate_set_clause(c *Update_set_clauseContext)

	// ExitColumn_based_update_set_clause is called when exiting the column_based_update_set_clause production.
	ExitColumn_based_update_set_clause(c *Column_based_update_set_clauseContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitInsert_statement is called when exiting the insert_statement production.
	ExitInsert_statement(c *Insert_statementContext)

	// ExitSingle_table_insert is called when exiting the single_table_insert production.
	ExitSingle_table_insert(c *Single_table_insertContext)

	// ExitMulti_table_insert is called when exiting the multi_table_insert production.
	ExitMulti_table_insert(c *Multi_table_insertContext)

	// ExitMulti_table_element is called when exiting the multi_table_element production.
	ExitMulti_table_element(c *Multi_table_elementContext)

	// ExitConditional_insert_clause is called when exiting the conditional_insert_clause production.
	ExitConditional_insert_clause(c *Conditional_insert_clauseContext)

	// ExitConditional_insert_when_part is called when exiting the conditional_insert_when_part production.
	ExitConditional_insert_when_part(c *Conditional_insert_when_partContext)

	// ExitConditional_insert_else_part is called when exiting the conditional_insert_else_part production.
	ExitConditional_insert_else_part(c *Conditional_insert_else_partContext)

	// ExitInsert_into_clause is called when exiting the insert_into_clause production.
	ExitInsert_into_clause(c *Insert_into_clauseContext)

	// ExitValues_clause is called when exiting the values_clause production.
	ExitValues_clause(c *Values_clauseContext)

	// ExitMerge_statement is called when exiting the merge_statement production.
	ExitMerge_statement(c *Merge_statementContext)

	// ExitMerge_update_clause is called when exiting the merge_update_clause production.
	ExitMerge_update_clause(c *Merge_update_clauseContext)

	// ExitMerge_element is called when exiting the merge_element production.
	ExitMerge_element(c *Merge_elementContext)

	// ExitMerge_update_delete_part is called when exiting the merge_update_delete_part production.
	ExitMerge_update_delete_part(c *Merge_update_delete_partContext)

	// ExitMerge_insert_clause is called when exiting the merge_insert_clause production.
	ExitMerge_insert_clause(c *Merge_insert_clauseContext)

	// ExitSelected_tableview is called when exiting the selected_tableview production.
	ExitSelected_tableview(c *Selected_tableviewContext)

	// ExitLock_table_statement is called when exiting the lock_table_statement production.
	ExitLock_table_statement(c *Lock_table_statementContext)

	// ExitWait_nowait_part is called when exiting the wait_nowait_part production.
	ExitWait_nowait_part(c *Wait_nowait_partContext)

	// ExitLock_table_element is called when exiting the lock_table_element production.
	ExitLock_table_element(c *Lock_table_elementContext)

	// ExitLock_mode is called when exiting the lock_mode production.
	ExitLock_mode(c *Lock_modeContext)

	// ExitGeneral_table_ref is called when exiting the general_table_ref production.
	ExitGeneral_table_ref(c *General_table_refContext)

	// ExitStatic_returning_clause is called when exiting the static_returning_clause production.
	ExitStatic_returning_clause(c *Static_returning_clauseContext)

	// ExitError_logging_clause is called when exiting the error_logging_clause production.
	ExitError_logging_clause(c *Error_logging_clauseContext)

	// ExitError_logging_into_part is called when exiting the error_logging_into_part production.
	ExitError_logging_into_part(c *Error_logging_into_partContext)

	// ExitError_logging_reject_part is called when exiting the error_logging_reject_part production.
	ExitError_logging_reject_part(c *Error_logging_reject_partContext)

	// ExitDml_table_expression_clause is called when exiting the dml_table_expression_clause production.
	ExitDml_table_expression_clause(c *Dml_table_expression_clauseContext)

	// ExitTable_collection_expression is called when exiting the table_collection_expression production.
	ExitTable_collection_expression(c *Table_collection_expressionContext)

	// ExitSubquery_restriction_clause is called when exiting the subquery_restriction_clause production.
	ExitSubquery_restriction_clause(c *Subquery_restriction_clauseContext)

	// ExitSample_clause is called when exiting the sample_clause production.
	ExitSample_clause(c *Sample_clauseContext)

	// ExitSeed_part is called when exiting the seed_part production.
	ExitSeed_part(c *Seed_partContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCursor_expression is called when exiting the cursor_expression production.
	ExitCursor_expression(c *Cursor_expressionContext)

	// ExitLogical_expression is called when exiting the logical_expression production.
	ExitLogical_expression(c *Logical_expressionContext)

	// ExitMultiset_expression is called when exiting the multiset_expression production.
	ExitMultiset_expression(c *Multiset_expressionContext)

	// ExitRelational_expression is called when exiting the relational_expression production.
	ExitRelational_expression(c *Relational_expressionContext)

	// ExitCompound_expression is called when exiting the compound_expression production.
	ExitCompound_expression(c *Compound_expressionContext)

	// ExitRelational_operator is called when exiting the relational_operator production.
	ExitRelational_operator(c *Relational_operatorContext)

	// ExitIn_elements is called when exiting the in_elements production.
	ExitIn_elements(c *In_elementsContext)

	// ExitBetween_elements is called when exiting the between_elements production.
	ExitBetween_elements(c *Between_elementsContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitInterval_expression is called when exiting the interval_expression production.
	ExitInterval_expression(c *Interval_expressionContext)

	// ExitModel_expression is called when exiting the model_expression production.
	ExitModel_expression(c *Model_expressionContext)

	// ExitModel_expression_element is called when exiting the model_expression_element production.
	ExitModel_expression_element(c *Model_expression_elementContext)

	// ExitSingle_column_for_loop is called when exiting the single_column_for_loop production.
	ExitSingle_column_for_loop(c *Single_column_for_loopContext)

	// ExitMulti_column_for_loop is called when exiting the multi_column_for_loop production.
	ExitMulti_column_for_loop(c *Multi_column_for_loopContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitSimple_case_statement is called when exiting the simple_case_statement production.
	ExitSimple_case_statement(c *Simple_case_statementContext)

	// ExitSimple_case_when_part is called when exiting the simple_case_when_part production.
	ExitSimple_case_when_part(c *Simple_case_when_partContext)

	// ExitSearched_case_statement is called when exiting the searched_case_statement production.
	ExitSearched_case_statement(c *Searched_case_statementContext)

	// ExitSearched_case_when_part is called when exiting the searched_case_when_part production.
	ExitSearched_case_when_part(c *Searched_case_when_partContext)

	// ExitCase_else_part is called when exiting the case_else_part production.
	ExitCase_else_part(c *Case_else_partContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitQuantified_expression is called when exiting the quantified_expression production.
	ExitQuantified_expression(c *Quantified_expressionContext)

	// ExitString_function is called when exiting the string_function production.
	ExitString_function(c *String_functionContext)

	// ExitStandard_function is called when exiting the standard_function production.
	ExitStandard_function(c *Standard_functionContext)

	// ExitNumeric_function_wrapper is called when exiting the numeric_function_wrapper production.
	ExitNumeric_function_wrapper(c *Numeric_function_wrapperContext)

	// ExitNumeric_function is called when exiting the numeric_function production.
	ExitNumeric_function(c *Numeric_functionContext)

	// ExitOther_function is called when exiting the other_function production.
	ExitOther_function(c *Other_functionContext)

	// ExitOver_clause_keyword is called when exiting the over_clause_keyword production.
	ExitOver_clause_keyword(c *Over_clause_keywordContext)

	// ExitWithin_or_over_clause_keyword is called when exiting the within_or_over_clause_keyword production.
	ExitWithin_or_over_clause_keyword(c *Within_or_over_clause_keywordContext)

	// ExitStandard_prediction_function_keyword is called when exiting the standard_prediction_function_keyword production.
	ExitStandard_prediction_function_keyword(c *Standard_prediction_function_keywordContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitWindowing_clause is called when exiting the windowing_clause production.
	ExitWindowing_clause(c *Windowing_clauseContext)

	// ExitWindowing_type is called when exiting the windowing_type production.
	ExitWindowing_type(c *Windowing_typeContext)

	// ExitWindowing_elements is called when exiting the windowing_elements production.
	ExitWindowing_elements(c *Windowing_elementsContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitUsing_element is called when exiting the using_element production.
	ExitUsing_element(c *Using_elementContext)

	// ExitCollect_order_by_part is called when exiting the collect_order_by_part production.
	ExitCollect_order_by_part(c *Collect_order_by_partContext)

	// ExitWithin_or_over_part is called when exiting the within_or_over_part production.
	ExitWithin_or_over_part(c *Within_or_over_partContext)

	// ExitCost_matrix_clause is called when exiting the cost_matrix_clause production.
	ExitCost_matrix_clause(c *Cost_matrix_clauseContext)

	// ExitXml_passing_clause is called when exiting the xml_passing_clause production.
	ExitXml_passing_clause(c *Xml_passing_clauseContext)

	// ExitXml_attributes_clause is called when exiting the xml_attributes_clause production.
	ExitXml_attributes_clause(c *Xml_attributes_clauseContext)

	// ExitXml_namespaces_clause is called when exiting the xml_namespaces_clause production.
	ExitXml_namespaces_clause(c *Xml_namespaces_clauseContext)

	// ExitXml_table_column is called when exiting the xml_table_column production.
	ExitXml_table_column(c *Xml_table_columnContext)

	// ExitXml_general_default_part is called when exiting the xml_general_default_part production.
	ExitXml_general_default_part(c *Xml_general_default_partContext)

	// ExitXml_multiuse_expression_element is called when exiting the xml_multiuse_expression_element production.
	ExitXml_multiuse_expression_element(c *Xml_multiuse_expression_elementContext)

	// ExitXmlroot_param_version_part is called when exiting the xmlroot_param_version_part production.
	ExitXmlroot_param_version_part(c *Xmlroot_param_version_partContext)

	// ExitXmlroot_param_standalone_part is called when exiting the xmlroot_param_standalone_part production.
	ExitXmlroot_param_standalone_part(c *Xmlroot_param_standalone_partContext)

	// ExitXmlserialize_param_enconding_part is called when exiting the xmlserialize_param_enconding_part production.
	ExitXmlserialize_param_enconding_part(c *Xmlserialize_param_enconding_partContext)

	// ExitXmlserialize_param_version_part is called when exiting the xmlserialize_param_version_part production.
	ExitXmlserialize_param_version_part(c *Xmlserialize_param_version_partContext)

	// ExitXmlserialize_param_ident_part is called when exiting the xmlserialize_param_ident_part production.
	ExitXmlserialize_param_ident_part(c *Xmlserialize_param_ident_partContext)

	// ExitSql_plus_command is called when exiting the sql_plus_command production.
	ExitSql_plus_command(c *Sql_plus_commandContext)

	// ExitWhenever_command is called when exiting the whenever_command production.
	ExitWhenever_command(c *Whenever_commandContext)

	// ExitSet_command is called when exiting the set_command production.
	ExitSet_command(c *Set_commandContext)

	// ExitPartition_extension_clause is called when exiting the partition_extension_clause production.
	ExitPartition_extension_clause(c *Partition_extension_clauseContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitInto_clause is called when exiting the into_clause production.
	ExitInto_clause(c *Into_clauseContext)

	// ExitXml_column_name is called when exiting the xml_column_name production.
	ExitXml_column_name(c *Xml_column_nameContext)

	// ExitCost_class_name is called when exiting the cost_class_name production.
	ExitCost_class_name(c *Cost_class_nameContext)

	// ExitAttribute_name is called when exiting the attribute_name production.
	ExitAttribute_name(c *Attribute_nameContext)

	// ExitSavepoint_name is called when exiting the savepoint_name production.
	ExitSavepoint_name(c *Savepoint_nameContext)

	// ExitRollback_segment_name is called when exiting the rollback_segment_name production.
	ExitRollback_segment_name(c *Rollback_segment_nameContext)

	// ExitTable_var_name is called when exiting the table_var_name production.
	ExitTable_var_name(c *Table_var_nameContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitRoutine_name is called when exiting the routine_name production.
	ExitRoutine_name(c *Routine_nameContext)

	// ExitPackage_name is called when exiting the package_name production.
	ExitPackage_name(c *Package_nameContext)

	// ExitImplementation_type_name is called when exiting the implementation_type_name production.
	ExitImplementation_type_name(c *Implementation_type_nameContext)

	// ExitParameter_name is called when exiting the parameter_name production.
	ExitParameter_name(c *Parameter_nameContext)

	// ExitReference_model_name is called when exiting the reference_model_name production.
	ExitReference_model_name(c *Reference_model_nameContext)

	// ExitMain_model_name is called when exiting the main_model_name production.
	ExitMain_model_name(c *Main_model_nameContext)

	// ExitAggregate_function_name is called when exiting the aggregate_function_name production.
	ExitAggregate_function_name(c *Aggregate_function_nameContext)

	// ExitQuery_name is called when exiting the query_name production.
	ExitQuery_name(c *Query_nameContext)

	// ExitConstraint_name is called when exiting the constraint_name production.
	ExitConstraint_name(c *Constraint_nameContext)

	// ExitLabel_name is called when exiting the label_name production.
	ExitLabel_name(c *Label_nameContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitSequence_name is called when exiting the sequence_name production.
	ExitSequence_name(c *Sequence_nameContext)

	// ExitException_name is called when exiting the exception_name production.
	ExitException_name(c *Exception_nameContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitProcedure_name is called when exiting the procedure_name production.
	ExitProcedure_name(c *Procedure_nameContext)

	// ExitTrigger_name is called when exiting the trigger_name production.
	ExitTrigger_name(c *Trigger_nameContext)

	// ExitVariable_name is called when exiting the variable_name production.
	ExitVariable_name(c *Variable_nameContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitCursor_name is called when exiting the cursor_name production.
	ExitCursor_name(c *Cursor_nameContext)

	// ExitRecord_name is called when exiting the record_name production.
	ExitRecord_name(c *Record_nameContext)

	// ExitCollection_name is called when exiting the collection_name production.
	ExitCollection_name(c *Collection_nameContext)

	// ExitLink_name is called when exiting the link_name production.
	ExitLink_name(c *Link_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitTableview_name is called when exiting the tableview_name production.
	ExitTableview_name(c *Tableview_nameContext)

	// ExitChar_set_name is called when exiting the char_set_name production.
	ExitChar_set_name(c *Char_set_nameContext)

	// ExitSynonym_name is called when exiting the synonym_name production.
	ExitSynonym_name(c *Synonym_nameContext)

	// ExitSchema_object_name is called when exiting the schema_object_name production.
	ExitSchema_object_name(c *Schema_object_nameContext)

	// ExitKeep_clause is called when exiting the keep_clause production.
	ExitKeep_clause(c *Keep_clauseContext)

	// ExitFunction_argument is called when exiting the function_argument production.
	ExitFunction_argument(c *Function_argumentContext)

	// ExitFunction_argument_analytic is called when exiting the function_argument_analytic production.
	ExitFunction_argument_analytic(c *Function_argument_analyticContext)

	// ExitFunction_argument_modeling is called when exiting the function_argument_modeling production.
	ExitFunction_argument_modeling(c *Function_argument_modelingContext)

	// ExitRespect_or_ignore_nulls is called when exiting the respect_or_ignore_nulls production.
	ExitRespect_or_ignore_nulls(c *Respect_or_ignore_nullsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitType_spec is called when exiting the type_spec production.
	ExitType_spec(c *Type_specContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitPrecision_part is called when exiting the precision_part production.
	ExitPrecision_part(c *Precision_partContext)

	// ExitNative_datatype_element is called when exiting the native_datatype_element production.
	ExitNative_datatype_element(c *Native_datatype_elementContext)

	// ExitBind_variable is called when exiting the bind_variable production.
	ExitBind_variable(c *Bind_variableContext)

	// ExitGeneral_element is called when exiting the general_element production.
	ExitGeneral_element(c *General_elementContext)

	// ExitGeneral_element_part is called when exiting the general_element_part production.
	ExitGeneral_element_part(c *General_element_partContext)

	// ExitTable_element is called when exiting the table_element production.
	ExitTable_element(c *Table_elementContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumeric is called when exiting the numeric production.
	ExitNumeric(c *NumericContext)

	// ExitNumeric_negative is called when exiting the numeric_negative production.
	ExitNumeric_negative(c *Numeric_negativeContext)

	// ExitQuoted_string is called when exiting the quoted_string production.
	ExitQuoted_string(c *Quoted_stringContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitId_expression is called when exiting the id_expression production.
	ExitId_expression(c *Id_expressionContext)

	// ExitOuter_join_sign is called when exiting the outer_join_sign production.
	ExitOuter_join_sign(c *Outer_join_signContext)

	// ExitRegular_id is called when exiting the regular_id production.
	ExitRegular_id(c *Regular_idContext)

	// ExitString_function_name is called when exiting the string_function_name production.
	ExitString_function_name(c *String_function_nameContext)

	// ExitNumeric_function_name is called when exiting the numeric_function_name production.
	ExitNumeric_function_name(c *Numeric_function_nameContext)
}
