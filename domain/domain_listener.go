// Code generated from domain.g4 by ANTLR 4.9.3. DO NOT EDIT.

package domain // domain
import "github.com/antlr/antlr4/runtime/Go/antlr"

// domainListener is a complete listener for a parse tree produced by domainParser.
type domainListener interface {
	antlr.ParseTreeListener

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterSubdomain is called when entering the subdomain production.
	EnterSubdomain(c *SubdomainContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitSubdomain is called when exiting the subdomain production.
	ExitSubdomain(c *SubdomainContext)
}
