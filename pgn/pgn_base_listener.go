// Generated from PGN.g4 by ANTLR 4.7.

package pgn // PGN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePGNListener is a complete listener for a parse tree produced by PGNParser.
type BasePGNListener struct{}

var _ PGNListener = &BasePGNListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePGNListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePGNListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePGNListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePGNListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BasePGNListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BasePGNListener) ExitParse(ctx *ParseContext) {}

// EnterPgn_database is called when production pgn_database is entered.
func (s *BasePGNListener) EnterPgn_database(ctx *Pgn_databaseContext) {}

// ExitPgn_database is called when production pgn_database is exited.
func (s *BasePGNListener) ExitPgn_database(ctx *Pgn_databaseContext) {}

// EnterPgn_game is called when production pgn_game is entered.
func (s *BasePGNListener) EnterPgn_game(ctx *Pgn_gameContext) {}

// ExitPgn_game is called when production pgn_game is exited.
func (s *BasePGNListener) ExitPgn_game(ctx *Pgn_gameContext) {}

// EnterTag_section is called when production tag_section is entered.
func (s *BasePGNListener) EnterTag_section(ctx *Tag_sectionContext) {}

// ExitTag_section is called when production tag_section is exited.
func (s *BasePGNListener) ExitTag_section(ctx *Tag_sectionContext) {}

// EnterTag_pair is called when production tag_pair is entered.
func (s *BasePGNListener) EnterTag_pair(ctx *Tag_pairContext) {}

// ExitTag_pair is called when production tag_pair is exited.
func (s *BasePGNListener) ExitTag_pair(ctx *Tag_pairContext) {}

// EnterTag_name is called when production tag_name is entered.
func (s *BasePGNListener) EnterTag_name(ctx *Tag_nameContext) {}

// ExitTag_name is called when production tag_name is exited.
func (s *BasePGNListener) ExitTag_name(ctx *Tag_nameContext) {}

// EnterTag_value is called when production tag_value is entered.
func (s *BasePGNListener) EnterTag_value(ctx *Tag_valueContext) {}

// ExitTag_value is called when production tag_value is exited.
func (s *BasePGNListener) ExitTag_value(ctx *Tag_valueContext) {}

// EnterMovetext_section is called when production movetext_section is entered.
func (s *BasePGNListener) EnterMovetext_section(ctx *Movetext_sectionContext) {}

// ExitMovetext_section is called when production movetext_section is exited.
func (s *BasePGNListener) ExitMovetext_section(ctx *Movetext_sectionContext) {}

// EnterElement_sequence is called when production element_sequence is entered.
func (s *BasePGNListener) EnterElement_sequence(ctx *Element_sequenceContext) {}

// ExitElement_sequence is called when production element_sequence is exited.
func (s *BasePGNListener) ExitElement_sequence(ctx *Element_sequenceContext) {}

// EnterElement is called when production element is entered.
func (s *BasePGNListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasePGNListener) ExitElement(ctx *ElementContext) {}

// EnterMove_number_indication is called when production move_number_indication is entered.
func (s *BasePGNListener) EnterMove_number_indication(ctx *Move_number_indicationContext) {}

// ExitMove_number_indication is called when production move_number_indication is exited.
func (s *BasePGNListener) ExitMove_number_indication(ctx *Move_number_indicationContext) {}

// EnterSan_move is called when production san_move is entered.
func (s *BasePGNListener) EnterSan_move(ctx *San_moveContext) {}

// ExitSan_move is called when production san_move is exited.
func (s *BasePGNListener) ExitSan_move(ctx *San_moveContext) {}

// EnterRecursive_variation is called when production recursive_variation is entered.
func (s *BasePGNListener) EnterRecursive_variation(ctx *Recursive_variationContext) {}

// ExitRecursive_variation is called when production recursive_variation is exited.
func (s *BasePGNListener) ExitRecursive_variation(ctx *Recursive_variationContext) {}

// EnterGame_termination is called when production game_termination is entered.
func (s *BasePGNListener) EnterGame_termination(ctx *Game_terminationContext) {}

// ExitGame_termination is called when production game_termination is exited.
func (s *BasePGNListener) ExitGame_termination(ctx *Game_terminationContext) {}
