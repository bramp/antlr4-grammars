// Generated from tinyc.g4 by ANTLR 4.7.

package tinyc // tinyc
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetinycVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetinycVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitParen_expr(ctx *Paren_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitSum(ctx *SumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}
