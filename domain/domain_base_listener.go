// Code generated from domain.g4 by ANTLR 4.9.3. DO NOT EDIT.

package domain // domain
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedomainListener is a complete listener for a parse tree produced by domainParser.
type BasedomainListener struct{}

var _ domainListener = &BasedomainListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedomainListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedomainListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedomainListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedomainListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDomain is called when production domain is entered.
func (s *BasedomainListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BasedomainListener) ExitDomain(ctx *DomainContext) {}

// EnterSubdomain is called when production subdomain is entered.
func (s *BasedomainListener) EnterSubdomain(ctx *SubdomainContext) {}

// ExitSubdomain is called when production subdomain is exited.
func (s *BasedomainListener) ExitSubdomain(ctx *SubdomainContext) {}
