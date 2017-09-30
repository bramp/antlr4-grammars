// Generated from TSqlParser.g4 by ANTLR 4.7.

package tsql // TSqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTSqlParserVisitor) VisitTsql_file(ctx *Tsql_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBatch(ctx *BatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSql_clauses(ctx *Sql_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSql_clause(ctx *Sql_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDml_clause(ctx *Dml_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDdl_clause(ctx *Ddl_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_statement(ctx *Backup_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCfl_statement(ctx *Cfl_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBlock_statement(ctx *Block_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGoto_statement(ctx *Goto_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitThrow_statement(ctx *Throw_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitThrow_error_number(ctx *Throw_error_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitThrow_message(ctx *Throw_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitThrow_state(ctx *Throw_stateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTry_catch_statement(ctx *Try_catch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWaitfor_statement(ctx *Waitfor_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPrint_statement(ctx *Print_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRaiseerror_statement(ctx *Raiseerror_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEmpty_statement(ctx *Empty_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAnother_statement(ctx *Another_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_application_role(ctx *Alter_application_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_application_role(ctx *Create_application_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_aggregate(ctx *Drop_aggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_application_role(ctx *Drop_application_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly(ctx *Alter_assemblyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_start(ctx *Alter_assembly_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_clause(ctx *Alter_assembly_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_from_clause(ctx *Alter_assembly_from_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_from_clause_start(ctx *Alter_assembly_from_clause_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_drop_clause(ctx *Alter_assembly_drop_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_drop_multiple_files(ctx *Alter_assembly_drop_multiple_filesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_drop(ctx *Alter_assembly_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_add_clause(ctx *Alter_assembly_add_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_asssembly_add_clause_start(ctx *Alter_asssembly_add_clause_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_client_file_clause(ctx *Alter_assembly_client_file_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_file_name(ctx *Alter_assembly_file_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_file_bits(ctx *Alter_assembly_file_bitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_as(ctx *Alter_assembly_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_with_clause(ctx *Alter_assembly_with_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_assembly_with(ctx *Alter_assembly_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClient_assembly_specifier(ctx *Client_assembly_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAssembly_option(ctx *Assembly_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNetwork_file_share(ctx *Network_file_shareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNetwork_computer(ctx *Network_computerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNetwork_file_start(ctx *Network_file_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFile_path(ctx *File_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFile_directory_path_separator(ctx *File_directory_path_separatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitLocal_file(ctx *Local_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitLocal_drive(ctx *Local_driveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMultiple_local_files(ctx *Multiple_local_filesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMultiple_local_file_start(ctx *Multiple_local_file_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_assembly(ctx *Create_assemblyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_assembly(ctx *Drop_assemblyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_asymmetric_key(ctx *Alter_asymmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_asymmetric_key_start(ctx *Alter_asymmetric_key_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAsymmetric_key_option(ctx *Asymmetric_key_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAsymmetric_key_option_start(ctx *Asymmetric_key_option_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAsymmetric_key_password_change_option(ctx *Asymmetric_key_password_change_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_asymmetric_key(ctx *Create_asymmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_asymmetric_key(ctx *Drop_asymmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_authorization(ctx *Alter_authorizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAuthorization_grantee(ctx *Authorization_granteeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEntity_to(ctx *Entity_toContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColon_colon(ctx *Colon_colonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_authorization_start(ctx *Alter_authorization_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_authorization_for_sql_database(ctx *Alter_authorization_for_sql_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_authorization_for_azure_dw(ctx *Alter_authorization_for_azure_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_authorization_for_parallel_dw(ctx *Alter_authorization_for_parallel_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClass_type(ctx *Class_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClass_type_for_sql_database(ctx *Class_type_for_sql_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClass_type_for_azure_dw(ctx *Class_type_for_azure_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClass_type_for_parallel_dw(ctx *Class_type_for_parallel_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_availability_group(ctx *Drop_availability_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_availability_group(ctx *Alter_availability_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_availability_group_start(ctx *Alter_availability_group_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_availability_group_options(ctx *Alter_availability_group_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_broker_priority(ctx *Create_or_alter_broker_priorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_broker_priority(ctx *Drop_broker_priorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_certificate(ctx *Alter_certificateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_column_encryption_key(ctx *Alter_column_encryption_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_column_encryption_key(ctx *Create_column_encryption_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_certificate(ctx *Drop_certificateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_column_encryption_key(ctx *Drop_column_encryption_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_column_master_key(ctx *Drop_column_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_contract(ctx *Drop_contractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_credential(ctx *Drop_credentialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_cryptograhic_provider(ctx *Drop_cryptograhic_providerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_database(ctx *Drop_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_database_audit_specification(ctx *Drop_database_audit_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_database_scoped_credential(ctx *Drop_database_scoped_credentialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_default(ctx *Drop_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_endpoint(ctx *Drop_endpointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_external_data_source(ctx *Drop_external_data_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_external_file_format(ctx *Drop_external_file_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_external_library(ctx *Drop_external_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_external_resource_pool(ctx *Drop_external_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_external_table(ctx *Drop_external_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_event_notifications(ctx *Drop_event_notificationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_event_session(ctx *Drop_event_sessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_fulltext_catalog(ctx *Drop_fulltext_catalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_fulltext_index(ctx *Drop_fulltext_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_fulltext_stoplist(ctx *Drop_fulltext_stoplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_login(ctx *Drop_loginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_master_key(ctx *Drop_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_message_type(ctx *Drop_message_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_partition_function(ctx *Drop_partition_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_partition_scheme(ctx *Drop_partition_schemeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_queue(ctx *Drop_queueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_remote_service_binding(ctx *Drop_remote_service_bindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_resource_pool(ctx *Drop_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_db_role(ctx *Drop_db_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_route(ctx *Drop_routeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_rule(ctx *Drop_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_schema(ctx *Drop_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_search_property_list(ctx *Drop_search_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_security_policy(ctx *Drop_security_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_sequence(ctx *Drop_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_server_audit(ctx *Drop_server_auditContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_server_audit_specification(ctx *Drop_server_audit_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_server_role(ctx *Drop_server_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_service(ctx *Drop_serviceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_signature(ctx *Drop_signatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_statistics_name_azure_dw_and_pdw(ctx *Drop_statistics_name_azure_dw_and_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_symmetric_key(ctx *Drop_symmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_synonym(ctx *Drop_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_user(ctx *Drop_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_workload_group(ctx *Drop_workload_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_xml_schema_collection(ctx *Drop_xml_schema_collectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDisable_trigger(ctx *Disable_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEnable_trigger(ctx *Enable_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTruncate_table(ctx *Truncate_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_column_master_key(ctx *Create_column_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_credential(ctx *Alter_credentialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_credential(ctx *Create_credentialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_cryptographic_provider(ctx *Alter_cryptographic_providerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_cryptographic_provider(ctx *Create_cryptographic_providerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_event_notification(ctx *Create_event_notificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_event_session(ctx *Create_or_alter_event_sessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEvent_session_predicate_expression(ctx *Event_session_predicate_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEvent_session_predicate_factor(ctx *Event_session_predicate_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEvent_session_predicate_leaf(ctx *Event_session_predicate_leafContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_external_data_source(ctx *Alter_external_data_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_external_library(ctx *Alter_external_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_external_library(ctx *Create_external_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_external_resource_pool(ctx *Alter_external_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_external_resource_pool(ctx *Create_external_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_fulltext_catalog(ctx *Alter_fulltext_catalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_fulltext_catalog(ctx *Create_fulltext_catalogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_fulltext_stoplist(ctx *Alter_fulltext_stoplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_fulltext_stoplist(ctx *Create_fulltext_stoplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_login_sql_server(ctx *Alter_login_sql_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_login_sql_server(ctx *Create_login_sql_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_login_azure_sql(ctx *Alter_login_azure_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_login_azure_sql(ctx *Create_login_azure_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_login_azure_sql_dw_and_pdw(ctx *Alter_login_azure_sql_dw_and_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_login_pdw(ctx *Create_login_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_master_key_sql_server(ctx *Alter_master_key_sql_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_master_key_sql_server(ctx *Create_master_key_sql_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_master_key_azure_sql(ctx *Alter_master_key_azure_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_master_key_azure_sql(ctx *Create_master_key_azure_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_message_type(ctx *Alter_message_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_partition_function(ctx *Alter_partition_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_partition_scheme(ctx *Alter_partition_schemeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_remote_service_binding(ctx *Alter_remote_service_bindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_remote_service_binding(ctx *Create_remote_service_bindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_resource_pool(ctx *Create_resource_poolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_resource_governor(ctx *Alter_resource_governorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_db_role(ctx *Alter_db_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_db_role(ctx *Create_db_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_route(ctx *Create_routeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_rule(ctx *Create_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_schema_sql(ctx *Alter_schema_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_schema(ctx *Create_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_schema_azure_sql_dw_and_pdw(ctx *Create_schema_azure_sql_dw_and_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_schema_azure_sql_dw_and_pdw(ctx *Alter_schema_azure_sql_dw_and_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_search_property_list(ctx *Create_search_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_security_policy(ctx *Create_security_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_sequence(ctx *Alter_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_sequence(ctx *Create_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_server_audit(ctx *Alter_server_auditContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_server_audit(ctx *Create_server_auditContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_server_audit_specification(ctx *Alter_server_audit_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_server_audit_specification(ctx *Create_server_audit_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_server_configuration(ctx *Alter_server_configurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_server_role(ctx *Alter_server_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_server_role(ctx *Create_server_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_server_role_pdw(ctx *Alter_server_role_pdwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_service(ctx *Alter_serviceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_service(ctx *Create_serviceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_service_master_key(ctx *Alter_service_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_symmetric_key(ctx *Alter_symmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_symmetric_key(ctx *Create_symmetric_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_synonym(ctx *Create_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_user(ctx *Alter_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_user(ctx *Create_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_user_azure_sql_dw(ctx *Create_user_azure_sql_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_user_azure_sql(ctx *Alter_user_azure_sqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_workload_group(ctx *Alter_workload_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_workload_group(ctx *Create_workload_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_xml_schema_collection(ctx *Create_xml_schema_collectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_queue(ctx *Create_queueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQueue_settings(ctx *Queue_settingsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_queue(ctx *Alter_queueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQueue_action(ctx *Queue_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQueue_rebuild_options(ctx *Queue_rebuild_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_contract(ctx *Create_contractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitConversation_statement(ctx *Conversation_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMessage_statement(ctx *Message_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMerge_statement(ctx *Merge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMerge_matched(ctx *Merge_matchedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMerge_not_matched(ctx *Merge_not_matchedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDelete_statement_from(ctx *Delete_statement_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitInsert_statement(ctx *Insert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitInsert_statement_value(ctx *Insert_statement_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitReceive_statement(ctx *Receive_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTime(ctx *TimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOutput_clause(ctx *Output_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOutput_dml_list_elem(ctx *Output_dml_list_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOutput_column_name(ctx *Output_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_database(ctx *Create_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_index(ctx *Create_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_procedure(ctx *Create_or_alter_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_trigger(ctx *Create_or_alter_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_dml_trigger(ctx *Create_or_alter_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDml_trigger_option(ctx *Dml_trigger_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDml_trigger_operation(ctx *Dml_trigger_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_ddl_trigger(ctx *Create_or_alter_ddl_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDdl_trigger_operation(ctx *Ddl_trigger_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_or_alter_function(ctx *Create_or_alter_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFunc_body_returns_select(ctx *Func_body_returns_selectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFunc_body_returns_table(ctx *Func_body_returns_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFunc_body_returns_scalar(ctx *Func_body_returns_scalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitProcedure_param(ctx *Procedure_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitProcedure_option(ctx *Procedure_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFunction_option(ctx *Function_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_statistics(ctx *Create_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_table(ctx *Create_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_options(ctx *Table_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_view(ctx *Create_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitView_attribute(ctx *View_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_table(ctx *Alter_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_database(ctx *Alter_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDatabase_optionspec(ctx *Database_optionspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAuto_option(ctx *Auto_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitChange_tracking_option(ctx *Change_tracking_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitChange_tracking_option_list(ctx *Change_tracking_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitContainment_option(ctx *Containment_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCursor_option(ctx *Cursor_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlter_endpoint(ctx *Alter_endpointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDatabase_mirroring_option(ctx *Database_mirroring_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMirroring_set_option(ctx *Mirroring_set_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMirroring_partner(ctx *Mirroring_partnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMirroring_witness(ctx *Mirroring_witnessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWitness_partner_equal(ctx *Witness_partner_equalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPartner_option(ctx *Partner_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWitness_option(ctx *Witness_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWitness_server(ctx *Witness_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPartner_server(ctx *Partner_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMirroring_host_port_seperator(ctx *Mirroring_host_port_seperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPartner_server_tcp_prefix(ctx *Partner_server_tcp_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPort_number(ctx *Port_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitHost(ctx *HostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDb_encryption_option(ctx *Db_encryption_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDb_state_option(ctx *Db_state_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDb_update_option(ctx *Db_update_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDb_user_access_option(ctx *Db_user_access_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDelayed_durability_option(ctx *Delayed_durability_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExternal_access_option(ctx *External_access_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitHadr_options(ctx *Hadr_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitParameterization_option(ctx *Parameterization_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRecovery_option(ctx *Recovery_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitService_broker_option(ctx *Service_broker_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSnapshot_option(ctx *Snapshot_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSql_option(ctx *Sql_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTermination(ctx *TerminationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_index(ctx *Drop_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_relational_or_xml_or_spatial_index(ctx *Drop_relational_or_xml_or_spatial_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_backward_compatible_index(ctx *Drop_backward_compatible_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_procedure(ctx *Drop_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_trigger(ctx *Drop_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_dml_trigger(ctx *Drop_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_ddl_trigger(ctx *Drop_ddl_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_function(ctx *Drop_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_statistics(ctx *Drop_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_table(ctx *Drop_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_view(ctx *Drop_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_type(ctx *Create_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDrop_type(ctx *Drop_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRowset_function_limited(ctx *Rowset_function_limitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOpenquery(ctx *OpenqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOpendatasource(ctx *OpendatasourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDeclare_statement(ctx *Declare_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCursor_statement(ctx *Cursor_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_database(ctx *Backup_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_log(ctx *Backup_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_certificate(ctx *Backup_certificateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_master_key(ctx *Backup_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBackup_service_master_key(ctx *Backup_service_master_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExecute_statement(ctx *Execute_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExecute_statement_arg(ctx *Execute_statement_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExecute_var_string(ctx *Execute_var_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSecurity_statement(ctx *Security_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_certificate(ctx *Create_certificateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExisting_keys(ctx *Existing_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPrivate_key_options(ctx *Private_key_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGenerate_new_keys(ctx *Generate_new_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDate_options(ctx *Date_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOpen_key(ctx *Open_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClose_key(ctx *Close_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_key(ctx *Create_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitKey_options(ctx *Key_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAlgorithm(ctx *AlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEncryption_mechanism(ctx *Encryption_mechanismContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDecryption_mechanism(ctx *Decryption_mechanismContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGrant_permission(ctx *Grant_permissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSet_statement(ctx *Set_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTransaction_statement(ctx *Transaction_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGo_statement(ctx *Go_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUse_statement(ctx *Use_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDbcc_clause(ctx *Dbcc_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDbcc_options(ctx *Dbcc_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExecute_clause(ctx *Execute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDeclare_local(ctx *Declare_localContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_type_definition(ctx *Table_type_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitXml_type_definition(ctx *Xml_type_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitXml_schema_collection(ctx *Xml_schema_collectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_def_table_constraint(ctx *Column_def_table_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_definition(ctx *Column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_constraint(ctx *Column_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_constraint(ctx *Table_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOn_delete(ctx *On_deleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOn_update(ctx *On_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitIndex_options(ctx *Index_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitIndex_option(ctx *Index_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDeclare_cursor(ctx *Declare_cursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDeclare_set_cursor_common_partial(ctx *Declare_set_cursor_common_partialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFetch_cursor(ctx *Fetch_cursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSet_special(ctx *Set_specialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPrimitive_expression(ctx *Primitive_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCase_expression(ctx *Case_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUnary_operator_expression(ctx *Unary_operator_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBracket_expression(ctx *Bracket_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitConstant_expression(ctx *Constant_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWith_expression(ctx *With_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUpdate_elem(ctx *Update_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSearch_condition_list(ctx *Search_condition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSearch_condition(ctx *Search_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSearch_condition_and(ctx *Search_condition_andContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSearch_condition_not(ctx *Search_condition_notContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQuery_expression(ctx *Query_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSql_union(ctx *Sql_unionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQuery_specification(ctx *Query_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTop_clause(ctx *Top_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTop_percent(ctx *Top_percentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTop_count(ctx *Top_countContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFor_clause(ctx *For_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitXml_common_directives(ctx *Xml_common_directivesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOrder_by_expression(ctx *Order_by_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGroup_by_item(ctx *Group_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOption_clause(ctx *Option_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOptimize_for_arg(ctx *Optimize_for_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSelect_list(ctx *Select_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUdt_method_arguments(ctx *Udt_method_argumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAsterisk(ctx *AsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_elem(ctx *Column_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUdt_elem(ctx *Udt_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExpression_elem(ctx *Expression_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSelect_list_elem(ctx *Select_list_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_sources(ctx *Table_sourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_source(ctx *Table_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_source_item_joined(ctx *Table_source_item_joinedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_source_item(ctx *Table_source_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOpen_xml(ctx *Open_xmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSchema_declaration(ctx *Schema_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_declaration(ctx *Column_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitChange_table(ctx *Change_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitJoin_part(ctx *Join_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitPivot_clause(ctx *Pivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFull_column_name_list(ctx *Full_column_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_name_with_hint(ctx *Table_name_with_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRowset_function(ctx *Rowset_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBulk_option(ctx *Bulk_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDerived_table(ctx *Derived_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRANKING_WINDOWED_FUNC(ctx *RANKING_WINDOWED_FUNCContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAGGREGATE_WINDOWED_FUNC(ctx *AGGREGATE_WINDOWED_FUNCContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitANALYTIC_WINDOWED_FUNC(ctx *ANALYTIC_WINDOWED_FUNCContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSCALAR_FUNCTION(ctx *SCALAR_FUNCTIONContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBINARY_CHECKSUM(ctx *BINARY_CHECKSUMContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCAST(ctx *CASTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCONVERT(ctx *CONVERTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCHECKSUM(ctx *CHECKSUMContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCOALESCE(ctx *COALESCEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCURRENT_TIMESTAMP(ctx *CURRENT_TIMESTAMPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCURRENT_USER(ctx *CURRENT_USERContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDATEADD(ctx *DATEADDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDATEDIFF(ctx *DATEDIFFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDATENAME(ctx *DATENAMEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDATEPART(ctx *DATEPARTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGETDATE(ctx *GETDATEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGETUTCDATE(ctx *GETUTCDATEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitIDENTITY(ctx *IDENTITYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitMIN_ACTIVE_ROWVERSION(ctx *MIN_ACTIVE_ROWVERSIONContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNULLIF(ctx *NULLIFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSTUFF(ctx *STUFFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSESSION_USER(ctx *SESSION_USERContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSYSTEM_USER(ctx *SYSTEM_USERContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitISNULL(ctx *ISNULLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitXML_DATA_TYPE_FUNC(ctx *XML_DATA_TYPE_FUNCContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitXml_data_type_methods(ctx *Xml_data_type_methodsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitValue_method(ctx *Value_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQuery_method(ctx *Query_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExist_method(ctx *Exist_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitModify_method(ctx *Modify_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNodes_method(ctx *Nodes_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSwitch_section(ctx *Switch_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAs_column_alias(ctx *As_column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAs_table_alias(ctx *As_table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWith_table_hints(ctx *With_table_hintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitInsert_with_table_hints(ctx *Insert_with_table_hintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_hint(ctx *Table_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitIndex_value(ctx *Index_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_alias_list(ctx *Column_alias_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_value_constructor(ctx *Table_value_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRanking_windowed_function(ctx *Ranking_windowed_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAnalytic_windowed_function(ctx *Analytic_windowed_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAll_distinct_expression(ctx *All_distinct_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitRow_or_range_clause(ctx *Row_or_range_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWindow_frame_extent(ctx *Window_frame_extentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWindow_frame_bound(ctx *Window_frame_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWindow_frame_preceding(ctx *Window_frame_precedingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWindow_frame_following(ctx *Window_frame_followingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCreate_database_option(ctx *Create_database_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDatabase_filestream_option(ctx *Database_filestream_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDatabase_file_spec(ctx *Database_file_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFile_group(ctx *File_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFile_spec(ctx *File_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEntity_name(ctx *Entity_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEntity_name_for_azure_dw(ctx *Entity_name_for_azure_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEntity_name_for_parallel_dw(ctx *Entity_name_for_parallel_dwContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFull_table_name(ctx *Full_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSimple_name(ctx *Simple_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFunc_proc_name(ctx *Func_proc_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDdl_object(ctx *Ddl_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFull_column_name(ctx *Full_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_name_list_with_order(ctx *Column_name_list_with_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitColumn_name_list(ctx *Column_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitOn_off(ctx *On_offContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitClustered(ctx *ClusteredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNull_notnull(ctx *Null_notnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitNull_or_default(ctx *Null_or_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitScalar_function_name(ctx *Scalar_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBegin_conversation_timer(ctx *Begin_conversation_timerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitBegin_conversation_dialog(ctx *Begin_conversation_dialogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitContract_name(ctx *Contract_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitService_name(ctx *Service_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitEnd_conversation(ctx *End_conversationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitWaitfor_conversation(ctx *Waitfor_conversationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitGet_conversation(ctx *Get_conversationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitQueue_id(ctx *Queue_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSend_conversation(ctx *Send_conversationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitData_type(ctx *Data_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitDefault_value(ctx *Default_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSign(ctx *SignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitSimple_id(ctx *Simple_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitComparison_operator(ctx *Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitAssignment_operator(ctx *Assignment_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTSqlParserVisitor) VisitFile_size(ctx *File_sizeContext) interface{} {
	return v.VisitChildren(ctx)
}
