// Code generated from bnf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bnf // bnf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bnfListener is a complete listener for a parse tree produced by bnfParser.
type bnfListener interface {
	antlr.ParseTreeListener

	// EnterRulelist is called when entering the rulelist production.
	EnterRulelist(c *RulelistContext)

	// EnterRule_ is called when entering the rule_ production.
	EnterRule_(c *Rule_Context)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterRhs is called when entering the rhs production.
	EnterRhs(c *RhsContext)

	// EnterAlternatives is called when entering the alternatives production.
	EnterAlternatives(c *AlternativesContext)

	// EnterAlternative is called when entering the alternative production.
	EnterAlternative(c *AlternativeContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterOptional_ is called when entering the optional_ production.
	EnterOptional_(c *Optional_Context)

	// EnterZeroormore is called when entering the zeroormore production.
	EnterZeroormore(c *ZeroormoreContext)

	// EnterOneormore is called when entering the oneormore production.
	EnterOneormore(c *OneormoreContext)

	// EnterText_ is called when entering the text_ production.
	EnterText_(c *Text_Context)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// EnterRuleid is called when entering the ruleid production.
	EnterRuleid(c *RuleidContext)

	// ExitRulelist is called when exiting the rulelist production.
	ExitRulelist(c *RulelistContext)

	// ExitRule_ is called when exiting the rule_ production.
	ExitRule_(c *Rule_Context)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitRhs is called when exiting the rhs production.
	ExitRhs(c *RhsContext)

	// ExitAlternatives is called when exiting the alternatives production.
	ExitAlternatives(c *AlternativesContext)

	// ExitAlternative is called when exiting the alternative production.
	ExitAlternative(c *AlternativeContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitOptional_ is called when exiting the optional_ production.
	ExitOptional_(c *Optional_Context)

	// ExitZeroormore is called when exiting the zeroormore production.
	ExitZeroormore(c *ZeroormoreContext)

	// ExitOneormore is called when exiting the oneormore production.
	ExitOneormore(c *OneormoreContext)

	// ExitText_ is called when exiting the text_ production.
	ExitText_(c *Text_Context)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)

	// ExitRuleid is called when exiting the ruleid production.
	ExitRuleid(c *RuleidContext)
}
