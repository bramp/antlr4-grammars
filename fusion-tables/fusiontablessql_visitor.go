// Generated from FusionTablesSql.g4 by ANTLR 4.7.

package fusion-tables // FusionTablesSql
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by FusionTablesSqlParser.
type FusionTablesSqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FusionTablesSqlParser#fusionTablesSql.
	VisitFusionTablesSql(ctx *FusionTablesSqlContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#table_name_in_ddl.
	VisitTable_name_in_ddl(ctx *Table_name_in_ddlContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#table_name_in_dml.
	VisitTable_name_in_dml(ctx *Table_name_in_dmlContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#create_table_as_select_stmt.
	VisitCreate_table_as_select_stmt(ctx *Create_table_as_select_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#describe_stmt.
	VisitDescribe_stmt(ctx *Describe_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#show_tables_stmt.
	VisitShow_tables_stmt(ctx *Show_tables_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#alter_table_stmt.
	VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#create_view_stmt.
	VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#drop_table_stmt.
	VisitDrop_table_stmt(ctx *Drop_table_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#insert_stmt.
	VisitInsert_stmt(ctx *Insert_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#update_stmt.
	VisitUpdate_stmt(ctx *Update_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#column_assignment.
	VisitColumn_assignment(ctx *Column_assignmentContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#delete_stmt.
	VisitDelete_stmt(ctx *Delete_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#eq_comparison.
	VisitEq_comparison(ctx *Eq_comparisonContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#table_name_with_alias.
	VisitTable_name_with_alias(ctx *Table_name_with_aliasContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#select_stmt.
	VisitSelect_stmt(ctx *Select_stmtContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#ordering_term.
	VisitOrdering_term(ctx *Ordering_termContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#join_clause.
	VisitJoin_clause(ctx *Join_clauseContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#result_column.
	VisitResult_column(ctx *Result_columnContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#qualified_column_name.
	VisitQualified_column_name(ctx *Qualified_column_nameContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#aggregate_exp.
	VisitAggregate_exp(ctx *Aggregate_expContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#column_name_beginning_expr.
	VisitColumn_name_beginning_expr(ctx *Column_name_beginning_exprContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#column_name_in_dml.
	VisitColumn_name_in_dml(ctx *Column_name_in_dmlContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#and_or_or.
	VisitAnd_or_or(ctx *And_or_orContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#geometry.
	VisitGeometry(ctx *GeometryContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#circle.
	VisitCircle(ctx *CircleContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#rectangle.
	VisitRectangle(ctx *RectangleContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#coordinate.
	VisitCoordinate(ctx *CoordinateContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#error_message.
	VisitError_message(ctx *Error_messageContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#new_table_name.
	VisitNew_table_name(ctx *New_table_nameContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#view_name.
	VisitView_name(ctx *View_nameContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#numeric_literal.
	VisitNumeric_literal(ctx *Numeric_literalContext) interface{}

	// Visit a parse tree produced by FusionTablesSqlParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

}