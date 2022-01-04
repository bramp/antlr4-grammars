// Code generated from callable_.g4 by ANTLR 4.9.3. DO NOT EDIT.

package callable_ // callable_
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 6, 2, 13, 10,
	2, 13, 2, 14, 2, 14, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3,
	3, 3, 3, 3, 3, 5, 3, 26, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 33,
	10, 4, 12, 4, 14, 4, 36, 11, 4, 3, 5, 3, 5, 5, 5, 40, 10, 5, 3, 5, 2, 2,
	6, 2, 4, 6, 8, 2, 2, 2, 42, 2, 10, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 29,
	3, 2, 2, 2, 8, 39, 3, 2, 2, 2, 10, 19, 5, 4, 3, 2, 11, 13, 7, 10, 2, 2,
	12, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 14, 15, 3,
	2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 12, 3, 2, 2, 2, 18,
	21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 3, 3, 2, 2,
	2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 6, 2, 2, 23, 25, 7, 3, 2, 2, 24, 26,
	5, 6, 4, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2,
	27, 28, 7, 4, 2, 2, 28, 5, 3, 2, 2, 2, 29, 34, 5, 8, 5, 2, 30, 31, 7, 5,
	2, 2, 31, 33, 5, 8, 5, 2, 32, 30, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32,
	3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 7, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2,
	37, 40, 5, 4, 3, 2, 38, 40, 7, 7, 2, 2, 39, 37, 3, 2, 2, 2, 39, 38, 3,
	2, 2, 2, 40, 9, 3, 2, 2, 2, 7, 14, 19, 25, 34, 39,
}
var literalNames = []string{
	"", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "ID", "STRING", "LETTER_UPPER", "LETTER_LOWER", "EOL",
	"WS",
}

var ruleNames = []string{
	"program", "line", "f_inner", "f_arg",
}

type callable_Parser struct {
	*antlr.BaseParser
}

// Newcallable_Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *callable_Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newcallable_Parser(input antlr.TokenStream) *callable_Parser {
	this := new(callable_Parser)
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
	this.GrammarFileName = "callable_.g4"

	return this
}

// callable_Parser tokens.
const (
	callable_ParserEOF          = antlr.TokenEOF
	callable_ParserT__0         = 1
	callable_ParserT__1         = 2
	callable_ParserT__2         = 3
	callable_ParserID           = 4
	callable_ParserSTRING       = 5
	callable_ParserLETTER_UPPER = 6
	callable_ParserLETTER_LOWER = 7
	callable_ParserEOL          = 8
	callable_ParserWS           = 9
)

// callable_Parser rules.
const (
	callable_ParserRULE_program = 0
	callable_ParserRULE_line    = 1
	callable_ParserRULE_f_inner = 2
	callable_ParserRULE_f_arg   = 3
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
	p.RuleIndex = callable_ParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = callable_ParserRULE_program

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

func (s *ProgramContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(callable_ParserEOL)
}

func (s *ProgramContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(callable_ParserEOL, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *callable_Parser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, callable_ParserRULE_program)
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
		p.SetState(8)
		p.Line()
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == callable_ParserEOL {
		p.SetState(10)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == callable_ParserEOL {
			{
				p.SetState(9)
				p.Match(callable_ParserEOL)
			}

			p.SetState(12)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(14)
			p.Line()
		}

		p.SetState(19)
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
	p.RuleIndex = callable_ParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = callable_ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) ID() antlr.TerminalNode {
	return s.GetToken(callable_ParserID, 0)
}

func (s *LineContext) F_inner() IF_innerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_innerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_innerContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *callable_Parser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, callable_ParserRULE_line)
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
		p.SetState(20)
		p.Match(callable_ParserID)
	}
	{
		p.SetState(21)
		p.Match(callable_ParserT__0)
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == callable_ParserID || _la == callable_ParserSTRING {
		{
			p.SetState(22)
			p.F_inner()
		}

	}
	{
		p.SetState(25)
		p.Match(callable_ParserT__1)
	}

	return localctx
}

// IF_innerContext is an interface to support dynamic dispatch.
type IF_innerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_innerContext differentiates from other interfaces.
	IsF_innerContext()
}

type F_innerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_innerContext() *F_innerContext {
	var p = new(F_innerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = callable_ParserRULE_f_inner
	return p
}

func (*F_innerContext) IsF_innerContext() {}

func NewF_innerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_innerContext {
	var p = new(F_innerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = callable_ParserRULE_f_inner

	return p
}

func (s *F_innerContext) GetParser() antlr.Parser { return s.parser }

func (s *F_innerContext) AllF_arg() []IF_argContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IF_argContext)(nil)).Elem())
	var tst = make([]IF_argContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IF_argContext)
		}
	}

	return tst
}

func (s *F_innerContext) F_arg(i int) IF_argContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_argContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IF_argContext)
}

func (s *F_innerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_innerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_innerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.EnterF_inner(s)
	}
}

func (s *F_innerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.ExitF_inner(s)
	}
}

func (p *callable_Parser) F_inner() (localctx IF_innerContext) {
	this := p
	_ = this

	localctx = NewF_innerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, callable_ParserRULE_f_inner)
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
		p.SetState(27)
		p.F_arg()
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == callable_ParserT__2 {
		{
			p.SetState(28)
			p.Match(callable_ParserT__2)
		}
		{
			p.SetState(29)
			p.F_arg()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IF_argContext is an interface to support dynamic dispatch.
type IF_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_argContext differentiates from other interfaces.
	IsF_argContext()
}

type F_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_argContext() *F_argContext {
	var p = new(F_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = callable_ParserRULE_f_arg
	return p
}

func (*F_argContext) IsF_argContext() {}

func NewF_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_argContext {
	var p = new(F_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = callable_ParserRULE_f_arg

	return p
}

func (s *F_argContext) GetParser() antlr.Parser { return s.parser }

func (s *F_argContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *F_argContext) STRING() antlr.TerminalNode {
	return s.GetToken(callable_ParserSTRING, 0)
}

func (s *F_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.EnterF_arg(s)
	}
}

func (s *F_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(callable_Listener); ok {
		listenerT.ExitF_arg(s)
	}
}

func (p *callable_Parser) F_arg() (localctx IF_argContext) {
	this := p
	_ = this

	localctx = NewF_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, callable_ParserRULE_f_arg)

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
	case callable_ParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Line()
		}

	case callable_ParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.Match(callable_ParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
