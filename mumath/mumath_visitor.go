// Generated from mumath.g4 by ANTLR 4.7.

package mumath // mumath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by mumathParser.
type mumathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mumathParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by mumathParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by mumathParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by mumathParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by mumathParser#actualParameter.
	VisitActualParameter(ctx *ActualParameterContext) interface{}

	// Visit a parse tree produced by mumathParser#statments.
	VisitStatments(ctx *StatmentsContext) interface{}

	// Visit a parse tree produced by mumathParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by mumathParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by mumathParser#when.
	VisitWhen(ctx *WhenContext) interface{}

	// Visit a parse tree produced by mumathParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by mumathParser#relationalOperator.
	VisitRelationalOperator(ctx *RelationalOperatorContext) interface{}

	// Visit a parse tree produced by mumathParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by mumathParser#addingOperator.
	VisitAddingOperator(ctx *AddingOperatorContext) interface{}

	// Visit a parse tree produced by mumathParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by mumathParser#multiplyingOperator.
	VisitMultiplyingOperator(ctx *MultiplyingOperatorContext) interface{}

	// Visit a parse tree produced by mumathParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by mumathParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by mumathParser#functionDesignator.
	VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{}

	// Visit a parse tree produced by mumathParser#equal.
	VisitEqual(ctx *EqualContext) interface{}
}
