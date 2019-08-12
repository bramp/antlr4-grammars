// Code generated from redcode.g4 by ANTLR 4.7.2. DO NOT EDIT.

package redcode // redcode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 62, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 3, 3,
	3, 5, 3, 26, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 33, 10, 4, 3, 4,
	5, 4, 36, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 4, 5, 4, 44, 10,
	4, 3, 4, 5, 4, 47, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 5,
	8, 56, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12,
	14, 16, 2, 6, 3, 2, 19, 34, 3, 2, 12, 18, 3, 2, 5, 9, 3, 2, 10, 11, 2,
	61, 2, 19, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 48, 3,
	2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 52, 3, 2, 2, 2, 14, 55, 3, 2, 2, 2, 16,
	59, 3, 2, 2, 2, 18, 20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2,
	2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 26, 5,
	16, 9, 2, 24, 26, 5, 6, 4, 2, 25, 23, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26,
	27, 3, 2, 2, 2, 27, 28, 7, 37, 2, 2, 28, 5, 3, 2, 2, 2, 29, 32, 5, 8, 5,
	2, 30, 31, 7, 3, 2, 2, 31, 33, 5, 10, 6, 2, 32, 30, 3, 2, 2, 2, 32, 33,
	3, 2, 2, 2, 33, 35, 3, 2, 2, 2, 34, 36, 5, 12, 7, 2, 35, 34, 3, 2, 2, 2,
	35, 36, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 43, 5, 14, 8, 2, 38, 40, 7,
	4, 2, 2, 39, 41, 5, 12, 7, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41,
	42, 3, 2, 2, 2, 42, 44, 5, 14, 8, 2, 43, 38, 3, 2, 2, 2, 43, 44, 3, 2,
	2, 2, 44, 46, 3, 2, 2, 2, 45, 47, 5, 16, 9, 2, 46, 45, 3, 2, 2, 2, 46,
	47, 3, 2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 49, 9, 2, 2, 2, 49, 9, 3, 2, 2,
	2, 50, 51, 9, 3, 2, 2, 51, 11, 3, 2, 2, 2, 52, 53, 9, 4, 2, 2, 53, 13,
	3, 2, 2, 2, 54, 56, 9, 5, 2, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 57, 3, 2, 2, 2, 57, 58, 7, 35, 2, 2, 58, 15, 3, 2, 2, 2, 59, 60, 7,
	36, 2, 2, 60, 17, 3, 2, 2, 2, 10, 21, 25, 32, 35, 40, 43, 46, 55,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "','", "'#'", "'$'", "'@'", "'<'", "'>'", "'+'", "'-'", "'A'",
	"'B'", "'AB'", "'BA'", "'F'", "'X'", "'I'", "'DAT'", "'MOV'", "'ADD'",
	"'SUB'", "'MUL'", "'DIV'", "'MOD'", "'JMP'", "'JMZ'", "'JMN'", "'DJN'",
	"'CMP'", "'SLT'", "'DJZ'", "'SPL'", "'ORG'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "A", "B", "AB", "BA", "F", "X",
	"I", "DAT", "MOV", "ADD", "SUB", "MUL", "DIV", "MOD", "JMP", "JMZ", "JMN",
	"DJN", "CMP", "SLT", "DJZ", "SPL", "ORG", "NUMBER", "COMMENT", "EOL", "WS",
}

var ruleNames = []string{
	"file", "line", "instruction", "opcode", "modifier", "mmode", "number",
	"comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type redcodeParser struct {
	*antlr.BaseParser
}

func NewredcodeParser(input antlr.TokenStream) *redcodeParser {
	this := new(redcodeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "redcode.g4"

	return this
}

// redcodeParser tokens.
const (
	redcodeParserEOF     = antlr.TokenEOF
	redcodeParserT__0    = 1
	redcodeParserT__1    = 2
	redcodeParserT__2    = 3
	redcodeParserT__3    = 4
	redcodeParserT__4    = 5
	redcodeParserT__5    = 6
	redcodeParserT__6    = 7
	redcodeParserT__7    = 8
	redcodeParserT__8    = 9
	redcodeParserA       = 10
	redcodeParserB       = 11
	redcodeParserAB      = 12
	redcodeParserBA      = 13
	redcodeParserF       = 14
	redcodeParserX       = 15
	redcodeParserI       = 16
	redcodeParserDAT     = 17
	redcodeParserMOV     = 18
	redcodeParserADD     = 19
	redcodeParserSUB     = 20
	redcodeParserMUL     = 21
	redcodeParserDIV     = 22
	redcodeParserMOD     = 23
	redcodeParserJMP     = 24
	redcodeParserJMZ     = 25
	redcodeParserJMN     = 26
	redcodeParserDJN     = 27
	redcodeParserCMP     = 28
	redcodeParserSLT     = 29
	redcodeParserDJZ     = 30
	redcodeParserSPL     = 31
	redcodeParserORG     = 32
	redcodeParserNUMBER  = 33
	redcodeParserCOMMENT = 34
	redcodeParserEOL     = 35
	redcodeParserWS      = 36
)

// redcodeParser rules.
const (
	redcodeParserRULE_file        = 0
	redcodeParserRULE_line        = 1
	redcodeParserRULE_instruction = 2
	redcodeParserRULE_opcode      = 3
	redcodeParserRULE_modifier    = 4
	redcodeParserRULE_mmode       = 5
	redcodeParserRULE_number      = 6
	redcodeParserRULE_comment     = 7
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = redcodeParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *FileContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *redcodeParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, redcodeParserRULE_file)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(redcodeParserDAT-17))|(1<<(redcodeParserMOV-17))|(1<<(redcodeParserADD-17))|(1<<(redcodeParserSUB-17))|(1<<(redcodeParserMUL-17))|(1<<(redcodeParserDIV-17))|(1<<(redcodeParserMOD-17))|(1<<(redcodeParserJMP-17))|(1<<(redcodeParserJMZ-17))|(1<<(redcodeParserJMN-17))|(1<<(redcodeParserDJN-17))|(1<<(redcodeParserCMP-17))|(1<<(redcodeParserSLT-17))|(1<<(redcodeParserDJZ-17))|(1<<(redcodeParserSPL-17))|(1<<(redcodeParserORG-17))|(1<<(redcodeParserCOMMENT-17)))) != 0) {
		{
			p.SetState(16)
			p.Line()
		}

		p.SetState(19)
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
	p.RuleIndex = redcodeParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) EOL() antlr.TerminalNode {
	return s.GetToken(redcodeParserEOL, 0)
}

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

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *redcodeParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, redcodeParserRULE_line)

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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case redcodeParserCOMMENT:
		{
			p.SetState(21)
			p.Comment()
		}

	case redcodeParserDAT, redcodeParserMOV, redcodeParserADD, redcodeParserSUB, redcodeParserMUL, redcodeParserDIV, redcodeParserMOD, redcodeParserJMP, redcodeParserJMZ, redcodeParserJMN, redcodeParserDJN, redcodeParserCMP, redcodeParserSLT, redcodeParserDJZ, redcodeParserSPL, redcodeParserORG:
		{
			p.SetState(22)
			p.Instruction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(25)
		p.Match(redcodeParserEOL)
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
	p.RuleIndex = redcodeParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_instruction

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

func (s *InstructionContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *InstructionContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *InstructionContext) Modifier() IModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *InstructionContext) AllMmode() []IMmodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMmodeContext)(nil)).Elem())
	var tst = make([]IMmodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMmodeContext)
		}
	}

	return tst
}

func (s *InstructionContext) Mmode(i int) IMmodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMmodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMmodeContext)
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
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *redcodeParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, redcodeParserRULE_instruction)
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
		p.SetState(27)
		p.Opcode()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == redcodeParserT__0 {
		{
			p.SetState(28)
			p.Match(redcodeParserT__0)
		}
		{
			p.SetState(29)
			p.Modifier()
		}

	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<redcodeParserT__2)|(1<<redcodeParserT__3)|(1<<redcodeParserT__4)|(1<<redcodeParserT__5)|(1<<redcodeParserT__6))) != 0 {
		{
			p.SetState(32)
			p.Mmode()
		}

	}
	{
		p.SetState(35)
		p.Number()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == redcodeParserT__1 {
		{
			p.SetState(36)
			p.Match(redcodeParserT__1)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<redcodeParserT__2)|(1<<redcodeParserT__3)|(1<<redcodeParserT__4)|(1<<redcodeParserT__5)|(1<<redcodeParserT__6))) != 0 {
			{
				p.SetState(37)
				p.Mmode()
			}

		}
		{
			p.SetState(40)
			p.Number()
		}

	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == redcodeParserCOMMENT {
		{
			p.SetState(43)
			p.Comment()
		}

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
	p.RuleIndex = redcodeParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) DAT() antlr.TerminalNode {
	return s.GetToken(redcodeParserDAT, 0)
}

func (s *OpcodeContext) MOV() antlr.TerminalNode {
	return s.GetToken(redcodeParserMOV, 0)
}

func (s *OpcodeContext) ADD() antlr.TerminalNode {
	return s.GetToken(redcodeParserADD, 0)
}

func (s *OpcodeContext) SUB() antlr.TerminalNode {
	return s.GetToken(redcodeParserSUB, 0)
}

func (s *OpcodeContext) MUL() antlr.TerminalNode {
	return s.GetToken(redcodeParserMUL, 0)
}

func (s *OpcodeContext) DIV() antlr.TerminalNode {
	return s.GetToken(redcodeParserDIV, 0)
}

func (s *OpcodeContext) MOD() antlr.TerminalNode {
	return s.GetToken(redcodeParserMOD, 0)
}

func (s *OpcodeContext) JMP() antlr.TerminalNode {
	return s.GetToken(redcodeParserJMP, 0)
}

func (s *OpcodeContext) JMZ() antlr.TerminalNode {
	return s.GetToken(redcodeParserJMZ, 0)
}

func (s *OpcodeContext) JMN() antlr.TerminalNode {
	return s.GetToken(redcodeParserJMN, 0)
}

func (s *OpcodeContext) DJN() antlr.TerminalNode {
	return s.GetToken(redcodeParserDJN, 0)
}

func (s *OpcodeContext) CMP() antlr.TerminalNode {
	return s.GetToken(redcodeParserCMP, 0)
}

func (s *OpcodeContext) SLT() antlr.TerminalNode {
	return s.GetToken(redcodeParserSLT, 0)
}

func (s *OpcodeContext) SPL() antlr.TerminalNode {
	return s.GetToken(redcodeParserSPL, 0)
}

func (s *OpcodeContext) ORG() antlr.TerminalNode {
	return s.GetToken(redcodeParserORG, 0)
}

func (s *OpcodeContext) DJZ() antlr.TerminalNode {
	return s.GetToken(redcodeParserDJZ, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *redcodeParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, redcodeParserRULE_opcode)
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
		p.SetState(46)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(redcodeParserDAT-17))|(1<<(redcodeParserMOV-17))|(1<<(redcodeParserADD-17))|(1<<(redcodeParserSUB-17))|(1<<(redcodeParserMUL-17))|(1<<(redcodeParserDIV-17))|(1<<(redcodeParserMOD-17))|(1<<(redcodeParserJMP-17))|(1<<(redcodeParserJMZ-17))|(1<<(redcodeParserJMN-17))|(1<<(redcodeParserDJN-17))|(1<<(redcodeParserCMP-17))|(1<<(redcodeParserSLT-17))|(1<<(redcodeParserDJZ-17))|(1<<(redcodeParserSPL-17))|(1<<(redcodeParserORG-17)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = redcodeParserRULE_modifier
	return p
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) A() antlr.TerminalNode {
	return s.GetToken(redcodeParserA, 0)
}

func (s *ModifierContext) B() antlr.TerminalNode {
	return s.GetToken(redcodeParserB, 0)
}

func (s *ModifierContext) AB() antlr.TerminalNode {
	return s.GetToken(redcodeParserAB, 0)
}

func (s *ModifierContext) BA() antlr.TerminalNode {
	return s.GetToken(redcodeParserBA, 0)
}

func (s *ModifierContext) F() antlr.TerminalNode {
	return s.GetToken(redcodeParserF, 0)
}

func (s *ModifierContext) X() antlr.TerminalNode {
	return s.GetToken(redcodeParserX, 0)
}

func (s *ModifierContext) I() antlr.TerminalNode {
	return s.GetToken(redcodeParserI, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (p *redcodeParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, redcodeParserRULE_modifier)
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
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<redcodeParserA)|(1<<redcodeParserB)|(1<<redcodeParserAB)|(1<<redcodeParserBA)|(1<<redcodeParserF)|(1<<redcodeParserX)|(1<<redcodeParserI))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMmodeContext is an interface to support dynamic dispatch.
type IMmodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMmodeContext differentiates from other interfaces.
	IsMmodeContext()
}

type MmodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMmodeContext() *MmodeContext {
	var p = new(MmodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = redcodeParserRULE_mmode
	return p
}

func (*MmodeContext) IsMmodeContext() {}

func NewMmodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MmodeContext {
	var p = new(MmodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_mmode

	return p
}

func (s *MmodeContext) GetParser() antlr.Parser { return s.parser }
func (s *MmodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MmodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MmodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterMmode(s)
	}
}

func (s *MmodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitMmode(s)
	}
}

func (p *redcodeParser) Mmode() (localctx IMmodeContext) {
	localctx = NewMmodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, redcodeParserRULE_mmode)
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
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<redcodeParserT__2)|(1<<redcodeParserT__3)|(1<<redcodeParserT__4)|(1<<redcodeParserT__5)|(1<<redcodeParserT__6))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = redcodeParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(redcodeParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *redcodeParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, redcodeParserRULE_number)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == redcodeParserT__7 || _la == redcodeParserT__8 {
		{
			p.SetState(52)
			_la = p.GetTokenStream().LA(1)

			if !(_la == redcodeParserT__7 || _la == redcodeParserT__8) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(55)
		p.Match(redcodeParserNUMBER)
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
	p.RuleIndex = redcodeParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = redcodeParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(redcodeParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(redcodeListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *redcodeParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, redcodeParserRULE_comment)

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
		p.Match(redcodeParserCOMMENT)
	}

	return localctx
}
