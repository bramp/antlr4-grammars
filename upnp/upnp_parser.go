// Code generated from Upnp.g4 by ANTLR 4.7.2. DO NOT EDIT.

package upnp // Upnp
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 105,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2,
	5, 2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3,
	24, 11, 3, 3, 3, 3, 3, 7, 3, 28, 10, 3, 12, 3, 14, 3, 31, 11, 3, 3, 3,
	3, 3, 5, 3, 35, 10, 3, 3, 3, 3, 3, 6, 3, 39, 10, 3, 13, 3, 14, 3, 40, 3,
	3, 3, 3, 6, 3, 45, 10, 3, 13, 3, 14, 3, 46, 3, 3, 7, 3, 50, 10, 3, 12,
	3, 14, 3, 53, 11, 3, 3, 4, 3, 4, 6, 4, 57, 10, 4, 13, 4, 14, 4, 58, 3,
	4, 3, 4, 6, 4, 63, 10, 4, 13, 4, 14, 4, 64, 3, 4, 3, 4, 3, 4, 6, 4, 70,
	10, 4, 13, 4, 14, 4, 71, 3, 4, 3, 4, 6, 4, 76, 10, 4, 13, 4, 14, 4, 77,
	3, 4, 5, 4, 81, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 7, 6, 88, 10, 6, 12,
	6, 14, 6, 91, 11, 6, 3, 6, 7, 6, 94, 10, 6, 12, 6, 14, 6, 97, 11, 6, 3,
	6, 7, 6, 100, 10, 6, 12, 6, 14, 6, 103, 11, 6, 3, 6, 2, 3, 4, 7, 2, 4,
	6, 8, 10, 2, 2, 2, 114, 2, 14, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 80, 3,
	2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2, 13,
	15, 7, 18, 2, 2, 14, 12, 3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 3, 3, 2, 2,
	2, 16, 17, 8, 3, 1, 2, 17, 35, 5, 6, 4, 2, 18, 22, 7, 3, 2, 2, 19, 21,
	7, 13, 2, 2, 20, 19, 3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	22, 23, 3, 2, 2, 2, 23, 25, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 29, 5,
	4, 3, 2, 26, 28, 7, 13, 2, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29,
	27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 32, 3, 2, 2, 2, 31, 29, 3, 2, 2,
	2, 32, 33, 7, 4, 2, 2, 33, 35, 3, 2, 2, 2, 34, 16, 3, 2, 2, 2, 34, 18,
	3, 2, 2, 2, 35, 51, 3, 2, 2, 2, 36, 38, 12, 4, 2, 2, 37, 39, 7, 13, 2,
	2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 7, 7, 2, 2, 43, 45, 7, 13, 2, 2,
	44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 5, 4, 3, 5, 49, 36, 3, 2, 2, 2, 50,
	53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 5, 3, 2, 2,
	2, 53, 51, 3, 2, 2, 2, 54, 56, 7, 14, 2, 2, 55, 57, 7, 13, 2, 2, 56, 55,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 60, 3, 2, 2, 2, 60, 62, 7, 8, 2, 2, 61, 63, 7, 13, 2, 2, 62, 61, 3,
	2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65,
	66, 3, 2, 2, 2, 66, 81, 5, 8, 5, 2, 67, 69, 7, 14, 2, 2, 68, 70, 7, 13,
	2, 2, 69, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 7, 11, 2, 2, 74, 76, 7, 13, 2,
	2, 75, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78,
	3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 7, 12, 2, 2, 80, 54, 3, 2, 2, 2,
	80, 67, 3, 2, 2, 2, 81, 7, 3, 2, 2, 2, 82, 83, 7, 17, 2, 2, 83, 84, 5,
	10, 6, 2, 84, 85, 7, 17, 2, 2, 85, 9, 3, 2, 2, 2, 86, 88, 7, 19, 2, 2,
	87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 95, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 94, 7, 13, 2, 2, 93,
	92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2,
	2, 96, 101, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 100, 7, 19, 2, 2, 99, 98,
	3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2,
	2, 2, 102, 11, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 17, 14, 22, 29, 34, 40,
	46, 51, 58, 64, 71, 77, 80, 89, 95, 101,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "", "", "", "", "'exists'", "", "", "", "'\t'",
	"' '", "'\"'", "'*'",
}
var symbolicNames = []string{
	"", "", "", "NUMBER", "WHITESPACE", "LOGOP", "BINOP", "RELOP", "STRINGOP",
	"EXISTSOP", "BOOLVAL", "WCHAR", "PROPERTY", "HTAB", "SPACE", "DQUOTE",
	"ASTERISK", "STRING_LITERAL",
}

var ruleNames = []string{
	"searchCrit", "searchExp", "relExp", "quotedVal", "escapedQuote",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type UpnpParser struct {
	*antlr.BaseParser
}

func NewUpnpParser(input antlr.TokenStream) *UpnpParser {
	this := new(UpnpParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Upnp.g4"

	return this
}

// UpnpParser tokens.
const (
	UpnpParserEOF            = antlr.TokenEOF
	UpnpParserT__0           = 1
	UpnpParserT__1           = 2
	UpnpParserNUMBER         = 3
	UpnpParserWHITESPACE     = 4
	UpnpParserLOGOP          = 5
	UpnpParserBINOP          = 6
	UpnpParserRELOP          = 7
	UpnpParserSTRINGOP       = 8
	UpnpParserEXISTSOP       = 9
	UpnpParserBOOLVAL        = 10
	UpnpParserWCHAR          = 11
	UpnpParserPROPERTY       = 12
	UpnpParserHTAB           = 13
	UpnpParserSPACE          = 14
	UpnpParserDQUOTE         = 15
	UpnpParserASTERISK       = 16
	UpnpParserSTRING_LITERAL = 17
)

// UpnpParser rules.
const (
	UpnpParserRULE_searchCrit   = 0
	UpnpParserRULE_searchExp    = 1
	UpnpParserRULE_relExp       = 2
	UpnpParserRULE_quotedVal    = 3
	UpnpParserRULE_escapedQuote = 4
)

// ISearchCritContext is an interface to support dynamic dispatch.
type ISearchCritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSearchCritContext differentiates from other interfaces.
	IsSearchCritContext()
}

type SearchCritContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchCritContext() *SearchCritContext {
	var p = new(SearchCritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpnpParserRULE_searchCrit
	return p
}

func (*SearchCritContext) IsSearchCritContext() {}

func NewSearchCritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchCritContext {
	var p = new(SearchCritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpnpParserRULE_searchCrit

	return p
}

func (s *SearchCritContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchCritContext) SearchExp() ISearchExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearchExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISearchExpContext)
}

func (s *SearchCritContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(UpnpParserASTERISK, 0)
}

func (s *SearchCritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchCritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchCritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.EnterSearchCrit(s)
	}
}

func (s *SearchCritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.ExitSearchCrit(s)
	}
}

func (p *UpnpParser) SearchCrit() (localctx ISearchCritContext) {
	localctx = NewSearchCritContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UpnpParserRULE_searchCrit)

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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpnpParserT__0, UpnpParserPROPERTY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.searchExp(0)
		}

	case UpnpParserASTERISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Match(UpnpParserASTERISK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISearchExpContext is an interface to support dynamic dispatch.
type ISearchExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSearchExpContext differentiates from other interfaces.
	IsSearchExpContext()
}

type SearchExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchExpContext() *SearchExpContext {
	var p = new(SearchExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpnpParserRULE_searchExp
	return p
}

func (*SearchExpContext) IsSearchExpContext() {}

func NewSearchExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchExpContext {
	var p = new(SearchExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpnpParserRULE_searchExp

	return p
}

func (s *SearchExpContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchExpContext) RelExp() IRelExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelExpContext)
}

func (s *SearchExpContext) AllSearchExp() []ISearchExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISearchExpContext)(nil)).Elem())
	var tst = make([]ISearchExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISearchExpContext)
		}
	}

	return tst
}

func (s *SearchExpContext) SearchExp(i int) ISearchExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISearchExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISearchExpContext)
}

func (s *SearchExpContext) AllWCHAR() []antlr.TerminalNode {
	return s.GetTokens(UpnpParserWCHAR)
}

func (s *SearchExpContext) WCHAR(i int) antlr.TerminalNode {
	return s.GetToken(UpnpParserWCHAR, i)
}

func (s *SearchExpContext) LOGOP() antlr.TerminalNode {
	return s.GetToken(UpnpParserLOGOP, 0)
}

func (s *SearchExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.EnterSearchExp(s)
	}
}

func (s *SearchExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.ExitSearchExp(s)
	}
}

func (p *UpnpParser) SearchExp() (localctx ISearchExpContext) {
	return p.searchExp(0)
}

func (p *UpnpParser) searchExp(_p int) (localctx ISearchExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSearchExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISearchExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, UpnpParserRULE_searchExp, _p)
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

	switch p.GetTokenStream().LA(1) {
	case UpnpParserPROPERTY:
		{
			p.SetState(15)
			p.RelExp()
		}

	case UpnpParserT__0:
		{
			p.SetState(16)
			p.Match(UpnpParserT__0)
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == UpnpParserWCHAR {
			{
				p.SetState(17)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(23)
			p.searchExp(0)
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == UpnpParserWCHAR {
			{
				p.SetState(24)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(30)
			p.Match(UpnpParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSearchExpContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, UpnpParserRULE_searchExp)
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == UpnpParserWCHAR {
				{
					p.SetState(35)
					p.Match(UpnpParserWCHAR)
				}

				p.SetState(38)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(40)
				p.Match(UpnpParserLOGOP)
			}
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == UpnpParserWCHAR {
				{
					p.SetState(41)
					p.Match(UpnpParserWCHAR)
				}

				p.SetState(44)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(46)
				p.searchExp(3)
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IRelExpContext is an interface to support dynamic dispatch.
type IRelExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelExpContext differentiates from other interfaces.
	IsRelExpContext()
}

type RelExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelExpContext() *RelExpContext {
	var p = new(RelExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpnpParserRULE_relExp
	return p
}

func (*RelExpContext) IsRelExpContext() {}

func NewRelExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelExpContext {
	var p = new(RelExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpnpParserRULE_relExp

	return p
}

func (s *RelExpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelExpContext) PROPERTY() antlr.TerminalNode {
	return s.GetToken(UpnpParserPROPERTY, 0)
}

func (s *RelExpContext) BINOP() antlr.TerminalNode {
	return s.GetToken(UpnpParserBINOP, 0)
}

func (s *RelExpContext) QuotedVal() IQuotedValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedValContext)
}

func (s *RelExpContext) AllWCHAR() []antlr.TerminalNode {
	return s.GetTokens(UpnpParserWCHAR)
}

func (s *RelExpContext) WCHAR(i int) antlr.TerminalNode {
	return s.GetToken(UpnpParserWCHAR, i)
}

func (s *RelExpContext) EXISTSOP() antlr.TerminalNode {
	return s.GetToken(UpnpParserEXISTSOP, 0)
}

func (s *RelExpContext) BOOLVAL() antlr.TerminalNode {
	return s.GetToken(UpnpParserBOOLVAL, 0)
}

func (s *RelExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.EnterRelExp(s)
	}
}

func (s *RelExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.ExitRelExp(s)
	}
}

func (p *UpnpParser) RelExp() (localctx IRelExpContext) {
	localctx = NewRelExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UpnpParserRULE_relExp)
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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(UpnpParserPROPERTY)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == UpnpParserWCHAR {
			{
				p.SetState(53)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(58)
			p.Match(UpnpParserBINOP)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == UpnpParserWCHAR {
			{
				p.SetState(59)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(64)
			p.QuotedVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(UpnpParserPROPERTY)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == UpnpParserWCHAR {
			{
				p.SetState(66)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(UpnpParserEXISTSOP)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == UpnpParserWCHAR {
			{
				p.SetState(72)
				p.Match(UpnpParserWCHAR)
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
			p.Match(UpnpParserBOOLVAL)
		}

	}

	return localctx
}

// IQuotedValContext is an interface to support dynamic dispatch.
type IQuotedValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedValContext differentiates from other interfaces.
	IsQuotedValContext()
}

type QuotedValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedValContext() *QuotedValContext {
	var p = new(QuotedValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpnpParserRULE_quotedVal
	return p
}

func (*QuotedValContext) IsQuotedValContext() {}

func NewQuotedValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedValContext {
	var p = new(QuotedValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpnpParserRULE_quotedVal

	return p
}

func (s *QuotedValContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedValContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(UpnpParserDQUOTE)
}

func (s *QuotedValContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(UpnpParserDQUOTE, i)
}

func (s *QuotedValContext) EscapedQuote() IEscapedQuoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscapedQuoteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscapedQuoteContext)
}

func (s *QuotedValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.EnterQuotedVal(s)
	}
}

func (s *QuotedValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.ExitQuotedVal(s)
	}
}

func (p *UpnpParser) QuotedVal() (localctx IQuotedValContext) {
	localctx = NewQuotedValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UpnpParserRULE_quotedVal)

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
		p.SetState(80)
		p.Match(UpnpParserDQUOTE)
	}
	{
		p.SetState(81)
		p.EscapedQuote()
	}
	{
		p.SetState(82)
		p.Match(UpnpParserDQUOTE)
	}

	return localctx
}

// IEscapedQuoteContext is an interface to support dynamic dispatch.
type IEscapedQuoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscapedQuoteContext differentiates from other interfaces.
	IsEscapedQuoteContext()
}

type EscapedQuoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscapedQuoteContext() *EscapedQuoteContext {
	var p = new(EscapedQuoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpnpParserRULE_escapedQuote
	return p
}

func (*EscapedQuoteContext) IsEscapedQuoteContext() {}

func NewEscapedQuoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EscapedQuoteContext {
	var p = new(EscapedQuoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpnpParserRULE_escapedQuote

	return p
}

func (s *EscapedQuoteContext) GetParser() antlr.Parser { return s.parser }

func (s *EscapedQuoteContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(UpnpParserSTRING_LITERAL)
}

func (s *EscapedQuoteContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(UpnpParserSTRING_LITERAL, i)
}

func (s *EscapedQuoteContext) AllWCHAR() []antlr.TerminalNode {
	return s.GetTokens(UpnpParserWCHAR)
}

func (s *EscapedQuoteContext) WCHAR(i int) antlr.TerminalNode {
	return s.GetToken(UpnpParserWCHAR, i)
}

func (s *EscapedQuoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscapedQuoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EscapedQuoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.EnterEscapedQuote(s)
	}
}

func (s *EscapedQuoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpnpListener); ok {
		listenerT.ExitEscapedQuote(s)
	}
}

func (p *UpnpParser) EscapedQuote() (localctx IEscapedQuoteContext) {
	localctx = NewEscapedQuoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UpnpParserRULE_escapedQuote)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.Match(UpnpParserSTRING_LITERAL)
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpnpParserWCHAR {
		{
			p.SetState(90)
			p.Match(UpnpParserWCHAR)
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpnpParserSTRING_LITERAL {
		{
			p.SetState(96)
			p.Match(UpnpParserSTRING_LITERAL)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *UpnpParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *SearchExpContext = nil
		if localctx != nil {
			t = localctx.(*SearchExpContext)
		}
		return p.SearchExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *UpnpParser) SearchExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
