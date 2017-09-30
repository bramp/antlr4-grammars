// Generated from PeopleCode.g4 by ANTLR 4.7.

package peoplecode // PeopleCode
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePeopleCodeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePeopleCodeVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtList(ctx *StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtAppClassImport(ctx *StmtAppClassImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtExternalFuncImport(ctx *StmtExternalFuncImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtClassDeclaration(ctx *StmtClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtMethodImpl(ctx *StmtMethodImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtGetImpl(ctx *StmtGetImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtSetImpl(ctx *StmtSetImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtFuncImpl(ctx *StmtFuncImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtVarDeclaration(ctx *StmtVarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtIf(ctx *StmtIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtFor(ctx *StmtForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtWhile(ctx *StmtWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtEvaluate(ctx *StmtEvaluateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtTryCatch(ctx *StmtTryCatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtExit(ctx *StmtExitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtBreak(ctx *StmtBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtError(ctx *StmtErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtWarning(ctx *StmtWarningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtReturn(ctx *StmtReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtThrow(ctx *StmtThrowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtAssign(ctx *StmtAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitStmtExpr(ctx *StmtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprComparison(ctx *ExprComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprConcat(ctx *ExprConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprCreate(ctx *ExprCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprBoolean(ctx *ExprBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprNot(ctx *ExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprAddSub(ctx *ExprAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprDotAccess(ctx *ExprDotAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprFnOrIdxCall(ctx *ExprFnOrIdxCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprParenthesized(ctx *ExprParenthesizedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprMulDiv(ctx *ExprMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprNegate(ctx *ExprNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprArrayIndex(ctx *ExprArrayIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprLiteral(ctx *ExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprEquality(ctx *ExprEqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprDynamicReference(ctx *ExprDynamicReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprId(ctx *ExprIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitVarDeclarator(ctx *VarDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitAppClassImport(ctx *AppClassImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitAppPkgPath(ctx *AppPkgPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitAppClassPath(ctx *AppClassPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitExtFuncImport(ctx *ExtFuncImportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitRecDefnPath(ctx *RecDefnPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitClassBlock(ctx *ClassBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitClassBlockStmt(ctx *ClassBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitMethod(ctx *MethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitInstance(ctx *InstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitMethodImpl(ctx *MethodImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitGetImpl(ctx *GetImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitSetImpl(ctx *SetImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitFuncImpl(ctx *FuncImplContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitFuncSignature(ctx *FuncSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitFormalParamList(ctx *FormalParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitReturnType(ctx *ReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitEvaluateStmt(ctx *EvaluateStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitWhenBranch(ctx *WhenBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitWhenOtherBranch(ctx *WhenOtherBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitTryCatchStmt(ctx *TryCatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitCatchSignature(ctx *CatchSignatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitCreateInvocation(ctx *CreateInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePeopleCodeVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}
