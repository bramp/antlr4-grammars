// Code generated from alpaca.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alpaca // alpaca
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 31, 191,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 55, 10, 3, 12,
	3, 14, 3, 58, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3, 5, 3,
	5, 5, 5, 68, 10, 5, 3, 5, 7, 5, 71, 10, 5, 12, 5, 14, 5, 74, 11, 5, 3,
	5, 5, 5, 77, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6, 82, 10, 6, 12, 6, 14, 6, 85,
	11, 6, 3, 6, 5, 6, 88, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 7, 13, 108, 10, 13, 12, 13, 14, 13, 111, 11, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 117, 10, 14, 3, 15, 3, 15, 3, 15, 5, 15, 122, 10, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 127, 10, 16, 12, 16, 14, 16, 130, 11, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 141, 10,
	17, 3, 18, 3, 18, 5, 18, 145, 10, 18, 3, 18, 3, 18, 5, 18, 149, 10, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 155, 10, 19, 5, 19, 157, 10, 19, 3,
	19, 3, 19, 5, 19, 161, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 167,
	10, 21, 12, 21, 14, 21, 170, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22,
	176, 10, 22, 12, 22, 14, 22, 179, 11, 22, 3, 23, 6, 23, 182, 10, 23, 13,
	23, 14, 23, 183, 3, 24, 6, 24, 187, 10, 24, 13, 24, 14, 24, 188, 3, 24,
	2, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 2, 6, 3, 2, 3, 4, 3, 2, 14, 16, 3, 2, 22, 24, 3,
	2, 27, 28, 2, 193, 2, 48, 3, 2, 2, 2, 4, 51, 3, 2, 2, 2, 6, 62, 3, 2, 2,
	2, 8, 64, 3, 2, 2, 2, 10, 78, 3, 2, 2, 2, 12, 89, 3, 2, 2, 2, 14, 93, 3,
	2, 2, 2, 16, 95, 3, 2, 2, 2, 18, 97, 3, 2, 2, 2, 20, 99, 3, 2, 2, 2, 22,
	101, 3, 2, 2, 2, 24, 104, 3, 2, 2, 2, 26, 112, 3, 2, 2, 2, 28, 121, 3,
	2, 2, 2, 30, 123, 3, 2, 2, 2, 32, 140, 3, 2, 2, 2, 34, 142, 3, 2, 2, 2,
	36, 150, 3, 2, 2, 2, 38, 162, 3, 2, 2, 2, 40, 164, 3, 2, 2, 2, 42, 173,
	3, 2, 2, 2, 44, 181, 3, 2, 2, 2, 46, 186, 3, 2, 2, 2, 48, 49, 5, 4, 3,
	2, 49, 50, 9, 2, 2, 2, 50, 3, 3, 2, 2, 2, 51, 56, 5, 6, 4, 2, 52, 53, 7,
	5, 2, 2, 53, 55, 5, 6, 4, 2, 54, 52, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56,
	54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 5, 3, 2, 2, 2, 58, 56, 3, 2, 2,
	2, 59, 63, 5, 8, 5, 2, 60, 63, 5, 10, 6, 2, 61, 63, 5, 12, 7, 2, 62, 59,
	3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 7, 3, 2, 2, 2,
	64, 65, 7, 6, 2, 2, 65, 67, 5, 16, 9, 2, 66, 68, 7, 25, 2, 2, 67, 66, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 72, 3, 2, 2, 2, 69, 71, 5, 20, 11, 2,
	70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3,
	2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 77, 5, 24, 13, 2,
	76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 9, 3, 2, 2, 2, 78, 79, 7, 7,
	2, 2, 79, 83, 5, 14, 8, 2, 80, 82, 5, 20, 11, 2, 81, 80, 3, 2, 2, 2, 82,
	85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3, 2, 2,
	2, 85, 83, 3, 2, 2, 2, 86, 88, 5, 24, 13, 2, 87, 86, 3, 2, 2, 2, 87, 88,
	3, 2, 2, 2, 88, 11, 3, 2, 2, 2, 89, 90, 7, 8, 2, 2, 90, 91, 5, 18, 10,
	2, 91, 92, 5, 40, 21, 2, 92, 13, 3, 2, 2, 2, 93, 94, 5, 42, 22, 2, 94,
	15, 3, 2, 2, 2, 95, 96, 5, 42, 22, 2, 96, 17, 3, 2, 2, 2, 97, 98, 5, 42,
	22, 2, 98, 19, 3, 2, 2, 2, 99, 100, 5, 22, 12, 2, 100, 21, 3, 2, 2, 2,
	101, 102, 7, 9, 2, 2, 102, 103, 5, 14, 8, 2, 103, 23, 3, 2, 2, 2, 104,
	109, 5, 26, 14, 2, 105, 106, 7, 10, 2, 2, 106, 108, 5, 26, 14, 2, 107,
	105, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110,
	3, 2, 2, 2, 110, 25, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 113, 7, 11,
	2, 2, 113, 116, 5, 28, 15, 2, 114, 115, 7, 12, 2, 2, 115, 117, 5, 30, 16,
	2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 27, 3, 2, 2, 2, 118,
	122, 5, 16, 9, 2, 119, 122, 5, 46, 24, 2, 120, 122, 7, 13, 2, 2, 121, 118,
	3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 29, 3, 2,
	2, 2, 123, 128, 5, 32, 17, 2, 124, 125, 9, 3, 2, 2, 125, 127, 5, 32, 17,
	2, 126, 124, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 31, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 141, 5,
	36, 19, 2, 132, 133, 7, 17, 2, 2, 133, 134, 5, 30, 16, 2, 134, 135, 7,
	18, 2, 2, 135, 141, 3, 2, 2, 2, 136, 137, 7, 19, 2, 2, 137, 141, 5, 32,
	17, 2, 138, 141, 5, 38, 20, 2, 139, 141, 5, 34, 18, 2, 140, 131, 3, 2,
	2, 2, 140, 132, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2,
	140, 139, 3, 2, 2, 2, 141, 33, 3, 2, 2, 2, 142, 148, 5, 28, 15, 2, 143,
	145, 7, 20, 2, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146,
	3, 2, 2, 2, 146, 149, 5, 28, 15, 2, 147, 149, 5, 22, 12, 2, 148, 144, 3,
	2, 2, 2, 148, 147, 3, 2, 2, 2, 149, 35, 3, 2, 2, 2, 150, 156, 5, 44, 23,
	2, 151, 154, 7, 21, 2, 2, 152, 155, 5, 40, 21, 2, 153, 155, 5, 18, 10,
	2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 156,
	151, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 161,
	5, 28, 15, 2, 159, 161, 5, 22, 12, 2, 160, 158, 3, 2, 2, 2, 160, 159, 3,
	2, 2, 2, 161, 37, 3, 2, 2, 2, 162, 163, 9, 4, 2, 2, 163, 39, 3, 2, 2, 2,
	164, 168, 7, 17, 2, 2, 165, 167, 5, 46, 24, 2, 166, 165, 3, 2, 2, 2, 167,
	170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171,
	3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171, 172, 7, 18, 2, 2, 172, 41, 3, 2,
	2, 2, 173, 177, 7, 27, 2, 2, 174, 176, 9, 5, 2, 2, 175, 174, 3, 2, 2, 2,
	176, 179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178,
	43, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 182, 7, 28, 2, 2, 181, 180,
	3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2,
	2, 2, 184, 45, 3, 2, 2, 2, 185, 187, 7, 29, 2, 2, 186, 185, 3, 2, 2, 2,
	187, 188, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189,
	47, 3, 2, 2, 2, 23, 56, 62, 67, 72, 76, 83, 87, 109, 116, 121, 128, 140,
	144, 148, 154, 156, 160, 168, 177, 183, 188,
}
var literalNames = []string{
	"", "'.'", "'begin'", "';'", "'state'", "'class'", "'neighbourhood'", "'is'",
	"','", "'to'", "'when'", "'me'", "'and'", "'or'", "'xor'", "'('", "')'",
	"'not'", "'='", "'in'", "'true'", "'false'", "'guess'", "", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "QUOTEDCHAR", "QUOTE", "ALPHA", "DIGIT", "ARROW", "COMMENT",
	"WS",
}

var ruleNames = []string{
	"prog", "defns", "defn", "stateDefn", "classDefn", "nbhdDefn", "classID",
	"stateID", "nbhdID", "membershipDecl", "classRef", "rules", "rule_", "stateRef",
	"expression", "term", "relationalPred", "adjacencyPred", "boolPrimitive",
	"neigbourhood", "identifier", "naturalnumber", "arrowchain",
}

type alpacaParser struct {
	*antlr.BaseParser
}

// NewalpacaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *alpacaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewalpacaParser(input antlr.TokenStream) *alpacaParser {
	this := new(alpacaParser)
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
	this.GrammarFileName = "alpaca.g4"

	return this
}

// alpacaParser tokens.
const (
	alpacaParserEOF        = antlr.TokenEOF
	alpacaParserT__0       = 1
	alpacaParserT__1       = 2
	alpacaParserT__2       = 3
	alpacaParserT__3       = 4
	alpacaParserT__4       = 5
	alpacaParserT__5       = 6
	alpacaParserT__6       = 7
	alpacaParserT__7       = 8
	alpacaParserT__8       = 9
	alpacaParserT__9       = 10
	alpacaParserT__10      = 11
	alpacaParserT__11      = 12
	alpacaParserT__12      = 13
	alpacaParserT__13      = 14
	alpacaParserT__14      = 15
	alpacaParserT__15      = 16
	alpacaParserT__16      = 17
	alpacaParserT__17      = 18
	alpacaParserT__18      = 19
	alpacaParserT__19      = 20
	alpacaParserT__20      = 21
	alpacaParserT__21      = 22
	alpacaParserQUOTEDCHAR = 23
	alpacaParserQUOTE      = 24
	alpacaParserALPHA      = 25
	alpacaParserDIGIT      = 26
	alpacaParserARROW      = 27
	alpacaParserCOMMENT    = 28
	alpacaParserWS         = 29
)

// alpacaParser rules.
const (
	alpacaParserRULE_prog           = 0
	alpacaParserRULE_defns          = 1
	alpacaParserRULE_defn           = 2
	alpacaParserRULE_stateDefn      = 3
	alpacaParserRULE_classDefn      = 4
	alpacaParserRULE_nbhdDefn       = 5
	alpacaParserRULE_classID        = 6
	alpacaParserRULE_stateID        = 7
	alpacaParserRULE_nbhdID         = 8
	alpacaParserRULE_membershipDecl = 9
	alpacaParserRULE_classRef       = 10
	alpacaParserRULE_rules          = 11
	alpacaParserRULE_rule_          = 12
	alpacaParserRULE_stateRef       = 13
	alpacaParserRULE_expression     = 14
	alpacaParserRULE_term           = 15
	alpacaParserRULE_relationalPred = 16
	alpacaParserRULE_adjacencyPred  = 17
	alpacaParserRULE_boolPrimitive  = 18
	alpacaParserRULE_neigbourhood   = 19
	alpacaParserRULE_identifier     = 20
	alpacaParserRULE_naturalnumber  = 21
	alpacaParserRULE_arrowchain     = 22
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
	p.RuleIndex = alpacaParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Defns() IDefnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefnsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *alpacaParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, alpacaParserRULE_prog)
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
		p.SetState(46)
		p.Defns()
	}
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(_la == alpacaParserT__0 || _la == alpacaParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDefnsContext is an interface to support dynamic dispatch.
type IDefnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefnsContext differentiates from other interfaces.
	IsDefnsContext()
}

type DefnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefnsContext() *DefnsContext {
	var p = new(DefnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_defns
	return p
}

func (*DefnsContext) IsDefnsContext() {}

func NewDefnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefnsContext {
	var p = new(DefnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_defns

	return p
}

func (s *DefnsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefnsContext) AllDefn() []IDefnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefnContext)(nil)).Elem())
	var tst = make([]IDefnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefnContext)
		}
	}

	return tst
}

func (s *DefnsContext) Defn(i int) IDefnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefnContext)
}

func (s *DefnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterDefns(s)
	}
}

func (s *DefnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitDefns(s)
	}
}

func (p *alpacaParser) Defns() (localctx IDefnsContext) {
	this := p
	_ = this

	localctx = NewDefnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, alpacaParserRULE_defns)
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
		p.Defn()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alpacaParserT__2 {
		{
			p.SetState(50)
			p.Match(alpacaParserT__2)
		}
		{
			p.SetState(51)
			p.Defn()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefnContext is an interface to support dynamic dispatch.
type IDefnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefnContext differentiates from other interfaces.
	IsDefnContext()
}

type DefnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefnContext() *DefnContext {
	var p = new(DefnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_defn
	return p
}

func (*DefnContext) IsDefnContext() {}

func NewDefnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefnContext {
	var p = new(DefnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_defn

	return p
}

func (s *DefnContext) GetParser() antlr.Parser { return s.parser }

func (s *DefnContext) StateDefn() IStateDefnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateDefnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateDefnContext)
}

func (s *DefnContext) ClassDefn() IClassDefnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDefnContext)
}

func (s *DefnContext) NbhdDefn() INbhdDefnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INbhdDefnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INbhdDefnContext)
}

func (s *DefnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterDefn(s)
	}
}

func (s *DefnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitDefn(s)
	}
}

func (p *alpacaParser) Defn() (localctx IDefnContext) {
	this := p
	_ = this

	localctx = NewDefnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, alpacaParserRULE_defn)

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
	case alpacaParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.StateDefn()
		}

	case alpacaParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.ClassDefn()
		}

	case alpacaParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.NbhdDefn()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStateDefnContext is an interface to support dynamic dispatch.
type IStateDefnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateDefnContext differentiates from other interfaces.
	IsStateDefnContext()
}

type StateDefnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateDefnContext() *StateDefnContext {
	var p = new(StateDefnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_stateDefn
	return p
}

func (*StateDefnContext) IsStateDefnContext() {}

func NewStateDefnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDefnContext {
	var p = new(StateDefnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_stateDefn

	return p
}

func (s *StateDefnContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDefnContext) StateID() IStateIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateIDContext)
}

func (s *StateDefnContext) QUOTEDCHAR() antlr.TerminalNode {
	return s.GetToken(alpacaParserQUOTEDCHAR, 0)
}

func (s *StateDefnContext) AllMembershipDecl() []IMembershipDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMembershipDeclContext)(nil)).Elem())
	var tst = make([]IMembershipDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMembershipDeclContext)
		}
	}

	return tst
}

func (s *StateDefnContext) MembershipDecl(i int) IMembershipDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembershipDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMembershipDeclContext)
}

func (s *StateDefnContext) Rules() IRulesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *StateDefnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDefnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateDefnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterStateDefn(s)
	}
}

func (s *StateDefnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitStateDefn(s)
	}
}

func (p *alpacaParser) StateDefn() (localctx IStateDefnContext) {
	this := p
	_ = this

	localctx = NewStateDefnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, alpacaParserRULE_stateDefn)
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
		p.Match(alpacaParserT__3)
	}
	{
		p.SetState(63)
		p.StateID()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alpacaParserQUOTEDCHAR {
		{
			p.SetState(64)
			p.Match(alpacaParserQUOTEDCHAR)
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alpacaParserT__6 {
		{
			p.SetState(67)
			p.MembershipDecl()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alpacaParserT__8 {
		{
			p.SetState(73)
			p.Rules()
		}

	}

	return localctx
}

// IClassDefnContext is an interface to support dynamic dispatch.
type IClassDefnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDefnContext differentiates from other interfaces.
	IsClassDefnContext()
}

type ClassDefnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefnContext() *ClassDefnContext {
	var p = new(ClassDefnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_classDefn
	return p
}

func (*ClassDefnContext) IsClassDefnContext() {}

func NewClassDefnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefnContext {
	var p = new(ClassDefnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_classDefn

	return p
}

func (s *ClassDefnContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefnContext) ClassID() IClassIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassIDContext)
}

func (s *ClassDefnContext) AllMembershipDecl() []IMembershipDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMembershipDeclContext)(nil)).Elem())
	var tst = make([]IMembershipDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMembershipDeclContext)
		}
	}

	return tst
}

func (s *ClassDefnContext) MembershipDecl(i int) IMembershipDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMembershipDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMembershipDeclContext)
}

func (s *ClassDefnContext) Rules() IRulesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *ClassDefnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterClassDefn(s)
	}
}

func (s *ClassDefnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitClassDefn(s)
	}
}

func (p *alpacaParser) ClassDefn() (localctx IClassDefnContext) {
	this := p
	_ = this

	localctx = NewClassDefnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, alpacaParserRULE_classDefn)
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
		p.Match(alpacaParserT__4)
	}
	{
		p.SetState(77)
		p.ClassID()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alpacaParserT__6 {
		{
			p.SetState(78)
			p.MembershipDecl()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alpacaParserT__8 {
		{
			p.SetState(84)
			p.Rules()
		}

	}

	return localctx
}

// INbhdDefnContext is an interface to support dynamic dispatch.
type INbhdDefnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNbhdDefnContext differentiates from other interfaces.
	IsNbhdDefnContext()
}

type NbhdDefnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNbhdDefnContext() *NbhdDefnContext {
	var p = new(NbhdDefnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_nbhdDefn
	return p
}

func (*NbhdDefnContext) IsNbhdDefnContext() {}

func NewNbhdDefnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NbhdDefnContext {
	var p = new(NbhdDefnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_nbhdDefn

	return p
}

func (s *NbhdDefnContext) GetParser() antlr.Parser { return s.parser }

func (s *NbhdDefnContext) NbhdID() INbhdIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INbhdIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INbhdIDContext)
}

func (s *NbhdDefnContext) Neigbourhood() INeigbourhoodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeigbourhoodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INeigbourhoodContext)
}

func (s *NbhdDefnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NbhdDefnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NbhdDefnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterNbhdDefn(s)
	}
}

func (s *NbhdDefnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitNbhdDefn(s)
	}
}

func (p *alpacaParser) NbhdDefn() (localctx INbhdDefnContext) {
	this := p
	_ = this

	localctx = NewNbhdDefnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, alpacaParserRULE_nbhdDefn)

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
		p.Match(alpacaParserT__5)
	}
	{
		p.SetState(88)
		p.NbhdID()
	}
	{
		p.SetState(89)
		p.Neigbourhood()
	}

	return localctx
}

// IClassIDContext is an interface to support dynamic dispatch.
type IClassIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassIDContext differentiates from other interfaces.
	IsClassIDContext()
}

type ClassIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassIDContext() *ClassIDContext {
	var p = new(ClassIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_classID
	return p
}

func (*ClassIDContext) IsClassIDContext() {}

func NewClassIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassIDContext {
	var p = new(ClassIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_classID

	return p
}

func (s *ClassIDContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassIDContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ClassIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterClassID(s)
	}
}

func (s *ClassIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitClassID(s)
	}
}

func (p *alpacaParser) ClassID() (localctx IClassIDContext) {
	this := p
	_ = this

	localctx = NewClassIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, alpacaParserRULE_classID)

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
		p.Identifier()
	}

	return localctx
}

// IStateIDContext is an interface to support dynamic dispatch.
type IStateIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateIDContext differentiates from other interfaces.
	IsStateIDContext()
}

type StateIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateIDContext() *StateIDContext {
	var p = new(StateIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_stateID
	return p
}

func (*StateIDContext) IsStateIDContext() {}

func NewStateIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateIDContext {
	var p = new(StateIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_stateID

	return p
}

func (s *StateIDContext) GetParser() antlr.Parser { return s.parser }

func (s *StateIDContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *StateIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterStateID(s)
	}
}

func (s *StateIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitStateID(s)
	}
}

func (p *alpacaParser) StateID() (localctx IStateIDContext) {
	this := p
	_ = this

	localctx = NewStateIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, alpacaParserRULE_stateID)

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
		p.SetState(93)
		p.Identifier()
	}

	return localctx
}

// INbhdIDContext is an interface to support dynamic dispatch.
type INbhdIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNbhdIDContext differentiates from other interfaces.
	IsNbhdIDContext()
}

type NbhdIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNbhdIDContext() *NbhdIDContext {
	var p = new(NbhdIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_nbhdID
	return p
}

func (*NbhdIDContext) IsNbhdIDContext() {}

func NewNbhdIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NbhdIDContext {
	var p = new(NbhdIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_nbhdID

	return p
}

func (s *NbhdIDContext) GetParser() antlr.Parser { return s.parser }

func (s *NbhdIDContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NbhdIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NbhdIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NbhdIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterNbhdID(s)
	}
}

func (s *NbhdIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitNbhdID(s)
	}
}

func (p *alpacaParser) NbhdID() (localctx INbhdIDContext) {
	this := p
	_ = this

	localctx = NewNbhdIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, alpacaParserRULE_nbhdID)

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
		p.SetState(95)
		p.Identifier()
	}

	return localctx
}

// IMembershipDeclContext is an interface to support dynamic dispatch.
type IMembershipDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMembershipDeclContext differentiates from other interfaces.
	IsMembershipDeclContext()
}

type MembershipDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembershipDeclContext() *MembershipDeclContext {
	var p = new(MembershipDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_membershipDecl
	return p
}

func (*MembershipDeclContext) IsMembershipDeclContext() {}

func NewMembershipDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembershipDeclContext {
	var p = new(MembershipDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_membershipDecl

	return p
}

func (s *MembershipDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MembershipDeclContext) ClassRef() IClassRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassRefContext)
}

func (s *MembershipDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembershipDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterMembershipDecl(s)
	}
}

func (s *MembershipDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitMembershipDecl(s)
	}
}

func (p *alpacaParser) MembershipDecl() (localctx IMembershipDeclContext) {
	this := p
	_ = this

	localctx = NewMembershipDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, alpacaParserRULE_membershipDecl)

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
		p.ClassRef()
	}

	return localctx
}

// IClassRefContext is an interface to support dynamic dispatch.
type IClassRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassRefContext differentiates from other interfaces.
	IsClassRefContext()
}

type ClassRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassRefContext() *ClassRefContext {
	var p = new(ClassRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_classRef
	return p
}

func (*ClassRefContext) IsClassRefContext() {}

func NewClassRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassRefContext {
	var p = new(ClassRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_classRef

	return p
}

func (s *ClassRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassRefContext) ClassID() IClassIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassIDContext)
}

func (s *ClassRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterClassRef(s)
	}
}

func (s *ClassRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitClassRef(s)
	}
}

func (p *alpacaParser) ClassRef() (localctx IClassRefContext) {
	this := p
	_ = this

	localctx = NewClassRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, alpacaParserRULE_classRef)

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
		p.SetState(99)
		p.Match(alpacaParserT__6)
	}
	{
		p.SetState(100)
		p.ClassID()
	}

	return localctx
}

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) AllRule_() []IRule_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRule_Context)(nil)).Elem())
	var tst = make([]IRule_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRule_Context)
		}
	}

	return tst
}

func (s *RulesContext) Rule_(i int) IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitRules(s)
	}
}

func (p *alpacaParser) Rules() (localctx IRulesContext) {
	this := p
	_ = this

	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, alpacaParserRULE_rules)
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
		p.SetState(102)
		p.Rule_()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alpacaParserT__7 {
		{
			p.SetState(103)
			p.Match(alpacaParserT__7)
		}
		{
			p.SetState(104)
			p.Rule_()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = alpacaParserRULE_rule_
	return p
}

func (*Rule_Context) IsRule_Context() {}

func NewRule_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_Context {
	var p = new(Rule_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_rule_

	return p
}

func (s *Rule_Context) GetParser() antlr.Parser { return s.parser }

func (s *Rule_Context) StateRef() IStateRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateRefContext)
}

func (s *Rule_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Rule_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterRule_(s)
	}
}

func (s *Rule_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitRule_(s)
	}
}

func (p *alpacaParser) Rule_() (localctx IRule_Context) {
	this := p
	_ = this

	localctx = NewRule_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, alpacaParserRULE_rule_)
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
		p.SetState(110)
		p.Match(alpacaParserT__8)
	}
	{
		p.SetState(111)
		p.StateRef()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alpacaParserT__9 {
		{
			p.SetState(112)
			p.Match(alpacaParserT__9)
		}
		{
			p.SetState(113)
			p.Expression()
		}

	}

	return localctx
}

// IStateRefContext is an interface to support dynamic dispatch.
type IStateRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateRefContext differentiates from other interfaces.
	IsStateRefContext()
}

type StateRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateRefContext() *StateRefContext {
	var p = new(StateRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_stateRef
	return p
}

func (*StateRefContext) IsStateRefContext() {}

func NewStateRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateRefContext {
	var p = new(StateRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_stateRef

	return p
}

func (s *StateRefContext) GetParser() antlr.Parser { return s.parser }

func (s *StateRefContext) StateID() IStateIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateIDContext)
}

func (s *StateRefContext) Arrowchain() IArrowchainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowchainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrowchainContext)
}

func (s *StateRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterStateRef(s)
	}
}

func (s *StateRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitStateRef(s)
	}
}

func (p *alpacaParser) StateRef() (localctx IStateRefContext) {
	this := p
	_ = this

	localctx = NewStateRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, alpacaParserRULE_stateRef)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alpacaParserALPHA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.StateID()
		}

	case alpacaParserARROW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Arrowchain()
		}

	case alpacaParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(alpacaParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = alpacaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_expression

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
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *alpacaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, alpacaParserRULE_expression)
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
		p.SetState(121)
		p.Term()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alpacaParserT__11)|(1<<alpacaParserT__12)|(1<<alpacaParserT__13))) != 0 {
		{
			p.SetState(122)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alpacaParserT__11)|(1<<alpacaParserT__12)|(1<<alpacaParserT__13))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.Term()
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = alpacaParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AdjacencyPred() IAdjacencyPredContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdjacencyPredContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdjacencyPredContext)
}

func (s *TermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) BoolPrimitive() IBoolPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolPrimitiveContext)
}

func (s *TermContext) RelationalPred() IRelationalPredContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalPredContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalPredContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *alpacaParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, alpacaParserRULE_term)

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
	case alpacaParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.AdjacencyPred()
		}

	case alpacaParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(alpacaParserT__14)
		}
		{
			p.SetState(131)
			p.Expression()
		}
		{
			p.SetState(132)
			p.Match(alpacaParserT__15)
		}

	case alpacaParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(alpacaParserT__16)
		}
		{
			p.SetState(135)
			p.Term()
		}

	case alpacaParserT__19, alpacaParserT__20, alpacaParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.BoolPrimitive()
		}

	case alpacaParserT__10, alpacaParserALPHA, alpacaParserARROW:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
			p.RelationalPred()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelationalPredContext is an interface to support dynamic dispatch.
type IRelationalPredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalPredContext differentiates from other interfaces.
	IsRelationalPredContext()
}

type RelationalPredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalPredContext() *RelationalPredContext {
	var p = new(RelationalPredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_relationalPred
	return p
}

func (*RelationalPredContext) IsRelationalPredContext() {}

func NewRelationalPredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalPredContext {
	var p = new(RelationalPredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_relationalPred

	return p
}

func (s *RelationalPredContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalPredContext) AllStateRef() []IStateRefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateRefContext)(nil)).Elem())
	var tst = make([]IStateRefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateRefContext)
		}
	}

	return tst
}

func (s *RelationalPredContext) StateRef(i int) IStateRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateRefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateRefContext)
}

func (s *RelationalPredContext) ClassRef() IClassRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassRefContext)
}

func (s *RelationalPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalPredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalPredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterRelationalPred(s)
	}
}

func (s *RelationalPredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitRelationalPred(s)
	}
}

func (p *alpacaParser) RelationalPred() (localctx IRelationalPredContext) {
	this := p
	_ = this

	localctx = NewRelationalPredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, alpacaParserRULE_relationalPred)
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
		p.SetState(140)
		p.StateRef()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alpacaParserT__10, alpacaParserT__17, alpacaParserALPHA, alpacaParserARROW:
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == alpacaParserT__17 {
			{
				p.SetState(141)
				p.Match(alpacaParserT__17)
			}

		}
		{
			p.SetState(144)
			p.StateRef()
		}

	case alpacaParserT__6:
		{
			p.SetState(145)
			p.ClassRef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAdjacencyPredContext is an interface to support dynamic dispatch.
type IAdjacencyPredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdjacencyPredContext differentiates from other interfaces.
	IsAdjacencyPredContext()
}

type AdjacencyPredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdjacencyPredContext() *AdjacencyPredContext {
	var p = new(AdjacencyPredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_adjacencyPred
	return p
}

func (*AdjacencyPredContext) IsAdjacencyPredContext() {}

func NewAdjacencyPredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdjacencyPredContext {
	var p = new(AdjacencyPredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_adjacencyPred

	return p
}

func (s *AdjacencyPredContext) GetParser() antlr.Parser { return s.parser }

func (s *AdjacencyPredContext) Naturalnumber() INaturalnumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INaturalnumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INaturalnumberContext)
}

func (s *AdjacencyPredContext) StateRef() IStateRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateRefContext)
}

func (s *AdjacencyPredContext) ClassRef() IClassRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassRefContext)
}

func (s *AdjacencyPredContext) Neigbourhood() INeigbourhoodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeigbourhoodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INeigbourhoodContext)
}

func (s *AdjacencyPredContext) NbhdID() INbhdIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INbhdIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INbhdIDContext)
}

func (s *AdjacencyPredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdjacencyPredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdjacencyPredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterAdjacencyPred(s)
	}
}

func (s *AdjacencyPredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitAdjacencyPred(s)
	}
}

func (p *alpacaParser) AdjacencyPred() (localctx IAdjacencyPredContext) {
	this := p
	_ = this

	localctx = NewAdjacencyPredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, alpacaParserRULE_adjacencyPred)
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
		p.SetState(148)
		p.Naturalnumber()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alpacaParserT__18 {
		{
			p.SetState(149)
			p.Match(alpacaParserT__18)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case alpacaParserT__14:
			{
				p.SetState(150)
				p.Neigbourhood()
			}

		case alpacaParserALPHA:
			{
				p.SetState(151)
				p.NbhdID()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alpacaParserT__10, alpacaParserALPHA, alpacaParserARROW:
		{
			p.SetState(156)
			p.StateRef()
		}

	case alpacaParserT__6:
		{
			p.SetState(157)
			p.ClassRef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolPrimitiveContext is an interface to support dynamic dispatch.
type IBoolPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolPrimitiveContext differentiates from other interfaces.
	IsBoolPrimitiveContext()
}

type BoolPrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolPrimitiveContext() *BoolPrimitiveContext {
	var p = new(BoolPrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_boolPrimitive
	return p
}

func (*BoolPrimitiveContext) IsBoolPrimitiveContext() {}

func NewBoolPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolPrimitiveContext {
	var p = new(BoolPrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_boolPrimitive

	return p
}

func (s *BoolPrimitiveContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolPrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolPrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolPrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterBoolPrimitive(s)
	}
}

func (s *BoolPrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitBoolPrimitive(s)
	}
}

func (p *alpacaParser) BoolPrimitive() (localctx IBoolPrimitiveContext) {
	this := p
	_ = this

	localctx = NewBoolPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, alpacaParserRULE_boolPrimitive)
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
		p.SetState(160)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alpacaParserT__19)|(1<<alpacaParserT__20)|(1<<alpacaParserT__21))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INeigbourhoodContext is an interface to support dynamic dispatch.
type INeigbourhoodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeigbourhoodContext differentiates from other interfaces.
	IsNeigbourhoodContext()
}

type NeigbourhoodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeigbourhoodContext() *NeigbourhoodContext {
	var p = new(NeigbourhoodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_neigbourhood
	return p
}

func (*NeigbourhoodContext) IsNeigbourhoodContext() {}

func NewNeigbourhoodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NeigbourhoodContext {
	var p = new(NeigbourhoodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_neigbourhood

	return p
}

func (s *NeigbourhoodContext) GetParser() antlr.Parser { return s.parser }

func (s *NeigbourhoodContext) AllArrowchain() []IArrowchainContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrowchainContext)(nil)).Elem())
	var tst = make([]IArrowchainContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrowchainContext)
		}
	}

	return tst
}

func (s *NeigbourhoodContext) Arrowchain(i int) IArrowchainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowchainContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrowchainContext)
}

func (s *NeigbourhoodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeigbourhoodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NeigbourhoodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterNeigbourhood(s)
	}
}

func (s *NeigbourhoodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitNeigbourhood(s)
	}
}

func (p *alpacaParser) Neigbourhood() (localctx INeigbourhoodContext) {
	this := p
	_ = this

	localctx = NewNeigbourhoodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, alpacaParserRULE_neigbourhood)
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
		p.SetState(162)
		p.Match(alpacaParserT__14)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alpacaParserARROW {
		{
			p.SetState(163)
			p.Arrowchain()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(alpacaParserT__15)
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
	p.RuleIndex = alpacaParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllALPHA() []antlr.TerminalNode {
	return s.GetTokens(alpacaParserALPHA)
}

func (s *IdentifierContext) ALPHA(i int) antlr.TerminalNode {
	return s.GetToken(alpacaParserALPHA, i)
}

func (s *IdentifierContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(alpacaParserDIGIT)
}

func (s *IdentifierContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(alpacaParserDIGIT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *alpacaParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, alpacaParserRULE_identifier)
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
		p.SetState(171)
		p.Match(alpacaParserALPHA)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				_la = p.GetTokenStream().LA(1)

				if !(_la == alpacaParserALPHA || _la == alpacaParserDIGIT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// INaturalnumberContext is an interface to support dynamic dispatch.
type INaturalnumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNaturalnumberContext differentiates from other interfaces.
	IsNaturalnumberContext()
}

type NaturalnumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNaturalnumberContext() *NaturalnumberContext {
	var p = new(NaturalnumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_naturalnumber
	return p
}

func (*NaturalnumberContext) IsNaturalnumberContext() {}

func NewNaturalnumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NaturalnumberContext {
	var p = new(NaturalnumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_naturalnumber

	return p
}

func (s *NaturalnumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NaturalnumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(alpacaParserDIGIT)
}

func (s *NaturalnumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(alpacaParserDIGIT, i)
}

func (s *NaturalnumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NaturalnumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NaturalnumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterNaturalnumber(s)
	}
}

func (s *NaturalnumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitNaturalnumber(s)
	}
}

func (p *alpacaParser) Naturalnumber() (localctx INaturalnumberContext) {
	this := p
	_ = this

	localctx = NewNaturalnumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, alpacaParserRULE_naturalnumber)
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
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == alpacaParserDIGIT {
		{
			p.SetState(178)
			p.Match(alpacaParserDIGIT)
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrowchainContext is an interface to support dynamic dispatch.
type IArrowchainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrowchainContext differentiates from other interfaces.
	IsArrowchainContext()
}

type ArrowchainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowchainContext() *ArrowchainContext {
	var p = new(ArrowchainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alpacaParserRULE_arrowchain
	return p
}

func (*ArrowchainContext) IsArrowchainContext() {}

func NewArrowchainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowchainContext {
	var p = new(ArrowchainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alpacaParserRULE_arrowchain

	return p
}

func (s *ArrowchainContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowchainContext) AllARROW() []antlr.TerminalNode {
	return s.GetTokens(alpacaParserARROW)
}

func (s *ArrowchainContext) ARROW(i int) antlr.TerminalNode {
	return s.GetToken(alpacaParserARROW, i)
}

func (s *ArrowchainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowchainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowchainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.EnterArrowchain(s)
	}
}

func (s *ArrowchainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alpacaListener); ok {
		listenerT.ExitArrowchain(s)
	}
}

func (p *alpacaParser) Arrowchain() (localctx IArrowchainContext) {
	this := p
	_ = this

	localctx = NewArrowchainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, alpacaParserRULE_arrowchain)

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
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(183)
				p.Match(alpacaParserARROW)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}
