// Generated from telephone.g4 by ANTLR 4.7.

package telephone // telephone
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 52, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 5, 2, 18, 10, 2, 3, 2, 5, 2, 21, 10, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 5, 3, 27, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 47, 2,
	17, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 28, 3, 2, 2, 2, 8, 33, 3, 2, 2, 2,
	10, 37, 3, 2, 2, 2, 12, 41, 3, 2, 2, 2, 14, 46, 3, 2, 2, 2, 16, 18, 7,
	3, 2, 2, 17, 16, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2, 19,
	21, 7, 4, 2, 2, 20, 19, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 22, 3, 2, 2,
	2, 22, 23, 5, 4, 3, 2, 23, 3, 3, 2, 2, 2, 24, 27, 5, 6, 4, 2, 25, 27, 5,
	14, 8, 2, 26, 24, 3, 2, 2, 2, 26, 25, 3, 2, 2, 2, 27, 5, 3, 2, 2, 2, 28,
	29, 7, 5, 2, 2, 29, 30, 5, 8, 5, 2, 30, 31, 5, 10, 6, 2, 31, 32, 5, 12,
	7, 2, 32, 7, 3, 2, 2, 2, 33, 34, 7, 7, 2, 2, 34, 35, 7, 7, 2, 2, 35, 36,
	7, 7, 2, 2, 36, 9, 3, 2, 2, 2, 37, 38, 7, 7, 2, 2, 38, 39, 7, 7, 2, 2,
	39, 40, 7, 7, 2, 2, 40, 11, 3, 2, 2, 2, 41, 42, 7, 7, 2, 2, 42, 43, 7,
	7, 2, 2, 43, 44, 7, 7, 2, 2, 44, 45, 7, 7, 2, 2, 45, 13, 3, 2, 2, 2, 46,
	47, 7, 6, 2, 2, 47, 48, 5, 8, 5, 2, 48, 49, 5, 10, 6, 2, 49, 50, 5, 12,
	7, 2, 50, 15, 3, 2, 2, 2, 5, 17, 20, 26,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+1'", "'+'", "'011'", "'010'",
}
var symbolicNames = []string{
	"", "", "", "", "", "DIGIT", "WS",
}

var ruleNames = []string{
	"number", "variation", "nanp", "areacode", "exchange", "subscriber", "japan",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type telephoneParser struct {
	*antlr.BaseParser
}

func NewtelephoneParser(input antlr.TokenStream) *telephoneParser {
	this := new(telephoneParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "telephone.g4"

	return this
}

// telephoneParser tokens.
const (
	telephoneParserEOF   = antlr.TokenEOF
	telephoneParserT__0  = 1
	telephoneParserT__1  = 2
	telephoneParserT__2  = 3
	telephoneParserT__3  = 4
	telephoneParserDIGIT = 5
	telephoneParserWS    = 6
)

// telephoneParser rules.
const (
	telephoneParserRULE_number     = 0
	telephoneParserRULE_variation  = 1
	telephoneParserRULE_nanp       = 2
	telephoneParserRULE_areacode   = 3
	telephoneParserRULE_exchange   = 4
	telephoneParserRULE_subscriber = 5
	telephoneParserRULE_japan      = 6
)

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
	p.RuleIndex = telephoneParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Variation() IVariationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariationContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *telephoneParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, telephoneParserRULE_number)
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

	if _la == telephoneParserT__0 {
		{
			p.SetState(14)
			p.Match(telephoneParserT__0)
		}

	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == telephoneParserT__1 {
		{
			p.SetState(17)
			p.Match(telephoneParserT__1)
		}

	}
	{
		p.SetState(20)
		p.Variation()
	}

	return localctx
}

// IVariationContext is an interface to support dynamic dispatch.
type IVariationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariationContext differentiates from other interfaces.
	IsVariationContext()
}

type VariationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariationContext() *VariationContext {
	var p = new(VariationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_variation
	return p
}

func (*VariationContext) IsVariationContext() {}

func NewVariationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariationContext {
	var p = new(VariationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_variation

	return p
}

func (s *VariationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariationContext) Nanp() INanpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INanpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INanpContext)
}

func (s *VariationContext) Japan() IJapanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJapanContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJapanContext)
}

func (s *VariationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterVariation(s)
	}
}

func (s *VariationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitVariation(s)
	}
}

func (p *telephoneParser) Variation() (localctx IVariationContext) {
	localctx = NewVariationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, telephoneParserRULE_variation)

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

	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case telephoneParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Nanp()
		}

	case telephoneParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Japan()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INanpContext is an interface to support dynamic dispatch.
type INanpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNanpContext differentiates from other interfaces.
	IsNanpContext()
}

type NanpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNanpContext() *NanpContext {
	var p = new(NanpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_nanp
	return p
}

func (*NanpContext) IsNanpContext() {}

func NewNanpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NanpContext {
	var p = new(NanpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_nanp

	return p
}

func (s *NanpContext) GetParser() antlr.Parser { return s.parser }

func (s *NanpContext) Areacode() IAreacodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAreacodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAreacodeContext)
}

func (s *NanpContext) Exchange() IExchangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExchangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExchangeContext)
}

func (s *NanpContext) Subscriber() ISubscriberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubscriberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubscriberContext)
}

func (s *NanpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NanpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NanpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterNanp(s)
	}
}

func (s *NanpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitNanp(s)
	}
}

func (p *telephoneParser) Nanp() (localctx INanpContext) {
	localctx = NewNanpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, telephoneParserRULE_nanp)

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
		p.Match(telephoneParserT__2)
	}
	{
		p.SetState(27)
		p.Areacode()
	}
	{
		p.SetState(28)
		p.Exchange()
	}
	{
		p.SetState(29)
		p.Subscriber()
	}

	return localctx
}

// IAreacodeContext is an interface to support dynamic dispatch.
type IAreacodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAreacodeContext differentiates from other interfaces.
	IsAreacodeContext()
}

type AreacodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAreacodeContext() *AreacodeContext {
	var p = new(AreacodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_areacode
	return p
}

func (*AreacodeContext) IsAreacodeContext() {}

func NewAreacodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AreacodeContext {
	var p = new(AreacodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_areacode

	return p
}

func (s *AreacodeContext) GetParser() antlr.Parser { return s.parser }

func (s *AreacodeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(telephoneParserDIGIT)
}

func (s *AreacodeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(telephoneParserDIGIT, i)
}

func (s *AreacodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AreacodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AreacodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterAreacode(s)
	}
}

func (s *AreacodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitAreacode(s)
	}
}

func (p *telephoneParser) Areacode() (localctx IAreacodeContext) {
	localctx = NewAreacodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, telephoneParserRULE_areacode)

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
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(32)
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(33)
		p.Match(telephoneParserDIGIT)
	}

	return localctx
}

// IExchangeContext is an interface to support dynamic dispatch.
type IExchangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExchangeContext differentiates from other interfaces.
	IsExchangeContext()
}

type ExchangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExchangeContext() *ExchangeContext {
	var p = new(ExchangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_exchange
	return p
}

func (*ExchangeContext) IsExchangeContext() {}

func NewExchangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExchangeContext {
	var p = new(ExchangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_exchange

	return p
}

func (s *ExchangeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExchangeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(telephoneParserDIGIT)
}

func (s *ExchangeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(telephoneParserDIGIT, i)
}

func (s *ExchangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExchangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExchangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterExchange(s)
	}
}

func (s *ExchangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitExchange(s)
	}
}

func (p *telephoneParser) Exchange() (localctx IExchangeContext) {
	localctx = NewExchangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, telephoneParserRULE_exchange)

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
		p.SetState(35)
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(36)
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(37)
		p.Match(telephoneParserDIGIT)
	}

	return localctx
}

// ISubscriberContext is an interface to support dynamic dispatch.
type ISubscriberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubscriberContext differentiates from other interfaces.
	IsSubscriberContext()
}

type SubscriberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscriberContext() *SubscriberContext {
	var p = new(SubscriberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_subscriber
	return p
}

func (*SubscriberContext) IsSubscriberContext() {}

func NewSubscriberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubscriberContext {
	var p = new(SubscriberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_subscriber

	return p
}

func (s *SubscriberContext) GetParser() antlr.Parser { return s.parser }

func (s *SubscriberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(telephoneParserDIGIT)
}

func (s *SubscriberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(telephoneParserDIGIT, i)
}

func (s *SubscriberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubscriberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterSubscriber(s)
	}
}

func (s *SubscriberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitSubscriber(s)
	}
}

func (p *telephoneParser) Subscriber() (localctx ISubscriberContext) {
	localctx = NewSubscriberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, telephoneParserRULE_subscriber)

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
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(40)
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(41)
		p.Match(telephoneParserDIGIT)
	}
	{
		p.SetState(42)
		p.Match(telephoneParserDIGIT)
	}

	return localctx
}

// IJapanContext is an interface to support dynamic dispatch.
type IJapanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJapanContext differentiates from other interfaces.
	IsJapanContext()
}

type JapanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJapanContext() *JapanContext {
	var p = new(JapanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = telephoneParserRULE_japan
	return p
}

func (*JapanContext) IsJapanContext() {}

func NewJapanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JapanContext {
	var p = new(JapanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = telephoneParserRULE_japan

	return p
}

func (s *JapanContext) GetParser() antlr.Parser { return s.parser }

func (s *JapanContext) Areacode() IAreacodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAreacodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAreacodeContext)
}

func (s *JapanContext) Exchange() IExchangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExchangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExchangeContext)
}

func (s *JapanContext) Subscriber() ISubscriberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubscriberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubscriberContext)
}

func (s *JapanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JapanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JapanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.EnterJapan(s)
	}
}

func (s *JapanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(telephoneListener); ok {
		listenerT.ExitJapan(s)
	}
}

func (p *telephoneParser) Japan() (localctx IJapanContext) {
	localctx = NewJapanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, telephoneParserRULE_japan)

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
		p.SetState(44)
		p.Match(telephoneParserT__3)
	}
	{
		p.SetState(45)
		p.Areacode()
	}
	{
		p.SetState(46)
		p.Exchange()
	}
	{
		p.SetState(47)
		p.Subscriber()
	}

	return localctx
}
