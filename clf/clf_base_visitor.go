// Generated from clf.g4 by ANTLR 4.7.

package clf // clf
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseclfVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseclfVisitor) VisitLog(ctx *LogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitHost(ctx *HostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitLogname(ctx *LognameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitUsername(ctx *UsernameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitDatetimetz(ctx *DatetimetzContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitReferer(ctx *RefererContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitRequest(ctx *RequestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitUseragent(ctx *UseragentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitStatuscode(ctx *StatuscodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseclfVisitor) VisitBytes(ctx *BytesContext) interface{} {
	return v.VisitChildren(ctx)
}
