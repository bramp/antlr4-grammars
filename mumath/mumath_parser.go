// Code generated from mumath.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mumath // mumath
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 204,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 5, 2, 44, 10, 2, 3, 2,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	6, 3, 57, 10, 3, 13, 3, 14, 3, 58, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 68, 10, 4, 12, 4, 14, 4, 71, 11, 4, 3, 4, 5, 4, 74, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 82, 10, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 5, 6, 88, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10,
	7, 3, 7, 3, 7, 7, 7, 100, 10, 7, 12, 7, 14, 7, 103, 11, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 113, 10, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 5, 10, 120, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 126,
	10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 134, 10, 11, 12,
	11, 14, 11, 137, 11, 11, 3, 12, 3, 12, 5, 12, 141, 10, 12, 3, 13, 5, 13,
	144, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 150, 10, 13, 12, 13, 14,
	13, 153, 11, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 161,
	10, 15, 12, 15, 14, 15, 164, 11, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 177, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 185, 10, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 7, 19, 192, 10, 19, 12, 19, 14, 19, 195, 11, 19, 3, 19, 5,
	19, 198, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 2, 2, 21, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 7, 3,
	2, 25, 26, 4, 2, 18, 18, 29, 33, 4, 2, 12, 12, 20, 21, 6, 2, 13, 13, 15,
	15, 22, 23, 36, 36, 4, 2, 7, 7, 28, 28, 2, 215, 2, 49, 3, 2, 2, 2, 4, 56,
	3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 87, 3, 2, 2, 2, 12,
	95, 3, 2, 2, 2, 14, 104, 3, 2, 2, 2, 16, 109, 3, 2, 2, 2, 18, 116, 3, 2,
	2, 2, 20, 129, 3, 2, 2, 2, 22, 140, 3, 2, 2, 2, 24, 143, 3, 2, 2, 2, 26,
	154, 3, 2, 2, 2, 28, 156, 3, 2, 2, 2, 30, 165, 3, 2, 2, 2, 32, 176, 3,
	2, 2, 2, 34, 184, 3, 2, 2, 2, 36, 186, 3, 2, 2, 2, 38, 201, 3, 2, 2, 2,
	40, 44, 5, 8, 5, 2, 41, 44, 5, 4, 3, 2, 42, 44, 5, 36, 19, 2, 43, 40, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45,
	46, 9, 2, 2, 2, 46, 48, 3, 2, 2, 2, 47, 43, 3, 2, 2, 2, 48, 51, 3, 2, 2,
	2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 52, 53, 7, 2, 2, 3, 53, 3, 3, 2, 2, 2, 54, 55, 7, 37, 2, 2,
	55, 57, 7, 27, 2, 2, 56, 54, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3,
	2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 5, 20, 11, 2,
	61, 5, 3, 2, 2, 2, 62, 73, 7, 34, 2, 2, 63, 74, 7, 35, 2, 2, 64, 69, 7,
	37, 2, 2, 65, 66, 7, 24, 2, 2, 66, 68, 7, 37, 2, 2, 67, 65, 3, 2, 2, 2,
	68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3,
	2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 74, 7, 35, 2, 2, 73, 63, 3, 2, 2, 2, 73,
	64, 3, 2, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7, 5, 2, 2, 76, 77, 7, 37, 2,
	2, 77, 78, 5, 6, 4, 2, 78, 79, 7, 24, 2, 2, 79, 81, 5, 12, 7, 2, 80, 82,
	7, 24, 2, 2, 81, 80, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 84, 7, 6, 2, 2, 84, 9, 3, 2, 2, 2, 85, 88, 5, 20, 11, 2, 86, 88, 5,
	4, 3, 2, 87, 85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 11, 3, 2, 2, 2, 89,
	96, 5, 16, 9, 2, 90, 96, 5, 18, 10, 2, 91, 96, 5, 14, 8, 2, 92, 96, 5,
	4, 3, 2, 93, 96, 5, 20, 11, 2, 94, 96, 5, 36, 19, 2, 95, 89, 3, 2, 2, 2,
	95, 90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3,
	2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 101, 3, 2, 2, 2, 97, 98, 7, 24, 2, 2,
	98, 100, 5, 12, 7, 2, 99, 97, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 13, 3, 2, 2, 2, 103, 101, 3, 2,
	2, 2, 104, 105, 7, 3, 2, 2, 105, 106, 5, 12, 7, 2, 106, 107, 7, 24, 2,
	2, 107, 108, 7, 4, 2, 2, 108, 15, 3, 2, 2, 2, 109, 110, 7, 8, 2, 2, 110,
	112, 5, 12, 7, 2, 111, 113, 7, 24, 2, 2, 112, 111, 3, 2, 2, 2, 112, 113,
	3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 7, 9, 2, 2, 115, 17, 3, 2,
	2, 2, 116, 117, 7, 10, 2, 2, 117, 119, 5, 20, 11, 2, 118, 120, 7, 24, 2,
	2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	122, 7, 11, 2, 2, 122, 123, 7, 24, 2, 2, 123, 125, 5, 12, 7, 2, 124, 126,
	7, 24, 2, 2, 125, 124, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2,
	2, 2, 127, 128, 7, 11, 2, 2, 128, 19, 3, 2, 2, 2, 129, 135, 5, 24, 13,
	2, 130, 131, 5, 22, 12, 2, 131, 132, 5, 24, 13, 2, 132, 134, 3, 2, 2, 2,
	133, 130, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 21, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 141, 5,
	38, 20, 2, 139, 141, 9, 3, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2,
	2, 2, 141, 23, 3, 2, 2, 2, 142, 144, 7, 21, 2, 2, 143, 142, 3, 2, 2, 2,
	143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 151, 5, 28, 15, 2, 146,
	147, 5, 26, 14, 2, 147, 148, 5, 28, 15, 2, 148, 150, 3, 2, 2, 2, 149, 146,
	3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2,
	2, 2, 152, 25, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 9, 4, 2, 2,
	155, 27, 3, 2, 2, 2, 156, 162, 5, 32, 17, 2, 157, 158, 5, 30, 16, 2, 158,
	159, 5, 32, 17, 2, 159, 161, 3, 2, 2, 2, 160, 157, 3, 2, 2, 2, 161, 164,
	3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 29, 3, 2,
	2, 2, 164, 162, 3, 2, 2, 2, 165, 166, 9, 5, 2, 2, 166, 31, 3, 2, 2, 2,
	167, 177, 7, 37, 2, 2, 168, 177, 5, 34, 18, 2, 169, 170, 7, 34, 2, 2, 170,
	171, 5, 20, 11, 2, 171, 172, 7, 35, 2, 2, 172, 177, 3, 2, 2, 2, 173, 177,
	5, 36, 19, 2, 174, 175, 7, 14, 2, 2, 175, 177, 5, 32, 17, 2, 176, 167,
	3, 2, 2, 2, 176, 168, 3, 2, 2, 2, 176, 169, 3, 2, 2, 2, 176, 173, 3, 2,
	2, 2, 176, 174, 3, 2, 2, 2, 177, 33, 3, 2, 2, 2, 178, 185, 7, 40, 2, 2,
	179, 185, 7, 39, 2, 2, 180, 181, 7, 19, 2, 2, 181, 185, 7, 37, 2, 2, 182,
	183, 7, 19, 2, 2, 183, 185, 7, 39, 2, 2, 184, 178, 3, 2, 2, 2, 184, 179,
	3, 2, 2, 2, 184, 180, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 35, 3, 2,
	2, 2, 186, 187, 7, 37, 2, 2, 187, 197, 7, 34, 2, 2, 188, 193, 5, 10, 6,
	2, 189, 190, 7, 24, 2, 2, 190, 192, 5, 10, 6, 2, 191, 189, 3, 2, 2, 2,
	192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194,
	198, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197, 188,
	3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 7, 35,
	2, 2, 200, 37, 3, 2, 2, 2, 201, 202, 9, 6, 2, 2, 202, 39, 3, 2, 2, 2, 23,
	43, 49, 58, 69, 73, 81, 87, 95, 101, 112, 119, 125, 135, 140, 143, 151,
	162, 176, 184, 193, 197,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'BLOCK'", "'ENDBLOCK'", "'FUNCTION'", "'ENDFUN'", "'EQ'", "'LOOP'",
	"'ENDLOOP'", "'WHEN'", "'EXIT'", "'OR'", "'AND'", "'NOT'", "'mod'", "",
	"", "'=='", "'''", "'+'", "'-'", "'*'", "'/'", "','", "';'", "'$'", "':'",
	"'='", "'<>'", "'<'", "'<='", "'>='", "'>'", "'('", "')'", "'^'",
}
var symbolicNames = []string{
	"", "BLOCK", "ENDBLOCK", "FUNCTION", "ENDFUN", "EQF", "LOOP", "ENDLOOP",
	"WHEN", "EXIT", "OR", "AND", "NOT", "MOD", "WS", "COMMENT", "EQUATION",
	"QUOTE", "PLUS", "MINUS", "STAR", "SLASH", "COMMA", "SEMI", "DOLLAR", "COLON",
	"EQC", "NOT_EQUAL", "LT", "LE", "GE", "GT", "LPAREN", "RPAREN", "POWER",
	"ID", "ARR", "STRING", "NUMBER",
}

var ruleNames = []string{
	"program", "assignment", "list", "functionDefinition", "actualParameter",
	"statments", "block", "loop", "when", "expression", "relationalOperator",
	"simpleExpression", "addingOperator", "term", "multiplyingOperator", "factor",
	"constant", "functionDesignator", "equal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type mumathParser struct {
	*antlr.BaseParser
}

func NewmumathParser(input antlr.TokenStream) *mumathParser {
	this := new(mumathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "mumath.g4"

	return this
}

// mumathParser tokens.
const (
	mumathParserEOF       = antlr.TokenEOF
	mumathParserBLOCK     = 1
	mumathParserENDBLOCK  = 2
	mumathParserFUNCTION  = 3
	mumathParserENDFUN    = 4
	mumathParserEQF       = 5
	mumathParserLOOP      = 6
	mumathParserENDLOOP   = 7
	mumathParserWHEN      = 8
	mumathParserEXIT      = 9
	mumathParserOR        = 10
	mumathParserAND       = 11
	mumathParserNOT       = 12
	mumathParserMOD       = 13
	mumathParserWS        = 14
	mumathParserCOMMENT   = 15
	mumathParserEQUATION  = 16
	mumathParserQUOTE     = 17
	mumathParserPLUS      = 18
	mumathParserMINUS     = 19
	mumathParserSTAR      = 20
	mumathParserSLASH     = 21
	mumathParserCOMMA     = 22
	mumathParserSEMI      = 23
	mumathParserDOLLAR    = 24
	mumathParserCOLON     = 25
	mumathParserEQC       = 26
	mumathParserNOT_EQUAL = 27
	mumathParserLT        = 28
	mumathParserLE        = 29
	mumathParserGE        = 30
	mumathParserGT        = 31
	mumathParserLPAREN    = 32
	mumathParserRPAREN    = 33
	mumathParserPOWER     = 34
	mumathParserID        = 35
	mumathParserARR       = 36
	mumathParserSTRING    = 37
	mumathParserNUMBER    = 38
)

// mumathParser rules.
const (
	mumathParserRULE_program             = 0
	mumathParserRULE_assignment          = 1
	mumathParserRULE_list                = 2
	mumathParserRULE_functionDefinition  = 3
	mumathParserRULE_actualParameter     = 4
	mumathParserRULE_statments           = 5
	mumathParserRULE_block               = 6
	mumathParserRULE_loop                = 7
	mumathParserRULE_when                = 8
	mumathParserRULE_expression          = 9
	mumathParserRULE_relationalOperator  = 10
	mumathParserRULE_simpleExpression    = 11
	mumathParserRULE_addingOperator      = 12
	mumathParserRULE_term                = 13
	mumathParserRULE_multiplyingOperator = 14
	mumathParserRULE_factor              = 15
	mumathParserRULE_constant            = 16
	mumathParserRULE_functionDesignator  = 17
	mumathParserRULE_equal               = 18
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
	p.RuleIndex = mumathParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(mumathParserEOF, 0)
}

func (s *ProgramContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(mumathParserSEMI)
}

func (s *ProgramContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserSEMI, i)
}

func (s *ProgramContext) AllDOLLAR() []antlr.TerminalNode {
	return s.GetTokens(mumathParserDOLLAR)
}

func (s *ProgramContext) DOLLAR(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserDOLLAR, i)
}

func (s *ProgramContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem())
	var tst = make([]IFunctionDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDefinitionContext)
		}
	}

	return tst
}

func (s *ProgramContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *ProgramContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ProgramContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ProgramContext) AllFunctionDesignator() []IFunctionDesignatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDesignatorContext)(nil)).Elem())
	var tst = make([]IFunctionDesignatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDesignatorContext)
		}
	}

	return tst
}

func (s *ProgramContext) FunctionDesignator(i int) IFunctionDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDesignatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDesignatorContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *mumathParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mumathParserRULE_program)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumathParserFUNCTION || _la == mumathParserID {
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(38)
				p.FunctionDefinition()
			}

		case 2:
			{
				p.SetState(39)
				p.Assignment()
			}

		case 3:
			{
				p.SetState(40)
				p.FunctionDesignator()
			}

		}
		{
			p.SetState(43)
			_la = p.GetTokenStream().LA(1)

			if !(_la == mumathParserSEMI || _la == mumathParserDOLLAR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(mumathParserEOF)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(mumathParserID)
}

func (s *AssignmentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserID, i)
}

func (s *AssignmentContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOLON)
}

func (s *AssignmentContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOLON, i)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *mumathParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mumathParserRULE_assignment)

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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(52)
				p.Match(mumathParserID)
			}
			{
				p.SetState(53)
				p.Match(mumathParserCOLON)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	{
		p.SetState(58)
		p.Expression()
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserLPAREN, 0)
}

func (s *ListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserRPAREN, 0)
}

func (s *ListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(mumathParserID)
}

func (s *ListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserID, i)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *mumathParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mumathParserRULE_list)
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
		p.SetState(60)
		p.Match(mumathParserLPAREN)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumathParserRPAREN:
		{
			p.SetState(61)
			p.Match(mumathParserRPAREN)
		}

	case mumathParserID:
		{
			p.SetState(62)
			p.Match(mumathParserID)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == mumathParserCOMMA {
			{
				p.SetState(63)
				p.Match(mumathParserCOMMA)
			}
			{
				p.SetState(64)
				p.Match(mumathParserID)
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(70)
			p.Match(mumathParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(mumathParserFUNCTION, 0)
}

func (s *FunctionDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(mumathParserID, 0)
}

func (s *FunctionDefinitionContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *FunctionDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOMMA)
}

func (s *FunctionDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, i)
}

func (s *FunctionDefinitionContext) Statments() IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
}

func (s *FunctionDefinitionContext) ENDFUN() antlr.TerminalNode {
	return s.GetToken(mumathParserENDFUN, 0)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *mumathParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mumathParserRULE_functionDefinition)
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
		p.SetState(73)
		p.Match(mumathParserFUNCTION)
	}
	{
		p.SetState(74)
		p.Match(mumathParserID)
	}
	{
		p.SetState(75)
		p.List()
	}
	{
		p.SetState(76)
		p.Match(mumathParserCOMMA)
	}
	{
		p.SetState(77)
		p.Statments()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumathParserCOMMA {
		{
			p.SetState(78)
			p.Match(mumathParserCOMMA)
		}

	}
	{
		p.SetState(81)
		p.Match(mumathParserENDFUN)
	}

	return localctx
}

// IActualParameterContext is an interface to support dynamic dispatch.
type IActualParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActualParameterContext differentiates from other interfaces.
	IsActualParameterContext()
}

type ActualParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActualParameterContext() *ActualParameterContext {
	var p = new(ActualParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_actualParameter
	return p
}

func (*ActualParameterContext) IsActualParameterContext() {}

func NewActualParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActualParameterContext {
	var p = new(ActualParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_actualParameter

	return p
}

func (s *ActualParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ActualParameterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ActualParameterContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ActualParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActualParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActualParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterActualParameter(s)
	}
}

func (s *ActualParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitActualParameter(s)
	}
}

func (p *mumathParser) ActualParameter() (localctx IActualParameterContext) {
	localctx = NewActualParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mumathParserRULE_actualParameter)

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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Assignment()
		}

	}

	return localctx
}

// IStatmentsContext is an interface to support dynamic dispatch.
type IStatmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatmentsContext differentiates from other interfaces.
	IsStatmentsContext()
}

type StatmentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatmentsContext() *StatmentsContext {
	var p = new(StatmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_statments
	return p
}

func (*StatmentsContext) IsStatmentsContext() {}

func NewStatmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatmentsContext {
	var p = new(StatmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_statments

	return p
}

func (s *StatmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatmentsContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatmentsContext) When() IWhenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhenContext)
}

func (s *StatmentsContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatmentsContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatmentsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatmentsContext) FunctionDesignator() IFunctionDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDesignatorContext)
}

func (s *StatmentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOMMA)
}

func (s *StatmentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, i)
}

func (s *StatmentsContext) AllStatments() []IStatmentsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatmentsContext)(nil)).Elem())
	var tst = make([]IStatmentsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatmentsContext)
		}
	}

	return tst
}

func (s *StatmentsContext) Statments(i int) IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
}

func (s *StatmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterStatments(s)
	}
}

func (s *StatmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitStatments(s)
	}
}

func (p *mumathParser) Statments() (localctx IStatmentsContext) {
	localctx = NewStatmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mumathParserRULE_statments)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(87)
			p.Loop()
		}

	case 2:
		{
			p.SetState(88)
			p.When()
		}

	case 3:
		{
			p.SetState(89)
			p.Block()
		}

	case 4:
		{
			p.SetState(90)
			p.Assignment()
		}

	case 5:
		{
			p.SetState(91)
			p.Expression()
		}

	case 6:
		{
			p.SetState(92)
			p.FunctionDesignator()
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(95)
				p.Match(mumathParserCOMMA)
			}
			{
				p.SetState(96)
				p.Statments()
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BLOCK() antlr.TerminalNode {
	return s.GetToken(mumathParserBLOCK, 0)
}

func (s *BlockContext) Statments() IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
}

func (s *BlockContext) COMMA() antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, 0)
}

func (s *BlockContext) ENDBLOCK() antlr.TerminalNode {
	return s.GetToken(mumathParserENDBLOCK, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *mumathParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mumathParserRULE_block)

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
		p.Match(mumathParserBLOCK)
	}
	{
		p.SetState(103)
		p.Statments()
	}
	{
		p.SetState(104)
		p.Match(mumathParserCOMMA)
	}
	{
		p.SetState(105)
		p.Match(mumathParserENDBLOCK)
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) LOOP() antlr.TerminalNode {
	return s.GetToken(mumathParserLOOP, 0)
}

func (s *LoopContext) Statments() IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
}

func (s *LoopContext) ENDLOOP() antlr.TerminalNode {
	return s.GetToken(mumathParserENDLOOP, 0)
}

func (s *LoopContext) COMMA() antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, 0)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *mumathParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mumathParserRULE_loop)
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
		p.SetState(107)
		p.Match(mumathParserLOOP)
	}
	{
		p.SetState(108)
		p.Statments()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumathParserCOMMA {
		{
			p.SetState(109)
			p.Match(mumathParserCOMMA)
		}

	}
	{
		p.SetState(112)
		p.Match(mumathParserENDLOOP)
	}

	return localctx
}

// IWhenContext is an interface to support dynamic dispatch.
type IWhenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhenContext differentiates from other interfaces.
	IsWhenContext()
}

type WhenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhenContext() *WhenContext {
	var p = new(WhenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_when
	return p
}

func (*WhenContext) IsWhenContext() {}

func NewWhenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenContext {
	var p = new(WhenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_when

	return p
}

func (s *WhenContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenContext) WHEN() antlr.TerminalNode {
	return s.GetToken(mumathParserWHEN, 0)
}

func (s *WhenContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhenContext) AllEXIT() []antlr.TerminalNode {
	return s.GetTokens(mumathParserEXIT)
}

func (s *WhenContext) EXIT(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserEXIT, i)
}

func (s *WhenContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOMMA)
}

func (s *WhenContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, i)
}

func (s *WhenContext) Statments() IStatmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatmentsContext)
}

func (s *WhenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterWhen(s)
	}
}

func (s *WhenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitWhen(s)
	}
}

func (p *mumathParser) When() (localctx IWhenContext) {
	localctx = NewWhenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mumathParserRULE_when)
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
		p.SetState(114)
		p.Match(mumathParserWHEN)
	}
	{
		p.SetState(115)
		p.Expression()
	}

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumathParserCOMMA {
		{
			p.SetState(116)
			p.Match(mumathParserCOMMA)
		}

	}
	{
		p.SetState(119)
		p.Match(mumathParserEXIT)
	}
	{
		p.SetState(120)
		p.Match(mumathParserCOMMA)
	}
	{
		p.SetState(121)
		p.Statments()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumathParserCOMMA {
		{
			p.SetState(122)
			p.Match(mumathParserCOMMA)
		}

	}
	{
		p.SetState(125)
		p.Match(mumathParserEXIT)
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
	p.RuleIndex = mumathParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllSimpleExpression() []ISimpleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimpleExpressionContext)(nil)).Elem())
	var tst = make([]ISimpleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimpleExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) SimpleExpression(i int) ISimpleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimpleExpressionContext)
}

func (s *ExpressionContext) AllRelationalOperator() []IRelationalOperatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationalOperatorContext)(nil)).Elem())
	var tst = make([]IRelationalOperatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationalOperatorContext)
		}
	}

	return tst
}

func (s *ExpressionContext) RelationalOperator(i int) IRelationalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalOperatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *mumathParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mumathParserRULE_expression)
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
		p.SetState(127)
		p.SimpleExpression()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumathParserEQF)|(1<<mumathParserEQUATION)|(1<<mumathParserEQC)|(1<<mumathParserNOT_EQUAL)|(1<<mumathParserLT)|(1<<mumathParserLE)|(1<<mumathParserGE)|(1<<mumathParserGT))) != 0 {
		{
			p.SetState(128)
			p.RelationalOperator()
		}
		{
			p.SetState(129)
			p.SimpleExpression()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationalOperatorContext is an interface to support dynamic dispatch.
type IRelationalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalOperatorContext differentiates from other interfaces.
	IsRelationalOperatorContext()
}

type RelationalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOperatorContext() *RelationalOperatorContext {
	var p = new(RelationalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_relationalOperator
	return p
}

func (*RelationalOperatorContext) IsRelationalOperatorContext() {}

func NewRelationalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOperatorContext {
	var p = new(RelationalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_relationalOperator

	return p
}

func (s *RelationalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOperatorContext) Equal() IEqualContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualContext)
}

func (s *RelationalOperatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(mumathParserNOT_EQUAL, 0)
}

func (s *RelationalOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(mumathParserLT, 0)
}

func (s *RelationalOperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(mumathParserLE, 0)
}

func (s *RelationalOperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(mumathParserGE, 0)
}

func (s *RelationalOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(mumathParserGT, 0)
}

func (s *RelationalOperatorContext) EQUATION() antlr.TerminalNode {
	return s.GetToken(mumathParserEQUATION, 0)
}

func (s *RelationalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterRelationalOperator(s)
	}
}

func (s *RelationalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitRelationalOperator(s)
	}
}

func (p *mumathParser) RelationalOperator() (localctx IRelationalOperatorContext) {
	localctx = NewRelationalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, mumathParserRULE_relationalOperator)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumathParserEQF, mumathParserEQC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Equal()
		}

	case mumathParserEQUATION, mumathParserNOT_EQUAL, mumathParserLT, mumathParserLE, mumathParserGE, mumathParserGT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumathParserEQUATION)|(1<<mumathParserNOT_EQUAL)|(1<<mumathParserLT)|(1<<mumathParserLE)|(1<<mumathParserGE)|(1<<mumathParserGT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISimpleExpressionContext is an interface to support dynamic dispatch.
type ISimpleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleExpressionContext differentiates from other interfaces.
	IsSimpleExpressionContext()
}

type SimpleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleExpressionContext() *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_simpleExpression
	return p
}

func (*SimpleExpressionContext) IsSimpleExpressionContext() {}

func NewSimpleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_simpleExpression

	return p
}

func (s *SimpleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleExpressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *SimpleExpressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(mumathParserMINUS, 0)
}

func (s *SimpleExpressionContext) AllAddingOperator() []IAddingOperatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddingOperatorContext)(nil)).Elem())
	var tst = make([]IAddingOperatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddingOperatorContext)
		}
	}

	return tst
}

func (s *SimpleExpressionContext) AddingOperator(i int) IAddingOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddingOperatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddingOperatorContext)
}

func (s *SimpleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitSimpleExpression(s)
	}
}

func (p *mumathParser) SimpleExpression() (localctx ISimpleExpressionContext) {
	localctx = NewSimpleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mumathParserRULE_simpleExpression)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumathParserMINUS {
		{
			p.SetState(140)
			p.Match(mumathParserMINUS)
		}

	}
	{
		p.SetState(143)
		p.Term()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumathParserOR)|(1<<mumathParserPLUS)|(1<<mumathParserMINUS))) != 0 {
		{
			p.SetState(144)
			p.AddingOperator()
		}
		{
			p.SetState(145)
			p.Term()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAddingOperatorContext is an interface to support dynamic dispatch.
type IAddingOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddingOperatorContext differentiates from other interfaces.
	IsAddingOperatorContext()
}

type AddingOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddingOperatorContext() *AddingOperatorContext {
	var p = new(AddingOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_addingOperator
	return p
}

func (*AddingOperatorContext) IsAddingOperatorContext() {}

func NewAddingOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddingOperatorContext {
	var p = new(AddingOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_addingOperator

	return p
}

func (s *AddingOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AddingOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(mumathParserPLUS, 0)
}

func (s *AddingOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(mumathParserMINUS, 0)
}

func (s *AddingOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(mumathParserOR, 0)
}

func (s *AddingOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddingOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddingOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterAddingOperator(s)
	}
}

func (s *AddingOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitAddingOperator(s)
	}
}

func (p *mumathParser) AddingOperator() (localctx IAddingOperatorContext) {
	localctx = NewAddingOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mumathParserRULE_addingOperator)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumathParserOR)|(1<<mumathParserPLUS)|(1<<mumathParserMINUS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMultiplyingOperator() []IMultiplyingOperatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingOperatorContext)(nil)).Elem())
	var tst = make([]IMultiplyingOperatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingOperatorContext)
		}
	}

	return tst
}

func (s *TermContext) MultiplyingOperator(i int) IMultiplyingOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingOperatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingOperatorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *mumathParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mumathParserRULE_term)
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
		p.SetState(154)
		p.Factor()
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(mumathParserAND-11))|(1<<(mumathParserMOD-11))|(1<<(mumathParserSTAR-11))|(1<<(mumathParserSLASH-11))|(1<<(mumathParserPOWER-11)))) != 0 {
		{
			p.SetState(155)
			p.MultiplyingOperator()
		}
		{
			p.SetState(156)
			p.Factor()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplyingOperatorContext is an interface to support dynamic dispatch.
type IMultiplyingOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingOperatorContext differentiates from other interfaces.
	IsMultiplyingOperatorContext()
}

type MultiplyingOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingOperatorContext() *MultiplyingOperatorContext {
	var p = new(MultiplyingOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_multiplyingOperator
	return p
}

func (*MultiplyingOperatorContext) IsMultiplyingOperatorContext() {}

func NewMultiplyingOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingOperatorContext {
	var p = new(MultiplyingOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_multiplyingOperator

	return p
}

func (s *MultiplyingOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingOperatorContext) STAR() antlr.TerminalNode {
	return s.GetToken(mumathParserSTAR, 0)
}

func (s *MultiplyingOperatorContext) SLASH() antlr.TerminalNode {
	return s.GetToken(mumathParserSLASH, 0)
}

func (s *MultiplyingOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(mumathParserMOD, 0)
}

func (s *MultiplyingOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(mumathParserAND, 0)
}

func (s *MultiplyingOperatorContext) POWER() antlr.TerminalNode {
	return s.GetToken(mumathParserPOWER, 0)
}

func (s *MultiplyingOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterMultiplyingOperator(s)
	}
}

func (s *MultiplyingOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitMultiplyingOperator(s)
	}
}

func (p *mumathParser) MultiplyingOperator() (localctx IMultiplyingOperatorContext) {
	localctx = NewMultiplyingOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mumathParserRULE_multiplyingOperator)
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
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(mumathParserAND-11))|(1<<(mumathParserMOD-11))|(1<<(mumathParserSTAR-11))|(1<<(mumathParserSLASH-11))|(1<<(mumathParserPOWER-11)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(mumathParserID, 0)
}

func (s *FactorContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserLPAREN, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserRPAREN, 0)
}

func (s *FactorContext) FunctionDesignator() IFunctionDesignatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDesignatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDesignatorContext)
}

func (s *FactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(mumathParserNOT, 0)
}

func (s *FactorContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *mumathParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mumathParserRULE_factor)

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

	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(mumathParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Constant()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Match(mumathParserLPAREN)
		}
		{
			p.SetState(168)
			p.Expression()
		}
		{
			p.SetState(169)
			p.Match(mumathParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)
			p.FunctionDesignator()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)
			p.Match(mumathParserNOT)
		}
		{
			p.SetState(173)
			p.Factor()
		}

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
	p.RuleIndex = mumathParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mumathParserNUMBER, 0)
}

func (s *ConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(mumathParserSTRING, 0)
}

func (s *ConstantContext) QUOTE() antlr.TerminalNode {
	return s.GetToken(mumathParserQUOTE, 0)
}

func (s *ConstantContext) ID() antlr.TerminalNode {
	return s.GetToken(mumathParserID, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *mumathParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, mumathParserRULE_constant)

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
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(176)
			p.Match(mumathParserNUMBER)
		}

	case 2:
		{
			p.SetState(177)
			p.Match(mumathParserSTRING)
		}

	case 3:
		{
			p.SetState(178)
			p.Match(mumathParserQUOTE)
		}
		{
			p.SetState(179)
			p.Match(mumathParserID)
		}

	case 4:
		{
			p.SetState(180)
			p.Match(mumathParserQUOTE)
		}
		{
			p.SetState(181)
			p.Match(mumathParserSTRING)
		}

	}

	return localctx
}

// IFunctionDesignatorContext is an interface to support dynamic dispatch.
type IFunctionDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDesignatorContext differentiates from other interfaces.
	IsFunctionDesignatorContext()
}

type FunctionDesignatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDesignatorContext() *FunctionDesignatorContext {
	var p = new(FunctionDesignatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_functionDesignator
	return p
}

func (*FunctionDesignatorContext) IsFunctionDesignatorContext() {}

func NewFunctionDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDesignatorContext {
	var p = new(FunctionDesignatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_functionDesignator

	return p
}

func (s *FunctionDesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDesignatorContext) ID() antlr.TerminalNode {
	return s.GetToken(mumathParserID, 0)
}

func (s *FunctionDesignatorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserLPAREN, 0)
}

func (s *FunctionDesignatorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumathParserRPAREN, 0)
}

func (s *FunctionDesignatorContext) AllActualParameter() []IActualParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IActualParameterContext)(nil)).Elem())
	var tst = make([]IActualParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IActualParameterContext)
		}
	}

	return tst
}

func (s *FunctionDesignatorContext) ActualParameter(i int) IActualParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActualParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IActualParameterContext)
}

func (s *FunctionDesignatorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumathParserCOMMA)
}

func (s *FunctionDesignatorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumathParserCOMMA, i)
}

func (s *FunctionDesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterFunctionDesignator(s)
	}
}

func (s *FunctionDesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitFunctionDesignator(s)
	}
}

func (p *mumathParser) FunctionDesignator() (localctx IFunctionDesignatorContext) {
	localctx = NewFunctionDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, mumathParserRULE_functionDesignator)
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
		p.SetState(184)
		p.Match(mumathParserID)
	}
	{
		p.SetState(185)
		p.Match(mumathParserLPAREN)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumathParserNOT, mumathParserQUOTE, mumathParserMINUS, mumathParserLPAREN, mumathParserID, mumathParserSTRING, mumathParserNUMBER:
		{
			p.SetState(186)
			p.ActualParameter()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == mumathParserCOMMA {
			{
				p.SetState(187)
				p.Match(mumathParserCOMMA)
			}
			{
				p.SetState(188)
				p.ActualParameter()
			}

			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case mumathParserRPAREN:

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(197)
		p.Match(mumathParserRPAREN)
	}

	return localctx
}

// IEqualContext is an interface to support dynamic dispatch.
type IEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualContext differentiates from other interfaces.
	IsEqualContext()
}

type EqualContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualContext() *EqualContext {
	var p = new(EqualContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumathParserRULE_equal
	return p
}

func (*EqualContext) IsEqualContext() {}

func NewEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualContext {
	var p = new(EqualContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumathParserRULE_equal

	return p
}

func (s *EqualContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualContext) EQF() antlr.TerminalNode {
	return s.GetToken(mumathParserEQF, 0)
}

func (s *EqualContext) EQC() antlr.TerminalNode {
	return s.GetToken(mumathParserEQC, 0)
}

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumathListener); ok {
		listenerT.ExitEqual(s)
	}
}

func (p *mumathParser) Equal() (localctx IEqualContext) {
	localctx = NewEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, mumathParserRULE_equal)
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
		p.SetState(199)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mumathParserEQF || _la == mumathParserEQC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
