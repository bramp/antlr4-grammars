// Code generated from Java8Parser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package java8 // Java8Parser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Java8ParserListener is a complete listener for a parse tree produced by Java8Parser.
type Java8ParserListener interface {
	antlr.ParseTreeListener

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterNumericType is called when entering the numericType production.
	EnterNumericType(c *NumericTypeContext)

	// EnterIntegralType is called when entering the integralType production.
	EnterIntegralType(c *IntegralTypeContext)

	// EnterFloatingPointType is called when entering the floatingPointType production.
	EnterFloatingPointType(c *FloatingPointTypeContext)

	// EnterReferenceType is called when entering the referenceType production.
	EnterReferenceType(c *ReferenceTypeContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterClassType_lf_classOrInterfaceType is called when entering the classType_lf_classOrInterfaceType production.
	EnterClassType_lf_classOrInterfaceType(c *ClassType_lf_classOrInterfaceTypeContext)

	// EnterClassType_lfno_classOrInterfaceType is called when entering the classType_lfno_classOrInterfaceType production.
	EnterClassType_lfno_classOrInterfaceType(c *ClassType_lfno_classOrInterfaceTypeContext)

	// EnterInterfaceType is called when entering the interfaceType production.
	EnterInterfaceType(c *InterfaceTypeContext)

	// EnterInterfaceType_lf_classOrInterfaceType is called when entering the interfaceType_lf_classOrInterfaceType production.
	EnterInterfaceType_lf_classOrInterfaceType(c *InterfaceType_lf_classOrInterfaceTypeContext)

	// EnterInterfaceType_lfno_classOrInterfaceType is called when entering the interfaceType_lfno_classOrInterfaceType production.
	EnterInterfaceType_lfno_classOrInterfaceType(c *InterfaceType_lfno_classOrInterfaceTypeContext)

	// EnterTypeVariable is called when entering the typeVariable production.
	EnterTypeVariable(c *TypeVariableContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterDims is called when entering the dims production.
	EnterDims(c *DimsContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterTypeParameterModifier is called when entering the typeParameterModifier production.
	EnterTypeParameterModifier(c *TypeParameterModifierContext)

	// EnterTypeBound is called when entering the typeBound production.
	EnterTypeBound(c *TypeBoundContext)

	// EnterAdditionalBound is called when entering the additionalBound production.
	EnterAdditionalBound(c *AdditionalBoundContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgumentList is called when entering the typeArgumentList production.
	EnterTypeArgumentList(c *TypeArgumentListContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterWildcard is called when entering the wildcard production.
	EnterWildcard(c *WildcardContext)

	// EnterWildcardBounds is called when entering the wildcardBounds production.
	EnterWildcardBounds(c *WildcardBoundsContext)

	// EnterPackageName is called when entering the packageName production.
	EnterPackageName(c *PackageNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterPackageOrTypeName is called when entering the packageOrTypeName production.
	EnterPackageOrTypeName(c *PackageOrTypeNameContext)

	// EnterExpressionName is called when entering the expressionName production.
	EnterExpressionName(c *ExpressionNameContext)

	// EnterMethodName is called when entering the methodName production.
	EnterMethodName(c *MethodNameContext)

	// EnterAmbiguousName is called when entering the ambiguousName production.
	EnterAmbiguousName(c *AmbiguousNameContext)

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPackageDeclaration is called when entering the packageDeclaration production.
	EnterPackageDeclaration(c *PackageDeclarationContext)

	// EnterPackageModifier is called when entering the packageModifier production.
	EnterPackageModifier(c *PackageModifierContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterSingleTypeImportDeclaration is called when entering the singleTypeImportDeclaration production.
	EnterSingleTypeImportDeclaration(c *SingleTypeImportDeclarationContext)

	// EnterTypeImportOnDemandDeclaration is called when entering the typeImportOnDemandDeclaration production.
	EnterTypeImportOnDemandDeclaration(c *TypeImportOnDemandDeclarationContext)

	// EnterSingleStaticImportDeclaration is called when entering the singleStaticImportDeclaration production.
	EnterSingleStaticImportDeclaration(c *SingleStaticImportDeclarationContext)

	// EnterStaticImportOnDemandDeclaration is called when entering the staticImportOnDemandDeclaration production.
	EnterStaticImportOnDemandDeclaration(c *StaticImportOnDemandDeclarationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterNormalClassDeclaration is called when entering the normalClassDeclaration production.
	EnterNormalClassDeclaration(c *NormalClassDeclarationContext)

	// EnterClassModifier is called when entering the classModifier production.
	EnterClassModifier(c *ClassModifierContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameterList is called when entering the typeParameterList production.
	EnterTypeParameterList(c *TypeParameterListContext)

	// EnterSuperclass is called when entering the superclass production.
	EnterSuperclass(c *SuperclassContext)

	// EnterSuperinterfaces is called when entering the superinterfaces production.
	EnterSuperinterfaces(c *SuperinterfacesContext)

	// EnterInterfaceTypeList is called when entering the interfaceTypeList production.
	EnterInterfaceTypeList(c *InterfaceTypeListContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterClassMemberDeclaration is called when entering the classMemberDeclaration production.
	EnterClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterFieldModifier is called when entering the fieldModifier production.
	EnterFieldModifier(c *FieldModifierContext)

	// EnterVariableDeclaratorList is called when entering the variableDeclaratorList production.
	EnterVariableDeclaratorList(c *VariableDeclaratorListContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterUnannType is called when entering the unannType production.
	EnterUnannType(c *UnannTypeContext)

	// EnterUnannPrimitiveType is called when entering the unannPrimitiveType production.
	EnterUnannPrimitiveType(c *UnannPrimitiveTypeContext)

	// EnterUnannReferenceType is called when entering the unannReferenceType production.
	EnterUnannReferenceType(c *UnannReferenceTypeContext)

	// EnterUnannClassOrInterfaceType is called when entering the unannClassOrInterfaceType production.
	EnterUnannClassOrInterfaceType(c *UnannClassOrInterfaceTypeContext)

	// EnterUnannClassType is called when entering the unannClassType production.
	EnterUnannClassType(c *UnannClassTypeContext)

	// EnterUnannClassType_lf_unannClassOrInterfaceType is called when entering the unannClassType_lf_unannClassOrInterfaceType production.
	EnterUnannClassType_lf_unannClassOrInterfaceType(c *UnannClassType_lf_unannClassOrInterfaceTypeContext)

	// EnterUnannClassType_lfno_unannClassOrInterfaceType is called when entering the unannClassType_lfno_unannClassOrInterfaceType production.
	EnterUnannClassType_lfno_unannClassOrInterfaceType(c *UnannClassType_lfno_unannClassOrInterfaceTypeContext)

	// EnterUnannInterfaceType is called when entering the unannInterfaceType production.
	EnterUnannInterfaceType(c *UnannInterfaceTypeContext)

	// EnterUnannInterfaceType_lf_unannClassOrInterfaceType is called when entering the unannInterfaceType_lf_unannClassOrInterfaceType production.
	EnterUnannInterfaceType_lf_unannClassOrInterfaceType(c *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext)

	// EnterUnannInterfaceType_lfno_unannClassOrInterfaceType is called when entering the unannInterfaceType_lfno_unannClassOrInterfaceType production.
	EnterUnannInterfaceType_lfno_unannClassOrInterfaceType(c *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext)

	// EnterUnannTypeVariable is called when entering the unannTypeVariable production.
	EnterUnannTypeVariable(c *UnannTypeVariableContext)

	// EnterUnannArrayType is called when entering the unannArrayType production.
	EnterUnannArrayType(c *UnannArrayTypeContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterMethodModifier is called when entering the methodModifier production.
	EnterMethodModifier(c *MethodModifierContext)

	// EnterMethodHeader is called when entering the methodHeader production.
	EnterMethodHeader(c *MethodHeaderContext)

	// EnterResult is called when entering the result production.
	EnterResult(c *ResultContext)

	// EnterMethodDeclarator is called when entering the methodDeclarator production.
	EnterMethodDeclarator(c *MethodDeclaratorContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterLastFormalParameter is called when entering the lastFormalParameter production.
	EnterLastFormalParameter(c *LastFormalParameterContext)

	// EnterReceiverParameter is called when entering the receiverParameter production.
	EnterReceiverParameter(c *ReceiverParameterContext)

	// EnterThrows_ is called when entering the throws_ production.
	EnterThrows_(c *Throws_Context)

	// EnterExceptionTypeList is called when entering the exceptionTypeList production.
	EnterExceptionTypeList(c *ExceptionTypeListContext)

	// EnterExceptionType is called when entering the exceptionType production.
	EnterExceptionType(c *ExceptionTypeContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterInstanceInitializer is called when entering the instanceInitializer production.
	EnterInstanceInitializer(c *InstanceInitializerContext)

	// EnterStaticInitializer is called when entering the staticInitializer production.
	EnterStaticInitializer(c *StaticInitializerContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterConstructorModifier is called when entering the constructorModifier production.
	EnterConstructorModifier(c *ConstructorModifierContext)

	// EnterConstructorDeclarator is called when entering the constructorDeclarator production.
	EnterConstructorDeclarator(c *ConstructorDeclaratorContext)

	// EnterSimpleTypeName is called when entering the simpleTypeName production.
	EnterSimpleTypeName(c *SimpleTypeNameContext)

	// EnterConstructorBody is called when entering the constructorBody production.
	EnterConstructorBody(c *ConstructorBodyContext)

	// EnterExplicitConstructorInvocation is called when entering the explicitConstructorInvocation production.
	EnterExplicitConstructorInvocation(c *ExplicitConstructorInvocationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumConstantList is called when entering the enumConstantList production.
	EnterEnumConstantList(c *EnumConstantListContext)

	// EnterEnumConstant is called when entering the enumConstant production.
	EnterEnumConstant(c *EnumConstantContext)

	// EnterEnumConstantModifier is called when entering the enumConstantModifier production.
	EnterEnumConstantModifier(c *EnumConstantModifierContext)

	// EnterEnumBodyDeclarations is called when entering the enumBodyDeclarations production.
	EnterEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterNormalInterfaceDeclaration is called when entering the normalInterfaceDeclaration production.
	EnterNormalInterfaceDeclaration(c *NormalInterfaceDeclarationContext)

	// EnterInterfaceModifier is called when entering the interfaceModifier production.
	EnterInterfaceModifier(c *InterfaceModifierContext)

	// EnterExtendsInterfaces is called when entering the extendsInterfaces production.
	EnterExtendsInterfaces(c *ExtendsInterfacesContext)

	// EnterInterfaceBody is called when entering the interfaceBody production.
	EnterInterfaceBody(c *InterfaceBodyContext)

	// EnterInterfaceMemberDeclaration is called when entering the interfaceMemberDeclaration production.
	EnterInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// EnterConstantDeclaration is called when entering the constantDeclaration production.
	EnterConstantDeclaration(c *ConstantDeclarationContext)

	// EnterConstantModifier is called when entering the constantModifier production.
	EnterConstantModifier(c *ConstantModifierContext)

	// EnterInterfaceMethodDeclaration is called when entering the interfaceMethodDeclaration production.
	EnterInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// EnterInterfaceMethodModifier is called when entering the interfaceMethodModifier production.
	EnterInterfaceMethodModifier(c *InterfaceMethodModifierContext)

	// EnterAnnotationTypeDeclaration is called when entering the annotationTypeDeclaration production.
	EnterAnnotationTypeDeclaration(c *AnnotationTypeDeclarationContext)

	// EnterAnnotationTypeBody is called when entering the annotationTypeBody production.
	EnterAnnotationTypeBody(c *AnnotationTypeBodyContext)

	// EnterAnnotationTypeMemberDeclaration is called when entering the annotationTypeMemberDeclaration production.
	EnterAnnotationTypeMemberDeclaration(c *AnnotationTypeMemberDeclarationContext)

	// EnterAnnotationTypeElementDeclaration is called when entering the annotationTypeElementDeclaration production.
	EnterAnnotationTypeElementDeclaration(c *AnnotationTypeElementDeclarationContext)

	// EnterAnnotationTypeElementModifier is called when entering the annotationTypeElementModifier production.
	EnterAnnotationTypeElementModifier(c *AnnotationTypeElementModifierContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterNormalAnnotation is called when entering the normalAnnotation production.
	EnterNormalAnnotation(c *NormalAnnotationContext)

	// EnterElementValuePairList is called when entering the elementValuePairList production.
	EnterElementValuePairList(c *ElementValuePairListContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterElementValue is called when entering the elementValue production.
	EnterElementValue(c *ElementValueContext)

	// EnterElementValueArrayInitializer is called when entering the elementValueArrayInitializer production.
	EnterElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// EnterElementValueList is called when entering the elementValueList production.
	EnterElementValueList(c *ElementValueListContext)

	// EnterMarkerAnnotation is called when entering the markerAnnotation production.
	EnterMarkerAnnotation(c *MarkerAnnotationContext)

	// EnterSingleElementAnnotation is called when entering the singleElementAnnotation production.
	EnterSingleElementAnnotation(c *SingleElementAnnotationContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterVariableInitializerList is called when entering the variableInitializerList production.
	EnterVariableInitializerList(c *VariableInitializerListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatements is called when entering the blockStatements production.
	EnterBlockStatements(c *BlockStatementsContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclarationStatement is called when entering the localVariableDeclarationStatement production.
	EnterLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatementNoShortIf is called when entering the statementNoShortIf production.
	EnterStatementNoShortIf(c *StatementNoShortIfContext)

	// EnterStatementWithoutTrailingSubstatement is called when entering the statementWithoutTrailingSubstatement production.
	EnterStatementWithoutTrailingSubstatement(c *StatementWithoutTrailingSubstatementContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterLabeledStatement is called when entering the labeledStatement production.
	EnterLabeledStatement(c *LabeledStatementContext)

	// EnterLabeledStatementNoShortIf is called when entering the labeledStatementNoShortIf production.
	EnterLabeledStatementNoShortIf(c *LabeledStatementNoShortIfContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterStatementExpression is called when entering the statementExpression production.
	EnterStatementExpression(c *StatementExpressionContext)

	// EnterIfThenStatement is called when entering the ifThenStatement production.
	EnterIfThenStatement(c *IfThenStatementContext)

	// EnterIfThenElseStatement is called when entering the ifThenElseStatement production.
	EnterIfThenElseStatement(c *IfThenElseStatementContext)

	// EnterIfThenElseStatementNoShortIf is called when entering the ifThenElseStatementNoShortIf production.
	EnterIfThenElseStatementNoShortIf(c *IfThenElseStatementNoShortIfContext)

	// EnterAssertStatement is called when entering the assertStatement production.
	EnterAssertStatement(c *AssertStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabels is called when entering the switchLabels production.
	EnterSwitchLabels(c *SwitchLabelsContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterEnumConstantName is called when entering the enumConstantName production.
	EnterEnumConstantName(c *EnumConstantNameContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterWhileStatementNoShortIf is called when entering the whileStatementNoShortIf production.
	EnterWhileStatementNoShortIf(c *WhileStatementNoShortIfContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForStatementNoShortIf is called when entering the forStatementNoShortIf production.
	EnterForStatementNoShortIf(c *ForStatementNoShortIfContext)

	// EnterBasicForStatement is called when entering the basicForStatement production.
	EnterBasicForStatement(c *BasicForStatementContext)

	// EnterBasicForStatementNoShortIf is called when entering the basicForStatementNoShortIf production.
	EnterBasicForStatementNoShortIf(c *BasicForStatementNoShortIfContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterStatementExpressionList is called when entering the statementExpressionList production.
	EnterStatementExpressionList(c *StatementExpressionListContext)

	// EnterEnhancedForStatement is called when entering the enhancedForStatement production.
	EnterEnhancedForStatement(c *EnhancedForStatementContext)

	// EnterEnhancedForStatementNoShortIf is called when entering the enhancedForStatementNoShortIf production.
	EnterEnhancedForStatementNoShortIf(c *EnhancedForStatementNoShortIfContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterSynchronizedStatement is called when entering the synchronizedStatement production.
	EnterSynchronizedStatement(c *SynchronizedStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterCatches is called when entering the catches production.
	EnterCatches(c *CatchesContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterCatchFormalParameter is called when entering the catchFormalParameter production.
	EnterCatchFormalParameter(c *CatchFormalParameterContext)

	// EnterCatchType is called when entering the catchType production.
	EnterCatchType(c *CatchTypeContext)

	// EnterFinally_ is called when entering the finally_ production.
	EnterFinally_(c *Finally_Context)

	// EnterTryWithResourcesStatement is called when entering the tryWithResourcesStatement production.
	EnterTryWithResourcesStatement(c *TryWithResourcesStatementContext)

	// EnterResourceSpecification is called when entering the resourceSpecification production.
	EnterResourceSpecification(c *ResourceSpecificationContext)

	// EnterResourceList is called when entering the resourceList production.
	EnterResourceList(c *ResourceListContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterPrimaryNoNewArray is called when entering the primaryNoNewArray production.
	EnterPrimaryNoNewArray(c *PrimaryNoNewArrayContext)

	// EnterPrimaryNoNewArray_lf_arrayAccess is called when entering the primaryNoNewArray_lf_arrayAccess production.
	EnterPrimaryNoNewArray_lf_arrayAccess(c *PrimaryNoNewArray_lf_arrayAccessContext)

	// EnterPrimaryNoNewArray_lfno_arrayAccess is called when entering the primaryNoNewArray_lfno_arrayAccess production.
	EnterPrimaryNoNewArray_lfno_arrayAccess(c *PrimaryNoNewArray_lfno_arrayAccessContext)

	// EnterPrimaryNoNewArray_lf_primary is called when entering the primaryNoNewArray_lf_primary production.
	EnterPrimaryNoNewArray_lf_primary(c *PrimaryNoNewArray_lf_primaryContext)

	// EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when entering the primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary production.
	EnterPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(c *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext)

	// EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when entering the primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary production.
	EnterPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(c *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext)

	// EnterPrimaryNoNewArray_lfno_primary is called when entering the primaryNoNewArray_lfno_primary production.
	EnterPrimaryNoNewArray_lfno_primary(c *PrimaryNoNewArray_lfno_primaryContext)

	// EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when entering the primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary production.
	EnterPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(c *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext)

	// EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when entering the primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary production.
	EnterPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(c *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext)

	// EnterClassInstanceCreationExpression is called when entering the classInstanceCreationExpression production.
	EnterClassInstanceCreationExpression(c *ClassInstanceCreationExpressionContext)

	// EnterClassInstanceCreationExpression_lf_primary is called when entering the classInstanceCreationExpression_lf_primary production.
	EnterClassInstanceCreationExpression_lf_primary(c *ClassInstanceCreationExpression_lf_primaryContext)

	// EnterClassInstanceCreationExpression_lfno_primary is called when entering the classInstanceCreationExpression_lfno_primary production.
	EnterClassInstanceCreationExpression_lfno_primary(c *ClassInstanceCreationExpression_lfno_primaryContext)

	// EnterTypeArgumentsOrDiamond is called when entering the typeArgumentsOrDiamond production.
	EnterTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// EnterFieldAccess is called when entering the fieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterFieldAccess_lf_primary is called when entering the fieldAccess_lf_primary production.
	EnterFieldAccess_lf_primary(c *FieldAccess_lf_primaryContext)

	// EnterFieldAccess_lfno_primary is called when entering the fieldAccess_lfno_primary production.
	EnterFieldAccess_lfno_primary(c *FieldAccess_lfno_primaryContext)

	// EnterArrayAccess is called when entering the arrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterArrayAccess_lf_primary is called when entering the arrayAccess_lf_primary production.
	EnterArrayAccess_lf_primary(c *ArrayAccess_lf_primaryContext)

	// EnterArrayAccess_lfno_primary is called when entering the arrayAccess_lfno_primary production.
	EnterArrayAccess_lfno_primary(c *ArrayAccess_lfno_primaryContext)

	// EnterMethodInvocation is called when entering the methodInvocation production.
	EnterMethodInvocation(c *MethodInvocationContext)

	// EnterMethodInvocation_lf_primary is called when entering the methodInvocation_lf_primary production.
	EnterMethodInvocation_lf_primary(c *MethodInvocation_lf_primaryContext)

	// EnterMethodInvocation_lfno_primary is called when entering the methodInvocation_lfno_primary production.
	EnterMethodInvocation_lfno_primary(c *MethodInvocation_lfno_primaryContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterMethodReference is called when entering the methodReference production.
	EnterMethodReference(c *MethodReferenceContext)

	// EnterMethodReference_lf_primary is called when entering the methodReference_lf_primary production.
	EnterMethodReference_lf_primary(c *MethodReference_lf_primaryContext)

	// EnterMethodReference_lfno_primary is called when entering the methodReference_lfno_primary production.
	EnterMethodReference_lfno_primary(c *MethodReference_lfno_primaryContext)

	// EnterArrayCreationExpression is called when entering the arrayCreationExpression production.
	EnterArrayCreationExpression(c *ArrayCreationExpressionContext)

	// EnterDimExprs is called when entering the dimExprs production.
	EnterDimExprs(c *DimExprsContext)

	// EnterDimExpr is called when entering the dimExpr production.
	EnterDimExpr(c *DimExprContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterInferredFormalParameterList is called when entering the inferredFormalParameterList production.
	EnterInferredFormalParameterList(c *InferredFormalParameterListContext)

	// EnterLambdaBody is called when entering the lambdaBody production.
	EnterLambdaBody(c *LambdaBodyContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterLeftHandSide is called when entering the leftHandSide production.
	EnterLeftHandSide(c *LeftHandSideContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterConditionalOrExpression is called when entering the conditionalOrExpression production.
	EnterConditionalOrExpression(c *ConditionalOrExpressionContext)

	// EnterConditionalAndExpression is called when entering the conditionalAndExpression production.
	EnterConditionalAndExpression(c *ConditionalAndExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPreIncrementExpression is called when entering the preIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterPreDecrementExpression is called when entering the preDecrementExpression production.
	EnterPreDecrementExpression(c *PreDecrementExpressionContext)

	// EnterUnaryExpressionNotPlusMinus is called when entering the unaryExpressionNotPlusMinus production.
	EnterUnaryExpressionNotPlusMinus(c *UnaryExpressionNotPlusMinusContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostIncrementExpression is called when entering the postIncrementExpression production.
	EnterPostIncrementExpression(c *PostIncrementExpressionContext)

	// EnterPostIncrementExpression_lf_postfixExpression is called when entering the postIncrementExpression_lf_postfixExpression production.
	EnterPostIncrementExpression_lf_postfixExpression(c *PostIncrementExpression_lf_postfixExpressionContext)

	// EnterPostDecrementExpression is called when entering the postDecrementExpression production.
	EnterPostDecrementExpression(c *PostDecrementExpressionContext)

	// EnterPostDecrementExpression_lf_postfixExpression is called when entering the postDecrementExpression_lf_postfixExpression production.
	EnterPostDecrementExpression_lf_postfixExpression(c *PostDecrementExpression_lf_postfixExpressionContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitNumericType is called when exiting the numericType production.
	ExitNumericType(c *NumericTypeContext)

	// ExitIntegralType is called when exiting the integralType production.
	ExitIntegralType(c *IntegralTypeContext)

	// ExitFloatingPointType is called when exiting the floatingPointType production.
	ExitFloatingPointType(c *FloatingPointTypeContext)

	// ExitReferenceType is called when exiting the referenceType production.
	ExitReferenceType(c *ReferenceTypeContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitClassType_lf_classOrInterfaceType is called when exiting the classType_lf_classOrInterfaceType production.
	ExitClassType_lf_classOrInterfaceType(c *ClassType_lf_classOrInterfaceTypeContext)

	// ExitClassType_lfno_classOrInterfaceType is called when exiting the classType_lfno_classOrInterfaceType production.
	ExitClassType_lfno_classOrInterfaceType(c *ClassType_lfno_classOrInterfaceTypeContext)

	// ExitInterfaceType is called when exiting the interfaceType production.
	ExitInterfaceType(c *InterfaceTypeContext)

	// ExitInterfaceType_lf_classOrInterfaceType is called when exiting the interfaceType_lf_classOrInterfaceType production.
	ExitInterfaceType_lf_classOrInterfaceType(c *InterfaceType_lf_classOrInterfaceTypeContext)

	// ExitInterfaceType_lfno_classOrInterfaceType is called when exiting the interfaceType_lfno_classOrInterfaceType production.
	ExitInterfaceType_lfno_classOrInterfaceType(c *InterfaceType_lfno_classOrInterfaceTypeContext)

	// ExitTypeVariable is called when exiting the typeVariable production.
	ExitTypeVariable(c *TypeVariableContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitDims is called when exiting the dims production.
	ExitDims(c *DimsContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitTypeParameterModifier is called when exiting the typeParameterModifier production.
	ExitTypeParameterModifier(c *TypeParameterModifierContext)

	// ExitTypeBound is called when exiting the typeBound production.
	ExitTypeBound(c *TypeBoundContext)

	// ExitAdditionalBound is called when exiting the additionalBound production.
	ExitAdditionalBound(c *AdditionalBoundContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgumentList is called when exiting the typeArgumentList production.
	ExitTypeArgumentList(c *TypeArgumentListContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitWildcard is called when exiting the wildcard production.
	ExitWildcard(c *WildcardContext)

	// ExitWildcardBounds is called when exiting the wildcardBounds production.
	ExitWildcardBounds(c *WildcardBoundsContext)

	// ExitPackageName is called when exiting the packageName production.
	ExitPackageName(c *PackageNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitPackageOrTypeName is called when exiting the packageOrTypeName production.
	ExitPackageOrTypeName(c *PackageOrTypeNameContext)

	// ExitExpressionName is called when exiting the expressionName production.
	ExitExpressionName(c *ExpressionNameContext)

	// ExitMethodName is called when exiting the methodName production.
	ExitMethodName(c *MethodNameContext)

	// ExitAmbiguousName is called when exiting the ambiguousName production.
	ExitAmbiguousName(c *AmbiguousNameContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitPackageModifier is called when exiting the packageModifier production.
	ExitPackageModifier(c *PackageModifierContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitSingleTypeImportDeclaration is called when exiting the singleTypeImportDeclaration production.
	ExitSingleTypeImportDeclaration(c *SingleTypeImportDeclarationContext)

	// ExitTypeImportOnDemandDeclaration is called when exiting the typeImportOnDemandDeclaration production.
	ExitTypeImportOnDemandDeclaration(c *TypeImportOnDemandDeclarationContext)

	// ExitSingleStaticImportDeclaration is called when exiting the singleStaticImportDeclaration production.
	ExitSingleStaticImportDeclaration(c *SingleStaticImportDeclarationContext)

	// ExitStaticImportOnDemandDeclaration is called when exiting the staticImportOnDemandDeclaration production.
	ExitStaticImportOnDemandDeclaration(c *StaticImportOnDemandDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitNormalClassDeclaration is called when exiting the normalClassDeclaration production.
	ExitNormalClassDeclaration(c *NormalClassDeclarationContext)

	// ExitClassModifier is called when exiting the classModifier production.
	ExitClassModifier(c *ClassModifierContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameterList is called when exiting the typeParameterList production.
	ExitTypeParameterList(c *TypeParameterListContext)

	// ExitSuperclass is called when exiting the superclass production.
	ExitSuperclass(c *SuperclassContext)

	// ExitSuperinterfaces is called when exiting the superinterfaces production.
	ExitSuperinterfaces(c *SuperinterfacesContext)

	// ExitInterfaceTypeList is called when exiting the interfaceTypeList production.
	ExitInterfaceTypeList(c *InterfaceTypeListContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitClassMemberDeclaration is called when exiting the classMemberDeclaration production.
	ExitClassMemberDeclaration(c *ClassMemberDeclarationContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitFieldModifier is called when exiting the fieldModifier production.
	ExitFieldModifier(c *FieldModifierContext)

	// ExitVariableDeclaratorList is called when exiting the variableDeclaratorList production.
	ExitVariableDeclaratorList(c *VariableDeclaratorListContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitUnannType is called when exiting the unannType production.
	ExitUnannType(c *UnannTypeContext)

	// ExitUnannPrimitiveType is called when exiting the unannPrimitiveType production.
	ExitUnannPrimitiveType(c *UnannPrimitiveTypeContext)

	// ExitUnannReferenceType is called when exiting the unannReferenceType production.
	ExitUnannReferenceType(c *UnannReferenceTypeContext)

	// ExitUnannClassOrInterfaceType is called when exiting the unannClassOrInterfaceType production.
	ExitUnannClassOrInterfaceType(c *UnannClassOrInterfaceTypeContext)

	// ExitUnannClassType is called when exiting the unannClassType production.
	ExitUnannClassType(c *UnannClassTypeContext)

	// ExitUnannClassType_lf_unannClassOrInterfaceType is called when exiting the unannClassType_lf_unannClassOrInterfaceType production.
	ExitUnannClassType_lf_unannClassOrInterfaceType(c *UnannClassType_lf_unannClassOrInterfaceTypeContext)

	// ExitUnannClassType_lfno_unannClassOrInterfaceType is called when exiting the unannClassType_lfno_unannClassOrInterfaceType production.
	ExitUnannClassType_lfno_unannClassOrInterfaceType(c *UnannClassType_lfno_unannClassOrInterfaceTypeContext)

	// ExitUnannInterfaceType is called when exiting the unannInterfaceType production.
	ExitUnannInterfaceType(c *UnannInterfaceTypeContext)

	// ExitUnannInterfaceType_lf_unannClassOrInterfaceType is called when exiting the unannInterfaceType_lf_unannClassOrInterfaceType production.
	ExitUnannInterfaceType_lf_unannClassOrInterfaceType(c *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext)

	// ExitUnannInterfaceType_lfno_unannClassOrInterfaceType is called when exiting the unannInterfaceType_lfno_unannClassOrInterfaceType production.
	ExitUnannInterfaceType_lfno_unannClassOrInterfaceType(c *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext)

	// ExitUnannTypeVariable is called when exiting the unannTypeVariable production.
	ExitUnannTypeVariable(c *UnannTypeVariableContext)

	// ExitUnannArrayType is called when exiting the unannArrayType production.
	ExitUnannArrayType(c *UnannArrayTypeContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitMethodModifier is called when exiting the methodModifier production.
	ExitMethodModifier(c *MethodModifierContext)

	// ExitMethodHeader is called when exiting the methodHeader production.
	ExitMethodHeader(c *MethodHeaderContext)

	// ExitResult is called when exiting the result production.
	ExitResult(c *ResultContext)

	// ExitMethodDeclarator is called when exiting the methodDeclarator production.
	ExitMethodDeclarator(c *MethodDeclaratorContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitLastFormalParameter is called when exiting the lastFormalParameter production.
	ExitLastFormalParameter(c *LastFormalParameterContext)

	// ExitReceiverParameter is called when exiting the receiverParameter production.
	ExitReceiverParameter(c *ReceiverParameterContext)

	// ExitThrows_ is called when exiting the throws_ production.
	ExitThrows_(c *Throws_Context)

	// ExitExceptionTypeList is called when exiting the exceptionTypeList production.
	ExitExceptionTypeList(c *ExceptionTypeListContext)

	// ExitExceptionType is called when exiting the exceptionType production.
	ExitExceptionType(c *ExceptionTypeContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitInstanceInitializer is called when exiting the instanceInitializer production.
	ExitInstanceInitializer(c *InstanceInitializerContext)

	// ExitStaticInitializer is called when exiting the staticInitializer production.
	ExitStaticInitializer(c *StaticInitializerContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitConstructorModifier is called when exiting the constructorModifier production.
	ExitConstructorModifier(c *ConstructorModifierContext)

	// ExitConstructorDeclarator is called when exiting the constructorDeclarator production.
	ExitConstructorDeclarator(c *ConstructorDeclaratorContext)

	// ExitSimpleTypeName is called when exiting the simpleTypeName production.
	ExitSimpleTypeName(c *SimpleTypeNameContext)

	// ExitConstructorBody is called when exiting the constructorBody production.
	ExitConstructorBody(c *ConstructorBodyContext)

	// ExitExplicitConstructorInvocation is called when exiting the explicitConstructorInvocation production.
	ExitExplicitConstructorInvocation(c *ExplicitConstructorInvocationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumConstantList is called when exiting the enumConstantList production.
	ExitEnumConstantList(c *EnumConstantListContext)

	// ExitEnumConstant is called when exiting the enumConstant production.
	ExitEnumConstant(c *EnumConstantContext)

	// ExitEnumConstantModifier is called when exiting the enumConstantModifier production.
	ExitEnumConstantModifier(c *EnumConstantModifierContext)

	// ExitEnumBodyDeclarations is called when exiting the enumBodyDeclarations production.
	ExitEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitNormalInterfaceDeclaration is called when exiting the normalInterfaceDeclaration production.
	ExitNormalInterfaceDeclaration(c *NormalInterfaceDeclarationContext)

	// ExitInterfaceModifier is called when exiting the interfaceModifier production.
	ExitInterfaceModifier(c *InterfaceModifierContext)

	// ExitExtendsInterfaces is called when exiting the extendsInterfaces production.
	ExitExtendsInterfaces(c *ExtendsInterfacesContext)

	// ExitInterfaceBody is called when exiting the interfaceBody production.
	ExitInterfaceBody(c *InterfaceBodyContext)

	// ExitInterfaceMemberDeclaration is called when exiting the interfaceMemberDeclaration production.
	ExitInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// ExitConstantDeclaration is called when exiting the constantDeclaration production.
	ExitConstantDeclaration(c *ConstantDeclarationContext)

	// ExitConstantModifier is called when exiting the constantModifier production.
	ExitConstantModifier(c *ConstantModifierContext)

	// ExitInterfaceMethodDeclaration is called when exiting the interfaceMethodDeclaration production.
	ExitInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// ExitInterfaceMethodModifier is called when exiting the interfaceMethodModifier production.
	ExitInterfaceMethodModifier(c *InterfaceMethodModifierContext)

	// ExitAnnotationTypeDeclaration is called when exiting the annotationTypeDeclaration production.
	ExitAnnotationTypeDeclaration(c *AnnotationTypeDeclarationContext)

	// ExitAnnotationTypeBody is called when exiting the annotationTypeBody production.
	ExitAnnotationTypeBody(c *AnnotationTypeBodyContext)

	// ExitAnnotationTypeMemberDeclaration is called when exiting the annotationTypeMemberDeclaration production.
	ExitAnnotationTypeMemberDeclaration(c *AnnotationTypeMemberDeclarationContext)

	// ExitAnnotationTypeElementDeclaration is called when exiting the annotationTypeElementDeclaration production.
	ExitAnnotationTypeElementDeclaration(c *AnnotationTypeElementDeclarationContext)

	// ExitAnnotationTypeElementModifier is called when exiting the annotationTypeElementModifier production.
	ExitAnnotationTypeElementModifier(c *AnnotationTypeElementModifierContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitNormalAnnotation is called when exiting the normalAnnotation production.
	ExitNormalAnnotation(c *NormalAnnotationContext)

	// ExitElementValuePairList is called when exiting the elementValuePairList production.
	ExitElementValuePairList(c *ElementValuePairListContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitElementValue is called when exiting the elementValue production.
	ExitElementValue(c *ElementValueContext)

	// ExitElementValueArrayInitializer is called when exiting the elementValueArrayInitializer production.
	ExitElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// ExitElementValueList is called when exiting the elementValueList production.
	ExitElementValueList(c *ElementValueListContext)

	// ExitMarkerAnnotation is called when exiting the markerAnnotation production.
	ExitMarkerAnnotation(c *MarkerAnnotationContext)

	// ExitSingleElementAnnotation is called when exiting the singleElementAnnotation production.
	ExitSingleElementAnnotation(c *SingleElementAnnotationContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitVariableInitializerList is called when exiting the variableInitializerList production.
	ExitVariableInitializerList(c *VariableInitializerListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatements is called when exiting the blockStatements production.
	ExitBlockStatements(c *BlockStatementsContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclarationStatement is called when exiting the localVariableDeclarationStatement production.
	ExitLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatementNoShortIf is called when exiting the statementNoShortIf production.
	ExitStatementNoShortIf(c *StatementNoShortIfContext)

	// ExitStatementWithoutTrailingSubstatement is called when exiting the statementWithoutTrailingSubstatement production.
	ExitStatementWithoutTrailingSubstatement(c *StatementWithoutTrailingSubstatementContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitLabeledStatement is called when exiting the labeledStatement production.
	ExitLabeledStatement(c *LabeledStatementContext)

	// ExitLabeledStatementNoShortIf is called when exiting the labeledStatementNoShortIf production.
	ExitLabeledStatementNoShortIf(c *LabeledStatementNoShortIfContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitStatementExpression is called when exiting the statementExpression production.
	ExitStatementExpression(c *StatementExpressionContext)

	// ExitIfThenStatement is called when exiting the ifThenStatement production.
	ExitIfThenStatement(c *IfThenStatementContext)

	// ExitIfThenElseStatement is called when exiting the ifThenElseStatement production.
	ExitIfThenElseStatement(c *IfThenElseStatementContext)

	// ExitIfThenElseStatementNoShortIf is called when exiting the ifThenElseStatementNoShortIf production.
	ExitIfThenElseStatementNoShortIf(c *IfThenElseStatementNoShortIfContext)

	// ExitAssertStatement is called when exiting the assertStatement production.
	ExitAssertStatement(c *AssertStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabels is called when exiting the switchLabels production.
	ExitSwitchLabels(c *SwitchLabelsContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitEnumConstantName is called when exiting the enumConstantName production.
	ExitEnumConstantName(c *EnumConstantNameContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitWhileStatementNoShortIf is called when exiting the whileStatementNoShortIf production.
	ExitWhileStatementNoShortIf(c *WhileStatementNoShortIfContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForStatementNoShortIf is called when exiting the forStatementNoShortIf production.
	ExitForStatementNoShortIf(c *ForStatementNoShortIfContext)

	// ExitBasicForStatement is called when exiting the basicForStatement production.
	ExitBasicForStatement(c *BasicForStatementContext)

	// ExitBasicForStatementNoShortIf is called when exiting the basicForStatementNoShortIf production.
	ExitBasicForStatementNoShortIf(c *BasicForStatementNoShortIfContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitStatementExpressionList is called when exiting the statementExpressionList production.
	ExitStatementExpressionList(c *StatementExpressionListContext)

	// ExitEnhancedForStatement is called when exiting the enhancedForStatement production.
	ExitEnhancedForStatement(c *EnhancedForStatementContext)

	// ExitEnhancedForStatementNoShortIf is called when exiting the enhancedForStatementNoShortIf production.
	ExitEnhancedForStatementNoShortIf(c *EnhancedForStatementNoShortIfContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitSynchronizedStatement is called when exiting the synchronizedStatement production.
	ExitSynchronizedStatement(c *SynchronizedStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitCatches is called when exiting the catches production.
	ExitCatches(c *CatchesContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitCatchFormalParameter is called when exiting the catchFormalParameter production.
	ExitCatchFormalParameter(c *CatchFormalParameterContext)

	// ExitCatchType is called when exiting the catchType production.
	ExitCatchType(c *CatchTypeContext)

	// ExitFinally_ is called when exiting the finally_ production.
	ExitFinally_(c *Finally_Context)

	// ExitTryWithResourcesStatement is called when exiting the tryWithResourcesStatement production.
	ExitTryWithResourcesStatement(c *TryWithResourcesStatementContext)

	// ExitResourceSpecification is called when exiting the resourceSpecification production.
	ExitResourceSpecification(c *ResourceSpecificationContext)

	// ExitResourceList is called when exiting the resourceList production.
	ExitResourceList(c *ResourceListContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitPrimaryNoNewArray is called when exiting the primaryNoNewArray production.
	ExitPrimaryNoNewArray(c *PrimaryNoNewArrayContext)

	// ExitPrimaryNoNewArray_lf_arrayAccess is called when exiting the primaryNoNewArray_lf_arrayAccess production.
	ExitPrimaryNoNewArray_lf_arrayAccess(c *PrimaryNoNewArray_lf_arrayAccessContext)

	// ExitPrimaryNoNewArray_lfno_arrayAccess is called when exiting the primaryNoNewArray_lfno_arrayAccess production.
	ExitPrimaryNoNewArray_lfno_arrayAccess(c *PrimaryNoNewArray_lfno_arrayAccessContext)

	// ExitPrimaryNoNewArray_lf_primary is called when exiting the primaryNoNewArray_lf_primary production.
	ExitPrimaryNoNewArray_lf_primary(c *PrimaryNoNewArray_lf_primaryContext)

	// ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary is called when exiting the primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary production.
	ExitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(c *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext)

	// ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary is called when exiting the primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary production.
	ExitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(c *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext)

	// ExitPrimaryNoNewArray_lfno_primary is called when exiting the primaryNoNewArray_lfno_primary production.
	ExitPrimaryNoNewArray_lfno_primary(c *PrimaryNoNewArray_lfno_primaryContext)

	// ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary is called when exiting the primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary production.
	ExitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(c *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext)

	// ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary is called when exiting the primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary production.
	ExitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(c *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext)

	// ExitClassInstanceCreationExpression is called when exiting the classInstanceCreationExpression production.
	ExitClassInstanceCreationExpression(c *ClassInstanceCreationExpressionContext)

	// ExitClassInstanceCreationExpression_lf_primary is called when exiting the classInstanceCreationExpression_lf_primary production.
	ExitClassInstanceCreationExpression_lf_primary(c *ClassInstanceCreationExpression_lf_primaryContext)

	// ExitClassInstanceCreationExpression_lfno_primary is called when exiting the classInstanceCreationExpression_lfno_primary production.
	ExitClassInstanceCreationExpression_lfno_primary(c *ClassInstanceCreationExpression_lfno_primaryContext)

	// ExitTypeArgumentsOrDiamond is called when exiting the typeArgumentsOrDiamond production.
	ExitTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// ExitFieldAccess is called when exiting the fieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitFieldAccess_lf_primary is called when exiting the fieldAccess_lf_primary production.
	ExitFieldAccess_lf_primary(c *FieldAccess_lf_primaryContext)

	// ExitFieldAccess_lfno_primary is called when exiting the fieldAccess_lfno_primary production.
	ExitFieldAccess_lfno_primary(c *FieldAccess_lfno_primaryContext)

	// ExitArrayAccess is called when exiting the arrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitArrayAccess_lf_primary is called when exiting the arrayAccess_lf_primary production.
	ExitArrayAccess_lf_primary(c *ArrayAccess_lf_primaryContext)

	// ExitArrayAccess_lfno_primary is called when exiting the arrayAccess_lfno_primary production.
	ExitArrayAccess_lfno_primary(c *ArrayAccess_lfno_primaryContext)

	// ExitMethodInvocation is called when exiting the methodInvocation production.
	ExitMethodInvocation(c *MethodInvocationContext)

	// ExitMethodInvocation_lf_primary is called when exiting the methodInvocation_lf_primary production.
	ExitMethodInvocation_lf_primary(c *MethodInvocation_lf_primaryContext)

	// ExitMethodInvocation_lfno_primary is called when exiting the methodInvocation_lfno_primary production.
	ExitMethodInvocation_lfno_primary(c *MethodInvocation_lfno_primaryContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitMethodReference is called when exiting the methodReference production.
	ExitMethodReference(c *MethodReferenceContext)

	// ExitMethodReference_lf_primary is called when exiting the methodReference_lf_primary production.
	ExitMethodReference_lf_primary(c *MethodReference_lf_primaryContext)

	// ExitMethodReference_lfno_primary is called when exiting the methodReference_lfno_primary production.
	ExitMethodReference_lfno_primary(c *MethodReference_lfno_primaryContext)

	// ExitArrayCreationExpression is called when exiting the arrayCreationExpression production.
	ExitArrayCreationExpression(c *ArrayCreationExpressionContext)

	// ExitDimExprs is called when exiting the dimExprs production.
	ExitDimExprs(c *DimExprsContext)

	// ExitDimExpr is called when exiting the dimExpr production.
	ExitDimExpr(c *DimExprContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitInferredFormalParameterList is called when exiting the inferredFormalParameterList production.
	ExitInferredFormalParameterList(c *InferredFormalParameterListContext)

	// ExitLambdaBody is called when exiting the lambdaBody production.
	ExitLambdaBody(c *LambdaBodyContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitLeftHandSide is called when exiting the leftHandSide production.
	ExitLeftHandSide(c *LeftHandSideContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitConditionalOrExpression is called when exiting the conditionalOrExpression production.
	ExitConditionalOrExpression(c *ConditionalOrExpressionContext)

	// ExitConditionalAndExpression is called when exiting the conditionalAndExpression production.
	ExitConditionalAndExpression(c *ConditionalAndExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPreIncrementExpression is called when exiting the preIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitPreDecrementExpression is called when exiting the preDecrementExpression production.
	ExitPreDecrementExpression(c *PreDecrementExpressionContext)

	// ExitUnaryExpressionNotPlusMinus is called when exiting the unaryExpressionNotPlusMinus production.
	ExitUnaryExpressionNotPlusMinus(c *UnaryExpressionNotPlusMinusContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostIncrementExpression is called when exiting the postIncrementExpression production.
	ExitPostIncrementExpression(c *PostIncrementExpressionContext)

	// ExitPostIncrementExpression_lf_postfixExpression is called when exiting the postIncrementExpression_lf_postfixExpression production.
	ExitPostIncrementExpression_lf_postfixExpression(c *PostIncrementExpression_lf_postfixExpressionContext)

	// ExitPostDecrementExpression is called when exiting the postDecrementExpression production.
	ExitPostDecrementExpression(c *PostDecrementExpressionContext)

	// ExitPostDecrementExpression_lf_postfixExpression is called when exiting the postDecrementExpression_lf_postfixExpression production.
	ExitPostDecrementExpression_lf_postfixExpression(c *PostDecrementExpression_lf_postfixExpressionContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)
}
