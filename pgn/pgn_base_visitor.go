// Generated from PGN.g4 by ANTLR 4.7.

package pgn // PGN
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePGNVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePGNVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitPgn_database(ctx *Pgn_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitPgn_game(ctx *Pgn_gameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitTag_section(ctx *Tag_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitTag_pair(ctx *Tag_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitTag_name(ctx *Tag_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitTag_value(ctx *Tag_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitMovetext_section(ctx *Movetext_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitElement_sequence(ctx *Element_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitMove_number_indication(ctx *Move_number_indicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitSan_move(ctx *San_moveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitRecursive_variation(ctx *Recursive_variationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePGNVisitor) VisitGame_termination(ctx *Game_terminationContext) interface{} {
	return v.VisitChildren(ctx)
}
