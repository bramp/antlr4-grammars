// Code generated from matlab.g4 by ANTLR 4.7.2. DO NOT EDIT.

package matlab // matlab
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasematlabListener is a complete listener for a parse tree produced by matlabParser.
type BasematlabListener struct{}

var _ matlabListener = &BasematlabListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasematlabListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasematlabListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasematlabListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasematlabListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BasematlabListener) EnterPrimary_expression(ctx *Primary_expressionContext) {}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BasematlabListener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterPostfix_expression is called when production postfix_expression is entered.
func (s *BasematlabListener) EnterPostfix_expression(ctx *Postfix_expressionContext) {}

// ExitPostfix_expression is called when production postfix_expression is exited.
func (s *BasematlabListener) ExitPostfix_expression(ctx *Postfix_expressionContext) {}

// EnterIndex_expression is called when production index_expression is entered.
func (s *BasematlabListener) EnterIndex_expression(ctx *Index_expressionContext) {}

// ExitIndex_expression is called when production index_expression is exited.
func (s *BasematlabListener) ExitIndex_expression(ctx *Index_expressionContext) {}

// EnterIndex_expression_list is called when production index_expression_list is entered.
func (s *BasematlabListener) EnterIndex_expression_list(ctx *Index_expression_listContext) {}

// ExitIndex_expression_list is called when production index_expression_list is exited.
func (s *BasematlabListener) ExitIndex_expression_list(ctx *Index_expression_listContext) {}

// EnterArray_expression is called when production array_expression is entered.
func (s *BasematlabListener) EnterArray_expression(ctx *Array_expressionContext) {}

// ExitArray_expression is called when production array_expression is exited.
func (s *BasematlabListener) ExitArray_expression(ctx *Array_expressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BasematlabListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BasematlabListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BasematlabListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BasematlabListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BasematlabListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BasematlabListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BasematlabListener) EnterAdditive_expression(ctx *Additive_expressionContext) {}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BasematlabListener) ExitAdditive_expression(ctx *Additive_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BasematlabListener) EnterRelational_expression(ctx *Relational_expressionContext) {}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BasematlabListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BasematlabListener) EnterEquality_expression(ctx *Equality_expressionContext) {}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BasematlabListener) ExitEquality_expression(ctx *Equality_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BasematlabListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BasematlabListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterOr_expression is called when production or_expression is entered.
func (s *BasematlabListener) EnterOr_expression(ctx *Or_expressionContext) {}

// ExitOr_expression is called when production or_expression is exited.
func (s *BasematlabListener) ExitOr_expression(ctx *Or_expressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasematlabListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasematlabListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignment_expression is called when production assignment_expression is entered.
func (s *BasematlabListener) EnterAssignment_expression(ctx *Assignment_expressionContext) {}

// ExitAssignment_expression is called when production assignment_expression is exited.
func (s *BasematlabListener) ExitAssignment_expression(ctx *Assignment_expressionContext) {}

// EnterEostmt is called when production eostmt is entered.
func (s *BasematlabListener) EnterEostmt(ctx *EostmtContext) {}

// ExitEostmt is called when production eostmt is exited.
func (s *BasematlabListener) ExitEostmt(ctx *EostmtContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasematlabListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasematlabListener) ExitStatement(ctx *StatementContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BasematlabListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BasematlabListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *BasematlabListener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *BasematlabListener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterGlobal_statement is called when production global_statement is entered.
func (s *BasematlabListener) EnterGlobal_statement(ctx *Global_statementContext) {}

// ExitGlobal_statement is called when production global_statement is exited.
func (s *BasematlabListener) ExitGlobal_statement(ctx *Global_statementContext) {}

// EnterClear_statement is called when production clear_statement is entered.
func (s *BasematlabListener) EnterClear_statement(ctx *Clear_statementContext) {}

// ExitClear_statement is called when production clear_statement is exited.
func (s *BasematlabListener) ExitClear_statement(ctx *Clear_statementContext) {}

// EnterExpression_statement is called when production expression_statement is entered.
func (s *BasematlabListener) EnterExpression_statement(ctx *Expression_statementContext) {}

// ExitExpression_statement is called when production expression_statement is exited.
func (s *BasematlabListener) ExitExpression_statement(ctx *Expression_statementContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BasematlabListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BasematlabListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterArray_element is called when production array_element is entered.
func (s *BasematlabListener) EnterArray_element(ctx *Array_elementContext) {}

// ExitArray_element is called when production array_element is exited.
func (s *BasematlabListener) ExitArray_element(ctx *Array_elementContext) {}

// EnterArray_list is called when production array_list is entered.
func (s *BasematlabListener) EnterArray_list(ctx *Array_listContext) {}

// ExitArray_list is called when production array_list is exited.
func (s *BasematlabListener) ExitArray_list(ctx *Array_listContext) {}

// EnterSelection_statement is called when production selection_statement is entered.
func (s *BasematlabListener) EnterSelection_statement(ctx *Selection_statementContext) {}

// ExitSelection_statement is called when production selection_statement is exited.
func (s *BasematlabListener) ExitSelection_statement(ctx *Selection_statementContext) {}

// EnterElseif_clause is called when production elseif_clause is entered.
func (s *BasematlabListener) EnterElseif_clause(ctx *Elseif_clauseContext) {}

// ExitElseif_clause is called when production elseif_clause is exited.
func (s *BasematlabListener) ExitElseif_clause(ctx *Elseif_clauseContext) {}

// EnterIteration_statement is called when production iteration_statement is entered.
func (s *BasematlabListener) EnterIteration_statement(ctx *Iteration_statementContext) {}

// ExitIteration_statement is called when production iteration_statement is exited.
func (s *BasematlabListener) ExitIteration_statement(ctx *Iteration_statementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BasematlabListener) EnterJump_statement(ctx *Jump_statementContext) {}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BasematlabListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterTranslation_unit is called when production translation_unit is entered.
func (s *BasematlabListener) EnterTranslation_unit(ctx *Translation_unitContext) {}

// ExitTranslation_unit is called when production translation_unit is exited.
func (s *BasematlabListener) ExitTranslation_unit(ctx *Translation_unitContext) {}

// EnterFunc_ident_list is called when production func_ident_list is entered.
func (s *BasematlabListener) EnterFunc_ident_list(ctx *Func_ident_listContext) {}

// ExitFunc_ident_list is called when production func_ident_list is exited.
func (s *BasematlabListener) ExitFunc_ident_list(ctx *Func_ident_listContext) {}

// EnterFunc_return_list is called when production func_return_list is entered.
func (s *BasematlabListener) EnterFunc_return_list(ctx *Func_return_listContext) {}

// ExitFunc_return_list is called when production func_return_list is exited.
func (s *BasematlabListener) ExitFunc_return_list(ctx *Func_return_listContext) {}

// EnterFunction_declare_lhs is called when production function_declare_lhs is entered.
func (s *BasematlabListener) EnterFunction_declare_lhs(ctx *Function_declare_lhsContext) {}

// ExitFunction_declare_lhs is called when production function_declare_lhs is exited.
func (s *BasematlabListener) ExitFunction_declare_lhs(ctx *Function_declare_lhsContext) {}

// EnterFunction_declare is called when production function_declare is entered.
func (s *BasematlabListener) EnterFunction_declare(ctx *Function_declareContext) {}

// ExitFunction_declare is called when production function_declare is exited.
func (s *BasematlabListener) ExitFunction_declare(ctx *Function_declareContext) {}
