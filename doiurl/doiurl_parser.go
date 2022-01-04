// Code generated from doiurl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package doiurl // doiurl
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 66, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 24, 10, 2, 3,
	2, 3, 2, 5, 2, 28, 10, 2, 3, 2, 5, 2, 31, 10, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 44, 10, 6, 12, 6, 14, 6,
	47, 11, 6, 3, 7, 6, 7, 50, 10, 7, 13, 7, 14, 7, 51, 3, 8, 7, 8, 55, 10,
	8, 12, 8, 14, 8, 58, 11, 8, 3, 9, 7, 9, 61, 10, 9, 12, 9, 14, 9, 64, 11,
	9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 3, 4, 2, 4, 4, 7, 8,
	2, 64, 2, 18, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 38,
	3, 2, 2, 2, 10, 40, 3, 2, 2, 2, 12, 49, 3, 2, 2, 2, 14, 56, 3, 2, 2, 2,
	16, 62, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 20, 7, 3, 2, 2, 20, 23, 5,
	6, 4, 2, 21, 22, 7, 4, 2, 2, 22, 24, 5, 14, 8, 2, 23, 21, 3, 2, 2, 2, 23,
	24, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 26, 7, 5, 2, 2, 26, 28, 5, 16,
	9, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 31,
	7, 2, 2, 3, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 3, 3, 2, 2, 2,
	32, 33, 7, 6, 2, 2, 33, 5, 3, 2, 2, 2, 34, 35, 5, 8, 5, 2, 35, 36, 7, 7,
	2, 2, 36, 37, 5, 10, 6, 2, 37, 7, 3, 2, 2, 2, 38, 39, 5, 12, 7, 2, 39,
	9, 3, 2, 2, 2, 40, 45, 5, 12, 7, 2, 41, 42, 7, 7, 2, 2, 42, 44, 5, 12,
	7, 2, 43, 41, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46,
	3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 50, 7, 8, 2, 2,
	49, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3,
	2, 2, 2, 52, 13, 3, 2, 2, 2, 53, 55, 9, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55,
	58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 15, 3, 2, 2,
	2, 58, 56, 3, 2, 2, 2, 59, 61, 9, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 64,
	3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 17, 3, 2, 2, 2,
	64, 62, 3, 2, 2, 2, 9, 23, 27, 30, 45, 51, 56, 62,
}
var literalNames = []string{
	"", "':'", "'?'", "'#'", "'doi'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "PCHAR", "WS",
}

var ruleNames = []string{
	"doiuri", "scheme", "encodeddoi", "prefix_", "suffix", "segment", "query",
	"fragment_",
}

type doiurlParser struct {
	*antlr.BaseParser
}

// NewdoiurlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *doiurlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdoiurlParser(input antlr.TokenStream) *doiurlParser {
	this := new(doiurlParser)
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
	this.GrammarFileName = "doiurl.g4"

	return this
}

// doiurlParser tokens.
const (
	doiurlParserEOF   = antlr.TokenEOF
	doiurlParserT__0  = 1
	doiurlParserT__1  = 2
	doiurlParserT__2  = 3
	doiurlParserT__3  = 4
	doiurlParserT__4  = 5
	doiurlParserPCHAR = 6
	doiurlParserWS    = 7
)

// doiurlParser rules.
const (
	doiurlParserRULE_doiuri     = 0
	doiurlParserRULE_scheme     = 1
	doiurlParserRULE_encodeddoi = 2
	doiurlParserRULE_prefix_    = 3
	doiurlParserRULE_suffix     = 4
	doiurlParserRULE_segment    = 5
	doiurlParserRULE_query      = 6
	doiurlParserRULE_fragment_  = 7
)

// IDoiuriContext is an interface to support dynamic dispatch.
type IDoiuriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoiuriContext differentiates from other interfaces.
	IsDoiuriContext()
}

type DoiuriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoiuriContext() *DoiuriContext {
	var p = new(DoiuriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_doiuri
	return p
}

func (*DoiuriContext) IsDoiuriContext() {}

func NewDoiuriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoiuriContext {
	var p = new(DoiuriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_doiuri

	return p
}

func (s *DoiuriContext) GetParser() antlr.Parser { return s.parser }

func (s *DoiuriContext) Scheme() ISchemeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemeContext)
}

func (s *DoiuriContext) Encodeddoi() IEncodeddoiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodeddoiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEncodeddoiContext)
}

func (s *DoiuriContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *DoiuriContext) Fragment_() IFragment_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragment_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragment_Context)
}

func (s *DoiuriContext) EOF() antlr.TerminalNode {
	return s.GetToken(doiurlParserEOF, 0)
}

func (s *DoiuriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoiuriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoiuriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterDoiuri(s)
	}
}

func (s *DoiuriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitDoiuri(s)
	}
}

func (p *doiurlParser) Doiuri() (localctx IDoiuriContext) {
	this := p
	_ = this

	localctx = NewDoiuriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, doiurlParserRULE_doiuri)
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
		p.Scheme()
	}
	{
		p.SetState(17)
		p.Match(doiurlParserT__0)
	}
	{
		p.SetState(18)
		p.Encodeddoi()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == doiurlParserT__1 {
		{
			p.SetState(19)
			p.Match(doiurlParserT__1)
		}
		{
			p.SetState(20)
			p.Query()
		}

	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == doiurlParserT__2 {
		{
			p.SetState(23)
			p.Match(doiurlParserT__2)
		}
		{
			p.SetState(24)
			p.Fragment_()
		}

	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(27)
			p.Match(doiurlParserEOF)
		}

	}

	return localctx
}

// ISchemeContext is an interface to support dynamic dispatch.
type ISchemeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemeContext differentiates from other interfaces.
	IsSchemeContext()
}

type SchemeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemeContext() *SchemeContext {
	var p = new(SchemeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_scheme
	return p
}

func (*SchemeContext) IsSchemeContext() {}

func NewSchemeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemeContext {
	var p = new(SchemeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_scheme

	return p
}

func (s *SchemeContext) GetParser() antlr.Parser { return s.parser }
func (s *SchemeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterScheme(s)
	}
}

func (s *SchemeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitScheme(s)
	}
}

func (p *doiurlParser) Scheme() (localctx ISchemeContext) {
	this := p
	_ = this

	localctx = NewSchemeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, doiurlParserRULE_scheme)

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
		p.SetState(30)
		p.Match(doiurlParserT__3)
	}

	return localctx
}

// IEncodeddoiContext is an interface to support dynamic dispatch.
type IEncodeddoiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodeddoiContext differentiates from other interfaces.
	IsEncodeddoiContext()
}

type EncodeddoiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodeddoiContext() *EncodeddoiContext {
	var p = new(EncodeddoiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_encodeddoi
	return p
}

func (*EncodeddoiContext) IsEncodeddoiContext() {}

func NewEncodeddoiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodeddoiContext {
	var p = new(EncodeddoiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_encodeddoi

	return p
}

func (s *EncodeddoiContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodeddoiContext) Prefix_() IPrefix_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_Context)
}

func (s *EncodeddoiContext) Suffix() ISuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuffixContext)
}

func (s *EncodeddoiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodeddoiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodeddoiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterEncodeddoi(s)
	}
}

func (s *EncodeddoiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitEncodeddoi(s)
	}
}

func (p *doiurlParser) Encodeddoi() (localctx IEncodeddoiContext) {
	this := p
	_ = this

	localctx = NewEncodeddoiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, doiurlParserRULE_encodeddoi)

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
		p.Prefix_()
	}
	{
		p.SetState(33)
		p.Match(doiurlParserT__4)
	}
	{
		p.SetState(34)
		p.Suffix()
	}

	return localctx
}

// IPrefix_Context is an interface to support dynamic dispatch.
type IPrefix_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefix_Context differentiates from other interfaces.
	IsPrefix_Context()
}

type Prefix_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_Context() *Prefix_Context {
	var p = new(Prefix_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_prefix_
	return p
}

func (*Prefix_Context) IsPrefix_Context() {}

func NewPrefix_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_Context {
	var p = new(Prefix_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_prefix_

	return p
}

func (s *Prefix_Context) GetParser() antlr.Parser { return s.parser }

func (s *Prefix_Context) Segment() ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *Prefix_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterPrefix_(s)
	}
}

func (s *Prefix_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitPrefix_(s)
	}
}

func (p *doiurlParser) Prefix_() (localctx IPrefix_Context) {
	this := p
	_ = this

	localctx = NewPrefix_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, doiurlParserRULE_prefix_)

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
		p.SetState(36)
		p.Segment()
	}

	return localctx
}

// ISuffixContext is an interface to support dynamic dispatch.
type ISuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuffixContext differentiates from other interfaces.
	IsSuffixContext()
}

type SuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuffixContext() *SuffixContext {
	var p = new(SuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_suffix
	return p
}

func (*SuffixContext) IsSuffixContext() {}

func NewSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuffixContext {
	var p = new(SuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_suffix

	return p
}

func (s *SuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *SuffixContext) AllSegment() []ISegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISegmentContext)(nil)).Elem())
	var tst = make([]ISegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISegmentContext)
		}
	}

	return tst
}

func (s *SuffixContext) Segment(i int) ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *SuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterSuffix(s)
	}
}

func (s *SuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitSuffix(s)
	}
}

func (p *doiurlParser) Suffix() (localctx ISuffixContext) {
	this := p
	_ = this

	localctx = NewSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, doiurlParserRULE_suffix)
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
		p.SetState(38)
		p.Segment()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == doiurlParserT__4 {
		{
			p.SetState(39)
			p.Match(doiurlParserT__4)
		}
		{
			p.SetState(40)
			p.Segment()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_segment
	return p
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) AllPCHAR() []antlr.TerminalNode {
	return s.GetTokens(doiurlParserPCHAR)
}

func (s *SegmentContext) PCHAR(i int) antlr.TerminalNode {
	return s.GetToken(doiurlParserPCHAR, i)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *doiurlParser) Segment() (localctx ISegmentContext) {
	this := p
	_ = this

	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, doiurlParserRULE_segment)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == doiurlParserPCHAR {
		{
			p.SetState(46)
			p.Match(doiurlParserPCHAR)
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) AllPCHAR() []antlr.TerminalNode {
	return s.GetTokens(doiurlParserPCHAR)
}

func (s *QueryContext) PCHAR(i int) antlr.TerminalNode {
	return s.GetToken(doiurlParserPCHAR, i)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *doiurlParser) Query() (localctx IQueryContext) {
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, doiurlParserRULE_query)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<doiurlParserT__1)|(1<<doiurlParserT__4)|(1<<doiurlParserPCHAR))) != 0 {
		{
			p.SetState(51)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<doiurlParserT__1)|(1<<doiurlParserT__4)|(1<<doiurlParserPCHAR))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFragment_Context is an interface to support dynamic dispatch.
type IFragment_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragment_Context differentiates from other interfaces.
	IsFragment_Context()
}

type Fragment_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragment_Context() *Fragment_Context {
	var p = new(Fragment_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = doiurlParserRULE_fragment_
	return p
}

func (*Fragment_Context) IsFragment_Context() {}

func NewFragment_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fragment_Context {
	var p = new(Fragment_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = doiurlParserRULE_fragment_

	return p
}

func (s *Fragment_Context) GetParser() antlr.Parser { return s.parser }

func (s *Fragment_Context) AllPCHAR() []antlr.TerminalNode {
	return s.GetTokens(doiurlParserPCHAR)
}

func (s *Fragment_Context) PCHAR(i int) antlr.TerminalNode {
	return s.GetToken(doiurlParserPCHAR, i)
}

func (s *Fragment_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fragment_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fragment_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.EnterFragment_(s)
	}
}

func (s *Fragment_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(doiurlListener); ok {
		listenerT.ExitFragment_(s)
	}
}

func (p *doiurlParser) Fragment_() (localctx IFragment_Context) {
	this := p
	_ = this

	localctx = NewFragment_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, doiurlParserRULE_fragment_)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<doiurlParserT__1)|(1<<doiurlParserT__4)|(1<<doiurlParserPCHAR))) != 0 {
		{
			p.SetState(57)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<doiurlParserT__1)|(1<<doiurlParserT__4)|(1<<doiurlParserPCHAR))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
