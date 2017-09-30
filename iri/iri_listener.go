// Generated from IRI.g4 by ANTLR 4.7.

package iri // IRI
import "github.com/antlr/antlr4/runtime/Go/antlr"

// IRIListener is a complete listener for a parse tree produced by IRIParser.
type IRIListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterIri is called when entering the iri production.
	EnterIri(c *IriContext)

	// EnterIhier_part is called when entering the ihier_part production.
	EnterIhier_part(c *Ihier_partContext)

	// EnterIri_reference is called when entering the iri_reference production.
	EnterIri_reference(c *Iri_referenceContext)

	// EnterAbsolute_iri is called when entering the absolute_iri production.
	EnterAbsolute_iri(c *Absolute_iriContext)

	// EnterIrelative_ref is called when entering the irelative_ref production.
	EnterIrelative_ref(c *Irelative_refContext)

	// EnterIrelative_part is called when entering the irelative_part production.
	EnterIrelative_part(c *Irelative_partContext)

	// EnterIauthority is called when entering the iauthority production.
	EnterIauthority(c *IauthorityContext)

	// EnterIuserinfo is called when entering the iuserinfo production.
	EnterIuserinfo(c *IuserinfoContext)

	// EnterIhost is called when entering the ihost production.
	EnterIhost(c *IhostContext)

	// EnterIreg_name is called when entering the ireg_name production.
	EnterIreg_name(c *Ireg_nameContext)

	// EnterIpath is called when entering the ipath production.
	EnterIpath(c *IpathContext)

	// EnterIpath_abempty is called when entering the ipath_abempty production.
	EnterIpath_abempty(c *Ipath_abemptyContext)

	// EnterIpath_absolute is called when entering the ipath_absolute production.
	EnterIpath_absolute(c *Ipath_absoluteContext)

	// EnterIpath_noscheme is called when entering the ipath_noscheme production.
	EnterIpath_noscheme(c *Ipath_noschemeContext)

	// EnterIpath_rootless is called when entering the ipath_rootless production.
	EnterIpath_rootless(c *Ipath_rootlessContext)

	// EnterIpath_empty is called when entering the ipath_empty production.
	EnterIpath_empty(c *Ipath_emptyContext)

	// EnterIsegment is called when entering the isegment production.
	EnterIsegment(c *IsegmentContext)

	// EnterIsegment_nz is called when entering the isegment_nz production.
	EnterIsegment_nz(c *Isegment_nzContext)

	// EnterIsegment_nz_nc is called when entering the isegment_nz_nc production.
	EnterIsegment_nz_nc(c *Isegment_nz_ncContext)

	// EnterIpchar is called when entering the ipchar production.
	EnterIpchar(c *IpcharContext)

	// EnterIquery is called when entering the iquery production.
	EnterIquery(c *IqueryContext)

	// EnterIfragment is called when entering the ifragment production.
	EnterIfragment(c *IfragmentContext)

	// EnterIunreserved is called when entering the iunreserved production.
	EnterIunreserved(c *IunreservedContext)

	// EnterScheme is called when entering the scheme production.
	EnterScheme(c *SchemeContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterIp_literal is called when entering the ip_literal production.
	EnterIp_literal(c *Ip_literalContext)

	// EnterIp_v_future is called when entering the ip_v_future production.
	EnterIp_v_future(c *Ip_v_futureContext)

	// EnterIp_v6_address is called when entering the ip_v6_address production.
	EnterIp_v6_address(c *Ip_v6_addressContext)

	// EnterH16 is called when entering the h16 production.
	EnterH16(c *H16Context)

	// EnterLs32 is called when entering the ls32 production.
	EnterLs32(c *Ls32Context)

	// EnterIp_v4_address is called when entering the ip_v4_address production.
	EnterIp_v4_address(c *Ip_v4_addressContext)

	// EnterDec_octet is called when entering the dec_octet production.
	EnterDec_octet(c *Dec_octetContext)

	// EnterPct_encoded is called when entering the pct_encoded production.
	EnterPct_encoded(c *Pct_encodedContext)

	// EnterUnreserved is called when entering the unreserved production.
	EnterUnreserved(c *UnreservedContext)

	// EnterReserved is called when entering the reserved production.
	EnterReserved(c *ReservedContext)

	// EnterGen_delims is called when entering the gen_delims production.
	EnterGen_delims(c *Gen_delimsContext)

	// EnterSub_delims is called when entering the sub_delims production.
	EnterSub_delims(c *Sub_delimsContext)

	// EnterAlpha is called when entering the alpha production.
	EnterAlpha(c *AlphaContext)

	// EnterHexdig is called when entering the hexdig production.
	EnterHexdig(c *HexdigContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterNon_zero_digit is called when entering the non_zero_digit production.
	EnterNon_zero_digit(c *Non_zero_digitContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitIri is called when exiting the iri production.
	ExitIri(c *IriContext)

	// ExitIhier_part is called when exiting the ihier_part production.
	ExitIhier_part(c *Ihier_partContext)

	// ExitIri_reference is called when exiting the iri_reference production.
	ExitIri_reference(c *Iri_referenceContext)

	// ExitAbsolute_iri is called when exiting the absolute_iri production.
	ExitAbsolute_iri(c *Absolute_iriContext)

	// ExitIrelative_ref is called when exiting the irelative_ref production.
	ExitIrelative_ref(c *Irelative_refContext)

	// ExitIrelative_part is called when exiting the irelative_part production.
	ExitIrelative_part(c *Irelative_partContext)

	// ExitIauthority is called when exiting the iauthority production.
	ExitIauthority(c *IauthorityContext)

	// ExitIuserinfo is called when exiting the iuserinfo production.
	ExitIuserinfo(c *IuserinfoContext)

	// ExitIhost is called when exiting the ihost production.
	ExitIhost(c *IhostContext)

	// ExitIreg_name is called when exiting the ireg_name production.
	ExitIreg_name(c *Ireg_nameContext)

	// ExitIpath is called when exiting the ipath production.
	ExitIpath(c *IpathContext)

	// ExitIpath_abempty is called when exiting the ipath_abempty production.
	ExitIpath_abempty(c *Ipath_abemptyContext)

	// ExitIpath_absolute is called when exiting the ipath_absolute production.
	ExitIpath_absolute(c *Ipath_absoluteContext)

	// ExitIpath_noscheme is called when exiting the ipath_noscheme production.
	ExitIpath_noscheme(c *Ipath_noschemeContext)

	// ExitIpath_rootless is called when exiting the ipath_rootless production.
	ExitIpath_rootless(c *Ipath_rootlessContext)

	// ExitIpath_empty is called when exiting the ipath_empty production.
	ExitIpath_empty(c *Ipath_emptyContext)

	// ExitIsegment is called when exiting the isegment production.
	ExitIsegment(c *IsegmentContext)

	// ExitIsegment_nz is called when exiting the isegment_nz production.
	ExitIsegment_nz(c *Isegment_nzContext)

	// ExitIsegment_nz_nc is called when exiting the isegment_nz_nc production.
	ExitIsegment_nz_nc(c *Isegment_nz_ncContext)

	// ExitIpchar is called when exiting the ipchar production.
	ExitIpchar(c *IpcharContext)

	// ExitIquery is called when exiting the iquery production.
	ExitIquery(c *IqueryContext)

	// ExitIfragment is called when exiting the ifragment production.
	ExitIfragment(c *IfragmentContext)

	// ExitIunreserved is called when exiting the iunreserved production.
	ExitIunreserved(c *IunreservedContext)

	// ExitScheme is called when exiting the scheme production.
	ExitScheme(c *SchemeContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitIp_literal is called when exiting the ip_literal production.
	ExitIp_literal(c *Ip_literalContext)

	// ExitIp_v_future is called when exiting the ip_v_future production.
	ExitIp_v_future(c *Ip_v_futureContext)

	// ExitIp_v6_address is called when exiting the ip_v6_address production.
	ExitIp_v6_address(c *Ip_v6_addressContext)

	// ExitH16 is called when exiting the h16 production.
	ExitH16(c *H16Context)

	// ExitLs32 is called when exiting the ls32 production.
	ExitLs32(c *Ls32Context)

	// ExitIp_v4_address is called when exiting the ip_v4_address production.
	ExitIp_v4_address(c *Ip_v4_addressContext)

	// ExitDec_octet is called when exiting the dec_octet production.
	ExitDec_octet(c *Dec_octetContext)

	// ExitPct_encoded is called when exiting the pct_encoded production.
	ExitPct_encoded(c *Pct_encodedContext)

	// ExitUnreserved is called when exiting the unreserved production.
	ExitUnreserved(c *UnreservedContext)

	// ExitReserved is called when exiting the reserved production.
	ExitReserved(c *ReservedContext)

	// ExitGen_delims is called when exiting the gen_delims production.
	ExitGen_delims(c *Gen_delimsContext)

	// ExitSub_delims is called when exiting the sub_delims production.
	ExitSub_delims(c *Sub_delimsContext)

	// ExitAlpha is called when exiting the alpha production.
	ExitAlpha(c *AlphaContext)

	// ExitHexdig is called when exiting the hexdig production.
	ExitHexdig(c *HexdigContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitNon_zero_digit is called when exiting the non_zero_digit production.
	ExitNon_zero_digit(c *Non_zero_digitContext)
}
