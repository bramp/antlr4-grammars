// Generated from IRI.g4 by ANTLR 4.7.

package iri // IRI
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by IRIParser.
type IRIVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by IRIParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by IRIParser#iri.
	VisitIri(ctx *IriContext) interface{}

	// Visit a parse tree produced by IRIParser#ihier_part.
	VisitIhier_part(ctx *Ihier_partContext) interface{}

	// Visit a parse tree produced by IRIParser#iri_reference.
	VisitIri_reference(ctx *Iri_referenceContext) interface{}

	// Visit a parse tree produced by IRIParser#absolute_iri.
	VisitAbsolute_iri(ctx *Absolute_iriContext) interface{}

	// Visit a parse tree produced by IRIParser#irelative_ref.
	VisitIrelative_ref(ctx *Irelative_refContext) interface{}

	// Visit a parse tree produced by IRIParser#irelative_part.
	VisitIrelative_part(ctx *Irelative_partContext) interface{}

	// Visit a parse tree produced by IRIParser#iauthority.
	VisitIauthority(ctx *IauthorityContext) interface{}

	// Visit a parse tree produced by IRIParser#iuserinfo.
	VisitIuserinfo(ctx *IuserinfoContext) interface{}

	// Visit a parse tree produced by IRIParser#ihost.
	VisitIhost(ctx *IhostContext) interface{}

	// Visit a parse tree produced by IRIParser#ireg_name.
	VisitIreg_name(ctx *Ireg_nameContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath.
	VisitIpath(ctx *IpathContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath_abempty.
	VisitIpath_abempty(ctx *Ipath_abemptyContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath_absolute.
	VisitIpath_absolute(ctx *Ipath_absoluteContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath_noscheme.
	VisitIpath_noscheme(ctx *Ipath_noschemeContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath_rootless.
	VisitIpath_rootless(ctx *Ipath_rootlessContext) interface{}

	// Visit a parse tree produced by IRIParser#ipath_empty.
	VisitIpath_empty(ctx *Ipath_emptyContext) interface{}

	// Visit a parse tree produced by IRIParser#isegment.
	VisitIsegment(ctx *IsegmentContext) interface{}

	// Visit a parse tree produced by IRIParser#isegment_nz.
	VisitIsegment_nz(ctx *Isegment_nzContext) interface{}

	// Visit a parse tree produced by IRIParser#isegment_nz_nc.
	VisitIsegment_nz_nc(ctx *Isegment_nz_ncContext) interface{}

	// Visit a parse tree produced by IRIParser#ipchar.
	VisitIpchar(ctx *IpcharContext) interface{}

	// Visit a parse tree produced by IRIParser#iquery.
	VisitIquery(ctx *IqueryContext) interface{}

	// Visit a parse tree produced by IRIParser#ifragment.
	VisitIfragment(ctx *IfragmentContext) interface{}

	// Visit a parse tree produced by IRIParser#iunreserved.
	VisitIunreserved(ctx *IunreservedContext) interface{}

	// Visit a parse tree produced by IRIParser#scheme.
	VisitScheme(ctx *SchemeContext) interface{}

	// Visit a parse tree produced by IRIParser#port.
	VisitPort(ctx *PortContext) interface{}

	// Visit a parse tree produced by IRIParser#ip_literal.
	VisitIp_literal(ctx *Ip_literalContext) interface{}

	// Visit a parse tree produced by IRIParser#ip_v_future.
	VisitIp_v_future(ctx *Ip_v_futureContext) interface{}

	// Visit a parse tree produced by IRIParser#ip_v6_address.
	VisitIp_v6_address(ctx *Ip_v6_addressContext) interface{}

	// Visit a parse tree produced by IRIParser#h16.
	VisitH16(ctx *H16Context) interface{}

	// Visit a parse tree produced by IRIParser#ls32.
	VisitLs32(ctx *Ls32Context) interface{}

	// Visit a parse tree produced by IRIParser#ip_v4_address.
	VisitIp_v4_address(ctx *Ip_v4_addressContext) interface{}

	// Visit a parse tree produced by IRIParser#dec_octet.
	VisitDec_octet(ctx *Dec_octetContext) interface{}

	// Visit a parse tree produced by IRIParser#pct_encoded.
	VisitPct_encoded(ctx *Pct_encodedContext) interface{}

	// Visit a parse tree produced by IRIParser#unreserved.
	VisitUnreserved(ctx *UnreservedContext) interface{}

	// Visit a parse tree produced by IRIParser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by IRIParser#gen_delims.
	VisitGen_delims(ctx *Gen_delimsContext) interface{}

	// Visit a parse tree produced by IRIParser#sub_delims.
	VisitSub_delims(ctx *Sub_delimsContext) interface{}

	// Visit a parse tree produced by IRIParser#alpha.
	VisitAlpha(ctx *AlphaContext) interface{}

	// Visit a parse tree produced by IRIParser#hexdig.
	VisitHexdig(ctx *HexdigContext) interface{}

	// Visit a parse tree produced by IRIParser#digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by IRIParser#non_zero_digit.
	VisitNon_zero_digit(ctx *Non_zero_digitContext) interface{}
}
