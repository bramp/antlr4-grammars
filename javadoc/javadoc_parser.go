// Generated from JavadocParser.g4 by ANTLR 4.7.

package javadoc // JavadocParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 242,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 7, 2, 50, 10, 2, 12, 2, 14, 2, 53, 11, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 2, 3, 2,
	3, 2, 5, 2, 68, 10, 2, 3, 3, 3, 3, 7, 3, 72, 10, 3, 12, 3, 14, 3, 75, 11,
	3, 3, 3, 7, 3, 78, 10, 3, 12, 3, 14, 3, 81, 11, 3, 3, 3, 3, 3, 3, 3, 6,
	3, 86, 10, 3, 13, 3, 14, 3, 87, 3, 3, 7, 3, 91, 10, 3, 12, 3, 14, 3, 94,
	11, 3, 3, 3, 3, 3, 5, 3, 98, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 104,
	10, 5, 13, 5, 14, 5, 105, 3, 5, 3, 5, 7, 5, 110, 10, 5, 12, 5, 14, 5, 113,
	11, 5, 3, 6, 3, 6, 7, 6, 117, 10, 6, 12, 6, 14, 6, 120, 11, 6, 3, 6, 3,
	6, 7, 6, 124, 10, 6, 12, 6, 14, 6, 127, 11, 6, 5, 6, 129, 10, 6, 3, 7,
	5, 7, 132, 10, 7, 3, 7, 6, 7, 135, 10, 7, 13, 7, 14, 7, 136, 3, 7, 3, 7,
	3, 7, 7, 7, 142, 10, 7, 12, 7, 14, 7, 145, 11, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 5, 9, 151, 10, 9, 3, 10, 3, 10, 3, 10, 6, 10, 156, 10, 10, 13, 10, 14,
	10, 157, 3, 11, 3, 11, 3, 12, 6, 12, 163, 10, 12, 13, 12, 14, 12, 164,
	3, 13, 5, 13, 168, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 173, 10, 13, 3,
	13, 7, 13, 176, 10, 13, 12, 13, 14, 13, 179, 11, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 15, 5, 15, 186, 10, 15, 3, 16, 6, 16, 189, 10, 16, 13, 16, 14,
	16, 190, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 198, 10, 18, 12, 18,
	14, 18, 201, 11, 18, 3, 18, 5, 18, 204, 10, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 20, 6, 20, 211, 10, 20, 13, 20, 14, 20, 212, 3, 21, 3, 21, 7, 21,
	217, 10, 21, 12, 21, 14, 21, 220, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 7, 22, 227, 10, 22, 12, 22, 14, 22, 230, 11, 22, 3, 22, 7, 22, 233,
	10, 22, 12, 22, 14, 22, 236, 11, 22, 5, 22, 238, 10, 22, 3, 23, 3, 23,
	3, 23, 2, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 2, 6, 3, 2, 4, 5, 6, 2, 3, 3, 6, 6, 8, 9, 13,
	14, 6, 2, 3, 3, 5, 6, 8, 9, 13, 14, 4, 2, 3, 6, 8, 9, 2, 257, 2, 67, 3,
	2, 2, 2, 4, 97, 3, 2, 2, 2, 6, 99, 3, 2, 2, 2, 8, 101, 3, 2, 2, 2, 10,
	128, 3, 2, 2, 2, 12, 131, 3, 2, 2, 2, 14, 146, 3, 2, 2, 2, 16, 150, 3,
	2, 2, 2, 18, 155, 3, 2, 2, 2, 20, 159, 3, 2, 2, 2, 22, 162, 3, 2, 2, 2,
	24, 167, 3, 2, 2, 2, 26, 180, 3, 2, 2, 2, 28, 185, 3, 2, 2, 2, 30, 188,
	3, 2, 2, 2, 32, 192, 3, 2, 2, 2, 34, 194, 3, 2, 2, 2, 36, 207, 3, 2, 2,
	2, 38, 210, 3, 2, 2, 2, 40, 214, 3, 2, 2, 2, 42, 237, 3, 2, 2, 2, 44, 239,
	3, 2, 2, 2, 46, 68, 7, 2, 2, 3, 47, 51, 7, 10, 2, 2, 48, 50, 5, 6, 4, 2,
	49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3,
	2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 5, 4, 3, 2, 55,
	56, 7, 11, 2, 2, 56, 57, 7, 2, 2, 3, 57, 68, 3, 2, 2, 2, 58, 60, 5, 6,
	4, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62,
	3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 5, 4, 3, 2,
	65, 66, 7, 2, 2, 3, 66, 68, 3, 2, 2, 2, 67, 46, 3, 2, 2, 2, 67, 47, 3,
	2, 2, 2, 67, 61, 3, 2, 2, 2, 68, 3, 3, 2, 2, 2, 69, 73, 5, 8, 5, 2, 70,
	72, 5, 6, 4, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2,
	2, 73, 74, 3, 2, 2, 2, 74, 98, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78,
	5, 6, 4, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2,
	79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 98, 5,
	22, 12, 2, 83, 85, 5, 8, 5, 2, 84, 86, 7, 4, 2, 2, 85, 84, 3, 2, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 92, 3,
	2, 2, 2, 89, 91, 5, 6, 4, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92,
	90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2,
	2, 95, 96, 5, 22, 12, 2, 96, 98, 3, 2, 2, 2, 97, 69, 3, 2, 2, 2, 97, 79,
	3, 2, 2, 2, 97, 83, 3, 2, 2, 2, 98, 5, 3, 2, 2, 2, 99, 100, 9, 2, 2, 2,
	100, 7, 3, 2, 2, 2, 101, 111, 5, 10, 6, 2, 102, 104, 5, 20, 11, 2, 103,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 5, 10, 6, 2, 108, 110, 3, 2,
	2, 2, 109, 103, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 9, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 118,
	5, 12, 7, 2, 115, 117, 5, 16, 9, 2, 116, 115, 3, 2, 2, 2, 117, 120, 3,
	2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 129, 3, 2, 2,
	2, 120, 118, 3, 2, 2, 2, 121, 125, 5, 34, 18, 2, 122, 124, 5, 16, 9, 2,
	123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128, 114,
	3, 2, 2, 2, 128, 121, 3, 2, 2, 2, 129, 11, 3, 2, 2, 2, 130, 132, 7, 5,
	2, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2,
	133, 135, 5, 14, 8, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 143, 3, 2, 2, 2, 138, 142,
	5, 14, 8, 2, 139, 142, 7, 5, 2, 2, 140, 142, 7, 7, 2, 2, 141, 138, 3, 2,
	2, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2,
	143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 13, 3, 2, 2, 2, 145, 143,
	3, 2, 2, 2, 146, 147, 9, 3, 2, 2, 147, 15, 3, 2, 2, 2, 148, 151, 5, 34,
	18, 2, 149, 151, 5, 18, 10, 2, 150, 148, 3, 2, 2, 2, 150, 149, 3, 2, 2,
	2, 151, 17, 3, 2, 2, 2, 152, 156, 5, 14, 8, 2, 153, 156, 7, 5, 2, 2, 154,
	156, 7, 7, 2, 2, 155, 152, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 154,
	3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 19, 3, 2, 2, 2, 159, 160, 7, 4, 2, 2, 160, 21, 3, 2, 2, 2, 161,
	163, 5, 24, 13, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162,
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 23, 3, 2, 2, 2, 166, 168, 7, 5,
	2, 2, 167, 166, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2,
	169, 170, 7, 7, 2, 2, 170, 172, 5, 26, 14, 2, 171, 173, 7, 5, 2, 2, 172,
	171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 177, 3, 2, 2, 2, 174, 176,
	5, 28, 15, 2, 175, 174, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3,
	2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 25, 3, 2, 2, 2, 179, 177, 3, 2, 2,
	2, 180, 181, 7, 3, 2, 2, 181, 27, 3, 2, 2, 2, 182, 186, 5, 30, 16, 2, 183,
	186, 5, 34, 18, 2, 184, 186, 7, 4, 2, 2, 185, 182, 3, 2, 2, 2, 185, 183,
	3, 2, 2, 2, 185, 184, 3, 2, 2, 2, 186, 29, 3, 2, 2, 2, 187, 189, 5, 32,
	17, 2, 188, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2,
	190, 191, 3, 2, 2, 2, 191, 31, 3, 2, 2, 2, 192, 193, 9, 4, 2, 2, 193, 33,
	3, 2, 2, 2, 194, 195, 7, 12, 2, 2, 195, 199, 5, 36, 19, 2, 196, 198, 7,
	5, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2,
	2, 199, 200, 3, 2, 2, 2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202,
	204, 5, 38, 20, 2, 203, 202, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 206, 7, 14, 2, 2, 206, 35, 3, 2, 2, 2, 207, 208, 7, 3,
	2, 2, 208, 37, 3, 2, 2, 2, 209, 211, 5, 42, 22, 2, 210, 209, 3, 2, 2, 2,
	211, 212, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213,
	39, 3, 2, 2, 2, 214, 218, 7, 13, 2, 2, 215, 217, 5, 42, 22, 2, 216, 215,
	3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2,
	2, 2, 219, 221, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222, 7, 14, 2, 2,
	222, 41, 3, 2, 2, 2, 223, 238, 5, 40, 21, 2, 224, 234, 5, 44, 23, 2, 225,
	227, 7, 4, 2, 2, 226, 225, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 231, 3, 2, 2, 2, 230, 228, 3, 2,
	2, 2, 231, 233, 5, 44, 23, 2, 232, 228, 3, 2, 2, 2, 233, 236, 3, 2, 2,
	2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236,
	234, 3, 2, 2, 2, 237, 223, 3, 2, 2, 2, 237, 224, 3, 2, 2, 2, 238, 43, 3,
	2, 2, 2, 239, 240, 9, 5, 2, 2, 240, 45, 3, 2, 2, 2, 35, 51, 61, 67, 73,
	79, 87, 92, 97, 105, 111, 118, 125, 128, 131, 136, 141, 143, 150, 155,
	157, 164, 167, 172, 177, 185, 190, 199, 203, 212, 218, 228, 234, 237,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'@'", "'*'", "'/'", "", "", "'{@'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "NAME", "NEWLINE", "SPACE", "TEXT_CONTENT", "AT", "STAR", "SLASH",
	"JAVADOC_START", "JAVADOC_END", "INLINE_TAG_START", "BRACE_OPEN", "BRACE_CLOSE",
}

var ruleNames = []string{
	"documentation", "documentationContent", "skipWhitespace", "description",
	"descriptionLine", "descriptionLineStart", "descriptionLineNoSpaceNoAt",
	"descriptionLineElement", "descriptionLineText", "descriptionNewline",
	"tagSection", "blockTag", "blockTagName", "blockTagContent", "blockTagText",
	"blockTagTextElement", "inlineTag", "inlineTagName", "inlineTagContent",
	"braceExpression", "braceContent", "braceText",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JavadocParser struct {
	*antlr.BaseParser
}

func NewJavadocParser(input antlr.TokenStream) *JavadocParser {
	this := new(JavadocParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JavadocParser.g4"

	return this
}

// JavadocParser tokens.
const (
	JavadocParserEOF              = antlr.TokenEOF
	JavadocParserNAME             = 1
	JavadocParserNEWLINE          = 2
	JavadocParserSPACE            = 3
	JavadocParserTEXT_CONTENT     = 4
	JavadocParserAT               = 5
	JavadocParserSTAR             = 6
	JavadocParserSLASH            = 7
	JavadocParserJAVADOC_START    = 8
	JavadocParserJAVADOC_END      = 9
	JavadocParserINLINE_TAG_START = 10
	JavadocParserBRACE_OPEN       = 11
	JavadocParserBRACE_CLOSE      = 12
)

// JavadocParser rules.
const (
	JavadocParserRULE_documentation              = 0
	JavadocParserRULE_documentationContent       = 1
	JavadocParserRULE_skipWhitespace             = 2
	JavadocParserRULE_description                = 3
	JavadocParserRULE_descriptionLine            = 4
	JavadocParserRULE_descriptionLineStart       = 5
	JavadocParserRULE_descriptionLineNoSpaceNoAt = 6
	JavadocParserRULE_descriptionLineElement     = 7
	JavadocParserRULE_descriptionLineText        = 8
	JavadocParserRULE_descriptionNewline         = 9
	JavadocParserRULE_tagSection                 = 10
	JavadocParserRULE_blockTag                   = 11
	JavadocParserRULE_blockTagName               = 12
	JavadocParserRULE_blockTagContent            = 13
	JavadocParserRULE_blockTagText               = 14
	JavadocParserRULE_blockTagTextElement        = 15
	JavadocParserRULE_inlineTag                  = 16
	JavadocParserRULE_inlineTagName              = 17
	JavadocParserRULE_inlineTagContent           = 18
	JavadocParserRULE_braceExpression            = 19
	JavadocParserRULE_braceContent               = 20
	JavadocParserRULE_braceText                  = 21
)

// IDocumentationContext is an interface to support dynamic dispatch.
type IDocumentationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentationContext differentiates from other interfaces.
	IsDocumentationContext()
}

type DocumentationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentationContext() *DocumentationContext {
	var p = new(DocumentationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_documentation
	return p
}

func (*DocumentationContext) IsDocumentationContext() {}

func NewDocumentationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentationContext {
	var p = new(DocumentationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_documentation

	return p
}

func (s *DocumentationContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentationContext) EOF() antlr.TerminalNode {
	return s.GetToken(JavadocParserEOF, 0)
}

func (s *DocumentationContext) JAVADOC_START() antlr.TerminalNode {
	return s.GetToken(JavadocParserJAVADOC_START, 0)
}

func (s *DocumentationContext) DocumentationContent() IDocumentationContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentationContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentationContentContext)
}

func (s *DocumentationContext) JAVADOC_END() antlr.TerminalNode {
	return s.GetToken(JavadocParserJAVADOC_END, 0)
}

func (s *DocumentationContext) AllSkipWhitespace() []ISkipWhitespaceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkipWhitespaceContext)(nil)).Elem())
	var tst = make([]ISkipWhitespaceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkipWhitespaceContext)
		}
	}

	return tst
}

func (s *DocumentationContext) SkipWhitespace(i int) ISkipWhitespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipWhitespaceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkipWhitespaceContext)
}

func (s *DocumentationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDocumentation(s)
	}
}

func (s *DocumentationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDocumentation(s)
	}
}

func (s *DocumentationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDocumentation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) Documentation() (localctx IDocumentationContext) {
	localctx = NewDocumentationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JavadocParserRULE_documentation)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavadocParserEOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(JavadocParserEOF)
		}

	case JavadocParserJAVADOC_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(JavadocParserJAVADOC_START)
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(46)
					p.SkipWhitespace()
				}

			}
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}
		{
			p.SetState(52)
			p.DocumentationContent()
		}
		{
			p.SetState(53)
			p.Match(JavadocParserJAVADOC_END)
		}
		{
			p.SetState(54)
			p.Match(JavadocParserEOF)
		}

	case JavadocParserNAME, JavadocParserNEWLINE, JavadocParserSPACE, JavadocParserTEXT_CONTENT, JavadocParserAT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserINLINE_TAG_START, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(56)
					p.SkipWhitespace()
				}

			}
			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		{
			p.SetState(62)
			p.DocumentationContent()
		}
		{
			p.SetState(63)
			p.Match(JavadocParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDocumentationContentContext is an interface to support dynamic dispatch.
type IDocumentationContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentationContentContext differentiates from other interfaces.
	IsDocumentationContentContext()
}

type DocumentationContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentationContentContext() *DocumentationContentContext {
	var p = new(DocumentationContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_documentationContent
	return p
}

func (*DocumentationContentContext) IsDocumentationContentContext() {}

func NewDocumentationContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentationContentContext {
	var p = new(DocumentationContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_documentationContent

	return p
}

func (s *DocumentationContentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentationContentContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DocumentationContentContext) AllSkipWhitespace() []ISkipWhitespaceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkipWhitespaceContext)(nil)).Elem())
	var tst = make([]ISkipWhitespaceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkipWhitespaceContext)
		}
	}

	return tst
}

func (s *DocumentationContentContext) SkipWhitespace(i int) ISkipWhitespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipWhitespaceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkipWhitespaceContext)
}

func (s *DocumentationContentContext) TagSection() ITagSectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagSectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagSectionContext)
}

func (s *DocumentationContentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserNEWLINE)
}

func (s *DocumentationContentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, i)
}

func (s *DocumentationContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentationContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentationContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDocumentationContent(s)
	}
}

func (s *DocumentationContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDocumentationContent(s)
	}
}

func (s *DocumentationContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDocumentationContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DocumentationContent() (localctx IDocumentationContentContext) {
	localctx = NewDocumentationContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JavadocParserRULE_documentationContent)
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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Description()
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JavadocParserNEWLINE || _la == JavadocParserSPACE {
			{
				p.SetState(68)
				p.SkipWhitespace()
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(74)
					p.SkipWhitespace()
				}

			}
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}
		{
			p.SetState(80)
			p.TagSection()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.Description()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(82)
					p.Match(JavadocParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(87)
					p.SkipWhitespace()
				}

			}
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}
		{
			p.SetState(93)
			p.TagSection()
		}

	}

	return localctx
}

// ISkipWhitespaceContext is an interface to support dynamic dispatch.
type ISkipWhitespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkipWhitespaceContext differentiates from other interfaces.
	IsSkipWhitespaceContext()
}

type SkipWhitespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipWhitespaceContext() *SkipWhitespaceContext {
	var p = new(SkipWhitespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_skipWhitespace
	return p
}

func (*SkipWhitespaceContext) IsSkipWhitespaceContext() {}

func NewSkipWhitespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipWhitespaceContext {
	var p = new(SkipWhitespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_skipWhitespace

	return p
}

func (s *SkipWhitespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipWhitespaceContext) SPACE() antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, 0)
}

func (s *SkipWhitespaceContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, 0)
}

func (s *SkipWhitespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipWhitespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipWhitespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterSkipWhitespace(s)
	}
}

func (s *SkipWhitespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitSkipWhitespace(s)
	}
}

func (s *SkipWhitespaceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitSkipWhitespace(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) SkipWhitespace() (localctx ISkipWhitespaceContext) {
	localctx = NewSkipWhitespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JavadocParserRULE_skipWhitespace)
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
	p.SetState(97)
	_la = p.GetTokenStream().LA(1)

	if !(_la == JavadocParserNEWLINE || _la == JavadocParserSPACE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) AllDescriptionLine() []IDescriptionLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionLineContext)(nil)).Elem())
	var tst = make([]IDescriptionLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionLineContext)
		}
	}

	return tst
}

func (s *DescriptionContext) DescriptionLine(i int) IDescriptionLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineContext)
}

func (s *DescriptionContext) AllDescriptionNewline() []IDescriptionNewlineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionNewlineContext)(nil)).Elem())
	var tst = make([]IDescriptionNewlineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionNewlineContext)
		}
	}

	return tst
}

func (s *DescriptionContext) DescriptionNewline(i int) IDescriptionNewlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionNewlineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionNewlineContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JavadocParserRULE_description)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.DescriptionLine()
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == JavadocParserNEWLINE {
				{
					p.SetState(100)
					p.DescriptionNewline()
				}

				p.SetState(103)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(105)
				p.DescriptionLine()
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IDescriptionLineContext is an interface to support dynamic dispatch.
type IDescriptionLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineContext differentiates from other interfaces.
	IsDescriptionLineContext()
}

type DescriptionLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineContext() *DescriptionLineContext {
	var p = new(DescriptionLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionLine
	return p
}

func (*DescriptionLineContext) IsDescriptionLineContext() {}

func NewDescriptionLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineContext {
	var p = new(DescriptionLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionLine

	return p
}

func (s *DescriptionLineContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineContext) DescriptionLineStart() IDescriptionLineStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineStartContext)
}

func (s *DescriptionLineContext) AllDescriptionLineElement() []IDescriptionLineElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionLineElementContext)(nil)).Elem())
	var tst = make([]IDescriptionLineElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionLineElementContext)
		}
	}

	return tst
}

func (s *DescriptionLineContext) DescriptionLineElement(i int) IDescriptionLineElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineElementContext)
}

func (s *DescriptionLineContext) InlineTag() IInlineTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineTagContext)
}

func (s *DescriptionLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionLine(s)
	}
}

func (s *DescriptionLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionLine(s)
	}
}

func (s *DescriptionLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionLine() (localctx IDescriptionLineContext) {
	localctx = NewDescriptionLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JavadocParserRULE_descriptionLine)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavadocParserNAME, JavadocParserSPACE, JavadocParserTEXT_CONTENT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.DescriptionLineStart()
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(113)
					p.DescriptionLineElement()
				}

			}
			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}

	case JavadocParserINLINE_TAG_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.InlineTag()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(120)
					p.DescriptionLineElement()
				}

			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDescriptionLineStartContext is an interface to support dynamic dispatch.
type IDescriptionLineStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineStartContext differentiates from other interfaces.
	IsDescriptionLineStartContext()
}

type DescriptionLineStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineStartContext() *DescriptionLineStartContext {
	var p = new(DescriptionLineStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionLineStart
	return p
}

func (*DescriptionLineStartContext) IsDescriptionLineStartContext() {}

func NewDescriptionLineStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineStartContext {
	var p = new(DescriptionLineStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionLineStart

	return p
}

func (s *DescriptionLineStartContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineStartContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserSPACE)
}

func (s *DescriptionLineStartContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, i)
}

func (s *DescriptionLineStartContext) AllDescriptionLineNoSpaceNoAt() []IDescriptionLineNoSpaceNoAtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionLineNoSpaceNoAtContext)(nil)).Elem())
	var tst = make([]IDescriptionLineNoSpaceNoAtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionLineNoSpaceNoAtContext)
		}
	}

	return tst
}

func (s *DescriptionLineStartContext) DescriptionLineNoSpaceNoAt(i int) IDescriptionLineNoSpaceNoAtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineNoSpaceNoAtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineNoSpaceNoAtContext)
}

func (s *DescriptionLineStartContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserAT)
}

func (s *DescriptionLineStartContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserAT, i)
}

func (s *DescriptionLineStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionLineStart(s)
	}
}

func (s *DescriptionLineStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionLineStart(s)
	}
}

func (s *DescriptionLineStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionLineStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionLineStart() (localctx IDescriptionLineStartContext) {
	localctx = NewDescriptionLineStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JavadocParserRULE_descriptionLineStart)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavadocParserSPACE {
		{
			p.SetState(128)
			p.Match(JavadocParserSPACE)
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(131)
				p.DescriptionLineNoSpaceNoAt()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(139)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case JavadocParserNAME, JavadocParserTEXT_CONTENT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
				{
					p.SetState(136)
					p.DescriptionLineNoSpaceNoAt()
				}

			case JavadocParserSPACE:
				{
					p.SetState(137)
					p.Match(JavadocParserSPACE)
				}

			case JavadocParserAT:
				{
					p.SetState(138)
					p.Match(JavadocParserAT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IDescriptionLineNoSpaceNoAtContext is an interface to support dynamic dispatch.
type IDescriptionLineNoSpaceNoAtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineNoSpaceNoAtContext differentiates from other interfaces.
	IsDescriptionLineNoSpaceNoAtContext()
}

type DescriptionLineNoSpaceNoAtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineNoSpaceNoAtContext() *DescriptionLineNoSpaceNoAtContext {
	var p = new(DescriptionLineNoSpaceNoAtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionLineNoSpaceNoAt
	return p
}

func (*DescriptionLineNoSpaceNoAtContext) IsDescriptionLineNoSpaceNoAtContext() {}

func NewDescriptionLineNoSpaceNoAtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineNoSpaceNoAtContext {
	var p = new(DescriptionLineNoSpaceNoAtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionLineNoSpaceNoAt

	return p
}

func (s *DescriptionLineNoSpaceNoAtContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineNoSpaceNoAtContext) TEXT_CONTENT() antlr.TerminalNode {
	return s.GetToken(JavadocParserTEXT_CONTENT, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) NAME() antlr.TerminalNode {
	return s.GetToken(JavadocParserNAME, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) STAR() antlr.TerminalNode {
	return s.GetToken(JavadocParserSTAR, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) SLASH() antlr.TerminalNode {
	return s.GetToken(JavadocParserSLASH, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) BRACE_OPEN() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_OPEN, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) BRACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_CLOSE, 0)
}

func (s *DescriptionLineNoSpaceNoAtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineNoSpaceNoAtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineNoSpaceNoAtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionLineNoSpaceNoAt(s)
	}
}

func (s *DescriptionLineNoSpaceNoAtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionLineNoSpaceNoAt(s)
	}
}

func (s *DescriptionLineNoSpaceNoAtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionLineNoSpaceNoAt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionLineNoSpaceNoAt() (localctx IDescriptionLineNoSpaceNoAtContext) {
	localctx = NewDescriptionLineNoSpaceNoAtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JavadocParserRULE_descriptionLineNoSpaceNoAt)
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
	p.SetState(144)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH)|(1<<JavadocParserBRACE_OPEN)|(1<<JavadocParserBRACE_CLOSE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDescriptionLineElementContext is an interface to support dynamic dispatch.
type IDescriptionLineElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineElementContext differentiates from other interfaces.
	IsDescriptionLineElementContext()
}

type DescriptionLineElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineElementContext() *DescriptionLineElementContext {
	var p = new(DescriptionLineElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionLineElement
	return p
}

func (*DescriptionLineElementContext) IsDescriptionLineElementContext() {}

func NewDescriptionLineElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineElementContext {
	var p = new(DescriptionLineElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionLineElement

	return p
}

func (s *DescriptionLineElementContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineElementContext) InlineTag() IInlineTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineTagContext)
}

func (s *DescriptionLineElementContext) DescriptionLineText() IDescriptionLineTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineTextContext)
}

func (s *DescriptionLineElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionLineElement(s)
	}
}

func (s *DescriptionLineElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionLineElement(s)
	}
}

func (s *DescriptionLineElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionLineElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionLineElement() (localctx IDescriptionLineElementContext) {
	localctx = NewDescriptionLineElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JavadocParserRULE_descriptionLineElement)

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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavadocParserINLINE_TAG_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.InlineTag()
		}

	case JavadocParserNAME, JavadocParserSPACE, JavadocParserTEXT_CONTENT, JavadocParserAT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.DescriptionLineText()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDescriptionLineTextContext is an interface to support dynamic dispatch.
type IDescriptionLineTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineTextContext differentiates from other interfaces.
	IsDescriptionLineTextContext()
}

type DescriptionLineTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineTextContext() *DescriptionLineTextContext {
	var p = new(DescriptionLineTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionLineText
	return p
}

func (*DescriptionLineTextContext) IsDescriptionLineTextContext() {}

func NewDescriptionLineTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineTextContext {
	var p = new(DescriptionLineTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionLineText

	return p
}

func (s *DescriptionLineTextContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineTextContext) AllDescriptionLineNoSpaceNoAt() []IDescriptionLineNoSpaceNoAtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionLineNoSpaceNoAtContext)(nil)).Elem())
	var tst = make([]IDescriptionLineNoSpaceNoAtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionLineNoSpaceNoAtContext)
		}
	}

	return tst
}

func (s *DescriptionLineTextContext) DescriptionLineNoSpaceNoAt(i int) IDescriptionLineNoSpaceNoAtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineNoSpaceNoAtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineNoSpaceNoAtContext)
}

func (s *DescriptionLineTextContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserSPACE)
}

func (s *DescriptionLineTextContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, i)
}

func (s *DescriptionLineTextContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserAT)
}

func (s *DescriptionLineTextContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserAT, i)
}

func (s *DescriptionLineTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionLineText(s)
	}
}

func (s *DescriptionLineTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionLineText(s)
	}
}

func (s *DescriptionLineTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionLineText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionLineText() (localctx IDescriptionLineTextContext) {
	localctx = NewDescriptionLineTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JavadocParserRULE_descriptionLineText)

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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(153)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case JavadocParserNAME, JavadocParserTEXT_CONTENT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
				{
					p.SetState(150)
					p.DescriptionLineNoSpaceNoAt()
				}

			case JavadocParserSPACE:
				{
					p.SetState(151)
					p.Match(JavadocParserSPACE)
				}

			case JavadocParserAT:
				{
					p.SetState(152)
					p.Match(JavadocParserAT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IDescriptionNewlineContext is an interface to support dynamic dispatch.
type IDescriptionNewlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionNewlineContext differentiates from other interfaces.
	IsDescriptionNewlineContext()
}

type DescriptionNewlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionNewlineContext() *DescriptionNewlineContext {
	var p = new(DescriptionNewlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_descriptionNewline
	return p
}

func (*DescriptionNewlineContext) IsDescriptionNewlineContext() {}

func NewDescriptionNewlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionNewlineContext {
	var p = new(DescriptionNewlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_descriptionNewline

	return p
}

func (s *DescriptionNewlineContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionNewlineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, 0)
}

func (s *DescriptionNewlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionNewlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionNewlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterDescriptionNewline(s)
	}
}

func (s *DescriptionNewlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitDescriptionNewline(s)
	}
}

func (s *DescriptionNewlineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitDescriptionNewline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) DescriptionNewline() (localctx IDescriptionNewlineContext) {
	localctx = NewDescriptionNewlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JavadocParserRULE_descriptionNewline)

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
		p.Match(JavadocParserNEWLINE)
	}

	return localctx
}

// ITagSectionContext is an interface to support dynamic dispatch.
type ITagSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagSectionContext differentiates from other interfaces.
	IsTagSectionContext()
}

type TagSectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagSectionContext() *TagSectionContext {
	var p = new(TagSectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_tagSection
	return p
}

func (*TagSectionContext) IsTagSectionContext() {}

func NewTagSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagSectionContext {
	var p = new(TagSectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_tagSection

	return p
}

func (s *TagSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *TagSectionContext) AllBlockTag() []IBlockTagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockTagContext)(nil)).Elem())
	var tst = make([]IBlockTagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockTagContext)
		}
	}

	return tst
}

func (s *TagSectionContext) BlockTag(i int) IBlockTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockTagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockTagContext)
}

func (s *TagSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterTagSection(s)
	}
}

func (s *TagSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitTagSection(s)
	}
}

func (s *TagSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitTagSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) TagSection() (localctx ITagSectionContext) {
	localctx = NewTagSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JavadocParserRULE_tagSection)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == JavadocParserSPACE || _la == JavadocParserAT {
		{
			p.SetState(159)
			p.BlockTag()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockTagContext is an interface to support dynamic dispatch.
type IBlockTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockTagContext differentiates from other interfaces.
	IsBlockTagContext()
}

type BlockTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockTagContext() *BlockTagContext {
	var p = new(BlockTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_blockTag
	return p
}

func (*BlockTagContext) IsBlockTagContext() {}

func NewBlockTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockTagContext {
	var p = new(BlockTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_blockTag

	return p
}

func (s *BlockTagContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockTagContext) AT() antlr.TerminalNode {
	return s.GetToken(JavadocParserAT, 0)
}

func (s *BlockTagContext) BlockTagName() IBlockTagNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockTagNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockTagNameContext)
}

func (s *BlockTagContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserSPACE)
}

func (s *BlockTagContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, i)
}

func (s *BlockTagContext) AllBlockTagContent() []IBlockTagContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockTagContentContext)(nil)).Elem())
	var tst = make([]IBlockTagContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockTagContentContext)
		}
	}

	return tst
}

func (s *BlockTagContext) BlockTagContent(i int) IBlockTagContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockTagContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockTagContentContext)
}

func (s *BlockTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBlockTag(s)
	}
}

func (s *BlockTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBlockTag(s)
	}
}

func (s *BlockTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBlockTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BlockTag() (localctx IBlockTagContext) {
	localctx = NewBlockTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JavadocParserRULE_blockTag)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JavadocParserSPACE {
		{
			p.SetState(164)
			p.Match(JavadocParserSPACE)
		}

	}
	{
		p.SetState(167)
		p.Match(JavadocParserAT)
	}
	{
		p.SetState(168)
		p.BlockTagName()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(169)
			p.Match(JavadocParserSPACE)
		}

	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				p.BlockTagContent()
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockTagNameContext is an interface to support dynamic dispatch.
type IBlockTagNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockTagNameContext differentiates from other interfaces.
	IsBlockTagNameContext()
}

type BlockTagNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockTagNameContext() *BlockTagNameContext {
	var p = new(BlockTagNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_blockTagName
	return p
}

func (*BlockTagNameContext) IsBlockTagNameContext() {}

func NewBlockTagNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockTagNameContext {
	var p = new(BlockTagNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_blockTagName

	return p
}

func (s *BlockTagNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockTagNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(JavadocParserNAME, 0)
}

func (s *BlockTagNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTagNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockTagNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBlockTagName(s)
	}
}

func (s *BlockTagNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBlockTagName(s)
	}
}

func (s *BlockTagNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBlockTagName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BlockTagName() (localctx IBlockTagNameContext) {
	localctx = NewBlockTagNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JavadocParserRULE_blockTagName)

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
		p.SetState(178)
		p.Match(JavadocParserNAME)
	}

	return localctx
}

// IBlockTagContentContext is an interface to support dynamic dispatch.
type IBlockTagContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockTagContentContext differentiates from other interfaces.
	IsBlockTagContentContext()
}

type BlockTagContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockTagContentContext() *BlockTagContentContext {
	var p = new(BlockTagContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_blockTagContent
	return p
}

func (*BlockTagContentContext) IsBlockTagContentContext() {}

func NewBlockTagContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockTagContentContext {
	var p = new(BlockTagContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_blockTagContent

	return p
}

func (s *BlockTagContentContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockTagContentContext) BlockTagText() IBlockTagTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockTagTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockTagTextContext)
}

func (s *BlockTagContentContext) InlineTag() IInlineTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineTagContext)
}

func (s *BlockTagContentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, 0)
}

func (s *BlockTagContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTagContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockTagContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBlockTagContent(s)
	}
}

func (s *BlockTagContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBlockTagContent(s)
	}
}

func (s *BlockTagContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBlockTagContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BlockTagContent() (localctx IBlockTagContentContext) {
	localctx = NewBlockTagContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JavadocParserRULE_blockTagContent)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavadocParserNAME, JavadocParserSPACE, JavadocParserTEXT_CONTENT, JavadocParserSTAR, JavadocParserSLASH, JavadocParserBRACE_OPEN, JavadocParserBRACE_CLOSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.BlockTagText()
		}

	case JavadocParserINLINE_TAG_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.InlineTag()
		}

	case JavadocParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(JavadocParserNEWLINE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockTagTextContext is an interface to support dynamic dispatch.
type IBlockTagTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockTagTextContext differentiates from other interfaces.
	IsBlockTagTextContext()
}

type BlockTagTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockTagTextContext() *BlockTagTextContext {
	var p = new(BlockTagTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_blockTagText
	return p
}

func (*BlockTagTextContext) IsBlockTagTextContext() {}

func NewBlockTagTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockTagTextContext {
	var p = new(BlockTagTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_blockTagText

	return p
}

func (s *BlockTagTextContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockTagTextContext) AllBlockTagTextElement() []IBlockTagTextElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockTagTextElementContext)(nil)).Elem())
	var tst = make([]IBlockTagTextElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockTagTextElementContext)
		}
	}

	return tst
}

func (s *BlockTagTextContext) BlockTagTextElement(i int) IBlockTagTextElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockTagTextElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockTagTextElementContext)
}

func (s *BlockTagTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTagTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockTagTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBlockTagText(s)
	}
}

func (s *BlockTagTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBlockTagText(s)
	}
}

func (s *BlockTagTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBlockTagText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BlockTagText() (localctx IBlockTagTextContext) {
	localctx = NewBlockTagTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JavadocParserRULE_blockTagText)

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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(185)
				p.BlockTagTextElement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockTagTextElementContext is an interface to support dynamic dispatch.
type IBlockTagTextElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockTagTextElementContext differentiates from other interfaces.
	IsBlockTagTextElementContext()
}

type BlockTagTextElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockTagTextElementContext() *BlockTagTextElementContext {
	var p = new(BlockTagTextElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_blockTagTextElement
	return p
}

func (*BlockTagTextElementContext) IsBlockTagTextElementContext() {}

func NewBlockTagTextElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockTagTextElementContext {
	var p = new(BlockTagTextElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_blockTagTextElement

	return p
}

func (s *BlockTagTextElementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockTagTextElementContext) TEXT_CONTENT() antlr.TerminalNode {
	return s.GetToken(JavadocParserTEXT_CONTENT, 0)
}

func (s *BlockTagTextElementContext) NAME() antlr.TerminalNode {
	return s.GetToken(JavadocParserNAME, 0)
}

func (s *BlockTagTextElementContext) SPACE() antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, 0)
}

func (s *BlockTagTextElementContext) STAR() antlr.TerminalNode {
	return s.GetToken(JavadocParserSTAR, 0)
}

func (s *BlockTagTextElementContext) SLASH() antlr.TerminalNode {
	return s.GetToken(JavadocParserSLASH, 0)
}

func (s *BlockTagTextElementContext) BRACE_OPEN() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_OPEN, 0)
}

func (s *BlockTagTextElementContext) BRACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_CLOSE, 0)
}

func (s *BlockTagTextElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockTagTextElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockTagTextElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBlockTagTextElement(s)
	}
}

func (s *BlockTagTextElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBlockTagTextElement(s)
	}
}

func (s *BlockTagTextElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBlockTagTextElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BlockTagTextElement() (localctx IBlockTagTextElementContext) {
	localctx = NewBlockTagTextElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JavadocParserRULE_blockTagTextElement)
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
	p.SetState(190)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserSPACE)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH)|(1<<JavadocParserBRACE_OPEN)|(1<<JavadocParserBRACE_CLOSE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IInlineTagContext is an interface to support dynamic dispatch.
type IInlineTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineTagContext differentiates from other interfaces.
	IsInlineTagContext()
}

type InlineTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTagContext() *InlineTagContext {
	var p = new(InlineTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_inlineTag
	return p
}

func (*InlineTagContext) IsInlineTagContext() {}

func NewInlineTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTagContext {
	var p = new(InlineTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_inlineTag

	return p
}

func (s *InlineTagContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTagContext) INLINE_TAG_START() antlr.TerminalNode {
	return s.GetToken(JavadocParserINLINE_TAG_START, 0)
}

func (s *InlineTagContext) InlineTagName() IInlineTagNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineTagNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineTagNameContext)
}

func (s *InlineTagContext) BRACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_CLOSE, 0)
}

func (s *InlineTagContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserSPACE)
}

func (s *InlineTagContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, i)
}

func (s *InlineTagContext) InlineTagContent() IInlineTagContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineTagContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineTagContentContext)
}

func (s *InlineTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterInlineTag(s)
	}
}

func (s *InlineTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitInlineTag(s)
	}
}

func (s *InlineTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitInlineTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) InlineTag() (localctx IInlineTagContext) {
	localctx = NewInlineTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JavadocParserRULE_inlineTag)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(JavadocParserINLINE_TAG_START)
	}
	{
		p.SetState(193)
		p.InlineTagName()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(194)
				p.Match(JavadocParserSPACE)
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserNEWLINE)|(1<<JavadocParserSPACE)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH)|(1<<JavadocParserBRACE_OPEN))) != 0 {
		{
			p.SetState(200)
			p.InlineTagContent()
		}

	}
	{
		p.SetState(203)
		p.Match(JavadocParserBRACE_CLOSE)
	}

	return localctx
}

// IInlineTagNameContext is an interface to support dynamic dispatch.
type IInlineTagNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineTagNameContext differentiates from other interfaces.
	IsInlineTagNameContext()
}

type InlineTagNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTagNameContext() *InlineTagNameContext {
	var p = new(InlineTagNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_inlineTagName
	return p
}

func (*InlineTagNameContext) IsInlineTagNameContext() {}

func NewInlineTagNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTagNameContext {
	var p = new(InlineTagNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_inlineTagName

	return p
}

func (s *InlineTagNameContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTagNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(JavadocParserNAME, 0)
}

func (s *InlineTagNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTagNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTagNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterInlineTagName(s)
	}
}

func (s *InlineTagNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitInlineTagName(s)
	}
}

func (s *InlineTagNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitInlineTagName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) InlineTagName() (localctx IInlineTagNameContext) {
	localctx = NewInlineTagNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JavadocParserRULE_inlineTagName)

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
		p.SetState(205)
		p.Match(JavadocParserNAME)
	}

	return localctx
}

// IInlineTagContentContext is an interface to support dynamic dispatch.
type IInlineTagContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineTagContentContext differentiates from other interfaces.
	IsInlineTagContentContext()
}

type InlineTagContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTagContentContext() *InlineTagContentContext {
	var p = new(InlineTagContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_inlineTagContent
	return p
}

func (*InlineTagContentContext) IsInlineTagContentContext() {}

func NewInlineTagContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTagContentContext {
	var p = new(InlineTagContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_inlineTagContent

	return p
}

func (s *InlineTagContentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTagContentContext) AllBraceContent() []IBraceContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBraceContentContext)(nil)).Elem())
	var tst = make([]IBraceContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBraceContentContext)
		}
	}

	return tst
}

func (s *InlineTagContentContext) BraceContent(i int) IBraceContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBraceContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBraceContentContext)
}

func (s *InlineTagContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTagContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTagContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterInlineTagContent(s)
	}
}

func (s *InlineTagContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitInlineTagContent(s)
	}
}

func (s *InlineTagContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitInlineTagContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) InlineTagContent() (localctx IInlineTagContentContext) {
	localctx = NewInlineTagContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JavadocParserRULE_inlineTagContent)
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
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserNEWLINE)|(1<<JavadocParserSPACE)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH)|(1<<JavadocParserBRACE_OPEN))) != 0) {
		{
			p.SetState(207)
			p.BraceContent()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBraceExpressionContext is an interface to support dynamic dispatch.
type IBraceExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBraceExpressionContext differentiates from other interfaces.
	IsBraceExpressionContext()
}

type BraceExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBraceExpressionContext() *BraceExpressionContext {
	var p = new(BraceExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_braceExpression
	return p
}

func (*BraceExpressionContext) IsBraceExpressionContext() {}

func NewBraceExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BraceExpressionContext {
	var p = new(BraceExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_braceExpression

	return p
}

func (s *BraceExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BraceExpressionContext) BRACE_OPEN() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_OPEN, 0)
}

func (s *BraceExpressionContext) BRACE_CLOSE() antlr.TerminalNode {
	return s.GetToken(JavadocParserBRACE_CLOSE, 0)
}

func (s *BraceExpressionContext) AllBraceContent() []IBraceContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBraceContentContext)(nil)).Elem())
	var tst = make([]IBraceContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBraceContentContext)
		}
	}

	return tst
}

func (s *BraceExpressionContext) BraceContent(i int) IBraceContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBraceContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBraceContentContext)
}

func (s *BraceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BraceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBraceExpression(s)
	}
}

func (s *BraceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBraceExpression(s)
	}
}

func (s *BraceExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBraceExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BraceExpression() (localctx IBraceExpressionContext) {
	localctx = NewBraceExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JavadocParserRULE_braceExpression)
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
		p.SetState(212)
		p.Match(JavadocParserBRACE_OPEN)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserNEWLINE)|(1<<JavadocParserSPACE)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH)|(1<<JavadocParserBRACE_OPEN))) != 0 {
		{
			p.SetState(213)
			p.BraceContent()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(JavadocParserBRACE_CLOSE)
	}

	return localctx
}

// IBraceContentContext is an interface to support dynamic dispatch.
type IBraceContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBraceContentContext differentiates from other interfaces.
	IsBraceContentContext()
}

type BraceContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBraceContentContext() *BraceContentContext {
	var p = new(BraceContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_braceContent
	return p
}

func (*BraceContentContext) IsBraceContentContext() {}

func NewBraceContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BraceContentContext {
	var p = new(BraceContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_braceContent

	return p
}

func (s *BraceContentContext) GetParser() antlr.Parser { return s.parser }

func (s *BraceContentContext) BraceExpression() IBraceExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBraceExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBraceExpressionContext)
}

func (s *BraceContentContext) AllBraceText() []IBraceTextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBraceTextContext)(nil)).Elem())
	var tst = make([]IBraceTextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBraceTextContext)
		}
	}

	return tst
}

func (s *BraceContentContext) BraceText(i int) IBraceTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBraceTextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBraceTextContext)
}

func (s *BraceContentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(JavadocParserNEWLINE)
}

func (s *BraceContentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, i)
}

func (s *BraceContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BraceContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBraceContent(s)
	}
}

func (s *BraceContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBraceContent(s)
	}
}

func (s *BraceContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBraceContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BraceContent() (localctx IBraceContentContext) {
	localctx = NewBraceContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JavadocParserRULE_braceContent)

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

	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JavadocParserBRACE_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.BraceExpression()
		}

	case JavadocParserNAME, JavadocParserNEWLINE, JavadocParserSPACE, JavadocParserTEXT_CONTENT, JavadocParserSTAR, JavadocParserSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.BraceText()
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(226)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(223)
							p.Match(JavadocParserNEWLINE)
						}

					}
					p.SetState(228)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
				}
				{
					p.SetState(229)
					p.BraceText()
				}

			}
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBraceTextContext is an interface to support dynamic dispatch.
type IBraceTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBraceTextContext differentiates from other interfaces.
	IsBraceTextContext()
}

type BraceTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBraceTextContext() *BraceTextContext {
	var p = new(BraceTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JavadocParserRULE_braceText
	return p
}

func (*BraceTextContext) IsBraceTextContext() {}

func NewBraceTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BraceTextContext {
	var p = new(BraceTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JavadocParserRULE_braceText

	return p
}

func (s *BraceTextContext) GetParser() antlr.Parser { return s.parser }

func (s *BraceTextContext) TEXT_CONTENT() antlr.TerminalNode {
	return s.GetToken(JavadocParserTEXT_CONTENT, 0)
}

func (s *BraceTextContext) NAME() antlr.TerminalNode {
	return s.GetToken(JavadocParserNAME, 0)
}

func (s *BraceTextContext) SPACE() antlr.TerminalNode {
	return s.GetToken(JavadocParserSPACE, 0)
}

func (s *BraceTextContext) STAR() antlr.TerminalNode {
	return s.GetToken(JavadocParserSTAR, 0)
}

func (s *BraceTextContext) SLASH() antlr.TerminalNode {
	return s.GetToken(JavadocParserSLASH, 0)
}

func (s *BraceTextContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(JavadocParserNEWLINE, 0)
}

func (s *BraceTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BraceTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.EnterBraceText(s)
	}
}

func (s *BraceTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JavadocParserListener); ok {
		listenerT.ExitBraceText(s)
	}
}

func (s *BraceTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JavadocParserVisitor:
		return t.VisitBraceText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JavadocParser) BraceText() (localctx IBraceTextContext) {
	localctx = NewBraceTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JavadocParserRULE_braceText)
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
	p.SetState(237)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JavadocParserNAME)|(1<<JavadocParserNEWLINE)|(1<<JavadocParserSPACE)|(1<<JavadocParserTEXT_CONTENT)|(1<<JavadocParserSTAR)|(1<<JavadocParserSLASH))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
