// Generated from emailaddress.g4 by ANTLR 4.7.

package emailaddress // emailaddress
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseemailaddressVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseemailaddressVisitor) VisitEmailaddress(ctx *EmailaddressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitMailbox(ctx *MailboxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitRouteaddr(ctx *RouteaddrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitRoute(ctx *RouteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitAddrspec(ctx *AddrspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitLocalpart(ctx *LocalpartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitDomain(ctx *DomainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitSubdomain(ctx *SubdomainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitDomainref(ctx *DomainrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitPhrase(ctx *PhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitWord(ctx *WordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitLwspchar(ctx *LwspcharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitLwsp(ctx *LwspContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitDelimeters(ctx *DelimetersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitQuotedpair(ctx *QuotedpairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitDomainliteral(ctx *DomainliteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitQuotedstring(ctx *QuotedstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseemailaddressVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
