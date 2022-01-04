// Code generated from DiceNotationParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dicenotation // DiceNotationParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiceNotationParserListener is a complete listener for a parse tree produced by DiceNotationParser.
type DiceNotationParserListener interface {
	antlr.ParseTreeListener

	// EnterNotation is called when entering the notation production.
	EnterNotation(c *NotationContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterMultOp is called when entering the multOp production.
	EnterMultOp(c *MultOpContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterDice is called when entering the dice production.
	EnterDice(c *DiceContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitNotation is called when exiting the notation production.
	ExitNotation(c *NotationContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitMultOp is called when exiting the multOp production.
	ExitMultOp(c *MultOpContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitDice is called when exiting the dice production.
	ExitDice(c *DiceContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
