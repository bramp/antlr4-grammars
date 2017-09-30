// Generated from lcc.g4 by ANTLR 4.7.

package lcc // lcc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by lccParser.
type lccVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by lccParser#lcc.
	VisitLcc(ctx *LccContext) interface{}

	// Visit a parse tree produced by lccParser#topic.
	VisitTopic(ctx *TopicContext) interface{}

	// Visit a parse tree produced by lccParser#subtopic.
	VisitSubtopic(ctx *SubtopicContext) interface{}

	// Visit a parse tree produced by lccParser#subclasses.
	VisitSubclasses(ctx *SubclassesContext) interface{}

	// Visit a parse tree produced by lccParser#subclass.
	VisitSubclass(ctx *SubclassContext) interface{}

	// Visit a parse tree produced by lccParser#cutters.
	VisitCutters(ctx *CuttersContext) interface{}

	// Visit a parse tree produced by lccParser#cutter.
	VisitCutter(ctx *CutterContext) interface{}

	// Visit a parse tree produced by lccParser#date.
	VisitDate(ctx *DateContext) interface{}
}
