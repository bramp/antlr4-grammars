// Code generated from dif.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dif // dif
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedifListener is a complete listener for a parse tree produced by difParser.
type BasedifListener struct{}

var _ difListener = &BasedifListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedifListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedifListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedifListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedifListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDif is called when production dif is entered.
func (s *BasedifListener) EnterDif(ctx *DifContext) {}

// ExitDif is called when production dif is exited.
func (s *BasedifListener) ExitDif(ctx *DifContext) {}

// EnterHeader is called when production header is entered.
func (s *BasedifListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BasedifListener) ExitHeader(ctx *HeaderContext) {}

// EnterTableheader is called when production tableheader is entered.
func (s *BasedifListener) EnterTableheader(ctx *TableheaderContext) {}

// ExitTableheader is called when production tableheader is exited.
func (s *BasedifListener) ExitTableheader(ctx *TableheaderContext) {}

// EnterVectorsheader is called when production vectorsheader is entered.
func (s *BasedifListener) EnterVectorsheader(ctx *VectorsheaderContext) {}

// ExitVectorsheader is called when production vectorsheader is exited.
func (s *BasedifListener) ExitVectorsheader(ctx *VectorsheaderContext) {}

// EnterTuplesheader is called when production tuplesheader is entered.
func (s *BasedifListener) EnterTuplesheader(ctx *TuplesheaderContext) {}

// ExitTuplesheader is called when production tuplesheader is exited.
func (s *BasedifListener) ExitTuplesheader(ctx *TuplesheaderContext) {}

// EnterDataheader is called when production dataheader is entered.
func (s *BasedifListener) EnterDataheader(ctx *DataheaderContext) {}

// ExitDataheader is called when production dataheader is exited.
func (s *BasedifListener) ExitDataheader(ctx *DataheaderContext) {}

// EnterData is called when production data is entered.
func (s *BasedifListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BasedifListener) ExitData(ctx *DataContext) {}

// EnterDatum is called when production datum is entered.
func (s *BasedifListener) EnterDatum(ctx *DatumContext) {}

// ExitDatum is called when production datum is exited.
func (s *BasedifListener) ExitDatum(ctx *DatumContext) {}

// EnterDirective is called when production directive is entered.
func (s *BasedifListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BasedifListener) ExitDirective(ctx *DirectiveContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BasedifListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BasedifListener) ExitString_(ctx *String_Context) {}

// EnterNumeric is called when production numeric is entered.
func (s *BasedifListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BasedifListener) ExitNumeric(ctx *NumericContext) {}

// EnterPair is called when production pair is entered.
func (s *BasedifListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasedifListener) ExitPair(ctx *PairContext) {}
