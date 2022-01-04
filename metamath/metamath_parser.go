// Code generated from metamath.g4 by ANTLR 4.9.3. DO NOT EDIT.

package metamath // metamath
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 189,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 3, 3, 3, 3, 3, 5, 3,
	56, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 64, 10, 5, 13, 5,
	14, 5, 65, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 75, 10, 6, 3,
	7, 3, 7, 7, 7, 79, 10, 7, 12, 7, 14, 7, 82, 11, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 6, 8, 88, 10, 8, 13, 8, 14, 8, 89, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 7, 9, 98, 10, 9, 12, 9, 14, 9, 101, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	5, 10, 107, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 7, 12, 119, 10, 12, 12, 12, 14, 12, 122, 11, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 5, 13, 128, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7,
	14, 134, 10, 14, 12, 14, 14, 14, 137, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 7, 15, 145, 10, 15, 12, 15, 14, 15, 148, 11, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 156, 10, 16, 3, 17, 6, 17, 159,
	10, 17, 13, 17, 14, 17, 160, 3, 18, 3, 18, 7, 18, 165, 10, 18, 12, 18,
	14, 18, 168, 11, 18, 3, 18, 3, 18, 6, 18, 172, 10, 18, 13, 18, 14, 18,
	173, 3, 19, 3, 19, 3, 20, 6, 20, 179, 10, 20, 13, 20, 14, 20, 180, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 4, 4,
	2, 16, 16, 20, 20, 3, 2, 17, 19, 2, 187, 2, 49, 3, 2, 2, 2, 4, 55, 3, 2,
	2, 2, 6, 57, 3, 2, 2, 2, 8, 61, 3, 2, 2, 2, 10, 74, 3, 2, 2, 2, 12, 76,
	3, 2, 2, 2, 14, 85, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 106, 3, 2, 2, 2,
	20, 108, 3, 2, 2, 2, 22, 114, 3, 2, 2, 2, 24, 127, 3, 2, 2, 2, 26, 129,
	3, 2, 2, 2, 28, 140, 3, 2, 2, 2, 30, 155, 3, 2, 2, 2, 32, 158, 3, 2, 2,
	2, 34, 162, 3, 2, 2, 2, 36, 175, 3, 2, 2, 2, 38, 178, 3, 2, 2, 2, 40, 182,
	3, 2, 2, 2, 42, 184, 3, 2, 2, 2, 44, 186, 3, 2, 2, 2, 46, 48, 5, 4, 3,
	2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 50, 3, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 56, 5, 6, 4, 2,
	53, 56, 5, 8, 5, 2, 54, 56, 5, 10, 6, 2, 55, 52, 3, 2, 2, 2, 55, 53, 3,
	2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 5, 3, 2, 2, 2, 57, 58, 7, 3, 2, 2, 58,
	59, 5, 40, 21, 2, 59, 60, 7, 4, 2, 2, 60, 7, 3, 2, 2, 2, 61, 63, 7, 5,
	2, 2, 62, 64, 5, 42, 22, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65,
	63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 7, 6, 2,
	2, 68, 9, 3, 2, 2, 2, 69, 75, 5, 12, 7, 2, 70, 75, 5, 14, 8, 2, 71, 75,
	5, 16, 9, 2, 72, 75, 5, 18, 10, 2, 73, 75, 5, 24, 13, 2, 74, 69, 3, 2,
	2, 2, 74, 70, 3, 2, 2, 2, 74, 71, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 73,
	3, 2, 2, 2, 75, 11, 3, 2, 2, 2, 76, 80, 7, 7, 2, 2, 77, 79, 5, 10, 6, 2,
	78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3,
	2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84, 7, 8, 2, 2, 84,
	13, 3, 2, 2, 2, 85, 87, 7, 9, 2, 2, 86, 88, 5, 44, 23, 2, 87, 86, 3, 2,
	2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 92, 7, 6, 2, 2, 92, 15, 3, 2, 2, 2, 93, 94, 7, 10, 2, 2,
	94, 95, 5, 44, 23, 2, 95, 99, 5, 44, 23, 2, 96, 98, 5, 44, 23, 2, 97, 96,
	3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2,
	2, 100, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 103, 7, 6, 2, 2, 103,
	17, 3, 2, 2, 2, 104, 107, 5, 20, 11, 2, 105, 107, 5, 22, 12, 2, 106, 104,
	3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 19, 3, 2, 2, 2, 108, 109, 7, 20,
	2, 2, 109, 110, 7, 11, 2, 2, 110, 111, 5, 36, 19, 2, 111, 112, 5, 44, 23,
	2, 112, 113, 7, 6, 2, 2, 113, 21, 3, 2, 2, 2, 114, 115, 7, 20, 2, 2, 115,
	116, 7, 12, 2, 2, 116, 120, 5, 36, 19, 2, 117, 119, 5, 38, 20, 2, 118,
	117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121,
	3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 6,
	2, 2, 124, 23, 3, 2, 2, 2, 125, 128, 5, 26, 14, 2, 126, 128, 5, 28, 15,
	2, 127, 125, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 25, 3, 2, 2, 2, 129,
	130, 7, 20, 2, 2, 130, 131, 7, 13, 2, 2, 131, 135, 5, 36, 19, 2, 132, 134,
	5, 38, 20, 2, 133, 132, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3,
	2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2, 137, 135, 3, 2, 2,
	2, 138, 139, 7, 6, 2, 2, 139, 27, 3, 2, 2, 2, 140, 141, 7, 20, 2, 2, 141,
	142, 7, 14, 2, 2, 142, 146, 5, 36, 19, 2, 143, 145, 5, 38, 20, 2, 144,
	143, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147,
	3, 2, 2, 2, 147, 149, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 15,
	2, 2, 150, 151, 5, 30, 16, 2, 151, 152, 7, 6, 2, 2, 152, 29, 3, 2, 2, 2,
	153, 156, 5, 32, 17, 2, 154, 156, 5, 34, 18, 2, 155, 153, 3, 2, 2, 2, 155,
	154, 3, 2, 2, 2, 156, 31, 3, 2, 2, 2, 157, 159, 9, 2, 2, 2, 158, 157, 3,
	2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2,
	2, 161, 33, 3, 2, 2, 2, 162, 166, 7, 17, 2, 2, 163, 165, 7, 20, 2, 2, 164,
	163, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167,
	3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 171, 7, 18,
	2, 2, 170, 172, 7, 21, 2, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2,
	173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 35, 3, 2, 2, 2, 175, 176,
	5, 42, 22, 2, 176, 37, 3, 2, 2, 2, 177, 179, 9, 3, 2, 2, 178, 177, 3, 2,
	2, 2, 179, 180, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2,
	181, 39, 3, 2, 2, 2, 182, 183, 5, 38, 20, 2, 183, 41, 3, 2, 2, 2, 184,
	185, 5, 38, 20, 2, 185, 43, 3, 2, 2, 2, 186, 187, 5, 38, 20, 2, 187, 45,
	3, 2, 2, 2, 19, 49, 55, 65, 74, 80, 89, 99, 106, 120, 127, 135, 146, 155,
	160, 166, 173, 180,
}
var literalNames = []string{
	"", "'$['", "'$]'", "'$c'", "'$.'", "'${'", "'$}'", "'$v'", "'$d'", "'$f'",
	"'$e'", "'$a'", "'$p'", "'$='", "'?'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LPAREN", "RPAREN",
	"PRINTABLECHARACTER", "LABEL", "COMPRESSEDPROOFBLOCK", "BLOCK_COMMENT",
	"WS",
}

var ruleNames = []string{
	"database", "outermostscopestmt", "includestmt", "constantstmt", "stmt",
	"block", "variablestmt", "disjointstmt", "hypothesisstmt", "floatingstmt",
	"essentialstmt", "assertstmt", "axiomstmt", "provablestmt", "proof", "uncompressedproof",
	"compressedproof", "typecode", "mathsymbol", "filename", "constant", "variable",
}

type metamathParser struct {
	*antlr.BaseParser
}

// NewmetamathParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *metamathParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmetamathParser(input antlr.TokenStream) *metamathParser {
	this := new(metamathParser)
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
	this.GrammarFileName = "metamath.g4"

	return this
}

// metamathParser tokens.
const (
	metamathParserEOF                  = antlr.TokenEOF
	metamathParserT__0                 = 1
	metamathParserT__1                 = 2
	metamathParserT__2                 = 3
	metamathParserT__3                 = 4
	metamathParserT__4                 = 5
	metamathParserT__5                 = 6
	metamathParserT__6                 = 7
	metamathParserT__7                 = 8
	metamathParserT__8                 = 9
	metamathParserT__9                 = 10
	metamathParserT__10                = 11
	metamathParserT__11                = 12
	metamathParserT__12                = 13
	metamathParserT__13                = 14
	metamathParserLPAREN               = 15
	metamathParserRPAREN               = 16
	metamathParserPRINTABLECHARACTER   = 17
	metamathParserLABEL                = 18
	metamathParserCOMPRESSEDPROOFBLOCK = 19
	metamathParserBLOCK_COMMENT        = 20
	metamathParserWS                   = 21
)

// metamathParser rules.
const (
	metamathParserRULE_database           = 0
	metamathParserRULE_outermostscopestmt = 1
	metamathParserRULE_includestmt        = 2
	metamathParserRULE_constantstmt       = 3
	metamathParserRULE_stmt               = 4
	metamathParserRULE_block              = 5
	metamathParserRULE_variablestmt       = 6
	metamathParserRULE_disjointstmt       = 7
	metamathParserRULE_hypothesisstmt     = 8
	metamathParserRULE_floatingstmt       = 9
	metamathParserRULE_essentialstmt      = 10
	metamathParserRULE_assertstmt         = 11
	metamathParserRULE_axiomstmt          = 12
	metamathParserRULE_provablestmt       = 13
	metamathParserRULE_proof              = 14
	metamathParserRULE_uncompressedproof  = 15
	metamathParserRULE_compressedproof    = 16
	metamathParserRULE_typecode           = 17
	metamathParserRULE_mathsymbol         = 18
	metamathParserRULE_filename           = 19
	metamathParserRULE_constant           = 20
	metamathParserRULE_variable           = 21
)

// IDatabaseContext is an interface to support dynamic dispatch.
type IDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabaseContext differentiates from other interfaces.
	IsDatabaseContext()
}

type DatabaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseContext() *DatabaseContext {
	var p = new(DatabaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_database
	return p
}

func (*DatabaseContext) IsDatabaseContext() {}

func NewDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseContext {
	var p = new(DatabaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_database

	return p
}

func (s *DatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseContext) AllOutermostscopestmt() []IOutermostscopestmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOutermostscopestmtContext)(nil)).Elem())
	var tst = make([]IOutermostscopestmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOutermostscopestmtContext)
		}
	}

	return tst
}

func (s *DatabaseContext) Outermostscopestmt(i int) IOutermostscopestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutermostscopestmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOutermostscopestmtContext)
}

func (s *DatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterDatabase(s)
	}
}

func (s *DatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitDatabase(s)
	}
}

func (p *metamathParser) Database() (localctx IDatabaseContext) {
	this := p
	_ = this

	localctx = NewDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, metamathParserRULE_database)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserT__0)|(1<<metamathParserT__2)|(1<<metamathParserT__4)|(1<<metamathParserT__6)|(1<<metamathParserT__7)|(1<<metamathParserLABEL))) != 0 {
		{
			p.SetState(44)
			p.Outermostscopestmt()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOutermostscopestmtContext is an interface to support dynamic dispatch.
type IOutermostscopestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutermostscopestmtContext differentiates from other interfaces.
	IsOutermostscopestmtContext()
}

type OutermostscopestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutermostscopestmtContext() *OutermostscopestmtContext {
	var p = new(OutermostscopestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_outermostscopestmt
	return p
}

func (*OutermostscopestmtContext) IsOutermostscopestmtContext() {}

func NewOutermostscopestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutermostscopestmtContext {
	var p = new(OutermostscopestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_outermostscopestmt

	return p
}

func (s *OutermostscopestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *OutermostscopestmtContext) Includestmt() IIncludestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludestmtContext)
}

func (s *OutermostscopestmtContext) Constantstmt() IConstantstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantstmtContext)
}

func (s *OutermostscopestmtContext) Stmt() IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *OutermostscopestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutermostscopestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutermostscopestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterOutermostscopestmt(s)
	}
}

func (s *OutermostscopestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitOutermostscopestmt(s)
	}
}

func (p *metamathParser) Outermostscopestmt() (localctx IOutermostscopestmtContext) {
	this := p
	_ = this

	localctx = NewOutermostscopestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, metamathParserRULE_outermostscopestmt)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case metamathParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Includestmt()
		}

	case metamathParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Constantstmt()
		}

	case metamathParserT__4, metamathParserT__6, metamathParserT__7, metamathParserLABEL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIncludestmtContext is an interface to support dynamic dispatch.
type IIncludestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludestmtContext differentiates from other interfaces.
	IsIncludestmtContext()
}

type IncludestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludestmtContext() *IncludestmtContext {
	var p = new(IncludestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_includestmt
	return p
}

func (*IncludestmtContext) IsIncludestmtContext() {}

func NewIncludestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludestmtContext {
	var p = new(IncludestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_includestmt

	return p
}

func (s *IncludestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludestmtContext) Filename() IFilenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilenameContext)
}

func (s *IncludestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterIncludestmt(s)
	}
}

func (s *IncludestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitIncludestmt(s)
	}
}

func (p *metamathParser) Includestmt() (localctx IIncludestmtContext) {
	this := p
	_ = this

	localctx = NewIncludestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, metamathParserRULE_includestmt)

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
		p.Match(metamathParserT__0)
	}
	{
		p.SetState(56)
		p.Filename()
	}
	{
		p.SetState(57)
		p.Match(metamathParserT__1)
	}

	return localctx
}

// IConstantstmtContext is an interface to support dynamic dispatch.
type IConstantstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantstmtContext differentiates from other interfaces.
	IsConstantstmtContext()
}

type ConstantstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantstmtContext() *ConstantstmtContext {
	var p = new(ConstantstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_constantstmt
	return p
}

func (*ConstantstmtContext) IsConstantstmtContext() {}

func NewConstantstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantstmtContext {
	var p = new(ConstantstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_constantstmt

	return p
}

func (s *ConstantstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantstmtContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *ConstantstmtContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterConstantstmt(s)
	}
}

func (s *ConstantstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitConstantstmt(s)
	}
}

func (p *metamathParser) Constantstmt() (localctx IConstantstmtContext) {
	this := p
	_ = this

	localctx = NewConstantstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, metamathParserRULE_constantstmt)
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
		p.Match(metamathParserT__2)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0) {
		{
			p.SetState(60)
			p.Constant()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StmtContext) Variablestmt() IVariablestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablestmtContext)
}

func (s *StmtContext) Disjointstmt() IDisjointstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjointstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjointstmtContext)
}

func (s *StmtContext) Hypothesisstmt() IHypothesisstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHypothesisstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHypothesisstmtContext)
}

func (s *StmtContext) Assertstmt() IAssertstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertstmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *metamathParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, metamathParserRULE_stmt)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Variablestmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Disjointstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Hypothesisstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Assertstmt()
		}

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
	p.RuleIndex = metamathParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *metamathParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, metamathParserRULE_block)
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
		p.Match(metamathParserT__4)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserT__4)|(1<<metamathParserT__6)|(1<<metamathParserT__7)|(1<<metamathParserLABEL))) != 0 {
		{
			p.SetState(75)
			p.Stmt()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(metamathParserT__5)
	}

	return localctx
}

// IVariablestmtContext is an interface to support dynamic dispatch.
type IVariablestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablestmtContext differentiates from other interfaces.
	IsVariablestmtContext()
}

type VariablestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablestmtContext() *VariablestmtContext {
	var p = new(VariablestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_variablestmt
	return p
}

func (*VariablestmtContext) IsVariablestmtContext() {}

func NewVariablestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablestmtContext {
	var p = new(VariablestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_variablestmt

	return p
}

func (s *VariablestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablestmtContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *VariablestmtContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariablestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterVariablestmt(s)
	}
}

func (s *VariablestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitVariablestmt(s)
	}
}

func (p *metamathParser) Variablestmt() (localctx IVariablestmtContext) {
	this := p
	_ = this

	localctx = NewVariablestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, metamathParserRULE_variablestmt)
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
		p.SetState(83)
		p.Match(metamathParserT__6)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0) {
		{
			p.SetState(84)
			p.Variable()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IDisjointstmtContext is an interface to support dynamic dispatch.
type IDisjointstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjointstmtContext differentiates from other interfaces.
	IsDisjointstmtContext()
}

type DisjointstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjointstmtContext() *DisjointstmtContext {
	var p = new(DisjointstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_disjointstmt
	return p
}

func (*DisjointstmtContext) IsDisjointstmtContext() {}

func NewDisjointstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjointstmtContext {
	var p = new(DisjointstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_disjointstmt

	return p
}

func (s *DisjointstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjointstmtContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *DisjointstmtContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *DisjointstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjointstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjointstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterDisjointstmt(s)
	}
}

func (s *DisjointstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitDisjointstmt(s)
	}
}

func (p *metamathParser) Disjointstmt() (localctx IDisjointstmtContext) {
	this := p
	_ = this

	localctx = NewDisjointstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, metamathParserRULE_disjointstmt)
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
		p.SetState(91)
		p.Match(metamathParserT__7)
	}
	{
		p.SetState(92)
		p.Variable()
	}
	{
		p.SetState(93)
		p.Variable()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0 {
		{
			p.SetState(94)
			p.Variable()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IHypothesisstmtContext is an interface to support dynamic dispatch.
type IHypothesisstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHypothesisstmtContext differentiates from other interfaces.
	IsHypothesisstmtContext()
}

type HypothesisstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHypothesisstmtContext() *HypothesisstmtContext {
	var p = new(HypothesisstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_hypothesisstmt
	return p
}

func (*HypothesisstmtContext) IsHypothesisstmtContext() {}

func NewHypothesisstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HypothesisstmtContext {
	var p = new(HypothesisstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_hypothesisstmt

	return p
}

func (s *HypothesisstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *HypothesisstmtContext) Floatingstmt() IFloatingstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatingstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatingstmtContext)
}

func (s *HypothesisstmtContext) Essentialstmt() IEssentialstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEssentialstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEssentialstmtContext)
}

func (s *HypothesisstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HypothesisstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HypothesisstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterHypothesisstmt(s)
	}
}

func (s *HypothesisstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitHypothesisstmt(s)
	}
}

func (p *metamathParser) Hypothesisstmt() (localctx IHypothesisstmtContext) {
	this := p
	_ = this

	localctx = NewHypothesisstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, metamathParserRULE_hypothesisstmt)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Floatingstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Essentialstmt()
		}

	}

	return localctx
}

// IFloatingstmtContext is an interface to support dynamic dispatch.
type IFloatingstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatingstmtContext differentiates from other interfaces.
	IsFloatingstmtContext()
}

type FloatingstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingstmtContext() *FloatingstmtContext {
	var p = new(FloatingstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_floatingstmt
	return p
}

func (*FloatingstmtContext) IsFloatingstmtContext() {}

func NewFloatingstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingstmtContext {
	var p = new(FloatingstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_floatingstmt

	return p
}

func (s *FloatingstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingstmtContext) LABEL() antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, 0)
}

func (s *FloatingstmtContext) Typecode() ITypecodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypecodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypecodeContext)
}

func (s *FloatingstmtContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FloatingstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterFloatingstmt(s)
	}
}

func (s *FloatingstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitFloatingstmt(s)
	}
}

func (p *metamathParser) Floatingstmt() (localctx IFloatingstmtContext) {
	this := p
	_ = this

	localctx = NewFloatingstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, metamathParserRULE_floatingstmt)

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
		p.SetState(106)
		p.Match(metamathParserLABEL)
	}
	{
		p.SetState(107)
		p.Match(metamathParserT__8)
	}
	{
		p.SetState(108)
		p.Typecode()
	}
	{
		p.SetState(109)
		p.Variable()
	}
	{
		p.SetState(110)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IEssentialstmtContext is an interface to support dynamic dispatch.
type IEssentialstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEssentialstmtContext differentiates from other interfaces.
	IsEssentialstmtContext()
}

type EssentialstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEssentialstmtContext() *EssentialstmtContext {
	var p = new(EssentialstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_essentialstmt
	return p
}

func (*EssentialstmtContext) IsEssentialstmtContext() {}

func NewEssentialstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EssentialstmtContext {
	var p = new(EssentialstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_essentialstmt

	return p
}

func (s *EssentialstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EssentialstmtContext) LABEL() antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, 0)
}

func (s *EssentialstmtContext) Typecode() ITypecodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypecodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypecodeContext)
}

func (s *EssentialstmtContext) AllMathsymbol() []IMathsymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem())
	var tst = make([]IMathsymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathsymbolContext)
		}
	}

	return tst
}

func (s *EssentialstmtContext) Mathsymbol(i int) IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *EssentialstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EssentialstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EssentialstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterEssentialstmt(s)
	}
}

func (s *EssentialstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitEssentialstmt(s)
	}
}

func (p *metamathParser) Essentialstmt() (localctx IEssentialstmtContext) {
	this := p
	_ = this

	localctx = NewEssentialstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, metamathParserRULE_essentialstmt)
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
		p.SetState(112)
		p.Match(metamathParserLABEL)
	}
	{
		p.SetState(113)
		p.Match(metamathParserT__9)
	}
	{
		p.SetState(114)
		p.Typecode()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0 {
		{
			p.SetState(115)
			p.Mathsymbol()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IAssertstmtContext is an interface to support dynamic dispatch.
type IAssertstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssertstmtContext differentiates from other interfaces.
	IsAssertstmtContext()
}

type AssertstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertstmtContext() *AssertstmtContext {
	var p = new(AssertstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_assertstmt
	return p
}

func (*AssertstmtContext) IsAssertstmtContext() {}

func NewAssertstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertstmtContext {
	var p = new(AssertstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_assertstmt

	return p
}

func (s *AssertstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertstmtContext) Axiomstmt() IAxiomstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxiomstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAxiomstmtContext)
}

func (s *AssertstmtContext) Provablestmt() IProvablestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProvablestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProvablestmtContext)
}

func (s *AssertstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterAssertstmt(s)
	}
}

func (s *AssertstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitAssertstmt(s)
	}
}

func (p *metamathParser) Assertstmt() (localctx IAssertstmtContext) {
	this := p
	_ = this

	localctx = NewAssertstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, metamathParserRULE_assertstmt)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Axiomstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Provablestmt()
		}

	}

	return localctx
}

// IAxiomstmtContext is an interface to support dynamic dispatch.
type IAxiomstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxiomstmtContext differentiates from other interfaces.
	IsAxiomstmtContext()
}

type AxiomstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxiomstmtContext() *AxiomstmtContext {
	var p = new(AxiomstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_axiomstmt
	return p
}

func (*AxiomstmtContext) IsAxiomstmtContext() {}

func NewAxiomstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AxiomstmtContext {
	var p = new(AxiomstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_axiomstmt

	return p
}

func (s *AxiomstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AxiomstmtContext) LABEL() antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, 0)
}

func (s *AxiomstmtContext) Typecode() ITypecodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypecodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypecodeContext)
}

func (s *AxiomstmtContext) AllMathsymbol() []IMathsymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem())
	var tst = make([]IMathsymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathsymbolContext)
		}
	}

	return tst
}

func (s *AxiomstmtContext) Mathsymbol(i int) IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *AxiomstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxiomstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AxiomstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterAxiomstmt(s)
	}
}

func (s *AxiomstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitAxiomstmt(s)
	}
}

func (p *metamathParser) Axiomstmt() (localctx IAxiomstmtContext) {
	this := p
	_ = this

	localctx = NewAxiomstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, metamathParserRULE_axiomstmt)
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
		p.Match(metamathParserLABEL)
	}
	{
		p.SetState(128)
		p.Match(metamathParserT__10)
	}
	{
		p.SetState(129)
		p.Typecode()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0 {
		{
			p.SetState(130)
			p.Mathsymbol()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IProvablestmtContext is an interface to support dynamic dispatch.
type IProvablestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProvablestmtContext differentiates from other interfaces.
	IsProvablestmtContext()
}

type ProvablestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProvablestmtContext() *ProvablestmtContext {
	var p = new(ProvablestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_provablestmt
	return p
}

func (*ProvablestmtContext) IsProvablestmtContext() {}

func NewProvablestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProvablestmtContext {
	var p = new(ProvablestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_provablestmt

	return p
}

func (s *ProvablestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ProvablestmtContext) LABEL() antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, 0)
}

func (s *ProvablestmtContext) Typecode() ITypecodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypecodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypecodeContext)
}

func (s *ProvablestmtContext) Proof() IProofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProofContext)
}

func (s *ProvablestmtContext) AllMathsymbol() []IMathsymbolContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem())
	var tst = make([]IMathsymbolContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathsymbolContext)
		}
	}

	return tst
}

func (s *ProvablestmtContext) Mathsymbol(i int) IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *ProvablestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProvablestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProvablestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterProvablestmt(s)
	}
}

func (s *ProvablestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitProvablestmt(s)
	}
}

func (p *metamathParser) Provablestmt() (localctx IProvablestmtContext) {
	this := p
	_ = this

	localctx = NewProvablestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, metamathParserRULE_provablestmt)
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
		p.Match(metamathParserLABEL)
	}
	{
		p.SetState(139)
		p.Match(metamathParserT__11)
	}
	{
		p.SetState(140)
		p.Typecode()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0 {
		{
			p.SetState(141)
			p.Mathsymbol()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(metamathParserT__12)
	}
	{
		p.SetState(148)
		p.Proof()
	}
	{
		p.SetState(149)
		p.Match(metamathParserT__3)
	}

	return localctx
}

// IProofContext is an interface to support dynamic dispatch.
type IProofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProofContext differentiates from other interfaces.
	IsProofContext()
}

type ProofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProofContext() *ProofContext {
	var p = new(ProofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_proof
	return p
}

func (*ProofContext) IsProofContext() {}

func NewProofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProofContext {
	var p = new(ProofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_proof

	return p
}

func (s *ProofContext) GetParser() antlr.Parser { return s.parser }

func (s *ProofContext) Uncompressedproof() IUncompressedproofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUncompressedproofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUncompressedproofContext)
}

func (s *ProofContext) Compressedproof() ICompressedproofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompressedproofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompressedproofContext)
}

func (s *ProofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterProof(s)
	}
}

func (s *ProofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitProof(s)
	}
}

func (p *metamathParser) Proof() (localctx IProofContext) {
	this := p
	_ = this

	localctx = NewProofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, metamathParserRULE_proof)

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

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case metamathParserT__13, metamathParserLABEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Uncompressedproof()
		}

	case metamathParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Compressedproof()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUncompressedproofContext is an interface to support dynamic dispatch.
type IUncompressedproofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUncompressedproofContext differentiates from other interfaces.
	IsUncompressedproofContext()
}

type UncompressedproofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUncompressedproofContext() *UncompressedproofContext {
	var p = new(UncompressedproofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_uncompressedproof
	return p
}

func (*UncompressedproofContext) IsUncompressedproofContext() {}

func NewUncompressedproofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UncompressedproofContext {
	var p = new(UncompressedproofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_uncompressedproof

	return p
}

func (s *UncompressedproofContext) GetParser() antlr.Parser { return s.parser }

func (s *UncompressedproofContext) AllLABEL() []antlr.TerminalNode {
	return s.GetTokens(metamathParserLABEL)
}

func (s *UncompressedproofContext) LABEL(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, i)
}

func (s *UncompressedproofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UncompressedproofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UncompressedproofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterUncompressedproof(s)
	}
}

func (s *UncompressedproofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitUncompressedproof(s)
	}
}

func (p *metamathParser) Uncompressedproof() (localctx IUncompressedproofContext) {
	this := p
	_ = this

	localctx = NewUncompressedproofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, metamathParserRULE_uncompressedproof)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == metamathParserT__13 || _la == metamathParserLABEL {
		{
			p.SetState(155)
			_la = p.GetTokenStream().LA(1)

			if !(_la == metamathParserT__13 || _la == metamathParserLABEL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICompressedproofContext is an interface to support dynamic dispatch.
type ICompressedproofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompressedproofContext differentiates from other interfaces.
	IsCompressedproofContext()
}

type CompressedproofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompressedproofContext() *CompressedproofContext {
	var p = new(CompressedproofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_compressedproof
	return p
}

func (*CompressedproofContext) IsCompressedproofContext() {}

func NewCompressedproofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompressedproofContext {
	var p = new(CompressedproofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_compressedproof

	return p
}

func (s *CompressedproofContext) GetParser() antlr.Parser { return s.parser }

func (s *CompressedproofContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(metamathParserLPAREN, 0)
}

func (s *CompressedproofContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(metamathParserRPAREN, 0)
}

func (s *CompressedproofContext) AllLABEL() []antlr.TerminalNode {
	return s.GetTokens(metamathParserLABEL)
}

func (s *CompressedproofContext) LABEL(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserLABEL, i)
}

func (s *CompressedproofContext) AllCOMPRESSEDPROOFBLOCK() []antlr.TerminalNode {
	return s.GetTokens(metamathParserCOMPRESSEDPROOFBLOCK)
}

func (s *CompressedproofContext) COMPRESSEDPROOFBLOCK(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserCOMPRESSEDPROOFBLOCK, i)
}

func (s *CompressedproofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompressedproofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompressedproofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterCompressedproof(s)
	}
}

func (s *CompressedproofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitCompressedproof(s)
	}
}

func (p *metamathParser) Compressedproof() (localctx ICompressedproofContext) {
	this := p
	_ = this

	localctx = NewCompressedproofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, metamathParserRULE_compressedproof)
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
		p.Match(metamathParserLPAREN)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == metamathParserLABEL {
		{
			p.SetState(161)
			p.Match(metamathParserLABEL)
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(metamathParserRPAREN)
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == metamathParserCOMPRESSEDPROOFBLOCK {
		{
			p.SetState(168)
			p.Match(metamathParserCOMPRESSEDPROOFBLOCK)
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypecodeContext is an interface to support dynamic dispatch.
type ITypecodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypecodeContext differentiates from other interfaces.
	IsTypecodeContext()
}

type TypecodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypecodeContext() *TypecodeContext {
	var p = new(TypecodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_typecode
	return p
}

func (*TypecodeContext) IsTypecodeContext() {}

func NewTypecodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypecodeContext {
	var p = new(TypecodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_typecode

	return p
}

func (s *TypecodeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypecodeContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *TypecodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypecodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypecodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterTypecode(s)
	}
}

func (s *TypecodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitTypecode(s)
	}
}

func (p *metamathParser) Typecode() (localctx ITypecodeContext) {
	this := p
	_ = this

	localctx = NewTypecodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, metamathParserRULE_typecode)

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
		p.SetState(173)
		p.Constant()
	}

	return localctx
}

// IMathsymbolContext is an interface to support dynamic dispatch.
type IMathsymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathsymbolContext differentiates from other interfaces.
	IsMathsymbolContext()
}

type MathsymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathsymbolContext() *MathsymbolContext {
	var p = new(MathsymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_mathsymbol
	return p
}

func (*MathsymbolContext) IsMathsymbolContext() {}

func NewMathsymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathsymbolContext {
	var p = new(MathsymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_mathsymbol

	return p
}

func (s *MathsymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *MathsymbolContext) AllPRINTABLECHARACTER() []antlr.TerminalNode {
	return s.GetTokens(metamathParserPRINTABLECHARACTER)
}

func (s *MathsymbolContext) PRINTABLECHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserPRINTABLECHARACTER, i)
}

func (s *MathsymbolContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(metamathParserLPAREN)
}

func (s *MathsymbolContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserLPAREN, i)
}

func (s *MathsymbolContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(metamathParserRPAREN)
}

func (s *MathsymbolContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(metamathParserRPAREN, i)
}

func (s *MathsymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathsymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathsymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterMathsymbol(s)
	}
}

func (s *MathsymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitMathsymbol(s)
	}
}

func (p *metamathParser) Mathsymbol() (localctx IMathsymbolContext) {
	this := p
	_ = this

	localctx = NewMathsymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, metamathParserRULE_mathsymbol)
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
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(175)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<metamathParserLPAREN)|(1<<metamathParserRPAREN)|(1<<metamathParserPRINTABLECHARACTER))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IFilenameContext is an interface to support dynamic dispatch.
type IFilenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilenameContext differentiates from other interfaces.
	IsFilenameContext()
}

type FilenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilenameContext() *FilenameContext {
	var p = new(FilenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_filename
	return p
}

func (*FilenameContext) IsFilenameContext() {}

func NewFilenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilenameContext {
	var p = new(FilenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_filename

	return p
}

func (s *FilenameContext) GetParser() antlr.Parser { return s.parser }

func (s *FilenameContext) Mathsymbol() IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *FilenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterFilename(s)
	}
}

func (s *FilenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitFilename(s)
	}
}

func (p *metamathParser) Filename() (localctx IFilenameContext) {
	this := p
	_ = this

	localctx = NewFilenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, metamathParserRULE_filename)

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
		p.SetState(180)
		p.Mathsymbol()
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = metamathParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Mathsymbol() IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *metamathParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, metamathParserRULE_constant)

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
		p.SetState(182)
		p.Mathsymbol()
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
	p.RuleIndex = metamathParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = metamathParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Mathsymbol() IMathsymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathsymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathsymbolContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(metamathListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *metamathParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, metamathParserRULE_variable)

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
		p.SetState(184)
		p.Mathsymbol()
	}

	return localctx
}
