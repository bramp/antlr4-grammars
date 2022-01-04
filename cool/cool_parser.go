// Code generated from COOL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cool // COOL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 217,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 22, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 28, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 34, 10, 4, 12,
	4, 14, 4, 37, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 46,
	10, 5, 12, 5, 14, 5, 49, 11, 5, 5, 5, 51, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 65, 10, 5, 5, 5,
	67, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	7, 7, 79, 10, 7, 12, 7, 14, 7, 82, 11, 7, 5, 7, 84, 10, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 105, 10, 7, 13, 7, 14, 7, 106, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 117, 10, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 125, 10, 7, 7, 7, 127, 10, 7, 12, 7, 14,
	7, 130, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 6, 7, 144, 10, 7, 13, 7, 14, 7, 145, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 170, 10, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 196, 10,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 204, 10, 7, 12, 7, 14, 7,
	207, 11, 7, 5, 7, 209, 10, 7, 3, 7, 7, 7, 212, 10, 7, 12, 7, 14, 7, 215,
	11, 7, 3, 7, 2, 3, 12, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 251, 2, 14, 3, 2,
	2, 2, 4, 21, 3, 2, 2, 2, 6, 23, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 68,
	3, 2, 2, 2, 12, 169, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 3, 3, 2, 2, 2,
	16, 17, 5, 6, 4, 2, 17, 18, 7, 3, 2, 2, 18, 19, 5, 4, 3, 2, 19, 22, 3,
	2, 2, 2, 20, 22, 7, 2, 2, 3, 21, 16, 3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22,
	5, 3, 2, 2, 2, 23, 24, 7, 12, 2, 2, 24, 27, 7, 33, 2, 2, 25, 26, 7, 18,
	2, 2, 26, 28, 7, 33, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28,
	29, 3, 2, 2, 2, 29, 35, 7, 4, 2, 2, 30, 31, 5, 8, 5, 2, 31, 32, 7, 3, 2,
	2, 32, 34, 3, 2, 2, 2, 33, 30, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33,
	3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2,
	38, 39, 7, 5, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 34, 2, 2, 41, 50, 7,
	6, 2, 2, 42, 47, 5, 10, 6, 2, 43, 44, 7, 7, 2, 2, 44, 46, 5, 10, 6, 2,
	45, 43, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3,
	2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 42, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 7, 8, 2, 2, 53, 54, 7, 9, 2,
	2, 54, 55, 7, 33, 2, 2, 55, 56, 7, 4, 2, 2, 56, 57, 5, 12, 7, 2, 57, 58,
	7, 5, 2, 2, 58, 67, 3, 2, 2, 2, 59, 60, 7, 34, 2, 2, 60, 61, 7, 9, 2, 2,
	61, 64, 7, 33, 2, 2, 62, 63, 7, 35, 2, 2, 63, 65, 5, 12, 7, 2, 64, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 40, 3, 2, 2, 2,
	66, 59, 3, 2, 2, 2, 67, 9, 3, 2, 2, 2, 68, 69, 7, 34, 2, 2, 69, 70, 7,
	9, 2, 2, 70, 71, 7, 33, 2, 2, 71, 11, 3, 2, 2, 2, 72, 73, 8, 7, 1, 2, 73,
	74, 7, 34, 2, 2, 74, 83, 7, 6, 2, 2, 75, 80, 5, 12, 7, 2, 76, 77, 7, 7,
	2, 2, 77, 79, 5, 12, 7, 2, 78, 76, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80,
	78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 83, 75, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 170,
	7, 8, 2, 2, 86, 87, 7, 16, 2, 2, 87, 88, 5, 12, 7, 2, 88, 89, 7, 23, 2,
	2, 89, 90, 5, 12, 7, 2, 90, 91, 7, 13, 2, 2, 91, 92, 5, 12, 7, 2, 92, 93,
	7, 15, 2, 2, 93, 170, 3, 2, 2, 2, 94, 95, 7, 24, 2, 2, 95, 96, 5, 12, 7,
	2, 96, 97, 7, 21, 2, 2, 97, 98, 5, 12, 7, 2, 98, 99, 7, 22, 2, 2, 99, 170,
	3, 2, 2, 2, 100, 104, 7, 4, 2, 2, 101, 102, 5, 12, 7, 2, 102, 103, 7, 3,
	2, 2, 103, 105, 3, 2, 2, 2, 104, 101, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2,
	106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	109, 7, 5, 2, 2, 109, 170, 3, 2, 2, 2, 110, 111, 7, 20, 2, 2, 111, 112,
	7, 34, 2, 2, 112, 113, 7, 9, 2, 2, 113, 116, 7, 33, 2, 2, 114, 115, 7,
	35, 2, 2, 115, 117, 5, 12, 7, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2,
	2, 2, 117, 128, 3, 2, 2, 2, 118, 119, 7, 7, 2, 2, 119, 120, 7, 34, 2, 2,
	120, 121, 7, 9, 2, 2, 121, 124, 7, 33, 2, 2, 122, 123, 7, 35, 2, 2, 123,
	125, 5, 12, 7, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127,
	3, 2, 2, 2, 126, 118, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2,
	2, 2, 128, 129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2,
	131, 132, 7, 17, 2, 2, 132, 170, 5, 12, 7, 22, 133, 134, 7, 25, 2, 2, 134,
	135, 5, 12, 7, 2, 135, 143, 7, 28, 2, 2, 136, 137, 7, 34, 2, 2, 137, 138,
	7, 9, 2, 2, 138, 139, 7, 33, 2, 2, 139, 140, 7, 36, 2, 2, 140, 141, 5,
	12, 7, 2, 141, 142, 7, 3, 2, 2, 142, 144, 3, 2, 2, 2, 143, 136, 3, 2, 2,
	2, 144, 145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 148, 7, 26, 2, 2, 148, 170, 3, 2, 2, 2, 149, 150,
	7, 27, 2, 2, 150, 170, 7, 33, 2, 2, 151, 152, 7, 44, 2, 2, 152, 170, 5,
	12, 7, 19, 153, 154, 7, 19, 2, 2, 154, 170, 5, 12, 7, 18, 155, 156, 7,
	29, 2, 2, 156, 170, 5, 12, 7, 10, 157, 158, 7, 6, 2, 2, 158, 159, 5, 12,
	7, 2, 159, 160, 7, 8, 2, 2, 160, 170, 3, 2, 2, 2, 161, 170, 7, 34, 2, 2,
	162, 170, 7, 32, 2, 2, 163, 170, 7, 31, 2, 2, 164, 170, 7, 30, 2, 2, 165,
	170, 7, 14, 2, 2, 166, 167, 7, 34, 2, 2, 167, 168, 7, 35, 2, 2, 168, 170,
	5, 12, 7, 3, 169, 72, 3, 2, 2, 2, 169, 86, 3, 2, 2, 2, 169, 94, 3, 2, 2,
	2, 169, 100, 3, 2, 2, 2, 169, 110, 3, 2, 2, 2, 169, 133, 3, 2, 2, 2, 169,
	149, 3, 2, 2, 2, 169, 151, 3, 2, 2, 2, 169, 153, 3, 2, 2, 2, 169, 155,
	3, 2, 2, 2, 169, 157, 3, 2, 2, 2, 169, 161, 3, 2, 2, 2, 169, 162, 3, 2,
	2, 2, 169, 163, 3, 2, 2, 2, 169, 164, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2,
	169, 166, 3, 2, 2, 2, 170, 213, 3, 2, 2, 2, 171, 172, 12, 17, 2, 2, 172,
	173, 7, 39, 2, 2, 173, 212, 5, 12, 7, 18, 174, 175, 12, 16, 2, 2, 175,
	176, 7, 40, 2, 2, 176, 212, 5, 12, 7, 17, 177, 178, 12, 15, 2, 2, 178,
	179, 7, 37, 2, 2, 179, 212, 5, 12, 7, 16, 180, 181, 12, 14, 2, 2, 181,
	182, 7, 38, 2, 2, 182, 212, 5, 12, 7, 15, 183, 184, 12, 13, 2, 2, 184,
	185, 7, 41, 2, 2, 185, 212, 5, 12, 7, 14, 186, 187, 12, 12, 2, 2, 187,
	188, 7, 42, 2, 2, 188, 212, 5, 12, 7, 13, 189, 190, 12, 11, 2, 2, 190,
	191, 7, 43, 2, 2, 191, 212, 5, 12, 7, 12, 192, 195, 12, 27, 2, 2, 193,
	194, 7, 10, 2, 2, 194, 196, 7, 33, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196,
	3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 7, 11, 2, 2, 198, 199, 7, 34,
	2, 2, 199, 208, 7, 6, 2, 2, 200, 205, 5, 12, 7, 2, 201, 202, 7, 7, 2, 2,
	202, 204, 5, 12, 7, 2, 203, 201, 3, 2, 2, 2, 204, 207, 3, 2, 2, 2, 205,
	203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205,
	3, 2, 2, 2, 208, 200, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 3, 2,
	2, 2, 210, 212, 7, 8, 2, 2, 211, 171, 3, 2, 2, 2, 211, 174, 3, 2, 2, 2,
	211, 177, 3, 2, 2, 2, 211, 180, 3, 2, 2, 2, 211, 183, 3, 2, 2, 2, 211,
	186, 3, 2, 2, 2, 211, 189, 3, 2, 2, 2, 211, 192, 3, 2, 2, 2, 212, 215,
	3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 13, 3, 2,
	2, 2, 215, 213, 3, 2, 2, 2, 22, 21, 27, 35, 47, 50, 64, 66, 80, 83, 106,
	116, 124, 128, 145, 169, 195, 205, 208, 211, 213,
}
var literalNames = []string{
	"", "';'", "'{'", "'}'", "'('", "','", "')'", "':'", "'@'", "'.'", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "'<-'", "'=>'", "'+'", "'-'", "'*'", "'/'", "'<'", "'<='",
	"'='", "'~'", "'(*'", "'*)'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "CLASS", "ELSE", "FALSE", "FI",
	"IF", "IN", "INHERITS", "ISVOID", "LET", "LOOP", "POOL", "THEN", "WHILE",
	"CASE", "ESAC", "NEW", "OF", "NOT", "TRUE", "STRING", "INT", "TYPEID",
	"OBJECTID", "ASSIGNMENT", "CASE_ARROW", "ADD", "MINUS", "MULTIPLY", "DIVISION",
	"LESS_THAN", "LESS_EQUAL", "EQUAL", "INTEGER_NEGATIVE", "OPEN_COMMENT",
	"CLOSE_COMMENT", "COMMENT", "ONE_LINE_COMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"program", "programBlocks", "classDefine", "feature", "formal", "expression",
}

type COOLParser struct {
	*antlr.BaseParser
}

// NewCOOLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *COOLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCOOLParser(input antlr.TokenStream) *COOLParser {
	this := new(COOLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "COOL.g4"

	return this
}

// COOLParser tokens.
const (
	COOLParserEOF              = antlr.TokenEOF
	COOLParserT__0             = 1
	COOLParserT__1             = 2
	COOLParserT__2             = 3
	COOLParserT__3             = 4
	COOLParserT__4             = 5
	COOLParserT__5             = 6
	COOLParserT__6             = 7
	COOLParserT__7             = 8
	COOLParserT__8             = 9
	COOLParserCLASS            = 10
	COOLParserELSE             = 11
	COOLParserFALSE            = 12
	COOLParserFI               = 13
	COOLParserIF               = 14
	COOLParserIN               = 15
	COOLParserINHERITS         = 16
	COOLParserISVOID           = 17
	COOLParserLET              = 18
	COOLParserLOOP             = 19
	COOLParserPOOL             = 20
	COOLParserTHEN             = 21
	COOLParserWHILE            = 22
	COOLParserCASE             = 23
	COOLParserESAC             = 24
	COOLParserNEW              = 25
	COOLParserOF               = 26
	COOLParserNOT              = 27
	COOLParserTRUE             = 28
	COOLParserSTRING           = 29
	COOLParserINT              = 30
	COOLParserTYPEID           = 31
	COOLParserOBJECTID         = 32
	COOLParserASSIGNMENT       = 33
	COOLParserCASE_ARROW       = 34
	COOLParserADD              = 35
	COOLParserMINUS            = 36
	COOLParserMULTIPLY         = 37
	COOLParserDIVISION         = 38
	COOLParserLESS_THAN        = 39
	COOLParserLESS_EQUAL       = 40
	COOLParserEQUAL            = 41
	COOLParserINTEGER_NEGATIVE = 42
	COOLParserOPEN_COMMENT     = 43
	COOLParserCLOSE_COMMENT    = 44
	COOLParserCOMMENT          = 45
	COOLParserONE_LINE_COMMENT = 46
	COOLParserWHITESPACE       = 47
)

// COOLParser rules.
const (
	COOLParserRULE_program       = 0
	COOLParserRULE_programBlocks = 1
	COOLParserRULE_classDefine   = 2
	COOLParserRULE_feature       = 3
	COOLParserRULE_formal        = 4
	COOLParserRULE_expression    = 5
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = COOLParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ProgramBlocks() IProgramBlocksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramBlocksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramBlocksContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *COOLParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, COOLParserRULE_program)

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
		p.SetState(12)
		p.ProgramBlocks()
	}

	return localctx
}

// IProgramBlocksContext is an interface to support dynamic dispatch.
type IProgramBlocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramBlocksContext differentiates from other interfaces.
	IsProgramBlocksContext()
}

type ProgramBlocksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBlocksContext() *ProgramBlocksContext {
	var p = new(ProgramBlocksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = COOLParserRULE_programBlocks
	return p
}

func (*ProgramBlocksContext) IsProgramBlocksContext() {}

func NewProgramBlocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBlocksContext {
	var p = new(ProgramBlocksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_programBlocks

	return p
}

func (s *ProgramBlocksContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBlocksContext) CopyFrom(ctx *ProgramBlocksContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ProgramBlocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBlocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ClassesContext struct {
	*ProgramBlocksContext
}

func NewClassesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassesContext {
	var p = new(ClassesContext)

	p.ProgramBlocksContext = NewEmptyProgramBlocksContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ProgramBlocksContext))

	return p
}

func (s *ClassesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassesContext) ClassDefine() IClassDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDefineContext)
}

func (s *ClassesContext) ProgramBlocks() IProgramBlocksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramBlocksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramBlocksContext)
}

func (s *ClassesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterClasses(s)
	}
}

func (s *ClassesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitClasses(s)
	}
}

type EofContext struct {
	*ProgramBlocksContext
}

func NewEofContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EofContext {
	var p = new(EofContext)

	p.ProgramBlocksContext = NewEmptyProgramBlocksContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ProgramBlocksContext))

	return p
}

func (s *EofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EofContext) EOF() antlr.TerminalNode {
	return s.GetToken(COOLParserEOF, 0)
}

func (s *EofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterEof(s)
	}
}

func (s *EofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitEof(s)
	}
}

func (p *COOLParser) ProgramBlocks() (localctx IProgramBlocksContext) {
	this := p
	_ = this

	localctx = NewProgramBlocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, COOLParserRULE_programBlocks)

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

	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case COOLParserCLASS:
		localctx = NewClassesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.ClassDefine()
		}
		{
			p.SetState(15)
			p.Match(COOLParserT__0)
		}
		{
			p.SetState(16)
			p.ProgramBlocks()
		}

	case COOLParserEOF:
		localctx = NewEofContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.Match(COOLParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassDefineContext is an interface to support dynamic dispatch.
type IClassDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDefineContext differentiates from other interfaces.
	IsClassDefineContext()
}

type ClassDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefineContext() *ClassDefineContext {
	var p = new(ClassDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = COOLParserRULE_classDefine
	return p
}

func (*ClassDefineContext) IsClassDefineContext() {}

func NewClassDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefineContext {
	var p = new(ClassDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_classDefine

	return p
}

func (s *ClassDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefineContext) CLASS() antlr.TerminalNode {
	return s.GetToken(COOLParserCLASS, 0)
}

func (s *ClassDefineContext) AllTYPEID() []antlr.TerminalNode {
	return s.GetTokens(COOLParserTYPEID)
}

func (s *ClassDefineContext) TYPEID(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, i)
}

func (s *ClassDefineContext) INHERITS() antlr.TerminalNode {
	return s.GetToken(COOLParserINHERITS, 0)
}

func (s *ClassDefineContext) AllFeature() []IFeatureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFeatureContext)(nil)).Elem())
	var tst = make([]IFeatureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFeatureContext)
		}
	}

	return tst
}

func (s *ClassDefineContext) Feature(i int) IFeatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFeatureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFeatureContext)
}

func (s *ClassDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterClassDefine(s)
	}
}

func (s *ClassDefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitClassDefine(s)
	}
}

func (p *COOLParser) ClassDefine() (localctx IClassDefineContext) {
	this := p
	_ = this

	localctx = NewClassDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, COOLParserRULE_classDefine)
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
		p.SetState(21)
		p.Match(COOLParserCLASS)
	}
	{
		p.SetState(22)
		p.Match(COOLParserTYPEID)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == COOLParserINHERITS {
		{
			p.SetState(23)
			p.Match(COOLParserINHERITS)
		}
		{
			p.SetState(24)
			p.Match(COOLParserTYPEID)
		}

	}
	{
		p.SetState(27)
		p.Match(COOLParserT__1)
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == COOLParserOBJECTID {
		{
			p.SetState(28)
			p.Feature()
		}
		{
			p.SetState(29)
			p.Match(COOLParserT__0)
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(COOLParserT__2)
	}

	return localctx
}

// IFeatureContext is an interface to support dynamic dispatch.
type IFeatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeatureContext differentiates from other interfaces.
	IsFeatureContext()
}

type FeatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureContext() *FeatureContext {
	var p = new(FeatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = COOLParserRULE_feature
	return p
}

func (*FeatureContext) IsFeatureContext() {}

func NewFeatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureContext {
	var p = new(FeatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_feature

	return p
}

func (s *FeatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureContext) CopyFrom(ctx *FeatureContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethodContext struct {
	*FeatureContext
}

func NewMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodContext {
	var p = new(MethodContext)

	p.FeatureContext = NewEmptyFeatureContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FeatureContext))

	return p
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *MethodContext) TYPEID() antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, 0)
}

func (s *MethodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MethodContext) AllFormal() []IFormalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalContext)(nil)).Elem())
	var tst = make([]IFormalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalContext)
		}
	}

	return tst
}

func (s *MethodContext) Formal(i int) IFormalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalContext)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitMethod(s)
	}
}

type PropertyContext struct {
	*FeatureContext
}

func NewPropertyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyContext {
	var p = new(PropertyContext)

	p.FeatureContext = NewEmptyFeatureContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FeatureContext))

	return p
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *PropertyContext) TYPEID() antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, 0)
}

func (s *PropertyContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(COOLParserASSIGNMENT, 0)
}

func (s *PropertyContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *COOLParser) Feature() (localctx IFeatureContext) {
	this := p
	_ = this

	localctx = NewFeatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, COOLParserRULE_feature)
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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(COOLParserOBJECTID)
		}
		{
			p.SetState(39)
			p.Match(COOLParserT__3)
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == COOLParserOBJECTID {
			{
				p.SetState(40)
				p.Formal()
			}
			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == COOLParserT__4 {
				{
					p.SetState(41)
					p.Match(COOLParserT__4)
				}
				{
					p.SetState(42)
					p.Formal()
				}

				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(50)
			p.Match(COOLParserT__5)
		}
		{
			p.SetState(51)
			p.Match(COOLParserT__6)
		}
		{
			p.SetState(52)
			p.Match(COOLParserTYPEID)
		}
		{
			p.SetState(53)
			p.Match(COOLParserT__1)
		}
		{
			p.SetState(54)
			p.expression(0)
		}
		{
			p.SetState(55)
			p.Match(COOLParserT__2)
		}

	case 2:
		localctx = NewPropertyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(COOLParserOBJECTID)
		}
		{
			p.SetState(58)
			p.Match(COOLParserT__6)
		}
		{
			p.SetState(59)
			p.Match(COOLParserTYPEID)
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == COOLParserASSIGNMENT {
			{
				p.SetState(60)
				p.Match(COOLParserASSIGNMENT)
			}
			{
				p.SetState(61)
				p.expression(0)
			}

		}

	}

	return localctx
}

// IFormalContext is an interface to support dynamic dispatch.
type IFormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalContext differentiates from other interfaces.
	IsFormalContext()
}

type FormalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalContext() *FormalContext {
	var p = new(FormalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = COOLParserRULE_formal
	return p
}

func (*FormalContext) IsFormalContext() {}

func NewFormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalContext {
	var p = new(FormalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_formal

	return p
}

func (s *FormalContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *FormalContext) TYPEID() antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, 0)
}

func (s *FormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterFormal(s)
	}
}

func (s *FormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitFormal(s)
	}
}

func (p *COOLParser) Formal() (localctx IFormalContext) {
	this := p
	_ = this

	localctx = NewFormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, COOLParserRULE_formal)

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
		p.Match(COOLParserOBJECTID)
	}
	{
		p.SetState(67)
		p.Match(COOLParserT__6)
	}
	{
		p.SetState(68)
		p.Match(COOLParserTYPEID)
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
	p.RuleIndex = COOLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = COOLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LetInContext struct {
	*ExpressionContext
}

func NewLetInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetInContext {
	var p = new(LetInContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LetInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetInContext) LET() antlr.TerminalNode {
	return s.GetToken(COOLParserLET, 0)
}

func (s *LetInContext) AllOBJECTID() []antlr.TerminalNode {
	return s.GetTokens(COOLParserOBJECTID)
}

func (s *LetInContext) OBJECTID(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, i)
}

func (s *LetInContext) AllTYPEID() []antlr.TerminalNode {
	return s.GetTokens(COOLParserTYPEID)
}

func (s *LetInContext) TYPEID(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, i)
}

func (s *LetInContext) IN() antlr.TerminalNode {
	return s.GetToken(COOLParserIN, 0)
}

func (s *LetInContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LetInContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetInContext) AllASSIGNMENT() []antlr.TerminalNode {
	return s.GetTokens(COOLParserASSIGNMENT)
}

func (s *LetInContext) ASSIGNMENT(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserASSIGNMENT, i)
}

func (s *LetInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterLetIn(s)
	}
}

func (s *LetInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitLetIn(s)
	}
}

type MinusContext struct {
	*ExpressionContext
}

func NewMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusContext {
	var p = new(MinusContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MinusContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(COOLParserMINUS, 0)
}

func (s *MinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterMinus(s)
	}
}

func (s *MinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitMinus(s)
	}
}

type StringContext struct {
	*ExpressionContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(COOLParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitString(s)
	}
}

type IsvoidContext struct {
	*ExpressionContext
}

func NewIsvoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsvoidContext {
	var p = new(IsvoidContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IsvoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsvoidContext) ISVOID() antlr.TerminalNode {
	return s.GetToken(COOLParserISVOID, 0)
}

func (s *IsvoidContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsvoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterIsvoid(s)
	}
}

func (s *IsvoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitIsvoid(s)
	}
}

type WhileContext struct {
	*ExpressionContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(COOLParserWHILE, 0)
}

func (s *WhileContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *WhileContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileContext) LOOP() antlr.TerminalNode {
	return s.GetToken(COOLParserLOOP, 0)
}

func (s *WhileContext) POOL() antlr.TerminalNode {
	return s.GetToken(COOLParserPOOL, 0)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitWhile(s)
	}
}

type DivisionContext struct {
	*ExpressionContext
}

func NewDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivisionContext {
	var p = new(DivisionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DivisionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DivisionContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(COOLParserDIVISION, 0)
}

func (s *DivisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterDivision(s)
	}
}

func (s *DivisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitDivision(s)
	}
}

type NegativeContext struct {
	*ExpressionContext
}

func NewNegativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeContext {
	var p = new(NegativeContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeContext) INTEGER_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(COOLParserINTEGER_NEGATIVE, 0)
}

func (s *NegativeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterNegative(s)
	}
}

func (s *NegativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitNegative(s)
	}
}

type BoolNotContext struct {
	*ExpressionContext
}

func NewBoolNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolNotContext {
	var p = new(BoolNotContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BoolNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolNotContext) NOT() antlr.TerminalNode {
	return s.GetToken(COOLParserNOT, 0)
}

func (s *BoolNotContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BoolNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterBoolNot(s)
	}
}

func (s *BoolNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitBoolNot(s)
	}
}

type LessThanContext struct {
	*ExpressionContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessThanContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LessThanContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(COOLParserLESS_THAN, 0)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitLessThan(s)
	}
}

type BlockContext struct {
	*ExpressionContext
}

func NewBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockContext {
	var p = new(BlockContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BlockContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitBlock(s)
	}
}

type IdContext struct {
	*ExpressionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitId(s)
	}
}

type MultiplyContext struct {
	*ExpressionContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplyContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(COOLParserMULTIPLY, 0)
}

func (s *MultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterMultiply(s)
	}
}

func (s *MultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitMultiply(s)
	}
}

type IfContext struct {
	*ExpressionContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(COOLParserIF, 0)
}

func (s *IfContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IfContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfContext) THEN() antlr.TerminalNode {
	return s.GetToken(COOLParserTHEN, 0)
}

func (s *IfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(COOLParserELSE, 0)
}

func (s *IfContext) FI() antlr.TerminalNode {
	return s.GetToken(COOLParserFI, 0)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitIf(s)
	}
}

type CaseContext struct {
	*ExpressionContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(COOLParserCASE, 0)
}

func (s *CaseContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CaseContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseContext) OF() antlr.TerminalNode {
	return s.GetToken(COOLParserOF, 0)
}

func (s *CaseContext) ESAC() antlr.TerminalNode {
	return s.GetToken(COOLParserESAC, 0)
}

func (s *CaseContext) AllOBJECTID() []antlr.TerminalNode {
	return s.GetTokens(COOLParserOBJECTID)
}

func (s *CaseContext) OBJECTID(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, i)
}

func (s *CaseContext) AllTYPEID() []antlr.TerminalNode {
	return s.GetTokens(COOLParserTYPEID)
}

func (s *CaseContext) TYPEID(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, i)
}

func (s *CaseContext) AllCASE_ARROW() []antlr.TerminalNode {
	return s.GetTokens(COOLParserCASE_ARROW)
}

func (s *CaseContext) CASE_ARROW(i int) antlr.TerminalNode {
	return s.GetToken(COOLParserCASE_ARROW, i)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitCase(s)
	}
}

type OwnMethodCallContext struct {
	*ExpressionContext
}

func NewOwnMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OwnMethodCallContext {
	var p = new(OwnMethodCallContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OwnMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OwnMethodCallContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *OwnMethodCallContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OwnMethodCallContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OwnMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterOwnMethodCall(s)
	}
}

func (s *OwnMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitOwnMethodCall(s)
	}
}

type AddContext struct {
	*ExpressionContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddContext) ADD() antlr.TerminalNode {
	return s.GetToken(COOLParserADD, 0)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitAdd(s)
	}
}

type NewContext struct {
	*ExpressionContext
}

func NewNewContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewContext {
	var p = new(NewContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewContext) NEW() antlr.TerminalNode {
	return s.GetToken(COOLParserNEW, 0)
}

func (s *NewContext) TYPEID() antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, 0)
}

func (s *NewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterNew(s)
	}
}

func (s *NewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitNew(s)
	}
}

type ParenthesesContext struct {
	*ExpressionContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitParentheses(s)
	}
}

type AssignmentContext struct {
	*ExpressionContext
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *AssignmentContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(COOLParserASSIGNMENT, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitAssignment(s)
	}
}

type FalseContext struct {
	*ExpressionContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(COOLParserFALSE, 0)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitFalse(s)
	}
}

type IntContext struct {
	*ExpressionContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(COOLParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitInt(s)
	}
}

type EqualContext struct {
	*ExpressionContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(COOLParserEQUAL, 0)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitEqual(s)
	}
}

type TrueContext struct {
	*ExpressionContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(COOLParserTRUE, 0)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitTrue(s)
	}
}

type LessEqualContext struct {
	*ExpressionContext
}

func NewLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessEqualContext {
	var p = new(LessEqualContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessEqualContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessEqualContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LessEqualContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(COOLParserLESS_EQUAL, 0)
}

func (s *LessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterLessEqual(s)
	}
}

func (s *LessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitLessEqual(s)
	}
}

type MethodCallContext struct {
	*ExpressionContext
}

func NewMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallContext {
	var p = new(MethodCallContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MethodCallContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MethodCallContext) OBJECTID() antlr.TerminalNode {
	return s.GetToken(COOLParserOBJECTID, 0)
}

func (s *MethodCallContext) TYPEID() antlr.TerminalNode {
	return s.GetToken(COOLParserTYPEID, 0)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(COOLListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (p *COOLParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *COOLParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, COOLParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOwnMethodCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(71)
			p.Match(COOLParserOBJECTID)
		}
		{
			p.SetState(72)
			p.Match(COOLParserT__3)
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<COOLParserT__1)|(1<<COOLParserT__3)|(1<<COOLParserFALSE)|(1<<COOLParserIF)|(1<<COOLParserISVOID)|(1<<COOLParserLET)|(1<<COOLParserWHILE)|(1<<COOLParserCASE)|(1<<COOLParserNEW)|(1<<COOLParserNOT)|(1<<COOLParserTRUE)|(1<<COOLParserSTRING)|(1<<COOLParserINT))) != 0) || _la == COOLParserOBJECTID || _la == COOLParserINTEGER_NEGATIVE {
			{
				p.SetState(73)
				p.expression(0)
			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == COOLParserT__4 {
				{
					p.SetState(74)
					p.Match(COOLParserT__4)
				}
				{
					p.SetState(75)
					p.expression(0)
				}

				p.SetState(80)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(83)
			p.Match(COOLParserT__5)
		}

	case 2:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(COOLParserIF)
		}
		{
			p.SetState(85)
			p.expression(0)
		}
		{
			p.SetState(86)
			p.Match(COOLParserTHEN)
		}
		{
			p.SetState(87)
			p.expression(0)
		}
		{
			p.SetState(88)
			p.Match(COOLParserELSE)
		}
		{
			p.SetState(89)
			p.expression(0)
		}
		{
			p.SetState(90)
			p.Match(COOLParserFI)
		}

	case 3:
		localctx = NewWhileContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.Match(COOLParserWHILE)
		}
		{
			p.SetState(93)
			p.expression(0)
		}
		{
			p.SetState(94)
			p.Match(COOLParserLOOP)
		}
		{
			p.SetState(95)
			p.expression(0)
		}
		{
			p.SetState(96)
			p.Match(COOLParserPOOL)
		}

	case 4:
		localctx = NewBlockContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(98)
			p.Match(COOLParserT__1)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<COOLParserT__1)|(1<<COOLParserT__3)|(1<<COOLParserFALSE)|(1<<COOLParserIF)|(1<<COOLParserISVOID)|(1<<COOLParserLET)|(1<<COOLParserWHILE)|(1<<COOLParserCASE)|(1<<COOLParserNEW)|(1<<COOLParserNOT)|(1<<COOLParserTRUE)|(1<<COOLParserSTRING)|(1<<COOLParserINT))) != 0) || _la == COOLParserOBJECTID || _la == COOLParserINTEGER_NEGATIVE {
			{
				p.SetState(99)
				p.expression(0)
			}
			{
				p.SetState(100)
				p.Match(COOLParserT__0)
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.Match(COOLParserT__2)
		}

	case 5:
		localctx = NewLetInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(108)
			p.Match(COOLParserLET)
		}
		{
			p.SetState(109)
			p.Match(COOLParserOBJECTID)
		}
		{
			p.SetState(110)
			p.Match(COOLParserT__6)
		}
		{
			p.SetState(111)
			p.Match(COOLParserTYPEID)
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == COOLParserASSIGNMENT {
			{
				p.SetState(112)
				p.Match(COOLParserASSIGNMENT)
			}
			{
				p.SetState(113)
				p.expression(0)
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == COOLParserT__4 {
			{
				p.SetState(116)
				p.Match(COOLParserT__4)
			}
			{
				p.SetState(117)
				p.Match(COOLParserOBJECTID)
			}
			{
				p.SetState(118)
				p.Match(COOLParserT__6)
			}
			{
				p.SetState(119)
				p.Match(COOLParserTYPEID)
			}
			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == COOLParserASSIGNMENT {
				{
					p.SetState(120)
					p.Match(COOLParserASSIGNMENT)
				}
				{
					p.SetState(121)
					p.expression(0)
				}

			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(COOLParserIN)
		}
		{
			p.SetState(130)
			p.expression(20)
		}

	case 6:
		localctx = NewCaseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(COOLParserCASE)
		}
		{
			p.SetState(132)
			p.expression(0)
		}
		{
			p.SetState(133)
			p.Match(COOLParserOF)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == COOLParserOBJECTID {
			{
				p.SetState(134)
				p.Match(COOLParserOBJECTID)
			}
			{
				p.SetState(135)
				p.Match(COOLParserT__6)
			}
			{
				p.SetState(136)
				p.Match(COOLParserTYPEID)
			}
			{
				p.SetState(137)
				p.Match(COOLParserCASE_ARROW)
			}
			{
				p.SetState(138)
				p.expression(0)
			}
			{
				p.SetState(139)
				p.Match(COOLParserT__0)
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(145)
			p.Match(COOLParserESAC)
		}

	case 7:
		localctx = NewNewContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(147)
			p.Match(COOLParserNEW)
		}
		{
			p.SetState(148)
			p.Match(COOLParserTYPEID)
		}

	case 8:
		localctx = NewNegativeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.Match(COOLParserINTEGER_NEGATIVE)
		}
		{
			p.SetState(150)
			p.expression(17)
		}

	case 9:
		localctx = NewIsvoidContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(151)
			p.Match(COOLParserISVOID)
		}
		{
			p.SetState(152)
			p.expression(16)
		}

	case 10:
		localctx = NewBoolNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(COOLParserNOT)
		}
		{
			p.SetState(154)
			p.expression(8)
		}

	case 11:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(155)
			p.Match(COOLParserT__3)
		}
		{
			p.SetState(156)
			p.expression(0)
		}
		{
			p.SetState(157)
			p.Match(COOLParserT__5)
		}

	case 12:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(COOLParserOBJECTID)
		}

	case 13:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(COOLParserINT)
		}

	case 14:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(161)
			p.Match(COOLParserSTRING)
		}

	case 15:
		localctx = NewTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(162)
			p.Match(COOLParserTRUE)
		}

	case 16:
		localctx = NewFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.Match(COOLParserFALSE)
		}

	case 17:
		localctx = NewAssignmentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(COOLParserOBJECTID)
		}
		{
			p.SetState(165)
			p.Match(COOLParserASSIGNMENT)
		}
		{
			p.SetState(166)
			p.expression(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(170)
					p.Match(COOLParserMULTIPLY)
				}
				{
					p.SetState(171)
					p.expression(16)
				}

			case 2:
				localctx = NewDivisionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(173)
					p.Match(COOLParserDIVISION)
				}
				{
					p.SetState(174)
					p.expression(15)
				}

			case 3:
				localctx = NewAddContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(176)
					p.Match(COOLParserADD)
				}
				{
					p.SetState(177)
					p.expression(14)
				}

			case 4:
				localctx = NewMinusContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(179)
					p.Match(COOLParserMINUS)
				}
				{
					p.SetState(180)
					p.expression(13)
				}

			case 5:
				localctx = NewLessThanContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(181)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(182)
					p.Match(COOLParserLESS_THAN)
				}
				{
					p.SetState(183)
					p.expression(12)
				}

			case 6:
				localctx = NewLessEqualContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(185)
					p.Match(COOLParserLESS_EQUAL)
				}
				{
					p.SetState(186)
					p.expression(11)
				}

			case 7:
				localctx = NewEqualContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(188)
					p.Match(COOLParserEQUAL)
				}
				{
					p.SetState(189)
					p.expression(10)
				}

			case 8:
				localctx = NewMethodCallContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, COOLParserRULE_expression)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				p.SetState(193)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == COOLParserT__7 {
					{
						p.SetState(191)
						p.Match(COOLParserT__7)
					}
					{
						p.SetState(192)
						p.Match(COOLParserTYPEID)
					}

				}
				{
					p.SetState(195)
					p.Match(COOLParserT__8)
				}
				{
					p.SetState(196)
					p.Match(COOLParserOBJECTID)
				}
				{
					p.SetState(197)
					p.Match(COOLParserT__3)
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<COOLParserT__1)|(1<<COOLParserT__3)|(1<<COOLParserFALSE)|(1<<COOLParserIF)|(1<<COOLParserISVOID)|(1<<COOLParserLET)|(1<<COOLParserWHILE)|(1<<COOLParserCASE)|(1<<COOLParserNEW)|(1<<COOLParserNOT)|(1<<COOLParserTRUE)|(1<<COOLParserSTRING)|(1<<COOLParserINT))) != 0) || _la == COOLParserOBJECTID || _la == COOLParserINTEGER_NEGATIVE {
					{
						p.SetState(198)
						p.expression(0)
					}
					p.SetState(203)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == COOLParserT__4 {
						{
							p.SetState(199)
							p.Match(COOLParserT__4)
						}
						{
							p.SetState(200)
							p.expression(0)
						}

						p.SetState(205)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(208)
					p.Match(COOLParserT__5)
				}

			}

		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

func (p *COOLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *COOLParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 25)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
