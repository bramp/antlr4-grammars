// Generated from cookie.g4 by ANTLR 4.7.

package cookie // cookie
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by cookieParser.
type cookieVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by cookieParser#cookie.
	VisitCookie(ctx *CookieContext) interface{}

	// Visit a parse tree produced by cookieParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by cookieParser#av_pairs.
	VisitAv_pairs(ctx *Av_pairsContext) interface{}

	// Visit a parse tree produced by cookieParser#av_pair.
	VisitAv_pair(ctx *Av_pairContext) interface{}

	// Visit a parse tree produced by cookieParser#attr.
	VisitAttr(ctx *AttrContext) interface{}

	// Visit a parse tree produced by cookieParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by cookieParser#word.
	VisitWord(ctx *WordContext) interface{}

	// Visit a parse tree produced by cookieParser#token.
	VisitToken(ctx *TokenContext) interface{}

	// Visit a parse tree produced by cookieParser#quoted_string.
	VisitQuoted_string(ctx *Quoted_stringContext) interface{}
}
