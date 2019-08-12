// Code generated from vhdl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package vhdl // vhdl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// vhdlListener is a complete listener for a parse tree produced by vhdlParser.
type vhdlListener interface {
	antlr.ParseTreeListener

	// EnterAbstract_literal is called when entering the abstract_literal production.
	EnterAbstract_literal(c *Abstract_literalContext)

	// EnterAccess_type_definition is called when entering the access_type_definition production.
	EnterAccess_type_definition(c *Access_type_definitionContext)

	// EnterAcross_aspect is called when entering the across_aspect production.
	EnterAcross_aspect(c *Across_aspectContext)

	// EnterActual_designator is called when entering the actual_designator production.
	EnterActual_designator(c *Actual_designatorContext)

	// EnterActual_parameter_part is called when entering the actual_parameter_part production.
	EnterActual_parameter_part(c *Actual_parameter_partContext)

	// EnterActual_part is called when entering the actual_part production.
	EnterActual_part(c *Actual_partContext)

	// EnterAdding_operator is called when entering the adding_operator production.
	EnterAdding_operator(c *Adding_operatorContext)

	// EnterAggregate is called when entering the aggregate production.
	EnterAggregate(c *AggregateContext)

	// EnterAlias_declaration is called when entering the alias_declaration production.
	EnterAlias_declaration(c *Alias_declarationContext)

	// EnterAlias_designator is called when entering the alias_designator production.
	EnterAlias_designator(c *Alias_designatorContext)

	// EnterAlias_indication is called when entering the alias_indication production.
	EnterAlias_indication(c *Alias_indicationContext)

	// EnterAllocator is called when entering the allocator production.
	EnterAllocator(c *AllocatorContext)

	// EnterArchitecture_body is called when entering the architecture_body production.
	EnterArchitecture_body(c *Architecture_bodyContext)

	// EnterArchitecture_declarative_part is called when entering the architecture_declarative_part production.
	EnterArchitecture_declarative_part(c *Architecture_declarative_partContext)

	// EnterArchitecture_statement is called when entering the architecture_statement production.
	EnterArchitecture_statement(c *Architecture_statementContext)

	// EnterArchitecture_statement_part is called when entering the architecture_statement_part production.
	EnterArchitecture_statement_part(c *Architecture_statement_partContext)

	// EnterArray_nature_definition is called when entering the array_nature_definition production.
	EnterArray_nature_definition(c *Array_nature_definitionContext)

	// EnterArray_type_definition is called when entering the array_type_definition production.
	EnterArray_type_definition(c *Array_type_definitionContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterAssertion_statement is called when entering the assertion_statement production.
	EnterAssertion_statement(c *Assertion_statementContext)

	// EnterAssociation_element is called when entering the association_element production.
	EnterAssociation_element(c *Association_elementContext)

	// EnterAssociation_list is called when entering the association_list production.
	EnterAssociation_list(c *Association_listContext)

	// EnterAttribute_declaration is called when entering the attribute_declaration production.
	EnterAttribute_declaration(c *Attribute_declarationContext)

	// EnterAttribute_designator is called when entering the attribute_designator production.
	EnterAttribute_designator(c *Attribute_designatorContext)

	// EnterAttribute_specification is called when entering the attribute_specification production.
	EnterAttribute_specification(c *Attribute_specificationContext)

	// EnterBase_unit_declaration is called when entering the base_unit_declaration production.
	EnterBase_unit_declaration(c *Base_unit_declarationContext)

	// EnterBinding_indication is called when entering the binding_indication production.
	EnterBinding_indication(c *Binding_indicationContext)

	// EnterBlock_configuration is called when entering the block_configuration production.
	EnterBlock_configuration(c *Block_configurationContext)

	// EnterBlock_declarative_item is called when entering the block_declarative_item production.
	EnterBlock_declarative_item(c *Block_declarative_itemContext)

	// EnterBlock_declarative_part is called when entering the block_declarative_part production.
	EnterBlock_declarative_part(c *Block_declarative_partContext)

	// EnterBlock_header is called when entering the block_header production.
	EnterBlock_header(c *Block_headerContext)

	// EnterBlock_specification is called when entering the block_specification production.
	EnterBlock_specification(c *Block_specificationContext)

	// EnterBlock_statement is called when entering the block_statement production.
	EnterBlock_statement(c *Block_statementContext)

	// EnterBlock_statement_part is called when entering the block_statement_part production.
	EnterBlock_statement_part(c *Block_statement_partContext)

	// EnterBranch_quantity_declaration is called when entering the branch_quantity_declaration production.
	EnterBranch_quantity_declaration(c *Branch_quantity_declarationContext)

	// EnterBreak_element is called when entering the break_element production.
	EnterBreak_element(c *Break_elementContext)

	// EnterBreak_list is called when entering the break_list production.
	EnterBreak_list(c *Break_listContext)

	// EnterBreak_selector_clause is called when entering the break_selector_clause production.
	EnterBreak_selector_clause(c *Break_selector_clauseContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_statement_alternative is called when entering the case_statement_alternative production.
	EnterCase_statement_alternative(c *Case_statement_alternativeContext)

	// EnterChoice is called when entering the choice production.
	EnterChoice(c *ChoiceContext)

	// EnterChoices is called when entering the choices production.
	EnterChoices(c *ChoicesContext)

	// EnterComponent_configuration is called when entering the component_configuration production.
	EnterComponent_configuration(c *Component_configurationContext)

	// EnterComponent_declaration is called when entering the component_declaration production.
	EnterComponent_declaration(c *Component_declarationContext)

	// EnterComponent_instantiation_statement is called when entering the component_instantiation_statement production.
	EnterComponent_instantiation_statement(c *Component_instantiation_statementContext)

	// EnterComponent_specification is called when entering the component_specification production.
	EnterComponent_specification(c *Component_specificationContext)

	// EnterComposite_nature_definition is called when entering the composite_nature_definition production.
	EnterComposite_nature_definition(c *Composite_nature_definitionContext)

	// EnterComposite_type_definition is called when entering the composite_type_definition production.
	EnterComposite_type_definition(c *Composite_type_definitionContext)

	// EnterConcurrent_assertion_statement is called when entering the concurrent_assertion_statement production.
	EnterConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// EnterConcurrent_break_statement is called when entering the concurrent_break_statement production.
	EnterConcurrent_break_statement(c *Concurrent_break_statementContext)

	// EnterConcurrent_procedure_call_statement is called when entering the concurrent_procedure_call_statement production.
	EnterConcurrent_procedure_call_statement(c *Concurrent_procedure_call_statementContext)

	// EnterConcurrent_signal_assignment_statement is called when entering the concurrent_signal_assignment_statement production.
	EnterConcurrent_signal_assignment_statement(c *Concurrent_signal_assignment_statementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterCondition_clause is called when entering the condition_clause production.
	EnterCondition_clause(c *Condition_clauseContext)

	// EnterConditional_signal_assignment is called when entering the conditional_signal_assignment production.
	EnterConditional_signal_assignment(c *Conditional_signal_assignmentContext)

	// EnterConditional_waveforms is called when entering the conditional_waveforms production.
	EnterConditional_waveforms(c *Conditional_waveformsContext)

	// EnterConfiguration_declaration is called when entering the configuration_declaration production.
	EnterConfiguration_declaration(c *Configuration_declarationContext)

	// EnterConfiguration_declarative_item is called when entering the configuration_declarative_item production.
	EnterConfiguration_declarative_item(c *Configuration_declarative_itemContext)

	// EnterConfiguration_declarative_part is called when entering the configuration_declarative_part production.
	EnterConfiguration_declarative_part(c *Configuration_declarative_partContext)

	// EnterConfiguration_item is called when entering the configuration_item production.
	EnterConfiguration_item(c *Configuration_itemContext)

	// EnterConfiguration_specification is called when entering the configuration_specification production.
	EnterConfiguration_specification(c *Configuration_specificationContext)

	// EnterConstant_declaration is called when entering the constant_declaration production.
	EnterConstant_declaration(c *Constant_declarationContext)

	// EnterConstrained_array_definition is called when entering the constrained_array_definition production.
	EnterConstrained_array_definition(c *Constrained_array_definitionContext)

	// EnterConstrained_nature_definition is called when entering the constrained_nature_definition production.
	EnterConstrained_nature_definition(c *Constrained_nature_definitionContext)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterContext_clause is called when entering the context_clause production.
	EnterContext_clause(c *Context_clauseContext)

	// EnterContext_item is called when entering the context_item production.
	EnterContext_item(c *Context_itemContext)

	// EnterDelay_mechanism is called when entering the delay_mechanism production.
	EnterDelay_mechanism(c *Delay_mechanismContext)

	// EnterDesign_file is called when entering the design_file production.
	EnterDesign_file(c *Design_fileContext)

	// EnterDesign_unit is called when entering the design_unit production.
	EnterDesign_unit(c *Design_unitContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// EnterDisconnection_specification is called when entering the disconnection_specification production.
	EnterDisconnection_specification(c *Disconnection_specificationContext)

	// EnterDiscrete_range is called when entering the discrete_range production.
	EnterDiscrete_range(c *Discrete_rangeContext)

	// EnterElement_association is called when entering the element_association production.
	EnterElement_association(c *Element_associationContext)

	// EnterElement_declaration is called when entering the element_declaration production.
	EnterElement_declaration(c *Element_declarationContext)

	// EnterElement_subnature_definition is called when entering the element_subnature_definition production.
	EnterElement_subnature_definition(c *Element_subnature_definitionContext)

	// EnterElement_subtype_definition is called when entering the element_subtype_definition production.
	EnterElement_subtype_definition(c *Element_subtype_definitionContext)

	// EnterEntity_aspect is called when entering the entity_aspect production.
	EnterEntity_aspect(c *Entity_aspectContext)

	// EnterEntity_class is called when entering the entity_class production.
	EnterEntity_class(c *Entity_classContext)

	// EnterEntity_class_entry is called when entering the entity_class_entry production.
	EnterEntity_class_entry(c *Entity_class_entryContext)

	// EnterEntity_class_entry_list is called when entering the entity_class_entry_list production.
	EnterEntity_class_entry_list(c *Entity_class_entry_listContext)

	// EnterEntity_declaration is called when entering the entity_declaration production.
	EnterEntity_declaration(c *Entity_declarationContext)

	// EnterEntity_declarative_item is called when entering the entity_declarative_item production.
	EnterEntity_declarative_item(c *Entity_declarative_itemContext)

	// EnterEntity_declarative_part is called when entering the entity_declarative_part production.
	EnterEntity_declarative_part(c *Entity_declarative_partContext)

	// EnterEntity_designator is called when entering the entity_designator production.
	EnterEntity_designator(c *Entity_designatorContext)

	// EnterEntity_header is called when entering the entity_header production.
	EnterEntity_header(c *Entity_headerContext)

	// EnterEntity_name_list is called when entering the entity_name_list production.
	EnterEntity_name_list(c *Entity_name_listContext)

	// EnterEntity_specification is called when entering the entity_specification production.
	EnterEntity_specification(c *Entity_specificationContext)

	// EnterEntity_statement is called when entering the entity_statement production.
	EnterEntity_statement(c *Entity_statementContext)

	// EnterEntity_statement_part is called when entering the entity_statement_part production.
	EnterEntity_statement_part(c *Entity_statement_partContext)

	// EnterEntity_tag is called when entering the entity_tag production.
	EnterEntity_tag(c *Entity_tagContext)

	// EnterEnumeration_literal is called when entering the enumeration_literal production.
	EnterEnumeration_literal(c *Enumeration_literalContext)

	// EnterEnumeration_type_definition is called when entering the enumeration_type_definition production.
	EnterEnumeration_type_definition(c *Enumeration_type_definitionContext)

	// EnterExit_statement is called when entering the exit_statement production.
	EnterExit_statement(c *Exit_statementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFile_declaration is called when entering the file_declaration production.
	EnterFile_declaration(c *File_declarationContext)

	// EnterFile_logical_name is called when entering the file_logical_name production.
	EnterFile_logical_name(c *File_logical_nameContext)

	// EnterFile_open_information is called when entering the file_open_information production.
	EnterFile_open_information(c *File_open_informationContext)

	// EnterFile_type_definition is called when entering the file_type_definition production.
	EnterFile_type_definition(c *File_type_definitionContext)

	// EnterFormal_parameter_list is called when entering the formal_parameter_list production.
	EnterFormal_parameter_list(c *Formal_parameter_listContext)

	// EnterFormal_part is called when entering the formal_part production.
	EnterFormal_part(c *Formal_partContext)

	// EnterFree_quantity_declaration is called when entering the free_quantity_declaration production.
	EnterFree_quantity_declaration(c *Free_quantity_declarationContext)

	// EnterGenerate_statement is called when entering the generate_statement production.
	EnterGenerate_statement(c *Generate_statementContext)

	// EnterGeneration_scheme is called when entering the generation_scheme production.
	EnterGeneration_scheme(c *Generation_schemeContext)

	// EnterGeneric_clause is called when entering the generic_clause production.
	EnterGeneric_clause(c *Generic_clauseContext)

	// EnterGeneric_list is called when entering the generic_list production.
	EnterGeneric_list(c *Generic_listContext)

	// EnterGeneric_map_aspect is called when entering the generic_map_aspect production.
	EnterGeneric_map_aspect(c *Generic_map_aspectContext)

	// EnterGroup_constituent is called when entering the group_constituent production.
	EnterGroup_constituent(c *Group_constituentContext)

	// EnterGroup_constituent_list is called when entering the group_constituent_list production.
	EnterGroup_constituent_list(c *Group_constituent_listContext)

	// EnterGroup_declaration is called when entering the group_declaration production.
	EnterGroup_declaration(c *Group_declarationContext)

	// EnterGroup_template_declaration is called when entering the group_template_declaration production.
	EnterGroup_template_declaration(c *Group_template_declarationContext)

	// EnterGuarded_signal_specification is called when entering the guarded_signal_specification production.
	EnterGuarded_signal_specification(c *Guarded_signal_specificationContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterIndex_constraint is called when entering the index_constraint production.
	EnterIndex_constraint(c *Index_constraintContext)

	// EnterIndex_specification is called when entering the index_specification production.
	EnterIndex_specification(c *Index_specificationContext)

	// EnterIndex_subtype_definition is called when entering the index_subtype_definition production.
	EnterIndex_subtype_definition(c *Index_subtype_definitionContext)

	// EnterInstantiated_unit is called when entering the instantiated_unit production.
	EnterInstantiated_unit(c *Instantiated_unitContext)

	// EnterInstantiation_list is called when entering the instantiation_list production.
	EnterInstantiation_list(c *Instantiation_listContext)

	// EnterInterface_constant_declaration is called when entering the interface_constant_declaration production.
	EnterInterface_constant_declaration(c *Interface_constant_declarationContext)

	// EnterInterface_declaration is called when entering the interface_declaration production.
	EnterInterface_declaration(c *Interface_declarationContext)

	// EnterInterface_element is called when entering the interface_element production.
	EnterInterface_element(c *Interface_elementContext)

	// EnterInterface_file_declaration is called when entering the interface_file_declaration production.
	EnterInterface_file_declaration(c *Interface_file_declarationContext)

	// EnterInterface_signal_list is called when entering the interface_signal_list production.
	EnterInterface_signal_list(c *Interface_signal_listContext)

	// EnterInterface_port_list is called when entering the interface_port_list production.
	EnterInterface_port_list(c *Interface_port_listContext)

	// EnterInterface_list is called when entering the interface_list production.
	EnterInterface_list(c *Interface_listContext)

	// EnterInterface_quantity_declaration is called when entering the interface_quantity_declaration production.
	EnterInterface_quantity_declaration(c *Interface_quantity_declarationContext)

	// EnterInterface_port_declaration is called when entering the interface_port_declaration production.
	EnterInterface_port_declaration(c *Interface_port_declarationContext)

	// EnterInterface_signal_declaration is called when entering the interface_signal_declaration production.
	EnterInterface_signal_declaration(c *Interface_signal_declarationContext)

	// EnterInterface_terminal_declaration is called when entering the interface_terminal_declaration production.
	EnterInterface_terminal_declaration(c *Interface_terminal_declarationContext)

	// EnterInterface_variable_declaration is called when entering the interface_variable_declaration production.
	EnterInterface_variable_declaration(c *Interface_variable_declarationContext)

	// EnterIteration_scheme is called when entering the iteration_scheme production.
	EnterIteration_scheme(c *Iteration_schemeContext)

	// EnterLabel_colon is called when entering the label_colon production.
	EnterLabel_colon(c *Label_colonContext)

	// EnterLibrary_clause is called when entering the library_clause production.
	EnterLibrary_clause(c *Library_clauseContext)

	// EnterLibrary_unit is called when entering the library_unit production.
	EnterLibrary_unit(c *Library_unitContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLogical_name is called when entering the logical_name production.
	EnterLogical_name(c *Logical_nameContext)

	// EnterLogical_name_list is called when entering the logical_name_list production.
	EnterLogical_name_list(c *Logical_name_listContext)

	// EnterLogical_operator is called when entering the logical_operator production.
	EnterLogical_operator(c *Logical_operatorContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterSignal_mode is called when entering the signal_mode production.
	EnterSignal_mode(c *Signal_modeContext)

	// EnterMultiplying_operator is called when entering the multiplying_operator production.
	EnterMultiplying_operator(c *Multiplying_operatorContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterName_part is called when entering the name_part production.
	EnterName_part(c *Name_partContext)

	// EnterSelected_name is called when entering the selected_name production.
	EnterSelected_name(c *Selected_nameContext)

	// EnterSelected_name_part is called when entering the selected_name_part production.
	EnterSelected_name_part(c *Selected_name_partContext)

	// EnterFunction_call_or_indexed_name_part is called when entering the function_call_or_indexed_name_part production.
	EnterFunction_call_or_indexed_name_part(c *Function_call_or_indexed_name_partContext)

	// EnterSlice_name_part is called when entering the slice_name_part production.
	EnterSlice_name_part(c *Slice_name_partContext)

	// EnterAttribute_name_part is called when entering the attribute_name_part production.
	EnterAttribute_name_part(c *Attribute_name_partContext)

	// EnterNature_declaration is called when entering the nature_declaration production.
	EnterNature_declaration(c *Nature_declarationContext)

	// EnterNature_definition is called when entering the nature_definition production.
	EnterNature_definition(c *Nature_definitionContext)

	// EnterNature_element_declaration is called when entering the nature_element_declaration production.
	EnterNature_element_declaration(c *Nature_element_declarationContext)

	// EnterNext_statement is called when entering the next_statement production.
	EnterNext_statement(c *Next_statementContext)

	// EnterNumeric_literal is called when entering the numeric_literal production.
	EnterNumeric_literal(c *Numeric_literalContext)

	// EnterObject_declaration is called when entering the object_declaration production.
	EnterObject_declaration(c *Object_declarationContext)

	// EnterOpts is called when entering the opts production.
	EnterOpts(c *OptsContext)

	// EnterPackage_body is called when entering the package_body production.
	EnterPackage_body(c *Package_bodyContext)

	// EnterPackage_body_declarative_item is called when entering the package_body_declarative_item production.
	EnterPackage_body_declarative_item(c *Package_body_declarative_itemContext)

	// EnterPackage_body_declarative_part is called when entering the package_body_declarative_part production.
	EnterPackage_body_declarative_part(c *Package_body_declarative_partContext)

	// EnterPackage_declaration is called when entering the package_declaration production.
	EnterPackage_declaration(c *Package_declarationContext)

	// EnterPackage_declarative_item is called when entering the package_declarative_item production.
	EnterPackage_declarative_item(c *Package_declarative_itemContext)

	// EnterPackage_declarative_part is called when entering the package_declarative_part production.
	EnterPackage_declarative_part(c *Package_declarative_partContext)

	// EnterParameter_specification is called when entering the parameter_specification production.
	EnterParameter_specification(c *Parameter_specificationContext)

	// EnterPhysical_literal is called when entering the physical_literal production.
	EnterPhysical_literal(c *Physical_literalContext)

	// EnterPhysical_type_definition is called when entering the physical_type_definition production.
	EnterPhysical_type_definition(c *Physical_type_definitionContext)

	// EnterPort_clause is called when entering the port_clause production.
	EnterPort_clause(c *Port_clauseContext)

	// EnterPort_list is called when entering the port_list production.
	EnterPort_list(c *Port_listContext)

	// EnterPort_map_aspect is called when entering the port_map_aspect production.
	EnterPort_map_aspect(c *Port_map_aspectContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterPrimary_unit is called when entering the primary_unit production.
	EnterPrimary_unit(c *Primary_unitContext)

	// EnterProcedural_declarative_item is called when entering the procedural_declarative_item production.
	EnterProcedural_declarative_item(c *Procedural_declarative_itemContext)

	// EnterProcedural_declarative_part is called when entering the procedural_declarative_part production.
	EnterProcedural_declarative_part(c *Procedural_declarative_partContext)

	// EnterProcedural_statement_part is called when entering the procedural_statement_part production.
	EnterProcedural_statement_part(c *Procedural_statement_partContext)

	// EnterProcedure_call is called when entering the procedure_call production.
	EnterProcedure_call(c *Procedure_callContext)

	// EnterProcedure_call_statement is called when entering the procedure_call_statement production.
	EnterProcedure_call_statement(c *Procedure_call_statementContext)

	// EnterProcess_declarative_item is called when entering the process_declarative_item production.
	EnterProcess_declarative_item(c *Process_declarative_itemContext)

	// EnterProcess_declarative_part is called when entering the process_declarative_part production.
	EnterProcess_declarative_part(c *Process_declarative_partContext)

	// EnterProcess_statement is called when entering the process_statement production.
	EnterProcess_statement(c *Process_statementContext)

	// EnterProcess_statement_part is called when entering the process_statement_part production.
	EnterProcess_statement_part(c *Process_statement_partContext)

	// EnterQualified_expression is called when entering the qualified_expression production.
	EnterQualified_expression(c *Qualified_expressionContext)

	// EnterQuantity_declaration is called when entering the quantity_declaration production.
	EnterQuantity_declaration(c *Quantity_declarationContext)

	// EnterQuantity_list is called when entering the quantity_list production.
	EnterQuantity_list(c *Quantity_listContext)

	// EnterQuantity_specification is called when entering the quantity_specification production.
	EnterQuantity_specification(c *Quantity_specificationContext)

	// EnterRange_decl is called when entering the range_decl production.
	EnterRange_decl(c *Range_declContext)

	// EnterExplicit_range is called when entering the explicit_range production.
	EnterExplicit_range(c *Explicit_rangeContext)

	// EnterRange_constraint is called when entering the range_constraint production.
	EnterRange_constraint(c *Range_constraintContext)

	// EnterRecord_nature_definition is called when entering the record_nature_definition production.
	EnterRecord_nature_definition(c *Record_nature_definitionContext)

	// EnterRecord_type_definition is called when entering the record_type_definition production.
	EnterRecord_type_definition(c *Record_type_definitionContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterRelational_operator is called when entering the relational_operator production.
	EnterRelational_operator(c *Relational_operatorContext)

	// EnterReport_statement is called when entering the report_statement production.
	EnterReport_statement(c *Report_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterScalar_nature_definition is called when entering the scalar_nature_definition production.
	EnterScalar_nature_definition(c *Scalar_nature_definitionContext)

	// EnterScalar_type_definition is called when entering the scalar_type_definition production.
	EnterScalar_type_definition(c *Scalar_type_definitionContext)

	// EnterSecondary_unit is called when entering the secondary_unit production.
	EnterSecondary_unit(c *Secondary_unitContext)

	// EnterSecondary_unit_declaration is called when entering the secondary_unit_declaration production.
	EnterSecondary_unit_declaration(c *Secondary_unit_declarationContext)

	// EnterSelected_signal_assignment is called when entering the selected_signal_assignment production.
	EnterSelected_signal_assignment(c *Selected_signal_assignmentContext)

	// EnterSelected_waveforms is called when entering the selected_waveforms production.
	EnterSelected_waveforms(c *Selected_waveformsContext)

	// EnterSensitivity_clause is called when entering the sensitivity_clause production.
	EnterSensitivity_clause(c *Sensitivity_clauseContext)

	// EnterSensitivity_list is called when entering the sensitivity_list production.
	EnterSensitivity_list(c *Sensitivity_listContext)

	// EnterSequence_of_statements is called when entering the sequence_of_statements production.
	EnterSequence_of_statements(c *Sequence_of_statementsContext)

	// EnterSequential_statement is called when entering the sequential_statement production.
	EnterSequential_statement(c *Sequential_statementContext)

	// EnterShift_expression is called when entering the shift_expression production.
	EnterShift_expression(c *Shift_expressionContext)

	// EnterShift_operator is called when entering the shift_operator production.
	EnterShift_operator(c *Shift_operatorContext)

	// EnterSignal_assignment_statement is called when entering the signal_assignment_statement production.
	EnterSignal_assignment_statement(c *Signal_assignment_statementContext)

	// EnterSignal_declaration is called when entering the signal_declaration production.
	EnterSignal_declaration(c *Signal_declarationContext)

	// EnterSignal_kind is called when entering the signal_kind production.
	EnterSignal_kind(c *Signal_kindContext)

	// EnterSignal_list is called when entering the signal_list production.
	EnterSignal_list(c *Signal_listContext)

	// EnterSignature is called when entering the signature production.
	EnterSignature(c *SignatureContext)

	// EnterSimple_expression is called when entering the simple_expression production.
	EnterSimple_expression(c *Simple_expressionContext)

	// EnterSimple_simultaneous_statement is called when entering the simple_simultaneous_statement production.
	EnterSimple_simultaneous_statement(c *Simple_simultaneous_statementContext)

	// EnterSimultaneous_alternative is called when entering the simultaneous_alternative production.
	EnterSimultaneous_alternative(c *Simultaneous_alternativeContext)

	// EnterSimultaneous_case_statement is called when entering the simultaneous_case_statement production.
	EnterSimultaneous_case_statement(c *Simultaneous_case_statementContext)

	// EnterSimultaneous_if_statement is called when entering the simultaneous_if_statement production.
	EnterSimultaneous_if_statement(c *Simultaneous_if_statementContext)

	// EnterSimultaneous_procedural_statement is called when entering the simultaneous_procedural_statement production.
	EnterSimultaneous_procedural_statement(c *Simultaneous_procedural_statementContext)

	// EnterSimultaneous_statement is called when entering the simultaneous_statement production.
	EnterSimultaneous_statement(c *Simultaneous_statementContext)

	// EnterSimultaneous_statement_part is called when entering the simultaneous_statement_part production.
	EnterSimultaneous_statement_part(c *Simultaneous_statement_partContext)

	// EnterSource_aspect is called when entering the source_aspect production.
	EnterSource_aspect(c *Source_aspectContext)

	// EnterSource_quantity_declaration is called when entering the source_quantity_declaration production.
	EnterSource_quantity_declaration(c *Source_quantity_declarationContext)

	// EnterStep_limit_specification is called when entering the step_limit_specification production.
	EnterStep_limit_specification(c *Step_limit_specificationContext)

	// EnterSubnature_declaration is called when entering the subnature_declaration production.
	EnterSubnature_declaration(c *Subnature_declarationContext)

	// EnterSubnature_indication is called when entering the subnature_indication production.
	EnterSubnature_indication(c *Subnature_indicationContext)

	// EnterSubprogram_body is called when entering the subprogram_body production.
	EnterSubprogram_body(c *Subprogram_bodyContext)

	// EnterSubprogram_declaration is called when entering the subprogram_declaration production.
	EnterSubprogram_declaration(c *Subprogram_declarationContext)

	// EnterSubprogram_declarative_item is called when entering the subprogram_declarative_item production.
	EnterSubprogram_declarative_item(c *Subprogram_declarative_itemContext)

	// EnterSubprogram_declarative_part is called when entering the subprogram_declarative_part production.
	EnterSubprogram_declarative_part(c *Subprogram_declarative_partContext)

	// EnterSubprogram_kind is called when entering the subprogram_kind production.
	EnterSubprogram_kind(c *Subprogram_kindContext)

	// EnterSubprogram_specification is called when entering the subprogram_specification production.
	EnterSubprogram_specification(c *Subprogram_specificationContext)

	// EnterProcedure_specification is called when entering the procedure_specification production.
	EnterProcedure_specification(c *Procedure_specificationContext)

	// EnterFunction_specification is called when entering the function_specification production.
	EnterFunction_specification(c *Function_specificationContext)

	// EnterSubprogram_statement_part is called when entering the subprogram_statement_part production.
	EnterSubprogram_statement_part(c *Subprogram_statement_partContext)

	// EnterSubtype_declaration is called when entering the subtype_declaration production.
	EnterSubtype_declaration(c *Subtype_declarationContext)

	// EnterSubtype_indication is called when entering the subtype_indication production.
	EnterSubtype_indication(c *Subtype_indicationContext)

	// EnterSuffix is called when entering the suffix production.
	EnterSuffix(c *SuffixContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTerminal_aspect is called when entering the terminal_aspect production.
	EnterTerminal_aspect(c *Terminal_aspectContext)

	// EnterTerminal_declaration is called when entering the terminal_declaration production.
	EnterTerminal_declaration(c *Terminal_declarationContext)

	// EnterThrough_aspect is called when entering the through_aspect production.
	EnterThrough_aspect(c *Through_aspectContext)

	// EnterTimeout_clause is called when entering the timeout_clause production.
	EnterTimeout_clause(c *Timeout_clauseContext)

	// EnterTolerance_aspect is called when entering the tolerance_aspect production.
	EnterTolerance_aspect(c *Tolerance_aspectContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterType_definition is called when entering the type_definition production.
	EnterType_definition(c *Type_definitionContext)

	// EnterUnconstrained_array_definition is called when entering the unconstrained_array_definition production.
	EnterUnconstrained_array_definition(c *Unconstrained_array_definitionContext)

	// EnterUnconstrained_nature_definition is called when entering the unconstrained_nature_definition production.
	EnterUnconstrained_nature_definition(c *Unconstrained_nature_definitionContext)

	// EnterUse_clause is called when entering the use_clause production.
	EnterUse_clause(c *Use_clauseContext)

	// EnterVariable_assignment_statement is called when entering the variable_assignment_statement production.
	EnterVariable_assignment_statement(c *Variable_assignment_statementContext)

	// EnterVariable_declaration is called when entering the variable_declaration production.
	EnterVariable_declaration(c *Variable_declarationContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterWaveform is called when entering the waveform production.
	EnterWaveform(c *WaveformContext)

	// EnterWaveform_element is called when entering the waveform_element production.
	EnterWaveform_element(c *Waveform_elementContext)

	// ExitAbstract_literal is called when exiting the abstract_literal production.
	ExitAbstract_literal(c *Abstract_literalContext)

	// ExitAccess_type_definition is called when exiting the access_type_definition production.
	ExitAccess_type_definition(c *Access_type_definitionContext)

	// ExitAcross_aspect is called when exiting the across_aspect production.
	ExitAcross_aspect(c *Across_aspectContext)

	// ExitActual_designator is called when exiting the actual_designator production.
	ExitActual_designator(c *Actual_designatorContext)

	// ExitActual_parameter_part is called when exiting the actual_parameter_part production.
	ExitActual_parameter_part(c *Actual_parameter_partContext)

	// ExitActual_part is called when exiting the actual_part production.
	ExitActual_part(c *Actual_partContext)

	// ExitAdding_operator is called when exiting the adding_operator production.
	ExitAdding_operator(c *Adding_operatorContext)

	// ExitAggregate is called when exiting the aggregate production.
	ExitAggregate(c *AggregateContext)

	// ExitAlias_declaration is called when exiting the alias_declaration production.
	ExitAlias_declaration(c *Alias_declarationContext)

	// ExitAlias_designator is called when exiting the alias_designator production.
	ExitAlias_designator(c *Alias_designatorContext)

	// ExitAlias_indication is called when exiting the alias_indication production.
	ExitAlias_indication(c *Alias_indicationContext)

	// ExitAllocator is called when exiting the allocator production.
	ExitAllocator(c *AllocatorContext)

	// ExitArchitecture_body is called when exiting the architecture_body production.
	ExitArchitecture_body(c *Architecture_bodyContext)

	// ExitArchitecture_declarative_part is called when exiting the architecture_declarative_part production.
	ExitArchitecture_declarative_part(c *Architecture_declarative_partContext)

	// ExitArchitecture_statement is called when exiting the architecture_statement production.
	ExitArchitecture_statement(c *Architecture_statementContext)

	// ExitArchitecture_statement_part is called when exiting the architecture_statement_part production.
	ExitArchitecture_statement_part(c *Architecture_statement_partContext)

	// ExitArray_nature_definition is called when exiting the array_nature_definition production.
	ExitArray_nature_definition(c *Array_nature_definitionContext)

	// ExitArray_type_definition is called when exiting the array_type_definition production.
	ExitArray_type_definition(c *Array_type_definitionContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitAssertion_statement is called when exiting the assertion_statement production.
	ExitAssertion_statement(c *Assertion_statementContext)

	// ExitAssociation_element is called when exiting the association_element production.
	ExitAssociation_element(c *Association_elementContext)

	// ExitAssociation_list is called when exiting the association_list production.
	ExitAssociation_list(c *Association_listContext)

	// ExitAttribute_declaration is called when exiting the attribute_declaration production.
	ExitAttribute_declaration(c *Attribute_declarationContext)

	// ExitAttribute_designator is called when exiting the attribute_designator production.
	ExitAttribute_designator(c *Attribute_designatorContext)

	// ExitAttribute_specification is called when exiting the attribute_specification production.
	ExitAttribute_specification(c *Attribute_specificationContext)

	// ExitBase_unit_declaration is called when exiting the base_unit_declaration production.
	ExitBase_unit_declaration(c *Base_unit_declarationContext)

	// ExitBinding_indication is called when exiting the binding_indication production.
	ExitBinding_indication(c *Binding_indicationContext)

	// ExitBlock_configuration is called when exiting the block_configuration production.
	ExitBlock_configuration(c *Block_configurationContext)

	// ExitBlock_declarative_item is called when exiting the block_declarative_item production.
	ExitBlock_declarative_item(c *Block_declarative_itemContext)

	// ExitBlock_declarative_part is called when exiting the block_declarative_part production.
	ExitBlock_declarative_part(c *Block_declarative_partContext)

	// ExitBlock_header is called when exiting the block_header production.
	ExitBlock_header(c *Block_headerContext)

	// ExitBlock_specification is called when exiting the block_specification production.
	ExitBlock_specification(c *Block_specificationContext)

	// ExitBlock_statement is called when exiting the block_statement production.
	ExitBlock_statement(c *Block_statementContext)

	// ExitBlock_statement_part is called when exiting the block_statement_part production.
	ExitBlock_statement_part(c *Block_statement_partContext)

	// ExitBranch_quantity_declaration is called when exiting the branch_quantity_declaration production.
	ExitBranch_quantity_declaration(c *Branch_quantity_declarationContext)

	// ExitBreak_element is called when exiting the break_element production.
	ExitBreak_element(c *Break_elementContext)

	// ExitBreak_list is called when exiting the break_list production.
	ExitBreak_list(c *Break_listContext)

	// ExitBreak_selector_clause is called when exiting the break_selector_clause production.
	ExitBreak_selector_clause(c *Break_selector_clauseContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_statement_alternative is called when exiting the case_statement_alternative production.
	ExitCase_statement_alternative(c *Case_statement_alternativeContext)

	// ExitChoice is called when exiting the choice production.
	ExitChoice(c *ChoiceContext)

	// ExitChoices is called when exiting the choices production.
	ExitChoices(c *ChoicesContext)

	// ExitComponent_configuration is called when exiting the component_configuration production.
	ExitComponent_configuration(c *Component_configurationContext)

	// ExitComponent_declaration is called when exiting the component_declaration production.
	ExitComponent_declaration(c *Component_declarationContext)

	// ExitComponent_instantiation_statement is called when exiting the component_instantiation_statement production.
	ExitComponent_instantiation_statement(c *Component_instantiation_statementContext)

	// ExitComponent_specification is called when exiting the component_specification production.
	ExitComponent_specification(c *Component_specificationContext)

	// ExitComposite_nature_definition is called when exiting the composite_nature_definition production.
	ExitComposite_nature_definition(c *Composite_nature_definitionContext)

	// ExitComposite_type_definition is called when exiting the composite_type_definition production.
	ExitComposite_type_definition(c *Composite_type_definitionContext)

	// ExitConcurrent_assertion_statement is called when exiting the concurrent_assertion_statement production.
	ExitConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// ExitConcurrent_break_statement is called when exiting the concurrent_break_statement production.
	ExitConcurrent_break_statement(c *Concurrent_break_statementContext)

	// ExitConcurrent_procedure_call_statement is called when exiting the concurrent_procedure_call_statement production.
	ExitConcurrent_procedure_call_statement(c *Concurrent_procedure_call_statementContext)

	// ExitConcurrent_signal_assignment_statement is called when exiting the concurrent_signal_assignment_statement production.
	ExitConcurrent_signal_assignment_statement(c *Concurrent_signal_assignment_statementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitCondition_clause is called when exiting the condition_clause production.
	ExitCondition_clause(c *Condition_clauseContext)

	// ExitConditional_signal_assignment is called when exiting the conditional_signal_assignment production.
	ExitConditional_signal_assignment(c *Conditional_signal_assignmentContext)

	// ExitConditional_waveforms is called when exiting the conditional_waveforms production.
	ExitConditional_waveforms(c *Conditional_waveformsContext)

	// ExitConfiguration_declaration is called when exiting the configuration_declaration production.
	ExitConfiguration_declaration(c *Configuration_declarationContext)

	// ExitConfiguration_declarative_item is called when exiting the configuration_declarative_item production.
	ExitConfiguration_declarative_item(c *Configuration_declarative_itemContext)

	// ExitConfiguration_declarative_part is called when exiting the configuration_declarative_part production.
	ExitConfiguration_declarative_part(c *Configuration_declarative_partContext)

	// ExitConfiguration_item is called when exiting the configuration_item production.
	ExitConfiguration_item(c *Configuration_itemContext)

	// ExitConfiguration_specification is called when exiting the configuration_specification production.
	ExitConfiguration_specification(c *Configuration_specificationContext)

	// ExitConstant_declaration is called when exiting the constant_declaration production.
	ExitConstant_declaration(c *Constant_declarationContext)

	// ExitConstrained_array_definition is called when exiting the constrained_array_definition production.
	ExitConstrained_array_definition(c *Constrained_array_definitionContext)

	// ExitConstrained_nature_definition is called when exiting the constrained_nature_definition production.
	ExitConstrained_nature_definition(c *Constrained_nature_definitionContext)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitContext_clause is called when exiting the context_clause production.
	ExitContext_clause(c *Context_clauseContext)

	// ExitContext_item is called when exiting the context_item production.
	ExitContext_item(c *Context_itemContext)

	// ExitDelay_mechanism is called when exiting the delay_mechanism production.
	ExitDelay_mechanism(c *Delay_mechanismContext)

	// ExitDesign_file is called when exiting the design_file production.
	ExitDesign_file(c *Design_fileContext)

	// ExitDesign_unit is called when exiting the design_unit production.
	ExitDesign_unit(c *Design_unitContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)

	// ExitDisconnection_specification is called when exiting the disconnection_specification production.
	ExitDisconnection_specification(c *Disconnection_specificationContext)

	// ExitDiscrete_range is called when exiting the discrete_range production.
	ExitDiscrete_range(c *Discrete_rangeContext)

	// ExitElement_association is called when exiting the element_association production.
	ExitElement_association(c *Element_associationContext)

	// ExitElement_declaration is called when exiting the element_declaration production.
	ExitElement_declaration(c *Element_declarationContext)

	// ExitElement_subnature_definition is called when exiting the element_subnature_definition production.
	ExitElement_subnature_definition(c *Element_subnature_definitionContext)

	// ExitElement_subtype_definition is called when exiting the element_subtype_definition production.
	ExitElement_subtype_definition(c *Element_subtype_definitionContext)

	// ExitEntity_aspect is called when exiting the entity_aspect production.
	ExitEntity_aspect(c *Entity_aspectContext)

	// ExitEntity_class is called when exiting the entity_class production.
	ExitEntity_class(c *Entity_classContext)

	// ExitEntity_class_entry is called when exiting the entity_class_entry production.
	ExitEntity_class_entry(c *Entity_class_entryContext)

	// ExitEntity_class_entry_list is called when exiting the entity_class_entry_list production.
	ExitEntity_class_entry_list(c *Entity_class_entry_listContext)

	// ExitEntity_declaration is called when exiting the entity_declaration production.
	ExitEntity_declaration(c *Entity_declarationContext)

	// ExitEntity_declarative_item is called when exiting the entity_declarative_item production.
	ExitEntity_declarative_item(c *Entity_declarative_itemContext)

	// ExitEntity_declarative_part is called when exiting the entity_declarative_part production.
	ExitEntity_declarative_part(c *Entity_declarative_partContext)

	// ExitEntity_designator is called when exiting the entity_designator production.
	ExitEntity_designator(c *Entity_designatorContext)

	// ExitEntity_header is called when exiting the entity_header production.
	ExitEntity_header(c *Entity_headerContext)

	// ExitEntity_name_list is called when exiting the entity_name_list production.
	ExitEntity_name_list(c *Entity_name_listContext)

	// ExitEntity_specification is called when exiting the entity_specification production.
	ExitEntity_specification(c *Entity_specificationContext)

	// ExitEntity_statement is called when exiting the entity_statement production.
	ExitEntity_statement(c *Entity_statementContext)

	// ExitEntity_statement_part is called when exiting the entity_statement_part production.
	ExitEntity_statement_part(c *Entity_statement_partContext)

	// ExitEntity_tag is called when exiting the entity_tag production.
	ExitEntity_tag(c *Entity_tagContext)

	// ExitEnumeration_literal is called when exiting the enumeration_literal production.
	ExitEnumeration_literal(c *Enumeration_literalContext)

	// ExitEnumeration_type_definition is called when exiting the enumeration_type_definition production.
	ExitEnumeration_type_definition(c *Enumeration_type_definitionContext)

	// ExitExit_statement is called when exiting the exit_statement production.
	ExitExit_statement(c *Exit_statementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFile_declaration is called when exiting the file_declaration production.
	ExitFile_declaration(c *File_declarationContext)

	// ExitFile_logical_name is called when exiting the file_logical_name production.
	ExitFile_logical_name(c *File_logical_nameContext)

	// ExitFile_open_information is called when exiting the file_open_information production.
	ExitFile_open_information(c *File_open_informationContext)

	// ExitFile_type_definition is called when exiting the file_type_definition production.
	ExitFile_type_definition(c *File_type_definitionContext)

	// ExitFormal_parameter_list is called when exiting the formal_parameter_list production.
	ExitFormal_parameter_list(c *Formal_parameter_listContext)

	// ExitFormal_part is called when exiting the formal_part production.
	ExitFormal_part(c *Formal_partContext)

	// ExitFree_quantity_declaration is called when exiting the free_quantity_declaration production.
	ExitFree_quantity_declaration(c *Free_quantity_declarationContext)

	// ExitGenerate_statement is called when exiting the generate_statement production.
	ExitGenerate_statement(c *Generate_statementContext)

	// ExitGeneration_scheme is called when exiting the generation_scheme production.
	ExitGeneration_scheme(c *Generation_schemeContext)

	// ExitGeneric_clause is called when exiting the generic_clause production.
	ExitGeneric_clause(c *Generic_clauseContext)

	// ExitGeneric_list is called when exiting the generic_list production.
	ExitGeneric_list(c *Generic_listContext)

	// ExitGeneric_map_aspect is called when exiting the generic_map_aspect production.
	ExitGeneric_map_aspect(c *Generic_map_aspectContext)

	// ExitGroup_constituent is called when exiting the group_constituent production.
	ExitGroup_constituent(c *Group_constituentContext)

	// ExitGroup_constituent_list is called when exiting the group_constituent_list production.
	ExitGroup_constituent_list(c *Group_constituent_listContext)

	// ExitGroup_declaration is called when exiting the group_declaration production.
	ExitGroup_declaration(c *Group_declarationContext)

	// ExitGroup_template_declaration is called when exiting the group_template_declaration production.
	ExitGroup_template_declaration(c *Group_template_declarationContext)

	// ExitGuarded_signal_specification is called when exiting the guarded_signal_specification production.
	ExitGuarded_signal_specification(c *Guarded_signal_specificationContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitIndex_constraint is called when exiting the index_constraint production.
	ExitIndex_constraint(c *Index_constraintContext)

	// ExitIndex_specification is called when exiting the index_specification production.
	ExitIndex_specification(c *Index_specificationContext)

	// ExitIndex_subtype_definition is called when exiting the index_subtype_definition production.
	ExitIndex_subtype_definition(c *Index_subtype_definitionContext)

	// ExitInstantiated_unit is called when exiting the instantiated_unit production.
	ExitInstantiated_unit(c *Instantiated_unitContext)

	// ExitInstantiation_list is called when exiting the instantiation_list production.
	ExitInstantiation_list(c *Instantiation_listContext)

	// ExitInterface_constant_declaration is called when exiting the interface_constant_declaration production.
	ExitInterface_constant_declaration(c *Interface_constant_declarationContext)

	// ExitInterface_declaration is called when exiting the interface_declaration production.
	ExitInterface_declaration(c *Interface_declarationContext)

	// ExitInterface_element is called when exiting the interface_element production.
	ExitInterface_element(c *Interface_elementContext)

	// ExitInterface_file_declaration is called when exiting the interface_file_declaration production.
	ExitInterface_file_declaration(c *Interface_file_declarationContext)

	// ExitInterface_signal_list is called when exiting the interface_signal_list production.
	ExitInterface_signal_list(c *Interface_signal_listContext)

	// ExitInterface_port_list is called when exiting the interface_port_list production.
	ExitInterface_port_list(c *Interface_port_listContext)

	// ExitInterface_list is called when exiting the interface_list production.
	ExitInterface_list(c *Interface_listContext)

	// ExitInterface_quantity_declaration is called when exiting the interface_quantity_declaration production.
	ExitInterface_quantity_declaration(c *Interface_quantity_declarationContext)

	// ExitInterface_port_declaration is called when exiting the interface_port_declaration production.
	ExitInterface_port_declaration(c *Interface_port_declarationContext)

	// ExitInterface_signal_declaration is called when exiting the interface_signal_declaration production.
	ExitInterface_signal_declaration(c *Interface_signal_declarationContext)

	// ExitInterface_terminal_declaration is called when exiting the interface_terminal_declaration production.
	ExitInterface_terminal_declaration(c *Interface_terminal_declarationContext)

	// ExitInterface_variable_declaration is called when exiting the interface_variable_declaration production.
	ExitInterface_variable_declaration(c *Interface_variable_declarationContext)

	// ExitIteration_scheme is called when exiting the iteration_scheme production.
	ExitIteration_scheme(c *Iteration_schemeContext)

	// ExitLabel_colon is called when exiting the label_colon production.
	ExitLabel_colon(c *Label_colonContext)

	// ExitLibrary_clause is called when exiting the library_clause production.
	ExitLibrary_clause(c *Library_clauseContext)

	// ExitLibrary_unit is called when exiting the library_unit production.
	ExitLibrary_unit(c *Library_unitContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLogical_name is called when exiting the logical_name production.
	ExitLogical_name(c *Logical_nameContext)

	// ExitLogical_name_list is called when exiting the logical_name_list production.
	ExitLogical_name_list(c *Logical_name_listContext)

	// ExitLogical_operator is called when exiting the logical_operator production.
	ExitLogical_operator(c *Logical_operatorContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitSignal_mode is called when exiting the signal_mode production.
	ExitSignal_mode(c *Signal_modeContext)

	// ExitMultiplying_operator is called when exiting the multiplying_operator production.
	ExitMultiplying_operator(c *Multiplying_operatorContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitName_part is called when exiting the name_part production.
	ExitName_part(c *Name_partContext)

	// ExitSelected_name is called when exiting the selected_name production.
	ExitSelected_name(c *Selected_nameContext)

	// ExitSelected_name_part is called when exiting the selected_name_part production.
	ExitSelected_name_part(c *Selected_name_partContext)

	// ExitFunction_call_or_indexed_name_part is called when exiting the function_call_or_indexed_name_part production.
	ExitFunction_call_or_indexed_name_part(c *Function_call_or_indexed_name_partContext)

	// ExitSlice_name_part is called when exiting the slice_name_part production.
	ExitSlice_name_part(c *Slice_name_partContext)

	// ExitAttribute_name_part is called when exiting the attribute_name_part production.
	ExitAttribute_name_part(c *Attribute_name_partContext)

	// ExitNature_declaration is called when exiting the nature_declaration production.
	ExitNature_declaration(c *Nature_declarationContext)

	// ExitNature_definition is called when exiting the nature_definition production.
	ExitNature_definition(c *Nature_definitionContext)

	// ExitNature_element_declaration is called when exiting the nature_element_declaration production.
	ExitNature_element_declaration(c *Nature_element_declarationContext)

	// ExitNext_statement is called when exiting the next_statement production.
	ExitNext_statement(c *Next_statementContext)

	// ExitNumeric_literal is called when exiting the numeric_literal production.
	ExitNumeric_literal(c *Numeric_literalContext)

	// ExitObject_declaration is called when exiting the object_declaration production.
	ExitObject_declaration(c *Object_declarationContext)

	// ExitOpts is called when exiting the opts production.
	ExitOpts(c *OptsContext)

	// ExitPackage_body is called when exiting the package_body production.
	ExitPackage_body(c *Package_bodyContext)

	// ExitPackage_body_declarative_item is called when exiting the package_body_declarative_item production.
	ExitPackage_body_declarative_item(c *Package_body_declarative_itemContext)

	// ExitPackage_body_declarative_part is called when exiting the package_body_declarative_part production.
	ExitPackage_body_declarative_part(c *Package_body_declarative_partContext)

	// ExitPackage_declaration is called when exiting the package_declaration production.
	ExitPackage_declaration(c *Package_declarationContext)

	// ExitPackage_declarative_item is called when exiting the package_declarative_item production.
	ExitPackage_declarative_item(c *Package_declarative_itemContext)

	// ExitPackage_declarative_part is called when exiting the package_declarative_part production.
	ExitPackage_declarative_part(c *Package_declarative_partContext)

	// ExitParameter_specification is called when exiting the parameter_specification production.
	ExitParameter_specification(c *Parameter_specificationContext)

	// ExitPhysical_literal is called when exiting the physical_literal production.
	ExitPhysical_literal(c *Physical_literalContext)

	// ExitPhysical_type_definition is called when exiting the physical_type_definition production.
	ExitPhysical_type_definition(c *Physical_type_definitionContext)

	// ExitPort_clause is called when exiting the port_clause production.
	ExitPort_clause(c *Port_clauseContext)

	// ExitPort_list is called when exiting the port_list production.
	ExitPort_list(c *Port_listContext)

	// ExitPort_map_aspect is called when exiting the port_map_aspect production.
	ExitPort_map_aspect(c *Port_map_aspectContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitPrimary_unit is called when exiting the primary_unit production.
	ExitPrimary_unit(c *Primary_unitContext)

	// ExitProcedural_declarative_item is called when exiting the procedural_declarative_item production.
	ExitProcedural_declarative_item(c *Procedural_declarative_itemContext)

	// ExitProcedural_declarative_part is called when exiting the procedural_declarative_part production.
	ExitProcedural_declarative_part(c *Procedural_declarative_partContext)

	// ExitProcedural_statement_part is called when exiting the procedural_statement_part production.
	ExitProcedural_statement_part(c *Procedural_statement_partContext)

	// ExitProcedure_call is called when exiting the procedure_call production.
	ExitProcedure_call(c *Procedure_callContext)

	// ExitProcedure_call_statement is called when exiting the procedure_call_statement production.
	ExitProcedure_call_statement(c *Procedure_call_statementContext)

	// ExitProcess_declarative_item is called when exiting the process_declarative_item production.
	ExitProcess_declarative_item(c *Process_declarative_itemContext)

	// ExitProcess_declarative_part is called when exiting the process_declarative_part production.
	ExitProcess_declarative_part(c *Process_declarative_partContext)

	// ExitProcess_statement is called when exiting the process_statement production.
	ExitProcess_statement(c *Process_statementContext)

	// ExitProcess_statement_part is called when exiting the process_statement_part production.
	ExitProcess_statement_part(c *Process_statement_partContext)

	// ExitQualified_expression is called when exiting the qualified_expression production.
	ExitQualified_expression(c *Qualified_expressionContext)

	// ExitQuantity_declaration is called when exiting the quantity_declaration production.
	ExitQuantity_declaration(c *Quantity_declarationContext)

	// ExitQuantity_list is called when exiting the quantity_list production.
	ExitQuantity_list(c *Quantity_listContext)

	// ExitQuantity_specification is called when exiting the quantity_specification production.
	ExitQuantity_specification(c *Quantity_specificationContext)

	// ExitRange_decl is called when exiting the range_decl production.
	ExitRange_decl(c *Range_declContext)

	// ExitExplicit_range is called when exiting the explicit_range production.
	ExitExplicit_range(c *Explicit_rangeContext)

	// ExitRange_constraint is called when exiting the range_constraint production.
	ExitRange_constraint(c *Range_constraintContext)

	// ExitRecord_nature_definition is called when exiting the record_nature_definition production.
	ExitRecord_nature_definition(c *Record_nature_definitionContext)

	// ExitRecord_type_definition is called when exiting the record_type_definition production.
	ExitRecord_type_definition(c *Record_type_definitionContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitRelational_operator is called when exiting the relational_operator production.
	ExitRelational_operator(c *Relational_operatorContext)

	// ExitReport_statement is called when exiting the report_statement production.
	ExitReport_statement(c *Report_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitScalar_nature_definition is called when exiting the scalar_nature_definition production.
	ExitScalar_nature_definition(c *Scalar_nature_definitionContext)

	// ExitScalar_type_definition is called when exiting the scalar_type_definition production.
	ExitScalar_type_definition(c *Scalar_type_definitionContext)

	// ExitSecondary_unit is called when exiting the secondary_unit production.
	ExitSecondary_unit(c *Secondary_unitContext)

	// ExitSecondary_unit_declaration is called when exiting the secondary_unit_declaration production.
	ExitSecondary_unit_declaration(c *Secondary_unit_declarationContext)

	// ExitSelected_signal_assignment is called when exiting the selected_signal_assignment production.
	ExitSelected_signal_assignment(c *Selected_signal_assignmentContext)

	// ExitSelected_waveforms is called when exiting the selected_waveforms production.
	ExitSelected_waveforms(c *Selected_waveformsContext)

	// ExitSensitivity_clause is called when exiting the sensitivity_clause production.
	ExitSensitivity_clause(c *Sensitivity_clauseContext)

	// ExitSensitivity_list is called when exiting the sensitivity_list production.
	ExitSensitivity_list(c *Sensitivity_listContext)

	// ExitSequence_of_statements is called when exiting the sequence_of_statements production.
	ExitSequence_of_statements(c *Sequence_of_statementsContext)

	// ExitSequential_statement is called when exiting the sequential_statement production.
	ExitSequential_statement(c *Sequential_statementContext)

	// ExitShift_expression is called when exiting the shift_expression production.
	ExitShift_expression(c *Shift_expressionContext)

	// ExitShift_operator is called when exiting the shift_operator production.
	ExitShift_operator(c *Shift_operatorContext)

	// ExitSignal_assignment_statement is called when exiting the signal_assignment_statement production.
	ExitSignal_assignment_statement(c *Signal_assignment_statementContext)

	// ExitSignal_declaration is called when exiting the signal_declaration production.
	ExitSignal_declaration(c *Signal_declarationContext)

	// ExitSignal_kind is called when exiting the signal_kind production.
	ExitSignal_kind(c *Signal_kindContext)

	// ExitSignal_list is called when exiting the signal_list production.
	ExitSignal_list(c *Signal_listContext)

	// ExitSignature is called when exiting the signature production.
	ExitSignature(c *SignatureContext)

	// ExitSimple_expression is called when exiting the simple_expression production.
	ExitSimple_expression(c *Simple_expressionContext)

	// ExitSimple_simultaneous_statement is called when exiting the simple_simultaneous_statement production.
	ExitSimple_simultaneous_statement(c *Simple_simultaneous_statementContext)

	// ExitSimultaneous_alternative is called when exiting the simultaneous_alternative production.
	ExitSimultaneous_alternative(c *Simultaneous_alternativeContext)

	// ExitSimultaneous_case_statement is called when exiting the simultaneous_case_statement production.
	ExitSimultaneous_case_statement(c *Simultaneous_case_statementContext)

	// ExitSimultaneous_if_statement is called when exiting the simultaneous_if_statement production.
	ExitSimultaneous_if_statement(c *Simultaneous_if_statementContext)

	// ExitSimultaneous_procedural_statement is called when exiting the simultaneous_procedural_statement production.
	ExitSimultaneous_procedural_statement(c *Simultaneous_procedural_statementContext)

	// ExitSimultaneous_statement is called when exiting the simultaneous_statement production.
	ExitSimultaneous_statement(c *Simultaneous_statementContext)

	// ExitSimultaneous_statement_part is called when exiting the simultaneous_statement_part production.
	ExitSimultaneous_statement_part(c *Simultaneous_statement_partContext)

	// ExitSource_aspect is called when exiting the source_aspect production.
	ExitSource_aspect(c *Source_aspectContext)

	// ExitSource_quantity_declaration is called when exiting the source_quantity_declaration production.
	ExitSource_quantity_declaration(c *Source_quantity_declarationContext)

	// ExitStep_limit_specification is called when exiting the step_limit_specification production.
	ExitStep_limit_specification(c *Step_limit_specificationContext)

	// ExitSubnature_declaration is called when exiting the subnature_declaration production.
	ExitSubnature_declaration(c *Subnature_declarationContext)

	// ExitSubnature_indication is called when exiting the subnature_indication production.
	ExitSubnature_indication(c *Subnature_indicationContext)

	// ExitSubprogram_body is called when exiting the subprogram_body production.
	ExitSubprogram_body(c *Subprogram_bodyContext)

	// ExitSubprogram_declaration is called when exiting the subprogram_declaration production.
	ExitSubprogram_declaration(c *Subprogram_declarationContext)

	// ExitSubprogram_declarative_item is called when exiting the subprogram_declarative_item production.
	ExitSubprogram_declarative_item(c *Subprogram_declarative_itemContext)

	// ExitSubprogram_declarative_part is called when exiting the subprogram_declarative_part production.
	ExitSubprogram_declarative_part(c *Subprogram_declarative_partContext)

	// ExitSubprogram_kind is called when exiting the subprogram_kind production.
	ExitSubprogram_kind(c *Subprogram_kindContext)

	// ExitSubprogram_specification is called when exiting the subprogram_specification production.
	ExitSubprogram_specification(c *Subprogram_specificationContext)

	// ExitProcedure_specification is called when exiting the procedure_specification production.
	ExitProcedure_specification(c *Procedure_specificationContext)

	// ExitFunction_specification is called when exiting the function_specification production.
	ExitFunction_specification(c *Function_specificationContext)

	// ExitSubprogram_statement_part is called when exiting the subprogram_statement_part production.
	ExitSubprogram_statement_part(c *Subprogram_statement_partContext)

	// ExitSubtype_declaration is called when exiting the subtype_declaration production.
	ExitSubtype_declaration(c *Subtype_declarationContext)

	// ExitSubtype_indication is called when exiting the subtype_indication production.
	ExitSubtype_indication(c *Subtype_indicationContext)

	// ExitSuffix is called when exiting the suffix production.
	ExitSuffix(c *SuffixContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTerminal_aspect is called when exiting the terminal_aspect production.
	ExitTerminal_aspect(c *Terminal_aspectContext)

	// ExitTerminal_declaration is called when exiting the terminal_declaration production.
	ExitTerminal_declaration(c *Terminal_declarationContext)

	// ExitThrough_aspect is called when exiting the through_aspect production.
	ExitThrough_aspect(c *Through_aspectContext)

	// ExitTimeout_clause is called when exiting the timeout_clause production.
	ExitTimeout_clause(c *Timeout_clauseContext)

	// ExitTolerance_aspect is called when exiting the tolerance_aspect production.
	ExitTolerance_aspect(c *Tolerance_aspectContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitType_definition is called when exiting the type_definition production.
	ExitType_definition(c *Type_definitionContext)

	// ExitUnconstrained_array_definition is called when exiting the unconstrained_array_definition production.
	ExitUnconstrained_array_definition(c *Unconstrained_array_definitionContext)

	// ExitUnconstrained_nature_definition is called when exiting the unconstrained_nature_definition production.
	ExitUnconstrained_nature_definition(c *Unconstrained_nature_definitionContext)

	// ExitUse_clause is called when exiting the use_clause production.
	ExitUse_clause(c *Use_clauseContext)

	// ExitVariable_assignment_statement is called when exiting the variable_assignment_statement production.
	ExitVariable_assignment_statement(c *Variable_assignment_statementContext)

	// ExitVariable_declaration is called when exiting the variable_declaration production.
	ExitVariable_declaration(c *Variable_declarationContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitWaveform is called when exiting the waveform production.
	ExitWaveform(c *WaveformContext)

	// ExitWaveform_element is called when exiting the waveform_element production.
	ExitWaveform_element(c *Waveform_elementContext)
}
