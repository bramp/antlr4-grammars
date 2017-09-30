// Generated from RCS.g4 by ANTLR 4.7.

package rcs // RCS
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RCSParser.
type RCSVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RCSParser#rcstext.
	VisitRcstext(ctx *RcstextContext) interface{}

	// Visit a parse tree produced by RCSParser#rcsheader.
	VisitRcsheader(ctx *RcsheaderContext) interface{}

	// Visit a parse tree produced by RCSParser#rcsrevisions.
	VisitRcsrevisions(ctx *RcsrevisionsContext) interface{}

	// Visit a parse tree produced by RCSParser#admin.
	VisitAdmin(ctx *AdminContext) interface{}

	// Visit a parse tree produced by RCSParser#head.
	VisitHead(ctx *HeadContext) interface{}

	// Visit a parse tree produced by RCSParser#branch.
	VisitBranch(ctx *BranchContext) interface{}

	// Visit a parse tree produced by RCSParser#access.
	VisitAccess(ctx *AccessContext) interface{}

	// Visit a parse tree produced by RCSParser#symbols.
	VisitSymbols(ctx *SymbolsContext) interface{}

	// Visit a parse tree produced by RCSParser#tags.
	VisitTags(ctx *TagsContext) interface{}

	// Visit a parse tree produced by RCSParser#locks.
	VisitLocks(ctx *LocksContext) interface{}

	// Visit a parse tree produced by RCSParser#strict.
	VisitStrict(ctx *StrictContext) interface{}

	// Visit a parse tree produced by RCSParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by RCSParser#expand.
	VisitExpand(ctx *ExpandContext) interface{}

	// Visit a parse tree produced by RCSParser#deltalist.
	VisitDeltalist(ctx *DeltalistContext) interface{}

	// Visit a parse tree produced by RCSParser#delta.
	VisitDelta(ctx *DeltaContext) interface{}

	// Visit a parse tree produced by RCSParser#delta_date.
	VisitDelta_date(ctx *Delta_dateContext) interface{}

	// Visit a parse tree produced by RCSParser#delta_author.
	VisitDelta_author(ctx *Delta_authorContext) interface{}

	// Visit a parse tree produced by RCSParser#delta_state.
	VisitDelta_state(ctx *Delta_stateContext) interface{}

	// Visit a parse tree produced by RCSParser#delta_branches.
	VisitDelta_branches(ctx *Delta_branchesContext) interface{}

	// Visit a parse tree produced by RCSParser#delta_next.
	VisitDelta_next(ctx *Delta_nextContext) interface{}

	// Visit a parse tree produced by RCSParser#desc.
	VisitDesc(ctx *DescContext) interface{}

	// Visit a parse tree produced by RCSParser#deltatextlist.
	VisitDeltatextlist(ctx *DeltatextlistContext) interface{}

	// Visit a parse tree produced by RCSParser#deltatext.
	VisitDeltatext(ctx *DeltatextContext) interface{}

	// Visit a parse tree produced by RCSParser#deltatext_log.
	VisitDeltatext_log(ctx *Deltatext_logContext) interface{}

	// Visit a parse tree produced by RCSParser#deltatext_text.
	VisitDeltatext_text(ctx *Deltatext_textContext) interface{}

	// Visit a parse tree produced by RCSParser#newphrase.
	VisitNewphrase(ctx *NewphraseContext) interface{}
}
