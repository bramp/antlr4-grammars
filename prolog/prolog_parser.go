// Code generated from prolog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prolog // prolog
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 68, 106,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 7, 2, 21, 10, 2, 12, 2, 14, 2, 24,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 7, 5, 38, 10, 5, 12, 5, 14, 5, 41, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 50, 10, 6, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 69, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 78, 10, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 84, 10, 6, 12, 6, 14, 6, 87, 11, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 102, 10, 8, 3, 9, 3, 9, 3, 9, 2, 3, 10, 10, 2, 4, 6, 8, 10, 12, 14,
	16, 2, 4, 6, 2, 3, 3, 5, 5, 8, 8, 14, 52, 3, 2, 56, 60, 2, 120, 2, 22,
	3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10,
	77, 3, 2, 2, 2, 12, 88, 3, 2, 2, 2, 14, 101, 3, 2, 2, 2, 16, 103, 3, 2,
	2, 2, 18, 21, 5, 4, 3, 2, 19, 21, 5, 6, 4, 2, 20, 18, 3, 2, 2, 2, 20, 19,
	3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2,
	23, 25, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2,
	2, 2, 27, 28, 7, 3, 2, 2, 28, 29, 5, 10, 6, 2, 29, 30, 7, 4, 2, 2, 30,
	5, 3, 2, 2, 2, 31, 32, 5, 10, 6, 2, 32, 33, 7, 4, 2, 2, 33, 7, 3, 2, 2,
	2, 34, 39, 5, 10, 6, 2, 35, 36, 7, 5, 2, 2, 36, 38, 5, 10, 6, 2, 37, 35,
	3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 9, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 8, 6, 1, 2, 43, 78, 7, 55,
	2, 2, 44, 45, 7, 6, 2, 2, 45, 46, 5, 10, 6, 2, 46, 47, 7, 7, 2, 2, 47,
	78, 3, 2, 2, 2, 48, 50, 7, 8, 2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 78, 5, 16, 9, 2, 52, 54, 7, 8, 2, 2, 53, 52,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 78, 7, 61, 2, 2,
	56, 57, 5, 14, 8, 2, 57, 58, 7, 6, 2, 2, 58, 59, 5, 8, 5, 2, 59, 60, 7,
	7, 2, 2, 60, 78, 3, 2, 2, 2, 61, 62, 5, 12, 7, 2, 62, 63, 5, 10, 6, 6,
	63, 78, 3, 2, 2, 2, 64, 65, 7, 9, 2, 2, 65, 68, 5, 8, 5, 2, 66, 67, 7,
	10, 2, 2, 67, 69, 5, 10, 6, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 71, 7, 11, 2, 2, 71, 78, 3, 2, 2, 2, 72, 73, 7,
	12, 2, 2, 73, 74, 5, 8, 5, 2, 74, 75, 7, 13, 2, 2, 75, 78, 3, 2, 2, 2,
	76, 78, 5, 14, 8, 2, 77, 42, 3, 2, 2, 2, 77, 44, 3, 2, 2, 2, 77, 49, 3,
	2, 2, 2, 77, 53, 3, 2, 2, 2, 77, 56, 3, 2, 2, 2, 77, 61, 3, 2, 2, 2, 77,
	64, 3, 2, 2, 2, 77, 72, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 85, 3, 2, 2,
	2, 79, 80, 12, 7, 2, 2, 80, 81, 5, 12, 7, 2, 81, 82, 5, 10, 6, 7, 82, 84,
	3, 2, 2, 2, 83, 79, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2,
	85, 86, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 9,
	2, 2, 2, 89, 13, 3, 2, 2, 2, 90, 91, 7, 9, 2, 2, 91, 102, 7, 11, 2, 2,
	92, 93, 7, 12, 2, 2, 93, 102, 7, 13, 2, 2, 94, 102, 7, 54, 2, 2, 95, 102,
	7, 62, 2, 2, 96, 102, 7, 63, 2, 2, 97, 102, 7, 64, 2, 2, 98, 102, 7, 65,
	2, 2, 99, 102, 7, 20, 2, 2, 100, 102, 7, 53, 2, 2, 101, 90, 3, 2, 2, 2,
	101, 92, 3, 2, 2, 2, 101, 94, 3, 2, 2, 2, 101, 95, 3, 2, 2, 2, 101, 96,
	3, 2, 2, 2, 101, 97, 3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 101, 100, 3, 2, 2, 2, 102, 15, 3, 2, 2, 2, 103, 104, 9, 3, 2, 2, 104,
	17, 3, 2, 2, 2, 11, 20, 22, 39, 49, 53, 68, 77, 85, 101,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':-'", "'.'", "','", "'('", "')'", "'-'", "'['", "'|'", "']'", "'{'",
	"'}'", "'-->'", "'?-'", "'dynamic'", "'multifile'", "'discontiguous'",
	"'public'", "';'", "'->'", "'\\+'", "'='", "'\\='", "'=='", "'\\=='", "'@<'",
	"'@=<'", "'@>'", "'@>='", "'=..'", "'is'", "'=:='", "'=\\='", "'<'", "'=<'",
	"'>'", "'>='", "':'", "'+'", "'/\\'", "'\\/'", "'*'", "'/'", "'//'", "'rem'",
	"'mod'", "'<<'", "'>>'", "'**'", "'^'", "'\\'", "'!'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LETTER_DIGIT",
	"VARIABLE", "DECIMAL", "BINARY", "OCTAL", "HEX", "CHARACTER_CODE_CONSTANT",
	"FLOAT", "GRAPHIC_TOKEN", "QUOTED", "DOUBLE_QUOTED_LIST", "BACK_QUOTED_STRING",
	"WS", "COMMENT", "MULTILINE_COMMENT",
}

var ruleNames = []string{
	"p_text", "directive", "clause", "termlist", "term", "operator", "atom",
	"integer",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type prologParser struct {
	*antlr.BaseParser
}

func NewprologParser(input antlr.TokenStream) *prologParser {
	this := new(prologParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "prolog.g4"

	return this
}

// prologParser tokens.
const (
	prologParserEOF                     = antlr.TokenEOF
	prologParserT__0                    = 1
	prologParserT__1                    = 2
	prologParserT__2                    = 3
	prologParserT__3                    = 4
	prologParserT__4                    = 5
	prologParserT__5                    = 6
	prologParserT__6                    = 7
	prologParserT__7                    = 8
	prologParserT__8                    = 9
	prologParserT__9                    = 10
	prologParserT__10                   = 11
	prologParserT__11                   = 12
	prologParserT__12                   = 13
	prologParserT__13                   = 14
	prologParserT__14                   = 15
	prologParserT__15                   = 16
	prologParserT__16                   = 17
	prologParserT__17                   = 18
	prologParserT__18                   = 19
	prologParserT__19                   = 20
	prologParserT__20                   = 21
	prologParserT__21                   = 22
	prologParserT__22                   = 23
	prologParserT__23                   = 24
	prologParserT__24                   = 25
	prologParserT__25                   = 26
	prologParserT__26                   = 27
	prologParserT__27                   = 28
	prologParserT__28                   = 29
	prologParserT__29                   = 30
	prologParserT__30                   = 31
	prologParserT__31                   = 32
	prologParserT__32                   = 33
	prologParserT__33                   = 34
	prologParserT__34                   = 35
	prologParserT__35                   = 36
	prologParserT__36                   = 37
	prologParserT__37                   = 38
	prologParserT__38                   = 39
	prologParserT__39                   = 40
	prologParserT__40                   = 41
	prologParserT__41                   = 42
	prologParserT__42                   = 43
	prologParserT__43                   = 44
	prologParserT__44                   = 45
	prologParserT__45                   = 46
	prologParserT__46                   = 47
	prologParserT__47                   = 48
	prologParserT__48                   = 49
	prologParserT__49                   = 50
	prologParserT__50                   = 51
	prologParserLETTER_DIGIT            = 52
	prologParserVARIABLE                = 53
	prologParserDECIMAL                 = 54
	prologParserBINARY                  = 55
	prologParserOCTAL                   = 56
	prologParserHEX                     = 57
	prologParserCHARACTER_CODE_CONSTANT = 58
	prologParserFLOAT                   = 59
	prologParserGRAPHIC_TOKEN           = 60
	prologParserQUOTED                  = 61
	prologParserDOUBLE_QUOTED_LIST      = 62
	prologParserBACK_QUOTED_STRING      = 63
	prologParserWS                      = 64
	prologParserCOMMENT                 = 65
	prologParserMULTILINE_COMMENT       = 66
)

// prologParser rules.
const (
	prologParserRULE_p_text    = 0
	prologParserRULE_directive = 1
	prologParserRULE_clause    = 2
	prologParserRULE_termlist  = 3
	prologParserRULE_term      = 4
	prologParserRULE_operator  = 5
	prologParserRULE_atom      = 6
	prologParserRULE_integer   = 7
)

// IP_textContext is an interface to support dynamic dispatch.
type IP_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsP_textContext differentiates from other interfaces.
	IsP_textContext()
}

type P_textContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyP_textContext() *P_textContext {
	var p = new(P_textContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = prologParserRULE_p_text
	return p
}

func (*P_textContext) IsP_textContext() {}

func NewP_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *P_textContext {
	var p = new(P_textContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_p_text

	return p
}

func (s *P_textContext) GetParser() antlr.Parser { return s.parser }

func (s *P_textContext) EOF() antlr.TerminalNode {
	return s.GetToken(prologParserEOF, 0)
}

func (s *P_textContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *P_textContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *P_textContext) AllClause() []IClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseContext)(nil)).Elem())
	var tst = make([]IClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseContext)
		}
	}

	return tst
}

func (s *P_textContext) Clause(i int) IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *P_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *P_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *P_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterP_text(s)
	}
}

func (s *P_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitP_text(s)
	}
}

func (p *prologParser) P_text() (localctx IP_textContext) {
	localctx = NewP_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, prologParserRULE_p_text)
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<prologParserT__0)|(1<<prologParserT__2)|(1<<prologParserT__3)|(1<<prologParserT__5)|(1<<prologParserT__6)|(1<<prologParserT__9)|(1<<prologParserT__11)|(1<<prologParserT__12)|(1<<prologParserT__13)|(1<<prologParserT__14)|(1<<prologParserT__15)|(1<<prologParserT__16)|(1<<prologParserT__17)|(1<<prologParserT__18)|(1<<prologParserT__19)|(1<<prologParserT__20)|(1<<prologParserT__21)|(1<<prologParserT__22)|(1<<prologParserT__23)|(1<<prologParserT__24)|(1<<prologParserT__25)|(1<<prologParserT__26)|(1<<prologParserT__27)|(1<<prologParserT__28)|(1<<prologParserT__29)|(1<<prologParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(prologParserT__31-32))|(1<<(prologParserT__32-32))|(1<<(prologParserT__33-32))|(1<<(prologParserT__34-32))|(1<<(prologParserT__35-32))|(1<<(prologParserT__36-32))|(1<<(prologParserT__37-32))|(1<<(prologParserT__38-32))|(1<<(prologParserT__39-32))|(1<<(prologParserT__40-32))|(1<<(prologParserT__41-32))|(1<<(prologParserT__42-32))|(1<<(prologParserT__43-32))|(1<<(prologParserT__44-32))|(1<<(prologParserT__45-32))|(1<<(prologParserT__46-32))|(1<<(prologParserT__47-32))|(1<<(prologParserT__48-32))|(1<<(prologParserT__49-32))|(1<<(prologParserT__50-32))|(1<<(prologParserLETTER_DIGIT-32))|(1<<(prologParserVARIABLE-32))|(1<<(prologParserDECIMAL-32))|(1<<(prologParserBINARY-32))|(1<<(prologParserOCTAL-32))|(1<<(prologParserHEX-32))|(1<<(prologParserCHARACTER_CODE_CONSTANT-32))|(1<<(prologParserFLOAT-32))|(1<<(prologParserGRAPHIC_TOKEN-32))|(1<<(prologParserQUOTED-32))|(1<<(prologParserDOUBLE_QUOTED_LIST-32))|(1<<(prologParserBACK_QUOTED_STRING-32)))) != 0) {
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(16)
				p.Directive()
			}

		case 2:
			{
				p.SetState(17)
				p.Clause()
			}

		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(prologParserEOF)
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
	p.RuleIndex = prologParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *prologParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, prologParserRULE_directive)

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
		p.SetState(25)
		p.Match(prologParserT__0)
	}
	{
		p.SetState(26)
		p.term(0)
	}
	{
		p.SetState(27)
		p.Match(prologParserT__1)
	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = prologParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *prologParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, prologParserRULE_clause)

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
		p.SetState(29)
		p.term(0)
	}
	{
		p.SetState(30)
		p.Match(prologParserT__1)
	}

	return localctx
}

// ITermlistContext is an interface to support dynamic dispatch.
type ITermlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermlistContext differentiates from other interfaces.
	IsTermlistContext()
}

type TermlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermlistContext() *TermlistContext {
	var p = new(TermlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = prologParserRULE_termlist
	return p
}

func (*TermlistContext) IsTermlistContext() {}

func NewTermlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermlistContext {
	var p = new(TermlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_termlist

	return p
}

func (s *TermlistContext) GetParser() antlr.Parser { return s.parser }

func (s *TermlistContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TermlistContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterTermlist(s)
	}
}

func (s *TermlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitTermlist(s)
	}
}

func (p *prologParser) Termlist() (localctx ITermlistContext) {
	localctx = NewTermlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, prologParserRULE_termlist)
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
		p.SetState(32)
		p.term(0)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == prologParserT__2 {
		{
			p.SetState(33)
			p.Match(prologParserT__2)
		}
		{
			p.SetState(34)
			p.term(0)
		}

		p.SetState(39)
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
	p.RuleIndex = prologParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Atom_termContext struct {
	*TermContext
}

func NewAtom_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atom_termContext {
	var p = new(Atom_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Atom_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_termContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Atom_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterAtom_term(s)
	}
}

func (s *Atom_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitAtom_term(s)
	}
}

type Binary_operatorContext struct {
	*TermContext
}

func NewBinary_operatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Binary_operatorContext {
	var p = new(Binary_operatorContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_operatorContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Binary_operatorContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Binary_operatorContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *Binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterBinary_operator(s)
	}
}

func (s *Binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitBinary_operator(s)
	}
}

type Unary_operatorContext struct {
	*TermContext
}

func NewUnary_operatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Unary_operatorContext {
	var p = new(Unary_operatorContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_operatorContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *Unary_operatorContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterUnary_operator(s)
	}
}

func (s *Unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitUnary_operator(s)
	}
}

type Braced_termContext struct {
	*TermContext
}

func NewBraced_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Braced_termContext {
	var p = new(Braced_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Braced_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Braced_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Braced_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterBraced_term(s)
	}
}

func (s *Braced_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitBraced_term(s)
	}
}

type List_termContext struct {
	*TermContext
}

func NewList_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *List_termContext {
	var p = new(List_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *List_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_termContext) Termlist() ITermlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermlistContext)
}

func (s *List_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *List_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterList_term(s)
	}
}

func (s *List_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitList_term(s)
	}
}

type VariableContext struct {
	*TermContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(prologParserVARIABLE, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitVariable(s)
	}
}

type FloatContext struct {
	*TermContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(prologParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitFloat(s)
	}
}

type Compound_termContext struct {
	*TermContext
}

func NewCompound_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Compound_termContext {
	var p = new(Compound_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Compound_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_termContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Compound_termContext) Termlist() ITermlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermlistContext)
}

func (s *Compound_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterCompound_term(s)
	}
}

func (s *Compound_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitCompound_term(s)
	}
}

type Integer_termContext struct {
	*TermContext
}

func NewInteger_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Integer_termContext {
	var p = new(Integer_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Integer_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_termContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Integer_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterInteger_term(s)
	}
}

func (s *Integer_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitInteger_term(s)
	}
}

type Curly_bracketed_termContext struct {
	*TermContext
}

func NewCurly_bracketed_termContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Curly_bracketed_termContext {
	var p = new(Curly_bracketed_termContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *Curly_bracketed_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Curly_bracketed_termContext) Termlist() ITermlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermlistContext)
}

func (s *Curly_bracketed_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterCurly_bracketed_term(s)
	}
}

func (s *Curly_bracketed_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitCurly_bracketed_term(s)
	}
}

func (p *prologParser) Term() (localctx ITermContext) {
	return p.term(0)
}

func (p *prologParser) term(_p int) (localctx ITermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, prologParserRULE_term, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(41)
			p.Match(prologParserVARIABLE)
		}

	case 2:
		localctx = NewBraced_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(42)
			p.Match(prologParserT__3)
		}
		{
			p.SetState(43)
			p.term(0)
		}
		{
			p.SetState(44)
			p.Match(prologParserT__4)
		}

	case 3:
		localctx = NewInteger_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == prologParserT__5 {
			{
				p.SetState(46)
				p.Match(prologParserT__5)
			}

		}
		{
			p.SetState(49)
			p.Integer()
		}

	case 4:
		localctx = NewFloatContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == prologParserT__5 {
			{
				p.SetState(50)
				p.Match(prologParserT__5)
			}

		}
		{
			p.SetState(53)
			p.Match(prologParserFLOAT)
		}

	case 5:
		localctx = NewCompound_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Atom()
		}
		{
			p.SetState(55)
			p.Match(prologParserT__3)
		}
		{
			p.SetState(56)
			p.Termlist()
		}
		{
			p.SetState(57)
			p.Match(prologParserT__4)
		}

	case 6:
		localctx = NewUnary_operatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.Operator()
		}
		{
			p.SetState(60)
			p.term(4)
		}

	case 7:
		localctx = NewList_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(prologParserT__6)
		}
		{
			p.SetState(63)
			p.Termlist()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == prologParserT__7 {
			{
				p.SetState(64)
				p.Match(prologParserT__7)
			}
			{
				p.SetState(65)
				p.term(0)
			}

		}
		{
			p.SetState(68)
			p.Match(prologParserT__8)
		}

	case 8:
		localctx = NewCurly_bracketed_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(prologParserT__9)
		}
		{
			p.SetState(71)
			p.Termlist()
		}
		{
			p.SetState(72)
			p.Match(prologParserT__10)
		}

	case 9:
		localctx = NewAtom_termContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Atom()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBinary_operatorContext(p, NewTermContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, prologParserRULE_term)
			p.SetState(77)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(78)
				p.Operator()
			}
			{
				p.SetState(79)
				p.term(5)
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.RuleIndex = prologParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *prologParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, prologParserRULE_operator)
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
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<prologParserT__0)|(1<<prologParserT__2)|(1<<prologParserT__5)|(1<<prologParserT__11)|(1<<prologParserT__12)|(1<<prologParserT__13)|(1<<prologParserT__14)|(1<<prologParserT__15)|(1<<prologParserT__16)|(1<<prologParserT__17)|(1<<prologParserT__18)|(1<<prologParserT__19)|(1<<prologParserT__20)|(1<<prologParserT__21)|(1<<prologParserT__22)|(1<<prologParserT__23)|(1<<prologParserT__24)|(1<<prologParserT__25)|(1<<prologParserT__26)|(1<<prologParserT__27)|(1<<prologParserT__28)|(1<<prologParserT__29)|(1<<prologParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(prologParserT__31-32))|(1<<(prologParserT__32-32))|(1<<(prologParserT__33-32))|(1<<(prologParserT__34-32))|(1<<(prologParserT__35-32))|(1<<(prologParserT__36-32))|(1<<(prologParserT__37-32))|(1<<(prologParserT__38-32))|(1<<(prologParserT__39-32))|(1<<(prologParserT__40-32))|(1<<(prologParserT__41-32))|(1<<(prologParserT__42-32))|(1<<(prologParserT__43-32))|(1<<(prologParserT__44-32))|(1<<(prologParserT__45-32))|(1<<(prologParserT__46-32))|(1<<(prologParserT__47-32))|(1<<(prologParserT__48-32))|(1<<(prologParserT__49-32)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = prologParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Backq_stringContext struct {
	*AtomContext
}

func NewBackq_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Backq_stringContext {
	var p = new(Backq_stringContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *Backq_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Backq_stringContext) BACK_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(prologParserBACK_QUOTED_STRING, 0)
}

func (s *Backq_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterBackq_string(s)
	}
}

func (s *Backq_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitBackq_string(s)
	}
}

type CutContext struct {
	*AtomContext
}

func NewCutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CutContext {
	var p = new(CutContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *CutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterCut(s)
	}
}

func (s *CutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitCut(s)
	}
}

type Empty_bracesContext struct {
	*AtomContext
}

func NewEmpty_bracesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Empty_bracesContext {
	var p = new(Empty_bracesContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *Empty_bracesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_bracesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterEmpty_braces(s)
	}
}

func (s *Empty_bracesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitEmpty_braces(s)
	}
}

type Dq_stringContext struct {
	*AtomContext
}

func NewDq_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Dq_stringContext {
	var p = new(Dq_stringContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *Dq_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dq_stringContext) DOUBLE_QUOTED_LIST() antlr.TerminalNode {
	return s.GetToken(prologParserDOUBLE_QUOTED_LIST, 0)
}

func (s *Dq_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterDq_string(s)
	}
}

func (s *Dq_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitDq_string(s)
	}
}

type NameContext struct {
	*AtomContext
}

func NewNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameContext {
	var p = new(NameContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) LETTER_DIGIT() antlr.TerminalNode {
	return s.GetToken(prologParserLETTER_DIGIT, 0)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitName(s)
	}
}

type Quoted_stringContext struct {
	*AtomContext
}

func NewQuoted_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Quoted_stringContext {
	var p = new(Quoted_stringContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *Quoted_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_stringContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(prologParserQUOTED, 0)
}

func (s *Quoted_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterQuoted_string(s)
	}
}

func (s *Quoted_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitQuoted_string(s)
	}
}

type Empty_listContext struct {
	*AtomContext
}

func NewEmpty_listContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Empty_listContext {
	var p = new(Empty_listContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *Empty_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Empty_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterEmpty_list(s)
	}
}

func (s *Empty_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitEmpty_list(s)
	}
}

type GraphicContext struct {
	*AtomContext
}

func NewGraphicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GraphicContext {
	var p = new(GraphicContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *GraphicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphicContext) GRAPHIC_TOKEN() antlr.TerminalNode {
	return s.GetToken(prologParserGRAPHIC_TOKEN, 0)
}

func (s *GraphicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterGraphic(s)
	}
}

func (s *GraphicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitGraphic(s)
	}
}

type SemicolonContext struct {
	*AtomContext
}

func NewSemicolonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SemicolonContext {
	var p = new(SemicolonContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *SemicolonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SemicolonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterSemicolon(s)
	}
}

func (s *SemicolonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitSemicolon(s)
	}
}

func (p *prologParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, prologParserRULE_atom)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case prologParserT__6:
		localctx = NewEmpty_listContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(prologParserT__6)
		}
		{
			p.SetState(89)
			p.Match(prologParserT__8)
		}

	case prologParserT__9:
		localctx = NewEmpty_bracesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(prologParserT__9)
		}
		{
			p.SetState(91)
			p.Match(prologParserT__10)
		}

	case prologParserLETTER_DIGIT:
		localctx = NewNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(prologParserLETTER_DIGIT)
		}

	case prologParserGRAPHIC_TOKEN:
		localctx = NewGraphicContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(prologParserGRAPHIC_TOKEN)
		}

	case prologParserQUOTED:
		localctx = NewQuoted_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(prologParserQUOTED)
		}

	case prologParserDOUBLE_QUOTED_LIST:
		localctx = NewDq_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(prologParserDOUBLE_QUOTED_LIST)
		}

	case prologParserBACK_QUOTED_STRING:
		localctx = NewBackq_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(96)
			p.Match(prologParserBACK_QUOTED_STRING)
		}

	case prologParserT__17:
		localctx = NewSemicolonContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(97)
			p.Match(prologParserT__17)
		}

	case prologParserT__50:
		localctx = NewCutContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(98)
			p.Match(prologParserT__50)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = prologParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = prologParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(prologParserDECIMAL, 0)
}

func (s *IntegerContext) CHARACTER_CODE_CONSTANT() antlr.TerminalNode {
	return s.GetToken(prologParserCHARACTER_CODE_CONSTANT, 0)
}

func (s *IntegerContext) BINARY() antlr.TerminalNode {
	return s.GetToken(prologParserBINARY, 0)
}

func (s *IntegerContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(prologParserOCTAL, 0)
}

func (s *IntegerContext) HEX() antlr.TerminalNode {
	return s.GetToken(prologParserHEX, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(prologListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *prologParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, prologParserRULE_integer)
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
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(prologParserDECIMAL-54))|(1<<(prologParserBINARY-54))|(1<<(prologParserOCTAL-54))|(1<<(prologParserHEX-54))|(1<<(prologParserCHARACTER_CODE_CONSTANT-54)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *prologParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *TermContext = nil
		if localctx != nil {
			t = localctx.(*TermContext)
		}
		return p.Term_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *prologParser) Term_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
