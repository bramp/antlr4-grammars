// Generated from emailaddress.g4 by ANTLR 4.7.

package emailaddress // emailaddress
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by emailaddressParser.
type emailaddressVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by emailaddressParser#emailaddress.
	VisitEmailaddress(ctx *EmailaddressContext) interface{}

	// Visit a parse tree produced by emailaddressParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by emailaddressParser#mailbox.
	VisitMailbox(ctx *MailboxContext) interface{}

	// Visit a parse tree produced by emailaddressParser#routeaddr.
	VisitRouteaddr(ctx *RouteaddrContext) interface{}

	// Visit a parse tree produced by emailaddressParser#route.
	VisitRoute(ctx *RouteContext) interface{}

	// Visit a parse tree produced by emailaddressParser#addrspec.
	VisitAddrspec(ctx *AddrspecContext) interface{}

	// Visit a parse tree produced by emailaddressParser#localpart.
	VisitLocalpart(ctx *LocalpartContext) interface{}

	// Visit a parse tree produced by emailaddressParser#domain.
	VisitDomain(ctx *DomainContext) interface{}

	// Visit a parse tree produced by emailaddressParser#subdomain.
	VisitSubdomain(ctx *SubdomainContext) interface{}

	// Visit a parse tree produced by emailaddressParser#domainref.
	VisitDomainref(ctx *DomainrefContext) interface{}

	// Visit a parse tree produced by emailaddressParser#phrase.
	VisitPhrase(ctx *PhraseContext) interface{}

	// Visit a parse tree produced by emailaddressParser#word.
	VisitWord(ctx *WordContext) interface{}

	// Visit a parse tree produced by emailaddressParser#lwspchar.
	VisitLwspchar(ctx *LwspcharContext) interface{}

	// Visit a parse tree produced by emailaddressParser#lwsp.
	VisitLwsp(ctx *LwspContext) interface{}

	// Visit a parse tree produced by emailaddressParser#delimeters.
	VisitDelimeters(ctx *DelimetersContext) interface{}

	// Visit a parse tree produced by emailaddressParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by emailaddressParser#quotedpair.
	VisitQuotedpair(ctx *QuotedpairContext) interface{}

	// Visit a parse tree produced by emailaddressParser#domainliteral.
	VisitDomainliteral(ctx *DomainliteralContext) interface{}

	// Visit a parse tree produced by emailaddressParser#quotedstring.
	VisitQuotedstring(ctx *QuotedstringContext) interface{}

	// Visit a parse tree produced by emailaddressParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
