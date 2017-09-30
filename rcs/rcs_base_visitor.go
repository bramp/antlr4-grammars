// Generated from RCS.g4 by ANTLR 4.7.

package rcs // RCS
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRCSVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRCSVisitor) VisitRcstext(ctx *RcstextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitRcsheader(ctx *RcsheaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitRcsrevisions(ctx *RcsrevisionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitAdmin(ctx *AdminContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitHead(ctx *HeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitBranch(ctx *BranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitAccess(ctx *AccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitSymbols(ctx *SymbolsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitLocks(ctx *LocksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitStrict(ctx *StrictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitExpand(ctx *ExpandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDeltalist(ctx *DeltalistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta(ctx *DeltaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta_date(ctx *Delta_dateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta_author(ctx *Delta_authorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta_state(ctx *Delta_stateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta_branches(ctx *Delta_branchesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDelta_next(ctx *Delta_nextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDesc(ctx *DescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDeltatextlist(ctx *DeltatextlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDeltatext(ctx *DeltatextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDeltatext_log(ctx *Deltatext_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitDeltatext_text(ctx *Deltatext_textContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRCSVisitor) VisitNewphrase(ctx *NewphraseContext) interface{} {
	return v.VisitChildren(ctx)
}
