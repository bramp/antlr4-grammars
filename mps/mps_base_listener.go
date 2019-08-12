// Code generated from mps.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mps // mps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasempsListener is a complete listener for a parse tree produced by mpsParser.
type BasempsListener struct{}

var _ mpsListener = &BasempsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasempsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasempsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasempsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasempsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModell is called when production modell is entered.
func (s *BasempsListener) EnterModell(ctx *ModellContext) {}

// ExitModell is called when production modell is exited.
func (s *BasempsListener) ExitModell(ctx *ModellContext) {}

// EnterFirstrow is called when production firstrow is entered.
func (s *BasempsListener) EnterFirstrow(ctx *FirstrowContext) {}

// ExitFirstrow is called when production firstrow is exited.
func (s *BasempsListener) ExitFirstrow(ctx *FirstrowContext) {}

// EnterRows is called when production rows is entered.
func (s *BasempsListener) EnterRows(ctx *RowsContext) {}

// ExitRows is called when production rows is exited.
func (s *BasempsListener) ExitRows(ctx *RowsContext) {}

// EnterColumns is called when production columns is entered.
func (s *BasempsListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BasempsListener) ExitColumns(ctx *ColumnsContext) {}

// EnterRhs is called when production rhs is entered.
func (s *BasempsListener) EnterRhs(ctx *RhsContext) {}

// ExitRhs is called when production rhs is exited.
func (s *BasempsListener) ExitRhs(ctx *RhsContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BasempsListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BasempsListener) ExitRanges(ctx *RangesContext) {}

// EnterBounds is called when production bounds is entered.
func (s *BasempsListener) EnterBounds(ctx *BoundsContext) {}

// ExitBounds is called when production bounds is exited.
func (s *BasempsListener) ExitBounds(ctx *BoundsContext) {}

// EnterEndata is called when production endata is entered.
func (s *BasempsListener) EnterEndata(ctx *EndataContext) {}

// ExitEndata is called when production endata is exited.
func (s *BasempsListener) ExitEndata(ctx *EndataContext) {}

// EnterRowdatacard is called when production rowdatacard is entered.
func (s *BasempsListener) EnterRowdatacard(ctx *RowdatacardContext) {}

// ExitRowdatacard is called when production rowdatacard is exited.
func (s *BasempsListener) ExitRowdatacard(ctx *RowdatacardContext) {}

// EnterColumndatacards is called when production columndatacards is entered.
func (s *BasempsListener) EnterColumndatacards(ctx *ColumndatacardsContext) {}

// ExitColumndatacards is called when production columndatacards is exited.
func (s *BasempsListener) ExitColumndatacards(ctx *ColumndatacardsContext) {}

// EnterRhsdatacards is called when production rhsdatacards is entered.
func (s *BasempsListener) EnterRhsdatacards(ctx *RhsdatacardsContext) {}

// ExitRhsdatacards is called when production rhsdatacards is exited.
func (s *BasempsListener) ExitRhsdatacards(ctx *RhsdatacardsContext) {}

// EnterRangesdatacards is called when production rangesdatacards is entered.
func (s *BasempsListener) EnterRangesdatacards(ctx *RangesdatacardsContext) {}

// ExitRangesdatacards is called when production rangesdatacards is exited.
func (s *BasempsListener) ExitRangesdatacards(ctx *RangesdatacardsContext) {}

// EnterBoundsdatacards is called when production boundsdatacards is entered.
func (s *BasempsListener) EnterBoundsdatacards(ctx *BoundsdatacardsContext) {}

// ExitBoundsdatacards is called when production boundsdatacards is exited.
func (s *BasempsListener) ExitBoundsdatacards(ctx *BoundsdatacardsContext) {}

// EnterColumndatacard is called when production columndatacard is entered.
func (s *BasempsListener) EnterColumndatacard(ctx *ColumndatacardContext) {}

// ExitColumndatacard is called when production columndatacard is exited.
func (s *BasempsListener) ExitColumndatacard(ctx *ColumndatacardContext) {}

// EnterRhsdatacard is called when production rhsdatacard is entered.
func (s *BasempsListener) EnterRhsdatacard(ctx *RhsdatacardContext) {}

// ExitRhsdatacard is called when production rhsdatacard is exited.
func (s *BasempsListener) ExitRhsdatacard(ctx *RhsdatacardContext) {}

// EnterRangesdatacard is called when production rangesdatacard is entered.
func (s *BasempsListener) EnterRangesdatacard(ctx *RangesdatacardContext) {}

// ExitRangesdatacard is called when production rangesdatacard is exited.
func (s *BasempsListener) ExitRangesdatacard(ctx *RangesdatacardContext) {}

// EnterBoundsdatacard is called when production boundsdatacard is entered.
func (s *BasempsListener) EnterBoundsdatacard(ctx *BoundsdatacardContext) {}

// ExitBoundsdatacard is called when production boundsdatacard is exited.
func (s *BasempsListener) ExitBoundsdatacard(ctx *BoundsdatacardContext) {}

// EnterIntblock is called when production intblock is entered.
func (s *BasempsListener) EnterIntblock(ctx *IntblockContext) {}

// ExitIntblock is called when production intblock is exited.
func (s *BasempsListener) ExitIntblock(ctx *IntblockContext) {}

// EnterStartmarker is called when production startmarker is entered.
func (s *BasempsListener) EnterStartmarker(ctx *StartmarkerContext) {}

// ExitStartmarker is called when production startmarker is exited.
func (s *BasempsListener) ExitStartmarker(ctx *StartmarkerContext) {}

// EnterEndmarker is called when production endmarker is entered.
func (s *BasempsListener) EnterEndmarker(ctx *EndmarkerContext) {}

// ExitEndmarker is called when production endmarker is exited.
func (s *BasempsListener) ExitEndmarker(ctx *EndmarkerContext) {}
