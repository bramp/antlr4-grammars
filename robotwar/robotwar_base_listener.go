// Code generated from robotwar.g4 by ANTLR 4.9.3. DO NOT EDIT.

package robotwar // robotwar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaserobotwarListener is a complete listener for a parse tree produced by robotwarParser.
type BaserobotwarListener struct{}

var _ robotwarListener = &BaserobotwarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaserobotwarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaserobotwarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaserobotwarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaserobotwarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaserobotwarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaserobotwarListener) ExitProgram(ctx *ProgramContext) {}

// EnterLine is called when production line is entered.
func (s *BaserobotwarListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaserobotwarListener) ExitLine(ctx *LineContext) {}

// EnterLabel is called when production label is entered.
func (s *BaserobotwarListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaserobotwarListener) ExitLabel(ctx *LabelContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaserobotwarListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaserobotwarListener) ExitStatement(ctx *StatementContext) {}

// EnterAccumstatement is called when production accumstatement is entered.
func (s *BaserobotwarListener) EnterAccumstatement(ctx *AccumstatementContext) {}

// ExitAccumstatement is called when production accumstatement is exited.
func (s *BaserobotwarListener) ExitAccumstatement(ctx *AccumstatementContext) {}

// EnterAccumexpression is called when production accumexpression is entered.
func (s *BaserobotwarListener) EnterAccumexpression(ctx *AccumexpressionContext) {}

// ExitAccumexpression is called when production accumexpression is exited.
func (s *BaserobotwarListener) ExitAccumexpression(ctx *AccumexpressionContext) {}

// EnterGosubstatement is called when production gosubstatement is entered.
func (s *BaserobotwarListener) EnterGosubstatement(ctx *GosubstatementContext) {}

// ExitGosubstatement is called when production gosubstatement is exited.
func (s *BaserobotwarListener) ExitGosubstatement(ctx *GosubstatementContext) {}

// EnterGotostatement is called when production gotostatement is entered.
func (s *BaserobotwarListener) EnterGotostatement(ctx *GotostatementContext) {}

// ExitGotostatement is called when production gotostatement is exited.
func (s *BaserobotwarListener) ExitGotostatement(ctx *GotostatementContext) {}

// EnterTostatement is called when production tostatement is entered.
func (s *BaserobotwarListener) EnterTostatement(ctx *TostatementContext) {}

// ExitTostatement is called when production tostatement is exited.
func (s *BaserobotwarListener) ExitTostatement(ctx *TostatementContext) {}

// EnterEndsubstatement is called when production endsubstatement is entered.
func (s *BaserobotwarListener) EnterEndsubstatement(ctx *EndsubstatementContext) {}

// ExitEndsubstatement is called when production endsubstatement is exited.
func (s *BaserobotwarListener) ExitEndsubstatement(ctx *EndsubstatementContext) {}

// EnterIfstatement is called when production ifstatement is entered.
func (s *BaserobotwarListener) EnterIfstatement(ctx *IfstatementContext) {}

// ExitIfstatement is called when production ifstatement is exited.
func (s *BaserobotwarListener) ExitIfstatement(ctx *IfstatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaserobotwarListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaserobotwarListener) ExitCondition(ctx *ConditionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaserobotwarListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaserobotwarListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaserobotwarListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaserobotwarListener) ExitOperation(ctx *OperationContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaserobotwarListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaserobotwarListener) ExitComparison(ctx *ComparisonContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaserobotwarListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaserobotwarListener) ExitArgument(ctx *ArgumentContext) {}

// EnterRegister_ is called when production register_ is entered.
func (s *BaserobotwarListener) EnterRegister_(ctx *Register_Context) {}

// ExitRegister_ is called when production register_ is exited.
func (s *BaserobotwarListener) ExitRegister_(ctx *Register_Context) {}

// EnterNumber is called when production number is entered.
func (s *BaserobotwarListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaserobotwarListener) ExitNumber(ctx *NumberContext) {}

// EnterComment is called when production comment is entered.
func (s *BaserobotwarListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaserobotwarListener) ExitComment(ctx *CommentContext) {}
