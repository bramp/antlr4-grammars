// Code generated from brainflak.g4 by ANTLR 4.9.3. DO NOT EDIT.

package brainflak // brainflak
import "github.com/antlr/antlr4/runtime/Go/antlr"

// brainflakListener is a complete listener for a parse tree produced by brainflakParser.
type brainflakListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterParenstmt is called when entering the parenstmt production.
	EnterParenstmt(c *ParenstmtContext)

	// EnterBracestmt is called when entering the bracestmt production.
	EnterBracestmt(c *BracestmtContext)

	// EnterBracketstmt is called when entering the bracketstmt production.
	EnterBracketstmt(c *BracketstmtContext)

	// EnterGtltstmt is called when entering the gtltstmt production.
	EnterGtltstmt(c *GtltstmtContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitParenstmt is called when exiting the parenstmt production.
	ExitParenstmt(c *ParenstmtContext)

	// ExitBracestmt is called when exiting the bracestmt production.
	ExitBracestmt(c *BracestmtContext)

	// ExitBracketstmt is called when exiting the bracketstmt production.
	ExitBracketstmt(c *BracketstmtContext)

	// ExitGtltstmt is called when exiting the gtltstmt production.
	ExitGtltstmt(c *GtltstmtContext)
}
