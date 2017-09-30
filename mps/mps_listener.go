// Generated from mps.g4 by ANTLR 4.7.

package mps // mps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// mpsListener is a complete listener for a parse tree produced by mpsParser.
type mpsListener interface {
	antlr.ParseTreeListener

	// EnterModell is called when entering the modell production.
	EnterModell(c *ModellContext)

	// EnterFirstrow is called when entering the firstrow production.
	EnterFirstrow(c *FirstrowContext)

	// EnterRows is called when entering the rows production.
	EnterRows(c *RowsContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

	// EnterRhs is called when entering the rhs production.
	EnterRhs(c *RhsContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

	// EnterBounds is called when entering the bounds production.
	EnterBounds(c *BoundsContext)

	// EnterEndata is called when entering the endata production.
	EnterEndata(c *EndataContext)

	// EnterRowdatacard is called when entering the rowdatacard production.
	EnterRowdatacard(c *RowdatacardContext)

	// EnterColumndatacards is called when entering the columndatacards production.
	EnterColumndatacards(c *ColumndatacardsContext)

	// EnterRhsdatacards is called when entering the rhsdatacards production.
	EnterRhsdatacards(c *RhsdatacardsContext)

	// EnterRangesdatacards is called when entering the rangesdatacards production.
	EnterRangesdatacards(c *RangesdatacardsContext)

	// EnterBoundsdatacards is called when entering the boundsdatacards production.
	EnterBoundsdatacards(c *BoundsdatacardsContext)

	// EnterColumndatacard is called when entering the columndatacard production.
	EnterColumndatacard(c *ColumndatacardContext)

	// EnterRhsdatacard is called when entering the rhsdatacard production.
	EnterRhsdatacard(c *RhsdatacardContext)

	// EnterRangesdatacard is called when entering the rangesdatacard production.
	EnterRangesdatacard(c *RangesdatacardContext)

	// EnterBoundsdatacard is called when entering the boundsdatacard production.
	EnterBoundsdatacard(c *BoundsdatacardContext)

	// EnterIntblock is called when entering the intblock production.
	EnterIntblock(c *IntblockContext)

	// EnterStartmarker is called when entering the startmarker production.
	EnterStartmarker(c *StartmarkerContext)

	// EnterEndmarker is called when entering the endmarker production.
	EnterEndmarker(c *EndmarkerContext)

	// ExitModell is called when exiting the modell production.
	ExitModell(c *ModellContext)

	// ExitFirstrow is called when exiting the firstrow production.
	ExitFirstrow(c *FirstrowContext)

	// ExitRows is called when exiting the rows production.
	ExitRows(c *RowsContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)

	// ExitRhs is called when exiting the rhs production.
	ExitRhs(c *RhsContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)

	// ExitBounds is called when exiting the bounds production.
	ExitBounds(c *BoundsContext)

	// ExitEndata is called when exiting the endata production.
	ExitEndata(c *EndataContext)

	// ExitRowdatacard is called when exiting the rowdatacard production.
	ExitRowdatacard(c *RowdatacardContext)

	// ExitColumndatacards is called when exiting the columndatacards production.
	ExitColumndatacards(c *ColumndatacardsContext)

	// ExitRhsdatacards is called when exiting the rhsdatacards production.
	ExitRhsdatacards(c *RhsdatacardsContext)

	// ExitRangesdatacards is called when exiting the rangesdatacards production.
	ExitRangesdatacards(c *RangesdatacardsContext)

	// ExitBoundsdatacards is called when exiting the boundsdatacards production.
	ExitBoundsdatacards(c *BoundsdatacardsContext)

	// ExitColumndatacard is called when exiting the columndatacard production.
	ExitColumndatacard(c *ColumndatacardContext)

	// ExitRhsdatacard is called when exiting the rhsdatacard production.
	ExitRhsdatacard(c *RhsdatacardContext)

	// ExitRangesdatacard is called when exiting the rangesdatacard production.
	ExitRangesdatacard(c *RangesdatacardContext)

	// ExitBoundsdatacard is called when exiting the boundsdatacard production.
	ExitBoundsdatacard(c *BoundsdatacardContext)

	// ExitIntblock is called when exiting the intblock production.
	ExitIntblock(c *IntblockContext)

	// ExitStartmarker is called when exiting the startmarker production.
	ExitStartmarker(c *StartmarkerContext)

	// ExitEndmarker is called when exiting the endmarker production.
	ExitEndmarker(c *EndmarkerContext)
}
