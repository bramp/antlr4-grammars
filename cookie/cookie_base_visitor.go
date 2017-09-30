// Generated from cookie.g4 by ANTLR 4.7.

package cookie // cookie
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasecookieVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecookieVisitor) VisitCookie(ctx *CookieContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitAv_pairs(ctx *Av_pairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitAv_pair(ctx *Av_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitAttr(ctx *AttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitWord(ctx *WordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitToken(ctx *TokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecookieVisitor) VisitQuoted_string(ctx *Quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}
