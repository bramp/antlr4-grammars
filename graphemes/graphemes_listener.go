// Code generated from Graphemes.g4 by ANTLR 4.9.3. DO NOT EDIT.

package graphemes // Graphemes
import "github.com/antlr/antlr4/runtime/Go/antlr"

// GraphemesListener is a complete listener for a parse tree produced by GraphemesParser.
type GraphemesListener interface {
	antlr.ParseTreeListener

	// EnterEmoji_sequence is called when entering the emoji_sequence production.
	EnterEmoji_sequence(c *Emoji_sequenceContext)

	// EnterGrapheme_cluster is called when entering the grapheme_cluster production.
	EnterGrapheme_cluster(c *Grapheme_clusterContext)

	// EnterGraphemes is called when entering the graphemes production.
	EnterGraphemes(c *GraphemesContext)

	// ExitEmoji_sequence is called when exiting the emoji_sequence production.
	ExitEmoji_sequence(c *Emoji_sequenceContext)

	// ExitGrapheme_cluster is called when exiting the grapheme_cluster production.
	ExitGrapheme_cluster(c *Grapheme_clusterContext)

	// ExitGraphemes is called when exiting the graphemes production.
	ExitGraphemes(c *GraphemesContext)
}
