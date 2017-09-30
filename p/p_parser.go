// Generated from p.g4 by ANTLR 4.7.

package p // p
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 30, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 6, 2, 12, 10, 2, 13,
	2, 14, 2, 13, 3, 3, 3, 3, 5, 3, 18, 10, 3, 3, 4, 3, 4, 6, 4, 22, 10, 4,
	13, 4, 14, 4, 23, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2,
	3, 3, 2, 5, 6, 2, 28, 2, 11, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 19, 3, 2,
	2, 2, 8, 27, 3, 2, 2, 2, 10, 12, 5, 4, 3, 2, 11, 10, 3, 2, 2, 2, 12, 13,
	3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 3, 3, 2, 2, 2,
	15, 18, 5, 6, 4, 2, 16, 18, 5, 8, 5, 2, 17, 15, 3, 2, 2, 2, 17, 16, 3,
	2, 2, 2, 18, 5, 3, 2, 2, 2, 19, 21, 7, 3, 2, 2, 20, 22, 5, 4, 3, 2, 21,
	20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2,
	2, 24, 25, 3, 2, 2, 2, 25, 26, 7, 4, 2, 2, 26, 7, 3, 2, 2, 2, 27, 28, 9,
	2, 2, 2, 28, 9, 3, 2, 2, 2, 5, 13, 17, 23,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'R'", "'\u03BB'",
}
var symbolicNames = []string{
	"", "", "", "R", "L", "WS",
}

var ruleNames = []string{
	"prog", "symbol", "iterate", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type pParser struct {
	*antlr.BaseParser
}

func NewpParser(input antlr.TokenStream) *pParser {
	this := new(pParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "p.g4"

	return this
}

// pParser tokens.
const (
	pParserEOF  = antlr.TokenEOF
	pParserT__0 = 1
	pParserT__1 = 2
	pParserR    = 3
	pParserL    = 4
	pParserWS   = 5
)

// pParser rules.
const (
	pParserRULE_prog    = 0
	pParserRULE_symbol  = 1
	pParserRULE_iterate = 2
	pParserRULE_atom    = 3
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
	p.RuleIndex = pParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllSymbol() []ISymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolContext)(nil)).Elem())
	var tst = make([]ISymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolContext)
		}
	}

	return tst
}

func (s *ProgContext) Symbol(i int) ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pParserRULE_prog)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pParserT__0)|(1<<pParserR)|(1<<pParserL))) != 0) {
		{
			p.SetState(8)
			p.Symbol()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Iterate() IIterateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterateContext)
}

func (s *SymbolContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (s *SymbolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pVisitor:
		return t.VisitSymbol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, pParserRULE_symbol)

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

	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(13)
			p.Iterate()
		}

	case pParserR, pParserL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(14)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIterateContext is an interface to support dynamic dispatch.
type IIterateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterateContext differentiates from other interfaces.
	IsIterateContext()
}

type IterateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterateContext() *IterateContext {
	var p = new(IterateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pParserRULE_iterate
	return p
}

func (*IterateContext) IsIterateContext() {}

func NewIterateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterateContext {
	var p = new(IterateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pParserRULE_iterate

	return p
}

func (s *IterateContext) GetParser() antlr.Parser { return s.parser }

func (s *IterateContext) AllSymbol() []ISymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolContext)(nil)).Elem())
	var tst = make([]ISymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolContext)
		}
	}

	return tst
}

func (s *IterateContext) Symbol(i int) ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *IterateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.EnterIterate(s)
	}
}

func (s *IterateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.ExitIterate(s)
	}
}

func (s *IterateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pVisitor:
		return t.VisitIterate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pParser) Iterate() (localctx IIterateContext) {
	localctx = NewIterateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pParserRULE_iterate)
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
		p.SetState(17)
		p.Match(pParserT__0)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pParserT__0)|(1<<pParserR)|(1<<pParserL))) != 0) {
		{
			p.SetState(18)
			p.Symbol()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(pParserT__1)
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
	p.RuleIndex = pParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) R() antlr.TerminalNode {
	return s.GetToken(pParserR, 0)
}

func (s *AtomContext) L() antlr.TerminalNode {
	return s.GetToken(pParserL, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case pVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *pParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, pParserRULE_atom)
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
	p.SetState(25)
	_la = p.GetTokenStream().LA(1)

	if !(_la == pParserR || _la == pParserL) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
