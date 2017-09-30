// Generated from emailaddress.g4 by ANTLR 4.7.

package emailaddress // emailaddress
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseemailaddressListener is a complete listener for a parse tree produced by emailaddressParser.
type BaseemailaddressListener struct{}

var _ emailaddressListener = &BaseemailaddressListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseemailaddressListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseemailaddressListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseemailaddressListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseemailaddressListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEmailaddress is called when production emailaddress is entered.
func (s *BaseemailaddressListener) EnterEmailaddress(ctx *EmailaddressContext) {}

// ExitEmailaddress is called when production emailaddress is exited.
func (s *BaseemailaddressListener) ExitEmailaddress(ctx *EmailaddressContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseemailaddressListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseemailaddressListener) ExitGroup(ctx *GroupContext) {}

// EnterMailbox is called when production mailbox is entered.
func (s *BaseemailaddressListener) EnterMailbox(ctx *MailboxContext) {}

// ExitMailbox is called when production mailbox is exited.
func (s *BaseemailaddressListener) ExitMailbox(ctx *MailboxContext) {}

// EnterRouteaddr is called when production routeaddr is entered.
func (s *BaseemailaddressListener) EnterRouteaddr(ctx *RouteaddrContext) {}

// ExitRouteaddr is called when production routeaddr is exited.
func (s *BaseemailaddressListener) ExitRouteaddr(ctx *RouteaddrContext) {}

// EnterRoute is called when production route is entered.
func (s *BaseemailaddressListener) EnterRoute(ctx *RouteContext) {}

// ExitRoute is called when production route is exited.
func (s *BaseemailaddressListener) ExitRoute(ctx *RouteContext) {}

// EnterAddrspec is called when production addrspec is entered.
func (s *BaseemailaddressListener) EnterAddrspec(ctx *AddrspecContext) {}

// ExitAddrspec is called when production addrspec is exited.
func (s *BaseemailaddressListener) ExitAddrspec(ctx *AddrspecContext) {}

// EnterLocalpart is called when production localpart is entered.
func (s *BaseemailaddressListener) EnterLocalpart(ctx *LocalpartContext) {}

// ExitLocalpart is called when production localpart is exited.
func (s *BaseemailaddressListener) ExitLocalpart(ctx *LocalpartContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseemailaddressListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseemailaddressListener) ExitDomain(ctx *DomainContext) {}

// EnterSubdomain is called when production subdomain is entered.
func (s *BaseemailaddressListener) EnterSubdomain(ctx *SubdomainContext) {}

// ExitSubdomain is called when production subdomain is exited.
func (s *BaseemailaddressListener) ExitSubdomain(ctx *SubdomainContext) {}

// EnterDomainref is called when production domainref is entered.
func (s *BaseemailaddressListener) EnterDomainref(ctx *DomainrefContext) {}

// ExitDomainref is called when production domainref is exited.
func (s *BaseemailaddressListener) ExitDomainref(ctx *DomainrefContext) {}

// EnterPhrase is called when production phrase is entered.
func (s *BaseemailaddressListener) EnterPhrase(ctx *PhraseContext) {}

// ExitPhrase is called when production phrase is exited.
func (s *BaseemailaddressListener) ExitPhrase(ctx *PhraseContext) {}

// EnterWord is called when production word is entered.
func (s *BaseemailaddressListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseemailaddressListener) ExitWord(ctx *WordContext) {}

// EnterLwspchar is called when production lwspchar is entered.
func (s *BaseemailaddressListener) EnterLwspchar(ctx *LwspcharContext) {}

// ExitLwspchar is called when production lwspchar is exited.
func (s *BaseemailaddressListener) ExitLwspchar(ctx *LwspcharContext) {}

// EnterLwsp is called when production lwsp is entered.
func (s *BaseemailaddressListener) EnterLwsp(ctx *LwspContext) {}

// ExitLwsp is called when production lwsp is exited.
func (s *BaseemailaddressListener) ExitLwsp(ctx *LwspContext) {}

// EnterDelimeters is called when production delimeters is entered.
func (s *BaseemailaddressListener) EnterDelimeters(ctx *DelimetersContext) {}

// ExitDelimeters is called when production delimeters is exited.
func (s *BaseemailaddressListener) ExitDelimeters(ctx *DelimetersContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseemailaddressListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseemailaddressListener) ExitAtom(ctx *AtomContext) {}

// EnterQuotedpair is called when production quotedpair is entered.
func (s *BaseemailaddressListener) EnterQuotedpair(ctx *QuotedpairContext) {}

// ExitQuotedpair is called when production quotedpair is exited.
func (s *BaseemailaddressListener) ExitQuotedpair(ctx *QuotedpairContext) {}

// EnterDomainliteral is called when production domainliteral is entered.
func (s *BaseemailaddressListener) EnterDomainliteral(ctx *DomainliteralContext) {}

// ExitDomainliteral is called when production domainliteral is exited.
func (s *BaseemailaddressListener) ExitDomainliteral(ctx *DomainliteralContext) {}

// EnterQuotedstring is called when production quotedstring is entered.
func (s *BaseemailaddressListener) EnterQuotedstring(ctx *QuotedstringContext) {}

// ExitQuotedstring is called when production quotedstring is exited.
func (s *BaseemailaddressListener) ExitQuotedstring(ctx *QuotedstringContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseemailaddressListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseemailaddressListener) ExitComment(ctx *CommentContext) {}
