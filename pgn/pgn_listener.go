// Generated from PGN.g4 by ANTLR 4.7.

package pgn // PGN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PGNListener is a complete listener for a parse tree produced by PGNParser.
type PGNListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterPgn_database is called when entering the pgn_database production.
	EnterPgn_database(c *Pgn_databaseContext)

	// EnterPgn_game is called when entering the pgn_game production.
	EnterPgn_game(c *Pgn_gameContext)

	// EnterTag_section is called when entering the tag_section production.
	EnterTag_section(c *Tag_sectionContext)

	// EnterTag_pair is called when entering the tag_pair production.
	EnterTag_pair(c *Tag_pairContext)

	// EnterTag_name is called when entering the tag_name production.
	EnterTag_name(c *Tag_nameContext)

	// EnterTag_value is called when entering the tag_value production.
	EnterTag_value(c *Tag_valueContext)

	// EnterMovetext_section is called when entering the movetext_section production.
	EnterMovetext_section(c *Movetext_sectionContext)

	// EnterElement_sequence is called when entering the element_sequence production.
	EnterElement_sequence(c *Element_sequenceContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterMove_number_indication is called when entering the move_number_indication production.
	EnterMove_number_indication(c *Move_number_indicationContext)

	// EnterSan_move is called when entering the san_move production.
	EnterSan_move(c *San_moveContext)

	// EnterRecursive_variation is called when entering the recursive_variation production.
	EnterRecursive_variation(c *Recursive_variationContext)

	// EnterGame_termination is called when entering the game_termination production.
	EnterGame_termination(c *Game_terminationContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitPgn_database is called when exiting the pgn_database production.
	ExitPgn_database(c *Pgn_databaseContext)

	// ExitPgn_game is called when exiting the pgn_game production.
	ExitPgn_game(c *Pgn_gameContext)

	// ExitTag_section is called when exiting the tag_section production.
	ExitTag_section(c *Tag_sectionContext)

	// ExitTag_pair is called when exiting the tag_pair production.
	ExitTag_pair(c *Tag_pairContext)

	// ExitTag_name is called when exiting the tag_name production.
	ExitTag_name(c *Tag_nameContext)

	// ExitTag_value is called when exiting the tag_value production.
	ExitTag_value(c *Tag_valueContext)

	// ExitMovetext_section is called when exiting the movetext_section production.
	ExitMovetext_section(c *Movetext_sectionContext)

	// ExitElement_sequence is called when exiting the element_sequence production.
	ExitElement_sequence(c *Element_sequenceContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitMove_number_indication is called when exiting the move_number_indication production.
	ExitMove_number_indication(c *Move_number_indicationContext)

	// ExitSan_move is called when exiting the san_move production.
	ExitSan_move(c *San_moveContext)

	// ExitRecursive_variation is called when exiting the recursive_variation production.
	ExitRecursive_variation(c *Recursive_variationContext)

	// ExitGame_termination is called when exiting the game_termination production.
	ExitGame_termination(c *Game_terminationContext)
}
