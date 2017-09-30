// Generated from RCS.g4 by ANTLR 4.7.

package rcs // RCS
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRCSListener is a complete listener for a parse tree produced by RCSParser.
type BaseRCSListener struct{}

var _ RCSListener = &BaseRCSListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRCSListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRCSListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRCSListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRCSListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRcstext is called when production rcstext is entered.
func (s *BaseRCSListener) EnterRcstext(ctx *RcstextContext) {}

// ExitRcstext is called when production rcstext is exited.
func (s *BaseRCSListener) ExitRcstext(ctx *RcstextContext) {}

// EnterRcsheader is called when production rcsheader is entered.
func (s *BaseRCSListener) EnterRcsheader(ctx *RcsheaderContext) {}

// ExitRcsheader is called when production rcsheader is exited.
func (s *BaseRCSListener) ExitRcsheader(ctx *RcsheaderContext) {}

// EnterRcsrevisions is called when production rcsrevisions is entered.
func (s *BaseRCSListener) EnterRcsrevisions(ctx *RcsrevisionsContext) {}

// ExitRcsrevisions is called when production rcsrevisions is exited.
func (s *BaseRCSListener) ExitRcsrevisions(ctx *RcsrevisionsContext) {}

// EnterAdmin is called when production admin is entered.
func (s *BaseRCSListener) EnterAdmin(ctx *AdminContext) {}

// ExitAdmin is called when production admin is exited.
func (s *BaseRCSListener) ExitAdmin(ctx *AdminContext) {}

// EnterHead is called when production head is entered.
func (s *BaseRCSListener) EnterHead(ctx *HeadContext) {}

// ExitHead is called when production head is exited.
func (s *BaseRCSListener) ExitHead(ctx *HeadContext) {}

// EnterBranch is called when production branch is entered.
func (s *BaseRCSListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BaseRCSListener) ExitBranch(ctx *BranchContext) {}

// EnterAccess is called when production access is entered.
func (s *BaseRCSListener) EnterAccess(ctx *AccessContext) {}

// ExitAccess is called when production access is exited.
func (s *BaseRCSListener) ExitAccess(ctx *AccessContext) {}

// EnterSymbols is called when production symbols is entered.
func (s *BaseRCSListener) EnterSymbols(ctx *SymbolsContext) {}

// ExitSymbols is called when production symbols is exited.
func (s *BaseRCSListener) ExitSymbols(ctx *SymbolsContext) {}

// EnterTags is called when production tags is entered.
func (s *BaseRCSListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BaseRCSListener) ExitTags(ctx *TagsContext) {}

// EnterLocks is called when production locks is entered.
func (s *BaseRCSListener) EnterLocks(ctx *LocksContext) {}

// ExitLocks is called when production locks is exited.
func (s *BaseRCSListener) ExitLocks(ctx *LocksContext) {}

// EnterStrict is called when production strict is entered.
func (s *BaseRCSListener) EnterStrict(ctx *StrictContext) {}

// ExitStrict is called when production strict is exited.
func (s *BaseRCSListener) ExitStrict(ctx *StrictContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseRCSListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseRCSListener) ExitComment(ctx *CommentContext) {}

// EnterExpand is called when production expand is entered.
func (s *BaseRCSListener) EnterExpand(ctx *ExpandContext) {}

// ExitExpand is called when production expand is exited.
func (s *BaseRCSListener) ExitExpand(ctx *ExpandContext) {}

// EnterDeltalist is called when production deltalist is entered.
func (s *BaseRCSListener) EnterDeltalist(ctx *DeltalistContext) {}

// ExitDeltalist is called when production deltalist is exited.
func (s *BaseRCSListener) ExitDeltalist(ctx *DeltalistContext) {}

// EnterDelta is called when production delta is entered.
func (s *BaseRCSListener) EnterDelta(ctx *DeltaContext) {}

// ExitDelta is called when production delta is exited.
func (s *BaseRCSListener) ExitDelta(ctx *DeltaContext) {}

// EnterDelta_date is called when production delta_date is entered.
func (s *BaseRCSListener) EnterDelta_date(ctx *Delta_dateContext) {}

// ExitDelta_date is called when production delta_date is exited.
func (s *BaseRCSListener) ExitDelta_date(ctx *Delta_dateContext) {}

// EnterDelta_author is called when production delta_author is entered.
func (s *BaseRCSListener) EnterDelta_author(ctx *Delta_authorContext) {}

// ExitDelta_author is called when production delta_author is exited.
func (s *BaseRCSListener) ExitDelta_author(ctx *Delta_authorContext) {}

// EnterDelta_state is called when production delta_state is entered.
func (s *BaseRCSListener) EnterDelta_state(ctx *Delta_stateContext) {}

// ExitDelta_state is called when production delta_state is exited.
func (s *BaseRCSListener) ExitDelta_state(ctx *Delta_stateContext) {}

// EnterDelta_branches is called when production delta_branches is entered.
func (s *BaseRCSListener) EnterDelta_branches(ctx *Delta_branchesContext) {}

// ExitDelta_branches is called when production delta_branches is exited.
func (s *BaseRCSListener) ExitDelta_branches(ctx *Delta_branchesContext) {}

// EnterDelta_next is called when production delta_next is entered.
func (s *BaseRCSListener) EnterDelta_next(ctx *Delta_nextContext) {}

// ExitDelta_next is called when production delta_next is exited.
func (s *BaseRCSListener) ExitDelta_next(ctx *Delta_nextContext) {}

// EnterDesc is called when production desc is entered.
func (s *BaseRCSListener) EnterDesc(ctx *DescContext) {}

// ExitDesc is called when production desc is exited.
func (s *BaseRCSListener) ExitDesc(ctx *DescContext) {}

// EnterDeltatextlist is called when production deltatextlist is entered.
func (s *BaseRCSListener) EnterDeltatextlist(ctx *DeltatextlistContext) {}

// ExitDeltatextlist is called when production deltatextlist is exited.
func (s *BaseRCSListener) ExitDeltatextlist(ctx *DeltatextlistContext) {}

// EnterDeltatext is called when production deltatext is entered.
func (s *BaseRCSListener) EnterDeltatext(ctx *DeltatextContext) {}

// ExitDeltatext is called when production deltatext is exited.
func (s *BaseRCSListener) ExitDeltatext(ctx *DeltatextContext) {}

// EnterDeltatext_log is called when production deltatext_log is entered.
func (s *BaseRCSListener) EnterDeltatext_log(ctx *Deltatext_logContext) {}

// ExitDeltatext_log is called when production deltatext_log is exited.
func (s *BaseRCSListener) ExitDeltatext_log(ctx *Deltatext_logContext) {}

// EnterDeltatext_text is called when production deltatext_text is entered.
func (s *BaseRCSListener) EnterDeltatext_text(ctx *Deltatext_textContext) {}

// ExitDeltatext_text is called when production deltatext_text is exited.
func (s *BaseRCSListener) ExitDeltatext_text(ctx *Deltatext_textContext) {}

// EnterNewphrase is called when production newphrase is entered.
func (s *BaseRCSListener) EnterNewphrase(ctx *NewphraseContext) {}

// ExitNewphrase is called when production newphrase is exited.
func (s *BaseRCSListener) ExitNewphrase(ctx *NewphraseContext) {}
