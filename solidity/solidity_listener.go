// Code generated from Solidity.g4 by ANTLR 4.7.2. DO NOT EDIT.

package solidity // Solidity
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SolidityListener is a complete listener for a parse tree produced by SolidityParser.
type SolidityListener interface {
	antlr.ParseTreeListener

	// EnterSourceUnit is called when entering the sourceUnit production.
	EnterSourceUnit(c *SourceUnitContext)

	// EnterPragmaDirective is called when entering the pragmaDirective production.
	EnterPragmaDirective(c *PragmaDirectiveContext)

	// EnterPragmaName is called when entering the pragmaName production.
	EnterPragmaName(c *PragmaNameContext)

	// EnterPragmaValue is called when entering the pragmaValue production.
	EnterPragmaValue(c *PragmaValueContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterVersionOperator is called when entering the versionOperator production.
	EnterVersionOperator(c *VersionOperatorContext)

	// EnterVersionConstraint is called when entering the versionConstraint production.
	EnterVersionConstraint(c *VersionConstraintContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterImportDirective is called when entering the importDirective production.
	EnterImportDirective(c *ImportDirectiveContext)

	// EnterContractDefinition is called when entering the contractDefinition production.
	EnterContractDefinition(c *ContractDefinitionContext)

	// EnterInheritanceSpecifier is called when entering the inheritanceSpecifier production.
	EnterInheritanceSpecifier(c *InheritanceSpecifierContext)

	// EnterContractPart is called when entering the contractPart production.
	EnterContractPart(c *ContractPartContext)

	// EnterStateVariableDeclaration is called when entering the stateVariableDeclaration production.
	EnterStateVariableDeclaration(c *StateVariableDeclarationContext)

	// EnterUsingForDeclaration is called when entering the usingForDeclaration production.
	EnterUsingForDeclaration(c *UsingForDeclarationContext)

	// EnterStructDefinition is called when entering the structDefinition production.
	EnterStructDefinition(c *StructDefinitionContext)

	// EnterConstructorDefinition is called when entering the constructorDefinition production.
	EnterConstructorDefinition(c *ConstructorDefinitionContext)

	// EnterModifierDefinition is called when entering the modifierDefinition production.
	EnterModifierDefinition(c *ModifierDefinitionContext)

	// EnterModifierInvocation is called when entering the modifierInvocation production.
	EnterModifierInvocation(c *ModifierInvocationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterReturnParameters is called when entering the returnParameters production.
	EnterReturnParameters(c *ReturnParametersContext)

	// EnterModifierList is called when entering the modifierList production.
	EnterModifierList(c *ModifierListContext)

	// EnterEventDefinition is called when entering the eventDefinition production.
	EnterEventDefinition(c *EventDefinitionContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterEnumDefinition is called when entering the enumDefinition production.
	EnterEnumDefinition(c *EnumDefinitionContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterEventParameterList is called when entering the eventParameterList production.
	EnterEventParameterList(c *EventParameterListContext)

	// EnterEventParameter is called when entering the eventParameter production.
	EnterEventParameter(c *EventParameterContext)

	// EnterFunctionTypeParameterList is called when entering the functionTypeParameterList production.
	EnterFunctionTypeParameterList(c *FunctionTypeParameterListContext)

	// EnterFunctionTypeParameter is called when entering the functionTypeParameter production.
	EnterFunctionTypeParameter(c *FunctionTypeParameterContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterUserDefinedTypeName is called when entering the userDefinedTypeName production.
	EnterUserDefinedTypeName(c *UserDefinedTypeNameContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterFunctionTypeName is called when entering the functionTypeName production.
	EnterFunctionTypeName(c *FunctionTypeNameContext)

	// EnterStorageLocation is called when entering the storageLocation production.
	EnterStorageLocation(c *StorageLocationContext)

	// EnterStateMutability is called when entering the stateMutability production.
	EnterStateMutability(c *StateMutabilityContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterInlineAssemblyStatement is called when entering the inlineAssemblyStatement production.
	EnterInlineAssemblyStatement(c *InlineAssemblyStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterEmitStatement is called when entering the emitStatement production.
	EnterEmitStatement(c *EmitStatementContext)

	// EnterVariableDeclarationStatement is called when entering the variableDeclarationStatement production.
	EnterVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterElementaryTypeName is called when entering the elementaryTypeName production.
	EnterElementaryTypeName(c *ElementaryTypeNameContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterNameValueList is called when entering the nameValueList production.
	EnterNameValueList(c *NameValueListContext)

	// EnterNameValue is called when entering the nameValue production.
	EnterNameValue(c *NameValueContext)

	// EnterFunctionCallArguments is called when entering the functionCallArguments production.
	EnterFunctionCallArguments(c *FunctionCallArgumentsContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAssemblyBlock is called when entering the assemblyBlock production.
	EnterAssemblyBlock(c *AssemblyBlockContext)

	// EnterAssemblyItem is called when entering the assemblyItem production.
	EnterAssemblyItem(c *AssemblyItemContext)

	// EnterAssemblyExpression is called when entering the assemblyExpression production.
	EnterAssemblyExpression(c *AssemblyExpressionContext)

	// EnterAssemblyCall is called when entering the assemblyCall production.
	EnterAssemblyCall(c *AssemblyCallContext)

	// EnterAssemblyLocalDefinition is called when entering the assemblyLocalDefinition production.
	EnterAssemblyLocalDefinition(c *AssemblyLocalDefinitionContext)

	// EnterAssemblyAssignment is called when entering the assemblyAssignment production.
	EnterAssemblyAssignment(c *AssemblyAssignmentContext)

	// EnterAssemblyIdentifierOrList is called when entering the assemblyIdentifierOrList production.
	EnterAssemblyIdentifierOrList(c *AssemblyIdentifierOrListContext)

	// EnterAssemblyIdentifierList is called when entering the assemblyIdentifierList production.
	EnterAssemblyIdentifierList(c *AssemblyIdentifierListContext)

	// EnterAssemblyStackAssignment is called when entering the assemblyStackAssignment production.
	EnterAssemblyStackAssignment(c *AssemblyStackAssignmentContext)

	// EnterLabelDefinition is called when entering the labelDefinition production.
	EnterLabelDefinition(c *LabelDefinitionContext)

	// EnterAssemblySwitch is called when entering the assemblySwitch production.
	EnterAssemblySwitch(c *AssemblySwitchContext)

	// EnterAssemblyCase is called when entering the assemblyCase production.
	EnterAssemblyCase(c *AssemblyCaseContext)

	// EnterAssemblyFunctionDefinition is called when entering the assemblyFunctionDefinition production.
	EnterAssemblyFunctionDefinition(c *AssemblyFunctionDefinitionContext)

	// EnterAssemblyFunctionReturns is called when entering the assemblyFunctionReturns production.
	EnterAssemblyFunctionReturns(c *AssemblyFunctionReturnsContext)

	// EnterAssemblyFor is called when entering the assemblyFor production.
	EnterAssemblyFor(c *AssemblyForContext)

	// EnterAssemblyIf is called when entering the assemblyIf production.
	EnterAssemblyIf(c *AssemblyIfContext)

	// EnterAssemblyLiteral is called when entering the assemblyLiteral production.
	EnterAssemblyLiteral(c *AssemblyLiteralContext)

	// EnterSubAssembly is called when entering the subAssembly production.
	EnterSubAssembly(c *SubAssemblyContext)

	// EnterTupleExpression is called when entering the tupleExpression production.
	EnterTupleExpression(c *TupleExpressionContext)

	// EnterElementaryTypeNameExpression is called when entering the elementaryTypeNameExpression production.
	EnterElementaryTypeNameExpression(c *ElementaryTypeNameExpressionContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitSourceUnit is called when exiting the sourceUnit production.
	ExitSourceUnit(c *SourceUnitContext)

	// ExitPragmaDirective is called when exiting the pragmaDirective production.
	ExitPragmaDirective(c *PragmaDirectiveContext)

	// ExitPragmaName is called when exiting the pragmaName production.
	ExitPragmaName(c *PragmaNameContext)

	// ExitPragmaValue is called when exiting the pragmaValue production.
	ExitPragmaValue(c *PragmaValueContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitVersionOperator is called when exiting the versionOperator production.
	ExitVersionOperator(c *VersionOperatorContext)

	// ExitVersionConstraint is called when exiting the versionConstraint production.
	ExitVersionConstraint(c *VersionConstraintContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitImportDirective is called when exiting the importDirective production.
	ExitImportDirective(c *ImportDirectiveContext)

	// ExitContractDefinition is called when exiting the contractDefinition production.
	ExitContractDefinition(c *ContractDefinitionContext)

	// ExitInheritanceSpecifier is called when exiting the inheritanceSpecifier production.
	ExitInheritanceSpecifier(c *InheritanceSpecifierContext)

	// ExitContractPart is called when exiting the contractPart production.
	ExitContractPart(c *ContractPartContext)

	// ExitStateVariableDeclaration is called when exiting the stateVariableDeclaration production.
	ExitStateVariableDeclaration(c *StateVariableDeclarationContext)

	// ExitUsingForDeclaration is called when exiting the usingForDeclaration production.
	ExitUsingForDeclaration(c *UsingForDeclarationContext)

	// ExitStructDefinition is called when exiting the structDefinition production.
	ExitStructDefinition(c *StructDefinitionContext)

	// ExitConstructorDefinition is called when exiting the constructorDefinition production.
	ExitConstructorDefinition(c *ConstructorDefinitionContext)

	// ExitModifierDefinition is called when exiting the modifierDefinition production.
	ExitModifierDefinition(c *ModifierDefinitionContext)

	// ExitModifierInvocation is called when exiting the modifierInvocation production.
	ExitModifierInvocation(c *ModifierInvocationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitReturnParameters is called when exiting the returnParameters production.
	ExitReturnParameters(c *ReturnParametersContext)

	// ExitModifierList is called when exiting the modifierList production.
	ExitModifierList(c *ModifierListContext)

	// ExitEventDefinition is called when exiting the eventDefinition production.
	ExitEventDefinition(c *EventDefinitionContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitEnumDefinition is called when exiting the enumDefinition production.
	ExitEnumDefinition(c *EnumDefinitionContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitEventParameterList is called when exiting the eventParameterList production.
	ExitEventParameterList(c *EventParameterListContext)

	// ExitEventParameter is called when exiting the eventParameter production.
	ExitEventParameter(c *EventParameterContext)

	// ExitFunctionTypeParameterList is called when exiting the functionTypeParameterList production.
	ExitFunctionTypeParameterList(c *FunctionTypeParameterListContext)

	// ExitFunctionTypeParameter is called when exiting the functionTypeParameter production.
	ExitFunctionTypeParameter(c *FunctionTypeParameterContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitUserDefinedTypeName is called when exiting the userDefinedTypeName production.
	ExitUserDefinedTypeName(c *UserDefinedTypeNameContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitFunctionTypeName is called when exiting the functionTypeName production.
	ExitFunctionTypeName(c *FunctionTypeNameContext)

	// ExitStorageLocation is called when exiting the storageLocation production.
	ExitStorageLocation(c *StorageLocationContext)

	// ExitStateMutability is called when exiting the stateMutability production.
	ExitStateMutability(c *StateMutabilityContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitInlineAssemblyStatement is called when exiting the inlineAssemblyStatement production.
	ExitInlineAssemblyStatement(c *InlineAssemblyStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitEmitStatement is called when exiting the emitStatement production.
	ExitEmitStatement(c *EmitStatementContext)

	// ExitVariableDeclarationStatement is called when exiting the variableDeclarationStatement production.
	ExitVariableDeclarationStatement(c *VariableDeclarationStatementContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitElementaryTypeName is called when exiting the elementaryTypeName production.
	ExitElementaryTypeName(c *ElementaryTypeNameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitNameValueList is called when exiting the nameValueList production.
	ExitNameValueList(c *NameValueListContext)

	// ExitNameValue is called when exiting the nameValue production.
	ExitNameValue(c *NameValueContext)

	// ExitFunctionCallArguments is called when exiting the functionCallArguments production.
	ExitFunctionCallArguments(c *FunctionCallArgumentsContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAssemblyBlock is called when exiting the assemblyBlock production.
	ExitAssemblyBlock(c *AssemblyBlockContext)

	// ExitAssemblyItem is called when exiting the assemblyItem production.
	ExitAssemblyItem(c *AssemblyItemContext)

	// ExitAssemblyExpression is called when exiting the assemblyExpression production.
	ExitAssemblyExpression(c *AssemblyExpressionContext)

	// ExitAssemblyCall is called when exiting the assemblyCall production.
	ExitAssemblyCall(c *AssemblyCallContext)

	// ExitAssemblyLocalDefinition is called when exiting the assemblyLocalDefinition production.
	ExitAssemblyLocalDefinition(c *AssemblyLocalDefinitionContext)

	// ExitAssemblyAssignment is called when exiting the assemblyAssignment production.
	ExitAssemblyAssignment(c *AssemblyAssignmentContext)

	// ExitAssemblyIdentifierOrList is called when exiting the assemblyIdentifierOrList production.
	ExitAssemblyIdentifierOrList(c *AssemblyIdentifierOrListContext)

	// ExitAssemblyIdentifierList is called when exiting the assemblyIdentifierList production.
	ExitAssemblyIdentifierList(c *AssemblyIdentifierListContext)

	// ExitAssemblyStackAssignment is called when exiting the assemblyStackAssignment production.
	ExitAssemblyStackAssignment(c *AssemblyStackAssignmentContext)

	// ExitLabelDefinition is called when exiting the labelDefinition production.
	ExitLabelDefinition(c *LabelDefinitionContext)

	// ExitAssemblySwitch is called when exiting the assemblySwitch production.
	ExitAssemblySwitch(c *AssemblySwitchContext)

	// ExitAssemblyCase is called when exiting the assemblyCase production.
	ExitAssemblyCase(c *AssemblyCaseContext)

	// ExitAssemblyFunctionDefinition is called when exiting the assemblyFunctionDefinition production.
	ExitAssemblyFunctionDefinition(c *AssemblyFunctionDefinitionContext)

	// ExitAssemblyFunctionReturns is called when exiting the assemblyFunctionReturns production.
	ExitAssemblyFunctionReturns(c *AssemblyFunctionReturnsContext)

	// ExitAssemblyFor is called when exiting the assemblyFor production.
	ExitAssemblyFor(c *AssemblyForContext)

	// ExitAssemblyIf is called when exiting the assemblyIf production.
	ExitAssemblyIf(c *AssemblyIfContext)

	// ExitAssemblyLiteral is called when exiting the assemblyLiteral production.
	ExitAssemblyLiteral(c *AssemblyLiteralContext)

	// ExitSubAssembly is called when exiting the subAssembly production.
	ExitSubAssembly(c *SubAssemblyContext)

	// ExitTupleExpression is called when exiting the tupleExpression production.
	ExitTupleExpression(c *TupleExpressionContext)

	// ExitElementaryTypeNameExpression is called when exiting the elementaryTypeNameExpression production.
	ExitElementaryTypeNameExpression(c *ElementaryTypeNameExpressionContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
