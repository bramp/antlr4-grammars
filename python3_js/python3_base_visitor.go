// Generated from Python3.g4 by ANTLR 4.7.

package python3_js // Python3
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePython3Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePython3Visitor) VisitSingle_input(ctx *Single_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitFile_input(ctx *File_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitEval_input(ctx *Eval_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDecorator(ctx *DecoratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDecorators(ctx *DecoratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDecorated(ctx *DecoratedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitFuncdef(ctx *FuncdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTypedargslist(ctx *TypedargslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTfpdef(ctx *TfpdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitVarargslist(ctx *VarargslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitVfpdef(ctx *VfpdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSimple_stmt(ctx *Simple_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSmall_stmt(ctx *Small_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitExpr_stmt(ctx *Expr_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTestlist_star_expr(ctx *Testlist_star_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitAugassign(ctx *AugassignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDel_stmt(ctx *Del_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitPass_stmt(ctx *Pass_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitFlow_stmt(ctx *Flow_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitBreak_stmt(ctx *Break_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitContinue_stmt(ctx *Continue_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitYield_stmt(ctx *Yield_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitRaise_stmt(ctx *Raise_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitImport_stmt(ctx *Import_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitImport_name(ctx *Import_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitImport_from(ctx *Import_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitImport_as_name(ctx *Import_as_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDotted_as_name(ctx *Dotted_as_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitImport_as_names(ctx *Import_as_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDotted_as_names(ctx *Dotted_as_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDotted_name(ctx *Dotted_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitGlobal_stmt(ctx *Global_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitNonlocal_stmt(ctx *Nonlocal_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitAssert_stmt(ctx *Assert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitCompound_stmt(ctx *Compound_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitWhile_stmt(ctx *While_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTry_stmt(ctx *Try_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitWith_stmt(ctx *With_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitWith_item(ctx *With_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitExcept_clause(ctx *Except_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSuite(ctx *SuiteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTest_nocond(ctx *Test_nocondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitLambdef(ctx *LambdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitLambdef_nocond(ctx *Lambdef_nocondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitOr_test(ctx *Or_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitAnd_test(ctx *And_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitNot_test(ctx *Not_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitComp_op(ctx *Comp_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitStar_expr(ctx *Star_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitXor_expr(ctx *Xor_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitAnd_expr(ctx *And_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitShift_expr(ctx *Shift_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitArith_expr(ctx *Arith_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTestlist_comp(ctx *Testlist_compContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTrailer(ctx *TrailerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSubscriptlist(ctx *SubscriptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitSliceop(ctx *SliceopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitTestlist(ctx *TestlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitDictorsetmaker(ctx *DictorsetmakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitClassdef(ctx *ClassdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitArglist(ctx *ArglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitComp_iter(ctx *Comp_iterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitComp_for(ctx *Comp_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitComp_if(ctx *Comp_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitYield_expr(ctx *Yield_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitYield_arg(ctx *Yield_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython3Visitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}
