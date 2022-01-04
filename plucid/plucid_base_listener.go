// Code generated from plucid.g4 by ANTLR 4.9.3. DO NOT EDIT.

package plucid // plucid
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseplucidListener is a complete listener for a parse tree produced by plucidParser.
type BaseplucidListener struct{}

var _ plucidListener = &BaseplucidListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseplucidListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseplucidListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseplucidListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseplucidListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseplucidListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseplucidListener) ExitProgram(ctx *ProgramContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseplucidListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseplucidListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseplucidListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseplucidListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumeric_constant is called when production numeric_constant is entered.
func (s *BaseplucidListener) EnterNumeric_constant(ctx *Numeric_constantContext) {}

// ExitNumeric_constant is called when production numeric_constant is exited.
func (s *BaseplucidListener) ExitNumeric_constant(ctx *Numeric_constantContext) {}

// EnterInteger_constant is called when production integer_constant is entered.
func (s *BaseplucidListener) EnterInteger_constant(ctx *Integer_constantContext) {}

// ExitInteger_constant is called when production integer_constant is exited.
func (s *BaseplucidListener) ExitInteger_constant(ctx *Integer_constantContext) {}

// EnterReal_constant is called when production real_constant is entered.
func (s *BaseplucidListener) EnterReal_constant(ctx *Real_constantContext) {}

// ExitReal_constant is called when production real_constant is exited.
func (s *BaseplucidListener) ExitReal_constant(ctx *Real_constantContext) {}

// EnterWord_constant is called when production word_constant is entered.
func (s *BaseplucidListener) EnterWord_constant(ctx *Word_constantContext) {}

// ExitWord_constant is called when production word_constant is exited.
func (s *BaseplucidListener) ExitWord_constant(ctx *Word_constantContext) {}

// EnterWord_constant_less_the_quotes is called when production word_constant_less_the_quotes is entered.
func (s *BaseplucidListener) EnterWord_constant_less_the_quotes(ctx *Word_constant_less_the_quotesContext) {
}

// ExitWord_constant_less_the_quotes is called when production word_constant_less_the_quotes is exited.
func (s *BaseplucidListener) ExitWord_constant_less_the_quotes(ctx *Word_constant_less_the_quotesContext) {
}

// EnterList_constant is called when production list_constant is entered.
func (s *BaseplucidListener) EnterList_constant(ctx *List_constantContext) {}

// ExitList_constant is called when production list_constant is exited.
func (s *BaseplucidListener) ExitList_constant(ctx *List_constantContext) {}

// EnterList_constant_element is called when production list_constant_element is entered.
func (s *BaseplucidListener) EnterList_constant_element(ctx *List_constant_elementContext) {}

// ExitList_constant_element is called when production list_constant_element is exited.
func (s *BaseplucidListener) ExitList_constant_element(ctx *List_constant_elementContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseplucidListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseplucidListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterPrefix_operator is called when production prefix_operator is entered.
func (s *BaseplucidListener) EnterPrefix_operator(ctx *Prefix_operatorContext) {}

// ExitPrefix_operator is called when production prefix_operator is exited.
func (s *BaseplucidListener) ExitPrefix_operator(ctx *Prefix_operatorContext) {}

// EnterInfix_operator is called when production infix_operator is entered.
func (s *BaseplucidListener) EnterInfix_operator(ctx *Infix_operatorContext) {}

// ExitInfix_operator is called when production infix_operator is exited.
func (s *BaseplucidListener) ExitInfix_operator(ctx *Infix_operatorContext) {}

// EnterList_expression is called when production list_expression is entered.
func (s *BaseplucidListener) EnterList_expression(ctx *List_expressionContext) {}

// ExitList_expression is called when production list_expression is exited.
func (s *BaseplucidListener) ExitList_expression(ctx *List_expressionContext) {}

// EnterExpressions_list is called when production expressions_list is entered.
func (s *BaseplucidListener) EnterExpressions_list(ctx *Expressions_listContext) {}

// ExitExpressions_list is called when production expressions_list is exited.
func (s *BaseplucidListener) ExitExpressions_list(ctx *Expressions_listContext) {}

// EnterExpression_item is called when production expression_item is entered.
func (s *BaseplucidListener) EnterExpression_item(ctx *Expression_itemContext) {}

// ExitExpression_item is called when production expression_item is exited.
func (s *BaseplucidListener) ExitExpression_item(ctx *Expression_itemContext) {}

// EnterIf_expression is called when production if_expression is entered.
func (s *BaseplucidListener) EnterIf_expression(ctx *If_expressionContext) {}

// ExitIf_expression is called when production if_expression is exited.
func (s *BaseplucidListener) ExitIf_expression(ctx *If_expressionContext) {}

// EnterEndif is called when production endif is entered.
func (s *BaseplucidListener) EnterEndif(ctx *EndifContext) {}

// ExitEndif is called when production endif is exited.
func (s *BaseplucidListener) ExitEndif(ctx *EndifContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BaseplucidListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BaseplucidListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterCond_expression is called when production cond_expression is entered.
func (s *BaseplucidListener) EnterCond_expression(ctx *Cond_expressionContext) {}

// ExitCond_expression is called when production cond_expression is exited.
func (s *BaseplucidListener) ExitCond_expression(ctx *Cond_expressionContext) {}

// EnterCbody is called when production cbody is entered.
func (s *BaseplucidListener) EnterCbody(ctx *CbodyContext) {}

// ExitCbody is called when production cbody is exited.
func (s *BaseplucidListener) ExitCbody(ctx *CbodyContext) {}

// EnterDefaultcase is called when production defaultcase is entered.
func (s *BaseplucidListener) EnterDefaultcase(ctx *DefaultcaseContext) {}

// ExitDefaultcase is called when production defaultcase is exited.
func (s *BaseplucidListener) ExitDefaultcase(ctx *DefaultcaseContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseplucidListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseplucidListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterActuals_list is called when production actuals_list is entered.
func (s *BaseplucidListener) EnterActuals_list(ctx *Actuals_listContext) {}

// ExitActuals_list is called when production actuals_list is exited.
func (s *BaseplucidListener) ExitActuals_list(ctx *Actuals_listContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseplucidListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseplucidListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterBody is called when production body is entered.
func (s *BaseplucidListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseplucidListener) ExitBody(ctx *BodyContext) {}

// EnterDeclarations_list is called when production declarations_list is entered.
func (s *BaseplucidListener) EnterDeclarations_list(ctx *Declarations_listContext) {}

// ExitDeclarations_list is called when production declarations_list is exited.
func (s *BaseplucidListener) ExitDeclarations_list(ctx *Declarations_listContext) {}

// EnterCurrent_declaration is called when production current_declaration is entered.
func (s *BaseplucidListener) EnterCurrent_declaration(ctx *Current_declarationContext) {}

// ExitCurrent_declaration is called when production current_declaration is exited.
func (s *BaseplucidListener) ExitCurrent_declaration(ctx *Current_declarationContext) {}

// EnterDefinitions_list is called when production definitions_list is entered.
func (s *BaseplucidListener) EnterDefinitions_list(ctx *Definitions_listContext) {}

// ExitDefinitions_list is called when production definitions_list is exited.
func (s *BaseplucidListener) ExitDefinitions_list(ctx *Definitions_listContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseplucidListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseplucidListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterSimple_definition is called when production simple_definition is entered.
func (s *BaseplucidListener) EnterSimple_definition(ctx *Simple_definitionContext) {}

// ExitSimple_definition is called when production simple_definition is exited.
func (s *BaseplucidListener) ExitSimple_definition(ctx *Simple_definitionContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseplucidListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseplucidListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterFormals_list is called when production formals_list is entered.
func (s *BaseplucidListener) EnterFormals_list(ctx *Formals_listContext) {}

// ExitFormals_list is called when production formals_list is exited.
func (s *BaseplucidListener) ExitFormals_list(ctx *Formals_listContext) {}
