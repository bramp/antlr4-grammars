// Code generated from FusionTablesSql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package fusiontablessql // FusionTablesSql
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FusionTablesSqlListener is a complete listener for a parse tree produced by FusionTablesSqlParser.
type FusionTablesSqlListener interface {
	antlr.ParseTreeListener

	// EnterFusionTablesSql is called when entering the fusionTablesSql production.
	EnterFusionTablesSql(c *FusionTablesSqlContext)

	// EnterSql_stmt is called when entering the sql_stmt production.
	EnterSql_stmt(c *Sql_stmtContext)

	// EnterTable_name_in_ddl is called when entering the table_name_in_ddl production.
	EnterTable_name_in_ddl(c *Table_name_in_ddlContext)

	// EnterTable_name_in_dml is called when entering the table_name_in_dml production.
	EnterTable_name_in_dml(c *Table_name_in_dmlContext)

	// EnterCreate_table_as_select_stmt is called when entering the create_table_as_select_stmt production.
	EnterCreate_table_as_select_stmt(c *Create_table_as_select_stmtContext)

	// EnterDescribe_stmt is called when entering the describe_stmt production.
	EnterDescribe_stmt(c *Describe_stmtContext)

	// EnterShow_tables_stmt is called when entering the show_tables_stmt production.
	EnterShow_tables_stmt(c *Show_tables_stmtContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterCreate_view_stmt is called when entering the create_view_stmt production.
	EnterCreate_view_stmt(c *Create_view_stmtContext)

	// EnterDrop_table_stmt is called when entering the drop_table_stmt production.
	EnterDrop_table_stmt(c *Drop_table_stmtContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterColumn_assignment is called when entering the column_assignment production.
	EnterColumn_assignment(c *Column_assignmentContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterEq_comparison is called when entering the eq_comparison production.
	EnterEq_comparison(c *Eq_comparisonContext)

	// EnterTable_name_with_alias is called when entering the table_name_with_alias production.
	EnterTable_name_with_alias(c *Table_name_with_aliasContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterOrdering_term is called when entering the ordering_term production.
	EnterOrdering_term(c *Ordering_termContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterResult_column is called when entering the result_column production.
	EnterResult_column(c *Result_columnContext)

	// EnterQualified_column_name is called when entering the qualified_column_name production.
	EnterQualified_column_name(c *Qualified_column_nameContext)

	// EnterAggregate_exp is called when entering the aggregate_exp production.
	EnterAggregate_exp(c *Aggregate_expContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterColumn_name_beginning_expr is called when entering the column_name_beginning_expr production.
	EnterColumn_name_beginning_expr(c *Column_name_beginning_exprContext)

	// EnterColumn_name_in_dml is called when entering the column_name_in_dml production.
	EnterColumn_name_in_dml(c *Column_name_in_dmlContext)

	// EnterAnd_or_or is called when entering the and_or_or production.
	EnterAnd_or_or(c *And_or_orContext)

	// EnterGeometry is called when entering the geometry production.
	EnterGeometry(c *GeometryContext)

	// EnterCircle is called when entering the circle production.
	EnterCircle(c *CircleContext)

	// EnterRectangle is called when entering the rectangle production.
	EnterRectangle(c *RectangleContext)

	// EnterCoordinate is called when entering the coordinate production.
	EnterCoordinate(c *CoordinateContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterError_message is called when entering the error_message production.
	EnterError_message(c *Error_messageContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterNew_table_name is called when entering the new_table_name production.
	EnterNew_table_name(c *New_table_nameContext)

	// EnterView_name is called when entering the view_name production.
	EnterView_name(c *View_nameContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterNumeric_literal is called when entering the numeric_literal production.
	EnterNumeric_literal(c *Numeric_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// ExitFusionTablesSql is called when exiting the fusionTablesSql production.
	ExitFusionTablesSql(c *FusionTablesSqlContext)

	// ExitSql_stmt is called when exiting the sql_stmt production.
	ExitSql_stmt(c *Sql_stmtContext)

	// ExitTable_name_in_ddl is called when exiting the table_name_in_ddl production.
	ExitTable_name_in_ddl(c *Table_name_in_ddlContext)

	// ExitTable_name_in_dml is called when exiting the table_name_in_dml production.
	ExitTable_name_in_dml(c *Table_name_in_dmlContext)

	// ExitCreate_table_as_select_stmt is called when exiting the create_table_as_select_stmt production.
	ExitCreate_table_as_select_stmt(c *Create_table_as_select_stmtContext)

	// ExitDescribe_stmt is called when exiting the describe_stmt production.
	ExitDescribe_stmt(c *Describe_stmtContext)

	// ExitShow_tables_stmt is called when exiting the show_tables_stmt production.
	ExitShow_tables_stmt(c *Show_tables_stmtContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitCreate_view_stmt is called when exiting the create_view_stmt production.
	ExitCreate_view_stmt(c *Create_view_stmtContext)

	// ExitDrop_table_stmt is called when exiting the drop_table_stmt production.
	ExitDrop_table_stmt(c *Drop_table_stmtContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitColumn_assignment is called when exiting the column_assignment production.
	ExitColumn_assignment(c *Column_assignmentContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitEq_comparison is called when exiting the eq_comparison production.
	ExitEq_comparison(c *Eq_comparisonContext)

	// ExitTable_name_with_alias is called when exiting the table_name_with_alias production.
	ExitTable_name_with_alias(c *Table_name_with_aliasContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitOrdering_term is called when exiting the ordering_term production.
	ExitOrdering_term(c *Ordering_termContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitResult_column is called when exiting the result_column production.
	ExitResult_column(c *Result_columnContext)

	// ExitQualified_column_name is called when exiting the qualified_column_name production.
	ExitQualified_column_name(c *Qualified_column_nameContext)

	// ExitAggregate_exp is called when exiting the aggregate_exp production.
	ExitAggregate_exp(c *Aggregate_expContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitColumn_name_beginning_expr is called when exiting the column_name_beginning_expr production.
	ExitColumn_name_beginning_expr(c *Column_name_beginning_exprContext)

	// ExitColumn_name_in_dml is called when exiting the column_name_in_dml production.
	ExitColumn_name_in_dml(c *Column_name_in_dmlContext)

	// ExitAnd_or_or is called when exiting the and_or_or production.
	ExitAnd_or_or(c *And_or_orContext)

	// ExitGeometry is called when exiting the geometry production.
	ExitGeometry(c *GeometryContext)

	// ExitCircle is called when exiting the circle production.
	ExitCircle(c *CircleContext)

	// ExitRectangle is called when exiting the rectangle production.
	ExitRectangle(c *RectangleContext)

	// ExitCoordinate is called when exiting the coordinate production.
	ExitCoordinate(c *CoordinateContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitError_message is called when exiting the error_message production.
	ExitError_message(c *Error_messageContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitNew_table_name is called when exiting the new_table_name production.
	ExitNew_table_name(c *New_table_nameContext)

	// ExitView_name is called when exiting the view_name production.
	ExitView_name(c *View_nameContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitNumeric_literal is called when exiting the numeric_literal production.
	ExitNumeric_literal(c *Numeric_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)
}
