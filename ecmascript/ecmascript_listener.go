// Generated from ECMAScript.g4 by ANTLR 4.7.

package ecmascript // ECMAScript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ECMAScriptListener is a complete listener for a parse tree produced by ECMAScriptParser.
type ECMAScriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSourceElements is called when entering the sourceElements production.
	EnterSourceElements(c *SourceElementsContext)

	// EnterSourceElement is called when entering the sourceElement production.
	EnterSourceElement(c *SourceElementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterInitialiser is called when entering the initialiser production.
	EnterInitialiser(c *InitialiserContext)

	// EnterVoidStatement is called when entering the voidStatement production.
	EnterVoidStatement(c *VoidStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterDoStatement is called when entering the DoStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterWhileStatement is called when entering the WhileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the ForStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForVarStatement is called when entering the ForVarStatement production.
	EnterForVarStatement(c *ForVarStatementContext)

	// EnterForInStatement is called when entering the ForInStatement production.
	EnterForInStatement(c *ForInStatementContext)

	// EnterForVarInStatement is called when entering the ForVarInStatement production.
	EnterForVarInStatement(c *ForVarInStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterWithStatement is called when entering the withStatement production.
	EnterWithStatement(c *WithStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterCaseClauses is called when entering the caseClauses production.
	EnterCaseClauses(c *CaseClausesContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterDefaultClause is called when entering the defaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterLabelledStatement is called when entering the labelledStatement production.
	EnterLabelledStatement(c *LabelledStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterCatchProduction is called when entering the catchProduction production.
	EnterCatchProduction(c *CatchProductionContext)

	// EnterFinallyProduction is called when entering the finallyProduction production.
	EnterFinallyProduction(c *FinallyProductionContext)

	// EnterDebuggerStatement is called when entering the debuggerStatement production.
	EnterDebuggerStatement(c *DebuggerStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterElision is called when entering the elision production.
	EnterElision(c *ElisionContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterPropertyNameAndValueList is called when entering the propertyNameAndValueList production.
	EnterPropertyNameAndValueList(c *PropertyNameAndValueListContext)

	// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
	EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// EnterPropertyGetter is called when entering the PropertyGetter production.
	EnterPropertyGetter(c *PropertyGetterContext)

	// EnterPropertySetter is called when entering the PropertySetter production.
	EnterPropertySetter(c *PropertySetterContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterPropertySetParameterList is called when entering the propertySetParameterList production.
	EnterPropertySetParameterList(c *PropertySetParameterListContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
	EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
	EnterPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterThisExpression is called when entering the ThisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterAssignmentExpression is called when entering the AssignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
	EnterPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// EnterTypeofExpression is called when entering the TypeofExpression production.
	EnterTypeofExpression(c *TypeofExpressionContext)

	// EnterInstanceofExpression is called when entering the InstanceofExpression production.
	EnterInstanceofExpression(c *InstanceofExpressionContext)

	// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
	EnterUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// EnterDeleteExpression is called when entering the DeleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterBitXOrExpression is called when entering the BitXOrExpression production.
	EnterBitXOrExpression(c *BitXOrExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterBitShiftExpression is called when entering the BitShiftExpression production.
	EnterBitShiftExpression(c *BitShiftExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the RelationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
	EnterPostIncrementExpression(c *PostIncrementExpressionContext)

	// EnterBitNotExpression is called when entering the BitNotExpression production.
	EnterBitNotExpression(c *BitNotExpressionContext)

	// EnterNewExpression is called when entering the NewExpression production.
	EnterNewExpression(c *NewExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
	EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// EnterMemberDotExpression is called when entering the MemberDotExpression production.
	EnterMemberDotExpression(c *MemberDotExpressionContext)

	// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
	EnterMemberIndexExpression(c *MemberIndexExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterBitAndExpression is called when entering the BitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterBitOrExpression is called when entering the BitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
	EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// EnterVoidExpression is called when entering the VoidExpression production.
	EnterVoidExpression(c *VoidExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterIdentifierName is called when entering the identifierName production.
	EnterIdentifierName(c *IdentifierNameContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterFutureReservedWord is called when entering the futureReservedWord production.
	EnterFutureReservedWord(c *FutureReservedWordContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// EnterEof is called when entering the eof production.
	EnterEof(c *EofContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSourceElements is called when exiting the sourceElements production.
	ExitSourceElements(c *SourceElementsContext)

	// ExitSourceElement is called when exiting the sourceElement production.
	ExitSourceElement(c *SourceElementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitInitialiser is called when exiting the initialiser production.
	ExitInitialiser(c *InitialiserContext)

	// ExitVoidStatement is called when exiting the voidStatement production.
	ExitVoidStatement(c *VoidStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitDoStatement is called when exiting the DoStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitWhileStatement is called when exiting the WhileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the ForStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForVarStatement is called when exiting the ForVarStatement production.
	ExitForVarStatement(c *ForVarStatementContext)

	// ExitForInStatement is called when exiting the ForInStatement production.
	ExitForInStatement(c *ForInStatementContext)

	// ExitForVarInStatement is called when exiting the ForVarInStatement production.
	ExitForVarInStatement(c *ForVarInStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitWithStatement is called when exiting the withStatement production.
	ExitWithStatement(c *WithStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitCaseClauses is called when exiting the caseClauses production.
	ExitCaseClauses(c *CaseClausesContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitDefaultClause is called when exiting the defaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitLabelledStatement is called when exiting the labelledStatement production.
	ExitLabelledStatement(c *LabelledStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitCatchProduction is called when exiting the catchProduction production.
	ExitCatchProduction(c *CatchProductionContext)

	// ExitFinallyProduction is called when exiting the finallyProduction production.
	ExitFinallyProduction(c *FinallyProductionContext)

	// ExitDebuggerStatement is called when exiting the debuggerStatement production.
	ExitDebuggerStatement(c *DebuggerStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitElision is called when exiting the elision production.
	ExitElision(c *ElisionContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitPropertyNameAndValueList is called when exiting the propertyNameAndValueList production.
	ExitPropertyNameAndValueList(c *PropertyNameAndValueListContext)

	// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
	ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// ExitPropertyGetter is called when exiting the PropertyGetter production.
	ExitPropertyGetter(c *PropertyGetterContext)

	// ExitPropertySetter is called when exiting the PropertySetter production.
	ExitPropertySetter(c *PropertySetterContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitPropertySetParameterList is called when exiting the propertySetParameterList production.
	ExitPropertySetParameterList(c *PropertySetParameterListContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
	ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
	ExitPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitThisExpression is called when exiting the ThisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
	ExitPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// ExitTypeofExpression is called when exiting the TypeofExpression production.
	ExitTypeofExpression(c *TypeofExpressionContext)

	// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
	ExitInstanceofExpression(c *InstanceofExpressionContext)

	// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
	ExitUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// ExitDeleteExpression is called when exiting the DeleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
	ExitBitXOrExpression(c *BitXOrExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
	ExitBitShiftExpression(c *BitShiftExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the RelationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
	ExitPostIncrementExpression(c *PostIncrementExpressionContext)

	// ExitBitNotExpression is called when exiting the BitNotExpression production.
	ExitBitNotExpression(c *BitNotExpressionContext)

	// ExitNewExpression is called when exiting the NewExpression production.
	ExitNewExpression(c *NewExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
	ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
	ExitMemberDotExpression(c *MemberDotExpressionContext)

	// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
	ExitMemberIndexExpression(c *MemberIndexExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitBitAndExpression is called when exiting the BitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitBitOrExpression is called when exiting the BitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
	ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// ExitVoidExpression is called when exiting the VoidExpression production.
	ExitVoidExpression(c *VoidExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitIdentifierName is called when exiting the identifierName production.
	ExitIdentifierName(c *IdentifierNameContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitFutureReservedWord is called when exiting the futureReservedWord production.
	ExitFutureReservedWord(c *FutureReservedWordContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)

	// ExitEof is called when exiting the eof production.
	ExitEof(c *EofContext)
}
