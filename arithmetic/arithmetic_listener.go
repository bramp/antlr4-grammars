// Code generated from arithmetic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package arithmetic // arithmetic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// arithmeticListener is a complete listener for a parse tree produced by arithmeticParser.
type arithmeticListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
