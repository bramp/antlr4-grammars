// Code generated from stellaris.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stellaris // stellaris
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 70, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14, 2,
	23, 3, 3, 6, 3, 27, 10, 3, 13, 3, 14, 3, 28, 3, 4, 3, 4, 6, 4, 33, 10,
	4, 13, 4, 14, 4, 34, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 41, 10, 5, 3, 6, 3,
	6, 3, 6, 5, 6, 46, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 52, 10, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 7, 9, 58, 10, 9, 12, 9, 14, 9, 61, 11, 9, 3, 9, 5, 9,
	64, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 2, 5, 3, 2, 3, 5, 3, 2, 6, 8, 3, 2, 11, 13, 2, 69, 2, 21, 3,
	2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 30, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 45,
	3, 2, 2, 2, 12, 47, 3, 2, 2, 2, 14, 53, 3, 2, 2, 2, 16, 55, 3, 2, 2, 2,
	18, 67, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 22, 23, 3,
	2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25,
	27, 5, 6, 4, 2, 26, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2,
	2, 28, 29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 32, 5, 8, 5, 2, 31, 33, 9,
	2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34,
	35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 5, 10, 6, 2, 37, 7, 3, 2, 2,
	2, 38, 41, 5, 18, 10, 2, 39, 41, 5, 12, 7, 2, 40, 38, 3, 2, 2, 2, 40, 39,
	3, 2, 2, 2, 41, 9, 3, 2, 2, 2, 42, 46, 5, 18, 10, 2, 43, 46, 5, 12, 7,
	2, 44, 46, 5, 16, 9, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44,
	3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 48, 5, 18, 10, 2, 48, 51, 5, 14, 8,
	2, 49, 52, 5, 12, 7, 2, 50, 52, 5, 18, 10, 2, 51, 49, 3, 2, 2, 2, 51, 50,
	3, 2, 2, 2, 52, 13, 3, 2, 2, 2, 53, 54, 9, 3, 2, 2, 54, 15, 3, 2, 2, 2,
	55, 63, 7, 9, 2, 2, 56, 58, 5, 4, 3, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3,
	2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 64, 3, 2, 2, 2, 61,
	59, 3, 2, 2, 2, 62, 64, 5, 18, 10, 2, 63, 59, 3, 2, 2, 2, 63, 62, 3, 2,
	2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 7, 10, 2, 2, 66, 17, 3, 2, 2, 2, 67,
	68, 9, 4, 2, 2, 68, 19, 3, 2, 2, 2, 10, 23, 28, 34, 40, 45, 51, 59, 63,
}
var literalNames = []string{
	"", "'='", "'>'", "'<'", "'.'", "'@'", "':'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "IDENTIFIER", "INTEGER", "STRING",
	"COMMENT", "SPACE", "NL",
}

var ruleNames = []string{
	"content", "expr", "keyval", "key", "val", "attrib", "accessor", "group",
	"id_",
}

type stellarisParser struct {
	*antlr.BaseParser
}

// NewstellarisParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *stellarisParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewstellarisParser(input antlr.TokenStream) *stellarisParser {
	this := new(stellarisParser)
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
	this.GrammarFileName = "stellaris.g4"

	return this
}

// stellarisParser tokens.
const (
	stellarisParserEOF        = antlr.TokenEOF
	stellarisParserT__0       = 1
	stellarisParserT__1       = 2
	stellarisParserT__2       = 3
	stellarisParserT__3       = 4
	stellarisParserT__4       = 5
	stellarisParserT__5       = 6
	stellarisParserT__6       = 7
	stellarisParserT__7       = 8
	stellarisParserIDENTIFIER = 9
	stellarisParserINTEGER    = 10
	stellarisParserSTRING     = 11
	stellarisParserCOMMENT    = 12
	stellarisParserSPACE      = 13
	stellarisParserNL         = 14
)

// stellarisParser rules.
const (
	stellarisParserRULE_content  = 0
	stellarisParserRULE_expr     = 1
	stellarisParserRULE_keyval   = 2
	stellarisParserRULE_key      = 3
	stellarisParserRULE_val      = 4
	stellarisParserRULE_attrib   = 5
	stellarisParserRULE_accessor = 6
	stellarisParserRULE_group    = 7
	stellarisParserRULE_id_      = 8
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ContentContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *stellarisParser) Content() (localctx IContentContext) {
	this := p
	_ = this

	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, stellarisParserRULE_content)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserIDENTIFIER)|(1<<stellarisParserINTEGER)|(1<<stellarisParserSTRING))) != 0) {
		{
			p.SetState(18)
			p.Expr()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllKeyval() []IKeyvalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyvalContext)(nil)).Elem())
	var tst = make([]IKeyvalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyvalContext)
		}
	}

	return tst
}

func (s *ExprContext) Keyval(i int) IKeyvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyvalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyvalContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *stellarisParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, stellarisParserRULE_expr)

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(23)
				p.Keyval()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IKeyvalContext is an interface to support dynamic dispatch.
type IKeyvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyvalContext differentiates from other interfaces.
	IsKeyvalContext()
}

type KeyvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyvalContext() *KeyvalContext {
	var p = new(KeyvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_keyval
	return p
}

func (*KeyvalContext) IsKeyvalContext() {}

func NewKeyvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvalContext {
	var p = new(KeyvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_keyval

	return p
}

func (s *KeyvalContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvalContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyvalContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *KeyvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterKeyval(s)
	}
}

func (s *KeyvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitKeyval(s)
	}
}

func (p *stellarisParser) Keyval() (localctx IKeyvalContext) {
	this := p
	_ = this

	localctx = NewKeyvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, stellarisParserRULE_keyval)
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
		p.SetState(28)
		p.Key()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserT__0)|(1<<stellarisParserT__1)|(1<<stellarisParserT__2))) != 0) {
		{
			p.SetState(29)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserT__0)|(1<<stellarisParserT__1)|(1<<stellarisParserT__2))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Val()
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
	p.RuleIndex = stellarisParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *KeyContext) Attrib() IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *stellarisParser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, stellarisParserRULE_key)

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

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Id_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Attrib()
		}

	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ValContext) Attrib() IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *ValContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *stellarisParser) Val() (localctx IValContext) {
	this := p
	_ = this

	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, stellarisParserRULE_val)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Id_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Attrib()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Group()
		}

	}

	return localctx
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) AllId_() []IId_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IId_Context)(nil)).Elem())
	var tst = make([]IId_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IId_Context)
		}
	}

	return tst
}

func (s *AttribContext) Id_(i int) IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *AttribContext) Accessor() IAccessorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessorContext)
}

func (s *AttribContext) Attrib() IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *stellarisParser) Attrib() (localctx IAttribContext) {
	this := p
	_ = this

	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, stellarisParserRULE_attrib)

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
		p.SetState(45)
		p.Id_()
	}
	{
		p.SetState(46)
		p.Accessor()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(47)
			p.Attrib()
		}

	case 2:
		{
			p.SetState(48)
			p.Id_()
		}

	}

	return localctx
}

// IAccessorContext is an interface to support dynamic dispatch.
type IAccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessorContext differentiates from other interfaces.
	IsAccessorContext()
}

type AccessorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorContext() *AccessorContext {
	var p = new(AccessorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_accessor
	return p
}

func (*AccessorContext) IsAccessorContext() {}

func NewAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorContext {
	var p = new(AccessorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_accessor

	return p
}

func (s *AccessorContext) GetParser() antlr.Parser { return s.parser }
func (s *AccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterAccessor(s)
	}
}

func (s *AccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitAccessor(s)
	}
}

func (p *stellarisParser) Accessor() (localctx IAccessorContext) {
	this := p
	_ = this

	localctx = NewAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, stellarisParserRULE_accessor)
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
		p.SetState(51)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserT__3)|(1<<stellarisParserT__4)|(1<<stellarisParserT__5))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *GroupContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GroupContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *stellarisParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, stellarisParserRULE_group)
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
		p.SetState(53)
		p.Match(stellarisParserT__6)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserIDENTIFIER)|(1<<stellarisParserINTEGER)|(1<<stellarisParserSTRING))) != 0 {
			{
				p.SetState(54)
				p.Expr()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		{
			p.SetState(60)
			p.Id_()
		}

	}
	{
		p.SetState(63)
		p.Match(stellarisParserT__7)
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = stellarisParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = stellarisParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(stellarisParserIDENTIFIER, 0)
}

func (s *Id_Context) STRING() antlr.TerminalNode {
	return s.GetToken(stellarisParserSTRING, 0)
}

func (s *Id_Context) INTEGER() antlr.TerminalNode {
	return s.GetToken(stellarisParserINTEGER, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stellarisListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *stellarisParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, stellarisParserRULE_id_)
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
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<stellarisParserIDENTIFIER)|(1<<stellarisParserINTEGER)|(1<<stellarisParserSTRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
