// Code generated from krl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package krl // krl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasekrlListener is a complete listener for a parse tree produced by krlParser.
type BasekrlListener struct{}

var _ krlListener = &BasekrlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekrlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekrlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekrlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekrlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BasekrlListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BasekrlListener) ExitModule(ctx *ModuleContext) {}

// EnterModuleRoutines is called when production moduleRoutines is entered.
func (s *BasekrlListener) EnterModuleRoutines(ctx *ModuleRoutinesContext) {}

// ExitModuleRoutines is called when production moduleRoutines is exited.
func (s *BasekrlListener) ExitModuleRoutines(ctx *ModuleRoutinesContext) {}

// EnterMainRoutine is called when production mainRoutine is entered.
func (s *BasekrlListener) EnterMainRoutine(ctx *MainRoutineContext) {}

// ExitMainRoutine is called when production mainRoutine is exited.
func (s *BasekrlListener) ExitMainRoutine(ctx *MainRoutineContext) {}

// EnterSubRoutine is called when production subRoutine is entered.
func (s *BasekrlListener) EnterSubRoutine(ctx *SubRoutineContext) {}

// ExitSubRoutine is called when production subRoutine is exited.
func (s *BasekrlListener) ExitSubRoutine(ctx *SubRoutineContext) {}

// EnterProcedureDefinition is called when production procedureDefinition is entered.
func (s *BasekrlListener) EnterProcedureDefinition(ctx *ProcedureDefinitionContext) {}

// ExitProcedureDefinition is called when production procedureDefinition is exited.
func (s *BasekrlListener) ExitProcedureDefinition(ctx *ProcedureDefinitionContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BasekrlListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BasekrlListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BasekrlListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BasekrlListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BasekrlListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BasekrlListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterModuleData is called when production moduleData is entered.
func (s *BasekrlListener) EnterModuleData(ctx *ModuleDataContext) {}

// ExitModuleData is called when production moduleData is exited.
func (s *BasekrlListener) ExitModuleData(ctx *ModuleDataContext) {}

// EnterModuleName is called when production moduleName is entered.
func (s *BasekrlListener) EnterModuleName(ctx *ModuleNameContext) {}

// ExitModuleName is called when production moduleName is exited.
func (s *BasekrlListener) ExitModuleName(ctx *ModuleNameContext) {}

// EnterDataList is called when production dataList is entered.
func (s *BasekrlListener) EnterDataList(ctx *DataListContext) {}

// ExitDataList is called when production dataList is exited.
func (s *BasekrlListener) ExitDataList(ctx *DataListContext) {}

// EnterArrayInitialisation is called when production arrayInitialisation is entered.
func (s *BasekrlListener) EnterArrayInitialisation(ctx *ArrayInitialisationContext) {}

// ExitArrayInitialisation is called when production arrayInitialisation is exited.
func (s *BasekrlListener) ExitArrayInitialisation(ctx *ArrayInitialisationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BasekrlListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BasekrlListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterStructureDefinition is called when production structureDefinition is entered.
func (s *BasekrlListener) EnterStructureDefinition(ctx *StructureDefinitionContext) {}

// ExitStructureDefinition is called when production structureDefinition is exited.
func (s *BasekrlListener) ExitStructureDefinition(ctx *StructureDefinitionContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BasekrlListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BasekrlListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BasekrlListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BasekrlListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BasekrlListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BasekrlListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterSignalDeclaration is called when production signalDeclaration is entered.
func (s *BasekrlListener) EnterSignalDeclaration(ctx *SignalDeclarationContext) {}

// ExitSignalDeclaration is called when production signalDeclaration is exited.
func (s *BasekrlListener) ExitSignalDeclaration(ctx *SignalDeclarationContext) {}

// EnterVariableDeclarationInDataList is called when production variableDeclarationInDataList is entered.
func (s *BasekrlListener) EnterVariableDeclarationInDataList(ctx *VariableDeclarationInDataListContext) {
}

// ExitVariableDeclarationInDataList is called when production variableDeclarationInDataList is exited.
func (s *BasekrlListener) ExitVariableDeclarationInDataList(ctx *VariableDeclarationInDataListContext) {
}

// EnterVariableListRest is called when production variableListRest is entered.
func (s *BasekrlListener) EnterVariableListRest(ctx *VariableListRestContext) {}

// ExitVariableListRest is called when production variableListRest is exited.
func (s *BasekrlListener) ExitVariableListRest(ctx *VariableListRestContext) {}

// EnterVariableInitialisation is called when production variableInitialisation is entered.
func (s *BasekrlListener) EnterVariableInitialisation(ctx *VariableInitialisationContext) {}

// ExitVariableInitialisation is called when production variableInitialisation is exited.
func (s *BasekrlListener) ExitVariableInitialisation(ctx *VariableInitialisationContext) {}

// EnterStructLiteral is called when production structLiteral is entered.
func (s *BasekrlListener) EnterStructLiteral(ctx *StructLiteralContext) {}

// ExitStructLiteral is called when production structLiteral is exited.
func (s *BasekrlListener) ExitStructLiteral(ctx *StructLiteralContext) {}

// EnterStructElementList is called when production structElementList is entered.
func (s *BasekrlListener) EnterStructElementList(ctx *StructElementListContext) {}

// ExitStructElementList is called when production structElementList is exited.
func (s *BasekrlListener) ExitStructElementList(ctx *StructElementListContext) {}

// EnterStructElement is called when production structElement is entered.
func (s *BasekrlListener) EnterStructElement(ctx *StructElementContext) {}

// ExitStructElement is called when production structElement is exited.
func (s *BasekrlListener) ExitStructElement(ctx *StructElementContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BasekrlListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BasekrlListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasekrlListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasekrlListener) ExitParameter(ctx *ParameterContext) {}

// EnterRoutineBody is called when production routineBody is entered.
func (s *BasekrlListener) EnterRoutineBody(ctx *RoutineBodyContext) {}

// ExitRoutineBody is called when production routineBody is exited.
func (s *BasekrlListener) ExitRoutineBody(ctx *RoutineBodyContext) {}

// EnterRoutineDataSection is called when production routineDataSection is entered.
func (s *BasekrlListener) EnterRoutineDataSection(ctx *RoutineDataSectionContext) {}

// ExitRoutineDataSection is called when production routineDataSection is exited.
func (s *BasekrlListener) ExitRoutineDataSection(ctx *RoutineDataSectionContext) {}

// EnterForwardDeclaration is called when production forwardDeclaration is entered.
func (s *BasekrlListener) EnterForwardDeclaration(ctx *ForwardDeclarationContext) {}

// ExitForwardDeclaration is called when production forwardDeclaration is exited.
func (s *BasekrlListener) ExitForwardDeclaration(ctx *ForwardDeclarationContext) {}

// EnterFormalParametersWithType is called when production formalParametersWithType is entered.
func (s *BasekrlListener) EnterFormalParametersWithType(ctx *FormalParametersWithTypeContext) {}

// ExitFormalParametersWithType is called when production formalParametersWithType is exited.
func (s *BasekrlListener) ExitFormalParametersWithType(ctx *FormalParametersWithTypeContext) {}

// EnterParameterWithType is called when production parameterWithType is entered.
func (s *BasekrlListener) EnterParameterWithType(ctx *ParameterWithTypeContext) {}

// ExitParameterWithType is called when production parameterWithType is exited.
func (s *BasekrlListener) ExitParameterWithType(ctx *ParameterWithTypeContext) {}

// EnterParameterCallType is called when production parameterCallType is entered.
func (s *BasekrlListener) EnterParameterCallType(ctx *ParameterCallTypeContext) {}

// ExitParameterCallType is called when production parameterCallType is exited.
func (s *BasekrlListener) ExitParameterCallType(ctx *ParameterCallTypeContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BasekrlListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BasekrlListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BasekrlListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BasekrlListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterArrayVariableSuffix is called when production arrayVariableSuffix is entered.
func (s *BasekrlListener) EnterArrayVariableSuffix(ctx *ArrayVariableSuffixContext) {}

// ExitArrayVariableSuffix is called when production arrayVariableSuffix is exited.
func (s *BasekrlListener) ExitArrayVariableSuffix(ctx *ArrayVariableSuffixContext) {}

// EnterRoutineImplementationSection is called when production routineImplementationSection is entered.
func (s *BasekrlListener) EnterRoutineImplementationSection(ctx *RoutineImplementationSectionContext) {
}

// ExitRoutineImplementationSection is called when production routineImplementationSection is exited.
func (s *BasekrlListener) ExitRoutineImplementationSection(ctx *RoutineImplementationSectionContext) {
}

// EnterStatementList is called when production statementList is entered.
func (s *BasekrlListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BasekrlListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasekrlListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasekrlListener) ExitStatement(ctx *StatementContext) {}

// EnterAnalogOutputStatement is called when production analogOutputStatement is entered.
func (s *BasekrlListener) EnterAnalogOutputStatement(ctx *AnalogOutputStatementContext) {}

// ExitAnalogOutputStatement is called when production analogOutputStatement is exited.
func (s *BasekrlListener) ExitAnalogOutputStatement(ctx *AnalogOutputStatementContext) {}

// EnterAnalogInputStatement is called when production analogInputStatement is entered.
func (s *BasekrlListener) EnterAnalogInputStatement(ctx *AnalogInputStatementContext) {}

// ExitAnalogInputStatement is called when production analogInputStatement is exited.
func (s *BasekrlListener) ExitAnalogInputStatement(ctx *AnalogInputStatementContext) {}

// EnterSwitchBlockStatementGroups is called when production switchBlockStatementGroups is entered.
func (s *BasekrlListener) EnterSwitchBlockStatementGroups(ctx *SwitchBlockStatementGroupsContext) {}

// ExitSwitchBlockStatementGroups is called when production switchBlockStatementGroups is exited.
func (s *BasekrlListener) ExitSwitchBlockStatementGroups(ctx *SwitchBlockStatementGroupsContext) {}

// EnterCaseLabel is called when production caseLabel is entered.
func (s *BasekrlListener) EnterCaseLabel(ctx *CaseLabelContext) {}

// ExitCaseLabel is called when production caseLabel is exited.
func (s *BasekrlListener) ExitCaseLabel(ctx *CaseLabelContext) {}

// EnterDefaultLabel is called when production defaultLabel is entered.
func (s *BasekrlListener) EnterDefaultLabel(ctx *DefaultLabelContext) {}

// ExitDefaultLabel is called when production defaultLabel is exited.
func (s *BasekrlListener) ExitDefaultLabel(ctx *DefaultLabelContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasekrlListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasekrlListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BasekrlListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BasekrlListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasekrlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasekrlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelationalOp is called when production relationalOp is entered.
func (s *BasekrlListener) EnterRelationalOp(ctx *RelationalOpContext) {}

// ExitRelationalOp is called when production relationalOp is exited.
func (s *BasekrlListener) ExitRelationalOp(ctx *RelationalOpContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BasekrlListener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BasekrlListener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BasekrlListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BasekrlListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BasekrlListener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BasekrlListener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BasekrlListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BasekrlListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BasekrlListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BasekrlListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterGeometricExpression is called when production geometricExpression is entered.
func (s *BasekrlListener) EnterGeometricExpression(ctx *GeometricExpressionContext) {}

// ExitGeometricExpression is called when production geometricExpression is exited.
func (s *BasekrlListener) ExitGeometricExpression(ctx *GeometricExpressionContext) {}

// EnterUnaryNotExpression is called when production unaryNotExpression is entered.
func (s *BasekrlListener) EnterUnaryNotExpression(ctx *UnaryNotExpressionContext) {}

// ExitUnaryNotExpression is called when production unaryNotExpression is exited.
func (s *BasekrlListener) ExitUnaryNotExpression(ctx *UnaryNotExpressionContext) {}

// EnterUnaryPlusMinuxExpression is called when production unaryPlusMinuxExpression is entered.
func (s *BasekrlListener) EnterUnaryPlusMinuxExpression(ctx *UnaryPlusMinuxExpressionContext) {}

// ExitUnaryPlusMinuxExpression is called when production unaryPlusMinuxExpression is exited.
func (s *BasekrlListener) ExitUnaryPlusMinuxExpression(ctx *UnaryPlusMinuxExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasekrlListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasekrlListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BasekrlListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BasekrlListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasekrlListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasekrlListener) ExitType_(ctx *Type_Context) {}

// EnterTypeName is called when production typeName is entered.
func (s *BasekrlListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BasekrlListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BasekrlListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BasekrlListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasekrlListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasekrlListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasekrlListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasekrlListener) ExitLiteral(ctx *LiteralContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BasekrlListener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BasekrlListener) ExitEnumElement(ctx *EnumElementContext) {}
