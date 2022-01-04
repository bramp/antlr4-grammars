// Code generated from DGSParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgs // DGSParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDGSParserListener is a complete listener for a parse tree produced by DGSParser.
type BaseDGSParserListener struct{}

var _ DGSParserListener = &BaseDGSParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDGSParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDGSParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDGSParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDGSParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDgs is called when production dgs is entered.
func (s *BaseDGSParserListener) EnterDgs(ctx *DgsContext) {}

// ExitDgs is called when production dgs is exited.
func (s *BaseDGSParserListener) ExitDgs(ctx *DgsContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseDGSParserListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseDGSParserListener) ExitHeader(ctx *HeaderContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseDGSParserListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseDGSParserListener) ExitEvent(ctx *EventContext) {}

// EnterAn is called when production an is entered.
func (s *BaseDGSParserListener) EnterAn(ctx *AnContext) {}

// ExitAn is called when production an is exited.
func (s *BaseDGSParserListener) ExitAn(ctx *AnContext) {}

// EnterCn is called when production cn is entered.
func (s *BaseDGSParserListener) EnterCn(ctx *CnContext) {}

// ExitCn is called when production cn is exited.
func (s *BaseDGSParserListener) ExitCn(ctx *CnContext) {}

// EnterDn is called when production dn is entered.
func (s *BaseDGSParserListener) EnterDn(ctx *DnContext) {}

// ExitDn is called when production dn is exited.
func (s *BaseDGSParserListener) ExitDn(ctx *DnContext) {}

// EnterAe is called when production ae is entered.
func (s *BaseDGSParserListener) EnterAe(ctx *AeContext) {}

// ExitAe is called when production ae is exited.
func (s *BaseDGSParserListener) ExitAe(ctx *AeContext) {}

// EnterCe is called when production ce is entered.
func (s *BaseDGSParserListener) EnterCe(ctx *CeContext) {}

// ExitCe is called when production ce is exited.
func (s *BaseDGSParserListener) ExitCe(ctx *CeContext) {}

// EnterDe is called when production de is entered.
func (s *BaseDGSParserListener) EnterDe(ctx *DeContext) {}

// ExitDe is called when production de is exited.
func (s *BaseDGSParserListener) ExitDe(ctx *DeContext) {}

// EnterCg is called when production cg is entered.
func (s *BaseDGSParserListener) EnterCg(ctx *CgContext) {}

// ExitCg is called when production cg is exited.
func (s *BaseDGSParserListener) ExitCg(ctx *CgContext) {}

// EnterSt is called when production st is entered.
func (s *BaseDGSParserListener) EnterSt(ctx *StContext) {}

// ExitSt is called when production st is exited.
func (s *BaseDGSParserListener) ExitSt(ctx *StContext) {}

// EnterCl is called when production cl is entered.
func (s *BaseDGSParserListener) EnterCl(ctx *ClContext) {}

// ExitCl is called when production cl is exited.
func (s *BaseDGSParserListener) ExitCl(ctx *ClContext) {}

// EnterAttributes is called when production attributes is entered.
func (s *BaseDGSParserListener) EnterAttributes(ctx *AttributesContext) {}

// ExitAttributes is called when production attributes is exited.
func (s *BaseDGSParserListener) ExitAttributes(ctx *AttributesContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseDGSParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseDGSParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseDGSParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseDGSParserListener) ExitValue(ctx *ValueContext) {}

// EnterArray_ is called when production array_ is entered.
func (s *BaseDGSParserListener) EnterArray_(ctx *Array_Context) {}

// ExitArray_ is called when production array_ is exited.
func (s *BaseDGSParserListener) ExitArray_(ctx *Array_Context) {}

// EnterA_map is called when production a_map is entered.
func (s *BaseDGSParserListener) EnterA_map(ctx *A_mapContext) {}

// ExitA_map is called when production a_map is exited.
func (s *BaseDGSParserListener) ExitA_map(ctx *A_mapContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BaseDGSParserListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BaseDGSParserListener) ExitMapping(ctx *MappingContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseDGSParserListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseDGSParserListener) ExitDirection(ctx *DirectionContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseDGSParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseDGSParserListener) ExitAssign(ctx *AssignContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseDGSParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseDGSParserListener) ExitIdentifier(ctx *IdentifierContext) {}
