// Generated from SUOKIF.g4 by ANTLR 4.7.

package suokif // SUOKIF
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 137,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14,
	2, 25, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 33, 10, 3, 3, 4,
	3, 4, 5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 42, 10, 5, 13, 5, 14, 5,
	43, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 6, 8, 65, 10, 8, 13,
	8, 14, 8, 66, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 6, 9, 79, 10, 9, 13, 9, 14, 9, 80, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6,
	9, 88, 10, 9, 13, 9, 14, 9, 89, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 109,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 115, 10, 10, 13, 10, 14, 10,
	116, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 127,
	10, 10, 13, 10, 14, 10, 128, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 135, 10,
	10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 4, 2, 8, 8,
	10, 10, 2, 150, 2, 23, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2,
	8, 38, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 61, 3, 2,
	2, 2, 16, 108, 3, 2, 2, 2, 18, 134, 3, 2, 2, 2, 20, 22, 5, 10, 6, 2, 21,
	20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2,
	2, 24, 3, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 33, 7, 10, 2, 2, 27, 33,
	7, 8, 2, 2, 28, 33, 7, 9, 2, 2, 29, 33, 5, 8, 5, 2, 30, 33, 7, 11, 2, 2,
	31, 33, 5, 10, 6, 2, 32, 26, 3, 2, 2, 2, 32, 27, 3, 2, 2, 2, 32, 28, 3,
	2, 2, 2, 32, 29, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33,
	5, 3, 2, 2, 2, 34, 37, 5, 10, 6, 2, 35, 37, 5, 4, 3, 2, 36, 34, 3, 2, 2,
	2, 36, 35, 3, 2, 2, 2, 37, 7, 3, 2, 2, 2, 38, 39, 7, 14, 2, 2, 39, 41,
	7, 8, 2, 2, 40, 42, 5, 6, 4, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 7,
	15, 2, 2, 46, 9, 3, 2, 2, 2, 47, 54, 7, 8, 2, 2, 48, 54, 5, 12, 7, 2, 49,
	54, 5, 14, 8, 2, 50, 54, 5, 16, 9, 2, 51, 54, 5, 18, 10, 2, 52, 54, 7,
	10, 2, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3, 2, 2, 2, 53, 49, 3, 2, 2, 2, 53,
	50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 11, 3, 2, 2,
	2, 55, 56, 7, 14, 2, 2, 56, 57, 7, 16, 2, 2, 57, 58, 5, 4, 3, 2, 58, 59,
	5, 4, 3, 2, 59, 60, 7, 15, 2, 2, 60, 13, 3, 2, 2, 2, 61, 62, 7, 14, 2,
	2, 62, 64, 9, 2, 2, 2, 63, 65, 5, 6, 4, 2, 64, 63, 3, 2, 2, 2, 65, 66,
	3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 69, 7, 15, 2, 2, 69, 15, 3, 2, 2, 2, 70, 71, 7, 14, 2, 2, 71, 72, 7,
	3, 2, 2, 72, 73, 5, 10, 6, 2, 73, 74, 7, 15, 2, 2, 74, 109, 3, 2, 2, 2,
	75, 76, 7, 14, 2, 2, 76, 78, 7, 4, 2, 2, 77, 79, 5, 10, 6, 2, 78, 77, 3,
	2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81,
	82, 3, 2, 2, 2, 82, 83, 7, 15, 2, 2, 83, 109, 3, 2, 2, 2, 84, 85, 7, 14,
	2, 2, 85, 87, 7, 5, 2, 2, 86, 88, 5, 10, 6, 2, 87, 86, 3, 2, 2, 2, 88,
	89, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91, 3, 2, 2,
	2, 91, 92, 7, 15, 2, 2, 92, 109, 3, 2, 2, 2, 93, 94, 7, 14, 2, 2, 94, 95,
	7, 16, 2, 2, 95, 96, 7, 17, 2, 2, 96, 97, 5, 10, 6, 2, 97, 98, 5, 10, 6,
	2, 98, 99, 7, 15, 2, 2, 99, 109, 3, 2, 2, 2, 100, 101, 7, 14, 2, 2, 101,
	102, 7, 18, 2, 2, 102, 103, 7, 16, 2, 2, 103, 104, 7, 17, 2, 2, 104, 105,
	5, 10, 6, 2, 105, 106, 5, 10, 6, 2, 106, 107, 7, 15, 2, 2, 107, 109, 3,
	2, 2, 2, 108, 70, 3, 2, 2, 2, 108, 75, 3, 2, 2, 2, 108, 84, 3, 2, 2, 2,
	108, 93, 3, 2, 2, 2, 108, 100, 3, 2, 2, 2, 109, 17, 3, 2, 2, 2, 110, 111,
	7, 14, 2, 2, 111, 112, 7, 6, 2, 2, 112, 114, 7, 14, 2, 2, 113, 115, 7,
	10, 2, 2, 114, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 114, 3, 2, 2,
	2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 7, 15, 2, 2, 119,
	120, 5, 10, 6, 2, 120, 121, 7, 15, 2, 2, 121, 135, 3, 2, 2, 2, 122, 123,
	7, 14, 2, 2, 123, 124, 7, 7, 2, 2, 124, 126, 7, 14, 2, 2, 125, 127, 7,
	10, 2, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2,
	2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 7, 15, 2, 2, 131,
	132, 5, 10, 6, 2, 132, 133, 7, 15, 2, 2, 133, 135, 3, 2, 2, 2, 134, 110,
	3, 2, 2, 2, 134, 122, 3, 2, 2, 2, 135, 19, 3, 2, 2, 2, 14, 23, 32, 36,
	43, 53, 66, 80, 89, 108, 116, 128, 134,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'not'", "'and'", "'or'", "'forall'", "'exists'", "", "", "", "", "",
	"", "'('", "')'", "'='", "'>'", "'<'", "'?'",
}
var symbolicNames = []string{
	"", "NOT", "AND", "OR", "FORALL", "EXISTS", "WORD", "STRING", "VARIABLE",
	"NUMBER", "WHITE", "COMMENT", "LPAREN", "RPAREN", "ASSIGN", "GT", "LT",
	"QUESTION",
}

var ruleNames = []string{
	"top_level", "term", "argument", "funterm", "sentence", "equation", "relsent",
	"logsent", "quantsent",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SUOKIFParser struct {
	*antlr.BaseParser
}

func NewSUOKIFParser(input antlr.TokenStream) *SUOKIFParser {
	this := new(SUOKIFParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SUOKIF.g4"

	return this
}

// SUOKIFParser tokens.
const (
	SUOKIFParserEOF      = antlr.TokenEOF
	SUOKIFParserNOT      = 1
	SUOKIFParserAND      = 2
	SUOKIFParserOR       = 3
	SUOKIFParserFORALL   = 4
	SUOKIFParserEXISTS   = 5
	SUOKIFParserWORD     = 6
	SUOKIFParserSTRING   = 7
	SUOKIFParserVARIABLE = 8
	SUOKIFParserNUMBER   = 9
	SUOKIFParserWHITE    = 10
	SUOKIFParserCOMMENT  = 11
	SUOKIFParserLPAREN   = 12
	SUOKIFParserRPAREN   = 13
	SUOKIFParserASSIGN   = 14
	SUOKIFParserGT       = 15
	SUOKIFParserLT       = 16
	SUOKIFParserQUESTION = 17
)

// SUOKIFParser rules.
const (
	SUOKIFParserRULE_top_level = 0
	SUOKIFParserRULE_term      = 1
	SUOKIFParserRULE_argument  = 2
	SUOKIFParserRULE_funterm   = 3
	SUOKIFParserRULE_sentence  = 4
	SUOKIFParserRULE_equation  = 5
	SUOKIFParserRULE_relsent   = 6
	SUOKIFParserRULE_logsent   = 7
	SUOKIFParserRULE_quantsent = 8
)

// ITop_levelContext is an interface to support dynamic dispatch.
type ITop_levelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTop_levelContext differentiates from other interfaces.
	IsTop_levelContext()
}

type Top_levelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_levelContext() *Top_levelContext {
	var p = new(Top_levelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SUOKIFParserRULE_top_level
	return p
}

func (*Top_levelContext) IsTop_levelContext() {}

func NewTop_levelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_levelContext {
	var p = new(Top_levelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_top_level

	return p
}

func (s *Top_levelContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_levelContext) AllSentence() []ISentenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISentenceContext)(nil)).Elem())
	var tst = make([]ISentenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISentenceContext)
		}
	}

	return tst
}

func (s *Top_levelContext) Sentence(i int) ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *Top_levelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_levelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Top_levelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterTop_level(s)
	}
}

func (s *Top_levelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitTop_level(s)
	}
}

func (p *SUOKIFParser) Top_level() (localctx ITop_levelContext) {
	localctx = NewTop_levelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SUOKIFParserRULE_top_level)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SUOKIFParserWORD)|(1<<SUOKIFParserVARIABLE)|(1<<SUOKIFParserLPAREN))) != 0 {
		{
			p.SetState(18)
			p.Sentence()
		}

		p.SetState(23)
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
	p.RuleIndex = SUOKIFParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserVARIABLE, 0)
}

func (s *TermContext) WORD() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserWORD, 0)
}

func (s *TermContext) STRING() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserSTRING, 0)
}

func (s *TermContext) Funterm() IFuntermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuntermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuntermContext)
}

func (s *TermContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserNUMBER, 0)
}

func (s *TermContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *SUOKIFParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SUOKIFParserRULE_term)

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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(SUOKIFParserVARIABLE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(SUOKIFParserWORD)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(SUOKIFParserSTRING)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Funterm()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.Match(SUOKIFParserNUMBER)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(29)
			p.Sentence()
		}

	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SUOKIFParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *ArgumentContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *SUOKIFParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SUOKIFParserRULE_argument)

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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Sentence()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Term()
		}

	}

	return localctx
}

// IFuntermContext is an interface to support dynamic dispatch.
type IFuntermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuntermContext differentiates from other interfaces.
	IsFuntermContext()
}

type FuntermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuntermContext() *FuntermContext {
	var p = new(FuntermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SUOKIFParserRULE_funterm
	return p
}

func (*FuntermContext) IsFuntermContext() {}

func NewFuntermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuntermContext {
	var p = new(FuntermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_funterm

	return p
}

func (s *FuntermContext) GetParser() antlr.Parser { return s.parser }

func (s *FuntermContext) WORD() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserWORD, 0)
}

func (s *FuntermContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *FuntermContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *FuntermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuntermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuntermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterFunterm(s)
	}
}

func (s *FuntermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitFunterm(s)
	}
}

func (p *SUOKIFParser) Funterm() (localctx IFuntermContext) {
	localctx = NewFuntermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SUOKIFParserRULE_funterm)
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
		p.SetState(36)
		p.Match(SUOKIFParserLPAREN)
	}
	{
		p.SetState(37)
		p.Match(SUOKIFParserWORD)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SUOKIFParserWORD)|(1<<SUOKIFParserSTRING)|(1<<SUOKIFParserVARIABLE)|(1<<SUOKIFParserNUMBER)|(1<<SUOKIFParserLPAREN))) != 0) {
		{
			p.SetState(38)
			p.Argument()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(SUOKIFParserRPAREN)
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
	p.RuleIndex = SUOKIFParserRULE_sentence
	return p
}

func (*SentenceContext) IsSentenceContext() {}

func NewSentenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenceContext {
	var p = new(SentenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_sentence

	return p
}

func (s *SentenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenceContext) WORD() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserWORD, 0)
}

func (s *SentenceContext) Equation() IEquationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *SentenceContext) Relsent() IRelsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelsentContext)
}

func (s *SentenceContext) Logsent() ILogsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogsentContext)
}

func (s *SentenceContext) Quantsent() IQuantsentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantsentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantsentContext)
}

func (s *SentenceContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserVARIABLE, 0)
}

func (s *SentenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterSentence(s)
	}
}

func (s *SentenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitSentence(s)
	}
}

func (p *SUOKIFParser) Sentence() (localctx ISentenceContext) {
	localctx = NewSentenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SUOKIFParserRULE_sentence)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(SUOKIFParserWORD)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Equation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.Relsent()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
			p.Logsent()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(49)
			p.Quantsent()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(50)
			p.Match(SUOKIFParserVARIABLE)
		}

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
	p.RuleIndex = SUOKIFParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

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

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *SUOKIFParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SUOKIFParserRULE_equation)

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
		p.Match(SUOKIFParserLPAREN)
	}
	{
		p.SetState(54)
		p.Match(SUOKIFParserASSIGN)
	}
	{
		p.SetState(55)
		p.Term()
	}
	{
		p.SetState(56)
		p.Term()
	}
	{
		p.SetState(57)
		p.Match(SUOKIFParserRPAREN)
	}

	return localctx
}

// IRelsentContext is an interface to support dynamic dispatch.
type IRelsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelsentContext differentiates from other interfaces.
	IsRelsentContext()
}

type RelsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelsentContext() *RelsentContext {
	var p = new(RelsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SUOKIFParserRULE_relsent
	return p
}

func (*RelsentContext) IsRelsentContext() {}

func NewRelsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelsentContext {
	var p = new(RelsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_relsent

	return p
}

func (s *RelsentContext) GetParser() antlr.Parser { return s.parser }

func (s *RelsentContext) WORD() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserWORD, 0)
}

func (s *RelsentContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserVARIABLE, 0)
}

func (s *RelsentContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *RelsentContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *RelsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterRelsent(s)
	}
}

func (s *RelsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitRelsent(s)
	}
}

func (p *SUOKIFParser) Relsent() (localctx IRelsentContext) {
	localctx = NewRelsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SUOKIFParserRULE_relsent)
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
		p.Match(SUOKIFParserLPAREN)
	}
	p.SetState(60)
	_la = p.GetTokenStream().LA(1)

	if !(_la == SUOKIFParserWORD || _la == SUOKIFParserVARIABLE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SUOKIFParserWORD)|(1<<SUOKIFParserSTRING)|(1<<SUOKIFParserVARIABLE)|(1<<SUOKIFParserNUMBER)|(1<<SUOKIFParserLPAREN))) != 0) {
		{
			p.SetState(61)
			p.Argument()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(SUOKIFParserRPAREN)
	}

	return localctx
}

// ILogsentContext is an interface to support dynamic dispatch.
type ILogsentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogsentContext differentiates from other interfaces.
	IsLogsentContext()
}

type LogsentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogsentContext() *LogsentContext {
	var p = new(LogsentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SUOKIFParserRULE_logsent
	return p
}

func (*LogsentContext) IsLogsentContext() {}

func NewLogsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogsentContext {
	var p = new(LogsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_logsent

	return p
}

func (s *LogsentContext) GetParser() antlr.Parser { return s.parser }

func (s *LogsentContext) NOT() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserNOT, 0)
}

func (s *LogsentContext) AllSentence() []ISentenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISentenceContext)(nil)).Elem())
	var tst = make([]ISentenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISentenceContext)
		}
	}

	return tst
}

func (s *LogsentContext) Sentence(i int) ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *LogsentContext) AND() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserAND, 0)
}

func (s *LogsentContext) OR() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserOR, 0)
}

func (s *LogsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterLogsent(s)
	}
}

func (s *LogsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitLogsent(s)
	}
}

func (p *SUOKIFParser) Logsent() (localctx ILogsentContext) {
	localctx = NewLogsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SUOKIFParserRULE_logsent)
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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(69)
			p.Match(SUOKIFParserNOT)
		}
		{
			p.SetState(70)
			p.Sentence()
		}
		{
			p.SetState(71)
			p.Match(SUOKIFParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(74)
			p.Match(SUOKIFParserAND)
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SUOKIFParserWORD)|(1<<SUOKIFParserVARIABLE)|(1<<SUOKIFParserLPAREN))) != 0) {
			{
				p.SetState(75)
				p.Sentence()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)
			p.Match(SUOKIFParserRPAREN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(83)
			p.Match(SUOKIFParserOR)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SUOKIFParserWORD)|(1<<SUOKIFParserVARIABLE)|(1<<SUOKIFParserLPAREN))) != 0) {
			{
				p.SetState(84)
				p.Sentence()
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(89)
			p.Match(SUOKIFParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(92)
			p.Match(SUOKIFParserASSIGN)
		}
		{
			p.SetState(93)
			p.Match(SUOKIFParserGT)
		}
		{
			p.SetState(94)
			p.Sentence()
		}
		{
			p.SetState(95)
			p.Sentence()
		}
		{
			p.SetState(96)
			p.Match(SUOKIFParserRPAREN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(99)
			p.Match(SUOKIFParserLT)
		}
		{
			p.SetState(100)
			p.Match(SUOKIFParserASSIGN)
		}
		{
			p.SetState(101)
			p.Match(SUOKIFParserGT)
		}
		{
			p.SetState(102)
			p.Sentence()
		}
		{
			p.SetState(103)
			p.Sentence()
		}
		{
			p.SetState(104)
			p.Match(SUOKIFParserRPAREN)
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
	p.RuleIndex = SUOKIFParserRULE_quantsent
	return p
}

func (*QuantsentContext) IsQuantsentContext() {}

func NewQuantsentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantsentContext {
	var p = new(QuantsentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SUOKIFParserRULE_quantsent

	return p
}

func (s *QuantsentContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantsentContext) FORALL() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserFORALL, 0)
}

func (s *QuantsentContext) Sentence() ISentenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISentenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISentenceContext)
}

func (s *QuantsentContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(SUOKIFParserVARIABLE)
}

func (s *QuantsentContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(SUOKIFParserVARIABLE, i)
}

func (s *QuantsentContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SUOKIFParserEXISTS, 0)
}

func (s *QuantsentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantsentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantsentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.EnterQuantsent(s)
	}
}

func (s *QuantsentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SUOKIFListener); ok {
		listenerT.ExitQuantsent(s)
	}
}

func (p *SUOKIFParser) Quantsent() (localctx IQuantsentContext) {
	localctx = NewQuantsentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SUOKIFParserRULE_quantsent)
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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(109)
			p.Match(SUOKIFParserFORALL)
		}
		{
			p.SetState(110)
			p.Match(SUOKIFParserLPAREN)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SUOKIFParserVARIABLE {
			{
				p.SetState(111)
				p.Match(SUOKIFParserVARIABLE)
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(SUOKIFParserRPAREN)
		}
		{
			p.SetState(117)
			p.Sentence()
		}
		{
			p.SetState(118)
			p.Match(SUOKIFParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(SUOKIFParserLPAREN)
		}
		{
			p.SetState(121)
			p.Match(SUOKIFParserEXISTS)
		}
		{
			p.SetState(122)
			p.Match(SUOKIFParserLPAREN)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SUOKIFParserVARIABLE {
			{
				p.SetState(123)
				p.Match(SUOKIFParserVARIABLE)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(128)
			p.Match(SUOKIFParserRPAREN)
		}
		{
			p.SetState(129)
			p.Sentence()
		}
		{
			p.SetState(130)
			p.Match(SUOKIFParserRPAREN)
		}

	}

	return localctx
}
