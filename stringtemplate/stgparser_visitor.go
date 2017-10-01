// Generated from STGParser.g4 by ANTLR 4.7.

package stringtemplate // STGParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by STGParser.
type STGParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by STGParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by STGParser#delimiters.
	VisitDelimiters(ctx *DelimitersContext) interface{}

	// Visit a parse tree produced by STGParser#imports.
	VisitImports(ctx *ImportsContext) interface{}

	// Visit a parse tree produced by STGParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by STGParser#formalArgs.
	VisitFormalArgs(ctx *FormalArgsContext) interface{}

	// Visit a parse tree produced by STGParser#formalArg.
	VisitFormalArg(ctx *FormalArgContext) interface{}

	// Visit a parse tree produced by STGParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by STGParser#dictPairs.
	VisitDictPairs(ctx *DictPairsContext) interface{}

	// Visit a parse tree produced by STGParser#keyValuePair.
	VisitKeyValuePair(ctx *KeyValuePairContext) interface{}

	// Visit a parse tree produced by STGParser#defaultValuePair.
	VisitDefaultValuePair(ctx *DefaultValuePairContext) interface{}

	// Visit a parse tree produced by STGParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}
}
