// Generated from FusionTablesSql.g4 by ANTLR 4.7.

package fusion-tables // FusionTablesSql
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFusionTablesSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFusionTablesSqlVisitor) VisitFusionTablesSql(ctx *FusionTablesSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitTable_name_in_ddl(ctx *Table_name_in_ddlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitTable_name_in_dml(ctx *Table_name_in_dmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitCreate_table_as_select_stmt(ctx *Create_table_as_select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitDescribe_stmt(ctx *Describe_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitShow_tables_stmt(ctx *Show_tables_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitDrop_table_stmt(ctx *Drop_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitColumn_assignment(ctx *Column_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitEq_comparison(ctx *Eq_comparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitTable_name_with_alias(ctx *Table_name_with_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitOrdering_term(ctx *Ordering_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitJoin_clause(ctx *Join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitResult_column(ctx *Result_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitQualified_column_name(ctx *Qualified_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitAggregate_exp(ctx *Aggregate_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitColumn_name_beginning_expr(ctx *Column_name_beginning_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitColumn_name_in_dml(ctx *Column_name_in_dmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitAnd_or_or(ctx *And_or_orContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitGeometry(ctx *GeometryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitCircle(ctx *CircleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitRectangle(ctx *RectangleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitCoordinate(ctx *CoordinateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitError_message(ctx *Error_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitNew_table_name(ctx *New_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitView_name(ctx *View_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitNumeric_literal(ctx *Numeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFusionTablesSqlVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}
