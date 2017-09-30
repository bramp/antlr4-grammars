// Generated from ECMAScript.g4 by ANTLR 4.7.

package ecmascript // ECMAScript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ECMAScriptParser.
type ECMAScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ECMAScriptParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#sourceElements.
	VisitSourceElements(ctx *SourceElementsContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#sourceElement.
	VisitSourceElement(ctx *SourceElementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclarationList.
	VisitVariableDeclarationList(ctx *VariableDeclarationListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#initialiser.
	VisitInitialiser(ctx *InitialiserContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#DoStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#WhileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ForVarStatement.
	VisitForVarStatement(ctx *ForVarStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ForInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ForVarInStatement.
	VisitForVarInStatement(ctx *ForVarInStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClauses.
	VisitCaseClauses(ctx *CaseClausesContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#caseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#defaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#labelledStatement.
	VisitLabelledStatement(ctx *LabelledStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#catchProduction.
	VisitCatchProduction(ctx *CatchProductionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#finallyProduction.
	VisitFinallyProduction(ctx *FinallyProductionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#debuggerStatement.
	VisitDebuggerStatement(ctx *DebuggerStatementContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#elision.
	VisitElision(ctx *ElisionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyNameAndValueList.
	VisitPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PropertyExpressionAssignment.
	VisitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PropertyGetter.
	VisitPropertyGetter(ctx *PropertyGetterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PropertySetter.
	VisitPropertySetter(ctx *PropertySetterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#propertySetParameterList.
	VisitPropertySetParameterList(ctx *PropertySetParameterListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#expressionSequence.
	VisitExpressionSequence(ctx *ExpressionSequenceContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#TernaryExpression.
	VisitTernaryExpression(ctx *TernaryExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#LogicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PreIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ObjectLiteralExpression.
	VisitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#InExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#LogicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#NotExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PreDecreaseExpression.
	VisitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ArgumentsExpression.
	VisitArgumentsExpression(ctx *ArgumentsExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ThisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#UnaryMinusExpression.
	VisitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#AssignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PostDecreaseExpression.
	VisitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#TypeofExpression.
	VisitTypeofExpression(ctx *TypeofExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#InstanceofExpression.
	VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#UnaryPlusExpression.
	VisitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#DeleteExpression.
	VisitDeleteExpression(ctx *DeleteExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#EqualityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#BitXOrExpression.
	VisitBitXOrExpression(ctx *BitXOrExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#MultiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#BitShiftExpression.
	VisitBitShiftExpression(ctx *BitShiftExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#AdditiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#RelationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#PostIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#BitNotExpression.
	VisitBitNotExpression(ctx *BitNotExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#NewExpression.
	VisitNewExpression(ctx *NewExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#ArrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#MemberDotExpression.
	VisitMemberDotExpression(ctx *MemberDotExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#MemberIndexExpression.
	VisitMemberIndexExpression(ctx *MemberIndexExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#BitAndExpression.
	VisitBitAndExpression(ctx *BitAndExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#BitOrExpression.
	VisitBitOrExpression(ctx *BitOrExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#AssignmentOperatorExpression.
	VisitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#VoidExpression.
	VisitVoidExpression(ctx *VoidExpressionContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#identifierName.
	VisitIdentifierName(ctx *IdentifierNameContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#futureReservedWord.
	VisitFutureReservedWord(ctx *FutureReservedWordContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by ECMAScriptParser#eof.
	VisitEof(ctx *EofContext) interface{}
}
