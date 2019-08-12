// Code generated from DGSParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dgs // DGSParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DGSParserListener is a complete listener for a parse tree produced by DGSParser.
type DGSParserListener interface {
	antlr.ParseTreeListener

	// EnterDgs is called when entering the dgs production.
	EnterDgs(c *DgsContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterAn is called when entering the an production.
	EnterAn(c *AnContext)

	// EnterCn is called when entering the cn production.
	EnterCn(c *CnContext)

	// EnterDn is called when entering the dn production.
	EnterDn(c *DnContext)

	// EnterAe is called when entering the ae production.
	EnterAe(c *AeContext)

	// EnterCe is called when entering the ce production.
	EnterCe(c *CeContext)

	// EnterDe is called when entering the de production.
	EnterDe(c *DeContext)

	// EnterCg is called when entering the cg production.
	EnterCg(c *CgContext)

	// EnterSt is called when entering the st production.
	EnterSt(c *StContext)

	// EnterCl is called when entering the cl production.
	EnterCl(c *ClContext)

	// EnterAttributes is called when entering the attributes production.
	EnterAttributes(c *AttributesContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterA_map is called when entering the a_map production.
	EnterA_map(c *A_mapContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitDgs is called when exiting the dgs production.
	ExitDgs(c *DgsContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitAn is called when exiting the an production.
	ExitAn(c *AnContext)

	// ExitCn is called when exiting the cn production.
	ExitCn(c *CnContext)

	// ExitDn is called when exiting the dn production.
	ExitDn(c *DnContext)

	// ExitAe is called when exiting the ae production.
	ExitAe(c *AeContext)

	// ExitCe is called when exiting the ce production.
	ExitCe(c *CeContext)

	// ExitDe is called when exiting the de production.
	ExitDe(c *DeContext)

	// ExitCg is called when exiting the cg production.
	ExitCg(c *CgContext)

	// ExitSt is called when exiting the st production.
	ExitSt(c *StContext)

	// ExitCl is called when exiting the cl production.
	ExitCl(c *ClContext)

	// ExitAttributes is called when exiting the attributes production.
	ExitAttributes(c *AttributesContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitA_map is called when exiting the a_map production.
	ExitA_map(c *A_mapContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
