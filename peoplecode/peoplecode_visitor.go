// Generated from PeopleCode.g4 by ANTLR 4.7.

package peoplecode // PeopleCode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PeopleCodeParser.
type PeopleCodeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PeopleCodeParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtAppClassImport.
	VisitStmtAppClassImport(ctx *StmtAppClassImportContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtExternalFuncImport.
	VisitStmtExternalFuncImport(ctx *StmtExternalFuncImportContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtClassDeclaration.
	VisitStmtClassDeclaration(ctx *StmtClassDeclarationContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtMethodImpl.
	VisitStmtMethodImpl(ctx *StmtMethodImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtGetImpl.
	VisitStmtGetImpl(ctx *StmtGetImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtSetImpl.
	VisitStmtSetImpl(ctx *StmtSetImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtFuncImpl.
	VisitStmtFuncImpl(ctx *StmtFuncImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtVarDeclaration.
	VisitStmtVarDeclaration(ctx *StmtVarDeclarationContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtIf.
	VisitStmtIf(ctx *StmtIfContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtFor.
	VisitStmtFor(ctx *StmtForContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtWhile.
	VisitStmtWhile(ctx *StmtWhileContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtEvaluate.
	VisitStmtEvaluate(ctx *StmtEvaluateContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtTryCatch.
	VisitStmtTryCatch(ctx *StmtTryCatchContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtExit.
	VisitStmtExit(ctx *StmtExitContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtBreak.
	VisitStmtBreak(ctx *StmtBreakContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtError.
	VisitStmtError(ctx *StmtErrorContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtWarning.
	VisitStmtWarning(ctx *StmtWarningContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtReturn.
	VisitStmtReturn(ctx *StmtReturnContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtThrow.
	VisitStmtThrow(ctx *StmtThrowContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtAssign.
	VisitStmtAssign(ctx *StmtAssignContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#StmtExpr.
	VisitStmtExpr(ctx *StmtExprContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprComparison.
	VisitExprComparison(ctx *ExprComparisonContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprConcat.
	VisitExprConcat(ctx *ExprConcatContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprCreate.
	VisitExprCreate(ctx *ExprCreateContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprBoolean.
	VisitExprBoolean(ctx *ExprBooleanContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprNot.
	VisitExprNot(ctx *ExprNotContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprAddSub.
	VisitExprAddSub(ctx *ExprAddSubContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprDotAccess.
	VisitExprDotAccess(ctx *ExprDotAccessContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprFnOrIdxCall.
	VisitExprFnOrIdxCall(ctx *ExprFnOrIdxCallContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprParenthesized.
	VisitExprParenthesized(ctx *ExprParenthesizedContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprMulDiv.
	VisitExprMulDiv(ctx *ExprMulDivContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprNegate.
	VisitExprNegate(ctx *ExprNegateContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprArrayIndex.
	VisitExprArrayIndex(ctx *ExprArrayIndexContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprLiteral.
	VisitExprLiteral(ctx *ExprLiteralContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprEquality.
	VisitExprEquality(ctx *ExprEqualityContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprDynamicReference.
	VisitExprDynamicReference(ctx *ExprDynamicReferenceContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ExprId.
	VisitExprId(ctx *ExprIdContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#varDeclarator.
	VisitVarDeclarator(ctx *VarDeclaratorContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#appClassImport.
	VisitAppClassImport(ctx *AppClassImportContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#appPkgPath.
	VisitAppPkgPath(ctx *AppPkgPathContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#appClassPath.
	VisitAppClassPath(ctx *AppClassPathContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#extFuncImport.
	VisitExtFuncImport(ctx *ExtFuncImportContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#recDefnPath.
	VisitRecDefnPath(ctx *RecDefnPathContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#classBlock.
	VisitClassBlock(ctx *ClassBlockContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#classBlockStmt.
	VisitClassBlockStmt(ctx *ClassBlockStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#instance.
	VisitInstance(ctx *InstanceContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#methodImpl.
	VisitMethodImpl(ctx *MethodImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#getImpl.
	VisitGetImpl(ctx *GetImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#setImpl.
	VisitSetImpl(ctx *SetImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#funcImpl.
	VisitFuncImpl(ctx *FuncImplContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#funcSignature.
	VisitFuncSignature(ctx *FuncSignatureContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#formalParamList.
	VisitFormalParamList(ctx *FormalParamListContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#returnType.
	VisitReturnType(ctx *ReturnTypeContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#evaluateStmt.
	VisitEvaluateStmt(ctx *EvaluateStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#whenBranch.
	VisitWhenBranch(ctx *WhenBranchContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#whenOtherBranch.
	VisitWhenOtherBranch(ctx *WhenOtherBranchContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#tryCatchStmt.
	VisitTryCatchStmt(ctx *TryCatchStmtContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#catchSignature.
	VisitCatchSignature(ctx *CatchSignatureContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#createInvocation.
	VisitCreateInvocation(ctx *CreateInvocationContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PeopleCodeParser#id.
	VisitId(ctx *IdContext) interface{}
}
