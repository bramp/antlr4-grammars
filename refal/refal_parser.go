// Code generated from refal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package refal // refal
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 170,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 5, 2, 50, 10, 2, 3, 2, 5, 2, 53, 10, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 59, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 65, 10, 2, 12,
	2, 14, 2, 68, 11, 2, 3, 3, 5, 3, 71, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 86, 10, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 93, 10, 7, 5, 7, 95, 10, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 103, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 113, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 124, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 5, 15, 132, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 145, 10, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 5, 19, 152, 10, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 5, 23, 168, 10, 23, 3, 23, 2, 3, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 4, 3, 2, 7,
	9, 3, 2, 20, 21, 2, 167, 2, 58, 3, 2, 2, 2, 4, 70, 3, 2, 2, 2, 6, 77, 3,
	2, 2, 2, 8, 80, 3, 2, 2, 2, 10, 87, 3, 2, 2, 2, 12, 89, 3, 2, 2, 2, 14,
	96, 3, 2, 2, 2, 16, 104, 3, 2, 2, 2, 18, 112, 3, 2, 2, 2, 20, 114, 3, 2,
	2, 2, 22, 116, 3, 2, 2, 2, 24, 118, 3, 2, 2, 2, 26, 123, 3, 2, 2, 2, 28,
	131, 3, 2, 2, 2, 30, 133, 3, 2, 2, 2, 32, 144, 3, 2, 2, 2, 34, 146, 3,
	2, 2, 2, 36, 151, 3, 2, 2, 2, 38, 153, 3, 2, 2, 2, 40, 157, 3, 2, 2, 2,
	42, 161, 3, 2, 2, 2, 44, 167, 3, 2, 2, 2, 46, 47, 8, 2, 1, 2, 47, 52, 5,
	4, 3, 2, 48, 50, 7, 3, 2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 53, 5, 2, 2, 2, 52, 49, 3, 2, 2, 2, 52, 53, 3, 2, 2,
	2, 53, 59, 3, 2, 2, 2, 54, 55, 5, 6, 4, 2, 55, 56, 7, 3, 2, 2, 56, 57,
	5, 2, 2, 4, 57, 59, 3, 2, 2, 2, 58, 46, 3, 2, 2, 2, 58, 54, 3, 2, 2, 2,
	59, 66, 3, 2, 2, 2, 60, 61, 12, 3, 2, 2, 61, 62, 5, 6, 4, 2, 62, 63, 7,
	3, 2, 2, 63, 65, 3, 2, 2, 2, 64, 60, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66,
	64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 3, 3, 2, 2, 2, 68, 66, 3, 2, 2,
	2, 69, 71, 7, 4, 2, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 73, 5, 10, 6, 2, 73, 74, 7, 5, 2, 2, 74, 75, 5, 12, 7,
	2, 75, 76, 7, 6, 2, 2, 76, 5, 3, 2, 2, 2, 77, 78, 9, 2, 2, 2, 78, 79, 5,
	8, 5, 2, 79, 7, 3, 2, 2, 2, 80, 85, 5, 10, 6, 2, 81, 82, 7, 10, 2, 2, 82,
	83, 5, 8, 5, 2, 83, 84, 7, 3, 2, 2, 84, 86, 3, 2, 2, 2, 85, 81, 3, 2, 2,
	2, 85, 86, 3, 2, 2, 2, 86, 9, 3, 2, 2, 2, 87, 88, 5, 34, 18, 2, 88, 11,
	3, 2, 2, 2, 89, 94, 5, 14, 8, 2, 90, 92, 7, 3, 2, 2, 91, 93, 5, 12, 7,
	2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 90,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 13, 3, 2, 2, 2, 96, 97, 5, 16, 9, 2,
	97, 102, 5, 18, 10, 2, 98, 99, 7, 11, 2, 2, 99, 103, 5, 24, 13, 2, 100,
	101, 7, 10, 2, 2, 101, 103, 5, 30, 16, 2, 102, 98, 3, 2, 2, 2, 102, 100,
	3, 2, 2, 2, 103, 15, 3, 2, 2, 2, 104, 105, 5, 20, 11, 2, 105, 17, 3, 2,
	2, 2, 106, 107, 7, 10, 2, 2, 107, 108, 5, 22, 12, 2, 108, 109, 7, 12, 2,
	2, 109, 110, 5, 20, 11, 2, 110, 111, 5, 18, 10, 2, 111, 113, 3, 2, 2, 2,
	112, 106, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 19, 3, 2, 2, 2, 114, 115,
	5, 26, 14, 2, 115, 21, 3, 2, 2, 2, 116, 117, 5, 26, 14, 2, 117, 23, 3,
	2, 2, 2, 118, 119, 5, 26, 14, 2, 119, 25, 3, 2, 2, 2, 120, 121, 5, 28,
	15, 2, 121, 122, 5, 26, 14, 2, 122, 124, 3, 2, 2, 2, 123, 120, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 27, 3, 2, 2, 2, 125, 132, 5, 32, 17, 2, 126,
	132, 5, 36, 19, 2, 127, 128, 7, 13, 2, 2, 128, 129, 5, 26, 14, 2, 129,
	130, 7, 14, 2, 2, 130, 132, 3, 2, 2, 2, 131, 125, 3, 2, 2, 2, 131, 126,
	3, 2, 2, 2, 131, 127, 3, 2, 2, 2, 132, 29, 3, 2, 2, 2, 133, 134, 5, 22,
	12, 2, 134, 135, 7, 12, 2, 2, 135, 136, 7, 5, 2, 2, 136, 137, 5, 12, 7,
	2, 137, 138, 7, 6, 2, 2, 138, 31, 3, 2, 2, 2, 139, 145, 5, 34, 18, 2, 140,
	145, 7, 19, 2, 2, 141, 145, 7, 21, 2, 2, 142, 145, 7, 22, 2, 2, 143, 145,
	7, 23, 2, 2, 144, 139, 3, 2, 2, 2, 144, 140, 3, 2, 2, 2, 144, 141, 3, 2,
	2, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 33, 3, 2, 2, 2,
	146, 147, 9, 3, 2, 2, 147, 35, 3, 2, 2, 2, 148, 152, 5, 38, 20, 2, 149,
	152, 5, 40, 21, 2, 150, 152, 5, 42, 22, 2, 151, 148, 3, 2, 2, 2, 151, 149,
	3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 37, 3, 2, 2, 2, 153, 154, 7, 15,
	2, 2, 154, 155, 7, 16, 2, 2, 155, 156, 5, 44, 23, 2, 156, 39, 3, 2, 2,
	2, 157, 158, 7, 17, 2, 2, 158, 159, 7, 16, 2, 2, 159, 160, 5, 44, 23, 2,
	160, 41, 3, 2, 2, 2, 161, 162, 7, 18, 2, 2, 162, 163, 7, 16, 2, 2, 163,
	164, 5, 44, 23, 2, 164, 43, 3, 2, 2, 2, 165, 168, 5, 34, 18, 2, 166, 168,
	7, 19, 2, 2, 167, 165, 3, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 45, 3, 2,
	2, 2, 17, 49, 52, 58, 66, 70, 85, 92, 94, 102, 112, 123, 131, 144, 151,
	167,
}
var literalNames = []string{
	"", "';'", "'$ENTRY'", "'{'", "'}'", "'$EXTERNAL'", "'$EXTERN'", "'$EXTRN'",
	"','", "'='", "':'", "'<'", "'>'", "'s'", "'.'", "'t'", "'e'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "DIGITS",
	"IDENTIFER", "STRING", "STRING2", "CHAR", "LINE_COMMENT", "BLOCK_COMMENT",
	"WS",
}

var ruleNames = []string{
	"program", "f_definition", "external_decl", "f_name_list", "f_name", "block_",
	"sentence", "left_side", "conditions", "pattern", "arg_", "right_side",
	"expression_", "term_", "block_ending", "symbol", "identifier", "variable",
	"svar", "tvar", "evar", "index",
}

type refalParser struct {
	*antlr.BaseParser
}

// NewrefalParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *refalParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewrefalParser(input antlr.TokenStream) *refalParser {
	this := new(refalParser)
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
	this.GrammarFileName = "refal.g4"

	return this
}

// refalParser tokens.
const (
	refalParserEOF           = antlr.TokenEOF
	refalParserT__0          = 1
	refalParserT__1          = 2
	refalParserT__2          = 3
	refalParserT__3          = 4
	refalParserT__4          = 5
	refalParserT__5          = 6
	refalParserT__6          = 7
	refalParserT__7          = 8
	refalParserT__8          = 9
	refalParserT__9          = 10
	refalParserT__10         = 11
	refalParserT__11         = 12
	refalParserT__12         = 13
	refalParserT__13         = 14
	refalParserT__14         = 15
	refalParserT__15         = 16
	refalParserDIGITS        = 17
	refalParserIDENTIFER     = 18
	refalParserSTRING        = 19
	refalParserSTRING2       = 20
	refalParserCHAR          = 21
	refalParserLINE_COMMENT  = 22
	refalParserBLOCK_COMMENT = 23
	refalParserWS            = 24
)

// refalParser rules.
const (
	refalParserRULE_program       = 0
	refalParserRULE_f_definition  = 1
	refalParserRULE_external_decl = 2
	refalParserRULE_f_name_list   = 3
	refalParserRULE_f_name        = 4
	refalParserRULE_block_        = 5
	refalParserRULE_sentence      = 6
	refalParserRULE_left_side     = 7
	refalParserRULE_conditions    = 8
	refalParserRULE_pattern       = 9
	refalParserRULE_arg_          = 10
	refalParserRULE_right_side    = 11
	refalParserRULE_expression_   = 12
	refalParserRULE_term_         = 13
	refalParserRULE_block_ending  = 14
	refalParserRULE_symbol        = 15
	refalParserRULE_identifier    = 16
	refalParserRULE_variable      = 17
	refalParserRULE_svar          = 18
	refalParserRULE_tvar          = 19
	refalParserRULE_evar          = 20
	refalParserRULE_index         = 21
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
	p.RuleIndex = refalParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) F_definition() IF_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_definitionContext)
}

func (s *ProgramContext) Program() IProgramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *ProgramContext) External_decl() IExternal_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternal_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternal_declContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *refalParser) Program() (localctx IProgramContext) {
	return p.program(0)
}

func (p *refalParser) program(_p int) (localctx IProgramContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewProgramContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IProgramContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, refalParserRULE_program, _p)
	var _la int

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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case refalParserT__1, refalParserIDENTIFER, refalParserSTRING:
		{
			p.SetState(45)
			p.F_definition()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == refalParserT__0 {
				{
					p.SetState(46)
					p.Match(refalParserT__0)
				}

			}
			{
				p.SetState(49)
				p.program(0)
			}

		}

	case refalParserT__4, refalParserT__5, refalParserT__6:
		{
			p.SetState(52)
			p.External_decl()
		}
		{
			p.SetState(53)
			p.Match(refalParserT__0)
		}
		{
			p.SetState(54)
			p.program(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewProgramContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, refalParserRULE_program)
			p.SetState(58)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(59)
				p.External_decl()
			}
			{
				p.SetState(60)
				p.Match(refalParserT__0)
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IF_definitionContext is an interface to support dynamic dispatch.
type IF_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_definitionContext differentiates from other interfaces.
	IsF_definitionContext()
}

type F_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_definitionContext() *F_definitionContext {
	var p = new(F_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_f_definition
	return p
}

func (*F_definitionContext) IsF_definitionContext() {}

func NewF_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_definitionContext {
	var p = new(F_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_f_definition

	return p
}

func (s *F_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *F_definitionContext) F_name() IF_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_nameContext)
}

func (s *F_definitionContext) Block_() IBlock_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_Context)
}

func (s *F_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterF_definition(s)
	}
}

func (s *F_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitF_definition(s)
	}
}

func (p *refalParser) F_definition() (localctx IF_definitionContext) {
	this := p
	_ = this

	localctx = NewF_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, refalParserRULE_f_definition)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == refalParserT__1 {
		{
			p.SetState(67)
			p.Match(refalParserT__1)
		}

	}
	{
		p.SetState(70)
		p.F_name()
	}
	{
		p.SetState(71)
		p.Match(refalParserT__2)
	}
	{
		p.SetState(72)
		p.Block_()
	}
	{
		p.SetState(73)
		p.Match(refalParserT__3)
	}

	return localctx
}

// IExternal_declContext is an interface to support dynamic dispatch.
type IExternal_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternal_declContext differentiates from other interfaces.
	IsExternal_declContext()
}

type External_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternal_declContext() *External_declContext {
	var p = new(External_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_external_decl
	return p
}

func (*External_declContext) IsExternal_declContext() {}

func NewExternal_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *External_declContext {
	var p = new(External_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_external_decl

	return p
}

func (s *External_declContext) GetParser() antlr.Parser { return s.parser }

func (s *External_declContext) F_name_list() IF_name_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_name_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_name_listContext)
}

func (s *External_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *External_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *External_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterExternal_decl(s)
	}
}

func (s *External_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitExternal_decl(s)
	}
}

func (p *refalParser) External_decl() (localctx IExternal_declContext) {
	this := p
	_ = this

	localctx = NewExternal_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, refalParserRULE_external_decl)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<refalParserT__4)|(1<<refalParserT__5)|(1<<refalParserT__6))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(76)
		p.F_name_list()
	}

	return localctx
}

// IF_name_listContext is an interface to support dynamic dispatch.
type IF_name_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_name_listContext differentiates from other interfaces.
	IsF_name_listContext()
}

type F_name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_name_listContext() *F_name_listContext {
	var p = new(F_name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_f_name_list
	return p
}

func (*F_name_listContext) IsF_name_listContext() {}

func NewF_name_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_name_listContext {
	var p = new(F_name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_f_name_list

	return p
}

func (s *F_name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *F_name_listContext) F_name() IF_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_nameContext)
}

func (s *F_name_listContext) F_name_list() IF_name_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IF_name_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IF_name_listContext)
}

func (s *F_name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_name_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterF_name_list(s)
	}
}

func (s *F_name_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitF_name_list(s)
	}
}

func (p *refalParser) F_name_list() (localctx IF_name_listContext) {
	this := p
	_ = this

	localctx = NewF_name_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, refalParserRULE_f_name_list)
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
		p.F_name()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == refalParserT__7 {
		{
			p.SetState(79)
			p.Match(refalParserT__7)
		}
		{
			p.SetState(80)
			p.F_name_list()
		}
		{
			p.SetState(81)
			p.Match(refalParserT__0)
		}

	}

	return localctx
}

// IF_nameContext is an interface to support dynamic dispatch.
type IF_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsF_nameContext differentiates from other interfaces.
	IsF_nameContext()
}

type F_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyF_nameContext() *F_nameContext {
	var p = new(F_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_f_name
	return p
}

func (*F_nameContext) IsF_nameContext() {}

func NewF_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *F_nameContext {
	var p = new(F_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_f_name

	return p
}

func (s *F_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *F_nameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *F_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *F_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *F_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterF_name(s)
	}
}

func (s *F_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitF_name(s)
	}
}

func (p *refalParser) F_name() (localctx IF_nameContext) {
	this := p
	_ = this

	localctx = NewF_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, refalParserRULE_f_name)

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
		p.SetState(85)
		p.Identifier()
	}

	return localctx
}

// IBlock_Context is an interface to support dynamic dispatch.
type IBlock_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlock_Context differentiates from other interfaces.
	IsBlock_Context()
}

type Block_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_Context() *Block_Context {
	var p = new(Block_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_block_
	return p
}

func (*Block_Context) IsBlock_Context() {}

func NewBlock_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_Context {
	var p = new(Block_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_block_

	return p
}

func (s *Block_Context) GetParser() antlr.Parser { return s.parser }

func (s *Block_Context) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *Block_Context) Block_() IBlock_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_Context)
}

func (s *Block_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterBlock_(s)
	}
}

func (s *Block_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitBlock_(s)
	}
}

func (p *refalParser) Block_() (localctx IBlock_Context) {
	this := p
	_ = this

	localctx = NewBlock_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, refalParserRULE_block_)
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
		p.SetState(87)
		p.Sentence()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == refalParserT__0 {
		{
			p.SetState(88)
			p.Match(refalParserT__0)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<refalParserT__7)|(1<<refalParserT__8)|(1<<refalParserT__10)|(1<<refalParserT__12)|(1<<refalParserT__14)|(1<<refalParserT__15)|(1<<refalParserDIGITS)|(1<<refalParserIDENTIFER)|(1<<refalParserSTRING)|(1<<refalParserSTRING2)|(1<<refalParserCHAR))) != 0 {
			{
				p.SetState(89)
				p.Block_()
			}

		}

	}

	return localctx
}

// ISentenceContext is an interface to support dynamic dispatch.
type ISentenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSentenceContext differentiates from other interfaces.
	IsSentenceContext()
}

type SentenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenceContext() *SentenceContext {
	var p = new(SentenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_sentence
	return p
}

func (*SentenceContext) IsSentenceContext() {}

func NewSentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceContext {
	var p = new(SentenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_sentence

	return p
}

func (s *SentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceContext) Left_side() ILeft_sideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeft_sideContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeft_sideContext)
}

func (s *SentenceContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *SentenceContext) Right_side() IRight_sideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRight_sideContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRight_sideContext)
}

func (s *SentenceContext) Block_ending() IBlock_endingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endingContext)
}

func (s *SentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterSentence(s)
	}
}

func (s *SentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitSentence(s)
	}
}

func (p *refalParser) Sentence() (localctx ISentenceContext) {
	this := p
	_ = this

	localctx = NewSentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, refalParserRULE_sentence)

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
		p.Left_side()
	}
	{
		p.SetState(95)
		p.Conditions()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case refalParserT__8:
		{
			p.SetState(96)
			p.Match(refalParserT__8)
		}
		{
			p.SetState(97)
			p.Right_side()
		}

	case refalParserT__7:
		{
			p.SetState(98)
			p.Match(refalParserT__7)
		}
		{
			p.SetState(99)
			p.Block_ending()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILeft_sideContext is an interface to support dynamic dispatch.
type ILeft_sideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeft_sideContext differentiates from other interfaces.
	IsLeft_sideContext()
}

type Left_sideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeft_sideContext() *Left_sideContext {
	var p = new(Left_sideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_left_side
	return p
}

func (*Left_sideContext) IsLeft_sideContext() {}

func NewLeft_sideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Left_sideContext {
	var p = new(Left_sideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_left_side

	return p
}

func (s *Left_sideContext) GetParser() antlr.Parser { return s.parser }

func (s *Left_sideContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Left_sideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Left_sideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Left_sideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterLeft_side(s)
	}
}

func (s *Left_sideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitLeft_side(s)
	}
}

func (p *refalParser) Left_side() (localctx ILeft_sideContext) {
	this := p
	_ = this

	localctx = NewLeft_sideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, refalParserRULE_left_side)

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
		p.Pattern()
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) Arg_() IArg_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_Context)
}

func (s *ConditionsContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *refalParser) Conditions() (localctx IConditionsContext) {
	this := p
	_ = this

	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, refalParserRULE_conditions)

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
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(104)
			p.Match(refalParserT__7)
		}
		{
			p.SetState(105)
			p.Arg_()
		}
		{
			p.SetState(106)
			p.Match(refalParserT__9)
		}
		{
			p.SetState(107)
			p.Pattern()
		}
		{
			p.SetState(108)
			p.Conditions()
		}

	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) Expression_() IExpression_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_Context)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *refalParser) Pattern() (localctx IPatternContext) {
	this := p
	_ = this

	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, refalParserRULE_pattern)

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
		p.SetState(112)
		p.Expression_()
	}

	return localctx
}

// IArg_Context is an interface to support dynamic dispatch.
type IArg_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_Context differentiates from other interfaces.
	IsArg_Context()
}

type Arg_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_Context() *Arg_Context {
	var p = new(Arg_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_arg_
	return p
}

func (*Arg_Context) IsArg_Context() {}

func NewArg_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_Context {
	var p = new(Arg_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_arg_

	return p
}

func (s *Arg_Context) GetParser() antlr.Parser { return s.parser }

func (s *Arg_Context) Expression_() IExpression_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_Context)
}

func (s *Arg_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterArg_(s)
	}
}

func (s *Arg_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitArg_(s)
	}
}

func (p *refalParser) Arg_() (localctx IArg_Context) {
	this := p
	_ = this

	localctx = NewArg_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, refalParserRULE_arg_)

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
		p.SetState(114)
		p.Expression_()
	}

	return localctx
}

// IRight_sideContext is an interface to support dynamic dispatch.
type IRight_sideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRight_sideContext differentiates from other interfaces.
	IsRight_sideContext()
}

type Right_sideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRight_sideContext() *Right_sideContext {
	var p = new(Right_sideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_right_side
	return p
}

func (*Right_sideContext) IsRight_sideContext() {}

func NewRight_sideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Right_sideContext {
	var p = new(Right_sideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_right_side

	return p
}

func (s *Right_sideContext) GetParser() antlr.Parser { return s.parser }

func (s *Right_sideContext) Expression_() IExpression_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_Context)
}

func (s *Right_sideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Right_sideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Right_sideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterRight_side(s)
	}
}

func (s *Right_sideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitRight_side(s)
	}
}

func (p *refalParser) Right_side() (localctx IRight_sideContext) {
	this := p
	_ = this

	localctx = NewRight_sideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, refalParserRULE_right_side)

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
		p.SetState(116)
		p.Expression_()
	}

	return localctx
}

// IExpression_Context is an interface to support dynamic dispatch.
type IExpression_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_Context differentiates from other interfaces.
	IsExpression_Context()
}

type Expression_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_Context() *Expression_Context {
	var p = new(Expression_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_expression_
	return p
}

func (*Expression_Context) IsExpression_Context() {}

func NewExpression_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_Context {
	var p = new(Expression_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_expression_

	return p
}

func (s *Expression_Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression_Context) Term_() ITerm_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm_Context)
}

func (s *Expression_Context) Expression_() IExpression_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_Context)
}

func (s *Expression_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterExpression_(s)
	}
}

func (s *Expression_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitExpression_(s)
	}
}

func (p *refalParser) Expression_() (localctx IExpression_Context) {
	this := p
	_ = this

	localctx = NewExpression_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, refalParserRULE_expression_)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<refalParserT__10)|(1<<refalParserT__12)|(1<<refalParserT__14)|(1<<refalParserT__15)|(1<<refalParserDIGITS)|(1<<refalParserIDENTIFER)|(1<<refalParserSTRING)|(1<<refalParserSTRING2)|(1<<refalParserCHAR))) != 0 {
		{
			p.SetState(118)
			p.Term_()
		}
		{
			p.SetState(119)
			p.Expression_()
		}

	}

	return localctx
}

// ITerm_Context is an interface to support dynamic dispatch.
type ITerm_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm_Context differentiates from other interfaces.
	IsTerm_Context()
}

type Term_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm_Context() *Term_Context {
	var p = new(Term_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_term_
	return p
}

func (*Term_Context) IsTerm_Context() {}

func NewTerm_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term_Context {
	var p = new(Term_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_term_

	return p
}

func (s *Term_Context) GetParser() antlr.Parser { return s.parser }

func (s *Term_Context) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Term_Context) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Term_Context) Expression_() IExpression_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_Context)
}

func (s *Term_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterTerm_(s)
	}
}

func (s *Term_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitTerm_(s)
	}
}

func (p *refalParser) Term_() (localctx ITerm_Context) {
	this := p
	_ = this

	localctx = NewTerm_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, refalParserRULE_term_)

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case refalParserDIGITS, refalParserIDENTIFER, refalParserSTRING, refalParserSTRING2, refalParserCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Symbol()
		}

	case refalParserT__12, refalParserT__14, refalParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Variable()
		}

	case refalParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Match(refalParserT__10)
		}
		{
			p.SetState(126)
			p.Expression_()
		}
		{
			p.SetState(127)
			p.Match(refalParserT__11)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlock_endingContext is an interface to support dynamic dispatch.
type IBlock_endingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlock_endingContext differentiates from other interfaces.
	IsBlock_endingContext()
}

type Block_endingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_endingContext() *Block_endingContext {
	var p = new(Block_endingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_block_ending
	return p
}

func (*Block_endingContext) IsBlock_endingContext() {}

func NewBlock_endingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_endingContext {
	var p = new(Block_endingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_block_ending

	return p
}

func (s *Block_endingContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_endingContext) Arg_() IArg_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_Context)
}

func (s *Block_endingContext) Block_() IBlock_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_Context)
}

func (s *Block_endingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_endingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_endingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterBlock_ending(s)
	}
}

func (s *Block_endingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitBlock_ending(s)
	}
}

func (p *refalParser) Block_ending() (localctx IBlock_endingContext) {
	this := p
	_ = this

	localctx = NewBlock_endingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, refalParserRULE_block_ending)

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
		p.SetState(131)
		p.Arg_()
	}
	{
		p.SetState(132)
		p.Match(refalParserT__9)
	}
	{
		p.SetState(133)
		p.Match(refalParserT__2)
	}
	{
		p.SetState(134)
		p.Block_()
	}
	{
		p.SetState(135)
		p.Match(refalParserT__3)
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SymbolContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(refalParserDIGITS, 0)
}

func (s *SymbolContext) STRING() antlr.TerminalNode {
	return s.GetToken(refalParserSTRING, 0)
}

func (s *SymbolContext) STRING2() antlr.TerminalNode {
	return s.GetToken(refalParserSTRING2, 0)
}

func (s *SymbolContext) CHAR() antlr.TerminalNode {
	return s.GetToken(refalParserCHAR, 0)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *refalParser) Symbol() (localctx ISymbolContext) {
	this := p
	_ = this

	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, refalParserRULE_symbol)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(refalParserDIGITS)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(refalParserSTRING)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(140)
			p.Match(refalParserSTRING2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)
			p.Match(refalParserCHAR)
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
	p.RuleIndex = refalParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFER() antlr.TerminalNode {
	return s.GetToken(refalParserIDENTIFER, 0)
}

func (s *IdentifierContext) STRING() antlr.TerminalNode {
	return s.GetToken(refalParserSTRING, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *refalParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, refalParserRULE_identifier)
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
		p.SetState(144)
		_la = p.GetTokenStream().LA(1)

		if !(_la == refalParserIDENTIFER || _la == refalParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Svar() ISvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISvarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISvarContext)
}

func (s *VariableContext) Tvar() ITvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITvarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITvarContext)
}

func (s *VariableContext) Evar() IEvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEvarContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *refalParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, refalParserRULE_variable)

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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case refalParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Svar()
		}

	case refalParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Tvar()
		}

	case refalParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.Evar()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISvarContext is an interface to support dynamic dispatch.
type ISvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSvarContext differentiates from other interfaces.
	IsSvarContext()
}

type SvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySvarContext() *SvarContext {
	var p = new(SvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_svar
	return p
}

func (*SvarContext) IsSvarContext() {}

func NewSvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SvarContext {
	var p = new(SvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_svar

	return p
}

func (s *SvarContext) GetParser() antlr.Parser { return s.parser }

func (s *SvarContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *SvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterSvar(s)
	}
}

func (s *SvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitSvar(s)
	}
}

func (p *refalParser) Svar() (localctx ISvarContext) {
	this := p
	_ = this

	localctx = NewSvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, refalParserRULE_svar)

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
		p.SetState(151)
		p.Match(refalParserT__12)
	}
	{
		p.SetState(152)
		p.Match(refalParserT__13)
	}
	{
		p.SetState(153)
		p.Index()
	}

	return localctx
}

// ITvarContext is an interface to support dynamic dispatch.
type ITvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTvarContext differentiates from other interfaces.
	IsTvarContext()
}

type TvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTvarContext() *TvarContext {
	var p = new(TvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_tvar
	return p
}

func (*TvarContext) IsTvarContext() {}

func NewTvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TvarContext {
	var p = new(TvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_tvar

	return p
}

func (s *TvarContext) GetParser() antlr.Parser { return s.parser }

func (s *TvarContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *TvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterTvar(s)
	}
}

func (s *TvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitTvar(s)
	}
}

func (p *refalParser) Tvar() (localctx ITvarContext) {
	this := p
	_ = this

	localctx = NewTvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, refalParserRULE_tvar)

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
		p.SetState(155)
		p.Match(refalParserT__14)
	}
	{
		p.SetState(156)
		p.Match(refalParserT__13)
	}
	{
		p.SetState(157)
		p.Index()
	}

	return localctx
}

// IEvarContext is an interface to support dynamic dispatch.
type IEvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEvarContext differentiates from other interfaces.
	IsEvarContext()
}

type EvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvarContext() *EvarContext {
	var p = new(EvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_evar
	return p
}

func (*EvarContext) IsEvarContext() {}

func NewEvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvarContext {
	var p = new(EvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_evar

	return p
}

func (s *EvarContext) GetParser() antlr.Parser { return s.parser }

func (s *EvarContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *EvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterEvar(s)
	}
}

func (s *EvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitEvar(s)
	}
}

func (p *refalParser) Evar() (localctx IEvarContext) {
	this := p
	_ = this

	localctx = NewEvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, refalParserRULE_evar)

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
		p.SetState(159)
		p.Match(refalParserT__15)
	}
	{
		p.SetState(160)
		p.Match(refalParserT__13)
	}
	{
		p.SetState(161)
		p.Index()
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = refalParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = refalParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IndexContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(refalParserDIGITS, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(refalListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *refalParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, refalParserRULE_index)

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

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case refalParserIDENTIFER, refalParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Identifier()
		}

	case refalParserDIGITS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(refalParserDIGITS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *refalParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ProgramContext = nil
		if localctx != nil {
			t = localctx.(*ProgramContext)
		}
		return p.Program_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *refalParser) Program_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
