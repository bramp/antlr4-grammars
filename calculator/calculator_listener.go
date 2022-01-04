// Code generated from calculator.g4 by ANTLR 4.9.3. DO NOT EDIT.

package calculator // calculator
import "github.com/antlr/antlr4/runtime/Go/antlr"

// calculatorListener is a complete listener for a parse tree produced by calculatorParser.
type calculatorListener interface {
	antlr.ParseTreeListener

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterPowExpression is called when entering the powExpression production.
	EnterPowExpression(c *PowExpressionContext)

	// EnterSignedAtom is called when entering the signedAtom production.
	EnterSignedAtom(c *SignedAtomContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterScientific is called when entering the scientific production.
	EnterScientific(c *ScientificContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunc_ is called when entering the func_ production.
	EnterFunc_(c *Func_Context)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitPowExpression is called when exiting the powExpression production.
	ExitPowExpression(c *PowExpressionContext)

	// ExitSignedAtom is called when exiting the signedAtom production.
	ExitSignedAtom(c *SignedAtomContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitScientific is called when exiting the scientific production.
	ExitScientific(c *ScientificContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunc_ is called when exiting the func_ production.
	ExitFunc_(c *Func_Context)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
