// Generated from StackTrace.g4 by ANTLR 4.7.

package stacktrace // StackTrace
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 127,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 41, 10, 3, 13, 3, 14, 3, 42, 3, 3, 7,
	3, 46, 10, 3, 12, 3, 14, 3, 49, 11, 3, 3, 3, 5, 3, 52, 10, 3, 3, 4, 3,
	4, 5, 4, 56, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8,
	77, 10, 8, 3, 9, 5, 9, 80, 10, 9, 3, 9, 3, 9, 7, 9, 84, 10, 9, 12, 9, 14,
	9, 87, 11, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 97, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 103, 10, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 6, 15, 112, 10, 15, 13, 15, 14,
	15, 113, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 7, 18, 122, 10, 18,
	12, 18, 14, 18, 125, 11, 18, 3, 18, 3, 123, 2, 19, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 2, 2, 123, 2, 36, 3, 2,
	2, 2, 4, 40, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10, 67,
	3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 74, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2,
	18, 88, 3, 2, 2, 2, 20, 96, 3, 2, 2, 2, 22, 98, 3, 2, 2, 2, 24, 104, 3,
	2, 2, 2, 26, 106, 3, 2, 2, 2, 28, 111, 3, 2, 2, 2, 30, 115, 3, 2, 2, 2,
	32, 117, 3, 2, 2, 2, 34, 119, 3, 2, 2, 2, 36, 37, 5, 4, 3, 2, 37, 38, 7,
	2, 2, 3, 38, 3, 3, 2, 2, 2, 39, 41, 5, 14, 8, 2, 40, 39, 3, 2, 2, 2, 41,
	42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 47, 3, 2, 2,
	2, 44, 46, 5, 6, 4, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45,
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2,
	50, 52, 5, 10, 6, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 5, 3,
	2, 2, 2, 53, 56, 5, 8, 5, 2, 54, 56, 5, 12, 7, 2, 55, 53, 3, 2, 2, 2, 55,
	54, 3, 2, 2, 2, 56, 7, 3, 2, 2, 2, 57, 58, 7, 10, 2, 2, 58, 59, 5, 22,
	12, 2, 59, 60, 7, 3, 2, 2, 60, 63, 5, 20, 11, 2, 61, 62, 7, 14, 2, 2, 62,
	64, 7, 7, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 3, 2, 2,
	2, 65, 66, 7, 4, 2, 2, 66, 9, 3, 2, 2, 2, 67, 68, 7, 11, 2, 2, 68, 69,
	5, 4, 3, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 13, 2, 2, 71, 72, 7, 7, 2, 2,
	72, 73, 7, 12, 2, 2, 73, 13, 3, 2, 2, 2, 74, 76, 5, 16, 9, 2, 75, 77, 5,
	34, 18, 2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 15, 3, 2, 2, 2,
	78, 80, 5, 28, 15, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 3,
	2, 2, 2, 81, 85, 5, 30, 16, 2, 82, 84, 5, 18, 10, 2, 83, 82, 3, 2, 2, 2,
	84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 17, 3,
	2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 7, 5, 2, 2, 89, 90, 5, 30, 16, 2,
	90, 19, 3, 2, 2, 2, 91, 92, 5, 32, 17, 2, 92, 93, 7, 6, 2, 2, 93, 97, 3,
	2, 2, 2, 94, 97, 7, 15, 2, 2, 95, 97, 7, 16, 2, 2, 96, 91, 3, 2, 2, 2,
	96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 21, 3, 2, 2, 2, 98, 99, 5,
	16, 9, 2, 99, 102, 7, 9, 2, 2, 100, 103, 5, 26, 14, 2, 101, 103, 5, 24,
	13, 2, 102, 100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2,
	103, 23, 3, 2, 2, 2, 104, 105, 7, 17, 2, 2, 105, 25, 3, 2, 2, 2, 106, 107,
	5, 32, 17, 2, 107, 27, 3, 2, 2, 2, 108, 109, 5, 32, 17, 2, 109, 110, 7,
	9, 2, 2, 110, 112, 3, 2, 2, 2, 111, 108, 3, 2, 2, 2, 112, 113, 3, 2, 2,
	2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 29, 3, 2, 2, 2, 115,
	116, 7, 8, 2, 2, 116, 31, 3, 2, 2, 2, 117, 118, 7, 8, 2, 2, 118, 33, 3,
	2, 2, 2, 119, 123, 7, 14, 2, 2, 120, 122, 11, 2, 2, 2, 121, 120, 3, 2,
	2, 2, 122, 125, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2,
	124, 35, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 14, 42, 47, 51, 55, 63, 76,
	79, 85, 96, 102, 113, 123,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'$'", "'.java'", "", "", "'.'", "'at'", "'Caused by:'",
	"'more'", "'...'", "':'", "'Native Method'", "'Unknown Source'", "'<init>'",
	"", "", "'_'",
}
var symbolicNames = []string{
	"", "", "", "", "", "Number", "JavaWord", "DOT", "AT", "CAUSED_BY", "MORE_",
	"ELLIPSIS", "COLON", "NATIVE_METHOD", "UNKNOWN_SOURCE", "INIT", "NonCapitalLetter",
	"CapitalLetter", "Symbol", "Digit", "WS",
}

var ruleNames = []string{
	"startRule", "stackTrace", "stackTraceLine", "atLine", "causedByLine",
	"ellipsisLine", "messageLine", "qualifiedClass", "innerClassName", "classFile",
	"qualifiedMethod", "constructor", "methodName", "packagePath", "className",
	"identifier", "message",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type StackTraceParser struct {
	*antlr.BaseParser
}

func NewStackTraceParser(input antlr.TokenStream) *StackTraceParser {
	this := new(StackTraceParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "StackTrace.g4"

	return this
}

// StackTraceParser tokens.
const (
	StackTraceParserEOF              = antlr.TokenEOF
	StackTraceParserT__0             = 1
	StackTraceParserT__1             = 2
	StackTraceParserT__2             = 3
	StackTraceParserT__3             = 4
	StackTraceParserNumber           = 5
	StackTraceParserJavaWord         = 6
	StackTraceParserDOT              = 7
	StackTraceParserAT               = 8
	StackTraceParserCAUSED_BY        = 9
	StackTraceParserMORE_            = 10
	StackTraceParserELLIPSIS         = 11
	StackTraceParserCOLON            = 12
	StackTraceParserNATIVE_METHOD    = 13
	StackTraceParserUNKNOWN_SOURCE   = 14
	StackTraceParserINIT             = 15
	StackTraceParserNonCapitalLetter = 16
	StackTraceParserCapitalLetter    = 17
	StackTraceParserSymbol           = 18
	StackTraceParserDigit            = 19
	StackTraceParserWS               = 20
)

// StackTraceParser rules.
const (
	StackTraceParserRULE_startRule       = 0
	StackTraceParserRULE_stackTrace      = 1
	StackTraceParserRULE_stackTraceLine  = 2
	StackTraceParserRULE_atLine          = 3
	StackTraceParserRULE_causedByLine    = 4
	StackTraceParserRULE_ellipsisLine    = 5
	StackTraceParserRULE_messageLine     = 6
	StackTraceParserRULE_qualifiedClass  = 7
	StackTraceParserRULE_innerClassName  = 8
	StackTraceParserRULE_classFile       = 9
	StackTraceParserRULE_qualifiedMethod = 10
	StackTraceParserRULE_constructor     = 11
	StackTraceParserRULE_methodName      = 12
	StackTraceParserRULE_packagePath     = 13
	StackTraceParserRULE_className       = 14
	StackTraceParserRULE_identifier      = 15
	StackTraceParserRULE_message         = 16
)

// IStartRuleContext is an interface to support dynamic dispatch.
type IStartRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartRuleContext differentiates from other interfaces.
	IsStartRuleContext()
}

type StartRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartRuleContext() *StartRuleContext {
	var p = new(StartRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_startRule
	return p
}

func (*StartRuleContext) IsStartRuleContext() {}

func NewStartRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartRuleContext {
	var p = new(StartRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_startRule

	return p
}

func (s *StartRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *StartRuleContext) StackTrace() IStackTraceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStackTraceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStackTraceContext)
}

func (s *StartRuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(StackTraceParserEOF, 0)
}

func (s *StartRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterStartRule(s)
	}
}

func (s *StartRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitStartRule(s)
	}
}

func (p *StackTraceParser) StartRule() (localctx IStartRuleContext) {
	localctx = NewStartRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StackTraceParserRULE_startRule)

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
		p.SetState(34)
		p.StackTrace()
	}
	{
		p.SetState(35)
		p.Match(StackTraceParserEOF)
	}

	return localctx
}

// IStackTraceContext is an interface to support dynamic dispatch.
type IStackTraceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStackTraceContext differentiates from other interfaces.
	IsStackTraceContext()
}

type StackTraceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStackTraceContext() *StackTraceContext {
	var p = new(StackTraceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_stackTrace
	return p
}

func (*StackTraceContext) IsStackTraceContext() {}

func NewStackTraceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StackTraceContext {
	var p = new(StackTraceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_stackTrace

	return p
}

func (s *StackTraceContext) GetParser() antlr.Parser { return s.parser }

func (s *StackTraceContext) AllMessageLine() []IMessageLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageLineContext)(nil)).Elem())
	var tst = make([]IMessageLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageLineContext)
		}
	}

	return tst
}

func (s *StackTraceContext) MessageLine(i int) IMessageLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageLineContext)
}

func (s *StackTraceContext) AllStackTraceLine() []IStackTraceLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStackTraceLineContext)(nil)).Elem())
	var tst = make([]IStackTraceLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStackTraceLineContext)
		}
	}

	return tst
}

func (s *StackTraceContext) StackTraceLine(i int) IStackTraceLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStackTraceLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStackTraceLineContext)
}

func (s *StackTraceContext) CausedByLine() ICausedByLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICausedByLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICausedByLineContext)
}

func (s *StackTraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StackTraceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StackTraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterStackTrace(s)
	}
}

func (s *StackTraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitStackTrace(s)
	}
}

func (p *StackTraceParser) StackTrace() (localctx IStackTraceContext) {
	localctx = NewStackTraceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StackTraceParserRULE_stackTrace)
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

	for ok := true; ok; ok = _la == StackTraceParserJavaWord {
		{
			p.SetState(37)
			p.MessageLine()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StackTraceParserAT || _la == StackTraceParserELLIPSIS {
		{
			p.SetState(42)
			p.StackTraceLine()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StackTraceParserCAUSED_BY {
		{
			p.SetState(48)
			p.CausedByLine()
		}

	}

	return localctx
}

// IStackTraceLineContext is an interface to support dynamic dispatch.
type IStackTraceLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStackTraceLineContext differentiates from other interfaces.
	IsStackTraceLineContext()
}

type StackTraceLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStackTraceLineContext() *StackTraceLineContext {
	var p = new(StackTraceLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_stackTraceLine
	return p
}

func (*StackTraceLineContext) IsStackTraceLineContext() {}

func NewStackTraceLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StackTraceLineContext {
	var p = new(StackTraceLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_stackTraceLine

	return p
}

func (s *StackTraceLineContext) GetParser() antlr.Parser { return s.parser }

func (s *StackTraceLineContext) AtLine() IAtLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtLineContext)
}

func (s *StackTraceLineContext) EllipsisLine() IEllipsisLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEllipsisLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEllipsisLineContext)
}

func (s *StackTraceLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StackTraceLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StackTraceLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterStackTraceLine(s)
	}
}

func (s *StackTraceLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitStackTraceLine(s)
	}
}

func (p *StackTraceParser) StackTraceLine() (localctx IStackTraceLineContext) {
	localctx = NewStackTraceLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StackTraceParserRULE_stackTraceLine)

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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StackTraceParserAT:
		{
			p.SetState(51)
			p.AtLine()
		}

	case StackTraceParserELLIPSIS:
		{
			p.SetState(52)
			p.EllipsisLine()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtLineContext is an interface to support dynamic dispatch.
type IAtLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtLineContext differentiates from other interfaces.
	IsAtLineContext()
}

type AtLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtLineContext() *AtLineContext {
	var p = new(AtLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_atLine
	return p
}

func (*AtLineContext) IsAtLineContext() {}

func NewAtLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtLineContext {
	var p = new(AtLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_atLine

	return p
}

func (s *AtLineContext) GetParser() antlr.Parser { return s.parser }

func (s *AtLineContext) AT() antlr.TerminalNode {
	return s.GetToken(StackTraceParserAT, 0)
}

func (s *AtLineContext) QualifiedMethod() IQualifiedMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedMethodContext)
}

func (s *AtLineContext) ClassFile() IClassFileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassFileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassFileContext)
}

func (s *AtLineContext) COLON() antlr.TerminalNode {
	return s.GetToken(StackTraceParserCOLON, 0)
}

func (s *AtLineContext) Number() antlr.TerminalNode {
	return s.GetToken(StackTraceParserNumber, 0)
}

func (s *AtLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterAtLine(s)
	}
}

func (s *AtLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitAtLine(s)
	}
}

func (p *StackTraceParser) AtLine() (localctx IAtLineContext) {
	localctx = NewAtLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StackTraceParserRULE_atLine)
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
		p.SetState(55)
		p.Match(StackTraceParserAT)
	}
	{
		p.SetState(56)
		p.QualifiedMethod()
	}
	{
		p.SetState(57)
		p.Match(StackTraceParserT__0)
	}
	{
		p.SetState(58)
		p.ClassFile()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StackTraceParserCOLON {
		{
			p.SetState(59)
			p.Match(StackTraceParserCOLON)
		}
		{
			p.SetState(60)
			p.Match(StackTraceParserNumber)
		}

	}
	{
		p.SetState(63)
		p.Match(StackTraceParserT__1)
	}

	return localctx
}

// ICausedByLineContext is an interface to support dynamic dispatch.
type ICausedByLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCausedByLineContext differentiates from other interfaces.
	IsCausedByLineContext()
}

type CausedByLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCausedByLineContext() *CausedByLineContext {
	var p = new(CausedByLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_causedByLine
	return p
}

func (*CausedByLineContext) IsCausedByLineContext() {}

func NewCausedByLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CausedByLineContext {
	var p = new(CausedByLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_causedByLine

	return p
}

func (s *CausedByLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CausedByLineContext) CAUSED_BY() antlr.TerminalNode {
	return s.GetToken(StackTraceParserCAUSED_BY, 0)
}

func (s *CausedByLineContext) StackTrace() IStackTraceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStackTraceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStackTraceContext)
}

func (s *CausedByLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CausedByLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CausedByLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterCausedByLine(s)
	}
}

func (s *CausedByLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitCausedByLine(s)
	}
}

func (p *StackTraceParser) CausedByLine() (localctx ICausedByLineContext) {
	localctx = NewCausedByLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StackTraceParserRULE_causedByLine)

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
		p.SetState(65)
		p.Match(StackTraceParserCAUSED_BY)
	}
	{
		p.SetState(66)
		p.StackTrace()
	}

	return localctx
}

// IEllipsisLineContext is an interface to support dynamic dispatch.
type IEllipsisLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEllipsisLineContext differentiates from other interfaces.
	IsEllipsisLineContext()
}

type EllipsisLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEllipsisLineContext() *EllipsisLineContext {
	var p = new(EllipsisLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_ellipsisLine
	return p
}

func (*EllipsisLineContext) IsEllipsisLineContext() {}

func NewEllipsisLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EllipsisLineContext {
	var p = new(EllipsisLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_ellipsisLine

	return p
}

func (s *EllipsisLineContext) GetParser() antlr.Parser { return s.parser }

func (s *EllipsisLineContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(StackTraceParserELLIPSIS, 0)
}

func (s *EllipsisLineContext) Number() antlr.TerminalNode {
	return s.GetToken(StackTraceParserNumber, 0)
}

func (s *EllipsisLineContext) MORE_() antlr.TerminalNode {
	return s.GetToken(StackTraceParserMORE_, 0)
}

func (s *EllipsisLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EllipsisLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EllipsisLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterEllipsisLine(s)
	}
}

func (s *EllipsisLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitEllipsisLine(s)
	}
}

func (p *StackTraceParser) EllipsisLine() (localctx IEllipsisLineContext) {
	localctx = NewEllipsisLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StackTraceParserRULE_ellipsisLine)

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
		p.Match(StackTraceParserELLIPSIS)
	}
	{
		p.SetState(69)
		p.Match(StackTraceParserNumber)
	}
	{
		p.SetState(70)
		p.Match(StackTraceParserMORE_)
	}

	return localctx
}

// IMessageLineContext is an interface to support dynamic dispatch.
type IMessageLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageLineContext differentiates from other interfaces.
	IsMessageLineContext()
}

type MessageLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageLineContext() *MessageLineContext {
	var p = new(MessageLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_messageLine
	return p
}

func (*MessageLineContext) IsMessageLineContext() {}

func NewMessageLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageLineContext {
	var p = new(MessageLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_messageLine

	return p
}

func (s *MessageLineContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageLineContext) QualifiedClass() IQualifiedClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedClassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedClassContext)
}

func (s *MessageLineContext) Message() IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterMessageLine(s)
	}
}

func (s *MessageLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitMessageLine(s)
	}
}

func (p *StackTraceParser) MessageLine() (localctx IMessageLineContext) {
	localctx = NewMessageLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StackTraceParserRULE_messageLine)
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
		p.SetState(72)
		p.QualifiedClass()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StackTraceParserCOLON {
		{
			p.SetState(73)
			p.Message()
		}

	}

	return localctx
}

// IQualifiedClassContext is an interface to support dynamic dispatch.
type IQualifiedClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedClassContext differentiates from other interfaces.
	IsQualifiedClassContext()
}

type QualifiedClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedClassContext() *QualifiedClassContext {
	var p = new(QualifiedClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_qualifiedClass
	return p
}

func (*QualifiedClassContext) IsQualifiedClassContext() {}

func NewQualifiedClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedClassContext {
	var p = new(QualifiedClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_qualifiedClass

	return p
}

func (s *QualifiedClassContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedClassContext) ClassName() IClassNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassNameContext)
}

func (s *QualifiedClassContext) PackagePath() IPackagePathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackagePathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackagePathContext)
}

func (s *QualifiedClassContext) AllInnerClassName() []IInnerClassNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInnerClassNameContext)(nil)).Elem())
	var tst = make([]IInnerClassNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInnerClassNameContext)
		}
	}

	return tst
}

func (s *QualifiedClassContext) InnerClassName(i int) IInnerClassNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInnerClassNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInnerClassNameContext)
}

func (s *QualifiedClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterQualifiedClass(s)
	}
}

func (s *QualifiedClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitQualifiedClass(s)
	}
}

func (p *StackTraceParser) QualifiedClass() (localctx IQualifiedClassContext) {
	localctx = NewQualifiedClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StackTraceParserRULE_qualifiedClass)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.PackagePath()
		}

	}
	{
		p.SetState(79)
		p.ClassName()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StackTraceParserT__2 {
		{
			p.SetState(80)
			p.InnerClassName()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInnerClassNameContext is an interface to support dynamic dispatch.
type IInnerClassNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerClassNameContext differentiates from other interfaces.
	IsInnerClassNameContext()
}

type InnerClassNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerClassNameContext() *InnerClassNameContext {
	var p = new(InnerClassNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_innerClassName
	return p
}

func (*InnerClassNameContext) IsInnerClassNameContext() {}

func NewInnerClassNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerClassNameContext {
	var p = new(InnerClassNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_innerClassName

	return p
}

func (s *InnerClassNameContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerClassNameContext) ClassName() IClassNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassNameContext)
}

func (s *InnerClassNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerClassNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerClassNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterInnerClassName(s)
	}
}

func (s *InnerClassNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitInnerClassName(s)
	}
}

func (p *StackTraceParser) InnerClassName() (localctx IInnerClassNameContext) {
	localctx = NewInnerClassNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StackTraceParserRULE_innerClassName)

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
		p.SetState(86)
		p.Match(StackTraceParserT__2)
	}
	{
		p.SetState(87)
		p.ClassName()
	}

	return localctx
}

// IClassFileContext is an interface to support dynamic dispatch.
type IClassFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassFileContext differentiates from other interfaces.
	IsClassFileContext()
}

type ClassFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassFileContext() *ClassFileContext {
	var p = new(ClassFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_classFile
	return p
}

func (*ClassFileContext) IsClassFileContext() {}

func NewClassFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassFileContext {
	var p = new(ClassFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_classFile

	return p
}

func (s *ClassFileContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassFileContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassFileContext) NATIVE_METHOD() antlr.TerminalNode {
	return s.GetToken(StackTraceParserNATIVE_METHOD, 0)
}

func (s *ClassFileContext) UNKNOWN_SOURCE() antlr.TerminalNode {
	return s.GetToken(StackTraceParserUNKNOWN_SOURCE, 0)
}

func (s *ClassFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterClassFile(s)
	}
}

func (s *ClassFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitClassFile(s)
	}
}

func (p *StackTraceParser) ClassFile() (localctx IClassFileContext) {
	localctx = NewClassFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, StackTraceParserRULE_classFile)

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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StackTraceParserJavaWord:
		{
			p.SetState(89)
			p.Identifier()
		}
		{
			p.SetState(90)
			p.Match(StackTraceParserT__3)
		}

	case StackTraceParserNATIVE_METHOD:
		{
			p.SetState(92)
			p.Match(StackTraceParserNATIVE_METHOD)
		}

	case StackTraceParserUNKNOWN_SOURCE:
		{
			p.SetState(93)
			p.Match(StackTraceParserUNKNOWN_SOURCE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualifiedMethodContext is an interface to support dynamic dispatch.
type IQualifiedMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedMethodContext differentiates from other interfaces.
	IsQualifiedMethodContext()
}

type QualifiedMethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedMethodContext() *QualifiedMethodContext {
	var p = new(QualifiedMethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_qualifiedMethod
	return p
}

func (*QualifiedMethodContext) IsQualifiedMethodContext() {}

func NewQualifiedMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedMethodContext {
	var p = new(QualifiedMethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_qualifiedMethod

	return p
}

func (s *QualifiedMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedMethodContext) QualifiedClass() IQualifiedClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedClassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedClassContext)
}

func (s *QualifiedMethodContext) DOT() antlr.TerminalNode {
	return s.GetToken(StackTraceParserDOT, 0)
}

func (s *QualifiedMethodContext) MethodName() IMethodNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodNameContext)
}

func (s *QualifiedMethodContext) Constructor() IConstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstructorContext)
}

func (s *QualifiedMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterQualifiedMethod(s)
	}
}

func (s *QualifiedMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitQualifiedMethod(s)
	}
}

func (p *StackTraceParser) QualifiedMethod() (localctx IQualifiedMethodContext) {
	localctx = NewQualifiedMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, StackTraceParserRULE_qualifiedMethod)

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
		p.SetState(96)
		p.QualifiedClass()
	}
	{
		p.SetState(97)
		p.Match(StackTraceParserDOT)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StackTraceParserJavaWord:
		{
			p.SetState(98)
			p.MethodName()
		}

	case StackTraceParserINIT:
		{
			p.SetState(99)
			p.Constructor()
		}

	case StackTraceParserT__0:

	default:
	}

	return localctx
}

// IConstructorContext is an interface to support dynamic dispatch.
type IConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructorContext differentiates from other interfaces.
	IsConstructorContext()
}

type ConstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructorContext() *ConstructorContext {
	var p = new(ConstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_constructor
	return p
}

func (*ConstructorContext) IsConstructorContext() {}

func NewConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructorContext {
	var p = new(ConstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_constructor

	return p
}

func (s *ConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructorContext) INIT() antlr.TerminalNode {
	return s.GetToken(StackTraceParserINIT, 0)
}

func (s *ConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterConstructor(s)
	}
}

func (s *ConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitConstructor(s)
	}
}

func (p *StackTraceParser) Constructor() (localctx IConstructorContext) {
	localctx = NewConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, StackTraceParserRULE_constructor)

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
		p.SetState(102)
		p.Match(StackTraceParserINIT)
	}

	return localctx
}

// IMethodNameContext is an interface to support dynamic dispatch.
type IMethodNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodNameContext differentiates from other interfaces.
	IsMethodNameContext()
}

type MethodNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodNameContext() *MethodNameContext {
	var p = new(MethodNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_methodName
	return p
}

func (*MethodNameContext) IsMethodNameContext() {}

func NewMethodNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodNameContext {
	var p = new(MethodNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_methodName

	return p
}

func (s *MethodNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodNameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MethodNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterMethodName(s)
	}
}

func (s *MethodNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitMethodName(s)
	}
}

func (p *StackTraceParser) MethodName() (localctx IMethodNameContext) {
	localctx = NewMethodNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StackTraceParserRULE_methodName)

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
		p.SetState(104)
		p.Identifier()
	}

	return localctx
}

// IPackagePathContext is an interface to support dynamic dispatch.
type IPackagePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackagePathContext differentiates from other interfaces.
	IsPackagePathContext()
}

type PackagePathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackagePathContext() *PackagePathContext {
	var p = new(PackagePathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_packagePath
	return p
}

func (*PackagePathContext) IsPackagePathContext() {}

func NewPackagePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackagePathContext {
	var p = new(PackagePathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_packagePath

	return p
}

func (s *PackagePathContext) GetParser() antlr.Parser { return s.parser }

func (s *PackagePathContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *PackagePathContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PackagePathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(StackTraceParserDOT)
}

func (s *PackagePathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(StackTraceParserDOT, i)
}

func (s *PackagePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackagePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackagePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterPackagePath(s)
	}
}

func (s *PackagePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitPackagePath(s)
	}
}

func (p *StackTraceParser) PackagePath() (localctx IPackagePathContext) {
	localctx = NewPackagePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, StackTraceParserRULE_packagePath)

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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(106)
				p.Identifier()
			}
			{
				p.SetState(107)
				p.Match(StackTraceParserDOT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IClassNameContext is an interface to support dynamic dispatch.
type IClassNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassNameContext differentiates from other interfaces.
	IsClassNameContext()
}

type ClassNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassNameContext() *ClassNameContext {
	var p = new(ClassNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_className
	return p
}

func (*ClassNameContext) IsClassNameContext() {}

func NewClassNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassNameContext {
	var p = new(ClassNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_className

	return p
}

func (s *ClassNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassNameContext) JavaWord() antlr.TerminalNode {
	return s.GetToken(StackTraceParserJavaWord, 0)
}

func (s *ClassNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterClassName(s)
	}
}

func (s *ClassNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitClassName(s)
	}
}

func (p *StackTraceParser) ClassName() (localctx IClassNameContext) {
	localctx = NewClassNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, StackTraceParserRULE_className)

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
		p.Match(StackTraceParserJavaWord)
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
	p.RuleIndex = StackTraceParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) JavaWord() antlr.TerminalNode {
	return s.GetToken(StackTraceParserJavaWord, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *StackTraceParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, StackTraceParserRULE_identifier)

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
		p.Match(StackTraceParserJavaWord)
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StackTraceParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StackTraceParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) COLON() antlr.TerminalNode {
	return s.GetToken(StackTraceParserCOLON, 0)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StackTraceListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *StackTraceParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, StackTraceParserRULE_message)

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
	{
		p.SetState(117)
		p.Match(StackTraceParserCOLON)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(118)
			p.MatchWildcard()

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}
