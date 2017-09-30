// Generated from romannumerals.g4 by ANTLR 4.7.

package romannumerals // romannumerals
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseromannumeralsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseromannumeralsVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitThousands(ctx *ThousandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitThous_part(ctx *Thous_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitHundreds(ctx *HundredsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitHun_part(ctx *Hun_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitHun_rep(ctx *Hun_repContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitTens(ctx *TensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitTens_part(ctx *Tens_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitTens_rep(ctx *Tens_repContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitOnes(ctx *OnesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseromannumeralsVisitor) VisitOnes_rep(ctx *Ones_repContext) interface{} {
	return v.VisitChildren(ctx)
}
