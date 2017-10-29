// Generated from xpath.g4 by ANTLR 4.7.

package xpath // xpath
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 242,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 3, 2, 3, 3, 3, 3, 5, 3, 61, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 67,
	10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 72, 10, 5, 12, 5, 14, 5, 75, 11, 5, 3, 6,
	3, 6, 3, 6, 7, 6, 80, 10, 6, 12, 6, 14, 6, 83, 11, 6, 3, 6, 5, 6, 86, 10,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 91, 10, 7, 5, 7, 93, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 103, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 121, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13,
	128, 10, 13, 12, 13, 14, 13, 131, 11, 13, 5, 13, 133, 10, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 5, 14, 140, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	145, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 151, 10, 15, 5, 15, 153,
	10, 15, 3, 16, 3, 16, 7, 16, 157, 10, 16, 12, 16, 14, 16, 160, 11, 16,
	3, 17, 3, 17, 3, 17, 7, 17, 165, 10, 17, 12, 17, 14, 17, 168, 11, 17, 3,
	18, 3, 18, 3, 18, 7, 18, 173, 10, 18, 12, 18, 14, 18, 176, 11, 18, 3, 19,
	3, 19, 3, 19, 7, 19, 181, 10, 19, 12, 19, 14, 19, 184, 11, 19, 3, 20, 3,
	20, 3, 20, 7, 20, 189, 10, 20, 12, 20, 14, 20, 192, 11, 20, 3, 21, 3, 21,
	3, 21, 7, 21, 197, 10, 21, 12, 21, 14, 21, 200, 11, 21, 3, 22, 3, 22, 3,
	22, 5, 22, 205, 10, 22, 3, 22, 3, 22, 3, 22, 5, 22, 210, 10, 22, 5, 22,
	212, 10, 22, 3, 23, 7, 23, 215, 10, 23, 12, 23, 14, 23, 218, 11, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 225, 10, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 238, 10,
	27, 3, 28, 3, 28, 3, 28, 2, 2, 29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 2,
	10, 3, 2, 14, 15, 4, 2, 22, 22, 24, 24, 3, 2, 6, 7, 3, 2, 28, 31, 3, 2,
	20, 21, 4, 2, 8, 9, 23, 23, 3, 2, 8, 9, 4, 2, 13, 13, 38, 38, 2, 246, 2,
	56, 3, 2, 2, 2, 4, 60, 3, 2, 2, 2, 6, 66, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2,
	10, 85, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14, 102, 3, 2, 2, 2, 16, 104, 3,
	2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 110, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2,
	24, 122, 3, 2, 2, 2, 26, 144, 3, 2, 2, 2, 28, 152, 3, 2, 2, 2, 30, 154,
	3, 2, 2, 2, 32, 161, 3, 2, 2, 2, 34, 169, 3, 2, 2, 2, 36, 177, 3, 2, 2,
	2, 38, 185, 3, 2, 2, 2, 40, 193, 3, 2, 2, 2, 42, 211, 3, 2, 2, 2, 44, 216,
	3, 2, 2, 2, 46, 221, 3, 2, 2, 2, 48, 226, 3, 2, 2, 2, 50, 228, 3, 2, 2,
	2, 52, 237, 3, 2, 2, 2, 54, 239, 3, 2, 2, 2, 56, 57, 5, 20, 11, 2, 57,
	3, 3, 2, 2, 2, 58, 61, 5, 8, 5, 2, 59, 61, 5, 6, 4, 2, 60, 58, 3, 2, 2,
	2, 60, 59, 3, 2, 2, 2, 61, 5, 3, 2, 2, 2, 62, 63, 7, 14, 2, 2, 63, 67,
	5, 8, 5, 2, 64, 65, 7, 15, 2, 2, 65, 67, 5, 8, 5, 2, 66, 62, 3, 2, 2, 2,
	66, 64, 3, 2, 2, 2, 67, 7, 3, 2, 2, 2, 68, 73, 5, 10, 6, 2, 69, 70, 9,
	2, 2, 2, 70, 72, 5, 10, 6, 2, 71, 69, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73,
	71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 9, 3, 2, 2, 2, 75, 73, 3, 2, 2,
	2, 76, 77, 5, 12, 7, 2, 77, 81, 5, 14, 8, 2, 78, 80, 5, 16, 9, 2, 79, 78,
	3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 86, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 86, 5, 18, 10, 2, 85, 76, 3,
	2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 88, 7, 13, 2, 2, 88,
	93, 7, 33, 2, 2, 89, 91, 7, 25, 2, 2, 90, 89, 3, 2, 2, 2, 90, 91, 3, 2,
	2, 2, 91, 93, 3, 2, 2, 2, 92, 87, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 13,
	3, 2, 2, 2, 94, 103, 5, 52, 27, 2, 95, 96, 7, 11, 2, 2, 96, 97, 7, 16,
	2, 2, 97, 103, 7, 17, 2, 2, 98, 99, 7, 3, 2, 2, 99, 100, 7, 16, 2, 2, 100,
	101, 7, 36, 2, 2, 101, 103, 7, 17, 2, 2, 102, 94, 3, 2, 2, 2, 102, 95,
	3, 2, 2, 2, 102, 98, 3, 2, 2, 2, 103, 15, 3, 2, 2, 2, 104, 105, 7, 18,
	2, 2, 105, 106, 5, 20, 11, 2, 106, 107, 7, 19, 2, 2, 107, 17, 3, 2, 2,
	2, 108, 109, 9, 3, 2, 2, 109, 19, 3, 2, 2, 2, 110, 111, 5, 32, 17, 2, 111,
	21, 3, 2, 2, 2, 112, 121, 5, 50, 26, 2, 113, 114, 7, 16, 2, 2, 114, 115,
	5, 20, 11, 2, 115, 116, 7, 17, 2, 2, 116, 121, 3, 2, 2, 2, 117, 121, 7,
	36, 2, 2, 118, 121, 7, 12, 2, 2, 119, 121, 5, 24, 13, 2, 120, 112, 3, 2,
	2, 2, 120, 113, 3, 2, 2, 2, 120, 117, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2,
	120, 119, 3, 2, 2, 2, 121, 23, 3, 2, 2, 2, 122, 123, 5, 48, 25, 2, 123,
	132, 7, 16, 2, 2, 124, 129, 5, 20, 11, 2, 125, 126, 7, 26, 2, 2, 126, 128,
	5, 20, 11, 2, 127, 125, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3,
	2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2,
	2, 132, 124, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134,
	135, 7, 17, 2, 2, 135, 25, 3, 2, 2, 2, 136, 139, 5, 28, 15, 2, 137, 138,
	7, 27, 2, 2, 138, 140, 5, 26, 14, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3,
	2, 2, 2, 140, 145, 3, 2, 2, 2, 141, 142, 7, 14, 2, 2, 142, 143, 7, 27,
	2, 2, 143, 145, 5, 26, 14, 2, 144, 136, 3, 2, 2, 2, 144, 141, 3, 2, 2,
	2, 145, 27, 3, 2, 2, 2, 146, 153, 5, 4, 3, 2, 147, 150, 5, 30, 16, 2, 148,
	149, 9, 2, 2, 2, 149, 151, 5, 8, 5, 2, 150, 148, 3, 2, 2, 2, 150, 151,
	3, 2, 2, 2, 151, 153, 3, 2, 2, 2, 152, 146, 3, 2, 2, 2, 152, 147, 3, 2,
	2, 2, 153, 29, 3, 2, 2, 2, 154, 158, 5, 22, 12, 2, 155, 157, 5, 16, 9,
	2, 156, 155, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 31, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 166, 5,
	34, 18, 2, 162, 163, 7, 4, 2, 2, 163, 165, 5, 34, 18, 2, 164, 162, 3, 2,
	2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2,
	167, 33, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 174, 5, 36, 19, 2, 170,
	171, 7, 5, 2, 2, 171, 173, 5, 36, 19, 2, 172, 170, 3, 2, 2, 2, 173, 176,
	3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 35, 3, 2,
	2, 2, 176, 174, 3, 2, 2, 2, 177, 182, 5, 38, 20, 2, 178, 179, 9, 4, 2,
	2, 179, 181, 5, 38, 20, 2, 180, 178, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2,
	182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 37, 3, 2, 2, 2, 184, 182,
	3, 2, 2, 2, 185, 190, 5, 40, 21, 2, 186, 187, 9, 5, 2, 2, 187, 189, 5,
	40, 21, 2, 188, 186, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2,
	2, 2, 190, 191, 3, 2, 2, 2, 191, 39, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2,
	193, 198, 5, 42, 22, 2, 194, 195, 9, 6, 2, 2, 195, 197, 5, 42, 22, 2, 196,
	194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199,
	3, 2, 2, 2, 199, 41, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 204, 5, 44,
	23, 2, 202, 203, 9, 7, 2, 2, 203, 205, 5, 42, 22, 2, 204, 202, 3, 2, 2,
	2, 204, 205, 3, 2, 2, 2, 205, 212, 3, 2, 2, 2, 206, 209, 7, 14, 2, 2, 207,
	208, 9, 8, 2, 2, 208, 210, 5, 42, 22, 2, 209, 207, 3, 2, 2, 2, 209, 210,
	3, 2, 2, 2, 210, 212, 3, 2, 2, 2, 211, 201, 3, 2, 2, 2, 211, 206, 3, 2,
	2, 2, 212, 43, 3, 2, 2, 2, 213, 215, 7, 20, 2, 2, 214, 213, 3, 2, 2, 2,
	215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217,
	219, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 220, 5, 26, 14, 2, 220, 45,
	3, 2, 2, 2, 221, 224, 5, 54, 28, 2, 222, 223, 7, 32, 2, 2, 223, 225, 5,
	54, 28, 2, 224, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 47, 3, 2, 2,
	2, 226, 227, 5, 46, 24, 2, 227, 49, 3, 2, 2, 2, 228, 229, 7, 10, 2, 2,
	229, 230, 5, 46, 24, 2, 230, 51, 3, 2, 2, 2, 231, 238, 7, 23, 2, 2, 232,
	233, 5, 54, 28, 2, 233, 234, 7, 32, 2, 2, 234, 235, 7, 23, 2, 2, 235, 238,
	3, 2, 2, 2, 236, 238, 5, 46, 24, 2, 237, 231, 3, 2, 2, 2, 237, 232, 3,
	2, 2, 2, 237, 236, 3, 2, 2, 2, 238, 53, 3, 2, 2, 2, 239, 240, 9, 9, 2,
	2, 240, 55, 3, 2, 2, 2, 29, 60, 66, 73, 81, 85, 90, 92, 102, 120, 129,
	132, 139, 144, 150, 152, 158, 166, 174, 182, 190, 198, 204, 209, 211, 216,
	224, 237,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'processing-instruction'", "'or'", "'and'", "'='", "'!='", "'div'",
	"'mod'", "'$'", "", "", "", "'/'", "'//'", "'('", "')'", "'['", "']'",
	"'-'", "'+'", "'.'", "'*'", "'..'", "'@'", "','", "'|'", "'<'", "'>'",
	"'<='", "'>='", "':'", "'::'", "'''", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "NodeType", "Number", "AxisName", "PATHSEP",
	"ABRPATH", "LPAR", "RPAR", "LBRAC", "RBRAC", "MINUS", "PLUS", "DOT", "MUL",
	"DOTDOT", "AT", "COMMA", "PIPE", "LESS", "MORE_", "LE", "GE", "COLON",
	"CC", "APOS", "QUOT", "Literal", "Whitespace", "NCName",
}

var ruleNames = []string{
	"main", "locationPath", "absoluteLocationPathNoroot", "relativeLocationPath",
	"step", "axisSpecifier", "nodeTest", "predicate", "abbreviatedStep", "expr",
	"primaryExpr", "functionCall", "unionExprNoRoot", "pathExprNoRoot", "filterExpr",
	"orExpr", "andExpr", "equalityExpr", "relationalExpr", "additiveExpr",
	"multiplicativeExpr", "unaryExprNoRoot", "qName", "functionName", "variableReference",
	"nameTest", "nCName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type xpathParser struct {
	*antlr.BaseParser
}

func NewxpathParser(input antlr.TokenStream) *xpathParser {
	this := new(xpathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "xpath.g4"

	return this
}

// xpathParser tokens.
const (
	xpathParserEOF        = antlr.TokenEOF
	xpathParserT__0       = 1
	xpathParserT__1       = 2
	xpathParserT__2       = 3
	xpathParserT__3       = 4
	xpathParserT__4       = 5
	xpathParserT__5       = 6
	xpathParserT__6       = 7
	xpathParserT__7       = 8
	xpathParserNodeType   = 9
	xpathParserNumber     = 10
	xpathParserAxisName   = 11
	xpathParserPATHSEP    = 12
	xpathParserABRPATH    = 13
	xpathParserLPAR       = 14
	xpathParserRPAR       = 15
	xpathParserLBRAC      = 16
	xpathParserRBRAC      = 17
	xpathParserMINUS      = 18
	xpathParserPLUS       = 19
	xpathParserDOT        = 20
	xpathParserMUL        = 21
	xpathParserDOTDOT     = 22
	xpathParserAT         = 23
	xpathParserCOMMA      = 24
	xpathParserPIPE       = 25
	xpathParserLESS       = 26
	xpathParserMORE_      = 27
	xpathParserLE         = 28
	xpathParserGE         = 29
	xpathParserCOLON      = 30
	xpathParserCC         = 31
	xpathParserAPOS       = 32
	xpathParserQUOT       = 33
	xpathParserLiteral    = 34
	xpathParserWhitespace = 35
	xpathParserNCName     = 36
)

// xpathParser rules.
const (
	xpathParserRULE_main                       = 0
	xpathParserRULE_locationPath               = 1
	xpathParserRULE_absoluteLocationPathNoroot = 2
	xpathParserRULE_relativeLocationPath       = 3
	xpathParserRULE_step                       = 4
	xpathParserRULE_axisSpecifier              = 5
	xpathParserRULE_nodeTest                   = 6
	xpathParserRULE_predicate                  = 7
	xpathParserRULE_abbreviatedStep            = 8
	xpathParserRULE_expr                       = 9
	xpathParserRULE_primaryExpr                = 10
	xpathParserRULE_functionCall               = 11
	xpathParserRULE_unionExprNoRoot            = 12
	xpathParserRULE_pathExprNoRoot             = 13
	xpathParserRULE_filterExpr                 = 14
	xpathParserRULE_orExpr                     = 15
	xpathParserRULE_andExpr                    = 16
	xpathParserRULE_equalityExpr               = 17
	xpathParserRULE_relationalExpr             = 18
	xpathParserRULE_additiveExpr               = 19
	xpathParserRULE_multiplicativeExpr         = 20
	xpathParserRULE_unaryExprNoRoot            = 21
	xpathParserRULE_qName                      = 22
	xpathParserRULE_functionName               = 23
	xpathParserRULE_variableReference          = 24
	xpathParserRULE_nameTest                   = 25
	xpathParserRULE_nCName                     = 26
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *xpathParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, xpathParserRULE_main)

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
		p.SetState(54)
		p.Expr()
	}

	return localctx
}

// ILocationPathContext is an interface to support dynamic dispatch.
type ILocationPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationPathContext differentiates from other interfaces.
	IsLocationPathContext()
}

type LocationPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationPathContext() *LocationPathContext {
	var p = new(LocationPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_locationPath
	return p
}

func (*LocationPathContext) IsLocationPathContext() {}

func NewLocationPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationPathContext {
	var p = new(LocationPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_locationPath

	return p
}

func (s *LocationPathContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationPathContext) RelativeLocationPath() IRelativeLocationPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelativeLocationPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelativeLocationPathContext)
}

func (s *LocationPathContext) AbsoluteLocationPathNoroot() IAbsoluteLocationPathNorootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsoluteLocationPathNorootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsoluteLocationPathNorootContext)
}

func (s *LocationPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterLocationPath(s)
	}
}

func (s *LocationPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitLocationPath(s)
	}
}

func (p *xpathParser) LocationPath() (localctx ILocationPathContext) {
	localctx = NewLocationPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, xpathParserRULE_locationPath)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xpathParserT__0, xpathParserNodeType, xpathParserAxisName, xpathParserDOT, xpathParserMUL, xpathParserDOTDOT, xpathParserAT, xpathParserNCName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.RelativeLocationPath()
		}

	case xpathParserPATHSEP, xpathParserABRPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.AbsoluteLocationPathNoroot()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAbsoluteLocationPathNorootContext is an interface to support dynamic dispatch.
type IAbsoluteLocationPathNorootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsoluteLocationPathNorootContext differentiates from other interfaces.
	IsAbsoluteLocationPathNorootContext()
}

type AbsoluteLocationPathNorootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsoluteLocationPathNorootContext() *AbsoluteLocationPathNorootContext {
	var p = new(AbsoluteLocationPathNorootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_absoluteLocationPathNoroot
	return p
}

func (*AbsoluteLocationPathNorootContext) IsAbsoluteLocationPathNorootContext() {}

func NewAbsoluteLocationPathNorootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsoluteLocationPathNorootContext {
	var p = new(AbsoluteLocationPathNorootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_absoluteLocationPathNoroot

	return p
}

func (s *AbsoluteLocationPathNorootContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsoluteLocationPathNorootContext) RelativeLocationPath() IRelativeLocationPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelativeLocationPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelativeLocationPathContext)
}

func (s *AbsoluteLocationPathNorootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsoluteLocationPathNorootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsoluteLocationPathNorootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterAbsoluteLocationPathNoroot(s)
	}
}

func (s *AbsoluteLocationPathNorootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitAbsoluteLocationPathNoroot(s)
	}
}

func (p *xpathParser) AbsoluteLocationPathNoroot() (localctx IAbsoluteLocationPathNorootContext) {
	localctx = NewAbsoluteLocationPathNorootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, xpathParserRULE_absoluteLocationPathNoroot)

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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xpathParserPATHSEP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(xpathParserPATHSEP)
		}
		{
			p.SetState(61)
			p.RelativeLocationPath()
		}

	case xpathParserABRPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(xpathParserABRPATH)
		}
		{
			p.SetState(63)
			p.RelativeLocationPath()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelativeLocationPathContext is an interface to support dynamic dispatch.
type IRelativeLocationPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelativeLocationPathContext differentiates from other interfaces.
	IsRelativeLocationPathContext()
}

type RelativeLocationPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelativeLocationPathContext() *RelativeLocationPathContext {
	var p = new(RelativeLocationPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_relativeLocationPath
	return p
}

func (*RelativeLocationPathContext) IsRelativeLocationPathContext() {}

func NewRelativeLocationPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelativeLocationPathContext {
	var p = new(RelativeLocationPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_relativeLocationPath

	return p
}

func (s *RelativeLocationPathContext) GetParser() antlr.Parser { return s.parser }

func (s *RelativeLocationPathContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *RelativeLocationPathContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *RelativeLocationPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeLocationPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelativeLocationPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterRelativeLocationPath(s)
	}
}

func (s *RelativeLocationPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitRelativeLocationPath(s)
	}
}

func (p *xpathParser) RelativeLocationPath() (localctx IRelativeLocationPathContext) {
	localctx = NewRelativeLocationPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, xpathParserRULE_relativeLocationPath)
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
		p.SetState(66)
		p.Step()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserPATHSEP || _la == xpathParserABRPATH {
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(_la == xpathParserPATHSEP || _la == xpathParserABRPATH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(68)
			p.Step()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) AxisSpecifier() IAxisSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxisSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAxisSpecifierContext)
}

func (s *StepContext) NodeTest() INodeTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INodeTestContext)
}

func (s *StepContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *StepContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *StepContext) AbbreviatedStep() IAbbreviatedStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbbreviatedStepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbbreviatedStepContext)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterStep(s)
	}
}

func (s *StepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitStep(s)
	}
}

func (p *xpathParser) Step() (localctx IStepContext) {
	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, xpathParserRULE_step)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xpathParserT__0, xpathParserNodeType, xpathParserAxisName, xpathParserMUL, xpathParserAT, xpathParserNCName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.AxisSpecifier()
		}
		{
			p.SetState(75)
			p.NodeTest()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == xpathParserLBRAC {
			{
				p.SetState(76)
				p.Predicate()
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case xpathParserDOT, xpathParserDOTDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.AbbreviatedStep()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAxisSpecifierContext is an interface to support dynamic dispatch.
type IAxisSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxisSpecifierContext differentiates from other interfaces.
	IsAxisSpecifierContext()
}

type AxisSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxisSpecifierContext() *AxisSpecifierContext {
	var p = new(AxisSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_axisSpecifier
	return p
}

func (*AxisSpecifierContext) IsAxisSpecifierContext() {}

func NewAxisSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AxisSpecifierContext {
	var p = new(AxisSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_axisSpecifier

	return p
}

func (s *AxisSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AxisSpecifierContext) AxisName() antlr.TerminalNode {
	return s.GetToken(xpathParserAxisName, 0)
}

func (s *AxisSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxisSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AxisSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterAxisSpecifier(s)
	}
}

func (s *AxisSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitAxisSpecifier(s)
	}
}

func (p *xpathParser) AxisSpecifier() (localctx IAxisSpecifierContext) {
	localctx = NewAxisSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, xpathParserRULE_axisSpecifier)
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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(xpathParserAxisName)
		}
		{
			p.SetState(86)
			p.Match(xpathParserCC)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xpathParserAT {
			{
				p.SetState(87)
				p.Match(xpathParserAT)
			}

		}

	}

	return localctx
}

// INodeTestContext is an interface to support dynamic dispatch.
type INodeTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeTestContext differentiates from other interfaces.
	IsNodeTestContext()
}

type NodeTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeTestContext() *NodeTestContext {
	var p = new(NodeTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_nodeTest
	return p
}

func (*NodeTestContext) IsNodeTestContext() {}

func NewNodeTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeTestContext {
	var p = new(NodeTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_nodeTest

	return p
}

func (s *NodeTestContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeTestContext) NameTest() INameTestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameTestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameTestContext)
}

func (s *NodeTestContext) NodeType() antlr.TerminalNode {
	return s.GetToken(xpathParserNodeType, 0)
}

func (s *NodeTestContext) Literal() antlr.TerminalNode {
	return s.GetToken(xpathParserLiteral, 0)
}

func (s *NodeTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterNodeTest(s)
	}
}

func (s *NodeTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitNodeTest(s)
	}
}

func (p *xpathParser) NodeTest() (localctx INodeTestContext) {
	localctx = NewNodeTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, xpathParserRULE_nodeTest)

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

	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xpathParserAxisName, xpathParserMUL, xpathParserNCName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.NameTest()
		}

	case xpathParserNodeType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(xpathParserNodeType)
		}
		{
			p.SetState(94)
			p.Match(xpathParserLPAR)
		}
		{
			p.SetState(95)
			p.Match(xpathParserRPAR)
		}

	case xpathParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(xpathParserT__0)
		}
		{
			p.SetState(97)
			p.Match(xpathParserLPAR)
		}
		{
			p.SetState(98)
			p.Match(xpathParserLiteral)
		}
		{
			p.SetState(99)
			p.Match(xpathParserRPAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *xpathParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, xpathParserRULE_predicate)

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
		p.SetState(102)
		p.Match(xpathParserLBRAC)
	}
	{
		p.SetState(103)
		p.Expr()
	}
	{
		p.SetState(104)
		p.Match(xpathParserRBRAC)
	}

	return localctx
}

// IAbbreviatedStepContext is an interface to support dynamic dispatch.
type IAbbreviatedStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbbreviatedStepContext differentiates from other interfaces.
	IsAbbreviatedStepContext()
}

type AbbreviatedStepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbbreviatedStepContext() *AbbreviatedStepContext {
	var p = new(AbbreviatedStepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_abbreviatedStep
	return p
}

func (*AbbreviatedStepContext) IsAbbreviatedStepContext() {}

func NewAbbreviatedStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbbreviatedStepContext {
	var p = new(AbbreviatedStepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_abbreviatedStep

	return p
}

func (s *AbbreviatedStepContext) GetParser() antlr.Parser { return s.parser }
func (s *AbbreviatedStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbbreviatedStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbbreviatedStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterAbbreviatedStep(s)
	}
}

func (s *AbbreviatedStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitAbbreviatedStep(s)
	}
}

func (p *xpathParser) AbbreviatedStep() (localctx IAbbreviatedStepContext) {
	localctx = NewAbbreviatedStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, xpathParserRULE_abbreviatedStep)
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
	p.SetState(106)
	_la = p.GetTokenStream().LA(1)

	if !(_la == xpathParserDOT || _la == xpathParserDOTDOT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) OrExpr() IOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *xpathParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, xpathParserRULE_expr)

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
		p.SetState(108)
		p.OrExpr()
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) VariableReference() IVariableReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableReferenceContext)
}

func (s *PrimaryExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryExprContext) Literal() antlr.TerminalNode {
	return s.GetToken(xpathParserLiteral, 0)
}

func (s *PrimaryExprContext) Number() antlr.TerminalNode {
	return s.GetToken(xpathParserNumber, 0)
}

func (s *PrimaryExprContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (p *xpathParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, xpathParserRULE_primaryExpr)

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xpathParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.VariableReference()
		}

	case xpathParserLPAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Match(xpathParserLPAR)
		}
		{
			p.SetState(112)
			p.Expr()
		}
		{
			p.SetState(113)
			p.Match(xpathParserRPAR)
		}

	case xpathParserLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(xpathParserLiteral)
		}

	case xpathParserNumber:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Match(xpathParserNumber)
		}

	case xpathParserAxisName, xpathParserNCName:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
			p.FunctionCall()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *xpathParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, xpathParserRULE_functionCall)
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
		p.SetState(120)
		p.FunctionName()
	}
	{
		p.SetState(121)
		p.Match(xpathParserLPAR)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xpathParserT__0)|(1<<xpathParserT__7)|(1<<xpathParserNodeType)|(1<<xpathParserNumber)|(1<<xpathParserAxisName)|(1<<xpathParserPATHSEP)|(1<<xpathParserABRPATH)|(1<<xpathParserLPAR)|(1<<xpathParserMINUS)|(1<<xpathParserDOT)|(1<<xpathParserMUL)|(1<<xpathParserDOTDOT)|(1<<xpathParserAT))) != 0) || _la == xpathParserLiteral || _la == xpathParserNCName {
		{
			p.SetState(122)
			p.Expr()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == xpathParserCOMMA {
			{
				p.SetState(123)
				p.Match(xpathParserCOMMA)
			}
			{
				p.SetState(124)
				p.Expr()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(132)
		p.Match(xpathParserRPAR)
	}

	return localctx
}

// IUnionExprNoRootContext is an interface to support dynamic dispatch.
type IUnionExprNoRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionExprNoRootContext differentiates from other interfaces.
	IsUnionExprNoRootContext()
}

type UnionExprNoRootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionExprNoRootContext() *UnionExprNoRootContext {
	var p = new(UnionExprNoRootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_unionExprNoRoot
	return p
}

func (*UnionExprNoRootContext) IsUnionExprNoRootContext() {}

func NewUnionExprNoRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionExprNoRootContext {
	var p = new(UnionExprNoRootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_unionExprNoRoot

	return p
}

func (s *UnionExprNoRootContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionExprNoRootContext) PathExprNoRoot() IPathExprNoRootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathExprNoRootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathExprNoRootContext)
}

func (s *UnionExprNoRootContext) UnionExprNoRoot() IUnionExprNoRootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionExprNoRootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionExprNoRootContext)
}

func (s *UnionExprNoRootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionExprNoRootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionExprNoRootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterUnionExprNoRoot(s)
	}
}

func (s *UnionExprNoRootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitUnionExprNoRoot(s)
	}
}

func (p *xpathParser) UnionExprNoRoot() (localctx IUnionExprNoRootContext) {
	localctx = NewUnionExprNoRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, xpathParserRULE_unionExprNoRoot)
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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.PathExprNoRoot()
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xpathParserPIPE {
			{
				p.SetState(135)
				p.Match(xpathParserPIPE)
			}
			{
				p.SetState(136)
				p.UnionExprNoRoot()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(xpathParserPATHSEP)
		}
		{
			p.SetState(140)
			p.Match(xpathParserPIPE)
		}
		{
			p.SetState(141)
			p.UnionExprNoRoot()
		}

	}

	return localctx
}

// IPathExprNoRootContext is an interface to support dynamic dispatch.
type IPathExprNoRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathExprNoRootContext differentiates from other interfaces.
	IsPathExprNoRootContext()
}

type PathExprNoRootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathExprNoRootContext() *PathExprNoRootContext {
	var p = new(PathExprNoRootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_pathExprNoRoot
	return p
}

func (*PathExprNoRootContext) IsPathExprNoRootContext() {}

func NewPathExprNoRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathExprNoRootContext {
	var p = new(PathExprNoRootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_pathExprNoRoot

	return p
}

func (s *PathExprNoRootContext) GetParser() antlr.Parser { return s.parser }

func (s *PathExprNoRootContext) LocationPath() ILocationPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationPathContext)
}

func (s *PathExprNoRootContext) FilterExpr() IFilterExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *PathExprNoRootContext) RelativeLocationPath() IRelativeLocationPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelativeLocationPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelativeLocationPathContext)
}

func (s *PathExprNoRootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathExprNoRootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathExprNoRootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterPathExprNoRoot(s)
	}
}

func (s *PathExprNoRootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitPathExprNoRoot(s)
	}
}

func (p *xpathParser) PathExprNoRoot() (localctx IPathExprNoRootContext) {
	localctx = NewPathExprNoRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, xpathParserRULE_pathExprNoRoot)
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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.LocationPath()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.FilterExpr()
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xpathParserPATHSEP || _la == xpathParserABRPATH {
			p.SetState(146)
			_la = p.GetTokenStream().LA(1)

			if !(_la == xpathParserPATHSEP || _la == xpathParserABRPATH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(147)
				p.RelativeLocationPath()
			}

		}

	}

	return localctx
}

// IFilterExprContext is an interface to support dynamic dispatch.
type IFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterExprContext differentiates from other interfaces.
	IsFilterExprContext()
}

type FilterExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExprContext() *FilterExprContext {
	var p = new(FilterExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_filterExpr
	return p
}

func (*FilterExprContext) IsFilterExprContext() {}

func NewFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExprContext {
	var p = new(FilterExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_filterExpr

	return p
}

func (s *FilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *FilterExprContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *FilterExprContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}

func (p *xpathParser) FilterExpr() (localctx IFilterExprContext) {
	localctx = NewFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, xpathParserRULE_filterExpr)
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
		p.SetState(152)
		p.PrimaryExpr()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserLBRAC {
		{
			p.SetState(153)
			p.Predicate()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_orExpr
	return p
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) AllAndExpr() []IAndExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndExprContext)(nil)).Elem())
	var tst = make([]IAndExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) AndExpr(i int) IAndExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (p *xpathParser) OrExpr() (localctx IOrExprContext) {
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, xpathParserRULE_orExpr)
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
		p.SetState(159)
		p.AndExpr()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserT__1 {
		{
			p.SetState(160)
			p.Match(xpathParserT__1)
		}
		{
			p.SetState(161)
			p.AndExpr()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_andExpr
	return p
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) AllEqualityExpr() []IEqualityExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEqualityExprContext)(nil)).Elem())
	var tst = make([]IEqualityExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEqualityExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) EqualityExpr(i int) IEqualityExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEqualityExprContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (p *xpathParser) AndExpr() (localctx IAndExprContext) {
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, xpathParserRULE_andExpr)
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
		p.SetState(167)
		p.EqualityExpr()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserT__2 {
		{
			p.SetState(168)
			p.Match(xpathParserT__2)
		}
		{
			p.SetState(169)
			p.EqualityExpr()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_equalityExpr
	return p
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) AllRelationalExpr() []IRelationalExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationalExprContext)(nil)).Elem())
	var tst = make([]IRelationalExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationalExprContext)
		}
	}

	return tst
}

func (s *EqualityExprContext) RelationalExpr(i int) IRelationalExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationalExprContext)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (p *xpathParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, xpathParserRULE_equalityExpr)
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
		p.SetState(175)
		p.RelationalExpr()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserT__3 || _la == xpathParserT__4 {
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !(_la == xpathParserT__3 || _la == xpathParserT__4) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(177)
			p.RelationalExpr()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationalExprContext is an interface to support dynamic dispatch.
type IRelationalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExprContext differentiates from other interfaces.
	IsRelationalExprContext()
}

type RelationalExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExprContext() *RelationalExprContext {
	var p = new(RelationalExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_relationalExpr
	return p
}

func (*RelationalExprContext) IsRelationalExprContext() {}

func NewRelationalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExprContext {
	var p = new(RelationalExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_relationalExpr

	return p
}

func (s *RelationalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditiveExprContext)(nil)).Elem())
	var tst = make([]IAdditiveExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditiveExprContext)
		}
	}

	return tst
}

func (s *RelationalExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (p *xpathParser) RelationalExpr() (localctx IRelationalExprContext) {
	localctx = NewRelationalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, xpathParserRULE_relationalExpr)
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
		p.SetState(183)
		p.AdditiveExpr()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xpathParserLESS)|(1<<xpathParserMORE_)|(1<<xpathParserLE)|(1<<xpathParserGE))) != 0 {
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xpathParserLESS)|(1<<xpathParserMORE_)|(1<<xpathParserLE)|(1<<xpathParserGE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(185)
			p.AdditiveExpr()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_additiveExpr
	return p
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicativeExprContext)(nil)).Elem())
	var tst = make([]IMultiplicativeExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicativeExprContext)
		}
	}

	return tst
}

func (s *AdditiveExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExprContext)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (p *xpathParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, xpathParserRULE_additiveExpr)
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
		p.SetState(191)
		p.MultiplicativeExpr()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserMINUS || _la == xpathParserPLUS {
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == xpathParserMINUS || _la == xpathParserPLUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(193)
			p.MultiplicativeExpr()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_multiplicativeExpr
	return p
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) UnaryExprNoRoot() IUnaryExprNoRootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprNoRootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprNoRootContext)
}

func (s *MultiplicativeExprContext) MultiplicativeExpr() IMultiplicativeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExprContext)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (p *xpathParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, xpathParserRULE_multiplicativeExpr)
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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.UnaryExprNoRoot()
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xpathParserT__5)|(1<<xpathParserT__6)|(1<<xpathParserMUL))) != 0 {
			p.SetState(200)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xpathParserT__5)|(1<<xpathParserT__6)|(1<<xpathParserMUL))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(201)
				p.MultiplicativeExpr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Match(xpathParserPATHSEP)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xpathParserT__5 || _la == xpathParserT__6 {
			p.SetState(205)
			_la = p.GetTokenStream().LA(1)

			if !(_la == xpathParserT__5 || _la == xpathParserT__6) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(206)
				p.MultiplicativeExpr()
			}

		}

	}

	return localctx
}

// IUnaryExprNoRootContext is an interface to support dynamic dispatch.
type IUnaryExprNoRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprNoRootContext differentiates from other interfaces.
	IsUnaryExprNoRootContext()
}

type UnaryExprNoRootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprNoRootContext() *UnaryExprNoRootContext {
	var p = new(UnaryExprNoRootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_unaryExprNoRoot
	return p
}

func (*UnaryExprNoRootContext) IsUnaryExprNoRootContext() {}

func NewUnaryExprNoRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprNoRootContext {
	var p = new(UnaryExprNoRootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_unaryExprNoRoot

	return p
}

func (s *UnaryExprNoRootContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprNoRootContext) UnionExprNoRoot() IUnionExprNoRootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionExprNoRootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionExprNoRootContext)
}

func (s *UnaryExprNoRootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprNoRootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprNoRootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterUnaryExprNoRoot(s)
	}
}

func (s *UnaryExprNoRootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitUnaryExprNoRoot(s)
	}
}

func (p *xpathParser) UnaryExprNoRoot() (localctx IUnaryExprNoRootContext) {
	localctx = NewUnaryExprNoRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, xpathParserRULE_unaryExprNoRoot)
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
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xpathParserMINUS {
		{
			p.SetState(211)
			p.Match(xpathParserMINUS)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
		p.UnionExprNoRoot()
	}

	return localctx
}

// IQNameContext is an interface to support dynamic dispatch.
type IQNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQNameContext differentiates from other interfaces.
	IsQNameContext()
}

type QNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQNameContext() *QNameContext {
	var p = new(QNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_qName
	return p
}

func (*QNameContext) IsQNameContext() {}

func NewQNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QNameContext {
	var p = new(QNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_qName

	return p
}

func (s *QNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QNameContext) AllNCName() []INCNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INCNameContext)(nil)).Elem())
	var tst = make([]INCNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INCNameContext)
		}
	}

	return tst
}

func (s *QNameContext) NCName(i int) INCNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INCNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INCNameContext)
}

func (s *QNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterQName(s)
	}
}

func (s *QNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitQName(s)
	}
}

func (p *xpathParser) QName() (localctx IQNameContext) {
	localctx = NewQNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, xpathParserRULE_qName)
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
		p.SetState(219)
		p.NCName()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == xpathParserCOLON {
		{
			p.SetState(220)
			p.Match(xpathParserCOLON)
		}
		{
			p.SetState(221)
			p.NCName()
		}

	}

	return localctx
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) QName() IQNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQNameContext)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (p *xpathParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, xpathParserRULE_functionName)

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
		p.SetState(224)
		p.QName()
	}

	return localctx
}

// IVariableReferenceContext is an interface to support dynamic dispatch.
type IVariableReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableReferenceContext differentiates from other interfaces.
	IsVariableReferenceContext()
}

type VariableReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableReferenceContext() *VariableReferenceContext {
	var p = new(VariableReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_variableReference
	return p
}

func (*VariableReferenceContext) IsVariableReferenceContext() {}

func NewVariableReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableReferenceContext {
	var p = new(VariableReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_variableReference

	return p
}

func (s *VariableReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableReferenceContext) QName() IQNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQNameContext)
}

func (s *VariableReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterVariableReference(s)
	}
}

func (s *VariableReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitVariableReference(s)
	}
}

func (p *xpathParser) VariableReference() (localctx IVariableReferenceContext) {
	localctx = NewVariableReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, xpathParserRULE_variableReference)

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
		p.SetState(226)
		p.Match(xpathParserT__7)
	}
	{
		p.SetState(227)
		p.QName()
	}

	return localctx
}

// INameTestContext is an interface to support dynamic dispatch.
type INameTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameTestContext differentiates from other interfaces.
	IsNameTestContext()
}

type NameTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameTestContext() *NameTestContext {
	var p = new(NameTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_nameTest
	return p
}

func (*NameTestContext) IsNameTestContext() {}

func NewNameTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameTestContext {
	var p = new(NameTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_nameTest

	return p
}

func (s *NameTestContext) GetParser() antlr.Parser { return s.parser }

func (s *NameTestContext) NCName() INCNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INCNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INCNameContext)
}

func (s *NameTestContext) QName() IQNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQNameContext)
}

func (s *NameTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterNameTest(s)
	}
}

func (s *NameTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitNameTest(s)
	}
}

func (p *xpathParser) NameTest() (localctx INameTestContext) {
	localctx = NewNameTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, xpathParserRULE_nameTest)

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

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(xpathParserMUL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.NCName()
		}
		{
			p.SetState(231)
			p.Match(xpathParserCOLON)
		}
		{
			p.SetState(232)
			p.Match(xpathParserMUL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.QName()
		}

	}

	return localctx
}

// INCNameContext is an interface to support dynamic dispatch.
type INCNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNCNameContext differentiates from other interfaces.
	IsNCNameContext()
}

type NCNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNCNameContext() *NCNameContext {
	var p = new(NCNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xpathParserRULE_nCName
	return p
}

func (*NCNameContext) IsNCNameContext() {}

func NewNCNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NCNameContext {
	var p = new(NCNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xpathParserRULE_nCName

	return p
}

func (s *NCNameContext) GetParser() antlr.Parser { return s.parser }

func (s *NCNameContext) NCName() antlr.TerminalNode {
	return s.GetToken(xpathParserNCName, 0)
}

func (s *NCNameContext) AxisName() antlr.TerminalNode {
	return s.GetToken(xpathParserAxisName, 0)
}

func (s *NCNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NCNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NCNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.EnterNCName(s)
	}
}

func (s *NCNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xpathListener); ok {
		listenerT.ExitNCName(s)
	}
}

func (p *xpathParser) NCName() (localctx INCNameContext) {
	localctx = NewNCNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, xpathParserRULE_nCName)
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
	p.SetState(237)
	_la = p.GetTokenStream().LA(1)

	if !(_la == xpathParserAxisName || _la == xpathParserNCName) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
