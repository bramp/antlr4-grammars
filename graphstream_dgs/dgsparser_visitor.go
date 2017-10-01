// Generated from DGSParser.g4 by ANTLR 4.7.

package graphstream_dgs // DGSParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DGSParser.
type DGSParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DGSParser#dgs.
	VisitDgs(ctx *DgsContext) interface{}

	// Visit a parse tree produced by DGSParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by DGSParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by DGSParser#an.
	VisitAn(ctx *AnContext) interface{}

	// Visit a parse tree produced by DGSParser#cn.
	VisitCn(ctx *CnContext) interface{}

	// Visit a parse tree produced by DGSParser#dn.
	VisitDn(ctx *DnContext) interface{}

	// Visit a parse tree produced by DGSParser#ae.
	VisitAe(ctx *AeContext) interface{}

	// Visit a parse tree produced by DGSParser#ce.
	VisitCe(ctx *CeContext) interface{}

	// Visit a parse tree produced by DGSParser#de.
	VisitDe(ctx *DeContext) interface{}

	// Visit a parse tree produced by DGSParser#cg.
	VisitCg(ctx *CgContext) interface{}

	// Visit a parse tree produced by DGSParser#st.
	VisitSt(ctx *StContext) interface{}

	// Visit a parse tree produced by DGSParser#cl.
	VisitCl(ctx *ClContext) interface{}

	// Visit a parse tree produced by DGSParser#attributes.
	VisitAttributes(ctx *AttributesContext) interface{}

	// Visit a parse tree produced by DGSParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by DGSParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by DGSParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by DGSParser#a_map.
	VisitA_map(ctx *A_mapContext) interface{}

	// Visit a parse tree produced by DGSParser#mapping.
	VisitMapping(ctx *MappingContext) interface{}

	// Visit a parse tree produced by DGSParser#direction.
	VisitDirection(ctx *DirectionContext) interface{}

	// Visit a parse tree produced by DGSParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by DGSParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
