// Code generated from IRI.g4 by ANTLR 4.9.3. DO NOT EDIT.

package iri // IRI
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIRIListener is a complete listener for a parse tree produced by IRIParser.
type BaseIRIListener struct{}

var _ IRIListener = &BaseIRIListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIRIListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIRIListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIRIListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIRIListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseIRIListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseIRIListener) ExitParse(ctx *ParseContext) {}

// EnterIri is called when production iri is entered.
func (s *BaseIRIListener) EnterIri(ctx *IriContext) {}

// ExitIri is called when production iri is exited.
func (s *BaseIRIListener) ExitIri(ctx *IriContext) {}

// EnterIhier_part is called when production ihier_part is entered.
func (s *BaseIRIListener) EnterIhier_part(ctx *Ihier_partContext) {}

// ExitIhier_part is called when production ihier_part is exited.
func (s *BaseIRIListener) ExitIhier_part(ctx *Ihier_partContext) {}

// EnterIri_reference is called when production iri_reference is entered.
func (s *BaseIRIListener) EnterIri_reference(ctx *Iri_referenceContext) {}

// ExitIri_reference is called when production iri_reference is exited.
func (s *BaseIRIListener) ExitIri_reference(ctx *Iri_referenceContext) {}

// EnterAbsolute_iri is called when production absolute_iri is entered.
func (s *BaseIRIListener) EnterAbsolute_iri(ctx *Absolute_iriContext) {}

// ExitAbsolute_iri is called when production absolute_iri is exited.
func (s *BaseIRIListener) ExitAbsolute_iri(ctx *Absolute_iriContext) {}

// EnterIrelative_ref is called when production irelative_ref is entered.
func (s *BaseIRIListener) EnterIrelative_ref(ctx *Irelative_refContext) {}

// ExitIrelative_ref is called when production irelative_ref is exited.
func (s *BaseIRIListener) ExitIrelative_ref(ctx *Irelative_refContext) {}

// EnterIrelative_part is called when production irelative_part is entered.
func (s *BaseIRIListener) EnterIrelative_part(ctx *Irelative_partContext) {}

// ExitIrelative_part is called when production irelative_part is exited.
func (s *BaseIRIListener) ExitIrelative_part(ctx *Irelative_partContext) {}

// EnterIauthority is called when production iauthority is entered.
func (s *BaseIRIListener) EnterIauthority(ctx *IauthorityContext) {}

// ExitIauthority is called when production iauthority is exited.
func (s *BaseIRIListener) ExitIauthority(ctx *IauthorityContext) {}

// EnterIuserinfo is called when production iuserinfo is entered.
func (s *BaseIRIListener) EnterIuserinfo(ctx *IuserinfoContext) {}

// ExitIuserinfo is called when production iuserinfo is exited.
func (s *BaseIRIListener) ExitIuserinfo(ctx *IuserinfoContext) {}

// EnterIhost is called when production ihost is entered.
func (s *BaseIRIListener) EnterIhost(ctx *IhostContext) {}

// ExitIhost is called when production ihost is exited.
func (s *BaseIRIListener) ExitIhost(ctx *IhostContext) {}

// EnterIreg_name is called when production ireg_name is entered.
func (s *BaseIRIListener) EnterIreg_name(ctx *Ireg_nameContext) {}

// ExitIreg_name is called when production ireg_name is exited.
func (s *BaseIRIListener) ExitIreg_name(ctx *Ireg_nameContext) {}

// EnterIpath is called when production ipath is entered.
func (s *BaseIRIListener) EnterIpath(ctx *IpathContext) {}

// ExitIpath is called when production ipath is exited.
func (s *BaseIRIListener) ExitIpath(ctx *IpathContext) {}

// EnterIpath_abempty is called when production ipath_abempty is entered.
func (s *BaseIRIListener) EnterIpath_abempty(ctx *Ipath_abemptyContext) {}

// ExitIpath_abempty is called when production ipath_abempty is exited.
func (s *BaseIRIListener) ExitIpath_abempty(ctx *Ipath_abemptyContext) {}

// EnterIpath_absolute is called when production ipath_absolute is entered.
func (s *BaseIRIListener) EnterIpath_absolute(ctx *Ipath_absoluteContext) {}

// ExitIpath_absolute is called when production ipath_absolute is exited.
func (s *BaseIRIListener) ExitIpath_absolute(ctx *Ipath_absoluteContext) {}

// EnterIpath_noscheme is called when production ipath_noscheme is entered.
func (s *BaseIRIListener) EnterIpath_noscheme(ctx *Ipath_noschemeContext) {}

// ExitIpath_noscheme is called when production ipath_noscheme is exited.
func (s *BaseIRIListener) ExitIpath_noscheme(ctx *Ipath_noschemeContext) {}

// EnterIpath_rootless is called when production ipath_rootless is entered.
func (s *BaseIRIListener) EnterIpath_rootless(ctx *Ipath_rootlessContext) {}

// ExitIpath_rootless is called when production ipath_rootless is exited.
func (s *BaseIRIListener) ExitIpath_rootless(ctx *Ipath_rootlessContext) {}

// EnterIpath_empty is called when production ipath_empty is entered.
func (s *BaseIRIListener) EnterIpath_empty(ctx *Ipath_emptyContext) {}

// ExitIpath_empty is called when production ipath_empty is exited.
func (s *BaseIRIListener) ExitIpath_empty(ctx *Ipath_emptyContext) {}

// EnterIsegment is called when production isegment is entered.
func (s *BaseIRIListener) EnterIsegment(ctx *IsegmentContext) {}

// ExitIsegment is called when production isegment is exited.
func (s *BaseIRIListener) ExitIsegment(ctx *IsegmentContext) {}

// EnterIsegment_nz is called when production isegment_nz is entered.
func (s *BaseIRIListener) EnterIsegment_nz(ctx *Isegment_nzContext) {}

// ExitIsegment_nz is called when production isegment_nz is exited.
func (s *BaseIRIListener) ExitIsegment_nz(ctx *Isegment_nzContext) {}

// EnterIsegment_nz_nc is called when production isegment_nz_nc is entered.
func (s *BaseIRIListener) EnterIsegment_nz_nc(ctx *Isegment_nz_ncContext) {}

// ExitIsegment_nz_nc is called when production isegment_nz_nc is exited.
func (s *BaseIRIListener) ExitIsegment_nz_nc(ctx *Isegment_nz_ncContext) {}

// EnterIpchar is called when production ipchar is entered.
func (s *BaseIRIListener) EnterIpchar(ctx *IpcharContext) {}

// ExitIpchar is called when production ipchar is exited.
func (s *BaseIRIListener) ExitIpchar(ctx *IpcharContext) {}

// EnterIquery is called when production iquery is entered.
func (s *BaseIRIListener) EnterIquery(ctx *IqueryContext) {}

// ExitIquery is called when production iquery is exited.
func (s *BaseIRIListener) ExitIquery(ctx *IqueryContext) {}

// EnterIfragment is called when production ifragment is entered.
func (s *BaseIRIListener) EnterIfragment(ctx *IfragmentContext) {}

// ExitIfragment is called when production ifragment is exited.
func (s *BaseIRIListener) ExitIfragment(ctx *IfragmentContext) {}

// EnterIunreserved is called when production iunreserved is entered.
func (s *BaseIRIListener) EnterIunreserved(ctx *IunreservedContext) {}

// ExitIunreserved is called when production iunreserved is exited.
func (s *BaseIRIListener) ExitIunreserved(ctx *IunreservedContext) {}

// EnterScheme is called when production scheme is entered.
func (s *BaseIRIListener) EnterScheme(ctx *SchemeContext) {}

// ExitScheme is called when production scheme is exited.
func (s *BaseIRIListener) ExitScheme(ctx *SchemeContext) {}

// EnterPort is called when production port is entered.
func (s *BaseIRIListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseIRIListener) ExitPort(ctx *PortContext) {}

// EnterIp_literal is called when production ip_literal is entered.
func (s *BaseIRIListener) EnterIp_literal(ctx *Ip_literalContext) {}

// ExitIp_literal is called when production ip_literal is exited.
func (s *BaseIRIListener) ExitIp_literal(ctx *Ip_literalContext) {}

// EnterIp_v_future is called when production ip_v_future is entered.
func (s *BaseIRIListener) EnterIp_v_future(ctx *Ip_v_futureContext) {}

// ExitIp_v_future is called when production ip_v_future is exited.
func (s *BaseIRIListener) ExitIp_v_future(ctx *Ip_v_futureContext) {}

// EnterIp_v6_address is called when production ip_v6_address is entered.
func (s *BaseIRIListener) EnterIp_v6_address(ctx *Ip_v6_addressContext) {}

// ExitIp_v6_address is called when production ip_v6_address is exited.
func (s *BaseIRIListener) ExitIp_v6_address(ctx *Ip_v6_addressContext) {}

// EnterH16 is called when production h16 is entered.
func (s *BaseIRIListener) EnterH16(ctx *H16Context) {}

// ExitH16 is called when production h16 is exited.
func (s *BaseIRIListener) ExitH16(ctx *H16Context) {}

// EnterLs32 is called when production ls32 is entered.
func (s *BaseIRIListener) EnterLs32(ctx *Ls32Context) {}

// ExitLs32 is called when production ls32 is exited.
func (s *BaseIRIListener) ExitLs32(ctx *Ls32Context) {}

// EnterIp_v4_address is called when production ip_v4_address is entered.
func (s *BaseIRIListener) EnterIp_v4_address(ctx *Ip_v4_addressContext) {}

// ExitIp_v4_address is called when production ip_v4_address is exited.
func (s *BaseIRIListener) ExitIp_v4_address(ctx *Ip_v4_addressContext) {}

// EnterDec_octet is called when production dec_octet is entered.
func (s *BaseIRIListener) EnterDec_octet(ctx *Dec_octetContext) {}

// ExitDec_octet is called when production dec_octet is exited.
func (s *BaseIRIListener) ExitDec_octet(ctx *Dec_octetContext) {}

// EnterPct_encoded is called when production pct_encoded is entered.
func (s *BaseIRIListener) EnterPct_encoded(ctx *Pct_encodedContext) {}

// ExitPct_encoded is called when production pct_encoded is exited.
func (s *BaseIRIListener) ExitPct_encoded(ctx *Pct_encodedContext) {}

// EnterUnreserved is called when production unreserved is entered.
func (s *BaseIRIListener) EnterUnreserved(ctx *UnreservedContext) {}

// ExitUnreserved is called when production unreserved is exited.
func (s *BaseIRIListener) ExitUnreserved(ctx *UnreservedContext) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseIRIListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseIRIListener) ExitReserved(ctx *ReservedContext) {}

// EnterGen_delims is called when production gen_delims is entered.
func (s *BaseIRIListener) EnterGen_delims(ctx *Gen_delimsContext) {}

// ExitGen_delims is called when production gen_delims is exited.
func (s *BaseIRIListener) ExitGen_delims(ctx *Gen_delimsContext) {}

// EnterSub_delims is called when production sub_delims is entered.
func (s *BaseIRIListener) EnterSub_delims(ctx *Sub_delimsContext) {}

// ExitSub_delims is called when production sub_delims is exited.
func (s *BaseIRIListener) ExitSub_delims(ctx *Sub_delimsContext) {}

// EnterAlpha is called when production alpha is entered.
func (s *BaseIRIListener) EnterAlpha(ctx *AlphaContext) {}

// ExitAlpha is called when production alpha is exited.
func (s *BaseIRIListener) ExitAlpha(ctx *AlphaContext) {}

// EnterHexdig is called when production hexdig is entered.
func (s *BaseIRIListener) EnterHexdig(ctx *HexdigContext) {}

// ExitHexdig is called when production hexdig is exited.
func (s *BaseIRIListener) ExitHexdig(ctx *HexdigContext) {}

// EnterDigit is called when production digit is entered.
func (s *BaseIRIListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BaseIRIListener) ExitDigit(ctx *DigitContext) {}

// EnterNon_zero_digit is called when production non_zero_digit is entered.
func (s *BaseIRIListener) EnterNon_zero_digit(ctx *Non_zero_digitContext) {}

// ExitNon_zero_digit is called when production non_zero_digit is exited.
func (s *BaseIRIListener) ExitNon_zero_digit(ctx *Non_zero_digitContext) {}
