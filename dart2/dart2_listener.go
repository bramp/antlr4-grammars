// Code generated from Dart2.g4 by ANTLR 4.7.2. DO NOT EDIT.

package dart2 // Dart2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Dart2Listener is a complete listener for a parse tree produced by Dart2Parser.
type Dart2Listener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterDeclaredIdentifier is called when entering the declaredIdentifier production.
	EnterDeclaredIdentifier(c *DeclaredIdentifierContext)

	// EnterFinalConstVarOrType is called when entering the finalConstVarOrType production.
	EnterFinalConstVarOrType(c *FinalConstVarOrTypeContext)

	// EnterVarOrType is called when entering the varOrType production.
	EnterVarOrType(c *VarOrTypeContext)

	// EnterInitializedVariableDeclaration is called when entering the initializedVariableDeclaration production.
	EnterInitializedVariableDeclaration(c *InitializedVariableDeclarationContext)

	// EnterInitializedIdentifier is called when entering the initializedIdentifier production.
	EnterInitializedIdentifier(c *InitializedIdentifierContext)

	// EnterInitializedIdentifierList is called when entering the initializedIdentifierList production.
	EnterInitializedIdentifierList(c *InitializedIdentifierListContext)

	// EnterFunctionSignature is called when entering the functionSignature production.
	EnterFunctionSignature(c *FunctionSignatureContext)

	// EnterFormalParameterPart is called when entering the formalParameterPart production.
	EnterFormalParameterPart(c *FormalParameterPartContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterNormalFormalParameters is called when entering the normalFormalParameters production.
	EnterNormalFormalParameters(c *NormalFormalParametersContext)

	// EnterOptionalFormalParameters is called when entering the optionalFormalParameters production.
	EnterOptionalFormalParameters(c *OptionalFormalParametersContext)

	// EnterOptionalPositionalFormalParameters is called when entering the optionalPositionalFormalParameters production.
	EnterOptionalPositionalFormalParameters(c *OptionalPositionalFormalParametersContext)

	// EnterNamedFormalParameters is called when entering the namedFormalParameters production.
	EnterNamedFormalParameters(c *NamedFormalParametersContext)

	// EnterNormalFormalParameter is called when entering the normalFormalParameter production.
	EnterNormalFormalParameter(c *NormalFormalParameterContext)

	// EnterFunctionFormalParameter is called when entering the functionFormalParameter production.
	EnterFunctionFormalParameter(c *FunctionFormalParameterContext)

	// EnterSimpleFormalParameter is called when entering the simpleFormalParameter production.
	EnterSimpleFormalParameter(c *SimpleFormalParameterContext)

	// EnterFieldFormalParameter is called when entering the fieldFormalParameter production.
	EnterFieldFormalParameter(c *FieldFormalParameterContext)

	// EnterDefaultFormalParameter is called when entering the defaultFormalParameter production.
	EnterDefaultFormalParameter(c *DefaultFormalParameterContext)

	// EnterDefaultNamedParameter is called when entering the defaultNamedParameter production.
	EnterDefaultNamedParameter(c *DefaultNamedParameterContext)

	// EnterClassDefinition is called when entering the classDefinition production.
	EnterClassDefinition(c *ClassDefinitionContext)

	// EnterMixins is called when entering the mixins production.
	EnterMixins(c *MixinsContext)

	// EnterClassMemberDefinition is called when entering the classMemberDefinition production.
	EnterClassMemberDefinition(c *ClassMemberDefinitionContext)

	// EnterMethodSignature is called when entering the methodSignature production.
	EnterMethodSignature(c *MethodSignatureContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterStaticFinalDeclarationList is called when entering the staticFinalDeclarationList production.
	EnterStaticFinalDeclarationList(c *StaticFinalDeclarationListContext)

	// EnterStaticFinalDeclaration is called when entering the staticFinalDeclaration production.
	EnterStaticFinalDeclaration(c *StaticFinalDeclarationContext)

	// EnterOperatorSignature is called when entering the operatorSignature production.
	EnterOperatorSignature(c *OperatorSignatureContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterGetterSignature is called when entering the getterSignature production.
	EnterGetterSignature(c *GetterSignatureContext)

	// EnterSetterSignature is called when entering the setterSignature production.
	EnterSetterSignature(c *SetterSignatureContext)

	// EnterConstructorSignature is called when entering the constructorSignature production.
	EnterConstructorSignature(c *ConstructorSignatureContext)

	// EnterRedirection is called when entering the redirection production.
	EnterRedirection(c *RedirectionContext)

	// EnterInitializers is called when entering the initializers production.
	EnterInitializers(c *InitializersContext)

	// EnterInitializerListEntry is called when entering the initializerListEntry production.
	EnterInitializerListEntry(c *InitializerListEntryContext)

	// EnterFieldInitializer is called when entering the fieldInitializer production.
	EnterFieldInitializer(c *FieldInitializerContext)

	// EnterFactoryConstructorSignature is called when entering the factoryConstructorSignature production.
	EnterFactoryConstructorSignature(c *FactoryConstructorSignatureContext)

	// EnterRedirectingFactoryConstructorSignature is called when entering the redirectingFactoryConstructorSignature production.
	EnterRedirectingFactoryConstructorSignature(c *RedirectingFactoryConstructorSignatureContext)

	// EnterConstantConstructorSignature is called when entering the constantConstructorSignature production.
	EnterConstantConstructorSignature(c *ConstantConstructorSignatureContext)

	// EnterSuperclass is called when entering the superclass production.
	EnterSuperclass(c *SuperclassContext)

	// EnterInterfaces is called when entering the interfaces production.
	EnterInterfaces(c *InterfacesContext)

	// EnterMixinApplicationClass is called when entering the mixinApplicationClass production.
	EnterMixinApplicationClass(c *MixinApplicationClassContext)

	// EnterMixinApplication is called when entering the mixinApplication production.
	EnterMixinApplication(c *MixinApplicationContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterEnumEntry is called when entering the enumEntry production.
	EnterEnumEntry(c *EnumEntryContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterMetadata is called when entering the metadata production.
	EnterMetadata(c *MetadataContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionWithoutCascade is called when entering the expressionWithoutCascade production.
	EnterExpressionWithoutCascade(c *ExpressionWithoutCascadeContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterStringInterpolation is called when entering the stringInterpolation production.
	EnterStringInterpolation(c *StringInterpolationContext)

	// EnterSymbolLiteral is called when entering the symbolLiteral production.
	EnterSymbolLiteral(c *SymbolLiteralContext)

	// EnterListLiteral is called when entering the listLiteral production.
	EnterListLiteral(c *ListLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterMapLiteralEntry is called when entering the mapLiteralEntry production.
	EnterMapLiteralEntry(c *MapLiteralEntryContext)

	// EnterThrowExpression is called when entering the throwExpression production.
	EnterThrowExpression(c *ThrowExpressionContext)

	// EnterThrowExpressionWithoutCascade is called when entering the throwExpressionWithoutCascade production.
	EnterThrowExpressionWithoutCascade(c *ThrowExpressionWithoutCascadeContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterThisExpression is called when entering the thisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterNayaExpression is called when entering the nayaExpression production.
	EnterNayaExpression(c *NayaExpressionContext)

	// EnterConstObjectExpression is called when entering the constObjectExpression production.
	EnterConstObjectExpression(c *ConstObjectExpressionContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterNamedArgument is called when entering the namedArgument production.
	EnterNamedArgument(c *NamedArgumentContext)

	// EnterCascadeSection is called when entering the cascadeSection production.
	EnterCascadeSection(c *CascadeSectionContext)

	// EnterCascadeSelector is called when entering the cascadeSelector production.
	EnterCascadeSelector(c *CascadeSelectorContext)

	// EnterArgumentPart is called when entering the argumentPart production.
	EnterArgumentPart(c *ArgumentPartContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterCompoundAssignmentOperator is called when entering the compoundAssignmentOperator production.
	EnterCompoundAssignmentOperator(c *CompoundAssignmentOperatorContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterIfNullExpression is called when entering the ifNullExpression production.
	EnterIfNullExpression(c *IfNullExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterEqualityOperator is called when entering the equalityOperator production.
	EnterEqualityOperator(c *EqualityOperatorContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterRelationalOperator is called when entering the relationalOperator production.
	EnterRelationalOperator(c *RelationalOperatorContext)

	// EnterBitwiseOrExpression is called when entering the bitwiseOrExpression production.
	EnterBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// EnterBitwiseXorExpression is called when entering the bitwiseXorExpression production.
	EnterBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// EnterBitwiseAndExpression is called when entering the bitwiseAndExpression production.
	EnterBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// EnterBitwiseOperator is called when entering the bitwiseOperator production.
	EnterBitwiseOperator(c *BitwiseOperatorContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterShiftOperator is called when entering the shiftOperator production.
	EnterShiftOperator(c *ShiftOperatorContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterAdditiveOperator is called when entering the additiveOperator production.
	EnterAdditiveOperator(c *AdditiveOperatorContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterMultiplicativeOperator is called when entering the multiplicativeOperator production.
	EnterMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPrefixOperator is called when entering the prefixOperator production.
	EnterPrefixOperator(c *PrefixOperatorContext)

	// EnterMinusOperator is called when entering the minusOperator production.
	EnterMinusOperator(c *MinusOperatorContext)

	// EnterNegationOperator is called when entering the negationOperator production.
	EnterNegationOperator(c *NegationOperatorContext)

	// EnterTildeOperator is called when entering the tildeOperator production.
	EnterTildeOperator(c *TildeOperatorContext)

	// EnterAwaitExpression is called when entering the awaitExpression production.
	EnterAwaitExpression(c *AwaitExpressionContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostfixOperator is called when entering the postfixOperator production.
	EnterPostfixOperator(c *PostfixOperatorContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterIncrementOperator is called when entering the incrementOperator production.
	EnterIncrementOperator(c *IncrementOperatorContext)

	// EnterAssignableExpression is called when entering the assignableExpression production.
	EnterAssignableExpression(c *AssignableExpressionContext)

	// EnterUnconditionalAssignableSelector is called when entering the unconditionalAssignableSelector production.
	EnterUnconditionalAssignableSelector(c *UnconditionalAssignableSelectorContext)

	// EnterAssignableSelector is called when entering the assignableSelector production.
	EnterAssignableSelector(c *AssignableSelectorContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterQualified is called when entering the qualified production.
	EnterQualified(c *QualifiedContext)

	// EnterTypeTest is called when entering the typeTest production.
	EnterTypeTest(c *TypeTestContext)

	// EnterIsOperator is called when entering the isOperator production.
	EnterIsOperator(c *IsOperatorContext)

	// EnterTypeCast is called when entering the typeCast production.
	EnterTypeCast(c *TypeCastContext)

	// EnterAsOperator is called when entering the asOperator production.
	EnterAsOperator(c *AsOperatorContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNonLabledStatment is called when entering the nonLabledStatment production.
	EnterNonLabledStatment(c *NonLabledStatmentContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterLocalFunctionDeclaration is called when entering the localFunctionDeclaration production.
	EnterLocalFunctionDeclaration(c *LocalFunctionDeclarationContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForLoopParts is called when entering the forLoopParts production.
	EnterForLoopParts(c *ForLoopPartsContext)

	// EnterForInitializerStatement is called when entering the forInitializerStatement production.
	EnterForInitializerStatement(c *ForInitializerStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterRethrowStatment is called when entering the rethrowStatment production.
	EnterRethrowStatment(c *RethrowStatmentContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterOnPart is called when entering the onPart production.
	EnterOnPart(c *OnPartContext)

	// EnterCatchPart is called when entering the catchPart production.
	EnterCatchPart(c *CatchPartContext)

	// EnterFinallyPart is called when entering the finallyPart production.
	EnterFinallyPart(c *FinallyPartContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterYieldStatement is called when entering the yieldStatement production.
	EnterYieldStatement(c *YieldStatementContext)

	// EnterYieldEachStatement is called when entering the yieldEachStatement production.
	EnterYieldEachStatement(c *YieldEachStatementContext)

	// EnterAssertStatement is called when entering the assertStatement production.
	EnterAssertStatement(c *AssertStatementContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterTopLevelDefinition is called when entering the topLevelDefinition production.
	EnterTopLevelDefinition(c *TopLevelDefinitionContext)

	// EnterGetOrSet is called when entering the getOrSet production.
	EnterGetOrSet(c *GetOrSetContext)

	// EnterLibraryDefinition is called when entering the libraryDefinition production.
	EnterLibraryDefinition(c *LibraryDefinitionContext)

	// EnterScriptTag is called when entering the scriptTag production.
	EnterScriptTag(c *ScriptTagContext)

	// EnterLibraryName is called when entering the libraryName production.
	EnterLibraryName(c *LibraryNameContext)

	// EnterImportOrExport is called when entering the importOrExport production.
	EnterImportOrExport(c *ImportOrExportContext)

	// EnterDottedIdentifierList is called when entering the dottedIdentifierList production.
	EnterDottedIdentifierList(c *DottedIdentifierListContext)

	// EnterLibraryimport is called when entering the libraryimport production.
	EnterLibraryimport(c *LibraryimportContext)

	// EnterImportSpecification is called when entering the importSpecification production.
	EnterImportSpecification(c *ImportSpecificationContext)

	// EnterCombinator is called when entering the combinator production.
	EnterCombinator(c *CombinatorContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterLibraryExport is called when entering the libraryExport production.
	EnterLibraryExport(c *LibraryExportContext)

	// EnterPartDirective is called when entering the partDirective production.
	EnterPartDirective(c *PartDirectiveContext)

	// EnterPartHeader is called when entering the partHeader production.
	EnterPartHeader(c *PartHeaderContext)

	// EnterPartDeclaration is called when entering the partDeclaration production.
	EnterPartDeclaration(c *PartDeclarationContext)

	// EnterUri is called when entering the uri production.
	EnterUri(c *UriContext)

	// EnterConfigurableUri is called when entering the configurableUri production.
	EnterConfigurableUri(c *ConfigurableUriContext)

	// EnterConfigurationUri is called when entering the configurationUri production.
	EnterConfigurationUri(c *ConfigurationUriContext)

	// EnterUriTest is called when entering the uriTest production.
	EnterUriTest(c *UriTestContext)

	// EnterDtype is called when entering the dtype production.
	EnterDtype(c *DtypeContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeAlias is called when entering the typeAlias production.
	EnterTypeAlias(c *TypeAliasContext)

	// EnterTypeAliasBody is called when entering the typeAliasBody production.
	EnterTypeAliasBody(c *TypeAliasBodyContext)

	// EnterFunctionTypeAlias is called when entering the functionTypeAlias production.
	EnterFunctionTypeAlias(c *FunctionTypeAliasContext)

	// EnterFunctionPrefix is called when entering the functionPrefix production.
	EnterFunctionPrefix(c *FunctionPrefixContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitDeclaredIdentifier is called when exiting the declaredIdentifier production.
	ExitDeclaredIdentifier(c *DeclaredIdentifierContext)

	// ExitFinalConstVarOrType is called when exiting the finalConstVarOrType production.
	ExitFinalConstVarOrType(c *FinalConstVarOrTypeContext)

	// ExitVarOrType is called when exiting the varOrType production.
	ExitVarOrType(c *VarOrTypeContext)

	// ExitInitializedVariableDeclaration is called when exiting the initializedVariableDeclaration production.
	ExitInitializedVariableDeclaration(c *InitializedVariableDeclarationContext)

	// ExitInitializedIdentifier is called when exiting the initializedIdentifier production.
	ExitInitializedIdentifier(c *InitializedIdentifierContext)

	// ExitInitializedIdentifierList is called when exiting the initializedIdentifierList production.
	ExitInitializedIdentifierList(c *InitializedIdentifierListContext)

	// ExitFunctionSignature is called when exiting the functionSignature production.
	ExitFunctionSignature(c *FunctionSignatureContext)

	// ExitFormalParameterPart is called when exiting the formalParameterPart production.
	ExitFormalParameterPart(c *FormalParameterPartContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitNormalFormalParameters is called when exiting the normalFormalParameters production.
	ExitNormalFormalParameters(c *NormalFormalParametersContext)

	// ExitOptionalFormalParameters is called when exiting the optionalFormalParameters production.
	ExitOptionalFormalParameters(c *OptionalFormalParametersContext)

	// ExitOptionalPositionalFormalParameters is called when exiting the optionalPositionalFormalParameters production.
	ExitOptionalPositionalFormalParameters(c *OptionalPositionalFormalParametersContext)

	// ExitNamedFormalParameters is called when exiting the namedFormalParameters production.
	ExitNamedFormalParameters(c *NamedFormalParametersContext)

	// ExitNormalFormalParameter is called when exiting the normalFormalParameter production.
	ExitNormalFormalParameter(c *NormalFormalParameterContext)

	// ExitFunctionFormalParameter is called when exiting the functionFormalParameter production.
	ExitFunctionFormalParameter(c *FunctionFormalParameterContext)

	// ExitSimpleFormalParameter is called when exiting the simpleFormalParameter production.
	ExitSimpleFormalParameter(c *SimpleFormalParameterContext)

	// ExitFieldFormalParameter is called when exiting the fieldFormalParameter production.
	ExitFieldFormalParameter(c *FieldFormalParameterContext)

	// ExitDefaultFormalParameter is called when exiting the defaultFormalParameter production.
	ExitDefaultFormalParameter(c *DefaultFormalParameterContext)

	// ExitDefaultNamedParameter is called when exiting the defaultNamedParameter production.
	ExitDefaultNamedParameter(c *DefaultNamedParameterContext)

	// ExitClassDefinition is called when exiting the classDefinition production.
	ExitClassDefinition(c *ClassDefinitionContext)

	// ExitMixins is called when exiting the mixins production.
	ExitMixins(c *MixinsContext)

	// ExitClassMemberDefinition is called when exiting the classMemberDefinition production.
	ExitClassMemberDefinition(c *ClassMemberDefinitionContext)

	// ExitMethodSignature is called when exiting the methodSignature production.
	ExitMethodSignature(c *MethodSignatureContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitStaticFinalDeclarationList is called when exiting the staticFinalDeclarationList production.
	ExitStaticFinalDeclarationList(c *StaticFinalDeclarationListContext)

	// ExitStaticFinalDeclaration is called when exiting the staticFinalDeclaration production.
	ExitStaticFinalDeclaration(c *StaticFinalDeclarationContext)

	// ExitOperatorSignature is called when exiting the operatorSignature production.
	ExitOperatorSignature(c *OperatorSignatureContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitGetterSignature is called when exiting the getterSignature production.
	ExitGetterSignature(c *GetterSignatureContext)

	// ExitSetterSignature is called when exiting the setterSignature production.
	ExitSetterSignature(c *SetterSignatureContext)

	// ExitConstructorSignature is called when exiting the constructorSignature production.
	ExitConstructorSignature(c *ConstructorSignatureContext)

	// ExitRedirection is called when exiting the redirection production.
	ExitRedirection(c *RedirectionContext)

	// ExitInitializers is called when exiting the initializers production.
	ExitInitializers(c *InitializersContext)

	// ExitInitializerListEntry is called when exiting the initializerListEntry production.
	ExitInitializerListEntry(c *InitializerListEntryContext)

	// ExitFieldInitializer is called when exiting the fieldInitializer production.
	ExitFieldInitializer(c *FieldInitializerContext)

	// ExitFactoryConstructorSignature is called when exiting the factoryConstructorSignature production.
	ExitFactoryConstructorSignature(c *FactoryConstructorSignatureContext)

	// ExitRedirectingFactoryConstructorSignature is called when exiting the redirectingFactoryConstructorSignature production.
	ExitRedirectingFactoryConstructorSignature(c *RedirectingFactoryConstructorSignatureContext)

	// ExitConstantConstructorSignature is called when exiting the constantConstructorSignature production.
	ExitConstantConstructorSignature(c *ConstantConstructorSignatureContext)

	// ExitSuperclass is called when exiting the superclass production.
	ExitSuperclass(c *SuperclassContext)

	// ExitInterfaces is called when exiting the interfaces production.
	ExitInterfaces(c *InterfacesContext)

	// ExitMixinApplicationClass is called when exiting the mixinApplicationClass production.
	ExitMixinApplicationClass(c *MixinApplicationClassContext)

	// ExitMixinApplication is called when exiting the mixinApplication production.
	ExitMixinApplication(c *MixinApplicationContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitEnumEntry is called when exiting the enumEntry production.
	ExitEnumEntry(c *EnumEntryContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitMetadata is called when exiting the metadata production.
	ExitMetadata(c *MetadataContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionWithoutCascade is called when exiting the expressionWithoutCascade production.
	ExitExpressionWithoutCascade(c *ExpressionWithoutCascadeContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitStringInterpolation is called when exiting the stringInterpolation production.
	ExitStringInterpolation(c *StringInterpolationContext)

	// ExitSymbolLiteral is called when exiting the symbolLiteral production.
	ExitSymbolLiteral(c *SymbolLiteralContext)

	// ExitListLiteral is called when exiting the listLiteral production.
	ExitListLiteral(c *ListLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitMapLiteralEntry is called when exiting the mapLiteralEntry production.
	ExitMapLiteralEntry(c *MapLiteralEntryContext)

	// ExitThrowExpression is called when exiting the throwExpression production.
	ExitThrowExpression(c *ThrowExpressionContext)

	// ExitThrowExpressionWithoutCascade is called when exiting the throwExpressionWithoutCascade production.
	ExitThrowExpressionWithoutCascade(c *ThrowExpressionWithoutCascadeContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitThisExpression is called when exiting the thisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitNayaExpression is called when exiting the nayaExpression production.
	ExitNayaExpression(c *NayaExpressionContext)

	// ExitConstObjectExpression is called when exiting the constObjectExpression production.
	ExitConstObjectExpression(c *ConstObjectExpressionContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitNamedArgument is called when exiting the namedArgument production.
	ExitNamedArgument(c *NamedArgumentContext)

	// ExitCascadeSection is called when exiting the cascadeSection production.
	ExitCascadeSection(c *CascadeSectionContext)

	// ExitCascadeSelector is called when exiting the cascadeSelector production.
	ExitCascadeSelector(c *CascadeSelectorContext)

	// ExitArgumentPart is called when exiting the argumentPart production.
	ExitArgumentPart(c *ArgumentPartContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitCompoundAssignmentOperator is called when exiting the compoundAssignmentOperator production.
	ExitCompoundAssignmentOperator(c *CompoundAssignmentOperatorContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitIfNullExpression is called when exiting the ifNullExpression production.
	ExitIfNullExpression(c *IfNullExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitEqualityOperator is called when exiting the equalityOperator production.
	ExitEqualityOperator(c *EqualityOperatorContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitRelationalOperator is called when exiting the relationalOperator production.
	ExitRelationalOperator(c *RelationalOperatorContext)

	// ExitBitwiseOrExpression is called when exiting the bitwiseOrExpression production.
	ExitBitwiseOrExpression(c *BitwiseOrExpressionContext)

	// ExitBitwiseXorExpression is called when exiting the bitwiseXorExpression production.
	ExitBitwiseXorExpression(c *BitwiseXorExpressionContext)

	// ExitBitwiseAndExpression is called when exiting the bitwiseAndExpression production.
	ExitBitwiseAndExpression(c *BitwiseAndExpressionContext)

	// ExitBitwiseOperator is called when exiting the bitwiseOperator production.
	ExitBitwiseOperator(c *BitwiseOperatorContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitShiftOperator is called when exiting the shiftOperator production.
	ExitShiftOperator(c *ShiftOperatorContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitAdditiveOperator is called when exiting the additiveOperator production.
	ExitAdditiveOperator(c *AdditiveOperatorContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitMultiplicativeOperator is called when exiting the multiplicativeOperator production.
	ExitMultiplicativeOperator(c *MultiplicativeOperatorContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPrefixOperator is called when exiting the prefixOperator production.
	ExitPrefixOperator(c *PrefixOperatorContext)

	// ExitMinusOperator is called when exiting the minusOperator production.
	ExitMinusOperator(c *MinusOperatorContext)

	// ExitNegationOperator is called when exiting the negationOperator production.
	ExitNegationOperator(c *NegationOperatorContext)

	// ExitTildeOperator is called when exiting the tildeOperator production.
	ExitTildeOperator(c *TildeOperatorContext)

	// ExitAwaitExpression is called when exiting the awaitExpression production.
	ExitAwaitExpression(c *AwaitExpressionContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostfixOperator is called when exiting the postfixOperator production.
	ExitPostfixOperator(c *PostfixOperatorContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitIncrementOperator is called when exiting the incrementOperator production.
	ExitIncrementOperator(c *IncrementOperatorContext)

	// ExitAssignableExpression is called when exiting the assignableExpression production.
	ExitAssignableExpression(c *AssignableExpressionContext)

	// ExitUnconditionalAssignableSelector is called when exiting the unconditionalAssignableSelector production.
	ExitUnconditionalAssignableSelector(c *UnconditionalAssignableSelectorContext)

	// ExitAssignableSelector is called when exiting the assignableSelector production.
	ExitAssignableSelector(c *AssignableSelectorContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitQualified is called when exiting the qualified production.
	ExitQualified(c *QualifiedContext)

	// ExitTypeTest is called when exiting the typeTest production.
	ExitTypeTest(c *TypeTestContext)

	// ExitIsOperator is called when exiting the isOperator production.
	ExitIsOperator(c *IsOperatorContext)

	// ExitTypeCast is called when exiting the typeCast production.
	ExitTypeCast(c *TypeCastContext)

	// ExitAsOperator is called when exiting the asOperator production.
	ExitAsOperator(c *AsOperatorContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNonLabledStatment is called when exiting the nonLabledStatment production.
	ExitNonLabledStatment(c *NonLabledStatmentContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitLocalFunctionDeclaration is called when exiting the localFunctionDeclaration production.
	ExitLocalFunctionDeclaration(c *LocalFunctionDeclarationContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForLoopParts is called when exiting the forLoopParts production.
	ExitForLoopParts(c *ForLoopPartsContext)

	// ExitForInitializerStatement is called when exiting the forInitializerStatement production.
	ExitForInitializerStatement(c *ForInitializerStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitRethrowStatment is called when exiting the rethrowStatment production.
	ExitRethrowStatment(c *RethrowStatmentContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitOnPart is called when exiting the onPart production.
	ExitOnPart(c *OnPartContext)

	// ExitCatchPart is called when exiting the catchPart production.
	ExitCatchPart(c *CatchPartContext)

	// ExitFinallyPart is called when exiting the finallyPart production.
	ExitFinallyPart(c *FinallyPartContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitYieldStatement is called when exiting the yieldStatement production.
	ExitYieldStatement(c *YieldStatementContext)

	// ExitYieldEachStatement is called when exiting the yieldEachStatement production.
	ExitYieldEachStatement(c *YieldEachStatementContext)

	// ExitAssertStatement is called when exiting the assertStatement production.
	ExitAssertStatement(c *AssertStatementContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitTopLevelDefinition is called when exiting the topLevelDefinition production.
	ExitTopLevelDefinition(c *TopLevelDefinitionContext)

	// ExitGetOrSet is called when exiting the getOrSet production.
	ExitGetOrSet(c *GetOrSetContext)

	// ExitLibraryDefinition is called when exiting the libraryDefinition production.
	ExitLibraryDefinition(c *LibraryDefinitionContext)

	// ExitScriptTag is called when exiting the scriptTag production.
	ExitScriptTag(c *ScriptTagContext)

	// ExitLibraryName is called when exiting the libraryName production.
	ExitLibraryName(c *LibraryNameContext)

	// ExitImportOrExport is called when exiting the importOrExport production.
	ExitImportOrExport(c *ImportOrExportContext)

	// ExitDottedIdentifierList is called when exiting the dottedIdentifierList production.
	ExitDottedIdentifierList(c *DottedIdentifierListContext)

	// ExitLibraryimport is called when exiting the libraryimport production.
	ExitLibraryimport(c *LibraryimportContext)

	// ExitImportSpecification is called when exiting the importSpecification production.
	ExitImportSpecification(c *ImportSpecificationContext)

	// ExitCombinator is called when exiting the combinator production.
	ExitCombinator(c *CombinatorContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitLibraryExport is called when exiting the libraryExport production.
	ExitLibraryExport(c *LibraryExportContext)

	// ExitPartDirective is called when exiting the partDirective production.
	ExitPartDirective(c *PartDirectiveContext)

	// ExitPartHeader is called when exiting the partHeader production.
	ExitPartHeader(c *PartHeaderContext)

	// ExitPartDeclaration is called when exiting the partDeclaration production.
	ExitPartDeclaration(c *PartDeclarationContext)

	// ExitUri is called when exiting the uri production.
	ExitUri(c *UriContext)

	// ExitConfigurableUri is called when exiting the configurableUri production.
	ExitConfigurableUri(c *ConfigurableUriContext)

	// ExitConfigurationUri is called when exiting the configurationUri production.
	ExitConfigurationUri(c *ConfigurationUriContext)

	// ExitUriTest is called when exiting the uriTest production.
	ExitUriTest(c *UriTestContext)

	// ExitDtype is called when exiting the dtype production.
	ExitDtype(c *DtypeContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeAlias is called when exiting the typeAlias production.
	ExitTypeAlias(c *TypeAliasContext)

	// ExitTypeAliasBody is called when exiting the typeAliasBody production.
	ExitTypeAliasBody(c *TypeAliasBodyContext)

	// ExitFunctionTypeAlias is called when exiting the functionTypeAlias production.
	ExitFunctionTypeAlias(c *FunctionTypeAliasContext)

	// ExitFunctionPrefix is called when exiting the functionPrefix production.
	ExitFunctionPrefix(c *FunctionPrefixContext)
}
