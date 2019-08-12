// Code generated from matlab.g4 by ANTLR 4.7.2. DO NOT EDIT.

package matlab // matlab
import "github.com/antlr/antlr4/runtime/Go/antlr"

// matlabListener is a complete listener for a parse tree produced by matlabParser.
type matlabListener interface {
	antlr.ParseTreeListener

	// EnterPrimary_expression is called when entering the primary_expression production.
	EnterPrimary_expression(c *Primary_expressionContext)

	// EnterPostfix_expression is called when entering the postfix_expression production.
	EnterPostfix_expression(c *Postfix_expressionContext)

	// EnterIndex_expression is called when entering the index_expression production.
	EnterIndex_expression(c *Index_expressionContext)

	// EnterIndex_expression_list is called when entering the index_expression_list production.
	EnterIndex_expression_list(c *Index_expression_listContext)

	// EnterArray_expression is called when entering the array_expression production.
	EnterArray_expression(c *Array_expressionContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterMultiplicative_expression is called when entering the multiplicative_expression production.
	EnterMultiplicative_expression(c *Multiplicative_expressionContext)

	// EnterAdditive_expression is called when entering the additive_expression production.
	EnterAdditive_expression(c *Additive_expressionContext)

	// EnterRelational_expression is called when entering the relational_expression production.
	EnterRelational_expression(c *Relational_expressionContext)

	// EnterEquality_expression is called when entering the equality_expression production.
	EnterEquality_expression(c *Equality_expressionContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterOr_expression is called when entering the or_expression production.
	EnterOr_expression(c *Or_expressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignment_expression is called when entering the assignment_expression production.
	EnterAssignment_expression(c *Assignment_expressionContext)

	// EnterEostmt is called when entering the eostmt production.
	EnterEostmt(c *EostmtContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterGlobal_statement is called when entering the global_statement production.
	EnterGlobal_statement(c *Global_statementContext)

	// EnterClear_statement is called when entering the clear_statement production.
	EnterClear_statement(c *Clear_statementContext)

	// EnterExpression_statement is called when entering the expression_statement production.
	EnterExpression_statement(c *Expression_statementContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterArray_element is called when entering the array_element production.
	EnterArray_element(c *Array_elementContext)

	// EnterArray_list is called when entering the array_list production.
	EnterArray_list(c *Array_listContext)

	// EnterSelection_statement is called when entering the selection_statement production.
	EnterSelection_statement(c *Selection_statementContext)

	// EnterElseif_clause is called when entering the elseif_clause production.
	EnterElseif_clause(c *Elseif_clauseContext)

	// EnterIteration_statement is called when entering the iteration_statement production.
	EnterIteration_statement(c *Iteration_statementContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterTranslation_unit is called when entering the translation_unit production.
	EnterTranslation_unit(c *Translation_unitContext)

	// EnterFunc_ident_list is called when entering the func_ident_list production.
	EnterFunc_ident_list(c *Func_ident_listContext)

	// EnterFunc_return_list is called when entering the func_return_list production.
	EnterFunc_return_list(c *Func_return_listContext)

	// EnterFunction_declare_lhs is called when entering the function_declare_lhs production.
	EnterFunction_declare_lhs(c *Function_declare_lhsContext)

	// EnterFunction_declare is called when entering the function_declare production.
	EnterFunction_declare(c *Function_declareContext)

	// ExitPrimary_expression is called when exiting the primary_expression production.
	ExitPrimary_expression(c *Primary_expressionContext)

	// ExitPostfix_expression is called when exiting the postfix_expression production.
	ExitPostfix_expression(c *Postfix_expressionContext)

	// ExitIndex_expression is called when exiting the index_expression production.
	ExitIndex_expression(c *Index_expressionContext)

	// ExitIndex_expression_list is called when exiting the index_expression_list production.
	ExitIndex_expression_list(c *Index_expression_listContext)

	// ExitArray_expression is called when exiting the array_expression production.
	ExitArray_expression(c *Array_expressionContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitMultiplicative_expression is called when exiting the multiplicative_expression production.
	ExitMultiplicative_expression(c *Multiplicative_expressionContext)

	// ExitAdditive_expression is called when exiting the additive_expression production.
	ExitAdditive_expression(c *Additive_expressionContext)

	// ExitRelational_expression is called when exiting the relational_expression production.
	ExitRelational_expression(c *Relational_expressionContext)

	// ExitEquality_expression is called when exiting the equality_expression production.
	ExitEquality_expression(c *Equality_expressionContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitOr_expression is called when exiting the or_expression production.
	ExitOr_expression(c *Or_expressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignment_expression is called when exiting the assignment_expression production.
	ExitAssignment_expression(c *Assignment_expressionContext)

	// ExitEostmt is called when exiting the eostmt production.
	ExitEostmt(c *EostmtContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitGlobal_statement is called when exiting the global_statement production.
	ExitGlobal_statement(c *Global_statementContext)

	// ExitClear_statement is called when exiting the clear_statement production.
	ExitClear_statement(c *Clear_statementContext)

	// ExitExpression_statement is called when exiting the expression_statement production.
	ExitExpression_statement(c *Expression_statementContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitArray_element is called when exiting the array_element production.
	ExitArray_element(c *Array_elementContext)

	// ExitArray_list is called when exiting the array_list production.
	ExitArray_list(c *Array_listContext)

	// ExitSelection_statement is called when exiting the selection_statement production.
	ExitSelection_statement(c *Selection_statementContext)

	// ExitElseif_clause is called when exiting the elseif_clause production.
	ExitElseif_clause(c *Elseif_clauseContext)

	// ExitIteration_statement is called when exiting the iteration_statement production.
	ExitIteration_statement(c *Iteration_statementContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitTranslation_unit is called when exiting the translation_unit production.
	ExitTranslation_unit(c *Translation_unitContext)

	// ExitFunc_ident_list is called when exiting the func_ident_list production.
	ExitFunc_ident_list(c *Func_ident_listContext)

	// ExitFunc_return_list is called when exiting the func_return_list production.
	ExitFunc_return_list(c *Func_return_listContext)

	// ExitFunction_declare_lhs is called when exiting the function_declare_lhs production.
	ExitFunction_declare_lhs(c *Function_declare_lhsContext)

	// ExitFunction_declare is called when exiting the function_declare production.
	ExitFunction_declare(c *Function_declareContext)
}
