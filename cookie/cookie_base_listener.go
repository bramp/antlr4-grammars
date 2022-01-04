// Code generated from cookie.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cookie // cookie
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecookieListener is a complete listener for a parse tree produced by cookieParser.
type BasecookieListener struct{}

var _ cookieListener = &BasecookieListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecookieListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecookieListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecookieListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecookieListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCookie is called when production cookie is entered.
func (s *BasecookieListener) EnterCookie(ctx *CookieContext) {}

// ExitCookie is called when production cookie is exited.
func (s *BasecookieListener) ExitCookie(ctx *CookieContext) {}

// EnterName is called when production name is entered.
func (s *BasecookieListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasecookieListener) ExitName(ctx *NameContext) {}

// EnterAv_pairs is called when production av_pairs is entered.
func (s *BasecookieListener) EnterAv_pairs(ctx *Av_pairsContext) {}

// ExitAv_pairs is called when production av_pairs is exited.
func (s *BasecookieListener) ExitAv_pairs(ctx *Av_pairsContext) {}

// EnterAv_pair is called when production av_pair is entered.
func (s *BasecookieListener) EnterAv_pair(ctx *Av_pairContext) {}

// ExitAv_pair is called when production av_pair is exited.
func (s *BasecookieListener) ExitAv_pair(ctx *Av_pairContext) {}

// EnterAttr is called when production attr is entered.
func (s *BasecookieListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BasecookieListener) ExitAttr(ctx *AttrContext) {}

// EnterValue is called when production value is entered.
func (s *BasecookieListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasecookieListener) ExitValue(ctx *ValueContext) {}

// EnterWord is called when production word is entered.
func (s *BasecookieListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BasecookieListener) ExitWord(ctx *WordContext) {}

// EnterToken is called when production token is entered.
func (s *BasecookieListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BasecookieListener) ExitToken(ctx *TokenContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BasecookieListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BasecookieListener) ExitQuoted_string(ctx *Quoted_stringContext) {}
