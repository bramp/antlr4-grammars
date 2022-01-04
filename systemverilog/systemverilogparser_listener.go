// Code generated from SystemVerilogParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package systemverilog // SystemVerilogParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SystemVerilogParserListener is a complete listener for a parse tree produced by SystemVerilogParser.
type SystemVerilogParserListener interface {
	antlr.ParseTreeListener

	// EnterLibrary_text is called when entering the library_text production.
	EnterLibrary_text(c *Library_textContext)

	// EnterLibrary_description is called when entering the library_description production.
	EnterLibrary_description(c *Library_descriptionContext)

	// EnterLibrary_declaration is called when entering the library_declaration production.
	EnterLibrary_declaration(c *Library_declarationContext)

	// EnterInclude_statement is called when entering the include_statement production.
	EnterInclude_statement(c *Include_statementContext)

	// EnterSource_text is called when entering the source_text production.
	EnterSource_text(c *Source_textContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterModule_nonansi_header is called when entering the module_nonansi_header production.
	EnterModule_nonansi_header(c *Module_nonansi_headerContext)

	// EnterModule_ansi_header is called when entering the module_ansi_header production.
	EnterModule_ansi_header(c *Module_ansi_headerContext)

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterModule_keyword is called when entering the module_keyword production.
	EnterModule_keyword(c *Module_keywordContext)

	// EnterInterface_declaration is called when entering the interface_declaration production.
	EnterInterface_declaration(c *Interface_declarationContext)

	// EnterInterface_nonansi_header is called when entering the interface_nonansi_header production.
	EnterInterface_nonansi_header(c *Interface_nonansi_headerContext)

	// EnterInterface_ansi_header is called when entering the interface_ansi_header production.
	EnterInterface_ansi_header(c *Interface_ansi_headerContext)

	// EnterProgram_declaration is called when entering the program_declaration production.
	EnterProgram_declaration(c *Program_declarationContext)

	// EnterProgram_nonansi_header is called when entering the program_nonansi_header production.
	EnterProgram_nonansi_header(c *Program_nonansi_headerContext)

	// EnterProgram_ansi_header is called when entering the program_ansi_header production.
	EnterProgram_ansi_header(c *Program_ansi_headerContext)

	// EnterChecker_declaration is called when entering the checker_declaration production.
	EnterChecker_declaration(c *Checker_declarationContext)

	// EnterClass_declaration is called when entering the class_declaration production.
	EnterClass_declaration(c *Class_declarationContext)

	// EnterInterface_class_type is called when entering the interface_class_type production.
	EnterInterface_class_type(c *Interface_class_typeContext)

	// EnterInterface_class_declaration is called when entering the interface_class_declaration production.
	EnterInterface_class_declaration(c *Interface_class_declarationContext)

	// EnterInterface_class_item is called when entering the interface_class_item production.
	EnterInterface_class_item(c *Interface_class_itemContext)

	// EnterInterface_class_method is called when entering the interface_class_method production.
	EnterInterface_class_method(c *Interface_class_methodContext)

	// EnterPackage_declaration is called when entering the package_declaration production.
	EnterPackage_declaration(c *Package_declarationContext)

	// EnterTimeunits_declaration is called when entering the timeunits_declaration production.
	EnterTimeunits_declaration(c *Timeunits_declarationContext)

	// EnterParameter_port_list is called when entering the parameter_port_list production.
	EnterParameter_port_list(c *Parameter_port_listContext)

	// EnterParameter_port_declaration is called when entering the parameter_port_declaration production.
	EnterParameter_port_declaration(c *Parameter_port_declarationContext)

	// EnterList_of_ports is called when entering the list_of_ports production.
	EnterList_of_ports(c *List_of_portsContext)

	// EnterList_of_port_declarations is called when entering the list_of_port_declarations production.
	EnterList_of_port_declarations(c *List_of_port_declarationsContext)

	// EnterPort_declaration is called when entering the port_declaration production.
	EnterPort_declaration(c *Port_declarationContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterPort_expression is called when entering the port_expression production.
	EnterPort_expression(c *Port_expressionContext)

	// EnterPort_reference is called when entering the port_reference production.
	EnterPort_reference(c *Port_referenceContext)

	// EnterPort_direction is called when entering the port_direction production.
	EnterPort_direction(c *Port_directionContext)

	// EnterNet_port_header is called when entering the net_port_header production.
	EnterNet_port_header(c *Net_port_headerContext)

	// EnterVariable_port_header is called when entering the variable_port_header production.
	EnterVariable_port_header(c *Variable_port_headerContext)

	// EnterInterface_port_header is called when entering the interface_port_header production.
	EnterInterface_port_header(c *Interface_port_headerContext)

	// EnterAnsi_port_declaration is called when entering the ansi_port_declaration production.
	EnterAnsi_port_declaration(c *Ansi_port_declarationContext)

	// EnterElaboration_system_task is called when entering the elaboration_system_task production.
	EnterElaboration_system_task(c *Elaboration_system_taskContext)

	// EnterModule_common_item is called when entering the module_common_item production.
	EnterModule_common_item(c *Module_common_itemContext)

	// EnterModule_item is called when entering the module_item production.
	EnterModule_item(c *Module_itemContext)

	// EnterModule_or_generate_item is called when entering the module_or_generate_item production.
	EnterModule_or_generate_item(c *Module_or_generate_itemContext)

	// EnterModule_or_generate_item_declaration is called when entering the module_or_generate_item_declaration production.
	EnterModule_or_generate_item_declaration(c *Module_or_generate_item_declarationContext)

	// EnterNon_port_module_item is called when entering the non_port_module_item production.
	EnterNon_port_module_item(c *Non_port_module_itemContext)

	// EnterParameter_override is called when entering the parameter_override production.
	EnterParameter_override(c *Parameter_overrideContext)

	// EnterBind_directive is called when entering the bind_directive production.
	EnterBind_directive(c *Bind_directiveContext)

	// EnterBind_target_scope is called when entering the bind_target_scope production.
	EnterBind_target_scope(c *Bind_target_scopeContext)

	// EnterBind_target_instance is called when entering the bind_target_instance production.
	EnterBind_target_instance(c *Bind_target_instanceContext)

	// EnterBind_target_instance_list is called when entering the bind_target_instance_list production.
	EnterBind_target_instance_list(c *Bind_target_instance_listContext)

	// EnterBind_instantiation is called when entering the bind_instantiation production.
	EnterBind_instantiation(c *Bind_instantiationContext)

	// EnterConfig_declaration is called when entering the config_declaration production.
	EnterConfig_declaration(c *Config_declarationContext)

	// EnterDesign_statement is called when entering the design_statement production.
	EnterDesign_statement(c *Design_statementContext)

	// EnterConfig_rule_statement is called when entering the config_rule_statement production.
	EnterConfig_rule_statement(c *Config_rule_statementContext)

	// EnterDefault_clause is called when entering the default_clause production.
	EnterDefault_clause(c *Default_clauseContext)

	// EnterInst_clause is called when entering the inst_clause production.
	EnterInst_clause(c *Inst_clauseContext)

	// EnterInst_name is called when entering the inst_name production.
	EnterInst_name(c *Inst_nameContext)

	// EnterCell_clause is called when entering the cell_clause production.
	EnterCell_clause(c *Cell_clauseContext)

	// EnterLiblist_clause is called when entering the liblist_clause production.
	EnterLiblist_clause(c *Liblist_clauseContext)

	// EnterUse_clause is called when entering the use_clause production.
	EnterUse_clause(c *Use_clauseContext)

	// EnterInterface_or_generate_item is called when entering the interface_or_generate_item production.
	EnterInterface_or_generate_item(c *Interface_or_generate_itemContext)

	// EnterExtern_tf_declaration is called when entering the extern_tf_declaration production.
	EnterExtern_tf_declaration(c *Extern_tf_declarationContext)

	// EnterInterface_item is called when entering the interface_item production.
	EnterInterface_item(c *Interface_itemContext)

	// EnterNon_port_interface_item is called when entering the non_port_interface_item production.
	EnterNon_port_interface_item(c *Non_port_interface_itemContext)

	// EnterProgram_item is called when entering the program_item production.
	EnterProgram_item(c *Program_itemContext)

	// EnterNon_port_program_item is called when entering the non_port_program_item production.
	EnterNon_port_program_item(c *Non_port_program_itemContext)

	// EnterProgram_generate_item is called when entering the program_generate_item production.
	EnterProgram_generate_item(c *Program_generate_itemContext)

	// EnterChecker_port_list is called when entering the checker_port_list production.
	EnterChecker_port_list(c *Checker_port_listContext)

	// EnterChecker_port_item is called when entering the checker_port_item production.
	EnterChecker_port_item(c *Checker_port_itemContext)

	// EnterChecker_port_direction is called when entering the checker_port_direction production.
	EnterChecker_port_direction(c *Checker_port_directionContext)

	// EnterChecker_or_generate_item is called when entering the checker_or_generate_item production.
	EnterChecker_or_generate_item(c *Checker_or_generate_itemContext)

	// EnterChecker_or_generate_item_declaration is called when entering the checker_or_generate_item_declaration production.
	EnterChecker_or_generate_item_declaration(c *Checker_or_generate_item_declarationContext)

	// EnterChecker_generate_item is called when entering the checker_generate_item production.
	EnterChecker_generate_item(c *Checker_generate_itemContext)

	// EnterClass_item is called when entering the class_item production.
	EnterClass_item(c *Class_itemContext)

	// EnterClass_property is called when entering the class_property production.
	EnterClass_property(c *Class_propertyContext)

	// EnterClass_method is called when entering the class_method production.
	EnterClass_method(c *Class_methodContext)

	// EnterClass_constructor_prototype is called when entering the class_constructor_prototype production.
	EnterClass_constructor_prototype(c *Class_constructor_prototypeContext)

	// EnterClass_constraint is called when entering the class_constraint production.
	EnterClass_constraint(c *Class_constraintContext)

	// EnterClass_item_qualifier is called when entering the class_item_qualifier production.
	EnterClass_item_qualifier(c *Class_item_qualifierContext)

	// EnterProperty_qualifier is called when entering the property_qualifier production.
	EnterProperty_qualifier(c *Property_qualifierContext)

	// EnterRandom_qualifier is called when entering the random_qualifier production.
	EnterRandom_qualifier(c *Random_qualifierContext)

	// EnterMethod_qualifier is called when entering the method_qualifier production.
	EnterMethod_qualifier(c *Method_qualifierContext)

	// EnterMethod_prototype is called when entering the method_prototype production.
	EnterMethod_prototype(c *Method_prototypeContext)

	// EnterClass_constructor_declaration is called when entering the class_constructor_declaration production.
	EnterClass_constructor_declaration(c *Class_constructor_declarationContext)

	// EnterConstraint_declaration is called when entering the constraint_declaration production.
	EnterConstraint_declaration(c *Constraint_declarationContext)

	// EnterConstraint_block is called when entering the constraint_block production.
	EnterConstraint_block(c *Constraint_blockContext)

	// EnterConstraint_block_item is called when entering the constraint_block_item production.
	EnterConstraint_block_item(c *Constraint_block_itemContext)

	// EnterSolve_before_list is called when entering the solve_before_list production.
	EnterSolve_before_list(c *Solve_before_listContext)

	// EnterConstraint_primary is called when entering the constraint_primary production.
	EnterConstraint_primary(c *Constraint_primaryContext)

	// EnterConstraint_expression is called when entering the constraint_expression production.
	EnterConstraint_expression(c *Constraint_expressionContext)

	// EnterUniqueness_constraint is called when entering the uniqueness_constraint production.
	EnterUniqueness_constraint(c *Uniqueness_constraintContext)

	// EnterConstraint_set is called when entering the constraint_set production.
	EnterConstraint_set(c *Constraint_setContext)

	// EnterDist_list is called when entering the dist_list production.
	EnterDist_list(c *Dist_listContext)

	// EnterDist_item is called when entering the dist_item production.
	EnterDist_item(c *Dist_itemContext)

	// EnterDist_weight is called when entering the dist_weight production.
	EnterDist_weight(c *Dist_weightContext)

	// EnterConstraint_prototype is called when entering the constraint_prototype production.
	EnterConstraint_prototype(c *Constraint_prototypeContext)

	// EnterConstraint_prototype_qualifier is called when entering the constraint_prototype_qualifier production.
	EnterConstraint_prototype_qualifier(c *Constraint_prototype_qualifierContext)

	// EnterExtern_constraint_declaration is called when entering the extern_constraint_declaration production.
	EnterExtern_constraint_declaration(c *Extern_constraint_declarationContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterPackage_item is called when entering the package_item production.
	EnterPackage_item(c *Package_itemContext)

	// EnterPackage_or_generate_item_declaration is called when entering the package_or_generate_item_declaration production.
	EnterPackage_or_generate_item_declaration(c *Package_or_generate_item_declarationContext)

	// EnterAnonymous_program is called when entering the anonymous_program production.
	EnterAnonymous_program(c *Anonymous_programContext)

	// EnterAnonymous_program_item is called when entering the anonymous_program_item production.
	EnterAnonymous_program_item(c *Anonymous_program_itemContext)

	// EnterLocal_parameter_declaration is called when entering the local_parameter_declaration production.
	EnterLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterSpecparam_declaration is called when entering the specparam_declaration production.
	EnterSpecparam_declaration(c *Specparam_declarationContext)

	// EnterInout_declaration is called when entering the inout_declaration production.
	EnterInout_declaration(c *Inout_declarationContext)

	// EnterInput_declaration is called when entering the input_declaration production.
	EnterInput_declaration(c *Input_declarationContext)

	// EnterOutput_declaration is called when entering the output_declaration production.
	EnterOutput_declaration(c *Output_declarationContext)

	// EnterInterface_port_declaration is called when entering the interface_port_declaration production.
	EnterInterface_port_declaration(c *Interface_port_declarationContext)

	// EnterRef_declaration is called when entering the ref_declaration production.
	EnterRef_declaration(c *Ref_declarationContext)

	// EnterData_declaration is called when entering the data_declaration production.
	EnterData_declaration(c *Data_declarationContext)

	// EnterPackage_import_declaration is called when entering the package_import_declaration production.
	EnterPackage_import_declaration(c *Package_import_declarationContext)

	// EnterPackage_import_item is called when entering the package_import_item production.
	EnterPackage_import_item(c *Package_import_itemContext)

	// EnterPackage_export_declaration is called when entering the package_export_declaration production.
	EnterPackage_export_declaration(c *Package_export_declarationContext)

	// EnterGenvar_declaration is called when entering the genvar_declaration production.
	EnterGenvar_declaration(c *Genvar_declarationContext)

	// EnterNet_declaration is called when entering the net_declaration production.
	EnterNet_declaration(c *Net_declarationContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterNet_type_declaration is called when entering the net_type_declaration production.
	EnterNet_type_declaration(c *Net_type_declarationContext)

	// EnterLifetime is called when entering the lifetime production.
	EnterLifetime(c *LifetimeContext)

	// EnterCasting_type is called when entering the casting_type production.
	EnterCasting_type(c *Casting_typeContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterData_type_or_implicit is called when entering the data_type_or_implicit production.
	EnterData_type_or_implicit(c *Data_type_or_implicitContext)

	// EnterImplicit_data_type is called when entering the implicit_data_type production.
	EnterImplicit_data_type(c *Implicit_data_typeContext)

	// EnterEnum_base_type is called when entering the enum_base_type production.
	EnterEnum_base_type(c *Enum_base_typeContext)

	// EnterEnum_name_declaration is called when entering the enum_name_declaration production.
	EnterEnum_name_declaration(c *Enum_name_declarationContext)

	// EnterClass_scope is called when entering the class_scope production.
	EnterClass_scope(c *Class_scopeContext)

	// EnterClass_type is called when entering the class_type production.
	EnterClass_type(c *Class_typeContext)

	// EnterInteger_type is called when entering the integer_type production.
	EnterInteger_type(c *Integer_typeContext)

	// EnterInteger_atom_type is called when entering the integer_atom_type production.
	EnterInteger_atom_type(c *Integer_atom_typeContext)

	// EnterInteger_vector_type is called when entering the integer_vector_type production.
	EnterInteger_vector_type(c *Integer_vector_typeContext)

	// EnterNon_integer_type is called when entering the non_integer_type production.
	EnterNon_integer_type(c *Non_integer_typeContext)

	// EnterNet_type is called when entering the net_type production.
	EnterNet_type(c *Net_typeContext)

	// EnterNet_port_type is called when entering the net_port_type production.
	EnterNet_port_type(c *Net_port_typeContext)

	// EnterVariable_port_type is called when entering the variable_port_type production.
	EnterVariable_port_type(c *Variable_port_typeContext)

	// EnterVar_data_type is called when entering the var_data_type production.
	EnterVar_data_type(c *Var_data_typeContext)

	// EnterSigning is called when entering the signing production.
	EnterSigning(c *SigningContext)

	// EnterSimple_type is called when entering the simple_type production.
	EnterSimple_type(c *Simple_typeContext)

	// EnterStruct_union_member is called when entering the struct_union_member production.
	EnterStruct_union_member(c *Struct_union_memberContext)

	// EnterData_type_or_void is called when entering the data_type_or_void production.
	EnterData_type_or_void(c *Data_type_or_voidContext)

	// EnterStruct_union is called when entering the struct_union production.
	EnterStruct_union(c *Struct_unionContext)

	// EnterType_reference is called when entering the type_reference production.
	EnterType_reference(c *Type_referenceContext)

	// EnterDrive_strength is called when entering the drive_strength production.
	EnterDrive_strength(c *Drive_strengthContext)

	// EnterStrength0 is called when entering the strength0 production.
	EnterStrength0(c *Strength0Context)

	// EnterStrength1 is called when entering the strength1 production.
	EnterStrength1(c *Strength1Context)

	// EnterCharge_strength is called when entering the charge_strength production.
	EnterCharge_strength(c *Charge_strengthContext)

	// EnterDelay3 is called when entering the delay3 production.
	EnterDelay3(c *Delay3Context)

	// EnterDelay2 is called when entering the delay2 production.
	EnterDelay2(c *Delay2Context)

	// EnterDelay_value is called when entering the delay_value production.
	EnterDelay_value(c *Delay_valueContext)

	// EnterList_of_defparam_assignments is called when entering the list_of_defparam_assignments production.
	EnterList_of_defparam_assignments(c *List_of_defparam_assignmentsContext)

	// EnterList_of_genvar_identifiers is called when entering the list_of_genvar_identifiers production.
	EnterList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// EnterList_of_interface_identifiers is called when entering the list_of_interface_identifiers production.
	EnterList_of_interface_identifiers(c *List_of_interface_identifiersContext)

	// EnterList_of_net_decl_assignments is called when entering the list_of_net_decl_assignments production.
	EnterList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// EnterList_of_param_assignments is called when entering the list_of_param_assignments production.
	EnterList_of_param_assignments(c *List_of_param_assignmentsContext)

	// EnterList_of_port_identifiers is called when entering the list_of_port_identifiers production.
	EnterList_of_port_identifiers(c *List_of_port_identifiersContext)

	// EnterList_of_udp_port_identifiers is called when entering the list_of_udp_port_identifiers production.
	EnterList_of_udp_port_identifiers(c *List_of_udp_port_identifiersContext)

	// EnterList_of_specparam_assignments is called when entering the list_of_specparam_assignments production.
	EnterList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// EnterList_of_tf_variable_identifiers is called when entering the list_of_tf_variable_identifiers production.
	EnterList_of_tf_variable_identifiers(c *List_of_tf_variable_identifiersContext)

	// EnterList_of_type_assignments is called when entering the list_of_type_assignments production.
	EnterList_of_type_assignments(c *List_of_type_assignmentsContext)

	// EnterList_of_variable_decl_assignments is called when entering the list_of_variable_decl_assignments production.
	EnterList_of_variable_decl_assignments(c *List_of_variable_decl_assignmentsContext)

	// EnterList_of_variable_identifiers is called when entering the list_of_variable_identifiers production.
	EnterList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// EnterList_of_variable_port_identifiers is called when entering the list_of_variable_port_identifiers production.
	EnterList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// EnterDefparam_assignment is called when entering the defparam_assignment production.
	EnterDefparam_assignment(c *Defparam_assignmentContext)

	// EnterNet_decl_assignment is called when entering the net_decl_assignment production.
	EnterNet_decl_assignment(c *Net_decl_assignmentContext)

	// EnterParam_assignment is called when entering the param_assignment production.
	EnterParam_assignment(c *Param_assignmentContext)

	// EnterSpecparam_assignment is called when entering the specparam_assignment production.
	EnterSpecparam_assignment(c *Specparam_assignmentContext)

	// EnterType_assignment is called when entering the type_assignment production.
	EnterType_assignment(c *Type_assignmentContext)

	// EnterPulse_control_specparam is called when entering the pulse_control_specparam production.
	EnterPulse_control_specparam(c *Pulse_control_specparamContext)

	// EnterError_limit_value is called when entering the error_limit_value production.
	EnterError_limit_value(c *Error_limit_valueContext)

	// EnterReject_limit_value is called when entering the reject_limit_value production.
	EnterReject_limit_value(c *Reject_limit_valueContext)

	// EnterLimit_value is called when entering the limit_value production.
	EnterLimit_value(c *Limit_valueContext)

	// EnterVariable_decl_assignment is called when entering the variable_decl_assignment production.
	EnterVariable_decl_assignment(c *Variable_decl_assignmentContext)

	// EnterClass_new is called when entering the class_new production.
	EnterClass_new(c *Class_newContext)

	// EnterDynamic_array_new is called when entering the dynamic_array_new production.
	EnterDynamic_array_new(c *Dynamic_array_newContext)

	// EnterUnpacked_dimension is called when entering the unpacked_dimension production.
	EnterUnpacked_dimension(c *Unpacked_dimensionContext)

	// EnterPacked_dimension is called when entering the packed_dimension production.
	EnterPacked_dimension(c *Packed_dimensionContext)

	// EnterAssociative_dimension is called when entering the associative_dimension production.
	EnterAssociative_dimension(c *Associative_dimensionContext)

	// EnterVariable_dimension is called when entering the variable_dimension production.
	EnterVariable_dimension(c *Variable_dimensionContext)

	// EnterQueue_dimension is called when entering the queue_dimension production.
	EnterQueue_dimension(c *Queue_dimensionContext)

	// EnterUnsized_dimension is called when entering the unsized_dimension production.
	EnterUnsized_dimension(c *Unsized_dimensionContext)

	// EnterFunction_data_type_or_implicit is called when entering the function_data_type_or_implicit production.
	EnterFunction_data_type_or_implicit(c *Function_data_type_or_implicitContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_body_declaration is called when entering the function_body_declaration production.
	EnterFunction_body_declaration(c *Function_body_declarationContext)

	// EnterFunction_prototype is called when entering the function_prototype production.
	EnterFunction_prototype(c *Function_prototypeContext)

	// EnterDpi_import_export is called when entering the dpi_import_export production.
	EnterDpi_import_export(c *Dpi_import_exportContext)

	// EnterDpi_spec_string is called when entering the dpi_spec_string production.
	EnterDpi_spec_string(c *Dpi_spec_stringContext)

	// EnterDpi_function_import_property is called when entering the dpi_function_import_property production.
	EnterDpi_function_import_property(c *Dpi_function_import_propertyContext)

	// EnterDpi_task_import_property is called when entering the dpi_task_import_property production.
	EnterDpi_task_import_property(c *Dpi_task_import_propertyContext)

	// EnterDpi_function_proto is called when entering the dpi_function_proto production.
	EnterDpi_function_proto(c *Dpi_function_protoContext)

	// EnterDpi_task_proto is called when entering the dpi_task_proto production.
	EnterDpi_task_proto(c *Dpi_task_protoContext)

	// EnterTask_declaration is called when entering the task_declaration production.
	EnterTask_declaration(c *Task_declarationContext)

	// EnterTask_body_declaration is called when entering the task_body_declaration production.
	EnterTask_body_declaration(c *Task_body_declarationContext)

	// EnterTf_item_declaration is called when entering the tf_item_declaration production.
	EnterTf_item_declaration(c *Tf_item_declarationContext)

	// EnterTf_port_list is called when entering the tf_port_list production.
	EnterTf_port_list(c *Tf_port_listContext)

	// EnterTf_port_item is called when entering the tf_port_item production.
	EnterTf_port_item(c *Tf_port_itemContext)

	// EnterTf_port_direction is called when entering the tf_port_direction production.
	EnterTf_port_direction(c *Tf_port_directionContext)

	// EnterTf_port_declaration is called when entering the tf_port_declaration production.
	EnterTf_port_declaration(c *Tf_port_declarationContext)

	// EnterTask_prototype is called when entering the task_prototype production.
	EnterTask_prototype(c *Task_prototypeContext)

	// EnterBlock_item_declaration is called when entering the block_item_declaration production.
	EnterBlock_item_declaration(c *Block_item_declarationContext)

	// EnterModport_declaration is called when entering the modport_declaration production.
	EnterModport_declaration(c *Modport_declarationContext)

	// EnterModport_item is called when entering the modport_item production.
	EnterModport_item(c *Modport_itemContext)

	// EnterModport_ports_declaration is called when entering the modport_ports_declaration production.
	EnterModport_ports_declaration(c *Modport_ports_declarationContext)

	// EnterModport_clocking_declaration is called when entering the modport_clocking_declaration production.
	EnterModport_clocking_declaration(c *Modport_clocking_declarationContext)

	// EnterModport_simple_ports_declaration is called when entering the modport_simple_ports_declaration production.
	EnterModport_simple_ports_declaration(c *Modport_simple_ports_declarationContext)

	// EnterModport_simple_port is called when entering the modport_simple_port production.
	EnterModport_simple_port(c *Modport_simple_portContext)

	// EnterModport_tf_ports_declaration is called when entering the modport_tf_ports_declaration production.
	EnterModport_tf_ports_declaration(c *Modport_tf_ports_declarationContext)

	// EnterModport_tf_port is called when entering the modport_tf_port production.
	EnterModport_tf_port(c *Modport_tf_portContext)

	// EnterImport_export is called when entering the import_export production.
	EnterImport_export(c *Import_exportContext)

	// EnterConcurrent_assertion_item is called when entering the concurrent_assertion_item production.
	EnterConcurrent_assertion_item(c *Concurrent_assertion_itemContext)

	// EnterConcurrent_assertion_statement is called when entering the concurrent_assertion_statement production.
	EnterConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// EnterAssert_property_statement is called when entering the assert_property_statement production.
	EnterAssert_property_statement(c *Assert_property_statementContext)

	// EnterAssume_property_statement is called when entering the assume_property_statement production.
	EnterAssume_property_statement(c *Assume_property_statementContext)

	// EnterCover_property_statement is called when entering the cover_property_statement production.
	EnterCover_property_statement(c *Cover_property_statementContext)

	// EnterExpect_property_statement is called when entering the expect_property_statement production.
	EnterExpect_property_statement(c *Expect_property_statementContext)

	// EnterCover_sequence_statement is called when entering the cover_sequence_statement production.
	EnterCover_sequence_statement(c *Cover_sequence_statementContext)

	// EnterRestrict_property_statement is called when entering the restrict_property_statement production.
	EnterRestrict_property_statement(c *Restrict_property_statementContext)

	// EnterProperty_instance is called when entering the property_instance production.
	EnterProperty_instance(c *Property_instanceContext)

	// EnterProperty_list_of_arguments is called when entering the property_list_of_arguments production.
	EnterProperty_list_of_arguments(c *Property_list_of_argumentsContext)

	// EnterProperty_actual_arg is called when entering the property_actual_arg production.
	EnterProperty_actual_arg(c *Property_actual_argContext)

	// EnterAssertion_item_declaration is called when entering the assertion_item_declaration production.
	EnterAssertion_item_declaration(c *Assertion_item_declarationContext)

	// EnterProperty_declaration is called when entering the property_declaration production.
	EnterProperty_declaration(c *Property_declarationContext)

	// EnterProperty_port_list is called when entering the property_port_list production.
	EnterProperty_port_list(c *Property_port_listContext)

	// EnterProperty_port_item is called when entering the property_port_item production.
	EnterProperty_port_item(c *Property_port_itemContext)

	// EnterProperty_lvar_port_direction is called when entering the property_lvar_port_direction production.
	EnterProperty_lvar_port_direction(c *Property_lvar_port_directionContext)

	// EnterProperty_formal_type is called when entering the property_formal_type production.
	EnterProperty_formal_type(c *Property_formal_typeContext)

	// EnterProperty_spec is called when entering the property_spec production.
	EnterProperty_spec(c *Property_specContext)

	// EnterProperty_expr is called when entering the property_expr production.
	EnterProperty_expr(c *Property_exprContext)

	// EnterProperty_case_item is called when entering the property_case_item production.
	EnterProperty_case_item(c *Property_case_itemContext)

	// EnterSequence_declaration is called when entering the sequence_declaration production.
	EnterSequence_declaration(c *Sequence_declarationContext)

	// EnterSequence_port_list is called when entering the sequence_port_list production.
	EnterSequence_port_list(c *Sequence_port_listContext)

	// EnterSequence_port_item is called when entering the sequence_port_item production.
	EnterSequence_port_item(c *Sequence_port_itemContext)

	// EnterSequence_lvar_port_direction is called when entering the sequence_lvar_port_direction production.
	EnterSequence_lvar_port_direction(c *Sequence_lvar_port_directionContext)

	// EnterSequence_formal_type is called when entering the sequence_formal_type production.
	EnterSequence_formal_type(c *Sequence_formal_typeContext)

	// EnterSequence_expr is called when entering the sequence_expr production.
	EnterSequence_expr(c *Sequence_exprContext)

	// EnterCycle_delay_range is called when entering the cycle_delay_range production.
	EnterCycle_delay_range(c *Cycle_delay_rangeContext)

	// EnterSequence_method_call is called when entering the sequence_method_call production.
	EnterSequence_method_call(c *Sequence_method_callContext)

	// EnterSequence_match_item is called when entering the sequence_match_item production.
	EnterSequence_match_item(c *Sequence_match_itemContext)

	// EnterSequence_instance is called when entering the sequence_instance production.
	EnterSequence_instance(c *Sequence_instanceContext)

	// EnterSequence_list_of_arguments is called when entering the sequence_list_of_arguments production.
	EnterSequence_list_of_arguments(c *Sequence_list_of_argumentsContext)

	// EnterSequence_actual_arg is called when entering the sequence_actual_arg production.
	EnterSequence_actual_arg(c *Sequence_actual_argContext)

	// EnterBoolean_abbrev is called when entering the boolean_abbrev production.
	EnterBoolean_abbrev(c *Boolean_abbrevContext)

	// EnterSequence_abbrev is called when entering the sequence_abbrev production.
	EnterSequence_abbrev(c *Sequence_abbrevContext)

	// EnterConsecutive_repetition is called when entering the consecutive_repetition production.
	EnterConsecutive_repetition(c *Consecutive_repetitionContext)

	// EnterNon_consecutive_repetition is called when entering the non_consecutive_repetition production.
	EnterNon_consecutive_repetition(c *Non_consecutive_repetitionContext)

	// EnterGoto_repetition is called when entering the goto_repetition production.
	EnterGoto_repetition(c *Goto_repetitionContext)

	// EnterConst_or_range_expression is called when entering the const_or_range_expression production.
	EnterConst_or_range_expression(c *Const_or_range_expressionContext)

	// EnterCycle_delay_const_range_expression is called when entering the cycle_delay_const_range_expression production.
	EnterCycle_delay_const_range_expression(c *Cycle_delay_const_range_expressionContext)

	// EnterExpression_or_dist is called when entering the expression_or_dist production.
	EnterExpression_or_dist(c *Expression_or_distContext)

	// EnterAssertion_variable_declaration is called when entering the assertion_variable_declaration production.
	EnterAssertion_variable_declaration(c *Assertion_variable_declarationContext)

	// EnterCovergroup_declaration is called when entering the covergroup_declaration production.
	EnterCovergroup_declaration(c *Covergroup_declarationContext)

	// EnterCoverage_spec_or_option is called when entering the coverage_spec_or_option production.
	EnterCoverage_spec_or_option(c *Coverage_spec_or_optionContext)

	// EnterCoverage_option is called when entering the coverage_option production.
	EnterCoverage_option(c *Coverage_optionContext)

	// EnterCoverage_spec is called when entering the coverage_spec production.
	EnterCoverage_spec(c *Coverage_specContext)

	// EnterCoverage_event is called when entering the coverage_event production.
	EnterCoverage_event(c *Coverage_eventContext)

	// EnterBlock_event_expression is called when entering the block_event_expression production.
	EnterBlock_event_expression(c *Block_event_expressionContext)

	// EnterHierarchical_btf_identifier is called when entering the hierarchical_btf_identifier production.
	EnterHierarchical_btf_identifier(c *Hierarchical_btf_identifierContext)

	// EnterCover_point is called when entering the cover_point production.
	EnterCover_point(c *Cover_pointContext)

	// EnterBins_or_empty is called when entering the bins_or_empty production.
	EnterBins_or_empty(c *Bins_or_emptyContext)

	// EnterBins_or_options is called when entering the bins_or_options production.
	EnterBins_or_options(c *Bins_or_optionsContext)

	// EnterBins_keyword is called when entering the bins_keyword production.
	EnterBins_keyword(c *Bins_keywordContext)

	// EnterTrans_list is called when entering the trans_list production.
	EnterTrans_list(c *Trans_listContext)

	// EnterTrans_set is called when entering the trans_set production.
	EnterTrans_set(c *Trans_setContext)

	// EnterTrans_range_list is called when entering the trans_range_list production.
	EnterTrans_range_list(c *Trans_range_listContext)

	// EnterTrans_item is called when entering the trans_item production.
	EnterTrans_item(c *Trans_itemContext)

	// EnterRepeat_range is called when entering the repeat_range production.
	EnterRepeat_range(c *Repeat_rangeContext)

	// EnterCover_cross is called when entering the cover_cross production.
	EnterCover_cross(c *Cover_crossContext)

	// EnterList_of_cross_items is called when entering the list_of_cross_items production.
	EnterList_of_cross_items(c *List_of_cross_itemsContext)

	// EnterCross_item is called when entering the cross_item production.
	EnterCross_item(c *Cross_itemContext)

	// EnterCross_body is called when entering the cross_body production.
	EnterCross_body(c *Cross_bodyContext)

	// EnterCross_body_item is called when entering the cross_body_item production.
	EnterCross_body_item(c *Cross_body_itemContext)

	// EnterBins_selection_or_option is called when entering the bins_selection_or_option production.
	EnterBins_selection_or_option(c *Bins_selection_or_optionContext)

	// EnterBins_selection is called when entering the bins_selection production.
	EnterBins_selection(c *Bins_selectionContext)

	// EnterSelect_expression is called when entering the select_expression production.
	EnterSelect_expression(c *Select_expressionContext)

	// EnterSelect_condition is called when entering the select_condition production.
	EnterSelect_condition(c *Select_conditionContext)

	// EnterBins_expression is called when entering the bins_expression production.
	EnterBins_expression(c *Bins_expressionContext)

	// EnterCovergroup_range_list is called when entering the covergroup_range_list production.
	EnterCovergroup_range_list(c *Covergroup_range_listContext)

	// EnterCovergroup_value_range is called when entering the covergroup_value_range production.
	EnterCovergroup_value_range(c *Covergroup_value_rangeContext)

	// EnterWith_covergroup_expression is called when entering the with_covergroup_expression production.
	EnterWith_covergroup_expression(c *With_covergroup_expressionContext)

	// EnterSet_covergroup_expression is called when entering the set_covergroup_expression production.
	EnterSet_covergroup_expression(c *Set_covergroup_expressionContext)

	// EnterInteger_covergroup_expression is called when entering the integer_covergroup_expression production.
	EnterInteger_covergroup_expression(c *Integer_covergroup_expressionContext)

	// EnterCross_set_expression is called when entering the cross_set_expression production.
	EnterCross_set_expression(c *Cross_set_expressionContext)

	// EnterCovergroup_expression is called when entering the covergroup_expression production.
	EnterCovergroup_expression(c *Covergroup_expressionContext)

	// EnterLet_declaration is called when entering the let_declaration production.
	EnterLet_declaration(c *Let_declarationContext)

	// EnterLet_identifier is called when entering the let_identifier production.
	EnterLet_identifier(c *Let_identifierContext)

	// EnterLet_port_list is called when entering the let_port_list production.
	EnterLet_port_list(c *Let_port_listContext)

	// EnterLet_port_item is called when entering the let_port_item production.
	EnterLet_port_item(c *Let_port_itemContext)

	// EnterLet_formal_type is called when entering the let_formal_type production.
	EnterLet_formal_type(c *Let_formal_typeContext)

	// EnterLet_expression is called when entering the let_expression production.
	EnterLet_expression(c *Let_expressionContext)

	// EnterLet_list_of_arguments is called when entering the let_list_of_arguments production.
	EnterLet_list_of_arguments(c *Let_list_of_argumentsContext)

	// EnterLet_actual_arg is called when entering the let_actual_arg production.
	EnterLet_actual_arg(c *Let_actual_argContext)

	// EnterGate_instantiation is called when entering the gate_instantiation production.
	EnterGate_instantiation(c *Gate_instantiationContext)

	// EnterCmos_switch_instance is called when entering the cmos_switch_instance production.
	EnterCmos_switch_instance(c *Cmos_switch_instanceContext)

	// EnterEnable_gate_instance is called when entering the enable_gate_instance production.
	EnterEnable_gate_instance(c *Enable_gate_instanceContext)

	// EnterMos_switch_instance is called when entering the mos_switch_instance production.
	EnterMos_switch_instance(c *Mos_switch_instanceContext)

	// EnterN_input_gate_instance is called when entering the n_input_gate_instance production.
	EnterN_input_gate_instance(c *N_input_gate_instanceContext)

	// EnterN_output_gate_instance is called when entering the n_output_gate_instance production.
	EnterN_output_gate_instance(c *N_output_gate_instanceContext)

	// EnterPass_switch_instance is called when entering the pass_switch_instance production.
	EnterPass_switch_instance(c *Pass_switch_instanceContext)

	// EnterPass_enable_switch_instance is called when entering the pass_enable_switch_instance production.
	EnterPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// EnterPull_gate_instance is called when entering the pull_gate_instance production.
	EnterPull_gate_instance(c *Pull_gate_instanceContext)

	// EnterPulldown_strength is called when entering the pulldown_strength production.
	EnterPulldown_strength(c *Pulldown_strengthContext)

	// EnterPullup_strength is called when entering the pullup_strength production.
	EnterPullup_strength(c *Pullup_strengthContext)

	// EnterEnable_terminal is called when entering the enable_terminal production.
	EnterEnable_terminal(c *Enable_terminalContext)

	// EnterInout_terminal is called when entering the inout_terminal production.
	EnterInout_terminal(c *Inout_terminalContext)

	// EnterInput_terminal is called when entering the input_terminal production.
	EnterInput_terminal(c *Input_terminalContext)

	// EnterNcontrol_terminal is called when entering the ncontrol_terminal production.
	EnterNcontrol_terminal(c *Ncontrol_terminalContext)

	// EnterOutput_terminal is called when entering the output_terminal production.
	EnterOutput_terminal(c *Output_terminalContext)

	// EnterPcontrol_terminal is called when entering the pcontrol_terminal production.
	EnterPcontrol_terminal(c *Pcontrol_terminalContext)

	// EnterCmos_switchtype is called when entering the cmos_switchtype production.
	EnterCmos_switchtype(c *Cmos_switchtypeContext)

	// EnterEnable_gatetype is called when entering the enable_gatetype production.
	EnterEnable_gatetype(c *Enable_gatetypeContext)

	// EnterMos_switchtype is called when entering the mos_switchtype production.
	EnterMos_switchtype(c *Mos_switchtypeContext)

	// EnterN_input_gatetype is called when entering the n_input_gatetype production.
	EnterN_input_gatetype(c *N_input_gatetypeContext)

	// EnterN_output_gatetype is called when entering the n_output_gatetype production.
	EnterN_output_gatetype(c *N_output_gatetypeContext)

	// EnterPass_en_switchtype is called when entering the pass_en_switchtype production.
	EnterPass_en_switchtype(c *Pass_en_switchtypeContext)

	// EnterPass_switchtype is called when entering the pass_switchtype production.
	EnterPass_switchtype(c *Pass_switchtypeContext)

	// EnterModule_instantiation is called when entering the module_instantiation production.
	EnterModule_instantiation(c *Module_instantiationContext)

	// EnterParameter_value_assignment is called when entering the parameter_value_assignment production.
	EnterParameter_value_assignment(c *Parameter_value_assignmentContext)

	// EnterList_of_parameter_assignments is called when entering the list_of_parameter_assignments production.
	EnterList_of_parameter_assignments(c *List_of_parameter_assignmentsContext)

	// EnterOrdered_parameter_assignment is called when entering the ordered_parameter_assignment production.
	EnterOrdered_parameter_assignment(c *Ordered_parameter_assignmentContext)

	// EnterNamed_parameter_assignment is called when entering the named_parameter_assignment production.
	EnterNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// EnterHierarchical_instance is called when entering the hierarchical_instance production.
	EnterHierarchical_instance(c *Hierarchical_instanceContext)

	// EnterName_of_instance is called when entering the name_of_instance production.
	EnterName_of_instance(c *Name_of_instanceContext)

	// EnterList_of_port_connections is called when entering the list_of_port_connections production.
	EnterList_of_port_connections(c *List_of_port_connectionsContext)

	// EnterOrdered_port_connection is called when entering the ordered_port_connection production.
	EnterOrdered_port_connection(c *Ordered_port_connectionContext)

	// EnterNamed_port_connection is called when entering the named_port_connection production.
	EnterNamed_port_connection(c *Named_port_connectionContext)

	// EnterInterface_instantiation is called when entering the interface_instantiation production.
	EnterInterface_instantiation(c *Interface_instantiationContext)

	// EnterProgram_instantiation is called when entering the program_instantiation production.
	EnterProgram_instantiation(c *Program_instantiationContext)

	// EnterChecker_instantiation is called when entering the checker_instantiation production.
	EnterChecker_instantiation(c *Checker_instantiationContext)

	// EnterList_of_checker_port_connections is called when entering the list_of_checker_port_connections production.
	EnterList_of_checker_port_connections(c *List_of_checker_port_connectionsContext)

	// EnterOrdered_checker_port_connection is called when entering the ordered_checker_port_connection production.
	EnterOrdered_checker_port_connection(c *Ordered_checker_port_connectionContext)

	// EnterNamed_checker_port_connection is called when entering the named_checker_port_connection production.
	EnterNamed_checker_port_connection(c *Named_checker_port_connectionContext)

	// EnterGenerate_region is called when entering the generate_region production.
	EnterGenerate_region(c *Generate_regionContext)

	// EnterLoop_generate_construct is called when entering the loop_generate_construct production.
	EnterLoop_generate_construct(c *Loop_generate_constructContext)

	// EnterGenvar_initialization is called when entering the genvar_initialization production.
	EnterGenvar_initialization(c *Genvar_initializationContext)

	// EnterGenvar_iteration is called when entering the genvar_iteration production.
	EnterGenvar_iteration(c *Genvar_iterationContext)

	// EnterConditional_generate_construct is called when entering the conditional_generate_construct production.
	EnterConditional_generate_construct(c *Conditional_generate_constructContext)

	// EnterIf_generate_construct is called when entering the if_generate_construct production.
	EnterIf_generate_construct(c *If_generate_constructContext)

	// EnterCase_generate_construct is called when entering the case_generate_construct production.
	EnterCase_generate_construct(c *Case_generate_constructContext)

	// EnterCase_generate_item is called when entering the case_generate_item production.
	EnterCase_generate_item(c *Case_generate_itemContext)

	// EnterGenerate_block is called when entering the generate_block production.
	EnterGenerate_block(c *Generate_blockContext)

	// EnterGenerate_item is called when entering the generate_item production.
	EnterGenerate_item(c *Generate_itemContext)

	// EnterContinuous_assign is called when entering the continuous_assign production.
	EnterContinuous_assign(c *Continuous_assignContext)

	// EnterList_of_net_assignments is called when entering the list_of_net_assignments production.
	EnterList_of_net_assignments(c *List_of_net_assignmentsContext)

	// EnterList_of_variable_assignments is called when entering the list_of_variable_assignments production.
	EnterList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// EnterNet_alias is called when entering the net_alias production.
	EnterNet_alias(c *Net_aliasContext)

	// EnterNet_assignment is called when entering the net_assignment production.
	EnterNet_assignment(c *Net_assignmentContext)

	// EnterInitial_construct is called when entering the initial_construct production.
	EnterInitial_construct(c *Initial_constructContext)

	// EnterAlways_construct is called when entering the always_construct production.
	EnterAlways_construct(c *Always_constructContext)

	// EnterAlways_keyword is called when entering the always_keyword production.
	EnterAlways_keyword(c *Always_keywordContext)

	// EnterFinal_construct is called when entering the final_construct production.
	EnterFinal_construct(c *Final_constructContext)

	// EnterBlocking_assignment is called when entering the blocking_assignment production.
	EnterBlocking_assignment(c *Blocking_assignmentContext)

	// EnterOperator_assignment is called when entering the operator_assignment production.
	EnterOperator_assignment(c *Operator_assignmentContext)

	// EnterAssignment_operator is called when entering the assignment_operator production.
	EnterAssignment_operator(c *Assignment_operatorContext)

	// EnterNonblocking_assignment is called when entering the nonblocking_assignment production.
	EnterNonblocking_assignment(c *Nonblocking_assignmentContext)

	// EnterProcedural_continuous_assignment is called when entering the procedural_continuous_assignment production.
	EnterProcedural_continuous_assignment(c *Procedural_continuous_assignmentContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterAction_block is called when entering the action_block production.
	EnterAction_block(c *Action_blockContext)

	// EnterSeq_block is called when entering the seq_block production.
	EnterSeq_block(c *Seq_blockContext)

	// EnterPar_block is called when entering the par_block production.
	EnterPar_block(c *Par_blockContext)

	// EnterJoin_keyword is called when entering the join_keyword production.
	EnterJoin_keyword(c *Join_keywordContext)

	// EnterStatement_or_null is called when entering the statement_or_null production.
	EnterStatement_or_null(c *Statement_or_nullContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatement_item is called when entering the statement_item production.
	EnterStatement_item(c *Statement_itemContext)

	// EnterFunction_statement is called when entering the function_statement production.
	EnterFunction_statement(c *Function_statementContext)

	// EnterFunction_statement_or_null is called when entering the function_statement_or_null production.
	EnterFunction_statement_or_null(c *Function_statement_or_nullContext)

	// EnterVariable_identifier_list is called when entering the variable_identifier_list production.
	EnterVariable_identifier_list(c *Variable_identifier_listContext)

	// EnterProcedural_timing_control_statement is called when entering the procedural_timing_control_statement production.
	EnterProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// EnterDelay_or_event_control is called when entering the delay_or_event_control production.
	EnterDelay_or_event_control(c *Delay_or_event_controlContext)

	// EnterDelay_control is called when entering the delay_control production.
	EnterDelay_control(c *Delay_controlContext)

	// EnterEvent_control is called when entering the event_control production.
	EnterEvent_control(c *Event_controlContext)

	// EnterEvent_expression is called when entering the event_expression production.
	EnterEvent_expression(c *Event_expressionContext)

	// EnterProcedural_timing_control is called when entering the procedural_timing_control production.
	EnterProcedural_timing_control(c *Procedural_timing_controlContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterEvent_trigger is called when entering the event_trigger production.
	EnterEvent_trigger(c *Event_triggerContext)

	// EnterDisable_statement is called when entering the disable_statement production.
	EnterDisable_statement(c *Disable_statementContext)

	// EnterConditional_statement is called when entering the conditional_statement production.
	EnterConditional_statement(c *Conditional_statementContext)

	// EnterUnique_priority is called when entering the unique_priority production.
	EnterUnique_priority(c *Unique_priorityContext)

	// EnterCond_predicate is called when entering the cond_predicate production.
	EnterCond_predicate(c *Cond_predicateContext)

	// EnterExpression_or_cond_pattern is called when entering the expression_or_cond_pattern production.
	EnterExpression_or_cond_pattern(c *Expression_or_cond_patternContext)

	// EnterCond_pattern is called when entering the cond_pattern production.
	EnterCond_pattern(c *Cond_patternContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_keyword is called when entering the case_keyword production.
	EnterCase_keyword(c *Case_keywordContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterCase_item is called when entering the case_item production.
	EnterCase_item(c *Case_itemContext)

	// EnterCase_pattern_item is called when entering the case_pattern_item production.
	EnterCase_pattern_item(c *Case_pattern_itemContext)

	// EnterCase_inside_item is called when entering the case_inside_item production.
	EnterCase_inside_item(c *Case_inside_itemContext)

	// EnterCase_item_expression is called when entering the case_item_expression production.
	EnterCase_item_expression(c *Case_item_expressionContext)

	// EnterRandcase_statement is called when entering the randcase_statement production.
	EnterRandcase_statement(c *Randcase_statementContext)

	// EnterRandcase_item is called when entering the randcase_item production.
	EnterRandcase_item(c *Randcase_itemContext)

	// EnterOpen_range_list is called when entering the open_range_list production.
	EnterOpen_range_list(c *Open_range_listContext)

	// EnterOpen_value_range is called when entering the open_value_range production.
	EnterOpen_value_range(c *Open_value_rangeContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterAssignment_pattern is called when entering the assignment_pattern production.
	EnterAssignment_pattern(c *Assignment_patternContext)

	// EnterStructure_pattern_key is called when entering the structure_pattern_key production.
	EnterStructure_pattern_key(c *Structure_pattern_keyContext)

	// EnterArray_pattern_key is called when entering the array_pattern_key production.
	EnterArray_pattern_key(c *Array_pattern_keyContext)

	// EnterAssignment_pattern_key is called when entering the assignment_pattern_key production.
	EnterAssignment_pattern_key(c *Assignment_pattern_keyContext)

	// EnterAssignment_pattern_expression is called when entering the assignment_pattern_expression production.
	EnterAssignment_pattern_expression(c *Assignment_pattern_expressionContext)

	// EnterAssignment_pattern_expression_type is called when entering the assignment_pattern_expression_type production.
	EnterAssignment_pattern_expression_type(c *Assignment_pattern_expression_typeContext)

	// EnterConstant_assignment_pattern_expression is called when entering the constant_assignment_pattern_expression production.
	EnterConstant_assignment_pattern_expression(c *Constant_assignment_pattern_expressionContext)

	// EnterAssignment_pattern_net_lvalue is called when entering the assignment_pattern_net_lvalue production.
	EnterAssignment_pattern_net_lvalue(c *Assignment_pattern_net_lvalueContext)

	// EnterAssignment_pattern_variable_lvalue is called when entering the assignment_pattern_variable_lvalue production.
	EnterAssignment_pattern_variable_lvalue(c *Assignment_pattern_variable_lvalueContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterFor_initialization is called when entering the for_initialization production.
	EnterFor_initialization(c *For_initializationContext)

	// EnterFor_variable_declaration is called when entering the for_variable_declaration production.
	EnterFor_variable_declaration(c *For_variable_declarationContext)

	// EnterFor_step is called when entering the for_step production.
	EnterFor_step(c *For_stepContext)

	// EnterFor_step_assignment is called when entering the for_step_assignment production.
	EnterFor_step_assignment(c *For_step_assignmentContext)

	// EnterLoop_variables is called when entering the loop_variables production.
	EnterLoop_variables(c *Loop_variablesContext)

	// EnterSubroutine_call_statement is called when entering the subroutine_call_statement production.
	EnterSubroutine_call_statement(c *Subroutine_call_statementContext)

	// EnterAssertion_item is called when entering the assertion_item production.
	EnterAssertion_item(c *Assertion_itemContext)

	// EnterDeferred_immediate_assertion_item is called when entering the deferred_immediate_assertion_item production.
	EnterDeferred_immediate_assertion_item(c *Deferred_immediate_assertion_itemContext)

	// EnterProcedural_assertion_statement is called when entering the procedural_assertion_statement production.
	EnterProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// EnterImmediate_assertion_statement is called when entering the immediate_assertion_statement production.
	EnterImmediate_assertion_statement(c *Immediate_assertion_statementContext)

	// EnterSimple_immediate_assertion_statement is called when entering the simple_immediate_assertion_statement production.
	EnterSimple_immediate_assertion_statement(c *Simple_immediate_assertion_statementContext)

	// EnterSimple_immediate_assert_statement is called when entering the simple_immediate_assert_statement production.
	EnterSimple_immediate_assert_statement(c *Simple_immediate_assert_statementContext)

	// EnterSimple_immediate_assume_statement is called when entering the simple_immediate_assume_statement production.
	EnterSimple_immediate_assume_statement(c *Simple_immediate_assume_statementContext)

	// EnterSimple_immediate_cover_statement is called when entering the simple_immediate_cover_statement production.
	EnterSimple_immediate_cover_statement(c *Simple_immediate_cover_statementContext)

	// EnterDeferred_immediate_assertion_statement is called when entering the deferred_immediate_assertion_statement production.
	EnterDeferred_immediate_assertion_statement(c *Deferred_immediate_assertion_statementContext)

	// EnterDeferred_immediate_assert_statement is called when entering the deferred_immediate_assert_statement production.
	EnterDeferred_immediate_assert_statement(c *Deferred_immediate_assert_statementContext)

	// EnterDeferred_immediate_assume_statement is called when entering the deferred_immediate_assume_statement production.
	EnterDeferred_immediate_assume_statement(c *Deferred_immediate_assume_statementContext)

	// EnterDeferred_immediate_cover_statement is called when entering the deferred_immediate_cover_statement production.
	EnterDeferred_immediate_cover_statement(c *Deferred_immediate_cover_statementContext)

	// EnterClocking_declaration is called when entering the clocking_declaration production.
	EnterClocking_declaration(c *Clocking_declarationContext)

	// EnterClocking_event is called when entering the clocking_event production.
	EnterClocking_event(c *Clocking_eventContext)

	// EnterClocking_item is called when entering the clocking_item production.
	EnterClocking_item(c *Clocking_itemContext)

	// EnterDefault_skew is called when entering the default_skew production.
	EnterDefault_skew(c *Default_skewContext)

	// EnterClocking_direction is called when entering the clocking_direction production.
	EnterClocking_direction(c *Clocking_directionContext)

	// EnterList_of_clocking_decl_assign is called when entering the list_of_clocking_decl_assign production.
	EnterList_of_clocking_decl_assign(c *List_of_clocking_decl_assignContext)

	// EnterClocking_decl_assign is called when entering the clocking_decl_assign production.
	EnterClocking_decl_assign(c *Clocking_decl_assignContext)

	// EnterClocking_skew is called when entering the clocking_skew production.
	EnterClocking_skew(c *Clocking_skewContext)

	// EnterClocking_drive is called when entering the clocking_drive production.
	EnterClocking_drive(c *Clocking_driveContext)

	// EnterCycle_delay is called when entering the cycle_delay production.
	EnterCycle_delay(c *Cycle_delayContext)

	// EnterClockvar is called when entering the clockvar production.
	EnterClockvar(c *ClockvarContext)

	// EnterClockvar_expression is called when entering the clockvar_expression production.
	EnterClockvar_expression(c *Clockvar_expressionContext)

	// EnterRandsequence_statement is called when entering the randsequence_statement production.
	EnterRandsequence_statement(c *Randsequence_statementContext)

	// EnterProduction is called when entering the production production.
	EnterProduction(c *ProductionContext)

	// EnterRs_rule is called when entering the rs_rule production.
	EnterRs_rule(c *Rs_ruleContext)

	// EnterRs_production_list is called when entering the rs_production_list production.
	EnterRs_production_list(c *Rs_production_listContext)

	// EnterWeight_specification is called when entering the weight_specification production.
	EnterWeight_specification(c *Weight_specificationContext)

	// EnterRs_code_block is called when entering the rs_code_block production.
	EnterRs_code_block(c *Rs_code_blockContext)

	// EnterRs_prod is called when entering the rs_prod production.
	EnterRs_prod(c *Rs_prodContext)

	// EnterProduction_item is called when entering the production_item production.
	EnterProduction_item(c *Production_itemContext)

	// EnterRs_if_else is called when entering the rs_if_else production.
	EnterRs_if_else(c *Rs_if_elseContext)

	// EnterRs_repeat is called when entering the rs_repeat production.
	EnterRs_repeat(c *Rs_repeatContext)

	// EnterRs_case is called when entering the rs_case production.
	EnterRs_case(c *Rs_caseContext)

	// EnterRs_case_item is called when entering the rs_case_item production.
	EnterRs_case_item(c *Rs_case_itemContext)

	// EnterSpecify_block is called when entering the specify_block production.
	EnterSpecify_block(c *Specify_blockContext)

	// EnterSpecify_item is called when entering the specify_item production.
	EnterSpecify_item(c *Specify_itemContext)

	// EnterPulsestyle_declaration is called when entering the pulsestyle_declaration production.
	EnterPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// EnterShowcancelled_declaration is called when entering the showcancelled_declaration production.
	EnterShowcancelled_declaration(c *Showcancelled_declarationContext)

	// EnterPath_declaration is called when entering the path_declaration production.
	EnterPath_declaration(c *Path_declarationContext)

	// EnterSimple_path_declaration is called when entering the simple_path_declaration production.
	EnterSimple_path_declaration(c *Simple_path_declarationContext)

	// EnterParallel_path_description is called when entering the parallel_path_description production.
	EnterParallel_path_description(c *Parallel_path_descriptionContext)

	// EnterFull_path_description is called when entering the full_path_description production.
	EnterFull_path_description(c *Full_path_descriptionContext)

	// EnterList_of_path_inputs is called when entering the list_of_path_inputs production.
	EnterList_of_path_inputs(c *List_of_path_inputsContext)

	// EnterList_of_path_outputs is called when entering the list_of_path_outputs production.
	EnterList_of_path_outputs(c *List_of_path_outputsContext)

	// EnterSpecify_input_terminal_descriptor is called when entering the specify_input_terminal_descriptor production.
	EnterSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// EnterSpecify_output_terminal_descriptor is called when entering the specify_output_terminal_descriptor production.
	EnterSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// EnterInput_identifier is called when entering the input_identifier production.
	EnterInput_identifier(c *Input_identifierContext)

	// EnterOutput_identifier is called when entering the output_identifier production.
	EnterOutput_identifier(c *Output_identifierContext)

	// EnterPath_delay_value is called when entering the path_delay_value production.
	EnterPath_delay_value(c *Path_delay_valueContext)

	// EnterList_of_path_delay_expressions is called when entering the list_of_path_delay_expressions production.
	EnterList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// EnterT_path_delay_expression is called when entering the t_path_delay_expression production.
	EnterT_path_delay_expression(c *T_path_delay_expressionContext)

	// EnterTrise_path_delay_expression is called when entering the trise_path_delay_expression production.
	EnterTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// EnterTfall_path_delay_expression is called when entering the tfall_path_delay_expression production.
	EnterTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// EnterTz_path_delay_expression is called when entering the tz_path_delay_expression production.
	EnterTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// EnterT01_path_delay_expression is called when entering the t01_path_delay_expression production.
	EnterT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// EnterT10_path_delay_expression is called when entering the t10_path_delay_expression production.
	EnterT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// EnterT0z_path_delay_expression is called when entering the t0z_path_delay_expression production.
	EnterT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// EnterTz1_path_delay_expression is called when entering the tz1_path_delay_expression production.
	EnterTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// EnterT1z_path_delay_expression is called when entering the t1z_path_delay_expression production.
	EnterT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// EnterTz0_path_delay_expression is called when entering the tz0_path_delay_expression production.
	EnterTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// EnterT0x_path_delay_expression is called when entering the t0x_path_delay_expression production.
	EnterT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// EnterTx1_path_delay_expression is called when entering the tx1_path_delay_expression production.
	EnterTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// EnterT1x_path_delay_expression is called when entering the t1x_path_delay_expression production.
	EnterT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// EnterTx0_path_delay_expression is called when entering the tx0_path_delay_expression production.
	EnterTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// EnterTxz_path_delay_expression is called when entering the txz_path_delay_expression production.
	EnterTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// EnterTzx_path_delay_expression is called when entering the tzx_path_delay_expression production.
	EnterTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// EnterPath_delay_expression is called when entering the path_delay_expression production.
	EnterPath_delay_expression(c *Path_delay_expressionContext)

	// EnterEdge_sensitive_path_declaration is called when entering the edge_sensitive_path_declaration production.
	EnterEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// EnterParallel_edge_sensitive_path_description is called when entering the parallel_edge_sensitive_path_description production.
	EnterParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// EnterFull_edge_sensitive_path_description is called when entering the full_edge_sensitive_path_description production.
	EnterFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// EnterData_source_expression is called when entering the data_source_expression production.
	EnterData_source_expression(c *Data_source_expressionContext)

	// EnterEdge_identifier is called when entering the edge_identifier production.
	EnterEdge_identifier(c *Edge_identifierContext)

	// EnterState_dependent_path_declaration is called when entering the state_dependent_path_declaration production.
	EnterState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// EnterPolarity_operator is called when entering the polarity_operator production.
	EnterPolarity_operator(c *Polarity_operatorContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterConstant_concatenation is called when entering the constant_concatenation production.
	EnterConstant_concatenation(c *Constant_concatenationContext)

	// EnterConstant_multiple_concatenation is called when entering the constant_multiple_concatenation production.
	EnterConstant_multiple_concatenation(c *Constant_multiple_concatenationContext)

	// EnterModule_path_concatenation is called when entering the module_path_concatenation production.
	EnterModule_path_concatenation(c *Module_path_concatenationContext)

	// EnterModule_path_multiple_concatenation is called when entering the module_path_multiple_concatenation production.
	EnterModule_path_multiple_concatenation(c *Module_path_multiple_concatenationContext)

	// EnterMultiple_concatenation is called when entering the multiple_concatenation production.
	EnterMultiple_concatenation(c *Multiple_concatenationContext)

	// EnterStreaming_concatenation is called when entering the streaming_concatenation production.
	EnterStreaming_concatenation(c *Streaming_concatenationContext)

	// EnterStream_operator is called when entering the stream_operator production.
	EnterStream_operator(c *Stream_operatorContext)

	// EnterSlice_size is called when entering the slice_size production.
	EnterSlice_size(c *Slice_sizeContext)

	// EnterStream_concatenation is called when entering the stream_concatenation production.
	EnterStream_concatenation(c *Stream_concatenationContext)

	// EnterStream_expression is called when entering the stream_expression production.
	EnterStream_expression(c *Stream_expressionContext)

	// EnterArray_range_expression is called when entering the array_range_expression production.
	EnterArray_range_expression(c *Array_range_expressionContext)

	// EnterEmpty_unpacked_array_concatenation is called when entering the empty_unpacked_array_concatenation production.
	EnterEmpty_unpacked_array_concatenation(c *Empty_unpacked_array_concatenationContext)

	// EnterConstant_function_call is called when entering the constant_function_call production.
	EnterConstant_function_call(c *Constant_function_callContext)

	// EnterTf_call is called when entering the tf_call production.
	EnterTf_call(c *Tf_callContext)

	// EnterSystem_tf_call is called when entering the system_tf_call production.
	EnterSystem_tf_call(c *System_tf_callContext)

	// EnterSubroutine_call is called when entering the subroutine_call production.
	EnterSubroutine_call(c *Subroutine_callContext)

	// EnterFunction_subroutine_call is called when entering the function_subroutine_call production.
	EnterFunction_subroutine_call(c *Function_subroutine_callContext)

	// EnterList_of_arguments is called when entering the list_of_arguments production.
	EnterList_of_arguments(c *List_of_argumentsContext)

	// EnterMethod_call_body is called when entering the method_call_body production.
	EnterMethod_call_body(c *Method_call_bodyContext)

	// EnterBuilt_in_method_call is called when entering the built_in_method_call production.
	EnterBuilt_in_method_call(c *Built_in_method_callContext)

	// EnterArray_manipulation_call is called when entering the array_manipulation_call production.
	EnterArray_manipulation_call(c *Array_manipulation_callContext)

	// EnterRandomize_call is called when entering the randomize_call production.
	EnterRandomize_call(c *Randomize_callContext)

	// EnterArray_method_name is called when entering the array_method_name production.
	EnterArray_method_name(c *Array_method_nameContext)

	// EnterInc_or_dec_expression is called when entering the inc_or_dec_expression production.
	EnterInc_or_dec_expression(c *Inc_or_dec_expressionContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterConstant_mintypmax_expression is called when entering the constant_mintypmax_expression production.
	EnterConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// EnterConstant_param_expression is called when entering the constant_param_expression production.
	EnterConstant_param_expression(c *Constant_param_expressionContext)

	// EnterParam_expression is called when entering the param_expression production.
	EnterParam_expression(c *Param_expressionContext)

	// EnterConstant_range_expression is called when entering the constant_range_expression production.
	EnterConstant_range_expression(c *Constant_range_expressionContext)

	// EnterConstant_part_select_range is called when entering the constant_part_select_range production.
	EnterConstant_part_select_range(c *Constant_part_select_rangeContext)

	// EnterConstant_range is called when entering the constant_range production.
	EnterConstant_range(c *Constant_rangeContext)

	// EnterConstant_indexed_range is called when entering the constant_indexed_range production.
	EnterConstant_indexed_range(c *Constant_indexed_rangeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterValue_range is called when entering the value_range production.
	EnterValue_range(c *Value_rangeContext)

	// EnterMintypmax_expression is called when entering the mintypmax_expression production.
	EnterMintypmax_expression(c *Mintypmax_expressionContext)

	// EnterModule_path_expression is called when entering the module_path_expression production.
	EnterModule_path_expression(c *Module_path_expressionContext)

	// EnterModule_path_mintypmax_expression is called when entering the module_path_mintypmax_expression production.
	EnterModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// EnterPart_select_range is called when entering the part_select_range production.
	EnterPart_select_range(c *Part_select_rangeContext)

	// EnterIndexed_range is called when entering the indexed_range production.
	EnterIndexed_range(c *Indexed_rangeContext)

	// EnterGenvar_expression is called when entering the genvar_expression production.
	EnterGenvar_expression(c *Genvar_expressionContext)

	// EnterConstant_primary is called when entering the constant_primary production.
	EnterConstant_primary(c *Constant_primaryContext)

	// EnterModule_path_primary is called when entering the module_path_primary production.
	EnterModule_path_primary(c *Module_path_primaryContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterClass_qualifier is called when entering the class_qualifier production.
	EnterClass_qualifier(c *Class_qualifierContext)

	// EnterRange_expression is called when entering the range_expression production.
	EnterRange_expression(c *Range_expressionContext)

	// EnterPrimary_literal is called when entering the primary_literal production.
	EnterPrimary_literal(c *Primary_literalContext)

	// EnterImplicit_class_handle is called when entering the implicit_class_handle production.
	EnterImplicit_class_handle(c *Implicit_class_handleContext)

	// EnterBit_select is called when entering the bit_select production.
	EnterBit_select(c *Bit_selectContext)

	// EnterSelect_ is called when entering the select_ production.
	EnterSelect_(c *Select_Context)

	// EnterNonrange_select is called when entering the nonrange_select production.
	EnterNonrange_select(c *Nonrange_selectContext)

	// EnterConstant_bit_select is called when entering the constant_bit_select production.
	EnterConstant_bit_select(c *Constant_bit_selectContext)

	// EnterConstant_select is called when entering the constant_select production.
	EnterConstant_select(c *Constant_selectContext)

	// EnterConstant_let_expression is called when entering the constant_let_expression production.
	EnterConstant_let_expression(c *Constant_let_expressionContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterNet_lvalue is called when entering the net_lvalue production.
	EnterNet_lvalue(c *Net_lvalueContext)

	// EnterVariable_lvalue is called when entering the variable_lvalue production.
	EnterVariable_lvalue(c *Variable_lvalueContext)

	// EnterNonrange_variable_lvalue is called when entering the nonrange_variable_lvalue production.
	EnterNonrange_variable_lvalue(c *Nonrange_variable_lvalueContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterInc_or_dec_operator is called when entering the inc_or_dec_operator production.
	EnterInc_or_dec_operator(c *Inc_or_dec_operatorContext)

	// EnterUnary_module_path_operator is called when entering the unary_module_path_operator production.
	EnterUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// EnterBinary_module_path_operator is called when entering the binary_module_path_operator production.
	EnterBinary_module_path_operator(c *Binary_module_path_operatorContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterAttribute_instance is called when entering the attribute_instance production.
	EnterAttribute_instance(c *Attribute_instanceContext)

	// EnterAttr_spec is called when entering the attr_spec production.
	EnterAttr_spec(c *Attr_specContext)

	// EnterAttr_name is called when entering the attr_name production.
	EnterAttr_name(c *Attr_nameContext)

	// EnterArray_identifier is called when entering the array_identifier production.
	EnterArray_identifier(c *Array_identifierContext)

	// EnterBlock_identifier is called when entering the block_identifier production.
	EnterBlock_identifier(c *Block_identifierContext)

	// EnterBin_identifier is called when entering the bin_identifier production.
	EnterBin_identifier(c *Bin_identifierContext)

	// EnterC_identifier is called when entering the c_identifier production.
	EnterC_identifier(c *C_identifierContext)

	// EnterCell_identifier is called when entering the cell_identifier production.
	EnterCell_identifier(c *Cell_identifierContext)

	// EnterChecker_identifier is called when entering the checker_identifier production.
	EnterChecker_identifier(c *Checker_identifierContext)

	// EnterClass_identifier is called when entering the class_identifier production.
	EnterClass_identifier(c *Class_identifierContext)

	// EnterClass_variable_identifier is called when entering the class_variable_identifier production.
	EnterClass_variable_identifier(c *Class_variable_identifierContext)

	// EnterClocking_identifier is called when entering the clocking_identifier production.
	EnterClocking_identifier(c *Clocking_identifierContext)

	// EnterConfig_identifier is called when entering the config_identifier production.
	EnterConfig_identifier(c *Config_identifierContext)

	// EnterConst_identifier is called when entering the const_identifier production.
	EnterConst_identifier(c *Const_identifierContext)

	// EnterConstraint_identifier is called when entering the constraint_identifier production.
	EnterConstraint_identifier(c *Constraint_identifierContext)

	// EnterCovergroup_identifier is called when entering the covergroup_identifier production.
	EnterCovergroup_identifier(c *Covergroup_identifierContext)

	// EnterCovergroup_variable_identifier is called when entering the covergroup_variable_identifier production.
	EnterCovergroup_variable_identifier(c *Covergroup_variable_identifierContext)

	// EnterCover_point_identifier is called when entering the cover_point_identifier production.
	EnterCover_point_identifier(c *Cover_point_identifierContext)

	// EnterCross_identifier is called when entering the cross_identifier production.
	EnterCross_identifier(c *Cross_identifierContext)

	// EnterDynamic_array_variable_identifier is called when entering the dynamic_array_variable_identifier production.
	EnterDynamic_array_variable_identifier(c *Dynamic_array_variable_identifierContext)

	// EnterEnum_identifier is called when entering the enum_identifier production.
	EnterEnum_identifier(c *Enum_identifierContext)

	// EnterFormal_port_identifier is called when entering the formal_port_identifier production.
	EnterFormal_port_identifier(c *Formal_port_identifierContext)

	// EnterFunction_identifier is called when entering the function_identifier production.
	EnterFunction_identifier(c *Function_identifierContext)

	// EnterGenerate_block_identifier is called when entering the generate_block_identifier production.
	EnterGenerate_block_identifier(c *Generate_block_identifierContext)

	// EnterGenvar_identifier is called when entering the genvar_identifier production.
	EnterGenvar_identifier(c *Genvar_identifierContext)

	// EnterHierarchical_array_identifier is called when entering the hierarchical_array_identifier production.
	EnterHierarchical_array_identifier(c *Hierarchical_array_identifierContext)

	// EnterHierarchical_block_identifier is called when entering the hierarchical_block_identifier production.
	EnterHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// EnterHierarchical_event_identifier is called when entering the hierarchical_event_identifier production.
	EnterHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// EnterHierarchical_identifier is called when entering the hierarchical_identifier production.
	EnterHierarchical_identifier(c *Hierarchical_identifierContext)

	// EnterHierarchical_net_identifier is called when entering the hierarchical_net_identifier production.
	EnterHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// EnterHierarchical_parameter_identifier is called when entering the hierarchical_parameter_identifier production.
	EnterHierarchical_parameter_identifier(c *Hierarchical_parameter_identifierContext)

	// EnterHierarchical_property_identifier is called when entering the hierarchical_property_identifier production.
	EnterHierarchical_property_identifier(c *Hierarchical_property_identifierContext)

	// EnterHierarchical_sequence_identifier is called when entering the hierarchical_sequence_identifier production.
	EnterHierarchical_sequence_identifier(c *Hierarchical_sequence_identifierContext)

	// EnterHierarchical_task_identifier is called when entering the hierarchical_task_identifier production.
	EnterHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// EnterHierarchical_tf_identifier is called when entering the hierarchical_tf_identifier production.
	EnterHierarchical_tf_identifier(c *Hierarchical_tf_identifierContext)

	// EnterHierarchical_variable_identifier is called when entering the hierarchical_variable_identifier production.
	EnterHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIndex_variable_identifier is called when entering the index_variable_identifier production.
	EnterIndex_variable_identifier(c *Index_variable_identifierContext)

	// EnterInterface_identifier is called when entering the interface_identifier production.
	EnterInterface_identifier(c *Interface_identifierContext)

	// EnterInterface_instance_identifier is called when entering the interface_instance_identifier production.
	EnterInterface_instance_identifier(c *Interface_instance_identifierContext)

	// EnterInout_port_identifier is called when entering the inout_port_identifier production.
	EnterInout_port_identifier(c *Inout_port_identifierContext)

	// EnterInput_port_identifier is called when entering the input_port_identifier production.
	EnterInput_port_identifier(c *Input_port_identifierContext)

	// EnterInstance_identifier is called when entering the instance_identifier production.
	EnterInstance_identifier(c *Instance_identifierContext)

	// EnterLibrary_identifier is called when entering the library_identifier production.
	EnterLibrary_identifier(c *Library_identifierContext)

	// EnterMember_identifier is called when entering the member_identifier production.
	EnterMember_identifier(c *Member_identifierContext)

	// EnterMethod_identifier is called when entering the method_identifier production.
	EnterMethod_identifier(c *Method_identifierContext)

	// EnterModport_identifier is called when entering the modport_identifier production.
	EnterModport_identifier(c *Modport_identifierContext)

	// EnterModule_identifier is called when entering the module_identifier production.
	EnterModule_identifier(c *Module_identifierContext)

	// EnterNet_identifier is called when entering the net_identifier production.
	EnterNet_identifier(c *Net_identifierContext)

	// EnterNet_type_identifier is called when entering the net_type_identifier production.
	EnterNet_type_identifier(c *Net_type_identifierContext)

	// EnterOutput_port_identifier is called when entering the output_port_identifier production.
	EnterOutput_port_identifier(c *Output_port_identifierContext)

	// EnterPackage_identifier is called when entering the package_identifier production.
	EnterPackage_identifier(c *Package_identifierContext)

	// EnterPackage_scope is called when entering the package_scope production.
	EnterPackage_scope(c *Package_scopeContext)

	// EnterParameter_identifier is called when entering the parameter_identifier production.
	EnterParameter_identifier(c *Parameter_identifierContext)

	// EnterPort_identifier is called when entering the port_identifier production.
	EnterPort_identifier(c *Port_identifierContext)

	// EnterProduction_identifier is called when entering the production_identifier production.
	EnterProduction_identifier(c *Production_identifierContext)

	// EnterProgram_identifier is called when entering the program_identifier production.
	EnterProgram_identifier(c *Program_identifierContext)

	// EnterProperty_identifier is called when entering the property_identifier production.
	EnterProperty_identifier(c *Property_identifierContext)

	// EnterPs_class_identifier is called when entering the ps_class_identifier production.
	EnterPs_class_identifier(c *Ps_class_identifierContext)

	// EnterPs_covergroup_identifier is called when entering the ps_covergroup_identifier production.
	EnterPs_covergroup_identifier(c *Ps_covergroup_identifierContext)

	// EnterPs_checker_identifier is called when entering the ps_checker_identifier production.
	EnterPs_checker_identifier(c *Ps_checker_identifierContext)

	// EnterPs_identifier is called when entering the ps_identifier production.
	EnterPs_identifier(c *Ps_identifierContext)

	// EnterPs_or_hierarchical_array_identifier is called when entering the ps_or_hierarchical_array_identifier production.
	EnterPs_or_hierarchical_array_identifier(c *Ps_or_hierarchical_array_identifierContext)

	// EnterPs_or_hierarchical_net_identifier is called when entering the ps_or_hierarchical_net_identifier production.
	EnterPs_or_hierarchical_net_identifier(c *Ps_or_hierarchical_net_identifierContext)

	// EnterPs_or_hierarchical_property_identifier is called when entering the ps_or_hierarchical_property_identifier production.
	EnterPs_or_hierarchical_property_identifier(c *Ps_or_hierarchical_property_identifierContext)

	// EnterPs_or_hierarchical_sequence_identifier is called when entering the ps_or_hierarchical_sequence_identifier production.
	EnterPs_or_hierarchical_sequence_identifier(c *Ps_or_hierarchical_sequence_identifierContext)

	// EnterPs_or_hierarchical_tf_identifier is called when entering the ps_or_hierarchical_tf_identifier production.
	EnterPs_or_hierarchical_tf_identifier(c *Ps_or_hierarchical_tf_identifierContext)

	// EnterPs_parameter_identifier is called when entering the ps_parameter_identifier production.
	EnterPs_parameter_identifier(c *Ps_parameter_identifierContext)

	// EnterPs_type_identifier is called when entering the ps_type_identifier production.
	EnterPs_type_identifier(c *Ps_type_identifierContext)

	// EnterSequence_identifier is called when entering the sequence_identifier production.
	EnterSequence_identifier(c *Sequence_identifierContext)

	// EnterSignal_identifier is called when entering the signal_identifier production.
	EnterSignal_identifier(c *Signal_identifierContext)

	// EnterSpecparam_identifier is called when entering the specparam_identifier production.
	EnterSpecparam_identifier(c *Specparam_identifierContext)

	// EnterTask_identifier is called when entering the task_identifier production.
	EnterTask_identifier(c *Task_identifierContext)

	// EnterTf_identifier is called when entering the tf_identifier production.
	EnterTf_identifier(c *Tf_identifierContext)

	// EnterTopmodule_identifier is called when entering the topmodule_identifier production.
	EnterTopmodule_identifier(c *Topmodule_identifierContext)

	// EnterType_identifier is called when entering the type_identifier production.
	EnterType_identifier(c *Type_identifierContext)

	// EnterVariable_identifier is called when entering the variable_identifier production.
	EnterVariable_identifier(c *Variable_identifierContext)

	// ExitLibrary_text is called when exiting the library_text production.
	ExitLibrary_text(c *Library_textContext)

	// ExitLibrary_description is called when exiting the library_description production.
	ExitLibrary_description(c *Library_descriptionContext)

	// ExitLibrary_declaration is called when exiting the library_declaration production.
	ExitLibrary_declaration(c *Library_declarationContext)

	// ExitInclude_statement is called when exiting the include_statement production.
	ExitInclude_statement(c *Include_statementContext)

	// ExitSource_text is called when exiting the source_text production.
	ExitSource_text(c *Source_textContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitModule_nonansi_header is called when exiting the module_nonansi_header production.
	ExitModule_nonansi_header(c *Module_nonansi_headerContext)

	// ExitModule_ansi_header is called when exiting the module_ansi_header production.
	ExitModule_ansi_header(c *Module_ansi_headerContext)

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitModule_keyword is called when exiting the module_keyword production.
	ExitModule_keyword(c *Module_keywordContext)

	// ExitInterface_declaration is called when exiting the interface_declaration production.
	ExitInterface_declaration(c *Interface_declarationContext)

	// ExitInterface_nonansi_header is called when exiting the interface_nonansi_header production.
	ExitInterface_nonansi_header(c *Interface_nonansi_headerContext)

	// ExitInterface_ansi_header is called when exiting the interface_ansi_header production.
	ExitInterface_ansi_header(c *Interface_ansi_headerContext)

	// ExitProgram_declaration is called when exiting the program_declaration production.
	ExitProgram_declaration(c *Program_declarationContext)

	// ExitProgram_nonansi_header is called when exiting the program_nonansi_header production.
	ExitProgram_nonansi_header(c *Program_nonansi_headerContext)

	// ExitProgram_ansi_header is called when exiting the program_ansi_header production.
	ExitProgram_ansi_header(c *Program_ansi_headerContext)

	// ExitChecker_declaration is called when exiting the checker_declaration production.
	ExitChecker_declaration(c *Checker_declarationContext)

	// ExitClass_declaration is called when exiting the class_declaration production.
	ExitClass_declaration(c *Class_declarationContext)

	// ExitInterface_class_type is called when exiting the interface_class_type production.
	ExitInterface_class_type(c *Interface_class_typeContext)

	// ExitInterface_class_declaration is called when exiting the interface_class_declaration production.
	ExitInterface_class_declaration(c *Interface_class_declarationContext)

	// ExitInterface_class_item is called when exiting the interface_class_item production.
	ExitInterface_class_item(c *Interface_class_itemContext)

	// ExitInterface_class_method is called when exiting the interface_class_method production.
	ExitInterface_class_method(c *Interface_class_methodContext)

	// ExitPackage_declaration is called when exiting the package_declaration production.
	ExitPackage_declaration(c *Package_declarationContext)

	// ExitTimeunits_declaration is called when exiting the timeunits_declaration production.
	ExitTimeunits_declaration(c *Timeunits_declarationContext)

	// ExitParameter_port_list is called when exiting the parameter_port_list production.
	ExitParameter_port_list(c *Parameter_port_listContext)

	// ExitParameter_port_declaration is called when exiting the parameter_port_declaration production.
	ExitParameter_port_declaration(c *Parameter_port_declarationContext)

	// ExitList_of_ports is called when exiting the list_of_ports production.
	ExitList_of_ports(c *List_of_portsContext)

	// ExitList_of_port_declarations is called when exiting the list_of_port_declarations production.
	ExitList_of_port_declarations(c *List_of_port_declarationsContext)

	// ExitPort_declaration is called when exiting the port_declaration production.
	ExitPort_declaration(c *Port_declarationContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitPort_expression is called when exiting the port_expression production.
	ExitPort_expression(c *Port_expressionContext)

	// ExitPort_reference is called when exiting the port_reference production.
	ExitPort_reference(c *Port_referenceContext)

	// ExitPort_direction is called when exiting the port_direction production.
	ExitPort_direction(c *Port_directionContext)

	// ExitNet_port_header is called when exiting the net_port_header production.
	ExitNet_port_header(c *Net_port_headerContext)

	// ExitVariable_port_header is called when exiting the variable_port_header production.
	ExitVariable_port_header(c *Variable_port_headerContext)

	// ExitInterface_port_header is called when exiting the interface_port_header production.
	ExitInterface_port_header(c *Interface_port_headerContext)

	// ExitAnsi_port_declaration is called when exiting the ansi_port_declaration production.
	ExitAnsi_port_declaration(c *Ansi_port_declarationContext)

	// ExitElaboration_system_task is called when exiting the elaboration_system_task production.
	ExitElaboration_system_task(c *Elaboration_system_taskContext)

	// ExitModule_common_item is called when exiting the module_common_item production.
	ExitModule_common_item(c *Module_common_itemContext)

	// ExitModule_item is called when exiting the module_item production.
	ExitModule_item(c *Module_itemContext)

	// ExitModule_or_generate_item is called when exiting the module_or_generate_item production.
	ExitModule_or_generate_item(c *Module_or_generate_itemContext)

	// ExitModule_or_generate_item_declaration is called when exiting the module_or_generate_item_declaration production.
	ExitModule_or_generate_item_declaration(c *Module_or_generate_item_declarationContext)

	// ExitNon_port_module_item is called when exiting the non_port_module_item production.
	ExitNon_port_module_item(c *Non_port_module_itemContext)

	// ExitParameter_override is called when exiting the parameter_override production.
	ExitParameter_override(c *Parameter_overrideContext)

	// ExitBind_directive is called when exiting the bind_directive production.
	ExitBind_directive(c *Bind_directiveContext)

	// ExitBind_target_scope is called when exiting the bind_target_scope production.
	ExitBind_target_scope(c *Bind_target_scopeContext)

	// ExitBind_target_instance is called when exiting the bind_target_instance production.
	ExitBind_target_instance(c *Bind_target_instanceContext)

	// ExitBind_target_instance_list is called when exiting the bind_target_instance_list production.
	ExitBind_target_instance_list(c *Bind_target_instance_listContext)

	// ExitBind_instantiation is called when exiting the bind_instantiation production.
	ExitBind_instantiation(c *Bind_instantiationContext)

	// ExitConfig_declaration is called when exiting the config_declaration production.
	ExitConfig_declaration(c *Config_declarationContext)

	// ExitDesign_statement is called when exiting the design_statement production.
	ExitDesign_statement(c *Design_statementContext)

	// ExitConfig_rule_statement is called when exiting the config_rule_statement production.
	ExitConfig_rule_statement(c *Config_rule_statementContext)

	// ExitDefault_clause is called when exiting the default_clause production.
	ExitDefault_clause(c *Default_clauseContext)

	// ExitInst_clause is called when exiting the inst_clause production.
	ExitInst_clause(c *Inst_clauseContext)

	// ExitInst_name is called when exiting the inst_name production.
	ExitInst_name(c *Inst_nameContext)

	// ExitCell_clause is called when exiting the cell_clause production.
	ExitCell_clause(c *Cell_clauseContext)

	// ExitLiblist_clause is called when exiting the liblist_clause production.
	ExitLiblist_clause(c *Liblist_clauseContext)

	// ExitUse_clause is called when exiting the use_clause production.
	ExitUse_clause(c *Use_clauseContext)

	// ExitInterface_or_generate_item is called when exiting the interface_or_generate_item production.
	ExitInterface_or_generate_item(c *Interface_or_generate_itemContext)

	// ExitExtern_tf_declaration is called when exiting the extern_tf_declaration production.
	ExitExtern_tf_declaration(c *Extern_tf_declarationContext)

	// ExitInterface_item is called when exiting the interface_item production.
	ExitInterface_item(c *Interface_itemContext)

	// ExitNon_port_interface_item is called when exiting the non_port_interface_item production.
	ExitNon_port_interface_item(c *Non_port_interface_itemContext)

	// ExitProgram_item is called when exiting the program_item production.
	ExitProgram_item(c *Program_itemContext)

	// ExitNon_port_program_item is called when exiting the non_port_program_item production.
	ExitNon_port_program_item(c *Non_port_program_itemContext)

	// ExitProgram_generate_item is called when exiting the program_generate_item production.
	ExitProgram_generate_item(c *Program_generate_itemContext)

	// ExitChecker_port_list is called when exiting the checker_port_list production.
	ExitChecker_port_list(c *Checker_port_listContext)

	// ExitChecker_port_item is called when exiting the checker_port_item production.
	ExitChecker_port_item(c *Checker_port_itemContext)

	// ExitChecker_port_direction is called when exiting the checker_port_direction production.
	ExitChecker_port_direction(c *Checker_port_directionContext)

	// ExitChecker_or_generate_item is called when exiting the checker_or_generate_item production.
	ExitChecker_or_generate_item(c *Checker_or_generate_itemContext)

	// ExitChecker_or_generate_item_declaration is called when exiting the checker_or_generate_item_declaration production.
	ExitChecker_or_generate_item_declaration(c *Checker_or_generate_item_declarationContext)

	// ExitChecker_generate_item is called when exiting the checker_generate_item production.
	ExitChecker_generate_item(c *Checker_generate_itemContext)

	// ExitClass_item is called when exiting the class_item production.
	ExitClass_item(c *Class_itemContext)

	// ExitClass_property is called when exiting the class_property production.
	ExitClass_property(c *Class_propertyContext)

	// ExitClass_method is called when exiting the class_method production.
	ExitClass_method(c *Class_methodContext)

	// ExitClass_constructor_prototype is called when exiting the class_constructor_prototype production.
	ExitClass_constructor_prototype(c *Class_constructor_prototypeContext)

	// ExitClass_constraint is called when exiting the class_constraint production.
	ExitClass_constraint(c *Class_constraintContext)

	// ExitClass_item_qualifier is called when exiting the class_item_qualifier production.
	ExitClass_item_qualifier(c *Class_item_qualifierContext)

	// ExitProperty_qualifier is called when exiting the property_qualifier production.
	ExitProperty_qualifier(c *Property_qualifierContext)

	// ExitRandom_qualifier is called when exiting the random_qualifier production.
	ExitRandom_qualifier(c *Random_qualifierContext)

	// ExitMethod_qualifier is called when exiting the method_qualifier production.
	ExitMethod_qualifier(c *Method_qualifierContext)

	// ExitMethod_prototype is called when exiting the method_prototype production.
	ExitMethod_prototype(c *Method_prototypeContext)

	// ExitClass_constructor_declaration is called when exiting the class_constructor_declaration production.
	ExitClass_constructor_declaration(c *Class_constructor_declarationContext)

	// ExitConstraint_declaration is called when exiting the constraint_declaration production.
	ExitConstraint_declaration(c *Constraint_declarationContext)

	// ExitConstraint_block is called when exiting the constraint_block production.
	ExitConstraint_block(c *Constraint_blockContext)

	// ExitConstraint_block_item is called when exiting the constraint_block_item production.
	ExitConstraint_block_item(c *Constraint_block_itemContext)

	// ExitSolve_before_list is called when exiting the solve_before_list production.
	ExitSolve_before_list(c *Solve_before_listContext)

	// ExitConstraint_primary is called when exiting the constraint_primary production.
	ExitConstraint_primary(c *Constraint_primaryContext)

	// ExitConstraint_expression is called when exiting the constraint_expression production.
	ExitConstraint_expression(c *Constraint_expressionContext)

	// ExitUniqueness_constraint is called when exiting the uniqueness_constraint production.
	ExitUniqueness_constraint(c *Uniqueness_constraintContext)

	// ExitConstraint_set is called when exiting the constraint_set production.
	ExitConstraint_set(c *Constraint_setContext)

	// ExitDist_list is called when exiting the dist_list production.
	ExitDist_list(c *Dist_listContext)

	// ExitDist_item is called when exiting the dist_item production.
	ExitDist_item(c *Dist_itemContext)

	// ExitDist_weight is called when exiting the dist_weight production.
	ExitDist_weight(c *Dist_weightContext)

	// ExitConstraint_prototype is called when exiting the constraint_prototype production.
	ExitConstraint_prototype(c *Constraint_prototypeContext)

	// ExitConstraint_prototype_qualifier is called when exiting the constraint_prototype_qualifier production.
	ExitConstraint_prototype_qualifier(c *Constraint_prototype_qualifierContext)

	// ExitExtern_constraint_declaration is called when exiting the extern_constraint_declaration production.
	ExitExtern_constraint_declaration(c *Extern_constraint_declarationContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitPackage_item is called when exiting the package_item production.
	ExitPackage_item(c *Package_itemContext)

	// ExitPackage_or_generate_item_declaration is called when exiting the package_or_generate_item_declaration production.
	ExitPackage_or_generate_item_declaration(c *Package_or_generate_item_declarationContext)

	// ExitAnonymous_program is called when exiting the anonymous_program production.
	ExitAnonymous_program(c *Anonymous_programContext)

	// ExitAnonymous_program_item is called when exiting the anonymous_program_item production.
	ExitAnonymous_program_item(c *Anonymous_program_itemContext)

	// ExitLocal_parameter_declaration is called when exiting the local_parameter_declaration production.
	ExitLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitSpecparam_declaration is called when exiting the specparam_declaration production.
	ExitSpecparam_declaration(c *Specparam_declarationContext)

	// ExitInout_declaration is called when exiting the inout_declaration production.
	ExitInout_declaration(c *Inout_declarationContext)

	// ExitInput_declaration is called when exiting the input_declaration production.
	ExitInput_declaration(c *Input_declarationContext)

	// ExitOutput_declaration is called when exiting the output_declaration production.
	ExitOutput_declaration(c *Output_declarationContext)

	// ExitInterface_port_declaration is called when exiting the interface_port_declaration production.
	ExitInterface_port_declaration(c *Interface_port_declarationContext)

	// ExitRef_declaration is called when exiting the ref_declaration production.
	ExitRef_declaration(c *Ref_declarationContext)

	// ExitData_declaration is called when exiting the data_declaration production.
	ExitData_declaration(c *Data_declarationContext)

	// ExitPackage_import_declaration is called when exiting the package_import_declaration production.
	ExitPackage_import_declaration(c *Package_import_declarationContext)

	// ExitPackage_import_item is called when exiting the package_import_item production.
	ExitPackage_import_item(c *Package_import_itemContext)

	// ExitPackage_export_declaration is called when exiting the package_export_declaration production.
	ExitPackage_export_declaration(c *Package_export_declarationContext)

	// ExitGenvar_declaration is called when exiting the genvar_declaration production.
	ExitGenvar_declaration(c *Genvar_declarationContext)

	// ExitNet_declaration is called when exiting the net_declaration production.
	ExitNet_declaration(c *Net_declarationContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitNet_type_declaration is called when exiting the net_type_declaration production.
	ExitNet_type_declaration(c *Net_type_declarationContext)

	// ExitLifetime is called when exiting the lifetime production.
	ExitLifetime(c *LifetimeContext)

	// ExitCasting_type is called when exiting the casting_type production.
	ExitCasting_type(c *Casting_typeContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitData_type_or_implicit is called when exiting the data_type_or_implicit production.
	ExitData_type_or_implicit(c *Data_type_or_implicitContext)

	// ExitImplicit_data_type is called when exiting the implicit_data_type production.
	ExitImplicit_data_type(c *Implicit_data_typeContext)

	// ExitEnum_base_type is called when exiting the enum_base_type production.
	ExitEnum_base_type(c *Enum_base_typeContext)

	// ExitEnum_name_declaration is called when exiting the enum_name_declaration production.
	ExitEnum_name_declaration(c *Enum_name_declarationContext)

	// ExitClass_scope is called when exiting the class_scope production.
	ExitClass_scope(c *Class_scopeContext)

	// ExitClass_type is called when exiting the class_type production.
	ExitClass_type(c *Class_typeContext)

	// ExitInteger_type is called when exiting the integer_type production.
	ExitInteger_type(c *Integer_typeContext)

	// ExitInteger_atom_type is called when exiting the integer_atom_type production.
	ExitInteger_atom_type(c *Integer_atom_typeContext)

	// ExitInteger_vector_type is called when exiting the integer_vector_type production.
	ExitInteger_vector_type(c *Integer_vector_typeContext)

	// ExitNon_integer_type is called when exiting the non_integer_type production.
	ExitNon_integer_type(c *Non_integer_typeContext)

	// ExitNet_type is called when exiting the net_type production.
	ExitNet_type(c *Net_typeContext)

	// ExitNet_port_type is called when exiting the net_port_type production.
	ExitNet_port_type(c *Net_port_typeContext)

	// ExitVariable_port_type is called when exiting the variable_port_type production.
	ExitVariable_port_type(c *Variable_port_typeContext)

	// ExitVar_data_type is called when exiting the var_data_type production.
	ExitVar_data_type(c *Var_data_typeContext)

	// ExitSigning is called when exiting the signing production.
	ExitSigning(c *SigningContext)

	// ExitSimple_type is called when exiting the simple_type production.
	ExitSimple_type(c *Simple_typeContext)

	// ExitStruct_union_member is called when exiting the struct_union_member production.
	ExitStruct_union_member(c *Struct_union_memberContext)

	// ExitData_type_or_void is called when exiting the data_type_or_void production.
	ExitData_type_or_void(c *Data_type_or_voidContext)

	// ExitStruct_union is called when exiting the struct_union production.
	ExitStruct_union(c *Struct_unionContext)

	// ExitType_reference is called when exiting the type_reference production.
	ExitType_reference(c *Type_referenceContext)

	// ExitDrive_strength is called when exiting the drive_strength production.
	ExitDrive_strength(c *Drive_strengthContext)

	// ExitStrength0 is called when exiting the strength0 production.
	ExitStrength0(c *Strength0Context)

	// ExitStrength1 is called when exiting the strength1 production.
	ExitStrength1(c *Strength1Context)

	// ExitCharge_strength is called when exiting the charge_strength production.
	ExitCharge_strength(c *Charge_strengthContext)

	// ExitDelay3 is called when exiting the delay3 production.
	ExitDelay3(c *Delay3Context)

	// ExitDelay2 is called when exiting the delay2 production.
	ExitDelay2(c *Delay2Context)

	// ExitDelay_value is called when exiting the delay_value production.
	ExitDelay_value(c *Delay_valueContext)

	// ExitList_of_defparam_assignments is called when exiting the list_of_defparam_assignments production.
	ExitList_of_defparam_assignments(c *List_of_defparam_assignmentsContext)

	// ExitList_of_genvar_identifiers is called when exiting the list_of_genvar_identifiers production.
	ExitList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// ExitList_of_interface_identifiers is called when exiting the list_of_interface_identifiers production.
	ExitList_of_interface_identifiers(c *List_of_interface_identifiersContext)

	// ExitList_of_net_decl_assignments is called when exiting the list_of_net_decl_assignments production.
	ExitList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// ExitList_of_param_assignments is called when exiting the list_of_param_assignments production.
	ExitList_of_param_assignments(c *List_of_param_assignmentsContext)

	// ExitList_of_port_identifiers is called when exiting the list_of_port_identifiers production.
	ExitList_of_port_identifiers(c *List_of_port_identifiersContext)

	// ExitList_of_udp_port_identifiers is called when exiting the list_of_udp_port_identifiers production.
	ExitList_of_udp_port_identifiers(c *List_of_udp_port_identifiersContext)

	// ExitList_of_specparam_assignments is called when exiting the list_of_specparam_assignments production.
	ExitList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// ExitList_of_tf_variable_identifiers is called when exiting the list_of_tf_variable_identifiers production.
	ExitList_of_tf_variable_identifiers(c *List_of_tf_variable_identifiersContext)

	// ExitList_of_type_assignments is called when exiting the list_of_type_assignments production.
	ExitList_of_type_assignments(c *List_of_type_assignmentsContext)

	// ExitList_of_variable_decl_assignments is called when exiting the list_of_variable_decl_assignments production.
	ExitList_of_variable_decl_assignments(c *List_of_variable_decl_assignmentsContext)

	// ExitList_of_variable_identifiers is called when exiting the list_of_variable_identifiers production.
	ExitList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// ExitList_of_variable_port_identifiers is called when exiting the list_of_variable_port_identifiers production.
	ExitList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// ExitDefparam_assignment is called when exiting the defparam_assignment production.
	ExitDefparam_assignment(c *Defparam_assignmentContext)

	// ExitNet_decl_assignment is called when exiting the net_decl_assignment production.
	ExitNet_decl_assignment(c *Net_decl_assignmentContext)

	// ExitParam_assignment is called when exiting the param_assignment production.
	ExitParam_assignment(c *Param_assignmentContext)

	// ExitSpecparam_assignment is called when exiting the specparam_assignment production.
	ExitSpecparam_assignment(c *Specparam_assignmentContext)

	// ExitType_assignment is called when exiting the type_assignment production.
	ExitType_assignment(c *Type_assignmentContext)

	// ExitPulse_control_specparam is called when exiting the pulse_control_specparam production.
	ExitPulse_control_specparam(c *Pulse_control_specparamContext)

	// ExitError_limit_value is called when exiting the error_limit_value production.
	ExitError_limit_value(c *Error_limit_valueContext)

	// ExitReject_limit_value is called when exiting the reject_limit_value production.
	ExitReject_limit_value(c *Reject_limit_valueContext)

	// ExitLimit_value is called when exiting the limit_value production.
	ExitLimit_value(c *Limit_valueContext)

	// ExitVariable_decl_assignment is called when exiting the variable_decl_assignment production.
	ExitVariable_decl_assignment(c *Variable_decl_assignmentContext)

	// ExitClass_new is called when exiting the class_new production.
	ExitClass_new(c *Class_newContext)

	// ExitDynamic_array_new is called when exiting the dynamic_array_new production.
	ExitDynamic_array_new(c *Dynamic_array_newContext)

	// ExitUnpacked_dimension is called when exiting the unpacked_dimension production.
	ExitUnpacked_dimension(c *Unpacked_dimensionContext)

	// ExitPacked_dimension is called when exiting the packed_dimension production.
	ExitPacked_dimension(c *Packed_dimensionContext)

	// ExitAssociative_dimension is called when exiting the associative_dimension production.
	ExitAssociative_dimension(c *Associative_dimensionContext)

	// ExitVariable_dimension is called when exiting the variable_dimension production.
	ExitVariable_dimension(c *Variable_dimensionContext)

	// ExitQueue_dimension is called when exiting the queue_dimension production.
	ExitQueue_dimension(c *Queue_dimensionContext)

	// ExitUnsized_dimension is called when exiting the unsized_dimension production.
	ExitUnsized_dimension(c *Unsized_dimensionContext)

	// ExitFunction_data_type_or_implicit is called when exiting the function_data_type_or_implicit production.
	ExitFunction_data_type_or_implicit(c *Function_data_type_or_implicitContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_body_declaration is called when exiting the function_body_declaration production.
	ExitFunction_body_declaration(c *Function_body_declarationContext)

	// ExitFunction_prototype is called when exiting the function_prototype production.
	ExitFunction_prototype(c *Function_prototypeContext)

	// ExitDpi_import_export is called when exiting the dpi_import_export production.
	ExitDpi_import_export(c *Dpi_import_exportContext)

	// ExitDpi_spec_string is called when exiting the dpi_spec_string production.
	ExitDpi_spec_string(c *Dpi_spec_stringContext)

	// ExitDpi_function_import_property is called when exiting the dpi_function_import_property production.
	ExitDpi_function_import_property(c *Dpi_function_import_propertyContext)

	// ExitDpi_task_import_property is called when exiting the dpi_task_import_property production.
	ExitDpi_task_import_property(c *Dpi_task_import_propertyContext)

	// ExitDpi_function_proto is called when exiting the dpi_function_proto production.
	ExitDpi_function_proto(c *Dpi_function_protoContext)

	// ExitDpi_task_proto is called when exiting the dpi_task_proto production.
	ExitDpi_task_proto(c *Dpi_task_protoContext)

	// ExitTask_declaration is called when exiting the task_declaration production.
	ExitTask_declaration(c *Task_declarationContext)

	// ExitTask_body_declaration is called when exiting the task_body_declaration production.
	ExitTask_body_declaration(c *Task_body_declarationContext)

	// ExitTf_item_declaration is called when exiting the tf_item_declaration production.
	ExitTf_item_declaration(c *Tf_item_declarationContext)

	// ExitTf_port_list is called when exiting the tf_port_list production.
	ExitTf_port_list(c *Tf_port_listContext)

	// ExitTf_port_item is called when exiting the tf_port_item production.
	ExitTf_port_item(c *Tf_port_itemContext)

	// ExitTf_port_direction is called when exiting the tf_port_direction production.
	ExitTf_port_direction(c *Tf_port_directionContext)

	// ExitTf_port_declaration is called when exiting the tf_port_declaration production.
	ExitTf_port_declaration(c *Tf_port_declarationContext)

	// ExitTask_prototype is called when exiting the task_prototype production.
	ExitTask_prototype(c *Task_prototypeContext)

	// ExitBlock_item_declaration is called when exiting the block_item_declaration production.
	ExitBlock_item_declaration(c *Block_item_declarationContext)

	// ExitModport_declaration is called when exiting the modport_declaration production.
	ExitModport_declaration(c *Modport_declarationContext)

	// ExitModport_item is called when exiting the modport_item production.
	ExitModport_item(c *Modport_itemContext)

	// ExitModport_ports_declaration is called when exiting the modport_ports_declaration production.
	ExitModport_ports_declaration(c *Modport_ports_declarationContext)

	// ExitModport_clocking_declaration is called when exiting the modport_clocking_declaration production.
	ExitModport_clocking_declaration(c *Modport_clocking_declarationContext)

	// ExitModport_simple_ports_declaration is called when exiting the modport_simple_ports_declaration production.
	ExitModport_simple_ports_declaration(c *Modport_simple_ports_declarationContext)

	// ExitModport_simple_port is called when exiting the modport_simple_port production.
	ExitModport_simple_port(c *Modport_simple_portContext)

	// ExitModport_tf_ports_declaration is called when exiting the modport_tf_ports_declaration production.
	ExitModport_tf_ports_declaration(c *Modport_tf_ports_declarationContext)

	// ExitModport_tf_port is called when exiting the modport_tf_port production.
	ExitModport_tf_port(c *Modport_tf_portContext)

	// ExitImport_export is called when exiting the import_export production.
	ExitImport_export(c *Import_exportContext)

	// ExitConcurrent_assertion_item is called when exiting the concurrent_assertion_item production.
	ExitConcurrent_assertion_item(c *Concurrent_assertion_itemContext)

	// ExitConcurrent_assertion_statement is called when exiting the concurrent_assertion_statement production.
	ExitConcurrent_assertion_statement(c *Concurrent_assertion_statementContext)

	// ExitAssert_property_statement is called when exiting the assert_property_statement production.
	ExitAssert_property_statement(c *Assert_property_statementContext)

	// ExitAssume_property_statement is called when exiting the assume_property_statement production.
	ExitAssume_property_statement(c *Assume_property_statementContext)

	// ExitCover_property_statement is called when exiting the cover_property_statement production.
	ExitCover_property_statement(c *Cover_property_statementContext)

	// ExitExpect_property_statement is called when exiting the expect_property_statement production.
	ExitExpect_property_statement(c *Expect_property_statementContext)

	// ExitCover_sequence_statement is called when exiting the cover_sequence_statement production.
	ExitCover_sequence_statement(c *Cover_sequence_statementContext)

	// ExitRestrict_property_statement is called when exiting the restrict_property_statement production.
	ExitRestrict_property_statement(c *Restrict_property_statementContext)

	// ExitProperty_instance is called when exiting the property_instance production.
	ExitProperty_instance(c *Property_instanceContext)

	// ExitProperty_list_of_arguments is called when exiting the property_list_of_arguments production.
	ExitProperty_list_of_arguments(c *Property_list_of_argumentsContext)

	// ExitProperty_actual_arg is called when exiting the property_actual_arg production.
	ExitProperty_actual_arg(c *Property_actual_argContext)

	// ExitAssertion_item_declaration is called when exiting the assertion_item_declaration production.
	ExitAssertion_item_declaration(c *Assertion_item_declarationContext)

	// ExitProperty_declaration is called when exiting the property_declaration production.
	ExitProperty_declaration(c *Property_declarationContext)

	// ExitProperty_port_list is called when exiting the property_port_list production.
	ExitProperty_port_list(c *Property_port_listContext)

	// ExitProperty_port_item is called when exiting the property_port_item production.
	ExitProperty_port_item(c *Property_port_itemContext)

	// ExitProperty_lvar_port_direction is called when exiting the property_lvar_port_direction production.
	ExitProperty_lvar_port_direction(c *Property_lvar_port_directionContext)

	// ExitProperty_formal_type is called when exiting the property_formal_type production.
	ExitProperty_formal_type(c *Property_formal_typeContext)

	// ExitProperty_spec is called when exiting the property_spec production.
	ExitProperty_spec(c *Property_specContext)

	// ExitProperty_expr is called when exiting the property_expr production.
	ExitProperty_expr(c *Property_exprContext)

	// ExitProperty_case_item is called when exiting the property_case_item production.
	ExitProperty_case_item(c *Property_case_itemContext)

	// ExitSequence_declaration is called when exiting the sequence_declaration production.
	ExitSequence_declaration(c *Sequence_declarationContext)

	// ExitSequence_port_list is called when exiting the sequence_port_list production.
	ExitSequence_port_list(c *Sequence_port_listContext)

	// ExitSequence_port_item is called when exiting the sequence_port_item production.
	ExitSequence_port_item(c *Sequence_port_itemContext)

	// ExitSequence_lvar_port_direction is called when exiting the sequence_lvar_port_direction production.
	ExitSequence_lvar_port_direction(c *Sequence_lvar_port_directionContext)

	// ExitSequence_formal_type is called when exiting the sequence_formal_type production.
	ExitSequence_formal_type(c *Sequence_formal_typeContext)

	// ExitSequence_expr is called when exiting the sequence_expr production.
	ExitSequence_expr(c *Sequence_exprContext)

	// ExitCycle_delay_range is called when exiting the cycle_delay_range production.
	ExitCycle_delay_range(c *Cycle_delay_rangeContext)

	// ExitSequence_method_call is called when exiting the sequence_method_call production.
	ExitSequence_method_call(c *Sequence_method_callContext)

	// ExitSequence_match_item is called when exiting the sequence_match_item production.
	ExitSequence_match_item(c *Sequence_match_itemContext)

	// ExitSequence_instance is called when exiting the sequence_instance production.
	ExitSequence_instance(c *Sequence_instanceContext)

	// ExitSequence_list_of_arguments is called when exiting the sequence_list_of_arguments production.
	ExitSequence_list_of_arguments(c *Sequence_list_of_argumentsContext)

	// ExitSequence_actual_arg is called when exiting the sequence_actual_arg production.
	ExitSequence_actual_arg(c *Sequence_actual_argContext)

	// ExitBoolean_abbrev is called when exiting the boolean_abbrev production.
	ExitBoolean_abbrev(c *Boolean_abbrevContext)

	// ExitSequence_abbrev is called when exiting the sequence_abbrev production.
	ExitSequence_abbrev(c *Sequence_abbrevContext)

	// ExitConsecutive_repetition is called when exiting the consecutive_repetition production.
	ExitConsecutive_repetition(c *Consecutive_repetitionContext)

	// ExitNon_consecutive_repetition is called when exiting the non_consecutive_repetition production.
	ExitNon_consecutive_repetition(c *Non_consecutive_repetitionContext)

	// ExitGoto_repetition is called when exiting the goto_repetition production.
	ExitGoto_repetition(c *Goto_repetitionContext)

	// ExitConst_or_range_expression is called when exiting the const_or_range_expression production.
	ExitConst_or_range_expression(c *Const_or_range_expressionContext)

	// ExitCycle_delay_const_range_expression is called when exiting the cycle_delay_const_range_expression production.
	ExitCycle_delay_const_range_expression(c *Cycle_delay_const_range_expressionContext)

	// ExitExpression_or_dist is called when exiting the expression_or_dist production.
	ExitExpression_or_dist(c *Expression_or_distContext)

	// ExitAssertion_variable_declaration is called when exiting the assertion_variable_declaration production.
	ExitAssertion_variable_declaration(c *Assertion_variable_declarationContext)

	// ExitCovergroup_declaration is called when exiting the covergroup_declaration production.
	ExitCovergroup_declaration(c *Covergroup_declarationContext)

	// ExitCoverage_spec_or_option is called when exiting the coverage_spec_or_option production.
	ExitCoverage_spec_or_option(c *Coverage_spec_or_optionContext)

	// ExitCoverage_option is called when exiting the coverage_option production.
	ExitCoverage_option(c *Coverage_optionContext)

	// ExitCoverage_spec is called when exiting the coverage_spec production.
	ExitCoverage_spec(c *Coverage_specContext)

	// ExitCoverage_event is called when exiting the coverage_event production.
	ExitCoverage_event(c *Coverage_eventContext)

	// ExitBlock_event_expression is called when exiting the block_event_expression production.
	ExitBlock_event_expression(c *Block_event_expressionContext)

	// ExitHierarchical_btf_identifier is called when exiting the hierarchical_btf_identifier production.
	ExitHierarchical_btf_identifier(c *Hierarchical_btf_identifierContext)

	// ExitCover_point is called when exiting the cover_point production.
	ExitCover_point(c *Cover_pointContext)

	// ExitBins_or_empty is called when exiting the bins_or_empty production.
	ExitBins_or_empty(c *Bins_or_emptyContext)

	// ExitBins_or_options is called when exiting the bins_or_options production.
	ExitBins_or_options(c *Bins_or_optionsContext)

	// ExitBins_keyword is called when exiting the bins_keyword production.
	ExitBins_keyword(c *Bins_keywordContext)

	// ExitTrans_list is called when exiting the trans_list production.
	ExitTrans_list(c *Trans_listContext)

	// ExitTrans_set is called when exiting the trans_set production.
	ExitTrans_set(c *Trans_setContext)

	// ExitTrans_range_list is called when exiting the trans_range_list production.
	ExitTrans_range_list(c *Trans_range_listContext)

	// ExitTrans_item is called when exiting the trans_item production.
	ExitTrans_item(c *Trans_itemContext)

	// ExitRepeat_range is called when exiting the repeat_range production.
	ExitRepeat_range(c *Repeat_rangeContext)

	// ExitCover_cross is called when exiting the cover_cross production.
	ExitCover_cross(c *Cover_crossContext)

	// ExitList_of_cross_items is called when exiting the list_of_cross_items production.
	ExitList_of_cross_items(c *List_of_cross_itemsContext)

	// ExitCross_item is called when exiting the cross_item production.
	ExitCross_item(c *Cross_itemContext)

	// ExitCross_body is called when exiting the cross_body production.
	ExitCross_body(c *Cross_bodyContext)

	// ExitCross_body_item is called when exiting the cross_body_item production.
	ExitCross_body_item(c *Cross_body_itemContext)

	// ExitBins_selection_or_option is called when exiting the bins_selection_or_option production.
	ExitBins_selection_or_option(c *Bins_selection_or_optionContext)

	// ExitBins_selection is called when exiting the bins_selection production.
	ExitBins_selection(c *Bins_selectionContext)

	// ExitSelect_expression is called when exiting the select_expression production.
	ExitSelect_expression(c *Select_expressionContext)

	// ExitSelect_condition is called when exiting the select_condition production.
	ExitSelect_condition(c *Select_conditionContext)

	// ExitBins_expression is called when exiting the bins_expression production.
	ExitBins_expression(c *Bins_expressionContext)

	// ExitCovergroup_range_list is called when exiting the covergroup_range_list production.
	ExitCovergroup_range_list(c *Covergroup_range_listContext)

	// ExitCovergroup_value_range is called when exiting the covergroup_value_range production.
	ExitCovergroup_value_range(c *Covergroup_value_rangeContext)

	// ExitWith_covergroup_expression is called when exiting the with_covergroup_expression production.
	ExitWith_covergroup_expression(c *With_covergroup_expressionContext)

	// ExitSet_covergroup_expression is called when exiting the set_covergroup_expression production.
	ExitSet_covergroup_expression(c *Set_covergroup_expressionContext)

	// ExitInteger_covergroup_expression is called when exiting the integer_covergroup_expression production.
	ExitInteger_covergroup_expression(c *Integer_covergroup_expressionContext)

	// ExitCross_set_expression is called when exiting the cross_set_expression production.
	ExitCross_set_expression(c *Cross_set_expressionContext)

	// ExitCovergroup_expression is called when exiting the covergroup_expression production.
	ExitCovergroup_expression(c *Covergroup_expressionContext)

	// ExitLet_declaration is called when exiting the let_declaration production.
	ExitLet_declaration(c *Let_declarationContext)

	// ExitLet_identifier is called when exiting the let_identifier production.
	ExitLet_identifier(c *Let_identifierContext)

	// ExitLet_port_list is called when exiting the let_port_list production.
	ExitLet_port_list(c *Let_port_listContext)

	// ExitLet_port_item is called when exiting the let_port_item production.
	ExitLet_port_item(c *Let_port_itemContext)

	// ExitLet_formal_type is called when exiting the let_formal_type production.
	ExitLet_formal_type(c *Let_formal_typeContext)

	// ExitLet_expression is called when exiting the let_expression production.
	ExitLet_expression(c *Let_expressionContext)

	// ExitLet_list_of_arguments is called when exiting the let_list_of_arguments production.
	ExitLet_list_of_arguments(c *Let_list_of_argumentsContext)

	// ExitLet_actual_arg is called when exiting the let_actual_arg production.
	ExitLet_actual_arg(c *Let_actual_argContext)

	// ExitGate_instantiation is called when exiting the gate_instantiation production.
	ExitGate_instantiation(c *Gate_instantiationContext)

	// ExitCmos_switch_instance is called when exiting the cmos_switch_instance production.
	ExitCmos_switch_instance(c *Cmos_switch_instanceContext)

	// ExitEnable_gate_instance is called when exiting the enable_gate_instance production.
	ExitEnable_gate_instance(c *Enable_gate_instanceContext)

	// ExitMos_switch_instance is called when exiting the mos_switch_instance production.
	ExitMos_switch_instance(c *Mos_switch_instanceContext)

	// ExitN_input_gate_instance is called when exiting the n_input_gate_instance production.
	ExitN_input_gate_instance(c *N_input_gate_instanceContext)

	// ExitN_output_gate_instance is called when exiting the n_output_gate_instance production.
	ExitN_output_gate_instance(c *N_output_gate_instanceContext)

	// ExitPass_switch_instance is called when exiting the pass_switch_instance production.
	ExitPass_switch_instance(c *Pass_switch_instanceContext)

	// ExitPass_enable_switch_instance is called when exiting the pass_enable_switch_instance production.
	ExitPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// ExitPull_gate_instance is called when exiting the pull_gate_instance production.
	ExitPull_gate_instance(c *Pull_gate_instanceContext)

	// ExitPulldown_strength is called when exiting the pulldown_strength production.
	ExitPulldown_strength(c *Pulldown_strengthContext)

	// ExitPullup_strength is called when exiting the pullup_strength production.
	ExitPullup_strength(c *Pullup_strengthContext)

	// ExitEnable_terminal is called when exiting the enable_terminal production.
	ExitEnable_terminal(c *Enable_terminalContext)

	// ExitInout_terminal is called when exiting the inout_terminal production.
	ExitInout_terminal(c *Inout_terminalContext)

	// ExitInput_terminal is called when exiting the input_terminal production.
	ExitInput_terminal(c *Input_terminalContext)

	// ExitNcontrol_terminal is called when exiting the ncontrol_terminal production.
	ExitNcontrol_terminal(c *Ncontrol_terminalContext)

	// ExitOutput_terminal is called when exiting the output_terminal production.
	ExitOutput_terminal(c *Output_terminalContext)

	// ExitPcontrol_terminal is called when exiting the pcontrol_terminal production.
	ExitPcontrol_terminal(c *Pcontrol_terminalContext)

	// ExitCmos_switchtype is called when exiting the cmos_switchtype production.
	ExitCmos_switchtype(c *Cmos_switchtypeContext)

	// ExitEnable_gatetype is called when exiting the enable_gatetype production.
	ExitEnable_gatetype(c *Enable_gatetypeContext)

	// ExitMos_switchtype is called when exiting the mos_switchtype production.
	ExitMos_switchtype(c *Mos_switchtypeContext)

	// ExitN_input_gatetype is called when exiting the n_input_gatetype production.
	ExitN_input_gatetype(c *N_input_gatetypeContext)

	// ExitN_output_gatetype is called when exiting the n_output_gatetype production.
	ExitN_output_gatetype(c *N_output_gatetypeContext)

	// ExitPass_en_switchtype is called when exiting the pass_en_switchtype production.
	ExitPass_en_switchtype(c *Pass_en_switchtypeContext)

	// ExitPass_switchtype is called when exiting the pass_switchtype production.
	ExitPass_switchtype(c *Pass_switchtypeContext)

	// ExitModule_instantiation is called when exiting the module_instantiation production.
	ExitModule_instantiation(c *Module_instantiationContext)

	// ExitParameter_value_assignment is called when exiting the parameter_value_assignment production.
	ExitParameter_value_assignment(c *Parameter_value_assignmentContext)

	// ExitList_of_parameter_assignments is called when exiting the list_of_parameter_assignments production.
	ExitList_of_parameter_assignments(c *List_of_parameter_assignmentsContext)

	// ExitOrdered_parameter_assignment is called when exiting the ordered_parameter_assignment production.
	ExitOrdered_parameter_assignment(c *Ordered_parameter_assignmentContext)

	// ExitNamed_parameter_assignment is called when exiting the named_parameter_assignment production.
	ExitNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// ExitHierarchical_instance is called when exiting the hierarchical_instance production.
	ExitHierarchical_instance(c *Hierarchical_instanceContext)

	// ExitName_of_instance is called when exiting the name_of_instance production.
	ExitName_of_instance(c *Name_of_instanceContext)

	// ExitList_of_port_connections is called when exiting the list_of_port_connections production.
	ExitList_of_port_connections(c *List_of_port_connectionsContext)

	// ExitOrdered_port_connection is called when exiting the ordered_port_connection production.
	ExitOrdered_port_connection(c *Ordered_port_connectionContext)

	// ExitNamed_port_connection is called when exiting the named_port_connection production.
	ExitNamed_port_connection(c *Named_port_connectionContext)

	// ExitInterface_instantiation is called when exiting the interface_instantiation production.
	ExitInterface_instantiation(c *Interface_instantiationContext)

	// ExitProgram_instantiation is called when exiting the program_instantiation production.
	ExitProgram_instantiation(c *Program_instantiationContext)

	// ExitChecker_instantiation is called when exiting the checker_instantiation production.
	ExitChecker_instantiation(c *Checker_instantiationContext)

	// ExitList_of_checker_port_connections is called when exiting the list_of_checker_port_connections production.
	ExitList_of_checker_port_connections(c *List_of_checker_port_connectionsContext)

	// ExitOrdered_checker_port_connection is called when exiting the ordered_checker_port_connection production.
	ExitOrdered_checker_port_connection(c *Ordered_checker_port_connectionContext)

	// ExitNamed_checker_port_connection is called when exiting the named_checker_port_connection production.
	ExitNamed_checker_port_connection(c *Named_checker_port_connectionContext)

	// ExitGenerate_region is called when exiting the generate_region production.
	ExitGenerate_region(c *Generate_regionContext)

	// ExitLoop_generate_construct is called when exiting the loop_generate_construct production.
	ExitLoop_generate_construct(c *Loop_generate_constructContext)

	// ExitGenvar_initialization is called when exiting the genvar_initialization production.
	ExitGenvar_initialization(c *Genvar_initializationContext)

	// ExitGenvar_iteration is called when exiting the genvar_iteration production.
	ExitGenvar_iteration(c *Genvar_iterationContext)

	// ExitConditional_generate_construct is called when exiting the conditional_generate_construct production.
	ExitConditional_generate_construct(c *Conditional_generate_constructContext)

	// ExitIf_generate_construct is called when exiting the if_generate_construct production.
	ExitIf_generate_construct(c *If_generate_constructContext)

	// ExitCase_generate_construct is called when exiting the case_generate_construct production.
	ExitCase_generate_construct(c *Case_generate_constructContext)

	// ExitCase_generate_item is called when exiting the case_generate_item production.
	ExitCase_generate_item(c *Case_generate_itemContext)

	// ExitGenerate_block is called when exiting the generate_block production.
	ExitGenerate_block(c *Generate_blockContext)

	// ExitGenerate_item is called when exiting the generate_item production.
	ExitGenerate_item(c *Generate_itemContext)

	// ExitContinuous_assign is called when exiting the continuous_assign production.
	ExitContinuous_assign(c *Continuous_assignContext)

	// ExitList_of_net_assignments is called when exiting the list_of_net_assignments production.
	ExitList_of_net_assignments(c *List_of_net_assignmentsContext)

	// ExitList_of_variable_assignments is called when exiting the list_of_variable_assignments production.
	ExitList_of_variable_assignments(c *List_of_variable_assignmentsContext)

	// ExitNet_alias is called when exiting the net_alias production.
	ExitNet_alias(c *Net_aliasContext)

	// ExitNet_assignment is called when exiting the net_assignment production.
	ExitNet_assignment(c *Net_assignmentContext)

	// ExitInitial_construct is called when exiting the initial_construct production.
	ExitInitial_construct(c *Initial_constructContext)

	// ExitAlways_construct is called when exiting the always_construct production.
	ExitAlways_construct(c *Always_constructContext)

	// ExitAlways_keyword is called when exiting the always_keyword production.
	ExitAlways_keyword(c *Always_keywordContext)

	// ExitFinal_construct is called when exiting the final_construct production.
	ExitFinal_construct(c *Final_constructContext)

	// ExitBlocking_assignment is called when exiting the blocking_assignment production.
	ExitBlocking_assignment(c *Blocking_assignmentContext)

	// ExitOperator_assignment is called when exiting the operator_assignment production.
	ExitOperator_assignment(c *Operator_assignmentContext)

	// ExitAssignment_operator is called when exiting the assignment_operator production.
	ExitAssignment_operator(c *Assignment_operatorContext)

	// ExitNonblocking_assignment is called when exiting the nonblocking_assignment production.
	ExitNonblocking_assignment(c *Nonblocking_assignmentContext)

	// ExitProcedural_continuous_assignment is called when exiting the procedural_continuous_assignment production.
	ExitProcedural_continuous_assignment(c *Procedural_continuous_assignmentContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitAction_block is called when exiting the action_block production.
	ExitAction_block(c *Action_blockContext)

	// ExitSeq_block is called when exiting the seq_block production.
	ExitSeq_block(c *Seq_blockContext)

	// ExitPar_block is called when exiting the par_block production.
	ExitPar_block(c *Par_blockContext)

	// ExitJoin_keyword is called when exiting the join_keyword production.
	ExitJoin_keyword(c *Join_keywordContext)

	// ExitStatement_or_null is called when exiting the statement_or_null production.
	ExitStatement_or_null(c *Statement_or_nullContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatement_item is called when exiting the statement_item production.
	ExitStatement_item(c *Statement_itemContext)

	// ExitFunction_statement is called when exiting the function_statement production.
	ExitFunction_statement(c *Function_statementContext)

	// ExitFunction_statement_or_null is called when exiting the function_statement_or_null production.
	ExitFunction_statement_or_null(c *Function_statement_or_nullContext)

	// ExitVariable_identifier_list is called when exiting the variable_identifier_list production.
	ExitVariable_identifier_list(c *Variable_identifier_listContext)

	// ExitProcedural_timing_control_statement is called when exiting the procedural_timing_control_statement production.
	ExitProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// ExitDelay_or_event_control is called when exiting the delay_or_event_control production.
	ExitDelay_or_event_control(c *Delay_or_event_controlContext)

	// ExitDelay_control is called when exiting the delay_control production.
	ExitDelay_control(c *Delay_controlContext)

	// ExitEvent_control is called when exiting the event_control production.
	ExitEvent_control(c *Event_controlContext)

	// ExitEvent_expression is called when exiting the event_expression production.
	ExitEvent_expression(c *Event_expressionContext)

	// ExitProcedural_timing_control is called when exiting the procedural_timing_control production.
	ExitProcedural_timing_control(c *Procedural_timing_controlContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitEvent_trigger is called when exiting the event_trigger production.
	ExitEvent_trigger(c *Event_triggerContext)

	// ExitDisable_statement is called when exiting the disable_statement production.
	ExitDisable_statement(c *Disable_statementContext)

	// ExitConditional_statement is called when exiting the conditional_statement production.
	ExitConditional_statement(c *Conditional_statementContext)

	// ExitUnique_priority is called when exiting the unique_priority production.
	ExitUnique_priority(c *Unique_priorityContext)

	// ExitCond_predicate is called when exiting the cond_predicate production.
	ExitCond_predicate(c *Cond_predicateContext)

	// ExitExpression_or_cond_pattern is called when exiting the expression_or_cond_pattern production.
	ExitExpression_or_cond_pattern(c *Expression_or_cond_patternContext)

	// ExitCond_pattern is called when exiting the cond_pattern production.
	ExitCond_pattern(c *Cond_patternContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_keyword is called when exiting the case_keyword production.
	ExitCase_keyword(c *Case_keywordContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitCase_item is called when exiting the case_item production.
	ExitCase_item(c *Case_itemContext)

	// ExitCase_pattern_item is called when exiting the case_pattern_item production.
	ExitCase_pattern_item(c *Case_pattern_itemContext)

	// ExitCase_inside_item is called when exiting the case_inside_item production.
	ExitCase_inside_item(c *Case_inside_itemContext)

	// ExitCase_item_expression is called when exiting the case_item_expression production.
	ExitCase_item_expression(c *Case_item_expressionContext)

	// ExitRandcase_statement is called when exiting the randcase_statement production.
	ExitRandcase_statement(c *Randcase_statementContext)

	// ExitRandcase_item is called when exiting the randcase_item production.
	ExitRandcase_item(c *Randcase_itemContext)

	// ExitOpen_range_list is called when exiting the open_range_list production.
	ExitOpen_range_list(c *Open_range_listContext)

	// ExitOpen_value_range is called when exiting the open_value_range production.
	ExitOpen_value_range(c *Open_value_rangeContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitAssignment_pattern is called when exiting the assignment_pattern production.
	ExitAssignment_pattern(c *Assignment_patternContext)

	// ExitStructure_pattern_key is called when exiting the structure_pattern_key production.
	ExitStructure_pattern_key(c *Structure_pattern_keyContext)

	// ExitArray_pattern_key is called when exiting the array_pattern_key production.
	ExitArray_pattern_key(c *Array_pattern_keyContext)

	// ExitAssignment_pattern_key is called when exiting the assignment_pattern_key production.
	ExitAssignment_pattern_key(c *Assignment_pattern_keyContext)

	// ExitAssignment_pattern_expression is called when exiting the assignment_pattern_expression production.
	ExitAssignment_pattern_expression(c *Assignment_pattern_expressionContext)

	// ExitAssignment_pattern_expression_type is called when exiting the assignment_pattern_expression_type production.
	ExitAssignment_pattern_expression_type(c *Assignment_pattern_expression_typeContext)

	// ExitConstant_assignment_pattern_expression is called when exiting the constant_assignment_pattern_expression production.
	ExitConstant_assignment_pattern_expression(c *Constant_assignment_pattern_expressionContext)

	// ExitAssignment_pattern_net_lvalue is called when exiting the assignment_pattern_net_lvalue production.
	ExitAssignment_pattern_net_lvalue(c *Assignment_pattern_net_lvalueContext)

	// ExitAssignment_pattern_variable_lvalue is called when exiting the assignment_pattern_variable_lvalue production.
	ExitAssignment_pattern_variable_lvalue(c *Assignment_pattern_variable_lvalueContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitFor_initialization is called when exiting the for_initialization production.
	ExitFor_initialization(c *For_initializationContext)

	// ExitFor_variable_declaration is called when exiting the for_variable_declaration production.
	ExitFor_variable_declaration(c *For_variable_declarationContext)

	// ExitFor_step is called when exiting the for_step production.
	ExitFor_step(c *For_stepContext)

	// ExitFor_step_assignment is called when exiting the for_step_assignment production.
	ExitFor_step_assignment(c *For_step_assignmentContext)

	// ExitLoop_variables is called when exiting the loop_variables production.
	ExitLoop_variables(c *Loop_variablesContext)

	// ExitSubroutine_call_statement is called when exiting the subroutine_call_statement production.
	ExitSubroutine_call_statement(c *Subroutine_call_statementContext)

	// ExitAssertion_item is called when exiting the assertion_item production.
	ExitAssertion_item(c *Assertion_itemContext)

	// ExitDeferred_immediate_assertion_item is called when exiting the deferred_immediate_assertion_item production.
	ExitDeferred_immediate_assertion_item(c *Deferred_immediate_assertion_itemContext)

	// ExitProcedural_assertion_statement is called when exiting the procedural_assertion_statement production.
	ExitProcedural_assertion_statement(c *Procedural_assertion_statementContext)

	// ExitImmediate_assertion_statement is called when exiting the immediate_assertion_statement production.
	ExitImmediate_assertion_statement(c *Immediate_assertion_statementContext)

	// ExitSimple_immediate_assertion_statement is called when exiting the simple_immediate_assertion_statement production.
	ExitSimple_immediate_assertion_statement(c *Simple_immediate_assertion_statementContext)

	// ExitSimple_immediate_assert_statement is called when exiting the simple_immediate_assert_statement production.
	ExitSimple_immediate_assert_statement(c *Simple_immediate_assert_statementContext)

	// ExitSimple_immediate_assume_statement is called when exiting the simple_immediate_assume_statement production.
	ExitSimple_immediate_assume_statement(c *Simple_immediate_assume_statementContext)

	// ExitSimple_immediate_cover_statement is called when exiting the simple_immediate_cover_statement production.
	ExitSimple_immediate_cover_statement(c *Simple_immediate_cover_statementContext)

	// ExitDeferred_immediate_assertion_statement is called when exiting the deferred_immediate_assertion_statement production.
	ExitDeferred_immediate_assertion_statement(c *Deferred_immediate_assertion_statementContext)

	// ExitDeferred_immediate_assert_statement is called when exiting the deferred_immediate_assert_statement production.
	ExitDeferred_immediate_assert_statement(c *Deferred_immediate_assert_statementContext)

	// ExitDeferred_immediate_assume_statement is called when exiting the deferred_immediate_assume_statement production.
	ExitDeferred_immediate_assume_statement(c *Deferred_immediate_assume_statementContext)

	// ExitDeferred_immediate_cover_statement is called when exiting the deferred_immediate_cover_statement production.
	ExitDeferred_immediate_cover_statement(c *Deferred_immediate_cover_statementContext)

	// ExitClocking_declaration is called when exiting the clocking_declaration production.
	ExitClocking_declaration(c *Clocking_declarationContext)

	// ExitClocking_event is called when exiting the clocking_event production.
	ExitClocking_event(c *Clocking_eventContext)

	// ExitClocking_item is called when exiting the clocking_item production.
	ExitClocking_item(c *Clocking_itemContext)

	// ExitDefault_skew is called when exiting the default_skew production.
	ExitDefault_skew(c *Default_skewContext)

	// ExitClocking_direction is called when exiting the clocking_direction production.
	ExitClocking_direction(c *Clocking_directionContext)

	// ExitList_of_clocking_decl_assign is called when exiting the list_of_clocking_decl_assign production.
	ExitList_of_clocking_decl_assign(c *List_of_clocking_decl_assignContext)

	// ExitClocking_decl_assign is called when exiting the clocking_decl_assign production.
	ExitClocking_decl_assign(c *Clocking_decl_assignContext)

	// ExitClocking_skew is called when exiting the clocking_skew production.
	ExitClocking_skew(c *Clocking_skewContext)

	// ExitClocking_drive is called when exiting the clocking_drive production.
	ExitClocking_drive(c *Clocking_driveContext)

	// ExitCycle_delay is called when exiting the cycle_delay production.
	ExitCycle_delay(c *Cycle_delayContext)

	// ExitClockvar is called when exiting the clockvar production.
	ExitClockvar(c *ClockvarContext)

	// ExitClockvar_expression is called when exiting the clockvar_expression production.
	ExitClockvar_expression(c *Clockvar_expressionContext)

	// ExitRandsequence_statement is called when exiting the randsequence_statement production.
	ExitRandsequence_statement(c *Randsequence_statementContext)

	// ExitProduction is called when exiting the production production.
	ExitProduction(c *ProductionContext)

	// ExitRs_rule is called when exiting the rs_rule production.
	ExitRs_rule(c *Rs_ruleContext)

	// ExitRs_production_list is called when exiting the rs_production_list production.
	ExitRs_production_list(c *Rs_production_listContext)

	// ExitWeight_specification is called when exiting the weight_specification production.
	ExitWeight_specification(c *Weight_specificationContext)

	// ExitRs_code_block is called when exiting the rs_code_block production.
	ExitRs_code_block(c *Rs_code_blockContext)

	// ExitRs_prod is called when exiting the rs_prod production.
	ExitRs_prod(c *Rs_prodContext)

	// ExitProduction_item is called when exiting the production_item production.
	ExitProduction_item(c *Production_itemContext)

	// ExitRs_if_else is called when exiting the rs_if_else production.
	ExitRs_if_else(c *Rs_if_elseContext)

	// ExitRs_repeat is called when exiting the rs_repeat production.
	ExitRs_repeat(c *Rs_repeatContext)

	// ExitRs_case is called when exiting the rs_case production.
	ExitRs_case(c *Rs_caseContext)

	// ExitRs_case_item is called when exiting the rs_case_item production.
	ExitRs_case_item(c *Rs_case_itemContext)

	// ExitSpecify_block is called when exiting the specify_block production.
	ExitSpecify_block(c *Specify_blockContext)

	// ExitSpecify_item is called when exiting the specify_item production.
	ExitSpecify_item(c *Specify_itemContext)

	// ExitPulsestyle_declaration is called when exiting the pulsestyle_declaration production.
	ExitPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// ExitShowcancelled_declaration is called when exiting the showcancelled_declaration production.
	ExitShowcancelled_declaration(c *Showcancelled_declarationContext)

	// ExitPath_declaration is called when exiting the path_declaration production.
	ExitPath_declaration(c *Path_declarationContext)

	// ExitSimple_path_declaration is called when exiting the simple_path_declaration production.
	ExitSimple_path_declaration(c *Simple_path_declarationContext)

	// ExitParallel_path_description is called when exiting the parallel_path_description production.
	ExitParallel_path_description(c *Parallel_path_descriptionContext)

	// ExitFull_path_description is called when exiting the full_path_description production.
	ExitFull_path_description(c *Full_path_descriptionContext)

	// ExitList_of_path_inputs is called when exiting the list_of_path_inputs production.
	ExitList_of_path_inputs(c *List_of_path_inputsContext)

	// ExitList_of_path_outputs is called when exiting the list_of_path_outputs production.
	ExitList_of_path_outputs(c *List_of_path_outputsContext)

	// ExitSpecify_input_terminal_descriptor is called when exiting the specify_input_terminal_descriptor production.
	ExitSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// ExitSpecify_output_terminal_descriptor is called when exiting the specify_output_terminal_descriptor production.
	ExitSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// ExitInput_identifier is called when exiting the input_identifier production.
	ExitInput_identifier(c *Input_identifierContext)

	// ExitOutput_identifier is called when exiting the output_identifier production.
	ExitOutput_identifier(c *Output_identifierContext)

	// ExitPath_delay_value is called when exiting the path_delay_value production.
	ExitPath_delay_value(c *Path_delay_valueContext)

	// ExitList_of_path_delay_expressions is called when exiting the list_of_path_delay_expressions production.
	ExitList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// ExitT_path_delay_expression is called when exiting the t_path_delay_expression production.
	ExitT_path_delay_expression(c *T_path_delay_expressionContext)

	// ExitTrise_path_delay_expression is called when exiting the trise_path_delay_expression production.
	ExitTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// ExitTfall_path_delay_expression is called when exiting the tfall_path_delay_expression production.
	ExitTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// ExitTz_path_delay_expression is called when exiting the tz_path_delay_expression production.
	ExitTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// ExitT01_path_delay_expression is called when exiting the t01_path_delay_expression production.
	ExitT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// ExitT10_path_delay_expression is called when exiting the t10_path_delay_expression production.
	ExitT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// ExitT0z_path_delay_expression is called when exiting the t0z_path_delay_expression production.
	ExitT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// ExitTz1_path_delay_expression is called when exiting the tz1_path_delay_expression production.
	ExitTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// ExitT1z_path_delay_expression is called when exiting the t1z_path_delay_expression production.
	ExitT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// ExitTz0_path_delay_expression is called when exiting the tz0_path_delay_expression production.
	ExitTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// ExitT0x_path_delay_expression is called when exiting the t0x_path_delay_expression production.
	ExitT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// ExitTx1_path_delay_expression is called when exiting the tx1_path_delay_expression production.
	ExitTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// ExitT1x_path_delay_expression is called when exiting the t1x_path_delay_expression production.
	ExitT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// ExitTx0_path_delay_expression is called when exiting the tx0_path_delay_expression production.
	ExitTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// ExitTxz_path_delay_expression is called when exiting the txz_path_delay_expression production.
	ExitTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// ExitTzx_path_delay_expression is called when exiting the tzx_path_delay_expression production.
	ExitTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// ExitPath_delay_expression is called when exiting the path_delay_expression production.
	ExitPath_delay_expression(c *Path_delay_expressionContext)

	// ExitEdge_sensitive_path_declaration is called when exiting the edge_sensitive_path_declaration production.
	ExitEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// ExitParallel_edge_sensitive_path_description is called when exiting the parallel_edge_sensitive_path_description production.
	ExitParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// ExitFull_edge_sensitive_path_description is called when exiting the full_edge_sensitive_path_description production.
	ExitFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// ExitData_source_expression is called when exiting the data_source_expression production.
	ExitData_source_expression(c *Data_source_expressionContext)

	// ExitEdge_identifier is called when exiting the edge_identifier production.
	ExitEdge_identifier(c *Edge_identifierContext)

	// ExitState_dependent_path_declaration is called when exiting the state_dependent_path_declaration production.
	ExitState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// ExitPolarity_operator is called when exiting the polarity_operator production.
	ExitPolarity_operator(c *Polarity_operatorContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitConstant_concatenation is called when exiting the constant_concatenation production.
	ExitConstant_concatenation(c *Constant_concatenationContext)

	// ExitConstant_multiple_concatenation is called when exiting the constant_multiple_concatenation production.
	ExitConstant_multiple_concatenation(c *Constant_multiple_concatenationContext)

	// ExitModule_path_concatenation is called when exiting the module_path_concatenation production.
	ExitModule_path_concatenation(c *Module_path_concatenationContext)

	// ExitModule_path_multiple_concatenation is called when exiting the module_path_multiple_concatenation production.
	ExitModule_path_multiple_concatenation(c *Module_path_multiple_concatenationContext)

	// ExitMultiple_concatenation is called when exiting the multiple_concatenation production.
	ExitMultiple_concatenation(c *Multiple_concatenationContext)

	// ExitStreaming_concatenation is called when exiting the streaming_concatenation production.
	ExitStreaming_concatenation(c *Streaming_concatenationContext)

	// ExitStream_operator is called when exiting the stream_operator production.
	ExitStream_operator(c *Stream_operatorContext)

	// ExitSlice_size is called when exiting the slice_size production.
	ExitSlice_size(c *Slice_sizeContext)

	// ExitStream_concatenation is called when exiting the stream_concatenation production.
	ExitStream_concatenation(c *Stream_concatenationContext)

	// ExitStream_expression is called when exiting the stream_expression production.
	ExitStream_expression(c *Stream_expressionContext)

	// ExitArray_range_expression is called when exiting the array_range_expression production.
	ExitArray_range_expression(c *Array_range_expressionContext)

	// ExitEmpty_unpacked_array_concatenation is called when exiting the empty_unpacked_array_concatenation production.
	ExitEmpty_unpacked_array_concatenation(c *Empty_unpacked_array_concatenationContext)

	// ExitConstant_function_call is called when exiting the constant_function_call production.
	ExitConstant_function_call(c *Constant_function_callContext)

	// ExitTf_call is called when exiting the tf_call production.
	ExitTf_call(c *Tf_callContext)

	// ExitSystem_tf_call is called when exiting the system_tf_call production.
	ExitSystem_tf_call(c *System_tf_callContext)

	// ExitSubroutine_call is called when exiting the subroutine_call production.
	ExitSubroutine_call(c *Subroutine_callContext)

	// ExitFunction_subroutine_call is called when exiting the function_subroutine_call production.
	ExitFunction_subroutine_call(c *Function_subroutine_callContext)

	// ExitList_of_arguments is called when exiting the list_of_arguments production.
	ExitList_of_arguments(c *List_of_argumentsContext)

	// ExitMethod_call_body is called when exiting the method_call_body production.
	ExitMethod_call_body(c *Method_call_bodyContext)

	// ExitBuilt_in_method_call is called when exiting the built_in_method_call production.
	ExitBuilt_in_method_call(c *Built_in_method_callContext)

	// ExitArray_manipulation_call is called when exiting the array_manipulation_call production.
	ExitArray_manipulation_call(c *Array_manipulation_callContext)

	// ExitRandomize_call is called when exiting the randomize_call production.
	ExitRandomize_call(c *Randomize_callContext)

	// ExitArray_method_name is called when exiting the array_method_name production.
	ExitArray_method_name(c *Array_method_nameContext)

	// ExitInc_or_dec_expression is called when exiting the inc_or_dec_expression production.
	ExitInc_or_dec_expression(c *Inc_or_dec_expressionContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitConstant_mintypmax_expression is called when exiting the constant_mintypmax_expression production.
	ExitConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// ExitConstant_param_expression is called when exiting the constant_param_expression production.
	ExitConstant_param_expression(c *Constant_param_expressionContext)

	// ExitParam_expression is called when exiting the param_expression production.
	ExitParam_expression(c *Param_expressionContext)

	// ExitConstant_range_expression is called when exiting the constant_range_expression production.
	ExitConstant_range_expression(c *Constant_range_expressionContext)

	// ExitConstant_part_select_range is called when exiting the constant_part_select_range production.
	ExitConstant_part_select_range(c *Constant_part_select_rangeContext)

	// ExitConstant_range is called when exiting the constant_range production.
	ExitConstant_range(c *Constant_rangeContext)

	// ExitConstant_indexed_range is called when exiting the constant_indexed_range production.
	ExitConstant_indexed_range(c *Constant_indexed_rangeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitValue_range is called when exiting the value_range production.
	ExitValue_range(c *Value_rangeContext)

	// ExitMintypmax_expression is called when exiting the mintypmax_expression production.
	ExitMintypmax_expression(c *Mintypmax_expressionContext)

	// ExitModule_path_expression is called when exiting the module_path_expression production.
	ExitModule_path_expression(c *Module_path_expressionContext)

	// ExitModule_path_mintypmax_expression is called when exiting the module_path_mintypmax_expression production.
	ExitModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// ExitPart_select_range is called when exiting the part_select_range production.
	ExitPart_select_range(c *Part_select_rangeContext)

	// ExitIndexed_range is called when exiting the indexed_range production.
	ExitIndexed_range(c *Indexed_rangeContext)

	// ExitGenvar_expression is called when exiting the genvar_expression production.
	ExitGenvar_expression(c *Genvar_expressionContext)

	// ExitConstant_primary is called when exiting the constant_primary production.
	ExitConstant_primary(c *Constant_primaryContext)

	// ExitModule_path_primary is called when exiting the module_path_primary production.
	ExitModule_path_primary(c *Module_path_primaryContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitClass_qualifier is called when exiting the class_qualifier production.
	ExitClass_qualifier(c *Class_qualifierContext)

	// ExitRange_expression is called when exiting the range_expression production.
	ExitRange_expression(c *Range_expressionContext)

	// ExitPrimary_literal is called when exiting the primary_literal production.
	ExitPrimary_literal(c *Primary_literalContext)

	// ExitImplicit_class_handle is called when exiting the implicit_class_handle production.
	ExitImplicit_class_handle(c *Implicit_class_handleContext)

	// ExitBit_select is called when exiting the bit_select production.
	ExitBit_select(c *Bit_selectContext)

	// ExitSelect_ is called when exiting the select_ production.
	ExitSelect_(c *Select_Context)

	// ExitNonrange_select is called when exiting the nonrange_select production.
	ExitNonrange_select(c *Nonrange_selectContext)

	// ExitConstant_bit_select is called when exiting the constant_bit_select production.
	ExitConstant_bit_select(c *Constant_bit_selectContext)

	// ExitConstant_select is called when exiting the constant_select production.
	ExitConstant_select(c *Constant_selectContext)

	// ExitConstant_let_expression is called when exiting the constant_let_expression production.
	ExitConstant_let_expression(c *Constant_let_expressionContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitNet_lvalue is called when exiting the net_lvalue production.
	ExitNet_lvalue(c *Net_lvalueContext)

	// ExitVariable_lvalue is called when exiting the variable_lvalue production.
	ExitVariable_lvalue(c *Variable_lvalueContext)

	// ExitNonrange_variable_lvalue is called when exiting the nonrange_variable_lvalue production.
	ExitNonrange_variable_lvalue(c *Nonrange_variable_lvalueContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitInc_or_dec_operator is called when exiting the inc_or_dec_operator production.
	ExitInc_or_dec_operator(c *Inc_or_dec_operatorContext)

	// ExitUnary_module_path_operator is called when exiting the unary_module_path_operator production.
	ExitUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// ExitBinary_module_path_operator is called when exiting the binary_module_path_operator production.
	ExitBinary_module_path_operator(c *Binary_module_path_operatorContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitAttribute_instance is called when exiting the attribute_instance production.
	ExitAttribute_instance(c *Attribute_instanceContext)

	// ExitAttr_spec is called when exiting the attr_spec production.
	ExitAttr_spec(c *Attr_specContext)

	// ExitAttr_name is called when exiting the attr_name production.
	ExitAttr_name(c *Attr_nameContext)

	// ExitArray_identifier is called when exiting the array_identifier production.
	ExitArray_identifier(c *Array_identifierContext)

	// ExitBlock_identifier is called when exiting the block_identifier production.
	ExitBlock_identifier(c *Block_identifierContext)

	// ExitBin_identifier is called when exiting the bin_identifier production.
	ExitBin_identifier(c *Bin_identifierContext)

	// ExitC_identifier is called when exiting the c_identifier production.
	ExitC_identifier(c *C_identifierContext)

	// ExitCell_identifier is called when exiting the cell_identifier production.
	ExitCell_identifier(c *Cell_identifierContext)

	// ExitChecker_identifier is called when exiting the checker_identifier production.
	ExitChecker_identifier(c *Checker_identifierContext)

	// ExitClass_identifier is called when exiting the class_identifier production.
	ExitClass_identifier(c *Class_identifierContext)

	// ExitClass_variable_identifier is called when exiting the class_variable_identifier production.
	ExitClass_variable_identifier(c *Class_variable_identifierContext)

	// ExitClocking_identifier is called when exiting the clocking_identifier production.
	ExitClocking_identifier(c *Clocking_identifierContext)

	// ExitConfig_identifier is called when exiting the config_identifier production.
	ExitConfig_identifier(c *Config_identifierContext)

	// ExitConst_identifier is called when exiting the const_identifier production.
	ExitConst_identifier(c *Const_identifierContext)

	// ExitConstraint_identifier is called when exiting the constraint_identifier production.
	ExitConstraint_identifier(c *Constraint_identifierContext)

	// ExitCovergroup_identifier is called when exiting the covergroup_identifier production.
	ExitCovergroup_identifier(c *Covergroup_identifierContext)

	// ExitCovergroup_variable_identifier is called when exiting the covergroup_variable_identifier production.
	ExitCovergroup_variable_identifier(c *Covergroup_variable_identifierContext)

	// ExitCover_point_identifier is called when exiting the cover_point_identifier production.
	ExitCover_point_identifier(c *Cover_point_identifierContext)

	// ExitCross_identifier is called when exiting the cross_identifier production.
	ExitCross_identifier(c *Cross_identifierContext)

	// ExitDynamic_array_variable_identifier is called when exiting the dynamic_array_variable_identifier production.
	ExitDynamic_array_variable_identifier(c *Dynamic_array_variable_identifierContext)

	// ExitEnum_identifier is called when exiting the enum_identifier production.
	ExitEnum_identifier(c *Enum_identifierContext)

	// ExitFormal_port_identifier is called when exiting the formal_port_identifier production.
	ExitFormal_port_identifier(c *Formal_port_identifierContext)

	// ExitFunction_identifier is called when exiting the function_identifier production.
	ExitFunction_identifier(c *Function_identifierContext)

	// ExitGenerate_block_identifier is called when exiting the generate_block_identifier production.
	ExitGenerate_block_identifier(c *Generate_block_identifierContext)

	// ExitGenvar_identifier is called when exiting the genvar_identifier production.
	ExitGenvar_identifier(c *Genvar_identifierContext)

	// ExitHierarchical_array_identifier is called when exiting the hierarchical_array_identifier production.
	ExitHierarchical_array_identifier(c *Hierarchical_array_identifierContext)

	// ExitHierarchical_block_identifier is called when exiting the hierarchical_block_identifier production.
	ExitHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// ExitHierarchical_event_identifier is called when exiting the hierarchical_event_identifier production.
	ExitHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// ExitHierarchical_identifier is called when exiting the hierarchical_identifier production.
	ExitHierarchical_identifier(c *Hierarchical_identifierContext)

	// ExitHierarchical_net_identifier is called when exiting the hierarchical_net_identifier production.
	ExitHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// ExitHierarchical_parameter_identifier is called when exiting the hierarchical_parameter_identifier production.
	ExitHierarchical_parameter_identifier(c *Hierarchical_parameter_identifierContext)

	// ExitHierarchical_property_identifier is called when exiting the hierarchical_property_identifier production.
	ExitHierarchical_property_identifier(c *Hierarchical_property_identifierContext)

	// ExitHierarchical_sequence_identifier is called when exiting the hierarchical_sequence_identifier production.
	ExitHierarchical_sequence_identifier(c *Hierarchical_sequence_identifierContext)

	// ExitHierarchical_task_identifier is called when exiting the hierarchical_task_identifier production.
	ExitHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// ExitHierarchical_tf_identifier is called when exiting the hierarchical_tf_identifier production.
	ExitHierarchical_tf_identifier(c *Hierarchical_tf_identifierContext)

	// ExitHierarchical_variable_identifier is called when exiting the hierarchical_variable_identifier production.
	ExitHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIndex_variable_identifier is called when exiting the index_variable_identifier production.
	ExitIndex_variable_identifier(c *Index_variable_identifierContext)

	// ExitInterface_identifier is called when exiting the interface_identifier production.
	ExitInterface_identifier(c *Interface_identifierContext)

	// ExitInterface_instance_identifier is called when exiting the interface_instance_identifier production.
	ExitInterface_instance_identifier(c *Interface_instance_identifierContext)

	// ExitInout_port_identifier is called when exiting the inout_port_identifier production.
	ExitInout_port_identifier(c *Inout_port_identifierContext)

	// ExitInput_port_identifier is called when exiting the input_port_identifier production.
	ExitInput_port_identifier(c *Input_port_identifierContext)

	// ExitInstance_identifier is called when exiting the instance_identifier production.
	ExitInstance_identifier(c *Instance_identifierContext)

	// ExitLibrary_identifier is called when exiting the library_identifier production.
	ExitLibrary_identifier(c *Library_identifierContext)

	// ExitMember_identifier is called when exiting the member_identifier production.
	ExitMember_identifier(c *Member_identifierContext)

	// ExitMethod_identifier is called when exiting the method_identifier production.
	ExitMethod_identifier(c *Method_identifierContext)

	// ExitModport_identifier is called when exiting the modport_identifier production.
	ExitModport_identifier(c *Modport_identifierContext)

	// ExitModule_identifier is called when exiting the module_identifier production.
	ExitModule_identifier(c *Module_identifierContext)

	// ExitNet_identifier is called when exiting the net_identifier production.
	ExitNet_identifier(c *Net_identifierContext)

	// ExitNet_type_identifier is called when exiting the net_type_identifier production.
	ExitNet_type_identifier(c *Net_type_identifierContext)

	// ExitOutput_port_identifier is called when exiting the output_port_identifier production.
	ExitOutput_port_identifier(c *Output_port_identifierContext)

	// ExitPackage_identifier is called when exiting the package_identifier production.
	ExitPackage_identifier(c *Package_identifierContext)

	// ExitPackage_scope is called when exiting the package_scope production.
	ExitPackage_scope(c *Package_scopeContext)

	// ExitParameter_identifier is called when exiting the parameter_identifier production.
	ExitParameter_identifier(c *Parameter_identifierContext)

	// ExitPort_identifier is called when exiting the port_identifier production.
	ExitPort_identifier(c *Port_identifierContext)

	// ExitProduction_identifier is called when exiting the production_identifier production.
	ExitProduction_identifier(c *Production_identifierContext)

	// ExitProgram_identifier is called when exiting the program_identifier production.
	ExitProgram_identifier(c *Program_identifierContext)

	// ExitProperty_identifier is called when exiting the property_identifier production.
	ExitProperty_identifier(c *Property_identifierContext)

	// ExitPs_class_identifier is called when exiting the ps_class_identifier production.
	ExitPs_class_identifier(c *Ps_class_identifierContext)

	// ExitPs_covergroup_identifier is called when exiting the ps_covergroup_identifier production.
	ExitPs_covergroup_identifier(c *Ps_covergroup_identifierContext)

	// ExitPs_checker_identifier is called when exiting the ps_checker_identifier production.
	ExitPs_checker_identifier(c *Ps_checker_identifierContext)

	// ExitPs_identifier is called when exiting the ps_identifier production.
	ExitPs_identifier(c *Ps_identifierContext)

	// ExitPs_or_hierarchical_array_identifier is called when exiting the ps_or_hierarchical_array_identifier production.
	ExitPs_or_hierarchical_array_identifier(c *Ps_or_hierarchical_array_identifierContext)

	// ExitPs_or_hierarchical_net_identifier is called when exiting the ps_or_hierarchical_net_identifier production.
	ExitPs_or_hierarchical_net_identifier(c *Ps_or_hierarchical_net_identifierContext)

	// ExitPs_or_hierarchical_property_identifier is called when exiting the ps_or_hierarchical_property_identifier production.
	ExitPs_or_hierarchical_property_identifier(c *Ps_or_hierarchical_property_identifierContext)

	// ExitPs_or_hierarchical_sequence_identifier is called when exiting the ps_or_hierarchical_sequence_identifier production.
	ExitPs_or_hierarchical_sequence_identifier(c *Ps_or_hierarchical_sequence_identifierContext)

	// ExitPs_or_hierarchical_tf_identifier is called when exiting the ps_or_hierarchical_tf_identifier production.
	ExitPs_or_hierarchical_tf_identifier(c *Ps_or_hierarchical_tf_identifierContext)

	// ExitPs_parameter_identifier is called when exiting the ps_parameter_identifier production.
	ExitPs_parameter_identifier(c *Ps_parameter_identifierContext)

	// ExitPs_type_identifier is called when exiting the ps_type_identifier production.
	ExitPs_type_identifier(c *Ps_type_identifierContext)

	// ExitSequence_identifier is called when exiting the sequence_identifier production.
	ExitSequence_identifier(c *Sequence_identifierContext)

	// ExitSignal_identifier is called when exiting the signal_identifier production.
	ExitSignal_identifier(c *Signal_identifierContext)

	// ExitSpecparam_identifier is called when exiting the specparam_identifier production.
	ExitSpecparam_identifier(c *Specparam_identifierContext)

	// ExitTask_identifier is called when exiting the task_identifier production.
	ExitTask_identifier(c *Task_identifierContext)

	// ExitTf_identifier is called when exiting the tf_identifier production.
	ExitTf_identifier(c *Tf_identifierContext)

	// ExitTopmodule_identifier is called when exiting the topmodule_identifier production.
	ExitTopmodule_identifier(c *Topmodule_identifierContext)

	// ExitType_identifier is called when exiting the type_identifier production.
	ExitType_identifier(c *Type_identifierContext)

	// ExitVariable_identifier is called when exiting the variable_identifier production.
	ExitVariable_identifier(c *Variable_identifierContext)
}
