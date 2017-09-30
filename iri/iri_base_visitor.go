// Generated from IRI.g4 by ANTLR 4.7.

package iri // IRI
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseIRIVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIRIVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIri(ctx *IriContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIhier_part(ctx *Ihier_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIri_reference(ctx *Iri_referenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitAbsolute_iri(ctx *Absolute_iriContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIrelative_ref(ctx *Irelative_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIrelative_part(ctx *Irelative_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIauthority(ctx *IauthorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIuserinfo(ctx *IuserinfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIhost(ctx *IhostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIreg_name(ctx *Ireg_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath(ctx *IpathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath_abempty(ctx *Ipath_abemptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath_absolute(ctx *Ipath_absoluteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath_noscheme(ctx *Ipath_noschemeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath_rootless(ctx *Ipath_rootlessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpath_empty(ctx *Ipath_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIsegment(ctx *IsegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIsegment_nz(ctx *Isegment_nzContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIsegment_nz_nc(ctx *Isegment_nz_ncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIpchar(ctx *IpcharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIquery(ctx *IqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIfragment(ctx *IfragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIunreserved(ctx *IunreservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitScheme(ctx *SchemeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitPort(ctx *PortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIp_literal(ctx *Ip_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIp_v_future(ctx *Ip_v_futureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIp_v6_address(ctx *Ip_v6_addressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitH16(ctx *H16Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitLs32(ctx *Ls32Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitIp_v4_address(ctx *Ip_v4_addressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitDec_octet(ctx *Dec_octetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitPct_encoded(ctx *Pct_encodedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitUnreserved(ctx *UnreservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitReserved(ctx *ReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitGen_delims(ctx *Gen_delimsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitSub_delims(ctx *Sub_delimsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitAlpha(ctx *AlphaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitHexdig(ctx *HexdigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIRIVisitor) VisitNon_zero_digit(ctx *Non_zero_digitContext) interface{} {
	return v.VisitChildren(ctx)
}
