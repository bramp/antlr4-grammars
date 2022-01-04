// Code generated from karel.g4 by ANTLR 4.9.3. DO NOT EDIT.

package karel // karel
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 88, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 7, 2, 25,
	10, 2, 12, 2, 14, 2, 28, 11, 2, 3, 2, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14,
	2, 35, 11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 50, 10, 4, 3, 5, 3, 5, 7, 5, 54, 10, 5, 12,
	5, 14, 5, 57, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 77, 10,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 84, 10, 11, 13, 11, 14, 11,
	85, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 4, 2, 18,
	22, 43, 43, 3, 2, 23, 42, 2, 86, 2, 22, 3, 2, 2, 2, 4, 39, 3, 2, 2, 2,
	6, 49, 3, 2, 2, 2, 8, 51, 3, 2, 2, 2, 10, 60, 3, 2, 2, 2, 12, 65, 3, 2,
	2, 2, 14, 70, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 83,
	3, 2, 2, 2, 22, 26, 7, 3, 2, 2, 23, 25, 5, 4, 3, 2, 24, 23, 3, 2, 2, 2,
	25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 29, 3,
	2, 2, 2, 28, 26, 3, 2, 2, 2, 29, 33, 7, 4, 2, 2, 30, 32, 5, 6, 4, 2, 31,
	30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2,
	2, 34, 36, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37, 7, 5, 2, 2, 37, 38,
	7, 6, 2, 2, 38, 3, 3, 2, 2, 2, 39, 40, 7, 7, 2, 2, 40, 41, 7, 43, 2, 2,
	41, 42, 7, 8, 2, 2, 42, 43, 5, 6, 4, 2, 43, 5, 3, 2, 2, 2, 44, 50, 5, 8,
	5, 2, 45, 50, 5, 10, 6, 2, 46, 50, 5, 12, 7, 2, 47, 50, 5, 14, 8, 2, 48,
	50, 5, 16, 9, 2, 49, 44, 3, 2, 2, 2, 49, 45, 3, 2, 2, 2, 49, 46, 3, 2,
	2, 2, 49, 47, 3, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 7, 3, 2, 2, 2, 51, 55,
	7, 9, 2, 2, 52, 54, 5, 6, 4, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2,
	55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 55, 3,
	2, 2, 2, 58, 59, 7, 10, 2, 2, 59, 9, 3, 2, 2, 2, 60, 61, 7, 11, 2, 2, 61,
	62, 5, 20, 11, 2, 62, 63, 7, 12, 2, 2, 63, 64, 5, 6, 4, 2, 64, 11, 3, 2,
	2, 2, 65, 66, 7, 13, 2, 2, 66, 67, 5, 18, 10, 2, 67, 68, 7, 14, 2, 2, 68,
	69, 5, 6, 4, 2, 69, 13, 3, 2, 2, 2, 70, 71, 7, 15, 2, 2, 71, 72, 5, 18,
	10, 2, 72, 73, 7, 16, 2, 2, 73, 76, 5, 6, 4, 2, 74, 75, 7, 17, 2, 2, 75,
	77, 5, 6, 4, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 15, 3, 2, 2,
	2, 78, 79, 9, 2, 2, 2, 79, 17, 3, 2, 2, 2, 80, 81, 9, 3, 2, 2, 81, 19,
	3, 2, 2, 2, 82, 84, 7, 45, 2, 2, 83, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2,
	85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 21, 3, 2, 2, 2, 8, 26, 33,
	49, 55, 76, 85,
}
var literalNames = []string{
	"", "'BEGINNING-OF-PROGRAM'", "'BEGINNING-OF-EXECUTION'", "'END-OF-EXECUTION'",
	"'END-OF-PROGRAM'", "'DEFINE'", "'AS'", "'BEGIN'", "'END'", "'ITERATE'",
	"'TIMES'", "'WHILE'", "'DO'", "'IF'", "'THEN'", "'ELSE'", "'MOVE'", "'TURNLEFT'",
	"'PICKBEEPER'", "'PUTBEEPER'", "'TURNOFF'", "'FRONT-IS-CLEAR'", "'FRONT-IS-BLOCKED'",
	"'LEFT-IS-CLEAR'", "'LEFT-IS-BLOCKED'", "'RIGHT-IS-CLEAR'", "'RIGHT-IS-BLOCKED'",
	"'BACK-IS-CLEAR'", "'BACK-IS-BLOCKED'", "'NEXT-TO-A-BEEPER'", "'NOT-NEXT-TO-A-BEEPER'",
	"'ANY-BEEPERS-IN-BEEPER-BAG'", "'NO-BEEPERS-IN-BEEPER-BAG'", "'FACING-NORTH'",
	"'NOT-FACING-NORTH'", "'FACING-SOUTH'", "'NOT-FACING-SOUTH'", "'FACING-EAST'",
	"'NOT-FACING-EAST'", "'FACING-WEST'", "'NOT-FACING-WEST'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "IDENTIFIER", "LETTER", "DIGIT", "WS",
}

var ruleNames = []string{
	"karel", "definition", "statement", "block", "iteration", "loop", "conditional",
	"instruction", "condition", "number",
}

type karelParser struct {
	*antlr.BaseParser
}

// NewkarelParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *karelParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewkarelParser(input antlr.TokenStream) *karelParser {
	this := new(karelParser)
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
	this.GrammarFileName = "karel.g4"

	return this
}

// karelParser tokens.
const (
	karelParserEOF        = antlr.TokenEOF
	karelParserT__0       = 1
	karelParserT__1       = 2
	karelParserT__2       = 3
	karelParserT__3       = 4
	karelParserT__4       = 5
	karelParserT__5       = 6
	karelParserT__6       = 7
	karelParserT__7       = 8
	karelParserT__8       = 9
	karelParserT__9       = 10
	karelParserT__10      = 11
	karelParserT__11      = 12
	karelParserT__12      = 13
	karelParserT__13      = 14
	karelParserT__14      = 15
	karelParserT__15      = 16
	karelParserT__16      = 17
	karelParserT__17      = 18
	karelParserT__18      = 19
	karelParserT__19      = 20
	karelParserT__20      = 21
	karelParserT__21      = 22
	karelParserT__22      = 23
	karelParserT__23      = 24
	karelParserT__24      = 25
	karelParserT__25      = 26
	karelParserT__26      = 27
	karelParserT__27      = 28
	karelParserT__28      = 29
	karelParserT__29      = 30
	karelParserT__30      = 31
	karelParserT__31      = 32
	karelParserT__32      = 33
	karelParserT__33      = 34
	karelParserT__34      = 35
	karelParserT__35      = 36
	karelParserT__36      = 37
	karelParserT__37      = 38
	karelParserT__38      = 39
	karelParserT__39      = 40
	karelParserIDENTIFIER = 41
	karelParserLETTER     = 42
	karelParserDIGIT      = 43
	karelParserWS         = 44
)

// karelParser rules.
const (
	karelParserRULE_karel       = 0
	karelParserRULE_definition  = 1
	karelParserRULE_statement   = 2
	karelParserRULE_block       = 3
	karelParserRULE_iteration   = 4
	karelParserRULE_loop        = 5
	karelParserRULE_conditional = 6
	karelParserRULE_instruction = 7
	karelParserRULE_condition   = 8
	karelParserRULE_number      = 9
)

// IKarelContext is an interface to support dynamic dispatch.
type IKarelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKarelContext differentiates from other interfaces.
	IsKarelContext()
}

type KarelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKarelContext() *KarelContext {
	var p = new(KarelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_karel
	return p
}

func (*KarelContext) IsKarelContext() {}

func NewKarelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KarelContext {
	var p = new(KarelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_karel

	return p
}

func (s *KarelContext) GetParser() antlr.Parser { return s.parser }

func (s *KarelContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *KarelContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *KarelContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *KarelContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *KarelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KarelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KarelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterKarel(s)
	}
}

func (s *KarelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitKarel(s)
	}
}

func (p *karelParser) Karel() (localctx IKarelContext) {
	this := p
	_ = this

	localctx = NewKarelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, karelParserRULE_karel)
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
		p.Match(karelParserT__0)
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == karelParserT__4 {
		{
			p.SetState(21)
			p.Definition()
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
		p.Match(karelParserT__1)
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<karelParserT__6)|(1<<karelParserT__8)|(1<<karelParserT__10)|(1<<karelParserT__12)|(1<<karelParserT__15)|(1<<karelParserT__16)|(1<<karelParserT__17)|(1<<karelParserT__18)|(1<<karelParserT__19))) != 0) || _la == karelParserIDENTIFIER {
		{
			p.SetState(28)
			p.Statement()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(karelParserT__2)
	}
	{
		p.SetState(35)
		p.Match(karelParserT__3)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(karelParserIDENTIFIER, 0)
}

func (s *DefinitionContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *karelParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, karelParserRULE_definition)

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
		p.Match(karelParserT__4)
	}
	{
		p.SetState(38)
		p.Match(karelParserIDENTIFIER)
	}
	{
		p.SetState(39)
		p.Match(karelParserT__5)
	}
	{
		p.SetState(40)
		p.Statement()
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
	p.RuleIndex = karelParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Iteration() IIterationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *karelParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, karelParserRULE_statement)

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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case karelParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Block()
		}

	case karelParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Iteration()
		}

	case karelParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Loop()
		}

	case karelParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Conditional()
		}

	case karelParserT__15, karelParserT__16, karelParserT__17, karelParserT__18, karelParserT__19, karelParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.Instruction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *karelParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, karelParserRULE_block)
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
		p.SetState(49)
		p.Match(karelParserT__6)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<karelParserT__6)|(1<<karelParserT__8)|(1<<karelParserT__10)|(1<<karelParserT__12)|(1<<karelParserT__15)|(1<<karelParserT__16)|(1<<karelParserT__17)|(1<<karelParserT__18)|(1<<karelParserT__19))) != 0) || _la == karelParserIDENTIFIER {
		{
			p.SetState(50)
			p.Statement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(karelParserT__7)
	}

	return localctx
}

// IIterationContext is an interface to support dynamic dispatch.
type IIterationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationContext differentiates from other interfaces.
	IsIterationContext()
}

type IterationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationContext() *IterationContext {
	var p = new(IterationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_iteration
	return p
}

func (*IterationContext) IsIterationContext() {}

func NewIterationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationContext {
	var p = new(IterationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_iteration

	return p
}

func (s *IterationContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *IterationContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IterationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterIteration(s)
	}
}

func (s *IterationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitIteration(s)
	}
}

func (p *karelParser) Iteration() (localctx IIterationContext) {
	this := p
	_ = this

	localctx = NewIterationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, karelParserRULE_iteration)

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
		p.SetState(58)
		p.Match(karelParserT__8)
	}
	{
		p.SetState(59)
		p.Number()
	}
	{
		p.SetState(60)
		p.Match(karelParserT__9)
	}
	{
		p.SetState(61)
		p.Statement()
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *LoopContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *karelParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, karelParserRULE_loop)

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
		p.Match(karelParserT__10)
	}
	{
		p.SetState(64)
		p.Condition()
	}
	{
		p.SetState(65)
		p.Match(karelParserT__11)
	}
	{
		p.SetState(66)
		p.Statement()
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionalContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ConditionalContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *karelParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, karelParserRULE_conditional)

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
		p.SetState(68)
		p.Match(karelParserT__12)
	}
	{
		p.SetState(69)
		p.Condition()
	}
	{
		p.SetState(70)
		p.Match(karelParserT__13)
	}
	{
		p.SetState(71)
		p.Statement()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(72)
			p.Match(karelParserT__14)
		}
		{
			p.SetState(73)
			p.Statement()
		}

	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(karelParserIDENTIFIER, 0)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *karelParser) Instruction() (localctx IInstructionContext) {
	this := p
	_ = this

	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, karelParserRULE_instruction)
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
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(karelParserT__15-16))|(1<<(karelParserT__16-16))|(1<<(karelParserT__17-16))|(1<<(karelParserT__18-16))|(1<<(karelParserT__19-16))|(1<<(karelParserIDENTIFIER-16)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = karelParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }
func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *karelParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, karelParserRULE_condition)
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
		p.SetState(78)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(karelParserT__20-21))|(1<<(karelParserT__21-21))|(1<<(karelParserT__22-21))|(1<<(karelParserT__23-21))|(1<<(karelParserT__24-21))|(1<<(karelParserT__25-21))|(1<<(karelParserT__26-21))|(1<<(karelParserT__27-21))|(1<<(karelParserT__28-21))|(1<<(karelParserT__29-21))|(1<<(karelParserT__30-21))|(1<<(karelParserT__31-21))|(1<<(karelParserT__32-21))|(1<<(karelParserT__33-21))|(1<<(karelParserT__34-21))|(1<<(karelParserT__35-21))|(1<<(karelParserT__36-21))|(1<<(karelParserT__37-21))|(1<<(karelParserT__38-21))|(1<<(karelParserT__39-21)))) != 0) {
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
	p.RuleIndex = karelParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = karelParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(karelParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(karelParserDIGIT, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(karelListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *karelParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, karelParserRULE_number)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == karelParserDIGIT {
		{
			p.SetState(80)
			p.Match(karelParserDIGIT)
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
