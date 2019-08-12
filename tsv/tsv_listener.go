// Code generated from tsv.g4 by ANTLR 4.7.2. DO NOT EDIT.

package tsv // tsv
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tsvListener is a complete listener for a parse tree produced by tsvParser.
type tsvListener interface {
	antlr.ParseTreeListener

	// EnterTsvFile is called when entering the tsvFile production.
	EnterTsvFile(c *TsvFileContext)

	// EnterHdr is called when entering the hdr production.
	EnterHdr(c *HdrContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// ExitTsvFile is called when exiting the tsvFile production.
	ExitTsvFile(c *TsvFileContext)

	// ExitHdr is called when exiting the hdr production.
	ExitHdr(c *HdrContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)
}
