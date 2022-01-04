// Code generated from edn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package edn // edn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ednListener is a complete listener for a parse tree produced by ednParser.
type ednListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterMap_ is called when entering the map_ production.
	EnterMap_(c *Map_Context)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitMap_ is called when exiting the map_ production.
	ExitMap_(c *Map_Context)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)
}
