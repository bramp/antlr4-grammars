// Code generated from PowerQueryParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerquery // PowerQueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePowerQueryParserListener is a complete listener for a parse tree produced by PowerQueryParser.
type BasePowerQueryParserListener struct{}

var _ PowerQueryParserListener = &BasePowerQueryParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePowerQueryParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePowerQueryParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePowerQueryParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePowerQueryParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BasePowerQueryParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BasePowerQueryParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterSection_document is called when production section_document is entered.
func (s *BasePowerQueryParserListener) EnterSection_document(ctx *Section_documentContext) {}

// ExitSection_document is called when production section_document is exited.
func (s *BasePowerQueryParserListener) ExitSection_document(ctx *Section_documentContext) {}

// EnterSection is called when production section is entered.
func (s *BasePowerQueryParserListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BasePowerQueryParserListener) ExitSection(ctx *SectionContext) {}

// EnterSection_name is called when production section_name is entered.
func (s *BasePowerQueryParserListener) EnterSection_name(ctx *Section_nameContext) {}

// ExitSection_name is called when production section_name is exited.
func (s *BasePowerQueryParserListener) ExitSection_name(ctx *Section_nameContext) {}

// EnterSection_members is called when production section_members is entered.
func (s *BasePowerQueryParserListener) EnterSection_members(ctx *Section_membersContext) {}

// ExitSection_members is called when production section_members is exited.
func (s *BasePowerQueryParserListener) ExitSection_members(ctx *Section_membersContext) {}

// EnterSection_member is called when production section_member is entered.
func (s *BasePowerQueryParserListener) EnterSection_member(ctx *Section_memberContext) {}

// ExitSection_member is called when production section_member is exited.
func (s *BasePowerQueryParserListener) ExitSection_member(ctx *Section_memberContext) {}

// EnterSection_member_name is called when production section_member_name is entered.
func (s *BasePowerQueryParserListener) EnterSection_member_name(ctx *Section_member_nameContext) {}

// ExitSection_member_name is called when production section_member_name is exited.
func (s *BasePowerQueryParserListener) ExitSection_member_name(ctx *Section_member_nameContext) {}

// EnterExpression_document is called when production expression_document is entered.
func (s *BasePowerQueryParserListener) EnterExpression_document(ctx *Expression_documentContext) {}

// ExitExpression_document is called when production expression_document is exited.
func (s *BasePowerQueryParserListener) ExitExpression_document(ctx *Expression_documentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePowerQueryParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePowerQueryParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogical_or_expression is called when production logical_or_expression is entered.
func (s *BasePowerQueryParserListener) EnterLogical_or_expression(ctx *Logical_or_expressionContext) {
}

// ExitLogical_or_expression is called when production logical_or_expression is exited.
func (s *BasePowerQueryParserListener) ExitLogical_or_expression(ctx *Logical_or_expressionContext) {}

// EnterLogical_and_expression is called when production logical_and_expression is entered.
func (s *BasePowerQueryParserListener) EnterLogical_and_expression(ctx *Logical_and_expressionContext) {
}

// ExitLogical_and_expression is called when production logical_and_expression is exited.
func (s *BasePowerQueryParserListener) ExitLogical_and_expression(ctx *Logical_and_expressionContext) {
}

// EnterIs_expression is called when production is_expression is entered.
func (s *BasePowerQueryParserListener) EnterIs_expression(ctx *Is_expressionContext) {}

// ExitIs_expression is called when production is_expression is exited.
func (s *BasePowerQueryParserListener) ExitIs_expression(ctx *Is_expressionContext) {}

// EnterNullable_primitive_type is called when production nullable_primitive_type is entered.
func (s *BasePowerQueryParserListener) EnterNullable_primitive_type(ctx *Nullable_primitive_typeContext) {
}

// ExitNullable_primitive_type is called when production nullable_primitive_type is exited.
func (s *BasePowerQueryParserListener) ExitNullable_primitive_type(ctx *Nullable_primitive_typeContext) {
}

// EnterAs_expression is called when production as_expression is entered.
func (s *BasePowerQueryParserListener) EnterAs_expression(ctx *As_expressionContext) {}

// ExitAs_expression is called when production as_expression is exited.
func (s *BasePowerQueryParserListener) ExitAs_expression(ctx *As_expressionContext) {}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BasePowerQueryParserListener) EnterEquality_expression(ctx *Equality_expressionContext) {}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BasePowerQueryParserListener) ExitEquality_expression(ctx *Equality_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BasePowerQueryParserListener) EnterRelational_expression(ctx *Relational_expressionContext) {
}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BasePowerQueryParserListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BasePowerQueryParserListener) EnterAdditive_expression(ctx *Additive_expressionContext) {}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BasePowerQueryParserListener) ExitAdditive_expression(ctx *Additive_expressionContext) {}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BasePowerQueryParserListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {
}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BasePowerQueryParserListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {
}

// EnterMetadata_expression is called when production metadata_expression is entered.
func (s *BasePowerQueryParserListener) EnterMetadata_expression(ctx *Metadata_expressionContext) {}

// ExitMetadata_expression is called when production metadata_expression is exited.
func (s *BasePowerQueryParserListener) ExitMetadata_expression(ctx *Metadata_expressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BasePowerQueryParserListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BasePowerQueryParserListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BasePowerQueryParserListener) EnterPrimary_expression(ctx *Primary_expressionContext) {}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BasePowerQueryParserListener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterLiteral_expression is called when production literal_expression is entered.
func (s *BasePowerQueryParserListener) EnterLiteral_expression(ctx *Literal_expressionContext) {}

// ExitLiteral_expression is called when production literal_expression is exited.
func (s *BasePowerQueryParserListener) ExitLiteral_expression(ctx *Literal_expressionContext) {}

// EnterIdentifier_expression is called when production identifier_expression is entered.
func (s *BasePowerQueryParserListener) EnterIdentifier_expression(ctx *Identifier_expressionContext) {
}

// ExitIdentifier_expression is called when production identifier_expression is exited.
func (s *BasePowerQueryParserListener) ExitIdentifier_expression(ctx *Identifier_expressionContext) {}

// EnterIdentifier_reference is called when production identifier_reference is entered.
func (s *BasePowerQueryParserListener) EnterIdentifier_reference(ctx *Identifier_referenceContext) {}

// ExitIdentifier_reference is called when production identifier_reference is exited.
func (s *BasePowerQueryParserListener) ExitIdentifier_reference(ctx *Identifier_referenceContext) {}

// EnterExclusive_identifier_reference is called when production exclusive_identifier_reference is entered.
func (s *BasePowerQueryParserListener) EnterExclusive_identifier_reference(ctx *Exclusive_identifier_referenceContext) {
}

// ExitExclusive_identifier_reference is called when production exclusive_identifier_reference is exited.
func (s *BasePowerQueryParserListener) ExitExclusive_identifier_reference(ctx *Exclusive_identifier_referenceContext) {
}

// EnterInclusive_identifier_reference is called when production inclusive_identifier_reference is entered.
func (s *BasePowerQueryParserListener) EnterInclusive_identifier_reference(ctx *Inclusive_identifier_referenceContext) {
}

// ExitInclusive_identifier_reference is called when production inclusive_identifier_reference is exited.
func (s *BasePowerQueryParserListener) ExitInclusive_identifier_reference(ctx *Inclusive_identifier_referenceContext) {
}

// EnterSection_access_expression is called when production section_access_expression is entered.
func (s *BasePowerQueryParserListener) EnterSection_access_expression(ctx *Section_access_expressionContext) {
}

// ExitSection_access_expression is called when production section_access_expression is exited.
func (s *BasePowerQueryParserListener) ExitSection_access_expression(ctx *Section_access_expressionContext) {
}

// EnterParenthesized_expression is called when production parenthesized_expression is entered.
func (s *BasePowerQueryParserListener) EnterParenthesized_expression(ctx *Parenthesized_expressionContext) {
}

// ExitParenthesized_expression is called when production parenthesized_expression is exited.
func (s *BasePowerQueryParserListener) ExitParenthesized_expression(ctx *Parenthesized_expressionContext) {
}

// EnterNot_implemented_expression is called when production not_implemented_expression is entered.
func (s *BasePowerQueryParserListener) EnterNot_implemented_expression(ctx *Not_implemented_expressionContext) {
}

// ExitNot_implemented_expression is called when production not_implemented_expression is exited.
func (s *BasePowerQueryParserListener) ExitNot_implemented_expression(ctx *Not_implemented_expressionContext) {
}

// EnterArgument_list is called when production argument_list is entered.
func (s *BasePowerQueryParserListener) EnterArgument_list(ctx *Argument_listContext) {}

// ExitArgument_list is called when production argument_list is exited.
func (s *BasePowerQueryParserListener) ExitArgument_list(ctx *Argument_listContext) {}

// EnterList_expression is called when production list_expression is entered.
func (s *BasePowerQueryParserListener) EnterList_expression(ctx *List_expressionContext) {}

// ExitList_expression is called when production list_expression is exited.
func (s *BasePowerQueryParserListener) ExitList_expression(ctx *List_expressionContext) {}

// EnterItem_list is called when production item_list is entered.
func (s *BasePowerQueryParserListener) EnterItem_list(ctx *Item_listContext) {}

// ExitItem_list is called when production item_list is exited.
func (s *BasePowerQueryParserListener) ExitItem_list(ctx *Item_listContext) {}

// EnterItem is called when production item is entered.
func (s *BasePowerQueryParserListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasePowerQueryParserListener) ExitItem(ctx *ItemContext) {}

// EnterRecord_expression is called when production record_expression is entered.
func (s *BasePowerQueryParserListener) EnterRecord_expression(ctx *Record_expressionContext) {}

// ExitRecord_expression is called when production record_expression is exited.
func (s *BasePowerQueryParserListener) ExitRecord_expression(ctx *Record_expressionContext) {}

// EnterField_list is called when production field_list is entered.
func (s *BasePowerQueryParserListener) EnterField_list(ctx *Field_listContext) {}

// ExitField_list is called when production field_list is exited.
func (s *BasePowerQueryParserListener) ExitField_list(ctx *Field_listContext) {}

// EnterField is called when production field is entered.
func (s *BasePowerQueryParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasePowerQueryParserListener) ExitField(ctx *FieldContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BasePowerQueryParserListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BasePowerQueryParserListener) ExitField_name(ctx *Field_nameContext) {}

// EnterItem_selector is called when production item_selector is entered.
func (s *BasePowerQueryParserListener) EnterItem_selector(ctx *Item_selectorContext) {}

// ExitItem_selector is called when production item_selector is exited.
func (s *BasePowerQueryParserListener) ExitItem_selector(ctx *Item_selectorContext) {}

// EnterField_selector is called when production field_selector is entered.
func (s *BasePowerQueryParserListener) EnterField_selector(ctx *Field_selectorContext) {}

// ExitField_selector is called when production field_selector is exited.
func (s *BasePowerQueryParserListener) ExitField_selector(ctx *Field_selectorContext) {}

// EnterRequired_field_selector is called when production required_field_selector is entered.
func (s *BasePowerQueryParserListener) EnterRequired_field_selector(ctx *Required_field_selectorContext) {
}

// ExitRequired_field_selector is called when production required_field_selector is exited.
func (s *BasePowerQueryParserListener) ExitRequired_field_selector(ctx *Required_field_selectorContext) {
}

// EnterOptional_field_selector is called when production optional_field_selector is entered.
func (s *BasePowerQueryParserListener) EnterOptional_field_selector(ctx *Optional_field_selectorContext) {
}

// ExitOptional_field_selector is called when production optional_field_selector is exited.
func (s *BasePowerQueryParserListener) ExitOptional_field_selector(ctx *Optional_field_selectorContext) {
}

// EnterImplicit_target_field_selection is called when production implicit_target_field_selection is entered.
func (s *BasePowerQueryParserListener) EnterImplicit_target_field_selection(ctx *Implicit_target_field_selectionContext) {
}

// ExitImplicit_target_field_selection is called when production implicit_target_field_selection is exited.
func (s *BasePowerQueryParserListener) ExitImplicit_target_field_selection(ctx *Implicit_target_field_selectionContext) {
}

// EnterRequired_projection is called when production required_projection is entered.
func (s *BasePowerQueryParserListener) EnterRequired_projection(ctx *Required_projectionContext) {}

// ExitRequired_projection is called when production required_projection is exited.
func (s *BasePowerQueryParserListener) ExitRequired_projection(ctx *Required_projectionContext) {}

// EnterOptional_projection is called when production optional_projection is entered.
func (s *BasePowerQueryParserListener) EnterOptional_projection(ctx *Optional_projectionContext) {}

// ExitOptional_projection is called when production optional_projection is exited.
func (s *BasePowerQueryParserListener) ExitOptional_projection(ctx *Optional_projectionContext) {}

// EnterRequired_selector_list is called when production required_selector_list is entered.
func (s *BasePowerQueryParserListener) EnterRequired_selector_list(ctx *Required_selector_listContext) {
}

// ExitRequired_selector_list is called when production required_selector_list is exited.
func (s *BasePowerQueryParserListener) ExitRequired_selector_list(ctx *Required_selector_listContext) {
}

// EnterImplicit_target_projection is called when production implicit_target_projection is entered.
func (s *BasePowerQueryParserListener) EnterImplicit_target_projection(ctx *Implicit_target_projectionContext) {
}

// ExitImplicit_target_projection is called when production implicit_target_projection is exited.
func (s *BasePowerQueryParserListener) ExitImplicit_target_projection(ctx *Implicit_target_projectionContext) {
}

// EnterFunction_expression is called when production function_expression is entered.
func (s *BasePowerQueryParserListener) EnterFunction_expression(ctx *Function_expressionContext) {}

// ExitFunction_expression is called when production function_expression is exited.
func (s *BasePowerQueryParserListener) ExitFunction_expression(ctx *Function_expressionContext) {}

// EnterFunction_body is called when production function_body is entered.
func (s *BasePowerQueryParserListener) EnterFunction_body(ctx *Function_bodyContext) {}

// ExitFunction_body is called when production function_body is exited.
func (s *BasePowerQueryParserListener) ExitFunction_body(ctx *Function_bodyContext) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BasePowerQueryParserListener) EnterParameter_list(ctx *Parameter_listContext) {}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BasePowerQueryParserListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterFixed_parameter_list is called when production fixed_parameter_list is entered.
func (s *BasePowerQueryParserListener) EnterFixed_parameter_list(ctx *Fixed_parameter_listContext) {}

// ExitFixed_parameter_list is called when production fixed_parameter_list is exited.
func (s *BasePowerQueryParserListener) ExitFixed_parameter_list(ctx *Fixed_parameter_listContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePowerQueryParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePowerQueryParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterParameter_name is called when production parameter_name is entered.
func (s *BasePowerQueryParserListener) EnterParameter_name(ctx *Parameter_nameContext) {}

// ExitParameter_name is called when production parameter_name is exited.
func (s *BasePowerQueryParserListener) ExitParameter_name(ctx *Parameter_nameContext) {}

// EnterParameter_type is called when production parameter_type is entered.
func (s *BasePowerQueryParserListener) EnterParameter_type(ctx *Parameter_typeContext) {}

// ExitParameter_type is called when production parameter_type is exited.
func (s *BasePowerQueryParserListener) ExitParameter_type(ctx *Parameter_typeContext) {}

// EnterReturn_type is called when production return_type is entered.
func (s *BasePowerQueryParserListener) EnterReturn_type(ctx *Return_typeContext) {}

// ExitReturn_type is called when production return_type is exited.
func (s *BasePowerQueryParserListener) ExitReturn_type(ctx *Return_typeContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BasePowerQueryParserListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BasePowerQueryParserListener) ExitAssertion(ctx *AssertionContext) {}

// EnterOptional_parameter_list is called when production optional_parameter_list is entered.
func (s *BasePowerQueryParserListener) EnterOptional_parameter_list(ctx *Optional_parameter_listContext) {
}

// ExitOptional_parameter_list is called when production optional_parameter_list is exited.
func (s *BasePowerQueryParserListener) ExitOptional_parameter_list(ctx *Optional_parameter_listContext) {
}

// EnterOptional_parameter is called when production optional_parameter is entered.
func (s *BasePowerQueryParserListener) EnterOptional_parameter(ctx *Optional_parameterContext) {}

// ExitOptional_parameter is called when production optional_parameter is exited.
func (s *BasePowerQueryParserListener) ExitOptional_parameter(ctx *Optional_parameterContext) {}

// EnterEach_expression is called when production each_expression is entered.
func (s *BasePowerQueryParserListener) EnterEach_expression(ctx *Each_expressionContext) {}

// ExitEach_expression is called when production each_expression is exited.
func (s *BasePowerQueryParserListener) ExitEach_expression(ctx *Each_expressionContext) {}

// EnterEach_expression_body is called when production each_expression_body is entered.
func (s *BasePowerQueryParserListener) EnterEach_expression_body(ctx *Each_expression_bodyContext) {}

// ExitEach_expression_body is called when production each_expression_body is exited.
func (s *BasePowerQueryParserListener) ExitEach_expression_body(ctx *Each_expression_bodyContext) {}

// EnterLet_expression is called when production let_expression is entered.
func (s *BasePowerQueryParserListener) EnterLet_expression(ctx *Let_expressionContext) {}

// ExitLet_expression is called when production let_expression is exited.
func (s *BasePowerQueryParserListener) ExitLet_expression(ctx *Let_expressionContext) {}

// EnterVariable_list is called when production variable_list is entered.
func (s *BasePowerQueryParserListener) EnterVariable_list(ctx *Variable_listContext) {}

// ExitVariable_list is called when production variable_list is exited.
func (s *BasePowerQueryParserListener) ExitVariable_list(ctx *Variable_listContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasePowerQueryParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasePowerQueryParserListener) ExitVariable(ctx *VariableContext) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BasePowerQueryParserListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BasePowerQueryParserListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterIf_expression is called when production if_expression is entered.
func (s *BasePowerQueryParserListener) EnterIf_expression(ctx *If_expressionContext) {}

// ExitIf_expression is called when production if_expression is exited.
func (s *BasePowerQueryParserListener) ExitIf_expression(ctx *If_expressionContext) {}

// EnterIf_condition is called when production if_condition is entered.
func (s *BasePowerQueryParserListener) EnterIf_condition(ctx *If_conditionContext) {}

// ExitIf_condition is called when production if_condition is exited.
func (s *BasePowerQueryParserListener) ExitIf_condition(ctx *If_conditionContext) {}

// EnterTrue_expression is called when production true_expression is entered.
func (s *BasePowerQueryParserListener) EnterTrue_expression(ctx *True_expressionContext) {}

// ExitTrue_expression is called when production true_expression is exited.
func (s *BasePowerQueryParserListener) ExitTrue_expression(ctx *True_expressionContext) {}

// EnterFalse_expression is called when production false_expression is entered.
func (s *BasePowerQueryParserListener) EnterFalse_expression(ctx *False_expressionContext) {}

// ExitFalse_expression is called when production false_expression is exited.
func (s *BasePowerQueryParserListener) ExitFalse_expression(ctx *False_expressionContext) {}

// EnterType_expression is called when production type_expression is entered.
func (s *BasePowerQueryParserListener) EnterType_expression(ctx *Type_expressionContext) {}

// ExitType_expression is called when production type_expression is exited.
func (s *BasePowerQueryParserListener) ExitType_expression(ctx *Type_expressionContext) {}

// EnterType_expr is called when production type_expr is entered.
func (s *BasePowerQueryParserListener) EnterType_expr(ctx *Type_exprContext) {}

// ExitType_expr is called when production type_expr is exited.
func (s *BasePowerQueryParserListener) ExitType_expr(ctx *Type_exprContext) {}

// EnterPrimary_type is called when production primary_type is entered.
func (s *BasePowerQueryParserListener) EnterPrimary_type(ctx *Primary_typeContext) {}

// ExitPrimary_type is called when production primary_type is exited.
func (s *BasePowerQueryParserListener) ExitPrimary_type(ctx *Primary_typeContext) {}

// EnterPrimitive_type is called when production primitive_type is entered.
func (s *BasePowerQueryParserListener) EnterPrimitive_type(ctx *Primitive_typeContext) {}

// ExitPrimitive_type is called when production primitive_type is exited.
func (s *BasePowerQueryParserListener) ExitPrimitive_type(ctx *Primitive_typeContext) {}

// EnterRecord_type is called when production record_type is entered.
func (s *BasePowerQueryParserListener) EnterRecord_type(ctx *Record_typeContext) {}

// ExitRecord_type is called when production record_type is exited.
func (s *BasePowerQueryParserListener) ExitRecord_type(ctx *Record_typeContext) {}

// EnterField_specification_list is called when production field_specification_list is entered.
func (s *BasePowerQueryParserListener) EnterField_specification_list(ctx *Field_specification_listContext) {
}

// ExitField_specification_list is called when production field_specification_list is exited.
func (s *BasePowerQueryParserListener) ExitField_specification_list(ctx *Field_specification_listContext) {
}

// EnterField_specification is called when production field_specification is entered.
func (s *BasePowerQueryParserListener) EnterField_specification(ctx *Field_specificationContext) {}

// ExitField_specification is called when production field_specification is exited.
func (s *BasePowerQueryParserListener) ExitField_specification(ctx *Field_specificationContext) {}

// EnterField_type_specification is called when production field_type_specification is entered.
func (s *BasePowerQueryParserListener) EnterField_type_specification(ctx *Field_type_specificationContext) {
}

// ExitField_type_specification is called when production field_type_specification is exited.
func (s *BasePowerQueryParserListener) ExitField_type_specification(ctx *Field_type_specificationContext) {
}

// EnterField_type is called when production field_type is entered.
func (s *BasePowerQueryParserListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BasePowerQueryParserListener) ExitField_type(ctx *Field_typeContext) {}

// EnterOpen_record_marker is called when production open_record_marker is entered.
func (s *BasePowerQueryParserListener) EnterOpen_record_marker(ctx *Open_record_markerContext) {}

// ExitOpen_record_marker is called when production open_record_marker is exited.
func (s *BasePowerQueryParserListener) ExitOpen_record_marker(ctx *Open_record_markerContext) {}

// EnterList_type is called when production list_type is entered.
func (s *BasePowerQueryParserListener) EnterList_type(ctx *List_typeContext) {}

// ExitList_type is called when production list_type is exited.
func (s *BasePowerQueryParserListener) ExitList_type(ctx *List_typeContext) {}

// EnterItem_type is called when production item_type is entered.
func (s *BasePowerQueryParserListener) EnterItem_type(ctx *Item_typeContext) {}

// ExitItem_type is called when production item_type is exited.
func (s *BasePowerQueryParserListener) ExitItem_type(ctx *Item_typeContext) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BasePowerQueryParserListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BasePowerQueryParserListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterParameter_specification_list is called when production parameter_specification_list is entered.
func (s *BasePowerQueryParserListener) EnterParameter_specification_list(ctx *Parameter_specification_listContext) {
}

// ExitParameter_specification_list is called when production parameter_specification_list is exited.
func (s *BasePowerQueryParserListener) ExitParameter_specification_list(ctx *Parameter_specification_listContext) {
}

// EnterRequired_parameter_specification_list is called when production required_parameter_specification_list is entered.
func (s *BasePowerQueryParserListener) EnterRequired_parameter_specification_list(ctx *Required_parameter_specification_listContext) {
}

// ExitRequired_parameter_specification_list is called when production required_parameter_specification_list is exited.
func (s *BasePowerQueryParserListener) ExitRequired_parameter_specification_list(ctx *Required_parameter_specification_listContext) {
}

// EnterRequired_parameter_specification is called when production required_parameter_specification is entered.
func (s *BasePowerQueryParserListener) EnterRequired_parameter_specification(ctx *Required_parameter_specificationContext) {
}

// ExitRequired_parameter_specification is called when production required_parameter_specification is exited.
func (s *BasePowerQueryParserListener) ExitRequired_parameter_specification(ctx *Required_parameter_specificationContext) {
}

// EnterOptional_parameter_specification_list is called when production optional_parameter_specification_list is entered.
func (s *BasePowerQueryParserListener) EnterOptional_parameter_specification_list(ctx *Optional_parameter_specification_listContext) {
}

// ExitOptional_parameter_specification_list is called when production optional_parameter_specification_list is exited.
func (s *BasePowerQueryParserListener) ExitOptional_parameter_specification_list(ctx *Optional_parameter_specification_listContext) {
}

// EnterOptional_parameter_specification is called when production optional_parameter_specification is entered.
func (s *BasePowerQueryParserListener) EnterOptional_parameter_specification(ctx *Optional_parameter_specificationContext) {
}

// ExitOptional_parameter_specification is called when production optional_parameter_specification is exited.
func (s *BasePowerQueryParserListener) ExitOptional_parameter_specification(ctx *Optional_parameter_specificationContext) {
}

// EnterParameter_specification is called when production parameter_specification is entered.
func (s *BasePowerQueryParserListener) EnterParameter_specification(ctx *Parameter_specificationContext) {
}

// ExitParameter_specification is called when production parameter_specification is exited.
func (s *BasePowerQueryParserListener) ExitParameter_specification(ctx *Parameter_specificationContext) {
}

// EnterTable_type is called when production table_type is entered.
func (s *BasePowerQueryParserListener) EnterTable_type(ctx *Table_typeContext) {}

// ExitTable_type is called when production table_type is exited.
func (s *BasePowerQueryParserListener) ExitTable_type(ctx *Table_typeContext) {}

// EnterRow_type is called when production row_type is entered.
func (s *BasePowerQueryParserListener) EnterRow_type(ctx *Row_typeContext) {}

// ExitRow_type is called when production row_type is exited.
func (s *BasePowerQueryParserListener) ExitRow_type(ctx *Row_typeContext) {}

// EnterNullable_type is called when production nullable_type is entered.
func (s *BasePowerQueryParserListener) EnterNullable_type(ctx *Nullable_typeContext) {}

// ExitNullable_type is called when production nullable_type is exited.
func (s *BasePowerQueryParserListener) ExitNullable_type(ctx *Nullable_typeContext) {}

// EnterError_raising_expression is called when production error_raising_expression is entered.
func (s *BasePowerQueryParserListener) EnterError_raising_expression(ctx *Error_raising_expressionContext) {
}

// ExitError_raising_expression is called when production error_raising_expression is exited.
func (s *BasePowerQueryParserListener) ExitError_raising_expression(ctx *Error_raising_expressionContext) {
}

// EnterError_handling_expression is called when production error_handling_expression is entered.
func (s *BasePowerQueryParserListener) EnterError_handling_expression(ctx *Error_handling_expressionContext) {
}

// ExitError_handling_expression is called when production error_handling_expression is exited.
func (s *BasePowerQueryParserListener) ExitError_handling_expression(ctx *Error_handling_expressionContext) {
}

// EnterProtected_expression is called when production protected_expression is entered.
func (s *BasePowerQueryParserListener) EnterProtected_expression(ctx *Protected_expressionContext) {}

// ExitProtected_expression is called when production protected_expression is exited.
func (s *BasePowerQueryParserListener) ExitProtected_expression(ctx *Protected_expressionContext) {}

// EnterOtherwise_clause is called when production otherwise_clause is entered.
func (s *BasePowerQueryParserListener) EnterOtherwise_clause(ctx *Otherwise_clauseContext) {}

// ExitOtherwise_clause is called when production otherwise_clause is exited.
func (s *BasePowerQueryParserListener) ExitOtherwise_clause(ctx *Otherwise_clauseContext) {}

// EnterDefault_expression is called when production default_expression is entered.
func (s *BasePowerQueryParserListener) EnterDefault_expression(ctx *Default_expressionContext) {}

// ExitDefault_expression is called when production default_expression is exited.
func (s *BasePowerQueryParserListener) ExitDefault_expression(ctx *Default_expressionContext) {}

// EnterLiteral_attribs is called when production literal_attribs is entered.
func (s *BasePowerQueryParserListener) EnterLiteral_attribs(ctx *Literal_attribsContext) {}

// ExitLiteral_attribs is called when production literal_attribs is exited.
func (s *BasePowerQueryParserListener) ExitLiteral_attribs(ctx *Literal_attribsContext) {}

// EnterRecord_literal is called when production record_literal is entered.
func (s *BasePowerQueryParserListener) EnterRecord_literal(ctx *Record_literalContext) {}

// ExitRecord_literal is called when production record_literal is exited.
func (s *BasePowerQueryParserListener) ExitRecord_literal(ctx *Record_literalContext) {}

// EnterLiteral_field_list is called when production literal_field_list is entered.
func (s *BasePowerQueryParserListener) EnterLiteral_field_list(ctx *Literal_field_listContext) {}

// ExitLiteral_field_list is called when production literal_field_list is exited.
func (s *BasePowerQueryParserListener) ExitLiteral_field_list(ctx *Literal_field_listContext) {}

// EnterLiteral_field is called when production literal_field is entered.
func (s *BasePowerQueryParserListener) EnterLiteral_field(ctx *Literal_fieldContext) {}

// ExitLiteral_field is called when production literal_field is exited.
func (s *BasePowerQueryParserListener) ExitLiteral_field(ctx *Literal_fieldContext) {}

// EnterList_literal is called when production list_literal is entered.
func (s *BasePowerQueryParserListener) EnterList_literal(ctx *List_literalContext) {}

// ExitList_literal is called when production list_literal is exited.
func (s *BasePowerQueryParserListener) ExitList_literal(ctx *List_literalContext) {}

// EnterLiteral_item_list is called when production literal_item_list is entered.
func (s *BasePowerQueryParserListener) EnterLiteral_item_list(ctx *Literal_item_listContext) {}

// ExitLiteral_item_list is called when production literal_item_list is exited.
func (s *BasePowerQueryParserListener) ExitLiteral_item_list(ctx *Literal_item_listContext) {}

// EnterAny_literal is called when production any_literal is entered.
func (s *BasePowerQueryParserListener) EnterAny_literal(ctx *Any_literalContext) {}

// ExitAny_literal is called when production any_literal is exited.
func (s *BasePowerQueryParserListener) ExitAny_literal(ctx *Any_literalContext) {}
