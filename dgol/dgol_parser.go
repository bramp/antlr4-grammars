// Code generated from dgol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgol // dgol
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 195,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 7, 2, 37, 10, 2, 12, 2, 14, 2, 40, 11, 2, 3, 2, 3, 2, 7, 2, 44, 10,
	2, 12, 2, 14, 2, 47, 11, 2, 3, 2, 3, 2, 5, 2, 51, 10, 2, 3, 2, 7, 2, 54,
	10, 2, 12, 2, 14, 2, 57, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 7, 4, 69, 10, 4, 12, 4, 14, 4, 72, 11, 4, 5, 4, 74,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 95, 10, 6, 12, 6,
	14, 6, 98, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	5, 8, 109, 10, 8, 3, 8, 7, 8, 112, 10, 8, 12, 8, 14, 8, 115, 11, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 123, 10, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 135, 10, 11, 3,
	12, 3, 12, 3, 12, 7, 12, 140, 10, 12, 12, 12, 14, 12, 143, 11, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 148, 10, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 164,
	10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 175, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 181, 10, 15, 12,
	15, 14, 15, 184, 11, 15, 5, 15, 186, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 4, 4, 2, 11, 11, 23, 23, 4, 2, 13, 13, 15, 15,
	2, 203, 2, 38, 3, 2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 82,
	3, 2, 2, 2, 10, 90, 3, 2, 2, 2, 12, 103, 3, 2, 2, 2, 14, 113, 3, 2, 2,
	2, 16, 122, 3, 2, 2, 2, 18, 124, 3, 2, 2, 2, 20, 126, 3, 2, 2, 2, 22, 136,
	3, 2, 2, 2, 24, 152, 3, 2, 2, 2, 26, 159, 3, 2, 2, 2, 28, 170, 3, 2, 2,
	2, 30, 189, 3, 2, 2, 2, 32, 191, 3, 2, 2, 2, 34, 37, 5, 4, 3, 2, 35, 37,
	7, 24, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2,
	38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 45, 3, 2, 2, 2, 40, 38, 3,
	2, 2, 2, 41, 44, 5, 6, 4, 2, 42, 44, 7, 24, 2, 2, 43, 41, 3, 2, 2, 2, 43,
	42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2,
	2, 46, 50, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 51, 5, 8, 5, 2, 49, 51,
	5, 10, 6, 2, 50, 48, 3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 55, 3, 2, 2, 2,
	52, 54, 7, 24, 2, 2, 53, 52, 3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3,
	2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 3, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58,
	59, 7, 3, 2, 2, 59, 60, 7, 23, 2, 2, 60, 61, 7, 24, 2, 2, 61, 5, 3, 2,
	2, 2, 62, 63, 7, 4, 2, 2, 63, 64, 7, 23, 2, 2, 64, 73, 7, 5, 2, 2, 65,
	70, 7, 23, 2, 2, 66, 67, 7, 6, 2, 2, 67, 69, 7, 23, 2, 2, 68, 66, 3, 2,
	2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 74,
	3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 65, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 76, 7, 7, 2, 2, 76, 77, 7, 24, 2, 2, 77, 78, 5,
	14, 8, 2, 78, 79, 7, 8, 2, 2, 79, 80, 7, 23, 2, 2, 80, 81, 7, 24, 2, 2,
	81, 7, 3, 2, 2, 2, 82, 83, 7, 9, 2, 2, 83, 84, 7, 23, 2, 2, 84, 85, 7,
	24, 2, 2, 85, 86, 5, 14, 8, 2, 86, 87, 7, 8, 2, 2, 87, 88, 7, 23, 2, 2,
	88, 89, 7, 24, 2, 2, 89, 9, 3, 2, 2, 2, 90, 91, 7, 10, 2, 2, 91, 96, 7,
	23, 2, 2, 92, 95, 5, 12, 7, 2, 93, 95, 7, 24, 2, 2, 94, 92, 3, 2, 2, 2,
	94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3,
	2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 100, 7, 8, 2, 2, 100,
	101, 7, 23, 2, 2, 101, 102, 7, 24, 2, 2, 102, 11, 3, 2, 2, 2, 103, 104,
	7, 4, 2, 2, 104, 105, 7, 23, 2, 2, 105, 106, 7, 24, 2, 2, 106, 13, 3, 2,
	2, 2, 107, 109, 5, 16, 9, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 112, 7, 24, 2, 2, 111, 108, 3, 2, 2, 2, 112,
	115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 15, 3,
	2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 123, 5, 20, 11, 2, 117, 123, 5, 22,
	12, 2, 118, 123, 5, 26, 14, 2, 119, 123, 5, 28, 15, 2, 120, 123, 5, 30,
	16, 2, 121, 123, 5, 32, 17, 2, 122, 116, 3, 2, 2, 2, 122, 117, 3, 2, 2,
	2, 122, 118, 3, 2, 2, 2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122,
	121, 3, 2, 2, 2, 123, 17, 3, 2, 2, 2, 124, 125, 9, 2, 2, 2, 125, 19, 3,
	2, 2, 2, 126, 127, 7, 12, 2, 2, 127, 134, 7, 23, 2, 2, 128, 129, 7, 13,
	2, 2, 129, 135, 5, 18, 10, 2, 130, 131, 7, 14, 2, 2, 131, 135, 7, 23, 2,
	2, 132, 133, 7, 15, 2, 2, 133, 135, 5, 18, 10, 2, 134, 128, 3, 2, 2, 2,
	134, 130, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 21, 3, 2, 2, 2, 136, 141,
	5, 24, 13, 2, 137, 138, 7, 16, 2, 2, 138, 140, 5, 24, 13, 2, 139, 137,
	3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 147, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 145, 7, 16, 2, 2,
	145, 146, 7, 24, 2, 2, 146, 148, 5, 14, 8, 2, 147, 144, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 7, 8, 2, 2, 150, 151,
	7, 17, 2, 2, 151, 23, 3, 2, 2, 2, 152, 153, 7, 17, 2, 2, 153, 154, 7, 23,
	2, 2, 154, 155, 9, 3, 2, 2, 155, 156, 7, 23, 2, 2, 156, 157, 7, 24, 2,
	2, 157, 158, 5, 14, 8, 2, 158, 25, 3, 2, 2, 2, 159, 160, 7, 18, 2, 2, 160,
	163, 7, 23, 2, 2, 161, 162, 7, 14, 2, 2, 162, 164, 7, 23, 2, 2, 163, 161,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 24,
	2, 2, 166, 167, 5, 14, 8, 2, 167, 168, 7, 8, 2, 2, 168, 169, 7, 18, 2,
	2, 169, 27, 3, 2, 2, 2, 170, 171, 7, 19, 2, 2, 171, 174, 7, 23, 2, 2, 172,
	173, 7, 20, 2, 2, 173, 175, 7, 23, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175,
	3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 185, 7, 5, 2, 2, 177, 182, 5, 18,
	10, 2, 178, 179, 7, 6, 2, 2, 179, 181, 5, 18, 10, 2, 180, 178, 3, 2, 2,
	2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 177, 3, 2, 2, 2, 185, 186,
	3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188, 7, 7, 2, 2, 188, 29, 3, 2,
	2, 2, 189, 190, 7, 21, 2, 2, 190, 31, 3, 2, 2, 2, 191, 192, 7, 22, 2, 2,
	192, 193, 7, 23, 2, 2, 193, 33, 3, 2, 2, 2, 22, 36, 38, 43, 45, 50, 55,
	70, 73, 94, 96, 108, 113, 122, 134, 141, 147, 163, 174, 182, 185,
}
var literalNames = []string{
	"", "'USE'", "'SUBROUTINE'", "'('", "','", "')'", "'END'", "'PROGRAM'",
	"'LIBRARY'", "'0'", "'LET'", "'='", "'<'", "'>'", "'ELSE'", "'IF'", "'DO'",
	"'CALL'", "'.'", "'RETURN'", "'EXIT'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "IDENTIFER", "NL", "WS",
}

var ruleNames = []string{
	"module", "usedeclaration", "subroutinedefinition", "programdefinition",
	"librarydefinition", "subroutinedeclaration", "statements", "statement",
	"identifierorzero", "letstatement", "ifstatement", "ifhead", "dostatement",
	"callstatement", "returnstatement", "exitstatement",
}

type dgolParser struct {
	*antlr.BaseParser
}

// NewdgolParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *dgolParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdgolParser(input antlr.TokenStream) *dgolParser {
	this := new(dgolParser)
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
	this.GrammarFileName = "dgol.g4"

	return this
}

// dgolParser tokens.
const (
	dgolParserEOF       = antlr.TokenEOF
	dgolParserT__0      = 1
	dgolParserT__1      = 2
	dgolParserT__2      = 3
	dgolParserT__3      = 4
	dgolParserT__4      = 5
	dgolParserT__5      = 6
	dgolParserT__6      = 7
	dgolParserT__7      = 8
	dgolParserT__8      = 9
	dgolParserT__9      = 10
	dgolParserT__10     = 11
	dgolParserT__11     = 12
	dgolParserT__12     = 13
	dgolParserT__13     = 14
	dgolParserT__14     = 15
	dgolParserT__15     = 16
	dgolParserT__16     = 17
	dgolParserT__17     = 18
	dgolParserT__18     = 19
	dgolParserT__19     = 20
	dgolParserIDENTIFER = 21
	dgolParserNL        = 22
	dgolParserWS        = 23
)

// dgolParser rules.
const (
	dgolParserRULE_module                = 0
	dgolParserRULE_usedeclaration        = 1
	dgolParserRULE_subroutinedefinition  = 2
	dgolParserRULE_programdefinition     = 3
	dgolParserRULE_librarydefinition     = 4
	dgolParserRULE_subroutinedeclaration = 5
	dgolParserRULE_statements            = 6
	dgolParserRULE_statement             = 7
	dgolParserRULE_identifierorzero      = 8
	dgolParserRULE_letstatement          = 9
	dgolParserRULE_ifstatement           = 10
	dgolParserRULE_ifhead                = 11
	dgolParserRULE_dostatement           = 12
	dgolParserRULE_callstatement         = 13
	dgolParserRULE_returnstatement       = 14
	dgolParserRULE_exitstatement         = 15
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) Programdefinition() IProgramdefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramdefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramdefinitionContext)
}

func (s *ModuleContext) Librarydefinition() ILibrarydefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibrarydefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibrarydefinitionContext)
}

func (s *ModuleContext) AllUsedeclaration() []IUsedeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUsedeclarationContext)(nil)).Elem())
	var tst = make([]IUsedeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUsedeclarationContext)
		}
	}

	return tst
}

func (s *ModuleContext) Usedeclaration(i int) IUsedeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsedeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUsedeclarationContext)
}

func (s *ModuleContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(dgolParserNL)
}

func (s *ModuleContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserNL, i)
}

func (s *ModuleContext) AllSubroutinedefinition() []ISubroutinedefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubroutinedefinitionContext)(nil)).Elem())
	var tst = make([]ISubroutinedefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubroutinedefinitionContext)
		}
	}

	return tst
}

func (s *ModuleContext) Subroutinedefinition(i int) ISubroutinedefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinedefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubroutinedefinitionContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *dgolParser) Module() (localctx IModuleContext) {
	this := p
	_ = this

	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, dgolParserRULE_module)
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
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(34)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case dgolParserT__0:
				{
					p.SetState(32)
					p.Usedeclaration()
				}

			case dgolParserNL:
				{
					p.SetState(33)
					p.Match(dgolParserNL)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dgolParserT__1 || _la == dgolParserNL {
		p.SetState(41)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case dgolParserT__1:
			{
				p.SetState(39)
				p.Subroutinedefinition()
			}

		case dgolParserNL:
			{
				p.SetState(40)
				p.Match(dgolParserNL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dgolParserT__6:
		{
			p.SetState(46)
			p.Programdefinition()
		}

	case dgolParserT__7:
		{
			p.SetState(47)
			p.Librarydefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dgolParserNL {
		{
			p.SetState(50)
			p.Match(dgolParserNL)
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUsedeclarationContext is an interface to support dynamic dispatch.
type IUsedeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUsedeclarationContext differentiates from other interfaces.
	IsUsedeclarationContext()
}

type UsedeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsedeclarationContext() *UsedeclarationContext {
	var p = new(UsedeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_usedeclaration
	return p
}

func (*UsedeclarationContext) IsUsedeclarationContext() {}

func NewUsedeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UsedeclarationContext {
	var p = new(UsedeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_usedeclaration

	return p
}

func (s *UsedeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *UsedeclarationContext) IDENTIFER() antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, 0)
}

func (s *UsedeclarationContext) NL() antlr.TerminalNode {
	return s.GetToken(dgolParserNL, 0)
}

func (s *UsedeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UsedeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UsedeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterUsedeclaration(s)
	}
}

func (s *UsedeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitUsedeclaration(s)
	}
}

func (p *dgolParser) Usedeclaration() (localctx IUsedeclarationContext) {
	this := p
	_ = this

	localctx = NewUsedeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, dgolParserRULE_usedeclaration)

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
		p.SetState(56)
		p.Match(dgolParserT__0)
	}
	{
		p.SetState(57)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(58)
		p.Match(dgolParserNL)
	}

	return localctx
}

// ISubroutinedefinitionContext is an interface to support dynamic dispatch.
type ISubroutinedefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinedefinitionContext differentiates from other interfaces.
	IsSubroutinedefinitionContext()
}

type SubroutinedefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinedefinitionContext() *SubroutinedefinitionContext {
	var p = new(SubroutinedefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_subroutinedefinition
	return p
}

func (*SubroutinedefinitionContext) IsSubroutinedefinitionContext() {}

func NewSubroutinedefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinedefinitionContext {
	var p = new(SubroutinedefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_subroutinedefinition

	return p
}

func (s *SubroutinedefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinedefinitionContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *SubroutinedefinitionContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *SubroutinedefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(dgolParserNL)
}

func (s *SubroutinedefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserNL, i)
}

func (s *SubroutinedefinitionContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *SubroutinedefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinedefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinedefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterSubroutinedefinition(s)
	}
}

func (s *SubroutinedefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitSubroutinedefinition(s)
	}
}

func (p *dgolParser) Subroutinedefinition() (localctx ISubroutinedefinitionContext) {
	this := p
	_ = this

	localctx = NewSubroutinedefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, dgolParserRULE_subroutinedefinition)
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
		p.SetState(60)
		p.Match(dgolParserT__1)
	}
	{
		p.SetState(61)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(62)
		p.Match(dgolParserT__2)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dgolParserIDENTIFER {
		{
			p.SetState(63)
			p.Match(dgolParserIDENTIFER)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == dgolParserT__3 {
			{
				p.SetState(64)
				p.Match(dgolParserT__3)
			}
			{
				p.SetState(65)
				p.Match(dgolParserIDENTIFER)
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(73)
		p.Match(dgolParserT__4)
	}
	{
		p.SetState(74)
		p.Match(dgolParserNL)
	}
	{
		p.SetState(75)
		p.Statements()
	}
	{
		p.SetState(76)
		p.Match(dgolParserT__5)
	}
	{
		p.SetState(77)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(78)
		p.Match(dgolParserNL)
	}

	return localctx
}

// IProgramdefinitionContext is an interface to support dynamic dispatch.
type IProgramdefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramdefinitionContext differentiates from other interfaces.
	IsProgramdefinitionContext()
}

type ProgramdefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramdefinitionContext() *ProgramdefinitionContext {
	var p = new(ProgramdefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_programdefinition
	return p
}

func (*ProgramdefinitionContext) IsProgramdefinitionContext() {}

func NewProgramdefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramdefinitionContext {
	var p = new(ProgramdefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_programdefinition

	return p
}

func (s *ProgramdefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramdefinitionContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *ProgramdefinitionContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *ProgramdefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(dgolParserNL)
}

func (s *ProgramdefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserNL, i)
}

func (s *ProgramdefinitionContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramdefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramdefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramdefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterProgramdefinition(s)
	}
}

func (s *ProgramdefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitProgramdefinition(s)
	}
}

func (p *dgolParser) Programdefinition() (localctx IProgramdefinitionContext) {
	this := p
	_ = this

	localctx = NewProgramdefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, dgolParserRULE_programdefinition)

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
		p.SetState(80)
		p.Match(dgolParserT__6)
	}
	{
		p.SetState(81)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(82)
		p.Match(dgolParserNL)
	}
	{
		p.SetState(83)
		p.Statements()
	}
	{
		p.SetState(84)
		p.Match(dgolParserT__5)
	}
	{
		p.SetState(85)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(86)
		p.Match(dgolParserNL)
	}

	return localctx
}

// ILibrarydefinitionContext is an interface to support dynamic dispatch.
type ILibrarydefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibrarydefinitionContext differentiates from other interfaces.
	IsLibrarydefinitionContext()
}

type LibrarydefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibrarydefinitionContext() *LibrarydefinitionContext {
	var p = new(LibrarydefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_librarydefinition
	return p
}

func (*LibrarydefinitionContext) IsLibrarydefinitionContext() {}

func NewLibrarydefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibrarydefinitionContext {
	var p = new(LibrarydefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_librarydefinition

	return p
}

func (s *LibrarydefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *LibrarydefinitionContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *LibrarydefinitionContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *LibrarydefinitionContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(dgolParserNL)
}

func (s *LibrarydefinitionContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserNL, i)
}

func (s *LibrarydefinitionContext) AllSubroutinedeclaration() []ISubroutinedeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubroutinedeclarationContext)(nil)).Elem())
	var tst = make([]ISubroutinedeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubroutinedeclarationContext)
		}
	}

	return tst
}

func (s *LibrarydefinitionContext) Subroutinedeclaration(i int) ISubroutinedeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubroutinedeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubroutinedeclarationContext)
}

func (s *LibrarydefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibrarydefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibrarydefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterLibrarydefinition(s)
	}
}

func (s *LibrarydefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitLibrarydefinition(s)
	}
}

func (p *dgolParser) Librarydefinition() (localctx ILibrarydefinitionContext) {
	this := p
	_ = this

	localctx = NewLibrarydefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, dgolParserRULE_librarydefinition)
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
		p.SetState(88)
		p.Match(dgolParserT__7)
	}
	{
		p.SetState(89)
		p.Match(dgolParserIDENTIFER)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dgolParserT__1 || _la == dgolParserNL {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case dgolParserT__1:
			{
				p.SetState(90)
				p.Subroutinedeclaration()
			}

		case dgolParserNL:
			{
				p.SetState(91)
				p.Match(dgolParserNL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(dgolParserT__5)
	}
	{
		p.SetState(98)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(99)
		p.Match(dgolParserNL)
	}

	return localctx
}

// ISubroutinedeclarationContext is an interface to support dynamic dispatch.
type ISubroutinedeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubroutinedeclarationContext differentiates from other interfaces.
	IsSubroutinedeclarationContext()
}

type SubroutinedeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubroutinedeclarationContext() *SubroutinedeclarationContext {
	var p = new(SubroutinedeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_subroutinedeclaration
	return p
}

func (*SubroutinedeclarationContext) IsSubroutinedeclarationContext() {}

func NewSubroutinedeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubroutinedeclarationContext {
	var p = new(SubroutinedeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_subroutinedeclaration

	return p
}

func (s *SubroutinedeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SubroutinedeclarationContext) IDENTIFER() antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, 0)
}

func (s *SubroutinedeclarationContext) NL() antlr.TerminalNode {
	return s.GetToken(dgolParserNL, 0)
}

func (s *SubroutinedeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubroutinedeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubroutinedeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterSubroutinedeclaration(s)
	}
}

func (s *SubroutinedeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitSubroutinedeclaration(s)
	}
}

func (p *dgolParser) Subroutinedeclaration() (localctx ISubroutinedeclarationContext) {
	this := p
	_ = this

	localctx = NewSubroutinedeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, dgolParserRULE_subroutinedeclaration)

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
		p.Match(dgolParserT__1)
	}
	{
		p.SetState(102)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(103)
		p.Match(dgolParserNL)
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
	p.RuleIndex = dgolParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(dgolParserNL)
}

func (s *StatementsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserNL, i)
}

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
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (p *dgolParser) Statements() (localctx IStatementsContext) {
	this := p
	_ = this

	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, dgolParserRULE_statements)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<dgolParserT__9)|(1<<dgolParserT__14)|(1<<dgolParserT__15)|(1<<dgolParserT__16)|(1<<dgolParserT__18)|(1<<dgolParserT__19)|(1<<dgolParserNL))) != 0 {
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<dgolParserT__9)|(1<<dgolParserT__14)|(1<<dgolParserT__15)|(1<<dgolParserT__16)|(1<<dgolParserT__18)|(1<<dgolParserT__19))) != 0 {
			{
				p.SetState(105)
				p.Statement()
			}

		}
		{
			p.SetState(108)
			p.Match(dgolParserNL)
		}

		p.SetState(113)
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
	p.RuleIndex = dgolParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Letstatement() ILetstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILetstatementContext)
}

func (s *StatementContext) Ifstatement() IIfstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatementContext)
}

func (s *StatementContext) Dostatement() IDostatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDostatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDostatementContext)
}

func (s *StatementContext) Callstatement() ICallstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallstatementContext)
}

func (s *StatementContext) Returnstatement() IReturnstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnstatementContext)
}

func (s *StatementContext) Exitstatement() IExitstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExitstatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExitstatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *dgolParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, dgolParserRULE_statement)

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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dgolParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Letstatement()
		}

	case dgolParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Ifstatement()
		}

	case dgolParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Dostatement()
		}

	case dgolParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Callstatement()
		}

	case dgolParserT__18:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.Returnstatement()
		}

	case dgolParserT__19:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.Exitstatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierorzeroContext is an interface to support dynamic dispatch.
type IIdentifierorzeroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierorzeroContext differentiates from other interfaces.
	IsIdentifierorzeroContext()
}

type IdentifierorzeroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierorzeroContext() *IdentifierorzeroContext {
	var p = new(IdentifierorzeroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_identifierorzero
	return p
}

func (*IdentifierorzeroContext) IsIdentifierorzeroContext() {}

func NewIdentifierorzeroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierorzeroContext {
	var p = new(IdentifierorzeroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_identifierorzero

	return p
}

func (s *IdentifierorzeroContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierorzeroContext) IDENTIFER() antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, 0)
}

func (s *IdentifierorzeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierorzeroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierorzeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterIdentifierorzero(s)
	}
}

func (s *IdentifierorzeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitIdentifierorzero(s)
	}
}

func (p *dgolParser) Identifierorzero() (localctx IIdentifierorzeroContext) {
	this := p
	_ = this

	localctx = NewIdentifierorzeroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, dgolParserRULE_identifierorzero)
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
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == dgolParserT__8 || _la == dgolParserIDENTIFER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILetstatementContext is an interface to support dynamic dispatch.
type ILetstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetstatementContext differentiates from other interfaces.
	IsLetstatementContext()
}

type LetstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetstatementContext() *LetstatementContext {
	var p = new(LetstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_letstatement
	return p
}

func (*LetstatementContext) IsLetstatementContext() {}

func NewLetstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetstatementContext {
	var p = new(LetstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_letstatement

	return p
}

func (s *LetstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LetstatementContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *LetstatementContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *LetstatementContext) Identifierorzero() IIdentifierorzeroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierorzeroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierorzeroContext)
}

func (s *LetstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterLetstatement(s)
	}
}

func (s *LetstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitLetstatement(s)
	}
}

func (p *dgolParser) Letstatement() (localctx ILetstatementContext) {
	this := p
	_ = this

	localctx = NewLetstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, dgolParserRULE_letstatement)

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
		p.SetState(124)
		p.Match(dgolParserT__9)
	}
	{
		p.SetState(125)
		p.Match(dgolParserIDENTIFER)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dgolParserT__10:
		{
			p.SetState(126)
			p.Match(dgolParserT__10)
		}
		{
			p.SetState(127)
			p.Identifierorzero()
		}

	case dgolParserT__11:
		{
			p.SetState(128)
			p.Match(dgolParserT__11)
		}
		{
			p.SetState(129)
			p.Match(dgolParserIDENTIFER)
		}

	case dgolParserT__12:
		{
			p.SetState(130)
			p.Match(dgolParserT__12)
		}
		{
			p.SetState(131)
			p.Identifierorzero()
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
	p.RuleIndex = dgolParserRULE_ifstatement
	return p
}

func (*IfstatementContext) IsIfstatementContext() {}

func NewIfstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatementContext {
	var p = new(IfstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_ifstatement

	return p
}

func (s *IfstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatementContext) AllIfhead() []IIfheadContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfheadContext)(nil)).Elem())
	var tst = make([]IIfheadContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfheadContext)
		}
	}

	return tst
}

func (s *IfstatementContext) Ifhead(i int) IIfheadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfheadContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfheadContext)
}

func (s *IfstatementContext) NL() antlr.TerminalNode {
	return s.GetToken(dgolParserNL, 0)
}

func (s *IfstatementContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterIfstatement(s)
	}
}

func (s *IfstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitIfstatement(s)
	}
}

func (p *dgolParser) Ifstatement() (localctx IIfstatementContext) {
	this := p
	_ = this

	localctx = NewIfstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, dgolParserRULE_ifstatement)
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
	{
		p.SetState(134)
		p.Ifhead()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(135)
				p.Match(dgolParserT__13)
			}
			{
				p.SetState(136)
				p.Ifhead()
			}

		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dgolParserT__13 {
		{
			p.SetState(142)
			p.Match(dgolParserT__13)
		}
		{
			p.SetState(143)
			p.Match(dgolParserNL)
		}
		{
			p.SetState(144)
			p.Statements()
		}

	}
	{
		p.SetState(147)
		p.Match(dgolParserT__5)
	}
	{
		p.SetState(148)
		p.Match(dgolParserT__14)
	}

	return localctx
}

// IIfheadContext is an interface to support dynamic dispatch.
type IIfheadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfheadContext differentiates from other interfaces.
	IsIfheadContext()
}

type IfheadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfheadContext() *IfheadContext {
	var p = new(IfheadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_ifhead
	return p
}

func (*IfheadContext) IsIfheadContext() {}

func NewIfheadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfheadContext {
	var p = new(IfheadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_ifhead

	return p
}

func (s *IfheadContext) GetParser() antlr.Parser { return s.parser }

func (s *IfheadContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *IfheadContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *IfheadContext) NL() antlr.TerminalNode {
	return s.GetToken(dgolParserNL, 0)
}

func (s *IfheadContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfheadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfheadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfheadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterIfhead(s)
	}
}

func (s *IfheadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitIfhead(s)
	}
}

func (p *dgolParser) Ifhead() (localctx IIfheadContext) {
	this := p
	_ = this

	localctx = NewIfheadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, dgolParserRULE_ifhead)
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
		p.SetState(150)
		p.Match(dgolParserT__14)
	}
	{
		p.SetState(151)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(152)
		_la = p.GetTokenStream().LA(1)

		if !(_la == dgolParserT__10 || _la == dgolParserT__12) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(153)
		p.Match(dgolParserIDENTIFER)
	}
	{
		p.SetState(154)
		p.Match(dgolParserNL)
	}
	{
		p.SetState(155)
		p.Statements()
	}

	return localctx
}

// IDostatementContext is an interface to support dynamic dispatch.
type IDostatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDostatementContext differentiates from other interfaces.
	IsDostatementContext()
}

type DostatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDostatementContext() *DostatementContext {
	var p = new(DostatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_dostatement
	return p
}

func (*DostatementContext) IsDostatementContext() {}

func NewDostatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DostatementContext {
	var p = new(DostatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_dostatement

	return p
}

func (s *DostatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DostatementContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *DostatementContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *DostatementContext) NL() antlr.TerminalNode {
	return s.GetToken(dgolParserNL, 0)
}

func (s *DostatementContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *DostatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DostatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DostatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterDostatement(s)
	}
}

func (s *DostatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitDostatement(s)
	}
}

func (p *dgolParser) Dostatement() (localctx IDostatementContext) {
	this := p
	_ = this

	localctx = NewDostatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, dgolParserRULE_dostatement)
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
		p.SetState(157)
		p.Match(dgolParserT__15)
	}
	{
		p.SetState(158)
		p.Match(dgolParserIDENTIFER)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dgolParserT__11 {
		{
			p.SetState(159)
			p.Match(dgolParserT__11)
		}
		{
			p.SetState(160)
			p.Match(dgolParserIDENTIFER)
		}

	}
	{
		p.SetState(163)
		p.Match(dgolParserNL)
	}
	{
		p.SetState(164)
		p.Statements()
	}
	{
		p.SetState(165)
		p.Match(dgolParserT__5)
	}
	{
		p.SetState(166)
		p.Match(dgolParserT__15)
	}

	return localctx
}

// ICallstatementContext is an interface to support dynamic dispatch.
type ICallstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallstatementContext differentiates from other interfaces.
	IsCallstatementContext()
}

type CallstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallstatementContext() *CallstatementContext {
	var p = new(CallstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_callstatement
	return p
}

func (*CallstatementContext) IsCallstatementContext() {}

func NewCallstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallstatementContext {
	var p = new(CallstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_callstatement

	return p
}

func (s *CallstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CallstatementContext) AllIDENTIFER() []antlr.TerminalNode {
	return s.GetTokens(dgolParserIDENTIFER)
}

func (s *CallstatementContext) IDENTIFER(i int) antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, i)
}

func (s *CallstatementContext) AllIdentifierorzero() []IIdentifierorzeroContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierorzeroContext)(nil)).Elem())
	var tst = make([]IIdentifierorzeroContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierorzeroContext)
		}
	}

	return tst
}

func (s *CallstatementContext) Identifierorzero(i int) IIdentifierorzeroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierorzeroContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierorzeroContext)
}

func (s *CallstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterCallstatement(s)
	}
}

func (s *CallstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitCallstatement(s)
	}
}

func (p *dgolParser) Callstatement() (localctx ICallstatementContext) {
	this := p
	_ = this

	localctx = NewCallstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, dgolParserRULE_callstatement)
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
		p.SetState(168)
		p.Match(dgolParserT__16)
	}
	{
		p.SetState(169)
		p.Match(dgolParserIDENTIFER)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dgolParserT__17 {
		{
			p.SetState(170)
			p.Match(dgolParserT__17)
		}
		{
			p.SetState(171)
			p.Match(dgolParserIDENTIFER)
		}

	}
	{
		p.SetState(174)
		p.Match(dgolParserT__2)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dgolParserT__8 || _la == dgolParserIDENTIFER {
		{
			p.SetState(175)
			p.Identifierorzero()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == dgolParserT__3 {
			{
				p.SetState(176)
				p.Match(dgolParserT__3)
			}
			{
				p.SetState(177)
				p.Identifierorzero()
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(185)
		p.Match(dgolParserT__4)
	}

	return localctx
}

// IReturnstatementContext is an interface to support dynamic dispatch.
type IReturnstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnstatementContext differentiates from other interfaces.
	IsReturnstatementContext()
}

type ReturnstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstatementContext() *ReturnstatementContext {
	var p = new(ReturnstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_returnstatement
	return p
}

func (*ReturnstatementContext) IsReturnstatementContext() {}

func NewReturnstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstatementContext {
	var p = new(ReturnstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_returnstatement

	return p
}

func (s *ReturnstatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ReturnstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterReturnstatement(s)
	}
}

func (s *ReturnstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitReturnstatement(s)
	}
}

func (p *dgolParser) Returnstatement() (localctx IReturnstatementContext) {
	this := p
	_ = this

	localctx = NewReturnstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, dgolParserRULE_returnstatement)

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
		p.SetState(187)
		p.Match(dgolParserT__18)
	}

	return localctx
}

// IExitstatementContext is an interface to support dynamic dispatch.
type IExitstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExitstatementContext differentiates from other interfaces.
	IsExitstatementContext()
}

type ExitstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExitstatementContext() *ExitstatementContext {
	var p = new(ExitstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dgolParserRULE_exitstatement
	return p
}

func (*ExitstatementContext) IsExitstatementContext() {}

func NewExitstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExitstatementContext {
	var p = new(ExitstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dgolParserRULE_exitstatement

	return p
}

func (s *ExitstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExitstatementContext) IDENTIFER() antlr.TerminalNode {
	return s.GetToken(dgolParserIDENTIFER, 0)
}

func (s *ExitstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExitstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExitstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.EnterExitstatement(s)
	}
}

func (s *ExitstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dgolListener); ok {
		listenerT.ExitExitstatement(s)
	}
}

func (p *dgolParser) Exitstatement() (localctx IExitstatementContext) {
	this := p
	_ = this

	localctx = NewExitstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, dgolParserRULE_exitstatement)

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
		p.SetState(189)
		p.Match(dgolParserT__19)
	}
	{
		p.SetState(190)
		p.Match(dgolParserIDENTIFER)
	}

	return localctx
}
