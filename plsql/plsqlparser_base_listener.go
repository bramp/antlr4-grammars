// Code generated from PlSqlParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package plsql // PlSqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePlSqlParserListener is a complete listener for a parse tree produced by PlSqlParser.
type BasePlSqlParserListener struct{}

var _ PlSqlParserListener = &BasePlSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSql_script is called when production sql_script is entered.
func (s *BasePlSqlParserListener) EnterSql_script(ctx *Sql_scriptContext) {}

// ExitSql_script is called when production sql_script is exited.
func (s *BasePlSqlParserListener) ExitSql_script(ctx *Sql_scriptContext) {}

// EnterUnit_statement is called when production unit_statement is entered.
func (s *BasePlSqlParserListener) EnterUnit_statement(ctx *Unit_statementContext) {}

// ExitUnit_statement is called when production unit_statement is exited.
func (s *BasePlSqlParserListener) ExitUnit_statement(ctx *Unit_statementContext) {}

// EnterDrop_function is called when production drop_function is entered.
func (s *BasePlSqlParserListener) EnterDrop_function(ctx *Drop_functionContext) {}

// ExitDrop_function is called when production drop_function is exited.
func (s *BasePlSqlParserListener) ExitDrop_function(ctx *Drop_functionContext) {}

// EnterAlter_function is called when production alter_function is entered.
func (s *BasePlSqlParserListener) EnterAlter_function(ctx *Alter_functionContext) {}

// ExitAlter_function is called when production alter_function is exited.
func (s *BasePlSqlParserListener) ExitAlter_function(ctx *Alter_functionContext) {}

// EnterCreate_function_body is called when production create_function_body is entered.
func (s *BasePlSqlParserListener) EnterCreate_function_body(ctx *Create_function_bodyContext) {}

// ExitCreate_function_body is called when production create_function_body is exited.
func (s *BasePlSqlParserListener) ExitCreate_function_body(ctx *Create_function_bodyContext) {}

// EnterParallel_enable_clause is called when production parallel_enable_clause is entered.
func (s *BasePlSqlParserListener) EnterParallel_enable_clause(ctx *Parallel_enable_clauseContext) {}

// ExitParallel_enable_clause is called when production parallel_enable_clause is exited.
func (s *BasePlSqlParserListener) ExitParallel_enable_clause(ctx *Parallel_enable_clauseContext) {}

// EnterPartition_by_clause is called when production partition_by_clause is entered.
func (s *BasePlSqlParserListener) EnterPartition_by_clause(ctx *Partition_by_clauseContext) {}

// ExitPartition_by_clause is called when production partition_by_clause is exited.
func (s *BasePlSqlParserListener) ExitPartition_by_clause(ctx *Partition_by_clauseContext) {}

// EnterResult_cache_clause is called when production result_cache_clause is entered.
func (s *BasePlSqlParserListener) EnterResult_cache_clause(ctx *Result_cache_clauseContext) {}

// ExitResult_cache_clause is called when production result_cache_clause is exited.
func (s *BasePlSqlParserListener) ExitResult_cache_clause(ctx *Result_cache_clauseContext) {}

// EnterRelies_on_part is called when production relies_on_part is entered.
func (s *BasePlSqlParserListener) EnterRelies_on_part(ctx *Relies_on_partContext) {}

// ExitRelies_on_part is called when production relies_on_part is exited.
func (s *BasePlSqlParserListener) ExitRelies_on_part(ctx *Relies_on_partContext) {}

// EnterStreaming_clause is called when production streaming_clause is entered.
func (s *BasePlSqlParserListener) EnterStreaming_clause(ctx *Streaming_clauseContext) {}

// ExitStreaming_clause is called when production streaming_clause is exited.
func (s *BasePlSqlParserListener) ExitStreaming_clause(ctx *Streaming_clauseContext) {}

// EnterDrop_package is called when production drop_package is entered.
func (s *BasePlSqlParserListener) EnterDrop_package(ctx *Drop_packageContext) {}

// ExitDrop_package is called when production drop_package is exited.
func (s *BasePlSqlParserListener) ExitDrop_package(ctx *Drop_packageContext) {}

// EnterAlter_package is called when production alter_package is entered.
func (s *BasePlSqlParserListener) EnterAlter_package(ctx *Alter_packageContext) {}

// ExitAlter_package is called when production alter_package is exited.
func (s *BasePlSqlParserListener) ExitAlter_package(ctx *Alter_packageContext) {}

// EnterCreate_package is called when production create_package is entered.
func (s *BasePlSqlParserListener) EnterCreate_package(ctx *Create_packageContext) {}

// ExitCreate_package is called when production create_package is exited.
func (s *BasePlSqlParserListener) ExitCreate_package(ctx *Create_packageContext) {}

// EnterCreate_package_body is called when production create_package_body is entered.
func (s *BasePlSqlParserListener) EnterCreate_package_body(ctx *Create_package_bodyContext) {}

// ExitCreate_package_body is called when production create_package_body is exited.
func (s *BasePlSqlParserListener) ExitCreate_package_body(ctx *Create_package_bodyContext) {}

// EnterPackage_obj_spec is called when production package_obj_spec is entered.
func (s *BasePlSqlParserListener) EnterPackage_obj_spec(ctx *Package_obj_specContext) {}

// ExitPackage_obj_spec is called when production package_obj_spec is exited.
func (s *BasePlSqlParserListener) ExitPackage_obj_spec(ctx *Package_obj_specContext) {}

// EnterProcedure_spec is called when production procedure_spec is entered.
func (s *BasePlSqlParserListener) EnterProcedure_spec(ctx *Procedure_specContext) {}

// ExitProcedure_spec is called when production procedure_spec is exited.
func (s *BasePlSqlParserListener) ExitProcedure_spec(ctx *Procedure_specContext) {}

// EnterFunction_spec is called when production function_spec is entered.
func (s *BasePlSqlParserListener) EnterFunction_spec(ctx *Function_specContext) {}

// ExitFunction_spec is called when production function_spec is exited.
func (s *BasePlSqlParserListener) ExitFunction_spec(ctx *Function_specContext) {}

// EnterPackage_obj_body is called when production package_obj_body is entered.
func (s *BasePlSqlParserListener) EnterPackage_obj_body(ctx *Package_obj_bodyContext) {}

// ExitPackage_obj_body is called when production package_obj_body is exited.
func (s *BasePlSqlParserListener) ExitPackage_obj_body(ctx *Package_obj_bodyContext) {}

// EnterDrop_procedure is called when production drop_procedure is entered.
func (s *BasePlSqlParserListener) EnterDrop_procedure(ctx *Drop_procedureContext) {}

// ExitDrop_procedure is called when production drop_procedure is exited.
func (s *BasePlSqlParserListener) ExitDrop_procedure(ctx *Drop_procedureContext) {}

// EnterAlter_procedure is called when production alter_procedure is entered.
func (s *BasePlSqlParserListener) EnterAlter_procedure(ctx *Alter_procedureContext) {}

// ExitAlter_procedure is called when production alter_procedure is exited.
func (s *BasePlSqlParserListener) ExitAlter_procedure(ctx *Alter_procedureContext) {}

// EnterFunction_body is called when production function_body is entered.
func (s *BasePlSqlParserListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BasePlSqlParserListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterProcedure_body is called when production procedure_body is entered.
func (s *BasePlSqlParserListener) EnterProcedure_body(ctx *Procedure_bodyContext) {}

// ExitProcedure_body is called when production procedure_body is exited.
func (s *BasePlSqlParserListener) ExitProcedure_body(ctx *Procedure_bodyContext) {}

// EnterCreate_procedure_body is called when production create_procedure_body is entered.
func (s *BasePlSqlParserListener) EnterCreate_procedure_body(ctx *Create_procedure_bodyContext) {}

// ExitCreate_procedure_body is called when production create_procedure_body is exited.
func (s *BasePlSqlParserListener) ExitCreate_procedure_body(ctx *Create_procedure_bodyContext) {}

// EnterDrop_trigger is called when production drop_trigger is entered.
func (s *BasePlSqlParserListener) EnterDrop_trigger(ctx *Drop_triggerContext) {}

// ExitDrop_trigger is called when production drop_trigger is exited.
func (s *BasePlSqlParserListener) ExitDrop_trigger(ctx *Drop_triggerContext) {}

// EnterAlter_trigger is called when production alter_trigger is entered.
func (s *BasePlSqlParserListener) EnterAlter_trigger(ctx *Alter_triggerContext) {}

// ExitAlter_trigger is called when production alter_trigger is exited.
func (s *BasePlSqlParserListener) ExitAlter_trigger(ctx *Alter_triggerContext) {}

// EnterCreate_trigger is called when production create_trigger is entered.
func (s *BasePlSqlParserListener) EnterCreate_trigger(ctx *Create_triggerContext) {}

// ExitCreate_trigger is called when production create_trigger is exited.
func (s *BasePlSqlParserListener) ExitCreate_trigger(ctx *Create_triggerContext) {}

// EnterTrigger_follows_clause is called when production trigger_follows_clause is entered.
func (s *BasePlSqlParserListener) EnterTrigger_follows_clause(ctx *Trigger_follows_clauseContext) {}

// ExitTrigger_follows_clause is called when production trigger_follows_clause is exited.
func (s *BasePlSqlParserListener) ExitTrigger_follows_clause(ctx *Trigger_follows_clauseContext) {}

// EnterTrigger_when_clause is called when production trigger_when_clause is entered.
func (s *BasePlSqlParserListener) EnterTrigger_when_clause(ctx *Trigger_when_clauseContext) {}

// ExitTrigger_when_clause is called when production trigger_when_clause is exited.
func (s *BasePlSqlParserListener) ExitTrigger_when_clause(ctx *Trigger_when_clauseContext) {}

// EnterSimple_dml_trigger is called when production simple_dml_trigger is entered.
func (s *BasePlSqlParserListener) EnterSimple_dml_trigger(ctx *Simple_dml_triggerContext) {}

// ExitSimple_dml_trigger is called when production simple_dml_trigger is exited.
func (s *BasePlSqlParserListener) ExitSimple_dml_trigger(ctx *Simple_dml_triggerContext) {}

// EnterFor_each_row is called when production for_each_row is entered.
func (s *BasePlSqlParserListener) EnterFor_each_row(ctx *For_each_rowContext) {}

// ExitFor_each_row is called when production for_each_row is exited.
func (s *BasePlSqlParserListener) ExitFor_each_row(ctx *For_each_rowContext) {}

// EnterCompound_dml_trigger is called when production compound_dml_trigger is entered.
func (s *BasePlSqlParserListener) EnterCompound_dml_trigger(ctx *Compound_dml_triggerContext) {}

// ExitCompound_dml_trigger is called when production compound_dml_trigger is exited.
func (s *BasePlSqlParserListener) ExitCompound_dml_trigger(ctx *Compound_dml_triggerContext) {}

// EnterNon_dml_trigger is called when production non_dml_trigger is entered.
func (s *BasePlSqlParserListener) EnterNon_dml_trigger(ctx *Non_dml_triggerContext) {}

// ExitNon_dml_trigger is called when production non_dml_trigger is exited.
func (s *BasePlSqlParserListener) ExitNon_dml_trigger(ctx *Non_dml_triggerContext) {}

// EnterTrigger_body is called when production trigger_body is entered.
func (s *BasePlSqlParserListener) EnterTrigger_body(ctx *Trigger_bodyContext) {}

// ExitTrigger_body is called when production trigger_body is exited.
func (s *BasePlSqlParserListener) ExitTrigger_body(ctx *Trigger_bodyContext) {}

// EnterRoutine_clause is called when production routine_clause is entered.
func (s *BasePlSqlParserListener) EnterRoutine_clause(ctx *Routine_clauseContext) {}

// ExitRoutine_clause is called when production routine_clause is exited.
func (s *BasePlSqlParserListener) ExitRoutine_clause(ctx *Routine_clauseContext) {}

// EnterCompound_trigger_block is called when production compound_trigger_block is entered.
func (s *BasePlSqlParserListener) EnterCompound_trigger_block(ctx *Compound_trigger_blockContext) {}

// ExitCompound_trigger_block is called when production compound_trigger_block is exited.
func (s *BasePlSqlParserListener) ExitCompound_trigger_block(ctx *Compound_trigger_blockContext) {}

// EnterTiming_point_section is called when production timing_point_section is entered.
func (s *BasePlSqlParserListener) EnterTiming_point_section(ctx *Timing_point_sectionContext) {}

// ExitTiming_point_section is called when production timing_point_section is exited.
func (s *BasePlSqlParserListener) ExitTiming_point_section(ctx *Timing_point_sectionContext) {}

// EnterNon_dml_event is called when production non_dml_event is entered.
func (s *BasePlSqlParserListener) EnterNon_dml_event(ctx *Non_dml_eventContext) {}

// ExitNon_dml_event is called when production non_dml_event is exited.
func (s *BasePlSqlParserListener) ExitNon_dml_event(ctx *Non_dml_eventContext) {}

// EnterDml_event_clause is called when production dml_event_clause is entered.
func (s *BasePlSqlParserListener) EnterDml_event_clause(ctx *Dml_event_clauseContext) {}

// ExitDml_event_clause is called when production dml_event_clause is exited.
func (s *BasePlSqlParserListener) ExitDml_event_clause(ctx *Dml_event_clauseContext) {}

// EnterDml_event_element is called when production dml_event_element is entered.
func (s *BasePlSqlParserListener) EnterDml_event_element(ctx *Dml_event_elementContext) {}

// ExitDml_event_element is called when production dml_event_element is exited.
func (s *BasePlSqlParserListener) ExitDml_event_element(ctx *Dml_event_elementContext) {}

// EnterDml_event_nested_clause is called when production dml_event_nested_clause is entered.
func (s *BasePlSqlParserListener) EnterDml_event_nested_clause(ctx *Dml_event_nested_clauseContext) {}

// ExitDml_event_nested_clause is called when production dml_event_nested_clause is exited.
func (s *BasePlSqlParserListener) ExitDml_event_nested_clause(ctx *Dml_event_nested_clauseContext) {}

// EnterReferencing_clause is called when production referencing_clause is entered.
func (s *BasePlSqlParserListener) EnterReferencing_clause(ctx *Referencing_clauseContext) {}

// ExitReferencing_clause is called when production referencing_clause is exited.
func (s *BasePlSqlParserListener) ExitReferencing_clause(ctx *Referencing_clauseContext) {}

// EnterReferencing_element is called when production referencing_element is entered.
func (s *BasePlSqlParserListener) EnterReferencing_element(ctx *Referencing_elementContext) {}

// ExitReferencing_element is called when production referencing_element is exited.
func (s *BasePlSqlParserListener) ExitReferencing_element(ctx *Referencing_elementContext) {}

// EnterDrop_type is called when production drop_type is entered.
func (s *BasePlSqlParserListener) EnterDrop_type(ctx *Drop_typeContext) {}

// ExitDrop_type is called when production drop_type is exited.
func (s *BasePlSqlParserListener) ExitDrop_type(ctx *Drop_typeContext) {}

// EnterAlter_type is called when production alter_type is entered.
func (s *BasePlSqlParserListener) EnterAlter_type(ctx *Alter_typeContext) {}

// ExitAlter_type is called when production alter_type is exited.
func (s *BasePlSqlParserListener) ExitAlter_type(ctx *Alter_typeContext) {}

// EnterCompile_type_clause is called when production compile_type_clause is entered.
func (s *BasePlSqlParserListener) EnterCompile_type_clause(ctx *Compile_type_clauseContext) {}

// ExitCompile_type_clause is called when production compile_type_clause is exited.
func (s *BasePlSqlParserListener) ExitCompile_type_clause(ctx *Compile_type_clauseContext) {}

// EnterReplace_type_clause is called when production replace_type_clause is entered.
func (s *BasePlSqlParserListener) EnterReplace_type_clause(ctx *Replace_type_clauseContext) {}

// ExitReplace_type_clause is called when production replace_type_clause is exited.
func (s *BasePlSqlParserListener) ExitReplace_type_clause(ctx *Replace_type_clauseContext) {}

// EnterAlter_method_spec is called when production alter_method_spec is entered.
func (s *BasePlSqlParserListener) EnterAlter_method_spec(ctx *Alter_method_specContext) {}

// ExitAlter_method_spec is called when production alter_method_spec is exited.
func (s *BasePlSqlParserListener) ExitAlter_method_spec(ctx *Alter_method_specContext) {}

// EnterAlter_method_element is called when production alter_method_element is entered.
func (s *BasePlSqlParserListener) EnterAlter_method_element(ctx *Alter_method_elementContext) {}

// ExitAlter_method_element is called when production alter_method_element is exited.
func (s *BasePlSqlParserListener) ExitAlter_method_element(ctx *Alter_method_elementContext) {}

// EnterAlter_attribute_definition is called when production alter_attribute_definition is entered.
func (s *BasePlSqlParserListener) EnterAlter_attribute_definition(ctx *Alter_attribute_definitionContext) {
}

// ExitAlter_attribute_definition is called when production alter_attribute_definition is exited.
func (s *BasePlSqlParserListener) ExitAlter_attribute_definition(ctx *Alter_attribute_definitionContext) {
}

// EnterAttribute_definition is called when production attribute_definition is entered.
func (s *BasePlSqlParserListener) EnterAttribute_definition(ctx *Attribute_definitionContext) {}

// ExitAttribute_definition is called when production attribute_definition is exited.
func (s *BasePlSqlParserListener) ExitAttribute_definition(ctx *Attribute_definitionContext) {}

// EnterAlter_collection_clauses is called when production alter_collection_clauses is entered.
func (s *BasePlSqlParserListener) EnterAlter_collection_clauses(ctx *Alter_collection_clausesContext) {
}

// ExitAlter_collection_clauses is called when production alter_collection_clauses is exited.
func (s *BasePlSqlParserListener) ExitAlter_collection_clauses(ctx *Alter_collection_clausesContext) {}

// EnterDependent_handling_clause is called when production dependent_handling_clause is entered.
func (s *BasePlSqlParserListener) EnterDependent_handling_clause(ctx *Dependent_handling_clauseContext) {
}

// ExitDependent_handling_clause is called when production dependent_handling_clause is exited.
func (s *BasePlSqlParserListener) ExitDependent_handling_clause(ctx *Dependent_handling_clauseContext) {
}

// EnterDependent_exceptions_part is called when production dependent_exceptions_part is entered.
func (s *BasePlSqlParserListener) EnterDependent_exceptions_part(ctx *Dependent_exceptions_partContext) {
}

// ExitDependent_exceptions_part is called when production dependent_exceptions_part is exited.
func (s *BasePlSqlParserListener) ExitDependent_exceptions_part(ctx *Dependent_exceptions_partContext) {
}

// EnterCreate_type is called when production create_type is entered.
func (s *BasePlSqlParserListener) EnterCreate_type(ctx *Create_typeContext) {}

// ExitCreate_type is called when production create_type is exited.
func (s *BasePlSqlParserListener) ExitCreate_type(ctx *Create_typeContext) {}

// EnterType_definition is called when production type_definition is entered.
func (s *BasePlSqlParserListener) EnterType_definition(ctx *Type_definitionContext) {}

// ExitType_definition is called when production type_definition is exited.
func (s *BasePlSqlParserListener) ExitType_definition(ctx *Type_definitionContext) {}

// EnterObject_type_def is called when production object_type_def is entered.
func (s *BasePlSqlParserListener) EnterObject_type_def(ctx *Object_type_defContext) {}

// ExitObject_type_def is called when production object_type_def is exited.
func (s *BasePlSqlParserListener) ExitObject_type_def(ctx *Object_type_defContext) {}

// EnterObject_as_part is called when production object_as_part is entered.
func (s *BasePlSqlParserListener) EnterObject_as_part(ctx *Object_as_partContext) {}

// ExitObject_as_part is called when production object_as_part is exited.
func (s *BasePlSqlParserListener) ExitObject_as_part(ctx *Object_as_partContext) {}

// EnterObject_under_part is called when production object_under_part is entered.
func (s *BasePlSqlParserListener) EnterObject_under_part(ctx *Object_under_partContext) {}

// ExitObject_under_part is called when production object_under_part is exited.
func (s *BasePlSqlParserListener) ExitObject_under_part(ctx *Object_under_partContext) {}

// EnterNested_table_type_def is called when production nested_table_type_def is entered.
func (s *BasePlSqlParserListener) EnterNested_table_type_def(ctx *Nested_table_type_defContext) {}

// ExitNested_table_type_def is called when production nested_table_type_def is exited.
func (s *BasePlSqlParserListener) ExitNested_table_type_def(ctx *Nested_table_type_defContext) {}

// EnterSqlj_object_type is called when production sqlj_object_type is entered.
func (s *BasePlSqlParserListener) EnterSqlj_object_type(ctx *Sqlj_object_typeContext) {}

// ExitSqlj_object_type is called when production sqlj_object_type is exited.
func (s *BasePlSqlParserListener) ExitSqlj_object_type(ctx *Sqlj_object_typeContext) {}

// EnterType_body is called when production type_body is entered.
func (s *BasePlSqlParserListener) EnterType_body(ctx *Type_bodyContext) {}

// ExitType_body is called when production type_body is exited.
func (s *BasePlSqlParserListener) ExitType_body(ctx *Type_bodyContext) {}

// EnterType_body_elements is called when production type_body_elements is entered.
func (s *BasePlSqlParserListener) EnterType_body_elements(ctx *Type_body_elementsContext) {}

// ExitType_body_elements is called when production type_body_elements is exited.
func (s *BasePlSqlParserListener) ExitType_body_elements(ctx *Type_body_elementsContext) {}

// EnterMap_order_func_declaration is called when production map_order_func_declaration is entered.
func (s *BasePlSqlParserListener) EnterMap_order_func_declaration(ctx *Map_order_func_declarationContext) {
}

// ExitMap_order_func_declaration is called when production map_order_func_declaration is exited.
func (s *BasePlSqlParserListener) ExitMap_order_func_declaration(ctx *Map_order_func_declarationContext) {
}

// EnterSubprog_decl_in_type is called when production subprog_decl_in_type is entered.
func (s *BasePlSqlParserListener) EnterSubprog_decl_in_type(ctx *Subprog_decl_in_typeContext) {}

// ExitSubprog_decl_in_type is called when production subprog_decl_in_type is exited.
func (s *BasePlSqlParserListener) ExitSubprog_decl_in_type(ctx *Subprog_decl_in_typeContext) {}

// EnterProc_decl_in_type is called when production proc_decl_in_type is entered.
func (s *BasePlSqlParserListener) EnterProc_decl_in_type(ctx *Proc_decl_in_typeContext) {}

// ExitProc_decl_in_type is called when production proc_decl_in_type is exited.
func (s *BasePlSqlParserListener) ExitProc_decl_in_type(ctx *Proc_decl_in_typeContext) {}

// EnterFunc_decl_in_type is called when production func_decl_in_type is entered.
func (s *BasePlSqlParserListener) EnterFunc_decl_in_type(ctx *Func_decl_in_typeContext) {}

// ExitFunc_decl_in_type is called when production func_decl_in_type is exited.
func (s *BasePlSqlParserListener) ExitFunc_decl_in_type(ctx *Func_decl_in_typeContext) {}

// EnterConstructor_declaration is called when production constructor_declaration is entered.
func (s *BasePlSqlParserListener) EnterConstructor_declaration(ctx *Constructor_declarationContext) {}

// ExitConstructor_declaration is called when production constructor_declaration is exited.
func (s *BasePlSqlParserListener) ExitConstructor_declaration(ctx *Constructor_declarationContext) {}

// EnterModifier_clause is called when production modifier_clause is entered.
func (s *BasePlSqlParserListener) EnterModifier_clause(ctx *Modifier_clauseContext) {}

// ExitModifier_clause is called when production modifier_clause is exited.
func (s *BasePlSqlParserListener) ExitModifier_clause(ctx *Modifier_clauseContext) {}

// EnterObject_member_spec is called when production object_member_spec is entered.
func (s *BasePlSqlParserListener) EnterObject_member_spec(ctx *Object_member_specContext) {}

// ExitObject_member_spec is called when production object_member_spec is exited.
func (s *BasePlSqlParserListener) ExitObject_member_spec(ctx *Object_member_specContext) {}

// EnterSqlj_object_type_attr is called when production sqlj_object_type_attr is entered.
func (s *BasePlSqlParserListener) EnterSqlj_object_type_attr(ctx *Sqlj_object_type_attrContext) {}

// ExitSqlj_object_type_attr is called when production sqlj_object_type_attr is exited.
func (s *BasePlSqlParserListener) ExitSqlj_object_type_attr(ctx *Sqlj_object_type_attrContext) {}

// EnterElement_spec is called when production element_spec is entered.
func (s *BasePlSqlParserListener) EnterElement_spec(ctx *Element_specContext) {}

// ExitElement_spec is called when production element_spec is exited.
func (s *BasePlSqlParserListener) ExitElement_spec(ctx *Element_specContext) {}

// EnterElement_spec_options is called when production element_spec_options is entered.
func (s *BasePlSqlParserListener) EnterElement_spec_options(ctx *Element_spec_optionsContext) {}

// ExitElement_spec_options is called when production element_spec_options is exited.
func (s *BasePlSqlParserListener) ExitElement_spec_options(ctx *Element_spec_optionsContext) {}

// EnterSubprogram_spec is called when production subprogram_spec is entered.
func (s *BasePlSqlParserListener) EnterSubprogram_spec(ctx *Subprogram_specContext) {}

// ExitSubprogram_spec is called when production subprogram_spec is exited.
func (s *BasePlSqlParserListener) ExitSubprogram_spec(ctx *Subprogram_specContext) {}

// EnterType_procedure_spec is called when production type_procedure_spec is entered.
func (s *BasePlSqlParserListener) EnterType_procedure_spec(ctx *Type_procedure_specContext) {}

// ExitType_procedure_spec is called when production type_procedure_spec is exited.
func (s *BasePlSqlParserListener) ExitType_procedure_spec(ctx *Type_procedure_specContext) {}

// EnterType_function_spec is called when production type_function_spec is entered.
func (s *BasePlSqlParserListener) EnterType_function_spec(ctx *Type_function_specContext) {}

// ExitType_function_spec is called when production type_function_spec is exited.
func (s *BasePlSqlParserListener) ExitType_function_spec(ctx *Type_function_specContext) {}

// EnterConstructor_spec is called when production constructor_spec is entered.
func (s *BasePlSqlParserListener) EnterConstructor_spec(ctx *Constructor_specContext) {}

// ExitConstructor_spec is called when production constructor_spec is exited.
func (s *BasePlSqlParserListener) ExitConstructor_spec(ctx *Constructor_specContext) {}

// EnterMap_order_function_spec is called when production map_order_function_spec is entered.
func (s *BasePlSqlParserListener) EnterMap_order_function_spec(ctx *Map_order_function_specContext) {}

// ExitMap_order_function_spec is called when production map_order_function_spec is exited.
func (s *BasePlSqlParserListener) ExitMap_order_function_spec(ctx *Map_order_function_specContext) {}

// EnterPragma_clause is called when production pragma_clause is entered.
func (s *BasePlSqlParserListener) EnterPragma_clause(ctx *Pragma_clauseContext) {}

// ExitPragma_clause is called when production pragma_clause is exited.
func (s *BasePlSqlParserListener) ExitPragma_clause(ctx *Pragma_clauseContext) {}

// EnterPragma_elements is called when production pragma_elements is entered.
func (s *BasePlSqlParserListener) EnterPragma_elements(ctx *Pragma_elementsContext) {}

// ExitPragma_elements is called when production pragma_elements is exited.
func (s *BasePlSqlParserListener) ExitPragma_elements(ctx *Pragma_elementsContext) {}

// EnterType_elements_parameter is called when production type_elements_parameter is entered.
func (s *BasePlSqlParserListener) EnterType_elements_parameter(ctx *Type_elements_parameterContext) {}

// ExitType_elements_parameter is called when production type_elements_parameter is exited.
func (s *BasePlSqlParserListener) ExitType_elements_parameter(ctx *Type_elements_parameterContext) {}

// EnterDrop_sequence is called when production drop_sequence is entered.
func (s *BasePlSqlParserListener) EnterDrop_sequence(ctx *Drop_sequenceContext) {}

// ExitDrop_sequence is called when production drop_sequence is exited.
func (s *BasePlSqlParserListener) ExitDrop_sequence(ctx *Drop_sequenceContext) {}

// EnterAlter_sequence is called when production alter_sequence is entered.
func (s *BasePlSqlParserListener) EnterAlter_sequence(ctx *Alter_sequenceContext) {}

// ExitAlter_sequence is called when production alter_sequence is exited.
func (s *BasePlSqlParserListener) ExitAlter_sequence(ctx *Alter_sequenceContext) {}

// EnterCreate_sequence is called when production create_sequence is entered.
func (s *BasePlSqlParserListener) EnterCreate_sequence(ctx *Create_sequenceContext) {}

// ExitCreate_sequence is called when production create_sequence is exited.
func (s *BasePlSqlParserListener) ExitCreate_sequence(ctx *Create_sequenceContext) {}

// EnterSequence_spec is called when production sequence_spec is entered.
func (s *BasePlSqlParserListener) EnterSequence_spec(ctx *Sequence_specContext) {}

// ExitSequence_spec is called when production sequence_spec is exited.
func (s *BasePlSqlParserListener) ExitSequence_spec(ctx *Sequence_specContext) {}

// EnterSequence_start_clause is called when production sequence_start_clause is entered.
func (s *BasePlSqlParserListener) EnterSequence_start_clause(ctx *Sequence_start_clauseContext) {}

// ExitSequence_start_clause is called when production sequence_start_clause is exited.
func (s *BasePlSqlParserListener) ExitSequence_start_clause(ctx *Sequence_start_clauseContext) {}

// EnterCreate_index is called when production create_index is entered.
func (s *BasePlSqlParserListener) EnterCreate_index(ctx *Create_indexContext) {}

// ExitCreate_index is called when production create_index is exited.
func (s *BasePlSqlParserListener) ExitCreate_index(ctx *Create_indexContext) {}

// EnterAlter_index is called when production alter_index is entered.
func (s *BasePlSqlParserListener) EnterAlter_index(ctx *Alter_indexContext) {}

// ExitAlter_index is called when production alter_index is exited.
func (s *BasePlSqlParserListener) ExitAlter_index(ctx *Alter_indexContext) {}

// EnterDrop_index is called when production drop_index is entered.
func (s *BasePlSqlParserListener) EnterDrop_index(ctx *Drop_indexContext) {}

// ExitDrop_index is called when production drop_index is exited.
func (s *BasePlSqlParserListener) ExitDrop_index(ctx *Drop_indexContext) {}

// EnterCreate_table is called when production create_table is entered.
func (s *BasePlSqlParserListener) EnterCreate_table(ctx *Create_tableContext) {}

// ExitCreate_table is called when production create_table is exited.
func (s *BasePlSqlParserListener) ExitCreate_table(ctx *Create_tableContext) {}

// EnterSize_clause is called when production size_clause is entered.
func (s *BasePlSqlParserListener) EnterSize_clause(ctx *Size_clauseContext) {}

// ExitSize_clause is called when production size_clause is exited.
func (s *BasePlSqlParserListener) ExitSize_clause(ctx *Size_clauseContext) {}

// EnterDrop_table is called when production drop_table is entered.
func (s *BasePlSqlParserListener) EnterDrop_table(ctx *Drop_tableContext) {}

// ExitDrop_table is called when production drop_table is exited.
func (s *BasePlSqlParserListener) ExitDrop_table(ctx *Drop_tableContext) {}

// EnterComment_on_column is called when production comment_on_column is entered.
func (s *BasePlSqlParserListener) EnterComment_on_column(ctx *Comment_on_columnContext) {}

// ExitComment_on_column is called when production comment_on_column is exited.
func (s *BasePlSqlParserListener) ExitComment_on_column(ctx *Comment_on_columnContext) {}

// EnterCreate_synonym is called when production create_synonym is entered.
func (s *BasePlSqlParserListener) EnterCreate_synonym(ctx *Create_synonymContext) {}

// ExitCreate_synonym is called when production create_synonym is exited.
func (s *BasePlSqlParserListener) ExitCreate_synonym(ctx *Create_synonymContext) {}

// EnterComment_on_table is called when production comment_on_table is entered.
func (s *BasePlSqlParserListener) EnterComment_on_table(ctx *Comment_on_tableContext) {}

// ExitComment_on_table is called when production comment_on_table is exited.
func (s *BasePlSqlParserListener) ExitComment_on_table(ctx *Comment_on_tableContext) {}

// EnterAlter_table is called when production alter_table is entered.
func (s *BasePlSqlParserListener) EnterAlter_table(ctx *Alter_tableContext) {}

// ExitAlter_table is called when production alter_table is exited.
func (s *BasePlSqlParserListener) ExitAlter_table(ctx *Alter_tableContext) {}

// EnterAdd_constraint is called when production add_constraint is entered.
func (s *BasePlSqlParserListener) EnterAdd_constraint(ctx *Add_constraintContext) {}

// ExitAdd_constraint is called when production add_constraint is exited.
func (s *BasePlSqlParserListener) ExitAdd_constraint(ctx *Add_constraintContext) {}

// EnterCheck_constraint is called when production check_constraint is entered.
func (s *BasePlSqlParserListener) EnterCheck_constraint(ctx *Check_constraintContext) {}

// ExitCheck_constraint is called when production check_constraint is exited.
func (s *BasePlSqlParserListener) ExitCheck_constraint(ctx *Check_constraintContext) {}

// EnterDrop_constraint is called when production drop_constraint is entered.
func (s *BasePlSqlParserListener) EnterDrop_constraint(ctx *Drop_constraintContext) {}

// ExitDrop_constraint is called when production drop_constraint is exited.
func (s *BasePlSqlParserListener) ExitDrop_constraint(ctx *Drop_constraintContext) {}

// EnterEnable_constraint is called when production enable_constraint is entered.
func (s *BasePlSqlParserListener) EnterEnable_constraint(ctx *Enable_constraintContext) {}

// ExitEnable_constraint is called when production enable_constraint is exited.
func (s *BasePlSqlParserListener) ExitEnable_constraint(ctx *Enable_constraintContext) {}

// EnterDisable_constraint is called when production disable_constraint is entered.
func (s *BasePlSqlParserListener) EnterDisable_constraint(ctx *Disable_constraintContext) {}

// ExitDisable_constraint is called when production disable_constraint is exited.
func (s *BasePlSqlParserListener) ExitDisable_constraint(ctx *Disable_constraintContext) {}

// EnterForeign_key_clause is called when production foreign_key_clause is entered.
func (s *BasePlSqlParserListener) EnterForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// ExitForeign_key_clause is called when production foreign_key_clause is exited.
func (s *BasePlSqlParserListener) ExitForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// EnterReferences_clause is called when production references_clause is entered.
func (s *BasePlSqlParserListener) EnterReferences_clause(ctx *References_clauseContext) {}

// ExitReferences_clause is called when production references_clause is exited.
func (s *BasePlSqlParserListener) ExitReferences_clause(ctx *References_clauseContext) {}

// EnterOn_delete_clause is called when production on_delete_clause is entered.
func (s *BasePlSqlParserListener) EnterOn_delete_clause(ctx *On_delete_clauseContext) {}

// ExitOn_delete_clause is called when production on_delete_clause is exited.
func (s *BasePlSqlParserListener) ExitOn_delete_clause(ctx *On_delete_clauseContext) {}

// EnterUnique_key_clause is called when production unique_key_clause is entered.
func (s *BasePlSqlParserListener) EnterUnique_key_clause(ctx *Unique_key_clauseContext) {}

// ExitUnique_key_clause is called when production unique_key_clause is exited.
func (s *BasePlSqlParserListener) ExitUnique_key_clause(ctx *Unique_key_clauseContext) {}

// EnterPrimary_key_clause is called when production primary_key_clause is entered.
func (s *BasePlSqlParserListener) EnterPrimary_key_clause(ctx *Primary_key_clauseContext) {}

// ExitPrimary_key_clause is called when production primary_key_clause is exited.
func (s *BasePlSqlParserListener) ExitPrimary_key_clause(ctx *Primary_key_clauseContext) {}

// EnterAnonymous_block is called when production anonymous_block is entered.
func (s *BasePlSqlParserListener) EnterAnonymous_block(ctx *Anonymous_blockContext) {}

// ExitAnonymous_block is called when production anonymous_block is exited.
func (s *BasePlSqlParserListener) ExitAnonymous_block(ctx *Anonymous_blockContext) {}

// EnterInvoker_rights_clause is called when production invoker_rights_clause is entered.
func (s *BasePlSqlParserListener) EnterInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// ExitInvoker_rights_clause is called when production invoker_rights_clause is exited.
func (s *BasePlSqlParserListener) ExitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// EnterCompiler_parameters_clause is called when production compiler_parameters_clause is entered.
func (s *BasePlSqlParserListener) EnterCompiler_parameters_clause(ctx *Compiler_parameters_clauseContext) {
}

// ExitCompiler_parameters_clause is called when production compiler_parameters_clause is exited.
func (s *BasePlSqlParserListener) ExitCompiler_parameters_clause(ctx *Compiler_parameters_clauseContext) {
}

// EnterCall_spec is called when production call_spec is entered.
func (s *BasePlSqlParserListener) EnterCall_spec(ctx *Call_specContext) {}

// ExitCall_spec is called when production call_spec is exited.
func (s *BasePlSqlParserListener) ExitCall_spec(ctx *Call_specContext) {}

// EnterJava_spec is called when production java_spec is entered.
func (s *BasePlSqlParserListener) EnterJava_spec(ctx *Java_specContext) {}

// ExitJava_spec is called when production java_spec is exited.
func (s *BasePlSqlParserListener) ExitJava_spec(ctx *Java_specContext) {}

// EnterC_spec is called when production c_spec is entered.
func (s *BasePlSqlParserListener) EnterC_spec(ctx *C_specContext) {}

// ExitC_spec is called when production c_spec is exited.
func (s *BasePlSqlParserListener) ExitC_spec(ctx *C_specContext) {}

// EnterC_agent_in_clause is called when production c_agent_in_clause is entered.
func (s *BasePlSqlParserListener) EnterC_agent_in_clause(ctx *C_agent_in_clauseContext) {}

// ExitC_agent_in_clause is called when production c_agent_in_clause is exited.
func (s *BasePlSqlParserListener) ExitC_agent_in_clause(ctx *C_agent_in_clauseContext) {}

// EnterC_parameters_clause is called when production c_parameters_clause is entered.
func (s *BasePlSqlParserListener) EnterC_parameters_clause(ctx *C_parameters_clauseContext) {}

// ExitC_parameters_clause is called when production c_parameters_clause is exited.
func (s *BasePlSqlParserListener) ExitC_parameters_clause(ctx *C_parameters_clauseContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePlSqlParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePlSqlParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterDefault_value_part is called when production default_value_part is entered.
func (s *BasePlSqlParserListener) EnterDefault_value_part(ctx *Default_value_partContext) {}

// ExitDefault_value_part is called when production default_value_part is exited.
func (s *BasePlSqlParserListener) ExitDefault_value_part(ctx *Default_value_partContext) {}

// EnterDeclare_spec is called when production declare_spec is entered.
func (s *BasePlSqlParserListener) EnterDeclare_spec(ctx *Declare_specContext) {}

// ExitDeclare_spec is called when production declare_spec is exited.
func (s *BasePlSqlParserListener) ExitDeclare_spec(ctx *Declare_specContext) {}

// EnterVariable_declaration is called when production variable_declaration is entered.
func (s *BasePlSqlParserListener) EnterVariable_declaration(ctx *Variable_declarationContext) {}

// ExitVariable_declaration is called when production variable_declaration is exited.
func (s *BasePlSqlParserListener) ExitVariable_declaration(ctx *Variable_declarationContext) {}

// EnterSubtype_declaration is called when production subtype_declaration is entered.
func (s *BasePlSqlParserListener) EnterSubtype_declaration(ctx *Subtype_declarationContext) {}

// ExitSubtype_declaration is called when production subtype_declaration is exited.
func (s *BasePlSqlParserListener) ExitSubtype_declaration(ctx *Subtype_declarationContext) {}

// EnterCursor_declaration is called when production cursor_declaration is entered.
func (s *BasePlSqlParserListener) EnterCursor_declaration(ctx *Cursor_declarationContext) {}

// ExitCursor_declaration is called when production cursor_declaration is exited.
func (s *BasePlSqlParserListener) ExitCursor_declaration(ctx *Cursor_declarationContext) {}

// EnterParameter_spec is called when production parameter_spec is entered.
func (s *BasePlSqlParserListener) EnterParameter_spec(ctx *Parameter_specContext) {}

// ExitParameter_spec is called when production parameter_spec is exited.
func (s *BasePlSqlParserListener) ExitParameter_spec(ctx *Parameter_specContext) {}

// EnterException_declaration is called when production exception_declaration is entered.
func (s *BasePlSqlParserListener) EnterException_declaration(ctx *Exception_declarationContext) {}

// ExitException_declaration is called when production exception_declaration is exited.
func (s *BasePlSqlParserListener) ExitException_declaration(ctx *Exception_declarationContext) {}

// EnterPragma_declaration is called when production pragma_declaration is entered.
func (s *BasePlSqlParserListener) EnterPragma_declaration(ctx *Pragma_declarationContext) {}

// ExitPragma_declaration is called when production pragma_declaration is exited.
func (s *BasePlSqlParserListener) ExitPragma_declaration(ctx *Pragma_declarationContext) {}

// EnterRecord_type_def is called when production record_type_def is entered.
func (s *BasePlSqlParserListener) EnterRecord_type_def(ctx *Record_type_defContext) {}

// ExitRecord_type_def is called when production record_type_def is exited.
func (s *BasePlSqlParserListener) ExitRecord_type_def(ctx *Record_type_defContext) {}

// EnterField_spec is called when production field_spec is entered.
func (s *BasePlSqlParserListener) EnterField_spec(ctx *Field_specContext) {}

// ExitField_spec is called when production field_spec is exited.
func (s *BasePlSqlParserListener) ExitField_spec(ctx *Field_specContext) {}

// EnterRef_cursor_type_def is called when production ref_cursor_type_def is entered.
func (s *BasePlSqlParserListener) EnterRef_cursor_type_def(ctx *Ref_cursor_type_defContext) {}

// ExitRef_cursor_type_def is called when production ref_cursor_type_def is exited.
func (s *BasePlSqlParserListener) ExitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) {}

// EnterType_declaration is called when production type_declaration is entered.
func (s *BasePlSqlParserListener) EnterType_declaration(ctx *Type_declarationContext) {}

// ExitType_declaration is called when production type_declaration is exited.
func (s *BasePlSqlParserListener) ExitType_declaration(ctx *Type_declarationContext) {}

// EnterTable_type_def is called when production table_type_def is entered.
func (s *BasePlSqlParserListener) EnterTable_type_def(ctx *Table_type_defContext) {}

// ExitTable_type_def is called when production table_type_def is exited.
func (s *BasePlSqlParserListener) ExitTable_type_def(ctx *Table_type_defContext) {}

// EnterTable_indexed_by_part is called when production table_indexed_by_part is entered.
func (s *BasePlSqlParserListener) EnterTable_indexed_by_part(ctx *Table_indexed_by_partContext) {}

// ExitTable_indexed_by_part is called when production table_indexed_by_part is exited.
func (s *BasePlSqlParserListener) ExitTable_indexed_by_part(ctx *Table_indexed_by_partContext) {}

// EnterVarray_type_def is called when production varray_type_def is entered.
func (s *BasePlSqlParserListener) EnterVarray_type_def(ctx *Varray_type_defContext) {}

// ExitVarray_type_def is called when production varray_type_def is exited.
func (s *BasePlSqlParserListener) ExitVarray_type_def(ctx *Varray_type_defContext) {}

// EnterSeq_of_statements is called when production seq_of_statements is entered.
func (s *BasePlSqlParserListener) EnterSeq_of_statements(ctx *Seq_of_statementsContext) {}

// ExitSeq_of_statements is called when production seq_of_statements is exited.
func (s *BasePlSqlParserListener) ExitSeq_of_statements(ctx *Seq_of_statementsContext) {}

// EnterLabel_declaration is called when production label_declaration is entered.
func (s *BasePlSqlParserListener) EnterLabel_declaration(ctx *Label_declarationContext) {}

// ExitLabel_declaration is called when production label_declaration is exited.
func (s *BasePlSqlParserListener) ExitLabel_declaration(ctx *Label_declarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePlSqlParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePlSqlParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSwallow_to_semi is called when production swallow_to_semi is entered.
func (s *BasePlSqlParserListener) EnterSwallow_to_semi(ctx *Swallow_to_semiContext) {}

// ExitSwallow_to_semi is called when production swallow_to_semi is exited.
func (s *BasePlSqlParserListener) ExitSwallow_to_semi(ctx *Swallow_to_semiContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BasePlSqlParserListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BasePlSqlParserListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BasePlSqlParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BasePlSqlParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterExit_statement is called when production exit_statement is entered.
func (s *BasePlSqlParserListener) EnterExit_statement(ctx *Exit_statementContext) {}

// ExitExit_statement is called when production exit_statement is exited.
func (s *BasePlSqlParserListener) ExitExit_statement(ctx *Exit_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *BasePlSqlParserListener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *BasePlSqlParserListener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasePlSqlParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasePlSqlParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElsif_part is called when production elsif_part is entered.
func (s *BasePlSqlParserListener) EnterElsif_part(ctx *Elsif_partContext) {}

// ExitElsif_part is called when production elsif_part is exited.
func (s *BasePlSqlParserListener) ExitElsif_part(ctx *Elsif_partContext) {}

// EnterElse_part is called when production else_part is entered.
func (s *BasePlSqlParserListener) EnterElse_part(ctx *Else_partContext) {}

// ExitElse_part is called when production else_part is exited.
func (s *BasePlSqlParserListener) ExitElse_part(ctx *Else_partContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BasePlSqlParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BasePlSqlParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterCursor_loop_param is called when production cursor_loop_param is entered.
func (s *BasePlSqlParserListener) EnterCursor_loop_param(ctx *Cursor_loop_paramContext) {}

// ExitCursor_loop_param is called when production cursor_loop_param is exited.
func (s *BasePlSqlParserListener) ExitCursor_loop_param(ctx *Cursor_loop_paramContext) {}

// EnterForall_statement is called when production forall_statement is entered.
func (s *BasePlSqlParserListener) EnterForall_statement(ctx *Forall_statementContext) {}

// ExitForall_statement is called when production forall_statement is exited.
func (s *BasePlSqlParserListener) ExitForall_statement(ctx *Forall_statementContext) {}

// EnterBounds_clause is called when production bounds_clause is entered.
func (s *BasePlSqlParserListener) EnterBounds_clause(ctx *Bounds_clauseContext) {}

// ExitBounds_clause is called when production bounds_clause is exited.
func (s *BasePlSqlParserListener) ExitBounds_clause(ctx *Bounds_clauseContext) {}

// EnterBetween_bound is called when production between_bound is entered.
func (s *BasePlSqlParserListener) EnterBetween_bound(ctx *Between_boundContext) {}

// ExitBetween_bound is called when production between_bound is exited.
func (s *BasePlSqlParserListener) ExitBetween_bound(ctx *Between_boundContext) {}

// EnterLower_bound is called when production lower_bound is entered.
func (s *BasePlSqlParserListener) EnterLower_bound(ctx *Lower_boundContext) {}

// ExitLower_bound is called when production lower_bound is exited.
func (s *BasePlSqlParserListener) ExitLower_bound(ctx *Lower_boundContext) {}

// EnterUpper_bound is called when production upper_bound is entered.
func (s *BasePlSqlParserListener) EnterUpper_bound(ctx *Upper_boundContext) {}

// ExitUpper_bound is called when production upper_bound is exited.
func (s *BasePlSqlParserListener) ExitUpper_bound(ctx *Upper_boundContext) {}

// EnterNull_statement is called when production null_statement is entered.
func (s *BasePlSqlParserListener) EnterNull_statement(ctx *Null_statementContext) {}

// ExitNull_statement is called when production null_statement is exited.
func (s *BasePlSqlParserListener) ExitNull_statement(ctx *Null_statementContext) {}

// EnterRaise_statement is called when production raise_statement is entered.
func (s *BasePlSqlParserListener) EnterRaise_statement(ctx *Raise_statementContext) {}

// ExitRaise_statement is called when production raise_statement is exited.
func (s *BasePlSqlParserListener) ExitRaise_statement(ctx *Raise_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BasePlSqlParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BasePlSqlParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BasePlSqlParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BasePlSqlParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterPipe_row_statement is called when production pipe_row_statement is entered.
func (s *BasePlSqlParserListener) EnterPipe_row_statement(ctx *Pipe_row_statementContext) {}

// ExitPipe_row_statement is called when production pipe_row_statement is exited.
func (s *BasePlSqlParserListener) ExitPipe_row_statement(ctx *Pipe_row_statementContext) {}

// EnterBody is called when production body is entered.
func (s *BasePlSqlParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasePlSqlParserListener) ExitBody(ctx *BodyContext) {}

// EnterException_handler is called when production exception_handler is entered.
func (s *BasePlSqlParserListener) EnterException_handler(ctx *Exception_handlerContext) {}

// ExitException_handler is called when production exception_handler is exited.
func (s *BasePlSqlParserListener) ExitException_handler(ctx *Exception_handlerContext) {}

// EnterTrigger_block is called when production trigger_block is entered.
func (s *BasePlSqlParserListener) EnterTrigger_block(ctx *Trigger_blockContext) {}

// ExitTrigger_block is called when production trigger_block is exited.
func (s *BasePlSqlParserListener) ExitTrigger_block(ctx *Trigger_blockContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePlSqlParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePlSqlParserListener) ExitBlock(ctx *BlockContext) {}

// EnterSql_statement is called when production sql_statement is entered.
func (s *BasePlSqlParserListener) EnterSql_statement(ctx *Sql_statementContext) {}

// ExitSql_statement is called when production sql_statement is exited.
func (s *BasePlSqlParserListener) ExitSql_statement(ctx *Sql_statementContext) {}

// EnterExecute_immediate is called when production execute_immediate is entered.
func (s *BasePlSqlParserListener) EnterExecute_immediate(ctx *Execute_immediateContext) {}

// ExitExecute_immediate is called when production execute_immediate is exited.
func (s *BasePlSqlParserListener) ExitExecute_immediate(ctx *Execute_immediateContext) {}

// EnterDynamic_returning_clause is called when production dynamic_returning_clause is entered.
func (s *BasePlSqlParserListener) EnterDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) {
}

// ExitDynamic_returning_clause is called when production dynamic_returning_clause is exited.
func (s *BasePlSqlParserListener) ExitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) {}

// EnterData_manipulation_language_statements is called when production data_manipulation_language_statements is entered.
func (s *BasePlSqlParserListener) EnterData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) {
}

// ExitData_manipulation_language_statements is called when production data_manipulation_language_statements is exited.
func (s *BasePlSqlParserListener) ExitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) {
}

// EnterCursor_manipulation_statements is called when production cursor_manipulation_statements is entered.
func (s *BasePlSqlParserListener) EnterCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) {
}

// ExitCursor_manipulation_statements is called when production cursor_manipulation_statements is exited.
func (s *BasePlSqlParserListener) ExitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) {
}

// EnterClose_statement is called when production close_statement is entered.
func (s *BasePlSqlParserListener) EnterClose_statement(ctx *Close_statementContext) {}

// ExitClose_statement is called when production close_statement is exited.
func (s *BasePlSqlParserListener) ExitClose_statement(ctx *Close_statementContext) {}

// EnterOpen_statement is called when production open_statement is entered.
func (s *BasePlSqlParserListener) EnterOpen_statement(ctx *Open_statementContext) {}

// ExitOpen_statement is called when production open_statement is exited.
func (s *BasePlSqlParserListener) ExitOpen_statement(ctx *Open_statementContext) {}

// EnterFetch_statement is called when production fetch_statement is entered.
func (s *BasePlSqlParserListener) EnterFetch_statement(ctx *Fetch_statementContext) {}

// ExitFetch_statement is called when production fetch_statement is exited.
func (s *BasePlSqlParserListener) ExitFetch_statement(ctx *Fetch_statementContext) {}

// EnterOpen_for_statement is called when production open_for_statement is entered.
func (s *BasePlSqlParserListener) EnterOpen_for_statement(ctx *Open_for_statementContext) {}

// ExitOpen_for_statement is called when production open_for_statement is exited.
func (s *BasePlSqlParserListener) ExitOpen_for_statement(ctx *Open_for_statementContext) {}

// EnterTransaction_control_statements is called when production transaction_control_statements is entered.
func (s *BasePlSqlParserListener) EnterTransaction_control_statements(ctx *Transaction_control_statementsContext) {
}

// ExitTransaction_control_statements is called when production transaction_control_statements is exited.
func (s *BasePlSqlParserListener) ExitTransaction_control_statements(ctx *Transaction_control_statementsContext) {
}

// EnterSet_transaction_command is called when production set_transaction_command is entered.
func (s *BasePlSqlParserListener) EnterSet_transaction_command(ctx *Set_transaction_commandContext) {}

// ExitSet_transaction_command is called when production set_transaction_command is exited.
func (s *BasePlSqlParserListener) ExitSet_transaction_command(ctx *Set_transaction_commandContext) {}

// EnterSet_constraint_command is called when production set_constraint_command is entered.
func (s *BasePlSqlParserListener) EnterSet_constraint_command(ctx *Set_constraint_commandContext) {}

// ExitSet_constraint_command is called when production set_constraint_command is exited.
func (s *BasePlSqlParserListener) ExitSet_constraint_command(ctx *Set_constraint_commandContext) {}

// EnterCommit_statement is called when production commit_statement is entered.
func (s *BasePlSqlParserListener) EnterCommit_statement(ctx *Commit_statementContext) {}

// ExitCommit_statement is called when production commit_statement is exited.
func (s *BasePlSqlParserListener) ExitCommit_statement(ctx *Commit_statementContext) {}

// EnterWrite_clause is called when production write_clause is entered.
func (s *BasePlSqlParserListener) EnterWrite_clause(ctx *Write_clauseContext) {}

// ExitWrite_clause is called when production write_clause is exited.
func (s *BasePlSqlParserListener) ExitWrite_clause(ctx *Write_clauseContext) {}

// EnterRollback_statement is called when production rollback_statement is entered.
func (s *BasePlSqlParserListener) EnterRollback_statement(ctx *Rollback_statementContext) {}

// ExitRollback_statement is called when production rollback_statement is exited.
func (s *BasePlSqlParserListener) ExitRollback_statement(ctx *Rollback_statementContext) {}

// EnterSavepoint_statement is called when production savepoint_statement is entered.
func (s *BasePlSqlParserListener) EnterSavepoint_statement(ctx *Savepoint_statementContext) {}

// ExitSavepoint_statement is called when production savepoint_statement is exited.
func (s *BasePlSqlParserListener) ExitSavepoint_statement(ctx *Savepoint_statementContext) {}

// EnterExplain_statement is called when production explain_statement is entered.
func (s *BasePlSqlParserListener) EnterExplain_statement(ctx *Explain_statementContext) {}

// ExitExplain_statement is called when production explain_statement is exited.
func (s *BasePlSqlParserListener) ExitExplain_statement(ctx *Explain_statementContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BasePlSqlParserListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BasePlSqlParserListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterSubquery_factoring_clause is called when production subquery_factoring_clause is entered.
func (s *BasePlSqlParserListener) EnterSubquery_factoring_clause(ctx *Subquery_factoring_clauseContext) {
}

// ExitSubquery_factoring_clause is called when production subquery_factoring_clause is exited.
func (s *BasePlSqlParserListener) ExitSubquery_factoring_clause(ctx *Subquery_factoring_clauseContext) {
}

// EnterFactoring_element is called when production factoring_element is entered.
func (s *BasePlSqlParserListener) EnterFactoring_element(ctx *Factoring_elementContext) {}

// ExitFactoring_element is called when production factoring_element is exited.
func (s *BasePlSqlParserListener) ExitFactoring_element(ctx *Factoring_elementContext) {}

// EnterSearch_clause is called when production search_clause is entered.
func (s *BasePlSqlParserListener) EnterSearch_clause(ctx *Search_clauseContext) {}

// ExitSearch_clause is called when production search_clause is exited.
func (s *BasePlSqlParserListener) ExitSearch_clause(ctx *Search_clauseContext) {}

// EnterCycle_clause is called when production cycle_clause is entered.
func (s *BasePlSqlParserListener) EnterCycle_clause(ctx *Cycle_clauseContext) {}

// ExitCycle_clause is called when production cycle_clause is exited.
func (s *BasePlSqlParserListener) ExitCycle_clause(ctx *Cycle_clauseContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BasePlSqlParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BasePlSqlParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSubquery_basic_elements is called when production subquery_basic_elements is entered.
func (s *BasePlSqlParserListener) EnterSubquery_basic_elements(ctx *Subquery_basic_elementsContext) {}

// ExitSubquery_basic_elements is called when production subquery_basic_elements is exited.
func (s *BasePlSqlParserListener) ExitSubquery_basic_elements(ctx *Subquery_basic_elementsContext) {}

// EnterSubquery_operation_part is called when production subquery_operation_part is entered.
func (s *BasePlSqlParserListener) EnterSubquery_operation_part(ctx *Subquery_operation_partContext) {}

// ExitSubquery_operation_part is called when production subquery_operation_part is exited.
func (s *BasePlSqlParserListener) ExitSubquery_operation_part(ctx *Subquery_operation_partContext) {}

// EnterQuery_block is called when production query_block is entered.
func (s *BasePlSqlParserListener) EnterQuery_block(ctx *Query_blockContext) {}

// ExitQuery_block is called when production query_block is exited.
func (s *BasePlSqlParserListener) ExitQuery_block(ctx *Query_blockContext) {}

// EnterSelected_element is called when production selected_element is entered.
func (s *BasePlSqlParserListener) EnterSelected_element(ctx *Selected_elementContext) {}

// ExitSelected_element is called when production selected_element is exited.
func (s *BasePlSqlParserListener) ExitSelected_element(ctx *Selected_elementContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BasePlSqlParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BasePlSqlParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterSelect_list_elements is called when production select_list_elements is entered.
func (s *BasePlSqlParserListener) EnterSelect_list_elements(ctx *Select_list_elementsContext) {}

// ExitSelect_list_elements is called when production select_list_elements is exited.
func (s *BasePlSqlParserListener) ExitSelect_list_elements(ctx *Select_list_elementsContext) {}

// EnterTable_ref_list is called when production table_ref_list is entered.
func (s *BasePlSqlParserListener) EnterTable_ref_list(ctx *Table_ref_listContext) {}

// ExitTable_ref_list is called when production table_ref_list is exited.
func (s *BasePlSqlParserListener) ExitTable_ref_list(ctx *Table_ref_listContext) {}

// EnterTable_ref is called when production table_ref is entered.
func (s *BasePlSqlParserListener) EnterTable_ref(ctx *Table_refContext) {}

// ExitTable_ref is called when production table_ref is exited.
func (s *BasePlSqlParserListener) ExitTable_ref(ctx *Table_refContext) {}

// EnterTable_ref_aux is called when production table_ref_aux is entered.
func (s *BasePlSqlParserListener) EnterTable_ref_aux(ctx *Table_ref_auxContext) {}

// ExitTable_ref_aux is called when production table_ref_aux is exited.
func (s *BasePlSqlParserListener) ExitTable_ref_aux(ctx *Table_ref_auxContext) {}

// EnterTable_ref_aux_internal_one is called when production table_ref_aux_internal_one is entered.
func (s *BasePlSqlParserListener) EnterTable_ref_aux_internal_one(ctx *Table_ref_aux_internal_oneContext) {
}

// ExitTable_ref_aux_internal_one is called when production table_ref_aux_internal_one is exited.
func (s *BasePlSqlParserListener) ExitTable_ref_aux_internal_one(ctx *Table_ref_aux_internal_oneContext) {
}

// EnterTable_ref_aux_internal_two is called when production table_ref_aux_internal_two is entered.
func (s *BasePlSqlParserListener) EnterTable_ref_aux_internal_two(ctx *Table_ref_aux_internal_twoContext) {
}

// ExitTable_ref_aux_internal_two is called when production table_ref_aux_internal_two is exited.
func (s *BasePlSqlParserListener) ExitTable_ref_aux_internal_two(ctx *Table_ref_aux_internal_twoContext) {
}

// EnterTable_ref_aux_internal_three is called when production table_ref_aux_internal_three is entered.
func (s *BasePlSqlParserListener) EnterTable_ref_aux_internal_three(ctx *Table_ref_aux_internal_threeContext) {
}

// ExitTable_ref_aux_internal_three is called when production table_ref_aux_internal_three is exited.
func (s *BasePlSqlParserListener) ExitTable_ref_aux_internal_three(ctx *Table_ref_aux_internal_threeContext) {
}

// EnterJoin_clause is called when production join_clause is entered.
func (s *BasePlSqlParserListener) EnterJoin_clause(ctx *Join_clauseContext) {}

// ExitJoin_clause is called when production join_clause is exited.
func (s *BasePlSqlParserListener) ExitJoin_clause(ctx *Join_clauseContext) {}

// EnterJoin_on_part is called when production join_on_part is entered.
func (s *BasePlSqlParserListener) EnterJoin_on_part(ctx *Join_on_partContext) {}

// ExitJoin_on_part is called when production join_on_part is exited.
func (s *BasePlSqlParserListener) ExitJoin_on_part(ctx *Join_on_partContext) {}

// EnterJoin_using_part is called when production join_using_part is entered.
func (s *BasePlSqlParserListener) EnterJoin_using_part(ctx *Join_using_partContext) {}

// ExitJoin_using_part is called when production join_using_part is exited.
func (s *BasePlSqlParserListener) ExitJoin_using_part(ctx *Join_using_partContext) {}

// EnterOuter_join_type is called when production outer_join_type is entered.
func (s *BasePlSqlParserListener) EnterOuter_join_type(ctx *Outer_join_typeContext) {}

// ExitOuter_join_type is called when production outer_join_type is exited.
func (s *BasePlSqlParserListener) ExitOuter_join_type(ctx *Outer_join_typeContext) {}

// EnterQuery_partition_clause is called when production query_partition_clause is entered.
func (s *BasePlSqlParserListener) EnterQuery_partition_clause(ctx *Query_partition_clauseContext) {}

// ExitQuery_partition_clause is called when production query_partition_clause is exited.
func (s *BasePlSqlParserListener) ExitQuery_partition_clause(ctx *Query_partition_clauseContext) {}

// EnterFlashback_query_clause is called when production flashback_query_clause is entered.
func (s *BasePlSqlParserListener) EnterFlashback_query_clause(ctx *Flashback_query_clauseContext) {}

// ExitFlashback_query_clause is called when production flashback_query_clause is exited.
func (s *BasePlSqlParserListener) ExitFlashback_query_clause(ctx *Flashback_query_clauseContext) {}

// EnterPivot_clause is called when production pivot_clause is entered.
func (s *BasePlSqlParserListener) EnterPivot_clause(ctx *Pivot_clauseContext) {}

// ExitPivot_clause is called when production pivot_clause is exited.
func (s *BasePlSqlParserListener) ExitPivot_clause(ctx *Pivot_clauseContext) {}

// EnterPivot_element is called when production pivot_element is entered.
func (s *BasePlSqlParserListener) EnterPivot_element(ctx *Pivot_elementContext) {}

// ExitPivot_element is called when production pivot_element is exited.
func (s *BasePlSqlParserListener) ExitPivot_element(ctx *Pivot_elementContext) {}

// EnterPivot_for_clause is called when production pivot_for_clause is entered.
func (s *BasePlSqlParserListener) EnterPivot_for_clause(ctx *Pivot_for_clauseContext) {}

// ExitPivot_for_clause is called when production pivot_for_clause is exited.
func (s *BasePlSqlParserListener) ExitPivot_for_clause(ctx *Pivot_for_clauseContext) {}

// EnterPivot_in_clause is called when production pivot_in_clause is entered.
func (s *BasePlSqlParserListener) EnterPivot_in_clause(ctx *Pivot_in_clauseContext) {}

// ExitPivot_in_clause is called when production pivot_in_clause is exited.
func (s *BasePlSqlParserListener) ExitPivot_in_clause(ctx *Pivot_in_clauseContext) {}

// EnterPivot_in_clause_element is called when production pivot_in_clause_element is entered.
func (s *BasePlSqlParserListener) EnterPivot_in_clause_element(ctx *Pivot_in_clause_elementContext) {}

// ExitPivot_in_clause_element is called when production pivot_in_clause_element is exited.
func (s *BasePlSqlParserListener) ExitPivot_in_clause_element(ctx *Pivot_in_clause_elementContext) {}

// EnterPivot_in_clause_elements is called when production pivot_in_clause_elements is entered.
func (s *BasePlSqlParserListener) EnterPivot_in_clause_elements(ctx *Pivot_in_clause_elementsContext) {
}

// ExitPivot_in_clause_elements is called when production pivot_in_clause_elements is exited.
func (s *BasePlSqlParserListener) ExitPivot_in_clause_elements(ctx *Pivot_in_clause_elementsContext) {}

// EnterUnpivot_clause is called when production unpivot_clause is entered.
func (s *BasePlSqlParserListener) EnterUnpivot_clause(ctx *Unpivot_clauseContext) {}

// ExitUnpivot_clause is called when production unpivot_clause is exited.
func (s *BasePlSqlParserListener) ExitUnpivot_clause(ctx *Unpivot_clauseContext) {}

// EnterUnpivot_in_clause is called when production unpivot_in_clause is entered.
func (s *BasePlSqlParserListener) EnterUnpivot_in_clause(ctx *Unpivot_in_clauseContext) {}

// ExitUnpivot_in_clause is called when production unpivot_in_clause is exited.
func (s *BasePlSqlParserListener) ExitUnpivot_in_clause(ctx *Unpivot_in_clauseContext) {}

// EnterUnpivot_in_elements is called when production unpivot_in_elements is entered.
func (s *BasePlSqlParserListener) EnterUnpivot_in_elements(ctx *Unpivot_in_elementsContext) {}

// ExitUnpivot_in_elements is called when production unpivot_in_elements is exited.
func (s *BasePlSqlParserListener) ExitUnpivot_in_elements(ctx *Unpivot_in_elementsContext) {}

// EnterHierarchical_query_clause is called when production hierarchical_query_clause is entered.
func (s *BasePlSqlParserListener) EnterHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) {
}

// ExitHierarchical_query_clause is called when production hierarchical_query_clause is exited.
func (s *BasePlSqlParserListener) ExitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) {
}

// EnterStart_part is called when production start_part is entered.
func (s *BasePlSqlParserListener) EnterStart_part(ctx *Start_partContext) {}

// ExitStart_part is called when production start_part is exited.
func (s *BasePlSqlParserListener) ExitStart_part(ctx *Start_partContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BasePlSqlParserListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BasePlSqlParserListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterGroup_by_elements is called when production group_by_elements is entered.
func (s *BasePlSqlParserListener) EnterGroup_by_elements(ctx *Group_by_elementsContext) {}

// ExitGroup_by_elements is called when production group_by_elements is exited.
func (s *BasePlSqlParserListener) ExitGroup_by_elements(ctx *Group_by_elementsContext) {}

// EnterRollup_cube_clause is called when production rollup_cube_clause is entered.
func (s *BasePlSqlParserListener) EnterRollup_cube_clause(ctx *Rollup_cube_clauseContext) {}

// ExitRollup_cube_clause is called when production rollup_cube_clause is exited.
func (s *BasePlSqlParserListener) ExitRollup_cube_clause(ctx *Rollup_cube_clauseContext) {}

// EnterGrouping_sets_clause is called when production grouping_sets_clause is entered.
func (s *BasePlSqlParserListener) EnterGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// ExitGrouping_sets_clause is called when production grouping_sets_clause is exited.
func (s *BasePlSqlParserListener) ExitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// EnterGrouping_sets_elements is called when production grouping_sets_elements is entered.
func (s *BasePlSqlParserListener) EnterGrouping_sets_elements(ctx *Grouping_sets_elementsContext) {}

// ExitGrouping_sets_elements is called when production grouping_sets_elements is exited.
func (s *BasePlSqlParserListener) ExitGrouping_sets_elements(ctx *Grouping_sets_elementsContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BasePlSqlParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BasePlSqlParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterModel_clause is called when production model_clause is entered.
func (s *BasePlSqlParserListener) EnterModel_clause(ctx *Model_clauseContext) {}

// ExitModel_clause is called when production model_clause is exited.
func (s *BasePlSqlParserListener) ExitModel_clause(ctx *Model_clauseContext) {}

// EnterCell_reference_options is called when production cell_reference_options is entered.
func (s *BasePlSqlParserListener) EnterCell_reference_options(ctx *Cell_reference_optionsContext) {}

// ExitCell_reference_options is called when production cell_reference_options is exited.
func (s *BasePlSqlParserListener) ExitCell_reference_options(ctx *Cell_reference_optionsContext) {}

// EnterReturn_rows_clause is called when production return_rows_clause is entered.
func (s *BasePlSqlParserListener) EnterReturn_rows_clause(ctx *Return_rows_clauseContext) {}

// ExitReturn_rows_clause is called when production return_rows_clause is exited.
func (s *BasePlSqlParserListener) ExitReturn_rows_clause(ctx *Return_rows_clauseContext) {}

// EnterReference_model is called when production reference_model is entered.
func (s *BasePlSqlParserListener) EnterReference_model(ctx *Reference_modelContext) {}

// ExitReference_model is called when production reference_model is exited.
func (s *BasePlSqlParserListener) ExitReference_model(ctx *Reference_modelContext) {}

// EnterMain_model is called when production main_model is entered.
func (s *BasePlSqlParserListener) EnterMain_model(ctx *Main_modelContext) {}

// ExitMain_model is called when production main_model is exited.
func (s *BasePlSqlParserListener) ExitMain_model(ctx *Main_modelContext) {}

// EnterModel_column_clauses is called when production model_column_clauses is entered.
func (s *BasePlSqlParserListener) EnterModel_column_clauses(ctx *Model_column_clausesContext) {}

// ExitModel_column_clauses is called when production model_column_clauses is exited.
func (s *BasePlSqlParserListener) ExitModel_column_clauses(ctx *Model_column_clausesContext) {}

// EnterModel_column_partition_part is called when production model_column_partition_part is entered.
func (s *BasePlSqlParserListener) EnterModel_column_partition_part(ctx *Model_column_partition_partContext) {
}

// ExitModel_column_partition_part is called when production model_column_partition_part is exited.
func (s *BasePlSqlParserListener) ExitModel_column_partition_part(ctx *Model_column_partition_partContext) {
}

// EnterModel_column_list is called when production model_column_list is entered.
func (s *BasePlSqlParserListener) EnterModel_column_list(ctx *Model_column_listContext) {}

// ExitModel_column_list is called when production model_column_list is exited.
func (s *BasePlSqlParserListener) ExitModel_column_list(ctx *Model_column_listContext) {}

// EnterModel_column is called when production model_column is entered.
func (s *BasePlSqlParserListener) EnterModel_column(ctx *Model_columnContext) {}

// ExitModel_column is called when production model_column is exited.
func (s *BasePlSqlParserListener) ExitModel_column(ctx *Model_columnContext) {}

// EnterModel_rules_clause is called when production model_rules_clause is entered.
func (s *BasePlSqlParserListener) EnterModel_rules_clause(ctx *Model_rules_clauseContext) {}

// ExitModel_rules_clause is called when production model_rules_clause is exited.
func (s *BasePlSqlParserListener) ExitModel_rules_clause(ctx *Model_rules_clauseContext) {}

// EnterModel_rules_part is called when production model_rules_part is entered.
func (s *BasePlSqlParserListener) EnterModel_rules_part(ctx *Model_rules_partContext) {}

// ExitModel_rules_part is called when production model_rules_part is exited.
func (s *BasePlSqlParserListener) ExitModel_rules_part(ctx *Model_rules_partContext) {}

// EnterModel_rules_element is called when production model_rules_element is entered.
func (s *BasePlSqlParserListener) EnterModel_rules_element(ctx *Model_rules_elementContext) {}

// ExitModel_rules_element is called when production model_rules_element is exited.
func (s *BasePlSqlParserListener) ExitModel_rules_element(ctx *Model_rules_elementContext) {}

// EnterCell_assignment is called when production cell_assignment is entered.
func (s *BasePlSqlParserListener) EnterCell_assignment(ctx *Cell_assignmentContext) {}

// ExitCell_assignment is called when production cell_assignment is exited.
func (s *BasePlSqlParserListener) ExitCell_assignment(ctx *Cell_assignmentContext) {}

// EnterModel_iterate_clause is called when production model_iterate_clause is entered.
func (s *BasePlSqlParserListener) EnterModel_iterate_clause(ctx *Model_iterate_clauseContext) {}

// ExitModel_iterate_clause is called when production model_iterate_clause is exited.
func (s *BasePlSqlParserListener) ExitModel_iterate_clause(ctx *Model_iterate_clauseContext) {}

// EnterUntil_part is called when production until_part is entered.
func (s *BasePlSqlParserListener) EnterUntil_part(ctx *Until_partContext) {}

// ExitUntil_part is called when production until_part is exited.
func (s *BasePlSqlParserListener) ExitUntil_part(ctx *Until_partContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BasePlSqlParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BasePlSqlParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterOrder_by_elements is called when production order_by_elements is entered.
func (s *BasePlSqlParserListener) EnterOrder_by_elements(ctx *Order_by_elementsContext) {}

// ExitOrder_by_elements is called when production order_by_elements is exited.
func (s *BasePlSqlParserListener) ExitOrder_by_elements(ctx *Order_by_elementsContext) {}

// EnterFor_update_clause is called when production for_update_clause is entered.
func (s *BasePlSqlParserListener) EnterFor_update_clause(ctx *For_update_clauseContext) {}

// ExitFor_update_clause is called when production for_update_clause is exited.
func (s *BasePlSqlParserListener) ExitFor_update_clause(ctx *For_update_clauseContext) {}

// EnterFor_update_of_part is called when production for_update_of_part is entered.
func (s *BasePlSqlParserListener) EnterFor_update_of_part(ctx *For_update_of_partContext) {}

// ExitFor_update_of_part is called when production for_update_of_part is exited.
func (s *BasePlSqlParserListener) ExitFor_update_of_part(ctx *For_update_of_partContext) {}

// EnterFor_update_options is called when production for_update_options is entered.
func (s *BasePlSqlParserListener) EnterFor_update_options(ctx *For_update_optionsContext) {}

// ExitFor_update_options is called when production for_update_options is exited.
func (s *BasePlSqlParserListener) ExitFor_update_options(ctx *For_update_optionsContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BasePlSqlParserListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BasePlSqlParserListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterUpdate_set_clause is called when production update_set_clause is entered.
func (s *BasePlSqlParserListener) EnterUpdate_set_clause(ctx *Update_set_clauseContext) {}

// ExitUpdate_set_clause is called when production update_set_clause is exited.
func (s *BasePlSqlParserListener) ExitUpdate_set_clause(ctx *Update_set_clauseContext) {}

// EnterColumn_based_update_set_clause is called when production column_based_update_set_clause is entered.
func (s *BasePlSqlParserListener) EnterColumn_based_update_set_clause(ctx *Column_based_update_set_clauseContext) {
}

// ExitColumn_based_update_set_clause is called when production column_based_update_set_clause is exited.
func (s *BasePlSqlParserListener) ExitColumn_based_update_set_clause(ctx *Column_based_update_set_clauseContext) {
}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BasePlSqlParserListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BasePlSqlParserListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *BasePlSqlParserListener) EnterInsert_statement(ctx *Insert_statementContext) {}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *BasePlSqlParserListener) ExitInsert_statement(ctx *Insert_statementContext) {}

// EnterSingle_table_insert is called when production single_table_insert is entered.
func (s *BasePlSqlParserListener) EnterSingle_table_insert(ctx *Single_table_insertContext) {}

// ExitSingle_table_insert is called when production single_table_insert is exited.
func (s *BasePlSqlParserListener) ExitSingle_table_insert(ctx *Single_table_insertContext) {}

// EnterMulti_table_insert is called when production multi_table_insert is entered.
func (s *BasePlSqlParserListener) EnterMulti_table_insert(ctx *Multi_table_insertContext) {}

// ExitMulti_table_insert is called when production multi_table_insert is exited.
func (s *BasePlSqlParserListener) ExitMulti_table_insert(ctx *Multi_table_insertContext) {}

// EnterMulti_table_element is called when production multi_table_element is entered.
func (s *BasePlSqlParserListener) EnterMulti_table_element(ctx *Multi_table_elementContext) {}

// ExitMulti_table_element is called when production multi_table_element is exited.
func (s *BasePlSqlParserListener) ExitMulti_table_element(ctx *Multi_table_elementContext) {}

// EnterConditional_insert_clause is called when production conditional_insert_clause is entered.
func (s *BasePlSqlParserListener) EnterConditional_insert_clause(ctx *Conditional_insert_clauseContext) {
}

// ExitConditional_insert_clause is called when production conditional_insert_clause is exited.
func (s *BasePlSqlParserListener) ExitConditional_insert_clause(ctx *Conditional_insert_clauseContext) {
}

// EnterConditional_insert_when_part is called when production conditional_insert_when_part is entered.
func (s *BasePlSqlParserListener) EnterConditional_insert_when_part(ctx *Conditional_insert_when_partContext) {
}

// ExitConditional_insert_when_part is called when production conditional_insert_when_part is exited.
func (s *BasePlSqlParserListener) ExitConditional_insert_when_part(ctx *Conditional_insert_when_partContext) {
}

// EnterConditional_insert_else_part is called when production conditional_insert_else_part is entered.
func (s *BasePlSqlParserListener) EnterConditional_insert_else_part(ctx *Conditional_insert_else_partContext) {
}

// ExitConditional_insert_else_part is called when production conditional_insert_else_part is exited.
func (s *BasePlSqlParserListener) ExitConditional_insert_else_part(ctx *Conditional_insert_else_partContext) {
}

// EnterInsert_into_clause is called when production insert_into_clause is entered.
func (s *BasePlSqlParserListener) EnterInsert_into_clause(ctx *Insert_into_clauseContext) {}

// ExitInsert_into_clause is called when production insert_into_clause is exited.
func (s *BasePlSqlParserListener) ExitInsert_into_clause(ctx *Insert_into_clauseContext) {}

// EnterValues_clause is called when production values_clause is entered.
func (s *BasePlSqlParserListener) EnterValues_clause(ctx *Values_clauseContext) {}

// ExitValues_clause is called when production values_clause is exited.
func (s *BasePlSqlParserListener) ExitValues_clause(ctx *Values_clauseContext) {}

// EnterMerge_statement is called when production merge_statement is entered.
func (s *BasePlSqlParserListener) EnterMerge_statement(ctx *Merge_statementContext) {}

// ExitMerge_statement is called when production merge_statement is exited.
func (s *BasePlSqlParserListener) ExitMerge_statement(ctx *Merge_statementContext) {}

// EnterMerge_update_clause is called when production merge_update_clause is entered.
func (s *BasePlSqlParserListener) EnterMerge_update_clause(ctx *Merge_update_clauseContext) {}

// ExitMerge_update_clause is called when production merge_update_clause is exited.
func (s *BasePlSqlParserListener) ExitMerge_update_clause(ctx *Merge_update_clauseContext) {}

// EnterMerge_element is called when production merge_element is entered.
func (s *BasePlSqlParserListener) EnterMerge_element(ctx *Merge_elementContext) {}

// ExitMerge_element is called when production merge_element is exited.
func (s *BasePlSqlParserListener) ExitMerge_element(ctx *Merge_elementContext) {}

// EnterMerge_update_delete_part is called when production merge_update_delete_part is entered.
func (s *BasePlSqlParserListener) EnterMerge_update_delete_part(ctx *Merge_update_delete_partContext) {
}

// ExitMerge_update_delete_part is called when production merge_update_delete_part is exited.
func (s *BasePlSqlParserListener) ExitMerge_update_delete_part(ctx *Merge_update_delete_partContext) {}

// EnterMerge_insert_clause is called when production merge_insert_clause is entered.
func (s *BasePlSqlParserListener) EnterMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// ExitMerge_insert_clause is called when production merge_insert_clause is exited.
func (s *BasePlSqlParserListener) ExitMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// EnterSelected_tableview is called when production selected_tableview is entered.
func (s *BasePlSqlParserListener) EnterSelected_tableview(ctx *Selected_tableviewContext) {}

// ExitSelected_tableview is called when production selected_tableview is exited.
func (s *BasePlSqlParserListener) ExitSelected_tableview(ctx *Selected_tableviewContext) {}

// EnterLock_table_statement is called when production lock_table_statement is entered.
func (s *BasePlSqlParserListener) EnterLock_table_statement(ctx *Lock_table_statementContext) {}

// ExitLock_table_statement is called when production lock_table_statement is exited.
func (s *BasePlSqlParserListener) ExitLock_table_statement(ctx *Lock_table_statementContext) {}

// EnterWait_nowait_part is called when production wait_nowait_part is entered.
func (s *BasePlSqlParserListener) EnterWait_nowait_part(ctx *Wait_nowait_partContext) {}

// ExitWait_nowait_part is called when production wait_nowait_part is exited.
func (s *BasePlSqlParserListener) ExitWait_nowait_part(ctx *Wait_nowait_partContext) {}

// EnterLock_table_element is called when production lock_table_element is entered.
func (s *BasePlSqlParserListener) EnterLock_table_element(ctx *Lock_table_elementContext) {}

// ExitLock_table_element is called when production lock_table_element is exited.
func (s *BasePlSqlParserListener) ExitLock_table_element(ctx *Lock_table_elementContext) {}

// EnterLock_mode is called when production lock_mode is entered.
func (s *BasePlSqlParserListener) EnterLock_mode(ctx *Lock_modeContext) {}

// ExitLock_mode is called when production lock_mode is exited.
func (s *BasePlSqlParserListener) ExitLock_mode(ctx *Lock_modeContext) {}

// EnterGeneral_table_ref is called when production general_table_ref is entered.
func (s *BasePlSqlParserListener) EnterGeneral_table_ref(ctx *General_table_refContext) {}

// ExitGeneral_table_ref is called when production general_table_ref is exited.
func (s *BasePlSqlParserListener) ExitGeneral_table_ref(ctx *General_table_refContext) {}

// EnterStatic_returning_clause is called when production static_returning_clause is entered.
func (s *BasePlSqlParserListener) EnterStatic_returning_clause(ctx *Static_returning_clauseContext) {}

// ExitStatic_returning_clause is called when production static_returning_clause is exited.
func (s *BasePlSqlParserListener) ExitStatic_returning_clause(ctx *Static_returning_clauseContext) {}

// EnterError_logging_clause is called when production error_logging_clause is entered.
func (s *BasePlSqlParserListener) EnterError_logging_clause(ctx *Error_logging_clauseContext) {}

// ExitError_logging_clause is called when production error_logging_clause is exited.
func (s *BasePlSqlParserListener) ExitError_logging_clause(ctx *Error_logging_clauseContext) {}

// EnterError_logging_into_part is called when production error_logging_into_part is entered.
func (s *BasePlSqlParserListener) EnterError_logging_into_part(ctx *Error_logging_into_partContext) {}

// ExitError_logging_into_part is called when production error_logging_into_part is exited.
func (s *BasePlSqlParserListener) ExitError_logging_into_part(ctx *Error_logging_into_partContext) {}

// EnterError_logging_reject_part is called when production error_logging_reject_part is entered.
func (s *BasePlSqlParserListener) EnterError_logging_reject_part(ctx *Error_logging_reject_partContext) {
}

// ExitError_logging_reject_part is called when production error_logging_reject_part is exited.
func (s *BasePlSqlParserListener) ExitError_logging_reject_part(ctx *Error_logging_reject_partContext) {
}

// EnterDml_table_expression_clause is called when production dml_table_expression_clause is entered.
func (s *BasePlSqlParserListener) EnterDml_table_expression_clause(ctx *Dml_table_expression_clauseContext) {
}

// ExitDml_table_expression_clause is called when production dml_table_expression_clause is exited.
func (s *BasePlSqlParserListener) ExitDml_table_expression_clause(ctx *Dml_table_expression_clauseContext) {
}

// EnterTable_collection_expression is called when production table_collection_expression is entered.
func (s *BasePlSqlParserListener) EnterTable_collection_expression(ctx *Table_collection_expressionContext) {
}

// ExitTable_collection_expression is called when production table_collection_expression is exited.
func (s *BasePlSqlParserListener) ExitTable_collection_expression(ctx *Table_collection_expressionContext) {
}

// EnterSubquery_restriction_clause is called when production subquery_restriction_clause is entered.
func (s *BasePlSqlParserListener) EnterSubquery_restriction_clause(ctx *Subquery_restriction_clauseContext) {
}

// ExitSubquery_restriction_clause is called when production subquery_restriction_clause is exited.
func (s *BasePlSqlParserListener) ExitSubquery_restriction_clause(ctx *Subquery_restriction_clauseContext) {
}

// EnterSample_clause is called when production sample_clause is entered.
func (s *BasePlSqlParserListener) EnterSample_clause(ctx *Sample_clauseContext) {}

// ExitSample_clause is called when production sample_clause is exited.
func (s *BasePlSqlParserListener) ExitSample_clause(ctx *Sample_clauseContext) {}

// EnterSeed_part is called when production seed_part is entered.
func (s *BasePlSqlParserListener) EnterSeed_part(ctx *Seed_partContext) {}

// ExitSeed_part is called when production seed_part is exited.
func (s *BasePlSqlParserListener) ExitSeed_part(ctx *Seed_partContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasePlSqlParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasePlSqlParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BasePlSqlParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BasePlSqlParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePlSqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePlSqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCursor_expression is called when production cursor_expression is entered.
func (s *BasePlSqlParserListener) EnterCursor_expression(ctx *Cursor_expressionContext) {}

// ExitCursor_expression is called when production cursor_expression is exited.
func (s *BasePlSqlParserListener) ExitCursor_expression(ctx *Cursor_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BasePlSqlParserListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BasePlSqlParserListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterMultiset_expression is called when production multiset_expression is entered.
func (s *BasePlSqlParserListener) EnterMultiset_expression(ctx *Multiset_expressionContext) {}

// ExitMultiset_expression is called when production multiset_expression is exited.
func (s *BasePlSqlParserListener) ExitMultiset_expression(ctx *Multiset_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BasePlSqlParserListener) EnterRelational_expression(ctx *Relational_expressionContext) {}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BasePlSqlParserListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterCompound_expression is called when production compound_expression is entered.
func (s *BasePlSqlParserListener) EnterCompound_expression(ctx *Compound_expressionContext) {}

// ExitCompound_expression is called when production compound_expression is exited.
func (s *BasePlSqlParserListener) ExitCompound_expression(ctx *Compound_expressionContext) {}

// EnterRelational_operator is called when production relational_operator is entered.
func (s *BasePlSqlParserListener) EnterRelational_operator(ctx *Relational_operatorContext) {}

// ExitRelational_operator is called when production relational_operator is exited.
func (s *BasePlSqlParserListener) ExitRelational_operator(ctx *Relational_operatorContext) {}

// EnterIn_elements is called when production in_elements is entered.
func (s *BasePlSqlParserListener) EnterIn_elements(ctx *In_elementsContext) {}

// ExitIn_elements is called when production in_elements is exited.
func (s *BasePlSqlParserListener) ExitIn_elements(ctx *In_elementsContext) {}

// EnterBetween_elements is called when production between_elements is entered.
func (s *BasePlSqlParserListener) EnterBetween_elements(ctx *Between_elementsContext) {}

// ExitBetween_elements is called when production between_elements is exited.
func (s *BasePlSqlParserListener) ExitBetween_elements(ctx *Between_elementsContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BasePlSqlParserListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BasePlSqlParserListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterInterval_expression is called when production interval_expression is entered.
func (s *BasePlSqlParserListener) EnterInterval_expression(ctx *Interval_expressionContext) {}

// ExitInterval_expression is called when production interval_expression is exited.
func (s *BasePlSqlParserListener) ExitInterval_expression(ctx *Interval_expressionContext) {}

// EnterModel_expression is called when production model_expression is entered.
func (s *BasePlSqlParserListener) EnterModel_expression(ctx *Model_expressionContext) {}

// ExitModel_expression is called when production model_expression is exited.
func (s *BasePlSqlParserListener) ExitModel_expression(ctx *Model_expressionContext) {}

// EnterModel_expression_element is called when production model_expression_element is entered.
func (s *BasePlSqlParserListener) EnterModel_expression_element(ctx *Model_expression_elementContext) {
}

// ExitModel_expression_element is called when production model_expression_element is exited.
func (s *BasePlSqlParserListener) ExitModel_expression_element(ctx *Model_expression_elementContext) {}

// EnterSingle_column_for_loop is called when production single_column_for_loop is entered.
func (s *BasePlSqlParserListener) EnterSingle_column_for_loop(ctx *Single_column_for_loopContext) {}

// ExitSingle_column_for_loop is called when production single_column_for_loop is exited.
func (s *BasePlSqlParserListener) ExitSingle_column_for_loop(ctx *Single_column_for_loopContext) {}

// EnterMulti_column_for_loop is called when production multi_column_for_loop is entered.
func (s *BasePlSqlParserListener) EnterMulti_column_for_loop(ctx *Multi_column_for_loopContext) {}

// ExitMulti_column_for_loop is called when production multi_column_for_loop is exited.
func (s *BasePlSqlParserListener) ExitMulti_column_for_loop(ctx *Multi_column_for_loopContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BasePlSqlParserListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BasePlSqlParserListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterCase_statement is called when production case_statement is entered.
func (s *BasePlSqlParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BasePlSqlParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterSimple_case_statement is called when production simple_case_statement is entered.
func (s *BasePlSqlParserListener) EnterSimple_case_statement(ctx *Simple_case_statementContext) {}

// ExitSimple_case_statement is called when production simple_case_statement is exited.
func (s *BasePlSqlParserListener) ExitSimple_case_statement(ctx *Simple_case_statementContext) {}

// EnterSimple_case_when_part is called when production simple_case_when_part is entered.
func (s *BasePlSqlParserListener) EnterSimple_case_when_part(ctx *Simple_case_when_partContext) {}

// ExitSimple_case_when_part is called when production simple_case_when_part is exited.
func (s *BasePlSqlParserListener) ExitSimple_case_when_part(ctx *Simple_case_when_partContext) {}

// EnterSearched_case_statement is called when production searched_case_statement is entered.
func (s *BasePlSqlParserListener) EnterSearched_case_statement(ctx *Searched_case_statementContext) {}

// ExitSearched_case_statement is called when production searched_case_statement is exited.
func (s *BasePlSqlParserListener) ExitSearched_case_statement(ctx *Searched_case_statementContext) {}

// EnterSearched_case_when_part is called when production searched_case_when_part is entered.
func (s *BasePlSqlParserListener) EnterSearched_case_when_part(ctx *Searched_case_when_partContext) {}

// ExitSearched_case_when_part is called when production searched_case_when_part is exited.
func (s *BasePlSqlParserListener) ExitSearched_case_when_part(ctx *Searched_case_when_partContext) {}

// EnterCase_else_part is called when production case_else_part is entered.
func (s *BasePlSqlParserListener) EnterCase_else_part(ctx *Case_else_partContext) {}

// ExitCase_else_part is called when production case_else_part is exited.
func (s *BasePlSqlParserListener) ExitCase_else_part(ctx *Case_else_partContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePlSqlParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePlSqlParserListener) ExitAtom(ctx *AtomContext) {}

// EnterQuantified_expression is called when production quantified_expression is entered.
func (s *BasePlSqlParserListener) EnterQuantified_expression(ctx *Quantified_expressionContext) {}

// ExitQuantified_expression is called when production quantified_expression is exited.
func (s *BasePlSqlParserListener) ExitQuantified_expression(ctx *Quantified_expressionContext) {}

// EnterString_function is called when production string_function is entered.
func (s *BasePlSqlParserListener) EnterString_function(ctx *String_functionContext) {}

// ExitString_function is called when production string_function is exited.
func (s *BasePlSqlParserListener) ExitString_function(ctx *String_functionContext) {}

// EnterStandard_function is called when production standard_function is entered.
func (s *BasePlSqlParserListener) EnterStandard_function(ctx *Standard_functionContext) {}

// ExitStandard_function is called when production standard_function is exited.
func (s *BasePlSqlParserListener) ExitStandard_function(ctx *Standard_functionContext) {}

// EnterNumeric_function_wrapper is called when production numeric_function_wrapper is entered.
func (s *BasePlSqlParserListener) EnterNumeric_function_wrapper(ctx *Numeric_function_wrapperContext) {
}

// ExitNumeric_function_wrapper is called when production numeric_function_wrapper is exited.
func (s *BasePlSqlParserListener) ExitNumeric_function_wrapper(ctx *Numeric_function_wrapperContext) {}

// EnterNumeric_function is called when production numeric_function is entered.
func (s *BasePlSqlParserListener) EnterNumeric_function(ctx *Numeric_functionContext) {}

// ExitNumeric_function is called when production numeric_function is exited.
func (s *BasePlSqlParserListener) ExitNumeric_function(ctx *Numeric_functionContext) {}

// EnterOther_function is called when production other_function is entered.
func (s *BasePlSqlParserListener) EnterOther_function(ctx *Other_functionContext) {}

// ExitOther_function is called when production other_function is exited.
func (s *BasePlSqlParserListener) ExitOther_function(ctx *Other_functionContext) {}

// EnterOver_clause_keyword is called when production over_clause_keyword is entered.
func (s *BasePlSqlParserListener) EnterOver_clause_keyword(ctx *Over_clause_keywordContext) {}

// ExitOver_clause_keyword is called when production over_clause_keyword is exited.
func (s *BasePlSqlParserListener) ExitOver_clause_keyword(ctx *Over_clause_keywordContext) {}

// EnterWithin_or_over_clause_keyword is called when production within_or_over_clause_keyword is entered.
func (s *BasePlSqlParserListener) EnterWithin_or_over_clause_keyword(ctx *Within_or_over_clause_keywordContext) {
}

// ExitWithin_or_over_clause_keyword is called when production within_or_over_clause_keyword is exited.
func (s *BasePlSqlParserListener) ExitWithin_or_over_clause_keyword(ctx *Within_or_over_clause_keywordContext) {
}

// EnterStandard_prediction_function_keyword is called when production standard_prediction_function_keyword is entered.
func (s *BasePlSqlParserListener) EnterStandard_prediction_function_keyword(ctx *Standard_prediction_function_keywordContext) {
}

// ExitStandard_prediction_function_keyword is called when production standard_prediction_function_keyword is exited.
func (s *BasePlSqlParserListener) ExitStandard_prediction_function_keyword(ctx *Standard_prediction_function_keywordContext) {
}

// EnterOver_clause is called when production over_clause is entered.
func (s *BasePlSqlParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BasePlSqlParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterWindowing_clause is called when production windowing_clause is entered.
func (s *BasePlSqlParserListener) EnterWindowing_clause(ctx *Windowing_clauseContext) {}

// ExitWindowing_clause is called when production windowing_clause is exited.
func (s *BasePlSqlParserListener) ExitWindowing_clause(ctx *Windowing_clauseContext) {}

// EnterWindowing_type is called when production windowing_type is entered.
func (s *BasePlSqlParserListener) EnterWindowing_type(ctx *Windowing_typeContext) {}

// ExitWindowing_type is called when production windowing_type is exited.
func (s *BasePlSqlParserListener) ExitWindowing_type(ctx *Windowing_typeContext) {}

// EnterWindowing_elements is called when production windowing_elements is entered.
func (s *BasePlSqlParserListener) EnterWindowing_elements(ctx *Windowing_elementsContext) {}

// ExitWindowing_elements is called when production windowing_elements is exited.
func (s *BasePlSqlParserListener) ExitWindowing_elements(ctx *Windowing_elementsContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BasePlSqlParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BasePlSqlParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterUsing_element is called when production using_element is entered.
func (s *BasePlSqlParserListener) EnterUsing_element(ctx *Using_elementContext) {}

// ExitUsing_element is called when production using_element is exited.
func (s *BasePlSqlParserListener) ExitUsing_element(ctx *Using_elementContext) {}

// EnterCollect_order_by_part is called when production collect_order_by_part is entered.
func (s *BasePlSqlParserListener) EnterCollect_order_by_part(ctx *Collect_order_by_partContext) {}

// ExitCollect_order_by_part is called when production collect_order_by_part is exited.
func (s *BasePlSqlParserListener) ExitCollect_order_by_part(ctx *Collect_order_by_partContext) {}

// EnterWithin_or_over_part is called when production within_or_over_part is entered.
func (s *BasePlSqlParserListener) EnterWithin_or_over_part(ctx *Within_or_over_partContext) {}

// ExitWithin_or_over_part is called when production within_or_over_part is exited.
func (s *BasePlSqlParserListener) ExitWithin_or_over_part(ctx *Within_or_over_partContext) {}

// EnterCost_matrix_clause is called when production cost_matrix_clause is entered.
func (s *BasePlSqlParserListener) EnterCost_matrix_clause(ctx *Cost_matrix_clauseContext) {}

// ExitCost_matrix_clause is called when production cost_matrix_clause is exited.
func (s *BasePlSqlParserListener) ExitCost_matrix_clause(ctx *Cost_matrix_clauseContext) {}

// EnterXml_passing_clause is called when production xml_passing_clause is entered.
func (s *BasePlSqlParserListener) EnterXml_passing_clause(ctx *Xml_passing_clauseContext) {}

// ExitXml_passing_clause is called when production xml_passing_clause is exited.
func (s *BasePlSqlParserListener) ExitXml_passing_clause(ctx *Xml_passing_clauseContext) {}

// EnterXml_attributes_clause is called when production xml_attributes_clause is entered.
func (s *BasePlSqlParserListener) EnterXml_attributes_clause(ctx *Xml_attributes_clauseContext) {}

// ExitXml_attributes_clause is called when production xml_attributes_clause is exited.
func (s *BasePlSqlParserListener) ExitXml_attributes_clause(ctx *Xml_attributes_clauseContext) {}

// EnterXml_namespaces_clause is called when production xml_namespaces_clause is entered.
func (s *BasePlSqlParserListener) EnterXml_namespaces_clause(ctx *Xml_namespaces_clauseContext) {}

// ExitXml_namespaces_clause is called when production xml_namespaces_clause is exited.
func (s *BasePlSqlParserListener) ExitXml_namespaces_clause(ctx *Xml_namespaces_clauseContext) {}

// EnterXml_table_column is called when production xml_table_column is entered.
func (s *BasePlSqlParserListener) EnterXml_table_column(ctx *Xml_table_columnContext) {}

// ExitXml_table_column is called when production xml_table_column is exited.
func (s *BasePlSqlParserListener) ExitXml_table_column(ctx *Xml_table_columnContext) {}

// EnterXml_general_default_part is called when production xml_general_default_part is entered.
func (s *BasePlSqlParserListener) EnterXml_general_default_part(ctx *Xml_general_default_partContext) {
}

// ExitXml_general_default_part is called when production xml_general_default_part is exited.
func (s *BasePlSqlParserListener) ExitXml_general_default_part(ctx *Xml_general_default_partContext) {}

// EnterXml_multiuse_expression_element is called when production xml_multiuse_expression_element is entered.
func (s *BasePlSqlParserListener) EnterXml_multiuse_expression_element(ctx *Xml_multiuse_expression_elementContext) {
}

// ExitXml_multiuse_expression_element is called when production xml_multiuse_expression_element is exited.
func (s *BasePlSqlParserListener) ExitXml_multiuse_expression_element(ctx *Xml_multiuse_expression_elementContext) {
}

// EnterXmlroot_param_version_part is called when production xmlroot_param_version_part is entered.
func (s *BasePlSqlParserListener) EnterXmlroot_param_version_part(ctx *Xmlroot_param_version_partContext) {
}

// ExitXmlroot_param_version_part is called when production xmlroot_param_version_part is exited.
func (s *BasePlSqlParserListener) ExitXmlroot_param_version_part(ctx *Xmlroot_param_version_partContext) {
}

// EnterXmlroot_param_standalone_part is called when production xmlroot_param_standalone_part is entered.
func (s *BasePlSqlParserListener) EnterXmlroot_param_standalone_part(ctx *Xmlroot_param_standalone_partContext) {
}

// ExitXmlroot_param_standalone_part is called when production xmlroot_param_standalone_part is exited.
func (s *BasePlSqlParserListener) ExitXmlroot_param_standalone_part(ctx *Xmlroot_param_standalone_partContext) {
}

// EnterXmlserialize_param_enconding_part is called when production xmlserialize_param_enconding_part is entered.
func (s *BasePlSqlParserListener) EnterXmlserialize_param_enconding_part(ctx *Xmlserialize_param_enconding_partContext) {
}

// ExitXmlserialize_param_enconding_part is called when production xmlserialize_param_enconding_part is exited.
func (s *BasePlSqlParserListener) ExitXmlserialize_param_enconding_part(ctx *Xmlserialize_param_enconding_partContext) {
}

// EnterXmlserialize_param_version_part is called when production xmlserialize_param_version_part is entered.
func (s *BasePlSqlParserListener) EnterXmlserialize_param_version_part(ctx *Xmlserialize_param_version_partContext) {
}

// ExitXmlserialize_param_version_part is called when production xmlserialize_param_version_part is exited.
func (s *BasePlSqlParserListener) ExitXmlserialize_param_version_part(ctx *Xmlserialize_param_version_partContext) {
}

// EnterXmlserialize_param_ident_part is called when production xmlserialize_param_ident_part is entered.
func (s *BasePlSqlParserListener) EnterXmlserialize_param_ident_part(ctx *Xmlserialize_param_ident_partContext) {
}

// ExitXmlserialize_param_ident_part is called when production xmlserialize_param_ident_part is exited.
func (s *BasePlSqlParserListener) ExitXmlserialize_param_ident_part(ctx *Xmlserialize_param_ident_partContext) {
}

// EnterSql_plus_command is called when production sql_plus_command is entered.
func (s *BasePlSqlParserListener) EnterSql_plus_command(ctx *Sql_plus_commandContext) {}

// ExitSql_plus_command is called when production sql_plus_command is exited.
func (s *BasePlSqlParserListener) ExitSql_plus_command(ctx *Sql_plus_commandContext) {}

// EnterWhenever_command is called when production whenever_command is entered.
func (s *BasePlSqlParserListener) EnterWhenever_command(ctx *Whenever_commandContext) {}

// ExitWhenever_command is called when production whenever_command is exited.
func (s *BasePlSqlParserListener) ExitWhenever_command(ctx *Whenever_commandContext) {}

// EnterSet_command is called when production set_command is entered.
func (s *BasePlSqlParserListener) EnterSet_command(ctx *Set_commandContext) {}

// ExitSet_command is called when production set_command is exited.
func (s *BasePlSqlParserListener) ExitSet_command(ctx *Set_commandContext) {}

// EnterPartition_extension_clause is called when production partition_extension_clause is entered.
func (s *BasePlSqlParserListener) EnterPartition_extension_clause(ctx *Partition_extension_clauseContext) {
}

// ExitPartition_extension_clause is called when production partition_extension_clause is exited.
func (s *BasePlSqlParserListener) ExitPartition_extension_clause(ctx *Partition_extension_clauseContext) {
}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BasePlSqlParserListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BasePlSqlParserListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BasePlSqlParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BasePlSqlParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BasePlSqlParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BasePlSqlParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterInto_clause is called when production into_clause is entered.
func (s *BasePlSqlParserListener) EnterInto_clause(ctx *Into_clauseContext) {}

// ExitInto_clause is called when production into_clause is exited.
func (s *BasePlSqlParserListener) ExitInto_clause(ctx *Into_clauseContext) {}

// EnterXml_column_name is called when production xml_column_name is entered.
func (s *BasePlSqlParserListener) EnterXml_column_name(ctx *Xml_column_nameContext) {}

// ExitXml_column_name is called when production xml_column_name is exited.
func (s *BasePlSqlParserListener) ExitXml_column_name(ctx *Xml_column_nameContext) {}

// EnterCost_class_name is called when production cost_class_name is entered.
func (s *BasePlSqlParserListener) EnterCost_class_name(ctx *Cost_class_nameContext) {}

// ExitCost_class_name is called when production cost_class_name is exited.
func (s *BasePlSqlParserListener) ExitCost_class_name(ctx *Cost_class_nameContext) {}

// EnterAttribute_name is called when production attribute_name is entered.
func (s *BasePlSqlParserListener) EnterAttribute_name(ctx *Attribute_nameContext) {}

// ExitAttribute_name is called when production attribute_name is exited.
func (s *BasePlSqlParserListener) ExitAttribute_name(ctx *Attribute_nameContext) {}

// EnterSavepoint_name is called when production savepoint_name is entered.
func (s *BasePlSqlParserListener) EnterSavepoint_name(ctx *Savepoint_nameContext) {}

// ExitSavepoint_name is called when production savepoint_name is exited.
func (s *BasePlSqlParserListener) ExitSavepoint_name(ctx *Savepoint_nameContext) {}

// EnterRollback_segment_name is called when production rollback_segment_name is entered.
func (s *BasePlSqlParserListener) EnterRollback_segment_name(ctx *Rollback_segment_nameContext) {}

// ExitRollback_segment_name is called when production rollback_segment_name is exited.
func (s *BasePlSqlParserListener) ExitRollback_segment_name(ctx *Rollback_segment_nameContext) {}

// EnterTable_var_name is called when production table_var_name is entered.
func (s *BasePlSqlParserListener) EnterTable_var_name(ctx *Table_var_nameContext) {}

// ExitTable_var_name is called when production table_var_name is exited.
func (s *BasePlSqlParserListener) ExitTable_var_name(ctx *Table_var_nameContext) {}

// EnterSchema_name is called when production schema_name is entered.
func (s *BasePlSqlParserListener) EnterSchema_name(ctx *Schema_nameContext) {}

// ExitSchema_name is called when production schema_name is exited.
func (s *BasePlSqlParserListener) ExitSchema_name(ctx *Schema_nameContext) {}

// EnterRoutine_name is called when production routine_name is entered.
func (s *BasePlSqlParserListener) EnterRoutine_name(ctx *Routine_nameContext) {}

// ExitRoutine_name is called when production routine_name is exited.
func (s *BasePlSqlParserListener) ExitRoutine_name(ctx *Routine_nameContext) {}

// EnterPackage_name is called when production package_name is entered.
func (s *BasePlSqlParserListener) EnterPackage_name(ctx *Package_nameContext) {}

// ExitPackage_name is called when production package_name is exited.
func (s *BasePlSqlParserListener) ExitPackage_name(ctx *Package_nameContext) {}

// EnterImplementation_type_name is called when production implementation_type_name is entered.
func (s *BasePlSqlParserListener) EnterImplementation_type_name(ctx *Implementation_type_nameContext) {
}

// ExitImplementation_type_name is called when production implementation_type_name is exited.
func (s *BasePlSqlParserListener) ExitImplementation_type_name(ctx *Implementation_type_nameContext) {}

// EnterParameter_name is called when production parameter_name is entered.
func (s *BasePlSqlParserListener) EnterParameter_name(ctx *Parameter_nameContext) {}

// ExitParameter_name is called when production parameter_name is exited.
func (s *BasePlSqlParserListener) ExitParameter_name(ctx *Parameter_nameContext) {}

// EnterReference_model_name is called when production reference_model_name is entered.
func (s *BasePlSqlParserListener) EnterReference_model_name(ctx *Reference_model_nameContext) {}

// ExitReference_model_name is called when production reference_model_name is exited.
func (s *BasePlSqlParserListener) ExitReference_model_name(ctx *Reference_model_nameContext) {}

// EnterMain_model_name is called when production main_model_name is entered.
func (s *BasePlSqlParserListener) EnterMain_model_name(ctx *Main_model_nameContext) {}

// ExitMain_model_name is called when production main_model_name is exited.
func (s *BasePlSqlParserListener) ExitMain_model_name(ctx *Main_model_nameContext) {}

// EnterAggregate_function_name is called when production aggregate_function_name is entered.
func (s *BasePlSqlParserListener) EnterAggregate_function_name(ctx *Aggregate_function_nameContext) {}

// ExitAggregate_function_name is called when production aggregate_function_name is exited.
func (s *BasePlSqlParserListener) ExitAggregate_function_name(ctx *Aggregate_function_nameContext) {}

// EnterQuery_name is called when production query_name is entered.
func (s *BasePlSqlParserListener) EnterQuery_name(ctx *Query_nameContext) {}

// ExitQuery_name is called when production query_name is exited.
func (s *BasePlSqlParserListener) ExitQuery_name(ctx *Query_nameContext) {}

// EnterConstraint_name is called when production constraint_name is entered.
func (s *BasePlSqlParserListener) EnterConstraint_name(ctx *Constraint_nameContext) {}

// ExitConstraint_name is called when production constraint_name is exited.
func (s *BasePlSqlParserListener) ExitConstraint_name(ctx *Constraint_nameContext) {}

// EnterLabel_name is called when production label_name is entered.
func (s *BasePlSqlParserListener) EnterLabel_name(ctx *Label_nameContext) {}

// ExitLabel_name is called when production label_name is exited.
func (s *BasePlSqlParserListener) ExitLabel_name(ctx *Label_nameContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BasePlSqlParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BasePlSqlParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterSequence_name is called when production sequence_name is entered.
func (s *BasePlSqlParserListener) EnterSequence_name(ctx *Sequence_nameContext) {}

// ExitSequence_name is called when production sequence_name is exited.
func (s *BasePlSqlParserListener) ExitSequence_name(ctx *Sequence_nameContext) {}

// EnterException_name is called when production exception_name is entered.
func (s *BasePlSqlParserListener) EnterException_name(ctx *Exception_nameContext) {}

// ExitException_name is called when production exception_name is exited.
func (s *BasePlSqlParserListener) ExitException_name(ctx *Exception_nameContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BasePlSqlParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BasePlSqlParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterProcedure_name is called when production procedure_name is entered.
func (s *BasePlSqlParserListener) EnterProcedure_name(ctx *Procedure_nameContext) {}

// ExitProcedure_name is called when production procedure_name is exited.
func (s *BasePlSqlParserListener) ExitProcedure_name(ctx *Procedure_nameContext) {}

// EnterTrigger_name is called when production trigger_name is entered.
func (s *BasePlSqlParserListener) EnterTrigger_name(ctx *Trigger_nameContext) {}

// ExitTrigger_name is called when production trigger_name is exited.
func (s *BasePlSqlParserListener) ExitTrigger_name(ctx *Trigger_nameContext) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BasePlSqlParserListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BasePlSqlParserListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BasePlSqlParserListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BasePlSqlParserListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BasePlSqlParserListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BasePlSqlParserListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterRecord_name is called when production record_name is entered.
func (s *BasePlSqlParserListener) EnterRecord_name(ctx *Record_nameContext) {}

// ExitRecord_name is called when production record_name is exited.
func (s *BasePlSqlParserListener) ExitRecord_name(ctx *Record_nameContext) {}

// EnterCollection_name is called when production collection_name is entered.
func (s *BasePlSqlParserListener) EnterCollection_name(ctx *Collection_nameContext) {}

// ExitCollection_name is called when production collection_name is exited.
func (s *BasePlSqlParserListener) ExitCollection_name(ctx *Collection_nameContext) {}

// EnterLink_name is called when production link_name is entered.
func (s *BasePlSqlParserListener) EnterLink_name(ctx *Link_nameContext) {}

// ExitLink_name is called when production link_name is exited.
func (s *BasePlSqlParserListener) ExitLink_name(ctx *Link_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BasePlSqlParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BasePlSqlParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterTableview_name is called when production tableview_name is entered.
func (s *BasePlSqlParserListener) EnterTableview_name(ctx *Tableview_nameContext) {}

// ExitTableview_name is called when production tableview_name is exited.
func (s *BasePlSqlParserListener) ExitTableview_name(ctx *Tableview_nameContext) {}

// EnterChar_set_name is called when production char_set_name is entered.
func (s *BasePlSqlParserListener) EnterChar_set_name(ctx *Char_set_nameContext) {}

// ExitChar_set_name is called when production char_set_name is exited.
func (s *BasePlSqlParserListener) ExitChar_set_name(ctx *Char_set_nameContext) {}

// EnterSynonym_name is called when production synonym_name is entered.
func (s *BasePlSqlParserListener) EnterSynonym_name(ctx *Synonym_nameContext) {}

// ExitSynonym_name is called when production synonym_name is exited.
func (s *BasePlSqlParserListener) ExitSynonym_name(ctx *Synonym_nameContext) {}

// EnterSchema_object_name is called when production schema_object_name is entered.
func (s *BasePlSqlParserListener) EnterSchema_object_name(ctx *Schema_object_nameContext) {}

// ExitSchema_object_name is called when production schema_object_name is exited.
func (s *BasePlSqlParserListener) ExitSchema_object_name(ctx *Schema_object_nameContext) {}

// EnterKeep_clause is called when production keep_clause is entered.
func (s *BasePlSqlParserListener) EnterKeep_clause(ctx *Keep_clauseContext) {}

// ExitKeep_clause is called when production keep_clause is exited.
func (s *BasePlSqlParserListener) ExitKeep_clause(ctx *Keep_clauseContext) {}

// EnterFunction_argument is called when production function_argument is entered.
func (s *BasePlSqlParserListener) EnterFunction_argument(ctx *Function_argumentContext) {}

// ExitFunction_argument is called when production function_argument is exited.
func (s *BasePlSqlParserListener) ExitFunction_argument(ctx *Function_argumentContext) {}

// EnterFunction_argument_analytic is called when production function_argument_analytic is entered.
func (s *BasePlSqlParserListener) EnterFunction_argument_analytic(ctx *Function_argument_analyticContext) {
}

// ExitFunction_argument_analytic is called when production function_argument_analytic is exited.
func (s *BasePlSqlParserListener) ExitFunction_argument_analytic(ctx *Function_argument_analyticContext) {
}

// EnterFunction_argument_modeling is called when production function_argument_modeling is entered.
func (s *BasePlSqlParserListener) EnterFunction_argument_modeling(ctx *Function_argument_modelingContext) {
}

// ExitFunction_argument_modeling is called when production function_argument_modeling is exited.
func (s *BasePlSqlParserListener) ExitFunction_argument_modeling(ctx *Function_argument_modelingContext) {
}

// EnterRespect_or_ignore_nulls is called when production respect_or_ignore_nulls is entered.
func (s *BasePlSqlParserListener) EnterRespect_or_ignore_nulls(ctx *Respect_or_ignore_nullsContext) {}

// ExitRespect_or_ignore_nulls is called when production respect_or_ignore_nulls is exited.
func (s *BasePlSqlParserListener) ExitRespect_or_ignore_nulls(ctx *Respect_or_ignore_nullsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasePlSqlParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasePlSqlParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterType_spec is called when production type_spec is entered.
func (s *BasePlSqlParserListener) EnterType_spec(ctx *Type_specContext) {}

// ExitType_spec is called when production type_spec is exited.
func (s *BasePlSqlParserListener) ExitType_spec(ctx *Type_specContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BasePlSqlParserListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BasePlSqlParserListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterPrecision_part is called when production precision_part is entered.
func (s *BasePlSqlParserListener) EnterPrecision_part(ctx *Precision_partContext) {}

// ExitPrecision_part is called when production precision_part is exited.
func (s *BasePlSqlParserListener) ExitPrecision_part(ctx *Precision_partContext) {}

// EnterNative_datatype_element is called when production native_datatype_element is entered.
func (s *BasePlSqlParserListener) EnterNative_datatype_element(ctx *Native_datatype_elementContext) {}

// ExitNative_datatype_element is called when production native_datatype_element is exited.
func (s *BasePlSqlParserListener) ExitNative_datatype_element(ctx *Native_datatype_elementContext) {}

// EnterBind_variable is called when production bind_variable is entered.
func (s *BasePlSqlParserListener) EnterBind_variable(ctx *Bind_variableContext) {}

// ExitBind_variable is called when production bind_variable is exited.
func (s *BasePlSqlParserListener) ExitBind_variable(ctx *Bind_variableContext) {}

// EnterGeneral_element is called when production general_element is entered.
func (s *BasePlSqlParserListener) EnterGeneral_element(ctx *General_elementContext) {}

// ExitGeneral_element is called when production general_element is exited.
func (s *BasePlSqlParserListener) ExitGeneral_element(ctx *General_elementContext) {}

// EnterGeneral_element_part is called when production general_element_part is entered.
func (s *BasePlSqlParserListener) EnterGeneral_element_part(ctx *General_element_partContext) {}

// ExitGeneral_element_part is called when production general_element_part is exited.
func (s *BasePlSqlParserListener) ExitGeneral_element_part(ctx *General_element_partContext) {}

// EnterTable_element is called when production table_element is entered.
func (s *BasePlSqlParserListener) EnterTable_element(ctx *Table_elementContext) {}

// ExitTable_element is called when production table_element is exited.
func (s *BasePlSqlParserListener) ExitTable_element(ctx *Table_elementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasePlSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasePlSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumeric is called when production numeric is entered.
func (s *BasePlSqlParserListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BasePlSqlParserListener) ExitNumeric(ctx *NumericContext) {}

// EnterNumeric_negative is called when production numeric_negative is entered.
func (s *BasePlSqlParserListener) EnterNumeric_negative(ctx *Numeric_negativeContext) {}

// ExitNumeric_negative is called when production numeric_negative is exited.
func (s *BasePlSqlParserListener) ExitNumeric_negative(ctx *Numeric_negativeContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BasePlSqlParserListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BasePlSqlParserListener) ExitQuoted_string(ctx *Quoted_stringContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePlSqlParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePlSqlParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterId_expression is called when production id_expression is entered.
func (s *BasePlSqlParserListener) EnterId_expression(ctx *Id_expressionContext) {}

// ExitId_expression is called when production id_expression is exited.
func (s *BasePlSqlParserListener) ExitId_expression(ctx *Id_expressionContext) {}

// EnterOuter_join_sign is called when production outer_join_sign is entered.
func (s *BasePlSqlParserListener) EnterOuter_join_sign(ctx *Outer_join_signContext) {}

// ExitOuter_join_sign is called when production outer_join_sign is exited.
func (s *BasePlSqlParserListener) ExitOuter_join_sign(ctx *Outer_join_signContext) {}

// EnterRegular_id is called when production regular_id is entered.
func (s *BasePlSqlParserListener) EnterRegular_id(ctx *Regular_idContext) {}

// ExitRegular_id is called when production regular_id is exited.
func (s *BasePlSqlParserListener) ExitRegular_id(ctx *Regular_idContext) {}

// EnterString_function_name is called when production string_function_name is entered.
func (s *BasePlSqlParserListener) EnterString_function_name(ctx *String_function_nameContext) {}

// ExitString_function_name is called when production string_function_name is exited.
func (s *BasePlSqlParserListener) ExitString_function_name(ctx *String_function_nameContext) {}

// EnterNumeric_function_name is called when production numeric_function_name is entered.
func (s *BasePlSqlParserListener) EnterNumeric_function_name(ctx *Numeric_function_nameContext) {}

// ExitNumeric_function_name is called when production numeric_function_name is exited.
func (s *BasePlSqlParserListener) ExitNumeric_function_name(ctx *Numeric_function_nameContext) {}
