// Code generated from CLIF.g4 by ANTLR 4.7.2. DO NOT EDIT.

package clif // CLIF
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 219,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 7, 2, 45, 10,
	2, 12, 2, 14, 2, 48, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 56,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 70, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 84, 10, 9, 3, 10, 3, 10, 5, 10, 88, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 107, 10, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 114, 10, 13, 12, 13, 14, 13, 117, 11,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 131, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 136, 10,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	5, 15, 148, 10, 15, 3, 15, 3, 15, 3, 15, 7, 15, 153, 10, 15, 12, 15, 14,
	15, 156, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 172, 10, 17, 12, 17, 14,
	17, 175, 11, 17, 3, 17, 5, 17, 178, 10, 17, 3, 17, 5, 17, 181, 10, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 196, 10, 18, 3, 18, 5, 18, 199, 10, 18, 3, 19,
	6, 19, 202, 10, 19, 13, 19, 14, 19, 203, 3, 20, 3, 20, 3, 20, 5, 20, 209,
	10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 215, 10, 21, 3, 21, 3, 21, 3,
	21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 2, 7, 4, 2, 8, 8, 10, 10, 4, 2, 11, 11, 27, 27, 3, 2, 13,
	14, 3, 2, 15, 16, 3, 2, 17, 18, 2, 227, 2, 46, 3, 2, 2, 2, 4, 49, 3, 2,
	2, 2, 6, 51, 3, 2, 2, 2, 8, 55, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 71,
	3, 2, 2, 2, 14, 73, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 87, 3, 2, 2, 2,
	20, 106, 3, 2, 2, 2, 22, 108, 3, 2, 2, 2, 24, 130, 3, 2, 2, 2, 26, 132,
	3, 2, 2, 2, 28, 141, 3, 2, 2, 2, 30, 159, 3, 2, 2, 2, 32, 165, 3, 2, 2,
	2, 34, 198, 3, 2, 2, 2, 36, 201, 3, 2, 2, 2, 38, 208, 3, 2, 2, 2, 40, 210,
	3, 2, 2, 2, 42, 45, 5, 10, 6, 2, 43, 45, 7, 9, 2, 2, 44, 42, 3, 2, 2, 2,
	44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 3, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 9, 2, 2, 2, 50,
	5, 3, 2, 2, 2, 51, 52, 9, 3, 2, 2, 52, 7, 3, 2, 2, 2, 53, 56, 5, 4, 3,
	2, 54, 56, 5, 6, 4, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 9, 3,
	2, 2, 2, 57, 70, 5, 8, 5, 2, 58, 59, 7, 3, 2, 2, 59, 60, 5, 12, 7, 2, 60,
	61, 5, 2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 70, 3, 2, 2, 2, 63, 64, 7, 3, 2,
	2, 64, 65, 7, 25, 2, 2, 65, 66, 7, 10, 2, 2, 66, 67, 5, 10, 6, 2, 67, 68,
	7, 4, 2, 2, 68, 70, 3, 2, 2, 2, 69, 57, 3, 2, 2, 2, 69, 58, 3, 2, 2, 2,
	69, 63, 3, 2, 2, 2, 70, 11, 3, 2, 2, 2, 71, 72, 5, 10, 6, 2, 72, 13, 3,
	2, 2, 2, 73, 74, 7, 3, 2, 2, 74, 75, 7, 12, 2, 2, 75, 76, 5, 10, 6, 2,
	76, 77, 5, 10, 6, 2, 77, 78, 7, 4, 2, 2, 78, 15, 3, 2, 2, 2, 79, 84, 5,
	18, 10, 2, 80, 84, 5, 24, 13, 2, 81, 84, 5, 26, 14, 2, 82, 84, 5, 30, 16,
	2, 83, 79, 3, 2, 2, 2, 83, 80, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 82,
	3, 2, 2, 2, 84, 17, 3, 2, 2, 2, 85, 88, 5, 14, 8, 2, 86, 88, 5, 20, 11,
	2, 87, 85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 19, 3, 2, 2, 2, 89, 90,
	7, 3, 2, 2, 90, 91, 5, 22, 12, 2, 91, 92, 5, 2, 2, 2, 92, 93, 7, 4, 2,
	2, 93, 107, 3, 2, 2, 2, 94, 95, 7, 3, 2, 2, 95, 96, 5, 10, 6, 2, 96, 97,
	7, 3, 2, 2, 97, 98, 7, 20, 2, 2, 98, 99, 7, 3, 2, 2, 99, 100, 5, 8, 5,
	2, 100, 101, 5, 10, 6, 2, 101, 102, 7, 4, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	104, 7, 4, 2, 2, 104, 105, 7, 4, 2, 2, 105, 107, 3, 2, 2, 2, 106, 89, 3,
	2, 2, 2, 106, 94, 3, 2, 2, 2, 107, 21, 3, 2, 2, 2, 108, 109, 5, 10, 6,
	2, 109, 23, 3, 2, 2, 2, 110, 111, 7, 3, 2, 2, 111, 115, 9, 4, 2, 2, 112,
	114, 5, 16, 9, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 115, 3, 2,
	2, 2, 118, 131, 7, 4, 2, 2, 119, 120, 7, 3, 2, 2, 120, 121, 9, 5, 2, 2,
	121, 122, 5, 16, 9, 2, 122, 123, 5, 16, 9, 2, 123, 124, 7, 4, 2, 2, 124,
	131, 3, 2, 2, 2, 125, 126, 7, 3, 2, 2, 126, 127, 7, 19, 2, 2, 127, 128,
	5, 16, 9, 2, 128, 129, 7, 4, 2, 2, 129, 131, 3, 2, 2, 2, 130, 110, 3, 2,
	2, 2, 130, 119, 3, 2, 2, 2, 130, 125, 3, 2, 2, 2, 131, 25, 3, 2, 2, 2,
	132, 133, 7, 3, 2, 2, 133, 135, 9, 6, 2, 2, 134, 136, 5, 6, 4, 2, 135,
	134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138,
	5, 28, 15, 2, 138, 139, 5, 16, 9, 2, 139, 140, 7, 4, 2, 2, 140, 27, 3,
	2, 2, 2, 141, 154, 7, 3, 2, 2, 142, 153, 5, 6, 4, 2, 143, 153, 7, 9, 2,
	2, 144, 147, 7, 3, 2, 2, 145, 148, 5, 6, 4, 2, 146, 148, 7, 9, 2, 2, 147,
	145, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150,
	5, 10, 6, 2, 150, 151, 7, 4, 2, 2, 151, 153, 3, 2, 2, 2, 152, 142, 3, 2,
	2, 2, 152, 143, 3, 2, 2, 2, 152, 144, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2,
	154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 156,
	154, 3, 2, 2, 2, 157, 158, 7, 4, 2, 2, 158, 29, 3, 2, 2, 2, 159, 160, 7,
	3, 2, 2, 160, 161, 7, 25, 2, 2, 161, 162, 7, 11, 2, 2, 162, 163, 5, 16,
	9, 2, 163, 164, 7, 4, 2, 2, 164, 31, 3, 2, 2, 2, 165, 166, 7, 3, 2, 2,
	166, 167, 7, 24, 2, 2, 167, 177, 5, 6, 4, 2, 168, 169, 7, 3, 2, 2, 169,
	173, 7, 23, 2, 2, 170, 172, 5, 8, 5, 2, 171, 170, 3, 2, 2, 2, 172, 175,
	3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 176, 3, 2,
	2, 2, 175, 173, 3, 2, 2, 2, 176, 178, 7, 4, 2, 2, 177, 168, 3, 2, 2, 2,
	177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179, 181, 5, 38, 20, 2, 180,
	179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183,
	7, 4, 2, 2, 183, 33, 3, 2, 2, 2, 184, 199, 5, 16, 9, 2, 185, 199, 5, 32,
	17, 2, 186, 187, 7, 3, 2, 2, 187, 188, 7, 22, 2, 2, 188, 189, 5, 6, 4,
	2, 189, 190, 7, 4, 2, 2, 190, 199, 3, 2, 2, 2, 191, 192, 7, 3, 2, 2, 192,
	193, 7, 25, 2, 2, 193, 195, 7, 11, 2, 2, 194, 196, 5, 38, 20, 2, 195, 194,
	3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 199, 7, 4,
	2, 2, 198, 184, 3, 2, 2, 2, 198, 185, 3, 2, 2, 2, 198, 186, 3, 2, 2, 2,
	198, 191, 3, 2, 2, 2, 199, 35, 3, 2, 2, 2, 200, 202, 5, 34, 18, 2, 201,
	200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204,
	3, 2, 2, 2, 204, 37, 3, 2, 2, 2, 205, 209, 5, 32, 17, 2, 206, 209, 5, 40,
	21, 2, 207, 209, 5, 36, 19, 2, 208, 205, 3, 2, 2, 2, 208, 206, 3, 2, 2,
	2, 208, 207, 3, 2, 2, 2, 209, 39, 3, 2, 2, 2, 210, 211, 7, 3, 2, 2, 211,
	212, 7, 21, 2, 2, 212, 214, 5, 6, 4, 2, 213, 215, 5, 36, 19, 2, 214, 213,
	3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 7, 4,
	2, 2, 217, 41, 3, 2, 2, 2, 23, 44, 46, 55, 69, 83, 87, 106, 115, 130, 135,
	147, 152, 154, 173, 177, 180, 195, 198, 203, 208, 214,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'''", "'\"'", "'\\'", "", "", "", "", "'='", "'and'",
	"'or'", "'iff'", "'if'", "'forall'", "'exists'", "'not'", "'cl-roleset'",
	"'cl-text'", "'cl-imports'", "'cl-excludes'", "'cl-module'", "'cl-comment'",
	"'cl-prefix'",
}
var symbolicNames = []string{
	"", "OPEN", "CLOSE", "STRINGQUOTE", "NAMEQUOTE", "BACKSLASH", "NUMERAL",
	"SEQMARK", "QUOTEDSTRING", "ENCLOSEDNAME", "EQUAL", "AND", "OR", "IFF",
	"IF", "FORALL", "EXISTS", "NOT", "CL_ROLESET", "CL_TEXT", "CL_IMPORTS",
	"CL_EXCLUDES", "CL_MODULE", "CL_COMMENT", "CL_PREFIX", "NAMECHARSEQUENCE",
	"WHITE", "BLOCKCOMMENT", "LineComment",
}

var ruleNames = []string{
	"termseq", "interpretedname", "interpretablename", "name", "term", "operator",
	"equation", "sentence", "atomsent", "atom", "predicate", "boolsent", "quantsent",
	"boundlist", "commentsent", "module", "phrase", "text", "cltext", "namedtext",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CLIFParser struct {
	*antlr.BaseParser
}

func NewCLIFParser(input antlr.TokenStream) *CLIFParser {
	this := new(CLIFParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CLIF.g4"

	return this
}

// CLIFParser tokens.
const (
	CLIFParserEOF              = antlr.TokenEOF
	CLIFParserOPEN             = 1
	CLIFParserCLOSE            = 2
	CLIFParserSTRINGQUOTE      = 3
	CLIFParserNAMEQUOTE        = 4
	CLIFParserBACKSLASH        = 5
	CLIFParserNUMERAL          = 6
	CLIFParserSEQMARK          = 7
	CLIFParserQUOTEDSTRING     = 8
	CLIFParserENCLOSEDNAME     = 9
	CLIFParserEQUAL            = 10
	CLIFParserAND              = 11
	CLIFParserOR               = 12
	CLIFParserIFF              = 13
	CLIFParserIF               = 14
	CLIFParserFORALL           = 15
	CLIFParserEXISTS           = 16
	CLIFParserNOT              = 17
	CLIFParserCL_ROLESET       = 18
	CLIFParserCL_TEXT          = 19
	CLIFParserCL_IMPORTS       = 20
	CLIFParserCL_EXCLUDES      = 21
	CLIFParserCL_MODULE        = 22
	CLIFParserCL_COMMENT       = 23
	CLIFParserCL_PREFIX        = 24
	CLIFParserNAMECHARSEQUENCE = 25
	CLIFParserWHITE            = 26
	CLIFParserBLOCKCOMMENT     = 27
	CLIFParserLineComment      = 28
)

// CLIFParser rules.
const (
	CLIFParserRULE_termseq           = 0
	CLIFParserRULE_interpretedname   = 1
	CLIFParserRULE_interpretablename = 2
	CLIFParserRULE_name              = 3
	CLIFParserRULE_term              = 4
	CLIFParserRULE_operator          = 5
	CLIFParserRULE_equation          = 6
	CLIFParserRULE_sentence          = 7
	CLIFParserRULE_atomsent          = 8
	CLIFParserRULE_atom              = 9
	CLIFParserRULE_predicate         = 10
	CLIFParserRULE_boolsent          = 11
	CLIFParserRULE_quantsent         = 12
	CLIFParserRULE_boundlist         = 13
	CLIFParserRULE_commentsent       = 14
	CLIFParserRULE_module            = 15
	CLIFParserRULE_phrase            = 16
	CLIFParserRULE_text              = 17
	CLIFParserRULE_cltext            = 18
	CLIFParserRULE_namedtext         = 19
)

// ITermseqContext is an interface to support dynamic dispatch.
type ITermseqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermseqContext differentiates from other interfaces.
	IsTermseqContext()
}

type TermseqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermseqContext() *TermseqContext {
	var p = new(TermseqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_termseq
	return p
}

func (*TermseqContext) IsTermseqContext() {}

func NewTermseqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermseqContext {
	var p = new(TermseqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_termseq

	return p
}

func (s *TermseqContext) GetParser() antlr.Parser { return s.parser }

func (s *TermseqContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TermseqContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermseqContext) AllSEQMARK() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserSEQMARK)
}

func (s *TermseqContext) SEQMARK(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserSEQMARK, i)
}

func (s *TermseqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermseqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermseqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterTermseq(s)
	}
}

func (s *TermseqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitTermseq(s)
	}
}

func (p *CLIFParser) Termseq() (localctx ITermseqContext) {
	localctx = NewTermseqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CLIFParserRULE_termseq)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CLIFParserOPEN)|(1<<CLIFParserNUMERAL)|(1<<CLIFParserSEQMARK)|(1<<CLIFParserQUOTEDSTRING)|(1<<CLIFParserENCLOSEDNAME)|(1<<CLIFParserNAMECHARSEQUENCE))) != 0 {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CLIFParserOPEN, CLIFParserNUMERAL, CLIFParserQUOTEDSTRING, CLIFParserENCLOSEDNAME, CLIFParserNAMECHARSEQUENCE:
			{
				p.SetState(40)
				p.Term()
			}

		case CLIFParserSEQMARK:
			{
				p.SetState(41)
				p.Match(CLIFParserSEQMARK)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInterpretednameContext is an interface to support dynamic dispatch.
type IInterpretednameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterpretednameContext differentiates from other interfaces.
	IsInterpretednameContext()
}

type InterpretednameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterpretednameContext() *InterpretednameContext {
	var p = new(InterpretednameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_interpretedname
	return p
}

func (*InterpretednameContext) IsInterpretednameContext() {}

func NewInterpretednameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterpretednameContext {
	var p = new(InterpretednameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_interpretedname

	return p
}

func (s *InterpretednameContext) GetParser() antlr.Parser { return s.parser }

func (s *InterpretednameContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(CLIFParserNUMERAL, 0)
}

func (s *InterpretednameContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(CLIFParserQUOTEDSTRING, 0)
}

func (s *InterpretednameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpretednameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterpretednameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterInterpretedname(s)
	}
}

func (s *InterpretednameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitInterpretedname(s)
	}
}

func (p *CLIFParser) Interpretedname() (localctx IInterpretednameContext) {
	localctx = NewInterpretednameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CLIFParserRULE_interpretedname)
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
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CLIFParserNUMERAL || _la == CLIFParserQUOTEDSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInterpretablenameContext is an interface to support dynamic dispatch.
type IInterpretablenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterpretablenameContext differentiates from other interfaces.
	IsInterpretablenameContext()
}

type InterpretablenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterpretablenameContext() *InterpretablenameContext {
	var p = new(InterpretablenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_interpretablename
	return p
}

func (*InterpretablenameContext) IsInterpretablenameContext() {}

func NewInterpretablenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterpretablenameContext {
	var p = new(InterpretablenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_interpretablename

	return p
}

func (s *InterpretablenameContext) GetParser() antlr.Parser { return s.parser }

func (s *InterpretablenameContext) NAMECHARSEQUENCE() antlr.TerminalNode {
	return s.GetToken(CLIFParserNAMECHARSEQUENCE, 0)
}

func (s *InterpretablenameContext) ENCLOSEDNAME() antlr.TerminalNode {
	return s.GetToken(CLIFParserENCLOSEDNAME, 0)
}

func (s *InterpretablenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpretablenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterpretablenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterInterpretablename(s)
	}
}

func (s *InterpretablenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitInterpretablename(s)
	}
}

func (p *CLIFParser) Interpretablename() (localctx IInterpretablenameContext) {
	localctx = NewInterpretablenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CLIFParserRULE_interpretablename)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == CLIFParserENCLOSEDNAME || _la == CLIFParserNAMECHARSEQUENCE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = CLIFParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) Interpretedname() IInterpretednameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretednameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretednameContext)
}

func (s *NameContext) Interpretablename() IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *CLIFParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CLIFParserRULE_name)

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
	case CLIFParserNUMERAL, CLIFParserQUOTEDSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Interpretedname()
		}

	case CLIFParserENCLOSEDNAME, CLIFParserNAMECHARSEQUENCE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Interpretablename()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = CLIFParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TermContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *TermContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *TermContext) Termseq() ITermseqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermseqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermseqContext)
}

func (s *TermContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *TermContext) CL_COMMENT() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_COMMENT, 0)
}

func (s *TermContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(CLIFParserQUOTEDSTRING, 0)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *CLIFParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CLIFParserRULE_term)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(57)
			p.Operator()
		}
		{
			p.SetState(58)
			p.Termseq()
		}
		{
			p.SetState(59)
			p.Match(CLIFParserCLOSE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(62)
			p.Match(CLIFParserCL_COMMENT)
		}
		{
			p.SetState(63)
			p.Match(CLIFParserQUOTEDSTRING)
		}
		{
			p.SetState(64)
			p.Term()
		}
		{
			p.SetState(65)
			p.Match(CLIFParserCLOSE)
		}

	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *CLIFParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CLIFParserRULE_operator)

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
		p.Term()
	}

	return localctx
}

// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *EquationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CLIFParserEQUAL, 0)
}

func (s *EquationContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *EquationContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *EquationContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *CLIFParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CLIFParserRULE_equation)

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
		p.Match(CLIFParserOPEN)
	}
	{
		p.SetState(72)
		p.Match(CLIFParserEQUAL)
	}
	{
		p.SetState(73)
		p.Term()
	}
	{
		p.SetState(74)
		p.Term()
	}
	{
		p.SetState(75)
		p.Match(CLIFParserCLOSE)
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
	p.RuleIndex = CLIFParserRULE_sentence
	return p
}

func (*SentenceContext) IsSentenceContext() {}

func NewSentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceContext {
	var p = new(SentenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_sentence

	return p
}

func (s *SentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceContext) Atomsent() IAtomsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomsentContext)
}

func (s *SentenceContext) Boolsent() IBoolsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolsentContext)
}

func (s *SentenceContext) Quantsent() IQuantsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantsentContext)
}

func (s *SentenceContext) Commentsent() ICommentsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentsentContext)
}

func (s *SentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterSentence(s)
	}
}

func (s *SentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitSentence(s)
	}
}

func (p *CLIFParser) Sentence() (localctx ISentenceContext) {
	localctx = NewSentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CLIFParserRULE_sentence)

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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Atomsent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Boolsent()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Quantsent()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Commentsent()
		}

	}

	return localctx
}

// IAtomsentContext is an interface to support dynamic dispatch.
type IAtomsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomsentContext differentiates from other interfaces.
	IsAtomsentContext()
}

type AtomsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomsentContext() *AtomsentContext {
	var p = new(AtomsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_atomsent
	return p
}

func (*AtomsentContext) IsAtomsentContext() {}

func NewAtomsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomsentContext {
	var p = new(AtomsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_atomsent

	return p
}

func (s *AtomsentContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomsentContext) Equation() IEquationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *AtomsentContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterAtomsent(s)
	}
}

func (s *AtomsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitAtomsent(s)
	}
}

func (p *CLIFParser) Atomsent() (localctx IAtomsentContext) {
	localctx = NewAtomsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CLIFParserRULE_atomsent)

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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Equation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Atom()
		}

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
	p.RuleIndex = CLIFParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserOPEN)
}

func (s *AtomContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, i)
}

func (s *AtomContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *AtomContext) Termseq() ITermseqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermseqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermseqContext)
}

func (s *AtomContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserCLOSE)
}

func (s *AtomContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, i)
}

func (s *AtomContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *AtomContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *AtomContext) CL_ROLESET() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_ROLESET, 0)
}

func (s *AtomContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *CLIFParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CLIFParserRULE_atom)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(88)
			p.Predicate()
		}
		{
			p.SetState(89)
			p.Termseq()
		}
		{
			p.SetState(90)
			p.Match(CLIFParserCLOSE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(93)
			p.Term()
		}
		{
			p.SetState(94)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(95)
			p.Match(CLIFParserCL_ROLESET)
		}

		{
			p.SetState(96)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(97)
			p.Name()
		}
		{
			p.SetState(98)
			p.Term()
		}
		{
			p.SetState(99)
			p.Match(CLIFParserCLOSE)
		}

		{
			p.SetState(101)
			p.Match(CLIFParserCLOSE)
		}
		{
			p.SetState(102)
			p.Match(CLIFParserCLOSE)
		}

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
	p.RuleIndex = CLIFParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *CLIFParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CLIFParserRULE_predicate)

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
		p.Term()
	}

	return localctx
}

// IBoolsentContext is an interface to support dynamic dispatch.
type IBoolsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolsentContext differentiates from other interfaces.
	IsBoolsentContext()
}

type BoolsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolsentContext() *BoolsentContext {
	var p = new(BoolsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_boolsent
	return p
}

func (*BoolsentContext) IsBoolsentContext() {}

func NewBoolsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolsentContext {
	var p = new(BoolsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_boolsent

	return p
}

func (s *BoolsentContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolsentContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *BoolsentContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *BoolsentContext) AND() antlr.TerminalNode {
	return s.GetToken(CLIFParserAND, 0)
}

func (s *BoolsentContext) OR() antlr.TerminalNode {
	return s.GetToken(CLIFParserOR, 0)
}

func (s *BoolsentContext) AllSentence() []ISentenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISentenceContext)(nil)).Elem())
	var tst = make([]ISentenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISentenceContext)
		}
	}

	return tst
}

func (s *BoolsentContext) Sentence(i int) ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *BoolsentContext) IF() antlr.TerminalNode {
	return s.GetToken(CLIFParserIF, 0)
}

func (s *BoolsentContext) IFF() antlr.TerminalNode {
	return s.GetToken(CLIFParserIFF, 0)
}

func (s *BoolsentContext) NOT() antlr.TerminalNode {
	return s.GetToken(CLIFParserNOT, 0)
}

func (s *BoolsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterBoolsent(s)
	}
}

func (s *BoolsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitBoolsent(s)
	}
}

func (p *CLIFParser) Boolsent() (localctx IBoolsentContext) {
	localctx = NewBoolsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CLIFParserRULE_boolsent)
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

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(109)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CLIFParserAND || _la == CLIFParserOR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CLIFParserOPEN {
			{
				p.SetState(110)
				p.Sentence()
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(CLIFParserCLOSE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(118)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CLIFParserIFF || _la == CLIFParserIF) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(119)
			p.Sentence()
		}
		{
			p.SetState(120)
			p.Sentence()
		}
		{
			p.SetState(121)
			p.Match(CLIFParserCLOSE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(124)
			p.Match(CLIFParserNOT)
		}
		{
			p.SetState(125)
			p.Sentence()
		}
		{
			p.SetState(126)
			p.Match(CLIFParserCLOSE)
		}

	}

	return localctx
}

// IQuantsentContext is an interface to support dynamic dispatch.
type IQuantsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantsentContext differentiates from other interfaces.
	IsQuantsentContext()
}

type QuantsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantsentContext() *QuantsentContext {
	var p = new(QuantsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_quantsent
	return p
}

func (*QuantsentContext) IsQuantsentContext() {}

func NewQuantsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantsentContext {
	var p = new(QuantsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_quantsent

	return p
}

func (s *QuantsentContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantsentContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *QuantsentContext) Boundlist() IBoundlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundlistContext)
}

func (s *QuantsentContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *QuantsentContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *QuantsentContext) FORALL() antlr.TerminalNode {
	return s.GetToken(CLIFParserFORALL, 0)
}

func (s *QuantsentContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(CLIFParserEXISTS, 0)
}

func (s *QuantsentContext) Interpretablename() IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *QuantsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterQuantsent(s)
	}
}

func (s *QuantsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitQuantsent(s)
	}
}

func (p *CLIFParser) Quantsent() (localctx IQuantsentContext) {
	localctx = NewQuantsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CLIFParserRULE_quantsent)
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
		p.SetState(130)
		p.Match(CLIFParserOPEN)
	}
	{
		p.SetState(131)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CLIFParserFORALL || _la == CLIFParserEXISTS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CLIFParserENCLOSEDNAME || _la == CLIFParserNAMECHARSEQUENCE {
		{
			p.SetState(132)
			p.Interpretablename()
		}

	}
	{
		p.SetState(135)
		p.Boundlist()
	}
	{
		p.SetState(136)
		p.Sentence()
	}
	{
		p.SetState(137)
		p.Match(CLIFParserCLOSE)
	}

	return localctx
}

// IBoundlistContext is an interface to support dynamic dispatch.
type IBoundlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundlistContext differentiates from other interfaces.
	IsBoundlistContext()
}

type BoundlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundlistContext() *BoundlistContext {
	var p = new(BoundlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_boundlist
	return p
}

func (*BoundlistContext) IsBoundlistContext() {}

func NewBoundlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundlistContext {
	var p = new(BoundlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_boundlist

	return p
}

func (s *BoundlistContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundlistContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserOPEN)
}

func (s *BoundlistContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, i)
}

func (s *BoundlistContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserCLOSE)
}

func (s *BoundlistContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, i)
}

func (s *BoundlistContext) AllInterpretablename() []IInterpretablenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem())
	var tst = make([]IInterpretablenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInterpretablenameContext)
		}
	}

	return tst
}

func (s *BoundlistContext) Interpretablename(i int) IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *BoundlistContext) AllSEQMARK() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserSEQMARK)
}

func (s *BoundlistContext) SEQMARK(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserSEQMARK, i)
}

func (s *BoundlistContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BoundlistContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BoundlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterBoundlist(s)
	}
}

func (s *BoundlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitBoundlist(s)
	}
}

func (p *CLIFParser) Boundlist() (localctx IBoundlistContext) {
	localctx = NewBoundlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CLIFParserRULE_boundlist)
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
		p.SetState(139)
		p.Match(CLIFParserOPEN)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CLIFParserOPEN)|(1<<CLIFParserSEQMARK)|(1<<CLIFParserENCLOSEDNAME)|(1<<CLIFParserNAMECHARSEQUENCE))) != 0 {
		p.SetState(150)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CLIFParserENCLOSEDNAME, CLIFParserNAMECHARSEQUENCE:
			{
				p.SetState(140)
				p.Interpretablename()
			}

		case CLIFParserSEQMARK:
			{
				p.SetState(141)
				p.Match(CLIFParserSEQMARK)
			}

		case CLIFParserOPEN:
			{
				p.SetState(142)
				p.Match(CLIFParserOPEN)
			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case CLIFParserENCLOSEDNAME, CLIFParserNAMECHARSEQUENCE:
				{
					p.SetState(143)
					p.Interpretablename()
				}

			case CLIFParserSEQMARK:
				{
					p.SetState(144)
					p.Match(CLIFParserSEQMARK)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(147)
				p.Term()
			}
			{
				p.SetState(148)
				p.Match(CLIFParserCLOSE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(CLIFParserCLOSE)
	}

	return localctx
}

// ICommentsentContext is an interface to support dynamic dispatch.
type ICommentsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentsentContext differentiates from other interfaces.
	IsCommentsentContext()
}

type CommentsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentsentContext() *CommentsentContext {
	var p = new(CommentsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_commentsent
	return p
}

func (*CommentsentContext) IsCommentsentContext() {}

func NewCommentsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentsentContext {
	var p = new(CommentsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_commentsent

	return p
}

func (s *CommentsentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentsentContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *CommentsentContext) CL_COMMENT() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_COMMENT, 0)
}

func (s *CommentsentContext) ENCLOSEDNAME() antlr.TerminalNode {
	return s.GetToken(CLIFParserENCLOSEDNAME, 0)
}

func (s *CommentsentContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *CommentsentContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *CommentsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterCommentsent(s)
	}
}

func (s *CommentsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitCommentsent(s)
	}
}

func (p *CLIFParser) Commentsent() (localctx ICommentsentContext) {
	localctx = NewCommentsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CLIFParserRULE_commentsent)

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
		p.Match(CLIFParserOPEN)
	}
	{
		p.SetState(158)
		p.Match(CLIFParserCL_COMMENT)
	}
	{
		p.SetState(159)
		p.Match(CLIFParserENCLOSEDNAME)
	}
	{
		p.SetState(160)
		p.Sentence()
	}
	{
		p.SetState(161)
		p.Match(CLIFParserCLOSE)
	}

	return localctx
}

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
	p.RuleIndex = CLIFParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserOPEN)
}

func (s *ModuleContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, i)
}

func (s *ModuleContext) CL_MODULE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_MODULE, 0)
}

func (s *ModuleContext) Interpretablename() IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *ModuleContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(CLIFParserCLOSE)
}

func (s *ModuleContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, i)
}

func (s *ModuleContext) CL_EXCLUDES() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_EXCLUDES, 0)
}

func (s *ModuleContext) Cltext() ICltextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICltextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICltextContext)
}

func (s *ModuleContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ModuleContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *CLIFParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CLIFParserRULE_module)
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
		p.SetState(163)
		p.Match(CLIFParserOPEN)
	}
	{
		p.SetState(164)
		p.Match(CLIFParserCL_MODULE)
	}
	{
		p.SetState(165)
		p.Interpretablename()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(166)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(167)
			p.Match(CLIFParserCL_EXCLUDES)
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CLIFParserNUMERAL)|(1<<CLIFParserQUOTEDSTRING)|(1<<CLIFParserENCLOSEDNAME)|(1<<CLIFParserNAMECHARSEQUENCE))) != 0 {
			{
				p.SetState(168)
				p.Name()
			}

			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(174)
			p.Match(CLIFParserCLOSE)
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CLIFParserOPEN {
		{
			p.SetState(177)
			p.Cltext()
		}

	}
	{
		p.SetState(180)
		p.Match(CLIFParserCLOSE)
	}

	return localctx
}

// IPhraseContext is an interface to support dynamic dispatch.
type IPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhraseContext differentiates from other interfaces.
	IsPhraseContext()
}

type PhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhraseContext() *PhraseContext {
	var p = new(PhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_phrase
	return p
}

func (*PhraseContext) IsPhraseContext() {}

func NewPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhraseContext {
	var p = new(PhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_phrase

	return p
}

func (s *PhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhraseContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *PhraseContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *PhraseContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *PhraseContext) CL_IMPORTS() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_IMPORTS, 0)
}

func (s *PhraseContext) Interpretablename() IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *PhraseContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *PhraseContext) CL_COMMENT() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_COMMENT, 0)
}

func (s *PhraseContext) ENCLOSEDNAME() antlr.TerminalNode {
	return s.GetToken(CLIFParserENCLOSEDNAME, 0)
}

func (s *PhraseContext) Cltext() ICltextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICltextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICltextContext)
}

func (s *PhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterPhrase(s)
	}
}

func (s *PhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitPhrase(s)
	}
}

func (p *CLIFParser) Phrase() (localctx IPhraseContext) {
	localctx = NewPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CLIFParserRULE_phrase)
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

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Sentence()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Module()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(185)
			p.Match(CLIFParserCL_IMPORTS)
		}
		{
			p.SetState(186)
			p.Interpretablename()
		}
		{
			p.SetState(187)
			p.Match(CLIFParserCLOSE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Match(CLIFParserOPEN)
		}
		{
			p.SetState(190)
			p.Match(CLIFParserCL_COMMENT)
		}
		{
			p.SetState(191)
			p.Match(CLIFParserENCLOSEDNAME)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CLIFParserOPEN {
			{
				p.SetState(192)
				p.Cltext()
			}

		}
		{
			p.SetState(195)
			p.Match(CLIFParserCLOSE)
		}

	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllPhrase() []IPhraseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPhraseContext)(nil)).Elem())
	var tst = make([]IPhraseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPhraseContext)
		}
	}

	return tst
}

func (s *TextContext) Phrase(i int) IPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhraseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPhraseContext)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *CLIFParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CLIFParserRULE_text)
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
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CLIFParserOPEN {
		{
			p.SetState(198)
			p.Phrase()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICltextContext is an interface to support dynamic dispatch.
type ICltextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCltextContext differentiates from other interfaces.
	IsCltextContext()
}

type CltextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCltextContext() *CltextContext {
	var p = new(CltextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_cltext
	return p
}

func (*CltextContext) IsCltextContext() {}

func NewCltextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CltextContext {
	var p = new(CltextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_cltext

	return p
}

func (s *CltextContext) GetParser() antlr.Parser { return s.parser }

func (s *CltextContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *CltextContext) Namedtext() INamedtextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedtextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedtextContext)
}

func (s *CltextContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *CltextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CltextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CltextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterCltext(s)
	}
}

func (s *CltextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitCltext(s)
	}
}

func (p *CLIFParser) Cltext() (localctx ICltextContext) {
	localctx = NewCltextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CLIFParserRULE_cltext)

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

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Module()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Namedtext()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Text()
		}

	}

	return localctx
}

// INamedtextContext is an interface to support dynamic dispatch.
type INamedtextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedtextContext differentiates from other interfaces.
	IsNamedtextContext()
}

type NamedtextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedtextContext() *NamedtextContext {
	var p = new(NamedtextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CLIFParserRULE_namedtext
	return p
}

func (*NamedtextContext) IsNamedtextContext() {}

func NewNamedtextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedtextContext {
	var p = new(NamedtextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CLIFParserRULE_namedtext

	return p
}

func (s *NamedtextContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedtextContext) OPEN() antlr.TerminalNode {
	return s.GetToken(CLIFParserOPEN, 0)
}

func (s *NamedtextContext) CL_TEXT() antlr.TerminalNode {
	return s.GetToken(CLIFParserCL_TEXT, 0)
}

func (s *NamedtextContext) Interpretablename() IInterpretablenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterpretablenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterpretablenameContext)
}

func (s *NamedtextContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CLIFParserCLOSE, 0)
}

func (s *NamedtextContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *NamedtextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedtextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedtextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.EnterNamedtext(s)
	}
}

func (s *NamedtextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CLIFListener); ok {
		listenerT.ExitNamedtext(s)
	}
}

func (p *CLIFParser) Namedtext() (localctx INamedtextContext) {
	localctx = NewNamedtextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CLIFParserRULE_namedtext)
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
		p.SetState(208)
		p.Match(CLIFParserOPEN)
	}
	{
		p.SetState(209)
		p.Match(CLIFParserCL_TEXT)
	}
	{
		p.SetState(210)
		p.Interpretablename()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CLIFParserOPEN {
		{
			p.SetState(211)
			p.Text()
		}

	}
	{
		p.SetState(214)
		p.Match(CLIFParserCLOSE)
	}

	return localctx
}
