// Generated from Java.g4 by ANTLR 4.7.

package java // Java
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJavaListener is a complete listener for a parse tree produced by JavaParser.
type BaseJavaListener struct{}

var _ JavaListener = &BaseJavaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseJavaListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseJavaListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseJavaListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseJavaListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseJavaListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseJavaListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseJavaListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseJavaListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseJavaListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseJavaListener) ExitModifier(ctx *ModifierContext) {}

// EnterClassOrInterfaceModifier is called when production classOrInterfaceModifier is entered.
func (s *BaseJavaListener) EnterClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {}

// ExitClassOrInterfaceModifier is called when production classOrInterfaceModifier is exited.
func (s *BaseJavaListener) ExitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseJavaListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseJavaListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseJavaListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseJavaListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseJavaListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseJavaListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseJavaListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseJavaListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseJavaListener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseJavaListener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseJavaListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseJavaListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumConstants is called when production enumConstants is entered.
func (s *BaseJavaListener) EnterEnumConstants(ctx *EnumConstantsContext) {}

// ExitEnumConstants is called when production enumConstants is exited.
func (s *BaseJavaListener) ExitEnumConstants(ctx *EnumConstantsContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseJavaListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseJavaListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEnumBodyDeclarations is called when production enumBodyDeclarations is entered.
func (s *BaseJavaListener) EnterEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// ExitEnumBodyDeclarations is called when production enumBodyDeclarations is exited.
func (s *BaseJavaListener) ExitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseJavaListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseJavaListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseJavaListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseJavaListener) ExitTypeList(ctx *TypeListContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseJavaListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseJavaListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterInterfaceBody is called when production interfaceBody is entered.
func (s *BaseJavaListener) EnterInterfaceBody(ctx *InterfaceBodyContext) {}

// ExitInterfaceBody is called when production interfaceBody is exited.
func (s *BaseJavaListener) ExitInterfaceBody(ctx *InterfaceBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseJavaListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseJavaListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseJavaListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseJavaListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseJavaListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseJavaListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterGenericMethodDeclaration is called when production genericMethodDeclaration is entered.
func (s *BaseJavaListener) EnterGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {}

// ExitGenericMethodDeclaration is called when production genericMethodDeclaration is exited.
func (s *BaseJavaListener) ExitGenericMethodDeclaration(ctx *GenericMethodDeclarationContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseJavaListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseJavaListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {}

// EnterGenericConstructorDeclaration is called when production genericConstructorDeclaration is entered.
func (s *BaseJavaListener) EnterGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// ExitGenericConstructorDeclaration is called when production genericConstructorDeclaration is exited.
func (s *BaseJavaListener) ExitGenericConstructorDeclaration(ctx *GenericConstructorDeclarationContext) {
}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseJavaListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseJavaListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is entered.
func (s *BaseJavaListener) EnterInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// ExitInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is exited.
func (s *BaseJavaListener) ExitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// EnterInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is entered.
func (s *BaseJavaListener) EnterInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// ExitInterfaceMemberDeclaration is called when production interfaceMemberDeclaration is exited.
func (s *BaseJavaListener) ExitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) {}

// EnterConstDeclaration is called when production constDeclaration is entered.
func (s *BaseJavaListener) EnterConstDeclaration(ctx *ConstDeclarationContext) {}

// ExitConstDeclaration is called when production constDeclaration is exited.
func (s *BaseJavaListener) ExitConstDeclaration(ctx *ConstDeclarationContext) {}

// EnterConstantDeclarator is called when production constantDeclarator is entered.
func (s *BaseJavaListener) EnterConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// ExitConstantDeclarator is called when production constantDeclarator is exited.
func (s *BaseJavaListener) ExitConstantDeclarator(ctx *ConstantDeclaratorContext) {}

// EnterInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is entered.
func (s *BaseJavaListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// ExitInterfaceMethodDeclaration is called when production interfaceMethodDeclaration is exited.
func (s *BaseJavaListener) ExitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {}

// EnterGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is entered.
func (s *BaseJavaListener) EnterGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// ExitGenericInterfaceMethodDeclaration is called when production genericInterfaceMethodDeclaration is exited.
func (s *BaseJavaListener) ExitGenericInterfaceMethodDeclaration(ctx *GenericInterfaceMethodDeclarationContext) {
}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseJavaListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseJavaListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseJavaListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseJavaListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseJavaListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseJavaListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseJavaListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseJavaListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseJavaListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseJavaListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterEnumConstantName is called when production enumConstantName is entered.
func (s *BaseJavaListener) EnterEnumConstantName(ctx *EnumConstantNameContext) {}

// ExitEnumConstantName is called when production enumConstantName is exited.
func (s *BaseJavaListener) ExitEnumConstantName(ctx *EnumConstantNameContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseJavaListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseJavaListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseJavaListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseJavaListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseJavaListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseJavaListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseJavaListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseJavaListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseJavaListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseJavaListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BaseJavaListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BaseJavaListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseJavaListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseJavaListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJavaListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJavaListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseJavaListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseJavaListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseJavaListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseJavaListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseJavaListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseJavaListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterConstructorBody is called when production constructorBody is entered.
func (s *BaseJavaListener) EnterConstructorBody(ctx *ConstructorBodyContext) {}

// ExitConstructorBody is called when production constructorBody is exited.
func (s *BaseJavaListener) ExitConstructorBody(ctx *ConstructorBodyContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseJavaListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseJavaListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJavaListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJavaListener) ExitLiteral(ctx *LiteralContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseJavaListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseJavaListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationName is called when production annotationName is entered.
func (s *BaseJavaListener) EnterAnnotationName(ctx *AnnotationNameContext) {}

// ExitAnnotationName is called when production annotationName is exited.
func (s *BaseJavaListener) ExitAnnotationName(ctx *AnnotationNameContext) {}

// EnterElementValuePairs is called when production elementValuePairs is entered.
func (s *BaseJavaListener) EnterElementValuePairs(ctx *ElementValuePairsContext) {}

// ExitElementValuePairs is called when production elementValuePairs is exited.
func (s *BaseJavaListener) ExitElementValuePairs(ctx *ElementValuePairsContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseJavaListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseJavaListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseJavaListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseJavaListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseJavaListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseJavaListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterAnnotationTypeDeclaration is called when production annotationTypeDeclaration is entered.
func (s *BaseJavaListener) EnterAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// ExitAnnotationTypeDeclaration is called when production annotationTypeDeclaration is exited.
func (s *BaseJavaListener) ExitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) {}

// EnterAnnotationTypeBody is called when production annotationTypeBody is entered.
func (s *BaseJavaListener) EnterAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// ExitAnnotationTypeBody is called when production annotationTypeBody is exited.
func (s *BaseJavaListener) ExitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) {}

// EnterAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is entered.
func (s *BaseJavaListener) EnterAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// ExitAnnotationTypeElementDeclaration is called when production annotationTypeElementDeclaration is exited.
func (s *BaseJavaListener) ExitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) {
}

// EnterAnnotationTypeElementRest is called when production annotationTypeElementRest is entered.
func (s *BaseJavaListener) EnterAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {}

// ExitAnnotationTypeElementRest is called when production annotationTypeElementRest is exited.
func (s *BaseJavaListener) ExitAnnotationTypeElementRest(ctx *AnnotationTypeElementRestContext) {}

// EnterAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is entered.
func (s *BaseJavaListener) EnterAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// ExitAnnotationMethodOrConstantRest is called when production annotationMethodOrConstantRest is exited.
func (s *BaseJavaListener) ExitAnnotationMethodOrConstantRest(ctx *AnnotationMethodOrConstantRestContext) {
}

// EnterAnnotationMethodRest is called when production annotationMethodRest is entered.
func (s *BaseJavaListener) EnterAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// ExitAnnotationMethodRest is called when production annotationMethodRest is exited.
func (s *BaseJavaListener) ExitAnnotationMethodRest(ctx *AnnotationMethodRestContext) {}

// EnterAnnotationConstantRest is called when production annotationConstantRest is entered.
func (s *BaseJavaListener) EnterAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// ExitAnnotationConstantRest is called when production annotationConstantRest is exited.
func (s *BaseJavaListener) ExitAnnotationConstantRest(ctx *AnnotationConstantRestContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseJavaListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseJavaListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJavaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJavaListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseJavaListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseJavaListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is entered.
func (s *BaseJavaListener) EnterLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// ExitLocalVariableDeclarationStatement is called when production localVariableDeclarationStatement is exited.
func (s *BaseJavaListener) ExitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) {
}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseJavaListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseJavaListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJavaListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJavaListener) ExitStatement(ctx *StatementContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseJavaListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseJavaListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseJavaListener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseJavaListener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseJavaListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseJavaListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterResourceSpecification is called when production resourceSpecification is entered.
func (s *BaseJavaListener) EnterResourceSpecification(ctx *ResourceSpecificationContext) {}

// ExitResourceSpecification is called when production resourceSpecification is exited.
func (s *BaseJavaListener) ExitResourceSpecification(ctx *ResourceSpecificationContext) {}

// EnterResources is called when production resources is entered.
func (s *BaseJavaListener) EnterResources(ctx *ResourcesContext) {}

// ExitResources is called when production resources is exited.
func (s *BaseJavaListener) ExitResources(ctx *ResourcesContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseJavaListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseJavaListener) ExitResource(ctx *ResourceContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseJavaListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseJavaListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseJavaListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseJavaListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseJavaListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseJavaListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseJavaListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseJavaListener) ExitForInit(ctx *ForInitContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BaseJavaListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BaseJavaListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseJavaListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseJavaListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseJavaListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseJavaListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseJavaListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseJavaListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterStatementExpression is called when production statementExpression is entered.
func (s *BaseJavaListener) EnterStatementExpression(ctx *StatementExpressionContext) {}

// ExitStatementExpression is called when production statementExpression is exited.
func (s *BaseJavaListener) ExitStatementExpression(ctx *StatementExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseJavaListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseJavaListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJavaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJavaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseJavaListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseJavaListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseJavaListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseJavaListener) ExitCreator(ctx *CreatorContext) {}

// EnterCreatedName is called when production createdName is entered.
func (s *BaseJavaListener) EnterCreatedName(ctx *CreatedNameContext) {}

// ExitCreatedName is called when production createdName is exited.
func (s *BaseJavaListener) ExitCreatedName(ctx *CreatedNameContext) {}

// EnterInnerCreator is called when production innerCreator is entered.
func (s *BaseJavaListener) EnterInnerCreator(ctx *InnerCreatorContext) {}

// ExitInnerCreator is called when production innerCreator is exited.
func (s *BaseJavaListener) ExitInnerCreator(ctx *InnerCreatorContext) {}

// EnterArrayCreatorRest is called when production arrayCreatorRest is entered.
func (s *BaseJavaListener) EnterArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// ExitArrayCreatorRest is called when production arrayCreatorRest is exited.
func (s *BaseJavaListener) ExitArrayCreatorRest(ctx *ArrayCreatorRestContext) {}

// EnterClassCreatorRest is called when production classCreatorRest is entered.
func (s *BaseJavaListener) EnterClassCreatorRest(ctx *ClassCreatorRestContext) {}

// ExitClassCreatorRest is called when production classCreatorRest is exited.
func (s *BaseJavaListener) ExitClassCreatorRest(ctx *ClassCreatorRestContext) {}

// EnterExplicitGenericInvocation is called when production explicitGenericInvocation is entered.
func (s *BaseJavaListener) EnterExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {}

// ExitExplicitGenericInvocation is called when production explicitGenericInvocation is exited.
func (s *BaseJavaListener) ExitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) {}

// EnterNonWildcardTypeArguments is called when production nonWildcardTypeArguments is entered.
func (s *BaseJavaListener) EnterNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {}

// ExitNonWildcardTypeArguments is called when production nonWildcardTypeArguments is exited.
func (s *BaseJavaListener) ExitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseJavaListener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseJavaListener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is entered.
func (s *BaseJavaListener) EnterNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// ExitNonWildcardTypeArgumentsOrDiamond is called when production nonWildcardTypeArgumentsOrDiamond is exited.
func (s *BaseJavaListener) ExitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) {
}

// EnterSuperSuffix is called when production superSuffix is entered.
func (s *BaseJavaListener) EnterSuperSuffix(ctx *SuperSuffixContext) {}

// ExitSuperSuffix is called when production superSuffix is exited.
func (s *BaseJavaListener) ExitSuperSuffix(ctx *SuperSuffixContext) {}

// EnterExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is entered.
func (s *BaseJavaListener) EnterExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// ExitExplicitGenericInvocationSuffix is called when production explicitGenericInvocationSuffix is exited.
func (s *BaseJavaListener) ExitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) {
}

// EnterArguments is called when production arguments is entered.
func (s *BaseJavaListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseJavaListener) ExitArguments(ctx *ArgumentsContext) {}
