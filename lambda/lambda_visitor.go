// Generated from lambda.g4 by ANTLR 4.7.

package lambda // lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by lambdaParser.
type lambdaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by lambdaParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by lambdaParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by lambdaParser#application.
	VisitApplication(ctx *ApplicationContext) interface{}

	// Visit a parse tree produced by lambdaParser#scope.
	VisitScope(ctx *ScopeContext) interface{}
}
