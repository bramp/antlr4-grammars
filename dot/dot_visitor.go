// Generated from DOT.g4 by ANTLR 4.7.

package dot // DOT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DOTParser.
type DOTVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DOTParser#graph.
	VisitGraph(ctx *GraphContext) interface{}

	// Visit a parse tree produced by DOTParser#stmt_list.
	VisitStmt_list(ctx *Stmt_listContext) interface{}

	// Visit a parse tree produced by DOTParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by DOTParser#attr_stmt.
	VisitAttr_stmt(ctx *Attr_stmtContext) interface{}

	// Visit a parse tree produced by DOTParser#attr_list.
	VisitAttr_list(ctx *Attr_listContext) interface{}

	// Visit a parse tree produced by DOTParser#a_list.
	VisitA_list(ctx *A_listContext) interface{}

	// Visit a parse tree produced by DOTParser#edge_stmt.
	VisitEdge_stmt(ctx *Edge_stmtContext) interface{}

	// Visit a parse tree produced by DOTParser#edgeRHS.
	VisitEdgeRHS(ctx *EdgeRHSContext) interface{}

	// Visit a parse tree produced by DOTParser#edgeop.
	VisitEdgeop(ctx *EdgeopContext) interface{}

	// Visit a parse tree produced by DOTParser#node_stmt.
	VisitNode_stmt(ctx *Node_stmtContext) interface{}

	// Visit a parse tree produced by DOTParser#node_id.
	VisitNode_id(ctx *Node_idContext) interface{}

	// Visit a parse tree produced by DOTParser#port.
	VisitPort(ctx *PortContext) interface{}

	// Visit a parse tree produced by DOTParser#subgraph.
	VisitSubgraph(ctx *SubgraphContext) interface{}

	// Visit a parse tree produced by DOTParser#id.
	VisitId(ctx *IdContext) interface{}
}
