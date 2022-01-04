// Code generated from wln.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wln // wln
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 26, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 2, 3, 2, 3, 3, 6, 3, 20, 10, 3, 13, 3, 14, 3, 21,
	3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3, 4, 2, 3, 24, 26, 26, 2, 24, 2,
	8, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 23, 3, 2, 2, 2, 8, 13, 5, 4, 3, 2,
	9, 10, 7, 27, 2, 2, 10, 12, 5, 4, 3, 2, 11, 9, 3, 2, 2, 2, 12, 15, 3, 2,
	2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 16, 3, 2, 2, 2, 15, 13,
	3, 2, 2, 2, 16, 17, 7, 2, 2, 3, 17, 3, 3, 2, 2, 2, 18, 20, 5, 6, 4, 2,
	19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3,
	2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 24, 9, 2, 2, 2, 24, 7, 3, 2, 2, 2, 4, 13,
	21,
}
var literalNames = []string{
	"", "'E'", "'G'", "'I'", "'R'", "'U'", "'V'", "'C'", "'K'", "'L'", "'M'",
	"'N'", "'O'", "'T'", "'W'", "'X'", "'Y'", "'Z'", "'F'", "'H'", "'Q'", "'S'",
	"'UU'", "'&'", "", "' '",
}
var symbolicNames = []string{
	"", "BROMINE", "CHLORINE", "IODINE", "BENZENE", "DOUBLEBOND", "CARBONYL",
	"CARBONNON", "NITROGEN3PLUS", "CARBOSYCLIC", "IMINO", "NITROGEN4LESS",
	"OXYGEN", "HETEROCYCLIC", "DIOXO", "CARBON4", "CARBON3", "AMINO", "FLOURINE",
	"HYDROGEN", "HYDROXYL", "SULFER", "TRIPLE", "AMP", "DIGIT", "SPACE", "WS",
}

var ruleNames = []string{
	"wln", "group", "symbol",
}

type wlnParser struct {
	*antlr.BaseParser
}

// NewwlnParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *wlnParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewwlnParser(input antlr.TokenStream) *wlnParser {
	this := new(wlnParser)
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
	this.GrammarFileName = "wln.g4"

	return this
}

// wlnParser tokens.
const (
	wlnParserEOF           = antlr.TokenEOF
	wlnParserBROMINE       = 1
	wlnParserCHLORINE      = 2
	wlnParserIODINE        = 3
	wlnParserBENZENE       = 4
	wlnParserDOUBLEBOND    = 5
	wlnParserCARBONYL      = 6
	wlnParserCARBONNON     = 7
	wlnParserNITROGEN3PLUS = 8
	wlnParserCARBOSYCLIC   = 9
	wlnParserIMINO         = 10
	wlnParserNITROGEN4LESS = 11
	wlnParserOXYGEN        = 12
	wlnParserHETEROCYCLIC  = 13
	wlnParserDIOXO         = 14
	wlnParserCARBON4       = 15
	wlnParserCARBON3       = 16
	wlnParserAMINO         = 17
	wlnParserFLOURINE      = 18
	wlnParserHYDROGEN      = 19
	wlnParserHYDROXYL      = 20
	wlnParserSULFER        = 21
	wlnParserTRIPLE        = 22
	wlnParserAMP           = 23
	wlnParserDIGIT         = 24
	wlnParserSPACE         = 25
	wlnParserWS            = 26
)

// wlnParser rules.
const (
	wlnParserRULE_wln    = 0
	wlnParserRULE_group  = 1
	wlnParserRULE_symbol = 2
)

// IWlnContext is an interface to support dynamic dispatch.
type IWlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWlnContext differentiates from other interfaces.
	IsWlnContext()
}

type WlnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWlnContext() *WlnContext {
	var p = new(WlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wlnParserRULE_wln
	return p
}

func (*WlnContext) IsWlnContext() {}

func NewWlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WlnContext {
	var p = new(WlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wlnParserRULE_wln

	return p
}

func (s *WlnContext) GetParser() antlr.Parser { return s.parser }

func (s *WlnContext) AllGroup() []IGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupContext)(nil)).Elem())
	var tst = make([]IGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupContext)
		}
	}

	return tst
}

func (s *WlnContext) Group(i int) IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *WlnContext) EOF() antlr.TerminalNode {
	return s.GetToken(wlnParserEOF, 0)
}

func (s *WlnContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(wlnParserSPACE)
}

func (s *WlnContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(wlnParserSPACE, i)
}

func (s *WlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.EnterWln(s)
	}
}

func (s *WlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.ExitWln(s)
	}
}

func (p *wlnParser) Wln() (localctx IWlnContext) {
	this := p
	_ = this

	localctx = NewWlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, wlnParserRULE_wln)
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
		p.SetState(6)
		p.Group()
	}
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wlnParserSPACE {
		{
			p.SetState(7)
			p.Match(wlnParserSPACE)
		}
		{
			p.SetState(8)
			p.Group()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(wlnParserEOF)
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wlnParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wlnParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) AllSymbol() []ISymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISymbolContext)(nil)).Elem())
	var tst = make([]ISymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISymbolContext)
		}
	}

	return tst
}

func (s *GroupContext) Symbol(i int) ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *wlnParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, wlnParserRULE_group)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<wlnParserBROMINE)|(1<<wlnParserCHLORINE)|(1<<wlnParserIODINE)|(1<<wlnParserBENZENE)|(1<<wlnParserDOUBLEBOND)|(1<<wlnParserCARBONYL)|(1<<wlnParserCARBONNON)|(1<<wlnParserNITROGEN3PLUS)|(1<<wlnParserCARBOSYCLIC)|(1<<wlnParserIMINO)|(1<<wlnParserNITROGEN4LESS)|(1<<wlnParserOXYGEN)|(1<<wlnParserHETEROCYCLIC)|(1<<wlnParserDIOXO)|(1<<wlnParserCARBON4)|(1<<wlnParserCARBON3)|(1<<wlnParserAMINO)|(1<<wlnParserFLOURINE)|(1<<wlnParserHYDROGEN)|(1<<wlnParserHYDROXYL)|(1<<wlnParserSULFER)|(1<<wlnParserTRIPLE)|(1<<wlnParserDIGIT))) != 0) {
		{
			p.SetState(16)
			p.Symbol()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wlnParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wlnParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) BROMINE() antlr.TerminalNode {
	return s.GetToken(wlnParserBROMINE, 0)
}

func (s *SymbolContext) CHLORINE() antlr.TerminalNode {
	return s.GetToken(wlnParserCHLORINE, 0)
}

func (s *SymbolContext) IODINE() antlr.TerminalNode {
	return s.GetToken(wlnParserIODINE, 0)
}

func (s *SymbolContext) BENZENE() antlr.TerminalNode {
	return s.GetToken(wlnParserBENZENE, 0)
}

func (s *SymbolContext) DOUBLEBOND() antlr.TerminalNode {
	return s.GetToken(wlnParserDOUBLEBOND, 0)
}

func (s *SymbolContext) CARBONYL() antlr.TerminalNode {
	return s.GetToken(wlnParserCARBONYL, 0)
}

func (s *SymbolContext) CARBONNON() antlr.TerminalNode {
	return s.GetToken(wlnParserCARBONNON, 0)
}

func (s *SymbolContext) NITROGEN3PLUS() antlr.TerminalNode {
	return s.GetToken(wlnParserNITROGEN3PLUS, 0)
}

func (s *SymbolContext) CARBOSYCLIC() antlr.TerminalNode {
	return s.GetToken(wlnParserCARBOSYCLIC, 0)
}

func (s *SymbolContext) IMINO() antlr.TerminalNode {
	return s.GetToken(wlnParserIMINO, 0)
}

func (s *SymbolContext) NITROGEN4LESS() antlr.TerminalNode {
	return s.GetToken(wlnParserNITROGEN4LESS, 0)
}

func (s *SymbolContext) OXYGEN() antlr.TerminalNode {
	return s.GetToken(wlnParserOXYGEN, 0)
}

func (s *SymbolContext) HETEROCYCLIC() antlr.TerminalNode {
	return s.GetToken(wlnParserHETEROCYCLIC, 0)
}

func (s *SymbolContext) DIOXO() antlr.TerminalNode {
	return s.GetToken(wlnParserDIOXO, 0)
}

func (s *SymbolContext) CARBON4() antlr.TerminalNode {
	return s.GetToken(wlnParserCARBON4, 0)
}

func (s *SymbolContext) CARBON3() antlr.TerminalNode {
	return s.GetToken(wlnParserCARBON3, 0)
}

func (s *SymbolContext) AMINO() antlr.TerminalNode {
	return s.GetToken(wlnParserAMINO, 0)
}

func (s *SymbolContext) FLOURINE() antlr.TerminalNode {
	return s.GetToken(wlnParserFLOURINE, 0)
}

func (s *SymbolContext) HYDROGEN() antlr.TerminalNode {
	return s.GetToken(wlnParserHYDROGEN, 0)
}

func (s *SymbolContext) HYDROXYL() antlr.TerminalNode {
	return s.GetToken(wlnParserHYDROXYL, 0)
}

func (s *SymbolContext) SULFER() antlr.TerminalNode {
	return s.GetToken(wlnParserSULFER, 0)
}

func (s *SymbolContext) TRIPLE() antlr.TerminalNode {
	return s.GetToken(wlnParserTRIPLE, 0)
}

func (s *SymbolContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(wlnParserDIGIT, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wlnListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *wlnParser) Symbol() (localctx ISymbolContext) {
	this := p
	_ = this

	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, wlnParserRULE_symbol)
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
		p.SetState(21)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<wlnParserBROMINE)|(1<<wlnParserCHLORINE)|(1<<wlnParserIODINE)|(1<<wlnParserBENZENE)|(1<<wlnParserDOUBLEBOND)|(1<<wlnParserCARBONYL)|(1<<wlnParserCARBONNON)|(1<<wlnParserNITROGEN3PLUS)|(1<<wlnParserCARBOSYCLIC)|(1<<wlnParserIMINO)|(1<<wlnParserNITROGEN4LESS)|(1<<wlnParserOXYGEN)|(1<<wlnParserHETEROCYCLIC)|(1<<wlnParserDIOXO)|(1<<wlnParserCARBON4)|(1<<wlnParserCARBON3)|(1<<wlnParserAMINO)|(1<<wlnParserFLOURINE)|(1<<wlnParserHYDROGEN)|(1<<wlnParserHYDROXYL)|(1<<wlnParserSULFER)|(1<<wlnParserTRIPLE)|(1<<wlnParserDIGIT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
