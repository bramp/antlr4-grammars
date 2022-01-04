// Code generated from wln.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wln // wln
import "github.com/antlr/antlr4/runtime/Go/antlr"

// wlnListener is a complete listener for a parse tree produced by wlnParser.
type wlnListener interface {
	antlr.ParseTreeListener

	// EnterWln is called when entering the wln production.
	EnterWln(c *WlnContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// ExitWln is called when exiting the wln production.
	ExitWln(c *WlnContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)
}
