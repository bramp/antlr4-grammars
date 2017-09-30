// Generated from clf.g4 by ANTLR 4.7.

package clf // clf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by clfParser.
type clfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by clfParser#log.
	VisitLog(ctx *LogContext) interface{}

	// Visit a parse tree produced by clfParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by clfParser#host.
	VisitHost(ctx *HostContext) interface{}

	// Visit a parse tree produced by clfParser#logname.
	VisitLogname(ctx *LognameContext) interface{}

	// Visit a parse tree produced by clfParser#username.
	VisitUsername(ctx *UsernameContext) interface{}

	// Visit a parse tree produced by clfParser#datetimetz.
	VisitDatetimetz(ctx *DatetimetzContext) interface{}

	// Visit a parse tree produced by clfParser#referer.
	VisitReferer(ctx *RefererContext) interface{}

	// Visit a parse tree produced by clfParser#request.
	VisitRequest(ctx *RequestContext) interface{}

	// Visit a parse tree produced by clfParser#useragent.
	VisitUseragent(ctx *UseragentContext) interface{}

	// Visit a parse tree produced by clfParser#statuscode.
	VisitStatuscode(ctx *StatuscodeContext) interface{}

	// Visit a parse tree produced by clfParser#bytes.
	VisitBytes(ctx *BytesContext) interface{}
}
