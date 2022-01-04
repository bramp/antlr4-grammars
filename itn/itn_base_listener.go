// Code generated from itn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package itn // itn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseitnListener is a complete listener for a parse tree produced by itnParser.
type BaseitnListener struct{}

var _ itnListener = &BaseitnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseitnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseitnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseitnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseitnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterItinerary is called when production itinerary is entered.
func (s *BaseitnListener) EnterItinerary(ctx *ItineraryContext) {}

// ExitItinerary is called when production itinerary is exited.
func (s *BaseitnListener) ExitItinerary(ctx *ItineraryContext) {}

// EnterLine is called when production line is entered.
func (s *BaseitnListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseitnListener) ExitLine(ctx *LineContext) {}

// EnterLongitude is called when production longitude is entered.
func (s *BaseitnListener) EnterLongitude(ctx *LongitudeContext) {}

// ExitLongitude is called when production longitude is exited.
func (s *BaseitnListener) ExitLongitude(ctx *LongitudeContext) {}

// EnterLatitude is called when production latitude is entered.
func (s *BaseitnListener) EnterLatitude(ctx *LatitudeContext) {}

// ExitLatitude is called when production latitude is exited.
func (s *BaseitnListener) ExitLatitude(ctx *LatitudeContext) {}

// EnterDescr is called when production descr is entered.
func (s *BaseitnListener) EnterDescr(ctx *DescrContext) {}

// ExitDescr is called when production descr is exited.
func (s *BaseitnListener) ExitDescr(ctx *DescrContext) {}

// EnterFlag is called when production flag is entered.
func (s *BaseitnListener) EnterFlag(ctx *FlagContext) {}

// ExitFlag is called when production flag is exited.
func (s *BaseitnListener) ExitFlag(ctx *FlagContext) {}
