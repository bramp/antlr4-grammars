// Code generated from JSON5.g4 by ANTLR 4.9.3. DO NOT EDIT.

package json5 // JSON5
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 76, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 5, 2, 18, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 5, 3, 32, 10, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 38, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 51, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7,
	7, 57, 10, 7, 12, 7, 14, 7, 60, 11, 7, 3, 7, 5, 7, 63, 10, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 69, 10, 7, 3, 8, 5, 8, 72, 10, 8, 3, 8, 3, 8, 3, 8,
	2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 5, 2, 11, 12, 14, 14, 16, 16, 3,
	2, 13, 14, 2, 80, 2, 17, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 39, 3, 2, 2,
	2, 8, 43, 3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 71, 3,
	2, 2, 2, 16, 18, 5, 10, 6, 2, 17, 16, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18,
	19, 3, 2, 2, 2, 19, 20, 7, 2, 2, 3, 20, 3, 3, 2, 2, 2, 21, 22, 7, 3, 2,
	2, 22, 27, 5, 6, 4, 2, 23, 24, 7, 4, 2, 2, 24, 26, 5, 6, 4, 2, 25, 23,
	3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2,
	28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32, 7, 4, 2, 2, 31, 30, 3,
	2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 34, 7, 5, 2, 2, 34,
	38, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 38, 7, 5, 2, 2, 37, 21, 3, 2, 2,
	2, 37, 35, 3, 2, 2, 2, 38, 5, 3, 2, 2, 2, 39, 40, 5, 8, 5, 2, 40, 41, 7,
	6, 2, 2, 41, 42, 5, 10, 6, 2, 42, 7, 3, 2, 2, 2, 43, 44, 9, 2, 2, 2, 44,
	9, 3, 2, 2, 2, 45, 51, 7, 12, 2, 2, 46, 51, 5, 14, 8, 2, 47, 51, 5, 4,
	3, 2, 48, 51, 5, 12, 7, 2, 49, 51, 7, 11, 2, 2, 50, 45, 3, 2, 2, 2, 50,
	46, 3, 2, 2, 2, 50, 47, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 49, 3, 2, 2,
	2, 51, 11, 3, 2, 2, 2, 52, 53, 7, 7, 2, 2, 53, 58, 5, 10, 6, 2, 54, 55,
	7, 4, 2, 2, 55, 57, 5, 10, 6, 2, 56, 54, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2,
	58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3,
	2, 2, 2, 61, 63, 7, 4, 2, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63,
	64, 3, 2, 2, 2, 64, 65, 7, 8, 2, 2, 65, 69, 3, 2, 2, 2, 66, 67, 7, 7, 2,
	2, 67, 69, 7, 8, 2, 2, 68, 52, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 13,
	3, 2, 2, 2, 70, 72, 7, 15, 2, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 73, 3, 2, 2, 2, 73, 74, 9, 3, 2, 2, 74, 15, 3, 2, 2, 2, 11, 17, 27,
	31, 37, 50, 58, 62, 68, 71,
}
var literalNames = []string{
	"", "'{'", "','", "'}'", "':'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT",
	"LITERAL", "STRING", "NUMBER", "NUMERIC_LITERAL", "SYMBOL", "IDENTIFIER",
	"WS",
}

var ruleNames = []string{
	"json5", "obj", "pair", "key", "value", "arr", "number",
}

type JSON5Parser struct {
	*antlr.BaseParser
}

// NewJSON5Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *JSON5Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJSON5Parser(input antlr.TokenStream) *JSON5Parser {
	this := new(JSON5Parser)
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
	this.GrammarFileName = "JSON5.g4"

	return this
}

// JSON5Parser tokens.
const (
	JSON5ParserEOF                 = antlr.TokenEOF
	JSON5ParserT__0                = 1
	JSON5ParserT__1                = 2
	JSON5ParserT__2                = 3
	JSON5ParserT__3                = 4
	JSON5ParserT__4                = 5
	JSON5ParserT__5                = 6
	JSON5ParserSINGLE_LINE_COMMENT = 7
	JSON5ParserMULTI_LINE_COMMENT  = 8
	JSON5ParserLITERAL             = 9
	JSON5ParserSTRING              = 10
	JSON5ParserNUMBER              = 11
	JSON5ParserNUMERIC_LITERAL     = 12
	JSON5ParserSYMBOL              = 13
	JSON5ParserIDENTIFIER          = 14
	JSON5ParserWS                  = 15
)

// JSON5Parser rules.
const (
	JSON5ParserRULE_json5  = 0
	JSON5ParserRULE_obj    = 1
	JSON5ParserRULE_pair   = 2
	JSON5ParserRULE_key    = 3
	JSON5ParserRULE_value  = 4
	JSON5ParserRULE_arr    = 5
	JSON5ParserRULE_number = 6
)

// IJson5Context is an interface to support dynamic dispatch.
type IJson5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJson5Context differentiates from other interfaces.
	IsJson5Context()
}

type Json5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson5Context() *Json5Context {
	var p = new(Json5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSON5ParserRULE_json5
	return p
}

func (*Json5Context) IsJson5Context() {}

func NewJson5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json5Context {
	var p = new(Json5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_json5

	return p
}

func (s *Json5Context) GetParser() antlr.Parser { return s.parser }

func (s *Json5Context) EOF() antlr.TerminalNode {
	return s.GetToken(JSON5ParserEOF, 0)
}

func (s *Json5Context) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Json5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterJson5(s)
	}
}

func (s *Json5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitJson5(s)
	}
}

func (p *JSON5Parser) Json5() (localctx IJson5Context) {
	this := p
	_ = this

	localctx = NewJson5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JSON5ParserRULE_json5)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSON5ParserT__0)|(1<<JSON5ParserT__4)|(1<<JSON5ParserLITERAL)|(1<<JSON5ParserSTRING)|(1<<JSON5ParserNUMBER)|(1<<JSON5ParserNUMERIC_LITERAL)|(1<<JSON5ParserSYMBOL))) != 0 {
		{
			p.SetState(14)
			p.Value()
		}

	}
	{
		p.SetState(17)
		p.Match(JSON5ParserEOF)
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSON5ParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *JSON5Parser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JSON5ParserRULE_obj)
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

	var _alt int

	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Match(JSON5ParserT__0)
		}
		{
			p.SetState(20)
			p.Pair()
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(21)
					p.Match(JSON5ParserT__1)
				}
				{
					p.SetState(22)
					p.Pair()
				}

			}
			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JSON5ParserT__1 {
			{
				p.SetState(28)
				p.Match(JSON5ParserT__1)
			}

		}
		{
			p.SetState(31)
			p.Match(JSON5ParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(JSON5ParserT__0)
		}
		{
			p.SetState(34)
			p.Match(JSON5ParserT__2)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSON5ParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *JSON5Parser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JSON5ParserRULE_pair)

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
		p.SetState(37)
		p.Key()
	}
	{
		p.SetState(38)
		p.Match(JSON5ParserT__3)
	}
	{
		p.SetState(39)
		p.Value()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSON5ParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSON5ParserSTRING, 0)
}

func (s *KeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(JSON5ParserIDENTIFIER, 0)
}

func (s *KeyContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(JSON5ParserLITERAL, 0)
}

func (s *KeyContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(JSON5ParserNUMERIC_LITERAL, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *JSON5Parser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JSON5ParserRULE_key)
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
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSON5ParserLITERAL)|(1<<JSON5ParserSTRING)|(1<<JSON5ParserNUMERIC_LITERAL)|(1<<JSON5ParserIDENTIFIER))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = JSON5ParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSON5ParserSTRING, 0)
}

func (s *ValueContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ValueContext) Obj() IObjContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Arr() IArrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *ValueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(JSON5ParserLITERAL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *JSON5Parser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JSON5ParserRULE_value)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSON5ParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Match(JSON5ParserSTRING)
		}

	case JSON5ParserNUMBER, JSON5ParserNUMERIC_LITERAL, JSON5ParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Number()
		}

	case JSON5ParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Obj()
		}

	case JSON5ParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.Arr()
		}

	case JSON5ParserLITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(47)
			p.Match(JSON5ParserLITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSON5ParserRULE_arr
	return p
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitArr(s)
	}
}

func (p *JSON5Parser) Arr() (localctx IArrContext) {
	this := p
	_ = this

	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JSON5ParserRULE_arr)
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

	var _alt int

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(JSON5ParserT__4)
		}
		{
			p.SetState(51)
			p.Value()
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(52)
					p.Match(JSON5ParserT__1)
				}
				{
					p.SetState(53)
					p.Value()
				}

			}
			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JSON5ParserT__1 {
			{
				p.SetState(59)
				p.Match(JSON5ParserT__1)
			}

		}
		{
			p.SetState(62)
			p.Match(JSON5ParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(JSON5ParserT__4)
		}
		{
			p.SetState(65)
			p.Match(JSON5ParserT__5)
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
	p.RuleIndex = JSON5ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSON5ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(JSON5ParserNUMERIC_LITERAL, 0)
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JSON5ParserNUMBER, 0)
}

func (s *NumberContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(JSON5ParserSYMBOL, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSON5Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *JSON5Parser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JSON5ParserRULE_number)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JSON5ParserSYMBOL {
		{
			p.SetState(68)
			p.Match(JSON5ParserSYMBOL)
		}

	}
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JSON5ParserNUMBER || _la == JSON5ParserNUMERIC_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
