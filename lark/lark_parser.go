// Code generated from LarkParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lark // LarkParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 202,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 5, 2, 38, 10, 2, 3, 2, 7, 2, 41, 10, 2, 12, 2, 14, 2, 44, 11,
	2, 3, 2, 5, 2, 47, 10, 2, 3, 3, 3, 3, 3, 3, 5, 3, 52, 10, 3, 3, 4, 3, 4,
	3, 4, 5, 4, 57, 10, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 5, 5, 65, 10,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14,
	6, 77, 11, 6, 3, 6, 5, 6, 80, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 86,
	10, 7, 12, 7, 14, 7, 89, 11, 7, 3, 7, 5, 7, 92, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 103, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 113, 10, 9, 13, 9, 14, 9, 114, 5, 9,
	117, 10, 9, 3, 10, 5, 10, 120, 10, 10, 3, 10, 3, 10, 3, 10, 7, 10, 125,
	10, 10, 12, 10, 14, 10, 128, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11,
	134, 10, 11, 12, 11, 14, 11, 137, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 7, 12, 144, 10, 12, 12, 12, 14, 12, 147, 11, 12, 3, 13, 3, 13, 3, 13,
	5, 13, 152, 10, 13, 3, 14, 7, 14, 155, 10, 14, 12, 14, 14, 14, 158, 11,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 166, 10, 15, 5, 15,
	168, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 179, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 7, 17, 191, 10, 17, 12, 17, 14, 17, 194, 11, 17, 3,
	17, 3, 17, 5, 17, 198, 10, 17, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 4, 3, 2, 23,
	24, 3, 2, 21, 22, 2, 216, 2, 42, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 53,
	3, 2, 2, 2, 8, 61, 3, 2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 91, 3, 2, 2, 2,
	14, 93, 3, 2, 2, 2, 16, 116, 3, 2, 2, 2, 18, 119, 3, 2, 2, 2, 20, 129,
	3, 2, 2, 2, 22, 140, 3, 2, 2, 2, 24, 148, 3, 2, 2, 2, 26, 156, 3, 2, 2,
	2, 28, 159, 3, 2, 2, 2, 30, 178, 3, 2, 2, 2, 32, 197, 3, 2, 2, 2, 34, 199,
	3, 2, 2, 2, 36, 38, 5, 4, 3, 2, 37, 36, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2,
	38, 39, 3, 2, 2, 2, 39, 41, 7, 25, 2, 2, 40, 37, 3, 2, 2, 2, 41, 44, 3,
	2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44,
	42, 3, 2, 2, 2, 45, 47, 5, 4, 3, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 3, 3, 2, 2, 2, 48, 52, 5, 6, 4, 2, 49, 52, 5, 8, 5, 2, 50, 52, 5,
	16, 9, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	5, 3, 2, 2, 2, 53, 54, 7, 21, 2, 2, 54, 56, 5, 10, 6, 2, 55, 57, 5, 14,
	8, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59,
	7, 3, 2, 2, 59, 60, 5, 22, 12, 2, 60, 7, 3, 2, 2, 2, 61, 62, 7, 22, 2,
	2, 62, 64, 5, 12, 7, 2, 63, 65, 5, 14, 8, 2, 64, 63, 3, 2, 2, 2, 64, 65,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 7, 3, 2, 2, 67, 68, 5, 22, 12,
	2, 68, 9, 3, 2, 2, 2, 69, 70, 7, 4, 2, 2, 70, 75, 7, 21, 2, 2, 71, 72,
	7, 10, 2, 2, 72, 74, 7, 21, 2, 2, 73, 71, 3, 2, 2, 2, 74, 77, 3, 2, 2,
	2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75,
	3, 2, 2, 2, 78, 80, 7, 5, 2, 2, 79, 69, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 11, 3, 2, 2, 2, 81, 82, 7, 4, 2, 2, 82, 87, 7, 22, 2, 2, 83, 84, 7,
	10, 2, 2, 84, 86, 7, 22, 2, 2, 85, 83, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2,
	87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3,
	2, 2, 2, 90, 92, 7, 5, 2, 2, 91, 81, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92,
	13, 3, 2, 2, 2, 93, 94, 7, 11, 2, 2, 94, 95, 7, 26, 2, 2, 95, 15, 3, 2,
	2, 2, 96, 97, 7, 13, 2, 2, 97, 117, 5, 22, 12, 2, 98, 99, 7, 14, 2, 2,
	99, 102, 5, 18, 10, 2, 100, 101, 7, 12, 2, 2, 101, 103, 5, 34, 18, 2, 102,
	100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 117, 3, 2, 2, 2, 104, 105,
	7, 14, 2, 2, 105, 106, 5, 18, 10, 2, 106, 107, 5, 20, 11, 2, 107, 117,
	3, 2, 2, 2, 108, 109, 7, 15, 2, 2, 109, 117, 5, 6, 4, 2, 110, 112, 7, 16,
	2, 2, 111, 113, 5, 34, 18, 2, 112, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2,
	2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 3, 2, 2, 2, 116,
	96, 3, 2, 2, 2, 116, 98, 3, 2, 2, 2, 116, 104, 3, 2, 2, 2, 116, 108, 3,
	2, 2, 2, 116, 110, 3, 2, 2, 2, 117, 17, 3, 2, 2, 2, 118, 120, 7, 11, 2,
	2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	126, 5, 34, 18, 2, 122, 123, 7, 11, 2, 2, 123, 125, 5, 34, 18, 2, 124,
	122, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 19, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 6,
	2, 2, 130, 135, 5, 34, 18, 2, 131, 132, 7, 10, 2, 2, 132, 134, 5, 34, 18,
	2, 133, 131, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139,
	7, 7, 2, 2, 139, 21, 3, 2, 2, 2, 140, 145, 5, 24, 13, 2, 141, 142, 7, 19,
	2, 2, 142, 144, 5, 24, 13, 2, 143, 141, 3, 2, 2, 2, 144, 147, 3, 2, 2,
	2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 23, 3, 2, 2, 2, 147,
	145, 3, 2, 2, 2, 148, 151, 5, 26, 14, 2, 149, 150, 7, 12, 2, 2, 150, 152,
	7, 21, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 25, 3, 2,
	2, 2, 153, 155, 5, 28, 15, 2, 154, 153, 3, 2, 2, 2, 155, 158, 3, 2, 2,
	2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 27, 3, 2, 2, 2, 158,
	156, 3, 2, 2, 2, 159, 167, 5, 30, 16, 2, 160, 168, 7, 20, 2, 2, 161, 162,
	7, 18, 2, 2, 162, 165, 7, 26, 2, 2, 163, 164, 7, 17, 2, 2, 164, 166, 7,
	26, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2,
	2, 167, 160, 3, 2, 2, 2, 167, 161, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	29, 3, 2, 2, 2, 169, 170, 7, 6, 2, 2, 170, 171, 5, 22, 12, 2, 171, 172,
	7, 7, 2, 2, 172, 179, 3, 2, 2, 2, 173, 174, 7, 8, 2, 2, 174, 175, 5, 22,
	12, 2, 175, 176, 7, 9, 2, 2, 176, 179, 3, 2, 2, 2, 177, 179, 5, 32, 17,
	2, 178, 169, 3, 2, 2, 2, 178, 173, 3, 2, 2, 2, 178, 177, 3, 2, 2, 2, 179,
	31, 3, 2, 2, 2, 180, 181, 7, 23, 2, 2, 181, 182, 7, 17, 2, 2, 182, 198,
	7, 23, 2, 2, 183, 198, 5, 34, 18, 2, 184, 198, 9, 2, 2, 2, 185, 186, 5,
	34, 18, 2, 186, 187, 7, 4, 2, 2, 187, 192, 5, 32, 17, 2, 188, 189, 7, 10,
	2, 2, 189, 191, 5, 32, 17, 2, 190, 188, 3, 2, 2, 2, 191, 194, 3, 2, 2,
	2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 195, 3, 2, 2, 2, 194,
	192, 3, 2, 2, 2, 195, 196, 7, 5, 2, 2, 196, 198, 3, 2, 2, 2, 197, 180,
	3, 2, 2, 2, 197, 183, 3, 2, 2, 2, 197, 184, 3, 2, 2, 2, 197, 185, 3, 2,
	2, 2, 198, 33, 3, 2, 2, 2, 199, 200, 9, 3, 2, 2, 200, 35, 3, 2, 2, 2, 26,
	37, 42, 46, 51, 56, 64, 75, 79, 87, 91, 102, 114, 116, 119, 126, 135, 145,
	151, 156, 165, 167, 178, 192, 197,
}
var literalNames = []string{
	"", "':'", "'{'", "'}'", "'('", "')'", "'['", "']'", "','", "'.'", "'->'",
	"'%ignore'", "'%import'", "'%override'", "'%declare'", "'..'", "'~'", "'|'",
}
var symbolicNames = []string{
	"", "COLON", "LC", "RC", "LP", "RP", "LB", "RB", "COMMA", "DOT", "ARROW",
	"IGNORE", "IMPORT", "OVERRIDE", "DECLARE", "DD", "SQ", "VBAR", "OP", "RULE",
	"TOKEN", "STRING", "REGEXP", "NL", "NUMBER", "WS_INLINE", "COMMENT",
}

var ruleNames = []string{
	"start", "item", "rule_", "token", "rule_params", "token_params", "priority",
	"statement", "import_path", "name_list", "expansions", "alias", "expansion",
	"expr", "atom", "value", "name",
}

type LarkParser struct {
	*antlr.BaseParser
}

// NewLarkParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LarkParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLarkParser(input antlr.TokenStream) *LarkParser {
	this := new(LarkParser)
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
	this.GrammarFileName = "LarkParser.g4"

	return this
}

// LarkParser tokens.
const (
	LarkParserEOF       = antlr.TokenEOF
	LarkParserCOLON     = 1
	LarkParserLC        = 2
	LarkParserRC        = 3
	LarkParserLP        = 4
	LarkParserRP        = 5
	LarkParserLB        = 6
	LarkParserRB        = 7
	LarkParserCOMMA     = 8
	LarkParserDOT       = 9
	LarkParserARROW     = 10
	LarkParserIGNORE    = 11
	LarkParserIMPORT    = 12
	LarkParserOVERRIDE  = 13
	LarkParserDECLARE   = 14
	LarkParserDD        = 15
	LarkParserSQ        = 16
	LarkParserVBAR      = 17
	LarkParserOP        = 18
	LarkParserRULE      = 19
	LarkParserTOKEN     = 20
	LarkParserSTRING    = 21
	LarkParserREGEXP    = 22
	LarkParserNL        = 23
	LarkParserNUMBER    = 24
	LarkParserWS_INLINE = 25
	LarkParserCOMMENT   = 26
)

// LarkParser rules.
const (
	LarkParserRULE_start        = 0
	LarkParserRULE_item         = 1
	LarkParserRULE_rule_        = 2
	LarkParserRULE_token        = 3
	LarkParserRULE_rule_params  = 4
	LarkParserRULE_token_params = 5
	LarkParserRULE_priority     = 6
	LarkParserRULE_statement    = 7
	LarkParserRULE_import_path  = 8
	LarkParserRULE_name_list    = 9
	LarkParserRULE_expansions   = 10
	LarkParserRULE_alias        = 11
	LarkParserRULE_expansion    = 12
	LarkParserRULE_expr         = 13
	LarkParserRULE_atom         = 14
	LarkParserRULE_value        = 15
	LarkParserRULE_name         = 16
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(LarkParserNL)
}

func (s *StartContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserNL, i)
}

func (s *StartContext) AllItem() []IItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemContext)(nil)).Elem())
	var tst = make([]IItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemContext)
		}
	}

	return tst
}

func (s *StartContext) Item(i int) IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *LarkParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LarkParserRULE_start)
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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LarkParserIGNORE)|(1<<LarkParserIMPORT)|(1<<LarkParserOVERRIDE)|(1<<LarkParserDECLARE)|(1<<LarkParserRULE)|(1<<LarkParserTOKEN))) != 0 {
				{
					p.SetState(34)
					p.Item()
				}

			}
			{
				p.SetState(37)
				p.Match(LarkParserNL)
			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LarkParserIGNORE)|(1<<LarkParserIMPORT)|(1<<LarkParserOVERRIDE)|(1<<LarkParserDECLARE)|(1<<LarkParserRULE)|(1<<LarkParserTOKEN))) != 0 {
		{
			p.SetState(43)
			p.Item()
		}

	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) Rule_() IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *ItemContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *ItemContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *LarkParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LarkParserRULE_item)

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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LarkParserRULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Rule_()
		}

	case LarkParserTOKEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Token()
		}

	case LarkParserIGNORE, LarkParserIMPORT, LarkParserOVERRIDE, LarkParserDECLARE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRule_Context is an interface to support dynamic dispatch.
type IRule_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_Context differentiates from other interfaces.
	IsRule_Context()
}

type Rule_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_Context() *Rule_Context {
	var p = new(Rule_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_rule_
	return p
}

func (*Rule_Context) IsRule_Context() {}

func NewRule_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_Context {
	var p = new(Rule_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_rule_

	return p
}

func (s *Rule_Context) GetParser() antlr.Parser { return s.parser }

func (s *Rule_Context) RULE() antlr.TerminalNode {
	return s.GetToken(LarkParserRULE, 0)
}

func (s *Rule_Context) Rule_params() IRule_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRule_paramsContext)
}

func (s *Rule_Context) COLON() antlr.TerminalNode {
	return s.GetToken(LarkParserCOLON, 0)
}

func (s *Rule_Context) Expansions() IExpansionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpansionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpansionsContext)
}

func (s *Rule_Context) Priority() IPriorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPriorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPriorityContext)
}

func (s *Rule_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterRule_(s)
	}
}

func (s *Rule_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitRule_(s)
	}
}

func (p *LarkParser) Rule_() (localctx IRule_Context) {
	this := p
	_ = this

	localctx = NewRule_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LarkParserRULE_rule_)
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
		p.SetState(51)
		p.Match(LarkParserRULE)
	}
	{
		p.SetState(52)
		p.Rule_params()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserDOT {
		{
			p.SetState(53)
			p.Priority()
		}

	}
	{
		p.SetState(56)
		p.Match(LarkParserCOLON)
	}
	{
		p.SetState(57)
		p.Expansions()
	}

	return localctx
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_token
	return p
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(LarkParserTOKEN, 0)
}

func (s *TokenContext) Token_params() IToken_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IToken_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IToken_paramsContext)
}

func (s *TokenContext) COLON() antlr.TerminalNode {
	return s.GetToken(LarkParserCOLON, 0)
}

func (s *TokenContext) Expansions() IExpansionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpansionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpansionsContext)
}

func (s *TokenContext) Priority() IPriorityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPriorityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPriorityContext)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitToken(s)
	}
}

func (p *LarkParser) Token() (localctx ITokenContext) {
	this := p
	_ = this

	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LarkParserRULE_token)
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
		p.SetState(59)
		p.Match(LarkParserTOKEN)
	}
	{
		p.SetState(60)
		p.Token_params()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserDOT {
		{
			p.SetState(61)
			p.Priority()
		}

	}
	{
		p.SetState(64)
		p.Match(LarkParserCOLON)
	}
	{
		p.SetState(65)
		p.Expansions()
	}

	return localctx
}

// IRule_paramsContext is an interface to support dynamic dispatch.
type IRule_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_paramsContext differentiates from other interfaces.
	IsRule_paramsContext()
}

type Rule_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_paramsContext() *Rule_paramsContext {
	var p = new(Rule_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_rule_params
	return p
}

func (*Rule_paramsContext) IsRule_paramsContext() {}

func NewRule_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_paramsContext {
	var p = new(Rule_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_rule_params

	return p
}

func (s *Rule_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_paramsContext) LC() antlr.TerminalNode {
	return s.GetToken(LarkParserLC, 0)
}

func (s *Rule_paramsContext) AllRULE() []antlr.TerminalNode {
	return s.GetTokens(LarkParserRULE)
}

func (s *Rule_paramsContext) RULE(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserRULE, i)
}

func (s *Rule_paramsContext) RC() antlr.TerminalNode {
	return s.GetToken(LarkParserRC, 0)
}

func (s *Rule_paramsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LarkParserCOMMA)
}

func (s *Rule_paramsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserCOMMA, i)
}

func (s *Rule_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterRule_params(s)
	}
}

func (s *Rule_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitRule_params(s)
	}
}

func (p *LarkParser) Rule_params() (localctx IRule_paramsContext) {
	this := p
	_ = this

	localctx = NewRule_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LarkParserRULE_rule_params)
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
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserLC {
		{
			p.SetState(67)
			p.Match(LarkParserLC)
		}
		{
			p.SetState(68)
			p.Match(LarkParserRULE)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LarkParserCOMMA {
			{
				p.SetState(69)
				p.Match(LarkParserCOMMA)
			}
			{
				p.SetState(70)
				p.Match(LarkParserRULE)
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(76)
			p.Match(LarkParserRC)
		}

	}

	return localctx
}

// IToken_paramsContext is an interface to support dynamic dispatch.
type IToken_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToken_paramsContext differentiates from other interfaces.
	IsToken_paramsContext()
}

type Token_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToken_paramsContext() *Token_paramsContext {
	var p = new(Token_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_token_params
	return p
}

func (*Token_paramsContext) IsToken_paramsContext() {}

func NewToken_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Token_paramsContext {
	var p = new(Token_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_token_params

	return p
}

func (s *Token_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Token_paramsContext) LC() antlr.TerminalNode {
	return s.GetToken(LarkParserLC, 0)
}

func (s *Token_paramsContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(LarkParserTOKEN)
}

func (s *Token_paramsContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserTOKEN, i)
}

func (s *Token_paramsContext) RC() antlr.TerminalNode {
	return s.GetToken(LarkParserRC, 0)
}

func (s *Token_paramsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LarkParserCOMMA)
}

func (s *Token_paramsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserCOMMA, i)
}

func (s *Token_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Token_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Token_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterToken_params(s)
	}
}

func (s *Token_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitToken_params(s)
	}
}

func (p *LarkParser) Token_params() (localctx IToken_paramsContext) {
	this := p
	_ = this

	localctx = NewToken_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LarkParserRULE_token_params)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserLC {
		{
			p.SetState(79)
			p.Match(LarkParserLC)
		}
		{
			p.SetState(80)
			p.Match(LarkParserTOKEN)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LarkParserCOMMA {
			{
				p.SetState(81)
				p.Match(LarkParserCOMMA)
			}
			{
				p.SetState(82)
				p.Match(LarkParserTOKEN)
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(88)
			p.Match(LarkParserRC)
		}

	}

	return localctx
}

// IPriorityContext is an interface to support dynamic dispatch.
type IPriorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPriorityContext differentiates from other interfaces.
	IsPriorityContext()
}

type PriorityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriorityContext() *PriorityContext {
	var p = new(PriorityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_priority
	return p
}

func (*PriorityContext) IsPriorityContext() {}

func NewPriorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PriorityContext {
	var p = new(PriorityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_priority

	return p
}

func (s *PriorityContext) GetParser() antlr.Parser { return s.parser }

func (s *PriorityContext) DOT() antlr.TerminalNode {
	return s.GetToken(LarkParserDOT, 0)
}

func (s *PriorityContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LarkParserNUMBER, 0)
}

func (s *PriorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PriorityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterPriority(s)
	}
}

func (s *PriorityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitPriority(s)
	}
}

func (p *LarkParser) Priority() (localctx IPriorityContext) {
	this := p
	_ = this

	localctx = NewPriorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LarkParserRULE_priority)

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
		p.SetState(91)
		p.Match(LarkParserDOT)
	}
	{
		p.SetState(92)
		p.Match(LarkParserNUMBER)
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
	p.RuleIndex = LarkParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IGNORE() antlr.TerminalNode {
	return s.GetToken(LarkParserIGNORE, 0)
}

func (s *StatementContext) Expansions() IExpansionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpansionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpansionsContext)
}

func (s *StatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(LarkParserIMPORT, 0)
}

func (s *StatementContext) Import_path() IImport_pathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_pathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImport_pathContext)
}

func (s *StatementContext) ARROW() antlr.TerminalNode {
	return s.GetToken(LarkParserARROW, 0)
}

func (s *StatementContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *StatementContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *StatementContext) Name_list() IName_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IName_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IName_listContext)
}

func (s *StatementContext) OVERRIDE() antlr.TerminalNode {
	return s.GetToken(LarkParserOVERRIDE, 0)
}

func (s *StatementContext) Rule_() IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *StatementContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(LarkParserDECLARE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *LarkParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LarkParserRULE_statement)
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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(LarkParserIGNORE)
		}
		{
			p.SetState(95)
			p.Expansions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(LarkParserIMPORT)
		}
		{
			p.SetState(97)
			p.Import_path()
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LarkParserARROW {
			{
				p.SetState(98)
				p.Match(LarkParserARROW)
			}
			{
				p.SetState(99)
				p.Name()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Match(LarkParserIMPORT)
		}
		{
			p.SetState(103)
			p.Import_path()
		}
		{
			p.SetState(104)
			p.Name_list()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(LarkParserOVERRIDE)
		}
		{
			p.SetState(107)
			p.Rule_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)
			p.Match(LarkParserDECLARE)
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LarkParserRULE || _la == LarkParserTOKEN {
			{
				p.SetState(109)
				p.Name()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IImport_pathContext is an interface to support dynamic dispatch.
type IImport_pathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_pathContext differentiates from other interfaces.
	IsImport_pathContext()
}

type Import_pathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_pathContext() *Import_pathContext {
	var p = new(Import_pathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_import_path
	return p
}

func (*Import_pathContext) IsImport_pathContext() {}

func NewImport_pathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_pathContext {
	var p = new(Import_pathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_import_path

	return p
}

func (s *Import_pathContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_pathContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *Import_pathContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Import_pathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(LarkParserDOT)
}

func (s *Import_pathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserDOT, i)
}

func (s *Import_pathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_pathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_pathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterImport_path(s)
	}
}

func (s *Import_pathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitImport_path(s)
	}
}

func (p *LarkParser) Import_path() (localctx IImport_pathContext) {
	this := p
	_ = this

	localctx = NewImport_pathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LarkParserRULE_import_path)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserDOT {
		{
			p.SetState(116)
			p.Match(LarkParserDOT)
		}

	}
	{
		p.SetState(119)
		p.Name()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LarkParserDOT {
		{
			p.SetState(120)
			p.Match(LarkParserDOT)
		}
		{
			p.SetState(121)
			p.Name()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IName_listContext is an interface to support dynamic dispatch.
type IName_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsName_listContext differentiates from other interfaces.
	IsName_listContext()
}

type Name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_listContext() *Name_listContext {
	var p = new(Name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_name_list
	return p
}

func (*Name_listContext) IsName_listContext() {}

func NewName_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_listContext {
	var p = new(Name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_name_list

	return p
}

func (s *Name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_listContext) LP() antlr.TerminalNode {
	return s.GetToken(LarkParserLP, 0)
}

func (s *Name_listContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *Name_listContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Name_listContext) RP() antlr.TerminalNode {
	return s.GetToken(LarkParserRP, 0)
}

func (s *Name_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LarkParserCOMMA)
}

func (s *Name_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserCOMMA, i)
}

func (s *Name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterName_list(s)
	}
}

func (s *Name_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitName_list(s)
	}
}

func (p *LarkParser) Name_list() (localctx IName_listContext) {
	this := p
	_ = this

	localctx = NewName_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LarkParserRULE_name_list)
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
		p.SetState(127)
		p.Match(LarkParserLP)
	}
	{
		p.SetState(128)
		p.Name()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LarkParserCOMMA {
		{
			p.SetState(129)
			p.Match(LarkParserCOMMA)
		}
		{
			p.SetState(130)
			p.Name()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(LarkParserRP)
	}

	return localctx
}

// IExpansionsContext is an interface to support dynamic dispatch.
type IExpansionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpansionsContext differentiates from other interfaces.
	IsExpansionsContext()
}

type ExpansionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpansionsContext() *ExpansionsContext {
	var p = new(ExpansionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_expansions
	return p
}

func (*ExpansionsContext) IsExpansionsContext() {}

func NewExpansionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpansionsContext {
	var p = new(ExpansionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_expansions

	return p
}

func (s *ExpansionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpansionsContext) AllAlias() []IAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasContext)(nil)).Elem())
	var tst = make([]IAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasContext)
		}
	}

	return tst
}

func (s *ExpansionsContext) Alias(i int) IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *ExpansionsContext) AllVBAR() []antlr.TerminalNode {
	return s.GetTokens(LarkParserVBAR)
}

func (s *ExpansionsContext) VBAR(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserVBAR, i)
}

func (s *ExpansionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpansionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpansionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterExpansions(s)
	}
}

func (s *ExpansionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitExpansions(s)
	}
}

func (p *LarkParser) Expansions() (localctx IExpansionsContext) {
	this := p
	_ = this

	localctx = NewExpansionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LarkParserRULE_expansions)
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
		p.SetState(138)
		p.Alias()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LarkParserVBAR {
		{
			p.SetState(139)
			p.Match(LarkParserVBAR)
		}
		{
			p.SetState(140)
			p.Alias()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Expansion() IExpansionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpansionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpansionContext)
}

func (s *AliasContext) ARROW() antlr.TerminalNode {
	return s.GetToken(LarkParserARROW, 0)
}

func (s *AliasContext) RULE() antlr.TerminalNode {
	return s.GetToken(LarkParserRULE, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *LarkParser) Alias() (localctx IAliasContext) {
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LarkParserRULE_alias)
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
		p.SetState(146)
		p.Expansion()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LarkParserARROW {
		{
			p.SetState(147)
			p.Match(LarkParserARROW)
		}
		{
			p.SetState(148)
			p.Match(LarkParserRULE)
		}

	}

	return localctx
}

// IExpansionContext is an interface to support dynamic dispatch.
type IExpansionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpansionContext differentiates from other interfaces.
	IsExpansionContext()
}

type ExpansionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpansionContext() *ExpansionContext {
	var p = new(ExpansionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_expansion
	return p
}

func (*ExpansionContext) IsExpansionContext() {}

func NewExpansionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpansionContext {
	var p = new(ExpansionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_expansion

	return p
}

func (s *ExpansionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpansionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExpansionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpansionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpansionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterExpansion(s)
	}
}

func (s *ExpansionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitExpansion(s)
	}
}

func (p *LarkParser) Expansion() (localctx IExpansionContext) {
	this := p
	_ = this

	localctx = NewExpansionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LarkParserRULE_expansion)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LarkParserLP)|(1<<LarkParserLB)|(1<<LarkParserRULE)|(1<<LarkParserTOKEN)|(1<<LarkParserSTRING)|(1<<LarkParserREGEXP))) != 0 {
		{
			p.SetState(151)
			p.Expr()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = LarkParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExprContext) OP() antlr.TerminalNode {
	return s.GetToken(LarkParserOP, 0)
}

func (s *ExprContext) SQ() antlr.TerminalNode {
	return s.GetToken(LarkParserSQ, 0)
}

func (s *ExprContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(LarkParserNUMBER)
}

func (s *ExprContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserNUMBER, i)
}

func (s *ExprContext) DD() antlr.TerminalNode {
	return s.GetToken(LarkParserDD, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *LarkParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LarkParserRULE_expr)
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
		p.Atom()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LarkParserOP:
		{
			p.SetState(158)
			p.Match(LarkParserOP)
		}

	case LarkParserSQ:
		{
			p.SetState(159)
			p.Match(LarkParserSQ)
		}
		{
			p.SetState(160)
			p.Match(LarkParserNUMBER)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LarkParserDD {
			{
				p.SetState(161)
				p.Match(LarkParserDD)
			}
			{
				p.SetState(162)
				p.Match(LarkParserNUMBER)
			}

		}

	case LarkParserEOF, LarkParserLP, LarkParserRP, LarkParserLB, LarkParserRB, LarkParserARROW, LarkParserVBAR, LarkParserRULE, LarkParserTOKEN, LarkParserSTRING, LarkParserREGEXP, LarkParserNL:

	default:
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) LP() antlr.TerminalNode {
	return s.GetToken(LarkParserLP, 0)
}

func (s *AtomContext) Expansions() IExpansionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpansionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpansionsContext)
}

func (s *AtomContext) RP() antlr.TerminalNode {
	return s.GetToken(LarkParserRP, 0)
}

func (s *AtomContext) LB() antlr.TerminalNode {
	return s.GetToken(LarkParserLB, 0)
}

func (s *AtomContext) RB() antlr.TerminalNode {
	return s.GetToken(LarkParserRB, 0)
}

func (s *AtomContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *LarkParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LarkParserRULE_atom)

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

	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LarkParserLP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(LarkParserLP)
		}
		{
			p.SetState(168)
			p.Expansions()
		}
		{
			p.SetState(169)
			p.Match(LarkParserRP)
		}

	case LarkParserLB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(LarkParserLB)
		}
		{
			p.SetState(172)
			p.Expansions()
		}
		{
			p.SetState(173)
			p.Match(LarkParserRB)
		}

	case LarkParserRULE, LarkParserTOKEN, LarkParserSTRING, LarkParserREGEXP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LarkParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(LarkParserSTRING)
}

func (s *ValueContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserSTRING, i)
}

func (s *ValueContext) DD() antlr.TerminalNode {
	return s.GetToken(LarkParserDD, 0)
}

func (s *ValueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ValueContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(LarkParserREGEXP, 0)
}

func (s *ValueContext) LC() antlr.TerminalNode {
	return s.GetToken(LarkParserLC, 0)
}

func (s *ValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueContext) RC() antlr.TerminalNode {
	return s.GetToken(LarkParserRC, 0)
}

func (s *ValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LarkParserCOMMA)
}

func (s *ValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LarkParserCOMMA, i)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *LarkParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LarkParserRULE_value)
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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Match(LarkParserSTRING)
		}
		{
			p.SetState(179)
			p.Match(LarkParserDD)
		}
		{
			p.SetState(180)
			p.Match(LarkParserSTRING)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LarkParserSTRING || _la == LarkParserREGEXP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Name()
		}
		{
			p.SetState(184)
			p.Match(LarkParserLC)
		}
		{
			p.SetState(185)
			p.Value()
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LarkParserCOMMA {
			{
				p.SetState(186)
				p.Match(LarkParserCOMMA)
			}
			{
				p.SetState(187)
				p.Value()
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(193)
			p.Match(LarkParserRC)
		}

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
	p.RuleIndex = LarkParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LarkParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) RULE() antlr.TerminalNode {
	return s.GetToken(LarkParserRULE, 0)
}

func (s *NameContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(LarkParserTOKEN, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LarkParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *LarkParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LarkParserRULE_name)
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
		p.SetState(197)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LarkParserRULE || _la == LarkParserTOKEN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
