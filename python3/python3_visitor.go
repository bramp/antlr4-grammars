// Generated from Python3.g4 by ANTLR 4.7.

package python3 // Python3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Python3Parser.
type Python3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Python3Parser#single_input.
	VisitSingle_input(ctx *Single_inputContext) interface{}

	// Visit a parse tree produced by Python3Parser#file_input.
	VisitFile_input(ctx *File_inputContext) interface{}

	// Visit a parse tree produced by Python3Parser#eval_input.
	VisitEval_input(ctx *Eval_inputContext) interface{}

	// Visit a parse tree produced by Python3Parser#decorator.
	VisitDecorator(ctx *DecoratorContext) interface{}

	// Visit a parse tree produced by Python3Parser#decorators.
	VisitDecorators(ctx *DecoratorsContext) interface{}

	// Visit a parse tree produced by Python3Parser#decorated.
	VisitDecorated(ctx *DecoratedContext) interface{}

	// Visit a parse tree produced by Python3Parser#async_funcdef.
	VisitAsync_funcdef(ctx *Async_funcdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#funcdef.
	VisitFuncdef(ctx *FuncdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by Python3Parser#typedargslist.
	VisitTypedargslist(ctx *TypedargslistContext) interface{}

	// Visit a parse tree produced by Python3Parser#tfpdef.
	VisitTfpdef(ctx *TfpdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#varargslist.
	VisitVarargslist(ctx *VarargslistContext) interface{}

	// Visit a parse tree produced by Python3Parser#vfpdef.
	VisitVfpdef(ctx *VfpdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#simple_stmt.
	VisitSimple_stmt(ctx *Simple_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#small_stmt.
	VisitSmall_stmt(ctx *Small_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#expr_stmt.
	VisitExpr_stmt(ctx *Expr_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#annassign.
	VisitAnnassign(ctx *AnnassignContext) interface{}

	// Visit a parse tree produced by Python3Parser#testlist_star_expr.
	VisitTestlist_star_expr(ctx *Testlist_star_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#augassign.
	VisitAugassign(ctx *AugassignContext) interface{}

	// Visit a parse tree produced by Python3Parser#del_stmt.
	VisitDel_stmt(ctx *Del_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#pass_stmt.
	VisitPass_stmt(ctx *Pass_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#flow_stmt.
	VisitFlow_stmt(ctx *Flow_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#break_stmt.
	VisitBreak_stmt(ctx *Break_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#continue_stmt.
	VisitContinue_stmt(ctx *Continue_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#yield_stmt.
	VisitYield_stmt(ctx *Yield_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#raise_stmt.
	VisitRaise_stmt(ctx *Raise_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#import_stmt.
	VisitImport_stmt(ctx *Import_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#import_name.
	VisitImport_name(ctx *Import_nameContext) interface{}

	// Visit a parse tree produced by Python3Parser#import_from.
	VisitImport_from(ctx *Import_fromContext) interface{}

	// Visit a parse tree produced by Python3Parser#import_as_name.
	VisitImport_as_name(ctx *Import_as_nameContext) interface{}

	// Visit a parse tree produced by Python3Parser#dotted_as_name.
	VisitDotted_as_name(ctx *Dotted_as_nameContext) interface{}

	// Visit a parse tree produced by Python3Parser#import_as_names.
	VisitImport_as_names(ctx *Import_as_namesContext) interface{}

	// Visit a parse tree produced by Python3Parser#dotted_as_names.
	VisitDotted_as_names(ctx *Dotted_as_namesContext) interface{}

	// Visit a parse tree produced by Python3Parser#dotted_name.
	VisitDotted_name(ctx *Dotted_nameContext) interface{}

	// Visit a parse tree produced by Python3Parser#global_stmt.
	VisitGlobal_stmt(ctx *Global_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#nonlocal_stmt.
	VisitNonlocal_stmt(ctx *Nonlocal_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#assert_stmt.
	VisitAssert_stmt(ctx *Assert_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#compound_stmt.
	VisitCompound_stmt(ctx *Compound_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#async_stmt.
	VisitAsync_stmt(ctx *Async_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#while_stmt.
	VisitWhile_stmt(ctx *While_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#try_stmt.
	VisitTry_stmt(ctx *Try_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#with_stmt.
	VisitWith_stmt(ctx *With_stmtContext) interface{}

	// Visit a parse tree produced by Python3Parser#with_item.
	VisitWith_item(ctx *With_itemContext) interface{}

	// Visit a parse tree produced by Python3Parser#except_clause.
	VisitExcept_clause(ctx *Except_clauseContext) interface{}

	// Visit a parse tree produced by Python3Parser#suite.
	VisitSuite(ctx *SuiteContext) interface{}

	// Visit a parse tree produced by Python3Parser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by Python3Parser#test_nocond.
	VisitTest_nocond(ctx *Test_nocondContext) interface{}

	// Visit a parse tree produced by Python3Parser#lambdef.
	VisitLambdef(ctx *LambdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#lambdef_nocond.
	VisitLambdef_nocond(ctx *Lambdef_nocondContext) interface{}

	// Visit a parse tree produced by Python3Parser#or_test.
	VisitOr_test(ctx *Or_testContext) interface{}

	// Visit a parse tree produced by Python3Parser#and_test.
	VisitAnd_test(ctx *And_testContext) interface{}

	// Visit a parse tree produced by Python3Parser#not_test.
	VisitNot_test(ctx *Not_testContext) interface{}

	// Visit a parse tree produced by Python3Parser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by Python3Parser#comp_op.
	VisitComp_op(ctx *Comp_opContext) interface{}

	// Visit a parse tree produced by Python3Parser#star_expr.
	VisitStar_expr(ctx *Star_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by Python3Parser#xor_expr.
	VisitXor_expr(ctx *Xor_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#and_expr.
	VisitAnd_expr(ctx *And_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#shift_expr.
	VisitShift_expr(ctx *Shift_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#arith_expr.
	VisitArith_expr(ctx *Arith_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by Python3Parser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by Python3Parser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by Python3Parser#atom_expr.
	VisitAtom_expr(ctx *Atom_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by Python3Parser#testlist_comp.
	VisitTestlist_comp(ctx *Testlist_compContext) interface{}

	// Visit a parse tree produced by Python3Parser#trailer.
	VisitTrailer(ctx *TrailerContext) interface{}

	// Visit a parse tree produced by Python3Parser#subscriptlist.
	VisitSubscriptlist(ctx *SubscriptlistContext) interface{}

	// Visit a parse tree produced by Python3Parser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by Python3Parser#sliceop.
	VisitSliceop(ctx *SliceopContext) interface{}

	// Visit a parse tree produced by Python3Parser#exprlist.
	VisitExprlist(ctx *ExprlistContext) interface{}

	// Visit a parse tree produced by Python3Parser#testlist.
	VisitTestlist(ctx *TestlistContext) interface{}

	// Visit a parse tree produced by Python3Parser#dictorsetmaker.
	VisitDictorsetmaker(ctx *DictorsetmakerContext) interface{}

	// Visit a parse tree produced by Python3Parser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by Python3Parser#arglist.
	VisitArglist(ctx *ArglistContext) interface{}

	// Visit a parse tree produced by Python3Parser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by Python3Parser#comp_iter.
	VisitComp_iter(ctx *Comp_iterContext) interface{}

	// Visit a parse tree produced by Python3Parser#comp_for.
	VisitComp_for(ctx *Comp_forContext) interface{}

	// Visit a parse tree produced by Python3Parser#comp_if.
	VisitComp_if(ctx *Comp_ifContext) interface{}

	// Visit a parse tree produced by Python3Parser#encoding_decl.
	VisitEncoding_decl(ctx *Encoding_declContext) interface{}

	// Visit a parse tree produced by Python3Parser#yield_expr.
	VisitYield_expr(ctx *Yield_exprContext) interface{}

	// Visit a parse tree produced by Python3Parser#yield_arg.
	VisitYield_arg(ctx *Yield_argContext) interface{}
}
