// Code generated from edn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package edn // edn
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 86, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 7, 2, 20, 10, 2, 12, 2, 14, 2, 23, 11, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 7, 6, 50, 10, 6, 12, 6, 14, 6, 53, 11, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 7, 7, 59, 10, 7, 12, 7, 14, 7, 62, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 7, 8, 70, 10, 8, 12, 8, 14, 8, 73, 11, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 7, 9, 79, 10, 9, 12, 9, 14, 9, 82, 11, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10,
	2, 4, 6, 8, 10, 12, 14, 16, 2, 3, 5, 2, 3, 4, 6, 7, 20, 20, 2, 94, 2, 21,
	3, 2, 2, 2, 4, 39, 3, 2, 2, 2, 6, 41, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10,
	47, 3, 2, 2, 2, 12, 56, 3, 2, 2, 2, 14, 65, 3, 2, 2, 2, 16, 76, 3, 2, 2,
	2, 18, 20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19,
	3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 24, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2,
	24, 25, 7, 2, 2, 3, 25, 3, 3, 2, 2, 2, 26, 40, 7, 3, 2, 2, 27, 40, 7, 4,
	2, 2, 28, 40, 7, 5, 2, 2, 29, 40, 7, 6, 2, 2, 30, 40, 7, 7, 2, 2, 31, 40,
	7, 8, 2, 2, 32, 40, 5, 8, 5, 2, 33, 40, 7, 20, 2, 2, 34, 40, 5, 6, 4, 2,
	35, 40, 5, 10, 6, 2, 36, 40, 5, 12, 7, 2, 37, 40, 5, 14, 8, 2, 38, 40,
	5, 16, 9, 2, 39, 26, 3, 2, 2, 2, 39, 27, 3, 2, 2, 2, 39, 28, 3, 2, 2, 2,
	39, 29, 3, 2, 2, 2, 39, 30, 3, 2, 2, 2, 39, 31, 3, 2, 2, 2, 39, 32, 3,
	2, 2, 2, 39, 33, 3, 2, 2, 2, 39, 34, 3, 2, 2, 2, 39, 35, 3, 2, 2, 2, 39,
	36, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 5, 3, 2, 2,
	2, 41, 42, 7, 16, 2, 2, 42, 43, 7, 20, 2, 2, 43, 7, 3, 2, 2, 2, 44, 45,
	7, 17, 2, 2, 45, 46, 9, 2, 2, 2, 46, 9, 3, 2, 2, 2, 47, 51, 7, 9, 2, 2,
	48, 50, 5, 4, 3, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	55, 7, 10, 2, 2, 55, 11, 3, 2, 2, 2, 56, 60, 7, 11, 2, 2, 57, 59, 5, 4,
	3, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 7, 12, 2, 2,
	64, 13, 3, 2, 2, 2, 65, 71, 7, 13, 2, 2, 66, 67, 5, 4, 3, 2, 67, 68, 5,
	4, 3, 2, 68, 70, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71,
	69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 73, 71, 3, 2, 2,
	2, 74, 75, 7, 14, 2, 2, 75, 15, 3, 2, 2, 2, 76, 80, 7, 15, 2, 2, 77, 79,
	5, 4, 3, 2, 78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2,
	80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84, 7,
	14, 2, 2, 84, 17, 3, 2, 2, 2, 8, 21, 39, 51, 60, 71, 80,
}
var literalNames = []string{
	"", "'nil'", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'#{'", "'#'", "':'",
}
var symbolicNames = []string{
	"", "NilLiteral", "BooleanLiteral", "StringLiteral", "IntegerLiteral",
	"FloatingPointLiteral", "CharacterLiteral", "LeftParenthesis", "RightParenthesis",
	"LeftBracket", "RightBracket", "LeftBrace", "RightBrace", "HashedLeftBrace",
	"Hash", "Colon", "FractionalPart", "Int", "Symbol", "Name", "Comment",
	"Whitespace",
}

var ruleNames = []string{
	"program", "value", "tag", "keyword", "list_", "vector", "map_", "set_",
}

type ednParser struct {
	*antlr.BaseParser
}

// NewednParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ednParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewednParser(input antlr.TokenStream) *ednParser {
	this := new(ednParser)
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
	this.GrammarFileName = "edn.g4"

	return this
}

// ednParser tokens.
const (
	ednParserEOF                  = antlr.TokenEOF
	ednParserNilLiteral           = 1
	ednParserBooleanLiteral       = 2
	ednParserStringLiteral        = 3
	ednParserIntegerLiteral       = 4
	ednParserFloatingPointLiteral = 5
	ednParserCharacterLiteral     = 6
	ednParserLeftParenthesis      = 7
	ednParserRightParenthesis     = 8
	ednParserLeftBracket          = 9
	ednParserRightBracket         = 10
	ednParserLeftBrace            = 11
	ednParserRightBrace           = 12
	ednParserHashedLeftBrace      = 13
	ednParserHash                 = 14
	ednParserColon                = 15
	ednParserFractionalPart       = 16
	ednParserInt                  = 17
	ednParserSymbol               = 18
	ednParserName                 = 19
	ednParserComment              = 20
	ednParserWhitespace           = 21
)

// ednParser rules.
const (
	ednParserRULE_program = 0
	ednParserRULE_value   = 1
	ednParserRULE_tag     = 2
	ednParserRULE_keyword = 3
	ednParserRULE_list_   = 4
	ednParserRULE_vector  = 5
	ednParserRULE_map_    = 6
	ednParserRULE_set_    = 7
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ednParserEOF, 0)
}

func (s *ProgramContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ProgramContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ednParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ednParserRULE_program)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserStringLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserCharacterLiteral)|(1<<ednParserLeftParenthesis)|(1<<ednParserLeftBracket)|(1<<ednParserLeftBrace)|(1<<ednParserHashedLeftBrace)|(1<<ednParserHash)|(1<<ednParserColon)|(1<<ednParserSymbol))) != 0 {
		{
			p.SetState(16)
			p.Value()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(ednParserEOF)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NilLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserNilLiteral, 0)
}

func (s *ValueContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserBooleanLiteral, 0)
}

func (s *ValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserStringLiteral, 0)
}

func (s *ValueContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserIntegerLiteral, 0)
}

func (s *ValueContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserFloatingPointLiteral, 0)
}

func (s *ValueContext) CharacterLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserCharacterLiteral, 0)
}

func (s *ValueContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ValueContext) Symbol() antlr.TerminalNode {
	return s.GetToken(ednParserSymbol, 0)
}

func (s *ValueContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *ValueContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *ValueContext) Vector() IVectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *ValueContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *ValueContext) Set_() ISet_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_Context)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ednParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ednParserRULE_value)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ednParserNilLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(ednParserNilLiteral)
		}

	case ednParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(ednParserBooleanLiteral)
		}

	case ednParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(ednParserStringLiteral)
		}

	case ednParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Match(ednParserIntegerLiteral)
		}

	case ednParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.Match(ednParserFloatingPointLiteral)
		}

	case ednParserCharacterLiteral:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(29)
			p.Match(ednParserCharacterLiteral)
		}

	case ednParserColon:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(30)
			p.Keyword()
		}

	case ednParserSymbol:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(31)
			p.Match(ednParserSymbol)
		}

	case ednParserHash:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(32)
			p.Tag()
		}

	case ednParserLeftParenthesis:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(33)
			p.List_()
		}

	case ednParserLeftBracket:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(34)
			p.Vector()
		}

	case ednParserLeftBrace:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(35)
			p.Map_()
		}

	case ednParserHashedLeftBrace:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(36)
			p.Set_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = ednParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) Hash() antlr.TerminalNode {
	return s.GetToken(ednParserHash, 0)
}

func (s *TagContext) Symbol() antlr.TerminalNode {
	return s.GetToken(ednParserSymbol, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *ednParser) Tag() (localctx ITagContext) {
	this := p
	_ = this

	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ednParserRULE_tag)

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
		p.Match(ednParserHash)
	}
	{
		p.SetState(40)
		p.Match(ednParserSymbol)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Colon() antlr.TerminalNode {
	return s.GetToken(ednParserColon, 0)
}

func (s *KeywordContext) Symbol() antlr.TerminalNode {
	return s.GetToken(ednParserSymbol, 0)
}

func (s *KeywordContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserIntegerLiteral, 0)
}

func (s *KeywordContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserFloatingPointLiteral, 0)
}

func (s *KeywordContext) NilLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserNilLiteral, 0)
}

func (s *KeywordContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ednParserBooleanLiteral, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *ednParser) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ednParserRULE_keyword)
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
		p.SetState(42)
		p.Match(ednParserColon)
	}
	{
		p.SetState(43)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserSymbol))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(ednParserLeftParenthesis, 0)
}

func (s *List_Context) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(ednParserRightParenthesis, 0)
}

func (s *List_Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *List_Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *ednParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ednParserRULE_list_)
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
		p.SetState(45)
		p.Match(ednParserLeftParenthesis)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserStringLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserCharacterLiteral)|(1<<ednParserLeftParenthesis)|(1<<ednParserLeftBracket)|(1<<ednParserLeftBrace)|(1<<ednParserHashedLeftBrace)|(1<<ednParserHash)|(1<<ednParserColon)|(1<<ednParserSymbol))) != 0 {
		{
			p.SetState(46)
			p.Value()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(ednParserRightParenthesis)
	}

	return localctx
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_vector
	return p
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(ednParserLeftBracket, 0)
}

func (s *VectorContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(ednParserRightBracket, 0)
}

func (s *VectorContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *VectorContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *ednParser) Vector() (localctx IVectorContext) {
	this := p
	_ = this

	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ednParserRULE_vector)
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
		p.SetState(54)
		p.Match(ednParserLeftBracket)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserStringLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserCharacterLiteral)|(1<<ednParserLeftParenthesis)|(1<<ednParserLeftBracket)|(1<<ednParserLeftBrace)|(1<<ednParserHashedLeftBrace)|(1<<ednParserHash)|(1<<ednParserColon)|(1<<ednParserSymbol))) != 0 {
		{
			p.SetState(55)
			p.Value()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Match(ednParserRightBracket)
	}

	return localctx
}

// IMap_Context is an interface to support dynamic dispatch.
type IMap_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_Context differentiates from other interfaces.
	IsMap_Context()
}

type Map_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_Context() *Map_Context {
	var p = new(Map_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_map_
	return p
}

func (*Map_Context) IsMap_Context() {}

func NewMap_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_Context {
	var p = new(Map_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_map_

	return p
}

func (s *Map_Context) GetParser() antlr.Parser { return s.parser }

func (s *Map_Context) LeftBrace() antlr.TerminalNode {
	return s.GetToken(ednParserLeftBrace, 0)
}

func (s *Map_Context) RightBrace() antlr.TerminalNode {
	return s.GetToken(ednParserRightBrace, 0)
}

func (s *Map_Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Map_Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Map_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterMap_(s)
	}
}

func (s *Map_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitMap_(s)
	}
}

func (p *ednParser) Map_() (localctx IMap_Context) {
	this := p
	_ = this

	localctx = NewMap_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ednParserRULE_map_)
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
		p.SetState(63)
		p.Match(ednParserLeftBrace)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserStringLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserCharacterLiteral)|(1<<ednParserLeftParenthesis)|(1<<ednParserLeftBracket)|(1<<ednParserLeftBrace)|(1<<ednParserHashedLeftBrace)|(1<<ednParserHash)|(1<<ednParserColon)|(1<<ednParserSymbol))) != 0 {
		{
			p.SetState(64)
			p.Value()
		}
		{
			p.SetState(65)
			p.Value()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(ednParserRightBrace)
	}

	return localctx
}

// ISet_Context is an interface to support dynamic dispatch.
type ISet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_Context differentiates from other interfaces.
	IsSet_Context()
}

type Set_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_Context() *Set_Context {
	var p = new(Set_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ednParserRULE_set_
	return p
}

func (*Set_Context) IsSet_Context() {}

func NewSet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_Context {
	var p = new(Set_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ednParserRULE_set_

	return p
}

func (s *Set_Context) GetParser() antlr.Parser { return s.parser }

func (s *Set_Context) HashedLeftBrace() antlr.TerminalNode {
	return s.GetToken(ednParserHashedLeftBrace, 0)
}

func (s *Set_Context) RightBrace() antlr.TerminalNode {
	return s.GetToken(ednParserRightBrace, 0)
}

func (s *Set_Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Set_Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Set_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.EnterSet_(s)
	}
}

func (s *Set_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ednListener); ok {
		listenerT.ExitSet_(s)
	}
}

func (p *ednParser) Set_() (localctx ISet_Context) {
	this := p
	_ = this

	localctx = NewSet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ednParserRULE_set_)
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
		p.SetState(74)
		p.Match(ednParserHashedLeftBrace)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ednParserNilLiteral)|(1<<ednParserBooleanLiteral)|(1<<ednParserStringLiteral)|(1<<ednParserIntegerLiteral)|(1<<ednParserFloatingPointLiteral)|(1<<ednParserCharacterLiteral)|(1<<ednParserLeftParenthesis)|(1<<ednParserLeftBracket)|(1<<ednParserLeftBrace)|(1<<ednParserHashedLeftBrace)|(1<<ednParserHash)|(1<<ednParserColon)|(1<<ednParserSymbol))) != 0 {
		{
			p.SetState(75)
			p.Value()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(ednParserRightBrace)
	}

	return localctx
}
