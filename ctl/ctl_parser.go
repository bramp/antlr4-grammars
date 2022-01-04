// Code generated from ctl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package ctl // ctl
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 45, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 35, 10, 2, 3, 2, 3,
	2, 3, 2, 7, 2, 40, 10, 2, 12, 2, 14, 2, 43, 11, 2, 3, 2, 2, 3, 2, 3, 2,
	2, 4, 3, 2, 9, 11, 3, 2, 15, 18, 2, 52, 2, 34, 3, 2, 2, 2, 4, 5, 8, 2,
	1, 2, 5, 35, 7, 19, 2, 2, 6, 35, 7, 20, 2, 2, 7, 35, 7, 7, 2, 2, 8, 9,
	7, 12, 2, 2, 9, 10, 9, 2, 2, 2, 10, 35, 5, 2, 2, 9, 11, 12, 7, 13, 2, 2,
	12, 13, 9, 2, 2, 2, 13, 35, 5, 2, 2, 8, 14, 15, 7, 12, 2, 2, 15, 16, 7,
	3, 2, 2, 16, 17, 5, 2, 2, 2, 17, 18, 7, 8, 2, 2, 18, 19, 5, 2, 2, 2, 19,
	20, 7, 4, 2, 2, 20, 35, 3, 2, 2, 2, 21, 22, 7, 13, 2, 2, 22, 23, 7, 3,
	2, 2, 23, 24, 5, 2, 2, 2, 24, 25, 7, 8, 2, 2, 25, 26, 5, 2, 2, 2, 26, 27,
	7, 4, 2, 2, 27, 35, 3, 2, 2, 2, 28, 29, 7, 5, 2, 2, 29, 30, 5, 2, 2, 2,
	30, 31, 7, 6, 2, 2, 31, 35, 3, 2, 2, 2, 32, 33, 7, 21, 2, 2, 33, 35, 5,
	2, 2, 3, 34, 4, 3, 2, 2, 2, 34, 6, 3, 2, 2, 2, 34, 7, 3, 2, 2, 2, 34, 8,
	3, 2, 2, 2, 34, 11, 3, 2, 2, 2, 34, 14, 3, 2, 2, 2, 34, 21, 3, 2, 2, 2,
	34, 28, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 41, 3, 2, 2, 2, 36, 37, 12,
	4, 2, 2, 37, 38, 9, 3, 2, 2, 38, 40, 5, 2, 2, 5, 39, 36, 3, 2, 2, 2, 40,
	43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 3, 3, 2, 2,
	2, 43, 41, 3, 2, 2, 2, 4, 34, 41,
}
var literalNames = []string{
	"", "'['", "']'", "'('", "')'", "", "'U'", "'G'", "'F'", "'X'", "'A'",
	"'E'", "'\u2205'", "'\u21D4'", "'\u21D2'", "'\u2227'", "'\u2228'", "'\u22A4'",
	"'\u22A5'", "'\u2310'",
}
var symbolicNames = []string{
	"", "", "", "", "", "ATOMIC", "CTL_UNTIL", "CTL_GLOBALLY", "CTL_FINALLY",
	"CTL_NEXT", "CTL_INEVITABLE", "CTL_EXISTS", "CTL_PROPOSITION", "CTL_LEFT_RIGHT_DOUBLE_ARROW",
	"CTL_RIGHTWARDS_DOUBLE_ARROW", "CTL_AND", "CTL_OR", "CTL_DOWNTACK", "CTL_UPTACK",
	"CTL_NOT", "WS",
}

var ruleNames = []string{
	"proposition",
}

type ctlParser struct {
	*antlr.BaseParser
}

// NewctlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ctlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewctlParser(input antlr.TokenStream) *ctlParser {
	this := new(ctlParser)
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
	this.GrammarFileName = "ctl.g4"

	return this
}

// ctlParser tokens.
const (
	ctlParserEOF                         = antlr.TokenEOF
	ctlParserT__0                        = 1
	ctlParserT__1                        = 2
	ctlParserT__2                        = 3
	ctlParserT__3                        = 4
	ctlParserATOMIC                      = 5
	ctlParserCTL_UNTIL                   = 6
	ctlParserCTL_GLOBALLY                = 7
	ctlParserCTL_FINALLY                 = 8
	ctlParserCTL_NEXT                    = 9
	ctlParserCTL_INEVITABLE              = 10
	ctlParserCTL_EXISTS                  = 11
	ctlParserCTL_PROPOSITION             = 12
	ctlParserCTL_LEFT_RIGHT_DOUBLE_ARROW = 13
	ctlParserCTL_RIGHTWARDS_DOUBLE_ARROW = 14
	ctlParserCTL_AND                     = 15
	ctlParserCTL_OR                      = 16
	ctlParserCTL_DOWNTACK                = 17
	ctlParserCTL_UPTACK                  = 18
	ctlParserCTL_NOT                     = 19
	ctlParserWS                          = 20
)

// ctlParserRULE_proposition is the ctlParser rule.
const ctlParserRULE_proposition = 0

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
	p.RuleIndex = ctlParserRULE_proposition
	return p
}

func (*PropositionContext) IsPropositionContext() {}

func NewPropositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropositionContext {
	var p = new(PropositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ctlParserRULE_proposition

	return p
}

func (s *PropositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropositionContext) CTL_DOWNTACK() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_DOWNTACK, 0)
}

func (s *PropositionContext) CTL_UPTACK() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_UPTACK, 0)
}

func (s *PropositionContext) ATOMIC() antlr.TerminalNode {
	return s.GetToken(ctlParserATOMIC, 0)
}

func (s *PropositionContext) CTL_INEVITABLE() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_INEVITABLE, 0)
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

func (s *PropositionContext) CTL_NEXT() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_NEXT, 0)
}

func (s *PropositionContext) CTL_FINALLY() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_FINALLY, 0)
}

func (s *PropositionContext) CTL_GLOBALLY() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_GLOBALLY, 0)
}

func (s *PropositionContext) CTL_EXISTS() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_EXISTS, 0)
}

func (s *PropositionContext) CTL_UNTIL() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_UNTIL, 0)
}

func (s *PropositionContext) CTL_NOT() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_NOT, 0)
}

func (s *PropositionContext) CTL_AND() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_AND, 0)
}

func (s *PropositionContext) CTL_OR() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_OR, 0)
}

func (s *PropositionContext) CTL_RIGHTWARDS_DOUBLE_ARROW() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_RIGHTWARDS_DOUBLE_ARROW, 0)
}

func (s *PropositionContext) CTL_LEFT_RIGHT_DOUBLE_ARROW() antlr.TerminalNode {
	return s.GetToken(ctlParserCTL_LEFT_RIGHT_DOUBLE_ARROW, 0)
}

func (s *PropositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ctlListener); ok {
		listenerT.EnterProposition(s)
	}
}

func (s *PropositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ctlListener); ok {
		listenerT.ExitProposition(s)
	}
}

func (p *ctlParser) Proposition() (localctx IPropositionContext) {
	return p.proposition(0)
}

func (p *ctlParser) proposition(_p int) (localctx IPropositionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPropositionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPropositionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ctlParserRULE_proposition, _p)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(3)
			p.Match(ctlParserCTL_DOWNTACK)
		}

	case 2:
		{
			p.SetState(4)
			p.Match(ctlParserCTL_UPTACK)
		}

	case 3:
		{
			p.SetState(5)
			p.Match(ctlParserATOMIC)
		}

	case 4:
		{
			p.SetState(6)
			p.Match(ctlParserCTL_INEVITABLE)
		}
		{
			p.SetState(7)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ctlParserCTL_GLOBALLY)|(1<<ctlParserCTL_FINALLY)|(1<<ctlParserCTL_NEXT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(8)
			p.proposition(7)
		}

	case 5:
		{
			p.SetState(9)
			p.Match(ctlParserCTL_EXISTS)
		}
		{
			p.SetState(10)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ctlParserCTL_GLOBALLY)|(1<<ctlParserCTL_FINALLY)|(1<<ctlParserCTL_NEXT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(11)
			p.proposition(6)
		}

	case 6:
		{
			p.SetState(12)
			p.Match(ctlParserCTL_INEVITABLE)
		}
		{
			p.SetState(13)
			p.Match(ctlParserT__0)
		}
		{
			p.SetState(14)
			p.proposition(0)
		}
		{
			p.SetState(15)
			p.Match(ctlParserCTL_UNTIL)
		}
		{
			p.SetState(16)
			p.proposition(0)
		}
		{
			p.SetState(17)
			p.Match(ctlParserT__1)
		}

	case 7:
		{
			p.SetState(19)
			p.Match(ctlParserCTL_EXISTS)
		}
		{
			p.SetState(20)
			p.Match(ctlParserT__0)
		}
		{
			p.SetState(21)
			p.proposition(0)
		}
		{
			p.SetState(22)
			p.Match(ctlParserCTL_UNTIL)
		}
		{
			p.SetState(23)
			p.proposition(0)
		}
		{
			p.SetState(24)
			p.Match(ctlParserT__1)
		}

	case 8:
		{
			p.SetState(26)
			p.Match(ctlParserT__2)
		}
		{
			p.SetState(27)
			p.proposition(0)
		}
		{
			p.SetState(28)
			p.Match(ctlParserT__3)
		}

	case 9:
		{
			p.SetState(30)
			p.Match(ctlParserCTL_NOT)
		}
		{
			p.SetState(31)
			p.proposition(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPropositionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ctlParserRULE_proposition)
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(35)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ctlParserCTL_LEFT_RIGHT_DOUBLE_ARROW)|(1<<ctlParserCTL_RIGHTWARDS_DOUBLE_ARROW)|(1<<ctlParserCTL_AND)|(1<<ctlParserCTL_OR))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(36)
				p.proposition(3)
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ctlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *ctlParser) Proposition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
