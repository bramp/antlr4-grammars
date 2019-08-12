// Code generated from JPA.g4 by ANTLR 4.7.2. DO NOT EDIT.

package jpa // JPA
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJPAListener is a complete listener for a parse tree produced by JPAParser.
type BaseJPAListener struct{}

var _ JPAListener = &BaseJPAListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJPAListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJPAListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJPAListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJPAListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQl_statement is called when production ql_statement is entered.
func (s *BaseJPAListener) EnterQl_statement(ctx *Ql_statementContext) {}

// ExitQl_statement is called when production ql_statement is exited.
func (s *BaseJPAListener) ExitQl_statement(ctx *Ql_statementContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BaseJPAListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BaseJPAListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BaseJPAListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BaseJPAListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BaseJPAListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BaseJPAListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseJPAListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseJPAListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterIdentification_variable_declaration is called when production identification_variable_declaration is entered.
func (s *BaseJPAListener) EnterIdentification_variable_declaration(ctx *Identification_variable_declarationContext) {
}

// ExitIdentification_variable_declaration is called when production identification_variable_declaration is exited.
func (s *BaseJPAListener) ExitIdentification_variable_declaration(ctx *Identification_variable_declarationContext) {
}

// EnterRange_variable_declaration is called when production range_variable_declaration is entered.
func (s *BaseJPAListener) EnterRange_variable_declaration(ctx *Range_variable_declarationContext) {}

// ExitRange_variable_declaration is called when production range_variable_declaration is exited.
func (s *BaseJPAListener) ExitRange_variable_declaration(ctx *Range_variable_declarationContext) {}

// EnterJoin is called when production join is entered.
func (s *BaseJPAListener) EnterJoin(ctx *JoinContext) {}

// ExitJoin is called when production join is exited.
func (s *BaseJPAListener) ExitJoin(ctx *JoinContext) {}

// EnterFetch_join is called when production fetch_join is entered.
func (s *BaseJPAListener) EnterFetch_join(ctx *Fetch_joinContext) {}

// ExitFetch_join is called when production fetch_join is exited.
func (s *BaseJPAListener) ExitFetch_join(ctx *Fetch_joinContext) {}

// EnterJoin_spec is called when production join_spec is entered.
func (s *BaseJPAListener) EnterJoin_spec(ctx *Join_specContext) {}

// ExitJoin_spec is called when production join_spec is exited.
func (s *BaseJPAListener) ExitJoin_spec(ctx *Join_specContext) {}

// EnterJoin_association_path_expression is called when production join_association_path_expression is entered.
func (s *BaseJPAListener) EnterJoin_association_path_expression(ctx *Join_association_path_expressionContext) {
}

// ExitJoin_association_path_expression is called when production join_association_path_expression is exited.
func (s *BaseJPAListener) ExitJoin_association_path_expression(ctx *Join_association_path_expressionContext) {
}

// EnterJoin_collection_valued_path_expression is called when production join_collection_valued_path_expression is entered.
func (s *BaseJPAListener) EnterJoin_collection_valued_path_expression(ctx *Join_collection_valued_path_expressionContext) {
}

// ExitJoin_collection_valued_path_expression is called when production join_collection_valued_path_expression is exited.
func (s *BaseJPAListener) ExitJoin_collection_valued_path_expression(ctx *Join_collection_valued_path_expressionContext) {
}

// EnterJoin_single_valued_association_path_expression is called when production join_single_valued_association_path_expression is entered.
func (s *BaseJPAListener) EnterJoin_single_valued_association_path_expression(ctx *Join_single_valued_association_path_expressionContext) {
}

// ExitJoin_single_valued_association_path_expression is called when production join_single_valued_association_path_expression is exited.
func (s *BaseJPAListener) ExitJoin_single_valued_association_path_expression(ctx *Join_single_valued_association_path_expressionContext) {
}

// EnterCollection_member_declaration is called when production collection_member_declaration is entered.
func (s *BaseJPAListener) EnterCollection_member_declaration(ctx *Collection_member_declarationContext) {
}

// ExitCollection_member_declaration is called when production collection_member_declaration is exited.
func (s *BaseJPAListener) ExitCollection_member_declaration(ctx *Collection_member_declarationContext) {
}

// EnterSingle_valued_path_expression is called when production single_valued_path_expression is entered.
func (s *BaseJPAListener) EnterSingle_valued_path_expression(ctx *Single_valued_path_expressionContext) {
}

// ExitSingle_valued_path_expression is called when production single_valued_path_expression is exited.
func (s *BaseJPAListener) ExitSingle_valued_path_expression(ctx *Single_valued_path_expressionContext) {
}

// EnterState_field_path_expression is called when production state_field_path_expression is entered.
func (s *BaseJPAListener) EnterState_field_path_expression(ctx *State_field_path_expressionContext) {}

// ExitState_field_path_expression is called when production state_field_path_expression is exited.
func (s *BaseJPAListener) ExitState_field_path_expression(ctx *State_field_path_expressionContext) {}

// EnterSingle_valued_association_path_expression is called when production single_valued_association_path_expression is entered.
func (s *BaseJPAListener) EnterSingle_valued_association_path_expression(ctx *Single_valued_association_path_expressionContext) {
}

// ExitSingle_valued_association_path_expression is called when production single_valued_association_path_expression is exited.
func (s *BaseJPAListener) ExitSingle_valued_association_path_expression(ctx *Single_valued_association_path_expressionContext) {
}

// EnterCollection_valued_path_expression is called when production collection_valued_path_expression is entered.
func (s *BaseJPAListener) EnterCollection_valued_path_expression(ctx *Collection_valued_path_expressionContext) {
}

// ExitCollection_valued_path_expression is called when production collection_valued_path_expression is exited.
func (s *BaseJPAListener) ExitCollection_valued_path_expression(ctx *Collection_valued_path_expressionContext) {
}

// EnterState_field is called when production state_field is entered.
func (s *BaseJPAListener) EnterState_field(ctx *State_fieldContext) {}

// ExitState_field is called when production state_field is exited.
func (s *BaseJPAListener) ExitState_field(ctx *State_fieldContext) {}

// EnterUpdate_clause is called when production update_clause is entered.
func (s *BaseJPAListener) EnterUpdate_clause(ctx *Update_clauseContext) {}

// ExitUpdate_clause is called when production update_clause is exited.
func (s *BaseJPAListener) ExitUpdate_clause(ctx *Update_clauseContext) {}

// EnterUpdate_item is called when production update_item is entered.
func (s *BaseJPAListener) EnterUpdate_item(ctx *Update_itemContext) {}

// ExitUpdate_item is called when production update_item is exited.
func (s *BaseJPAListener) ExitUpdate_item(ctx *Update_itemContext) {}

// EnterNew_value is called when production new_value is entered.
func (s *BaseJPAListener) EnterNew_value(ctx *New_valueContext) {}

// ExitNew_value is called when production new_value is exited.
func (s *BaseJPAListener) ExitNew_value(ctx *New_valueContext) {}

// EnterDelete_clause is called when production delete_clause is entered.
func (s *BaseJPAListener) EnterDelete_clause(ctx *Delete_clauseContext) {}

// ExitDelete_clause is called when production delete_clause is exited.
func (s *BaseJPAListener) ExitDelete_clause(ctx *Delete_clauseContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseJPAListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseJPAListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSelect_expression is called when production select_expression is entered.
func (s *BaseJPAListener) EnterSelect_expression(ctx *Select_expressionContext) {}

// ExitSelect_expression is called when production select_expression is exited.
func (s *BaseJPAListener) ExitSelect_expression(ctx *Select_expressionContext) {}

// EnterConstructor_expression is called when production constructor_expression is entered.
func (s *BaseJPAListener) EnterConstructor_expression(ctx *Constructor_expressionContext) {}

// ExitConstructor_expression is called when production constructor_expression is exited.
func (s *BaseJPAListener) ExitConstructor_expression(ctx *Constructor_expressionContext) {}

// EnterConstructor_item is called when production constructor_item is entered.
func (s *BaseJPAListener) EnterConstructor_item(ctx *Constructor_itemContext) {}

// ExitConstructor_item is called when production constructor_item is exited.
func (s *BaseJPAListener) ExitConstructor_item(ctx *Constructor_itemContext) {}

// EnterAggregate_expression is called when production aggregate_expression is entered.
func (s *BaseJPAListener) EnterAggregate_expression(ctx *Aggregate_expressionContext) {}

// ExitAggregate_expression is called when production aggregate_expression is exited.
func (s *BaseJPAListener) ExitAggregate_expression(ctx *Aggregate_expressionContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseJPAListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseJPAListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterGroupby_clause is called when production groupby_clause is entered.
func (s *BaseJPAListener) EnterGroupby_clause(ctx *Groupby_clauseContext) {}

// ExitGroupby_clause is called when production groupby_clause is exited.
func (s *BaseJPAListener) ExitGroupby_clause(ctx *Groupby_clauseContext) {}

// EnterGroupby_item is called when production groupby_item is entered.
func (s *BaseJPAListener) EnterGroupby_item(ctx *Groupby_itemContext) {}

// ExitGroupby_item is called when production groupby_item is exited.
func (s *BaseJPAListener) ExitGroupby_item(ctx *Groupby_itemContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseJPAListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseJPAListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterOrderby_clause is called when production orderby_clause is entered.
func (s *BaseJPAListener) EnterOrderby_clause(ctx *Orderby_clauseContext) {}

// ExitOrderby_clause is called when production orderby_clause is exited.
func (s *BaseJPAListener) ExitOrderby_clause(ctx *Orderby_clauseContext) {}

// EnterOrderby_item is called when production orderby_item is entered.
func (s *BaseJPAListener) EnterOrderby_item(ctx *Orderby_itemContext) {}

// ExitOrderby_item is called when production orderby_item is exited.
func (s *BaseJPAListener) ExitOrderby_item(ctx *Orderby_itemContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseJPAListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseJPAListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSubquery_from_clause is called when production subquery_from_clause is entered.
func (s *BaseJPAListener) EnterSubquery_from_clause(ctx *Subquery_from_clauseContext) {}

// ExitSubquery_from_clause is called when production subquery_from_clause is exited.
func (s *BaseJPAListener) ExitSubquery_from_clause(ctx *Subquery_from_clauseContext) {}

// EnterSubselect_identification_variable_declaration is called when production subselect_identification_variable_declaration is entered.
func (s *BaseJPAListener) EnterSubselect_identification_variable_declaration(ctx *Subselect_identification_variable_declarationContext) {
}

// ExitSubselect_identification_variable_declaration is called when production subselect_identification_variable_declaration is exited.
func (s *BaseJPAListener) ExitSubselect_identification_variable_declaration(ctx *Subselect_identification_variable_declarationContext) {
}

// EnterAssociation_path_expression is called when production association_path_expression is entered.
func (s *BaseJPAListener) EnterAssociation_path_expression(ctx *Association_path_expressionContext) {}

// ExitAssociation_path_expression is called when production association_path_expression is exited.
func (s *BaseJPAListener) ExitAssociation_path_expression(ctx *Association_path_expressionContext) {}

// EnterSimple_select_clause is called when production simple_select_clause is entered.
func (s *BaseJPAListener) EnterSimple_select_clause(ctx *Simple_select_clauseContext) {}

// ExitSimple_select_clause is called when production simple_select_clause is exited.
func (s *BaseJPAListener) ExitSimple_select_clause(ctx *Simple_select_clauseContext) {}

// EnterSimple_select_expression is called when production simple_select_expression is entered.
func (s *BaseJPAListener) EnterSimple_select_expression(ctx *Simple_select_expressionContext) {}

// ExitSimple_select_expression is called when production simple_select_expression is exited.
func (s *BaseJPAListener) ExitSimple_select_expression(ctx *Simple_select_expressionContext) {}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseJPAListener) EnterConditional_expression(ctx *Conditional_expressionContext) {}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseJPAListener) ExitConditional_expression(ctx *Conditional_expressionContext) {}

// EnterConditional_term is called when production conditional_term is entered.
func (s *BaseJPAListener) EnterConditional_term(ctx *Conditional_termContext) {}

// ExitConditional_term is called when production conditional_term is exited.
func (s *BaseJPAListener) ExitConditional_term(ctx *Conditional_termContext) {}

// EnterConditional_factor is called when production conditional_factor is entered.
func (s *BaseJPAListener) EnterConditional_factor(ctx *Conditional_factorContext) {}

// ExitConditional_factor is called when production conditional_factor is exited.
func (s *BaseJPAListener) ExitConditional_factor(ctx *Conditional_factorContext) {}

// EnterConditional_primary is called when production conditional_primary is entered.
func (s *BaseJPAListener) EnterConditional_primary(ctx *Conditional_primaryContext) {}

// ExitConditional_primary is called when production conditional_primary is exited.
func (s *BaseJPAListener) ExitConditional_primary(ctx *Conditional_primaryContext) {}

// EnterSimple_cond_expression is called when production simple_cond_expression is entered.
func (s *BaseJPAListener) EnterSimple_cond_expression(ctx *Simple_cond_expressionContext) {}

// ExitSimple_cond_expression is called when production simple_cond_expression is exited.
func (s *BaseJPAListener) ExitSimple_cond_expression(ctx *Simple_cond_expressionContext) {}

// EnterBetween_expression is called when production between_expression is entered.
func (s *BaseJPAListener) EnterBetween_expression(ctx *Between_expressionContext) {}

// ExitBetween_expression is called when production between_expression is exited.
func (s *BaseJPAListener) ExitBetween_expression(ctx *Between_expressionContext) {}

// EnterIn_expression is called when production in_expression is entered.
func (s *BaseJPAListener) EnterIn_expression(ctx *In_expressionContext) {}

// ExitIn_expression is called when production in_expression is exited.
func (s *BaseJPAListener) ExitIn_expression(ctx *In_expressionContext) {}

// EnterIn_item is called when production in_item is entered.
func (s *BaseJPAListener) EnterIn_item(ctx *In_itemContext) {}

// ExitIn_item is called when production in_item is exited.
func (s *BaseJPAListener) ExitIn_item(ctx *In_itemContext) {}

// EnterLike_expression is called when production like_expression is entered.
func (s *BaseJPAListener) EnterLike_expression(ctx *Like_expressionContext) {}

// ExitLike_expression is called when production like_expression is exited.
func (s *BaseJPAListener) ExitLike_expression(ctx *Like_expressionContext) {}

// EnterNull_comparison_expression is called when production null_comparison_expression is entered.
func (s *BaseJPAListener) EnterNull_comparison_expression(ctx *Null_comparison_expressionContext) {}

// ExitNull_comparison_expression is called when production null_comparison_expression is exited.
func (s *BaseJPAListener) ExitNull_comparison_expression(ctx *Null_comparison_expressionContext) {}

// EnterEmpty_collection_comparison_expression is called when production empty_collection_comparison_expression is entered.
func (s *BaseJPAListener) EnterEmpty_collection_comparison_expression(ctx *Empty_collection_comparison_expressionContext) {
}

// ExitEmpty_collection_comparison_expression is called when production empty_collection_comparison_expression is exited.
func (s *BaseJPAListener) ExitEmpty_collection_comparison_expression(ctx *Empty_collection_comparison_expressionContext) {
}

// EnterCollection_member_expression is called when production collection_member_expression is entered.
func (s *BaseJPAListener) EnterCollection_member_expression(ctx *Collection_member_expressionContext) {
}

// ExitCollection_member_expression is called when production collection_member_expression is exited.
func (s *BaseJPAListener) ExitCollection_member_expression(ctx *Collection_member_expressionContext) {}

// EnterExists_expression is called when production exists_expression is entered.
func (s *BaseJPAListener) EnterExists_expression(ctx *Exists_expressionContext) {}

// ExitExists_expression is called when production exists_expression is exited.
func (s *BaseJPAListener) ExitExists_expression(ctx *Exists_expressionContext) {}

// EnterAll_or_any_expression is called when production all_or_any_expression is entered.
func (s *BaseJPAListener) EnterAll_or_any_expression(ctx *All_or_any_expressionContext) {}

// ExitAll_or_any_expression is called when production all_or_any_expression is exited.
func (s *BaseJPAListener) ExitAll_or_any_expression(ctx *All_or_any_expressionContext) {}

// EnterComparison_expression is called when production comparison_expression is entered.
func (s *BaseJPAListener) EnterComparison_expression(ctx *Comparison_expressionContext) {}

// ExitComparison_expression is called when production comparison_expression is exited.
func (s *BaseJPAListener) ExitComparison_expression(ctx *Comparison_expressionContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseJPAListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseJPAListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterArithmetic_expression is called when production arithmetic_expression is entered.
func (s *BaseJPAListener) EnterArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// ExitArithmetic_expression is called when production arithmetic_expression is exited.
func (s *BaseJPAListener) ExitArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// EnterSimple_arithmetic_expression is called when production simple_arithmetic_expression is entered.
func (s *BaseJPAListener) EnterSimple_arithmetic_expression(ctx *Simple_arithmetic_expressionContext) {
}

// ExitSimple_arithmetic_expression is called when production simple_arithmetic_expression is exited.
func (s *BaseJPAListener) ExitSimple_arithmetic_expression(ctx *Simple_arithmetic_expressionContext) {}

// EnterArithmetic_term is called when production arithmetic_term is entered.
func (s *BaseJPAListener) EnterArithmetic_term(ctx *Arithmetic_termContext) {}

// ExitArithmetic_term is called when production arithmetic_term is exited.
func (s *BaseJPAListener) ExitArithmetic_term(ctx *Arithmetic_termContext) {}

// EnterArithmetic_factor is called when production arithmetic_factor is entered.
func (s *BaseJPAListener) EnterArithmetic_factor(ctx *Arithmetic_factorContext) {}

// ExitArithmetic_factor is called when production arithmetic_factor is exited.
func (s *BaseJPAListener) ExitArithmetic_factor(ctx *Arithmetic_factorContext) {}

// EnterArithmetic_primary is called when production arithmetic_primary is entered.
func (s *BaseJPAListener) EnterArithmetic_primary(ctx *Arithmetic_primaryContext) {}

// ExitArithmetic_primary is called when production arithmetic_primary is exited.
func (s *BaseJPAListener) ExitArithmetic_primary(ctx *Arithmetic_primaryContext) {}

// EnterString_expression is called when production string_expression is entered.
func (s *BaseJPAListener) EnterString_expression(ctx *String_expressionContext) {}

// ExitString_expression is called when production string_expression is exited.
func (s *BaseJPAListener) ExitString_expression(ctx *String_expressionContext) {}

// EnterString_primary is called when production string_primary is entered.
func (s *BaseJPAListener) EnterString_primary(ctx *String_primaryContext) {}

// ExitString_primary is called when production string_primary is exited.
func (s *BaseJPAListener) ExitString_primary(ctx *String_primaryContext) {}

// EnterDatetime_expression is called when production datetime_expression is entered.
func (s *BaseJPAListener) EnterDatetime_expression(ctx *Datetime_expressionContext) {}

// ExitDatetime_expression is called when production datetime_expression is exited.
func (s *BaseJPAListener) ExitDatetime_expression(ctx *Datetime_expressionContext) {}

// EnterDatetime_primary is called when production datetime_primary is entered.
func (s *BaseJPAListener) EnterDatetime_primary(ctx *Datetime_primaryContext) {}

// ExitDatetime_primary is called when production datetime_primary is exited.
func (s *BaseJPAListener) ExitDatetime_primary(ctx *Datetime_primaryContext) {}

// EnterBoolean_expression is called when production boolean_expression is entered.
func (s *BaseJPAListener) EnterBoolean_expression(ctx *Boolean_expressionContext) {}

// ExitBoolean_expression is called when production boolean_expression is exited.
func (s *BaseJPAListener) ExitBoolean_expression(ctx *Boolean_expressionContext) {}

// EnterBoolean_primary is called when production boolean_primary is entered.
func (s *BaseJPAListener) EnterBoolean_primary(ctx *Boolean_primaryContext) {}

// ExitBoolean_primary is called when production boolean_primary is exited.
func (s *BaseJPAListener) ExitBoolean_primary(ctx *Boolean_primaryContext) {}

// EnterEnum_expression is called when production enum_expression is entered.
func (s *BaseJPAListener) EnterEnum_expression(ctx *Enum_expressionContext) {}

// ExitEnum_expression is called when production enum_expression is exited.
func (s *BaseJPAListener) ExitEnum_expression(ctx *Enum_expressionContext) {}

// EnterEnum_primary is called when production enum_primary is entered.
func (s *BaseJPAListener) EnterEnum_primary(ctx *Enum_primaryContext) {}

// ExitEnum_primary is called when production enum_primary is exited.
func (s *BaseJPAListener) ExitEnum_primary(ctx *Enum_primaryContext) {}

// EnterEntity_expression is called when production entity_expression is entered.
func (s *BaseJPAListener) EnterEntity_expression(ctx *Entity_expressionContext) {}

// ExitEntity_expression is called when production entity_expression is exited.
func (s *BaseJPAListener) ExitEntity_expression(ctx *Entity_expressionContext) {}

// EnterSimple_entity_expression is called when production simple_entity_expression is entered.
func (s *BaseJPAListener) EnterSimple_entity_expression(ctx *Simple_entity_expressionContext) {}

// ExitSimple_entity_expression is called when production simple_entity_expression is exited.
func (s *BaseJPAListener) ExitSimple_entity_expression(ctx *Simple_entity_expressionContext) {}

// EnterFunctions_returning_numerics is called when production functions_returning_numerics is entered.
func (s *BaseJPAListener) EnterFunctions_returning_numerics(ctx *Functions_returning_numericsContext) {
}

// ExitFunctions_returning_numerics is called when production functions_returning_numerics is exited.
func (s *BaseJPAListener) ExitFunctions_returning_numerics(ctx *Functions_returning_numericsContext) {}

// EnterFunctions_returning_datetime is called when production functions_returning_datetime is entered.
func (s *BaseJPAListener) EnterFunctions_returning_datetime(ctx *Functions_returning_datetimeContext) {
}

// ExitFunctions_returning_datetime is called when production functions_returning_datetime is exited.
func (s *BaseJPAListener) ExitFunctions_returning_datetime(ctx *Functions_returning_datetimeContext) {}

// EnterFunctions_returning_strings is called when production functions_returning_strings is entered.
func (s *BaseJPAListener) EnterFunctions_returning_strings(ctx *Functions_returning_stringsContext) {}

// ExitFunctions_returning_strings is called when production functions_returning_strings is exited.
func (s *BaseJPAListener) ExitFunctions_returning_strings(ctx *Functions_returning_stringsContext) {}

// EnterTrim_specification is called when production trim_specification is entered.
func (s *BaseJPAListener) EnterTrim_specification(ctx *Trim_specificationContext) {}

// ExitTrim_specification is called when production trim_specification is exited.
func (s *BaseJPAListener) ExitTrim_specification(ctx *Trim_specificationContext) {}

// EnterNumeric_literal is called when production numeric_literal is entered.
func (s *BaseJPAListener) EnterNumeric_literal(ctx *Numeric_literalContext) {}

// ExitNumeric_literal is called when production numeric_literal is exited.
func (s *BaseJPAListener) ExitNumeric_literal(ctx *Numeric_literalContext) {}

// EnterPattern_value is called when production pattern_value is entered.
func (s *BaseJPAListener) EnterPattern_value(ctx *Pattern_valueContext) {}

// ExitPattern_value is called when production pattern_value is exited.
func (s *BaseJPAListener) ExitPattern_value(ctx *Pattern_valueContext) {}

// EnterInput_parameter is called when production input_parameter is entered.
func (s *BaseJPAListener) EnterInput_parameter(ctx *Input_parameterContext) {}

// ExitInput_parameter is called when production input_parameter is exited.
func (s *BaseJPAListener) ExitInput_parameter(ctx *Input_parameterContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJPAListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJPAListener) ExitLiteral(ctx *LiteralContext) {}

// EnterConstructor_name is called when production constructor_name is entered.
func (s *BaseJPAListener) EnterConstructor_name(ctx *Constructor_nameContext) {}

// ExitConstructor_name is called when production constructor_name is exited.
func (s *BaseJPAListener) ExitConstructor_name(ctx *Constructor_nameContext) {}

// EnterEnum_literal is called when production enum_literal is entered.
func (s *BaseJPAListener) EnterEnum_literal(ctx *Enum_literalContext) {}

// ExitEnum_literal is called when production enum_literal is exited.
func (s *BaseJPAListener) ExitEnum_literal(ctx *Enum_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseJPAListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseJPAListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterSimple_state_field is called when production simple_state_field is entered.
func (s *BaseJPAListener) EnterSimple_state_field(ctx *Simple_state_fieldContext) {}

// ExitSimple_state_field is called when production simple_state_field is exited.
func (s *BaseJPAListener) ExitSimple_state_field(ctx *Simple_state_fieldContext) {}

// EnterEmbedded_class_state_field is called when production embedded_class_state_field is entered.
func (s *BaseJPAListener) EnterEmbedded_class_state_field(ctx *Embedded_class_state_fieldContext) {}

// ExitEmbedded_class_state_field is called when production embedded_class_state_field is exited.
func (s *BaseJPAListener) ExitEmbedded_class_state_field(ctx *Embedded_class_state_fieldContext) {}

// EnterSingle_valued_association_field is called when production single_valued_association_field is entered.
func (s *BaseJPAListener) EnterSingle_valued_association_field(ctx *Single_valued_association_fieldContext) {
}

// ExitSingle_valued_association_field is called when production single_valued_association_field is exited.
func (s *BaseJPAListener) ExitSingle_valued_association_field(ctx *Single_valued_association_fieldContext) {
}

// EnterCollection_valued_association_field is called when production collection_valued_association_field is entered.
func (s *BaseJPAListener) EnterCollection_valued_association_field(ctx *Collection_valued_association_fieldContext) {
}

// ExitCollection_valued_association_field is called when production collection_valued_association_field is exited.
func (s *BaseJPAListener) ExitCollection_valued_association_field(ctx *Collection_valued_association_fieldContext) {
}

// EnterAbstract_schema_name is called when production abstract_schema_name is entered.
func (s *BaseJPAListener) EnterAbstract_schema_name(ctx *Abstract_schema_nameContext) {}

// ExitAbstract_schema_name is called when production abstract_schema_name is exited.
func (s *BaseJPAListener) ExitAbstract_schema_name(ctx *Abstract_schema_nameContext) {}
