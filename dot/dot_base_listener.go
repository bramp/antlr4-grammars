// Code generated from DOT.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dot // DOT
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDOTListener is a complete listener for a parse tree produced by DOTParser.
type BaseDOTListener struct{}

var _ DOTListener = &BaseDOTListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDOTListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDOTListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDOTListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDOTListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGraph is called when production graph is entered.
func (s *BaseDOTListener) EnterGraph(ctx *GraphContext) {}

// ExitGraph is called when production graph is exited.
func (s *BaseDOTListener) ExitGraph(ctx *GraphContext) {}

// EnterStmt_list is called when production stmt_list is entered.
func (s *BaseDOTListener) EnterStmt_list(ctx *Stmt_listContext) {}

// ExitStmt_list is called when production stmt_list is exited.
func (s *BaseDOTListener) ExitStmt_list(ctx *Stmt_listContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseDOTListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseDOTListener) ExitStmt(ctx *StmtContext) {}

// EnterAttr_stmt is called when production attr_stmt is entered.
func (s *BaseDOTListener) EnterAttr_stmt(ctx *Attr_stmtContext) {}

// ExitAttr_stmt is called when production attr_stmt is exited.
func (s *BaseDOTListener) ExitAttr_stmt(ctx *Attr_stmtContext) {}

// EnterAttr_list is called when production attr_list is entered.
func (s *BaseDOTListener) EnterAttr_list(ctx *Attr_listContext) {}

// ExitAttr_list is called when production attr_list is exited.
func (s *BaseDOTListener) ExitAttr_list(ctx *Attr_listContext) {}

// EnterA_list is called when production a_list is entered.
func (s *BaseDOTListener) EnterA_list(ctx *A_listContext) {}

// ExitA_list is called when production a_list is exited.
func (s *BaseDOTListener) ExitA_list(ctx *A_listContext) {}

// EnterEdge_stmt is called when production edge_stmt is entered.
func (s *BaseDOTListener) EnterEdge_stmt(ctx *Edge_stmtContext) {}

// ExitEdge_stmt is called when production edge_stmt is exited.
func (s *BaseDOTListener) ExitEdge_stmt(ctx *Edge_stmtContext) {}

// EnterEdgeRHS is called when production edgeRHS is entered.
func (s *BaseDOTListener) EnterEdgeRHS(ctx *EdgeRHSContext) {}

// ExitEdgeRHS is called when production edgeRHS is exited.
func (s *BaseDOTListener) ExitEdgeRHS(ctx *EdgeRHSContext) {}

// EnterEdgeop is called when production edgeop is entered.
func (s *BaseDOTListener) EnterEdgeop(ctx *EdgeopContext) {}

// ExitEdgeop is called when production edgeop is exited.
func (s *BaseDOTListener) ExitEdgeop(ctx *EdgeopContext) {}

// EnterNode_stmt is called when production node_stmt is entered.
func (s *BaseDOTListener) EnterNode_stmt(ctx *Node_stmtContext) {}

// ExitNode_stmt is called when production node_stmt is exited.
func (s *BaseDOTListener) ExitNode_stmt(ctx *Node_stmtContext) {}

// EnterNode_id is called when production node_id is entered.
func (s *BaseDOTListener) EnterNode_id(ctx *Node_idContext) {}

// ExitNode_id is called when production node_id is exited.
func (s *BaseDOTListener) ExitNode_id(ctx *Node_idContext) {}

// EnterPort is called when production port is entered.
func (s *BaseDOTListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseDOTListener) ExitPort(ctx *PortContext) {}

// EnterSubgraph is called when production subgraph is entered.
func (s *BaseDOTListener) EnterSubgraph(ctx *SubgraphContext) {}

// ExitSubgraph is called when production subgraph is exited.
func (s *BaseDOTListener) ExitSubgraph(ctx *SubgraphContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseDOTListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseDOTListener) ExitId_(ctx *Id_Context) {}
