// Code generated from newick.g4 by ANTLR 4.9.3. DO NOT EDIT.

package newick // newick
import "github.com/antlr/antlr4/runtime/Go/antlr"

// newickListener is a complete listener for a parse tree produced by newickParser.
type newickListener interface {
	antlr.ParseTreeListener

	// EnterTree_ is called when entering the tree_ production.
	EnterTree_(c *Tree_Context)

	// EnterRootLeaf is called when entering the rootLeaf production.
	EnterRootLeaf(c *RootLeafContext)

	// EnterRootInternal is called when entering the rootInternal production.
	EnterRootInternal(c *RootInternalContext)

	// EnterSubtree is called when entering the subtree production.
	EnterSubtree(c *SubtreeContext)

	// EnterLeaf is called when entering the leaf production.
	EnterLeaf(c *LeafContext)

	// EnterInternal_ is called when entering the internal_ production.
	EnterInternal_(c *Internal_Context)

	// EnterBranchSet is called when entering the branchSet production.
	EnterBranchSet(c *BranchSetContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// ExitTree_ is called when exiting the tree_ production.
	ExitTree_(c *Tree_Context)

	// ExitRootLeaf is called when exiting the rootLeaf production.
	ExitRootLeaf(c *RootLeafContext)

	// ExitRootInternal is called when exiting the rootInternal production.
	ExitRootInternal(c *RootInternalContext)

	// ExitSubtree is called when exiting the subtree production.
	ExitSubtree(c *SubtreeContext)

	// ExitLeaf is called when exiting the leaf production.
	ExitLeaf(c *LeafContext)

	// ExitInternal_ is called when exiting the internal_ production.
	ExitInternal_(c *Internal_Context)

	// ExitBranchSet is called when exiting the branchSet production.
	ExitBranchSet(c *BranchSetContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)
}
