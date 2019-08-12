// Code generated from CMake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cmake // CMake
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 42, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 23, 10,
	3, 12, 3, 14, 3, 26, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7,
	5, 35, 10, 5, 12, 5, 14, 5, 38, 11, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4,
	6, 8, 2, 3, 4, 2, 5, 6, 8, 9, 2, 42, 2, 13, 3, 2, 2, 2, 4, 18, 3, 2, 2,
	2, 6, 29, 3, 2, 2, 2, 8, 31, 3, 2, 2, 2, 10, 12, 5, 4, 3, 2, 11, 10, 3,
	2, 2, 2, 12, 15, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14,
	16, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 16, 17, 7, 2, 2, 3, 17, 3, 3, 2, 2,
	2, 18, 19, 7, 5, 2, 2, 19, 24, 7, 3, 2, 2, 20, 23, 5, 6, 4, 2, 21, 23,
	5, 8, 5, 2, 22, 20, 3, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2,
	24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24, 3,
	2, 2, 2, 27, 28, 7, 4, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 9, 2, 2, 2, 30,
	7, 3, 2, 2, 2, 31, 36, 7, 3, 2, 2, 32, 35, 5, 6, 4, 2, 33, 35, 5, 8, 5,
	2, 34, 32, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34,
	3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2,
	39, 40, 7, 4, 2, 2, 40, 9, 3, 2, 2, 2, 7, 13, 22, 24, 34, 36,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "Identifier", "Unquoted_argument", "Escape_sequence", "Quoted_argument",
	"Bracket_argument", "Bracket_comment", "Line_comment", "Newline", "Space",
}

var ruleNames = []string{
	"file", "command_invocation", "single_argument", "compound_argument",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CMakeParser struct {
	*antlr.BaseParser
}

func NewCMakeParser(input antlr.TokenStream) *CMakeParser {
	this := new(CMakeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CMake.g4"

	return this
}

// CMakeParser tokens.
const (
	CMakeParserEOF               = antlr.TokenEOF
	CMakeParserT__0              = 1
	CMakeParserT__1              = 2
	CMakeParserIdentifier        = 3
	CMakeParserUnquoted_argument = 4
	CMakeParserEscape_sequence   = 5
	CMakeParserQuoted_argument   = 6
	CMakeParserBracket_argument  = 7
	CMakeParserBracket_comment   = 8
	CMakeParserLine_comment      = 9
	CMakeParserNewline           = 10
	CMakeParserSpace             = 11
)

// CMakeParser rules.
const (
	CMakeParserRULE_file               = 0
	CMakeParserRULE_command_invocation = 1
	CMakeParserRULE_single_argument    = 2
	CMakeParserRULE_compound_argument  = 3
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMakeParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMakeParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(CMakeParserEOF, 0)
}

func (s *FileContext) AllCommand_invocation() []ICommand_invocationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommand_invocationContext)(nil)).Elem())
	var tst = make([]ICommand_invocationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommand_invocationContext)
		}
	}

	return tst
}

func (s *FileContext) Command_invocation(i int) ICommand_invocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommand_invocationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommand_invocationContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *CMakeParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CMakeParserRULE_file)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CMakeParserIdentifier {
		{
			p.SetState(8)
			p.Command_invocation()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(CMakeParserEOF)
	}

	return localctx
}

// ICommand_invocationContext is an interface to support dynamic dispatch.
type ICommand_invocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommand_invocationContext differentiates from other interfaces.
	IsCommand_invocationContext()
}

type Command_invocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_invocationContext() *Command_invocationContext {
	var p = new(Command_invocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMakeParserRULE_command_invocation
	return p
}

func (*Command_invocationContext) IsCommand_invocationContext() {}

func NewCommand_invocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_invocationContext {
	var p = new(Command_invocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMakeParserRULE_command_invocation

	return p
}

func (s *Command_invocationContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_invocationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CMakeParserIdentifier, 0)
}

func (s *Command_invocationContext) AllSingle_argument() []ISingle_argumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingle_argumentContext)(nil)).Elem())
	var tst = make([]ISingle_argumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingle_argumentContext)
		}
	}

	return tst
}

func (s *Command_invocationContext) Single_argument(i int) ISingle_argumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_argumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingle_argumentContext)
}

func (s *Command_invocationContext) AllCompound_argument() []ICompound_argumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompound_argumentContext)(nil)).Elem())
	var tst = make([]ICompound_argumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompound_argumentContext)
		}
	}

	return tst
}

func (s *Command_invocationContext) Compound_argument(i int) ICompound_argumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_argumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompound_argumentContext)
}

func (s *Command_invocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_invocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Command_invocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.EnterCommand_invocation(s)
	}
}

func (s *Command_invocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.ExitCommand_invocation(s)
	}
}

func (p *CMakeParser) Command_invocation() (localctx ICommand_invocationContext) {
	localctx = NewCommand_invocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CMakeParserRULE_command_invocation)
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
		p.SetState(16)
		p.Match(CMakeParserIdentifier)
	}
	{
		p.SetState(17)
		p.Match(CMakeParserT__0)
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CMakeParserT__0)|(1<<CMakeParserIdentifier)|(1<<CMakeParserUnquoted_argument)|(1<<CMakeParserQuoted_argument)|(1<<CMakeParserBracket_argument))) != 0 {
		p.SetState(20)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CMakeParserIdentifier, CMakeParserUnquoted_argument, CMakeParserQuoted_argument, CMakeParserBracket_argument:
			{
				p.SetState(18)
				p.Single_argument()
			}

		case CMakeParserT__0:
			{
				p.SetState(19)
				p.Compound_argument()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(CMakeParserT__1)
	}

	return localctx
}

// ISingle_argumentContext is an interface to support dynamic dispatch.
type ISingle_argumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_argumentContext differentiates from other interfaces.
	IsSingle_argumentContext()
}

type Single_argumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_argumentContext() *Single_argumentContext {
	var p = new(Single_argumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMakeParserRULE_single_argument
	return p
}

func (*Single_argumentContext) IsSingle_argumentContext() {}

func NewSingle_argumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_argumentContext {
	var p = new(Single_argumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMakeParserRULE_single_argument

	return p
}

func (s *Single_argumentContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_argumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CMakeParserIdentifier, 0)
}

func (s *Single_argumentContext) Unquoted_argument() antlr.TerminalNode {
	return s.GetToken(CMakeParserUnquoted_argument, 0)
}

func (s *Single_argumentContext) Bracket_argument() antlr.TerminalNode {
	return s.GetToken(CMakeParserBracket_argument, 0)
}

func (s *Single_argumentContext) Quoted_argument() antlr.TerminalNode {
	return s.GetToken(CMakeParserQuoted_argument, 0)
}

func (s *Single_argumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_argumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_argumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.EnterSingle_argument(s)
	}
}

func (s *Single_argumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.ExitSingle_argument(s)
	}
}

func (p *CMakeParser) Single_argument() (localctx ISingle_argumentContext) {
	localctx = NewSingle_argumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CMakeParserRULE_single_argument)
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
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CMakeParserIdentifier)|(1<<CMakeParserUnquoted_argument)|(1<<CMakeParserQuoted_argument)|(1<<CMakeParserBracket_argument))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompound_argumentContext is an interface to support dynamic dispatch.
type ICompound_argumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompound_argumentContext differentiates from other interfaces.
	IsCompound_argumentContext()
}

type Compound_argumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_argumentContext() *Compound_argumentContext {
	var p = new(Compound_argumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CMakeParserRULE_compound_argument
	return p
}

func (*Compound_argumentContext) IsCompound_argumentContext() {}

func NewCompound_argumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_argumentContext {
	var p = new(Compound_argumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CMakeParserRULE_compound_argument

	return p
}

func (s *Compound_argumentContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_argumentContext) AllSingle_argument() []ISingle_argumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingle_argumentContext)(nil)).Elem())
	var tst = make([]ISingle_argumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingle_argumentContext)
		}
	}

	return tst
}

func (s *Compound_argumentContext) Single_argument(i int) ISingle_argumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_argumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingle_argumentContext)
}

func (s *Compound_argumentContext) AllCompound_argument() []ICompound_argumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompound_argumentContext)(nil)).Elem())
	var tst = make([]ICompound_argumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompound_argumentContext)
		}
	}

	return tst
}

func (s *Compound_argumentContext) Compound_argument(i int) ICompound_argumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_argumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompound_argumentContext)
}

func (s *Compound_argumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_argumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_argumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.EnterCompound_argument(s)
	}
}

func (s *Compound_argumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CMakeListener); ok {
		listenerT.ExitCompound_argument(s)
	}
}

func (p *CMakeParser) Compound_argument() (localctx ICompound_argumentContext) {
	localctx = NewCompound_argumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CMakeParserRULE_compound_argument)
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
		p.SetState(29)
		p.Match(CMakeParserT__0)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CMakeParserT__0)|(1<<CMakeParserIdentifier)|(1<<CMakeParserUnquoted_argument)|(1<<CMakeParserQuoted_argument)|(1<<CMakeParserBracket_argument))) != 0 {
		p.SetState(32)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CMakeParserIdentifier, CMakeParserUnquoted_argument, CMakeParserQuoted_argument, CMakeParserBracket_argument:
			{
				p.SetState(30)
				p.Single_argument()
			}

		case CMakeParserT__0:
			{
				p.SetState(31)
				p.Compound_argument()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(CMakeParserT__1)
	}

	return localctx
}
