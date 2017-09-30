// Generated from useragent.g4 by ANTLR 4.7.

package useragent // useragent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// useragentListener is a complete listener for a parse tree produced by useragentParser.
type useragentListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterProduct is called when entering the product production.
	EnterProduct(c *ProductContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitProduct is called when exiting the product production.
	ExitProduct(c *ProductContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
