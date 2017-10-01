// Generated from Cobol85Preprocessor.g4 by ANTLR 4.7.

package cobol85 // Cobol85Preprocessor
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 265,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 55, 10, 2, 12, 2,
	14, 2, 58, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 67, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 74, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 81, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6, 86, 10, 6, 12, 6, 14,
	6, 89, 11, 6, 3, 6, 5, 6, 92, 10, 6, 3, 6, 7, 6, 95, 10, 6, 12, 6, 14,
	6, 98, 11, 6, 3, 6, 5, 6, 101, 10, 6, 3, 6, 7, 6, 104, 10, 6, 12, 6, 14,
	6, 107, 11, 6, 3, 6, 5, 6, 110, 10, 6, 3, 6, 5, 6, 113, 10, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 5, 7, 119, 10, 7, 3, 8, 3, 8, 7, 8, 123, 10, 8, 12, 8, 14,
	8, 126, 11, 8, 3, 8, 3, 8, 6, 8, 130, 10, 8, 13, 8, 14, 8, 131, 3, 8, 7,
	8, 135, 10, 8, 12, 8, 14, 8, 138, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 143, 10,
	9, 12, 9, 14, 9, 146, 11, 9, 3, 9, 5, 9, 149, 10, 9, 3, 10, 3, 10, 7, 10,
	153, 10, 10, 12, 10, 14, 10, 156, 11, 10, 3, 10, 6, 10, 159, 10, 10, 13,
	10, 14, 10, 160, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	7, 12, 171, 10, 12, 12, 12, 14, 12, 174, 11, 12, 3, 12, 3, 12, 7, 12, 178,
	10, 12, 12, 12, 14, 12, 181, 11, 12, 3, 12, 3, 12, 7, 12, 185, 10, 12,
	12, 12, 14, 12, 188, 11, 12, 3, 12, 5, 12, 191, 10, 12, 3, 12, 7, 12, 194,
	10, 12, 12, 12, 14, 12, 197, 11, 12, 3, 12, 5, 12, 200, 10, 12, 3, 13,
	3, 13, 7, 13, 204, 10, 13, 12, 13, 14, 13, 207, 11, 13, 3, 13, 3, 13, 5,
	13, 211, 10, 13, 3, 14, 3, 14, 7, 14, 215, 10, 14, 12, 14, 14, 14, 218,
	11, 14, 3, 14, 3, 14, 5, 14, 222, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 228, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 234, 10, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 244, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 6, 21, 250, 10, 21, 13, 21, 14, 21, 251, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 6, 22, 259, 10, 22, 13, 22, 14, 22, 260, 3,
	23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 5, 3, 2, 9, 10, 4, 2, 6, 6,
	17, 19, 5, 2, 3, 3, 9, 12, 14, 14, 2, 293, 2, 56, 3, 2, 2, 2, 4, 61, 3,
	2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 82, 3, 2, 2, 2, 12,
	118, 3, 2, 2, 2, 14, 120, 3, 2, 2, 2, 16, 139, 3, 2, 2, 2, 18, 150, 3,
	2, 2, 2, 20, 164, 3, 2, 2, 2, 22, 168, 3, 2, 2, 2, 24, 201, 3, 2, 2, 2,
	26, 212, 3, 2, 2, 2, 28, 227, 3, 2, 2, 2, 30, 233, 3, 2, 2, 2, 32, 235,
	3, 2, 2, 2, 34, 237, 3, 2, 2, 2, 36, 239, 3, 2, 2, 2, 38, 241, 3, 2, 2,
	2, 40, 249, 3, 2, 2, 2, 42, 258, 3, 2, 2, 2, 44, 262, 3, 2, 2, 2, 46, 55,
	5, 10, 6, 2, 47, 55, 5, 4, 3, 2, 48, 55, 5, 6, 4, 2, 49, 55, 5, 8, 5, 2,
	50, 55, 5, 20, 11, 2, 51, 55, 5, 16, 9, 2, 52, 55, 5, 40, 21, 2, 53, 55,
	5, 32, 17, 2, 54, 46, 3, 2, 2, 2, 54, 47, 3, 2, 2, 2, 54, 48, 3, 2, 2,
	2, 54, 49, 3, 2, 2, 2, 54, 50, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2,
	56, 57, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7,
	2, 2, 3, 60, 3, 3, 2, 2, 2, 61, 62, 7, 8, 2, 2, 62, 63, 7, 4, 2, 2, 63,
	64, 5, 40, 21, 2, 64, 66, 7, 7, 2, 2, 65, 67, 7, 22, 2, 2, 66, 65, 3, 2,
	2, 2, 66, 67, 3, 2, 2, 2, 67, 5, 3, 2, 2, 2, 68, 69, 7, 8, 2, 2, 69, 70,
	7, 15, 2, 2, 70, 71, 5, 40, 21, 2, 71, 73, 7, 7, 2, 2, 72, 74, 7, 22, 2,
	2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7,
	8, 2, 2, 76, 77, 7, 16, 2, 2, 77, 78, 5, 40, 21, 2, 78, 80, 7, 7, 2, 2,
	79, 81, 7, 22, 2, 2, 80, 79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 9, 3,
	2, 2, 2, 82, 83, 7, 5, 2, 2, 83, 91, 5, 12, 7, 2, 84, 86, 7, 26, 2, 2,
	85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 92, 5, 24, 13, 2,
	91, 87, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 100, 3, 2, 2, 2, 93, 95, 7,
	26, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96,
	97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 101, 5, 26,
	14, 2, 100, 96, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 109, 3, 2, 2, 2,
	102, 104, 7, 26, 2, 2, 103, 102, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105,
	103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 105,
	3, 2, 2, 2, 108, 110, 5, 14, 8, 2, 109, 105, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 110, 112, 3, 2, 2, 2, 111, 113, 7, 20, 2, 2, 112, 111, 3, 2, 2, 2,
	112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 7, 22, 2, 2, 115,
	11, 3, 2, 2, 2, 116, 119, 5, 36, 19, 2, 117, 119, 5, 34, 18, 2, 118, 116,
	3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 119, 13, 3, 2, 2, 2, 120, 124, 7, 14,
	2, 2, 121, 123, 7, 26, 2, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2,
	124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 136, 5, 22, 12, 2, 128, 130, 7, 26, 2, 2, 129, 128,
	3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 5, 22, 12, 2, 134, 129, 3, 2, 2,
	2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	15, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 144, 5, 18, 10, 2, 140, 143,
	5, 10, 6, 2, 141, 143, 5, 40, 21, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3,
	2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 149, 5, 20, 11, 2,
	148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 17, 3, 2, 2, 2, 150, 158,
	7, 13, 2, 2, 151, 153, 7, 26, 2, 2, 152, 151, 3, 2, 2, 2, 153, 156, 3,
	2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 157, 3, 2, 2,
	2, 156, 154, 3, 2, 2, 2, 157, 159, 5, 22, 12, 2, 158, 154, 3, 2, 2, 2,
	159, 160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161,
	162, 3, 2, 2, 2, 162, 163, 7, 22, 2, 2, 163, 19, 3, 2, 2, 2, 164, 165,
	7, 13, 2, 2, 165, 166, 7, 11, 2, 2, 166, 167, 7, 22, 2, 2, 167, 21, 3,
	2, 2, 2, 168, 172, 5, 28, 15, 2, 169, 171, 7, 26, 2, 2, 170, 169, 3, 2,
	2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2,
	173, 175, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 179, 7, 3, 2, 2, 176,
	178, 7, 26, 2, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179, 177,
	3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181, 179, 3, 2,
	2, 2, 182, 190, 5, 30, 16, 2, 183, 185, 7, 26, 2, 2, 184, 183, 3, 2, 2,
	2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187,
	189, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 191, 5, 24, 13, 2, 190, 186,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 199, 3, 2, 2, 2, 192, 194, 7, 26,
	2, 2, 193, 192, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2,
	195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198,
	200, 5, 26, 14, 2, 199, 195, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 23,
	3, 2, 2, 2, 201, 205, 9, 2, 2, 2, 202, 204, 7, 26, 2, 2, 203, 202, 3, 2,
	2, 2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2,
	206, 210, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 211, 5, 36, 19, 2, 209,
	211, 5, 34, 18, 2, 210, 208, 3, 2, 2, 2, 210, 209, 3, 2, 2, 2, 211, 25,
	3, 2, 2, 2, 212, 216, 7, 12, 2, 2, 213, 215, 7, 26, 2, 2, 214, 213, 3,
	2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2,
	2, 217, 221, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 222, 5, 36, 19, 2,
	220, 222, 5, 34, 18, 2, 221, 219, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222,
	27, 3, 2, 2, 2, 223, 228, 5, 36, 19, 2, 224, 228, 5, 34, 18, 2, 225, 228,
	5, 38, 20, 2, 226, 228, 5, 42, 22, 2, 227, 223, 3, 2, 2, 2, 227, 224, 3,
	2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 2, 2, 228, 29, 3, 2, 2,
	2, 229, 234, 5, 36, 19, 2, 230, 234, 5, 34, 18, 2, 231, 234, 5, 38, 20,
	2, 232, 234, 5, 42, 22, 2, 233, 229, 3, 2, 2, 2, 233, 230, 3, 2, 2, 2,
	233, 231, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234, 31, 3, 2, 2, 2, 235, 236,
	9, 3, 2, 2, 236, 33, 3, 2, 2, 2, 237, 238, 7, 25, 2, 2, 238, 35, 3, 2,
	2, 2, 239, 240, 7, 24, 2, 2, 240, 37, 3, 2, 2, 2, 241, 243, 7, 23, 2, 2,
	242, 244, 5, 40, 21, 2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244,
	245, 3, 2, 2, 2, 245, 246, 7, 23, 2, 2, 246, 39, 3, 2, 2, 2, 247, 250,
	5, 42, 22, 2, 248, 250, 7, 26, 2, 2, 249, 247, 3, 2, 2, 2, 249, 248, 3,
	2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2,
	2, 252, 41, 3, 2, 2, 2, 253, 259, 5, 44, 23, 2, 254, 259, 5, 34, 18, 2,
	255, 259, 5, 36, 19, 2, 256, 259, 7, 29, 2, 2, 257, 259, 7, 22, 2, 2, 258,
	253, 3, 2, 2, 2, 258, 254, 3, 2, 2, 2, 258, 255, 3, 2, 2, 2, 258, 256,
	3, 2, 2, 2, 258, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 258, 3, 2,
	2, 2, 260, 261, 3, 2, 2, 2, 261, 43, 3, 2, 2, 2, 262, 263, 9, 4, 2, 2,
	263, 45, 3, 2, 2, 2, 40, 54, 56, 66, 73, 80, 87, 91, 96, 100, 105, 109,
	112, 118, 124, 131, 136, 142, 144, 148, 154, 160, 172, 179, 186, 190, 195,
	199, 205, 210, 216, 221, 227, 233, 243, 249, 251, 258, 260,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "'>*'", "'.'", "'=='",
}
var symbolicNames = []string{
	"", "BY", "CICS", "COPY", "EJECT", "END_EXEC", "EXEC", "IN", "OF", "OFF",
	"ON", "REPLACE", "REPLACING", "SQL", "SQLIMS", "SKIP1", "SKIP2", "SKIP3",
	"SUPPRESS", "COMMENTTAG", "DOT", "DOUBLEEQUALCHAR", "NONNUMERICLITERAL",
	"IDENTIFIER", "NEWLINE", "COMMENTLINE", "WS", "TEXT",
}

var ruleNames = []string{
	"startRule", "execCicsStatement", "execSqlStatement", "execSqlImsStatement",
	"copyStatement", "copySource", "replacingPhrase", "replaceArea", "replaceByStatement",
	"replaceOffStatement", "replaceClause", "directoryPhrase", "familyPhrase",
	"replaceable", "replacement", "controlSpacingStatement", "cobolWord", "literal",
	"pseudoText", "charData", "charDataLine", "charDataKeyword",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Cobol85PreprocessorParser struct {
	*antlr.BaseParser
}

func NewCobol85PreprocessorParser(input antlr.TokenStream) *Cobol85PreprocessorParser {
	this := new(Cobol85PreprocessorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Cobol85Preprocessor.g4"

	return this
}

// Cobol85PreprocessorParser tokens.
const (
	Cobol85PreprocessorParserEOF               = antlr.TokenEOF
	Cobol85PreprocessorParserBY                = 1
	Cobol85PreprocessorParserCICS              = 2
	Cobol85PreprocessorParserCOPY              = 3
	Cobol85PreprocessorParserEJECT             = 4
	Cobol85PreprocessorParserEND_EXEC          = 5
	Cobol85PreprocessorParserEXEC              = 6
	Cobol85PreprocessorParserIN                = 7
	Cobol85PreprocessorParserOF                = 8
	Cobol85PreprocessorParserOFF               = 9
	Cobol85PreprocessorParserON                = 10
	Cobol85PreprocessorParserREPLACE           = 11
	Cobol85PreprocessorParserREPLACING         = 12
	Cobol85PreprocessorParserSQL               = 13
	Cobol85PreprocessorParserSQLIMS            = 14
	Cobol85PreprocessorParserSKIP1             = 15
	Cobol85PreprocessorParserSKIP2             = 16
	Cobol85PreprocessorParserSKIP3             = 17
	Cobol85PreprocessorParserSUPPRESS          = 18
	Cobol85PreprocessorParserCOMMENTTAG        = 19
	Cobol85PreprocessorParserDOT               = 20
	Cobol85PreprocessorParserDOUBLEEQUALCHAR   = 21
	Cobol85PreprocessorParserNONNUMERICLITERAL = 22
	Cobol85PreprocessorParserIDENTIFIER        = 23
	Cobol85PreprocessorParserNEWLINE           = 24
	Cobol85PreprocessorParserCOMMENTLINE       = 25
	Cobol85PreprocessorParserWS                = 26
	Cobol85PreprocessorParserTEXT              = 27
)

// Cobol85PreprocessorParser rules.
const (
	Cobol85PreprocessorParserRULE_startRule               = 0
	Cobol85PreprocessorParserRULE_execCicsStatement       = 1
	Cobol85PreprocessorParserRULE_execSqlStatement        = 2
	Cobol85PreprocessorParserRULE_execSqlImsStatement     = 3
	Cobol85PreprocessorParserRULE_copyStatement           = 4
	Cobol85PreprocessorParserRULE_copySource              = 5
	Cobol85PreprocessorParserRULE_replacingPhrase         = 6
	Cobol85PreprocessorParserRULE_replaceArea             = 7
	Cobol85PreprocessorParserRULE_replaceByStatement      = 8
	Cobol85PreprocessorParserRULE_replaceOffStatement     = 9
	Cobol85PreprocessorParserRULE_replaceClause           = 10
	Cobol85PreprocessorParserRULE_directoryPhrase         = 11
	Cobol85PreprocessorParserRULE_familyPhrase            = 12
	Cobol85PreprocessorParserRULE_replaceable             = 13
	Cobol85PreprocessorParserRULE_replacement             = 14
	Cobol85PreprocessorParserRULE_controlSpacingStatement = 15
	Cobol85PreprocessorParserRULE_cobolWord               = 16
	Cobol85PreprocessorParserRULE_literal                 = 17
	Cobol85PreprocessorParserRULE_pseudoText              = 18
	Cobol85PreprocessorParserRULE_charData                = 19
	Cobol85PreprocessorParserRULE_charDataLine            = 20
	Cobol85PreprocessorParserRULE_charDataKeyword         = 21
)

// IStartRuleContext is an interface to support dynamic dispatch.
type IStartRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartRuleContext differentiates from other interfaces.
	IsStartRuleContext()
}

type StartRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartRuleContext() *StartRuleContext {
	var p = new(StartRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_startRule
	return p
}

func (*StartRuleContext) IsStartRuleContext() {}

func NewStartRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartRuleContext {
	var p = new(StartRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_startRule

	return p
}

func (s *StartRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StartRuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEOF, 0)
}

func (s *StartRuleContext) AllCopyStatement() []ICopyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICopyStatementContext)(nil)).Elem())
	var tst = make([]ICopyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICopyStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) CopyStatement(i int) ICopyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICopyStatementContext)
}

func (s *StartRuleContext) AllExecCicsStatement() []IExecCicsStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecCicsStatementContext)(nil)).Elem())
	var tst = make([]IExecCicsStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecCicsStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ExecCicsStatement(i int) IExecCicsStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecCicsStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecCicsStatementContext)
}

func (s *StartRuleContext) AllExecSqlStatement() []IExecSqlStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecSqlStatementContext)(nil)).Elem())
	var tst = make([]IExecSqlStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecSqlStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ExecSqlStatement(i int) IExecSqlStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecSqlStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecSqlStatementContext)
}

func (s *StartRuleContext) AllExecSqlImsStatement() []IExecSqlImsStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExecSqlImsStatementContext)(nil)).Elem())
	var tst = make([]IExecSqlImsStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExecSqlImsStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ExecSqlImsStatement(i int) IExecSqlImsStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecSqlImsStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExecSqlImsStatementContext)
}

func (s *StartRuleContext) AllReplaceOffStatement() []IReplaceOffStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReplaceOffStatementContext)(nil)).Elem())
	var tst = make([]IReplaceOffStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReplaceOffStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ReplaceOffStatement(i int) IReplaceOffStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceOffStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReplaceOffStatementContext)
}

func (s *StartRuleContext) AllReplaceArea() []IReplaceAreaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReplaceAreaContext)(nil)).Elem())
	var tst = make([]IReplaceAreaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReplaceAreaContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ReplaceArea(i int) IReplaceAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceAreaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReplaceAreaContext)
}

func (s *StartRuleContext) AllCharData() []ICharDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharDataContext)(nil)).Elem())
	var tst = make([]ICharDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharDataContext)
		}
	}

	return tst
}

func (s *StartRuleContext) CharData(i int) ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *StartRuleContext) AllControlSpacingStatement() []IControlSpacingStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IControlSpacingStatementContext)(nil)).Elem())
	var tst = make([]IControlSpacingStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IControlSpacingStatementContext)
		}
	}

	return tst
}

func (s *StartRuleContext) ControlSpacingStatement(i int) IControlSpacingStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlSpacingStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IControlSpacingStatementContext)
}

func (s *StartRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterStartRule(s)
	}
}

func (s *StartRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitStartRule(s)
	}
}

func (s *StartRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitStartRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) StartRule() (localctx IStartRuleContext) {
	localctx = NewStartRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Cobol85PreprocessorParserRULE_startRule)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Cobol85PreprocessorParserBY)|(1<<Cobol85PreprocessorParserCOPY)|(1<<Cobol85PreprocessorParserEJECT)|(1<<Cobol85PreprocessorParserEXEC)|(1<<Cobol85PreprocessorParserIN)|(1<<Cobol85PreprocessorParserOF)|(1<<Cobol85PreprocessorParserOFF)|(1<<Cobol85PreprocessorParserON)|(1<<Cobol85PreprocessorParserREPLACE)|(1<<Cobol85PreprocessorParserREPLACING)|(1<<Cobol85PreprocessorParserSKIP1)|(1<<Cobol85PreprocessorParserSKIP2)|(1<<Cobol85PreprocessorParserSKIP3)|(1<<Cobol85PreprocessorParserDOT)|(1<<Cobol85PreprocessorParserNONNUMERICLITERAL)|(1<<Cobol85PreprocessorParserIDENTIFIER)|(1<<Cobol85PreprocessorParserNEWLINE)|(1<<Cobol85PreprocessorParserTEXT))) != 0 {
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(44)
				p.CopyStatement()
			}

		case 2:
			{
				p.SetState(45)
				p.ExecCicsStatement()
			}

		case 3:
			{
				p.SetState(46)
				p.ExecSqlStatement()
			}

		case 4:
			{
				p.SetState(47)
				p.ExecSqlImsStatement()
			}

		case 5:
			{
				p.SetState(48)
				p.ReplaceOffStatement()
			}

		case 6:
			{
				p.SetState(49)
				p.ReplaceArea()
			}

		case 7:
			{
				p.SetState(50)
				p.CharData()
			}

		case 8:
			{
				p.SetState(51)
				p.ControlSpacingStatement()
			}

		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(Cobol85PreprocessorParserEOF)
	}

	return localctx
}

// IExecCicsStatementContext is an interface to support dynamic dispatch.
type IExecCicsStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecCicsStatementContext differentiates from other interfaces.
	IsExecCicsStatementContext()
}

type ExecCicsStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecCicsStatementContext() *ExecCicsStatementContext {
	var p = new(ExecCicsStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_execCicsStatement
	return p
}

func (*ExecCicsStatementContext) IsExecCicsStatementContext() {}

func NewExecCicsStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecCicsStatementContext {
	var p = new(ExecCicsStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_execCicsStatement

	return p
}

func (s *ExecCicsStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecCicsStatementContext) EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEXEC, 0)
}

func (s *ExecCicsStatementContext) CICS() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserCICS, 0)
}

func (s *ExecCicsStatementContext) CharData() ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *ExecCicsStatementContext) END_EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEND_EXEC, 0)
}

func (s *ExecCicsStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *ExecCicsStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecCicsStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecCicsStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterExecCicsStatement(s)
	}
}

func (s *ExecCicsStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitExecCicsStatement(s)
	}
}

func (s *ExecCicsStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitExecCicsStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ExecCicsStatement() (localctx IExecCicsStatementContext) {
	localctx = NewExecCicsStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Cobol85PreprocessorParserRULE_execCicsStatement)

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
		p.SetState(59)
		p.Match(Cobol85PreprocessorParserEXEC)
	}
	{
		p.SetState(60)
		p.Match(Cobol85PreprocessorParserCICS)
	}
	{
		p.SetState(61)
		p.CharData()
	}
	{
		p.SetState(62)
		p.Match(Cobol85PreprocessorParserEND_EXEC)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(63)
			p.Match(Cobol85PreprocessorParserDOT)
		}

	}

	return localctx
}

// IExecSqlStatementContext is an interface to support dynamic dispatch.
type IExecSqlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecSqlStatementContext differentiates from other interfaces.
	IsExecSqlStatementContext()
}

type ExecSqlStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecSqlStatementContext() *ExecSqlStatementContext {
	var p = new(ExecSqlStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_execSqlStatement
	return p
}

func (*ExecSqlStatementContext) IsExecSqlStatementContext() {}

func NewExecSqlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecSqlStatementContext {
	var p = new(ExecSqlStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_execSqlStatement

	return p
}

func (s *ExecSqlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecSqlStatementContext) EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEXEC, 0)
}

func (s *ExecSqlStatementContext) SQL() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSQL, 0)
}

func (s *ExecSqlStatementContext) CharData() ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *ExecSqlStatementContext) END_EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEND_EXEC, 0)
}

func (s *ExecSqlStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *ExecSqlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecSqlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecSqlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterExecSqlStatement(s)
	}
}

func (s *ExecSqlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitExecSqlStatement(s)
	}
}

func (s *ExecSqlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitExecSqlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ExecSqlStatement() (localctx IExecSqlStatementContext) {
	localctx = NewExecSqlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Cobol85PreprocessorParserRULE_execSqlStatement)

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
		p.Match(Cobol85PreprocessorParserEXEC)
	}
	{
		p.SetState(67)
		p.Match(Cobol85PreprocessorParserSQL)
	}
	{
		p.SetState(68)
		p.CharData()
	}
	{
		p.SetState(69)
		p.Match(Cobol85PreprocessorParserEND_EXEC)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(70)
			p.Match(Cobol85PreprocessorParserDOT)
		}

	}

	return localctx
}

// IExecSqlImsStatementContext is an interface to support dynamic dispatch.
type IExecSqlImsStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecSqlImsStatementContext differentiates from other interfaces.
	IsExecSqlImsStatementContext()
}

type ExecSqlImsStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecSqlImsStatementContext() *ExecSqlImsStatementContext {
	var p = new(ExecSqlImsStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_execSqlImsStatement
	return p
}

func (*ExecSqlImsStatementContext) IsExecSqlImsStatementContext() {}

func NewExecSqlImsStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecSqlImsStatementContext {
	var p = new(ExecSqlImsStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_execSqlImsStatement

	return p
}

func (s *ExecSqlImsStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecSqlImsStatementContext) EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEXEC, 0)
}

func (s *ExecSqlImsStatementContext) SQLIMS() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSQLIMS, 0)
}

func (s *ExecSqlImsStatementContext) CharData() ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *ExecSqlImsStatementContext) END_EXEC() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEND_EXEC, 0)
}

func (s *ExecSqlImsStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *ExecSqlImsStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecSqlImsStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecSqlImsStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterExecSqlImsStatement(s)
	}
}

func (s *ExecSqlImsStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitExecSqlImsStatement(s)
	}
}

func (s *ExecSqlImsStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitExecSqlImsStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ExecSqlImsStatement() (localctx IExecSqlImsStatementContext) {
	localctx = NewExecSqlImsStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Cobol85PreprocessorParserRULE_execSqlImsStatement)

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
		p.SetState(73)
		p.Match(Cobol85PreprocessorParserEXEC)
	}
	{
		p.SetState(74)
		p.Match(Cobol85PreprocessorParserSQLIMS)
	}
	{
		p.SetState(75)
		p.CharData()
	}
	{
		p.SetState(76)
		p.Match(Cobol85PreprocessorParserEND_EXEC)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(77)
			p.Match(Cobol85PreprocessorParserDOT)
		}

	}

	return localctx
}

// ICopyStatementContext is an interface to support dynamic dispatch.
type ICopyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopyStatementContext differentiates from other interfaces.
	IsCopyStatementContext()
}

type CopyStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopyStatementContext() *CopyStatementContext {
	var p = new(CopyStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_copyStatement
	return p
}

func (*CopyStatementContext) IsCopyStatementContext() {}

func NewCopyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyStatementContext {
	var p = new(CopyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_copyStatement

	return p
}

func (s *CopyStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyStatementContext) COPY() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserCOPY, 0)
}

func (s *CopyStatementContext) CopySource() ICopySourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopySourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopySourceContext)
}

func (s *CopyStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *CopyStatementContext) DirectoryPhrase() IDirectoryPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectoryPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectoryPhraseContext)
}

func (s *CopyStatementContext) FamilyPhrase() IFamilyPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFamilyPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFamilyPhraseContext)
}

func (s *CopyStatementContext) ReplacingPhrase() IReplacingPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplacingPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplacingPhraseContext)
}

func (s *CopyStatementContext) SUPPRESS() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSUPPRESS, 0)
}

func (s *CopyStatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *CopyStatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *CopyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCopyStatement(s)
	}
}

func (s *CopyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCopyStatement(s)
	}
}

func (s *CopyStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCopyStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CopyStatement() (localctx ICopyStatementContext) {
	localctx = NewCopyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Cobol85PreprocessorParserRULE_copyStatement)
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
		p.SetState(80)
		p.Match(Cobol85PreprocessorParserCOPY)
	}
	{
		p.SetState(81)
		p.CopySource()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(82)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(88)
			p.DirectoryPhrase()
		}

	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(91)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.FamilyPhrase()
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Cobol85PreprocessorParserREPLACING || _la == Cobol85PreprocessorParserNEWLINE {
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(100)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.ReplacingPhrase()
		}

	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Cobol85PreprocessorParserSUPPRESS {
		{
			p.SetState(109)
			p.Match(Cobol85PreprocessorParserSUPPRESS)
		}

	}
	{
		p.SetState(112)
		p.Match(Cobol85PreprocessorParserDOT)
	}

	return localctx
}

// ICopySourceContext is an interface to support dynamic dispatch.
type ICopySourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopySourceContext differentiates from other interfaces.
	IsCopySourceContext()
}

type CopySourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopySourceContext() *CopySourceContext {
	var p = new(CopySourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_copySource
	return p
}

func (*CopySourceContext) IsCopySourceContext() {}

func NewCopySourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopySourceContext {
	var p = new(CopySourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_copySource

	return p
}

func (s *CopySourceContext) GetParser() antlr.Parser { return s.parser }

func (s *CopySourceContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *CopySourceContext) CobolWord() ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *CopySourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopySourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopySourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCopySource(s)
	}
}

func (s *CopySourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCopySource(s)
	}
}

func (s *CopySourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCopySource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CopySource() (localctx ICopySourceContext) {
	localctx = NewCopySourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Cobol85PreprocessorParserRULE_copySource)

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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Cobol85PreprocessorParserNONNUMERICLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Literal()
		}

	case Cobol85PreprocessorParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.CobolWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReplacingPhraseContext is an interface to support dynamic dispatch.
type IReplacingPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplacingPhraseContext differentiates from other interfaces.
	IsReplacingPhraseContext()
}

type ReplacingPhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplacingPhraseContext() *ReplacingPhraseContext {
	var p = new(ReplacingPhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replacingPhrase
	return p
}

func (*ReplacingPhraseContext) IsReplacingPhraseContext() {}

func NewReplacingPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplacingPhraseContext {
	var p = new(ReplacingPhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replacingPhrase

	return p
}

func (s *ReplacingPhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplacingPhraseContext) REPLACING() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserREPLACING, 0)
}

func (s *ReplacingPhraseContext) AllReplaceClause() []IReplaceClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReplaceClauseContext)(nil)).Elem())
	var tst = make([]IReplaceClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReplaceClauseContext)
		}
	}

	return tst
}

func (s *ReplacingPhraseContext) ReplaceClause(i int) IReplaceClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReplaceClauseContext)
}

func (s *ReplacingPhraseContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *ReplacingPhraseContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *ReplacingPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplacingPhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplacingPhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplacingPhrase(s)
	}
}

func (s *ReplacingPhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplacingPhrase(s)
	}
}

func (s *ReplacingPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplacingPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ReplacingPhrase() (localctx IReplacingPhraseContext) {
	localctx = NewReplacingPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Cobol85PreprocessorParserRULE_replacingPhrase)
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
		p.SetState(118)
		p.Match(Cobol85PreprocessorParserREPLACING)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		{
			p.SetState(119)
			p.Match(Cobol85PreprocessorParserNEWLINE)
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.ReplaceClause()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(126)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(131)
			p.ReplaceClause()
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReplaceAreaContext is an interface to support dynamic dispatch.
type IReplaceAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceAreaContext differentiates from other interfaces.
	IsReplaceAreaContext()
}

type ReplaceAreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceAreaContext() *ReplaceAreaContext {
	var p = new(ReplaceAreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceArea
	return p
}

func (*ReplaceAreaContext) IsReplaceAreaContext() {}

func NewReplaceAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceAreaContext {
	var p = new(ReplaceAreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceArea

	return p
}

func (s *ReplaceAreaContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceAreaContext) ReplaceByStatement() IReplaceByStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceByStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplaceByStatementContext)
}

func (s *ReplaceAreaContext) AllCopyStatement() []ICopyStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICopyStatementContext)(nil)).Elem())
	var tst = make([]ICopyStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICopyStatementContext)
		}
	}

	return tst
}

func (s *ReplaceAreaContext) CopyStatement(i int) ICopyStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICopyStatementContext)
}

func (s *ReplaceAreaContext) AllCharData() []ICharDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharDataContext)(nil)).Elem())
	var tst = make([]ICharDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharDataContext)
		}
	}

	return tst
}

func (s *ReplaceAreaContext) CharData(i int) ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *ReplaceAreaContext) ReplaceOffStatement() IReplaceOffStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceOffStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplaceOffStatementContext)
}

func (s *ReplaceAreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceAreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceAreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplaceArea(s)
	}
}

func (s *ReplaceAreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplaceArea(s)
	}
}

func (s *ReplaceAreaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplaceArea(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ReplaceArea() (localctx IReplaceAreaContext) {
	localctx = NewReplaceAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Cobol85PreprocessorParserRULE_replaceArea)

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
	{
		p.SetState(137)
		p.ReplaceByStatement()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(140)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Cobol85PreprocessorParserCOPY:
				{
					p.SetState(138)
					p.CopyStatement()
				}

			case Cobol85PreprocessorParserBY, Cobol85PreprocessorParserIN, Cobol85PreprocessorParserOF, Cobol85PreprocessorParserOFF, Cobol85PreprocessorParserON, Cobol85PreprocessorParserREPLACING, Cobol85PreprocessorParserDOT, Cobol85PreprocessorParserNONNUMERICLITERAL, Cobol85PreprocessorParserIDENTIFIER, Cobol85PreprocessorParserNEWLINE, Cobol85PreprocessorParserTEXT:
				{
					p.SetState(139)
					p.CharData()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(145)
			p.ReplaceOffStatement()
		}

	}

	return localctx
}

// IReplaceByStatementContext is an interface to support dynamic dispatch.
type IReplaceByStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceByStatementContext differentiates from other interfaces.
	IsReplaceByStatementContext()
}

type ReplaceByStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceByStatementContext() *ReplaceByStatementContext {
	var p = new(ReplaceByStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceByStatement
	return p
}

func (*ReplaceByStatementContext) IsReplaceByStatementContext() {}

func NewReplaceByStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceByStatementContext {
	var p = new(ReplaceByStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceByStatement

	return p
}

func (s *ReplaceByStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceByStatementContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserREPLACE, 0)
}

func (s *ReplaceByStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *ReplaceByStatementContext) AllReplaceClause() []IReplaceClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReplaceClauseContext)(nil)).Elem())
	var tst = make([]IReplaceClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReplaceClauseContext)
		}
	}

	return tst
}

func (s *ReplaceByStatementContext) ReplaceClause(i int) IReplaceClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReplaceClauseContext)
}

func (s *ReplaceByStatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *ReplaceByStatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *ReplaceByStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceByStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceByStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplaceByStatement(s)
	}
}

func (s *ReplaceByStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplaceByStatement(s)
	}
}

func (s *ReplaceByStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplaceByStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ReplaceByStatement() (localctx IReplaceByStatementContext) {
	localctx = NewReplaceByStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Cobol85PreprocessorParserRULE_replaceByStatement)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(Cobol85PreprocessorParserREPLACE)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == Cobol85PreprocessorParserNEWLINE {
				{
					p.SetState(149)
					p.Match(Cobol85PreprocessorParserNEWLINE)
				}

				p.SetState(154)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(155)
				p.ReplaceClause()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	{
		p.SetState(160)
		p.Match(Cobol85PreprocessorParserDOT)
	}

	return localctx
}

// IReplaceOffStatementContext is an interface to support dynamic dispatch.
type IReplaceOffStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceOffStatementContext differentiates from other interfaces.
	IsReplaceOffStatementContext()
}

type ReplaceOffStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceOffStatementContext() *ReplaceOffStatementContext {
	var p = new(ReplaceOffStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceOffStatement
	return p
}

func (*ReplaceOffStatementContext) IsReplaceOffStatementContext() {}

func NewReplaceOffStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceOffStatementContext {
	var p = new(ReplaceOffStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceOffStatement

	return p
}

func (s *ReplaceOffStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceOffStatementContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserREPLACE, 0)
}

func (s *ReplaceOffStatementContext) OFF() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserOFF, 0)
}

func (s *ReplaceOffStatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, 0)
}

func (s *ReplaceOffStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceOffStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceOffStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplaceOffStatement(s)
	}
}

func (s *ReplaceOffStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplaceOffStatement(s)
	}
}

func (s *ReplaceOffStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplaceOffStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ReplaceOffStatement() (localctx IReplaceOffStatementContext) {
	localctx = NewReplaceOffStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Cobol85PreprocessorParserRULE_replaceOffStatement)

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
		p.SetState(162)
		p.Match(Cobol85PreprocessorParserREPLACE)
	}
	{
		p.SetState(163)
		p.Match(Cobol85PreprocessorParserOFF)
	}
	{
		p.SetState(164)
		p.Match(Cobol85PreprocessorParserDOT)
	}

	return localctx
}

// IReplaceClauseContext is an interface to support dynamic dispatch.
type IReplaceClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceClauseContext differentiates from other interfaces.
	IsReplaceClauseContext()
}

type ReplaceClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceClauseContext() *ReplaceClauseContext {
	var p = new(ReplaceClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceClause
	return p
}

func (*ReplaceClauseContext) IsReplaceClauseContext() {}

func NewReplaceClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceClauseContext {
	var p = new(ReplaceClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceClause

	return p
}

func (s *ReplaceClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceClauseContext) Replaceable() IReplaceableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplaceableContext)
}

func (s *ReplaceClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserBY, 0)
}

func (s *ReplaceClauseContext) Replacement() IReplacementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplacementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplacementContext)
}

func (s *ReplaceClauseContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *ReplaceClauseContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *ReplaceClauseContext) DirectoryPhrase() IDirectoryPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectoryPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectoryPhraseContext)
}

func (s *ReplaceClauseContext) FamilyPhrase() IFamilyPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFamilyPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFamilyPhraseContext)
}

func (s *ReplaceClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplaceClause(s)
	}
}

func (s *ReplaceClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplaceClause(s)
	}
}

func (s *ReplaceClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplaceClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ReplaceClause() (localctx IReplaceClauseContext) {
	localctx = NewReplaceClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Cobol85PreprocessorParserRULE_replaceClause)
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
		p.SetState(166)
		p.Replaceable()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		{
			p.SetState(167)
			p.Match(Cobol85PreprocessorParserNEWLINE)
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.Match(Cobol85PreprocessorParserBY)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		{
			p.SetState(174)
			p.Match(Cobol85PreprocessorParserNEWLINE)
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(180)
		p.Replacement()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(181)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(187)
			p.DirectoryPhrase()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Cobol85PreprocessorParserNEWLINE {
			{
				p.SetState(190)
				p.Match(Cobol85PreprocessorParserNEWLINE)
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(196)
			p.FamilyPhrase()
		}

	}

	return localctx
}

// IDirectoryPhraseContext is an interface to support dynamic dispatch.
type IDirectoryPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectoryPhraseContext differentiates from other interfaces.
	IsDirectoryPhraseContext()
}

type DirectoryPhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectoryPhraseContext() *DirectoryPhraseContext {
	var p = new(DirectoryPhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_directoryPhrase
	return p
}

func (*DirectoryPhraseContext) IsDirectoryPhraseContext() {}

func NewDirectoryPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectoryPhraseContext {
	var p = new(DirectoryPhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_directoryPhrase

	return p
}

func (s *DirectoryPhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectoryPhraseContext) OF() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserOF, 0)
}

func (s *DirectoryPhraseContext) IN() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserIN, 0)
}

func (s *DirectoryPhraseContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *DirectoryPhraseContext) CobolWord() ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *DirectoryPhraseContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *DirectoryPhraseContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *DirectoryPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectoryPhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectoryPhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterDirectoryPhrase(s)
	}
}

func (s *DirectoryPhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitDirectoryPhrase(s)
	}
}

func (s *DirectoryPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitDirectoryPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) DirectoryPhrase() (localctx IDirectoryPhraseContext) {
	localctx = NewDirectoryPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Cobol85PreprocessorParserRULE_directoryPhrase)
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
	p.SetState(199)
	_la = p.GetTokenStream().LA(1)

	if !(_la == Cobol85PreprocessorParserIN || _la == Cobol85PreprocessorParserOF) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		{
			p.SetState(200)
			p.Match(Cobol85PreprocessorParserNEWLINE)
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Cobol85PreprocessorParserNONNUMERICLITERAL:
		{
			p.SetState(206)
			p.Literal()
		}

	case Cobol85PreprocessorParserIDENTIFIER:
		{
			p.SetState(207)
			p.CobolWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFamilyPhraseContext is an interface to support dynamic dispatch.
type IFamilyPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFamilyPhraseContext differentiates from other interfaces.
	IsFamilyPhraseContext()
}

type FamilyPhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFamilyPhraseContext() *FamilyPhraseContext {
	var p = new(FamilyPhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_familyPhrase
	return p
}

func (*FamilyPhraseContext) IsFamilyPhraseContext() {}

func NewFamilyPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FamilyPhraseContext {
	var p = new(FamilyPhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_familyPhrase

	return p
}

func (s *FamilyPhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *FamilyPhraseContext) ON() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserON, 0)
}

func (s *FamilyPhraseContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *FamilyPhraseContext) CobolWord() ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *FamilyPhraseContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *FamilyPhraseContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *FamilyPhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FamilyPhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FamilyPhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterFamilyPhrase(s)
	}
}

func (s *FamilyPhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitFamilyPhrase(s)
	}
}

func (s *FamilyPhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitFamilyPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) FamilyPhrase() (localctx IFamilyPhraseContext) {
	localctx = NewFamilyPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Cobol85PreprocessorParserRULE_familyPhrase)
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
		p.Match(Cobol85PreprocessorParserON)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Cobol85PreprocessorParserNEWLINE {
		{
			p.SetState(211)
			p.Match(Cobol85PreprocessorParserNEWLINE)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Cobol85PreprocessorParserNONNUMERICLITERAL:
		{
			p.SetState(217)
			p.Literal()
		}

	case Cobol85PreprocessorParserIDENTIFIER:
		{
			p.SetState(218)
			p.CobolWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReplaceableContext is an interface to support dynamic dispatch.
type IReplaceableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceableContext differentiates from other interfaces.
	IsReplaceableContext()
}

type ReplaceableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceableContext() *ReplaceableContext {
	var p = new(ReplaceableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceable
	return p
}

func (*ReplaceableContext) IsReplaceableContext() {}

func NewReplaceableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceableContext {
	var p = new(ReplaceableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replaceable

	return p
}

func (s *ReplaceableContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceableContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ReplaceableContext) CobolWord() ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *ReplaceableContext) PseudoText() IPseudoTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoTextContext)
}

func (s *ReplaceableContext) CharDataLine() ICharDataLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataLineContext)
}

func (s *ReplaceableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplaceable(s)
	}
}

func (s *ReplaceableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplaceable(s)
	}
}

func (s *ReplaceableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplaceable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) Replaceable() (localctx IReplaceableContext) {
	localctx = NewReplaceableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Cobol85PreprocessorParserRULE_replaceable)

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

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.CobolWord()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.PseudoText()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
			p.CharDataLine()
		}

	}

	return localctx
}

// IReplacementContext is an interface to support dynamic dispatch.
type IReplacementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplacementContext differentiates from other interfaces.
	IsReplacementContext()
}

type ReplacementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplacementContext() *ReplacementContext {
	var p = new(ReplacementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_replacement
	return p
}

func (*ReplacementContext) IsReplacementContext() {}

func NewReplacementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplacementContext {
	var p = new(ReplacementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_replacement

	return p
}

func (s *ReplacementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplacementContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ReplacementContext) CobolWord() ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *ReplacementContext) PseudoText() IPseudoTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoTextContext)
}

func (s *ReplacementContext) CharDataLine() ICharDataLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataLineContext)
}

func (s *ReplacementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplacementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplacementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterReplacement(s)
	}
}

func (s *ReplacementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitReplacement(s)
	}
}

func (s *ReplacementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitReplacement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) Replacement() (localctx IReplacementContext) {
	localctx = NewReplacementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Cobol85PreprocessorParserRULE_replacement)

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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.CobolWord()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.PseudoText()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(230)
			p.CharDataLine()
		}

	}

	return localctx
}

// IControlSpacingStatementContext is an interface to support dynamic dispatch.
type IControlSpacingStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlSpacingStatementContext differentiates from other interfaces.
	IsControlSpacingStatementContext()
}

type ControlSpacingStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlSpacingStatementContext() *ControlSpacingStatementContext {
	var p = new(ControlSpacingStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_controlSpacingStatement
	return p
}

func (*ControlSpacingStatementContext) IsControlSpacingStatementContext() {}

func NewControlSpacingStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlSpacingStatementContext {
	var p = new(ControlSpacingStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_controlSpacingStatement

	return p
}

func (s *ControlSpacingStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlSpacingStatementContext) SKIP1() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSKIP1, 0)
}

func (s *ControlSpacingStatementContext) SKIP2() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSKIP2, 0)
}

func (s *ControlSpacingStatementContext) SKIP3() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserSKIP3, 0)
}

func (s *ControlSpacingStatementContext) EJECT() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserEJECT, 0)
}

func (s *ControlSpacingStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlSpacingStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlSpacingStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterControlSpacingStatement(s)
	}
}

func (s *ControlSpacingStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitControlSpacingStatement(s)
	}
}

func (s *ControlSpacingStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitControlSpacingStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) ControlSpacingStatement() (localctx IControlSpacingStatementContext) {
	localctx = NewControlSpacingStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Cobol85PreprocessorParserRULE_controlSpacingStatement)
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
	p.SetState(233)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Cobol85PreprocessorParserEJECT)|(1<<Cobol85PreprocessorParserSKIP1)|(1<<Cobol85PreprocessorParserSKIP2)|(1<<Cobol85PreprocessorParserSKIP3))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ICobolWordContext is an interface to support dynamic dispatch.
type ICobolWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCobolWordContext differentiates from other interfaces.
	IsCobolWordContext()
}

type CobolWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCobolWordContext() *CobolWordContext {
	var p = new(CobolWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_cobolWord
	return p
}

func (*CobolWordContext) IsCobolWordContext() {}

func NewCobolWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CobolWordContext {
	var p = new(CobolWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_cobolWord

	return p
}

func (s *CobolWordContext) GetParser() antlr.Parser { return s.parser }

func (s *CobolWordContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserIDENTIFIER, 0)
}

func (s *CobolWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CobolWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CobolWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCobolWord(s)
	}
}

func (s *CobolWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCobolWord(s)
	}
}

func (s *CobolWordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCobolWord(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CobolWord() (localctx ICobolWordContext) {
	localctx = NewCobolWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Cobol85PreprocessorParserRULE_cobolWord)

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
		p.Match(Cobol85PreprocessorParserIDENTIFIER)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NONNUMERICLITERAL() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNONNUMERICLITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Cobol85PreprocessorParserRULE_literal)

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
		p.SetState(237)
		p.Match(Cobol85PreprocessorParserNONNUMERICLITERAL)
	}

	return localctx
}

// IPseudoTextContext is an interface to support dynamic dispatch.
type IPseudoTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudoTextContext differentiates from other interfaces.
	IsPseudoTextContext()
}

type PseudoTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudoTextContext() *PseudoTextContext {
	var p = new(PseudoTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_pseudoText
	return p
}

func (*PseudoTextContext) IsPseudoTextContext() {}

func NewPseudoTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PseudoTextContext {
	var p = new(PseudoTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_pseudoText

	return p
}

func (s *PseudoTextContext) GetParser() antlr.Parser { return s.parser }

func (s *PseudoTextContext) AllDOUBLEEQUALCHAR() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserDOUBLEEQUALCHAR)
}

func (s *PseudoTextContext) DOUBLEEQUALCHAR(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOUBLEEQUALCHAR, i)
}

func (s *PseudoTextContext) CharData() ICharDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharDataContext)
}

func (s *PseudoTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PseudoTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PseudoTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterPseudoText(s)
	}
}

func (s *PseudoTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitPseudoText(s)
	}
}

func (s *PseudoTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitPseudoText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) PseudoText() (localctx IPseudoTextContext) {
	localctx = NewPseudoTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Cobol85PreprocessorParserRULE_pseudoText)
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
		p.SetState(239)
		p.Match(Cobol85PreprocessorParserDOUBLEEQUALCHAR)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Cobol85PreprocessorParserBY)|(1<<Cobol85PreprocessorParserIN)|(1<<Cobol85PreprocessorParserOF)|(1<<Cobol85PreprocessorParserOFF)|(1<<Cobol85PreprocessorParserON)|(1<<Cobol85PreprocessorParserREPLACING)|(1<<Cobol85PreprocessorParserDOT)|(1<<Cobol85PreprocessorParserNONNUMERICLITERAL)|(1<<Cobol85PreprocessorParserIDENTIFIER)|(1<<Cobol85PreprocessorParserNEWLINE)|(1<<Cobol85PreprocessorParserTEXT))) != 0 {
		{
			p.SetState(240)
			p.CharData()
		}

	}
	{
		p.SetState(243)
		p.Match(Cobol85PreprocessorParserDOUBLEEQUALCHAR)
	}

	return localctx
}

// ICharDataContext is an interface to support dynamic dispatch.
type ICharDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharDataContext differentiates from other interfaces.
	IsCharDataContext()
}

type CharDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharDataContext() *CharDataContext {
	var p = new(CharDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_charData
	return p
}

func (*CharDataContext) IsCharDataContext() {}

func NewCharDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharDataContext {
	var p = new(CharDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_charData

	return p
}

func (s *CharDataContext) GetParser() antlr.Parser { return s.parser }

func (s *CharDataContext) AllCharDataLine() []ICharDataLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharDataLineContext)(nil)).Elem())
	var tst = make([]ICharDataLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharDataLineContext)
		}
	}

	return tst
}

func (s *CharDataContext) CharDataLine(i int) ICharDataLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharDataLineContext)
}

func (s *CharDataContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserNEWLINE)
}

func (s *CharDataContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserNEWLINE, i)
}

func (s *CharDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCharData(s)
	}
}

func (s *CharDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCharData(s)
	}
}

func (s *CharDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCharData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CharData() (localctx ICharDataContext) {
	localctx = NewCharDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Cobol85PreprocessorParserRULE_charData)

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
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(247)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Cobol85PreprocessorParserBY, Cobol85PreprocessorParserIN, Cobol85PreprocessorParserOF, Cobol85PreprocessorParserOFF, Cobol85PreprocessorParserON, Cobol85PreprocessorParserREPLACING, Cobol85PreprocessorParserDOT, Cobol85PreprocessorParserNONNUMERICLITERAL, Cobol85PreprocessorParserIDENTIFIER, Cobol85PreprocessorParserTEXT:
				{
					p.SetState(245)
					p.CharDataLine()
				}

			case Cobol85PreprocessorParserNEWLINE:
				{
					p.SetState(246)
					p.Match(Cobol85PreprocessorParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// ICharDataLineContext is an interface to support dynamic dispatch.
type ICharDataLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharDataLineContext differentiates from other interfaces.
	IsCharDataLineContext()
}

type CharDataLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharDataLineContext() *CharDataLineContext {
	var p = new(CharDataLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_charDataLine
	return p
}

func (*CharDataLineContext) IsCharDataLineContext() {}

func NewCharDataLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharDataLineContext {
	var p = new(CharDataLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_charDataLine

	return p
}

func (s *CharDataLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CharDataLineContext) AllCharDataKeyword() []ICharDataKeywordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharDataKeywordContext)(nil)).Elem())
	var tst = make([]ICharDataKeywordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharDataKeywordContext)
		}
	}

	return tst
}

func (s *CharDataLineContext) CharDataKeyword(i int) ICharDataKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharDataKeywordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharDataKeywordContext)
}

func (s *CharDataLineContext) AllCobolWord() []ICobolWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICobolWordContext)(nil)).Elem())
	var tst = make([]ICobolWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICobolWordContext)
		}
	}

	return tst
}

func (s *CharDataLineContext) CobolWord(i int) ICobolWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICobolWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICobolWordContext)
}

func (s *CharDataLineContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *CharDataLineContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *CharDataLineContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserTEXT)
}

func (s *CharDataLineContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserTEXT, i)
}

func (s *CharDataLineContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Cobol85PreprocessorParserDOT)
}

func (s *CharDataLineContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserDOT, i)
}

func (s *CharDataLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharDataLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharDataLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCharDataLine(s)
	}
}

func (s *CharDataLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCharDataLine(s)
	}
}

func (s *CharDataLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCharDataLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CharDataLine() (localctx ICharDataLineContext) {
	localctx = NewCharDataLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Cobol85PreprocessorParserRULE_charDataLine)

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
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(256)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Cobol85PreprocessorParserBY, Cobol85PreprocessorParserIN, Cobol85PreprocessorParserOF, Cobol85PreprocessorParserOFF, Cobol85PreprocessorParserON, Cobol85PreprocessorParserREPLACING:
				{
					p.SetState(251)
					p.CharDataKeyword()
				}

			case Cobol85PreprocessorParserIDENTIFIER:
				{
					p.SetState(252)
					p.CobolWord()
				}

			case Cobol85PreprocessorParserNONNUMERICLITERAL:
				{
					p.SetState(253)
					p.Literal()
				}

			case Cobol85PreprocessorParserTEXT:
				{
					p.SetState(254)
					p.Match(Cobol85PreprocessorParserTEXT)
				}

			case Cobol85PreprocessorParserDOT:
				{
					p.SetState(255)
					p.Match(Cobol85PreprocessorParserDOT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// ICharDataKeywordContext is an interface to support dynamic dispatch.
type ICharDataKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharDataKeywordContext differentiates from other interfaces.
	IsCharDataKeywordContext()
}

type CharDataKeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharDataKeywordContext() *CharDataKeywordContext {
	var p = new(CharDataKeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Cobol85PreprocessorParserRULE_charDataKeyword
	return p
}

func (*CharDataKeywordContext) IsCharDataKeywordContext() {}

func NewCharDataKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharDataKeywordContext {
	var p = new(CharDataKeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Cobol85PreprocessorParserRULE_charDataKeyword

	return p
}

func (s *CharDataKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *CharDataKeywordContext) BY() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserBY, 0)
}

func (s *CharDataKeywordContext) IN() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserIN, 0)
}

func (s *CharDataKeywordContext) OF() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserOF, 0)
}

func (s *CharDataKeywordContext) OFF() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserOFF, 0)
}

func (s *CharDataKeywordContext) ON() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserON, 0)
}

func (s *CharDataKeywordContext) REPLACING() antlr.TerminalNode {
	return s.GetToken(Cobol85PreprocessorParserREPLACING, 0)
}

func (s *CharDataKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharDataKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharDataKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.EnterCharDataKeyword(s)
	}
}

func (s *CharDataKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Cobol85PreprocessorListener); ok {
		listenerT.ExitCharDataKeyword(s)
	}
}

func (s *CharDataKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Cobol85PreprocessorVisitor:
		return t.VisitCharDataKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Cobol85PreprocessorParser) CharDataKeyword() (localctx ICharDataKeywordContext) {
	localctx = NewCharDataKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Cobol85PreprocessorParserRULE_charDataKeyword)
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
	p.SetState(260)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Cobol85PreprocessorParserBY)|(1<<Cobol85PreprocessorParserIN)|(1<<Cobol85PreprocessorParserOF)|(1<<Cobol85PreprocessorParserOFF)|(1<<Cobol85PreprocessorParserON)|(1<<Cobol85PreprocessorParserREPLACING))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
