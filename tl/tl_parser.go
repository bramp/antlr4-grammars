// Code generated from tl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tl // tl
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 26, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 16, 10, 2, 3, 2, 3, 2, 3, 2, 7, 2, 21, 10, 2, 12, 2, 14, 2, 24,
	11, 2, 3, 2, 2, 3, 2, 3, 2, 2, 3, 3, 2, 6, 7, 2, 30, 2, 15, 3, 2, 2, 2,
	4, 16, 8, 2, 1, 2, 5, 16, 7, 9, 2, 2, 6, 16, 7, 5, 2, 2, 7, 8, 7, 10, 2,
	2, 8, 16, 5, 2, 2, 6, 9, 10, 9, 2, 2, 2, 10, 16, 5, 2, 2, 4, 11, 12, 7,
	3, 2, 2, 12, 13, 5, 2, 2, 2, 13, 14, 7, 4, 2, 2, 14, 16, 3, 2, 2, 2, 15,
	4, 3, 2, 2, 2, 15, 5, 3, 2, 2, 2, 15, 6, 3, 2, 2, 2, 15, 7, 3, 2, 2, 2,
	15, 9, 3, 2, 2, 2, 15, 11, 3, 2, 2, 2, 16, 22, 3, 2, 2, 2, 17, 18, 12,
	5, 2, 2, 18, 19, 7, 8, 2, 2, 19, 21, 5, 2, 2, 6, 20, 17, 3, 2, 2, 2, 21,
	24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 3, 3, 2, 2,
	2, 24, 22, 3, 2, 2, 2, 4, 15, 22,
}
var literalNames = []string{
	"", "'('", "')'", "", "'G'", "'H'", "'\u2228'", "'\u22A5'", "'\u2310'",
}
var symbolicNames = []string{
	"", "", "", "ATOMIC", "TL_ALWAYS", "TL_WAS", "TL_OR", "TL_UPTACK", "TL_NOT",
	"WS",
}

var ruleNames = []string{
	"proposition",
}

type tlParser struct {
	*antlr.BaseParser
}

// NewtlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *tlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtlParser(input antlr.TokenStream) *tlParser {
	this := new(tlParser)
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
	this.GrammarFileName = "tl.g4"

	return this
}

// tlParser tokens.
const (
	tlParserEOF       = antlr.TokenEOF
	tlParserT__0      = 1
	tlParserT__1      = 2
	tlParserATOMIC    = 3
	tlParserTL_ALWAYS = 4
	tlParserTL_WAS    = 5
	tlParserTL_OR     = 6
	tlParserTL_UPTACK = 7
	tlParserTL_NOT    = 8
	tlParserWS        = 9
)

// tlParserRULE_proposition is the tlParser rule.
const tlParserRULE_proposition = 0

// IPropositionContext is an interface to support dynamic dispatch.
type IPropositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropositionContext differentiates from other interfaces.
	IsPropositionContext()
}

type PropositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropositionContext() *PropositionContext {
	var p = new(PropositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tlParserRULE_proposition
	return p
}

func (*PropositionContext) IsPropositionContext() {}

func NewPropositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropositionContext {
	var p = new(PropositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tlParserRULE_proposition

	return p
}

func (s *PropositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropositionContext) TL_UPTACK() antlr.TerminalNode {
	return s.GetToken(tlParserTL_UPTACK, 0)
}

func (s *PropositionContext) ATOMIC() antlr.TerminalNode {
	return s.GetToken(tlParserATOMIC, 0)
}

func (s *PropositionContext) TL_NOT() antlr.TerminalNode {
	return s.GetToken(tlParserTL_NOT, 0)
}

func (s *PropositionContext) AllProposition() []IPropositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropositionContext)(nil)).Elem())
	var tst = make([]IPropositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropositionContext)
		}
	}

	return tst
}

func (s *PropositionContext) Proposition(i int) IPropositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropositionContext)
}

func (s *PropositionContext) TL_ALWAYS() antlr.TerminalNode {
	return s.GetToken(tlParserTL_ALWAYS, 0)
}

func (s *PropositionContext) TL_WAS() antlr.TerminalNode {
	return s.GetToken(tlParserTL_WAS, 0)
}

func (s *PropositionContext) TL_OR() antlr.TerminalNode {
	return s.GetToken(tlParserTL_OR, 0)
}

func (s *PropositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.EnterProposition(s)
	}
}

func (s *PropositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tlListener); ok {
		listenerT.ExitProposition(s)
	}
}

func (p *tlParser) Proposition() (localctx IPropositionContext) {
	return p.proposition(0)
}

func (p *tlParser) proposition(_p int) (localctx IPropositionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPropositionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPropositionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, tlParserRULE_proposition, _p)
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:

	case 2:
		{
			p.SetState(3)
			p.Match(tlParserTL_UPTACK)
		}

	case 3:
		{
			p.SetState(4)
			p.Match(tlParserATOMIC)
		}

	case 4:
		{
			p.SetState(5)
			p.Match(tlParserTL_NOT)
		}
		{
			p.SetState(6)
			p.proposition(4)
		}

	case 5:
		{
			p.SetState(7)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tlParserTL_ALWAYS || _la == tlParserTL_WAS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(8)
			p.proposition(2)
		}

	case 6:
		{
			p.SetState(9)
			p.Match(tlParserT__0)
		}
		{
			p.SetState(10)
			p.proposition(0)
		}
		{
			p.SetState(11)
			p.Match(tlParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPropositionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tlParserRULE_proposition)
			p.SetState(15)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(16)
				p.Match(tlParserTL_OR)
			}
			{
				p.SetState(17)
				p.proposition(4)
			}

		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

func (p *tlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *PropositionContext = nil
		if localctx != nil {
			t = localctx.(*PropositionContext)
		}
		return p.Proposition_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tlParser) Proposition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
