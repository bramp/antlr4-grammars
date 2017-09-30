// Generated from Python3.g4 by ANTLR 4.7.

package python3-js // Python3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePython3Listener is a complete listener for a parse tree produced by Python3Parser.
type BasePython3Listener struct{}

var _ Python3Listener = &BasePython3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePython3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePython3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePython3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePython3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingle_input is called when production single_input is entered.
func (s *BasePython3Listener) EnterSingle_input(ctx *Single_inputContext) {}

// ExitSingle_input is called when production single_input is exited.
func (s *BasePython3Listener) ExitSingle_input(ctx *Single_inputContext) {}

// EnterFile_input is called when production file_input is entered.
func (s *BasePython3Listener) EnterFile_input(ctx *File_inputContext) {}

// ExitFile_input is called when production file_input is exited.
func (s *BasePython3Listener) ExitFile_input(ctx *File_inputContext) {}

// EnterEval_input is called when production eval_input is entered.
func (s *BasePython3Listener) EnterEval_input(ctx *Eval_inputContext) {}

// ExitEval_input is called when production eval_input is exited.
func (s *BasePython3Listener) ExitEval_input(ctx *Eval_inputContext) {}

// EnterDecorator is called when production decorator is entered.
func (s *BasePython3Listener) EnterDecorator(ctx *DecoratorContext) {}

// ExitDecorator is called when production decorator is exited.
func (s *BasePython3Listener) ExitDecorator(ctx *DecoratorContext) {}

// EnterDecorators is called when production decorators is entered.
func (s *BasePython3Listener) EnterDecorators(ctx *DecoratorsContext) {}

// ExitDecorators is called when production decorators is exited.
func (s *BasePython3Listener) ExitDecorators(ctx *DecoratorsContext) {}

// EnterDecorated is called when production decorated is entered.
func (s *BasePython3Listener) EnterDecorated(ctx *DecoratedContext) {}

// ExitDecorated is called when production decorated is exited.
func (s *BasePython3Listener) ExitDecorated(ctx *DecoratedContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BasePython3Listener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BasePython3Listener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BasePython3Listener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BasePython3Listener) ExitParameters(ctx *ParametersContext) {}

// EnterTypedargslist is called when production typedargslist is entered.
func (s *BasePython3Listener) EnterTypedargslist(ctx *TypedargslistContext) {}

// ExitTypedargslist is called when production typedargslist is exited.
func (s *BasePython3Listener) ExitTypedargslist(ctx *TypedargslistContext) {}

// EnterTfpdef is called when production tfpdef is entered.
func (s *BasePython3Listener) EnterTfpdef(ctx *TfpdefContext) {}

// ExitTfpdef is called when production tfpdef is exited.
func (s *BasePython3Listener) ExitTfpdef(ctx *TfpdefContext) {}

// EnterVarargslist is called when production varargslist is entered.
func (s *BasePython3Listener) EnterVarargslist(ctx *VarargslistContext) {}

// ExitVarargslist is called when production varargslist is exited.
func (s *BasePython3Listener) ExitVarargslist(ctx *VarargslistContext) {}

// EnterVfpdef is called when production vfpdef is entered.
func (s *BasePython3Listener) EnterVfpdef(ctx *VfpdefContext) {}

// ExitVfpdef is called when production vfpdef is exited.
func (s *BasePython3Listener) ExitVfpdef(ctx *VfpdefContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePython3Listener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePython3Listener) ExitStmt(ctx *StmtContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BasePython3Listener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BasePython3Listener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterSmall_stmt is called when production small_stmt is entered.
func (s *BasePython3Listener) EnterSmall_stmt(ctx *Small_stmtContext) {}

// ExitSmall_stmt is called when production small_stmt is exited.
func (s *BasePython3Listener) ExitSmall_stmt(ctx *Small_stmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BasePython3Listener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BasePython3Listener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterTestlist_star_expr is called when production testlist_star_expr is entered.
func (s *BasePython3Listener) EnterTestlist_star_expr(ctx *Testlist_star_exprContext) {}

// ExitTestlist_star_expr is called when production testlist_star_expr is exited.
func (s *BasePython3Listener) ExitTestlist_star_expr(ctx *Testlist_star_exprContext) {}

// EnterAugassign is called when production augassign is entered.
func (s *BasePython3Listener) EnterAugassign(ctx *AugassignContext) {}

// ExitAugassign is called when production augassign is exited.
func (s *BasePython3Listener) ExitAugassign(ctx *AugassignContext) {}

// EnterDel_stmt is called when production del_stmt is entered.
func (s *BasePython3Listener) EnterDel_stmt(ctx *Del_stmtContext) {}

// ExitDel_stmt is called when production del_stmt is exited.
func (s *BasePython3Listener) ExitDel_stmt(ctx *Del_stmtContext) {}

// EnterPass_stmt is called when production pass_stmt is entered.
func (s *BasePython3Listener) EnterPass_stmt(ctx *Pass_stmtContext) {}

// ExitPass_stmt is called when production pass_stmt is exited.
func (s *BasePython3Listener) ExitPass_stmt(ctx *Pass_stmtContext) {}

// EnterFlow_stmt is called when production flow_stmt is entered.
func (s *BasePython3Listener) EnterFlow_stmt(ctx *Flow_stmtContext) {}

// ExitFlow_stmt is called when production flow_stmt is exited.
func (s *BasePython3Listener) ExitFlow_stmt(ctx *Flow_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BasePython3Listener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BasePython3Listener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BasePython3Listener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BasePython3Listener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BasePython3Listener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BasePython3Listener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterYield_stmt is called when production yield_stmt is entered.
func (s *BasePython3Listener) EnterYield_stmt(ctx *Yield_stmtContext) {}

// ExitYield_stmt is called when production yield_stmt is exited.
func (s *BasePython3Listener) ExitYield_stmt(ctx *Yield_stmtContext) {}

// EnterRaise_stmt is called when production raise_stmt is entered.
func (s *BasePython3Listener) EnterRaise_stmt(ctx *Raise_stmtContext) {}

// ExitRaise_stmt is called when production raise_stmt is exited.
func (s *BasePython3Listener) ExitRaise_stmt(ctx *Raise_stmtContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BasePython3Listener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BasePython3Listener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterImport_name is called when production import_name is entered.
func (s *BasePython3Listener) EnterImport_name(ctx *Import_nameContext) {}

// ExitImport_name is called when production import_name is exited.
func (s *BasePython3Listener) ExitImport_name(ctx *Import_nameContext) {}

// EnterImport_from is called when production import_from is entered.
func (s *BasePython3Listener) EnterImport_from(ctx *Import_fromContext) {}

// ExitImport_from is called when production import_from is exited.
func (s *BasePython3Listener) ExitImport_from(ctx *Import_fromContext) {}

// EnterImport_as_name is called when production import_as_name is entered.
func (s *BasePython3Listener) EnterImport_as_name(ctx *Import_as_nameContext) {}

// ExitImport_as_name is called when production import_as_name is exited.
func (s *BasePython3Listener) ExitImport_as_name(ctx *Import_as_nameContext) {}

// EnterDotted_as_name is called when production dotted_as_name is entered.
func (s *BasePython3Listener) EnterDotted_as_name(ctx *Dotted_as_nameContext) {}

// ExitDotted_as_name is called when production dotted_as_name is exited.
func (s *BasePython3Listener) ExitDotted_as_name(ctx *Dotted_as_nameContext) {}

// EnterImport_as_names is called when production import_as_names is entered.
func (s *BasePython3Listener) EnterImport_as_names(ctx *Import_as_namesContext) {}

// ExitImport_as_names is called when production import_as_names is exited.
func (s *BasePython3Listener) ExitImport_as_names(ctx *Import_as_namesContext) {}

// EnterDotted_as_names is called when production dotted_as_names is entered.
func (s *BasePython3Listener) EnterDotted_as_names(ctx *Dotted_as_namesContext) {}

// ExitDotted_as_names is called when production dotted_as_names is exited.
func (s *BasePython3Listener) ExitDotted_as_names(ctx *Dotted_as_namesContext) {}

// EnterDotted_name is called when production dotted_name is entered.
func (s *BasePython3Listener) EnterDotted_name(ctx *Dotted_nameContext) {}

// ExitDotted_name is called when production dotted_name is exited.
func (s *BasePython3Listener) ExitDotted_name(ctx *Dotted_nameContext) {}

// EnterGlobal_stmt is called when production global_stmt is entered.
func (s *BasePython3Listener) EnterGlobal_stmt(ctx *Global_stmtContext) {}

// ExitGlobal_stmt is called when production global_stmt is exited.
func (s *BasePython3Listener) ExitGlobal_stmt(ctx *Global_stmtContext) {}

// EnterNonlocal_stmt is called when production nonlocal_stmt is entered.
func (s *BasePython3Listener) EnterNonlocal_stmt(ctx *Nonlocal_stmtContext) {}

// ExitNonlocal_stmt is called when production nonlocal_stmt is exited.
func (s *BasePython3Listener) ExitNonlocal_stmt(ctx *Nonlocal_stmtContext) {}

// EnterAssert_stmt is called when production assert_stmt is entered.
func (s *BasePython3Listener) EnterAssert_stmt(ctx *Assert_stmtContext) {}

// ExitAssert_stmt is called when production assert_stmt is exited.
func (s *BasePython3Listener) ExitAssert_stmt(ctx *Assert_stmtContext) {}

// EnterCompound_stmt is called when production compound_stmt is entered.
func (s *BasePython3Listener) EnterCompound_stmt(ctx *Compound_stmtContext) {}

// ExitCompound_stmt is called when production compound_stmt is exited.
func (s *BasePython3Listener) ExitCompound_stmt(ctx *Compound_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BasePython3Listener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BasePython3Listener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BasePython3Listener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BasePython3Listener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BasePython3Listener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BasePython3Listener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterTry_stmt is called when production try_stmt is entered.
func (s *BasePython3Listener) EnterTry_stmt(ctx *Try_stmtContext) {}

// ExitTry_stmt is called when production try_stmt is exited.
func (s *BasePython3Listener) ExitTry_stmt(ctx *Try_stmtContext) {}

// EnterWith_stmt is called when production with_stmt is entered.
func (s *BasePython3Listener) EnterWith_stmt(ctx *With_stmtContext) {}

// ExitWith_stmt is called when production with_stmt is exited.
func (s *BasePython3Listener) ExitWith_stmt(ctx *With_stmtContext) {}

// EnterWith_item is called when production with_item is entered.
func (s *BasePython3Listener) EnterWith_item(ctx *With_itemContext) {}

// ExitWith_item is called when production with_item is exited.
func (s *BasePython3Listener) ExitWith_item(ctx *With_itemContext) {}

// EnterExcept_clause is called when production except_clause is entered.
func (s *BasePython3Listener) EnterExcept_clause(ctx *Except_clauseContext) {}

// ExitExcept_clause is called when production except_clause is exited.
func (s *BasePython3Listener) ExitExcept_clause(ctx *Except_clauseContext) {}

// EnterSuite is called when production suite is entered.
func (s *BasePython3Listener) EnterSuite(ctx *SuiteContext) {}

// ExitSuite is called when production suite is exited.
func (s *BasePython3Listener) ExitSuite(ctx *SuiteContext) {}

// EnterTest is called when production test is entered.
func (s *BasePython3Listener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasePython3Listener) ExitTest(ctx *TestContext) {}

// EnterTest_nocond is called when production test_nocond is entered.
func (s *BasePython3Listener) EnterTest_nocond(ctx *Test_nocondContext) {}

// ExitTest_nocond is called when production test_nocond is exited.
func (s *BasePython3Listener) ExitTest_nocond(ctx *Test_nocondContext) {}

// EnterLambdef is called when production lambdef is entered.
func (s *BasePython3Listener) EnterLambdef(ctx *LambdefContext) {}

// ExitLambdef is called when production lambdef is exited.
func (s *BasePython3Listener) ExitLambdef(ctx *LambdefContext) {}

// EnterLambdef_nocond is called when production lambdef_nocond is entered.
func (s *BasePython3Listener) EnterLambdef_nocond(ctx *Lambdef_nocondContext) {}

// ExitLambdef_nocond is called when production lambdef_nocond is exited.
func (s *BasePython3Listener) ExitLambdef_nocond(ctx *Lambdef_nocondContext) {}

// EnterOr_test is called when production or_test is entered.
func (s *BasePython3Listener) EnterOr_test(ctx *Or_testContext) {}

// ExitOr_test is called when production or_test is exited.
func (s *BasePython3Listener) ExitOr_test(ctx *Or_testContext) {}

// EnterAnd_test is called when production and_test is entered.
func (s *BasePython3Listener) EnterAnd_test(ctx *And_testContext) {}

// ExitAnd_test is called when production and_test is exited.
func (s *BasePython3Listener) ExitAnd_test(ctx *And_testContext) {}

// EnterNot_test is called when production not_test is entered.
func (s *BasePython3Listener) EnterNot_test(ctx *Not_testContext) {}

// ExitNot_test is called when production not_test is exited.
func (s *BasePython3Listener) ExitNot_test(ctx *Not_testContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasePython3Listener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasePython3Listener) ExitComparison(ctx *ComparisonContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BasePython3Listener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BasePython3Listener) ExitComp_op(ctx *Comp_opContext) {}

// EnterStar_expr is called when production star_expr is entered.
func (s *BasePython3Listener) EnterStar_expr(ctx *Star_exprContext) {}

// ExitStar_expr is called when production star_expr is exited.
func (s *BasePython3Listener) ExitStar_expr(ctx *Star_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePython3Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePython3Listener) ExitExpr(ctx *ExprContext) {}

// EnterXor_expr is called when production xor_expr is entered.
func (s *BasePython3Listener) EnterXor_expr(ctx *Xor_exprContext) {}

// ExitXor_expr is called when production xor_expr is exited.
func (s *BasePython3Listener) ExitXor_expr(ctx *Xor_exprContext) {}

// EnterAnd_expr is called when production and_expr is entered.
func (s *BasePython3Listener) EnterAnd_expr(ctx *And_exprContext) {}

// ExitAnd_expr is called when production and_expr is exited.
func (s *BasePython3Listener) ExitAnd_expr(ctx *And_exprContext) {}

// EnterShift_expr is called when production shift_expr is entered.
func (s *BasePython3Listener) EnterShift_expr(ctx *Shift_exprContext) {}

// ExitShift_expr is called when production shift_expr is exited.
func (s *BasePython3Listener) ExitShift_expr(ctx *Shift_exprContext) {}

// EnterArith_expr is called when production arith_expr is entered.
func (s *BasePython3Listener) EnterArith_expr(ctx *Arith_exprContext) {}

// ExitArith_expr is called when production arith_expr is exited.
func (s *BasePython3Listener) ExitArith_expr(ctx *Arith_exprContext) {}

// EnterTerm is called when production term is entered.
func (s *BasePython3Listener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasePython3Listener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasePython3Listener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasePython3Listener) ExitFactor(ctx *FactorContext) {}

// EnterPower is called when production power is entered.
func (s *BasePython3Listener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BasePython3Listener) ExitPower(ctx *PowerContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePython3Listener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePython3Listener) ExitAtom(ctx *AtomContext) {}

// EnterTestlist_comp is called when production testlist_comp is entered.
func (s *BasePython3Listener) EnterTestlist_comp(ctx *Testlist_compContext) {}

// ExitTestlist_comp is called when production testlist_comp is exited.
func (s *BasePython3Listener) ExitTestlist_comp(ctx *Testlist_compContext) {}

// EnterTrailer is called when production trailer is entered.
func (s *BasePython3Listener) EnterTrailer(ctx *TrailerContext) {}

// ExitTrailer is called when production trailer is exited.
func (s *BasePython3Listener) ExitTrailer(ctx *TrailerContext) {}

// EnterSubscriptlist is called when production subscriptlist is entered.
func (s *BasePython3Listener) EnterSubscriptlist(ctx *SubscriptlistContext) {}

// ExitSubscriptlist is called when production subscriptlist is exited.
func (s *BasePython3Listener) ExitSubscriptlist(ctx *SubscriptlistContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BasePython3Listener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BasePython3Listener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSliceop is called when production sliceop is entered.
func (s *BasePython3Listener) EnterSliceop(ctx *SliceopContext) {}

// ExitSliceop is called when production sliceop is exited.
func (s *BasePython3Listener) ExitSliceop(ctx *SliceopContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BasePython3Listener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BasePython3Listener) ExitExprlist(ctx *ExprlistContext) {}

// EnterTestlist is called when production testlist is entered.
func (s *BasePython3Listener) EnterTestlist(ctx *TestlistContext) {}

// ExitTestlist is called when production testlist is exited.
func (s *BasePython3Listener) ExitTestlist(ctx *TestlistContext) {}

// EnterDictorsetmaker is called when production dictorsetmaker is entered.
func (s *BasePython3Listener) EnterDictorsetmaker(ctx *DictorsetmakerContext) {}

// ExitDictorsetmaker is called when production dictorsetmaker is exited.
func (s *BasePython3Listener) ExitDictorsetmaker(ctx *DictorsetmakerContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BasePython3Listener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BasePython3Listener) ExitClassdef(ctx *ClassdefContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BasePython3Listener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BasePython3Listener) ExitArglist(ctx *ArglistContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasePython3Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasePython3Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterComp_iter is called when production comp_iter is entered.
func (s *BasePython3Listener) EnterComp_iter(ctx *Comp_iterContext) {}

// ExitComp_iter is called when production comp_iter is exited.
func (s *BasePython3Listener) ExitComp_iter(ctx *Comp_iterContext) {}

// EnterComp_for is called when production comp_for is entered.
func (s *BasePython3Listener) EnterComp_for(ctx *Comp_forContext) {}

// ExitComp_for is called when production comp_for is exited.
func (s *BasePython3Listener) ExitComp_for(ctx *Comp_forContext) {}

// EnterComp_if is called when production comp_if is entered.
func (s *BasePython3Listener) EnterComp_if(ctx *Comp_ifContext) {}

// ExitComp_if is called when production comp_if is exited.
func (s *BasePython3Listener) ExitComp_if(ctx *Comp_ifContext) {}

// EnterYield_expr is called when production yield_expr is entered.
func (s *BasePython3Listener) EnterYield_expr(ctx *Yield_exprContext) {}

// ExitYield_expr is called when production yield_expr is exited.
func (s *BasePython3Listener) ExitYield_expr(ctx *Yield_exprContext) {}

// EnterYield_arg is called when production yield_arg is entered.
func (s *BasePython3Listener) EnterYield_arg(ctx *Yield_argContext) {}

// ExitYield_arg is called when production yield_arg is exited.
func (s *BasePython3Listener) ExitYield_arg(ctx *Yield_argContext) {}

// EnterStr is called when production str is entered.
func (s *BasePython3Listener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BasePython3Listener) ExitStr(ctx *StrContext) {}

// EnterNumber is called when production number is entered.
func (s *BasePython3Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasePython3Listener) ExitNumber(ctx *NumberContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasePython3Listener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasePython3Listener) ExitInteger(ctx *IntegerContext) {}
