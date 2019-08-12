// Code generated from gtin.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gtin // gtin
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 258,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 82, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 5, 8, 105, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 5, 16, 135, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 150, 10, 19, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 178, 10, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 2, 2, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 2, 5, 3, 2, 4, 12, 3, 2, 3, 11, 3, 2, 3, 12,
	2, 230, 2, 81, 3, 2, 2, 2, 4, 85, 3, 2, 2, 2, 6, 87, 3, 2, 2, 2, 8, 96,
	3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2, 14, 104, 3, 2, 2,
	2, 16, 106, 3, 2, 2, 2, 18, 111, 3, 2, 2, 2, 20, 113, 3, 2, 2, 2, 22, 115,
	3, 2, 2, 2, 24, 121, 3, 2, 2, 2, 26, 128, 3, 2, 2, 2, 28, 130, 3, 2, 2,
	2, 30, 134, 3, 2, 2, 2, 32, 136, 3, 2, 2, 2, 34, 142, 3, 2, 2, 2, 36, 149,
	3, 2, 2, 2, 38, 151, 3, 2, 2, 2, 40, 155, 3, 2, 2, 2, 42, 160, 3, 2, 2,
	2, 44, 165, 3, 2, 2, 2, 46, 170, 3, 2, 2, 2, 48, 177, 3, 2, 2, 2, 50, 181,
	3, 2, 2, 2, 52, 191, 3, 2, 2, 2, 54, 196, 3, 2, 2, 2, 56, 201, 3, 2, 2,
	2, 58, 205, 3, 2, 2, 2, 60, 209, 3, 2, 2, 2, 62, 219, 3, 2, 2, 2, 64, 229,
	3, 2, 2, 2, 66, 233, 3, 2, 2, 2, 68, 237, 3, 2, 2, 2, 70, 240, 3, 2, 2,
	2, 72, 242, 3, 2, 2, 2, 74, 255, 3, 2, 2, 2, 76, 82, 5, 4, 3, 2, 77, 82,
	5, 8, 5, 2, 78, 82, 5, 10, 6, 2, 79, 82, 5, 12, 7, 2, 80, 82, 5, 30, 16,
	2, 81, 76, 3, 2, 2, 2, 81, 77, 3, 2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 7, 2, 2, 3,
	84, 3, 3, 2, 2, 2, 85, 86, 5, 6, 4, 2, 86, 5, 3, 2, 2, 2, 87, 88, 5, 74,
	38, 2, 88, 89, 5, 74, 38, 2, 89, 90, 5, 74, 38, 2, 90, 91, 5, 74, 38, 2,
	91, 92, 5, 74, 38, 2, 92, 93, 5, 74, 38, 2, 93, 94, 5, 74, 38, 2, 94, 95,
	5, 74, 38, 2, 95, 7, 3, 2, 2, 2, 96, 97, 5, 14, 8, 2, 97, 9, 3, 2, 2, 2,
	98, 99, 5, 36, 19, 2, 99, 11, 3, 2, 2, 2, 100, 101, 5, 66, 34, 2, 101,
	13, 3, 2, 2, 2, 102, 105, 5, 16, 9, 2, 103, 105, 5, 24, 13, 2, 104, 102,
	3, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 15, 3, 2, 2, 2, 106, 107, 5, 26,
	14, 2, 107, 108, 5, 18, 10, 2, 108, 109, 5, 20, 11, 2, 109, 110, 5, 28,
	15, 2, 110, 17, 3, 2, 2, 2, 111, 112, 5, 22, 12, 2, 112, 19, 3, 2, 2, 2,
	113, 114, 5, 22, 12, 2, 114, 21, 3, 2, 2, 2, 115, 116, 5, 74, 38, 2, 116,
	117, 5, 74, 38, 2, 117, 118, 5, 74, 38, 2, 118, 119, 5, 74, 38, 2, 119,
	120, 5, 74, 38, 2, 120, 23, 3, 2, 2, 2, 121, 122, 5, 74, 38, 2, 122, 123,
	5, 74, 38, 2, 123, 124, 5, 74, 38, 2, 124, 125, 5, 74, 38, 2, 125, 126,
	5, 74, 38, 2, 126, 127, 5, 74, 38, 2, 127, 25, 3, 2, 2, 2, 128, 129, 5,
	74, 38, 2, 129, 27, 3, 2, 2, 2, 130, 131, 5, 74, 38, 2, 131, 29, 3, 2,
	2, 2, 132, 135, 5, 32, 17, 2, 133, 135, 5, 34, 18, 2, 134, 132, 3, 2, 2,
	2, 134, 133, 3, 2, 2, 2, 135, 31, 3, 2, 2, 2, 136, 137, 5, 74, 38, 2, 137,
	138, 5, 74, 38, 2, 138, 139, 5, 74, 38, 2, 139, 140, 5, 74, 38, 2, 140,
	141, 5, 74, 38, 2, 141, 33, 3, 2, 2, 2, 142, 143, 5, 74, 38, 2, 143, 144,
	5, 74, 38, 2, 144, 35, 3, 2, 2, 2, 145, 150, 5, 40, 21, 2, 146, 150, 5,
	58, 30, 2, 147, 150, 5, 48, 25, 2, 148, 150, 5, 38, 20, 2, 149, 145, 3,
	2, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2,
	2, 150, 37, 3, 2, 2, 2, 151, 152, 5, 64, 33, 2, 152, 153, 5, 62, 32, 2,
	153, 154, 5, 28, 15, 2, 154, 39, 3, 2, 2, 2, 155, 156, 5, 42, 22, 2, 156,
	157, 5, 44, 23, 2, 157, 158, 5, 46, 24, 2, 158, 159, 5, 28, 15, 2, 159,
	41, 3, 2, 2, 2, 160, 161, 7, 12, 2, 2, 161, 162, 7, 10, 2, 2, 162, 163,
	7, 12, 2, 2, 163, 164, 7, 3, 2, 2, 164, 43, 3, 2, 2, 2, 165, 166, 5, 74,
	38, 2, 166, 167, 5, 74, 38, 2, 167, 168, 5, 74, 38, 2, 168, 169, 5, 74,
	38, 2, 169, 45, 3, 2, 2, 2, 170, 171, 5, 74, 38, 2, 171, 172, 5, 74, 38,
	2, 172, 173, 5, 74, 38, 2, 173, 174, 5, 74, 38, 2, 174, 47, 3, 2, 2, 2,
	175, 178, 5, 52, 27, 2, 176, 178, 5, 54, 28, 2, 177, 175, 3, 2, 2, 2, 177,
	176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 5, 50, 26, 2, 180, 49,
	3, 2, 2, 2, 181, 182, 5, 74, 38, 2, 182, 183, 5, 74, 38, 2, 183, 184, 5,
	74, 38, 2, 184, 185, 5, 74, 38, 2, 185, 186, 5, 74, 38, 2, 186, 187, 5,
	74, 38, 2, 187, 188, 5, 74, 38, 2, 188, 189, 5, 74, 38, 2, 189, 190, 5,
	74, 38, 2, 190, 51, 3, 2, 2, 2, 191, 192, 7, 12, 2, 2, 192, 193, 7, 10,
	2, 2, 193, 194, 7, 12, 2, 2, 194, 195, 9, 2, 2, 2, 195, 53, 3, 2, 2, 2,
	196, 197, 7, 12, 2, 2, 197, 198, 7, 10, 2, 2, 198, 199, 7, 11, 2, 2, 199,
	200, 5, 74, 38, 2, 200, 55, 3, 2, 2, 2, 201, 202, 7, 12, 2, 2, 202, 203,
	7, 10, 2, 2, 203, 204, 7, 10, 2, 2, 204, 57, 3, 2, 2, 2, 205, 206, 5, 56,
	29, 2, 206, 207, 5, 60, 31, 2, 207, 208, 5, 28, 15, 2, 208, 59, 3, 2, 2,
	2, 209, 210, 5, 74, 38, 2, 210, 211, 5, 74, 38, 2, 211, 212, 5, 74, 38,
	2, 212, 213, 5, 74, 38, 2, 213, 214, 5, 74, 38, 2, 214, 215, 5, 74, 38,
	2, 215, 216, 5, 74, 38, 2, 216, 217, 5, 74, 38, 2, 217, 218, 5, 74, 38,
	2, 218, 61, 3, 2, 2, 2, 219, 220, 5, 74, 38, 2, 220, 221, 5, 74, 38, 2,
	221, 222, 5, 74, 38, 2, 222, 223, 5, 74, 38, 2, 223, 224, 5, 74, 38, 2,
	224, 225, 5, 74, 38, 2, 225, 226, 5, 74, 38, 2, 226, 227, 5, 74, 38, 2,
	227, 228, 5, 74, 38, 2, 228, 63, 3, 2, 2, 2, 229, 230, 5, 74, 38, 2, 230,
	231, 5, 74, 38, 2, 231, 232, 5, 74, 38, 2, 232, 65, 3, 2, 2, 2, 233, 234,
	5, 70, 36, 2, 234, 235, 5, 72, 37, 2, 235, 236, 5, 28, 15, 2, 236, 67,
	3, 2, 2, 2, 237, 238, 5, 74, 38, 2, 238, 239, 5, 74, 38, 2, 239, 69, 3,
	2, 2, 2, 240, 241, 9, 3, 2, 2, 241, 71, 3, 2, 2, 2, 242, 243, 5, 74, 38,
	2, 243, 244, 5, 74, 38, 2, 244, 245, 5, 74, 38, 2, 245, 246, 5, 74, 38,
	2, 246, 247, 5, 74, 38, 2, 247, 248, 5, 74, 38, 2, 248, 249, 5, 74, 38,
	2, 249, 250, 5, 74, 38, 2, 250, 251, 5, 74, 38, 2, 251, 252, 5, 74, 38,
	2, 252, 253, 5, 74, 38, 2, 253, 254, 5, 74, 38, 2, 254, 73, 3, 2, 2, 2,
	255, 256, 9, 4, 2, 2, 256, 75, 3, 2, 2, 2, 7, 81, 104, 134, 149, 177,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'0'", "'1'", "'2'", "'3'", "'4'", "'5'", "'6'", "'7'", "'8'", "'9'",
	"'-'",
}
var symbolicNames = []string{
	"", "DIGIT_0", "DIGIT_1", "DIGIT_2", "DIGIT_3", "DIGIT_4", "DIGIT_5", "DIGIT_6",
	"DIGIT_7", "DIGIT_8", "DIGIT_9", "HYPHEN", "WS",
}

var ruleNames = []string{
	"gtin", "gtin8", "ean8", "gtin12", "gtin13", "gtin14", "upc", "upc_a",
	"upc_a_manufacturer", "upc_a_product", "upc_a_5", "upc_e", "num_system",
	"check_code", "supplemental_code", "supplemental_code_5", "supplemental_code_2",
	"ean13", "ean13_generic", "ean13_ismn", "gs1_prefix_ismn", "ismn_publisher_number",
	"ismn_item_number", "ean13_bookland", "bookland_isbn", "gs1_prefix_bookland_1",
	"gs1_prefix_bookland_2", "gs1_prefix_issn", "ean13_issn", "issn", "ean_13_manprod",
	"gs1_prefix", "ean14", "ean14_appid", "ean14_packaging", "ean14_product",
	"any_digit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type gtinParser struct {
	*antlr.BaseParser
}

func NewgtinParser(input antlr.TokenStream) *gtinParser {
	this := new(gtinParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gtin.g4"

	return this
}

// gtinParser tokens.
const (
	gtinParserEOF     = antlr.TokenEOF
	gtinParserDIGIT_0 = 1
	gtinParserDIGIT_1 = 2
	gtinParserDIGIT_2 = 3
	gtinParserDIGIT_3 = 4
	gtinParserDIGIT_4 = 5
	gtinParserDIGIT_5 = 6
	gtinParserDIGIT_6 = 7
	gtinParserDIGIT_7 = 8
	gtinParserDIGIT_8 = 9
	gtinParserDIGIT_9 = 10
	gtinParserHYPHEN  = 11
	gtinParserWS      = 12
)

// gtinParser rules.
const (
	gtinParserRULE_gtin                  = 0
	gtinParserRULE_gtin8                 = 1
	gtinParserRULE_ean8                  = 2
	gtinParserRULE_gtin12                = 3
	gtinParserRULE_gtin13                = 4
	gtinParserRULE_gtin14                = 5
	gtinParserRULE_upc                   = 6
	gtinParserRULE_upc_a                 = 7
	gtinParserRULE_upc_a_manufacturer    = 8
	gtinParserRULE_upc_a_product         = 9
	gtinParserRULE_upc_a_5               = 10
	gtinParserRULE_upc_e                 = 11
	gtinParserRULE_num_system            = 12
	gtinParserRULE_check_code            = 13
	gtinParserRULE_supplemental_code     = 14
	gtinParserRULE_supplemental_code_5   = 15
	gtinParserRULE_supplemental_code_2   = 16
	gtinParserRULE_ean13                 = 17
	gtinParserRULE_ean13_generic         = 18
	gtinParserRULE_ean13_ismn            = 19
	gtinParserRULE_gs1_prefix_ismn       = 20
	gtinParserRULE_ismn_publisher_number = 21
	gtinParserRULE_ismn_item_number      = 22
	gtinParserRULE_ean13_bookland        = 23
	gtinParserRULE_bookland_isbn         = 24
	gtinParserRULE_gs1_prefix_bookland_1 = 25
	gtinParserRULE_gs1_prefix_bookland_2 = 26
	gtinParserRULE_gs1_prefix_issn       = 27
	gtinParserRULE_ean13_issn            = 28
	gtinParserRULE_issn                  = 29
	gtinParserRULE_ean_13_manprod        = 30
	gtinParserRULE_gs1_prefix            = 31
	gtinParserRULE_ean14                 = 32
	gtinParserRULE_ean14_appid           = 33
	gtinParserRULE_ean14_packaging       = 34
	gtinParserRULE_ean14_product         = 35
	gtinParserRULE_any_digit             = 36
)

// IGtinContext is an interface to support dynamic dispatch.
type IGtinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtinContext differentiates from other interfaces.
	IsGtinContext()
}

type GtinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtinContext() *GtinContext {
	var p = new(GtinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gtin
	return p
}

func (*GtinContext) IsGtinContext() {}

func NewGtinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GtinContext {
	var p = new(GtinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gtin

	return p
}

func (s *GtinContext) GetParser() antlr.Parser { return s.parser }

func (s *GtinContext) EOF() antlr.TerminalNode {
	return s.GetToken(gtinParserEOF, 0)
}

func (s *GtinContext) Gtin8() IGtin8Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtin8Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtin8Context)
}

func (s *GtinContext) Gtin12() IGtin12Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtin12Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtin12Context)
}

func (s *GtinContext) Gtin13() IGtin13Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtin13Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtin13Context)
}

func (s *GtinContext) Gtin14() IGtin14Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtin14Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtin14Context)
}

func (s *GtinContext) Supplemental_code() ISupplemental_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISupplemental_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISupplemental_codeContext)
}

func (s *GtinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GtinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGtin(s)
	}
}

func (s *GtinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGtin(s)
	}
}

func (p *gtinParser) Gtin() (localctx IGtinContext) {
	localctx = NewGtinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gtinParserRULE_gtin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(74)
			p.Gtin8()
		}

	case 2:
		{
			p.SetState(75)
			p.Gtin12()
		}

	case 3:
		{
			p.SetState(76)
			p.Gtin13()
		}

	case 4:
		{
			p.SetState(77)
			p.Gtin14()
		}

	case 5:
		{
			p.SetState(78)
			p.Supplemental_code()
		}

	}
	{
		p.SetState(81)
		p.Match(gtinParserEOF)
	}

	return localctx
}

// IGtin8Context is an interface to support dynamic dispatch.
type IGtin8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtin8Context differentiates from other interfaces.
	IsGtin8Context()
}

type Gtin8Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtin8Context() *Gtin8Context {
	var p = new(Gtin8Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gtin8
	return p
}

func (*Gtin8Context) IsGtin8Context() {}

func NewGtin8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gtin8Context {
	var p = new(Gtin8Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gtin8

	return p
}

func (s *Gtin8Context) GetParser() antlr.Parser { return s.parser }

func (s *Gtin8Context) Ean8() IEan8Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan8Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan8Context)
}

func (s *Gtin8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gtin8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gtin8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGtin8(s)
	}
}

func (s *Gtin8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGtin8(s)
	}
}

func (p *gtinParser) Gtin8() (localctx IGtin8Context) {
	localctx = NewGtin8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gtinParserRULE_gtin8)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Ean8()
	}

	return localctx
}

// IEan8Context is an interface to support dynamic dispatch.
type IEan8Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan8Context differentiates from other interfaces.
	IsEan8Context()
}

type Ean8Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan8Context() *Ean8Context {
	var p = new(Ean8Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean8
	return p
}

func (*Ean8Context) IsEan8Context() {}

func NewEan8Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean8Context {
	var p = new(Ean8Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean8

	return p
}

func (s *Ean8Context) GetParser() antlr.Parser { return s.parser }

func (s *Ean8Context) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ean8Context) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ean8Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean8Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean8Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan8(s)
	}
}

func (s *Ean8Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan8(s)
	}
}

func (p *gtinParser) Ean8() (localctx IEan8Context) {
	localctx = NewEan8Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gtinParserRULE_ean8)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Any_digit()
	}
	{
		p.SetState(86)
		p.Any_digit()
	}
	{
		p.SetState(87)
		p.Any_digit()
	}
	{
		p.SetState(88)
		p.Any_digit()
	}
	{
		p.SetState(89)
		p.Any_digit()
	}
	{
		p.SetState(90)
		p.Any_digit()
	}
	{
		p.SetState(91)
		p.Any_digit()
	}
	{
		p.SetState(92)
		p.Any_digit()
	}

	return localctx
}

// IGtin12Context is an interface to support dynamic dispatch.
type IGtin12Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtin12Context differentiates from other interfaces.
	IsGtin12Context()
}

type Gtin12Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtin12Context() *Gtin12Context {
	var p = new(Gtin12Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gtin12
	return p
}

func (*Gtin12Context) IsGtin12Context() {}

func NewGtin12Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gtin12Context {
	var p = new(Gtin12Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gtin12

	return p
}

func (s *Gtin12Context) GetParser() antlr.Parser { return s.parser }

func (s *Gtin12Context) Upc() IUpcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpcContext)
}

func (s *Gtin12Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gtin12Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gtin12Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGtin12(s)
	}
}

func (s *Gtin12Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGtin12(s)
	}
}

func (p *gtinParser) Gtin12() (localctx IGtin12Context) {
	localctx = NewGtin12Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gtinParserRULE_gtin12)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Upc()
	}

	return localctx
}

// IGtin13Context is an interface to support dynamic dispatch.
type IGtin13Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtin13Context differentiates from other interfaces.
	IsGtin13Context()
}

type Gtin13Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtin13Context() *Gtin13Context {
	var p = new(Gtin13Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gtin13
	return p
}

func (*Gtin13Context) IsGtin13Context() {}

func NewGtin13Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gtin13Context {
	var p = new(Gtin13Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gtin13

	return p
}

func (s *Gtin13Context) GetParser() antlr.Parser { return s.parser }

func (s *Gtin13Context) Ean13() IEan13Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan13Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan13Context)
}

func (s *Gtin13Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gtin13Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gtin13Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGtin13(s)
	}
}

func (s *Gtin13Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGtin13(s)
	}
}

func (p *gtinParser) Gtin13() (localctx IGtin13Context) {
	localctx = NewGtin13Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gtinParserRULE_gtin13)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Ean13()
	}

	return localctx
}

// IGtin14Context is an interface to support dynamic dispatch.
type IGtin14Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtin14Context differentiates from other interfaces.
	IsGtin14Context()
}

type Gtin14Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtin14Context() *Gtin14Context {
	var p = new(Gtin14Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gtin14
	return p
}

func (*Gtin14Context) IsGtin14Context() {}

func NewGtin14Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gtin14Context {
	var p = new(Gtin14Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gtin14

	return p
}

func (s *Gtin14Context) GetParser() antlr.Parser { return s.parser }

func (s *Gtin14Context) Ean14() IEan14Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan14Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan14Context)
}

func (s *Gtin14Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gtin14Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gtin14Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGtin14(s)
	}
}

func (s *Gtin14Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGtin14(s)
	}
}

func (p *gtinParser) Gtin14() (localctx IGtin14Context) {
	localctx = NewGtin14Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gtinParserRULE_gtin14)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Ean14()
	}

	return localctx
}

// IUpcContext is an interface to support dynamic dispatch.
type IUpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpcContext differentiates from other interfaces.
	IsUpcContext()
}

type UpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpcContext() *UpcContext {
	var p = new(UpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc
	return p
}

func (*UpcContext) IsUpcContext() {}

func NewUpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpcContext {
	var p = new(UpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc

	return p
}

func (s *UpcContext) GetParser() antlr.Parser { return s.parser }

func (s *UpcContext) Upc_a() IUpc_aContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_aContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_aContext)
}

func (s *UpcContext) Upc_e() IUpc_eContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_eContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_eContext)
}

func (s *UpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc(s)
	}
}

func (s *UpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc(s)
	}
}

func (p *gtinParser) Upc() (localctx IUpcContext) {
	localctx = NewUpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gtinParserRULE_upc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(100)
			p.Upc_a()
		}

	case 2:
		{
			p.SetState(101)
			p.Upc_e()
		}

	}

	return localctx
}

// IUpc_aContext is an interface to support dynamic dispatch.
type IUpc_aContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpc_aContext differentiates from other interfaces.
	IsUpc_aContext()
}

type Upc_aContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpc_aContext() *Upc_aContext {
	var p = new(Upc_aContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc_a
	return p
}

func (*Upc_aContext) IsUpc_aContext() {}

func NewUpc_aContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upc_aContext {
	var p = new(Upc_aContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc_a

	return p
}

func (s *Upc_aContext) GetParser() antlr.Parser { return s.parser }

func (s *Upc_aContext) Num_system() INum_systemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INum_systemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INum_systemContext)
}

func (s *Upc_aContext) Upc_a_manufacturer() IUpc_a_manufacturerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_a_manufacturerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_a_manufacturerContext)
}

func (s *Upc_aContext) Upc_a_product() IUpc_a_productContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_a_productContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_a_productContext)
}

func (s *Upc_aContext) Check_code() ICheck_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheck_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheck_codeContext)
}

func (s *Upc_aContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upc_aContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upc_aContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc_a(s)
	}
}

func (s *Upc_aContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc_a(s)
	}
}

func (p *gtinParser) Upc_a() (localctx IUpc_aContext) {
	localctx = NewUpc_aContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gtinParserRULE_upc_a)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Num_system()
	}
	{
		p.SetState(105)
		p.Upc_a_manufacturer()
	}
	{
		p.SetState(106)
		p.Upc_a_product()
	}
	{
		p.SetState(107)
		p.Check_code()
	}

	return localctx
}

// IUpc_a_manufacturerContext is an interface to support dynamic dispatch.
type IUpc_a_manufacturerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpc_a_manufacturerContext differentiates from other interfaces.
	IsUpc_a_manufacturerContext()
}

type Upc_a_manufacturerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpc_a_manufacturerContext() *Upc_a_manufacturerContext {
	var p = new(Upc_a_manufacturerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc_a_manufacturer
	return p
}

func (*Upc_a_manufacturerContext) IsUpc_a_manufacturerContext() {}

func NewUpc_a_manufacturerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upc_a_manufacturerContext {
	var p = new(Upc_a_manufacturerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc_a_manufacturer

	return p
}

func (s *Upc_a_manufacturerContext) GetParser() antlr.Parser { return s.parser }

func (s *Upc_a_manufacturerContext) Upc_a_5() IUpc_a_5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_a_5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_a_5Context)
}

func (s *Upc_a_manufacturerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upc_a_manufacturerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upc_a_manufacturerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc_a_manufacturer(s)
	}
}

func (s *Upc_a_manufacturerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc_a_manufacturer(s)
	}
}

func (p *gtinParser) Upc_a_manufacturer() (localctx IUpc_a_manufacturerContext) {
	localctx = NewUpc_a_manufacturerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gtinParserRULE_upc_a_manufacturer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Upc_a_5()
	}

	return localctx
}

// IUpc_a_productContext is an interface to support dynamic dispatch.
type IUpc_a_productContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpc_a_productContext differentiates from other interfaces.
	IsUpc_a_productContext()
}

type Upc_a_productContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpc_a_productContext() *Upc_a_productContext {
	var p = new(Upc_a_productContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc_a_product
	return p
}

func (*Upc_a_productContext) IsUpc_a_productContext() {}

func NewUpc_a_productContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upc_a_productContext {
	var p = new(Upc_a_productContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc_a_product

	return p
}

func (s *Upc_a_productContext) GetParser() antlr.Parser { return s.parser }

func (s *Upc_a_productContext) Upc_a_5() IUpc_a_5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpc_a_5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpc_a_5Context)
}

func (s *Upc_a_productContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upc_a_productContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upc_a_productContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc_a_product(s)
	}
}

func (s *Upc_a_productContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc_a_product(s)
	}
}

func (p *gtinParser) Upc_a_product() (localctx IUpc_a_productContext) {
	localctx = NewUpc_a_productContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gtinParserRULE_upc_a_product)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Upc_a_5()
	}

	return localctx
}

// IUpc_a_5Context is an interface to support dynamic dispatch.
type IUpc_a_5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpc_a_5Context differentiates from other interfaces.
	IsUpc_a_5Context()
}

type Upc_a_5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpc_a_5Context() *Upc_a_5Context {
	var p = new(Upc_a_5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc_a_5
	return p
}

func (*Upc_a_5Context) IsUpc_a_5Context() {}

func NewUpc_a_5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upc_a_5Context {
	var p = new(Upc_a_5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc_a_5

	return p
}

func (s *Upc_a_5Context) GetParser() antlr.Parser { return s.parser }

func (s *Upc_a_5Context) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Upc_a_5Context) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Upc_a_5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upc_a_5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upc_a_5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc_a_5(s)
	}
}

func (s *Upc_a_5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc_a_5(s)
	}
}

func (p *gtinParser) Upc_a_5() (localctx IUpc_a_5Context) {
	localctx = NewUpc_a_5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gtinParserRULE_upc_a_5)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Any_digit()
	}
	{
		p.SetState(114)
		p.Any_digit()
	}
	{
		p.SetState(115)
		p.Any_digit()
	}
	{
		p.SetState(116)
		p.Any_digit()
	}
	{
		p.SetState(117)
		p.Any_digit()
	}

	return localctx
}

// IUpc_eContext is an interface to support dynamic dispatch.
type IUpc_eContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpc_eContext differentiates from other interfaces.
	IsUpc_eContext()
}

type Upc_eContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpc_eContext() *Upc_eContext {
	var p = new(Upc_eContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_upc_e
	return p
}

func (*Upc_eContext) IsUpc_eContext() {}

func NewUpc_eContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Upc_eContext {
	var p = new(Upc_eContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_upc_e

	return p
}

func (s *Upc_eContext) GetParser() antlr.Parser { return s.parser }

func (s *Upc_eContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Upc_eContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Upc_eContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Upc_eContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Upc_eContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterUpc_e(s)
	}
}

func (s *Upc_eContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitUpc_e(s)
	}
}

func (p *gtinParser) Upc_e() (localctx IUpc_eContext) {
	localctx = NewUpc_eContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gtinParserRULE_upc_e)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Any_digit()
	}
	{
		p.SetState(120)
		p.Any_digit()
	}
	{
		p.SetState(121)
		p.Any_digit()
	}
	{
		p.SetState(122)
		p.Any_digit()
	}
	{
		p.SetState(123)
		p.Any_digit()
	}
	{
		p.SetState(124)
		p.Any_digit()
	}

	return localctx
}

// INum_systemContext is an interface to support dynamic dispatch.
type INum_systemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNum_systemContext differentiates from other interfaces.
	IsNum_systemContext()
}

type Num_systemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNum_systemContext() *Num_systemContext {
	var p = new(Num_systemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_num_system
	return p
}

func (*Num_systemContext) IsNum_systemContext() {}

func NewNum_systemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Num_systemContext {
	var p = new(Num_systemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_num_system

	return p
}

func (s *Num_systemContext) GetParser() antlr.Parser { return s.parser }

func (s *Num_systemContext) Any_digit() IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Num_systemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Num_systemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Num_systemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterNum_system(s)
	}
}

func (s *Num_systemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitNum_system(s)
	}
}

func (p *gtinParser) Num_system() (localctx INum_systemContext) {
	localctx = NewNum_systemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gtinParserRULE_num_system)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Any_digit()
	}

	return localctx
}

// ICheck_codeContext is an interface to support dynamic dispatch.
type ICheck_codeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheck_codeContext differentiates from other interfaces.
	IsCheck_codeContext()
}

type Check_codeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheck_codeContext() *Check_codeContext {
	var p = new(Check_codeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_check_code
	return p
}

func (*Check_codeContext) IsCheck_codeContext() {}

func NewCheck_codeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Check_codeContext {
	var p = new(Check_codeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_check_code

	return p
}

func (s *Check_codeContext) GetParser() antlr.Parser { return s.parser }

func (s *Check_codeContext) Any_digit() IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Check_codeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Check_codeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Check_codeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterCheck_code(s)
	}
}

func (s *Check_codeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitCheck_code(s)
	}
}

func (p *gtinParser) Check_code() (localctx ICheck_codeContext) {
	localctx = NewCheck_codeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gtinParserRULE_check_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Any_digit()
	}

	return localctx
}

// ISupplemental_codeContext is an interface to support dynamic dispatch.
type ISupplemental_codeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSupplemental_codeContext differentiates from other interfaces.
	IsSupplemental_codeContext()
}

type Supplemental_codeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupplemental_codeContext() *Supplemental_codeContext {
	var p = new(Supplemental_codeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_supplemental_code
	return p
}

func (*Supplemental_codeContext) IsSupplemental_codeContext() {}

func NewSupplemental_codeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Supplemental_codeContext {
	var p = new(Supplemental_codeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_supplemental_code

	return p
}

func (s *Supplemental_codeContext) GetParser() antlr.Parser { return s.parser }

func (s *Supplemental_codeContext) Supplemental_code_5() ISupplemental_code_5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISupplemental_code_5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISupplemental_code_5Context)
}

func (s *Supplemental_codeContext) Supplemental_code_2() ISupplemental_code_2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISupplemental_code_2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISupplemental_code_2Context)
}

func (s *Supplemental_codeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Supplemental_codeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Supplemental_codeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterSupplemental_code(s)
	}
}

func (s *Supplemental_codeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitSupplemental_code(s)
	}
}

func (p *gtinParser) Supplemental_code() (localctx ISupplemental_codeContext) {
	localctx = NewSupplemental_codeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gtinParserRULE_supplemental_code)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Supplemental_code_5()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Supplemental_code_2()
		}

	}

	return localctx
}

// ISupplemental_code_5Context is an interface to support dynamic dispatch.
type ISupplemental_code_5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSupplemental_code_5Context differentiates from other interfaces.
	IsSupplemental_code_5Context()
}

type Supplemental_code_5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupplemental_code_5Context() *Supplemental_code_5Context {
	var p = new(Supplemental_code_5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_supplemental_code_5
	return p
}

func (*Supplemental_code_5Context) IsSupplemental_code_5Context() {}

func NewSupplemental_code_5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Supplemental_code_5Context {
	var p = new(Supplemental_code_5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_supplemental_code_5

	return p
}

func (s *Supplemental_code_5Context) GetParser() antlr.Parser { return s.parser }

func (s *Supplemental_code_5Context) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Supplemental_code_5Context) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Supplemental_code_5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Supplemental_code_5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Supplemental_code_5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterSupplemental_code_5(s)
	}
}

func (s *Supplemental_code_5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitSupplemental_code_5(s)
	}
}

func (p *gtinParser) Supplemental_code_5() (localctx ISupplemental_code_5Context) {
	localctx = NewSupplemental_code_5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gtinParserRULE_supplemental_code_5)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Any_digit()
	}
	{
		p.SetState(135)
		p.Any_digit()
	}
	{
		p.SetState(136)
		p.Any_digit()
	}
	{
		p.SetState(137)
		p.Any_digit()
	}
	{
		p.SetState(138)
		p.Any_digit()
	}

	return localctx
}

// ISupplemental_code_2Context is an interface to support dynamic dispatch.
type ISupplemental_code_2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSupplemental_code_2Context differentiates from other interfaces.
	IsSupplemental_code_2Context()
}

type Supplemental_code_2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySupplemental_code_2Context() *Supplemental_code_2Context {
	var p = new(Supplemental_code_2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_supplemental_code_2
	return p
}

func (*Supplemental_code_2Context) IsSupplemental_code_2Context() {}

func NewSupplemental_code_2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Supplemental_code_2Context {
	var p = new(Supplemental_code_2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_supplemental_code_2

	return p
}

func (s *Supplemental_code_2Context) GetParser() antlr.Parser { return s.parser }

func (s *Supplemental_code_2Context) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Supplemental_code_2Context) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Supplemental_code_2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Supplemental_code_2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Supplemental_code_2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterSupplemental_code_2(s)
	}
}

func (s *Supplemental_code_2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitSupplemental_code_2(s)
	}
}

func (p *gtinParser) Supplemental_code_2() (localctx ISupplemental_code_2Context) {
	localctx = NewSupplemental_code_2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gtinParserRULE_supplemental_code_2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Any_digit()
	}
	{
		p.SetState(141)
		p.Any_digit()
	}

	return localctx
}

// IEan13Context is an interface to support dynamic dispatch.
type IEan13Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan13Context differentiates from other interfaces.
	IsEan13Context()
}

type Ean13Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan13Context() *Ean13Context {
	var p = new(Ean13Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean13
	return p
}

func (*Ean13Context) IsEan13Context() {}

func NewEan13Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean13Context {
	var p = new(Ean13Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean13

	return p
}

func (s *Ean13Context) GetParser() antlr.Parser { return s.parser }

func (s *Ean13Context) Ean13_ismn() IEan13_ismnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan13_ismnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan13_ismnContext)
}

func (s *Ean13Context) Ean13_issn() IEan13_issnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan13_issnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan13_issnContext)
}

func (s *Ean13Context) Ean13_bookland() IEan13_booklandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan13_booklandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan13_booklandContext)
}

func (s *Ean13Context) Ean13_generic() IEan13_genericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan13_genericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan13_genericContext)
}

func (s *Ean13Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean13Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean13Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan13(s)
	}
}

func (s *Ean13Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan13(s)
	}
}

func (p *gtinParser) Ean13() (localctx IEan13Context) {
	localctx = NewEan13Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gtinParserRULE_ean13)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(143)
			p.Ean13_ismn()
		}

	case 2:
		{
			p.SetState(144)
			p.Ean13_issn()
		}

	case 3:
		{
			p.SetState(145)
			p.Ean13_bookland()
		}

	case 4:
		{
			p.SetState(146)
			p.Ean13_generic()
		}

	}

	return localctx
}

// IEan13_genericContext is an interface to support dynamic dispatch.
type IEan13_genericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan13_genericContext differentiates from other interfaces.
	IsEan13_genericContext()
}

type Ean13_genericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan13_genericContext() *Ean13_genericContext {
	var p = new(Ean13_genericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean13_generic
	return p
}

func (*Ean13_genericContext) IsEan13_genericContext() {}

func NewEan13_genericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean13_genericContext {
	var p = new(Ean13_genericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean13_generic

	return p
}

func (s *Ean13_genericContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean13_genericContext) Gs1_prefix() IGs1_prefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGs1_prefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGs1_prefixContext)
}

func (s *Ean13_genericContext) Ean_13_manprod() IEan_13_manprodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan_13_manprodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan_13_manprodContext)
}

func (s *Ean13_genericContext) Check_code() ICheck_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheck_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheck_codeContext)
}

func (s *Ean13_genericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean13_genericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean13_genericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan13_generic(s)
	}
}

func (s *Ean13_genericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan13_generic(s)
	}
}

func (p *gtinParser) Ean13_generic() (localctx IEan13_genericContext) {
	localctx = NewEan13_genericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gtinParserRULE_ean13_generic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Gs1_prefix()
	}
	{
		p.SetState(150)
		p.Ean_13_manprod()
	}
	{
		p.SetState(151)
		p.Check_code()
	}

	return localctx
}

// IEan13_ismnContext is an interface to support dynamic dispatch.
type IEan13_ismnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan13_ismnContext differentiates from other interfaces.
	IsEan13_ismnContext()
}

type Ean13_ismnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan13_ismnContext() *Ean13_ismnContext {
	var p = new(Ean13_ismnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean13_ismn
	return p
}

func (*Ean13_ismnContext) IsEan13_ismnContext() {}

func NewEan13_ismnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean13_ismnContext {
	var p = new(Ean13_ismnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean13_ismn

	return p
}

func (s *Ean13_ismnContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean13_ismnContext) Gs1_prefix_ismn() IGs1_prefix_ismnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGs1_prefix_ismnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGs1_prefix_ismnContext)
}

func (s *Ean13_ismnContext) Ismn_publisher_number() IIsmn_publisher_numberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsmn_publisher_numberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsmn_publisher_numberContext)
}

func (s *Ean13_ismnContext) Ismn_item_number() IIsmn_item_numberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsmn_item_numberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsmn_item_numberContext)
}

func (s *Ean13_ismnContext) Check_code() ICheck_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheck_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheck_codeContext)
}

func (s *Ean13_ismnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean13_ismnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean13_ismnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan13_ismn(s)
	}
}

func (s *Ean13_ismnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan13_ismn(s)
	}
}

func (p *gtinParser) Ean13_ismn() (localctx IEan13_ismnContext) {
	localctx = NewEan13_ismnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gtinParserRULE_ean13_ismn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Gs1_prefix_ismn()
	}
	{
		p.SetState(154)
		p.Ismn_publisher_number()
	}
	{
		p.SetState(155)
		p.Ismn_item_number()
	}
	{
		p.SetState(156)
		p.Check_code()
	}

	return localctx
}

// IGs1_prefix_ismnContext is an interface to support dynamic dispatch.
type IGs1_prefix_ismnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGs1_prefix_ismnContext differentiates from other interfaces.
	IsGs1_prefix_ismnContext()
}

type Gs1_prefix_ismnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGs1_prefix_ismnContext() *Gs1_prefix_ismnContext {
	var p = new(Gs1_prefix_ismnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gs1_prefix_ismn
	return p
}

func (*Gs1_prefix_ismnContext) IsGs1_prefix_ismnContext() {}

func NewGs1_prefix_ismnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gs1_prefix_ismnContext {
	var p = new(Gs1_prefix_ismnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gs1_prefix_ismn

	return p
}

func (s *Gs1_prefix_ismnContext) GetParser() antlr.Parser { return s.parser }

func (s *Gs1_prefix_ismnContext) AllDIGIT_9() []antlr.TerminalNode {
	return s.GetTokens(gtinParserDIGIT_9)
}

func (s *Gs1_prefix_ismnContext) DIGIT_9(i int) antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_9, i)
}

func (s *Gs1_prefix_ismnContext) DIGIT_7() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, 0)
}

func (s *Gs1_prefix_ismnContext) DIGIT_0() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_0, 0)
}

func (s *Gs1_prefix_ismnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gs1_prefix_ismnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gs1_prefix_ismnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGs1_prefix_ismn(s)
	}
}

func (s *Gs1_prefix_ismnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGs1_prefix_ismn(s)
	}
}

func (p *gtinParser) Gs1_prefix_ismn() (localctx IGs1_prefix_ismnContext) {
	localctx = NewGs1_prefix_ismnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gtinParserRULE_gs1_prefix_ismn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(159)
		p.Match(gtinParserDIGIT_7)
	}
	{
		p.SetState(160)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(161)
		p.Match(gtinParserDIGIT_0)
	}

	return localctx
}

// IIsmn_publisher_numberContext is an interface to support dynamic dispatch.
type IIsmn_publisher_numberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsmn_publisher_numberContext differentiates from other interfaces.
	IsIsmn_publisher_numberContext()
}

type Ismn_publisher_numberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsmn_publisher_numberContext() *Ismn_publisher_numberContext {
	var p = new(Ismn_publisher_numberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ismn_publisher_number
	return p
}

func (*Ismn_publisher_numberContext) IsIsmn_publisher_numberContext() {}

func NewIsmn_publisher_numberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ismn_publisher_numberContext {
	var p = new(Ismn_publisher_numberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ismn_publisher_number

	return p
}

func (s *Ismn_publisher_numberContext) GetParser() antlr.Parser { return s.parser }

func (s *Ismn_publisher_numberContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ismn_publisher_numberContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ismn_publisher_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ismn_publisher_numberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ismn_publisher_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterIsmn_publisher_number(s)
	}
}

func (s *Ismn_publisher_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitIsmn_publisher_number(s)
	}
}

func (p *gtinParser) Ismn_publisher_number() (localctx IIsmn_publisher_numberContext) {
	localctx = NewIsmn_publisher_numberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gtinParserRULE_ismn_publisher_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Any_digit()
	}
	{
		p.SetState(164)
		p.Any_digit()
	}
	{
		p.SetState(165)
		p.Any_digit()
	}
	{
		p.SetState(166)
		p.Any_digit()
	}

	return localctx
}

// IIsmn_item_numberContext is an interface to support dynamic dispatch.
type IIsmn_item_numberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsmn_item_numberContext differentiates from other interfaces.
	IsIsmn_item_numberContext()
}

type Ismn_item_numberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsmn_item_numberContext() *Ismn_item_numberContext {
	var p = new(Ismn_item_numberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ismn_item_number
	return p
}

func (*Ismn_item_numberContext) IsIsmn_item_numberContext() {}

func NewIsmn_item_numberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ismn_item_numberContext {
	var p = new(Ismn_item_numberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ismn_item_number

	return p
}

func (s *Ismn_item_numberContext) GetParser() antlr.Parser { return s.parser }

func (s *Ismn_item_numberContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ismn_item_numberContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ismn_item_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ismn_item_numberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ismn_item_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterIsmn_item_number(s)
	}
}

func (s *Ismn_item_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitIsmn_item_number(s)
	}
}

func (p *gtinParser) Ismn_item_number() (localctx IIsmn_item_numberContext) {
	localctx = NewIsmn_item_numberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gtinParserRULE_ismn_item_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Any_digit()
	}
	{
		p.SetState(169)
		p.Any_digit()
	}
	{
		p.SetState(170)
		p.Any_digit()
	}
	{
		p.SetState(171)
		p.Any_digit()
	}

	return localctx
}

// IEan13_booklandContext is an interface to support dynamic dispatch.
type IEan13_booklandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan13_booklandContext differentiates from other interfaces.
	IsEan13_booklandContext()
}

type Ean13_booklandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan13_booklandContext() *Ean13_booklandContext {
	var p = new(Ean13_booklandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean13_bookland
	return p
}

func (*Ean13_booklandContext) IsEan13_booklandContext() {}

func NewEan13_booklandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean13_booklandContext {
	var p = new(Ean13_booklandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean13_bookland

	return p
}

func (s *Ean13_booklandContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean13_booklandContext) Bookland_isbn() IBookland_isbnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBookland_isbnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBookland_isbnContext)
}

func (s *Ean13_booklandContext) Gs1_prefix_bookland_1() IGs1_prefix_bookland_1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGs1_prefix_bookland_1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGs1_prefix_bookland_1Context)
}

func (s *Ean13_booklandContext) Gs1_prefix_bookland_2() IGs1_prefix_bookland_2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGs1_prefix_bookland_2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGs1_prefix_bookland_2Context)
}

func (s *Ean13_booklandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean13_booklandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean13_booklandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan13_bookland(s)
	}
}

func (s *Ean13_booklandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan13_bookland(s)
	}
}

func (p *gtinParser) Ean13_bookland() (localctx IEan13_booklandContext) {
	localctx = NewEan13_booklandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gtinParserRULE_ean13_bookland)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(173)
			p.Gs1_prefix_bookland_1()
		}

	case 2:
		{
			p.SetState(174)
			p.Gs1_prefix_bookland_2()
		}

	}
	{
		p.SetState(177)
		p.Bookland_isbn()
	}

	return localctx
}

// IBookland_isbnContext is an interface to support dynamic dispatch.
type IBookland_isbnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBookland_isbnContext differentiates from other interfaces.
	IsBookland_isbnContext()
}

type Bookland_isbnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBookland_isbnContext() *Bookland_isbnContext {
	var p = new(Bookland_isbnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_bookland_isbn
	return p
}

func (*Bookland_isbnContext) IsBookland_isbnContext() {}

func NewBookland_isbnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bookland_isbnContext {
	var p = new(Bookland_isbnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_bookland_isbn

	return p
}

func (s *Bookland_isbnContext) GetParser() antlr.Parser { return s.parser }

func (s *Bookland_isbnContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Bookland_isbnContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Bookland_isbnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bookland_isbnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bookland_isbnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterBookland_isbn(s)
	}
}

func (s *Bookland_isbnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitBookland_isbn(s)
	}
}

func (p *gtinParser) Bookland_isbn() (localctx IBookland_isbnContext) {
	localctx = NewBookland_isbnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gtinParserRULE_bookland_isbn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Any_digit()
	}
	{
		p.SetState(180)
		p.Any_digit()
	}
	{
		p.SetState(181)
		p.Any_digit()
	}
	{
		p.SetState(182)
		p.Any_digit()
	}
	{
		p.SetState(183)
		p.Any_digit()
	}
	{
		p.SetState(184)
		p.Any_digit()
	}
	{
		p.SetState(185)
		p.Any_digit()
	}
	{
		p.SetState(186)
		p.Any_digit()
	}
	{
		p.SetState(187)
		p.Any_digit()
	}

	return localctx
}

// IGs1_prefix_bookland_1Context is an interface to support dynamic dispatch.
type IGs1_prefix_bookland_1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGs1_prefix_bookland_1Context differentiates from other interfaces.
	IsGs1_prefix_bookland_1Context()
}

type Gs1_prefix_bookland_1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGs1_prefix_bookland_1Context() *Gs1_prefix_bookland_1Context {
	var p = new(Gs1_prefix_bookland_1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gs1_prefix_bookland_1
	return p
}

func (*Gs1_prefix_bookland_1Context) IsGs1_prefix_bookland_1Context() {}

func NewGs1_prefix_bookland_1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gs1_prefix_bookland_1Context {
	var p = new(Gs1_prefix_bookland_1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gs1_prefix_bookland_1

	return p
}

func (s *Gs1_prefix_bookland_1Context) GetParser() antlr.Parser { return s.parser }

func (s *Gs1_prefix_bookland_1Context) AllDIGIT_9() []antlr.TerminalNode {
	return s.GetTokens(gtinParserDIGIT_9)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_9(i int) antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_9, i)
}

func (s *Gs1_prefix_bookland_1Context) AllDIGIT_7() []antlr.TerminalNode {
	return s.GetTokens(gtinParserDIGIT_7)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_7(i int) antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, i)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_1() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_1, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_2() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_2, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_3() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_3, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_4() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_4, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_5() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_5, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_6() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_6, 0)
}

func (s *Gs1_prefix_bookland_1Context) DIGIT_8() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_8, 0)
}

func (s *Gs1_prefix_bookland_1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gs1_prefix_bookland_1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gs1_prefix_bookland_1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGs1_prefix_bookland_1(s)
	}
}

func (s *Gs1_prefix_bookland_1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGs1_prefix_bookland_1(s)
	}
}

func (p *gtinParser) Gs1_prefix_bookland_1() (localctx IGs1_prefix_bookland_1Context) {
	localctx = NewGs1_prefix_bookland_1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gtinParserRULE_gs1_prefix_bookland_1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(190)
		p.Match(gtinParserDIGIT_7)
	}
	{
		p.SetState(191)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gtinParserDIGIT_1)|(1<<gtinParserDIGIT_2)|(1<<gtinParserDIGIT_3)|(1<<gtinParserDIGIT_4)|(1<<gtinParserDIGIT_5)|(1<<gtinParserDIGIT_6)|(1<<gtinParserDIGIT_7)|(1<<gtinParserDIGIT_8)|(1<<gtinParserDIGIT_9))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGs1_prefix_bookland_2Context is an interface to support dynamic dispatch.
type IGs1_prefix_bookland_2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGs1_prefix_bookland_2Context differentiates from other interfaces.
	IsGs1_prefix_bookland_2Context()
}

type Gs1_prefix_bookland_2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGs1_prefix_bookland_2Context() *Gs1_prefix_bookland_2Context {
	var p = new(Gs1_prefix_bookland_2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gs1_prefix_bookland_2
	return p
}

func (*Gs1_prefix_bookland_2Context) IsGs1_prefix_bookland_2Context() {}

func NewGs1_prefix_bookland_2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gs1_prefix_bookland_2Context {
	var p = new(Gs1_prefix_bookland_2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gs1_prefix_bookland_2

	return p
}

func (s *Gs1_prefix_bookland_2Context) GetParser() antlr.Parser { return s.parser }

func (s *Gs1_prefix_bookland_2Context) DIGIT_9() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_9, 0)
}

func (s *Gs1_prefix_bookland_2Context) DIGIT_7() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, 0)
}

func (s *Gs1_prefix_bookland_2Context) DIGIT_8() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_8, 0)
}

func (s *Gs1_prefix_bookland_2Context) Any_digit() IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Gs1_prefix_bookland_2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gs1_prefix_bookland_2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gs1_prefix_bookland_2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGs1_prefix_bookland_2(s)
	}
}

func (s *Gs1_prefix_bookland_2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGs1_prefix_bookland_2(s)
	}
}

func (p *gtinParser) Gs1_prefix_bookland_2() (localctx IGs1_prefix_bookland_2Context) {
	localctx = NewGs1_prefix_bookland_2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gtinParserRULE_gs1_prefix_bookland_2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(195)
		p.Match(gtinParserDIGIT_7)
	}
	{
		p.SetState(196)
		p.Match(gtinParserDIGIT_8)
	}
	{
		p.SetState(197)
		p.Any_digit()
	}

	return localctx
}

// IGs1_prefix_issnContext is an interface to support dynamic dispatch.
type IGs1_prefix_issnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGs1_prefix_issnContext differentiates from other interfaces.
	IsGs1_prefix_issnContext()
}

type Gs1_prefix_issnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGs1_prefix_issnContext() *Gs1_prefix_issnContext {
	var p = new(Gs1_prefix_issnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gs1_prefix_issn
	return p
}

func (*Gs1_prefix_issnContext) IsGs1_prefix_issnContext() {}

func NewGs1_prefix_issnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gs1_prefix_issnContext {
	var p = new(Gs1_prefix_issnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gs1_prefix_issn

	return p
}

func (s *Gs1_prefix_issnContext) GetParser() antlr.Parser { return s.parser }

func (s *Gs1_prefix_issnContext) DIGIT_9() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_9, 0)
}

func (s *Gs1_prefix_issnContext) AllDIGIT_7() []antlr.TerminalNode {
	return s.GetTokens(gtinParserDIGIT_7)
}

func (s *Gs1_prefix_issnContext) DIGIT_7(i int) antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, i)
}

func (s *Gs1_prefix_issnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gs1_prefix_issnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gs1_prefix_issnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGs1_prefix_issn(s)
	}
}

func (s *Gs1_prefix_issnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGs1_prefix_issn(s)
	}
}

func (p *gtinParser) Gs1_prefix_issn() (localctx IGs1_prefix_issnContext) {
	localctx = NewGs1_prefix_issnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gtinParserRULE_gs1_prefix_issn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(gtinParserDIGIT_9)
	}
	{
		p.SetState(200)
		p.Match(gtinParserDIGIT_7)
	}
	{
		p.SetState(201)
		p.Match(gtinParserDIGIT_7)
	}

	return localctx
}

// IEan13_issnContext is an interface to support dynamic dispatch.
type IEan13_issnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan13_issnContext differentiates from other interfaces.
	IsEan13_issnContext()
}

type Ean13_issnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan13_issnContext() *Ean13_issnContext {
	var p = new(Ean13_issnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean13_issn
	return p
}

func (*Ean13_issnContext) IsEan13_issnContext() {}

func NewEan13_issnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean13_issnContext {
	var p = new(Ean13_issnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean13_issn

	return p
}

func (s *Ean13_issnContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean13_issnContext) Gs1_prefix_issn() IGs1_prefix_issnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGs1_prefix_issnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGs1_prefix_issnContext)
}

func (s *Ean13_issnContext) Issn() IIssnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIssnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIssnContext)
}

func (s *Ean13_issnContext) Check_code() ICheck_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheck_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheck_codeContext)
}

func (s *Ean13_issnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean13_issnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean13_issnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan13_issn(s)
	}
}

func (s *Ean13_issnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan13_issn(s)
	}
}

func (p *gtinParser) Ean13_issn() (localctx IEan13_issnContext) {
	localctx = NewEan13_issnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gtinParserRULE_ean13_issn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Gs1_prefix_issn()
	}
	{
		p.SetState(204)
		p.Issn()
	}
	{
		p.SetState(205)
		p.Check_code()
	}

	return localctx
}

// IIssnContext is an interface to support dynamic dispatch.
type IIssnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIssnContext differentiates from other interfaces.
	IsIssnContext()
}

type IssnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIssnContext() *IssnContext {
	var p = new(IssnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_issn
	return p
}

func (*IssnContext) IsIssnContext() {}

func NewIssnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IssnContext {
	var p = new(IssnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_issn

	return p
}

func (s *IssnContext) GetParser() antlr.Parser { return s.parser }

func (s *IssnContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *IssnContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *IssnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IssnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IssnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterIssn(s)
	}
}

func (s *IssnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitIssn(s)
	}
}

func (p *gtinParser) Issn() (localctx IIssnContext) {
	localctx = NewIssnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gtinParserRULE_issn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Any_digit()
	}
	{
		p.SetState(208)
		p.Any_digit()
	}
	{
		p.SetState(209)
		p.Any_digit()
	}
	{
		p.SetState(210)
		p.Any_digit()
	}
	{
		p.SetState(211)
		p.Any_digit()
	}
	{
		p.SetState(212)
		p.Any_digit()
	}
	{
		p.SetState(213)
		p.Any_digit()
	}
	{
		p.SetState(214)
		p.Any_digit()
	}
	{
		p.SetState(215)
		p.Any_digit()
	}

	return localctx
}

// IEan_13_manprodContext is an interface to support dynamic dispatch.
type IEan_13_manprodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan_13_manprodContext differentiates from other interfaces.
	IsEan_13_manprodContext()
}

type Ean_13_manprodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan_13_manprodContext() *Ean_13_manprodContext {
	var p = new(Ean_13_manprodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean_13_manprod
	return p
}

func (*Ean_13_manprodContext) IsEan_13_manprodContext() {}

func NewEan_13_manprodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean_13_manprodContext {
	var p = new(Ean_13_manprodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean_13_manprod

	return p
}

func (s *Ean_13_manprodContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean_13_manprodContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ean_13_manprodContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ean_13_manprodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean_13_manprodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean_13_manprodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan_13_manprod(s)
	}
}

func (s *Ean_13_manprodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan_13_manprod(s)
	}
}

func (p *gtinParser) Ean_13_manprod() (localctx IEan_13_manprodContext) {
	localctx = NewEan_13_manprodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gtinParserRULE_ean_13_manprod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Any_digit()
	}
	{
		p.SetState(218)
		p.Any_digit()
	}
	{
		p.SetState(219)
		p.Any_digit()
	}
	{
		p.SetState(220)
		p.Any_digit()
	}
	{
		p.SetState(221)
		p.Any_digit()
	}
	{
		p.SetState(222)
		p.Any_digit()
	}
	{
		p.SetState(223)
		p.Any_digit()
	}
	{
		p.SetState(224)
		p.Any_digit()
	}
	{
		p.SetState(225)
		p.Any_digit()
	}

	return localctx
}

// IGs1_prefixContext is an interface to support dynamic dispatch.
type IGs1_prefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGs1_prefixContext differentiates from other interfaces.
	IsGs1_prefixContext()
}

type Gs1_prefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGs1_prefixContext() *Gs1_prefixContext {
	var p = new(Gs1_prefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_gs1_prefix
	return p
}

func (*Gs1_prefixContext) IsGs1_prefixContext() {}

func NewGs1_prefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gs1_prefixContext {
	var p = new(Gs1_prefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_gs1_prefix

	return p
}

func (s *Gs1_prefixContext) GetParser() antlr.Parser { return s.parser }

func (s *Gs1_prefixContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Gs1_prefixContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Gs1_prefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gs1_prefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gs1_prefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterGs1_prefix(s)
	}
}

func (s *Gs1_prefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitGs1_prefix(s)
	}
}

func (p *gtinParser) Gs1_prefix() (localctx IGs1_prefixContext) {
	localctx = NewGs1_prefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gtinParserRULE_gs1_prefix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Any_digit()
	}
	{
		p.SetState(228)
		p.Any_digit()
	}
	{
		p.SetState(229)
		p.Any_digit()
	}

	return localctx
}

// IEan14Context is an interface to support dynamic dispatch.
type IEan14Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan14Context differentiates from other interfaces.
	IsEan14Context()
}

type Ean14Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan14Context() *Ean14Context {
	var p = new(Ean14Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean14
	return p
}

func (*Ean14Context) IsEan14Context() {}

func NewEan14Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean14Context {
	var p = new(Ean14Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean14

	return p
}

func (s *Ean14Context) GetParser() antlr.Parser { return s.parser }

func (s *Ean14Context) Ean14_packaging() IEan14_packagingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan14_packagingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan14_packagingContext)
}

func (s *Ean14Context) Ean14_product() IEan14_productContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEan14_productContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEan14_productContext)
}

func (s *Ean14Context) Check_code() ICheck_codeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheck_codeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheck_codeContext)
}

func (s *Ean14Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean14Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean14Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan14(s)
	}
}

func (s *Ean14Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan14(s)
	}
}

func (p *gtinParser) Ean14() (localctx IEan14Context) {
	localctx = NewEan14Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gtinParserRULE_ean14)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Ean14_packaging()
	}
	{
		p.SetState(232)
		p.Ean14_product()
	}
	{
		p.SetState(233)
		p.Check_code()
	}

	return localctx
}

// IEan14_appidContext is an interface to support dynamic dispatch.
type IEan14_appidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan14_appidContext differentiates from other interfaces.
	IsEan14_appidContext()
}

type Ean14_appidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan14_appidContext() *Ean14_appidContext {
	var p = new(Ean14_appidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean14_appid
	return p
}

func (*Ean14_appidContext) IsEan14_appidContext() {}

func NewEan14_appidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean14_appidContext {
	var p = new(Ean14_appidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean14_appid

	return p
}

func (s *Ean14_appidContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean14_appidContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ean14_appidContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ean14_appidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean14_appidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean14_appidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan14_appid(s)
	}
}

func (s *Ean14_appidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan14_appid(s)
	}
}

func (p *gtinParser) Ean14_appid() (localctx IEan14_appidContext) {
	localctx = NewEan14_appidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gtinParserRULE_ean14_appid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Any_digit()
	}
	{
		p.SetState(236)
		p.Any_digit()
	}

	return localctx
}

// IEan14_packagingContext is an interface to support dynamic dispatch.
type IEan14_packagingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan14_packagingContext differentiates from other interfaces.
	IsEan14_packagingContext()
}

type Ean14_packagingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan14_packagingContext() *Ean14_packagingContext {
	var p = new(Ean14_packagingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean14_packaging
	return p
}

func (*Ean14_packagingContext) IsEan14_packagingContext() {}

func NewEan14_packagingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean14_packagingContext {
	var p = new(Ean14_packagingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean14_packaging

	return p
}

func (s *Ean14_packagingContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean14_packagingContext) DIGIT_0() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_0, 0)
}

func (s *Ean14_packagingContext) DIGIT_1() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_1, 0)
}

func (s *Ean14_packagingContext) DIGIT_2() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_2, 0)
}

func (s *Ean14_packagingContext) DIGIT_3() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_3, 0)
}

func (s *Ean14_packagingContext) DIGIT_4() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_4, 0)
}

func (s *Ean14_packagingContext) DIGIT_5() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_5, 0)
}

func (s *Ean14_packagingContext) DIGIT_6() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_6, 0)
}

func (s *Ean14_packagingContext) DIGIT_7() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, 0)
}

func (s *Ean14_packagingContext) DIGIT_8() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_8, 0)
}

func (s *Ean14_packagingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean14_packagingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean14_packagingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan14_packaging(s)
	}
}

func (s *Ean14_packagingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan14_packaging(s)
	}
}

func (p *gtinParser) Ean14_packaging() (localctx IEan14_packagingContext) {
	localctx = NewEan14_packagingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gtinParserRULE_ean14_packaging)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gtinParserDIGIT_0)|(1<<gtinParserDIGIT_1)|(1<<gtinParserDIGIT_2)|(1<<gtinParserDIGIT_3)|(1<<gtinParserDIGIT_4)|(1<<gtinParserDIGIT_5)|(1<<gtinParserDIGIT_6)|(1<<gtinParserDIGIT_7)|(1<<gtinParserDIGIT_8))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEan14_productContext is an interface to support dynamic dispatch.
type IEan14_productContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEan14_productContext differentiates from other interfaces.
	IsEan14_productContext()
}

type Ean14_productContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEan14_productContext() *Ean14_productContext {
	var p = new(Ean14_productContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_ean14_product
	return p
}

func (*Ean14_productContext) IsEan14_productContext() {}

func NewEan14_productContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ean14_productContext {
	var p = new(Ean14_productContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_ean14_product

	return p
}

func (s *Ean14_productContext) GetParser() antlr.Parser { return s.parser }

func (s *Ean14_productContext) AllAny_digit() []IAny_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAny_digitContext)(nil)).Elem())
	var tst = make([]IAny_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAny_digitContext)
		}
	}

	return tst
}

func (s *Ean14_productContext) Any_digit(i int) IAny_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAny_digitContext)
}

func (s *Ean14_productContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ean14_productContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ean14_productContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterEan14_product(s)
	}
}

func (s *Ean14_productContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitEan14_product(s)
	}
}

func (p *gtinParser) Ean14_product() (localctx IEan14_productContext) {
	localctx = NewEan14_productContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gtinParserRULE_ean14_product)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Any_digit()
	}
	{
		p.SetState(241)
		p.Any_digit()
	}
	{
		p.SetState(242)
		p.Any_digit()
	}
	{
		p.SetState(243)
		p.Any_digit()
	}
	{
		p.SetState(244)
		p.Any_digit()
	}
	{
		p.SetState(245)
		p.Any_digit()
	}
	{
		p.SetState(246)
		p.Any_digit()
	}
	{
		p.SetState(247)
		p.Any_digit()
	}
	{
		p.SetState(248)
		p.Any_digit()
	}
	{
		p.SetState(249)
		p.Any_digit()
	}
	{
		p.SetState(250)
		p.Any_digit()
	}
	{
		p.SetState(251)
		p.Any_digit()
	}

	return localctx
}

// IAny_digitContext is an interface to support dynamic dispatch.
type IAny_digitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_digitContext differentiates from other interfaces.
	IsAny_digitContext()
}

type Any_digitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_digitContext() *Any_digitContext {
	var p = new(Any_digitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gtinParserRULE_any_digit
	return p
}

func (*Any_digitContext) IsAny_digitContext() {}

func NewAny_digitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_digitContext {
	var p = new(Any_digitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gtinParserRULE_any_digit

	return p
}

func (s *Any_digitContext) GetParser() antlr.Parser { return s.parser }

func (s *Any_digitContext) DIGIT_0() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_0, 0)
}

func (s *Any_digitContext) DIGIT_1() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_1, 0)
}

func (s *Any_digitContext) DIGIT_2() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_2, 0)
}

func (s *Any_digitContext) DIGIT_3() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_3, 0)
}

func (s *Any_digitContext) DIGIT_4() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_4, 0)
}

func (s *Any_digitContext) DIGIT_5() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_5, 0)
}

func (s *Any_digitContext) DIGIT_6() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_6, 0)
}

func (s *Any_digitContext) DIGIT_7() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_7, 0)
}

func (s *Any_digitContext) DIGIT_8() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_8, 0)
}

func (s *Any_digitContext) DIGIT_9() antlr.TerminalNode {
	return s.GetToken(gtinParserDIGIT_9, 0)
}

func (s *Any_digitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_digitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_digitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.EnterAny_digit(s)
	}
}

func (s *Any_digitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gtinListener); ok {
		listenerT.ExitAny_digit(s)
	}
}

func (p *gtinParser) Any_digit() (localctx IAny_digitContext) {
	localctx = NewAny_digitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gtinParserRULE_any_digit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gtinParserDIGIT_0)|(1<<gtinParserDIGIT_1)|(1<<gtinParserDIGIT_2)|(1<<gtinParserDIGIT_3)|(1<<gtinParserDIGIT_4)|(1<<gtinParserDIGIT_5)|(1<<gtinParserDIGIT_6)|(1<<gtinParserDIGIT_7)|(1<<gtinParserDIGIT_8)|(1<<gtinParserDIGIT_9))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
