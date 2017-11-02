// Generated from ECMAScript.g4 by ANTLR 4.7.

package ecmascript // ECMAScript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseECMAScriptListener is a complete listener for a parse tree produced by ECMAScriptParser.
type BaseECMAScriptListener struct{}

var _ ECMAScriptListener = &BaseECMAScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseECMAScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseECMAScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseECMAScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseECMAScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseECMAScriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseECMAScriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseECMAScriptListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseECMAScriptListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseECMAScriptListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseECMAScriptListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseECMAScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseECMAScriptListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseECMAScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseECMAScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseECMAScriptListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseECMAScriptListener) ExitStatementList(ctx *StatementListContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseECMAScriptListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseECMAScriptListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseECMAScriptListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseECMAScriptListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterInitialiser is called when production initialiser is entered.
func (s *BaseECMAScriptListener) EnterInitialiser(ctx *InitialiserContext) {}

// ExitInitialiser is called when production initialiser is exited.
func (s *BaseECMAScriptListener) ExitInitialiser(ctx *InitialiserContext) {}

// EnterVoidStatement is called when production voidStatement is entered.
func (s *BaseECMAScriptListener) EnterVoidStatement(ctx *VoidStatementContext) {}

// ExitVoidStatement is called when production voidStatement is exited.
func (s *BaseECMAScriptListener) ExitVoidStatement(ctx *VoidStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseECMAScriptListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseECMAScriptListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseECMAScriptListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseECMAScriptListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterDoStatement is called when production DoStatement is entered.
func (s *BaseECMAScriptListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production DoStatement is exited.
func (s *BaseECMAScriptListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseECMAScriptListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseECMAScriptListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseECMAScriptListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseECMAScriptListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForVarStatement is called when production ForVarStatement is entered.
func (s *BaseECMAScriptListener) EnterForVarStatement(ctx *ForVarStatementContext) {}

// ExitForVarStatement is called when production ForVarStatement is exited.
func (s *BaseECMAScriptListener) ExitForVarStatement(ctx *ForVarStatementContext) {}

// EnterForInStatement is called when production ForInStatement is entered.
func (s *BaseECMAScriptListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production ForInStatement is exited.
func (s *BaseECMAScriptListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterForVarInStatement is called when production ForVarInStatement is entered.
func (s *BaseECMAScriptListener) EnterForVarInStatement(ctx *ForVarInStatementContext) {}

// ExitForVarInStatement is called when production ForVarInStatement is exited.
func (s *BaseECMAScriptListener) ExitForVarInStatement(ctx *ForVarInStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseECMAScriptListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseECMAScriptListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseECMAScriptListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseECMAScriptListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseECMAScriptListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseECMAScriptListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseECMAScriptListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseECMAScriptListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseECMAScriptListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseECMAScriptListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseECMAScriptListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseECMAScriptListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseECMAScriptListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseECMAScriptListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseECMAScriptListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseECMAScriptListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseECMAScriptListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseECMAScriptListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterLabelledStatement is called when production labelledStatement is entered.
func (s *BaseECMAScriptListener) EnterLabelledStatement(ctx *LabelledStatementContext) {}

// ExitLabelledStatement is called when production labelledStatement is exited.
func (s *BaseECMAScriptListener) ExitLabelledStatement(ctx *LabelledStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseECMAScriptListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseECMAScriptListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseECMAScriptListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseECMAScriptListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchProduction is called when production catchProduction is entered.
func (s *BaseECMAScriptListener) EnterCatchProduction(ctx *CatchProductionContext) {}

// ExitCatchProduction is called when production catchProduction is exited.
func (s *BaseECMAScriptListener) ExitCatchProduction(ctx *CatchProductionContext) {}

// EnterFinallyProduction is called when production finallyProduction is entered.
func (s *BaseECMAScriptListener) EnterFinallyProduction(ctx *FinallyProductionContext) {}

// ExitFinallyProduction is called when production finallyProduction is exited.
func (s *BaseECMAScriptListener) ExitFinallyProduction(ctx *FinallyProductionContext) {}

// EnterDebuggerStatement is called when production debuggerStatement is entered.
func (s *BaseECMAScriptListener) EnterDebuggerStatement(ctx *DebuggerStatementContext) {}

// ExitDebuggerStatement is called when production debuggerStatement is exited.
func (s *BaseECMAScriptListener) ExitDebuggerStatement(ctx *DebuggerStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseECMAScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseECMAScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseECMAScriptListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseECMAScriptListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseECMAScriptListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseECMAScriptListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseECMAScriptListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseECMAScriptListener) ExitElementList(ctx *ElementListContext) {}

// EnterElision is called when production elision is entered.
func (s *BaseECMAScriptListener) EnterElision(ctx *ElisionContext) {}

// ExitElision is called when production elision is exited.
func (s *BaseECMAScriptListener) ExitElision(ctx *ElisionContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropertyNameAndValueList is called when production propertyNameAndValueList is entered.
func (s *BaseECMAScriptListener) EnterPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) {}

// ExitPropertyNameAndValueList is called when production propertyNameAndValueList is exited.
func (s *BaseECMAScriptListener) ExitPropertyNameAndValueList(ctx *PropertyNameAndValueListContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseECMAScriptListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseECMAScriptListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterPropertyGetter is called when production PropertyGetter is entered.
func (s *BaseECMAScriptListener) EnterPropertyGetter(ctx *PropertyGetterContext) {}

// ExitPropertyGetter is called when production PropertyGetter is exited.
func (s *BaseECMAScriptListener) ExitPropertyGetter(ctx *PropertyGetterContext) {}

// EnterPropertySetter is called when production PropertySetter is entered.
func (s *BaseECMAScriptListener) EnterPropertySetter(ctx *PropertySetterContext) {}

// ExitPropertySetter is called when production PropertySetter is exited.
func (s *BaseECMAScriptListener) ExitPropertySetter(ctx *PropertySetterContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseECMAScriptListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseECMAScriptListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterPropertySetParameterList is called when production propertySetParameterList is entered.
func (s *BaseECMAScriptListener) EnterPropertySetParameterList(ctx *PropertySetParameterListContext) {}

// ExitPropertySetParameterList is called when production propertySetParameterList is exited.
func (s *BaseECMAScriptListener) ExitPropertySetParameterList(ctx *PropertySetParameterListContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseECMAScriptListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseECMAScriptListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseECMAScriptListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseECMAScriptListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseECMAScriptListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseECMAScriptListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseECMAScriptListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseECMAScriptListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseECMAScriptListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseECMAScriptListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterPreIncrementExpression is called when production PreIncrementExpression is entered.
func (s *BaseECMAScriptListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// ExitPreIncrementExpression is called when production PreIncrementExpression is exited.
func (s *BaseECMAScriptListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseECMAScriptListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseECMAScriptListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseECMAScriptListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseECMAScriptListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseECMAScriptListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseECMAScriptListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseECMAScriptListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseECMAScriptListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPreDecreaseExpression is called when production PreDecreaseExpression is entered.
func (s *BaseECMAScriptListener) EnterPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// ExitPreDecreaseExpression is called when production PreDecreaseExpression is exited.
func (s *BaseECMAScriptListener) ExitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseECMAScriptListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseECMAScriptListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterThisExpression is called when production ThisExpression is entered.
func (s *BaseECMAScriptListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production ThisExpression is exited.
func (s *BaseECMAScriptListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseECMAScriptListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseECMAScriptListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseECMAScriptListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseECMAScriptListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BaseECMAScriptListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BaseECMAScriptListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostDecreaseExpression is called when production PostDecreaseExpression is entered.
func (s *BaseECMAScriptListener) EnterPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {}

// ExitPostDecreaseExpression is called when production PostDecreaseExpression is exited.
func (s *BaseECMAScriptListener) ExitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {}

// EnterTypeofExpression is called when production TypeofExpression is entered.
func (s *BaseECMAScriptListener) EnterTypeofExpression(ctx *TypeofExpressionContext) {}

// ExitTypeofExpression is called when production TypeofExpression is exited.
func (s *BaseECMAScriptListener) ExitTypeofExpression(ctx *TypeofExpressionContext) {}

// EnterInstanceofExpression is called when production InstanceofExpression is entered.
func (s *BaseECMAScriptListener) EnterInstanceofExpression(ctx *InstanceofExpressionContext) {}

// ExitInstanceofExpression is called when production InstanceofExpression is exited.
func (s *BaseECMAScriptListener) ExitInstanceofExpression(ctx *InstanceofExpressionContext) {}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseECMAScriptListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseECMAScriptListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterDeleteExpression is called when production DeleteExpression is entered.
func (s *BaseECMAScriptListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production DeleteExpression is exited.
func (s *BaseECMAScriptListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseECMAScriptListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseECMAScriptListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterBitXOrExpression is called when production BitXOrExpression is entered.
func (s *BaseECMAScriptListener) EnterBitXOrExpression(ctx *BitXOrExpressionContext) {}

// ExitBitXOrExpression is called when production BitXOrExpression is exited.
func (s *BaseECMAScriptListener) ExitBitXOrExpression(ctx *BitXOrExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseECMAScriptListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseECMAScriptListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterBitShiftExpression is called when production BitShiftExpression is entered.
func (s *BaseECMAScriptListener) EnterBitShiftExpression(ctx *BitShiftExpressionContext) {}

// ExitBitShiftExpression is called when production BitShiftExpression is exited.
func (s *BaseECMAScriptListener) ExitBitShiftExpression(ctx *BitShiftExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseECMAScriptListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseECMAScriptListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseECMAScriptListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseECMAScriptListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseECMAScriptListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseECMAScriptListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterPostIncrementExpression is called when production PostIncrementExpression is entered.
func (s *BaseECMAScriptListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// ExitPostIncrementExpression is called when production PostIncrementExpression is exited.
func (s *BaseECMAScriptListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseECMAScriptListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseECMAScriptListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewExpression is called when production NewExpression is entered.
func (s *BaseECMAScriptListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production NewExpression is exited.
func (s *BaseECMAScriptListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseECMAScriptListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseECMAScriptListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseECMAScriptListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseECMAScriptListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// EnterMemberDotExpression is called when production MemberDotExpression is entered.
func (s *BaseECMAScriptListener) EnterMemberDotExpression(ctx *MemberDotExpressionContext) {}

// ExitMemberDotExpression is called when production MemberDotExpression is exited.
func (s *BaseECMAScriptListener) ExitMemberDotExpression(ctx *MemberDotExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseECMAScriptListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseECMAScriptListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseECMAScriptListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseECMAScriptListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterBitAndExpression is called when production BitAndExpression is entered.
func (s *BaseECMAScriptListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production BitAndExpression is exited.
func (s *BaseECMAScriptListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterBitOrExpression is called when production BitOrExpression is entered.
func (s *BaseECMAScriptListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production BitOrExpression is exited.
func (s *BaseECMAScriptListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterAssignmentOperatorExpression is called when production AssignmentOperatorExpression is entered.
func (s *BaseECMAScriptListener) EnterAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// ExitAssignmentOperatorExpression is called when production AssignmentOperatorExpression is exited.
func (s *BaseECMAScriptListener) ExitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// EnterVoidExpression is called when production VoidExpression is entered.
func (s *BaseECMAScriptListener) EnterVoidExpression(ctx *VoidExpressionContext) {}

// ExitVoidExpression is called when production VoidExpression is exited.
func (s *BaseECMAScriptListener) ExitVoidExpression(ctx *VoidExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseECMAScriptListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseECMAScriptListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseECMAScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseECMAScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseECMAScriptListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseECMAScriptListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseECMAScriptListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseECMAScriptListener) ExitIdentifierName(ctx *IdentifierNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseECMAScriptListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseECMAScriptListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseECMAScriptListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseECMAScriptListener) ExitKeyword(ctx *KeywordContext) {}

// EnterFutureReservedWord is called when production futureReservedWord is entered.
func (s *BaseECMAScriptListener) EnterFutureReservedWord(ctx *FutureReservedWordContext) {}

// ExitFutureReservedWord is called when production futureReservedWord is exited.
func (s *BaseECMAScriptListener) ExitFutureReservedWord(ctx *FutureReservedWordContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseECMAScriptListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseECMAScriptListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseECMAScriptListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseECMAScriptListener) ExitSetter(ctx *SetterContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseECMAScriptListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseECMAScriptListener) ExitEos(ctx *EosContext) {}

// EnterEof is called when production eof is entered.
func (s *BaseECMAScriptListener) EnterEof(ctx *EofContext) {}

// ExitEof is called when production eof is exited.
func (s *BaseECMAScriptListener) ExitEof(ctx *EofContext) {}
