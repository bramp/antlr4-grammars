// Code generated from Dart2.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dart2 // Dart2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDart2Listener is a complete listener for a parse tree produced by Dart2Parser.
type BaseDart2Listener struct{}

var _ Dart2Listener = &BaseDart2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDart2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDart2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDart2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDart2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseDart2Listener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseDart2Listener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseDart2Listener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseDart2Listener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterDeclaredIdentifier is called when production declaredIdentifier is entered.
func (s *BaseDart2Listener) EnterDeclaredIdentifier(ctx *DeclaredIdentifierContext) {}

// ExitDeclaredIdentifier is called when production declaredIdentifier is exited.
func (s *BaseDart2Listener) ExitDeclaredIdentifier(ctx *DeclaredIdentifierContext) {}

// EnterFinalConstVarOrType is called when production finalConstVarOrType is entered.
func (s *BaseDart2Listener) EnterFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) {}

// ExitFinalConstVarOrType is called when production finalConstVarOrType is exited.
func (s *BaseDart2Listener) ExitFinalConstVarOrType(ctx *FinalConstVarOrTypeContext) {}

// EnterVarOrType is called when production varOrType is entered.
func (s *BaseDart2Listener) EnterVarOrType(ctx *VarOrTypeContext) {}

// ExitVarOrType is called when production varOrType is exited.
func (s *BaseDart2Listener) ExitVarOrType(ctx *VarOrTypeContext) {}

// EnterInitializedVariableDeclaration is called when production initializedVariableDeclaration is entered.
func (s *BaseDart2Listener) EnterInitializedVariableDeclaration(ctx *InitializedVariableDeclarationContext) {
}

// ExitInitializedVariableDeclaration is called when production initializedVariableDeclaration is exited.
func (s *BaseDart2Listener) ExitInitializedVariableDeclaration(ctx *InitializedVariableDeclarationContext) {
}

// EnterInitializedIdentifier is called when production initializedIdentifier is entered.
func (s *BaseDart2Listener) EnterInitializedIdentifier(ctx *InitializedIdentifierContext) {}

// ExitInitializedIdentifier is called when production initializedIdentifier is exited.
func (s *BaseDart2Listener) ExitInitializedIdentifier(ctx *InitializedIdentifierContext) {}

// EnterInitializedIdentifierList is called when production initializedIdentifierList is entered.
func (s *BaseDart2Listener) EnterInitializedIdentifierList(ctx *InitializedIdentifierListContext) {}

// ExitInitializedIdentifierList is called when production initializedIdentifierList is exited.
func (s *BaseDart2Listener) ExitInitializedIdentifierList(ctx *InitializedIdentifierListContext) {}

// EnterFunctionSignature is called when production functionSignature is entered.
func (s *BaseDart2Listener) EnterFunctionSignature(ctx *FunctionSignatureContext) {}

// ExitFunctionSignature is called when production functionSignature is exited.
func (s *BaseDart2Listener) ExitFunctionSignature(ctx *FunctionSignatureContext) {}

// EnterFormalParameterPart is called when production formalParameterPart is entered.
func (s *BaseDart2Listener) EnterFormalParameterPart(ctx *FormalParameterPartContext) {}

// ExitFormalParameterPart is called when production formalParameterPart is exited.
func (s *BaseDart2Listener) ExitFormalParameterPart(ctx *FormalParameterPartContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseDart2Listener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseDart2Listener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseDart2Listener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseDart2Listener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDart2Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDart2Listener) ExitBlock(ctx *BlockContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseDart2Listener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseDart2Listener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterNormalFormalParameters is called when production normalFormalParameters is entered.
func (s *BaseDart2Listener) EnterNormalFormalParameters(ctx *NormalFormalParametersContext) {}

// ExitNormalFormalParameters is called when production normalFormalParameters is exited.
func (s *BaseDart2Listener) ExitNormalFormalParameters(ctx *NormalFormalParametersContext) {}

// EnterOptionalFormalParameters is called when production optionalFormalParameters is entered.
func (s *BaseDart2Listener) EnterOptionalFormalParameters(ctx *OptionalFormalParametersContext) {}

// ExitOptionalFormalParameters is called when production optionalFormalParameters is exited.
func (s *BaseDart2Listener) ExitOptionalFormalParameters(ctx *OptionalFormalParametersContext) {}

// EnterOptionalPositionalFormalParameters is called when production optionalPositionalFormalParameters is entered.
func (s *BaseDart2Listener) EnterOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) {
}

// ExitOptionalPositionalFormalParameters is called when production optionalPositionalFormalParameters is exited.
func (s *BaseDart2Listener) ExitOptionalPositionalFormalParameters(ctx *OptionalPositionalFormalParametersContext) {
}

// EnterNamedFormalParameters is called when production namedFormalParameters is entered.
func (s *BaseDart2Listener) EnterNamedFormalParameters(ctx *NamedFormalParametersContext) {}

// ExitNamedFormalParameters is called when production namedFormalParameters is exited.
func (s *BaseDart2Listener) ExitNamedFormalParameters(ctx *NamedFormalParametersContext) {}

// EnterNormalFormalParameter is called when production normalFormalParameter is entered.
func (s *BaseDart2Listener) EnterNormalFormalParameter(ctx *NormalFormalParameterContext) {}

// ExitNormalFormalParameter is called when production normalFormalParameter is exited.
func (s *BaseDart2Listener) ExitNormalFormalParameter(ctx *NormalFormalParameterContext) {}

// EnterFunctionFormalParameter is called when production functionFormalParameter is entered.
func (s *BaseDart2Listener) EnterFunctionFormalParameter(ctx *FunctionFormalParameterContext) {}

// ExitFunctionFormalParameter is called when production functionFormalParameter is exited.
func (s *BaseDart2Listener) ExitFunctionFormalParameter(ctx *FunctionFormalParameterContext) {}

// EnterSimpleFormalParameter is called when production simpleFormalParameter is entered.
func (s *BaseDart2Listener) EnterSimpleFormalParameter(ctx *SimpleFormalParameterContext) {}

// ExitSimpleFormalParameter is called when production simpleFormalParameter is exited.
func (s *BaseDart2Listener) ExitSimpleFormalParameter(ctx *SimpleFormalParameterContext) {}

// EnterFieldFormalParameter is called when production fieldFormalParameter is entered.
func (s *BaseDart2Listener) EnterFieldFormalParameter(ctx *FieldFormalParameterContext) {}

// ExitFieldFormalParameter is called when production fieldFormalParameter is exited.
func (s *BaseDart2Listener) ExitFieldFormalParameter(ctx *FieldFormalParameterContext) {}

// EnterDefaultFormalParameter is called when production defaultFormalParameter is entered.
func (s *BaseDart2Listener) EnterDefaultFormalParameter(ctx *DefaultFormalParameterContext) {}

// ExitDefaultFormalParameter is called when production defaultFormalParameter is exited.
func (s *BaseDart2Listener) ExitDefaultFormalParameter(ctx *DefaultFormalParameterContext) {}

// EnterDefaultNamedParameter is called when production defaultNamedParameter is entered.
func (s *BaseDart2Listener) EnterDefaultNamedParameter(ctx *DefaultNamedParameterContext) {}

// ExitDefaultNamedParameter is called when production defaultNamedParameter is exited.
func (s *BaseDart2Listener) ExitDefaultNamedParameter(ctx *DefaultNamedParameterContext) {}

// EnterClassDefinition is called when production classDefinition is entered.
func (s *BaseDart2Listener) EnterClassDefinition(ctx *ClassDefinitionContext) {}

// ExitClassDefinition is called when production classDefinition is exited.
func (s *BaseDart2Listener) ExitClassDefinition(ctx *ClassDefinitionContext) {}

// EnterMixins is called when production mixins is entered.
func (s *BaseDart2Listener) EnterMixins(ctx *MixinsContext) {}

// ExitMixins is called when production mixins is exited.
func (s *BaseDart2Listener) ExitMixins(ctx *MixinsContext) {}

// EnterClassMemberDefinition is called when production classMemberDefinition is entered.
func (s *BaseDart2Listener) EnterClassMemberDefinition(ctx *ClassMemberDefinitionContext) {}

// ExitClassMemberDefinition is called when production classMemberDefinition is exited.
func (s *BaseDart2Listener) ExitClassMemberDefinition(ctx *ClassMemberDefinitionContext) {}

// EnterMethodSignature is called when production methodSignature is entered.
func (s *BaseDart2Listener) EnterMethodSignature(ctx *MethodSignatureContext) {}

// ExitMethodSignature is called when production methodSignature is exited.
func (s *BaseDart2Listener) ExitMethodSignature(ctx *MethodSignatureContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDart2Listener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDart2Listener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterStaticFinalDeclarationList is called when production staticFinalDeclarationList is entered.
func (s *BaseDart2Listener) EnterStaticFinalDeclarationList(ctx *StaticFinalDeclarationListContext) {}

// ExitStaticFinalDeclarationList is called when production staticFinalDeclarationList is exited.
func (s *BaseDart2Listener) ExitStaticFinalDeclarationList(ctx *StaticFinalDeclarationListContext) {}

// EnterStaticFinalDeclaration is called when production staticFinalDeclaration is entered.
func (s *BaseDart2Listener) EnterStaticFinalDeclaration(ctx *StaticFinalDeclarationContext) {}

// ExitStaticFinalDeclaration is called when production staticFinalDeclaration is exited.
func (s *BaseDart2Listener) ExitStaticFinalDeclaration(ctx *StaticFinalDeclarationContext) {}

// EnterOperatorSignature is called when production operatorSignature is entered.
func (s *BaseDart2Listener) EnterOperatorSignature(ctx *OperatorSignatureContext) {}

// ExitOperatorSignature is called when production operatorSignature is exited.
func (s *BaseDart2Listener) ExitOperatorSignature(ctx *OperatorSignatureContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseDart2Listener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseDart2Listener) ExitOperator(ctx *OperatorContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BaseDart2Listener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BaseDart2Listener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterGetterSignature is called when production getterSignature is entered.
func (s *BaseDart2Listener) EnterGetterSignature(ctx *GetterSignatureContext) {}

// ExitGetterSignature is called when production getterSignature is exited.
func (s *BaseDart2Listener) ExitGetterSignature(ctx *GetterSignatureContext) {}

// EnterSetterSignature is called when production setterSignature is entered.
func (s *BaseDart2Listener) EnterSetterSignature(ctx *SetterSignatureContext) {}

// ExitSetterSignature is called when production setterSignature is exited.
func (s *BaseDart2Listener) ExitSetterSignature(ctx *SetterSignatureContext) {}

// EnterConstructorSignature is called when production constructorSignature is entered.
func (s *BaseDart2Listener) EnterConstructorSignature(ctx *ConstructorSignatureContext) {}

// ExitConstructorSignature is called when production constructorSignature is exited.
func (s *BaseDart2Listener) ExitConstructorSignature(ctx *ConstructorSignatureContext) {}

// EnterRedirection is called when production redirection is entered.
func (s *BaseDart2Listener) EnterRedirection(ctx *RedirectionContext) {}

// ExitRedirection is called when production redirection is exited.
func (s *BaseDart2Listener) ExitRedirection(ctx *RedirectionContext) {}

// EnterInitializers is called when production initializers is entered.
func (s *BaseDart2Listener) EnterInitializers(ctx *InitializersContext) {}

// ExitInitializers is called when production initializers is exited.
func (s *BaseDart2Listener) ExitInitializers(ctx *InitializersContext) {}

// EnterInitializerListEntry is called when production initializerListEntry is entered.
func (s *BaseDart2Listener) EnterInitializerListEntry(ctx *InitializerListEntryContext) {}

// ExitInitializerListEntry is called when production initializerListEntry is exited.
func (s *BaseDart2Listener) ExitInitializerListEntry(ctx *InitializerListEntryContext) {}

// EnterFieldInitializer is called when production fieldInitializer is entered.
func (s *BaseDart2Listener) EnterFieldInitializer(ctx *FieldInitializerContext) {}

// ExitFieldInitializer is called when production fieldInitializer is exited.
func (s *BaseDart2Listener) ExitFieldInitializer(ctx *FieldInitializerContext) {}

// EnterFactoryConstructorSignature is called when production factoryConstructorSignature is entered.
func (s *BaseDart2Listener) EnterFactoryConstructorSignature(ctx *FactoryConstructorSignatureContext) {
}

// ExitFactoryConstructorSignature is called when production factoryConstructorSignature is exited.
func (s *BaseDart2Listener) ExitFactoryConstructorSignature(ctx *FactoryConstructorSignatureContext) {}

// EnterRedirectingFactoryConstructorSignature is called when production redirectingFactoryConstructorSignature is entered.
func (s *BaseDart2Listener) EnterRedirectingFactoryConstructorSignature(ctx *RedirectingFactoryConstructorSignatureContext) {
}

// ExitRedirectingFactoryConstructorSignature is called when production redirectingFactoryConstructorSignature is exited.
func (s *BaseDart2Listener) ExitRedirectingFactoryConstructorSignature(ctx *RedirectingFactoryConstructorSignatureContext) {
}

// EnterConstantConstructorSignature is called when production constantConstructorSignature is entered.
func (s *BaseDart2Listener) EnterConstantConstructorSignature(ctx *ConstantConstructorSignatureContext) {
}

// ExitConstantConstructorSignature is called when production constantConstructorSignature is exited.
func (s *BaseDart2Listener) ExitConstantConstructorSignature(ctx *ConstantConstructorSignatureContext) {
}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseDart2Listener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseDart2Listener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterInterfaces is called when production interfaces is entered.
func (s *BaseDart2Listener) EnterInterfaces(ctx *InterfacesContext) {}

// ExitInterfaces is called when production interfaces is exited.
func (s *BaseDart2Listener) ExitInterfaces(ctx *InterfacesContext) {}

// EnterMixinApplicationClass is called when production mixinApplicationClass is entered.
func (s *BaseDart2Listener) EnterMixinApplicationClass(ctx *MixinApplicationClassContext) {}

// ExitMixinApplicationClass is called when production mixinApplicationClass is exited.
func (s *BaseDart2Listener) ExitMixinApplicationClass(ctx *MixinApplicationClassContext) {}

// EnterMixinApplication is called when production mixinApplication is entered.
func (s *BaseDart2Listener) EnterMixinApplication(ctx *MixinApplicationContext) {}

// ExitMixinApplication is called when production mixinApplication is exited.
func (s *BaseDart2Listener) ExitMixinApplication(ctx *MixinApplicationContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseDart2Listener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseDart2Listener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterEnumEntry is called when production enumEntry is entered.
func (s *BaseDart2Listener) EnterEnumEntry(ctx *EnumEntryContext) {}

// ExitEnumEntry is called when production enumEntry is exited.
func (s *BaseDart2Listener) ExitEnumEntry(ctx *EnumEntryContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseDart2Listener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseDart2Listener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseDart2Listener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseDart2Listener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterMetadata is called when production metadata is entered.
func (s *BaseDart2Listener) EnterMetadata(ctx *MetadataContext) {}

// ExitMetadata is called when production metadata is exited.
func (s *BaseDart2Listener) ExitMetadata(ctx *MetadataContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDart2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDart2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionWithoutCascade is called when production expressionWithoutCascade is entered.
func (s *BaseDart2Listener) EnterExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) {}

// ExitExpressionWithoutCascade is called when production expressionWithoutCascade is exited.
func (s *BaseDart2Listener) ExitExpressionWithoutCascade(ctx *ExpressionWithoutCascadeContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseDart2Listener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseDart2Listener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseDart2Listener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseDart2Listener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDart2Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDart2Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseDart2Listener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseDart2Listener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseDart2Listener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseDart2Listener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseDart2Listener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseDart2Listener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseDart2Listener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseDart2Listener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterStringInterpolation is called when production stringInterpolation is entered.
func (s *BaseDart2Listener) EnterStringInterpolation(ctx *StringInterpolationContext) {}

// ExitStringInterpolation is called when production stringInterpolation is exited.
func (s *BaseDart2Listener) ExitStringInterpolation(ctx *StringInterpolationContext) {}

// EnterSymbolLiteral is called when production symbolLiteral is entered.
func (s *BaseDart2Listener) EnterSymbolLiteral(ctx *SymbolLiteralContext) {}

// ExitSymbolLiteral is called when production symbolLiteral is exited.
func (s *BaseDart2Listener) ExitSymbolLiteral(ctx *SymbolLiteralContext) {}

// EnterListLiteral is called when production listLiteral is entered.
func (s *BaseDart2Listener) EnterListLiteral(ctx *ListLiteralContext) {}

// ExitListLiteral is called when production listLiteral is exited.
func (s *BaseDart2Listener) ExitListLiteral(ctx *ListLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseDart2Listener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseDart2Listener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterMapLiteralEntry is called when production mapLiteralEntry is entered.
func (s *BaseDart2Listener) EnterMapLiteralEntry(ctx *MapLiteralEntryContext) {}

// ExitMapLiteralEntry is called when production mapLiteralEntry is exited.
func (s *BaseDart2Listener) ExitMapLiteralEntry(ctx *MapLiteralEntryContext) {}

// EnterThrowExpression is called when production throwExpression is entered.
func (s *BaseDart2Listener) EnterThrowExpression(ctx *ThrowExpressionContext) {}

// ExitThrowExpression is called when production throwExpression is exited.
func (s *BaseDart2Listener) ExitThrowExpression(ctx *ThrowExpressionContext) {}

// EnterThrowExpressionWithoutCascade is called when production throwExpressionWithoutCascade is entered.
func (s *BaseDart2Listener) EnterThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) {
}

// ExitThrowExpressionWithoutCascade is called when production throwExpressionWithoutCascade is exited.
func (s *BaseDart2Listener) ExitThrowExpressionWithoutCascade(ctx *ThrowExpressionWithoutCascadeContext) {
}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseDart2Listener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseDart2Listener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *BaseDart2Listener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *BaseDart2Listener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterNayaExpression is called when production nayaExpression is entered.
func (s *BaseDart2Listener) EnterNayaExpression(ctx *NayaExpressionContext) {}

// ExitNayaExpression is called when production nayaExpression is exited.
func (s *BaseDart2Listener) ExitNayaExpression(ctx *NayaExpressionContext) {}

// EnterConstObjectExpression is called when production constObjectExpression is entered.
func (s *BaseDart2Listener) EnterConstObjectExpression(ctx *ConstObjectExpressionContext) {}

// ExitConstObjectExpression is called when production constObjectExpression is exited.
func (s *BaseDart2Listener) ExitConstObjectExpression(ctx *ConstObjectExpressionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseDart2Listener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseDart2Listener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseDart2Listener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseDart2Listener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgument is called when production namedArgument is entered.
func (s *BaseDart2Listener) EnterNamedArgument(ctx *NamedArgumentContext) {}

// ExitNamedArgument is called when production namedArgument is exited.
func (s *BaseDart2Listener) ExitNamedArgument(ctx *NamedArgumentContext) {}

// EnterCascadeSection is called when production cascadeSection is entered.
func (s *BaseDart2Listener) EnterCascadeSection(ctx *CascadeSectionContext) {}

// ExitCascadeSection is called when production cascadeSection is exited.
func (s *BaseDart2Listener) ExitCascadeSection(ctx *CascadeSectionContext) {}

// EnterCascadeSelector is called when production cascadeSelector is entered.
func (s *BaseDart2Listener) EnterCascadeSelector(ctx *CascadeSelectorContext) {}

// ExitCascadeSelector is called when production cascadeSelector is exited.
func (s *BaseDart2Listener) ExitCascadeSelector(ctx *CascadeSelectorContext) {}

// EnterArgumentPart is called when production argumentPart is entered.
func (s *BaseDart2Listener) EnterArgumentPart(ctx *ArgumentPartContext) {}

// ExitArgumentPart is called when production argumentPart is exited.
func (s *BaseDart2Listener) ExitArgumentPart(ctx *ArgumentPartContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseDart2Listener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseDart2Listener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterCompoundAssignmentOperator is called when production compoundAssignmentOperator is entered.
func (s *BaseDart2Listener) EnterCompoundAssignmentOperator(ctx *CompoundAssignmentOperatorContext) {}

// ExitCompoundAssignmentOperator is called when production compoundAssignmentOperator is exited.
func (s *BaseDart2Listener) ExitCompoundAssignmentOperator(ctx *CompoundAssignmentOperatorContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseDart2Listener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseDart2Listener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterIfNullExpression is called when production ifNullExpression is entered.
func (s *BaseDart2Listener) EnterIfNullExpression(ctx *IfNullExpressionContext) {}

// ExitIfNullExpression is called when production ifNullExpression is exited.
func (s *BaseDart2Listener) ExitIfNullExpression(ctx *IfNullExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseDart2Listener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseDart2Listener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseDart2Listener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseDart2Listener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseDart2Listener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseDart2Listener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterEqualityOperator is called when production equalityOperator is entered.
func (s *BaseDart2Listener) EnterEqualityOperator(ctx *EqualityOperatorContext) {}

// ExitEqualityOperator is called when production equalityOperator is exited.
func (s *BaseDart2Listener) ExitEqualityOperator(ctx *EqualityOperatorContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseDart2Listener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseDart2Listener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterRelationalOperator is called when production relationalOperator is entered.
func (s *BaseDart2Listener) EnterRelationalOperator(ctx *RelationalOperatorContext) {}

// ExitRelationalOperator is called when production relationalOperator is exited.
func (s *BaseDart2Listener) ExitRelationalOperator(ctx *RelationalOperatorContext) {}

// EnterBitwiseOrExpression is called when production bitwiseOrExpression is entered.
func (s *BaseDart2Listener) EnterBitwiseOrExpression(ctx *BitwiseOrExpressionContext) {}

// ExitBitwiseOrExpression is called when production bitwiseOrExpression is exited.
func (s *BaseDart2Listener) ExitBitwiseOrExpression(ctx *BitwiseOrExpressionContext) {}

// EnterBitwiseXorExpression is called when production bitwiseXorExpression is entered.
func (s *BaseDart2Listener) EnterBitwiseXorExpression(ctx *BitwiseXorExpressionContext) {}

// ExitBitwiseXorExpression is called when production bitwiseXorExpression is exited.
func (s *BaseDart2Listener) ExitBitwiseXorExpression(ctx *BitwiseXorExpressionContext) {}

// EnterBitwiseAndExpression is called when production bitwiseAndExpression is entered.
func (s *BaseDart2Listener) EnterBitwiseAndExpression(ctx *BitwiseAndExpressionContext) {}

// ExitBitwiseAndExpression is called when production bitwiseAndExpression is exited.
func (s *BaseDart2Listener) ExitBitwiseAndExpression(ctx *BitwiseAndExpressionContext) {}

// EnterBitwiseOperator is called when production bitwiseOperator is entered.
func (s *BaseDart2Listener) EnterBitwiseOperator(ctx *BitwiseOperatorContext) {}

// ExitBitwiseOperator is called when production bitwiseOperator is exited.
func (s *BaseDart2Listener) ExitBitwiseOperator(ctx *BitwiseOperatorContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseDart2Listener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseDart2Listener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterShiftOperator is called when production shiftOperator is entered.
func (s *BaseDart2Listener) EnterShiftOperator(ctx *ShiftOperatorContext) {}

// ExitShiftOperator is called when production shiftOperator is exited.
func (s *BaseDart2Listener) ExitShiftOperator(ctx *ShiftOperatorContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseDart2Listener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseDart2Listener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterAdditiveOperator is called when production additiveOperator is entered.
func (s *BaseDart2Listener) EnterAdditiveOperator(ctx *AdditiveOperatorContext) {}

// ExitAdditiveOperator is called when production additiveOperator is exited.
func (s *BaseDart2Listener) ExitAdditiveOperator(ctx *AdditiveOperatorContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseDart2Listener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseDart2Listener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterMultiplicativeOperator is called when production multiplicativeOperator is entered.
func (s *BaseDart2Listener) EnterMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// ExitMultiplicativeOperator is called when production multiplicativeOperator is exited.
func (s *BaseDart2Listener) ExitMultiplicativeOperator(ctx *MultiplicativeOperatorContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseDart2Listener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseDart2Listener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrefixOperator is called when production prefixOperator is entered.
func (s *BaseDart2Listener) EnterPrefixOperator(ctx *PrefixOperatorContext) {}

// ExitPrefixOperator is called when production prefixOperator is exited.
func (s *BaseDart2Listener) ExitPrefixOperator(ctx *PrefixOperatorContext) {}

// EnterMinusOperator is called when production minusOperator is entered.
func (s *BaseDart2Listener) EnterMinusOperator(ctx *MinusOperatorContext) {}

// ExitMinusOperator is called when production minusOperator is exited.
func (s *BaseDart2Listener) ExitMinusOperator(ctx *MinusOperatorContext) {}

// EnterNegationOperator is called when production negationOperator is entered.
func (s *BaseDart2Listener) EnterNegationOperator(ctx *NegationOperatorContext) {}

// ExitNegationOperator is called when production negationOperator is exited.
func (s *BaseDart2Listener) ExitNegationOperator(ctx *NegationOperatorContext) {}

// EnterTildeOperator is called when production tildeOperator is entered.
func (s *BaseDart2Listener) EnterTildeOperator(ctx *TildeOperatorContext) {}

// ExitTildeOperator is called when production tildeOperator is exited.
func (s *BaseDart2Listener) ExitTildeOperator(ctx *TildeOperatorContext) {}

// EnterAwaitExpression is called when production awaitExpression is entered.
func (s *BaseDart2Listener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production awaitExpression is exited.
func (s *BaseDart2Listener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseDart2Listener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseDart2Listener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostfixOperator is called when production postfixOperator is entered.
func (s *BaseDart2Listener) EnterPostfixOperator(ctx *PostfixOperatorContext) {}

// ExitPostfixOperator is called when production postfixOperator is exited.
func (s *BaseDart2Listener) ExitPostfixOperator(ctx *PostfixOperatorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseDart2Listener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseDart2Listener) ExitSelector(ctx *SelectorContext) {}

// EnterIncrementOperator is called when production incrementOperator is entered.
func (s *BaseDart2Listener) EnterIncrementOperator(ctx *IncrementOperatorContext) {}

// ExitIncrementOperator is called when production incrementOperator is exited.
func (s *BaseDart2Listener) ExitIncrementOperator(ctx *IncrementOperatorContext) {}

// EnterAssignableExpression is called when production assignableExpression is entered.
func (s *BaseDart2Listener) EnterAssignableExpression(ctx *AssignableExpressionContext) {}

// ExitAssignableExpression is called when production assignableExpression is exited.
func (s *BaseDart2Listener) ExitAssignableExpression(ctx *AssignableExpressionContext) {}

// EnterUnconditionalAssignableSelector is called when production unconditionalAssignableSelector is entered.
func (s *BaseDart2Listener) EnterUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) {
}

// ExitUnconditionalAssignableSelector is called when production unconditionalAssignableSelector is exited.
func (s *BaseDart2Listener) ExitUnconditionalAssignableSelector(ctx *UnconditionalAssignableSelectorContext) {
}

// EnterAssignableSelector is called when production assignableSelector is entered.
func (s *BaseDart2Listener) EnterAssignableSelector(ctx *AssignableSelectorContext) {}

// ExitAssignableSelector is called when production assignableSelector is exited.
func (s *BaseDart2Listener) ExitAssignableSelector(ctx *AssignableSelectorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseDart2Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseDart2Listener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterQualified is called when production qualified is entered.
func (s *BaseDart2Listener) EnterQualified(ctx *QualifiedContext) {}

// ExitQualified is called when production qualified is exited.
func (s *BaseDart2Listener) ExitQualified(ctx *QualifiedContext) {}

// EnterTypeTest is called when production typeTest is entered.
func (s *BaseDart2Listener) EnterTypeTest(ctx *TypeTestContext) {}

// ExitTypeTest is called when production typeTest is exited.
func (s *BaseDart2Listener) ExitTypeTest(ctx *TypeTestContext) {}

// EnterIsOperator is called when production isOperator is entered.
func (s *BaseDart2Listener) EnterIsOperator(ctx *IsOperatorContext) {}

// ExitIsOperator is called when production isOperator is exited.
func (s *BaseDart2Listener) ExitIsOperator(ctx *IsOperatorContext) {}

// EnterTypeCast is called when production typeCast is entered.
func (s *BaseDart2Listener) EnterTypeCast(ctx *TypeCastContext) {}

// ExitTypeCast is called when production typeCast is exited.
func (s *BaseDart2Listener) ExitTypeCast(ctx *TypeCastContext) {}

// EnterAsOperator is called when production asOperator is entered.
func (s *BaseDart2Listener) EnterAsOperator(ctx *AsOperatorContext) {}

// ExitAsOperator is called when production asOperator is exited.
func (s *BaseDart2Listener) ExitAsOperator(ctx *AsOperatorContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseDart2Listener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseDart2Listener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDart2Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDart2Listener) ExitStatement(ctx *StatementContext) {}

// EnterNonLabledStatment is called when production nonLabledStatment is entered.
func (s *BaseDart2Listener) EnterNonLabledStatment(ctx *NonLabledStatmentContext) {}

// ExitNonLabledStatment is called when production nonLabledStatment is exited.
func (s *BaseDart2Listener) ExitNonLabledStatment(ctx *NonLabledStatmentContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseDart2Listener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseDart2Listener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseDart2Listener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseDart2Listener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterLocalFunctionDeclaration is called when production localFunctionDeclaration is entered.
func (s *BaseDart2Listener) EnterLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) {}

// ExitLocalFunctionDeclaration is called when production localFunctionDeclaration is exited.
func (s *BaseDart2Listener) ExitLocalFunctionDeclaration(ctx *LocalFunctionDeclarationContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseDart2Listener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseDart2Listener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseDart2Listener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseDart2Listener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForLoopParts is called when production forLoopParts is entered.
func (s *BaseDart2Listener) EnterForLoopParts(ctx *ForLoopPartsContext) {}

// ExitForLoopParts is called when production forLoopParts is exited.
func (s *BaseDart2Listener) ExitForLoopParts(ctx *ForLoopPartsContext) {}

// EnterForInitializerStatement is called when production forInitializerStatement is entered.
func (s *BaseDart2Listener) EnterForInitializerStatement(ctx *ForInitializerStatementContext) {}

// ExitForInitializerStatement is called when production forInitializerStatement is exited.
func (s *BaseDart2Listener) ExitForInitializerStatement(ctx *ForInitializerStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseDart2Listener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseDart2Listener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseDart2Listener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseDart2Listener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseDart2Listener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseDart2Listener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseDart2Listener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseDart2Listener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseDart2Listener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseDart2Listener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterRethrowStatment is called when production rethrowStatment is entered.
func (s *BaseDart2Listener) EnterRethrowStatment(ctx *RethrowStatmentContext) {}

// ExitRethrowStatment is called when production rethrowStatment is exited.
func (s *BaseDart2Listener) ExitRethrowStatment(ctx *RethrowStatmentContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseDart2Listener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseDart2Listener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterOnPart is called when production onPart is entered.
func (s *BaseDart2Listener) EnterOnPart(ctx *OnPartContext) {}

// ExitOnPart is called when production onPart is exited.
func (s *BaseDart2Listener) ExitOnPart(ctx *OnPartContext) {}

// EnterCatchPart is called when production catchPart is entered.
func (s *BaseDart2Listener) EnterCatchPart(ctx *CatchPartContext) {}

// ExitCatchPart is called when production catchPart is exited.
func (s *BaseDart2Listener) ExitCatchPart(ctx *CatchPartContext) {}

// EnterFinallyPart is called when production finallyPart is entered.
func (s *BaseDart2Listener) EnterFinallyPart(ctx *FinallyPartContext) {}

// ExitFinallyPart is called when production finallyPart is exited.
func (s *BaseDart2Listener) ExitFinallyPart(ctx *FinallyPartContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseDart2Listener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseDart2Listener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDart2Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDart2Listener) ExitLabel(ctx *LabelContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseDart2Listener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseDart2Listener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseDart2Listener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseDart2Listener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterYieldStatement is called when production yieldStatement is entered.
func (s *BaseDart2Listener) EnterYieldStatement(ctx *YieldStatementContext) {}

// ExitYieldStatement is called when production yieldStatement is exited.
func (s *BaseDart2Listener) ExitYieldStatement(ctx *YieldStatementContext) {}

// EnterYieldEachStatement is called when production yieldEachStatement is entered.
func (s *BaseDart2Listener) EnterYieldEachStatement(ctx *YieldEachStatementContext) {}

// ExitYieldEachStatement is called when production yieldEachStatement is exited.
func (s *BaseDart2Listener) ExitYieldEachStatement(ctx *YieldEachStatementContext) {}

// EnterAssertStatement is called when production assertStatement is entered.
func (s *BaseDart2Listener) EnterAssertStatement(ctx *AssertStatementContext) {}

// ExitAssertStatement is called when production assertStatement is exited.
func (s *BaseDart2Listener) ExitAssertStatement(ctx *AssertStatementContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseDart2Listener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseDart2Listener) ExitAssertion(ctx *AssertionContext) {}

// EnterTopLevelDefinition is called when production topLevelDefinition is entered.
func (s *BaseDart2Listener) EnterTopLevelDefinition(ctx *TopLevelDefinitionContext) {}

// ExitTopLevelDefinition is called when production topLevelDefinition is exited.
func (s *BaseDart2Listener) ExitTopLevelDefinition(ctx *TopLevelDefinitionContext) {}

// EnterGetOrSet is called when production getOrSet is entered.
func (s *BaseDart2Listener) EnterGetOrSet(ctx *GetOrSetContext) {}

// ExitGetOrSet is called when production getOrSet is exited.
func (s *BaseDart2Listener) ExitGetOrSet(ctx *GetOrSetContext) {}

// EnterLibraryDefinition is called when production libraryDefinition is entered.
func (s *BaseDart2Listener) EnterLibraryDefinition(ctx *LibraryDefinitionContext) {}

// ExitLibraryDefinition is called when production libraryDefinition is exited.
func (s *BaseDart2Listener) ExitLibraryDefinition(ctx *LibraryDefinitionContext) {}

// EnterScriptTag is called when production scriptTag is entered.
func (s *BaseDart2Listener) EnterScriptTag(ctx *ScriptTagContext) {}

// ExitScriptTag is called when production scriptTag is exited.
func (s *BaseDart2Listener) ExitScriptTag(ctx *ScriptTagContext) {}

// EnterLibraryName is called when production libraryName is entered.
func (s *BaseDart2Listener) EnterLibraryName(ctx *LibraryNameContext) {}

// ExitLibraryName is called when production libraryName is exited.
func (s *BaseDart2Listener) ExitLibraryName(ctx *LibraryNameContext) {}

// EnterImportOrExport is called when production importOrExport is entered.
func (s *BaseDart2Listener) EnterImportOrExport(ctx *ImportOrExportContext) {}

// ExitImportOrExport is called when production importOrExport is exited.
func (s *BaseDart2Listener) ExitImportOrExport(ctx *ImportOrExportContext) {}

// EnterDottedIdentifierList is called when production dottedIdentifierList is entered.
func (s *BaseDart2Listener) EnterDottedIdentifierList(ctx *DottedIdentifierListContext) {}

// ExitDottedIdentifierList is called when production dottedIdentifierList is exited.
func (s *BaseDart2Listener) ExitDottedIdentifierList(ctx *DottedIdentifierListContext) {}

// EnterLibraryimport is called when production libraryimport is entered.
func (s *BaseDart2Listener) EnterLibraryimport(ctx *LibraryimportContext) {}

// ExitLibraryimport is called when production libraryimport is exited.
func (s *BaseDart2Listener) ExitLibraryimport(ctx *LibraryimportContext) {}

// EnterImportSpecification is called when production importSpecification is entered.
func (s *BaseDart2Listener) EnterImportSpecification(ctx *ImportSpecificationContext) {}

// ExitImportSpecification is called when production importSpecification is exited.
func (s *BaseDart2Listener) ExitImportSpecification(ctx *ImportSpecificationContext) {}

// EnterCombinator is called when production combinator is entered.
func (s *BaseDart2Listener) EnterCombinator(ctx *CombinatorContext) {}

// ExitCombinator is called when production combinator is exited.
func (s *BaseDart2Listener) ExitCombinator(ctx *CombinatorContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseDart2Listener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseDart2Listener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterLibraryExport is called when production libraryExport is entered.
func (s *BaseDart2Listener) EnterLibraryExport(ctx *LibraryExportContext) {}

// ExitLibraryExport is called when production libraryExport is exited.
func (s *BaseDart2Listener) ExitLibraryExport(ctx *LibraryExportContext) {}

// EnterPartDirective is called when production partDirective is entered.
func (s *BaseDart2Listener) EnterPartDirective(ctx *PartDirectiveContext) {}

// ExitPartDirective is called when production partDirective is exited.
func (s *BaseDart2Listener) ExitPartDirective(ctx *PartDirectiveContext) {}

// EnterPartHeader is called when production partHeader is entered.
func (s *BaseDart2Listener) EnterPartHeader(ctx *PartHeaderContext) {}

// ExitPartHeader is called when production partHeader is exited.
func (s *BaseDart2Listener) ExitPartHeader(ctx *PartHeaderContext) {}

// EnterPartDeclaration is called when production partDeclaration is entered.
func (s *BaseDart2Listener) EnterPartDeclaration(ctx *PartDeclarationContext) {}

// ExitPartDeclaration is called when production partDeclaration is exited.
func (s *BaseDart2Listener) ExitPartDeclaration(ctx *PartDeclarationContext) {}

// EnterUri is called when production uri is entered.
func (s *BaseDart2Listener) EnterUri(ctx *UriContext) {}

// ExitUri is called when production uri is exited.
func (s *BaseDart2Listener) ExitUri(ctx *UriContext) {}

// EnterConfigurableUri is called when production configurableUri is entered.
func (s *BaseDart2Listener) EnterConfigurableUri(ctx *ConfigurableUriContext) {}

// ExitConfigurableUri is called when production configurableUri is exited.
func (s *BaseDart2Listener) ExitConfigurableUri(ctx *ConfigurableUriContext) {}

// EnterConfigurationUri is called when production configurationUri is entered.
func (s *BaseDart2Listener) EnterConfigurationUri(ctx *ConfigurationUriContext) {}

// ExitConfigurationUri is called when production configurationUri is exited.
func (s *BaseDart2Listener) ExitConfigurationUri(ctx *ConfigurationUriContext) {}

// EnterUriTest is called when production uriTest is entered.
func (s *BaseDart2Listener) EnterUriTest(ctx *UriTestContext) {}

// ExitUriTest is called when production uriTest is exited.
func (s *BaseDart2Listener) ExitUriTest(ctx *UriTestContext) {}

// EnterDtype is called when production dtype is entered.
func (s *BaseDart2Listener) EnterDtype(ctx *DtypeContext) {}

// ExitDtype is called when production dtype is exited.
func (s *BaseDart2Listener) ExitDtype(ctx *DtypeContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseDart2Listener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseDart2Listener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseDart2Listener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseDart2Listener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseDart2Listener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseDart2Listener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *BaseDart2Listener) EnterTypeAlias(ctx *TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *BaseDart2Listener) ExitTypeAlias(ctx *TypeAliasContext) {}

// EnterTypeAliasBody is called when production typeAliasBody is entered.
func (s *BaseDart2Listener) EnterTypeAliasBody(ctx *TypeAliasBodyContext) {}

// ExitTypeAliasBody is called when production typeAliasBody is exited.
func (s *BaseDart2Listener) ExitTypeAliasBody(ctx *TypeAliasBodyContext) {}

// EnterFunctionTypeAlias is called when production functionTypeAlias is entered.
func (s *BaseDart2Listener) EnterFunctionTypeAlias(ctx *FunctionTypeAliasContext) {}

// ExitFunctionTypeAlias is called when production functionTypeAlias is exited.
func (s *BaseDart2Listener) ExitFunctionTypeAlias(ctx *FunctionTypeAliasContext) {}

// EnterFunctionPrefix is called when production functionPrefix is entered.
func (s *BaseDart2Listener) EnterFunctionPrefix(ctx *FunctionPrefixContext) {}

// ExitFunctionPrefix is called when production functionPrefix is exited.
func (s *BaseDart2Listener) ExitFunctionPrefix(ctx *FunctionPrefixContext) {}
