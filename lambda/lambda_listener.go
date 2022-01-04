// Code generated from lambda.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lambda // lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lambdaListener is a complete listener for a parse tree produced by lambdaParser.
type lambdaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterApplication is called when entering the application production.
	EnterApplication(c *ApplicationContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitApplication is called when exiting the application production.
	ExitApplication(c *ApplicationContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)
}
