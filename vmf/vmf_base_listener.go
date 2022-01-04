// Code generated from vmf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vmf // vmf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasevmfListener is a complete listener for a parse tree produced by vmfParser.
type BasevmfListener struct{}

var _ vmfListener = &BasevmfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasevmfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasevmfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasevmfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasevmfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterVmf is called when production vmf is entered.
func (s *BasevmfListener) EnterVmf(ctx *VmfContext) {}

// ExitVmf is called when production vmf is exited.
func (s *BasevmfListener) ExitVmf(ctx *VmfContext) {}

// EnterKeyvalue is called when production keyvalue is entered.
func (s *BasevmfListener) EnterKeyvalue(ctx *KeyvalueContext) {}

// ExitKeyvalue is called when production keyvalue is exited.
func (s *BasevmfListener) ExitKeyvalue(ctx *KeyvalueContext) {}

// EnterKey is called when production key is entered.
func (s *BasevmfListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasevmfListener) ExitKey(ctx *KeyContext) {}

// EnterAtomicvalue is called when production atomicvalue is entered.
func (s *BasevmfListener) EnterAtomicvalue(ctx *AtomicvalueContext) {}

// ExitAtomicvalue is called when production atomicvalue is exited.
func (s *BasevmfListener) ExitAtomicvalue(ctx *AtomicvalueContext) {}

// EnterVal is called when production val is entered.
func (s *BasevmfListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BasevmfListener) ExitVal(ctx *ValContext) {}

// EnterListvalue is called when production listvalue is entered.
func (s *BasevmfListener) EnterListvalue(ctx *ListvalueContext) {}

// ExitListvalue is called when production listvalue is exited.
func (s *BasevmfListener) ExitListvalue(ctx *ListvalueContext) {}
