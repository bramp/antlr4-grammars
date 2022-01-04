// Code generated from RCS.g4 by ANTLR 4.9.3. DO NOT EDIT.

package rcs // RCS
import "github.com/antlr/antlr4/runtime/Go/antlr"

// RCSListener is a complete listener for a parse tree produced by RCSParser.
type RCSListener interface {
	antlr.ParseTreeListener

	// EnterRcstext is called when entering the rcstext production.
	EnterRcstext(c *RcstextContext)

	// EnterRcsheader is called when entering the rcsheader production.
	EnterRcsheader(c *RcsheaderContext)

	// EnterRcsrevisions is called when entering the rcsrevisions production.
	EnterRcsrevisions(c *RcsrevisionsContext)

	// EnterAdmin is called when entering the admin production.
	EnterAdmin(c *AdminContext)

	// EnterHead is called when entering the head production.
	EnterHead(c *HeadContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// EnterAccess is called when entering the access production.
	EnterAccess(c *AccessContext)

	// EnterSymbols is called when entering the symbols production.
	EnterSymbols(c *SymbolsContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterLocks is called when entering the locks production.
	EnterLocks(c *LocksContext)

	// EnterStrict is called when entering the strict production.
	EnterStrict(c *StrictContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterExpand is called when entering the expand production.
	EnterExpand(c *ExpandContext)

	// EnterDeltalist is called when entering the deltalist production.
	EnterDeltalist(c *DeltalistContext)

	// EnterDelta is called when entering the delta production.
	EnterDelta(c *DeltaContext)

	// EnterDelta_date is called when entering the delta_date production.
	EnterDelta_date(c *Delta_dateContext)

	// EnterDelta_author is called when entering the delta_author production.
	EnterDelta_author(c *Delta_authorContext)

	// EnterDelta_state is called when entering the delta_state production.
	EnterDelta_state(c *Delta_stateContext)

	// EnterDelta_branches is called when entering the delta_branches production.
	EnterDelta_branches(c *Delta_branchesContext)

	// EnterDelta_next is called when entering the delta_next production.
	EnterDelta_next(c *Delta_nextContext)

	// EnterDesc is called when entering the desc production.
	EnterDesc(c *DescContext)

	// EnterDeltatextlist is called when entering the deltatextlist production.
	EnterDeltatextlist(c *DeltatextlistContext)

	// EnterDeltatext is called when entering the deltatext production.
	EnterDeltatext(c *DeltatextContext)

	// EnterDeltatext_log is called when entering the deltatext_log production.
	EnterDeltatext_log(c *Deltatext_logContext)

	// EnterDeltatext_text is called when entering the deltatext_text production.
	EnterDeltatext_text(c *Deltatext_textContext)

	// EnterNewphrase is called when entering the newphrase production.
	EnterNewphrase(c *NewphraseContext)

	// ExitRcstext is called when exiting the rcstext production.
	ExitRcstext(c *RcstextContext)

	// ExitRcsheader is called when exiting the rcsheader production.
	ExitRcsheader(c *RcsheaderContext)

	// ExitRcsrevisions is called when exiting the rcsrevisions production.
	ExitRcsrevisions(c *RcsrevisionsContext)

	// ExitAdmin is called when exiting the admin production.
	ExitAdmin(c *AdminContext)

	// ExitHead is called when exiting the head production.
	ExitHead(c *HeadContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)

	// ExitAccess is called when exiting the access production.
	ExitAccess(c *AccessContext)

	// ExitSymbols is called when exiting the symbols production.
	ExitSymbols(c *SymbolsContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitLocks is called when exiting the locks production.
	ExitLocks(c *LocksContext)

	// ExitStrict is called when exiting the strict production.
	ExitStrict(c *StrictContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitExpand is called when exiting the expand production.
	ExitExpand(c *ExpandContext)

	// ExitDeltalist is called when exiting the deltalist production.
	ExitDeltalist(c *DeltalistContext)

	// ExitDelta is called when exiting the delta production.
	ExitDelta(c *DeltaContext)

	// ExitDelta_date is called when exiting the delta_date production.
	ExitDelta_date(c *Delta_dateContext)

	// ExitDelta_author is called when exiting the delta_author production.
	ExitDelta_author(c *Delta_authorContext)

	// ExitDelta_state is called when exiting the delta_state production.
	ExitDelta_state(c *Delta_stateContext)

	// ExitDelta_branches is called when exiting the delta_branches production.
	ExitDelta_branches(c *Delta_branchesContext)

	// ExitDelta_next is called when exiting the delta_next production.
	ExitDelta_next(c *Delta_nextContext)

	// ExitDesc is called when exiting the desc production.
	ExitDesc(c *DescContext)

	// ExitDeltatextlist is called when exiting the deltatextlist production.
	ExitDeltatextlist(c *DeltatextlistContext)

	// ExitDeltatext is called when exiting the deltatext production.
	ExitDeltatext(c *DeltatextContext)

	// ExitDeltatext_log is called when exiting the deltatext_log production.
	ExitDeltatext_log(c *Deltatext_logContext)

	// ExitDeltatext_text is called when exiting the deltatext_text production.
	ExitDeltatext_text(c *Deltatext_textContext)

	// ExitNewphrase is called when exiting the newphrase production.
	ExitNewphrase(c *NewphraseContext)
}
