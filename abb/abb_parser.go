// Code generated from abbParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package abb // abbParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 154,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46,
	11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 61, 10, 5, 12, 5, 14, 5, 64, 11, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 72, 10, 6, 12, 6, 14, 6, 75, 11, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 5, 7, 81, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 87, 10,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 95, 10, 10, 12, 10, 14,
	10, 98, 11, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	107, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 114, 10, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 5, 14, 121, 10, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 5, 15, 127, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 133, 10,
	16, 12, 16, 14, 16, 136, 11, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 5, 17, 145, 10, 17, 3, 17, 3, 17, 5, 17, 149, 10, 17, 3, 17,
	5, 17, 152, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 5, 3, 2, 11, 16, 3, 2, 8, 10, 3, 2, 38, 39,
	2, 159, 2, 34, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 62,
	3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 78, 3, 2, 2, 2, 14, 82, 3, 2, 2, 2,
	16, 84, 3, 2, 2, 2, 18, 90, 3, 2, 2, 2, 20, 106, 3, 2, 2, 2, 22, 108, 3,
	2, 2, 2, 24, 117, 3, 2, 2, 2, 26, 120, 3, 2, 2, 2, 28, 126, 3, 2, 2, 2,
	30, 128, 3, 2, 2, 2, 32, 151, 3, 2, 2, 2, 34, 35, 5, 4, 3, 2, 35, 36, 7,
	2, 2, 3, 36, 3, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5, 6, 4, 2, 39,
	40, 7, 44, 2, 2, 40, 44, 5, 8, 5, 2, 41, 43, 7, 44, 2, 2, 42, 41, 3, 2,
	2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47,
	3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 7, 4, 2, 2, 48, 5, 3, 2, 2, 2,
	49, 52, 7, 51, 2, 2, 50, 52, 5, 12, 7, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3,
	2, 2, 2, 52, 7, 3, 2, 2, 2, 53, 61, 7, 44, 2, 2, 54, 55, 5, 22, 12, 2,
	55, 56, 7, 44, 2, 2, 56, 61, 3, 2, 2, 2, 57, 58, 5, 10, 6, 2, 58, 59, 7,
	44, 2, 2, 59, 61, 3, 2, 2, 2, 60, 53, 3, 2, 2, 2, 60, 54, 3, 2, 2, 2, 60,
	57, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2,
	2, 63, 9, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 7, 5, 2, 2, 66, 67, 5,
	12, 7, 2, 67, 73, 7, 44, 2, 2, 68, 69, 5, 18, 10, 2, 69, 70, 7, 44, 2,
	2, 70, 72, 3, 2, 2, 2, 71, 68, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71,
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	76, 77, 7, 6, 2, 2, 77, 11, 3, 2, 2, 2, 78, 80, 5, 14, 8, 2, 79, 81, 5,
	16, 9, 2, 80, 79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 13, 3, 2, 2, 2, 82,
	83, 7, 51, 2, 2, 83, 15, 3, 2, 2, 2, 84, 86, 7, 26, 2, 2, 85, 87, 7, 51,
	2, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89,
	7, 27, 2, 2, 89, 17, 3, 2, 2, 2, 90, 96, 7, 51, 2, 2, 91, 92, 5, 20, 11,
	2, 92, 93, 7, 21, 2, 2, 93, 95, 3, 2, 2, 2, 94, 91, 3, 2, 2, 2, 95, 98,
	3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2,
	98, 96, 3, 2, 2, 2, 99, 100, 5, 20, 11, 2, 100, 101, 7, 25, 2, 2, 101,
	19, 3, 2, 2, 2, 102, 107, 7, 17, 2, 2, 103, 107, 7, 18, 2, 2, 104, 107,
	5, 32, 17, 2, 105, 107, 7, 51, 2, 2, 106, 102, 3, 2, 2, 2, 106, 103, 3,
	2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 21, 3, 2, 2,
	2, 108, 109, 5, 26, 14, 2, 109, 110, 5, 24, 13, 2, 110, 113, 7, 51, 2,
	2, 111, 112, 7, 20, 2, 2, 112, 114, 5, 28, 15, 2, 113, 111, 3, 2, 2, 2,
	113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 7, 25, 2, 2, 116,
	23, 3, 2, 2, 2, 117, 118, 9, 2, 2, 2, 118, 25, 3, 2, 2, 2, 119, 121, 7,
	7, 2, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2,
	2, 122, 123, 9, 3, 2, 2, 123, 27, 3, 2, 2, 2, 124, 127, 5, 30, 16, 2, 125,
	127, 5, 32, 17, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 29,
	3, 2, 2, 2, 128, 134, 7, 28, 2, 2, 129, 130, 5, 28, 15, 2, 130, 131, 7,
	21, 2, 2, 131, 133, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 133, 136, 3, 2, 2,
	2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 137, 138, 5, 28, 15, 2, 138, 139, 7, 29, 2, 2, 139, 31,
	3, 2, 2, 2, 140, 152, 7, 46, 2, 2, 141, 152, 7, 47, 2, 2, 142, 152, 7,
	48, 2, 2, 143, 145, 9, 4, 2, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2,
	2, 145, 146, 3, 2, 2, 2, 146, 152, 7, 49, 2, 2, 147, 149, 9, 4, 2, 2, 148,
	147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152,
	7, 50, 2, 2, 151, 140, 3, 2, 2, 2, 151, 141, 3, 2, 2, 2, 151, 142, 3, 2,
	2, 2, 151, 144, 3, 2, 2, 2, 151, 148, 3, 2, 2, 2, 152, 33, 3, 2, 2, 2,
	18, 44, 51, 60, 62, 73, 80, 86, 96, 106, 113, 120, 126, 134, 144, 148,
	151,
}
var literalNames = []string{
	"", "'module'", "'endmodule'", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "'/'", "':='", "','", "'{'", "'}'", "':'", "';'", "'('",
	"')'", "'['", "']'", "'.'", "'..'", "'>'", "'>='", "'<'", "'<='", "'=='",
	"'<>'", "'+'", "'-'", "'*'", "'%'", "'#'",
}
var symbolicNames = []string{
	"", "MODULE", "ENDMODULE", "PROC", "ENDPROC", "LOCAL", "CONST", "PERS",
	"VAR", "TOOLDATA", "WOBJDATA", "SPEEDDATA", "ZONEDATA", "CLOCK", "BOOL",
	"ON_CALL", "OFF_CALL", "SLASH", "EQUALS", "COMMA", "CURLY_OPEN", "CURLY_CLOSE",
	"COLON", "SEMICOLON", "BRACKET_OPEN", "BRACKET_CLOSE", "SQUARE_OPEN", "SQUARE_CLOSE",
	"DOT", "DOUBLEDOT", "REL_BIGGER", "REL_BIGGER_OR_EQUAL", "REL_SMALLER",
	"REL_SMALLER_OR_EQUAL", "REL_EQUAL", "REL_NOTEQUAL", "PLUS", "MINUS", "MULTIPLY",
	"PERCENT", "HASH", "WS", "NEWLINE", "LINE_COMMENT", "BOOLLITERAL", "CHARLITERAL",
	"STRINGLITERAL", "FLOATLITERAL", "INTLITERAL", "IDENTIFIER",
}

var ruleNames = []string{
	"module", "moduleData", "moduleName", "dataList", "procedure", "procCall",
	"procName", "procParameter", "functionCall", "functionParameter", "declaration",
	"type_", "init_", "expression", "array_", "primitive",
}

type abbParser struct {
	*antlr.BaseParser
}

// NewabbParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *abbParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewabbParser(input antlr.TokenStream) *abbParser {
	this := new(abbParser)
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
	this.GrammarFileName = "abbParser.g4"

	return this
}

// abbParser tokens.
const (
	abbParserEOF                  = antlr.TokenEOF
	abbParserMODULE               = 1
	abbParserENDMODULE            = 2
	abbParserPROC                 = 3
	abbParserENDPROC              = 4
	abbParserLOCAL                = 5
	abbParserCONST                = 6
	abbParserPERS                 = 7
	abbParserVAR                  = 8
	abbParserTOOLDATA             = 9
	abbParserWOBJDATA             = 10
	abbParserSPEEDDATA            = 11
	abbParserZONEDATA             = 12
	abbParserCLOCK                = 13
	abbParserBOOL                 = 14
	abbParserON_CALL              = 15
	abbParserOFF_CALL             = 16
	abbParserSLASH                = 17
	abbParserEQUALS               = 18
	abbParserCOMMA                = 19
	abbParserCURLY_OPEN           = 20
	abbParserCURLY_CLOSE          = 21
	abbParserCOLON                = 22
	abbParserSEMICOLON            = 23
	abbParserBRACKET_OPEN         = 24
	abbParserBRACKET_CLOSE        = 25
	abbParserSQUARE_OPEN          = 26
	abbParserSQUARE_CLOSE         = 27
	abbParserDOT                  = 28
	abbParserDOUBLEDOT            = 29
	abbParserREL_BIGGER           = 30
	abbParserREL_BIGGER_OR_EQUAL  = 31
	abbParserREL_SMALLER          = 32
	abbParserREL_SMALLER_OR_EQUAL = 33
	abbParserREL_EQUAL            = 34
	abbParserREL_NOTEQUAL         = 35
	abbParserPLUS                 = 36
	abbParserMINUS                = 37
	abbParserMULTIPLY             = 38
	abbParserPERCENT              = 39
	abbParserHASH                 = 40
	abbParserWS                   = 41
	abbParserNEWLINE              = 42
	abbParserLINE_COMMENT         = 43
	abbParserBOOLLITERAL          = 44
	abbParserCHARLITERAL          = 45
	abbParserSTRINGLITERAL        = 46
	abbParserFLOATLITERAL         = 47
	abbParserINTLITERAL           = 48
	abbParserIDENTIFIER           = 49
)

// abbParser rules.
const (
	abbParserRULE_module            = 0
	abbParserRULE_moduleData        = 1
	abbParserRULE_moduleName        = 2
	abbParserRULE_dataList          = 3
	abbParserRULE_procedure         = 4
	abbParserRULE_procCall          = 5
	abbParserRULE_procName          = 6
	abbParserRULE_procParameter     = 7
	abbParserRULE_functionCall      = 8
	abbParserRULE_functionParameter = 9
	abbParserRULE_declaration       = 10
	abbParserRULE_type_             = 11
	abbParserRULE_init_             = 12
	abbParserRULE_expression        = 13
	abbParserRULE_array_            = 14
	abbParserRULE_primitive         = 15
)

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
	p.RuleIndex = abbParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) ModuleData() IModuleDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDataContext)
}

func (s *ModuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(abbParserEOF, 0)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *abbParser) Module() (localctx IModuleContext) {
	this := p
	_ = this

	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, abbParserRULE_module)

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
		p.ModuleData()
	}
	{
		p.SetState(33)
		p.Match(abbParserEOF)
	}

	return localctx
}

// IModuleDataContext is an interface to support dynamic dispatch.
type IModuleDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDataContext differentiates from other interfaces.
	IsModuleDataContext()
}

type ModuleDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDataContext() *ModuleDataContext {
	var p = new(ModuleDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_moduleData
	return p
}

func (*ModuleDataContext) IsModuleDataContext() {}

func NewModuleDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDataContext {
	var p = new(ModuleDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_moduleData

	return p
}

func (s *ModuleDataContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDataContext) MODULE() antlr.TerminalNode {
	return s.GetToken(abbParserMODULE, 0)
}

func (s *ModuleDataContext) ModuleName() IModuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleNameContext)
}

func (s *ModuleDataContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(abbParserNEWLINE)
}

func (s *ModuleDataContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(abbParserNEWLINE, i)
}

func (s *ModuleDataContext) DataList() IDataListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataListContext)
}

func (s *ModuleDataContext) ENDMODULE() antlr.TerminalNode {
	return s.GetToken(abbParserENDMODULE, 0)
}

func (s *ModuleDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterModuleData(s)
	}
}

func (s *ModuleDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitModuleData(s)
	}
}

func (p *abbParser) ModuleData() (localctx IModuleDataContext) {
	this := p
	_ = this

	localctx = NewModuleDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, abbParserRULE_moduleData)
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
		p.SetState(35)
		p.Match(abbParserMODULE)
	}
	{
		p.SetState(36)
		p.ModuleName()
	}
	{
		p.SetState(37)
		p.Match(abbParserNEWLINE)
	}
	{
		p.SetState(38)
		p.DataList()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == abbParserNEWLINE {
		{
			p.SetState(39)
			p.Match(abbParserNEWLINE)
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(abbParserENDMODULE)
	}

	return localctx
}

// IModuleNameContext is an interface to support dynamic dispatch.
type IModuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleNameContext differentiates from other interfaces.
	IsModuleNameContext()
}

type ModuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleNameContext() *ModuleNameContext {
	var p = new(ModuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_moduleName
	return p
}

func (*ModuleNameContext) IsModuleNameContext() {}

func NewModuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleNameContext {
	var p = new(ModuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_moduleName

	return p
}

func (s *ModuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *ModuleNameContext) ProcCall() IProcCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcCallContext)
}

func (s *ModuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterModuleName(s)
	}
}

func (s *ModuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitModuleName(s)
	}
}

func (p *abbParser) ModuleName() (localctx IModuleNameContext) {
	this := p
	_ = this

	localctx = NewModuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, abbParserRULE_moduleName)

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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(abbParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.ProcCall()
		}

	}

	return localctx
}

// IDataListContext is an interface to support dynamic dispatch.
type IDataListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataListContext differentiates from other interfaces.
	IsDataListContext()
}

type DataListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataListContext() *DataListContext {
	var p = new(DataListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_dataList
	return p
}

func (*DataListContext) IsDataListContext() {}

func NewDataListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataListContext {
	var p = new(DataListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_dataList

	return p
}

func (s *DataListContext) GetParser() antlr.Parser { return s.parser }

func (s *DataListContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(abbParserNEWLINE)
}

func (s *DataListContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(abbParserNEWLINE, i)
}

func (s *DataListContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *DataListContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DataListContext) AllProcedure() []IProcedureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProcedureContext)(nil)).Elem())
	var tst = make([]IProcedureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProcedureContext)
		}
	}

	return tst
}

func (s *DataListContext) Procedure(i int) IProcedureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProcedureContext)
}

func (s *DataListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterDataList(s)
	}
}

func (s *DataListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitDataList(s)
	}
}

func (p *abbParser) DataList() (localctx IDataListContext) {
	this := p
	_ = this

	localctx = NewDataListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, abbParserRULE_dataList)

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(58)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case abbParserNEWLINE:
				{
					p.SetState(51)
					p.Match(abbParserNEWLINE)
				}

			case abbParserLOCAL, abbParserCONST, abbParserPERS, abbParserVAR:
				{
					p.SetState(52)
					p.Declaration()
				}
				{
					p.SetState(53)
					p.Match(abbParserNEWLINE)
				}

			case abbParserPROC:
				{
					p.SetState(55)
					p.Procedure()
				}
				{
					p.SetState(56)
					p.Match(abbParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = abbParserRULE_procedure
	return p
}

func (*ProcedureContext) IsProcedureContext() {}

func NewProcedureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureContext {
	var p = new(ProcedureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_procedure

	return p
}

func (s *ProcedureContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureContext) PROC() antlr.TerminalNode {
	return s.GetToken(abbParserPROC, 0)
}

func (s *ProcedureContext) ProcCall() IProcCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcCallContext)
}

func (s *ProcedureContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(abbParserNEWLINE)
}

func (s *ProcedureContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(abbParserNEWLINE, i)
}

func (s *ProcedureContext) ENDPROC() antlr.TerminalNode {
	return s.GetToken(abbParserENDPROC, 0)
}

func (s *ProcedureContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *ProcedureContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ProcedureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterProcedure(s)
	}
}

func (s *ProcedureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitProcedure(s)
	}
}

func (p *abbParser) Procedure() (localctx IProcedureContext) {
	this := p
	_ = this

	localctx = NewProcedureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, abbParserRULE_procedure)
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
		p.Match(abbParserPROC)
	}
	{
		p.SetState(64)
		p.ProcCall()
	}
	{
		p.SetState(65)
		p.Match(abbParserNEWLINE)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == abbParserIDENTIFIER {
		{
			p.SetState(66)
			p.FunctionCall()
		}
		{
			p.SetState(67)
			p.Match(abbParserNEWLINE)
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(abbParserENDPROC)
	}

	return localctx
}

// IProcCallContext is an interface to support dynamic dispatch.
type IProcCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcCallContext differentiates from other interfaces.
	IsProcCallContext()
}

type ProcCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcCallContext() *ProcCallContext {
	var p = new(ProcCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_procCall
	return p
}

func (*ProcCallContext) IsProcCallContext() {}

func NewProcCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcCallContext {
	var p = new(ProcCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_procCall

	return p
}

func (s *ProcCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcCallContext) ProcName() IProcNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcNameContext)
}

func (s *ProcCallContext) ProcParameter() IProcParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcParameterContext)
}

func (s *ProcCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterProcCall(s)
	}
}

func (s *ProcCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitProcCall(s)
	}
}

func (p *abbParser) ProcCall() (localctx IProcCallContext) {
	this := p
	_ = this

	localctx = NewProcCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, abbParserRULE_procCall)
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
		p.SetState(76)
		p.ProcName()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == abbParserBRACKET_OPEN {
		{
			p.SetState(77)
			p.ProcParameter()
		}

	}

	return localctx
}

// IProcNameContext is an interface to support dynamic dispatch.
type IProcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcNameContext differentiates from other interfaces.
	IsProcNameContext()
}

type ProcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcNameContext() *ProcNameContext {
	var p = new(ProcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_procName
	return p
}

func (*ProcNameContext) IsProcNameContext() {}

func NewProcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcNameContext {
	var p = new(ProcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_procName

	return p
}

func (s *ProcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *ProcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterProcName(s)
	}
}

func (s *ProcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitProcName(s)
	}
}

func (p *abbParser) ProcName() (localctx IProcNameContext) {
	this := p
	_ = this

	localctx = NewProcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, abbParserRULE_procName)

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
		p.SetState(80)
		p.Match(abbParserIDENTIFIER)
	}

	return localctx
}

// IProcParameterContext is an interface to support dynamic dispatch.
type IProcParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcParameterContext differentiates from other interfaces.
	IsProcParameterContext()
}

type ProcParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcParameterContext() *ProcParameterContext {
	var p = new(ProcParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_procParameter
	return p
}

func (*ProcParameterContext) IsProcParameterContext() {}

func NewProcParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcParameterContext {
	var p = new(ProcParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_procParameter

	return p
}

func (s *ProcParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcParameterContext) BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(abbParserBRACKET_OPEN, 0)
}

func (s *ProcParameterContext) BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(abbParserBRACKET_CLOSE, 0)
}

func (s *ProcParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *ProcParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterProcParameter(s)
	}
}

func (s *ProcParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitProcParameter(s)
	}
}

func (p *abbParser) ProcParameter() (localctx IProcParameterContext) {
	this := p
	_ = this

	localctx = NewProcParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, abbParserRULE_procParameter)
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
		p.SetState(82)
		p.Match(abbParserBRACKET_OPEN)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == abbParserIDENTIFIER {
		{
			p.SetState(83)
			p.Match(abbParserIDENTIFIER)
		}

	}
	{
		p.SetState(86)
		p.Match(abbParserBRACKET_CLOSE)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) AllFunctionParameter() []IFunctionParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem())
	var tst = make([]IFunctionParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionParameterContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) FunctionParameter(i int) IFunctionParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionCallContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(abbParserSEMICOLON, 0)
}

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(abbParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(abbParserCOMMA, i)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *abbParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, abbParserRULE_functionCall)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(abbParserIDENTIFIER)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(89)
				p.FunctionParameter()
			}
			{
				p.SetState(90)
				p.Match(abbParserCOMMA)
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	{
		p.SetState(97)
		p.FunctionParameter()
	}
	{
		p.SetState(98)
		p.Match(abbParserSEMICOLON)
	}

	return localctx
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_functionParameter
	return p
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) ON_CALL() antlr.TerminalNode {
	return s.GetToken(abbParserON_CALL, 0)
}

func (s *FunctionParameterContext) OFF_CALL() antlr.TerminalNode {
	return s.GetToken(abbParserOFF_CALL, 0)
}

func (s *FunctionParameterContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *FunctionParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (p *abbParser) FunctionParameter() (localctx IFunctionParameterContext) {
	this := p
	_ = this

	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, abbParserRULE_functionParameter)

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

	switch p.GetTokenStream().LA(1) {
	case abbParserON_CALL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(abbParserON_CALL)
		}

	case abbParserOFF_CALL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(abbParserOFF_CALL)
		}

	case abbParserPLUS, abbParserMINUS, abbParserBOOLLITERAL, abbParserCHARLITERAL, abbParserSTRINGLITERAL, abbParserFLOATLITERAL, abbParserINTLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Primitive()
		}

	case abbParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(abbParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Init_() IInit_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInit_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInit_Context)
}

func (s *DeclarationContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(abbParserIDENTIFIER, 0)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(abbParserSEMICOLON, 0)
}

func (s *DeclarationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(abbParserEQUALS, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *abbParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, abbParserRULE_declaration)
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
		p.SetState(106)
		p.Init_()
	}
	{
		p.SetState(107)
		p.Type_()
	}
	{
		p.SetState(108)
		p.Match(abbParserIDENTIFIER)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == abbParserEQUALS {
		{
			p.SetState(109)
			p.Match(abbParserEQUALS)
		}
		{
			p.SetState(110)
			p.Expression()
		}

	}
	{
		p.SetState(113)
		p.Match(abbParserSEMICOLON)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TOOLDATA() antlr.TerminalNode {
	return s.GetToken(abbParserTOOLDATA, 0)
}

func (s *Type_Context) WOBJDATA() antlr.TerminalNode {
	return s.GetToken(abbParserWOBJDATA, 0)
}

func (s *Type_Context) SPEEDDATA() antlr.TerminalNode {
	return s.GetToken(abbParserSPEEDDATA, 0)
}

func (s *Type_Context) ZONEDATA() antlr.TerminalNode {
	return s.GetToken(abbParserZONEDATA, 0)
}

func (s *Type_Context) CLOCK() antlr.TerminalNode {
	return s.GetToken(abbParserCLOCK, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(abbParserBOOL, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *abbParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, abbParserRULE_type_)
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
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<abbParserTOOLDATA)|(1<<abbParserWOBJDATA)|(1<<abbParserSPEEDDATA)|(1<<abbParserZONEDATA)|(1<<abbParserCLOCK)|(1<<abbParserBOOL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInit_Context is an interface to support dynamic dispatch.
type IInit_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInit_Context differentiates from other interfaces.
	IsInit_Context()
}

type Init_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInit_Context() *Init_Context {
	var p = new(Init_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_init_
	return p
}

func (*Init_Context) IsInit_Context() {}

func NewInit_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Init_Context {
	var p = new(Init_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_init_

	return p
}

func (s *Init_Context) GetParser() antlr.Parser { return s.parser }

func (s *Init_Context) CONST() antlr.TerminalNode {
	return s.GetToken(abbParserCONST, 0)
}

func (s *Init_Context) PERS() antlr.TerminalNode {
	return s.GetToken(abbParserPERS, 0)
}

func (s *Init_Context) VAR() antlr.TerminalNode {
	return s.GetToken(abbParserVAR, 0)
}

func (s *Init_Context) LOCAL() antlr.TerminalNode {
	return s.GetToken(abbParserLOCAL, 0)
}

func (s *Init_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Init_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Init_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterInit_(s)
	}
}

func (s *Init_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitInit_(s)
	}
}

func (p *abbParser) Init_() (localctx IInit_Context) {
	this := p
	_ = this

	localctx = NewInit_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, abbParserRULE_init_)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == abbParserLOCAL {
		{
			p.SetState(117)
			p.Match(abbParserLOCAL)
		}

	}
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<abbParserCONST)|(1<<abbParserPERS)|(1<<abbParserVAR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = abbParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Array_() IArray_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_Context)
}

func (s *ExpressionContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *abbParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, abbParserRULE_expression)

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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case abbParserSQUARE_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Array_()
		}

	case abbParserPLUS, abbParserMINUS, abbParserBOOLLITERAL, abbParserCHARLITERAL, abbParserSTRINGLITERAL, abbParserFLOATLITERAL, abbParserINTLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Primitive()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArray_Context is an interface to support dynamic dispatch.
type IArray_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_Context differentiates from other interfaces.
	IsArray_Context()
}

type Array_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_Context() *Array_Context {
	var p = new(Array_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_array_
	return p
}

func (*Array_Context) IsArray_Context() {}

func NewArray_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_Context {
	var p = new(Array_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_array_

	return p
}

func (s *Array_Context) GetParser() antlr.Parser { return s.parser }

func (s *Array_Context) SQUARE_OPEN() antlr.TerminalNode {
	return s.GetToken(abbParserSQUARE_OPEN, 0)
}

func (s *Array_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Array_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Array_Context) SQUARE_CLOSE() antlr.TerminalNode {
	return s.GetToken(abbParserSQUARE_CLOSE, 0)
}

func (s *Array_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(abbParserCOMMA)
}

func (s *Array_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(abbParserCOMMA, i)
}

func (s *Array_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterArray_(s)
	}
}

func (s *Array_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitArray_(s)
	}
}

func (p *abbParser) Array_() (localctx IArray_Context) {
	this := p
	_ = this

	localctx = NewArray_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, abbParserRULE_array_)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(abbParserSQUARE_OPEN)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(127)
				p.Expression()
			}
			{
				p.SetState(128)
				p.Match(abbParserCOMMA)
			}

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	{
		p.SetState(135)
		p.Expression()
	}
	{
		p.SetState(136)
		p.Match(abbParserSQUARE_CLOSE)
	}

	return localctx
}

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	var p = new(PrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = abbParserRULE_primitive
	return p
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	var p = new(PrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = abbParserRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveContext) BOOLLITERAL() antlr.TerminalNode {
	return s.GetToken(abbParserBOOLLITERAL, 0)
}

func (s *PrimitiveContext) CHARLITERAL() antlr.TerminalNode {
	return s.GetToken(abbParserCHARLITERAL, 0)
}

func (s *PrimitiveContext) STRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(abbParserSTRINGLITERAL, 0)
}

func (s *PrimitiveContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(abbParserFLOATLITERAL, 0)
}

func (s *PrimitiveContext) PLUS() antlr.TerminalNode {
	return s.GetToken(abbParserPLUS, 0)
}

func (s *PrimitiveContext) MINUS() antlr.TerminalNode {
	return s.GetToken(abbParserMINUS, 0)
}

func (s *PrimitiveContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(abbParserINTLITERAL, 0)
}

func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(abbParserListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (p *abbParser) Primitive() (localctx IPrimitiveContext) {
	this := p
	_ = this

	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, abbParserRULE_primitive)
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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Match(abbParserBOOLLITERAL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(abbParserCHARLITERAL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Match(abbParserSTRINGLITERAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == abbParserPLUS || _la == abbParserMINUS {
			{
				p.SetState(141)
				_la = p.GetTokenStream().LA(1)

				if !(_la == abbParserPLUS || _la == abbParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(144)
			p.Match(abbParserFLOATLITERAL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == abbParserPLUS || _la == abbParserMINUS {
			{
				p.SetState(145)
				_la = p.GetTokenStream().LA(1)

				if !(_la == abbParserPLUS || _la == abbParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(148)
			p.Match(abbParserINTLITERAL)
		}

	}

	return localctx
}
