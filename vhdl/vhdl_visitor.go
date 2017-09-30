// Generated from vhdl.g4 by ANTLR 4.7.

package vhdl // vhdl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by vhdlParser.
type vhdlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by vhdlParser#abstract_literal.
	VisitAbstract_literal(ctx *Abstract_literalContext) interface{}

	// Visit a parse tree produced by vhdlParser#access_type_definition.
	VisitAccess_type_definition(ctx *Access_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#across_aspect.
	VisitAcross_aspect(ctx *Across_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#actual_designator.
	VisitActual_designator(ctx *Actual_designatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#actual_parameter_part.
	VisitActual_parameter_part(ctx *Actual_parameter_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#actual_part.
	VisitActual_part(ctx *Actual_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#adding_operator.
	VisitAdding_operator(ctx *Adding_operatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#aggregate.
	VisitAggregate(ctx *AggregateContext) interface{}

	// Visit a parse tree produced by vhdlParser#alias_declaration.
	VisitAlias_declaration(ctx *Alias_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#alias_designator.
	VisitAlias_designator(ctx *Alias_designatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#alias_indication.
	VisitAlias_indication(ctx *Alias_indicationContext) interface{}

	// Visit a parse tree produced by vhdlParser#allocator.
	VisitAllocator(ctx *AllocatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#architecture_body.
	VisitArchitecture_body(ctx *Architecture_bodyContext) interface{}

	// Visit a parse tree produced by vhdlParser#architecture_declarative_part.
	VisitArchitecture_declarative_part(ctx *Architecture_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#architecture_statement.
	VisitArchitecture_statement(ctx *Architecture_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#architecture_statement_part.
	VisitArchitecture_statement_part(ctx *Architecture_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#array_nature_definition.
	VisitArray_nature_definition(ctx *Array_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#array_type_definition.
	VisitArray_type_definition(ctx *Array_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#assertion.
	VisitAssertion(ctx *AssertionContext) interface{}

	// Visit a parse tree produced by vhdlParser#assertion_statement.
	VisitAssertion_statement(ctx *Assertion_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#association_element.
	VisitAssociation_element(ctx *Association_elementContext) interface{}

	// Visit a parse tree produced by vhdlParser#association_list.
	VisitAssociation_list(ctx *Association_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#attribute_declaration.
	VisitAttribute_declaration(ctx *Attribute_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#attribute_designator.
	VisitAttribute_designator(ctx *Attribute_designatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#attribute_specification.
	VisitAttribute_specification(ctx *Attribute_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#base_unit_declaration.
	VisitBase_unit_declaration(ctx *Base_unit_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#binding_indication.
	VisitBinding_indication(ctx *Binding_indicationContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_configuration.
	VisitBlock_configuration(ctx *Block_configurationContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_declarative_item.
	VisitBlock_declarative_item(ctx *Block_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_declarative_part.
	VisitBlock_declarative_part(ctx *Block_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_header.
	VisitBlock_header(ctx *Block_headerContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_specification.
	VisitBlock_specification(ctx *Block_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_statement.
	VisitBlock_statement(ctx *Block_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#block_statement_part.
	VisitBlock_statement_part(ctx *Block_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#branch_quantity_declaration.
	VisitBranch_quantity_declaration(ctx *Branch_quantity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#break_element.
	VisitBreak_element(ctx *Break_elementContext) interface{}

	// Visit a parse tree produced by vhdlParser#break_list.
	VisitBreak_list(ctx *Break_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#break_selector_clause.
	VisitBreak_selector_clause(ctx *Break_selector_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#case_statement_alternative.
	VisitCase_statement_alternative(ctx *Case_statement_alternativeContext) interface{}

	// Visit a parse tree produced by vhdlParser#choice.
	VisitChoice(ctx *ChoiceContext) interface{}

	// Visit a parse tree produced by vhdlParser#choices.
	VisitChoices(ctx *ChoicesContext) interface{}

	// Visit a parse tree produced by vhdlParser#component_configuration.
	VisitComponent_configuration(ctx *Component_configurationContext) interface{}

	// Visit a parse tree produced by vhdlParser#component_declaration.
	VisitComponent_declaration(ctx *Component_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#component_instantiation_statement.
	VisitComponent_instantiation_statement(ctx *Component_instantiation_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#component_specification.
	VisitComponent_specification(ctx *Component_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#composite_nature_definition.
	VisitComposite_nature_definition(ctx *Composite_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#composite_type_definition.
	VisitComposite_type_definition(ctx *Composite_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#concurrent_assertion_statement.
	VisitConcurrent_assertion_statement(ctx *Concurrent_assertion_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#concurrent_break_statement.
	VisitConcurrent_break_statement(ctx *Concurrent_break_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#concurrent_procedure_call_statement.
	VisitConcurrent_procedure_call_statement(ctx *Concurrent_procedure_call_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#concurrent_signal_assignment_statement.
	VisitConcurrent_signal_assignment_statement(ctx *Concurrent_signal_assignment_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by vhdlParser#condition_clause.
	VisitCondition_clause(ctx *Condition_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#conditional_signal_assignment.
	VisitConditional_signal_assignment(ctx *Conditional_signal_assignmentContext) interface{}

	// Visit a parse tree produced by vhdlParser#conditional_waveforms.
	VisitConditional_waveforms(ctx *Conditional_waveformsContext) interface{}

	// Visit a parse tree produced by vhdlParser#configuration_declaration.
	VisitConfiguration_declaration(ctx *Configuration_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#configuration_declarative_item.
	VisitConfiguration_declarative_item(ctx *Configuration_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#configuration_declarative_part.
	VisitConfiguration_declarative_part(ctx *Configuration_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#configuration_item.
	VisitConfiguration_item(ctx *Configuration_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#configuration_specification.
	VisitConfiguration_specification(ctx *Configuration_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#constant_declaration.
	VisitConstant_declaration(ctx *Constant_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#constrained_array_definition.
	VisitConstrained_array_definition(ctx *Constrained_array_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#constrained_nature_definition.
	VisitConstrained_nature_definition(ctx *Constrained_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by vhdlParser#context_clause.
	VisitContext_clause(ctx *Context_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#context_item.
	VisitContext_item(ctx *Context_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#delay_mechanism.
	VisitDelay_mechanism(ctx *Delay_mechanismContext) interface{}

	// Visit a parse tree produced by vhdlParser#design_file.
	VisitDesign_file(ctx *Design_fileContext) interface{}

	// Visit a parse tree produced by vhdlParser#design_unit.
	VisitDesign_unit(ctx *Design_unitContext) interface{}

	// Visit a parse tree produced by vhdlParser#designator.
	VisitDesignator(ctx *DesignatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#direction.
	VisitDirection(ctx *DirectionContext) interface{}

	// Visit a parse tree produced by vhdlParser#disconnection_specification.
	VisitDisconnection_specification(ctx *Disconnection_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#discrete_range.
	VisitDiscrete_range(ctx *Discrete_rangeContext) interface{}

	// Visit a parse tree produced by vhdlParser#element_association.
	VisitElement_association(ctx *Element_associationContext) interface{}

	// Visit a parse tree produced by vhdlParser#element_declaration.
	VisitElement_declaration(ctx *Element_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#element_subnature_definition.
	VisitElement_subnature_definition(ctx *Element_subnature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#element_subtype_definition.
	VisitElement_subtype_definition(ctx *Element_subtype_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_aspect.
	VisitEntity_aspect(ctx *Entity_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_class.
	VisitEntity_class(ctx *Entity_classContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_class_entry.
	VisitEntity_class_entry(ctx *Entity_class_entryContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_class_entry_list.
	VisitEntity_class_entry_list(ctx *Entity_class_entry_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_declaration.
	VisitEntity_declaration(ctx *Entity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_declarative_item.
	VisitEntity_declarative_item(ctx *Entity_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_declarative_part.
	VisitEntity_declarative_part(ctx *Entity_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_designator.
	VisitEntity_designator(ctx *Entity_designatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_header.
	VisitEntity_header(ctx *Entity_headerContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_name_list.
	VisitEntity_name_list(ctx *Entity_name_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_specification.
	VisitEntity_specification(ctx *Entity_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_statement.
	VisitEntity_statement(ctx *Entity_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_statement_part.
	VisitEntity_statement_part(ctx *Entity_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#entity_tag.
	VisitEntity_tag(ctx *Entity_tagContext) interface{}

	// Visit a parse tree produced by vhdlParser#enumeration_literal.
	VisitEnumeration_literal(ctx *Enumeration_literalContext) interface{}

	// Visit a parse tree produced by vhdlParser#enumeration_type_definition.
	VisitEnumeration_type_definition(ctx *Enumeration_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#exit_statement.
	VisitExit_statement(ctx *Exit_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by vhdlParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by vhdlParser#file_declaration.
	VisitFile_declaration(ctx *File_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#file_logical_name.
	VisitFile_logical_name(ctx *File_logical_nameContext) interface{}

	// Visit a parse tree produced by vhdlParser#file_open_information.
	VisitFile_open_information(ctx *File_open_informationContext) interface{}

	// Visit a parse tree produced by vhdlParser#file_type_definition.
	VisitFile_type_definition(ctx *File_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#formal_parameter_list.
	VisitFormal_parameter_list(ctx *Formal_parameter_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#formal_part.
	VisitFormal_part(ctx *Formal_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#free_quantity_declaration.
	VisitFree_quantity_declaration(ctx *Free_quantity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#generate_statement.
	VisitGenerate_statement(ctx *Generate_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#generation_scheme.
	VisitGeneration_scheme(ctx *Generation_schemeContext) interface{}

	// Visit a parse tree produced by vhdlParser#generic_clause.
	VisitGeneric_clause(ctx *Generic_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#generic_list.
	VisitGeneric_list(ctx *Generic_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#generic_map_aspect.
	VisitGeneric_map_aspect(ctx *Generic_map_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#group_constituent.
	VisitGroup_constituent(ctx *Group_constituentContext) interface{}

	// Visit a parse tree produced by vhdlParser#group_constituent_list.
	VisitGroup_constituent_list(ctx *Group_constituent_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#group_declaration.
	VisitGroup_declaration(ctx *Group_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#group_template_declaration.
	VisitGroup_template_declaration(ctx *Group_template_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#guarded_signal_specification.
	VisitGuarded_signal_specification(ctx *Guarded_signal_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by vhdlParser#identifier_list.
	VisitIdentifier_list(ctx *Identifier_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#index_constraint.
	VisitIndex_constraint(ctx *Index_constraintContext) interface{}

	// Visit a parse tree produced by vhdlParser#index_specification.
	VisitIndex_specification(ctx *Index_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#index_subtype_definition.
	VisitIndex_subtype_definition(ctx *Index_subtype_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#instantiated_unit.
	VisitInstantiated_unit(ctx *Instantiated_unitContext) interface{}

	// Visit a parse tree produced by vhdlParser#instantiation_list.
	VisitInstantiation_list(ctx *Instantiation_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_constant_declaration.
	VisitInterface_constant_declaration(ctx *Interface_constant_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_declaration.
	VisitInterface_declaration(ctx *Interface_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_element.
	VisitInterface_element(ctx *Interface_elementContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_file_declaration.
	VisitInterface_file_declaration(ctx *Interface_file_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_signal_list.
	VisitInterface_signal_list(ctx *Interface_signal_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_port_list.
	VisitInterface_port_list(ctx *Interface_port_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_list.
	VisitInterface_list(ctx *Interface_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_quantity_declaration.
	VisitInterface_quantity_declaration(ctx *Interface_quantity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_port_declaration.
	VisitInterface_port_declaration(ctx *Interface_port_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_signal_declaration.
	VisitInterface_signal_declaration(ctx *Interface_signal_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_terminal_declaration.
	VisitInterface_terminal_declaration(ctx *Interface_terminal_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#interface_variable_declaration.
	VisitInterface_variable_declaration(ctx *Interface_variable_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#iteration_scheme.
	VisitIteration_scheme(ctx *Iteration_schemeContext) interface{}

	// Visit a parse tree produced by vhdlParser#label_colon.
	VisitLabel_colon(ctx *Label_colonContext) interface{}

	// Visit a parse tree produced by vhdlParser#library_clause.
	VisitLibrary_clause(ctx *Library_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#library_unit.
	VisitLibrary_unit(ctx *Library_unitContext) interface{}

	// Visit a parse tree produced by vhdlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by vhdlParser#logical_name.
	VisitLogical_name(ctx *Logical_nameContext) interface{}

	// Visit a parse tree produced by vhdlParser#logical_name_list.
	VisitLogical_name_list(ctx *Logical_name_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#logical_operator.
	VisitLogical_operator(ctx *Logical_operatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#signal_mode.
	VisitSignal_mode(ctx *Signal_modeContext) interface{}

	// Visit a parse tree produced by vhdlParser#multiplying_operator.
	VisitMultiplying_operator(ctx *Multiplying_operatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by vhdlParser#name_part.
	VisitName_part(ctx *Name_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#name_attribute_part.
	VisitName_attribute_part(ctx *Name_attribute_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#name_function_call_or_indexed_part.
	VisitName_function_call_or_indexed_part(ctx *Name_function_call_or_indexed_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#name_slice_part.
	VisitName_slice_part(ctx *Name_slice_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#selected_name.
	VisitSelected_name(ctx *Selected_nameContext) interface{}

	// Visit a parse tree produced by vhdlParser#nature_declaration.
	VisitNature_declaration(ctx *Nature_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#nature_definition.
	VisitNature_definition(ctx *Nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#nature_element_declaration.
	VisitNature_element_declaration(ctx *Nature_element_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#next_statement.
	VisitNext_statement(ctx *Next_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#numeric_literal.
	VisitNumeric_literal(ctx *Numeric_literalContext) interface{}

	// Visit a parse tree produced by vhdlParser#object_declaration.
	VisitObject_declaration(ctx *Object_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#opts.
	VisitOpts(ctx *OptsContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_body.
	VisitPackage_body(ctx *Package_bodyContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_body_declarative_item.
	VisitPackage_body_declarative_item(ctx *Package_body_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_body_declarative_part.
	VisitPackage_body_declarative_part(ctx *Package_body_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_declaration.
	VisitPackage_declaration(ctx *Package_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_declarative_item.
	VisitPackage_declarative_item(ctx *Package_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#package_declarative_part.
	VisitPackage_declarative_part(ctx *Package_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#parameter_specification.
	VisitParameter_specification(ctx *Parameter_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#physical_literal.
	VisitPhysical_literal(ctx *Physical_literalContext) interface{}

	// Visit a parse tree produced by vhdlParser#physical_type_definition.
	VisitPhysical_type_definition(ctx *Physical_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#port_clause.
	VisitPort_clause(ctx *Port_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#port_list.
	VisitPort_list(ctx *Port_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#port_map_aspect.
	VisitPort_map_aspect(ctx *Port_map_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by vhdlParser#primary_unit.
	VisitPrimary_unit(ctx *Primary_unitContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedural_declarative_item.
	VisitProcedural_declarative_item(ctx *Procedural_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedural_declarative_part.
	VisitProcedural_declarative_part(ctx *Procedural_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedural_statement_part.
	VisitProcedural_statement_part(ctx *Procedural_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedure_call.
	VisitProcedure_call(ctx *Procedure_callContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedure_call_statement.
	VisitProcedure_call_statement(ctx *Procedure_call_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#process_declarative_item.
	VisitProcess_declarative_item(ctx *Process_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#process_declarative_part.
	VisitProcess_declarative_part(ctx *Process_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#process_statement.
	VisitProcess_statement(ctx *Process_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#process_statement_part.
	VisitProcess_statement_part(ctx *Process_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#qualified_expression.
	VisitQualified_expression(ctx *Qualified_expressionContext) interface{}

	// Visit a parse tree produced by vhdlParser#quantity_declaration.
	VisitQuantity_declaration(ctx *Quantity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#quantity_list.
	VisitQuantity_list(ctx *Quantity_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#quantity_specification.
	VisitQuantity_specification(ctx *Quantity_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#range_decl.
	VisitRange_decl(ctx *Range_declContext) interface{}

	// Visit a parse tree produced by vhdlParser#explicit_range.
	VisitExplicit_range(ctx *Explicit_rangeContext) interface{}

	// Visit a parse tree produced by vhdlParser#range_constraint.
	VisitRange_constraint(ctx *Range_constraintContext) interface{}

	// Visit a parse tree produced by vhdlParser#record_nature_definition.
	VisitRecord_nature_definition(ctx *Record_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#record_type_definition.
	VisitRecord_type_definition(ctx *Record_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by vhdlParser#relational_operator.
	VisitRelational_operator(ctx *Relational_operatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#report_statement.
	VisitReport_statement(ctx *Report_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#scalar_nature_definition.
	VisitScalar_nature_definition(ctx *Scalar_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#scalar_type_definition.
	VisitScalar_type_definition(ctx *Scalar_type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#secondary_unit.
	VisitSecondary_unit(ctx *Secondary_unitContext) interface{}

	// Visit a parse tree produced by vhdlParser#secondary_unit_declaration.
	VisitSecondary_unit_declaration(ctx *Secondary_unit_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#selected_signal_assignment.
	VisitSelected_signal_assignment(ctx *Selected_signal_assignmentContext) interface{}

	// Visit a parse tree produced by vhdlParser#selected_waveforms.
	VisitSelected_waveforms(ctx *Selected_waveformsContext) interface{}

	// Visit a parse tree produced by vhdlParser#sensitivity_clause.
	VisitSensitivity_clause(ctx *Sensitivity_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#sensitivity_list.
	VisitSensitivity_list(ctx *Sensitivity_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#sequence_of_statements.
	VisitSequence_of_statements(ctx *Sequence_of_statementsContext) interface{}

	// Visit a parse tree produced by vhdlParser#sequential_statement.
	VisitSequential_statement(ctx *Sequential_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#shift_expression.
	VisitShift_expression(ctx *Shift_expressionContext) interface{}

	// Visit a parse tree produced by vhdlParser#shift_operator.
	VisitShift_operator(ctx *Shift_operatorContext) interface{}

	// Visit a parse tree produced by vhdlParser#signal_assignment_statement.
	VisitSignal_assignment_statement(ctx *Signal_assignment_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#signal_declaration.
	VisitSignal_declaration(ctx *Signal_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#signal_kind.
	VisitSignal_kind(ctx *Signal_kindContext) interface{}

	// Visit a parse tree produced by vhdlParser#signal_list.
	VisitSignal_list(ctx *Signal_listContext) interface{}

	// Visit a parse tree produced by vhdlParser#signature.
	VisitSignature(ctx *SignatureContext) interface{}

	// Visit a parse tree produced by vhdlParser#simple_expression.
	VisitSimple_expression(ctx *Simple_expressionContext) interface{}

	// Visit a parse tree produced by vhdlParser#simple_simultaneous_statement.
	VisitSimple_simultaneous_statement(ctx *Simple_simultaneous_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_alternative.
	VisitSimultaneous_alternative(ctx *Simultaneous_alternativeContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_case_statement.
	VisitSimultaneous_case_statement(ctx *Simultaneous_case_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_if_statement.
	VisitSimultaneous_if_statement(ctx *Simultaneous_if_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_procedural_statement.
	VisitSimultaneous_procedural_statement(ctx *Simultaneous_procedural_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_statement.
	VisitSimultaneous_statement(ctx *Simultaneous_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#simultaneous_statement_part.
	VisitSimultaneous_statement_part(ctx *Simultaneous_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#source_aspect.
	VisitSource_aspect(ctx *Source_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#source_quantity_declaration.
	VisitSource_quantity_declaration(ctx *Source_quantity_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#step_limit_specification.
	VisitStep_limit_specification(ctx *Step_limit_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subnature_declaration.
	VisitSubnature_declaration(ctx *Subnature_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subnature_indication.
	VisitSubnature_indication(ctx *Subnature_indicationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_body.
	VisitSubprogram_body(ctx *Subprogram_bodyContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_declaration.
	VisitSubprogram_declaration(ctx *Subprogram_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_declarative_item.
	VisitSubprogram_declarative_item(ctx *Subprogram_declarative_itemContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_declarative_part.
	VisitSubprogram_declarative_part(ctx *Subprogram_declarative_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_kind.
	VisitSubprogram_kind(ctx *Subprogram_kindContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_specification.
	VisitSubprogram_specification(ctx *Subprogram_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#procedure_specification.
	VisitProcedure_specification(ctx *Procedure_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#function_specification.
	VisitFunction_specification(ctx *Function_specificationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subprogram_statement_part.
	VisitSubprogram_statement_part(ctx *Subprogram_statement_partContext) interface{}

	// Visit a parse tree produced by vhdlParser#subtype_declaration.
	VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#subtype_indication.
	VisitSubtype_indication(ctx *Subtype_indicationContext) interface{}

	// Visit a parse tree produced by vhdlParser#suffix.
	VisitSuffix(ctx *SuffixContext) interface{}

	// Visit a parse tree produced by vhdlParser#target.
	VisitTarget(ctx *TargetContext) interface{}

	// Visit a parse tree produced by vhdlParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by vhdlParser#terminal_aspect.
	VisitTerminal_aspect(ctx *Terminal_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#terminal_declaration.
	VisitTerminal_declaration(ctx *Terminal_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#through_aspect.
	VisitThrough_aspect(ctx *Through_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#timeout_clause.
	VisitTimeout_clause(ctx *Timeout_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#tolerance_aspect.
	VisitTolerance_aspect(ctx *Tolerance_aspectContext) interface{}

	// Visit a parse tree produced by vhdlParser#type_declaration.
	VisitType_declaration(ctx *Type_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#type_definition.
	VisitType_definition(ctx *Type_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#unconstrained_array_definition.
	VisitUnconstrained_array_definition(ctx *Unconstrained_array_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#unconstrained_nature_definition.
	VisitUnconstrained_nature_definition(ctx *Unconstrained_nature_definitionContext) interface{}

	// Visit a parse tree produced by vhdlParser#use_clause.
	VisitUse_clause(ctx *Use_clauseContext) interface{}

	// Visit a parse tree produced by vhdlParser#variable_assignment_statement.
	VisitVariable_assignment_statement(ctx *Variable_assignment_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#variable_declaration.
	VisitVariable_declaration(ctx *Variable_declarationContext) interface{}

	// Visit a parse tree produced by vhdlParser#wait_statement.
	VisitWait_statement(ctx *Wait_statementContext) interface{}

	// Visit a parse tree produced by vhdlParser#waveform.
	VisitWaveform(ctx *WaveformContext) interface{}

	// Visit a parse tree produced by vhdlParser#waveform_element.
	VisitWaveform_element(ctx *Waveform_elementContext) interface{}
}
