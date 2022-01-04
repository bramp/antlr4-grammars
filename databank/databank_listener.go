// Code generated from databank.g4 by ANTLR 4.9.3. DO NOT EDIT.

package databank // databank
import "github.com/antlr/antlr4/runtime/Go/antlr"

// databankListener is a complete listener for a parse tree produced by databankParser.
type databankListener interface {
	antlr.ParseTreeListener

	// EnterDatabank is called when entering the databank production.
	EnterDatabank(c *DatabankContext)

	// EnterDatedseries is called when entering the datedseries production.
	EnterDatedseries(c *DatedseriesContext)

	// EnterUndatedseries is called when entering the undatedseries production.
	EnterUndatedseries(c *UndatedseriesContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterDateline is called when entering the dateline production.
	EnterDateline(c *DatelineContext)

	// EnterSample is called when entering the sample production.
	EnterSample(c *SampleContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitDatabank is called when exiting the databank production.
	ExitDatabank(c *DatabankContext)

	// ExitDatedseries is called when exiting the datedseries production.
	ExitDatedseries(c *DatedseriesContext)

	// ExitUndatedseries is called when exiting the undatedseries production.
	ExitUndatedseries(c *UndatedseriesContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitDateline is called when exiting the dateline production.
	ExitDateline(c *DatelineContext)

	// ExitSample is called when exiting the sample production.
	ExitSample(c *SampleContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
