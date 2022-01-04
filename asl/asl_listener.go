// Code generated from ASL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asl // ASL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ASLListener is a complete listener for a parse tree produced by ASLParser.
type ASLListener interface {
	antlr.ParseTreeListener

	// EnterAsl is called when entering the asl production.
	EnterAsl(c *AslContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimple_statement is called when entering the simple_statement production.
	EnterSimple_statement(c *Simple_statementContext)

	// EnterSequential_logic_statement is called when entering the sequential_logic_statement production.
	EnterSequential_logic_statement(c *Sequential_logic_statementContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterConstant_assignment is called when entering the constant_assignment production.
	EnterConstant_assignment(c *Constant_assignmentContext)

	// EnterDate_time_assignment is called when entering the date_time_assignment production.
	EnterDate_time_assignment(c *Date_time_assignmentContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterObject_attribute_assignment is called when entering the object_attribute_assignment production.
	EnterObject_attribute_assignment(c *Object_attribute_assignmentContext)

	// EnterArithmetic_assignment is called when entering the arithmetic_assignment production.
	EnterArithmetic_assignment(c *Arithmetic_assignmentContext)

	// EnterArithmetic_expression is called when entering the arithmetic_expression production.
	EnterArithmetic_expression(c *Arithmetic_expressionContext)

	// EnterArithmetic_component is called when entering the arithmetic_component production.
	EnterArithmetic_component(c *Arithmetic_componentContext)

	// EnterCountof_function is called when entering the countof_function production.
	EnterCountof_function(c *Countof_functionContext)

	// EnterCreate_statement is called when entering the create_statement production.
	EnterCreate_statement(c *Create_statementContext)

	// EnterAttribute_assignments is called when entering the attribute_assignments production.
	EnterAttribute_assignments(c *Attribute_assignmentsContext)

	// EnterAttribute_assignment is called when entering the attribute_assignment production.
	EnterAttribute_assignment(c *Attribute_assignmentContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterFind_statement is called when entering the find_statement production.
	EnterFind_statement(c *Find_statementContext)

	// EnterFind_instance_statement is called when entering the find_instance_statement production.
	EnterFind_instance_statement(c *Find_instance_statementContext)

	// EnterFind_set_statement is called when entering the find_set_statement production.
	EnterFind_set_statement(c *Find_set_statementContext)

	// EnterObject_condition is called when entering the object_condition production.
	EnterObject_condition(c *Object_conditionContext)

	// EnterEquality_operator is called when entering the equality_operator production.
	EnterEquality_operator(c *Equality_operatorContext)

	// EnterLogical_operator is called when entering the logical_operator production.
	EnterLogical_operator(c *Logical_operatorContext)

	// EnterRelationship_navigation is called when entering the relationship_navigation production.
	EnterRelationship_navigation(c *Relationship_navigationContext)

	// EnterAssociative_relationship_navigation is called when entering the associative_relationship_navigation production.
	EnterAssociative_relationship_navigation(c *Associative_relationship_navigationContext)

	// EnterRelationship_link is called when entering the relationship_link production.
	EnterRelationship_link(c *Relationship_linkContext)

	// EnterRelationship_unlink is called when entering the relationship_unlink production.
	EnterRelationship_unlink(c *Relationship_unlinkContext)

	// EnterRelationship_spec is called when entering the relationship_spec production.
	EnterRelationship_spec(c *Relationship_specContext)

	// EnterQualified_relationship is called when entering the qualified_relationship production.
	EnterQualified_relationship(c *Qualified_relationshipContext)

	// EnterRelationship_role is called when entering the relationship_role production.
	EnterRelationship_role(c *Relationship_roleContext)

	// EnterEvent_generation is called when entering the event_generation production.
	EnterEvent_generation(c *Event_generationContext)

	// EnterInternal_event_generation is called when entering the internal_event_generation production.
	EnterInternal_event_generation(c *Internal_event_generationContext)

	// EnterExternal_event_generation is called when entering the external_event_generation production.
	EnterExternal_event_generation(c *External_event_generationContext)

	// EnterEvent_specification is called when entering the event_specification production.
	EnterEvent_specification(c *Event_specificationContext)

	// EnterExternal_event_specification is called when entering the external_event_specification production.
	EnterExternal_event_specification(c *External_event_specificationContext)

	// EnterEvent_parameters is called when entering the event_parameters production.
	EnterEvent_parameters(c *Event_parametersContext)

	// EnterEvent_parameter is called when entering the event_parameter production.
	EnterEvent_parameter(c *Event_parameterContext)

	// EnterOperation_invocation is called when entering the operation_invocation production.
	EnterOperation_invocation(c *Operation_invocationContext)

	// EnterObject_operation is called when entering the object_operation production.
	EnterObject_operation(c *Object_operationContext)

	// EnterDomain_operation is called when entering the domain_operation production.
	EnterDomain_operation(c *Domain_operationContext)

	// EnterBridge_operation is called when entering the bridge_operation production.
	EnterBridge_operation(c *Bridge_operationContext)

	// EnterOo_specification is called when entering the oo_specification production.
	EnterOo_specification(c *Oo_specificationContext)

	// EnterDo_specification is called when entering the do_specification production.
	EnterDo_specification(c *Do_specificationContext)

	// EnterBo_specification is called when entering the bo_specification production.
	EnterBo_specification(c *Bo_specificationContext)

	// EnterOutput_parameter_list is called when entering the output_parameter_list production.
	EnterOutput_parameter_list(c *Output_parameter_listContext)

	// EnterOutput_parameter is called when entering the output_parameter production.
	EnterOutput_parameter(c *Output_parameterContext)

	// EnterInput_parameter_list is called when entering the input_parameter_list production.
	EnterInput_parameter_list(c *Input_parameter_listContext)

	// EnterInput_parameter is called when entering the input_parameter production.
	EnterInput_parameter(c *Input_parameterContext)

	// EnterTimer_operation is called when entering the timer_operation production.
	EnterTimer_operation(c *Timer_operationContext)

	// EnterTimer_creation is called when entering the timer_creation production.
	EnterTimer_creation(c *Timer_creationContext)

	// EnterTimer_deletion is called when entering the timer_deletion production.
	EnterTimer_deletion(c *Timer_deletionContext)

	// EnterSet_relative_timer is called when entering the set_relative_timer production.
	EnterSet_relative_timer(c *Set_relative_timerContext)

	// EnterSet_absolute_timer is called when entering the set_absolute_timer production.
	EnterSet_absolute_timer(c *Set_absolute_timerContext)

	// EnterSet_recurring_timer is called when entering the set_recurring_timer production.
	EnterSet_recurring_timer(c *Set_recurring_timerContext)

	// EnterTimer_reset is called when entering the timer_reset production.
	EnterTimer_reset(c *Timer_resetContext)

	// EnterTimer_id is called when entering the timer_id production.
	EnterTimer_id(c *Timer_idContext)

	// EnterTime_to_fire is called when entering the time_to_fire production.
	EnterTime_to_fire(c *Time_to_fireContext)

	// EnterReturn_event is called when entering the return_event production.
	EnterReturn_event(c *Return_eventContext)

	// EnterTarget_instance is called when entering the target_instance production.
	EnterTarget_instance(c *Target_instanceContext)

	// EnterFire_year is called when entering the fire_year production.
	EnterFire_year(c *Fire_yearContext)

	// EnterFire_month is called when entering the fire_month production.
	EnterFire_month(c *Fire_monthContext)

	// EnterFire_date is called when entering the fire_date production.
	EnterFire_date(c *Fire_dateContext)

	// EnterFire_hour is called when entering the fire_hour production.
	EnterFire_hour(c *Fire_hourContext)

	// EnterFire_minute is called when entering the fire_minute production.
	EnterFire_minute(c *Fire_minuteContext)

	// EnterFire_second is called when entering the fire_second production.
	EnterFire_second(c *Fire_secondContext)

	// EnterStart_time is called when entering the start_time production.
	EnterStart_time(c *Start_timeContext)

	// EnterPeriod is called when entering the period production.
	EnterPeriod(c *PeriodContext)

	// EnterSet_operation is called when entering the set_operation production.
	EnterSet_operation(c *Set_operationContext)

	// EnterUnique_set is called when entering the unique_set production.
	EnterUnique_set(c *Unique_setContext)

	// EnterSet_combination is called when entering the set_combination production.
	EnterSet_combination(c *Set_combinationContext)

	// EnterSet_difference is called when entering the set_difference production.
	EnterSet_difference(c *Set_differenceContext)

	// EnterStruct_statement is called when entering the struct_statement production.
	EnterStruct_statement(c *Struct_statementContext)

	// EnterStruct_definition is called when entering the struct_definition production.
	EnterStruct_definition(c *Struct_definitionContext)

	// EnterMember_name is called when entering the member_name production.
	EnterMember_name(c *Member_nameContext)

	// EnterMember_type is called when entering the member_type production.
	EnterMember_type(c *Member_typeContext)

	// EnterStruct_instantiation is called when entering the struct_instantiation production.
	EnterStruct_instantiation(c *Struct_instantiationContext)

	// EnterStruct_assembly is called when entering the struct_assembly production.
	EnterStruct_assembly(c *Struct_assemblyContext)

	// EnterStruct_for_loop is called when entering the struct_for_loop production.
	EnterStruct_for_loop(c *Struct_for_loopContext)

	// EnterOrder_struct is called when entering the order_struct production.
	EnterOrder_struct(c *Order_structContext)

	// EnterStruct_subset is called when entering the struct_subset production.
	EnterStruct_subset(c *Struct_subsetContext)

	// EnterStruct_condition is called when entering the struct_condition production.
	EnterStruct_condition(c *Struct_conditionContext)

	// EnterStruct_ is called when entering the struct_ production.
	EnterStruct_(c *Struct_Context)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterValue_list is called when entering the value_list production.
	EnterValue_list(c *Value_listContext)

	// EnterStruct_value is called when entering the struct_value production.
	EnterStruct_value(c *Struct_valueContext)

	// EnterVariable_list is called when entering the variable_list production.
	EnterVariable_list(c *Variable_listContext)

	// EnterSwitch_statement is called when entering the switch_statement production.
	EnterSwitch_statement(c *Switch_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterLogical_condition is called when entering the logical_condition production.
	EnterLogical_condition(c *Logical_conditionContext)

	// EnterCompound_logical_condition is called when entering the compound_logical_condition production.
	EnterCompound_logical_condition(c *Compound_logical_conditionContext)

	// EnterSimple_logical_condition is called when entering the simple_logical_condition production.
	EnterSimple_logical_condition(c *Simple_logical_conditionContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterFor_loop is called when entering the for_loop production.
	EnterFor_loop(c *For_loopContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterNative_language_section is called when entering the native_language_section production.
	EnterNative_language_section(c *Native_language_sectionContext)

	// EnterSet_variable is called when entering the set_variable production.
	EnterSet_variable(c *Set_variableContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterObject_name is called when entering the object_name production.
	EnterObject_name(c *Object_nameContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterObject_attribute is called when entering the object_attribute production.
	EnterObject_attribute(c *Object_attributeContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterText_string is called when entering the text_string production.
	EnterText_string(c *Text_stringContext)

	// EnterEnum_value is called when entering the enum_value production.
	EnterEnum_value(c *Enum_valueContext)

	// EnterStructure_field is called when entering the structure_field production.
	EnterStructure_field(c *Structure_fieldContext)

	// EnterUser_defined_type is called when entering the user_defined_type production.
	EnterUser_defined_type(c *User_defined_typeContext)

	// ExitAsl is called when exiting the asl production.
	ExitAsl(c *AslContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimple_statement is called when exiting the simple_statement production.
	ExitSimple_statement(c *Simple_statementContext)

	// ExitSequential_logic_statement is called when exiting the sequential_logic_statement production.
	ExitSequential_logic_statement(c *Sequential_logic_statementContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitConstant_assignment is called when exiting the constant_assignment production.
	ExitConstant_assignment(c *Constant_assignmentContext)

	// ExitDate_time_assignment is called when exiting the date_time_assignment production.
	ExitDate_time_assignment(c *Date_time_assignmentContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitObject_attribute_assignment is called when exiting the object_attribute_assignment production.
	ExitObject_attribute_assignment(c *Object_attribute_assignmentContext)

	// ExitArithmetic_assignment is called when exiting the arithmetic_assignment production.
	ExitArithmetic_assignment(c *Arithmetic_assignmentContext)

	// ExitArithmetic_expression is called when exiting the arithmetic_expression production.
	ExitArithmetic_expression(c *Arithmetic_expressionContext)

	// ExitArithmetic_component is called when exiting the arithmetic_component production.
	ExitArithmetic_component(c *Arithmetic_componentContext)

	// ExitCountof_function is called when exiting the countof_function production.
	ExitCountof_function(c *Countof_functionContext)

	// ExitCreate_statement is called when exiting the create_statement production.
	ExitCreate_statement(c *Create_statementContext)

	// ExitAttribute_assignments is called when exiting the attribute_assignments production.
	ExitAttribute_assignments(c *Attribute_assignmentsContext)

	// ExitAttribute_assignment is called when exiting the attribute_assignment production.
	ExitAttribute_assignment(c *Attribute_assignmentContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitFind_statement is called when exiting the find_statement production.
	ExitFind_statement(c *Find_statementContext)

	// ExitFind_instance_statement is called when exiting the find_instance_statement production.
	ExitFind_instance_statement(c *Find_instance_statementContext)

	// ExitFind_set_statement is called when exiting the find_set_statement production.
	ExitFind_set_statement(c *Find_set_statementContext)

	// ExitObject_condition is called when exiting the object_condition production.
	ExitObject_condition(c *Object_conditionContext)

	// ExitEquality_operator is called when exiting the equality_operator production.
	ExitEquality_operator(c *Equality_operatorContext)

	// ExitLogical_operator is called when exiting the logical_operator production.
	ExitLogical_operator(c *Logical_operatorContext)

	// ExitRelationship_navigation is called when exiting the relationship_navigation production.
	ExitRelationship_navigation(c *Relationship_navigationContext)

	// ExitAssociative_relationship_navigation is called when exiting the associative_relationship_navigation production.
	ExitAssociative_relationship_navigation(c *Associative_relationship_navigationContext)

	// ExitRelationship_link is called when exiting the relationship_link production.
	ExitRelationship_link(c *Relationship_linkContext)

	// ExitRelationship_unlink is called when exiting the relationship_unlink production.
	ExitRelationship_unlink(c *Relationship_unlinkContext)

	// ExitRelationship_spec is called when exiting the relationship_spec production.
	ExitRelationship_spec(c *Relationship_specContext)

	// ExitQualified_relationship is called when exiting the qualified_relationship production.
	ExitQualified_relationship(c *Qualified_relationshipContext)

	// ExitRelationship_role is called when exiting the relationship_role production.
	ExitRelationship_role(c *Relationship_roleContext)

	// ExitEvent_generation is called when exiting the event_generation production.
	ExitEvent_generation(c *Event_generationContext)

	// ExitInternal_event_generation is called when exiting the internal_event_generation production.
	ExitInternal_event_generation(c *Internal_event_generationContext)

	// ExitExternal_event_generation is called when exiting the external_event_generation production.
	ExitExternal_event_generation(c *External_event_generationContext)

	// ExitEvent_specification is called when exiting the event_specification production.
	ExitEvent_specification(c *Event_specificationContext)

	// ExitExternal_event_specification is called when exiting the external_event_specification production.
	ExitExternal_event_specification(c *External_event_specificationContext)

	// ExitEvent_parameters is called when exiting the event_parameters production.
	ExitEvent_parameters(c *Event_parametersContext)

	// ExitEvent_parameter is called when exiting the event_parameter production.
	ExitEvent_parameter(c *Event_parameterContext)

	// ExitOperation_invocation is called when exiting the operation_invocation production.
	ExitOperation_invocation(c *Operation_invocationContext)

	// ExitObject_operation is called when exiting the object_operation production.
	ExitObject_operation(c *Object_operationContext)

	// ExitDomain_operation is called when exiting the domain_operation production.
	ExitDomain_operation(c *Domain_operationContext)

	// ExitBridge_operation is called when exiting the bridge_operation production.
	ExitBridge_operation(c *Bridge_operationContext)

	// ExitOo_specification is called when exiting the oo_specification production.
	ExitOo_specification(c *Oo_specificationContext)

	// ExitDo_specification is called when exiting the do_specification production.
	ExitDo_specification(c *Do_specificationContext)

	// ExitBo_specification is called when exiting the bo_specification production.
	ExitBo_specification(c *Bo_specificationContext)

	// ExitOutput_parameter_list is called when exiting the output_parameter_list production.
	ExitOutput_parameter_list(c *Output_parameter_listContext)

	// ExitOutput_parameter is called when exiting the output_parameter production.
	ExitOutput_parameter(c *Output_parameterContext)

	// ExitInput_parameter_list is called when exiting the input_parameter_list production.
	ExitInput_parameter_list(c *Input_parameter_listContext)

	// ExitInput_parameter is called when exiting the input_parameter production.
	ExitInput_parameter(c *Input_parameterContext)

	// ExitTimer_operation is called when exiting the timer_operation production.
	ExitTimer_operation(c *Timer_operationContext)

	// ExitTimer_creation is called when exiting the timer_creation production.
	ExitTimer_creation(c *Timer_creationContext)

	// ExitTimer_deletion is called when exiting the timer_deletion production.
	ExitTimer_deletion(c *Timer_deletionContext)

	// ExitSet_relative_timer is called when exiting the set_relative_timer production.
	ExitSet_relative_timer(c *Set_relative_timerContext)

	// ExitSet_absolute_timer is called when exiting the set_absolute_timer production.
	ExitSet_absolute_timer(c *Set_absolute_timerContext)

	// ExitSet_recurring_timer is called when exiting the set_recurring_timer production.
	ExitSet_recurring_timer(c *Set_recurring_timerContext)

	// ExitTimer_reset is called when exiting the timer_reset production.
	ExitTimer_reset(c *Timer_resetContext)

	// ExitTimer_id is called when exiting the timer_id production.
	ExitTimer_id(c *Timer_idContext)

	// ExitTime_to_fire is called when exiting the time_to_fire production.
	ExitTime_to_fire(c *Time_to_fireContext)

	// ExitReturn_event is called when exiting the return_event production.
	ExitReturn_event(c *Return_eventContext)

	// ExitTarget_instance is called when exiting the target_instance production.
	ExitTarget_instance(c *Target_instanceContext)

	// ExitFire_year is called when exiting the fire_year production.
	ExitFire_year(c *Fire_yearContext)

	// ExitFire_month is called when exiting the fire_month production.
	ExitFire_month(c *Fire_monthContext)

	// ExitFire_date is called when exiting the fire_date production.
	ExitFire_date(c *Fire_dateContext)

	// ExitFire_hour is called when exiting the fire_hour production.
	ExitFire_hour(c *Fire_hourContext)

	// ExitFire_minute is called when exiting the fire_minute production.
	ExitFire_minute(c *Fire_minuteContext)

	// ExitFire_second is called when exiting the fire_second production.
	ExitFire_second(c *Fire_secondContext)

	// ExitStart_time is called when exiting the start_time production.
	ExitStart_time(c *Start_timeContext)

	// ExitPeriod is called when exiting the period production.
	ExitPeriod(c *PeriodContext)

	// ExitSet_operation is called when exiting the set_operation production.
	ExitSet_operation(c *Set_operationContext)

	// ExitUnique_set is called when exiting the unique_set production.
	ExitUnique_set(c *Unique_setContext)

	// ExitSet_combination is called when exiting the set_combination production.
	ExitSet_combination(c *Set_combinationContext)

	// ExitSet_difference is called when exiting the set_difference production.
	ExitSet_difference(c *Set_differenceContext)

	// ExitStruct_statement is called when exiting the struct_statement production.
	ExitStruct_statement(c *Struct_statementContext)

	// ExitStruct_definition is called when exiting the struct_definition production.
	ExitStruct_definition(c *Struct_definitionContext)

	// ExitMember_name is called when exiting the member_name production.
	ExitMember_name(c *Member_nameContext)

	// ExitMember_type is called when exiting the member_type production.
	ExitMember_type(c *Member_typeContext)

	// ExitStruct_instantiation is called when exiting the struct_instantiation production.
	ExitStruct_instantiation(c *Struct_instantiationContext)

	// ExitStruct_assembly is called when exiting the struct_assembly production.
	ExitStruct_assembly(c *Struct_assemblyContext)

	// ExitStruct_for_loop is called when exiting the struct_for_loop production.
	ExitStruct_for_loop(c *Struct_for_loopContext)

	// ExitOrder_struct is called when exiting the order_struct production.
	ExitOrder_struct(c *Order_structContext)

	// ExitStruct_subset is called when exiting the struct_subset production.
	ExitStruct_subset(c *Struct_subsetContext)

	// ExitStruct_condition is called when exiting the struct_condition production.
	ExitStruct_condition(c *Struct_conditionContext)

	// ExitStruct_ is called when exiting the struct_ production.
	ExitStruct_(c *Struct_Context)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitValue_list is called when exiting the value_list production.
	ExitValue_list(c *Value_listContext)

	// ExitStruct_value is called when exiting the struct_value production.
	ExitStruct_value(c *Struct_valueContext)

	// ExitVariable_list is called when exiting the variable_list production.
	ExitVariable_list(c *Variable_listContext)

	// ExitSwitch_statement is called when exiting the switch_statement production.
	ExitSwitch_statement(c *Switch_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitLogical_condition is called when exiting the logical_condition production.
	ExitLogical_condition(c *Logical_conditionContext)

	// ExitCompound_logical_condition is called when exiting the compound_logical_condition production.
	ExitCompound_logical_condition(c *Compound_logical_conditionContext)

	// ExitSimple_logical_condition is called when exiting the simple_logical_condition production.
	ExitSimple_logical_condition(c *Simple_logical_conditionContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitFor_loop is called when exiting the for_loop production.
	ExitFor_loop(c *For_loopContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitNative_language_section is called when exiting the native_language_section production.
	ExitNative_language_section(c *Native_language_sectionContext)

	// ExitSet_variable is called when exiting the set_variable production.
	ExitSet_variable(c *Set_variableContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitObject_name is called when exiting the object_name production.
	ExitObject_name(c *Object_nameContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitObject_attribute is called when exiting the object_attribute production.
	ExitObject_attribute(c *Object_attributeContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitText_string is called when exiting the text_string production.
	ExitText_string(c *Text_stringContext)

	// ExitEnum_value is called when exiting the enum_value production.
	ExitEnum_value(c *Enum_valueContext)

	// ExitStructure_field is called when exiting the structure_field production.
	ExitStructure_field(c *Structure_fieldContext)

	// ExitUser_defined_type is called when exiting the user_defined_type production.
	ExitUser_defined_type(c *User_defined_typeContext)
}
