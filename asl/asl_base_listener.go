// Code generated from ASL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asl // ASL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseASLListener is a complete listener for a parse tree produced by ASLParser.
type BaseASLListener struct{}

var _ ASLListener = &BaseASLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseASLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseASLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseASLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseASLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAsl is called when production asl is entered.
func (s *BaseASLListener) EnterAsl(ctx *AslContext) {}

// ExitAsl is called when production asl is exited.
func (s *BaseASLListener) ExitAsl(ctx *AslContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseASLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseASLListener) ExitStatement(ctx *StatementContext) {}

// EnterSimple_statement is called when production simple_statement is entered.
func (s *BaseASLListener) EnterSimple_statement(ctx *Simple_statementContext) {}

// ExitSimple_statement is called when production simple_statement is exited.
func (s *BaseASLListener) ExitSimple_statement(ctx *Simple_statementContext) {}

// EnterSequential_logic_statement is called when production sequential_logic_statement is entered.
func (s *BaseASLListener) EnterSequential_logic_statement(ctx *Sequential_logic_statementContext) {}

// ExitSequential_logic_statement is called when production sequential_logic_statement is exited.
func (s *BaseASLListener) ExitSequential_logic_statement(ctx *Sequential_logic_statementContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BaseASLListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BaseASLListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterConstant_assignment is called when production constant_assignment is entered.
func (s *BaseASLListener) EnterConstant_assignment(ctx *Constant_assignmentContext) {}

// ExitConstant_assignment is called when production constant_assignment is exited.
func (s *BaseASLListener) ExitConstant_assignment(ctx *Constant_assignmentContext) {}

// EnterDate_time_assignment is called when production date_time_assignment is entered.
func (s *BaseASLListener) EnterDate_time_assignment(ctx *Date_time_assignmentContext) {}

// ExitDate_time_assignment is called when production date_time_assignment is exited.
func (s *BaseASLListener) ExitDate_time_assignment(ctx *Date_time_assignmentContext) {}

// EnterVariable_assignment is called when production variable_assignment is entered.
func (s *BaseASLListener) EnterVariable_assignment(ctx *Variable_assignmentContext) {}

// ExitVariable_assignment is called when production variable_assignment is exited.
func (s *BaseASLListener) ExitVariable_assignment(ctx *Variable_assignmentContext) {}

// EnterObject_attribute_assignment is called when production object_attribute_assignment is entered.
func (s *BaseASLListener) EnterObject_attribute_assignment(ctx *Object_attribute_assignmentContext) {}

// ExitObject_attribute_assignment is called when production object_attribute_assignment is exited.
func (s *BaseASLListener) ExitObject_attribute_assignment(ctx *Object_attribute_assignmentContext) {}

// EnterArithmetic_assignment is called when production arithmetic_assignment is entered.
func (s *BaseASLListener) EnterArithmetic_assignment(ctx *Arithmetic_assignmentContext) {}

// ExitArithmetic_assignment is called when production arithmetic_assignment is exited.
func (s *BaseASLListener) ExitArithmetic_assignment(ctx *Arithmetic_assignmentContext) {}

// EnterArithmetic_expression is called when production arithmetic_expression is entered.
func (s *BaseASLListener) EnterArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// ExitArithmetic_expression is called when production arithmetic_expression is exited.
func (s *BaseASLListener) ExitArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// EnterArithmetic_component is called when production arithmetic_component is entered.
func (s *BaseASLListener) EnterArithmetic_component(ctx *Arithmetic_componentContext) {}

// ExitArithmetic_component is called when production arithmetic_component is exited.
func (s *BaseASLListener) ExitArithmetic_component(ctx *Arithmetic_componentContext) {}

// EnterCountof_function is called when production countof_function is entered.
func (s *BaseASLListener) EnterCountof_function(ctx *Countof_functionContext) {}

// ExitCountof_function is called when production countof_function is exited.
func (s *BaseASLListener) ExitCountof_function(ctx *Countof_functionContext) {}

// EnterCreate_statement is called when production create_statement is entered.
func (s *BaseASLListener) EnterCreate_statement(ctx *Create_statementContext) {}

// ExitCreate_statement is called when production create_statement is exited.
func (s *BaseASLListener) ExitCreate_statement(ctx *Create_statementContext) {}

// EnterAttribute_assignments is called when production attribute_assignments is entered.
func (s *BaseASLListener) EnterAttribute_assignments(ctx *Attribute_assignmentsContext) {}

// ExitAttribute_assignments is called when production attribute_assignments is exited.
func (s *BaseASLListener) ExitAttribute_assignments(ctx *Attribute_assignmentsContext) {}

// EnterAttribute_assignment is called when production attribute_assignment is entered.
func (s *BaseASLListener) EnterAttribute_assignment(ctx *Attribute_assignmentContext) {}

// ExitAttribute_assignment is called when production attribute_assignment is exited.
func (s *BaseASLListener) ExitAttribute_assignment(ctx *Attribute_assignmentContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BaseASLListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BaseASLListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterFind_statement is called when production find_statement is entered.
func (s *BaseASLListener) EnterFind_statement(ctx *Find_statementContext) {}

// ExitFind_statement is called when production find_statement is exited.
func (s *BaseASLListener) ExitFind_statement(ctx *Find_statementContext) {}

// EnterFind_instance_statement is called when production find_instance_statement is entered.
func (s *BaseASLListener) EnterFind_instance_statement(ctx *Find_instance_statementContext) {}

// ExitFind_instance_statement is called when production find_instance_statement is exited.
func (s *BaseASLListener) ExitFind_instance_statement(ctx *Find_instance_statementContext) {}

// EnterFind_set_statement is called when production find_set_statement is entered.
func (s *BaseASLListener) EnterFind_set_statement(ctx *Find_set_statementContext) {}

// ExitFind_set_statement is called when production find_set_statement is exited.
func (s *BaseASLListener) ExitFind_set_statement(ctx *Find_set_statementContext) {}

// EnterObject_condition is called when production object_condition is entered.
func (s *BaseASLListener) EnterObject_condition(ctx *Object_conditionContext) {}

// ExitObject_condition is called when production object_condition is exited.
func (s *BaseASLListener) ExitObject_condition(ctx *Object_conditionContext) {}

// EnterEquality_operator is called when production equality_operator is entered.
func (s *BaseASLListener) EnterEquality_operator(ctx *Equality_operatorContext) {}

// ExitEquality_operator is called when production equality_operator is exited.
func (s *BaseASLListener) ExitEquality_operator(ctx *Equality_operatorContext) {}

// EnterLogical_operator is called when production logical_operator is entered.
func (s *BaseASLListener) EnterLogical_operator(ctx *Logical_operatorContext) {}

// ExitLogical_operator is called when production logical_operator is exited.
func (s *BaseASLListener) ExitLogical_operator(ctx *Logical_operatorContext) {}

// EnterRelationship_navigation is called when production relationship_navigation is entered.
func (s *BaseASLListener) EnterRelationship_navigation(ctx *Relationship_navigationContext) {}

// ExitRelationship_navigation is called when production relationship_navigation is exited.
func (s *BaseASLListener) ExitRelationship_navigation(ctx *Relationship_navigationContext) {}

// EnterAssociative_relationship_navigation is called when production associative_relationship_navigation is entered.
func (s *BaseASLListener) EnterAssociative_relationship_navigation(ctx *Associative_relationship_navigationContext) {
}

// ExitAssociative_relationship_navigation is called when production associative_relationship_navigation is exited.
func (s *BaseASLListener) ExitAssociative_relationship_navigation(ctx *Associative_relationship_navigationContext) {
}

// EnterRelationship_link is called when production relationship_link is entered.
func (s *BaseASLListener) EnterRelationship_link(ctx *Relationship_linkContext) {}

// ExitRelationship_link is called when production relationship_link is exited.
func (s *BaseASLListener) ExitRelationship_link(ctx *Relationship_linkContext) {}

// EnterRelationship_unlink is called when production relationship_unlink is entered.
func (s *BaseASLListener) EnterRelationship_unlink(ctx *Relationship_unlinkContext) {}

// ExitRelationship_unlink is called when production relationship_unlink is exited.
func (s *BaseASLListener) ExitRelationship_unlink(ctx *Relationship_unlinkContext) {}

// EnterRelationship_spec is called when production relationship_spec is entered.
func (s *BaseASLListener) EnterRelationship_spec(ctx *Relationship_specContext) {}

// ExitRelationship_spec is called when production relationship_spec is exited.
func (s *BaseASLListener) ExitRelationship_spec(ctx *Relationship_specContext) {}

// EnterQualified_relationship is called when production qualified_relationship is entered.
func (s *BaseASLListener) EnterQualified_relationship(ctx *Qualified_relationshipContext) {}

// ExitQualified_relationship is called when production qualified_relationship is exited.
func (s *BaseASLListener) ExitQualified_relationship(ctx *Qualified_relationshipContext) {}

// EnterRelationship_role is called when production relationship_role is entered.
func (s *BaseASLListener) EnterRelationship_role(ctx *Relationship_roleContext) {}

// ExitRelationship_role is called when production relationship_role is exited.
func (s *BaseASLListener) ExitRelationship_role(ctx *Relationship_roleContext) {}

// EnterEvent_generation is called when production event_generation is entered.
func (s *BaseASLListener) EnterEvent_generation(ctx *Event_generationContext) {}

// ExitEvent_generation is called when production event_generation is exited.
func (s *BaseASLListener) ExitEvent_generation(ctx *Event_generationContext) {}

// EnterInternal_event_generation is called when production internal_event_generation is entered.
func (s *BaseASLListener) EnterInternal_event_generation(ctx *Internal_event_generationContext) {}

// ExitInternal_event_generation is called when production internal_event_generation is exited.
func (s *BaseASLListener) ExitInternal_event_generation(ctx *Internal_event_generationContext) {}

// EnterExternal_event_generation is called when production external_event_generation is entered.
func (s *BaseASLListener) EnterExternal_event_generation(ctx *External_event_generationContext) {}

// ExitExternal_event_generation is called when production external_event_generation is exited.
func (s *BaseASLListener) ExitExternal_event_generation(ctx *External_event_generationContext) {}

// EnterEvent_specification is called when production event_specification is entered.
func (s *BaseASLListener) EnterEvent_specification(ctx *Event_specificationContext) {}

// ExitEvent_specification is called when production event_specification is exited.
func (s *BaseASLListener) ExitEvent_specification(ctx *Event_specificationContext) {}

// EnterExternal_event_specification is called when production external_event_specification is entered.
func (s *BaseASLListener) EnterExternal_event_specification(ctx *External_event_specificationContext) {
}

// ExitExternal_event_specification is called when production external_event_specification is exited.
func (s *BaseASLListener) ExitExternal_event_specification(ctx *External_event_specificationContext) {
}

// EnterEvent_parameters is called when production event_parameters is entered.
func (s *BaseASLListener) EnterEvent_parameters(ctx *Event_parametersContext) {}

// ExitEvent_parameters is called when production event_parameters is exited.
func (s *BaseASLListener) ExitEvent_parameters(ctx *Event_parametersContext) {}

// EnterEvent_parameter is called when production event_parameter is entered.
func (s *BaseASLListener) EnterEvent_parameter(ctx *Event_parameterContext) {}

// ExitEvent_parameter is called when production event_parameter is exited.
func (s *BaseASLListener) ExitEvent_parameter(ctx *Event_parameterContext) {}

// EnterOperation_invocation is called when production operation_invocation is entered.
func (s *BaseASLListener) EnterOperation_invocation(ctx *Operation_invocationContext) {}

// ExitOperation_invocation is called when production operation_invocation is exited.
func (s *BaseASLListener) ExitOperation_invocation(ctx *Operation_invocationContext) {}

// EnterObject_operation is called when production object_operation is entered.
func (s *BaseASLListener) EnterObject_operation(ctx *Object_operationContext) {}

// ExitObject_operation is called when production object_operation is exited.
func (s *BaseASLListener) ExitObject_operation(ctx *Object_operationContext) {}

// EnterDomain_operation is called when production domain_operation is entered.
func (s *BaseASLListener) EnterDomain_operation(ctx *Domain_operationContext) {}

// ExitDomain_operation is called when production domain_operation is exited.
func (s *BaseASLListener) ExitDomain_operation(ctx *Domain_operationContext) {}

// EnterBridge_operation is called when production bridge_operation is entered.
func (s *BaseASLListener) EnterBridge_operation(ctx *Bridge_operationContext) {}

// ExitBridge_operation is called when production bridge_operation is exited.
func (s *BaseASLListener) ExitBridge_operation(ctx *Bridge_operationContext) {}

// EnterOo_specification is called when production oo_specification is entered.
func (s *BaseASLListener) EnterOo_specification(ctx *Oo_specificationContext) {}

// ExitOo_specification is called when production oo_specification is exited.
func (s *BaseASLListener) ExitOo_specification(ctx *Oo_specificationContext) {}

// EnterDo_specification is called when production do_specification is entered.
func (s *BaseASLListener) EnterDo_specification(ctx *Do_specificationContext) {}

// ExitDo_specification is called when production do_specification is exited.
func (s *BaseASLListener) ExitDo_specification(ctx *Do_specificationContext) {}

// EnterBo_specification is called when production bo_specification is entered.
func (s *BaseASLListener) EnterBo_specification(ctx *Bo_specificationContext) {}

// ExitBo_specification is called when production bo_specification is exited.
func (s *BaseASLListener) ExitBo_specification(ctx *Bo_specificationContext) {}

// EnterOutput_parameter_list is called when production output_parameter_list is entered.
func (s *BaseASLListener) EnterOutput_parameter_list(ctx *Output_parameter_listContext) {}

// ExitOutput_parameter_list is called when production output_parameter_list is exited.
func (s *BaseASLListener) ExitOutput_parameter_list(ctx *Output_parameter_listContext) {}

// EnterOutput_parameter is called when production output_parameter is entered.
func (s *BaseASLListener) EnterOutput_parameter(ctx *Output_parameterContext) {}

// ExitOutput_parameter is called when production output_parameter is exited.
func (s *BaseASLListener) ExitOutput_parameter(ctx *Output_parameterContext) {}

// EnterInput_parameter_list is called when production input_parameter_list is entered.
func (s *BaseASLListener) EnterInput_parameter_list(ctx *Input_parameter_listContext) {}

// ExitInput_parameter_list is called when production input_parameter_list is exited.
func (s *BaseASLListener) ExitInput_parameter_list(ctx *Input_parameter_listContext) {}

// EnterInput_parameter is called when production input_parameter is entered.
func (s *BaseASLListener) EnterInput_parameter(ctx *Input_parameterContext) {}

// ExitInput_parameter is called when production input_parameter is exited.
func (s *BaseASLListener) ExitInput_parameter(ctx *Input_parameterContext) {}

// EnterTimer_operation is called when production timer_operation is entered.
func (s *BaseASLListener) EnterTimer_operation(ctx *Timer_operationContext) {}

// ExitTimer_operation is called when production timer_operation is exited.
func (s *BaseASLListener) ExitTimer_operation(ctx *Timer_operationContext) {}

// EnterTimer_creation is called when production timer_creation is entered.
func (s *BaseASLListener) EnterTimer_creation(ctx *Timer_creationContext) {}

// ExitTimer_creation is called when production timer_creation is exited.
func (s *BaseASLListener) ExitTimer_creation(ctx *Timer_creationContext) {}

// EnterTimer_deletion is called when production timer_deletion is entered.
func (s *BaseASLListener) EnterTimer_deletion(ctx *Timer_deletionContext) {}

// ExitTimer_deletion is called when production timer_deletion is exited.
func (s *BaseASLListener) ExitTimer_deletion(ctx *Timer_deletionContext) {}

// EnterSet_relative_timer is called when production set_relative_timer is entered.
func (s *BaseASLListener) EnterSet_relative_timer(ctx *Set_relative_timerContext) {}

// ExitSet_relative_timer is called when production set_relative_timer is exited.
func (s *BaseASLListener) ExitSet_relative_timer(ctx *Set_relative_timerContext) {}

// EnterSet_absolute_timer is called when production set_absolute_timer is entered.
func (s *BaseASLListener) EnterSet_absolute_timer(ctx *Set_absolute_timerContext) {}

// ExitSet_absolute_timer is called when production set_absolute_timer is exited.
func (s *BaseASLListener) ExitSet_absolute_timer(ctx *Set_absolute_timerContext) {}

// EnterSet_recurring_timer is called when production set_recurring_timer is entered.
func (s *BaseASLListener) EnterSet_recurring_timer(ctx *Set_recurring_timerContext) {}

// ExitSet_recurring_timer is called when production set_recurring_timer is exited.
func (s *BaseASLListener) ExitSet_recurring_timer(ctx *Set_recurring_timerContext) {}

// EnterTimer_reset is called when production timer_reset is entered.
func (s *BaseASLListener) EnterTimer_reset(ctx *Timer_resetContext) {}

// ExitTimer_reset is called when production timer_reset is exited.
func (s *BaseASLListener) ExitTimer_reset(ctx *Timer_resetContext) {}

// EnterTimer_id is called when production timer_id is entered.
func (s *BaseASLListener) EnterTimer_id(ctx *Timer_idContext) {}

// ExitTimer_id is called when production timer_id is exited.
func (s *BaseASLListener) ExitTimer_id(ctx *Timer_idContext) {}

// EnterTime_to_fire is called when production time_to_fire is entered.
func (s *BaseASLListener) EnterTime_to_fire(ctx *Time_to_fireContext) {}

// ExitTime_to_fire is called when production time_to_fire is exited.
func (s *BaseASLListener) ExitTime_to_fire(ctx *Time_to_fireContext) {}

// EnterReturn_event is called when production return_event is entered.
func (s *BaseASLListener) EnterReturn_event(ctx *Return_eventContext) {}

// ExitReturn_event is called when production return_event is exited.
func (s *BaseASLListener) ExitReturn_event(ctx *Return_eventContext) {}

// EnterTarget_instance is called when production target_instance is entered.
func (s *BaseASLListener) EnterTarget_instance(ctx *Target_instanceContext) {}

// ExitTarget_instance is called when production target_instance is exited.
func (s *BaseASLListener) ExitTarget_instance(ctx *Target_instanceContext) {}

// EnterFire_year is called when production fire_year is entered.
func (s *BaseASLListener) EnterFire_year(ctx *Fire_yearContext) {}

// ExitFire_year is called when production fire_year is exited.
func (s *BaseASLListener) ExitFire_year(ctx *Fire_yearContext) {}

// EnterFire_month is called when production fire_month is entered.
func (s *BaseASLListener) EnterFire_month(ctx *Fire_monthContext) {}

// ExitFire_month is called when production fire_month is exited.
func (s *BaseASLListener) ExitFire_month(ctx *Fire_monthContext) {}

// EnterFire_date is called when production fire_date is entered.
func (s *BaseASLListener) EnterFire_date(ctx *Fire_dateContext) {}

// ExitFire_date is called when production fire_date is exited.
func (s *BaseASLListener) ExitFire_date(ctx *Fire_dateContext) {}

// EnterFire_hour is called when production fire_hour is entered.
func (s *BaseASLListener) EnterFire_hour(ctx *Fire_hourContext) {}

// ExitFire_hour is called when production fire_hour is exited.
func (s *BaseASLListener) ExitFire_hour(ctx *Fire_hourContext) {}

// EnterFire_minute is called when production fire_minute is entered.
func (s *BaseASLListener) EnterFire_minute(ctx *Fire_minuteContext) {}

// ExitFire_minute is called when production fire_minute is exited.
func (s *BaseASLListener) ExitFire_minute(ctx *Fire_minuteContext) {}

// EnterFire_second is called when production fire_second is entered.
func (s *BaseASLListener) EnterFire_second(ctx *Fire_secondContext) {}

// ExitFire_second is called when production fire_second is exited.
func (s *BaseASLListener) ExitFire_second(ctx *Fire_secondContext) {}

// EnterStart_time is called when production start_time is entered.
func (s *BaseASLListener) EnterStart_time(ctx *Start_timeContext) {}

// ExitStart_time is called when production start_time is exited.
func (s *BaseASLListener) ExitStart_time(ctx *Start_timeContext) {}

// EnterPeriod is called when production period is entered.
func (s *BaseASLListener) EnterPeriod(ctx *PeriodContext) {}

// ExitPeriod is called when production period is exited.
func (s *BaseASLListener) ExitPeriod(ctx *PeriodContext) {}

// EnterSet_operation is called when production set_operation is entered.
func (s *BaseASLListener) EnterSet_operation(ctx *Set_operationContext) {}

// ExitSet_operation is called when production set_operation is exited.
func (s *BaseASLListener) ExitSet_operation(ctx *Set_operationContext) {}

// EnterUnique_set is called when production unique_set is entered.
func (s *BaseASLListener) EnterUnique_set(ctx *Unique_setContext) {}

// ExitUnique_set is called when production unique_set is exited.
func (s *BaseASLListener) ExitUnique_set(ctx *Unique_setContext) {}

// EnterSet_combination is called when production set_combination is entered.
func (s *BaseASLListener) EnterSet_combination(ctx *Set_combinationContext) {}

// ExitSet_combination is called when production set_combination is exited.
func (s *BaseASLListener) ExitSet_combination(ctx *Set_combinationContext) {}

// EnterSet_difference is called when production set_difference is entered.
func (s *BaseASLListener) EnterSet_difference(ctx *Set_differenceContext) {}

// ExitSet_difference is called when production set_difference is exited.
func (s *BaseASLListener) ExitSet_difference(ctx *Set_differenceContext) {}

// EnterStruct_statement is called when production struct_statement is entered.
func (s *BaseASLListener) EnterStruct_statement(ctx *Struct_statementContext) {}

// ExitStruct_statement is called when production struct_statement is exited.
func (s *BaseASLListener) ExitStruct_statement(ctx *Struct_statementContext) {}

// EnterStruct_definition is called when production struct_definition is entered.
func (s *BaseASLListener) EnterStruct_definition(ctx *Struct_definitionContext) {}

// ExitStruct_definition is called when production struct_definition is exited.
func (s *BaseASLListener) ExitStruct_definition(ctx *Struct_definitionContext) {}

// EnterMember_name is called when production member_name is entered.
func (s *BaseASLListener) EnterMember_name(ctx *Member_nameContext) {}

// ExitMember_name is called when production member_name is exited.
func (s *BaseASLListener) ExitMember_name(ctx *Member_nameContext) {}

// EnterMember_type is called when production member_type is entered.
func (s *BaseASLListener) EnterMember_type(ctx *Member_typeContext) {}

// ExitMember_type is called when production member_type is exited.
func (s *BaseASLListener) ExitMember_type(ctx *Member_typeContext) {}

// EnterStruct_instantiation is called when production struct_instantiation is entered.
func (s *BaseASLListener) EnterStruct_instantiation(ctx *Struct_instantiationContext) {}

// ExitStruct_instantiation is called when production struct_instantiation is exited.
func (s *BaseASLListener) ExitStruct_instantiation(ctx *Struct_instantiationContext) {}

// EnterStruct_assembly is called when production struct_assembly is entered.
func (s *BaseASLListener) EnterStruct_assembly(ctx *Struct_assemblyContext) {}

// ExitStruct_assembly is called when production struct_assembly is exited.
func (s *BaseASLListener) ExitStruct_assembly(ctx *Struct_assemblyContext) {}

// EnterStruct_for_loop is called when production struct_for_loop is entered.
func (s *BaseASLListener) EnterStruct_for_loop(ctx *Struct_for_loopContext) {}

// ExitStruct_for_loop is called when production struct_for_loop is exited.
func (s *BaseASLListener) ExitStruct_for_loop(ctx *Struct_for_loopContext) {}

// EnterOrder_struct is called when production order_struct is entered.
func (s *BaseASLListener) EnterOrder_struct(ctx *Order_structContext) {}

// ExitOrder_struct is called when production order_struct is exited.
func (s *BaseASLListener) ExitOrder_struct(ctx *Order_structContext) {}

// EnterStruct_subset is called when production struct_subset is entered.
func (s *BaseASLListener) EnterStruct_subset(ctx *Struct_subsetContext) {}

// ExitStruct_subset is called when production struct_subset is exited.
func (s *BaseASLListener) ExitStruct_subset(ctx *Struct_subsetContext) {}

// EnterStruct_condition is called when production struct_condition is entered.
func (s *BaseASLListener) EnterStruct_condition(ctx *Struct_conditionContext) {}

// ExitStruct_condition is called when production struct_condition is exited.
func (s *BaseASLListener) ExitStruct_condition(ctx *Struct_conditionContext) {}

// EnterStruct_ is called when production struct_ is entered.
func (s *BaseASLListener) EnterStruct_(ctx *Struct_Context) {}

// ExitStruct_ is called when production struct_ is exited.
func (s *BaseASLListener) ExitStruct_(ctx *Struct_Context) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseASLListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseASLListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterValue_list is called when production value_list is entered.
func (s *BaseASLListener) EnterValue_list(ctx *Value_listContext) {}

// ExitValue_list is called when production value_list is exited.
func (s *BaseASLListener) ExitValue_list(ctx *Value_listContext) {}

// EnterStruct_value is called when production struct_value is entered.
func (s *BaseASLListener) EnterStruct_value(ctx *Struct_valueContext) {}

// ExitStruct_value is called when production struct_value is exited.
func (s *BaseASLListener) ExitStruct_value(ctx *Struct_valueContext) {}

// EnterVariable_list is called when production variable_list is entered.
func (s *BaseASLListener) EnterVariable_list(ctx *Variable_listContext) {}

// ExitVariable_list is called when production variable_list is exited.
func (s *BaseASLListener) ExitVariable_list(ctx *Variable_listContext) {}

// EnterSwitch_statement is called when production switch_statement is entered.
func (s *BaseASLListener) EnterSwitch_statement(ctx *Switch_statementContext) {}

// ExitSwitch_statement is called when production switch_statement is exited.
func (s *BaseASLListener) ExitSwitch_statement(ctx *Switch_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseASLListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseASLListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterLogical_condition is called when production logical_condition is entered.
func (s *BaseASLListener) EnterLogical_condition(ctx *Logical_conditionContext) {}

// ExitLogical_condition is called when production logical_condition is exited.
func (s *BaseASLListener) ExitLogical_condition(ctx *Logical_conditionContext) {}

// EnterCompound_logical_condition is called when production compound_logical_condition is entered.
func (s *BaseASLListener) EnterCompound_logical_condition(ctx *Compound_logical_conditionContext) {}

// ExitCompound_logical_condition is called when production compound_logical_condition is exited.
func (s *BaseASLListener) ExitCompound_logical_condition(ctx *Compound_logical_conditionContext) {}

// EnterSimple_logical_condition is called when production simple_logical_condition is entered.
func (s *BaseASLListener) EnterSimple_logical_condition(ctx *Simple_logical_conditionContext) {}

// ExitSimple_logical_condition is called when production simple_logical_condition is exited.
func (s *BaseASLListener) ExitSimple_logical_condition(ctx *Simple_logical_conditionContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseASLListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseASLListener) ExitComponent(ctx *ComponentContext) {}

// EnterFor_loop is called when production for_loop is entered.
func (s *BaseASLListener) EnterFor_loop(ctx *For_loopContext) {}

// ExitFor_loop is called when production for_loop is exited.
func (s *BaseASLListener) ExitFor_loop(ctx *For_loopContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BaseASLListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BaseASLListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseASLListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseASLListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterNative_language_section is called when production native_language_section is entered.
func (s *BaseASLListener) EnterNative_language_section(ctx *Native_language_sectionContext) {}

// ExitNative_language_section is called when production native_language_section is exited.
func (s *BaseASLListener) ExitNative_language_section(ctx *Native_language_sectionContext) {}

// EnterSet_variable is called when production set_variable is entered.
func (s *BaseASLListener) EnterSet_variable(ctx *Set_variableContext) {}

// ExitSet_variable is called when production set_variable is exited.
func (s *BaseASLListener) ExitSet_variable(ctx *Set_variableContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseASLListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseASLListener) ExitVariable(ctx *VariableContext) {}

// EnterObject_name is called when production object_name is entered.
func (s *BaseASLListener) EnterObject_name(ctx *Object_nameContext) {}

// ExitObject_name is called when production object_name is exited.
func (s *BaseASLListener) ExitObject_name(ctx *Object_nameContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseASLListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseASLListener) ExitAttribute(ctx *AttributeContext) {}

// EnterObject_attribute is called when production object_attribute is entered.
func (s *BaseASLListener) EnterObject_attribute(ctx *Object_attributeContext) {}

// ExitObject_attribute is called when production object_attribute is exited.
func (s *BaseASLListener) ExitObject_attribute(ctx *Object_attributeContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseASLListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseASLListener) ExitConstant(ctx *ConstantContext) {}

// EnterText_string is called when production text_string is entered.
func (s *BaseASLListener) EnterText_string(ctx *Text_stringContext) {}

// ExitText_string is called when production text_string is exited.
func (s *BaseASLListener) ExitText_string(ctx *Text_stringContext) {}

// EnterEnum_value is called when production enum_value is entered.
func (s *BaseASLListener) EnterEnum_value(ctx *Enum_valueContext) {}

// ExitEnum_value is called when production enum_value is exited.
func (s *BaseASLListener) ExitEnum_value(ctx *Enum_valueContext) {}

// EnterStructure_field is called when production structure_field is entered.
func (s *BaseASLListener) EnterStructure_field(ctx *Structure_fieldContext) {}

// ExitStructure_field is called when production structure_field is exited.
func (s *BaseASLListener) ExitStructure_field(ctx *Structure_fieldContext) {}

// EnterUser_defined_type is called when production user_defined_type is entered.
func (s *BaseASLListener) EnterUser_defined_type(ctx *User_defined_typeContext) {}

// ExitUser_defined_type is called when production user_defined_type is exited.
func (s *BaseASLListener) ExitUser_defined_type(ctx *User_defined_typeContext) {}
