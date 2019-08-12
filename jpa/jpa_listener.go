// Code generated from JPA.g4 by ANTLR 4.7.2. DO NOT EDIT.

package jpa // JPA
import "github.com/antlr/antlr4/runtime/Go/antlr"

// JPAListener is a complete listener for a parse tree produced by JPAParser.
type JPAListener interface {
	antlr.ParseTreeListener

	// EnterQl_statement is called when entering the ql_statement production.
	EnterQl_statement(c *Ql_statementContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterIdentification_variable_declaration is called when entering the identification_variable_declaration production.
	EnterIdentification_variable_declaration(c *Identification_variable_declarationContext)

	// EnterRange_variable_declaration is called when entering the range_variable_declaration production.
	EnterRange_variable_declaration(c *Range_variable_declarationContext)

	// EnterJoin is called when entering the join production.
	EnterJoin(c *JoinContext)

	// EnterFetch_join is called when entering the fetch_join production.
	EnterFetch_join(c *Fetch_joinContext)

	// EnterJoin_spec is called when entering the join_spec production.
	EnterJoin_spec(c *Join_specContext)

	// EnterJoin_association_path_expression is called when entering the join_association_path_expression production.
	EnterJoin_association_path_expression(c *Join_association_path_expressionContext)

	// EnterJoin_collection_valued_path_expression is called when entering the join_collection_valued_path_expression production.
	EnterJoin_collection_valued_path_expression(c *Join_collection_valued_path_expressionContext)

	// EnterJoin_single_valued_association_path_expression is called when entering the join_single_valued_association_path_expression production.
	EnterJoin_single_valued_association_path_expression(c *Join_single_valued_association_path_expressionContext)

	// EnterCollection_member_declaration is called when entering the collection_member_declaration production.
	EnterCollection_member_declaration(c *Collection_member_declarationContext)

	// EnterSingle_valued_path_expression is called when entering the single_valued_path_expression production.
	EnterSingle_valued_path_expression(c *Single_valued_path_expressionContext)

	// EnterState_field_path_expression is called when entering the state_field_path_expression production.
	EnterState_field_path_expression(c *State_field_path_expressionContext)

	// EnterSingle_valued_association_path_expression is called when entering the single_valued_association_path_expression production.
	EnterSingle_valued_association_path_expression(c *Single_valued_association_path_expressionContext)

	// EnterCollection_valued_path_expression is called when entering the collection_valued_path_expression production.
	EnterCollection_valued_path_expression(c *Collection_valued_path_expressionContext)

	// EnterState_field is called when entering the state_field production.
	EnterState_field(c *State_fieldContext)

	// EnterUpdate_clause is called when entering the update_clause production.
	EnterUpdate_clause(c *Update_clauseContext)

	// EnterUpdate_item is called when entering the update_item production.
	EnterUpdate_item(c *Update_itemContext)

	// EnterNew_value is called when entering the new_value production.
	EnterNew_value(c *New_valueContext)

	// EnterDelete_clause is called when entering the delete_clause production.
	EnterDelete_clause(c *Delete_clauseContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSelect_expression is called when entering the select_expression production.
	EnterSelect_expression(c *Select_expressionContext)

	// EnterConstructor_expression is called when entering the constructor_expression production.
	EnterConstructor_expression(c *Constructor_expressionContext)

	// EnterConstructor_item is called when entering the constructor_item production.
	EnterConstructor_item(c *Constructor_itemContext)

	// EnterAggregate_expression is called when entering the aggregate_expression production.
	EnterAggregate_expression(c *Aggregate_expressionContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterGroupby_clause is called when entering the groupby_clause production.
	EnterGroupby_clause(c *Groupby_clauseContext)

	// EnterGroupby_item is called when entering the groupby_item production.
	EnterGroupby_item(c *Groupby_itemContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterOrderby_clause is called when entering the orderby_clause production.
	EnterOrderby_clause(c *Orderby_clauseContext)

	// EnterOrderby_item is called when entering the orderby_item production.
	EnterOrderby_item(c *Orderby_itemContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterSubquery_from_clause is called when entering the subquery_from_clause production.
	EnterSubquery_from_clause(c *Subquery_from_clauseContext)

	// EnterSubselect_identification_variable_declaration is called when entering the subselect_identification_variable_declaration production.
	EnterSubselect_identification_variable_declaration(c *Subselect_identification_variable_declarationContext)

	// EnterAssociation_path_expression is called when entering the association_path_expression production.
	EnterAssociation_path_expression(c *Association_path_expressionContext)

	// EnterSimple_select_clause is called when entering the simple_select_clause production.
	EnterSimple_select_clause(c *Simple_select_clauseContext)

	// EnterSimple_select_expression is called when entering the simple_select_expression production.
	EnterSimple_select_expression(c *Simple_select_expressionContext)

	// EnterConditional_expression is called when entering the conditional_expression production.
	EnterConditional_expression(c *Conditional_expressionContext)

	// EnterConditional_term is called when entering the conditional_term production.
	EnterConditional_term(c *Conditional_termContext)

	// EnterConditional_factor is called when entering the conditional_factor production.
	EnterConditional_factor(c *Conditional_factorContext)

	// EnterConditional_primary is called when entering the conditional_primary production.
	EnterConditional_primary(c *Conditional_primaryContext)

	// EnterSimple_cond_expression is called when entering the simple_cond_expression production.
	EnterSimple_cond_expression(c *Simple_cond_expressionContext)

	// EnterBetween_expression is called when entering the between_expression production.
	EnterBetween_expression(c *Between_expressionContext)

	// EnterIn_expression is called when entering the in_expression production.
	EnterIn_expression(c *In_expressionContext)

	// EnterIn_item is called when entering the in_item production.
	EnterIn_item(c *In_itemContext)

	// EnterLike_expression is called when entering the like_expression production.
	EnterLike_expression(c *Like_expressionContext)

	// EnterNull_comparison_expression is called when entering the null_comparison_expression production.
	EnterNull_comparison_expression(c *Null_comparison_expressionContext)

	// EnterEmpty_collection_comparison_expression is called when entering the empty_collection_comparison_expression production.
	EnterEmpty_collection_comparison_expression(c *Empty_collection_comparison_expressionContext)

	// EnterCollection_member_expression is called when entering the collection_member_expression production.
	EnterCollection_member_expression(c *Collection_member_expressionContext)

	// EnterExists_expression is called when entering the exists_expression production.
	EnterExists_expression(c *Exists_expressionContext)

	// EnterAll_or_any_expression is called when entering the all_or_any_expression production.
	EnterAll_or_any_expression(c *All_or_any_expressionContext)

	// EnterComparison_expression is called when entering the comparison_expression production.
	EnterComparison_expression(c *Comparison_expressionContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterArithmetic_expression is called when entering the arithmetic_expression production.
	EnterArithmetic_expression(c *Arithmetic_expressionContext)

	// EnterSimple_arithmetic_expression is called when entering the simple_arithmetic_expression production.
	EnterSimple_arithmetic_expression(c *Simple_arithmetic_expressionContext)

	// EnterArithmetic_term is called when entering the arithmetic_term production.
	EnterArithmetic_term(c *Arithmetic_termContext)

	// EnterArithmetic_factor is called when entering the arithmetic_factor production.
	EnterArithmetic_factor(c *Arithmetic_factorContext)

	// EnterArithmetic_primary is called when entering the arithmetic_primary production.
	EnterArithmetic_primary(c *Arithmetic_primaryContext)

	// EnterString_expression is called when entering the string_expression production.
	EnterString_expression(c *String_expressionContext)

	// EnterString_primary is called when entering the string_primary production.
	EnterString_primary(c *String_primaryContext)

	// EnterDatetime_expression is called when entering the datetime_expression production.
	EnterDatetime_expression(c *Datetime_expressionContext)

	// EnterDatetime_primary is called when entering the datetime_primary production.
	EnterDatetime_primary(c *Datetime_primaryContext)

	// EnterBoolean_expression is called when entering the boolean_expression production.
	EnterBoolean_expression(c *Boolean_expressionContext)

	// EnterBoolean_primary is called when entering the boolean_primary production.
	EnterBoolean_primary(c *Boolean_primaryContext)

	// EnterEnum_expression is called when entering the enum_expression production.
	EnterEnum_expression(c *Enum_expressionContext)

	// EnterEnum_primary is called when entering the enum_primary production.
	EnterEnum_primary(c *Enum_primaryContext)

	// EnterEntity_expression is called when entering the entity_expression production.
	EnterEntity_expression(c *Entity_expressionContext)

	// EnterSimple_entity_expression is called when entering the simple_entity_expression production.
	EnterSimple_entity_expression(c *Simple_entity_expressionContext)

	// EnterFunctions_returning_numerics is called when entering the functions_returning_numerics production.
	EnterFunctions_returning_numerics(c *Functions_returning_numericsContext)

	// EnterFunctions_returning_datetime is called when entering the functions_returning_datetime production.
	EnterFunctions_returning_datetime(c *Functions_returning_datetimeContext)

	// EnterFunctions_returning_strings is called when entering the functions_returning_strings production.
	EnterFunctions_returning_strings(c *Functions_returning_stringsContext)

	// EnterTrim_specification is called when entering the trim_specification production.
	EnterTrim_specification(c *Trim_specificationContext)

	// EnterNumeric_literal is called when entering the numeric_literal production.
	EnterNumeric_literal(c *Numeric_literalContext)

	// EnterPattern_value is called when entering the pattern_value production.
	EnterPattern_value(c *Pattern_valueContext)

	// EnterInput_parameter is called when entering the input_parameter production.
	EnterInput_parameter(c *Input_parameterContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterConstructor_name is called when entering the constructor_name production.
	EnterConstructor_name(c *Constructor_nameContext)

	// EnterEnum_literal is called when entering the enum_literal production.
	EnterEnum_literal(c *Enum_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterSimple_state_field is called when entering the simple_state_field production.
	EnterSimple_state_field(c *Simple_state_fieldContext)

	// EnterEmbedded_class_state_field is called when entering the embedded_class_state_field production.
	EnterEmbedded_class_state_field(c *Embedded_class_state_fieldContext)

	// EnterSingle_valued_association_field is called when entering the single_valued_association_field production.
	EnterSingle_valued_association_field(c *Single_valued_association_fieldContext)

	// EnterCollection_valued_association_field is called when entering the collection_valued_association_field production.
	EnterCollection_valued_association_field(c *Collection_valued_association_fieldContext)

	// EnterAbstract_schema_name is called when entering the abstract_schema_name production.
	EnterAbstract_schema_name(c *Abstract_schema_nameContext)

	// ExitQl_statement is called when exiting the ql_statement production.
	ExitQl_statement(c *Ql_statementContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitIdentification_variable_declaration is called when exiting the identification_variable_declaration production.
	ExitIdentification_variable_declaration(c *Identification_variable_declarationContext)

	// ExitRange_variable_declaration is called when exiting the range_variable_declaration production.
	ExitRange_variable_declaration(c *Range_variable_declarationContext)

	// ExitJoin is called when exiting the join production.
	ExitJoin(c *JoinContext)

	// ExitFetch_join is called when exiting the fetch_join production.
	ExitFetch_join(c *Fetch_joinContext)

	// ExitJoin_spec is called when exiting the join_spec production.
	ExitJoin_spec(c *Join_specContext)

	// ExitJoin_association_path_expression is called when exiting the join_association_path_expression production.
	ExitJoin_association_path_expression(c *Join_association_path_expressionContext)

	// ExitJoin_collection_valued_path_expression is called when exiting the join_collection_valued_path_expression production.
	ExitJoin_collection_valued_path_expression(c *Join_collection_valued_path_expressionContext)

	// ExitJoin_single_valued_association_path_expression is called when exiting the join_single_valued_association_path_expression production.
	ExitJoin_single_valued_association_path_expression(c *Join_single_valued_association_path_expressionContext)

	// ExitCollection_member_declaration is called when exiting the collection_member_declaration production.
	ExitCollection_member_declaration(c *Collection_member_declarationContext)

	// ExitSingle_valued_path_expression is called when exiting the single_valued_path_expression production.
	ExitSingle_valued_path_expression(c *Single_valued_path_expressionContext)

	// ExitState_field_path_expression is called when exiting the state_field_path_expression production.
	ExitState_field_path_expression(c *State_field_path_expressionContext)

	// ExitSingle_valued_association_path_expression is called when exiting the single_valued_association_path_expression production.
	ExitSingle_valued_association_path_expression(c *Single_valued_association_path_expressionContext)

	// ExitCollection_valued_path_expression is called when exiting the collection_valued_path_expression production.
	ExitCollection_valued_path_expression(c *Collection_valued_path_expressionContext)

	// ExitState_field is called when exiting the state_field production.
	ExitState_field(c *State_fieldContext)

	// ExitUpdate_clause is called when exiting the update_clause production.
	ExitUpdate_clause(c *Update_clauseContext)

	// ExitUpdate_item is called when exiting the update_item production.
	ExitUpdate_item(c *Update_itemContext)

	// ExitNew_value is called when exiting the new_value production.
	ExitNew_value(c *New_valueContext)

	// ExitDelete_clause is called when exiting the delete_clause production.
	ExitDelete_clause(c *Delete_clauseContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSelect_expression is called when exiting the select_expression production.
	ExitSelect_expression(c *Select_expressionContext)

	// ExitConstructor_expression is called when exiting the constructor_expression production.
	ExitConstructor_expression(c *Constructor_expressionContext)

	// ExitConstructor_item is called when exiting the constructor_item production.
	ExitConstructor_item(c *Constructor_itemContext)

	// ExitAggregate_expression is called when exiting the aggregate_expression production.
	ExitAggregate_expression(c *Aggregate_expressionContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitGroupby_clause is called when exiting the groupby_clause production.
	ExitGroupby_clause(c *Groupby_clauseContext)

	// ExitGroupby_item is called when exiting the groupby_item production.
	ExitGroupby_item(c *Groupby_itemContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitOrderby_clause is called when exiting the orderby_clause production.
	ExitOrderby_clause(c *Orderby_clauseContext)

	// ExitOrderby_item is called when exiting the orderby_item production.
	ExitOrderby_item(c *Orderby_itemContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitSubquery_from_clause is called when exiting the subquery_from_clause production.
	ExitSubquery_from_clause(c *Subquery_from_clauseContext)

	// ExitSubselect_identification_variable_declaration is called when exiting the subselect_identification_variable_declaration production.
	ExitSubselect_identification_variable_declaration(c *Subselect_identification_variable_declarationContext)

	// ExitAssociation_path_expression is called when exiting the association_path_expression production.
	ExitAssociation_path_expression(c *Association_path_expressionContext)

	// ExitSimple_select_clause is called when exiting the simple_select_clause production.
	ExitSimple_select_clause(c *Simple_select_clauseContext)

	// ExitSimple_select_expression is called when exiting the simple_select_expression production.
	ExitSimple_select_expression(c *Simple_select_expressionContext)

	// ExitConditional_expression is called when exiting the conditional_expression production.
	ExitConditional_expression(c *Conditional_expressionContext)

	// ExitConditional_term is called when exiting the conditional_term production.
	ExitConditional_term(c *Conditional_termContext)

	// ExitConditional_factor is called when exiting the conditional_factor production.
	ExitConditional_factor(c *Conditional_factorContext)

	// ExitConditional_primary is called when exiting the conditional_primary production.
	ExitConditional_primary(c *Conditional_primaryContext)

	// ExitSimple_cond_expression is called when exiting the simple_cond_expression production.
	ExitSimple_cond_expression(c *Simple_cond_expressionContext)

	// ExitBetween_expression is called when exiting the between_expression production.
	ExitBetween_expression(c *Between_expressionContext)

	// ExitIn_expression is called when exiting the in_expression production.
	ExitIn_expression(c *In_expressionContext)

	// ExitIn_item is called when exiting the in_item production.
	ExitIn_item(c *In_itemContext)

	// ExitLike_expression is called when exiting the like_expression production.
	ExitLike_expression(c *Like_expressionContext)

	// ExitNull_comparison_expression is called when exiting the null_comparison_expression production.
	ExitNull_comparison_expression(c *Null_comparison_expressionContext)

	// ExitEmpty_collection_comparison_expression is called when exiting the empty_collection_comparison_expression production.
	ExitEmpty_collection_comparison_expression(c *Empty_collection_comparison_expressionContext)

	// ExitCollection_member_expression is called when exiting the collection_member_expression production.
	ExitCollection_member_expression(c *Collection_member_expressionContext)

	// ExitExists_expression is called when exiting the exists_expression production.
	ExitExists_expression(c *Exists_expressionContext)

	// ExitAll_or_any_expression is called when exiting the all_or_any_expression production.
	ExitAll_or_any_expression(c *All_or_any_expressionContext)

	// ExitComparison_expression is called when exiting the comparison_expression production.
	ExitComparison_expression(c *Comparison_expressionContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitArithmetic_expression is called when exiting the arithmetic_expression production.
	ExitArithmetic_expression(c *Arithmetic_expressionContext)

	// ExitSimple_arithmetic_expression is called when exiting the simple_arithmetic_expression production.
	ExitSimple_arithmetic_expression(c *Simple_arithmetic_expressionContext)

	// ExitArithmetic_term is called when exiting the arithmetic_term production.
	ExitArithmetic_term(c *Arithmetic_termContext)

	// ExitArithmetic_factor is called when exiting the arithmetic_factor production.
	ExitArithmetic_factor(c *Arithmetic_factorContext)

	// ExitArithmetic_primary is called when exiting the arithmetic_primary production.
	ExitArithmetic_primary(c *Arithmetic_primaryContext)

	// ExitString_expression is called when exiting the string_expression production.
	ExitString_expression(c *String_expressionContext)

	// ExitString_primary is called when exiting the string_primary production.
	ExitString_primary(c *String_primaryContext)

	// ExitDatetime_expression is called when exiting the datetime_expression production.
	ExitDatetime_expression(c *Datetime_expressionContext)

	// ExitDatetime_primary is called when exiting the datetime_primary production.
	ExitDatetime_primary(c *Datetime_primaryContext)

	// ExitBoolean_expression is called when exiting the boolean_expression production.
	ExitBoolean_expression(c *Boolean_expressionContext)

	// ExitBoolean_primary is called when exiting the boolean_primary production.
	ExitBoolean_primary(c *Boolean_primaryContext)

	// ExitEnum_expression is called when exiting the enum_expression production.
	ExitEnum_expression(c *Enum_expressionContext)

	// ExitEnum_primary is called when exiting the enum_primary production.
	ExitEnum_primary(c *Enum_primaryContext)

	// ExitEntity_expression is called when exiting the entity_expression production.
	ExitEntity_expression(c *Entity_expressionContext)

	// ExitSimple_entity_expression is called when exiting the simple_entity_expression production.
	ExitSimple_entity_expression(c *Simple_entity_expressionContext)

	// ExitFunctions_returning_numerics is called when exiting the functions_returning_numerics production.
	ExitFunctions_returning_numerics(c *Functions_returning_numericsContext)

	// ExitFunctions_returning_datetime is called when exiting the functions_returning_datetime production.
	ExitFunctions_returning_datetime(c *Functions_returning_datetimeContext)

	// ExitFunctions_returning_strings is called when exiting the functions_returning_strings production.
	ExitFunctions_returning_strings(c *Functions_returning_stringsContext)

	// ExitTrim_specification is called when exiting the trim_specification production.
	ExitTrim_specification(c *Trim_specificationContext)

	// ExitNumeric_literal is called when exiting the numeric_literal production.
	ExitNumeric_literal(c *Numeric_literalContext)

	// ExitPattern_value is called when exiting the pattern_value production.
	ExitPattern_value(c *Pattern_valueContext)

	// ExitInput_parameter is called when exiting the input_parameter production.
	ExitInput_parameter(c *Input_parameterContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitConstructor_name is called when exiting the constructor_name production.
	ExitConstructor_name(c *Constructor_nameContext)

	// ExitEnum_literal is called when exiting the enum_literal production.
	ExitEnum_literal(c *Enum_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitSimple_state_field is called when exiting the simple_state_field production.
	ExitSimple_state_field(c *Simple_state_fieldContext)

	// ExitEmbedded_class_state_field is called when exiting the embedded_class_state_field production.
	ExitEmbedded_class_state_field(c *Embedded_class_state_fieldContext)

	// ExitSingle_valued_association_field is called when exiting the single_valued_association_field production.
	ExitSingle_valued_association_field(c *Single_valued_association_fieldContext)

	// ExitCollection_valued_association_field is called when exiting the collection_valued_association_field production.
	ExitCollection_valued_association_field(c *Collection_valued_association_fieldContext)

	// ExitAbstract_schema_name is called when exiting the abstract_schema_name production.
	ExitAbstract_schema_name(c *Abstract_schema_nameContext)
}
