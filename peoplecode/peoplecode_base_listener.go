// Code generated from PeopleCode.g4 by ANTLR 4.7.2. DO NOT EDIT.

package peoplecode // PeopleCode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePeopleCodeListener is a complete listener for a parse tree produced by PeopleCodeParser.
type BasePeopleCodeListener struct{}

var _ PeopleCodeListener = &BasePeopleCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePeopleCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePeopleCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePeopleCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePeopleCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePeopleCodeListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePeopleCodeListener) ExitProgram(ctx *ProgramContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BasePeopleCodeListener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BasePeopleCodeListener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmtAppClassImport is called when production StmtAppClassImport is entered.
func (s *BasePeopleCodeListener) EnterStmtAppClassImport(ctx *StmtAppClassImportContext) {}

// ExitStmtAppClassImport is called when production StmtAppClassImport is exited.
func (s *BasePeopleCodeListener) ExitStmtAppClassImport(ctx *StmtAppClassImportContext) {}

// EnterStmtExternalFuncImport is called when production StmtExternalFuncImport is entered.
func (s *BasePeopleCodeListener) EnterStmtExternalFuncImport(ctx *StmtExternalFuncImportContext) {}

// ExitStmtExternalFuncImport is called when production StmtExternalFuncImport is exited.
func (s *BasePeopleCodeListener) ExitStmtExternalFuncImport(ctx *StmtExternalFuncImportContext) {}

// EnterStmtClassDeclaration is called when production StmtClassDeclaration is entered.
func (s *BasePeopleCodeListener) EnterStmtClassDeclaration(ctx *StmtClassDeclarationContext) {}

// ExitStmtClassDeclaration is called when production StmtClassDeclaration is exited.
func (s *BasePeopleCodeListener) ExitStmtClassDeclaration(ctx *StmtClassDeclarationContext) {}

// EnterStmtMethodImpl is called when production StmtMethodImpl is entered.
func (s *BasePeopleCodeListener) EnterStmtMethodImpl(ctx *StmtMethodImplContext) {}

// ExitStmtMethodImpl is called when production StmtMethodImpl is exited.
func (s *BasePeopleCodeListener) ExitStmtMethodImpl(ctx *StmtMethodImplContext) {}

// EnterStmtGetImpl is called when production StmtGetImpl is entered.
func (s *BasePeopleCodeListener) EnterStmtGetImpl(ctx *StmtGetImplContext) {}

// ExitStmtGetImpl is called when production StmtGetImpl is exited.
func (s *BasePeopleCodeListener) ExitStmtGetImpl(ctx *StmtGetImplContext) {}

// EnterStmtSetImpl is called when production StmtSetImpl is entered.
func (s *BasePeopleCodeListener) EnterStmtSetImpl(ctx *StmtSetImplContext) {}

// ExitStmtSetImpl is called when production StmtSetImpl is exited.
func (s *BasePeopleCodeListener) ExitStmtSetImpl(ctx *StmtSetImplContext) {}

// EnterStmtFuncImpl is called when production StmtFuncImpl is entered.
func (s *BasePeopleCodeListener) EnterStmtFuncImpl(ctx *StmtFuncImplContext) {}

// ExitStmtFuncImpl is called when production StmtFuncImpl is exited.
func (s *BasePeopleCodeListener) ExitStmtFuncImpl(ctx *StmtFuncImplContext) {}

// EnterStmtVarDeclaration is called when production StmtVarDeclaration is entered.
func (s *BasePeopleCodeListener) EnterStmtVarDeclaration(ctx *StmtVarDeclarationContext) {}

// ExitStmtVarDeclaration is called when production StmtVarDeclaration is exited.
func (s *BasePeopleCodeListener) ExitStmtVarDeclaration(ctx *StmtVarDeclarationContext) {}

// EnterStmtIf is called when production StmtIf is entered.
func (s *BasePeopleCodeListener) EnterStmtIf(ctx *StmtIfContext) {}

// ExitStmtIf is called when production StmtIf is exited.
func (s *BasePeopleCodeListener) ExitStmtIf(ctx *StmtIfContext) {}

// EnterStmtFor is called when production StmtFor is entered.
func (s *BasePeopleCodeListener) EnterStmtFor(ctx *StmtForContext) {}

// ExitStmtFor is called when production StmtFor is exited.
func (s *BasePeopleCodeListener) ExitStmtFor(ctx *StmtForContext) {}

// EnterStmtWhile is called when production StmtWhile is entered.
func (s *BasePeopleCodeListener) EnterStmtWhile(ctx *StmtWhileContext) {}

// ExitStmtWhile is called when production StmtWhile is exited.
func (s *BasePeopleCodeListener) ExitStmtWhile(ctx *StmtWhileContext) {}

// EnterStmtEvaluate is called when production StmtEvaluate is entered.
func (s *BasePeopleCodeListener) EnterStmtEvaluate(ctx *StmtEvaluateContext) {}

// ExitStmtEvaluate is called when production StmtEvaluate is exited.
func (s *BasePeopleCodeListener) ExitStmtEvaluate(ctx *StmtEvaluateContext) {}

// EnterStmtTryCatch is called when production StmtTryCatch is entered.
func (s *BasePeopleCodeListener) EnterStmtTryCatch(ctx *StmtTryCatchContext) {}

// ExitStmtTryCatch is called when production StmtTryCatch is exited.
func (s *BasePeopleCodeListener) ExitStmtTryCatch(ctx *StmtTryCatchContext) {}

// EnterStmtExit is called when production StmtExit is entered.
func (s *BasePeopleCodeListener) EnterStmtExit(ctx *StmtExitContext) {}

// ExitStmtExit is called when production StmtExit is exited.
func (s *BasePeopleCodeListener) ExitStmtExit(ctx *StmtExitContext) {}

// EnterStmtBreak is called when production StmtBreak is entered.
func (s *BasePeopleCodeListener) EnterStmtBreak(ctx *StmtBreakContext) {}

// ExitStmtBreak is called when production StmtBreak is exited.
func (s *BasePeopleCodeListener) ExitStmtBreak(ctx *StmtBreakContext) {}

// EnterStmtError is called when production StmtError is entered.
func (s *BasePeopleCodeListener) EnterStmtError(ctx *StmtErrorContext) {}

// ExitStmtError is called when production StmtError is exited.
func (s *BasePeopleCodeListener) ExitStmtError(ctx *StmtErrorContext) {}

// EnterStmtWarning is called when production StmtWarning is entered.
func (s *BasePeopleCodeListener) EnterStmtWarning(ctx *StmtWarningContext) {}

// ExitStmtWarning is called when production StmtWarning is exited.
func (s *BasePeopleCodeListener) ExitStmtWarning(ctx *StmtWarningContext) {}

// EnterStmtReturn is called when production StmtReturn is entered.
func (s *BasePeopleCodeListener) EnterStmtReturn(ctx *StmtReturnContext) {}

// ExitStmtReturn is called when production StmtReturn is exited.
func (s *BasePeopleCodeListener) ExitStmtReturn(ctx *StmtReturnContext) {}

// EnterStmtThrow is called when production StmtThrow is entered.
func (s *BasePeopleCodeListener) EnterStmtThrow(ctx *StmtThrowContext) {}

// ExitStmtThrow is called when production StmtThrow is exited.
func (s *BasePeopleCodeListener) ExitStmtThrow(ctx *StmtThrowContext) {}

// EnterStmtAssign is called when production StmtAssign is entered.
func (s *BasePeopleCodeListener) EnterStmtAssign(ctx *StmtAssignContext) {}

// ExitStmtAssign is called when production StmtAssign is exited.
func (s *BasePeopleCodeListener) ExitStmtAssign(ctx *StmtAssignContext) {}

// EnterStmtExpr is called when production StmtExpr is entered.
func (s *BasePeopleCodeListener) EnterStmtExpr(ctx *StmtExprContext) {}

// ExitStmtExpr is called when production StmtExpr is exited.
func (s *BasePeopleCodeListener) ExitStmtExpr(ctx *StmtExprContext) {}

// EnterExprComparison is called when production ExprComparison is entered.
func (s *BasePeopleCodeListener) EnterExprComparison(ctx *ExprComparisonContext) {}

// ExitExprComparison is called when production ExprComparison is exited.
func (s *BasePeopleCodeListener) ExitExprComparison(ctx *ExprComparisonContext) {}

// EnterExprConcat is called when production ExprConcat is entered.
func (s *BasePeopleCodeListener) EnterExprConcat(ctx *ExprConcatContext) {}

// ExitExprConcat is called when production ExprConcat is exited.
func (s *BasePeopleCodeListener) ExitExprConcat(ctx *ExprConcatContext) {}

// EnterExprCreate is called when production ExprCreate is entered.
func (s *BasePeopleCodeListener) EnterExprCreate(ctx *ExprCreateContext) {}

// ExitExprCreate is called when production ExprCreate is exited.
func (s *BasePeopleCodeListener) ExitExprCreate(ctx *ExprCreateContext) {}

// EnterExprBoolean is called when production ExprBoolean is entered.
func (s *BasePeopleCodeListener) EnterExprBoolean(ctx *ExprBooleanContext) {}

// ExitExprBoolean is called when production ExprBoolean is exited.
func (s *BasePeopleCodeListener) ExitExprBoolean(ctx *ExprBooleanContext) {}

// EnterExprNot is called when production ExprNot is entered.
func (s *BasePeopleCodeListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production ExprNot is exited.
func (s *BasePeopleCodeListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprAddSub is called when production ExprAddSub is entered.
func (s *BasePeopleCodeListener) EnterExprAddSub(ctx *ExprAddSubContext) {}

// ExitExprAddSub is called when production ExprAddSub is exited.
func (s *BasePeopleCodeListener) ExitExprAddSub(ctx *ExprAddSubContext) {}

// EnterExprDotAccess is called when production ExprDotAccess is entered.
func (s *BasePeopleCodeListener) EnterExprDotAccess(ctx *ExprDotAccessContext) {}

// ExitExprDotAccess is called when production ExprDotAccess is exited.
func (s *BasePeopleCodeListener) ExitExprDotAccess(ctx *ExprDotAccessContext) {}

// EnterExprFnOrIdxCall is called when production ExprFnOrIdxCall is entered.
func (s *BasePeopleCodeListener) EnterExprFnOrIdxCall(ctx *ExprFnOrIdxCallContext) {}

// ExitExprFnOrIdxCall is called when production ExprFnOrIdxCall is exited.
func (s *BasePeopleCodeListener) ExitExprFnOrIdxCall(ctx *ExprFnOrIdxCallContext) {}

// EnterExprParenthesized is called when production ExprParenthesized is entered.
func (s *BasePeopleCodeListener) EnterExprParenthesized(ctx *ExprParenthesizedContext) {}

// ExitExprParenthesized is called when production ExprParenthesized is exited.
func (s *BasePeopleCodeListener) ExitExprParenthesized(ctx *ExprParenthesizedContext) {}

// EnterExprMulDiv is called when production ExprMulDiv is entered.
func (s *BasePeopleCodeListener) EnterExprMulDiv(ctx *ExprMulDivContext) {}

// ExitExprMulDiv is called when production ExprMulDiv is exited.
func (s *BasePeopleCodeListener) ExitExprMulDiv(ctx *ExprMulDivContext) {}

// EnterExprNegate is called when production ExprNegate is entered.
func (s *BasePeopleCodeListener) EnterExprNegate(ctx *ExprNegateContext) {}

// ExitExprNegate is called when production ExprNegate is exited.
func (s *BasePeopleCodeListener) ExitExprNegate(ctx *ExprNegateContext) {}

// EnterExprArrayIndex is called when production ExprArrayIndex is entered.
func (s *BasePeopleCodeListener) EnterExprArrayIndex(ctx *ExprArrayIndexContext) {}

// ExitExprArrayIndex is called when production ExprArrayIndex is exited.
func (s *BasePeopleCodeListener) ExitExprArrayIndex(ctx *ExprArrayIndexContext) {}

// EnterExprLiteral is called when production ExprLiteral is entered.
func (s *BasePeopleCodeListener) EnterExprLiteral(ctx *ExprLiteralContext) {}

// ExitExprLiteral is called when production ExprLiteral is exited.
func (s *BasePeopleCodeListener) ExitExprLiteral(ctx *ExprLiteralContext) {}

// EnterExprEquality is called when production ExprEquality is entered.
func (s *BasePeopleCodeListener) EnterExprEquality(ctx *ExprEqualityContext) {}

// ExitExprEquality is called when production ExprEquality is exited.
func (s *BasePeopleCodeListener) ExitExprEquality(ctx *ExprEqualityContext) {}

// EnterExprDynamicReference is called when production ExprDynamicReference is entered.
func (s *BasePeopleCodeListener) EnterExprDynamicReference(ctx *ExprDynamicReferenceContext) {}

// ExitExprDynamicReference is called when production ExprDynamicReference is exited.
func (s *BasePeopleCodeListener) ExitExprDynamicReference(ctx *ExprDynamicReferenceContext) {}

// EnterExprId is called when production ExprId is entered.
func (s *BasePeopleCodeListener) EnterExprId(ctx *ExprIdContext) {}

// ExitExprId is called when production ExprId is exited.
func (s *BasePeopleCodeListener) ExitExprId(ctx *ExprIdContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BasePeopleCodeListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BasePeopleCodeListener) ExitExprList(ctx *ExprListContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BasePeopleCodeListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BasePeopleCodeListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterVarDeclarator is called when production varDeclarator is entered.
func (s *BasePeopleCodeListener) EnterVarDeclarator(ctx *VarDeclaratorContext) {}

// ExitVarDeclarator is called when production varDeclarator is exited.
func (s *BasePeopleCodeListener) ExitVarDeclarator(ctx *VarDeclaratorContext) {}

// EnterVarType is called when production varType is entered.
func (s *BasePeopleCodeListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BasePeopleCodeListener) ExitVarType(ctx *VarTypeContext) {}

// EnterAppClassImport is called when production appClassImport is entered.
func (s *BasePeopleCodeListener) EnterAppClassImport(ctx *AppClassImportContext) {}

// ExitAppClassImport is called when production appClassImport is exited.
func (s *BasePeopleCodeListener) ExitAppClassImport(ctx *AppClassImportContext) {}

// EnterAppPkgPath is called when production appPkgPath is entered.
func (s *BasePeopleCodeListener) EnterAppPkgPath(ctx *AppPkgPathContext) {}

// ExitAppPkgPath is called when production appPkgPath is exited.
func (s *BasePeopleCodeListener) ExitAppPkgPath(ctx *AppPkgPathContext) {}

// EnterAppClassPath is called when production appClassPath is entered.
func (s *BasePeopleCodeListener) EnterAppClassPath(ctx *AppClassPathContext) {}

// ExitAppClassPath is called when production appClassPath is exited.
func (s *BasePeopleCodeListener) ExitAppClassPath(ctx *AppClassPathContext) {}

// EnterExtFuncImport is called when production extFuncImport is entered.
func (s *BasePeopleCodeListener) EnterExtFuncImport(ctx *ExtFuncImportContext) {}

// ExitExtFuncImport is called when production extFuncImport is exited.
func (s *BasePeopleCodeListener) ExitExtFuncImport(ctx *ExtFuncImportContext) {}

// EnterRecDefnPath is called when production recDefnPath is entered.
func (s *BasePeopleCodeListener) EnterRecDefnPath(ctx *RecDefnPathContext) {}

// ExitRecDefnPath is called when production recDefnPath is exited.
func (s *BasePeopleCodeListener) ExitRecDefnPath(ctx *RecDefnPathContext) {}

// EnterEvent is called when production event is entered.
func (s *BasePeopleCodeListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BasePeopleCodeListener) ExitEvent(ctx *EventContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasePeopleCodeListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasePeopleCodeListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBlock is called when production classBlock is entered.
func (s *BasePeopleCodeListener) EnterClassBlock(ctx *ClassBlockContext) {}

// ExitClassBlock is called when production classBlock is exited.
func (s *BasePeopleCodeListener) ExitClassBlock(ctx *ClassBlockContext) {}

// EnterClassBlockStmt is called when production classBlockStmt is entered.
func (s *BasePeopleCodeListener) EnterClassBlockStmt(ctx *ClassBlockStmtContext) {}

// ExitClassBlockStmt is called when production classBlockStmt is exited.
func (s *BasePeopleCodeListener) ExitClassBlockStmt(ctx *ClassBlockStmtContext) {}

// EnterMethod is called when production method is entered.
func (s *BasePeopleCodeListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BasePeopleCodeListener) ExitMethod(ctx *MethodContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasePeopleCodeListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasePeopleCodeListener) ExitConstant(ctx *ConstantContext) {}

// EnterProperty is called when production property is entered.
func (s *BasePeopleCodeListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BasePeopleCodeListener) ExitProperty(ctx *PropertyContext) {}

// EnterInstance is called when production instance is entered.
func (s *BasePeopleCodeListener) EnterInstance(ctx *InstanceContext) {}

// ExitInstance is called when production instance is exited.
func (s *BasePeopleCodeListener) ExitInstance(ctx *InstanceContext) {}

// EnterMethodImpl is called when production methodImpl is entered.
func (s *BasePeopleCodeListener) EnterMethodImpl(ctx *MethodImplContext) {}

// ExitMethodImpl is called when production methodImpl is exited.
func (s *BasePeopleCodeListener) ExitMethodImpl(ctx *MethodImplContext) {}

// EnterGetImpl is called when production getImpl is entered.
func (s *BasePeopleCodeListener) EnterGetImpl(ctx *GetImplContext) {}

// ExitGetImpl is called when production getImpl is exited.
func (s *BasePeopleCodeListener) ExitGetImpl(ctx *GetImplContext) {}

// EnterSetImpl is called when production setImpl is entered.
func (s *BasePeopleCodeListener) EnterSetImpl(ctx *SetImplContext) {}

// ExitSetImpl is called when production setImpl is exited.
func (s *BasePeopleCodeListener) ExitSetImpl(ctx *SetImplContext) {}

// EnterFuncImpl is called when production funcImpl is entered.
func (s *BasePeopleCodeListener) EnterFuncImpl(ctx *FuncImplContext) {}

// ExitFuncImpl is called when production funcImpl is exited.
func (s *BasePeopleCodeListener) ExitFuncImpl(ctx *FuncImplContext) {}

// EnterFuncSignature is called when production funcSignature is entered.
func (s *BasePeopleCodeListener) EnterFuncSignature(ctx *FuncSignatureContext) {}

// ExitFuncSignature is called when production funcSignature is exited.
func (s *BasePeopleCodeListener) ExitFuncSignature(ctx *FuncSignatureContext) {}

// EnterFormalParamList is called when production formalParamList is entered.
func (s *BasePeopleCodeListener) EnterFormalParamList(ctx *FormalParamListContext) {}

// ExitFormalParamList is called when production formalParamList is exited.
func (s *BasePeopleCodeListener) ExitFormalParamList(ctx *FormalParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BasePeopleCodeListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasePeopleCodeListener) ExitParam(ctx *ParamContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BasePeopleCodeListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BasePeopleCodeListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BasePeopleCodeListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BasePeopleCodeListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BasePeopleCodeListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BasePeopleCodeListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BasePeopleCodeListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BasePeopleCodeListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterEvaluateStmt is called when production evaluateStmt is entered.
func (s *BasePeopleCodeListener) EnterEvaluateStmt(ctx *EvaluateStmtContext) {}

// ExitEvaluateStmt is called when production evaluateStmt is exited.
func (s *BasePeopleCodeListener) ExitEvaluateStmt(ctx *EvaluateStmtContext) {}

// EnterWhenBranch is called when production whenBranch is entered.
func (s *BasePeopleCodeListener) EnterWhenBranch(ctx *WhenBranchContext) {}

// ExitWhenBranch is called when production whenBranch is exited.
func (s *BasePeopleCodeListener) ExitWhenBranch(ctx *WhenBranchContext) {}

// EnterWhenOtherBranch is called when production whenOtherBranch is entered.
func (s *BasePeopleCodeListener) EnterWhenOtherBranch(ctx *WhenOtherBranchContext) {}

// ExitWhenOtherBranch is called when production whenOtherBranch is exited.
func (s *BasePeopleCodeListener) ExitWhenOtherBranch(ctx *WhenOtherBranchContext) {}

// EnterTryCatchStmt is called when production tryCatchStmt is entered.
func (s *BasePeopleCodeListener) EnterTryCatchStmt(ctx *TryCatchStmtContext) {}

// ExitTryCatchStmt is called when production tryCatchStmt is exited.
func (s *BasePeopleCodeListener) ExitTryCatchStmt(ctx *TryCatchStmtContext) {}

// EnterCatchSignature is called when production catchSignature is entered.
func (s *BasePeopleCodeListener) EnterCatchSignature(ctx *CatchSignatureContext) {}

// ExitCatchSignature is called when production catchSignature is exited.
func (s *BasePeopleCodeListener) ExitCatchSignature(ctx *CatchSignatureContext) {}

// EnterCreateInvocation is called when production createInvocation is entered.
func (s *BasePeopleCodeListener) EnterCreateInvocation(ctx *CreateInvocationContext) {}

// ExitCreateInvocation is called when production createInvocation is exited.
func (s *BasePeopleCodeListener) ExitCreateInvocation(ctx *CreateInvocationContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePeopleCodeListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePeopleCodeListener) ExitLiteral(ctx *LiteralContext) {}

// EnterId is called when production id is entered.
func (s *BasePeopleCodeListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BasePeopleCodeListener) ExitId(ctx *IdContext) {}
