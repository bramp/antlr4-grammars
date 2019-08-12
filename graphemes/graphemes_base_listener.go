// Code generated from Graphemes.g4 by ANTLR 4.7.2. DO NOT EDIT.

package graphemes // Graphemes
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraphemesListener is a complete listener for a parse tree produced by GraphemesParser.
type BaseGraphemesListener struct{}

var _ GraphemesListener = &BaseGraphemesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphemesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphemesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphemesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphemesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEmoji_sequence is called when production emoji_sequence is entered.
func (s *BaseGraphemesListener) EnterEmoji_sequence(ctx *Emoji_sequenceContext) {}

// ExitEmoji_sequence is called when production emoji_sequence is exited.
func (s *BaseGraphemesListener) ExitEmoji_sequence(ctx *Emoji_sequenceContext) {}

// EnterGrapheme_cluster is called when production grapheme_cluster is entered.
func (s *BaseGraphemesListener) EnterGrapheme_cluster(ctx *Grapheme_clusterContext) {}

// ExitGrapheme_cluster is called when production grapheme_cluster is exited.
func (s *BaseGraphemesListener) ExitGrapheme_cluster(ctx *Grapheme_clusterContext) {}

// EnterGraphemes is called when production graphemes is entered.
func (s *BaseGraphemesListener) EnterGraphemes(ctx *GraphemesContext) {}

// ExitGraphemes is called when production graphemes is exited.
func (s *BaseGraphemesListener) ExitGraphemes(ctx *GraphemesContext) {}
