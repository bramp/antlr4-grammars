// Code generated from PMMN.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pmmn // PMMN
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 68, 10, 6, 12, 6, 14,
	6, 71, 11, 6, 3, 6, 2, 4, 2, 10, 7, 2, 4, 6, 8, 10, 2, 2, 2, 73, 2, 12,
	3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10,
	62, 3, 2, 2, 2, 12, 13, 8, 2, 1, 2, 13, 14, 5, 6, 4, 2, 14, 19, 3, 2, 2,
	2, 15, 16, 12, 3, 2, 2, 16, 18, 5, 6, 4, 2, 17, 15, 3, 2, 2, 2, 18, 21,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 3, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 22, 23, 7, 3, 2, 2, 23, 24, 5, 2, 2, 2, 24, 25, 7,
	4, 2, 2, 25, 5, 3, 2, 2, 2, 26, 27, 7, 5, 2, 2, 27, 28, 7, 6, 2, 2, 28,
	29, 5, 10, 6, 2, 29, 30, 7, 7, 2, 2, 30, 31, 7, 8, 2, 2, 31, 56, 3, 2,
	2, 2, 32, 33, 5, 8, 5, 2, 33, 34, 7, 8, 2, 2, 34, 56, 3, 2, 2, 2, 35, 36,
	7, 9, 2, 2, 36, 37, 7, 6, 2, 2, 37, 38, 5, 8, 5, 2, 38, 39, 7, 7, 2, 2,
	39, 40, 5, 4, 3, 2, 40, 56, 3, 2, 2, 2, 41, 42, 7, 9, 2, 2, 42, 43, 7,
	6, 2, 2, 43, 44, 5, 8, 5, 2, 44, 45, 7, 7, 2, 2, 45, 46, 5, 4, 3, 2, 46,
	47, 7, 10, 2, 2, 47, 48, 5, 4, 3, 2, 48, 56, 3, 2, 2, 2, 49, 50, 7, 11,
	2, 2, 50, 51, 7, 6, 2, 2, 51, 52, 5, 8, 5, 2, 52, 53, 7, 7, 2, 2, 53, 54,
	5, 4, 3, 2, 54, 56, 3, 2, 2, 2, 55, 26, 3, 2, 2, 2, 55, 32, 3, 2, 2, 2,
	55, 35, 3, 2, 2, 2, 55, 41, 3, 2, 2, 2, 55, 49, 3, 2, 2, 2, 56, 7, 3, 2,
	2, 2, 57, 58, 7, 12, 2, 2, 58, 59, 7, 6, 2, 2, 59, 60, 5, 10, 6, 2, 60,
	61, 7, 7, 2, 2, 61, 9, 3, 2, 2, 2, 62, 63, 8, 6, 1, 2, 63, 64, 7, 13, 2,
	2, 64, 69, 3, 2, 2, 2, 65, 66, 12, 3, 2, 2, 66, 68, 7, 13, 2, 2, 67, 65,
	3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2,
	70, 11, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 5, 19, 55, 69,
}
var literalNames = []string{
	"", "'{'", "'}'", "'inc'", "'('", "')'", "';'", "'if'", "'else'", "'while'",
	"'dec'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "DIGIT", "WS",
}

var ruleNames = []string{
	"commandlist", "block", "command", "test", "counter",
}

type PMMNParser struct {
	*antlr.BaseParser
}

// NewPMMNParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PMMNParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPMMNParser(input antlr.TokenStream) *PMMNParser {
	this := new(PMMNParser)
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
	this.GrammarFileName = "PMMN.g4"

	return this
}

// PMMNParser tokens.
const (
	PMMNParserEOF   = antlr.TokenEOF
	PMMNParserT__0  = 1
	PMMNParserT__1  = 2
	PMMNParserT__2  = 3
	PMMNParserT__3  = 4
	PMMNParserT__4  = 5
	PMMNParserT__5  = 6
	PMMNParserT__6  = 7
	PMMNParserT__7  = 8
	PMMNParserT__8  = 9
	PMMNParserT__9  = 10
	PMMNParserDIGIT = 11
	PMMNParserWS    = 12
)

// PMMNParser rules.
const (
	PMMNParserRULE_commandlist = 0
	PMMNParserRULE_block       = 1
	PMMNParserRULE_command     = 2
	PMMNParserRULE_test        = 3
	PMMNParserRULE_counter     = 4
)

// ICommandlistContext is an interface to support dynamic dispatch.
type ICommandlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandlistContext differentiates from other interfaces.
	IsCommandlistContext()
}

type CommandlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandlistContext() *CommandlistContext {
	var p = new(CommandlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMMNParserRULE_commandlist
	return p
}

func (*CommandlistContext) IsCommandlistContext() {}

func NewCommandlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandlistContext {
	var p = new(CommandlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMMNParserRULE_commandlist

	return p
}

func (s *CommandlistContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandlistContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandlistContext) Commandlist() ICommandlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandlistContext)
}

func (s *CommandlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.EnterCommandlist(s)
	}
}

func (s *CommandlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.ExitCommandlist(s)
	}
}

func (p *PMMNParser) Commandlist() (localctx ICommandlistContext) {
	return p.commandlist(0)
}

func (p *PMMNParser) commandlist(_p int) (localctx ICommandlistContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCommandlistContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICommandlistContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, PMMNParserRULE_commandlist, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(11)
		p.Command()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCommandlistContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PMMNParserRULE_commandlist)
			p.SetState(13)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(14)
				p.Command()
			}

		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMMNParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMMNParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Commandlist() ICommandlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandlistContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *PMMNParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PMMNParserRULE_block)

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
		p.SetState(20)
		p.Match(PMMNParserT__0)
	}
	{
		p.SetState(21)
		p.commandlist(0)
	}
	{
		p.SetState(22)
		p.Match(PMMNParserT__1)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMMNParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMMNParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Counter() ICounterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICounterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICounterContext)
}

func (s *CommandContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *CommandContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *CommandContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *PMMNParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PMMNParserRULE_command)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(PMMNParserT__2)
		}
		{
			p.SetState(25)
			p.Match(PMMNParserT__3)
		}
		{
			p.SetState(26)
			p.counter(0)
		}
		{
			p.SetState(27)
			p.Match(PMMNParserT__4)
		}
		{
			p.SetState(28)
			p.Match(PMMNParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Test()
		}
		{
			p.SetState(31)
			p.Match(PMMNParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(PMMNParserT__6)
		}
		{
			p.SetState(34)
			p.Match(PMMNParserT__3)
		}
		{
			p.SetState(35)
			p.Test()
		}
		{
			p.SetState(36)
			p.Match(PMMNParserT__4)
		}
		{
			p.SetState(37)
			p.Block()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Match(PMMNParserT__6)
		}
		{
			p.SetState(40)
			p.Match(PMMNParserT__3)
		}
		{
			p.SetState(41)
			p.Test()
		}
		{
			p.SetState(42)
			p.Match(PMMNParserT__4)
		}
		{
			p.SetState(43)
			p.Block()
		}
		{
			p.SetState(44)
			p.Match(PMMNParserT__7)
		}
		{
			p.SetState(45)
			p.Block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(47)
			p.Match(PMMNParserT__8)
		}
		{
			p.SetState(48)
			p.Match(PMMNParserT__3)
		}
		{
			p.SetState(49)
			p.Test()
		}
		{
			p.SetState(50)
			p.Match(PMMNParserT__4)
		}
		{
			p.SetState(51)
			p.Block()
		}

	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMMNParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMMNParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) Counter() ICounterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICounterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICounterContext)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.ExitTest(s)
	}
}

func (p *PMMNParser) Test() (localctx ITestContext) {
	this := p
	_ = this

	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PMMNParserRULE_test)

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
		p.Match(PMMNParserT__9)
	}
	{
		p.SetState(56)
		p.Match(PMMNParserT__3)
	}
	{
		p.SetState(57)
		p.counter(0)
	}
	{
		p.SetState(58)
		p.Match(PMMNParserT__4)
	}

	return localctx
}

// ICounterContext is an interface to support dynamic dispatch.
type ICounterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCounterContext differentiates from other interfaces.
	IsCounterContext()
}

type CounterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCounterContext() *CounterContext {
	var p = new(CounterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMMNParserRULE_counter
	return p
}

func (*CounterContext) IsCounterContext() {}

func NewCounterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CounterContext {
	var p = new(CounterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMMNParserRULE_counter

	return p
}

func (s *CounterContext) GetParser() antlr.Parser { return s.parser }

func (s *CounterContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(PMMNParserDIGIT, 0)
}

func (s *CounterContext) Counter() ICounterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICounterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICounterContext)
}

func (s *CounterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CounterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CounterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.EnterCounter(s)
	}
}

func (s *CounterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMMNListener); ok {
		listenerT.ExitCounter(s)
	}
}

func (p *PMMNParser) Counter() (localctx ICounterContext) {
	return p.counter(0)
}

func (p *PMMNParser) counter(_p int) (localctx ICounterContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCounterContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICounterContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, PMMNParserRULE_counter, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
		p.SetState(61)
		p.Match(PMMNParserDIGIT)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCounterContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PMMNParserRULE_counter)
			p.SetState(63)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(64)
				p.Match(PMMNParserDIGIT)
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *PMMNParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *CommandlistContext = nil
		if localctx != nil {
			t = localctx.(*CommandlistContext)
		}
		return p.Commandlist_Sempred(t, predIndex)

	case 4:
		var t *CounterContext = nil
		if localctx != nil {
			t = localctx.(*CounterContext)
		}
		return p.Counter_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PMMNParser) Commandlist_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PMMNParser) Counter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
