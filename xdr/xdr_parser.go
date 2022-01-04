// Code generated from xdr.g4 by ANTLR 4.9.3. DO NOT EDIT.

package xdr // xdr
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 213,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 46, 10, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	5, 2, 60, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 67, 10, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 75, 10, 2, 3, 3, 3, 3, 5, 3, 79, 10, 3,
	3, 4, 3, 4, 3, 5, 5, 5, 84, 10, 5, 3, 5, 3, 5, 5, 5, 88, 10, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 99, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 113,
	10, 7, 12, 7, 14, 7, 116, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 130, 10, 9, 12, 9, 14, 9, 133, 11,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 7, 11, 147, 10, 11, 12, 11, 14, 11, 150, 11, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 157, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 169, 10, 12, 12, 12, 14,
	12, 172, 11, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	202, 10, 14, 3, 15, 3, 15, 5, 15, 206, 10, 15, 3, 16, 6, 16, 209, 10, 16,
	13, 16, 14, 16, 210, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 2, 3, 3, 2, 35, 37, 2, 229, 2, 74, 3, 2, 2, 2, 4, 78,
	3, 2, 2, 2, 6, 80, 3, 2, 2, 2, 8, 98, 3, 2, 2, 2, 10, 100, 3, 2, 2, 2,
	12, 103, 3, 2, 2, 2, 14, 119, 3, 2, 2, 2, 16, 122, 3, 2, 2, 2, 18, 136,
	3, 2, 2, 2, 20, 139, 3, 2, 2, 2, 22, 160, 3, 2, 2, 2, 24, 176, 3, 2, 2,
	2, 26, 201, 3, 2, 2, 2, 28, 205, 3, 2, 2, 2, 30, 208, 3, 2, 2, 2, 32, 33,
	5, 8, 5, 2, 33, 34, 7, 38, 2, 2, 34, 75, 3, 2, 2, 2, 35, 36, 5, 8, 5, 2,
	36, 37, 7, 38, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40, 7,
	4, 2, 2, 40, 75, 3, 2, 2, 2, 41, 42, 5, 8, 5, 2, 42, 43, 7, 38, 2, 2, 43,
	45, 7, 5, 2, 2, 44, 46, 5, 4, 3, 2, 45, 44, 3, 2, 2, 2, 45, 46, 3, 2, 2,
	2, 46, 47, 3, 2, 2, 2, 47, 48, 7, 6, 2, 2, 48, 75, 3, 2, 2, 2, 49, 50,
	7, 7, 2, 2, 50, 51, 7, 38, 2, 2, 51, 52, 7, 3, 2, 2, 52, 53, 5, 4, 3, 2,
	53, 54, 7, 4, 2, 2, 54, 75, 3, 2, 2, 2, 55, 56, 7, 7, 2, 2, 56, 57, 7,
	38, 2, 2, 57, 59, 7, 5, 2, 2, 58, 60, 5, 4, 3, 2, 59, 58, 3, 2, 2, 2, 59,
	60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 75, 7, 6, 2, 2, 62, 63, 7, 8, 2,
	2, 63, 64, 7, 38, 2, 2, 64, 66, 7, 5, 2, 2, 65, 67, 5, 4, 3, 2, 66, 65,
	3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 75, 7, 6, 2, 2,
	69, 70, 5, 8, 5, 2, 70, 71, 7, 9, 2, 2, 71, 72, 7, 38, 2, 2, 72, 75, 3,
	2, 2, 2, 73, 75, 7, 10, 2, 2, 74, 32, 3, 2, 2, 2, 74, 35, 3, 2, 2, 2, 74,
	41, 3, 2, 2, 2, 74, 49, 3, 2, 2, 2, 74, 55, 3, 2, 2, 2, 74, 62, 3, 2, 2,
	2, 74, 69, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 3, 3, 2, 2, 2, 76, 79, 5,
	6, 4, 2, 77, 79, 7, 38, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79,
	5, 3, 2, 2, 2, 80, 81, 9, 2, 2, 2, 81, 7, 3, 2, 2, 2, 82, 84, 7, 11, 2,
	2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 99,
	7, 12, 2, 2, 86, 88, 7, 11, 2, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3, 2, 2,
	2, 88, 89, 3, 2, 2, 2, 89, 99, 7, 13, 2, 2, 90, 99, 7, 14, 2, 2, 91, 99,
	7, 15, 2, 2, 92, 99, 7, 16, 2, 2, 93, 99, 7, 17, 2, 2, 94, 99, 5, 10, 6,
	2, 95, 99, 5, 14, 8, 2, 96, 99, 5, 18, 10, 2, 97, 99, 7, 38, 2, 2, 98,
	83, 3, 2, 2, 2, 98, 87, 3, 2, 2, 2, 98, 90, 3, 2, 2, 2, 98, 91, 3, 2, 2,
	2, 98, 92, 3, 2, 2, 2, 98, 93, 3, 2, 2, 2, 98, 94, 3, 2, 2, 2, 98, 95,
	3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 9, 3, 2, 2, 2,
	100, 101, 7, 18, 2, 2, 101, 102, 5, 12, 7, 2, 102, 11, 3, 2, 2, 2, 103,
	104, 7, 19, 2, 2, 104, 105, 7, 38, 2, 2, 105, 106, 7, 20, 2, 2, 106, 107,
	5, 4, 3, 2, 107, 114, 3, 2, 2, 2, 108, 109, 7, 21, 2, 2, 109, 110, 7, 38,
	2, 2, 110, 111, 7, 20, 2, 2, 111, 113, 5, 4, 3, 2, 112, 108, 3, 2, 2, 2,
	113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115,
	117, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 118, 7, 22, 2, 2, 118, 13,
	3, 2, 2, 2, 119, 120, 7, 23, 2, 2, 120, 121, 5, 16, 9, 2, 121, 15, 3, 2,
	2, 2, 122, 123, 7, 19, 2, 2, 123, 124, 5, 2, 2, 2, 124, 125, 7, 24, 2,
	2, 125, 131, 3, 2, 2, 2, 126, 127, 5, 2, 2, 2, 127, 128, 7, 24, 2, 2, 128,
	130, 3, 2, 2, 2, 129, 126, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2, 133, 131, 3, 2,
	2, 2, 134, 135, 7, 22, 2, 2, 135, 17, 3, 2, 2, 2, 136, 137, 7, 25, 2, 2,
	137, 138, 5, 20, 11, 2, 138, 19, 3, 2, 2, 2, 139, 140, 7, 26, 2, 2, 140,
	141, 7, 27, 2, 2, 141, 142, 5, 2, 2, 2, 142, 143, 7, 28, 2, 2, 143, 144,
	7, 19, 2, 2, 144, 148, 5, 22, 12, 2, 145, 147, 5, 22, 12, 2, 146, 145,
	3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2,
	2, 2, 149, 156, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152, 7, 29, 2, 2,
	152, 153, 7, 30, 2, 2, 153, 154, 5, 2, 2, 2, 154, 155, 7, 24, 2, 2, 155,
	157, 3, 2, 2, 2, 156, 151, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158,
	3, 2, 2, 2, 158, 159, 7, 22, 2, 2, 159, 21, 3, 2, 2, 2, 160, 161, 7, 31,
	2, 2, 161, 162, 5, 4, 3, 2, 162, 163, 7, 30, 2, 2, 163, 170, 3, 2, 2, 2,
	164, 165, 7, 31, 2, 2, 165, 166, 5, 4, 3, 2, 166, 167, 7, 30, 2, 2, 167,
	169, 3, 2, 2, 2, 168, 164, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168,
	3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170, 3, 2,
	2, 2, 173, 174, 5, 2, 2, 2, 174, 175, 7, 24, 2, 2, 175, 23, 3, 2, 2, 2,
	176, 177, 7, 32, 2, 2, 177, 178, 7, 38, 2, 2, 178, 179, 7, 20, 2, 2, 179,
	180, 5, 6, 4, 2, 180, 181, 7, 24, 2, 2, 181, 25, 3, 2, 2, 2, 182, 183,
	7, 33, 2, 2, 183, 184, 5, 2, 2, 2, 184, 185, 7, 24, 2, 2, 185, 202, 3,
	2, 2, 2, 186, 187, 7, 18, 2, 2, 187, 188, 7, 38, 2, 2, 188, 189, 5, 12,
	7, 2, 189, 190, 7, 24, 2, 2, 190, 202, 3, 2, 2, 2, 191, 192, 7, 23, 2,
	2, 192, 193, 7, 38, 2, 2, 193, 194, 5, 16, 9, 2, 194, 195, 7, 24, 2, 2,
	195, 202, 3, 2, 2, 2, 196, 197, 7, 25, 2, 2, 197, 198, 7, 38, 2, 2, 198,
	199, 5, 20, 11, 2, 199, 200, 7, 24, 2, 2, 200, 202, 3, 2, 2, 2, 201, 182,
	3, 2, 2, 2, 201, 186, 3, 2, 2, 2, 201, 191, 3, 2, 2, 2, 201, 196, 3, 2,
	2, 2, 202, 27, 3, 2, 2, 2, 203, 206, 5, 26, 14, 2, 204, 206, 5, 24, 13,
	2, 205, 203, 3, 2, 2, 2, 205, 204, 3, 2, 2, 2, 206, 29, 3, 2, 2, 2, 207,
	209, 5, 28, 15, 2, 208, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 208,
	3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 31, 3, 2, 2, 2, 18, 45, 59, 66,
	74, 78, 83, 87, 98, 114, 131, 148, 156, 170, 201, 205, 210,
}
var literalNames = []string{
	"", "'['", "']'", "'<'", "'>'", "'opaque'", "'string'", "'*'", "'void'",
	"'unsigned'", "'int'", "'hyper'", "'float'", "'double'", "'quadruple'",
	"'bool'", "'enum'", "'{'", "'='", "','", "'}'", "'struct'", "';'", "'union'",
	"'switch'", "'('", "')'", "'default'", "':'", "'case'", "'const'", "'typedef'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMENT", "OCTAL",
	"DECIMAL", "HEXADECIMAL", "IDENTIFIER", "WS",
}

var ruleNames = []string{
	"declaration", "value", "constant", "typeSpecifier", "enumTypeSpec", "enumBody",
	"structTypeSpec", "structBody", "unionTypeSpec", "unionBody", "caseSpec",
	"constantDef", "typeDef", "definition", "xdrSpecification",
}

type xdrParser struct {
	*antlr.BaseParser
}

// NewxdrParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *xdrParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewxdrParser(input antlr.TokenStream) *xdrParser {
	this := new(xdrParser)
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
	this.GrammarFileName = "xdr.g4"

	return this
}

// xdrParser tokens.
const (
	xdrParserEOF         = antlr.TokenEOF
	xdrParserT__0        = 1
	xdrParserT__1        = 2
	xdrParserT__2        = 3
	xdrParserT__3        = 4
	xdrParserT__4        = 5
	xdrParserT__5        = 6
	xdrParserT__6        = 7
	xdrParserT__7        = 8
	xdrParserT__8        = 9
	xdrParserT__9        = 10
	xdrParserT__10       = 11
	xdrParserT__11       = 12
	xdrParserT__12       = 13
	xdrParserT__13       = 14
	xdrParserT__14       = 15
	xdrParserT__15       = 16
	xdrParserT__16       = 17
	xdrParserT__17       = 18
	xdrParserT__18       = 19
	xdrParserT__19       = 20
	xdrParserT__20       = 21
	xdrParserT__21       = 22
	xdrParserT__22       = 23
	xdrParserT__23       = 24
	xdrParserT__24       = 25
	xdrParserT__25       = 26
	xdrParserT__26       = 27
	xdrParserT__27       = 28
	xdrParserT__28       = 29
	xdrParserT__29       = 30
	xdrParserT__30       = 31
	xdrParserCOMMENT     = 32
	xdrParserOCTAL       = 33
	xdrParserDECIMAL     = 34
	xdrParserHEXADECIMAL = 35
	xdrParserIDENTIFIER  = 36
	xdrParserWS          = 37
)

// xdrParser rules.
const (
	xdrParserRULE_declaration      = 0
	xdrParserRULE_value            = 1
	xdrParserRULE_constant         = 2
	xdrParserRULE_typeSpecifier    = 3
	xdrParserRULE_enumTypeSpec     = 4
	xdrParserRULE_enumBody         = 5
	xdrParserRULE_structTypeSpec   = 6
	xdrParserRULE_structBody       = 7
	xdrParserRULE_unionTypeSpec    = 8
	xdrParserRULE_unionBody        = 9
	xdrParserRULE_caseSpec         = 10
	xdrParserRULE_constantDef      = 11
	xdrParserRULE_typeDef          = 12
	xdrParserRULE_definition       = 13
	xdrParserRULE_xdrSpecification = 14
)

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
	p.RuleIndex = xdrParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, 0)
}

func (s *DeclarationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *xdrParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, xdrParserRULE_declaration)
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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.TypeSpecifier()
		}
		{
			p.SetState(31)
			p.Match(xdrParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.TypeSpecifier()
		}
		{
			p.SetState(34)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(35)
			p.Match(xdrParserT__0)
		}
		{
			p.SetState(36)
			p.Value()
		}
		{
			p.SetState(37)
			p.Match(xdrParserT__1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.TypeSpecifier()
		}
		{
			p.SetState(40)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(41)
			p.Match(xdrParserT__2)
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(xdrParserOCTAL-33))|(1<<(xdrParserDECIMAL-33))|(1<<(xdrParserHEXADECIMAL-33))|(1<<(xdrParserIDENTIFIER-33)))) != 0 {
			{
				p.SetState(42)
				p.Value()
			}

		}
		{
			p.SetState(45)
			p.Match(xdrParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Match(xdrParserT__4)
		}
		{
			p.SetState(48)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(49)
			p.Match(xdrParserT__0)
		}
		{
			p.SetState(50)
			p.Value()
		}
		{
			p.SetState(51)
			p.Match(xdrParserT__1)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)
			p.Match(xdrParserT__4)
		}
		{
			p.SetState(54)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(55)
			p.Match(xdrParserT__2)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(xdrParserOCTAL-33))|(1<<(xdrParserDECIMAL-33))|(1<<(xdrParserHEXADECIMAL-33))|(1<<(xdrParserIDENTIFIER-33)))) != 0 {
			{
				p.SetState(56)
				p.Value()
			}

		}
		{
			p.SetState(59)
			p.Match(xdrParserT__3)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Match(xdrParserT__5)
		}
		{
			p.SetState(61)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(62)
			p.Match(xdrParserT__2)
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(xdrParserOCTAL-33))|(1<<(xdrParserDECIMAL-33))|(1<<(xdrParserHEXADECIMAL-33))|(1<<(xdrParserIDENTIFIER-33)))) != 0 {
			{
				p.SetState(63)
				p.Value()
			}

		}
		{
			p.SetState(66)
			p.Match(xdrParserT__3)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(67)
			p.TypeSpecifier()
		}
		{
			p.SetState(68)
			p.Match(xdrParserT__6)
		}
		{
			p.SetState(69)
			p.Match(xdrParserIDENTIFIER)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.Match(xdrParserT__7)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *xdrParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, xdrParserRULE_value)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xdrParserOCTAL, xdrParserDECIMAL, xdrParserHEXADECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Constant()
		}

	case xdrParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(xdrParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(xdrParserDECIMAL, 0)
}

func (s *ConstantContext) HEXADECIMAL() antlr.TerminalNode {
	return s.GetToken(xdrParserHEXADECIMAL, 0)
}

func (s *ConstantContext) OCTAL() antlr.TerminalNode {
	return s.GetToken(xdrParserOCTAL, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *xdrParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, xdrParserRULE_constant)
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
		p.SetState(78)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(xdrParserOCTAL-33))|(1<<(xdrParserDECIMAL-33))|(1<<(xdrParserHEXADECIMAL-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) EnumTypeSpec() IEnumTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeSpecContext)
}

func (s *TypeSpecifierContext) StructTypeSpec() IStructTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTypeSpecContext)
}

func (s *TypeSpecifierContext) UnionTypeSpec() IUnionTypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeSpecContext)
}

func (s *TypeSpecifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *xdrParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	this := p
	_ = this

	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, xdrParserRULE_typeSpecifier)
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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xdrParserT__8 {
			{
				p.SetState(80)
				p.Match(xdrParserT__8)
			}

		}
		{
			p.SetState(83)
			p.Match(xdrParserT__9)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == xdrParserT__8 {
			{
				p.SetState(84)
				p.Match(xdrParserT__8)
			}

		}
		{
			p.SetState(87)
			p.Match(xdrParserT__10)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(xdrParserT__11)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(xdrParserT__12)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.Match(xdrParserT__13)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.Match(xdrParserT__14)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.EnumTypeSpec()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(93)
			p.StructTypeSpec()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(94)
			p.UnionTypeSpec()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(95)
			p.Match(xdrParserIDENTIFIER)
		}

	}

	return localctx
}

// IEnumTypeSpecContext is an interface to support dynamic dispatch.
type IEnumTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeSpecContext differentiates from other interfaces.
	IsEnumTypeSpecContext()
}

type EnumTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeSpecContext() *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_enumTypeSpec
	return p
}

func (*EnumTypeSpecContext) IsEnumTypeSpecContext() {}

func NewEnumTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeSpecContext {
	var p = new(EnumTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_enumTypeSpec

	return p
}

func (s *EnumTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeSpecContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterEnumTypeSpec(s)
	}
}

func (s *EnumTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitEnumTypeSpec(s)
	}
}

func (p *xdrParser) EnumTypeSpec() (localctx IEnumTypeSpecContext) {
	this := p
	_ = this

	localctx = NewEnumTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, xdrParserRULE_enumTypeSpec)

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
		p.SetState(98)
		p.Match(xdrParserT__15)
	}
	{
		p.SetState(99)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(xdrParserIDENTIFIER)
}

func (s *EnumBodyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, i)
}

func (s *EnumBodyContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *EnumBodyContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *xdrParser) EnumBody() (localctx IEnumBodyContext) {
	this := p
	_ = this

	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, xdrParserRULE_enumBody)
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
		p.Match(xdrParserT__16)
	}

	{
		p.SetState(102)
		p.Match(xdrParserIDENTIFIER)
	}
	{
		p.SetState(103)
		p.Match(xdrParserT__17)
	}
	{
		p.SetState(104)
		p.Value()
	}

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xdrParserT__18 {
		{
			p.SetState(106)
			p.Match(xdrParserT__18)
		}
		{
			p.SetState(107)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(108)
			p.Match(xdrParserT__17)
		}
		{
			p.SetState(109)
			p.Value()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(xdrParserT__19)
	}

	return localctx
}

// IStructTypeSpecContext is an interface to support dynamic dispatch.
type IStructTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTypeSpecContext differentiates from other interfaces.
	IsStructTypeSpecContext()
}

type StructTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeSpecContext() *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_structTypeSpec
	return p
}

func (*StructTypeSpecContext) IsStructTypeSpecContext() {}

func NewStructTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_structTypeSpec

	return p
}

func (s *StructTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeSpecContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitStructTypeSpec(s)
	}
}

func (p *xdrParser) StructTypeSpec() (localctx IStructTypeSpecContext) {
	this := p
	_ = this

	localctx = NewStructTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, xdrParserRULE_structTypeSpec)

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
		p.Match(xdrParserT__20)
	}
	{
		p.SetState(118)
		p.StructBody()
	}

	return localctx
}

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_structBody
	return p
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *StructBodyContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (p *xdrParser) StructBody() (localctx IStructBodyContext) {
	this := p
	_ = this

	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, xdrParserRULE_structBody)
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
		p.SetState(120)
		p.Match(xdrParserT__16)
	}

	{
		p.SetState(121)
		p.Declaration()
	}
	{
		p.SetState(122)
		p.Match(xdrParserT__21)
	}

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(xdrParserT__4-5))|(1<<(xdrParserT__5-5))|(1<<(xdrParserT__7-5))|(1<<(xdrParserT__8-5))|(1<<(xdrParserT__9-5))|(1<<(xdrParserT__10-5))|(1<<(xdrParserT__11-5))|(1<<(xdrParserT__12-5))|(1<<(xdrParserT__13-5))|(1<<(xdrParserT__14-5))|(1<<(xdrParserT__15-5))|(1<<(xdrParserT__20-5))|(1<<(xdrParserT__22-5))|(1<<(xdrParserIDENTIFIER-5)))) != 0 {
		{
			p.SetState(124)
			p.Declaration()
		}
		{
			p.SetState(125)
			p.Match(xdrParserT__21)
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(xdrParserT__19)
	}

	return localctx
}

// IUnionTypeSpecContext is an interface to support dynamic dispatch.
type IUnionTypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeSpecContext differentiates from other interfaces.
	IsUnionTypeSpecContext()
}

type UnionTypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeSpecContext() *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_unionTypeSpec
	return p
}

func (*UnionTypeSpecContext) IsUnionTypeSpecContext() {}

func NewUnionTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeSpecContext {
	var p = new(UnionTypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_unionTypeSpec

	return p
}

func (s *UnionTypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeSpecContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *UnionTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterUnionTypeSpec(s)
	}
}

func (s *UnionTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitUnionTypeSpec(s)
	}
}

func (p *xdrParser) UnionTypeSpec() (localctx IUnionTypeSpecContext) {
	this := p
	_ = this

	localctx = NewUnionTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, xdrParserRULE_unionTypeSpec)

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
		p.SetState(134)
		p.Match(xdrParserT__22)
	}
	{
		p.SetState(135)
		p.UnionBody()
	}

	return localctx
}

// IUnionBodyContext is an interface to support dynamic dispatch.
type IUnionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionBodyContext differentiates from other interfaces.
	IsUnionBodyContext()
}

type UnionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionBodyContext() *UnionBodyContext {
	var p = new(UnionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_unionBody
	return p
}

func (*UnionBodyContext) IsUnionBodyContext() {}

func NewUnionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionBodyContext {
	var p = new(UnionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_unionBody

	return p
}

func (s *UnionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionBodyContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *UnionBodyContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *UnionBodyContext) AllCaseSpec() []ICaseSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem())
	var tst = make([]ICaseSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseSpecContext)
		}
	}

	return tst
}

func (s *UnionBodyContext) CaseSpec(i int) ICaseSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseSpecContext)
}

func (s *UnionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterUnionBody(s)
	}
}

func (s *UnionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitUnionBody(s)
	}
}

func (p *xdrParser) UnionBody() (localctx IUnionBodyContext) {
	this := p
	_ = this

	localctx = NewUnionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, xdrParserRULE_unionBody)
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
		p.SetState(137)
		p.Match(xdrParserT__23)
	}
	{
		p.SetState(138)
		p.Match(xdrParserT__24)
	}
	{
		p.SetState(139)
		p.Declaration()
	}
	{
		p.SetState(140)
		p.Match(xdrParserT__25)
	}
	{
		p.SetState(141)
		p.Match(xdrParserT__16)
	}
	{
		p.SetState(142)
		p.CaseSpec()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xdrParserT__28 {
		{
			p.SetState(143)
			p.CaseSpec()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == xdrParserT__26 {
		{
			p.SetState(149)
			p.Match(xdrParserT__26)
		}
		{
			p.SetState(150)
			p.Match(xdrParserT__27)
		}
		{
			p.SetState(151)
			p.Declaration()
		}
		{
			p.SetState(152)
			p.Match(xdrParserT__21)
		}

	}
	{
		p.SetState(156)
		p.Match(xdrParserT__19)
	}

	return localctx
}

// ICaseSpecContext is an interface to support dynamic dispatch.
type ICaseSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseSpecContext differentiates from other interfaces.
	IsCaseSpecContext()
}

type CaseSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseSpecContext() *CaseSpecContext {
	var p = new(CaseSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_caseSpec
	return p
}

func (*CaseSpecContext) IsCaseSpecContext() {}

func NewCaseSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseSpecContext {
	var p = new(CaseSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_caseSpec

	return p
}

func (s *CaseSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseSpecContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *CaseSpecContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *CaseSpecContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CaseSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterCaseSpec(s)
	}
}

func (s *CaseSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitCaseSpec(s)
	}
}

func (p *xdrParser) CaseSpec() (localctx ICaseSpecContext) {
	this := p
	_ = this

	localctx = NewCaseSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, xdrParserRULE_caseSpec)
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
		p.SetState(158)
		p.Match(xdrParserT__28)
	}
	{
		p.SetState(159)
		p.Value()
	}
	{
		p.SetState(160)
		p.Match(xdrParserT__27)
	}

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == xdrParserT__28 {
		{
			p.SetState(162)
			p.Match(xdrParserT__28)
		}
		{
			p.SetState(163)
			p.Value()
		}
		{
			p.SetState(164)
			p.Match(xdrParserT__27)
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
		p.Declaration()
	}
	{
		p.SetState(172)
		p.Match(xdrParserT__21)
	}

	return localctx
}

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	var p = new(ConstantDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_constantDef
	return p
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	var p = new(ConstantDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (p *xdrParser) ConstantDef() (localctx IConstantDefContext) {
	this := p
	_ = this

	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, xdrParserRULE_constantDef)

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
		p.SetState(174)
		p.Match(xdrParserT__29)
	}
	{
		p.SetState(175)
		p.Match(xdrParserIDENTIFIER)
	}
	{
		p.SetState(176)
		p.Match(xdrParserT__17)
	}
	{
		p.SetState(177)
		p.Constant()
	}
	{
		p.SetState(178)
		p.Match(xdrParserT__21)
	}

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TypeDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(xdrParserIDENTIFIER, 0)
}

func (s *TypeDefContext) EnumBody() IEnumBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *TypeDefContext) StructBody() IStructBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *TypeDefContext) UnionBody() IUnionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionBodyContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (p *xdrParser) TypeDef() (localctx ITypeDefContext) {
	this := p
	_ = this

	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, xdrParserRULE_typeDef)

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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xdrParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(xdrParserT__30)
		}
		{
			p.SetState(181)
			p.Declaration()
		}
		{
			p.SetState(182)
			p.Match(xdrParserT__21)
		}

	case xdrParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(xdrParserT__15)
		}
		{
			p.SetState(185)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(186)
			p.EnumBody()
		}
		{
			p.SetState(187)
			p.Match(xdrParserT__21)
		}

	case xdrParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.Match(xdrParserT__20)
		}
		{
			p.SetState(190)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(191)
			p.StructBody()
		}
		{
			p.SetState(192)
			p.Match(xdrParserT__21)
		}

	case xdrParserT__22:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(xdrParserT__22)
		}
		{
			p.SetState(195)
			p.Match(xdrParserIDENTIFIER)
		}
		{
			p.SetState(196)
			p.UnionBody()
		}
		{
			p.SetState(197)
			p.Match(xdrParserT__21)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) TypeDef() ITypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *DefinitionContext) ConstantDef() IConstantDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantDefContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *xdrParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, xdrParserRULE_definition)

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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case xdrParserT__15, xdrParserT__20, xdrParserT__22, xdrParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.TypeDef()
		}

	case xdrParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.ConstantDef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IXdrSpecificationContext is an interface to support dynamic dispatch.
type IXdrSpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXdrSpecificationContext differentiates from other interfaces.
	IsXdrSpecificationContext()
}

type XdrSpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXdrSpecificationContext() *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = xdrParserRULE_xdrSpecification
	return p
}

func (*XdrSpecificationContext) IsXdrSpecificationContext() {}

func NewXdrSpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XdrSpecificationContext {
	var p = new(XdrSpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = xdrParserRULE_xdrSpecification

	return p
}

func (s *XdrSpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *XdrSpecificationContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *XdrSpecificationContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *XdrSpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XdrSpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XdrSpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.EnterXdrSpecification(s)
	}
}

func (s *XdrSpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(xdrListener); ok {
		listenerT.ExitXdrSpecification(s)
	}
}

func (p *xdrParser) XdrSpecification() (localctx IXdrSpecificationContext) {
	this := p
	_ = this

	localctx = NewXdrSpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, xdrParserRULE_xdrSpecification)
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
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<xdrParserT__15)|(1<<xdrParserT__20)|(1<<xdrParserT__22)|(1<<xdrParserT__29)|(1<<xdrParserT__30))) != 0) {
		{
			p.SetState(205)
			p.Definition()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
