// Generated from robotwar.g4 by ANTLR 4.7.

package robotwars // robotwar
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 138,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 6, 2, 42, 10, 2, 13, 2, 14, 2, 43,
	3, 3, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 10, 5, 10, 82, 10, 10, 3, 10, 3, 10, 6, 10, 86, 10, 10,
	13, 10, 14, 10, 87, 3, 11, 3, 11, 3, 12, 5, 12, 93, 10, 12, 3, 12, 3, 12,
	5, 12, 97, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 7, 14, 109, 10, 14, 12, 14, 14, 14, 112, 11, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 118, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 5, 17, 127, 10, 17, 3, 18, 3, 18, 3, 19, 5, 19, 132,
	10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 7, 4, 2, 51, 52,
	56, 56, 3, 2, 12, 15, 3, 2, 3, 6, 3, 2, 16, 50, 3, 2, 12, 13, 2, 140, 2,
	41, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2,
	10, 62, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 74, 3, 2, 2, 2, 16, 77, 3,
	2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 89, 3, 2, 2, 2, 22, 92, 3, 2, 2, 2, 24,
	100, 3, 2, 2, 2, 26, 117, 3, 2, 2, 2, 28, 119, 3, 2, 2, 2, 30, 121, 3,
	2, 2, 2, 32, 126, 3, 2, 2, 2, 34, 128, 3, 2, 2, 2, 36, 131, 3, 2, 2, 2,
	38, 135, 3, 2, 2, 2, 40, 42, 5, 4, 3, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 3, 3, 2, 2, 2, 45,
	49, 5, 6, 4, 2, 46, 49, 5, 38, 20, 2, 47, 49, 5, 8, 5, 2, 48, 45, 3, 2,
	2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 50, 51, 7, 56, 2, 2, 51, 5, 3, 2, 2, 2, 52, 53, 7, 53, 2, 2,
	53, 7, 3, 2, 2, 2, 54, 61, 5, 22, 12, 2, 55, 61, 5, 18, 10, 2, 56, 61,
	5, 14, 8, 2, 57, 61, 5, 16, 9, 2, 58, 61, 5, 20, 11, 2, 59, 61, 5, 10,
	6, 2, 60, 54, 3, 2, 2, 2, 60, 55, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 57,
	3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2,
	62, 63, 5, 12, 7, 2, 63, 11, 3, 2, 2, 2, 64, 71, 7, 3, 2, 2, 65, 71, 7,
	4, 2, 2, 66, 67, 7, 5, 2, 2, 67, 71, 5, 26, 14, 2, 68, 69, 7, 6, 2, 2,
	69, 71, 5, 26, 14, 2, 70, 64, 3, 2, 2, 2, 70, 65, 3, 2, 2, 2, 70, 66, 3,
	2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 5, 8, 5, 2, 73,
	13, 3, 2, 2, 2, 74, 75, 7, 7, 2, 2, 75, 76, 5, 6, 4, 2, 76, 15, 3, 2, 2,
	2, 77, 78, 7, 8, 2, 2, 78, 79, 5, 6, 4, 2, 79, 17, 3, 2, 2, 2, 80, 82,
	5, 26, 14, 2, 81, 80, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 85, 3, 2, 2,
	2, 83, 84, 7, 9, 2, 2, 84, 86, 5, 34, 18, 2, 85, 83, 3, 2, 2, 2, 86, 87,
	3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 19, 3, 2, 2, 2,
	89, 90, 7, 10, 2, 2, 90, 21, 3, 2, 2, 2, 91, 93, 7, 11, 2, 2, 92, 91, 3,
	2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 5, 24, 13, 2,
	95, 97, 9, 2, 2, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3,
	2, 2, 2, 98, 99, 5, 8, 5, 2, 99, 23, 3, 2, 2, 2, 100, 101, 5, 26, 14, 2,
	101, 102, 5, 30, 16, 2, 102, 103, 5, 26, 14, 2, 103, 25, 3, 2, 2, 2, 104,
	110, 5, 32, 17, 2, 105, 106, 5, 28, 15, 2, 106, 107, 5, 32, 17, 2, 107,
	109, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108,
	3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 118, 3, 2, 2, 2, 112, 110, 3, 2,
	2, 2, 113, 114, 5, 28, 15, 2, 114, 115, 5, 32, 17, 2, 115, 118, 3, 2, 2,
	2, 116, 118, 5, 32, 17, 2, 117, 104, 3, 2, 2, 2, 117, 113, 3, 2, 2, 2,
	117, 116, 3, 2, 2, 2, 118, 27, 3, 2, 2, 2, 119, 120, 9, 3, 2, 2, 120, 29,
	3, 2, 2, 2, 121, 122, 9, 4, 2, 2, 122, 31, 3, 2, 2, 2, 123, 127, 5, 36,
	19, 2, 124, 127, 5, 34, 18, 2, 125, 127, 7, 50, 2, 2, 126, 123, 3, 2, 2,
	2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 33, 3, 2, 2, 2, 128,
	129, 9, 5, 2, 2, 129, 35, 3, 2, 2, 2, 130, 132, 9, 6, 2, 2, 131, 130, 3,
	2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134, 7, 54, 2,
	2, 134, 37, 3, 2, 2, 2, 135, 136, 7, 55, 2, 2, 136, 39, 3, 2, 2, 2, 14,
	43, 48, 60, 70, 81, 87, 92, 96, 110, 117, 126, 131,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'#'", "'<'", "'>'", "'GOSUB'", "'GOTO'", "'TO'", "'ENDSUB'",
	"'IF'", "'+'", "'-'", "'*'", "'/'", "'A'", "'B'", "'C'", "'D'", "'E'",
	"'F'", "'G'", "'H'", "'I'", "'J'", "'K'", "'L'", "'M'", "'N'", "'O'", "'P'",
	"'Q'", "'R'", "'S'", "'T'", "'U'", "'V'", "'W'", "'X'", "'Y'", "'Z'", "'AIM'",
	"'SHOT'", "'RADAR'", "'DAMAGE'", "'SPEEDX'", "'SPEEDY'", "'RANDOM'", "'INDEX'",
	"'DATA'", "'.'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "A", "B", "C",
	"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z", "AIM", "SHOT", "RADAR", "DAMAGE",
	"SPEEDX", "SPEEDY", "RANDOM", "INDEX", "DATA", "DOT", "COMMA", "ID", "NUMBER",
	"COMMENT", "EOL", "WS",
}

var ruleNames = []string{
	"program", "line", "label", "statement", "accumstatement", "accumexpression",
	"gosubstatement", "gotostatement", "tostatement", "endsubstatement", "ifstatement",
	"condition", "expression", "operation", "comparison", "argument", "register",
	"number", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type robotwarParser struct {
	*antlr.BaseParser
}

func NewrobotwarParser(input antlr.TokenStream) *robotwarParser {
	this := new(robotwarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "robotwar.g4"

	return this
}

// robotwarParser tokens.
const (
	robotwarParserEOF     = antlr.TokenEOF
	robotwarParserT__0    = 1
	robotwarParserT__1    = 2
	robotwarParserT__2    = 3
	robotwarParserT__3    = 4
	robotwarParserT__4    = 5
	robotwarParserT__5    = 6
	robotwarParserT__6    = 7
	robotwarParserT__7    = 8
	robotwarParserT__8    = 9
	robotwarParserT__9    = 10
	robotwarParserT__10   = 11
	robotwarParserT__11   = 12
	robotwarParserT__12   = 13
	robotwarParserA       = 14
	robotwarParserB       = 15
	robotwarParserC       = 16
	robotwarParserD       = 17
	robotwarParserE       = 18
	robotwarParserF       = 19
	robotwarParserG       = 20
	robotwarParserH       = 21
	robotwarParserI       = 22
	robotwarParserJ       = 23
	robotwarParserK       = 24
	robotwarParserL       = 25
	robotwarParserM       = 26
	robotwarParserN       = 27
	robotwarParserO       = 28
	robotwarParserP       = 29
	robotwarParserQ       = 30
	robotwarParserR       = 31
	robotwarParserS       = 32
	robotwarParserT       = 33
	robotwarParserU       = 34
	robotwarParserV       = 35
	robotwarParserW       = 36
	robotwarParserX       = 37
	robotwarParserY       = 38
	robotwarParserZ       = 39
	robotwarParserAIM     = 40
	robotwarParserSHOT    = 41
	robotwarParserRADAR   = 42
	robotwarParserDAMAGE  = 43
	robotwarParserSPEEDX  = 44
	robotwarParserSPEEDY  = 45
	robotwarParserRANDOM  = 46
	robotwarParserINDEX   = 47
	robotwarParserDATA    = 48
	robotwarParserDOT     = 49
	robotwarParserCOMMA   = 50
	robotwarParserID      = 51
	robotwarParserNUMBER  = 52
	robotwarParserCOMMENT = 53
	robotwarParserEOL     = 54
	robotwarParserWS      = 55
)

// robotwarParser rules.
const (
	robotwarParserRULE_program         = 0
	robotwarParserRULE_line            = 1
	robotwarParserRULE_label           = 2
	robotwarParserRULE_statement       = 3
	robotwarParserRULE_accumstatement  = 4
	robotwarParserRULE_accumexpression = 5
	robotwarParserRULE_gosubstatement  = 6
	robotwarParserRULE_gotostatement   = 7
	robotwarParserRULE_tostatement     = 8
	robotwarParserRULE_endsubstatement = 9
	robotwarParserRULE_ifstatement     = 10
	robotwarParserRULE_condition       = 11
	robotwarParserRULE_expression      = 12
	robotwarParserRULE_operation       = 13
	robotwarParserRULE_comparison      = 14
	robotwarParserRULE_argument        = 15
	robotwarParserRULE_register        = 16
	robotwarParserRULE_number          = 17
	robotwarParserRULE_comment         = 18
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
	p.RuleIndex = robotwarParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgramContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *robotwarParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, robotwarParserRULE_program)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<robotwarParserT__0)|(1<<robotwarParserT__1)|(1<<robotwarParserT__2)|(1<<robotwarParserT__3)|(1<<robotwarParserT__4)|(1<<robotwarParserT__5)|(1<<robotwarParserT__6)|(1<<robotwarParserT__7)|(1<<robotwarParserT__8)|(1<<robotwarParserT__9)|(1<<robotwarParserT__10)|(1<<robotwarParserT__11)|(1<<robotwarParserT__12)|(1<<robotwarParserA)|(1<<robotwarParserB)|(1<<robotwarParserC)|(1<<robotwarParserD)|(1<<robotwarParserE)|(1<<robotwarParserF)|(1<<robotwarParserG)|(1<<robotwarParserH)|(1<<robotwarParserI)|(1<<robotwarParserJ)|(1<<robotwarParserK)|(1<<robotwarParserL)|(1<<robotwarParserM)|(1<<robotwarParserN)|(1<<robotwarParserO)|(1<<robotwarParserP)|(1<<robotwarParserQ)|(1<<robotwarParserR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(robotwarParserS-32))|(1<<(robotwarParserT-32))|(1<<(robotwarParserU-32))|(1<<(robotwarParserV-32))|(1<<(robotwarParserW-32))|(1<<(robotwarParserX-32))|(1<<(robotwarParserY-32))|(1<<(robotwarParserZ-32))|(1<<(robotwarParserAIM-32))|(1<<(robotwarParserSHOT-32))|(1<<(robotwarParserRADAR-32))|(1<<(robotwarParserDAMAGE-32))|(1<<(robotwarParserSPEEDX-32))|(1<<(robotwarParserSPEEDY-32))|(1<<(robotwarParserRANDOM-32))|(1<<(robotwarParserINDEX-32))|(1<<(robotwarParserDATA-32))|(1<<(robotwarParserID-32))|(1<<(robotwarParserNUMBER-32))|(1<<(robotwarParserCOMMENT-32))|(1<<(robotwarParserEOL-32)))) != 0) {
		{
			p.SetState(38)
			p.Line()
		}

		p.SetState(41)
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
	p.RuleIndex = robotwarParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) EOL() antlr.TerminalNode {
	return s.GetToken(robotwarParserEOL, 0)
}

func (s *LineContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *robotwarParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, robotwarParserRULE_line)

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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case robotwarParserID:
		{
			p.SetState(43)
			p.Label()
		}

	case robotwarParserCOMMENT:
		{
			p.SetState(44)
			p.Comment()
		}

	case robotwarParserT__0, robotwarParserT__1, robotwarParserT__2, robotwarParserT__3, robotwarParserT__4, robotwarParserT__5, robotwarParserT__6, robotwarParserT__7, robotwarParserT__8, robotwarParserT__9, robotwarParserT__10, robotwarParserT__11, robotwarParserT__12, robotwarParserA, robotwarParserB, robotwarParserC, robotwarParserD, robotwarParserE, robotwarParserF, robotwarParserG, robotwarParserH, robotwarParserI, robotwarParserJ, robotwarParserK, robotwarParserL, robotwarParserM, robotwarParserN, robotwarParserO, robotwarParserP, robotwarParserQ, robotwarParserR, robotwarParserS, robotwarParserT, robotwarParserU, robotwarParserV, robotwarParserW, robotwarParserX, robotwarParserY, robotwarParserZ, robotwarParserAIM, robotwarParserSHOT, robotwarParserRADAR, robotwarParserDAMAGE, robotwarParserSPEEDX, robotwarParserSPEEDY, robotwarParserRANDOM, robotwarParserINDEX, robotwarParserDATA, robotwarParserNUMBER:
		{
			p.SetState(45)
			p.Statement()
		}

	case robotwarParserEOL:

	default:
	}
	{
		p.SetState(48)
		p.Match(robotwarParserEOL)
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
	p.RuleIndex = robotwarParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) ID() antlr.TerminalNode {
	return s.GetToken(robotwarParserID, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *robotwarParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, robotwarParserRULE_label)

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
		p.SetState(50)
		p.Match(robotwarParserID)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Ifstatement() IIfstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatementContext)
}

func (s *StatementContext) Tostatement() ITostatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITostatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITostatementContext)
}

func (s *StatementContext) Gosubstatement() IGosubstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGosubstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGosubstatementContext)
}

func (s *StatementContext) Gotostatement() IGotostatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotostatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotostatementContext)
}

func (s *StatementContext) Endsubstatement() IEndsubstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndsubstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndsubstatementContext)
}

func (s *StatementContext) Accumstatement() IAccumstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccumstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccumstatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *robotwarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, robotwarParserRULE_statement)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Ifstatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Tostatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Gosubstatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Gotostatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(56)
			p.Endsubstatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(57)
			p.Accumstatement()
		}

	}

	return localctx
}

// IAccumstatementContext is an interface to support dynamic dispatch.
type IAccumstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccumstatementContext differentiates from other interfaces.
	IsAccumstatementContext()
}

type AccumstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccumstatementContext() *AccumstatementContext {
	var p = new(AccumstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_accumstatement
	return p
}

func (*AccumstatementContext) IsAccumstatementContext() {}

func NewAccumstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccumstatementContext {
	var p = new(AccumstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_accumstatement

	return p
}

func (s *AccumstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AccumstatementContext) Accumexpression() IAccumexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccumexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccumexpressionContext)
}

func (s *AccumstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccumstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccumstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterAccumstatement(s)
	}
}

func (s *AccumstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitAccumstatement(s)
	}
}

func (p *robotwarParser) Accumstatement() (localctx IAccumstatementContext) {
	localctx = NewAccumstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, robotwarParserRULE_accumstatement)

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
		p.Accumexpression()
	}

	return localctx
}

// IAccumexpressionContext is an interface to support dynamic dispatch.
type IAccumexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccumexpressionContext differentiates from other interfaces.
	IsAccumexpressionContext()
}

type AccumexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccumexpressionContext() *AccumexpressionContext {
	var p = new(AccumexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_accumexpression
	return p
}

func (*AccumexpressionContext) IsAccumexpressionContext() {}

func NewAccumexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccumexpressionContext {
	var p = new(AccumexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_accumexpression

	return p
}

func (s *AccumexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AccumexpressionContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *AccumexpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AccumexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccumexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccumexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterAccumexpression(s)
	}
}

func (s *AccumexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitAccumexpression(s)
	}
}

func (p *robotwarParser) Accumexpression() (localctx IAccumexpressionContext) {
	localctx = NewAccumexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, robotwarParserRULE_accumexpression)

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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case robotwarParserT__0:
		{
			p.SetState(62)
			p.Match(robotwarParserT__0)
		}

	case robotwarParserT__1:
		{
			p.SetState(63)
			p.Match(robotwarParserT__1)
		}

	case robotwarParserT__2:
		{
			p.SetState(64)
			p.Match(robotwarParserT__2)
		}
		{
			p.SetState(65)
			p.Expression()
		}

	case robotwarParserT__3:
		{
			p.SetState(66)
			p.Match(robotwarParserT__3)
		}
		{
			p.SetState(67)
			p.Expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(70)
		p.Statement()
	}

	return localctx
}

// IGosubstatementContext is an interface to support dynamic dispatch.
type IGosubstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGosubstatementContext differentiates from other interfaces.
	IsGosubstatementContext()
}

type GosubstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGosubstatementContext() *GosubstatementContext {
	var p = new(GosubstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_gosubstatement
	return p
}

func (*GosubstatementContext) IsGosubstatementContext() {}

func NewGosubstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GosubstatementContext {
	var p = new(GosubstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_gosubstatement

	return p
}

func (s *GosubstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GosubstatementContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *GosubstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GosubstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GosubstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterGosubstatement(s)
	}
}

func (s *GosubstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitGosubstatement(s)
	}
}

func (p *robotwarParser) Gosubstatement() (localctx IGosubstatementContext) {
	localctx = NewGosubstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, robotwarParserRULE_gosubstatement)

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
		p.Match(robotwarParserT__4)
	}
	{
		p.SetState(73)
		p.Label()
	}

	return localctx
}

// IGotostatementContext is an interface to support dynamic dispatch.
type IGotostatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotostatementContext differentiates from other interfaces.
	IsGotostatementContext()
}

type GotostatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotostatementContext() *GotostatementContext {
	var p = new(GotostatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_gotostatement
	return p
}

func (*GotostatementContext) IsGotostatementContext() {}

func NewGotostatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotostatementContext {
	var p = new(GotostatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_gotostatement

	return p
}

func (s *GotostatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GotostatementContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *GotostatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotostatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotostatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterGotostatement(s)
	}
}

func (s *GotostatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitGotostatement(s)
	}
}

func (p *robotwarParser) Gotostatement() (localctx IGotostatementContext) {
	localctx = NewGotostatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, robotwarParserRULE_gotostatement)

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
		p.SetState(75)
		p.Match(robotwarParserT__5)
	}
	{
		p.SetState(76)
		p.Label()
	}

	return localctx
}

// ITostatementContext is an interface to support dynamic dispatch.
type ITostatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTostatementContext differentiates from other interfaces.
	IsTostatementContext()
}

type TostatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTostatementContext() *TostatementContext {
	var p = new(TostatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_tostatement
	return p
}

func (*TostatementContext) IsTostatementContext() {}

func NewTostatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TostatementContext {
	var p = new(TostatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_tostatement

	return p
}

func (s *TostatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TostatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TostatementContext) AllRegister() []IRegisterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegisterContext)(nil)).Elem())
	var tst = make([]IRegisterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegisterContext)
		}
	}

	return tst
}

func (s *TostatementContext) Register(i int) IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *TostatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TostatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TostatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterTostatement(s)
	}
}

func (s *TostatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitTostatement(s)
	}
}

func (p *robotwarParser) Tostatement() (localctx ITostatementContext) {
	localctx = NewTostatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, robotwarParserRULE_tostatement)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<robotwarParserT__9)|(1<<robotwarParserT__10)|(1<<robotwarParserT__11)|(1<<robotwarParserT__12)|(1<<robotwarParserA)|(1<<robotwarParserB)|(1<<robotwarParserC)|(1<<robotwarParserD)|(1<<robotwarParserE)|(1<<robotwarParserF)|(1<<robotwarParserG)|(1<<robotwarParserH)|(1<<robotwarParserI)|(1<<robotwarParserJ)|(1<<robotwarParserK)|(1<<robotwarParserL)|(1<<robotwarParserM)|(1<<robotwarParserN)|(1<<robotwarParserO)|(1<<robotwarParserP)|(1<<robotwarParserQ)|(1<<robotwarParserR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(robotwarParserS-32))|(1<<(robotwarParserT-32))|(1<<(robotwarParserU-32))|(1<<(robotwarParserV-32))|(1<<(robotwarParserW-32))|(1<<(robotwarParserX-32))|(1<<(robotwarParserY-32))|(1<<(robotwarParserZ-32))|(1<<(robotwarParserAIM-32))|(1<<(robotwarParserSHOT-32))|(1<<(robotwarParserRADAR-32))|(1<<(robotwarParserDAMAGE-32))|(1<<(robotwarParserSPEEDX-32))|(1<<(robotwarParserSPEEDY-32))|(1<<(robotwarParserRANDOM-32))|(1<<(robotwarParserINDEX-32))|(1<<(robotwarParserDATA-32))|(1<<(robotwarParserNUMBER-32)))) != 0) {
		{
			p.SetState(78)
			p.Expression()
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == robotwarParserT__6 {
		{
			p.SetState(81)
			p.Match(robotwarParserT__6)
		}
		{
			p.SetState(82)
			p.Register()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEndsubstatementContext is an interface to support dynamic dispatch.
type IEndsubstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndsubstatementContext differentiates from other interfaces.
	IsEndsubstatementContext()
}

type EndsubstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndsubstatementContext() *EndsubstatementContext {
	var p = new(EndsubstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_endsubstatement
	return p
}

func (*EndsubstatementContext) IsEndsubstatementContext() {}

func NewEndsubstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndsubstatementContext {
	var p = new(EndsubstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_endsubstatement

	return p
}

func (s *EndsubstatementContext) GetParser() antlr.Parser { return s.parser }
func (s *EndsubstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndsubstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndsubstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterEndsubstatement(s)
	}
}

func (s *EndsubstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitEndsubstatement(s)
	}
}

func (p *robotwarParser) Endsubstatement() (localctx IEndsubstatementContext) {
	localctx = NewEndsubstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, robotwarParserRULE_endsubstatement)

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
		p.SetState(87)
		p.Match(robotwarParserT__7)
	}

	return localctx
}

// IIfstatementContext is an interface to support dynamic dispatch.
type IIfstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatementContext differentiates from other interfaces.
	IsIfstatementContext()
}

type IfstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatementContext() *IfstatementContext {
	var p = new(IfstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_ifstatement
	return p
}

func (*IfstatementContext) IsIfstatementContext() {}

func NewIfstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatementContext {
	var p = new(IfstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_ifstatement

	return p
}

func (s *IfstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfstatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstatementContext) EOL() antlr.TerminalNode {
	return s.GetToken(robotwarParserEOL, 0)
}

func (s *IfstatementContext) COMMA() antlr.TerminalNode {
	return s.GetToken(robotwarParserCOMMA, 0)
}

func (s *IfstatementContext) DOT() antlr.TerminalNode {
	return s.GetToken(robotwarParserDOT, 0)
}

func (s *IfstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterIfstatement(s)
	}
}

func (s *IfstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitIfstatement(s)
	}
}

func (p *robotwarParser) Ifstatement() (localctx IIfstatementContext) {
	localctx = NewIfstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, robotwarParserRULE_ifstatement)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == robotwarParserT__8 {
		{
			p.SetState(89)
			p.Match(robotwarParserT__8)
		}

	}
	{
		p.SetState(92)
		p.Condition()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(robotwarParserDOT-49))|(1<<(robotwarParserCOMMA-49))|(1<<(robotwarParserEOL-49)))) != 0 {
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(robotwarParserDOT-49))|(1<<(robotwarParserCOMMA-49))|(1<<(robotwarParserEOL-49)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(96)
		p.Statement()
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *robotwarParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, robotwarParserRULE_condition)

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
		p.Expression()
	}
	{
		p.SetState(99)
		p.Comparison()
	}
	{
		p.SetState(100)
		p.Expression()
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
	p.RuleIndex = robotwarParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ExpressionContext) AllOperation() []IOperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationContext)(nil)).Elem())
	var tst = make([]IOperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Operation(i int) IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *robotwarParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, robotwarParserRULE_expression)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Argument()
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(103)
					p.Operation()
				}
				{
					p.SetState(104)
					p.Argument()
				}

			}
			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.Operation()
		}
		{
			p.SetState(112)
			p.Argument()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(114)
			p.Argument()
		}

	}

	return localctx
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *robotwarParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, robotwarParserRULE_operation)
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
	p.SetState(117)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<robotwarParserT__9)|(1<<robotwarParserT__10)|(1<<robotwarParserT__11)|(1<<robotwarParserT__12))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (p *robotwarParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, robotwarParserRULE_comparison)
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
	p.SetState(119)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<robotwarParserT__0)|(1<<robotwarParserT__1)|(1<<robotwarParserT__2)|(1<<robotwarParserT__3))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
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
	p.RuleIndex = robotwarParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgumentContext) Register() IRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegisterContext)
}

func (s *ArgumentContext) DATA() antlr.TerminalNode {
	return s.GetToken(robotwarParserDATA, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *robotwarParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, robotwarParserRULE_argument)

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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Register()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Match(robotwarParserDATA)
		}

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
	p.RuleIndex = robotwarParserRULE_register
	return p
}

func (*RegisterContext) IsRegisterContext() {}

func NewRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegisterContext {
	var p = new(RegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_register

	return p
}

func (s *RegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *RegisterContext) A() antlr.TerminalNode {
	return s.GetToken(robotwarParserA, 0)
}

func (s *RegisterContext) B() antlr.TerminalNode {
	return s.GetToken(robotwarParserB, 0)
}

func (s *RegisterContext) C() antlr.TerminalNode {
	return s.GetToken(robotwarParserC, 0)
}

func (s *RegisterContext) D() antlr.TerminalNode {
	return s.GetToken(robotwarParserD, 0)
}

func (s *RegisterContext) E() antlr.TerminalNode {
	return s.GetToken(robotwarParserE, 0)
}

func (s *RegisterContext) F() antlr.TerminalNode {
	return s.GetToken(robotwarParserF, 0)
}

func (s *RegisterContext) G() antlr.TerminalNode {
	return s.GetToken(robotwarParserG, 0)
}

func (s *RegisterContext) H() antlr.TerminalNode {
	return s.GetToken(robotwarParserH, 0)
}

func (s *RegisterContext) I() antlr.TerminalNode {
	return s.GetToken(robotwarParserI, 0)
}

func (s *RegisterContext) J() antlr.TerminalNode {
	return s.GetToken(robotwarParserJ, 0)
}

func (s *RegisterContext) K() antlr.TerminalNode {
	return s.GetToken(robotwarParserK, 0)
}

func (s *RegisterContext) L() antlr.TerminalNode {
	return s.GetToken(robotwarParserL, 0)
}

func (s *RegisterContext) M() antlr.TerminalNode {
	return s.GetToken(robotwarParserM, 0)
}

func (s *RegisterContext) N() antlr.TerminalNode {
	return s.GetToken(robotwarParserN, 0)
}

func (s *RegisterContext) O() antlr.TerminalNode {
	return s.GetToken(robotwarParserO, 0)
}

func (s *RegisterContext) P() antlr.TerminalNode {
	return s.GetToken(robotwarParserP, 0)
}

func (s *RegisterContext) Q() antlr.TerminalNode {
	return s.GetToken(robotwarParserQ, 0)
}

func (s *RegisterContext) R() antlr.TerminalNode {
	return s.GetToken(robotwarParserR, 0)
}

func (s *RegisterContext) S() antlr.TerminalNode {
	return s.GetToken(robotwarParserS, 0)
}

func (s *RegisterContext) T() antlr.TerminalNode {
	return s.GetToken(robotwarParserT, 0)
}

func (s *RegisterContext) U() antlr.TerminalNode {
	return s.GetToken(robotwarParserU, 0)
}

func (s *RegisterContext) V() antlr.TerminalNode {
	return s.GetToken(robotwarParserV, 0)
}

func (s *RegisterContext) W() antlr.TerminalNode {
	return s.GetToken(robotwarParserW, 0)
}

func (s *RegisterContext) X() antlr.TerminalNode {
	return s.GetToken(robotwarParserX, 0)
}

func (s *RegisterContext) Y() antlr.TerminalNode {
	return s.GetToken(robotwarParserY, 0)
}

func (s *RegisterContext) Z() antlr.TerminalNode {
	return s.GetToken(robotwarParserZ, 0)
}

func (s *RegisterContext) AIM() antlr.TerminalNode {
	return s.GetToken(robotwarParserAIM, 0)
}

func (s *RegisterContext) SHOT() antlr.TerminalNode {
	return s.GetToken(robotwarParserSHOT, 0)
}

func (s *RegisterContext) RADAR() antlr.TerminalNode {
	return s.GetToken(robotwarParserRADAR, 0)
}

func (s *RegisterContext) SPEEDX() antlr.TerminalNode {
	return s.GetToken(robotwarParserSPEEDX, 0)
}

func (s *RegisterContext) SPEEDY() antlr.TerminalNode {
	return s.GetToken(robotwarParserSPEEDY, 0)
}

func (s *RegisterContext) RANDOM() antlr.TerminalNode {
	return s.GetToken(robotwarParserRANDOM, 0)
}

func (s *RegisterContext) INDEX() antlr.TerminalNode {
	return s.GetToken(robotwarParserINDEX, 0)
}

func (s *RegisterContext) DATA() antlr.TerminalNode {
	return s.GetToken(robotwarParserDATA, 0)
}

func (s *RegisterContext) DAMAGE() antlr.TerminalNode {
	return s.GetToken(robotwarParserDAMAGE, 0)
}

func (s *RegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterRegister(s)
	}
}

func (s *RegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitRegister(s)
	}
}

func (p *robotwarParser) Register() (localctx IRegisterContext) {
	localctx = NewRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, robotwarParserRULE_register)
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
	p.SetState(126)
	_la = p.GetTokenStream().LA(1)

	if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<robotwarParserA)|(1<<robotwarParserB)|(1<<robotwarParserC)|(1<<robotwarParserD)|(1<<robotwarParserE)|(1<<robotwarParserF)|(1<<robotwarParserG)|(1<<robotwarParserH)|(1<<robotwarParserI)|(1<<robotwarParserJ)|(1<<robotwarParserK)|(1<<robotwarParserL)|(1<<robotwarParserM)|(1<<robotwarParserN)|(1<<robotwarParserO)|(1<<robotwarParserP)|(1<<robotwarParserQ)|(1<<robotwarParserR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(robotwarParserS-32))|(1<<(robotwarParserT-32))|(1<<(robotwarParserU-32))|(1<<(robotwarParserV-32))|(1<<(robotwarParserW-32))|(1<<(robotwarParserX-32))|(1<<(robotwarParserY-32))|(1<<(robotwarParserZ-32))|(1<<(robotwarParserAIM-32))|(1<<(robotwarParserSHOT-32))|(1<<(robotwarParserRADAR-32))|(1<<(robotwarParserDAMAGE-32))|(1<<(robotwarParserSPEEDX-32))|(1<<(robotwarParserSPEEDY-32))|(1<<(robotwarParserRANDOM-32))|(1<<(robotwarParserINDEX-32))|(1<<(robotwarParserDATA-32)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = robotwarParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(robotwarParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *robotwarParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, robotwarParserRULE_number)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == robotwarParserT__9 || _la == robotwarParserT__10 {
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !(_la == robotwarParserT__9 || _la == robotwarParserT__10) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(131)
		p.Match(robotwarParserNUMBER)
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
	p.RuleIndex = robotwarParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = robotwarParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(robotwarParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(robotwarListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *robotwarParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, robotwarParserRULE_comment)

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
		p.SetState(133)
		p.Match(robotwarParserCOMMENT)
	}

	return localctx
}
