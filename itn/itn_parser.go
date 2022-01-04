// Code generated from itn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package itn // itn
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 40, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 34, 2, 17, 3,
	2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 33, 3, 2, 2, 2, 10, 35,
	3, 2, 2, 2, 12, 37, 3, 2, 2, 2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2,
	16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 20, 3,
	2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 21, 7, 2, 2, 3, 21, 3, 3, 2, 2, 2, 22,
	23, 5, 6, 4, 2, 23, 24, 7, 3, 2, 2, 24, 25, 5, 8, 5, 2, 25, 26, 7, 3, 2,
	2, 26, 27, 5, 10, 6, 2, 27, 28, 7, 3, 2, 2, 28, 29, 5, 12, 7, 2, 29, 30,
	7, 3, 2, 2, 30, 5, 3, 2, 2, 2, 31, 32, 7, 4, 2, 2, 32, 7, 3, 2, 2, 2, 33,
	34, 7, 4, 2, 2, 34, 9, 3, 2, 2, 2, 35, 36, 7, 5, 2, 2, 36, 11, 3, 2, 2,
	2, 37, 38, 7, 4, 2, 2, 38, 13, 3, 2, 2, 2, 3, 17,
}
var literalNames = []string{
	"", "'|'",
}
var symbolicNames = []string{
	"", "", "NUM", "TEXT", "WS",
}

var ruleNames = []string{
	"itinerary", "line", "longitude", "latitude", "descr", "flag",
}

type itnParser struct {
	*antlr.BaseParser
}

// NewitnParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *itnParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewitnParser(input antlr.TokenStream) *itnParser {
	this := new(itnParser)
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
	this.GrammarFileName = "itn.g4"

	return this
}

// itnParser tokens.
const (
	itnParserEOF  = antlr.TokenEOF
	itnParserT__0 = 1
	itnParserNUM  = 2
	itnParserTEXT = 3
	itnParserWS   = 4
)

// itnParser rules.
const (
	itnParserRULE_itinerary = 0
	itnParserRULE_line      = 1
	itnParserRULE_longitude = 2
	itnParserRULE_latitude  = 3
	itnParserRULE_descr     = 4
	itnParserRULE_flag      = 5
)

// IItineraryContext is an interface to support dynamic dispatch.
type IItineraryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItineraryContext differentiates from other interfaces.
	IsItineraryContext()
}

type ItineraryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItineraryContext() *ItineraryContext {
	var p = new(ItineraryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_itinerary
	return p
}

func (*ItineraryContext) IsItineraryContext() {}

func NewItineraryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItineraryContext {
	var p = new(ItineraryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_itinerary

	return p
}

func (s *ItineraryContext) GetParser() antlr.Parser { return s.parser }

func (s *ItineraryContext) EOF() antlr.TerminalNode {
	return s.GetToken(itnParserEOF, 0)
}

func (s *ItineraryContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ItineraryContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ItineraryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItineraryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItineraryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterItinerary(s)
	}
}

func (s *ItineraryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitItinerary(s)
	}
}

func (p *itnParser) Itinerary() (localctx IItineraryContext) {
	this := p
	_ = this

	localctx = NewItineraryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, itnParserRULE_itinerary)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == itnParserNUM {
		{
			p.SetState(12)
			p.Line()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(itnParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Longitude() ILongitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILongitudeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILongitudeContext)
}

func (s *LineContext) Latitude() ILatitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILatitudeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILatitudeContext)
}

func (s *LineContext) Descr() IDescrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescrContext)
}

func (s *LineContext) Flag() IFlagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlagContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *itnParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, itnParserRULE_line)

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
		p.Longitude()
	}
	{
		p.SetState(21)
		p.Match(itnParserT__0)
	}
	{
		p.SetState(22)
		p.Latitude()
	}
	{
		p.SetState(23)
		p.Match(itnParserT__0)
	}
	{
		p.SetState(24)
		p.Descr()
	}
	{
		p.SetState(25)
		p.Match(itnParserT__0)
	}
	{
		p.SetState(26)
		p.Flag()
	}
	{
		p.SetState(27)
		p.Match(itnParserT__0)
	}

	return localctx
}

// ILongitudeContext is an interface to support dynamic dispatch.
type ILongitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLongitudeContext differentiates from other interfaces.
	IsLongitudeContext()
}

type LongitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongitudeContext() *LongitudeContext {
	var p = new(LongitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_longitude
	return p
}

func (*LongitudeContext) IsLongitudeContext() {}

func NewLongitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongitudeContext {
	var p = new(LongitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_longitude

	return p
}

func (s *LongitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LongitudeContext) NUM() antlr.TerminalNode {
	return s.GetToken(itnParserNUM, 0)
}

func (s *LongitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterLongitude(s)
	}
}

func (s *LongitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitLongitude(s)
	}
}

func (p *itnParser) Longitude() (localctx ILongitudeContext) {
	this := p
	_ = this

	localctx = NewLongitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, itnParserRULE_longitude)

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
		p.SetState(29)
		p.Match(itnParserNUM)
	}

	return localctx
}

// ILatitudeContext is an interface to support dynamic dispatch.
type ILatitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLatitudeContext differentiates from other interfaces.
	IsLatitudeContext()
}

type LatitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatitudeContext() *LatitudeContext {
	var p = new(LatitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_latitude
	return p
}

func (*LatitudeContext) IsLatitudeContext() {}

func NewLatitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatitudeContext {
	var p = new(LatitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_latitude

	return p
}

func (s *LatitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LatitudeContext) NUM() antlr.TerminalNode {
	return s.GetToken(itnParserNUM, 0)
}

func (s *LatitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterLatitude(s)
	}
}

func (s *LatitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitLatitude(s)
	}
}

func (p *itnParser) Latitude() (localctx ILatitudeContext) {
	this := p
	_ = this

	localctx = NewLatitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, itnParserRULE_latitude)

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
		p.SetState(31)
		p.Match(itnParserNUM)
	}

	return localctx
}

// IDescrContext is an interface to support dynamic dispatch.
type IDescrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescrContext differentiates from other interfaces.
	IsDescrContext()
}

type DescrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescrContext() *DescrContext {
	var p = new(DescrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_descr
	return p
}

func (*DescrContext) IsDescrContext() {}

func NewDescrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescrContext {
	var p = new(DescrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_descr

	return p
}

func (s *DescrContext) GetParser() antlr.Parser { return s.parser }

func (s *DescrContext) TEXT() antlr.TerminalNode {
	return s.GetToken(itnParserTEXT, 0)
}

func (s *DescrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterDescr(s)
	}
}

func (s *DescrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitDescr(s)
	}
}

func (p *itnParser) Descr() (localctx IDescrContext) {
	this := p
	_ = this

	localctx = NewDescrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, itnParserRULE_descr)

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
		p.SetState(33)
		p.Match(itnParserTEXT)
	}

	return localctx
}

// IFlagContext is an interface to support dynamic dispatch.
type IFlagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlagContext differentiates from other interfaces.
	IsFlagContext()
}

type FlagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlagContext() *FlagContext {
	var p = new(FlagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = itnParserRULE_flag
	return p
}

func (*FlagContext) IsFlagContext() {}

func NewFlagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlagContext {
	var p = new(FlagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = itnParserRULE_flag

	return p
}

func (s *FlagContext) GetParser() antlr.Parser { return s.parser }

func (s *FlagContext) NUM() antlr.TerminalNode {
	return s.GetToken(itnParserNUM, 0)
}

func (s *FlagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.EnterFlag(s)
	}
}

func (s *FlagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(itnListener); ok {
		listenerT.ExitFlag(s)
	}
}

func (p *itnParser) Flag() (localctx IFlagContext) {
	this := p
	_ = this

	localctx = NewFlagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, itnParserRULE_flag)

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
		p.SetState(35)
		p.Match(itnParserNUM)
	}

	return localctx
}
