// Generated from gtin.g4 by ANTLR 4.7.

package gtin // gtin
import "github.com/antlr/antlr4/runtime/Go/antlr"

// gtinListener is a complete listener for a parse tree produced by gtinParser.
type gtinListener interface {
	antlr.ParseTreeListener

	// EnterGtin is called when entering the gtin production.
	EnterGtin(c *GtinContext)

	// EnterGtin8 is called when entering the gtin8 production.
	EnterGtin8(c *Gtin8Context)

	// EnterEan8 is called when entering the ean8 production.
	EnterEan8(c *Ean8Context)

	// EnterGtin12 is called when entering the gtin12 production.
	EnterGtin12(c *Gtin12Context)

	// EnterGtin13 is called when entering the gtin13 production.
	EnterGtin13(c *Gtin13Context)

	// EnterGtin14 is called when entering the gtin14 production.
	EnterGtin14(c *Gtin14Context)

	// EnterUpc is called when entering the upc production.
	EnterUpc(c *UpcContext)

	// EnterUpc_a is called when entering the upc_a production.
	EnterUpc_a(c *Upc_aContext)

	// EnterUpc_a_manufacturer is called when entering the upc_a_manufacturer production.
	EnterUpc_a_manufacturer(c *Upc_a_manufacturerContext)

	// EnterUpc_a_product is called when entering the upc_a_product production.
	EnterUpc_a_product(c *Upc_a_productContext)

	// EnterUpc_a_5 is called when entering the upc_a_5 production.
	EnterUpc_a_5(c *Upc_a_5Context)

	// EnterUpc_e is called when entering the upc_e production.
	EnterUpc_e(c *Upc_eContext)

	// EnterNum_system is called when entering the num_system production.
	EnterNum_system(c *Num_systemContext)

	// EnterCheck_code is called when entering the check_code production.
	EnterCheck_code(c *Check_codeContext)

	// EnterSupplemental_code is called when entering the supplemental_code production.
	EnterSupplemental_code(c *Supplemental_codeContext)

	// EnterSupplemental_code_5 is called when entering the supplemental_code_5 production.
	EnterSupplemental_code_5(c *Supplemental_code_5Context)

	// EnterSupplemental_code_2 is called when entering the supplemental_code_2 production.
	EnterSupplemental_code_2(c *Supplemental_code_2Context)

	// EnterEan13 is called when entering the ean13 production.
	EnterEan13(c *Ean13Context)

	// EnterEan13_generic is called when entering the ean13_generic production.
	EnterEan13_generic(c *Ean13_genericContext)

	// EnterEan13_ismn is called when entering the ean13_ismn production.
	EnterEan13_ismn(c *Ean13_ismnContext)

	// EnterGs1_prefix_ismn is called when entering the gs1_prefix_ismn production.
	EnterGs1_prefix_ismn(c *Gs1_prefix_ismnContext)

	// EnterIsmn_publisher_number is called when entering the ismn_publisher_number production.
	EnterIsmn_publisher_number(c *Ismn_publisher_numberContext)

	// EnterIsmn_item_number is called when entering the ismn_item_number production.
	EnterIsmn_item_number(c *Ismn_item_numberContext)

	// EnterEan13_bookland is called when entering the ean13_bookland production.
	EnterEan13_bookland(c *Ean13_booklandContext)

	// EnterBookland_isbn is called when entering the bookland_isbn production.
	EnterBookland_isbn(c *Bookland_isbnContext)

	// EnterGs1_prefix_bookland_1 is called when entering the gs1_prefix_bookland_1 production.
	EnterGs1_prefix_bookland_1(c *Gs1_prefix_bookland_1Context)

	// EnterGs1_prefix_bookland_2 is called when entering the gs1_prefix_bookland_2 production.
	EnterGs1_prefix_bookland_2(c *Gs1_prefix_bookland_2Context)

	// EnterGs1_prefix_issn is called when entering the gs1_prefix_issn production.
	EnterGs1_prefix_issn(c *Gs1_prefix_issnContext)

	// EnterEan13_issn is called when entering the ean13_issn production.
	EnterEan13_issn(c *Ean13_issnContext)

	// EnterIssn is called when entering the issn production.
	EnterIssn(c *IssnContext)

	// EnterEan_13_manprod is called when entering the ean_13_manprod production.
	EnterEan_13_manprod(c *Ean_13_manprodContext)

	// EnterGs1_prefix is called when entering the gs1_prefix production.
	EnterGs1_prefix(c *Gs1_prefixContext)

	// EnterEan14 is called when entering the ean14 production.
	EnterEan14(c *Ean14Context)

	// EnterEan14_appid is called when entering the ean14_appid production.
	EnterEan14_appid(c *Ean14_appidContext)

	// EnterEan14_packaging is called when entering the ean14_packaging production.
	EnterEan14_packaging(c *Ean14_packagingContext)

	// EnterEan14_product is called when entering the ean14_product production.
	EnterEan14_product(c *Ean14_productContext)

	// EnterAny_digit is called when entering the any_digit production.
	EnterAny_digit(c *Any_digitContext)

	// ExitGtin is called when exiting the gtin production.
	ExitGtin(c *GtinContext)

	// ExitGtin8 is called when exiting the gtin8 production.
	ExitGtin8(c *Gtin8Context)

	// ExitEan8 is called when exiting the ean8 production.
	ExitEan8(c *Ean8Context)

	// ExitGtin12 is called when exiting the gtin12 production.
	ExitGtin12(c *Gtin12Context)

	// ExitGtin13 is called when exiting the gtin13 production.
	ExitGtin13(c *Gtin13Context)

	// ExitGtin14 is called when exiting the gtin14 production.
	ExitGtin14(c *Gtin14Context)

	// ExitUpc is called when exiting the upc production.
	ExitUpc(c *UpcContext)

	// ExitUpc_a is called when exiting the upc_a production.
	ExitUpc_a(c *Upc_aContext)

	// ExitUpc_a_manufacturer is called when exiting the upc_a_manufacturer production.
	ExitUpc_a_manufacturer(c *Upc_a_manufacturerContext)

	// ExitUpc_a_product is called when exiting the upc_a_product production.
	ExitUpc_a_product(c *Upc_a_productContext)

	// ExitUpc_a_5 is called when exiting the upc_a_5 production.
	ExitUpc_a_5(c *Upc_a_5Context)

	// ExitUpc_e is called when exiting the upc_e production.
	ExitUpc_e(c *Upc_eContext)

	// ExitNum_system is called when exiting the num_system production.
	ExitNum_system(c *Num_systemContext)

	// ExitCheck_code is called when exiting the check_code production.
	ExitCheck_code(c *Check_codeContext)

	// ExitSupplemental_code is called when exiting the supplemental_code production.
	ExitSupplemental_code(c *Supplemental_codeContext)

	// ExitSupplemental_code_5 is called when exiting the supplemental_code_5 production.
	ExitSupplemental_code_5(c *Supplemental_code_5Context)

	// ExitSupplemental_code_2 is called when exiting the supplemental_code_2 production.
	ExitSupplemental_code_2(c *Supplemental_code_2Context)

	// ExitEan13 is called when exiting the ean13 production.
	ExitEan13(c *Ean13Context)

	// ExitEan13_generic is called when exiting the ean13_generic production.
	ExitEan13_generic(c *Ean13_genericContext)

	// ExitEan13_ismn is called when exiting the ean13_ismn production.
	ExitEan13_ismn(c *Ean13_ismnContext)

	// ExitGs1_prefix_ismn is called when exiting the gs1_prefix_ismn production.
	ExitGs1_prefix_ismn(c *Gs1_prefix_ismnContext)

	// ExitIsmn_publisher_number is called when exiting the ismn_publisher_number production.
	ExitIsmn_publisher_number(c *Ismn_publisher_numberContext)

	// ExitIsmn_item_number is called when exiting the ismn_item_number production.
	ExitIsmn_item_number(c *Ismn_item_numberContext)

	// ExitEan13_bookland is called when exiting the ean13_bookland production.
	ExitEan13_bookland(c *Ean13_booklandContext)

	// ExitBookland_isbn is called when exiting the bookland_isbn production.
	ExitBookland_isbn(c *Bookland_isbnContext)

	// ExitGs1_prefix_bookland_1 is called when exiting the gs1_prefix_bookland_1 production.
	ExitGs1_prefix_bookland_1(c *Gs1_prefix_bookland_1Context)

	// ExitGs1_prefix_bookland_2 is called when exiting the gs1_prefix_bookland_2 production.
	ExitGs1_prefix_bookland_2(c *Gs1_prefix_bookland_2Context)

	// ExitGs1_prefix_issn is called when exiting the gs1_prefix_issn production.
	ExitGs1_prefix_issn(c *Gs1_prefix_issnContext)

	// ExitEan13_issn is called when exiting the ean13_issn production.
	ExitEan13_issn(c *Ean13_issnContext)

	// ExitIssn is called when exiting the issn production.
	ExitIssn(c *IssnContext)

	// ExitEan_13_manprod is called when exiting the ean_13_manprod production.
	ExitEan_13_manprod(c *Ean_13_manprodContext)

	// ExitGs1_prefix is called when exiting the gs1_prefix production.
	ExitGs1_prefix(c *Gs1_prefixContext)

	// ExitEan14 is called when exiting the ean14 production.
	ExitEan14(c *Ean14Context)

	// ExitEan14_appid is called when exiting the ean14_appid production.
	ExitEan14_appid(c *Ean14_appidContext)

	// ExitEan14_packaging is called when exiting the ean14_packaging production.
	ExitEan14_packaging(c *Ean14_packagingContext)

	// ExitEan14_product is called when exiting the ean14_product production.
	ExitEan14_product(c *Ean14_productContext)

	// ExitAny_digit is called when exiting the any_digit production.
	ExitAny_digit(c *Any_digitContext)
}
