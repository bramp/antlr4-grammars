// Generated from lambda.g4 by ANTLR 4.7.

package lambda // lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lambdaListener is a complete listener for a parse tree produced by lambdaParser.
type lambdaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterApplication is called when entering the application production.
	EnterApplication(c *ApplicationContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitApplication is called when exiting the application production.
	ExitApplication(c *ApplicationContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)
}
