// Code generated from ATL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package atl // ATL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseATLListener is a complete listener for a parse tree produced by ATLParser.
type BaseATLListener struct{}

var _ ATLListener = &BaseATLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseATLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseATLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseATLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseATLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseATLListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseATLListener) ExitUnit(ctx *UnitContext) {}

// EnterModule is called when production module is entered.
func (s *BaseATLListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseATLListener) ExitModule(ctx *ModuleContext) {}

// EnterTargetModelPattern is called when production targetModelPattern is entered.
func (s *BaseATLListener) EnterTargetModelPattern(ctx *TargetModelPatternContext) {}

// ExitTargetModelPattern is called when production targetModelPattern is exited.
func (s *BaseATLListener) ExitTargetModelPattern(ctx *TargetModelPatternContext) {}

// EnterSourceModelPattern is called when production sourceModelPattern is entered.
func (s *BaseATLListener) EnterSourceModelPattern(ctx *SourceModelPatternContext) {}

// ExitSourceModelPattern is called when production sourceModelPattern is exited.
func (s *BaseATLListener) ExitSourceModelPattern(ctx *SourceModelPatternContext) {}

// EnterTransformationMode is called when production transformationMode is entered.
func (s *BaseATLListener) EnterTransformationMode(ctx *TransformationModeContext) {}

// ExitTransformationMode is called when production transformationMode is exited.
func (s *BaseATLListener) ExitTransformationMode(ctx *TransformationModeContext) {}

// EnterLibrary is called when production library is entered.
func (s *BaseATLListener) EnterLibrary(ctx *LibraryContext) {}

// ExitLibrary is called when production library is exited.
func (s *BaseATLListener) ExitLibrary(ctx *LibraryContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseATLListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseATLListener) ExitQuery(ctx *QueryContext) {}

// EnterLibraryRef is called when production libraryRef is entered.
func (s *BaseATLListener) EnterLibraryRef(ctx *LibraryRefContext) {}

// ExitLibraryRef is called when production libraryRef is exited.
func (s *BaseATLListener) ExitLibraryRef(ctx *LibraryRefContext) {}

// EnterModuleElement is called when production moduleElement is entered.
func (s *BaseATLListener) EnterModuleElement(ctx *ModuleElementContext) {}

// ExitModuleElement is called when production moduleElement is exited.
func (s *BaseATLListener) ExitModuleElement(ctx *ModuleElementContext) {}

// EnterHelper is called when production helper is entered.
func (s *BaseATLListener) EnterHelper(ctx *HelperContext) {}

// ExitHelper is called when production helper is exited.
func (s *BaseATLListener) ExitHelper(ctx *HelperContext) {}

// EnterOclFeatureDefinition is called when production oclFeatureDefinition is entered.
func (s *BaseATLListener) EnterOclFeatureDefinition(ctx *OclFeatureDefinitionContext) {}

// ExitOclFeatureDefinition is called when production oclFeatureDefinition is exited.
func (s *BaseATLListener) ExitOclFeatureDefinition(ctx *OclFeatureDefinitionContext) {}

// EnterOclContextDefinition is called when production oclContextDefinition is entered.
func (s *BaseATLListener) EnterOclContextDefinition(ctx *OclContextDefinitionContext) {}

// ExitOclContextDefinition is called when production oclContextDefinition is exited.
func (s *BaseATLListener) ExitOclContextDefinition(ctx *OclContextDefinitionContext) {}

// EnterOclFeature is called when production oclFeature is entered.
func (s *BaseATLListener) EnterOclFeature(ctx *OclFeatureContext) {}

// ExitOclFeature is called when production oclFeature is exited.
func (s *BaseATLListener) ExitOclFeature(ctx *OclFeatureContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseATLListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseATLListener) ExitOperation(ctx *OperationContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseATLListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseATLListener) ExitParameter(ctx *ParameterContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseATLListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseATLListener) ExitAttribute(ctx *AttributeContext) {}

// EnterArule is called when production arule is entered.
func (s *BaseATLListener) EnterArule(ctx *AruleContext) {}

// ExitArule is called when production arule is exited.
func (s *BaseATLListener) ExitArule(ctx *AruleContext) {}

// EnterMatchedRule is called when production matchedRule is entered.
func (s *BaseATLListener) EnterMatchedRule(ctx *MatchedRuleContext) {}

// ExitMatchedRule is called when production matchedRule is exited.
func (s *BaseATLListener) ExitMatchedRule(ctx *MatchedRuleContext) {}

// EnterLazyMatchedRule is called when production lazyMatchedRule is entered.
func (s *BaseATLListener) EnterLazyMatchedRule(ctx *LazyMatchedRuleContext) {}

// ExitLazyMatchedRule is called when production lazyMatchedRule is exited.
func (s *BaseATLListener) ExitLazyMatchedRule(ctx *LazyMatchedRuleContext) {}

// EnterRuleVariableDeclaration is called when production ruleVariableDeclaration is entered.
func (s *BaseATLListener) EnterRuleVariableDeclaration(ctx *RuleVariableDeclarationContext) {}

// ExitRuleVariableDeclaration is called when production ruleVariableDeclaration is exited.
func (s *BaseATLListener) ExitRuleVariableDeclaration(ctx *RuleVariableDeclarationContext) {}

// EnterCalledRule is called when production calledRule is entered.
func (s *BaseATLListener) EnterCalledRule(ctx *CalledRuleContext) {}

// ExitCalledRule is called when production calledRule is exited.
func (s *BaseATLListener) ExitCalledRule(ctx *CalledRuleContext) {}

// EnterInPattern is called when production inPattern is entered.
func (s *BaseATLListener) EnterInPattern(ctx *InPatternContext) {}

// ExitInPattern is called when production inPattern is exited.
func (s *BaseATLListener) ExitInPattern(ctx *InPatternContext) {}

// EnterInPatternElement is called when production inPatternElement is entered.
func (s *BaseATLListener) EnterInPatternElement(ctx *InPatternElementContext) {}

// ExitInPatternElement is called when production inPatternElement is exited.
func (s *BaseATLListener) ExitInPatternElement(ctx *InPatternElementContext) {}

// EnterSimpleInPatternElement is called when production simpleInPatternElement is entered.
func (s *BaseATLListener) EnterSimpleInPatternElement(ctx *SimpleInPatternElementContext) {}

// ExitSimpleInPatternElement is called when production simpleInPatternElement is exited.
func (s *BaseATLListener) ExitSimpleInPatternElement(ctx *SimpleInPatternElementContext) {}

// EnterOutPattern is called when production outPattern is entered.
func (s *BaseATLListener) EnterOutPattern(ctx *OutPatternContext) {}

// ExitOutPattern is called when production outPattern is exited.
func (s *BaseATLListener) ExitOutPattern(ctx *OutPatternContext) {}

// EnterOutPatternElement is called when production outPatternElement is entered.
func (s *BaseATLListener) EnterOutPatternElement(ctx *OutPatternElementContext) {}

// ExitOutPatternElement is called when production outPatternElement is exited.
func (s *BaseATLListener) ExitOutPatternElement(ctx *OutPatternElementContext) {}

// EnterSimpleOutPatternElement is called when production simpleOutPatternElement is entered.
func (s *BaseATLListener) EnterSimpleOutPatternElement(ctx *SimpleOutPatternElementContext) {}

// ExitSimpleOutPatternElement is called when production simpleOutPatternElement is exited.
func (s *BaseATLListener) ExitSimpleOutPatternElement(ctx *SimpleOutPatternElementContext) {}

// EnterForEachOutPatternElement is called when production forEachOutPatternElement is entered.
func (s *BaseATLListener) EnterForEachOutPatternElement(ctx *ForEachOutPatternElementContext) {}

// ExitForEachOutPatternElement is called when production forEachOutPatternElement is exited.
func (s *BaseATLListener) ExitForEachOutPatternElement(ctx *ForEachOutPatternElementContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseATLListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseATLListener) ExitBinding(ctx *BindingContext) {}

// EnterActionBlock is called when production actionBlock is entered.
func (s *BaseATLListener) EnterActionBlock(ctx *ActionBlockContext) {}

// ExitActionBlock is called when production actionBlock is exited.
func (s *BaseATLListener) ExitActionBlock(ctx *ActionBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseATLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseATLListener) ExitStatement(ctx *StatementContext) {}

// EnterBindingStat is called when production bindingStat is entered.
func (s *BaseATLListener) EnterBindingStat(ctx *BindingStatContext) {}

// ExitBindingStat is called when production bindingStat is exited.
func (s *BaseATLListener) ExitBindingStat(ctx *BindingStatContext) {}

// EnterExpressionStat is called when production expressionStat is entered.
func (s *BaseATLListener) EnterExpressionStat(ctx *ExpressionStatContext) {}

// ExitExpressionStat is called when production expressionStat is exited.
func (s *BaseATLListener) ExitExpressionStat(ctx *ExpressionStatContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseATLListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseATLListener) ExitIfStat(ctx *IfStatContext) {}

// EnterForStat is called when production forStat is entered.
func (s *BaseATLListener) EnterForStat(ctx *ForStatContext) {}

// ExitForStat is called when production forStat is exited.
func (s *BaseATLListener) ExitForStat(ctx *ForStatContext) {}

// EnterOclModel is called when production oclModel is entered.
func (s *BaseATLListener) EnterOclModel(ctx *OclModelContext) {}

// ExitOclModel is called when production oclModel is exited.
func (s *BaseATLListener) ExitOclModel(ctx *OclModelContext) {}

// EnterOclModelElement is called when production oclModelElement is entered.
func (s *BaseATLListener) EnterOclModelElement(ctx *OclModelElementContext) {}

// ExitOclModelElement is called when production oclModelElement is exited.
func (s *BaseATLListener) ExitOclModelElement(ctx *OclModelElementContext) {}

// EnterOclExpression is called when production oclExpression is entered.
func (s *BaseATLListener) EnterOclExpression(ctx *OclExpressionContext) {}

// ExitOclExpression is called when production oclExpression is exited.
func (s *BaseATLListener) ExitOclExpression(ctx *OclExpressionContext) {}

// EnterIteratorExp is called when production iteratorExp is entered.
func (s *BaseATLListener) EnterIteratorExp(ctx *IteratorExpContext) {}

// ExitIteratorExp is called when production iteratorExp is exited.
func (s *BaseATLListener) ExitIteratorExp(ctx *IteratorExpContext) {}

// EnterIterateExp is called when production iterateExp is entered.
func (s *BaseATLListener) EnterIterateExp(ctx *IterateExpContext) {}

// ExitIterateExp is called when production iterateExp is exited.
func (s *BaseATLListener) ExitIterateExp(ctx *IterateExpContext) {}

// EnterCollectionOperationCallExp is called when production collectionOperationCallExp is entered.
func (s *BaseATLListener) EnterCollectionOperationCallExp(ctx *CollectionOperationCallExpContext) {}

// ExitCollectionOperationCallExp is called when production collectionOperationCallExp is exited.
func (s *BaseATLListener) ExitCollectionOperationCallExp(ctx *CollectionOperationCallExpContext) {}

// EnterOperationCallExp is called when production operationCallExp is entered.
func (s *BaseATLListener) EnterOperationCallExp(ctx *OperationCallExpContext) {}

// ExitOperationCallExp is called when production operationCallExp is exited.
func (s *BaseATLListener) ExitOperationCallExp(ctx *OperationCallExpContext) {}

// EnterNavigationOrAttributeCallExp is called when production navigationOrAttributeCallExp is entered.
func (s *BaseATLListener) EnterNavigationOrAttributeCallExp(ctx *NavigationOrAttributeCallExpContext) {
}

// ExitNavigationOrAttributeCallExp is called when production navigationOrAttributeCallExp is exited.
func (s *BaseATLListener) ExitNavigationOrAttributeCallExp(ctx *NavigationOrAttributeCallExpContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BaseATLListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BaseATLListener) ExitIterator(ctx *IteratorContext) {}

// EnterOclUndefinedExp is called when production oclUndefinedExp is entered.
func (s *BaseATLListener) EnterOclUndefinedExp(ctx *OclUndefinedExpContext) {}

// ExitOclUndefinedExp is called when production oclUndefinedExp is exited.
func (s *BaseATLListener) ExitOclUndefinedExp(ctx *OclUndefinedExpContext) {}

// EnterPrimitiveExp is called when production primitiveExp is entered.
func (s *BaseATLListener) EnterPrimitiveExp(ctx *PrimitiveExpContext) {}

// ExitPrimitiveExp is called when production primitiveExp is exited.
func (s *BaseATLListener) ExitPrimitiveExp(ctx *PrimitiveExpContext) {}

// EnterNumericExp is called when production numericExp is entered.
func (s *BaseATLListener) EnterNumericExp(ctx *NumericExpContext) {}

// ExitNumericExp is called when production numericExp is exited.
func (s *BaseATLListener) ExitNumericExp(ctx *NumericExpContext) {}

// EnterBooleanExp is called when production booleanExp is entered.
func (s *BaseATLListener) EnterBooleanExp(ctx *BooleanExpContext) {}

// ExitBooleanExp is called when production booleanExp is exited.
func (s *BaseATLListener) ExitBooleanExp(ctx *BooleanExpContext) {}

// EnterIntegerExp is called when production integerExp is entered.
func (s *BaseATLListener) EnterIntegerExp(ctx *IntegerExpContext) {}

// ExitIntegerExp is called when production integerExp is exited.
func (s *BaseATLListener) ExitIntegerExp(ctx *IntegerExpContext) {}

// EnterRealExp is called when production realExp is entered.
func (s *BaseATLListener) EnterRealExp(ctx *RealExpContext) {}

// ExitRealExp is called when production realExp is exited.
func (s *BaseATLListener) ExitRealExp(ctx *RealExpContext) {}

// EnterStringExp is called when production stringExp is entered.
func (s *BaseATLListener) EnterStringExp(ctx *StringExpContext) {}

// ExitStringExp is called when production stringExp is exited.
func (s *BaseATLListener) ExitStringExp(ctx *StringExpContext) {}

// EnterIfExp is called when production ifExp is entered.
func (s *BaseATLListener) EnterIfExp(ctx *IfExpContext) {}

// ExitIfExp is called when production ifExp is exited.
func (s *BaseATLListener) ExitIfExp(ctx *IfExpContext) {}

// EnterVariableExp is called when production variableExp is entered.
func (s *BaseATLListener) EnterVariableExp(ctx *VariableExpContext) {}

// ExitVariableExp is called when production variableExp is exited.
func (s *BaseATLListener) ExitVariableExp(ctx *VariableExpContext) {}

// EnterSuperExp is called when production superExp is entered.
func (s *BaseATLListener) EnterSuperExp(ctx *SuperExpContext) {}

// ExitSuperExp is called when production superExp is exited.
func (s *BaseATLListener) ExitSuperExp(ctx *SuperExpContext) {}

// EnterLetExp is called when production letExp is entered.
func (s *BaseATLListener) EnterLetExp(ctx *LetExpContext) {}

// ExitLetExp is called when production letExp is exited.
func (s *BaseATLListener) ExitLetExp(ctx *LetExpContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseATLListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseATLListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterEnumLiteralExp is called when production enumLiteralExp is entered.
func (s *BaseATLListener) EnterEnumLiteralExp(ctx *EnumLiteralExpContext) {}

// ExitEnumLiteralExp is called when production enumLiteralExp is exited.
func (s *BaseATLListener) ExitEnumLiteralExp(ctx *EnumLiteralExpContext) {}

// EnterCollectionExp is called when production collectionExp is entered.
func (s *BaseATLListener) EnterCollectionExp(ctx *CollectionExpContext) {}

// ExitCollectionExp is called when production collectionExp is exited.
func (s *BaseATLListener) ExitCollectionExp(ctx *CollectionExpContext) {}

// EnterBagExp is called when production bagExp is entered.
func (s *BaseATLListener) EnterBagExp(ctx *BagExpContext) {}

// ExitBagExp is called when production bagExp is exited.
func (s *BaseATLListener) ExitBagExp(ctx *BagExpContext) {}

// EnterSetExp is called when production setExp is entered.
func (s *BaseATLListener) EnterSetExp(ctx *SetExpContext) {}

// ExitSetExp is called when production setExp is exited.
func (s *BaseATLListener) ExitSetExp(ctx *SetExpContext) {}

// EnterOrderedSetExp is called when production orderedSetExp is entered.
func (s *BaseATLListener) EnterOrderedSetExp(ctx *OrderedSetExpContext) {}

// ExitOrderedSetExp is called when production orderedSetExp is exited.
func (s *BaseATLListener) ExitOrderedSetExp(ctx *OrderedSetExpContext) {}

// EnterSequenceExp is called when production sequenceExp is entered.
func (s *BaseATLListener) EnterSequenceExp(ctx *SequenceExpContext) {}

// ExitSequenceExp is called when production sequenceExp is exited.
func (s *BaseATLListener) ExitSequenceExp(ctx *SequenceExpContext) {}

// EnterMapExp is called when production mapExp is entered.
func (s *BaseATLListener) EnterMapExp(ctx *MapExpContext) {}

// ExitMapExp is called when production mapExp is exited.
func (s *BaseATLListener) ExitMapExp(ctx *MapExpContext) {}

// EnterMapElement is called when production mapElement is entered.
func (s *BaseATLListener) EnterMapElement(ctx *MapElementContext) {}

// ExitMapElement is called when production mapElement is exited.
func (s *BaseATLListener) ExitMapElement(ctx *MapElementContext) {}

// EnterTupleExp is called when production tupleExp is entered.
func (s *BaseATLListener) EnterTupleExp(ctx *TupleExpContext) {}

// ExitTupleExp is called when production tupleExp is exited.
func (s *BaseATLListener) ExitTupleExp(ctx *TupleExpContext) {}

// EnterTuplePart is called when production tuplePart is entered.
func (s *BaseATLListener) EnterTuplePart(ctx *TuplePartContext) {}

// ExitTuplePart is called when production tuplePart is exited.
func (s *BaseATLListener) ExitTuplePart(ctx *TuplePartContext) {}

// EnterOclType is called when production oclType is entered.
func (s *BaseATLListener) EnterOclType(ctx *OclTypeContext) {}

// ExitOclType is called when production oclType is exited.
func (s *BaseATLListener) ExitOclType(ctx *OclTypeContext) {}

// EnterOclAnyType is called when production oclAnyType is entered.
func (s *BaseATLListener) EnterOclAnyType(ctx *OclAnyTypeContext) {}

// ExitOclAnyType is called when production oclAnyType is exited.
func (s *BaseATLListener) ExitOclAnyType(ctx *OclAnyTypeContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseATLListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseATLListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterTupleTypeAttribute is called when production tupleTypeAttribute is entered.
func (s *BaseATLListener) EnterTupleTypeAttribute(ctx *TupleTypeAttributeContext) {}

// ExitTupleTypeAttribute is called when production tupleTypeAttribute is exited.
func (s *BaseATLListener) ExitTupleTypeAttribute(ctx *TupleTypeAttributeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseATLListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseATLListener) ExitMapType(ctx *MapTypeContext) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseATLListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseATLListener) ExitPrimitive(ctx *PrimitiveContext) {}

// EnterNumericType is called when production numericType is entered.
func (s *BaseATLListener) EnterNumericType(ctx *NumericTypeContext) {}

// ExitNumericType is called when production numericType is exited.
func (s *BaseATLListener) ExitNumericType(ctx *NumericTypeContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseATLListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseATLListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterRealType is called when production realType is entered.
func (s *BaseATLListener) EnterRealType(ctx *RealTypeContext) {}

// ExitRealType is called when production realType is exited.
func (s *BaseATLListener) ExitRealType(ctx *RealTypeContext) {}

// EnterBooleanType is called when production booleanType is entered.
func (s *BaseATLListener) EnterBooleanType(ctx *BooleanTypeContext) {}

// ExitBooleanType is called when production booleanType is exited.
func (s *BaseATLListener) ExitBooleanType(ctx *BooleanTypeContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseATLListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseATLListener) ExitStringType(ctx *StringTypeContext) {}

// EnterCollectionType is called when production collectionType is entered.
func (s *BaseATLListener) EnterCollectionType(ctx *CollectionTypeContext) {}

// ExitCollectionType is called when production collectionType is exited.
func (s *BaseATLListener) ExitCollectionType(ctx *CollectionTypeContext) {}

// EnterBagType is called when production bagType is entered.
func (s *BaseATLListener) EnterBagType(ctx *BagTypeContext) {}

// ExitBagType is called when production bagType is exited.
func (s *BaseATLListener) ExitBagType(ctx *BagTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseATLListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseATLListener) ExitSetType(ctx *SetTypeContext) {}

// EnterOrderedSetType is called when production orderedSetType is entered.
func (s *BaseATLListener) EnterOrderedSetType(ctx *OrderedSetTypeContext) {}

// ExitOrderedSetType is called when production orderedSetType is exited.
func (s *BaseATLListener) ExitOrderedSetType(ctx *OrderedSetTypeContext) {}

// EnterSequenceType is called when production sequenceType is entered.
func (s *BaseATLListener) EnterSequenceType(ctx *SequenceTypeContext) {}

// ExitSequenceType is called when production sequenceType is exited.
func (s *BaseATLListener) ExitSequenceType(ctx *SequenceTypeContext) {}

// EnterPriority_0 is called when production priority_0 is entered.
func (s *BaseATLListener) EnterPriority_0(ctx *Priority_0Context) {}

// ExitPriority_0 is called when production priority_0 is exited.
func (s *BaseATLListener) ExitPriority_0(ctx *Priority_0Context) {}

// EnterPriority_1 is called when production priority_1 is entered.
func (s *BaseATLListener) EnterPriority_1(ctx *Priority_1Context) {}

// ExitPriority_1 is called when production priority_1 is exited.
func (s *BaseATLListener) ExitPriority_1(ctx *Priority_1Context) {}

// EnterPriority_2 is called when production priority_2 is entered.
func (s *BaseATLListener) EnterPriority_2(ctx *Priority_2Context) {}

// ExitPriority_2 is called when production priority_2 is exited.
func (s *BaseATLListener) ExitPriority_2(ctx *Priority_2Context) {}

// EnterPriority_3 is called when production priority_3 is entered.
func (s *BaseATLListener) EnterPriority_3(ctx *Priority_3Context) {}

// ExitPriority_3 is called when production priority_3 is exited.
func (s *BaseATLListener) ExitPriority_3(ctx *Priority_3Context) {}

// EnterPriority_4 is called when production priority_4 is entered.
func (s *BaseATLListener) EnterPriority_4(ctx *Priority_4Context) {}

// ExitPriority_4 is called when production priority_4 is exited.
func (s *BaseATLListener) ExitPriority_4(ctx *Priority_4Context) {}

// EnterPriority_5 is called when production priority_5 is entered.
func (s *BaseATLListener) EnterPriority_5(ctx *Priority_5Context) {}

// ExitPriority_5 is called when production priority_5 is exited.
func (s *BaseATLListener) ExitPriority_5(ctx *Priority_5Context) {}

// EnterMatchedRule_abstractContents is called when production matchedRule_abstractContents is entered.
func (s *BaseATLListener) EnterMatchedRule_abstractContents(ctx *MatchedRule_abstractContentsContext) {
}

// ExitMatchedRule_abstractContents is called when production matchedRule_abstractContents is exited.
func (s *BaseATLListener) ExitMatchedRule_abstractContents(ctx *MatchedRule_abstractContentsContext) {}

// EnterOclType_abstractContents is called when production oclType_abstractContents is entered.
func (s *BaseATLListener) EnterOclType_abstractContents(ctx *OclType_abstractContentsContext) {}

// ExitOclType_abstractContents is called when production oclType_abstractContents is exited.
func (s *BaseATLListener) ExitOclType_abstractContents(ctx *OclType_abstractContentsContext) {}

// EnterOclAnyType_abstractContents is called when production oclAnyType_abstractContents is entered.
func (s *BaseATLListener) EnterOclAnyType_abstractContents(ctx *OclAnyType_abstractContentsContext) {}

// ExitOclAnyType_abstractContents is called when production oclAnyType_abstractContents is exited.
func (s *BaseATLListener) ExitOclAnyType_abstractContents(ctx *OclAnyType_abstractContentsContext) {}

// EnterCollectionType_abstractContents is called when production collectionType_abstractContents is entered.
func (s *BaseATLListener) EnterCollectionType_abstractContents(ctx *CollectionType_abstractContentsContext) {
}

// ExitCollectionType_abstractContents is called when production collectionType_abstractContents is exited.
func (s *BaseATLListener) ExitCollectionType_abstractContents(ctx *CollectionType_abstractContentsContext) {
}

// EnterPrimary_oclExpression is called when production primary_oclExpression is entered.
func (s *BaseATLListener) EnterPrimary_oclExpression(ctx *Primary_oclExpressionContext) {}

// ExitPrimary_oclExpression is called when production primary_oclExpression is exited.
func (s *BaseATLListener) ExitPrimary_oclExpression(ctx *Primary_oclExpressionContext) {}
