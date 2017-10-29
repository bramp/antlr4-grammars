// Generated from postalcode.g4 by ANTLR 4.7.

package postalcode // postalcode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 12, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2,
	2, 2, 10, 2, 4, 3, 2, 2, 2, 4, 5, 7, 4, 2, 2, 5, 6, 7, 3, 2, 2, 6, 7, 7,
	4, 2, 2, 7, 8, 7, 3, 2, 2, 8, 9, 7, 4, 2, 2, 9, 10, 7, 3, 2, 2, 10, 3,
	3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames []string

var symbolicNames = []string{
	"", "DIGIT", "LETTER", "WS",
}

var ruleNames = []string{
	"postalcode",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type postalcodeParser struct {
	*antlr.BaseParser
}

func NewpostalcodeParser(input antlr.TokenStream) *postalcodeParser {
	this := new(postalcodeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "postalcode.g4"

	return this
}

// postalcodeParser tokens.
const (
	postalcodeParserEOF    = antlr.TokenEOF
	postalcodeParserDIGIT  = 1
	postalcodeParserLETTER = 2
	postalcodeParserWS     = 3
)

// postalcodeParserRULE_postalcode is the postalcodeParser rule.
const postalcodeParserRULE_postalcode = 0

// IPostalcodeContext is an interface to support dynamic dispatch.
type IPostalcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostalcodeContext differentiates from other interfaces.
	IsPostalcodeContext()
}

type PostalcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostalcodeContext() *PostalcodeContext {
	var p = new(PostalcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = postalcodeParserRULE_postalcode
	return p
}

func (*PostalcodeContext) IsPostalcodeContext() {}

func NewPostalcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostalcodeContext {
	var p = new(PostalcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = postalcodeParserRULE_postalcode

	return p
}

func (s *PostalcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *PostalcodeContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(postalcodeParserLETTER)
}

func (s *PostalcodeContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(postalcodeParserLETTER, i)
}

func (s *PostalcodeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(postalcodeParserDIGIT)
}

func (s *PostalcodeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(postalcodeParserDIGIT, i)
}

func (s *PostalcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostalcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostalcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(postalcodeListener); ok {
		listenerT.EnterPostalcode(s)
	}
}

func (s *PostalcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(postalcodeListener); ok {
		listenerT.ExitPostalcode(s)
	}
}

func (p *postalcodeParser) Postalcode() (localctx IPostalcodeContext) {
	localctx = NewPostalcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, postalcodeParserRULE_postalcode)

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
		p.SetState(2)
		p.Match(postalcodeParserLETTER)
	}
	{
		p.SetState(3)
		p.Match(postalcodeParserDIGIT)
	}
	{
		p.SetState(4)
		p.Match(postalcodeParserLETTER)
	}
	{
		p.SetState(5)
		p.Match(postalcodeParserDIGIT)
	}
	{
		p.SetState(6)
		p.Match(postalcodeParserLETTER)
	}
	{
		p.SetState(7)
		p.Match(postalcodeParserDIGIT)
	}

	return localctx
}
