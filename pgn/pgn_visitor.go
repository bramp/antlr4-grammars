// Generated from PGN.g4 by ANTLR 4.7.

package pgn // PGN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PGNParser.
type PGNVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PGNParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by PGNParser#pgn_database.
	VisitPgn_database(ctx *Pgn_databaseContext) interface{}

	// Visit a parse tree produced by PGNParser#pgn_game.
	VisitPgn_game(ctx *Pgn_gameContext) interface{}

	// Visit a parse tree produced by PGNParser#tag_section.
	VisitTag_section(ctx *Tag_sectionContext) interface{}

	// Visit a parse tree produced by PGNParser#tag_pair.
	VisitTag_pair(ctx *Tag_pairContext) interface{}

	// Visit a parse tree produced by PGNParser#tag_name.
	VisitTag_name(ctx *Tag_nameContext) interface{}

	// Visit a parse tree produced by PGNParser#tag_value.
	VisitTag_value(ctx *Tag_valueContext) interface{}

	// Visit a parse tree produced by PGNParser#movetext_section.
	VisitMovetext_section(ctx *Movetext_sectionContext) interface{}

	// Visit a parse tree produced by PGNParser#element_sequence.
	VisitElement_sequence(ctx *Element_sequenceContext) interface{}

	// Visit a parse tree produced by PGNParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by PGNParser#move_number_indication.
	VisitMove_number_indication(ctx *Move_number_indicationContext) interface{}

	// Visit a parse tree produced by PGNParser#san_move.
	VisitSan_move(ctx *San_moveContext) interface{}

	// Visit a parse tree produced by PGNParser#recursive_variation.
	VisitRecursive_variation(ctx *Recursive_variationContext) interface{}

	// Visit a parse tree produced by PGNParser#game_termination.
	VisitGame_termination(ctx *Game_terminationContext) interface{}
}
