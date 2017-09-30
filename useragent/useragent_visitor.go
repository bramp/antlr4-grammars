// Generated from useragent.g4 by ANTLR 4.7.

package useragent // useragent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by useragentParser.
type useragentVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by useragentParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by useragentParser#product.
	VisitProduct(ctx *ProductContext) interface{}

	// Visit a parse tree produced by useragentParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by useragentParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by useragentParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
