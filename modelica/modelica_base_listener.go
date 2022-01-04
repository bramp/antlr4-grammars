// Code generated from modelica.g4 by ANTLR 4.9.3. DO NOT EDIT.

package modelica // modelica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemodelicaListener is a complete listener for a parse tree produced by modelicaParser.
type BasemodelicaListener struct{}

var _ modelicaListener = &BasemodelicaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemodelicaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemodelicaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemodelicaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemodelicaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStored_definition is called when production stored_definition is entered.
func (s *BasemodelicaListener) EnterStored_definition(ctx *Stored_definitionContext) {}

// ExitStored_definition is called when production stored_definition is exited.
func (s *BasemodelicaListener) ExitStored_definition(ctx *Stored_definitionContext) {}

// EnterClass_definition is called when production class_definition is entered.
func (s *BasemodelicaListener) EnterClass_definition(ctx *Class_definitionContext) {}

// ExitClass_definition is called when production class_definition is exited.
func (s *BasemodelicaListener) ExitClass_definition(ctx *Class_definitionContext) {}

// EnterClass_specifier is called when production class_specifier is entered.
func (s *BasemodelicaListener) EnterClass_specifier(ctx *Class_specifierContext) {}

// ExitClass_specifier is called when production class_specifier is exited.
func (s *BasemodelicaListener) ExitClass_specifier(ctx *Class_specifierContext) {}

// EnterClass_prefixes is called when production class_prefixes is entered.
func (s *BasemodelicaListener) EnterClass_prefixes(ctx *Class_prefixesContext) {}

// ExitClass_prefixes is called when production class_prefixes is exited.
func (s *BasemodelicaListener) ExitClass_prefixes(ctx *Class_prefixesContext) {}

// EnterLong_class_specifier is called when production long_class_specifier is entered.
func (s *BasemodelicaListener) EnterLong_class_specifier(ctx *Long_class_specifierContext) {}

// ExitLong_class_specifier is called when production long_class_specifier is exited.
func (s *BasemodelicaListener) ExitLong_class_specifier(ctx *Long_class_specifierContext) {}

// EnterShort_class_specifier is called when production short_class_specifier is entered.
func (s *BasemodelicaListener) EnterShort_class_specifier(ctx *Short_class_specifierContext) {}

// ExitShort_class_specifier is called when production short_class_specifier is exited.
func (s *BasemodelicaListener) ExitShort_class_specifier(ctx *Short_class_specifierContext) {}

// EnterDer_class_specifier is called when production der_class_specifier is entered.
func (s *BasemodelicaListener) EnterDer_class_specifier(ctx *Der_class_specifierContext) {}

// ExitDer_class_specifier is called when production der_class_specifier is exited.
func (s *BasemodelicaListener) ExitDer_class_specifier(ctx *Der_class_specifierContext) {}

// EnterBase_prefix is called when production base_prefix is entered.
func (s *BasemodelicaListener) EnterBase_prefix(ctx *Base_prefixContext) {}

// ExitBase_prefix is called when production base_prefix is exited.
func (s *BasemodelicaListener) ExitBase_prefix(ctx *Base_prefixContext) {}

// EnterEnum_list is called when production enum_list is entered.
func (s *BasemodelicaListener) EnterEnum_list(ctx *Enum_listContext) {}

// ExitEnum_list is called when production enum_list is exited.
func (s *BasemodelicaListener) ExitEnum_list(ctx *Enum_listContext) {}

// EnterEnumeration_literal is called when production enumeration_literal is entered.
func (s *BasemodelicaListener) EnterEnumeration_literal(ctx *Enumeration_literalContext) {}

// ExitEnumeration_literal is called when production enumeration_literal is exited.
func (s *BasemodelicaListener) ExitEnumeration_literal(ctx *Enumeration_literalContext) {}

// EnterComposition is called when production composition is entered.
func (s *BasemodelicaListener) EnterComposition(ctx *CompositionContext) {}

// ExitComposition is called when production composition is exited.
func (s *BasemodelicaListener) ExitComposition(ctx *CompositionContext) {}

// EnterLanguage_specification is called when production language_specification is entered.
func (s *BasemodelicaListener) EnterLanguage_specification(ctx *Language_specificationContext) {}

// ExitLanguage_specification is called when production language_specification is exited.
func (s *BasemodelicaListener) ExitLanguage_specification(ctx *Language_specificationContext) {}

// EnterExternal_function_call is called when production external_function_call is entered.
func (s *BasemodelicaListener) EnterExternal_function_call(ctx *External_function_callContext) {}

// ExitExternal_function_call is called when production external_function_call is exited.
func (s *BasemodelicaListener) ExitExternal_function_call(ctx *External_function_callContext) {}

// EnterElement_list is called when production element_list is entered.
func (s *BasemodelicaListener) EnterElement_list(ctx *Element_listContext) {}

// ExitElement_list is called when production element_list is exited.
func (s *BasemodelicaListener) ExitElement_list(ctx *Element_listContext) {}

// EnterElement is called when production element is entered.
func (s *BasemodelicaListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasemodelicaListener) ExitElement(ctx *ElementContext) {}

// EnterImport_clause is called when production import_clause is entered.
func (s *BasemodelicaListener) EnterImport_clause(ctx *Import_clauseContext) {}

// ExitImport_clause is called when production import_clause is exited.
func (s *BasemodelicaListener) ExitImport_clause(ctx *Import_clauseContext) {}

// EnterImport_list is called when production import_list is entered.
func (s *BasemodelicaListener) EnterImport_list(ctx *Import_listContext) {}

// ExitImport_list is called when production import_list is exited.
func (s *BasemodelicaListener) ExitImport_list(ctx *Import_listContext) {}

// EnterExtends_clause is called when production extends_clause is entered.
func (s *BasemodelicaListener) EnterExtends_clause(ctx *Extends_clauseContext) {}

// ExitExtends_clause is called when production extends_clause is exited.
func (s *BasemodelicaListener) ExitExtends_clause(ctx *Extends_clauseContext) {}

// EnterConstraining_clause is called when production constraining_clause is entered.
func (s *BasemodelicaListener) EnterConstraining_clause(ctx *Constraining_clauseContext) {}

// ExitConstraining_clause is called when production constraining_clause is exited.
func (s *BasemodelicaListener) ExitConstraining_clause(ctx *Constraining_clauseContext) {}

// EnterComponent_clause is called when production component_clause is entered.
func (s *BasemodelicaListener) EnterComponent_clause(ctx *Component_clauseContext) {}

// ExitComponent_clause is called when production component_clause is exited.
func (s *BasemodelicaListener) ExitComponent_clause(ctx *Component_clauseContext) {}

// EnterType_prefix is called when production type_prefix is entered.
func (s *BasemodelicaListener) EnterType_prefix(ctx *Type_prefixContext) {}

// ExitType_prefix is called when production type_prefix is exited.
func (s *BasemodelicaListener) ExitType_prefix(ctx *Type_prefixContext) {}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BasemodelicaListener) EnterType_specifier(ctx *Type_specifierContext) {}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BasemodelicaListener) ExitType_specifier(ctx *Type_specifierContext) {}

// EnterComponent_list is called when production component_list is entered.
func (s *BasemodelicaListener) EnterComponent_list(ctx *Component_listContext) {}

// ExitComponent_list is called when production component_list is exited.
func (s *BasemodelicaListener) ExitComponent_list(ctx *Component_listContext) {}

// EnterComponent_declaration is called when production component_declaration is entered.
func (s *BasemodelicaListener) EnterComponent_declaration(ctx *Component_declarationContext) {}

// ExitComponent_declaration is called when production component_declaration is exited.
func (s *BasemodelicaListener) ExitComponent_declaration(ctx *Component_declarationContext) {}

// EnterCondition_attribute is called when production condition_attribute is entered.
func (s *BasemodelicaListener) EnterCondition_attribute(ctx *Condition_attributeContext) {}

// ExitCondition_attribute is called when production condition_attribute is exited.
func (s *BasemodelicaListener) ExitCondition_attribute(ctx *Condition_attributeContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasemodelicaListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasemodelicaListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterModification is called when production modification is entered.
func (s *BasemodelicaListener) EnterModification(ctx *ModificationContext) {}

// ExitModification is called when production modification is exited.
func (s *BasemodelicaListener) ExitModification(ctx *ModificationContext) {}

// EnterClass_modification is called when production class_modification is entered.
func (s *BasemodelicaListener) EnterClass_modification(ctx *Class_modificationContext) {}

// ExitClass_modification is called when production class_modification is exited.
func (s *BasemodelicaListener) ExitClass_modification(ctx *Class_modificationContext) {}

// EnterArgument_list is called when production argument_list is entered.
func (s *BasemodelicaListener) EnterArgument_list(ctx *Argument_listContext) {}

// ExitArgument_list is called when production argument_list is exited.
func (s *BasemodelicaListener) ExitArgument_list(ctx *Argument_listContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasemodelicaListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasemodelicaListener) ExitArgument(ctx *ArgumentContext) {}

// EnterElement_modification_or_replaceable is called when production element_modification_or_replaceable is entered.
func (s *BasemodelicaListener) EnterElement_modification_or_replaceable(ctx *Element_modification_or_replaceableContext) {
}

// ExitElement_modification_or_replaceable is called when production element_modification_or_replaceable is exited.
func (s *BasemodelicaListener) ExitElement_modification_or_replaceable(ctx *Element_modification_or_replaceableContext) {
}

// EnterElement_modification is called when production element_modification is entered.
func (s *BasemodelicaListener) EnterElement_modification(ctx *Element_modificationContext) {}

// ExitElement_modification is called when production element_modification is exited.
func (s *BasemodelicaListener) ExitElement_modification(ctx *Element_modificationContext) {}

// EnterElement_redeclaration is called when production element_redeclaration is entered.
func (s *BasemodelicaListener) EnterElement_redeclaration(ctx *Element_redeclarationContext) {}

// ExitElement_redeclaration is called when production element_redeclaration is exited.
func (s *BasemodelicaListener) ExitElement_redeclaration(ctx *Element_redeclarationContext) {}

// EnterElement_replaceable is called when production element_replaceable is entered.
func (s *BasemodelicaListener) EnterElement_replaceable(ctx *Element_replaceableContext) {}

// ExitElement_replaceable is called when production element_replaceable is exited.
func (s *BasemodelicaListener) ExitElement_replaceable(ctx *Element_replaceableContext) {}

// EnterComponent_clause1 is called when production component_clause1 is entered.
func (s *BasemodelicaListener) EnterComponent_clause1(ctx *Component_clause1Context) {}

// ExitComponent_clause1 is called when production component_clause1 is exited.
func (s *BasemodelicaListener) ExitComponent_clause1(ctx *Component_clause1Context) {}

// EnterComponent_declaration1 is called when production component_declaration1 is entered.
func (s *BasemodelicaListener) EnterComponent_declaration1(ctx *Component_declaration1Context) {}

// ExitComponent_declaration1 is called when production component_declaration1 is exited.
func (s *BasemodelicaListener) ExitComponent_declaration1(ctx *Component_declaration1Context) {}

// EnterShort_class_definition is called when production short_class_definition is entered.
func (s *BasemodelicaListener) EnterShort_class_definition(ctx *Short_class_definitionContext) {}

// ExitShort_class_definition is called when production short_class_definition is exited.
func (s *BasemodelicaListener) ExitShort_class_definition(ctx *Short_class_definitionContext) {}

// EnterEquation_section is called when production equation_section is entered.
func (s *BasemodelicaListener) EnterEquation_section(ctx *Equation_sectionContext) {}

// ExitEquation_section is called when production equation_section is exited.
func (s *BasemodelicaListener) ExitEquation_section(ctx *Equation_sectionContext) {}

// EnterAlgorithm_section is called when production algorithm_section is entered.
func (s *BasemodelicaListener) EnterAlgorithm_section(ctx *Algorithm_sectionContext) {}

// ExitAlgorithm_section is called when production algorithm_section is exited.
func (s *BasemodelicaListener) ExitAlgorithm_section(ctx *Algorithm_sectionContext) {}

// EnterEquation is called when production equation is entered.
func (s *BasemodelicaListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BasemodelicaListener) ExitEquation(ctx *EquationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasemodelicaListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasemodelicaListener) ExitStatement(ctx *StatementContext) {}

// EnterIf_equation is called when production if_equation is entered.
func (s *BasemodelicaListener) EnterIf_equation(ctx *If_equationContext) {}

// ExitIf_equation is called when production if_equation is exited.
func (s *BasemodelicaListener) ExitIf_equation(ctx *If_equationContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasemodelicaListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasemodelicaListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterFor_equation is called when production for_equation is entered.
func (s *BasemodelicaListener) EnterFor_equation(ctx *For_equationContext) {}

// ExitFor_equation is called when production for_equation is exited.
func (s *BasemodelicaListener) ExitFor_equation(ctx *For_equationContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BasemodelicaListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BasemodelicaListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterFor_indices is called when production for_indices is entered.
func (s *BasemodelicaListener) EnterFor_indices(ctx *For_indicesContext) {}

// ExitFor_indices is called when production for_indices is exited.
func (s *BasemodelicaListener) ExitFor_indices(ctx *For_indicesContext) {}

// EnterFor_index is called when production for_index is entered.
func (s *BasemodelicaListener) EnterFor_index(ctx *For_indexContext) {}

// ExitFor_index is called when production for_index is exited.
func (s *BasemodelicaListener) ExitFor_index(ctx *For_indexContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BasemodelicaListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BasemodelicaListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterWhen_equation is called when production when_equation is entered.
func (s *BasemodelicaListener) EnterWhen_equation(ctx *When_equationContext) {}

// ExitWhen_equation is called when production when_equation is exited.
func (s *BasemodelicaListener) ExitWhen_equation(ctx *When_equationContext) {}

// EnterWhen_statement is called when production when_statement is entered.
func (s *BasemodelicaListener) EnterWhen_statement(ctx *When_statementContext) {}

// ExitWhen_statement is called when production when_statement is exited.
func (s *BasemodelicaListener) ExitWhen_statement(ctx *When_statementContext) {}

// EnterConnect_clause is called when production connect_clause is entered.
func (s *BasemodelicaListener) EnterConnect_clause(ctx *Connect_clauseContext) {}

// ExitConnect_clause is called when production connect_clause is exited.
func (s *BasemodelicaListener) ExitConnect_clause(ctx *Connect_clauseContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasemodelicaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasemodelicaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSimple_expression is called when production simple_expression is entered.
func (s *BasemodelicaListener) EnterSimple_expression(ctx *Simple_expressionContext) {}

// ExitSimple_expression is called when production simple_expression is exited.
func (s *BasemodelicaListener) ExitSimple_expression(ctx *Simple_expressionContext) {}

// EnterLogical_expression is called when production logical_expression is entered.
func (s *BasemodelicaListener) EnterLogical_expression(ctx *Logical_expressionContext) {}

// ExitLogical_expression is called when production logical_expression is exited.
func (s *BasemodelicaListener) ExitLogical_expression(ctx *Logical_expressionContext) {}

// EnterLogical_term is called when production logical_term is entered.
func (s *BasemodelicaListener) EnterLogical_term(ctx *Logical_termContext) {}

// ExitLogical_term is called when production logical_term is exited.
func (s *BasemodelicaListener) ExitLogical_term(ctx *Logical_termContext) {}

// EnterLogical_factor is called when production logical_factor is entered.
func (s *BasemodelicaListener) EnterLogical_factor(ctx *Logical_factorContext) {}

// ExitLogical_factor is called when production logical_factor is exited.
func (s *BasemodelicaListener) ExitLogical_factor(ctx *Logical_factorContext) {}

// EnterRelation is called when production relation is entered.
func (s *BasemodelicaListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BasemodelicaListener) ExitRelation(ctx *RelationContext) {}

// EnterRel_op is called when production rel_op is entered.
func (s *BasemodelicaListener) EnterRel_op(ctx *Rel_opContext) {}

// ExitRel_op is called when production rel_op is exited.
func (s *BasemodelicaListener) ExitRel_op(ctx *Rel_opContext) {}

// EnterArithmetic_expression is called when production arithmetic_expression is entered.
func (s *BasemodelicaListener) EnterArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// ExitArithmetic_expression is called when production arithmetic_expression is exited.
func (s *BasemodelicaListener) ExitArithmetic_expression(ctx *Arithmetic_expressionContext) {}

// EnterAdd_op is called when production add_op is entered.
func (s *BasemodelicaListener) EnterAdd_op(ctx *Add_opContext) {}

// ExitAdd_op is called when production add_op is exited.
func (s *BasemodelicaListener) ExitAdd_op(ctx *Add_opContext) {}

// EnterTerm is called when production term is entered.
func (s *BasemodelicaListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasemodelicaListener) ExitTerm(ctx *TermContext) {}

// EnterMul_op is called when production mul_op is entered.
func (s *BasemodelicaListener) EnterMul_op(ctx *Mul_opContext) {}

// ExitMul_op is called when production mul_op is exited.
func (s *BasemodelicaListener) ExitMul_op(ctx *Mul_opContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasemodelicaListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasemodelicaListener) ExitFactor(ctx *FactorContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasemodelicaListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasemodelicaListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterName is called when production name is entered.
func (s *BasemodelicaListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasemodelicaListener) ExitName(ctx *NameContext) {}

// EnterComponent_reference is called when production component_reference is entered.
func (s *BasemodelicaListener) EnterComponent_reference(ctx *Component_referenceContext) {}

// ExitComponent_reference is called when production component_reference is exited.
func (s *BasemodelicaListener) ExitComponent_reference(ctx *Component_referenceContext) {}

// EnterFunction_call_args is called when production function_call_args is entered.
func (s *BasemodelicaListener) EnterFunction_call_args(ctx *Function_call_argsContext) {}

// ExitFunction_call_args is called when production function_call_args is exited.
func (s *BasemodelicaListener) ExitFunction_call_args(ctx *Function_call_argsContext) {}

// EnterFunction_arguments is called when production function_arguments is entered.
func (s *BasemodelicaListener) EnterFunction_arguments(ctx *Function_argumentsContext) {}

// ExitFunction_arguments is called when production function_arguments is exited.
func (s *BasemodelicaListener) ExitFunction_arguments(ctx *Function_argumentsContext) {}

// EnterNamed_arguments is called when production named_arguments is entered.
func (s *BasemodelicaListener) EnterNamed_arguments(ctx *Named_argumentsContext) {}

// ExitNamed_arguments is called when production named_arguments is exited.
func (s *BasemodelicaListener) ExitNamed_arguments(ctx *Named_argumentsContext) {}

// EnterNamed_argument is called when production named_argument is entered.
func (s *BasemodelicaListener) EnterNamed_argument(ctx *Named_argumentContext) {}

// ExitNamed_argument is called when production named_argument is exited.
func (s *BasemodelicaListener) ExitNamed_argument(ctx *Named_argumentContext) {}

// EnterFunction_argument is called when production function_argument is entered.
func (s *BasemodelicaListener) EnterFunction_argument(ctx *Function_argumentContext) {}

// ExitFunction_argument is called when production function_argument is exited.
func (s *BasemodelicaListener) ExitFunction_argument(ctx *Function_argumentContext) {}

// EnterOutput_expression_list is called when production output_expression_list is entered.
func (s *BasemodelicaListener) EnterOutput_expression_list(ctx *Output_expression_listContext) {}

// ExitOutput_expression_list is called when production output_expression_list is exited.
func (s *BasemodelicaListener) ExitOutput_expression_list(ctx *Output_expression_listContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasemodelicaListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasemodelicaListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterArray_subscripts is called when production array_subscripts is entered.
func (s *BasemodelicaListener) EnterArray_subscripts(ctx *Array_subscriptsContext) {}

// ExitArray_subscripts is called when production array_subscripts is exited.
func (s *BasemodelicaListener) ExitArray_subscripts(ctx *Array_subscriptsContext) {}

// EnterSubscript_ is called when production subscript_ is entered.
func (s *BasemodelicaListener) EnterSubscript_(ctx *Subscript_Context) {}

// ExitSubscript_ is called when production subscript_ is exited.
func (s *BasemodelicaListener) ExitSubscript_(ctx *Subscript_Context) {}

// EnterComment is called when production comment is entered.
func (s *BasemodelicaListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasemodelicaListener) ExitComment(ctx *CommentContext) {}

// EnterString_comment is called when production string_comment is entered.
func (s *BasemodelicaListener) EnterString_comment(ctx *String_commentContext) {}

// ExitString_comment is called when production string_comment is exited.
func (s *BasemodelicaListener) ExitString_comment(ctx *String_commentContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BasemodelicaListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BasemodelicaListener) ExitAnnotation(ctx *AnnotationContext) {}
