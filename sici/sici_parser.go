// Code generated from sici.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sici // sici
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 111,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 7, 2, 34, 10, 2, 12,
	2, 14, 2, 37, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 57, 10,
	5, 13, 5, 14, 5, 58, 3, 5, 3, 5, 3, 6, 6, 6, 64, 10, 6, 13, 6, 14, 6, 65,
	3, 6, 3, 6, 6, 6, 70, 10, 6, 13, 6, 14, 6, 71, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 6, 8, 81, 10, 8, 13, 8, 14, 8, 82, 3, 9, 6, 9, 86, 10,
	9, 13, 9, 14, 9, 87, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 2, 4, 4, 2, 5, 5, 13, 13, 4, 2, 3, 3, 13, 13, 2, 102, 2,
	30, 3, 2, 2, 2, 4, 40, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2,
	10, 63, 3, 2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 80, 3, 2, 2, 2, 16, 85, 3,
	2, 2, 2, 18, 89, 3, 2, 2, 2, 20, 98, 3, 2, 2, 2, 22, 100, 3, 2, 2, 2, 24,
	102, 3, 2, 2, 2, 26, 105, 3, 2, 2, 2, 28, 108, 3, 2, 2, 2, 30, 31, 5, 4,
	3, 2, 31, 35, 5, 12, 7, 2, 32, 34, 5, 18, 10, 2, 33, 32, 3, 2, 2, 2, 34,
	37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2,
	2, 37, 35, 3, 2, 2, 2, 38, 39, 7, 2, 2, 3, 39, 3, 3, 2, 2, 2, 40, 41, 5,
	6, 4, 2, 41, 42, 5, 8, 5, 2, 42, 43, 5, 10, 6, 2, 43, 5, 3, 2, 2, 2, 44,
	45, 7, 13, 2, 2, 45, 46, 7, 13, 2, 2, 46, 47, 7, 13, 2, 2, 47, 48, 7, 13,
	2, 2, 48, 49, 7, 3, 2, 2, 49, 50, 7, 13, 2, 2, 50, 51, 7, 13, 2, 2, 51,
	52, 7, 13, 2, 2, 52, 53, 7, 13, 2, 2, 53, 7, 3, 2, 2, 2, 54, 56, 7, 4,
	2, 2, 55, 57, 9, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2,
	61, 9, 3, 2, 2, 2, 62, 64, 7, 13, 2, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3,
	2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67,
	69, 7, 7, 2, 2, 68, 70, 7, 13, 2, 2, 69, 68, 3, 2, 2, 2, 70, 71, 3, 2,
	2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 11, 3, 2, 2, 2, 73, 74,
	7, 8, 2, 2, 74, 75, 5, 14, 8, 2, 75, 76, 7, 7, 2, 2, 76, 77, 5, 16, 9,
	2, 77, 78, 7, 9, 2, 2, 78, 13, 3, 2, 2, 2, 79, 81, 7, 13, 2, 2, 80, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 15, 3, 2, 2, 2, 84, 86, 7, 12, 2, 2, 85, 84, 3, 2, 2, 2, 86, 87, 3,
	2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 17, 3, 2, 2, 2, 89,
	90, 5, 20, 11, 2, 90, 91, 7, 10, 2, 2, 91, 92, 5, 22, 12, 2, 92, 93, 7,
	10, 2, 2, 93, 94, 5, 24, 13, 2, 94, 95, 7, 11, 2, 2, 95, 96, 5, 26, 14,
	2, 96, 97, 5, 28, 15, 2, 97, 19, 3, 2, 2, 2, 98, 99, 7, 13, 2, 2, 99, 21,
	3, 2, 2, 2, 100, 101, 7, 13, 2, 2, 101, 23, 3, 2, 2, 2, 102, 103, 7, 12,
	2, 2, 103, 104, 7, 12, 2, 2, 104, 25, 3, 2, 2, 2, 105, 106, 7, 13, 2, 2,
	106, 107, 9, 3, 2, 2, 107, 27, 3, 2, 2, 2, 108, 109, 7, 12, 2, 2, 109,
	29, 3, 2, 2, 2, 8, 35, 58, 65, 71, 82, 87,
}
var literalNames = []string{
	"", "'-'", "'('", "'/'", "')'", "':'", "'<'", "'>'", "'.'", "';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "LETTER", "DIGIT", "WS",
}

var ruleNames = []string{
	"sici", "item", "issn", "chronology", "enumeration", "contribution", "location",
	"title", "control", "csi", "dpi", "mfi", "version", "check",
}

type siciParser struct {
	*antlr.BaseParser
}

// NewsiciParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *siciParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsiciParser(input antlr.TokenStream) *siciParser {
	this := new(siciParser)
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
	this.GrammarFileName = "sici.g4"

	return this
}

// siciParser tokens.
const (
	siciParserEOF    = antlr.TokenEOF
	siciParserT__0   = 1
	siciParserT__1   = 2
	siciParserT__2   = 3
	siciParserT__3   = 4
	siciParserT__4   = 5
	siciParserT__5   = 6
	siciParserT__6   = 7
	siciParserT__7   = 8
	siciParserT__8   = 9
	siciParserLETTER = 10
	siciParserDIGIT  = 11
	siciParserWS     = 12
)

// siciParser rules.
const (
	siciParserRULE_sici         = 0
	siciParserRULE_item         = 1
	siciParserRULE_issn         = 2
	siciParserRULE_chronology   = 3
	siciParserRULE_enumeration  = 4
	siciParserRULE_contribution = 5
	siciParserRULE_location     = 6
	siciParserRULE_title        = 7
	siciParserRULE_control      = 8
	siciParserRULE_csi          = 9
	siciParserRULE_dpi          = 10
	siciParserRULE_mfi          = 11
	siciParserRULE_version      = 12
	siciParserRULE_check        = 13
)

// ISiciContext is an interface to support dynamic dispatch.
type ISiciContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSiciContext differentiates from other interfaces.
	IsSiciContext()
}

type SiciContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySiciContext() *SiciContext {
	var p = new(SiciContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_sici
	return p
}

func (*SiciContext) IsSiciContext() {}

func NewSiciContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SiciContext {
	var p = new(SiciContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_sici

	return p
}

func (s *SiciContext) GetParser() antlr.Parser { return s.parser }

func (s *SiciContext) Item() IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *SiciContext) Contribution() IContributionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContributionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContributionContext)
}

func (s *SiciContext) EOF() antlr.TerminalNode {
	return s.GetToken(siciParserEOF, 0)
}

func (s *SiciContext) AllControl() []IControlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IControlContext)(nil)).Elem())
	var tst = make([]IControlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IControlContext)
		}
	}

	return tst
}

func (s *SiciContext) Control(i int) IControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IControlContext)
}

func (s *SiciContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SiciContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SiciContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterSici(s)
	}
}

func (s *SiciContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitSici(s)
	}
}

func (p *siciParser) Sici() (localctx ISiciContext) {
	this := p
	_ = this

	localctx = NewSiciContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, siciParserRULE_sici)
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
		p.Item()
	}
	{
		p.SetState(29)
		p.Contribution()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == siciParserDIGIT {
		{
			p.SetState(30)
			p.Control()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(siciParserEOF)
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
	p.RuleIndex = siciParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) Issn() IIssnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIssnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIssnContext)
}

func (s *ItemContext) Chronology() IChronologyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChronologyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChronologyContext)
}

func (s *ItemContext) Enumeration() IEnumerationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumerationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumerationContext)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *siciParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, siciParserRULE_item)

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
		p.Issn()
	}
	{
		p.SetState(39)
		p.Chronology()
	}
	{
		p.SetState(40)
		p.Enumeration()
	}

	return localctx
}

// IIssnContext is an interface to support dynamic dispatch.
type IIssnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIssnContext differentiates from other interfaces.
	IsIssnContext()
}

type IssnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIssnContext() *IssnContext {
	var p = new(IssnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_issn
	return p
}

func (*IssnContext) IsIssnContext() {}

func NewIssnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IssnContext {
	var p = new(IssnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_issn

	return p
}

func (s *IssnContext) GetParser() antlr.Parser { return s.parser }

func (s *IssnContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(siciParserDIGIT)
}

func (s *IssnContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, i)
}

func (s *IssnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IssnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IssnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterIssn(s)
	}
}

func (s *IssnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitIssn(s)
	}
}

func (p *siciParser) Issn() (localctx IIssnContext) {
	this := p
	_ = this

	localctx = NewIssnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, siciParserRULE_issn)

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
		p.SetState(42)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(43)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(44)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(45)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(46)
		p.Match(siciParserT__0)
	}
	{
		p.SetState(47)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(48)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(49)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(50)
		p.Match(siciParserDIGIT)
	}

	return localctx
}

// IChronologyContext is an interface to support dynamic dispatch.
type IChronologyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChronologyContext differentiates from other interfaces.
	IsChronologyContext()
}

type ChronologyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChronologyContext() *ChronologyContext {
	var p = new(ChronologyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_chronology
	return p
}

func (*ChronologyContext) IsChronologyContext() {}

func NewChronologyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChronologyContext {
	var p = new(ChronologyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_chronology

	return p
}

func (s *ChronologyContext) GetParser() antlr.Parser { return s.parser }

func (s *ChronologyContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(siciParserDIGIT)
}

func (s *ChronologyContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, i)
}

func (s *ChronologyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChronologyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChronologyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterChronology(s)
	}
}

func (s *ChronologyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitChronology(s)
	}
}

func (p *siciParser) Chronology() (localctx IChronologyContext) {
	this := p
	_ = this

	localctx = NewChronologyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, siciParserRULE_chronology)
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
		p.SetState(52)
		p.Match(siciParserT__1)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == siciParserT__2 || _la == siciParserDIGIT {
		{
			p.SetState(53)
			_la = p.GetTokenStream().LA(1)

			if !(_la == siciParserT__2 || _la == siciParserDIGIT) {
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
	{
		p.SetState(58)
		p.Match(siciParserT__3)
	}

	return localctx
}

// IEnumerationContext is an interface to support dynamic dispatch.
type IEnumerationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumerationContext differentiates from other interfaces.
	IsEnumerationContext()
}

type EnumerationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumerationContext() *EnumerationContext {
	var p = new(EnumerationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_enumeration
	return p
}

func (*EnumerationContext) IsEnumerationContext() {}

func NewEnumerationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumerationContext {
	var p = new(EnumerationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_enumeration

	return p
}

func (s *EnumerationContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumerationContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(siciParserDIGIT)
}

func (s *EnumerationContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, i)
}

func (s *EnumerationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumerationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumerationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterEnumeration(s)
	}
}

func (s *EnumerationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitEnumeration(s)
	}
}

func (p *siciParser) Enumeration() (localctx IEnumerationContext) {
	this := p
	_ = this

	localctx = NewEnumerationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, siciParserRULE_enumeration)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == siciParserDIGIT {
		{
			p.SetState(60)
			p.Match(siciParserDIGIT)
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(siciParserT__4)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == siciParserDIGIT {
		{
			p.SetState(66)
			p.Match(siciParserDIGIT)
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IContributionContext is an interface to support dynamic dispatch.
type IContributionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContributionContext differentiates from other interfaces.
	IsContributionContext()
}

type ContributionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContributionContext() *ContributionContext {
	var p = new(ContributionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_contribution
	return p
}

func (*ContributionContext) IsContributionContext() {}

func NewContributionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContributionContext {
	var p = new(ContributionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_contribution

	return p
}

func (s *ContributionContext) GetParser() antlr.Parser { return s.parser }

func (s *ContributionContext) Location() ILocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *ContributionContext) Title() ITitleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITitleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *ContributionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContributionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContributionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterContribution(s)
	}
}

func (s *ContributionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitContribution(s)
	}
}

func (p *siciParser) Contribution() (localctx IContributionContext) {
	this := p
	_ = this

	localctx = NewContributionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, siciParserRULE_contribution)

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
		p.SetState(71)
		p.Match(siciParserT__5)
	}
	{
		p.SetState(72)
		p.Location()
	}
	{
		p.SetState(73)
		p.Match(siciParserT__4)
	}
	{
		p.SetState(74)
		p.Title()
	}
	{
		p.SetState(75)
		p.Match(siciParserT__6)
	}

	return localctx
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_location
	return p
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(siciParserDIGIT)
}

func (s *LocationContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, i)
}

func (s *LocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterLocation(s)
	}
}

func (s *LocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitLocation(s)
	}
}

func (p *siciParser) Location() (localctx ILocationContext) {
	this := p
	_ = this

	localctx = NewLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, siciParserRULE_location)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == siciParserDIGIT {
		{
			p.SetState(77)
			p.Match(siciParserDIGIT)
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITitleContext is an interface to support dynamic dispatch.
type ITitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTitleContext differentiates from other interfaces.
	IsTitleContext()
}

type TitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitleContext() *TitleContext {
	var p = new(TitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TitleContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(siciParserLETTER)
}

func (s *TitleContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(siciParserLETTER, i)
}

func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *siciParser) Title() (localctx ITitleContext) {
	this := p
	_ = this

	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, siciParserRULE_title)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == siciParserLETTER {
		{
			p.SetState(82)
			p.Match(siciParserLETTER)
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IControlContext is an interface to support dynamic dispatch.
type IControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlContext differentiates from other interfaces.
	IsControlContext()
}

type ControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlContext() *ControlContext {
	var p = new(ControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_control
	return p
}

func (*ControlContext) IsControlContext() {}

func NewControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlContext {
	var p = new(ControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_control

	return p
}

func (s *ControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlContext) Csi() ICsiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICsiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICsiContext)
}

func (s *ControlContext) Dpi() IDpiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDpiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDpiContext)
}

func (s *ControlContext) Mfi() IMfiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMfiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMfiContext)
}

func (s *ControlContext) Version() IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *ControlContext) Check() ICheckContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *ControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterControl(s)
	}
}

func (s *ControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitControl(s)
	}
}

func (p *siciParser) Control() (localctx IControlContext) {
	this := p
	_ = this

	localctx = NewControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, siciParserRULE_control)

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
		p.SetState(87)
		p.Csi()
	}
	{
		p.SetState(88)
		p.Match(siciParserT__7)
	}
	{
		p.SetState(89)
		p.Dpi()
	}
	{
		p.SetState(90)
		p.Match(siciParserT__7)
	}
	{
		p.SetState(91)
		p.Mfi()
	}
	{
		p.SetState(92)
		p.Match(siciParserT__8)
	}
	{
		p.SetState(93)
		p.Version()
	}
	{
		p.SetState(94)
		p.Check()
	}

	return localctx
}

// ICsiContext is an interface to support dynamic dispatch.
type ICsiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCsiContext differentiates from other interfaces.
	IsCsiContext()
}

type CsiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsiContext() *CsiContext {
	var p = new(CsiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_csi
	return p
}

func (*CsiContext) IsCsiContext() {}

func NewCsiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsiContext {
	var p = new(CsiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_csi

	return p
}

func (s *CsiContext) GetParser() antlr.Parser { return s.parser }

func (s *CsiContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, 0)
}

func (s *CsiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CsiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterCsi(s)
	}
}

func (s *CsiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitCsi(s)
	}
}

func (p *siciParser) Csi() (localctx ICsiContext) {
	this := p
	_ = this

	localctx = NewCsiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, siciParserRULE_csi)

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
		p.SetState(96)
		p.Match(siciParserDIGIT)
	}

	return localctx
}

// IDpiContext is an interface to support dynamic dispatch.
type IDpiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDpiContext differentiates from other interfaces.
	IsDpiContext()
}

type DpiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDpiContext() *DpiContext {
	var p = new(DpiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_dpi
	return p
}

func (*DpiContext) IsDpiContext() {}

func NewDpiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DpiContext {
	var p = new(DpiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_dpi

	return p
}

func (s *DpiContext) GetParser() antlr.Parser { return s.parser }

func (s *DpiContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, 0)
}

func (s *DpiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DpiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DpiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterDpi(s)
	}
}

func (s *DpiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitDpi(s)
	}
}

func (p *siciParser) Dpi() (localctx IDpiContext) {
	this := p
	_ = this

	localctx = NewDpiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, siciParserRULE_dpi)

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
		p.SetState(98)
		p.Match(siciParserDIGIT)
	}

	return localctx
}

// IMfiContext is an interface to support dynamic dispatch.
type IMfiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMfiContext differentiates from other interfaces.
	IsMfiContext()
}

type MfiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMfiContext() *MfiContext {
	var p = new(MfiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_mfi
	return p
}

func (*MfiContext) IsMfiContext() {}

func NewMfiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MfiContext {
	var p = new(MfiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_mfi

	return p
}

func (s *MfiContext) GetParser() antlr.Parser { return s.parser }

func (s *MfiContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(siciParserLETTER)
}

func (s *MfiContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(siciParserLETTER, i)
}

func (s *MfiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MfiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MfiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterMfi(s)
	}
}

func (s *MfiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitMfi(s)
	}
}

func (p *siciParser) Mfi() (localctx IMfiContext) {
	this := p
	_ = this

	localctx = NewMfiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, siciParserRULE_mfi)

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
		p.SetState(100)
		p.Match(siciParserLETTER)
	}
	{
		p.SetState(101)
		p.Match(siciParserLETTER)
	}

	return localctx
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_version
	return p
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(siciParserDIGIT)
}

func (s *VersionContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(siciParserDIGIT, i)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *siciParser) Version() (localctx IVersionContext) {
	this := p
	_ = this

	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, siciParserRULE_version)
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
		p.SetState(103)
		p.Match(siciParserDIGIT)
	}
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(_la == siciParserT__0 || _la == siciParserDIGIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICheckContext is an interface to support dynamic dispatch.
type ICheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckContext differentiates from other interfaces.
	IsCheckContext()
}

type CheckContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckContext() *CheckContext {
	var p = new(CheckContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = siciParserRULE_check
	return p
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = siciParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) LETTER() antlr.TerminalNode {
	return s.GetToken(siciParserLETTER, 0)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(siciListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *siciParser) Check() (localctx ICheckContext) {
	this := p
	_ = this

	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, siciParserRULE_check)

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
		p.SetState(106)
		p.Match(siciParserLETTER)
	}

	return localctx
}
