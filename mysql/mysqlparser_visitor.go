// Generated from MySqlParser.g4 by ANTLR 4.7.

package mysql // MySqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MySqlParser.
type MySqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySqlParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MySqlParser#sql_statements.
	VisitSql_statements(ctx *Sql_statementsContext) interface{}

	// Visit a parse tree produced by MySqlParser#sql_statement.
	VisitSql_statement(ctx *Sql_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#empty_statement.
	VisitEmpty_statement(ctx *Empty_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#ddl_statement.
	VisitDdl_statement(ctx *Ddl_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#dml_statement.
	VisitDml_statement(ctx *Dml_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#transaction_statement.
	VisitTransaction_statement(ctx *Transaction_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#replication_statement.
	VisitReplication_statement(ctx *Replication_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#prepared_statement.
	VisitPrepared_statement(ctx *Prepared_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#compound_statement.
	VisitCompound_statement(ctx *Compound_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#administration_statement.
	VisitAdministration_statement(ctx *Administration_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#utility_statement.
	VisitUtility_statement(ctx *Utility_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_database.
	VisitCreate_database(ctx *Create_databaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_event.
	VisitCreate_event(ctx *Create_eventContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_index.
	VisitCreate_index(ctx *Create_indexContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_logfile_group.
	VisitCreate_logfile_group(ctx *Create_logfile_groupContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_procedure.
	VisitCreate_procedure(ctx *Create_procedureContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_function.
	VisitCreate_function(ctx *Create_functionContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_server.
	VisitCreate_server(ctx *Create_serverContext) interface{}

	// Visit a parse tree produced by MySqlParser#copyCreateTable.
	VisitCopyCreateTable(ctx *CopyCreateTableContext) interface{}

	// Visit a parse tree produced by MySqlParser#queryCreateTable.
	VisitQueryCreateTable(ctx *QueryCreateTableContext) interface{}

	// Visit a parse tree produced by MySqlParser#colCreateTable.
	VisitColCreateTable(ctx *ColCreateTableContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_tablespace_innodb.
	VisitCreate_tablespace_innodb(ctx *Create_tablespace_innodbContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_tablespace_ndb.
	VisitCreate_tablespace_ndb(ctx *Create_tablespace_ndbContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_trigger.
	VisitCreate_trigger(ctx *Create_triggerContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_view.
	VisitCreate_view(ctx *Create_viewContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_database_option.
	VisitCreate_database_option(ctx *Create_database_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#owner_statement.
	VisitOwner_statement(ctx *Owner_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#preciseSchedule.
	VisitPreciseSchedule(ctx *PreciseScheduleContext) interface{}

	// Visit a parse tree produced by MySqlParser#intervalSchedule.
	VisitIntervalSchedule(ctx *IntervalScheduleContext) interface{}

	// Visit a parse tree produced by MySqlParser#timestamp_value.
	VisitTimestamp_value(ctx *Timestamp_valueContext) interface{}

	// Visit a parse tree produced by MySqlParser#interval_expr.
	VisitInterval_expr(ctx *Interval_exprContext) interface{}

	// Visit a parse tree produced by MySqlParser#interval_type.
	VisitInterval_type(ctx *Interval_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#index_type.
	VisitIndex_type(ctx *Index_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#index_option.
	VisitIndex_option(ctx *Index_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#proc_param.
	VisitProc_param(ctx *Proc_paramContext) interface{}

	// Visit a parse tree produced by MySqlParser#func_param.
	VisitFunc_param(ctx *Func_paramContext) interface{}

	// Visit a parse tree produced by MySqlParser#rcComment.
	VisitRcComment(ctx *RcCommentContext) interface{}

	// Visit a parse tree produced by MySqlParser#rcSqllang.
	VisitRcSqllang(ctx *RcSqllangContext) interface{}

	// Visit a parse tree produced by MySqlParser#rcDeterm.
	VisitRcDeterm(ctx *RcDetermContext) interface{}

	// Visit a parse tree produced by MySqlParser#rcSqldata.
	VisitRcSqldata(ctx *RcSqldataContext) interface{}

	// Visit a parse tree produced by MySqlParser#rcSecurestmt.
	VisitRcSecurestmt(ctx *RcSecurestmtContext) interface{}

	// Visit a parse tree produced by MySqlParser#server_option.
	VisitServer_option(ctx *Server_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#column_def_table_constraints.
	VisitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) interface{}

	// Visit a parse tree produced by MySqlParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#constraintDefinition.
	VisitConstraintDefinition(ctx *ConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#indexDefinition.
	VisitIndexDefinition(ctx *IndexDefinitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#column_definition.
	VisitColumn_definition(ctx *Column_definitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrNull.
	VisitColConstrNull(ctx *ColConstrNullContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrDflt.
	VisitColConstrDflt(ctx *ColConstrDfltContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrAuInc.
	VisitColConstrAuInc(ctx *ColConstrAuIncContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrPK.
	VisitColConstrPK(ctx *ColConstrPKContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrUK.
	VisitColConstrUK(ctx *ColConstrUKContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrComment.
	VisitColConstrComment(ctx *ColConstrCommentContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrForm.
	VisitColConstrForm(ctx *ColConstrFormContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrStorage.
	VisitColConstrStorage(ctx *ColConstrStorageContext) interface{}

	// Visit a parse tree produced by MySqlParser#colConstrRefdef.
	VisitColConstrRefdef(ctx *ColConstrRefdefContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblConstrPK.
	VisitTblConstrPK(ctx *TblConstrPKContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblConstrUK.
	VisitTblConstrUK(ctx *TblConstrUKContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblConstrFK.
	VisitTblConstrFK(ctx *TblConstrFKContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblConstCheck.
	VisitTblConstCheck(ctx *TblConstCheckContext) interface{}

	// Visit a parse tree produced by MySqlParser#reference_definition.
	VisitReference_definition(ctx *Reference_definitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#on_delete_action.
	VisitOn_delete_action(ctx *On_delete_actionContext) interface{}

	// Visit a parse tree produced by MySqlParser#on_update_action.
	VisitOn_update_action(ctx *On_update_actionContext) interface{}

	// Visit a parse tree produced by MySqlParser#reference_action_control_type.
	VisitReference_action_control_type(ctx *Reference_action_control_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleIndex.
	VisitSimpleIndex(ctx *SimpleIndexContext) interface{}

	// Visit a parse tree produced by MySqlParser#specIndex.
	VisitSpecIndex(ctx *SpecIndexContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptEngine.
	VisitTblOptEngine(ctx *TblOptEngineContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptAuInc.
	VisitTblOptAuInc(ctx *TblOptAuIncContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptAvgRLen.
	VisitTblOptAvgRLen(ctx *TblOptAvgRLenContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptDefCharSet.
	VisitTblOptDefCharSet(ctx *TblOptDefCharSetContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptChkSum.
	VisitTblOptChkSum(ctx *TblOptChkSumContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptDefCollate.
	VisitTblOptDefCollate(ctx *TblOptDefCollateContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptComment.
	VisitTblOptComment(ctx *TblOptCommentContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptCompr.
	VisitTblOptCompr(ctx *TblOptComprContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptConn.
	VisitTblOptConn(ctx *TblOptConnContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptDataDir.
	VisitTblOptDataDir(ctx *TblOptDataDirContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptDelKW.
	VisitTblOptDelKW(ctx *TblOptDelKWContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptEncr.
	VisitTblOptEncr(ctx *TblOptEncrContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptIndexDir.
	VisitTblOptIndexDir(ctx *TblOptIndexDirContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptInsMeth.
	VisitTblOptInsMeth(ctx *TblOptInsMethContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptKeyBlockSz.
	VisitTblOptKeyBlockSz(ctx *TblOptKeyBlockSzContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptMaxRows.
	VisitTblOptMaxRows(ctx *TblOptMaxRowsContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptMinRows.
	VisitTblOptMinRows(ctx *TblOptMinRowsContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptPackK.
	VisitTblOptPackK(ctx *TblOptPackKContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptPasswd.
	VisitTblOptPasswd(ctx *TblOptPasswdContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptRowFormat.
	VisitTblOptRowFormat(ctx *TblOptRowFormatContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptStatAutoR.
	VisitTblOptStatAutoR(ctx *TblOptStatAutoRContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptStatPersist.
	VisitTblOptStatPersist(ctx *TblOptStatPersistContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptStatSamplPg.
	VisitTblOptStatSamplPg(ctx *TblOptStatSamplPgContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptTablespace.
	VisitTblOptTablespace(ctx *TblOptTablespaceContext) interface{}

	// Visit a parse tree produced by MySqlParser#tblOptUnion.
	VisitTblOptUnion(ctx *TblOptUnionContext) interface{}

	// Visit a parse tree produced by MySqlParser#partition_options.
	VisitPartition_options(ctx *Partition_optionsContext) interface{}

	// Visit a parse tree produced by MySqlParser#partition_function_definition.
	VisitPartition_function_definition(ctx *Partition_function_definitionContext) interface{}

	// Visit a parse tree produced by MySqlParser#linear_partition_func_def.
	VisitLinear_partition_func_def(ctx *Linear_partition_func_defContext) interface{}

	// Visit a parse tree produced by MySqlParser#partition_def.
	VisitPartition_def(ctx *Partition_defContext) interface{}

	// Visit a parse tree produced by MySqlParser#subpartition_def.
	VisitSubpartition_def(ctx *Subpartition_defContext) interface{}

	// Visit a parse tree produced by MySqlParser#alterDb.
	VisitAlterDb(ctx *AlterDbContext) interface{}

	// Visit a parse tree produced by MySqlParser#alterDbUpgradeName.
	VisitAlterDbUpgradeName(ctx *AlterDbUpgradeNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_event.
	VisitAlter_event(ctx *Alter_eventContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_function.
	VisitAlter_function(ctx *Alter_functionContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_instance.
	VisitAlter_instance(ctx *Alter_instanceContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_logfile_group.
	VisitAlter_logfile_group(ctx *Alter_logfile_groupContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_procedure.
	VisitAlter_procedure(ctx *Alter_procedureContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_server.
	VisitAlter_server(ctx *Alter_serverContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_table.
	VisitAlter_table(ctx *Alter_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_tablespace.
	VisitAlter_tablespace(ctx *Alter_tablespaceContext) interface{}

	// Visit a parse tree produced by MySqlParser#alter_view.
	VisitAlter_view(ctx *Alter_viewContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblTableOpt.
	VisitAltblTableOpt(ctx *AltblTableOptContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddCol.
	VisitAltblAddCol(ctx *AltblAddColContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddCols.
	VisitAltblAddCols(ctx *AltblAddColsContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddIndex.
	VisitAltblAddIndex(ctx *AltblAddIndexContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddPK.
	VisitAltblAddPK(ctx *AltblAddPKContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddUK.
	VisitAltblAddUK(ctx *AltblAddUKContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddSpecIndex.
	VisitAltblAddSpecIndex(ctx *AltblAddSpecIndexContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddFK.
	VisitAltblAddFK(ctx *AltblAddFKContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAlg.
	VisitAltblAlg(ctx *AltblAlgContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblColDef.
	VisitAltblColDef(ctx *AltblColDefContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblColChange.
	VisitAltblColChange(ctx *AltblColChangeContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblLock.
	VisitAltblLock(ctx *AltblLockContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblColMod.
	VisitAltblColMod(ctx *AltblColModContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblColDrop.
	VisitAltblColDrop(ctx *AltblColDropContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDropPK.
	VisitAltblDropPK(ctx *AltblDropPKContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDropIndex.
	VisitAltblDropIndex(ctx *AltblDropIndexContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDropFK.
	VisitAltblDropFK(ctx *AltblDropFKContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDisKey.
	VisitAltblDisKey(ctx *AltblDisKeyContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblEnKey.
	VisitAltblEnKey(ctx *AltblEnKeyContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblRenameTbl.
	VisitAltblRenameTbl(ctx *AltblRenameTblContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblResort.
	VisitAltblResort(ctx *AltblResortContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblConvert.
	VisitAltblConvert(ctx *AltblConvertContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDefCharset.
	VisitAltblDefCharset(ctx *AltblDefCharsetContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDisTblspace.
	VisitAltblDisTblspace(ctx *AltblDisTblspaceContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblImpTblSpace.
	VisitAltblImpTblSpace(ctx *AltblImpTblSpaceContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblForce.
	VisitAltblForce(ctx *AltblForceContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblValid.
	VisitAltblValid(ctx *AltblValidContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAddPart.
	VisitAltblAddPart(ctx *AltblAddPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDropPart.
	VisitAltblDropPart(ctx *AltblDropPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblDiscartPart.
	VisitAltblDiscartPart(ctx *AltblDiscartPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblImportPart.
	VisitAltblImportPart(ctx *AltblImportPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblTruncPart.
	VisitAltblTruncPart(ctx *AltblTruncPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblCoalPart.
	VisitAltblCoalPart(ctx *AltblCoalPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblReorgPart.
	VisitAltblReorgPart(ctx *AltblReorgPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblExchPart.
	VisitAltblExchPart(ctx *AltblExchPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblAnalPart.
	VisitAltblAnalPart(ctx *AltblAnalPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblCheckPart.
	VisitAltblCheckPart(ctx *AltblCheckPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblOptimPart.
	VisitAltblOptimPart(ctx *AltblOptimPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblRebuildPart.
	VisitAltblRebuildPart(ctx *AltblRebuildPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblRepairPart.
	VisitAltblRepairPart(ctx *AltblRepairPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblRemovePart.
	VisitAltblRemovePart(ctx *AltblRemovePartContext) interface{}

	// Visit a parse tree produced by MySqlParser#altblUpgrPart.
	VisitAltblUpgrPart(ctx *AltblUpgrPartContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_database.
	VisitDrop_database(ctx *Drop_databaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_event.
	VisitDrop_event(ctx *Drop_eventContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_index.
	VisitDrop_index(ctx *Drop_indexContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_logfile_group.
	VisitDrop_logfile_group(ctx *Drop_logfile_groupContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_procedure.
	VisitDrop_procedure(ctx *Drop_procedureContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_function.
	VisitDrop_function(ctx *Drop_functionContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_server.
	VisitDrop_server(ctx *Drop_serverContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_table.
	VisitDrop_table(ctx *Drop_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_tablespace.
	VisitDrop_tablespace(ctx *Drop_tablespaceContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_trigger.
	VisitDrop_trigger(ctx *Drop_triggerContext) interface{}

	// Visit a parse tree produced by MySqlParser#drop_view.
	VisitDrop_view(ctx *Drop_viewContext) interface{}

	// Visit a parse tree produced by MySqlParser#rename_table.
	VisitRename_table(ctx *Rename_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#truncate_table.
	VisitTruncate_table(ctx *Truncate_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#call_statement.
	VisitCall_statement(ctx *Call_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#delete_statement.
	VisitDelete_statement(ctx *Delete_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#do_statement.
	VisitDo_statement(ctx *Do_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_statement.
	VisitHandler_statement(ctx *Handler_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#insert_statement.
	VisitInsert_statement(ctx *Insert_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#load_data_statement.
	VisitLoad_data_statement(ctx *Load_data_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#load_xml_statement.
	VisitLoad_xml_statement(ctx *Load_xml_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#replace_statement.
	VisitReplace_statement(ctx *Replace_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleSelect.
	VisitSimpleSelect(ctx *SimpleSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#parenSelect.
	VisitParenSelect(ctx *ParenSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionParenSelect.
	VisitUnionParenSelect(ctx *UnionParenSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#update_statement.
	VisitUpdate_statement(ctx *Update_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#insert_statement_value.
	VisitInsert_statement_value(ctx *Insert_statement_valueContext) interface{}

	// Visit a parse tree produced by MySqlParser#update_elem.
	VisitUpdate_elem(ctx *Update_elemContext) interface{}

	// Visit a parse tree produced by MySqlParser#col_or_uservar.
	VisitCol_or_uservar(ctx *Col_or_uservarContext) interface{}

	// Visit a parse tree produced by MySqlParser#single_delete_statement.
	VisitSingle_delete_statement(ctx *Single_delete_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#multiple_delete_statement.
	VisitMultiple_delete_statement(ctx *Multiple_delete_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_open_statement.
	VisitHandler_open_statement(ctx *Handler_open_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_read_index_statement.
	VisitHandler_read_index_statement(ctx *Handler_read_index_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_read_statement.
	VisitHandler_read_statement(ctx *Handler_read_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_close_statement.
	VisitHandler_close_statement(ctx *Handler_close_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#single_update_statement.
	VisitSingle_update_statement(ctx *Single_update_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#multiple_update_statement.
	VisitMultiple_update_statement(ctx *Multiple_update_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#order_by_clause.
	VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#order_by_expression.
	VisitOrder_by_expression(ctx *Order_by_expressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#table_sources.
	VisitTable_sources(ctx *Table_sourcesContext) interface{}

	// Visit a parse tree produced by MySqlParser#table_source.
	VisitTable_source(ctx *Table_sourceContext) interface{}

	// Visit a parse tree produced by MySqlParser#atomTableItem.
	VisitAtomTableItem(ctx *AtomTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryTableItem.
	VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourcesItem.
	VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#index_hint.
	VisitIndex_hint(ctx *Index_hintContext) interface{}

	// Visit a parse tree produced by MySqlParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#straightJoin.
	VisitStraightJoin(ctx *StraightJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#outerJoin.
	VisitOuterJoin(ctx *OuterJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#naturalJoin.
	VisitNaturalJoin(ctx *NaturalJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by MySqlParser#query_expression.
	VisitQuery_expression(ctx *Query_expressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#query_expression_nointo.
	VisitQuery_expression_nointo(ctx *Query_expression_nointoContext) interface{}

	// Visit a parse tree produced by MySqlParser#query_specification.
	VisitQuery_specification(ctx *Query_specificationContext) interface{}

	// Visit a parse tree produced by MySqlParser#query_specification_nointo.
	VisitQuery_specification_nointo(ctx *Query_specification_nointoContext) interface{}

	// Visit a parse tree produced by MySqlParser#union_parenth.
	VisitUnion_parenth(ctx *Union_parenthContext) interface{}

	// Visit a parse tree produced by MySqlParser#union_statement.
	VisitUnion_statement(ctx *Union_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#select_spec.
	VisitSelect_spec(ctx *Select_specContext) interface{}

	// Visit a parse tree produced by MySqlParser#select_list.
	VisitSelect_list(ctx *Select_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#sellistelAllCol.
	VisitSellistelAllCol(ctx *SellistelAllColContext) interface{}

	// Visit a parse tree produced by MySqlParser#sellistelCol.
	VisitSellistelCol(ctx *SellistelColContext) interface{}

	// Visit a parse tree produced by MySqlParser#sellistelFunc.
	VisitSellistelFunc(ctx *SellistelFuncContext) interface{}

	// Visit a parse tree produced by MySqlParser#sellistelExpr.
	VisitSellistelExpr(ctx *SellistelExprContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoVars.
	VisitSelectIntoVars(ctx *SelectIntoVarsContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoDump.
	VisitSelectIntoDump(ctx *SelectIntoDumpContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoOutfile.
	VisitSelectIntoOutfile(ctx *SelectIntoOutfileContext) interface{}

	// Visit a parse tree produced by MySqlParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#group_by_item.
	VisitGroup_by_item(ctx *Group_by_itemContext) interface{}

	// Visit a parse tree produced by MySqlParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#start_transaction.
	VisitStart_transaction(ctx *Start_transactionContext) interface{}

	// Visit a parse tree produced by MySqlParser#begin_work.
	VisitBegin_work(ctx *Begin_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#commit_work.
	VisitCommit_work(ctx *Commit_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#rollback_work.
	VisitRollback_work(ctx *Rollback_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#savepoint_statement.
	VisitSavepoint_statement(ctx *Savepoint_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#rollback_statement.
	VisitRollback_statement(ctx *Rollback_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#release_statement.
	VisitRelease_statement(ctx *Release_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#lock_tables.
	VisitLock_tables(ctx *Lock_tablesContext) interface{}

	// Visit a parse tree produced by MySqlParser#unlock_tables.
	VisitUnlock_tables(ctx *Unlock_tablesContext) interface{}

	// Visit a parse tree produced by MySqlParser#set_autocommit_statement.
	VisitSet_autocommit_statement(ctx *Set_autocommit_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#set_transaction_statement.
	VisitSet_transaction_statement(ctx *Set_transaction_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#transact_option.
	VisitTransact_option(ctx *Transact_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#lock_table_element.
	VisitLock_table_element(ctx *Lock_table_elementContext) interface{}

	// Visit a parse tree produced by MySqlParser#trans_characteristic.
	VisitTrans_characteristic(ctx *Trans_characteristicContext) interface{}

	// Visit a parse tree produced by MySqlParser#transaction_level.
	VisitTransaction_level(ctx *Transaction_levelContext) interface{}

	// Visit a parse tree produced by MySqlParser#change_master.
	VisitChange_master(ctx *Change_masterContext) interface{}

	// Visit a parse tree produced by MySqlParser#change_repl_filter.
	VisitChange_repl_filter(ctx *Change_repl_filterContext) interface{}

	// Visit a parse tree produced by MySqlParser#purge_binary_logs.
	VisitPurge_binary_logs(ctx *Purge_binary_logsContext) interface{}

	// Visit a parse tree produced by MySqlParser#reset_master.
	VisitReset_master(ctx *Reset_masterContext) interface{}

	// Visit a parse tree produced by MySqlParser#reset_slave.
	VisitReset_slave(ctx *Reset_slaveContext) interface{}

	// Visit a parse tree produced by MySqlParser#start_slave.
	VisitStart_slave(ctx *Start_slaveContext) interface{}

	// Visit a parse tree produced by MySqlParser#stop_slave.
	VisitStop_slave(ctx *Stop_slaveContext) interface{}

	// Visit a parse tree produced by MySqlParser#start_group_repl.
	VisitStart_group_repl(ctx *Start_group_replContext) interface{}

	// Visit a parse tree produced by MySqlParser#stop_group_repl.
	VisitStop_group_repl(ctx *Stop_group_replContext) interface{}

	// Visit a parse tree produced by MySqlParser#masterOptString.
	VisitMasterOptString(ctx *MasterOptStringContext) interface{}

	// Visit a parse tree produced by MySqlParser#masterOptDecimal.
	VisitMasterOptDecimal(ctx *MasterOptDecimalContext) interface{}

	// Visit a parse tree produced by MySqlParser#masterOptBool.
	VisitMasterOptBool(ctx *MasterOptBoolContext) interface{}

	// Visit a parse tree produced by MySqlParser#masterOptReal.
	VisitMasterOptReal(ctx *MasterOptRealContext) interface{}

	// Visit a parse tree produced by MySqlParser#masterOptIdList.
	VisitMasterOptIdList(ctx *MasterOptIdListContext) interface{}

	// Visit a parse tree produced by MySqlParser#string_master_option.
	VisitString_master_option(ctx *String_master_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#decimal_master_option.
	VisitDecimal_master_option(ctx *Decimal_master_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#bool_master_option.
	VisitBool_master_option(ctx *Bool_master_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#channel_option.
	VisitChannel_option(ctx *Channel_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#replfilterDbList.
	VisitReplfilterDbList(ctx *ReplfilterDbListContext) interface{}

	// Visit a parse tree produced by MySqlParser#replfilterTableList.
	VisitReplfilterTableList(ctx *ReplfilterTableListContext) interface{}

	// Visit a parse tree produced by MySqlParser#replfilterStableList.
	VisitReplfilterStableList(ctx *ReplfilterStableListContext) interface{}

	// Visit a parse tree produced by MySqlParser#replfilterTablepairList.
	VisitReplfilterTablepairList(ctx *ReplfilterTablepairListContext) interface{}

	// Visit a parse tree produced by MySqlParser#thread_type.
	VisitThread_type(ctx *Thread_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#untilGtidSset.
	VisitUntilGtidSset(ctx *UntilGtidSsetContext) interface{}

	// Visit a parse tree produced by MySqlParser#untilMasterLog.
	VisitUntilMasterLog(ctx *UntilMasterLogContext) interface{}

	// Visit a parse tree produced by MySqlParser#untilRelayLog.
	VisitUntilRelayLog(ctx *UntilRelayLogContext) interface{}

	// Visit a parse tree produced by MySqlParser#untilSqlGaps.
	VisitUntilSqlGaps(ctx *UntilSqlGapsContext) interface{}

	// Visit a parse tree produced by MySqlParser#start_slave_connection_option.
	VisitStart_slave_connection_option(ctx *Start_slave_connection_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#gtid_set.
	VisitGtid_set(ctx *Gtid_setContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_start_transaction.
	VisitXa_start_transaction(ctx *Xa_start_transactionContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_end_transaction.
	VisitXa_end_transaction(ctx *Xa_end_transactionContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_prepare.
	VisitXa_prepare(ctx *Xa_prepareContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_commit_work.
	VisitXa_commit_work(ctx *Xa_commit_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_rollback_work.
	VisitXa_rollback_work(ctx *Xa_rollback_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#xa_recover_work.
	VisitXa_recover_work(ctx *Xa_recover_workContext) interface{}

	// Visit a parse tree produced by MySqlParser#prepare_statement.
	VisitPrepare_statement(ctx *Prepare_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#execute_statement.
	VisitExecute_statement(ctx *Execute_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#deallocate_prepare.
	VisitDeallocate_prepare(ctx *Deallocate_prepareContext) interface{}

	// Visit a parse tree produced by MySqlParser#routine_body.
	VisitRoutine_body(ctx *Routine_bodyContext) interface{}

	// Visit a parse tree produced by MySqlParser#block_statement.
	VisitBlock_statement(ctx *Block_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#iterate_statement.
	VisitIterate_statement(ctx *Iterate_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#leave_statement.
	VisitLeave_statement(ctx *Leave_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#repeat_statement.
	VisitRepeat_statement(ctx *Repeat_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#cursor_statement.
	VisitCursor_statement(ctx *Cursor_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#declare_variable.
	VisitDeclare_variable(ctx *Declare_variableContext) interface{}

	// Visit a parse tree produced by MySqlParser#declare_condition.
	VisitDeclare_condition(ctx *Declare_conditionContext) interface{}

	// Visit a parse tree produced by MySqlParser#declare_cursor.
	VisitDeclare_cursor(ctx *Declare_cursorContext) interface{}

	// Visit a parse tree produced by MySqlParser#declare_handler.
	VisitDeclare_handler(ctx *Declare_handlerContext) interface{}

	// Visit a parse tree produced by MySqlParser#handler_condition_value.
	VisitHandler_condition_value(ctx *Handler_condition_valueContext) interface{}

	// Visit a parse tree produced by MySqlParser#procedure_sql_statement.
	VisitProcedure_sql_statement(ctx *Procedure_sql_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#alterUserMysql56.
	VisitAlterUserMysql56(ctx *AlterUserMysql56Context) interface{}

	// Visit a parse tree produced by MySqlParser#alterUserMysql57.
	VisitAlterUserMysql57(ctx *AlterUserMysql57Context) interface{}

	// Visit a parse tree produced by MySqlParser#createUserMysql56.
	VisitCreateUserMysql56(ctx *CreateUserMysql56Context) interface{}

	// Visit a parse tree produced by MySqlParser#createUserMysql57.
	VisitCreateUserMysql57(ctx *CreateUserMysql57Context) interface{}

	// Visit a parse tree produced by MySqlParser#drop_user.
	VisitDrop_user(ctx *Drop_userContext) interface{}

	// Visit a parse tree produced by MySqlParser#grant_statement.
	VisitGrant_statement(ctx *Grant_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#grant_proxy.
	VisitGrant_proxy(ctx *Grant_proxyContext) interface{}

	// Visit a parse tree produced by MySqlParser#rename_user.
	VisitRename_user(ctx *Rename_userContext) interface{}

	// Visit a parse tree produced by MySqlParser#detailRevoke.
	VisitDetailRevoke(ctx *DetailRevokeContext) interface{}

	// Visit a parse tree produced by MySqlParser#shortRevoke.
	VisitShortRevoke(ctx *ShortRevokeContext) interface{}

	// Visit a parse tree produced by MySqlParser#revoke_proxy.
	VisitRevoke_proxy(ctx *Revoke_proxyContext) interface{}

	// Visit a parse tree produced by MySqlParser#set_password_statement.
	VisitSet_password_statement(ctx *Set_password_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#user_password_option.
	VisitUser_password_option(ctx *User_password_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#authByPassword.
	VisitAuthByPassword(ctx *AuthByPasswordContext) interface{}

	// Visit a parse tree produced by MySqlParser#authByString.
	VisitAuthByString(ctx *AuthByStringContext) interface{}

	// Visit a parse tree produced by MySqlParser#authByHash.
	VisitAuthByHash(ctx *AuthByHashContext) interface{}

	// Visit a parse tree produced by MySqlParser#tls_option.
	VisitTls_option(ctx *Tls_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#user_resource_option.
	VisitUser_resource_option(ctx *User_resource_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#user_lock_option.
	VisitUser_lock_option(ctx *User_lock_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#privelege_clause.
	VisitPrivelege_clause(ctx *Privelege_clauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) interface{}

	// Visit a parse tree produced by MySqlParser#privilege_level.
	VisitPrivilege_level(ctx *Privilege_levelContext) interface{}

	// Visit a parse tree produced by MySqlParser#set_password_option.
	VisitSet_password_option(ctx *Set_password_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#analyze_table.
	VisitAnalyze_table(ctx *Analyze_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#check_table.
	VisitCheck_table(ctx *Check_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#checksum_table.
	VisitChecksum_table(ctx *Checksum_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#optimize_table.
	VisitOptimize_table(ctx *Optimize_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#repair_table.
	VisitRepair_table(ctx *Repair_tableContext) interface{}

	// Visit a parse tree produced by MySqlParser#check_table_option.
	VisitCheck_table_option(ctx *Check_table_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#create_udfunction.
	VisitCreate_udfunction(ctx *Create_udfunctionContext) interface{}

	// Visit a parse tree produced by MySqlParser#install_plugin.
	VisitInstall_plugin(ctx *Install_pluginContext) interface{}

	// Visit a parse tree produced by MySqlParser#uninstall_plugin.
	VisitUninstall_plugin(ctx *Uninstall_pluginContext) interface{}

	// Visit a parse tree produced by MySqlParser#setVariableAssignment.
	VisitSetVariableAssignment(ctx *SetVariableAssignmentContext) interface{}

	// Visit a parse tree produced by MySqlParser#setCharset.
	VisitSetCharset(ctx *SetCharsetContext) interface{}

	// Visit a parse tree produced by MySqlParser#setNames.
	VisitSetNames(ctx *SetNamesContext) interface{}

	// Visit a parse tree produced by MySqlParser#setPasswordStatement.
	VisitSetPasswordStatement(ctx *SetPasswordStatementContext) interface{}

	// Visit a parse tree produced by MySqlParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by MySqlParser#setAutocommit.
	VisitSetAutocommit(ctx *SetAutocommitContext) interface{}

	// Visit a parse tree produced by MySqlParser#showMasterlogs.
	VisitShowMasterlogs(ctx *ShowMasterlogsContext) interface{}

	// Visit a parse tree produced by MySqlParser#showLogevents.
	VisitShowLogevents(ctx *ShowLogeventsContext) interface{}

	// Visit a parse tree produced by MySqlParser#showObjWithFilter.
	VisitShowObjWithFilter(ctx *ShowObjWithFilterContext) interface{}

	// Visit a parse tree produced by MySqlParser#showColumns.
	VisitShowColumns(ctx *ShowColumnsContext) interface{}

	// Visit a parse tree produced by MySqlParser#showCreateDb.
	VisitShowCreateDb(ctx *ShowCreateDbContext) interface{}

	// Visit a parse tree produced by MySqlParser#showCreateFullidobj.
	VisitShowCreateFullidobj(ctx *ShowCreateFullidobjContext) interface{}

	// Visit a parse tree produced by MySqlParser#showCreateUser.
	VisitShowCreateUser(ctx *ShowCreateUserContext) interface{}

	// Visit a parse tree produced by MySqlParser#showEngine.
	VisitShowEngine(ctx *ShowEngineContext) interface{}

	// Visit a parse tree produced by MySqlParser#showGlobalinfo.
	VisitShowGlobalinfo(ctx *ShowGlobalinfoContext) interface{}

	// Visit a parse tree produced by MySqlParser#showErrWarn.
	VisitShowErrWarn(ctx *ShowErrWarnContext) interface{}

	// Visit a parse tree produced by MySqlParser#showCountErrWarn.
	VisitShowCountErrWarn(ctx *ShowCountErrWarnContext) interface{}

	// Visit a parse tree produced by MySqlParser#showFromschemaFilter.
	VisitShowFromschemaFilter(ctx *ShowFromschemaFilterContext) interface{}

	// Visit a parse tree produced by MySqlParser#showRoutinecode.
	VisitShowRoutinecode(ctx *ShowRoutinecodeContext) interface{}

	// Visit a parse tree produced by MySqlParser#showGrants.
	VisitShowGrants(ctx *ShowGrantsContext) interface{}

	// Visit a parse tree produced by MySqlParser#showIndexes.
	VisitShowIndexes(ctx *ShowIndexesContext) interface{}

	// Visit a parse tree produced by MySqlParser#showOpentables.
	VisitShowOpentables(ctx *ShowOpentablesContext) interface{}

	// Visit a parse tree produced by MySqlParser#showProfile.
	VisitShowProfile(ctx *ShowProfileContext) interface{}

	// Visit a parse tree produced by MySqlParser#showSlavestatus.
	VisitShowSlavestatus(ctx *ShowSlavestatusContext) interface{}

	// Visit a parse tree produced by MySqlParser#variable_clause.
	VisitVariable_clause(ctx *Variable_clauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#show_filter.
	VisitShow_filter(ctx *Show_filterContext) interface{}

	// Visit a parse tree produced by MySqlParser#show_profile_type.
	VisitShow_profile_type(ctx *Show_profile_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#binlog_statement.
	VisitBinlog_statement(ctx *Binlog_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#cache_index_statement.
	VisitCache_index_statement(ctx *Cache_index_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#flush_statement.
	VisitFlush_statement(ctx *Flush_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#kill_statement.
	VisitKill_statement(ctx *Kill_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#load_index_into_cache.
	VisitLoad_index_into_cache(ctx *Load_index_into_cacheContext) interface{}

	// Visit a parse tree produced by MySqlParser#reset_statement.
	VisitReset_statement(ctx *Reset_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#shutdown_statement.
	VisitShutdown_statement(ctx *Shutdown_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#tbl_index_list.
	VisitTbl_index_list(ctx *Tbl_index_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#flush_option.
	VisitFlush_option(ctx *Flush_optionContext) interface{}

	// Visit a parse tree produced by MySqlParser#load_tbl_index_list.
	VisitLoad_tbl_index_list(ctx *Load_tbl_index_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#simple_describe_statement.
	VisitSimple_describe_statement(ctx *Simple_describe_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#full_describe_statement.
	VisitFull_describe_statement(ctx *Full_describe_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#help_statement.
	VisitHelp_statement(ctx *Help_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#use_statement.
	VisitUse_statement(ctx *Use_statementContext) interface{}

	// Visit a parse tree produced by MySqlParser#descstmtDescObj.
	VisitDescstmtDescObj(ctx *DescstmtDescObjContext) interface{}

	// Visit a parse tree produced by MySqlParser#connectionDescObj.
	VisitConnectionDescObj(ctx *ConnectionDescObjContext) interface{}

	// Visit a parse tree produced by MySqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#full_id.
	VisitFull_id(ctx *Full_idContext) interface{}

	// Visit a parse tree produced by MySqlParser#full_column_name.
	VisitFull_column_name(ctx *Full_column_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#index_col_name.
	VisitIndex_col_name(ctx *Index_col_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#user_name.
	VisitUser_name(ctx *User_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#mysql_variable.
	VisitMysql_variable(ctx *Mysql_variableContext) interface{}

	// Visit a parse tree produced by MySqlParser#charset_name.
	VisitCharset_name(ctx *Charset_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#collation_name.
	VisitCollation_name(ctx *Collation_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#engine_name.
	VisitEngine_name(ctx *Engine_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#uuid_set.
	VisitUuid_set(ctx *Uuid_setContext) interface{}

	// Visit a parse tree produced by MySqlParser#xid.
	VisitXid(ctx *XidContext) interface{}

	// Visit a parse tree produced by MySqlParser#xid_string_id.
	VisitXid_string_id(ctx *Xid_string_idContext) interface{}

	// Visit a parse tree produced by MySqlParser#auth_plugin.
	VisitAuth_plugin(ctx *Auth_pluginContext) interface{}

	// Visit a parse tree produced by MySqlParser#id_.
	VisitId_(ctx *Id_Context) interface{}

	// Visit a parse tree produced by MySqlParser#simple_id.
	VisitSimple_id(ctx *Simple_idContext) interface{}

	// Visit a parse tree produced by MySqlParser#dot_ext_id.
	VisitDot_ext_id(ctx *Dot_ext_idContext) interface{}

	// Visit a parse tree produced by MySqlParser#decimal_literal.
	VisitDecimal_literal(ctx *Decimal_literalContext) interface{}

	// Visit a parse tree produced by MySqlParser#filesize_literal.
	VisitFilesize_literal(ctx *Filesize_literalContext) interface{}

	// Visit a parse tree produced by MySqlParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

	// Visit a parse tree produced by MySqlParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by MySqlParser#hexadecimal_literal.
	VisitHexadecimal_literal(ctx *Hexadecimal_literalContext) interface{}

	// Visit a parse tree produced by MySqlParser#null_notnull.
	VisitNull_notnull(ctx *Null_notnullContext) interface{}

	// Visit a parse tree produced by MySqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by MySqlParser#charDatatype.
	VisitCharDatatype(ctx *CharDatatypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#dimensionDatatype.
	VisitDimensionDatatype(ctx *DimensionDatatypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleDatatype.
	VisitSimpleDatatype(ctx *SimpleDatatypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#collectCharDatatype.
	VisitCollectCharDatatype(ctx *CollectCharDatatypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#spatialDatatype.
	VisitSpatialDatatype(ctx *SpatialDatatypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#data_type_to_convert.
	VisitData_type_to_convert(ctx *Data_type_to_convertContext) interface{}

	// Visit a parse tree produced by MySqlParser#spatial_data_type.
	VisitSpatial_data_type(ctx *Spatial_data_typeContext) interface{}

	// Visit a parse tree produced by MySqlParser#length_one_dimension.
	VisitLength_one_dimension(ctx *Length_one_dimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#length_two_dimension.
	VisitLength_two_dimension(ctx *Length_two_dimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#length_two_optional_dimension.
	VisitLength_two_optional_dimension(ctx *Length_two_optional_dimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#id_list.
	VisitId_list(ctx *Id_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#table_list.
	VisitTable_list(ctx *Table_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#table_pair_list.
	VisitTable_pair_list(ctx *Table_pair_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#index_colname_list.
	VisitIndex_colname_list(ctx *Index_colname_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#constant_list.
	VisitConstant_list(ctx *Constant_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#simple_string_list.
	VisitSimple_string_list(ctx *Simple_string_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#user_var_list.
	VisitUser_var_list(ctx *User_var_listContext) interface{}

	// Visit a parse tree produced by MySqlParser#default_value.
	VisitDefault_value(ctx *Default_valueContext) interface{}

	// Visit a parse tree produced by MySqlParser#if_exists.
	VisitIf_exists(ctx *If_existsContext) interface{}

	// Visit a parse tree produced by MySqlParser#if_not_exists.
	VisitIf_not_exists(ctx *If_not_existsContext) interface{}

	// Visit a parse tree produced by MySqlParser#specificFunctionCall.
	VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateFunctionCall.
	VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#scalarFunctionCall.
	VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#udfFunctionCall.
	VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleSpecificFCall.
	VisitSimpleSpecificFCall(ctx *SimpleSpecificFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#convertDataTypeFCall.
	VisitConvertDataTypeFCall(ctx *ConvertDataTypeFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#valuesFCall.
	VisitValuesFCall(ctx *ValuesFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseFCall.
	VisitCaseFCall(ctx *CaseFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#charFCall.
	VisitCharFCall(ctx *CharFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#positionFCall.
	VisitPositionFCall(ctx *PositionFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#substrFCall.
	VisitSubstrFCall(ctx *SubstrFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#trimFCall.
	VisitTrimFCall(ctx *TrimFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#weightFCall.
	VisitWeightFCall(ctx *WeightFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#extractFCall.
	VisitExtractFCall(ctx *ExtractFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#getFormatFCall.
	VisitGetFormatFCall(ctx *GetFormatFCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#levelWeightFList.
	VisitLevelWeightFList(ctx *LevelWeightFListContext) interface{}

	// Visit a parse tree produced by MySqlParser#levelWeightFRange.
	VisitLevelWeightFRange(ctx *LevelWeightFRangeContext) interface{}

	// Visit a parse tree produced by MySqlParser#aggregate_windowed_function.
	VisitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) interface{}

	// Visit a parse tree produced by MySqlParser#scalar_function_name.
	VisitScalar_function_name(ctx *Scalar_function_nameContext) interface{}

	// Visit a parse tree produced by MySqlParser#function_args.
	VisitFunction_args(ctx *Function_argsContext) interface{}

	// Visit a parse tree produced by MySqlParser#function_arg.
	VisitFunction_arg(ctx *Function_argContext) interface{}

	// Visit a parse tree produced by MySqlParser#isExpression.
	VisitIsExpression(ctx *IsExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#soundsLikePredicate.
	VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryComparasionPredicate.
	VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#binaryComparasionPredicate.
	VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#existsExpessionAtom.
	VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#mysqlVariableExpressionAtom.
	VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#binaryExpressionAtom.
	VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullColumnNameExpressionAtom.
	VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#defaultExpressionAtom.
	VisitDefaultExpressionAtom(ctx *DefaultExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#bitExpressionAtom.
	VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#nestedExpressionAtom.
	VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#intervalExpressionAtom.
	VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#comparison_operator.
	VisitComparison_operator(ctx *Comparison_operatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#logical_operator.
	VisitLogical_operator(ctx *Logical_operatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#bit_operator.
	VisitBit_operator(ctx *Bit_operatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#math_operator.
	VisitMath_operator(ctx *Math_operatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#charset_name_base.
	VisitCharset_name_base(ctx *Charset_name_baseContext) interface{}

	// Visit a parse tree produced by MySqlParser#transaction_level_base.
	VisitTransaction_level_base(ctx *Transaction_level_baseContext) interface{}

	// Visit a parse tree produced by MySqlParser#privileges_base.
	VisitPrivileges_base(ctx *Privileges_baseContext) interface{}

	// Visit a parse tree produced by MySqlParser#interval_type_base.
	VisitInterval_type_base(ctx *Interval_type_baseContext) interface{}

	// Visit a parse tree produced by MySqlParser#data_type_base.
	VisitData_type_base(ctx *Data_type_baseContext) interface{}

	// Visit a parse tree produced by MySqlParser#keywords_can_be_id.
	VisitKeywords_can_be_id(ctx *Keywords_can_be_idContext) interface{}

	// Visit a parse tree produced by MySqlParser#function_name_base.
	VisitFunction_name_base(ctx *Function_name_baseContext) interface{}
}
