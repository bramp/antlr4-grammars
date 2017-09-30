// Generated from mumath.g4 by ANTLR 4.7.

package mumath // mumath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// mumathListener is a complete listener for a parse tree produced by mumathParser.
type mumathListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterActualParameter is called when entering the actualParameter production.
	EnterActualParameter(c *ActualParameterContext)

	// EnterStatments is called when entering the statments production.
	EnterStatments(c *StatmentsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterWhen is called when entering the when production.
	EnterWhen(c *WhenContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelationalOperator is called when entering the relationalOperator production.
	EnterRelationalOperator(c *RelationalOperatorContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterAddingOperator is called when entering the addingOperator production.
	EnterAddingOperator(c *AddingOperatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMultiplyingOperator is called when entering the multiplyingOperator production.
	EnterMultiplyingOperator(c *MultiplyingOperatorContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterFunctionDesignator is called when entering the functionDesignator production.
	EnterFunctionDesignator(c *FunctionDesignatorContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitActualParameter is called when exiting the actualParameter production.
	ExitActualParameter(c *ActualParameterContext)

	// ExitStatments is called when exiting the statments production.
	ExitStatments(c *StatmentsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitWhen is called when exiting the when production.
	ExitWhen(c *WhenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelationalOperator is called when exiting the relationalOperator production.
	ExitRelationalOperator(c *RelationalOperatorContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitAddingOperator is called when exiting the addingOperator production.
	ExitAddingOperator(c *AddingOperatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMultiplyingOperator is called when exiting the multiplyingOperator production.
	ExitMultiplyingOperator(c *MultiplyingOperatorContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitFunctionDesignator is called when exiting the functionDesignator production.
	ExitFunctionDesignator(c *FunctionDesignatorContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)
}
