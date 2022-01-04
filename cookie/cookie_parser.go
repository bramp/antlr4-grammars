// Code generated from cookie.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cookie // cookie
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2,
	25, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 32, 10, 4, 12, 4, 14, 4,
	35, 11, 4, 3, 5, 3, 5, 3, 5, 5, 5, 40, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 5, 8, 48, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 48, 2, 23, 3, 2, 2, 2, 4, 26, 3,
	2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12,
	43, 3, 2, 2, 2, 14, 47, 3, 2, 2, 2, 16, 49, 3, 2, 2, 2, 18, 51, 3, 2, 2,
	2, 20, 22, 5, 6, 4, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21,
	3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2,
	26, 27, 5, 10, 6, 2, 27, 5, 3, 2, 2, 2, 28, 33, 5, 8, 5, 2, 29, 30, 7,
	3, 2, 2, 30, 32, 5, 8, 5, 2, 31, 29, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33,
	31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 7, 3, 2, 2, 2, 35, 33, 3, 2, 2,
	2, 36, 39, 5, 10, 6, 2, 37, 38, 7, 4, 2, 2, 38, 40, 5, 12, 7, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 9, 3, 2, 2, 2, 41, 42, 5, 16, 9, 2,
	42, 11, 3, 2, 2, 2, 43, 44, 5, 14, 8, 2, 44, 13, 3, 2, 2, 2, 45, 48, 5,
	16, 9, 2, 46, 48, 5, 18, 10, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2,
	48, 15, 3, 2, 2, 2, 49, 50, 7, 6, 2, 2, 50, 17, 3, 2, 2, 2, 51, 52, 7,
	5, 2, 2, 52, 19, 3, 2, 2, 2, 6, 23, 33, 39, 47,
}
var literalNames = []string{
	"", "';'", "'='",
}
var symbolicNames = []string{
	"", "", "", "STRING", "TOKEN", "DIGIT", "WS",
}

var ruleNames = []string{
	"cookie", "name", "av_pairs", "av_pair", "attr", "value", "word", "token",
	"quoted_string",
}

type cookieParser struct {
	*antlr.BaseParser
}

// NewcookieParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *cookieParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcookieParser(input antlr.TokenStream) *cookieParser {
	this := new(cookieParser)
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
	this.GrammarFileName = "cookie.g4"

	return this
}

// cookieParser tokens.
const (
	cookieParserEOF    = antlr.TokenEOF
	cookieParserT__0   = 1
	cookieParserT__1   = 2
	cookieParserSTRING = 3
	cookieParserTOKEN  = 4
	cookieParserDIGIT  = 5
	cookieParserWS     = 6
)

// cookieParser rules.
const (
	cookieParserRULE_cookie        = 0
	cookieParserRULE_name          = 1
	cookieParserRULE_av_pairs      = 2
	cookieParserRULE_av_pair       = 3
	cookieParserRULE_attr          = 4
	cookieParserRULE_value         = 5
	cookieParserRULE_word          = 6
	cookieParserRULE_token         = 7
	cookieParserRULE_quoted_string = 8
)

// ICookieContext is an interface to support dynamic dispatch.
type ICookieContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCookieContext differentiates from other interfaces.
	IsCookieContext()
}

type CookieContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCookieContext() *CookieContext {
	var p = new(CookieContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_cookie
	return p
}

func (*CookieContext) IsCookieContext() {}

func NewCookieContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CookieContext {
	var p = new(CookieContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_cookie

	return p
}

func (s *CookieContext) GetParser() antlr.Parser { return s.parser }

func (s *CookieContext) AllAv_pairs() []IAv_pairsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAv_pairsContext)(nil)).Elem())
	var tst = make([]IAv_pairsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAv_pairsContext)
		}
	}

	return tst
}

func (s *CookieContext) Av_pairs(i int) IAv_pairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAv_pairsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAv_pairsContext)
}

func (s *CookieContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CookieContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CookieContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterCookie(s)
	}
}

func (s *CookieContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitCookie(s)
	}
}

func (p *cookieParser) Cookie() (localctx ICookieContext) {
	this := p
	_ = this

	localctx = NewCookieContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cookieParserRULE_cookie)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == cookieParserTOKEN {
		{
			p.SetState(18)
			p.Av_pairs()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = cookieParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *cookieParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, cookieParserRULE_name)

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
		p.SetState(24)
		p.Attr()
	}

	return localctx
}

// IAv_pairsContext is an interface to support dynamic dispatch.
type IAv_pairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAv_pairsContext differentiates from other interfaces.
	IsAv_pairsContext()
}

type Av_pairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAv_pairsContext() *Av_pairsContext {
	var p = new(Av_pairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_av_pairs
	return p
}

func (*Av_pairsContext) IsAv_pairsContext() {}

func NewAv_pairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Av_pairsContext {
	var p = new(Av_pairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_av_pairs

	return p
}

func (s *Av_pairsContext) GetParser() antlr.Parser { return s.parser }

func (s *Av_pairsContext) AllAv_pair() []IAv_pairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAv_pairContext)(nil)).Elem())
	var tst = make([]IAv_pairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAv_pairContext)
		}
	}

	return tst
}

func (s *Av_pairsContext) Av_pair(i int) IAv_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAv_pairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAv_pairContext)
}

func (s *Av_pairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Av_pairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Av_pairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterAv_pairs(s)
	}
}

func (s *Av_pairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitAv_pairs(s)
	}
}

func (p *cookieParser) Av_pairs() (localctx IAv_pairsContext) {
	this := p
	_ = this

	localctx = NewAv_pairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cookieParserRULE_av_pairs)
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
		p.SetState(26)
		p.Av_pair()
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == cookieParserT__0 {
		{
			p.SetState(27)
			p.Match(cookieParserT__0)
		}
		{
			p.SetState(28)
			p.Av_pair()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAv_pairContext is an interface to support dynamic dispatch.
type IAv_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAv_pairContext differentiates from other interfaces.
	IsAv_pairContext()
}

type Av_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAv_pairContext() *Av_pairContext {
	var p = new(Av_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_av_pair
	return p
}

func (*Av_pairContext) IsAv_pairContext() {}

func NewAv_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Av_pairContext {
	var p = new(Av_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_av_pair

	return p
}

func (s *Av_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Av_pairContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Av_pairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Av_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Av_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Av_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterAv_pair(s)
	}
}

func (s *Av_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitAv_pair(s)
	}
}

func (p *cookieParser) Av_pair() (localctx IAv_pairContext) {
	this := p
	_ = this

	localctx = NewAv_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, cookieParserRULE_av_pair)
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
		p.SetState(34)
		p.Attr()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cookieParserT__1 {
		{
			p.SetState(35)
			p.Match(cookieParserT__1)
		}
		{
			p.SetState(36)
			p.Value()
		}

	}

	return localctx
}

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *cookieParser) Attr() (localctx IAttrContext) {
	this := p
	_ = this

	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cookieParserRULE_attr)

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
		p.SetState(39)
		p.Token()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Word() IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *cookieParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cookieParserRULE_value)

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
		p.SetState(41)
		p.Word()
	}

	return localctx
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *WordContext) Quoted_string() IQuoted_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_stringContext)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitWord(s)
	}
}

func (p *cookieParser) Word() (localctx IWordContext) {
	this := p
	_ = this

	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, cookieParserRULE_word)

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

	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cookieParserTOKEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Token()
		}

	case cookieParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Quoted_string()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_token
	return p
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(cookieParserTOKEN, 0)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitToken(s)
	}
}

func (p *cookieParser) Token() (localctx ITokenContext) {
	this := p
	_ = this

	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, cookieParserRULE_token)

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
		p.SetState(47)
		p.Match(cookieParserTOKEN)
	}

	return localctx
}

// IQuoted_stringContext is an interface to support dynamic dispatch.
type IQuoted_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoted_stringContext differentiates from other interfaces.
	IsQuoted_stringContext()
}

type Quoted_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_stringContext() *Quoted_stringContext {
	var p = new(Quoted_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cookieParserRULE_quoted_string
	return p
}

func (*Quoted_stringContext) IsQuoted_stringContext() {}

func NewQuoted_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_stringContext {
	var p = new(Quoted_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cookieParserRULE_quoted_string

	return p
}

func (s *Quoted_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Quoted_stringContext) STRING() antlr.TerminalNode {
	return s.GetToken(cookieParserSTRING, 0)
}

func (s *Quoted_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quoted_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.EnterQuoted_string(s)
	}
}

func (s *Quoted_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cookieListener); ok {
		listenerT.ExitQuoted_string(s)
	}
}

func (p *cookieParser) Quoted_string() (localctx IQuoted_stringContext) {
	this := p
	_ = this

	localctx = NewQuoted_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, cookieParserRULE_quoted_string)

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
		p.SetState(49)
		p.Match(cookieParserSTRING)
	}

	return localctx
}
