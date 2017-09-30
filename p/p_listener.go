// Generated from p.g4 by ANTLR 4.7.

package p // p
import "github.com/antlr/antlr4/runtime/Go/antlr"

// pListener is a complete listener for a parse tree produced by pParser.
type pListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterIterate is called when entering the iterate production.
	EnterIterate(c *IterateContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitIterate is called when exiting the iterate production.
	ExitIterate(c *IterateContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}
