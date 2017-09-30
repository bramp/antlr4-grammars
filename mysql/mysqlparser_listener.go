// Generated from MySqlParser.g4 by ANTLR 4.7.

package mysql // MySqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type MySqlParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSql_statements is called when entering the sql_statements production.
	EnterSql_statements(c *Sql_statementsContext)

	// EnterSql_statement is called when entering the sql_statement production.
	EnterSql_statement(c *Sql_statementContext)

	// EnterEmpty_statement is called when entering the empty_statement production.
	EnterEmpty_statement(c *Empty_statementContext)

	// EnterDdl_statement is called when entering the ddl_statement production.
	EnterDdl_statement(c *Ddl_statementContext)

	// EnterDml_statement is called when entering the dml_statement production.
	EnterDml_statement(c *Dml_statementContext)

	// EnterTransaction_statement is called when entering the transaction_statement production.
	EnterTransaction_statement(c *Transaction_statementContext)

	// EnterReplication_statement is called when entering the replication_statement production.
	EnterReplication_statement(c *Replication_statementContext)

	// EnterPrepared_statement is called when entering the prepared_statement production.
	EnterPrepared_statement(c *Prepared_statementContext)

	// EnterCompound_statement is called when entering the compound_statement production.
	EnterCompound_statement(c *Compound_statementContext)

	// EnterAdministration_statement is called when entering the administration_statement production.
	EnterAdministration_statement(c *Administration_statementContext)

	// EnterUtility_statement is called when entering the utility_statement production.
	EnterUtility_statement(c *Utility_statementContext)

	// EnterCreate_database is called when entering the create_database production.
	EnterCreate_database(c *Create_databaseContext)

	// EnterCreate_event is called when entering the create_event production.
	EnterCreate_event(c *Create_eventContext)

	// EnterCreate_index is called when entering the create_index production.
	EnterCreate_index(c *Create_indexContext)

	// EnterCreate_logfile_group is called when entering the create_logfile_group production.
	EnterCreate_logfile_group(c *Create_logfile_groupContext)

	// EnterCreate_procedure is called when entering the create_procedure production.
	EnterCreate_procedure(c *Create_procedureContext)

	// EnterCreate_function is called when entering the create_function production.
	EnterCreate_function(c *Create_functionContext)

	// EnterCreate_server is called when entering the create_server production.
	EnterCreate_server(c *Create_serverContext)

	// EnterCopyCreateTable is called when entering the copyCreateTable production.
	EnterCopyCreateTable(c *CopyCreateTableContext)

	// EnterQueryCreateTable is called when entering the queryCreateTable production.
	EnterQueryCreateTable(c *QueryCreateTableContext)

	// EnterColCreateTable is called when entering the colCreateTable production.
	EnterColCreateTable(c *ColCreateTableContext)

	// EnterCreate_tablespace_innodb is called when entering the create_tablespace_innodb production.
	EnterCreate_tablespace_innodb(c *Create_tablespace_innodbContext)

	// EnterCreate_tablespace_ndb is called when entering the create_tablespace_ndb production.
	EnterCreate_tablespace_ndb(c *Create_tablespace_ndbContext)

	// EnterCreate_trigger is called when entering the create_trigger production.
	EnterCreate_trigger(c *Create_triggerContext)

	// EnterCreate_view is called when entering the create_view production.
	EnterCreate_view(c *Create_viewContext)

	// EnterCreate_database_option is called when entering the create_database_option production.
	EnterCreate_database_option(c *Create_database_optionContext)

	// EnterOwner_statement is called when entering the owner_statement production.
	EnterOwner_statement(c *Owner_statementContext)

	// EnterPreciseSchedule is called when entering the preciseSchedule production.
	EnterPreciseSchedule(c *PreciseScheduleContext)

	// EnterIntervalSchedule is called when entering the intervalSchedule production.
	EnterIntervalSchedule(c *IntervalScheduleContext)

	// EnterTimestamp_value is called when entering the timestamp_value production.
	EnterTimestamp_value(c *Timestamp_valueContext)

	// EnterInterval_expr is called when entering the interval_expr production.
	EnterInterval_expr(c *Interval_exprContext)

	// EnterInterval_type is called when entering the interval_type production.
	EnterInterval_type(c *Interval_typeContext)

	// EnterIndex_type is called when entering the index_type production.
	EnterIndex_type(c *Index_typeContext)

	// EnterIndex_option is called when entering the index_option production.
	EnterIndex_option(c *Index_optionContext)

	// EnterProc_param is called when entering the proc_param production.
	EnterProc_param(c *Proc_paramContext)

	// EnterFunc_param is called when entering the func_param production.
	EnterFunc_param(c *Func_paramContext)

	// EnterRcComment is called when entering the rcComment production.
	EnterRcComment(c *RcCommentContext)

	// EnterRcSqllang is called when entering the rcSqllang production.
	EnterRcSqllang(c *RcSqllangContext)

	// EnterRcDeterm is called when entering the rcDeterm production.
	EnterRcDeterm(c *RcDetermContext)

	// EnterRcSqldata is called when entering the rcSqldata production.
	EnterRcSqldata(c *RcSqldataContext)

	// EnterRcSecurestmt is called when entering the rcSecurestmt production.
	EnterRcSecurestmt(c *RcSecurestmtContext)

	// EnterServer_option is called when entering the server_option production.
	EnterServer_option(c *Server_optionContext)

	// EnterColumn_def_table_constraints is called when entering the column_def_table_constraints production.
	EnterColumn_def_table_constraints(c *Column_def_table_constraintsContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterConstraintDefinition is called when entering the constraintDefinition production.
	EnterConstraintDefinition(c *ConstraintDefinitionContext)

	// EnterIndexDefinition is called when entering the indexDefinition production.
	EnterIndexDefinition(c *IndexDefinitionContext)

	// EnterColumn_definition is called when entering the column_definition production.
	EnterColumn_definition(c *Column_definitionContext)

	// EnterColConstrNull is called when entering the colConstrNull production.
	EnterColConstrNull(c *ColConstrNullContext)

	// EnterColConstrDflt is called when entering the colConstrDflt production.
	EnterColConstrDflt(c *ColConstrDfltContext)

	// EnterColConstrAuInc is called when entering the colConstrAuInc production.
	EnterColConstrAuInc(c *ColConstrAuIncContext)

	// EnterColConstrPK is called when entering the colConstrPK production.
	EnterColConstrPK(c *ColConstrPKContext)

	// EnterColConstrUK is called when entering the colConstrUK production.
	EnterColConstrUK(c *ColConstrUKContext)

	// EnterColConstrComment is called when entering the colConstrComment production.
	EnterColConstrComment(c *ColConstrCommentContext)

	// EnterColConstrForm is called when entering the colConstrForm production.
	EnterColConstrForm(c *ColConstrFormContext)

	// EnterColConstrStorage is called when entering the colConstrStorage production.
	EnterColConstrStorage(c *ColConstrStorageContext)

	// EnterColConstrRefdef is called when entering the colConstrRefdef production.
	EnterColConstrRefdef(c *ColConstrRefdefContext)

	// EnterTblConstrPK is called when entering the tblConstrPK production.
	EnterTblConstrPK(c *TblConstrPKContext)

	// EnterTblConstrUK is called when entering the tblConstrUK production.
	EnterTblConstrUK(c *TblConstrUKContext)

	// EnterTblConstrFK is called when entering the tblConstrFK production.
	EnterTblConstrFK(c *TblConstrFKContext)

	// EnterTblConstCheck is called when entering the tblConstCheck production.
	EnterTblConstCheck(c *TblConstCheckContext)

	// EnterReference_definition is called when entering the reference_definition production.
	EnterReference_definition(c *Reference_definitionContext)

	// EnterOn_delete_action is called when entering the on_delete_action production.
	EnterOn_delete_action(c *On_delete_actionContext)

	// EnterOn_update_action is called when entering the on_update_action production.
	EnterOn_update_action(c *On_update_actionContext)

	// EnterReference_action_control_type is called when entering the reference_action_control_type production.
	EnterReference_action_control_type(c *Reference_action_control_typeContext)

	// EnterSimpleIndex is called when entering the simpleIndex production.
	EnterSimpleIndex(c *SimpleIndexContext)

	// EnterSpecIndex is called when entering the specIndex production.
	EnterSpecIndex(c *SpecIndexContext)

	// EnterTblOptEngine is called when entering the tblOptEngine production.
	EnterTblOptEngine(c *TblOptEngineContext)

	// EnterTblOptAuInc is called when entering the tblOptAuInc production.
	EnterTblOptAuInc(c *TblOptAuIncContext)

	// EnterTblOptAvgRLen is called when entering the tblOptAvgRLen production.
	EnterTblOptAvgRLen(c *TblOptAvgRLenContext)

	// EnterTblOptDefCharSet is called when entering the tblOptDefCharSet production.
	EnterTblOptDefCharSet(c *TblOptDefCharSetContext)

	// EnterTblOptChkSum is called when entering the tblOptChkSum production.
	EnterTblOptChkSum(c *TblOptChkSumContext)

	// EnterTblOptDefCollate is called when entering the tblOptDefCollate production.
	EnterTblOptDefCollate(c *TblOptDefCollateContext)

	// EnterTblOptComment is called when entering the tblOptComment production.
	EnterTblOptComment(c *TblOptCommentContext)

	// EnterTblOptCompr is called when entering the tblOptCompr production.
	EnterTblOptCompr(c *TblOptComprContext)

	// EnterTblOptConn is called when entering the tblOptConn production.
	EnterTblOptConn(c *TblOptConnContext)

	// EnterTblOptDataDir is called when entering the tblOptDataDir production.
	EnterTblOptDataDir(c *TblOptDataDirContext)

	// EnterTblOptDelKW is called when entering the tblOptDelKW production.
	EnterTblOptDelKW(c *TblOptDelKWContext)

	// EnterTblOptEncr is called when entering the tblOptEncr production.
	EnterTblOptEncr(c *TblOptEncrContext)

	// EnterTblOptIndexDir is called when entering the tblOptIndexDir production.
	EnterTblOptIndexDir(c *TblOptIndexDirContext)

	// EnterTblOptInsMeth is called when entering the tblOptInsMeth production.
	EnterTblOptInsMeth(c *TblOptInsMethContext)

	// EnterTblOptKeyBlockSz is called when entering the tblOptKeyBlockSz production.
	EnterTblOptKeyBlockSz(c *TblOptKeyBlockSzContext)

	// EnterTblOptMaxRows is called when entering the tblOptMaxRows production.
	EnterTblOptMaxRows(c *TblOptMaxRowsContext)

	// EnterTblOptMinRows is called when entering the tblOptMinRows production.
	EnterTblOptMinRows(c *TblOptMinRowsContext)

	// EnterTblOptPackK is called when entering the tblOptPackK production.
	EnterTblOptPackK(c *TblOptPackKContext)

	// EnterTblOptPasswd is called when entering the tblOptPasswd production.
	EnterTblOptPasswd(c *TblOptPasswdContext)

	// EnterTblOptRowFormat is called when entering the tblOptRowFormat production.
	EnterTblOptRowFormat(c *TblOptRowFormatContext)

	// EnterTblOptStatAutoR is called when entering the tblOptStatAutoR production.
	EnterTblOptStatAutoR(c *TblOptStatAutoRContext)

	// EnterTblOptStatPersist is called when entering the tblOptStatPersist production.
	EnterTblOptStatPersist(c *TblOptStatPersistContext)

	// EnterTblOptStatSamplPg is called when entering the tblOptStatSamplPg production.
	EnterTblOptStatSamplPg(c *TblOptStatSamplPgContext)

	// EnterTblOptTablespace is called when entering the tblOptTablespace production.
	EnterTblOptTablespace(c *TblOptTablespaceContext)

	// EnterTblOptUnion is called when entering the tblOptUnion production.
	EnterTblOptUnion(c *TblOptUnionContext)

	// EnterPartition_options is called when entering the partition_options production.
	EnterPartition_options(c *Partition_optionsContext)

	// EnterPartition_function_definition is called when entering the partition_function_definition production.
	EnterPartition_function_definition(c *Partition_function_definitionContext)

	// EnterLinear_partition_func_def is called when entering the linear_partition_func_def production.
	EnterLinear_partition_func_def(c *Linear_partition_func_defContext)

	// EnterPartition_def is called when entering the partition_def production.
	EnterPartition_def(c *Partition_defContext)

	// EnterSubpartition_def is called when entering the subpartition_def production.
	EnterSubpartition_def(c *Subpartition_defContext)

	// EnterAlterDb is called when entering the alterDb production.
	EnterAlterDb(c *AlterDbContext)

	// EnterAlterDbUpgradeName is called when entering the alterDbUpgradeName production.
	EnterAlterDbUpgradeName(c *AlterDbUpgradeNameContext)

	// EnterAlter_event is called when entering the alter_event production.
	EnterAlter_event(c *Alter_eventContext)

	// EnterAlter_function is called when entering the alter_function production.
	EnterAlter_function(c *Alter_functionContext)

	// EnterAlter_instance is called when entering the alter_instance production.
	EnterAlter_instance(c *Alter_instanceContext)

	// EnterAlter_logfile_group is called when entering the alter_logfile_group production.
	EnterAlter_logfile_group(c *Alter_logfile_groupContext)

	// EnterAlter_procedure is called when entering the alter_procedure production.
	EnterAlter_procedure(c *Alter_procedureContext)

	// EnterAlter_server is called when entering the alter_server production.
	EnterAlter_server(c *Alter_serverContext)

	// EnterAlter_table is called when entering the alter_table production.
	EnterAlter_table(c *Alter_tableContext)

	// EnterAlter_tablespace is called when entering the alter_tablespace production.
	EnterAlter_tablespace(c *Alter_tablespaceContext)

	// EnterAlter_view is called when entering the alter_view production.
	EnterAlter_view(c *Alter_viewContext)

	// EnterAltblTableOpt is called when entering the altblTableOpt production.
	EnterAltblTableOpt(c *AltblTableOptContext)

	// EnterAltblAddCol is called when entering the altblAddCol production.
	EnterAltblAddCol(c *AltblAddColContext)

	// EnterAltblAddCols is called when entering the altblAddCols production.
	EnterAltblAddCols(c *AltblAddColsContext)

	// EnterAltblAddIndex is called when entering the altblAddIndex production.
	EnterAltblAddIndex(c *AltblAddIndexContext)

	// EnterAltblAddPK is called when entering the altblAddPK production.
	EnterAltblAddPK(c *AltblAddPKContext)

	// EnterAltblAddUK is called when entering the altblAddUK production.
	EnterAltblAddUK(c *AltblAddUKContext)

	// EnterAltblAddSpecIndex is called when entering the altblAddSpecIndex production.
	EnterAltblAddSpecIndex(c *AltblAddSpecIndexContext)

	// EnterAltblAddFK is called when entering the altblAddFK production.
	EnterAltblAddFK(c *AltblAddFKContext)

	// EnterAltblAlg is called when entering the altblAlg production.
	EnterAltblAlg(c *AltblAlgContext)

	// EnterAltblColDef is called when entering the altblColDef production.
	EnterAltblColDef(c *AltblColDefContext)

	// EnterAltblColChange is called when entering the altblColChange production.
	EnterAltblColChange(c *AltblColChangeContext)

	// EnterAltblLock is called when entering the altblLock production.
	EnterAltblLock(c *AltblLockContext)

	// EnterAltblColMod is called when entering the altblColMod production.
	EnterAltblColMod(c *AltblColModContext)

	// EnterAltblColDrop is called when entering the altblColDrop production.
	EnterAltblColDrop(c *AltblColDropContext)

	// EnterAltblDropPK is called when entering the altblDropPK production.
	EnterAltblDropPK(c *AltblDropPKContext)

	// EnterAltblDropIndex is called when entering the altblDropIndex production.
	EnterAltblDropIndex(c *AltblDropIndexContext)

	// EnterAltblDropFK is called when entering the altblDropFK production.
	EnterAltblDropFK(c *AltblDropFKContext)

	// EnterAltblDisKey is called when entering the altblDisKey production.
	EnterAltblDisKey(c *AltblDisKeyContext)

	// EnterAltblEnKey is called when entering the altblEnKey production.
	EnterAltblEnKey(c *AltblEnKeyContext)

	// EnterAltblRenameTbl is called when entering the altblRenameTbl production.
	EnterAltblRenameTbl(c *AltblRenameTblContext)

	// EnterAltblResort is called when entering the altblResort production.
	EnterAltblResort(c *AltblResortContext)

	// EnterAltblConvert is called when entering the altblConvert production.
	EnterAltblConvert(c *AltblConvertContext)

	// EnterAltblDefCharset is called when entering the altblDefCharset production.
	EnterAltblDefCharset(c *AltblDefCharsetContext)

	// EnterAltblDisTblspace is called when entering the altblDisTblspace production.
	EnterAltblDisTblspace(c *AltblDisTblspaceContext)

	// EnterAltblImpTblSpace is called when entering the altblImpTblSpace production.
	EnterAltblImpTblSpace(c *AltblImpTblSpaceContext)

	// EnterAltblForce is called when entering the altblForce production.
	EnterAltblForce(c *AltblForceContext)

	// EnterAltblValid is called when entering the altblValid production.
	EnterAltblValid(c *AltblValidContext)

	// EnterAltblAddPart is called when entering the altblAddPart production.
	EnterAltblAddPart(c *AltblAddPartContext)

	// EnterAltblDropPart is called when entering the altblDropPart production.
	EnterAltblDropPart(c *AltblDropPartContext)

	// EnterAltblDiscartPart is called when entering the altblDiscartPart production.
	EnterAltblDiscartPart(c *AltblDiscartPartContext)

	// EnterAltblImportPart is called when entering the altblImportPart production.
	EnterAltblImportPart(c *AltblImportPartContext)

	// EnterAltblTruncPart is called when entering the altblTruncPart production.
	EnterAltblTruncPart(c *AltblTruncPartContext)

	// EnterAltblCoalPart is called when entering the altblCoalPart production.
	EnterAltblCoalPart(c *AltblCoalPartContext)

	// EnterAltblReorgPart is called when entering the altblReorgPart production.
	EnterAltblReorgPart(c *AltblReorgPartContext)

	// EnterAltblExchPart is called when entering the altblExchPart production.
	EnterAltblExchPart(c *AltblExchPartContext)

	// EnterAltblAnalPart is called when entering the altblAnalPart production.
	EnterAltblAnalPart(c *AltblAnalPartContext)

	// EnterAltblCheckPart is called when entering the altblCheckPart production.
	EnterAltblCheckPart(c *AltblCheckPartContext)

	// EnterAltblOptimPart is called when entering the altblOptimPart production.
	EnterAltblOptimPart(c *AltblOptimPartContext)

	// EnterAltblRebuildPart is called when entering the altblRebuildPart production.
	EnterAltblRebuildPart(c *AltblRebuildPartContext)

	// EnterAltblRepairPart is called when entering the altblRepairPart production.
	EnterAltblRepairPart(c *AltblRepairPartContext)

	// EnterAltblRemovePart is called when entering the altblRemovePart production.
	EnterAltblRemovePart(c *AltblRemovePartContext)

	// EnterAltblUpgrPart is called when entering the altblUpgrPart production.
	EnterAltblUpgrPart(c *AltblUpgrPartContext)

	// EnterDrop_database is called when entering the drop_database production.
	EnterDrop_database(c *Drop_databaseContext)

	// EnterDrop_event is called when entering the drop_event production.
	EnterDrop_event(c *Drop_eventContext)

	// EnterDrop_index is called when entering the drop_index production.
	EnterDrop_index(c *Drop_indexContext)

	// EnterDrop_logfile_group is called when entering the drop_logfile_group production.
	EnterDrop_logfile_group(c *Drop_logfile_groupContext)

	// EnterDrop_procedure is called when entering the drop_procedure production.
	EnterDrop_procedure(c *Drop_procedureContext)

	// EnterDrop_function is called when entering the drop_function production.
	EnterDrop_function(c *Drop_functionContext)

	// EnterDrop_server is called when entering the drop_server production.
	EnterDrop_server(c *Drop_serverContext)

	// EnterDrop_table is called when entering the drop_table production.
	EnterDrop_table(c *Drop_tableContext)

	// EnterDrop_tablespace is called when entering the drop_tablespace production.
	EnterDrop_tablespace(c *Drop_tablespaceContext)

	// EnterDrop_trigger is called when entering the drop_trigger production.
	EnterDrop_trigger(c *Drop_triggerContext)

	// EnterDrop_view is called when entering the drop_view production.
	EnterDrop_view(c *Drop_viewContext)

	// EnterRename_table is called when entering the rename_table production.
	EnterRename_table(c *Rename_tableContext)

	// EnterTruncate_table is called when entering the truncate_table production.
	EnterTruncate_table(c *Truncate_tableContext)

	// EnterCall_statement is called when entering the call_statement production.
	EnterCall_statement(c *Call_statementContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterDo_statement is called when entering the do_statement production.
	EnterDo_statement(c *Do_statementContext)

	// EnterHandler_statement is called when entering the handler_statement production.
	EnterHandler_statement(c *Handler_statementContext)

	// EnterInsert_statement is called when entering the insert_statement production.
	EnterInsert_statement(c *Insert_statementContext)

	// EnterLoad_data_statement is called when entering the load_data_statement production.
	EnterLoad_data_statement(c *Load_data_statementContext)

	// EnterLoad_xml_statement is called when entering the load_xml_statement production.
	EnterLoad_xml_statement(c *Load_xml_statementContext)

	// EnterReplace_statement is called when entering the replace_statement production.
	EnterReplace_statement(c *Replace_statementContext)

	// EnterSimpleSelect is called when entering the simpleSelect production.
	EnterSimpleSelect(c *SimpleSelectContext)

	// EnterParenSelect is called when entering the parenSelect production.
	EnterParenSelect(c *ParenSelectContext)

	// EnterUnionSelect is called when entering the unionSelect production.
	EnterUnionSelect(c *UnionSelectContext)

	// EnterUnionParenSelect is called when entering the unionParenSelect production.
	EnterUnionParenSelect(c *UnionParenSelectContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterInsert_statement_value is called when entering the insert_statement_value production.
	EnterInsert_statement_value(c *Insert_statement_valueContext)

	// EnterUpdate_elem is called when entering the update_elem production.
	EnterUpdate_elem(c *Update_elemContext)

	// EnterCol_or_uservar is called when entering the col_or_uservar production.
	EnterCol_or_uservar(c *Col_or_uservarContext)

	// EnterSingle_delete_statement is called when entering the single_delete_statement production.
	EnterSingle_delete_statement(c *Single_delete_statementContext)

	// EnterMultiple_delete_statement is called when entering the multiple_delete_statement production.
	EnterMultiple_delete_statement(c *Multiple_delete_statementContext)

	// EnterHandler_open_statement is called when entering the handler_open_statement production.
	EnterHandler_open_statement(c *Handler_open_statementContext)

	// EnterHandler_read_index_statement is called when entering the handler_read_index_statement production.
	EnterHandler_read_index_statement(c *Handler_read_index_statementContext)

	// EnterHandler_read_statement is called when entering the handler_read_statement production.
	EnterHandler_read_statement(c *Handler_read_statementContext)

	// EnterHandler_close_statement is called when entering the handler_close_statement production.
	EnterHandler_close_statement(c *Handler_close_statementContext)

	// EnterSingle_update_statement is called when entering the single_update_statement production.
	EnterSingle_update_statement(c *Single_update_statementContext)

	// EnterMultiple_update_statement is called when entering the multiple_update_statement production.
	EnterMultiple_update_statement(c *Multiple_update_statementContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterOrder_by_expression is called when entering the order_by_expression production.
	EnterOrder_by_expression(c *Order_by_expressionContext)

	// EnterTable_sources is called when entering the table_sources production.
	EnterTable_sources(c *Table_sourcesContext)

	// EnterTable_source is called when entering the table_source production.
	EnterTable_source(c *Table_sourceContext)

	// EnterAtomTableItem is called when entering the atomTableItem production.
	EnterAtomTableItem(c *AtomTableItemContext)

	// EnterSubqueryTableItem is called when entering the subqueryTableItem production.
	EnterSubqueryTableItem(c *SubqueryTableItemContext)

	// EnterTableSourcesItem is called when entering the tableSourcesItem production.
	EnterTableSourcesItem(c *TableSourcesItemContext)

	// EnterIndex_hint is called when entering the index_hint production.
	EnterIndex_hint(c *Index_hintContext)

	// EnterInnerJoin is called when entering the innerJoin production.
	EnterInnerJoin(c *InnerJoinContext)

	// EnterStraightJoin is called when entering the straightJoin production.
	EnterStraightJoin(c *StraightJoinContext)

	// EnterOuterJoin is called when entering the outerJoin production.
	EnterOuterJoin(c *OuterJoinContext)

	// EnterNaturalJoin is called when entering the naturalJoin production.
	EnterNaturalJoin(c *NaturalJoinContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterQuery_expression is called when entering the query_expression production.
	EnterQuery_expression(c *Query_expressionContext)

	// EnterQuery_expression_nointo is called when entering the query_expression_nointo production.
	EnterQuery_expression_nointo(c *Query_expression_nointoContext)

	// EnterQuery_specification is called when entering the query_specification production.
	EnterQuery_specification(c *Query_specificationContext)

	// EnterQuery_specification_nointo is called when entering the query_specification_nointo production.
	EnterQuery_specification_nointo(c *Query_specification_nointoContext)

	// EnterUnion_parenth is called when entering the union_parenth production.
	EnterUnion_parenth(c *Union_parenthContext)

	// EnterUnion_statement is called when entering the union_statement production.
	EnterUnion_statement(c *Union_statementContext)

	// EnterSelect_spec is called when entering the select_spec production.
	EnterSelect_spec(c *Select_specContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSellistelAllCol is called when entering the sellistelAllCol production.
	EnterSellistelAllCol(c *SellistelAllColContext)

	// EnterSellistelCol is called when entering the sellistelCol production.
	EnterSellistelCol(c *SellistelColContext)

	// EnterSellistelFunc is called when entering the sellistelFunc production.
	EnterSellistelFunc(c *SellistelFuncContext)

	// EnterSellistelExpr is called when entering the sellistelExpr production.
	EnterSellistelExpr(c *SellistelExprContext)

	// EnterSelectIntoVars is called when entering the selectIntoVars production.
	EnterSelectIntoVars(c *SelectIntoVarsContext)

	// EnterSelectIntoDump is called when entering the selectIntoDump production.
	EnterSelectIntoDump(c *SelectIntoDumpContext)

	// EnterSelectIntoOutfile is called when entering the selectIntoOutfile production.
	EnterSelectIntoOutfile(c *SelectIntoOutfileContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterGroup_by_item is called when entering the group_by_item production.
	EnterGroup_by_item(c *Group_by_itemContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterStart_transaction is called when entering the start_transaction production.
	EnterStart_transaction(c *Start_transactionContext)

	// EnterBegin_work is called when entering the begin_work production.
	EnterBegin_work(c *Begin_workContext)

	// EnterCommit_work is called when entering the commit_work production.
	EnterCommit_work(c *Commit_workContext)

	// EnterRollback_work is called when entering the rollback_work production.
	EnterRollback_work(c *Rollback_workContext)

	// EnterSavepoint_statement is called when entering the savepoint_statement production.
	EnterSavepoint_statement(c *Savepoint_statementContext)

	// EnterRollback_statement is called when entering the rollback_statement production.
	EnterRollback_statement(c *Rollback_statementContext)

	// EnterRelease_statement is called when entering the release_statement production.
	EnterRelease_statement(c *Release_statementContext)

	// EnterLock_tables is called when entering the lock_tables production.
	EnterLock_tables(c *Lock_tablesContext)

	// EnterUnlock_tables is called when entering the unlock_tables production.
	EnterUnlock_tables(c *Unlock_tablesContext)

	// EnterSet_autocommit_statement is called when entering the set_autocommit_statement production.
	EnterSet_autocommit_statement(c *Set_autocommit_statementContext)

	// EnterSet_transaction_statement is called when entering the set_transaction_statement production.
	EnterSet_transaction_statement(c *Set_transaction_statementContext)

	// EnterTransact_option is called when entering the transact_option production.
	EnterTransact_option(c *Transact_optionContext)

	// EnterLock_table_element is called when entering the lock_table_element production.
	EnterLock_table_element(c *Lock_table_elementContext)

	// EnterTrans_characteristic is called when entering the trans_characteristic production.
	EnterTrans_characteristic(c *Trans_characteristicContext)

	// EnterTransaction_level is called when entering the transaction_level production.
	EnterTransaction_level(c *Transaction_levelContext)

	// EnterChange_master is called when entering the change_master production.
	EnterChange_master(c *Change_masterContext)

	// EnterChange_repl_filter is called when entering the change_repl_filter production.
	EnterChange_repl_filter(c *Change_repl_filterContext)

	// EnterPurge_binary_logs is called when entering the purge_binary_logs production.
	EnterPurge_binary_logs(c *Purge_binary_logsContext)

	// EnterReset_master is called when entering the reset_master production.
	EnterReset_master(c *Reset_masterContext)

	// EnterReset_slave is called when entering the reset_slave production.
	EnterReset_slave(c *Reset_slaveContext)

	// EnterStart_slave is called when entering the start_slave production.
	EnterStart_slave(c *Start_slaveContext)

	// EnterStop_slave is called when entering the stop_slave production.
	EnterStop_slave(c *Stop_slaveContext)

	// EnterStart_group_repl is called when entering the start_group_repl production.
	EnterStart_group_repl(c *Start_group_replContext)

	// EnterStop_group_repl is called when entering the stop_group_repl production.
	EnterStop_group_repl(c *Stop_group_replContext)

	// EnterMasterOptString is called when entering the masterOptString production.
	EnterMasterOptString(c *MasterOptStringContext)

	// EnterMasterOptDecimal is called when entering the masterOptDecimal production.
	EnterMasterOptDecimal(c *MasterOptDecimalContext)

	// EnterMasterOptBool is called when entering the masterOptBool production.
	EnterMasterOptBool(c *MasterOptBoolContext)

	// EnterMasterOptReal is called when entering the masterOptReal production.
	EnterMasterOptReal(c *MasterOptRealContext)

	// EnterMasterOptIdList is called when entering the masterOptIdList production.
	EnterMasterOptIdList(c *MasterOptIdListContext)

	// EnterString_master_option is called when entering the string_master_option production.
	EnterString_master_option(c *String_master_optionContext)

	// EnterDecimal_master_option is called when entering the decimal_master_option production.
	EnterDecimal_master_option(c *Decimal_master_optionContext)

	// EnterBool_master_option is called when entering the bool_master_option production.
	EnterBool_master_option(c *Bool_master_optionContext)

	// EnterChannel_option is called when entering the channel_option production.
	EnterChannel_option(c *Channel_optionContext)

	// EnterReplfilterDbList is called when entering the replfilterDbList production.
	EnterReplfilterDbList(c *ReplfilterDbListContext)

	// EnterReplfilterTableList is called when entering the replfilterTableList production.
	EnterReplfilterTableList(c *ReplfilterTableListContext)

	// EnterReplfilterStableList is called when entering the replfilterStableList production.
	EnterReplfilterStableList(c *ReplfilterStableListContext)

	// EnterReplfilterTablepairList is called when entering the replfilterTablepairList production.
	EnterReplfilterTablepairList(c *ReplfilterTablepairListContext)

	// EnterThread_type is called when entering the thread_type production.
	EnterThread_type(c *Thread_typeContext)

	// EnterUntilGtidSset is called when entering the untilGtidSset production.
	EnterUntilGtidSset(c *UntilGtidSsetContext)

	// EnterUntilMasterLog is called when entering the untilMasterLog production.
	EnterUntilMasterLog(c *UntilMasterLogContext)

	// EnterUntilRelayLog is called when entering the untilRelayLog production.
	EnterUntilRelayLog(c *UntilRelayLogContext)

	// EnterUntilSqlGaps is called when entering the untilSqlGaps production.
	EnterUntilSqlGaps(c *UntilSqlGapsContext)

	// EnterStart_slave_connection_option is called when entering the start_slave_connection_option production.
	EnterStart_slave_connection_option(c *Start_slave_connection_optionContext)

	// EnterGtid_set is called when entering the gtid_set production.
	EnterGtid_set(c *Gtid_setContext)

	// EnterXa_start_transaction is called when entering the xa_start_transaction production.
	EnterXa_start_transaction(c *Xa_start_transactionContext)

	// EnterXa_end_transaction is called when entering the xa_end_transaction production.
	EnterXa_end_transaction(c *Xa_end_transactionContext)

	// EnterXa_prepare is called when entering the xa_prepare production.
	EnterXa_prepare(c *Xa_prepareContext)

	// EnterXa_commit_work is called when entering the xa_commit_work production.
	EnterXa_commit_work(c *Xa_commit_workContext)

	// EnterXa_rollback_work is called when entering the xa_rollback_work production.
	EnterXa_rollback_work(c *Xa_rollback_workContext)

	// EnterXa_recover_work is called when entering the xa_recover_work production.
	EnterXa_recover_work(c *Xa_recover_workContext)

	// EnterPrepare_statement is called when entering the prepare_statement production.
	EnterPrepare_statement(c *Prepare_statementContext)

	// EnterExecute_statement is called when entering the execute_statement production.
	EnterExecute_statement(c *Execute_statementContext)

	// EnterDeallocate_prepare is called when entering the deallocate_prepare production.
	EnterDeallocate_prepare(c *Deallocate_prepareContext)

	// EnterRoutine_body is called when entering the routine_body production.
	EnterRoutine_body(c *Routine_bodyContext)

	// EnterBlock_statement is called when entering the block_statement production.
	EnterBlock_statement(c *Block_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterIterate_statement is called when entering the iterate_statement production.
	EnterIterate_statement(c *Iterate_statementContext)

	// EnterLeave_statement is called when entering the leave_statement production.
	EnterLeave_statement(c *Leave_statementContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterRepeat_statement is called when entering the repeat_statement production.
	EnterRepeat_statement(c *Repeat_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterCursor_statement is called when entering the cursor_statement production.
	EnterCursor_statement(c *Cursor_statementContext)

	// EnterDeclare_variable is called when entering the declare_variable production.
	EnterDeclare_variable(c *Declare_variableContext)

	// EnterDeclare_condition is called when entering the declare_condition production.
	EnterDeclare_condition(c *Declare_conditionContext)

	// EnterDeclare_cursor is called when entering the declare_cursor production.
	EnterDeclare_cursor(c *Declare_cursorContext)

	// EnterDeclare_handler is called when entering the declare_handler production.
	EnterDeclare_handler(c *Declare_handlerContext)

	// EnterHandler_condition_value is called when entering the handler_condition_value production.
	EnterHandler_condition_value(c *Handler_condition_valueContext)

	// EnterProcedure_sql_statement is called when entering the procedure_sql_statement production.
	EnterProcedure_sql_statement(c *Procedure_sql_statementContext)

	// EnterAlterUserMysql56 is called when entering the alterUserMysql56 production.
	EnterAlterUserMysql56(c *AlterUserMysql56Context)

	// EnterAlterUserMysql57 is called when entering the alterUserMysql57 production.
	EnterAlterUserMysql57(c *AlterUserMysql57Context)

	// EnterCreateUserMysql56 is called when entering the createUserMysql56 production.
	EnterCreateUserMysql56(c *CreateUserMysql56Context)

	// EnterCreateUserMysql57 is called when entering the createUserMysql57 production.
	EnterCreateUserMysql57(c *CreateUserMysql57Context)

	// EnterDrop_user is called when entering the drop_user production.
	EnterDrop_user(c *Drop_userContext)

	// EnterGrant_statement is called when entering the grant_statement production.
	EnterGrant_statement(c *Grant_statementContext)

	// EnterGrant_proxy is called when entering the grant_proxy production.
	EnterGrant_proxy(c *Grant_proxyContext)

	// EnterRename_user is called when entering the rename_user production.
	EnterRename_user(c *Rename_userContext)

	// EnterDetailRevoke is called when entering the detailRevoke production.
	EnterDetailRevoke(c *DetailRevokeContext)

	// EnterShortRevoke is called when entering the shortRevoke production.
	EnterShortRevoke(c *ShortRevokeContext)

	// EnterRevoke_proxy is called when entering the revoke_proxy production.
	EnterRevoke_proxy(c *Revoke_proxyContext)

	// EnterSet_password_statement is called when entering the set_password_statement production.
	EnterSet_password_statement(c *Set_password_statementContext)

	// EnterUser_password_option is called when entering the user_password_option production.
	EnterUser_password_option(c *User_password_optionContext)

	// EnterAuthByPassword is called when entering the authByPassword production.
	EnterAuthByPassword(c *AuthByPasswordContext)

	// EnterAuthByString is called when entering the authByString production.
	EnterAuthByString(c *AuthByStringContext)

	// EnterAuthByHash is called when entering the authByHash production.
	EnterAuthByHash(c *AuthByHashContext)

	// EnterTls_option is called when entering the tls_option production.
	EnterTls_option(c *Tls_optionContext)

	// EnterUser_resource_option is called when entering the user_resource_option production.
	EnterUser_resource_option(c *User_resource_optionContext)

	// EnterUser_lock_option is called when entering the user_lock_option production.
	EnterUser_lock_option(c *User_lock_optionContext)

	// EnterPrivelege_clause is called when entering the privelege_clause production.
	EnterPrivelege_clause(c *Privelege_clauseContext)

	// EnterPrivilege is called when entering the privilege production.
	EnterPrivilege(c *PrivilegeContext)

	// EnterPrivilege_level is called when entering the privilege_level production.
	EnterPrivilege_level(c *Privilege_levelContext)

	// EnterSet_password_option is called when entering the set_password_option production.
	EnterSet_password_option(c *Set_password_optionContext)

	// EnterAnalyze_table is called when entering the analyze_table production.
	EnterAnalyze_table(c *Analyze_tableContext)

	// EnterCheck_table is called when entering the check_table production.
	EnterCheck_table(c *Check_tableContext)

	// EnterChecksum_table is called when entering the checksum_table production.
	EnterChecksum_table(c *Checksum_tableContext)

	// EnterOptimize_table is called when entering the optimize_table production.
	EnterOptimize_table(c *Optimize_tableContext)

	// EnterRepair_table is called when entering the repair_table production.
	EnterRepair_table(c *Repair_tableContext)

	// EnterCheck_table_option is called when entering the check_table_option production.
	EnterCheck_table_option(c *Check_table_optionContext)

	// EnterCreate_udfunction is called when entering the create_udfunction production.
	EnterCreate_udfunction(c *Create_udfunctionContext)

	// EnterInstall_plugin is called when entering the install_plugin production.
	EnterInstall_plugin(c *Install_pluginContext)

	// EnterUninstall_plugin is called when entering the uninstall_plugin production.
	EnterUninstall_plugin(c *Uninstall_pluginContext)

	// EnterSetVariableAssignment is called when entering the setVariableAssignment production.
	EnterSetVariableAssignment(c *SetVariableAssignmentContext)

	// EnterSetCharset is called when entering the setCharset production.
	EnterSetCharset(c *SetCharsetContext)

	// EnterSetNames is called when entering the setNames production.
	EnterSetNames(c *SetNamesContext)

	// EnterSetPasswordStatement is called when entering the setPasswordStatement production.
	EnterSetPasswordStatement(c *SetPasswordStatementContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterSetAutocommit is called when entering the setAutocommit production.
	EnterSetAutocommit(c *SetAutocommitContext)

	// EnterShowMasterlogs is called when entering the showMasterlogs production.
	EnterShowMasterlogs(c *ShowMasterlogsContext)

	// EnterShowLogevents is called when entering the showLogevents production.
	EnterShowLogevents(c *ShowLogeventsContext)

	// EnterShowObjWithFilter is called when entering the showObjWithFilter production.
	EnterShowObjWithFilter(c *ShowObjWithFilterContext)

	// EnterShowColumns is called when entering the showColumns production.
	EnterShowColumns(c *ShowColumnsContext)

	// EnterShowCreateDb is called when entering the showCreateDb production.
	EnterShowCreateDb(c *ShowCreateDbContext)

	// EnterShowCreateFullidobj is called when entering the showCreateFullidobj production.
	EnterShowCreateFullidobj(c *ShowCreateFullidobjContext)

	// EnterShowCreateUser is called when entering the showCreateUser production.
	EnterShowCreateUser(c *ShowCreateUserContext)

	// EnterShowEngine is called when entering the showEngine production.
	EnterShowEngine(c *ShowEngineContext)

	// EnterShowGlobalinfo is called when entering the showGlobalinfo production.
	EnterShowGlobalinfo(c *ShowGlobalinfoContext)

	// EnterShowErrWarn is called when entering the showErrWarn production.
	EnterShowErrWarn(c *ShowErrWarnContext)

	// EnterShowCountErrWarn is called when entering the showCountErrWarn production.
	EnterShowCountErrWarn(c *ShowCountErrWarnContext)

	// EnterShowFromschemaFilter is called when entering the showFromschemaFilter production.
	EnterShowFromschemaFilter(c *ShowFromschemaFilterContext)

	// EnterShowRoutinecode is called when entering the showRoutinecode production.
	EnterShowRoutinecode(c *ShowRoutinecodeContext)

	// EnterShowGrants is called when entering the showGrants production.
	EnterShowGrants(c *ShowGrantsContext)

	// EnterShowIndexes is called when entering the showIndexes production.
	EnterShowIndexes(c *ShowIndexesContext)

	// EnterShowOpentables is called when entering the showOpentables production.
	EnterShowOpentables(c *ShowOpentablesContext)

	// EnterShowProfile is called when entering the showProfile production.
	EnterShowProfile(c *ShowProfileContext)

	// EnterShowSlavestatus is called when entering the showSlavestatus production.
	EnterShowSlavestatus(c *ShowSlavestatusContext)

	// EnterVariable_clause is called when entering the variable_clause production.
	EnterVariable_clause(c *Variable_clauseContext)

	// EnterShow_filter is called when entering the show_filter production.
	EnterShow_filter(c *Show_filterContext)

	// EnterShow_profile_type is called when entering the show_profile_type production.
	EnterShow_profile_type(c *Show_profile_typeContext)

	// EnterBinlog_statement is called when entering the binlog_statement production.
	EnterBinlog_statement(c *Binlog_statementContext)

	// EnterCache_index_statement is called when entering the cache_index_statement production.
	EnterCache_index_statement(c *Cache_index_statementContext)

	// EnterFlush_statement is called when entering the flush_statement production.
	EnterFlush_statement(c *Flush_statementContext)

	// EnterKill_statement is called when entering the kill_statement production.
	EnterKill_statement(c *Kill_statementContext)

	// EnterLoad_index_into_cache is called when entering the load_index_into_cache production.
	EnterLoad_index_into_cache(c *Load_index_into_cacheContext)

	// EnterReset_statement is called when entering the reset_statement production.
	EnterReset_statement(c *Reset_statementContext)

	// EnterShutdown_statement is called when entering the shutdown_statement production.
	EnterShutdown_statement(c *Shutdown_statementContext)

	// EnterTbl_index_list is called when entering the tbl_index_list production.
	EnterTbl_index_list(c *Tbl_index_listContext)

	// EnterFlush_option is called when entering the flush_option production.
	EnterFlush_option(c *Flush_optionContext)

	// EnterLoad_tbl_index_list is called when entering the load_tbl_index_list production.
	EnterLoad_tbl_index_list(c *Load_tbl_index_listContext)

	// EnterSimple_describe_statement is called when entering the simple_describe_statement production.
	EnterSimple_describe_statement(c *Simple_describe_statementContext)

	// EnterFull_describe_statement is called when entering the full_describe_statement production.
	EnterFull_describe_statement(c *Full_describe_statementContext)

	// EnterHelp_statement is called when entering the help_statement production.
	EnterHelp_statement(c *Help_statementContext)

	// EnterUse_statement is called when entering the use_statement production.
	EnterUse_statement(c *Use_statementContext)

	// EnterDescstmtDescObj is called when entering the descstmtDescObj production.
	EnterDescstmtDescObj(c *DescstmtDescObjContext)

	// EnterConnectionDescObj is called when entering the connectionDescObj production.
	EnterConnectionDescObj(c *ConnectionDescObjContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterFull_id is called when entering the full_id production.
	EnterFull_id(c *Full_idContext)

	// EnterFull_column_name is called when entering the full_column_name production.
	EnterFull_column_name(c *Full_column_nameContext)

	// EnterIndex_col_name is called when entering the index_col_name production.
	EnterIndex_col_name(c *Index_col_nameContext)

	// EnterUser_name is called when entering the user_name production.
	EnterUser_name(c *User_nameContext)

	// EnterMysql_variable is called when entering the mysql_variable production.
	EnterMysql_variable(c *Mysql_variableContext)

	// EnterCharset_name is called when entering the charset_name production.
	EnterCharset_name(c *Charset_nameContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterEngine_name is called when entering the engine_name production.
	EnterEngine_name(c *Engine_nameContext)

	// EnterUuid_set is called when entering the uuid_set production.
	EnterUuid_set(c *Uuid_setContext)

	// EnterXid is called when entering the xid production.
	EnterXid(c *XidContext)

	// EnterXid_string_id is called when entering the xid_string_id production.
	EnterXid_string_id(c *Xid_string_idContext)

	// EnterAuth_plugin is called when entering the auth_plugin production.
	EnterAuth_plugin(c *Auth_pluginContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// EnterSimple_id is called when entering the simple_id production.
	EnterSimple_id(c *Simple_idContext)

	// EnterDot_ext_id is called when entering the dot_ext_id production.
	EnterDot_ext_id(c *Dot_ext_idContext)

	// EnterDecimal_literal is called when entering the decimal_literal production.
	EnterDecimal_literal(c *Decimal_literalContext)

	// EnterFilesize_literal is called when entering the filesize_literal production.
	EnterFilesize_literal(c *Filesize_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterHexadecimal_literal is called when entering the hexadecimal_literal production.
	EnterHexadecimal_literal(c *Hexadecimal_literalContext)

	// EnterNull_notnull is called when entering the null_notnull production.
	EnterNull_notnull(c *Null_notnullContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterCharDatatype is called when entering the charDatatype production.
	EnterCharDatatype(c *CharDatatypeContext)

	// EnterDimensionDatatype is called when entering the dimensionDatatype production.
	EnterDimensionDatatype(c *DimensionDatatypeContext)

	// EnterSimpleDatatype is called when entering the simpleDatatype production.
	EnterSimpleDatatype(c *SimpleDatatypeContext)

	// EnterCollectCharDatatype is called when entering the collectCharDatatype production.
	EnterCollectCharDatatype(c *CollectCharDatatypeContext)

	// EnterSpatialDatatype is called when entering the spatialDatatype production.
	EnterSpatialDatatype(c *SpatialDatatypeContext)

	// EnterData_type_to_convert is called when entering the data_type_to_convert production.
	EnterData_type_to_convert(c *Data_type_to_convertContext)

	// EnterSpatial_data_type is called when entering the spatial_data_type production.
	EnterSpatial_data_type(c *Spatial_data_typeContext)

	// EnterLength_one_dimension is called when entering the length_one_dimension production.
	EnterLength_one_dimension(c *Length_one_dimensionContext)

	// EnterLength_two_dimension is called when entering the length_two_dimension production.
	EnterLength_two_dimension(c *Length_two_dimensionContext)

	// EnterLength_two_optional_dimension is called when entering the length_two_optional_dimension production.
	EnterLength_two_optional_dimension(c *Length_two_optional_dimensionContext)

	// EnterId_list is called when entering the id_list production.
	EnterId_list(c *Id_listContext)

	// EnterTable_list is called when entering the table_list production.
	EnterTable_list(c *Table_listContext)

	// EnterTable_pair_list is called when entering the table_pair_list production.
	EnterTable_pair_list(c *Table_pair_listContext)

	// EnterIndex_colname_list is called when entering the index_colname_list production.
	EnterIndex_colname_list(c *Index_colname_listContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterConstant_list is called when entering the constant_list production.
	EnterConstant_list(c *Constant_listContext)

	// EnterSimple_string_list is called when entering the simple_string_list production.
	EnterSimple_string_list(c *Simple_string_listContext)

	// EnterUser_var_list is called when entering the user_var_list production.
	EnterUser_var_list(c *User_var_listContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterIf_exists is called when entering the if_exists production.
	EnterIf_exists(c *If_existsContext)

	// EnterIf_not_exists is called when entering the if_not_exists production.
	EnterIf_not_exists(c *If_not_existsContext)

	// EnterSpecificFunctionCall is called when entering the specificFunctionCall production.
	EnterSpecificFunctionCall(c *SpecificFunctionCallContext)

	// EnterAggregateFunctionCall is called when entering the aggregateFunctionCall production.
	EnterAggregateFunctionCall(c *AggregateFunctionCallContext)

	// EnterScalarFunctionCall is called when entering the scalarFunctionCall production.
	EnterScalarFunctionCall(c *ScalarFunctionCallContext)

	// EnterUdfFunctionCall is called when entering the udfFunctionCall production.
	EnterUdfFunctionCall(c *UdfFunctionCallContext)

	// EnterSimpleSpecificFCall is called when entering the simpleSpecificFCall production.
	EnterSimpleSpecificFCall(c *SimpleSpecificFCallContext)

	// EnterConvertDataTypeFCall is called when entering the convertDataTypeFCall production.
	EnterConvertDataTypeFCall(c *ConvertDataTypeFCallContext)

	// EnterValuesFCall is called when entering the valuesFCall production.
	EnterValuesFCall(c *ValuesFCallContext)

	// EnterCaseFCall is called when entering the caseFCall production.
	EnterCaseFCall(c *CaseFCallContext)

	// EnterCharFCall is called when entering the charFCall production.
	EnterCharFCall(c *CharFCallContext)

	// EnterPositionFCall is called when entering the positionFCall production.
	EnterPositionFCall(c *PositionFCallContext)

	// EnterSubstrFCall is called when entering the substrFCall production.
	EnterSubstrFCall(c *SubstrFCallContext)

	// EnterTrimFCall is called when entering the trimFCall production.
	EnterTrimFCall(c *TrimFCallContext)

	// EnterWeightFCall is called when entering the weightFCall production.
	EnterWeightFCall(c *WeightFCallContext)

	// EnterExtractFCall is called when entering the extractFCall production.
	EnterExtractFCall(c *ExtractFCallContext)

	// EnterGetFormatFCall is called when entering the getFormatFCall production.
	EnterGetFormatFCall(c *GetFormatFCallContext)

	// EnterLevelWeightFList is called when entering the levelWeightFList production.
	EnterLevelWeightFList(c *LevelWeightFListContext)

	// EnterLevelWeightFRange is called when entering the levelWeightFRange production.
	EnterLevelWeightFRange(c *LevelWeightFRangeContext)

	// EnterAggregate_windowed_function is called when entering the aggregate_windowed_function production.
	EnterAggregate_windowed_function(c *Aggregate_windowed_functionContext)

	// EnterScalar_function_name is called when entering the scalar_function_name production.
	EnterScalar_function_name(c *Scalar_function_nameContext)

	// EnterFunction_args is called when entering the function_args production.
	EnterFunction_args(c *Function_argsContext)

	// EnterFunction_arg is called when entering the function_arg production.
	EnterFunction_arg(c *Function_argContext)

	// EnterIsExpression is called when entering the isExpression production.
	EnterIsExpression(c *IsExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterSoundsLikePredicate is called when entering the soundsLikePredicate production.
	EnterSoundsLikePredicate(c *SoundsLikePredicateContext)

	// EnterExpressionAtomPredicate is called when entering the expressionAtomPredicate production.
	EnterExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterSubqueryComparasionPredicate is called when entering the subqueryComparasionPredicate production.
	EnterSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterBinaryComparasionPredicate is called when entering the binaryComparasionPredicate production.
	EnterBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterRegexpPredicate is called when entering the regexpPredicate production.
	EnterRegexpPredicate(c *RegexpPredicateContext)

	// EnterUnaryExpressionAtom is called when entering the unaryExpressionAtom production.
	EnterUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// EnterExistsExpessionAtom is called when entering the existsExpessionAtom production.
	EnterExistsExpessionAtom(c *ExistsExpessionAtomContext)

	// EnterConstantExpressionAtom is called when entering the constantExpressionAtom production.
	EnterConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// EnterFunctionCallExpressionAtom is called when entering the functionCallExpressionAtom production.
	EnterFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// EnterMysqlVariableExpressionAtom is called when entering the mysqlVariableExpressionAtom production.
	EnterMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// EnterBinaryExpressionAtom is called when entering the binaryExpressionAtom production.
	EnterBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// EnterFullColumnNameExpressionAtom is called when entering the fullColumnNameExpressionAtom production.
	EnterFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// EnterDefaultExpressionAtom is called when entering the defaultExpressionAtom production.
	EnterDefaultExpressionAtom(c *DefaultExpressionAtomContext)

	// EnterBitExpressionAtom is called when entering the bitExpressionAtom production.
	EnterBitExpressionAtom(c *BitExpressionAtomContext)

	// EnterNestedExpressionAtom is called when entering the nestedExpressionAtom production.
	EnterNestedExpressionAtom(c *NestedExpressionAtomContext)

	// EnterMathExpressionAtom is called when entering the mathExpressionAtom production.
	EnterMathExpressionAtom(c *MathExpressionAtomContext)

	// EnterIntervalExpressionAtom is called when entering the intervalExpressionAtom production.
	EnterIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterLogical_operator is called when entering the logical_operator production.
	EnterLogical_operator(c *Logical_operatorContext)

	// EnterBit_operator is called when entering the bit_operator production.
	EnterBit_operator(c *Bit_operatorContext)

	// EnterMath_operator is called when entering the math_operator production.
	EnterMath_operator(c *Math_operatorContext)

	// EnterCharset_name_base is called when entering the charset_name_base production.
	EnterCharset_name_base(c *Charset_name_baseContext)

	// EnterTransaction_level_base is called when entering the transaction_level_base production.
	EnterTransaction_level_base(c *Transaction_level_baseContext)

	// EnterPrivileges_base is called when entering the privileges_base production.
	EnterPrivileges_base(c *Privileges_baseContext)

	// EnterInterval_type_base is called when entering the interval_type_base production.
	EnterInterval_type_base(c *Interval_type_baseContext)

	// EnterData_type_base is called when entering the data_type_base production.
	EnterData_type_base(c *Data_type_baseContext)

	// EnterKeywords_can_be_id is called when entering the keywords_can_be_id production.
	EnterKeywords_can_be_id(c *Keywords_can_be_idContext)

	// EnterFunction_name_base is called when entering the function_name_base production.
	EnterFunction_name_base(c *Function_name_baseContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSql_statements is called when exiting the sql_statements production.
	ExitSql_statements(c *Sql_statementsContext)

	// ExitSql_statement is called when exiting the sql_statement production.
	ExitSql_statement(c *Sql_statementContext)

	// ExitEmpty_statement is called when exiting the empty_statement production.
	ExitEmpty_statement(c *Empty_statementContext)

	// ExitDdl_statement is called when exiting the ddl_statement production.
	ExitDdl_statement(c *Ddl_statementContext)

	// ExitDml_statement is called when exiting the dml_statement production.
	ExitDml_statement(c *Dml_statementContext)

	// ExitTransaction_statement is called when exiting the transaction_statement production.
	ExitTransaction_statement(c *Transaction_statementContext)

	// ExitReplication_statement is called when exiting the replication_statement production.
	ExitReplication_statement(c *Replication_statementContext)

	// ExitPrepared_statement is called when exiting the prepared_statement production.
	ExitPrepared_statement(c *Prepared_statementContext)

	// ExitCompound_statement is called when exiting the compound_statement production.
	ExitCompound_statement(c *Compound_statementContext)

	// ExitAdministration_statement is called when exiting the administration_statement production.
	ExitAdministration_statement(c *Administration_statementContext)

	// ExitUtility_statement is called when exiting the utility_statement production.
	ExitUtility_statement(c *Utility_statementContext)

	// ExitCreate_database is called when exiting the create_database production.
	ExitCreate_database(c *Create_databaseContext)

	// ExitCreate_event is called when exiting the create_event production.
	ExitCreate_event(c *Create_eventContext)

	// ExitCreate_index is called when exiting the create_index production.
	ExitCreate_index(c *Create_indexContext)

	// ExitCreate_logfile_group is called when exiting the create_logfile_group production.
	ExitCreate_logfile_group(c *Create_logfile_groupContext)

	// ExitCreate_procedure is called when exiting the create_procedure production.
	ExitCreate_procedure(c *Create_procedureContext)

	// ExitCreate_function is called when exiting the create_function production.
	ExitCreate_function(c *Create_functionContext)

	// ExitCreate_server is called when exiting the create_server production.
	ExitCreate_server(c *Create_serverContext)

	// ExitCopyCreateTable is called when exiting the copyCreateTable production.
	ExitCopyCreateTable(c *CopyCreateTableContext)

	// ExitQueryCreateTable is called when exiting the queryCreateTable production.
	ExitQueryCreateTable(c *QueryCreateTableContext)

	// ExitColCreateTable is called when exiting the colCreateTable production.
	ExitColCreateTable(c *ColCreateTableContext)

	// ExitCreate_tablespace_innodb is called when exiting the create_tablespace_innodb production.
	ExitCreate_tablespace_innodb(c *Create_tablespace_innodbContext)

	// ExitCreate_tablespace_ndb is called when exiting the create_tablespace_ndb production.
	ExitCreate_tablespace_ndb(c *Create_tablespace_ndbContext)

	// ExitCreate_trigger is called when exiting the create_trigger production.
	ExitCreate_trigger(c *Create_triggerContext)

	// ExitCreate_view is called when exiting the create_view production.
	ExitCreate_view(c *Create_viewContext)

	// ExitCreate_database_option is called when exiting the create_database_option production.
	ExitCreate_database_option(c *Create_database_optionContext)

	// ExitOwner_statement is called when exiting the owner_statement production.
	ExitOwner_statement(c *Owner_statementContext)

	// ExitPreciseSchedule is called when exiting the preciseSchedule production.
	ExitPreciseSchedule(c *PreciseScheduleContext)

	// ExitIntervalSchedule is called when exiting the intervalSchedule production.
	ExitIntervalSchedule(c *IntervalScheduleContext)

	// ExitTimestamp_value is called when exiting the timestamp_value production.
	ExitTimestamp_value(c *Timestamp_valueContext)

	// ExitInterval_expr is called when exiting the interval_expr production.
	ExitInterval_expr(c *Interval_exprContext)

	// ExitInterval_type is called when exiting the interval_type production.
	ExitInterval_type(c *Interval_typeContext)

	// ExitIndex_type is called when exiting the index_type production.
	ExitIndex_type(c *Index_typeContext)

	// ExitIndex_option is called when exiting the index_option production.
	ExitIndex_option(c *Index_optionContext)

	// ExitProc_param is called when exiting the proc_param production.
	ExitProc_param(c *Proc_paramContext)

	// ExitFunc_param is called when exiting the func_param production.
	ExitFunc_param(c *Func_paramContext)

	// ExitRcComment is called when exiting the rcComment production.
	ExitRcComment(c *RcCommentContext)

	// ExitRcSqllang is called when exiting the rcSqllang production.
	ExitRcSqllang(c *RcSqllangContext)

	// ExitRcDeterm is called when exiting the rcDeterm production.
	ExitRcDeterm(c *RcDetermContext)

	// ExitRcSqldata is called when exiting the rcSqldata production.
	ExitRcSqldata(c *RcSqldataContext)

	// ExitRcSecurestmt is called when exiting the rcSecurestmt production.
	ExitRcSecurestmt(c *RcSecurestmtContext)

	// ExitServer_option is called when exiting the server_option production.
	ExitServer_option(c *Server_optionContext)

	// ExitColumn_def_table_constraints is called when exiting the column_def_table_constraints production.
	ExitColumn_def_table_constraints(c *Column_def_table_constraintsContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitConstraintDefinition is called when exiting the constraintDefinition production.
	ExitConstraintDefinition(c *ConstraintDefinitionContext)

	// ExitIndexDefinition is called when exiting the indexDefinition production.
	ExitIndexDefinition(c *IndexDefinitionContext)

	// ExitColumn_definition is called when exiting the column_definition production.
	ExitColumn_definition(c *Column_definitionContext)

	// ExitColConstrNull is called when exiting the colConstrNull production.
	ExitColConstrNull(c *ColConstrNullContext)

	// ExitColConstrDflt is called when exiting the colConstrDflt production.
	ExitColConstrDflt(c *ColConstrDfltContext)

	// ExitColConstrAuInc is called when exiting the colConstrAuInc production.
	ExitColConstrAuInc(c *ColConstrAuIncContext)

	// ExitColConstrPK is called when exiting the colConstrPK production.
	ExitColConstrPK(c *ColConstrPKContext)

	// ExitColConstrUK is called when exiting the colConstrUK production.
	ExitColConstrUK(c *ColConstrUKContext)

	// ExitColConstrComment is called when exiting the colConstrComment production.
	ExitColConstrComment(c *ColConstrCommentContext)

	// ExitColConstrForm is called when exiting the colConstrForm production.
	ExitColConstrForm(c *ColConstrFormContext)

	// ExitColConstrStorage is called when exiting the colConstrStorage production.
	ExitColConstrStorage(c *ColConstrStorageContext)

	// ExitColConstrRefdef is called when exiting the colConstrRefdef production.
	ExitColConstrRefdef(c *ColConstrRefdefContext)

	// ExitTblConstrPK is called when exiting the tblConstrPK production.
	ExitTblConstrPK(c *TblConstrPKContext)

	// ExitTblConstrUK is called when exiting the tblConstrUK production.
	ExitTblConstrUK(c *TblConstrUKContext)

	// ExitTblConstrFK is called when exiting the tblConstrFK production.
	ExitTblConstrFK(c *TblConstrFKContext)

	// ExitTblConstCheck is called when exiting the tblConstCheck production.
	ExitTblConstCheck(c *TblConstCheckContext)

	// ExitReference_definition is called when exiting the reference_definition production.
	ExitReference_definition(c *Reference_definitionContext)

	// ExitOn_delete_action is called when exiting the on_delete_action production.
	ExitOn_delete_action(c *On_delete_actionContext)

	// ExitOn_update_action is called when exiting the on_update_action production.
	ExitOn_update_action(c *On_update_actionContext)

	// ExitReference_action_control_type is called when exiting the reference_action_control_type production.
	ExitReference_action_control_type(c *Reference_action_control_typeContext)

	// ExitSimpleIndex is called when exiting the simpleIndex production.
	ExitSimpleIndex(c *SimpleIndexContext)

	// ExitSpecIndex is called when exiting the specIndex production.
	ExitSpecIndex(c *SpecIndexContext)

	// ExitTblOptEngine is called when exiting the tblOptEngine production.
	ExitTblOptEngine(c *TblOptEngineContext)

	// ExitTblOptAuInc is called when exiting the tblOptAuInc production.
	ExitTblOptAuInc(c *TblOptAuIncContext)

	// ExitTblOptAvgRLen is called when exiting the tblOptAvgRLen production.
	ExitTblOptAvgRLen(c *TblOptAvgRLenContext)

	// ExitTblOptDefCharSet is called when exiting the tblOptDefCharSet production.
	ExitTblOptDefCharSet(c *TblOptDefCharSetContext)

	// ExitTblOptChkSum is called when exiting the tblOptChkSum production.
	ExitTblOptChkSum(c *TblOptChkSumContext)

	// ExitTblOptDefCollate is called when exiting the tblOptDefCollate production.
	ExitTblOptDefCollate(c *TblOptDefCollateContext)

	// ExitTblOptComment is called when exiting the tblOptComment production.
	ExitTblOptComment(c *TblOptCommentContext)

	// ExitTblOptCompr is called when exiting the tblOptCompr production.
	ExitTblOptCompr(c *TblOptComprContext)

	// ExitTblOptConn is called when exiting the tblOptConn production.
	ExitTblOptConn(c *TblOptConnContext)

	// ExitTblOptDataDir is called when exiting the tblOptDataDir production.
	ExitTblOptDataDir(c *TblOptDataDirContext)

	// ExitTblOptDelKW is called when exiting the tblOptDelKW production.
	ExitTblOptDelKW(c *TblOptDelKWContext)

	// ExitTblOptEncr is called when exiting the tblOptEncr production.
	ExitTblOptEncr(c *TblOptEncrContext)

	// ExitTblOptIndexDir is called when exiting the tblOptIndexDir production.
	ExitTblOptIndexDir(c *TblOptIndexDirContext)

	// ExitTblOptInsMeth is called when exiting the tblOptInsMeth production.
	ExitTblOptInsMeth(c *TblOptInsMethContext)

	// ExitTblOptKeyBlockSz is called when exiting the tblOptKeyBlockSz production.
	ExitTblOptKeyBlockSz(c *TblOptKeyBlockSzContext)

	// ExitTblOptMaxRows is called when exiting the tblOptMaxRows production.
	ExitTblOptMaxRows(c *TblOptMaxRowsContext)

	// ExitTblOptMinRows is called when exiting the tblOptMinRows production.
	ExitTblOptMinRows(c *TblOptMinRowsContext)

	// ExitTblOptPackK is called when exiting the tblOptPackK production.
	ExitTblOptPackK(c *TblOptPackKContext)

	// ExitTblOptPasswd is called when exiting the tblOptPasswd production.
	ExitTblOptPasswd(c *TblOptPasswdContext)

	// ExitTblOptRowFormat is called when exiting the tblOptRowFormat production.
	ExitTblOptRowFormat(c *TblOptRowFormatContext)

	// ExitTblOptStatAutoR is called when exiting the tblOptStatAutoR production.
	ExitTblOptStatAutoR(c *TblOptStatAutoRContext)

	// ExitTblOptStatPersist is called when exiting the tblOptStatPersist production.
	ExitTblOptStatPersist(c *TblOptStatPersistContext)

	// ExitTblOptStatSamplPg is called when exiting the tblOptStatSamplPg production.
	ExitTblOptStatSamplPg(c *TblOptStatSamplPgContext)

	// ExitTblOptTablespace is called when exiting the tblOptTablespace production.
	ExitTblOptTablespace(c *TblOptTablespaceContext)

	// ExitTblOptUnion is called when exiting the tblOptUnion production.
	ExitTblOptUnion(c *TblOptUnionContext)

	// ExitPartition_options is called when exiting the partition_options production.
	ExitPartition_options(c *Partition_optionsContext)

	// ExitPartition_function_definition is called when exiting the partition_function_definition production.
	ExitPartition_function_definition(c *Partition_function_definitionContext)

	// ExitLinear_partition_func_def is called when exiting the linear_partition_func_def production.
	ExitLinear_partition_func_def(c *Linear_partition_func_defContext)

	// ExitPartition_def is called when exiting the partition_def production.
	ExitPartition_def(c *Partition_defContext)

	// ExitSubpartition_def is called when exiting the subpartition_def production.
	ExitSubpartition_def(c *Subpartition_defContext)

	// ExitAlterDb is called when exiting the alterDb production.
	ExitAlterDb(c *AlterDbContext)

	// ExitAlterDbUpgradeName is called when exiting the alterDbUpgradeName production.
	ExitAlterDbUpgradeName(c *AlterDbUpgradeNameContext)

	// ExitAlter_event is called when exiting the alter_event production.
	ExitAlter_event(c *Alter_eventContext)

	// ExitAlter_function is called when exiting the alter_function production.
	ExitAlter_function(c *Alter_functionContext)

	// ExitAlter_instance is called when exiting the alter_instance production.
	ExitAlter_instance(c *Alter_instanceContext)

	// ExitAlter_logfile_group is called when exiting the alter_logfile_group production.
	ExitAlter_logfile_group(c *Alter_logfile_groupContext)

	// ExitAlter_procedure is called when exiting the alter_procedure production.
	ExitAlter_procedure(c *Alter_procedureContext)

	// ExitAlter_server is called when exiting the alter_server production.
	ExitAlter_server(c *Alter_serverContext)

	// ExitAlter_table is called when exiting the alter_table production.
	ExitAlter_table(c *Alter_tableContext)

	// ExitAlter_tablespace is called when exiting the alter_tablespace production.
	ExitAlter_tablespace(c *Alter_tablespaceContext)

	// ExitAlter_view is called when exiting the alter_view production.
	ExitAlter_view(c *Alter_viewContext)

	// ExitAltblTableOpt is called when exiting the altblTableOpt production.
	ExitAltblTableOpt(c *AltblTableOptContext)

	// ExitAltblAddCol is called when exiting the altblAddCol production.
	ExitAltblAddCol(c *AltblAddColContext)

	// ExitAltblAddCols is called when exiting the altblAddCols production.
	ExitAltblAddCols(c *AltblAddColsContext)

	// ExitAltblAddIndex is called when exiting the altblAddIndex production.
	ExitAltblAddIndex(c *AltblAddIndexContext)

	// ExitAltblAddPK is called when exiting the altblAddPK production.
	ExitAltblAddPK(c *AltblAddPKContext)

	// ExitAltblAddUK is called when exiting the altblAddUK production.
	ExitAltblAddUK(c *AltblAddUKContext)

	// ExitAltblAddSpecIndex is called when exiting the altblAddSpecIndex production.
	ExitAltblAddSpecIndex(c *AltblAddSpecIndexContext)

	// ExitAltblAddFK is called when exiting the altblAddFK production.
	ExitAltblAddFK(c *AltblAddFKContext)

	// ExitAltblAlg is called when exiting the altblAlg production.
	ExitAltblAlg(c *AltblAlgContext)

	// ExitAltblColDef is called when exiting the altblColDef production.
	ExitAltblColDef(c *AltblColDefContext)

	// ExitAltblColChange is called when exiting the altblColChange production.
	ExitAltblColChange(c *AltblColChangeContext)

	// ExitAltblLock is called when exiting the altblLock production.
	ExitAltblLock(c *AltblLockContext)

	// ExitAltblColMod is called when exiting the altblColMod production.
	ExitAltblColMod(c *AltblColModContext)

	// ExitAltblColDrop is called when exiting the altblColDrop production.
	ExitAltblColDrop(c *AltblColDropContext)

	// ExitAltblDropPK is called when exiting the altblDropPK production.
	ExitAltblDropPK(c *AltblDropPKContext)

	// ExitAltblDropIndex is called when exiting the altblDropIndex production.
	ExitAltblDropIndex(c *AltblDropIndexContext)

	// ExitAltblDropFK is called when exiting the altblDropFK production.
	ExitAltblDropFK(c *AltblDropFKContext)

	// ExitAltblDisKey is called when exiting the altblDisKey production.
	ExitAltblDisKey(c *AltblDisKeyContext)

	// ExitAltblEnKey is called when exiting the altblEnKey production.
	ExitAltblEnKey(c *AltblEnKeyContext)

	// ExitAltblRenameTbl is called when exiting the altblRenameTbl production.
	ExitAltblRenameTbl(c *AltblRenameTblContext)

	// ExitAltblResort is called when exiting the altblResort production.
	ExitAltblResort(c *AltblResortContext)

	// ExitAltblConvert is called when exiting the altblConvert production.
	ExitAltblConvert(c *AltblConvertContext)

	// ExitAltblDefCharset is called when exiting the altblDefCharset production.
	ExitAltblDefCharset(c *AltblDefCharsetContext)

	// ExitAltblDisTblspace is called when exiting the altblDisTblspace production.
	ExitAltblDisTblspace(c *AltblDisTblspaceContext)

	// ExitAltblImpTblSpace is called when exiting the altblImpTblSpace production.
	ExitAltblImpTblSpace(c *AltblImpTblSpaceContext)

	// ExitAltblForce is called when exiting the altblForce production.
	ExitAltblForce(c *AltblForceContext)

	// ExitAltblValid is called when exiting the altblValid production.
	ExitAltblValid(c *AltblValidContext)

	// ExitAltblAddPart is called when exiting the altblAddPart production.
	ExitAltblAddPart(c *AltblAddPartContext)

	// ExitAltblDropPart is called when exiting the altblDropPart production.
	ExitAltblDropPart(c *AltblDropPartContext)

	// ExitAltblDiscartPart is called when exiting the altblDiscartPart production.
	ExitAltblDiscartPart(c *AltblDiscartPartContext)

	// ExitAltblImportPart is called when exiting the altblImportPart production.
	ExitAltblImportPart(c *AltblImportPartContext)

	// ExitAltblTruncPart is called when exiting the altblTruncPart production.
	ExitAltblTruncPart(c *AltblTruncPartContext)

	// ExitAltblCoalPart is called when exiting the altblCoalPart production.
	ExitAltblCoalPart(c *AltblCoalPartContext)

	// ExitAltblReorgPart is called when exiting the altblReorgPart production.
	ExitAltblReorgPart(c *AltblReorgPartContext)

	// ExitAltblExchPart is called when exiting the altblExchPart production.
	ExitAltblExchPart(c *AltblExchPartContext)

	// ExitAltblAnalPart is called when exiting the altblAnalPart production.
	ExitAltblAnalPart(c *AltblAnalPartContext)

	// ExitAltblCheckPart is called when exiting the altblCheckPart production.
	ExitAltblCheckPart(c *AltblCheckPartContext)

	// ExitAltblOptimPart is called when exiting the altblOptimPart production.
	ExitAltblOptimPart(c *AltblOptimPartContext)

	// ExitAltblRebuildPart is called when exiting the altblRebuildPart production.
	ExitAltblRebuildPart(c *AltblRebuildPartContext)

	// ExitAltblRepairPart is called when exiting the altblRepairPart production.
	ExitAltblRepairPart(c *AltblRepairPartContext)

	// ExitAltblRemovePart is called when exiting the altblRemovePart production.
	ExitAltblRemovePart(c *AltblRemovePartContext)

	// ExitAltblUpgrPart is called when exiting the altblUpgrPart production.
	ExitAltblUpgrPart(c *AltblUpgrPartContext)

	// ExitDrop_database is called when exiting the drop_database production.
	ExitDrop_database(c *Drop_databaseContext)

	// ExitDrop_event is called when exiting the drop_event production.
	ExitDrop_event(c *Drop_eventContext)

	// ExitDrop_index is called when exiting the drop_index production.
	ExitDrop_index(c *Drop_indexContext)

	// ExitDrop_logfile_group is called when exiting the drop_logfile_group production.
	ExitDrop_logfile_group(c *Drop_logfile_groupContext)

	// ExitDrop_procedure is called when exiting the drop_procedure production.
	ExitDrop_procedure(c *Drop_procedureContext)

	// ExitDrop_function is called when exiting the drop_function production.
	ExitDrop_function(c *Drop_functionContext)

	// ExitDrop_server is called when exiting the drop_server production.
	ExitDrop_server(c *Drop_serverContext)

	// ExitDrop_table is called when exiting the drop_table production.
	ExitDrop_table(c *Drop_tableContext)

	// ExitDrop_tablespace is called when exiting the drop_tablespace production.
	ExitDrop_tablespace(c *Drop_tablespaceContext)

	// ExitDrop_trigger is called when exiting the drop_trigger production.
	ExitDrop_trigger(c *Drop_triggerContext)

	// ExitDrop_view is called when exiting the drop_view production.
	ExitDrop_view(c *Drop_viewContext)

	// ExitRename_table is called when exiting the rename_table production.
	ExitRename_table(c *Rename_tableContext)

	// ExitTruncate_table is called when exiting the truncate_table production.
	ExitTruncate_table(c *Truncate_tableContext)

	// ExitCall_statement is called when exiting the call_statement production.
	ExitCall_statement(c *Call_statementContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitDo_statement is called when exiting the do_statement production.
	ExitDo_statement(c *Do_statementContext)

	// ExitHandler_statement is called when exiting the handler_statement production.
	ExitHandler_statement(c *Handler_statementContext)

	// ExitInsert_statement is called when exiting the insert_statement production.
	ExitInsert_statement(c *Insert_statementContext)

	// ExitLoad_data_statement is called when exiting the load_data_statement production.
	ExitLoad_data_statement(c *Load_data_statementContext)

	// ExitLoad_xml_statement is called when exiting the load_xml_statement production.
	ExitLoad_xml_statement(c *Load_xml_statementContext)

	// ExitReplace_statement is called when exiting the replace_statement production.
	ExitReplace_statement(c *Replace_statementContext)

	// ExitSimpleSelect is called when exiting the simpleSelect production.
	ExitSimpleSelect(c *SimpleSelectContext)

	// ExitParenSelect is called when exiting the parenSelect production.
	ExitParenSelect(c *ParenSelectContext)

	// ExitUnionSelect is called when exiting the unionSelect production.
	ExitUnionSelect(c *UnionSelectContext)

	// ExitUnionParenSelect is called when exiting the unionParenSelect production.
	ExitUnionParenSelect(c *UnionParenSelectContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitInsert_statement_value is called when exiting the insert_statement_value production.
	ExitInsert_statement_value(c *Insert_statement_valueContext)

	// ExitUpdate_elem is called when exiting the update_elem production.
	ExitUpdate_elem(c *Update_elemContext)

	// ExitCol_or_uservar is called when exiting the col_or_uservar production.
	ExitCol_or_uservar(c *Col_or_uservarContext)

	// ExitSingle_delete_statement is called when exiting the single_delete_statement production.
	ExitSingle_delete_statement(c *Single_delete_statementContext)

	// ExitMultiple_delete_statement is called when exiting the multiple_delete_statement production.
	ExitMultiple_delete_statement(c *Multiple_delete_statementContext)

	// ExitHandler_open_statement is called when exiting the handler_open_statement production.
	ExitHandler_open_statement(c *Handler_open_statementContext)

	// ExitHandler_read_index_statement is called when exiting the handler_read_index_statement production.
	ExitHandler_read_index_statement(c *Handler_read_index_statementContext)

	// ExitHandler_read_statement is called when exiting the handler_read_statement production.
	ExitHandler_read_statement(c *Handler_read_statementContext)

	// ExitHandler_close_statement is called when exiting the handler_close_statement production.
	ExitHandler_close_statement(c *Handler_close_statementContext)

	// ExitSingle_update_statement is called when exiting the single_update_statement production.
	ExitSingle_update_statement(c *Single_update_statementContext)

	// ExitMultiple_update_statement is called when exiting the multiple_update_statement production.
	ExitMultiple_update_statement(c *Multiple_update_statementContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitOrder_by_expression is called when exiting the order_by_expression production.
	ExitOrder_by_expression(c *Order_by_expressionContext)

	// ExitTable_sources is called when exiting the table_sources production.
	ExitTable_sources(c *Table_sourcesContext)

	// ExitTable_source is called when exiting the table_source production.
	ExitTable_source(c *Table_sourceContext)

	// ExitAtomTableItem is called when exiting the atomTableItem production.
	ExitAtomTableItem(c *AtomTableItemContext)

	// ExitSubqueryTableItem is called when exiting the subqueryTableItem production.
	ExitSubqueryTableItem(c *SubqueryTableItemContext)

	// ExitTableSourcesItem is called when exiting the tableSourcesItem production.
	ExitTableSourcesItem(c *TableSourcesItemContext)

	// ExitIndex_hint is called when exiting the index_hint production.
	ExitIndex_hint(c *Index_hintContext)

	// ExitInnerJoin is called when exiting the innerJoin production.
	ExitInnerJoin(c *InnerJoinContext)

	// ExitStraightJoin is called when exiting the straightJoin production.
	ExitStraightJoin(c *StraightJoinContext)

	// ExitOuterJoin is called when exiting the outerJoin production.
	ExitOuterJoin(c *OuterJoinContext)

	// ExitNaturalJoin is called when exiting the naturalJoin production.
	ExitNaturalJoin(c *NaturalJoinContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitQuery_expression is called when exiting the query_expression production.
	ExitQuery_expression(c *Query_expressionContext)

	// ExitQuery_expression_nointo is called when exiting the query_expression_nointo production.
	ExitQuery_expression_nointo(c *Query_expression_nointoContext)

	// ExitQuery_specification is called when exiting the query_specification production.
	ExitQuery_specification(c *Query_specificationContext)

	// ExitQuery_specification_nointo is called when exiting the query_specification_nointo production.
	ExitQuery_specification_nointo(c *Query_specification_nointoContext)

	// ExitUnion_parenth is called when exiting the union_parenth production.
	ExitUnion_parenth(c *Union_parenthContext)

	// ExitUnion_statement is called when exiting the union_statement production.
	ExitUnion_statement(c *Union_statementContext)

	// ExitSelect_spec is called when exiting the select_spec production.
	ExitSelect_spec(c *Select_specContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSellistelAllCol is called when exiting the sellistelAllCol production.
	ExitSellistelAllCol(c *SellistelAllColContext)

	// ExitSellistelCol is called when exiting the sellistelCol production.
	ExitSellistelCol(c *SellistelColContext)

	// ExitSellistelFunc is called when exiting the sellistelFunc production.
	ExitSellistelFunc(c *SellistelFuncContext)

	// ExitSellistelExpr is called when exiting the sellistelExpr production.
	ExitSellistelExpr(c *SellistelExprContext)

	// ExitSelectIntoVars is called when exiting the selectIntoVars production.
	ExitSelectIntoVars(c *SelectIntoVarsContext)

	// ExitSelectIntoDump is called when exiting the selectIntoDump production.
	ExitSelectIntoDump(c *SelectIntoDumpContext)

	// ExitSelectIntoOutfile is called when exiting the selectIntoOutfile production.
	ExitSelectIntoOutfile(c *SelectIntoOutfileContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitGroup_by_item is called when exiting the group_by_item production.
	ExitGroup_by_item(c *Group_by_itemContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitStart_transaction is called when exiting the start_transaction production.
	ExitStart_transaction(c *Start_transactionContext)

	// ExitBegin_work is called when exiting the begin_work production.
	ExitBegin_work(c *Begin_workContext)

	// ExitCommit_work is called when exiting the commit_work production.
	ExitCommit_work(c *Commit_workContext)

	// ExitRollback_work is called when exiting the rollback_work production.
	ExitRollback_work(c *Rollback_workContext)

	// ExitSavepoint_statement is called when exiting the savepoint_statement production.
	ExitSavepoint_statement(c *Savepoint_statementContext)

	// ExitRollback_statement is called when exiting the rollback_statement production.
	ExitRollback_statement(c *Rollback_statementContext)

	// ExitRelease_statement is called when exiting the release_statement production.
	ExitRelease_statement(c *Release_statementContext)

	// ExitLock_tables is called when exiting the lock_tables production.
	ExitLock_tables(c *Lock_tablesContext)

	// ExitUnlock_tables is called when exiting the unlock_tables production.
	ExitUnlock_tables(c *Unlock_tablesContext)

	// ExitSet_autocommit_statement is called when exiting the set_autocommit_statement production.
	ExitSet_autocommit_statement(c *Set_autocommit_statementContext)

	// ExitSet_transaction_statement is called when exiting the set_transaction_statement production.
	ExitSet_transaction_statement(c *Set_transaction_statementContext)

	// ExitTransact_option is called when exiting the transact_option production.
	ExitTransact_option(c *Transact_optionContext)

	// ExitLock_table_element is called when exiting the lock_table_element production.
	ExitLock_table_element(c *Lock_table_elementContext)

	// ExitTrans_characteristic is called when exiting the trans_characteristic production.
	ExitTrans_characteristic(c *Trans_characteristicContext)

	// ExitTransaction_level is called when exiting the transaction_level production.
	ExitTransaction_level(c *Transaction_levelContext)

	// ExitChange_master is called when exiting the change_master production.
	ExitChange_master(c *Change_masterContext)

	// ExitChange_repl_filter is called when exiting the change_repl_filter production.
	ExitChange_repl_filter(c *Change_repl_filterContext)

	// ExitPurge_binary_logs is called when exiting the purge_binary_logs production.
	ExitPurge_binary_logs(c *Purge_binary_logsContext)

	// ExitReset_master is called when exiting the reset_master production.
	ExitReset_master(c *Reset_masterContext)

	// ExitReset_slave is called when exiting the reset_slave production.
	ExitReset_slave(c *Reset_slaveContext)

	// ExitStart_slave is called when exiting the start_slave production.
	ExitStart_slave(c *Start_slaveContext)

	// ExitStop_slave is called when exiting the stop_slave production.
	ExitStop_slave(c *Stop_slaveContext)

	// ExitStart_group_repl is called when exiting the start_group_repl production.
	ExitStart_group_repl(c *Start_group_replContext)

	// ExitStop_group_repl is called when exiting the stop_group_repl production.
	ExitStop_group_repl(c *Stop_group_replContext)

	// ExitMasterOptString is called when exiting the masterOptString production.
	ExitMasterOptString(c *MasterOptStringContext)

	// ExitMasterOptDecimal is called when exiting the masterOptDecimal production.
	ExitMasterOptDecimal(c *MasterOptDecimalContext)

	// ExitMasterOptBool is called when exiting the masterOptBool production.
	ExitMasterOptBool(c *MasterOptBoolContext)

	// ExitMasterOptReal is called when exiting the masterOptReal production.
	ExitMasterOptReal(c *MasterOptRealContext)

	// ExitMasterOptIdList is called when exiting the masterOptIdList production.
	ExitMasterOptIdList(c *MasterOptIdListContext)

	// ExitString_master_option is called when exiting the string_master_option production.
	ExitString_master_option(c *String_master_optionContext)

	// ExitDecimal_master_option is called when exiting the decimal_master_option production.
	ExitDecimal_master_option(c *Decimal_master_optionContext)

	// ExitBool_master_option is called when exiting the bool_master_option production.
	ExitBool_master_option(c *Bool_master_optionContext)

	// ExitChannel_option is called when exiting the channel_option production.
	ExitChannel_option(c *Channel_optionContext)

	// ExitReplfilterDbList is called when exiting the replfilterDbList production.
	ExitReplfilterDbList(c *ReplfilterDbListContext)

	// ExitReplfilterTableList is called when exiting the replfilterTableList production.
	ExitReplfilterTableList(c *ReplfilterTableListContext)

	// ExitReplfilterStableList is called when exiting the replfilterStableList production.
	ExitReplfilterStableList(c *ReplfilterStableListContext)

	// ExitReplfilterTablepairList is called when exiting the replfilterTablepairList production.
	ExitReplfilterTablepairList(c *ReplfilterTablepairListContext)

	// ExitThread_type is called when exiting the thread_type production.
	ExitThread_type(c *Thread_typeContext)

	// ExitUntilGtidSset is called when exiting the untilGtidSset production.
	ExitUntilGtidSset(c *UntilGtidSsetContext)

	// ExitUntilMasterLog is called when exiting the untilMasterLog production.
	ExitUntilMasterLog(c *UntilMasterLogContext)

	// ExitUntilRelayLog is called when exiting the untilRelayLog production.
	ExitUntilRelayLog(c *UntilRelayLogContext)

	// ExitUntilSqlGaps is called when exiting the untilSqlGaps production.
	ExitUntilSqlGaps(c *UntilSqlGapsContext)

	// ExitStart_slave_connection_option is called when exiting the start_slave_connection_option production.
	ExitStart_slave_connection_option(c *Start_slave_connection_optionContext)

	// ExitGtid_set is called when exiting the gtid_set production.
	ExitGtid_set(c *Gtid_setContext)

	// ExitXa_start_transaction is called when exiting the xa_start_transaction production.
	ExitXa_start_transaction(c *Xa_start_transactionContext)

	// ExitXa_end_transaction is called when exiting the xa_end_transaction production.
	ExitXa_end_transaction(c *Xa_end_transactionContext)

	// ExitXa_prepare is called when exiting the xa_prepare production.
	ExitXa_prepare(c *Xa_prepareContext)

	// ExitXa_commit_work is called when exiting the xa_commit_work production.
	ExitXa_commit_work(c *Xa_commit_workContext)

	// ExitXa_rollback_work is called when exiting the xa_rollback_work production.
	ExitXa_rollback_work(c *Xa_rollback_workContext)

	// ExitXa_recover_work is called when exiting the xa_recover_work production.
	ExitXa_recover_work(c *Xa_recover_workContext)

	// ExitPrepare_statement is called when exiting the prepare_statement production.
	ExitPrepare_statement(c *Prepare_statementContext)

	// ExitExecute_statement is called when exiting the execute_statement production.
	ExitExecute_statement(c *Execute_statementContext)

	// ExitDeallocate_prepare is called when exiting the deallocate_prepare production.
	ExitDeallocate_prepare(c *Deallocate_prepareContext)

	// ExitRoutine_body is called when exiting the routine_body production.
	ExitRoutine_body(c *Routine_bodyContext)

	// ExitBlock_statement is called when exiting the block_statement production.
	ExitBlock_statement(c *Block_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitIterate_statement is called when exiting the iterate_statement production.
	ExitIterate_statement(c *Iterate_statementContext)

	// ExitLeave_statement is called when exiting the leave_statement production.
	ExitLeave_statement(c *Leave_statementContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitRepeat_statement is called when exiting the repeat_statement production.
	ExitRepeat_statement(c *Repeat_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitCursor_statement is called when exiting the cursor_statement production.
	ExitCursor_statement(c *Cursor_statementContext)

	// ExitDeclare_variable is called when exiting the declare_variable production.
	ExitDeclare_variable(c *Declare_variableContext)

	// ExitDeclare_condition is called when exiting the declare_condition production.
	ExitDeclare_condition(c *Declare_conditionContext)

	// ExitDeclare_cursor is called when exiting the declare_cursor production.
	ExitDeclare_cursor(c *Declare_cursorContext)

	// ExitDeclare_handler is called when exiting the declare_handler production.
	ExitDeclare_handler(c *Declare_handlerContext)

	// ExitHandler_condition_value is called when exiting the handler_condition_value production.
	ExitHandler_condition_value(c *Handler_condition_valueContext)

	// ExitProcedure_sql_statement is called when exiting the procedure_sql_statement production.
	ExitProcedure_sql_statement(c *Procedure_sql_statementContext)

	// ExitAlterUserMysql56 is called when exiting the alterUserMysql56 production.
	ExitAlterUserMysql56(c *AlterUserMysql56Context)

	// ExitAlterUserMysql57 is called when exiting the alterUserMysql57 production.
	ExitAlterUserMysql57(c *AlterUserMysql57Context)

	// ExitCreateUserMysql56 is called when exiting the createUserMysql56 production.
	ExitCreateUserMysql56(c *CreateUserMysql56Context)

	// ExitCreateUserMysql57 is called when exiting the createUserMysql57 production.
	ExitCreateUserMysql57(c *CreateUserMysql57Context)

	// ExitDrop_user is called when exiting the drop_user production.
	ExitDrop_user(c *Drop_userContext)

	// ExitGrant_statement is called when exiting the grant_statement production.
	ExitGrant_statement(c *Grant_statementContext)

	// ExitGrant_proxy is called when exiting the grant_proxy production.
	ExitGrant_proxy(c *Grant_proxyContext)

	// ExitRename_user is called when exiting the rename_user production.
	ExitRename_user(c *Rename_userContext)

	// ExitDetailRevoke is called when exiting the detailRevoke production.
	ExitDetailRevoke(c *DetailRevokeContext)

	// ExitShortRevoke is called when exiting the shortRevoke production.
	ExitShortRevoke(c *ShortRevokeContext)

	// ExitRevoke_proxy is called when exiting the revoke_proxy production.
	ExitRevoke_proxy(c *Revoke_proxyContext)

	// ExitSet_password_statement is called when exiting the set_password_statement production.
	ExitSet_password_statement(c *Set_password_statementContext)

	// ExitUser_password_option is called when exiting the user_password_option production.
	ExitUser_password_option(c *User_password_optionContext)

	// ExitAuthByPassword is called when exiting the authByPassword production.
	ExitAuthByPassword(c *AuthByPasswordContext)

	// ExitAuthByString is called when exiting the authByString production.
	ExitAuthByString(c *AuthByStringContext)

	// ExitAuthByHash is called when exiting the authByHash production.
	ExitAuthByHash(c *AuthByHashContext)

	// ExitTls_option is called when exiting the tls_option production.
	ExitTls_option(c *Tls_optionContext)

	// ExitUser_resource_option is called when exiting the user_resource_option production.
	ExitUser_resource_option(c *User_resource_optionContext)

	// ExitUser_lock_option is called when exiting the user_lock_option production.
	ExitUser_lock_option(c *User_lock_optionContext)

	// ExitPrivelege_clause is called when exiting the privelege_clause production.
	ExitPrivelege_clause(c *Privelege_clauseContext)

	// ExitPrivilege is called when exiting the privilege production.
	ExitPrivilege(c *PrivilegeContext)

	// ExitPrivilege_level is called when exiting the privilege_level production.
	ExitPrivilege_level(c *Privilege_levelContext)

	// ExitSet_password_option is called when exiting the set_password_option production.
	ExitSet_password_option(c *Set_password_optionContext)

	// ExitAnalyze_table is called when exiting the analyze_table production.
	ExitAnalyze_table(c *Analyze_tableContext)

	// ExitCheck_table is called when exiting the check_table production.
	ExitCheck_table(c *Check_tableContext)

	// ExitChecksum_table is called when exiting the checksum_table production.
	ExitChecksum_table(c *Checksum_tableContext)

	// ExitOptimize_table is called when exiting the optimize_table production.
	ExitOptimize_table(c *Optimize_tableContext)

	// ExitRepair_table is called when exiting the repair_table production.
	ExitRepair_table(c *Repair_tableContext)

	// ExitCheck_table_option is called when exiting the check_table_option production.
	ExitCheck_table_option(c *Check_table_optionContext)

	// ExitCreate_udfunction is called when exiting the create_udfunction production.
	ExitCreate_udfunction(c *Create_udfunctionContext)

	// ExitInstall_plugin is called when exiting the install_plugin production.
	ExitInstall_plugin(c *Install_pluginContext)

	// ExitUninstall_plugin is called when exiting the uninstall_plugin production.
	ExitUninstall_plugin(c *Uninstall_pluginContext)

	// ExitSetVariableAssignment is called when exiting the setVariableAssignment production.
	ExitSetVariableAssignment(c *SetVariableAssignmentContext)

	// ExitSetCharset is called when exiting the setCharset production.
	ExitSetCharset(c *SetCharsetContext)

	// ExitSetNames is called when exiting the setNames production.
	ExitSetNames(c *SetNamesContext)

	// ExitSetPasswordStatement is called when exiting the setPasswordStatement production.
	ExitSetPasswordStatement(c *SetPasswordStatementContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitSetAutocommit is called when exiting the setAutocommit production.
	ExitSetAutocommit(c *SetAutocommitContext)

	// ExitShowMasterlogs is called when exiting the showMasterlogs production.
	ExitShowMasterlogs(c *ShowMasterlogsContext)

	// ExitShowLogevents is called when exiting the showLogevents production.
	ExitShowLogevents(c *ShowLogeventsContext)

	// ExitShowObjWithFilter is called when exiting the showObjWithFilter production.
	ExitShowObjWithFilter(c *ShowObjWithFilterContext)

	// ExitShowColumns is called when exiting the showColumns production.
	ExitShowColumns(c *ShowColumnsContext)

	// ExitShowCreateDb is called when exiting the showCreateDb production.
	ExitShowCreateDb(c *ShowCreateDbContext)

	// ExitShowCreateFullidobj is called when exiting the showCreateFullidobj production.
	ExitShowCreateFullidobj(c *ShowCreateFullidobjContext)

	// ExitShowCreateUser is called when exiting the showCreateUser production.
	ExitShowCreateUser(c *ShowCreateUserContext)

	// ExitShowEngine is called when exiting the showEngine production.
	ExitShowEngine(c *ShowEngineContext)

	// ExitShowGlobalinfo is called when exiting the showGlobalinfo production.
	ExitShowGlobalinfo(c *ShowGlobalinfoContext)

	// ExitShowErrWarn is called when exiting the showErrWarn production.
	ExitShowErrWarn(c *ShowErrWarnContext)

	// ExitShowCountErrWarn is called when exiting the showCountErrWarn production.
	ExitShowCountErrWarn(c *ShowCountErrWarnContext)

	// ExitShowFromschemaFilter is called when exiting the showFromschemaFilter production.
	ExitShowFromschemaFilter(c *ShowFromschemaFilterContext)

	// ExitShowRoutinecode is called when exiting the showRoutinecode production.
	ExitShowRoutinecode(c *ShowRoutinecodeContext)

	// ExitShowGrants is called when exiting the showGrants production.
	ExitShowGrants(c *ShowGrantsContext)

	// ExitShowIndexes is called when exiting the showIndexes production.
	ExitShowIndexes(c *ShowIndexesContext)

	// ExitShowOpentables is called when exiting the showOpentables production.
	ExitShowOpentables(c *ShowOpentablesContext)

	// ExitShowProfile is called when exiting the showProfile production.
	ExitShowProfile(c *ShowProfileContext)

	// ExitShowSlavestatus is called when exiting the showSlavestatus production.
	ExitShowSlavestatus(c *ShowSlavestatusContext)

	// ExitVariable_clause is called when exiting the variable_clause production.
	ExitVariable_clause(c *Variable_clauseContext)

	// ExitShow_filter is called when exiting the show_filter production.
	ExitShow_filter(c *Show_filterContext)

	// ExitShow_profile_type is called when exiting the show_profile_type production.
	ExitShow_profile_type(c *Show_profile_typeContext)

	// ExitBinlog_statement is called when exiting the binlog_statement production.
	ExitBinlog_statement(c *Binlog_statementContext)

	// ExitCache_index_statement is called when exiting the cache_index_statement production.
	ExitCache_index_statement(c *Cache_index_statementContext)

	// ExitFlush_statement is called when exiting the flush_statement production.
	ExitFlush_statement(c *Flush_statementContext)

	// ExitKill_statement is called when exiting the kill_statement production.
	ExitKill_statement(c *Kill_statementContext)

	// ExitLoad_index_into_cache is called when exiting the load_index_into_cache production.
	ExitLoad_index_into_cache(c *Load_index_into_cacheContext)

	// ExitReset_statement is called when exiting the reset_statement production.
	ExitReset_statement(c *Reset_statementContext)

	// ExitShutdown_statement is called when exiting the shutdown_statement production.
	ExitShutdown_statement(c *Shutdown_statementContext)

	// ExitTbl_index_list is called when exiting the tbl_index_list production.
	ExitTbl_index_list(c *Tbl_index_listContext)

	// ExitFlush_option is called when exiting the flush_option production.
	ExitFlush_option(c *Flush_optionContext)

	// ExitLoad_tbl_index_list is called when exiting the load_tbl_index_list production.
	ExitLoad_tbl_index_list(c *Load_tbl_index_listContext)

	// ExitSimple_describe_statement is called when exiting the simple_describe_statement production.
	ExitSimple_describe_statement(c *Simple_describe_statementContext)

	// ExitFull_describe_statement is called when exiting the full_describe_statement production.
	ExitFull_describe_statement(c *Full_describe_statementContext)

	// ExitHelp_statement is called when exiting the help_statement production.
	ExitHelp_statement(c *Help_statementContext)

	// ExitUse_statement is called when exiting the use_statement production.
	ExitUse_statement(c *Use_statementContext)

	// ExitDescstmtDescObj is called when exiting the descstmtDescObj production.
	ExitDescstmtDescObj(c *DescstmtDescObjContext)

	// ExitConnectionDescObj is called when exiting the connectionDescObj production.
	ExitConnectionDescObj(c *ConnectionDescObjContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitFull_id is called when exiting the full_id production.
	ExitFull_id(c *Full_idContext)

	// ExitFull_column_name is called when exiting the full_column_name production.
	ExitFull_column_name(c *Full_column_nameContext)

	// ExitIndex_col_name is called when exiting the index_col_name production.
	ExitIndex_col_name(c *Index_col_nameContext)

	// ExitUser_name is called when exiting the user_name production.
	ExitUser_name(c *User_nameContext)

	// ExitMysql_variable is called when exiting the mysql_variable production.
	ExitMysql_variable(c *Mysql_variableContext)

	// ExitCharset_name is called when exiting the charset_name production.
	ExitCharset_name(c *Charset_nameContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitEngine_name is called when exiting the engine_name production.
	ExitEngine_name(c *Engine_nameContext)

	// ExitUuid_set is called when exiting the uuid_set production.
	ExitUuid_set(c *Uuid_setContext)

	// ExitXid is called when exiting the xid production.
	ExitXid(c *XidContext)

	// ExitXid_string_id is called when exiting the xid_string_id production.
	ExitXid_string_id(c *Xid_string_idContext)

	// ExitAuth_plugin is called when exiting the auth_plugin production.
	ExitAuth_plugin(c *Auth_pluginContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)

	// ExitSimple_id is called when exiting the simple_id production.
	ExitSimple_id(c *Simple_idContext)

	// ExitDot_ext_id is called when exiting the dot_ext_id production.
	ExitDot_ext_id(c *Dot_ext_idContext)

	// ExitDecimal_literal is called when exiting the decimal_literal production.
	ExitDecimal_literal(c *Decimal_literalContext)

	// ExitFilesize_literal is called when exiting the filesize_literal production.
	ExitFilesize_literal(c *Filesize_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitHexadecimal_literal is called when exiting the hexadecimal_literal production.
	ExitHexadecimal_literal(c *Hexadecimal_literalContext)

	// ExitNull_notnull is called when exiting the null_notnull production.
	ExitNull_notnull(c *Null_notnullContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitCharDatatype is called when exiting the charDatatype production.
	ExitCharDatatype(c *CharDatatypeContext)

	// ExitDimensionDatatype is called when exiting the dimensionDatatype production.
	ExitDimensionDatatype(c *DimensionDatatypeContext)

	// ExitSimpleDatatype is called when exiting the simpleDatatype production.
	ExitSimpleDatatype(c *SimpleDatatypeContext)

	// ExitCollectCharDatatype is called when exiting the collectCharDatatype production.
	ExitCollectCharDatatype(c *CollectCharDatatypeContext)

	// ExitSpatialDatatype is called when exiting the spatialDatatype production.
	ExitSpatialDatatype(c *SpatialDatatypeContext)

	// ExitData_type_to_convert is called when exiting the data_type_to_convert production.
	ExitData_type_to_convert(c *Data_type_to_convertContext)

	// ExitSpatial_data_type is called when exiting the spatial_data_type production.
	ExitSpatial_data_type(c *Spatial_data_typeContext)

	// ExitLength_one_dimension is called when exiting the length_one_dimension production.
	ExitLength_one_dimension(c *Length_one_dimensionContext)

	// ExitLength_two_dimension is called when exiting the length_two_dimension production.
	ExitLength_two_dimension(c *Length_two_dimensionContext)

	// ExitLength_two_optional_dimension is called when exiting the length_two_optional_dimension production.
	ExitLength_two_optional_dimension(c *Length_two_optional_dimensionContext)

	// ExitId_list is called when exiting the id_list production.
	ExitId_list(c *Id_listContext)

	// ExitTable_list is called when exiting the table_list production.
	ExitTable_list(c *Table_listContext)

	// ExitTable_pair_list is called when exiting the table_pair_list production.
	ExitTable_pair_list(c *Table_pair_listContext)

	// ExitIndex_colname_list is called when exiting the index_colname_list production.
	ExitIndex_colname_list(c *Index_colname_listContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitConstant_list is called when exiting the constant_list production.
	ExitConstant_list(c *Constant_listContext)

	// ExitSimple_string_list is called when exiting the simple_string_list production.
	ExitSimple_string_list(c *Simple_string_listContext)

	// ExitUser_var_list is called when exiting the user_var_list production.
	ExitUser_var_list(c *User_var_listContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitIf_exists is called when exiting the if_exists production.
	ExitIf_exists(c *If_existsContext)

	// ExitIf_not_exists is called when exiting the if_not_exists production.
	ExitIf_not_exists(c *If_not_existsContext)

	// ExitSpecificFunctionCall is called when exiting the specificFunctionCall production.
	ExitSpecificFunctionCall(c *SpecificFunctionCallContext)

	// ExitAggregateFunctionCall is called when exiting the aggregateFunctionCall production.
	ExitAggregateFunctionCall(c *AggregateFunctionCallContext)

	// ExitScalarFunctionCall is called when exiting the scalarFunctionCall production.
	ExitScalarFunctionCall(c *ScalarFunctionCallContext)

	// ExitUdfFunctionCall is called when exiting the udfFunctionCall production.
	ExitUdfFunctionCall(c *UdfFunctionCallContext)

	// ExitSimpleSpecificFCall is called when exiting the simpleSpecificFCall production.
	ExitSimpleSpecificFCall(c *SimpleSpecificFCallContext)

	// ExitConvertDataTypeFCall is called when exiting the convertDataTypeFCall production.
	ExitConvertDataTypeFCall(c *ConvertDataTypeFCallContext)

	// ExitValuesFCall is called when exiting the valuesFCall production.
	ExitValuesFCall(c *ValuesFCallContext)

	// ExitCaseFCall is called when exiting the caseFCall production.
	ExitCaseFCall(c *CaseFCallContext)

	// ExitCharFCall is called when exiting the charFCall production.
	ExitCharFCall(c *CharFCallContext)

	// ExitPositionFCall is called when exiting the positionFCall production.
	ExitPositionFCall(c *PositionFCallContext)

	// ExitSubstrFCall is called when exiting the substrFCall production.
	ExitSubstrFCall(c *SubstrFCallContext)

	// ExitTrimFCall is called when exiting the trimFCall production.
	ExitTrimFCall(c *TrimFCallContext)

	// ExitWeightFCall is called when exiting the weightFCall production.
	ExitWeightFCall(c *WeightFCallContext)

	// ExitExtractFCall is called when exiting the extractFCall production.
	ExitExtractFCall(c *ExtractFCallContext)

	// ExitGetFormatFCall is called when exiting the getFormatFCall production.
	ExitGetFormatFCall(c *GetFormatFCallContext)

	// ExitLevelWeightFList is called when exiting the levelWeightFList production.
	ExitLevelWeightFList(c *LevelWeightFListContext)

	// ExitLevelWeightFRange is called when exiting the levelWeightFRange production.
	ExitLevelWeightFRange(c *LevelWeightFRangeContext)

	// ExitAggregate_windowed_function is called when exiting the aggregate_windowed_function production.
	ExitAggregate_windowed_function(c *Aggregate_windowed_functionContext)

	// ExitScalar_function_name is called when exiting the scalar_function_name production.
	ExitScalar_function_name(c *Scalar_function_nameContext)

	// ExitFunction_args is called when exiting the function_args production.
	ExitFunction_args(c *Function_argsContext)

	// ExitFunction_arg is called when exiting the function_arg production.
	ExitFunction_arg(c *Function_argContext)

	// ExitIsExpression is called when exiting the isExpression production.
	ExitIsExpression(c *IsExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitSoundsLikePredicate is called when exiting the soundsLikePredicate production.
	ExitSoundsLikePredicate(c *SoundsLikePredicateContext)

	// ExitExpressionAtomPredicate is called when exiting the expressionAtomPredicate production.
	ExitExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitSubqueryComparasionPredicate is called when exiting the subqueryComparasionPredicate production.
	ExitSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitBinaryComparasionPredicate is called when exiting the binaryComparasionPredicate production.
	ExitBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitRegexpPredicate is called when exiting the regexpPredicate production.
	ExitRegexpPredicate(c *RegexpPredicateContext)

	// ExitUnaryExpressionAtom is called when exiting the unaryExpressionAtom production.
	ExitUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// ExitExistsExpessionAtom is called when exiting the existsExpessionAtom production.
	ExitExistsExpessionAtom(c *ExistsExpessionAtomContext)

	// ExitConstantExpressionAtom is called when exiting the constantExpressionAtom production.
	ExitConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// ExitFunctionCallExpressionAtom is called when exiting the functionCallExpressionAtom production.
	ExitFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// ExitMysqlVariableExpressionAtom is called when exiting the mysqlVariableExpressionAtom production.
	ExitMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// ExitBinaryExpressionAtom is called when exiting the binaryExpressionAtom production.
	ExitBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// ExitFullColumnNameExpressionAtom is called when exiting the fullColumnNameExpressionAtom production.
	ExitFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// ExitDefaultExpressionAtom is called when exiting the defaultExpressionAtom production.
	ExitDefaultExpressionAtom(c *DefaultExpressionAtomContext)

	// ExitBitExpressionAtom is called when exiting the bitExpressionAtom production.
	ExitBitExpressionAtom(c *BitExpressionAtomContext)

	// ExitNestedExpressionAtom is called when exiting the nestedExpressionAtom production.
	ExitNestedExpressionAtom(c *NestedExpressionAtomContext)

	// ExitMathExpressionAtom is called when exiting the mathExpressionAtom production.
	ExitMathExpressionAtom(c *MathExpressionAtomContext)

	// ExitIntervalExpressionAtom is called when exiting the intervalExpressionAtom production.
	ExitIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitLogical_operator is called when exiting the logical_operator production.
	ExitLogical_operator(c *Logical_operatorContext)

	// ExitBit_operator is called when exiting the bit_operator production.
	ExitBit_operator(c *Bit_operatorContext)

	// ExitMath_operator is called when exiting the math_operator production.
	ExitMath_operator(c *Math_operatorContext)

	// ExitCharset_name_base is called when exiting the charset_name_base production.
	ExitCharset_name_base(c *Charset_name_baseContext)

	// ExitTransaction_level_base is called when exiting the transaction_level_base production.
	ExitTransaction_level_base(c *Transaction_level_baseContext)

	// ExitPrivileges_base is called when exiting the privileges_base production.
	ExitPrivileges_base(c *Privileges_baseContext)

	// ExitInterval_type_base is called when exiting the interval_type_base production.
	ExitInterval_type_base(c *Interval_type_baseContext)

	// ExitData_type_base is called when exiting the data_type_base production.
	ExitData_type_base(c *Data_type_baseContext)

	// ExitKeywords_can_be_id is called when exiting the keywords_can_be_id production.
	ExitKeywords_can_be_id(c *Keywords_can_be_idContext)

	// ExitFunction_name_base is called when exiting the function_name_base production.
	ExitFunction_name_base(c *Function_name_baseContext)
}
