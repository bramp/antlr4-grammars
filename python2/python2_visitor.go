// Generated from Python2.g4 by ANTLR 4.7.

package python2 // Python2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Python2Parser.
type Python2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Python2Parser#single_input.
	VisitSingle_input(ctx *Single_inputContext) interface{}

	// Visit a parse tree produced by Python2Parser#file_input.
	VisitFile_input(ctx *File_inputContext) interface{}

	// Visit a parse tree produced by Python2Parser#eval_input.
	VisitEval_input(ctx *Eval_inputContext) interface{}

	// Visit a parse tree produced by Python2Parser#decorator.
	VisitDecorator(ctx *DecoratorContext) interface{}

	// Visit a parse tree produced by Python2Parser#decorators.
	VisitDecorators(ctx *DecoratorsContext) interface{}

	// Visit a parse tree produced by Python2Parser#decorated.
	VisitDecorated(ctx *DecoratedContext) interface{}

	// Visit a parse tree produced by Python2Parser#funcdef.
	VisitFuncdef(ctx *FuncdefContext) interface{}

	// Visit a parse tree produced by Python2Parser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by Python2Parser#varargslist.
	VisitVarargslist(ctx *VarargslistContext) interface{}

	// Visit a parse tree produced by Python2Parser#fpdef.
	VisitFpdef(ctx *FpdefContext) interface{}

	// Visit a parse tree produced by Python2Parser#fplist.
	VisitFplist(ctx *FplistContext) interface{}

	// Visit a parse tree produced by Python2Parser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#simple_stmt.
	VisitSimple_stmt(ctx *Simple_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#small_stmt.
	VisitSmall_stmt(ctx *Small_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#expr_stmt.
	VisitExpr_stmt(ctx *Expr_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#augassign.
	VisitAugassign(ctx *AugassignContext) interface{}

	// Visit a parse tree produced by Python2Parser#print_stmt.
	VisitPrint_stmt(ctx *Print_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#del_stmt.
	VisitDel_stmt(ctx *Del_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#pass_stmt.
	VisitPass_stmt(ctx *Pass_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#flow_stmt.
	VisitFlow_stmt(ctx *Flow_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#break_stmt.
	VisitBreak_stmt(ctx *Break_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#continue_stmt.
	VisitContinue_stmt(ctx *Continue_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#yield_stmt.
	VisitYield_stmt(ctx *Yield_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#raise_stmt.
	VisitRaise_stmt(ctx *Raise_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#import_stmt.
	VisitImport_stmt(ctx *Import_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#import_name.
	VisitImport_name(ctx *Import_nameContext) interface{}

	// Visit a parse tree produced by Python2Parser#import_from.
	VisitImport_from(ctx *Import_fromContext) interface{}

	// Visit a parse tree produced by Python2Parser#import_as_name.
	VisitImport_as_name(ctx *Import_as_nameContext) interface{}

	// Visit a parse tree produced by Python2Parser#dotted_as_name.
	VisitDotted_as_name(ctx *Dotted_as_nameContext) interface{}

	// Visit a parse tree produced by Python2Parser#import_as_names.
	VisitImport_as_names(ctx *Import_as_namesContext) interface{}

	// Visit a parse tree produced by Python2Parser#dotted_as_names.
	VisitDotted_as_names(ctx *Dotted_as_namesContext) interface{}

	// Visit a parse tree produced by Python2Parser#dotted_name.
	VisitDotted_name(ctx *Dotted_nameContext) interface{}

	// Visit a parse tree produced by Python2Parser#global_stmt.
	VisitGlobal_stmt(ctx *Global_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#exec_stmt.
	VisitExec_stmt(ctx *Exec_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#assert_stmt.
	VisitAssert_stmt(ctx *Assert_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#compound_stmt.
	VisitCompound_stmt(ctx *Compound_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#while_stmt.
	VisitWhile_stmt(ctx *While_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#try_stmt.
	VisitTry_stmt(ctx *Try_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#with_stmt.
	VisitWith_stmt(ctx *With_stmtContext) interface{}

	// Visit a parse tree produced by Python2Parser#with_item.
	VisitWith_item(ctx *With_itemContext) interface{}

	// Visit a parse tree produced by Python2Parser#except_clause.
	VisitExcept_clause(ctx *Except_clauseContext) interface{}

	// Visit a parse tree produced by Python2Parser#suite.
	VisitSuite(ctx *SuiteContext) interface{}

	// Visit a parse tree produced by Python2Parser#testlist_safe.
	VisitTestlist_safe(ctx *Testlist_safeContext) interface{}

	// Visit a parse tree produced by Python2Parser#old_test.
	VisitOld_test(ctx *Old_testContext) interface{}

	// Visit a parse tree produced by Python2Parser#old_lambdef.
	VisitOld_lambdef(ctx *Old_lambdefContext) interface{}

	// Visit a parse tree produced by Python2Parser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by Python2Parser#or_test.
	VisitOr_test(ctx *Or_testContext) interface{}

	// Visit a parse tree produced by Python2Parser#and_test.
	VisitAnd_test(ctx *And_testContext) interface{}

	// Visit a parse tree produced by Python2Parser#not_test.
	VisitNot_test(ctx *Not_testContext) interface{}

	// Visit a parse tree produced by Python2Parser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by Python2Parser#comp_op.
	VisitComp_op(ctx *Comp_opContext) interface{}

	// Visit a parse tree produced by Python2Parser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by Python2Parser#xor_expr.
	VisitXor_expr(ctx *Xor_exprContext) interface{}

	// Visit a parse tree produced by Python2Parser#and_expr.
	VisitAnd_expr(ctx *And_exprContext) interface{}

	// Visit a parse tree produced by Python2Parser#shift_expr.
	VisitShift_expr(ctx *Shift_exprContext) interface{}

	// Visit a parse tree produced by Python2Parser#arith_expr.
	VisitArith_expr(ctx *Arith_exprContext) interface{}

	// Visit a parse tree produced by Python2Parser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by Python2Parser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by Python2Parser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by Python2Parser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by Python2Parser#listmaker.
	VisitListmaker(ctx *ListmakerContext) interface{}

	// Visit a parse tree produced by Python2Parser#testlist_comp.
	VisitTestlist_comp(ctx *Testlist_compContext) interface{}

	// Visit a parse tree produced by Python2Parser#lambdef.
	VisitLambdef(ctx *LambdefContext) interface{}

	// Visit a parse tree produced by Python2Parser#trailer.
	VisitTrailer(ctx *TrailerContext) interface{}

	// Visit a parse tree produced by Python2Parser#subscriptlist.
	VisitSubscriptlist(ctx *SubscriptlistContext) interface{}

	// Visit a parse tree produced by Python2Parser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by Python2Parser#sliceop.
	VisitSliceop(ctx *SliceopContext) interface{}

	// Visit a parse tree produced by Python2Parser#exprlist.
	VisitExprlist(ctx *ExprlistContext) interface{}

	// Visit a parse tree produced by Python2Parser#testlist.
	VisitTestlist(ctx *TestlistContext) interface{}

	// Visit a parse tree produced by Python2Parser#dictorsetmaker.
	VisitDictorsetmaker(ctx *DictorsetmakerContext) interface{}

	// Visit a parse tree produced by Python2Parser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by Python2Parser#arglist.
	VisitArglist(ctx *ArglistContext) interface{}

	// Visit a parse tree produced by Python2Parser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by Python2Parser#list_iter.
	VisitList_iter(ctx *List_iterContext) interface{}

	// Visit a parse tree produced by Python2Parser#list_for.
	VisitList_for(ctx *List_forContext) interface{}

	// Visit a parse tree produced by Python2Parser#list_if.
	VisitList_if(ctx *List_ifContext) interface{}

	// Visit a parse tree produced by Python2Parser#comp_iter.
	VisitComp_iter(ctx *Comp_iterContext) interface{}

	// Visit a parse tree produced by Python2Parser#comp_for.
	VisitComp_for(ctx *Comp_forContext) interface{}

	// Visit a parse tree produced by Python2Parser#comp_if.
	VisitComp_if(ctx *Comp_ifContext) interface{}

	// Visit a parse tree produced by Python2Parser#testlist1.
	VisitTestlist1(ctx *Testlist1Context) interface{}

	// Visit a parse tree produced by Python2Parser#encoding_decl.
	VisitEncoding_decl(ctx *Encoding_declContext) interface{}

	// Visit a parse tree produced by Python2Parser#yield_expr.
	VisitYield_expr(ctx *Yield_exprContext) interface{}
}
