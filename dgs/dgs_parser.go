// Code generated from DGSParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgs // DGSParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 185,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 70, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 89, 10, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 7, 14, 110, 10, 14,
	12, 14, 14, 14, 113, 11, 14, 3, 15, 5, 15, 116, 10, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 7, 15, 123, 10, 15, 12, 15, 14, 15, 126, 11, 15, 5, 15,
	128, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 137,
	10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 143, 10, 17, 12, 17, 14, 17,
	146, 11, 17, 5, 17, 148, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 7, 18, 156, 10, 18, 12, 18, 14, 18, 159, 11, 18, 5, 18, 161, 10, 18,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 178, 10, 22, 12, 22, 14, 22, 181,
	11, 22, 5, 22, 183, 10, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 6, 4, 2, 27, 27,
	33, 33, 3, 2, 15, 16, 3, 2, 23, 24, 3, 2, 25, 26, 2, 192, 2, 44, 3, 2,
	2, 2, 4, 53, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2, 10, 77,
	3, 2, 2, 2, 12, 81, 3, 2, 2, 2, 14, 84, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2,
	18, 97, 3, 2, 2, 2, 20, 100, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 106,
	3, 2, 2, 2, 26, 111, 3, 2, 2, 2, 28, 115, 3, 2, 2, 2, 30, 136, 3, 2, 2,
	2, 32, 138, 3, 2, 2, 2, 34, 151, 3, 2, 2, 2, 36, 164, 3, 2, 2, 2, 38, 168,
	3, 2, 2, 2, 40, 170, 3, 2, 2, 2, 42, 182, 3, 2, 2, 2, 44, 50, 5, 4, 3,
	2, 45, 49, 5, 6, 4, 2, 46, 49, 7, 33, 2, 2, 47, 49, 7, 27, 2, 2, 48, 45,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 3, 3, 2, 2, 2, 52, 50, 3, 2,
	2, 2, 53, 54, 7, 3, 2, 2, 54, 55, 7, 27, 2, 2, 55, 56, 5, 42, 22, 2, 56,
	57, 7, 13, 2, 2, 57, 58, 7, 13, 2, 2, 58, 59, 7, 27, 2, 2, 59, 5, 3, 2,
	2, 2, 60, 70, 5, 8, 5, 2, 61, 70, 5, 10, 6, 2, 62, 70, 5, 12, 7, 2, 63,
	70, 5, 14, 8, 2, 64, 70, 5, 16, 9, 2, 65, 70, 5, 18, 10, 2, 66, 70, 5,
	20, 11, 2, 67, 70, 5, 22, 12, 2, 68, 70, 5, 24, 13, 2, 69, 60, 3, 2, 2,
	2, 69, 61, 3, 2, 2, 2, 69, 62, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 69, 64,
	3, 2, 2, 2, 69, 65, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	69, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 9, 2, 2, 2, 72, 7, 3, 2,
	2, 2, 73, 74, 7, 4, 2, 2, 74, 75, 5, 42, 22, 2, 75, 76, 5, 26, 14, 2, 76,
	9, 3, 2, 2, 2, 77, 78, 7, 5, 2, 2, 78, 79, 5, 42, 22, 2, 79, 80, 5, 26,
	14, 2, 80, 11, 3, 2, 2, 2, 81, 82, 7, 6, 2, 2, 82, 83, 5, 42, 22, 2, 83,
	13, 3, 2, 2, 2, 84, 85, 7, 7, 2, 2, 85, 86, 5, 42, 22, 2, 86, 88, 5, 42,
	22, 2, 87, 89, 5, 38, 20, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89,
	90, 3, 2, 2, 2, 90, 91, 5, 42, 22, 2, 91, 92, 5, 26, 14, 2, 92, 15, 3,
	2, 2, 2, 93, 94, 7, 8, 2, 2, 94, 95, 5, 42, 22, 2, 95, 96, 5, 26, 14, 2,
	96, 17, 3, 2, 2, 2, 97, 98, 7, 9, 2, 2, 98, 99, 5, 42, 22, 2, 99, 19, 3,
	2, 2, 2, 100, 101, 7, 10, 2, 2, 101, 102, 5, 26, 14, 2, 102, 21, 3, 2,
	2, 2, 103, 104, 7, 11, 2, 2, 104, 105, 7, 14, 2, 2, 105, 23, 3, 2, 2, 2,
	106, 107, 7, 12, 2, 2, 107, 25, 3, 2, 2, 2, 108, 110, 5, 28, 15, 2, 109,
	108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 27, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 116, 9, 3,
	2, 2, 115, 114, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2,
	117, 127, 5, 42, 22, 2, 118, 119, 5, 40, 21, 2, 119, 124, 5, 30, 16, 2,
	120, 121, 7, 17, 2, 2, 121, 123, 5, 30, 16, 2, 122, 120, 3, 2, 2, 2, 123,
	126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 128,
	3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 118, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 29, 3, 2, 2, 2, 129, 137, 7, 29, 2, 2, 130, 137, 7, 13, 2, 2,
	131, 137, 7, 14, 2, 2, 132, 137, 7, 30, 2, 2, 133, 137, 5, 32, 17, 2, 134,
	137, 5, 34, 18, 2, 135, 137, 5, 42, 22, 2, 136, 129, 3, 2, 2, 2, 136, 130,
	3, 2, 2, 2, 136, 131, 3, 2, 2, 2, 136, 132, 3, 2, 2, 2, 136, 133, 3, 2,
	2, 2, 136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 31, 3, 2, 2, 2,
	138, 147, 7, 18, 2, 2, 139, 144, 5, 30, 16, 2, 140, 141, 7, 17, 2, 2, 141,
	143, 5, 30, 16, 2, 142, 140, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2,
	2, 2, 147, 139, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 7, 19, 2, 2, 150, 33, 3, 2, 2, 2, 151, 160, 7, 20, 2, 2, 152,
	157, 5, 36, 19, 2, 153, 154, 7, 17, 2, 2, 154, 156, 5, 36, 19, 2, 155,
	153, 3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158,
	3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 152, 3, 2,
	2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 7, 21, 2, 2,
	163, 35, 3, 2, 2, 2, 164, 165, 5, 42, 22, 2, 165, 166, 5, 40, 21, 2, 166,
	167, 5, 30, 16, 2, 167, 37, 3, 2, 2, 2, 168, 169, 9, 4, 2, 2, 169, 39,
	3, 2, 2, 2, 170, 171, 9, 5, 2, 2, 171, 41, 3, 2, 2, 2, 172, 183, 7, 29,
	2, 2, 173, 183, 7, 13, 2, 2, 174, 179, 7, 28, 2, 2, 175, 176, 7, 22, 2,
	2, 176, 178, 7, 28, 2, 2, 177, 175, 3, 2, 2, 2, 178, 181, 3, 2, 2, 2, 179,
	177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179,
	3, 2, 2, 2, 182, 172, 3, 2, 2, 2, 182, 173, 3, 2, 2, 2, 182, 174, 3, 2,
	2, 2, 183, 43, 3, 2, 2, 2, 17, 48, 50, 69, 88, 111, 115, 124, 127, 136,
	144, 147, 157, 160, 179, 182,
}
var literalNames = []string{
	"", "", "'an'", "'cn'", "'dn'", "'ae'", "'ce'", "'de'", "'cg'", "'st'",
	"'cl'", "", "", "'+'", "'-'", "','", "'{'", "'}'", "'['", "']'", "'.'",
	"'<'", "'>'", "'='", "':'", "", "", "", "", "'#'",
}
var symbolicNames = []string{
	"", "MAGIC", "AN", "CN", "DN", "AE", "CE", "DE", "CG", "ST", "CL", "INT",
	"REAL", "PLUS", "MINUS", "COMMA", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
	"DOT", "LANGLE", "RANGLE", "EQUALS", "COLON", "EOL", "WORD", "STRING",
	"COLOR", "START_COMMENT", "WS", "COMMENT",
}

var ruleNames = []string{
	"dgs", "header", "event", "an", "cn", "dn", "ae", "ce", "de", "cg", "st",
	"cl", "attributes", "attribute", "value", "array_", "a_map", "mapping",
	"direction", "assign", "identifier",
}

type DGSParser struct {
	*antlr.BaseParser
}

// NewDGSParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DGSParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDGSParser(input antlr.TokenStream) *DGSParser {
	this := new(DGSParser)
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
	this.GrammarFileName = "DGSParser.g4"

	return this
}

// DGSParser tokens.
const (
	DGSParserEOF           = antlr.TokenEOF
	DGSParserMAGIC         = 1
	DGSParserAN            = 2
	DGSParserCN            = 3
	DGSParserDN            = 4
	DGSParserAE            = 5
	DGSParserCE            = 6
	DGSParserDE            = 7
	DGSParserCG            = 8
	DGSParserST            = 9
	DGSParserCL            = 10
	DGSParserINT           = 11
	DGSParserREAL          = 12
	DGSParserPLUS          = 13
	DGSParserMINUS         = 14
	DGSParserCOMMA         = 15
	DGSParserLBRACE        = 16
	DGSParserRBRACE        = 17
	DGSParserLBRACK        = 18
	DGSParserRBRACK        = 19
	DGSParserDOT           = 20
	DGSParserLANGLE        = 21
	DGSParserRANGLE        = 22
	DGSParserEQUALS        = 23
	DGSParserCOLON         = 24
	DGSParserEOL           = 25
	DGSParserWORD          = 26
	DGSParserSTRING        = 27
	DGSParserCOLOR         = 28
	DGSParserSTART_COMMENT = 29
	DGSParserWS            = 30
	DGSParserCOMMENT       = 31
)

// DGSParser rules.
const (
	DGSParserRULE_dgs        = 0
	DGSParserRULE_header     = 1
	DGSParserRULE_event      = 2
	DGSParserRULE_an         = 3
	DGSParserRULE_cn         = 4
	DGSParserRULE_dn         = 5
	DGSParserRULE_ae         = 6
	DGSParserRULE_ce         = 7
	DGSParserRULE_de         = 8
	DGSParserRULE_cg         = 9
	DGSParserRULE_st         = 10
	DGSParserRULE_cl         = 11
	DGSParserRULE_attributes = 12
	DGSParserRULE_attribute  = 13
	DGSParserRULE_value      = 14
	DGSParserRULE_array_     = 15
	DGSParserRULE_a_map      = 16
	DGSParserRULE_mapping    = 17
	DGSParserRULE_direction  = 18
	DGSParserRULE_assign     = 19
	DGSParserRULE_identifier = 20
)

// IDgsContext is an interface to support dynamic dispatch.
type IDgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDgsContext differentiates from other interfaces.
	IsDgsContext()
}

type DgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDgsContext() *DgsContext {
	var p = new(DgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_dgs
	return p
}

func (*DgsContext) IsDgsContext() {}

func NewDgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DgsContext {
	var p = new(DgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_dgs

	return p
}

func (s *DgsContext) GetParser() antlr.Parser { return s.parser }

func (s *DgsContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DgsContext) AllEvent() []IEventContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEventContext)(nil)).Elem())
	var tst = make([]IEventContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEventContext)
		}
	}

	return tst
}

func (s *DgsContext) Event(i int) IEventContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *DgsContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(DGSParserCOMMENT)
}

func (s *DgsContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserCOMMENT, i)
}

func (s *DgsContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(DGSParserEOL)
}

func (s *DgsContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserEOL, i)
}

func (s *DgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterDgs(s)
	}
}

func (s *DgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitDgs(s)
	}
}

func (p *DGSParser) Dgs() (localctx IDgsContext) {
	this := p
	_ = this

	localctx = NewDgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DGSParserRULE_dgs)
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
		p.SetState(42)
		p.Header()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGSParserAN)|(1<<DGSParserCN)|(1<<DGSParserDN)|(1<<DGSParserAE)|(1<<DGSParserCE)|(1<<DGSParserDE)|(1<<DGSParserCG)|(1<<DGSParserST)|(1<<DGSParserCL)|(1<<DGSParserEOL)|(1<<DGSParserCOMMENT))) != 0 {
		p.SetState(46)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DGSParserAN, DGSParserCN, DGSParserDN, DGSParserAE, DGSParserCE, DGSParserDE, DGSParserCG, DGSParserST, DGSParserCL:
			{
				p.SetState(43)
				p.Event()
			}

		case DGSParserCOMMENT:
			{
				p.SetState(44)
				p.Match(DGSParserCOMMENT)
			}

		case DGSParserEOL:
			{
				p.SetState(45)
				p.Match(DGSParserEOL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) MAGIC() antlr.TerminalNode {
	return s.GetToken(DGSParserMAGIC, 0)
}

func (s *HeaderContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(DGSParserEOL)
}

func (s *HeaderContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserEOL, i)
}

func (s *HeaderContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *HeaderContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(DGSParserINT)
}

func (s *HeaderContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserINT, i)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *DGSParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DGSParserRULE_header)

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
		p.Match(DGSParserMAGIC)
	}
	{
		p.SetState(52)
		p.Match(DGSParserEOL)
	}
	{
		p.SetState(53)
		p.Identifier()
	}
	{
		p.SetState(54)
		p.Match(DGSParserINT)
	}
	{
		p.SetState(55)
		p.Match(DGSParserINT)
	}
	{
		p.SetState(56)
		p.Match(DGSParserEOL)
	}

	return localctx
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_event
	return p
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }

func (s *EventContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DGSParserCOMMENT, 0)
}

func (s *EventContext) EOL() antlr.TerminalNode {
	return s.GetToken(DGSParserEOL, 0)
}

func (s *EventContext) An() IAnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnContext)
}

func (s *EventContext) Cn() ICnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICnContext)
}

func (s *EventContext) Dn() IDnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDnContext)
}

func (s *EventContext) Ae() IAeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAeContext)
}

func (s *EventContext) Ce() ICeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICeContext)
}

func (s *EventContext) De() IDeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeContext)
}

func (s *EventContext) Cg() ICgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICgContext)
}

func (s *EventContext) St() IStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStContext)
}

func (s *EventContext) Cl() IClContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClContext)
}

func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterEvent(s)
	}
}

func (s *EventContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitEvent(s)
	}
}

func (p *DGSParser) Event() (localctx IEventContext) {
	this := p
	_ = this

	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DGSParserRULE_event)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DGSParserAN:
		{
			p.SetState(58)
			p.An()
		}

	case DGSParserCN:
		{
			p.SetState(59)
			p.Cn()
		}

	case DGSParserDN:
		{
			p.SetState(60)
			p.Dn()
		}

	case DGSParserAE:
		{
			p.SetState(61)
			p.Ae()
		}

	case DGSParserCE:
		{
			p.SetState(62)
			p.Ce()
		}

	case DGSParserDE:
		{
			p.SetState(63)
			p.De()
		}

	case DGSParserCG:
		{
			p.SetState(64)
			p.Cg()
		}

	case DGSParserST:
		{
			p.SetState(65)
			p.St()
		}

	case DGSParserCL:
		{
			p.SetState(66)
			p.Cl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DGSParserEOL || _la == DGSParserCOMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAnContext is an interface to support dynamic dispatch.
type IAnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnContext differentiates from other interfaces.
	IsAnContext()
}

type AnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnContext() *AnContext {
	var p = new(AnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_an
	return p
}

func (*AnContext) IsAnContext() {}

func NewAnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnContext {
	var p = new(AnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_an

	return p
}

func (s *AnContext) GetParser() antlr.Parser { return s.parser }

func (s *AnContext) AN() antlr.TerminalNode {
	return s.GetToken(DGSParserAN, 0)
}

func (s *AnContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AnContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *AnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterAn(s)
	}
}

func (s *AnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitAn(s)
	}
}

func (p *DGSParser) An() (localctx IAnContext) {
	this := p
	_ = this

	localctx = NewAnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DGSParserRULE_an)

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
		p.SetState(71)
		p.Match(DGSParserAN)
	}
	{
		p.SetState(72)
		p.Identifier()
	}
	{
		p.SetState(73)
		p.Attributes()
	}

	return localctx
}

// ICnContext is an interface to support dynamic dispatch.
type ICnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCnContext differentiates from other interfaces.
	IsCnContext()
}

type CnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCnContext() *CnContext {
	var p = new(CnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_cn
	return p
}

func (*CnContext) IsCnContext() {}

func NewCnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CnContext {
	var p = new(CnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_cn

	return p
}

func (s *CnContext) GetParser() antlr.Parser { return s.parser }

func (s *CnContext) CN() antlr.TerminalNode {
	return s.GetToken(DGSParserCN, 0)
}

func (s *CnContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CnContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *CnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterCn(s)
	}
}

func (s *CnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitCn(s)
	}
}

func (p *DGSParser) Cn() (localctx ICnContext) {
	this := p
	_ = this

	localctx = NewCnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DGSParserRULE_cn)

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
		p.Match(DGSParserCN)
	}
	{
		p.SetState(76)
		p.Identifier()
	}
	{
		p.SetState(77)
		p.Attributes()
	}

	return localctx
}

// IDnContext is an interface to support dynamic dispatch.
type IDnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDnContext differentiates from other interfaces.
	IsDnContext()
}

type DnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDnContext() *DnContext {
	var p = new(DnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_dn
	return p
}

func (*DnContext) IsDnContext() {}

func NewDnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DnContext {
	var p = new(DnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_dn

	return p
}

func (s *DnContext) GetParser() antlr.Parser { return s.parser }

func (s *DnContext) DN() antlr.TerminalNode {
	return s.GetToken(DGSParserDN, 0)
}

func (s *DnContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterDn(s)
	}
}

func (s *DnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitDn(s)
	}
}

func (p *DGSParser) Dn() (localctx IDnContext) {
	this := p
	_ = this

	localctx = NewDnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DGSParserRULE_dn)

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
		p.SetState(79)
		p.Match(DGSParserDN)
	}
	{
		p.SetState(80)
		p.Identifier()
	}

	return localctx
}

// IAeContext is an interface to support dynamic dispatch.
type IAeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAeContext differentiates from other interfaces.
	IsAeContext()
}

type AeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAeContext() *AeContext {
	var p = new(AeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_ae
	return p
}

func (*AeContext) IsAeContext() {}

func NewAeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AeContext {
	var p = new(AeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_ae

	return p
}

func (s *AeContext) GetParser() antlr.Parser { return s.parser }

func (s *AeContext) AE() antlr.TerminalNode {
	return s.GetToken(DGSParserAE, 0)
}

func (s *AeContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *AeContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AeContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *AeContext) Direction() IDirectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectionContext)
}

func (s *AeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterAe(s)
	}
}

func (s *AeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitAe(s)
	}
}

func (p *DGSParser) Ae() (localctx IAeContext) {
	this := p
	_ = this

	localctx = NewAeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DGSParserRULE_ae)
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
		p.Match(DGSParserAE)
	}
	{
		p.SetState(83)
		p.Identifier()
	}
	{
		p.SetState(84)
		p.Identifier()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGSParserLANGLE || _la == DGSParserRANGLE {
		{
			p.SetState(85)
			p.Direction()
		}

	}
	{
		p.SetState(88)
		p.Identifier()
	}
	{
		p.SetState(89)
		p.Attributes()
	}

	return localctx
}

// ICeContext is an interface to support dynamic dispatch.
type ICeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCeContext differentiates from other interfaces.
	IsCeContext()
}

type CeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCeContext() *CeContext {
	var p = new(CeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_ce
	return p
}

func (*CeContext) IsCeContext() {}

func NewCeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeContext {
	var p = new(CeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_ce

	return p
}

func (s *CeContext) GetParser() antlr.Parser { return s.parser }

func (s *CeContext) CE() antlr.TerminalNode {
	return s.GetToken(DGSParserCE, 0)
}

func (s *CeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CeContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *CeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterCe(s)
	}
}

func (s *CeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitCe(s)
	}
}

func (p *DGSParser) Ce() (localctx ICeContext) {
	this := p
	_ = this

	localctx = NewCeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DGSParserRULE_ce)

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
		p.Match(DGSParserCE)
	}
	{
		p.SetState(92)
		p.Identifier()
	}
	{
		p.SetState(93)
		p.Attributes()
	}

	return localctx
}

// IDeContext is an interface to support dynamic dispatch.
type IDeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeContext differentiates from other interfaces.
	IsDeContext()
}

type DeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeContext() *DeContext {
	var p = new(DeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_de
	return p
}

func (*DeContext) IsDeContext() {}

func NewDeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeContext {
	var p = new(DeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_de

	return p
}

func (s *DeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeContext) DE() antlr.TerminalNode {
	return s.GetToken(DGSParserDE, 0)
}

func (s *DeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterDe(s)
	}
}

func (s *DeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitDe(s)
	}
}

func (p *DGSParser) De() (localctx IDeContext) {
	this := p
	_ = this

	localctx = NewDeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DGSParserRULE_de)

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
		p.Match(DGSParserDE)
	}
	{
		p.SetState(96)
		p.Identifier()
	}

	return localctx
}

// ICgContext is an interface to support dynamic dispatch.
type ICgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCgContext differentiates from other interfaces.
	IsCgContext()
}

type CgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCgContext() *CgContext {
	var p = new(CgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_cg
	return p
}

func (*CgContext) IsCgContext() {}

func NewCgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CgContext {
	var p = new(CgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_cg

	return p
}

func (s *CgContext) GetParser() antlr.Parser { return s.parser }

func (s *CgContext) CG() antlr.TerminalNode {
	return s.GetToken(DGSParserCG, 0)
}

func (s *CgContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *CgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterCg(s)
	}
}

func (s *CgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitCg(s)
	}
}

func (p *DGSParser) Cg() (localctx ICgContext) {
	this := p
	_ = this

	localctx = NewCgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DGSParserRULE_cg)

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
		p.SetState(98)
		p.Match(DGSParserCG)
	}
	{
		p.SetState(99)
		p.Attributes()
	}

	return localctx
}

// IStContext is an interface to support dynamic dispatch.
type IStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStContext differentiates from other interfaces.
	IsStContext()
}

type StContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStContext() *StContext {
	var p = new(StContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_st
	return p
}

func (*StContext) IsStContext() {}

func NewStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StContext {
	var p = new(StContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_st

	return p
}

func (s *StContext) GetParser() antlr.Parser { return s.parser }

func (s *StContext) ST() antlr.TerminalNode {
	return s.GetToken(DGSParserST, 0)
}

func (s *StContext) REAL() antlr.TerminalNode {
	return s.GetToken(DGSParserREAL, 0)
}

func (s *StContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterSt(s)
	}
}

func (s *StContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitSt(s)
	}
}

func (p *DGSParser) St() (localctx IStContext) {
	this := p
	_ = this

	localctx = NewStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DGSParserRULE_st)

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
		p.Match(DGSParserST)
	}
	{
		p.SetState(102)
		p.Match(DGSParserREAL)
	}

	return localctx
}

// IClContext is an interface to support dynamic dispatch.
type IClContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClContext differentiates from other interfaces.
	IsClContext()
}

type ClContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClContext() *ClContext {
	var p = new(ClContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_cl
	return p
}

func (*ClContext) IsClContext() {}

func NewClContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClContext {
	var p = new(ClContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_cl

	return p
}

func (s *ClContext) GetParser() antlr.Parser { return s.parser }

func (s *ClContext) CL() antlr.TerminalNode {
	return s.GetToken(DGSParserCL, 0)
}

func (s *ClContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterCl(s)
	}
}

func (s *ClContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitCl(s)
	}
}

func (p *DGSParser) Cl() (localctx IClContext) {
	this := p
	_ = this

	localctx = NewClContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DGSParserRULE_cl)

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
		p.Match(DGSParserCL)
	}

	return localctx
}

// IAttributesContext is an interface to support dynamic dispatch.
type IAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributesContext differentiates from other interfaces.
	IsAttributesContext()
}

type AttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesContext() *AttributesContext {
	var p = new(AttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_attributes
	return p
}

func (*AttributesContext) IsAttributesContext() {}

func NewAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesContext {
	var p = new(AttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_attributes

	return p
}

func (s *AttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *AttributesContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *AttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterAttributes(s)
	}
}

func (s *AttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitAttributes(s)
	}
}

func (p *DGSParser) Attributes() (localctx IAttributesContext) {
	this := p
	_ = this

	localctx = NewAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DGSParserRULE_attributes)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGSParserINT)|(1<<DGSParserPLUS)|(1<<DGSParserMINUS)|(1<<DGSParserWORD)|(1<<DGSParserSTRING))) != 0 {
		{
			p.SetState(106)
			p.Attribute()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AttributeContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AttributeContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *AttributeContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AttributeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(DGSParserPLUS, 0)
}

func (s *AttributeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(DGSParserMINUS, 0)
}

func (s *AttributeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DGSParserCOMMA)
}

func (s *AttributeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserCOMMA, i)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *DGSParser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DGSParserRULE_attribute)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGSParserPLUS || _la == DGSParserMINUS {
		{
			p.SetState(112)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DGSParserPLUS || _la == DGSParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(115)
		p.Identifier()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGSParserEQUALS || _la == DGSParserCOLON {
		{
			p.SetState(116)
			p.Assign()
		}
		{
			p.SetState(117)
			p.Value()
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DGSParserCOMMA {
			{
				p.SetState(118)
				p.Match(DGSParserCOMMA)
			}
			{
				p.SetState(119)
				p.Value()
			}

			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

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
	p.RuleIndex = DGSParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(DGSParserSTRING, 0)
}

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(DGSParserINT, 0)
}

func (s *ValueContext) REAL() antlr.TerminalNode {
	return s.GetToken(DGSParserREAL, 0)
}

func (s *ValueContext) COLOR() antlr.TerminalNode {
	return s.GetToken(DGSParserCOLOR, 0)
}

func (s *ValueContext) Array_() IArray_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_Context)
}

func (s *ValueContext) A_map() IA_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IA_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IA_mapContext)
}

func (s *ValueContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *DGSParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DGSParserRULE_value)

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

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(DGSParserSTRING)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(DGSParserINT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Match(DGSParserREAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Match(DGSParserCOLOR)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Array_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.A_map()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Identifier()
		}

	}

	return localctx
}

// IArray_Context is an interface to support dynamic dispatch.
type IArray_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_Context differentiates from other interfaces.
	IsArray_Context()
}

type Array_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_Context() *Array_Context {
	var p = new(Array_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_array_
	return p
}

func (*Array_Context) IsArray_Context() {}

func NewArray_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_Context {
	var p = new(Array_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_array_

	return p
}

func (s *Array_Context) GetParser() antlr.Parser { return s.parser }

func (s *Array_Context) LBRACE() antlr.TerminalNode {
	return s.GetToken(DGSParserLBRACE, 0)
}

func (s *Array_Context) RBRACE() antlr.TerminalNode {
	return s.GetToken(DGSParserRBRACE, 0)
}

func (s *Array_Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Array_Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Array_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DGSParserCOMMA)
}

func (s *Array_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserCOMMA, i)
}

func (s *Array_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterArray_(s)
	}
}

func (s *Array_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitArray_(s)
	}
}

func (p *DGSParser) Array_() (localctx IArray_Context) {
	this := p
	_ = this

	localctx = NewArray_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DGSParserRULE_array_)
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
		p.SetState(136)
		p.Match(DGSParserLBRACE)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGSParserINT)|(1<<DGSParserREAL)|(1<<DGSParserLBRACE)|(1<<DGSParserLBRACK)|(1<<DGSParserWORD)|(1<<DGSParserSTRING)|(1<<DGSParserCOLOR))) != 0 {
		{
			p.SetState(137)
			p.Value()
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DGSParserCOMMA {
			{
				p.SetState(138)
				p.Match(DGSParserCOMMA)
			}
			{
				p.SetState(139)
				p.Value()
			}

			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(147)
		p.Match(DGSParserRBRACE)
	}

	return localctx
}

// IA_mapContext is an interface to support dynamic dispatch.
type IA_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsA_mapContext differentiates from other interfaces.
	IsA_mapContext()
}

type A_mapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_mapContext() *A_mapContext {
	var p = new(A_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_a_map
	return p
}

func (*A_mapContext) IsA_mapContext() {}

func NewA_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_mapContext {
	var p = new(A_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_a_map

	return p
}

func (s *A_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *A_mapContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DGSParserLBRACK, 0)
}

func (s *A_mapContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DGSParserRBRACK, 0)
}

func (s *A_mapContext) AllMapping() []IMappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMappingContext)(nil)).Elem())
	var tst = make([]IMappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMappingContext)
		}
	}

	return tst
}

func (s *A_mapContext) Mapping(i int) IMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMappingContext)
}

func (s *A_mapContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DGSParserCOMMA)
}

func (s *A_mapContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserCOMMA, i)
}

func (s *A_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterA_map(s)
	}
}

func (s *A_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitA_map(s)
	}
}

func (p *DGSParser) A_map() (localctx IA_mapContext) {
	this := p
	_ = this

	localctx = NewA_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DGSParserRULE_a_map)
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
		p.SetState(149)
		p.Match(DGSParserLBRACK)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGSParserINT)|(1<<DGSParserWORD)|(1<<DGSParserSTRING))) != 0 {
		{
			p.SetState(150)
			p.Mapping()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DGSParserCOMMA {
			{
				p.SetState(151)
				p.Match(DGSParserCOMMA)
			}
			{
				p.SetState(152)
				p.Mapping()
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(160)
		p.Match(DGSParserRBRACK)
	}

	return localctx
}

// IMappingContext is an interface to support dynamic dispatch.
type IMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMappingContext differentiates from other interfaces.
	IsMappingContext()
}

type MappingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMappingContext() *MappingContext {
	var p = new(MappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_mapping
	return p
}

func (*MappingContext) IsMappingContext() {}

func NewMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingContext {
	var p = new(MappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_mapping

	return p
}

func (s *MappingContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MappingContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *MappingContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterMapping(s)
	}
}

func (s *MappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitMapping(s)
	}
}

func (p *DGSParser) Mapping() (localctx IMappingContext) {
	this := p
	_ = this

	localctx = NewMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DGSParserRULE_mapping)

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
		p.Identifier()
	}
	{
		p.SetState(163)
		p.Assign()
	}
	{
		p.SetState(164)
		p.Value()
	}

	return localctx
}

// IDirectionContext is an interface to support dynamic dispatch.
type IDirectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectionContext differentiates from other interfaces.
	IsDirectionContext()
}

type DirectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectionContext() *DirectionContext {
	var p = new(DirectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_direction
	return p
}

func (*DirectionContext) IsDirectionContext() {}

func NewDirectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectionContext {
	var p = new(DirectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_direction

	return p
}

func (s *DirectionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectionContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(DGSParserLANGLE, 0)
}

func (s *DirectionContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(DGSParserRANGLE, 0)
}

func (s *DirectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterDirection(s)
	}
}

func (s *DirectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitDirection(s)
	}
}

func (p *DGSParser) Direction() (localctx IDirectionContext) {
	this := p
	_ = this

	localctx = NewDirectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DGSParserRULE_direction)
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
		p.SetState(166)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DGSParserLANGLE || _la == DGSParserRANGLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGSParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DGSParserEQUALS, 0)
}

func (s *AssignContext) COLON() antlr.TerminalNode {
	return s.GetToken(DGSParserCOLON, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *DGSParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DGSParserRULE_assign)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == DGSParserEQUALS || _la == DGSParserCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = DGSParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGSParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) STRING() antlr.TerminalNode {
	return s.GetToken(DGSParserSTRING, 0)
}

func (s *IdentifierContext) INT() antlr.TerminalNode {
	return s.GetToken(DGSParserINT, 0)
}

func (s *IdentifierContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(DGSParserWORD)
}

func (s *IdentifierContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserWORD, i)
}

func (s *IdentifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(DGSParserDOT)
}

func (s *IdentifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(DGSParserDOT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGSParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *DGSParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DGSParserRULE_identifier)
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

	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DGSParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(DGSParserSTRING)
		}

	case DGSParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(DGSParserINT)
		}

	case DGSParserWORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(172)
			p.Match(DGSParserWORD)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DGSParserDOT {
			{
				p.SetState(173)
				p.Match(DGSParserDOT)
			}
			{
				p.SetState(174)
				p.Match(DGSParserWORD)
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
