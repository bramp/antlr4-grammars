// Generated from agc.g4 by ANTLR 4.7.

package agc // agc
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 198, 216,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 6, 2, 52, 10, 2, 13, 2, 14, 2, 53, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 61, 10, 3, 3, 4, 5, 4, 64, 10, 4, 3, 4,
	3, 4, 3, 5, 5, 5, 69, 10, 5, 3, 5, 3, 5, 3, 5, 3, 6, 5, 6, 75, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 82, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7,
	6, 88, 10, 6, 12, 6, 14, 6, 91, 11, 6, 3, 6, 3, 6, 3, 7, 5, 7, 96, 10,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 101, 10, 7, 3, 7, 7, 7, 104, 10, 7, 12, 7, 14,
	7, 107, 11, 7, 3, 7, 3, 7, 3, 7, 5, 7, 112, 10, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 5, 8, 118, 10, 8, 3, 8, 3, 8, 5, 8, 122, 10, 8, 3, 8, 7, 8, 125, 10,
	8, 12, 8, 14, 8, 128, 11, 8, 3, 8, 3, 8, 3, 8, 5, 8, 133, 10, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 141, 10, 9, 3, 10, 3, 10, 3, 10, 7,
	10, 146, 10, 10, 12, 10, 14, 10, 149, 11, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	154, 10, 10, 3, 11, 3, 11, 3, 12, 5, 12, 159, 10, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 171, 10, 15,
	3, 16, 3, 16, 3, 16, 7, 16, 176, 10, 16, 12, 16, 14, 16, 179, 11, 16, 3,
	17, 3, 17, 3, 17, 7, 17, 184, 10, 17, 12, 17, 14, 17, 187, 11, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 194, 10, 18, 3, 19, 3, 19, 3, 20, 5,
	20, 199, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5, 22,
	208, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 2, 9, 3, 2, 4, 5, 3, 2, 190, 191, 3, 2, 192, 193, 3,
	2, 6, 54, 3, 2, 55, 56, 3, 2, 57, 84, 3, 2, 85, 185, 2, 223, 2, 51, 3,
	2, 2, 2, 4, 60, 3, 2, 2, 2, 6, 63, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 74,
	3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 115, 3, 2, 2, 2, 16, 136, 3, 2, 2,
	2, 18, 147, 3, 2, 2, 2, 20, 155, 3, 2, 2, 2, 22, 158, 3, 2, 2, 2, 24, 162,
	3, 2, 2, 2, 26, 164, 3, 2, 2, 2, 28, 170, 3, 2, 2, 2, 30, 172, 3, 2, 2,
	2, 32, 180, 3, 2, 2, 2, 34, 193, 3, 2, 2, 2, 36, 195, 3, 2, 2, 2, 38, 198,
	3, 2, 2, 2, 40, 202, 3, 2, 2, 2, 42, 207, 3, 2, 2, 2, 44, 209, 3, 2, 2,
	2, 46, 211, 3, 2, 2, 2, 48, 213, 3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50,
	3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2,
	54, 3, 3, 2, 2, 2, 55, 61, 5, 8, 5, 2, 56, 61, 5, 6, 4, 2, 57, 61, 5, 10,
	6, 2, 58, 61, 5, 12, 7, 2, 59, 61, 5, 14, 8, 2, 60, 55, 3, 2, 2, 2, 60,
	56, 3, 2, 2, 2, 60, 57, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2,
	2, 61, 5, 3, 2, 2, 2, 62, 64, 5, 26, 14, 2, 63, 62, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 5, 22, 12, 2, 66, 7, 3, 2, 2, 2,
	67, 69, 5, 20, 11, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 3,
	2, 2, 2, 70, 71, 5, 24, 13, 2, 71, 72, 5, 22, 12, 2, 72, 9, 3, 2, 2, 2,
	73, 75, 5, 26, 14, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 3,
	2, 2, 2, 76, 77, 5, 20, 11, 2, 77, 81, 5, 16, 9, 2, 78, 79, 5, 22, 12,
	2, 79, 80, 5, 8, 5, 2, 80, 82, 3, 2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 89, 5, 18, 10, 2, 84, 85, 5, 22, 12,
	2, 85, 86, 5, 18, 10, 2, 86, 88, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 88, 91,
	3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2,
	91, 89, 3, 2, 2, 2, 92, 93, 5, 22, 12, 2, 93, 11, 3, 2, 2, 2, 94, 96, 5,
	28, 15, 2, 95, 94, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 98, 5, 20, 11, 2, 98, 105, 7, 3, 2, 2, 99, 101, 5, 20, 11, 2, 100,
	99, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 5,
	30, 16, 2, 103, 100, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2,
	2, 2, 105, 106, 3, 2, 2, 2, 106, 111, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2,
	108, 109, 5, 20, 11, 2, 109, 110, 5, 24, 13, 2, 110, 112, 3, 2, 2, 2, 111,
	108, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114,
	5, 22, 12, 2, 114, 13, 3, 2, 2, 2, 115, 117, 5, 28, 15, 2, 116, 118, 5,
	20, 11, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2,
	2, 2, 119, 126, 9, 2, 2, 2, 120, 122, 5, 20, 11, 2, 121, 120, 3, 2, 2,
	2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 5, 30, 16, 2,
	124, 121, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126,
	127, 3, 2, 2, 2, 127, 132, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130,
	5, 20, 11, 2, 130, 131, 5, 24, 13, 2, 131, 133, 3, 2, 2, 2, 132, 129, 3,
	2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 5, 22, 12,
	2, 135, 15, 3, 2, 2, 2, 136, 140, 5, 42, 22, 2, 137, 138, 5, 20, 11, 2,
	138, 139, 5, 42, 22, 2, 139, 141, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140,
	141, 3, 2, 2, 2, 141, 17, 3, 2, 2, 2, 142, 143, 5, 20, 11, 2, 143, 144,
	5, 30, 16, 2, 144, 146, 3, 2, 2, 2, 145, 142, 3, 2, 2, 2, 146, 149, 3,
	2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 153, 3, 2, 2,
	2, 149, 147, 3, 2, 2, 2, 150, 151, 5, 20, 11, 2, 151, 152, 5, 24, 13, 2,
	152, 154, 3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154,
	19, 3, 2, 2, 2, 155, 156, 7, 198, 2, 2, 156, 21, 3, 2, 2, 2, 157, 159,
	7, 198, 2, 2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3,
	2, 2, 2, 160, 161, 7, 197, 2, 2, 161, 23, 3, 2, 2, 2, 162, 163, 7, 189,
	2, 2, 163, 25, 3, 2, 2, 2, 164, 165, 7, 186, 2, 2, 165, 27, 3, 2, 2, 2,
	166, 171, 7, 186, 2, 2, 167, 168, 7, 195, 2, 2, 168, 169, 7, 186, 2, 2,
	169, 171, 7, 196, 2, 2, 170, 166, 3, 2, 2, 2, 170, 167, 3, 2, 2, 2, 171,
	29, 3, 2, 2, 2, 172, 177, 5, 32, 17, 2, 173, 174, 9, 3, 2, 2, 174, 176,
	5, 32, 17, 2, 175, 173, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3,
	2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 31, 3, 2, 2, 2, 179, 177, 3, 2, 2,
	2, 180, 185, 5, 34, 18, 2, 181, 182, 9, 4, 2, 2, 182, 184, 5, 34, 18, 2,
	183, 181, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 33, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 194, 5,
	36, 19, 2, 189, 194, 5, 38, 20, 2, 190, 194, 5, 28, 15, 2, 191, 194, 5,
	26, 14, 2, 192, 194, 5, 40, 21, 2, 193, 188, 3, 2, 2, 2, 193, 189, 3, 2,
	2, 2, 193, 190, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2, 2,
	194, 35, 3, 2, 2, 2, 195, 196, 7, 187, 2, 2, 196, 37, 3, 2, 2, 2, 197,
	199, 9, 3, 2, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200,
	3, 2, 2, 2, 200, 201, 7, 188, 2, 2, 201, 39, 3, 2, 2, 2, 202, 203, 9, 5,
	2, 2, 203, 41, 3, 2, 2, 2, 204, 208, 5, 48, 25, 2, 205, 208, 5, 46, 24,
	2, 206, 208, 5, 44, 23, 2, 207, 204, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2,
	207, 206, 3, 2, 2, 2, 208, 43, 3, 2, 2, 2, 209, 210, 9, 6, 2, 2, 210, 45,
	3, 2, 2, 2, 211, 212, 9, 7, 2, 2, 212, 47, 3, 2, 2, 2, 213, 214, 9, 8,
	2, 2, 214, 49, 3, 2, 2, 2, 27, 53, 60, 63, 68, 74, 81, 89, 95, 100, 105,
	111, 117, 121, 126, 132, 140, 147, 153, 158, 170, 177, 185, 193, 198, 207,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'ERASE'", "'='", "'EQUALS'", "'A'", "'L'", "'Q'", "'EB'", "'FB'",
	"'Z'", "'BB'", "'ARUPT'", "'LRUPT'", "'QRUPT'", "'BBRUPT'", "'BRUPT'",
	"'CYR'", "'SR'", "'CYL'", "'EDOP'", "'TIME2'", "'TIME1'", "'TIME3'", "'TIME4'",
	"'TIME5'", "'TIME6'", "'CDUX'", "'CDUY'", "'CDUZ'", "'OPTY'", "'OPTX'",
	"'PIPAX'", "'PIPAY'", "'PIPAZ'", "'Q-RHCCTR'", "'RHCP'", "'P-RHCCTR'",
	"'RHCY'", "'R-RHCCTR'", "'RHCR'", "'INLINK'", "'RNRAD'", "'GYROCTR'", "'GYROCMD'",
	"'CDUXCMD'", "'CDUYCMD'", "'CDUZCMD'", "'OPTYCMD'", "'OPTXCMD'", "'THRUST'",
	"'LEMONM'", "'OUTLINK'", "'ALTM'", "'AXT,1'", "'AXT,2'", "'1DNADR'", "'2DNADR'",
	"'3DNADR'", "'4DNADR'", "'5DNADR'", "'6DNADR'", "'DNCHAN'", "'DNPTR'",
	"'-1DNADR'", "'-2DNADR'", "'-3DNADR'", "'-4DNADR'", "'-5DNADR'", "'-6DNADR'",
	"'-DNCHAN'", "'-DNPTR'", "'2DEC'", "'2DEC*'", "'2FCADR'", "'BANK'", "'BLOCK'",
	"'BNKSUM'", "'COUNT'", "'COUNT*'", "'DEC'", "'OCT'", "'SETLOC'", "'SUBRO'",
	"'TC'", "'TCR'", "'CCS'", "'TCF'", "'DAS'", "'LXCH'", "'INCR'", "'AD'",
	"'ADS'", "'CA'", "'CS'", "'INDEX'", "'DXCH'", "'TS'", "'XCH'", "'MASK'",
	"'MSK'", "'READ'", "'WRITE'", "'RAND'", "'WAND'", "'ROR'", "'WOR'", "'RXOR'",
	"'EDRUPT'", "'BZF'", "'MSU'", "'QXCH'", "'AUG'", "'DIM'", "'DCA'", "'DCS'",
	"'SU'", "'BZMF'", "'MP'", "'XXALQ'", "'XLQ'", "'RETURN'", "'RELINT'", "'INHINT'",
	"'EXTEND'", "'NOOP'", "'DDOUBL'", "'DTCF'", "'COM'", "'ZL'", "'RESUME'",
	"'DTCB'", "'OVSK'", "'TCAA'", "'DOUBLE'", "'ZQ'", "'DCOM'", "'SQUARE'",
	"'PINC'", "'PCDU'", "'MINC'", "'MCDU'", "'DINC'", "'SHINC'", "'SHANC'",
	"'INOTRD'", "'INOTLD'", "'FETCH'", "'STORE'", "'GOJ'", "'TCSAJ'", "'CAF'",
	"'CAE'", "'CADR'", "'DMOVE'", "'VMOVE'", "'SMOVE'", "'DSU'", "'RTB'", "'ITC'",
	"'NOLOD'", "'EXIT'", "'BPL'", "'SIN'", "'COS'", "'CAD'", "'TEST'", "'VXSC'",
	"'DAD'", "'VXV'", "'VAD'", "'DMP'", "'BOV'", "'UNIT'", "'OCTAL'", "'ADRES'",
	"'ABVAL'", "'COMP'", "'DV'", "'NDX'", "'POUT'", "'MOUT'", "'ZOUT'", "'LODON'",
	"'TSLT'", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "','", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "LABEL", "INTE", "DECIMAL", "COMMENT", "PLUS", "MINUS",
	"TIMES", "DIV", "COMMA", "LPAREN", "RPAREN", "EOL", "WS",
}

var ruleNames = []string{
	"prog", "line", "blank_line", "comment_line", "instruction_line", "erase_line",
	"assignment_line", "opcodes", "argument", "ws", "eol", "comment", "label",
	"variable", "expression", "multiplyingExpression", "atom", "inte", "decimal",
	"register", "opcode", "axt_opcode", "pseudo_opcode", "standard_opcode",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type agcParser struct {
	*antlr.BaseParser
}

func NewagcParser(input antlr.TokenStream) *agcParser {
	this := new(agcParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "agc.g4"

	return this
}

// agcParser tokens.
const (
	agcParserEOF     = antlr.TokenEOF
	agcParserT__0    = 1
	agcParserT__1    = 2
	agcParserT__2    = 3
	agcParserT__3    = 4
	agcParserT__4    = 5
	agcParserT__5    = 6
	agcParserT__6    = 7
	agcParserT__7    = 8
	agcParserT__8    = 9
	agcParserT__9    = 10
	agcParserT__10   = 11
	agcParserT__11   = 12
	agcParserT__12   = 13
	agcParserT__13   = 14
	agcParserT__14   = 15
	agcParserT__15   = 16
	agcParserT__16   = 17
	agcParserT__17   = 18
	agcParserT__18   = 19
	agcParserT__19   = 20
	agcParserT__20   = 21
	agcParserT__21   = 22
	agcParserT__22   = 23
	agcParserT__23   = 24
	agcParserT__24   = 25
	agcParserT__25   = 26
	agcParserT__26   = 27
	agcParserT__27   = 28
	agcParserT__28   = 29
	agcParserT__29   = 30
	agcParserT__30   = 31
	agcParserT__31   = 32
	agcParserT__32   = 33
	agcParserT__33   = 34
	agcParserT__34   = 35
	agcParserT__35   = 36
	agcParserT__36   = 37
	agcParserT__37   = 38
	agcParserT__38   = 39
	agcParserT__39   = 40
	agcParserT__40   = 41
	agcParserT__41   = 42
	agcParserT__42   = 43
	agcParserT__43   = 44
	agcParserT__44   = 45
	agcParserT__45   = 46
	agcParserT__46   = 47
	agcParserT__47   = 48
	agcParserT__48   = 49
	agcParserT__49   = 50
	agcParserT__50   = 51
	agcParserT__51   = 52
	agcParserT__52   = 53
	agcParserT__53   = 54
	agcParserT__54   = 55
	agcParserT__55   = 56
	agcParserT__56   = 57
	agcParserT__57   = 58
	agcParserT__58   = 59
	agcParserT__59   = 60
	agcParserT__60   = 61
	agcParserT__61   = 62
	agcParserT__62   = 63
	agcParserT__63   = 64
	agcParserT__64   = 65
	agcParserT__65   = 66
	agcParserT__66   = 67
	agcParserT__67   = 68
	agcParserT__68   = 69
	agcParserT__69   = 70
	agcParserT__70   = 71
	agcParserT__71   = 72
	agcParserT__72   = 73
	agcParserT__73   = 74
	agcParserT__74   = 75
	agcParserT__75   = 76
	agcParserT__76   = 77
	agcParserT__77   = 78
	agcParserT__78   = 79
	agcParserT__79   = 80
	agcParserT__80   = 81
	agcParserT__81   = 82
	agcParserT__82   = 83
	agcParserT__83   = 84
	agcParserT__84   = 85
	agcParserT__85   = 86
	agcParserT__86   = 87
	agcParserT__87   = 88
	agcParserT__88   = 89
	agcParserT__89   = 90
	agcParserT__90   = 91
	agcParserT__91   = 92
	agcParserT__92   = 93
	agcParserT__93   = 94
	agcParserT__94   = 95
	agcParserT__95   = 96
	agcParserT__96   = 97
	agcParserT__97   = 98
	agcParserT__98   = 99
	agcParserT__99   = 100
	agcParserT__100  = 101
	agcParserT__101  = 102
	agcParserT__102  = 103
	agcParserT__103  = 104
	agcParserT__104  = 105
	agcParserT__105  = 106
	agcParserT__106  = 107
	agcParserT__107  = 108
	agcParserT__108  = 109
	agcParserT__109  = 110
	agcParserT__110  = 111
	agcParserT__111  = 112
	agcParserT__112  = 113
	agcParserT__113  = 114
	agcParserT__114  = 115
	agcParserT__115  = 116
	agcParserT__116  = 117
	agcParserT__117  = 118
	agcParserT__118  = 119
	agcParserT__119  = 120
	agcParserT__120  = 121
	agcParserT__121  = 122
	agcParserT__122  = 123
	agcParserT__123  = 124
	agcParserT__124  = 125
	agcParserT__125  = 126
	agcParserT__126  = 127
	agcParserT__127  = 128
	agcParserT__128  = 129
	agcParserT__129  = 130
	agcParserT__130  = 131
	agcParserT__131  = 132
	agcParserT__132  = 133
	agcParserT__133  = 134
	agcParserT__134  = 135
	agcParserT__135  = 136
	agcParserT__136  = 137
	agcParserT__137  = 138
	agcParserT__138  = 139
	agcParserT__139  = 140
	agcParserT__140  = 141
	agcParserT__141  = 142
	agcParserT__142  = 143
	agcParserT__143  = 144
	agcParserT__144  = 145
	agcParserT__145  = 146
	agcParserT__146  = 147
	agcParserT__147  = 148
	agcParserT__148  = 149
	agcParserT__149  = 150
	agcParserT__150  = 151
	agcParserT__151  = 152
	agcParserT__152  = 153
	agcParserT__153  = 154
	agcParserT__154  = 155
	agcParserT__155  = 156
	agcParserT__156  = 157
	agcParserT__157  = 158
	agcParserT__158  = 159
	agcParserT__159  = 160
	agcParserT__160  = 161
	agcParserT__161  = 162
	agcParserT__162  = 163
	agcParserT__163  = 164
	agcParserT__164  = 165
	agcParserT__165  = 166
	agcParserT__166  = 167
	agcParserT__167  = 168
	agcParserT__168  = 169
	agcParserT__169  = 170
	agcParserT__170  = 171
	agcParserT__171  = 172
	agcParserT__172  = 173
	agcParserT__173  = 174
	agcParserT__174  = 175
	agcParserT__175  = 176
	agcParserT__176  = 177
	agcParserT__177  = 178
	agcParserT__178  = 179
	agcParserT__179  = 180
	agcParserT__180  = 181
	agcParserT__181  = 182
	agcParserT__182  = 183
	agcParserLABEL   = 184
	agcParserINTE    = 185
	agcParserDECIMAL = 186
	agcParserCOMMENT = 187
	agcParserPLUS    = 188
	agcParserMINUS   = 189
	agcParserTIMES   = 190
	agcParserDIV     = 191
	agcParserCOMMA   = 192
	agcParserLPAREN  = 193
	agcParserRPAREN  = 194
	agcParserEOL     = 195
	agcParserWS      = 196
)

// agcParser rules.
const (
	agcParserRULE_prog                  = 0
	agcParserRULE_line                  = 1
	agcParserRULE_blank_line            = 2
	agcParserRULE_comment_line          = 3
	agcParserRULE_instruction_line      = 4
	agcParserRULE_erase_line            = 5
	agcParserRULE_assignment_line       = 6
	agcParserRULE_opcodes               = 7
	agcParserRULE_argument              = 8
	agcParserRULE_ws                    = 9
	agcParserRULE_eol                   = 10
	agcParserRULE_comment               = 11
	agcParserRULE_label                 = 12
	agcParserRULE_variable              = 13
	agcParserRULE_expression            = 14
	agcParserRULE_multiplyingExpression = 15
	agcParserRULE_atom                  = 16
	agcParserRULE_inte                  = 17
	agcParserRULE_decimal               = 18
	agcParserRULE_register              = 19
	agcParserRULE_opcode                = 20
	agcParserRULE_axt_opcode            = 21
	agcParserRULE_pseudo_opcode         = 22
	agcParserRULE_standard_opcode       = 23
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *agcParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, agcParserRULE_prog)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-184)&-(0x1f+1)) == 0 && ((1<<uint((_la-184)))&((1<<(agcParserLABEL-184))|(1<<(agcParserCOMMENT-184))|(1<<(agcParserLPAREN-184))|(1<<(agcParserEOL-184))|(1<<(agcParserWS-184)))) != 0) {
		{
			p.SetState(48)
			p.Line()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Comment_line() IComment_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComment_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComment_lineContext)
}

func (s *LineContext) Blank_line() IBlank_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlank_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlank_lineContext)
}

func (s *LineContext) Instruction_line() IInstruction_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruction_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruction_lineContext)
}

func (s *LineContext) Erase_line() IErase_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErase_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErase_lineContext)
}

func (s *LineContext) Assignment_line() IAssignment_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignment_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignment_lineContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *agcParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, agcParserRULE_line)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Comment_line()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Blank_line()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Instruction_line()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Erase_line()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Assignment_line()
		}

	}

	return localctx
}

// IBlank_lineContext is an interface to support dynamic dispatch.
type IBlank_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlank_lineContext differentiates from other interfaces.
	IsBlank_lineContext()
}

type Blank_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlank_lineContext() *Blank_lineContext {
	var p = new(Blank_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_blank_line
	return p
}

func (*Blank_lineContext) IsBlank_lineContext() {}

func NewBlank_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Blank_lineContext {
	var p = new(Blank_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_blank_line

	return p
}

func (s *Blank_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Blank_lineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Blank_lineContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *Blank_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Blank_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Blank_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterBlank_line(s)
	}
}

func (s *Blank_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitBlank_line(s)
	}
}

func (p *agcParser) Blank_line() (localctx IBlank_lineContext) {
	localctx = NewBlank_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, agcParserRULE_blank_line)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserLABEL {
		{
			p.SetState(60)
			p.Label()
		}

	}
	{
		p.SetState(63)
		p.Eol()
	}

	return localctx
}

// IComment_lineContext is an interface to support dynamic dispatch.
type IComment_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComment_lineContext differentiates from other interfaces.
	IsComment_lineContext()
}

type Comment_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_lineContext() *Comment_lineContext {
	var p = new(Comment_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_comment_line
	return p
}

func (*Comment_lineContext) IsComment_lineContext() {}

func NewComment_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_lineContext {
	var p = new(Comment_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_comment_line

	return p
}

func (s *Comment_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_lineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Comment_lineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Comment_lineContext) Ws() IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *Comment_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterComment_line(s)
	}
}

func (s *Comment_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitComment_line(s)
	}
}

func (p *agcParser) Comment_line() (localctx IComment_lineContext) {
	localctx = NewComment_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, agcParserRULE_comment_line)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserWS {
		{
			p.SetState(65)
			p.Ws()
		}

	}
	{
		p.SetState(68)
		p.Comment()
	}
	{
		p.SetState(69)
		p.Eol()
	}

	return localctx
}

// IInstruction_lineContext is an interface to support dynamic dispatch.
type IInstruction_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruction_lineContext differentiates from other interfaces.
	IsInstruction_lineContext()
}

type Instruction_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruction_lineContext() *Instruction_lineContext {
	var p = new(Instruction_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_instruction_line
	return p
}

func (*Instruction_lineContext) IsInstruction_lineContext() {}

func NewInstruction_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruction_lineContext {
	var p = new(Instruction_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_instruction_line

	return p
}

func (s *Instruction_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruction_lineContext) Ws() IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *Instruction_lineContext) Opcodes() IOpcodesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodesContext)
}

func (s *Instruction_lineContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *Instruction_lineContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *Instruction_lineContext) AllEol() []IEolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEolContext)(nil)).Elem())
	var tst = make([]IEolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEolContext)
		}
	}

	return tst
}

func (s *Instruction_lineContext) Eol(i int) IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Instruction_lineContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *Instruction_lineContext) Comment_line() IComment_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComment_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComment_lineContext)
}

func (s *Instruction_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruction_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruction_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterInstruction_line(s)
	}
}

func (s *Instruction_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitInstruction_line(s)
	}
}

func (p *agcParser) Instruction_line() (localctx IInstruction_lineContext) {
	localctx = NewInstruction_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, agcParserRULE_instruction_line)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserLABEL {
		{
			p.SetState(71)
			p.Label()
		}

	}
	{
		p.SetState(74)
		p.Ws()
	}
	{
		p.SetState(75)
		p.Opcodes()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.Eol()
		}
		{
			p.SetState(77)
			p.Comment_line()
		}

	}
	{
		p.SetState(81)
		p.Argument()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(82)
				p.Eol()
			}
			{
				p.SetState(83)
				p.Argument()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	{
		p.SetState(90)
		p.Eol()
	}

	return localctx
}

// IErase_lineContext is an interface to support dynamic dispatch.
type IErase_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsErase_lineContext differentiates from other interfaces.
	IsErase_lineContext()
}

type Erase_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErase_lineContext() *Erase_lineContext {
	var p = new(Erase_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_erase_line
	return p
}

func (*Erase_lineContext) IsErase_lineContext() {}

func NewErase_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Erase_lineContext {
	var p = new(Erase_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_erase_line

	return p
}

func (s *Erase_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Erase_lineContext) AllWs() []IWsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWsContext)(nil)).Elem())
	var tst = make([]IWsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWsContext)
		}
	}

	return tst
}

func (s *Erase_lineContext) Ws(i int) IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *Erase_lineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Erase_lineContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Erase_lineContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Erase_lineContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Erase_lineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Erase_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Erase_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Erase_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterErase_line(s)
	}
}

func (s *Erase_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitErase_line(s)
	}
}

func (p *agcParser) Erase_line() (localctx IErase_lineContext) {
	localctx = NewErase_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, agcParserRULE_erase_line)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserLABEL || _la == agcParserLPAREN {
		{
			p.SetState(92)
			p.Variable()
		}

	}
	{
		p.SetState(95)
		p.Ws()
	}
	{
		p.SetState(96)
		p.Match(agcParserT__0)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == agcParserWS {
				{
					p.SetState(97)
					p.Ws()
				}

			}
			{
				p.SetState(100)
				p.Expression()
			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(106)
			p.Ws()
		}
		{
			p.SetState(107)
			p.Comment()
		}

	}
	{
		p.SetState(111)
		p.Eol()
	}

	return localctx
}

// IAssignment_lineContext is an interface to support dynamic dispatch.
type IAssignment_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignment_lineContext differentiates from other interfaces.
	IsAssignment_lineContext()
}

type Assignment_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignment_lineContext() *Assignment_lineContext {
	var p = new(Assignment_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_assignment_line
	return p
}

func (*Assignment_lineContext) IsAssignment_lineContext() {}

func NewAssignment_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_lineContext {
	var p = new(Assignment_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_assignment_line

	return p
}

func (s *Assignment_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_lineContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Assignment_lineContext) Eol() IEolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEolContext)
}

func (s *Assignment_lineContext) AllWs() []IWsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWsContext)(nil)).Elem())
	var tst = make([]IWsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWsContext)
		}
	}

	return tst
}

func (s *Assignment_lineContext) Ws(i int) IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *Assignment_lineContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Assignment_lineContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assignment_lineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Assignment_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterAssignment_line(s)
	}
}

func (s *Assignment_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitAssignment_line(s)
	}
}

func (p *agcParser) Assignment_line() (localctx IAssignment_lineContext) {
	localctx = NewAssignment_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, agcParserRULE_assignment_line)
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
		p.SetState(113)
		p.Variable()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserWS {
		{
			p.SetState(114)
			p.Ws()
		}

	}
	p.SetState(117)
	_la = p.GetTokenStream().LA(1)

	if !(_la == agcParserT__1 || _la == agcParserT__2) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == agcParserWS {
				{
					p.SetState(118)
					p.Ws()
				}

			}
			{
				p.SetState(121)
				p.Expression()
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(127)
			p.Ws()
		}
		{
			p.SetState(128)
			p.Comment()
		}

	}
	{
		p.SetState(132)
		p.Eol()
	}

	return localctx
}

// IOpcodesContext is an interface to support dynamic dispatch.
type IOpcodesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodesContext differentiates from other interfaces.
	IsOpcodesContext()
}

type OpcodesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodesContext() *OpcodesContext {
	var p = new(OpcodesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_opcodes
	return p
}

func (*OpcodesContext) IsOpcodesContext() {}

func NewOpcodesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodesContext {
	var p = new(OpcodesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_opcodes

	return p
}

func (s *OpcodesContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodesContext) AllOpcode() []IOpcodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpcodeContext)(nil)).Elem())
	var tst = make([]IOpcodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpcodeContext)
		}
	}

	return tst
}

func (s *OpcodesContext) Opcode(i int) IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *OpcodesContext) Ws() IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *OpcodesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterOpcodes(s)
	}
}

func (s *OpcodesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitOpcodes(s)
	}
}

func (p *agcParser) Opcodes() (localctx IOpcodesContext) {
	localctx = NewOpcodesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, agcParserRULE_opcodes)

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
		p.Opcode()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(135)
			p.Ws()
		}
		{
			p.SetState(136)
			p.Opcode()
		}

	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) AllWs() []IWsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWsContext)(nil)).Elem())
	var tst = make([]IWsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWsContext)
		}
	}

	return tst
}

func (s *ArgumentContext) Ws(i int) IWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWsContext)
}

func (s *ArgumentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *agcParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, agcParserRULE_argument)

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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(140)
				p.Ws()
			}
			{
				p.SetState(141)
				p.Expression()
			}

		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
			p.Ws()
		}
		{
			p.SetState(149)
			p.Comment()
		}

	}

	return localctx
}

// IWsContext is an interface to support dynamic dispatch.
type IWsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWsContext differentiates from other interfaces.
	IsWsContext()
}

type WsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWsContext() *WsContext {
	var p = new(WsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_ws
	return p
}

func (*WsContext) IsWsContext() {}

func NewWsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WsContext {
	var p = new(WsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_ws

	return p
}

func (s *WsContext) GetParser() antlr.Parser { return s.parser }

func (s *WsContext) WS() antlr.TerminalNode {
	return s.GetToken(agcParserWS, 0)
}

func (s *WsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterWs(s)
	}
}

func (s *WsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitWs(s)
	}
}

func (p *agcParser) Ws() (localctx IWsContext) {
	localctx = NewWsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, agcParserRULE_ws)

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
		p.Match(agcParserWS)
	}

	return localctx
}

// IEolContext is an interface to support dynamic dispatch.
type IEolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEolContext differentiates from other interfaces.
	IsEolContext()
}

type EolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEolContext() *EolContext {
	var p = new(EolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_eol
	return p
}

func (*EolContext) IsEolContext() {}

func NewEolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EolContext {
	var p = new(EolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_eol

	return p
}

func (s *EolContext) GetParser() antlr.Parser { return s.parser }

func (s *EolContext) EOL() antlr.TerminalNode {
	return s.GetToken(agcParserEOL, 0)
}

func (s *EolContext) WS() antlr.TerminalNode {
	return s.GetToken(agcParserWS, 0)
}

func (s *EolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterEol(s)
	}
}

func (s *EolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitEol(s)
	}
}

func (p *agcParser) Eol() (localctx IEolContext) {
	localctx = NewEolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, agcParserRULE_eol)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserWS {
		{
			p.SetState(155)
			p.Match(agcParserWS)
		}

	}
	{
		p.SetState(158)
		p.Match(agcParserEOL)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(agcParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *agcParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, agcParserRULE_comment)

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
		p.SetState(160)
		p.Match(agcParserCOMMENT)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LABEL() antlr.TerminalNode {
	return s.GetToken(agcParserLABEL, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *agcParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, agcParserRULE_label)

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
		p.Match(agcParserLABEL)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) LABEL() antlr.TerminalNode {
	return s.GetToken(agcParserLABEL, 0)
}

func (s *VariableContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(agcParserLPAREN, 0)
}

func (s *VariableContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(agcParserRPAREN, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *agcParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, agcParserRULE_variable)

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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case agcParserLABEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(agcParserLABEL)
		}

	case agcParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(agcParserLPAREN)
		}
		{
			p.SetState(166)
			p.Match(agcParserLABEL)
		}
		{
			p.SetState(167)
			p.Match(agcParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(agcParserPLUS)
}

func (s *ExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(agcParserPLUS, i)
}

func (s *ExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(agcParserMINUS)
}

func (s *ExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(agcParserMINUS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *agcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, agcParserRULE_expression)
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
		p.SetState(170)
		p.MultiplyingExpression()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(171)
			_la = p.GetTokenStream().LA(1)

			if !(_la == agcParserPLUS || _la == agcParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(172)
				p.MultiplyingExpression()
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *MultiplyingExpressionContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(agcParserTIMES)
}

func (s *MultiplyingExpressionContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(agcParserTIMES, i)
}

func (s *MultiplyingExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(agcParserDIV)
}

func (s *MultiplyingExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(agcParserDIV, i)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *agcParser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, agcParserRULE_multiplyingExpression)
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
		p.SetState(178)
		p.Atom()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == agcParserTIMES || _la == agcParserDIV {
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(_la == agcParserTIMES || _la == agcParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(180)
			p.Atom()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Inte() IInteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteContext)
}

func (s *AtomContext) Decimal() IDecimalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalContext)
}

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *AtomContext) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *agcParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, agcParserRULE_atom)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Inte()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.Decimal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Variable()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.Register()
		}

	}

	return localctx
}

// IInteContext is an interface to support dynamic dispatch.
type IInteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteContext differentiates from other interfaces.
	IsInteContext()
}

type InteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteContext() *InteContext {
	var p = new(InteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_inte
	return p
}

func (*InteContext) IsInteContext() {}

func NewInteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InteContext {
	var p = new(InteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_inte

	return p
}

func (s *InteContext) GetParser() antlr.Parser { return s.parser }

func (s *InteContext) INTE() antlr.TerminalNode {
	return s.GetToken(agcParserINTE, 0)
}

func (s *InteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterInte(s)
	}
}

func (s *InteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitInte(s)
	}
}

func (p *agcParser) Inte() (localctx IInteContext) {
	localctx = NewInteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, agcParserRULE_inte)

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
		p.Match(agcParserINTE)
	}

	return localctx
}

// IDecimalContext is an interface to support dynamic dispatch.
type IDecimalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalContext differentiates from other interfaces.
	IsDecimalContext()
}

type DecimalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalContext() *DecimalContext {
	var p = new(DecimalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_decimal
	return p
}

func (*DecimalContext) IsDecimalContext() {}

func NewDecimalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalContext {
	var p = new(DecimalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_decimal

	return p
}

func (s *DecimalContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(agcParserDECIMAL, 0)
}

func (s *DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitDecimal(s)
	}
}

func (p *agcParser) Decimal() (localctx IDecimalContext) {
	localctx = NewDecimalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, agcParserRULE_decimal)
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
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == agcParserPLUS || _la == agcParserMINUS {
		p.SetState(195)
		_la = p.GetTokenStream().LA(1)

		if !(_la == agcParserPLUS || _la == agcParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(198)
		p.Match(agcParserDECIMAL)
	}

	return localctx
}

// IRegisterContext is an interface to support dynamic dispatch.
type IRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegisterContext differentiates from other interfaces.
	IsRegisterContext()
}

type RegisterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegisterContext() *RegisterContext {
	var p = new(RegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_register
	return p
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }
func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitRegister(s)
	}
}

func (p *agcParser) Register() (localctx IRegisterContext) {
	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, agcParserRULE_register)
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
	p.SetState(200)
	_la = p.GetTokenStream().LA(1)

	if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<agcParserT__3)|(1<<agcParserT__4)|(1<<agcParserT__5)|(1<<agcParserT__6)|(1<<agcParserT__7)|(1<<agcParserT__8)|(1<<agcParserT__9)|(1<<agcParserT__10)|(1<<agcParserT__11)|(1<<agcParserT__12)|(1<<agcParserT__13)|(1<<agcParserT__14)|(1<<agcParserT__15)|(1<<agcParserT__16)|(1<<agcParserT__17)|(1<<agcParserT__18)|(1<<agcParserT__19)|(1<<agcParserT__20)|(1<<agcParserT__21)|(1<<agcParserT__22)|(1<<agcParserT__23)|(1<<agcParserT__24)|(1<<agcParserT__25)|(1<<agcParserT__26)|(1<<agcParserT__27)|(1<<agcParserT__28)|(1<<agcParserT__29)|(1<<agcParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(agcParserT__31-32))|(1<<(agcParserT__32-32))|(1<<(agcParserT__33-32))|(1<<(agcParserT__34-32))|(1<<(agcParserT__35-32))|(1<<(agcParserT__36-32))|(1<<(agcParserT__37-32))|(1<<(agcParserT__38-32))|(1<<(agcParserT__39-32))|(1<<(agcParserT__40-32))|(1<<(agcParserT__41-32))|(1<<(agcParserT__42-32))|(1<<(agcParserT__43-32))|(1<<(agcParserT__44-32))|(1<<(agcParserT__45-32))|(1<<(agcParserT__46-32))|(1<<(agcParserT__47-32))|(1<<(agcParserT__48-32))|(1<<(agcParserT__49-32))|(1<<(agcParserT__50-32))|(1<<(agcParserT__51-32)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) Standard_opcode() IStandard_opcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStandard_opcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStandard_opcodeContext)
}

func (s *OpcodeContext) Pseudo_opcode() IPseudo_opcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudo_opcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudo_opcodeContext)
}

func (s *OpcodeContext) Axt_opcode() IAxt_opcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxt_opcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAxt_opcodeContext)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *agcParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, agcParserRULE_opcode)

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

	p.SetState(205)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case agcParserT__82, agcParserT__83, agcParserT__84, agcParserT__85, agcParserT__86, agcParserT__87, agcParserT__88, agcParserT__89, agcParserT__90, agcParserT__91, agcParserT__92, agcParserT__93, agcParserT__94, agcParserT__95, agcParserT__96, agcParserT__97, agcParserT__98, agcParserT__99, agcParserT__100, agcParserT__101, agcParserT__102, agcParserT__103, agcParserT__104, agcParserT__105, agcParserT__106, agcParserT__107, agcParserT__108, agcParserT__109, agcParserT__110, agcParserT__111, agcParserT__112, agcParserT__113, agcParserT__114, agcParserT__115, agcParserT__116, agcParserT__117, agcParserT__118, agcParserT__119, agcParserT__120, agcParserT__121, agcParserT__122, agcParserT__123, agcParserT__124, agcParserT__125, agcParserT__126, agcParserT__127, agcParserT__128, agcParserT__129, agcParserT__130, agcParserT__131, agcParserT__132, agcParserT__133, agcParserT__134, agcParserT__135, agcParserT__136, agcParserT__137, agcParserT__138, agcParserT__139, agcParserT__140, agcParserT__141, agcParserT__142, agcParserT__143, agcParserT__144, agcParserT__145, agcParserT__146, agcParserT__147, agcParserT__148, agcParserT__149, agcParserT__150, agcParserT__151, agcParserT__152, agcParserT__153, agcParserT__154, agcParserT__155, agcParserT__156, agcParserT__157, agcParserT__158, agcParserT__159, agcParserT__160, agcParserT__161, agcParserT__162, agcParserT__163, agcParserT__164, agcParserT__165, agcParserT__166, agcParserT__167, agcParserT__168, agcParserT__169, agcParserT__170, agcParserT__171, agcParserT__172, agcParserT__173, agcParserT__174, agcParserT__175, agcParserT__176, agcParserT__177, agcParserT__178, agcParserT__179, agcParserT__180, agcParserT__181, agcParserT__182:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Standard_opcode()
		}

	case agcParserT__54, agcParserT__55, agcParserT__56, agcParserT__57, agcParserT__58, agcParserT__59, agcParserT__60, agcParserT__61, agcParserT__62, agcParserT__63, agcParserT__64, agcParserT__65, agcParserT__66, agcParserT__67, agcParserT__68, agcParserT__69, agcParserT__70, agcParserT__71, agcParserT__72, agcParserT__73, agcParserT__74, agcParserT__75, agcParserT__76, agcParserT__77, agcParserT__78, agcParserT__79, agcParserT__80, agcParserT__81:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Pseudo_opcode()
		}

	case agcParserT__52, agcParserT__53:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(204)
			p.Axt_opcode()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAxt_opcodeContext is an interface to support dynamic dispatch.
type IAxt_opcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxt_opcodeContext differentiates from other interfaces.
	IsAxt_opcodeContext()
}

type Axt_opcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxt_opcodeContext() *Axt_opcodeContext {
	var p = new(Axt_opcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_axt_opcode
	return p
}

func (*Axt_opcodeContext) IsAxt_opcodeContext() {}

func NewAxt_opcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Axt_opcodeContext {
	var p = new(Axt_opcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_axt_opcode

	return p
}

func (s *Axt_opcodeContext) GetParser() antlr.Parser { return s.parser }
func (s *Axt_opcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Axt_opcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Axt_opcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterAxt_opcode(s)
	}
}

func (s *Axt_opcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitAxt_opcode(s)
	}
}

func (p *agcParser) Axt_opcode() (localctx IAxt_opcodeContext) {
	localctx = NewAxt_opcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, agcParserRULE_axt_opcode)
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
	p.SetState(207)
	_la = p.GetTokenStream().LA(1)

	if !(_la == agcParserT__52 || _la == agcParserT__53) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPseudo_opcodeContext is an interface to support dynamic dispatch.
type IPseudo_opcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudo_opcodeContext differentiates from other interfaces.
	IsPseudo_opcodeContext()
}

type Pseudo_opcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudo_opcodeContext() *Pseudo_opcodeContext {
	var p = new(Pseudo_opcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_pseudo_opcode
	return p
}

func (*Pseudo_opcodeContext) IsPseudo_opcodeContext() {}

func NewPseudo_opcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pseudo_opcodeContext {
	var p = new(Pseudo_opcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_pseudo_opcode

	return p
}

func (s *Pseudo_opcodeContext) GetParser() antlr.Parser { return s.parser }
func (s *Pseudo_opcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pseudo_opcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pseudo_opcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterPseudo_opcode(s)
	}
}

func (s *Pseudo_opcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitPseudo_opcode(s)
	}
}

func (p *agcParser) Pseudo_opcode() (localctx IPseudo_opcodeContext) {
	localctx = NewPseudo_opcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, agcParserRULE_pseudo_opcode)
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
	p.SetState(209)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(agcParserT__54-55))|(1<<(agcParserT__55-55))|(1<<(agcParserT__56-55))|(1<<(agcParserT__57-55))|(1<<(agcParserT__58-55))|(1<<(agcParserT__59-55))|(1<<(agcParserT__60-55))|(1<<(agcParserT__61-55))|(1<<(agcParserT__62-55))|(1<<(agcParserT__63-55))|(1<<(agcParserT__64-55))|(1<<(agcParserT__65-55))|(1<<(agcParserT__66-55))|(1<<(agcParserT__67-55))|(1<<(agcParserT__68-55))|(1<<(agcParserT__69-55))|(1<<(agcParserT__70-55))|(1<<(agcParserT__71-55))|(1<<(agcParserT__72-55))|(1<<(agcParserT__73-55))|(1<<(agcParserT__74-55))|(1<<(agcParserT__75-55))|(1<<(agcParserT__76-55))|(1<<(agcParserT__77-55))|(1<<(agcParserT__78-55))|(1<<(agcParserT__79-55))|(1<<(agcParserT__80-55))|(1<<(agcParserT__81-55)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IStandard_opcodeContext is an interface to support dynamic dispatch.
type IStandard_opcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStandard_opcodeContext differentiates from other interfaces.
	IsStandard_opcodeContext()
}

type Standard_opcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStandard_opcodeContext() *Standard_opcodeContext {
	var p = new(Standard_opcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = agcParserRULE_standard_opcode
	return p
}

func (*Standard_opcodeContext) IsStandard_opcodeContext() {}

func NewStandard_opcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Standard_opcodeContext {
	var p = new(Standard_opcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = agcParserRULE_standard_opcode

	return p
}

func (s *Standard_opcodeContext) GetParser() antlr.Parser { return s.parser }
func (s *Standard_opcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Standard_opcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Standard_opcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.EnterStandard_opcode(s)
	}
}

func (s *Standard_opcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agcListener); ok {
		listenerT.ExitStandard_opcode(s)
	}
}

func (p *agcParser) Standard_opcode() (localctx IStandard_opcodeContext) {
	localctx = NewStandard_opcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, agcParserRULE_standard_opcode)
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
	p.SetState(211)
	_la = p.GetTokenStream().LA(1)

	if !((((_la-83)&-(0x1f+1)) == 0 && ((1<<uint((_la-83)))&((1<<(agcParserT__82-83))|(1<<(agcParserT__83-83))|(1<<(agcParserT__84-83))|(1<<(agcParserT__85-83))|(1<<(agcParserT__86-83))|(1<<(agcParserT__87-83))|(1<<(agcParserT__88-83))|(1<<(agcParserT__89-83))|(1<<(agcParserT__90-83))|(1<<(agcParserT__91-83))|(1<<(agcParserT__92-83))|(1<<(agcParserT__93-83))|(1<<(agcParserT__94-83))|(1<<(agcParserT__95-83))|(1<<(agcParserT__96-83))|(1<<(agcParserT__97-83))|(1<<(agcParserT__98-83))|(1<<(agcParserT__99-83))|(1<<(agcParserT__100-83))|(1<<(agcParserT__101-83))|(1<<(agcParserT__102-83))|(1<<(agcParserT__103-83))|(1<<(agcParserT__104-83))|(1<<(agcParserT__105-83))|(1<<(agcParserT__106-83))|(1<<(agcParserT__107-83))|(1<<(agcParserT__108-83))|(1<<(agcParserT__109-83))|(1<<(agcParserT__110-83))|(1<<(agcParserT__111-83))|(1<<(agcParserT__112-83))|(1<<(agcParserT__113-83)))) != 0) || (((_la-115)&-(0x1f+1)) == 0 && ((1<<uint((_la-115)))&((1<<(agcParserT__114-115))|(1<<(agcParserT__115-115))|(1<<(agcParserT__116-115))|(1<<(agcParserT__117-115))|(1<<(agcParserT__118-115))|(1<<(agcParserT__119-115))|(1<<(agcParserT__120-115))|(1<<(agcParserT__121-115))|(1<<(agcParserT__122-115))|(1<<(agcParserT__123-115))|(1<<(agcParserT__124-115))|(1<<(agcParserT__125-115))|(1<<(agcParserT__126-115))|(1<<(agcParserT__127-115))|(1<<(agcParserT__128-115))|(1<<(agcParserT__129-115))|(1<<(agcParserT__130-115))|(1<<(agcParserT__131-115))|(1<<(agcParserT__132-115))|(1<<(agcParserT__133-115))|(1<<(agcParserT__134-115))|(1<<(agcParserT__135-115))|(1<<(agcParserT__136-115))|(1<<(agcParserT__137-115))|(1<<(agcParserT__138-115))|(1<<(agcParserT__139-115))|(1<<(agcParserT__140-115))|(1<<(agcParserT__141-115))|(1<<(agcParserT__142-115))|(1<<(agcParserT__143-115))|(1<<(agcParserT__144-115))|(1<<(agcParserT__145-115)))) != 0) || (((_la-147)&-(0x1f+1)) == 0 && ((1<<uint((_la-147)))&((1<<(agcParserT__146-147))|(1<<(agcParserT__147-147))|(1<<(agcParserT__148-147))|(1<<(agcParserT__149-147))|(1<<(agcParserT__150-147))|(1<<(agcParserT__151-147))|(1<<(agcParserT__152-147))|(1<<(agcParserT__153-147))|(1<<(agcParserT__154-147))|(1<<(agcParserT__155-147))|(1<<(agcParserT__156-147))|(1<<(agcParserT__157-147))|(1<<(agcParserT__158-147))|(1<<(agcParserT__159-147))|(1<<(agcParserT__160-147))|(1<<(agcParserT__161-147))|(1<<(agcParserT__162-147))|(1<<(agcParserT__163-147))|(1<<(agcParserT__164-147))|(1<<(agcParserT__165-147))|(1<<(agcParserT__166-147))|(1<<(agcParserT__167-147))|(1<<(agcParserT__168-147))|(1<<(agcParserT__169-147))|(1<<(agcParserT__170-147))|(1<<(agcParserT__171-147))|(1<<(agcParserT__172-147))|(1<<(agcParserT__173-147))|(1<<(agcParserT__174-147))|(1<<(agcParserT__175-147))|(1<<(agcParserT__176-147))|(1<<(agcParserT__177-147)))) != 0) || (((_la-179)&-(0x1f+1)) == 0 && ((1<<uint((_la-179)))&((1<<(agcParserT__178-179))|(1<<(agcParserT__179-179))|(1<<(agcParserT__180-179))|(1<<(agcParserT__181-179))|(1<<(agcParserT__182-179)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
