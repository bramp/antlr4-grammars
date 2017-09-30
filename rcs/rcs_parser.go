// Generated from RCS.g4 by ANTLR 4.7.

package rcs // RCS
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 221,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 68,
	10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 5, 5, 5, 76, 10, 5, 3, 5,
	5, 5, 79, 10, 5, 3, 5, 5, 5, 82, 10, 5, 3, 5, 7, 5, 85, 10, 5, 12, 5, 14,
	5, 88, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 7, 8, 100, 10, 8, 12, 8, 14, 8, 103, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	7, 9, 109, 10, 9, 12, 9, 14, 9, 112, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 7, 11, 122, 10, 11, 12, 11, 14, 11, 125, 11, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 134, 10, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 5, 14, 140, 10, 14, 3, 14, 3, 14, 3, 15, 7, 15,
	145, 10, 15, 12, 15, 14, 15, 148, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 7, 16, 157, 10, 16, 12, 16, 14, 16, 160, 11, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 7, 20, 176, 10, 20, 12, 20, 14, 20, 179, 11, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 5, 21, 185, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 23, 7, 23, 193, 10, 23, 12, 23, 14, 23, 196, 11, 23, 3, 24, 3, 24,
	3, 24, 7, 24, 201, 10, 24, 12, 24, 14, 24, 204, 11, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 6, 27, 215, 10, 27, 13, 27,
	14, 27, 216, 3, 27, 3, 27, 3, 27, 2, 2, 28, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	2, 2, 2, 212, 2, 54, 3, 2, 2, 2, 4, 60, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2,
	8, 65, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12, 93, 3, 2, 2, 2, 14, 97, 3, 2,
	2, 2, 16, 106, 3, 2, 2, 2, 18, 115, 3, 2, 2, 2, 20, 119, 3, 2, 2, 2, 22,
	128, 3, 2, 2, 2, 24, 131, 3, 2, 2, 2, 26, 137, 3, 2, 2, 2, 28, 146, 3,
	2, 2, 2, 30, 149, 3, 2, 2, 2, 32, 161, 3, 2, 2, 2, 34, 165, 3, 2, 2, 2,
	36, 169, 3, 2, 2, 2, 38, 173, 3, 2, 2, 2, 40, 182, 3, 2, 2, 2, 42, 188,
	3, 2, 2, 2, 44, 194, 3, 2, 2, 2, 46, 197, 3, 2, 2, 2, 48, 207, 3, 2, 2,
	2, 50, 210, 3, 2, 2, 2, 52, 214, 3, 2, 2, 2, 54, 55, 5, 8, 5, 2, 55, 56,
	5, 28, 15, 2, 56, 57, 5, 42, 22, 2, 57, 58, 5, 44, 23, 2, 58, 59, 7, 2,
	2, 3, 59, 3, 3, 2, 2, 2, 60, 61, 5, 8, 5, 2, 61, 5, 3, 2, 2, 2, 62, 63,
	5, 8, 5, 2, 63, 64, 5, 28, 15, 2, 64, 7, 3, 2, 2, 2, 65, 67, 5, 10, 6,
	2, 66, 68, 5, 12, 7, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69,
	3, 2, 2, 2, 69, 70, 5, 14, 8, 2, 70, 72, 5, 16, 9, 2, 71, 73, 5, 20, 11,
	2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 76,
	5, 22, 12, 2, 75, 74, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2,
	2, 77, 79, 5, 24, 13, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81,
	3, 2, 2, 2, 80, 82, 5, 26, 14, 2, 81, 80, 3, 2, 2, 2, 81, 82, 3, 2, 2,
	2, 82, 86, 3, 2, 2, 2, 83, 85, 5, 52, 27, 2, 84, 83, 3, 2, 2, 2, 85, 88,
	3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 9, 3, 2, 2, 2,
	88, 86, 3, 2, 2, 2, 89, 90, 7, 8, 2, 2, 90, 91, 7, 26, 2, 2, 91, 92, 7,
	28, 2, 2, 92, 11, 3, 2, 2, 2, 93, 94, 7, 9, 2, 2, 94, 95, 7, 4, 2, 2, 95,
	96, 7, 28, 2, 2, 96, 13, 3, 2, 2, 2, 97, 101, 7, 10, 2, 2, 98, 100, 7,
	24, 2, 2, 99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2,
	101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104,
	105, 7, 28, 2, 2, 105, 15, 3, 2, 2, 2, 106, 110, 7, 11, 2, 2, 107, 109,
	5, 18, 10, 2, 108, 107, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3,
	2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2,
	2, 113, 114, 7, 28, 2, 2, 114, 17, 3, 2, 2, 2, 115, 116, 7, 24, 2, 2, 116,
	117, 7, 29, 2, 2, 117, 118, 7, 26, 2, 2, 118, 19, 3, 2, 2, 2, 119, 123,
	7, 12, 2, 2, 120, 122, 7, 24, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3,
	2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2,
	2, 125, 123, 3, 2, 2, 2, 126, 127, 7, 28, 2, 2, 127, 21, 3, 2, 2, 2, 128,
	129, 7, 13, 2, 2, 129, 130, 7, 28, 2, 2, 130, 23, 3, 2, 2, 2, 131, 133,
	7, 14, 2, 2, 132, 134, 7, 27, 2, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3,
	2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 28, 2, 2, 136, 25, 3, 2, 2,
	2, 137, 139, 7, 15, 2, 2, 138, 140, 7, 27, 2, 2, 139, 138, 3, 2, 2, 2,
	139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 7, 28, 2, 2, 142,
	27, 3, 2, 2, 2, 143, 145, 5, 30, 16, 2, 144, 143, 3, 2, 2, 2, 145, 148,
	3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 29, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 26, 2, 2, 150, 151, 5, 32, 17,
	2, 151, 152, 5, 34, 18, 2, 152, 153, 5, 36, 19, 2, 153, 154, 5, 38, 20,
	2, 154, 158, 5, 40, 21, 2, 155, 157, 5, 52, 27, 2, 156, 155, 3, 2, 2, 2,
	157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159,
	31, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 7, 16, 2, 2, 162, 163,
	7, 26, 2, 2, 163, 164, 7, 28, 2, 2, 164, 33, 3, 2, 2, 2, 165, 166, 7, 17,
	2, 2, 166, 167, 7, 24, 2, 2, 167, 168, 7, 28, 2, 2, 168, 35, 3, 2, 2, 2,
	169, 170, 7, 18, 2, 2, 170, 171, 7, 24, 2, 2, 171, 172, 7, 28, 2, 2, 172,
	37, 3, 2, 2, 2, 173, 177, 7, 19, 2, 2, 174, 176, 7, 26, 2, 2, 175, 174,
	3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2,
	2, 2, 178, 180, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 181, 7, 28, 2, 2,
	181, 39, 3, 2, 2, 2, 182, 184, 7, 20, 2, 2, 183, 185, 7, 26, 2, 2, 184,
	183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187,
	7, 28, 2, 2, 187, 41, 3, 2, 2, 2, 188, 189, 7, 21, 2, 2, 189, 190, 7, 27,
	2, 2, 190, 43, 3, 2, 2, 2, 191, 193, 5, 46, 24, 2, 192, 191, 3, 2, 2, 2,
	193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	45, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 198, 7, 26, 2, 2, 198, 202,
	5, 48, 25, 2, 199, 201, 5, 52, 27, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3,
	2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 3, 2, 2,
	2, 204, 202, 3, 2, 2, 2, 205, 206, 5, 50, 26, 2, 206, 47, 3, 2, 2, 2, 207,
	208, 7, 22, 2, 2, 208, 209, 7, 27, 2, 2, 209, 49, 3, 2, 2, 2, 210, 211,
	7, 23, 2, 2, 211, 212, 7, 27, 2, 2, 212, 51, 3, 2, 2, 2, 213, 215, 7, 24,
	2, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2,
	216, 217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 7, 28, 2, 2, 219,
	53, 3, 2, 2, 2, 20, 67, 72, 75, 78, 81, 86, 101, 110, 123, 133, 139, 146,
	158, 177, 184, 194, 202, 216,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'COMMA'", "'BRANCH'", "'LOGS'", "'ADMIN'", "'DELTAS'", "'head'", "'branch'",
	"'access'", "'symbols'", "'locks'", "'strict'", "'comment'", "'expand'",
	"'date'", "'author'", "'state'", "'branches'", "'next'", "'desc'", "'log'",
	"'text'", "", "", "", "", "';'", "':'",
}
var symbolicNames = []string{
	"", "COMMA", "BRANCH", "LOGS", "ADMIN", "DELTAS", "LITERAL_HEAD", "LITERAL_BRANCH",
	"LITERAL_ACCESS", "LITERAL_SYMBOLS", "LITERAL_LOCKS", "LITERAL_STRICT",
	"LITERAL_COMMENT", "LITERAL_EXPAND", "LITERAL_DATE", "LITERAL_AUTHOR",
	"LITERAL_STATE", "LITERAL_BRANCHES", "LITERAL_NEXT", "LITERAL_DESC", "LITERAL_LOG",
	"LITERAL_TEXT", "IDENT", "INT", "REVISION", "STRING", "SEMI", "COLON",
	"WS",
}

var ruleNames = []string{
	"rcstext", "rcsheader", "rcsrevisions", "admin", "head", "branch", "access",
	"symbols", "tags", "locks", "strict", "comment", "expand", "deltalist",
	"delta", "delta_date", "delta_author", "delta_state", "delta_branches",
	"delta_next", "desc", "deltatextlist", "deltatext", "deltatext_log", "deltatext_text",
	"newphrase",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RCSParser struct {
	*antlr.BaseParser
}

func NewRCSParser(input antlr.TokenStream) *RCSParser {
	this := new(RCSParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RCS.g4"

	return this
}

// RCSParser tokens.
const (
	RCSParserEOF              = antlr.TokenEOF
	RCSParserCOMMA            = 1
	RCSParserBRANCH           = 2
	RCSParserLOGS             = 3
	RCSParserADMIN            = 4
	RCSParserDELTAS           = 5
	RCSParserLITERAL_HEAD     = 6
	RCSParserLITERAL_BRANCH   = 7
	RCSParserLITERAL_ACCESS   = 8
	RCSParserLITERAL_SYMBOLS  = 9
	RCSParserLITERAL_LOCKS    = 10
	RCSParserLITERAL_STRICT   = 11
	RCSParserLITERAL_COMMENT  = 12
	RCSParserLITERAL_EXPAND   = 13
	RCSParserLITERAL_DATE     = 14
	RCSParserLITERAL_AUTHOR   = 15
	RCSParserLITERAL_STATE    = 16
	RCSParserLITERAL_BRANCHES = 17
	RCSParserLITERAL_NEXT     = 18
	RCSParserLITERAL_DESC     = 19
	RCSParserLITERAL_LOG      = 20
	RCSParserLITERAL_TEXT     = 21
	RCSParserIDENT            = 22
	RCSParserINT              = 23
	RCSParserREVISION         = 24
	RCSParserSTRING           = 25
	RCSParserSEMI             = 26
	RCSParserCOLON            = 27
	RCSParserWS               = 28
)

// RCSParser rules.
const (
	RCSParserRULE_rcstext        = 0
	RCSParserRULE_rcsheader      = 1
	RCSParserRULE_rcsrevisions   = 2
	RCSParserRULE_admin          = 3
	RCSParserRULE_head           = 4
	RCSParserRULE_branch         = 5
	RCSParserRULE_access         = 6
	RCSParserRULE_symbols        = 7
	RCSParserRULE_tags           = 8
	RCSParserRULE_locks          = 9
	RCSParserRULE_strict         = 10
	RCSParserRULE_comment        = 11
	RCSParserRULE_expand         = 12
	RCSParserRULE_deltalist      = 13
	RCSParserRULE_delta          = 14
	RCSParserRULE_delta_date     = 15
	RCSParserRULE_delta_author   = 16
	RCSParserRULE_delta_state    = 17
	RCSParserRULE_delta_branches = 18
	RCSParserRULE_delta_next     = 19
	RCSParserRULE_desc           = 20
	RCSParserRULE_deltatextlist  = 21
	RCSParserRULE_deltatext      = 22
	RCSParserRULE_deltatext_log  = 23
	RCSParserRULE_deltatext_text = 24
	RCSParserRULE_newphrase      = 25
)

// IRcstextContext is an interface to support dynamic dispatch.
type IRcstextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRcstextContext differentiates from other interfaces.
	IsRcstextContext()
}

type RcstextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRcstextContext() *RcstextContext {
	var p = new(RcstextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_rcstext
	return p
}

func (*RcstextContext) IsRcstextContext() {}

func NewRcstextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RcstextContext {
	var p = new(RcstextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_rcstext

	return p
}

func (s *RcstextContext) GetParser() antlr.Parser { return s.parser }

func (s *RcstextContext) Admin() IAdminContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminContext)
}

func (s *RcstextContext) Deltalist() IDeltalistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltalistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeltalistContext)
}

func (s *RcstextContext) Desc() IDescContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescContext)
}

func (s *RcstextContext) Deltatextlist() IDeltatextlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltatextlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeltatextlistContext)
}

func (s *RcstextContext) EOF() antlr.TerminalNode {
	return s.GetToken(RCSParserEOF, 0)
}

func (s *RcstextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RcstextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RcstextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterRcstext(s)
	}
}

func (s *RcstextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitRcstext(s)
	}
}

func (s *RcstextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitRcstext(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Rcstext() (localctx IRcstextContext) {
	localctx = NewRcstextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RCSParserRULE_rcstext)

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
		p.SetState(52)
		p.Admin()
	}
	{
		p.SetState(53)
		p.Deltalist()
	}
	{
		p.SetState(54)
		p.Desc()
	}
	{
		p.SetState(55)
		p.Deltatextlist()
	}
	{
		p.SetState(56)
		p.Match(RCSParserEOF)
	}

	return localctx
}

// IRcsheaderContext is an interface to support dynamic dispatch.
type IRcsheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRcsheaderContext differentiates from other interfaces.
	IsRcsheaderContext()
}

type RcsheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRcsheaderContext() *RcsheaderContext {
	var p = new(RcsheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_rcsheader
	return p
}

func (*RcsheaderContext) IsRcsheaderContext() {}

func NewRcsheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RcsheaderContext {
	var p = new(RcsheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_rcsheader

	return p
}

func (s *RcsheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *RcsheaderContext) Admin() IAdminContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminContext)
}

func (s *RcsheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RcsheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RcsheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterRcsheader(s)
	}
}

func (s *RcsheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitRcsheader(s)
	}
}

func (s *RcsheaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitRcsheader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Rcsheader() (localctx IRcsheaderContext) {
	localctx = NewRcsheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RCSParserRULE_rcsheader)

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
		p.SetState(58)
		p.Admin()
	}

	return localctx
}

// IRcsrevisionsContext is an interface to support dynamic dispatch.
type IRcsrevisionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRcsrevisionsContext differentiates from other interfaces.
	IsRcsrevisionsContext()
}

type RcsrevisionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRcsrevisionsContext() *RcsrevisionsContext {
	var p = new(RcsrevisionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_rcsrevisions
	return p
}

func (*RcsrevisionsContext) IsRcsrevisionsContext() {}

func NewRcsrevisionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RcsrevisionsContext {
	var p = new(RcsrevisionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_rcsrevisions

	return p
}

func (s *RcsrevisionsContext) GetParser() antlr.Parser { return s.parser }

func (s *RcsrevisionsContext) Admin() IAdminContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminContext)
}

func (s *RcsrevisionsContext) Deltalist() IDeltalistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltalistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeltalistContext)
}

func (s *RcsrevisionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RcsrevisionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RcsrevisionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterRcsrevisions(s)
	}
}

func (s *RcsrevisionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitRcsrevisions(s)
	}
}

func (s *RcsrevisionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitRcsrevisions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Rcsrevisions() (localctx IRcsrevisionsContext) {
	localctx = NewRcsrevisionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RCSParserRULE_rcsrevisions)

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
		p.Admin()
	}
	{
		p.SetState(61)
		p.Deltalist()
	}

	return localctx
}

// IAdminContext is an interface to support dynamic dispatch.
type IAdminContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminContext differentiates from other interfaces.
	IsAdminContext()
}

type AdminContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminContext() *AdminContext {
	var p = new(AdminContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_admin
	return p
}

func (*AdminContext) IsAdminContext() {}

func NewAdminContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminContext {
	var p = new(AdminContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_admin

	return p
}

func (s *AdminContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminContext) Head() IHeadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeadContext)
}

func (s *AdminContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *AdminContext) Symbols() ISymbolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolsContext)
}

func (s *AdminContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *AdminContext) Locks() ILocksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocksContext)
}

func (s *AdminContext) Strict() IStrictContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrictContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrictContext)
}

func (s *AdminContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *AdminContext) Expand() IExpandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpandContext)
}

func (s *AdminContext) AllNewphrase() []INewphraseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewphraseContext)(nil)).Elem())
	var tst = make([]INewphraseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewphraseContext)
		}
	}

	return tst
}

func (s *AdminContext) Newphrase(i int) INewphraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewphraseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewphraseContext)
}

func (s *AdminContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterAdmin(s)
	}
}

func (s *AdminContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitAdmin(s)
	}
}

func (s *AdminContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitAdmin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Admin() (localctx IAdminContext) {
	localctx = NewAdminContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RCSParserRULE_admin)
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
		p.SetState(63)
		p.Head()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserLITERAL_BRANCH {
		{
			p.SetState(64)
			p.Branch()
		}

	}
	{
		p.SetState(67)
		p.Access()
	}
	{
		p.SetState(68)
		p.Symbols()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserLITERAL_LOCKS {
		{
			p.SetState(69)
			p.Locks()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserLITERAL_STRICT {
		{
			p.SetState(72)
			p.Strict()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserLITERAL_COMMENT {
		{
			p.SetState(75)
			p.Comment()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserLITERAL_EXPAND {
		{
			p.SetState(78)
			p.Expand()
		}

	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(81)
			p.Newphrase()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeadContext is an interface to support dynamic dispatch.
type IHeadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeadContext differentiates from other interfaces.
	IsHeadContext()
}

type HeadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeadContext() *HeadContext {
	var p = new(HeadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_head
	return p
}

func (*HeadContext) IsHeadContext() {}

func NewHeadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadContext {
	var p = new(HeadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_head

	return p
}

func (s *HeadContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadContext) LITERAL_HEAD() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_HEAD, 0)
}

func (s *HeadContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *HeadContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *HeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterHead(s)
	}
}

func (s *HeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitHead(s)
	}
}

func (s *HeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Head() (localctx IHeadContext) {
	localctx = NewHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RCSParserRULE_head)

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
		p.Match(RCSParserLITERAL_HEAD)
	}
	{
		p.SetState(88)
		p.Match(RCSParserREVISION)
	}
	{
		p.SetState(89)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) LITERAL_BRANCH() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_BRANCH, 0)
}

func (s *BranchContext) BRANCH() antlr.TerminalNode {
	return s.GetToken(RCSParserBRANCH, 0)
}

func (s *BranchContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (s *BranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Branch() (localctx IBranchContext) {
	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RCSParserRULE_branch)

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
		p.Match(RCSParserLITERAL_BRANCH)
	}
	{
		p.SetState(92)
		p.Match(RCSParserBRANCH)
	}
	{
		p.SetState(93)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IAccessContext is an interface to support dynamic dispatch.
type IAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessContext differentiates from other interfaces.
	IsAccessContext()
}

type AccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessContext() *AccessContext {
	var p = new(AccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_access
	return p
}

func (*AccessContext) IsAccessContext() {}

func NewAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessContext {
	var p = new(AccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_access

	return p
}

func (s *AccessContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessContext) LITERAL_ACCESS() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_ACCESS, 0)
}

func (s *AccessContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *AccessContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(RCSParserIDENT)
}

func (s *AccessContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, i)
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitAccess(s)
	}
}

func (s *AccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Access() (localctx IAccessContext) {
	localctx = NewAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RCSParserRULE_access)
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
		p.SetState(95)
		p.Match(RCSParserLITERAL_ACCESS)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(96)
			p.Match(RCSParserIDENT)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// ISymbolsContext is an interface to support dynamic dispatch.
type ISymbolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolsContext differentiates from other interfaces.
	IsSymbolsContext()
}

type SymbolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolsContext() *SymbolsContext {
	var p = new(SymbolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_symbols
	return p
}

func (*SymbolsContext) IsSymbolsContext() {}

func NewSymbolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolsContext {
	var p = new(SymbolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_symbols

	return p
}

func (s *SymbolsContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolsContext) LITERAL_SYMBOLS() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_SYMBOLS, 0)
}

func (s *SymbolsContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *SymbolsContext) AllTags() []ITagsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagsContext)(nil)).Elem())
	var tst = make([]ITagsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagsContext)
		}
	}

	return tst
}

func (s *SymbolsContext) Tags(i int) ITagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagsContext)
}

func (s *SymbolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterSymbols(s)
	}
}

func (s *SymbolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitSymbols(s)
	}
}

func (s *SymbolsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitSymbols(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Symbols() (localctx ISymbolsContext) {
	localctx = NewSymbolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RCSParserRULE_symbols)
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
		p.SetState(104)
		p.Match(RCSParserLITERAL_SYMBOLS)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(105)
			p.Tags()
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(111)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// ITagsContext is an interface to support dynamic dispatch.
type ITagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagsContext differentiates from other interfaces.
	IsTagsContext()
}

type TagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagsContext() *TagsContext {
	var p = new(TagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_tags
	return p
}

func (*TagsContext) IsTagsContext() {}

func NewTagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagsContext {
	var p = new(TagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_tags

	return p
}

func (s *TagsContext) GetParser() antlr.Parser { return s.parser }

func (s *TagsContext) IDENT() antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, 0)
}

func (s *TagsContext) COLON() antlr.TerminalNode {
	return s.GetToken(RCSParserCOLON, 0)
}

func (s *TagsContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Tags() (localctx ITagsContext) {
	localctx = NewTagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RCSParserRULE_tags)

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
		p.Match(RCSParserIDENT)
	}
	{
		p.SetState(114)
		p.Match(RCSParserCOLON)
	}
	{
		p.SetState(115)
		p.Match(RCSParserREVISION)
	}

	return localctx
}

// ILocksContext is an interface to support dynamic dispatch.
type ILocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocksContext differentiates from other interfaces.
	IsLocksContext()
}

type LocksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocksContext() *LocksContext {
	var p = new(LocksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_locks
	return p
}

func (*LocksContext) IsLocksContext() {}

func NewLocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocksContext {
	var p = new(LocksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_locks

	return p
}

func (s *LocksContext) GetParser() antlr.Parser { return s.parser }

func (s *LocksContext) LITERAL_LOCKS() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_LOCKS, 0)
}

func (s *LocksContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *LocksContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(RCSParserIDENT)
}

func (s *LocksContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, i)
}

func (s *LocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterLocks(s)
	}
}

func (s *LocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitLocks(s)
	}
}

func (s *LocksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitLocks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Locks() (localctx ILocksContext) {
	localctx = NewLocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RCSParserRULE_locks)
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
		p.SetState(117)
		p.Match(RCSParserLITERAL_LOCKS)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(118)
			p.Match(RCSParserIDENT)
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IStrictContext is an interface to support dynamic dispatch.
type IStrictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrictContext differentiates from other interfaces.
	IsStrictContext()
}

type StrictContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrictContext() *StrictContext {
	var p = new(StrictContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_strict
	return p
}

func (*StrictContext) IsStrictContext() {}

func NewStrictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrictContext {
	var p = new(StrictContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_strict

	return p
}

func (s *StrictContext) GetParser() antlr.Parser { return s.parser }

func (s *StrictContext) LITERAL_STRICT() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_STRICT, 0)
}

func (s *StrictContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *StrictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterStrict(s)
	}
}

func (s *StrictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitStrict(s)
	}
}

func (s *StrictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitStrict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Strict() (localctx IStrictContext) {
	localctx = NewStrictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RCSParserRULE_strict)

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
		p.SetState(126)
		p.Match(RCSParserLITERAL_STRICT)
	}
	{
		p.SetState(127)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LITERAL_COMMENT() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_COMMENT, 0)
}

func (s *CommentContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *CommentContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCSParserSTRING, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RCSParserRULE_comment)
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
		p.Match(RCSParserLITERAL_COMMENT)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserSTRING {
		{
			p.SetState(130)
			p.Match(RCSParserSTRING)
		}

	}
	{
		p.SetState(133)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IExpandContext is an interface to support dynamic dispatch.
type IExpandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpandContext differentiates from other interfaces.
	IsExpandContext()
}

type ExpandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpandContext() *ExpandContext {
	var p = new(ExpandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_expand
	return p
}

func (*ExpandContext) IsExpandContext() {}

func NewExpandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpandContext {
	var p = new(ExpandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_expand

	return p
}

func (s *ExpandContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpandContext) LITERAL_EXPAND() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_EXPAND, 0)
}

func (s *ExpandContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *ExpandContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCSParserSTRING, 0)
}

func (s *ExpandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterExpand(s)
	}
}

func (s *ExpandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitExpand(s)
	}
}

func (s *ExpandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitExpand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Expand() (localctx IExpandContext) {
	localctx = NewExpandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RCSParserRULE_expand)
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
		p.SetState(135)
		p.Match(RCSParserLITERAL_EXPAND)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserSTRING {
		{
			p.SetState(136)
			p.Match(RCSParserSTRING)
		}

	}
	{
		p.SetState(139)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDeltalistContext is an interface to support dynamic dispatch.
type IDeltalistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltalistContext differentiates from other interfaces.
	IsDeltalistContext()
}

type DeltalistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltalistContext() *DeltalistContext {
	var p = new(DeltalistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_deltalist
	return p
}

func (*DeltalistContext) IsDeltalistContext() {}

func NewDeltalistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeltalistContext {
	var p = new(DeltalistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_deltalist

	return p
}

func (s *DeltalistContext) GetParser() antlr.Parser { return s.parser }

func (s *DeltalistContext) AllDelta() []IDeltaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeltaContext)(nil)).Elem())
	var tst = make([]IDeltaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeltaContext)
		}
	}

	return tst
}

func (s *DeltalistContext) Delta(i int) IDeltaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeltaContext)
}

func (s *DeltalistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeltalistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeltalistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDeltalist(s)
	}
}

func (s *DeltalistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDeltalist(s)
	}
}

func (s *DeltalistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDeltalist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Deltalist() (localctx IDeltalistContext) {
	localctx = NewDeltalistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RCSParserRULE_deltalist)
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
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserREVISION {
		{
			p.SetState(141)
			p.Delta()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeltaContext is an interface to support dynamic dispatch.
type IDeltaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltaContext differentiates from other interfaces.
	IsDeltaContext()
}

type DeltaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltaContext() *DeltaContext {
	var p = new(DeltaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta
	return p
}

func (*DeltaContext) IsDeltaContext() {}

func NewDeltaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeltaContext {
	var p = new(DeltaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta

	return p
}

func (s *DeltaContext) GetParser() antlr.Parser { return s.parser }

func (s *DeltaContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *DeltaContext) Delta_date() IDelta_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelta_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelta_dateContext)
}

func (s *DeltaContext) Delta_author() IDelta_authorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelta_authorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelta_authorContext)
}

func (s *DeltaContext) Delta_state() IDelta_stateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelta_stateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelta_stateContext)
}

func (s *DeltaContext) Delta_branches() IDelta_branchesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelta_branchesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelta_branchesContext)
}

func (s *DeltaContext) Delta_next() IDelta_nextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelta_nextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelta_nextContext)
}

func (s *DeltaContext) AllNewphrase() []INewphraseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewphraseContext)(nil)).Elem())
	var tst = make([]INewphraseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewphraseContext)
		}
	}

	return tst
}

func (s *DeltaContext) Newphrase(i int) INewphraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewphraseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewphraseContext)
}

func (s *DeltaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeltaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeltaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta(s)
	}
}

func (s *DeltaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta(s)
	}
}

func (s *DeltaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta() (localctx IDeltaContext) {
	localctx = NewDeltaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RCSParserRULE_delta)
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
		p.SetState(147)
		p.Match(RCSParserREVISION)
	}
	{
		p.SetState(148)
		p.Delta_date()
	}
	{
		p.SetState(149)
		p.Delta_author()
	}
	{
		p.SetState(150)
		p.Delta_state()
	}
	{
		p.SetState(151)
		p.Delta_branches()
	}
	{
		p.SetState(152)
		p.Delta_next()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(153)
			p.Newphrase()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDelta_dateContext is an interface to support dynamic dispatch.
type IDelta_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelta_dateContext differentiates from other interfaces.
	IsDelta_dateContext()
}

type Delta_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelta_dateContext() *Delta_dateContext {
	var p = new(Delta_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta_date
	return p
}

func (*Delta_dateContext) IsDelta_dateContext() {}

func NewDelta_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delta_dateContext {
	var p = new(Delta_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta_date

	return p
}

func (s *Delta_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Delta_dateContext) LITERAL_DATE() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_DATE, 0)
}

func (s *Delta_dateContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *Delta_dateContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *Delta_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delta_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delta_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta_date(s)
	}
}

func (s *Delta_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta_date(s)
	}
}

func (s *Delta_dateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta_date(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta_date() (localctx IDelta_dateContext) {
	localctx = NewDelta_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RCSParserRULE_delta_date)

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
		p.Match(RCSParserLITERAL_DATE)
	}
	{
		p.SetState(160)
		p.Match(RCSParserREVISION)
	}
	{
		p.SetState(161)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDelta_authorContext is an interface to support dynamic dispatch.
type IDelta_authorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelta_authorContext differentiates from other interfaces.
	IsDelta_authorContext()
}

type Delta_authorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelta_authorContext() *Delta_authorContext {
	var p = new(Delta_authorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta_author
	return p
}

func (*Delta_authorContext) IsDelta_authorContext() {}

func NewDelta_authorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delta_authorContext {
	var p = new(Delta_authorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta_author

	return p
}

func (s *Delta_authorContext) GetParser() antlr.Parser { return s.parser }

func (s *Delta_authorContext) LITERAL_AUTHOR() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_AUTHOR, 0)
}

func (s *Delta_authorContext) IDENT() antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, 0)
}

func (s *Delta_authorContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *Delta_authorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delta_authorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delta_authorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta_author(s)
	}
}

func (s *Delta_authorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta_author(s)
	}
}

func (s *Delta_authorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta_author(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta_author() (localctx IDelta_authorContext) {
	localctx = NewDelta_authorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RCSParserRULE_delta_author)

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
		p.SetState(163)
		p.Match(RCSParserLITERAL_AUTHOR)
	}
	{
		p.SetState(164)
		p.Match(RCSParserIDENT)
	}
	{
		p.SetState(165)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDelta_stateContext is an interface to support dynamic dispatch.
type IDelta_stateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelta_stateContext differentiates from other interfaces.
	IsDelta_stateContext()
}

type Delta_stateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelta_stateContext() *Delta_stateContext {
	var p = new(Delta_stateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta_state
	return p
}

func (*Delta_stateContext) IsDelta_stateContext() {}

func NewDelta_stateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delta_stateContext {
	var p = new(Delta_stateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta_state

	return p
}

func (s *Delta_stateContext) GetParser() antlr.Parser { return s.parser }

func (s *Delta_stateContext) LITERAL_STATE() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_STATE, 0)
}

func (s *Delta_stateContext) IDENT() antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, 0)
}

func (s *Delta_stateContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *Delta_stateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delta_stateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delta_stateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta_state(s)
	}
}

func (s *Delta_stateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta_state(s)
	}
}

func (s *Delta_stateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta_state(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta_state() (localctx IDelta_stateContext) {
	localctx = NewDelta_stateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RCSParserRULE_delta_state)

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
		p.SetState(167)
		p.Match(RCSParserLITERAL_STATE)
	}
	{
		p.SetState(168)
		p.Match(RCSParserIDENT)
	}
	{
		p.SetState(169)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDelta_branchesContext is an interface to support dynamic dispatch.
type IDelta_branchesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelta_branchesContext differentiates from other interfaces.
	IsDelta_branchesContext()
}

type Delta_branchesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelta_branchesContext() *Delta_branchesContext {
	var p = new(Delta_branchesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta_branches
	return p
}

func (*Delta_branchesContext) IsDelta_branchesContext() {}

func NewDelta_branchesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delta_branchesContext {
	var p = new(Delta_branchesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta_branches

	return p
}

func (s *Delta_branchesContext) GetParser() antlr.Parser { return s.parser }

func (s *Delta_branchesContext) LITERAL_BRANCHES() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_BRANCHES, 0)
}

func (s *Delta_branchesContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *Delta_branchesContext) AllREVISION() []antlr.TerminalNode {
	return s.GetTokens(RCSParserREVISION)
}

func (s *Delta_branchesContext) REVISION(i int) antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, i)
}

func (s *Delta_branchesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delta_branchesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delta_branchesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta_branches(s)
	}
}

func (s *Delta_branchesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta_branches(s)
	}
}

func (s *Delta_branchesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta_branches(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta_branches() (localctx IDelta_branchesContext) {
	localctx = NewDelta_branchesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RCSParserRULE_delta_branches)
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
		p.SetState(171)
		p.Match(RCSParserLITERAL_BRANCHES)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserREVISION {
		{
			p.SetState(172)
			p.Match(RCSParserREVISION)
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDelta_nextContext is an interface to support dynamic dispatch.
type IDelta_nextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelta_nextContext differentiates from other interfaces.
	IsDelta_nextContext()
}

type Delta_nextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelta_nextContext() *Delta_nextContext {
	var p = new(Delta_nextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_delta_next
	return p
}

func (*Delta_nextContext) IsDelta_nextContext() {}

func NewDelta_nextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delta_nextContext {
	var p = new(Delta_nextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_delta_next

	return p
}

func (s *Delta_nextContext) GetParser() antlr.Parser { return s.parser }

func (s *Delta_nextContext) LITERAL_NEXT() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_NEXT, 0)
}

func (s *Delta_nextContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *Delta_nextContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *Delta_nextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delta_nextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delta_nextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDelta_next(s)
	}
}

func (s *Delta_nextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDelta_next(s)
	}
}

func (s *Delta_nextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDelta_next(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Delta_next() (localctx IDelta_nextContext) {
	localctx = NewDelta_nextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RCSParserRULE_delta_next)
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
		p.SetState(180)
		p.Match(RCSParserLITERAL_NEXT)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RCSParserREVISION {
		{
			p.SetState(181)
			p.Match(RCSParserREVISION)
		}

	}
	{
		p.SetState(184)
		p.Match(RCSParserSEMI)
	}

	return localctx
}

// IDescContext is an interface to support dynamic dispatch.
type IDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescContext differentiates from other interfaces.
	IsDescContext()
}

type DescContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescContext() *DescContext {
	var p = new(DescContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_desc
	return p
}

func (*DescContext) IsDescContext() {}

func NewDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescContext {
	var p = new(DescContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_desc

	return p
}

func (s *DescContext) GetParser() antlr.Parser { return s.parser }

func (s *DescContext) LITERAL_DESC() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_DESC, 0)
}

func (s *DescContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCSParserSTRING, 0)
}

func (s *DescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDesc(s)
	}
}

func (s *DescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDesc(s)
	}
}

func (s *DescContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDesc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Desc() (localctx IDescContext) {
	localctx = NewDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RCSParserRULE_desc)

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
		p.SetState(186)
		p.Match(RCSParserLITERAL_DESC)
	}
	{
		p.SetState(187)
		p.Match(RCSParserSTRING)
	}

	return localctx
}

// IDeltatextlistContext is an interface to support dynamic dispatch.
type IDeltatextlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltatextlistContext differentiates from other interfaces.
	IsDeltatextlistContext()
}

type DeltatextlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltatextlistContext() *DeltatextlistContext {
	var p = new(DeltatextlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_deltatextlist
	return p
}

func (*DeltatextlistContext) IsDeltatextlistContext() {}

func NewDeltatextlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeltatextlistContext {
	var p = new(DeltatextlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_deltatextlist

	return p
}

func (s *DeltatextlistContext) GetParser() antlr.Parser { return s.parser }

func (s *DeltatextlistContext) AllDeltatext() []IDeltatextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeltatextContext)(nil)).Elem())
	var tst = make([]IDeltatextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeltatextContext)
		}
	}

	return tst
}

func (s *DeltatextlistContext) Deltatext(i int) IDeltatextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltatextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeltatextContext)
}

func (s *DeltatextlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeltatextlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeltatextlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDeltatextlist(s)
	}
}

func (s *DeltatextlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDeltatextlist(s)
	}
}

func (s *DeltatextlistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDeltatextlist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Deltatextlist() (localctx IDeltatextlistContext) {
	localctx = NewDeltatextlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RCSParserRULE_deltatextlist)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserREVISION {
		{
			p.SetState(189)
			p.Deltatext()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeltatextContext is an interface to support dynamic dispatch.
type IDeltatextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltatextContext differentiates from other interfaces.
	IsDeltatextContext()
}

type DeltatextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltatextContext() *DeltatextContext {
	var p = new(DeltatextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_deltatext
	return p
}

func (*DeltatextContext) IsDeltatextContext() {}

func NewDeltatextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeltatextContext {
	var p = new(DeltatextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_deltatext

	return p
}

func (s *DeltatextContext) GetParser() antlr.Parser { return s.parser }

func (s *DeltatextContext) REVISION() antlr.TerminalNode {
	return s.GetToken(RCSParserREVISION, 0)
}

func (s *DeltatextContext) Deltatext_log() IDeltatext_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltatext_logContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeltatext_logContext)
}

func (s *DeltatextContext) Deltatext_text() IDeltatext_textContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeltatext_textContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeltatext_textContext)
}

func (s *DeltatextContext) AllNewphrase() []INewphraseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewphraseContext)(nil)).Elem())
	var tst = make([]INewphraseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewphraseContext)
		}
	}

	return tst
}

func (s *DeltatextContext) Newphrase(i int) INewphraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewphraseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewphraseContext)
}

func (s *DeltatextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeltatextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeltatextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDeltatext(s)
	}
}

func (s *DeltatextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDeltatext(s)
	}
}

func (s *DeltatextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDeltatext(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Deltatext() (localctx IDeltatextContext) {
	localctx = NewDeltatextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RCSParserRULE_deltatext)
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
		p.SetState(195)
		p.Match(RCSParserREVISION)
	}
	{
		p.SetState(196)
		p.Deltatext_log()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RCSParserIDENT {
		{
			p.SetState(197)
			p.Newphrase()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Deltatext_text()
	}

	return localctx
}

// IDeltatext_logContext is an interface to support dynamic dispatch.
type IDeltatext_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltatext_logContext differentiates from other interfaces.
	IsDeltatext_logContext()
}

type Deltatext_logContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltatext_logContext() *Deltatext_logContext {
	var p = new(Deltatext_logContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_deltatext_log
	return p
}

func (*Deltatext_logContext) IsDeltatext_logContext() {}

func NewDeltatext_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Deltatext_logContext {
	var p = new(Deltatext_logContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_deltatext_log

	return p
}

func (s *Deltatext_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Deltatext_logContext) LITERAL_LOG() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_LOG, 0)
}

func (s *Deltatext_logContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCSParserSTRING, 0)
}

func (s *Deltatext_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Deltatext_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Deltatext_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDeltatext_log(s)
	}
}

func (s *Deltatext_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDeltatext_log(s)
	}
}

func (s *Deltatext_logContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDeltatext_log(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Deltatext_log() (localctx IDeltatext_logContext) {
	localctx = NewDeltatext_logContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RCSParserRULE_deltatext_log)

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
		p.SetState(205)
		p.Match(RCSParserLITERAL_LOG)
	}
	{
		p.SetState(206)
		p.Match(RCSParserSTRING)
	}

	return localctx
}

// IDeltatext_textContext is an interface to support dynamic dispatch.
type IDeltatext_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltatext_textContext differentiates from other interfaces.
	IsDeltatext_textContext()
}

type Deltatext_textContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltatext_textContext() *Deltatext_textContext {
	var p = new(Deltatext_textContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_deltatext_text
	return p
}

func (*Deltatext_textContext) IsDeltatext_textContext() {}

func NewDeltatext_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Deltatext_textContext {
	var p = new(Deltatext_textContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_deltatext_text

	return p
}

func (s *Deltatext_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Deltatext_textContext) LITERAL_TEXT() antlr.TerminalNode {
	return s.GetToken(RCSParserLITERAL_TEXT, 0)
}

func (s *Deltatext_textContext) STRING() antlr.TerminalNode {
	return s.GetToken(RCSParserSTRING, 0)
}

func (s *Deltatext_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Deltatext_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Deltatext_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterDeltatext_text(s)
	}
}

func (s *Deltatext_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitDeltatext_text(s)
	}
}

func (s *Deltatext_textContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitDeltatext_text(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Deltatext_text() (localctx IDeltatext_textContext) {
	localctx = NewDeltatext_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RCSParserRULE_deltatext_text)

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
		p.SetState(208)
		p.Match(RCSParserLITERAL_TEXT)
	}
	{
		p.SetState(209)
		p.Match(RCSParserSTRING)
	}

	return localctx
}

// INewphraseContext is an interface to support dynamic dispatch.
type INewphraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewphraseContext differentiates from other interfaces.
	IsNewphraseContext()
}

type NewphraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewphraseContext() *NewphraseContext {
	var p = new(NewphraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCSParserRULE_newphrase
	return p
}

func (*NewphraseContext) IsNewphraseContext() {}

func NewNewphraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewphraseContext {
	var p = new(NewphraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCSParserRULE_newphrase

	return p
}

func (s *NewphraseContext) GetParser() antlr.Parser { return s.parser }

func (s *NewphraseContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RCSParserSEMI, 0)
}

func (s *NewphraseContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(RCSParserIDENT)
}

func (s *NewphraseContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(RCSParserIDENT, i)
}

func (s *NewphraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewphraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewphraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.EnterNewphrase(s)
	}
}

func (s *NewphraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCSListener); ok {
		listenerT.ExitNewphrase(s)
	}
}

func (s *NewphraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RCSVisitor:
		return t.VisitNewphrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RCSParser) Newphrase() (localctx INewphraseContext) {
	localctx = NewNewphraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RCSParserRULE_newphrase)
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
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RCSParserIDENT {
		{
			p.SetState(211)
			p.Match(RCSParserIDENT)
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(RCSParserSEMI)
	}

	return localctx
}
