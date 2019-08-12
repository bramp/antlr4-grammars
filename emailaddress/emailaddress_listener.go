// Code generated from emailaddress.g4 by ANTLR 4.7.2. DO NOT EDIT.

package emailaddress // emailaddress
import "github.com/antlr/antlr4/runtime/Go/antlr"

// emailaddressListener is a complete listener for a parse tree produced by emailaddressParser.
type emailaddressListener interface {
	antlr.ParseTreeListener

	// EnterEmailaddress is called when entering the emailaddress production.
	EnterEmailaddress(c *EmailaddressContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterMailbox is called when entering the mailbox production.
	EnterMailbox(c *MailboxContext)

	// EnterRouteaddr is called when entering the routeaddr production.
	EnterRouteaddr(c *RouteaddrContext)

	// EnterRoute is called when entering the route production.
	EnterRoute(c *RouteContext)

	// EnterAddrspec is called when entering the addrspec production.
	EnterAddrspec(c *AddrspecContext)

	// EnterLocalpart is called when entering the localpart production.
	EnterLocalpart(c *LocalpartContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterSubdomain is called when entering the subdomain production.
	EnterSubdomain(c *SubdomainContext)

	// EnterDomainref is called when entering the domainref production.
	EnterDomainref(c *DomainrefContext)

	// EnterPhrase is called when entering the phrase production.
	EnterPhrase(c *PhraseContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterLwspchar is called when entering the lwspchar production.
	EnterLwspchar(c *LwspcharContext)

	// EnterLwsp is called when entering the lwsp production.
	EnterLwsp(c *LwspContext)

	// EnterDelimeters is called when entering the delimeters production.
	EnterDelimeters(c *DelimetersContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterQuotedpair is called when entering the quotedpair production.
	EnterQuotedpair(c *QuotedpairContext)

	// EnterDomainliteral is called when entering the domainliteral production.
	EnterDomainliteral(c *DomainliteralContext)

	// EnterQuotedstring is called when entering the quotedstring production.
	EnterQuotedstring(c *QuotedstringContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitEmailaddress is called when exiting the emailaddress production.
	ExitEmailaddress(c *EmailaddressContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitMailbox is called when exiting the mailbox production.
	ExitMailbox(c *MailboxContext)

	// ExitRouteaddr is called when exiting the routeaddr production.
	ExitRouteaddr(c *RouteaddrContext)

	// ExitRoute is called when exiting the route production.
	ExitRoute(c *RouteContext)

	// ExitAddrspec is called when exiting the addrspec production.
	ExitAddrspec(c *AddrspecContext)

	// ExitLocalpart is called when exiting the localpart production.
	ExitLocalpart(c *LocalpartContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitSubdomain is called when exiting the subdomain production.
	ExitSubdomain(c *SubdomainContext)

	// ExitDomainref is called when exiting the domainref production.
	ExitDomainref(c *DomainrefContext)

	// ExitPhrase is called when exiting the phrase production.
	ExitPhrase(c *PhraseContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitLwspchar is called when exiting the lwspchar production.
	ExitLwspchar(c *LwspcharContext)

	// ExitLwsp is called when exiting the lwsp production.
	ExitLwsp(c *LwspContext)

	// ExitDelimeters is called when exiting the delimeters production.
	ExitDelimeters(c *DelimetersContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitQuotedpair is called when exiting the quotedpair production.
	ExitQuotedpair(c *QuotedpairContext)

	// ExitDomainliteral is called when exiting the domainliteral production.
	ExitDomainliteral(c *DomainliteralContext)

	// ExitQuotedstring is called when exiting the quotedstring production.
	ExitQuotedstring(c *QuotedstringContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
