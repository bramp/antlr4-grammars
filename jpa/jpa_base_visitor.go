// Generated from JPA.g4 by ANTLR 4.7.

package jpa // JPA
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJPAVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJPAVisitor) VisitQl_statement(ctx *Ql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitIdentification_variable_declaration(ctx *Identification_variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitRange_variable_declaration(ctx *Range_variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitJoin(ctx *JoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitFetch_join(ctx *Fetch_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitJoin_spec(ctx *Join_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitJoin_association_path_expression(ctx *Join_association_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitJoin_collection_valued_path_expression(ctx *Join_collection_valued_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitJoin_single_valued_association_path_expression(ctx *Join_single_valued_association_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitCollection_member_declaration(ctx *Collection_member_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSingle_valued_path_expression(ctx *Single_valued_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitState_field_path_expression(ctx *State_field_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSingle_valued_association_path_expression(ctx *Single_valued_association_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitCollection_valued_path_expression(ctx *Collection_valued_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitState_field(ctx *State_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitUpdate_clause(ctx *Update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitUpdate_item(ctx *Update_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitNew_value(ctx *New_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitDelete_clause(ctx *Delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSelect_expression(ctx *Select_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConstructor_expression(ctx *Constructor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConstructor_item(ctx *Constructor_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitAggregate_expression(ctx *Aggregate_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitGroupby_clause(ctx *Groupby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitGroupby_item(ctx *Groupby_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitOrderby_clause(ctx *Orderby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitOrderby_item(ctx *Orderby_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSubquery_from_clause(ctx *Subquery_from_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSubselect_identification_variable_declaration(ctx *Subselect_identification_variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitAssociation_path_expression(ctx *Association_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_select_clause(ctx *Simple_select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_select_expression(ctx *Simple_select_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConditional_expression(ctx *Conditional_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConditional_term(ctx *Conditional_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConditional_factor(ctx *Conditional_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConditional_primary(ctx *Conditional_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_cond_expression(ctx *Simple_cond_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitBetween_expression(ctx *Between_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitIn_expression(ctx *In_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitIn_item(ctx *In_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitLike_expression(ctx *Like_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitNull_comparison_expression(ctx *Null_comparison_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEmpty_collection_comparison_expression(ctx *Empty_collection_comparison_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitCollection_member_expression(ctx *Collection_member_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitExists_expression(ctx *Exists_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitAll_or_any_expression(ctx *All_or_any_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitComparison_expression(ctx *Comparison_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitComparison_operator(ctx *Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitArithmetic_expression(ctx *Arithmetic_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_arithmetic_expression(ctx *Simple_arithmetic_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitArithmetic_term(ctx *Arithmetic_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitArithmetic_factor(ctx *Arithmetic_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitArithmetic_primary(ctx *Arithmetic_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitString_expression(ctx *String_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitString_primary(ctx *String_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitDatetime_expression(ctx *Datetime_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitDatetime_primary(ctx *Datetime_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitBoolean_expression(ctx *Boolean_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitBoolean_primary(ctx *Boolean_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEnum_expression(ctx *Enum_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEnum_primary(ctx *Enum_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEntity_expression(ctx *Entity_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_entity_expression(ctx *Simple_entity_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitFunctions_returning_numerics(ctx *Functions_returning_numericsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitFunctions_returning_datetime(ctx *Functions_returning_datetimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitFunctions_returning_strings(ctx *Functions_returning_stringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitTrim_specification(ctx *Trim_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitNumeric_literal(ctx *Numeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitPattern_value(ctx *Pattern_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitInput_parameter(ctx *Input_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitConstructor_name(ctx *Constructor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEnum_literal(ctx *Enum_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSimple_state_field(ctx *Simple_state_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitEmbedded_class_state_field(ctx *Embedded_class_state_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitSingle_valued_association_field(ctx *Single_valued_association_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitCollection_valued_association_field(ctx *Collection_valued_association_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJPAVisitor) VisitAbstract_schema_name(ctx *Abstract_schema_nameContext) interface{} {
	return v.VisitChildren(ctx)
}
