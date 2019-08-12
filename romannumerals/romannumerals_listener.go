// Code generated from romannumerals.g4 by ANTLR 4.7.2. DO NOT EDIT.

package romannumerals // romannumerals
import "github.com/antlr/antlr4/runtime/Go/antlr"

// romannumeralsListener is a complete listener for a parse tree produced by romannumeralsParser.
type romannumeralsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterThousands is called when entering the thousands production.
	EnterThousands(c *ThousandsContext)

	// EnterThous_part is called when entering the thous_part production.
	EnterThous_part(c *Thous_partContext)

	// EnterHundreds is called when entering the hundreds production.
	EnterHundreds(c *HundredsContext)

	// EnterHun_part is called when entering the hun_part production.
	EnterHun_part(c *Hun_partContext)

	// EnterHun_rep is called when entering the hun_rep production.
	EnterHun_rep(c *Hun_repContext)

	// EnterTens is called when entering the tens production.
	EnterTens(c *TensContext)

	// EnterTens_part is called when entering the tens_part production.
	EnterTens_part(c *Tens_partContext)

	// EnterTens_rep is called when entering the tens_rep production.
	EnterTens_rep(c *Tens_repContext)

	// EnterOnes is called when entering the ones production.
	EnterOnes(c *OnesContext)

	// EnterOnes_rep is called when entering the ones_rep production.
	EnterOnes_rep(c *Ones_repContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitThousands is called when exiting the thousands production.
	ExitThousands(c *ThousandsContext)

	// ExitThous_part is called when exiting the thous_part production.
	ExitThous_part(c *Thous_partContext)

	// ExitHundreds is called when exiting the hundreds production.
	ExitHundreds(c *HundredsContext)

	// ExitHun_part is called when exiting the hun_part production.
	ExitHun_part(c *Hun_partContext)

	// ExitHun_rep is called when exiting the hun_rep production.
	ExitHun_rep(c *Hun_repContext)

	// ExitTens is called when exiting the tens production.
	ExitTens(c *TensContext)

	// ExitTens_part is called when exiting the tens_part production.
	ExitTens_part(c *Tens_partContext)

	// ExitTens_rep is called when exiting the tens_rep production.
	ExitTens_rep(c *Tens_repContext)

	// ExitOnes is called when exiting the ones production.
	ExitOnes(c *OnesContext)

	// ExitOnes_rep is called when exiting the ones_rep production.
	ExitOnes_rep(c *Ones_repContext)
}
