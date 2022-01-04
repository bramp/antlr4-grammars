// Code generated from nanofuck.g4 by ANTLR 4.9.3. DO NOT EDIT.

package nanofuck // nanofuck
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 21, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 9, 10, 2, 3, 2, 5, 2, 12, 10, 2,
	3, 2, 3, 2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 2, 3, 2, 3,
	2, 2, 2, 2, 22, 2, 11, 3, 2, 2, 2, 4, 5, 8, 2, 1, 2, 5, 12, 7, 3, 2, 2,
	6, 8, 7, 4, 2, 2, 7, 9, 5, 2, 2, 2, 8, 7, 3, 2, 2, 2, 8, 9, 3, 2, 2, 2,
	9, 10, 3, 2, 2, 2, 10, 12, 7, 5, 2, 2, 11, 4, 3, 2, 2, 2, 11, 6, 3, 2,
	2, 2, 12, 17, 3, 2, 2, 2, 13, 14, 12, 3, 2, 2, 14, 16, 5, 2, 2, 4, 15,
	13, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2,
	2, 18, 3, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 5, 8, 11, 17,
}
var literalNames = []string{
	"", "'*'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "WS",
}

var ruleNames = []string{
	"exp",
}

type nanofuckParser struct {
	*antlr.BaseParser
}

// NewnanofuckParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *nanofuckParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnanofuckParser(input antlr.TokenStream) *nanofuckParser {
	this := new(nanofuckParser)
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
	this.GrammarFileName = "nanofuck.g4"

	return this
}

// nanofuckParser tokens.
const (
	nanofuckParserEOF  = antlr.TokenEOF
	nanofuckParserT__0 = 1
	nanofuckParserT__1 = 2
	nanofuckParserT__2 = 3
	nanofuckParserWS   = 4
)

// nanofuckParserRULE_exp is the nanofuckParser rule.
const nanofuckParserRULE_exp = 0

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = nanofuckParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = nanofuckParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(nanofuckListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(nanofuckListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *nanofuckParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *nanofuckParser) exp(_p int) (localctx IExpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, nanofuckParserRULE_exp, _p)
	var _la int

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
	p.SetState(9)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case nanofuckParserT__0:
		{
			p.SetState(3)
			p.Match(nanofuckParserT__0)
		}

	case nanofuckParserT__1:
		{
			p.SetState(4)
			p.Match(nanofuckParserT__1)
		}
		p.SetState(6)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == nanofuckParserT__0 || _la == nanofuckParserT__1 {
			{
				p.SetState(5)
				p.exp(0)
			}

		}
		{
			p.SetState(8)
			p.Match(nanofuckParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, nanofuckParserRULE_exp)
			p.SetState(11)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(12)
				p.exp(2)
			}

		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *nanofuckParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *nanofuckParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
