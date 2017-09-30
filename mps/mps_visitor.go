// Generated from mps.g4 by ANTLR 4.7.

package mps // mps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by mpsParser.
type mpsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mpsParser#modell.
	VisitModell(ctx *ModellContext) interface{}

	// Visit a parse tree produced by mpsParser#firstrow.
	VisitFirstrow(ctx *FirstrowContext) interface{}

	// Visit a parse tree produced by mpsParser#rows.
	VisitRows(ctx *RowsContext) interface{}

	// Visit a parse tree produced by mpsParser#columns.
	VisitColumns(ctx *ColumnsContext) interface{}

	// Visit a parse tree produced by mpsParser#rhs.
	VisitRhs(ctx *RhsContext) interface{}

	// Visit a parse tree produced by mpsParser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by mpsParser#bounds.
	VisitBounds(ctx *BoundsContext) interface{}

	// Visit a parse tree produced by mpsParser#endata.
	VisitEndata(ctx *EndataContext) interface{}

	// Visit a parse tree produced by mpsParser#rowdatacard.
	VisitRowdatacard(ctx *RowdatacardContext) interface{}

	// Visit a parse tree produced by mpsParser#columndatacards.
	VisitColumndatacards(ctx *ColumndatacardsContext) interface{}

	// Visit a parse tree produced by mpsParser#rhsdatacards.
	VisitRhsdatacards(ctx *RhsdatacardsContext) interface{}

	// Visit a parse tree produced by mpsParser#rangesdatacards.
	VisitRangesdatacards(ctx *RangesdatacardsContext) interface{}

	// Visit a parse tree produced by mpsParser#boundsdatacards.
	VisitBoundsdatacards(ctx *BoundsdatacardsContext) interface{}

	// Visit a parse tree produced by mpsParser#columndatacard.
	VisitColumndatacard(ctx *ColumndatacardContext) interface{}

	// Visit a parse tree produced by mpsParser#rhsdatacard.
	VisitRhsdatacard(ctx *RhsdatacardContext) interface{}

	// Visit a parse tree produced by mpsParser#rangesdatacard.
	VisitRangesdatacard(ctx *RangesdatacardContext) interface{}

	// Visit a parse tree produced by mpsParser#boundsdatacard.
	VisitBoundsdatacard(ctx *BoundsdatacardContext) interface{}

	// Visit a parse tree produced by mpsParser#intblock.
	VisitIntblock(ctx *IntblockContext) interface{}

	// Visit a parse tree produced by mpsParser#startmarker.
	VisitStartmarker(ctx *StartmarkerContext) interface{}

	// Visit a parse tree produced by mpsParser#endmarker.
	VisitEndmarker(ctx *EndmarkerContext) interface{}
}
