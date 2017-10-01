// Generated from TURTLE.g4 by ANTLR 4.7.

package turtle // TURTLE
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 149,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14, 2, 45,
	11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 51, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 57, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 80, 10, 9, 5, 9, 82, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 90, 10, 10, 7, 10, 92, 10, 10, 12, 10, 14, 10, 95, 11, 10, 3,
	11, 3, 11, 3, 11, 7, 11, 100, 10, 11, 12, 11, 14, 11, 103, 11, 11, 3, 12,
	3, 12, 5, 12, 107, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 112, 10, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 121, 10, 15, 3, 16,
	3, 16, 3, 16, 5, 16, 126, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 7, 18, 134, 10, 18, 12, 18, 14, 18, 137, 11, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 145, 10, 19, 3, 20, 3, 20, 3, 20, 2, 2, 21,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	2, 3, 4, 2, 22, 22, 24, 24, 2, 151, 2, 43, 3, 2, 2, 2, 4, 50, 3, 2, 2,
	2, 6, 56, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12, 67, 3,
	2, 2, 2, 14, 70, 3, 2, 2, 2, 16, 81, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20,
	96, 3, 2, 2, 2, 22, 106, 3, 2, 2, 2, 24, 111, 3, 2, 2, 2, 26, 113, 3, 2,
	2, 2, 28, 120, 3, 2, 2, 2, 30, 125, 3, 2, 2, 2, 32, 127, 3, 2, 2, 2, 34,
	131, 3, 2, 2, 2, 36, 140, 3, 2, 2, 2, 38, 146, 3, 2, 2, 2, 40, 42, 5, 4,
	3, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 3, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 51, 5, 6, 4, 2,
	47, 48, 5, 16, 9, 2, 48, 49, 7, 3, 2, 2, 49, 51, 3, 2, 2, 2, 50, 46, 3,
	2, 2, 2, 50, 47, 3, 2, 2, 2, 51, 5, 3, 2, 2, 2, 52, 57, 5, 8, 5, 2, 53,
	57, 5, 10, 6, 2, 54, 57, 5, 14, 8, 2, 55, 57, 5, 12, 7, 2, 56, 52, 3, 2,
	2, 2, 56, 53, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 7,
	3, 2, 2, 2, 58, 59, 7, 4, 2, 2, 59, 60, 7, 23, 2, 2, 60, 61, 7, 22, 2,
	2, 61, 62, 7, 3, 2, 2, 62, 9, 3, 2, 2, 2, 63, 64, 7, 5, 2, 2, 64, 65, 7,
	22, 2, 2, 65, 66, 7, 3, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68, 7, 6, 2, 2, 68,
	69, 7, 22, 2, 2, 69, 13, 3, 2, 2, 2, 70, 71, 7, 7, 2, 2, 71, 72, 7, 23,
	2, 2, 72, 73, 7, 22, 2, 2, 73, 15, 3, 2, 2, 2, 74, 75, 5, 24, 13, 2, 75,
	76, 5, 18, 10, 2, 76, 82, 3, 2, 2, 2, 77, 79, 5, 32, 17, 2, 78, 80, 5,
	18, 10, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2,
	81, 74, 3, 2, 2, 2, 81, 77, 3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 84, 5,
	22, 12, 2, 84, 93, 5, 20, 11, 2, 85, 89, 7, 8, 2, 2, 86, 87, 5, 22, 12,
	2, 87, 88, 5, 20, 11, 2, 88, 90, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 85, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2,
	93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 19, 3, 2, 2, 2, 95, 93, 3,
	2, 2, 2, 96, 101, 5, 28, 15, 2, 97, 98, 7, 9, 2, 2, 98, 100, 5, 28, 15,
	2, 99, 97, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 21, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 107, 5,
	26, 14, 2, 105, 107, 7, 10, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2,
	2, 2, 107, 23, 3, 2, 2, 2, 108, 112, 5, 38, 20, 2, 109, 112, 7, 19, 2,
	2, 110, 112, 5, 34, 18, 2, 111, 108, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2,
	111, 110, 3, 2, 2, 2, 112, 25, 3, 2, 2, 2, 113, 114, 5, 38, 20, 2, 114,
	27, 3, 2, 2, 2, 115, 121, 5, 38, 20, 2, 116, 121, 7, 19, 2, 2, 117, 121,
	5, 34, 18, 2, 118, 121, 5, 32, 17, 2, 119, 121, 5, 30, 16, 2, 120, 115,
	3, 2, 2, 2, 120, 116, 3, 2, 2, 2, 120, 117, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 120, 119, 3, 2, 2, 2, 121, 29, 3, 2, 2, 2, 122, 126, 5, 36, 19, 2,
	123, 126, 7, 16, 2, 2, 124, 126, 7, 17, 2, 2, 125, 122, 3, 2, 2, 2, 125,
	123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 31, 3, 2, 2, 2, 127, 128, 7,
	11, 2, 2, 128, 129, 5, 18, 10, 2, 129, 130, 7, 12, 2, 2, 130, 33, 3, 2,
	2, 2, 131, 135, 7, 13, 2, 2, 132, 134, 5, 28, 15, 2, 133, 132, 3, 2, 2,
	2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	138, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 7, 14, 2, 2, 139, 35,
	3, 2, 2, 2, 140, 144, 7, 18, 2, 2, 141, 145, 7, 27, 2, 2, 142, 143, 7,
	15, 2, 2, 143, 145, 5, 38, 20, 2, 144, 141, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 144, 145, 3, 2, 2, 2, 145, 37, 3, 2, 2, 2, 146, 147, 9, 2, 2, 2,
	147, 39, 3, 2, 2, 2, 16, 43, 50, 56, 79, 81, 89, 93, 101, 106, 111, 120,
	125, 135, 144,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'@prefix'", "'@base'", "'BASE'", "'PREFIX'", "';'", "','",
	"'a'", "'['", "']'", "'('", "')'", "'^^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "NumericLiteral",
	"BooleanLiteral", "String", "BlankNode", "WS", "PN_PREFIX", "IRIREF", "PNAME_NS",
	"PrefixedName", "PNAME_LN", "BLANK_NODE_LABEL", "LANGTAG", "INTEGER", "DECIMAL",
	"DOUBLE", "EXPONENT", "STRING_LITERAL_LONG_SINGLE_QUOTE", "STRING_LITERAL_LONG_QUOTE",
	"STRING_LITERAL_QUOTE", "STRING_LITERAL_SINGLE_QUOTE", "UCHAR", "ECHAR",
	"ANON_WS", "ANON", "PN_CHARS_BASE", "PN_CHARS_U", "PN_CHARS", "PN_LOCAL",
	"PLX", "PERCENT", "HEX", "PN_LOCAL_ESC",
}

var ruleNames = []string{
	"turtleDoc", "statement", "directive", "prefixID", "base", "sparqlBase",
	"sparqlPrefix", "triples", "predicateObjectList", "objectList", "verb",
	"subject", "predicate", "object", "literal", "blankNodePropertyList", "collection",
	"rdfLiteral", "iri",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TURTLEParser struct {
	*antlr.BaseParser
}

func NewTURTLEParser(input antlr.TokenStream) *TURTLEParser {
	this := new(TURTLEParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TURTLE.g4"

	return this
}

// TURTLEParser tokens.
const (
	TURTLEParserEOF                              = antlr.TokenEOF
	TURTLEParserT__0                             = 1
	TURTLEParserT__1                             = 2
	TURTLEParserT__2                             = 3
	TURTLEParserT__3                             = 4
	TURTLEParserT__4                             = 5
	TURTLEParserT__5                             = 6
	TURTLEParserT__6                             = 7
	TURTLEParserT__7                             = 8
	TURTLEParserT__8                             = 9
	TURTLEParserT__9                             = 10
	TURTLEParserT__10                            = 11
	TURTLEParserT__11                            = 12
	TURTLEParserT__12                            = 13
	TURTLEParserNumericLiteral                   = 14
	TURTLEParserBooleanLiteral                   = 15
	TURTLEParserString                           = 16
	TURTLEParserBlankNode                        = 17
	TURTLEParserWS                               = 18
	TURTLEParserPN_PREFIX                        = 19
	TURTLEParserIRIREF                           = 20
	TURTLEParserPNAME_NS                         = 21
	TURTLEParserPrefixedName                     = 22
	TURTLEParserPNAME_LN                         = 23
	TURTLEParserBLANK_NODE_LABEL                 = 24
	TURTLEParserLANGTAG                          = 25
	TURTLEParserINTEGER                          = 26
	TURTLEParserDECIMAL                          = 27
	TURTLEParserDOUBLE                           = 28
	TURTLEParserEXPONENT                         = 29
	TURTLEParserSTRING_LITERAL_LONG_SINGLE_QUOTE = 30
	TURTLEParserSTRING_LITERAL_LONG_QUOTE        = 31
	TURTLEParserSTRING_LITERAL_QUOTE             = 32
	TURTLEParserSTRING_LITERAL_SINGLE_QUOTE      = 33
	TURTLEParserUCHAR                            = 34
	TURTLEParserECHAR                            = 35
	TURTLEParserANON_WS                          = 36
	TURTLEParserANON                             = 37
	TURTLEParserPN_CHARS_BASE                    = 38
	TURTLEParserPN_CHARS_U                       = 39
	TURTLEParserPN_CHARS                         = 40
	TURTLEParserPN_LOCAL                         = 41
	TURTLEParserPLX                              = 42
	TURTLEParserPERCENT                          = 43
	TURTLEParserHEX                              = 44
	TURTLEParserPN_LOCAL_ESC                     = 45
)

// TURTLEParser rules.
const (
	TURTLEParserRULE_turtleDoc             = 0
	TURTLEParserRULE_statement             = 1
	TURTLEParserRULE_directive             = 2
	TURTLEParserRULE_prefixID              = 3
	TURTLEParserRULE_base                  = 4
	TURTLEParserRULE_sparqlBase            = 5
	TURTLEParserRULE_sparqlPrefix          = 6
	TURTLEParserRULE_triples               = 7
	TURTLEParserRULE_predicateObjectList   = 8
	TURTLEParserRULE_objectList            = 9
	TURTLEParserRULE_verb                  = 10
	TURTLEParserRULE_subject               = 11
	TURTLEParserRULE_predicate             = 12
	TURTLEParserRULE_object                = 13
	TURTLEParserRULE_literal               = 14
	TURTLEParserRULE_blankNodePropertyList = 15
	TURTLEParserRULE_collection            = 16
	TURTLEParserRULE_rdfLiteral            = 17
	TURTLEParserRULE_iri                   = 18
)

// ITurtleDocContext is an interface to support dynamic dispatch.
type ITurtleDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTurtleDocContext differentiates from other interfaces.
	IsTurtleDocContext()
}

type TurtleDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTurtleDocContext() *TurtleDocContext {
	var p = new(TurtleDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_turtleDoc
	return p
}

func (*TurtleDocContext) IsTurtleDocContext() {}

func NewTurtleDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TurtleDocContext {
	var p = new(TurtleDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_turtleDoc

	return p
}

func (s *TurtleDocContext) GetParser() antlr.Parser { return s.parser }

func (s *TurtleDocContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *TurtleDocContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *TurtleDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TurtleDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TurtleDocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterTurtleDoc(s)
	}
}

func (s *TurtleDocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitTurtleDoc(s)
	}
}

func (s *TurtleDocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitTurtleDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) TurtleDoc() (localctx ITurtleDocContext) {
	localctx = NewTurtleDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TURTLEParserRULE_turtleDoc)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TURTLEParserT__1)|(1<<TURTLEParserT__2)|(1<<TURTLEParserT__3)|(1<<TURTLEParserT__4)|(1<<TURTLEParserT__8)|(1<<TURTLEParserT__10)|(1<<TURTLEParserBlankNode)|(1<<TURTLEParserIRIREF)|(1<<TURTLEParserPrefixedName))) != 0 {
		{
			p.SetState(38)
			p.Statement()
		}

		p.SetState(43)
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
	p.RuleIndex = TURTLEParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *StatementContext) Triples() ITriplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITriplesContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TURTLEParserRULE_statement)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserT__1, TURTLEParserT__2, TURTLEParserT__3, TURTLEParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Directive()
		}

	case TURTLEParserT__8, TURTLEParserT__10, TURTLEParserBlankNode, TURTLEParserIRIREF, TURTLEParserPrefixedName:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Triples()
		}
		{
			p.SetState(46)
			p.Match(TURTLEParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = TURTLEParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) PrefixID() IPrefixIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixIDContext)
}

func (s *DirectiveContext) Base() IBaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseContext)
}

func (s *DirectiveContext) SparqlPrefix() ISparqlPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISparqlPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISparqlPrefixContext)
}

func (s *DirectiveContext) SparqlBase() ISparqlBaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISparqlBaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISparqlBaseContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (s *DirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TURTLEParserRULE_directive)

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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.PrefixID()
		}

	case TURTLEParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Base()
		}

	case TURTLEParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.SparqlPrefix()
		}

	case TURTLEParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.SparqlBase()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrefixIDContext is an interface to support dynamic dispatch.
type IPrefixIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixIDContext differentiates from other interfaces.
	IsPrefixIDContext()
}

type PrefixIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixIDContext() *PrefixIDContext {
	var p = new(PrefixIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_prefixID
	return p
}

func (*PrefixIDContext) IsPrefixIDContext() {}

func NewPrefixIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixIDContext {
	var p = new(PrefixIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_prefixID

	return p
}

func (s *PrefixIDContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixIDContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(TURTLEParserPNAME_NS, 0)
}

func (s *PrefixIDContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(TURTLEParserIRIREF, 0)
}

func (s *PrefixIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterPrefixID(s)
	}
}

func (s *PrefixIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitPrefixID(s)
	}
}

func (s *PrefixIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitPrefixID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) PrefixID() (localctx IPrefixIDContext) {
	localctx = NewPrefixIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TURTLEParserRULE_prefixID)

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
		p.Match(TURTLEParserT__1)
	}
	{
		p.SetState(57)
		p.Match(TURTLEParserPNAME_NS)
	}
	{
		p.SetState(58)
		p.Match(TURTLEParserIRIREF)
	}
	{
		p.SetState(59)
		p.Match(TURTLEParserT__0)
	}

	return localctx
}

// IBaseContext is an interface to support dynamic dispatch.
type IBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseContext differentiates from other interfaces.
	IsBaseContext()
}

type BaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseContext() *BaseContext {
	var p = new(BaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_base
	return p
}

func (*BaseContext) IsBaseContext() {}

func NewBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseContext {
	var p = new(BaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_base

	return p
}

func (s *BaseContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(TURTLEParserIRIREF, 0)
}

func (s *BaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterBase(s)
	}
}

func (s *BaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitBase(s)
	}
}

func (s *BaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitBase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Base() (localctx IBaseContext) {
	localctx = NewBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TURTLEParserRULE_base)

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
		p.SetState(61)
		p.Match(TURTLEParserT__2)
	}
	{
		p.SetState(62)
		p.Match(TURTLEParserIRIREF)
	}
	{
		p.SetState(63)
		p.Match(TURTLEParserT__0)
	}

	return localctx
}

// ISparqlBaseContext is an interface to support dynamic dispatch.
type ISparqlBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSparqlBaseContext differentiates from other interfaces.
	IsSparqlBaseContext()
}

type SparqlBaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySparqlBaseContext() *SparqlBaseContext {
	var p = new(SparqlBaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_sparqlBase
	return p
}

func (*SparqlBaseContext) IsSparqlBaseContext() {}

func NewSparqlBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SparqlBaseContext {
	var p = new(SparqlBaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_sparqlBase

	return p
}

func (s *SparqlBaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SparqlBaseContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(TURTLEParserIRIREF, 0)
}

func (s *SparqlBaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SparqlBaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SparqlBaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterSparqlBase(s)
	}
}

func (s *SparqlBaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitSparqlBase(s)
	}
}

func (s *SparqlBaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitSparqlBase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) SparqlBase() (localctx ISparqlBaseContext) {
	localctx = NewSparqlBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TURTLEParserRULE_sparqlBase)

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
		p.Match(TURTLEParserT__3)
	}
	{
		p.SetState(66)
		p.Match(TURTLEParserIRIREF)
	}

	return localctx
}

// ISparqlPrefixContext is an interface to support dynamic dispatch.
type ISparqlPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSparqlPrefixContext differentiates from other interfaces.
	IsSparqlPrefixContext()
}

type SparqlPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySparqlPrefixContext() *SparqlPrefixContext {
	var p = new(SparqlPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_sparqlPrefix
	return p
}

func (*SparqlPrefixContext) IsSparqlPrefixContext() {}

func NewSparqlPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SparqlPrefixContext {
	var p = new(SparqlPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_sparqlPrefix

	return p
}

func (s *SparqlPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SparqlPrefixContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(TURTLEParserPNAME_NS, 0)
}

func (s *SparqlPrefixContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(TURTLEParserIRIREF, 0)
}

func (s *SparqlPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SparqlPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SparqlPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterSparqlPrefix(s)
	}
}

func (s *SparqlPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitSparqlPrefix(s)
	}
}

func (s *SparqlPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitSparqlPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) SparqlPrefix() (localctx ISparqlPrefixContext) {
	localctx = NewSparqlPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TURTLEParserRULE_sparqlPrefix)

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
		p.Match(TURTLEParserT__4)
	}
	{
		p.SetState(69)
		p.Match(TURTLEParserPNAME_NS)
	}
	{
		p.SetState(70)
		p.Match(TURTLEParserIRIREF)
	}

	return localctx
}

// ITriplesContext is an interface to support dynamic dispatch.
type ITriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTriplesContext differentiates from other interfaces.
	IsTriplesContext()
}

type TriplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesContext() *TriplesContext {
	var p = new(TriplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_triples
	return p
}

func (*TriplesContext) IsTriplesContext() {}

func NewTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesContext {
	var p = new(TriplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_triples

	return p
}

func (s *TriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesContext) Subject() ISubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubjectContext)
}

func (s *TriplesContext) PredicateObjectList() IPredicateObjectListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateObjectListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateObjectListContext)
}

func (s *TriplesContext) BlankNodePropertyList() IBlankNodePropertyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodePropertyListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlankNodePropertyListContext)
}

func (s *TriplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterTriples(s)
	}
}

func (s *TriplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitTriples(s)
	}
}

func (s *TriplesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitTriples(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Triples() (localctx ITriplesContext) {
	localctx = NewTriplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TURTLEParserRULE_triples)
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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserT__10, TURTLEParserBlankNode, TURTLEParserIRIREF, TURTLEParserPrefixedName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Subject()
		}
		{
			p.SetState(73)
			p.PredicateObjectList()
		}

	case TURTLEParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.BlankNodePropertyList()
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TURTLEParserT__7)|(1<<TURTLEParserIRIREF)|(1<<TURTLEParserPrefixedName))) != 0 {
			{
				p.SetState(76)
				p.PredicateObjectList()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPredicateObjectListContext is an interface to support dynamic dispatch.
type IPredicateObjectListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateObjectListContext differentiates from other interfaces.
	IsPredicateObjectListContext()
}

type PredicateObjectListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateObjectListContext() *PredicateObjectListContext {
	var p = new(PredicateObjectListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_predicateObjectList
	return p
}

func (*PredicateObjectListContext) IsPredicateObjectListContext() {}

func NewPredicateObjectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateObjectListContext {
	var p = new(PredicateObjectListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_predicateObjectList

	return p
}

func (s *PredicateObjectListContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateObjectListContext) AllVerb() []IVerbContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVerbContext)(nil)).Elem())
	var tst = make([]IVerbContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVerbContext)
		}
	}

	return tst
}

func (s *PredicateObjectListContext) Verb(i int) IVerbContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVerbContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVerbContext)
}

func (s *PredicateObjectListContext) AllObjectList() []IObjectListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectListContext)(nil)).Elem())
	var tst = make([]IObjectListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectListContext)
		}
	}

	return tst
}

func (s *PredicateObjectListContext) ObjectList(i int) IObjectListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectListContext)
}

func (s *PredicateObjectListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateObjectListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateObjectListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterPredicateObjectList(s)
	}
}

func (s *PredicateObjectListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitPredicateObjectList(s)
	}
}

func (s *PredicateObjectListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitPredicateObjectList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) PredicateObjectList() (localctx IPredicateObjectListContext) {
	localctx = NewPredicateObjectListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TURTLEParserRULE_predicateObjectList)
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
		p.SetState(81)
		p.Verb()
	}
	{
		p.SetState(82)
		p.ObjectList()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TURTLEParserT__5 {
		{
			p.SetState(83)
			p.Match(TURTLEParserT__5)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TURTLEParserT__7)|(1<<TURTLEParserIRIREF)|(1<<TURTLEParserPrefixedName))) != 0 {
			{
				p.SetState(84)
				p.Verb()
			}
			{
				p.SetState(85)
				p.ObjectList()
			}

		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjectListContext is an interface to support dynamic dispatch.
type IObjectListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectListContext differentiates from other interfaces.
	IsObjectListContext()
}

type ObjectListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectListContext() *ObjectListContext {
	var p = new(ObjectListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_objectList
	return p
}

func (*ObjectListContext) IsObjectListContext() {}

func NewObjectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectListContext {
	var p = new(ObjectListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_objectList

	return p
}

func (s *ObjectListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectListContext) AllObject() []IObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectContext)(nil)).Elem())
	var tst = make([]IObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectContext)
		}
	}

	return tst
}

func (s *ObjectListContext) Object(i int) IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ObjectListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterObjectList(s)
	}
}

func (s *ObjectListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitObjectList(s)
	}
}

func (s *ObjectListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitObjectList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) ObjectList() (localctx IObjectListContext) {
	localctx = NewObjectListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TURTLEParserRULE_objectList)
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
		p.SetState(94)
		p.Object()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TURTLEParserT__6 {
		{
			p.SetState(95)
			p.Match(TURTLEParserT__6)
		}
		{
			p.SetState(96)
			p.Object()
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVerbContext is an interface to support dynamic dispatch.
type IVerbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVerbContext differentiates from other interfaces.
	IsVerbContext()
}

type VerbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbContext() *VerbContext {
	var p = new(VerbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_verb
	return p
}

func (*VerbContext) IsVerbContext() {}

func NewVerbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbContext {
	var p = new(VerbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_verb

	return p
}

func (s *VerbContext) GetParser() antlr.Parser { return s.parser }

func (s *VerbContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *VerbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterVerb(s)
	}
}

func (s *VerbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitVerb(s)
	}
}

func (s *VerbContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitVerb(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Verb() (localctx IVerbContext) {
	localctx = NewVerbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TURTLEParserRULE_verb)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserIRIREF, TURTLEParserPrefixedName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Predicate()
		}

	case TURTLEParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(TURTLEParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubjectContext is an interface to support dynamic dispatch.
type ISubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubjectContext differentiates from other interfaces.
	IsSubjectContext()
}

type SubjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubjectContext() *SubjectContext {
	var p = new(SubjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_subject
	return p
}

func (*SubjectContext) IsSubjectContext() {}

func NewSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubjectContext {
	var p = new(SubjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_subject

	return p
}

func (s *SubjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SubjectContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *SubjectContext) BlankNode() antlr.TerminalNode {
	return s.GetToken(TURTLEParserBlankNode, 0)
}

func (s *SubjectContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *SubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterSubject(s)
	}
}

func (s *SubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitSubject(s)
	}
}

func (s *SubjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitSubject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Subject() (localctx ISubjectContext) {
	localctx = NewSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TURTLEParserRULE_subject)

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
	case TURTLEParserIRIREF, TURTLEParserPrefixedName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Iri()
		}

	case TURTLEParserBlankNode:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(TURTLEParserBlankNode)
		}

	case TURTLEParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Collection()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TURTLEParserRULE_predicate)

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
		p.Iri()
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *ObjectContext) BlankNode() antlr.TerminalNode {
	return s.GetToken(TURTLEParserBlankNode, 0)
}

func (s *ObjectContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *ObjectContext) BlankNodePropertyList() IBlankNodePropertyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodePropertyListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlankNodePropertyListContext)
}

func (s *ObjectContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitObject(s)
	}
}

func (s *ObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TURTLEParserRULE_object)

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserIRIREF, TURTLEParserPrefixedName:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Iri()
		}

	case TURTLEParserBlankNode:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(TURTLEParserBlankNode)
		}

	case TURTLEParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Collection()
		}

	case TURTLEParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.BlankNodePropertyList()
		}

	case TURTLEParserNumericLiteral, TURTLEParserBooleanLiteral, TURTLEParserString:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) RdfLiteral() IRdfLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *LiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(TURTLEParserNumericLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(TURTLEParserBooleanLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TURTLEParserRULE_literal)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserString:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.RdfLiteral()
		}

	case TURTLEParserNumericLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Match(TURTLEParserNumericLiteral)
		}

	case TURTLEParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.Match(TURTLEParserBooleanLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlankNodePropertyListContext is an interface to support dynamic dispatch.
type IBlankNodePropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankNodePropertyListContext differentiates from other interfaces.
	IsBlankNodePropertyListContext()
}

type BlankNodePropertyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodePropertyListContext() *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_blankNodePropertyList
	return p
}

func (*BlankNodePropertyListContext) IsBlankNodePropertyListContext() {}

func NewBlankNodePropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_blankNodePropertyList

	return p
}

func (s *BlankNodePropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodePropertyListContext) PredicateObjectList() IPredicateObjectListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateObjectListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateObjectListContext)
}

func (s *BlankNodePropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodePropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankNodePropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterBlankNodePropertyList(s)
	}
}

func (s *BlankNodePropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitBlankNodePropertyList(s)
	}
}

func (s *BlankNodePropertyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitBlankNodePropertyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) BlankNodePropertyList() (localctx IBlankNodePropertyListContext) {
	localctx = NewBlankNodePropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TURTLEParserRULE_blankNodePropertyList)

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
		p.SetState(125)
		p.Match(TURTLEParserT__8)
	}
	{
		p.SetState(126)
		p.PredicateObjectList()
	}
	{
		p.SetState(127)
		p.Match(TURTLEParserT__9)
	}

	return localctx
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_collection
	return p
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) AllObject() []IObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectContext)(nil)).Elem())
	var tst = make([]IObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectContext)
		}
	}

	return tst
}

func (s *CollectionContext) Object(i int) IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (s *CollectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitCollection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TURTLEParserRULE_collection)
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
		p.SetState(129)
		p.Match(TURTLEParserT__10)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TURTLEParserT__8)|(1<<TURTLEParserT__10)|(1<<TURTLEParserNumericLiteral)|(1<<TURTLEParserBooleanLiteral)|(1<<TURTLEParserString)|(1<<TURTLEParserBlankNode)|(1<<TURTLEParserIRIREF)|(1<<TURTLEParserPrefixedName))) != 0 {
		{
			p.SetState(130)
			p.Object()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(TURTLEParserT__11)
	}

	return localctx
}

// IRdfLiteralContext is an interface to support dynamic dispatch.
type IRdfLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRdfLiteralContext differentiates from other interfaces.
	IsRdfLiteralContext()
}

type RdfLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfLiteralContext() *RdfLiteralContext {
	var p = new(RdfLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_rdfLiteral
	return p
}

func (*RdfLiteralContext) IsRdfLiteralContext() {}

func NewRdfLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfLiteralContext {
	var p = new(RdfLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_rdfLiteral

	return p
}

func (s *RdfLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfLiteralContext) String() antlr.TerminalNode {
	return s.GetToken(TURTLEParserString, 0)
}

func (s *RdfLiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(TURTLEParserLANGTAG, 0)
}

func (s *RdfLiteralContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *RdfLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RdfLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitRdfLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) RdfLiteral() (localctx IRdfLiteralContext) {
	localctx = NewRdfLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TURTLEParserRULE_rdfLiteral)

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
		p.Match(TURTLEParserString)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TURTLEParserLANGTAG:
		{
			p.SetState(139)
			p.Match(TURTLEParserLANGTAG)
		}

	case TURTLEParserT__12:
		{
			p.SetState(140)
			p.Match(TURTLEParserT__12)
		}
		{
			p.SetState(141)
			p.Iri()
		}

	case TURTLEParserT__0, TURTLEParserT__5, TURTLEParserT__6, TURTLEParserT__8, TURTLEParserT__9, TURTLEParserT__10, TURTLEParserT__11, TURTLEParserNumericLiteral, TURTLEParserBooleanLiteral, TURTLEParserString, TURTLEParserBlankNode, TURTLEParserIRIREF, TURTLEParserPrefixedName:

	default:
	}

	return localctx
}

// IIriContext is an interface to support dynamic dispatch.
type IIriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriContext differentiates from other interfaces.
	IsIriContext()
}

type IriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriContext() *IriContext {
	var p = new(IriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TURTLEParserRULE_iri
	return p
}

func (*IriContext) IsIriContext() {}

func NewIriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriContext {
	var p = new(IriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TURTLEParserRULE_iri

	return p
}

func (s *IriContext) GetParser() antlr.Parser { return s.parser }

func (s *IriContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(TURTLEParserIRIREF, 0)
}

func (s *IriContext) PrefixedName() antlr.TerminalNode {
	return s.GetToken(TURTLEParserPrefixedName, 0)
}

func (s *IriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.EnterIri(s)
	}
}

func (s *IriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TURTLEListener); ok {
		listenerT.ExitIri(s)
	}
}

func (s *IriContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TURTLEVisitor:
		return t.VisitIri(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TURTLEParser) Iri() (localctx IIriContext) {
	localctx = NewIriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TURTLEParserRULE_iri)
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
	p.SetState(144)
	_la = p.GetTokenStream().LA(1)

	if !(_la == TURTLEParserIRIREF || _la == TURTLEParserPrefixedName) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
