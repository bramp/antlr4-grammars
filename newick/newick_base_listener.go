// Code generated from newick.g4 by ANTLR 4.9.3. DO NOT EDIT.

package newick // newick
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasenewickListener is a complete listener for a parse tree produced by newickParser.
type BasenewickListener struct{}

var _ newickListener = &BasenewickListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasenewickListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasenewickListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasenewickListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasenewickListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTree_ is called when production tree_ is entered.
func (s *BasenewickListener) EnterTree_(ctx *Tree_Context) {}

// ExitTree_ is called when production tree_ is exited.
func (s *BasenewickListener) ExitTree_(ctx *Tree_Context) {}

// EnterRootLeaf is called when production rootLeaf is entered.
func (s *BasenewickListener) EnterRootLeaf(ctx *RootLeafContext) {}

// ExitRootLeaf is called when production rootLeaf is exited.
func (s *BasenewickListener) ExitRootLeaf(ctx *RootLeafContext) {}

// EnterRootInternal is called when production rootInternal is entered.
func (s *BasenewickListener) EnterRootInternal(ctx *RootInternalContext) {}

// ExitRootInternal is called when production rootInternal is exited.
func (s *BasenewickListener) ExitRootInternal(ctx *RootInternalContext) {}

// EnterSubtree is called when production subtree is entered.
func (s *BasenewickListener) EnterSubtree(ctx *SubtreeContext) {}

// ExitSubtree is called when production subtree is exited.
func (s *BasenewickListener) ExitSubtree(ctx *SubtreeContext) {}

// EnterLeaf is called when production leaf is entered.
func (s *BasenewickListener) EnterLeaf(ctx *LeafContext) {}

// ExitLeaf is called when production leaf is exited.
func (s *BasenewickListener) ExitLeaf(ctx *LeafContext) {}

// EnterInternal_ is called when production internal_ is entered.
func (s *BasenewickListener) EnterInternal_(ctx *Internal_Context) {}

// ExitInternal_ is called when production internal_ is exited.
func (s *BasenewickListener) ExitInternal_(ctx *Internal_Context) {}

// EnterBranchSet is called when production branchSet is entered.
func (s *BasenewickListener) EnterBranchSet(ctx *BranchSetContext) {}

// ExitBranchSet is called when production branchSet is exited.
func (s *BasenewickListener) ExitBranchSet(ctx *BranchSetContext) {}

// EnterBranch is called when production branch is entered.
func (s *BasenewickListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BasenewickListener) ExitBranch(ctx *BranchContext) {}

// EnterName is called when production name is entered.
func (s *BasenewickListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasenewickListener) ExitName(ctx *NameContext) {}

// EnterLength is called when production length is entered.
func (s *BasenewickListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BasenewickListener) ExitLength(ctx *LengthContext) {}
