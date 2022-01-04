// Code generated from datetime.g4 by ANTLR 4.9.3. DO NOT EDIT.

package datetime // datetime
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 78, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 5, 2,
	26, 10, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 6, 4, 34, 10, 4, 13, 4,
	14, 4, 35, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 51, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 66, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 2, 2, 12,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 6, 3, 2, 4, 10, 3, 2, 11, 22, 3,
	2, 34, 35, 3, 2, 36, 37, 2, 81, 2, 25, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6,
	33, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 42, 3, 2, 2, 2, 12, 45, 3, 2, 2,
	2, 14, 65, 3, 2, 2, 2, 16, 67, 3, 2, 2, 2, 18, 70, 3, 2, 2, 2, 20, 75,
	3, 2, 2, 2, 22, 23, 5, 4, 3, 2, 23, 24, 7, 3, 2, 2, 24, 26, 3, 2, 2, 2,
	25, 22, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 28, 5,
	6, 4, 2, 28, 29, 5, 10, 6, 2, 29, 3, 3, 2, 2, 2, 30, 31, 9, 2, 2, 2, 31,
	5, 3, 2, 2, 2, 32, 34, 5, 16, 9, 2, 33, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2,
	2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38,
	5, 8, 5, 2, 38, 39, 5, 16, 9, 2, 39, 7, 3, 2, 2, 2, 40, 41, 9, 3, 2, 2,
	41, 9, 3, 2, 2, 2, 42, 43, 5, 12, 7, 2, 43, 44, 5, 14, 8, 2, 44, 11, 3,
	2, 2, 2, 45, 46, 5, 16, 9, 2, 46, 47, 7, 23, 2, 2, 47, 50, 5, 16, 9, 2,
	48, 49, 7, 23, 2, 2, 49, 51, 5, 16, 9, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3,
	2, 2, 2, 51, 13, 3, 2, 2, 2, 52, 66, 7, 24, 2, 2, 53, 66, 7, 25, 2, 2,
	54, 66, 7, 26, 2, 2, 55, 66, 7, 27, 2, 2, 56, 66, 7, 28, 2, 2, 57, 66,
	7, 29, 2, 2, 58, 66, 7, 30, 2, 2, 59, 66, 7, 31, 2, 2, 60, 66, 7, 32, 2,
	2, 61, 66, 7, 33, 2, 2, 62, 66, 7, 36, 2, 2, 63, 64, 9, 4, 2, 2, 64, 66,
	5, 18, 10, 2, 65, 52, 3, 2, 2, 2, 65, 53, 3, 2, 2, 2, 65, 54, 3, 2, 2,
	2, 65, 55, 3, 2, 2, 2, 65, 56, 3, 2, 2, 2, 65, 57, 3, 2, 2, 2, 65, 58,
	3, 2, 2, 2, 65, 59, 3, 2, 2, 2, 65, 60, 3, 2, 2, 2, 65, 61, 3, 2, 2, 2,
	65, 62, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 15, 3, 2, 2, 2, 67, 68, 5,
	20, 11, 2, 68, 69, 5, 20, 11, 2, 69, 17, 3, 2, 2, 2, 70, 71, 5, 20, 11,
	2, 71, 72, 5, 20, 11, 2, 72, 73, 5, 20, 11, 2, 73, 74, 5, 20, 11, 2, 74,
	19, 3, 2, 2, 2, 75, 76, 9, 5, 2, 2, 76, 21, 3, 2, 2, 2, 6, 25, 35, 50,
	65,
}
var literalNames = []string{
	"", "','", "'Mon'", "'Tue'", "'Wed'", "'Thu'", "'Fri'", "'Sat'", "'Sun'",
	"'Jan'", "'Feb'", "'Mar'", "'Apr'", "'May'", "'Jun'", "'Jul'", "'Aug'",
	"'Sep'", "'Oct'", "'Nov'", "'Dec'", "':'", "'UT'", "'GMT'", "'EST'", "'EDT'",
	"'CST'", "'CDT'", "'MST'", "'MDT'", "'PST'", "'PDT'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ALPHA",
	"DIGIT", "WS",
}

var ruleNames = []string{
	"date_time", "day", "date", "month", "time", "hour", "zone", "two_digit",
	"four_digit", "alphanumeric",
}

type datetimeParser struct {
	*antlr.BaseParser
}

// NewdatetimeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *datetimeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdatetimeParser(input antlr.TokenStream) *datetimeParser {
	this := new(datetimeParser)
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
	this.GrammarFileName = "datetime.g4"

	return this
}

// datetimeParser tokens.
const (
	datetimeParserEOF   = antlr.TokenEOF
	datetimeParserT__0  = 1
	datetimeParserT__1  = 2
	datetimeParserT__2  = 3
	datetimeParserT__3  = 4
	datetimeParserT__4  = 5
	datetimeParserT__5  = 6
	datetimeParserT__6  = 7
	datetimeParserT__7  = 8
	datetimeParserT__8  = 9
	datetimeParserT__9  = 10
	datetimeParserT__10 = 11
	datetimeParserT__11 = 12
	datetimeParserT__12 = 13
	datetimeParserT__13 = 14
	datetimeParserT__14 = 15
	datetimeParserT__15 = 16
	datetimeParserT__16 = 17
	datetimeParserT__17 = 18
	datetimeParserT__18 = 19
	datetimeParserT__19 = 20
	datetimeParserT__20 = 21
	datetimeParserT__21 = 22
	datetimeParserT__22 = 23
	datetimeParserT__23 = 24
	datetimeParserT__24 = 25
	datetimeParserT__25 = 26
	datetimeParserT__26 = 27
	datetimeParserT__27 = 28
	datetimeParserT__28 = 29
	datetimeParserT__29 = 30
	datetimeParserT__30 = 31
	datetimeParserT__31 = 32
	datetimeParserT__32 = 33
	datetimeParserALPHA = 34
	datetimeParserDIGIT = 35
	datetimeParserWS    = 36
)

// datetimeParser rules.
const (
	datetimeParserRULE_date_time    = 0
	datetimeParserRULE_day          = 1
	datetimeParserRULE_date         = 2
	datetimeParserRULE_month        = 3
	datetimeParserRULE_time         = 4
	datetimeParserRULE_hour         = 5
	datetimeParserRULE_zone         = 6
	datetimeParserRULE_two_digit    = 7
	datetimeParserRULE_four_digit   = 8
	datetimeParserRULE_alphanumeric = 9
)

// IDate_timeContext is an interface to support dynamic dispatch.
type IDate_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_timeContext differentiates from other interfaces.
	IsDate_timeContext()
}

type Date_timeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_timeContext() *Date_timeContext {
	var p = new(Date_timeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_date_time
	return p
}

func (*Date_timeContext) IsDate_timeContext() {}

func NewDate_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_timeContext {
	var p = new(Date_timeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_date_time

	return p
}

func (s *Date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_timeContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *Date_timeContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *Date_timeContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *Date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterDate_time(s)
	}
}

func (s *Date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitDate_time(s)
	}
}

func (p *datetimeParser) Date_time() (localctx IDate_timeContext) {
	this := p
	_ = this

	localctx = NewDate_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, datetimeParserRULE_date_time)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<datetimeParserT__1)|(1<<datetimeParserT__2)|(1<<datetimeParserT__3)|(1<<datetimeParserT__4)|(1<<datetimeParserT__5)|(1<<datetimeParserT__6)|(1<<datetimeParserT__7))) != 0 {
		{
			p.SetState(20)
			p.Day()
		}
		{
			p.SetState(21)
			p.Match(datetimeParserT__0)
		}

	}
	{
		p.SetState(25)
		p.Date()
	}
	{
		p.SetState(26)
		p.Time()
	}

	return localctx
}

// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_day
	return p
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }
func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitDay(s)
	}
}

func (p *datetimeParser) Day() (localctx IDayContext) {
	this := p
	_ = this

	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, datetimeParserRULE_day)
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
		p.SetState(28)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<datetimeParserT__1)|(1<<datetimeParserT__2)|(1<<datetimeParserT__3)|(1<<datetimeParserT__4)|(1<<datetimeParserT__5)|(1<<datetimeParserT__6)|(1<<datetimeParserT__7))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *DateContext) AllTwo_digit() []ITwo_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITwo_digitContext)(nil)).Elem())
	var tst = make([]ITwo_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITwo_digitContext)
		}
	}

	return tst
}

func (s *DateContext) Two_digit(i int) ITwo_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwo_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITwo_digitContext)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *datetimeParser) Date() (localctx IDateContext) {
	this := p
	_ = this

	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, datetimeParserRULE_date)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == datetimeParserALPHA || _la == datetimeParserDIGIT {
		{
			p.SetState(30)
			p.Two_digit()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Month()
	}
	{
		p.SetState(36)
		p.Two_digit()
	}

	return localctx
}

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_month
	return p
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }
func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitMonth(s)
	}
}

func (p *datetimeParser) Month() (localctx IMonthContext) {
	this := p
	_ = this

	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, datetimeParserRULE_month)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<datetimeParserT__8)|(1<<datetimeParserT__9)|(1<<datetimeParserT__10)|(1<<datetimeParserT__11)|(1<<datetimeParserT__12)|(1<<datetimeParserT__13)|(1<<datetimeParserT__14)|(1<<datetimeParserT__15)|(1<<datetimeParserT__16)|(1<<datetimeParserT__17)|(1<<datetimeParserT__18)|(1<<datetimeParserT__19))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *TimeContext) Zone() IZoneContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZoneContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZoneContext)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitTime(s)
	}
}

func (p *datetimeParser) Time() (localctx ITimeContext) {
	this := p
	_ = this

	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, datetimeParserRULE_time)

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
		p.SetState(40)
		p.Hour()
	}
	{
		p.SetState(41)
		p.Zone()
	}

	return localctx
}

// IHourContext is an interface to support dynamic dispatch.
type IHourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHourContext differentiates from other interfaces.
	IsHourContext()
}

type HourContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHourContext() *HourContext {
	var p = new(HourContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_hour
	return p
}

func (*HourContext) IsHourContext() {}

func NewHourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HourContext {
	var p = new(HourContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_hour

	return p
}

func (s *HourContext) GetParser() antlr.Parser { return s.parser }

func (s *HourContext) AllTwo_digit() []ITwo_digitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITwo_digitContext)(nil)).Elem())
	var tst = make([]ITwo_digitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITwo_digitContext)
		}
	}

	return tst
}

func (s *HourContext) Two_digit(i int) ITwo_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwo_digitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITwo_digitContext)
}

func (s *HourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterHour(s)
	}
}

func (s *HourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitHour(s)
	}
}

func (p *datetimeParser) Hour() (localctx IHourContext) {
	this := p
	_ = this

	localctx = NewHourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, datetimeParserRULE_hour)
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
		p.SetState(43)
		p.Two_digit()
	}
	{
		p.SetState(44)
		p.Match(datetimeParserT__20)
	}
	{
		p.SetState(45)
		p.Two_digit()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == datetimeParserT__20 {
		{
			p.SetState(46)
			p.Match(datetimeParserT__20)
		}
		{
			p.SetState(47)
			p.Two_digit()
		}

	}

	return localctx
}

// IZoneContext is an interface to support dynamic dispatch.
type IZoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZoneContext differentiates from other interfaces.
	IsZoneContext()
}

type ZoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZoneContext() *ZoneContext {
	var p = new(ZoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_zone
	return p
}

func (*ZoneContext) IsZoneContext() {}

func NewZoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZoneContext {
	var p = new(ZoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_zone

	return p
}

func (s *ZoneContext) GetParser() antlr.Parser { return s.parser }

func (s *ZoneContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(datetimeParserALPHA, 0)
}

func (s *ZoneContext) Four_digit() IFour_digitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFour_digitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFour_digitContext)
}

func (s *ZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterZone(s)
	}
}

func (s *ZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitZone(s)
	}
}

func (p *datetimeParser) Zone() (localctx IZoneContext) {
	this := p
	_ = this

	localctx = NewZoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, datetimeParserRULE_zone)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case datetimeParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(datetimeParserT__21)
		}

	case datetimeParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Match(datetimeParserT__22)
		}

	case datetimeParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Match(datetimeParserT__23)
		}

	case datetimeParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Match(datetimeParserT__24)
		}

	case datetimeParserT__25:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.Match(datetimeParserT__25)
		}

	case datetimeParserT__26:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(55)
			p.Match(datetimeParserT__26)
		}

	case datetimeParserT__27:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(56)
			p.Match(datetimeParserT__27)
		}

	case datetimeParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(57)
			p.Match(datetimeParserT__28)
		}

	case datetimeParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(58)
			p.Match(datetimeParserT__29)
		}

	case datetimeParserT__30:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(59)
			p.Match(datetimeParserT__30)
		}

	case datetimeParserALPHA:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(60)
			p.Match(datetimeParserALPHA)
		}

	case datetimeParserT__31, datetimeParserT__32:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(61)
			_la = p.GetTokenStream().LA(1)

			if !(_la == datetimeParserT__31 || _la == datetimeParserT__32) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(62)
			p.Four_digit()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITwo_digitContext is an interface to support dynamic dispatch.
type ITwo_digitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTwo_digitContext differentiates from other interfaces.
	IsTwo_digitContext()
}

type Two_digitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTwo_digitContext() *Two_digitContext {
	var p = new(Two_digitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_two_digit
	return p
}

func (*Two_digitContext) IsTwo_digitContext() {}

func NewTwo_digitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Two_digitContext {
	var p = new(Two_digitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_two_digit

	return p
}

func (s *Two_digitContext) GetParser() antlr.Parser { return s.parser }

func (s *Two_digitContext) AllAlphanumeric() []IAlphanumericContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlphanumericContext)(nil)).Elem())
	var tst = make([]IAlphanumericContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlphanumericContext)
		}
	}

	return tst
}

func (s *Two_digitContext) Alphanumeric(i int) IAlphanumericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphanumericContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlphanumericContext)
}

func (s *Two_digitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Two_digitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Two_digitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterTwo_digit(s)
	}
}

func (s *Two_digitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitTwo_digit(s)
	}
}

func (p *datetimeParser) Two_digit() (localctx ITwo_digitContext) {
	this := p
	_ = this

	localctx = NewTwo_digitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, datetimeParserRULE_two_digit)

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
		p.SetState(65)
		p.Alphanumeric()
	}
	{
		p.SetState(66)
		p.Alphanumeric()
	}

	return localctx
}

// IFour_digitContext is an interface to support dynamic dispatch.
type IFour_digitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFour_digitContext differentiates from other interfaces.
	IsFour_digitContext()
}

type Four_digitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFour_digitContext() *Four_digitContext {
	var p = new(Four_digitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_four_digit
	return p
}

func (*Four_digitContext) IsFour_digitContext() {}

func NewFour_digitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Four_digitContext {
	var p = new(Four_digitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_four_digit

	return p
}

func (s *Four_digitContext) GetParser() antlr.Parser { return s.parser }

func (s *Four_digitContext) AllAlphanumeric() []IAlphanumericContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlphanumericContext)(nil)).Elem())
	var tst = make([]IAlphanumericContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlphanumericContext)
		}
	}

	return tst
}

func (s *Four_digitContext) Alphanumeric(i int) IAlphanumericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphanumericContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlphanumericContext)
}

func (s *Four_digitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Four_digitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Four_digitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterFour_digit(s)
	}
}

func (s *Four_digitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitFour_digit(s)
	}
}

func (p *datetimeParser) Four_digit() (localctx IFour_digitContext) {
	this := p
	_ = this

	localctx = NewFour_digitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, datetimeParserRULE_four_digit)

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
		p.Alphanumeric()
	}
	{
		p.SetState(69)
		p.Alphanumeric()
	}
	{
		p.SetState(70)
		p.Alphanumeric()
	}
	{
		p.SetState(71)
		p.Alphanumeric()
	}

	return localctx
}

// IAlphanumericContext is an interface to support dynamic dispatch.
type IAlphanumericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlphanumericContext differentiates from other interfaces.
	IsAlphanumericContext()
}

type AlphanumericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlphanumericContext() *AlphanumericContext {
	var p = new(AlphanumericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = datetimeParserRULE_alphanumeric
	return p
}

func (*AlphanumericContext) IsAlphanumericContext() {}

func NewAlphanumericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlphanumericContext {
	var p = new(AlphanumericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = datetimeParserRULE_alphanumeric

	return p
}

func (s *AlphanumericContext) GetParser() antlr.Parser { return s.parser }

func (s *AlphanumericContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(datetimeParserALPHA, 0)
}

func (s *AlphanumericContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(datetimeParserDIGIT, 0)
}

func (s *AlphanumericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlphanumericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlphanumericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.EnterAlphanumeric(s)
	}
}

func (s *AlphanumericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(datetimeListener); ok {
		listenerT.ExitAlphanumeric(s)
	}
}

func (p *datetimeParser) Alphanumeric() (localctx IAlphanumericContext) {
	this := p
	_ = this

	localctx = NewAlphanumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, datetimeParserRULE_alphanumeric)
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

		if !(_la == datetimeParserALPHA || _la == datetimeParserDIGIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
