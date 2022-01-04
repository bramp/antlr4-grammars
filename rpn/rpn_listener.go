// Code generated from rpn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package rpn // rpn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// rpnListener is a complete listener for a parse tree produced by rpnParser.
type rpnListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterOper is called when entering the oper production.
	EnterOper(c *OperContext)

	// EnterSignedAtom is called when entering the signedAtom production.
	EnterSignedAtom(c *SignedAtomContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitOper is called when exiting the oper production.
	ExitOper(c *OperContext)

	// ExitSignedAtom is called when exiting the signedAtom production.
	ExitSignedAtom(c *SignedAtomContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)
}
