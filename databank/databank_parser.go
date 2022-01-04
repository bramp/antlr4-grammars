// Code generated from databank.g4 by ANTLR 4.9.3. DO NOT EDIT.

package databank // databank
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 5,
	2, 25, 10, 2, 3, 2, 6, 2, 28, 10, 2, 13, 2, 14, 2, 29, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	5, 7, 47, 10, 7, 3, 7, 5, 7, 50, 10, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4,
	6, 8, 10, 12, 14, 2, 3, 3, 2, 3, 5, 2, 51, 2, 19, 3, 2, 2, 2, 4, 31, 3,
	2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12,
	46, 3, 2, 2, 2, 14, 51, 3, 2, 2, 2, 16, 18, 7, 9, 2, 2, 17, 16, 3, 2, 2,
	2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 24,
	3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 25, 5, 4, 3, 2, 23, 25, 5, 6, 4, 2,
	24, 22, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 28, 5,
	12, 7, 2, 27, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 3, 3, 2, 2, 2, 31, 32, 5, 8, 5, 2, 32, 33, 5, 10, 6,
	2, 33, 34, 5, 10, 6, 2, 34, 5, 3, 2, 2, 2, 35, 36, 5, 10, 6, 2, 36, 37,
	5, 10, 6, 2, 37, 7, 3, 2, 2, 2, 38, 39, 9, 2, 2, 2, 39, 40, 7, 9, 2, 2,
	40, 9, 3, 2, 2, 2, 41, 42, 5, 14, 8, 2, 42, 43, 7, 9, 2, 2, 43, 11, 3,
	2, 2, 2, 44, 47, 5, 14, 8, 2, 45, 47, 7, 6, 2, 2, 46, 44, 3, 2, 2, 2, 46,
	45, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 50, 7, 9, 2, 2, 49, 48, 3, 2, 2,
	2, 49, 50, 3, 2, 2, 2, 50, 13, 3, 2, 2, 2, 51, 52, 7, 7, 2, 2, 52, 15,
	3, 2, 2, 2, 7, 19, 24, 29, 46, 49,
}
var literalNames = []string{
	"", "'-1'", "'-4'", "'-12'", "'NA'",
}
var symbolicNames = []string{
	"", "", "", "", "", "FLOATINGPOINT", "COMMENT", "EOL", "WS",
}

var ruleNames = []string{
	"databank", "datedseries", "undatedseries", "datatype", "dateline", "sample",
	"number",
}

type databankParser struct {
	*antlr.BaseParser
}

// NewdatabankParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *databankParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdatabankParser(input antlr.TokenStream) *databankParser {
	this := new(databankParser)
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
	this.GrammarFileName = "databank.g4"

	return this
}

// databankParser tokens.
const (
	databankParserEOF           = antlr.TokenEOF
	databankParserT__0          = 1
	databankParserT__1          = 2
	databankParserT__2          = 3
	databankParserT__3          = 4
	databankParserFLOATINGPOINT = 5
	databankParserCOMMENT       = 6
	databankParserEOL           = 7
	databankParserWS            = 8
)

// databankParser rules.
const (
	databankParserRULE_databank      = 0
	databankParserRULE_datedseries   = 1
	databankParserRULE_undatedseries = 2
	databankParserRULE_datatype      = 3
	databankParserRULE_dateline      = 4
	databankParserRULE_sample        = 5
	databankParserRULE_number        = 6
)

// IDatabankContext is an interface to support dynamic dispatch.
type IDatabankContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabankContext differentiates from other interfaces.
	IsDatabankContext()
}

type DatabankContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabankContext() *DatabankContext {
	var p = new(DatabankContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_databank
	return p
}

func (*DatabankContext) IsDatabankContext() {}

func NewDatabankContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabankContext {
	var p = new(DatabankContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_databank

	return p
}

func (s *DatabankContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabankContext) Datedseries() IDatedseriesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatedseriesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatedseriesContext)
}

func (s *DatabankContext) Undatedseries() IUndatedseriesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUndatedseriesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUndatedseriesContext)
}

func (s *DatabankContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(databankParserEOL)
}

func (s *DatabankContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(databankParserEOL, i)
}

func (s *DatabankContext) AllSample() []ISampleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISampleContext)(nil)).Elem())
	var tst = make([]ISampleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISampleContext)
		}
	}

	return tst
}

func (s *DatabankContext) Sample(i int) ISampleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISampleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISampleContext)
}

func (s *DatabankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabankContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterDatabank(s)
	}
}

func (s *DatabankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitDatabank(s)
	}
}

func (p *databankParser) Databank() (localctx IDatabankContext) {
	this := p
	_ = this

	localctx = NewDatabankContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, databankParserRULE_databank)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == databankParserEOL {
		{
			p.SetState(14)
			p.Match(databankParserEOL)
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case databankParserT__0, databankParserT__1, databankParserT__2:
		{
			p.SetState(20)
			p.Datedseries()
		}

	case databankParserFLOATINGPOINT:
		{
			p.SetState(21)
			p.Undatedseries()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == databankParserT__3 || _la == databankParserFLOATINGPOINT {
		{
			p.SetState(24)
			p.Sample()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDatedseriesContext is an interface to support dynamic dispatch.
type IDatedseriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatedseriesContext differentiates from other interfaces.
	IsDatedseriesContext()
}

type DatedseriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatedseriesContext() *DatedseriesContext {
	var p = new(DatedseriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_datedseries
	return p
}

func (*DatedseriesContext) IsDatedseriesContext() {}

func NewDatedseriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatedseriesContext {
	var p = new(DatedseriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_datedseries

	return p
}

func (s *DatedseriesContext) GetParser() antlr.Parser { return s.parser }

func (s *DatedseriesContext) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *DatedseriesContext) AllDateline() []IDatelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatelineContext)(nil)).Elem())
	var tst = make([]IDatelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatelineContext)
		}
	}

	return tst
}

func (s *DatedseriesContext) Dateline(i int) IDatelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatelineContext)
}

func (s *DatedseriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatedseriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatedseriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterDatedseries(s)
	}
}

func (s *DatedseriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitDatedseries(s)
	}
}

func (p *databankParser) Datedseries() (localctx IDatedseriesContext) {
	this := p
	_ = this

	localctx = NewDatedseriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, databankParserRULE_datedseries)

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
		p.Datatype()
	}
	{
		p.SetState(30)
		p.Dateline()
	}
	{
		p.SetState(31)
		p.Dateline()
	}

	return localctx
}

// IUndatedseriesContext is an interface to support dynamic dispatch.
type IUndatedseriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUndatedseriesContext differentiates from other interfaces.
	IsUndatedseriesContext()
}

type UndatedseriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndatedseriesContext() *UndatedseriesContext {
	var p = new(UndatedseriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_undatedseries
	return p
}

func (*UndatedseriesContext) IsUndatedseriesContext() {}

func NewUndatedseriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UndatedseriesContext {
	var p = new(UndatedseriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_undatedseries

	return p
}

func (s *UndatedseriesContext) GetParser() antlr.Parser { return s.parser }

func (s *UndatedseriesContext) AllDateline() []IDatelineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatelineContext)(nil)).Elem())
	var tst = make([]IDatelineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatelineContext)
		}
	}

	return tst
}

func (s *UndatedseriesContext) Dateline(i int) IDatelineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatelineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatelineContext)
}

func (s *UndatedseriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndatedseriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UndatedseriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterUndatedseries(s)
	}
}

func (s *UndatedseriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitUndatedseries(s)
	}
}

func (p *databankParser) Undatedseries() (localctx IUndatedseriesContext) {
	this := p
	_ = this

	localctx = NewUndatedseriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, databankParserRULE_undatedseries)

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
		p.Dateline()
	}
	{
		p.SetState(34)
		p.Dateline()
	}

	return localctx
}

// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) EOL() antlr.TerminalNode {
	return s.GetToken(databankParserEOL, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (p *databankParser) Datatype() (localctx IDatatypeContext) {
	this := p
	_ = this

	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, databankParserRULE_datatype)
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
		p.SetState(36)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<databankParserT__0)|(1<<databankParserT__1)|(1<<databankParserT__2))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(37)
		p.Match(databankParserEOL)
	}

	return localctx
}

// IDatelineContext is an interface to support dynamic dispatch.
type IDatelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatelineContext differentiates from other interfaces.
	IsDatelineContext()
}

type DatelineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatelineContext() *DatelineContext {
	var p = new(DatelineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_dateline
	return p
}

func (*DatelineContext) IsDatelineContext() {}

func NewDatelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatelineContext {
	var p = new(DatelineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_dateline

	return p
}

func (s *DatelineContext) GetParser() antlr.Parser { return s.parser }

func (s *DatelineContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *DatelineContext) EOL() antlr.TerminalNode {
	return s.GetToken(databankParserEOL, 0)
}

func (s *DatelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterDateline(s)
	}
}

func (s *DatelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitDateline(s)
	}
}

func (p *databankParser) Dateline() (localctx IDatelineContext) {
	this := p
	_ = this

	localctx = NewDatelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, databankParserRULE_dateline)

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
		p.SetState(39)
		p.Number()
	}
	{
		p.SetState(40)
		p.Match(databankParserEOL)
	}

	return localctx
}

// ISampleContext is an interface to support dynamic dispatch.
type ISampleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSampleContext differentiates from other interfaces.
	IsSampleContext()
}

type SampleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySampleContext() *SampleContext {
	var p = new(SampleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_sample
	return p
}

func (*SampleContext) IsSampleContext() {}

func NewSampleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SampleContext {
	var p = new(SampleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_sample

	return p
}

func (s *SampleContext) GetParser() antlr.Parser { return s.parser }

func (s *SampleContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *SampleContext) EOL() antlr.TerminalNode {
	return s.GetToken(databankParserEOL, 0)
}

func (s *SampleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SampleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SampleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterSample(s)
	}
}

func (s *SampleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitSample(s)
	}
}

func (p *databankParser) Sample() (localctx ISampleContext) {
	this := p
	_ = this

	localctx = NewSampleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, databankParserRULE_sample)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case databankParserFLOATINGPOINT:
		{
			p.SetState(42)
			p.Number()
		}

	case databankParserT__3:
		{
			p.SetState(43)
			p.Match(databankParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == databankParserEOL {
		{
			p.SetState(46)
			p.Match(databankParserEOL)
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = databankParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = databankParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FLOATINGPOINT() antlr.TerminalNode {
	return s.GetToken(databankParserFLOATINGPOINT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(databankListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *databankParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, databankParserRULE_number)

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
		p.SetState(49)
		p.Match(databankParserFLOATINGPOINT)
	}

	return localctx
}
