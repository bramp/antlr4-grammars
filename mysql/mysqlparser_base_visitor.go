// Generated from MySqlParser.g4 by ANTLR 4.7.

package mysql // MySqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMySqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySqlParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSql_statements(ctx *Sql_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSql_statement(ctx *Sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEmpty_statement(ctx *Empty_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDdl_statement(ctx *Ddl_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDml_statement(ctx *Dml_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransaction_statement(ctx *Transaction_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplication_statement(ctx *Replication_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrepared_statement(ctx *Prepared_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCompound_statement(ctx *Compound_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAdministration_statement(ctx *Administration_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUtility_statement(ctx *Utility_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_database(ctx *Create_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_event(ctx *Create_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_index(ctx *Create_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_logfile_group(ctx *Create_logfile_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_procedure(ctx *Create_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_function(ctx *Create_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_server(ctx *Create_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCopyCreateTable(ctx *CopyCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryCreateTable(ctx *QueryCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColCreateTable(ctx *ColCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_tablespace_innodb(ctx *Create_tablespace_innodbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_tablespace_ndb(ctx *Create_tablespace_ndbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_trigger(ctx *Create_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_view(ctx *Create_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_database_option(ctx *Create_database_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOwner_statement(ctx *Owner_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPreciseSchedule(ctx *PreciseScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalSchedule(ctx *IntervalScheduleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTimestamp_value(ctx *Timestamp_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInterval_expr(ctx *Interval_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInterval_type(ctx *Interval_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndex_type(ctx *Index_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndex_option(ctx *Index_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitProc_param(ctx *Proc_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunc_param(ctx *Func_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRcComment(ctx *RcCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRcSqllang(ctx *RcSqllangContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRcDeterm(ctx *RcDetermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRcSqldata(ctx *RcSqldataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRcSecurestmt(ctx *RcSecurestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitServer_option(ctx *Server_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstraintDefinition(ctx *ConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexDefinition(ctx *IndexDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumn_definition(ctx *Column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrNull(ctx *ColConstrNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrDflt(ctx *ColConstrDfltContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrAuInc(ctx *ColConstrAuIncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrPK(ctx *ColConstrPKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrUK(ctx *ColConstrUKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrComment(ctx *ColConstrCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrForm(ctx *ColConstrFormContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrStorage(ctx *ColConstrStorageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColConstrRefdef(ctx *ColConstrRefdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblConstrPK(ctx *TblConstrPKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblConstrUK(ctx *TblConstrUKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblConstrFK(ctx *TblConstrFKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblConstCheck(ctx *TblConstCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReference_definition(ctx *Reference_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOn_delete_action(ctx *On_delete_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOn_update_action(ctx *On_update_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReference_action_control_type(ctx *Reference_action_control_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleIndex(ctx *SimpleIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecIndex(ctx *SpecIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptEngine(ctx *TblOptEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptAuInc(ctx *TblOptAuIncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptAvgRLen(ctx *TblOptAvgRLenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptDefCharSet(ctx *TblOptDefCharSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptChkSum(ctx *TblOptChkSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptDefCollate(ctx *TblOptDefCollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptComment(ctx *TblOptCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptCompr(ctx *TblOptComprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptConn(ctx *TblOptConnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptDataDir(ctx *TblOptDataDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptDelKW(ctx *TblOptDelKWContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptEncr(ctx *TblOptEncrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptIndexDir(ctx *TblOptIndexDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptInsMeth(ctx *TblOptInsMethContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptKeyBlockSz(ctx *TblOptKeyBlockSzContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptMaxRows(ctx *TblOptMaxRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptMinRows(ctx *TblOptMinRowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptPackK(ctx *TblOptPackKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptPasswd(ctx *TblOptPasswdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptRowFormat(ctx *TblOptRowFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptStatAutoR(ctx *TblOptStatAutoRContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptStatPersist(ctx *TblOptStatPersistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptStatSamplPg(ctx *TblOptStatSamplPgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptTablespace(ctx *TblOptTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTblOptUnion(ctx *TblOptUnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartition_options(ctx *Partition_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartition_function_definition(ctx *Partition_function_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLinear_partition_func_def(ctx *Linear_partition_func_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartition_def(ctx *Partition_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubpartition_def(ctx *Subpartition_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterDb(ctx *AlterDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterDbUpgradeName(ctx *AlterDbUpgradeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_event(ctx *Alter_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_function(ctx *Alter_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_instance(ctx *Alter_instanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_logfile_group(ctx *Alter_logfile_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_procedure(ctx *Alter_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_server(ctx *Alter_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_table(ctx *Alter_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_tablespace(ctx *Alter_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlter_view(ctx *Alter_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblTableOpt(ctx *AltblTableOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddCol(ctx *AltblAddColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddCols(ctx *AltblAddColsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddIndex(ctx *AltblAddIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddPK(ctx *AltblAddPKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddUK(ctx *AltblAddUKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddSpecIndex(ctx *AltblAddSpecIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddFK(ctx *AltblAddFKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAlg(ctx *AltblAlgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblColDef(ctx *AltblColDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblColChange(ctx *AltblColChangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblLock(ctx *AltblLockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblColMod(ctx *AltblColModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblColDrop(ctx *AltblColDropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDropPK(ctx *AltblDropPKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDropIndex(ctx *AltblDropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDropFK(ctx *AltblDropFKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDisKey(ctx *AltblDisKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblEnKey(ctx *AltblEnKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblRenameTbl(ctx *AltblRenameTblContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblResort(ctx *AltblResortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblConvert(ctx *AltblConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDefCharset(ctx *AltblDefCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDisTblspace(ctx *AltblDisTblspaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblImpTblSpace(ctx *AltblImpTblSpaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblForce(ctx *AltblForceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblValid(ctx *AltblValidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAddPart(ctx *AltblAddPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDropPart(ctx *AltblDropPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblDiscartPart(ctx *AltblDiscartPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblImportPart(ctx *AltblImportPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblTruncPart(ctx *AltblTruncPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblCoalPart(ctx *AltblCoalPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblReorgPart(ctx *AltblReorgPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblExchPart(ctx *AltblExchPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblAnalPart(ctx *AltblAnalPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblCheckPart(ctx *AltblCheckPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblOptimPart(ctx *AltblOptimPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblRebuildPart(ctx *AltblRebuildPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblRepairPart(ctx *AltblRepairPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblRemovePart(ctx *AltblRemovePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAltblUpgrPart(ctx *AltblUpgrPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_database(ctx *Drop_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_event(ctx *Drop_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_index(ctx *Drop_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_logfile_group(ctx *Drop_logfile_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_procedure(ctx *Drop_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_function(ctx *Drop_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_server(ctx *Drop_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_table(ctx *Drop_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_tablespace(ctx *Drop_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_trigger(ctx *Drop_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_view(ctx *Drop_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRename_table(ctx *Rename_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTruncate_table(ctx *Truncate_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCall_statement(ctx *Call_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDo_statement(ctx *Do_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_statement(ctx *Handler_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsert_statement(ctx *Insert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoad_data_statement(ctx *Load_data_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoad_xml_statement(ctx *Load_xml_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplace_statement(ctx *Replace_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitParenSelect(ctx *ParenSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionParenSelect(ctx *UnionParenSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsert_statement_value(ctx *Insert_statement_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdate_elem(ctx *Update_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCol_or_uservar(ctx *Col_or_uservarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSingle_delete_statement(ctx *Single_delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMultiple_delete_statement(ctx *Multiple_delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_open_statement(ctx *Handler_open_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_read_index_statement(ctx *Handler_read_index_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_read_statement(ctx *Handler_read_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_close_statement(ctx *Handler_close_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSingle_update_statement(ctx *Single_update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMultiple_update_statement(ctx *Multiple_update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrder_by_expression(ctx *Order_by_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTable_sources(ctx *Table_sourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTable_source(ctx *Table_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndex_hint(ctx *Index_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuery_expression(ctx *Query_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuery_expression_nointo(ctx *Query_expression_nointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuery_specification(ctx *Query_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuery_specification_nointo(ctx *Query_specification_nointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnion_parenth(ctx *Union_parenthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnion_statement(ctx *Union_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelect_spec(ctx *Select_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelect_list(ctx *Select_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSellistelAllCol(ctx *SellistelAllColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSellistelCol(ctx *SellistelColContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSellistelFunc(ctx *SellistelFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSellistelExpr(ctx *SellistelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoVars(ctx *SelectIntoVarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoDump(ctx *SelectIntoDumpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoOutfile(ctx *SelectIntoOutfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGroup_by_item(ctx *Group_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStart_transaction(ctx *Start_transactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBegin_work(ctx *Begin_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCommit_work(ctx *Commit_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRollback_work(ctx *Rollback_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSavepoint_statement(ctx *Savepoint_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRollback_statement(ctx *Rollback_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRelease_statement(ctx *Release_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLock_tables(ctx *Lock_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnlock_tables(ctx *Unlock_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSet_autocommit_statement(ctx *Set_autocommit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSet_transaction_statement(ctx *Set_transaction_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransact_option(ctx *Transact_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLock_table_element(ctx *Lock_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTrans_characteristic(ctx *Trans_characteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransaction_level(ctx *Transaction_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChange_master(ctx *Change_masterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChange_repl_filter(ctx *Change_repl_filterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPurge_binary_logs(ctx *Purge_binary_logsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReset_master(ctx *Reset_masterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReset_slave(ctx *Reset_slaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStart_slave(ctx *Start_slaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStop_slave(ctx *Stop_slaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStart_group_repl(ctx *Start_group_replContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStop_group_repl(ctx *Stop_group_replContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterOptString(ctx *MasterOptStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterOptDecimal(ctx *MasterOptDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterOptBool(ctx *MasterOptBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterOptReal(ctx *MasterOptRealContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterOptIdList(ctx *MasterOptIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitString_master_option(ctx *String_master_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDecimal_master_option(ctx *Decimal_master_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBool_master_option(ctx *Bool_master_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChannel_option(ctx *Channel_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplfilterDbList(ctx *ReplfilterDbListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplfilterTableList(ctx *ReplfilterTableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplfilterStableList(ctx *ReplfilterStableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplfilterTablepairList(ctx *ReplfilterTablepairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitThread_type(ctx *Thread_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUntilGtidSset(ctx *UntilGtidSsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUntilMasterLog(ctx *UntilMasterLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUntilRelayLog(ctx *UntilRelayLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUntilSqlGaps(ctx *UntilSqlGapsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStart_slave_connection_option(ctx *Start_slave_connection_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGtid_set(ctx *Gtid_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_start_transaction(ctx *Xa_start_transactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_end_transaction(ctx *Xa_end_transactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_prepare(ctx *Xa_prepareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_commit_work(ctx *Xa_commit_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_rollback_work(ctx *Xa_rollback_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXa_recover_work(ctx *Xa_recover_workContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrepare_statement(ctx *Prepare_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExecute_statement(ctx *Execute_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeallocate_prepare(ctx *Deallocate_prepareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutine_body(ctx *Routine_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBlock_statement(ctx *Block_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIterate_statement(ctx *Iterate_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLeave_statement(ctx *Leave_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRepeat_statement(ctx *Repeat_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCursor_statement(ctx *Cursor_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclare_variable(ctx *Declare_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclare_condition(ctx *Declare_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclare_cursor(ctx *Declare_cursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclare_handler(ctx *Declare_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandler_condition_value(ctx *Handler_condition_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitProcedure_sql_statement(ctx *Procedure_sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterUserMysql56(ctx *AlterUserMysql56Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterUserMysql57(ctx *AlterUserMysql57Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateUserMysql56(ctx *CreateUserMysql56Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateUserMysql57(ctx *CreateUserMysql57Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDrop_user(ctx *Drop_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGrant_statement(ctx *Grant_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGrant_proxy(ctx *Grant_proxyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRename_user(ctx *Rename_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDetailRevoke(ctx *DetailRevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShortRevoke(ctx *ShortRevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRevoke_proxy(ctx *Revoke_proxyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSet_password_statement(ctx *Set_password_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUser_password_option(ctx *User_password_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAuthByPassword(ctx *AuthByPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAuthByString(ctx *AuthByStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAuthByHash(ctx *AuthByHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTls_option(ctx *Tls_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUser_resource_option(ctx *User_resource_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUser_lock_option(ctx *User_lock_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivelege_clause(ctx *Privelege_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivilege_level(ctx *Privilege_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSet_password_option(ctx *Set_password_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAnalyze_table(ctx *Analyze_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheck_table(ctx *Check_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChecksum_table(ctx *Checksum_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOptimize_table(ctx *Optimize_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRepair_table(ctx *Repair_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheck_table_option(ctx *Check_table_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreate_udfunction(ctx *Create_udfunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInstall_plugin(ctx *Install_pluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUninstall_plugin(ctx *Uninstall_pluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetVariableAssignment(ctx *SetVariableAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetCharset(ctx *SetCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetPasswordStatement(ctx *SetPasswordStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetAutocommit(ctx *SetAutocommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowMasterlogs(ctx *ShowMasterlogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowLogevents(ctx *ShowLogeventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowObjWithFilter(ctx *ShowObjWithFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowColumns(ctx *ShowColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateDb(ctx *ShowCreateDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateFullidobj(ctx *ShowCreateFullidobjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateUser(ctx *ShowCreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowEngine(ctx *ShowEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowGlobalinfo(ctx *ShowGlobalinfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowErrWarn(ctx *ShowErrWarnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCountErrWarn(ctx *ShowCountErrWarnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowFromschemaFilter(ctx *ShowFromschemaFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowRoutinecode(ctx *ShowRoutinecodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowGrants(ctx *ShowGrantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowIndexes(ctx *ShowIndexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowOpentables(ctx *ShowOpentablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowProfile(ctx *ShowProfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowSlavestatus(ctx *ShowSlavestatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitVariable_clause(ctx *Variable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShow_filter(ctx *Show_filterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShow_profile_type(ctx *Show_profile_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinlog_statement(ctx *Binlog_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCache_index_statement(ctx *Cache_index_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFlush_statement(ctx *Flush_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitKill_statement(ctx *Kill_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoad_index_into_cache(ctx *Load_index_into_cacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReset_statement(ctx *Reset_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShutdown_statement(ctx *Shutdown_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTbl_index_list(ctx *Tbl_index_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFlush_option(ctx *Flush_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoad_tbl_index_list(ctx *Load_tbl_index_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimple_describe_statement(ctx *Simple_describe_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFull_describe_statement(ctx *Full_describe_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHelp_statement(ctx *Help_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUse_statement(ctx *Use_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDescstmtDescObj(ctx *DescstmtDescObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConnectionDescObj(ctx *ConnectionDescObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFull_id(ctx *Full_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFull_column_name(ctx *Full_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndex_col_name(ctx *Index_col_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUser_name(ctx *User_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysql_variable(ctx *Mysql_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharset_name(ctx *Charset_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEngine_name(ctx *Engine_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUuid_set(ctx *Uuid_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXid(ctx *XidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXid_string_id(ctx *Xid_string_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAuth_plugin(ctx *Auth_pluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitId_(ctx *Id_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimple_id(ctx *Simple_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDot_ext_id(ctx *Dot_ext_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDecimal_literal(ctx *Decimal_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFilesize_literal(ctx *Filesize_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHexadecimal_literal(ctx *Hexadecimal_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNull_notnull(ctx *Null_notnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharDatatype(ctx *CharDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDimensionDatatype(ctx *DimensionDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleDatatype(ctx *SimpleDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollectCharDatatype(ctx *CollectCharDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpatialDatatype(ctx *SpatialDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitData_type_to_convert(ctx *Data_type_to_convertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpatial_data_type(ctx *Spatial_data_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLength_one_dimension(ctx *Length_one_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLength_two_dimension(ctx *Length_two_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLength_two_optional_dimension(ctx *Length_two_optional_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitId_list(ctx *Id_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTable_list(ctx *Table_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTable_pair_list(ctx *Table_pair_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndex_colname_list(ctx *Index_colname_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstant_list(ctx *Constant_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimple_string_list(ctx *Simple_string_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUser_var_list(ctx *User_var_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefault_value(ctx *Default_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIf_exists(ctx *If_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIf_not_exists(ctx *If_not_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleSpecificFCall(ctx *SimpleSpecificFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertDataTypeFCall(ctx *ConvertDataTypeFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitValuesFCall(ctx *ValuesFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFCall(ctx *CaseFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharFCall(ctx *CharFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPositionFCall(ctx *PositionFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubstrFCall(ctx *SubstrFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTrimFCall(ctx *TrimFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWeightFCall(ctx *WeightFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExtractFCall(ctx *ExtractFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGetFormatFCall(ctx *GetFormatFCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightFList(ctx *LevelWeightFListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightFRange(ctx *LevelWeightFRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalar_function_name(ctx *Scalar_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunction_args(ctx *Function_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunction_arg(ctx *Function_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefaultExpressionAtom(ctx *DefaultExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitComparison_operator(ctx *Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogical_operator(ctx *Logical_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBit_operator(ctx *Bit_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMath_operator(ctx *Math_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharset_name_base(ctx *Charset_name_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransaction_level_base(ctx *Transaction_level_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivileges_base(ctx *Privileges_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInterval_type_base(ctx *Interval_type_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitData_type_base(ctx *Data_type_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitKeywords_can_be_id(ctx *Keywords_can_be_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunction_name_base(ctx *Function_name_baseContext) interface{} {
	return v.VisitChildren(ctx)
}
