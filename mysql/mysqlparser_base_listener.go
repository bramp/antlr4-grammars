// Generated from MySqlParser.g4 by ANTLR 4.7.

package mysql // MySqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type BaseMySqlParserListener struct{}

var _ MySqlParserListener = &BaseMySqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMySqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMySqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterSql_statements is called when production sql_statements is entered.
func (s *BaseMySqlParserListener) EnterSql_statements(ctx *Sql_statementsContext) {}

// ExitSql_statements is called when production sql_statements is exited.
func (s *BaseMySqlParserListener) ExitSql_statements(ctx *Sql_statementsContext) {}

// EnterSql_statement is called when production sql_statement is entered.
func (s *BaseMySqlParserListener) EnterSql_statement(ctx *Sql_statementContext) {}

// ExitSql_statement is called when production sql_statement is exited.
func (s *BaseMySqlParserListener) ExitSql_statement(ctx *Sql_statementContext) {}

// EnterEmpty_statement is called when production empty_statement is entered.
func (s *BaseMySqlParserListener) EnterEmpty_statement(ctx *Empty_statementContext) {}

// ExitEmpty_statement is called when production empty_statement is exited.
func (s *BaseMySqlParserListener) ExitEmpty_statement(ctx *Empty_statementContext) {}

// EnterDdl_statement is called when production ddl_statement is entered.
func (s *BaseMySqlParserListener) EnterDdl_statement(ctx *Ddl_statementContext) {}

// ExitDdl_statement is called when production ddl_statement is exited.
func (s *BaseMySqlParserListener) ExitDdl_statement(ctx *Ddl_statementContext) {}

// EnterDml_statement is called when production dml_statement is entered.
func (s *BaseMySqlParserListener) EnterDml_statement(ctx *Dml_statementContext) {}

// ExitDml_statement is called when production dml_statement is exited.
func (s *BaseMySqlParserListener) ExitDml_statement(ctx *Dml_statementContext) {}

// EnterTransaction_statement is called when production transaction_statement is entered.
func (s *BaseMySqlParserListener) EnterTransaction_statement(ctx *Transaction_statementContext) {}

// ExitTransaction_statement is called when production transaction_statement is exited.
func (s *BaseMySqlParserListener) ExitTransaction_statement(ctx *Transaction_statementContext) {}

// EnterReplication_statement is called when production replication_statement is entered.
func (s *BaseMySqlParserListener) EnterReplication_statement(ctx *Replication_statementContext) {}

// ExitReplication_statement is called when production replication_statement is exited.
func (s *BaseMySqlParserListener) ExitReplication_statement(ctx *Replication_statementContext) {}

// EnterPrepared_statement is called when production prepared_statement is entered.
func (s *BaseMySqlParserListener) EnterPrepared_statement(ctx *Prepared_statementContext) {}

// ExitPrepared_statement is called when production prepared_statement is exited.
func (s *BaseMySqlParserListener) ExitPrepared_statement(ctx *Prepared_statementContext) {}

// EnterCompound_statement is called when production compound_statement is entered.
func (s *BaseMySqlParserListener) EnterCompound_statement(ctx *Compound_statementContext) {}

// ExitCompound_statement is called when production compound_statement is exited.
func (s *BaseMySqlParserListener) ExitCompound_statement(ctx *Compound_statementContext) {}

// EnterAdministration_statement is called when production administration_statement is entered.
func (s *BaseMySqlParserListener) EnterAdministration_statement(ctx *Administration_statementContext) {
}

// ExitAdministration_statement is called when production administration_statement is exited.
func (s *BaseMySqlParserListener) ExitAdministration_statement(ctx *Administration_statementContext) {}

// EnterUtility_statement is called when production utility_statement is entered.
func (s *BaseMySqlParserListener) EnterUtility_statement(ctx *Utility_statementContext) {}

// ExitUtility_statement is called when production utility_statement is exited.
func (s *BaseMySqlParserListener) ExitUtility_statement(ctx *Utility_statementContext) {}

// EnterCreate_database is called when production create_database is entered.
func (s *BaseMySqlParserListener) EnterCreate_database(ctx *Create_databaseContext) {}

// ExitCreate_database is called when production create_database is exited.
func (s *BaseMySqlParserListener) ExitCreate_database(ctx *Create_databaseContext) {}

// EnterCreate_event is called when production create_event is entered.
func (s *BaseMySqlParserListener) EnterCreate_event(ctx *Create_eventContext) {}

// ExitCreate_event is called when production create_event is exited.
func (s *BaseMySqlParserListener) ExitCreate_event(ctx *Create_eventContext) {}

// EnterCreate_index is called when production create_index is entered.
func (s *BaseMySqlParserListener) EnterCreate_index(ctx *Create_indexContext) {}

// ExitCreate_index is called when production create_index is exited.
func (s *BaseMySqlParserListener) ExitCreate_index(ctx *Create_indexContext) {}

// EnterCreate_logfile_group is called when production create_logfile_group is entered.
func (s *BaseMySqlParserListener) EnterCreate_logfile_group(ctx *Create_logfile_groupContext) {}

// ExitCreate_logfile_group is called when production create_logfile_group is exited.
func (s *BaseMySqlParserListener) ExitCreate_logfile_group(ctx *Create_logfile_groupContext) {}

// EnterCreate_procedure is called when production create_procedure is entered.
func (s *BaseMySqlParserListener) EnterCreate_procedure(ctx *Create_procedureContext) {}

// ExitCreate_procedure is called when production create_procedure is exited.
func (s *BaseMySqlParserListener) ExitCreate_procedure(ctx *Create_procedureContext) {}

// EnterCreate_function is called when production create_function is entered.
func (s *BaseMySqlParserListener) EnterCreate_function(ctx *Create_functionContext) {}

// ExitCreate_function is called when production create_function is exited.
func (s *BaseMySqlParserListener) ExitCreate_function(ctx *Create_functionContext) {}

// EnterCreate_server is called when production create_server is entered.
func (s *BaseMySqlParserListener) EnterCreate_server(ctx *Create_serverContext) {}

// ExitCreate_server is called when production create_server is exited.
func (s *BaseMySqlParserListener) ExitCreate_server(ctx *Create_serverContext) {}

// EnterCopyCreateTable is called when production copyCreateTable is entered.
func (s *BaseMySqlParserListener) EnterCopyCreateTable(ctx *CopyCreateTableContext) {}

// ExitCopyCreateTable is called when production copyCreateTable is exited.
func (s *BaseMySqlParserListener) ExitCopyCreateTable(ctx *CopyCreateTableContext) {}

// EnterQueryCreateTable is called when production queryCreateTable is entered.
func (s *BaseMySqlParserListener) EnterQueryCreateTable(ctx *QueryCreateTableContext) {}

// ExitQueryCreateTable is called when production queryCreateTable is exited.
func (s *BaseMySqlParserListener) ExitQueryCreateTable(ctx *QueryCreateTableContext) {}

// EnterColCreateTable is called when production colCreateTable is entered.
func (s *BaseMySqlParserListener) EnterColCreateTable(ctx *ColCreateTableContext) {}

// ExitColCreateTable is called when production colCreateTable is exited.
func (s *BaseMySqlParserListener) ExitColCreateTable(ctx *ColCreateTableContext) {}

// EnterCreate_tablespace_innodb is called when production create_tablespace_innodb is entered.
func (s *BaseMySqlParserListener) EnterCreate_tablespace_innodb(ctx *Create_tablespace_innodbContext) {
}

// ExitCreate_tablespace_innodb is called when production create_tablespace_innodb is exited.
func (s *BaseMySqlParserListener) ExitCreate_tablespace_innodb(ctx *Create_tablespace_innodbContext) {}

// EnterCreate_tablespace_ndb is called when production create_tablespace_ndb is entered.
func (s *BaseMySqlParserListener) EnterCreate_tablespace_ndb(ctx *Create_tablespace_ndbContext) {}

// ExitCreate_tablespace_ndb is called when production create_tablespace_ndb is exited.
func (s *BaseMySqlParserListener) ExitCreate_tablespace_ndb(ctx *Create_tablespace_ndbContext) {}

// EnterCreate_trigger is called when production create_trigger is entered.
func (s *BaseMySqlParserListener) EnterCreate_trigger(ctx *Create_triggerContext) {}

// ExitCreate_trigger is called when production create_trigger is exited.
func (s *BaseMySqlParserListener) ExitCreate_trigger(ctx *Create_triggerContext) {}

// EnterCreate_view is called when production create_view is entered.
func (s *BaseMySqlParserListener) EnterCreate_view(ctx *Create_viewContext) {}

// ExitCreate_view is called when production create_view is exited.
func (s *BaseMySqlParserListener) ExitCreate_view(ctx *Create_viewContext) {}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *BaseMySqlParserListener) EnterCreate_database_option(ctx *Create_database_optionContext) {}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *BaseMySqlParserListener) ExitCreate_database_option(ctx *Create_database_optionContext) {}

// EnterOwner_statement is called when production owner_statement is entered.
func (s *BaseMySqlParserListener) EnterOwner_statement(ctx *Owner_statementContext) {}

// ExitOwner_statement is called when production owner_statement is exited.
func (s *BaseMySqlParserListener) ExitOwner_statement(ctx *Owner_statementContext) {}

// EnterPreciseSchedule is called when production preciseSchedule is entered.
func (s *BaseMySqlParserListener) EnterPreciseSchedule(ctx *PreciseScheduleContext) {}

// ExitPreciseSchedule is called when production preciseSchedule is exited.
func (s *BaseMySqlParserListener) ExitPreciseSchedule(ctx *PreciseScheduleContext) {}

// EnterIntervalSchedule is called when production intervalSchedule is entered.
func (s *BaseMySqlParserListener) EnterIntervalSchedule(ctx *IntervalScheduleContext) {}

// ExitIntervalSchedule is called when production intervalSchedule is exited.
func (s *BaseMySqlParserListener) ExitIntervalSchedule(ctx *IntervalScheduleContext) {}

// EnterTimestamp_value is called when production timestamp_value is entered.
func (s *BaseMySqlParserListener) EnterTimestamp_value(ctx *Timestamp_valueContext) {}

// ExitTimestamp_value is called when production timestamp_value is exited.
func (s *BaseMySqlParserListener) ExitTimestamp_value(ctx *Timestamp_valueContext) {}

// EnterInterval_expr is called when production interval_expr is entered.
func (s *BaseMySqlParserListener) EnterInterval_expr(ctx *Interval_exprContext) {}

// ExitInterval_expr is called when production interval_expr is exited.
func (s *BaseMySqlParserListener) ExitInterval_expr(ctx *Interval_exprContext) {}

// EnterInterval_type is called when production interval_type is entered.
func (s *BaseMySqlParserListener) EnterInterval_type(ctx *Interval_typeContext) {}

// ExitInterval_type is called when production interval_type is exited.
func (s *BaseMySqlParserListener) ExitInterval_type(ctx *Interval_typeContext) {}

// EnterIndex_type is called when production index_type is entered.
func (s *BaseMySqlParserListener) EnterIndex_type(ctx *Index_typeContext) {}

// ExitIndex_type is called when production index_type is exited.
func (s *BaseMySqlParserListener) ExitIndex_type(ctx *Index_typeContext) {}

// EnterIndex_option is called when production index_option is entered.
func (s *BaseMySqlParserListener) EnterIndex_option(ctx *Index_optionContext) {}

// ExitIndex_option is called when production index_option is exited.
func (s *BaseMySqlParserListener) ExitIndex_option(ctx *Index_optionContext) {}

// EnterProc_param is called when production proc_param is entered.
func (s *BaseMySqlParserListener) EnterProc_param(ctx *Proc_paramContext) {}

// ExitProc_param is called when production proc_param is exited.
func (s *BaseMySqlParserListener) ExitProc_param(ctx *Proc_paramContext) {}

// EnterFunc_param is called when production func_param is entered.
func (s *BaseMySqlParserListener) EnterFunc_param(ctx *Func_paramContext) {}

// ExitFunc_param is called when production func_param is exited.
func (s *BaseMySqlParserListener) ExitFunc_param(ctx *Func_paramContext) {}

// EnterRcComment is called when production rcComment is entered.
func (s *BaseMySqlParserListener) EnterRcComment(ctx *RcCommentContext) {}

// ExitRcComment is called when production rcComment is exited.
func (s *BaseMySqlParserListener) ExitRcComment(ctx *RcCommentContext) {}

// EnterRcSqllang is called when production rcSqllang is entered.
func (s *BaseMySqlParserListener) EnterRcSqllang(ctx *RcSqllangContext) {}

// ExitRcSqllang is called when production rcSqllang is exited.
func (s *BaseMySqlParserListener) ExitRcSqllang(ctx *RcSqllangContext) {}

// EnterRcDeterm is called when production rcDeterm is entered.
func (s *BaseMySqlParserListener) EnterRcDeterm(ctx *RcDetermContext) {}

// ExitRcDeterm is called when production rcDeterm is exited.
func (s *BaseMySqlParserListener) ExitRcDeterm(ctx *RcDetermContext) {}

// EnterRcSqldata is called when production rcSqldata is entered.
func (s *BaseMySqlParserListener) EnterRcSqldata(ctx *RcSqldataContext) {}

// ExitRcSqldata is called when production rcSqldata is exited.
func (s *BaseMySqlParserListener) ExitRcSqldata(ctx *RcSqldataContext) {}

// EnterRcSecurestmt is called when production rcSecurestmt is entered.
func (s *BaseMySqlParserListener) EnterRcSecurestmt(ctx *RcSecurestmtContext) {}

// ExitRcSecurestmt is called when production rcSecurestmt is exited.
func (s *BaseMySqlParserListener) ExitRcSecurestmt(ctx *RcSecurestmtContext) {}

// EnterServer_option is called when production server_option is entered.
func (s *BaseMySqlParserListener) EnterServer_option(ctx *Server_optionContext) {}

// ExitServer_option is called when production server_option is exited.
func (s *BaseMySqlParserListener) ExitServer_option(ctx *Server_optionContext) {}

// EnterColumn_def_table_constraints is called when production column_def_table_constraints is entered.
func (s *BaseMySqlParserListener) EnterColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// ExitColumn_def_table_constraints is called when production column_def_table_constraints is exited.
func (s *BaseMySqlParserListener) ExitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseMySqlParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseMySqlParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterConstraintDefinition is called when production constraintDefinition is entered.
func (s *BaseMySqlParserListener) EnterConstraintDefinition(ctx *ConstraintDefinitionContext) {}

// ExitConstraintDefinition is called when production constraintDefinition is exited.
func (s *BaseMySqlParserListener) ExitConstraintDefinition(ctx *ConstraintDefinitionContext) {}

// EnterIndexDefinition is called when production indexDefinition is entered.
func (s *BaseMySqlParserListener) EnterIndexDefinition(ctx *IndexDefinitionContext) {}

// ExitIndexDefinition is called when production indexDefinition is exited.
func (s *BaseMySqlParserListener) ExitIndexDefinition(ctx *IndexDefinitionContext) {}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BaseMySqlParserListener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BaseMySqlParserListener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterColConstrNull is called when production colConstrNull is entered.
func (s *BaseMySqlParserListener) EnterColConstrNull(ctx *ColConstrNullContext) {}

// ExitColConstrNull is called when production colConstrNull is exited.
func (s *BaseMySqlParserListener) ExitColConstrNull(ctx *ColConstrNullContext) {}

// EnterColConstrDflt is called when production colConstrDflt is entered.
func (s *BaseMySqlParserListener) EnterColConstrDflt(ctx *ColConstrDfltContext) {}

// ExitColConstrDflt is called when production colConstrDflt is exited.
func (s *BaseMySqlParserListener) ExitColConstrDflt(ctx *ColConstrDfltContext) {}

// EnterColConstrAuInc is called when production colConstrAuInc is entered.
func (s *BaseMySqlParserListener) EnterColConstrAuInc(ctx *ColConstrAuIncContext) {}

// ExitColConstrAuInc is called when production colConstrAuInc is exited.
func (s *BaseMySqlParserListener) ExitColConstrAuInc(ctx *ColConstrAuIncContext) {}

// EnterColConstrPK is called when production colConstrPK is entered.
func (s *BaseMySqlParserListener) EnterColConstrPK(ctx *ColConstrPKContext) {}

// ExitColConstrPK is called when production colConstrPK is exited.
func (s *BaseMySqlParserListener) ExitColConstrPK(ctx *ColConstrPKContext) {}

// EnterColConstrUK is called when production colConstrUK is entered.
func (s *BaseMySqlParserListener) EnterColConstrUK(ctx *ColConstrUKContext) {}

// ExitColConstrUK is called when production colConstrUK is exited.
func (s *BaseMySqlParserListener) ExitColConstrUK(ctx *ColConstrUKContext) {}

// EnterColConstrComment is called when production colConstrComment is entered.
func (s *BaseMySqlParserListener) EnterColConstrComment(ctx *ColConstrCommentContext) {}

// ExitColConstrComment is called when production colConstrComment is exited.
func (s *BaseMySqlParserListener) ExitColConstrComment(ctx *ColConstrCommentContext) {}

// EnterColConstrForm is called when production colConstrForm is entered.
func (s *BaseMySqlParserListener) EnterColConstrForm(ctx *ColConstrFormContext) {}

// ExitColConstrForm is called when production colConstrForm is exited.
func (s *BaseMySqlParserListener) ExitColConstrForm(ctx *ColConstrFormContext) {}

// EnterColConstrStorage is called when production colConstrStorage is entered.
func (s *BaseMySqlParserListener) EnterColConstrStorage(ctx *ColConstrStorageContext) {}

// ExitColConstrStorage is called when production colConstrStorage is exited.
func (s *BaseMySqlParserListener) ExitColConstrStorage(ctx *ColConstrStorageContext) {}

// EnterColConstrRefdef is called when production colConstrRefdef is entered.
func (s *BaseMySqlParserListener) EnterColConstrRefdef(ctx *ColConstrRefdefContext) {}

// ExitColConstrRefdef is called when production colConstrRefdef is exited.
func (s *BaseMySqlParserListener) ExitColConstrRefdef(ctx *ColConstrRefdefContext) {}

// EnterTblConstrPK is called when production tblConstrPK is entered.
func (s *BaseMySqlParserListener) EnterTblConstrPK(ctx *TblConstrPKContext) {}

// ExitTblConstrPK is called when production tblConstrPK is exited.
func (s *BaseMySqlParserListener) ExitTblConstrPK(ctx *TblConstrPKContext) {}

// EnterTblConstrUK is called when production tblConstrUK is entered.
func (s *BaseMySqlParserListener) EnterTblConstrUK(ctx *TblConstrUKContext) {}

// ExitTblConstrUK is called when production tblConstrUK is exited.
func (s *BaseMySqlParserListener) ExitTblConstrUK(ctx *TblConstrUKContext) {}

// EnterTblConstrFK is called when production tblConstrFK is entered.
func (s *BaseMySqlParserListener) EnterTblConstrFK(ctx *TblConstrFKContext) {}

// ExitTblConstrFK is called when production tblConstrFK is exited.
func (s *BaseMySqlParserListener) ExitTblConstrFK(ctx *TblConstrFKContext) {}

// EnterTblConstCheck is called when production tblConstCheck is entered.
func (s *BaseMySqlParserListener) EnterTblConstCheck(ctx *TblConstCheckContext) {}

// ExitTblConstCheck is called when production tblConstCheck is exited.
func (s *BaseMySqlParserListener) ExitTblConstCheck(ctx *TblConstCheckContext) {}

// EnterReference_definition is called when production reference_definition is entered.
func (s *BaseMySqlParserListener) EnterReference_definition(ctx *Reference_definitionContext) {}

// ExitReference_definition is called when production reference_definition is exited.
func (s *BaseMySqlParserListener) ExitReference_definition(ctx *Reference_definitionContext) {}

// EnterOn_delete_action is called when production on_delete_action is entered.
func (s *BaseMySqlParserListener) EnterOn_delete_action(ctx *On_delete_actionContext) {}

// ExitOn_delete_action is called when production on_delete_action is exited.
func (s *BaseMySqlParserListener) ExitOn_delete_action(ctx *On_delete_actionContext) {}

// EnterOn_update_action is called when production on_update_action is entered.
func (s *BaseMySqlParserListener) EnterOn_update_action(ctx *On_update_actionContext) {}

// ExitOn_update_action is called when production on_update_action is exited.
func (s *BaseMySqlParserListener) ExitOn_update_action(ctx *On_update_actionContext) {}

// EnterReference_action_control_type is called when production reference_action_control_type is entered.
func (s *BaseMySqlParserListener) EnterReference_action_control_type(ctx *Reference_action_control_typeContext) {
}

// ExitReference_action_control_type is called when production reference_action_control_type is exited.
func (s *BaseMySqlParserListener) ExitReference_action_control_type(ctx *Reference_action_control_typeContext) {
}

// EnterSimpleIndex is called when production simpleIndex is entered.
func (s *BaseMySqlParserListener) EnterSimpleIndex(ctx *SimpleIndexContext) {}

// ExitSimpleIndex is called when production simpleIndex is exited.
func (s *BaseMySqlParserListener) ExitSimpleIndex(ctx *SimpleIndexContext) {}

// EnterSpecIndex is called when production specIndex is entered.
func (s *BaseMySqlParserListener) EnterSpecIndex(ctx *SpecIndexContext) {}

// ExitSpecIndex is called when production specIndex is exited.
func (s *BaseMySqlParserListener) ExitSpecIndex(ctx *SpecIndexContext) {}

// EnterTblOptEngine is called when production tblOptEngine is entered.
func (s *BaseMySqlParserListener) EnterTblOptEngine(ctx *TblOptEngineContext) {}

// ExitTblOptEngine is called when production tblOptEngine is exited.
func (s *BaseMySqlParserListener) ExitTblOptEngine(ctx *TblOptEngineContext) {}

// EnterTblOptAuInc is called when production tblOptAuInc is entered.
func (s *BaseMySqlParserListener) EnterTblOptAuInc(ctx *TblOptAuIncContext) {}

// ExitTblOptAuInc is called when production tblOptAuInc is exited.
func (s *BaseMySqlParserListener) ExitTblOptAuInc(ctx *TblOptAuIncContext) {}

// EnterTblOptAvgRLen is called when production tblOptAvgRLen is entered.
func (s *BaseMySqlParserListener) EnterTblOptAvgRLen(ctx *TblOptAvgRLenContext) {}

// ExitTblOptAvgRLen is called when production tblOptAvgRLen is exited.
func (s *BaseMySqlParserListener) ExitTblOptAvgRLen(ctx *TblOptAvgRLenContext) {}

// EnterTblOptDefCharSet is called when production tblOptDefCharSet is entered.
func (s *BaseMySqlParserListener) EnterTblOptDefCharSet(ctx *TblOptDefCharSetContext) {}

// ExitTblOptDefCharSet is called when production tblOptDefCharSet is exited.
func (s *BaseMySqlParserListener) ExitTblOptDefCharSet(ctx *TblOptDefCharSetContext) {}

// EnterTblOptChkSum is called when production tblOptChkSum is entered.
func (s *BaseMySqlParserListener) EnterTblOptChkSum(ctx *TblOptChkSumContext) {}

// ExitTblOptChkSum is called when production tblOptChkSum is exited.
func (s *BaseMySqlParserListener) ExitTblOptChkSum(ctx *TblOptChkSumContext) {}

// EnterTblOptDefCollate is called when production tblOptDefCollate is entered.
func (s *BaseMySqlParserListener) EnterTblOptDefCollate(ctx *TblOptDefCollateContext) {}

// ExitTblOptDefCollate is called when production tblOptDefCollate is exited.
func (s *BaseMySqlParserListener) ExitTblOptDefCollate(ctx *TblOptDefCollateContext) {}

// EnterTblOptComment is called when production tblOptComment is entered.
func (s *BaseMySqlParserListener) EnterTblOptComment(ctx *TblOptCommentContext) {}

// ExitTblOptComment is called when production tblOptComment is exited.
func (s *BaseMySqlParserListener) ExitTblOptComment(ctx *TblOptCommentContext) {}

// EnterTblOptCompr is called when production tblOptCompr is entered.
func (s *BaseMySqlParserListener) EnterTblOptCompr(ctx *TblOptComprContext) {}

// ExitTblOptCompr is called when production tblOptCompr is exited.
func (s *BaseMySqlParserListener) ExitTblOptCompr(ctx *TblOptComprContext) {}

// EnterTblOptConn is called when production tblOptConn is entered.
func (s *BaseMySqlParserListener) EnterTblOptConn(ctx *TblOptConnContext) {}

// ExitTblOptConn is called when production tblOptConn is exited.
func (s *BaseMySqlParserListener) ExitTblOptConn(ctx *TblOptConnContext) {}

// EnterTblOptDataDir is called when production tblOptDataDir is entered.
func (s *BaseMySqlParserListener) EnterTblOptDataDir(ctx *TblOptDataDirContext) {}

// ExitTblOptDataDir is called when production tblOptDataDir is exited.
func (s *BaseMySqlParserListener) ExitTblOptDataDir(ctx *TblOptDataDirContext) {}

// EnterTblOptDelKW is called when production tblOptDelKW is entered.
func (s *BaseMySqlParserListener) EnterTblOptDelKW(ctx *TblOptDelKWContext) {}

// ExitTblOptDelKW is called when production tblOptDelKW is exited.
func (s *BaseMySqlParserListener) ExitTblOptDelKW(ctx *TblOptDelKWContext) {}

// EnterTblOptEncr is called when production tblOptEncr is entered.
func (s *BaseMySqlParserListener) EnterTblOptEncr(ctx *TblOptEncrContext) {}

// ExitTblOptEncr is called when production tblOptEncr is exited.
func (s *BaseMySqlParserListener) ExitTblOptEncr(ctx *TblOptEncrContext) {}

// EnterTblOptIndexDir is called when production tblOptIndexDir is entered.
func (s *BaseMySqlParserListener) EnterTblOptIndexDir(ctx *TblOptIndexDirContext) {}

// ExitTblOptIndexDir is called when production tblOptIndexDir is exited.
func (s *BaseMySqlParserListener) ExitTblOptIndexDir(ctx *TblOptIndexDirContext) {}

// EnterTblOptInsMeth is called when production tblOptInsMeth is entered.
func (s *BaseMySqlParserListener) EnterTblOptInsMeth(ctx *TblOptInsMethContext) {}

// ExitTblOptInsMeth is called when production tblOptInsMeth is exited.
func (s *BaseMySqlParserListener) ExitTblOptInsMeth(ctx *TblOptInsMethContext) {}

// EnterTblOptKeyBlockSz is called when production tblOptKeyBlockSz is entered.
func (s *BaseMySqlParserListener) EnterTblOptKeyBlockSz(ctx *TblOptKeyBlockSzContext) {}

// ExitTblOptKeyBlockSz is called when production tblOptKeyBlockSz is exited.
func (s *BaseMySqlParserListener) ExitTblOptKeyBlockSz(ctx *TblOptKeyBlockSzContext) {}

// EnterTblOptMaxRows is called when production tblOptMaxRows is entered.
func (s *BaseMySqlParserListener) EnterTblOptMaxRows(ctx *TblOptMaxRowsContext) {}

// ExitTblOptMaxRows is called when production tblOptMaxRows is exited.
func (s *BaseMySqlParserListener) ExitTblOptMaxRows(ctx *TblOptMaxRowsContext) {}

// EnterTblOptMinRows is called when production tblOptMinRows is entered.
func (s *BaseMySqlParserListener) EnterTblOptMinRows(ctx *TblOptMinRowsContext) {}

// ExitTblOptMinRows is called when production tblOptMinRows is exited.
func (s *BaseMySqlParserListener) ExitTblOptMinRows(ctx *TblOptMinRowsContext) {}

// EnterTblOptPackK is called when production tblOptPackK is entered.
func (s *BaseMySqlParserListener) EnterTblOptPackK(ctx *TblOptPackKContext) {}

// ExitTblOptPackK is called when production tblOptPackK is exited.
func (s *BaseMySqlParserListener) ExitTblOptPackK(ctx *TblOptPackKContext) {}

// EnterTblOptPasswd is called when production tblOptPasswd is entered.
func (s *BaseMySqlParserListener) EnterTblOptPasswd(ctx *TblOptPasswdContext) {}

// ExitTblOptPasswd is called when production tblOptPasswd is exited.
func (s *BaseMySqlParserListener) ExitTblOptPasswd(ctx *TblOptPasswdContext) {}

// EnterTblOptRowFormat is called when production tblOptRowFormat is entered.
func (s *BaseMySqlParserListener) EnterTblOptRowFormat(ctx *TblOptRowFormatContext) {}

// ExitTblOptRowFormat is called when production tblOptRowFormat is exited.
func (s *BaseMySqlParserListener) ExitTblOptRowFormat(ctx *TblOptRowFormatContext) {}

// EnterTblOptStatAutoR is called when production tblOptStatAutoR is entered.
func (s *BaseMySqlParserListener) EnterTblOptStatAutoR(ctx *TblOptStatAutoRContext) {}

// ExitTblOptStatAutoR is called when production tblOptStatAutoR is exited.
func (s *BaseMySqlParserListener) ExitTblOptStatAutoR(ctx *TblOptStatAutoRContext) {}

// EnterTblOptStatPersist is called when production tblOptStatPersist is entered.
func (s *BaseMySqlParserListener) EnterTblOptStatPersist(ctx *TblOptStatPersistContext) {}

// ExitTblOptStatPersist is called when production tblOptStatPersist is exited.
func (s *BaseMySqlParserListener) ExitTblOptStatPersist(ctx *TblOptStatPersistContext) {}

// EnterTblOptStatSamplPg is called when production tblOptStatSamplPg is entered.
func (s *BaseMySqlParserListener) EnterTblOptStatSamplPg(ctx *TblOptStatSamplPgContext) {}

// ExitTblOptStatSamplPg is called when production tblOptStatSamplPg is exited.
func (s *BaseMySqlParserListener) ExitTblOptStatSamplPg(ctx *TblOptStatSamplPgContext) {}

// EnterTblOptTablespace is called when production tblOptTablespace is entered.
func (s *BaseMySqlParserListener) EnterTblOptTablespace(ctx *TblOptTablespaceContext) {}

// ExitTblOptTablespace is called when production tblOptTablespace is exited.
func (s *BaseMySqlParserListener) ExitTblOptTablespace(ctx *TblOptTablespaceContext) {}

// EnterTblOptUnion is called when production tblOptUnion is entered.
func (s *BaseMySqlParserListener) EnterTblOptUnion(ctx *TblOptUnionContext) {}

// ExitTblOptUnion is called when production tblOptUnion is exited.
func (s *BaseMySqlParserListener) ExitTblOptUnion(ctx *TblOptUnionContext) {}

// EnterPartition_options is called when production partition_options is entered.
func (s *BaseMySqlParserListener) EnterPartition_options(ctx *Partition_optionsContext) {}

// ExitPartition_options is called when production partition_options is exited.
func (s *BaseMySqlParserListener) ExitPartition_options(ctx *Partition_optionsContext) {}

// EnterPartition_function_definition is called when production partition_function_definition is entered.
func (s *BaseMySqlParserListener) EnterPartition_function_definition(ctx *Partition_function_definitionContext) {
}

// ExitPartition_function_definition is called when production partition_function_definition is exited.
func (s *BaseMySqlParserListener) ExitPartition_function_definition(ctx *Partition_function_definitionContext) {
}

// EnterLinear_partition_func_def is called when production linear_partition_func_def is entered.
func (s *BaseMySqlParserListener) EnterLinear_partition_func_def(ctx *Linear_partition_func_defContext) {
}

// ExitLinear_partition_func_def is called when production linear_partition_func_def is exited.
func (s *BaseMySqlParserListener) ExitLinear_partition_func_def(ctx *Linear_partition_func_defContext) {
}

// EnterPartition_def is called when production partition_def is entered.
func (s *BaseMySqlParserListener) EnterPartition_def(ctx *Partition_defContext) {}

// ExitPartition_def is called when production partition_def is exited.
func (s *BaseMySqlParserListener) ExitPartition_def(ctx *Partition_defContext) {}

// EnterSubpartition_def is called when production subpartition_def is entered.
func (s *BaseMySqlParserListener) EnterSubpartition_def(ctx *Subpartition_defContext) {}

// ExitSubpartition_def is called when production subpartition_def is exited.
func (s *BaseMySqlParserListener) ExitSubpartition_def(ctx *Subpartition_defContext) {}

// EnterAlterDb is called when production alterDb is entered.
func (s *BaseMySqlParserListener) EnterAlterDb(ctx *AlterDbContext) {}

// ExitAlterDb is called when production alterDb is exited.
func (s *BaseMySqlParserListener) ExitAlterDb(ctx *AlterDbContext) {}

// EnterAlterDbUpgradeName is called when production alterDbUpgradeName is entered.
func (s *BaseMySqlParserListener) EnterAlterDbUpgradeName(ctx *AlterDbUpgradeNameContext) {}

// ExitAlterDbUpgradeName is called when production alterDbUpgradeName is exited.
func (s *BaseMySqlParserListener) ExitAlterDbUpgradeName(ctx *AlterDbUpgradeNameContext) {}

// EnterAlter_event is called when production alter_event is entered.
func (s *BaseMySqlParserListener) EnterAlter_event(ctx *Alter_eventContext) {}

// ExitAlter_event is called when production alter_event is exited.
func (s *BaseMySqlParserListener) ExitAlter_event(ctx *Alter_eventContext) {}

// EnterAlter_function is called when production alter_function is entered.
func (s *BaseMySqlParserListener) EnterAlter_function(ctx *Alter_functionContext) {}

// ExitAlter_function is called when production alter_function is exited.
func (s *BaseMySqlParserListener) ExitAlter_function(ctx *Alter_functionContext) {}

// EnterAlter_instance is called when production alter_instance is entered.
func (s *BaseMySqlParserListener) EnterAlter_instance(ctx *Alter_instanceContext) {}

// ExitAlter_instance is called when production alter_instance is exited.
func (s *BaseMySqlParserListener) ExitAlter_instance(ctx *Alter_instanceContext) {}

// EnterAlter_logfile_group is called when production alter_logfile_group is entered.
func (s *BaseMySqlParserListener) EnterAlter_logfile_group(ctx *Alter_logfile_groupContext) {}

// ExitAlter_logfile_group is called when production alter_logfile_group is exited.
func (s *BaseMySqlParserListener) ExitAlter_logfile_group(ctx *Alter_logfile_groupContext) {}

// EnterAlter_procedure is called when production alter_procedure is entered.
func (s *BaseMySqlParserListener) EnterAlter_procedure(ctx *Alter_procedureContext) {}

// ExitAlter_procedure is called when production alter_procedure is exited.
func (s *BaseMySqlParserListener) ExitAlter_procedure(ctx *Alter_procedureContext) {}

// EnterAlter_server is called when production alter_server is entered.
func (s *BaseMySqlParserListener) EnterAlter_server(ctx *Alter_serverContext) {}

// ExitAlter_server is called when production alter_server is exited.
func (s *BaseMySqlParserListener) ExitAlter_server(ctx *Alter_serverContext) {}

// EnterAlter_table is called when production alter_table is entered.
func (s *BaseMySqlParserListener) EnterAlter_table(ctx *Alter_tableContext) {}

// ExitAlter_table is called when production alter_table is exited.
func (s *BaseMySqlParserListener) ExitAlter_table(ctx *Alter_tableContext) {}

// EnterAlter_tablespace is called when production alter_tablespace is entered.
func (s *BaseMySqlParserListener) EnterAlter_tablespace(ctx *Alter_tablespaceContext) {}

// ExitAlter_tablespace is called when production alter_tablespace is exited.
func (s *BaseMySqlParserListener) ExitAlter_tablespace(ctx *Alter_tablespaceContext) {}

// EnterAlter_view is called when production alter_view is entered.
func (s *BaseMySqlParserListener) EnterAlter_view(ctx *Alter_viewContext) {}

// ExitAlter_view is called when production alter_view is exited.
func (s *BaseMySqlParserListener) ExitAlter_view(ctx *Alter_viewContext) {}

// EnterAltblTableOpt is called when production altblTableOpt is entered.
func (s *BaseMySqlParserListener) EnterAltblTableOpt(ctx *AltblTableOptContext) {}

// ExitAltblTableOpt is called when production altblTableOpt is exited.
func (s *BaseMySqlParserListener) ExitAltblTableOpt(ctx *AltblTableOptContext) {}

// EnterAltblAddCol is called when production altblAddCol is entered.
func (s *BaseMySqlParserListener) EnterAltblAddCol(ctx *AltblAddColContext) {}

// ExitAltblAddCol is called when production altblAddCol is exited.
func (s *BaseMySqlParserListener) ExitAltblAddCol(ctx *AltblAddColContext) {}

// EnterAltblAddCols is called when production altblAddCols is entered.
func (s *BaseMySqlParserListener) EnterAltblAddCols(ctx *AltblAddColsContext) {}

// ExitAltblAddCols is called when production altblAddCols is exited.
func (s *BaseMySqlParserListener) ExitAltblAddCols(ctx *AltblAddColsContext) {}

// EnterAltblAddIndex is called when production altblAddIndex is entered.
func (s *BaseMySqlParserListener) EnterAltblAddIndex(ctx *AltblAddIndexContext) {}

// ExitAltblAddIndex is called when production altblAddIndex is exited.
func (s *BaseMySqlParserListener) ExitAltblAddIndex(ctx *AltblAddIndexContext) {}

// EnterAltblAddPK is called when production altblAddPK is entered.
func (s *BaseMySqlParserListener) EnterAltblAddPK(ctx *AltblAddPKContext) {}

// ExitAltblAddPK is called when production altblAddPK is exited.
func (s *BaseMySqlParserListener) ExitAltblAddPK(ctx *AltblAddPKContext) {}

// EnterAltblAddUK is called when production altblAddUK is entered.
func (s *BaseMySqlParserListener) EnterAltblAddUK(ctx *AltblAddUKContext) {}

// ExitAltblAddUK is called when production altblAddUK is exited.
func (s *BaseMySqlParserListener) ExitAltblAddUK(ctx *AltblAddUKContext) {}

// EnterAltblAddSpecIndex is called when production altblAddSpecIndex is entered.
func (s *BaseMySqlParserListener) EnterAltblAddSpecIndex(ctx *AltblAddSpecIndexContext) {}

// ExitAltblAddSpecIndex is called when production altblAddSpecIndex is exited.
func (s *BaseMySqlParserListener) ExitAltblAddSpecIndex(ctx *AltblAddSpecIndexContext) {}

// EnterAltblAddFK is called when production altblAddFK is entered.
func (s *BaseMySqlParserListener) EnterAltblAddFK(ctx *AltblAddFKContext) {}

// ExitAltblAddFK is called when production altblAddFK is exited.
func (s *BaseMySqlParserListener) ExitAltblAddFK(ctx *AltblAddFKContext) {}

// EnterAltblAlg is called when production altblAlg is entered.
func (s *BaseMySqlParserListener) EnterAltblAlg(ctx *AltblAlgContext) {}

// ExitAltblAlg is called when production altblAlg is exited.
func (s *BaseMySqlParserListener) ExitAltblAlg(ctx *AltblAlgContext) {}

// EnterAltblColDef is called when production altblColDef is entered.
func (s *BaseMySqlParserListener) EnterAltblColDef(ctx *AltblColDefContext) {}

// ExitAltblColDef is called when production altblColDef is exited.
func (s *BaseMySqlParserListener) ExitAltblColDef(ctx *AltblColDefContext) {}

// EnterAltblColChange is called when production altblColChange is entered.
func (s *BaseMySqlParserListener) EnterAltblColChange(ctx *AltblColChangeContext) {}

// ExitAltblColChange is called when production altblColChange is exited.
func (s *BaseMySqlParserListener) ExitAltblColChange(ctx *AltblColChangeContext) {}

// EnterAltblLock is called when production altblLock is entered.
func (s *BaseMySqlParserListener) EnterAltblLock(ctx *AltblLockContext) {}

// ExitAltblLock is called when production altblLock is exited.
func (s *BaseMySqlParserListener) ExitAltblLock(ctx *AltblLockContext) {}

// EnterAltblColMod is called when production altblColMod is entered.
func (s *BaseMySqlParserListener) EnterAltblColMod(ctx *AltblColModContext) {}

// ExitAltblColMod is called when production altblColMod is exited.
func (s *BaseMySqlParserListener) ExitAltblColMod(ctx *AltblColModContext) {}

// EnterAltblColDrop is called when production altblColDrop is entered.
func (s *BaseMySqlParserListener) EnterAltblColDrop(ctx *AltblColDropContext) {}

// ExitAltblColDrop is called when production altblColDrop is exited.
func (s *BaseMySqlParserListener) ExitAltblColDrop(ctx *AltblColDropContext) {}

// EnterAltblDropPK is called when production altblDropPK is entered.
func (s *BaseMySqlParserListener) EnterAltblDropPK(ctx *AltblDropPKContext) {}

// ExitAltblDropPK is called when production altblDropPK is exited.
func (s *BaseMySqlParserListener) ExitAltblDropPK(ctx *AltblDropPKContext) {}

// EnterAltblDropIndex is called when production altblDropIndex is entered.
func (s *BaseMySqlParserListener) EnterAltblDropIndex(ctx *AltblDropIndexContext) {}

// ExitAltblDropIndex is called when production altblDropIndex is exited.
func (s *BaseMySqlParserListener) ExitAltblDropIndex(ctx *AltblDropIndexContext) {}

// EnterAltblDropFK is called when production altblDropFK is entered.
func (s *BaseMySqlParserListener) EnterAltblDropFK(ctx *AltblDropFKContext) {}

// ExitAltblDropFK is called when production altblDropFK is exited.
func (s *BaseMySqlParserListener) ExitAltblDropFK(ctx *AltblDropFKContext) {}

// EnterAltblDisKey is called when production altblDisKey is entered.
func (s *BaseMySqlParserListener) EnterAltblDisKey(ctx *AltblDisKeyContext) {}

// ExitAltblDisKey is called when production altblDisKey is exited.
func (s *BaseMySqlParserListener) ExitAltblDisKey(ctx *AltblDisKeyContext) {}

// EnterAltblEnKey is called when production altblEnKey is entered.
func (s *BaseMySqlParserListener) EnterAltblEnKey(ctx *AltblEnKeyContext) {}

// ExitAltblEnKey is called when production altblEnKey is exited.
func (s *BaseMySqlParserListener) ExitAltblEnKey(ctx *AltblEnKeyContext) {}

// EnterAltblRenameTbl is called when production altblRenameTbl is entered.
func (s *BaseMySqlParserListener) EnterAltblRenameTbl(ctx *AltblRenameTblContext) {}

// ExitAltblRenameTbl is called when production altblRenameTbl is exited.
func (s *BaseMySqlParserListener) ExitAltblRenameTbl(ctx *AltblRenameTblContext) {}

// EnterAltblResort is called when production altblResort is entered.
func (s *BaseMySqlParserListener) EnterAltblResort(ctx *AltblResortContext) {}

// ExitAltblResort is called when production altblResort is exited.
func (s *BaseMySqlParserListener) ExitAltblResort(ctx *AltblResortContext) {}

// EnterAltblConvert is called when production altblConvert is entered.
func (s *BaseMySqlParserListener) EnterAltblConvert(ctx *AltblConvertContext) {}

// ExitAltblConvert is called when production altblConvert is exited.
func (s *BaseMySqlParserListener) ExitAltblConvert(ctx *AltblConvertContext) {}

// EnterAltblDefCharset is called when production altblDefCharset is entered.
func (s *BaseMySqlParserListener) EnterAltblDefCharset(ctx *AltblDefCharsetContext) {}

// ExitAltblDefCharset is called when production altblDefCharset is exited.
func (s *BaseMySqlParserListener) ExitAltblDefCharset(ctx *AltblDefCharsetContext) {}

// EnterAltblDisTblspace is called when production altblDisTblspace is entered.
func (s *BaseMySqlParserListener) EnterAltblDisTblspace(ctx *AltblDisTblspaceContext) {}

// ExitAltblDisTblspace is called when production altblDisTblspace is exited.
func (s *BaseMySqlParserListener) ExitAltblDisTblspace(ctx *AltblDisTblspaceContext) {}

// EnterAltblImpTblSpace is called when production altblImpTblSpace is entered.
func (s *BaseMySqlParserListener) EnterAltblImpTblSpace(ctx *AltblImpTblSpaceContext) {}

// ExitAltblImpTblSpace is called when production altblImpTblSpace is exited.
func (s *BaseMySqlParserListener) ExitAltblImpTblSpace(ctx *AltblImpTblSpaceContext) {}

// EnterAltblForce is called when production altblForce is entered.
func (s *BaseMySqlParserListener) EnterAltblForce(ctx *AltblForceContext) {}

// ExitAltblForce is called when production altblForce is exited.
func (s *BaseMySqlParserListener) ExitAltblForce(ctx *AltblForceContext) {}

// EnterAltblValid is called when production altblValid is entered.
func (s *BaseMySqlParserListener) EnterAltblValid(ctx *AltblValidContext) {}

// ExitAltblValid is called when production altblValid is exited.
func (s *BaseMySqlParserListener) ExitAltblValid(ctx *AltblValidContext) {}

// EnterAltblAddPart is called when production altblAddPart is entered.
func (s *BaseMySqlParserListener) EnterAltblAddPart(ctx *AltblAddPartContext) {}

// ExitAltblAddPart is called when production altblAddPart is exited.
func (s *BaseMySqlParserListener) ExitAltblAddPart(ctx *AltblAddPartContext) {}

// EnterAltblDropPart is called when production altblDropPart is entered.
func (s *BaseMySqlParserListener) EnterAltblDropPart(ctx *AltblDropPartContext) {}

// ExitAltblDropPart is called when production altblDropPart is exited.
func (s *BaseMySqlParserListener) ExitAltblDropPart(ctx *AltblDropPartContext) {}

// EnterAltblDiscartPart is called when production altblDiscartPart is entered.
func (s *BaseMySqlParserListener) EnterAltblDiscartPart(ctx *AltblDiscartPartContext) {}

// ExitAltblDiscartPart is called when production altblDiscartPart is exited.
func (s *BaseMySqlParserListener) ExitAltblDiscartPart(ctx *AltblDiscartPartContext) {}

// EnterAltblImportPart is called when production altblImportPart is entered.
func (s *BaseMySqlParserListener) EnterAltblImportPart(ctx *AltblImportPartContext) {}

// ExitAltblImportPart is called when production altblImportPart is exited.
func (s *BaseMySqlParserListener) ExitAltblImportPart(ctx *AltblImportPartContext) {}

// EnterAltblTruncPart is called when production altblTruncPart is entered.
func (s *BaseMySqlParserListener) EnterAltblTruncPart(ctx *AltblTruncPartContext) {}

// ExitAltblTruncPart is called when production altblTruncPart is exited.
func (s *BaseMySqlParserListener) ExitAltblTruncPart(ctx *AltblTruncPartContext) {}

// EnterAltblCoalPart is called when production altblCoalPart is entered.
func (s *BaseMySqlParserListener) EnterAltblCoalPart(ctx *AltblCoalPartContext) {}

// ExitAltblCoalPart is called when production altblCoalPart is exited.
func (s *BaseMySqlParserListener) ExitAltblCoalPart(ctx *AltblCoalPartContext) {}

// EnterAltblReorgPart is called when production altblReorgPart is entered.
func (s *BaseMySqlParserListener) EnterAltblReorgPart(ctx *AltblReorgPartContext) {}

// ExitAltblReorgPart is called when production altblReorgPart is exited.
func (s *BaseMySqlParserListener) ExitAltblReorgPart(ctx *AltblReorgPartContext) {}

// EnterAltblExchPart is called when production altblExchPart is entered.
func (s *BaseMySqlParserListener) EnterAltblExchPart(ctx *AltblExchPartContext) {}

// ExitAltblExchPart is called when production altblExchPart is exited.
func (s *BaseMySqlParserListener) ExitAltblExchPart(ctx *AltblExchPartContext) {}

// EnterAltblAnalPart is called when production altblAnalPart is entered.
func (s *BaseMySqlParserListener) EnterAltblAnalPart(ctx *AltblAnalPartContext) {}

// ExitAltblAnalPart is called when production altblAnalPart is exited.
func (s *BaseMySqlParserListener) ExitAltblAnalPart(ctx *AltblAnalPartContext) {}

// EnterAltblCheckPart is called when production altblCheckPart is entered.
func (s *BaseMySqlParserListener) EnterAltblCheckPart(ctx *AltblCheckPartContext) {}

// ExitAltblCheckPart is called when production altblCheckPart is exited.
func (s *BaseMySqlParserListener) ExitAltblCheckPart(ctx *AltblCheckPartContext) {}

// EnterAltblOptimPart is called when production altblOptimPart is entered.
func (s *BaseMySqlParserListener) EnterAltblOptimPart(ctx *AltblOptimPartContext) {}

// ExitAltblOptimPart is called when production altblOptimPart is exited.
func (s *BaseMySqlParserListener) ExitAltblOptimPart(ctx *AltblOptimPartContext) {}

// EnterAltblRebuildPart is called when production altblRebuildPart is entered.
func (s *BaseMySqlParserListener) EnterAltblRebuildPart(ctx *AltblRebuildPartContext) {}

// ExitAltblRebuildPart is called when production altblRebuildPart is exited.
func (s *BaseMySqlParserListener) ExitAltblRebuildPart(ctx *AltblRebuildPartContext) {}

// EnterAltblRepairPart is called when production altblRepairPart is entered.
func (s *BaseMySqlParserListener) EnterAltblRepairPart(ctx *AltblRepairPartContext) {}

// ExitAltblRepairPart is called when production altblRepairPart is exited.
func (s *BaseMySqlParserListener) ExitAltblRepairPart(ctx *AltblRepairPartContext) {}

// EnterAltblRemovePart is called when production altblRemovePart is entered.
func (s *BaseMySqlParserListener) EnterAltblRemovePart(ctx *AltblRemovePartContext) {}

// ExitAltblRemovePart is called when production altblRemovePart is exited.
func (s *BaseMySqlParserListener) ExitAltblRemovePart(ctx *AltblRemovePartContext) {}

// EnterAltblUpgrPart is called when production altblUpgrPart is entered.
func (s *BaseMySqlParserListener) EnterAltblUpgrPart(ctx *AltblUpgrPartContext) {}

// ExitAltblUpgrPart is called when production altblUpgrPart is exited.
func (s *BaseMySqlParserListener) ExitAltblUpgrPart(ctx *AltblUpgrPartContext) {}

// EnterDrop_database is called when production drop_database is entered.
func (s *BaseMySqlParserListener) EnterDrop_database(ctx *Drop_databaseContext) {}

// ExitDrop_database is called when production drop_database is exited.
func (s *BaseMySqlParserListener) ExitDrop_database(ctx *Drop_databaseContext) {}

// EnterDrop_event is called when production drop_event is entered.
func (s *BaseMySqlParserListener) EnterDrop_event(ctx *Drop_eventContext) {}

// ExitDrop_event is called when production drop_event is exited.
func (s *BaseMySqlParserListener) ExitDrop_event(ctx *Drop_eventContext) {}

// EnterDrop_index is called when production drop_index is entered.
func (s *BaseMySqlParserListener) EnterDrop_index(ctx *Drop_indexContext) {}

// ExitDrop_index is called when production drop_index is exited.
func (s *BaseMySqlParserListener) ExitDrop_index(ctx *Drop_indexContext) {}

// EnterDrop_logfile_group is called when production drop_logfile_group is entered.
func (s *BaseMySqlParserListener) EnterDrop_logfile_group(ctx *Drop_logfile_groupContext) {}

// ExitDrop_logfile_group is called when production drop_logfile_group is exited.
func (s *BaseMySqlParserListener) ExitDrop_logfile_group(ctx *Drop_logfile_groupContext) {}

// EnterDrop_procedure is called when production drop_procedure is entered.
func (s *BaseMySqlParserListener) EnterDrop_procedure(ctx *Drop_procedureContext) {}

// ExitDrop_procedure is called when production drop_procedure is exited.
func (s *BaseMySqlParserListener) ExitDrop_procedure(ctx *Drop_procedureContext) {}

// EnterDrop_function is called when production drop_function is entered.
func (s *BaseMySqlParserListener) EnterDrop_function(ctx *Drop_functionContext) {}

// ExitDrop_function is called when production drop_function is exited.
func (s *BaseMySqlParserListener) ExitDrop_function(ctx *Drop_functionContext) {}

// EnterDrop_server is called when production drop_server is entered.
func (s *BaseMySqlParserListener) EnterDrop_server(ctx *Drop_serverContext) {}

// ExitDrop_server is called when production drop_server is exited.
func (s *BaseMySqlParserListener) ExitDrop_server(ctx *Drop_serverContext) {}

// EnterDrop_table is called when production drop_table is entered.
func (s *BaseMySqlParserListener) EnterDrop_table(ctx *Drop_tableContext) {}

// ExitDrop_table is called when production drop_table is exited.
func (s *BaseMySqlParserListener) ExitDrop_table(ctx *Drop_tableContext) {}

// EnterDrop_tablespace is called when production drop_tablespace is entered.
func (s *BaseMySqlParserListener) EnterDrop_tablespace(ctx *Drop_tablespaceContext) {}

// ExitDrop_tablespace is called when production drop_tablespace is exited.
func (s *BaseMySqlParserListener) ExitDrop_tablespace(ctx *Drop_tablespaceContext) {}

// EnterDrop_trigger is called when production drop_trigger is entered.
func (s *BaseMySqlParserListener) EnterDrop_trigger(ctx *Drop_triggerContext) {}

// ExitDrop_trigger is called when production drop_trigger is exited.
func (s *BaseMySqlParserListener) ExitDrop_trigger(ctx *Drop_triggerContext) {}

// EnterDrop_view is called when production drop_view is entered.
func (s *BaseMySqlParserListener) EnterDrop_view(ctx *Drop_viewContext) {}

// ExitDrop_view is called when production drop_view is exited.
func (s *BaseMySqlParserListener) ExitDrop_view(ctx *Drop_viewContext) {}

// EnterRename_table is called when production rename_table is entered.
func (s *BaseMySqlParserListener) EnterRename_table(ctx *Rename_tableContext) {}

// ExitRename_table is called when production rename_table is exited.
func (s *BaseMySqlParserListener) ExitRename_table(ctx *Rename_tableContext) {}

// EnterTruncate_table is called when production truncate_table is entered.
func (s *BaseMySqlParserListener) EnterTruncate_table(ctx *Truncate_tableContext) {}

// ExitTruncate_table is called when production truncate_table is exited.
func (s *BaseMySqlParserListener) ExitTruncate_table(ctx *Truncate_tableContext) {}

// EnterCall_statement is called when production call_statement is entered.
func (s *BaseMySqlParserListener) EnterCall_statement(ctx *Call_statementContext) {}

// ExitCall_statement is called when production call_statement is exited.
func (s *BaseMySqlParserListener) ExitCall_statement(ctx *Call_statementContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BaseMySqlParserListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BaseMySqlParserListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterDo_statement is called when production do_statement is entered.
func (s *BaseMySqlParserListener) EnterDo_statement(ctx *Do_statementContext) {}

// ExitDo_statement is called when production do_statement is exited.
func (s *BaseMySqlParserListener) ExitDo_statement(ctx *Do_statementContext) {}

// EnterHandler_statement is called when production handler_statement is entered.
func (s *BaseMySqlParserListener) EnterHandler_statement(ctx *Handler_statementContext) {}

// ExitHandler_statement is called when production handler_statement is exited.
func (s *BaseMySqlParserListener) ExitHandler_statement(ctx *Handler_statementContext) {}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *BaseMySqlParserListener) EnterInsert_statement(ctx *Insert_statementContext) {}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *BaseMySqlParserListener) ExitInsert_statement(ctx *Insert_statementContext) {}

// EnterLoad_data_statement is called when production load_data_statement is entered.
func (s *BaseMySqlParserListener) EnterLoad_data_statement(ctx *Load_data_statementContext) {}

// ExitLoad_data_statement is called when production load_data_statement is exited.
func (s *BaseMySqlParserListener) ExitLoad_data_statement(ctx *Load_data_statementContext) {}

// EnterLoad_xml_statement is called when production load_xml_statement is entered.
func (s *BaseMySqlParserListener) EnterLoad_xml_statement(ctx *Load_xml_statementContext) {}

// ExitLoad_xml_statement is called when production load_xml_statement is exited.
func (s *BaseMySqlParserListener) ExitLoad_xml_statement(ctx *Load_xml_statementContext) {}

// EnterReplace_statement is called when production replace_statement is entered.
func (s *BaseMySqlParserListener) EnterReplace_statement(ctx *Replace_statementContext) {}

// ExitReplace_statement is called when production replace_statement is exited.
func (s *BaseMySqlParserListener) ExitReplace_statement(ctx *Replace_statementContext) {}

// EnterSimpleSelect is called when production simpleSelect is entered.
func (s *BaseMySqlParserListener) EnterSimpleSelect(ctx *SimpleSelectContext) {}

// ExitSimpleSelect is called when production simpleSelect is exited.
func (s *BaseMySqlParserListener) ExitSimpleSelect(ctx *SimpleSelectContext) {}

// EnterParenSelect is called when production parenSelect is entered.
func (s *BaseMySqlParserListener) EnterParenSelect(ctx *ParenSelectContext) {}

// ExitParenSelect is called when production parenSelect is exited.
func (s *BaseMySqlParserListener) ExitParenSelect(ctx *ParenSelectContext) {}

// EnterUnionSelect is called when production unionSelect is entered.
func (s *BaseMySqlParserListener) EnterUnionSelect(ctx *UnionSelectContext) {}

// ExitUnionSelect is called when production unionSelect is exited.
func (s *BaseMySqlParserListener) ExitUnionSelect(ctx *UnionSelectContext) {}

// EnterUnionParenSelect is called when production unionParenSelect is entered.
func (s *BaseMySqlParserListener) EnterUnionParenSelect(ctx *UnionParenSelectContext) {}

// ExitUnionParenSelect is called when production unionParenSelect is exited.
func (s *BaseMySqlParserListener) ExitUnionParenSelect(ctx *UnionParenSelectContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BaseMySqlParserListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BaseMySqlParserListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterInsert_statement_value is called when production insert_statement_value is entered.
func (s *BaseMySqlParserListener) EnterInsert_statement_value(ctx *Insert_statement_valueContext) {}

// ExitInsert_statement_value is called when production insert_statement_value is exited.
func (s *BaseMySqlParserListener) ExitInsert_statement_value(ctx *Insert_statement_valueContext) {}

// EnterUpdate_elem is called when production update_elem is entered.
func (s *BaseMySqlParserListener) EnterUpdate_elem(ctx *Update_elemContext) {}

// ExitUpdate_elem is called when production update_elem is exited.
func (s *BaseMySqlParserListener) ExitUpdate_elem(ctx *Update_elemContext) {}

// EnterCol_or_uservar is called when production col_or_uservar is entered.
func (s *BaseMySqlParserListener) EnterCol_or_uservar(ctx *Col_or_uservarContext) {}

// ExitCol_or_uservar is called when production col_or_uservar is exited.
func (s *BaseMySqlParserListener) ExitCol_or_uservar(ctx *Col_or_uservarContext) {}

// EnterSingle_delete_statement is called when production single_delete_statement is entered.
func (s *BaseMySqlParserListener) EnterSingle_delete_statement(ctx *Single_delete_statementContext) {}

// ExitSingle_delete_statement is called when production single_delete_statement is exited.
func (s *BaseMySqlParserListener) ExitSingle_delete_statement(ctx *Single_delete_statementContext) {}

// EnterMultiple_delete_statement is called when production multiple_delete_statement is entered.
func (s *BaseMySqlParserListener) EnterMultiple_delete_statement(ctx *Multiple_delete_statementContext) {
}

// ExitMultiple_delete_statement is called when production multiple_delete_statement is exited.
func (s *BaseMySqlParserListener) ExitMultiple_delete_statement(ctx *Multiple_delete_statementContext) {
}

// EnterHandler_open_statement is called when production handler_open_statement is entered.
func (s *BaseMySqlParserListener) EnterHandler_open_statement(ctx *Handler_open_statementContext) {}

// ExitHandler_open_statement is called when production handler_open_statement is exited.
func (s *BaseMySqlParserListener) ExitHandler_open_statement(ctx *Handler_open_statementContext) {}

// EnterHandler_read_index_statement is called when production handler_read_index_statement is entered.
func (s *BaseMySqlParserListener) EnterHandler_read_index_statement(ctx *Handler_read_index_statementContext) {
}

// ExitHandler_read_index_statement is called when production handler_read_index_statement is exited.
func (s *BaseMySqlParserListener) ExitHandler_read_index_statement(ctx *Handler_read_index_statementContext) {
}

// EnterHandler_read_statement is called when production handler_read_statement is entered.
func (s *BaseMySqlParserListener) EnterHandler_read_statement(ctx *Handler_read_statementContext) {}

// ExitHandler_read_statement is called when production handler_read_statement is exited.
func (s *BaseMySqlParserListener) ExitHandler_read_statement(ctx *Handler_read_statementContext) {}

// EnterHandler_close_statement is called when production handler_close_statement is entered.
func (s *BaseMySqlParserListener) EnterHandler_close_statement(ctx *Handler_close_statementContext) {}

// ExitHandler_close_statement is called when production handler_close_statement is exited.
func (s *BaseMySqlParserListener) ExitHandler_close_statement(ctx *Handler_close_statementContext) {}

// EnterSingle_update_statement is called when production single_update_statement is entered.
func (s *BaseMySqlParserListener) EnterSingle_update_statement(ctx *Single_update_statementContext) {}

// ExitSingle_update_statement is called when production single_update_statement is exited.
func (s *BaseMySqlParserListener) ExitSingle_update_statement(ctx *Single_update_statementContext) {}

// EnterMultiple_update_statement is called when production multiple_update_statement is entered.
func (s *BaseMySqlParserListener) EnterMultiple_update_statement(ctx *Multiple_update_statementContext) {
}

// ExitMultiple_update_statement is called when production multiple_update_statement is exited.
func (s *BaseMySqlParserListener) ExitMultiple_update_statement(ctx *Multiple_update_statementContext) {
}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BaseMySqlParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BaseMySqlParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterOrder_by_expression is called when production order_by_expression is entered.
func (s *BaseMySqlParserListener) EnterOrder_by_expression(ctx *Order_by_expressionContext) {}

// ExitOrder_by_expression is called when production order_by_expression is exited.
func (s *BaseMySqlParserListener) ExitOrder_by_expression(ctx *Order_by_expressionContext) {}

// EnterTable_sources is called when production table_sources is entered.
func (s *BaseMySqlParserListener) EnterTable_sources(ctx *Table_sourcesContext) {}

// ExitTable_sources is called when production table_sources is exited.
func (s *BaseMySqlParserListener) ExitTable_sources(ctx *Table_sourcesContext) {}

// EnterTable_source is called when production table_source is entered.
func (s *BaseMySqlParserListener) EnterTable_source(ctx *Table_sourceContext) {}

// ExitTable_source is called when production table_source is exited.
func (s *BaseMySqlParserListener) ExitTable_source(ctx *Table_sourceContext) {}

// EnterAtomTableItem is called when production atomTableItem is entered.
func (s *BaseMySqlParserListener) EnterAtomTableItem(ctx *AtomTableItemContext) {}

// ExitAtomTableItem is called when production atomTableItem is exited.
func (s *BaseMySqlParserListener) ExitAtomTableItem(ctx *AtomTableItemContext) {}

// EnterSubqueryTableItem is called when production subqueryTableItem is entered.
func (s *BaseMySqlParserListener) EnterSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// ExitSubqueryTableItem is called when production subqueryTableItem is exited.
func (s *BaseMySqlParserListener) ExitSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// EnterTableSourcesItem is called when production tableSourcesItem is entered.
func (s *BaseMySqlParserListener) EnterTableSourcesItem(ctx *TableSourcesItemContext) {}

// ExitTableSourcesItem is called when production tableSourcesItem is exited.
func (s *BaseMySqlParserListener) ExitTableSourcesItem(ctx *TableSourcesItemContext) {}

// EnterIndex_hint is called when production index_hint is entered.
func (s *BaseMySqlParserListener) EnterIndex_hint(ctx *Index_hintContext) {}

// ExitIndex_hint is called when production index_hint is exited.
func (s *BaseMySqlParserListener) ExitIndex_hint(ctx *Index_hintContext) {}

// EnterInnerJoin is called when production innerJoin is entered.
func (s *BaseMySqlParserListener) EnterInnerJoin(ctx *InnerJoinContext) {}

// ExitInnerJoin is called when production innerJoin is exited.
func (s *BaseMySqlParserListener) ExitInnerJoin(ctx *InnerJoinContext) {}

// EnterStraightJoin is called when production straightJoin is entered.
func (s *BaseMySqlParserListener) EnterStraightJoin(ctx *StraightJoinContext) {}

// ExitStraightJoin is called when production straightJoin is exited.
func (s *BaseMySqlParserListener) ExitStraightJoin(ctx *StraightJoinContext) {}

// EnterOuterJoin is called when production outerJoin is entered.
func (s *BaseMySqlParserListener) EnterOuterJoin(ctx *OuterJoinContext) {}

// ExitOuterJoin is called when production outerJoin is exited.
func (s *BaseMySqlParserListener) ExitOuterJoin(ctx *OuterJoinContext) {}

// EnterNaturalJoin is called when production naturalJoin is entered.
func (s *BaseMySqlParserListener) EnterNaturalJoin(ctx *NaturalJoinContext) {}

// ExitNaturalJoin is called when production naturalJoin is exited.
func (s *BaseMySqlParserListener) ExitNaturalJoin(ctx *NaturalJoinContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseMySqlParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseMySqlParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterQuery_expression is called when production query_expression is entered.
func (s *BaseMySqlParserListener) EnterQuery_expression(ctx *Query_expressionContext) {}

// ExitQuery_expression is called when production query_expression is exited.
func (s *BaseMySqlParserListener) ExitQuery_expression(ctx *Query_expressionContext) {}

// EnterQuery_expression_nointo is called when production query_expression_nointo is entered.
func (s *BaseMySqlParserListener) EnterQuery_expression_nointo(ctx *Query_expression_nointoContext) {}

// ExitQuery_expression_nointo is called when production query_expression_nointo is exited.
func (s *BaseMySqlParserListener) ExitQuery_expression_nointo(ctx *Query_expression_nointoContext) {}

// EnterQuery_specification is called when production query_specification is entered.
func (s *BaseMySqlParserListener) EnterQuery_specification(ctx *Query_specificationContext) {}

// ExitQuery_specification is called when production query_specification is exited.
func (s *BaseMySqlParserListener) ExitQuery_specification(ctx *Query_specificationContext) {}

// EnterQuery_specification_nointo is called when production query_specification_nointo is entered.
func (s *BaseMySqlParserListener) EnterQuery_specification_nointo(ctx *Query_specification_nointoContext) {
}

// ExitQuery_specification_nointo is called when production query_specification_nointo is exited.
func (s *BaseMySqlParserListener) ExitQuery_specification_nointo(ctx *Query_specification_nointoContext) {
}

// EnterUnion_parenth is called when production union_parenth is entered.
func (s *BaseMySqlParserListener) EnterUnion_parenth(ctx *Union_parenthContext) {}

// ExitUnion_parenth is called when production union_parenth is exited.
func (s *BaseMySqlParserListener) ExitUnion_parenth(ctx *Union_parenthContext) {}

// EnterUnion_statement is called when production union_statement is entered.
func (s *BaseMySqlParserListener) EnterUnion_statement(ctx *Union_statementContext) {}

// ExitUnion_statement is called when production union_statement is exited.
func (s *BaseMySqlParserListener) ExitUnion_statement(ctx *Union_statementContext) {}

// EnterSelect_spec is called when production select_spec is entered.
func (s *BaseMySqlParserListener) EnterSelect_spec(ctx *Select_specContext) {}

// ExitSelect_spec is called when production select_spec is exited.
func (s *BaseMySqlParserListener) ExitSelect_spec(ctx *Select_specContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseMySqlParserListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseMySqlParserListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterSellistelAllCol is called when production sellistelAllCol is entered.
func (s *BaseMySqlParserListener) EnterSellistelAllCol(ctx *SellistelAllColContext) {}

// ExitSellistelAllCol is called when production sellistelAllCol is exited.
func (s *BaseMySqlParserListener) ExitSellistelAllCol(ctx *SellistelAllColContext) {}

// EnterSellistelCol is called when production sellistelCol is entered.
func (s *BaseMySqlParserListener) EnterSellistelCol(ctx *SellistelColContext) {}

// ExitSellistelCol is called when production sellistelCol is exited.
func (s *BaseMySqlParserListener) ExitSellistelCol(ctx *SellistelColContext) {}

// EnterSellistelFunc is called when production sellistelFunc is entered.
func (s *BaseMySqlParserListener) EnterSellistelFunc(ctx *SellistelFuncContext) {}

// ExitSellistelFunc is called when production sellistelFunc is exited.
func (s *BaseMySqlParserListener) ExitSellistelFunc(ctx *SellistelFuncContext) {}

// EnterSellistelExpr is called when production sellistelExpr is entered.
func (s *BaseMySqlParserListener) EnterSellistelExpr(ctx *SellistelExprContext) {}

// ExitSellistelExpr is called when production sellistelExpr is exited.
func (s *BaseMySqlParserListener) ExitSellistelExpr(ctx *SellistelExprContext) {}

// EnterSelectIntoVars is called when production selectIntoVars is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoVars(ctx *SelectIntoVarsContext) {}

// ExitSelectIntoVars is called when production selectIntoVars is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoVars(ctx *SelectIntoVarsContext) {}

// EnterSelectIntoDump is called when production selectIntoDump is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoDump(ctx *SelectIntoDumpContext) {}

// ExitSelectIntoDump is called when production selectIntoDump is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoDump(ctx *SelectIntoDumpContext) {}

// EnterSelectIntoOutfile is called when production selectIntoOutfile is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoOutfile(ctx *SelectIntoOutfileContext) {}

// ExitSelectIntoOutfile is called when production selectIntoOutfile is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoOutfile(ctx *SelectIntoOutfileContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseMySqlParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseMySqlParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *BaseMySqlParserListener) EnterGroup_by_item(ctx *Group_by_itemContext) {}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *BaseMySqlParserListener) ExitGroup_by_item(ctx *Group_by_itemContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseMySqlParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseMySqlParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterStart_transaction is called when production start_transaction is entered.
func (s *BaseMySqlParserListener) EnterStart_transaction(ctx *Start_transactionContext) {}

// ExitStart_transaction is called when production start_transaction is exited.
func (s *BaseMySqlParserListener) ExitStart_transaction(ctx *Start_transactionContext) {}

// EnterBegin_work is called when production begin_work is entered.
func (s *BaseMySqlParserListener) EnterBegin_work(ctx *Begin_workContext) {}

// ExitBegin_work is called when production begin_work is exited.
func (s *BaseMySqlParserListener) ExitBegin_work(ctx *Begin_workContext) {}

// EnterCommit_work is called when production commit_work is entered.
func (s *BaseMySqlParserListener) EnterCommit_work(ctx *Commit_workContext) {}

// ExitCommit_work is called when production commit_work is exited.
func (s *BaseMySqlParserListener) ExitCommit_work(ctx *Commit_workContext) {}

// EnterRollback_work is called when production rollback_work is entered.
func (s *BaseMySqlParserListener) EnterRollback_work(ctx *Rollback_workContext) {}

// ExitRollback_work is called when production rollback_work is exited.
func (s *BaseMySqlParserListener) ExitRollback_work(ctx *Rollback_workContext) {}

// EnterSavepoint_statement is called when production savepoint_statement is entered.
func (s *BaseMySqlParserListener) EnterSavepoint_statement(ctx *Savepoint_statementContext) {}

// ExitSavepoint_statement is called when production savepoint_statement is exited.
func (s *BaseMySqlParserListener) ExitSavepoint_statement(ctx *Savepoint_statementContext) {}

// EnterRollback_statement is called when production rollback_statement is entered.
func (s *BaseMySqlParserListener) EnterRollback_statement(ctx *Rollback_statementContext) {}

// ExitRollback_statement is called when production rollback_statement is exited.
func (s *BaseMySqlParserListener) ExitRollback_statement(ctx *Rollback_statementContext) {}

// EnterRelease_statement is called when production release_statement is entered.
func (s *BaseMySqlParserListener) EnterRelease_statement(ctx *Release_statementContext) {}

// ExitRelease_statement is called when production release_statement is exited.
func (s *BaseMySqlParserListener) ExitRelease_statement(ctx *Release_statementContext) {}

// EnterLock_tables is called when production lock_tables is entered.
func (s *BaseMySqlParserListener) EnterLock_tables(ctx *Lock_tablesContext) {}

// ExitLock_tables is called when production lock_tables is exited.
func (s *BaseMySqlParserListener) ExitLock_tables(ctx *Lock_tablesContext) {}

// EnterUnlock_tables is called when production unlock_tables is entered.
func (s *BaseMySqlParserListener) EnterUnlock_tables(ctx *Unlock_tablesContext) {}

// ExitUnlock_tables is called when production unlock_tables is exited.
func (s *BaseMySqlParserListener) ExitUnlock_tables(ctx *Unlock_tablesContext) {}

// EnterSet_autocommit_statement is called when production set_autocommit_statement is entered.
func (s *BaseMySqlParserListener) EnterSet_autocommit_statement(ctx *Set_autocommit_statementContext) {
}

// ExitSet_autocommit_statement is called when production set_autocommit_statement is exited.
func (s *BaseMySqlParserListener) ExitSet_autocommit_statement(ctx *Set_autocommit_statementContext) {}

// EnterSet_transaction_statement is called when production set_transaction_statement is entered.
func (s *BaseMySqlParserListener) EnterSet_transaction_statement(ctx *Set_transaction_statementContext) {
}

// ExitSet_transaction_statement is called when production set_transaction_statement is exited.
func (s *BaseMySqlParserListener) ExitSet_transaction_statement(ctx *Set_transaction_statementContext) {
}

// EnterTransact_option is called when production transact_option is entered.
func (s *BaseMySqlParserListener) EnterTransact_option(ctx *Transact_optionContext) {}

// ExitTransact_option is called when production transact_option is exited.
func (s *BaseMySqlParserListener) ExitTransact_option(ctx *Transact_optionContext) {}

// EnterLock_table_element is called when production lock_table_element is entered.
func (s *BaseMySqlParserListener) EnterLock_table_element(ctx *Lock_table_elementContext) {}

// ExitLock_table_element is called when production lock_table_element is exited.
func (s *BaseMySqlParserListener) ExitLock_table_element(ctx *Lock_table_elementContext) {}

// EnterTrans_characteristic is called when production trans_characteristic is entered.
func (s *BaseMySqlParserListener) EnterTrans_characteristic(ctx *Trans_characteristicContext) {}

// ExitTrans_characteristic is called when production trans_characteristic is exited.
func (s *BaseMySqlParserListener) ExitTrans_characteristic(ctx *Trans_characteristicContext) {}

// EnterTransaction_level is called when production transaction_level is entered.
func (s *BaseMySqlParserListener) EnterTransaction_level(ctx *Transaction_levelContext) {}

// ExitTransaction_level is called when production transaction_level is exited.
func (s *BaseMySqlParserListener) ExitTransaction_level(ctx *Transaction_levelContext) {}

// EnterChange_master is called when production change_master is entered.
func (s *BaseMySqlParserListener) EnterChange_master(ctx *Change_masterContext) {}

// ExitChange_master is called when production change_master is exited.
func (s *BaseMySqlParserListener) ExitChange_master(ctx *Change_masterContext) {}

// EnterChange_repl_filter is called when production change_repl_filter is entered.
func (s *BaseMySqlParserListener) EnterChange_repl_filter(ctx *Change_repl_filterContext) {}

// ExitChange_repl_filter is called when production change_repl_filter is exited.
func (s *BaseMySqlParserListener) ExitChange_repl_filter(ctx *Change_repl_filterContext) {}

// EnterPurge_binary_logs is called when production purge_binary_logs is entered.
func (s *BaseMySqlParserListener) EnterPurge_binary_logs(ctx *Purge_binary_logsContext) {}

// ExitPurge_binary_logs is called when production purge_binary_logs is exited.
func (s *BaseMySqlParserListener) ExitPurge_binary_logs(ctx *Purge_binary_logsContext) {}

// EnterReset_master is called when production reset_master is entered.
func (s *BaseMySqlParserListener) EnterReset_master(ctx *Reset_masterContext) {}

// ExitReset_master is called when production reset_master is exited.
func (s *BaseMySqlParserListener) ExitReset_master(ctx *Reset_masterContext) {}

// EnterReset_slave is called when production reset_slave is entered.
func (s *BaseMySqlParserListener) EnterReset_slave(ctx *Reset_slaveContext) {}

// ExitReset_slave is called when production reset_slave is exited.
func (s *BaseMySqlParserListener) ExitReset_slave(ctx *Reset_slaveContext) {}

// EnterStart_slave is called when production start_slave is entered.
func (s *BaseMySqlParserListener) EnterStart_slave(ctx *Start_slaveContext) {}

// ExitStart_slave is called when production start_slave is exited.
func (s *BaseMySqlParserListener) ExitStart_slave(ctx *Start_slaveContext) {}

// EnterStop_slave is called when production stop_slave is entered.
func (s *BaseMySqlParserListener) EnterStop_slave(ctx *Stop_slaveContext) {}

// ExitStop_slave is called when production stop_slave is exited.
func (s *BaseMySqlParserListener) ExitStop_slave(ctx *Stop_slaveContext) {}

// EnterStart_group_repl is called when production start_group_repl is entered.
func (s *BaseMySqlParserListener) EnterStart_group_repl(ctx *Start_group_replContext) {}

// ExitStart_group_repl is called when production start_group_repl is exited.
func (s *BaseMySqlParserListener) ExitStart_group_repl(ctx *Start_group_replContext) {}

// EnterStop_group_repl is called when production stop_group_repl is entered.
func (s *BaseMySqlParserListener) EnterStop_group_repl(ctx *Stop_group_replContext) {}

// ExitStop_group_repl is called when production stop_group_repl is exited.
func (s *BaseMySqlParserListener) ExitStop_group_repl(ctx *Stop_group_replContext) {}

// EnterMasterOptString is called when production masterOptString is entered.
func (s *BaseMySqlParserListener) EnterMasterOptString(ctx *MasterOptStringContext) {}

// ExitMasterOptString is called when production masterOptString is exited.
func (s *BaseMySqlParserListener) ExitMasterOptString(ctx *MasterOptStringContext) {}

// EnterMasterOptDecimal is called when production masterOptDecimal is entered.
func (s *BaseMySqlParserListener) EnterMasterOptDecimal(ctx *MasterOptDecimalContext) {}

// ExitMasterOptDecimal is called when production masterOptDecimal is exited.
func (s *BaseMySqlParserListener) ExitMasterOptDecimal(ctx *MasterOptDecimalContext) {}

// EnterMasterOptBool is called when production masterOptBool is entered.
func (s *BaseMySqlParserListener) EnterMasterOptBool(ctx *MasterOptBoolContext) {}

// ExitMasterOptBool is called when production masterOptBool is exited.
func (s *BaseMySqlParserListener) ExitMasterOptBool(ctx *MasterOptBoolContext) {}

// EnterMasterOptReal is called when production masterOptReal is entered.
func (s *BaseMySqlParserListener) EnterMasterOptReal(ctx *MasterOptRealContext) {}

// ExitMasterOptReal is called when production masterOptReal is exited.
func (s *BaseMySqlParserListener) ExitMasterOptReal(ctx *MasterOptRealContext) {}

// EnterMasterOptIdList is called when production masterOptIdList is entered.
func (s *BaseMySqlParserListener) EnterMasterOptIdList(ctx *MasterOptIdListContext) {}

// ExitMasterOptIdList is called when production masterOptIdList is exited.
func (s *BaseMySqlParserListener) ExitMasterOptIdList(ctx *MasterOptIdListContext) {}

// EnterString_master_option is called when production string_master_option is entered.
func (s *BaseMySqlParserListener) EnterString_master_option(ctx *String_master_optionContext) {}

// ExitString_master_option is called when production string_master_option is exited.
func (s *BaseMySqlParserListener) ExitString_master_option(ctx *String_master_optionContext) {}

// EnterDecimal_master_option is called when production decimal_master_option is entered.
func (s *BaseMySqlParserListener) EnterDecimal_master_option(ctx *Decimal_master_optionContext) {}

// ExitDecimal_master_option is called when production decimal_master_option is exited.
func (s *BaseMySqlParserListener) ExitDecimal_master_option(ctx *Decimal_master_optionContext) {}

// EnterBool_master_option is called when production bool_master_option is entered.
func (s *BaseMySqlParserListener) EnterBool_master_option(ctx *Bool_master_optionContext) {}

// ExitBool_master_option is called when production bool_master_option is exited.
func (s *BaseMySqlParserListener) ExitBool_master_option(ctx *Bool_master_optionContext) {}

// EnterChannel_option is called when production channel_option is entered.
func (s *BaseMySqlParserListener) EnterChannel_option(ctx *Channel_optionContext) {}

// ExitChannel_option is called when production channel_option is exited.
func (s *BaseMySqlParserListener) ExitChannel_option(ctx *Channel_optionContext) {}

// EnterReplfilterDbList is called when production replfilterDbList is entered.
func (s *BaseMySqlParserListener) EnterReplfilterDbList(ctx *ReplfilterDbListContext) {}

// ExitReplfilterDbList is called when production replfilterDbList is exited.
func (s *BaseMySqlParserListener) ExitReplfilterDbList(ctx *ReplfilterDbListContext) {}

// EnterReplfilterTableList is called when production replfilterTableList is entered.
func (s *BaseMySqlParserListener) EnterReplfilterTableList(ctx *ReplfilterTableListContext) {}

// ExitReplfilterTableList is called when production replfilterTableList is exited.
func (s *BaseMySqlParserListener) ExitReplfilterTableList(ctx *ReplfilterTableListContext) {}

// EnterReplfilterStableList is called when production replfilterStableList is entered.
func (s *BaseMySqlParserListener) EnterReplfilterStableList(ctx *ReplfilterStableListContext) {}

// ExitReplfilterStableList is called when production replfilterStableList is exited.
func (s *BaseMySqlParserListener) ExitReplfilterStableList(ctx *ReplfilterStableListContext) {}

// EnterReplfilterTablepairList is called when production replfilterTablepairList is entered.
func (s *BaseMySqlParserListener) EnterReplfilterTablepairList(ctx *ReplfilterTablepairListContext) {}

// ExitReplfilterTablepairList is called when production replfilterTablepairList is exited.
func (s *BaseMySqlParserListener) ExitReplfilterTablepairList(ctx *ReplfilterTablepairListContext) {}

// EnterThread_type is called when production thread_type is entered.
func (s *BaseMySqlParserListener) EnterThread_type(ctx *Thread_typeContext) {}

// ExitThread_type is called when production thread_type is exited.
func (s *BaseMySqlParserListener) ExitThread_type(ctx *Thread_typeContext) {}

// EnterUntilGtidSset is called when production untilGtidSset is entered.
func (s *BaseMySqlParserListener) EnterUntilGtidSset(ctx *UntilGtidSsetContext) {}

// ExitUntilGtidSset is called when production untilGtidSset is exited.
func (s *BaseMySqlParserListener) ExitUntilGtidSset(ctx *UntilGtidSsetContext) {}

// EnterUntilMasterLog is called when production untilMasterLog is entered.
func (s *BaseMySqlParserListener) EnterUntilMasterLog(ctx *UntilMasterLogContext) {}

// ExitUntilMasterLog is called when production untilMasterLog is exited.
func (s *BaseMySqlParserListener) ExitUntilMasterLog(ctx *UntilMasterLogContext) {}

// EnterUntilRelayLog is called when production untilRelayLog is entered.
func (s *BaseMySqlParserListener) EnterUntilRelayLog(ctx *UntilRelayLogContext) {}

// ExitUntilRelayLog is called when production untilRelayLog is exited.
func (s *BaseMySqlParserListener) ExitUntilRelayLog(ctx *UntilRelayLogContext) {}

// EnterUntilSqlGaps is called when production untilSqlGaps is entered.
func (s *BaseMySqlParserListener) EnterUntilSqlGaps(ctx *UntilSqlGapsContext) {}

// ExitUntilSqlGaps is called when production untilSqlGaps is exited.
func (s *BaseMySqlParserListener) ExitUntilSqlGaps(ctx *UntilSqlGapsContext) {}

// EnterStart_slave_connection_option is called when production start_slave_connection_option is entered.
func (s *BaseMySqlParserListener) EnterStart_slave_connection_option(ctx *Start_slave_connection_optionContext) {
}

// ExitStart_slave_connection_option is called when production start_slave_connection_option is exited.
func (s *BaseMySqlParserListener) ExitStart_slave_connection_option(ctx *Start_slave_connection_optionContext) {
}

// EnterGtid_set is called when production gtid_set is entered.
func (s *BaseMySqlParserListener) EnterGtid_set(ctx *Gtid_setContext) {}

// ExitGtid_set is called when production gtid_set is exited.
func (s *BaseMySqlParserListener) ExitGtid_set(ctx *Gtid_setContext) {}

// EnterXa_start_transaction is called when production xa_start_transaction is entered.
func (s *BaseMySqlParserListener) EnterXa_start_transaction(ctx *Xa_start_transactionContext) {}

// ExitXa_start_transaction is called when production xa_start_transaction is exited.
func (s *BaseMySqlParserListener) ExitXa_start_transaction(ctx *Xa_start_transactionContext) {}

// EnterXa_end_transaction is called when production xa_end_transaction is entered.
func (s *BaseMySqlParserListener) EnterXa_end_transaction(ctx *Xa_end_transactionContext) {}

// ExitXa_end_transaction is called when production xa_end_transaction is exited.
func (s *BaseMySqlParserListener) ExitXa_end_transaction(ctx *Xa_end_transactionContext) {}

// EnterXa_prepare is called when production xa_prepare is entered.
func (s *BaseMySqlParserListener) EnterXa_prepare(ctx *Xa_prepareContext) {}

// ExitXa_prepare is called when production xa_prepare is exited.
func (s *BaseMySqlParserListener) ExitXa_prepare(ctx *Xa_prepareContext) {}

// EnterXa_commit_work is called when production xa_commit_work is entered.
func (s *BaseMySqlParserListener) EnterXa_commit_work(ctx *Xa_commit_workContext) {}

// ExitXa_commit_work is called when production xa_commit_work is exited.
func (s *BaseMySqlParserListener) ExitXa_commit_work(ctx *Xa_commit_workContext) {}

// EnterXa_rollback_work is called when production xa_rollback_work is entered.
func (s *BaseMySqlParserListener) EnterXa_rollback_work(ctx *Xa_rollback_workContext) {}

// ExitXa_rollback_work is called when production xa_rollback_work is exited.
func (s *BaseMySqlParserListener) ExitXa_rollback_work(ctx *Xa_rollback_workContext) {}

// EnterXa_recover_work is called when production xa_recover_work is entered.
func (s *BaseMySqlParserListener) EnterXa_recover_work(ctx *Xa_recover_workContext) {}

// ExitXa_recover_work is called when production xa_recover_work is exited.
func (s *BaseMySqlParserListener) ExitXa_recover_work(ctx *Xa_recover_workContext) {}

// EnterPrepare_statement is called when production prepare_statement is entered.
func (s *BaseMySqlParserListener) EnterPrepare_statement(ctx *Prepare_statementContext) {}

// ExitPrepare_statement is called when production prepare_statement is exited.
func (s *BaseMySqlParserListener) ExitPrepare_statement(ctx *Prepare_statementContext) {}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *BaseMySqlParserListener) EnterExecute_statement(ctx *Execute_statementContext) {}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *BaseMySqlParserListener) ExitExecute_statement(ctx *Execute_statementContext) {}

// EnterDeallocate_prepare is called when production deallocate_prepare is entered.
func (s *BaseMySqlParserListener) EnterDeallocate_prepare(ctx *Deallocate_prepareContext) {}

// ExitDeallocate_prepare is called when production deallocate_prepare is exited.
func (s *BaseMySqlParserListener) ExitDeallocate_prepare(ctx *Deallocate_prepareContext) {}

// EnterRoutine_body is called when production routine_body is entered.
func (s *BaseMySqlParserListener) EnterRoutine_body(ctx *Routine_bodyContext) {}

// ExitRoutine_body is called when production routine_body is exited.
func (s *BaseMySqlParserListener) ExitRoutine_body(ctx *Routine_bodyContext) {}

// EnterBlock_statement is called when production block_statement is entered.
func (s *BaseMySqlParserListener) EnterBlock_statement(ctx *Block_statementContext) {}

// ExitBlock_statement is called when production block_statement is exited.
func (s *BaseMySqlParserListener) ExitBlock_statement(ctx *Block_statementContext) {}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseMySqlParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseMySqlParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseMySqlParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseMySqlParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterIterate_statement is called when production iterate_statement is entered.
func (s *BaseMySqlParserListener) EnterIterate_statement(ctx *Iterate_statementContext) {}

// ExitIterate_statement is called when production iterate_statement is exited.
func (s *BaseMySqlParserListener) ExitIterate_statement(ctx *Iterate_statementContext) {}

// EnterLeave_statement is called when production leave_statement is entered.
func (s *BaseMySqlParserListener) EnterLeave_statement(ctx *Leave_statementContext) {}

// ExitLeave_statement is called when production leave_statement is exited.
func (s *BaseMySqlParserListener) ExitLeave_statement(ctx *Leave_statementContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseMySqlParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseMySqlParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterRepeat_statement is called when production repeat_statement is entered.
func (s *BaseMySqlParserListener) EnterRepeat_statement(ctx *Repeat_statementContext) {}

// ExitRepeat_statement is called when production repeat_statement is exited.
func (s *BaseMySqlParserListener) ExitRepeat_statement(ctx *Repeat_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseMySqlParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseMySqlParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseMySqlParserListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseMySqlParserListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterCursor_statement is called when production cursor_statement is entered.
func (s *BaseMySqlParserListener) EnterCursor_statement(ctx *Cursor_statementContext) {}

// ExitCursor_statement is called when production cursor_statement is exited.
func (s *BaseMySqlParserListener) ExitCursor_statement(ctx *Cursor_statementContext) {}

// EnterDeclare_variable is called when production declare_variable is entered.
func (s *BaseMySqlParserListener) EnterDeclare_variable(ctx *Declare_variableContext) {}

// ExitDeclare_variable is called when production declare_variable is exited.
func (s *BaseMySqlParserListener) ExitDeclare_variable(ctx *Declare_variableContext) {}

// EnterDeclare_condition is called when production declare_condition is entered.
func (s *BaseMySqlParserListener) EnterDeclare_condition(ctx *Declare_conditionContext) {}

// ExitDeclare_condition is called when production declare_condition is exited.
func (s *BaseMySqlParserListener) ExitDeclare_condition(ctx *Declare_conditionContext) {}

// EnterDeclare_cursor is called when production declare_cursor is entered.
func (s *BaseMySqlParserListener) EnterDeclare_cursor(ctx *Declare_cursorContext) {}

// ExitDeclare_cursor is called when production declare_cursor is exited.
func (s *BaseMySqlParserListener) ExitDeclare_cursor(ctx *Declare_cursorContext) {}

// EnterDeclare_handler is called when production declare_handler is entered.
func (s *BaseMySqlParserListener) EnterDeclare_handler(ctx *Declare_handlerContext) {}

// ExitDeclare_handler is called when production declare_handler is exited.
func (s *BaseMySqlParserListener) ExitDeclare_handler(ctx *Declare_handlerContext) {}

// EnterHandler_condition_value is called when production handler_condition_value is entered.
func (s *BaseMySqlParserListener) EnterHandler_condition_value(ctx *Handler_condition_valueContext) {}

// ExitHandler_condition_value is called when production handler_condition_value is exited.
func (s *BaseMySqlParserListener) ExitHandler_condition_value(ctx *Handler_condition_valueContext) {}

// EnterProcedure_sql_statement is called when production procedure_sql_statement is entered.
func (s *BaseMySqlParserListener) EnterProcedure_sql_statement(ctx *Procedure_sql_statementContext) {}

// ExitProcedure_sql_statement is called when production procedure_sql_statement is exited.
func (s *BaseMySqlParserListener) ExitProcedure_sql_statement(ctx *Procedure_sql_statementContext) {}

// EnterAlterUserMysql56 is called when production alterUserMysql56 is entered.
func (s *BaseMySqlParserListener) EnterAlterUserMysql56(ctx *AlterUserMysql56Context) {}

// ExitAlterUserMysql56 is called when production alterUserMysql56 is exited.
func (s *BaseMySqlParserListener) ExitAlterUserMysql56(ctx *AlterUserMysql56Context) {}

// EnterAlterUserMysql57 is called when production alterUserMysql57 is entered.
func (s *BaseMySqlParserListener) EnterAlterUserMysql57(ctx *AlterUserMysql57Context) {}

// ExitAlterUserMysql57 is called when production alterUserMysql57 is exited.
func (s *BaseMySqlParserListener) ExitAlterUserMysql57(ctx *AlterUserMysql57Context) {}

// EnterCreateUserMysql56 is called when production createUserMysql56 is entered.
func (s *BaseMySqlParserListener) EnterCreateUserMysql56(ctx *CreateUserMysql56Context) {}

// ExitCreateUserMysql56 is called when production createUserMysql56 is exited.
func (s *BaseMySqlParserListener) ExitCreateUserMysql56(ctx *CreateUserMysql56Context) {}

// EnterCreateUserMysql57 is called when production createUserMysql57 is entered.
func (s *BaseMySqlParserListener) EnterCreateUserMysql57(ctx *CreateUserMysql57Context) {}

// ExitCreateUserMysql57 is called when production createUserMysql57 is exited.
func (s *BaseMySqlParserListener) ExitCreateUserMysql57(ctx *CreateUserMysql57Context) {}

// EnterDrop_user is called when production drop_user is entered.
func (s *BaseMySqlParserListener) EnterDrop_user(ctx *Drop_userContext) {}

// ExitDrop_user is called when production drop_user is exited.
func (s *BaseMySqlParserListener) ExitDrop_user(ctx *Drop_userContext) {}

// EnterGrant_statement is called when production grant_statement is entered.
func (s *BaseMySqlParserListener) EnterGrant_statement(ctx *Grant_statementContext) {}

// ExitGrant_statement is called when production grant_statement is exited.
func (s *BaseMySqlParserListener) ExitGrant_statement(ctx *Grant_statementContext) {}

// EnterGrant_proxy is called when production grant_proxy is entered.
func (s *BaseMySqlParserListener) EnterGrant_proxy(ctx *Grant_proxyContext) {}

// ExitGrant_proxy is called when production grant_proxy is exited.
func (s *BaseMySqlParserListener) ExitGrant_proxy(ctx *Grant_proxyContext) {}

// EnterRename_user is called when production rename_user is entered.
func (s *BaseMySqlParserListener) EnterRename_user(ctx *Rename_userContext) {}

// ExitRename_user is called when production rename_user is exited.
func (s *BaseMySqlParserListener) ExitRename_user(ctx *Rename_userContext) {}

// EnterDetailRevoke is called when production detailRevoke is entered.
func (s *BaseMySqlParserListener) EnterDetailRevoke(ctx *DetailRevokeContext) {}

// ExitDetailRevoke is called when production detailRevoke is exited.
func (s *BaseMySqlParserListener) ExitDetailRevoke(ctx *DetailRevokeContext) {}

// EnterShortRevoke is called when production shortRevoke is entered.
func (s *BaseMySqlParserListener) EnterShortRevoke(ctx *ShortRevokeContext) {}

// ExitShortRevoke is called when production shortRevoke is exited.
func (s *BaseMySqlParserListener) ExitShortRevoke(ctx *ShortRevokeContext) {}

// EnterRevoke_proxy is called when production revoke_proxy is entered.
func (s *BaseMySqlParserListener) EnterRevoke_proxy(ctx *Revoke_proxyContext) {}

// ExitRevoke_proxy is called when production revoke_proxy is exited.
func (s *BaseMySqlParserListener) ExitRevoke_proxy(ctx *Revoke_proxyContext) {}

// EnterSet_password_statement is called when production set_password_statement is entered.
func (s *BaseMySqlParserListener) EnterSet_password_statement(ctx *Set_password_statementContext) {}

// ExitSet_password_statement is called when production set_password_statement is exited.
func (s *BaseMySqlParserListener) ExitSet_password_statement(ctx *Set_password_statementContext) {}

// EnterUser_password_option is called when production user_password_option is entered.
func (s *BaseMySqlParserListener) EnterUser_password_option(ctx *User_password_optionContext) {}

// ExitUser_password_option is called when production user_password_option is exited.
func (s *BaseMySqlParserListener) ExitUser_password_option(ctx *User_password_optionContext) {}

// EnterAuthByPassword is called when production authByPassword is entered.
func (s *BaseMySqlParserListener) EnterAuthByPassword(ctx *AuthByPasswordContext) {}

// ExitAuthByPassword is called when production authByPassword is exited.
func (s *BaseMySqlParserListener) ExitAuthByPassword(ctx *AuthByPasswordContext) {}

// EnterAuthByString is called when production authByString is entered.
func (s *BaseMySqlParserListener) EnterAuthByString(ctx *AuthByStringContext) {}

// ExitAuthByString is called when production authByString is exited.
func (s *BaseMySqlParserListener) ExitAuthByString(ctx *AuthByStringContext) {}

// EnterAuthByHash is called when production authByHash is entered.
func (s *BaseMySqlParserListener) EnterAuthByHash(ctx *AuthByHashContext) {}

// ExitAuthByHash is called when production authByHash is exited.
func (s *BaseMySqlParserListener) ExitAuthByHash(ctx *AuthByHashContext) {}

// EnterTls_option is called when production tls_option is entered.
func (s *BaseMySqlParserListener) EnterTls_option(ctx *Tls_optionContext) {}

// ExitTls_option is called when production tls_option is exited.
func (s *BaseMySqlParserListener) ExitTls_option(ctx *Tls_optionContext) {}

// EnterUser_resource_option is called when production user_resource_option is entered.
func (s *BaseMySqlParserListener) EnterUser_resource_option(ctx *User_resource_optionContext) {}

// ExitUser_resource_option is called when production user_resource_option is exited.
func (s *BaseMySqlParserListener) ExitUser_resource_option(ctx *User_resource_optionContext) {}

// EnterUser_lock_option is called when production user_lock_option is entered.
func (s *BaseMySqlParserListener) EnterUser_lock_option(ctx *User_lock_optionContext) {}

// ExitUser_lock_option is called when production user_lock_option is exited.
func (s *BaseMySqlParserListener) ExitUser_lock_option(ctx *User_lock_optionContext) {}

// EnterPrivelege_clause is called when production privelege_clause is entered.
func (s *BaseMySqlParserListener) EnterPrivelege_clause(ctx *Privelege_clauseContext) {}

// ExitPrivelege_clause is called when production privelege_clause is exited.
func (s *BaseMySqlParserListener) ExitPrivelege_clause(ctx *Privelege_clauseContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseMySqlParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseMySqlParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterPrivilege_level is called when production privilege_level is entered.
func (s *BaseMySqlParserListener) EnterPrivilege_level(ctx *Privilege_levelContext) {}

// ExitPrivilege_level is called when production privilege_level is exited.
func (s *BaseMySqlParserListener) ExitPrivilege_level(ctx *Privilege_levelContext) {}

// EnterSet_password_option is called when production set_password_option is entered.
func (s *BaseMySqlParserListener) EnterSet_password_option(ctx *Set_password_optionContext) {}

// ExitSet_password_option is called when production set_password_option is exited.
func (s *BaseMySqlParserListener) ExitSet_password_option(ctx *Set_password_optionContext) {}

// EnterAnalyze_table is called when production analyze_table is entered.
func (s *BaseMySqlParserListener) EnterAnalyze_table(ctx *Analyze_tableContext) {}

// ExitAnalyze_table is called when production analyze_table is exited.
func (s *BaseMySqlParserListener) ExitAnalyze_table(ctx *Analyze_tableContext) {}

// EnterCheck_table is called when production check_table is entered.
func (s *BaseMySqlParserListener) EnterCheck_table(ctx *Check_tableContext) {}

// ExitCheck_table is called when production check_table is exited.
func (s *BaseMySqlParserListener) ExitCheck_table(ctx *Check_tableContext) {}

// EnterChecksum_table is called when production checksum_table is entered.
func (s *BaseMySqlParserListener) EnterChecksum_table(ctx *Checksum_tableContext) {}

// ExitChecksum_table is called when production checksum_table is exited.
func (s *BaseMySqlParserListener) ExitChecksum_table(ctx *Checksum_tableContext) {}

// EnterOptimize_table is called when production optimize_table is entered.
func (s *BaseMySqlParserListener) EnterOptimize_table(ctx *Optimize_tableContext) {}

// ExitOptimize_table is called when production optimize_table is exited.
func (s *BaseMySqlParserListener) ExitOptimize_table(ctx *Optimize_tableContext) {}

// EnterRepair_table is called when production repair_table is entered.
func (s *BaseMySqlParserListener) EnterRepair_table(ctx *Repair_tableContext) {}

// ExitRepair_table is called when production repair_table is exited.
func (s *BaseMySqlParserListener) ExitRepair_table(ctx *Repair_tableContext) {}

// EnterCheck_table_option is called when production check_table_option is entered.
func (s *BaseMySqlParserListener) EnterCheck_table_option(ctx *Check_table_optionContext) {}

// ExitCheck_table_option is called when production check_table_option is exited.
func (s *BaseMySqlParserListener) ExitCheck_table_option(ctx *Check_table_optionContext) {}

// EnterCreate_udfunction is called when production create_udfunction is entered.
func (s *BaseMySqlParserListener) EnterCreate_udfunction(ctx *Create_udfunctionContext) {}

// ExitCreate_udfunction is called when production create_udfunction is exited.
func (s *BaseMySqlParserListener) ExitCreate_udfunction(ctx *Create_udfunctionContext) {}

// EnterInstall_plugin is called when production install_plugin is entered.
func (s *BaseMySqlParserListener) EnterInstall_plugin(ctx *Install_pluginContext) {}

// ExitInstall_plugin is called when production install_plugin is exited.
func (s *BaseMySqlParserListener) ExitInstall_plugin(ctx *Install_pluginContext) {}

// EnterUninstall_plugin is called when production uninstall_plugin is entered.
func (s *BaseMySqlParserListener) EnterUninstall_plugin(ctx *Uninstall_pluginContext) {}

// ExitUninstall_plugin is called when production uninstall_plugin is exited.
func (s *BaseMySqlParserListener) ExitUninstall_plugin(ctx *Uninstall_pluginContext) {}

// EnterSetVariableAssignment is called when production setVariableAssignment is entered.
func (s *BaseMySqlParserListener) EnterSetVariableAssignment(ctx *SetVariableAssignmentContext) {}

// ExitSetVariableAssignment is called when production setVariableAssignment is exited.
func (s *BaseMySqlParserListener) ExitSetVariableAssignment(ctx *SetVariableAssignmentContext) {}

// EnterSetCharset is called when production setCharset is entered.
func (s *BaseMySqlParserListener) EnterSetCharset(ctx *SetCharsetContext) {}

// ExitSetCharset is called when production setCharset is exited.
func (s *BaseMySqlParserListener) ExitSetCharset(ctx *SetCharsetContext) {}

// EnterSetNames is called when production setNames is entered.
func (s *BaseMySqlParserListener) EnterSetNames(ctx *SetNamesContext) {}

// ExitSetNames is called when production setNames is exited.
func (s *BaseMySqlParserListener) ExitSetNames(ctx *SetNamesContext) {}

// EnterSetPasswordStatement is called when production setPasswordStatement is entered.
func (s *BaseMySqlParserListener) EnterSetPasswordStatement(ctx *SetPasswordStatementContext) {}

// ExitSetPasswordStatement is called when production setPasswordStatement is exited.
func (s *BaseMySqlParserListener) ExitSetPasswordStatement(ctx *SetPasswordStatementContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseMySqlParserListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseMySqlParserListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterSetAutocommit is called when production setAutocommit is entered.
func (s *BaseMySqlParserListener) EnterSetAutocommit(ctx *SetAutocommitContext) {}

// ExitSetAutocommit is called when production setAutocommit is exited.
func (s *BaseMySqlParserListener) ExitSetAutocommit(ctx *SetAutocommitContext) {}

// EnterShowMasterlogs is called when production showMasterlogs is entered.
func (s *BaseMySqlParserListener) EnterShowMasterlogs(ctx *ShowMasterlogsContext) {}

// ExitShowMasterlogs is called when production showMasterlogs is exited.
func (s *BaseMySqlParserListener) ExitShowMasterlogs(ctx *ShowMasterlogsContext) {}

// EnterShowLogevents is called when production showLogevents is entered.
func (s *BaseMySqlParserListener) EnterShowLogevents(ctx *ShowLogeventsContext) {}

// ExitShowLogevents is called when production showLogevents is exited.
func (s *BaseMySqlParserListener) ExitShowLogevents(ctx *ShowLogeventsContext) {}

// EnterShowObjWithFilter is called when production showObjWithFilter is entered.
func (s *BaseMySqlParserListener) EnterShowObjWithFilter(ctx *ShowObjWithFilterContext) {}

// ExitShowObjWithFilter is called when production showObjWithFilter is exited.
func (s *BaseMySqlParserListener) ExitShowObjWithFilter(ctx *ShowObjWithFilterContext) {}

// EnterShowColumns is called when production showColumns is entered.
func (s *BaseMySqlParserListener) EnterShowColumns(ctx *ShowColumnsContext) {}

// ExitShowColumns is called when production showColumns is exited.
func (s *BaseMySqlParserListener) ExitShowColumns(ctx *ShowColumnsContext) {}

// EnterShowCreateDb is called when production showCreateDb is entered.
func (s *BaseMySqlParserListener) EnterShowCreateDb(ctx *ShowCreateDbContext) {}

// ExitShowCreateDb is called when production showCreateDb is exited.
func (s *BaseMySqlParserListener) ExitShowCreateDb(ctx *ShowCreateDbContext) {}

// EnterShowCreateFullidobj is called when production showCreateFullidobj is entered.
func (s *BaseMySqlParserListener) EnterShowCreateFullidobj(ctx *ShowCreateFullidobjContext) {}

// ExitShowCreateFullidobj is called when production showCreateFullidobj is exited.
func (s *BaseMySqlParserListener) ExitShowCreateFullidobj(ctx *ShowCreateFullidobjContext) {}

// EnterShowCreateUser is called when production showCreateUser is entered.
func (s *BaseMySqlParserListener) EnterShowCreateUser(ctx *ShowCreateUserContext) {}

// ExitShowCreateUser is called when production showCreateUser is exited.
func (s *BaseMySqlParserListener) ExitShowCreateUser(ctx *ShowCreateUserContext) {}

// EnterShowEngine is called when production showEngine is entered.
func (s *BaseMySqlParserListener) EnterShowEngine(ctx *ShowEngineContext) {}

// ExitShowEngine is called when production showEngine is exited.
func (s *BaseMySqlParserListener) ExitShowEngine(ctx *ShowEngineContext) {}

// EnterShowGlobalinfo is called when production showGlobalinfo is entered.
func (s *BaseMySqlParserListener) EnterShowGlobalinfo(ctx *ShowGlobalinfoContext) {}

// ExitShowGlobalinfo is called when production showGlobalinfo is exited.
func (s *BaseMySqlParserListener) ExitShowGlobalinfo(ctx *ShowGlobalinfoContext) {}

// EnterShowErrWarn is called when production showErrWarn is entered.
func (s *BaseMySqlParserListener) EnterShowErrWarn(ctx *ShowErrWarnContext) {}

// ExitShowErrWarn is called when production showErrWarn is exited.
func (s *BaseMySqlParserListener) ExitShowErrWarn(ctx *ShowErrWarnContext) {}

// EnterShowCountErrWarn is called when production showCountErrWarn is entered.
func (s *BaseMySqlParserListener) EnterShowCountErrWarn(ctx *ShowCountErrWarnContext) {}

// ExitShowCountErrWarn is called when production showCountErrWarn is exited.
func (s *BaseMySqlParserListener) ExitShowCountErrWarn(ctx *ShowCountErrWarnContext) {}

// EnterShowFromschemaFilter is called when production showFromschemaFilter is entered.
func (s *BaseMySqlParserListener) EnterShowFromschemaFilter(ctx *ShowFromschemaFilterContext) {}

// ExitShowFromschemaFilter is called when production showFromschemaFilter is exited.
func (s *BaseMySqlParserListener) ExitShowFromschemaFilter(ctx *ShowFromschemaFilterContext) {}

// EnterShowRoutinecode is called when production showRoutinecode is entered.
func (s *BaseMySqlParserListener) EnterShowRoutinecode(ctx *ShowRoutinecodeContext) {}

// ExitShowRoutinecode is called when production showRoutinecode is exited.
func (s *BaseMySqlParserListener) ExitShowRoutinecode(ctx *ShowRoutinecodeContext) {}

// EnterShowGrants is called when production showGrants is entered.
func (s *BaseMySqlParserListener) EnterShowGrants(ctx *ShowGrantsContext) {}

// ExitShowGrants is called when production showGrants is exited.
func (s *BaseMySqlParserListener) ExitShowGrants(ctx *ShowGrantsContext) {}

// EnterShowIndexes is called when production showIndexes is entered.
func (s *BaseMySqlParserListener) EnterShowIndexes(ctx *ShowIndexesContext) {}

// ExitShowIndexes is called when production showIndexes is exited.
func (s *BaseMySqlParserListener) ExitShowIndexes(ctx *ShowIndexesContext) {}

// EnterShowOpentables is called when production showOpentables is entered.
func (s *BaseMySqlParserListener) EnterShowOpentables(ctx *ShowOpentablesContext) {}

// ExitShowOpentables is called when production showOpentables is exited.
func (s *BaseMySqlParserListener) ExitShowOpentables(ctx *ShowOpentablesContext) {}

// EnterShowProfile is called when production showProfile is entered.
func (s *BaseMySqlParserListener) EnterShowProfile(ctx *ShowProfileContext) {}

// ExitShowProfile is called when production showProfile is exited.
func (s *BaseMySqlParserListener) ExitShowProfile(ctx *ShowProfileContext) {}

// EnterShowSlavestatus is called when production showSlavestatus is entered.
func (s *BaseMySqlParserListener) EnterShowSlavestatus(ctx *ShowSlavestatusContext) {}

// ExitShowSlavestatus is called when production showSlavestatus is exited.
func (s *BaseMySqlParserListener) ExitShowSlavestatus(ctx *ShowSlavestatusContext) {}

// EnterVariable_clause is called when production variable_clause is entered.
func (s *BaseMySqlParserListener) EnterVariable_clause(ctx *Variable_clauseContext) {}

// ExitVariable_clause is called when production variable_clause is exited.
func (s *BaseMySqlParserListener) ExitVariable_clause(ctx *Variable_clauseContext) {}

// EnterShow_filter is called when production show_filter is entered.
func (s *BaseMySqlParserListener) EnterShow_filter(ctx *Show_filterContext) {}

// ExitShow_filter is called when production show_filter is exited.
func (s *BaseMySqlParserListener) ExitShow_filter(ctx *Show_filterContext) {}

// EnterShow_profile_type is called when production show_profile_type is entered.
func (s *BaseMySqlParserListener) EnterShow_profile_type(ctx *Show_profile_typeContext) {}

// ExitShow_profile_type is called when production show_profile_type is exited.
func (s *BaseMySqlParserListener) ExitShow_profile_type(ctx *Show_profile_typeContext) {}

// EnterBinlog_statement is called when production binlog_statement is entered.
func (s *BaseMySqlParserListener) EnterBinlog_statement(ctx *Binlog_statementContext) {}

// ExitBinlog_statement is called when production binlog_statement is exited.
func (s *BaseMySqlParserListener) ExitBinlog_statement(ctx *Binlog_statementContext) {}

// EnterCache_index_statement is called when production cache_index_statement is entered.
func (s *BaseMySqlParserListener) EnterCache_index_statement(ctx *Cache_index_statementContext) {}

// ExitCache_index_statement is called when production cache_index_statement is exited.
func (s *BaseMySqlParserListener) ExitCache_index_statement(ctx *Cache_index_statementContext) {}

// EnterFlush_statement is called when production flush_statement is entered.
func (s *BaseMySqlParserListener) EnterFlush_statement(ctx *Flush_statementContext) {}

// ExitFlush_statement is called when production flush_statement is exited.
func (s *BaseMySqlParserListener) ExitFlush_statement(ctx *Flush_statementContext) {}

// EnterKill_statement is called when production kill_statement is entered.
func (s *BaseMySqlParserListener) EnterKill_statement(ctx *Kill_statementContext) {}

// ExitKill_statement is called when production kill_statement is exited.
func (s *BaseMySqlParserListener) ExitKill_statement(ctx *Kill_statementContext) {}

// EnterLoad_index_into_cache is called when production load_index_into_cache is entered.
func (s *BaseMySqlParserListener) EnterLoad_index_into_cache(ctx *Load_index_into_cacheContext) {}

// ExitLoad_index_into_cache is called when production load_index_into_cache is exited.
func (s *BaseMySqlParserListener) ExitLoad_index_into_cache(ctx *Load_index_into_cacheContext) {}

// EnterReset_statement is called when production reset_statement is entered.
func (s *BaseMySqlParserListener) EnterReset_statement(ctx *Reset_statementContext) {}

// ExitReset_statement is called when production reset_statement is exited.
func (s *BaseMySqlParserListener) ExitReset_statement(ctx *Reset_statementContext) {}

// EnterShutdown_statement is called when production shutdown_statement is entered.
func (s *BaseMySqlParserListener) EnterShutdown_statement(ctx *Shutdown_statementContext) {}

// ExitShutdown_statement is called when production shutdown_statement is exited.
func (s *BaseMySqlParserListener) ExitShutdown_statement(ctx *Shutdown_statementContext) {}

// EnterTbl_index_list is called when production tbl_index_list is entered.
func (s *BaseMySqlParserListener) EnterTbl_index_list(ctx *Tbl_index_listContext) {}

// ExitTbl_index_list is called when production tbl_index_list is exited.
func (s *BaseMySqlParserListener) ExitTbl_index_list(ctx *Tbl_index_listContext) {}

// EnterFlush_option is called when production flush_option is entered.
func (s *BaseMySqlParserListener) EnterFlush_option(ctx *Flush_optionContext) {}

// ExitFlush_option is called when production flush_option is exited.
func (s *BaseMySqlParserListener) ExitFlush_option(ctx *Flush_optionContext) {}

// EnterLoad_tbl_index_list is called when production load_tbl_index_list is entered.
func (s *BaseMySqlParserListener) EnterLoad_tbl_index_list(ctx *Load_tbl_index_listContext) {}

// ExitLoad_tbl_index_list is called when production load_tbl_index_list is exited.
func (s *BaseMySqlParserListener) ExitLoad_tbl_index_list(ctx *Load_tbl_index_listContext) {}

// EnterSimple_describe_statement is called when production simple_describe_statement is entered.
func (s *BaseMySqlParserListener) EnterSimple_describe_statement(ctx *Simple_describe_statementContext) {
}

// ExitSimple_describe_statement is called when production simple_describe_statement is exited.
func (s *BaseMySqlParserListener) ExitSimple_describe_statement(ctx *Simple_describe_statementContext) {
}

// EnterFull_describe_statement is called when production full_describe_statement is entered.
func (s *BaseMySqlParserListener) EnterFull_describe_statement(ctx *Full_describe_statementContext) {}

// ExitFull_describe_statement is called when production full_describe_statement is exited.
func (s *BaseMySqlParserListener) ExitFull_describe_statement(ctx *Full_describe_statementContext) {}

// EnterHelp_statement is called when production help_statement is entered.
func (s *BaseMySqlParserListener) EnterHelp_statement(ctx *Help_statementContext) {}

// ExitHelp_statement is called when production help_statement is exited.
func (s *BaseMySqlParserListener) ExitHelp_statement(ctx *Help_statementContext) {}

// EnterUse_statement is called when production use_statement is entered.
func (s *BaseMySqlParserListener) EnterUse_statement(ctx *Use_statementContext) {}

// ExitUse_statement is called when production use_statement is exited.
func (s *BaseMySqlParserListener) ExitUse_statement(ctx *Use_statementContext) {}

// EnterDescstmtDescObj is called when production descstmtDescObj is entered.
func (s *BaseMySqlParserListener) EnterDescstmtDescObj(ctx *DescstmtDescObjContext) {}

// ExitDescstmtDescObj is called when production descstmtDescObj is exited.
func (s *BaseMySqlParserListener) ExitDescstmtDescObj(ctx *DescstmtDescObjContext) {}

// EnterConnectionDescObj is called when production connectionDescObj is entered.
func (s *BaseMySqlParserListener) EnterConnectionDescObj(ctx *ConnectionDescObjContext) {}

// ExitConnectionDescObj is called when production connectionDescObj is exited.
func (s *BaseMySqlParserListener) ExitConnectionDescObj(ctx *ConnectionDescObjContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseMySqlParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseMySqlParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterFull_id is called when production full_id is entered.
func (s *BaseMySqlParserListener) EnterFull_id(ctx *Full_idContext) {}

// ExitFull_id is called when production full_id is exited.
func (s *BaseMySqlParserListener) ExitFull_id(ctx *Full_idContext) {}

// EnterFull_column_name is called when production full_column_name is entered.
func (s *BaseMySqlParserListener) EnterFull_column_name(ctx *Full_column_nameContext) {}

// ExitFull_column_name is called when production full_column_name is exited.
func (s *BaseMySqlParserListener) ExitFull_column_name(ctx *Full_column_nameContext) {}

// EnterIndex_col_name is called when production index_col_name is entered.
func (s *BaseMySqlParserListener) EnterIndex_col_name(ctx *Index_col_nameContext) {}

// ExitIndex_col_name is called when production index_col_name is exited.
func (s *BaseMySqlParserListener) ExitIndex_col_name(ctx *Index_col_nameContext) {}

// EnterUser_name is called when production user_name is entered.
func (s *BaseMySqlParserListener) EnterUser_name(ctx *User_nameContext) {}

// ExitUser_name is called when production user_name is exited.
func (s *BaseMySqlParserListener) ExitUser_name(ctx *User_nameContext) {}

// EnterMysql_variable is called when production mysql_variable is entered.
func (s *BaseMySqlParserListener) EnterMysql_variable(ctx *Mysql_variableContext) {}

// ExitMysql_variable is called when production mysql_variable is exited.
func (s *BaseMySqlParserListener) ExitMysql_variable(ctx *Mysql_variableContext) {}

// EnterCharset_name is called when production charset_name is entered.
func (s *BaseMySqlParserListener) EnterCharset_name(ctx *Charset_nameContext) {}

// ExitCharset_name is called when production charset_name is exited.
func (s *BaseMySqlParserListener) ExitCharset_name(ctx *Charset_nameContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseMySqlParserListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseMySqlParserListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterEngine_name is called when production engine_name is entered.
func (s *BaseMySqlParserListener) EnterEngine_name(ctx *Engine_nameContext) {}

// ExitEngine_name is called when production engine_name is exited.
func (s *BaseMySqlParserListener) ExitEngine_name(ctx *Engine_nameContext) {}

// EnterUuid_set is called when production uuid_set is entered.
func (s *BaseMySqlParserListener) EnterUuid_set(ctx *Uuid_setContext) {}

// ExitUuid_set is called when production uuid_set is exited.
func (s *BaseMySqlParserListener) ExitUuid_set(ctx *Uuid_setContext) {}

// EnterXid is called when production xid is entered.
func (s *BaseMySqlParserListener) EnterXid(ctx *XidContext) {}

// ExitXid is called when production xid is exited.
func (s *BaseMySqlParserListener) ExitXid(ctx *XidContext) {}

// EnterXid_string_id is called when production xid_string_id is entered.
func (s *BaseMySqlParserListener) EnterXid_string_id(ctx *Xid_string_idContext) {}

// ExitXid_string_id is called when production xid_string_id is exited.
func (s *BaseMySqlParserListener) ExitXid_string_id(ctx *Xid_string_idContext) {}

// EnterAuth_plugin is called when production auth_plugin is entered.
func (s *BaseMySqlParserListener) EnterAuth_plugin(ctx *Auth_pluginContext) {}

// ExitAuth_plugin is called when production auth_plugin is exited.
func (s *BaseMySqlParserListener) ExitAuth_plugin(ctx *Auth_pluginContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseMySqlParserListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseMySqlParserListener) ExitId_(ctx *Id_Context) {}

// EnterSimple_id is called when production simple_id is entered.
func (s *BaseMySqlParserListener) EnterSimple_id(ctx *Simple_idContext) {}

// ExitSimple_id is called when production simple_id is exited.
func (s *BaseMySqlParserListener) ExitSimple_id(ctx *Simple_idContext) {}

// EnterDot_ext_id is called when production dot_ext_id is entered.
func (s *BaseMySqlParserListener) EnterDot_ext_id(ctx *Dot_ext_idContext) {}

// ExitDot_ext_id is called when production dot_ext_id is exited.
func (s *BaseMySqlParserListener) ExitDot_ext_id(ctx *Dot_ext_idContext) {}

// EnterDecimal_literal is called when production decimal_literal is entered.
func (s *BaseMySqlParserListener) EnterDecimal_literal(ctx *Decimal_literalContext) {}

// ExitDecimal_literal is called when production decimal_literal is exited.
func (s *BaseMySqlParserListener) ExitDecimal_literal(ctx *Decimal_literalContext) {}

// EnterFilesize_literal is called when production filesize_literal is entered.
func (s *BaseMySqlParserListener) EnterFilesize_literal(ctx *Filesize_literalContext) {}

// ExitFilesize_literal is called when production filesize_literal is exited.
func (s *BaseMySqlParserListener) ExitFilesize_literal(ctx *Filesize_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseMySqlParserListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseMySqlParserListener) ExitString_literal(ctx *String_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseMySqlParserListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseMySqlParserListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterHexadecimal_literal is called when production hexadecimal_literal is entered.
func (s *BaseMySqlParserListener) EnterHexadecimal_literal(ctx *Hexadecimal_literalContext) {}

// ExitHexadecimal_literal is called when production hexadecimal_literal is exited.
func (s *BaseMySqlParserListener) ExitHexadecimal_literal(ctx *Hexadecimal_literalContext) {}

// EnterNull_notnull is called when production null_notnull is entered.
func (s *BaseMySqlParserListener) EnterNull_notnull(ctx *Null_notnullContext) {}

// ExitNull_notnull is called when production null_notnull is exited.
func (s *BaseMySqlParserListener) ExitNull_notnull(ctx *Null_notnullContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseMySqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseMySqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterCharDatatype is called when production charDatatype is entered.
func (s *BaseMySqlParserListener) EnterCharDatatype(ctx *CharDatatypeContext) {}

// ExitCharDatatype is called when production charDatatype is exited.
func (s *BaseMySqlParserListener) ExitCharDatatype(ctx *CharDatatypeContext) {}

// EnterDimensionDatatype is called when production dimensionDatatype is entered.
func (s *BaseMySqlParserListener) EnterDimensionDatatype(ctx *DimensionDatatypeContext) {}

// ExitDimensionDatatype is called when production dimensionDatatype is exited.
func (s *BaseMySqlParserListener) ExitDimensionDatatype(ctx *DimensionDatatypeContext) {}

// EnterSimpleDatatype is called when production simpleDatatype is entered.
func (s *BaseMySqlParserListener) EnterSimpleDatatype(ctx *SimpleDatatypeContext) {}

// ExitSimpleDatatype is called when production simpleDatatype is exited.
func (s *BaseMySqlParserListener) ExitSimpleDatatype(ctx *SimpleDatatypeContext) {}

// EnterCollectCharDatatype is called when production collectCharDatatype is entered.
func (s *BaseMySqlParserListener) EnterCollectCharDatatype(ctx *CollectCharDatatypeContext) {}

// ExitCollectCharDatatype is called when production collectCharDatatype is exited.
func (s *BaseMySqlParserListener) ExitCollectCharDatatype(ctx *CollectCharDatatypeContext) {}

// EnterSpatialDatatype is called when production spatialDatatype is entered.
func (s *BaseMySqlParserListener) EnterSpatialDatatype(ctx *SpatialDatatypeContext) {}

// ExitSpatialDatatype is called when production spatialDatatype is exited.
func (s *BaseMySqlParserListener) ExitSpatialDatatype(ctx *SpatialDatatypeContext) {}

// EnterData_type_to_convert is called when production data_type_to_convert is entered.
func (s *BaseMySqlParserListener) EnterData_type_to_convert(ctx *Data_type_to_convertContext) {}

// ExitData_type_to_convert is called when production data_type_to_convert is exited.
func (s *BaseMySqlParserListener) ExitData_type_to_convert(ctx *Data_type_to_convertContext) {}

// EnterSpatial_data_type is called when production spatial_data_type is entered.
func (s *BaseMySqlParserListener) EnterSpatial_data_type(ctx *Spatial_data_typeContext) {}

// ExitSpatial_data_type is called when production spatial_data_type is exited.
func (s *BaseMySqlParserListener) ExitSpatial_data_type(ctx *Spatial_data_typeContext) {}

// EnterLength_one_dimension is called when production length_one_dimension is entered.
func (s *BaseMySqlParserListener) EnterLength_one_dimension(ctx *Length_one_dimensionContext) {}

// ExitLength_one_dimension is called when production length_one_dimension is exited.
func (s *BaseMySqlParserListener) ExitLength_one_dimension(ctx *Length_one_dimensionContext) {}

// EnterLength_two_dimension is called when production length_two_dimension is entered.
func (s *BaseMySqlParserListener) EnterLength_two_dimension(ctx *Length_two_dimensionContext) {}

// ExitLength_two_dimension is called when production length_two_dimension is exited.
func (s *BaseMySqlParserListener) ExitLength_two_dimension(ctx *Length_two_dimensionContext) {}

// EnterLength_two_optional_dimension is called when production length_two_optional_dimension is entered.
func (s *BaseMySqlParserListener) EnterLength_two_optional_dimension(ctx *Length_two_optional_dimensionContext) {
}

// ExitLength_two_optional_dimension is called when production length_two_optional_dimension is exited.
func (s *BaseMySqlParserListener) ExitLength_two_optional_dimension(ctx *Length_two_optional_dimensionContext) {
}

// EnterId_list is called when production id_list is entered.
func (s *BaseMySqlParserListener) EnterId_list(ctx *Id_listContext) {}

// ExitId_list is called when production id_list is exited.
func (s *BaseMySqlParserListener) ExitId_list(ctx *Id_listContext) {}

// EnterTable_list is called when production table_list is entered.
func (s *BaseMySqlParserListener) EnterTable_list(ctx *Table_listContext) {}

// ExitTable_list is called when production table_list is exited.
func (s *BaseMySqlParserListener) ExitTable_list(ctx *Table_listContext) {}

// EnterTable_pair_list is called when production table_pair_list is entered.
func (s *BaseMySqlParserListener) EnterTable_pair_list(ctx *Table_pair_listContext) {}

// ExitTable_pair_list is called when production table_pair_list is exited.
func (s *BaseMySqlParserListener) ExitTable_pair_list(ctx *Table_pair_listContext) {}

// EnterIndex_colname_list is called when production index_colname_list is entered.
func (s *BaseMySqlParserListener) EnterIndex_colname_list(ctx *Index_colname_listContext) {}

// ExitIndex_colname_list is called when production index_colname_list is exited.
func (s *BaseMySqlParserListener) ExitIndex_colname_list(ctx *Index_colname_listContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseMySqlParserListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseMySqlParserListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterConstant_list is called when production constant_list is entered.
func (s *BaseMySqlParserListener) EnterConstant_list(ctx *Constant_listContext) {}

// ExitConstant_list is called when production constant_list is exited.
func (s *BaseMySqlParserListener) ExitConstant_list(ctx *Constant_listContext) {}

// EnterSimple_string_list is called when production simple_string_list is entered.
func (s *BaseMySqlParserListener) EnterSimple_string_list(ctx *Simple_string_listContext) {}

// ExitSimple_string_list is called when production simple_string_list is exited.
func (s *BaseMySqlParserListener) ExitSimple_string_list(ctx *Simple_string_listContext) {}

// EnterUser_var_list is called when production user_var_list is entered.
func (s *BaseMySqlParserListener) EnterUser_var_list(ctx *User_var_listContext) {}

// ExitUser_var_list is called when production user_var_list is exited.
func (s *BaseMySqlParserListener) ExitUser_var_list(ctx *User_var_listContext) {}

// EnterDefault_value is called when production default_value is entered.
func (s *BaseMySqlParserListener) EnterDefault_value(ctx *Default_valueContext) {}

// ExitDefault_value is called when production default_value is exited.
func (s *BaseMySqlParserListener) ExitDefault_value(ctx *Default_valueContext) {}

// EnterIf_exists is called when production if_exists is entered.
func (s *BaseMySqlParserListener) EnterIf_exists(ctx *If_existsContext) {}

// ExitIf_exists is called when production if_exists is exited.
func (s *BaseMySqlParserListener) ExitIf_exists(ctx *If_existsContext) {}

// EnterIf_not_exists is called when production if_not_exists is entered.
func (s *BaseMySqlParserListener) EnterIf_not_exists(ctx *If_not_existsContext) {}

// ExitIf_not_exists is called when production if_not_exists is exited.
func (s *BaseMySqlParserListener) ExitIf_not_exists(ctx *If_not_existsContext) {}

// EnterSpecificFunctionCall is called when production specificFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// ExitSpecificFunctionCall is called when production specificFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// EnterAggregateFunctionCall is called when production aggregateFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// ExitAggregateFunctionCall is called when production aggregateFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// EnterScalarFunctionCall is called when production scalarFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// ExitScalarFunctionCall is called when production scalarFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// EnterUdfFunctionCall is called when production udfFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// ExitUdfFunctionCall is called when production udfFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// EnterSimpleSpecificFCall is called when production simpleSpecificFCall is entered.
func (s *BaseMySqlParserListener) EnterSimpleSpecificFCall(ctx *SimpleSpecificFCallContext) {}

// ExitSimpleSpecificFCall is called when production simpleSpecificFCall is exited.
func (s *BaseMySqlParserListener) ExitSimpleSpecificFCall(ctx *SimpleSpecificFCallContext) {}

// EnterConvertDataTypeFCall is called when production convertDataTypeFCall is entered.
func (s *BaseMySqlParserListener) EnterConvertDataTypeFCall(ctx *ConvertDataTypeFCallContext) {}

// ExitConvertDataTypeFCall is called when production convertDataTypeFCall is exited.
func (s *BaseMySqlParserListener) ExitConvertDataTypeFCall(ctx *ConvertDataTypeFCallContext) {}

// EnterValuesFCall is called when production valuesFCall is entered.
func (s *BaseMySqlParserListener) EnterValuesFCall(ctx *ValuesFCallContext) {}

// ExitValuesFCall is called when production valuesFCall is exited.
func (s *BaseMySqlParserListener) ExitValuesFCall(ctx *ValuesFCallContext) {}

// EnterCaseFCall is called when production caseFCall is entered.
func (s *BaseMySqlParserListener) EnterCaseFCall(ctx *CaseFCallContext) {}

// ExitCaseFCall is called when production caseFCall is exited.
func (s *BaseMySqlParserListener) ExitCaseFCall(ctx *CaseFCallContext) {}

// EnterCharFCall is called when production charFCall is entered.
func (s *BaseMySqlParserListener) EnterCharFCall(ctx *CharFCallContext) {}

// ExitCharFCall is called when production charFCall is exited.
func (s *BaseMySqlParserListener) ExitCharFCall(ctx *CharFCallContext) {}

// EnterPositionFCall is called when production positionFCall is entered.
func (s *BaseMySqlParserListener) EnterPositionFCall(ctx *PositionFCallContext) {}

// ExitPositionFCall is called when production positionFCall is exited.
func (s *BaseMySqlParserListener) ExitPositionFCall(ctx *PositionFCallContext) {}

// EnterSubstrFCall is called when production substrFCall is entered.
func (s *BaseMySqlParserListener) EnterSubstrFCall(ctx *SubstrFCallContext) {}

// ExitSubstrFCall is called when production substrFCall is exited.
func (s *BaseMySqlParserListener) ExitSubstrFCall(ctx *SubstrFCallContext) {}

// EnterTrimFCall is called when production trimFCall is entered.
func (s *BaseMySqlParserListener) EnterTrimFCall(ctx *TrimFCallContext) {}

// ExitTrimFCall is called when production trimFCall is exited.
func (s *BaseMySqlParserListener) ExitTrimFCall(ctx *TrimFCallContext) {}

// EnterWeightFCall is called when production weightFCall is entered.
func (s *BaseMySqlParserListener) EnterWeightFCall(ctx *WeightFCallContext) {}

// ExitWeightFCall is called when production weightFCall is exited.
func (s *BaseMySqlParserListener) ExitWeightFCall(ctx *WeightFCallContext) {}

// EnterExtractFCall is called when production extractFCall is entered.
func (s *BaseMySqlParserListener) EnterExtractFCall(ctx *ExtractFCallContext) {}

// ExitExtractFCall is called when production extractFCall is exited.
func (s *BaseMySqlParserListener) ExitExtractFCall(ctx *ExtractFCallContext) {}

// EnterGetFormatFCall is called when production getFormatFCall is entered.
func (s *BaseMySqlParserListener) EnterGetFormatFCall(ctx *GetFormatFCallContext) {}

// ExitGetFormatFCall is called when production getFormatFCall is exited.
func (s *BaseMySqlParserListener) ExitGetFormatFCall(ctx *GetFormatFCallContext) {}

// EnterLevelWeightFList is called when production levelWeightFList is entered.
func (s *BaseMySqlParserListener) EnterLevelWeightFList(ctx *LevelWeightFListContext) {}

// ExitLevelWeightFList is called when production levelWeightFList is exited.
func (s *BaseMySqlParserListener) ExitLevelWeightFList(ctx *LevelWeightFListContext) {}

// EnterLevelWeightFRange is called when production levelWeightFRange is entered.
func (s *BaseMySqlParserListener) EnterLevelWeightFRange(ctx *LevelWeightFRangeContext) {}

// ExitLevelWeightFRange is called when production levelWeightFRange is exited.
func (s *BaseMySqlParserListener) ExitLevelWeightFRange(ctx *LevelWeightFRangeContext) {}

// EnterAggregate_windowed_function is called when production aggregate_windowed_function is entered.
func (s *BaseMySqlParserListener) EnterAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
}

// ExitAggregate_windowed_function is called when production aggregate_windowed_function is exited.
func (s *BaseMySqlParserListener) ExitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
}

// EnterScalar_function_name is called when production scalar_function_name is entered.
func (s *BaseMySqlParserListener) EnterScalar_function_name(ctx *Scalar_function_nameContext) {}

// ExitScalar_function_name is called when production scalar_function_name is exited.
func (s *BaseMySqlParserListener) ExitScalar_function_name(ctx *Scalar_function_nameContext) {}

// EnterFunction_args is called when production function_args is entered.
func (s *BaseMySqlParserListener) EnterFunction_args(ctx *Function_argsContext) {}

// ExitFunction_args is called when production function_args is exited.
func (s *BaseMySqlParserListener) ExitFunction_args(ctx *Function_argsContext) {}

// EnterFunction_arg is called when production function_arg is entered.
func (s *BaseMySqlParserListener) EnterFunction_arg(ctx *Function_argContext) {}

// ExitFunction_arg is called when production function_arg is exited.
func (s *BaseMySqlParserListener) ExitFunction_arg(ctx *Function_argContext) {}

// EnterIsExpression is called when production isExpression is entered.
func (s *BaseMySqlParserListener) EnterIsExpression(ctx *IsExpressionContext) {}

// ExitIsExpression is called when production isExpression is exited.
func (s *BaseMySqlParserListener) ExitIsExpression(ctx *IsExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseMySqlParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseMySqlParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseMySqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseMySqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseMySqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseMySqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterSoundsLikePredicate is called when production soundsLikePredicate is entered.
func (s *BaseMySqlParserListener) EnterSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// ExitSoundsLikePredicate is called when production soundsLikePredicate is exited.
func (s *BaseMySqlParserListener) ExitSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseMySqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseMySqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseMySqlParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseMySqlParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {
}

// ExitSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {
}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseMySqlParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseMySqlParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterBinaryComparasionPredicate is called when production binaryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// ExitBinaryComparasionPredicate is called when production binaryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseMySqlParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseMySqlParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseMySqlParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseMySqlParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterRegexpPredicate is called when production regexpPredicate is entered.
func (s *BaseMySqlParserListener) EnterRegexpPredicate(ctx *RegexpPredicateContext) {}

// ExitRegexpPredicate is called when production regexpPredicate is exited.
func (s *BaseMySqlParserListener) ExitRegexpPredicate(ctx *RegexpPredicateContext) {}

// EnterUnaryExpressionAtom is called when production unaryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// ExitUnaryExpressionAtom is called when production unaryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// EnterExistsExpessionAtom is called when production existsExpessionAtom is entered.
func (s *BaseMySqlParserListener) EnterExistsExpessionAtom(ctx *ExistsExpessionAtomContext) {}

// ExitExistsExpessionAtom is called when production existsExpessionAtom is exited.
func (s *BaseMySqlParserListener) ExitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) {}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterFunctionCallExpressionAtom is called when production functionCallExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// ExitFunctionCallExpressionAtom is called when production functionCallExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// EnterMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// ExitMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// EnterBinaryExpressionAtom is called when production binaryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// ExitBinaryExpressionAtom is called when production binaryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// EnterFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// ExitFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// EnterDefaultExpressionAtom is called when production defaultExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterDefaultExpressionAtom(ctx *DefaultExpressionAtomContext) {}

// ExitDefaultExpressionAtom is called when production defaultExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitDefaultExpressionAtom(ctx *DefaultExpressionAtomContext) {}

// EnterBitExpressionAtom is called when production bitExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// ExitBitExpressionAtom is called when production bitExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// EnterNestedExpressionAtom is called when production nestedExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// ExitNestedExpressionAtom is called when production nestedExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// EnterMathExpressionAtom is called when production mathExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// ExitMathExpressionAtom is called when production mathExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// EnterIntervalExpressionAtom is called when production intervalExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// ExitIntervalExpressionAtom is called when production intervalExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseMySqlParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseMySqlParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseMySqlParserListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseMySqlParserListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterLogical_operator is called when production logical_operator is entered.
func (s *BaseMySqlParserListener) EnterLogical_operator(ctx *Logical_operatorContext) {}

// ExitLogical_operator is called when production logical_operator is exited.
func (s *BaseMySqlParserListener) ExitLogical_operator(ctx *Logical_operatorContext) {}

// EnterBit_operator is called when production bit_operator is entered.
func (s *BaseMySqlParserListener) EnterBit_operator(ctx *Bit_operatorContext) {}

// ExitBit_operator is called when production bit_operator is exited.
func (s *BaseMySqlParserListener) ExitBit_operator(ctx *Bit_operatorContext) {}

// EnterMath_operator is called when production math_operator is entered.
func (s *BaseMySqlParserListener) EnterMath_operator(ctx *Math_operatorContext) {}

// ExitMath_operator is called when production math_operator is exited.
func (s *BaseMySqlParserListener) ExitMath_operator(ctx *Math_operatorContext) {}

// EnterCharset_name_base is called when production charset_name_base is entered.
func (s *BaseMySqlParserListener) EnterCharset_name_base(ctx *Charset_name_baseContext) {}

// ExitCharset_name_base is called when production charset_name_base is exited.
func (s *BaseMySqlParserListener) ExitCharset_name_base(ctx *Charset_name_baseContext) {}

// EnterTransaction_level_base is called when production transaction_level_base is entered.
func (s *BaseMySqlParserListener) EnterTransaction_level_base(ctx *Transaction_level_baseContext) {}

// ExitTransaction_level_base is called when production transaction_level_base is exited.
func (s *BaseMySqlParserListener) ExitTransaction_level_base(ctx *Transaction_level_baseContext) {}

// EnterPrivileges_base is called when production privileges_base is entered.
func (s *BaseMySqlParserListener) EnterPrivileges_base(ctx *Privileges_baseContext) {}

// ExitPrivileges_base is called when production privileges_base is exited.
func (s *BaseMySqlParserListener) ExitPrivileges_base(ctx *Privileges_baseContext) {}

// EnterInterval_type_base is called when production interval_type_base is entered.
func (s *BaseMySqlParserListener) EnterInterval_type_base(ctx *Interval_type_baseContext) {}

// ExitInterval_type_base is called when production interval_type_base is exited.
func (s *BaseMySqlParserListener) ExitInterval_type_base(ctx *Interval_type_baseContext) {}

// EnterData_type_base is called when production data_type_base is entered.
func (s *BaseMySqlParserListener) EnterData_type_base(ctx *Data_type_baseContext) {}

// ExitData_type_base is called when production data_type_base is exited.
func (s *BaseMySqlParserListener) ExitData_type_base(ctx *Data_type_baseContext) {}

// EnterKeywords_can_be_id is called when production keywords_can_be_id is entered.
func (s *BaseMySqlParserListener) EnterKeywords_can_be_id(ctx *Keywords_can_be_idContext) {}

// ExitKeywords_can_be_id is called when production keywords_can_be_id is exited.
func (s *BaseMySqlParserListener) ExitKeywords_can_be_id(ctx *Keywords_can_be_idContext) {}

// EnterFunction_name_base is called when production function_name_base is entered.
func (s *BaseMySqlParserListener) EnterFunction_name_base(ctx *Function_name_baseContext) {}

// ExitFunction_name_base is called when production function_name_base is exited.
func (s *BaseMySqlParserListener) ExitFunction_name_base(ctx *Function_name_baseContext) {}
