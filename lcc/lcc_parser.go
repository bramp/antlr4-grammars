// Generated from lcc.g4 by ANTLR 4.7.

package lcc // lcc
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 79, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 5, 2, 21, 10, 2, 3, 2, 5, 2, 24, 10, 2,
	3, 2, 5, 2, 27, 10, 2, 3, 2, 3, 2, 5, 2, 31, 10, 2, 3, 2, 3, 2, 5, 2, 35,
	10, 2, 3, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11, 5, 3, 6, 6, 6,
	56, 10, 6, 13, 6, 14, 6, 57, 3, 7, 3, 7, 3, 7, 7, 7, 63, 10, 7, 12, 7,
	14, 7, 66, 11, 7, 3, 8, 3, 8, 6, 8, 70, 10, 8, 13, 8, 14, 8, 71, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2,
	2, 2, 80, 2, 18, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 46,
	3, 2, 2, 2, 10, 55, 3, 2, 2, 2, 12, 59, 3, 2, 2, 2, 14, 67, 3, 2, 2, 2,
	16, 73, 3, 2, 2, 2, 18, 20, 5, 4, 3, 2, 19, 21, 7, 3, 2, 2, 20, 19, 3,
	2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 24, 5, 6, 4, 2, 23,
	22, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 27, 7, 3, 2,
	2, 26, 25, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30,
	5, 8, 5, 2, 29, 31, 7, 3, 2, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2,
	31, 34, 3, 2, 2, 2, 32, 33, 7, 4, 2, 2, 33, 35, 5, 12, 7, 2, 34, 32, 3,
	2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 37, 7, 3, 2, 2, 37,
	39, 5, 16, 9, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2,
	2, 2, 40, 41, 7, 2, 2, 3, 41, 3, 3, 2, 2, 2, 42, 43, 7, 6, 2, 2, 43, 5,
	3, 2, 2, 2, 44, 45, 7, 6, 2, 2, 45, 7, 3, 2, 2, 2, 46, 51, 5, 10, 6, 2,
	47, 48, 7, 4, 2, 2, 48, 50, 5, 10, 6, 2, 49, 47, 3, 2, 2, 2, 50, 53, 3,
	2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 9, 3, 2, 2, 2, 53,
	51, 3, 2, 2, 2, 54, 56, 7, 5, 2, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2,
	2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 11, 3, 2, 2, 2, 59, 64,
	5, 14, 8, 2, 60, 61, 7, 3, 2, 2, 61, 63, 5, 14, 8, 2, 62, 60, 3, 2, 2,
	2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 13,
	3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 69, 7, 6, 2, 2, 68, 70, 7, 5, 2, 2,
	69, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3,
	2, 2, 2, 72, 15, 3, 2, 2, 2, 73, 74, 7, 5, 2, 2, 74, 75, 7, 5, 2, 2, 75,
	76, 7, 5, 2, 2, 76, 77, 7, 5, 2, 2, 77, 17, 3, 2, 2, 2, 12, 20, 23, 26,
	30, 34, 38, 51, 57, 64, 71,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "' '", "'.'",
}
var symbolicNames = []string{
	"", "", "", "DIGIT", "LETTER", "WS",
}

var ruleNames = []string{
	"lcc", "topic", "subtopic", "subclasses", "subclass", "cutters", "cutter",
	"date",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type lccParser struct {
	*antlr.BaseParser
}

func NewlccParser(input antlr.TokenStream) *lccParser {
	this := new(lccParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "lcc.g4"

	return this
}

// lccParser tokens.
const (
	lccParserEOF    = antlr.TokenEOF
	lccParserT__0   = 1
	lccParserT__1   = 2
	lccParserDIGIT  = 3
	lccParserLETTER = 4
	lccParserWS     = 5
)

// lccParser rules.
const (
	lccParserRULE_lcc        = 0
	lccParserRULE_topic      = 1
	lccParserRULE_subtopic   = 2
	lccParserRULE_subclasses = 3
	lccParserRULE_subclass   = 4
	lccParserRULE_cutters    = 5
	lccParserRULE_cutter     = 6
	lccParserRULE_date       = 7
)

// ILccContext is an interface to support dynamic dispatch.
type ILccContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLccContext differentiates from other interfaces.
	IsLccContext()
}

type LccContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLccContext() *LccContext {
	var p = new(LccContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_lcc
	return p
}

func (*LccContext) IsLccContext() {}

func NewLccContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LccContext {
	var p = new(LccContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_lcc

	return p
}

func (s *LccContext) GetParser() antlr.Parser { return s.parser }

func (s *LccContext) Topic() ITopicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITopicContext)
}

func (s *LccContext) Subclasses() ISubclassesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubclassesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubclassesContext)
}

func (s *LccContext) EOF() antlr.TerminalNode {
	return s.GetToken(lccParserEOF, 0)
}

func (s *LccContext) Subtopic() ISubtopicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtopicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtopicContext)
}

func (s *LccContext) Cutters() ICuttersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICuttersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICuttersContext)
}

func (s *LccContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *LccContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LccContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LccContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterLcc(s)
	}
}

func (s *LccContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitLcc(s)
	}
}

func (p *lccParser) Lcc() (localctx ILccContext) {
	localctx = NewLccContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lccParserRULE_lcc)
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
		p.SetState(16)
		p.Topic()
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(17)
			p.Match(lccParserT__0)
		}

	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lccParserLETTER {
		{
			p.SetState(20)
			p.Subtopic()
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lccParserT__0 {
		{
			p.SetState(23)
			p.Match(lccParserT__0)
		}

	}
	{
		p.SetState(26)
		p.Subclasses()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(27)
			p.Match(lccParserT__0)
		}

	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lccParserT__1 {
		{
			p.SetState(30)
			p.Match(lccParserT__1)
		}
		{
			p.SetState(31)
			p.Cutters()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == lccParserT__0 {
		{
			p.SetState(34)
			p.Match(lccParserT__0)
		}
		{
			p.SetState(35)
			p.Date()
		}

	}
	{
		p.SetState(38)
		p.Match(lccParserEOF)
	}

	return localctx
}

// ITopicContext is an interface to support dynamic dispatch.
type ITopicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopicContext differentiates from other interfaces.
	IsTopicContext()
}

type TopicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopicContext() *TopicContext {
	var p = new(TopicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_topic
	return p
}

func (*TopicContext) IsTopicContext() {}

func NewTopicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopicContext {
	var p = new(TopicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_topic

	return p
}

func (s *TopicContext) GetParser() antlr.Parser { return s.parser }

func (s *TopicContext) LETTER() antlr.TerminalNode {
	return s.GetToken(lccParserLETTER, 0)
}

func (s *TopicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterTopic(s)
	}
}

func (s *TopicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitTopic(s)
	}
}

func (p *lccParser) Topic() (localctx ITopicContext) {
	localctx = NewTopicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lccParserRULE_topic)

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
		p.SetState(40)
		p.Match(lccParserLETTER)
	}

	return localctx
}

// ISubtopicContext is an interface to support dynamic dispatch.
type ISubtopicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtopicContext differentiates from other interfaces.
	IsSubtopicContext()
}

type SubtopicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtopicContext() *SubtopicContext {
	var p = new(SubtopicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_subtopic
	return p
}

func (*SubtopicContext) IsSubtopicContext() {}

func NewSubtopicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtopicContext {
	var p = new(SubtopicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_subtopic

	return p
}

func (s *SubtopicContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtopicContext) LETTER() antlr.TerminalNode {
	return s.GetToken(lccParserLETTER, 0)
}

func (s *SubtopicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtopicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtopicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterSubtopic(s)
	}
}

func (s *SubtopicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitSubtopic(s)
	}
}

func (p *lccParser) Subtopic() (localctx ISubtopicContext) {
	localctx = NewSubtopicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lccParserRULE_subtopic)

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
		p.SetState(42)
		p.Match(lccParserLETTER)
	}

	return localctx
}

// ISubclassesContext is an interface to support dynamic dispatch.
type ISubclassesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubclassesContext differentiates from other interfaces.
	IsSubclassesContext()
}

type SubclassesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubclassesContext() *SubclassesContext {
	var p = new(SubclassesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_subclasses
	return p
}

func (*SubclassesContext) IsSubclassesContext() {}

func NewSubclassesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubclassesContext {
	var p = new(SubclassesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_subclasses

	return p
}

func (s *SubclassesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubclassesContext) AllSubclass() []ISubclassContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubclassContext)(nil)).Elem())
	var tst = make([]ISubclassContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubclassContext)
		}
	}

	return tst
}

func (s *SubclassesContext) Subclass(i int) ISubclassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubclassContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubclassContext)
}

func (s *SubclassesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubclassesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubclassesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterSubclasses(s)
	}
}

func (s *SubclassesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitSubclasses(s)
	}
}

func (p *lccParser) Subclasses() (localctx ISubclassesContext) {
	localctx = NewSubclassesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lccParserRULE_subclasses)

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
		p.SetState(44)
		p.Subclass()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(45)
				p.Match(lccParserT__1)
			}
			{
				p.SetState(46)
				p.Subclass()
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ISubclassContext is an interface to support dynamic dispatch.
type ISubclassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubclassContext differentiates from other interfaces.
	IsSubclassContext()
}

type SubclassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubclassContext() *SubclassContext {
	var p = new(SubclassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_subclass
	return p
}

func (*SubclassContext) IsSubclassContext() {}

func NewSubclassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubclassContext {
	var p = new(SubclassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_subclass

	return p
}

func (s *SubclassContext) GetParser() antlr.Parser { return s.parser }

func (s *SubclassContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(lccParserDIGIT)
}

func (s *SubclassContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(lccParserDIGIT, i)
}

func (s *SubclassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubclassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubclassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterSubclass(s)
	}
}

func (s *SubclassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitSubclass(s)
	}
}

func (p *lccParser) Subclass() (localctx ISubclassContext) {
	localctx = NewSubclassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, lccParserRULE_subclass)
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == lccParserDIGIT {
		{
			p.SetState(52)
			p.Match(lccParserDIGIT)
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICuttersContext is an interface to support dynamic dispatch.
type ICuttersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCuttersContext differentiates from other interfaces.
	IsCuttersContext()
}

type CuttersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCuttersContext() *CuttersContext {
	var p = new(CuttersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_cutters
	return p
}

func (*CuttersContext) IsCuttersContext() {}

func NewCuttersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CuttersContext {
	var p = new(CuttersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_cutters

	return p
}

func (s *CuttersContext) GetParser() antlr.Parser { return s.parser }

func (s *CuttersContext) AllCutter() []ICutterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICutterContext)(nil)).Elem())
	var tst = make([]ICutterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICutterContext)
		}
	}

	return tst
}

func (s *CuttersContext) Cutter(i int) ICutterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICutterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICutterContext)
}

func (s *CuttersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CuttersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CuttersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterCutters(s)
	}
}

func (s *CuttersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitCutters(s)
	}
}

func (p *lccParser) Cutters() (localctx ICuttersContext) {
	localctx = NewCuttersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, lccParserRULE_cutters)

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
		p.SetState(57)
		p.Cutter()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(58)
				p.Match(lccParserT__0)
			}
			{
				p.SetState(59)
				p.Cutter()
			}

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// ICutterContext is an interface to support dynamic dispatch.
type ICutterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCutterContext differentiates from other interfaces.
	IsCutterContext()
}

type CutterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCutterContext() *CutterContext {
	var p = new(CutterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_cutter
	return p
}

func (*CutterContext) IsCutterContext() {}

func NewCutterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CutterContext {
	var p = new(CutterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_cutter

	return p
}

func (s *CutterContext) GetParser() antlr.Parser { return s.parser }

func (s *CutterContext) LETTER() antlr.TerminalNode {
	return s.GetToken(lccParserLETTER, 0)
}

func (s *CutterContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(lccParserDIGIT)
}

func (s *CutterContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(lccParserDIGIT, i)
}

func (s *CutterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CutterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CutterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterCutter(s)
	}
}

func (s *CutterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitCutter(s)
	}
}

func (p *lccParser) Cutter() (localctx ICutterContext) {
	localctx = NewCutterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, lccParserRULE_cutter)
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
		p.Match(lccParserLETTER)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == lccParserDIGIT {
		{
			p.SetState(66)
			p.Match(lccParserDIGIT)
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lccParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lccParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(lccParserDIGIT)
}

func (s *DateContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(lccParserDIGIT, i)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lccListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *lccParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, lccParserRULE_date)

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
		p.SetState(71)
		p.Match(lccParserDIGIT)
	}
	{
		p.SetState(72)
		p.Match(lccParserDIGIT)
	}
	{
		p.SetState(73)
		p.Match(lccParserDIGIT)
	}
	{
		p.SetState(74)
		p.Match(lccParserDIGIT)
	}

	return localctx
}
