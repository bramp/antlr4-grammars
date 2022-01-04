// Code generated from lisa.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lisa // lisa
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselisaListener is a complete listener for a parse tree produced by lisaParser.
type BaselisaListener struct{}

var _ lisaListener = &BaselisaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselisaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselisaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselisaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselisaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselisaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaselisaListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration_block is called when production declaration_block is entered.
func (s *BaselisaListener) EnterDeclaration_block(ctx *Declaration_blockContext) {}

// ExitDeclaration_block is called when production declaration_block is exited.
func (s *BaselisaListener) ExitDeclaration_block(ctx *Declaration_blockContext) {}

// EnterDeclaration_statements is called when production declaration_statements is entered.
func (s *BaselisaListener) EnterDeclaration_statements(ctx *Declaration_statementsContext) {}

// ExitDeclaration_statements is called when production declaration_statements is exited.
func (s *BaselisaListener) ExitDeclaration_statements(ctx *Declaration_statementsContext) {}

// EnterDeclaration_statement is called when production declaration_statement is entered.
func (s *BaselisaListener) EnterDeclaration_statement(ctx *Declaration_statementContext) {}

// ExitDeclaration_statement is called when production declaration_statement is exited.
func (s *BaselisaListener) ExitDeclaration_statement(ctx *Declaration_statementContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaselisaListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaselisaListener) ExitType_(ctx *Type_Context) {}

// EnterProgram_block is called when production program_block is entered.
func (s *BaselisaListener) EnterProgram_block(ctx *Program_blockContext) {}

// ExitProgram_block is called when production program_block is exited.
func (s *BaselisaListener) ExitProgram_block(ctx *Program_blockContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaselisaListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaselisaListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaselisaListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaselisaListener) ExitStatement(ctx *StatementContext) {}

// EnterGenerating_statement is called when production generating_statement is entered.
func (s *BaselisaListener) EnterGenerating_statement(ctx *Generating_statementContext) {}

// ExitGenerating_statement is called when production generating_statement is exited.
func (s *BaselisaListener) ExitGenerating_statement(ctx *Generating_statementContext) {}

// EnterGenerator_type is called when production generator_type is entered.
func (s *BaselisaListener) EnterGenerator_type(ctx *Generator_typeContext) {}

// ExitGenerator_type is called when production generator_type is exited.
func (s *BaselisaListener) ExitGenerator_type(ctx *Generator_typeContext) {}

// EnterExpression_statement is called when production expression_statement is entered.
func (s *BaselisaListener) EnterExpression_statement(ctx *Expression_statementContext) {}

// ExitExpression_statement is called when production expression_statement is exited.
func (s *BaselisaListener) ExitExpression_statement(ctx *Expression_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaselisaListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaselisaListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaselisaListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaselisaListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaselisaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaselisaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExprop is called when production exprop is entered.
func (s *BaselisaListener) EnterExprop(ctx *ExpropContext) {}

// ExitExprop is called when production exprop is exited.
func (s *BaselisaListener) ExitExprop(ctx *ExpropContext) {}

// EnterSimple_expression is called when production simple_expression is entered.
func (s *BaselisaListener) EnterSimple_expression(ctx *Simple_expressionContext) {}

// ExitSimple_expression is called when production simple_expression is exited.
func (s *BaselisaListener) ExitSimple_expression(ctx *Simple_expressionContext) {}

// EnterOr_expression is called when production or_expression is entered.
func (s *BaselisaListener) EnterOr_expression(ctx *Or_expressionContext) {}

// ExitOr_expression is called when production or_expression is exited.
func (s *BaselisaListener) ExitOr_expression(ctx *Or_expressionContext) {}

// EnterUnary_relationexpression is called when production unary_relationexpression is entered.
func (s *BaselisaListener) EnterUnary_relationexpression(ctx *Unary_relationexpressionContext) {}

// ExitUnary_relationexpression is called when production unary_relationexpression is exited.
func (s *BaselisaListener) ExitUnary_relationexpression(ctx *Unary_relationexpressionContext) {}

// EnterRelation_expression is called when production relation_expression is entered.
func (s *BaselisaListener) EnterRelation_expression(ctx *Relation_expressionContext) {}

// ExitRelation_expression is called when production relation_expression is exited.
func (s *BaselisaListener) ExitRelation_expression(ctx *Relation_expressionContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaselisaListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaselisaListener) ExitRelop(ctx *RelopContext) {}

// EnterAdd_expression is called when production add_expression is entered.
func (s *BaselisaListener) EnterAdd_expression(ctx *Add_expressionContext) {}

// ExitAdd_expression is called when production add_expression is exited.
func (s *BaselisaListener) ExitAdd_expression(ctx *Add_expressionContext) {}

// EnterAddop is called when production addop is entered.
func (s *BaselisaListener) EnterAddop(ctx *AddopContext) {}

// ExitAddop is called when production addop is exited.
func (s *BaselisaListener) ExitAddop(ctx *AddopContext) {}

// EnterTerm is called when production term is entered.
func (s *BaselisaListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaselisaListener) ExitTerm(ctx *TermContext) {}

// EnterMultop is called when production multop is entered.
func (s *BaselisaListener) EnterMultop(ctx *MultopContext) {}

// ExitMultop is called when production multop is exited.
func (s *BaselisaListener) ExitMultop(ctx *MultopContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaselisaListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaselisaListener) ExitFactor(ctx *FactorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaselisaListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaselisaListener) ExitConstant(ctx *ConstantContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaselisaListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaselisaListener) ExitInteger(ctx *IntegerContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaselisaListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaselisaListener) ExitFunction_(ctx *Function_Context) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaselisaListener) EnterParameter_list(ctx *Parameter_listContext) {}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaselisaListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaselisaListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaselisaListener) ExitVariable(ctx *VariableContext) {}
