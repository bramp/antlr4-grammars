// Code generated from PowerQueryParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerquery // PowerQueryParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PowerQueryParserListener is a complete listener for a parse tree produced by PowerQueryParser.
type PowerQueryParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterSection_document is called when entering the section_document production.
	EnterSection_document(c *Section_documentContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterSection_name is called when entering the section_name production.
	EnterSection_name(c *Section_nameContext)

	// EnterSection_members is called when entering the section_members production.
	EnterSection_members(c *Section_membersContext)

	// EnterSection_member is called when entering the section_member production.
	EnterSection_member(c *Section_memberContext)

	// EnterSection_member_name is called when entering the section_member_name production.
	EnterSection_member_name(c *Section_member_nameContext)

	// EnterExpression_document is called when entering the expression_document production.
	EnterExpression_document(c *Expression_documentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogical_or_expression is called when entering the logical_or_expression production.
	EnterLogical_or_expression(c *Logical_or_expressionContext)

	// EnterLogical_and_expression is called when entering the logical_and_expression production.
	EnterLogical_and_expression(c *Logical_and_expressionContext)

	// EnterIs_expression is called when entering the is_expression production.
	EnterIs_expression(c *Is_expressionContext)

	// EnterNullable_primitive_type is called when entering the nullable_primitive_type production.
	EnterNullable_primitive_type(c *Nullable_primitive_typeContext)

	// EnterAs_expression is called when entering the as_expression production.
	EnterAs_expression(c *As_expressionContext)

	// EnterEquality_expression is called when entering the equality_expression production.
	EnterEquality_expression(c *Equality_expressionContext)

	// EnterRelational_expression is called when entering the relational_expression production.
	EnterRelational_expression(c *Relational_expressionContext)

	// EnterAdditive_expression is called when entering the additive_expression production.
	EnterAdditive_expression(c *Additive_expressionContext)

	// EnterMultiplicative_expression is called when entering the multiplicative_expression production.
	EnterMultiplicative_expression(c *Multiplicative_expressionContext)

	// EnterMetadata_expression is called when entering the metadata_expression production.
	EnterMetadata_expression(c *Metadata_expressionContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterPrimary_expression is called when entering the primary_expression production.
	EnterPrimary_expression(c *Primary_expressionContext)

	// EnterLiteral_expression is called when entering the literal_expression production.
	EnterLiteral_expression(c *Literal_expressionContext)

	// EnterIdentifier_expression is called when entering the identifier_expression production.
	EnterIdentifier_expression(c *Identifier_expressionContext)

	// EnterIdentifier_reference is called when entering the identifier_reference production.
	EnterIdentifier_reference(c *Identifier_referenceContext)

	// EnterExclusive_identifier_reference is called when entering the exclusive_identifier_reference production.
	EnterExclusive_identifier_reference(c *Exclusive_identifier_referenceContext)

	// EnterInclusive_identifier_reference is called when entering the inclusive_identifier_reference production.
	EnterInclusive_identifier_reference(c *Inclusive_identifier_referenceContext)

	// EnterSection_access_expression is called when entering the section_access_expression production.
	EnterSection_access_expression(c *Section_access_expressionContext)

	// EnterParenthesized_expression is called when entering the parenthesized_expression production.
	EnterParenthesized_expression(c *Parenthesized_expressionContext)

	// EnterNot_implemented_expression is called when entering the not_implemented_expression production.
	EnterNot_implemented_expression(c *Not_implemented_expressionContext)

	// EnterArgument_list is called when entering the argument_list production.
	EnterArgument_list(c *Argument_listContext)

	// EnterList_expression is called when entering the list_expression production.
	EnterList_expression(c *List_expressionContext)

	// EnterItem_list is called when entering the item_list production.
	EnterItem_list(c *Item_listContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterRecord_expression is called when entering the record_expression production.
	EnterRecord_expression(c *Record_expressionContext)

	// EnterField_list is called when entering the field_list production.
	EnterField_list(c *Field_listContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterItem_selector is called when entering the item_selector production.
	EnterItem_selector(c *Item_selectorContext)

	// EnterField_selector is called when entering the field_selector production.
	EnterField_selector(c *Field_selectorContext)

	// EnterRequired_field_selector is called when entering the required_field_selector production.
	EnterRequired_field_selector(c *Required_field_selectorContext)

	// EnterOptional_field_selector is called when entering the optional_field_selector production.
	EnterOptional_field_selector(c *Optional_field_selectorContext)

	// EnterImplicit_target_field_selection is called when entering the implicit_target_field_selection production.
	EnterImplicit_target_field_selection(c *Implicit_target_field_selectionContext)

	// EnterRequired_projection is called when entering the required_projection production.
	EnterRequired_projection(c *Required_projectionContext)

	// EnterOptional_projection is called when entering the optional_projection production.
	EnterOptional_projection(c *Optional_projectionContext)

	// EnterRequired_selector_list is called when entering the required_selector_list production.
	EnterRequired_selector_list(c *Required_selector_listContext)

	// EnterImplicit_target_projection is called when entering the implicit_target_projection production.
	EnterImplicit_target_projection(c *Implicit_target_projectionContext)

	// EnterFunction_expression is called when entering the function_expression production.
	EnterFunction_expression(c *Function_expressionContext)

	// EnterFunction_body is called when entering the function_body production.
	EnterFunction_body(c *Function_bodyContext)

	// EnterParameter_list is called when entering the parameter_list production.
	EnterParameter_list(c *Parameter_listContext)

	// EnterFixed_parameter_list is called when entering the fixed_parameter_list production.
	EnterFixed_parameter_list(c *Fixed_parameter_listContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterParameter_name is called when entering the parameter_name production.
	EnterParameter_name(c *Parameter_nameContext)

	// EnterParameter_type is called when entering the parameter_type production.
	EnterParameter_type(c *Parameter_typeContext)

	// EnterReturn_type is called when entering the return_type production.
	EnterReturn_type(c *Return_typeContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterOptional_parameter_list is called when entering the optional_parameter_list production.
	EnterOptional_parameter_list(c *Optional_parameter_listContext)

	// EnterOptional_parameter is called when entering the optional_parameter production.
	EnterOptional_parameter(c *Optional_parameterContext)

	// EnterEach_expression is called when entering the each_expression production.
	EnterEach_expression(c *Each_expressionContext)

	// EnterEach_expression_body is called when entering the each_expression_body production.
	EnterEach_expression_body(c *Each_expression_bodyContext)

	// EnterLet_expression is called when entering the let_expression production.
	EnterLet_expression(c *Let_expressionContext)

	// EnterVariable_list is called when entering the variable_list production.
	EnterVariable_list(c *Variable_listContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVariable_name is called when entering the variable_name production.
	EnterVariable_name(c *Variable_nameContext)

	// EnterIf_expression is called when entering the if_expression production.
	EnterIf_expression(c *If_expressionContext)

	// EnterIf_condition is called when entering the if_condition production.
	EnterIf_condition(c *If_conditionContext)

	// EnterTrue_expression is called when entering the true_expression production.
	EnterTrue_expression(c *True_expressionContext)

	// EnterFalse_expression is called when entering the false_expression production.
	EnterFalse_expression(c *False_expressionContext)

	// EnterType_expression is called when entering the type_expression production.
	EnterType_expression(c *Type_expressionContext)

	// EnterType_expr is called when entering the type_expr production.
	EnterType_expr(c *Type_exprContext)

	// EnterPrimary_type is called when entering the primary_type production.
	EnterPrimary_type(c *Primary_typeContext)

	// EnterPrimitive_type is called when entering the primitive_type production.
	EnterPrimitive_type(c *Primitive_typeContext)

	// EnterRecord_type is called when entering the record_type production.
	EnterRecord_type(c *Record_typeContext)

	// EnterField_specification_list is called when entering the field_specification_list production.
	EnterField_specification_list(c *Field_specification_listContext)

	// EnterField_specification is called when entering the field_specification production.
	EnterField_specification(c *Field_specificationContext)

	// EnterField_type_specification is called when entering the field_type_specification production.
	EnterField_type_specification(c *Field_type_specificationContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// EnterOpen_record_marker is called when entering the open_record_marker production.
	EnterOpen_record_marker(c *Open_record_markerContext)

	// EnterList_type is called when entering the list_type production.
	EnterList_type(c *List_typeContext)

	// EnterItem_type is called when entering the item_type production.
	EnterItem_type(c *Item_typeContext)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterParameter_specification_list is called when entering the parameter_specification_list production.
	EnterParameter_specification_list(c *Parameter_specification_listContext)

	// EnterRequired_parameter_specification_list is called when entering the required_parameter_specification_list production.
	EnterRequired_parameter_specification_list(c *Required_parameter_specification_listContext)

	// EnterRequired_parameter_specification is called when entering the required_parameter_specification production.
	EnterRequired_parameter_specification(c *Required_parameter_specificationContext)

	// EnterOptional_parameter_specification_list is called when entering the optional_parameter_specification_list production.
	EnterOptional_parameter_specification_list(c *Optional_parameter_specification_listContext)

	// EnterOptional_parameter_specification is called when entering the optional_parameter_specification production.
	EnterOptional_parameter_specification(c *Optional_parameter_specificationContext)

	// EnterParameter_specification is called when entering the parameter_specification production.
	EnterParameter_specification(c *Parameter_specificationContext)

	// EnterTable_type is called when entering the table_type production.
	EnterTable_type(c *Table_typeContext)

	// EnterRow_type is called when entering the row_type production.
	EnterRow_type(c *Row_typeContext)

	// EnterNullable_type is called when entering the nullable_type production.
	EnterNullable_type(c *Nullable_typeContext)

	// EnterError_raising_expression is called when entering the error_raising_expression production.
	EnterError_raising_expression(c *Error_raising_expressionContext)

	// EnterError_handling_expression is called when entering the error_handling_expression production.
	EnterError_handling_expression(c *Error_handling_expressionContext)

	// EnterProtected_expression is called when entering the protected_expression production.
	EnterProtected_expression(c *Protected_expressionContext)

	// EnterOtherwise_clause is called when entering the otherwise_clause production.
	EnterOtherwise_clause(c *Otherwise_clauseContext)

	// EnterDefault_expression is called when entering the default_expression production.
	EnterDefault_expression(c *Default_expressionContext)

	// EnterLiteral_attribs is called when entering the literal_attribs production.
	EnterLiteral_attribs(c *Literal_attribsContext)

	// EnterRecord_literal is called when entering the record_literal production.
	EnterRecord_literal(c *Record_literalContext)

	// EnterLiteral_field_list is called when entering the literal_field_list production.
	EnterLiteral_field_list(c *Literal_field_listContext)

	// EnterLiteral_field is called when entering the literal_field production.
	EnterLiteral_field(c *Literal_fieldContext)

	// EnterList_literal is called when entering the list_literal production.
	EnterList_literal(c *List_literalContext)

	// EnterLiteral_item_list is called when entering the literal_item_list production.
	EnterLiteral_item_list(c *Literal_item_listContext)

	// EnterAny_literal is called when entering the any_literal production.
	EnterAny_literal(c *Any_literalContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitSection_document is called when exiting the section_document production.
	ExitSection_document(c *Section_documentContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitSection_name is called when exiting the section_name production.
	ExitSection_name(c *Section_nameContext)

	// ExitSection_members is called when exiting the section_members production.
	ExitSection_members(c *Section_membersContext)

	// ExitSection_member is called when exiting the section_member production.
	ExitSection_member(c *Section_memberContext)

	// ExitSection_member_name is called when exiting the section_member_name production.
	ExitSection_member_name(c *Section_member_nameContext)

	// ExitExpression_document is called when exiting the expression_document production.
	ExitExpression_document(c *Expression_documentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogical_or_expression is called when exiting the logical_or_expression production.
	ExitLogical_or_expression(c *Logical_or_expressionContext)

	// ExitLogical_and_expression is called when exiting the logical_and_expression production.
	ExitLogical_and_expression(c *Logical_and_expressionContext)

	// ExitIs_expression is called when exiting the is_expression production.
	ExitIs_expression(c *Is_expressionContext)

	// ExitNullable_primitive_type is called when exiting the nullable_primitive_type production.
	ExitNullable_primitive_type(c *Nullable_primitive_typeContext)

	// ExitAs_expression is called when exiting the as_expression production.
	ExitAs_expression(c *As_expressionContext)

	// ExitEquality_expression is called when exiting the equality_expression production.
	ExitEquality_expression(c *Equality_expressionContext)

	// ExitRelational_expression is called when exiting the relational_expression production.
	ExitRelational_expression(c *Relational_expressionContext)

	// ExitAdditive_expression is called when exiting the additive_expression production.
	ExitAdditive_expression(c *Additive_expressionContext)

	// ExitMultiplicative_expression is called when exiting the multiplicative_expression production.
	ExitMultiplicative_expression(c *Multiplicative_expressionContext)

	// ExitMetadata_expression is called when exiting the metadata_expression production.
	ExitMetadata_expression(c *Metadata_expressionContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitPrimary_expression is called when exiting the primary_expression production.
	ExitPrimary_expression(c *Primary_expressionContext)

	// ExitLiteral_expression is called when exiting the literal_expression production.
	ExitLiteral_expression(c *Literal_expressionContext)

	// ExitIdentifier_expression is called when exiting the identifier_expression production.
	ExitIdentifier_expression(c *Identifier_expressionContext)

	// ExitIdentifier_reference is called when exiting the identifier_reference production.
	ExitIdentifier_reference(c *Identifier_referenceContext)

	// ExitExclusive_identifier_reference is called when exiting the exclusive_identifier_reference production.
	ExitExclusive_identifier_reference(c *Exclusive_identifier_referenceContext)

	// ExitInclusive_identifier_reference is called when exiting the inclusive_identifier_reference production.
	ExitInclusive_identifier_reference(c *Inclusive_identifier_referenceContext)

	// ExitSection_access_expression is called when exiting the section_access_expression production.
	ExitSection_access_expression(c *Section_access_expressionContext)

	// ExitParenthesized_expression is called when exiting the parenthesized_expression production.
	ExitParenthesized_expression(c *Parenthesized_expressionContext)

	// ExitNot_implemented_expression is called when exiting the not_implemented_expression production.
	ExitNot_implemented_expression(c *Not_implemented_expressionContext)

	// ExitArgument_list is called when exiting the argument_list production.
	ExitArgument_list(c *Argument_listContext)

	// ExitList_expression is called when exiting the list_expression production.
	ExitList_expression(c *List_expressionContext)

	// ExitItem_list is called when exiting the item_list production.
	ExitItem_list(c *Item_listContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitRecord_expression is called when exiting the record_expression production.
	ExitRecord_expression(c *Record_expressionContext)

	// ExitField_list is called when exiting the field_list production.
	ExitField_list(c *Field_listContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitItem_selector is called when exiting the item_selector production.
	ExitItem_selector(c *Item_selectorContext)

	// ExitField_selector is called when exiting the field_selector production.
	ExitField_selector(c *Field_selectorContext)

	// ExitRequired_field_selector is called when exiting the required_field_selector production.
	ExitRequired_field_selector(c *Required_field_selectorContext)

	// ExitOptional_field_selector is called when exiting the optional_field_selector production.
	ExitOptional_field_selector(c *Optional_field_selectorContext)

	// ExitImplicit_target_field_selection is called when exiting the implicit_target_field_selection production.
	ExitImplicit_target_field_selection(c *Implicit_target_field_selectionContext)

	// ExitRequired_projection is called when exiting the required_projection production.
	ExitRequired_projection(c *Required_projectionContext)

	// ExitOptional_projection is called when exiting the optional_projection production.
	ExitOptional_projection(c *Optional_projectionContext)

	// ExitRequired_selector_list is called when exiting the required_selector_list production.
	ExitRequired_selector_list(c *Required_selector_listContext)

	// ExitImplicit_target_projection is called when exiting the implicit_target_projection production.
	ExitImplicit_target_projection(c *Implicit_target_projectionContext)

	// ExitFunction_expression is called when exiting the function_expression production.
	ExitFunction_expression(c *Function_expressionContext)

	// ExitFunction_body is called when exiting the function_body production.
	ExitFunction_body(c *Function_bodyContext)

	// ExitParameter_list is called when exiting the parameter_list production.
	ExitParameter_list(c *Parameter_listContext)

	// ExitFixed_parameter_list is called when exiting the fixed_parameter_list production.
	ExitFixed_parameter_list(c *Fixed_parameter_listContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitParameter_name is called when exiting the parameter_name production.
	ExitParameter_name(c *Parameter_nameContext)

	// ExitParameter_type is called when exiting the parameter_type production.
	ExitParameter_type(c *Parameter_typeContext)

	// ExitReturn_type is called when exiting the return_type production.
	ExitReturn_type(c *Return_typeContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitOptional_parameter_list is called when exiting the optional_parameter_list production.
	ExitOptional_parameter_list(c *Optional_parameter_listContext)

	// ExitOptional_parameter is called when exiting the optional_parameter production.
	ExitOptional_parameter(c *Optional_parameterContext)

	// ExitEach_expression is called when exiting the each_expression production.
	ExitEach_expression(c *Each_expressionContext)

	// ExitEach_expression_body is called when exiting the each_expression_body production.
	ExitEach_expression_body(c *Each_expression_bodyContext)

	// ExitLet_expression is called when exiting the let_expression production.
	ExitLet_expression(c *Let_expressionContext)

	// ExitVariable_list is called when exiting the variable_list production.
	ExitVariable_list(c *Variable_listContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVariable_name is called when exiting the variable_name production.
	ExitVariable_name(c *Variable_nameContext)

	// ExitIf_expression is called when exiting the if_expression production.
	ExitIf_expression(c *If_expressionContext)

	// ExitIf_condition is called when exiting the if_condition production.
	ExitIf_condition(c *If_conditionContext)

	// ExitTrue_expression is called when exiting the true_expression production.
	ExitTrue_expression(c *True_expressionContext)

	// ExitFalse_expression is called when exiting the false_expression production.
	ExitFalse_expression(c *False_expressionContext)

	// ExitType_expression is called when exiting the type_expression production.
	ExitType_expression(c *Type_expressionContext)

	// ExitType_expr is called when exiting the type_expr production.
	ExitType_expr(c *Type_exprContext)

	// ExitPrimary_type is called when exiting the primary_type production.
	ExitPrimary_type(c *Primary_typeContext)

	// ExitPrimitive_type is called when exiting the primitive_type production.
	ExitPrimitive_type(c *Primitive_typeContext)

	// ExitRecord_type is called when exiting the record_type production.
	ExitRecord_type(c *Record_typeContext)

	// ExitField_specification_list is called when exiting the field_specification_list production.
	ExitField_specification_list(c *Field_specification_listContext)

	// ExitField_specification is called when exiting the field_specification production.
	ExitField_specification(c *Field_specificationContext)

	// ExitField_type_specification is called when exiting the field_type_specification production.
	ExitField_type_specification(c *Field_type_specificationContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)

	// ExitOpen_record_marker is called when exiting the open_record_marker production.
	ExitOpen_record_marker(c *Open_record_markerContext)

	// ExitList_type is called when exiting the list_type production.
	ExitList_type(c *List_typeContext)

	// ExitItem_type is called when exiting the item_type production.
	ExitItem_type(c *Item_typeContext)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitParameter_specification_list is called when exiting the parameter_specification_list production.
	ExitParameter_specification_list(c *Parameter_specification_listContext)

	// ExitRequired_parameter_specification_list is called when exiting the required_parameter_specification_list production.
	ExitRequired_parameter_specification_list(c *Required_parameter_specification_listContext)

	// ExitRequired_parameter_specification is called when exiting the required_parameter_specification production.
	ExitRequired_parameter_specification(c *Required_parameter_specificationContext)

	// ExitOptional_parameter_specification_list is called when exiting the optional_parameter_specification_list production.
	ExitOptional_parameter_specification_list(c *Optional_parameter_specification_listContext)

	// ExitOptional_parameter_specification is called when exiting the optional_parameter_specification production.
	ExitOptional_parameter_specification(c *Optional_parameter_specificationContext)

	// ExitParameter_specification is called when exiting the parameter_specification production.
	ExitParameter_specification(c *Parameter_specificationContext)

	// ExitTable_type is called when exiting the table_type production.
	ExitTable_type(c *Table_typeContext)

	// ExitRow_type is called when exiting the row_type production.
	ExitRow_type(c *Row_typeContext)

	// ExitNullable_type is called when exiting the nullable_type production.
	ExitNullable_type(c *Nullable_typeContext)

	// ExitError_raising_expression is called when exiting the error_raising_expression production.
	ExitError_raising_expression(c *Error_raising_expressionContext)

	// ExitError_handling_expression is called when exiting the error_handling_expression production.
	ExitError_handling_expression(c *Error_handling_expressionContext)

	// ExitProtected_expression is called when exiting the protected_expression production.
	ExitProtected_expression(c *Protected_expressionContext)

	// ExitOtherwise_clause is called when exiting the otherwise_clause production.
	ExitOtherwise_clause(c *Otherwise_clauseContext)

	// ExitDefault_expression is called when exiting the default_expression production.
	ExitDefault_expression(c *Default_expressionContext)

	// ExitLiteral_attribs is called when exiting the literal_attribs production.
	ExitLiteral_attribs(c *Literal_attribsContext)

	// ExitRecord_literal is called when exiting the record_literal production.
	ExitRecord_literal(c *Record_literalContext)

	// ExitLiteral_field_list is called when exiting the literal_field_list production.
	ExitLiteral_field_list(c *Literal_field_listContext)

	// ExitLiteral_field is called when exiting the literal_field production.
	ExitLiteral_field(c *Literal_fieldContext)

	// ExitList_literal is called when exiting the list_literal production.
	ExitList_literal(c *List_literalContext)

	// ExitLiteral_item_list is called when exiting the literal_item_list production.
	ExitLiteral_item_list(c *Literal_item_listContext)

	// ExitAny_literal is called when exiting the any_literal production.
	ExitAny_literal(c *Any_literalContext)
}
