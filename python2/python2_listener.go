// Generated from Python2.g4 by ANTLR 4.7.

package python2 // Python2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Python2Listener is a complete listener for a parse tree produced by Python2Parser.
type Python2Listener interface {
	antlr.ParseTreeListener

	// EnterSingle_input is called when entering the single_input production.
	EnterSingle_input(c *Single_inputContext)

	// EnterFile_input is called when entering the file_input production.
	EnterFile_input(c *File_inputContext)

	// EnterEval_input is called when entering the eval_input production.
	EnterEval_input(c *Eval_inputContext)

	// EnterDecorator is called when entering the decorator production.
	EnterDecorator(c *DecoratorContext)

	// EnterDecorators is called when entering the decorators production.
	EnterDecorators(c *DecoratorsContext)

	// EnterDecorated is called when entering the decorated production.
	EnterDecorated(c *DecoratedContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterVarargslist is called when entering the varargslist production.
	EnterVarargslist(c *VarargslistContext)

	// EnterFpdef is called when entering the fpdef production.
	EnterFpdef(c *FpdefContext)

	// EnterFplist is called when entering the fplist production.
	EnterFplist(c *FplistContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterSmall_stmt is called when entering the small_stmt production.
	EnterSmall_stmt(c *Small_stmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterAugassign is called when entering the augassign production.
	EnterAugassign(c *AugassignContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterDel_stmt is called when entering the del_stmt production.
	EnterDel_stmt(c *Del_stmtContext)

	// EnterPass_stmt is called when entering the pass_stmt production.
	EnterPass_stmt(c *Pass_stmtContext)

	// EnterFlow_stmt is called when entering the flow_stmt production.
	EnterFlow_stmt(c *Flow_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterYield_stmt is called when entering the yield_stmt production.
	EnterYield_stmt(c *Yield_stmtContext)

	// EnterRaise_stmt is called when entering the raise_stmt production.
	EnterRaise_stmt(c *Raise_stmtContext)

	// EnterImport_stmt is called when entering the import_stmt production.
	EnterImport_stmt(c *Import_stmtContext)

	// EnterImport_name is called when entering the import_name production.
	EnterImport_name(c *Import_nameContext)

	// EnterImport_from is called when entering the import_from production.
	EnterImport_from(c *Import_fromContext)

	// EnterImport_as_name is called when entering the import_as_name production.
	EnterImport_as_name(c *Import_as_nameContext)

	// EnterDotted_as_name is called when entering the dotted_as_name production.
	EnterDotted_as_name(c *Dotted_as_nameContext)

	// EnterImport_as_names is called when entering the import_as_names production.
	EnterImport_as_names(c *Import_as_namesContext)

	// EnterDotted_as_names is called when entering the dotted_as_names production.
	EnterDotted_as_names(c *Dotted_as_namesContext)

	// EnterDotted_name is called when entering the dotted_name production.
	EnterDotted_name(c *Dotted_nameContext)

	// EnterGlobal_stmt is called when entering the global_stmt production.
	EnterGlobal_stmt(c *Global_stmtContext)

	// EnterExec_stmt is called when entering the exec_stmt production.
	EnterExec_stmt(c *Exec_stmtContext)

	// EnterAssert_stmt is called when entering the assert_stmt production.
	EnterAssert_stmt(c *Assert_stmtContext)

	// EnterCompound_stmt is called when entering the compound_stmt production.
	EnterCompound_stmt(c *Compound_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterTry_stmt is called when entering the try_stmt production.
	EnterTry_stmt(c *Try_stmtContext)

	// EnterWith_stmt is called when entering the with_stmt production.
	EnterWith_stmt(c *With_stmtContext)

	// EnterWith_item is called when entering the with_item production.
	EnterWith_item(c *With_itemContext)

	// EnterExcept_clause is called when entering the except_clause production.
	EnterExcept_clause(c *Except_clauseContext)

	// EnterSuite is called when entering the suite production.
	EnterSuite(c *SuiteContext)

	// EnterTestlist_safe is called when entering the testlist_safe production.
	EnterTestlist_safe(c *Testlist_safeContext)

	// EnterOld_test is called when entering the old_test production.
	EnterOld_test(c *Old_testContext)

	// EnterOld_lambdef is called when entering the old_lambdef production.
	EnterOld_lambdef(c *Old_lambdefContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterOr_test is called when entering the or_test production.
	EnterOr_test(c *Or_testContext)

	// EnterAnd_test is called when entering the and_test production.
	EnterAnd_test(c *And_testContext)

	// EnterNot_test is called when entering the not_test production.
	EnterNot_test(c *Not_testContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterXor_expr is called when entering the xor_expr production.
	EnterXor_expr(c *Xor_exprContext)

	// EnterAnd_expr is called when entering the and_expr production.
	EnterAnd_expr(c *And_exprContext)

	// EnterShift_expr is called when entering the shift_expr production.
	EnterShift_expr(c *Shift_exprContext)

	// EnterArith_expr is called when entering the arith_expr production.
	EnterArith_expr(c *Arith_exprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterListmaker is called when entering the listmaker production.
	EnterListmaker(c *ListmakerContext)

	// EnterTestlist_comp is called when entering the testlist_comp production.
	EnterTestlist_comp(c *Testlist_compContext)

	// EnterLambdef is called when entering the lambdef production.
	EnterLambdef(c *LambdefContext)

	// EnterTrailer is called when entering the trailer production.
	EnterTrailer(c *TrailerContext)

	// EnterSubscriptlist is called when entering the subscriptlist production.
	EnterSubscriptlist(c *SubscriptlistContext)

	// EnterSubscript is called when entering the subscript production.
	EnterSubscript(c *SubscriptContext)

	// EnterSliceop is called when entering the sliceop production.
	EnterSliceop(c *SliceopContext)

	// EnterExprlist is called when entering the exprlist production.
	EnterExprlist(c *ExprlistContext)

	// EnterTestlist is called when entering the testlist production.
	EnterTestlist(c *TestlistContext)

	// EnterDictorsetmaker is called when entering the dictorsetmaker production.
	EnterDictorsetmaker(c *DictorsetmakerContext)

	// EnterClassdef is called when entering the classdef production.
	EnterClassdef(c *ClassdefContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterList_iter is called when entering the list_iter production.
	EnterList_iter(c *List_iterContext)

	// EnterList_for is called when entering the list_for production.
	EnterList_for(c *List_forContext)

	// EnterList_if is called when entering the list_if production.
	EnterList_if(c *List_ifContext)

	// EnterComp_iter is called when entering the comp_iter production.
	EnterComp_iter(c *Comp_iterContext)

	// EnterComp_for is called when entering the comp_for production.
	EnterComp_for(c *Comp_forContext)

	// EnterComp_if is called when entering the comp_if production.
	EnterComp_if(c *Comp_ifContext)

	// EnterTestlist1 is called when entering the testlist1 production.
	EnterTestlist1(c *Testlist1Context)

	// EnterEncoding_decl is called when entering the encoding_decl production.
	EnterEncoding_decl(c *Encoding_declContext)

	// EnterYield_expr is called when entering the yield_expr production.
	EnterYield_expr(c *Yield_exprContext)

	// ExitSingle_input is called when exiting the single_input production.
	ExitSingle_input(c *Single_inputContext)

	// ExitFile_input is called when exiting the file_input production.
	ExitFile_input(c *File_inputContext)

	// ExitEval_input is called when exiting the eval_input production.
	ExitEval_input(c *Eval_inputContext)

	// ExitDecorator is called when exiting the decorator production.
	ExitDecorator(c *DecoratorContext)

	// ExitDecorators is called when exiting the decorators production.
	ExitDecorators(c *DecoratorsContext)

	// ExitDecorated is called when exiting the decorated production.
	ExitDecorated(c *DecoratedContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitVarargslist is called when exiting the varargslist production.
	ExitVarargslist(c *VarargslistContext)

	// ExitFpdef is called when exiting the fpdef production.
	ExitFpdef(c *FpdefContext)

	// ExitFplist is called when exiting the fplist production.
	ExitFplist(c *FplistContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitSmall_stmt is called when exiting the small_stmt production.
	ExitSmall_stmt(c *Small_stmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitAugassign is called when exiting the augassign production.
	ExitAugassign(c *AugassignContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitDel_stmt is called when exiting the del_stmt production.
	ExitDel_stmt(c *Del_stmtContext)

	// ExitPass_stmt is called when exiting the pass_stmt production.
	ExitPass_stmt(c *Pass_stmtContext)

	// ExitFlow_stmt is called when exiting the flow_stmt production.
	ExitFlow_stmt(c *Flow_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitYield_stmt is called when exiting the yield_stmt production.
	ExitYield_stmt(c *Yield_stmtContext)

	// ExitRaise_stmt is called when exiting the raise_stmt production.
	ExitRaise_stmt(c *Raise_stmtContext)

	// ExitImport_stmt is called when exiting the import_stmt production.
	ExitImport_stmt(c *Import_stmtContext)

	// ExitImport_name is called when exiting the import_name production.
	ExitImport_name(c *Import_nameContext)

	// ExitImport_from is called when exiting the import_from production.
	ExitImport_from(c *Import_fromContext)

	// ExitImport_as_name is called when exiting the import_as_name production.
	ExitImport_as_name(c *Import_as_nameContext)

	// ExitDotted_as_name is called when exiting the dotted_as_name production.
	ExitDotted_as_name(c *Dotted_as_nameContext)

	// ExitImport_as_names is called when exiting the import_as_names production.
	ExitImport_as_names(c *Import_as_namesContext)

	// ExitDotted_as_names is called when exiting the dotted_as_names production.
	ExitDotted_as_names(c *Dotted_as_namesContext)

	// ExitDotted_name is called when exiting the dotted_name production.
	ExitDotted_name(c *Dotted_nameContext)

	// ExitGlobal_stmt is called when exiting the global_stmt production.
	ExitGlobal_stmt(c *Global_stmtContext)

	// ExitExec_stmt is called when exiting the exec_stmt production.
	ExitExec_stmt(c *Exec_stmtContext)

	// ExitAssert_stmt is called when exiting the assert_stmt production.
	ExitAssert_stmt(c *Assert_stmtContext)

	// ExitCompound_stmt is called when exiting the compound_stmt production.
	ExitCompound_stmt(c *Compound_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitTry_stmt is called when exiting the try_stmt production.
	ExitTry_stmt(c *Try_stmtContext)

	// ExitWith_stmt is called when exiting the with_stmt production.
	ExitWith_stmt(c *With_stmtContext)

	// ExitWith_item is called when exiting the with_item production.
	ExitWith_item(c *With_itemContext)

	// ExitExcept_clause is called when exiting the except_clause production.
	ExitExcept_clause(c *Except_clauseContext)

	// ExitSuite is called when exiting the suite production.
	ExitSuite(c *SuiteContext)

	// ExitTestlist_safe is called when exiting the testlist_safe production.
	ExitTestlist_safe(c *Testlist_safeContext)

	// ExitOld_test is called when exiting the old_test production.
	ExitOld_test(c *Old_testContext)

	// ExitOld_lambdef is called when exiting the old_lambdef production.
	ExitOld_lambdef(c *Old_lambdefContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitOr_test is called when exiting the or_test production.
	ExitOr_test(c *Or_testContext)

	// ExitAnd_test is called when exiting the and_test production.
	ExitAnd_test(c *And_testContext)

	// ExitNot_test is called when exiting the not_test production.
	ExitNot_test(c *Not_testContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitXor_expr is called when exiting the xor_expr production.
	ExitXor_expr(c *Xor_exprContext)

	// ExitAnd_expr is called when exiting the and_expr production.
	ExitAnd_expr(c *And_exprContext)

	// ExitShift_expr is called when exiting the shift_expr production.
	ExitShift_expr(c *Shift_exprContext)

	// ExitArith_expr is called when exiting the arith_expr production.
	ExitArith_expr(c *Arith_exprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitListmaker is called when exiting the listmaker production.
	ExitListmaker(c *ListmakerContext)

	// ExitTestlist_comp is called when exiting the testlist_comp production.
	ExitTestlist_comp(c *Testlist_compContext)

	// ExitLambdef is called when exiting the lambdef production.
	ExitLambdef(c *LambdefContext)

	// ExitTrailer is called when exiting the trailer production.
	ExitTrailer(c *TrailerContext)

	// ExitSubscriptlist is called when exiting the subscriptlist production.
	ExitSubscriptlist(c *SubscriptlistContext)

	// ExitSubscript is called when exiting the subscript production.
	ExitSubscript(c *SubscriptContext)

	// ExitSliceop is called when exiting the sliceop production.
	ExitSliceop(c *SliceopContext)

	// ExitExprlist is called when exiting the exprlist production.
	ExitExprlist(c *ExprlistContext)

	// ExitTestlist is called when exiting the testlist production.
	ExitTestlist(c *TestlistContext)

	// ExitDictorsetmaker is called when exiting the dictorsetmaker production.
	ExitDictorsetmaker(c *DictorsetmakerContext)

	// ExitClassdef is called when exiting the classdef production.
	ExitClassdef(c *ClassdefContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitList_iter is called when exiting the list_iter production.
	ExitList_iter(c *List_iterContext)

	// ExitList_for is called when exiting the list_for production.
	ExitList_for(c *List_forContext)

	// ExitList_if is called when exiting the list_if production.
	ExitList_if(c *List_ifContext)

	// ExitComp_iter is called when exiting the comp_iter production.
	ExitComp_iter(c *Comp_iterContext)

	// ExitComp_for is called when exiting the comp_for production.
	ExitComp_for(c *Comp_forContext)

	// ExitComp_if is called when exiting the comp_if production.
	ExitComp_if(c *Comp_ifContext)

	// ExitTestlist1 is called when exiting the testlist1 production.
	ExitTestlist1(c *Testlist1Context)

	// ExitEncoding_decl is called when exiting the encoding_decl production.
	ExitEncoding_decl(c *Encoding_declContext)

	// ExitYield_expr is called when exiting the yield_expr production.
	ExitYield_expr(c *Yield_exprContext)
}
