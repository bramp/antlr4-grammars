// Generated from Python2.g4 by ANTLR 4.7.

package python2 // Python2
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePython2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePython2Visitor) VisitSingle_input(ctx *Single_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFile_input(ctx *File_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitEval_input(ctx *Eval_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDecorator(ctx *DecoratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDecorators(ctx *DecoratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDecorated(ctx *DecoratedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFuncdef(ctx *FuncdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitVarargslist(ctx *VarargslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFpdef(ctx *FpdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFplist(ctx *FplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSimple_stmt(ctx *Simple_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSmall_stmt(ctx *Small_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitExpr_stmt(ctx *Expr_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitAugassign(ctx *AugassignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitPrint_stmt(ctx *Print_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDel_stmt(ctx *Del_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitPass_stmt(ctx *Pass_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFlow_stmt(ctx *Flow_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitBreak_stmt(ctx *Break_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitContinue_stmt(ctx *Continue_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitYield_stmt(ctx *Yield_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitRaise_stmt(ctx *Raise_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitImport_stmt(ctx *Import_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitImport_name(ctx *Import_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitImport_from(ctx *Import_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitImport_as_name(ctx *Import_as_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDotted_as_name(ctx *Dotted_as_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitImport_as_names(ctx *Import_as_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDotted_as_names(ctx *Dotted_as_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDotted_name(ctx *Dotted_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitGlobal_stmt(ctx *Global_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitExec_stmt(ctx *Exec_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitAssert_stmt(ctx *Assert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitCompound_stmt(ctx *Compound_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitWhile_stmt(ctx *While_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTry_stmt(ctx *Try_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitWith_stmt(ctx *With_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitWith_item(ctx *With_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitExcept_clause(ctx *Except_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSuite(ctx *SuiteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTestlist_safe(ctx *Testlist_safeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitOld_test(ctx *Old_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitOld_lambdef(ctx *Old_lambdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitOr_test(ctx *Or_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitAnd_test(ctx *And_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitNot_test(ctx *Not_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitComp_op(ctx *Comp_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitXor_expr(ctx *Xor_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitAnd_expr(ctx *And_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitShift_expr(ctx *Shift_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitArith_expr(ctx *Arith_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitListmaker(ctx *ListmakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTestlist_comp(ctx *Testlist_compContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitLambdef(ctx *LambdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTrailer(ctx *TrailerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSubscriptlist(ctx *SubscriptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSubscript(ctx *SubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitSliceop(ctx *SliceopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitExprlist(ctx *ExprlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTestlist(ctx *TestlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitDictorsetmaker(ctx *DictorsetmakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitClassdef(ctx *ClassdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitArglist(ctx *ArglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitList_iter(ctx *List_iterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitList_for(ctx *List_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitList_if(ctx *List_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitComp_iter(ctx *Comp_iterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitComp_for(ctx *Comp_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitComp_if(ctx *Comp_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitTestlist1(ctx *Testlist1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitEncoding_decl(ctx *Encoding_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePython2Visitor) VisitYield_expr(ctx *Yield_exprContext) interface{} {
	return v.VisitChildren(ctx)
}
