// Code generated from cookie.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cookie // cookie
import "github.com/antlr/antlr4/runtime/Go/antlr"

// cookieListener is a complete listener for a parse tree produced by cookieParser.
type cookieListener interface {
	antlr.ParseTreeListener

	// EnterCookie is called when entering the cookie production.
	EnterCookie(c *CookieContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterAv_pairs is called when entering the av_pairs production.
	EnterAv_pairs(c *Av_pairsContext)

	// EnterAv_pair is called when entering the av_pair production.
	EnterAv_pair(c *Av_pairContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterQuoted_string is called when entering the quoted_string production.
	EnterQuoted_string(c *Quoted_stringContext)

	// ExitCookie is called when exiting the cookie production.
	ExitCookie(c *CookieContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitAv_pairs is called when exiting the av_pairs production.
	ExitAv_pairs(c *Av_pairsContext)

	// ExitAv_pair is called when exiting the av_pair production.
	ExitAv_pair(c *Av_pairContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitQuoted_string is called when exiting the quoted_string production.
	ExitQuoted_string(c *Quoted_stringContext)
}
