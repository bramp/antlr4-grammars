// Code generated from tnt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tnt // tnt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tntListener is a complete listener for a parse tree produced by tntParser.
type tntListener interface {
	antlr.ParseTreeListener

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterForevery is called when entering the forevery production.
	EnterForevery(c *ForeveryContext)

	// EnterExists is called when entering the exists production.
	EnterExists(c *ExistsContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitForevery is called when exiting the forevery production.
	ExitForevery(c *ForeveryContext)

	// ExitExists is called when exiting the exists production.
	ExitExists(c *ExistsContext)
}
