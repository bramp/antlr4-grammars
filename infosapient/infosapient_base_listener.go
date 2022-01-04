// Code generated from infosapient.g4 by ANTLR 4.9.3. DO NOT EDIT.

package infosapient // infosapient
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseinfosapientListener is a complete listener for a parse tree produced by infosapientParser.
type BaseinfosapientListener struct{}

var _ infosapientListener = &BaseinfosapientListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseinfosapientListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseinfosapientListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseinfosapientListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseinfosapientListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConditionalRule is called when production conditionalRule is entered.
func (s *BaseinfosapientListener) EnterConditionalRule(ctx *ConditionalRuleContext) {}

// ExitConditionalRule is called when production conditionalRule is exited.
func (s *BaseinfosapientListener) ExitConditionalRule(ctx *ConditionalRuleContext) {}

// EnterPremise is called when production premise is entered.
func (s *BaseinfosapientListener) EnterPremise(ctx *PremiseContext) {}

// ExitPremise is called when production premise is exited.
func (s *BaseinfosapientListener) ExitPremise(ctx *PremiseContext) {}

// EnterConsequent is called when production consequent is entered.
func (s *BaseinfosapientListener) EnterConsequent(ctx *ConsequentContext) {}

// ExitConsequent is called when production consequent is exited.
func (s *BaseinfosapientListener) ExitConsequent(ctx *ConsequentContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseinfosapientListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseinfosapientListener) ExitClause(ctx *ClauseContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseinfosapientListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseinfosapientListener) ExitExpr(ctx *ExprContext) {}

// EnterAttribClause is called when production attribClause is entered.
func (s *BaseinfosapientListener) EnterAttribClause(ctx *AttribClauseContext) {}

// ExitAttribClause is called when production attribClause is exited.
func (s *BaseinfosapientListener) ExitAttribClause(ctx *AttribClauseContext) {}

// EnterHedgeCollection is called when production hedgeCollection is entered.
func (s *BaseinfosapientListener) EnterHedgeCollection(ctx *HedgeCollectionContext) {}

// ExitHedgeCollection is called when production hedgeCollection is exited.
func (s *BaseinfosapientListener) ExitHedgeCollection(ctx *HedgeCollectionContext) {}

// EnterRestrictionHedge is called when production restrictionHedge is entered.
func (s *BaseinfosapientListener) EnterRestrictionHedge(ctx *RestrictionHedgeContext) {}

// ExitRestrictionHedge is called when production restrictionHedge is exited.
func (s *BaseinfosapientListener) ExitRestrictionHedge(ctx *RestrictionHedgeContext) {}

// EnterHedge is called when production hedge is entered.
func (s *BaseinfosapientListener) EnterHedge(ctx *HedgeContext) {}

// ExitHedge is called when production hedge is exited.
func (s *BaseinfosapientListener) ExitHedge(ctx *HedgeContext) {}

// EnterOperator_ is called when production operator_ is entered.
func (s *BaseinfosapientListener) EnterOperator_(ctx *Operator_Context) {}

// ExitOperator_ is called when production operator_ is exited.
func (s *BaseinfosapientListener) ExitOperator_(ctx *Operator_Context) {}

// EnterNLiteral is called when production nLiteral is entered.
func (s *BaseinfosapientListener) EnterNLiteral(ctx *NLiteralContext) {}

// ExitNLiteral is called when production nLiteral is exited.
func (s *BaseinfosapientListener) ExitNLiteral(ctx *NLiteralContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseinfosapientListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseinfosapientListener) ExitId_(ctx *Id_Context) {}
