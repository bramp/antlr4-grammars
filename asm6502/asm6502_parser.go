// Code generated from asm6502.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm6502 // asm6502
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 110,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 5, 2, 34, 10, 2,
	3, 2, 6, 2, 37, 10, 2, 13, 2, 14, 2, 38, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	45, 10, 3, 3, 4, 5, 4, 48, 10, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 4, 5,
	4, 55, 10, 4, 3, 5, 5, 5, 58, 10, 5, 3, 5, 3, 5, 5, 5, 62, 10, 5, 3, 5,
	5, 5, 65, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8,
	75, 10, 8, 3, 9, 3, 9, 3, 10, 5, 10, 80, 10, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 86, 10, 10, 3, 10, 3, 10, 5, 10, 90, 10, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 96, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 4, 3, 2, 6, 7, 3, 2, 12,
	72, 2, 112, 2, 36, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2, 6, 47, 3, 2, 2, 2, 8,
	57, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 71, 3, 2, 2,
	2, 16, 76, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 97, 3, 2, 2, 2, 22, 99,
	3, 2, 2, 2, 24, 101, 3, 2, 2, 2, 26, 103, 3, 2, 2, 2, 28, 105, 3, 2, 2,
	2, 30, 107, 3, 2, 2, 2, 32, 34, 5, 4, 3, 2, 33, 32, 3, 2, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 7, 77, 2, 2, 36, 33, 3, 2, 2, 2,
	37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 3, 3, 2,
	2, 2, 40, 45, 5, 28, 15, 2, 41, 45, 5, 6, 4, 2, 42, 45, 5, 8, 5, 2, 43,
	45, 5, 12, 7, 2, 44, 40, 3, 2, 2, 2, 44, 41, 3, 2, 2, 2, 44, 42, 3, 2,
	2, 2, 44, 43, 3, 2, 2, 2, 45, 5, 3, 2, 2, 2, 46, 48, 5, 16, 9, 2, 47, 46,
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 5, 30, 16,
	2, 50, 52, 5, 14, 8, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54,
	3, 2, 2, 2, 53, 55, 5, 28, 15, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2,
	2, 55, 7, 3, 2, 2, 2, 56, 58, 5, 18, 10, 2, 57, 56, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 5, 10, 6, 2, 60, 62, 5, 14, 8,
	2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 65,
	5, 28, 15, 2, 64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 9, 3, 2, 2, 2,
	66, 67, 7, 11, 2, 2, 67, 11, 3, 2, 2, 2, 68, 69, 5, 16, 9, 2, 69, 70, 7,
	3, 2, 2, 70, 13, 3, 2, 2, 2, 71, 74, 5, 18, 10, 2, 72, 73, 7, 4, 2, 2,
	73, 75, 5, 14, 8, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 15, 3,
	2, 2, 2, 76, 77, 5, 24, 13, 2, 77, 17, 3, 2, 2, 2, 78, 80, 5, 20, 11, 2,
	79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 85, 3, 2, 2, 2, 81, 86, 5,
	26, 14, 2, 82, 86, 5, 24, 13, 2, 83, 86, 5, 22, 12, 2, 84, 86, 7, 5, 2,
	2, 85, 81, 3, 2, 2, 2, 85, 82, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 84,
	3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 88, 9, 2, 2, 2, 88, 90, 5, 26, 14,
	2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 96, 3, 2, 2, 2, 91, 92,
	7, 8, 2, 2, 92, 93, 5, 18, 10, 2, 93, 94, 7, 9, 2, 2, 94, 96, 3, 2, 2,
	2, 95, 79, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 96, 19, 3, 2, 2, 2, 97, 98,
	7, 10, 2, 2, 98, 21, 3, 2, 2, 2, 99, 100, 7, 76, 2, 2, 100, 23, 3, 2, 2,
	2, 101, 102, 7, 73, 2, 2, 102, 25, 3, 2, 2, 2, 103, 104, 7, 74, 2, 2, 104,
	27, 3, 2, 2, 2, 105, 106, 7, 75, 2, 2, 106, 29, 3, 2, 2, 2, 107, 108, 9,
	3, 2, 2, 108, 31, 3, 2, 2, 2, 16, 33, 38, 44, 47, 51, 54, 57, 61, 64, 74,
	79, 85, 89, 95,
}
var literalNames = []string{
	"", "':'", "','", "'*'", "'+'", "'-'", "'('", "')'", "'#'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ASSEMBLER_INSTRUCTION", "ADC", "AND",
	"ASL", "BCC", "BCS", "BEQ", "BIT", "BMI", "BNE", "BPL", "BRA", "BRK", "BVC",
	"BVS", "CLC", "CLD", "CLI", "CLV", "CMP", "CPX", "CPY", "DEC", "DEX", "DEY",
	"EOR", "INC", "INX", "INY", "JMP", "JSR", "LDA", "LDY", "LDX", "LSR", "NOP",
	"ORA", "PHA", "PHX", "PHY", "PHP", "PLA", "PLP", "PLY", "ROL", "ROR", "RTI",
	"RTS", "SBC", "SEC", "SED", "SEI", "STA", "STX", "STY", "STZ", "TAX", "TAY",
	"TSX", "TXA", "TXS", "TYA", "NAME", "NUMBER", "COMMENT", "STRING", "EOL",
	"WS",
}

var ruleNames = []string{
	"prog", "line", "instruction", "assemblerinstruction", "assembleropcode",
	"lbl", "argumentlist", "label", "argument", "prefix_", "string_", "name",
	"number", "comment", "opcode",
}

type asm6502Parser struct {
	*antlr.BaseParser
}

// Newasm6502Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *asm6502Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newasm6502Parser(input antlr.TokenStream) *asm6502Parser {
	this := new(asm6502Parser)
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
	this.GrammarFileName = "asm6502.g4"

	return this
}

// asm6502Parser tokens.
const (
	asm6502ParserEOF                   = antlr.TokenEOF
	asm6502ParserT__0                  = 1
	asm6502ParserT__1                  = 2
	asm6502ParserT__2                  = 3
	asm6502ParserT__3                  = 4
	asm6502ParserT__4                  = 5
	asm6502ParserT__5                  = 6
	asm6502ParserT__6                  = 7
	asm6502ParserT__7                  = 8
	asm6502ParserASSEMBLER_INSTRUCTION = 9
	asm6502ParserADC                   = 10
	asm6502ParserAND                   = 11
	asm6502ParserASL                   = 12
	asm6502ParserBCC                   = 13
	asm6502ParserBCS                   = 14
	asm6502ParserBEQ                   = 15
	asm6502ParserBIT                   = 16
	asm6502ParserBMI                   = 17
	asm6502ParserBNE                   = 18
	asm6502ParserBPL                   = 19
	asm6502ParserBRA                   = 20
	asm6502ParserBRK                   = 21
	asm6502ParserBVC                   = 22
	asm6502ParserBVS                   = 23
	asm6502ParserCLC                   = 24
	asm6502ParserCLD                   = 25
	asm6502ParserCLI                   = 26
	asm6502ParserCLV                   = 27
	asm6502ParserCMP                   = 28
	asm6502ParserCPX                   = 29
	asm6502ParserCPY                   = 30
	asm6502ParserDEC                   = 31
	asm6502ParserDEX                   = 32
	asm6502ParserDEY                   = 33
	asm6502ParserEOR                   = 34
	asm6502ParserINC                   = 35
	asm6502ParserINX                   = 36
	asm6502ParserINY                   = 37
	asm6502ParserJMP                   = 38
	asm6502ParserJSR                   = 39
	asm6502ParserLDA                   = 40
	asm6502ParserLDY                   = 41
	asm6502ParserLDX                   = 42
	asm6502ParserLSR                   = 43
	asm6502ParserNOP                   = 44
	asm6502ParserORA                   = 45
	asm6502ParserPHA                   = 46
	asm6502ParserPHX                   = 47
	asm6502ParserPHY                   = 48
	asm6502ParserPHP                   = 49
	asm6502ParserPLA                   = 50
	asm6502ParserPLP                   = 51
	asm6502ParserPLY                   = 52
	asm6502ParserROL                   = 53
	asm6502ParserROR                   = 54
	asm6502ParserRTI                   = 55
	asm6502ParserRTS                   = 56
	asm6502ParserSBC                   = 57
	asm6502ParserSEC                   = 58
	asm6502ParserSED                   = 59
	asm6502ParserSEI                   = 60
	asm6502ParserSTA                   = 61
	asm6502ParserSTX                   = 62
	asm6502ParserSTY                   = 63
	asm6502ParserSTZ                   = 64
	asm6502ParserTAX                   = 65
	asm6502ParserTAY                   = 66
	asm6502ParserTSX                   = 67
	asm6502ParserTXA                   = 68
	asm6502ParserTXS                   = 69
	asm6502ParserTYA                   = 70
	asm6502ParserNAME                  = 71
	asm6502ParserNUMBER                = 72
	asm6502ParserCOMMENT               = 73
	asm6502ParserSTRING                = 74
	asm6502ParserEOL                   = 75
	asm6502ParserWS                    = 76
)

// asm6502Parser rules.
const (
	asm6502ParserRULE_prog                 = 0
	asm6502ParserRULE_line                 = 1
	asm6502ParserRULE_instruction          = 2
	asm6502ParserRULE_assemblerinstruction = 3
	asm6502ParserRULE_assembleropcode      = 4
	asm6502ParserRULE_lbl                  = 5
	asm6502ParserRULE_argumentlist         = 6
	asm6502ParserRULE_label                = 7
	asm6502ParserRULE_argument             = 8
	asm6502ParserRULE_prefix_              = 9
	asm6502ParserRULE_string_              = 10
	asm6502ParserRULE_name                 = 11
	asm6502ParserRULE_number               = 12
	asm6502ParserRULE_comment              = 13
	asm6502ParserRULE_opcode               = 14
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
	p.RuleIndex = asm6502ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(asm6502ParserEOL)
}

func (s *ProgContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(asm6502ParserEOL, i)
}

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
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *asm6502Parser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, asm6502ParserRULE_prog)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm6502ParserT__2)|(1<<asm6502ParserT__5)|(1<<asm6502ParserT__7)|(1<<asm6502ParserASSEMBLER_INSTRUCTION)|(1<<asm6502ParserADC)|(1<<asm6502ParserAND)|(1<<asm6502ParserASL)|(1<<asm6502ParserBCC)|(1<<asm6502ParserBCS)|(1<<asm6502ParserBEQ)|(1<<asm6502ParserBIT)|(1<<asm6502ParserBMI)|(1<<asm6502ParserBNE)|(1<<asm6502ParserBPL)|(1<<asm6502ParserBRA)|(1<<asm6502ParserBRK)|(1<<asm6502ParserBVC)|(1<<asm6502ParserBVS)|(1<<asm6502ParserCLC)|(1<<asm6502ParserCLD)|(1<<asm6502ParserCLI)|(1<<asm6502ParserCLV)|(1<<asm6502ParserCMP)|(1<<asm6502ParserCPX)|(1<<asm6502ParserCPY)|(1<<asm6502ParserDEC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(asm6502ParserDEX-32))|(1<<(asm6502ParserDEY-32))|(1<<(asm6502ParserEOR-32))|(1<<(asm6502ParserINC-32))|(1<<(asm6502ParserINX-32))|(1<<(asm6502ParserINY-32))|(1<<(asm6502ParserJMP-32))|(1<<(asm6502ParserJSR-32))|(1<<(asm6502ParserLDA-32))|(1<<(asm6502ParserLDY-32))|(1<<(asm6502ParserLDX-32))|(1<<(asm6502ParserLSR-32))|(1<<(asm6502ParserNOP-32))|(1<<(asm6502ParserORA-32))|(1<<(asm6502ParserPHA-32))|(1<<(asm6502ParserPHX-32))|(1<<(asm6502ParserPHY-32))|(1<<(asm6502ParserPHP-32))|(1<<(asm6502ParserPLA-32))|(1<<(asm6502ParserPLP-32))|(1<<(asm6502ParserPLY-32))|(1<<(asm6502ParserROL-32))|(1<<(asm6502ParserROR-32))|(1<<(asm6502ParserRTI-32))|(1<<(asm6502ParserRTS-32))|(1<<(asm6502ParserSBC-32))|(1<<(asm6502ParserSEC-32))|(1<<(asm6502ParserSED-32))|(1<<(asm6502ParserSEI-32))|(1<<(asm6502ParserSTA-32))|(1<<(asm6502ParserSTX-32))|(1<<(asm6502ParserSTY-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(asm6502ParserSTZ-64))|(1<<(asm6502ParserTAX-64))|(1<<(asm6502ParserTAY-64))|(1<<(asm6502ParserTSX-64))|(1<<(asm6502ParserTXA-64))|(1<<(asm6502ParserTXS-64))|(1<<(asm6502ParserTYA-64))|(1<<(asm6502ParserNAME-64))|(1<<(asm6502ParserNUMBER-64))|(1<<(asm6502ParserCOMMENT-64))|(1<<(asm6502ParserSTRING-64))|(1<<(asm6502ParserEOL-64)))) != 0) {
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm6502ParserT__2)|(1<<asm6502ParserT__5)|(1<<asm6502ParserT__7)|(1<<asm6502ParserASSEMBLER_INSTRUCTION)|(1<<asm6502ParserADC)|(1<<asm6502ParserAND)|(1<<asm6502ParserASL)|(1<<asm6502ParserBCC)|(1<<asm6502ParserBCS)|(1<<asm6502ParserBEQ)|(1<<asm6502ParserBIT)|(1<<asm6502ParserBMI)|(1<<asm6502ParserBNE)|(1<<asm6502ParserBPL)|(1<<asm6502ParserBRA)|(1<<asm6502ParserBRK)|(1<<asm6502ParserBVC)|(1<<asm6502ParserBVS)|(1<<asm6502ParserCLC)|(1<<asm6502ParserCLD)|(1<<asm6502ParserCLI)|(1<<asm6502ParserCLV)|(1<<asm6502ParserCMP)|(1<<asm6502ParserCPX)|(1<<asm6502ParserCPY)|(1<<asm6502ParserDEC))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(asm6502ParserDEX-32))|(1<<(asm6502ParserDEY-32))|(1<<(asm6502ParserEOR-32))|(1<<(asm6502ParserINC-32))|(1<<(asm6502ParserINX-32))|(1<<(asm6502ParserINY-32))|(1<<(asm6502ParserJMP-32))|(1<<(asm6502ParserJSR-32))|(1<<(asm6502ParserLDA-32))|(1<<(asm6502ParserLDY-32))|(1<<(asm6502ParserLDX-32))|(1<<(asm6502ParserLSR-32))|(1<<(asm6502ParserNOP-32))|(1<<(asm6502ParserORA-32))|(1<<(asm6502ParserPHA-32))|(1<<(asm6502ParserPHX-32))|(1<<(asm6502ParserPHY-32))|(1<<(asm6502ParserPHP-32))|(1<<(asm6502ParserPLA-32))|(1<<(asm6502ParserPLP-32))|(1<<(asm6502ParserPLY-32))|(1<<(asm6502ParserROL-32))|(1<<(asm6502ParserROR-32))|(1<<(asm6502ParserRTI-32))|(1<<(asm6502ParserRTS-32))|(1<<(asm6502ParserSBC-32))|(1<<(asm6502ParserSEC-32))|(1<<(asm6502ParserSED-32))|(1<<(asm6502ParserSEI-32))|(1<<(asm6502ParserSTA-32))|(1<<(asm6502ParserSTX-32))|(1<<(asm6502ParserSTY-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(asm6502ParserSTZ-64))|(1<<(asm6502ParserTAX-64))|(1<<(asm6502ParserTAY-64))|(1<<(asm6502ParserTSX-64))|(1<<(asm6502ParserTXA-64))|(1<<(asm6502ParserTXS-64))|(1<<(asm6502ParserTYA-64))|(1<<(asm6502ParserNAME-64))|(1<<(asm6502ParserNUMBER-64))|(1<<(asm6502ParserCOMMENT-64))|(1<<(asm6502ParserSTRING-64)))) != 0) {
			{
				p.SetState(30)
				p.Line()
			}

		}
		{
			p.SetState(33)
			p.Match(asm6502ParserEOL)
		}

		p.SetState(36)
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
	p.RuleIndex = asm6502ParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *LineContext) Assemblerinstruction() IAssemblerinstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerinstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblerinstructionContext)
}

func (s *LineContext) Lbl() ILblContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *asm6502Parser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, asm6502ParserRULE_line)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Comment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Instruction()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Assemblerinstruction()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Lbl()
		}

	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *InstructionContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *InstructionContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *InstructionContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *asm6502Parser) Instruction() (localctx IInstructionContext) {
	this := p
	_ = this

	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, asm6502ParserRULE_instruction)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm6502ParserNAME {
		{
			p.SetState(44)
			p.Label()
		}

	}
	{
		p.SetState(47)
		p.Opcode()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm6502ParserT__2)|(1<<asm6502ParserT__5)|(1<<asm6502ParserT__7))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(asm6502ParserNAME-71))|(1<<(asm6502ParserNUMBER-71))|(1<<(asm6502ParserSTRING-71)))) != 0) {
		{
			p.SetState(48)
			p.Argumentlist()
		}

	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm6502ParserCOMMENT {
		{
			p.SetState(51)
			p.Comment()
		}

	}

	return localctx
}

// IAssemblerinstructionContext is an interface to support dynamic dispatch.
type IAssemblerinstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblerinstructionContext differentiates from other interfaces.
	IsAssemblerinstructionContext()
}

type AssemblerinstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblerinstructionContext() *AssemblerinstructionContext {
	var p = new(AssemblerinstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_assemblerinstruction
	return p
}

func (*AssemblerinstructionContext) IsAssemblerinstructionContext() {}

func NewAssemblerinstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblerinstructionContext {
	var p = new(AssemblerinstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_assemblerinstruction

	return p
}

func (s *AssemblerinstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssemblerinstructionContext) Assembleropcode() IAssembleropcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssembleropcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssembleropcodeContext)
}

func (s *AssemblerinstructionContext) Argument() IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *AssemblerinstructionContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *AssemblerinstructionContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *AssemblerinstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblerinstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblerinstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterAssemblerinstruction(s)
	}
}

func (s *AssemblerinstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitAssemblerinstruction(s)
	}
}

func (p *asm6502Parser) Assemblerinstruction() (localctx IAssemblerinstructionContext) {
	this := p
	_ = this

	localctx = NewAssemblerinstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, asm6502ParserRULE_assemblerinstruction)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm6502ParserT__2)|(1<<asm6502ParserT__5)|(1<<asm6502ParserT__7))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(asm6502ParserNAME-71))|(1<<(asm6502ParserNUMBER-71))|(1<<(asm6502ParserSTRING-71)))) != 0) {
		{
			p.SetState(54)
			p.Argument()
		}

	}
	{
		p.SetState(57)
		p.Assembleropcode()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm6502ParserT__2)|(1<<asm6502ParserT__5)|(1<<asm6502ParserT__7))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(asm6502ParserNAME-71))|(1<<(asm6502ParserNUMBER-71))|(1<<(asm6502ParserSTRING-71)))) != 0) {
		{
			p.SetState(58)
			p.Argumentlist()
		}

	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm6502ParserCOMMENT {
		{
			p.SetState(61)
			p.Comment()
		}

	}

	return localctx
}

// IAssembleropcodeContext is an interface to support dynamic dispatch.
type IAssembleropcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssembleropcodeContext differentiates from other interfaces.
	IsAssembleropcodeContext()
}

type AssembleropcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssembleropcodeContext() *AssembleropcodeContext {
	var p = new(AssembleropcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_assembleropcode
	return p
}

func (*AssembleropcodeContext) IsAssembleropcodeContext() {}

func NewAssembleropcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssembleropcodeContext {
	var p = new(AssembleropcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_assembleropcode

	return p
}

func (s *AssembleropcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *AssembleropcodeContext) ASSEMBLER_INSTRUCTION() antlr.TerminalNode {
	return s.GetToken(asm6502ParserASSEMBLER_INSTRUCTION, 0)
}

func (s *AssembleropcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssembleropcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssembleropcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterAssembleropcode(s)
	}
}

func (s *AssembleropcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitAssembleropcode(s)
	}
}

func (p *asm6502Parser) Assembleropcode() (localctx IAssembleropcodeContext) {
	this := p
	_ = this

	localctx = NewAssembleropcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, asm6502ParserRULE_assembleropcode)

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
		p.SetState(64)
		p.Match(asm6502ParserASSEMBLER_INSTRUCTION)
	}

	return localctx
}

// ILblContext is an interface to support dynamic dispatch.
type ILblContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLblContext differentiates from other interfaces.
	IsLblContext()
}

type LblContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLblContext() *LblContext {
	var p = new(LblContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_lbl
	return p
}

func (*LblContext) IsLblContext() {}

func NewLblContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LblContext {
	var p = new(LblContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_lbl

	return p
}

func (s *LblContext) GetParser() antlr.Parser { return s.parser }

func (s *LblContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LblContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LblContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LblContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterLbl(s)
	}
}

func (s *LblContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitLbl(s)
	}
}

func (p *asm6502Parser) Lbl() (localctx ILblContext) {
	this := p
	_ = this

	localctx = NewLblContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, asm6502ParserRULE_lbl)

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
		p.Label()
	}
	{
		p.SetState(67)
		p.Match(asm6502ParserT__0)
	}

	return localctx
}

// IArgumentlistContext is an interface to support dynamic dispatch.
type IArgumentlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentlistContext differentiates from other interfaces.
	IsArgumentlistContext()
}

type ArgumentlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentlistContext() *ArgumentlistContext {
	var p = new(ArgumentlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_argumentlist
	return p
}

func (*ArgumentlistContext) IsArgumentlistContext() {}

func NewArgumentlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentlistContext {
	var p = new(ArgumentlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_argumentlist

	return p
}

func (s *ArgumentlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentlistContext) Argument() IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentlistContext) Argumentlist() IArgumentlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentlistContext)
}

func (s *ArgumentlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterArgumentlist(s)
	}
}

func (s *ArgumentlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitArgumentlist(s)
	}
}

func (p *asm6502Parser) Argumentlist() (localctx IArgumentlistContext) {
	this := p
	_ = this

	localctx = NewArgumentlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, asm6502ParserRULE_argumentlist)
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
		p.SetState(69)
		p.Argument()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm6502ParserT__1 {
		{
			p.SetState(70)
			p.Match(asm6502ParserT__1)
		}
		{
			p.SetState(71)
			p.Argumentlist()
		}

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
	p.RuleIndex = asm6502ParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *asm6502Parser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, asm6502ParserRULE_label)

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
		p.SetState(74)
		p.Name()
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
	p.RuleIndex = asm6502ParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *ArgumentContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ArgumentContext) Prefix_() IPrefix_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_Context)
}

func (s *ArgumentContext) Argument() IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *asm6502Parser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, asm6502ParserRULE_argument)
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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm6502ParserT__2, asm6502ParserT__7, asm6502ParserNAME, asm6502ParserNUMBER, asm6502ParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == asm6502ParserT__7 {
			{
				p.SetState(76)
				p.Prefix_()
			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case asm6502ParserNUMBER:
			{
				p.SetState(79)
				p.Number()
			}

		case asm6502ParserNAME:
			{
				p.SetState(80)
				p.Name()
			}

		case asm6502ParserSTRING:
			{
				p.SetState(81)
				p.String_()
			}

		case asm6502ParserT__2:
			{
				p.SetState(82)
				p.Match(asm6502ParserT__2)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == asm6502ParserT__3 || _la == asm6502ParserT__4 {
			{
				p.SetState(85)
				_la = p.GetTokenStream().LA(1)

				if !(_la == asm6502ParserT__3 || _la == asm6502ParserT__4) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(86)
				p.Number()
			}

		}

	case asm6502ParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(asm6502ParserT__5)
		}
		{
			p.SetState(90)
			p.Argument()
		}
		{
			p.SetState(91)
			p.Match(asm6502ParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrefix_Context is an interface to support dynamic dispatch.
type IPrefix_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefix_Context differentiates from other interfaces.
	IsPrefix_Context()
}

type Prefix_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_Context() *Prefix_Context {
	var p = new(Prefix_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_prefix_
	return p
}

func (*Prefix_Context) IsPrefix_Context() {}

func NewPrefix_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_Context {
	var p = new(Prefix_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_prefix_

	return p
}

func (s *Prefix_Context) GetParser() antlr.Parser { return s.parser }
func (s *Prefix_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterPrefix_(s)
	}
}

func (s *Prefix_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitPrefix_(s)
	}
}

func (p *asm6502Parser) Prefix_() (localctx IPrefix_Context) {
	this := p
	_ = this

	localctx = NewPrefix_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, asm6502ParserRULE_prefix_)

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
		p.SetState(95)
		p.Match(asm6502ParserT__7)
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *asm6502Parser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, asm6502ParserRULE_string_)

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
		p.SetState(97)
		p.Match(asm6502ParserSTRING)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm6502ParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(asm6502ParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitName(s)
	}
}

func (p *asm6502Parser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, asm6502ParserRULE_name)

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
		p.SetState(99)
		p.Match(asm6502ParserNAME)
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
	p.RuleIndex = asm6502ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(asm6502ParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *asm6502Parser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, asm6502ParserRULE_number)

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
		p.SetState(101)
		p.Match(asm6502ParserNUMBER)
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
	p.RuleIndex = asm6502ParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *asm6502Parser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, asm6502ParserRULE_comment)

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
		p.SetState(103)
		p.Match(asm6502ParserCOMMENT)
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
	p.RuleIndex = asm6502ParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm6502ParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) ADC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserADC, 0)
}

func (s *OpcodeContext) AND() antlr.TerminalNode {
	return s.GetToken(asm6502ParserAND, 0)
}

func (s *OpcodeContext) ASL() antlr.TerminalNode {
	return s.GetToken(asm6502ParserASL, 0)
}

func (s *OpcodeContext) BCC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBCC, 0)
}

func (s *OpcodeContext) BCS() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBCS, 0)
}

func (s *OpcodeContext) BEQ() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBEQ, 0)
}

func (s *OpcodeContext) BIT() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBIT, 0)
}

func (s *OpcodeContext) BMI() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBMI, 0)
}

func (s *OpcodeContext) BNE() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBNE, 0)
}

func (s *OpcodeContext) BPL() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBPL, 0)
}

func (s *OpcodeContext) BRA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBRA, 0)
}

func (s *OpcodeContext) BRK() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBRK, 0)
}

func (s *OpcodeContext) BVC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBVC, 0)
}

func (s *OpcodeContext) BVS() antlr.TerminalNode {
	return s.GetToken(asm6502ParserBVS, 0)
}

func (s *OpcodeContext) CLC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCLC, 0)
}

func (s *OpcodeContext) CLD() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCLD, 0)
}

func (s *OpcodeContext) CLI() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCLI, 0)
}

func (s *OpcodeContext) CLV() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCLV, 0)
}

func (s *OpcodeContext) CMP() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCMP, 0)
}

func (s *OpcodeContext) CPX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCPX, 0)
}

func (s *OpcodeContext) CPY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserCPY, 0)
}

func (s *OpcodeContext) DEC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserDEC, 0)
}

func (s *OpcodeContext) DEX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserDEX, 0)
}

func (s *OpcodeContext) DEY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserDEY, 0)
}

func (s *OpcodeContext) EOR() antlr.TerminalNode {
	return s.GetToken(asm6502ParserEOR, 0)
}

func (s *OpcodeContext) INC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserINC, 0)
}

func (s *OpcodeContext) INX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserINX, 0)
}

func (s *OpcodeContext) INY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserINY, 0)
}

func (s *OpcodeContext) JMP() antlr.TerminalNode {
	return s.GetToken(asm6502ParserJMP, 0)
}

func (s *OpcodeContext) JSR() antlr.TerminalNode {
	return s.GetToken(asm6502ParserJSR, 0)
}

func (s *OpcodeContext) LDA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserLDA, 0)
}

func (s *OpcodeContext) LDY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserLDY, 0)
}

func (s *OpcodeContext) LDX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserLDX, 0)
}

func (s *OpcodeContext) LSR() antlr.TerminalNode {
	return s.GetToken(asm6502ParserLSR, 0)
}

func (s *OpcodeContext) NOP() antlr.TerminalNode {
	return s.GetToken(asm6502ParserNOP, 0)
}

func (s *OpcodeContext) ORA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserORA, 0)
}

func (s *OpcodeContext) PHA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPHA, 0)
}

func (s *OpcodeContext) PHX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPHX, 0)
}

func (s *OpcodeContext) PHY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPHY, 0)
}

func (s *OpcodeContext) PHP() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPHP, 0)
}

func (s *OpcodeContext) PLA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPLA, 0)
}

func (s *OpcodeContext) PLP() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPLP, 0)
}

func (s *OpcodeContext) PLY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserPLY, 0)
}

func (s *OpcodeContext) ROL() antlr.TerminalNode {
	return s.GetToken(asm6502ParserROL, 0)
}

func (s *OpcodeContext) ROR() antlr.TerminalNode {
	return s.GetToken(asm6502ParserROR, 0)
}

func (s *OpcodeContext) RTI() antlr.TerminalNode {
	return s.GetToken(asm6502ParserRTI, 0)
}

func (s *OpcodeContext) RTS() antlr.TerminalNode {
	return s.GetToken(asm6502ParserRTS, 0)
}

func (s *OpcodeContext) SBC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSBC, 0)
}

func (s *OpcodeContext) SEC() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSEC, 0)
}

func (s *OpcodeContext) SED() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSED, 0)
}

func (s *OpcodeContext) SEI() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSEI, 0)
}

func (s *OpcodeContext) STA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSTA, 0)
}

func (s *OpcodeContext) STX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSTX, 0)
}

func (s *OpcodeContext) STY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSTY, 0)
}

func (s *OpcodeContext) STZ() antlr.TerminalNode {
	return s.GetToken(asm6502ParserSTZ, 0)
}

func (s *OpcodeContext) TAX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTAX, 0)
}

func (s *OpcodeContext) TAY() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTAY, 0)
}

func (s *OpcodeContext) TSX() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTSX, 0)
}

func (s *OpcodeContext) TXA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTXA, 0)
}

func (s *OpcodeContext) TXS() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTXS, 0)
}

func (s *OpcodeContext) TYA() antlr.TerminalNode {
	return s.GetToken(asm6502ParserTYA, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm6502Listener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *asm6502Parser) Opcode() (localctx IOpcodeContext) {
	this := p
	_ = this

	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, asm6502ParserRULE_opcode)
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
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(asm6502ParserADC-10))|(1<<(asm6502ParserAND-10))|(1<<(asm6502ParserASL-10))|(1<<(asm6502ParserBCC-10))|(1<<(asm6502ParserBCS-10))|(1<<(asm6502ParserBEQ-10))|(1<<(asm6502ParserBIT-10))|(1<<(asm6502ParserBMI-10))|(1<<(asm6502ParserBNE-10))|(1<<(asm6502ParserBPL-10))|(1<<(asm6502ParserBRA-10))|(1<<(asm6502ParserBRK-10))|(1<<(asm6502ParserBVC-10))|(1<<(asm6502ParserBVS-10))|(1<<(asm6502ParserCLC-10))|(1<<(asm6502ParserCLD-10))|(1<<(asm6502ParserCLI-10))|(1<<(asm6502ParserCLV-10))|(1<<(asm6502ParserCMP-10))|(1<<(asm6502ParserCPX-10))|(1<<(asm6502ParserCPY-10))|(1<<(asm6502ParserDEC-10))|(1<<(asm6502ParserDEX-10))|(1<<(asm6502ParserDEY-10))|(1<<(asm6502ParserEOR-10))|(1<<(asm6502ParserINC-10))|(1<<(asm6502ParserINX-10))|(1<<(asm6502ParserINY-10))|(1<<(asm6502ParserJMP-10))|(1<<(asm6502ParserJSR-10))|(1<<(asm6502ParserLDA-10))|(1<<(asm6502ParserLDY-10)))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(asm6502ParserLDX-42))|(1<<(asm6502ParserLSR-42))|(1<<(asm6502ParserNOP-42))|(1<<(asm6502ParserORA-42))|(1<<(asm6502ParserPHA-42))|(1<<(asm6502ParserPHX-42))|(1<<(asm6502ParserPHY-42))|(1<<(asm6502ParserPHP-42))|(1<<(asm6502ParserPLA-42))|(1<<(asm6502ParserPLP-42))|(1<<(asm6502ParserPLY-42))|(1<<(asm6502ParserROL-42))|(1<<(asm6502ParserROR-42))|(1<<(asm6502ParserRTI-42))|(1<<(asm6502ParserRTS-42))|(1<<(asm6502ParserSBC-42))|(1<<(asm6502ParserSEC-42))|(1<<(asm6502ParserSED-42))|(1<<(asm6502ParserSEI-42))|(1<<(asm6502ParserSTA-42))|(1<<(asm6502ParserSTX-42))|(1<<(asm6502ParserSTY-42))|(1<<(asm6502ParserSTZ-42))|(1<<(asm6502ParserTAX-42))|(1<<(asm6502ParserTAY-42))|(1<<(asm6502ParserTSX-42))|(1<<(asm6502ParserTXA-42))|(1<<(asm6502ParserTXS-42))|(1<<(asm6502ParserTYA-42)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
