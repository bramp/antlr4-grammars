// Generated from arithmetic.g4 by ANTLR 4.7.

package arithmetic // arithmetic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// arithmeticListener is a complete listener for a parse tree produced by arithmeticParser.
type arithmeticListener interface {
	antlr.ParseTreeListener

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterSignedAtom is called when entering the signedAtom production.
	EnterSignedAtom(c *SignedAtomContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitSignedAtom is called when exiting the signedAtom production.
	ExitSignedAtom(c *SignedAtomContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
