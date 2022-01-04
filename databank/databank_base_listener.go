// Code generated from databank.g4 by ANTLR 4.9.3. DO NOT EDIT.

package databank // databank
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedatabankListener is a complete listener for a parse tree produced by databankParser.
type BasedatabankListener struct{}

var _ databankListener = &BasedatabankListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedatabankListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedatabankListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedatabankListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedatabankListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDatabank is called when production databank is entered.
func (s *BasedatabankListener) EnterDatabank(ctx *DatabankContext) {}

// ExitDatabank is called when production databank is exited.
func (s *BasedatabankListener) ExitDatabank(ctx *DatabankContext) {}

// EnterDatedseries is called when production datedseries is entered.
func (s *BasedatabankListener) EnterDatedseries(ctx *DatedseriesContext) {}

// ExitDatedseries is called when production datedseries is exited.
func (s *BasedatabankListener) ExitDatedseries(ctx *DatedseriesContext) {}

// EnterUndatedseries is called when production undatedseries is entered.
func (s *BasedatabankListener) EnterUndatedseries(ctx *UndatedseriesContext) {}

// ExitUndatedseries is called when production undatedseries is exited.
func (s *BasedatabankListener) ExitUndatedseries(ctx *UndatedseriesContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BasedatabankListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BasedatabankListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterDateline is called when production dateline is entered.
func (s *BasedatabankListener) EnterDateline(ctx *DatelineContext) {}

// ExitDateline is called when production dateline is exited.
func (s *BasedatabankListener) ExitDateline(ctx *DatelineContext) {}

// EnterSample is called when production sample is entered.
func (s *BasedatabankListener) EnterSample(ctx *SampleContext) {}

// ExitSample is called when production sample is exited.
func (s *BasedatabankListener) ExitSample(ctx *SampleContext) {}

// EnterNumber is called when production number is entered.
func (s *BasedatabankListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasedatabankListener) ExitNumber(ctx *NumberContext) {}
