// Code generated from janus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package janus // janus
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 153,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	5, 2, 37, 10, 2, 7, 2, 39, 10, 2, 12, 2, 14, 2, 42, 11, 2, 3, 2, 3, 2,
	3, 2, 7, 2, 47, 10, 2, 12, 2, 14, 2, 50, 11, 2, 3, 3, 6, 3, 53, 10, 3,
	13, 3, 14, 3, 54, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 69, 10, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 82, 10, 6, 3, 6, 3, 6, 5,
	6, 86, 10, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 95, 10, 7,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 109, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 119, 10, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 129, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 141, 10, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 149, 10, 15, 3, 16, 3, 16, 3, 16, 2, 2, 17,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 2, 2, 161, 2,
	40, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 64, 3, 2, 2, 2,
	10, 77, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 96, 3, 2, 2, 2, 16, 99, 3,
	2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 118, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2,
	24, 128, 3, 2, 2, 2, 26, 140, 3, 2, 2, 2, 28, 148, 3, 2, 2, 2, 30, 150,
	3, 2, 2, 2, 32, 36, 7, 28, 2, 2, 33, 34, 7, 3, 2, 2, 34, 35, 7, 29, 2,
	2, 35, 37, 7, 4, 2, 2, 36, 33, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39,
	3, 2, 2, 2, 38, 32, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 48, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 7,
	5, 2, 2, 44, 45, 7, 28, 2, 2, 45, 47, 5, 4, 3, 2, 46, 43, 3, 2, 2, 2, 47,
	50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 3, 3, 2, 2,
	2, 50, 48, 3, 2, 2, 2, 51, 53, 5, 6, 4, 2, 52, 51, 3, 2, 2, 2, 53, 54,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 5, 3, 2, 2, 2,
	56, 63, 5, 8, 5, 2, 57, 63, 5, 10, 6, 2, 58, 63, 5, 12, 7, 2, 59, 63, 5,
	14, 8, 2, 60, 63, 5, 16, 9, 2, 61, 63, 5, 18, 10, 2, 62, 56, 3, 2, 2, 2,
	62, 57, 3, 2, 2, 2, 62, 58, 3, 2, 2, 2, 62, 59, 3, 2, 2, 2, 62, 60, 3,
	2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 7, 3, 2, 2, 2, 64, 65, 7, 6, 2, 2, 65,
	68, 5, 24, 13, 2, 66, 67, 7, 7, 2, 2, 67, 69, 5, 4, 3, 2, 68, 66, 3, 2,
	2, 2, 68, 69, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 71, 7, 8, 2, 2, 71, 73,
	5, 4, 3, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 75, 7, 9, 2, 2, 75, 76, 5, 24, 13, 2, 76, 9, 3, 2, 2, 2, 77, 78, 7,
	10, 2, 2, 78, 81, 5, 24, 13, 2, 79, 80, 7, 11, 2, 2, 80, 82, 5, 4, 3, 2,
	81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 84, 7,
	12, 2, 2, 84, 86, 5, 4, 3, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 88, 7, 13, 2, 2, 88, 89, 5, 24, 13, 2, 89, 11, 3, 2,
	2, 2, 90, 91, 7, 14, 2, 2, 91, 95, 7, 28, 2, 2, 92, 93, 7, 15, 2, 2, 93,
	95, 7, 28, 2, 2, 94, 90, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 13, 3, 2,
	2, 2, 96, 97, 7, 16, 2, 2, 97, 98, 7, 28, 2, 2, 98, 15, 3, 2, 2, 2, 99,
	100, 7, 17, 2, 2, 100, 101, 7, 28, 2, 2, 101, 17, 3, 2, 2, 2, 102, 103,
	5, 28, 15, 2, 103, 104, 5, 20, 11, 2, 104, 109, 3, 2, 2, 2, 105, 106, 5,
	28, 15, 2, 106, 107, 5, 22, 12, 2, 107, 109, 3, 2, 2, 2, 108, 102, 3, 2,
	2, 2, 108, 105, 3, 2, 2, 2, 109, 19, 3, 2, 2, 2, 110, 111, 7, 18, 2, 2,
	111, 119, 5, 24, 13, 2, 112, 113, 7, 19, 2, 2, 113, 119, 5, 24, 13, 2,
	114, 115, 7, 20, 2, 2, 115, 119, 5, 24, 13, 2, 116, 117, 7, 21, 2, 2, 117,
	119, 5, 24, 13, 2, 118, 110, 3, 2, 2, 2, 118, 112, 3, 2, 2, 2, 118, 114,
	3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 21, 3, 2, 2, 2, 120, 121, 7, 22,
	2, 2, 121, 122, 5, 28, 15, 2, 122, 23, 3, 2, 2, 2, 123, 129, 5, 26, 14,
	2, 124, 125, 5, 26, 14, 2, 125, 126, 7, 27, 2, 2, 126, 127, 5, 24, 13,
	2, 127, 129, 3, 2, 2, 2, 128, 123, 3, 2, 2, 2, 128, 124, 3, 2, 2, 2, 129,
	25, 3, 2, 2, 2, 130, 131, 7, 23, 2, 2, 131, 132, 5, 24, 13, 2, 132, 133,
	7, 24, 2, 2, 133, 141, 3, 2, 2, 2, 134, 135, 7, 25, 2, 2, 135, 141, 5,
	24, 13, 2, 136, 137, 7, 26, 2, 2, 137, 141, 5, 24, 13, 2, 138, 141, 5,
	28, 15, 2, 139, 141, 5, 30, 16, 2, 140, 130, 3, 2, 2, 2, 140, 134, 3, 2,
	2, 2, 140, 136, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2,
	141, 27, 3, 2, 2, 2, 142, 149, 7, 28, 2, 2, 143, 144, 7, 28, 2, 2, 144,
	145, 7, 3, 2, 2, 145, 146, 5, 24, 13, 2, 146, 147, 7, 4, 2, 2, 147, 149,
	3, 2, 2, 2, 148, 142, 3, 2, 2, 2, 148, 143, 3, 2, 2, 2, 149, 29, 3, 2,
	2, 2, 150, 151, 7, 29, 2, 2, 151, 31, 3, 2, 2, 2, 17, 36, 40, 48, 54, 62,
	68, 72, 81, 85, 94, 108, 118, 128, 140, 148,
}
var literalNames = []string{
	"", "'['", "']'", "'PROCEDURE'", "'IF'", "'THEN'", "'ELSE'", "'FI'", "'FROM'",
	"'DO'", "'LOOP'", "'UNTIL'", "'CALL'", "'UNCALL'", "'READ'", "'WRITE'",
	"'+='", "'-='", "'!='", "'<=>'", "':'", "'('", "')'", "'-'", "'~'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "BINOP", "IDENT", "NUM", "COMMENT", "WS",
}

var ruleNames = []string{
	"program", "statements", "statement", "ifstmt", "dostmt", "callstmt", "readstmt",
	"writestmt", "lvalstmt", "modstmt", "swapstmt", "expression", "minexp",
	"lvalue", "constant",
}

type janusParser struct {
	*antlr.BaseParser
}

// NewjanusParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *janusParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewjanusParser(input antlr.TokenStream) *janusParser {
	this := new(janusParser)
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
	this.GrammarFileName = "janus.g4"

	return this
}

// janusParser tokens.
const (
	janusParserEOF     = antlr.TokenEOF
	janusParserT__0    = 1
	janusParserT__1    = 2
	janusParserT__2    = 3
	janusParserT__3    = 4
	janusParserT__4    = 5
	janusParserT__5    = 6
	janusParserT__6    = 7
	janusParserT__7    = 8
	janusParserT__8    = 9
	janusParserT__9    = 10
	janusParserT__10   = 11
	janusParserT__11   = 12
	janusParserT__12   = 13
	janusParserT__13   = 14
	janusParserT__14   = 15
	janusParserT__15   = 16
	janusParserT__16   = 17
	janusParserT__17   = 18
	janusParserT__18   = 19
	janusParserT__19   = 20
	janusParserT__20   = 21
	janusParserT__21   = 22
	janusParserT__22   = 23
	janusParserT__23   = 24
	janusParserBINOP   = 25
	janusParserIDENT   = 26
	janusParserNUM     = 27
	janusParserCOMMENT = 28
	janusParserWS      = 29
)

// janusParser rules.
const (
	janusParserRULE_program    = 0
	janusParserRULE_statements = 1
	janusParserRULE_statement  = 2
	janusParserRULE_ifstmt     = 3
	janusParserRULE_dostmt     = 4
	janusParserRULE_callstmt   = 5
	janusParserRULE_readstmt   = 6
	janusParserRULE_writestmt  = 7
	janusParserRULE_lvalstmt   = 8
	janusParserRULE_modstmt    = 9
	janusParserRULE_swapstmt   = 10
	janusParserRULE_expression = 11
	janusParserRULE_minexp     = 12
	janusParserRULE_lvalue     = 13
	janusParserRULE_constant   = 14
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
	p.RuleIndex = janusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(janusParserIDENT)
}

func (s *ProgramContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(janusParserIDENT, i)
}

func (s *ProgramContext) AllStatements() []IStatementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementsContext)(nil)).Elem())
	var tst = make([]IStatementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementsContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statements(i int) IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(janusParserNUM)
}

func (s *ProgramContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(janusParserNUM, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *janusParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, janusParserRULE_program)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == janusParserIDENT {
		{
			p.SetState(30)
			p.Match(janusParserIDENT)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == janusParserT__0 {
			{
				p.SetState(31)
				p.Match(janusParserT__0)
			}
			{
				p.SetState(32)
				p.Match(janusParserNUM)
			}
			{
				p.SetState(33)
				p.Match(janusParserT__1)
			}

		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == janusParserT__2 {
		{
			p.SetState(41)
			p.Match(janusParserT__2)
		}
		{
			p.SetState(42)
			p.Match(janusParserIDENT)
		}
		{
			p.SetState(43)
			p.Statements()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *janusParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, janusParserRULE_statements)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<janusParserT__3)|(1<<janusParserT__7)|(1<<janusParserT__11)|(1<<janusParserT__12)|(1<<janusParserT__13)|(1<<janusParserT__14)|(1<<janusParserIDENT))) != 0) {
		{
			p.SetState(49)
			p.Statement()
		}

		p.SetState(52)
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
	p.RuleIndex = janusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Ifstmt() IIfstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StatementContext) Dostmt() IDostmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDostmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDostmtContext)
}

func (s *StatementContext) Callstmt() ICallstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallstmtContext)
}

func (s *StatementContext) Readstmt() IReadstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadstmtContext)
}

func (s *StatementContext) Writestmt() IWritestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWritestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWritestmtContext)
}

func (s *StatementContext) Lvalstmt() ILvalstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalstmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *janusParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, janusParserRULE_statement)

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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case janusParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Ifstmt()
		}

	case janusParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Dostmt()
		}

	case janusParserT__11, janusParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Callstmt()
		}

	case janusParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Readstmt()
		}

	case janusParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(58)
			p.Writestmt()
		}

	case janusParserIDENT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(59)
			p.Lvalstmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_ifstmt
	return p
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IfstmtContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfstmtContext) AllStatements() []IStatementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementsContext)(nil)).Elem())
	var tst = make([]IStatementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementsContext)
		}
	}

	return tst
}

func (s *IfstmtContext) Statements(i int) IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *janusParser) Ifstmt() (localctx IIfstmtContext) {
	this := p
	_ = this

	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, janusParserRULE_ifstmt)
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
		p.SetState(62)
		p.Match(janusParserT__3)
	}
	{
		p.SetState(63)
		p.Expression()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == janusParserT__4 {
		{
			p.SetState(64)
			p.Match(janusParserT__4)
		}
		{
			p.SetState(65)
			p.Statements()
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == janusParserT__5 {
		{
			p.SetState(68)
			p.Match(janusParserT__5)
		}
		{
			p.SetState(69)
			p.Statements()
		}

	}
	{
		p.SetState(72)
		p.Match(janusParserT__6)
	}
	{
		p.SetState(73)
		p.Expression()
	}

	return localctx
}

// IDostmtContext is an interface to support dynamic dispatch.
type IDostmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDostmtContext differentiates from other interfaces.
	IsDostmtContext()
}

type DostmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDostmtContext() *DostmtContext {
	var p = new(DostmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_dostmt
	return p
}

func (*DostmtContext) IsDostmtContext() {}

func NewDostmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DostmtContext {
	var p = new(DostmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_dostmt

	return p
}

func (s *DostmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DostmtContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DostmtContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DostmtContext) AllStatements() []IStatementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementsContext)(nil)).Elem())
	var tst = make([]IStatementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementsContext)
		}
	}

	return tst
}

func (s *DostmtContext) Statements(i int) IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *DostmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DostmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DostmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterDostmt(s)
	}
}

func (s *DostmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitDostmt(s)
	}
}

func (p *janusParser) Dostmt() (localctx IDostmtContext) {
	this := p
	_ = this

	localctx = NewDostmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, janusParserRULE_dostmt)
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
		p.SetState(75)
		p.Match(janusParserT__7)
	}
	{
		p.SetState(76)
		p.Expression()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == janusParserT__8 {
		{
			p.SetState(77)
			p.Match(janusParserT__8)
		}
		{
			p.SetState(78)
			p.Statements()
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == janusParserT__9 {
		{
			p.SetState(81)
			p.Match(janusParserT__9)
		}
		{
			p.SetState(82)
			p.Statements()
		}

	}
	{
		p.SetState(85)
		p.Match(janusParserT__10)
	}
	{
		p.SetState(86)
		p.Expression()
	}

	return localctx
}

// ICallstmtContext is an interface to support dynamic dispatch.
type ICallstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallstmtContext differentiates from other interfaces.
	IsCallstmtContext()
}

type CallstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallstmtContext() *CallstmtContext {
	var p = new(CallstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_callstmt
	return p
}

func (*CallstmtContext) IsCallstmtContext() {}

func NewCallstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallstmtContext {
	var p = new(CallstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_callstmt

	return p
}

func (s *CallstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CallstmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(janusParserIDENT, 0)
}

func (s *CallstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterCallstmt(s)
	}
}

func (s *CallstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitCallstmt(s)
	}
}

func (p *janusParser) Callstmt() (localctx ICallstmtContext) {
	this := p
	_ = this

	localctx = NewCallstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, janusParserRULE_callstmt)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case janusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(janusParserT__11)
		}
		{
			p.SetState(89)
			p.Match(janusParserIDENT)
		}

	case janusParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(janusParserT__12)
		}
		{
			p.SetState(91)
			p.Match(janusParserIDENT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReadstmtContext is an interface to support dynamic dispatch.
type IReadstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadstmtContext differentiates from other interfaces.
	IsReadstmtContext()
}

type ReadstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadstmtContext() *ReadstmtContext {
	var p = new(ReadstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_readstmt
	return p
}

func (*ReadstmtContext) IsReadstmtContext() {}

func NewReadstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadstmtContext {
	var p = new(ReadstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_readstmt

	return p
}

func (s *ReadstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadstmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(janusParserIDENT, 0)
}

func (s *ReadstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterReadstmt(s)
	}
}

func (s *ReadstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitReadstmt(s)
	}
}

func (p *janusParser) Readstmt() (localctx IReadstmtContext) {
	this := p
	_ = this

	localctx = NewReadstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, janusParserRULE_readstmt)

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
		p.SetState(94)
		p.Match(janusParserT__13)
	}
	{
		p.SetState(95)
		p.Match(janusParserIDENT)
	}

	return localctx
}

// IWritestmtContext is an interface to support dynamic dispatch.
type IWritestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWritestmtContext differentiates from other interfaces.
	IsWritestmtContext()
}

type WritestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWritestmtContext() *WritestmtContext {
	var p = new(WritestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_writestmt
	return p
}

func (*WritestmtContext) IsWritestmtContext() {}

func NewWritestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WritestmtContext {
	var p = new(WritestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_writestmt

	return p
}

func (s *WritestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WritestmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(janusParserIDENT, 0)
}

func (s *WritestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WritestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WritestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterWritestmt(s)
	}
}

func (s *WritestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitWritestmt(s)
	}
}

func (p *janusParser) Writestmt() (localctx IWritestmtContext) {
	this := p
	_ = this

	localctx = NewWritestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, janusParserRULE_writestmt)

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
		p.Match(janusParserT__14)
	}
	{
		p.SetState(98)
		p.Match(janusParserIDENT)
	}

	return localctx
}

// ILvalstmtContext is an interface to support dynamic dispatch.
type ILvalstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalstmtContext differentiates from other interfaces.
	IsLvalstmtContext()
}

type LvalstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalstmtContext() *LvalstmtContext {
	var p = new(LvalstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_lvalstmt
	return p
}

func (*LvalstmtContext) IsLvalstmtContext() {}

func NewLvalstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalstmtContext {
	var p = new(LvalstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_lvalstmt

	return p
}

func (s *LvalstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalstmtContext) Lvalue() ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LvalstmtContext) Modstmt() IModstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModstmtContext)
}

func (s *LvalstmtContext) Swapstmt() ISwapstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwapstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwapstmtContext)
}

func (s *LvalstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterLvalstmt(s)
	}
}

func (s *LvalstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitLvalstmt(s)
	}
}

func (p *janusParser) Lvalstmt() (localctx ILvalstmtContext) {
	this := p
	_ = this

	localctx = NewLvalstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, janusParserRULE_lvalstmt)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Lvalue()
		}
		{
			p.SetState(101)
			p.Modstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Lvalue()
		}
		{
			p.SetState(104)
			p.Swapstmt()
		}

	}

	return localctx
}

// IModstmtContext is an interface to support dynamic dispatch.
type IModstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModstmtContext differentiates from other interfaces.
	IsModstmtContext()
}

type ModstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModstmtContext() *ModstmtContext {
	var p = new(ModstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_modstmt
	return p
}

func (*ModstmtContext) IsModstmtContext() {}

func NewModstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModstmtContext {
	var p = new(ModstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_modstmt

	return p
}

func (s *ModstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ModstmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterModstmt(s)
	}
}

func (s *ModstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitModstmt(s)
	}
}

func (p *janusParser) Modstmt() (localctx IModstmtContext) {
	this := p
	_ = this

	localctx = NewModstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, janusParserRULE_modstmt)

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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case janusParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(janusParserT__15)
		}
		{
			p.SetState(109)
			p.Expression()
		}

	case janusParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(janusParserT__16)
		}
		{
			p.SetState(111)
			p.Expression()
		}

	case janusParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(janusParserT__17)
		}
		{
			p.SetState(113)
			p.Expression()
		}

	case janusParserT__18:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Match(janusParserT__18)
		}
		{
			p.SetState(115)
			p.Expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISwapstmtContext is an interface to support dynamic dispatch.
type ISwapstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwapstmtContext differentiates from other interfaces.
	IsSwapstmtContext()
}

type SwapstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwapstmtContext() *SwapstmtContext {
	var p = new(SwapstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_swapstmt
	return p
}

func (*SwapstmtContext) IsSwapstmtContext() {}

func NewSwapstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwapstmtContext {
	var p = new(SwapstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_swapstmt

	return p
}

func (s *SwapstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwapstmtContext) Lvalue() ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SwapstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwapstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwapstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterSwapstmt(s)
	}
}

func (s *SwapstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitSwapstmt(s)
	}
}

func (p *janusParser) Swapstmt() (localctx ISwapstmtContext) {
	this := p
	_ = this

	localctx = NewSwapstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, janusParserRULE_swapstmt)

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
		p.SetState(118)
		p.Match(janusParserT__19)
	}
	{
		p.SetState(119)
		p.Lvalue()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Minexp() IMinexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinexpContext)
}

func (s *ExpressionContext) BINOP() antlr.TerminalNode {
	return s.GetToken(janusParserBINOP, 0)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *janusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, janusParserRULE_expression)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Minexp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Minexp()
		}
		{
			p.SetState(123)
			p.Match(janusParserBINOP)
		}
		{
			p.SetState(124)
			p.Expression()
		}

	}

	return localctx
}

// IMinexpContext is an interface to support dynamic dispatch.
type IMinexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinexpContext differentiates from other interfaces.
	IsMinexpContext()
}

type MinexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinexpContext() *MinexpContext {
	var p = new(MinexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_minexp
	return p
}

func (*MinexpContext) IsMinexpContext() {}

func NewMinexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinexpContext {
	var p = new(MinexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_minexp

	return p
}

func (s *MinexpContext) GetParser() antlr.Parser { return s.parser }

func (s *MinexpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MinexpContext) Lvalue() ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *MinexpContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *MinexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterMinexp(s)
	}
}

func (s *MinexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitMinexp(s)
	}
}

func (p *janusParser) Minexp() (localctx IMinexpContext) {
	this := p
	_ = this

	localctx = NewMinexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, janusParserRULE_minexp)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case janusParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(janusParserT__20)
		}
		{
			p.SetState(129)
			p.Expression()
		}
		{
			p.SetState(130)
			p.Match(janusParserT__21)
		}

	case janusParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(janusParserT__22)
		}
		{
			p.SetState(133)
			p.Expression()
		}

	case janusParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(janusParserT__23)
		}
		{
			p.SetState(135)
			p.Expression()
		}

	case janusParserIDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Lvalue()
		}

	case janusParserNUM:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
			p.Constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) IDENT() antlr.TerminalNode {
	return s.GetToken(janusParserIDENT, 0)
}

func (s *LvalueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *janusParser) Lvalue() (localctx ILvalueContext) {
	this := p
	_ = this

	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, janusParserRULE_lvalue)

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(janusParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(janusParserIDENT)
		}
		{
			p.SetState(142)
			p.Match(janusParserT__0)
		}
		{
			p.SetState(143)
			p.Expression()
		}
		{
			p.SetState(144)
			p.Match(janusParserT__1)
		}

	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = janusParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = janusParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) NUM() antlr.TerminalNode {
	return s.GetToken(janusParserNUM, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(janusListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *janusParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, janusParserRULE_constant)

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
		p.SetState(148)
		p.Match(janusParserNUM)
	}

	return localctx
}
