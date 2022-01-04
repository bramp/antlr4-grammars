// Code generated from loop.g4 by ANTLR 4.9.3. DO NOT EDIT.

package loop // loop
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 5, 3, 29, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 34, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2,
	4, 6, 8, 10, 12, 14, 16, 2, 3, 3, 2, 5, 6, 2, 52, 2, 19, 3, 2, 2, 2, 4,
	25, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2, 10, 40, 3, 2, 2,
	2, 12, 46, 3, 2, 2, 2, 14, 52, 3, 2, 2, 2, 16, 54, 3, 2, 2, 2, 18, 20,
	5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2,
	21, 22, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 24, 7, 2, 2, 3, 24, 3, 3, 2,
	2, 2, 25, 28, 5, 6, 4, 2, 26, 27, 7, 3, 2, 2, 27, 29, 5, 6, 4, 2, 28, 26,
	3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 34, 5, 8, 5, 2,
	31, 34, 5, 10, 6, 2, 32, 34, 5, 12, 7, 2, 33, 30, 3, 2, 2, 2, 33, 31, 3,
	2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 7, 3, 2, 2, 2, 35, 36, 5, 14, 8, 2, 36,
	37, 7, 4, 2, 2, 37, 38, 5, 16, 9, 2, 38, 39, 7, 3, 2, 2, 39, 9, 3, 2, 2,
	2, 40, 41, 5, 14, 8, 2, 41, 42, 7, 4, 2, 2, 42, 43, 5, 14, 8, 2, 43, 44,
	9, 2, 2, 2, 44, 45, 5, 16, 9, 2, 45, 11, 3, 2, 2, 2, 46, 47, 7, 7, 2, 2,
	47, 48, 5, 14, 8, 2, 48, 49, 7, 8, 2, 2, 49, 50, 5, 4, 3, 2, 50, 51, 7,
	9, 2, 2, 51, 13, 3, 2, 2, 2, 52, 53, 7, 10, 2, 2, 53, 15, 3, 2, 2, 2, 54,
	55, 7, 11, 2, 2, 55, 17, 3, 2, 2, 2, 5, 21, 28, 33,
}
var literalNames = []string{
	"", "';'", "':='", "'+'", "'-'", "'LOOP'", "'DO'", "'END'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "ID", "NUMBER", "COMMENT", "WS",
}

var ruleNames = []string{
	"prog", "statementlist", "statement", "assignstmt", "incrementstmt", "loopstmt",
	"var_", "number",
}

type loopParser struct {
	*antlr.BaseParser
}

// NewloopParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *loopParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewloopParser(input antlr.TokenStream) *loopParser {
	this := new(loopParser)
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
	this.GrammarFileName = "loop.g4"

	return this
}

// loopParser tokens.
const (
	loopParserEOF     = antlr.TokenEOF
	loopParserT__0    = 1
	loopParserT__1    = 2
	loopParserT__2    = 3
	loopParserT__3    = 4
	loopParserT__4    = 5
	loopParserT__5    = 6
	loopParserT__6    = 7
	loopParserID      = 8
	loopParserNUMBER  = 9
	loopParserCOMMENT = 10
	loopParserWS      = 11
)

// loopParser rules.
const (
	loopParserRULE_prog          = 0
	loopParserRULE_statementlist = 1
	loopParserRULE_statement     = 2
	loopParserRULE_assignstmt    = 3
	loopParserRULE_incrementstmt = 4
	loopParserRULE_loopstmt      = 5
	loopParserRULE_var_          = 6
	loopParserRULE_number        = 7
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(loopParserEOF, 0)
}

func (s *ProgContext) AllStatementlist() []IStatementlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementlistContext)(nil)).Elem())
	var tst = make([]IStatementlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementlistContext)
		}
	}

	return tst
}

func (s *ProgContext) Statementlist(i int) IStatementlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementlistContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *loopParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, loopParserRULE_prog)
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

	for ok := true; ok; ok = _la == loopParserT__4 || _la == loopParserID {
		{
			p.SetState(16)
			p.Statementlist()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(loopParserEOF)
	}

	return localctx
}

// IStatementlistContext is an interface to support dynamic dispatch.
type IStatementlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementlistContext differentiates from other interfaces.
	IsStatementlistContext()
}

type StatementlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementlistContext() *StatementlistContext {
	var p = new(StatementlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_statementlist
	return p
}

func (*StatementlistContext) IsStatementlistContext() {}

func NewStatementlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementlistContext {
	var p = new(StatementlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_statementlist

	return p
}

func (s *StatementlistContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementlistContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementlistContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterStatementlist(s)
	}
}

func (s *StatementlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitStatementlist(s)
	}
}

func (p *loopParser) Statementlist() (localctx IStatementlistContext) {
	this := p
	_ = this

	localctx = NewStatementlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, loopParserRULE_statementlist)
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
		p.Statement()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == loopParserT__0 {
		{
			p.SetState(24)
			p.Match(loopParserT__0)
		}
		{
			p.SetState(25)
			p.Statement()
		}

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
	p.RuleIndex = loopParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignstmt() IAssignstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignstmtContext)
}

func (s *StatementContext) Incrementstmt() IIncrementstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncrementstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncrementstmtContext)
}

func (s *StatementContext) Loopstmt() ILoopstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopstmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *loopParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, loopParserRULE_statement)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Assignstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Incrementstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Loopstmt()
		}

	}

	return localctx
}

// IAssignstmtContext is an interface to support dynamic dispatch.
type IAssignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignstmtContext differentiates from other interfaces.
	IsAssignstmtContext()
}

type AssignstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstmtContext() *AssignstmtContext {
	var p = new(AssignstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_assignstmt
	return p
}

func (*AssignstmtContext) IsAssignstmtContext() {}

func NewAssignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstmtContext {
	var p = new(AssignstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_assignstmt

	return p
}

func (s *AssignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstmtContext) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *AssignstmtContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AssignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterAssignstmt(s)
	}
}

func (s *AssignstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitAssignstmt(s)
	}
}

func (p *loopParser) Assignstmt() (localctx IAssignstmtContext) {
	this := p
	_ = this

	localctx = NewAssignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, loopParserRULE_assignstmt)

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
		p.SetState(33)
		p.Var_()
	}
	{
		p.SetState(34)
		p.Match(loopParserT__1)
	}
	{
		p.SetState(35)
		p.Number()
	}
	{
		p.SetState(36)
		p.Match(loopParserT__0)
	}

	return localctx
}

// IIncrementstmtContext is an interface to support dynamic dispatch.
type IIncrementstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncrementstmtContext differentiates from other interfaces.
	IsIncrementstmtContext()
}

type IncrementstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrementstmtContext() *IncrementstmtContext {
	var p = new(IncrementstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_incrementstmt
	return p
}

func (*IncrementstmtContext) IsIncrementstmtContext() {}

func NewIncrementstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrementstmtContext {
	var p = new(IncrementstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_incrementstmt

	return p
}

func (s *IncrementstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncrementstmtContext) AllVar_() []IVar_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_Context)(nil)).Elem())
	var tst = make([]IVar_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_Context)
		}
	}

	return tst
}

func (s *IncrementstmtContext) Var_(i int) IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *IncrementstmtContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *IncrementstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrementstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterIncrementstmt(s)
	}
}

func (s *IncrementstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitIncrementstmt(s)
	}
}

func (p *loopParser) Incrementstmt() (localctx IIncrementstmtContext) {
	this := p
	_ = this

	localctx = NewIncrementstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, loopParserRULE_incrementstmt)
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
		p.Var_()
	}
	{
		p.SetState(39)
		p.Match(loopParserT__1)
	}
	{
		p.SetState(40)
		p.Var_()
	}
	{
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(_la == loopParserT__2 || _la == loopParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(42)
		p.Number()
	}

	return localctx
}

// ILoopstmtContext is an interface to support dynamic dispatch.
type ILoopstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopstmtContext differentiates from other interfaces.
	IsLoopstmtContext()
}

type LoopstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopstmtContext() *LoopstmtContext {
	var p = new(LoopstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_loopstmt
	return p
}

func (*LoopstmtContext) IsLoopstmtContext() {}

func NewLoopstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopstmtContext {
	var p = new(LoopstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_loopstmt

	return p
}

func (s *LoopstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopstmtContext) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *LoopstmtContext) Statementlist() IStatementlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementlistContext)
}

func (s *LoopstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterLoopstmt(s)
	}
}

func (s *LoopstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitLoopstmt(s)
	}
}

func (p *loopParser) Loopstmt() (localctx ILoopstmtContext) {
	this := p
	_ = this

	localctx = NewLoopstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, loopParserRULE_loopstmt)

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
		p.SetState(44)
		p.Match(loopParserT__4)
	}
	{
		p.SetState(45)
		p.Var_()
	}
	{
		p.SetState(46)
		p.Match(loopParserT__5)
	}
	{
		p.SetState(47)
		p.Statementlist()
	}
	{
		p.SetState(48)
		p.Match(loopParserT__6)
	}

	return localctx
}

// IVar_Context is an interface to support dynamic dispatch.
type IVar_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_Context differentiates from other interfaces.
	IsVar_Context()
}

type Var_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_Context() *Var_Context {
	var p = new(Var_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_var_
	return p
}

func (*Var_Context) IsVar_Context() {}

func NewVar_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_Context {
	var p = new(Var_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_var_

	return p
}

func (s *Var_Context) GetParser() antlr.Parser { return s.parser }

func (s *Var_Context) ID() antlr.TerminalNode {
	return s.GetToken(loopParserID, 0)
}

func (s *Var_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterVar_(s)
	}
}

func (s *Var_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitVar_(s)
	}
}

func (p *loopParser) Var_() (localctx IVar_Context) {
	this := p
	_ = this

	localctx = NewVar_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, loopParserRULE_var_)

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
		p.Match(loopParserID)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = loopParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = loopParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(loopParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(loopListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *loopParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, loopParserRULE_number)

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
		p.Match(loopParserNUMBER)
	}

	return localctx
}
