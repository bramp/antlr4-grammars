// Generated from oncrpcv2.g4 by ANTLR 4.7.

package oncrpc // oncrpcv2
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 286,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 7, 2, 50, 10, 2, 12, 2, 14, 2, 53, 11, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 65, 10, 3, 12, 3,
	14, 3, 68, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 81, 10, 4, 12, 4, 14, 4, 84, 11, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 93, 10, 5, 3, 6, 3, 6, 5, 6, 97, 10, 6, 3,
	7, 3, 7, 7, 7, 101, 10, 7, 12, 7, 14, 7, 104, 11, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 119,
	10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 133, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 140, 10, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 148, 10, 8, 3, 9, 3, 9, 5, 9,
	152, 10, 9, 3, 10, 3, 10, 3, 11, 5, 11, 157, 10, 11, 3, 11, 3, 11, 5, 11,
	161, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 172, 10, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 186, 10, 13, 12, 13, 14, 13,
	189, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 7, 15, 203, 10, 15, 12, 15, 14, 15, 206, 11, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 7, 17, 220, 10, 17, 12, 17, 14, 17, 223, 11, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 230, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 242, 10, 18, 12, 18, 14,
	18, 245, 11, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	275, 10, 20, 3, 21, 3, 21, 5, 21, 279, 10, 21, 3, 22, 6, 22, 282, 10, 22,
	13, 22, 14, 22, 283, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 3, 3, 2, 37, 39, 2, 303,
	2, 44, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 92, 3, 2, 2,
	2, 10, 96, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2, 14, 147, 3, 2, 2, 2, 16, 151,
	3, 2, 2, 2, 18, 153, 3, 2, 2, 2, 20, 171, 3, 2, 2, 2, 22, 173, 3, 2, 2,
	2, 24, 176, 3, 2, 2, 2, 26, 192, 3, 2, 2, 2, 28, 195, 3, 2, 2, 2, 30, 209,
	3, 2, 2, 2, 32, 212, 3, 2, 2, 2, 34, 233, 3, 2, 2, 2, 36, 249, 3, 2, 2,
	2, 38, 274, 3, 2, 2, 2, 40, 278, 3, 2, 2, 2, 42, 281, 3, 2, 2, 2, 44, 45,
	7, 3, 2, 2, 45, 46, 7, 40, 2, 2, 46, 47, 7, 4, 2, 2, 47, 51, 5, 4, 3, 2,
	48, 50, 5, 4, 3, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	55, 7, 5, 2, 2, 55, 56, 7, 6, 2, 2, 56, 57, 5, 18, 10, 2, 57, 58, 7, 7,
	2, 2, 58, 3, 3, 2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 61, 7, 40, 2, 2, 61, 62,
	7, 4, 2, 2, 62, 66, 5, 6, 4, 2, 63, 65, 5, 6, 4, 2, 64, 63, 3, 2, 2, 2,
	65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3,
	2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 5, 2, 2, 70, 71, 7, 6, 2, 2, 71,
	72, 5, 18, 10, 2, 72, 73, 7, 7, 2, 2, 73, 5, 3, 2, 2, 2, 74, 75, 5, 8,
	5, 2, 75, 76, 7, 40, 2, 2, 76, 77, 7, 9, 2, 2, 77, 82, 5, 10, 6, 2, 78,
	79, 7, 10, 2, 2, 79, 81, 5, 20, 11, 2, 80, 78, 3, 2, 2, 2, 81, 84, 3, 2,
	2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82,
	3, 2, 2, 2, 85, 86, 7, 11, 2, 2, 86, 87, 7, 6, 2, 2, 87, 88, 5, 18, 10,
	2, 88, 89, 7, 7, 2, 2, 89, 7, 3, 2, 2, 2, 90, 93, 7, 12, 2, 2, 91, 93,
	5, 20, 11, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 9, 3, 2, 2, 2,
	94, 97, 7, 12, 2, 2, 95, 97, 5, 20, 11, 2, 96, 94, 3, 2, 2, 2, 96, 95,
	3, 2, 2, 2, 97, 11, 3, 2, 2, 2, 98, 101, 5, 42, 22, 2, 99, 101, 5, 2, 2,
	2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102,
	100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 13, 3, 2, 2, 2, 104, 102, 3,
	2, 2, 2, 105, 106, 5, 20, 11, 2, 106, 107, 7, 40, 2, 2, 107, 148, 3, 2,
	2, 2, 108, 109, 5, 20, 11, 2, 109, 110, 7, 40, 2, 2, 110, 111, 7, 13, 2,
	2, 111, 112, 5, 16, 9, 2, 112, 113, 7, 14, 2, 2, 113, 148, 3, 2, 2, 2,
	114, 115, 5, 20, 11, 2, 115, 116, 7, 40, 2, 2, 116, 118, 7, 15, 2, 2, 117,
	119, 5, 16, 9, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120,
	3, 2, 2, 2, 120, 121, 7, 16, 2, 2, 121, 148, 3, 2, 2, 2, 122, 123, 7, 17,
	2, 2, 123, 124, 7, 40, 2, 2, 124, 125, 7, 13, 2, 2, 125, 126, 5, 16, 9,
	2, 126, 127, 7, 14, 2, 2, 127, 148, 3, 2, 2, 2, 128, 129, 7, 17, 2, 2,
	129, 130, 7, 40, 2, 2, 130, 132, 7, 15, 2, 2, 131, 133, 5, 16, 9, 2, 132,
	131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 148,
	7, 16, 2, 2, 135, 136, 7, 18, 2, 2, 136, 137, 7, 40, 2, 2, 137, 139, 7,
	15, 2, 2, 138, 140, 5, 16, 9, 2, 139, 138, 3, 2, 2, 2, 139, 140, 3, 2,
	2, 2, 140, 141, 3, 2, 2, 2, 141, 148, 7, 16, 2, 2, 142, 143, 5, 20, 11,
	2, 143, 144, 7, 19, 2, 2, 144, 145, 7, 40, 2, 2, 145, 148, 3, 2, 2, 2,
	146, 148, 7, 12, 2, 2, 147, 105, 3, 2, 2, 2, 147, 108, 3, 2, 2, 2, 147,
	114, 3, 2, 2, 2, 147, 122, 3, 2, 2, 2, 147, 128, 3, 2, 2, 2, 147, 135,
	3, 2, 2, 2, 147, 142, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 15, 3, 2,
	2, 2, 149, 152, 5, 18, 10, 2, 150, 152, 7, 40, 2, 2, 151, 149, 3, 2, 2,
	2, 151, 150, 3, 2, 2, 2, 152, 17, 3, 2, 2, 2, 153, 154, 9, 2, 2, 2, 154,
	19, 3, 2, 2, 2, 155, 157, 7, 20, 2, 2, 156, 155, 3, 2, 2, 2, 156, 157,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 172, 7, 21, 2, 2, 159, 161, 7, 20,
	2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2,
	162, 172, 7, 22, 2, 2, 163, 172, 7, 23, 2, 2, 164, 172, 7, 24, 2, 2, 165,
	172, 7, 25, 2, 2, 166, 172, 7, 26, 2, 2, 167, 172, 5, 22, 12, 2, 168, 172,
	5, 26, 14, 2, 169, 172, 5, 30, 16, 2, 170, 172, 7, 40, 2, 2, 171, 156,
	3, 2, 2, 2, 171, 160, 3, 2, 2, 2, 171, 163, 3, 2, 2, 2, 171, 164, 3, 2,
	2, 2, 171, 165, 3, 2, 2, 2, 171, 166, 3, 2, 2, 2, 171, 167, 3, 2, 2, 2,
	171, 168, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172,
	21, 3, 2, 2, 2, 173, 174, 7, 27, 2, 2, 174, 175, 5, 24, 13, 2, 175, 23,
	3, 2, 2, 2, 176, 177, 7, 4, 2, 2, 177, 178, 7, 40, 2, 2, 178, 179, 7, 6,
	2, 2, 179, 180, 5, 16, 9, 2, 180, 187, 3, 2, 2, 2, 181, 182, 7, 10, 2,
	2, 182, 183, 7, 40, 2, 2, 183, 184, 7, 6, 2, 2, 184, 186, 5, 16, 9, 2,
	185, 181, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191,
	7, 5, 2, 2, 191, 25, 3, 2, 2, 2, 192, 193, 7, 28, 2, 2, 193, 194, 5, 28,
	15, 2, 194, 27, 3, 2, 2, 2, 195, 196, 7, 4, 2, 2, 196, 197, 5, 14, 8, 2,
	197, 198, 7, 7, 2, 2, 198, 204, 3, 2, 2, 2, 199, 200, 5, 14, 8, 2, 200,
	201, 7, 7, 2, 2, 201, 203, 3, 2, 2, 2, 202, 199, 3, 2, 2, 2, 203, 206,
	3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 207, 3, 2,
	2, 2, 206, 204, 3, 2, 2, 2, 207, 208, 7, 5, 2, 2, 208, 29, 3, 2, 2, 2,
	209, 210, 7, 29, 2, 2, 210, 211, 5, 32, 17, 2, 211, 31, 3, 2, 2, 2, 212,
	213, 7, 30, 2, 2, 213, 214, 7, 9, 2, 2, 214, 215, 5, 14, 8, 2, 215, 216,
	7, 11, 2, 2, 216, 217, 7, 4, 2, 2, 217, 221, 5, 34, 18, 2, 218, 220, 5,
	34, 18, 2, 219, 218, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2,
	2, 2, 221, 222, 3, 2, 2, 2, 222, 229, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2,
	224, 225, 7, 31, 2, 2, 225, 226, 7, 32, 2, 2, 226, 227, 5, 14, 8, 2, 227,
	228, 7, 7, 2, 2, 228, 230, 3, 2, 2, 2, 229, 224, 3, 2, 2, 2, 229, 230,
	3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 7, 5, 2, 2, 232, 33, 3, 2,
	2, 2, 233, 234, 7, 33, 2, 2, 234, 235, 5, 16, 9, 2, 235, 236, 7, 32, 2,
	2, 236, 243, 3, 2, 2, 2, 237, 238, 7, 33, 2, 2, 238, 239, 5, 16, 9, 2,
	239, 240, 7, 32, 2, 2, 240, 242, 3, 2, 2, 2, 241, 237, 3, 2, 2, 2, 242,
	245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 246,
	3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 247, 5, 14, 8, 2, 247, 248, 7, 7,
	2, 2, 248, 35, 3, 2, 2, 2, 249, 250, 7, 34, 2, 2, 250, 251, 7, 40, 2, 2,
	251, 252, 7, 6, 2, 2, 252, 253, 5, 18, 10, 2, 253, 254, 7, 7, 2, 2, 254,
	37, 3, 2, 2, 2, 255, 256, 7, 35, 2, 2, 256, 257, 5, 14, 8, 2, 257, 258,
	7, 7, 2, 2, 258, 275, 3, 2, 2, 2, 259, 260, 7, 27, 2, 2, 260, 261, 7, 40,
	2, 2, 261, 262, 5, 24, 13, 2, 262, 263, 7, 7, 2, 2, 263, 275, 3, 2, 2,
	2, 264, 265, 7, 28, 2, 2, 265, 266, 7, 40, 2, 2, 266, 267, 5, 28, 15, 2,
	267, 268, 7, 7, 2, 2, 268, 275, 3, 2, 2, 2, 269, 270, 7, 29, 2, 2, 270,
	271, 7, 40, 2, 2, 271, 272, 5, 32, 17, 2, 272, 273, 7, 7, 2, 2, 273, 275,
	3, 2, 2, 2, 274, 255, 3, 2, 2, 2, 274, 259, 3, 2, 2, 2, 274, 264, 3, 2,
	2, 2, 274, 269, 3, 2, 2, 2, 275, 39, 3, 2, 2, 2, 276, 279, 5, 38, 20, 2,
	277, 279, 5, 36, 19, 2, 278, 276, 3, 2, 2, 2, 278, 277, 3, 2, 2, 2, 279,
	41, 3, 2, 2, 2, 280, 282, 5, 40, 21, 2, 281, 280, 3, 2, 2, 2, 282, 283,
	3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 43, 3, 2,
	2, 2, 25, 51, 66, 82, 92, 96, 100, 102, 118, 132, 139, 147, 151, 156, 160,
	171, 187, 204, 221, 229, 243, 274, 278, 283,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'program'", "'{'", "'}'", "'='", "';'", "'version'", "'('", "','",
	"')'", "'void'", "'['", "']'", "'<'", "'>'", "'opaque'", "'string'", "'*'",
	"'unsigned'", "'int'", "'hyper'", "'float'", "'double'", "'quadruple'",
	"'bool'", "'enum'", "'struct'", "'union'", "'switch'", "'default'", "':'",
	"'case'", "'const'", "'typedef'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT",
	"OCTAL", "DECIMAL", "HEXADECIMAL", "IDENTIFIER", "WS",
}

var ruleNames = []string{
	"programDef", "versionDef", "procedureDef", "procReturn", "procFirstArg",
	"oncrpcv2Specification", "declaration", "value", "constant", "typeSpecifier",
	"enumTypeSpec", "enumBody", "structTypeSpec", "structBody", "unionTypeSpec",
	"unionBody", "caseSpec", "constantDef", "typeDef", "definition", "xdrSpecification",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type oncrpcv2Parser struct {
	*antlr.BaseParser
}

func Newoncrpcv2Parser(input antlr.TokenStream) *oncrpcv2Parser {
	this := new(oncrpcv2Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "oncrpcv2.g4"

	return this
}

// oncrpcv2Parser tokens.
const (
	oncrpcv2ParserEOF         = antlr.TokenEOF
	oncrpcv2ParserT__0        = 1
	oncrpcv2ParserT__1        = 2
	oncrpcv2ParserT__2        = 3
	oncrpcv2ParserT__3        = 4
	oncrpcv2ParserT__4        = 5
	oncrpcv2ParserT__5        = 6
	oncrpcv2ParserT__6        = 7
	oncrpcv2ParserT__7        = 8
	oncrpcv2ParserT__8        = 9
	oncrpcv2ParserT__9        = 10
	oncrpcv2ParserT__10       = 11
	oncrpcv2ParserT__11       = 12
	oncrpcv2ParserT__12       = 13
	oncrpcv2ParserT__13       = 14
	oncrpcv2ParserT__14       = 15
	oncrpcv2ParserT__15       = 16
	oncrpcv2ParserT__16       = 17
	oncrpcv2ParserT__17       = 18
	oncrpcv2ParserT__18       = 19
	oncrpcv2ParserT__19       = 20
	oncrpcv2ParserT__20       = 21
	oncrpcv2ParserT__21       = 22
	oncrpcv2ParserT__22       = 23
	oncrpcv2ParserT__23       = 24
	oncrpcv2ParserT__24       = 25
	oncrpcv2ParserT__25       = 26
	oncrpcv2ParserT__26       = 27
	oncrpcv2ParserT__27       = 28
	oncrpcv2ParserT__28       = 29
	oncrpcv2ParserT__29       = 30
	oncrpcv2ParserT__30       = 31
	oncrpcv2ParserT__31       = 32
	oncrpcv2ParserT__32       = 33
	oncrpcv2ParserCOMMENT     = 34
	oncrpcv2ParserOCTAL       = 35
	oncrpcv2ParserDECIMAL     = 36
	oncrpcv2ParserHEXADECIMAL = 37
	oncrpcv2ParserIDENTIFIER  = 38
	oncrpcv2ParserWS          = 39
)

// oncrpcv2Parser rules.
const (
	oncrpcv2ParserRULE_programDef            = 0
	oncrpcv2ParserRULE_versionDef            = 1
	oncrpcv2ParserRULE_procedureDef          = 2
	oncrpcv2ParserRULE_procReturn            = 3
	oncrpcv2ParserRULE_procFirstArg          = 4
	oncrpcv2ParserRULE_oncrpcv2Specification = 5
	oncrpcv2ParserRULE_declaration           = 6
	oncrpcv2ParserRULE_value                 = 7
	oncrpcv2ParserRULE_constant              = 8
	oncrpcv2ParserRULE_typeSpecifier         = 9
	oncrpcv2ParserRULE_enumTypeSpec          = 10
	oncrpcv2ParserRULE_enumBody              = 11
	oncrpcv2ParserRULE_structTypeSpec        = 12
	oncrpcv2ParserRULE_structBody            = 13
	oncrpcv2ParserRULE_unionTypeSpec         = 14
	oncrpcv2ParserRULE_unionBody             = 15
	oncrpcv2ParserRULE_caseSpec              = 16
	oncrpcv2ParserRULE_constantDef           = 17
	oncrpcv2ParserRULE_typeDef               = 18
	oncrpcv2ParserRULE_definition            = 19
	oncrpcv2ParserRULE_xdrSpecification      = 20
)

// IProgramDefContext is an interface to support dynamic dispatch.
type IProgramDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramDefContext differentiates from other interfaces.
	IsProgramDefContext()
}

type ProgramDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramDefContext() *ProgramDefContext {
	var p = new(ProgramDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_programDef
	return p
}

func (*ProgramDefContext) IsProgramDefContext() {}

func NewProgramDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramDefContext {
	var p = new(ProgramDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_programDef

	return p
}

func (s *ProgramDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *ProgramDefContext) AllVersionDef() []IVersionDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVersionDefContext)(nil)).Elem())
	var tst = make([]IVersionDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVersionDefContext)
		}
	}

	return tst
}

func (s *ProgramDefContext) VersionDef(i int) IVersionDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVersionDefContext)
}

func (s *ProgramDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProgramDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterProgramDef(s)
	}
}

func (s *ProgramDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitProgramDef(s)
	}
}

func (s *ProgramDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitProgramDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) ProgramDef() (localctx IProgramDefContext) {
	localctx = NewProgramDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, oncrpcv2ParserRULE_programDef)
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
		p.SetState(42)
		p.Match(oncrpcv2ParserT__0)
	}
	{
		p.SetState(43)
		p.Match(oncrpcv2ParserIDENTIFIER)
	}
	{
		p.SetState(44)
		p.Match(oncrpcv2ParserT__1)
	}
	{
		p.SetState(45)
		p.VersionDef()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == oncrpcv2ParserT__5 {
		{
			p.SetState(46)
			p.VersionDef()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(oncrpcv2ParserT__2)
	}
	{
		p.SetState(53)
		p.Match(oncrpcv2ParserT__3)
	}
	{
		p.SetState(54)
		p.Constant()
	}
	{
		p.SetState(55)
		p.Match(oncrpcv2ParserT__4)
	}

	return localctx
}

// IVersionDefContext is an interface to support dynamic dispatch.
type IVersionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionDefContext differentiates from other interfaces.
	IsVersionDefContext()
}

type VersionDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionDefContext() *VersionDefContext {
	var p = new(VersionDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_versionDef
	return p
}

func (*VersionDefContext) IsVersionDefContext() {}

func NewVersionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionDefContext {
	var p = new(VersionDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_versionDef

	return p
}

func (s *VersionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *VersionDefContext) AllProcedureDef() []IProcedureDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcedureDefContext)(nil)).Elem())
	var tst = make([]IProcedureDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcedureDefContext)
		}
	}

	return tst
}

func (s *VersionDefContext) ProcedureDef(i int) IProcedureDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedureDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcedureDefContext)
}

func (s *VersionDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *VersionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterVersionDef(s)
	}
}

func (s *VersionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitVersionDef(s)
	}
}

func (s *VersionDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitVersionDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) VersionDef() (localctx IVersionDefContext) {
	localctx = NewVersionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, oncrpcv2ParserRULE_versionDef)
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
		p.SetState(57)
		p.Match(oncrpcv2ParserT__5)
	}
	{
		p.SetState(58)
		p.Match(oncrpcv2ParserIDENTIFIER)
	}
	{
		p.SetState(59)
		p.Match(oncrpcv2ParserT__1)
	}
	{
		p.SetState(60)
		p.ProcedureDef()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(oncrpcv2ParserT__9-10))|(1<<(oncrpcv2ParserT__17-10))|(1<<(oncrpcv2ParserT__18-10))|(1<<(oncrpcv2ParserT__19-10))|(1<<(oncrpcv2ParserT__20-10))|(1<<(oncrpcv2ParserT__21-10))|(1<<(oncrpcv2ParserT__22-10))|(1<<(oncrpcv2ParserT__23-10))|(1<<(oncrpcv2ParserT__24-10))|(1<<(oncrpcv2ParserT__25-10))|(1<<(oncrpcv2ParserT__26-10))|(1<<(oncrpcv2ParserIDENTIFIER-10)))) != 0 {
		{
			p.SetState(61)
			p.ProcedureDef()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(oncrpcv2ParserT__2)
	}
	{
		p.SetState(68)
		p.Match(oncrpcv2ParserT__3)
	}
	{
		p.SetState(69)
		p.Constant()
	}
	{
		p.SetState(70)
		p.Match(oncrpcv2ParserT__4)
	}

	return localctx
}

// IProcedureDefContext is an interface to support dynamic dispatch.
type IProcedureDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedureDefContext differentiates from other interfaces.
	IsProcedureDefContext()
}

type ProcedureDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureDefContext() *ProcedureDefContext {
	var p = new(ProcedureDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_procedureDef
	return p
}

func (*ProcedureDefContext) IsProcedureDefContext() {}

func NewProcedureDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureDefContext {
	var p = new(ProcedureDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_procedureDef

	return p
}

func (s *ProcedureDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureDefContext) ProcReturn() IProcReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcReturnContext)
}

func (s *ProcedureDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *ProcedureDefContext) ProcFirstArg() IProcFirstArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcFirstArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcFirstArgContext)
}

func (s *ProcedureDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProcedureDefContext) AllTypeSpecifier() []ITypeSpecifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem())
	var tst = make([]ITypeSpecifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecifierContext)
		}
	}

	return tst
}

func (s *ProcedureDefContext) TypeSpecifier(i int) ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcedureDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterProcedureDef(s)
	}
}

func (s *ProcedureDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitProcedureDef(s)
	}
}

func (s *ProcedureDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitProcedureDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) ProcedureDef() (localctx IProcedureDefContext) {
	localctx = NewProcedureDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, oncrpcv2ParserRULE_procedureDef)
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
		p.SetState(72)
		p.ProcReturn()
	}
	{
		p.SetState(73)
		p.Match(oncrpcv2ParserIDENTIFIER)
	}
	{
		p.SetState(74)
		p.Match(oncrpcv2ParserT__6)
	}
	{
		p.SetState(75)
		p.ProcFirstArg()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == oncrpcv2ParserT__7 {
		{
			p.SetState(76)
			p.Match(oncrpcv2ParserT__7)
		}
		{
			p.SetState(77)
			p.TypeSpecifier()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
		p.Match(oncrpcv2ParserT__8)
	}
	{
		p.SetState(84)
		p.Match(oncrpcv2ParserT__3)
	}
	{
		p.SetState(85)
		p.Constant()
	}
	{
		p.SetState(86)
		p.Match(oncrpcv2ParserT__4)
	}

	return localctx
}

// IProcReturnContext is an interface to support dynamic dispatch.
type IProcReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcReturnContext differentiates from other interfaces.
	IsProcReturnContext()
}

type ProcReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcReturnContext() *ProcReturnContext {
	var p = new(ProcReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_procReturn
	return p
}

func (*ProcReturnContext) IsProcReturnContext() {}

func NewProcReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcReturnContext {
	var p = new(ProcReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_procReturn

	return p
}

func (s *ProcReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcReturnContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterProcReturn(s)
	}
}

func (s *ProcReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitProcReturn(s)
	}
}

func (s *ProcReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitProcReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) ProcReturn() (localctx IProcReturnContext) {
	localctx = NewProcReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, oncrpcv2ParserRULE_procReturn)

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

	switch p.GetTokenStream().LA(1) {
	case oncrpcv2ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(oncrpcv2ParserT__9)
		}

	case oncrpcv2ParserT__17, oncrpcv2ParserT__18, oncrpcv2ParserT__19, oncrpcv2ParserT__20, oncrpcv2ParserT__21, oncrpcv2ParserT__22, oncrpcv2ParserT__23, oncrpcv2ParserT__24, oncrpcv2ParserT__25, oncrpcv2ParserT__26, oncrpcv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.TypeSpecifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IProcFirstArgContext is an interface to support dynamic dispatch.
type IProcFirstArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcFirstArgContext differentiates from other interfaces.
	IsProcFirstArgContext()
}

type ProcFirstArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcFirstArgContext() *ProcFirstArgContext {
	var p = new(ProcFirstArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_procFirstArg
	return p
}

func (*ProcFirstArgContext) IsProcFirstArgContext() {}

func NewProcFirstArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcFirstArgContext {
	var p = new(ProcFirstArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_procFirstArg

	return p
}

func (s *ProcFirstArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcFirstArgContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ProcFirstArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcFirstArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcFirstArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterProcFirstArg(s)
	}
}

func (s *ProcFirstArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitProcFirstArg(s)
	}
}

func (s *ProcFirstArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitProcFirstArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) ProcFirstArg() (localctx IProcFirstArgContext) {
	localctx = NewProcFirstArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, oncrpcv2ParserRULE_procFirstArg)

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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case oncrpcv2ParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(oncrpcv2ParserT__9)
		}

	case oncrpcv2ParserT__17, oncrpcv2ParserT__18, oncrpcv2ParserT__19, oncrpcv2ParserT__20, oncrpcv2ParserT__21, oncrpcv2ParserT__22, oncrpcv2ParserT__23, oncrpcv2ParserT__24, oncrpcv2ParserT__25, oncrpcv2ParserT__26, oncrpcv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.TypeSpecifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOncrpcv2SpecificationContext is an interface to support dynamic dispatch.
type IOncrpcv2SpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOncrpcv2SpecificationContext differentiates from other interfaces.
	IsOncrpcv2SpecificationContext()
}

type Oncrpcv2SpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOncrpcv2SpecificationContext() *Oncrpcv2SpecificationContext {
	var p = new(Oncrpcv2SpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_oncrpcv2Specification
	return p
}

func (*Oncrpcv2SpecificationContext) IsOncrpcv2SpecificationContext() {}

func NewOncrpcv2SpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Oncrpcv2SpecificationContext {
	var p = new(Oncrpcv2SpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_oncrpcv2Specification

	return p
}

func (s *Oncrpcv2SpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Oncrpcv2SpecificationContext) AllXdrSpecification() []IXdrSpecificationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXdrSpecificationContext)(nil)).Elem())
	var tst = make([]IXdrSpecificationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXdrSpecificationContext)
		}
	}

	return tst
}

func (s *Oncrpcv2SpecificationContext) XdrSpecification(i int) IXdrSpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXdrSpecificationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXdrSpecificationContext)
}

func (s *Oncrpcv2SpecificationContext) AllProgramDef() []IProgramDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProgramDefContext)(nil)).Elem())
	var tst = make([]IProgramDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProgramDefContext)
		}
	}

	return tst
}

func (s *Oncrpcv2SpecificationContext) ProgramDef(i int) IProgramDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProgramDefContext)
}

func (s *Oncrpcv2SpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Oncrpcv2SpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Oncrpcv2SpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterOncrpcv2Specification(s)
	}
}

func (s *Oncrpcv2SpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitOncrpcv2Specification(s)
	}
}

func (s *Oncrpcv2SpecificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitOncrpcv2Specification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) Oncrpcv2Specification() (localctx IOncrpcv2SpecificationContext) {
	localctx = NewOncrpcv2SpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, oncrpcv2ParserRULE_oncrpcv2Specification)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<oncrpcv2ParserT__0)|(1<<oncrpcv2ParserT__24)|(1<<oncrpcv2ParserT__25)|(1<<oncrpcv2ParserT__26))) != 0) || _la == oncrpcv2ParserT__31 || _la == oncrpcv2ParserT__32 {
		p.SetState(98)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case oncrpcv2ParserT__24, oncrpcv2ParserT__25, oncrpcv2ParserT__26, oncrpcv2ParserT__31, oncrpcv2ParserT__32:
			{
				p.SetState(96)
				p.XdrSpecification()
			}

		case oncrpcv2ParserT__0:
			{
				p.SetState(97)
				p.ProgramDef()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *DeclarationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, oncrpcv2ParserRULE_declaration)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.TypeSpecifier()
		}
		{
			p.SetState(104)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.TypeSpecifier()
		}
		{
			p.SetState(107)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(108)
			p.Match(oncrpcv2ParserT__10)
		}
		{
			p.SetState(109)
			p.Value()
		}
		{
			p.SetState(110)
			p.Match(oncrpcv2ParserT__11)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.TypeSpecifier()
		}
		{
			p.SetState(113)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(114)
			p.Match(oncrpcv2ParserT__12)
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(oncrpcv2ParserOCTAL-35))|(1<<(oncrpcv2ParserDECIMAL-35))|(1<<(oncrpcv2ParserHEXADECIMAL-35))|(1<<(oncrpcv2ParserIDENTIFIER-35)))) != 0 {
			{
				p.SetState(115)
				p.Value()
			}

		}
		{
			p.SetState(118)
			p.Match(oncrpcv2ParserT__13)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Match(oncrpcv2ParserT__14)
		}
		{
			p.SetState(121)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(122)
			p.Match(oncrpcv2ParserT__10)
		}
		{
			p.SetState(123)
			p.Value()
		}
		{
			p.SetState(124)
			p.Match(oncrpcv2ParserT__11)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.Match(oncrpcv2ParserT__14)
		}
		{
			p.SetState(127)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(128)
			p.Match(oncrpcv2ParserT__12)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(oncrpcv2ParserOCTAL-35))|(1<<(oncrpcv2ParserDECIMAL-35))|(1<<(oncrpcv2ParserHEXADECIMAL-35))|(1<<(oncrpcv2ParserIDENTIFIER-35)))) != 0 {
			{
				p.SetState(129)
				p.Value()
			}

		}
		{
			p.SetState(132)
			p.Match(oncrpcv2ParserT__13)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Match(oncrpcv2ParserT__15)
		}
		{
			p.SetState(134)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(135)
			p.Match(oncrpcv2ParserT__12)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(oncrpcv2ParserOCTAL-35))|(1<<(oncrpcv2ParserDECIMAL-35))|(1<<(oncrpcv2ParserHEXADECIMAL-35))|(1<<(oncrpcv2ParserIDENTIFIER-35)))) != 0 {
			{
				p.SetState(136)
				p.Value()
			}

		}
		{
			p.SetState(139)
			p.Match(oncrpcv2ParserT__13)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)
			p.TypeSpecifier()
		}
		{
			p.SetState(141)
			p.Match(oncrpcv2ParserT__16)
		}
		{
			p.SetState(142)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(144)
			p.Match(oncrpcv2ParserT__9)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, oncrpcv2ParserRULE_value)

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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case oncrpcv2ParserOCTAL, oncrpcv2ParserDECIMAL, oncrpcv2ParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Constant()
		}

	case oncrpcv2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserDECIMAL, 0)
}

func (s *ConstantContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserHEXADECIMAL, 0)
}

func (s *ConstantContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserOCTAL, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, oncrpcv2ParserRULE_constant)
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
	p.SetState(151)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(oncrpcv2ParserOCTAL-35))|(1<<(oncrpcv2ParserDECIMAL-35))|(1<<(oncrpcv2ParserHEXADECIMAL-35)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) EnumTypeSpec() IEnumTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeSpecContext)
}

func (s *TypeSpecifierContext) StructTypeSpec() IStructTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTypeSpecContext)
}

func (s *TypeSpecifierContext) UnionTypeSpec() IUnionTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeSpecContext)
}

func (s *TypeSpecifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitTypeSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, oncrpcv2ParserRULE_typeSpecifier)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == oncrpcv2ParserT__17 {
			{
				p.SetState(153)
				p.Match(oncrpcv2ParserT__17)
			}

		}
		{
			p.SetState(156)
			p.Match(oncrpcv2ParserT__18)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == oncrpcv2ParserT__17 {
			{
				p.SetState(157)
				p.Match(oncrpcv2ParserT__17)
			}

		}
		{
			p.SetState(160)
			p.Match(oncrpcv2ParserT__19)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Match(oncrpcv2ParserT__20)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Match(oncrpcv2ParserT__21)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(163)
			p.Match(oncrpcv2ParserT__22)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(164)
			p.Match(oncrpcv2ParserT__23)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(165)
			p.EnumTypeSpec()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(166)
			p.StructTypeSpec()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(167)
			p.UnionTypeSpec()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(168)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}

	}

	return localctx
}

// IEnumTypeSpecContext is an interface to support dynamic dispatch.
type IEnumTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeSpecContext differentiates from other interfaces.
	IsEnumTypeSpecContext()
}

type EnumTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeSpecContext() *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_enumTypeSpec
	return p
}

func (*EnumTypeSpecContext) IsEnumTypeSpecContext() {}

func NewEnumTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_enumTypeSpec

	return p
}

func (s *EnumTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeSpecContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterEnumTypeSpec(s)
	}
}

func (s *EnumTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitEnumTypeSpec(s)
	}
}

func (s *EnumTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitEnumTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) EnumTypeSpec() (localctx IEnumTypeSpecContext) {
	localctx = NewEnumTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, oncrpcv2ParserRULE_enumTypeSpec)

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
		p.SetState(171)
		p.Match(oncrpcv2ParserT__24)
	}
	{
		p.SetState(172)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(oncrpcv2ParserIDENTIFIER)
}

func (s *EnumBodyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, i)
}

func (s *EnumBodyContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, oncrpcv2ParserRULE_enumBody)
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
		p.SetState(174)
		p.Match(oncrpcv2ParserT__1)
	}

	{
		p.SetState(175)
		p.Match(oncrpcv2ParserIDENTIFIER)
	}
	{
		p.SetState(176)
		p.Match(oncrpcv2ParserT__3)
	}
	{
		p.SetState(177)
		p.Value()
	}

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == oncrpcv2ParserT__7 {
		{
			p.SetState(179)
			p.Match(oncrpcv2ParserT__7)
		}
		{
			p.SetState(180)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(181)
			p.Match(oncrpcv2ParserT__3)
		}
		{
			p.SetState(182)
			p.Value()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(oncrpcv2ParserT__2)
	}

	return localctx
}

// IStructTypeSpecContext is an interface to support dynamic dispatch.
type IStructTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTypeSpecContext differentiates from other interfaces.
	IsStructTypeSpecContext()
}

type StructTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeSpecContext() *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_structTypeSpec
	return p
}

func (*StructTypeSpecContext) IsStructTypeSpecContext() {}

func NewStructTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_structTypeSpec

	return p
}

func (s *StructTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeSpecContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitStructTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) StructTypeSpec() (localctx IStructTypeSpecContext) {
	localctx = NewStructTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, oncrpcv2ParserRULE_structTypeSpec)

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
		p.SetState(190)
		p.Match(oncrpcv2ParserT__25)
	}
	{
		p.SetState(191)
		p.StructBody()
	}

	return localctx
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_structBody
	return p
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *StructBodyContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (s *StructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, oncrpcv2ParserRULE_structBody)
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
		p.SetState(193)
		p.Match(oncrpcv2ParserT__1)
	}

	{
		p.SetState(194)
		p.Declaration()
	}
	{
		p.SetState(195)
		p.Match(oncrpcv2ParserT__4)
	}

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(oncrpcv2ParserT__9-10))|(1<<(oncrpcv2ParserT__14-10))|(1<<(oncrpcv2ParserT__15-10))|(1<<(oncrpcv2ParserT__17-10))|(1<<(oncrpcv2ParserT__18-10))|(1<<(oncrpcv2ParserT__19-10))|(1<<(oncrpcv2ParserT__20-10))|(1<<(oncrpcv2ParserT__21-10))|(1<<(oncrpcv2ParserT__22-10))|(1<<(oncrpcv2ParserT__23-10))|(1<<(oncrpcv2ParserT__24-10))|(1<<(oncrpcv2ParserT__25-10))|(1<<(oncrpcv2ParserT__26-10))|(1<<(oncrpcv2ParserIDENTIFIER-10)))) != 0 {
		{
			p.SetState(197)
			p.Declaration()
		}
		{
			p.SetState(198)
			p.Match(oncrpcv2ParserT__4)
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(oncrpcv2ParserT__2)
	}

	return localctx
}

// IUnionTypeSpecContext is an interface to support dynamic dispatch.
type IUnionTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeSpecContext differentiates from other interfaces.
	IsUnionTypeSpecContext()
}

type UnionTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeSpecContext() *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_unionTypeSpec
	return p
}

func (*UnionTypeSpecContext) IsUnionTypeSpecContext() {}

func NewUnionTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_unionTypeSpec

	return p
}

func (s *UnionTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeSpecContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *UnionTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterUnionTypeSpec(s)
	}
}

func (s *UnionTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitUnionTypeSpec(s)
	}
}

func (s *UnionTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitUnionTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) UnionTypeSpec() (localctx IUnionTypeSpecContext) {
	localctx = NewUnionTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, oncrpcv2ParserRULE_unionTypeSpec)

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
		p.Match(oncrpcv2ParserT__26)
	}
	{
		p.SetState(208)
		p.UnionBody()
	}

	return localctx
}

// IUnionBodyContext is an interface to support dynamic dispatch.
type IUnionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionBodyContext differentiates from other interfaces.
	IsUnionBodyContext()
}

type UnionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionBodyContext() *UnionBodyContext {
	var p = new(UnionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_unionBody
	return p
}

func (*UnionBodyContext) IsUnionBodyContext() {}

func NewUnionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionBodyContext {
	var p = new(UnionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_unionBody

	return p
}

func (s *UnionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionBodyContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *UnionBodyContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *UnionBodyContext) AllCaseSpec() []ICaseSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem())
	var tst = make([]ICaseSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseSpecContext)
		}
	}

	return tst
}

func (s *UnionBodyContext) CaseSpec(i int) ICaseSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseSpecContext)
}

func (s *UnionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterUnionBody(s)
	}
}

func (s *UnionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitUnionBody(s)
	}
}

func (s *UnionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitUnionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) UnionBody() (localctx IUnionBodyContext) {
	localctx = NewUnionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, oncrpcv2ParserRULE_unionBody)
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
		p.SetState(210)
		p.Match(oncrpcv2ParserT__27)
	}
	{
		p.SetState(211)
		p.Match(oncrpcv2ParserT__6)
	}
	{
		p.SetState(212)
		p.Declaration()
	}
	{
		p.SetState(213)
		p.Match(oncrpcv2ParserT__8)
	}
	{
		p.SetState(214)
		p.Match(oncrpcv2ParserT__1)
	}
	{
		p.SetState(215)
		p.CaseSpec()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == oncrpcv2ParserT__30 {
		{
			p.SetState(216)
			p.CaseSpec()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == oncrpcv2ParserT__28 {
		{
			p.SetState(222)
			p.Match(oncrpcv2ParserT__28)
		}
		{
			p.SetState(223)
			p.Match(oncrpcv2ParserT__29)
		}
		{
			p.SetState(224)
			p.Declaration()
		}
		{
			p.SetState(225)
			p.Match(oncrpcv2ParserT__4)
		}

	}
	{
		p.SetState(229)
		p.Match(oncrpcv2ParserT__2)
	}

	return localctx
}

// ICaseSpecContext is an interface to support dynamic dispatch.
type ICaseSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseSpecContext differentiates from other interfaces.
	IsCaseSpecContext()
}

type CaseSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseSpecContext() *CaseSpecContext {
	var p = new(CaseSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_caseSpec
	return p
}

func (*CaseSpecContext) IsCaseSpecContext() {}

func NewCaseSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseSpecContext {
	var p = new(CaseSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_caseSpec

	return p
}

func (s *CaseSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseSpecContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *CaseSpecContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *CaseSpecContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CaseSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterCaseSpec(s)
	}
}

func (s *CaseSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitCaseSpec(s)
	}
}

func (s *CaseSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitCaseSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) CaseSpec() (localctx ICaseSpecContext) {
	localctx = NewCaseSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, oncrpcv2ParserRULE_caseSpec)
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
		p.SetState(231)
		p.Match(oncrpcv2ParserT__30)
	}
	{
		p.SetState(232)
		p.Value()
	}
	{
		p.SetState(233)
		p.Match(oncrpcv2ParserT__29)
	}

	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == oncrpcv2ParserT__30 {
		{
			p.SetState(235)
			p.Match(oncrpcv2ParserT__30)
		}
		{
			p.SetState(236)
			p.Value()
		}
		{
			p.SetState(237)
			p.Match(oncrpcv2ParserT__29)
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(244)
		p.Declaration()
	}
	{
		p.SetState(245)
		p.Match(oncrpcv2ParserT__4)
	}

	return localctx
}

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	var p = new(ConstantDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_constantDef
	return p
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	var p = new(ConstantDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (s *ConstantDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitConstantDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) ConstantDef() (localctx IConstantDefContext) {
	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, oncrpcv2ParserRULE_constantDef)

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
		p.SetState(247)
		p.Match(oncrpcv2ParserT__31)
	}
	{
		p.SetState(248)
		p.Match(oncrpcv2ParserIDENTIFIER)
	}
	{
		p.SetState(249)
		p.Match(oncrpcv2ParserT__3)
	}
	{
		p.SetState(250)
		p.Constant()
	}
	{
		p.SetState(251)
		p.Match(oncrpcv2ParserT__4)
	}

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TypeDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(oncrpcv2ParserIDENTIFIER, 0)
}

func (s *TypeDefContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TypeDefContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *TypeDefContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (s *TypeDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitTypeDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, oncrpcv2ParserRULE_typeDef)

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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case oncrpcv2ParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(oncrpcv2ParserT__32)
		}
		{
			p.SetState(254)
			p.Declaration()
		}
		{
			p.SetState(255)
			p.Match(oncrpcv2ParserT__4)
		}

	case oncrpcv2ParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Match(oncrpcv2ParserT__24)
		}
		{
			p.SetState(258)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(259)
			p.EnumBody()
		}
		{
			p.SetState(260)
			p.Match(oncrpcv2ParserT__4)
		}

	case oncrpcv2ParserT__25:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.Match(oncrpcv2ParserT__25)
		}
		{
			p.SetState(263)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(264)
			p.StructBody()
		}
		{
			p.SetState(265)
			p.Match(oncrpcv2ParserT__4)
		}

	case oncrpcv2ParserT__26:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(267)
			p.Match(oncrpcv2ParserT__26)
		}
		{
			p.SetState(268)
			p.Match(oncrpcv2ParserIDENTIFIER)
		}
		{
			p.SetState(269)
			p.UnionBody()
		}
		{
			p.SetState(270)
			p.Match(oncrpcv2ParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) TypeDef() ITypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *DefinitionContext) ConstantDef() IConstantDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantDefContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, oncrpcv2ParserRULE_definition)

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

	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case oncrpcv2ParserT__24, oncrpcv2ParserT__25, oncrpcv2ParserT__26, oncrpcv2ParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.TypeDef()
		}

	case oncrpcv2ParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.ConstantDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IXdrSpecificationContext is an interface to support dynamic dispatch.
type IXdrSpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXdrSpecificationContext differentiates from other interfaces.
	IsXdrSpecificationContext()
}

type XdrSpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXdrSpecificationContext() *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = oncrpcv2ParserRULE_xdrSpecification
	return p
}

func (*XdrSpecificationContext) IsXdrSpecificationContext() {}

func NewXdrSpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = oncrpcv2ParserRULE_xdrSpecification

	return p
}

func (s *XdrSpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *XdrSpecificationContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *XdrSpecificationContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *XdrSpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XdrSpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XdrSpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.EnterXdrSpecification(s)
	}
}

func (s *XdrSpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oncrpcv2Listener); ok {
		listenerT.ExitXdrSpecification(s)
	}
}

func (s *XdrSpecificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oncrpcv2Visitor:
		return t.VisitXdrSpecification(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oncrpcv2Parser) XdrSpecification() (localctx IXdrSpecificationContext) {
	localctx = NewXdrSpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, oncrpcv2ParserRULE_xdrSpecification)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(278)
				p.Definition()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}
