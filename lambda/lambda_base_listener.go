// Code generated from lambda.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lambda // lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselambdaListener is a complete listener for a parse tree produced by lambdaParser.
type BaselambdaListener struct{}

var _ lambdaListener = &BaselambdaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselambdaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselambdaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselambdaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselambdaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaselambdaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaselambdaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaselambdaListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaselambdaListener) ExitFunction_(ctx *Function_Context) {}

// EnterApplication is called when production application is entered.
func (s *BaselambdaListener) EnterApplication(ctx *ApplicationContext) {}

// ExitApplication is called when production application is exited.
func (s *BaselambdaListener) ExitApplication(ctx *ApplicationContext) {}

// EnterScope is called when production scope is entered.
func (s *BaselambdaListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaselambdaListener) ExitScope(ctx *ScopeContext) {}
