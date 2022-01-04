// Code generated from brainfuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 28, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 7, 2, 10, 10, 2, 12, 2, 14, 2, 13,
	11, 2, 3, 3, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3, 3,
	5, 3, 24, 10, 3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3, 3, 2, 3, 8,
	2, 27, 2, 11, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 10,
	5, 4, 3, 2, 9, 8, 3, 2, 2, 2, 10, 13, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 11,
	12, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 14, 24, 5, 6, 4,
	2, 15, 19, 7, 9, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2, 18, 21,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 22, 24, 7, 10, 2, 2, 23, 14, 3, 2, 2, 2, 23, 15, 3,
	2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26, 9, 2, 2, 2, 26, 7, 3, 2, 2, 2, 5, 11,
	19, 23,
}
var literalNames = []string{
	"", "'>'", "'<'", "'+'", "'-'", "'.'", "','", "'['", "']'",
}
var symbolicNames = []string{
	"", "GT", "LT", "PLUS", "MINUS", "DOT", "COMMA", "LPAREN", "RPAREN", "WS",
}

var ruleNames = []string{
	"file_", "statement", "opcode",
}

type brainfuckParser struct {
	*antlr.BaseParser
}

// NewbrainfuckParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *brainfuckParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbrainfuckParser(input antlr.TokenStream) *brainfuckParser {
	this := new(brainfuckParser)
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
	brainfuckParserRULE_file_     = 0
	brainfuckParserRULE_statement = 1
	brainfuckParserRULE_opcode    = 2
)

// IFile_Context is an interface to support dynamic dispatch.
type IFile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_Context differentiates from other interfaces.
	IsFile_Context()
}

type File_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_Context() *File_Context {
	var p = new(File_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainfuckParserRULE_file_
	return p
}

func (*File_Context) IsFile_Context() {}

func NewFile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_Context {
	var p = new(File_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainfuckParserRULE_file_

	return p
}

func (s *File_Context) GetParser() antlr.Parser { return s.parser }

func (s *File_Context) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *File_Context) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *File_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.EnterFile_(s)
	}
}

func (s *File_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.ExitFile_(s)
	}
}

func (p *brainfuckParser) File_() (localctx IFile_Context) {
	this := p
	_ = this

	localctx = NewFile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, brainfuckParserRULE_file_)
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
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainfuckParserGT)|(1<<brainfuckParserLT)|(1<<brainfuckParserPLUS)|(1<<brainfuckParserMINUS)|(1<<brainfuckParserDOT)|(1<<brainfuckParserCOMMA)|(1<<brainfuckParserLPAREN))) != 0 {
		{
			p.SetState(6)
			p.Statement()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = brainfuckParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainfuckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(brainfuckParserLPAREN, 0)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(brainfuckParserRPAREN, 0)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainfuckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *brainfuckParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, brainfuckParserRULE_statement)
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

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case brainfuckParserGT, brainfuckParserLT, brainfuckParserPLUS, brainfuckParserMINUS, brainfuckParserDOT, brainfuckParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Opcode()
		}

	case brainfuckParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Match(brainfuckParserLPAREN)
		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainfuckParserGT)|(1<<brainfuckParserLT)|(1<<brainfuckParserPLUS)|(1<<brainfuckParserMINUS)|(1<<brainfuckParserDOT)|(1<<brainfuckParserCOMMA)|(1<<brainfuckParserLPAREN))) != 0 {
			{
				p.SetState(14)
				p.Statement()
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(20)
			p.Match(brainfuckParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (p *brainfuckParser) Opcode() (localctx IOpcodeContext) {
	this := p
	_ = this

	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, brainfuckParserRULE_opcode)
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
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainfuckParserGT)|(1<<brainfuckParserLT)|(1<<brainfuckParserPLUS)|(1<<brainfuckParserMINUS)|(1<<brainfuckParserDOT)|(1<<brainfuckParserCOMMA))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
