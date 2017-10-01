// Generated from JPA.g4 by ANTLR 4.7.

package jpa // JPA
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JPAParser.
type JPAVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JPAParser#ql_statement.
	VisitQl_statement(ctx *Ql_statementContext) interface{}

	// Visit a parse tree produced by JPAParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by JPAParser#update_statement.
	VisitUpdate_statement(ctx *Update_statementContext) interface{}

	// Visit a parse tree produced by JPAParser#delete_statement.
	VisitDelete_statement(ctx *Delete_statementContext) interface{}

	// Visit a parse tree produced by JPAParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#identification_variable_declaration.
	VisitIdentification_variable_declaration(ctx *Identification_variable_declarationContext) interface{}

	// Visit a parse tree produced by JPAParser#range_variable_declaration.
	VisitRange_variable_declaration(ctx *Range_variable_declarationContext) interface{}

	// Visit a parse tree produced by JPAParser#join.
	VisitJoin(ctx *JoinContext) interface{}

	// Visit a parse tree produced by JPAParser#fetch_join.
	VisitFetch_join(ctx *Fetch_joinContext) interface{}

	// Visit a parse tree produced by JPAParser#join_spec.
	VisitJoin_spec(ctx *Join_specContext) interface{}

	// Visit a parse tree produced by JPAParser#join_association_path_expression.
	VisitJoin_association_path_expression(ctx *Join_association_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#join_collection_valued_path_expression.
	VisitJoin_collection_valued_path_expression(ctx *Join_collection_valued_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#join_single_valued_association_path_expression.
	VisitJoin_single_valued_association_path_expression(ctx *Join_single_valued_association_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#collection_member_declaration.
	VisitCollection_member_declaration(ctx *Collection_member_declarationContext) interface{}

	// Visit a parse tree produced by JPAParser#single_valued_path_expression.
	VisitSingle_valued_path_expression(ctx *Single_valued_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#state_field_path_expression.
	VisitState_field_path_expression(ctx *State_field_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#single_valued_association_path_expression.
	VisitSingle_valued_association_path_expression(ctx *Single_valued_association_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#collection_valued_path_expression.
	VisitCollection_valued_path_expression(ctx *Collection_valued_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#state_field.
	VisitState_field(ctx *State_fieldContext) interface{}

	// Visit a parse tree produced by JPAParser#update_clause.
	VisitUpdate_clause(ctx *Update_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#update_item.
	VisitUpdate_item(ctx *Update_itemContext) interface{}

	// Visit a parse tree produced by JPAParser#new_value.
	VisitNew_value(ctx *New_valueContext) interface{}

	// Visit a parse tree produced by JPAParser#delete_clause.
	VisitDelete_clause(ctx *Delete_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#select_expression.
	VisitSelect_expression(ctx *Select_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#constructor_expression.
	VisitConstructor_expression(ctx *Constructor_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#constructor_item.
	VisitConstructor_item(ctx *Constructor_itemContext) interface{}

	// Visit a parse tree produced by JPAParser#aggregate_expression.
	VisitAggregate_expression(ctx *Aggregate_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#groupby_clause.
	VisitGroupby_clause(ctx *Groupby_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#groupby_item.
	VisitGroupby_item(ctx *Groupby_itemContext) interface{}

	// Visit a parse tree produced by JPAParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#orderby_clause.
	VisitOrderby_clause(ctx *Orderby_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#orderby_item.
	VisitOrderby_item(ctx *Orderby_itemContext) interface{}

	// Visit a parse tree produced by JPAParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by JPAParser#subquery_from_clause.
	VisitSubquery_from_clause(ctx *Subquery_from_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#subselect_identification_variable_declaration.
	VisitSubselect_identification_variable_declaration(ctx *Subselect_identification_variable_declarationContext) interface{}

	// Visit a parse tree produced by JPAParser#association_path_expression.
	VisitAssociation_path_expression(ctx *Association_path_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_select_clause.
	VisitSimple_select_clause(ctx *Simple_select_clauseContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_select_expression.
	VisitSimple_select_expression(ctx *Simple_select_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#conditional_expression.
	VisitConditional_expression(ctx *Conditional_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#conditional_term.
	VisitConditional_term(ctx *Conditional_termContext) interface{}

	// Visit a parse tree produced by JPAParser#conditional_factor.
	VisitConditional_factor(ctx *Conditional_factorContext) interface{}

	// Visit a parse tree produced by JPAParser#conditional_primary.
	VisitConditional_primary(ctx *Conditional_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_cond_expression.
	VisitSimple_cond_expression(ctx *Simple_cond_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#between_expression.
	VisitBetween_expression(ctx *Between_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#in_expression.
	VisitIn_expression(ctx *In_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#in_item.
	VisitIn_item(ctx *In_itemContext) interface{}

	// Visit a parse tree produced by JPAParser#like_expression.
	VisitLike_expression(ctx *Like_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#null_comparison_expression.
	VisitNull_comparison_expression(ctx *Null_comparison_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#empty_collection_comparison_expression.
	VisitEmpty_collection_comparison_expression(ctx *Empty_collection_comparison_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#collection_member_expression.
	VisitCollection_member_expression(ctx *Collection_member_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#exists_expression.
	VisitExists_expression(ctx *Exists_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#all_or_any_expression.
	VisitAll_or_any_expression(ctx *All_or_any_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#comparison_expression.
	VisitComparison_expression(ctx *Comparison_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#comparison_operator.
	VisitComparison_operator(ctx *Comparison_operatorContext) interface{}

	// Visit a parse tree produced by JPAParser#arithmetic_expression.
	VisitArithmetic_expression(ctx *Arithmetic_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_arithmetic_expression.
	VisitSimple_arithmetic_expression(ctx *Simple_arithmetic_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#arithmetic_term.
	VisitArithmetic_term(ctx *Arithmetic_termContext) interface{}

	// Visit a parse tree produced by JPAParser#arithmetic_factor.
	VisitArithmetic_factor(ctx *Arithmetic_factorContext) interface{}

	// Visit a parse tree produced by JPAParser#arithmetic_primary.
	VisitArithmetic_primary(ctx *Arithmetic_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#string_expression.
	VisitString_expression(ctx *String_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#string_primary.
	VisitString_primary(ctx *String_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#datetime_expression.
	VisitDatetime_expression(ctx *Datetime_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#datetime_primary.
	VisitDatetime_primary(ctx *Datetime_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#boolean_expression.
	VisitBoolean_expression(ctx *Boolean_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#boolean_primary.
	VisitBoolean_primary(ctx *Boolean_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#enum_expression.
	VisitEnum_expression(ctx *Enum_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#enum_primary.
	VisitEnum_primary(ctx *Enum_primaryContext) interface{}

	// Visit a parse tree produced by JPAParser#entity_expression.
	VisitEntity_expression(ctx *Entity_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_entity_expression.
	VisitSimple_entity_expression(ctx *Simple_entity_expressionContext) interface{}

	// Visit a parse tree produced by JPAParser#functions_returning_numerics.
	VisitFunctions_returning_numerics(ctx *Functions_returning_numericsContext) interface{}

	// Visit a parse tree produced by JPAParser#functions_returning_datetime.
	VisitFunctions_returning_datetime(ctx *Functions_returning_datetimeContext) interface{}

	// Visit a parse tree produced by JPAParser#functions_returning_strings.
	VisitFunctions_returning_strings(ctx *Functions_returning_stringsContext) interface{}

	// Visit a parse tree produced by JPAParser#trim_specification.
	VisitTrim_specification(ctx *Trim_specificationContext) interface{}

	// Visit a parse tree produced by JPAParser#numeric_literal.
	VisitNumeric_literal(ctx *Numeric_literalContext) interface{}

	// Visit a parse tree produced by JPAParser#pattern_value.
	VisitPattern_value(ctx *Pattern_valueContext) interface{}

	// Visit a parse tree produced by JPAParser#input_parameter.
	VisitInput_parameter(ctx *Input_parameterContext) interface{}

	// Visit a parse tree produced by JPAParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by JPAParser#constructor_name.
	VisitConstructor_name(ctx *Constructor_nameContext) interface{}

	// Visit a parse tree produced by JPAParser#enum_literal.
	VisitEnum_literal(ctx *Enum_literalContext) interface{}

	// Visit a parse tree produced by JPAParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by JPAParser#simple_state_field.
	VisitSimple_state_field(ctx *Simple_state_fieldContext) interface{}

	// Visit a parse tree produced by JPAParser#embedded_class_state_field.
	VisitEmbedded_class_state_field(ctx *Embedded_class_state_fieldContext) interface{}

	// Visit a parse tree produced by JPAParser#single_valued_association_field.
	VisitSingle_valued_association_field(ctx *Single_valued_association_fieldContext) interface{}

	// Visit a parse tree produced by JPAParser#collection_valued_association_field.
	VisitCollection_valued_association_field(ctx *Collection_valued_association_fieldContext) interface{}

	// Visit a parse tree produced by JPAParser#abstract_schema_name.
	VisitAbstract_schema_name(ctx *Abstract_schema_nameContext) interface{}
}
