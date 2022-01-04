// Code generated from asm8080.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm8080 // asm8080
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 124,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 5, 2, 40, 10, 2, 3, 2, 6, 2, 43, 10, 2, 13, 2,
	14, 2, 44, 3, 3, 5, 3, 48, 10, 3, 3, 3, 3, 3, 5, 3, 52, 10, 3, 3, 3, 5,
	3, 55, 10, 3, 3, 4, 3, 4, 5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	5, 7, 66, 10, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 75, 10,
	9, 3, 10, 3, 10, 3, 10, 7, 10, 80, 10, 10, 12, 10, 14, 10, 83, 11, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 90, 10, 12, 12, 12, 14, 12, 93,
	11, 12, 3, 13, 3, 13, 3, 13, 7, 13, 98, 10, 13, 12, 13, 14, 13, 101, 11,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	112, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 2, 2, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 2, 4, 3, 2, 5, 6, 3, 2, 7, 8, 2, 122, 2, 42, 3,
	2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 62,
	3, 2, 2, 2, 12, 65, 3, 2, 2, 2, 14, 70, 3, 2, 2, 2, 16, 72, 3, 2, 2, 2,
	18, 76, 3, 2, 2, 2, 20, 84, 3, 2, 2, 2, 22, 86, 3, 2, 2, 2, 24, 94, 3,
	2, 2, 2, 26, 111, 3, 2, 2, 2, 28, 113, 3, 2, 2, 2, 30, 115, 3, 2, 2, 2,
	32, 117, 3, 2, 2, 2, 34, 119, 3, 2, 2, 2, 36, 121, 3, 2, 2, 2, 38, 40,
	5, 4, 3, 2, 39, 38, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 43, 7, 19, 2, 2, 42, 39, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 42, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 48, 5, 16, 9, 2, 47,
	46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 52, 5, 6, 4,
	2, 50, 52, 5, 12, 7, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 51, 52,
	3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 55, 5, 36, 19, 2, 54, 53, 3, 2, 2,
	2, 54, 55, 3, 2, 2, 2, 55, 5, 3, 2, 2, 2, 56, 58, 5, 8, 5, 2, 57, 59, 5,
	18, 10, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2, 60,
	61, 7, 14, 2, 2, 61, 9, 3, 2, 2, 2, 62, 63, 7, 13, 2, 2, 63, 11, 3, 2,
	2, 2, 64, 66, 5, 26, 14, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66,
	67, 3, 2, 2, 2, 67, 68, 5, 14, 8, 2, 68, 69, 5, 18, 10, 2, 69, 13, 3, 2,
	2, 2, 70, 71, 7, 12, 2, 2, 71, 15, 3, 2, 2, 2, 72, 74, 5, 20, 11, 2, 73,
	75, 7, 3, 2, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 17, 3, 2, 2,
	2, 76, 81, 5, 22, 12, 2, 77, 78, 7, 4, 2, 2, 78, 80, 5, 22, 12, 2, 79,
	77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2,
	2, 82, 19, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 5, 32, 17, 2, 85, 21,
	3, 2, 2, 2, 86, 91, 5, 24, 13, 2, 87, 88, 9, 2, 2, 2, 88, 90, 5, 24, 13,
	2, 89, 87, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 23, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 99, 5, 26, 14,
	2, 95, 96, 9, 3, 2, 2, 96, 98, 5, 26, 14, 2, 97, 95, 3, 2, 2, 2, 98, 101,
	3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 25, 3, 2, 2,
	2, 101, 99, 3, 2, 2, 2, 102, 112, 5, 34, 18, 2, 103, 112, 5, 10, 6, 2,
	104, 112, 5, 28, 15, 2, 105, 112, 5, 32, 17, 2, 106, 112, 5, 30, 16, 2,
	107, 108, 7, 9, 2, 2, 108, 109, 5, 22, 12, 2, 109, 110, 7, 10, 2, 2, 110,
	112, 3, 2, 2, 2, 111, 102, 3, 2, 2, 2, 111, 103, 3, 2, 2, 2, 111, 104,
	3, 2, 2, 2, 111, 105, 3, 2, 2, 2, 111, 106, 3, 2, 2, 2, 111, 107, 3, 2,
	2, 2, 112, 27, 3, 2, 2, 2, 113, 114, 7, 11, 2, 2, 114, 29, 3, 2, 2, 2,
	115, 116, 7, 18, 2, 2, 116, 31, 3, 2, 2, 2, 117, 118, 7, 15, 2, 2, 118,
	33, 3, 2, 2, 2, 119, 120, 7, 16, 2, 2, 120, 35, 3, 2, 2, 2, 121, 122, 7,
	17, 2, 2, 122, 37, 3, 2, 2, 2, 14, 39, 44, 47, 51, 54, 58, 65, 74, 81,
	91, 99, 111,
}
var literalNames = []string{
	"", "':'", "','", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'$'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "ASSEMBLER_DIRECTIVE", "REGISTER",
	"OPCODE", "NAME", "NUMBER", "COMMENT", "STRING", "EOL", "WS",
}

var ruleNames = []string{
	"prog", "line", "instruction", "opcode", "register_", "directive", "assemblerdirective",
	"lbl", "expressionlist", "label", "expression", "multiplyingExpression",
	"argument", "dollar", "string_", "name", "number", "comment",
}

type asm8080Parser struct {
	*antlr.BaseParser
}

// Newasm8080Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *asm8080Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newasm8080Parser(input antlr.TokenStream) *asm8080Parser {
	this := new(asm8080Parser)
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
	this.GrammarFileName = "asm8080.g4"

	return this
}

// asm8080Parser tokens.
const (
	asm8080ParserEOF                 = antlr.TokenEOF
	asm8080ParserT__0                = 1
	asm8080ParserT__1                = 2
	asm8080ParserT__2                = 3
	asm8080ParserT__3                = 4
	asm8080ParserT__4                = 5
	asm8080ParserT__5                = 6
	asm8080ParserT__6                = 7
	asm8080ParserT__7                = 8
	asm8080ParserT__8                = 9
	asm8080ParserASSEMBLER_DIRECTIVE = 10
	asm8080ParserREGISTER            = 11
	asm8080ParserOPCODE              = 12
	asm8080ParserNAME                = 13
	asm8080ParserNUMBER              = 14
	asm8080ParserCOMMENT             = 15
	asm8080ParserSTRING              = 16
	asm8080ParserEOL                 = 17
	asm8080ParserWS                  = 18
)

// asm8080Parser rules.
const (
	asm8080ParserRULE_prog                  = 0
	asm8080ParserRULE_line                  = 1
	asm8080ParserRULE_instruction           = 2
	asm8080ParserRULE_opcode                = 3
	asm8080ParserRULE_register_             = 4
	asm8080ParserRULE_directive             = 5
	asm8080ParserRULE_assemblerdirective    = 6
	asm8080ParserRULE_lbl                   = 7
	asm8080ParserRULE_expressionlist        = 8
	asm8080ParserRULE_label                 = 9
	asm8080ParserRULE_expression            = 10
	asm8080ParserRULE_multiplyingExpression = 11
	asm8080ParserRULE_argument              = 12
	asm8080ParserRULE_dollar                = 13
	asm8080ParserRULE_string_               = 14
	asm8080ParserRULE_name                  = 15
	asm8080ParserRULE_number                = 16
	asm8080ParserRULE_comment               = 17
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
	p.RuleIndex = asm8080ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(asm8080ParserEOL)
}

func (s *ProgContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(asm8080ParserEOL, i)
}

func (s *ProgContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *asm8080Parser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, asm8080ParserRULE_prog)
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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8080ParserT__6)|(1<<asm8080ParserT__8)|(1<<asm8080ParserASSEMBLER_DIRECTIVE)|(1<<asm8080ParserREGISTER)|(1<<asm8080ParserOPCODE)|(1<<asm8080ParserNAME)|(1<<asm8080ParserNUMBER)|(1<<asm8080ParserCOMMENT)|(1<<asm8080ParserSTRING)|(1<<asm8080ParserEOL))) != 0) {
		p.SetState(37)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(36)
				p.Line()
			}

		}
		{
			p.SetState(39)
			p.Match(asm8080ParserEOL)
		}

		p.SetState(42)
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
	p.RuleIndex = asm8080ParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Lbl() ILblContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblContext)
}

func (s *LineContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *LineContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *asm8080Parser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, asm8080ParserRULE_line)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Lbl()
		}

	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm8080ParserOPCODE:
		{
			p.SetState(47)
			p.Instruction()
		}

	case asm8080ParserT__6, asm8080ParserT__8, asm8080ParserASSEMBLER_DIRECTIVE, asm8080ParserREGISTER, asm8080ParserNAME, asm8080ParserNUMBER, asm8080ParserSTRING:
		{
			p.SetState(48)
			p.Directive()
		}

	case asm8080ParserCOMMENT, asm8080ParserEOL:

	default:
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8080ParserCOMMENT {
		{
			p.SetState(51)
			p.Comment()
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
	p.RuleIndex = asm8080ParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *InstructionContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *asm8080Parser) Instruction() (localctx IInstructionContext) {
	this := p
	_ = this

	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, asm8080ParserRULE_instruction)
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
		p.Opcode()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8080ParserT__6)|(1<<asm8080ParserT__8)|(1<<asm8080ParserREGISTER)|(1<<asm8080ParserNAME)|(1<<asm8080ParserNUMBER)|(1<<asm8080ParserSTRING))) != 0 {
		{
			p.SetState(55)
			p.Expressionlist()
		}

	}

	return localctx
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) OPCODE() antlr.TerminalNode {
	return s.GetToken(asm8080ParserOPCODE, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *asm8080Parser) Opcode() (localctx IOpcodeContext) {
	this := p
	_ = this

	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, asm8080ParserRULE_opcode)

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
		p.Match(asm8080ParserOPCODE)
	}

	return localctx
}

// IRegister_Context is an interface to support dynamic dispatch.
type IRegister_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegister_Context differentiates from other interfaces.
	IsRegister_Context()
}

type Register_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegister_Context() *Register_Context {
	var p = new(Register_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_register_
	return p
}

func (*Register_Context) IsRegister_Context() {}

func NewRegister_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Register_Context {
	var p = new(Register_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_register_

	return p
}

func (s *Register_Context) GetParser() antlr.Parser { return s.parser }

func (s *Register_Context) REGISTER() antlr.TerminalNode {
	return s.GetToken(asm8080ParserREGISTER, 0)
}

func (s *Register_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Register_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Register_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterRegister_(s)
	}
}

func (s *Register_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitRegister_(s)
	}
}

func (p *asm8080Parser) Register_() (localctx IRegister_Context) {
	this := p
	_ = this

	localctx = NewRegister_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, asm8080ParserRULE_register_)

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
		p.Match(asm8080ParserREGISTER)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Assemblerdirective() IAssemblerdirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerdirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblerdirectiveContext)
}

func (s *DirectiveContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *DirectiveContext) Argument() IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *asm8080Parser) Directive() (localctx IDirectiveContext) {
	this := p
	_ = this

	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, asm8080ParserRULE_directive)
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8080ParserT__6)|(1<<asm8080ParserT__8)|(1<<asm8080ParserREGISTER)|(1<<asm8080ParserNAME)|(1<<asm8080ParserNUMBER)|(1<<asm8080ParserSTRING))) != 0 {
		{
			p.SetState(62)
			p.Argument()
		}

	}
	{
		p.SetState(65)
		p.Assemblerdirective()
	}
	{
		p.SetState(66)
		p.Expressionlist()
	}

	return localctx
}

// IAssemblerdirectiveContext is an interface to support dynamic dispatch.
type IAssemblerdirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblerdirectiveContext differentiates from other interfaces.
	IsAssemblerdirectiveContext()
}

type AssemblerdirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblerdirectiveContext() *AssemblerdirectiveContext {
	var p = new(AssemblerdirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_assemblerdirective
	return p
}

func (*AssemblerdirectiveContext) IsAssemblerdirectiveContext() {}

func NewAssemblerdirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblerdirectiveContext {
	var p = new(AssemblerdirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_assemblerdirective

	return p
}

func (s *AssemblerdirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AssemblerdirectiveContext) ASSEMBLER_DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(asm8080ParserASSEMBLER_DIRECTIVE, 0)
}

func (s *AssemblerdirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblerdirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblerdirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterAssemblerdirective(s)
	}
}

func (s *AssemblerdirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitAssemblerdirective(s)
	}
}

func (p *asm8080Parser) Assemblerdirective() (localctx IAssemblerdirectiveContext) {
	this := p
	_ = this

	localctx = NewAssemblerdirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, asm8080ParserRULE_assemblerdirective)

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
		p.Match(asm8080ParserASSEMBLER_DIRECTIVE)
	}

	return localctx
}

// ILblContext is an interface to support dynamic dispatch.
type ILblContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLblContext differentiates from other interfaces.
	IsLblContext()
}

type LblContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLblContext() *LblContext {
	var p = new(LblContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_lbl
	return p
}

func (*LblContext) IsLblContext() {}

func NewLblContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LblContext {
	var p = new(LblContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_lbl

	return p
}

func (s *LblContext) GetParser() antlr.Parser { return s.parser }

func (s *LblContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LblContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LblContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LblContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterLbl(s)
	}
}

func (s *LblContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitLbl(s)
	}
}

func (p *asm8080Parser) Lbl() (localctx ILblContext) {
	this := p
	_ = this

	localctx = NewLblContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, asm8080ParserRULE_lbl)
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
		p.Label()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8080ParserT__0 {
		{
			p.SetState(71)
			p.Match(asm8080ParserT__0)
		}

	}

	return localctx
}

// IExpressionlistContext is an interface to support dynamic dispatch.
type IExpressionlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionlistContext differentiates from other interfaces.
	IsExpressionlistContext()
}

type ExpressionlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionlistContext() *ExpressionlistContext {
	var p = new(ExpressionlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_expressionlist
	return p
}

func (*ExpressionlistContext) IsExpressionlistContext() {}

func NewExpressionlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionlistContext {
	var p = new(ExpressionlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_expressionlist

	return p
}

func (s *ExpressionlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionlistContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionlistContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterExpressionlist(s)
	}
}

func (s *ExpressionlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitExpressionlist(s)
	}
}

func (p *asm8080Parser) Expressionlist() (localctx IExpressionlistContext) {
	this := p
	_ = this

	localctx = NewExpressionlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, asm8080ParserRULE_expressionlist)
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
		p.Expression()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == asm8080ParserT__1 {
		{
			p.SetState(75)
			p.Match(asm8080ParserT__1)
		}
		{
			p.SetState(76)
			p.Expression()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *asm8080Parser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, asm8080ParserRULE_label)

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
		p.Name()
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
	p.RuleIndex = asm8080ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *asm8080Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, asm8080ParserRULE_expression)
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
		p.SetState(84)
		p.MultiplyingExpression()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == asm8080ParserT__2 || _la == asm8080ParserT__3 {
		{
			p.SetState(85)
			_la = p.GetTokenStream().LA(1)

			if !(_la == asm8080ParserT__2 || _la == asm8080ParserT__3) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(86)
			p.MultiplyingExpression()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *asm8080Parser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	this := p
	_ = this

	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, asm8080ParserRULE_multiplyingExpression)
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
		p.SetState(92)
		p.Argument()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == asm8080ParserT__4 || _la == asm8080ParserT__5 {
		{
			p.SetState(93)
			_la = p.GetTokenStream().LA(1)

			if !(_la == asm8080ParserT__4 || _la == asm8080ParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(94)
			p.Argument()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgumentContext) Register_() IRegister_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegister_Context)
}

func (s *ArgumentContext) Dollar() IDollarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDollarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDollarContext)
}

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *asm8080Parser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, asm8080ParserRULE_argument)

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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm8080ParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Number()
		}

	case asm8080ParserREGISTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Register_()
		}

	case asm8080ParserT__8:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Dollar()
		}

	case asm8080ParserNAME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Name()
		}

	case asm8080ParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.String_()
		}

	case asm8080ParserT__6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.Match(asm8080ParserT__6)
		}
		{
			p.SetState(106)
			p.Expression()
		}
		{
			p.SetState(107)
			p.Match(asm8080ParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDollarContext is an interface to support dynamic dispatch.
type IDollarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDollarContext differentiates from other interfaces.
	IsDollarContext()
}

type DollarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDollarContext() *DollarContext {
	var p = new(DollarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_dollar
	return p
}

func (*DollarContext) IsDollarContext() {}

func NewDollarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DollarContext {
	var p = new(DollarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_dollar

	return p
}

func (s *DollarContext) GetParser() antlr.Parser { return s.parser }
func (s *DollarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DollarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DollarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterDollar(s)
	}
}

func (s *DollarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitDollar(s)
	}
}

func (p *asm8080Parser) Dollar() (localctx IDollarContext) {
	this := p
	_ = this

	localctx = NewDollarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, asm8080ParserRULE_dollar)

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
		p.SetState(111)
		p.Match(asm8080ParserT__8)
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(asm8080ParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *asm8080Parser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, asm8080ParserRULE_string_)

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
		p.Match(asm8080ParserSTRING)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(asm8080ParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitName(s)
	}
}

func (p *asm8080Parser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, asm8080ParserRULE_name)

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
		p.SetState(115)
		p.Match(asm8080ParserNAME)
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
	p.RuleIndex = asm8080ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(asm8080ParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *asm8080Parser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, asm8080ParserRULE_number)

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
		p.SetState(117)
		p.Match(asm8080ParserNUMBER)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8080ParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8080ParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(asm8080ParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8080Listener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *asm8080Parser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, asm8080ParserRULE_comment)

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
		p.SetState(119)
		p.Match(asm8080ParserCOMMENT)
	}

	return localctx
}
