// Code generated from dif.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dif // dif
import "github.com/antlr/antlr4/runtime/Go/antlr"

// difListener is a complete listener for a parse tree produced by difParser.
type difListener interface {
	antlr.ParseTreeListener

	// EnterDif is called when entering the dif production.
	EnterDif(c *DifContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterTableheader is called when entering the tableheader production.
	EnterTableheader(c *TableheaderContext)

	// EnterVectorsheader is called when entering the vectorsheader production.
	EnterVectorsheader(c *VectorsheaderContext)

	// EnterTuplesheader is called when entering the tuplesheader production.
	EnterTuplesheader(c *TuplesheaderContext)

	// EnterDataheader is called when entering the dataheader production.
	EnterDataheader(c *DataheaderContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterDatum is called when entering the datum production.
	EnterDatum(c *DatumContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterNumeric is called when entering the numeric production.
	EnterNumeric(c *NumericContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// ExitDif is called when exiting the dif production.
	ExitDif(c *DifContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitTableheader is called when exiting the tableheader production.
	ExitTableheader(c *TableheaderContext)

	// ExitVectorsheader is called when exiting the vectorsheader production.
	ExitVectorsheader(c *VectorsheaderContext)

	// ExitTuplesheader is called when exiting the tuplesheader production.
	ExitTuplesheader(c *TuplesheaderContext)

	// ExitDataheader is called when exiting the dataheader production.
	ExitDataheader(c *DataheaderContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitDatum is called when exiting the datum production.
	ExitDatum(c *DatumContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitNumeric is called when exiting the numeric production.
	ExitNumeric(c *NumericContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)
}
