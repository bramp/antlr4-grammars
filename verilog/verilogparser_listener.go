// Code generated from VerilogParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package verilog // VerilogParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// VerilogParserListener is a complete listener for a parse tree produced by VerilogParser.
type VerilogParserListener interface {
	antlr.ParseTreeListener

	// EnterDisplay_tasks is called when entering the display_tasks production.
	EnterDisplay_tasks(c *Display_tasksContext)

	// EnterDisplay_task_name is called when entering the display_task_name production.
	EnterDisplay_task_name(c *Display_task_nameContext)

	// EnterList_of_arguments is called when entering the list_of_arguments production.
	EnterList_of_arguments(c *List_of_argumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterStrobe_tasks is called when entering the strobe_tasks production.
	EnterStrobe_tasks(c *Strobe_tasksContext)

	// EnterStrobe_task_name is called when entering the strobe_task_name production.
	EnterStrobe_task_name(c *Strobe_task_nameContext)

	// EnterMonitor_tasks is called when entering the monitor_tasks production.
	EnterMonitor_tasks(c *Monitor_tasksContext)

	// EnterMonitor_task_name is called when entering the monitor_task_name production.
	EnterMonitor_task_name(c *Monitor_task_nameContext)

	// EnterFile_open_function is called when entering the file_open_function production.
	EnterFile_open_function(c *File_open_functionContext)

	// EnterFile_close_task is called when entering the file_close_task production.
	EnterFile_close_task(c *File_close_taskContext)

	// EnterMulti_channel_descriptor is called when entering the multi_channel_descriptor production.
	EnterMulti_channel_descriptor(c *Multi_channel_descriptorContext)

	// EnterFd is called when entering the fd production.
	EnterFd(c *FdContext)

	// EnterFile_name is called when entering the file_name production.
	EnterFile_name(c *File_nameContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFile_output_tasks is called when entering the file_output_tasks production.
	EnterFile_output_tasks(c *File_output_tasksContext)

	// EnterFile_output_task_name is called when entering the file_output_task_name production.
	EnterFile_output_task_name(c *File_output_task_nameContext)

	// EnterLoad_memory_tasks is called when entering the load_memory_tasks production.
	EnterLoad_memory_tasks(c *Load_memory_tasksContext)

	// EnterMemory_name is called when entering the memory_name production.
	EnterMemory_name(c *Memory_nameContext)

	// EnterStart_addr is called when entering the start_addr production.
	EnterStart_addr(c *Start_addrContext)

	// EnterFinish_addr is called when entering the finish_addr production.
	EnterFinish_addr(c *Finish_addrContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterFinish_task is called when entering the finish_task production.
	EnterFinish_task(c *Finish_taskContext)

	// EnterFinish_number is called when entering the finish_number production.
	EnterFinish_number(c *Finish_numberContext)

	// EnterStop_task is called when entering the stop_task production.
	EnterStop_task(c *Stop_taskContext)

	// EnterTime_function is called when entering the time_function production.
	EnterTime_function(c *Time_functionContext)

	// EnterStime_function is called when entering the stime_function production.
	EnterStime_function(c *Stime_functionContext)

	// EnterRealtime_function is called when entering the realtime_function production.
	EnterRealtime_function(c *Realtime_functionContext)

	// EnterConversion_functions is called when entering the conversion_functions production.
	EnterConversion_functions(c *Conversion_functionsContext)

	// EnterConversion_function_name is called when entering the conversion_function_name production.
	EnterConversion_function_name(c *Conversion_function_nameContext)

	// EnterConstant_argument is called when entering the constant_argument production.
	EnterConstant_argument(c *Constant_argumentContext)

	// EnterRandom_function is called when entering the random_function production.
	EnterRandom_function(c *Random_functionContext)

	// EnterSeed is called when entering the seed production.
	EnterSeed(c *SeedContext)

	// EnterDist_functions is called when entering the dist_functions production.
	EnterDist_functions(c *Dist_functionsContext)

	// EnterStart_ is called when entering the start_ production.
	EnterStart_(c *Start_Context)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterMean is called when entering the mean production.
	EnterMean(c *MeanContext)

	// EnterStandard_deviation is called when entering the standard_deviation production.
	EnterStandard_deviation(c *Standard_deviationContext)

	// EnterDegree_of_freedom is called when entering the degree_of_freedom production.
	EnterDegree_of_freedom(c *Degree_of_freedomContext)

	// EnterK_stage is called when entering the k_stage production.
	EnterK_stage(c *K_stageContext)

	// EnterMath_functions is called when entering the math_functions production.
	EnterMath_functions(c *Math_functionsContext)

	// EnterInteger_math_functions is called when entering the integer_math_functions production.
	EnterInteger_math_functions(c *Integer_math_functionsContext)

	// EnterReal_math_functions is called when entering the real_math_functions production.
	EnterReal_math_functions(c *Real_math_functionsContext)

	// EnterSingle_argument_real_math_function_name is called when entering the single_argument_real_math_function_name production.
	EnterSingle_argument_real_math_function_name(c *Single_argument_real_math_function_nameContext)

	// EnterDouble_argument_real_math_function_name is called when entering the double_argument_real_math_function_name production.
	EnterDouble_argument_real_math_function_name(c *Double_argument_real_math_function_nameContext)

	// EnterDumpfile_task is called when entering the dumpfile_task production.
	EnterDumpfile_task(c *Dumpfile_taskContext)

	// EnterDumpvars_task is called when entering the dumpvars_task production.
	EnterDumpvars_task(c *Dumpvars_taskContext)

	// EnterList_of_modules_or_variables is called when entering the list_of_modules_or_variables production.
	EnterList_of_modules_or_variables(c *List_of_modules_or_variablesContext)

	// EnterModule_or_variable is called when entering the module_or_variable production.
	EnterModule_or_variable(c *Module_or_variableContext)

	// EnterLevels is called when entering the levels production.
	EnterLevels(c *LevelsContext)

	// EnterDumpoff_task is called when entering the dumpoff_task production.
	EnterDumpoff_task(c *Dumpoff_taskContext)

	// EnterDumpon_task is called when entering the dumpon_task production.
	EnterDumpon_task(c *Dumpon_taskContext)

	// EnterDumpall_task is called when entering the dumpall_task production.
	EnterDumpall_task(c *Dumpall_taskContext)

	// EnterDumplimit_task is called when entering the dumplimit_task production.
	EnterDumplimit_task(c *Dumplimit_taskContext)

	// EnterFile_size is called when entering the file_size production.
	EnterFile_size(c *File_sizeContext)

	// EnterDumpflush_task is called when entering the dumpflush_task production.
	EnterDumpflush_task(c *Dumpflush_taskContext)

	// EnterDumpports_task is called when entering the dumpports_task production.
	EnterDumpports_task(c *Dumpports_taskContext)

	// EnterScope_list is called when entering the scope_list production.
	EnterScope_list(c *Scope_listContext)

	// EnterFile_pathname is called when entering the file_pathname production.
	EnterFile_pathname(c *File_pathnameContext)

	// EnterDumpportsoff_task is called when entering the dumpportsoff_task production.
	EnterDumpportsoff_task(c *Dumpportsoff_taskContext)

	// EnterDumpportson_task is called when entering the dumpportson_task production.
	EnterDumpportson_task(c *Dumpportson_taskContext)

	// EnterDumpportsall_task is called when entering the dumpportsall_task production.
	EnterDumpportsall_task(c *Dumpportsall_taskContext)

	// EnterDumpportslimit_task is called when entering the dumpportslimit_task production.
	EnterDumpportslimit_task(c *Dumpportslimit_taskContext)

	// EnterDumpportsflush_task is called when entering the dumpportsflush_task production.
	EnterDumpportsflush_task(c *Dumpportsflush_taskContext)

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

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterModule_keyword is called when entering the module_keyword production.
	EnterModule_keyword(c *Module_keywordContext)

	// EnterModule_parameter_port_list is called when entering the module_parameter_port_list production.
	EnterModule_parameter_port_list(c *Module_parameter_port_listContext)

	// EnterList_of_ports is called when entering the list_of_ports production.
	EnterList_of_ports(c *List_of_portsContext)

	// EnterList_of_port_declarations is called when entering the list_of_port_declarations production.
	EnterList_of_port_declarations(c *List_of_port_declarationsContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterPort_expression is called when entering the port_expression production.
	EnterPort_expression(c *Port_expressionContext)

	// EnterPort_reference is called when entering the port_reference production.
	EnterPort_reference(c *Port_referenceContext)

	// EnterPort_declaration is called when entering the port_declaration production.
	EnterPort_declaration(c *Port_declarationContext)

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

	// EnterLocal_parameter_declaration is called when entering the local_parameter_declaration production.
	EnterLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterSpecparam_declaration is called when entering the specparam_declaration production.
	EnterSpecparam_declaration(c *Specparam_declarationContext)

	// EnterParameter_type is called when entering the parameter_type production.
	EnterParameter_type(c *Parameter_typeContext)

	// EnterInout_declaration is called when entering the inout_declaration production.
	EnterInout_declaration(c *Inout_declarationContext)

	// EnterInput_declaration is called when entering the input_declaration production.
	EnterInput_declaration(c *Input_declarationContext)

	// EnterOutput_declaration is called when entering the output_declaration production.
	EnterOutput_declaration(c *Output_declarationContext)

	// EnterEvent_declaration is called when entering the event_declaration production.
	EnterEvent_declaration(c *Event_declarationContext)

	// EnterInteger_declaration is called when entering the integer_declaration production.
	EnterInteger_declaration(c *Integer_declarationContext)

	// EnterNet_declaration is called when entering the net_declaration production.
	EnterNet_declaration(c *Net_declarationContext)

	// EnterReal_declaration is called when entering the real_declaration production.
	EnterReal_declaration(c *Real_declarationContext)

	// EnterRealtime_declaration is called when entering the realtime_declaration production.
	EnterRealtime_declaration(c *Realtime_declarationContext)

	// EnterReg_declaration is called when entering the reg_declaration production.
	EnterReg_declaration(c *Reg_declarationContext)

	// EnterTime_declaration is called when entering the time_declaration production.
	EnterTime_declaration(c *Time_declarationContext)

	// EnterNet_type is called when entering the net_type production.
	EnterNet_type(c *Net_typeContext)

	// EnterOutput_variable_type is called when entering the output_variable_type production.
	EnterOutput_variable_type(c *Output_variable_typeContext)

	// EnterReal_type is called when entering the real_type production.
	EnterReal_type(c *Real_typeContext)

	// EnterVariable_type is called when entering the variable_type production.
	EnterVariable_type(c *Variable_typeContext)

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

	// EnterList_of_event_identifiers is called when entering the list_of_event_identifiers production.
	EnterList_of_event_identifiers(c *List_of_event_identifiersContext)

	// EnterList_of_net_decl_assignments is called when entering the list_of_net_decl_assignments production.
	EnterList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// EnterList_of_net_identifiers is called when entering the list_of_net_identifiers production.
	EnterList_of_net_identifiers(c *List_of_net_identifiersContext)

	// EnterList_of_param_assignments is called when entering the list_of_param_assignments production.
	EnterList_of_param_assignments(c *List_of_param_assignmentsContext)

	// EnterList_of_port_identifiers is called when entering the list_of_port_identifiers production.
	EnterList_of_port_identifiers(c *List_of_port_identifiersContext)

	// EnterList_of_real_identifiers is called when entering the list_of_real_identifiers production.
	EnterList_of_real_identifiers(c *List_of_real_identifiersContext)

	// EnterList_of_specparam_assignments is called when entering the list_of_specparam_assignments production.
	EnterList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

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

	// EnterPulse_control_specparam is called when entering the pulse_control_specparam production.
	EnterPulse_control_specparam(c *Pulse_control_specparamContext)

	// EnterError_limit_value is called when entering the error_limit_value production.
	EnterError_limit_value(c *Error_limit_valueContext)

	// EnterReject_limit_value is called when entering the reject_limit_value production.
	EnterReject_limit_value(c *Reject_limit_valueContext)

	// EnterLimit_value is called when entering the limit_value production.
	EnterLimit_value(c *Limit_valueContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterRange_ is called when entering the range_ production.
	EnterRange_(c *Range_Context)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_item_declaration is called when entering the function_item_declaration production.
	EnterFunction_item_declaration(c *Function_item_declarationContext)

	// EnterFunction_port_list is called when entering the function_port_list production.
	EnterFunction_port_list(c *Function_port_listContext)

	// EnterFunction_range_or_type is called when entering the function_range_or_type production.
	EnterFunction_range_or_type(c *Function_range_or_typeContext)

	// EnterTask_declaration is called when entering the task_declaration production.
	EnterTask_declaration(c *Task_declarationContext)

	// EnterTask_item_declaration is called when entering the task_item_declaration production.
	EnterTask_item_declaration(c *Task_item_declarationContext)

	// EnterTask_port_list is called when entering the task_port_list production.
	EnterTask_port_list(c *Task_port_listContext)

	// EnterTask_port_item is called when entering the task_port_item production.
	EnterTask_port_item(c *Task_port_itemContext)

	// EnterTf_input_declaration is called when entering the tf_input_declaration production.
	EnterTf_input_declaration(c *Tf_input_declarationContext)

	// EnterTf_output_declaration is called when entering the tf_output_declaration production.
	EnterTf_output_declaration(c *Tf_output_declarationContext)

	// EnterTf_inout_declaration is called when entering the tf_inout_declaration production.
	EnterTf_inout_declaration(c *Tf_inout_declarationContext)

	// EnterTask_port_type is called when entering the task_port_type production.
	EnterTask_port_type(c *Task_port_typeContext)

	// EnterBlock_item_declaration is called when entering the block_item_declaration production.
	EnterBlock_item_declaration(c *Block_item_declarationContext)

	// EnterList_of_block_variable_identifiers is called when entering the list_of_block_variable_identifiers production.
	EnterList_of_block_variable_identifiers(c *List_of_block_variable_identifiersContext)

	// EnterList_of_block_real_identifiers is called when entering the list_of_block_real_identifiers production.
	EnterList_of_block_real_identifiers(c *List_of_block_real_identifiersContext)

	// EnterBlock_variable_type is called when entering the block_variable_type production.
	EnterBlock_variable_type(c *Block_variable_typeContext)

	// EnterBlock_real_type is called when entering the block_real_type production.
	EnterBlock_real_type(c *Block_real_typeContext)

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

	// EnterName_of_gate_instance is called when entering the name_of_gate_instance production.
	EnterName_of_gate_instance(c *Name_of_gate_instanceContext)

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

	// EnterModule_instance is called when entering the module_instance production.
	EnterModule_instance(c *Module_instanceContext)

	// EnterName_of_module_instance is called when entering the name_of_module_instance production.
	EnterName_of_module_instance(c *Name_of_module_instanceContext)

	// EnterList_of_port_connections is called when entering the list_of_port_connections production.
	EnterList_of_port_connections(c *List_of_port_connectionsContext)

	// EnterOrdered_port_connection is called when entering the ordered_port_connection production.
	EnterOrdered_port_connection(c *Ordered_port_connectionContext)

	// EnterNamed_port_connection is called when entering the named_port_connection production.
	EnterNamed_port_connection(c *Named_port_connectionContext)

	// EnterGenerate_region is called when entering the generate_region production.
	EnterGenerate_region(c *Generate_regionContext)

	// EnterGenvar_declaration is called when entering the genvar_declaration production.
	EnterGenvar_declaration(c *Genvar_declarationContext)

	// EnterList_of_genvar_identifiers is called when entering the list_of_genvar_identifiers production.
	EnterList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// EnterLoop_generate_construct is called when entering the loop_generate_construct production.
	EnterLoop_generate_construct(c *Loop_generate_constructContext)

	// EnterGenvar_initialization is called when entering the genvar_initialization production.
	EnterGenvar_initialization(c *Genvar_initializationContext)

	// EnterGenvar_expression is called when entering the genvar_expression production.
	EnterGenvar_expression(c *Genvar_expressionContext)

	// EnterGenvar_iteration is called when entering the genvar_iteration production.
	EnterGenvar_iteration(c *Genvar_iterationContext)

	// EnterGenvar_primary is called when entering the genvar_primary production.
	EnterGenvar_primary(c *Genvar_primaryContext)

	// EnterConditional_generate_construct is called when entering the conditional_generate_construct production.
	EnterConditional_generate_construct(c *Conditional_generate_constructContext)

	// EnterIf_generate_construct is called when entering the if_generate_construct production.
	EnterIf_generate_construct(c *If_generate_constructContext)

	// EnterCase_generate_construct is called when entering the case_generate_construct production.
	EnterCase_generate_construct(c *Case_generate_constructContext)

	// EnterGenerate_block is called when entering the generate_block production.
	EnterGenerate_block(c *Generate_blockContext)

	// EnterGenerate_block_or_null is called when entering the generate_block_or_null production.
	EnterGenerate_block_or_null(c *Generate_block_or_nullContext)

	// EnterContinuous_assign is called when entering the continuous_assign production.
	EnterContinuous_assign(c *Continuous_assignContext)

	// EnterList_of_net_assignments is called when entering the list_of_net_assignments production.
	EnterList_of_net_assignments(c *List_of_net_assignmentsContext)

	// EnterNet_assignment is called when entering the net_assignment production.
	EnterNet_assignment(c *Net_assignmentContext)

	// EnterInitial_construct is called when entering the initial_construct production.
	EnterInitial_construct(c *Initial_constructContext)

	// EnterAlways_construct is called when entering the always_construct production.
	EnterAlways_construct(c *Always_constructContext)

	// EnterBlocking_assignment is called when entering the blocking_assignment production.
	EnterBlocking_assignment(c *Blocking_assignmentContext)

	// EnterNonblocking_assignment is called when entering the nonblocking_assignment production.
	EnterNonblocking_assignment(c *Nonblocking_assignmentContext)

	// EnterProcedural_continuous_assignments is called when entering the procedural_continuous_assignments production.
	EnterProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterPar_block is called when entering the par_block production.
	EnterPar_block(c *Par_blockContext)

	// EnterSeq_block is called when entering the seq_block production.
	EnterSeq_block(c *Seq_blockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatement_or_null is called when entering the statement_or_null production.
	EnterStatement_or_null(c *Statement_or_nullContext)

	// EnterFunction_statement is called when entering the function_statement production.
	EnterFunction_statement(c *Function_statementContext)

	// EnterDelay_control is called when entering the delay_control production.
	EnterDelay_control(c *Delay_controlContext)

	// EnterDelay_or_event_control is called when entering the delay_or_event_control production.
	EnterDelay_or_event_control(c *Delay_or_event_controlContext)

	// EnterDisable_statement is called when entering the disable_statement production.
	EnterDisable_statement(c *Disable_statementContext)

	// EnterEvent_control is called when entering the event_control production.
	EnterEvent_control(c *Event_controlContext)

	// EnterEvent_trigger is called when entering the event_trigger production.
	EnterEvent_trigger(c *Event_triggerContext)

	// EnterEvent_expression is called when entering the event_expression production.
	EnterEvent_expression(c *Event_expressionContext)

	// EnterEvent_primary is called when entering the event_primary production.
	EnterEvent_primary(c *Event_primaryContext)

	// EnterProcedural_timing_control is called when entering the procedural_timing_control production.
	EnterProcedural_timing_control(c *Procedural_timing_controlContext)

	// EnterProcedural_timing_control_statement is called when entering the procedural_timing_control_statement production.
	EnterProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterConditional_statement is called when entering the conditional_statement production.
	EnterConditional_statement(c *Conditional_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_item is called when entering the case_item production.
	EnterCase_item(c *Case_itemContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterSystem_task_enable is called when entering the system_task_enable production.
	EnterSystem_task_enable(c *System_task_enableContext)

	// EnterTask_enable is called when entering the task_enable production.
	EnterTask_enable(c *Task_enableContext)

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

	// EnterConstant_function_call is called when entering the constant_function_call production.
	EnterConstant_function_call(c *Constant_function_callContext)

	// EnterConstant_system_function_call is called when entering the constant_system_function_call production.
	EnterConstant_system_function_call(c *Constant_system_function_callContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterSystem_function_call is called when entering the system_function_call production.
	EnterSystem_function_call(c *System_function_callContext)

	// EnterBase_expression is called when entering the base_expression production.
	EnterBase_expression(c *Base_expressionContext)

	// EnterConstant_base_expression is called when entering the constant_base_expression production.
	EnterConstant_base_expression(c *Constant_base_expressionContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterConstant_mintypmax_expression is called when entering the constant_mintypmax_expression production.
	EnterConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// EnterConstant_range_expression is called when entering the constant_range_expression production.
	EnterConstant_range_expression(c *Constant_range_expressionContext)

	// EnterDimension_constant_expression is called when entering the dimension_constant_expression production.
	EnterDimension_constant_expression(c *Dimension_constant_expressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLsb_constant_expression is called when entering the lsb_constant_expression production.
	EnterLsb_constant_expression(c *Lsb_constant_expressionContext)

	// EnterMintypmax_expression is called when entering the mintypmax_expression production.
	EnterMintypmax_expression(c *Mintypmax_expressionContext)

	// EnterModule_path_expression is called when entering the module_path_expression production.
	EnterModule_path_expression(c *Module_path_expressionContext)

	// EnterModule_path_mintypmax_expression is called when entering the module_path_mintypmax_expression production.
	EnterModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// EnterMsb_constant_expression is called when entering the msb_constant_expression production.
	EnterMsb_constant_expression(c *Msb_constant_expressionContext)

	// EnterRange_expression is called when entering the range_expression production.
	EnterRange_expression(c *Range_expressionContext)

	// EnterWidth_constant_expression is called when entering the width_constant_expression production.
	EnterWidth_constant_expression(c *Width_constant_expressionContext)

	// EnterConstant_primary is called when entering the constant_primary production.
	EnterConstant_primary(c *Constant_primaryContext)

	// EnterModule_path_primary is called when entering the module_path_primary production.
	EnterModule_path_primary(c *Module_path_primaryContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterNet_lvalue is called when entering the net_lvalue production.
	EnterNet_lvalue(c *Net_lvalueContext)

	// EnterVariable_lvalue is called when entering the variable_lvalue production.
	EnterVariable_lvalue(c *Variable_lvalueContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

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

	// EnterBlock_identifier is called when entering the block_identifier production.
	EnterBlock_identifier(c *Block_identifierContext)

	// EnterCell_identifier is called when entering the cell_identifier production.
	EnterCell_identifier(c *Cell_identifierContext)

	// EnterConfig_identifier is called when entering the config_identifier production.
	EnterConfig_identifier(c *Config_identifierContext)

	// EnterEvent_identifier is called when entering the event_identifier production.
	EnterEvent_identifier(c *Event_identifierContext)

	// EnterFunction_identifier is called when entering the function_identifier production.
	EnterFunction_identifier(c *Function_identifierContext)

	// EnterGate_instance_identifier is called when entering the gate_instance_identifier production.
	EnterGate_instance_identifier(c *Gate_instance_identifierContext)

	// EnterGenerate_block_identifier is called when entering the generate_block_identifier production.
	EnterGenerate_block_identifier(c *Generate_block_identifierContext)

	// EnterGenvar_identifier is called when entering the genvar_identifier production.
	EnterGenvar_identifier(c *Genvar_identifierContext)

	// EnterHierarchical_block_identifier is called when entering the hierarchical_block_identifier production.
	EnterHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// EnterHierarchical_event_identifier is called when entering the hierarchical_event_identifier production.
	EnterHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// EnterHierarchical_function_identifier is called when entering the hierarchical_function_identifier production.
	EnterHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// EnterHierarchical_identifier is called when entering the hierarchical_identifier production.
	EnterHierarchical_identifier(c *Hierarchical_identifierContext)

	// EnterHierarchical_net_identifier is called when entering the hierarchical_net_identifier production.
	EnterHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// EnterHierarchical_parameter_identifier is called when entering the hierarchical_parameter_identifier production.
	EnterHierarchical_parameter_identifier(c *Hierarchical_parameter_identifierContext)

	// EnterHierarchical_variable_identifier is called when entering the hierarchical_variable_identifier production.
	EnterHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// EnterHierarchical_task_identifier is called when entering the hierarchical_task_identifier production.
	EnterHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterInout_port_identifier is called when entering the inout_port_identifier production.
	EnterInout_port_identifier(c *Inout_port_identifierContext)

	// EnterInput_port_identifier is called when entering the input_port_identifier production.
	EnterInput_port_identifier(c *Input_port_identifierContext)

	// EnterInstance_identifier is called when entering the instance_identifier production.
	EnterInstance_identifier(c *Instance_identifierContext)

	// EnterLibrary_identifier is called when entering the library_identifier production.
	EnterLibrary_identifier(c *Library_identifierContext)

	// EnterModule_identifier is called when entering the module_identifier production.
	EnterModule_identifier(c *Module_identifierContext)

	// EnterModule_instance_identifier is called when entering the module_instance_identifier production.
	EnterModule_instance_identifier(c *Module_instance_identifierContext)

	// EnterNet_identifier is called when entering the net_identifier production.
	EnterNet_identifier(c *Net_identifierContext)

	// EnterOutput_port_identifier is called when entering the output_port_identifier production.
	EnterOutput_port_identifier(c *Output_port_identifierContext)

	// EnterParameter_identifier is called when entering the parameter_identifier production.
	EnterParameter_identifier(c *Parameter_identifierContext)

	// EnterPort_identifier is called when entering the port_identifier production.
	EnterPort_identifier(c *Port_identifierContext)

	// EnterReal_identifier is called when entering the real_identifier production.
	EnterReal_identifier(c *Real_identifierContext)

	// EnterSpecparam_identifier is called when entering the specparam_identifier production.
	EnterSpecparam_identifier(c *Specparam_identifierContext)

	// EnterSystem_function_identifier is called when entering the system_function_identifier production.
	EnterSystem_function_identifier(c *System_function_identifierContext)

	// EnterSystem_task_identifier is called when entering the system_task_identifier production.
	EnterSystem_task_identifier(c *System_task_identifierContext)

	// EnterTask_identifier is called when entering the task_identifier production.
	EnterTask_identifier(c *Task_identifierContext)

	// EnterTerminal_identifier is called when entering the terminal_identifier production.
	EnterTerminal_identifier(c *Terminal_identifierContext)

	// EnterTopmodule_identifier is called when entering the topmodule_identifier production.
	EnterTopmodule_identifier(c *Topmodule_identifierContext)

	// EnterUdp_identifier is called when entering the udp_identifier production.
	EnterUdp_identifier(c *Udp_identifierContext)

	// EnterUdp_instance_identifier is called when entering the udp_instance_identifier production.
	EnterUdp_instance_identifier(c *Udp_instance_identifierContext)

	// EnterVariable_identifier is called when entering the variable_identifier production.
	EnterVariable_identifier(c *Variable_identifierContext)

	// ExitDisplay_tasks is called when exiting the display_tasks production.
	ExitDisplay_tasks(c *Display_tasksContext)

	// ExitDisplay_task_name is called when exiting the display_task_name production.
	ExitDisplay_task_name(c *Display_task_nameContext)

	// ExitList_of_arguments is called when exiting the list_of_arguments production.
	ExitList_of_arguments(c *List_of_argumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitStrobe_tasks is called when exiting the strobe_tasks production.
	ExitStrobe_tasks(c *Strobe_tasksContext)

	// ExitStrobe_task_name is called when exiting the strobe_task_name production.
	ExitStrobe_task_name(c *Strobe_task_nameContext)

	// ExitMonitor_tasks is called when exiting the monitor_tasks production.
	ExitMonitor_tasks(c *Monitor_tasksContext)

	// ExitMonitor_task_name is called when exiting the monitor_task_name production.
	ExitMonitor_task_name(c *Monitor_task_nameContext)

	// ExitFile_open_function is called when exiting the file_open_function production.
	ExitFile_open_function(c *File_open_functionContext)

	// ExitFile_close_task is called when exiting the file_close_task production.
	ExitFile_close_task(c *File_close_taskContext)

	// ExitMulti_channel_descriptor is called when exiting the multi_channel_descriptor production.
	ExitMulti_channel_descriptor(c *Multi_channel_descriptorContext)

	// ExitFd is called when exiting the fd production.
	ExitFd(c *FdContext)

	// ExitFile_name is called when exiting the file_name production.
	ExitFile_name(c *File_nameContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFile_output_tasks is called when exiting the file_output_tasks production.
	ExitFile_output_tasks(c *File_output_tasksContext)

	// ExitFile_output_task_name is called when exiting the file_output_task_name production.
	ExitFile_output_task_name(c *File_output_task_nameContext)

	// ExitLoad_memory_tasks is called when exiting the load_memory_tasks production.
	ExitLoad_memory_tasks(c *Load_memory_tasksContext)

	// ExitMemory_name is called when exiting the memory_name production.
	ExitMemory_name(c *Memory_nameContext)

	// ExitStart_addr is called when exiting the start_addr production.
	ExitStart_addr(c *Start_addrContext)

	// ExitFinish_addr is called when exiting the finish_addr production.
	ExitFinish_addr(c *Finish_addrContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitFinish_task is called when exiting the finish_task production.
	ExitFinish_task(c *Finish_taskContext)

	// ExitFinish_number is called when exiting the finish_number production.
	ExitFinish_number(c *Finish_numberContext)

	// ExitStop_task is called when exiting the stop_task production.
	ExitStop_task(c *Stop_taskContext)

	// ExitTime_function is called when exiting the time_function production.
	ExitTime_function(c *Time_functionContext)

	// ExitStime_function is called when exiting the stime_function production.
	ExitStime_function(c *Stime_functionContext)

	// ExitRealtime_function is called when exiting the realtime_function production.
	ExitRealtime_function(c *Realtime_functionContext)

	// ExitConversion_functions is called when exiting the conversion_functions production.
	ExitConversion_functions(c *Conversion_functionsContext)

	// ExitConversion_function_name is called when exiting the conversion_function_name production.
	ExitConversion_function_name(c *Conversion_function_nameContext)

	// ExitConstant_argument is called when exiting the constant_argument production.
	ExitConstant_argument(c *Constant_argumentContext)

	// ExitRandom_function is called when exiting the random_function production.
	ExitRandom_function(c *Random_functionContext)

	// ExitSeed is called when exiting the seed production.
	ExitSeed(c *SeedContext)

	// ExitDist_functions is called when exiting the dist_functions production.
	ExitDist_functions(c *Dist_functionsContext)

	// ExitStart_ is called when exiting the start_ production.
	ExitStart_(c *Start_Context)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitMean is called when exiting the mean production.
	ExitMean(c *MeanContext)

	// ExitStandard_deviation is called when exiting the standard_deviation production.
	ExitStandard_deviation(c *Standard_deviationContext)

	// ExitDegree_of_freedom is called when exiting the degree_of_freedom production.
	ExitDegree_of_freedom(c *Degree_of_freedomContext)

	// ExitK_stage is called when exiting the k_stage production.
	ExitK_stage(c *K_stageContext)

	// ExitMath_functions is called when exiting the math_functions production.
	ExitMath_functions(c *Math_functionsContext)

	// ExitInteger_math_functions is called when exiting the integer_math_functions production.
	ExitInteger_math_functions(c *Integer_math_functionsContext)

	// ExitReal_math_functions is called when exiting the real_math_functions production.
	ExitReal_math_functions(c *Real_math_functionsContext)

	// ExitSingle_argument_real_math_function_name is called when exiting the single_argument_real_math_function_name production.
	ExitSingle_argument_real_math_function_name(c *Single_argument_real_math_function_nameContext)

	// ExitDouble_argument_real_math_function_name is called when exiting the double_argument_real_math_function_name production.
	ExitDouble_argument_real_math_function_name(c *Double_argument_real_math_function_nameContext)

	// ExitDumpfile_task is called when exiting the dumpfile_task production.
	ExitDumpfile_task(c *Dumpfile_taskContext)

	// ExitDumpvars_task is called when exiting the dumpvars_task production.
	ExitDumpvars_task(c *Dumpvars_taskContext)

	// ExitList_of_modules_or_variables is called when exiting the list_of_modules_or_variables production.
	ExitList_of_modules_or_variables(c *List_of_modules_or_variablesContext)

	// ExitModule_or_variable is called when exiting the module_or_variable production.
	ExitModule_or_variable(c *Module_or_variableContext)

	// ExitLevels is called when exiting the levels production.
	ExitLevels(c *LevelsContext)

	// ExitDumpoff_task is called when exiting the dumpoff_task production.
	ExitDumpoff_task(c *Dumpoff_taskContext)

	// ExitDumpon_task is called when exiting the dumpon_task production.
	ExitDumpon_task(c *Dumpon_taskContext)

	// ExitDumpall_task is called when exiting the dumpall_task production.
	ExitDumpall_task(c *Dumpall_taskContext)

	// ExitDumplimit_task is called when exiting the dumplimit_task production.
	ExitDumplimit_task(c *Dumplimit_taskContext)

	// ExitFile_size is called when exiting the file_size production.
	ExitFile_size(c *File_sizeContext)

	// ExitDumpflush_task is called when exiting the dumpflush_task production.
	ExitDumpflush_task(c *Dumpflush_taskContext)

	// ExitDumpports_task is called when exiting the dumpports_task production.
	ExitDumpports_task(c *Dumpports_taskContext)

	// ExitScope_list is called when exiting the scope_list production.
	ExitScope_list(c *Scope_listContext)

	// ExitFile_pathname is called when exiting the file_pathname production.
	ExitFile_pathname(c *File_pathnameContext)

	// ExitDumpportsoff_task is called when exiting the dumpportsoff_task production.
	ExitDumpportsoff_task(c *Dumpportsoff_taskContext)

	// ExitDumpportson_task is called when exiting the dumpportson_task production.
	ExitDumpportson_task(c *Dumpportson_taskContext)

	// ExitDumpportsall_task is called when exiting the dumpportsall_task production.
	ExitDumpportsall_task(c *Dumpportsall_taskContext)

	// ExitDumpportslimit_task is called when exiting the dumpportslimit_task production.
	ExitDumpportslimit_task(c *Dumpportslimit_taskContext)

	// ExitDumpportsflush_task is called when exiting the dumpportsflush_task production.
	ExitDumpportsflush_task(c *Dumpportsflush_taskContext)

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

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitModule_keyword is called when exiting the module_keyword production.
	ExitModule_keyword(c *Module_keywordContext)

	// ExitModule_parameter_port_list is called when exiting the module_parameter_port_list production.
	ExitModule_parameter_port_list(c *Module_parameter_port_listContext)

	// ExitList_of_ports is called when exiting the list_of_ports production.
	ExitList_of_ports(c *List_of_portsContext)

	// ExitList_of_port_declarations is called when exiting the list_of_port_declarations production.
	ExitList_of_port_declarations(c *List_of_port_declarationsContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitPort_expression is called when exiting the port_expression production.
	ExitPort_expression(c *Port_expressionContext)

	// ExitPort_reference is called when exiting the port_reference production.
	ExitPort_reference(c *Port_referenceContext)

	// ExitPort_declaration is called when exiting the port_declaration production.
	ExitPort_declaration(c *Port_declarationContext)

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

	// ExitLocal_parameter_declaration is called when exiting the local_parameter_declaration production.
	ExitLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitSpecparam_declaration is called when exiting the specparam_declaration production.
	ExitSpecparam_declaration(c *Specparam_declarationContext)

	// ExitParameter_type is called when exiting the parameter_type production.
	ExitParameter_type(c *Parameter_typeContext)

	// ExitInout_declaration is called when exiting the inout_declaration production.
	ExitInout_declaration(c *Inout_declarationContext)

	// ExitInput_declaration is called when exiting the input_declaration production.
	ExitInput_declaration(c *Input_declarationContext)

	// ExitOutput_declaration is called when exiting the output_declaration production.
	ExitOutput_declaration(c *Output_declarationContext)

	// ExitEvent_declaration is called when exiting the event_declaration production.
	ExitEvent_declaration(c *Event_declarationContext)

	// ExitInteger_declaration is called when exiting the integer_declaration production.
	ExitInteger_declaration(c *Integer_declarationContext)

	// ExitNet_declaration is called when exiting the net_declaration production.
	ExitNet_declaration(c *Net_declarationContext)

	// ExitReal_declaration is called when exiting the real_declaration production.
	ExitReal_declaration(c *Real_declarationContext)

	// ExitRealtime_declaration is called when exiting the realtime_declaration production.
	ExitRealtime_declaration(c *Realtime_declarationContext)

	// ExitReg_declaration is called when exiting the reg_declaration production.
	ExitReg_declaration(c *Reg_declarationContext)

	// ExitTime_declaration is called when exiting the time_declaration production.
	ExitTime_declaration(c *Time_declarationContext)

	// ExitNet_type is called when exiting the net_type production.
	ExitNet_type(c *Net_typeContext)

	// ExitOutput_variable_type is called when exiting the output_variable_type production.
	ExitOutput_variable_type(c *Output_variable_typeContext)

	// ExitReal_type is called when exiting the real_type production.
	ExitReal_type(c *Real_typeContext)

	// ExitVariable_type is called when exiting the variable_type production.
	ExitVariable_type(c *Variable_typeContext)

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

	// ExitList_of_event_identifiers is called when exiting the list_of_event_identifiers production.
	ExitList_of_event_identifiers(c *List_of_event_identifiersContext)

	// ExitList_of_net_decl_assignments is called when exiting the list_of_net_decl_assignments production.
	ExitList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// ExitList_of_net_identifiers is called when exiting the list_of_net_identifiers production.
	ExitList_of_net_identifiers(c *List_of_net_identifiersContext)

	// ExitList_of_param_assignments is called when exiting the list_of_param_assignments production.
	ExitList_of_param_assignments(c *List_of_param_assignmentsContext)

	// ExitList_of_port_identifiers is called when exiting the list_of_port_identifiers production.
	ExitList_of_port_identifiers(c *List_of_port_identifiersContext)

	// ExitList_of_real_identifiers is called when exiting the list_of_real_identifiers production.
	ExitList_of_real_identifiers(c *List_of_real_identifiersContext)

	// ExitList_of_specparam_assignments is called when exiting the list_of_specparam_assignments production.
	ExitList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

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

	// ExitPulse_control_specparam is called when exiting the pulse_control_specparam production.
	ExitPulse_control_specparam(c *Pulse_control_specparamContext)

	// ExitError_limit_value is called when exiting the error_limit_value production.
	ExitError_limit_value(c *Error_limit_valueContext)

	// ExitReject_limit_value is called when exiting the reject_limit_value production.
	ExitReject_limit_value(c *Reject_limit_valueContext)

	// ExitLimit_value is called when exiting the limit_value production.
	ExitLimit_value(c *Limit_valueContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitRange_ is called when exiting the range_ production.
	ExitRange_(c *Range_Context)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_item_declaration is called when exiting the function_item_declaration production.
	ExitFunction_item_declaration(c *Function_item_declarationContext)

	// ExitFunction_port_list is called when exiting the function_port_list production.
	ExitFunction_port_list(c *Function_port_listContext)

	// ExitFunction_range_or_type is called when exiting the function_range_or_type production.
	ExitFunction_range_or_type(c *Function_range_or_typeContext)

	// ExitTask_declaration is called when exiting the task_declaration production.
	ExitTask_declaration(c *Task_declarationContext)

	// ExitTask_item_declaration is called when exiting the task_item_declaration production.
	ExitTask_item_declaration(c *Task_item_declarationContext)

	// ExitTask_port_list is called when exiting the task_port_list production.
	ExitTask_port_list(c *Task_port_listContext)

	// ExitTask_port_item is called when exiting the task_port_item production.
	ExitTask_port_item(c *Task_port_itemContext)

	// ExitTf_input_declaration is called when exiting the tf_input_declaration production.
	ExitTf_input_declaration(c *Tf_input_declarationContext)

	// ExitTf_output_declaration is called when exiting the tf_output_declaration production.
	ExitTf_output_declaration(c *Tf_output_declarationContext)

	// ExitTf_inout_declaration is called when exiting the tf_inout_declaration production.
	ExitTf_inout_declaration(c *Tf_inout_declarationContext)

	// ExitTask_port_type is called when exiting the task_port_type production.
	ExitTask_port_type(c *Task_port_typeContext)

	// ExitBlock_item_declaration is called when exiting the block_item_declaration production.
	ExitBlock_item_declaration(c *Block_item_declarationContext)

	// ExitList_of_block_variable_identifiers is called when exiting the list_of_block_variable_identifiers production.
	ExitList_of_block_variable_identifiers(c *List_of_block_variable_identifiersContext)

	// ExitList_of_block_real_identifiers is called when exiting the list_of_block_real_identifiers production.
	ExitList_of_block_real_identifiers(c *List_of_block_real_identifiersContext)

	// ExitBlock_variable_type is called when exiting the block_variable_type production.
	ExitBlock_variable_type(c *Block_variable_typeContext)

	// ExitBlock_real_type is called when exiting the block_real_type production.
	ExitBlock_real_type(c *Block_real_typeContext)

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

	// ExitName_of_gate_instance is called when exiting the name_of_gate_instance production.
	ExitName_of_gate_instance(c *Name_of_gate_instanceContext)

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

	// ExitModule_instance is called when exiting the module_instance production.
	ExitModule_instance(c *Module_instanceContext)

	// ExitName_of_module_instance is called when exiting the name_of_module_instance production.
	ExitName_of_module_instance(c *Name_of_module_instanceContext)

	// ExitList_of_port_connections is called when exiting the list_of_port_connections production.
	ExitList_of_port_connections(c *List_of_port_connectionsContext)

	// ExitOrdered_port_connection is called when exiting the ordered_port_connection production.
	ExitOrdered_port_connection(c *Ordered_port_connectionContext)

	// ExitNamed_port_connection is called when exiting the named_port_connection production.
	ExitNamed_port_connection(c *Named_port_connectionContext)

	// ExitGenerate_region is called when exiting the generate_region production.
	ExitGenerate_region(c *Generate_regionContext)

	// ExitGenvar_declaration is called when exiting the genvar_declaration production.
	ExitGenvar_declaration(c *Genvar_declarationContext)

	// ExitList_of_genvar_identifiers is called when exiting the list_of_genvar_identifiers production.
	ExitList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// ExitLoop_generate_construct is called when exiting the loop_generate_construct production.
	ExitLoop_generate_construct(c *Loop_generate_constructContext)

	// ExitGenvar_initialization is called when exiting the genvar_initialization production.
	ExitGenvar_initialization(c *Genvar_initializationContext)

	// ExitGenvar_expression is called when exiting the genvar_expression production.
	ExitGenvar_expression(c *Genvar_expressionContext)

	// ExitGenvar_iteration is called when exiting the genvar_iteration production.
	ExitGenvar_iteration(c *Genvar_iterationContext)

	// ExitGenvar_primary is called when exiting the genvar_primary production.
	ExitGenvar_primary(c *Genvar_primaryContext)

	// ExitConditional_generate_construct is called when exiting the conditional_generate_construct production.
	ExitConditional_generate_construct(c *Conditional_generate_constructContext)

	// ExitIf_generate_construct is called when exiting the if_generate_construct production.
	ExitIf_generate_construct(c *If_generate_constructContext)

	// ExitCase_generate_construct is called when exiting the case_generate_construct production.
	ExitCase_generate_construct(c *Case_generate_constructContext)

	// ExitGenerate_block is called when exiting the generate_block production.
	ExitGenerate_block(c *Generate_blockContext)

	// ExitGenerate_block_or_null is called when exiting the generate_block_or_null production.
	ExitGenerate_block_or_null(c *Generate_block_or_nullContext)

	// ExitContinuous_assign is called when exiting the continuous_assign production.
	ExitContinuous_assign(c *Continuous_assignContext)

	// ExitList_of_net_assignments is called when exiting the list_of_net_assignments production.
	ExitList_of_net_assignments(c *List_of_net_assignmentsContext)

	// ExitNet_assignment is called when exiting the net_assignment production.
	ExitNet_assignment(c *Net_assignmentContext)

	// ExitInitial_construct is called when exiting the initial_construct production.
	ExitInitial_construct(c *Initial_constructContext)

	// ExitAlways_construct is called when exiting the always_construct production.
	ExitAlways_construct(c *Always_constructContext)

	// ExitBlocking_assignment is called when exiting the blocking_assignment production.
	ExitBlocking_assignment(c *Blocking_assignmentContext)

	// ExitNonblocking_assignment is called when exiting the nonblocking_assignment production.
	ExitNonblocking_assignment(c *Nonblocking_assignmentContext)

	// ExitProcedural_continuous_assignments is called when exiting the procedural_continuous_assignments production.
	ExitProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitPar_block is called when exiting the par_block production.
	ExitPar_block(c *Par_blockContext)

	// ExitSeq_block is called when exiting the seq_block production.
	ExitSeq_block(c *Seq_blockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatement_or_null is called when exiting the statement_or_null production.
	ExitStatement_or_null(c *Statement_or_nullContext)

	// ExitFunction_statement is called when exiting the function_statement production.
	ExitFunction_statement(c *Function_statementContext)

	// ExitDelay_control is called when exiting the delay_control production.
	ExitDelay_control(c *Delay_controlContext)

	// ExitDelay_or_event_control is called when exiting the delay_or_event_control production.
	ExitDelay_or_event_control(c *Delay_or_event_controlContext)

	// ExitDisable_statement is called when exiting the disable_statement production.
	ExitDisable_statement(c *Disable_statementContext)

	// ExitEvent_control is called when exiting the event_control production.
	ExitEvent_control(c *Event_controlContext)

	// ExitEvent_trigger is called when exiting the event_trigger production.
	ExitEvent_trigger(c *Event_triggerContext)

	// ExitEvent_expression is called when exiting the event_expression production.
	ExitEvent_expression(c *Event_expressionContext)

	// ExitEvent_primary is called when exiting the event_primary production.
	ExitEvent_primary(c *Event_primaryContext)

	// ExitProcedural_timing_control is called when exiting the procedural_timing_control production.
	ExitProcedural_timing_control(c *Procedural_timing_controlContext)

	// ExitProcedural_timing_control_statement is called when exiting the procedural_timing_control_statement production.
	ExitProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitConditional_statement is called when exiting the conditional_statement production.
	ExitConditional_statement(c *Conditional_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_item is called when exiting the case_item production.
	ExitCase_item(c *Case_itemContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitSystem_task_enable is called when exiting the system_task_enable production.
	ExitSystem_task_enable(c *System_task_enableContext)

	// ExitTask_enable is called when exiting the task_enable production.
	ExitTask_enable(c *Task_enableContext)

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

	// ExitConstant_function_call is called when exiting the constant_function_call production.
	ExitConstant_function_call(c *Constant_function_callContext)

	// ExitConstant_system_function_call is called when exiting the constant_system_function_call production.
	ExitConstant_system_function_call(c *Constant_system_function_callContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitSystem_function_call is called when exiting the system_function_call production.
	ExitSystem_function_call(c *System_function_callContext)

	// ExitBase_expression is called when exiting the base_expression production.
	ExitBase_expression(c *Base_expressionContext)

	// ExitConstant_base_expression is called when exiting the constant_base_expression production.
	ExitConstant_base_expression(c *Constant_base_expressionContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitConstant_mintypmax_expression is called when exiting the constant_mintypmax_expression production.
	ExitConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// ExitConstant_range_expression is called when exiting the constant_range_expression production.
	ExitConstant_range_expression(c *Constant_range_expressionContext)

	// ExitDimension_constant_expression is called when exiting the dimension_constant_expression production.
	ExitDimension_constant_expression(c *Dimension_constant_expressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLsb_constant_expression is called when exiting the lsb_constant_expression production.
	ExitLsb_constant_expression(c *Lsb_constant_expressionContext)

	// ExitMintypmax_expression is called when exiting the mintypmax_expression production.
	ExitMintypmax_expression(c *Mintypmax_expressionContext)

	// ExitModule_path_expression is called when exiting the module_path_expression production.
	ExitModule_path_expression(c *Module_path_expressionContext)

	// ExitModule_path_mintypmax_expression is called when exiting the module_path_mintypmax_expression production.
	ExitModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// ExitMsb_constant_expression is called when exiting the msb_constant_expression production.
	ExitMsb_constant_expression(c *Msb_constant_expressionContext)

	// ExitRange_expression is called when exiting the range_expression production.
	ExitRange_expression(c *Range_expressionContext)

	// ExitWidth_constant_expression is called when exiting the width_constant_expression production.
	ExitWidth_constant_expression(c *Width_constant_expressionContext)

	// ExitConstant_primary is called when exiting the constant_primary production.
	ExitConstant_primary(c *Constant_primaryContext)

	// ExitModule_path_primary is called when exiting the module_path_primary production.
	ExitModule_path_primary(c *Module_path_primaryContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitNet_lvalue is called when exiting the net_lvalue production.
	ExitNet_lvalue(c *Net_lvalueContext)

	// ExitVariable_lvalue is called when exiting the variable_lvalue production.
	ExitVariable_lvalue(c *Variable_lvalueContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

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

	// ExitBlock_identifier is called when exiting the block_identifier production.
	ExitBlock_identifier(c *Block_identifierContext)

	// ExitCell_identifier is called when exiting the cell_identifier production.
	ExitCell_identifier(c *Cell_identifierContext)

	// ExitConfig_identifier is called when exiting the config_identifier production.
	ExitConfig_identifier(c *Config_identifierContext)

	// ExitEvent_identifier is called when exiting the event_identifier production.
	ExitEvent_identifier(c *Event_identifierContext)

	// ExitFunction_identifier is called when exiting the function_identifier production.
	ExitFunction_identifier(c *Function_identifierContext)

	// ExitGate_instance_identifier is called when exiting the gate_instance_identifier production.
	ExitGate_instance_identifier(c *Gate_instance_identifierContext)

	// ExitGenerate_block_identifier is called when exiting the generate_block_identifier production.
	ExitGenerate_block_identifier(c *Generate_block_identifierContext)

	// ExitGenvar_identifier is called when exiting the genvar_identifier production.
	ExitGenvar_identifier(c *Genvar_identifierContext)

	// ExitHierarchical_block_identifier is called when exiting the hierarchical_block_identifier production.
	ExitHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// ExitHierarchical_event_identifier is called when exiting the hierarchical_event_identifier production.
	ExitHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// ExitHierarchical_function_identifier is called when exiting the hierarchical_function_identifier production.
	ExitHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// ExitHierarchical_identifier is called when exiting the hierarchical_identifier production.
	ExitHierarchical_identifier(c *Hierarchical_identifierContext)

	// ExitHierarchical_net_identifier is called when exiting the hierarchical_net_identifier production.
	ExitHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// ExitHierarchical_parameter_identifier is called when exiting the hierarchical_parameter_identifier production.
	ExitHierarchical_parameter_identifier(c *Hierarchical_parameter_identifierContext)

	// ExitHierarchical_variable_identifier is called when exiting the hierarchical_variable_identifier production.
	ExitHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// ExitHierarchical_task_identifier is called when exiting the hierarchical_task_identifier production.
	ExitHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitInout_port_identifier is called when exiting the inout_port_identifier production.
	ExitInout_port_identifier(c *Inout_port_identifierContext)

	// ExitInput_port_identifier is called when exiting the input_port_identifier production.
	ExitInput_port_identifier(c *Input_port_identifierContext)

	// ExitInstance_identifier is called when exiting the instance_identifier production.
	ExitInstance_identifier(c *Instance_identifierContext)

	// ExitLibrary_identifier is called when exiting the library_identifier production.
	ExitLibrary_identifier(c *Library_identifierContext)

	// ExitModule_identifier is called when exiting the module_identifier production.
	ExitModule_identifier(c *Module_identifierContext)

	// ExitModule_instance_identifier is called when exiting the module_instance_identifier production.
	ExitModule_instance_identifier(c *Module_instance_identifierContext)

	// ExitNet_identifier is called when exiting the net_identifier production.
	ExitNet_identifier(c *Net_identifierContext)

	// ExitOutput_port_identifier is called when exiting the output_port_identifier production.
	ExitOutput_port_identifier(c *Output_port_identifierContext)

	// ExitParameter_identifier is called when exiting the parameter_identifier production.
	ExitParameter_identifier(c *Parameter_identifierContext)

	// ExitPort_identifier is called when exiting the port_identifier production.
	ExitPort_identifier(c *Port_identifierContext)

	// ExitReal_identifier is called when exiting the real_identifier production.
	ExitReal_identifier(c *Real_identifierContext)

	// ExitSpecparam_identifier is called when exiting the specparam_identifier production.
	ExitSpecparam_identifier(c *Specparam_identifierContext)

	// ExitSystem_function_identifier is called when exiting the system_function_identifier production.
	ExitSystem_function_identifier(c *System_function_identifierContext)

	// ExitSystem_task_identifier is called when exiting the system_task_identifier production.
	ExitSystem_task_identifier(c *System_task_identifierContext)

	// ExitTask_identifier is called when exiting the task_identifier production.
	ExitTask_identifier(c *Task_identifierContext)

	// ExitTerminal_identifier is called when exiting the terminal_identifier production.
	ExitTerminal_identifier(c *Terminal_identifierContext)

	// ExitTopmodule_identifier is called when exiting the topmodule_identifier production.
	ExitTopmodule_identifier(c *Topmodule_identifierContext)

	// ExitUdp_identifier is called when exiting the udp_identifier production.
	ExitUdp_identifier(c *Udp_identifierContext)

	// ExitUdp_instance_identifier is called when exiting the udp_instance_identifier production.
	ExitUdp_instance_identifier(c *Udp_instance_identifierContext)

	// ExitVariable_identifier is called when exiting the variable_identifier production.
	ExitVariable_identifier(c *Variable_identifierContext)
}
