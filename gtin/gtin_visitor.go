// Generated from gtin.g4 by ANTLR 4.7.

package gtin // gtin
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by gtinParser.
type gtinVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by gtinParser#gtin.
	VisitGtin(ctx *GtinContext) interface{}

	// Visit a parse tree produced by gtinParser#gtin8.
	VisitGtin8(ctx *Gtin8Context) interface{}

	// Visit a parse tree produced by gtinParser#ean8.
	VisitEan8(ctx *Ean8Context) interface{}

	// Visit a parse tree produced by gtinParser#gtin12.
	VisitGtin12(ctx *Gtin12Context) interface{}

	// Visit a parse tree produced by gtinParser#gtin13.
	VisitGtin13(ctx *Gtin13Context) interface{}

	// Visit a parse tree produced by gtinParser#gtin14.
	VisitGtin14(ctx *Gtin14Context) interface{}

	// Visit a parse tree produced by gtinParser#upc.
	VisitUpc(ctx *UpcContext) interface{}

	// Visit a parse tree produced by gtinParser#upc_a.
	VisitUpc_a(ctx *Upc_aContext) interface{}

	// Visit a parse tree produced by gtinParser#upc_a_manufacturer.
	VisitUpc_a_manufacturer(ctx *Upc_a_manufacturerContext) interface{}

	// Visit a parse tree produced by gtinParser#upc_a_product.
	VisitUpc_a_product(ctx *Upc_a_productContext) interface{}

	// Visit a parse tree produced by gtinParser#upc_a_5.
	VisitUpc_a_5(ctx *Upc_a_5Context) interface{}

	// Visit a parse tree produced by gtinParser#upc_e.
	VisitUpc_e(ctx *Upc_eContext) interface{}

	// Visit a parse tree produced by gtinParser#num_system.
	VisitNum_system(ctx *Num_systemContext) interface{}

	// Visit a parse tree produced by gtinParser#check_code.
	VisitCheck_code(ctx *Check_codeContext) interface{}

	// Visit a parse tree produced by gtinParser#supplemental_code.
	VisitSupplemental_code(ctx *Supplemental_codeContext) interface{}

	// Visit a parse tree produced by gtinParser#supplemental_code_5.
	VisitSupplemental_code_5(ctx *Supplemental_code_5Context) interface{}

	// Visit a parse tree produced by gtinParser#supplemental_code_2.
	VisitSupplemental_code_2(ctx *Supplemental_code_2Context) interface{}

	// Visit a parse tree produced by gtinParser#ean13.
	VisitEan13(ctx *Ean13Context) interface{}

	// Visit a parse tree produced by gtinParser#ean13_generic.
	VisitEan13_generic(ctx *Ean13_genericContext) interface{}

	// Visit a parse tree produced by gtinParser#ean13_ismn.
	VisitEan13_ismn(ctx *Ean13_ismnContext) interface{}

	// Visit a parse tree produced by gtinParser#gs1_prefix_ismn.
	VisitGs1_prefix_ismn(ctx *Gs1_prefix_ismnContext) interface{}

	// Visit a parse tree produced by gtinParser#ismn_publisher_number.
	VisitIsmn_publisher_number(ctx *Ismn_publisher_numberContext) interface{}

	// Visit a parse tree produced by gtinParser#ismn_item_number.
	VisitIsmn_item_number(ctx *Ismn_item_numberContext) interface{}

	// Visit a parse tree produced by gtinParser#ean13_bookland.
	VisitEan13_bookland(ctx *Ean13_booklandContext) interface{}

	// Visit a parse tree produced by gtinParser#bookland_isbn.
	VisitBookland_isbn(ctx *Bookland_isbnContext) interface{}

	// Visit a parse tree produced by gtinParser#gs1_prefix_bookland_1.
	VisitGs1_prefix_bookland_1(ctx *Gs1_prefix_bookland_1Context) interface{}

	// Visit a parse tree produced by gtinParser#gs1_prefix_bookland_2.
	VisitGs1_prefix_bookland_2(ctx *Gs1_prefix_bookland_2Context) interface{}

	// Visit a parse tree produced by gtinParser#gs1_prefix_issn.
	VisitGs1_prefix_issn(ctx *Gs1_prefix_issnContext) interface{}

	// Visit a parse tree produced by gtinParser#ean13_issn.
	VisitEan13_issn(ctx *Ean13_issnContext) interface{}

	// Visit a parse tree produced by gtinParser#issn.
	VisitIssn(ctx *IssnContext) interface{}

	// Visit a parse tree produced by gtinParser#ean_13_manprod.
	VisitEan_13_manprod(ctx *Ean_13_manprodContext) interface{}

	// Visit a parse tree produced by gtinParser#gs1_prefix.
	VisitGs1_prefix(ctx *Gs1_prefixContext) interface{}

	// Visit a parse tree produced by gtinParser#ean14.
	VisitEan14(ctx *Ean14Context) interface{}

	// Visit a parse tree produced by gtinParser#ean14_appid.
	VisitEan14_appid(ctx *Ean14_appidContext) interface{}

	// Visit a parse tree produced by gtinParser#ean14_packaging.
	VisitEan14_packaging(ctx *Ean14_packagingContext) interface{}

	// Visit a parse tree produced by gtinParser#ean14_product.
	VisitEan14_product(ctx *Ean14_productContext) interface{}

	// Visit a parse tree produced by gtinParser#any_digit.
	VisitAny_digit(ctx *Any_digitContext) interface{}
}
