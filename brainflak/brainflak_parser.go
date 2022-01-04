// Code generated from brainflak.g4 by ANTLR 4.9.3. DO NOT EDIT.

package brainflak // brainflak
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 62, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 6, 2, 16, 10, 2, 13, 2, 14, 2, 17, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 24,
	10, 3, 3, 4, 3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 7, 5, 37, 10, 5, 12, 5, 14, 5, 40, 11, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 7, 6, 46, 10, 6, 12, 6, 14, 6, 49, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	7, 7, 55, 10, 7, 12, 7, 14, 7, 58, 11, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2,
	4, 6, 8, 10, 12, 2, 2, 2, 63, 2, 15, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6,
	25, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 52, 3, 2, 2,
	2, 14, 16, 5, 4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15,
	3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 3, 3, 2, 2, 2, 19, 24, 5, 6, 4, 2,
	20, 24, 5, 8, 5, 2, 21, 24, 5, 10, 6, 2, 22, 24, 5, 12, 7, 2, 23, 19, 3,
	2, 2, 2, 23, 20, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 22, 3, 2, 2, 2, 24,
	5, 3, 2, 2, 2, 25, 29, 7, 3, 2, 2, 26, 28, 5, 4, 3, 2, 27, 26, 3, 2, 2,
	2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 32,
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 7, 4, 2, 2, 33, 7, 3, 2, 2, 2,
	34, 38, 7, 7, 2, 2, 35, 37, 5, 4, 3, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3,
	2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 41, 42, 7, 8, 2, 2, 42, 9, 3, 2, 2, 2, 43, 47, 7, 9, 2,
	2, 44, 46, 5, 4, 3, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45,
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2,
	50, 51, 7, 10, 2, 2, 51, 11, 3, 2, 2, 2, 52, 56, 7, 6, 2, 2, 53, 55, 5,
	4, 3, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56,
	57, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 5, 2,
	2, 60, 13, 3, 2, 2, 2, 8, 17, 23, 29, 38, 47, 56,
}
var literalNames = []string{
	"", "'('", "')'", "'>'", "'<'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "GT", "LT", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"WS",
}

var ruleNames = []string{
	"file_", "statement", "parenstmt", "bracestmt", "bracketstmt", "gtltstmt",
}

type brainflakParser struct {
	*antlr.BaseParser
}

// NewbrainflakParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *brainflakParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbrainflakParser(input antlr.TokenStream) *brainflakParser {
	this := new(brainflakParser)
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
	this.GrammarFileName = "brainflak.g4"

	return this
}

// brainflakParser tokens.
const (
	brainflakParserEOF    = antlr.TokenEOF
	brainflakParserLPAREN = 1
	brainflakParserRPAREN = 2
	brainflakParserGT     = 3
	brainflakParserLT     = 4
	brainflakParserLBRACE = 5
	brainflakParserRBRACE = 6
	brainflakParserLBRACK = 7
	brainflakParserRBRACK = 8
	brainflakParserWS     = 9
)

// brainflakParser rules.
const (
	brainflakParserRULE_file_       = 0
	brainflakParserRULE_statement   = 1
	brainflakParserRULE_parenstmt   = 2
	brainflakParserRULE_bracestmt   = 3
	brainflakParserRULE_bracketstmt = 4
	brainflakParserRULE_gtltstmt    = 5
)

// IFile_Context is an interface to support dynamic dispatch.
type IFile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_Context differentiates from other interfaces.
	IsFile_Context()
}

type File_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_Context() *File_Context {
	var p = new(File_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_file_
	return p
}

func (*File_Context) IsFile_Context() {}

func NewFile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_Context {
	var p = new(File_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_file_

	return p
}

func (s *File_Context) GetParser() antlr.Parser { return s.parser }

func (s *File_Context) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *File_Context) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *File_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterFile_(s)
	}
}

func (s *File_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitFile_(s)
	}
}

func (p *brainflakParser) File_() (localctx IFile_Context) {
	this := p
	_ = this

	localctx = NewFile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, brainflakParserRULE_file_)
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainflakParserLPAREN)|(1<<brainflakParserLT)|(1<<brainflakParserLBRACE)|(1<<brainflakParserLBRACK))) != 0) {
		{
			p.SetState(12)
			p.Statement()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Parenstmt() IParenstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenstmtContext)
}

func (s *StatementContext) Bracestmt() IBracestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracestmtContext)
}

func (s *StatementContext) Bracketstmt() IBracketstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketstmtContext)
}

func (s *StatementContext) Gtltstmt() IGtltstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtltstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtltstmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *brainflakParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, brainflakParserRULE_statement)

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

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case brainflakParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.Parenstmt()
		}

	case brainflakParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.Bracestmt()
		}

	case brainflakParserLBRACK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)
			p.Bracketstmt()
		}

	case brainflakParserLT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(20)
			p.Gtltstmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParenstmtContext is an interface to support dynamic dispatch.
type IParenstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParenstmtContext differentiates from other interfaces.
	IsParenstmtContext()
}

type ParenstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenstmtContext() *ParenstmtContext {
	var p = new(ParenstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_parenstmt
	return p
}

func (*ParenstmtContext) IsParenstmtContext() {}

func NewParenstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenstmtContext {
	var p = new(ParenstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_parenstmt

	return p
}

func (s *ParenstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenstmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(brainflakParserLPAREN, 0)
}

func (s *ParenstmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(brainflakParserRPAREN, 0)
}

func (s *ParenstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ParenstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ParenstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterParenstmt(s)
	}
}

func (s *ParenstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitParenstmt(s)
	}
}

func (p *brainflakParser) Parenstmt() (localctx IParenstmtContext) {
	this := p
	_ = this

	localctx = NewParenstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, brainflakParserRULE_parenstmt)
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
		p.SetState(23)
		p.Match(brainflakParserLPAREN)
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainflakParserLPAREN)|(1<<brainflakParserLT)|(1<<brainflakParserLBRACE)|(1<<brainflakParserLBRACK))) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(brainflakParserRPAREN)
	}

	return localctx
}

// IBracestmtContext is an interface to support dynamic dispatch.
type IBracestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracestmtContext differentiates from other interfaces.
	IsBracestmtContext()
}

type BracestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracestmtContext() *BracestmtContext {
	var p = new(BracestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_bracestmt
	return p
}

func (*BracestmtContext) IsBracestmtContext() {}

func NewBracestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracestmtContext {
	var p = new(BracestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_bracestmt

	return p
}

func (s *BracestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BracestmtContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(brainflakParserLBRACE, 0)
}

func (s *BracestmtContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(brainflakParserRBRACE, 0)
}

func (s *BracestmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BracestmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BracestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterBracestmt(s)
	}
}

func (s *BracestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitBracestmt(s)
	}
}

func (p *brainflakParser) Bracestmt() (localctx IBracestmtContext) {
	this := p
	_ = this

	localctx = NewBracestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, brainflakParserRULE_bracestmt)
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
		p.Match(brainflakParserLBRACE)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainflakParserLPAREN)|(1<<brainflakParserLT)|(1<<brainflakParserLBRACE)|(1<<brainflakParserLBRACK))) != 0 {
		{
			p.SetState(33)
			p.Statement()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(brainflakParserRBRACE)
	}

	return localctx
}

// IBracketstmtContext is an interface to support dynamic dispatch.
type IBracketstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketstmtContext differentiates from other interfaces.
	IsBracketstmtContext()
}

type BracketstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketstmtContext() *BracketstmtContext {
	var p = new(BracketstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_bracketstmt
	return p
}

func (*BracketstmtContext) IsBracketstmtContext() {}

func NewBracketstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketstmtContext {
	var p = new(BracketstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_bracketstmt

	return p
}

func (s *BracketstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketstmtContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(brainflakParserLBRACK, 0)
}

func (s *BracketstmtContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(brainflakParserRBRACK, 0)
}

func (s *BracketstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BracketstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BracketstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterBracketstmt(s)
	}
}

func (s *BracketstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitBracketstmt(s)
	}
}

func (p *brainflakParser) Bracketstmt() (localctx IBracketstmtContext) {
	this := p
	_ = this

	localctx = NewBracketstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, brainflakParserRULE_bracketstmt)
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
		p.SetState(41)
		p.Match(brainflakParserLBRACK)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainflakParserLPAREN)|(1<<brainflakParserLT)|(1<<brainflakParserLBRACE)|(1<<brainflakParserLBRACK))) != 0 {
		{
			p.SetState(42)
			p.Statement()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(brainflakParserRBRACK)
	}

	return localctx
}

// IGtltstmtContext is an interface to support dynamic dispatch.
type IGtltstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtltstmtContext differentiates from other interfaces.
	IsGtltstmtContext()
}

type GtltstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtltstmtContext() *GtltstmtContext {
	var p = new(GtltstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = brainflakParserRULE_gtltstmt
	return p
}

func (*GtltstmtContext) IsGtltstmtContext() {}

func NewGtltstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GtltstmtContext {
	var p = new(GtltstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = brainflakParserRULE_gtltstmt

	return p
}

func (s *GtltstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GtltstmtContext) LT() antlr.TerminalNode {
	return s.GetToken(brainflakParserLT, 0)
}

func (s *GtltstmtContext) GT() antlr.TerminalNode {
	return s.GetToken(brainflakParserGT, 0)
}

func (s *GtltstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *GtltstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *GtltstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtltstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GtltstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.EnterGtltstmt(s)
	}
}

func (s *GtltstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(brainflakListener); ok {
		listenerT.ExitGtltstmt(s)
	}
}

func (p *brainflakParser) Gtltstmt() (localctx IGtltstmtContext) {
	this := p
	_ = this

	localctx = NewGtltstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, brainflakParserRULE_gtltstmt)
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
		p.SetState(50)
		p.Match(brainflakParserLT)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<brainflakParserLPAREN)|(1<<brainflakParserLT)|(1<<brainflakParserLBRACE)|(1<<brainflakParserLBRACK))) != 0 {
		{
			p.SetState(51)
			p.Statement()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(brainflakParserGT)
	}

	return localctx
}
