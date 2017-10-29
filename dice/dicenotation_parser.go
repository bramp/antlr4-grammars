// Generated from DiceNotation.g4 by ANTLR 4.7.

package dice // DiceNotation
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 31, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 5, 3, 16, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 25,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 2, 2, 29,
	2, 10, 3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 26, 3, 2, 2,
	2, 10, 11, 5, 4, 3, 2, 11, 3, 3, 2, 2, 2, 12, 16, 5, 8, 5, 2, 13, 16, 5,
	6, 4, 2, 14, 16, 7, 7, 2, 2, 15, 12, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15,
	14, 3, 2, 2, 2, 16, 5, 3, 2, 2, 2, 17, 18, 5, 8, 5, 2, 18, 19, 7, 3, 2,
	2, 19, 20, 5, 4, 3, 2, 20, 25, 3, 2, 2, 2, 21, 22, 7, 7, 2, 2, 22, 23,
	7, 3, 2, 2, 23, 25, 5, 4, 3, 2, 24, 17, 3, 2, 2, 2, 24, 21, 3, 2, 2, 2,
	25, 7, 3, 2, 2, 2, 26, 27, 7, 7, 2, 2, 27, 28, 7, 6, 2, 2, 28, 29, 7, 7,
	2, 2, 29, 9, 3, 2, 2, 2, 4, 15, 24,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "OPERATOR", "ADD", "SUB", "DSEPARATOR", "DIGIT", "WS",
}

var ruleNames = []string{
	"parse", "function", "binaryOp", "dice",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DiceNotationParser struct {
	*antlr.BaseParser
}

func NewDiceNotationParser(input antlr.TokenStream) *DiceNotationParser {
	this := new(DiceNotationParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DiceNotation.g4"

	return this
}

// DiceNotationParser tokens.
const (
	DiceNotationParserEOF        = antlr.TokenEOF
	DiceNotationParserOPERATOR   = 1
	DiceNotationParserADD        = 2
	DiceNotationParserSUB        = 3
	DiceNotationParserDSEPARATOR = 4
	DiceNotationParserDIGIT      = 5
	DiceNotationParserWS         = 6
)

// DiceNotationParser rules.
const (
	DiceNotationParserRULE_parse    = 0
	DiceNotationParserRULE_function = 1
	DiceNotationParserRULE_binaryOp = 2
	DiceNotationParserRULE_dice     = 3
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *DiceNotationParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiceNotationParserRULE_parse)

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
		p.SetState(8)
		p.Function()
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Dice() IDiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *FunctionContext) BinaryOp() IBinaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *FunctionContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDIGIT, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *DiceNotationParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiceNotationParserRULE_function)

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

	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Dice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.BinaryOp()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Match(DiceNotationParserDIGIT)
		}

	}

	return localctx
}

// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOpContext differentiates from other interfaces.
	IsBinaryOpContext()
}

type BinaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpContext() *BinaryOpContext {
	var p = new(BinaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_binaryOp
	return p
}

func (*BinaryOpContext) IsBinaryOpContext() {}

func NewBinaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_binaryOp

	return p
}

func (s *BinaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpContext) Dice() IDiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *BinaryOpContext) OPERATOR() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserOPERATOR, 0)
}

func (s *BinaryOpContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *BinaryOpContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDIGIT, 0)
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.EnterBinaryOp(s)
	}
}

func (s *BinaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.ExitBinaryOp(s)
	}
}

func (p *DiceNotationParser) BinaryOp() (localctx IBinaryOpContext) {
	localctx = NewBinaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DiceNotationParserRULE_binaryOp)

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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Dice()
		}
		{
			p.SetState(16)
			p.Match(DiceNotationParserOPERATOR)
		}
		{
			p.SetState(17)
			p.Function()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Match(DiceNotationParserDIGIT)
		}
		{
			p.SetState(20)
			p.Match(DiceNotationParserOPERATOR)
		}
		{
			p.SetState(21)
			p.Function()
		}

	}

	return localctx
}

// IDiceContext is an interface to support dynamic dispatch.
type IDiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiceContext differentiates from other interfaces.
	IsDiceContext()
}

type DiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiceContext() *DiceContext {
	var p = new(DiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_dice
	return p
}

func (*DiceContext) IsDiceContext() {}

func NewDiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiceContext {
	var p = new(DiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_dice

	return p
}

func (s *DiceContext) GetParser() antlr.Parser { return s.parser }

func (s *DiceContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(DiceNotationParserDIGIT)
}

func (s *DiceContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDIGIT, i)
}

func (s *DiceContext) DSEPARATOR() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDSEPARATOR, 0)
}

func (s *DiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.EnterDice(s)
	}
}

func (s *DiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationListener); ok {
		listenerT.ExitDice(s)
	}
}

func (p *DiceNotationParser) Dice() (localctx IDiceContext) {
	localctx = NewDiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DiceNotationParserRULE_dice)

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
		p.SetState(24)
		p.Match(DiceNotationParserDIGIT)
	}
	{
		p.SetState(25)
		p.Match(DiceNotationParserDSEPARATOR)
	}
	{
		p.SetState(26)
		p.Match(DiceNotationParserDIGIT)
	}

	return localctx
}
