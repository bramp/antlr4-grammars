// Code generated from ObjectiveCParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package objectivec // ObjectiveCParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObjectiveCParserListener is a complete listener for a parse tree produced by ObjectiveCParser.
type BaseObjectiveCParserListener struct{}

var _ ObjectiveCParserListener = &BaseObjectiveCParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjectiveCParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjectiveCParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjectiveCParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjectiveCParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *BaseObjectiveCParserListener) EnterTranslationUnit(ctx *TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *BaseObjectiveCParserListener) ExitTranslationUnit(ctx *TranslationUnitContext) {}

// EnterTopLevelDeclaration is called when production topLevelDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterTopLevelDeclaration(ctx *TopLevelDeclarationContext) {}

// ExitTopLevelDeclaration is called when production topLevelDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitTopLevelDeclaration(ctx *TopLevelDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterClassInterface is called when production classInterface is entered.
func (s *BaseObjectiveCParserListener) EnterClassInterface(ctx *ClassInterfaceContext) {}

// ExitClassInterface is called when production classInterface is exited.
func (s *BaseObjectiveCParserListener) ExitClassInterface(ctx *ClassInterfaceContext) {}

// EnterCategoryInterface is called when production categoryInterface is entered.
func (s *BaseObjectiveCParserListener) EnterCategoryInterface(ctx *CategoryInterfaceContext) {}

// ExitCategoryInterface is called when production categoryInterface is exited.
func (s *BaseObjectiveCParserListener) ExitCategoryInterface(ctx *CategoryInterfaceContext) {}

// EnterClassImplementation is called when production classImplementation is entered.
func (s *BaseObjectiveCParserListener) EnterClassImplementation(ctx *ClassImplementationContext) {}

// ExitClassImplementation is called when production classImplementation is exited.
func (s *BaseObjectiveCParserListener) ExitClassImplementation(ctx *ClassImplementationContext) {}

// EnterCategoryImplementation is called when production categoryImplementation is entered.
func (s *BaseObjectiveCParserListener) EnterCategoryImplementation(ctx *CategoryImplementationContext) {
}

// ExitCategoryImplementation is called when production categoryImplementation is exited.
func (s *BaseObjectiveCParserListener) ExitCategoryImplementation(ctx *CategoryImplementationContext) {
}

// EnterGenericTypeSpecifier is called when production genericTypeSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterGenericTypeSpecifier(ctx *GenericTypeSpecifierContext) {}

// ExitGenericTypeSpecifier is called when production genericTypeSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitGenericTypeSpecifier(ctx *GenericTypeSpecifierContext) {}

// EnterProtocolDeclaration is called when production protocolDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolDeclaration(ctx *ProtocolDeclarationContext) {}

// ExitProtocolDeclaration is called when production protocolDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolDeclaration(ctx *ProtocolDeclarationContext) {}

// EnterProtocolDeclarationSection is called when production protocolDeclarationSection is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolDeclarationSection(ctx *ProtocolDeclarationSectionContext) {
}

// ExitProtocolDeclarationSection is called when production protocolDeclarationSection is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolDeclarationSection(ctx *ProtocolDeclarationSectionContext) {
}

// EnterProtocolDeclarationList is called when production protocolDeclarationList is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolDeclarationList(ctx *ProtocolDeclarationListContext) {
}

// ExitProtocolDeclarationList is called when production protocolDeclarationList is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolDeclarationList(ctx *ProtocolDeclarationListContext) {
}

// EnterClassDeclarationList is called when production classDeclarationList is entered.
func (s *BaseObjectiveCParserListener) EnterClassDeclarationList(ctx *ClassDeclarationListContext) {}

// ExitClassDeclarationList is called when production classDeclarationList is exited.
func (s *BaseObjectiveCParserListener) ExitClassDeclarationList(ctx *ClassDeclarationListContext) {}

// EnterProtocolList is called when production protocolList is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolList(ctx *ProtocolListContext) {}

// ExitProtocolList is called when production protocolList is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolList(ctx *ProtocolListContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterPropertyAttributesList is called when production propertyAttributesList is entered.
func (s *BaseObjectiveCParserListener) EnterPropertyAttributesList(ctx *PropertyAttributesListContext) {
}

// ExitPropertyAttributesList is called when production propertyAttributesList is exited.
func (s *BaseObjectiveCParserListener) ExitPropertyAttributesList(ctx *PropertyAttributesListContext) {
}

// EnterPropertyAttribute is called when production propertyAttribute is entered.
func (s *BaseObjectiveCParserListener) EnterPropertyAttribute(ctx *PropertyAttributeContext) {}

// ExitPropertyAttribute is called when production propertyAttribute is exited.
func (s *BaseObjectiveCParserListener) ExitPropertyAttribute(ctx *PropertyAttributeContext) {}

// EnterProtocolName is called when production protocolName is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolName(ctx *ProtocolNameContext) {}

// ExitProtocolName is called when production protocolName is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolName(ctx *ProtocolNameContext) {}

// EnterInstanceVariables is called when production instanceVariables is entered.
func (s *BaseObjectiveCParserListener) EnterInstanceVariables(ctx *InstanceVariablesContext) {}

// ExitInstanceVariables is called when production instanceVariables is exited.
func (s *BaseObjectiveCParserListener) ExitInstanceVariables(ctx *InstanceVariablesContext) {}

// EnterVisibilitySection is called when production visibilitySection is entered.
func (s *BaseObjectiveCParserListener) EnterVisibilitySection(ctx *VisibilitySectionContext) {}

// ExitVisibilitySection is called when production visibilitySection is exited.
func (s *BaseObjectiveCParserListener) ExitVisibilitySection(ctx *VisibilitySectionContext) {}

// EnterAccessModifier is called when production accessModifier is entered.
func (s *BaseObjectiveCParserListener) EnterAccessModifier(ctx *AccessModifierContext) {}

// ExitAccessModifier is called when production accessModifier is exited.
func (s *BaseObjectiveCParserListener) ExitAccessModifier(ctx *AccessModifierContext) {}

// EnterInterfaceDeclarationList is called when production interfaceDeclarationList is entered.
func (s *BaseObjectiveCParserListener) EnterInterfaceDeclarationList(ctx *InterfaceDeclarationListContext) {
}

// ExitInterfaceDeclarationList is called when production interfaceDeclarationList is exited.
func (s *BaseObjectiveCParserListener) ExitInterfaceDeclarationList(ctx *InterfaceDeclarationListContext) {
}

// EnterClassMethodDeclaration is called when production classMethodDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterClassMethodDeclaration(ctx *ClassMethodDeclarationContext) {
}

// ExitClassMethodDeclaration is called when production classMethodDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitClassMethodDeclaration(ctx *ClassMethodDeclarationContext) {
}

// EnterInstanceMethodDeclaration is called when production instanceMethodDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterInstanceMethodDeclaration(ctx *InstanceMethodDeclarationContext) {
}

// ExitInstanceMethodDeclaration is called when production instanceMethodDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitInstanceMethodDeclaration(ctx *InstanceMethodDeclarationContext) {
}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterImplementationDefinitionList is called when production implementationDefinitionList is entered.
func (s *BaseObjectiveCParserListener) EnterImplementationDefinitionList(ctx *ImplementationDefinitionListContext) {
}

// ExitImplementationDefinitionList is called when production implementationDefinitionList is exited.
func (s *BaseObjectiveCParserListener) ExitImplementationDefinitionList(ctx *ImplementationDefinitionListContext) {
}

// EnterClassMethodDefinition is called when production classMethodDefinition is entered.
func (s *BaseObjectiveCParserListener) EnterClassMethodDefinition(ctx *ClassMethodDefinitionContext) {
}

// ExitClassMethodDefinition is called when production classMethodDefinition is exited.
func (s *BaseObjectiveCParserListener) ExitClassMethodDefinition(ctx *ClassMethodDefinitionContext) {}

// EnterInstanceMethodDefinition is called when production instanceMethodDefinition is entered.
func (s *BaseObjectiveCParserListener) EnterInstanceMethodDefinition(ctx *InstanceMethodDefinitionContext) {
}

// ExitInstanceMethodDefinition is called when production instanceMethodDefinition is exited.
func (s *BaseObjectiveCParserListener) ExitInstanceMethodDefinition(ctx *InstanceMethodDefinitionContext) {
}

// EnterMethodDefinition is called when production methodDefinition is entered.
func (s *BaseObjectiveCParserListener) EnterMethodDefinition(ctx *MethodDefinitionContext) {}

// ExitMethodDefinition is called when production methodDefinition is exited.
func (s *BaseObjectiveCParserListener) ExitMethodDefinition(ctx *MethodDefinitionContext) {}

// EnterMethodSelector is called when production methodSelector is entered.
func (s *BaseObjectiveCParserListener) EnterMethodSelector(ctx *MethodSelectorContext) {}

// ExitMethodSelector is called when production methodSelector is exited.
func (s *BaseObjectiveCParserListener) ExitMethodSelector(ctx *MethodSelectorContext) {}

// EnterKeywordDeclarator is called when production keywordDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterKeywordDeclarator(ctx *KeywordDeclaratorContext) {}

// ExitKeywordDeclarator is called when production keywordDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitKeywordDeclarator(ctx *KeywordDeclaratorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseObjectiveCParserListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseObjectiveCParserListener) ExitSelector(ctx *SelectorContext) {}

// EnterMethodType is called when production methodType is entered.
func (s *BaseObjectiveCParserListener) EnterMethodType(ctx *MethodTypeContext) {}

// ExitMethodType is called when production methodType is exited.
func (s *BaseObjectiveCParserListener) ExitMethodType(ctx *MethodTypeContext) {}

// EnterPropertyImplementation is called when production propertyImplementation is entered.
func (s *BaseObjectiveCParserListener) EnterPropertyImplementation(ctx *PropertyImplementationContext) {
}

// ExitPropertyImplementation is called when production propertyImplementation is exited.
func (s *BaseObjectiveCParserListener) ExitPropertyImplementation(ctx *PropertyImplementationContext) {
}

// EnterPropertySynthesizeList is called when production propertySynthesizeList is entered.
func (s *BaseObjectiveCParserListener) EnterPropertySynthesizeList(ctx *PropertySynthesizeListContext) {
}

// ExitPropertySynthesizeList is called when production propertySynthesizeList is exited.
func (s *BaseObjectiveCParserListener) ExitPropertySynthesizeList(ctx *PropertySynthesizeListContext) {
}

// EnterPropertySynthesizeItem is called when production propertySynthesizeItem is entered.
func (s *BaseObjectiveCParserListener) EnterPropertySynthesizeItem(ctx *PropertySynthesizeItemContext) {
}

// ExitPropertySynthesizeItem is called when production propertySynthesizeItem is exited.
func (s *BaseObjectiveCParserListener) ExitPropertySynthesizeItem(ctx *PropertySynthesizeItemContext) {
}

// EnterBlockType is called when production blockType is entered.
func (s *BaseObjectiveCParserListener) EnterBlockType(ctx *BlockTypeContext) {}

// ExitBlockType is called when production blockType is exited.
func (s *BaseObjectiveCParserListener) ExitBlockType(ctx *BlockTypeContext) {}

// EnterGenericsSpecifier is called when production genericsSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterGenericsSpecifier(ctx *GenericsSpecifierContext) {}

// ExitGenericsSpecifier is called when production genericsSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitGenericsSpecifier(ctx *GenericsSpecifierContext) {}

// EnterTypeSpecifierWithPrefixes is called when production typeSpecifierWithPrefixes is entered.
func (s *BaseObjectiveCParserListener) EnterTypeSpecifierWithPrefixes(ctx *TypeSpecifierWithPrefixesContext) {
}

// ExitTypeSpecifierWithPrefixes is called when production typeSpecifierWithPrefixes is exited.
func (s *BaseObjectiveCParserListener) ExitTypeSpecifierWithPrefixes(ctx *TypeSpecifierWithPrefixesContext) {
}

// EnterDictionaryExpression is called when production dictionaryExpression is entered.
func (s *BaseObjectiveCParserListener) EnterDictionaryExpression(ctx *DictionaryExpressionContext) {}

// ExitDictionaryExpression is called when production dictionaryExpression is exited.
func (s *BaseObjectiveCParserListener) ExitDictionaryExpression(ctx *DictionaryExpressionContext) {}

// EnterDictionaryPair is called when production dictionaryPair is entered.
func (s *BaseObjectiveCParserListener) EnterDictionaryPair(ctx *DictionaryPairContext) {}

// ExitDictionaryPair is called when production dictionaryPair is exited.
func (s *BaseObjectiveCParserListener) ExitDictionaryPair(ctx *DictionaryPairContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseObjectiveCParserListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseObjectiveCParserListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterBoxExpression is called when production boxExpression is entered.
func (s *BaseObjectiveCParserListener) EnterBoxExpression(ctx *BoxExpressionContext) {}

// ExitBoxExpression is called when production boxExpression is exited.
func (s *BaseObjectiveCParserListener) ExitBoxExpression(ctx *BoxExpressionContext) {}

// EnterBlockParameters is called when production blockParameters is entered.
func (s *BaseObjectiveCParserListener) EnterBlockParameters(ctx *BlockParametersContext) {}

// ExitBlockParameters is called when production blockParameters is exited.
func (s *BaseObjectiveCParserListener) ExitBlockParameters(ctx *BlockParametersContext) {}

// EnterTypeVariableDeclaratorOrName is called when production typeVariableDeclaratorOrName is entered.
func (s *BaseObjectiveCParserListener) EnterTypeVariableDeclaratorOrName(ctx *TypeVariableDeclaratorOrNameContext) {
}

// ExitTypeVariableDeclaratorOrName is called when production typeVariableDeclaratorOrName is exited.
func (s *BaseObjectiveCParserListener) ExitTypeVariableDeclaratorOrName(ctx *TypeVariableDeclaratorOrNameContext) {
}

// EnterBlockExpression is called when production blockExpression is entered.
func (s *BaseObjectiveCParserListener) EnterBlockExpression(ctx *BlockExpressionContext) {}

// ExitBlockExpression is called when production blockExpression is exited.
func (s *BaseObjectiveCParserListener) ExitBlockExpression(ctx *BlockExpressionContext) {}

// EnterMessageExpression is called when production messageExpression is entered.
func (s *BaseObjectiveCParserListener) EnterMessageExpression(ctx *MessageExpressionContext) {}

// ExitMessageExpression is called when production messageExpression is exited.
func (s *BaseObjectiveCParserListener) ExitMessageExpression(ctx *MessageExpressionContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseObjectiveCParserListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseObjectiveCParserListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterMessageSelector is called when production messageSelector is entered.
func (s *BaseObjectiveCParserListener) EnterMessageSelector(ctx *MessageSelectorContext) {}

// ExitMessageSelector is called when production messageSelector is exited.
func (s *BaseObjectiveCParserListener) ExitMessageSelector(ctx *MessageSelectorContext) {}

// EnterKeywordArgument is called when production keywordArgument is entered.
func (s *BaseObjectiveCParserListener) EnterKeywordArgument(ctx *KeywordArgumentContext) {}

// ExitKeywordArgument is called when production keywordArgument is exited.
func (s *BaseObjectiveCParserListener) ExitKeywordArgument(ctx *KeywordArgumentContext) {}

// EnterKeywordArgumentType is called when production keywordArgumentType is entered.
func (s *BaseObjectiveCParserListener) EnterKeywordArgumentType(ctx *KeywordArgumentTypeContext) {}

// ExitKeywordArgumentType is called when production keywordArgumentType is exited.
func (s *BaseObjectiveCParserListener) ExitKeywordArgumentType(ctx *KeywordArgumentTypeContext) {}

// EnterSelectorExpression is called when production selectorExpression is entered.
func (s *BaseObjectiveCParserListener) EnterSelectorExpression(ctx *SelectorExpressionContext) {}

// ExitSelectorExpression is called when production selectorExpression is exited.
func (s *BaseObjectiveCParserListener) ExitSelectorExpression(ctx *SelectorExpressionContext) {}

// EnterSelectorName is called when production selectorName is entered.
func (s *BaseObjectiveCParserListener) EnterSelectorName(ctx *SelectorNameContext) {}

// ExitSelectorName is called when production selectorName is exited.
func (s *BaseObjectiveCParserListener) ExitSelectorName(ctx *SelectorNameContext) {}

// EnterProtocolExpression is called when production protocolExpression is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolExpression(ctx *ProtocolExpressionContext) {}

// ExitProtocolExpression is called when production protocolExpression is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolExpression(ctx *ProtocolExpressionContext) {}

// EnterEncodeExpression is called when production encodeExpression is entered.
func (s *BaseObjectiveCParserListener) EnterEncodeExpression(ctx *EncodeExpressionContext) {}

// ExitEncodeExpression is called when production encodeExpression is exited.
func (s *BaseObjectiveCParserListener) ExitEncodeExpression(ctx *EncodeExpressionContext) {}

// EnterTypeVariableDeclarator is called when production typeVariableDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterTypeVariableDeclarator(ctx *TypeVariableDeclaratorContext) {
}

// ExitTypeVariableDeclarator is called when production typeVariableDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitTypeVariableDeclarator(ctx *TypeVariableDeclaratorContext) {
}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseObjectiveCParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseObjectiveCParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryBlock is called when production tryBlock is entered.
func (s *BaseObjectiveCParserListener) EnterTryBlock(ctx *TryBlockContext) {}

// ExitTryBlock is called when production tryBlock is exited.
func (s *BaseObjectiveCParserListener) ExitTryBlock(ctx *TryBlockContext) {}

// EnterCatchStatement is called when production catchStatement is entered.
func (s *BaseObjectiveCParserListener) EnterCatchStatement(ctx *CatchStatementContext) {}

// ExitCatchStatement is called when production catchStatement is exited.
func (s *BaseObjectiveCParserListener) ExitCatchStatement(ctx *CatchStatementContext) {}

// EnterSynchronizedStatement is called when production synchronizedStatement is entered.
func (s *BaseObjectiveCParserListener) EnterSynchronizedStatement(ctx *SynchronizedStatementContext) {
}

// ExitSynchronizedStatement is called when production synchronizedStatement is exited.
func (s *BaseObjectiveCParserListener) ExitSynchronizedStatement(ctx *SynchronizedStatementContext) {}

// EnterAutoreleaseStatement is called when production autoreleaseStatement is entered.
func (s *BaseObjectiveCParserListener) EnterAutoreleaseStatement(ctx *AutoreleaseStatementContext) {}

// ExitAutoreleaseStatement is called when production autoreleaseStatement is exited.
func (s *BaseObjectiveCParserListener) ExitAutoreleaseStatement(ctx *AutoreleaseStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseObjectiveCParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseObjectiveCParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionSignature is called when production functionSignature is entered.
func (s *BaseObjectiveCParserListener) EnterFunctionSignature(ctx *FunctionSignatureContext) {}

// ExitFunctionSignature is called when production functionSignature is exited.
func (s *BaseObjectiveCParserListener) ExitFunctionSignature(ctx *FunctionSignatureContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseObjectiveCParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseObjectiveCParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterAttributeParameters is called when production attributeParameters is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeParameters(ctx *AttributeParametersContext) {}

// ExitAttributeParameters is called when production attributeParameters is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeParameters(ctx *AttributeParametersContext) {}

// EnterAttributeParameterList is called when production attributeParameterList is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeParameterList(ctx *AttributeParameterListContext) {
}

// ExitAttributeParameterList is called when production attributeParameterList is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeParameterList(ctx *AttributeParameterListContext) {
}

// EnterAttributeParameter is called when production attributeParameter is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeParameter(ctx *AttributeParameterContext) {}

// ExitAttributeParameter is called when production attributeParameter is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeParameter(ctx *AttributeParameterContext) {}

// EnterAttributeParameterAssignment is called when production attributeParameterAssignment is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeParameterAssignment(ctx *AttributeParameterAssignmentContext) {
}

// ExitAttributeParameterAssignment is called when production attributeParameterAssignment is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeParameterAssignment(ctx *AttributeParameterAssignmentContext) {
}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseObjectiveCParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseObjectiveCParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseObjectiveCParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseObjectiveCParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterTypedefDeclaration is called when production typedefDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterTypedefDeclaration(ctx *TypedefDeclarationContext) {}

// ExitTypedefDeclaration is called when production typedefDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitTypedefDeclaration(ctx *TypedefDeclarationContext) {}

// EnterTypeDeclaratorList is called when production typeDeclaratorList is entered.
func (s *BaseObjectiveCParserListener) EnterTypeDeclaratorList(ctx *TypeDeclaratorListContext) {}

// ExitTypeDeclaratorList is called when production typeDeclaratorList is exited.
func (s *BaseObjectiveCParserListener) ExitTypeDeclaratorList(ctx *TypeDeclaratorListContext) {}

// EnterTypeDeclarator is called when production typeDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterTypeDeclarator(ctx *TypeDeclaratorContext) {}

// ExitTypeDeclarator is called when production typeDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitTypeDeclarator(ctx *TypeDeclaratorContext) {}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *BaseObjectiveCParserListener) EnterDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {
}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *BaseObjectiveCParserListener) ExitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {}

// EnterAttributeSpecifier is called when production attributeSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// ExitAttributeSpecifier is called when production attributeSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitAttributeSpecifier(ctx *AttributeSpecifierContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *BaseObjectiveCParserListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *BaseObjectiveCParserListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {
}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *BaseObjectiveCParserListener) EnterSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *BaseObjectiveCParserListener) ExitSpecifierQualifierList(ctx *SpecifierQualifierListContext) {
}

// EnterIbOutletQualifier is called when production ibOutletQualifier is entered.
func (s *BaseObjectiveCParserListener) EnterIbOutletQualifier(ctx *IbOutletQualifierContext) {}

// ExitIbOutletQualifier is called when production ibOutletQualifier is exited.
func (s *BaseObjectiveCParserListener) ExitIbOutletQualifier(ctx *IbOutletQualifierContext) {}

// EnterArcBehaviourSpecifier is called when production arcBehaviourSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterArcBehaviourSpecifier(ctx *ArcBehaviourSpecifierContext) {
}

// ExitArcBehaviourSpecifier is called when production arcBehaviourSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitArcBehaviourSpecifier(ctx *ArcBehaviourSpecifierContext) {}

// EnterNullabilitySpecifier is called when production nullabilitySpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterNullabilitySpecifier(ctx *NullabilitySpecifierContext) {}

// ExitNullabilitySpecifier is called when production nullabilitySpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitNullabilitySpecifier(ctx *NullabilitySpecifierContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {
}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// EnterTypePrefix is called when production typePrefix is entered.
func (s *BaseObjectiveCParserListener) EnterTypePrefix(ctx *TypePrefixContext) {}

// ExitTypePrefix is called when production typePrefix is exited.
func (s *BaseObjectiveCParserListener) ExitTypePrefix(ctx *TypePrefixContext) {}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *BaseObjectiveCParserListener) EnterTypeQualifier(ctx *TypeQualifierContext) {}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *BaseObjectiveCParserListener) ExitTypeQualifier(ctx *TypeQualifierContext) {}

// EnterProtocolQualifier is called when production protocolQualifier is entered.
func (s *BaseObjectiveCParserListener) EnterProtocolQualifier(ctx *ProtocolQualifierContext) {}

// ExitProtocolQualifier is called when production protocolQualifier is exited.
func (s *BaseObjectiveCParserListener) ExitProtocolQualifier(ctx *ProtocolQualifierContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterTypeofExpression is called when production typeofExpression is entered.
func (s *BaseObjectiveCParserListener) EnterTypeofExpression(ctx *TypeofExpressionContext) {}

// ExitTypeofExpression is called when production typeofExpression is exited.
func (s *BaseObjectiveCParserListener) ExitTypeofExpression(ctx *TypeofExpressionContext) {}

// EnterFieldDeclaratorList is called when production fieldDeclaratorList is entered.
func (s *BaseObjectiveCParserListener) EnterFieldDeclaratorList(ctx *FieldDeclaratorListContext) {}

// ExitFieldDeclaratorList is called when production fieldDeclaratorList is exited.
func (s *BaseObjectiveCParserListener) ExitFieldDeclaratorList(ctx *FieldDeclaratorListContext) {}

// EnterFieldDeclarator is called when production fieldDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterFieldDeclarator(ctx *FieldDeclaratorContext) {}

// ExitFieldDeclarator is called when production fieldDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitFieldDeclarator(ctx *FieldDeclaratorContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *BaseObjectiveCParserListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *BaseObjectiveCParserListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *BaseObjectiveCParserListener) EnterEnumeratorList(ctx *EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *BaseObjectiveCParserListener) ExitEnumeratorList(ctx *EnumeratorListContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseObjectiveCParserListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseObjectiveCParserListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterEnumeratorIdentifier is called when production enumeratorIdentifier is entered.
func (s *BaseObjectiveCParserListener) EnterEnumeratorIdentifier(ctx *EnumeratorIdentifierContext) {}

// ExitEnumeratorIdentifier is called when production enumeratorIdentifier is exited.
func (s *BaseObjectiveCParserListener) ExitEnumeratorIdentifier(ctx *EnumeratorIdentifierContext) {}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterDirectDeclarator(ctx *DirectDeclaratorContext) {}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {}

// EnterDeclaratorSuffix is called when production declaratorSuffix is entered.
func (s *BaseObjectiveCParserListener) EnterDeclaratorSuffix(ctx *DeclaratorSuffixContext) {}

// ExitDeclaratorSuffix is called when production declaratorSuffix is exited.
func (s *BaseObjectiveCParserListener) ExitDeclaratorSuffix(ctx *DeclaratorSuffixContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseObjectiveCParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseObjectiveCParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseObjectiveCParserListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BaseObjectiveCParserListener) ExitPointer(ctx *PointerContext) {}

// EnterMacro is called when production macro is entered.
func (s *BaseObjectiveCParserListener) EnterMacro(ctx *MacroContext) {}

// ExitMacro is called when production macro is exited.
func (s *BaseObjectiveCParserListener) ExitMacro(ctx *MacroContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseObjectiveCParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseObjectiveCParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterStructInitializer is called when production structInitializer is entered.
func (s *BaseObjectiveCParserListener) EnterStructInitializer(ctx *StructInitializerContext) {}

// ExitStructInitializer is called when production structInitializer is exited.
func (s *BaseObjectiveCParserListener) ExitStructInitializer(ctx *StructInitializerContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BaseObjectiveCParserListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BaseObjectiveCParserListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseObjectiveCParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseObjectiveCParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *BaseObjectiveCParserListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *BaseObjectiveCParserListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// EnterAbstractDeclaratorSuffix is called when production abstractDeclaratorSuffix is entered.
func (s *BaseObjectiveCParserListener) EnterAbstractDeclaratorSuffix(ctx *AbstractDeclaratorSuffixContext) {
}

// ExitAbstractDeclaratorSuffix is called when production abstractDeclaratorSuffix is exited.
func (s *BaseObjectiveCParserListener) ExitAbstractDeclaratorSuffix(ctx *AbstractDeclaratorSuffixContext) {
}

// EnterParameterDeclarationList is called when production parameterDeclarationList is entered.
func (s *BaseObjectiveCParserListener) EnterParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// ExitParameterDeclarationList is called when production parameterDeclarationList is exited.
func (s *BaseObjectiveCParserListener) ExitParameterDeclarationList(ctx *ParameterDeclarationListContext) {
}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseObjectiveCParserListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseObjectiveCParserListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseObjectiveCParserListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseObjectiveCParserListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseObjectiveCParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseObjectiveCParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseObjectiveCParserListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseObjectiveCParserListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseObjectiveCParserListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseObjectiveCParserListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseObjectiveCParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseObjectiveCParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseObjectiveCParserListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseObjectiveCParserListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseObjectiveCParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseObjectiveCParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchBlock is called when production switchBlock is entered.
func (s *BaseObjectiveCParserListener) EnterSwitchBlock(ctx *SwitchBlockContext) {}

// ExitSwitchBlock is called when production switchBlock is exited.
func (s *BaseObjectiveCParserListener) ExitSwitchBlock(ctx *SwitchBlockContext) {}

// EnterSwitchSection is called when production switchSection is entered.
func (s *BaseObjectiveCParserListener) EnterSwitchSection(ctx *SwitchSectionContext) {}

// ExitSwitchSection is called when production switchSection is exited.
func (s *BaseObjectiveCParserListener) ExitSwitchSection(ctx *SwitchSectionContext) {}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseObjectiveCParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseObjectiveCParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseObjectiveCParserListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseObjectiveCParserListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseObjectiveCParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseObjectiveCParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseObjectiveCParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseObjectiveCParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseObjectiveCParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseObjectiveCParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForLoopInitializer is called when production forLoopInitializer is entered.
func (s *BaseObjectiveCParserListener) EnterForLoopInitializer(ctx *ForLoopInitializerContext) {}

// ExitForLoopInitializer is called when production forLoopInitializer is exited.
func (s *BaseObjectiveCParserListener) ExitForLoopInitializer(ctx *ForLoopInitializerContext) {}

// EnterForInStatement is called when production forInStatement is entered.
func (s *BaseObjectiveCParserListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production forInStatement is exited.
func (s *BaseObjectiveCParserListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BaseObjectiveCParserListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BaseObjectiveCParserListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseObjectiveCParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseObjectiveCParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseObjectiveCParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseObjectiveCParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseObjectiveCParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseObjectiveCParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseObjectiveCParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseObjectiveCParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseObjectiveCParserListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseObjectiveCParserListener) ExitInitializer(ctx *InitializerContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseObjectiveCParserListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseObjectiveCParserListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseObjectiveCParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseObjectiveCParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseObjectiveCParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseObjectiveCParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseObjectiveCParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseObjectiveCParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostfix is called when production postfix is entered.
func (s *BaseObjectiveCParserListener) EnterPostfix(ctx *PostfixContext) {}

// ExitPostfix is called when production postfix is exited.
func (s *BaseObjectiveCParserListener) ExitPostfix(ctx *PostfixContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *BaseObjectiveCParserListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {
}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *BaseObjectiveCParserListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {
}

// EnterArgumentExpression is called when production argumentExpression is entered.
func (s *BaseObjectiveCParserListener) EnterArgumentExpression(ctx *ArgumentExpressionContext) {}

// ExitArgumentExpression is called when production argumentExpression is exited.
func (s *BaseObjectiveCParserListener) ExitArgumentExpression(ctx *ArgumentExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseObjectiveCParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseObjectiveCParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseObjectiveCParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseObjectiveCParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseObjectiveCParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseObjectiveCParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseObjectiveCParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseObjectiveCParserListener) ExitIdentifier(ctx *IdentifierContext) {}
