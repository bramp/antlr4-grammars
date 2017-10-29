// Generated from Abnf.g4 by ANTLR 4.7.

package abnf // Abnf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 83, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 7, 5, 40, 10, 5, 12, 5, 14, 5, 43, 11, 5, 3, 6, 3, 6,
	7, 6, 47, 10, 6, 12, 6, 14, 6, 50, 11, 6, 3, 7, 5, 7, 53, 10, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 5, 8, 59, 10, 8, 3, 8, 3, 8, 5, 8, 63, 10, 8, 5, 8, 65,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 73, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 2, 3, 3, 2, 3, 4, 2, 84, 2, 25, 3, 2, 2, 2, 4,
	30, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 44, 3, 2, 2,
	2, 12, 52, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2, 18, 74,
	3, 2, 2, 2, 20, 78, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2, 23, 22, 3, 2, 2, 2,
	24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 3,
	2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2, 30,
	31, 7, 13, 2, 2, 31, 32, 9, 2, 2, 2, 32, 33, 5, 6, 4, 2, 33, 5, 3, 2, 2,
	2, 34, 35, 5, 8, 5, 2, 35, 7, 3, 2, 2, 2, 36, 41, 5, 10, 6, 2, 37, 38,
	7, 5, 2, 2, 38, 40, 5, 10, 6, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2,
	41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 9, 3, 2, 2, 2, 43, 41, 3, 2,
	2, 2, 44, 48, 5, 12, 7, 2, 45, 47, 5, 12, 7, 2, 46, 45, 3, 2, 2, 2, 47,
	50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 11, 3, 2, 2,
	2, 50, 48, 3, 2, 2, 2, 51, 53, 5, 14, 8, 2, 52, 51, 3, 2, 2, 2, 52, 53,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 5, 16, 9, 2, 55, 13, 3, 2, 2, 2,
	56, 65, 7, 14, 2, 2, 57, 59, 7, 14, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 7, 6, 2, 2, 61, 63, 7, 14, 2, 2, 62,
	61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 56, 3, 2, 2,
	2, 64, 58, 3, 2, 2, 2, 65, 15, 3, 2, 2, 2, 66, 73, 7, 13, 2, 2, 67, 73,
	5, 18, 10, 2, 68, 73, 5, 20, 11, 2, 69, 73, 7, 17, 2, 2, 70, 73, 7, 11,
	2, 2, 71, 73, 7, 12, 2, 2, 72, 66, 3, 2, 2, 2, 72, 67, 3, 2, 2, 2, 72,
	68, 3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 71, 3, 2, 2,
	2, 73, 17, 3, 2, 2, 2, 74, 75, 7, 7, 2, 2, 75, 76, 5, 8, 5, 2, 76, 77,
	7, 8, 2, 2, 77, 19, 3, 2, 2, 2, 78, 79, 7, 9, 2, 2, 79, 80, 5, 8, 5, 2,
	80, 81, 7, 10, 2, 2, 81, 21, 3, 2, 2, 2, 10, 25, 41, 48, 52, 58, 62, 64,
	72,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'=/'", "'/'", "'*'", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "NumberValue", "ProseValue", "ID",
	"INT", "COMMENT", "WS", "STRING",
}

var ruleNames = []string{
	"rulelist", "rule_", "elements", "alternation", "concatenation", "repetition",
	"repeat", "element", "group", "option",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AbnfParser struct {
	*antlr.BaseParser
}

func NewAbnfParser(input antlr.TokenStream) *AbnfParser {
	this := new(AbnfParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Abnf.g4"

	return this
}

// AbnfParser tokens.
const (
	AbnfParserEOF         = antlr.TokenEOF
	AbnfParserT__0        = 1
	AbnfParserT__1        = 2
	AbnfParserT__2        = 3
	AbnfParserT__3        = 4
	AbnfParserT__4        = 5
	AbnfParserT__5        = 6
	AbnfParserT__6        = 7
	AbnfParserT__7        = 8
	AbnfParserNumberValue = 9
	AbnfParserProseValue  = 10
	AbnfParserID          = 11
	AbnfParserINT         = 12
	AbnfParserCOMMENT     = 13
	AbnfParserWS          = 14
	AbnfParserSTRING      = 15
)

// AbnfParser rules.
const (
	AbnfParserRULE_rulelist      = 0
	AbnfParserRULE_rule_         = 1
	AbnfParserRULE_elements      = 2
	AbnfParserRULE_alternation   = 3
	AbnfParserRULE_concatenation = 4
	AbnfParserRULE_repetition    = 5
	AbnfParserRULE_repeat        = 6
	AbnfParserRULE_element       = 7
	AbnfParserRULE_group         = 8
	AbnfParserRULE_option        = 9
)

// IRulelistContext is an interface to support dynamic dispatch.
type IRulelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulelistContext differentiates from other interfaces.
	IsRulelistContext()
}

type RulelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulelistContext() *RulelistContext {
	var p = new(RulelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_rulelist
	return p
}

func (*RulelistContext) IsRulelistContext() {}

func NewRulelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulelistContext {
	var p = new(RulelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_rulelist

	return p
}

func (s *RulelistContext) GetParser() antlr.Parser { return s.parser }

func (s *RulelistContext) EOF() antlr.TerminalNode {
	return s.GetToken(AbnfParserEOF, 0)
}

func (s *RulelistContext) AllRule_() []IRule_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRule_Context)(nil)).Elem())
	var tst = make([]IRule_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRule_Context)
		}
	}

	return tst
}

func (s *RulelistContext) Rule_(i int) IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *RulelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterRulelist(s)
	}
}

func (s *RulelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitRulelist(s)
	}
}

func (p *AbnfParser) Rulelist() (localctx IRulelistContext) {
	localctx = NewRulelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbnfParserRULE_rulelist)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AbnfParserID {
		{
			p.SetState(20)
			p.Rule_()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
		p.Match(AbnfParserEOF)
	}

	return localctx
}

// IRule_Context is an interface to support dynamic dispatch.
type IRule_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_Context differentiates from other interfaces.
	IsRule_Context()
}

type Rule_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_Context() *Rule_Context {
	var p = new(Rule_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_rule_
	return p
}

func (*Rule_Context) IsRule_Context() {}

func NewRule_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_Context {
	var p = new(Rule_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_rule_

	return p
}

func (s *Rule_Context) GetParser() antlr.Parser { return s.parser }

func (s *Rule_Context) ID() antlr.TerminalNode {
	return s.GetToken(AbnfParserID, 0)
}

func (s *Rule_Context) Elements() IElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *Rule_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterRule_(s)
	}
}

func (s *Rule_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitRule_(s)
	}
}

func (p *AbnfParser) Rule_() (localctx IRule_Context) {
	localctx = NewRule_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AbnfParserRULE_rule_)
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
		p.Match(AbnfParserID)
	}
	p.SetState(29)
	_la = p.GetTokenStream().LA(1)

	if !(_la == AbnfParserT__0 || _la == AbnfParserT__1) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(30)
		p.Elements()
	}

	return localctx
}

// IElementsContext is an interface to support dynamic dispatch.
type IElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementsContext differentiates from other interfaces.
	IsElementsContext()
}

type ElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementsContext() *ElementsContext {
	var p = new(ElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_elements
	return p
}

func (*ElementsContext) IsElementsContext() {}

func NewElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementsContext {
	var p = new(ElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_elements

	return p
}

func (s *ElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementsContext) Alternation() IAlternationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternationContext)
}

func (s *ElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterElements(s)
	}
}

func (s *ElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitElements(s)
	}
}

func (p *AbnfParser) Elements() (localctx IElementsContext) {
	localctx = NewElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AbnfParserRULE_elements)

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
		p.SetState(32)
		p.Alternation()
	}

	return localctx
}

// IAlternationContext is an interface to support dynamic dispatch.
type IAlternationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlternationContext differentiates from other interfaces.
	IsAlternationContext()
}

type AlternationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlternationContext() *AlternationContext {
	var p = new(AlternationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_alternation
	return p
}

func (*AlternationContext) IsAlternationContext() {}

func NewAlternationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlternationContext {
	var p = new(AlternationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_alternation

	return p
}

func (s *AlternationContext) GetParser() antlr.Parser { return s.parser }

func (s *AlternationContext) AllConcatenation() []IConcatenationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConcatenationContext)(nil)).Elem())
	var tst = make([]IConcatenationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConcatenationContext)
		}
	}

	return tst
}

func (s *AlternationContext) Concatenation(i int) IConcatenationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConcatenationContext)
}

func (s *AlternationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlternationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterAlternation(s)
	}
}

func (s *AlternationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitAlternation(s)
	}
}

func (p *AbnfParser) Alternation() (localctx IAlternationContext) {
	localctx = NewAlternationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AbnfParserRULE_alternation)
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
		p.Concatenation()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AbnfParserT__2 {
		{
			p.SetState(35)
			p.Match(AbnfParserT__2)
		}
		{
			p.SetState(36)
			p.Concatenation()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConcatenationContext is an interface to support dynamic dispatch.
type IConcatenationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatenationContext differentiates from other interfaces.
	IsConcatenationContext()
}

type ConcatenationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatenationContext() *ConcatenationContext {
	var p = new(ConcatenationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_concatenation
	return p
}

func (*ConcatenationContext) IsConcatenationContext() {}

func NewConcatenationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatenationContext {
	var p = new(ConcatenationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_concatenation

	return p
}

func (s *ConcatenationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatenationContext) AllRepetition() []IRepetitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRepetitionContext)(nil)).Elem())
	var tst = make([]IRepetitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRepetitionContext)
		}
	}

	return tst
}

func (s *ConcatenationContext) Repetition(i int) IRepetitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepetitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRepetitionContext)
}

func (s *ConcatenationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcatenationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterConcatenation(s)
	}
}

func (s *ConcatenationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitConcatenation(s)
	}
}

func (p *AbnfParser) Concatenation() (localctx IConcatenationContext) {
	localctx = NewConcatenationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AbnfParserRULE_concatenation)

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
	{
		p.SetState(42)
		p.Repetition()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Repetition()
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IRepetitionContext is an interface to support dynamic dispatch.
type IRepetitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepetitionContext differentiates from other interfaces.
	IsRepetitionContext()
}

type RepetitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepetitionContext() *RepetitionContext {
	var p = new(RepetitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_repetition
	return p
}

func (*RepetitionContext) IsRepetitionContext() {}

func NewRepetitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepetitionContext {
	var p = new(RepetitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_repetition

	return p
}

func (s *RepetitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RepetitionContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *RepetitionContext) Repeat() IRepeatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatContext)
}

func (s *RepetitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepetitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepetitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterRepetition(s)
	}
}

func (s *RepetitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitRepetition(s)
	}
}

func (p *AbnfParser) Repetition() (localctx IRepetitionContext) {
	localctx = NewRepetitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AbnfParserRULE_repetition)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbnfParserT__3 || _la == AbnfParserINT {
		{
			p.SetState(49)
			p.Repeat()
		}

	}
	{
		p.SetState(52)
		p.Element()
	}

	return localctx
}

// IRepeatContext is an interface to support dynamic dispatch.
type IRepeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatContext differentiates from other interfaces.
	IsRepeatContext()
}

type RepeatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatContext() *RepeatContext {
	var p = new(RepeatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_repeat
	return p
}

func (*RepeatContext) IsRepeatContext() {}

func NewRepeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatContext {
	var p = new(RepeatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_repeat

	return p
}

func (s *RepeatContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(AbnfParserINT)
}

func (s *RepeatContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(AbnfParserINT, i)
}

func (s *RepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterRepeat(s)
	}
}

func (s *RepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitRepeat(s)
	}
}

func (p *AbnfParser) Repeat() (localctx IRepeatContext) {
	localctx = NewRepeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AbnfParserRULE_repeat)
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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(AbnfParserINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbnfParserINT {
			{
				p.SetState(55)
				p.Match(AbnfParserINT)
			}

		}
		{
			p.SetState(58)
			p.Match(AbnfParserT__3)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbnfParserINT {
			{
				p.SetState(59)
				p.Match(AbnfParserINT)
			}

		}

	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) ID() antlr.TerminalNode {
	return s.GetToken(AbnfParserID, 0)
}

func (s *ElementContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ElementContext) Option() IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ElementContext) STRING() antlr.TerminalNode {
	return s.GetToken(AbnfParserSTRING, 0)
}

func (s *ElementContext) NumberValue() antlr.TerminalNode {
	return s.GetToken(AbnfParserNumberValue, 0)
}

func (s *ElementContext) ProseValue() antlr.TerminalNode {
	return s.GetToken(AbnfParserProseValue, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *AbnfParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AbnfParserRULE_element)

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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbnfParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(AbnfParserID)
		}

	case AbnfParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Group()
		}

	case AbnfParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Option()
		}

	case AbnfParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Match(AbnfParserSTRING)
		}

	case AbnfParserNumberValue:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.Match(AbnfParserNumberValue)
		}

	case AbnfParserProseValue:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.Match(AbnfParserProseValue)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = AbnfParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) Alternation() IAlternationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternationContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *AbnfParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AbnfParserRULE_group)

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
		p.SetState(72)
		p.Match(AbnfParserT__4)
	}
	{
		p.SetState(73)
		p.Alternation()
	}
	{
		p.SetState(74)
		p.Match(AbnfParserT__5)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbnfParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbnfParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) Alternation() IAlternationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternationContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbnfListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *AbnfParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AbnfParserRULE_option)

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
		p.SetState(76)
		p.Match(AbnfParserT__6)
	}
	{
		p.SetState(77)
		p.Alternation()
	}
	{
		p.SetState(78)
		p.Match(AbnfParserT__7)
	}

	return localctx
}
