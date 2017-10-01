// Generated from DOT.g4 by ANTLR 4.7.

package dot // DOT
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDOTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDOTVisitor) VisitGraph(ctx *GraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitStmt_list(ctx *Stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitAttr_stmt(ctx *Attr_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitAttr_list(ctx *Attr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitA_list(ctx *A_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitEdge_stmt(ctx *Edge_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitEdgeRHS(ctx *EdgeRHSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitEdgeop(ctx *EdgeopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitNode_stmt(ctx *Node_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitNode_id(ctx *Node_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitPort(ctx *PortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitSubgraph(ctx *SubgraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDOTVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}
