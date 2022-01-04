// Code generated from fen.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fen // fen
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 71, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	6, 4, 39, 10, 4, 13, 4, 14, 4, 40, 5, 4, 43, 10, 4, 3, 5, 3, 5, 5, 5, 47,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 6, 9,
	59, 10, 9, 13, 9, 14, 9, 60, 3, 10, 3, 10, 6, 10, 65, 10, 10, 13, 10, 14,
	10, 66, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 2, 6, 3, 2, 4, 5, 3, 2, 7, 10, 4, 2, 5, 5, 11, 17, 5, 2, 5, 5, 7, 10,
	19, 25, 2, 66, 2, 22, 3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2,
	8, 46, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 51, 3, 2, 2, 2, 14, 53, 3, 2,
	2, 2, 16, 55, 3, 2, 2, 2, 18, 64, 3, 2, 2, 2, 20, 68, 3, 2, 2, 2, 22, 23,
	5, 16, 9, 2, 23, 24, 7, 3, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 7, 3, 2, 2,
	26, 27, 5, 6, 4, 2, 27, 28, 7, 3, 2, 2, 28, 29, 5, 8, 5, 2, 29, 30, 7,
	3, 2, 2, 30, 31, 5, 12, 7, 2, 31, 32, 7, 3, 2, 2, 32, 33, 5, 14, 8, 2,
	33, 3, 3, 2, 2, 2, 34, 35, 9, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36, 43, 7, 6,
	2, 2, 37, 39, 9, 3, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38,
	3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 36, 3, 2, 2, 2,
	42, 38, 3, 2, 2, 2, 43, 7, 3, 2, 2, 2, 44, 47, 7, 6, 2, 2, 45, 47, 5, 10,
	6, 2, 46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 9, 3, 2, 2, 2, 48, 49,
	9, 4, 2, 2, 49, 50, 7, 26, 2, 2, 50, 11, 3, 2, 2, 2, 51, 52, 7, 26, 2,
	2, 52, 13, 3, 2, 2, 2, 53, 54, 7, 26, 2, 2, 54, 15, 3, 2, 2, 2, 55, 58,
	5, 18, 10, 2, 56, 57, 7, 18, 2, 2, 57, 59, 5, 18, 10, 2, 58, 56, 3, 2,
	2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 17,
	3, 2, 2, 2, 62, 65, 5, 20, 11, 2, 63, 65, 7, 26, 2, 2, 64, 62, 3, 2, 2,
	2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 19, 3, 2, 2, 2, 68, 69, 9, 5, 2, 2, 69, 21, 3, 2, 2, 2,
	8, 40, 42, 46, 60, 64, 66,
}
var literalNames = []string{
	"", "' '", "'w'", "'b'", "'-'", "'Q'", "'K'", "'k'", "'q'", "'a'", "'c'",
	"'d'", "'e'", "'f'", "'g'", "'h'", "'/'", "'p'", "'r'", "'n'", "'P'", "'R'",
	"'N'", "'B'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "NUMBER", "WS",
}

var ruleNames = []string{
	"fen", "color", "castling", "enpassant", "position", "halfmoveclock", "fullmoveclock",
	"placement", "rank", "piece",
}

type fenParser struct {
	*antlr.BaseParser
}

// NewfenParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *fenParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfenParser(input antlr.TokenStream) *fenParser {
	this := new(fenParser)
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
	this.GrammarFileName = "fen.g4"

	return this
}

// fenParser tokens.
const (
	fenParserEOF    = antlr.TokenEOF
	fenParserT__0   = 1
	fenParserT__1   = 2
	fenParserT__2   = 3
	fenParserT__3   = 4
	fenParserT__4   = 5
	fenParserT__5   = 6
	fenParserT__6   = 7
	fenParserT__7   = 8
	fenParserT__8   = 9
	fenParserT__9   = 10
	fenParserT__10  = 11
	fenParserT__11  = 12
	fenParserT__12  = 13
	fenParserT__13  = 14
	fenParserT__14  = 15
	fenParserT__15  = 16
	fenParserT__16  = 17
	fenParserT__17  = 18
	fenParserT__18  = 19
	fenParserT__19  = 20
	fenParserT__20  = 21
	fenParserT__21  = 22
	fenParserT__22  = 23
	fenParserNUMBER = 24
	fenParserWS     = 25
)

// fenParser rules.
const (
	fenParserRULE_fen           = 0
	fenParserRULE_color         = 1
	fenParserRULE_castling      = 2
	fenParserRULE_enpassant     = 3
	fenParserRULE_position      = 4
	fenParserRULE_halfmoveclock = 5
	fenParserRULE_fullmoveclock = 6
	fenParserRULE_placement     = 7
	fenParserRULE_rank          = 8
	fenParserRULE_piece         = 9
)

// IFenContext is an interface to support dynamic dispatch.
type IFenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFenContext differentiates from other interfaces.
	IsFenContext()
}

type FenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFenContext() *FenContext {
	var p = new(FenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_fen
	return p
}

func (*FenContext) IsFenContext() {}

func NewFenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FenContext {
	var p = new(FenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_fen

	return p
}

func (s *FenContext) GetParser() antlr.Parser { return s.parser }

func (s *FenContext) Placement() IPlacementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlacementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlacementContext)
}

func (s *FenContext) Color() IColorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColorContext)
}

func (s *FenContext) Castling() ICastlingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICastlingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICastlingContext)
}

func (s *FenContext) Enpassant() IEnpassantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnpassantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnpassantContext)
}

func (s *FenContext) Halfmoveclock() IHalfmoveclockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHalfmoveclockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHalfmoveclockContext)
}

func (s *FenContext) Fullmoveclock() IFullmoveclockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullmoveclockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullmoveclockContext)
}

func (s *FenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterFen(s)
	}
}

func (s *FenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitFen(s)
	}
}

func (p *fenParser) Fen() (localctx IFenContext) {
	this := p
	_ = this

	localctx = NewFenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, fenParserRULE_fen)

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
		p.Placement()
	}
	{
		p.SetState(21)
		p.Match(fenParserT__0)
	}
	{
		p.SetState(22)
		p.Color()
	}
	{
		p.SetState(23)
		p.Match(fenParserT__0)
	}
	{
		p.SetState(24)
		p.Castling()
	}
	{
		p.SetState(25)
		p.Match(fenParserT__0)
	}
	{
		p.SetState(26)
		p.Enpassant()
	}
	{
		p.SetState(27)
		p.Match(fenParserT__0)
	}
	{
		p.SetState(28)
		p.Halfmoveclock()
	}
	{
		p.SetState(29)
		p.Match(fenParserT__0)
	}
	{
		p.SetState(30)
		p.Fullmoveclock()
	}

	return localctx
}

// IColorContext is an interface to support dynamic dispatch.
type IColorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColorContext differentiates from other interfaces.
	IsColorContext()
}

type ColorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColorContext() *ColorContext {
	var p = new(ColorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_color
	return p
}

func (*ColorContext) IsColorContext() {}

func NewColorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColorContext {
	var p = new(ColorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_color

	return p
}

func (s *ColorContext) GetParser() antlr.Parser { return s.parser }
func (s *ColorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterColor(s)
	}
}

func (s *ColorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitColor(s)
	}
}

func (p *fenParser) Color() (localctx IColorContext) {
	this := p
	_ = this

	localctx = NewColorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, fenParserRULE_color)
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
		p.SetState(32)
		_la = p.GetTokenStream().LA(1)

		if !(_la == fenParserT__1 || _la == fenParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICastlingContext is an interface to support dynamic dispatch.
type ICastlingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCastlingContext differentiates from other interfaces.
	IsCastlingContext()
}

type CastlingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastlingContext() *CastlingContext {
	var p = new(CastlingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_castling
	return p
}

func (*CastlingContext) IsCastlingContext() {}

func NewCastlingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastlingContext {
	var p = new(CastlingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_castling

	return p
}

func (s *CastlingContext) GetParser() antlr.Parser { return s.parser }
func (s *CastlingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastlingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastlingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterCastling(s)
	}
}

func (s *CastlingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitCastling(s)
	}
}

func (p *fenParser) Castling() (localctx ICastlingContext) {
	this := p
	_ = this

	localctx = NewCastlingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fenParserRULE_castling)
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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case fenParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(fenParserT__3)
		}

	case fenParserT__4, fenParserT__5, fenParserT__6, fenParserT__7:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fenParserT__4)|(1<<fenParserT__5)|(1<<fenParserT__6)|(1<<fenParserT__7))) != 0) {
			{
				p.SetState(35)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fenParserT__4)|(1<<fenParserT__5)|(1<<fenParserT__6)|(1<<fenParserT__7))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnpassantContext is an interface to support dynamic dispatch.
type IEnpassantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnpassantContext differentiates from other interfaces.
	IsEnpassantContext()
}

type EnpassantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnpassantContext() *EnpassantContext {
	var p = new(EnpassantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_enpassant
	return p
}

func (*EnpassantContext) IsEnpassantContext() {}

func NewEnpassantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnpassantContext {
	var p = new(EnpassantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_enpassant

	return p
}

func (s *EnpassantContext) GetParser() antlr.Parser { return s.parser }

func (s *EnpassantContext) Position() IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *EnpassantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnpassantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnpassantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterEnpassant(s)
	}
}

func (s *EnpassantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitEnpassant(s)
	}
}

func (p *fenParser) Enpassant() (localctx IEnpassantContext) {
	this := p
	_ = this

	localctx = NewEnpassantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fenParserRULE_enpassant)

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

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case fenParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Match(fenParserT__3)
		}

	case fenParserT__2, fenParserT__8, fenParserT__9, fenParserT__10, fenParserT__11, fenParserT__12, fenParserT__13, fenParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Position()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPositionContext is an interface to support dynamic dispatch.
type IPositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositionContext differentiates from other interfaces.
	IsPositionContext()
}

type PositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositionContext() *PositionContext {
	var p = new(PositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_position
	return p
}

func (*PositionContext) IsPositionContext() {}

func NewPositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositionContext {
	var p = new(PositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_position

	return p
}

func (s *PositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PositionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(fenParserNUMBER, 0)
}

func (s *PositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterPosition(s)
	}
}

func (s *PositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitPosition(s)
	}
}

func (p *fenParser) Position() (localctx IPositionContext) {
	this := p
	_ = this

	localctx = NewPositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fenParserRULE_position)
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
		p.SetState(46)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fenParserT__2)|(1<<fenParserT__8)|(1<<fenParserT__9)|(1<<fenParserT__10)|(1<<fenParserT__11)|(1<<fenParserT__12)|(1<<fenParserT__13)|(1<<fenParserT__14))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(47)
		p.Match(fenParserNUMBER)
	}

	return localctx
}

// IHalfmoveclockContext is an interface to support dynamic dispatch.
type IHalfmoveclockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHalfmoveclockContext differentiates from other interfaces.
	IsHalfmoveclockContext()
}

type HalfmoveclockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHalfmoveclockContext() *HalfmoveclockContext {
	var p = new(HalfmoveclockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_halfmoveclock
	return p
}

func (*HalfmoveclockContext) IsHalfmoveclockContext() {}

func NewHalfmoveclockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HalfmoveclockContext {
	var p = new(HalfmoveclockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_halfmoveclock

	return p
}

func (s *HalfmoveclockContext) GetParser() antlr.Parser { return s.parser }

func (s *HalfmoveclockContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(fenParserNUMBER, 0)
}

func (s *HalfmoveclockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HalfmoveclockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HalfmoveclockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterHalfmoveclock(s)
	}
}

func (s *HalfmoveclockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitHalfmoveclock(s)
	}
}

func (p *fenParser) Halfmoveclock() (localctx IHalfmoveclockContext) {
	this := p
	_ = this

	localctx = NewHalfmoveclockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, fenParserRULE_halfmoveclock)

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
		p.Match(fenParserNUMBER)
	}

	return localctx
}

// IFullmoveclockContext is an interface to support dynamic dispatch.
type IFullmoveclockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullmoveclockContext differentiates from other interfaces.
	IsFullmoveclockContext()
}

type FullmoveclockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullmoveclockContext() *FullmoveclockContext {
	var p = new(FullmoveclockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_fullmoveclock
	return p
}

func (*FullmoveclockContext) IsFullmoveclockContext() {}

func NewFullmoveclockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullmoveclockContext {
	var p = new(FullmoveclockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_fullmoveclock

	return p
}

func (s *FullmoveclockContext) GetParser() antlr.Parser { return s.parser }

func (s *FullmoveclockContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(fenParserNUMBER, 0)
}

func (s *FullmoveclockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullmoveclockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullmoveclockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterFullmoveclock(s)
	}
}

func (s *FullmoveclockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitFullmoveclock(s)
	}
}

func (p *fenParser) Fullmoveclock() (localctx IFullmoveclockContext) {
	this := p
	_ = this

	localctx = NewFullmoveclockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, fenParserRULE_fullmoveclock)

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
		p.SetState(51)
		p.Match(fenParserNUMBER)
	}

	return localctx
}

// IPlacementContext is an interface to support dynamic dispatch.
type IPlacementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlacementContext differentiates from other interfaces.
	IsPlacementContext()
}

type PlacementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlacementContext() *PlacementContext {
	var p = new(PlacementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_placement
	return p
}

func (*PlacementContext) IsPlacementContext() {}

func NewPlacementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlacementContext {
	var p = new(PlacementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_placement

	return p
}

func (s *PlacementContext) GetParser() antlr.Parser { return s.parser }

func (s *PlacementContext) AllRank() []IRankContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRankContext)(nil)).Elem())
	var tst = make([]IRankContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRankContext)
		}
	}

	return tst
}

func (s *PlacementContext) Rank(i int) IRankContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRankContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRankContext)
}

func (s *PlacementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlacementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlacementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterPlacement(s)
	}
}

func (s *PlacementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitPlacement(s)
	}
}

func (p *fenParser) Placement() (localctx IPlacementContext) {
	this := p
	_ = this

	localctx = NewPlacementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, fenParserRULE_placement)
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
		p.SetState(53)
		p.Rank()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == fenParserT__15 {
		{
			p.SetState(54)
			p.Match(fenParserT__15)
		}
		{
			p.SetState(55)
			p.Rank()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRankContext is an interface to support dynamic dispatch.
type IRankContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRankContext differentiates from other interfaces.
	IsRankContext()
}

type RankContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRankContext() *RankContext {
	var p = new(RankContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_rank
	return p
}

func (*RankContext) IsRankContext() {}

func NewRankContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RankContext {
	var p = new(RankContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_rank

	return p
}

func (s *RankContext) GetParser() antlr.Parser { return s.parser }

func (s *RankContext) AllPiece() []IPieceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPieceContext)(nil)).Elem())
	var tst = make([]IPieceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPieceContext)
		}
	}

	return tst
}

func (s *RankContext) Piece(i int) IPieceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPieceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPieceContext)
}

func (s *RankContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(fenParserNUMBER)
}

func (s *RankContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(fenParserNUMBER, i)
}

func (s *RankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RankContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterRank(s)
	}
}

func (s *RankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitRank(s)
	}
}

func (p *fenParser) Rank() (localctx IRankContext) {
	this := p
	_ = this

	localctx = NewRankContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, fenParserRULE_rank)
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
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fenParserT__2)|(1<<fenParserT__4)|(1<<fenParserT__5)|(1<<fenParserT__6)|(1<<fenParserT__7)|(1<<fenParserT__16)|(1<<fenParserT__17)|(1<<fenParserT__18)|(1<<fenParserT__19)|(1<<fenParserT__20)|(1<<fenParserT__21)|(1<<fenParserT__22)|(1<<fenParserNUMBER))) != 0) {
		p.SetState(62)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case fenParserT__2, fenParserT__4, fenParserT__5, fenParserT__6, fenParserT__7, fenParserT__16, fenParserT__17, fenParserT__18, fenParserT__19, fenParserT__20, fenParserT__21, fenParserT__22:
			{
				p.SetState(60)
				p.Piece()
			}

		case fenParserNUMBER:
			{
				p.SetState(61)
				p.Match(fenParserNUMBER)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPieceContext is an interface to support dynamic dispatch.
type IPieceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPieceContext differentiates from other interfaces.
	IsPieceContext()
}

type PieceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPieceContext() *PieceContext {
	var p = new(PieceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fenParserRULE_piece
	return p
}

func (*PieceContext) IsPieceContext() {}

func NewPieceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PieceContext {
	var p = new(PieceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fenParserRULE_piece

	return p
}

func (s *PieceContext) GetParser() antlr.Parser { return s.parser }
func (s *PieceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PieceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PieceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.EnterPiece(s)
	}
}

func (s *PieceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fenListener); ok {
		listenerT.ExitPiece(s)
	}
}

func (p *fenParser) Piece() (localctx IPieceContext) {
	this := p
	_ = this

	localctx = NewPieceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, fenParserRULE_piece)
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
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fenParserT__2)|(1<<fenParserT__4)|(1<<fenParserT__5)|(1<<fenParserT__6)|(1<<fenParserT__7)|(1<<fenParserT__16)|(1<<fenParserT__17)|(1<<fenParserT__18)|(1<<fenParserT__19)|(1<<fenParserT__20)|(1<<fenParserT__21)|(1<<fenParserT__22))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
