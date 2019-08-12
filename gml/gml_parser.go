// Code generated from gml.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gml // gml
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14, 2,
	23, 3, 3, 3, 3, 6, 3, 28, 10, 3, 13, 3, 14, 3, 29, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 3, 6, 3,
	7, 5, 7, 47, 10, 7, 3, 7, 6, 7, 50, 10, 7, 13, 7, 14, 7, 51, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 2, 2, 2, 58, 2, 21, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 33, 3, 2, 2,
	2, 8, 41, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 46, 3, 2, 2, 2, 14, 53, 3,
	2, 2, 2, 16, 55, 3, 2, 2, 2, 18, 57, 3, 2, 2, 2, 20, 22, 5, 6, 4, 2, 21,
	20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2,
	2, 24, 3, 3, 2, 2, 2, 25, 27, 7, 3, 2, 2, 26, 28, 5, 6, 4, 2, 27, 26, 3,
	2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30,
	31, 3, 2, 2, 2, 31, 32, 7, 4, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 5, 10, 6,
	2, 34, 35, 5, 8, 5, 2, 35, 7, 3, 2, 2, 2, 36, 42, 5, 12, 7, 2, 37, 42,
	5, 14, 8, 2, 38, 42, 5, 18, 10, 2, 39, 42, 5, 16, 9, 2, 40, 42, 5, 4, 3,
	2, 41, 36, 3, 2, 2, 2, 41, 37, 3, 2, 2, 2, 41, 38, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 9, 3, 2, 2, 2, 43, 44, 7, 10, 2, 2,
	44, 11, 3, 2, 2, 2, 45, 47, 7, 7, 2, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 50, 7, 8, 2, 2, 49, 48, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 13, 3, 2, 2,
	2, 53, 54, 7, 6, 2, 2, 54, 15, 3, 2, 2, 2, 55, 56, 7, 10, 2, 2, 56, 17,
	3, 2, 2, 2, 57, 58, 7, 5, 2, 2, 58, 19, 3, 2, 2, 2, 7, 23, 29, 41, 46,
	51,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "STRINGLITERAL", "REAL", "SIGN", "DIGIT", "MANTISSA", "VALUE",
	"WS",
}

var ruleNames = []string{
	"graph", "list", "kv", "value", "key", "integer", "realnum", "str", "stringliteral",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type gmlParser struct {
	*antlr.BaseParser
}

func NewgmlParser(input antlr.TokenStream) *gmlParser {
	this := new(gmlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gml.g4"

	return this
}

// gmlParser tokens.
const (
	gmlParserEOF           = antlr.TokenEOF
	gmlParserT__0          = 1
	gmlParserT__1          = 2
	gmlParserSTRINGLITERAL = 3
	gmlParserREAL          = 4
	gmlParserSIGN          = 5
	gmlParserDIGIT         = 6
	gmlParserMANTISSA      = 7
	gmlParserVALUE         = 8
	gmlParserWS            = 9
)

// gmlParser rules.
const (
	gmlParserRULE_graph         = 0
	gmlParserRULE_list          = 1
	gmlParserRULE_kv            = 2
	gmlParserRULE_value         = 3
	gmlParserRULE_key           = 4
	gmlParserRULE_integer       = 5
	gmlParserRULE_realnum       = 6
	gmlParserRULE_str           = 7
	gmlParserRULE_stringliteral = 8
)

// IGraphContext is an interface to support dynamic dispatch.
type IGraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphContext differentiates from other interfaces.
	IsGraphContext()
}

type GraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphContext() *GraphContext {
	var p = new(GraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_graph
	return p
}

func (*GraphContext) IsGraphContext() {}

func NewGraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphContext {
	var p = new(GraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_graph

	return p
}

func (s *GraphContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphContext) AllKv() []IKvContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvContext)(nil)).Elem())
	var tst = make([]IKvContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvContext)
		}
	}

	return tst
}

func (s *GraphContext) Kv(i int) IKvContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvContext)
}

func (s *GraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterGraph(s)
	}
}

func (s *GraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitGraph(s)
	}
}

func (p *gmlParser) Graph() (localctx IGraphContext) {
	localctx = NewGraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gmlParserRULE_graph)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gmlParserVALUE {
		{
			p.SetState(18)
			p.Kv()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllKv() []IKvContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvContext)(nil)).Elem())
	var tst = make([]IKvContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvContext)
		}
	}

	return tst
}

func (s *ListContext) Kv(i int) IKvContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *gmlParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gmlParserRULE_list)
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
		p.Match(gmlParserT__0)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gmlParserVALUE {
		{
			p.SetState(24)
			p.Kv()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(29)
		p.Match(gmlParserT__1)
	}

	return localctx
}

// IKvContext is an interface to support dynamic dispatch.
type IKvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKvContext differentiates from other interfaces.
	IsKvContext()
}

type KvContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKvContext() *KvContext {
	var p = new(KvContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_kv
	return p
}

func (*KvContext) IsKvContext() {}

func NewKvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvContext {
	var p = new(KvContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_kv

	return p
}

func (s *KvContext) GetParser() antlr.Parser { return s.parser }

func (s *KvContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KvContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *KvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterKv(s)
	}
}

func (s *KvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitKv(s)
	}
}

func (p *gmlParser) Kv() (localctx IKvContext) {
	localctx = NewKvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gmlParserRULE_kv)

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
		p.SetState(31)
		p.Key()
	}
	{
		p.SetState(32)
		p.Value()
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
	p.RuleIndex = gmlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ValueContext) Realnum() IRealnumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealnumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealnumContext)
}

func (s *ValueContext) Stringliteral() IStringliteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringliteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringliteralContext)
}

func (s *ValueContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ValueContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *gmlParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gmlParserRULE_value)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gmlParserSIGN, gmlParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Integer()
		}

	case gmlParserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Realnum()
		}

	case gmlParserSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Stringliteral()
		}

	case gmlParserVALUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Str()
		}

	case gmlParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(38)
			p.List()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = gmlParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) VALUE() antlr.TerminalNode {
	return s.GetToken(gmlParserVALUE, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *gmlParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gmlParserRULE_key)

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
		p.Match(gmlParserVALUE)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) SIGN() antlr.TerminalNode {
	return s.GetToken(gmlParserSIGN, 0)
}

func (s *IntegerContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(gmlParserDIGIT)
}

func (s *IntegerContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(gmlParserDIGIT, i)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *gmlParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gmlParserRULE_integer)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gmlParserSIGN {
		{
			p.SetState(43)
			p.Match(gmlParserSIGN)
		}

	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gmlParserDIGIT {
		{
			p.SetState(46)
			p.Match(gmlParserDIGIT)
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRealnumContext is an interface to support dynamic dispatch.
type IRealnumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealnumContext differentiates from other interfaces.
	IsRealnumContext()
}

type RealnumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealnumContext() *RealnumContext {
	var p = new(RealnumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_realnum
	return p
}

func (*RealnumContext) IsRealnumContext() {}

func NewRealnumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealnumContext {
	var p = new(RealnumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_realnum

	return p
}

func (s *RealnumContext) GetParser() antlr.Parser { return s.parser }

func (s *RealnumContext) REAL() antlr.TerminalNode {
	return s.GetToken(gmlParserREAL, 0)
}

func (s *RealnumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealnumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealnumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterRealnum(s)
	}
}

func (s *RealnumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitRealnum(s)
	}
}

func (p *gmlParser) Realnum() (localctx IRealnumContext) {
	localctx = NewRealnumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gmlParserRULE_realnum)

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
		p.SetState(51)
		p.Match(gmlParserREAL)
	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) VALUE() antlr.TerminalNode {
	return s.GetToken(gmlParserVALUE, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitStr(s)
	}
}

func (p *gmlParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gmlParserRULE_str)

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
		p.SetState(53)
		p.Match(gmlParserVALUE)
	}

	return localctx
}

// IStringliteralContext is an interface to support dynamic dispatch.
type IStringliteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringliteralContext differentiates from other interfaces.
	IsStringliteralContext()
}

type StringliteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringliteralContext() *StringliteralContext {
	var p = new(StringliteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gmlParserRULE_stringliteral
	return p
}

func (*StringliteralContext) IsStringliteralContext() {}

func NewStringliteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringliteralContext {
	var p = new(StringliteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gmlParserRULE_stringliteral

	return p
}

func (s *StringliteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringliteralContext) STRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(gmlParserSTRINGLITERAL, 0)
}

func (s *StringliteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringliteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringliteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.EnterStringliteral(s)
	}
}

func (s *StringliteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gmlListener); ok {
		listenerT.ExitStringliteral(s)
	}
}

func (p *gmlParser) Stringliteral() (localctx IStringliteralContext) {
	localctx = NewStringliteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gmlParserRULE_stringliteral)

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
		p.SetState(55)
		p.Match(gmlParserSTRINGLITERAL)
	}

	return localctx
}
