// Generated from metric.g4 by ANTLR 4.7.

package metric // metric
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 45, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	3, 5, 3, 26, 10, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 5, 6, 39, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2,
	9, 2, 4, 6, 8, 10, 12, 14, 2, 6, 3, 2, 3, 4, 3, 2, 6, 21, 4, 2, 16, 16,
	22, 27, 4, 2, 8, 8, 28, 48, 2, 41, 2, 16, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2,
	6, 31, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 38, 3, 2, 2, 2, 12, 40, 3, 2,
	2, 2, 14, 42, 3, 2, 2, 2, 16, 21, 5, 4, 3, 2, 17, 18, 9, 2, 2, 2, 18, 20,
	5, 4, 3, 2, 19, 17, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2,
	21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 26, 5, 8,
	5, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 29,
	5, 10, 6, 2, 28, 30, 5, 6, 4, 2, 29, 28, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2,
	30, 5, 3, 2, 2, 2, 31, 32, 7, 5, 2, 2, 32, 33, 7, 49, 2, 2, 33, 7, 3, 2,
	2, 2, 34, 35, 9, 3, 2, 2, 35, 9, 3, 2, 2, 2, 36, 39, 5, 12, 7, 2, 37, 39,
	5, 14, 8, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 11, 3, 2, 2, 2,
	40, 41, 9, 4, 2, 2, 41, 13, 3, 2, 2, 2, 42, 43, 9, 5, 2, 2, 43, 15, 3,
	2, 2, 2, 6, 21, 25, 29, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'/'", "'^'", "'E'", "'P'", "'T'", "'G'", "'M'", "'k'", "'h'",
	"'da'", "'d'", "'c'", "'m'", "'\u00B5'", "'n'", "'p'", "'f'", "'a'", "'g'",
	"'s'", "'A'", "'K'", "'mol'", "'cd'", "'rad'", "'Hz'", "'sr'", "'N'", "'Pa'",
	"'J'", "'W'", "'C'", "'V'", "'F'", "'\u03A9'", "'S'", "'Wb'", "'H'", "'\u02DAC'",
	"'lm'", "'lx'", "'Bq'", "'Gy'", "'Sv'", "'kat'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "INTE", "WS",
}

var ruleNames = []string{
	"uom", "measure", "exponent", "prefix", "unit", "baseunit", "derivedunit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type metricParser struct {
	*antlr.BaseParser
}

func NewmetricParser(input antlr.TokenStream) *metricParser {
	this := new(metricParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "metric.g4"

	return this
}

// metricParser tokens.
const (
	metricParserEOF   = antlr.TokenEOF
	metricParserT__0  = 1
	metricParserT__1  = 2
	metricParserT__2  = 3
	metricParserT__3  = 4
	metricParserT__4  = 5
	metricParserT__5  = 6
	metricParserT__6  = 7
	metricParserT__7  = 8
	metricParserT__8  = 9
	metricParserT__9  = 10
	metricParserT__10 = 11
	metricParserT__11 = 12
	metricParserT__12 = 13
	metricParserT__13 = 14
	metricParserT__14 = 15
	metricParserT__15 = 16
	metricParserT__16 = 17
	metricParserT__17 = 18
	metricParserT__18 = 19
	metricParserT__19 = 20
	metricParserT__20 = 21
	metricParserT__21 = 22
	metricParserT__22 = 23
	metricParserT__23 = 24
	metricParserT__24 = 25
	metricParserT__25 = 26
	metricParserT__26 = 27
	metricParserT__27 = 28
	metricParserT__28 = 29
	metricParserT__29 = 30
	metricParserT__30 = 31
	metricParserT__31 = 32
	metricParserT__32 = 33
	metricParserT__33 = 34
	metricParserT__34 = 35
	metricParserT__35 = 36
	metricParserT__36 = 37
	metricParserT__37 = 38
	metricParserT__38 = 39
	metricParserT__39 = 40
	metricParserT__40 = 41
	metricParserT__41 = 42
	metricParserT__42 = 43
	metricParserT__43 = 44
	metricParserT__44 = 45
	metricParserT__45 = 46
	metricParserINTE  = 47
	metricParserWS    = 48
)

// metricParser rules.
const (
	metricParserRULE_uom         = 0
	metricParserRULE_measure     = 1
	metricParserRULE_exponent    = 2
	metricParserRULE_prefix      = 3
	metricParserRULE_unit        = 4
	metricParserRULE_baseunit    = 5
	metricParserRULE_derivedunit = 6
)

// IUomContext is an interface to support dynamic dispatch.
type IUomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUomContext differentiates from other interfaces.
	IsUomContext()
}

type UomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUomContext() *UomContext {
	var p = new(UomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_uom
	return p
}

func (*UomContext) IsUomContext() {}

func NewUomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UomContext {
	var p = new(UomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_uom

	return p
}

func (s *UomContext) GetParser() antlr.Parser { return s.parser }

func (s *UomContext) AllMeasure() []IMeasureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMeasureContext)(nil)).Elem())
	var tst = make([]IMeasureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMeasureContext)
		}
	}

	return tst
}

func (s *UomContext) Measure(i int) IMeasureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeasureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *UomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterUom(s)
	}
}

func (s *UomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitUom(s)
	}
}

func (p *metricParser) Uom() (localctx IUomContext) {
	localctx = NewUomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, metricParserRULE_uom)
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
		p.SetState(14)
		p.Measure()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == metricParserT__0 || _la == metricParserT__1 {
		p.SetState(15)
		_la = p.GetTokenStream().LA(1)

		if !(_la == metricParserT__0 || _la == metricParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(16)
			p.Measure()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMeasureContext is an interface to support dynamic dispatch.
type IMeasureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasureContext differentiates from other interfaces.
	IsMeasureContext()
}

type MeasureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasureContext() *MeasureContext {
	var p = new(MeasureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_measure
	return p
}

func (*MeasureContext) IsMeasureContext() {}

func NewMeasureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasureContext {
	var p = new(MeasureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_measure

	return p
}

func (s *MeasureContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasureContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *MeasureContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *MeasureContext) Exponent() IExponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExponentContext)
}

func (s *MeasureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterMeasure(s)
	}
}

func (s *MeasureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitMeasure(s)
	}
}

func (p *metricParser) Measure() (localctx IMeasureContext) {
	localctx = NewMeasureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, metricParserRULE_measure)
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

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(22)
			p.Prefix()
		}

	}
	{
		p.SetState(25)
		p.Unit()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == metricParserT__2 {
		{
			p.SetState(26)
			p.Exponent()
		}

	}

	return localctx
}

// IExponentContext is an interface to support dynamic dispatch.
type IExponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExponentContext differentiates from other interfaces.
	IsExponentContext()
}

type ExponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExponentContext() *ExponentContext {
	var p = new(ExponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_exponent
	return p
}

func (*ExponentContext) IsExponentContext() {}

func NewExponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExponentContext {
	var p = new(ExponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_exponent

	return p
}

func (s *ExponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ExponentContext) INTE() antlr.TerminalNode {
	return s.GetToken(metricParserINTE, 0)
}

func (s *ExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterExponent(s)
	}
}

func (s *ExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitExponent(s)
	}
}

func (p *metricParser) Exponent() (localctx IExponentContext) {
	localctx = NewExponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, metricParserRULE_exponent)

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
		p.Match(metricParserT__2)
	}
	{
		p.SetState(30)
		p.Match(metricParserINTE)
	}

	return localctx
}

// IPrefixContext is an interface to support dynamic dispatch.
type IPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixContext differentiates from other interfaces.
	IsPrefixContext()
}

type PrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixContext() *PrefixContext {
	var p = new(PrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_prefix
	return p
}

func (*PrefixContext) IsPrefixContext() {}

func NewPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixContext {
	var p = new(PrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_prefix

	return p
}

func (s *PrefixContext) GetParser() antlr.Parser { return s.parser }
func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (p *metricParser) Prefix() (localctx IPrefixContext) {
	localctx = NewPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, metricParserRULE_prefix)
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
	p.SetState(32)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metricParserT__3)|(1<<metricParserT__4)|(1<<metricParserT__5)|(1<<metricParserT__6)|(1<<metricParserT__7)|(1<<metricParserT__8)|(1<<metricParserT__9)|(1<<metricParserT__10)|(1<<metricParserT__11)|(1<<metricParserT__12)|(1<<metricParserT__13)|(1<<metricParserT__14)|(1<<metricParserT__15)|(1<<metricParserT__16)|(1<<metricParserT__17)|(1<<metricParserT__18))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) Baseunit() IBaseunitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseunitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseunitContext)
}

func (s *UnitContext) Derivedunit() IDerivedunitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDerivedunitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDerivedunitContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *metricParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, metricParserRULE_unit)

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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case metricParserT__13, metricParserT__19, metricParserT__20, metricParserT__21, metricParserT__22, metricParserT__23, metricParserT__24:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Baseunit()
		}

	case metricParserT__5, metricParserT__25, metricParserT__26, metricParserT__27, metricParserT__28, metricParserT__29, metricParserT__30, metricParserT__31, metricParserT__32, metricParserT__33, metricParserT__34, metricParserT__35, metricParserT__36, metricParserT__37, metricParserT__38, metricParserT__39, metricParserT__40, metricParserT__41, metricParserT__42, metricParserT__43, metricParserT__44, metricParserT__45:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Derivedunit()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseunitContext is an interface to support dynamic dispatch.
type IBaseunitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseunitContext differentiates from other interfaces.
	IsBaseunitContext()
}

type BaseunitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseunitContext() *BaseunitContext {
	var p = new(BaseunitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_baseunit
	return p
}

func (*BaseunitContext) IsBaseunitContext() {}

func NewBaseunitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseunitContext {
	var p = new(BaseunitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_baseunit

	return p
}

func (s *BaseunitContext) GetParser() antlr.Parser { return s.parser }
func (s *BaseunitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseunitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseunitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterBaseunit(s)
	}
}

func (s *BaseunitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitBaseunit(s)
	}
}

func (p *metricParser) Baseunit() (localctx IBaseunitContext) {
	localctx = NewBaseunitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, metricParserRULE_baseunit)
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
	p.SetState(38)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metricParserT__13)|(1<<metricParserT__19)|(1<<metricParserT__20)|(1<<metricParserT__21)|(1<<metricParserT__22)|(1<<metricParserT__23)|(1<<metricParserT__24))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDerivedunitContext is an interface to support dynamic dispatch.
type IDerivedunitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDerivedunitContext differentiates from other interfaces.
	IsDerivedunitContext()
}

type DerivedunitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDerivedunitContext() *DerivedunitContext {
	var p = new(DerivedunitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metricParserRULE_derivedunit
	return p
}

func (*DerivedunitContext) IsDerivedunitContext() {}

func NewDerivedunitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DerivedunitContext {
	var p = new(DerivedunitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metricParserRULE_derivedunit

	return p
}

func (s *DerivedunitContext) GetParser() antlr.Parser { return s.parser }
func (s *DerivedunitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivedunitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DerivedunitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.EnterDerivedunit(s)
	}
}

func (s *DerivedunitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metricListener); ok {
		listenerT.ExitDerivedunit(s)
	}
}

func (p *metricParser) Derivedunit() (localctx IDerivedunitContext) {
	localctx = NewDerivedunitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, metricParserRULE_derivedunit)
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
	p.SetState(40)
	_la = p.GetTokenStream().LA(1)

	if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metricParserT__5)|(1<<metricParserT__25)|(1<<metricParserT__26)|(1<<metricParserT__27)|(1<<metricParserT__28)|(1<<metricParserT__29)|(1<<metricParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(metricParserT__31-32))|(1<<(metricParserT__32-32))|(1<<(metricParserT__33-32))|(1<<(metricParserT__34-32))|(1<<(metricParserT__35-32))|(1<<(metricParserT__36-32))|(1<<(metricParserT__37-32))|(1<<(metricParserT__38-32))|(1<<(metricParserT__39-32))|(1<<(metricParserT__40-32))|(1<<(metricParserT__41-32))|(1<<(metricParserT__42-32))|(1<<(metricParserT__43-32))|(1<<(metricParserT__44-32))|(1<<(metricParserT__45-32)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
