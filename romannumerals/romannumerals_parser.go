// Code generated from romannumerals.g4 by ANTLR 4.9.3. DO NOT EDIT.

package romannumerals // romannumerals
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 88, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 57,
	10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 66, 10, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 74, 10, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 84, 10, 11, 3, 12, 3, 12, 3, 12,
	2, 3, 6, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 5, 3, 2, 7, 9,
	3, 2, 13, 15, 3, 2, 19, 21, 2, 95, 2, 24, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2,
	6, 33, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 58, 3, 2,
	2, 2, 14, 65, 3, 2, 2, 2, 16, 73, 3, 2, 2, 2, 18, 75, 3, 2, 2, 2, 20, 83,
	3, 2, 2, 2, 22, 85, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 3, 3, 2, 2, 2,
	26, 27, 5, 6, 4, 2, 27, 28, 5, 8, 5, 2, 28, 32, 3, 2, 2, 2, 29, 32, 5,
	6, 4, 2, 30, 32, 5, 8, 5, 2, 31, 26, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31,
	30, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 8, 4, 1, 2, 34, 35, 7, 3, 2,
	2, 35, 40, 3, 2, 2, 2, 36, 37, 12, 4, 2, 2, 37, 39, 7, 3, 2, 2, 38, 36,
	3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 7, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 5, 10, 6, 2, 44, 45, 5,
	14, 8, 2, 45, 49, 3, 2, 2, 2, 46, 49, 5, 10, 6, 2, 47, 49, 5, 14, 8, 2,
	48, 43, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 9, 3, 2,
	2, 2, 50, 57, 5, 12, 7, 2, 51, 57, 7, 4, 2, 2, 52, 57, 7, 5, 2, 2, 53,
	54, 7, 5, 2, 2, 54, 57, 5, 12, 7, 2, 55, 57, 7, 6, 2, 2, 56, 50, 3, 2,
	2, 2, 56, 51, 3, 2, 2, 2, 56, 52, 3, 2, 2, 2, 56, 53, 3, 2, 2, 2, 56, 55,
	3, 2, 2, 2, 57, 11, 3, 2, 2, 2, 58, 59, 9, 2, 2, 2, 59, 13, 3, 2, 2, 2,
	60, 61, 5, 16, 9, 2, 61, 62, 5, 20, 11, 2, 62, 66, 3, 2, 2, 2, 63, 66,
	5, 16, 9, 2, 64, 66, 5, 20, 11, 2, 65, 60, 3, 2, 2, 2, 65, 63, 3, 2, 2,
	2, 65, 64, 3, 2, 2, 2, 66, 15, 3, 2, 2, 2, 67, 74, 5, 18, 10, 2, 68, 74,
	7, 10, 2, 2, 69, 74, 7, 11, 2, 2, 70, 71, 7, 11, 2, 2, 71, 74, 5, 18, 10,
	2, 72, 74, 7, 12, 2, 2, 73, 67, 3, 2, 2, 2, 73, 68, 3, 2, 2, 2, 73, 69,
	3, 2, 2, 2, 73, 70, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 17, 3, 2, 2, 2,
	75, 76, 9, 3, 2, 2, 76, 19, 3, 2, 2, 2, 77, 84, 5, 22, 12, 2, 78, 84, 7,
	16, 2, 2, 79, 84, 7, 17, 2, 2, 80, 81, 7, 17, 2, 2, 81, 84, 5, 22, 12,
	2, 82, 84, 7, 18, 2, 2, 83, 77, 3, 2, 2, 2, 83, 78, 3, 2, 2, 2, 83, 79,
	3, 2, 2, 2, 83, 80, 3, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 21, 3, 2, 2, 2,
	85, 86, 9, 4, 2, 2, 86, 23, 3, 2, 2, 2, 9, 31, 40, 48, 56, 65, 73, 83,
}
var literalNames = []string{
	"", "'M'", "'CD'", "'D'", "'CM'", "'C'", "'CC'", "'CCC'", "'XL'", "'L'",
	"'XC'", "'X'", "'XX'", "'XXX'", "'IV'", "'V'", "'IX'", "'I'", "'II'", "'III'",
}
var symbolicNames = []string{
	"", "M", "CD", "D", "CM", "C", "CC", "CCC", "XL", "L", "XC", "X", "XX",
	"XXX", "IV", "V", "IX", "I", "II", "III", "WS",
}

var ruleNames = []string{
	"expression", "thousands", "thous_part", "hundreds", "hun_part", "hun_rep",
	"tens", "tens_part", "tens_rep", "ones", "ones_rep",
}

type romannumeralsParser struct {
	*antlr.BaseParser
}

// NewromannumeralsParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *romannumeralsParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewromannumeralsParser(input antlr.TokenStream) *romannumeralsParser {
	this := new(romannumeralsParser)
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
	this.GrammarFileName = "romannumerals.g4"

	return this
}

// romannumeralsParser tokens.
const (
	romannumeralsParserEOF = antlr.TokenEOF
	romannumeralsParserM   = 1
	romannumeralsParserCD  = 2
	romannumeralsParserD   = 3
	romannumeralsParserCM  = 4
	romannumeralsParserC   = 5
	romannumeralsParserCC  = 6
	romannumeralsParserCCC = 7
	romannumeralsParserXL  = 8
	romannumeralsParserL   = 9
	romannumeralsParserXC  = 10
	romannumeralsParserX   = 11
	romannumeralsParserXX  = 12
	romannumeralsParserXXX = 13
	romannumeralsParserIV  = 14
	romannumeralsParserV   = 15
	romannumeralsParserIX  = 16
	romannumeralsParserI   = 17
	romannumeralsParserII  = 18
	romannumeralsParserIII = 19
	romannumeralsParserWS  = 20
)

// romannumeralsParser rules.
const (
	romannumeralsParserRULE_expression = 0
	romannumeralsParserRULE_thousands  = 1
	romannumeralsParserRULE_thous_part = 2
	romannumeralsParserRULE_hundreds   = 3
	romannumeralsParserRULE_hun_part   = 4
	romannumeralsParserRULE_hun_rep    = 5
	romannumeralsParserRULE_tens       = 6
	romannumeralsParserRULE_tens_part  = 7
	romannumeralsParserRULE_tens_rep   = 8
	romannumeralsParserRULE_ones       = 9
	romannumeralsParserRULE_ones_rep   = 10
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Thousands() IThousandsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThousandsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThousandsContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *romannumeralsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, romannumeralsParserRULE_expression)

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
		p.SetState(22)
		p.Thousands()
	}

	return localctx
}

// IThousandsContext is an interface to support dynamic dispatch.
type IThousandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThousandsContext differentiates from other interfaces.
	IsThousandsContext()
}

type ThousandsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThousandsContext() *ThousandsContext {
	var p = new(ThousandsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_thousands
	return p
}

func (*ThousandsContext) IsThousandsContext() {}

func NewThousandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThousandsContext {
	var p = new(ThousandsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_thousands

	return p
}

func (s *ThousandsContext) GetParser() antlr.Parser { return s.parser }

func (s *ThousandsContext) Thous_part() IThous_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThous_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThous_partContext)
}

func (s *ThousandsContext) Hundreds() IHundredsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHundredsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHundredsContext)
}

func (s *ThousandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThousandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThousandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterThousands(s)
	}
}

func (s *ThousandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitThousands(s)
	}
}

func (p *romannumeralsParser) Thousands() (localctx IThousandsContext) {
	this := p
	_ = this

	localctx = NewThousandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, romannumeralsParserRULE_thousands)

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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.thous_part(0)
		}
		{
			p.SetState(25)
			p.Hundreds()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.thous_part(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Hundreds()
		}

	}

	return localctx
}

// IThous_partContext is an interface to support dynamic dispatch.
type IThous_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThous_partContext differentiates from other interfaces.
	IsThous_partContext()
}

type Thous_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThous_partContext() *Thous_partContext {
	var p = new(Thous_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_thous_part
	return p
}

func (*Thous_partContext) IsThous_partContext() {}

func NewThous_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Thous_partContext {
	var p = new(Thous_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_thous_part

	return p
}

func (s *Thous_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Thous_partContext) M() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserM, 0)
}

func (s *Thous_partContext) Thous_part() IThous_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThous_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThous_partContext)
}

func (s *Thous_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Thous_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Thous_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterThous_part(s)
	}
}

func (s *Thous_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitThous_part(s)
	}
}

func (p *romannumeralsParser) Thous_part() (localctx IThous_partContext) {
	return p.thous_part(0)
}

func (p *romannumeralsParser) thous_part(_p int) (localctx IThous_partContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewThous_partContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IThous_partContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, romannumeralsParserRULE_thous_part, _p)

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
		p.SetState(32)
		p.Match(romannumeralsParserM)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewThous_partContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, romannumeralsParserRULE_thous_part)
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(35)
				p.Match(romannumeralsParserM)
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IHundredsContext is an interface to support dynamic dispatch.
type IHundredsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHundredsContext differentiates from other interfaces.
	IsHundredsContext()
}

type HundredsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHundredsContext() *HundredsContext {
	var p = new(HundredsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_hundreds
	return p
}

func (*HundredsContext) IsHundredsContext() {}

func NewHundredsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HundredsContext {
	var p = new(HundredsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_hundreds

	return p
}

func (s *HundredsContext) GetParser() antlr.Parser { return s.parser }

func (s *HundredsContext) Hun_part() IHun_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHun_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHun_partContext)
}

func (s *HundredsContext) Tens() ITensContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITensContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITensContext)
}

func (s *HundredsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HundredsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HundredsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterHundreds(s)
	}
}

func (s *HundredsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitHundreds(s)
	}
}

func (p *romannumeralsParser) Hundreds() (localctx IHundredsContext) {
	this := p
	_ = this

	localctx = NewHundredsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, romannumeralsParserRULE_hundreds)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Hun_part()
		}
		{
			p.SetState(42)
			p.Tens()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Hun_part()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Tens()
		}

	}

	return localctx
}

// IHun_partContext is an interface to support dynamic dispatch.
type IHun_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHun_partContext differentiates from other interfaces.
	IsHun_partContext()
}

type Hun_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHun_partContext() *Hun_partContext {
	var p = new(Hun_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_hun_part
	return p
}

func (*Hun_partContext) IsHun_partContext() {}

func NewHun_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hun_partContext {
	var p = new(Hun_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_hun_part

	return p
}

func (s *Hun_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Hun_partContext) Hun_rep() IHun_repContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHun_repContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHun_repContext)
}

func (s *Hun_partContext) CD() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserCD, 0)
}

func (s *Hun_partContext) D() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserD, 0)
}

func (s *Hun_partContext) CM() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserCM, 0)
}

func (s *Hun_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hun_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hun_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterHun_part(s)
	}
}

func (s *Hun_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitHun_part(s)
	}
}

func (p *romannumeralsParser) Hun_part() (localctx IHun_partContext) {
	this := p
	_ = this

	localctx = NewHun_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, romannumeralsParserRULE_hun_part)

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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Hun_rep()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(romannumeralsParserCD)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(romannumeralsParserD)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Match(romannumeralsParserD)
		}
		{
			p.SetState(52)
			p.Hun_rep()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)
			p.Match(romannumeralsParserCM)
		}

	}

	return localctx
}

// IHun_repContext is an interface to support dynamic dispatch.
type IHun_repContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHun_repContext differentiates from other interfaces.
	IsHun_repContext()
}

type Hun_repContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHun_repContext() *Hun_repContext {
	var p = new(Hun_repContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_hun_rep
	return p
}

func (*Hun_repContext) IsHun_repContext() {}

func NewHun_repContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hun_repContext {
	var p = new(Hun_repContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_hun_rep

	return p
}

func (s *Hun_repContext) GetParser() antlr.Parser { return s.parser }

func (s *Hun_repContext) C() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserC, 0)
}

func (s *Hun_repContext) CC() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserCC, 0)
}

func (s *Hun_repContext) CCC() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserCCC, 0)
}

func (s *Hun_repContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hun_repContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hun_repContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterHun_rep(s)
	}
}

func (s *Hun_repContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitHun_rep(s)
	}
}

func (p *romannumeralsParser) Hun_rep() (localctx IHun_repContext) {
	this := p
	_ = this

	localctx = NewHun_repContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, romannumeralsParserRULE_hun_rep)
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
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<romannumeralsParserC)|(1<<romannumeralsParserCC)|(1<<romannumeralsParserCCC))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITensContext is an interface to support dynamic dispatch.
type ITensContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTensContext differentiates from other interfaces.
	IsTensContext()
}

type TensContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTensContext() *TensContext {
	var p = new(TensContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_tens
	return p
}

func (*TensContext) IsTensContext() {}

func NewTensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TensContext {
	var p = new(TensContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_tens

	return p
}

func (s *TensContext) GetParser() antlr.Parser { return s.parser }

func (s *TensContext) Tens_part() ITens_partContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITens_partContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITens_partContext)
}

func (s *TensContext) Ones() IOnesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnesContext)
}

func (s *TensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterTens(s)
	}
}

func (s *TensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitTens(s)
	}
}

func (p *romannumeralsParser) Tens() (localctx ITensContext) {
	this := p
	_ = this

	localctx = NewTensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, romannumeralsParserRULE_tens)

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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Tens_part()
		}
		{
			p.SetState(59)
			p.Ones()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Tens_part()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Ones()
		}

	}

	return localctx
}

// ITens_partContext is an interface to support dynamic dispatch.
type ITens_partContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTens_partContext differentiates from other interfaces.
	IsTens_partContext()
}

type Tens_partContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTens_partContext() *Tens_partContext {
	var p = new(Tens_partContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_tens_part
	return p
}

func (*Tens_partContext) IsTens_partContext() {}

func NewTens_partContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tens_partContext {
	var p = new(Tens_partContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_tens_part

	return p
}

func (s *Tens_partContext) GetParser() antlr.Parser { return s.parser }

func (s *Tens_partContext) Tens_rep() ITens_repContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITens_repContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITens_repContext)
}

func (s *Tens_partContext) XL() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserXL, 0)
}

func (s *Tens_partContext) L() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserL, 0)
}

func (s *Tens_partContext) XC() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserXC, 0)
}

func (s *Tens_partContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tens_partContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tens_partContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterTens_part(s)
	}
}

func (s *Tens_partContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitTens_part(s)
	}
}

func (p *romannumeralsParser) Tens_part() (localctx ITens_partContext) {
	this := p
	_ = this

	localctx = NewTens_partContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, romannumeralsParserRULE_tens_part)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Tens_rep()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(romannumeralsParserXL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(romannumeralsParserL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(romannumeralsParserL)
		}
		{
			p.SetState(69)
			p.Tens_rep()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.Match(romannumeralsParserXC)
		}

	}

	return localctx
}

// ITens_repContext is an interface to support dynamic dispatch.
type ITens_repContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTens_repContext differentiates from other interfaces.
	IsTens_repContext()
}

type Tens_repContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTens_repContext() *Tens_repContext {
	var p = new(Tens_repContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_tens_rep
	return p
}

func (*Tens_repContext) IsTens_repContext() {}

func NewTens_repContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tens_repContext {
	var p = new(Tens_repContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_tens_rep

	return p
}

func (s *Tens_repContext) GetParser() antlr.Parser { return s.parser }

func (s *Tens_repContext) X() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserX, 0)
}

func (s *Tens_repContext) XX() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserXX, 0)
}

func (s *Tens_repContext) XXX() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserXXX, 0)
}

func (s *Tens_repContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tens_repContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tens_repContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterTens_rep(s)
	}
}

func (s *Tens_repContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitTens_rep(s)
	}
}

func (p *romannumeralsParser) Tens_rep() (localctx ITens_repContext) {
	this := p
	_ = this

	localctx = NewTens_repContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, romannumeralsParserRULE_tens_rep)
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
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<romannumeralsParserX)|(1<<romannumeralsParserXX)|(1<<romannumeralsParserXXX))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOnesContext is an interface to support dynamic dispatch.
type IOnesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnesContext differentiates from other interfaces.
	IsOnesContext()
}

type OnesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnesContext() *OnesContext {
	var p = new(OnesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_ones
	return p
}

func (*OnesContext) IsOnesContext() {}

func NewOnesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnesContext {
	var p = new(OnesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_ones

	return p
}

func (s *OnesContext) GetParser() antlr.Parser { return s.parser }

func (s *OnesContext) Ones_rep() IOnes_repContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnes_repContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnes_repContext)
}

func (s *OnesContext) IV() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserIV, 0)
}

func (s *OnesContext) V() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserV, 0)
}

func (s *OnesContext) IX() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserIX, 0)
}

func (s *OnesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterOnes(s)
	}
}

func (s *OnesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitOnes(s)
	}
}

func (p *romannumeralsParser) Ones() (localctx IOnesContext) {
	this := p
	_ = this

	localctx = NewOnesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, romannumeralsParserRULE_ones)

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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Ones_rep()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(romannumeralsParserIV)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.Match(romannumeralsParserV)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.Match(romannumeralsParserV)
		}
		{
			p.SetState(79)
			p.Ones_rep()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)
			p.Match(romannumeralsParserIX)
		}

	}

	return localctx
}

// IOnes_repContext is an interface to support dynamic dispatch.
type IOnes_repContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnes_repContext differentiates from other interfaces.
	IsOnes_repContext()
}

type Ones_repContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnes_repContext() *Ones_repContext {
	var p = new(Ones_repContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = romannumeralsParserRULE_ones_rep
	return p
}

func (*Ones_repContext) IsOnes_repContext() {}

func NewOnes_repContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ones_repContext {
	var p = new(Ones_repContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = romannumeralsParserRULE_ones_rep

	return p
}

func (s *Ones_repContext) GetParser() antlr.Parser { return s.parser }

func (s *Ones_repContext) I() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserI, 0)
}

func (s *Ones_repContext) II() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserII, 0)
}

func (s *Ones_repContext) III() antlr.TerminalNode {
	return s.GetToken(romannumeralsParserIII, 0)
}

func (s *Ones_repContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ones_repContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ones_repContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.EnterOnes_rep(s)
	}
}

func (s *Ones_repContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(romannumeralsListener); ok {
		listenerT.ExitOnes_rep(s)
	}
}

func (p *romannumeralsParser) Ones_rep() (localctx IOnes_repContext) {
	this := p
	_ = this

	localctx = NewOnes_repContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, romannumeralsParserRULE_ones_rep)
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
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<romannumeralsParserI)|(1<<romannumeralsParserII)|(1<<romannumeralsParserIII))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *romannumeralsParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *Thous_partContext = nil
		if localctx != nil {
			t = localctx.(*Thous_partContext)
		}
		return p.Thous_part_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *romannumeralsParser) Thous_part_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
