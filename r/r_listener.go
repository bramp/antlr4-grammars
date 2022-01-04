// Code generated from R.g4 by ANTLR 4.9.3. DO NOT EDIT.

package r // R
import "github.com/antlr/antlr4/runtime/Go/antlr"

// RListener is a complete listener for a parse tree produced by RParser.
type RListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprlist is called when entering the exprlist production.
	EnterExprlist(c *ExprlistContext)

	// EnterFormlist is called when entering the formlist production.
	EnterFormlist(c *FormlistContext)

	// EnterForm is called when entering the form production.
	EnterForm(c *FormContext)

	// EnterSublist is called when entering the sublist production.
	EnterSublist(c *SublistContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprlist is called when exiting the exprlist production.
	ExitExprlist(c *ExprlistContext)

	// ExitFormlist is called when exiting the formlist production.
	ExitFormlist(c *FormlistContext)

	// ExitForm is called when exiting the form production.
	ExitForm(c *FormContext)

	// ExitSublist is called when exiting the sublist production.
	ExitSublist(c *SublistContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)
}
