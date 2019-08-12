// Code generated from propcalc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package propcalc // propcalc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// propcalcListener is a complete listener for a parse tree produced by propcalcParser.
type propcalcListener interface {
	antlr.ParseTreeListener

	// EnterProposition is called when entering the proposition production.
	EnterProposition(c *PropositionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelExpression is called when entering the relExpression production.
	EnterRelExpression(c *RelExpressionContext)

	// EnterAtoms is called when entering the atoms production.
	EnterAtoms(c *AtomsContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterEquiv is called when entering the equiv production.
	EnterEquiv(c *EquivContext)

	// EnterImplies is called when entering the implies production.
	EnterImplies(c *ImpliesContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// ExitProposition is called when exiting the proposition production.
	ExitProposition(c *PropositionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelExpression is called when exiting the relExpression production.
	ExitRelExpression(c *RelExpressionContext)

	// ExitAtoms is called when exiting the atoms production.
	ExitAtoms(c *AtomsContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitEquiv is called when exiting the equiv production.
	ExitEquiv(c *EquivContext)

	// ExitImplies is called when exiting the implies production.
	ExitImplies(c *ImpliesContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)
}
