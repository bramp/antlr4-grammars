// Code generated from mps.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mps // mps
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 150,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 48, 10, 2, 3, 2, 5, 2, 51, 10, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	5, 3, 58, 10, 3, 3, 3, 5, 3, 61, 10, 3, 3, 4, 3, 4, 6, 4, 65, 10, 4, 13,
	4, 14, 4, 66, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 6, 11, 88,
	10, 11, 13, 11, 14, 11, 89, 3, 12, 6, 12, 93, 10, 12, 13, 12, 14, 12, 94,
	3, 13, 6, 13, 98, 10, 13, 13, 13, 14, 13, 99, 3, 14, 6, 14, 103, 10, 14,
	13, 14, 14, 14, 104, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 112, 10,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 119, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 5, 17, 126, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 132, 10, 18, 3, 19, 3, 19, 6, 19, 136, 10, 19, 13, 19, 14, 19, 137,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 2, 5, 4, 2, 6, 6, 16, 16, 4, 2, 7, 7, 16, 16, 4, 2, 8,
	8, 16, 16, 2, 144, 2, 42, 3, 2, 2, 2, 4, 55, 3, 2, 2, 2, 6, 62, 3, 2, 2,
	2, 8, 68, 3, 2, 2, 2, 10, 71, 3, 2, 2, 2, 12, 74, 3, 2, 2, 2, 14, 77, 3,
	2, 2, 2, 16, 80, 3, 2, 2, 2, 18, 82, 3, 2, 2, 2, 20, 87, 3, 2, 2, 2, 22,
	92, 3, 2, 2, 2, 24, 97, 3, 2, 2, 2, 26, 102, 3, 2, 2, 2, 28, 106, 3, 2,
	2, 2, 30, 113, 3, 2, 2, 2, 32, 120, 3, 2, 2, 2, 34, 127, 3, 2, 2, 2, 36,
	133, 3, 2, 2, 2, 38, 141, 3, 2, 2, 2, 40, 145, 3, 2, 2, 2, 42, 43, 5, 4,
	3, 2, 43, 44, 5, 6, 4, 2, 44, 45, 5, 8, 5, 2, 45, 47, 5, 10, 6, 2, 46,
	48, 5, 12, 7, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2,
	2, 2, 49, 51, 5, 14, 8, 2, 50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 53, 5, 16, 9, 2, 53, 54, 7, 2, 2, 3, 54, 3, 3, 2, 2,
	2, 55, 57, 7, 3, 2, 2, 56, 58, 7, 16, 2, 2, 57, 56, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59, 61, 7, 13, 2, 2, 60, 59, 3, 2, 2, 2,
	60, 61, 3, 2, 2, 2, 61, 5, 3, 2, 2, 2, 62, 64, 7, 4, 2, 2, 63, 65, 5, 18,
	10, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66,
	67, 3, 2, 2, 2, 67, 7, 3, 2, 2, 2, 68, 69, 7, 5, 2, 2, 69, 70, 5, 20, 11,
	2, 70, 9, 3, 2, 2, 2, 71, 72, 7, 6, 2, 2, 72, 73, 5, 22, 12, 2, 73, 11,
	3, 2, 2, 2, 74, 75, 7, 7, 2, 2, 75, 76, 5, 24, 13, 2, 76, 13, 3, 2, 2,
	2, 77, 78, 7, 8, 2, 2, 78, 79, 5, 26, 14, 2, 79, 15, 3, 2, 2, 2, 80, 81,
	7, 9, 2, 2, 81, 17, 3, 2, 2, 2, 82, 83, 7, 15, 2, 2, 83, 84, 7, 16, 2,
	2, 84, 19, 3, 2, 2, 2, 85, 88, 5, 28, 15, 2, 86, 88, 5, 36, 19, 2, 87,
	85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 21, 3, 2, 2, 2, 91, 93, 5, 30, 16, 2, 92, 91,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 23, 3, 2, 2, 2, 96, 98, 5, 32, 17, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 25, 3, 2, 2, 2,
	101, 103, 5, 34, 18, 2, 102, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 27, 3, 2, 2, 2, 106, 107, 7,
	16, 2, 2, 107, 108, 7, 16, 2, 2, 108, 111, 7, 17, 2, 2, 109, 110, 7, 16,
	2, 2, 110, 112, 7, 17, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2,
	112, 29, 3, 2, 2, 2, 113, 114, 9, 2, 2, 2, 114, 115, 7, 16, 2, 2, 115,
	118, 7, 17, 2, 2, 116, 117, 7, 16, 2, 2, 117, 119, 7, 17, 2, 2, 118, 116,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 31, 3, 2, 2, 2, 120, 121, 9, 3,
	2, 2, 121, 122, 7, 16, 2, 2, 122, 125, 7, 17, 2, 2, 123, 124, 7, 16, 2,
	2, 124, 126, 7, 17, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	33, 3, 2, 2, 2, 127, 128, 7, 14, 2, 2, 128, 129, 9, 4, 2, 2, 129, 131,
	7, 16, 2, 2, 130, 132, 7, 17, 2, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3,
	2, 2, 2, 132, 35, 3, 2, 2, 2, 133, 135, 5, 38, 20, 2, 134, 136, 5, 28,
	15, 2, 135, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2,
	137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 5, 40, 21, 2, 140,
	37, 3, 2, 2, 2, 141, 142, 7, 16, 2, 2, 142, 143, 7, 10, 2, 2, 143, 144,
	7, 11, 2, 2, 144, 39, 3, 2, 2, 2, 145, 146, 7, 16, 2, 2, 146, 147, 7, 10,
	2, 2, 147, 148, 7, 12, 2, 2, 148, 41, 3, 2, 2, 2, 17, 47, 50, 57, 60, 66,
	87, 89, 94, 99, 104, 111, 118, 125, 131, 137,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'NAME'", "'ROWS'", "'COLUMNS'", "'RHS'", "'RANGES'", "'BOUNDS'", "'ENDATA'",
	"''MARKER''", "''INTORG''", "''INTEND''", "'FREE'",
}
var symbolicNames = []string{
	"", "NAMEINDICATORCARD", "ROWINDICATORCARD", "COLUMNINDICATORCARD", "RHSINDICATORCARD",
	"RANGESINDICATORCARD", "BOUNDSINDICATORCARD", "ENDATAINDICATORCARD", "KEYWORDMARKER",
	"STARTMARKER", "ENDMARKER", "KEYWORDFREE", "BOUNDKEY", "ROWTYPE", "IDENTIFIER",
	"NUMERICALVALUE", "WS", "LINE_COMMENT",
}

var ruleNames = []string{
	"modell", "firstrow", "rows", "columns", "rhs", "ranges", "bounds", "endata",
	"rowdatacard", "columndatacards", "rhsdatacards", "rangesdatacards", "boundsdatacards",
	"columndatacard", "rhsdatacard", "rangesdatacard", "boundsdatacard", "intblock",
	"startmarker", "endmarker",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type mpsParser struct {
	*antlr.BaseParser
}

func NewmpsParser(input antlr.TokenStream) *mpsParser {
	this := new(mpsParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "mps.g4"

	return this
}

// mpsParser tokens.
const (
	mpsParserEOF                 = antlr.TokenEOF
	mpsParserNAMEINDICATORCARD   = 1
	mpsParserROWINDICATORCARD    = 2
	mpsParserCOLUMNINDICATORCARD = 3
	mpsParserRHSINDICATORCARD    = 4
	mpsParserRANGESINDICATORCARD = 5
	mpsParserBOUNDSINDICATORCARD = 6
	mpsParserENDATAINDICATORCARD = 7
	mpsParserKEYWORDMARKER       = 8
	mpsParserSTARTMARKER         = 9
	mpsParserENDMARKER           = 10
	mpsParserKEYWORDFREE         = 11
	mpsParserBOUNDKEY            = 12
	mpsParserROWTYPE             = 13
	mpsParserIDENTIFIER          = 14
	mpsParserNUMERICALVALUE      = 15
	mpsParserWS                  = 16
	mpsParserLINE_COMMENT        = 17
)

// mpsParser rules.
const (
	mpsParserRULE_modell          = 0
	mpsParserRULE_firstrow        = 1
	mpsParserRULE_rows            = 2
	mpsParserRULE_columns         = 3
	mpsParserRULE_rhs             = 4
	mpsParserRULE_ranges          = 5
	mpsParserRULE_bounds          = 6
	mpsParserRULE_endata          = 7
	mpsParserRULE_rowdatacard     = 8
	mpsParserRULE_columndatacards = 9
	mpsParserRULE_rhsdatacards    = 10
	mpsParserRULE_rangesdatacards = 11
	mpsParserRULE_boundsdatacards = 12
	mpsParserRULE_columndatacard  = 13
	mpsParserRULE_rhsdatacard     = 14
	mpsParserRULE_rangesdatacard  = 15
	mpsParserRULE_boundsdatacard  = 16
	mpsParserRULE_intblock        = 17
	mpsParserRULE_startmarker     = 18
	mpsParserRULE_endmarker       = 19
)

// IModellContext is an interface to support dynamic dispatch.
type IModellContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModellContext differentiates from other interfaces.
	IsModellContext()
}

type ModellContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModellContext() *ModellContext {
	var p = new(ModellContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_modell
	return p
}

func (*ModellContext) IsModellContext() {}

func NewModellContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModellContext {
	var p = new(ModellContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_modell

	return p
}

func (s *ModellContext) GetParser() antlr.Parser { return s.parser }

func (s *ModellContext) Firstrow() IFirstrowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFirstrowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFirstrowContext)
}

func (s *ModellContext) Rows() IRowsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowsContext)
}

func (s *ModellContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *ModellContext) Rhs() IRhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRhsContext)
}

func (s *ModellContext) Endata() IEndataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndataContext)
}

func (s *ModellContext) EOF() antlr.TerminalNode {
	return s.GetToken(mpsParserEOF, 0)
}

func (s *ModellContext) Ranges() IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ModellContext) Bounds() IBoundsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundsContext)
}

func (s *ModellContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModellContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModellContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterModell(s)
	}
}

func (s *ModellContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitModell(s)
	}
}

func (p *mpsParser) Modell() (localctx IModellContext) {
	localctx = NewModellContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mpsParserRULE_modell)
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
		p.SetState(40)
		p.Firstrow()
	}
	{
		p.SetState(41)
		p.Rows()
	}
	{
		p.SetState(42)
		p.Columns()
	}
	{
		p.SetState(43)
		p.Rhs()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mpsParserRANGESINDICATORCARD {
		{
			p.SetState(44)
			p.Ranges()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mpsParserBOUNDSINDICATORCARD {
		{
			p.SetState(47)
			p.Bounds()
		}

	}
	{
		p.SetState(50)
		p.Endata()
	}
	{
		p.SetState(51)
		p.Match(mpsParserEOF)
	}

	return localctx
}

// IFirstrowContext is an interface to support dynamic dispatch.
type IFirstrowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFirstrowContext differentiates from other interfaces.
	IsFirstrowContext()
}

type FirstrowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFirstrowContext() *FirstrowContext {
	var p = new(FirstrowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_firstrow
	return p
}

func (*FirstrowContext) IsFirstrowContext() {}

func NewFirstrowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FirstrowContext {
	var p = new(FirstrowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_firstrow

	return p
}

func (s *FirstrowContext) GetParser() antlr.Parser { return s.parser }

func (s *FirstrowContext) NAMEINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserNAMEINDICATORCARD, 0)
}

func (s *FirstrowContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, 0)
}

func (s *FirstrowContext) KEYWORDFREE() antlr.TerminalNode {
	return s.GetToken(mpsParserKEYWORDFREE, 0)
}

func (s *FirstrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FirstrowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FirstrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterFirstrow(s)
	}
}

func (s *FirstrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitFirstrow(s)
	}
}

func (p *mpsParser) Firstrow() (localctx IFirstrowContext) {
	localctx = NewFirstrowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mpsParserRULE_firstrow)
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
		p.SetState(53)
		p.Match(mpsParserNAMEINDICATORCARD)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mpsParserIDENTIFIER {
		{
			p.SetState(54)
			p.Match(mpsParserIDENTIFIER)
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mpsParserKEYWORDFREE {
		{
			p.SetState(57)
			p.Match(mpsParserKEYWORDFREE)
		}

	}

	return localctx
}

// IRowsContext is an interface to support dynamic dispatch.
type IRowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowsContext differentiates from other interfaces.
	IsRowsContext()
}

type RowsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowsContext() *RowsContext {
	var p = new(RowsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rows
	return p
}

func (*RowsContext) IsRowsContext() {}

func NewRowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowsContext {
	var p = new(RowsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rows

	return p
}

func (s *RowsContext) GetParser() antlr.Parser { return s.parser }

func (s *RowsContext) ROWINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserROWINDICATORCARD, 0)
}

func (s *RowsContext) AllRowdatacard() []IRowdatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowdatacardContext)(nil)).Elem())
	var tst = make([]IRowdatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowdatacardContext)
		}
	}

	return tst
}

func (s *RowsContext) Rowdatacard(i int) IRowdatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowdatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowdatacardContext)
}

func (s *RowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRows(s)
	}
}

func (s *RowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRows(s)
	}
}

func (p *mpsParser) Rows() (localctx IRowsContext) {
	localctx = NewRowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mpsParserRULE_rows)
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
		p.Match(mpsParserROWINDICATORCARD)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mpsParserROWTYPE {
		{
			p.SetState(61)
			p.Rowdatacard()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) COLUMNINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserCOLUMNINDICATORCARD, 0)
}

func (s *ColumnsContext) Columndatacards() IColumndatacardsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumndatacardsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumndatacardsContext)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *mpsParser) Columns() (localctx IColumnsContext) {
	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mpsParserRULE_columns)

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
		p.SetState(66)
		p.Match(mpsParserCOLUMNINDICATORCARD)
	}
	{
		p.SetState(67)
		p.Columndatacards()
	}

	return localctx
}

// IRhsContext is an interface to support dynamic dispatch.
type IRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsContext differentiates from other interfaces.
	IsRhsContext()
}

type RhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsContext() *RhsContext {
	var p = new(RhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rhs
	return p
}

func (*RhsContext) IsRhsContext() {}

func NewRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsContext {
	var p = new(RhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rhs

	return p
}

func (s *RhsContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsContext) RHSINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserRHSINDICATORCARD, 0)
}

func (s *RhsContext) Rhsdatacards() IRhsdatacardsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRhsdatacardsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRhsdatacardsContext)
}

func (s *RhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRhs(s)
	}
}

func (s *RhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRhs(s)
	}
}

func (p *mpsParser) Rhs() (localctx IRhsContext) {
	localctx = NewRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mpsParserRULE_rhs)

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
		p.SetState(69)
		p.Match(mpsParserRHSINDICATORCARD)
	}
	{
		p.SetState(70)
		p.Rhsdatacards()
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) RANGESINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserRANGESINDICATORCARD, 0)
}

func (s *RangesContext) Rangesdatacards() IRangesdatacardsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesdatacardsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangesdatacardsContext)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *mpsParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mpsParserRULE_ranges)

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
		p.Match(mpsParserRANGESINDICATORCARD)
	}
	{
		p.SetState(73)
		p.Rangesdatacards()
	}

	return localctx
}

// IBoundsContext is an interface to support dynamic dispatch.
type IBoundsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundsContext differentiates from other interfaces.
	IsBoundsContext()
}

type BoundsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundsContext() *BoundsContext {
	var p = new(BoundsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_bounds
	return p
}

func (*BoundsContext) IsBoundsContext() {}

func NewBoundsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundsContext {
	var p = new(BoundsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_bounds

	return p
}

func (s *BoundsContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundsContext) BOUNDSINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserBOUNDSINDICATORCARD, 0)
}

func (s *BoundsContext) Boundsdatacards() IBoundsdatacardsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundsdatacardsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundsdatacardsContext)
}

func (s *BoundsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterBounds(s)
	}
}

func (s *BoundsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitBounds(s)
	}
}

func (p *mpsParser) Bounds() (localctx IBoundsContext) {
	localctx = NewBoundsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mpsParserRULE_bounds)

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
		p.Match(mpsParserBOUNDSINDICATORCARD)
	}
	{
		p.SetState(76)
		p.Boundsdatacards()
	}

	return localctx
}

// IEndataContext is an interface to support dynamic dispatch.
type IEndataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndataContext differentiates from other interfaces.
	IsEndataContext()
}

type EndataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndataContext() *EndataContext {
	var p = new(EndataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_endata
	return p
}

func (*EndataContext) IsEndataContext() {}

func NewEndataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndataContext {
	var p = new(EndataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_endata

	return p
}

func (s *EndataContext) GetParser() antlr.Parser { return s.parser }

func (s *EndataContext) ENDATAINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserENDATAINDICATORCARD, 0)
}

func (s *EndataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterEndata(s)
	}
}

func (s *EndataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitEndata(s)
	}
}

func (p *mpsParser) Endata() (localctx IEndataContext) {
	localctx = NewEndataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mpsParserRULE_endata)

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
		p.Match(mpsParserENDATAINDICATORCARD)
	}

	return localctx
}

// IRowdatacardContext is an interface to support dynamic dispatch.
type IRowdatacardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowdatacardContext differentiates from other interfaces.
	IsRowdatacardContext()
}

type RowdatacardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowdatacardContext() *RowdatacardContext {
	var p = new(RowdatacardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rowdatacard
	return p
}

func (*RowdatacardContext) IsRowdatacardContext() {}

func NewRowdatacardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowdatacardContext {
	var p = new(RowdatacardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rowdatacard

	return p
}

func (s *RowdatacardContext) GetParser() antlr.Parser { return s.parser }

func (s *RowdatacardContext) ROWTYPE() antlr.TerminalNode {
	return s.GetToken(mpsParserROWTYPE, 0)
}

func (s *RowdatacardContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, 0)
}

func (s *RowdatacardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowdatacardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowdatacardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRowdatacard(s)
	}
}

func (s *RowdatacardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRowdatacard(s)
	}
}

func (p *mpsParser) Rowdatacard() (localctx IRowdatacardContext) {
	localctx = NewRowdatacardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mpsParserRULE_rowdatacard)

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
		p.Match(mpsParserROWTYPE)
	}
	{
		p.SetState(81)
		p.Match(mpsParserIDENTIFIER)
	}

	return localctx
}

// IColumndatacardsContext is an interface to support dynamic dispatch.
type IColumndatacardsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumndatacardsContext differentiates from other interfaces.
	IsColumndatacardsContext()
}

type ColumndatacardsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumndatacardsContext() *ColumndatacardsContext {
	var p = new(ColumndatacardsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_columndatacards
	return p
}

func (*ColumndatacardsContext) IsColumndatacardsContext() {}

func NewColumndatacardsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumndatacardsContext {
	var p = new(ColumndatacardsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_columndatacards

	return p
}

func (s *ColumndatacardsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumndatacardsContext) AllColumndatacard() []IColumndatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumndatacardContext)(nil)).Elem())
	var tst = make([]IColumndatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumndatacardContext)
		}
	}

	return tst
}

func (s *ColumndatacardsContext) Columndatacard(i int) IColumndatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumndatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumndatacardContext)
}

func (s *ColumndatacardsContext) AllIntblock() []IIntblockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntblockContext)(nil)).Elem())
	var tst = make([]IIntblockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntblockContext)
		}
	}

	return tst
}

func (s *ColumndatacardsContext) Intblock(i int) IIntblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntblockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntblockContext)
}

func (s *ColumndatacardsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumndatacardsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumndatacardsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterColumndatacards(s)
	}
}

func (s *ColumndatacardsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitColumndatacards(s)
	}
}

func (p *mpsParser) Columndatacards() (localctx IColumndatacardsContext) {
	localctx = NewColumndatacardsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mpsParserRULE_columndatacards)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mpsParserIDENTIFIER {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(83)
				p.Columndatacard()
			}

		case 2:
			{
				p.SetState(84)
				p.Intblock()
			}

		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRhsdatacardsContext is an interface to support dynamic dispatch.
type IRhsdatacardsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsdatacardsContext differentiates from other interfaces.
	IsRhsdatacardsContext()
}

type RhsdatacardsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsdatacardsContext() *RhsdatacardsContext {
	var p = new(RhsdatacardsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rhsdatacards
	return p
}

func (*RhsdatacardsContext) IsRhsdatacardsContext() {}

func NewRhsdatacardsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsdatacardsContext {
	var p = new(RhsdatacardsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rhsdatacards

	return p
}

func (s *RhsdatacardsContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsdatacardsContext) AllRhsdatacard() []IRhsdatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRhsdatacardContext)(nil)).Elem())
	var tst = make([]IRhsdatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRhsdatacardContext)
		}
	}

	return tst
}

func (s *RhsdatacardsContext) Rhsdatacard(i int) IRhsdatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRhsdatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRhsdatacardContext)
}

func (s *RhsdatacardsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsdatacardsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsdatacardsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRhsdatacards(s)
	}
}

func (s *RhsdatacardsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRhsdatacards(s)
	}
}

func (p *mpsParser) Rhsdatacards() (localctx IRhsdatacardsContext) {
	localctx = NewRhsdatacardsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, mpsParserRULE_rhsdatacards)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mpsParserRHSINDICATORCARD || _la == mpsParserIDENTIFIER {
		{
			p.SetState(89)
			p.Rhsdatacard()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRangesdatacardsContext is an interface to support dynamic dispatch.
type IRangesdatacardsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesdatacardsContext differentiates from other interfaces.
	IsRangesdatacardsContext()
}

type RangesdatacardsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesdatacardsContext() *RangesdatacardsContext {
	var p = new(RangesdatacardsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rangesdatacards
	return p
}

func (*RangesdatacardsContext) IsRangesdatacardsContext() {}

func NewRangesdatacardsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesdatacardsContext {
	var p = new(RangesdatacardsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rangesdatacards

	return p
}

func (s *RangesdatacardsContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesdatacardsContext) AllRangesdatacard() []IRangesdatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRangesdatacardContext)(nil)).Elem())
	var tst = make([]IRangesdatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRangesdatacardContext)
		}
	}

	return tst
}

func (s *RangesdatacardsContext) Rangesdatacard(i int) IRangesdatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesdatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRangesdatacardContext)
}

func (s *RangesdatacardsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesdatacardsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesdatacardsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRangesdatacards(s)
	}
}

func (s *RangesdatacardsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRangesdatacards(s)
	}
}

func (p *mpsParser) Rangesdatacards() (localctx IRangesdatacardsContext) {
	localctx = NewRangesdatacardsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mpsParserRULE_rangesdatacards)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mpsParserRANGESINDICATORCARD || _la == mpsParserIDENTIFIER {
		{
			p.SetState(94)
			p.Rangesdatacard()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoundsdatacardsContext is an interface to support dynamic dispatch.
type IBoundsdatacardsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundsdatacardsContext differentiates from other interfaces.
	IsBoundsdatacardsContext()
}

type BoundsdatacardsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundsdatacardsContext() *BoundsdatacardsContext {
	var p = new(BoundsdatacardsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_boundsdatacards
	return p
}

func (*BoundsdatacardsContext) IsBoundsdatacardsContext() {}

func NewBoundsdatacardsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundsdatacardsContext {
	var p = new(BoundsdatacardsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_boundsdatacards

	return p
}

func (s *BoundsdatacardsContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundsdatacardsContext) AllBoundsdatacard() []IBoundsdatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoundsdatacardContext)(nil)).Elem())
	var tst = make([]IBoundsdatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoundsdatacardContext)
		}
	}

	return tst
}

func (s *BoundsdatacardsContext) Boundsdatacard(i int) IBoundsdatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundsdatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoundsdatacardContext)
}

func (s *BoundsdatacardsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundsdatacardsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundsdatacardsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterBoundsdatacards(s)
	}
}

func (s *BoundsdatacardsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitBoundsdatacards(s)
	}
}

func (p *mpsParser) Boundsdatacards() (localctx IBoundsdatacardsContext) {
	localctx = NewBoundsdatacardsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mpsParserRULE_boundsdatacards)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mpsParserBOUNDKEY {
		{
			p.SetState(99)
			p.Boundsdatacard()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumndatacardContext is an interface to support dynamic dispatch.
type IColumndatacardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumndatacardContext differentiates from other interfaces.
	IsColumndatacardContext()
}

type ColumndatacardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumndatacardContext() *ColumndatacardContext {
	var p = new(ColumndatacardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_columndatacard
	return p
}

func (*ColumndatacardContext) IsColumndatacardContext() {}

func NewColumndatacardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumndatacardContext {
	var p = new(ColumndatacardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_columndatacard

	return p
}

func (s *ColumndatacardContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumndatacardContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(mpsParserIDENTIFIER)
}

func (s *ColumndatacardContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, i)
}

func (s *ColumndatacardContext) AllNUMERICALVALUE() []antlr.TerminalNode {
	return s.GetTokens(mpsParserNUMERICALVALUE)
}

func (s *ColumndatacardContext) NUMERICALVALUE(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserNUMERICALVALUE, i)
}

func (s *ColumndatacardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumndatacardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumndatacardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterColumndatacard(s)
	}
}

func (s *ColumndatacardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitColumndatacard(s)
	}
}

func (p *mpsParser) Columndatacard() (localctx IColumndatacardContext) {
	localctx = NewColumndatacardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mpsParserRULE_columndatacard)

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
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(105)
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(106)
		p.Match(mpsParserNUMERICALVALUE)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
			p.Match(mpsParserIDENTIFIER)
		}
		{
			p.SetState(108)
			p.Match(mpsParserNUMERICALVALUE)
		}

	}

	return localctx
}

// IRhsdatacardContext is an interface to support dynamic dispatch.
type IRhsdatacardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsdatacardContext differentiates from other interfaces.
	IsRhsdatacardContext()
}

type RhsdatacardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsdatacardContext() *RhsdatacardContext {
	var p = new(RhsdatacardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rhsdatacard
	return p
}

func (*RhsdatacardContext) IsRhsdatacardContext() {}

func NewRhsdatacardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsdatacardContext {
	var p = new(RhsdatacardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rhsdatacard

	return p
}

func (s *RhsdatacardContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsdatacardContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(mpsParserIDENTIFIER)
}

func (s *RhsdatacardContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, i)
}

func (s *RhsdatacardContext) AllNUMERICALVALUE() []antlr.TerminalNode {
	return s.GetTokens(mpsParserNUMERICALVALUE)
}

func (s *RhsdatacardContext) NUMERICALVALUE(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserNUMERICALVALUE, i)
}

func (s *RhsdatacardContext) RHSINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserRHSINDICATORCARD, 0)
}

func (s *RhsdatacardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsdatacardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsdatacardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRhsdatacard(s)
	}
}

func (s *RhsdatacardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRhsdatacard(s)
	}
}

func (p *mpsParser) Rhsdatacard() (localctx IRhsdatacardContext) {
	localctx = NewRhsdatacardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mpsParserRULE_rhsdatacard)
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
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mpsParserRHSINDICATORCARD || _la == mpsParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(112)
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(113)
		p.Match(mpsParserNUMERICALVALUE)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.Match(mpsParserIDENTIFIER)
		}
		{
			p.SetState(115)
			p.Match(mpsParserNUMERICALVALUE)
		}

	}

	return localctx
}

// IRangesdatacardContext is an interface to support dynamic dispatch.
type IRangesdatacardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesdatacardContext differentiates from other interfaces.
	IsRangesdatacardContext()
}

type RangesdatacardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesdatacardContext() *RangesdatacardContext {
	var p = new(RangesdatacardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_rangesdatacard
	return p
}

func (*RangesdatacardContext) IsRangesdatacardContext() {}

func NewRangesdatacardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesdatacardContext {
	var p = new(RangesdatacardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_rangesdatacard

	return p
}

func (s *RangesdatacardContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesdatacardContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(mpsParserIDENTIFIER)
}

func (s *RangesdatacardContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, i)
}

func (s *RangesdatacardContext) AllNUMERICALVALUE() []antlr.TerminalNode {
	return s.GetTokens(mpsParserNUMERICALVALUE)
}

func (s *RangesdatacardContext) NUMERICALVALUE(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserNUMERICALVALUE, i)
}

func (s *RangesdatacardContext) RANGESINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserRANGESINDICATORCARD, 0)
}

func (s *RangesdatacardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesdatacardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesdatacardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterRangesdatacard(s)
	}
}

func (s *RangesdatacardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitRangesdatacard(s)
	}
}

func (p *mpsParser) Rangesdatacard() (localctx IRangesdatacardContext) {
	localctx = NewRangesdatacardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mpsParserRULE_rangesdatacard)
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
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mpsParserRANGESINDICATORCARD || _la == mpsParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(119)
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(120)
		p.Match(mpsParserNUMERICALVALUE)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(121)
			p.Match(mpsParserIDENTIFIER)
		}
		{
			p.SetState(122)
			p.Match(mpsParserNUMERICALVALUE)
		}

	}

	return localctx
}

// IBoundsdatacardContext is an interface to support dynamic dispatch.
type IBoundsdatacardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundsdatacardContext differentiates from other interfaces.
	IsBoundsdatacardContext()
}

type BoundsdatacardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundsdatacardContext() *BoundsdatacardContext {
	var p = new(BoundsdatacardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_boundsdatacard
	return p
}

func (*BoundsdatacardContext) IsBoundsdatacardContext() {}

func NewBoundsdatacardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundsdatacardContext {
	var p = new(BoundsdatacardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_boundsdatacard

	return p
}

func (s *BoundsdatacardContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundsdatacardContext) BOUNDKEY() antlr.TerminalNode {
	return s.GetToken(mpsParserBOUNDKEY, 0)
}

func (s *BoundsdatacardContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(mpsParserIDENTIFIER)
}

func (s *BoundsdatacardContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, i)
}

func (s *BoundsdatacardContext) BOUNDSINDICATORCARD() antlr.TerminalNode {
	return s.GetToken(mpsParserBOUNDSINDICATORCARD, 0)
}

func (s *BoundsdatacardContext) NUMERICALVALUE() antlr.TerminalNode {
	return s.GetToken(mpsParserNUMERICALVALUE, 0)
}

func (s *BoundsdatacardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundsdatacardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundsdatacardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterBoundsdatacard(s)
	}
}

func (s *BoundsdatacardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitBoundsdatacard(s)
	}
}

func (p *mpsParser) Boundsdatacard() (localctx IBoundsdatacardContext) {
	localctx = NewBoundsdatacardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, mpsParserRULE_boundsdatacard)
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
		p.SetState(125)
		p.Match(mpsParserBOUNDKEY)
	}
	{
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mpsParserBOUNDSINDICATORCARD || _la == mpsParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(127)
		p.Match(mpsParserIDENTIFIER)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mpsParserNUMERICALVALUE {
		{
			p.SetState(128)
			p.Match(mpsParserNUMERICALVALUE)
		}

	}

	return localctx
}

// IIntblockContext is an interface to support dynamic dispatch.
type IIntblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntblockContext differentiates from other interfaces.
	IsIntblockContext()
}

type IntblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntblockContext() *IntblockContext {
	var p = new(IntblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_intblock
	return p
}

func (*IntblockContext) IsIntblockContext() {}

func NewIntblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntblockContext {
	var p = new(IntblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_intblock

	return p
}

func (s *IntblockContext) GetParser() antlr.Parser { return s.parser }

func (s *IntblockContext) Startmarker() IStartmarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartmarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartmarkerContext)
}

func (s *IntblockContext) Endmarker() IEndmarkerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndmarkerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndmarkerContext)
}

func (s *IntblockContext) AllColumndatacard() []IColumndatacardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumndatacardContext)(nil)).Elem())
	var tst = make([]IColumndatacardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumndatacardContext)
		}
	}

	return tst
}

func (s *IntblockContext) Columndatacard(i int) IColumndatacardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumndatacardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumndatacardContext)
}

func (s *IntblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterIntblock(s)
	}
}

func (s *IntblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitIntblock(s)
	}
}

func (p *mpsParser) Intblock() (localctx IIntblockContext) {
	localctx = NewIntblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, mpsParserRULE_intblock)

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
		p.SetState(131)
		p.Startmarker()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(132)
				p.Columndatacard()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	{
		p.SetState(137)
		p.Endmarker()
	}

	return localctx
}

// IStartmarkerContext is an interface to support dynamic dispatch.
type IStartmarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartmarkerContext differentiates from other interfaces.
	IsStartmarkerContext()
}

type StartmarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartmarkerContext() *StartmarkerContext {
	var p = new(StartmarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_startmarker
	return p
}

func (*StartmarkerContext) IsStartmarkerContext() {}

func NewStartmarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartmarkerContext {
	var p = new(StartmarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_startmarker

	return p
}

func (s *StartmarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *StartmarkerContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, 0)
}

func (s *StartmarkerContext) KEYWORDMARKER() antlr.TerminalNode {
	return s.GetToken(mpsParserKEYWORDMARKER, 0)
}

func (s *StartmarkerContext) STARTMARKER() antlr.TerminalNode {
	return s.GetToken(mpsParserSTARTMARKER, 0)
}

func (s *StartmarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartmarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartmarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterStartmarker(s)
	}
}

func (s *StartmarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitStartmarker(s)
	}
}

func (p *mpsParser) Startmarker() (localctx IStartmarkerContext) {
	localctx = NewStartmarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, mpsParserRULE_startmarker)

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
		p.SetState(139)
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(140)
		p.Match(mpsParserKEYWORDMARKER)
	}
	{
		p.SetState(141)
		p.Match(mpsParserSTARTMARKER)
	}

	return localctx
}

// IEndmarkerContext is an interface to support dynamic dispatch.
type IEndmarkerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndmarkerContext differentiates from other interfaces.
	IsEndmarkerContext()
}

type EndmarkerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndmarkerContext() *EndmarkerContext {
	var p = new(EndmarkerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mpsParserRULE_endmarker
	return p
}

func (*EndmarkerContext) IsEndmarkerContext() {}

func NewEndmarkerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndmarkerContext {
	var p = new(EndmarkerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mpsParserRULE_endmarker

	return p
}

func (s *EndmarkerContext) GetParser() antlr.Parser { return s.parser }

func (s *EndmarkerContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mpsParserIDENTIFIER, 0)
}

func (s *EndmarkerContext) KEYWORDMARKER() antlr.TerminalNode {
	return s.GetToken(mpsParserKEYWORDMARKER, 0)
}

func (s *EndmarkerContext) ENDMARKER() antlr.TerminalNode {
	return s.GetToken(mpsParserENDMARKER, 0)
}

func (s *EndmarkerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndmarkerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndmarkerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.EnterEndmarker(s)
	}
}

func (s *EndmarkerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mpsListener); ok {
		listenerT.ExitEndmarker(s)
	}
}

func (p *mpsParser) Endmarker() (localctx IEndmarkerContext) {
	localctx = NewEndmarkerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, mpsParserRULE_endmarker)

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
		p.SetState(143)
		p.Match(mpsParserIDENTIFIER)
	}
	{
		p.SetState(144)
		p.Match(mpsParserKEYWORDMARKER)
	}
	{
		p.SetState(145)
		p.Match(mpsParserENDMARKER)
	}

	return localctx
}
