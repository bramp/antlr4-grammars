// Generated from STGParser.g4 by ANTLR 4.7.

package stringtemplate // STGParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 133,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	5, 2, 26, 10, 2, 3, 2, 5, 2, 29, 10, 2, 3, 2, 3, 2, 6, 2, 33, 10, 2, 13,
	2, 14, 2, 34, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 6,
	4, 46, 10, 4, 13, 4, 14, 4, 47, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 59, 10, 5, 3, 5, 5, 5, 62, 10, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 69, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6,
	14, 6, 77, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 91, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 102, 10, 9, 12, 9, 14, 9, 105, 11, 9, 3, 9, 3,
	9, 5, 9, 109, 10, 9, 3, 9, 5, 9, 112, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 131, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 2, 3, 3, 2, 10, 12, 2, 145, 2, 25, 3, 2, 2, 2,
	4, 38, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 70, 3, 2,
	2, 2, 12, 78, 3, 2, 2, 2, 14, 92, 3, 2, 2, 2, 16, 111, 3, 2, 2, 2, 18,
	113, 3, 2, 2, 2, 20, 117, 3, 2, 2, 2, 22, 130, 3, 2, 2, 2, 24, 26, 5, 4,
	3, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 29,
	5, 6, 4, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2,
	30, 33, 5, 8, 5, 2, 31, 33, 5, 14, 8, 2, 32, 30, 3, 2, 2, 2, 32, 31, 3,
	2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35,
	36, 3, 2, 2, 2, 36, 37, 7, 2, 2, 3, 37, 3, 3, 2, 2, 2, 38, 39, 7, 27, 2,
	2, 39, 40, 7, 10, 2, 2, 40, 41, 7, 17, 2, 2, 41, 42, 7, 10, 2, 2, 42, 5,
	3, 2, 2, 2, 43, 44, 7, 28, 2, 2, 44, 46, 7, 10, 2, 2, 45, 43, 3, 2, 2,
	2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 7, 3,
	2, 2, 2, 49, 50, 7, 23, 2, 2, 50, 51, 7, 9, 2, 2, 51, 52, 7, 16, 2, 2,
	52, 53, 7, 9, 2, 2, 53, 54, 7, 19, 2, 2, 54, 62, 7, 20, 2, 2, 55, 56, 7,
	9, 2, 2, 56, 58, 7, 19, 2, 2, 57, 59, 5, 10, 6, 2, 58, 57, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 7, 20, 2, 2, 61, 49, 3,
	2, 2, 2, 61, 55, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 14, 2, 2, 64,
	69, 9, 2, 2, 2, 65, 66, 7, 9, 2, 2, 66, 67, 7, 14, 2, 2, 67, 69, 7, 9,
	2, 2, 68, 61, 3, 2, 2, 2, 68, 65, 3, 2, 2, 2, 69, 9, 3, 2, 2, 2, 70, 75,
	5, 12, 7, 2, 71, 72, 7, 17, 2, 2, 72, 74, 5, 12, 7, 2, 73, 71, 3, 2, 2,
	2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 11,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 90, 7, 9, 2, 2, 79, 80, 7, 15, 2, 2,
	80, 91, 7, 10, 2, 2, 81, 82, 7, 15, 2, 2, 82, 91, 7, 13, 2, 2, 83, 84,
	7, 15, 2, 2, 84, 91, 7, 24, 2, 2, 85, 86, 7, 15, 2, 2, 86, 91, 7, 25, 2,
	2, 87, 88, 7, 15, 2, 2, 88, 89, 7, 21, 2, 2, 89, 91, 7, 22, 2, 2, 90, 79,
	3, 2, 2, 2, 90, 81, 3, 2, 2, 2, 90, 83, 3, 2, 2, 2, 90, 85, 3, 2, 2, 2,
	90, 87, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 13, 3, 2, 2, 2, 92, 93, 7,
	9, 2, 2, 93, 94, 7, 14, 2, 2, 94, 95, 7, 21, 2, 2, 95, 96, 5, 16, 9, 2,
	96, 97, 7, 22, 2, 2, 97, 15, 3, 2, 2, 2, 98, 103, 5, 18, 10, 2, 99, 100,
	7, 17, 2, 2, 100, 102, 5, 18, 10, 2, 101, 99, 3, 2, 2, 2, 102, 105, 3,
	2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 108, 3, 2, 2,
	2, 105, 103, 3, 2, 2, 2, 106, 107, 7, 17, 2, 2, 107, 109, 5, 20, 11, 2,
	108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110,
	112, 5, 20, 11, 2, 111, 98, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 17,
	3, 2, 2, 2, 113, 114, 7, 10, 2, 2, 114, 115, 7, 18, 2, 2, 115, 116, 5,
	22, 12, 2, 116, 19, 3, 2, 2, 2, 117, 118, 7, 29, 2, 2, 118, 119, 7, 18,
	2, 2, 119, 120, 5, 22, 12, 2, 120, 21, 3, 2, 2, 2, 121, 131, 7, 11, 2,
	2, 122, 131, 7, 12, 2, 2, 123, 131, 7, 13, 2, 2, 124, 131, 7, 10, 2, 2,
	125, 131, 7, 24, 2, 2, 126, 131, 7, 25, 2, 2, 127, 128, 7, 21, 2, 2, 128,
	131, 7, 22, 2, 2, 129, 131, 7, 30, 2, 2, 130, 121, 3, 2, 2, 2, 130, 122,
	3, 2, 2, 2, 130, 123, 3, 2, 2, 2, 130, 124, 3, 2, 2, 2, 130, 125, 3, 2,
	2, 2, 130, 126, 3, 2, 2, 2, 130, 127, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2,
	131, 23, 3, 2, 2, 2, 16, 25, 28, 32, 34, 47, 58, 61, 68, 75, 90, 103, 108,
	111, 130,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "'delimiters'", "'import'", "'default'", "'key'",
	"'value'", "'first'", "'last'", "'rest'", "'trunc'", "'strip'", "'trim'",
	"'length'", "'strlen'", "'reverse'", "'group'", "'wrap'", "'anchor'", "'separator'",
}
var symbolicNames = []string{
	"", "DOC_COMMENT", "BLOCK_COMMENT", "LINE_COMMENT", "TMPL_COMMENT", "HORZ_WS",
	"VERT_WS", "ID", "STRING", "BIGSTRING", "BIGSTRING_NO_NL", "ANON_TEMPLATE",
	"TMPL_ASSIGN", "ASSIGN", "DOT", "COMMA", "COLON", "LPAREN", "RPAREN", "LBRACK",
	"RBRACK", "AT", "TRUE", "FALSE", "ELLIPSIS", "DELIMITERS", "IMPORT", "DEFAULT",
	"KEY", "VALUE", "FIRST", "LAST", "REST", "TRUNC", "STRIP", "TRIM", "LENGTH",
	"STRLEN", "REVERSE", "GROUP", "WRAP", "ANCHOR", "SEPARATOR",
}

var ruleNames = []string{
	"group", "delimiters", "imports", "template", "formalArgs", "formalArg",
	"dict", "dictPairs", "keyValuePair", "defaultValuePair", "keyValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type STGParser struct {
	*antlr.BaseParser
}

func NewSTGParser(input antlr.TokenStream) *STGParser {
	this := new(STGParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "STGParser.g4"

	return this
}

// STGParser tokens.
const (
	STGParserEOF             = antlr.TokenEOF
	STGParserDOC_COMMENT     = 1
	STGParserBLOCK_COMMENT   = 2
	STGParserLINE_COMMENT    = 3
	STGParserTMPL_COMMENT    = 4
	STGParserHORZ_WS         = 5
	STGParserVERT_WS         = 6
	STGParserID              = 7
	STGParserSTRING          = 8
	STGParserBIGSTRING       = 9
	STGParserBIGSTRING_NO_NL = 10
	STGParserANON_TEMPLATE   = 11
	STGParserTMPL_ASSIGN     = 12
	STGParserASSIGN          = 13
	STGParserDOT             = 14
	STGParserCOMMA           = 15
	STGParserCOLON           = 16
	STGParserLPAREN          = 17
	STGParserRPAREN          = 18
	STGParserLBRACK          = 19
	STGParserRBRACK          = 20
	STGParserAT              = 21
	STGParserTRUE            = 22
	STGParserFALSE           = 23
	STGParserELLIPSIS        = 24
	STGParserDELIMITERS      = 25
	STGParserIMPORT          = 26
	STGParserDEFAULT         = 27
	STGParserKEY             = 28
	STGParserVALUE           = 29
	STGParserFIRST           = 30
	STGParserLAST            = 31
	STGParserREST            = 32
	STGParserTRUNC           = 33
	STGParserSTRIP           = 34
	STGParserTRIM            = 35
	STGParserLENGTH          = 36
	STGParserSTRLEN          = 37
	STGParserREVERSE         = 38
	STGParserGROUP           = 39
	STGParserWRAP            = 40
	STGParserANCHOR          = 41
	STGParserSEPARATOR       = 42
)

// STGParser rules.
const (
	STGParserRULE_group            = 0
	STGParserRULE_delimiters       = 1
	STGParserRULE_imports          = 2
	STGParserRULE_template         = 3
	STGParserRULE_formalArgs       = 4
	STGParserRULE_formalArg        = 5
	STGParserRULE_dict             = 6
	STGParserRULE_dictPairs        = 7
	STGParserRULE_keyValuePair     = 8
	STGParserRULE_defaultValuePair = 9
	STGParserRULE_keyValue         = 10
)

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) EOF() antlr.TerminalNode {
	return s.GetToken(STGParserEOF, 0)
}

func (s *GroupContext) Delimiters() IDelimitersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelimitersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelimitersContext)
}

func (s *GroupContext) Imports() IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *GroupContext) AllTemplate() []ITemplateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITemplateContext)(nil)).Elem())
	var tst = make([]ITemplateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITemplateContext)
		}
	}

	return tst
}

func (s *GroupContext) Template(i int) ITemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITemplateContext)
}

func (s *GroupContext) AllDict() []IDictContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDictContext)(nil)).Elem())
	var tst = make([]IDictContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDictContext)
		}
	}

	return tst
}

func (s *GroupContext) Dict(i int) IDictContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDictContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STGParserRULE_group)
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STGParserDELIMITERS {
		{
			p.SetState(22)
			p.Delimiters()
		}

	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STGParserIMPORT {
		{
			p.SetState(25)
			p.Imports()
		}

	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STGParserID || _la == STGParserAT {
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(28)
				p.Template()
			}

		case 2:
			{
				p.SetState(29)
				p.Dict()
			}

		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(STGParserEOF)
	}

	return localctx
}

// IDelimitersContext is an interface to support dynamic dispatch.
type IDelimitersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelimitersContext differentiates from other interfaces.
	IsDelimitersContext()
}

type DelimitersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimitersContext() *DelimitersContext {
	var p = new(DelimitersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_delimiters
	return p
}

func (*DelimitersContext) IsDelimitersContext() {}

func NewDelimitersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimitersContext {
	var p = new(DelimitersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_delimiters

	return p
}

func (s *DelimitersContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimitersContext) DELIMITERS() antlr.TerminalNode {
	return s.GetToken(STGParserDELIMITERS, 0)
}

func (s *DelimitersContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(STGParserSTRING)
}

func (s *DelimitersContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, i)
}

func (s *DelimitersContext) COMMA() antlr.TerminalNode {
	return s.GetToken(STGParserCOMMA, 0)
}

func (s *DelimitersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimitersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimitersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterDelimiters(s)
	}
}

func (s *DelimitersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitDelimiters(s)
	}
}

func (s *DelimitersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitDelimiters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) Delimiters() (localctx IDelimitersContext) {
	localctx = NewDelimitersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STGParserRULE_delimiters)

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
		p.SetState(36)
		p.Match(STGParserDELIMITERS)
	}
	{
		p.SetState(37)
		p.Match(STGParserSTRING)
	}
	{
		p.SetState(38)
		p.Match(STGParserCOMMA)
	}
	{
		p.SetState(39)
		p.Match(STGParserSTRING)
	}

	return localctx
}

// IImportsContext is an interface to support dynamic dispatch.
type IImportsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportsContext differentiates from other interfaces.
	IsImportsContext()
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportsContext) AllIMPORT() []antlr.TerminalNode {
	return s.GetTokens(STGParserIMPORT)
}

func (s *ImportsContext) IMPORT(i int) antlr.TerminalNode {
	return s.GetToken(STGParserIMPORT, i)
}

func (s *ImportsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(STGParserSTRING)
}

func (s *ImportsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, i)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterImports(s)
	}
}

func (s *ImportsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitImports(s)
	}
}

func (s *ImportsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitImports(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, STGParserRULE_imports)
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STGParserIMPORT {
		{
			p.SetState(41)
			p.Match(STGParserIMPORT)
		}
		{
			p.SetState(42)
			p.Match(STGParserSTRING)
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_template
	return p
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) TMPL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(STGParserTMPL_ASSIGN, 0)
}

func (s *TemplateContext) STRING() antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, 0)
}

func (s *TemplateContext) BIGSTRING() antlr.TerminalNode {
	return s.GetToken(STGParserBIGSTRING, 0)
}

func (s *TemplateContext) BIGSTRING_NO_NL() antlr.TerminalNode {
	return s.GetToken(STGParserBIGSTRING_NO_NL, 0)
}

func (s *TemplateContext) AT() antlr.TerminalNode {
	return s.GetToken(STGParserAT, 0)
}

func (s *TemplateContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(STGParserID)
}

func (s *TemplateContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(STGParserID, i)
}

func (s *TemplateContext) DOT() antlr.TerminalNode {
	return s.GetToken(STGParserDOT, 0)
}

func (s *TemplateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STGParserLPAREN, 0)
}

func (s *TemplateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STGParserRPAREN, 0)
}

func (s *TemplateContext) FormalArgs() IFormalArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalArgsContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STGParserRULE_template)
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

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case STGParserAT:
			{
				p.SetState(47)
				p.Match(STGParserAT)
			}
			{
				p.SetState(48)
				p.Match(STGParserID)
			}
			{
				p.SetState(49)
				p.Match(STGParserDOT)
			}
			{
				p.SetState(50)
				p.Match(STGParserID)
			}
			{
				p.SetState(51)
				p.Match(STGParserLPAREN)
			}
			{
				p.SetState(52)
				p.Match(STGParserRPAREN)
			}

		case STGParserID:
			{
				p.SetState(53)
				p.Match(STGParserID)
			}
			{
				p.SetState(54)
				p.Match(STGParserLPAREN)
			}
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == STGParserID {
				{
					p.SetState(55)
					p.FormalArgs()
				}

			}
			{
				p.SetState(58)
				p.Match(STGParserRPAREN)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(61)
			p.Match(STGParserTMPL_ASSIGN)
		}
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<STGParserSTRING)|(1<<STGParserBIGSTRING)|(1<<STGParserBIGSTRING_NO_NL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(STGParserID)
		}
		{
			p.SetState(64)
			p.Match(STGParserTMPL_ASSIGN)
		}
		{
			p.SetState(65)
			p.Match(STGParserID)
		}

	}

	return localctx
}

// IFormalArgsContext is an interface to support dynamic dispatch.
type IFormalArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalArgsContext differentiates from other interfaces.
	IsFormalArgsContext()
}

type FormalArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalArgsContext() *FormalArgsContext {
	var p = new(FormalArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_formalArgs
	return p
}

func (*FormalArgsContext) IsFormalArgsContext() {}

func NewFormalArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalArgsContext {
	var p = new(FormalArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_formalArgs

	return p
}

func (s *FormalArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalArgsContext) AllFormalArg() []IFormalArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalArgContext)(nil)).Elem())
	var tst = make([]IFormalArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalArgContext)
		}
	}

	return tst
}

func (s *FormalArgsContext) FormalArg(i int) IFormalArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalArgContext)
}

func (s *FormalArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STGParserCOMMA)
}

func (s *FormalArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STGParserCOMMA, i)
}

func (s *FormalArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterFormalArgs(s)
	}
}

func (s *FormalArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitFormalArgs(s)
	}
}

func (s *FormalArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitFormalArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) FormalArgs() (localctx IFormalArgsContext) {
	localctx = NewFormalArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, STGParserRULE_formalArgs)
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
		p.SetState(68)
		p.FormalArg()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STGParserCOMMA {
		{
			p.SetState(69)
			p.Match(STGParserCOMMA)
		}
		{
			p.SetState(70)
			p.FormalArg()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormalArgContext is an interface to support dynamic dispatch.
type IFormalArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalArgContext differentiates from other interfaces.
	IsFormalArgContext()
}

type FormalArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalArgContext() *FormalArgContext {
	var p = new(FormalArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_formalArg
	return p
}

func (*FormalArgContext) IsFormalArgContext() {}

func NewFormalArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalArgContext {
	var p = new(FormalArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_formalArg

	return p
}

func (s *FormalArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalArgContext) ID() antlr.TerminalNode {
	return s.GetToken(STGParserID, 0)
}

func (s *FormalArgContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(STGParserASSIGN, 0)
}

func (s *FormalArgContext) STRING() antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, 0)
}

func (s *FormalArgContext) ANON_TEMPLATE() antlr.TerminalNode {
	return s.GetToken(STGParserANON_TEMPLATE, 0)
}

func (s *FormalArgContext) TRUE() antlr.TerminalNode {
	return s.GetToken(STGParserTRUE, 0)
}

func (s *FormalArgContext) FALSE() antlr.TerminalNode {
	return s.GetToken(STGParserFALSE, 0)
}

func (s *FormalArgContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserLBRACK, 0)
}

func (s *FormalArgContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserRBRACK, 0)
}

func (s *FormalArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterFormalArg(s)
	}
}

func (s *FormalArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitFormalArg(s)
	}
}

func (s *FormalArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitFormalArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) FormalArg() (localctx IFormalArgContext) {
	localctx = NewFormalArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, STGParserRULE_formalArg)

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
		p.Match(STGParserID)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(77)
			p.Match(STGParserASSIGN)
		}
		{
			p.SetState(78)
			p.Match(STGParserSTRING)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(79)
			p.Match(STGParserASSIGN)
		}
		{
			p.SetState(80)
			p.Match(STGParserANON_TEMPLATE)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(81)
			p.Match(STGParserASSIGN)
		}
		{
			p.SetState(82)
			p.Match(STGParserTRUE)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(83)
			p.Match(STGParserASSIGN)
		}
		{
			p.SetState(84)
			p.Match(STGParserFALSE)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 5 {
		{
			p.SetState(85)
			p.Match(STGParserASSIGN)
		}
		{
			p.SetState(86)
			p.Match(STGParserLBRACK)
		}
		{
			p.SetState(87)
			p.Match(STGParserRBRACK)
		}

	}

	return localctx
}

// IDictContext is an interface to support dynamic dispatch.
type IDictContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictContext differentiates from other interfaces.
	IsDictContext()
}

type DictContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictContext() *DictContext {
	var p = new(DictContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_dict
	return p
}

func (*DictContext) IsDictContext() {}

func NewDictContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictContext {
	var p = new(DictContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_dict

	return p
}

func (s *DictContext) GetParser() antlr.Parser { return s.parser }

func (s *DictContext) ID() antlr.TerminalNode {
	return s.GetToken(STGParserID, 0)
}

func (s *DictContext) TMPL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(STGParserTMPL_ASSIGN, 0)
}

func (s *DictContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserLBRACK, 0)
}

func (s *DictContext) DictPairs() IDictPairsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictPairsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictPairsContext)
}

func (s *DictContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserRBRACK, 0)
}

func (s *DictContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterDict(s)
	}
}

func (s *DictContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitDict(s)
	}
}

func (s *DictContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitDict(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) Dict() (localctx IDictContext) {
	localctx = NewDictContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, STGParserRULE_dict)

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
		p.SetState(90)
		p.Match(STGParserID)
	}
	{
		p.SetState(91)
		p.Match(STGParserTMPL_ASSIGN)
	}
	{
		p.SetState(92)
		p.Match(STGParserLBRACK)
	}
	{
		p.SetState(93)
		p.DictPairs()
	}
	{
		p.SetState(94)
		p.Match(STGParserRBRACK)
	}

	return localctx
}

// IDictPairsContext is an interface to support dynamic dispatch.
type IDictPairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictPairsContext differentiates from other interfaces.
	IsDictPairsContext()
}

type DictPairsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictPairsContext() *DictPairsContext {
	var p = new(DictPairsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_dictPairs
	return p
}

func (*DictPairsContext) IsDictPairsContext() {}

func NewDictPairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictPairsContext {
	var p = new(DictPairsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_dictPairs

	return p
}

func (s *DictPairsContext) GetParser() antlr.Parser { return s.parser }

func (s *DictPairsContext) AllKeyValuePair() []IKeyValuePairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValuePairContext)(nil)).Elem())
	var tst = make([]IKeyValuePairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValuePairContext)
		}
	}

	return tst
}

func (s *DictPairsContext) KeyValuePair(i int) IKeyValuePairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValuePairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValuePairContext)
}

func (s *DictPairsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STGParserCOMMA)
}

func (s *DictPairsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STGParserCOMMA, i)
}

func (s *DictPairsContext) DefaultValuePair() IDefaultValuePairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValuePairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValuePairContext)
}

func (s *DictPairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictPairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictPairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterDictPairs(s)
	}
}

func (s *DictPairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitDictPairs(s)
	}
}

func (s *DictPairsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitDictPairs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) DictPairs() (localctx IDictPairsContext) {
	localctx = NewDictPairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, STGParserRULE_dictPairs)
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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STGParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.KeyValuePair()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(97)
					p.Match(STGParserCOMMA)
				}
				{
					p.SetState(98)
					p.KeyValuePair()
				}

			}
			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STGParserCOMMA {
			{
				p.SetState(104)
				p.Match(STGParserCOMMA)
			}
			{
				p.SetState(105)
				p.DefaultValuePair()
			}

		}

	case STGParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.DefaultValuePair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeyValuePairContext is an interface to support dynamic dispatch.
type IKeyValuePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValuePairContext differentiates from other interfaces.
	IsKeyValuePairContext()
}

type KeyValuePairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValuePairContext() *KeyValuePairContext {
	var p = new(KeyValuePairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_keyValuePair
	return p
}

func (*KeyValuePairContext) IsKeyValuePairContext() {}

func NewKeyValuePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValuePairContext {
	var p = new(KeyValuePairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_keyValuePair

	return p
}

func (s *KeyValuePairContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValuePairContext) STRING() antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, 0)
}

func (s *KeyValuePairContext) COLON() antlr.TerminalNode {
	return s.GetToken(STGParserCOLON, 0)
}

func (s *KeyValuePairContext) KeyValue() IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *KeyValuePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValuePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValuePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterKeyValuePair(s)
	}
}

func (s *KeyValuePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitKeyValuePair(s)
	}
}

func (s *KeyValuePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitKeyValuePair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) KeyValuePair() (localctx IKeyValuePairContext) {
	localctx = NewKeyValuePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, STGParserRULE_keyValuePair)

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
		p.Match(STGParserSTRING)
	}
	{
		p.SetState(112)
		p.Match(STGParserCOLON)
	}
	{
		p.SetState(113)
		p.KeyValue()
	}

	return localctx
}

// IDefaultValuePairContext is an interface to support dynamic dispatch.
type IDefaultValuePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValuePairContext differentiates from other interfaces.
	IsDefaultValuePairContext()
}

type DefaultValuePairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValuePairContext() *DefaultValuePairContext {
	var p = new(DefaultValuePairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_defaultValuePair
	return p
}

func (*DefaultValuePairContext) IsDefaultValuePairContext() {}

func NewDefaultValuePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValuePairContext {
	var p = new(DefaultValuePairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_defaultValuePair

	return p
}

func (s *DefaultValuePairContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValuePairContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(STGParserDEFAULT, 0)
}

func (s *DefaultValuePairContext) COLON() antlr.TerminalNode {
	return s.GetToken(STGParserCOLON, 0)
}

func (s *DefaultValuePairContext) KeyValue() IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *DefaultValuePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValuePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValuePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterDefaultValuePair(s)
	}
}

func (s *DefaultValuePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitDefaultValuePair(s)
	}
}

func (s *DefaultValuePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitDefaultValuePair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) DefaultValuePair() (localctx IDefaultValuePairContext) {
	localctx = NewDefaultValuePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, STGParserRULE_defaultValuePair)

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
		p.SetState(115)
		p.Match(STGParserDEFAULT)
	}
	{
		p.SetState(116)
		p.Match(STGParserCOLON)
	}
	{
		p.SetState(117)
		p.KeyValue()
	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STGParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STGParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) BIGSTRING() antlr.TerminalNode {
	return s.GetToken(STGParserBIGSTRING, 0)
}

func (s *KeyValueContext) BIGSTRING_NO_NL() antlr.TerminalNode {
	return s.GetToken(STGParserBIGSTRING_NO_NL, 0)
}

func (s *KeyValueContext) ANON_TEMPLATE() antlr.TerminalNode {
	return s.GetToken(STGParserANON_TEMPLATE, 0)
}

func (s *KeyValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(STGParserSTRING, 0)
}

func (s *KeyValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(STGParserTRUE, 0)
}

func (s *KeyValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(STGParserFALSE, 0)
}

func (s *KeyValueContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserLBRACK, 0)
}

func (s *KeyValueContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STGParserRBRACK, 0)
}

func (s *KeyValueContext) KEY() antlr.TerminalNode {
	return s.GetToken(STGParserKEY, 0)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STGParserListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STGParserVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STGParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STGParserRULE_keyValue)

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

	switch p.GetTokenStream().LA(1) {
	case STGParserBIGSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(STGParserBIGSTRING)
		}

	case STGParserBIGSTRING_NO_NL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(STGParserBIGSTRING_NO_NL)
		}

	case STGParserANON_TEMPLATE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Match(STGParserANON_TEMPLATE)
		}

	case STGParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Match(STGParserSTRING)
		}

	case STGParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Match(STGParserTRUE)
		}

	case STGParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Match(STGParserFALSE)
		}

	case STGParserLBRACK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(125)
			p.Match(STGParserLBRACK)
		}
		{
			p.SetState(126)
			p.Match(STGParserRBRACK)
		}

	case STGParserKEY:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(127)
			p.Match(STGParserKEY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
