// Code generated from scotty.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scotty // scotty
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 92, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 65, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 70, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 78, 10, 8, 3, 9, 3, 9, 3, 9, 5,
	9, 83, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 90, 10, 10, 3,
	10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 97, 2, 20, 3, 2,
	2, 2, 4, 29, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2, 10, 64,
	3, 2, 2, 2, 12, 69, 3, 2, 2, 2, 14, 77, 3, 2, 2, 2, 16, 82, 3, 2, 2, 2,
	18, 89, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 3, 3, 2, 2, 2, 22, 30, 5, 10,
	6, 2, 23, 24, 5, 8, 5, 2, 24, 25, 5, 4, 3, 2, 25, 30, 3, 2, 2, 2, 26, 27,
	5, 6, 4, 2, 27, 28, 5, 4, 3, 2, 28, 30, 3, 2, 2, 2, 29, 22, 3, 2, 2, 2,
	29, 23, 3, 2, 2, 2, 29, 26, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31, 32, 5, 12,
	7, 2, 32, 33, 7, 3, 2, 2, 33, 34, 5, 10, 6, 2, 34, 7, 3, 2, 2, 2, 35, 36,
	7, 4, 2, 2, 36, 37, 5, 12, 7, 2, 37, 38, 5, 12, 7, 2, 38, 39, 7, 3, 2,
	2, 39, 40, 5, 10, 6, 2, 40, 9, 3, 2, 2, 2, 41, 42, 7, 5, 2, 2, 42, 43,
	5, 10, 6, 2, 43, 44, 5, 10, 6, 2, 44, 65, 3, 2, 2, 2, 45, 46, 7, 6, 2,
	2, 46, 47, 5, 10, 6, 2, 47, 48, 5, 10, 6, 2, 48, 65, 3, 2, 2, 2, 49, 50,
	7, 7, 2, 2, 50, 51, 5, 10, 6, 2, 51, 52, 5, 10, 6, 2, 52, 65, 3, 2, 2,
	2, 53, 54, 7, 8, 2, 2, 54, 55, 5, 10, 6, 2, 55, 56, 5, 10, 6, 2, 56, 65,
	3, 2, 2, 2, 57, 58, 7, 9, 2, 2, 58, 59, 5, 12, 7, 2, 59, 60, 5, 10, 6,
	2, 60, 61, 7, 10, 2, 2, 61, 65, 3, 2, 2, 2, 62, 65, 5, 12, 7, 2, 63, 65,
	5, 16, 9, 2, 64, 41, 3, 2, 2, 2, 64, 45, 3, 2, 2, 2, 64, 49, 3, 2, 2, 2,
	64, 53, 3, 2, 2, 2, 64, 57, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 63, 3,
	2, 2, 2, 65, 11, 3, 2, 2, 2, 66, 67, 7, 13, 2, 2, 67, 70, 5, 14, 8, 2,
	68, 70, 7, 13, 2, 2, 69, 66, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 13, 3,
	2, 2, 2, 71, 72, 7, 13, 2, 2, 72, 78, 5, 14, 8, 2, 73, 74, 7, 12, 2, 2,
	74, 78, 5, 14, 8, 2, 75, 78, 7, 13, 2, 2, 76, 78, 7, 12, 2, 2, 77, 71,
	3, 2, 2, 2, 77, 73, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2,
	78, 15, 3, 2, 2, 2, 79, 80, 7, 6, 2, 2, 80, 83, 5, 18, 10, 2, 81, 83, 5,
	18, 10, 2, 82, 79, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 17, 3, 2, 2, 2,
	84, 85, 7, 12, 2, 2, 85, 90, 5, 18, 10, 2, 86, 87, 7, 11, 2, 2, 87, 90,
	5, 18, 10, 2, 88, 90, 7, 12, 2, 2, 89, 84, 3, 2, 2, 2, 89, 86, 3, 2, 2,
	2, 89, 88, 3, 2, 2, 2, 90, 19, 3, 2, 2, 2, 8, 29, 64, 69, 77, 82, 89,
}
var literalNames = []string{
	"", "'='", "'fun'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "DIGIT", "LETTER", "WS",
}

var ruleNames = []string{
	"prog", "program_lines", "var_assign", "fn_def", "prefix_exp", "identifier",
	"id_tail", "number", "digits",
}

type scottyParser struct {
	*antlr.BaseParser
}

// NewscottyParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *scottyParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewscottyParser(input antlr.TokenStream) *scottyParser {
	this := new(scottyParser)
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
	this.GrammarFileName = "scotty.g4"

	return this
}

// scottyParser tokens.
const (
	scottyParserEOF    = antlr.TokenEOF
	scottyParserT__0   = 1
	scottyParserT__1   = 2
	scottyParserT__2   = 3
	scottyParserT__3   = 4
	scottyParserT__4   = 5
	scottyParserT__5   = 6
	scottyParserT__6   = 7
	scottyParserT__7   = 8
	scottyParserT__8   = 9
	scottyParserDIGIT  = 10
	scottyParserLETTER = 11
	scottyParserWS     = 12
)

// scottyParser rules.
const (
	scottyParserRULE_prog          = 0
	scottyParserRULE_program_lines = 1
	scottyParserRULE_var_assign    = 2
	scottyParserRULE_fn_def        = 3
	scottyParserRULE_prefix_exp    = 4
	scottyParserRULE_identifier    = 5
	scottyParserRULE_id_tail       = 6
	scottyParserRULE_number        = 7
	scottyParserRULE_digits        = 8
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
	p.RuleIndex = scottyParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Program_lines() IProgram_linesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram_linesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram_linesContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *scottyParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, scottyParserRULE_prog)

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
		p.SetState(18)
		p.Program_lines()
	}

	return localctx
}

// IProgram_linesContext is an interface to support dynamic dispatch.
type IProgram_linesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgram_linesContext differentiates from other interfaces.
	IsProgram_linesContext()
}

type Program_linesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgram_linesContext() *Program_linesContext {
	var p = new(Program_linesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_program_lines
	return p
}

func (*Program_linesContext) IsProgram_linesContext() {}

func NewProgram_linesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program_linesContext {
	var p = new(Program_linesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_program_lines

	return p
}

func (s *Program_linesContext) GetParser() antlr.Parser { return s.parser }

func (s *Program_linesContext) Prefix_exp() IPrefix_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_expContext)
}

func (s *Program_linesContext) Fn_def() IFn_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFn_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFn_defContext)
}

func (s *Program_linesContext) Program_lines() IProgram_linesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram_linesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram_linesContext)
}

func (s *Program_linesContext) Var_assign() IVar_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_assignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_assignContext)
}

func (s *Program_linesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Program_linesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Program_linesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterProgram_lines(s)
	}
}

func (s *Program_linesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitProgram_lines(s)
	}
}

func (p *scottyParser) Program_lines() (localctx IProgram_linesContext) {
	this := p
	_ = this

	localctx = NewProgram_linesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, scottyParserRULE_program_lines)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Prefix_exp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Fn_def()
		}
		{
			p.SetState(22)
			p.Program_lines()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Var_assign()
		}
		{
			p.SetState(25)
			p.Program_lines()
		}

	}

	return localctx
}

// IVar_assignContext is an interface to support dynamic dispatch.
type IVar_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_assignContext differentiates from other interfaces.
	IsVar_assignContext()
}

type Var_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignContext() *Var_assignContext {
	var p = new(Var_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_var_assign
	return p
}

func (*Var_assignContext) IsVar_assignContext() {}

func NewVar_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignContext {
	var p = new(Var_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_var_assign

	return p
}

func (s *Var_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Var_assignContext) Prefix_exp() IPrefix_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_expContext)
}

func (s *Var_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterVar_assign(s)
	}
}

func (s *Var_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitVar_assign(s)
	}
}

func (p *scottyParser) Var_assign() (localctx IVar_assignContext) {
	this := p
	_ = this

	localctx = NewVar_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, scottyParserRULE_var_assign)

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
		p.SetState(29)
		p.Identifier()
	}
	{
		p.SetState(30)
		p.Match(scottyParserT__0)
	}
	{
		p.SetState(31)
		p.Prefix_exp()
	}

	return localctx
}

// IFn_defContext is an interface to support dynamic dispatch.
type IFn_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFn_defContext differentiates from other interfaces.
	IsFn_defContext()
}

type Fn_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_defContext() *Fn_defContext {
	var p = new(Fn_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_fn_def
	return p
}

func (*Fn_defContext) IsFn_defContext() {}

func NewFn_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_defContext {
	var p = new(Fn_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_fn_def

	return p
}

func (s *Fn_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_defContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *Fn_defContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Fn_defContext) Prefix_exp() IPrefix_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefix_expContext)
}

func (s *Fn_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterFn_def(s)
	}
}

func (s *Fn_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitFn_def(s)
	}
}

func (p *scottyParser) Fn_def() (localctx IFn_defContext) {
	this := p
	_ = this

	localctx = NewFn_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, scottyParserRULE_fn_def)

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
		p.Match(scottyParserT__1)
	}
	{
		p.SetState(34)
		p.Identifier()
	}
	{
		p.SetState(35)
		p.Identifier()
	}
	{
		p.SetState(36)
		p.Match(scottyParserT__0)
	}
	{
		p.SetState(37)
		p.Prefix_exp()
	}

	return localctx
}

// IPrefix_expContext is an interface to support dynamic dispatch.
type IPrefix_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefix_expContext differentiates from other interfaces.
	IsPrefix_expContext()
}

type Prefix_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_expContext() *Prefix_expContext {
	var p = new(Prefix_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_prefix_exp
	return p
}

func (*Prefix_expContext) IsPrefix_expContext() {}

func NewPrefix_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_expContext {
	var p = new(Prefix_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_prefix_exp

	return p
}

func (s *Prefix_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Prefix_expContext) AllPrefix_exp() []IPrefix_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrefix_expContext)(nil)).Elem())
	var tst = make([]IPrefix_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrefix_expContext)
		}
	}

	return tst
}

func (s *Prefix_expContext) Prefix_exp(i int) IPrefix_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefix_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrefix_expContext)
}

func (s *Prefix_expContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Prefix_expContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Prefix_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prefix_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterPrefix_exp(s)
	}
}

func (s *Prefix_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitPrefix_exp(s)
	}
}

func (p *scottyParser) Prefix_exp() (localctx IPrefix_expContext) {
	this := p
	_ = this

	localctx = NewPrefix_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, scottyParserRULE_prefix_exp)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(scottyParserT__2)
		}
		{
			p.SetState(40)
			p.Prefix_exp()
		}
		{
			p.SetState(41)
			p.Prefix_exp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Match(scottyParserT__3)
		}
		{
			p.SetState(44)
			p.Prefix_exp()
		}
		{
			p.SetState(45)
			p.Prefix_exp()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.Match(scottyParserT__4)
		}
		{
			p.SetState(48)
			p.Prefix_exp()
		}
		{
			p.SetState(49)
			p.Prefix_exp()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Match(scottyParserT__5)
		}
		{
			p.SetState(52)
			p.Prefix_exp()
		}
		{
			p.SetState(53)
			p.Prefix_exp()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.Match(scottyParserT__6)
		}
		{
			p.SetState(56)
			p.Identifier()
		}
		{
			p.SetState(57)
			p.Prefix_exp()
		}
		{
			p.SetState(58)
			p.Match(scottyParserT__7)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Identifier()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Number()
		}

	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) LETTER() antlr.TerminalNode {
	return s.GetToken(scottyParserLETTER, 0)
}

func (s *IdentifierContext) Id_tail() IId_tailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_tailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_tailContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *scottyParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, scottyParserRULE_identifier)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(scottyParserLETTER)
		}
		{
			p.SetState(65)
			p.Id_tail()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(scottyParserLETTER)
		}

	}

	return localctx
}

// IId_tailContext is an interface to support dynamic dispatch.
type IId_tailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_tailContext differentiates from other interfaces.
	IsId_tailContext()
}

type Id_tailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_tailContext() *Id_tailContext {
	var p = new(Id_tailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_id_tail
	return p
}

func (*Id_tailContext) IsId_tailContext() {}

func NewId_tailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_tailContext {
	var p = new(Id_tailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_id_tail

	return p
}

func (s *Id_tailContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_tailContext) LETTER() antlr.TerminalNode {
	return s.GetToken(scottyParserLETTER, 0)
}

func (s *Id_tailContext) Id_tail() IId_tailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_tailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_tailContext)
}

func (s *Id_tailContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(scottyParserDIGIT, 0)
}

func (s *Id_tailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_tailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_tailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterId_tail(s)
	}
}

func (s *Id_tailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitId_tail(s)
	}
}

func (p *scottyParser) Id_tail() (localctx IId_tailContext) {
	this := p
	_ = this

	localctx = NewId_tailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, scottyParserRULE_id_tail)

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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(scottyParserLETTER)
		}
		{
			p.SetState(70)
			p.Id_tail()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(scottyParserDIGIT)
		}
		{
			p.SetState(72)
			p.Id_tail()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(scottyParserLETTER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Match(scottyParserDIGIT)
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
	p.RuleIndex = scottyParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Digits() IDigitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitsContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *scottyParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, scottyParserRULE_number)

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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case scottyParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(scottyParserT__3)
		}
		{
			p.SetState(78)
			p.Digits()
		}

	case scottyParserT__8, scottyParserDIGIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Digits()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDigitsContext is an interface to support dynamic dispatch.
type IDigitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDigitsContext differentiates from other interfaces.
	IsDigitsContext()
}

type DigitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDigitsContext() *DigitsContext {
	var p = new(DigitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = scottyParserRULE_digits
	return p
}

func (*DigitsContext) IsDigitsContext() {}

func NewDigitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitsContext {
	var p = new(DigitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = scottyParserRULE_digits

	return p
}

func (s *DigitsContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitsContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(scottyParserDIGIT, 0)
}

func (s *DigitsContext) Digits() IDigitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDigitsContext)
}

func (s *DigitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DigitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.EnterDigits(s)
	}
}

func (s *DigitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(scottyListener); ok {
		listenerT.ExitDigits(s)
	}
}

func (p *scottyParser) Digits() (localctx IDigitsContext) {
	this := p
	_ = this

	localctx = NewDigitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, scottyParserRULE_digits)

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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(scottyParserDIGIT)
		}
		{
			p.SetState(83)
			p.Digits()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Match(scottyParserT__8)
		}
		{
			p.SetState(85)
			p.Digits()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Match(scottyParserDIGIT)
		}

	}

	return localctx
}
