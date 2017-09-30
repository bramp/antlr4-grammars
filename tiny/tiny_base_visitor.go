// Generated from tiny.g4 by ANTLR 4.7.

package tiny // tiny
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetinyVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetinyVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitStmt_list(ctx *Stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitRead_stmt(ctx *Read_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitWrite_stmt(ctx *Write_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitId_list(ctx *Id_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitExpr_list(ctx *Expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinyVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}
