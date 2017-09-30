// Generated from gml.g4 by ANTLR 4.7.

package gml // gml
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by gmlParser.
type gmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by gmlParser#graph.
	VisitGraph(ctx *GraphContext) interface{}

	// Visit a parse tree produced by gmlParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by gmlParser#kv.
	VisitKv(ctx *KvContext) interface{}

	// Visit a parse tree produced by gmlParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by gmlParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by gmlParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by gmlParser#realnum.
	VisitRealnum(ctx *RealnumContext) interface{}

	// Visit a parse tree produced by gmlParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by gmlParser#stringliteral.
	VisitStringliteral(ctx *StringliteralContext) interface{}
}
