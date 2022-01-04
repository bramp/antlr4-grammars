// Code generated from filter.g4 by ANTLR 4.9.3. DO NOT EDIT.

package filter // filter
import "github.com/antlr/antlr4/runtime/Go/antlr"

// filterListener is a complete listener for a parse tree produced by filterParser.
type filterListener interface {
	antlr.ParseTreeListener

	// EnterFilter_ is called when entering the filter_ production.
	EnterFilter_(c *Filter_Context)

	// EnterFiltercomp is called when entering the filtercomp production.
	EnterFiltercomp(c *FiltercompContext)

	// EnterAnd_ is called when entering the and_ production.
	EnterAnd_(c *And_Context)

	// EnterOr_ is called when entering the or_ production.
	EnterOr_(c *Or_Context)

	// EnterNot_ is called when entering the not_ production.
	EnterNot_(c *Not_Context)

	// EnterFilterlist is called when entering the filterlist production.
	EnterFilterlist(c *FilterlistContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterSimple is called when entering the simple production.
	EnterSimple(c *SimpleContext)

	// EnterFiltertype is called when entering the filtertype production.
	EnterFiltertype(c *FiltertypeContext)

	// EnterPresent is called when entering the present production.
	EnterPresent(c *PresentContext)

	// EnterSubstring is called when entering the substring production.
	EnterSubstring(c *SubstringContext)

	// EnterInitial is called when entering the initial production.
	EnterInitial(c *InitialContext)

	// EnterAny_ is called when entering the any_ production.
	EnterAny_(c *Any_Context)

	// EnterStarval is called when entering the starval production.
	EnterStarval(c *StarvalContext)

	// EnterFinal_ is called when entering the final_ production.
	EnterFinal_(c *Final_Context)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitFilter_ is called when exiting the filter_ production.
	ExitFilter_(c *Filter_Context)

	// ExitFiltercomp is called when exiting the filtercomp production.
	ExitFiltercomp(c *FiltercompContext)

	// ExitAnd_ is called when exiting the and_ production.
	ExitAnd_(c *And_Context)

	// ExitOr_ is called when exiting the or_ production.
	ExitOr_(c *Or_Context)

	// ExitNot_ is called when exiting the not_ production.
	ExitNot_(c *Not_Context)

	// ExitFilterlist is called when exiting the filterlist production.
	ExitFilterlist(c *FilterlistContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitSimple is called when exiting the simple production.
	ExitSimple(c *SimpleContext)

	// ExitFiltertype is called when exiting the filtertype production.
	ExitFiltertype(c *FiltertypeContext)

	// ExitPresent is called when exiting the present production.
	ExitPresent(c *PresentContext)

	// ExitSubstring is called when exiting the substring production.
	ExitSubstring(c *SubstringContext)

	// ExitInitial is called when exiting the initial production.
	ExitInitial(c *InitialContext)

	// ExitAny_ is called when exiting the any_ production.
	ExitAny_(c *Any_Context)

	// ExitStarval is called when exiting the starval production.
	ExitStarval(c *StarvalContext)

	// ExitFinal_ is called when exiting the final_ production.
	ExitFinal_(c *Final_Context)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
