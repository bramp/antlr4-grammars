// Code generated from pl0.g4 by ANTLR 4.7.2. DO NOT EDIT.

package pl0 // pl0
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 179,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 5,
	3, 47, 10, 3, 3, 3, 5, 3, 50, 10, 3, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14,
	3, 56, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 69, 10, 4, 12, 4, 14, 4, 72, 11, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 7, 5, 80, 10, 5, 12, 5, 14, 5, 83, 11, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 101, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 7, 13, 123, 10, 13, 12, 13, 14, 13, 126, 11, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 146, 10, 16, 3, 17, 5,
	17, 149, 10, 17, 3, 17, 3, 17, 3, 17, 7, 17, 154, 10, 17, 12, 17, 14, 17,
	157, 11, 17, 3, 18, 3, 18, 3, 18, 7, 18, 162, 10, 18, 12, 18, 14, 18, 165,
	11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 173, 10, 19, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 5, 4, 2, 4, 4, 10, 14, 3,
	2, 15, 16, 3, 2, 17, 18, 2, 178, 2, 42, 3, 2, 2, 2, 4, 46, 3, 2, 2, 2,
	6, 59, 3, 2, 2, 2, 8, 75, 3, 2, 2, 2, 10, 86, 3, 2, 2, 2, 12, 100, 3, 2,
	2, 2, 14, 102, 3, 2, 2, 2, 16, 106, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20,
	112, 3, 2, 2, 2, 22, 115, 3, 2, 2, 2, 24, 118, 3, 2, 2, 2, 26, 129, 3,
	2, 2, 2, 28, 134, 3, 2, 2, 2, 30, 145, 3, 2, 2, 2, 32, 148, 3, 2, 2, 2,
	34, 158, 3, 2, 2, 2, 36, 172, 3, 2, 2, 2, 38, 174, 3, 2, 2, 2, 40, 176,
	3, 2, 2, 2, 42, 43, 5, 4, 3, 2, 43, 44, 7, 3, 2, 2, 44, 3, 3, 2, 2, 2,
	45, 47, 5, 6, 4, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3,
	2, 2, 2, 48, 50, 5, 8, 5, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	54, 3, 2, 2, 2, 51, 53, 5, 10, 6, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2,
	2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54,
	3, 2, 2, 2, 57, 58, 5, 12, 7, 2, 58, 5, 3, 2, 2, 2, 59, 60, 7, 30, 2, 2,
	60, 61, 5, 38, 20, 2, 61, 62, 7, 4, 2, 2, 62, 70, 5, 40, 21, 2, 63, 64,
	7, 5, 2, 2, 64, 65, 5, 38, 20, 2, 65, 66, 7, 4, 2, 2, 66, 67, 5, 40, 21,
	2, 67, 69, 3, 2, 2, 2, 68, 63, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2,
	73, 74, 7, 6, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7, 31, 2, 2, 76, 81, 5,
	38, 20, 2, 77, 78, 7, 5, 2, 2, 78, 80, 5, 38, 20, 2, 79, 77, 3, 2, 2, 2,
	80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 6, 2, 2, 85, 9, 3, 2, 2, 2, 86,
	87, 7, 32, 2, 2, 87, 88, 5, 38, 20, 2, 88, 89, 7, 6, 2, 2, 89, 90, 5, 4,
	3, 2, 90, 91, 7, 6, 2, 2, 91, 11, 3, 2, 2, 2, 92, 101, 5, 14, 8, 2, 93,
	101, 5, 16, 9, 2, 94, 101, 5, 18, 10, 2, 95, 101, 5, 20, 11, 2, 96, 101,
	5, 22, 12, 2, 97, 101, 5, 24, 13, 2, 98, 101, 5, 26, 14, 2, 99, 101, 5,
	28, 15, 2, 100, 92, 3, 2, 2, 2, 100, 93, 3, 2, 2, 2, 100, 94, 3, 2, 2,
	2, 100, 95, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 100, 97, 3, 2, 2, 2, 100,
	98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 13, 3,
	2, 2, 2, 102, 103, 5, 38, 20, 2, 103, 104, 7, 7, 2, 2, 104, 105, 5, 32,
	17, 2, 105, 15, 3, 2, 2, 2, 106, 107, 7, 29, 2, 2, 107, 108, 5, 38, 20,
	2, 108, 17, 3, 2, 2, 2, 109, 110, 7, 21, 2, 2, 110, 111, 5, 38, 20, 2,
	111, 19, 3, 2, 2, 2, 112, 113, 7, 8, 2, 2, 113, 114, 5, 38, 20, 2, 114,
	21, 3, 2, 2, 2, 115, 116, 7, 9, 2, 2, 116, 117, 5, 32, 17, 2, 117, 23,
	3, 2, 2, 2, 118, 119, 7, 27, 2, 2, 119, 124, 5, 12, 7, 2, 120, 121, 7,
	6, 2, 2, 121, 123, 5, 12, 7, 2, 122, 120, 3, 2, 2, 2, 123, 126, 3, 2, 2,
	2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 128, 7, 28, 2, 2, 128, 25, 3, 2, 2, 2, 129, 130,
	7, 24, 2, 2, 130, 131, 5, 30, 16, 2, 131, 132, 7, 25, 2, 2, 132, 133, 5,
	12, 7, 2, 133, 27, 3, 2, 2, 2, 134, 135, 7, 22, 2, 2, 135, 136, 5, 30,
	16, 2, 136, 137, 7, 23, 2, 2, 137, 138, 5, 12, 7, 2, 138, 29, 3, 2, 2,
	2, 139, 140, 7, 26, 2, 2, 140, 146, 5, 32, 17, 2, 141, 142, 5, 32, 17,
	2, 142, 143, 9, 2, 2, 2, 143, 144, 5, 32, 17, 2, 144, 146, 3, 2, 2, 2,
	145, 139, 3, 2, 2, 2, 145, 141, 3, 2, 2, 2, 146, 31, 3, 2, 2, 2, 147, 149,
	9, 3, 2, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 3, 2,
	2, 2, 150, 155, 5, 34, 18, 2, 151, 152, 9, 3, 2, 2, 152, 154, 5, 34, 18,
	2, 153, 151, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 33, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 163, 5,
	36, 19, 2, 159, 160, 9, 4, 2, 2, 160, 162, 5, 36, 19, 2, 161, 159, 3, 2,
	2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2,
	164, 35, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 173, 5, 38, 20, 2, 167,
	173, 5, 40, 21, 2, 168, 169, 7, 19, 2, 2, 169, 170, 5, 32, 17, 2, 170,
	171, 7, 20, 2, 2, 171, 173, 3, 2, 2, 2, 172, 166, 3, 2, 2, 2, 172, 167,
	3, 2, 2, 2, 172, 168, 3, 2, 2, 2, 173, 37, 3, 2, 2, 2, 174, 175, 7, 33,
	2, 2, 175, 39, 3, 2, 2, 2, 176, 177, 7, 34, 2, 2, 177, 41, 3, 2, 2, 2,
	14, 46, 49, 54, 70, 81, 100, 124, 145, 148, 155, 163, 172,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'='", "','", "';'", "':='", "'?'", "'!'", "'#'", "'<'", "'<='",
	"'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "WRITE", "WHILE", "DO", "IF", "THEN", "ODD", "BEGIN", "END", "CALL",
	"CONST", "VAR", "PROCEDURE", "STRING", "NUMBER", "WS",
}

var ruleNames = []string{
	"program", "block", "consts", "vars", "procedure", "statement", "assignstmt",
	"callstmt", "writestmt", "qstmt", "bangstmt", "beginstmt", "ifstmt", "whilestmt",
	"condition", "expression", "term", "factor", "ident", "number",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type pl0Parser struct {
	*antlr.BaseParser
}

func Newpl0Parser(input antlr.TokenStream) *pl0Parser {
	this := new(pl0Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "pl0.g4"

	return this
}

// pl0Parser tokens.
const (
	pl0ParserEOF       = antlr.TokenEOF
	pl0ParserT__0      = 1
	pl0ParserT__1      = 2
	pl0ParserT__2      = 3
	pl0ParserT__3      = 4
	pl0ParserT__4      = 5
	pl0ParserT__5      = 6
	pl0ParserT__6      = 7
	pl0ParserT__7      = 8
	pl0ParserT__8      = 9
	pl0ParserT__9      = 10
	pl0ParserT__10     = 11
	pl0ParserT__11     = 12
	pl0ParserT__12     = 13
	pl0ParserT__13     = 14
	pl0ParserT__14     = 15
	pl0ParserT__15     = 16
	pl0ParserT__16     = 17
	pl0ParserT__17     = 18
	pl0ParserWRITE     = 19
	pl0ParserWHILE     = 20
	pl0ParserDO        = 21
	pl0ParserIF        = 22
	pl0ParserTHEN      = 23
	pl0ParserODD       = 24
	pl0ParserBEGIN     = 25
	pl0ParserEND       = 26
	pl0ParserCALL      = 27
	pl0ParserCONST     = 28
	pl0ParserVAR       = 29
	pl0ParserPROCEDURE = 30
	pl0ParserSTRING    = 31
	pl0ParserNUMBER    = 32
	pl0ParserWS        = 33
)

// pl0Parser rules.
const (
	pl0ParserRULE_program    = 0
	pl0ParserRULE_block      = 1
	pl0ParserRULE_consts     = 2
	pl0ParserRULE_vars       = 3
	pl0ParserRULE_procedure  = 4
	pl0ParserRULE_statement  = 5
	pl0ParserRULE_assignstmt = 6
	pl0ParserRULE_callstmt   = 7
	pl0ParserRULE_writestmt  = 8
	pl0ParserRULE_qstmt      = 9
	pl0ParserRULE_bangstmt   = 10
	pl0ParserRULE_beginstmt  = 11
	pl0ParserRULE_ifstmt     = 12
	pl0ParserRULE_whilestmt  = 13
	pl0ParserRULE_condition  = 14
	pl0ParserRULE_expression = 15
	pl0ParserRULE_term       = 16
	pl0ParserRULE_factor     = 17
	pl0ParserRULE_ident      = 18
	pl0ParserRULE_number     = 19
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
	p.RuleIndex = pl0ParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *pl0Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pl0ParserRULE_program)

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
		p.Block()
	}
	{
		p.SetState(41)
		p.Match(pl0ParserT__0)
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
	p.RuleIndex = pl0ParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) Consts() IConstsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstsContext)
}

func (s *BlockContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *BlockContext) AllProcedure() []IProcedureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcedureContext)(nil)).Elem())
	var tst = make([]IProcedureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcedureContext)
		}
	}

	return tst
}

func (s *BlockContext) Procedure(i int) IProcedureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcedureContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *pl0Parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, pl0ParserRULE_block)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pl0ParserCONST {
		{
			p.SetState(43)
			p.Consts()
		}

	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pl0ParserVAR {
		{
			p.SetState(46)
			p.Vars()
		}

	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserPROCEDURE {
		{
			p.SetState(49)
			p.Procedure()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Statement()
	}

	return localctx
}

// IConstsContext is an interface to support dynamic dispatch.
type IConstsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstsContext differentiates from other interfaces.
	IsConstsContext()
}

type ConstsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstsContext() *ConstsContext {
	var p = new(ConstsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_consts
	return p
}

func (*ConstsContext) IsConstsContext() {}

func NewConstsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstsContext {
	var p = new(ConstsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_consts

	return p
}

func (s *ConstsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstsContext) CONST() antlr.TerminalNode {
	return s.GetToken(pl0ParserCONST, 0)
}

func (s *ConstsContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *ConstsContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ConstsContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *ConstsContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ConstsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterConsts(s)
	}
}

func (s *ConstsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitConsts(s)
	}
}

func (p *pl0Parser) Consts() (localctx IConstsContext) {
	localctx = NewConstsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pl0ParserRULE_consts)
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
		p.SetState(57)
		p.Match(pl0ParserCONST)
	}
	{
		p.SetState(58)
		p.Ident()
	}
	{
		p.SetState(59)
		p.Match(pl0ParserT__1)
	}
	{
		p.SetState(60)
		p.Number()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserT__2 {
		{
			p.SetState(61)
			p.Match(pl0ParserT__2)
		}
		{
			p.SetState(62)
			p.Ident()
		}
		{
			p.SetState(63)
			p.Match(pl0ParserT__1)
		}
		{
			p.SetState(64)
			p.Number()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(pl0ParserT__3)
	}

	return localctx
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) VAR() antlr.TerminalNode {
	return s.GetToken(pl0ParserVAR, 0)
}

func (s *VarsContext) AllIdent() []IIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentContext)(nil)).Elem())
	var tst = make([]IIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentContext)
		}
	}

	return tst
}

func (s *VarsContext) Ident(i int) IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *pl0Parser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, pl0ParserRULE_vars)
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
		p.SetState(73)
		p.Match(pl0ParserVAR)
	}
	{
		p.SetState(74)
		p.Ident()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserT__2 {
		{
			p.SetState(75)
			p.Match(pl0ParserT__2)
		}
		{
			p.SetState(76)
			p.Ident()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(pl0ParserT__3)
	}

	return localctx
}

// IProcedureContext is an interface to support dynamic dispatch.
type IProcedureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedureContext differentiates from other interfaces.
	IsProcedureContext()
}

type ProcedureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureContext() *ProcedureContext {
	var p = new(ProcedureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_procedure
	return p
}

func (*ProcedureContext) IsProcedureContext() {}

func NewProcedureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureContext {
	var p = new(ProcedureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_procedure

	return p
}

func (s *ProcedureContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(pl0ParserPROCEDURE, 0)
}

func (s *ProcedureContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ProcedureContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProcedureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterProcedure(s)
	}
}

func (s *ProcedureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitProcedure(s)
	}
}

func (p *pl0Parser) Procedure() (localctx IProcedureContext) {
	localctx = NewProcedureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, pl0ParserRULE_procedure)

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
		p.Match(pl0ParserPROCEDURE)
	}
	{
		p.SetState(85)
		p.Ident()
	}
	{
		p.SetState(86)
		p.Match(pl0ParserT__3)
	}
	{
		p.SetState(87)
		p.Block()
	}
	{
		p.SetState(88)
		p.Match(pl0ParserT__3)
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
	p.RuleIndex = pl0ParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignstmt() IAssignstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignstmtContext)
}

func (s *StatementContext) Callstmt() ICallstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallstmtContext)
}

func (s *StatementContext) Writestmt() IWritestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWritestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWritestmtContext)
}

func (s *StatementContext) Qstmt() IQstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQstmtContext)
}

func (s *StatementContext) Bangstmt() IBangstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBangstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBangstmtContext)
}

func (s *StatementContext) Beginstmt() IBeginstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginstmtContext)
}

func (s *StatementContext) Ifstmt() IIfstmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StatementContext) Whilestmt() IWhilestmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhilestmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *pl0Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, pl0ParserRULE_statement)

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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pl0ParserSTRING:
		{
			p.SetState(90)
			p.Assignstmt()
		}

	case pl0ParserCALL:
		{
			p.SetState(91)
			p.Callstmt()
		}

	case pl0ParserWRITE:
		{
			p.SetState(92)
			p.Writestmt()
		}

	case pl0ParserT__5:
		{
			p.SetState(93)
			p.Qstmt()
		}

	case pl0ParserT__6:
		{
			p.SetState(94)
			p.Bangstmt()
		}

	case pl0ParserBEGIN:
		{
			p.SetState(95)
			p.Beginstmt()
		}

	case pl0ParserIF:
		{
			p.SetState(96)
			p.Ifstmt()
		}

	case pl0ParserWHILE:
		{
			p.SetState(97)
			p.Whilestmt()
		}

	case pl0ParserT__0, pl0ParserT__3, pl0ParserEND:

	default:
	}

	return localctx
}

// IAssignstmtContext is an interface to support dynamic dispatch.
type IAssignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignstmtContext differentiates from other interfaces.
	IsAssignstmtContext()
}

type AssignstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstmtContext() *AssignstmtContext {
	var p = new(AssignstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_assignstmt
	return p
}

func (*AssignstmtContext) IsAssignstmtContext() {}

func NewAssignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstmtContext {
	var p = new(AssignstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_assignstmt

	return p
}

func (s *AssignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *AssignstmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterAssignstmt(s)
	}
}

func (s *AssignstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitAssignstmt(s)
	}
}

func (p *pl0Parser) Assignstmt() (localctx IAssignstmtContext) {
	localctx = NewAssignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, pl0ParserRULE_assignstmt)

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
		p.SetState(100)
		p.Ident()
	}
	{
		p.SetState(101)
		p.Match(pl0ParserT__4)
	}
	{
		p.SetState(102)
		p.Expression()
	}

	return localctx
}

// ICallstmtContext is an interface to support dynamic dispatch.
type ICallstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallstmtContext differentiates from other interfaces.
	IsCallstmtContext()
}

type CallstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallstmtContext() *CallstmtContext {
	var p = new(CallstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_callstmt
	return p
}

func (*CallstmtContext) IsCallstmtContext() {}

func NewCallstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallstmtContext {
	var p = new(CallstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_callstmt

	return p
}

func (s *CallstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CallstmtContext) CALL() antlr.TerminalNode {
	return s.GetToken(pl0ParserCALL, 0)
}

func (s *CallstmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *CallstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterCallstmt(s)
	}
}

func (s *CallstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitCallstmt(s)
	}
}

func (p *pl0Parser) Callstmt() (localctx ICallstmtContext) {
	localctx = NewCallstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, pl0ParserRULE_callstmt)

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
		p.Match(pl0ParserCALL)
	}
	{
		p.SetState(105)
		p.Ident()
	}

	return localctx
}

// IWritestmtContext is an interface to support dynamic dispatch.
type IWritestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWritestmtContext differentiates from other interfaces.
	IsWritestmtContext()
}

type WritestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWritestmtContext() *WritestmtContext {
	var p = new(WritestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_writestmt
	return p
}

func (*WritestmtContext) IsWritestmtContext() {}

func NewWritestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WritestmtContext {
	var p = new(WritestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_writestmt

	return p
}

func (s *WritestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WritestmtContext) WRITE() antlr.TerminalNode {
	return s.GetToken(pl0ParserWRITE, 0)
}

func (s *WritestmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *WritestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WritestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WritestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterWritestmt(s)
	}
}

func (s *WritestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitWritestmt(s)
	}
}

func (p *pl0Parser) Writestmt() (localctx IWritestmtContext) {
	localctx = NewWritestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, pl0ParserRULE_writestmt)

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
		p.SetState(107)
		p.Match(pl0ParserWRITE)
	}
	{
		p.SetState(108)
		p.Ident()
	}

	return localctx
}

// IQstmtContext is an interface to support dynamic dispatch.
type IQstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQstmtContext differentiates from other interfaces.
	IsQstmtContext()
}

type QstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQstmtContext() *QstmtContext {
	var p = new(QstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_qstmt
	return p
}

func (*QstmtContext) IsQstmtContext() {}

func NewQstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QstmtContext {
	var p = new(QstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_qstmt

	return p
}

func (s *QstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *QstmtContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *QstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterQstmt(s)
	}
}

func (s *QstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitQstmt(s)
	}
}

func (p *pl0Parser) Qstmt() (localctx IQstmtContext) {
	localctx = NewQstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, pl0ParserRULE_qstmt)

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
		p.Match(pl0ParserT__5)
	}
	{
		p.SetState(111)
		p.Ident()
	}

	return localctx
}

// IBangstmtContext is an interface to support dynamic dispatch.
type IBangstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBangstmtContext differentiates from other interfaces.
	IsBangstmtContext()
}

type BangstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBangstmtContext() *BangstmtContext {
	var p = new(BangstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_bangstmt
	return p
}

func (*BangstmtContext) IsBangstmtContext() {}

func NewBangstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BangstmtContext {
	var p = new(BangstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_bangstmt

	return p
}

func (s *BangstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BangstmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BangstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BangstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BangstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterBangstmt(s)
	}
}

func (s *BangstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitBangstmt(s)
	}
}

func (p *pl0Parser) Bangstmt() (localctx IBangstmtContext) {
	localctx = NewBangstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, pl0ParserRULE_bangstmt)

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
		p.Match(pl0ParserT__6)
	}
	{
		p.SetState(114)
		p.Expression()
	}

	return localctx
}

// IBeginstmtContext is an interface to support dynamic dispatch.
type IBeginstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeginstmtContext differentiates from other interfaces.
	IsBeginstmtContext()
}

type BeginstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeginstmtContext() *BeginstmtContext {
	var p = new(BeginstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_beginstmt
	return p
}

func (*BeginstmtContext) IsBeginstmtContext() {}

func NewBeginstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginstmtContext {
	var p = new(BeginstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_beginstmt

	return p
}

func (s *BeginstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginstmtContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(pl0ParserBEGIN, 0)
}

func (s *BeginstmtContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BeginstmtContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BeginstmtContext) END() antlr.TerminalNode {
	return s.GetToken(pl0ParserEND, 0)
}

func (s *BeginstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterBeginstmt(s)
	}
}

func (s *BeginstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitBeginstmt(s)
	}
}

func (p *pl0Parser) Beginstmt() (localctx IBeginstmtContext) {
	localctx = NewBeginstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, pl0ParserRULE_beginstmt)
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
		p.SetState(116)
		p.Match(pl0ParserBEGIN)
	}
	{
		p.SetState(117)
		p.Statement()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserT__3 {
		{
			p.SetState(118)
			p.Match(pl0ParserT__3)
		}
		{
			p.SetState(119)
			p.Statement()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(pl0ParserEND)
	}

	return localctx
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_ifstmt
	return p
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(pl0ParserIF, 0)
}

func (s *IfstmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfstmtContext) THEN() antlr.TerminalNode {
	return s.GetToken(pl0ParserTHEN, 0)
}

func (s *IfstmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *pl0Parser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, pl0ParserRULE_ifstmt)

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
		p.Match(pl0ParserIF)
	}
	{
		p.SetState(128)
		p.Condition()
	}
	{
		p.SetState(129)
		p.Match(pl0ParserTHEN)
	}
	{
		p.SetState(130)
		p.Statement()
	}

	return localctx
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_whilestmt
	return p
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(pl0ParserWHILE, 0)
}

func (s *WhilestmtContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhilestmtContext) DO() antlr.TerminalNode {
	return s.GetToken(pl0ParserDO, 0)
}

func (s *WhilestmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *pl0Parser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, pl0ParserRULE_whilestmt)

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
		p.SetState(132)
		p.Match(pl0ParserWHILE)
	}
	{
		p.SetState(133)
		p.Condition()
	}
	{
		p.SetState(134)
		p.Match(pl0ParserDO)
	}
	{
		p.SetState(135)
		p.Statement()
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) ODD() antlr.TerminalNode {
	return s.GetToken(pl0ParserODD, 0)
}

func (s *ConditionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *pl0Parser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, pl0ParserRULE_condition)
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pl0ParserODD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(pl0ParserODD)
		}
		{
			p.SetState(138)
			p.Expression()
		}

	case pl0ParserT__12, pl0ParserT__13, pl0ParserT__16, pl0ParserSTRING, pl0ParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Expression()
		}
		{
			p.SetState(140)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pl0ParserT__1)|(1<<pl0ParserT__7)|(1<<pl0ParserT__8)|(1<<pl0ParserT__9)|(1<<pl0ParserT__10)|(1<<pl0ParserT__11))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(141)
			p.Expression()
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
	p.RuleIndex = pl0ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_expression

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
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *pl0Parser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, pl0ParserRULE_expression)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pl0ParserT__12 || _la == pl0ParserT__13 {
		{
			p.SetState(145)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pl0ParserT__12 || _la == pl0ParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(148)
		p.Term()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserT__12 || _la == pl0ParserT__13 {
		{
			p.SetState(149)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pl0ParserT__12 || _la == pl0ParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)
			p.Term()
		}

		p.SetState(155)
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
	p.RuleIndex = pl0ParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *pl0Parser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, pl0ParserRULE_term)
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
		p.SetState(156)
		p.Factor()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pl0ParserT__14 || _la == pl0ParserT__15 {
		{
			p.SetState(157)
			_la = p.GetTokenStream().LA(1)

			if !(_la == pl0ParserT__14 || _la == pl0ParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(158)
			p.Factor()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FactorContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *pl0Parser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, pl0ParserRULE_factor)

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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pl0ParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Ident()
		}

	case pl0ParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Number()
		}

	case pl0ParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(pl0ParserT__16)
		}
		{
			p.SetState(167)
			p.Expression()
		}
		{
			p.SetState(168)
			p.Match(pl0ParserT__17)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pl0ParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) STRING() antlr.TerminalNode {
	return s.GetToken(pl0ParserSTRING, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitIdent(s)
	}
}

func (p *pl0Parser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, pl0ParserRULE_ident)

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
		p.SetState(172)
		p.Match(pl0ParserSTRING)
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
	p.RuleIndex = pl0ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pl0ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(pl0ParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pl0Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *pl0Parser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, pl0ParserRULE_number)

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
		p.SetState(174)
		p.Match(pl0ParserNUMBER)
	}

	return localctx
}
