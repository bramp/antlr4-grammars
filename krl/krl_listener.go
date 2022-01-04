// Code generated from krl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package krl // krl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// krlListener is a complete listener for a parse tree produced by krlParser.
type krlListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterModuleRoutines is called when entering the moduleRoutines production.
	EnterModuleRoutines(c *ModuleRoutinesContext)

	// EnterMainRoutine is called when entering the mainRoutine production.
	EnterMainRoutine(c *MainRoutineContext)

	// EnterSubRoutine is called when entering the subRoutine production.
	EnterSubRoutine(c *SubRoutineContext)

	// EnterProcedureDefinition is called when entering the procedureDefinition production.
	EnterProcedureDefinition(c *ProcedureDefinitionContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterModuleData is called when entering the moduleData production.
	EnterModuleData(c *ModuleDataContext)

	// EnterModuleName is called when entering the moduleName production.
	EnterModuleName(c *ModuleNameContext)

	// EnterDataList is called when entering the dataList production.
	EnterDataList(c *DataListContext)

	// EnterArrayInitialisation is called when entering the arrayInitialisation production.
	EnterArrayInitialisation(c *ArrayInitialisationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterStructureDefinition is called when entering the structureDefinition production.
	EnterStructureDefinition(c *StructureDefinitionContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterSignalDeclaration is called when entering the signalDeclaration production.
	EnterSignalDeclaration(c *SignalDeclarationContext)

	// EnterVariableDeclarationInDataList is called when entering the variableDeclarationInDataList production.
	EnterVariableDeclarationInDataList(c *VariableDeclarationInDataListContext)

	// EnterVariableListRest is called when entering the variableListRest production.
	EnterVariableListRest(c *VariableListRestContext)

	// EnterVariableInitialisation is called when entering the variableInitialisation production.
	EnterVariableInitialisation(c *VariableInitialisationContext)

	// EnterStructLiteral is called when entering the structLiteral production.
	EnterStructLiteral(c *StructLiteralContext)

	// EnterStructElementList is called when entering the structElementList production.
	EnterStructElementList(c *StructElementListContext)

	// EnterStructElement is called when entering the structElement production.
	EnterStructElement(c *StructElementContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterRoutineBody is called when entering the routineBody production.
	EnterRoutineBody(c *RoutineBodyContext)

	// EnterRoutineDataSection is called when entering the routineDataSection production.
	EnterRoutineDataSection(c *RoutineDataSectionContext)

	// EnterForwardDeclaration is called when entering the forwardDeclaration production.
	EnterForwardDeclaration(c *ForwardDeclarationContext)

	// EnterFormalParametersWithType is called when entering the formalParametersWithType production.
	EnterFormalParametersWithType(c *FormalParametersWithTypeContext)

	// EnterParameterWithType is called when entering the parameterWithType production.
	EnterParameterWithType(c *ParameterWithTypeContext)

	// EnterParameterCallType is called when entering the parameterCallType production.
	EnterParameterCallType(c *ParameterCallTypeContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterArrayVariableSuffix is called when entering the arrayVariableSuffix production.
	EnterArrayVariableSuffix(c *ArrayVariableSuffixContext)

	// EnterRoutineImplementationSection is called when entering the routineImplementationSection production.
	EnterRoutineImplementationSection(c *RoutineImplementationSectionContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAnalogOutputStatement is called when entering the analogOutputStatement production.
	EnterAnalogOutputStatement(c *AnalogOutputStatementContext)

	// EnterAnalogInputStatement is called when entering the analogInputStatement production.
	EnterAnalogInputStatement(c *AnalogInputStatementContext)

	// EnterSwitchBlockStatementGroups is called when entering the switchBlockStatementGroups production.
	EnterSwitchBlockStatementGroups(c *SwitchBlockStatementGroupsContext)

	// EnterCaseLabel is called when entering the caseLabel production.
	EnterCaseLabel(c *CaseLabelContext)

	// EnterDefaultLabel is called when entering the defaultLabel production.
	EnterDefaultLabel(c *DefaultLabelContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelationalOp is called when entering the relationalOp production.
	EnterRelationalOp(c *RelationalOpContext)

	// EnterConditionalOrExpression is called when entering the conditionalOrExpression production.
	EnterConditionalOrExpression(c *ConditionalOrExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterConditionalAndExpression is called when entering the conditionalAndExpression production.
	EnterConditionalAndExpression(c *ConditionalAndExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterGeometricExpression is called when entering the geometricExpression production.
	EnterGeometricExpression(c *GeometricExpressionContext)

	// EnterUnaryNotExpression is called when entering the unaryNotExpression production.
	EnterUnaryNotExpression(c *UnaryNotExpressionContext)

	// EnterUnaryPlusMinuxExpression is called when entering the unaryPlusMinuxExpression production.
	EnterUnaryPlusMinuxExpression(c *UnaryPlusMinuxExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterEnumElement is called when entering the enumElement production.
	EnterEnumElement(c *EnumElementContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitModuleRoutines is called when exiting the moduleRoutines production.
	ExitModuleRoutines(c *ModuleRoutinesContext)

	// ExitMainRoutine is called when exiting the mainRoutine production.
	ExitMainRoutine(c *MainRoutineContext)

	// ExitSubRoutine is called when exiting the subRoutine production.
	ExitSubRoutine(c *SubRoutineContext)

	// ExitProcedureDefinition is called when exiting the procedureDefinition production.
	ExitProcedureDefinition(c *ProcedureDefinitionContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitModuleData is called when exiting the moduleData production.
	ExitModuleData(c *ModuleDataContext)

	// ExitModuleName is called when exiting the moduleName production.
	ExitModuleName(c *ModuleNameContext)

	// ExitDataList is called when exiting the dataList production.
	ExitDataList(c *DataListContext)

	// ExitArrayInitialisation is called when exiting the arrayInitialisation production.
	ExitArrayInitialisation(c *ArrayInitialisationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitStructureDefinition is called when exiting the structureDefinition production.
	ExitStructureDefinition(c *StructureDefinitionContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitSignalDeclaration is called when exiting the signalDeclaration production.
	ExitSignalDeclaration(c *SignalDeclarationContext)

	// ExitVariableDeclarationInDataList is called when exiting the variableDeclarationInDataList production.
	ExitVariableDeclarationInDataList(c *VariableDeclarationInDataListContext)

	// ExitVariableListRest is called when exiting the variableListRest production.
	ExitVariableListRest(c *VariableListRestContext)

	// ExitVariableInitialisation is called when exiting the variableInitialisation production.
	ExitVariableInitialisation(c *VariableInitialisationContext)

	// ExitStructLiteral is called when exiting the structLiteral production.
	ExitStructLiteral(c *StructLiteralContext)

	// ExitStructElementList is called when exiting the structElementList production.
	ExitStructElementList(c *StructElementListContext)

	// ExitStructElement is called when exiting the structElement production.
	ExitStructElement(c *StructElementContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitRoutineBody is called when exiting the routineBody production.
	ExitRoutineBody(c *RoutineBodyContext)

	// ExitRoutineDataSection is called when exiting the routineDataSection production.
	ExitRoutineDataSection(c *RoutineDataSectionContext)

	// ExitForwardDeclaration is called when exiting the forwardDeclaration production.
	ExitForwardDeclaration(c *ForwardDeclarationContext)

	// ExitFormalParametersWithType is called when exiting the formalParametersWithType production.
	ExitFormalParametersWithType(c *FormalParametersWithTypeContext)

	// ExitParameterWithType is called when exiting the parameterWithType production.
	ExitParameterWithType(c *ParameterWithTypeContext)

	// ExitParameterCallType is called when exiting the parameterCallType production.
	ExitParameterCallType(c *ParameterCallTypeContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitArrayVariableSuffix is called when exiting the arrayVariableSuffix production.
	ExitArrayVariableSuffix(c *ArrayVariableSuffixContext)

	// ExitRoutineImplementationSection is called when exiting the routineImplementationSection production.
	ExitRoutineImplementationSection(c *RoutineImplementationSectionContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAnalogOutputStatement is called when exiting the analogOutputStatement production.
	ExitAnalogOutputStatement(c *AnalogOutputStatementContext)

	// ExitAnalogInputStatement is called when exiting the analogInputStatement production.
	ExitAnalogInputStatement(c *AnalogInputStatementContext)

	// ExitSwitchBlockStatementGroups is called when exiting the switchBlockStatementGroups production.
	ExitSwitchBlockStatementGroups(c *SwitchBlockStatementGroupsContext)

	// ExitCaseLabel is called when exiting the caseLabel production.
	ExitCaseLabel(c *CaseLabelContext)

	// ExitDefaultLabel is called when exiting the defaultLabel production.
	ExitDefaultLabel(c *DefaultLabelContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelationalOp is called when exiting the relationalOp production.
	ExitRelationalOp(c *RelationalOpContext)

	// ExitConditionalOrExpression is called when exiting the conditionalOrExpression production.
	ExitConditionalOrExpression(c *ConditionalOrExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitConditionalAndExpression is called when exiting the conditionalAndExpression production.
	ExitConditionalAndExpression(c *ConditionalAndExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitGeometricExpression is called when exiting the geometricExpression production.
	ExitGeometricExpression(c *GeometricExpressionContext)

	// ExitUnaryNotExpression is called when exiting the unaryNotExpression production.
	ExitUnaryNotExpression(c *UnaryNotExpressionContext)

	// ExitUnaryPlusMinuxExpression is called when exiting the unaryPlusMinuxExpression production.
	ExitUnaryPlusMinuxExpression(c *UnaryPlusMinuxExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitEnumElement is called when exiting the enumElement production.
	ExitEnumElement(c *EnumElementContext)
}
