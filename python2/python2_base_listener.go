// Generated from Python2.g4 by ANTLR 4.7.

package python2 // Python2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePython2Listener is a complete listener for a parse tree produced by Python2Parser.
type BasePython2Listener struct{}

var _ Python2Listener = &BasePython2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePython2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePython2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePython2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePython2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingle_input is called when production single_input is entered.
func (s *BasePython2Listener) EnterSingle_input(ctx *Single_inputContext) {}

// ExitSingle_input is called when production single_input is exited.
func (s *BasePython2Listener) ExitSingle_input(ctx *Single_inputContext) {}

// EnterFile_input is called when production file_input is entered.
func (s *BasePython2Listener) EnterFile_input(ctx *File_inputContext) {}

// ExitFile_input is called when production file_input is exited.
func (s *BasePython2Listener) ExitFile_input(ctx *File_inputContext) {}

// EnterEval_input is called when production eval_input is entered.
func (s *BasePython2Listener) EnterEval_input(ctx *Eval_inputContext) {}

// ExitEval_input is called when production eval_input is exited.
func (s *BasePython2Listener) ExitEval_input(ctx *Eval_inputContext) {}

// EnterDecorator is called when production decorator is entered.
func (s *BasePython2Listener) EnterDecorator(ctx *DecoratorContext) {}

// ExitDecorator is called when production decorator is exited.
func (s *BasePython2Listener) ExitDecorator(ctx *DecoratorContext) {}

// EnterDecorators is called when production decorators is entered.
func (s *BasePython2Listener) EnterDecorators(ctx *DecoratorsContext) {}

// ExitDecorators is called when production decorators is exited.
func (s *BasePython2Listener) ExitDecorators(ctx *DecoratorsContext) {}

// EnterDecorated is called when production decorated is entered.
func (s *BasePython2Listener) EnterDecorated(ctx *DecoratedContext) {}

// ExitDecorated is called when production decorated is exited.
func (s *BasePython2Listener) ExitDecorated(ctx *DecoratedContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BasePython2Listener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BasePython2Listener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BasePython2Listener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BasePython2Listener) ExitParameters(ctx *ParametersContext) {}

// EnterVarargslist is called when production varargslist is entered.
func (s *BasePython2Listener) EnterVarargslist(ctx *VarargslistContext) {}

// ExitVarargslist is called when production varargslist is exited.
func (s *BasePython2Listener) ExitVarargslist(ctx *VarargslistContext) {}

// EnterFpdef is called when production fpdef is entered.
func (s *BasePython2Listener) EnterFpdef(ctx *FpdefContext) {}

// ExitFpdef is called when production fpdef is exited.
func (s *BasePython2Listener) ExitFpdef(ctx *FpdefContext) {}

// EnterFplist is called when production fplist is entered.
func (s *BasePython2Listener) EnterFplist(ctx *FplistContext) {}

// ExitFplist is called when production fplist is exited.
func (s *BasePython2Listener) ExitFplist(ctx *FplistContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePython2Listener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePython2Listener) ExitStmt(ctx *StmtContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BasePython2Listener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BasePython2Listener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterSmall_stmt is called when production small_stmt is entered.
func (s *BasePython2Listener) EnterSmall_stmt(ctx *Small_stmtContext) {}

// ExitSmall_stmt is called when production small_stmt is exited.
func (s *BasePython2Listener) ExitSmall_stmt(ctx *Small_stmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BasePython2Listener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BasePython2Listener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterAugassign is called when production augassign is entered.
func (s *BasePython2Listener) EnterAugassign(ctx *AugassignContext) {}

// ExitAugassign is called when production augassign is exited.
func (s *BasePython2Listener) ExitAugassign(ctx *AugassignContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BasePython2Listener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BasePython2Listener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterDel_stmt is called when production del_stmt is entered.
func (s *BasePython2Listener) EnterDel_stmt(ctx *Del_stmtContext) {}

// ExitDel_stmt is called when production del_stmt is exited.
func (s *BasePython2Listener) ExitDel_stmt(ctx *Del_stmtContext) {}

// EnterPass_stmt is called when production pass_stmt is entered.
func (s *BasePython2Listener) EnterPass_stmt(ctx *Pass_stmtContext) {}

// ExitPass_stmt is called when production pass_stmt is exited.
func (s *BasePython2Listener) ExitPass_stmt(ctx *Pass_stmtContext) {}

// EnterFlow_stmt is called when production flow_stmt is entered.
func (s *BasePython2Listener) EnterFlow_stmt(ctx *Flow_stmtContext) {}

// ExitFlow_stmt is called when production flow_stmt is exited.
func (s *BasePython2Listener) ExitFlow_stmt(ctx *Flow_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BasePython2Listener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BasePython2Listener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BasePython2Listener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BasePython2Listener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BasePython2Listener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BasePython2Listener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterYield_stmt is called when production yield_stmt is entered.
func (s *BasePython2Listener) EnterYield_stmt(ctx *Yield_stmtContext) {}

// ExitYield_stmt is called when production yield_stmt is exited.
func (s *BasePython2Listener) ExitYield_stmt(ctx *Yield_stmtContext) {}

// EnterRaise_stmt is called when production raise_stmt is entered.
func (s *BasePython2Listener) EnterRaise_stmt(ctx *Raise_stmtContext) {}

// ExitRaise_stmt is called when production raise_stmt is exited.
func (s *BasePython2Listener) ExitRaise_stmt(ctx *Raise_stmtContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BasePython2Listener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BasePython2Listener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterImport_name is called when production import_name is entered.
func (s *BasePython2Listener) EnterImport_name(ctx *Import_nameContext) {}

// ExitImport_name is called when production import_name is exited.
func (s *BasePython2Listener) ExitImport_name(ctx *Import_nameContext) {}

// EnterImport_from is called when production import_from is entered.
func (s *BasePython2Listener) EnterImport_from(ctx *Import_fromContext) {}

// ExitImport_from is called when production import_from is exited.
func (s *BasePython2Listener) ExitImport_from(ctx *Import_fromContext) {}

// EnterImport_as_name is called when production import_as_name is entered.
func (s *BasePython2Listener) EnterImport_as_name(ctx *Import_as_nameContext) {}

// ExitImport_as_name is called when production import_as_name is exited.
func (s *BasePython2Listener) ExitImport_as_name(ctx *Import_as_nameContext) {}

// EnterDotted_as_name is called when production dotted_as_name is entered.
func (s *BasePython2Listener) EnterDotted_as_name(ctx *Dotted_as_nameContext) {}

// ExitDotted_as_name is called when production dotted_as_name is exited.
func (s *BasePython2Listener) ExitDotted_as_name(ctx *Dotted_as_nameContext) {}

// EnterImport_as_names is called when production import_as_names is entered.
func (s *BasePython2Listener) EnterImport_as_names(ctx *Import_as_namesContext) {}

// ExitImport_as_names is called when production import_as_names is exited.
func (s *BasePython2Listener) ExitImport_as_names(ctx *Import_as_namesContext) {}

// EnterDotted_as_names is called when production dotted_as_names is entered.
func (s *BasePython2Listener) EnterDotted_as_names(ctx *Dotted_as_namesContext) {}

// ExitDotted_as_names is called when production dotted_as_names is exited.
func (s *BasePython2Listener) ExitDotted_as_names(ctx *Dotted_as_namesContext) {}

// EnterDotted_name is called when production dotted_name is entered.
func (s *BasePython2Listener) EnterDotted_name(ctx *Dotted_nameContext) {}

// ExitDotted_name is called when production dotted_name is exited.
func (s *BasePython2Listener) ExitDotted_name(ctx *Dotted_nameContext) {}

// EnterGlobal_stmt is called when production global_stmt is entered.
func (s *BasePython2Listener) EnterGlobal_stmt(ctx *Global_stmtContext) {}

// ExitGlobal_stmt is called when production global_stmt is exited.
func (s *BasePython2Listener) ExitGlobal_stmt(ctx *Global_stmtContext) {}

// EnterExec_stmt is called when production exec_stmt is entered.
func (s *BasePython2Listener) EnterExec_stmt(ctx *Exec_stmtContext) {}

// ExitExec_stmt is called when production exec_stmt is exited.
func (s *BasePython2Listener) ExitExec_stmt(ctx *Exec_stmtContext) {}

// EnterAssert_stmt is called when production assert_stmt is entered.
func (s *BasePython2Listener) EnterAssert_stmt(ctx *Assert_stmtContext) {}

// ExitAssert_stmt is called when production assert_stmt is exited.
func (s *BasePython2Listener) ExitAssert_stmt(ctx *Assert_stmtContext) {}

// EnterCompound_stmt is called when production compound_stmt is entered.
func (s *BasePython2Listener) EnterCompound_stmt(ctx *Compound_stmtContext) {}

// ExitCompound_stmt is called when production compound_stmt is exited.
func (s *BasePython2Listener) ExitCompound_stmt(ctx *Compound_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BasePython2Listener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BasePython2Listener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BasePython2Listener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BasePython2Listener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BasePython2Listener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BasePython2Listener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterTry_stmt is called when production try_stmt is entered.
func (s *BasePython2Listener) EnterTry_stmt(ctx *Try_stmtContext) {}

// ExitTry_stmt is called when production try_stmt is exited.
func (s *BasePython2Listener) ExitTry_stmt(ctx *Try_stmtContext) {}

// EnterWith_stmt is called when production with_stmt is entered.
func (s *BasePython2Listener) EnterWith_stmt(ctx *With_stmtContext) {}

// ExitWith_stmt is called when production with_stmt is exited.
func (s *BasePython2Listener) ExitWith_stmt(ctx *With_stmtContext) {}

// EnterWith_item is called when production with_item is entered.
func (s *BasePython2Listener) EnterWith_item(ctx *With_itemContext) {}

// ExitWith_item is called when production with_item is exited.
func (s *BasePython2Listener) ExitWith_item(ctx *With_itemContext) {}

// EnterExcept_clause is called when production except_clause is entered.
func (s *BasePython2Listener) EnterExcept_clause(ctx *Except_clauseContext) {}

// ExitExcept_clause is called when production except_clause is exited.
func (s *BasePython2Listener) ExitExcept_clause(ctx *Except_clauseContext) {}

// EnterSuite is called when production suite is entered.
func (s *BasePython2Listener) EnterSuite(ctx *SuiteContext) {}

// ExitSuite is called when production suite is exited.
func (s *BasePython2Listener) ExitSuite(ctx *SuiteContext) {}

// EnterTestlist_safe is called when production testlist_safe is entered.
func (s *BasePython2Listener) EnterTestlist_safe(ctx *Testlist_safeContext) {}

// ExitTestlist_safe is called when production testlist_safe is exited.
func (s *BasePython2Listener) ExitTestlist_safe(ctx *Testlist_safeContext) {}

// EnterOld_test is called when production old_test is entered.
func (s *BasePython2Listener) EnterOld_test(ctx *Old_testContext) {}

// ExitOld_test is called when production old_test is exited.
func (s *BasePython2Listener) ExitOld_test(ctx *Old_testContext) {}

// EnterOld_lambdef is called when production old_lambdef is entered.
func (s *BasePython2Listener) EnterOld_lambdef(ctx *Old_lambdefContext) {}

// ExitOld_lambdef is called when production old_lambdef is exited.
func (s *BasePython2Listener) ExitOld_lambdef(ctx *Old_lambdefContext) {}

// EnterTest is called when production test is entered.
func (s *BasePython2Listener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasePython2Listener) ExitTest(ctx *TestContext) {}

// EnterOr_test is called when production or_test is entered.
func (s *BasePython2Listener) EnterOr_test(ctx *Or_testContext) {}

// ExitOr_test is called when production or_test is exited.
func (s *BasePython2Listener) ExitOr_test(ctx *Or_testContext) {}

// EnterAnd_test is called when production and_test is entered.
func (s *BasePython2Listener) EnterAnd_test(ctx *And_testContext) {}

// ExitAnd_test is called when production and_test is exited.
func (s *BasePython2Listener) ExitAnd_test(ctx *And_testContext) {}

// EnterNot_test is called when production not_test is entered.
func (s *BasePython2Listener) EnterNot_test(ctx *Not_testContext) {}

// ExitNot_test is called when production not_test is exited.
func (s *BasePython2Listener) ExitNot_test(ctx *Not_testContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasePython2Listener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasePython2Listener) ExitComparison(ctx *ComparisonContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BasePython2Listener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BasePython2Listener) ExitComp_op(ctx *Comp_opContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePython2Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePython2Listener) ExitExpr(ctx *ExprContext) {}

// EnterXor_expr is called when production xor_expr is entered.
func (s *BasePython2Listener) EnterXor_expr(ctx *Xor_exprContext) {}

// ExitXor_expr is called when production xor_expr is exited.
func (s *BasePython2Listener) ExitXor_expr(ctx *Xor_exprContext) {}

// EnterAnd_expr is called when production and_expr is entered.
func (s *BasePython2Listener) EnterAnd_expr(ctx *And_exprContext) {}

// ExitAnd_expr is called when production and_expr is exited.
func (s *BasePython2Listener) ExitAnd_expr(ctx *And_exprContext) {}

// EnterShift_expr is called when production shift_expr is entered.
func (s *BasePython2Listener) EnterShift_expr(ctx *Shift_exprContext) {}

// ExitShift_expr is called when production shift_expr is exited.
func (s *BasePython2Listener) ExitShift_expr(ctx *Shift_exprContext) {}

// EnterArith_expr is called when production arith_expr is entered.
func (s *BasePython2Listener) EnterArith_expr(ctx *Arith_exprContext) {}

// ExitArith_expr is called when production arith_expr is exited.
func (s *BasePython2Listener) ExitArith_expr(ctx *Arith_exprContext) {}

// EnterTerm is called when production term is entered.
func (s *BasePython2Listener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasePython2Listener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasePython2Listener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasePython2Listener) ExitFactor(ctx *FactorContext) {}

// EnterPower is called when production power is entered.
func (s *BasePython2Listener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BasePython2Listener) ExitPower(ctx *PowerContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePython2Listener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePython2Listener) ExitAtom(ctx *AtomContext) {}

// EnterListmaker is called when production listmaker is entered.
func (s *BasePython2Listener) EnterListmaker(ctx *ListmakerContext) {}

// ExitListmaker is called when production listmaker is exited.
func (s *BasePython2Listener) ExitListmaker(ctx *ListmakerContext) {}

// EnterTestlist_comp is called when production testlist_comp is entered.
func (s *BasePython2Listener) EnterTestlist_comp(ctx *Testlist_compContext) {}

// ExitTestlist_comp is called when production testlist_comp is exited.
func (s *BasePython2Listener) ExitTestlist_comp(ctx *Testlist_compContext) {}

// EnterLambdef is called when production lambdef is entered.
func (s *BasePython2Listener) EnterLambdef(ctx *LambdefContext) {}

// ExitLambdef is called when production lambdef is exited.
func (s *BasePython2Listener) ExitLambdef(ctx *LambdefContext) {}

// EnterTrailer is called when production trailer is entered.
func (s *BasePython2Listener) EnterTrailer(ctx *TrailerContext) {}

// ExitTrailer is called when production trailer is exited.
func (s *BasePython2Listener) ExitTrailer(ctx *TrailerContext) {}

// EnterSubscriptlist is called when production subscriptlist is entered.
func (s *BasePython2Listener) EnterSubscriptlist(ctx *SubscriptlistContext) {}

// ExitSubscriptlist is called when production subscriptlist is exited.
func (s *BasePython2Listener) ExitSubscriptlist(ctx *SubscriptlistContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BasePython2Listener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BasePython2Listener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSliceop is called when production sliceop is entered.
func (s *BasePython2Listener) EnterSliceop(ctx *SliceopContext) {}

// ExitSliceop is called when production sliceop is exited.
func (s *BasePython2Listener) ExitSliceop(ctx *SliceopContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BasePython2Listener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BasePython2Listener) ExitExprlist(ctx *ExprlistContext) {}

// EnterTestlist is called when production testlist is entered.
func (s *BasePython2Listener) EnterTestlist(ctx *TestlistContext) {}

// ExitTestlist is called when production testlist is exited.
func (s *BasePython2Listener) ExitTestlist(ctx *TestlistContext) {}

// EnterDictorsetmaker is called when production dictorsetmaker is entered.
func (s *BasePython2Listener) EnterDictorsetmaker(ctx *DictorsetmakerContext) {}

// ExitDictorsetmaker is called when production dictorsetmaker is exited.
func (s *BasePython2Listener) ExitDictorsetmaker(ctx *DictorsetmakerContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BasePython2Listener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BasePython2Listener) ExitClassdef(ctx *ClassdefContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BasePython2Listener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BasePython2Listener) ExitArglist(ctx *ArglistContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasePython2Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasePython2Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterList_iter is called when production list_iter is entered.
func (s *BasePython2Listener) EnterList_iter(ctx *List_iterContext) {}

// ExitList_iter is called when production list_iter is exited.
func (s *BasePython2Listener) ExitList_iter(ctx *List_iterContext) {}

// EnterList_for is called when production list_for is entered.
func (s *BasePython2Listener) EnterList_for(ctx *List_forContext) {}

// ExitList_for is called when production list_for is exited.
func (s *BasePython2Listener) ExitList_for(ctx *List_forContext) {}

// EnterList_if is called when production list_if is entered.
func (s *BasePython2Listener) EnterList_if(ctx *List_ifContext) {}

// ExitList_if is called when production list_if is exited.
func (s *BasePython2Listener) ExitList_if(ctx *List_ifContext) {}

// EnterComp_iter is called when production comp_iter is entered.
func (s *BasePython2Listener) EnterComp_iter(ctx *Comp_iterContext) {}

// ExitComp_iter is called when production comp_iter is exited.
func (s *BasePython2Listener) ExitComp_iter(ctx *Comp_iterContext) {}

// EnterComp_for is called when production comp_for is entered.
func (s *BasePython2Listener) EnterComp_for(ctx *Comp_forContext) {}

// ExitComp_for is called when production comp_for is exited.
func (s *BasePython2Listener) ExitComp_for(ctx *Comp_forContext) {}

// EnterComp_if is called when production comp_if is entered.
func (s *BasePython2Listener) EnterComp_if(ctx *Comp_ifContext) {}

// ExitComp_if is called when production comp_if is exited.
func (s *BasePython2Listener) ExitComp_if(ctx *Comp_ifContext) {}

// EnterTestlist1 is called when production testlist1 is entered.
func (s *BasePython2Listener) EnterTestlist1(ctx *Testlist1Context) {}

// ExitTestlist1 is called when production testlist1 is exited.
func (s *BasePython2Listener) ExitTestlist1(ctx *Testlist1Context) {}

// EnterEncoding_decl is called when production encoding_decl is entered.
func (s *BasePython2Listener) EnterEncoding_decl(ctx *Encoding_declContext) {}

// ExitEncoding_decl is called when production encoding_decl is exited.
func (s *BasePython2Listener) ExitEncoding_decl(ctx *Encoding_declContext) {}

// EnterYield_expr is called when production yield_expr is entered.
func (s *BasePython2Listener) EnterYield_expr(ctx *Yield_exprContext) {}

// ExitYield_expr is called when production yield_expr is exited.
func (s *BasePython2Listener) ExitYield_expr(ctx *Yield_exprContext) {}
