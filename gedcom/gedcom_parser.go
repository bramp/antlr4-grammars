// Code generated from gedcom.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gedcom // gedcom
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 116,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 6,
	2, 36, 10, 2, 13, 2, 14, 2, 37, 3, 3, 3, 3, 5, 3, 42, 10, 3, 3, 3, 3, 3,
	5, 3, 46, 10, 3, 3, 3, 3, 3, 3, 4, 6, 4, 51, 10, 4, 13, 4, 14, 4, 52, 3,
	5, 3, 5, 3, 6, 6, 6, 58, 10, 6, 13, 6, 14, 6, 59, 3, 7, 6, 7, 63, 10, 7,
	13, 7, 14, 7, 64, 3, 8, 3, 8, 3, 8, 5, 8, 70, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 82, 10, 10, 3, 11,
	6, 11, 85, 10, 11, 13, 11, 14, 11, 86, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 6, 13, 95, 10, 13, 13, 13, 14, 13, 96, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 103, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 112, 10, 16, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 19, 20, 3, 2, 6, 18, 2,
	119, 2, 35, 3, 2, 2, 2, 4, 39, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 54, 3,
	2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 62, 3, 2, 2, 2, 14, 69, 3, 2, 2, 2, 16,
	71, 3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 84, 3, 2, 2, 2, 22, 88, 3, 2, 2,
	2, 24, 94, 3, 2, 2, 2, 26, 102, 3, 2, 2, 2, 28, 104, 3, 2, 2, 2, 30, 111,
	3, 2, 2, 2, 32, 113, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34, 3, 2, 2, 2,
	36, 37, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 3, 3, 2,
	2, 2, 39, 41, 5, 6, 4, 2, 40, 42, 5, 8, 5, 2, 41, 40, 3, 2, 2, 2, 41, 42,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45, 5, 10, 6, 2, 44, 46, 5, 12, 7,
	2, 45, 44, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48,
	7, 21, 2, 2, 48, 5, 3, 2, 2, 2, 49, 51, 7, 20, 2, 2, 50, 49, 3, 2, 2, 2,
	51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 7, 3, 2,
	2, 2, 54, 55, 5, 22, 12, 2, 55, 9, 3, 2, 2, 2, 56, 58, 5, 28, 15, 2, 57,
	56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 11, 3, 2, 2, 2, 61, 63, 5, 14, 8, 2, 62, 61, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 13, 3, 2, 2, 2,
	66, 70, 5, 22, 12, 2, 67, 70, 5, 16, 9, 2, 68, 70, 5, 30, 16, 2, 69, 66,
	3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 15, 3, 2, 2, 2,
	71, 72, 7, 3, 2, 2, 72, 73, 7, 4, 2, 2, 73, 74, 5, 20, 11, 2, 74, 75, 7,
	3, 2, 2, 75, 76, 5, 18, 10, 2, 76, 17, 3, 2, 2, 2, 77, 82, 7, 19, 2, 2,
	78, 82, 7, 20, 2, 2, 79, 82, 5, 32, 17, 2, 80, 82, 7, 4, 2, 2, 81, 77,
	3, 2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2,
	82, 19, 3, 2, 2, 2, 83, 85, 5, 30, 16, 2, 84, 83, 3, 2, 2, 2, 85, 86, 3,
	2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 21, 3, 2, 2, 2, 88,
	89, 7, 3, 2, 2, 89, 90, 5, 28, 15, 2, 90, 91, 5, 24, 13, 2, 91, 92, 7,
	3, 2, 2, 92, 23, 3, 2, 2, 2, 93, 95, 5, 26, 14, 2, 94, 93, 3, 2, 2, 2,
	95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 25, 3,
	2, 2, 2, 98, 103, 7, 19, 2, 2, 99, 103, 7, 20, 2, 2, 100, 103, 5, 32, 17,
	2, 101, 103, 7, 4, 2, 2, 102, 98, 3, 2, 2, 2, 102, 99, 3, 2, 2, 2, 102,
	100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 27, 3, 2, 2, 2, 104, 105, 9,
	2, 2, 2, 105, 29, 3, 2, 2, 2, 106, 112, 7, 19, 2, 2, 107, 112, 7, 20, 2,
	2, 108, 112, 5, 32, 17, 2, 109, 112, 7, 4, 2, 2, 110, 112, 7, 5, 2, 2,
	111, 106, 3, 2, 2, 2, 111, 107, 3, 2, 2, 2, 111, 108, 3, 2, 2, 2, 111,
	109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 31, 3, 2, 2, 2, 113, 114, 9,
	3, 2, 2, 114, 33, 3, 2, 2, 2, 14, 37, 41, 45, 52, 59, 64, 69, 81, 86, 96,
	102, 111,
}
var literalNames = []string{
	"", "'@'", "'#'", "'@@'", "'!'", "'\"'", "'$'", "'&'", "'''", "'('", "')'",
	"'*'", "'+'", "'-'", "','", "'.'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ALPHA",
	"DIGIT", "EOL", "WS",
}

var ruleNames = []string{
	"gedcom", "line", "level", "opt_xref_id", "tag", "line_value", "line_item",
	"escape", "non_at", "escape_text", "pointer", "pointer_string", "pointer_char",
	"alphanum", "anychar", "otherchar",
}

type gedcomParser struct {
	*antlr.BaseParser
}

// NewgedcomParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gedcomParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgedcomParser(input antlr.TokenStream) *gedcomParser {
	this := new(gedcomParser)
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
	this.GrammarFileName = "gedcom.g4"

	return this
}

// gedcomParser tokens.
const (
	gedcomParserEOF   = antlr.TokenEOF
	gedcomParserT__0  = 1
	gedcomParserT__1  = 2
	gedcomParserT__2  = 3
	gedcomParserT__3  = 4
	gedcomParserT__4  = 5
	gedcomParserT__5  = 6
	gedcomParserT__6  = 7
	gedcomParserT__7  = 8
	gedcomParserT__8  = 9
	gedcomParserT__9  = 10
	gedcomParserT__10 = 11
	gedcomParserT__11 = 12
	gedcomParserT__12 = 13
	gedcomParserT__13 = 14
	gedcomParserT__14 = 15
	gedcomParserT__15 = 16
	gedcomParserALPHA = 17
	gedcomParserDIGIT = 18
	gedcomParserEOL   = 19
	gedcomParserWS    = 20
)

// gedcomParser rules.
const (
	gedcomParserRULE_gedcom         = 0
	gedcomParserRULE_line           = 1
	gedcomParserRULE_level          = 2
	gedcomParserRULE_opt_xref_id    = 3
	gedcomParserRULE_tag            = 4
	gedcomParserRULE_line_value     = 5
	gedcomParserRULE_line_item      = 6
	gedcomParserRULE_escape         = 7
	gedcomParserRULE_non_at         = 8
	gedcomParserRULE_escape_text    = 9
	gedcomParserRULE_pointer        = 10
	gedcomParserRULE_pointer_string = 11
	gedcomParserRULE_pointer_char   = 12
	gedcomParserRULE_alphanum       = 13
	gedcomParserRULE_anychar        = 14
	gedcomParserRULE_otherchar      = 15
)

// IGedcomContext is an interface to support dynamic dispatch.
type IGedcomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGedcomContext differentiates from other interfaces.
	IsGedcomContext()
}

type GedcomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGedcomContext() *GedcomContext {
	var p = new(GedcomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_gedcom
	return p
}

func (*GedcomContext) IsGedcomContext() {}

func NewGedcomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GedcomContext {
	var p = new(GedcomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_gedcom

	return p
}

func (s *GedcomContext) GetParser() antlr.Parser { return s.parser }

func (s *GedcomContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *GedcomContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *GedcomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GedcomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GedcomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterGedcom(s)
	}
}

func (s *GedcomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitGedcom(s)
	}
}

func (p *gedcomParser) Gedcom() (localctx IGedcomContext) {
	this := p
	_ = this

	localctx = NewGedcomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gedcomParserRULE_gedcom)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gedcomParserDIGIT {
		{
			p.SetState(32)
			p.Line()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = gedcomParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Level() ILevelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILevelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILevelContext)
}

func (s *LineContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *LineContext) EOL() antlr.TerminalNode {
	return s.GetToken(gedcomParserEOL, 0)
}

func (s *LineContext) Opt_xref_id() IOpt_xref_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpt_xref_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpt_xref_idContext)
}

func (s *LineContext) Line_value() ILine_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILine_valueContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *gedcomParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gedcomParserRULE_line)
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
		p.SetState(37)
		p.Level()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gedcomParserT__0 {
		{
			p.SetState(38)
			p.Opt_xref_id()
		}

	}
	{
		p.SetState(41)
		p.Tag()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gedcomParserT__0)|(1<<gedcomParserT__1)|(1<<gedcomParserT__2)|(1<<gedcomParserT__3)|(1<<gedcomParserT__4)|(1<<gedcomParserT__5)|(1<<gedcomParserT__6)|(1<<gedcomParserT__7)|(1<<gedcomParserT__8)|(1<<gedcomParserT__9)|(1<<gedcomParserT__10)|(1<<gedcomParserT__11)|(1<<gedcomParserT__12)|(1<<gedcomParserT__13)|(1<<gedcomParserT__14)|(1<<gedcomParserT__15)|(1<<gedcomParserALPHA)|(1<<gedcomParserDIGIT))) != 0 {
		{
			p.SetState(42)
			p.Line_value()
		}

	}
	{
		p.SetState(45)
		p.Match(gedcomParserEOL)
	}

	return localctx
}

// ILevelContext is an interface to support dynamic dispatch.
type ILevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLevelContext differentiates from other interfaces.
	IsLevelContext()
}

type LevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLevelContext() *LevelContext {
	var p = new(LevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_level
	return p
}

func (*LevelContext) IsLevelContext() {}

func NewLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LevelContext {
	var p = new(LevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_level

	return p
}

func (s *LevelContext) GetParser() antlr.Parser { return s.parser }

func (s *LevelContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(gedcomParserDIGIT)
}

func (s *LevelContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(gedcomParserDIGIT, i)
}

func (s *LevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterLevel(s)
	}
}

func (s *LevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitLevel(s)
	}
}

func (p *gedcomParser) Level() (localctx ILevelContext) {
	this := p
	_ = this

	localctx = NewLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gedcomParserRULE_level)

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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(47)
				p.Match(gedcomParserDIGIT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IOpt_xref_idContext is an interface to support dynamic dispatch.
type IOpt_xref_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpt_xref_idContext differentiates from other interfaces.
	IsOpt_xref_idContext()
}

type Opt_xref_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpt_xref_idContext() *Opt_xref_idContext {
	var p = new(Opt_xref_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_opt_xref_id
	return p
}

func (*Opt_xref_idContext) IsOpt_xref_idContext() {}

func NewOpt_xref_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Opt_xref_idContext {
	var p = new(Opt_xref_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_opt_xref_id

	return p
}

func (s *Opt_xref_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Opt_xref_idContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *Opt_xref_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Opt_xref_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Opt_xref_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterOpt_xref_id(s)
	}
}

func (s *Opt_xref_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitOpt_xref_id(s)
	}
}

func (p *gedcomParser) Opt_xref_id() (localctx IOpt_xref_idContext) {
	this := p
	_ = this

	localctx = NewOpt_xref_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gedcomParserRULE_opt_xref_id)

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
		p.Pointer()
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AllAlphanum() []IAlphanumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlphanumContext)(nil)).Elem())
	var tst = make([]IAlphanumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlphanumContext)
		}
	}

	return tst
}

func (s *TagContext) Alphanum(i int) IAlphanumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphanumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlphanumContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *gedcomParser) Tag() (localctx ITagContext) {
	this := p
	_ = this

	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gedcomParserRULE_tag)

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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(54)
				p.Alphanum()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ILine_valueContext is an interface to support dynamic dispatch.
type ILine_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLine_valueContext differentiates from other interfaces.
	IsLine_valueContext()
}

type Line_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_valueContext() *Line_valueContext {
	var p = new(Line_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_line_value
	return p
}

func (*Line_valueContext) IsLine_valueContext() {}

func NewLine_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_valueContext {
	var p = new(Line_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_line_value

	return p
}

func (s *Line_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_valueContext) AllLine_item() []ILine_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILine_itemContext)(nil)).Elem())
	var tst = make([]ILine_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILine_itemContext)
		}
	}

	return tst
}

func (s *Line_valueContext) Line_item(i int) ILine_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILine_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILine_itemContext)
}

func (s *Line_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterLine_value(s)
	}
}

func (s *Line_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitLine_value(s)
	}
}

func (p *gedcomParser) Line_value() (localctx ILine_valueContext) {
	this := p
	_ = this

	localctx = NewLine_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gedcomParserRULE_line_value)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gedcomParserT__0)|(1<<gedcomParserT__1)|(1<<gedcomParserT__2)|(1<<gedcomParserT__3)|(1<<gedcomParserT__4)|(1<<gedcomParserT__5)|(1<<gedcomParserT__6)|(1<<gedcomParserT__7)|(1<<gedcomParserT__8)|(1<<gedcomParserT__9)|(1<<gedcomParserT__10)|(1<<gedcomParserT__11)|(1<<gedcomParserT__12)|(1<<gedcomParserT__13)|(1<<gedcomParserT__14)|(1<<gedcomParserT__15)|(1<<gedcomParserALPHA)|(1<<gedcomParserDIGIT))) != 0) {
		{
			p.SetState(59)
			p.Line_item()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILine_itemContext is an interface to support dynamic dispatch.
type ILine_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLine_itemContext differentiates from other interfaces.
	IsLine_itemContext()
}

type Line_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_itemContext() *Line_itemContext {
	var p = new(Line_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_line_item
	return p
}

func (*Line_itemContext) IsLine_itemContext() {}

func NewLine_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_itemContext {
	var p = new(Line_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_line_item

	return p
}

func (s *Line_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_itemContext) Pointer() IPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointerContext)
}

func (s *Line_itemContext) Escape() IEscapeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscapeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscapeContext)
}

func (s *Line_itemContext) Anychar() IAnycharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnycharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnycharContext)
}

func (s *Line_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterLine_item(s)
	}
}

func (s *Line_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitLine_item(s)
	}
}

func (p *gedcomParser) Line_item() (localctx ILine_itemContext) {
	this := p
	_ = this

	localctx = NewLine_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gedcomParserRULE_line_item)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Pointer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Escape()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Anychar()
		}

	}

	return localctx
}

// IEscapeContext is an interface to support dynamic dispatch.
type IEscapeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscapeContext differentiates from other interfaces.
	IsEscapeContext()
}

type EscapeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscapeContext() *EscapeContext {
	var p = new(EscapeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_escape
	return p
}

func (*EscapeContext) IsEscapeContext() {}

func NewEscapeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EscapeContext {
	var p = new(EscapeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_escape

	return p
}

func (s *EscapeContext) GetParser() antlr.Parser { return s.parser }

func (s *EscapeContext) Escape_text() IEscape_textContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEscape_textContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEscape_textContext)
}

func (s *EscapeContext) Non_at() INon_atContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INon_atContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INon_atContext)
}

func (s *EscapeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscapeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EscapeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterEscape(s)
	}
}

func (s *EscapeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitEscape(s)
	}
}

func (p *gedcomParser) Escape() (localctx IEscapeContext) {
	this := p
	_ = this

	localctx = NewEscapeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gedcomParserRULE_escape)

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
		p.SetState(69)
		p.Match(gedcomParserT__0)
	}
	{
		p.SetState(70)
		p.Match(gedcomParserT__1)
	}
	{
		p.SetState(71)
		p.Escape_text()
	}
	{
		p.SetState(72)
		p.Match(gedcomParserT__0)
	}
	{
		p.SetState(73)
		p.Non_at()
	}

	return localctx
}

// INon_atContext is an interface to support dynamic dispatch.
type INon_atContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNon_atContext differentiates from other interfaces.
	IsNon_atContext()
}

type Non_atContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_atContext() *Non_atContext {
	var p = new(Non_atContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_non_at
	return p
}

func (*Non_atContext) IsNon_atContext() {}

func NewNon_atContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_atContext {
	var p = new(Non_atContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_non_at

	return p
}

func (s *Non_atContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_atContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(gedcomParserALPHA, 0)
}

func (s *Non_atContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(gedcomParserDIGIT, 0)
}

func (s *Non_atContext) Otherchar() IOthercharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOthercharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOthercharContext)
}

func (s *Non_atContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_atContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Non_atContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterNon_at(s)
	}
}

func (s *Non_atContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitNon_at(s)
	}
}

func (p *gedcomParser) Non_at() (localctx INon_atContext) {
	this := p
	_ = this

	localctx = NewNon_atContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gedcomParserRULE_non_at)

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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gedcomParserALPHA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(gedcomParserALPHA)
		}

	case gedcomParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(gedcomParserDIGIT)
		}

	case gedcomParserT__3, gedcomParserT__4, gedcomParserT__5, gedcomParserT__6, gedcomParserT__7, gedcomParserT__8, gedcomParserT__9, gedcomParserT__10, gedcomParserT__11, gedcomParserT__12, gedcomParserT__13, gedcomParserT__14, gedcomParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.Otherchar()
		}

	case gedcomParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.Match(gedcomParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEscape_textContext is an interface to support dynamic dispatch.
type IEscape_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEscape_textContext differentiates from other interfaces.
	IsEscape_textContext()
}

type Escape_textContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscape_textContext() *Escape_textContext {
	var p = new(Escape_textContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_escape_text
	return p
}

func (*Escape_textContext) IsEscape_textContext() {}

func NewEscape_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Escape_textContext {
	var p = new(Escape_textContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_escape_text

	return p
}

func (s *Escape_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Escape_textContext) AllAnychar() []IAnycharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnycharContext)(nil)).Elem())
	var tst = make([]IAnycharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnycharContext)
		}
	}

	return tst
}

func (s *Escape_textContext) Anychar(i int) IAnycharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnycharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnycharContext)
}

func (s *Escape_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Escape_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Escape_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterEscape_text(s)
	}
}

func (s *Escape_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitEscape_text(s)
	}
}

func (p *gedcomParser) Escape_text() (localctx IEscape_textContext) {
	this := p
	_ = this

	localctx = NewEscape_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gedcomParserRULE_escape_text)
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
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gedcomParserT__1)|(1<<gedcomParserT__2)|(1<<gedcomParserT__3)|(1<<gedcomParserT__4)|(1<<gedcomParserT__5)|(1<<gedcomParserT__6)|(1<<gedcomParserT__7)|(1<<gedcomParserT__8)|(1<<gedcomParserT__9)|(1<<gedcomParserT__10)|(1<<gedcomParserT__11)|(1<<gedcomParserT__12)|(1<<gedcomParserT__13)|(1<<gedcomParserT__14)|(1<<gedcomParserT__15)|(1<<gedcomParserALPHA)|(1<<gedcomParserDIGIT))) != 0) {
		{
			p.SetState(81)
			p.Anychar()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPointerContext is an interface to support dynamic dispatch.
type IPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointerContext differentiates from other interfaces.
	IsPointerContext()
}

type PointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerContext() *PointerContext {
	var p = new(PointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_pointer
	return p
}

func (*PointerContext) IsPointerContext() {}

func NewPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerContext {
	var p = new(PointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_pointer

	return p
}

func (s *PointerContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerContext) Alphanum() IAlphanumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlphanumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlphanumContext)
}

func (s *PointerContext) Pointer_string() IPointer_stringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointer_stringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointer_stringContext)
}

func (s *PointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterPointer(s)
	}
}

func (s *PointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitPointer(s)
	}
}

func (p *gedcomParser) Pointer() (localctx IPointerContext) {
	this := p
	_ = this

	localctx = NewPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gedcomParserRULE_pointer)

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
		p.SetState(86)
		p.Match(gedcomParserT__0)
	}
	{
		p.SetState(87)
		p.Alphanum()
	}
	{
		p.SetState(88)
		p.Pointer_string()
	}
	{
		p.SetState(89)
		p.Match(gedcomParserT__0)
	}

	return localctx
}

// IPointer_stringContext is an interface to support dynamic dispatch.
type IPointer_stringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointer_stringContext differentiates from other interfaces.
	IsPointer_stringContext()
}

type Pointer_stringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointer_stringContext() *Pointer_stringContext {
	var p = new(Pointer_stringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_pointer_string
	return p
}

func (*Pointer_stringContext) IsPointer_stringContext() {}

func NewPointer_stringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pointer_stringContext {
	var p = new(Pointer_stringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_pointer_string

	return p
}

func (s *Pointer_stringContext) GetParser() antlr.Parser { return s.parser }

func (s *Pointer_stringContext) AllPointer_char() []IPointer_charContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointer_charContext)(nil)).Elem())
	var tst = make([]IPointer_charContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointer_charContext)
		}
	}

	return tst
}

func (s *Pointer_stringContext) Pointer_char(i int) IPointer_charContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointer_charContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointer_charContext)
}

func (s *Pointer_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pointer_stringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pointer_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterPointer_string(s)
	}
}

func (s *Pointer_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitPointer_string(s)
	}
}

func (p *gedcomParser) Pointer_string() (localctx IPointer_stringContext) {
	this := p
	_ = this

	localctx = NewPointer_stringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gedcomParserRULE_pointer_string)
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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gedcomParserT__1)|(1<<gedcomParserT__3)|(1<<gedcomParserT__4)|(1<<gedcomParserT__5)|(1<<gedcomParserT__6)|(1<<gedcomParserT__7)|(1<<gedcomParserT__8)|(1<<gedcomParserT__9)|(1<<gedcomParserT__10)|(1<<gedcomParserT__11)|(1<<gedcomParserT__12)|(1<<gedcomParserT__13)|(1<<gedcomParserT__14)|(1<<gedcomParserT__15)|(1<<gedcomParserALPHA)|(1<<gedcomParserDIGIT))) != 0) {
		{
			p.SetState(91)
			p.Pointer_char()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPointer_charContext is an interface to support dynamic dispatch.
type IPointer_charContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointer_charContext differentiates from other interfaces.
	IsPointer_charContext()
}

type Pointer_charContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointer_charContext() *Pointer_charContext {
	var p = new(Pointer_charContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_pointer_char
	return p
}

func (*Pointer_charContext) IsPointer_charContext() {}

func NewPointer_charContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pointer_charContext {
	var p = new(Pointer_charContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_pointer_char

	return p
}

func (s *Pointer_charContext) GetParser() antlr.Parser { return s.parser }

func (s *Pointer_charContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(gedcomParserALPHA, 0)
}

func (s *Pointer_charContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(gedcomParserDIGIT, 0)
}

func (s *Pointer_charContext) Otherchar() IOthercharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOthercharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOthercharContext)
}

func (s *Pointer_charContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pointer_charContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pointer_charContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterPointer_char(s)
	}
}

func (s *Pointer_charContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitPointer_char(s)
	}
}

func (p *gedcomParser) Pointer_char() (localctx IPointer_charContext) {
	this := p
	_ = this

	localctx = NewPointer_charContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gedcomParserRULE_pointer_char)

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

	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gedcomParserALPHA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(gedcomParserALPHA)
		}

	case gedcomParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(gedcomParserDIGIT)
		}

	case gedcomParserT__3, gedcomParserT__4, gedcomParserT__5, gedcomParserT__6, gedcomParserT__7, gedcomParserT__8, gedcomParserT__9, gedcomParserT__10, gedcomParserT__11, gedcomParserT__12, gedcomParserT__13, gedcomParserT__14, gedcomParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Otherchar()
		}

	case gedcomParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(99)
			p.Match(gedcomParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAlphanumContext is an interface to support dynamic dispatch.
type IAlphanumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlphanumContext differentiates from other interfaces.
	IsAlphanumContext()
}

type AlphanumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlphanumContext() *AlphanumContext {
	var p = new(AlphanumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_alphanum
	return p
}

func (*AlphanumContext) IsAlphanumContext() {}

func NewAlphanumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlphanumContext {
	var p = new(AlphanumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_alphanum

	return p
}

func (s *AlphanumContext) GetParser() antlr.Parser { return s.parser }

func (s *AlphanumContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(gedcomParserALPHA, 0)
}

func (s *AlphanumContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(gedcomParserDIGIT, 0)
}

func (s *AlphanumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlphanumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlphanumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterAlphanum(s)
	}
}

func (s *AlphanumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitAlphanum(s)
	}
}

func (p *gedcomParser) Alphanum() (localctx IAlphanumContext) {
	this := p
	_ = this

	localctx = NewAlphanumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gedcomParserRULE_alphanum)
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
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gedcomParserALPHA || _la == gedcomParserDIGIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAnycharContext is an interface to support dynamic dispatch.
type IAnycharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnycharContext differentiates from other interfaces.
	IsAnycharContext()
}

type AnycharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnycharContext() *AnycharContext {
	var p = new(AnycharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_anychar
	return p
}

func (*AnycharContext) IsAnycharContext() {}

func NewAnycharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnycharContext {
	var p = new(AnycharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_anychar

	return p
}

func (s *AnycharContext) GetParser() antlr.Parser { return s.parser }

func (s *AnycharContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(gedcomParserALPHA, 0)
}

func (s *AnycharContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(gedcomParserDIGIT, 0)
}

func (s *AnycharContext) Otherchar() IOthercharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOthercharContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOthercharContext)
}

func (s *AnycharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnycharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnycharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterAnychar(s)
	}
}

func (s *AnycharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitAnychar(s)
	}
}

func (p *gedcomParser) Anychar() (localctx IAnycharContext) {
	this := p
	_ = this

	localctx = NewAnycharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gedcomParserRULE_anychar)

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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gedcomParserALPHA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(gedcomParserALPHA)
		}

	case gedcomParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(gedcomParserDIGIT)
		}

	case gedcomParserT__3, gedcomParserT__4, gedcomParserT__5, gedcomParserT__6, gedcomParserT__7, gedcomParserT__8, gedcomParserT__9, gedcomParserT__10, gedcomParserT__11, gedcomParserT__12, gedcomParserT__13, gedcomParserT__14, gedcomParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Otherchar()
		}

	case gedcomParserT__1:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(107)
			p.Match(gedcomParserT__1)
		}

	case gedcomParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)
			p.Match(gedcomParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOthercharContext is an interface to support dynamic dispatch.
type IOthercharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOthercharContext differentiates from other interfaces.
	IsOthercharContext()
}

type OthercharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOthercharContext() *OthercharContext {
	var p = new(OthercharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gedcomParserRULE_otherchar
	return p
}

func (*OthercharContext) IsOthercharContext() {}

func NewOthercharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OthercharContext {
	var p = new(OthercharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gedcomParserRULE_otherchar

	return p
}

func (s *OthercharContext) GetParser() antlr.Parser { return s.parser }
func (s *OthercharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OthercharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OthercharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.EnterOtherchar(s)
	}
}

func (s *OthercharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gedcomListener); ok {
		listenerT.ExitOtherchar(s)
	}
}

func (p *gedcomParser) Otherchar() (localctx IOthercharContext) {
	this := p
	_ = this

	localctx = NewOthercharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gedcomParserRULE_otherchar)
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
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gedcomParserT__3)|(1<<gedcomParserT__4)|(1<<gedcomParserT__5)|(1<<gedcomParserT__6)|(1<<gedcomParserT__7)|(1<<gedcomParserT__8)|(1<<gedcomParserT__9)|(1<<gedcomParserT__10)|(1<<gedcomParserT__11)|(1<<gedcomParserT__12)|(1<<gedcomParserT__13)|(1<<gedcomParserT__14)|(1<<gedcomParserT__15))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
