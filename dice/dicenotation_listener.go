// Generated from DiceNotation.g4 by ANTLR 4.7.

package dice // DiceNotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiceNotationListener is a complete listener for a parse tree produced by DiceNotationParser.
type DiceNotationListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterDice is called when entering the dice production.
	EnterDice(c *DiceContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitDice is called when exiting the dice production.
	ExitDice(c *DiceContext)
}
