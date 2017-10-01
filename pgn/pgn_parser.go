// Generated from PGN.g4 by ANTLR 4.7.

package pgn // PGN
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 85, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 35, 10, 3,
	12, 3, 14, 3, 38, 11, 3, 3, 4, 3, 4, 3, 4, 3, 5, 7, 5, 44, 10, 5, 12, 5,
	14, 5, 47, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 7, 10, 63, 10, 10, 12, 10, 14, 10, 66,
	11, 10, 3, 11, 3, 11, 3, 11, 5, 11, 71, 10, 11, 3, 12, 3, 12, 5, 12, 75,
	10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 3, 4,
	2, 3, 5, 13, 13, 2, 77, 2, 30, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 39, 3,
	2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 53, 3, 2, 2, 2, 14,
	55, 3, 2, 2, 2, 16, 57, 3, 2, 2, 2, 18, 64, 3, 2, 2, 2, 20, 70, 3, 2, 2,
	2, 22, 72, 3, 2, 2, 2, 24, 76, 3, 2, 2, 2, 26, 78, 3, 2, 2, 2, 28, 82,
	3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3, 32, 3, 3, 2, 2, 2,
	33, 35, 5, 6, 4, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 5, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39,
	40, 5, 8, 5, 2, 40, 41, 5, 16, 9, 2, 41, 7, 3, 2, 2, 2, 42, 44, 5, 10,
	6, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46,
	3, 2, 2, 2, 46, 9, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 14, 2, 2,
	49, 50, 5, 12, 7, 2, 50, 51, 5, 14, 8, 2, 51, 52, 7, 15, 2, 2, 52, 11,
	3, 2, 2, 2, 53, 54, 7, 21, 2, 2, 54, 13, 3, 2, 2, 2, 55, 56, 7, 10, 2,
	2, 56, 15, 3, 2, 2, 2, 57, 58, 5, 18, 10, 2, 58, 59, 5, 28, 15, 2, 59,
	17, 3, 2, 2, 2, 60, 63, 5, 20, 11, 2, 61, 63, 5, 26, 14, 2, 62, 60, 3,
	2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64,
	65, 3, 2, 2, 2, 65, 19, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 71, 5, 22,
	12, 2, 68, 71, 5, 24, 13, 2, 69, 71, 7, 20, 2, 2, 70, 67, 3, 2, 2, 2, 70,
	68, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 21, 3, 2, 2, 2, 72, 74, 7, 11,
	2, 2, 73, 75, 7, 12, 2, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	23, 3, 2, 2, 2, 76, 77, 7, 21, 2, 2, 77, 25, 3, 2, 2, 2, 78, 79, 7, 16,
	2, 2, 79, 80, 5, 18, 10, 2, 80, 81, 7, 17, 2, 2, 81, 27, 3, 2, 2, 2, 82,
	83, 9, 2, 2, 2, 83, 29, 3, 2, 2, 2, 8, 36, 45, 62, 64, 70, 74,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'1-0'", "'0-1'", "'1/2-1/2'", "", "", "", "", "", "", "'.'", "'*'",
	"'['", "']'", "'('", "')'", "'<'", "'>'",
}
var symbolicNames = []string{
	"", "WHITE_WINS", "BLACK_WINS", "DRAWN_GAME", "REST_OF_LINE_COMMENT", "BRACE_COMMENT",
	"ESCAPE", "SPACES", "STRING", "INTEGER", "PERIOD", "ASTERISK", "LEFT_BRACKET",
	"RIGHT_BRACKET", "LEFT_PARENTHESIS", "RIGHT_PARENTHESIS", "LEFT_ANGLE_BRACKET",
	"RIGHT_ANGLE_BRACKET", "NUMERIC_ANNOTATION_GLYPH", "SYMBOL", "SUFFIX_ANNOTATION",
	"UNEXPECTED_CHAR",
}

var ruleNames = []string{
	"parse", "pgn_database", "pgn_game", "tag_section", "tag_pair", "tag_name",
	"tag_value", "movetext_section", "element_sequence", "element", "move_number_indication",
	"san_move", "recursive_variation", "game_termination",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PGNParser struct {
	*antlr.BaseParser
}

func NewPGNParser(input antlr.TokenStream) *PGNParser {
	this := new(PGNParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PGN.g4"

	return this
}

// PGNParser tokens.
const (
	PGNParserEOF                      = antlr.TokenEOF
	PGNParserWHITE_WINS               = 1
	PGNParserBLACK_WINS               = 2
	PGNParserDRAWN_GAME               = 3
	PGNParserREST_OF_LINE_COMMENT     = 4
	PGNParserBRACE_COMMENT            = 5
	PGNParserESCAPE                   = 6
	PGNParserSPACES                   = 7
	PGNParserSTRING                   = 8
	PGNParserINTEGER                  = 9
	PGNParserPERIOD                   = 10
	PGNParserASTERISK                 = 11
	PGNParserLEFT_BRACKET             = 12
	PGNParserRIGHT_BRACKET            = 13
	PGNParserLEFT_PARENTHESIS         = 14
	PGNParserRIGHT_PARENTHESIS        = 15
	PGNParserLEFT_ANGLE_BRACKET       = 16
	PGNParserRIGHT_ANGLE_BRACKET      = 17
	PGNParserNUMERIC_ANNOTATION_GLYPH = 18
	PGNParserSYMBOL                   = 19
	PGNParserSUFFIX_ANNOTATION        = 20
	PGNParserUNEXPECTED_CHAR          = 21
)

// PGNParser rules.
const (
	PGNParserRULE_parse                  = 0
	PGNParserRULE_pgn_database           = 1
	PGNParserRULE_pgn_game               = 2
	PGNParserRULE_tag_section            = 3
	PGNParserRULE_tag_pair               = 4
	PGNParserRULE_tag_name               = 5
	PGNParserRULE_tag_value              = 6
	PGNParserRULE_movetext_section       = 7
	PGNParserRULE_element_sequence       = 8
	PGNParserRULE_element                = 9
	PGNParserRULE_move_number_indication = 10
	PGNParserRULE_san_move               = 11
	PGNParserRULE_recursive_variation    = 12
	PGNParserRULE_game_termination       = 13
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Pgn_database() IPgn_databaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPgn_databaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPgn_databaseContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(PGNParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PGNParserRULE_parse)

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
		p.Pgn_database()
	}
	{
		p.SetState(29)
		p.Match(PGNParserEOF)
	}

	return localctx
}

// IPgn_databaseContext is an interface to support dynamic dispatch.
type IPgn_databaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPgn_databaseContext differentiates from other interfaces.
	IsPgn_databaseContext()
}

type Pgn_databaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPgn_databaseContext() *Pgn_databaseContext {
	var p = new(Pgn_databaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_pgn_database
	return p
}

func (*Pgn_databaseContext) IsPgn_databaseContext() {}

func NewPgn_databaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pgn_databaseContext {
	var p = new(Pgn_databaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_pgn_database

	return p
}

func (s *Pgn_databaseContext) GetParser() antlr.Parser { return s.parser }

func (s *Pgn_databaseContext) AllPgn_game() []IPgn_gameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPgn_gameContext)(nil)).Elem())
	var tst = make([]IPgn_gameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPgn_gameContext)
		}
	}

	return tst
}

func (s *Pgn_databaseContext) Pgn_game(i int) IPgn_gameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPgn_gameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPgn_gameContext)
}

func (s *Pgn_databaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pgn_databaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pgn_databaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterPgn_database(s)
	}
}

func (s *Pgn_databaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitPgn_database(s)
	}
}

func (s *Pgn_databaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitPgn_database(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Pgn_database() (localctx IPgn_databaseContext) {
	localctx = NewPgn_databaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PGNParserRULE_pgn_database)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PGNParserWHITE_WINS)|(1<<PGNParserBLACK_WINS)|(1<<PGNParserDRAWN_GAME)|(1<<PGNParserINTEGER)|(1<<PGNParserASTERISK)|(1<<PGNParserLEFT_BRACKET)|(1<<PGNParserLEFT_PARENTHESIS)|(1<<PGNParserNUMERIC_ANNOTATION_GLYPH)|(1<<PGNParserSYMBOL))) != 0 {
		{
			p.SetState(31)
			p.Pgn_game()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPgn_gameContext is an interface to support dynamic dispatch.
type IPgn_gameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPgn_gameContext differentiates from other interfaces.
	IsPgn_gameContext()
}

type Pgn_gameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPgn_gameContext() *Pgn_gameContext {
	var p = new(Pgn_gameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_pgn_game
	return p
}

func (*Pgn_gameContext) IsPgn_gameContext() {}

func NewPgn_gameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pgn_gameContext {
	var p = new(Pgn_gameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_pgn_game

	return p
}

func (s *Pgn_gameContext) GetParser() antlr.Parser { return s.parser }

func (s *Pgn_gameContext) Tag_section() ITag_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITag_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITag_sectionContext)
}

func (s *Pgn_gameContext) Movetext_section() IMovetext_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMovetext_sectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMovetext_sectionContext)
}

func (s *Pgn_gameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pgn_gameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pgn_gameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterPgn_game(s)
	}
}

func (s *Pgn_gameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitPgn_game(s)
	}
}

func (s *Pgn_gameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitPgn_game(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Pgn_game() (localctx IPgn_gameContext) {
	localctx = NewPgn_gameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PGNParserRULE_pgn_game)

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
		p.Tag_section()
	}
	{
		p.SetState(38)
		p.Movetext_section()
	}

	return localctx
}

// ITag_sectionContext is an interface to support dynamic dispatch.
type ITag_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTag_sectionContext differentiates from other interfaces.
	IsTag_sectionContext()
}

type Tag_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTag_sectionContext() *Tag_sectionContext {
	var p = new(Tag_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_tag_section
	return p
}

func (*Tag_sectionContext) IsTag_sectionContext() {}

func NewTag_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tag_sectionContext {
	var p = new(Tag_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_tag_section

	return p
}

func (s *Tag_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Tag_sectionContext) AllTag_pair() []ITag_pairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITag_pairContext)(nil)).Elem())
	var tst = make([]ITag_pairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITag_pairContext)
		}
	}

	return tst
}

func (s *Tag_sectionContext) Tag_pair(i int) ITag_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITag_pairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITag_pairContext)
}

func (s *Tag_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tag_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tag_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterTag_section(s)
	}
}

func (s *Tag_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitTag_section(s)
	}
}

func (s *Tag_sectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitTag_section(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Tag_section() (localctx ITag_sectionContext) {
	localctx = NewTag_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PGNParserRULE_tag_section)
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PGNParserLEFT_BRACKET {
		{
			p.SetState(40)
			p.Tag_pair()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITag_pairContext is an interface to support dynamic dispatch.
type ITag_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTag_pairContext differentiates from other interfaces.
	IsTag_pairContext()
}

type Tag_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTag_pairContext() *Tag_pairContext {
	var p = new(Tag_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_tag_pair
	return p
}

func (*Tag_pairContext) IsTag_pairContext() {}

func NewTag_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tag_pairContext {
	var p = new(Tag_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_tag_pair

	return p
}

func (s *Tag_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Tag_pairContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PGNParserLEFT_BRACKET, 0)
}

func (s *Tag_pairContext) Tag_name() ITag_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITag_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITag_nameContext)
}

func (s *Tag_pairContext) Tag_value() ITag_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITag_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITag_valueContext)
}

func (s *Tag_pairContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PGNParserRIGHT_BRACKET, 0)
}

func (s *Tag_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tag_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tag_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterTag_pair(s)
	}
}

func (s *Tag_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitTag_pair(s)
	}
}

func (s *Tag_pairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitTag_pair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Tag_pair() (localctx ITag_pairContext) {
	localctx = NewTag_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PGNParserRULE_tag_pair)

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
		p.Match(PGNParserLEFT_BRACKET)
	}
	{
		p.SetState(47)
		p.Tag_name()
	}
	{
		p.SetState(48)
		p.Tag_value()
	}
	{
		p.SetState(49)
		p.Match(PGNParserRIGHT_BRACKET)
	}

	return localctx
}

// ITag_nameContext is an interface to support dynamic dispatch.
type ITag_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTag_nameContext differentiates from other interfaces.
	IsTag_nameContext()
}

type Tag_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTag_nameContext() *Tag_nameContext {
	var p = new(Tag_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_tag_name
	return p
}

func (*Tag_nameContext) IsTag_nameContext() {}

func NewTag_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tag_nameContext {
	var p = new(Tag_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_tag_name

	return p
}

func (s *Tag_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Tag_nameContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(PGNParserSYMBOL, 0)
}

func (s *Tag_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tag_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tag_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterTag_name(s)
	}
}

func (s *Tag_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitTag_name(s)
	}
}

func (s *Tag_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitTag_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Tag_name() (localctx ITag_nameContext) {
	localctx = NewTag_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PGNParserRULE_tag_name)

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
		p.Match(PGNParserSYMBOL)
	}

	return localctx
}

// ITag_valueContext is an interface to support dynamic dispatch.
type ITag_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTag_valueContext differentiates from other interfaces.
	IsTag_valueContext()
}

type Tag_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTag_valueContext() *Tag_valueContext {
	var p = new(Tag_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_tag_value
	return p
}

func (*Tag_valueContext) IsTag_valueContext() {}

func NewTag_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tag_valueContext {
	var p = new(Tag_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_tag_value

	return p
}

func (s *Tag_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Tag_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PGNParserSTRING, 0)
}

func (s *Tag_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tag_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tag_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterTag_value(s)
	}
}

func (s *Tag_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitTag_value(s)
	}
}

func (s *Tag_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitTag_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Tag_value() (localctx ITag_valueContext) {
	localctx = NewTag_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PGNParserRULE_tag_value)

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
		p.Match(PGNParserSTRING)
	}

	return localctx
}

// IMovetext_sectionContext is an interface to support dynamic dispatch.
type IMovetext_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMovetext_sectionContext differentiates from other interfaces.
	IsMovetext_sectionContext()
}

type Movetext_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMovetext_sectionContext() *Movetext_sectionContext {
	var p = new(Movetext_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_movetext_section
	return p
}

func (*Movetext_sectionContext) IsMovetext_sectionContext() {}

func NewMovetext_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Movetext_sectionContext {
	var p = new(Movetext_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_movetext_section

	return p
}

func (s *Movetext_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Movetext_sectionContext) Element_sequence() IElement_sequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElement_sequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElement_sequenceContext)
}

func (s *Movetext_sectionContext) Game_termination() IGame_terminationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGame_terminationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGame_terminationContext)
}

func (s *Movetext_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Movetext_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Movetext_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterMovetext_section(s)
	}
}

func (s *Movetext_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitMovetext_section(s)
	}
}

func (s *Movetext_sectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitMovetext_section(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Movetext_section() (localctx IMovetext_sectionContext) {
	localctx = NewMovetext_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PGNParserRULE_movetext_section)

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
		p.SetState(55)
		p.Element_sequence()
	}
	{
		p.SetState(56)
		p.Game_termination()
	}

	return localctx
}

// IElement_sequenceContext is an interface to support dynamic dispatch.
type IElement_sequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElement_sequenceContext differentiates from other interfaces.
	IsElement_sequenceContext()
}

type Element_sequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElement_sequenceContext() *Element_sequenceContext {
	var p = new(Element_sequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_element_sequence
	return p
}

func (*Element_sequenceContext) IsElement_sequenceContext() {}

func NewElement_sequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Element_sequenceContext {
	var p = new(Element_sequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_element_sequence

	return p
}

func (s *Element_sequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Element_sequenceContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *Element_sequenceContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *Element_sequenceContext) AllRecursive_variation() []IRecursive_variationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRecursive_variationContext)(nil)).Elem())
	var tst = make([]IRecursive_variationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRecursive_variationContext)
		}
	}

	return tst
}

func (s *Element_sequenceContext) Recursive_variation(i int) IRecursive_variationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecursive_variationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRecursive_variationContext)
}

func (s *Element_sequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Element_sequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Element_sequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterElement_sequence(s)
	}
}

func (s *Element_sequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitElement_sequence(s)
	}
}

func (s *Element_sequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitElement_sequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Element_sequence() (localctx IElement_sequenceContext) {
	localctx = NewElement_sequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PGNParserRULE_element_sequence)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PGNParserINTEGER)|(1<<PGNParserLEFT_PARENTHESIS)|(1<<PGNParserNUMERIC_ANNOTATION_GLYPH)|(1<<PGNParserSYMBOL))) != 0 {
		p.SetState(60)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PGNParserINTEGER, PGNParserNUMERIC_ANNOTATION_GLYPH, PGNParserSYMBOL:
			{
				p.SetState(58)
				p.Element()
			}

		case PGNParserLEFT_PARENTHESIS:
			{
				p.SetState(59)
				p.Recursive_variation()
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

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Move_number_indication() IMove_number_indicationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMove_number_indicationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMove_number_indicationContext)
}

func (s *ElementContext) San_move() ISan_moveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISan_moveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISan_moveContext)
}

func (s *ElementContext) NUMERIC_ANNOTATION_GLYPH() antlr.TerminalNode {
	return s.GetToken(PGNParserNUMERIC_ANNOTATION_GLYPH, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PGNParserRULE_element)

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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PGNParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Move_number_indication()
		}

	case PGNParserSYMBOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.San_move()
		}

	case PGNParserNUMERIC_ANNOTATION_GLYPH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(PGNParserNUMERIC_ANNOTATION_GLYPH)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMove_number_indicationContext is an interface to support dynamic dispatch.
type IMove_number_indicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMove_number_indicationContext differentiates from other interfaces.
	IsMove_number_indicationContext()
}

type Move_number_indicationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMove_number_indicationContext() *Move_number_indicationContext {
	var p = new(Move_number_indicationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_move_number_indication
	return p
}

func (*Move_number_indicationContext) IsMove_number_indicationContext() {}

func NewMove_number_indicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Move_number_indicationContext {
	var p = new(Move_number_indicationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_move_number_indication

	return p
}

func (s *Move_number_indicationContext) GetParser() antlr.Parser { return s.parser }

func (s *Move_number_indicationContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PGNParserINTEGER, 0)
}

func (s *Move_number_indicationContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(PGNParserPERIOD, 0)
}

func (s *Move_number_indicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Move_number_indicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Move_number_indicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterMove_number_indication(s)
	}
}

func (s *Move_number_indicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitMove_number_indication(s)
	}
}

func (s *Move_number_indicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitMove_number_indication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Move_number_indication() (localctx IMove_number_indicationContext) {
	localctx = NewMove_number_indicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PGNParserRULE_move_number_indication)
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
		p.SetState(70)
		p.Match(PGNParserINTEGER)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PGNParserPERIOD {
		{
			p.SetState(71)
			p.Match(PGNParserPERIOD)
		}

	}

	return localctx
}

// ISan_moveContext is an interface to support dynamic dispatch.
type ISan_moveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSan_moveContext differentiates from other interfaces.
	IsSan_moveContext()
}

type San_moveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySan_moveContext() *San_moveContext {
	var p = new(San_moveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_san_move
	return p
}

func (*San_moveContext) IsSan_moveContext() {}

func NewSan_moveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *San_moveContext {
	var p = new(San_moveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_san_move

	return p
}

func (s *San_moveContext) GetParser() antlr.Parser { return s.parser }

func (s *San_moveContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(PGNParserSYMBOL, 0)
}

func (s *San_moveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *San_moveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *San_moveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterSan_move(s)
	}
}

func (s *San_moveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitSan_move(s)
	}
}

func (s *San_moveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitSan_move(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) San_move() (localctx ISan_moveContext) {
	localctx = NewSan_moveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PGNParserRULE_san_move)

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
		p.Match(PGNParserSYMBOL)
	}

	return localctx
}

// IRecursive_variationContext is an interface to support dynamic dispatch.
type IRecursive_variationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecursive_variationContext differentiates from other interfaces.
	IsRecursive_variationContext()
}

type Recursive_variationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecursive_variationContext() *Recursive_variationContext {
	var p = new(Recursive_variationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_recursive_variation
	return p
}

func (*Recursive_variationContext) IsRecursive_variationContext() {}

func NewRecursive_variationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Recursive_variationContext {
	var p = new(Recursive_variationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_recursive_variation

	return p
}

func (s *Recursive_variationContext) GetParser() antlr.Parser { return s.parser }

func (s *Recursive_variationContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PGNParserLEFT_PARENTHESIS, 0)
}

func (s *Recursive_variationContext) Element_sequence() IElement_sequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElement_sequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElement_sequenceContext)
}

func (s *Recursive_variationContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PGNParserRIGHT_PARENTHESIS, 0)
}

func (s *Recursive_variationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Recursive_variationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Recursive_variationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterRecursive_variation(s)
	}
}

func (s *Recursive_variationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitRecursive_variation(s)
	}
}

func (s *Recursive_variationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitRecursive_variation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Recursive_variation() (localctx IRecursive_variationContext) {
	localctx = NewRecursive_variationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PGNParserRULE_recursive_variation)

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
		p.SetState(76)
		p.Match(PGNParserLEFT_PARENTHESIS)
	}
	{
		p.SetState(77)
		p.Element_sequence()
	}
	{
		p.SetState(78)
		p.Match(PGNParserRIGHT_PARENTHESIS)
	}

	return localctx
}

// IGame_terminationContext is an interface to support dynamic dispatch.
type IGame_terminationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGame_terminationContext differentiates from other interfaces.
	IsGame_terminationContext()
}

type Game_terminationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGame_terminationContext() *Game_terminationContext {
	var p = new(Game_terminationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PGNParserRULE_game_termination
	return p
}

func (*Game_terminationContext) IsGame_terminationContext() {}

func NewGame_terminationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Game_terminationContext {
	var p = new(Game_terminationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PGNParserRULE_game_termination

	return p
}

func (s *Game_terminationContext) GetParser() antlr.Parser { return s.parser }

func (s *Game_terminationContext) WHITE_WINS() antlr.TerminalNode {
	return s.GetToken(PGNParserWHITE_WINS, 0)
}

func (s *Game_terminationContext) BLACK_WINS() antlr.TerminalNode {
	return s.GetToken(PGNParserBLACK_WINS, 0)
}

func (s *Game_terminationContext) DRAWN_GAME() antlr.TerminalNode {
	return s.GetToken(PGNParserDRAWN_GAME, 0)
}

func (s *Game_terminationContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(PGNParserASTERISK, 0)
}

func (s *Game_terminationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Game_terminationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Game_terminationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.EnterGame_termination(s)
	}
}

func (s *Game_terminationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PGNListener); ok {
		listenerT.ExitGame_termination(s)
	}
}

func (s *Game_terminationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PGNVisitor:
		return t.VisitGame_termination(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PGNParser) Game_termination() (localctx IGame_terminationContext) {
	localctx = NewGame_terminationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PGNParserRULE_game_termination)
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
	p.SetState(80)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PGNParserWHITE_WINS)|(1<<PGNParserBLACK_WINS)|(1<<PGNParserDRAWN_GAME)|(1<<PGNParserASTERISK))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
