// Code generated from modelica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package modelica // modelica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// modelicaListener is a complete listener for a parse tree produced by modelicaParser.
type modelicaListener interface {
	antlr.ParseTreeListener

	// EnterStored_definition is called when entering the stored_definition production.
	EnterStored_definition(c *Stored_definitionContext)

	// EnterClass_definition is called when entering the class_definition production.
	EnterClass_definition(c *Class_definitionContext)

	// EnterClass_specifier is called when entering the class_specifier production.
	EnterClass_specifier(c *Class_specifierContext)

	// EnterClass_prefixes is called when entering the class_prefixes production.
	EnterClass_prefixes(c *Class_prefixesContext)

	// EnterLong_class_specifier is called when entering the long_class_specifier production.
	EnterLong_class_specifier(c *Long_class_specifierContext)

	// EnterShort_class_specifier is called when entering the short_class_specifier production.
	EnterShort_class_specifier(c *Short_class_specifierContext)

	// EnterDer_class_specifier is called when entering the der_class_specifier production.
	EnterDer_class_specifier(c *Der_class_specifierContext)

	// EnterBase_prefix is called when entering the base_prefix production.
	EnterBase_prefix(c *Base_prefixContext)

	// EnterEnum_list is called when entering the enum_list production.
	EnterEnum_list(c *Enum_listContext)

	// EnterEnumeration_literal is called when entering the enumeration_literal production.
	EnterEnumeration_literal(c *Enumeration_literalContext)

	// EnterComposition is called when entering the composition production.
	EnterComposition(c *CompositionContext)

	// EnterLanguage_specification is called when entering the language_specification production.
	EnterLanguage_specification(c *Language_specificationContext)

	// EnterExternal_function_call is called when entering the external_function_call production.
	EnterExternal_function_call(c *External_function_callContext)

	// EnterElement_list is called when entering the element_list production.
	EnterElement_list(c *Element_listContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterImport_clause is called when entering the import_clause production.
	EnterImport_clause(c *Import_clauseContext)

	// EnterImport_list is called when entering the import_list production.
	EnterImport_list(c *Import_listContext)

	// EnterExtends_clause is called when entering the extends_clause production.
	EnterExtends_clause(c *Extends_clauseContext)

	// EnterConstraining_clause is called when entering the constraining_clause production.
	EnterConstraining_clause(c *Constraining_clauseContext)

	// EnterComponent_clause is called when entering the component_clause production.
	EnterComponent_clause(c *Component_clauseContext)

	// EnterType_prefix is called when entering the type_prefix production.
	EnterType_prefix(c *Type_prefixContext)

	// EnterType_specifier is called when entering the type_specifier production.
	EnterType_specifier(c *Type_specifierContext)

	// EnterComponent_list is called when entering the component_list production.
	EnterComponent_list(c *Component_listContext)

	// EnterComponent_declaration is called when entering the component_declaration production.
	EnterComponent_declaration(c *Component_declarationContext)

	// EnterCondition_attribute is called when entering the condition_attribute production.
	EnterCondition_attribute(c *Condition_attributeContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterModification is called when entering the modification production.
	EnterModification(c *ModificationContext)

	// EnterClass_modification is called when entering the class_modification production.
	EnterClass_modification(c *Class_modificationContext)

	// EnterArgument_list is called when entering the argument_list production.
	EnterArgument_list(c *Argument_listContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterElement_modification_or_replaceable is called when entering the element_modification_or_replaceable production.
	EnterElement_modification_or_replaceable(c *Element_modification_or_replaceableContext)

	// EnterElement_modification is called when entering the element_modification production.
	EnterElement_modification(c *Element_modificationContext)

	// EnterElement_redeclaration is called when entering the element_redeclaration production.
	EnterElement_redeclaration(c *Element_redeclarationContext)

	// EnterElement_replaceable is called when entering the element_replaceable production.
	EnterElement_replaceable(c *Element_replaceableContext)

	// EnterComponent_clause1 is called when entering the component_clause1 production.
	EnterComponent_clause1(c *Component_clause1Context)

	// EnterComponent_declaration1 is called when entering the component_declaration1 production.
	EnterComponent_declaration1(c *Component_declaration1Context)

	// EnterShort_class_definition is called when entering the short_class_definition production.
	EnterShort_class_definition(c *Short_class_definitionContext)

	// EnterEquation_section is called when entering the equation_section production.
	EnterEquation_section(c *Equation_sectionContext)

	// EnterAlgorithm_section is called when entering the algorithm_section production.
	EnterAlgorithm_section(c *Algorithm_sectionContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIf_equation is called when entering the if_equation production.
	EnterIf_equation(c *If_equationContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterFor_equation is called when entering the for_equation production.
	EnterFor_equation(c *For_equationContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterFor_indices is called when entering the for_indices production.
	EnterFor_indices(c *For_indicesContext)

	// EnterFor_index is called when entering the for_index production.
	EnterFor_index(c *For_indexContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterWhen_equation is called when entering the when_equation production.
	EnterWhen_equation(c *When_equationContext)

	// EnterWhen_statement is called when entering the when_statement production.
	EnterWhen_statement(c *When_statementContext)

	// EnterConnect_clause is called when entering the connect_clause production.
	EnterConnect_clause(c *Connect_clauseContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSimple_expression is called when entering the simple_expression production.
	EnterSimple_expression(c *Simple_expressionContext)

	// EnterLogical_expression is called when entering the logical_expression production.
	EnterLogical_expression(c *Logical_expressionContext)

	// EnterLogical_term is called when entering the logical_term production.
	EnterLogical_term(c *Logical_termContext)

	// EnterLogical_factor is called when entering the logical_factor production.
	EnterLogical_factor(c *Logical_factorContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterRel_op is called when entering the rel_op production.
	EnterRel_op(c *Rel_opContext)

	// EnterArithmetic_expression is called when entering the arithmetic_expression production.
	EnterArithmetic_expression(c *Arithmetic_expressionContext)

	// EnterAdd_op is called when entering the add_op production.
	EnterAdd_op(c *Add_opContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMul_op is called when entering the mul_op production.
	EnterMul_op(c *Mul_opContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterComponent_reference is called when entering the component_reference production.
	EnterComponent_reference(c *Component_referenceContext)

	// EnterFunction_call_args is called when entering the function_call_args production.
	EnterFunction_call_args(c *Function_call_argsContext)

	// EnterFunction_arguments is called when entering the function_arguments production.
	EnterFunction_arguments(c *Function_argumentsContext)

	// EnterNamed_arguments is called when entering the named_arguments production.
	EnterNamed_arguments(c *Named_argumentsContext)

	// EnterNamed_argument is called when entering the named_argument production.
	EnterNamed_argument(c *Named_argumentContext)

	// EnterFunction_argument is called when entering the function_argument production.
	EnterFunction_argument(c *Function_argumentContext)

	// EnterOutput_expression_list is called when entering the output_expression_list production.
	EnterOutput_expression_list(c *Output_expression_listContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterArray_subscripts is called when entering the array_subscripts production.
	EnterArray_subscripts(c *Array_subscriptsContext)

	// EnterSubscript is called when entering the subscript production.
	EnterSubscript(c *SubscriptContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterString_comment is called when entering the string_comment production.
	EnterString_comment(c *String_commentContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// ExitStored_definition is called when exiting the stored_definition production.
	ExitStored_definition(c *Stored_definitionContext)

	// ExitClass_definition is called when exiting the class_definition production.
	ExitClass_definition(c *Class_definitionContext)

	// ExitClass_specifier is called when exiting the class_specifier production.
	ExitClass_specifier(c *Class_specifierContext)

	// ExitClass_prefixes is called when exiting the class_prefixes production.
	ExitClass_prefixes(c *Class_prefixesContext)

	// ExitLong_class_specifier is called when exiting the long_class_specifier production.
	ExitLong_class_specifier(c *Long_class_specifierContext)

	// ExitShort_class_specifier is called when exiting the short_class_specifier production.
	ExitShort_class_specifier(c *Short_class_specifierContext)

	// ExitDer_class_specifier is called when exiting the der_class_specifier production.
	ExitDer_class_specifier(c *Der_class_specifierContext)

	// ExitBase_prefix is called when exiting the base_prefix production.
	ExitBase_prefix(c *Base_prefixContext)

	// ExitEnum_list is called when exiting the enum_list production.
	ExitEnum_list(c *Enum_listContext)

	// ExitEnumeration_literal is called when exiting the enumeration_literal production.
	ExitEnumeration_literal(c *Enumeration_literalContext)

	// ExitComposition is called when exiting the composition production.
	ExitComposition(c *CompositionContext)

	// ExitLanguage_specification is called when exiting the language_specification production.
	ExitLanguage_specification(c *Language_specificationContext)

	// ExitExternal_function_call is called when exiting the external_function_call production.
	ExitExternal_function_call(c *External_function_callContext)

	// ExitElement_list is called when exiting the element_list production.
	ExitElement_list(c *Element_listContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitImport_clause is called when exiting the import_clause production.
	ExitImport_clause(c *Import_clauseContext)

	// ExitImport_list is called when exiting the import_list production.
	ExitImport_list(c *Import_listContext)

	// ExitExtends_clause is called when exiting the extends_clause production.
	ExitExtends_clause(c *Extends_clauseContext)

	// ExitConstraining_clause is called when exiting the constraining_clause production.
	ExitConstraining_clause(c *Constraining_clauseContext)

	// ExitComponent_clause is called when exiting the component_clause production.
	ExitComponent_clause(c *Component_clauseContext)

	// ExitType_prefix is called when exiting the type_prefix production.
	ExitType_prefix(c *Type_prefixContext)

	// ExitType_specifier is called when exiting the type_specifier production.
	ExitType_specifier(c *Type_specifierContext)

	// ExitComponent_list is called when exiting the component_list production.
	ExitComponent_list(c *Component_listContext)

	// ExitComponent_declaration is called when exiting the component_declaration production.
	ExitComponent_declaration(c *Component_declarationContext)

	// ExitCondition_attribute is called when exiting the condition_attribute production.
	ExitCondition_attribute(c *Condition_attributeContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitModification is called when exiting the modification production.
	ExitModification(c *ModificationContext)

	// ExitClass_modification is called when exiting the class_modification production.
	ExitClass_modification(c *Class_modificationContext)

	// ExitArgument_list is called when exiting the argument_list production.
	ExitArgument_list(c *Argument_listContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitElement_modification_or_replaceable is called when exiting the element_modification_or_replaceable production.
	ExitElement_modification_or_replaceable(c *Element_modification_or_replaceableContext)

	// ExitElement_modification is called when exiting the element_modification production.
	ExitElement_modification(c *Element_modificationContext)

	// ExitElement_redeclaration is called when exiting the element_redeclaration production.
	ExitElement_redeclaration(c *Element_redeclarationContext)

	// ExitElement_replaceable is called when exiting the element_replaceable production.
	ExitElement_replaceable(c *Element_replaceableContext)

	// ExitComponent_clause1 is called when exiting the component_clause1 production.
	ExitComponent_clause1(c *Component_clause1Context)

	// ExitComponent_declaration1 is called when exiting the component_declaration1 production.
	ExitComponent_declaration1(c *Component_declaration1Context)

	// ExitShort_class_definition is called when exiting the short_class_definition production.
	ExitShort_class_definition(c *Short_class_definitionContext)

	// ExitEquation_section is called when exiting the equation_section production.
	ExitEquation_section(c *Equation_sectionContext)

	// ExitAlgorithm_section is called when exiting the algorithm_section production.
	ExitAlgorithm_section(c *Algorithm_sectionContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIf_equation is called when exiting the if_equation production.
	ExitIf_equation(c *If_equationContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitFor_equation is called when exiting the for_equation production.
	ExitFor_equation(c *For_equationContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitFor_indices is called when exiting the for_indices production.
	ExitFor_indices(c *For_indicesContext)

	// ExitFor_index is called when exiting the for_index production.
	ExitFor_index(c *For_indexContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitWhen_equation is called when exiting the when_equation production.
	ExitWhen_equation(c *When_equationContext)

	// ExitWhen_statement is called when exiting the when_statement production.
	ExitWhen_statement(c *When_statementContext)

	// ExitConnect_clause is called when exiting the connect_clause production.
	ExitConnect_clause(c *Connect_clauseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSimple_expression is called when exiting the simple_expression production.
	ExitSimple_expression(c *Simple_expressionContext)

	// ExitLogical_expression is called when exiting the logical_expression production.
	ExitLogical_expression(c *Logical_expressionContext)

	// ExitLogical_term is called when exiting the logical_term production.
	ExitLogical_term(c *Logical_termContext)

	// ExitLogical_factor is called when exiting the logical_factor production.
	ExitLogical_factor(c *Logical_factorContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitRel_op is called when exiting the rel_op production.
	ExitRel_op(c *Rel_opContext)

	// ExitArithmetic_expression is called when exiting the arithmetic_expression production.
	ExitArithmetic_expression(c *Arithmetic_expressionContext)

	// ExitAdd_op is called when exiting the add_op production.
	ExitAdd_op(c *Add_opContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMul_op is called when exiting the mul_op production.
	ExitMul_op(c *Mul_opContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitComponent_reference is called when exiting the component_reference production.
	ExitComponent_reference(c *Component_referenceContext)

	// ExitFunction_call_args is called when exiting the function_call_args production.
	ExitFunction_call_args(c *Function_call_argsContext)

	// ExitFunction_arguments is called when exiting the function_arguments production.
	ExitFunction_arguments(c *Function_argumentsContext)

	// ExitNamed_arguments is called when exiting the named_arguments production.
	ExitNamed_arguments(c *Named_argumentsContext)

	// ExitNamed_argument is called when exiting the named_argument production.
	ExitNamed_argument(c *Named_argumentContext)

	// ExitFunction_argument is called when exiting the function_argument production.
	ExitFunction_argument(c *Function_argumentContext)

	// ExitOutput_expression_list is called when exiting the output_expression_list production.
	ExitOutput_expression_list(c *Output_expression_listContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitArray_subscripts is called when exiting the array_subscripts production.
	ExitArray_subscripts(c *Array_subscriptsContext)

	// ExitSubscript is called when exiting the subscript production.
	ExitSubscript(c *SubscriptContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitString_comment is called when exiting the string_comment production.
	ExitString_comment(c *String_commentContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)
}
