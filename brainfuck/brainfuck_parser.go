// Generated from brainfuck.g4 by ANTLR 4.7.

package brainfuck // brainfuck
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 14, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 6, 2, 8, 10, 2, 13, 2, 14, 2, 9, 3, 3, 3, 3,
	3, 3, 2, 2, 4, 2, 4, 2, 3, 3, 2, 3, 10, 2, 12, 2, 7, 3, 2, 2, 2, 4, 11,
	3, 2, 2, 2, 6, 8, 5, 4, 3, 2, 7, 6, 3, 2, 2, 2, 8, 9, 3, 2, 2, 2, 9, 7,
	3, 2, 2, 2, 9, 10, 3, 2, 2, 2, 10, 3, 3, 2, 2, 2, 11, 12, 9, 2, 2, 2, 12,
	5, 3, 2, 2, 2, 3, 9,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'>'", "'<'", "'+'", "'-'", "'.'", "','", "'['", "']'",
}
var symbolicNames = []string{
	"", "GT", "LT", "PLUS", "MINUS", "DOT", "COMMA", "LPAREN", "RPAREN", "WS",
}

var ruleNames = []string{
	"file", "opcode",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type brainfuckParser struct {
	*antlr.BaseParser
}

func NewbrainfuckParser(input antlr.TokenStream) *brainfuckParser {
	this := new(brainfuckParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "brainfuck.g4"

	return this
}

// brainfuckParser tokens.
const (
	brainfuckParserEOF    = antlr.TokenEOF
	brainfuckParserGT     = 1
	brainfuckParserLT     = 2
	brainfuckParserPLUS   = 3
	brainfuckParserMINUS  = 4
	brainfuckParserDOT    = 5
	brainfuckParserCOMMA  = 6
	brainfuckParserLPAREN = 7
	brainfuckParserRPAREN = 8
	brainfuckParserWS     = 9
)

// brainfuckParser rules.
const (
	brainfuckParserRULE_file   = 0
	brainfuckParserRULE_opcode = 1
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
	p.RuleIndex = brainfuckParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainfuckParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllOpcode() []IOpcodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpcodeContext)(nil)).Elem())
	var tst = make([]IOpcodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpcodeContext)
		}
	}

	return tst
}

func (s *FileContext) Opcode(i int) IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case brainfuckVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *brainfuckParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, brainfuckParserRULE_file)
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
	p.SetState(5)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainfuckParserGT)|(1<<brainfuckParserLT)|(1<<brainfuckParserPLUS)|(1<<brainfuckParserMINUS)|(1<<brainfuckParserDOT)|(1<<brainfuckParserCOMMA)|(1<<brainfuckParserLPAREN)|(1<<brainfuckParserRPAREN))) != 0) {
		{
			p.SetState(4)
			p.Opcode()
		}

		p.SetState(7)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = brainfuckParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainfuckParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) GT() antlr.TerminalNode {
	return s.GetToken(brainfuckParserGT, 0)
}

func (s *OpcodeContext) LT() antlr.TerminalNode {
	return s.GetToken(brainfuckParserLT, 0)
}

func (s *OpcodeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(brainfuckParserPLUS, 0)
}

func (s *OpcodeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(brainfuckParserMINUS, 0)
}

func (s *OpcodeContext) DOT() antlr.TerminalNode {
	return s.GetToken(brainfuckParserDOT, 0)
}

func (s *OpcodeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(brainfuckParserCOMMA, 0)
}

func (s *OpcodeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(brainfuckParserLPAREN, 0)
}

func (s *OpcodeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(brainfuckParserRPAREN, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (s *OpcodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case brainfuckVisitor:
		return t.VisitOpcode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *brainfuckParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, brainfuckParserRULE_opcode)
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
	p.SetState(9)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainfuckParserGT)|(1<<brainfuckParserLT)|(1<<brainfuckParserPLUS)|(1<<brainfuckParserMINUS)|(1<<brainfuckParserDOT)|(1<<brainfuckParserCOMMA)|(1<<brainfuckParserLPAREN)|(1<<brainfuckParserRPAREN))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
