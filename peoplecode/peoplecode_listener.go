// Code generated from PeopleCode.g4 by ANTLR 4.7.2. DO NOT EDIT.

package peoplecode // PeopleCode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PeopleCodeListener is a complete listener for a parse tree produced by PeopleCodeParser.
type PeopleCodeListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmtAppClassImport is called when entering the StmtAppClassImport production.
	EnterStmtAppClassImport(c *StmtAppClassImportContext)

	// EnterStmtExternalFuncImport is called when entering the StmtExternalFuncImport production.
	EnterStmtExternalFuncImport(c *StmtExternalFuncImportContext)

	// EnterStmtClassDeclaration is called when entering the StmtClassDeclaration production.
	EnterStmtClassDeclaration(c *StmtClassDeclarationContext)

	// EnterStmtMethodImpl is called when entering the StmtMethodImpl production.
	EnterStmtMethodImpl(c *StmtMethodImplContext)

	// EnterStmtGetImpl is called when entering the StmtGetImpl production.
	EnterStmtGetImpl(c *StmtGetImplContext)

	// EnterStmtSetImpl is called when entering the StmtSetImpl production.
	EnterStmtSetImpl(c *StmtSetImplContext)

	// EnterStmtFuncImpl is called when entering the StmtFuncImpl production.
	EnterStmtFuncImpl(c *StmtFuncImplContext)

	// EnterStmtVarDeclaration is called when entering the StmtVarDeclaration production.
	EnterStmtVarDeclaration(c *StmtVarDeclarationContext)

	// EnterStmtIf is called when entering the StmtIf production.
	EnterStmtIf(c *StmtIfContext)

	// EnterStmtFor is called when entering the StmtFor production.
	EnterStmtFor(c *StmtForContext)

	// EnterStmtWhile is called when entering the StmtWhile production.
	EnterStmtWhile(c *StmtWhileContext)

	// EnterStmtEvaluate is called when entering the StmtEvaluate production.
	EnterStmtEvaluate(c *StmtEvaluateContext)

	// EnterStmtTryCatch is called when entering the StmtTryCatch production.
	EnterStmtTryCatch(c *StmtTryCatchContext)

	// EnterStmtExit is called when entering the StmtExit production.
	EnterStmtExit(c *StmtExitContext)

	// EnterStmtBreak is called when entering the StmtBreak production.
	EnterStmtBreak(c *StmtBreakContext)

	// EnterStmtError is called when entering the StmtError production.
	EnterStmtError(c *StmtErrorContext)

	// EnterStmtWarning is called when entering the StmtWarning production.
	EnterStmtWarning(c *StmtWarningContext)

	// EnterStmtReturn is called when entering the StmtReturn production.
	EnterStmtReturn(c *StmtReturnContext)

	// EnterStmtThrow is called when entering the StmtThrow production.
	EnterStmtThrow(c *StmtThrowContext)

	// EnterStmtAssign is called when entering the StmtAssign production.
	EnterStmtAssign(c *StmtAssignContext)

	// EnterStmtExpr is called when entering the StmtExpr production.
	EnterStmtExpr(c *StmtExprContext)

	// EnterExprComparison is called when entering the ExprComparison production.
	EnterExprComparison(c *ExprComparisonContext)

	// EnterExprConcat is called when entering the ExprConcat production.
	EnterExprConcat(c *ExprConcatContext)

	// EnterExprCreate is called when entering the ExprCreate production.
	EnterExprCreate(c *ExprCreateContext)

	// EnterExprBoolean is called when entering the ExprBoolean production.
	EnterExprBoolean(c *ExprBooleanContext)

	// EnterExprNot is called when entering the ExprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprAddSub is called when entering the ExprAddSub production.
	EnterExprAddSub(c *ExprAddSubContext)

	// EnterExprDotAccess is called when entering the ExprDotAccess production.
	EnterExprDotAccess(c *ExprDotAccessContext)

	// EnterExprFnOrIdxCall is called when entering the ExprFnOrIdxCall production.
	EnterExprFnOrIdxCall(c *ExprFnOrIdxCallContext)

	// EnterExprParenthesized is called when entering the ExprParenthesized production.
	EnterExprParenthesized(c *ExprParenthesizedContext)

	// EnterExprMulDiv is called when entering the ExprMulDiv production.
	EnterExprMulDiv(c *ExprMulDivContext)

	// EnterExprNegate is called when entering the ExprNegate production.
	EnterExprNegate(c *ExprNegateContext)

	// EnterExprArrayIndex is called when entering the ExprArrayIndex production.
	EnterExprArrayIndex(c *ExprArrayIndexContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprEquality is called when entering the ExprEquality production.
	EnterExprEquality(c *ExprEqualityContext)

	// EnterExprDynamicReference is called when entering the ExprDynamicReference production.
	EnterExprDynamicReference(c *ExprDynamicReferenceContext)

	// EnterExprId is called when entering the ExprId production.
	EnterExprId(c *ExprIdContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVarDeclarator is called when entering the varDeclarator production.
	EnterVarDeclarator(c *VarDeclaratorContext)

	// EnterVarType is called when entering the varType production.
	EnterVarType(c *VarTypeContext)

	// EnterAppClassImport is called when entering the appClassImport production.
	EnterAppClassImport(c *AppClassImportContext)

	// EnterAppPkgPath is called when entering the appPkgPath production.
	EnterAppPkgPath(c *AppPkgPathContext)

	// EnterAppClassPath is called when entering the appClassPath production.
	EnterAppClassPath(c *AppClassPathContext)

	// EnterExtFuncImport is called when entering the extFuncImport production.
	EnterExtFuncImport(c *ExtFuncImportContext)

	// EnterRecDefnPath is called when entering the recDefnPath production.
	EnterRecDefnPath(c *RecDefnPathContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBlock is called when entering the classBlock production.
	EnterClassBlock(c *ClassBlockContext)

	// EnterClassBlockStmt is called when entering the classBlockStmt production.
	EnterClassBlockStmt(c *ClassBlockStmtContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterInstance is called when entering the instance production.
	EnterInstance(c *InstanceContext)

	// EnterMethodImpl is called when entering the methodImpl production.
	EnterMethodImpl(c *MethodImplContext)

	// EnterGetImpl is called when entering the getImpl production.
	EnterGetImpl(c *GetImplContext)

	// EnterSetImpl is called when entering the setImpl production.
	EnterSetImpl(c *SetImplContext)

	// EnterFuncImpl is called when entering the funcImpl production.
	EnterFuncImpl(c *FuncImplContext)

	// EnterFuncSignature is called when entering the funcSignature production.
	EnterFuncSignature(c *FuncSignatureContext)

	// EnterFormalParamList is called when entering the formalParamList production.
	EnterFormalParamList(c *FormalParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterEvaluateStmt is called when entering the evaluateStmt production.
	EnterEvaluateStmt(c *EvaluateStmtContext)

	// EnterWhenBranch is called when entering the whenBranch production.
	EnterWhenBranch(c *WhenBranchContext)

	// EnterWhenOtherBranch is called when entering the whenOtherBranch production.
	EnterWhenOtherBranch(c *WhenOtherBranchContext)

	// EnterTryCatchStmt is called when entering the tryCatchStmt production.
	EnterTryCatchStmt(c *TryCatchStmtContext)

	// EnterCatchSignature is called when entering the catchSignature production.
	EnterCatchSignature(c *CatchSignatureContext)

	// EnterCreateInvocation is called when entering the createInvocation production.
	EnterCreateInvocation(c *CreateInvocationContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmtAppClassImport is called when exiting the StmtAppClassImport production.
	ExitStmtAppClassImport(c *StmtAppClassImportContext)

	// ExitStmtExternalFuncImport is called when exiting the StmtExternalFuncImport production.
	ExitStmtExternalFuncImport(c *StmtExternalFuncImportContext)

	// ExitStmtClassDeclaration is called when exiting the StmtClassDeclaration production.
	ExitStmtClassDeclaration(c *StmtClassDeclarationContext)

	// ExitStmtMethodImpl is called when exiting the StmtMethodImpl production.
	ExitStmtMethodImpl(c *StmtMethodImplContext)

	// ExitStmtGetImpl is called when exiting the StmtGetImpl production.
	ExitStmtGetImpl(c *StmtGetImplContext)

	// ExitStmtSetImpl is called when exiting the StmtSetImpl production.
	ExitStmtSetImpl(c *StmtSetImplContext)

	// ExitStmtFuncImpl is called when exiting the StmtFuncImpl production.
	ExitStmtFuncImpl(c *StmtFuncImplContext)

	// ExitStmtVarDeclaration is called when exiting the StmtVarDeclaration production.
	ExitStmtVarDeclaration(c *StmtVarDeclarationContext)

	// ExitStmtIf is called when exiting the StmtIf production.
	ExitStmtIf(c *StmtIfContext)

	// ExitStmtFor is called when exiting the StmtFor production.
	ExitStmtFor(c *StmtForContext)

	// ExitStmtWhile is called when exiting the StmtWhile production.
	ExitStmtWhile(c *StmtWhileContext)

	// ExitStmtEvaluate is called when exiting the StmtEvaluate production.
	ExitStmtEvaluate(c *StmtEvaluateContext)

	// ExitStmtTryCatch is called when exiting the StmtTryCatch production.
	ExitStmtTryCatch(c *StmtTryCatchContext)

	// ExitStmtExit is called when exiting the StmtExit production.
	ExitStmtExit(c *StmtExitContext)

	// ExitStmtBreak is called when exiting the StmtBreak production.
	ExitStmtBreak(c *StmtBreakContext)

	// ExitStmtError is called when exiting the StmtError production.
	ExitStmtError(c *StmtErrorContext)

	// ExitStmtWarning is called when exiting the StmtWarning production.
	ExitStmtWarning(c *StmtWarningContext)

	// ExitStmtReturn is called when exiting the StmtReturn production.
	ExitStmtReturn(c *StmtReturnContext)

	// ExitStmtThrow is called when exiting the StmtThrow production.
	ExitStmtThrow(c *StmtThrowContext)

	// ExitStmtAssign is called when exiting the StmtAssign production.
	ExitStmtAssign(c *StmtAssignContext)

	// ExitStmtExpr is called when exiting the StmtExpr production.
	ExitStmtExpr(c *StmtExprContext)

	// ExitExprComparison is called when exiting the ExprComparison production.
	ExitExprComparison(c *ExprComparisonContext)

	// ExitExprConcat is called when exiting the ExprConcat production.
	ExitExprConcat(c *ExprConcatContext)

	// ExitExprCreate is called when exiting the ExprCreate production.
	ExitExprCreate(c *ExprCreateContext)

	// ExitExprBoolean is called when exiting the ExprBoolean production.
	ExitExprBoolean(c *ExprBooleanContext)

	// ExitExprNot is called when exiting the ExprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprAddSub is called when exiting the ExprAddSub production.
	ExitExprAddSub(c *ExprAddSubContext)

	// ExitExprDotAccess is called when exiting the ExprDotAccess production.
	ExitExprDotAccess(c *ExprDotAccessContext)

	// ExitExprFnOrIdxCall is called when exiting the ExprFnOrIdxCall production.
	ExitExprFnOrIdxCall(c *ExprFnOrIdxCallContext)

	// ExitExprParenthesized is called when exiting the ExprParenthesized production.
	ExitExprParenthesized(c *ExprParenthesizedContext)

	// ExitExprMulDiv is called when exiting the ExprMulDiv production.
	ExitExprMulDiv(c *ExprMulDivContext)

	// ExitExprNegate is called when exiting the ExprNegate production.
	ExitExprNegate(c *ExprNegateContext)

	// ExitExprArrayIndex is called when exiting the ExprArrayIndex production.
	ExitExprArrayIndex(c *ExprArrayIndexContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprEquality is called when exiting the ExprEquality production.
	ExitExprEquality(c *ExprEqualityContext)

	// ExitExprDynamicReference is called when exiting the ExprDynamicReference production.
	ExitExprDynamicReference(c *ExprDynamicReferenceContext)

	// ExitExprId is called when exiting the ExprId production.
	ExitExprId(c *ExprIdContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVarDeclarator is called when exiting the varDeclarator production.
	ExitVarDeclarator(c *VarDeclaratorContext)

	// ExitVarType is called when exiting the varType production.
	ExitVarType(c *VarTypeContext)

	// ExitAppClassImport is called when exiting the appClassImport production.
	ExitAppClassImport(c *AppClassImportContext)

	// ExitAppPkgPath is called when exiting the appPkgPath production.
	ExitAppPkgPath(c *AppPkgPathContext)

	// ExitAppClassPath is called when exiting the appClassPath production.
	ExitAppClassPath(c *AppClassPathContext)

	// ExitExtFuncImport is called when exiting the extFuncImport production.
	ExitExtFuncImport(c *ExtFuncImportContext)

	// ExitRecDefnPath is called when exiting the recDefnPath production.
	ExitRecDefnPath(c *RecDefnPathContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBlock is called when exiting the classBlock production.
	ExitClassBlock(c *ClassBlockContext)

	// ExitClassBlockStmt is called when exiting the classBlockStmt production.
	ExitClassBlockStmt(c *ClassBlockStmtContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitInstance is called when exiting the instance production.
	ExitInstance(c *InstanceContext)

	// ExitMethodImpl is called when exiting the methodImpl production.
	ExitMethodImpl(c *MethodImplContext)

	// ExitGetImpl is called when exiting the getImpl production.
	ExitGetImpl(c *GetImplContext)

	// ExitSetImpl is called when exiting the setImpl production.
	ExitSetImpl(c *SetImplContext)

	// ExitFuncImpl is called when exiting the funcImpl production.
	ExitFuncImpl(c *FuncImplContext)

	// ExitFuncSignature is called when exiting the funcSignature production.
	ExitFuncSignature(c *FuncSignatureContext)

	// ExitFormalParamList is called when exiting the formalParamList production.
	ExitFormalParamList(c *FormalParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitEvaluateStmt is called when exiting the evaluateStmt production.
	ExitEvaluateStmt(c *EvaluateStmtContext)

	// ExitWhenBranch is called when exiting the whenBranch production.
	ExitWhenBranch(c *WhenBranchContext)

	// ExitWhenOtherBranch is called when exiting the whenOtherBranch production.
	ExitWhenOtherBranch(c *WhenOtherBranchContext)

	// ExitTryCatchStmt is called when exiting the tryCatchStmt production.
	ExitTryCatchStmt(c *TryCatchStmtContext)

	// ExitCatchSignature is called when exiting the catchSignature production.
	ExitCatchSignature(c *CatchSignatureContext)

	// ExitCreateInvocation is called when exiting the createInvocation production.
	ExitCreateInvocation(c *CreateInvocationContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}
