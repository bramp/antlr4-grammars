// Code generated from regexParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package regex // regexParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 156,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 7, 3, 49, 10, 3, 12, 3, 14, 3, 52, 11, 3, 3, 4, 7, 4, 55, 10,
	4, 12, 4, 14, 4, 58, 11, 4, 3, 5, 3, 5, 5, 5, 62, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 71, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 76,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 91, 10, 10, 3, 11, 3, 11, 3, 11, 5, 11, 96, 10,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13, 103, 10, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 114, 10, 13, 3,
	13, 5, 13, 117, 10, 13, 3, 14, 5, 14, 120, 10, 14, 3, 14, 3, 14, 6, 14,
	124, 10, 14, 13, 14, 14, 14, 125, 3, 15, 3, 15, 5, 15, 130, 10, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 5, 18, 144, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 7, 4, 2, 16, 17, 35, 36,
	4, 2, 12, 12, 39, 39, 4, 2, 14, 14, 33, 33, 4, 2, 15, 15, 34, 34, 4, 2,
	22, 22, 30, 30, 2, 161, 2, 42, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 6, 56, 3,
	2, 2, 2, 8, 59, 3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14,
	77, 3, 2, 2, 2, 16, 81, 3, 2, 2, 2, 18, 90, 3, 2, 2, 2, 20, 95, 3, 2, 2,
	2, 22, 97, 3, 2, 2, 2, 24, 116, 3, 2, 2, 2, 26, 119, 3, 2, 2, 2, 28, 129,
	3, 2, 2, 2, 30, 131, 3, 2, 2, 2, 32, 135, 3, 2, 2, 2, 34, 143, 3, 2, 2,
	2, 36, 145, 3, 2, 2, 2, 38, 149, 3, 2, 2, 2, 40, 153, 3, 2, 2, 2, 42, 43,
	5, 4, 3, 2, 43, 44, 7, 2, 2, 3, 44, 3, 3, 2, 2, 2, 45, 50, 5, 6, 4, 2,
	46, 47, 7, 5, 2, 2, 47, 49, 5, 6, 4, 2, 48, 46, 3, 2, 2, 2, 49, 52, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 5, 3, 2, 2, 2, 52,
	50, 3, 2, 2, 2, 53, 55, 5, 8, 5, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2,
	2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 7, 3, 2, 2, 2, 58, 56, 3,
	2, 2, 2, 59, 61, 5, 18, 10, 2, 60, 62, 5, 10, 6, 2, 61, 60, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 62, 9, 3, 2, 2, 2, 63, 71, 7, 7, 2, 2, 64, 71, 7, 8,
	2, 2, 65, 71, 7, 6, 2, 2, 66, 67, 7, 11, 2, 2, 67, 68, 5, 12, 7, 2, 68,
	69, 7, 18, 2, 2, 69, 71, 3, 2, 2, 2, 70, 63, 3, 2, 2, 2, 70, 64, 3, 2,
	2, 2, 70, 65, 3, 2, 2, 2, 70, 66, 3, 2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 76,
	5, 14, 8, 2, 73, 76, 5, 16, 9, 2, 74, 76, 7, 19, 2, 2, 75, 72, 3, 2, 2,
	2, 75, 73, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 13, 3, 2, 2, 2, 77, 78,
	7, 19, 2, 2, 78, 79, 7, 20, 2, 2, 79, 80, 7, 19, 2, 2, 80, 15, 3, 2, 2,
	2, 81, 82, 7, 19, 2, 2, 82, 83, 7, 20, 2, 2, 83, 17, 3, 2, 2, 2, 84, 91,
	7, 10, 2, 2, 85, 91, 5, 20, 11, 2, 86, 87, 7, 3, 2, 2, 87, 88, 5, 4, 3,
	2, 88, 89, 7, 4, 2, 2, 89, 91, 3, 2, 2, 2, 90, 84, 3, 2, 2, 2, 90, 85,
	3, 2, 2, 2, 90, 86, 3, 2, 2, 2, 91, 19, 3, 2, 2, 2, 92, 96, 5, 34, 18,
	2, 93, 96, 5, 22, 12, 2, 94, 96, 7, 9, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93,
	3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 21, 3, 2, 2, 2, 97, 98, 9, 2, 2, 2,
	98, 99, 5, 24, 13, 2, 99, 100, 7, 37, 2, 2, 100, 23, 3, 2, 2, 2, 101, 103,
	5, 26, 14, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3,
	2, 2, 2, 104, 105, 7, 38, 2, 2, 105, 106, 7, 38, 2, 2, 106, 117, 5, 22,
	12, 2, 107, 108, 5, 26, 14, 2, 108, 109, 7, 38, 2, 2, 109, 110, 5, 22,
	12, 2, 110, 117, 3, 2, 2, 2, 111, 113, 5, 26, 14, 2, 112, 114, 7, 38, 2,
	2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115,
	117, 7, 38, 2, 2, 116, 102, 3, 2, 2, 2, 116, 107, 3, 2, 2, 2, 116, 111,
	3, 2, 2, 2, 116, 115, 3, 2, 2, 2, 117, 25, 3, 2, 2, 2, 118, 120, 7, 38,
	2, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2,
	121, 124, 5, 28, 15, 2, 122, 124, 5, 34, 18, 2, 123, 121, 3, 2, 2, 2, 123,
	122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 27, 3, 2, 2, 2, 127, 130, 5, 30, 16, 2, 128, 130, 7, 39,
	2, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3, 2, 2, 2, 130, 29, 3, 2, 2, 2,
	131, 132, 5, 32, 17, 2, 132, 133, 7, 38, 2, 2, 133, 134, 5, 32, 17, 2,
	134, 31, 3, 2, 2, 2, 135, 136, 9, 3, 2, 2, 136, 33, 3, 2, 2, 2, 137, 144,
	7, 12, 2, 2, 138, 144, 7, 31, 2, 2, 139, 144, 7, 13, 2, 2, 140, 144, 7,
	32, 2, 2, 141, 144, 5, 36, 19, 2, 142, 144, 5, 38, 20, 2, 143, 137, 3,
	2, 2, 2, 143, 138, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2, 143, 140, 3, 2, 2,
	2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 35, 3, 2, 2, 2, 145,
	146, 9, 4, 2, 2, 146, 147, 5, 40, 21, 2, 147, 148, 7, 21, 2, 2, 148, 37,
	3, 2, 2, 2, 149, 150, 9, 5, 2, 2, 150, 151, 5, 40, 21, 2, 151, 152, 7,
	21, 2, 2, 152, 39, 3, 2, 2, 2, 153, 154, 9, 6, 2, 2, 154, 41, 3, 2, 2,
	2, 17, 50, 56, 61, 70, 75, 90, 95, 102, 113, 116, 119, 123, 125, 129, 143,
}
var literalNames = []string{
	"", "'('", "')'", "'|'", "'+'", "'?'", "'*'", "'.'", "", "'{'", "", "",
	"", "", "", "", "", "", "','", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "']'", "'-'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc",
	"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc", "ComplEsc",
	"NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact", "COMMA", "EndCategory",
	"IsCategory", "Letters", "Marks", "Numbers", "Punctuation", "Separators",
	"Symbols", "Others", "IsBlock", "NestedSingleCharEsc", "NestedMultiCharEsc",
	"NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup", "NestedPosCharGroup",
	"EndCharGroup", "DASH", "XmlChar",
}

var ruleNames = []string{
	"root", "regExp", "branch", "piece", "quantifier", "quantity", "quantRange",
	"quantMin", "atom", "charClass", "charClassExpr", "charGroup", "posCharGroup",
	"charRange", "seRange", "charOrEsc", "charClassEsc", "catEsc", "complEsc",
	"charProp",
}

type regexParser struct {
	*antlr.BaseParser
}

// NewregexParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *regexParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewregexParser(input antlr.TokenStream) *regexParser {
	this := new(regexParser)
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
	this.GrammarFileName = "regexParser.g4"

	return this
}

// regexParser tokens.
const (
	regexParserEOF                 = antlr.TokenEOF
	regexParserLPAREN              = 1
	regexParserRPAREN              = 2
	regexParserPIPE                = 3
	regexParserPLUS                = 4
	regexParserQUESTION            = 5
	regexParserSTAR                = 6
	regexParserWildcardEsc         = 7
	regexParserChar                = 8
	regexParserStartQuantity       = 9
	regexParserSingleCharEsc       = 10
	regexParserMultiCharEsc        = 11
	regexParserCatEsc              = 12
	regexParserComplEsc            = 13
	regexParserNegCharGroup        = 14
	regexParserPosCharGroup        = 15
	regexParserEndQuantity         = 16
	regexParserQuantExact          = 17
	regexParserCOMMA               = 18
	regexParserEndCategory         = 19
	regexParserIsCategory          = 20
	regexParserLetters             = 21
	regexParserMarks               = 22
	regexParserNumbers             = 23
	regexParserPunctuation         = 24
	regexParserSeparators          = 25
	regexParserSymbols             = 26
	regexParserOthers              = 27
	regexParserIsBlock             = 28
	regexParserNestedSingleCharEsc = 29
	regexParserNestedMultiCharEsc  = 30
	regexParserNestedCatEsc        = 31
	regexParserNestedComplEsc      = 32
	regexParserNestedNegCharGroup  = 33
	regexParserNestedPosCharGroup  = 34
	regexParserEndCharGroup        = 35
	regexParserDASH                = 36
	regexParserXmlChar             = 37
)

// regexParser rules.
const (
	regexParserRULE_root          = 0
	regexParserRULE_regExp        = 1
	regexParserRULE_branch        = 2
	regexParserRULE_piece         = 3
	regexParserRULE_quantifier    = 4
	regexParserRULE_quantity      = 5
	regexParserRULE_quantRange    = 6
	regexParserRULE_quantMin      = 7
	regexParserRULE_atom          = 8
	regexParserRULE_charClass     = 9
	regexParserRULE_charClassExpr = 10
	regexParserRULE_charGroup     = 11
	regexParserRULE_posCharGroup  = 12
	regexParserRULE_charRange     = 13
	regexParserRULE_seRange       = 14
	regexParserRULE_charOrEsc     = 15
	regexParserRULE_charClassEsc  = 16
	regexParserRULE_catEsc        = 17
	regexParserRULE_complEsc      = 18
	regexParserRULE_charProp      = 19
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) RegExp() IRegExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegExpContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(regexParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *regexParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, regexParserRULE_root)

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
		p.RegExp()
	}
	{
		p.SetState(41)
		p.Match(regexParserEOF)
	}

	return localctx
}

// IRegExpContext is an interface to support dynamic dispatch.
type IRegExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegExpContext differentiates from other interfaces.
	IsRegExpContext()
}

type RegExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegExpContext() *RegExpContext {
	var p = new(RegExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_regExp
	return p
}

func (*RegExpContext) IsRegExpContext() {}

func NewRegExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegExpContext {
	var p = new(RegExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_regExp

	return p
}

func (s *RegExpContext) GetParser() antlr.Parser { return s.parser }

func (s *RegExpContext) AllBranch() []IBranchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBranchContext)(nil)).Elem())
	var tst = make([]IBranchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBranchContext)
		}
	}

	return tst
}

func (s *RegExpContext) Branch(i int) IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *RegExpContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(regexParserPIPE)
}

func (s *RegExpContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(regexParserPIPE, i)
}

func (s *RegExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterRegExp(s)
	}
}

func (s *RegExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitRegExp(s)
	}
}

func (p *regexParser) RegExp() (localctx IRegExpContext) {
	this := p
	_ = this

	localctx = NewRegExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, regexParserRULE_regExp)
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
		p.SetState(43)
		p.Branch()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == regexParserPIPE {
		{
			p.SetState(44)
			p.Match(regexParserPIPE)
		}
		{
			p.SetState(45)
			p.Branch()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = regexParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) AllPiece() []IPieceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPieceContext)(nil)).Elem())
	var tst = make([]IPieceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPieceContext)
		}
	}

	return tst
}

func (s *BranchContext) Piece(i int) IPieceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPieceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPieceContext)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (p *regexParser) Branch() (localctx IBranchContext) {
	this := p
	_ = this

	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, regexParserRULE_branch)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<regexParserLPAREN)|(1<<regexParserWildcardEsc)|(1<<regexParserChar)|(1<<regexParserSingleCharEsc)|(1<<regexParserMultiCharEsc)|(1<<regexParserCatEsc)|(1<<regexParserComplEsc)|(1<<regexParserNegCharGroup)|(1<<regexParserPosCharGroup)|(1<<regexParserNestedSingleCharEsc)|(1<<regexParserNestedMultiCharEsc)|(1<<regexParserNestedCatEsc))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(regexParserNestedComplEsc-32))|(1<<(regexParserNestedNegCharGroup-32))|(1<<(regexParserNestedPosCharGroup-32)))) != 0) {
		{
			p.SetState(51)
			p.Piece()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPieceContext is an interface to support dynamic dispatch.
type IPieceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPieceContext differentiates from other interfaces.
	IsPieceContext()
}

type PieceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPieceContext() *PieceContext {
	var p = new(PieceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_piece
	return p
}

func (*PieceContext) IsPieceContext() {}

func NewPieceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PieceContext {
	var p = new(PieceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_piece

	return p
}

func (s *PieceContext) GetParser() antlr.Parser { return s.parser }

func (s *PieceContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *PieceContext) Quantifier() IQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantifierContext)
}

func (s *PieceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PieceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PieceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterPiece(s)
	}
}

func (s *PieceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitPiece(s)
	}
}

func (p *regexParser) Piece() (localctx IPieceContext) {
	this := p
	_ = this

	localctx = NewPieceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, regexParserRULE_piece)
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
		p.Atom()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<regexParserPLUS)|(1<<regexParserQUESTION)|(1<<regexParserSTAR)|(1<<regexParserStartQuantity))) != 0 {
		{
			p.SetState(58)
			p.Quantifier()
		}

	}

	return localctx
}

// IQuantifierContext is an interface to support dynamic dispatch.
type IQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantifierContext differentiates from other interfaces.
	IsQuantifierContext()
}

type QuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantifierContext() *QuantifierContext {
	var p = new(QuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_quantifier
	return p
}

func (*QuantifierContext) IsQuantifierContext() {}

func NewQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantifierContext {
	var p = new(QuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantifier

	return p
}

func (s *QuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantifierContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(regexParserQUESTION, 0)
}

func (s *QuantifierContext) STAR() antlr.TerminalNode {
	return s.GetToken(regexParserSTAR, 0)
}

func (s *QuantifierContext) PLUS() antlr.TerminalNode {
	return s.GetToken(regexParserPLUS, 0)
}

func (s *QuantifierContext) StartQuantity() antlr.TerminalNode {
	return s.GetToken(regexParserStartQuantity, 0)
}

func (s *QuantifierContext) Quantity() IQuantityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantityContext)
}

func (s *QuantifierContext) EndQuantity() antlr.TerminalNode {
	return s.GetToken(regexParserEndQuantity, 0)
}

func (s *QuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantifier(s)
	}
}

func (s *QuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantifier(s)
	}
}

func (p *regexParser) Quantifier() (localctx IQuantifierContext) {
	this := p
	_ = this

	localctx = NewQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, regexParserRULE_quantifier)

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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case regexParserQUESTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(regexParserQUESTION)
		}

	case regexParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(regexParserSTAR)
		}

	case regexParserPLUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(regexParserPLUS)
		}

	case regexParserStartQuantity:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Match(regexParserStartQuantity)
		}
		{
			p.SetState(65)
			p.Quantity()
		}
		{
			p.SetState(66)
			p.Match(regexParserEndQuantity)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuantityContext is an interface to support dynamic dispatch.
type IQuantityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantityContext differentiates from other interfaces.
	IsQuantityContext()
}

type QuantityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantityContext() *QuantityContext {
	var p = new(QuantityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_quantity
	return p
}

func (*QuantityContext) IsQuantityContext() {}

func NewQuantityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantityContext {
	var p = new(QuantityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantity

	return p
}

func (s *QuantityContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantityContext) QuantRange() IQuantRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantRangeContext)
}

func (s *QuantityContext) QuantMin() IQuantMinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantMinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantMinContext)
}

func (s *QuantityContext) QuantExact() antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, 0)
}

func (s *QuantityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantity(s)
	}
}

func (s *QuantityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantity(s)
	}
}

func (p *regexParser) Quantity() (localctx IQuantityContext) {
	this := p
	_ = this

	localctx = NewQuantityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, regexParserRULE_quantity)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.QuantRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.QuantMin()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(regexParserQuantExact)
		}

	}

	return localctx
}

// IQuantRangeContext is an interface to support dynamic dispatch.
type IQuantRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantRangeContext differentiates from other interfaces.
	IsQuantRangeContext()
}

type QuantRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantRangeContext() *QuantRangeContext {
	var p = new(QuantRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_quantRange
	return p
}

func (*QuantRangeContext) IsQuantRangeContext() {}

func NewQuantRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantRangeContext {
	var p = new(QuantRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantRange

	return p
}

func (s *QuantRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantRangeContext) AllQuantExact() []antlr.TerminalNode {
	return s.GetTokens(regexParserQuantExact)
}

func (s *QuantRangeContext) QuantExact(i int) antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, i)
}

func (s *QuantRangeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(regexParserCOMMA, 0)
}

func (s *QuantRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantRange(s)
	}
}

func (s *QuantRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantRange(s)
	}
}

func (p *regexParser) QuantRange() (localctx IQuantRangeContext) {
	this := p
	_ = this

	localctx = NewQuantRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, regexParserRULE_quantRange)

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
		p.Match(regexParserQuantExact)
	}
	{
		p.SetState(76)
		p.Match(regexParserCOMMA)
	}
	{
		p.SetState(77)
		p.Match(regexParserQuantExact)
	}

	return localctx
}

// IQuantMinContext is an interface to support dynamic dispatch.
type IQuantMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantMinContext differentiates from other interfaces.
	IsQuantMinContext()
}

type QuantMinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantMinContext() *QuantMinContext {
	var p = new(QuantMinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_quantMin
	return p
}

func (*QuantMinContext) IsQuantMinContext() {}

func NewQuantMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantMinContext {
	var p = new(QuantMinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantMin

	return p
}

func (s *QuantMinContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantMinContext) QuantExact() antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, 0)
}

func (s *QuantMinContext) COMMA() antlr.TerminalNode {
	return s.GetToken(regexParserCOMMA, 0)
}

func (s *QuantMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantMinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantMin(s)
	}
}

func (s *QuantMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantMin(s)
	}
}

func (p *regexParser) QuantMin() (localctx IQuantMinContext) {
	this := p
	_ = this

	localctx = NewQuantMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, regexParserRULE_quantMin)

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
		p.Match(regexParserQuantExact)
	}
	{
		p.SetState(80)
		p.Match(regexParserCOMMA)
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
	p.RuleIndex = regexParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Char() antlr.TerminalNode {
	return s.GetToken(regexParserChar, 0)
}

func (s *AtomContext) CharClass() ICharClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharClassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharClassContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(regexParserLPAREN, 0)
}

func (s *AtomContext) RegExp() IRegExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegExpContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(regexParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *regexParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, regexParserRULE_atom)

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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case regexParserChar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(regexParserChar)
		}

	case regexParserWildcardEsc, regexParserSingleCharEsc, regexParserMultiCharEsc, regexParserCatEsc, regexParserComplEsc, regexParserNegCharGroup, regexParserPosCharGroup, regexParserNestedSingleCharEsc, regexParserNestedMultiCharEsc, regexParserNestedCatEsc, regexParserNestedComplEsc, regexParserNestedNegCharGroup, regexParserNestedPosCharGroup:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.CharClass()
		}

	case regexParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(regexParserLPAREN)
		}
		{
			p.SetState(85)
			p.RegExp()
		}
		{
			p.SetState(86)
			p.Match(regexParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICharClassContext is an interface to support dynamic dispatch.
type ICharClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharClassContext differentiates from other interfaces.
	IsCharClassContext()
}

type CharClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassContext() *CharClassContext {
	var p = new(CharClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charClass
	return p
}

func (*CharClassContext) IsCharClassContext() {}

func NewCharClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassContext {
	var p = new(CharClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClass

	return p
}

func (s *CharClassContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassContext) CharClassEsc() ICharClassEscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharClassEscContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharClassEscContext)
}

func (s *CharClassContext) CharClassExpr() ICharClassExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharClassExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharClassExprContext)
}

func (s *CharClassContext) WildcardEsc() antlr.TerminalNode {
	return s.GetToken(regexParserWildcardEsc, 0)
}

func (s *CharClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClass(s)
	}
}

func (s *CharClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClass(s)
	}
}

func (p *regexParser) CharClass() (localctx ICharClassContext) {
	this := p
	_ = this

	localctx = NewCharClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, regexParserRULE_charClass)

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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case regexParserSingleCharEsc, regexParserMultiCharEsc, regexParserCatEsc, regexParserComplEsc, regexParserNestedSingleCharEsc, regexParserNestedMultiCharEsc, regexParserNestedCatEsc, regexParserNestedComplEsc:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.CharClassEsc()
		}

	case regexParserNegCharGroup, regexParserPosCharGroup, regexParserNestedNegCharGroup, regexParserNestedPosCharGroup:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.CharClassExpr()
		}

	case regexParserWildcardEsc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(regexParserWildcardEsc)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICharClassExprContext is an interface to support dynamic dispatch.
type ICharClassExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharClassExprContext differentiates from other interfaces.
	IsCharClassExprContext()
}

type CharClassExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassExprContext() *CharClassExprContext {
	var p = new(CharClassExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charClassExpr
	return p
}

func (*CharClassExprContext) IsCharClassExprContext() {}

func NewCharClassExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassExprContext {
	var p = new(CharClassExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClassExpr

	return p
}

func (s *CharClassExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassExprContext) CharGroup() ICharGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharGroupContext)
}

func (s *CharClassExprContext) EndCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserEndCharGroup, 0)
}

func (s *CharClassExprContext) NegCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNegCharGroup, 0)
}

func (s *CharClassExprContext) NestedNegCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNestedNegCharGroup, 0)
}

func (s *CharClassExprContext) PosCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserPosCharGroup, 0)
}

func (s *CharClassExprContext) NestedPosCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNestedPosCharGroup, 0)
}

func (s *CharClassExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClassExpr(s)
	}
}

func (s *CharClassExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClassExpr(s)
	}
}

func (p *regexParser) CharClassExpr() (localctx ICharClassExprContext) {
	this := p
	_ = this

	localctx = NewCharClassExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, regexParserRULE_charClassExpr)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(regexParserNegCharGroup-14))|(1<<(regexParserPosCharGroup-14))|(1<<(regexParserNestedNegCharGroup-14))|(1<<(regexParserNestedPosCharGroup-14)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(96)
		p.CharGroup()
	}
	{
		p.SetState(97)
		p.Match(regexParserEndCharGroup)
	}

	return localctx
}

// ICharGroupContext is an interface to support dynamic dispatch.
type ICharGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharGroupContext differentiates from other interfaces.
	IsCharGroupContext()
}

type CharGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharGroupContext() *CharGroupContext {
	var p = new(CharGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charGroup
	return p
}

func (*CharGroupContext) IsCharGroupContext() {}

func NewCharGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharGroupContext {
	var p = new(CharGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charGroup

	return p
}

func (s *CharGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *CharGroupContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(regexParserDASH)
}

func (s *CharGroupContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(regexParserDASH, i)
}

func (s *CharGroupContext) CharClassExpr() ICharClassExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharClassExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharClassExprContext)
}

func (s *CharGroupContext) PosCharGroup() IPosCharGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPosCharGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPosCharGroupContext)
}

func (s *CharGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharGroup(s)
	}
}

func (s *CharGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharGroup(s)
	}
}

func (p *regexParser) CharGroup() (localctx ICharGroupContext) {
	this := p
	_ = this

	localctx = NewCharGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, regexParserRULE_charGroup)
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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(99)
				p.PosCharGroup()
			}

		}
		{
			p.SetState(102)
			p.Match(regexParserDASH)
		}
		{
			p.SetState(103)
			p.Match(regexParserDASH)
		}
		{
			p.SetState(104)
			p.CharClassExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.PosCharGroup()
		}
		{
			p.SetState(106)
			p.Match(regexParserDASH)
		}
		{
			p.SetState(107)
			p.CharClassExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.PosCharGroup()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == regexParserDASH {
			{
				p.SetState(110)
				p.Match(regexParserDASH)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(regexParserDASH)
		}

	}

	return localctx
}

// IPosCharGroupContext is an interface to support dynamic dispatch.
type IPosCharGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPosCharGroupContext differentiates from other interfaces.
	IsPosCharGroupContext()
}

type PosCharGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPosCharGroupContext() *PosCharGroupContext {
	var p = new(PosCharGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_posCharGroup
	return p
}

func (*PosCharGroupContext) IsPosCharGroupContext() {}

func NewPosCharGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PosCharGroupContext {
	var p = new(PosCharGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_posCharGroup

	return p
}

func (s *PosCharGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *PosCharGroupContext) DASH() antlr.TerminalNode {
	return s.GetToken(regexParserDASH, 0)
}

func (s *PosCharGroupContext) AllCharRange() []ICharRangeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharRangeContext)(nil)).Elem())
	var tst = make([]ICharRangeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharRangeContext)
		}
	}

	return tst
}

func (s *PosCharGroupContext) CharRange(i int) ICharRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharRangeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharRangeContext)
}

func (s *PosCharGroupContext) AllCharClassEsc() []ICharClassEscContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharClassEscContext)(nil)).Elem())
	var tst = make([]ICharClassEscContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharClassEscContext)
		}
	}

	return tst
}

func (s *PosCharGroupContext) CharClassEsc(i int) ICharClassEscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharClassEscContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharClassEscContext)
}

func (s *PosCharGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosCharGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PosCharGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterPosCharGroup(s)
	}
}

func (s *PosCharGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitPosCharGroup(s)
	}
}

func (p *regexParser) PosCharGroup() (localctx IPosCharGroupContext) {
	this := p
	_ = this

	localctx = NewPosCharGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, regexParserRULE_posCharGroup)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == regexParserDASH {
		{
			p.SetState(116)
			p.Match(regexParserDASH)
		}

	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(regexParserSingleCharEsc-10))|(1<<(regexParserMultiCharEsc-10))|(1<<(regexParserCatEsc-10))|(1<<(regexParserComplEsc-10))|(1<<(regexParserNestedSingleCharEsc-10))|(1<<(regexParserNestedMultiCharEsc-10))|(1<<(regexParserNestedCatEsc-10))|(1<<(regexParserNestedComplEsc-10))|(1<<(regexParserXmlChar-10)))) != 0) {
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(119)
				p.CharRange()
			}

		case 2:
			{
				p.SetState(120)
				p.CharClassEsc()
			}

		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICharRangeContext is an interface to support dynamic dispatch.
type ICharRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharRangeContext differentiates from other interfaces.
	IsCharRangeContext()
}

type CharRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharRangeContext() *CharRangeContext {
	var p = new(CharRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charRange
	return p
}

func (*CharRangeContext) IsCharRangeContext() {}

func NewCharRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharRangeContext {
	var p = new(CharRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charRange

	return p
}

func (s *CharRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *CharRangeContext) SeRange() ISeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeRangeContext)
}

func (s *CharRangeContext) XmlChar() antlr.TerminalNode {
	return s.GetToken(regexParserXmlChar, 0)
}

func (s *CharRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharRange(s)
	}
}

func (s *CharRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharRange(s)
	}
}

func (p *regexParser) CharRange() (localctx ICharRangeContext) {
	this := p
	_ = this

	localctx = NewCharRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, regexParserRULE_charRange)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.SeRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(regexParserXmlChar)
		}

	}

	return localctx
}

// ISeRangeContext is an interface to support dynamic dispatch.
type ISeRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeRangeContext differentiates from other interfaces.
	IsSeRangeContext()
}

type SeRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeRangeContext() *SeRangeContext {
	var p = new(SeRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_seRange
	return p
}

func (*SeRangeContext) IsSeRangeContext() {}

func NewSeRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeRangeContext {
	var p = new(SeRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_seRange

	return p
}

func (s *SeRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *SeRangeContext) AllCharOrEsc() []ICharOrEscContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharOrEscContext)(nil)).Elem())
	var tst = make([]ICharOrEscContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharOrEscContext)
		}
	}

	return tst
}

func (s *SeRangeContext) CharOrEsc(i int) ICharOrEscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharOrEscContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharOrEscContext)
}

func (s *SeRangeContext) DASH() antlr.TerminalNode {
	return s.GetToken(regexParserDASH, 0)
}

func (s *SeRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterSeRange(s)
	}
}

func (s *SeRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitSeRange(s)
	}
}

func (p *regexParser) SeRange() (localctx ISeRangeContext) {
	this := p
	_ = this

	localctx = NewSeRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, regexParserRULE_seRange)

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
		p.CharOrEsc()
	}
	{
		p.SetState(130)
		p.Match(regexParserDASH)
	}
	{
		p.SetState(131)
		p.CharOrEsc()
	}

	return localctx
}

// ICharOrEscContext is an interface to support dynamic dispatch.
type ICharOrEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharOrEscContext differentiates from other interfaces.
	IsCharOrEscContext()
}

type CharOrEscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharOrEscContext() *CharOrEscContext {
	var p = new(CharOrEscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charOrEsc
	return p
}

func (*CharOrEscContext) IsCharOrEscContext() {}

func NewCharOrEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharOrEscContext {
	var p = new(CharOrEscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charOrEsc

	return p
}

func (s *CharOrEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CharOrEscContext) XmlChar() antlr.TerminalNode {
	return s.GetToken(regexParserXmlChar, 0)
}

func (s *CharOrEscContext) SingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserSingleCharEsc, 0)
}

func (s *CharOrEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharOrEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharOrEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharOrEsc(s)
	}
}

func (s *CharOrEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharOrEsc(s)
	}
}

func (p *regexParser) CharOrEsc() (localctx ICharOrEscContext) {
	this := p
	_ = this

	localctx = NewCharOrEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, regexParserRULE_charOrEsc)
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
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserSingleCharEsc || _la == regexParserXmlChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICharClassEscContext is an interface to support dynamic dispatch.
type ICharClassEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharClassEscContext differentiates from other interfaces.
	IsCharClassEscContext()
}

type CharClassEscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassEscContext() *CharClassEscContext {
	var p = new(CharClassEscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charClassEsc
	return p
}

func (*CharClassEscContext) IsCharClassEscContext() {}

func NewCharClassEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassEscContext {
	var p = new(CharClassEscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClassEsc

	return p
}

func (s *CharClassEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassEscContext) SingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserSingleCharEsc, 0)
}

func (s *CharClassEscContext) NestedSingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedSingleCharEsc, 0)
}

func (s *CharClassEscContext) MultiCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserMultiCharEsc, 0)
}

func (s *CharClassEscContext) NestedMultiCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedMultiCharEsc, 0)
}

func (s *CharClassEscContext) CatEsc() ICatEscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatEscContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatEscContext)
}

func (s *CharClassEscContext) ComplEsc() IComplEscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplEscContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplEscContext)
}

func (s *CharClassEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClassEsc(s)
	}
}

func (s *CharClassEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClassEsc(s)
	}
}

func (p *regexParser) CharClassEsc() (localctx ICharClassEscContext) {
	this := p
	_ = this

	localctx = NewCharClassEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, regexParserRULE_charClassEsc)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case regexParserSingleCharEsc:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(regexParserSingleCharEsc)
		}

	case regexParserNestedSingleCharEsc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(regexParserNestedSingleCharEsc)
		}

	case regexParserMultiCharEsc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(regexParserMultiCharEsc)
		}

	case regexParserNestedMultiCharEsc:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Match(regexParserNestedMultiCharEsc)
		}

	case regexParserCatEsc, regexParserNestedCatEsc:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.CatEsc()
		}

	case regexParserComplEsc, regexParserNestedComplEsc:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.ComplEsc()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICatEscContext is an interface to support dynamic dispatch.
type ICatEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatEscContext differentiates from other interfaces.
	IsCatEscContext()
}

type CatEscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatEscContext() *CatEscContext {
	var p = new(CatEscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_catEsc
	return p
}

func (*CatEscContext) IsCatEscContext() {}

func NewCatEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatEscContext {
	var p = new(CatEscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_catEsc

	return p
}

func (s *CatEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CatEscContext) CharProp() ICharPropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharPropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharPropContext)
}

func (s *CatEscContext) EndCategory() antlr.TerminalNode {
	return s.GetToken(regexParserEndCategory, 0)
}

func (s *CatEscContext) CatEsc() antlr.TerminalNode {
	return s.GetToken(regexParserCatEsc, 0)
}

func (s *CatEscContext) NestedCatEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedCatEsc, 0)
}

func (s *CatEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCatEsc(s)
	}
}

func (s *CatEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCatEsc(s)
	}
}

func (p *regexParser) CatEsc() (localctx ICatEscContext) {
	this := p
	_ = this

	localctx = NewCatEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, regexParserRULE_catEsc)
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
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserCatEsc || _la == regexParserNestedCatEsc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(144)
		p.CharProp()
	}
	{
		p.SetState(145)
		p.Match(regexParserEndCategory)
	}

	return localctx
}

// IComplEscContext is an interface to support dynamic dispatch.
type IComplEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComplEscContext differentiates from other interfaces.
	IsComplEscContext()
}

type ComplEscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComplEscContext() *ComplEscContext {
	var p = new(ComplEscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_complEsc
	return p
}

func (*ComplEscContext) IsComplEscContext() {}

func NewComplEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComplEscContext {
	var p = new(ComplEscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_complEsc

	return p
}

func (s *ComplEscContext) GetParser() antlr.Parser { return s.parser }

func (s *ComplEscContext) CharProp() ICharPropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharPropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharPropContext)
}

func (s *ComplEscContext) EndCategory() antlr.TerminalNode {
	return s.GetToken(regexParserEndCategory, 0)
}

func (s *ComplEscContext) ComplEsc() antlr.TerminalNode {
	return s.GetToken(regexParserComplEsc, 0)
}

func (s *ComplEscContext) NestedComplEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedComplEsc, 0)
}

func (s *ComplEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComplEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComplEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterComplEsc(s)
	}
}

func (s *ComplEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitComplEsc(s)
	}
}

func (p *regexParser) ComplEsc() (localctx IComplEscContext) {
	this := p
	_ = this

	localctx = NewComplEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, regexParserRULE_complEsc)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserComplEsc || _la == regexParserNestedComplEsc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(148)
		p.CharProp()
	}
	{
		p.SetState(149)
		p.Match(regexParserEndCategory)
	}

	return localctx
}

// ICharPropContext is an interface to support dynamic dispatch.
type ICharPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharPropContext differentiates from other interfaces.
	IsCharPropContext()
}

type CharPropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharPropContext() *CharPropContext {
	var p = new(CharPropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = regexParserRULE_charProp
	return p
}

func (*CharPropContext) IsCharPropContext() {}

func NewCharPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharPropContext {
	var p = new(CharPropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charProp

	return p
}

func (s *CharPropContext) GetParser() antlr.Parser { return s.parser }

func (s *CharPropContext) IsCategory() antlr.TerminalNode {
	return s.GetToken(regexParserIsCategory, 0)
}

func (s *CharPropContext) IsBlock() antlr.TerminalNode {
	return s.GetToken(regexParserIsBlock, 0)
}

func (s *CharPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharProp(s)
	}
}

func (s *CharPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharProp(s)
	}
}

func (p *regexParser) CharProp() (localctx ICharPropContext) {
	this := p
	_ = this

	localctx = NewCharPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, regexParserRULE_charProp)
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
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserIsCategory || _la == regexParserIsBlock) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
