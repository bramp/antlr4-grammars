// Code generated from SUOKIF.g4 by ANTLR 4.7.2. DO NOT EDIT.

package suokif // SUOKIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SUOKIFListener is a complete listener for a parse tree produced by SUOKIFParser.
type SUOKIFListener interface {
	antlr.ParseTreeListener

	// EnterTop_level is called when entering the top_level production.
	EnterTop_level(c *Top_levelContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterFunterm is called when entering the funterm production.
	EnterFunterm(c *FuntermContext)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterRelsent is called when entering the relsent production.
	EnterRelsent(c *RelsentContext)

	// EnterLogsent is called when entering the logsent production.
	EnterLogsent(c *LogsentContext)

	// EnterQuantsent is called when entering the quantsent production.
	EnterQuantsent(c *QuantsentContext)

	// ExitTop_level is called when exiting the top_level production.
	ExitTop_level(c *Top_levelContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitFunterm is called when exiting the funterm production.
	ExitFunterm(c *FuntermContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitRelsent is called when exiting the relsent production.
	ExitRelsent(c *RelsentContext)

	// ExitLogsent is called when exiting the logsent production.
	ExitLogsent(c *LogsentContext)

	// ExitQuantsent is called when exiting the quantsent production.
	ExitQuantsent(c *QuantsentContext)
}
