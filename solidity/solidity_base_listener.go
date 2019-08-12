// Code generated from Solidity.g4 by ANTLR 4.7.2. DO NOT EDIT.

package solidity // Solidity
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSolidityListener is a complete listener for a parse tree produced by SolidityParser.
type BaseSolidityListener struct{}

var _ SolidityListener = &BaseSolidityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSolidityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSolidityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSolidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSolidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceUnit is called when production sourceUnit is entered.
func (s *BaseSolidityListener) EnterSourceUnit(ctx *SourceUnitContext) {}

// ExitSourceUnit is called when production sourceUnit is exited.
func (s *BaseSolidityListener) ExitSourceUnit(ctx *SourceUnitContext) {}

// EnterPragmaDirective is called when production pragmaDirective is entered.
func (s *BaseSolidityListener) EnterPragmaDirective(ctx *PragmaDirectiveContext) {}

// ExitPragmaDirective is called when production pragmaDirective is exited.
func (s *BaseSolidityListener) ExitPragmaDirective(ctx *PragmaDirectiveContext) {}

// EnterPragmaName is called when production pragmaName is entered.
func (s *BaseSolidityListener) EnterPragmaName(ctx *PragmaNameContext) {}

// ExitPragmaName is called when production pragmaName is exited.
func (s *BaseSolidityListener) ExitPragmaName(ctx *PragmaNameContext) {}

// EnterPragmaValue is called when production pragmaValue is entered.
func (s *BaseSolidityListener) EnterPragmaValue(ctx *PragmaValueContext) {}

// ExitPragmaValue is called when production pragmaValue is exited.
func (s *BaseSolidityListener) ExitPragmaValue(ctx *PragmaValueContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseSolidityListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseSolidityListener) ExitVersion(ctx *VersionContext) {}

// EnterVersionOperator is called when production versionOperator is entered.
func (s *BaseSolidityListener) EnterVersionOperator(ctx *VersionOperatorContext) {}

// ExitVersionOperator is called when production versionOperator is exited.
func (s *BaseSolidityListener) ExitVersionOperator(ctx *VersionOperatorContext) {}

// EnterVersionConstraint is called when production versionConstraint is entered.
func (s *BaseSolidityListener) EnterVersionConstraint(ctx *VersionConstraintContext) {}

// ExitVersionConstraint is called when production versionConstraint is exited.
func (s *BaseSolidityListener) ExitVersionConstraint(ctx *VersionConstraintContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseSolidityListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseSolidityListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterImportDirective is called when production importDirective is entered.
func (s *BaseSolidityListener) EnterImportDirective(ctx *ImportDirectiveContext) {}

// ExitImportDirective is called when production importDirective is exited.
func (s *BaseSolidityListener) ExitImportDirective(ctx *ImportDirectiveContext) {}

// EnterContractDefinition is called when production contractDefinition is entered.
func (s *BaseSolidityListener) EnterContractDefinition(ctx *ContractDefinitionContext) {}

// ExitContractDefinition is called when production contractDefinition is exited.
func (s *BaseSolidityListener) ExitContractDefinition(ctx *ContractDefinitionContext) {}

// EnterInheritanceSpecifier is called when production inheritanceSpecifier is entered.
func (s *BaseSolidityListener) EnterInheritanceSpecifier(ctx *InheritanceSpecifierContext) {}

// ExitInheritanceSpecifier is called when production inheritanceSpecifier is exited.
func (s *BaseSolidityListener) ExitInheritanceSpecifier(ctx *InheritanceSpecifierContext) {}

// EnterContractPart is called when production contractPart is entered.
func (s *BaseSolidityListener) EnterContractPart(ctx *ContractPartContext) {}

// ExitContractPart is called when production contractPart is exited.
func (s *BaseSolidityListener) ExitContractPart(ctx *ContractPartContext) {}

// EnterStateVariableDeclaration is called when production stateVariableDeclaration is entered.
func (s *BaseSolidityListener) EnterStateVariableDeclaration(ctx *StateVariableDeclarationContext) {}

// ExitStateVariableDeclaration is called when production stateVariableDeclaration is exited.
func (s *BaseSolidityListener) ExitStateVariableDeclaration(ctx *StateVariableDeclarationContext) {}

// EnterUsingForDeclaration is called when production usingForDeclaration is entered.
func (s *BaseSolidityListener) EnterUsingForDeclaration(ctx *UsingForDeclarationContext) {}

// ExitUsingForDeclaration is called when production usingForDeclaration is exited.
func (s *BaseSolidityListener) ExitUsingForDeclaration(ctx *UsingForDeclarationContext) {}

// EnterStructDefinition is called when production structDefinition is entered.
func (s *BaseSolidityListener) EnterStructDefinition(ctx *StructDefinitionContext) {}

// ExitStructDefinition is called when production structDefinition is exited.
func (s *BaseSolidityListener) ExitStructDefinition(ctx *StructDefinitionContext) {}

// EnterConstructorDefinition is called when production constructorDefinition is entered.
func (s *BaseSolidityListener) EnterConstructorDefinition(ctx *ConstructorDefinitionContext) {}

// ExitConstructorDefinition is called when production constructorDefinition is exited.
func (s *BaseSolidityListener) ExitConstructorDefinition(ctx *ConstructorDefinitionContext) {}

// EnterModifierDefinition is called when production modifierDefinition is entered.
func (s *BaseSolidityListener) EnterModifierDefinition(ctx *ModifierDefinitionContext) {}

// ExitModifierDefinition is called when production modifierDefinition is exited.
func (s *BaseSolidityListener) ExitModifierDefinition(ctx *ModifierDefinitionContext) {}

// EnterModifierInvocation is called when production modifierInvocation is entered.
func (s *BaseSolidityListener) EnterModifierInvocation(ctx *ModifierInvocationContext) {}

// ExitModifierInvocation is called when production modifierInvocation is exited.
func (s *BaseSolidityListener) ExitModifierInvocation(ctx *ModifierInvocationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseSolidityListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseSolidityListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterReturnParameters is called when production returnParameters is entered.
func (s *BaseSolidityListener) EnterReturnParameters(ctx *ReturnParametersContext) {}

// ExitReturnParameters is called when production returnParameters is exited.
func (s *BaseSolidityListener) ExitReturnParameters(ctx *ReturnParametersContext) {}

// EnterModifierList is called when production modifierList is entered.
func (s *BaseSolidityListener) EnterModifierList(ctx *ModifierListContext) {}

// ExitModifierList is called when production modifierList is exited.
func (s *BaseSolidityListener) ExitModifierList(ctx *ModifierListContext) {}

// EnterEventDefinition is called when production eventDefinition is entered.
func (s *BaseSolidityListener) EnterEventDefinition(ctx *EventDefinitionContext) {}

// ExitEventDefinition is called when production eventDefinition is exited.
func (s *BaseSolidityListener) ExitEventDefinition(ctx *EventDefinitionContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseSolidityListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseSolidityListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterEnumDefinition is called when production enumDefinition is entered.
func (s *BaseSolidityListener) EnterEnumDefinition(ctx *EnumDefinitionContext) {}

// ExitEnumDefinition is called when production enumDefinition is exited.
func (s *BaseSolidityListener) ExitEnumDefinition(ctx *EnumDefinitionContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseSolidityListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseSolidityListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseSolidityListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseSolidityListener) ExitParameter(ctx *ParameterContext) {}

// EnterEventParameterList is called when production eventParameterList is entered.
func (s *BaseSolidityListener) EnterEventParameterList(ctx *EventParameterListContext) {}

// ExitEventParameterList is called when production eventParameterList is exited.
func (s *BaseSolidityListener) ExitEventParameterList(ctx *EventParameterListContext) {}

// EnterEventParameter is called when production eventParameter is entered.
func (s *BaseSolidityListener) EnterEventParameter(ctx *EventParameterContext) {}

// ExitEventParameter is called when production eventParameter is exited.
func (s *BaseSolidityListener) ExitEventParameter(ctx *EventParameterContext) {}

// EnterFunctionTypeParameterList is called when production functionTypeParameterList is entered.
func (s *BaseSolidityListener) EnterFunctionTypeParameterList(ctx *FunctionTypeParameterListContext) {}

// ExitFunctionTypeParameterList is called when production functionTypeParameterList is exited.
func (s *BaseSolidityListener) ExitFunctionTypeParameterList(ctx *FunctionTypeParameterListContext) {}

// EnterFunctionTypeParameter is called when production functionTypeParameter is entered.
func (s *BaseSolidityListener) EnterFunctionTypeParameter(ctx *FunctionTypeParameterContext) {}

// ExitFunctionTypeParameter is called when production functionTypeParameter is exited.
func (s *BaseSolidityListener) ExitFunctionTypeParameter(ctx *FunctionTypeParameterContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseSolidityListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseSolidityListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseSolidityListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseSolidityListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterUserDefinedTypeName is called when production userDefinedTypeName is entered.
func (s *BaseSolidityListener) EnterUserDefinedTypeName(ctx *UserDefinedTypeNameContext) {}

// ExitUserDefinedTypeName is called when production userDefinedTypeName is exited.
func (s *BaseSolidityListener) ExitUserDefinedTypeName(ctx *UserDefinedTypeNameContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BaseSolidityListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BaseSolidityListener) ExitMapping(ctx *MappingContext) {}

// EnterFunctionTypeName is called when production functionTypeName is entered.
func (s *BaseSolidityListener) EnterFunctionTypeName(ctx *FunctionTypeNameContext) {}

// ExitFunctionTypeName is called when production functionTypeName is exited.
func (s *BaseSolidityListener) ExitFunctionTypeName(ctx *FunctionTypeNameContext) {}

// EnterStorageLocation is called when production storageLocation is entered.
func (s *BaseSolidityListener) EnterStorageLocation(ctx *StorageLocationContext) {}

// ExitStorageLocation is called when production storageLocation is exited.
func (s *BaseSolidityListener) ExitStorageLocation(ctx *StorageLocationContext) {}

// EnterStateMutability is called when production stateMutability is entered.
func (s *BaseSolidityListener) EnterStateMutability(ctx *StateMutabilityContext) {}

// ExitStateMutability is called when production stateMutability is exited.
func (s *BaseSolidityListener) ExitStateMutability(ctx *StateMutabilityContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSolidityListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSolidityListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSolidityListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSolidityListener) ExitStatement(ctx *StatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseSolidityListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseSolidityListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseSolidityListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseSolidityListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseSolidityListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseSolidityListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseSolidityListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseSolidityListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseSolidityListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseSolidityListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterInlineAssemblyStatement is called when production inlineAssemblyStatement is entered.
func (s *BaseSolidityListener) EnterInlineAssemblyStatement(ctx *InlineAssemblyStatementContext) {}

// ExitInlineAssemblyStatement is called when production inlineAssemblyStatement is exited.
func (s *BaseSolidityListener) ExitInlineAssemblyStatement(ctx *InlineAssemblyStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseSolidityListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseSolidityListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseSolidityListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseSolidityListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseSolidityListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseSolidityListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseSolidityListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseSolidityListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseSolidityListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseSolidityListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterEmitStatement is called when production emitStatement is entered.
func (s *BaseSolidityListener) EnterEmitStatement(ctx *EmitStatementContext) {}

// ExitEmitStatement is called when production emitStatement is exited.
func (s *BaseSolidityListener) ExitEmitStatement(ctx *EmitStatementContext) {}

// EnterVariableDeclarationStatement is called when production variableDeclarationStatement is entered.
func (s *BaseSolidityListener) EnterVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// ExitVariableDeclarationStatement is called when production variableDeclarationStatement is exited.
func (s *BaseSolidityListener) ExitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) {
}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseSolidityListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseSolidityListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseSolidityListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseSolidityListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterElementaryTypeName is called when production elementaryTypeName is entered.
func (s *BaseSolidityListener) EnterElementaryTypeName(ctx *ElementaryTypeNameContext) {}

// ExitElementaryTypeName is called when production elementaryTypeName is exited.
func (s *BaseSolidityListener) ExitElementaryTypeName(ctx *ElementaryTypeNameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSolidityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSolidityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSolidityListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSolidityListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseSolidityListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseSolidityListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterNameValueList is called when production nameValueList is entered.
func (s *BaseSolidityListener) EnterNameValueList(ctx *NameValueListContext) {}

// ExitNameValueList is called when production nameValueList is exited.
func (s *BaseSolidityListener) ExitNameValueList(ctx *NameValueListContext) {}

// EnterNameValue is called when production nameValue is entered.
func (s *BaseSolidityListener) EnterNameValue(ctx *NameValueContext) {}

// ExitNameValue is called when production nameValue is exited.
func (s *BaseSolidityListener) ExitNameValue(ctx *NameValueContext) {}

// EnterFunctionCallArguments is called when production functionCallArguments is entered.
func (s *BaseSolidityListener) EnterFunctionCallArguments(ctx *FunctionCallArgumentsContext) {}

// ExitFunctionCallArguments is called when production functionCallArguments is exited.
func (s *BaseSolidityListener) ExitFunctionCallArguments(ctx *FunctionCallArgumentsContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSolidityListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSolidityListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAssemblyBlock is called when production assemblyBlock is entered.
func (s *BaseSolidityListener) EnterAssemblyBlock(ctx *AssemblyBlockContext) {}

// ExitAssemblyBlock is called when production assemblyBlock is exited.
func (s *BaseSolidityListener) ExitAssemblyBlock(ctx *AssemblyBlockContext) {}

// EnterAssemblyItem is called when production assemblyItem is entered.
func (s *BaseSolidityListener) EnterAssemblyItem(ctx *AssemblyItemContext) {}

// ExitAssemblyItem is called when production assemblyItem is exited.
func (s *BaseSolidityListener) ExitAssemblyItem(ctx *AssemblyItemContext) {}

// EnterAssemblyExpression is called when production assemblyExpression is entered.
func (s *BaseSolidityListener) EnterAssemblyExpression(ctx *AssemblyExpressionContext) {}

// ExitAssemblyExpression is called when production assemblyExpression is exited.
func (s *BaseSolidityListener) ExitAssemblyExpression(ctx *AssemblyExpressionContext) {}

// EnterAssemblyCall is called when production assemblyCall is entered.
func (s *BaseSolidityListener) EnterAssemblyCall(ctx *AssemblyCallContext) {}

// ExitAssemblyCall is called when production assemblyCall is exited.
func (s *BaseSolidityListener) ExitAssemblyCall(ctx *AssemblyCallContext) {}

// EnterAssemblyLocalDefinition is called when production assemblyLocalDefinition is entered.
func (s *BaseSolidityListener) EnterAssemblyLocalDefinition(ctx *AssemblyLocalDefinitionContext) {}

// ExitAssemblyLocalDefinition is called when production assemblyLocalDefinition is exited.
func (s *BaseSolidityListener) ExitAssemblyLocalDefinition(ctx *AssemblyLocalDefinitionContext) {}

// EnterAssemblyAssignment is called when production assemblyAssignment is entered.
func (s *BaseSolidityListener) EnterAssemblyAssignment(ctx *AssemblyAssignmentContext) {}

// ExitAssemblyAssignment is called when production assemblyAssignment is exited.
func (s *BaseSolidityListener) ExitAssemblyAssignment(ctx *AssemblyAssignmentContext) {}

// EnterAssemblyIdentifierOrList is called when production assemblyIdentifierOrList is entered.
func (s *BaseSolidityListener) EnterAssemblyIdentifierOrList(ctx *AssemblyIdentifierOrListContext) {}

// ExitAssemblyIdentifierOrList is called when production assemblyIdentifierOrList is exited.
func (s *BaseSolidityListener) ExitAssemblyIdentifierOrList(ctx *AssemblyIdentifierOrListContext) {}

// EnterAssemblyIdentifierList is called when production assemblyIdentifierList is entered.
func (s *BaseSolidityListener) EnterAssemblyIdentifierList(ctx *AssemblyIdentifierListContext) {}

// ExitAssemblyIdentifierList is called when production assemblyIdentifierList is exited.
func (s *BaseSolidityListener) ExitAssemblyIdentifierList(ctx *AssemblyIdentifierListContext) {}

// EnterAssemblyStackAssignment is called when production assemblyStackAssignment is entered.
func (s *BaseSolidityListener) EnterAssemblyStackAssignment(ctx *AssemblyStackAssignmentContext) {}

// ExitAssemblyStackAssignment is called when production assemblyStackAssignment is exited.
func (s *BaseSolidityListener) ExitAssemblyStackAssignment(ctx *AssemblyStackAssignmentContext) {}

// EnterLabelDefinition is called when production labelDefinition is entered.
func (s *BaseSolidityListener) EnterLabelDefinition(ctx *LabelDefinitionContext) {}

// ExitLabelDefinition is called when production labelDefinition is exited.
func (s *BaseSolidityListener) ExitLabelDefinition(ctx *LabelDefinitionContext) {}

// EnterAssemblySwitch is called when production assemblySwitch is entered.
func (s *BaseSolidityListener) EnterAssemblySwitch(ctx *AssemblySwitchContext) {}

// ExitAssemblySwitch is called when production assemblySwitch is exited.
func (s *BaseSolidityListener) ExitAssemblySwitch(ctx *AssemblySwitchContext) {}

// EnterAssemblyCase is called when production assemblyCase is entered.
func (s *BaseSolidityListener) EnterAssemblyCase(ctx *AssemblyCaseContext) {}

// ExitAssemblyCase is called when production assemblyCase is exited.
func (s *BaseSolidityListener) ExitAssemblyCase(ctx *AssemblyCaseContext) {}

// EnterAssemblyFunctionDefinition is called when production assemblyFunctionDefinition is entered.
func (s *BaseSolidityListener) EnterAssemblyFunctionDefinition(ctx *AssemblyFunctionDefinitionContext) {
}

// ExitAssemblyFunctionDefinition is called when production assemblyFunctionDefinition is exited.
func (s *BaseSolidityListener) ExitAssemblyFunctionDefinition(ctx *AssemblyFunctionDefinitionContext) {
}

// EnterAssemblyFunctionReturns is called when production assemblyFunctionReturns is entered.
func (s *BaseSolidityListener) EnterAssemblyFunctionReturns(ctx *AssemblyFunctionReturnsContext) {}

// ExitAssemblyFunctionReturns is called when production assemblyFunctionReturns is exited.
func (s *BaseSolidityListener) ExitAssemblyFunctionReturns(ctx *AssemblyFunctionReturnsContext) {}

// EnterAssemblyFor is called when production assemblyFor is entered.
func (s *BaseSolidityListener) EnterAssemblyFor(ctx *AssemblyForContext) {}

// ExitAssemblyFor is called when production assemblyFor is exited.
func (s *BaseSolidityListener) ExitAssemblyFor(ctx *AssemblyForContext) {}

// EnterAssemblyIf is called when production assemblyIf is entered.
func (s *BaseSolidityListener) EnterAssemblyIf(ctx *AssemblyIfContext) {}

// ExitAssemblyIf is called when production assemblyIf is exited.
func (s *BaseSolidityListener) ExitAssemblyIf(ctx *AssemblyIfContext) {}

// EnterAssemblyLiteral is called when production assemblyLiteral is entered.
func (s *BaseSolidityListener) EnterAssemblyLiteral(ctx *AssemblyLiteralContext) {}

// ExitAssemblyLiteral is called when production assemblyLiteral is exited.
func (s *BaseSolidityListener) ExitAssemblyLiteral(ctx *AssemblyLiteralContext) {}

// EnterSubAssembly is called when production subAssembly is entered.
func (s *BaseSolidityListener) EnterSubAssembly(ctx *SubAssemblyContext) {}

// ExitSubAssembly is called when production subAssembly is exited.
func (s *BaseSolidityListener) ExitSubAssembly(ctx *SubAssemblyContext) {}

// EnterTupleExpression is called when production tupleExpression is entered.
func (s *BaseSolidityListener) EnterTupleExpression(ctx *TupleExpressionContext) {}

// ExitTupleExpression is called when production tupleExpression is exited.
func (s *BaseSolidityListener) ExitTupleExpression(ctx *TupleExpressionContext) {}

// EnterElementaryTypeNameExpression is called when production elementaryTypeNameExpression is entered.
func (s *BaseSolidityListener) EnterElementaryTypeNameExpression(ctx *ElementaryTypeNameExpressionContext) {
}

// ExitElementaryTypeNameExpression is called when production elementaryTypeNameExpression is exited.
func (s *BaseSolidityListener) ExitElementaryTypeNameExpression(ctx *ElementaryTypeNameExpressionContext) {
}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseSolidityListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseSolidityListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSolidityListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSolidityListener) ExitIdentifier(ctx *IdentifierContext) {}
