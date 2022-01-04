// Code generated from microc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package microc // microc
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 6, 2, 30, 10, 2, 13, 2, 14, 2, 31, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 38, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 45,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 53, 10, 6, 12, 6, 14,
	6, 56, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 72, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 79, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 7, 11, 90, 10, 11, 12, 11, 14, 11, 93, 11, 11, 3, 12,
	3, 12, 3, 12, 5, 12, 98, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2,
	3, 20, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 2, 2, 102,
	2, 29, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 46, 3, 2, 2,
	2, 10, 50, 3, 2, 2, 2, 12, 59, 3, 2, 2, 2, 14, 62, 3, 2, 2, 2, 16, 71,
	3, 2, 2, 2, 18, 78, 3, 2, 2, 2, 20, 80, 3, 2, 2, 2, 22, 97, 3, 2, 2, 2,
	24, 99, 3, 2, 2, 2, 26, 101, 3, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 28, 3,
	2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	3, 3, 2, 2, 2, 33, 38, 5, 6, 4, 2, 34, 38, 5, 8, 5, 2, 35, 38, 5, 10, 6,
	2, 36, 38, 5, 12, 7, 2, 37, 33, 3, 2, 2, 2, 37, 34, 3, 2, 2, 2, 37, 35,
	3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 5, 3, 2, 2, 2, 39, 40, 7, 3, 2, 2,
	40, 41, 5, 14, 8, 2, 41, 44, 5, 4, 3, 2, 42, 43, 7, 4, 2, 2, 43, 45, 5,
	4, 3, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 7, 3, 2, 2, 2, 46,
	47, 7, 5, 2, 2, 47, 48, 5, 14, 8, 2, 48, 49, 5, 4, 3, 2, 49, 9, 3, 2, 2,
	2, 50, 54, 7, 6, 2, 2, 51, 53, 5, 4, 3, 2, 52, 51, 3, 2, 2, 2, 53, 56,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2,
	56, 54, 3, 2, 2, 2, 57, 58, 7, 7, 2, 2, 58, 11, 3, 2, 2, 2, 59, 60, 5,
	16, 9, 2, 60, 61, 7, 8, 2, 2, 61, 13, 3, 2, 2, 2, 62, 63, 7, 9, 2, 2, 63,
	64, 5, 16, 9, 2, 64, 65, 7, 10, 2, 2, 65, 15, 3, 2, 2, 2, 66, 72, 5, 18,
	10, 2, 67, 68, 5, 24, 13, 2, 68, 69, 7, 11, 2, 2, 69, 70, 5, 16, 9, 2,
	70, 72, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 71, 67, 3, 2, 2, 2, 72, 17, 3,
	2, 2, 2, 73, 79, 5, 20, 11, 2, 74, 75, 5, 20, 11, 2, 75, 76, 7, 12, 2,
	2, 76, 77, 5, 20, 11, 2, 77, 79, 3, 2, 2, 2, 78, 73, 3, 2, 2, 2, 78, 74,
	3, 2, 2, 2, 79, 19, 3, 2, 2, 2, 80, 81, 8, 11, 1, 2, 81, 82, 5, 22, 12,
	2, 82, 91, 3, 2, 2, 2, 83, 84, 12, 4, 2, 2, 84, 85, 7, 13, 2, 2, 85, 90,
	5, 22, 12, 2, 86, 87, 12, 3, 2, 2, 87, 88, 7, 14, 2, 2, 88, 90, 5, 22,
	12, 2, 89, 83, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91,
	89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 21, 3, 2, 2, 2, 93, 91, 3, 2, 2,
	2, 94, 98, 5, 24, 13, 2, 95, 98, 5, 26, 14, 2, 96, 98, 5, 14, 8, 2, 97,
	94, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 23, 3, 2, 2,
	2, 99, 100, 7, 15, 2, 2, 100, 25, 3, 2, 2, 2, 101, 102, 7, 16, 2, 2, 102,
	27, 3, 2, 2, 2, 11, 31, 37, 44, 54, 71, 78, 89, 91, 97,
}
var literalNames = []string{
	"", "'if'", "'else'", "'while'", "'{'", "'}'", "';'", "'('", "')'", "'='",
	"'<'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT", "WS",
}

var ruleNames = []string{
	"program", "statement", "ifstatement", "whilestatement", "blockstatement",
	"exprstatement", "paren_expr", "expr", "test", "sum_", "term", "id_", "integer",
}

type microcParser struct {
	*antlr.BaseParser
}

// NewmicrocParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *microcParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmicrocParser(input antlr.TokenStream) *microcParser {
	this := new(microcParser)
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
	this.GrammarFileName = "microc.g4"

	return this
}

// microcParser tokens.
const (
	microcParserEOF    = antlr.TokenEOF
	microcParserT__0   = 1
	microcParserT__1   = 2
	microcParserT__2   = 3
	microcParserT__3   = 4
	microcParserT__4   = 5
	microcParserT__5   = 6
	microcParserT__6   = 7
	microcParserT__7   = 8
	microcParserT__8   = 9
	microcParserT__9   = 10
	microcParserT__10  = 11
	microcParserT__11  = 12
	microcParserSTRING = 13
	microcParserINT    = 14
	microcParserWS     = 15
)

// microcParser rules.
const (
	microcParserRULE_program        = 0
	microcParserRULE_statement      = 1
	microcParserRULE_ifstatement    = 2
	microcParserRULE_whilestatement = 3
	microcParserRULE_blockstatement = 4
	microcParserRULE_exprstatement  = 5
	microcParserRULE_paren_expr     = 6
	microcParserRULE_expr           = 7
	microcParserRULE_test           = 8
	microcParserRULE_sum_           = 9
	microcParserRULE_term           = 10
	microcParserRULE_id_            = 11
	microcParserRULE_integer        = 12
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
	p.RuleIndex = microcParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *microcParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, microcParserRULE_program)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<microcParserT__0)|(1<<microcParserT__2)|(1<<microcParserT__3)|(1<<microcParserT__6)|(1<<microcParserSTRING)|(1<<microcParserINT))) != 0) {
		{
			p.SetState(26)
			p.Statement()
		}

		p.SetState(29)
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
	p.RuleIndex = microcParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Ifstatement() IIfstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatementContext)
}

func (s *StatementContext) Whilestatement() IWhilestatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhilestatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhilestatementContext)
}

func (s *StatementContext) Blockstatement() IBlockstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockstatementContext)
}

func (s *StatementContext) Exprstatement() IExprstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprstatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *microcParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, microcParserRULE_statement)

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

	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case microcParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Ifstatement()
		}

	case microcParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Whilestatement()
		}

	case microcParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Blockstatement()
		}

	case microcParserT__6, microcParserSTRING, microcParserINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(34)
			p.Exprstatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfstatementContext is an interface to support dynamic dispatch.
type IIfstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatementContext differentiates from other interfaces.
	IsIfstatementContext()
}

type IfstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatementContext() *IfstatementContext {
	var p = new(IfstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_ifstatement
	return p
}

func (*IfstatementContext) IsIfstatementContext() {}

func NewIfstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatementContext {
	var p = new(IfstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_ifstatement

	return p
}

func (s *IfstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatementContext) Paren_expr() IParen_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *IfstatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfstatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterIfstatement(s)
	}
}

func (s *IfstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitIfstatement(s)
	}
}

func (p *microcParser) Ifstatement() (localctx IIfstatementContext) {
	this := p
	_ = this

	localctx = NewIfstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, microcParserRULE_ifstatement)

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
		p.Match(microcParserT__0)
	}
	{
		p.SetState(38)
		p.Paren_expr()
	}
	{
		p.SetState(39)
		p.Statement()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(40)
			p.Match(microcParserT__1)
		}
		{
			p.SetState(41)
			p.Statement()
		}

	}

	return localctx
}

// IWhilestatementContext is an interface to support dynamic dispatch.
type IWhilestatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhilestatementContext differentiates from other interfaces.
	IsWhilestatementContext()
}

type WhilestatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestatementContext() *WhilestatementContext {
	var p = new(WhilestatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_whilestatement
	return p
}

func (*WhilestatementContext) IsWhilestatementContext() {}

func NewWhilestatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestatementContext {
	var p = new(WhilestatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_whilestatement

	return p
}

func (s *WhilestatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestatementContext) Paren_expr() IParen_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *WhilestatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhilestatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterWhilestatement(s)
	}
}

func (s *WhilestatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitWhilestatement(s)
	}
}

func (p *microcParser) Whilestatement() (localctx IWhilestatementContext) {
	this := p
	_ = this

	localctx = NewWhilestatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, microcParserRULE_whilestatement)

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
		p.Match(microcParserT__2)
	}
	{
		p.SetState(45)
		p.Paren_expr()
	}
	{
		p.SetState(46)
		p.Statement()
	}

	return localctx
}

// IBlockstatementContext is an interface to support dynamic dispatch.
type IBlockstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockstatementContext differentiates from other interfaces.
	IsBlockstatementContext()
}

type BlockstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockstatementContext() *BlockstatementContext {
	var p = new(BlockstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_blockstatement
	return p
}

func (*BlockstatementContext) IsBlockstatementContext() {}

func NewBlockstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockstatementContext {
	var p = new(BlockstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_blockstatement

	return p
}

func (s *BlockstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockstatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockstatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterBlockstatement(s)
	}
}

func (s *BlockstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitBlockstatement(s)
	}
}

func (p *microcParser) Blockstatement() (localctx IBlockstatementContext) {
	this := p
	_ = this

	localctx = NewBlockstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, microcParserRULE_blockstatement)
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
		p.SetState(48)
		p.Match(microcParserT__3)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<microcParserT__0)|(1<<microcParserT__2)|(1<<microcParserT__3)|(1<<microcParserT__6)|(1<<microcParserSTRING)|(1<<microcParserINT))) != 0 {
		{
			p.SetState(49)
			p.Statement()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(microcParserT__4)
	}

	return localctx
}

// IExprstatementContext is an interface to support dynamic dispatch.
type IExprstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprstatementContext differentiates from other interfaces.
	IsExprstatementContext()
}

type ExprstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprstatementContext() *ExprstatementContext {
	var p = new(ExprstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_exprstatement
	return p
}

func (*ExprstatementContext) IsExprstatementContext() {}

func NewExprstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprstatementContext {
	var p = new(ExprstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_exprstatement

	return p
}

func (s *ExprstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprstatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterExprstatement(s)
	}
}

func (s *ExprstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitExprstatement(s)
	}
}

func (p *microcParser) Exprstatement() (localctx IExprstatementContext) {
	this := p
	_ = this

	localctx = NewExprstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, microcParserRULE_exprstatement)

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
		p.SetState(57)
		p.Expr()
	}
	{
		p.SetState(58)
		p.Match(microcParserT__5)
	}

	return localctx
}

// IParen_exprContext is an interface to support dynamic dispatch.
type IParen_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParen_exprContext differentiates from other interfaces.
	IsParen_exprContext()
}

type Paren_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParen_exprContext() *Paren_exprContext {
	var p = new(Paren_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_paren_expr
	return p
}

func (*Paren_exprContext) IsParen_exprContext() {}

func NewParen_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Paren_exprContext {
	var p = new(Paren_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_paren_expr

	return p
}

func (s *Paren_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Paren_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Paren_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Paren_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Paren_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterParen_expr(s)
	}
}

func (s *Paren_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitParen_expr(s)
	}
}

func (p *microcParser) Paren_expr() (localctx IParen_exprContext) {
	this := p
	_ = this

	localctx = NewParen_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, microcParserRULE_paren_expr)

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
		p.SetState(60)
		p.Match(microcParserT__6)
	}
	{
		p.SetState(61)
		p.Expr()
	}
	{
		p.SetState(62)
		p.Match(microcParserT__7)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExprContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *microcParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, microcParserRULE_expr)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Test()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Id_()
		}
		{
			p.SetState(66)
			p.Match(microcParserT__8)
		}
		{
			p.SetState(67)
			p.Expr()
		}

	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllSum_() []ISum_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISum_Context)(nil)).Elem())
	var tst = make([]ISum_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISum_Context)
		}
	}

	return tst
}

func (s *TestContext) Sum_(i int) ISum_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISum_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISum_Context)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitTest(s)
	}
}

func (p *microcParser) Test() (localctx ITestContext) {
	this := p
	_ = this

	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, microcParserRULE_test)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.sum_(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.sum_(0)
		}
		{
			p.SetState(73)
			p.Match(microcParserT__9)
		}
		{
			p.SetState(74)
			p.sum_(0)
		}

	}

	return localctx
}

// ISum_Context is an interface to support dynamic dispatch.
type ISum_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSum_Context differentiates from other interfaces.
	IsSum_Context()
}

type Sum_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySum_Context() *Sum_Context {
	var p = new(Sum_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_sum_
	return p
}

func (*Sum_Context) IsSum_Context() {}

func NewSum_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sum_Context {
	var p = new(Sum_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_sum_

	return p
}

func (s *Sum_Context) GetParser() antlr.Parser { return s.parser }

func (s *Sum_Context) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Sum_Context) Sum_() ISum_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISum_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISum_Context)
}

func (s *Sum_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sum_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sum_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterSum_(s)
	}
}

func (s *Sum_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitSum_(s)
	}
}

func (p *microcParser) Sum_() (localctx ISum_Context) {
	return p.sum_(0)
}

func (p *microcParser) sum_(_p int) (localctx ISum_Context) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSum_Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISum_Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, microcParserRULE_sum_, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	{
		p.SetState(79)
		p.Term()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, microcParserRULE_sum_)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(82)
					p.Match(microcParserT__10)
				}
				{
					p.SetState(83)
					p.Term()
				}

			case 2:
				localctx = NewSum_Context(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, microcParserRULE_sum_)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(85)
					p.Match(microcParserT__11)
				}
				{
					p.SetState(86)
					p.Term()
				}

			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *TermContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Paren_expr() IParen_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParen_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParen_exprContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *microcParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, microcParserRULE_term)

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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case microcParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Id_()
		}

	case microcParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Integer()
		}

	case microcParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Paren_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) STRING() antlr.TerminalNode {
	return s.GetToken(microcParserSTRING, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *microcParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, microcParserRULE_id_)

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
		p.SetState(97)
		p.Match(microcParserSTRING)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = microcParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = microcParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(microcParserINT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(microcListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *microcParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, microcParserRULE_integer)

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
		p.SetState(99)
		p.Match(microcParserINT)
	}

	return localctx
}

func (p *microcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *Sum_Context = nil
		if localctx != nil {
			t = localctx.(*Sum_Context)
		}
		return p.Sum__Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *microcParser) Sum__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
