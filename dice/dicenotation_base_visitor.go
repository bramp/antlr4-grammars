// Generated from DiceNotation.g4 by ANTLR 4.7.

package dice // DiceNotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDiceNotationVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDiceNotationVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceNotationVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceNotationVisitor) VisitBinaryOp(ctx *BinaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDiceNotationVisitor) VisitDice(ctx *DiceContext) interface{} {
	return v.VisitChildren(ctx)
}
