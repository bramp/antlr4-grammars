// Generated from Abnf.g4 by ANTLR 4.7.

package abnf // Abnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbnfListener is a complete listener for a parse tree produced by AbnfParser.
type AbnfListener interface {
	antlr.ParseTreeListener

	// EnterRulelist is called when entering the rulelist production.
	EnterRulelist(c *RulelistContext)

	// EnterRule_ is called when entering the rule_ production.
	EnterRule_(c *Rule_Context)

	// EnterElements is called when entering the elements production.
	EnterElements(c *ElementsContext)

	// EnterAlternation is called when entering the alternation production.
	EnterAlternation(c *AlternationContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterRepetition is called when entering the repetition production.
	EnterRepetition(c *RepetitionContext)

	// EnterRepeat is called when entering the repeat production.
	EnterRepeat(c *RepeatContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// ExitRulelist is called when exiting the rulelist production.
	ExitRulelist(c *RulelistContext)

	// ExitRule_ is called when exiting the rule_ production.
	ExitRule_(c *Rule_Context)

	// ExitElements is called when exiting the elements production.
	ExitElements(c *ElementsContext)

	// ExitAlternation is called when exiting the alternation production.
	ExitAlternation(c *AlternationContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitRepetition is called when exiting the repetition production.
	ExitRepetition(c *RepetitionContext)

	// ExitRepeat is called when exiting the repeat production.
	ExitRepeat(c *RepeatContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)
}
