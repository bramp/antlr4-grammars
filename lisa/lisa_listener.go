// Code generated from lisa.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lisa // lisa
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lisaListener is a complete listener for a parse tree produced by lisaParser.
type lisaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDeclaration_block is called when entering the declaration_block production.
	EnterDeclaration_block(c *Declaration_blockContext)

	// EnterDeclaration_statements is called when entering the declaration_statements production.
	EnterDeclaration_statements(c *Declaration_statementsContext)

	// EnterDeclaration_statement is called when entering the declaration_statement production.
	EnterDeclaration_statement(c *Declaration_statementContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterProgram_block is called when entering the program_block production.
	EnterProgram_block(c *Program_blockContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterGenerating_statement is called when entering the generating_statement production.
	EnterGenerating_statement(c *Generating_statementContext)

	// EnterGenerator_type is called when entering the generator_type production.
	EnterGenerator_type(c *Generator_typeContext)

	// EnterExpression_statement is called when entering the expression_statement production.
	EnterExpression_statement(c *Expression_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExprop is called when entering the exprop production.
	EnterExprop(c *ExpropContext)

	// EnterSimple_expression is called when entering the simple_expression production.
	EnterSimple_expression(c *Simple_expressionContext)

	// EnterOr_expression is called when entering the or_expression production.
	EnterOr_expression(c *Or_expressionContext)

	// EnterUnary_relationexpression is called when entering the unary_relationexpression production.
	EnterUnary_relationexpression(c *Unary_relationexpressionContext)

	// EnterRelation_expression is called when entering the relation_expression production.
	EnterRelation_expression(c *Relation_expressionContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// EnterAdd_expression is called when entering the add_expression production.
	EnterAdd_expression(c *Add_expressionContext)

	// EnterAddop is called when entering the addop production.
	EnterAddop(c *AddopContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMultop is called when entering the multop production.
	EnterMultop(c *MultopContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterParameter_list is called when entering the parameter_list production.
	EnterParameter_list(c *Parameter_listContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDeclaration_block is called when exiting the declaration_block production.
	ExitDeclaration_block(c *Declaration_blockContext)

	// ExitDeclaration_statements is called when exiting the declaration_statements production.
	ExitDeclaration_statements(c *Declaration_statementsContext)

	// ExitDeclaration_statement is called when exiting the declaration_statement production.
	ExitDeclaration_statement(c *Declaration_statementContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitProgram_block is called when exiting the program_block production.
	ExitProgram_block(c *Program_blockContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitGenerating_statement is called when exiting the generating_statement production.
	ExitGenerating_statement(c *Generating_statementContext)

	// ExitGenerator_type is called when exiting the generator_type production.
	ExitGenerator_type(c *Generator_typeContext)

	// ExitExpression_statement is called when exiting the expression_statement production.
	ExitExpression_statement(c *Expression_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExprop is called when exiting the exprop production.
	ExitExprop(c *ExpropContext)

	// ExitSimple_expression is called when exiting the simple_expression production.
	ExitSimple_expression(c *Simple_expressionContext)

	// ExitOr_expression is called when exiting the or_expression production.
	ExitOr_expression(c *Or_expressionContext)

	// ExitUnary_relationexpression is called when exiting the unary_relationexpression production.
	ExitUnary_relationexpression(c *Unary_relationexpressionContext)

	// ExitRelation_expression is called when exiting the relation_expression production.
	ExitRelation_expression(c *Relation_expressionContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)

	// ExitAdd_expression is called when exiting the add_expression production.
	ExitAdd_expression(c *Add_expressionContext)

	// ExitAddop is called when exiting the addop production.
	ExitAddop(c *AddopContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMultop is called when exiting the multop production.
	ExitMultop(c *MultopContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitParameter_list is called when exiting the parameter_list production.
	ExitParameter_list(c *Parameter_listContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)
}
