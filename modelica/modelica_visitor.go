// Generated from modelica.g4 by ANTLR 4.7.

package modelica // modelica
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by modelicaParser.
type modelicaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by modelicaParser#stored_definition.
	VisitStored_definition(ctx *Stored_definitionContext) interface{}

	// Visit a parse tree produced by modelicaParser#class_definition.
	VisitClass_definition(ctx *Class_definitionContext) interface{}

	// Visit a parse tree produced by modelicaParser#class_specifier.
	VisitClass_specifier(ctx *Class_specifierContext) interface{}

	// Visit a parse tree produced by modelicaParser#class_prefixes.
	VisitClass_prefixes(ctx *Class_prefixesContext) interface{}

	// Visit a parse tree produced by modelicaParser#long_class_specifier.
	VisitLong_class_specifier(ctx *Long_class_specifierContext) interface{}

	// Visit a parse tree produced by modelicaParser#short_class_specifier.
	VisitShort_class_specifier(ctx *Short_class_specifierContext) interface{}

	// Visit a parse tree produced by modelicaParser#der_class_specifier.
	VisitDer_class_specifier(ctx *Der_class_specifierContext) interface{}

	// Visit a parse tree produced by modelicaParser#base_prefix.
	VisitBase_prefix(ctx *Base_prefixContext) interface{}

	// Visit a parse tree produced by modelicaParser#enum_list.
	VisitEnum_list(ctx *Enum_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#enumeration_literal.
	VisitEnumeration_literal(ctx *Enumeration_literalContext) interface{}

	// Visit a parse tree produced by modelicaParser#composition.
	VisitComposition(ctx *CompositionContext) interface{}

	// Visit a parse tree produced by modelicaParser#language_specification.
	VisitLanguage_specification(ctx *Language_specificationContext) interface{}

	// Visit a parse tree produced by modelicaParser#external_function_call.
	VisitExternal_function_call(ctx *External_function_callContext) interface{}

	// Visit a parse tree produced by modelicaParser#element_list.
	VisitElement_list(ctx *Element_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by modelicaParser#import_clause.
	VisitImport_clause(ctx *Import_clauseContext) interface{}

	// Visit a parse tree produced by modelicaParser#import_list.
	VisitImport_list(ctx *Import_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#extends_clause.
	VisitExtends_clause(ctx *Extends_clauseContext) interface{}

	// Visit a parse tree produced by modelicaParser#constraining_clause.
	VisitConstraining_clause(ctx *Constraining_clauseContext) interface{}

	// Visit a parse tree produced by modelicaParser#component_clause.
	VisitComponent_clause(ctx *Component_clauseContext) interface{}

	// Visit a parse tree produced by modelicaParser#type_prefix.
	VisitType_prefix(ctx *Type_prefixContext) interface{}

	// Visit a parse tree produced by modelicaParser#type_specifier.
	VisitType_specifier(ctx *Type_specifierContext) interface{}

	// Visit a parse tree produced by modelicaParser#component_list.
	VisitComponent_list(ctx *Component_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#component_declaration.
	VisitComponent_declaration(ctx *Component_declarationContext) interface{}

	// Visit a parse tree produced by modelicaParser#condition_attribute.
	VisitCondition_attribute(ctx *Condition_attributeContext) interface{}

	// Visit a parse tree produced by modelicaParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by modelicaParser#modification.
	VisitModification(ctx *ModificationContext) interface{}

	// Visit a parse tree produced by modelicaParser#class_modification.
	VisitClass_modification(ctx *Class_modificationContext) interface{}

	// Visit a parse tree produced by modelicaParser#argument_list.
	VisitArgument_list(ctx *Argument_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by modelicaParser#element_modification_or_replaceable.
	VisitElement_modification_or_replaceable(ctx *Element_modification_or_replaceableContext) interface{}

	// Visit a parse tree produced by modelicaParser#element_modification.
	VisitElement_modification(ctx *Element_modificationContext) interface{}

	// Visit a parse tree produced by modelicaParser#element_redeclaration.
	VisitElement_redeclaration(ctx *Element_redeclarationContext) interface{}

	// Visit a parse tree produced by modelicaParser#element_replaceable.
	VisitElement_replaceable(ctx *Element_replaceableContext) interface{}

	// Visit a parse tree produced by modelicaParser#component_clause1.
	VisitComponent_clause1(ctx *Component_clause1Context) interface{}

	// Visit a parse tree produced by modelicaParser#component_declaration1.
	VisitComponent_declaration1(ctx *Component_declaration1Context) interface{}

	// Visit a parse tree produced by modelicaParser#short_class_definition.
	VisitShort_class_definition(ctx *Short_class_definitionContext) interface{}

	// Visit a parse tree produced by modelicaParser#equation_section.
	VisitEquation_section(ctx *Equation_sectionContext) interface{}

	// Visit a parse tree produced by modelicaParser#algorithm_section.
	VisitAlgorithm_section(ctx *Algorithm_sectionContext) interface{}

	// Visit a parse tree produced by modelicaParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by modelicaParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by modelicaParser#if_equation.
	VisitIf_equation(ctx *If_equationContext) interface{}

	// Visit a parse tree produced by modelicaParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by modelicaParser#for_equation.
	VisitFor_equation(ctx *For_equationContext) interface{}

	// Visit a parse tree produced by modelicaParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by modelicaParser#for_indices.
	VisitFor_indices(ctx *For_indicesContext) interface{}

	// Visit a parse tree produced by modelicaParser#for_index.
	VisitFor_index(ctx *For_indexContext) interface{}

	// Visit a parse tree produced by modelicaParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by modelicaParser#when_equation.
	VisitWhen_equation(ctx *When_equationContext) interface{}

	// Visit a parse tree produced by modelicaParser#when_statement.
	VisitWhen_statement(ctx *When_statementContext) interface{}

	// Visit a parse tree produced by modelicaParser#connect_clause.
	VisitConnect_clause(ctx *Connect_clauseContext) interface{}

	// Visit a parse tree produced by modelicaParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by modelicaParser#simple_expression.
	VisitSimple_expression(ctx *Simple_expressionContext) interface{}

	// Visit a parse tree produced by modelicaParser#logical_expression.
	VisitLogical_expression(ctx *Logical_expressionContext) interface{}

	// Visit a parse tree produced by modelicaParser#logical_term.
	VisitLogical_term(ctx *Logical_termContext) interface{}

	// Visit a parse tree produced by modelicaParser#logical_factor.
	VisitLogical_factor(ctx *Logical_factorContext) interface{}

	// Visit a parse tree produced by modelicaParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by modelicaParser#rel_op.
	VisitRel_op(ctx *Rel_opContext) interface{}

	// Visit a parse tree produced by modelicaParser#arithmetic_expression.
	VisitArithmetic_expression(ctx *Arithmetic_expressionContext) interface{}

	// Visit a parse tree produced by modelicaParser#add_op.
	VisitAdd_op(ctx *Add_opContext) interface{}

	// Visit a parse tree produced by modelicaParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by modelicaParser#mul_op.
	VisitMul_op(ctx *Mul_opContext) interface{}

	// Visit a parse tree produced by modelicaParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by modelicaParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by modelicaParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by modelicaParser#component_reference.
	VisitComponent_reference(ctx *Component_referenceContext) interface{}

	// Visit a parse tree produced by modelicaParser#function_call_args.
	VisitFunction_call_args(ctx *Function_call_argsContext) interface{}

	// Visit a parse tree produced by modelicaParser#function_arguments.
	VisitFunction_arguments(ctx *Function_argumentsContext) interface{}

	// Visit a parse tree produced by modelicaParser#named_arguments.
	VisitNamed_arguments(ctx *Named_argumentsContext) interface{}

	// Visit a parse tree produced by modelicaParser#named_argument.
	VisitNamed_argument(ctx *Named_argumentContext) interface{}

	// Visit a parse tree produced by modelicaParser#function_argument.
	VisitFunction_argument(ctx *Function_argumentContext) interface{}

	// Visit a parse tree produced by modelicaParser#output_expression_list.
	VisitOutput_expression_list(ctx *Output_expression_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by modelicaParser#array_subscripts.
	VisitArray_subscripts(ctx *Array_subscriptsContext) interface{}

	// Visit a parse tree produced by modelicaParser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by modelicaParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by modelicaParser#string_comment.
	VisitString_comment(ctx *String_commentContext) interface{}

	// Visit a parse tree produced by modelicaParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}
}
