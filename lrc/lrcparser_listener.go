// Code generated from lrcParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lrc // lrcParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lrcParserListener is a complete listener for a parse tree produced by lrcParser.
type lrcParserListener interface {
	antlr.ParseTreeListener

	// EnterLrc is called when entering the lrc production.
	EnterLrc(c *LrcContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// ExitLrc is called when exiting the lrc production.
	ExitLrc(c *LrcContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)
}
