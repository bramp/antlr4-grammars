// Code generated from DOT.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dot // DOT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DOTListener is a complete listener for a parse tree produced by DOTParser.
type DOTListener interface {
	antlr.ParseTreeListener

	// EnterGraph is called when entering the graph production.
	EnterGraph(c *GraphContext)

	// EnterStmt_list is called when entering the stmt_list production.
	EnterStmt_list(c *Stmt_listContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAttr_stmt is called when entering the attr_stmt production.
	EnterAttr_stmt(c *Attr_stmtContext)

	// EnterAttr_list is called when entering the attr_list production.
	EnterAttr_list(c *Attr_listContext)

	// EnterA_list is called when entering the a_list production.
	EnterA_list(c *A_listContext)

	// EnterEdge_stmt is called when entering the edge_stmt production.
	EnterEdge_stmt(c *Edge_stmtContext)

	// EnterEdgeRHS is called when entering the edgeRHS production.
	EnterEdgeRHS(c *EdgeRHSContext)

	// EnterEdgeop is called when entering the edgeop production.
	EnterEdgeop(c *EdgeopContext)

	// EnterNode_stmt is called when entering the node_stmt production.
	EnterNode_stmt(c *Node_stmtContext)

	// EnterNode_id is called when entering the node_id production.
	EnterNode_id(c *Node_idContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterSubgraph is called when entering the subgraph production.
	EnterSubgraph(c *SubgraphContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitGraph is called when exiting the graph production.
	ExitGraph(c *GraphContext)

	// ExitStmt_list is called when exiting the stmt_list production.
	ExitStmt_list(c *Stmt_listContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAttr_stmt is called when exiting the attr_stmt production.
	ExitAttr_stmt(c *Attr_stmtContext)

	// ExitAttr_list is called when exiting the attr_list production.
	ExitAttr_list(c *Attr_listContext)

	// ExitA_list is called when exiting the a_list production.
	ExitA_list(c *A_listContext)

	// ExitEdge_stmt is called when exiting the edge_stmt production.
	ExitEdge_stmt(c *Edge_stmtContext)

	// ExitEdgeRHS is called when exiting the edgeRHS production.
	ExitEdgeRHS(c *EdgeRHSContext)

	// ExitEdgeop is called when exiting the edgeop production.
	ExitEdgeop(c *EdgeopContext)

	// ExitNode_stmt is called when exiting the node_stmt production.
	ExitNode_stmt(c *Node_stmtContext)

	// ExitNode_id is called when exiting the node_id production.
	ExitNode_id(c *Node_idContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitSubgraph is called when exiting the subgraph production.
	ExitSubgraph(c *SubgraphContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}
