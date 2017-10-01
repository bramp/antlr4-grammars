// Generated from FusionTablesSql.g4 by ANTLR 4.7.

package fusion_tables // FusionTablesSql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFusionTablesSqlListener is a complete listener for a parse tree produced by FusionTablesSqlParser.
type BaseFusionTablesSqlListener struct{}

var _ FusionTablesSqlListener = &BaseFusionTablesSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFusionTablesSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFusionTablesSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFusionTablesSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFusionTablesSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFusionTablesSql is called when production fusionTablesSql is entered.
func (s *BaseFusionTablesSqlListener) EnterFusionTablesSql(ctx *FusionTablesSqlContext) {}

// ExitFusionTablesSql is called when production fusionTablesSql is exited.
func (s *BaseFusionTablesSqlListener) ExitFusionTablesSql(ctx *FusionTablesSqlContext) {}

// EnterSql_stmt is called when production sql_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterSql_stmt(ctx *Sql_stmtContext) {}

// ExitSql_stmt is called when production sql_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitSql_stmt(ctx *Sql_stmtContext) {}

// EnterTable_name_in_ddl is called when production table_name_in_ddl is entered.
func (s *BaseFusionTablesSqlListener) EnterTable_name_in_ddl(ctx *Table_name_in_ddlContext) {}

// ExitTable_name_in_ddl is called when production table_name_in_ddl is exited.
func (s *BaseFusionTablesSqlListener) ExitTable_name_in_ddl(ctx *Table_name_in_ddlContext) {}

// EnterTable_name_in_dml is called when production table_name_in_dml is entered.
func (s *BaseFusionTablesSqlListener) EnterTable_name_in_dml(ctx *Table_name_in_dmlContext) {}

// ExitTable_name_in_dml is called when production table_name_in_dml is exited.
func (s *BaseFusionTablesSqlListener) ExitTable_name_in_dml(ctx *Table_name_in_dmlContext) {}

// EnterCreate_table_as_select_stmt is called when production create_table_as_select_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterCreate_table_as_select_stmt(ctx *Create_table_as_select_stmtContext) {
}

// ExitCreate_table_as_select_stmt is called when production create_table_as_select_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitCreate_table_as_select_stmt(ctx *Create_table_as_select_stmtContext) {
}

// EnterDescribe_stmt is called when production describe_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterDescribe_stmt(ctx *Describe_stmtContext) {}

// ExitDescribe_stmt is called when production describe_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitDescribe_stmt(ctx *Describe_stmtContext) {}

// EnterShow_tables_stmt is called when production show_tables_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterShow_tables_stmt(ctx *Show_tables_stmtContext) {}

// ExitShow_tables_stmt is called when production show_tables_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitShow_tables_stmt(ctx *Show_tables_stmtContext) {}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterCreate_view_stmt is called when production create_view_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterCreate_view_stmt(ctx *Create_view_stmtContext) {}

// ExitCreate_view_stmt is called when production create_view_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitCreate_view_stmt(ctx *Create_view_stmtContext) {}

// EnterDrop_table_stmt is called when production drop_table_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// ExitDrop_table_stmt is called when production drop_table_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterColumn_assignment is called when production column_assignment is entered.
func (s *BaseFusionTablesSqlListener) EnterColumn_assignment(ctx *Column_assignmentContext) {}

// ExitColumn_assignment is called when production column_assignment is exited.
func (s *BaseFusionTablesSqlListener) ExitColumn_assignment(ctx *Column_assignmentContext) {}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterEq_comparison is called when production eq_comparison is entered.
func (s *BaseFusionTablesSqlListener) EnterEq_comparison(ctx *Eq_comparisonContext) {}

// ExitEq_comparison is called when production eq_comparison is exited.
func (s *BaseFusionTablesSqlListener) ExitEq_comparison(ctx *Eq_comparisonContext) {}

// EnterTable_name_with_alias is called when production table_name_with_alias is entered.
func (s *BaseFusionTablesSqlListener) EnterTable_name_with_alias(ctx *Table_name_with_aliasContext) {}

// ExitTable_name_with_alias is called when production table_name_with_alias is exited.
func (s *BaseFusionTablesSqlListener) ExitTable_name_with_alias(ctx *Table_name_with_aliasContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseFusionTablesSqlListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseFusionTablesSqlListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterOrdering_term is called when production ordering_term is entered.
func (s *BaseFusionTablesSqlListener) EnterOrdering_term(ctx *Ordering_termContext) {}

// ExitOrdering_term is called when production ordering_term is exited.
func (s *BaseFusionTablesSqlListener) ExitOrdering_term(ctx *Ordering_termContext) {}

// EnterJoin_clause is called when production join_clause is entered.
func (s *BaseFusionTablesSqlListener) EnterJoin_clause(ctx *Join_clauseContext) {}

// ExitJoin_clause is called when production join_clause is exited.
func (s *BaseFusionTablesSqlListener) ExitJoin_clause(ctx *Join_clauseContext) {}

// EnterResult_column is called when production result_column is entered.
func (s *BaseFusionTablesSqlListener) EnterResult_column(ctx *Result_columnContext) {}

// ExitResult_column is called when production result_column is exited.
func (s *BaseFusionTablesSqlListener) ExitResult_column(ctx *Result_columnContext) {}

// EnterQualified_column_name is called when production qualified_column_name is entered.
func (s *BaseFusionTablesSqlListener) EnterQualified_column_name(ctx *Qualified_column_nameContext) {}

// ExitQualified_column_name is called when production qualified_column_name is exited.
func (s *BaseFusionTablesSqlListener) ExitQualified_column_name(ctx *Qualified_column_nameContext) {}

// EnterAggregate_exp is called when production aggregate_exp is entered.
func (s *BaseFusionTablesSqlListener) EnterAggregate_exp(ctx *Aggregate_expContext) {}

// ExitAggregate_exp is called when production aggregate_exp is exited.
func (s *BaseFusionTablesSqlListener) ExitAggregate_exp(ctx *Aggregate_expContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFusionTablesSqlListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFusionTablesSqlListener) ExitExpr(ctx *ExprContext) {}

// EnterColumn_name_beginning_expr is called when production column_name_beginning_expr is entered.
func (s *BaseFusionTablesSqlListener) EnterColumn_name_beginning_expr(ctx *Column_name_beginning_exprContext) {
}

// ExitColumn_name_beginning_expr is called when production column_name_beginning_expr is exited.
func (s *BaseFusionTablesSqlListener) ExitColumn_name_beginning_expr(ctx *Column_name_beginning_exprContext) {
}

// EnterColumn_name_in_dml is called when production column_name_in_dml is entered.
func (s *BaseFusionTablesSqlListener) EnterColumn_name_in_dml(ctx *Column_name_in_dmlContext) {}

// ExitColumn_name_in_dml is called when production column_name_in_dml is exited.
func (s *BaseFusionTablesSqlListener) ExitColumn_name_in_dml(ctx *Column_name_in_dmlContext) {}

// EnterAnd_or_or is called when production and_or_or is entered.
func (s *BaseFusionTablesSqlListener) EnterAnd_or_or(ctx *And_or_orContext) {}

// ExitAnd_or_or is called when production and_or_or is exited.
func (s *BaseFusionTablesSqlListener) ExitAnd_or_or(ctx *And_or_orContext) {}

// EnterGeometry is called when production geometry is entered.
func (s *BaseFusionTablesSqlListener) EnterGeometry(ctx *GeometryContext) {}

// ExitGeometry is called when production geometry is exited.
func (s *BaseFusionTablesSqlListener) ExitGeometry(ctx *GeometryContext) {}

// EnterCircle is called when production circle is entered.
func (s *BaseFusionTablesSqlListener) EnterCircle(ctx *CircleContext) {}

// ExitCircle is called when production circle is exited.
func (s *BaseFusionTablesSqlListener) ExitCircle(ctx *CircleContext) {}

// EnterRectangle is called when production rectangle is entered.
func (s *BaseFusionTablesSqlListener) EnterRectangle(ctx *RectangleContext) {}

// ExitRectangle is called when production rectangle is exited.
func (s *BaseFusionTablesSqlListener) ExitRectangle(ctx *RectangleContext) {}

// EnterCoordinate is called when production coordinate is entered.
func (s *BaseFusionTablesSqlListener) EnterCoordinate(ctx *CoordinateContext) {}

// ExitCoordinate is called when production coordinate is exited.
func (s *BaseFusionTablesSqlListener) ExitCoordinate(ctx *CoordinateContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseFusionTablesSqlListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseFusionTablesSqlListener) ExitKeyword(ctx *KeywordContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseFusionTablesSqlListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseFusionTablesSqlListener) ExitOperator(ctx *OperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseFusionTablesSqlListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseFusionTablesSqlListener) ExitLiteral(ctx *LiteralContext) {}

// EnterError_message is called when production error_message is entered.
func (s *BaseFusionTablesSqlListener) EnterError_message(ctx *Error_messageContext) {}

// ExitError_message is called when production error_message is exited.
func (s *BaseFusionTablesSqlListener) ExitError_message(ctx *Error_messageContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseFusionTablesSqlListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseFusionTablesSqlListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BaseFusionTablesSqlListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BaseFusionTablesSqlListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseFusionTablesSqlListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseFusionTablesSqlListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseFusionTablesSqlListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseFusionTablesSqlListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterNew_table_name is called when production new_table_name is entered.
func (s *BaseFusionTablesSqlListener) EnterNew_table_name(ctx *New_table_nameContext) {}

// ExitNew_table_name is called when production new_table_name is exited.
func (s *BaseFusionTablesSqlListener) ExitNew_table_name(ctx *New_table_nameContext) {}

// EnterView_name is called when production view_name is entered.
func (s *BaseFusionTablesSqlListener) EnterView_name(ctx *View_nameContext) {}

// ExitView_name is called when production view_name is exited.
func (s *BaseFusionTablesSqlListener) ExitView_name(ctx *View_nameContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseFusionTablesSqlListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseFusionTablesSqlListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterNumeric_literal is called when production numeric_literal is entered.
func (s *BaseFusionTablesSqlListener) EnterNumeric_literal(ctx *Numeric_literalContext) {}

// ExitNumeric_literal is called when production numeric_literal is exited.
func (s *BaseFusionTablesSqlListener) ExitNumeric_literal(ctx *Numeric_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseFusionTablesSqlListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseFusionTablesSqlListener) ExitString_literal(ctx *String_literalContext) {}
