// Generated from ATL.g4 by ANTLR 4.7.

package atl // ATL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ATLListener is a complete listener for a parse tree produced by ATLParser.
type ATLListener interface {
	antlr.ParseTreeListener

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterTargetModelPattern is called when entering the targetModelPattern production.
	EnterTargetModelPattern(c *TargetModelPatternContext)

	// EnterSourceModelPattern is called when entering the sourceModelPattern production.
	EnterSourceModelPattern(c *SourceModelPatternContext)

	// EnterTransformationMode is called when entering the transformationMode production.
	EnterTransformationMode(c *TransformationModeContext)

	// EnterLibrary is called when entering the library production.
	EnterLibrary(c *LibraryContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterLibraryRef is called when entering the libraryRef production.
	EnterLibraryRef(c *LibraryRefContext)

	// EnterModuleElement is called when entering the moduleElement production.
	EnterModuleElement(c *ModuleElementContext)

	// EnterHelper is called when entering the helper production.
	EnterHelper(c *HelperContext)

	// EnterOclFeatureDefinition is called when entering the oclFeatureDefinition production.
	EnterOclFeatureDefinition(c *OclFeatureDefinitionContext)

	// EnterOclContextDefinition is called when entering the oclContextDefinition production.
	EnterOclContextDefinition(c *OclContextDefinitionContext)

	// EnterOclFeature is called when entering the oclFeature production.
	EnterOclFeature(c *OclFeatureContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterArule is called when entering the arule production.
	EnterArule(c *AruleContext)

	// EnterMatchedRule is called when entering the matchedRule production.
	EnterMatchedRule(c *MatchedRuleContext)

	// EnterLazyMatchedRule is called when entering the lazyMatchedRule production.
	EnterLazyMatchedRule(c *LazyMatchedRuleContext)

	// EnterRuleVariableDeclaration is called when entering the ruleVariableDeclaration production.
	EnterRuleVariableDeclaration(c *RuleVariableDeclarationContext)

	// EnterCalledRule is called when entering the calledRule production.
	EnterCalledRule(c *CalledRuleContext)

	// EnterInPattern is called when entering the inPattern production.
	EnterInPattern(c *InPatternContext)

	// EnterInPatternElement is called when entering the inPatternElement production.
	EnterInPatternElement(c *InPatternElementContext)

	// EnterSimpleInPatternElement is called when entering the simpleInPatternElement production.
	EnterSimpleInPatternElement(c *SimpleInPatternElementContext)

	// EnterOutPattern is called when entering the outPattern production.
	EnterOutPattern(c *OutPatternContext)

	// EnterOutPatternElement is called when entering the outPatternElement production.
	EnterOutPatternElement(c *OutPatternElementContext)

	// EnterSimpleOutPatternElement is called when entering the simpleOutPatternElement production.
	EnterSimpleOutPatternElement(c *SimpleOutPatternElementContext)

	// EnterForEachOutPatternElement is called when entering the forEachOutPatternElement production.
	EnterForEachOutPatternElement(c *ForEachOutPatternElementContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterActionBlock is called when entering the actionBlock production.
	EnterActionBlock(c *ActionBlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBindingStat is called when entering the bindingStat production.
	EnterBindingStat(c *BindingStatContext)

	// EnterExpressionStat is called when entering the expressionStat production.
	EnterExpressionStat(c *ExpressionStatContext)

	// EnterIfStat is called when entering the ifStat production.
	EnterIfStat(c *IfStatContext)

	// EnterForStat is called when entering the forStat production.
	EnterForStat(c *ForStatContext)

	// EnterOclModel is called when entering the oclModel production.
	EnterOclModel(c *OclModelContext)

	// EnterOclModelElement is called when entering the oclModelElement production.
	EnterOclModelElement(c *OclModelElementContext)

	// EnterOclExpression is called when entering the oclExpression production.
	EnterOclExpression(c *OclExpressionContext)

	// EnterIteratorExp is called when entering the iteratorExp production.
	EnterIteratorExp(c *IteratorExpContext)

	// EnterIterateExp is called when entering the iterateExp production.
	EnterIterateExp(c *IterateExpContext)

	// EnterCollectionOperationCallExp is called when entering the collectionOperationCallExp production.
	EnterCollectionOperationCallExp(c *CollectionOperationCallExpContext)

	// EnterOperationCallExp is called when entering the operationCallExp production.
	EnterOperationCallExp(c *OperationCallExpContext)

	// EnterNavigationOrAttributeCallExp is called when entering the navigationOrAttributeCallExp production.
	EnterNavigationOrAttributeCallExp(c *NavigationOrAttributeCallExpContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterOclUndefinedExp is called when entering the oclUndefinedExp production.
	EnterOclUndefinedExp(c *OclUndefinedExpContext)

	// EnterPrimitiveExp is called when entering the primitiveExp production.
	EnterPrimitiveExp(c *PrimitiveExpContext)

	// EnterNumericExp is called when entering the numericExp production.
	EnterNumericExp(c *NumericExpContext)

	// EnterBooleanExp is called when entering the booleanExp production.
	EnterBooleanExp(c *BooleanExpContext)

	// EnterIntegerExp is called when entering the integerExp production.
	EnterIntegerExp(c *IntegerExpContext)

	// EnterRealExp is called when entering the realExp production.
	EnterRealExp(c *RealExpContext)

	// EnterStringExp is called when entering the stringExp production.
	EnterStringExp(c *StringExpContext)

	// EnterIfExp is called when entering the ifExp production.
	EnterIfExp(c *IfExpContext)

	// EnterVariableExp is called when entering the variableExp production.
	EnterVariableExp(c *VariableExpContext)

	// EnterSuperExp is called when entering the superExp production.
	EnterSuperExp(c *SuperExpContext)

	// EnterLetExp is called when entering the letExp production.
	EnterLetExp(c *LetExpContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterEnumLiteralExp is called when entering the enumLiteralExp production.
	EnterEnumLiteralExp(c *EnumLiteralExpContext)

	// EnterCollectionExp is called when entering the collectionExp production.
	EnterCollectionExp(c *CollectionExpContext)

	// EnterBagExp is called when entering the bagExp production.
	EnterBagExp(c *BagExpContext)

	// EnterSetExp is called when entering the setExp production.
	EnterSetExp(c *SetExpContext)

	// EnterOrderedSetExp is called when entering the orderedSetExp production.
	EnterOrderedSetExp(c *OrderedSetExpContext)

	// EnterSequenceExp is called when entering the sequenceExp production.
	EnterSequenceExp(c *SequenceExpContext)

	// EnterMapExp is called when entering the mapExp production.
	EnterMapExp(c *MapExpContext)

	// EnterMapElement is called when entering the mapElement production.
	EnterMapElement(c *MapElementContext)

	// EnterTupleExp is called when entering the tupleExp production.
	EnterTupleExp(c *TupleExpContext)

	// EnterTuplePart is called when entering the tuplePart production.
	EnterTuplePart(c *TuplePartContext)

	// EnterOclType is called when entering the oclType production.
	EnterOclType(c *OclTypeContext)

	// EnterOclAnyType is called when entering the oclAnyType production.
	EnterOclAnyType(c *OclAnyTypeContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterTupleTypeAttribute is called when entering the tupleTypeAttribute production.
	EnterTupleTypeAttribute(c *TupleTypeAttributeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// EnterNumericType is called when entering the numericType production.
	EnterNumericType(c *NumericTypeContext)

	// EnterIntegerType is called when entering the integerType production.
	EnterIntegerType(c *IntegerTypeContext)

	// EnterRealType is called when entering the realType production.
	EnterRealType(c *RealTypeContext)

	// EnterBooleanType is called when entering the booleanType production.
	EnterBooleanType(c *BooleanTypeContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterCollectionType is called when entering the collectionType production.
	EnterCollectionType(c *CollectionTypeContext)

	// EnterBagType is called when entering the bagType production.
	EnterBagType(c *BagTypeContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterOrderedSetType is called when entering the orderedSetType production.
	EnterOrderedSetType(c *OrderedSetTypeContext)

	// EnterSequenceType is called when entering the sequenceType production.
	EnterSequenceType(c *SequenceTypeContext)

	// EnterPriority_0 is called when entering the priority_0 production.
	EnterPriority_0(c *Priority_0Context)

	// EnterPriority_1 is called when entering the priority_1 production.
	EnterPriority_1(c *Priority_1Context)

	// EnterPriority_2 is called when entering the priority_2 production.
	EnterPriority_2(c *Priority_2Context)

	// EnterPriority_3 is called when entering the priority_3 production.
	EnterPriority_3(c *Priority_3Context)

	// EnterPriority_4 is called when entering the priority_4 production.
	EnterPriority_4(c *Priority_4Context)

	// EnterPriority_5 is called when entering the priority_5 production.
	EnterPriority_5(c *Priority_5Context)

	// EnterMatchedRule_abstractContents is called when entering the matchedRule_abstractContents production.
	EnterMatchedRule_abstractContents(c *MatchedRule_abstractContentsContext)

	// EnterOclType_abstractContents is called when entering the oclType_abstractContents production.
	EnterOclType_abstractContents(c *OclType_abstractContentsContext)

	// EnterOclAnyType_abstractContents is called when entering the oclAnyType_abstractContents production.
	EnterOclAnyType_abstractContents(c *OclAnyType_abstractContentsContext)

	// EnterCollectionType_abstractContents is called when entering the collectionType_abstractContents production.
	EnterCollectionType_abstractContents(c *CollectionType_abstractContentsContext)

	// EnterPrimary_oclExpression is called when entering the primary_oclExpression production.
	EnterPrimary_oclExpression(c *Primary_oclExpressionContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitTargetModelPattern is called when exiting the targetModelPattern production.
	ExitTargetModelPattern(c *TargetModelPatternContext)

	// ExitSourceModelPattern is called when exiting the sourceModelPattern production.
	ExitSourceModelPattern(c *SourceModelPatternContext)

	// ExitTransformationMode is called when exiting the transformationMode production.
	ExitTransformationMode(c *TransformationModeContext)

	// ExitLibrary is called when exiting the library production.
	ExitLibrary(c *LibraryContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitLibraryRef is called when exiting the libraryRef production.
	ExitLibraryRef(c *LibraryRefContext)

	// ExitModuleElement is called when exiting the moduleElement production.
	ExitModuleElement(c *ModuleElementContext)

	// ExitHelper is called when exiting the helper production.
	ExitHelper(c *HelperContext)

	// ExitOclFeatureDefinition is called when exiting the oclFeatureDefinition production.
	ExitOclFeatureDefinition(c *OclFeatureDefinitionContext)

	// ExitOclContextDefinition is called when exiting the oclContextDefinition production.
	ExitOclContextDefinition(c *OclContextDefinitionContext)

	// ExitOclFeature is called when exiting the oclFeature production.
	ExitOclFeature(c *OclFeatureContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitArule is called when exiting the arule production.
	ExitArule(c *AruleContext)

	// ExitMatchedRule is called when exiting the matchedRule production.
	ExitMatchedRule(c *MatchedRuleContext)

	// ExitLazyMatchedRule is called when exiting the lazyMatchedRule production.
	ExitLazyMatchedRule(c *LazyMatchedRuleContext)

	// ExitRuleVariableDeclaration is called when exiting the ruleVariableDeclaration production.
	ExitRuleVariableDeclaration(c *RuleVariableDeclarationContext)

	// ExitCalledRule is called when exiting the calledRule production.
	ExitCalledRule(c *CalledRuleContext)

	// ExitInPattern is called when exiting the inPattern production.
	ExitInPattern(c *InPatternContext)

	// ExitInPatternElement is called when exiting the inPatternElement production.
	ExitInPatternElement(c *InPatternElementContext)

	// ExitSimpleInPatternElement is called when exiting the simpleInPatternElement production.
	ExitSimpleInPatternElement(c *SimpleInPatternElementContext)

	// ExitOutPattern is called when exiting the outPattern production.
	ExitOutPattern(c *OutPatternContext)

	// ExitOutPatternElement is called when exiting the outPatternElement production.
	ExitOutPatternElement(c *OutPatternElementContext)

	// ExitSimpleOutPatternElement is called when exiting the simpleOutPatternElement production.
	ExitSimpleOutPatternElement(c *SimpleOutPatternElementContext)

	// ExitForEachOutPatternElement is called when exiting the forEachOutPatternElement production.
	ExitForEachOutPatternElement(c *ForEachOutPatternElementContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitActionBlock is called when exiting the actionBlock production.
	ExitActionBlock(c *ActionBlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBindingStat is called when exiting the bindingStat production.
	ExitBindingStat(c *BindingStatContext)

	// ExitExpressionStat is called when exiting the expressionStat production.
	ExitExpressionStat(c *ExpressionStatContext)

	// ExitIfStat is called when exiting the ifStat production.
	ExitIfStat(c *IfStatContext)

	// ExitForStat is called when exiting the forStat production.
	ExitForStat(c *ForStatContext)

	// ExitOclModel is called when exiting the oclModel production.
	ExitOclModel(c *OclModelContext)

	// ExitOclModelElement is called when exiting the oclModelElement production.
	ExitOclModelElement(c *OclModelElementContext)

	// ExitOclExpression is called when exiting the oclExpression production.
	ExitOclExpression(c *OclExpressionContext)

	// ExitIteratorExp is called when exiting the iteratorExp production.
	ExitIteratorExp(c *IteratorExpContext)

	// ExitIterateExp is called when exiting the iterateExp production.
	ExitIterateExp(c *IterateExpContext)

	// ExitCollectionOperationCallExp is called when exiting the collectionOperationCallExp production.
	ExitCollectionOperationCallExp(c *CollectionOperationCallExpContext)

	// ExitOperationCallExp is called when exiting the operationCallExp production.
	ExitOperationCallExp(c *OperationCallExpContext)

	// ExitNavigationOrAttributeCallExp is called when exiting the navigationOrAttributeCallExp production.
	ExitNavigationOrAttributeCallExp(c *NavigationOrAttributeCallExpContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitOclUndefinedExp is called when exiting the oclUndefinedExp production.
	ExitOclUndefinedExp(c *OclUndefinedExpContext)

	// ExitPrimitiveExp is called when exiting the primitiveExp production.
	ExitPrimitiveExp(c *PrimitiveExpContext)

	// ExitNumericExp is called when exiting the numericExp production.
	ExitNumericExp(c *NumericExpContext)

	// ExitBooleanExp is called when exiting the booleanExp production.
	ExitBooleanExp(c *BooleanExpContext)

	// ExitIntegerExp is called when exiting the integerExp production.
	ExitIntegerExp(c *IntegerExpContext)

	// ExitRealExp is called when exiting the realExp production.
	ExitRealExp(c *RealExpContext)

	// ExitStringExp is called when exiting the stringExp production.
	ExitStringExp(c *StringExpContext)

	// ExitIfExp is called when exiting the ifExp production.
	ExitIfExp(c *IfExpContext)

	// ExitVariableExp is called when exiting the variableExp production.
	ExitVariableExp(c *VariableExpContext)

	// ExitSuperExp is called when exiting the superExp production.
	ExitSuperExp(c *SuperExpContext)

	// ExitLetExp is called when exiting the letExp production.
	ExitLetExp(c *LetExpContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitEnumLiteralExp is called when exiting the enumLiteralExp production.
	ExitEnumLiteralExp(c *EnumLiteralExpContext)

	// ExitCollectionExp is called when exiting the collectionExp production.
	ExitCollectionExp(c *CollectionExpContext)

	// ExitBagExp is called when exiting the bagExp production.
	ExitBagExp(c *BagExpContext)

	// ExitSetExp is called when exiting the setExp production.
	ExitSetExp(c *SetExpContext)

	// ExitOrderedSetExp is called when exiting the orderedSetExp production.
	ExitOrderedSetExp(c *OrderedSetExpContext)

	// ExitSequenceExp is called when exiting the sequenceExp production.
	ExitSequenceExp(c *SequenceExpContext)

	// ExitMapExp is called when exiting the mapExp production.
	ExitMapExp(c *MapExpContext)

	// ExitMapElement is called when exiting the mapElement production.
	ExitMapElement(c *MapElementContext)

	// ExitTupleExp is called when exiting the tupleExp production.
	ExitTupleExp(c *TupleExpContext)

	// ExitTuplePart is called when exiting the tuplePart production.
	ExitTuplePart(c *TuplePartContext)

	// ExitOclType is called when exiting the oclType production.
	ExitOclType(c *OclTypeContext)

	// ExitOclAnyType is called when exiting the oclAnyType production.
	ExitOclAnyType(c *OclAnyTypeContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitTupleTypeAttribute is called when exiting the tupleTypeAttribute production.
	ExitTupleTypeAttribute(c *TupleTypeAttributeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)

	// ExitNumericType is called when exiting the numericType production.
	ExitNumericType(c *NumericTypeContext)

	// ExitIntegerType is called when exiting the integerType production.
	ExitIntegerType(c *IntegerTypeContext)

	// ExitRealType is called when exiting the realType production.
	ExitRealType(c *RealTypeContext)

	// ExitBooleanType is called when exiting the booleanType production.
	ExitBooleanType(c *BooleanTypeContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitCollectionType is called when exiting the collectionType production.
	ExitCollectionType(c *CollectionTypeContext)

	// ExitBagType is called when exiting the bagType production.
	ExitBagType(c *BagTypeContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitOrderedSetType is called when exiting the orderedSetType production.
	ExitOrderedSetType(c *OrderedSetTypeContext)

	// ExitSequenceType is called when exiting the sequenceType production.
	ExitSequenceType(c *SequenceTypeContext)

	// ExitPriority_0 is called when exiting the priority_0 production.
	ExitPriority_0(c *Priority_0Context)

	// ExitPriority_1 is called when exiting the priority_1 production.
	ExitPriority_1(c *Priority_1Context)

	// ExitPriority_2 is called when exiting the priority_2 production.
	ExitPriority_2(c *Priority_2Context)

	// ExitPriority_3 is called when exiting the priority_3 production.
	ExitPriority_3(c *Priority_3Context)

	// ExitPriority_4 is called when exiting the priority_4 production.
	ExitPriority_4(c *Priority_4Context)

	// ExitPriority_5 is called when exiting the priority_5 production.
	ExitPriority_5(c *Priority_5Context)

	// ExitMatchedRule_abstractContents is called when exiting the matchedRule_abstractContents production.
	ExitMatchedRule_abstractContents(c *MatchedRule_abstractContentsContext)

	// ExitOclType_abstractContents is called when exiting the oclType_abstractContents production.
	ExitOclType_abstractContents(c *OclType_abstractContentsContext)

	// ExitOclAnyType_abstractContents is called when exiting the oclAnyType_abstractContents production.
	ExitOclAnyType_abstractContents(c *OclAnyType_abstractContentsContext)

	// ExitCollectionType_abstractContents is called when exiting the collectionType_abstractContents production.
	ExitCollectionType_abstractContents(c *CollectionType_abstractContentsContext)

	// ExitPrimary_oclExpression is called when exiting the primary_oclExpression production.
	ExitPrimary_oclExpression(c *Primary_oclExpressionContext)
}
