// Code generated from ObjectiveCParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package objectivec // ObjectiveCParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ObjectiveCParserListener is a complete listener for a parse tree produced by ObjectiveCParser.
type ObjectiveCParserListener interface {
	antlr.ParseTreeListener

	// EnterTranslationUnit is called when entering the translationUnit production.
	EnterTranslationUnit(c *TranslationUnitContext)

	// EnterTopLevelDeclaration is called when entering the topLevelDeclaration production.
	EnterTopLevelDeclaration(c *TopLevelDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterClassInterface is called when entering the classInterface production.
	EnterClassInterface(c *ClassInterfaceContext)

	// EnterCategoryInterface is called when entering the categoryInterface production.
	EnterCategoryInterface(c *CategoryInterfaceContext)

	// EnterClassImplementation is called when entering the classImplementation production.
	EnterClassImplementation(c *ClassImplementationContext)

	// EnterCategoryImplementation is called when entering the categoryImplementation production.
	EnterCategoryImplementation(c *CategoryImplementationContext)

	// EnterGenericTypeSpecifier is called when entering the genericTypeSpecifier production.
	EnterGenericTypeSpecifier(c *GenericTypeSpecifierContext)

	// EnterProtocolDeclaration is called when entering the protocolDeclaration production.
	EnterProtocolDeclaration(c *ProtocolDeclarationContext)

	// EnterProtocolDeclarationSection is called when entering the protocolDeclarationSection production.
	EnterProtocolDeclarationSection(c *ProtocolDeclarationSectionContext)

	// EnterProtocolDeclarationList is called when entering the protocolDeclarationList production.
	EnterProtocolDeclarationList(c *ProtocolDeclarationListContext)

	// EnterClassDeclarationList is called when entering the classDeclarationList production.
	EnterClassDeclarationList(c *ClassDeclarationListContext)

	// EnterProtocolList is called when entering the protocolList production.
	EnterProtocolList(c *ProtocolListContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyAttributesList is called when entering the propertyAttributesList production.
	EnterPropertyAttributesList(c *PropertyAttributesListContext)

	// EnterPropertyAttribute is called when entering the propertyAttribute production.
	EnterPropertyAttribute(c *PropertyAttributeContext)

	// EnterProtocolName is called when entering the protocolName production.
	EnterProtocolName(c *ProtocolNameContext)

	// EnterInstanceVariables is called when entering the instanceVariables production.
	EnterInstanceVariables(c *InstanceVariablesContext)

	// EnterVisibilitySection is called when entering the visibilitySection production.
	EnterVisibilitySection(c *VisibilitySectionContext)

	// EnterAccessModifier is called when entering the accessModifier production.
	EnterAccessModifier(c *AccessModifierContext)

	// EnterInterfaceDeclarationList is called when entering the interfaceDeclarationList production.
	EnterInterfaceDeclarationList(c *InterfaceDeclarationListContext)

	// EnterClassMethodDeclaration is called when entering the classMethodDeclaration production.
	EnterClassMethodDeclaration(c *ClassMethodDeclarationContext)

	// EnterInstanceMethodDeclaration is called when entering the instanceMethodDeclaration production.
	EnterInstanceMethodDeclaration(c *InstanceMethodDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterImplementationDefinitionList is called when entering the implementationDefinitionList production.
	EnterImplementationDefinitionList(c *ImplementationDefinitionListContext)

	// EnterClassMethodDefinition is called when entering the classMethodDefinition production.
	EnterClassMethodDefinition(c *ClassMethodDefinitionContext)

	// EnterInstanceMethodDefinition is called when entering the instanceMethodDefinition production.
	EnterInstanceMethodDefinition(c *InstanceMethodDefinitionContext)

	// EnterMethodDefinition is called when entering the methodDefinition production.
	EnterMethodDefinition(c *MethodDefinitionContext)

	// EnterMethodSelector is called when entering the methodSelector production.
	EnterMethodSelector(c *MethodSelectorContext)

	// EnterKeywordDeclarator is called when entering the keywordDeclarator production.
	EnterKeywordDeclarator(c *KeywordDeclaratorContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterMethodType is called when entering the methodType production.
	EnterMethodType(c *MethodTypeContext)

	// EnterPropertyImplementation is called when entering the propertyImplementation production.
	EnterPropertyImplementation(c *PropertyImplementationContext)

	// EnterPropertySynthesizeList is called when entering the propertySynthesizeList production.
	EnterPropertySynthesizeList(c *PropertySynthesizeListContext)

	// EnterPropertySynthesizeItem is called when entering the propertySynthesizeItem production.
	EnterPropertySynthesizeItem(c *PropertySynthesizeItemContext)

	// EnterBlockType is called when entering the blockType production.
	EnterBlockType(c *BlockTypeContext)

	// EnterGenericsSpecifier is called when entering the genericsSpecifier production.
	EnterGenericsSpecifier(c *GenericsSpecifierContext)

	// EnterTypeSpecifierWithPrefixes is called when entering the typeSpecifierWithPrefixes production.
	EnterTypeSpecifierWithPrefixes(c *TypeSpecifierWithPrefixesContext)

	// EnterDictionaryExpression is called when entering the dictionaryExpression production.
	EnterDictionaryExpression(c *DictionaryExpressionContext)

	// EnterDictionaryPair is called when entering the dictionaryPair production.
	EnterDictionaryPair(c *DictionaryPairContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterBoxExpression is called when entering the boxExpression production.
	EnterBoxExpression(c *BoxExpressionContext)

	// EnterBlockParameters is called when entering the blockParameters production.
	EnterBlockParameters(c *BlockParametersContext)

	// EnterTypeVariableDeclaratorOrName is called when entering the typeVariableDeclaratorOrName production.
	EnterTypeVariableDeclaratorOrName(c *TypeVariableDeclaratorOrNameContext)

	// EnterBlockExpression is called when entering the blockExpression production.
	EnterBlockExpression(c *BlockExpressionContext)

	// EnterMessageExpression is called when entering the messageExpression production.
	EnterMessageExpression(c *MessageExpressionContext)

	// EnterReceiver is called when entering the receiver production.
	EnterReceiver(c *ReceiverContext)

	// EnterMessageSelector is called when entering the messageSelector production.
	EnterMessageSelector(c *MessageSelectorContext)

	// EnterKeywordArgument is called when entering the keywordArgument production.
	EnterKeywordArgument(c *KeywordArgumentContext)

	// EnterKeywordArgumentType is called when entering the keywordArgumentType production.
	EnterKeywordArgumentType(c *KeywordArgumentTypeContext)

	// EnterSelectorExpression is called when entering the selectorExpression production.
	EnterSelectorExpression(c *SelectorExpressionContext)

	// EnterSelectorName is called when entering the selectorName production.
	EnterSelectorName(c *SelectorNameContext)

	// EnterProtocolExpression is called when entering the protocolExpression production.
	EnterProtocolExpression(c *ProtocolExpressionContext)

	// EnterEncodeExpression is called when entering the encodeExpression production.
	EnterEncodeExpression(c *EncodeExpressionContext)

	// EnterTypeVariableDeclarator is called when entering the typeVariableDeclarator production.
	EnterTypeVariableDeclarator(c *TypeVariableDeclaratorContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterTryBlock is called when entering the tryBlock production.
	EnterTryBlock(c *TryBlockContext)

	// EnterCatchStatement is called when entering the catchStatement production.
	EnterCatchStatement(c *CatchStatementContext)

	// EnterSynchronizedStatement is called when entering the synchronizedStatement production.
	EnterSynchronizedStatement(c *SynchronizedStatementContext)

	// EnterAutoreleaseStatement is called when entering the autoreleaseStatement production.
	EnterAutoreleaseStatement(c *AutoreleaseStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionSignature is called when entering the functionSignature production.
	EnterFunctionSignature(c *FunctionSignatureContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttributeName is called when entering the attributeName production.
	EnterAttributeName(c *AttributeNameContext)

	// EnterAttributeParameters is called when entering the attributeParameters production.
	EnterAttributeParameters(c *AttributeParametersContext)

	// EnterAttributeParameterList is called when entering the attributeParameterList production.
	EnterAttributeParameterList(c *AttributeParameterListContext)

	// EnterAttributeParameter is called when entering the attributeParameter production.
	EnterAttributeParameter(c *AttributeParameterContext)

	// EnterAttributeParameterAssignment is called when entering the attributeParameterAssignment production.
	EnterAttributeParameterAssignment(c *AttributeParameterAssignmentContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterTypedefDeclaration is called when entering the typedefDeclaration production.
	EnterTypedefDeclaration(c *TypedefDeclarationContext)

	// EnterTypeDeclaratorList is called when entering the typeDeclaratorList production.
	EnterTypeDeclaratorList(c *TypeDeclaratorListContext)

	// EnterTypeDeclarator is called when entering the typeDeclarator production.
	EnterTypeDeclarator(c *TypeDeclaratorContext)

	// EnterDeclarationSpecifiers is called when entering the declarationSpecifiers production.
	EnterDeclarationSpecifiers(c *DeclarationSpecifiersContext)

	// EnterAttributeSpecifier is called when entering the attributeSpecifier production.
	EnterAttributeSpecifier(c *AttributeSpecifierContext)

	// EnterInitDeclaratorList is called when entering the initDeclaratorList production.
	EnterInitDeclaratorList(c *InitDeclaratorListContext)

	// EnterInitDeclarator is called when entering the initDeclarator production.
	EnterInitDeclarator(c *InitDeclaratorContext)

	// EnterStructOrUnionSpecifier is called when entering the structOrUnionSpecifier production.
	EnterStructOrUnionSpecifier(c *StructOrUnionSpecifierContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterSpecifierQualifierList is called when entering the specifierQualifierList production.
	EnterSpecifierQualifierList(c *SpecifierQualifierListContext)

	// EnterIbOutletQualifier is called when entering the ibOutletQualifier production.
	EnterIbOutletQualifier(c *IbOutletQualifierContext)

	// EnterArcBehaviourSpecifier is called when entering the arcBehaviourSpecifier production.
	EnterArcBehaviourSpecifier(c *ArcBehaviourSpecifierContext)

	// EnterNullabilitySpecifier is called when entering the nullabilitySpecifier production.
	EnterNullabilitySpecifier(c *NullabilitySpecifierContext)

	// EnterStorageClassSpecifier is called when entering the storageClassSpecifier production.
	EnterStorageClassSpecifier(c *StorageClassSpecifierContext)

	// EnterTypePrefix is called when entering the typePrefix production.
	EnterTypePrefix(c *TypePrefixContext)

	// EnterTypeQualifier is called when entering the typeQualifier production.
	EnterTypeQualifier(c *TypeQualifierContext)

	// EnterProtocolQualifier is called when entering the protocolQualifier production.
	EnterProtocolQualifier(c *ProtocolQualifierContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterTypeofExpression is called when entering the typeofExpression production.
	EnterTypeofExpression(c *TypeofExpressionContext)

	// EnterFieldDeclaratorList is called when entering the fieldDeclaratorList production.
	EnterFieldDeclaratorList(c *FieldDeclaratorListContext)

	// EnterFieldDeclarator is called when entering the fieldDeclarator production.
	EnterFieldDeclarator(c *FieldDeclaratorContext)

	// EnterEnumSpecifier is called when entering the enumSpecifier production.
	EnterEnumSpecifier(c *EnumSpecifierContext)

	// EnterEnumeratorList is called when entering the enumeratorList production.
	EnterEnumeratorList(c *EnumeratorListContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterEnumeratorIdentifier is called when entering the enumeratorIdentifier production.
	EnterEnumeratorIdentifier(c *EnumeratorIdentifierContext)

	// EnterDirectDeclarator is called when entering the directDeclarator production.
	EnterDirectDeclarator(c *DirectDeclaratorContext)

	// EnterDeclaratorSuffix is called when entering the declaratorSuffix production.
	EnterDeclaratorSuffix(c *DeclaratorSuffixContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterMacro is called when entering the macro production.
	EnterMacro(c *MacroContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterStructInitializer is called when entering the structInitializer production.
	EnterStructInitializer(c *StructInitializerContext)

	// EnterInitializerList is called when entering the initializerList production.
	EnterInitializerList(c *InitializerListContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterAbstractDeclarator is called when entering the abstractDeclarator production.
	EnterAbstractDeclarator(c *AbstractDeclaratorContext)

	// EnterAbstractDeclaratorSuffix is called when entering the abstractDeclaratorSuffix production.
	EnterAbstractDeclaratorSuffix(c *AbstractDeclaratorSuffixContext)

	// EnterParameterDeclarationList is called when entering the parameterDeclarationList production.
	EnterParameterDeclarationList(c *ParameterDeclarationListContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeledStatement is called when entering the labeledStatement production.
	EnterLabeledStatement(c *LabeledStatementContext)

	// EnterRangeExpression is called when entering the rangeExpression production.
	EnterRangeExpression(c *RangeExpressionContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchBlock is called when entering the switchBlock production.
	EnterSwitchBlock(c *SwitchBlockContext)

	// EnterSwitchSection is called when entering the switchSection production.
	EnterSwitchSection(c *SwitchSectionContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForLoopInitializer is called when entering the forLoopInitializer production.
	EnterForLoopInitializer(c *ForLoopInitializerContext)

	// EnterForInStatement is called when entering the forInStatement production.
	EnterForInStatement(c *ForInStatementContext)

	// EnterJumpStatement is called when entering the jumpStatement production.
	EnterJumpStatement(c *JumpStatementContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostfix is called when entering the postfix production.
	EnterPostfix(c *PostfixContext)

	// EnterArgumentExpressionList is called when entering the argumentExpressionList production.
	EnterArgumentExpressionList(c *ArgumentExpressionListContext)

	// EnterArgumentExpression is called when entering the argumentExpression production.
	EnterArgumentExpression(c *ArgumentExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitTranslationUnit is called when exiting the translationUnit production.
	ExitTranslationUnit(c *TranslationUnitContext)

	// ExitTopLevelDeclaration is called when exiting the topLevelDeclaration production.
	ExitTopLevelDeclaration(c *TopLevelDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitClassInterface is called when exiting the classInterface production.
	ExitClassInterface(c *ClassInterfaceContext)

	// ExitCategoryInterface is called when exiting the categoryInterface production.
	ExitCategoryInterface(c *CategoryInterfaceContext)

	// ExitClassImplementation is called when exiting the classImplementation production.
	ExitClassImplementation(c *ClassImplementationContext)

	// ExitCategoryImplementation is called when exiting the categoryImplementation production.
	ExitCategoryImplementation(c *CategoryImplementationContext)

	// ExitGenericTypeSpecifier is called when exiting the genericTypeSpecifier production.
	ExitGenericTypeSpecifier(c *GenericTypeSpecifierContext)

	// ExitProtocolDeclaration is called when exiting the protocolDeclaration production.
	ExitProtocolDeclaration(c *ProtocolDeclarationContext)

	// ExitProtocolDeclarationSection is called when exiting the protocolDeclarationSection production.
	ExitProtocolDeclarationSection(c *ProtocolDeclarationSectionContext)

	// ExitProtocolDeclarationList is called when exiting the protocolDeclarationList production.
	ExitProtocolDeclarationList(c *ProtocolDeclarationListContext)

	// ExitClassDeclarationList is called when exiting the classDeclarationList production.
	ExitClassDeclarationList(c *ClassDeclarationListContext)

	// ExitProtocolList is called when exiting the protocolList production.
	ExitProtocolList(c *ProtocolListContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyAttributesList is called when exiting the propertyAttributesList production.
	ExitPropertyAttributesList(c *PropertyAttributesListContext)

	// ExitPropertyAttribute is called when exiting the propertyAttribute production.
	ExitPropertyAttribute(c *PropertyAttributeContext)

	// ExitProtocolName is called when exiting the protocolName production.
	ExitProtocolName(c *ProtocolNameContext)

	// ExitInstanceVariables is called when exiting the instanceVariables production.
	ExitInstanceVariables(c *InstanceVariablesContext)

	// ExitVisibilitySection is called when exiting the visibilitySection production.
	ExitVisibilitySection(c *VisibilitySectionContext)

	// ExitAccessModifier is called when exiting the accessModifier production.
	ExitAccessModifier(c *AccessModifierContext)

	// ExitInterfaceDeclarationList is called when exiting the interfaceDeclarationList production.
	ExitInterfaceDeclarationList(c *InterfaceDeclarationListContext)

	// ExitClassMethodDeclaration is called when exiting the classMethodDeclaration production.
	ExitClassMethodDeclaration(c *ClassMethodDeclarationContext)

	// ExitInstanceMethodDeclaration is called when exiting the instanceMethodDeclaration production.
	ExitInstanceMethodDeclaration(c *InstanceMethodDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitImplementationDefinitionList is called when exiting the implementationDefinitionList production.
	ExitImplementationDefinitionList(c *ImplementationDefinitionListContext)

	// ExitClassMethodDefinition is called when exiting the classMethodDefinition production.
	ExitClassMethodDefinition(c *ClassMethodDefinitionContext)

	// ExitInstanceMethodDefinition is called when exiting the instanceMethodDefinition production.
	ExitInstanceMethodDefinition(c *InstanceMethodDefinitionContext)

	// ExitMethodDefinition is called when exiting the methodDefinition production.
	ExitMethodDefinition(c *MethodDefinitionContext)

	// ExitMethodSelector is called when exiting the methodSelector production.
	ExitMethodSelector(c *MethodSelectorContext)

	// ExitKeywordDeclarator is called when exiting the keywordDeclarator production.
	ExitKeywordDeclarator(c *KeywordDeclaratorContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitMethodType is called when exiting the methodType production.
	ExitMethodType(c *MethodTypeContext)

	// ExitPropertyImplementation is called when exiting the propertyImplementation production.
	ExitPropertyImplementation(c *PropertyImplementationContext)

	// ExitPropertySynthesizeList is called when exiting the propertySynthesizeList production.
	ExitPropertySynthesizeList(c *PropertySynthesizeListContext)

	// ExitPropertySynthesizeItem is called when exiting the propertySynthesizeItem production.
	ExitPropertySynthesizeItem(c *PropertySynthesizeItemContext)

	// ExitBlockType is called when exiting the blockType production.
	ExitBlockType(c *BlockTypeContext)

	// ExitGenericsSpecifier is called when exiting the genericsSpecifier production.
	ExitGenericsSpecifier(c *GenericsSpecifierContext)

	// ExitTypeSpecifierWithPrefixes is called when exiting the typeSpecifierWithPrefixes production.
	ExitTypeSpecifierWithPrefixes(c *TypeSpecifierWithPrefixesContext)

	// ExitDictionaryExpression is called when exiting the dictionaryExpression production.
	ExitDictionaryExpression(c *DictionaryExpressionContext)

	// ExitDictionaryPair is called when exiting the dictionaryPair production.
	ExitDictionaryPair(c *DictionaryPairContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitBoxExpression is called when exiting the boxExpression production.
	ExitBoxExpression(c *BoxExpressionContext)

	// ExitBlockParameters is called when exiting the blockParameters production.
	ExitBlockParameters(c *BlockParametersContext)

	// ExitTypeVariableDeclaratorOrName is called when exiting the typeVariableDeclaratorOrName production.
	ExitTypeVariableDeclaratorOrName(c *TypeVariableDeclaratorOrNameContext)

	// ExitBlockExpression is called when exiting the blockExpression production.
	ExitBlockExpression(c *BlockExpressionContext)

	// ExitMessageExpression is called when exiting the messageExpression production.
	ExitMessageExpression(c *MessageExpressionContext)

	// ExitReceiver is called when exiting the receiver production.
	ExitReceiver(c *ReceiverContext)

	// ExitMessageSelector is called when exiting the messageSelector production.
	ExitMessageSelector(c *MessageSelectorContext)

	// ExitKeywordArgument is called when exiting the keywordArgument production.
	ExitKeywordArgument(c *KeywordArgumentContext)

	// ExitKeywordArgumentType is called when exiting the keywordArgumentType production.
	ExitKeywordArgumentType(c *KeywordArgumentTypeContext)

	// ExitSelectorExpression is called when exiting the selectorExpression production.
	ExitSelectorExpression(c *SelectorExpressionContext)

	// ExitSelectorName is called when exiting the selectorName production.
	ExitSelectorName(c *SelectorNameContext)

	// ExitProtocolExpression is called when exiting the protocolExpression production.
	ExitProtocolExpression(c *ProtocolExpressionContext)

	// ExitEncodeExpression is called when exiting the encodeExpression production.
	ExitEncodeExpression(c *EncodeExpressionContext)

	// ExitTypeVariableDeclarator is called when exiting the typeVariableDeclarator production.
	ExitTypeVariableDeclarator(c *TypeVariableDeclaratorContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitTryBlock is called when exiting the tryBlock production.
	ExitTryBlock(c *TryBlockContext)

	// ExitCatchStatement is called when exiting the catchStatement production.
	ExitCatchStatement(c *CatchStatementContext)

	// ExitSynchronizedStatement is called when exiting the synchronizedStatement production.
	ExitSynchronizedStatement(c *SynchronizedStatementContext)

	// ExitAutoreleaseStatement is called when exiting the autoreleaseStatement production.
	ExitAutoreleaseStatement(c *AutoreleaseStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionSignature is called when exiting the functionSignature production.
	ExitFunctionSignature(c *FunctionSignatureContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttributeName is called when exiting the attributeName production.
	ExitAttributeName(c *AttributeNameContext)

	// ExitAttributeParameters is called when exiting the attributeParameters production.
	ExitAttributeParameters(c *AttributeParametersContext)

	// ExitAttributeParameterList is called when exiting the attributeParameterList production.
	ExitAttributeParameterList(c *AttributeParameterListContext)

	// ExitAttributeParameter is called when exiting the attributeParameter production.
	ExitAttributeParameter(c *AttributeParameterContext)

	// ExitAttributeParameterAssignment is called when exiting the attributeParameterAssignment production.
	ExitAttributeParameterAssignment(c *AttributeParameterAssignmentContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitTypedefDeclaration is called when exiting the typedefDeclaration production.
	ExitTypedefDeclaration(c *TypedefDeclarationContext)

	// ExitTypeDeclaratorList is called when exiting the typeDeclaratorList production.
	ExitTypeDeclaratorList(c *TypeDeclaratorListContext)

	// ExitTypeDeclarator is called when exiting the typeDeclarator production.
	ExitTypeDeclarator(c *TypeDeclaratorContext)

	// ExitDeclarationSpecifiers is called when exiting the declarationSpecifiers production.
	ExitDeclarationSpecifiers(c *DeclarationSpecifiersContext)

	// ExitAttributeSpecifier is called when exiting the attributeSpecifier production.
	ExitAttributeSpecifier(c *AttributeSpecifierContext)

	// ExitInitDeclaratorList is called when exiting the initDeclaratorList production.
	ExitInitDeclaratorList(c *InitDeclaratorListContext)

	// ExitInitDeclarator is called when exiting the initDeclarator production.
	ExitInitDeclarator(c *InitDeclaratorContext)

	// ExitStructOrUnionSpecifier is called when exiting the structOrUnionSpecifier production.
	ExitStructOrUnionSpecifier(c *StructOrUnionSpecifierContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitSpecifierQualifierList is called when exiting the specifierQualifierList production.
	ExitSpecifierQualifierList(c *SpecifierQualifierListContext)

	// ExitIbOutletQualifier is called when exiting the ibOutletQualifier production.
	ExitIbOutletQualifier(c *IbOutletQualifierContext)

	// ExitArcBehaviourSpecifier is called when exiting the arcBehaviourSpecifier production.
	ExitArcBehaviourSpecifier(c *ArcBehaviourSpecifierContext)

	// ExitNullabilitySpecifier is called when exiting the nullabilitySpecifier production.
	ExitNullabilitySpecifier(c *NullabilitySpecifierContext)

	// ExitStorageClassSpecifier is called when exiting the storageClassSpecifier production.
	ExitStorageClassSpecifier(c *StorageClassSpecifierContext)

	// ExitTypePrefix is called when exiting the typePrefix production.
	ExitTypePrefix(c *TypePrefixContext)

	// ExitTypeQualifier is called when exiting the typeQualifier production.
	ExitTypeQualifier(c *TypeQualifierContext)

	// ExitProtocolQualifier is called when exiting the protocolQualifier production.
	ExitProtocolQualifier(c *ProtocolQualifierContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitTypeofExpression is called when exiting the typeofExpression production.
	ExitTypeofExpression(c *TypeofExpressionContext)

	// ExitFieldDeclaratorList is called when exiting the fieldDeclaratorList production.
	ExitFieldDeclaratorList(c *FieldDeclaratorListContext)

	// ExitFieldDeclarator is called when exiting the fieldDeclarator production.
	ExitFieldDeclarator(c *FieldDeclaratorContext)

	// ExitEnumSpecifier is called when exiting the enumSpecifier production.
	ExitEnumSpecifier(c *EnumSpecifierContext)

	// ExitEnumeratorList is called when exiting the enumeratorList production.
	ExitEnumeratorList(c *EnumeratorListContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitEnumeratorIdentifier is called when exiting the enumeratorIdentifier production.
	ExitEnumeratorIdentifier(c *EnumeratorIdentifierContext)

	// ExitDirectDeclarator is called when exiting the directDeclarator production.
	ExitDirectDeclarator(c *DirectDeclaratorContext)

	// ExitDeclaratorSuffix is called when exiting the declaratorSuffix production.
	ExitDeclaratorSuffix(c *DeclaratorSuffixContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitMacro is called when exiting the macro production.
	ExitMacro(c *MacroContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitStructInitializer is called when exiting the structInitializer production.
	ExitStructInitializer(c *StructInitializerContext)

	// ExitInitializerList is called when exiting the initializerList production.
	ExitInitializerList(c *InitializerListContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitAbstractDeclarator is called when exiting the abstractDeclarator production.
	ExitAbstractDeclarator(c *AbstractDeclaratorContext)

	// ExitAbstractDeclaratorSuffix is called when exiting the abstractDeclaratorSuffix production.
	ExitAbstractDeclaratorSuffix(c *AbstractDeclaratorSuffixContext)

	// ExitParameterDeclarationList is called when exiting the parameterDeclarationList production.
	ExitParameterDeclarationList(c *ParameterDeclarationListContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeledStatement is called when exiting the labeledStatement production.
	ExitLabeledStatement(c *LabeledStatementContext)

	// ExitRangeExpression is called when exiting the rangeExpression production.
	ExitRangeExpression(c *RangeExpressionContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchBlock is called when exiting the switchBlock production.
	ExitSwitchBlock(c *SwitchBlockContext)

	// ExitSwitchSection is called when exiting the switchSection production.
	ExitSwitchSection(c *SwitchSectionContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForLoopInitializer is called when exiting the forLoopInitializer production.
	ExitForLoopInitializer(c *ForLoopInitializerContext)

	// ExitForInStatement is called when exiting the forInStatement production.
	ExitForInStatement(c *ForInStatementContext)

	// ExitJumpStatement is called when exiting the jumpStatement production.
	ExitJumpStatement(c *JumpStatementContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostfix is called when exiting the postfix production.
	ExitPostfix(c *PostfixContext)

	// ExitArgumentExpressionList is called when exiting the argumentExpressionList production.
	ExitArgumentExpressionList(c *ArgumentExpressionListContext)

	// ExitArgumentExpression is called when exiting the argumentExpression production.
	ExitArgumentExpression(c *ArgumentExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
