// Code generated from infosapient.g4 by ANTLR 4.9.3. DO NOT EDIT.

package infosapient // infosapient
import "github.com/antlr/antlr4/runtime/Go/antlr"

// infosapientListener is a complete listener for a parse tree produced by infosapientParser.
type infosapientListener interface {
	antlr.ParseTreeListener

	// EnterConditionalRule is called when entering the conditionalRule production.
	EnterConditionalRule(c *ConditionalRuleContext)

	// EnterPremise is called when entering the premise production.
	EnterPremise(c *PremiseContext)

	// EnterConsequent is called when entering the consequent production.
	EnterConsequent(c *ConsequentContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAttribClause is called when entering the attribClause production.
	EnterAttribClause(c *AttribClauseContext)

	// EnterHedgeCollection is called when entering the hedgeCollection production.
	EnterHedgeCollection(c *HedgeCollectionContext)

	// EnterRestrictionHedge is called when entering the restrictionHedge production.
	EnterRestrictionHedge(c *RestrictionHedgeContext)

	// EnterHedge is called when entering the hedge production.
	EnterHedge(c *HedgeContext)

	// EnterOperator_ is called when entering the operator_ production.
	EnterOperator_(c *Operator_Context)

	// EnterNLiteral is called when entering the nLiteral production.
	EnterNLiteral(c *NLiteralContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// ExitConditionalRule is called when exiting the conditionalRule production.
	ExitConditionalRule(c *ConditionalRuleContext)

	// ExitPremise is called when exiting the premise production.
	ExitPremise(c *PremiseContext)

	// ExitConsequent is called when exiting the consequent production.
	ExitConsequent(c *ConsequentContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAttribClause is called when exiting the attribClause production.
	ExitAttribClause(c *AttribClauseContext)

	// ExitHedgeCollection is called when exiting the hedgeCollection production.
	ExitHedgeCollection(c *HedgeCollectionContext)

	// ExitRestrictionHedge is called when exiting the restrictionHedge production.
	ExitRestrictionHedge(c *RestrictionHedgeContext)

	// ExitHedge is called when exiting the hedge production.
	ExitHedge(c *HedgeContext)

	// ExitOperator_ is called when exiting the operator_ production.
	ExitOperator_(c *Operator_Context)

	// ExitNLiteral is called when exiting the nLiteral production.
	ExitNLiteral(c *NLiteralContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)
}
