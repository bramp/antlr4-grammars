// Generated from C.g4 by ANTLR 4.7.

package c // C
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCListener is a complete listener for a parse tree produced by CParser.
type BaseCListener struct{}

var _ CListener = &BaseCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseCListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseCListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterGenericSelection is called when production genericSelection is entered.
func (s *BaseCListener) EnterGenericSelection(ctx *GenericSelectionContext) {}

// ExitGenericSelection is called when production genericSelection is exited.
func (s *BaseCListener) ExitGenericSelection(ctx *GenericSelectionContext) {}

// EnterGenericAssocList is called when production genericAssocList is entered.
func (s *BaseCListener) EnterGenericAssocList(ctx *GenericAssocListContext) {}

// ExitGenericAssocList is called when production genericAssocList is exited.
func (s *BaseCListener) ExitGenericAssocList(ctx *GenericAssocListContext) {}

// EnterGenericAssociation is called when production genericAssociation is entered.
func (s *BaseCListener) EnterGenericAssociation(ctx *GenericAssociationContext) {}

// ExitGenericAssociation is called when production genericAssociation is exited.
func (s *BaseCListener) ExitGenericAssociation(ctx *GenericAssociationContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseCListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseCListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterArgumentExpressionList is called when production argumentExpressionList is entered.
func (s *BaseCListener) EnterArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// ExitArgumentExpressionList is called when production argumentExpressionList is exited.
func (s *BaseCListener) ExitArgumentExpressionList(ctx *ArgumentExpressionListContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseCListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseCListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseCListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseCListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseCListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseCListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseCListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseCListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseCListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseCListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterShiftExpression is called when production shiftExpression is entered.
func (s *BaseCListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production shiftExpression is exited.
func (s *BaseCListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseCListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseCListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseCListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseCListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseCListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseCListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterExclusiveOrExpression is called when production exclusiveOrExpression is entered.
func (s *BaseCListener) EnterExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// ExitExclusiveOrExpression is called when production exclusiveOrExpression is exited.
func (s *BaseCListener) ExitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) {}

// EnterInclusiveOrExpression is called when production inclusiveOrExpression is entered.
func (s *BaseCListener) EnterInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// ExitInclusiveOrExpression is called when production inclusiveOrExpression is exited.
func (s *BaseCListener) ExitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) {}

// EnterLogicalAndExpression is called when production logicalAndExpression is entered.
func (s *BaseCListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production logicalAndExpression is exited.
func (s *BaseCListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterLogicalOrExpression is called when production logicalOrExpression is entered.
func (s *BaseCListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production logicalOrExpression is exited.
func (s *BaseCListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseCListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseCListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseCListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseCListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseCListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseCListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstantExpression is called when production constantExpression is entered.
func (s *BaseCListener) EnterConstantExpression(ctx *ConstantExpressionContext) {}

// ExitConstantExpression is called when production constantExpression is exited.
func (s *BaseCListener) ExitConstantExpression(ctx *ConstantExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclarationSpecifiers is called when production declarationSpecifiers is entered.
func (s *BaseCListener) EnterDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {}

// ExitDeclarationSpecifiers is called when production declarationSpecifiers is exited.
func (s *BaseCListener) ExitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) {}

// EnterDeclarationSpecifiers2 is called when production declarationSpecifiers2 is entered.
func (s *BaseCListener) EnterDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {}

// ExitDeclarationSpecifiers2 is called when production declarationSpecifiers2 is exited.
func (s *BaseCListener) ExitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) {}

// EnterDeclarationSpecifier is called when production declarationSpecifier is entered.
func (s *BaseCListener) EnterDeclarationSpecifier(ctx *DeclarationSpecifierContext) {}

// ExitDeclarationSpecifier is called when production declarationSpecifier is exited.
func (s *BaseCListener) ExitDeclarationSpecifier(ctx *DeclarationSpecifierContext) {}

// EnterInitDeclaratorList is called when production initDeclaratorList is entered.
func (s *BaseCListener) EnterInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// ExitInitDeclaratorList is called when production initDeclaratorList is exited.
func (s *BaseCListener) ExitInitDeclaratorList(ctx *InitDeclaratorListContext) {}

// EnterInitDeclarator is called when production initDeclarator is entered.
func (s *BaseCListener) EnterInitDeclarator(ctx *InitDeclaratorContext) {}

// ExitInitDeclarator is called when production initDeclarator is exited.
func (s *BaseCListener) ExitInitDeclarator(ctx *InitDeclaratorContext) {}

// EnterStorageClassSpecifier is called when production storageClassSpecifier is entered.
func (s *BaseCListener) EnterStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// ExitStorageClassSpecifier is called when production storageClassSpecifier is exited.
func (s *BaseCListener) ExitStorageClassSpecifier(ctx *StorageClassSpecifierContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseCListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseCListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterStructOrUnionSpecifier is called when production structOrUnionSpecifier is entered.
func (s *BaseCListener) EnterStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {}

// ExitStructOrUnionSpecifier is called when production structOrUnionSpecifier is exited.
func (s *BaseCListener) ExitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) {}

// EnterStructOrUnion is called when production structOrUnion is entered.
func (s *BaseCListener) EnterStructOrUnion(ctx *StructOrUnionContext) {}

// ExitStructOrUnion is called when production structOrUnion is exited.
func (s *BaseCListener) ExitStructOrUnion(ctx *StructOrUnionContext) {}

// EnterStructDeclarationList is called when production structDeclarationList is entered.
func (s *BaseCListener) EnterStructDeclarationList(ctx *StructDeclarationListContext) {}

// ExitStructDeclarationList is called when production structDeclarationList is exited.
func (s *BaseCListener) ExitStructDeclarationList(ctx *StructDeclarationListContext) {}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *BaseCListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *BaseCListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterSpecifierQualifierList is called when production specifierQualifierList is entered.
func (s *BaseCListener) EnterSpecifierQualifierList(ctx *SpecifierQualifierListContext) {}

// ExitSpecifierQualifierList is called when production specifierQualifierList is exited.
func (s *BaseCListener) ExitSpecifierQualifierList(ctx *SpecifierQualifierListContext) {}

// EnterStructDeclaratorList is called when production structDeclaratorList is entered.
func (s *BaseCListener) EnterStructDeclaratorList(ctx *StructDeclaratorListContext) {}

// ExitStructDeclaratorList is called when production structDeclaratorList is exited.
func (s *BaseCListener) ExitStructDeclaratorList(ctx *StructDeclaratorListContext) {}

// EnterStructDeclarator is called when production structDeclarator is entered.
func (s *BaseCListener) EnterStructDeclarator(ctx *StructDeclaratorContext) {}

// ExitStructDeclarator is called when production structDeclarator is exited.
func (s *BaseCListener) ExitStructDeclarator(ctx *StructDeclaratorContext) {}

// EnterEnumSpecifier is called when production enumSpecifier is entered.
func (s *BaseCListener) EnterEnumSpecifier(ctx *EnumSpecifierContext) {}

// ExitEnumSpecifier is called when production enumSpecifier is exited.
func (s *BaseCListener) ExitEnumSpecifier(ctx *EnumSpecifierContext) {}

// EnterEnumeratorList is called when production enumeratorList is entered.
func (s *BaseCListener) EnterEnumeratorList(ctx *EnumeratorListContext) {}

// ExitEnumeratorList is called when production enumeratorList is exited.
func (s *BaseCListener) ExitEnumeratorList(ctx *EnumeratorListContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseCListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseCListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterEnumerationConstant is called when production enumerationConstant is entered.
func (s *BaseCListener) EnterEnumerationConstant(ctx *EnumerationConstantContext) {}

// ExitEnumerationConstant is called when production enumerationConstant is exited.
func (s *BaseCListener) ExitEnumerationConstant(ctx *EnumerationConstantContext) {}

// EnterAtomicTypeSpecifier is called when production atomicTypeSpecifier is entered.
func (s *BaseCListener) EnterAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {}

// ExitAtomicTypeSpecifier is called when production atomicTypeSpecifier is exited.
func (s *BaseCListener) ExitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) {}

// EnterTypeQualifier is called when production typeQualifier is entered.
func (s *BaseCListener) EnterTypeQualifier(ctx *TypeQualifierContext) {}

// ExitTypeQualifier is called when production typeQualifier is exited.
func (s *BaseCListener) ExitTypeQualifier(ctx *TypeQualifierContext) {}

// EnterFunctionSpecifier is called when production functionSpecifier is entered.
func (s *BaseCListener) EnterFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// ExitFunctionSpecifier is called when production functionSpecifier is exited.
func (s *BaseCListener) ExitFunctionSpecifier(ctx *FunctionSpecifierContext) {}

// EnterAlignmentSpecifier is called when production alignmentSpecifier is entered.
func (s *BaseCListener) EnterAlignmentSpecifier(ctx *AlignmentSpecifierContext) {}

// ExitAlignmentSpecifier is called when production alignmentSpecifier is exited.
func (s *BaseCListener) ExitAlignmentSpecifier(ctx *AlignmentSpecifierContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseCListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseCListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDirectDeclarator is called when production directDeclarator is entered.
func (s *BaseCListener) EnterDirectDeclarator(ctx *DirectDeclaratorContext) {}

// ExitDirectDeclarator is called when production directDeclarator is exited.
func (s *BaseCListener) ExitDirectDeclarator(ctx *DirectDeclaratorContext) {}

// EnterGccDeclaratorExtension is called when production gccDeclaratorExtension is entered.
func (s *BaseCListener) EnterGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {}

// ExitGccDeclaratorExtension is called when production gccDeclaratorExtension is exited.
func (s *BaseCListener) ExitGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) {}

// EnterGccAttributeSpecifier is called when production gccAttributeSpecifier is entered.
func (s *BaseCListener) EnterGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {}

// ExitGccAttributeSpecifier is called when production gccAttributeSpecifier is exited.
func (s *BaseCListener) ExitGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) {}

// EnterGccAttributeList is called when production gccAttributeList is entered.
func (s *BaseCListener) EnterGccAttributeList(ctx *GccAttributeListContext) {}

// ExitGccAttributeList is called when production gccAttributeList is exited.
func (s *BaseCListener) ExitGccAttributeList(ctx *GccAttributeListContext) {}

// EnterGccAttribute is called when production gccAttribute is entered.
func (s *BaseCListener) EnterGccAttribute(ctx *GccAttributeContext) {}

// ExitGccAttribute is called when production gccAttribute is exited.
func (s *BaseCListener) ExitGccAttribute(ctx *GccAttributeContext) {}

// EnterNestedParenthesesBlock is called when production nestedParenthesesBlock is entered.
func (s *BaseCListener) EnterNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {}

// ExitNestedParenthesesBlock is called when production nestedParenthesesBlock is exited.
func (s *BaseCListener) ExitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseCListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BaseCListener) ExitPointer(ctx *PointerContext) {}

// EnterTypeQualifierList is called when production typeQualifierList is entered.
func (s *BaseCListener) EnterTypeQualifierList(ctx *TypeQualifierListContext) {}

// ExitTypeQualifierList is called when production typeQualifierList is exited.
func (s *BaseCListener) ExitTypeQualifierList(ctx *TypeQualifierListContext) {}

// EnterParameterTypeList is called when production parameterTypeList is entered.
func (s *BaseCListener) EnterParameterTypeList(ctx *ParameterTypeListContext) {}

// ExitParameterTypeList is called when production parameterTypeList is exited.
func (s *BaseCListener) ExitParameterTypeList(ctx *ParameterTypeListContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseCListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseCListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseCListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseCListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseCListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseCListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseCListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseCListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterAbstractDeclarator is called when production abstractDeclarator is entered.
func (s *BaseCListener) EnterAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// ExitAbstractDeclarator is called when production abstractDeclarator is exited.
func (s *BaseCListener) ExitAbstractDeclarator(ctx *AbstractDeclaratorContext) {}

// EnterDirectAbstractDeclarator is called when production directAbstractDeclarator is entered.
func (s *BaseCListener) EnterDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {}

// ExitDirectAbstractDeclarator is called when production directAbstractDeclarator is exited.
func (s *BaseCListener) ExitDirectAbstractDeclarator(ctx *DirectAbstractDeclaratorContext) {}

// EnterTypedefName is called when production typedefName is entered.
func (s *BaseCListener) EnterTypedefName(ctx *TypedefNameContext) {}

// ExitTypedefName is called when production typedefName is exited.
func (s *BaseCListener) ExitTypedefName(ctx *TypedefNameContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseCListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseCListener) ExitInitializer(ctx *InitializerContext) {}

// EnterInitializerList is called when production initializerList is entered.
func (s *BaseCListener) EnterInitializerList(ctx *InitializerListContext) {}

// ExitInitializerList is called when production initializerList is exited.
func (s *BaseCListener) ExitInitializerList(ctx *InitializerListContext) {}

// EnterDesignation is called when production designation is entered.
func (s *BaseCListener) EnterDesignation(ctx *DesignationContext) {}

// ExitDesignation is called when production designation is exited.
func (s *BaseCListener) ExitDesignation(ctx *DesignationContext) {}

// EnterDesignatorList is called when production designatorList is entered.
func (s *BaseCListener) EnterDesignatorList(ctx *DesignatorListContext) {}

// ExitDesignatorList is called when production designatorList is exited.
func (s *BaseCListener) ExitDesignatorList(ctx *DesignatorListContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseCListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseCListener) ExitDesignator(ctx *DesignatorContext) {}

// EnterStaticAssertDeclaration is called when production staticAssertDeclaration is entered.
func (s *BaseCListener) EnterStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// ExitStaticAssertDeclaration is called when production staticAssertDeclaration is exited.
func (s *BaseCListener) ExitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledStatement is called when production labeledStatement is entered.
func (s *BaseCListener) EnterLabeledStatement(ctx *LabeledStatementContext) {}

// ExitLabeledStatement is called when production labeledStatement is exited.
func (s *BaseCListener) ExitLabeledStatement(ctx *LabeledStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseCListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseCListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterBlockItemList is called when production blockItemList is entered.
func (s *BaseCListener) EnterBlockItemList(ctx *BlockItemListContext) {}

// ExitBlockItemList is called when production blockItemList is exited.
func (s *BaseCListener) ExitBlockItemList(ctx *BlockItemListContext) {}

// EnterBlockItem is called when production blockItem is entered.
func (s *BaseCListener) EnterBlockItem(ctx *BlockItemContext) {}

// ExitBlockItem is called when production blockItem is exited.
func (s *BaseCListener) ExitBlockItem(ctx *BlockItemContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseCListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseCListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterSelectionStatement is called when production selectionStatement is entered.
func (s *BaseCListener) EnterSelectionStatement(ctx *SelectionStatementContext) {}

// ExitSelectionStatement is called when production selectionStatement is exited.
func (s *BaseCListener) ExitSelectionStatement(ctx *SelectionStatementContext) {}

// EnterIterationStatement is called when production iterationStatement is entered.
func (s *BaseCListener) EnterIterationStatement(ctx *IterationStatementContext) {}

// ExitIterationStatement is called when production iterationStatement is exited.
func (s *BaseCListener) ExitIterationStatement(ctx *IterationStatementContext) {}

// EnterForCondition is called when production forCondition is entered.
func (s *BaseCListener) EnterForCondition(ctx *ForConditionContext) {}

// ExitForCondition is called when production forCondition is exited.
func (s *BaseCListener) ExitForCondition(ctx *ForConditionContext) {}

// EnterForDeclaration is called when production forDeclaration is entered.
func (s *BaseCListener) EnterForDeclaration(ctx *ForDeclarationContext) {}

// ExitForDeclaration is called when production forDeclaration is exited.
func (s *BaseCListener) ExitForDeclaration(ctx *ForDeclarationContext) {}

// EnterForExpression is called when production forExpression is entered.
func (s *BaseCListener) EnterForExpression(ctx *ForExpressionContext) {}

// ExitForExpression is called when production forExpression is exited.
func (s *BaseCListener) ExitForExpression(ctx *ForExpressionContext) {}

// EnterJumpStatement is called when production jumpStatement is entered.
func (s *BaseCListener) EnterJumpStatement(ctx *JumpStatementContext) {}

// ExitJumpStatement is called when production jumpStatement is exited.
func (s *BaseCListener) ExitJumpStatement(ctx *JumpStatementContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseCListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseCListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterTranslationUnit is called when production translationUnit is entered.
func (s *BaseCListener) EnterTranslationUnit(ctx *TranslationUnitContext) {}

// ExitTranslationUnit is called when production translationUnit is exited.
func (s *BaseCListener) ExitTranslationUnit(ctx *TranslationUnitContext) {}

// EnterExternalDeclaration is called when production externalDeclaration is entered.
func (s *BaseCListener) EnterExternalDeclaration(ctx *ExternalDeclarationContext) {}

// ExitExternalDeclaration is called when production externalDeclaration is exited.
func (s *BaseCListener) ExitExternalDeclaration(ctx *ExternalDeclarationContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseCListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseCListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *BaseCListener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *BaseCListener) ExitDeclarationList(ctx *DeclarationListContext) {}
