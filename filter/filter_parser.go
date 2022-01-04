// Code generated from filter.g4 by ANTLR 4.9.3. DO NOT EDIT.

package filter // filter
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 102,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 45, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 60, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 79,
	10, 12, 3, 12, 3, 12, 5, 12, 83, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 5,
	14, 89, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 94, 10, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 2, 3, 3, 2, 10, 13, 2, 94, 2, 36, 3, 2,
	2, 2, 4, 44, 3, 2, 2, 2, 6, 46, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 52,
	3, 2, 2, 2, 12, 59, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2, 16, 66, 3, 2, 2, 2,
	18, 70, 3, 2, 2, 2, 20, 72, 3, 2, 2, 2, 22, 75, 3, 2, 2, 2, 24, 84, 3,
	2, 2, 2, 26, 86, 3, 2, 2, 2, 28, 90, 3, 2, 2, 2, 30, 95, 3, 2, 2, 2, 32,
	97, 3, 2, 2, 2, 34, 99, 3, 2, 2, 2, 36, 37, 7, 3, 2, 2, 37, 38, 5, 4, 3,
	2, 38, 39, 7, 4, 2, 2, 39, 3, 3, 2, 2, 2, 40, 45, 5, 6, 4, 2, 41, 45, 5,
	8, 5, 2, 42, 45, 5, 10, 6, 2, 43, 45, 5, 14, 8, 2, 44, 40, 3, 2, 2, 2,
	44, 41, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 5, 3, 2,
	2, 2, 46, 47, 7, 5, 2, 2, 47, 48, 5, 12, 7, 2, 48, 7, 3, 2, 2, 2, 49, 50,
	7, 6, 2, 2, 50, 51, 5, 12, 7, 2, 51, 9, 3, 2, 2, 2, 52, 53, 7, 7, 2, 2,
	53, 54, 5, 2, 2, 2, 54, 11, 3, 2, 2, 2, 55, 60, 5, 2, 2, 2, 56, 57, 5,
	2, 2, 2, 57, 58, 5, 12, 7, 2, 58, 60, 3, 2, 2, 2, 59, 55, 3, 2, 2, 2, 59,
	56, 3, 2, 2, 2, 60, 13, 3, 2, 2, 2, 61, 65, 5, 16, 9, 2, 62, 65, 5, 20,
	11, 2, 63, 65, 5, 22, 12, 2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64,
	63, 3, 2, 2, 2, 65, 15, 3, 2, 2, 2, 66, 67, 5, 32, 17, 2, 67, 68, 5, 18,
	10, 2, 68, 69, 5, 34, 18, 2, 69, 17, 3, 2, 2, 2, 70, 71, 9, 2, 2, 2, 71,
	19, 3, 2, 2, 2, 72, 73, 5, 32, 17, 2, 73, 74, 7, 8, 2, 2, 74, 21, 3, 2,
	2, 2, 75, 76, 5, 32, 17, 2, 76, 78, 7, 10, 2, 2, 77, 79, 5, 24, 13, 2,
	78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 5,
	26, 14, 2, 81, 83, 5, 30, 16, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 23, 3, 2, 2, 2, 84, 85, 5, 34, 18, 2, 85, 25, 3, 2, 2, 2, 86, 88, 7,
	9, 2, 2, 87, 89, 5, 28, 15, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2,
	89, 27, 3, 2, 2, 2, 90, 91, 5, 34, 18, 2, 91, 93, 7, 9, 2, 2, 92, 94, 5,
	28, 15, 2, 93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 29, 3, 2, 2, 2,
	95, 96, 5, 34, 18, 2, 96, 31, 3, 2, 2, 2, 97, 98, 7, 14, 2, 2, 98, 33,
	3, 2, 2, 2, 99, 100, 7, 14, 2, 2, 100, 35, 3, 2, 2, 2, 9, 44, 59, 64, 78,
	82, 88, 93,
}
var literalNames = []string{
	"", "'('", "')'", "'&'", "'|'", "'!'", "'=*'", "'*'", "'='", "'~='", "'>='",
	"'<='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "EQUAL", "APPROX", "GREATER", "LESS", "OCTETSTRING",
}

var ruleNames = []string{
	"filter_", "filtercomp", "and_", "or_", "not_", "filterlist", "item", "simple",
	"filtertype", "present", "substring", "initial", "any_", "starval", "final_",
	"attr", "value",
}

type filterParser struct {
	*antlr.BaseParser
}

// NewfilterParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *filterParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfilterParser(input antlr.TokenStream) *filterParser {
	this := new(filterParser)
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
	this.GrammarFileName = "filter.g4"

	return this
}

// filterParser tokens.
const (
	filterParserEOF         = antlr.TokenEOF
	filterParserT__0        = 1
	filterParserT__1        = 2
	filterParserT__2        = 3
	filterParserT__3        = 4
	filterParserT__4        = 5
	filterParserT__5        = 6
	filterParserT__6        = 7
	filterParserEQUAL       = 8
	filterParserAPPROX      = 9
	filterParserGREATER     = 10
	filterParserLESS        = 11
	filterParserOCTETSTRING = 12
)

// filterParser rules.
const (
	filterParserRULE_filter_    = 0
	filterParserRULE_filtercomp = 1
	filterParserRULE_and_       = 2
	filterParserRULE_or_        = 3
	filterParserRULE_not_       = 4
	filterParserRULE_filterlist = 5
	filterParserRULE_item       = 6
	filterParserRULE_simple     = 7
	filterParserRULE_filtertype = 8
	filterParserRULE_present    = 9
	filterParserRULE_substring  = 10
	filterParserRULE_initial    = 11
	filterParserRULE_any_       = 12
	filterParserRULE_starval    = 13
	filterParserRULE_final_     = 14
	filterParserRULE_attr       = 15
	filterParserRULE_value      = 16
)

// IFilter_Context is an interface to support dynamic dispatch.
type IFilter_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_Context differentiates from other interfaces.
	IsFilter_Context()
}

type Filter_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_Context() *Filter_Context {
	var p = new(Filter_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_filter_
	return p
}

func (*Filter_Context) IsFilter_Context() {}

func NewFilter_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_Context {
	var p = new(Filter_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_filter_

	return p
}

func (s *Filter_Context) GetParser() antlr.Parser { return s.parser }

func (s *Filter_Context) Filtercomp() IFiltercompContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFiltercompContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFiltercompContext)
}

func (s *Filter_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterFilter_(s)
	}
}

func (s *Filter_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitFilter_(s)
	}
}

func (p *filterParser) Filter_() (localctx IFilter_Context) {
	this := p
	_ = this

	localctx = NewFilter_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, filterParserRULE_filter_)

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
		p.Match(filterParserT__0)
	}
	{
		p.SetState(35)
		p.Filtercomp()
	}
	{
		p.SetState(36)
		p.Match(filterParserT__1)
	}

	return localctx
}

// IFiltercompContext is an interface to support dynamic dispatch.
type IFiltercompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiltercompContext differentiates from other interfaces.
	IsFiltercompContext()
}

type FiltercompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiltercompContext() *FiltercompContext {
	var p = new(FiltercompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_filtercomp
	return p
}

func (*FiltercompContext) IsFiltercompContext() {}

func NewFiltercompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiltercompContext {
	var p = new(FiltercompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_filtercomp

	return p
}

func (s *FiltercompContext) GetParser() antlr.Parser { return s.parser }

func (s *FiltercompContext) And_() IAnd_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_Context)
}

func (s *FiltercompContext) Or_() IOr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_Context)
}

func (s *FiltercompContext) Not_() INot_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_Context)
}

func (s *FiltercompContext) Item() IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *FiltercompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiltercompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiltercompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterFiltercomp(s)
	}
}

func (s *FiltercompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitFiltercomp(s)
	}
}

func (p *filterParser) Filtercomp() (localctx IFiltercompContext) {
	this := p
	_ = this

	localctx = NewFiltercompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, filterParserRULE_filtercomp)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case filterParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.And_()
		}

	case filterParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Or_()
		}

	case filterParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Not_()
		}

	case filterParserOCTETSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Item()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAnd_Context is an interface to support dynamic dispatch.
type IAnd_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_Context differentiates from other interfaces.
	IsAnd_Context()
}

type And_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_Context() *And_Context {
	var p = new(And_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_and_
	return p
}

func (*And_Context) IsAnd_Context() {}

func NewAnd_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_Context {
	var p = new(And_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_and_

	return p
}

func (s *And_Context) GetParser() antlr.Parser { return s.parser }

func (s *And_Context) Filterlist() IFilterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterlistContext)
}

func (s *And_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterAnd_(s)
	}
}

func (s *And_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitAnd_(s)
	}
}

func (p *filterParser) And_() (localctx IAnd_Context) {
	this := p
	_ = this

	localctx = NewAnd_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, filterParserRULE_and_)

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
		p.Match(filterParserT__2)
	}
	{
		p.SetState(45)
		p.Filterlist()
	}

	return localctx
}

// IOr_Context is an interface to support dynamic dispatch.
type IOr_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_Context differentiates from other interfaces.
	IsOr_Context()
}

type Or_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_Context() *Or_Context {
	var p = new(Or_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_or_
	return p
}

func (*Or_Context) IsOr_Context() {}

func NewOr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_Context {
	var p = new(Or_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_or_

	return p
}

func (s *Or_Context) GetParser() antlr.Parser { return s.parser }

func (s *Or_Context) Filterlist() IFilterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterlistContext)
}

func (s *Or_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterOr_(s)
	}
}

func (s *Or_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitOr_(s)
	}
}

func (p *filterParser) Or_() (localctx IOr_Context) {
	this := p
	_ = this

	localctx = NewOr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, filterParserRULE_or_)

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
		p.SetState(47)
		p.Match(filterParserT__3)
	}
	{
		p.SetState(48)
		p.Filterlist()
	}

	return localctx
}

// INot_Context is an interface to support dynamic dispatch.
type INot_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_Context differentiates from other interfaces.
	IsNot_Context()
}

type Not_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_Context() *Not_Context {
	var p = new(Not_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_not_
	return p
}

func (*Not_Context) IsNot_Context() {}

func NewNot_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_Context {
	var p = new(Not_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_not_

	return p
}

func (s *Not_Context) GetParser() antlr.Parser { return s.parser }

func (s *Not_Context) Filter_() IFilter_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilter_Context)
}

func (s *Not_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterNot_(s)
	}
}

func (s *Not_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitNot_(s)
	}
}

func (p *filterParser) Not_() (localctx INot_Context) {
	this := p
	_ = this

	localctx = NewNot_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, filterParserRULE_not_)

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
		p.SetState(50)
		p.Match(filterParserT__4)
	}
	{
		p.SetState(51)
		p.Filter_()
	}

	return localctx
}

// IFilterlistContext is an interface to support dynamic dispatch.
type IFilterlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterlistContext differentiates from other interfaces.
	IsFilterlistContext()
}

type FilterlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterlistContext() *FilterlistContext {
	var p = new(FilterlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_filterlist
	return p
}

func (*FilterlistContext) IsFilterlistContext() {}

func NewFilterlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterlistContext {
	var p = new(FilterlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_filterlist

	return p
}

func (s *FilterlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterlistContext) Filter_() IFilter_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilter_Context)
}

func (s *FilterlistContext) Filterlist() IFilterlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterlistContext)
}

func (s *FilterlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterFilterlist(s)
	}
}

func (s *FilterlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitFilterlist(s)
	}
}

func (p *filterParser) Filterlist() (localctx IFilterlistContext) {
	this := p
	_ = this

	localctx = NewFilterlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, filterParserRULE_filterlist)

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

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Filter_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Filter_()
		}
		{
			p.SetState(55)
			p.Filterlist()
		}

	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) Simple() ISimpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleContext)
}

func (s *ItemContext) Present() IPresentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPresentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPresentContext)
}

func (s *ItemContext) Substring() ISubstringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubstringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubstringContext)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *filterParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, filterParserRULE_item)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Simple()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Present()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Substring()
		}

	}

	return localctx
}

// ISimpleContext is an interface to support dynamic dispatch.
type ISimpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleContext differentiates from other interfaces.
	IsSimpleContext()
}

type SimpleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleContext() *SimpleContext {
	var p = new(SimpleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_simple
	return p
}

func (*SimpleContext) IsSimpleContext() {}

func NewSimpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleContext {
	var p = new(SimpleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_simple

	return p
}

func (s *SimpleContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *SimpleContext) Filtertype() IFiltertypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFiltertypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFiltertypeContext)
}

func (s *SimpleContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterSimple(s)
	}
}

func (s *SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitSimple(s)
	}
}

func (p *filterParser) Simple() (localctx ISimpleContext) {
	this := p
	_ = this

	localctx = NewSimpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, filterParserRULE_simple)

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
		p.SetState(64)
		p.Attr()
	}
	{
		p.SetState(65)
		p.Filtertype()
	}
	{
		p.SetState(66)
		p.Value()
	}

	return localctx
}

// IFiltertypeContext is an interface to support dynamic dispatch.
type IFiltertypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiltertypeContext differentiates from other interfaces.
	IsFiltertypeContext()
}

type FiltertypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiltertypeContext() *FiltertypeContext {
	var p = new(FiltertypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_filtertype
	return p
}

func (*FiltertypeContext) IsFiltertypeContext() {}

func NewFiltertypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiltertypeContext {
	var p = new(FiltertypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_filtertype

	return p
}

func (s *FiltertypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FiltertypeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(filterParserEQUAL, 0)
}

func (s *FiltertypeContext) APPROX() antlr.TerminalNode {
	return s.GetToken(filterParserAPPROX, 0)
}

func (s *FiltertypeContext) GREATER() antlr.TerminalNode {
	return s.GetToken(filterParserGREATER, 0)
}

func (s *FiltertypeContext) LESS() antlr.TerminalNode {
	return s.GetToken(filterParserLESS, 0)
}

func (s *FiltertypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiltertypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiltertypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterFiltertype(s)
	}
}

func (s *FiltertypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitFiltertype(s)
	}
}

func (p *filterParser) Filtertype() (localctx IFiltertypeContext) {
	this := p
	_ = this

	localctx = NewFiltertypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, filterParserRULE_filtertype)
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
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<filterParserEQUAL)|(1<<filterParserAPPROX)|(1<<filterParserGREATER)|(1<<filterParserLESS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPresentContext is an interface to support dynamic dispatch.
type IPresentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPresentContext differentiates from other interfaces.
	IsPresentContext()
}

type PresentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPresentContext() *PresentContext {
	var p = new(PresentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_present
	return p
}

func (*PresentContext) IsPresentContext() {}

func NewPresentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PresentContext {
	var p = new(PresentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_present

	return p
}

func (s *PresentContext) GetParser() antlr.Parser { return s.parser }

func (s *PresentContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *PresentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PresentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterPresent(s)
	}
}

func (s *PresentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitPresent(s)
	}
}

func (p *filterParser) Present() (localctx IPresentContext) {
	this := p
	_ = this

	localctx = NewPresentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, filterParserRULE_present)

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
		p.SetState(70)
		p.Attr()
	}
	{
		p.SetState(71)
		p.Match(filterParserT__5)
	}

	return localctx
}

// ISubstringContext is an interface to support dynamic dispatch.
type ISubstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubstringContext differentiates from other interfaces.
	IsSubstringContext()
}

type SubstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubstringContext() *SubstringContext {
	var p = new(SubstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_substring
	return p
}

func (*SubstringContext) IsSubstringContext() {}

func NewSubstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubstringContext {
	var p = new(SubstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_substring

	return p
}

func (s *SubstringContext) GetParser() antlr.Parser { return s.parser }

func (s *SubstringContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *SubstringContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(filterParserEQUAL, 0)
}

func (s *SubstringContext) Any_() IAny_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_Context)
}

func (s *SubstringContext) Initial() IInitialContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitialContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitialContext)
}

func (s *SubstringContext) Final_() IFinal_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFinal_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFinal_Context)
}

func (s *SubstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterSubstring(s)
	}
}

func (s *SubstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitSubstring(s)
	}
}

func (p *filterParser) Substring() (localctx ISubstringContext) {
	this := p
	_ = this

	localctx = NewSubstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, filterParserRULE_substring)
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
		p.Attr()
	}
	{
		p.SetState(74)
		p.Match(filterParserEQUAL)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserOCTETSTRING {
		{
			p.SetState(75)
			p.Initial()
		}

	}
	{
		p.SetState(78)
		p.Any_()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == filterParserOCTETSTRING {
		{
			p.SetState(79)
			p.Final_()
		}

	}

	return localctx
}

// IInitialContext is an interface to support dynamic dispatch.
type IInitialContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitialContext differentiates from other interfaces.
	IsInitialContext()
}

type InitialContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitialContext() *InitialContext {
	var p = new(InitialContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_initial
	return p
}

func (*InitialContext) IsInitialContext() {}

func NewInitialContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitialContext {
	var p = new(InitialContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_initial

	return p
}

func (s *InitialContext) GetParser() antlr.Parser { return s.parser }

func (s *InitialContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitialContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterInitial(s)
	}
}

func (s *InitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitInitial(s)
	}
}

func (p *filterParser) Initial() (localctx IInitialContext) {
	this := p
	_ = this

	localctx = NewInitialContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, filterParserRULE_initial)

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
		p.SetState(82)
		p.Value()
	}

	return localctx
}

// IAny_Context is an interface to support dynamic dispatch.
type IAny_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_Context differentiates from other interfaces.
	IsAny_Context()
}

type Any_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_Context() *Any_Context {
	var p = new(Any_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_any_
	return p
}

func (*Any_Context) IsAny_Context() {}

func NewAny_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_Context {
	var p = new(Any_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_any_

	return p
}

func (s *Any_Context) GetParser() antlr.Parser { return s.parser }

func (s *Any_Context) Starval() IStarvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarvalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarvalContext)
}

func (s *Any_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterAny_(s)
	}
}

func (s *Any_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitAny_(s)
	}
}

func (p *filterParser) Any_() (localctx IAny_Context) {
	this := p
	_ = this

	localctx = NewAny_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, filterParserRULE_any_)

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
		p.SetState(84)
		p.Match(filterParserT__6)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(85)
			p.Starval()
		}

	}

	return localctx
}

// IStarvalContext is an interface to support dynamic dispatch.
type IStarvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStarvalContext differentiates from other interfaces.
	IsStarvalContext()
}

type StarvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarvalContext() *StarvalContext {
	var p = new(StarvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_starval
	return p
}

func (*StarvalContext) IsStarvalContext() {}

func NewStarvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarvalContext {
	var p = new(StarvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_starval

	return p
}

func (s *StarvalContext) GetParser() antlr.Parser { return s.parser }

func (s *StarvalContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *StarvalContext) Starval() IStarvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarvalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarvalContext)
}

func (s *StarvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterStarval(s)
	}
}

func (s *StarvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitStarval(s)
	}
}

func (p *filterParser) Starval() (localctx IStarvalContext) {
	this := p
	_ = this

	localctx = NewStarvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, filterParserRULE_starval)

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
		p.SetState(88)
		p.Value()
	}
	{
		p.SetState(89)
		p.Match(filterParserT__6)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(90)
			p.Starval()
		}

	}

	return localctx
}

// IFinal_Context is an interface to support dynamic dispatch.
type IFinal_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFinal_Context differentiates from other interfaces.
	IsFinal_Context()
}

type Final_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinal_Context() *Final_Context {
	var p = new(Final_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_final_
	return p
}

func (*Final_Context) IsFinal_Context() {}

func NewFinal_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Final_Context {
	var p = new(Final_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_final_

	return p
}

func (s *Final_Context) GetParser() antlr.Parser { return s.parser }

func (s *Final_Context) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Final_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Final_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Final_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterFinal_(s)
	}
}

func (s *Final_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitFinal_(s)
	}
}

func (p *filterParser) Final_() (localctx IFinal_Context) {
	this := p
	_ = this

	localctx = NewFinal_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, filterParserRULE_final_)

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
		p.SetState(93)
		p.Value()
	}

	return localctx
}

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) OCTETSTRING() antlr.TerminalNode {
	return s.GetToken(filterParserOCTETSTRING, 0)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *filterParser) Attr() (localctx IAttrContext) {
	this := p
	_ = this

	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, filterParserRULE_attr)

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
		p.SetState(95)
		p.Match(filterParserOCTETSTRING)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = filterParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = filterParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) OCTETSTRING() antlr.TerminalNode {
	return s.GetToken(filterParserOCTETSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(filterListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *filterParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, filterParserRULE_value)

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
		p.SetState(97)
		p.Match(filterParserOCTETSTRING)
	}

	return localctx
}
