// Code generated from lrcParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lrc // lrcParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 20, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 7, 2, 8, 10, 2, 12, 2, 14, 2, 11, 11, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 4, 2, 4, 2, 2, 2, 18,
	2, 9, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 8, 5, 4, 3, 2, 7, 6, 3, 2, 2, 2,
	8, 11, 3, 2, 2, 2, 9, 7, 3, 2, 2, 2, 9, 10, 3, 2, 2, 2, 10, 12, 3, 2, 2,
	2, 11, 9, 3, 2, 2, 2, 12, 13, 7, 2, 2, 3, 13, 3, 3, 2, 2, 2, 14, 15, 7,
	3, 2, 2, 15, 16, 7, 5, 2, 2, 16, 17, 7, 4, 2, 2, 17, 18, 7, 8, 2, 2, 18,
	5, 3, 2, 2, 2, 3, 9,
}
var literalNames = []string{
	"", "'['", "']'",
}
var symbolicNames = []string{
	"", "LB", "RB", "TIMESTAMP", "NUM", "WS", "TEXT", "EOL",
}

var ruleNames = []string{
	"lrc", "line",
}

type lrcParser struct {
	*antlr.BaseParser
}

// NewlrcParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *lrcParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlrcParser(input antlr.TokenStream) *lrcParser {
	this := new(lrcParser)
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
	this.GrammarFileName = "lrcParser.g4"

	return this
}

// lrcParser tokens.
const (
	lrcParserEOF       = antlr.TokenEOF
	lrcParserLB        = 1
	lrcParserRB        = 2
	lrcParserTIMESTAMP = 3
	lrcParserNUM       = 4
	lrcParserWS        = 5
	lrcParserTEXT      = 6
	lrcParserEOL       = 7
)

// lrcParser rules.
const (
	lrcParserRULE_lrc  = 0
	lrcParserRULE_line = 1
)

// ILrcContext is an interface to support dynamic dispatch.
type ILrcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLrcContext differentiates from other interfaces.
	IsLrcContext()
}

type LrcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLrcContext() *LrcContext {
	var p = new(LrcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lrcParserRULE_lrc
	return p
}

func (*LrcContext) IsLrcContext() {}

func NewLrcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LrcContext {
	var p = new(LrcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lrcParserRULE_lrc

	return p
}

func (s *LrcContext) GetParser() antlr.Parser { return s.parser }

func (s *LrcContext) EOF() antlr.TerminalNode {
	return s.GetToken(lrcParserEOF, 0)
}

func (s *LrcContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *LrcContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *LrcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LrcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lrcParserListener); ok {
		listenerT.EnterLrc(s)
	}
}

func (s *LrcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lrcParserListener); ok {
		listenerT.ExitLrc(s)
	}
}

func (p *lrcParser) Lrc() (localctx ILrcContext) {
	this := p
	_ = this

	localctx = NewLrcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lrcParserRULE_lrc)
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
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == lrcParserLB {
		{
			p.SetState(4)
			p.Line()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(10)
		p.Match(lrcParserEOF)
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
	p.RuleIndex = lrcParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lrcParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) LB() antlr.TerminalNode {
	return s.GetToken(lrcParserLB, 0)
}

func (s *LineContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(lrcParserTIMESTAMP, 0)
}

func (s *LineContext) RB() antlr.TerminalNode {
	return s.GetToken(lrcParserRB, 0)
}

func (s *LineContext) TEXT() antlr.TerminalNode {
	return s.GetToken(lrcParserTEXT, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lrcParserListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lrcParserListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *lrcParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lrcParserRULE_line)

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
		p.Match(lrcParserLB)
	}
	{
		p.SetState(13)
		p.Match(lrcParserTIMESTAMP)
	}
	{
		p.SetState(14)
		p.Match(lrcParserRB)
	}
	{
		p.SetState(15)
		p.Match(lrcParserTEXT)
	}

	return localctx
}
