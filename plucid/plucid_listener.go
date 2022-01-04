// Code generated from plucid.g4 by ANTLR 4.9.3. DO NOT EDIT.

package plucid // plucid
import "github.com/antlr/antlr4/runtime/Go/antlr"

// plucidListener is a complete listener for a parse tree produced by plucidParser.
type plucidListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumeric_constant is called when entering the numeric_constant production.
	EnterNumeric_constant(c *Numeric_constantContext)

	// EnterInteger_constant is called when entering the integer_constant production.
	EnterInteger_constant(c *Integer_constantContext)

	// EnterReal_constant is called when entering the real_constant production.
	EnterReal_constant(c *Real_constantContext)

	// EnterWord_constant is called when entering the word_constant production.
	EnterWord_constant(c *Word_constantContext)

	// EnterWord_constant_less_the_quotes is called when entering the word_constant_less_the_quotes production.
	EnterWord_constant_less_the_quotes(c *Word_constant_less_the_quotesContext)

	// EnterList_constant is called when entering the list_constant production.
	EnterList_constant(c *List_constantContext)

	// EnterList_constant_element is called when entering the list_constant_element production.
	EnterList_constant_element(c *List_constant_elementContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterPrefix_operator is called when entering the prefix_operator production.
	EnterPrefix_operator(c *Prefix_operatorContext)

	// EnterInfix_operator is called when entering the infix_operator production.
	EnterInfix_operator(c *Infix_operatorContext)

	// EnterList_expression is called when entering the list_expression production.
	EnterList_expression(c *List_expressionContext)

	// EnterExpressions_list is called when entering the expressions_list production.
	EnterExpressions_list(c *Expressions_listContext)

	// EnterExpression_item is called when entering the expression_item production.
	EnterExpression_item(c *Expression_itemContext)

	// EnterIf_expression is called when entering the if_expression production.
	EnterIf_expression(c *If_expressionContext)

	// EnterEndif is called when entering the endif production.
	EnterEndif(c *EndifContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterCond_expression is called when entering the cond_expression production.
	EnterCond_expression(c *Cond_expressionContext)

	// EnterCbody is called when entering the cbody production.
	EnterCbody(c *CbodyContext)

	// EnterDefaultcase is called when entering the defaultcase production.
	EnterDefaultcase(c *DefaultcaseContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterActuals_list is called when entering the actuals_list production.
	EnterActuals_list(c *Actuals_listContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterDeclarations_list is called when entering the declarations_list production.
	EnterDeclarations_list(c *Declarations_listContext)

	// EnterCurrent_declaration is called when entering the current_declaration production.
	EnterCurrent_declaration(c *Current_declarationContext)

	// EnterDefinitions_list is called when entering the definitions_list production.
	EnterDefinitions_list(c *Definitions_listContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterSimple_definition is called when entering the simple_definition production.
	EnterSimple_definition(c *Simple_definitionContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterFormals_list is called when entering the formals_list production.
	EnterFormals_list(c *Formals_listContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumeric_constant is called when exiting the numeric_constant production.
	ExitNumeric_constant(c *Numeric_constantContext)

	// ExitInteger_constant is called when exiting the integer_constant production.
	ExitInteger_constant(c *Integer_constantContext)

	// ExitReal_constant is called when exiting the real_constant production.
	ExitReal_constant(c *Real_constantContext)

	// ExitWord_constant is called when exiting the word_constant production.
	ExitWord_constant(c *Word_constantContext)

	// ExitWord_constant_less_the_quotes is called when exiting the word_constant_less_the_quotes production.
	ExitWord_constant_less_the_quotes(c *Word_constant_less_the_quotesContext)

	// ExitList_constant is called when exiting the list_constant production.
	ExitList_constant(c *List_constantContext)

	// ExitList_constant_element is called when exiting the list_constant_element production.
	ExitList_constant_element(c *List_constant_elementContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitPrefix_operator is called when exiting the prefix_operator production.
	ExitPrefix_operator(c *Prefix_operatorContext)

	// ExitInfix_operator is called when exiting the infix_operator production.
	ExitInfix_operator(c *Infix_operatorContext)

	// ExitList_expression is called when exiting the list_expression production.
	ExitList_expression(c *List_expressionContext)

	// ExitExpressions_list is called when exiting the expressions_list production.
	ExitExpressions_list(c *Expressions_listContext)

	// ExitExpression_item is called when exiting the expression_item production.
	ExitExpression_item(c *Expression_itemContext)

	// ExitIf_expression is called when exiting the if_expression production.
	ExitIf_expression(c *If_expressionContext)

	// ExitEndif is called when exiting the endif production.
	ExitEndif(c *EndifContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitCond_expression is called when exiting the cond_expression production.
	ExitCond_expression(c *Cond_expressionContext)

	// ExitCbody is called when exiting the cbody production.
	ExitCbody(c *CbodyContext)

	// ExitDefaultcase is called when exiting the defaultcase production.
	ExitDefaultcase(c *DefaultcaseContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitActuals_list is called when exiting the actuals_list production.
	ExitActuals_list(c *Actuals_listContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitDeclarations_list is called when exiting the declarations_list production.
	ExitDeclarations_list(c *Declarations_listContext)

	// ExitCurrent_declaration is called when exiting the current_declaration production.
	ExitCurrent_declaration(c *Current_declarationContext)

	// ExitDefinitions_list is called when exiting the definitions_list production.
	ExitDefinitions_list(c *Definitions_listContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitSimple_definition is called when exiting the simple_definition production.
	ExitSimple_definition(c *Simple_definitionContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitFormals_list is called when exiting the formals_list production.
	ExitFormals_list(c *Formals_listContext)
}
