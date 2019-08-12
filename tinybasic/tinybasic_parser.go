// Code generated from tinybasic.g4 by ANTLR 4.7.2. DO NOT EDIT.

package tinybasic // tinybasic
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 136,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 38, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 47, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 69, 10, 4, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 78,
	10, 5, 7, 5, 80, 10, 5, 12, 5, 14, 5, 83, 11, 5, 3, 6, 3, 6, 3, 6, 7, 6,
	88, 10, 6, 12, 6, 14, 6, 91, 11, 6, 3, 7, 5, 7, 94, 10, 7, 3, 7, 3, 7,
	3, 7, 7, 7, 99, 10, 7, 12, 7, 14, 7, 102, 11, 7, 3, 8, 3, 8, 3, 8, 7, 8,
	107, 10, 8, 12, 8, 14, 8, 110, 11, 8, 3, 9, 3, 9, 5, 9, 114, 10, 9, 3,
	10, 3, 10, 3, 11, 6, 11, 119, 10, 11, 13, 11, 14, 11, 120, 3, 12, 3, 12,
	5, 12, 125, 10, 12, 3, 12, 3, 12, 5, 12, 129, 10, 12, 3, 12, 3, 12, 3,
	12, 5, 12, 134, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 2, 8, 3, 2, 17, 19, 3, 2, 17, 18, 3, 2, 20, 21, 4, 2, 24, 24, 26,
	26, 5, 2, 9, 9, 19, 19, 23, 23, 5, 2, 9, 9, 19, 19, 22, 22, 2, 153, 2,
	27, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 72, 3, 2, 2, 2,
	10, 84, 3, 2, 2, 2, 12, 93, 3, 2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 113, 3,
	2, 2, 2, 18, 115, 3, 2, 2, 2, 20, 118, 3, 2, 2, 2, 22, 133, 3, 2, 2, 2,
	24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 3, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30,
	31, 5, 20, 11, 2, 31, 32, 5, 6, 4, 2, 32, 33, 7, 27, 2, 2, 33, 38, 3, 2,
	2, 2, 34, 35, 5, 6, 4, 2, 35, 36, 7, 27, 2, 2, 36, 38, 3, 2, 2, 2, 37,
	30, 3, 2, 2, 2, 37, 34, 3, 2, 2, 2, 38, 5, 3, 2, 2, 2, 39, 40, 7, 3, 2,
	2, 40, 69, 5, 8, 5, 2, 41, 42, 7, 4, 2, 2, 42, 43, 5, 12, 7, 2, 43, 44,
	5, 22, 12, 2, 44, 46, 5, 12, 7, 2, 45, 47, 7, 5, 2, 2, 46, 45, 3, 2, 2,
	2, 46, 47, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 5, 6, 4, 2, 49, 69,
	3, 2, 2, 2, 50, 51, 7, 6, 2, 2, 51, 69, 5, 20, 11, 2, 52, 53, 7, 7, 2,
	2, 53, 69, 5, 10, 6, 2, 54, 56, 7, 8, 2, 2, 55, 54, 3, 2, 2, 2, 55, 56,
	3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 5, 18, 10, 2, 58, 59, 7, 9, 2,
	2, 59, 60, 5, 12, 7, 2, 60, 69, 3, 2, 2, 2, 61, 62, 7, 10, 2, 2, 62, 69,
	5, 12, 7, 2, 63, 69, 7, 11, 2, 2, 64, 69, 7, 12, 2, 2, 65, 69, 7, 13, 2,
	2, 66, 69, 7, 14, 2, 2, 67, 69, 7, 15, 2, 2, 68, 39, 3, 2, 2, 2, 68, 41,
	3, 2, 2, 2, 68, 50, 3, 2, 2, 2, 68, 52, 3, 2, 2, 2, 68, 55, 3, 2, 2, 2,
	68, 61, 3, 2, 2, 2, 68, 63, 3, 2, 2, 2, 68, 64, 3, 2, 2, 2, 68, 65, 3,
	2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 7, 3, 2, 2, 2, 70,
	73, 7, 24, 2, 2, 71, 73, 5, 12, 7, 2, 72, 70, 3, 2, 2, 2, 72, 71, 3, 2,
	2, 2, 73, 81, 3, 2, 2, 2, 74, 77, 7, 16, 2, 2, 75, 78, 7, 24, 2, 2, 76,
	78, 5, 12, 7, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 80, 3, 2,
	2, 2, 79, 74, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 9, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 89, 5, 18, 10, 2,
	85, 86, 7, 16, 2, 2, 86, 88, 5, 18, 10, 2, 87, 85, 3, 2, 2, 2, 88, 91,
	3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 11, 3, 2, 2, 2,
	91, 89, 3, 2, 2, 2, 92, 94, 9, 2, 2, 2, 93, 92, 3, 2, 2, 2, 93, 94, 3,
	2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 100, 5, 14, 8, 2, 96, 97, 9, 3, 2, 2,
	97, 99, 5, 14, 8, 2, 98, 96, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98,
	3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 13, 3, 2, 2, 2, 102, 100, 3, 2,
	2, 2, 103, 108, 5, 16, 9, 2, 104, 105, 9, 4, 2, 2, 105, 107, 5, 16, 9,
	2, 106, 104, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 15, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 114, 5,
	18, 10, 2, 112, 114, 5, 20, 11, 2, 113, 111, 3, 2, 2, 2, 113, 112, 3, 2,
	2, 2, 114, 17, 3, 2, 2, 2, 115, 116, 9, 5, 2, 2, 116, 19, 3, 2, 2, 2, 117,
	119, 7, 25, 2, 2, 118, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118,
	3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 21, 3, 2, 2, 2, 122, 124, 7, 22,
	2, 2, 123, 125, 9, 6, 2, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2,
	125, 134, 3, 2, 2, 2, 126, 128, 7, 23, 2, 2, 127, 129, 9, 7, 2, 2, 128,
	127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 134, 3, 2, 2, 2, 130, 134,
	7, 9, 2, 2, 131, 134, 7, 17, 2, 2, 132, 134, 7, 18, 2, 2, 133, 122, 3,
	2, 2, 2, 133, 126, 3, 2, 2, 2, 133, 130, 3, 2, 2, 2, 133, 131, 3, 2, 2,
	2, 133, 132, 3, 2, 2, 2, 134, 23, 3, 2, 2, 2, 19, 27, 37, 46, 55, 68, 72,
	77, 81, 89, 93, 100, 108, 113, 120, 124, 128, 133,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'PRINT'", "'IF'", "'THEN'", "'GOTO'", "'INPUT'", "'LET'", "'='", "'GOSUB'",
	"'RETURN'", "'CLEAR'", "'LIST'", "'RUN'", "'END'", "','", "'+'", "'-'",
	"'\u03B5'", "'*'", "'/'", "'<'", "'>'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "STRING", "DIGIT", "VAR", "CR", "WS",
}

var ruleNames = []string{
	"program", "line", "statement", "exprlist", "varlist", "expression", "term",
	"factor", "vara", "number", "relop",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tinybasicParser struct {
	*antlr.BaseParser
}

func NewtinybasicParser(input antlr.TokenStream) *tinybasicParser {
	this := new(tinybasicParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tinybasic.g4"

	return this
}

// tinybasicParser tokens.
const (
	tinybasicParserEOF    = antlr.TokenEOF
	tinybasicParserT__0   = 1
	tinybasicParserT__1   = 2
	tinybasicParserT__2   = 3
	tinybasicParserT__3   = 4
	tinybasicParserT__4   = 5
	tinybasicParserT__5   = 6
	tinybasicParserT__6   = 7
	tinybasicParserT__7   = 8
	tinybasicParserT__8   = 9
	tinybasicParserT__9   = 10
	tinybasicParserT__10  = 11
	tinybasicParserT__11  = 12
	tinybasicParserT__12  = 13
	tinybasicParserT__13  = 14
	tinybasicParserT__14  = 15
	tinybasicParserT__15  = 16
	tinybasicParserT__16  = 17
	tinybasicParserT__17  = 18
	tinybasicParserT__18  = 19
	tinybasicParserT__19  = 20
	tinybasicParserT__20  = 21
	tinybasicParserSTRING = 22
	tinybasicParserDIGIT  = 23
	tinybasicParserVAR    = 24
	tinybasicParserCR     = 25
	tinybasicParserWS     = 26
)

// tinybasicParser rules.
const (
	tinybasicParserRULE_program    = 0
	tinybasicParserRULE_line       = 1
	tinybasicParserRULE_statement  = 2
	tinybasicParserRULE_exprlist   = 3
	tinybasicParserRULE_varlist    = 4
	tinybasicParserRULE_expression = 5
	tinybasicParserRULE_term       = 6
	tinybasicParserRULE_factor     = 7
	tinybasicParserRULE_vara       = 8
	tinybasicParserRULE_number     = 9
	tinybasicParserRULE_relop      = 10
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
	p.RuleIndex = tinybasicParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgramContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *tinybasicParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tinybasicParserRULE_program)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinybasicParserT__0)|(1<<tinybasicParserT__1)|(1<<tinybasicParserT__3)|(1<<tinybasicParserT__4)|(1<<tinybasicParserT__5)|(1<<tinybasicParserT__7)|(1<<tinybasicParserT__8)|(1<<tinybasicParserT__9)|(1<<tinybasicParserT__10)|(1<<tinybasicParserT__11)|(1<<tinybasicParserT__12)|(1<<tinybasicParserSTRING)|(1<<tinybasicParserDIGIT)|(1<<tinybasicParserVAR))) != 0 {
		{
			p.SetState(22)
			p.Line()
		}

		p.SetState(27)
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
	p.RuleIndex = tinybasicParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LineContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LineContext) CR() antlr.TerminalNode {
	return s.GetToken(tinybasicParserCR, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *tinybasicParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tinybasicParserRULE_line)

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
	case tinybasicParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Number()
		}
		{
			p.SetState(29)
			p.Statement()
		}
		{
			p.SetState(30)
			p.Match(tinybasicParserCR)
		}

	case tinybasicParserT__0, tinybasicParserT__1, tinybasicParserT__3, tinybasicParserT__4, tinybasicParserT__5, tinybasicParserT__7, tinybasicParserT__8, tinybasicParserT__9, tinybasicParserT__10, tinybasicParserT__11, tinybasicParserT__12, tinybasicParserSTRING, tinybasicParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Statement()
		}
		{
			p.SetState(33)
			p.Match(tinybasicParserCR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = tinybasicParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Exprlist() IExprlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprlistContext)
}

func (s *StatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *StatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *StatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *StatementContext) Varlist() IVarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *StatementContext) Vara() IVaraContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaraContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *tinybasicParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tinybasicParserRULE_statement)
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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tinybasicParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(tinybasicParserT__0)
		}
		{
			p.SetState(38)
			p.Exprlist()
		}

	case tinybasicParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(tinybasicParserT__1)
		}
		{
			p.SetState(40)
			p.Expression()
		}
		{
			p.SetState(41)
			p.Relop()
		}
		{
			p.SetState(42)
			p.Expression()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tinybasicParserT__2 {
			{
				p.SetState(43)
				p.Match(tinybasicParserT__2)
			}

		}
		{
			p.SetState(46)
			p.Statement()
		}

	case tinybasicParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Match(tinybasicParserT__3)
		}
		{
			p.SetState(49)
			p.Number()
		}

	case tinybasicParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)
			p.Match(tinybasicParserT__4)
		}
		{
			p.SetState(51)
			p.Varlist()
		}

	case tinybasicParserT__5, tinybasicParserSTRING, tinybasicParserVAR:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == tinybasicParserT__5 {
			{
				p.SetState(52)
				p.Match(tinybasicParserT__5)
			}

		}
		{
			p.SetState(55)
			p.Vara()
		}
		{
			p.SetState(56)
			p.Match(tinybasicParserT__6)
		}
		{
			p.SetState(57)
			p.Expression()
		}

	case tinybasicParserT__7:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(59)
			p.Match(tinybasicParserT__7)
		}
		{
			p.SetState(60)
			p.Expression()
		}

	case tinybasicParserT__8:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Match(tinybasicParserT__8)
		}

	case tinybasicParserT__9:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.Match(tinybasicParserT__9)
		}

	case tinybasicParserT__10:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.Match(tinybasicParserT__10)
		}

	case tinybasicParserT__11:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(64)
			p.Match(tinybasicParserT__11)
		}

	case tinybasicParserT__12:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(65)
			p.Match(tinybasicParserT__12)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprlistContext is an interface to support dynamic dispatch.
type IExprlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprlistContext differentiates from other interfaces.
	IsExprlistContext()
}

type ExprlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprlistContext() *ExprlistContext {
	var p = new(ExprlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinybasicParserRULE_exprlist
	return p
}

func (*ExprlistContext) IsExprlistContext() {}

func NewExprlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprlistContext {
	var p = new(ExprlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_exprlist

	return p
}

func (s *ExprlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprlistContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(tinybasicParserSTRING)
}

func (s *ExprlistContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(tinybasicParserSTRING, i)
}

func (s *ExprlistContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprlistContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterExprlist(s)
	}
}

func (s *ExprlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitExprlist(s)
	}
}

func (p *tinybasicParser) Exprlist() (localctx IExprlistContext) {
	localctx = NewExprlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tinybasicParserRULE_exprlist)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(68)
			p.Match(tinybasicParserSTRING)
		}

	case 2:
		{
			p.SetState(69)
			p.Expression()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tinybasicParserT__13 {
		{
			p.SetState(72)
			p.Match(tinybasicParserT__13)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(73)
				p.Match(tinybasicParserSTRING)
			}

		case 2:
			{
				p.SetState(74)
				p.Expression()
			}

		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinybasicParserRULE_varlist
	return p
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVara() []IVaraContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVaraContext)(nil)).Elem())
	var tst = make([]IVaraContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVaraContext)
		}
	}

	return tst
}

func (s *VarlistContext) Vara(i int) IVaraContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaraContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterVarlist(s)
	}
}

func (s *VarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitVarlist(s)
	}
}

func (p *tinybasicParser) Varlist() (localctx IVarlistContext) {
	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tinybasicParserRULE_varlist)
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
		p.SetState(82)
		p.Vara()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tinybasicParserT__13 {
		{
			p.SetState(83)
			p.Match(tinybasicParserT__13)
		}
		{
			p.SetState(84)
			p.Vara()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = tinybasicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *tinybasicParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tinybasicParserRULE_expression)
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinybasicParserT__14)|(1<<tinybasicParserT__15)|(1<<tinybasicParserT__16))) != 0 {
		{
			p.SetState(90)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinybasicParserT__14)|(1<<tinybasicParserT__15)|(1<<tinybasicParserT__16))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(93)
		p.Term()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(94)
				_la = p.GetTokenStream().LA(1)

				if !(_la == tinybasicParserT__14 || _la == tinybasicParserT__15) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(95)
				p.Term()
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.RuleIndex = tinybasicParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *tinybasicParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tinybasicParserRULE_term)
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
		p.SetState(101)
		p.Factor()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tinybasicParserT__17 || _la == tinybasicParserT__18 {
		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == tinybasicParserT__17 || _la == tinybasicParserT__18) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(103)
			p.Factor()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinybasicParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Vara() IVaraContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaraContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVaraContext)
}

func (s *FactorContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *tinybasicParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tinybasicParserRULE_factor)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tinybasicParserSTRING, tinybasicParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Vara()
		}

	case tinybasicParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Number()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVaraContext is an interface to support dynamic dispatch.
type IVaraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVaraContext differentiates from other interfaces.
	IsVaraContext()
}

type VaraContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVaraContext() *VaraContext {
	var p = new(VaraContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinybasicParserRULE_vara
	return p
}

func (*VaraContext) IsVaraContext() {}

func NewVaraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VaraContext {
	var p = new(VaraContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_vara

	return p
}

func (s *VaraContext) GetParser() antlr.Parser { return s.parser }

func (s *VaraContext) VAR() antlr.TerminalNode {
	return s.GetToken(tinybasicParserVAR, 0)
}

func (s *VaraContext) STRING() antlr.TerminalNode {
	return s.GetToken(tinybasicParserSTRING, 0)
}

func (s *VaraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VaraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VaraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterVara(s)
	}
}

func (s *VaraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitVara(s)
	}
}

func (p *tinybasicParser) Vara() (localctx IVaraContext) {
	localctx = NewVaraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tinybasicParserRULE_vara)
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
		p.SetState(113)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tinybasicParserSTRING || _la == tinybasicParserVAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = tinybasicParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(tinybasicParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(tinybasicParserDIGIT, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *tinybasicParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tinybasicParserRULE_number)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tinybasicParserDIGIT {
		{
			p.SetState(115)
			p.Match(tinybasicParserDIGIT)
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tinybasicParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tinybasicParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }
func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tinybasicListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *tinybasicParser) Relop() (localctx IRelopContext) {
	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tinybasicParserRULE_relop)
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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tinybasicParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(tinybasicParserT__19)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(121)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinybasicParserT__6)|(1<<tinybasicParserT__16)|(1<<tinybasicParserT__20))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case tinybasicParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Match(tinybasicParserT__20)
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(125)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tinybasicParserT__6)|(1<<tinybasicParserT__16)|(1<<tinybasicParserT__19))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case tinybasicParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(tinybasicParserT__6)
		}

	case tinybasicParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.Match(tinybasicParserT__14)
		}

	case tinybasicParserT__15:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(130)
			p.Match(tinybasicParserT__15)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
