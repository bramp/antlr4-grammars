// Generated from mumath.g4 by ANTLR 4.7.

package mumath // mumath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemumathListener is a complete listener for a parse tree produced by mumathParser.
type BasemumathListener struct{}

var _ mumathListener = &BasemumathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemumathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemumathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemumathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemumathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasemumathListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasemumathListener) ExitProgram(ctx *ProgramContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasemumathListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasemumathListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterList is called when production list is entered.
func (s *BasemumathListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasemumathListener) ExitList(ctx *ListContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BasemumathListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BasemumathListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterActualParameter is called when production actualParameter is entered.
func (s *BasemumathListener) EnterActualParameter(ctx *ActualParameterContext) {}

// ExitActualParameter is called when production actualParameter is exited.
func (s *BasemumathListener) ExitActualParameter(ctx *ActualParameterContext) {}

// EnterStatments is called when production statments is entered.
func (s *BasemumathListener) EnterStatments(ctx *StatmentsContext) {}

// ExitStatments is called when production statments is exited.
func (s *BasemumathListener) ExitStatments(ctx *StatmentsContext) {}

// EnterBlock is called when production block is entered.
func (s *BasemumathListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasemumathListener) ExitBlock(ctx *BlockContext) {}

// EnterLoop is called when production loop is entered.
func (s *BasemumathListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BasemumathListener) ExitLoop(ctx *LoopContext) {}

// EnterWhen is called when production when is entered.
func (s *BasemumathListener) EnterWhen(ctx *WhenContext) {}

// ExitWhen is called when production when is exited.
func (s *BasemumathListener) ExitWhen(ctx *WhenContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasemumathListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasemumathListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelationalOperator is called when production relationalOperator is entered.
func (s *BasemumathListener) EnterRelationalOperator(ctx *RelationalOperatorContext) {}

// ExitRelationalOperator is called when production relationalOperator is exited.
func (s *BasemumathListener) ExitRelationalOperator(ctx *RelationalOperatorContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BasemumathListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BasemumathListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterAddingOperator is called when production addingOperator is entered.
func (s *BasemumathListener) EnterAddingOperator(ctx *AddingOperatorContext) {}

// ExitAddingOperator is called when production addingOperator is exited.
func (s *BasemumathListener) ExitAddingOperator(ctx *AddingOperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BasemumathListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasemumathListener) ExitTerm(ctx *TermContext) {}

// EnterMultiplyingOperator is called when production multiplyingOperator is entered.
func (s *BasemumathListener) EnterMultiplyingOperator(ctx *MultiplyingOperatorContext) {}

// ExitMultiplyingOperator is called when production multiplyingOperator is exited.
func (s *BasemumathListener) ExitMultiplyingOperator(ctx *MultiplyingOperatorContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasemumathListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasemumathListener) ExitFactor(ctx *FactorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasemumathListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasemumathListener) ExitConstant(ctx *ConstantContext) {}

// EnterFunctionDesignator is called when production functionDesignator is entered.
func (s *BasemumathListener) EnterFunctionDesignator(ctx *FunctionDesignatorContext) {}

// ExitFunctionDesignator is called when production functionDesignator is exited.
func (s *BasemumathListener) ExitFunctionDesignator(ctx *FunctionDesignatorContext) {}

// EnterEqual is called when production equal is entered.
func (s *BasemumathListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BasemumathListener) ExitEqual(ctx *EqualContext) {}
