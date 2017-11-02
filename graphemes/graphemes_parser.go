// Generated from Graphemes.g4 by ANTLR 4.7.

package graphemes // Graphemes
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 44, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 7, 2, 11, 10, 2, 12, 2, 14,
	2, 14, 11, 2, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 26, 10, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32,
	11, 3, 5, 3, 34, 10, 3, 3, 4, 7, 4, 37, 10, 4, 12, 4, 14, 4, 40, 11, 4,
	3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 4, 4, 2, 6, 7, 12, 12, 3, 2, 3,
	5, 2, 47, 2, 8, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 12,
	9, 2, 2, 2, 9, 11, 9, 3, 2, 2, 10, 9, 3, 2, 2, 2, 11, 14, 3, 2, 2, 2, 12,
	10, 3, 2, 2, 2, 12, 13, 3, 2, 2, 2, 13, 3, 3, 2, 2, 2, 14, 12, 3, 2, 2,
	2, 15, 34, 7, 10, 2, 2, 16, 18, 7, 8, 2, 2, 17, 16, 3, 2, 2, 2, 18, 21,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 25, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 22, 26, 5, 2, 2, 2, 23, 26, 7, 11, 2, 2, 24, 26, 7,
	9, 2, 2, 25, 22, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26,
	30, 3, 2, 2, 2, 27, 29, 9, 3, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2,
	2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 33, 15, 3, 2, 2, 2, 33, 19, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2,
	35, 37, 5, 4, 3, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3,
	2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41,
	42, 7, 2, 2, 3, 42, 7, 3, 2, 2, 2, 8, 12, 19, 25, 30, 33, 38,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'\u200D'",
}
var symbolicNames = []string{
	"", "Extend", "ZWJ", "SpacingMark", "EmojiCoreSequence", "EmojiZWJSequence",
	"Prepend", "NonControl", "CRLF", "HangulSyllable", "EmojiTagSequence",
}

var ruleNames = []string{
	"emoji_sequence", "grapheme_cluster", "graphemes",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GraphemesParser struct {
	*antlr.BaseParser
}

func NewGraphemesParser(input antlr.TokenStream) *GraphemesParser {
	this := new(GraphemesParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Graphemes.g4"

	return this
}

// GraphemesParser tokens.
const (
	GraphemesParserEOF               = antlr.TokenEOF
	GraphemesParserExtend            = 1
	GraphemesParserZWJ               = 2
	GraphemesParserSpacingMark       = 3
	GraphemesParserEmojiCoreSequence = 4
	GraphemesParserEmojiZWJSequence  = 5
	GraphemesParserPrepend           = 6
	GraphemesParserNonControl        = 7
	GraphemesParserCRLF              = 8
	GraphemesParserHangulSyllable    = 9
	GraphemesParserEmojiTagSequence  = 10
)

// GraphemesParser rules.
const (
	GraphemesParserRULE_emoji_sequence   = 0
	GraphemesParserRULE_grapheme_cluster = 1
	GraphemesParserRULE_graphemes        = 2
)

// IEmoji_sequenceContext is an interface to support dynamic dispatch.
type IEmoji_sequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmoji_sequenceContext differentiates from other interfaces.
	IsEmoji_sequenceContext()
}

type Emoji_sequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmoji_sequenceContext() *Emoji_sequenceContext {
	var p = new(Emoji_sequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphemesParserRULE_emoji_sequence
	return p
}

func (*Emoji_sequenceContext) IsEmoji_sequenceContext() {}

func NewEmoji_sequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Emoji_sequenceContext {
	var p = new(Emoji_sequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphemesParserRULE_emoji_sequence

	return p
}

func (s *Emoji_sequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Emoji_sequenceContext) EmojiZWJSequence() antlr.TerminalNode {
	return s.GetToken(GraphemesParserEmojiZWJSequence, 0)
}

func (s *Emoji_sequenceContext) EmojiCoreSequence() antlr.TerminalNode {
	return s.GetToken(GraphemesParserEmojiCoreSequence, 0)
}

func (s *Emoji_sequenceContext) EmojiTagSequence() antlr.TerminalNode {
	return s.GetToken(GraphemesParserEmojiTagSequence, 0)
}

func (s *Emoji_sequenceContext) AllExtend() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserExtend)
}

func (s *Emoji_sequenceContext) Extend(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserExtend, i)
}

func (s *Emoji_sequenceContext) AllZWJ() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserZWJ)
}

func (s *Emoji_sequenceContext) ZWJ(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserZWJ, i)
}

func (s *Emoji_sequenceContext) AllSpacingMark() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserSpacingMark)
}

func (s *Emoji_sequenceContext) SpacingMark(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserSpacingMark, i)
}

func (s *Emoji_sequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Emoji_sequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Emoji_sequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.EnterEmoji_sequence(s)
	}
}

func (s *Emoji_sequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.ExitEmoji_sequence(s)
	}
}

func (p *GraphemesParser) Emoji_sequence() (localctx IEmoji_sequenceContext) {
	localctx = NewEmoji_sequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphemesParserRULE_emoji_sequence)
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
	p.SetState(6)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphemesParserEmojiCoreSequence)|(1<<GraphemesParserEmojiZWJSequence)|(1<<GraphemesParserEmojiTagSequence))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(7)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphemesParserExtend)|(1<<GraphemesParserZWJ)|(1<<GraphemesParserSpacingMark))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IGrapheme_clusterContext is an interface to support dynamic dispatch.
type IGrapheme_clusterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGrapheme_clusterContext differentiates from other interfaces.
	IsGrapheme_clusterContext()
}

type Grapheme_clusterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrapheme_clusterContext() *Grapheme_clusterContext {
	var p = new(Grapheme_clusterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphemesParserRULE_grapheme_cluster
	return p
}

func (*Grapheme_clusterContext) IsGrapheme_clusterContext() {}

func NewGrapheme_clusterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Grapheme_clusterContext {
	var p = new(Grapheme_clusterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphemesParserRULE_grapheme_cluster

	return p
}

func (s *Grapheme_clusterContext) GetParser() antlr.Parser { return s.parser }

func (s *Grapheme_clusterContext) CRLF() antlr.TerminalNode {
	return s.GetToken(GraphemesParserCRLF, 0)
}

func (s *Grapheme_clusterContext) Emoji_sequence() IEmoji_sequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmoji_sequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEmoji_sequenceContext)
}

func (s *Grapheme_clusterContext) HangulSyllable() antlr.TerminalNode {
	return s.GetToken(GraphemesParserHangulSyllable, 0)
}

func (s *Grapheme_clusterContext) NonControl() antlr.TerminalNode {
	return s.GetToken(GraphemesParserNonControl, 0)
}

func (s *Grapheme_clusterContext) AllPrepend() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserPrepend)
}

func (s *Grapheme_clusterContext) Prepend(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserPrepend, i)
}

func (s *Grapheme_clusterContext) AllExtend() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserExtend)
}

func (s *Grapheme_clusterContext) Extend(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserExtend, i)
}

func (s *Grapheme_clusterContext) AllZWJ() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserZWJ)
}

func (s *Grapheme_clusterContext) ZWJ(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserZWJ, i)
}

func (s *Grapheme_clusterContext) AllSpacingMark() []antlr.TerminalNode {
	return s.GetTokens(GraphemesParserSpacingMark)
}

func (s *Grapheme_clusterContext) SpacingMark(i int) antlr.TerminalNode {
	return s.GetToken(GraphemesParserSpacingMark, i)
}

func (s *Grapheme_clusterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Grapheme_clusterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Grapheme_clusterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.EnterGrapheme_cluster(s)
	}
}

func (s *Grapheme_clusterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.ExitGrapheme_cluster(s)
	}
}

func (p *GraphemesParser) Grapheme_cluster() (localctx IGrapheme_clusterContext) {
	localctx = NewGrapheme_clusterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphemesParserRULE_grapheme_cluster)
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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphemesParserCRLF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(13)
			p.Match(GraphemesParserCRLF)
		}

	case GraphemesParserEmojiCoreSequence, GraphemesParserEmojiZWJSequence, GraphemesParserPrepend, GraphemesParserNonControl, GraphemesParserHangulSyllable, GraphemesParserEmojiTagSequence:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraphemesParserPrepend {
			{
				p.SetState(14)
				p.Match(GraphemesParserPrepend)
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GraphemesParserEmojiCoreSequence, GraphemesParserEmojiZWJSequence, GraphemesParserEmojiTagSequence:
			{
				p.SetState(20)
				p.Emoji_sequence()
			}

		case GraphemesParserHangulSyllable:
			{
				p.SetState(21)
				p.Match(GraphemesParserHangulSyllable)
			}

		case GraphemesParserNonControl:
			{
				p.SetState(22)
				p.Match(GraphemesParserNonControl)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphemesParserExtend)|(1<<GraphemesParserZWJ)|(1<<GraphemesParserSpacingMark))) != 0 {
			p.SetState(25)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphemesParserExtend)|(1<<GraphemesParserZWJ)|(1<<GraphemesParserSpacingMark))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGraphemesContext is an interface to support dynamic dispatch.
type IGraphemesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphemesContext differentiates from other interfaces.
	IsGraphemesContext()
}

type GraphemesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphemesContext() *GraphemesContext {
	var p = new(GraphemesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphemesParserRULE_graphemes
	return p
}

func (*GraphemesContext) IsGraphemesContext() {}

func NewGraphemesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphemesContext {
	var p = new(GraphemesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphemesParserRULE_graphemes

	return p
}

func (s *GraphemesContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphemesContext) EOF() antlr.TerminalNode {
	return s.GetToken(GraphemesParserEOF, 0)
}

func (s *GraphemesContext) AllGrapheme_cluster() []IGrapheme_clusterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGrapheme_clusterContext)(nil)).Elem())
	var tst = make([]IGrapheme_clusterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGrapheme_clusterContext)
		}
	}

	return tst
}

func (s *GraphemesContext) Grapheme_cluster(i int) IGrapheme_clusterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGrapheme_clusterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGrapheme_clusterContext)
}

func (s *GraphemesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphemesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphemesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.EnterGraphemes(s)
	}
}

func (s *GraphemesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphemesListener); ok {
		listenerT.ExitGraphemes(s)
	}
}

func (p *GraphemesParser) Graphemes() (localctx IGraphemesContext) {
	localctx = NewGraphemesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphemesParserRULE_graphemes)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphemesParserEmojiCoreSequence)|(1<<GraphemesParserEmojiZWJSequence)|(1<<GraphemesParserPrepend)|(1<<GraphemesParserNonControl)|(1<<GraphemesParserCRLF)|(1<<GraphemesParserHangulSyllable)|(1<<GraphemesParserEmojiTagSequence))) != 0 {
		{
			p.SetState(33)
			p.Grapheme_cluster()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(GraphemesParserEOF)
	}

	return localctx
}
