// Code generated from vmf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vmf // vmf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// vmfListener is a complete listener for a parse tree produced by vmfParser.
type vmfListener interface {
	antlr.ParseTreeListener

	// EnterVmf is called when entering the vmf production.
	EnterVmf(c *VmfContext)

	// EnterKeyvalue is called when entering the keyvalue production.
	EnterKeyvalue(c *KeyvalueContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterAtomicvalue is called when entering the atomicvalue production.
	EnterAtomicvalue(c *AtomicvalueContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// EnterListvalue is called when entering the listvalue production.
	EnterListvalue(c *ListvalueContext)

	// ExitVmf is called when exiting the vmf production.
	ExitVmf(c *VmfContext)

	// ExitKeyvalue is called when exiting the keyvalue production.
	ExitKeyvalue(c *KeyvalueContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitAtomicvalue is called when exiting the atomicvalue production.
	ExitAtomicvalue(c *AtomicvalueContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)

	// ExitListvalue is called when exiting the listvalue production.
	ExitListvalue(c *ListvalueContext)
}
