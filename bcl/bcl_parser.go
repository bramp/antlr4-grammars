// Code generated from bcl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcl // bcl
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 23, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 7, 2, 8, 10, 2, 12, 2, 14, 2, 11, 11, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 2,
	2, 4, 2, 4, 2, 2, 2, 23, 2, 9, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 8, 5,
	4, 3, 2, 7, 6, 3, 2, 2, 2, 8, 11, 3, 2, 2, 2, 9, 7, 3, 2, 2, 2, 9, 10,
	3, 2, 2, 2, 10, 12, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 12, 13, 7, 2, 2, 3,
	13, 3, 3, 2, 2, 2, 14, 21, 7, 3, 2, 2, 15, 21, 7, 4, 2, 2, 16, 17, 7, 5,
	2, 2, 17, 18, 5, 4, 3, 2, 18, 19, 5, 4, 3, 2, 19, 21, 3, 2, 2, 2, 20, 14,
	3, 2, 2, 2, 20, 15, 3, 2, 2, 2, 20, 16, 3, 2, 2, 2, 21, 5, 3, 2, 2, 2,
	4, 9, 20,
}
var literalNames = []string{
	"", "'00'", "'01'", "'1'",
}
var symbolicNames = []string{
	"", "", "", "", "WS",
}

var ruleNames = []string{
	"bcl", "term",
}

type bclParser struct {
	*antlr.BaseParser
}

// NewbclParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *bclParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbclParser(input antlr.TokenStream) *bclParser {
	this := new(bclParser)
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
	this.GrammarFileName = "bcl.g4"

	return this
}

// bclParser tokens.
const (
	bclParserEOF  = antlr.TokenEOF
	bclParserT__0 = 1
	bclParserT__1 = 2
	bclParserT__2 = 3
	bclParserWS   = 4
)

// bclParser rules.
const (
	bclParserRULE_bcl  = 0
	bclParserRULE_term = 1
)

// IBclContext is an interface to support dynamic dispatch.
type IBclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBclContext differentiates from other interfaces.
	IsBclContext()
}

type BclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBclContext() *BclContext {
	var p = new(BclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bclParserRULE_bcl
	return p
}

func (*BclContext) IsBclContext() {}

func NewBclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BclContext {
	var p = new(BclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bclParserRULE_bcl

	return p
}

func (s *BclContext) GetParser() antlr.Parser { return s.parser }

func (s *BclContext) EOF() antlr.TerminalNode {
	return s.GetToken(bclParserEOF, 0)
}

func (s *BclContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BclContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bclListener); ok {
		listenerT.EnterBcl(s)
	}
}

func (s *BclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bclListener); ok {
		listenerT.ExitBcl(s)
	}
}

func (p *bclParser) Bcl() (localctx IBclContext) {
	this := p
	_ = this

	localctx = NewBclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bclParserRULE_bcl)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<bclParserT__0)|(1<<bclParserT__1)|(1<<bclParserT__2))) != 0 {
		{
			p.SetState(4)
			p.Term()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(10)
		p.Match(bclParserEOF)
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
	p.RuleIndex = bclParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bclParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TermContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bclListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bclListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *bclParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bclParserRULE_term)

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

	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bclParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Match(bclParserT__0)
		}

	case bclParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Match(bclParserT__1)
		}

	case bclParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.Match(bclParserT__2)
		}
		{
			p.SetState(15)
			p.Term()
		}
		{
			p.SetState(16)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
