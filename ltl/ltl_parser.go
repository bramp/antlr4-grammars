// Code generated from ltl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ltl // ltl
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 30, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 25,
	10, 2, 12, 2, 14, 2, 28, 11, 2, 3, 2, 2, 3, 2, 3, 2, 2, 5, 3, 2, 9, 11,
	3, 2, 14, 16, 4, 2, 8, 8, 12, 13, 2, 35, 2, 16, 3, 2, 2, 2, 4, 5, 8, 2,
	1, 2, 5, 17, 7, 3, 2, 2, 6, 17, 7, 4, 2, 2, 7, 17, 7, 7, 2, 2, 8, 9, 7,
	5, 2, 2, 9, 10, 5, 2, 2, 2, 10, 11, 7, 6, 2, 2, 11, 17, 3, 2, 2, 2, 12,
	13, 7, 17, 2, 2, 13, 17, 5, 2, 2, 5, 14, 15, 9, 2, 2, 2, 15, 17, 5, 2,
	2, 4, 16, 4, 3, 2, 2, 2, 16, 6, 3, 2, 2, 2, 16, 7, 3, 2, 2, 2, 16, 8, 3,
	2, 2, 2, 16, 12, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 17, 26, 3, 2, 2, 2, 18,
	19, 12, 6, 2, 2, 19, 20, 9, 3, 2, 2, 20, 25, 5, 2, 2, 7, 21, 22, 12, 3,
	2, 2, 22, 23, 9, 4, 2, 2, 23, 25, 5, 2, 2, 4, 24, 18, 3, 2, 2, 2, 24, 21,
	3, 2, 2, 2, 25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2,
	27, 3, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 5, 16, 24, 26,
}
var literalNames = []string{
	"", "'true'", "'false'", "'('", "')'", "", "'U'", "'G'", "'F'", "'X'",
	"'W'", "'R'", "'\u2192'", "'\u2227'", "'\u2228'", "'\u2310'",
}
var symbolicNames = []string{
	"", "", "", "", "", "ATOMIC", "LTL_UNTIL", "LTL_GLOBALLY", "LTL_FINALLY",
	"LTL_NEXT", "LTL_WEAK", "LTL_RELEASE", "LTL_RIGHTWARDS_SINGLE_ARROW", "LTL_AND",
	"LTL_OR", "LTL_NOT", "WS",
}

var ruleNames = []string{
	"proposition",
}

type ltlParser struct {
	*antlr.BaseParser
}

// NewltlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ltlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewltlParser(input antlr.TokenStream) *ltlParser {
	this := new(ltlParser)
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
	this.GrammarFileName = "ltl.g4"

	return this
}

// ltlParser tokens.
const (
	ltlParserEOF                         = antlr.TokenEOF
	ltlParserT__0                        = 1
	ltlParserT__1                        = 2
	ltlParserT__2                        = 3
	ltlParserT__3                        = 4
	ltlParserATOMIC                      = 5
	ltlParserLTL_UNTIL                   = 6
	ltlParserLTL_GLOBALLY                = 7
	ltlParserLTL_FINALLY                 = 8
	ltlParserLTL_NEXT                    = 9
	ltlParserLTL_WEAK                    = 10
	ltlParserLTL_RELEASE                 = 11
	ltlParserLTL_RIGHTWARDS_SINGLE_ARROW = 12
	ltlParserLTL_AND                     = 13
	ltlParserLTL_OR                      = 14
	ltlParserLTL_NOT                     = 15
	ltlParserWS                          = 16
)

// ltlParserRULE_proposition is the ltlParser rule.
const ltlParserRULE_proposition = 0

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
	p.RuleIndex = ltlParserRULE_proposition
	return p
}

func (*PropositionContext) IsPropositionContext() {}

func NewPropositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropositionContext {
	var p = new(PropositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ltlParserRULE_proposition

	return p
}

func (s *PropositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropositionContext) ATOMIC() antlr.TerminalNode {
	return s.GetToken(ltlParserATOMIC, 0)
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

func (s *PropositionContext) LTL_NOT() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_NOT, 0)
}

func (s *PropositionContext) LTL_GLOBALLY() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_GLOBALLY, 0)
}

func (s *PropositionContext) LTL_FINALLY() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_FINALLY, 0)
}

func (s *PropositionContext) LTL_NEXT() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_NEXT, 0)
}

func (s *PropositionContext) LTL_AND() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_AND, 0)
}

func (s *PropositionContext) LTL_OR() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_OR, 0)
}

func (s *PropositionContext) LTL_RIGHTWARDS_SINGLE_ARROW() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_RIGHTWARDS_SINGLE_ARROW, 0)
}

func (s *PropositionContext) LTL_UNTIL() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_UNTIL, 0)
}

func (s *PropositionContext) LTL_WEAK() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_WEAK, 0)
}

func (s *PropositionContext) LTL_RELEASE() antlr.TerminalNode {
	return s.GetToken(ltlParserLTL_RELEASE, 0)
}

func (s *PropositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ltlListener); ok {
		listenerT.EnterProposition(s)
	}
}

func (s *PropositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ltlListener); ok {
		listenerT.ExitProposition(s)
	}
}

func (p *ltlParser) Proposition() (localctx IPropositionContext) {
	return p.proposition(0)
}

func (p *ltlParser) proposition(_p int) (localctx IPropositionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPropositionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPropositionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ltlParserRULE_proposition, _p)
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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ltlParserT__0:
		{
			p.SetState(3)
			p.Match(ltlParserT__0)
		}

	case ltlParserT__1:
		{
			p.SetState(4)
			p.Match(ltlParserT__1)
		}

	case ltlParserATOMIC:
		{
			p.SetState(5)
			p.Match(ltlParserATOMIC)
		}

	case ltlParserT__2:
		{
			p.SetState(6)
			p.Match(ltlParserT__2)
		}
		{
			p.SetState(7)
			p.proposition(0)
		}
		{
			p.SetState(8)
			p.Match(ltlParserT__3)
		}

	case ltlParserLTL_NOT:
		{
			p.SetState(10)
			p.Match(ltlParserLTL_NOT)
		}
		{
			p.SetState(11)
			p.proposition(3)
		}

	case ltlParserLTL_GLOBALLY, ltlParserLTL_FINALLY, ltlParserLTL_NEXT:
		{
			p.SetState(12)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ltlParserLTL_GLOBALLY)|(1<<ltlParserLTL_FINALLY)|(1<<ltlParserLTL_NEXT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(13)
			p.proposition(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPropositionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ltlParserRULE_proposition)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(17)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ltlParserLTL_RIGHTWARDS_SINGLE_ARROW)|(1<<ltlParserLTL_AND)|(1<<ltlParserLTL_OR))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(18)
					p.proposition(5)
				}

			case 2:
				localctx = NewPropositionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ltlParserRULE_proposition)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(20)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ltlParserLTL_UNTIL)|(1<<ltlParserLTL_WEAK)|(1<<ltlParserLTL_RELEASE))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(21)
					p.proposition(2)
				}

			}

		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ltlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *ltlParser) Proposition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
