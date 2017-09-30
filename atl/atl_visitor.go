// Generated from ATL.g4 by ANTLR 4.7.

package atl // ATL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ATLParser.
type ATLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ATLParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by ATLParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by ATLParser#targetModelPattern.
	VisitTargetModelPattern(ctx *TargetModelPatternContext) interface{}

	// Visit a parse tree produced by ATLParser#sourceModelPattern.
	VisitSourceModelPattern(ctx *SourceModelPatternContext) interface{}

	// Visit a parse tree produced by ATLParser#transformationMode.
	VisitTransformationMode(ctx *TransformationModeContext) interface{}

	// Visit a parse tree produced by ATLParser#library.
	VisitLibrary(ctx *LibraryContext) interface{}

	// Visit a parse tree produced by ATLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by ATLParser#libraryRef.
	VisitLibraryRef(ctx *LibraryRefContext) interface{}

	// Visit a parse tree produced by ATLParser#moduleElement.
	VisitModuleElement(ctx *ModuleElementContext) interface{}

	// Visit a parse tree produced by ATLParser#helper.
	VisitHelper(ctx *HelperContext) interface{}

	// Visit a parse tree produced by ATLParser#oclFeatureDefinition.
	VisitOclFeatureDefinition(ctx *OclFeatureDefinitionContext) interface{}

	// Visit a parse tree produced by ATLParser#oclContextDefinition.
	VisitOclContextDefinition(ctx *OclContextDefinitionContext) interface{}

	// Visit a parse tree produced by ATLParser#oclFeature.
	VisitOclFeature(ctx *OclFeatureContext) interface{}

	// Visit a parse tree produced by ATLParser#operation.
	VisitOperation(ctx *OperationContext) interface{}

	// Visit a parse tree produced by ATLParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by ATLParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by ATLParser#arule.
	VisitArule(ctx *AruleContext) interface{}

	// Visit a parse tree produced by ATLParser#matchedRule.
	VisitMatchedRule(ctx *MatchedRuleContext) interface{}

	// Visit a parse tree produced by ATLParser#lazyMatchedRule.
	VisitLazyMatchedRule(ctx *LazyMatchedRuleContext) interface{}

	// Visit a parse tree produced by ATLParser#ruleVariableDeclaration.
	VisitRuleVariableDeclaration(ctx *RuleVariableDeclarationContext) interface{}

	// Visit a parse tree produced by ATLParser#calledRule.
	VisitCalledRule(ctx *CalledRuleContext) interface{}

	// Visit a parse tree produced by ATLParser#inPattern.
	VisitInPattern(ctx *InPatternContext) interface{}

	// Visit a parse tree produced by ATLParser#inPatternElement.
	VisitInPatternElement(ctx *InPatternElementContext) interface{}

	// Visit a parse tree produced by ATLParser#simpleInPatternElement.
	VisitSimpleInPatternElement(ctx *SimpleInPatternElementContext) interface{}

	// Visit a parse tree produced by ATLParser#outPattern.
	VisitOutPattern(ctx *OutPatternContext) interface{}

	// Visit a parse tree produced by ATLParser#outPatternElement.
	VisitOutPatternElement(ctx *OutPatternElementContext) interface{}

	// Visit a parse tree produced by ATLParser#simpleOutPatternElement.
	VisitSimpleOutPatternElement(ctx *SimpleOutPatternElementContext) interface{}

	// Visit a parse tree produced by ATLParser#forEachOutPatternElement.
	VisitForEachOutPatternElement(ctx *ForEachOutPatternElementContext) interface{}

	// Visit a parse tree produced by ATLParser#binding.
	VisitBinding(ctx *BindingContext) interface{}

	// Visit a parse tree produced by ATLParser#actionBlock.
	VisitActionBlock(ctx *ActionBlockContext) interface{}

	// Visit a parse tree produced by ATLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ATLParser#bindingStat.
	VisitBindingStat(ctx *BindingStatContext) interface{}

	// Visit a parse tree produced by ATLParser#expressionStat.
	VisitExpressionStat(ctx *ExpressionStatContext) interface{}

	// Visit a parse tree produced by ATLParser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by ATLParser#forStat.
	VisitForStat(ctx *ForStatContext) interface{}

	// Visit a parse tree produced by ATLParser#oclModel.
	VisitOclModel(ctx *OclModelContext) interface{}

	// Visit a parse tree produced by ATLParser#oclModelElement.
	VisitOclModelElement(ctx *OclModelElementContext) interface{}

	// Visit a parse tree produced by ATLParser#oclExpression.
	VisitOclExpression(ctx *OclExpressionContext) interface{}

	// Visit a parse tree produced by ATLParser#iteratorExp.
	VisitIteratorExp(ctx *IteratorExpContext) interface{}

	// Visit a parse tree produced by ATLParser#iterateExp.
	VisitIterateExp(ctx *IterateExpContext) interface{}

	// Visit a parse tree produced by ATLParser#collectionOperationCallExp.
	VisitCollectionOperationCallExp(ctx *CollectionOperationCallExpContext) interface{}

	// Visit a parse tree produced by ATLParser#operationCallExp.
	VisitOperationCallExp(ctx *OperationCallExpContext) interface{}

	// Visit a parse tree produced by ATLParser#navigationOrAttributeCallExp.
	VisitNavigationOrAttributeCallExp(ctx *NavigationOrAttributeCallExpContext) interface{}

	// Visit a parse tree produced by ATLParser#iterator.
	VisitIterator(ctx *IteratorContext) interface{}

	// Visit a parse tree produced by ATLParser#oclUndefinedExp.
	VisitOclUndefinedExp(ctx *OclUndefinedExpContext) interface{}

	// Visit a parse tree produced by ATLParser#primitiveExp.
	VisitPrimitiveExp(ctx *PrimitiveExpContext) interface{}

	// Visit a parse tree produced by ATLParser#numericExp.
	VisitNumericExp(ctx *NumericExpContext) interface{}

	// Visit a parse tree produced by ATLParser#booleanExp.
	VisitBooleanExp(ctx *BooleanExpContext) interface{}

	// Visit a parse tree produced by ATLParser#integerExp.
	VisitIntegerExp(ctx *IntegerExpContext) interface{}

	// Visit a parse tree produced by ATLParser#realExp.
	VisitRealExp(ctx *RealExpContext) interface{}

	// Visit a parse tree produced by ATLParser#stringExp.
	VisitStringExp(ctx *StringExpContext) interface{}

	// Visit a parse tree produced by ATLParser#ifExp.
	VisitIfExp(ctx *IfExpContext) interface{}

	// Visit a parse tree produced by ATLParser#variableExp.
	VisitVariableExp(ctx *VariableExpContext) interface{}

	// Visit a parse tree produced by ATLParser#superExp.
	VisitSuperExp(ctx *SuperExpContext) interface{}

	// Visit a parse tree produced by ATLParser#letExp.
	VisitLetExp(ctx *LetExpContext) interface{}

	// Visit a parse tree produced by ATLParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ATLParser#enumLiteralExp.
	VisitEnumLiteralExp(ctx *EnumLiteralExpContext) interface{}

	// Visit a parse tree produced by ATLParser#collectionExp.
	VisitCollectionExp(ctx *CollectionExpContext) interface{}

	// Visit a parse tree produced by ATLParser#bagExp.
	VisitBagExp(ctx *BagExpContext) interface{}

	// Visit a parse tree produced by ATLParser#setExp.
	VisitSetExp(ctx *SetExpContext) interface{}

	// Visit a parse tree produced by ATLParser#orderedSetExp.
	VisitOrderedSetExp(ctx *OrderedSetExpContext) interface{}

	// Visit a parse tree produced by ATLParser#sequenceExp.
	VisitSequenceExp(ctx *SequenceExpContext) interface{}

	// Visit a parse tree produced by ATLParser#mapExp.
	VisitMapExp(ctx *MapExpContext) interface{}

	// Visit a parse tree produced by ATLParser#mapElement.
	VisitMapElement(ctx *MapElementContext) interface{}

	// Visit a parse tree produced by ATLParser#tupleExp.
	VisitTupleExp(ctx *TupleExpContext) interface{}

	// Visit a parse tree produced by ATLParser#tuplePart.
	VisitTuplePart(ctx *TuplePartContext) interface{}

	// Visit a parse tree produced by ATLParser#oclType.
	VisitOclType(ctx *OclTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#oclAnyType.
	VisitOclAnyType(ctx *OclAnyTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#tupleTypeAttribute.
	VisitTupleTypeAttribute(ctx *TupleTypeAttributeContext) interface{}

	// Visit a parse tree produced by ATLParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#primitive.
	VisitPrimitive(ctx *PrimitiveContext) interface{}

	// Visit a parse tree produced by ATLParser#numericType.
	VisitNumericType(ctx *NumericTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#integerType.
	VisitIntegerType(ctx *IntegerTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#realType.
	VisitRealType(ctx *RealTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#booleanType.
	VisitBooleanType(ctx *BooleanTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#collectionType.
	VisitCollectionType(ctx *CollectionTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#bagType.
	VisitBagType(ctx *BagTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#setType.
	VisitSetType(ctx *SetTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#orderedSetType.
	VisitOrderedSetType(ctx *OrderedSetTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#sequenceType.
	VisitSequenceType(ctx *SequenceTypeContext) interface{}

	// Visit a parse tree produced by ATLParser#priority_0.
	VisitPriority_0(ctx *Priority_0Context) interface{}

	// Visit a parse tree produced by ATLParser#priority_1.
	VisitPriority_1(ctx *Priority_1Context) interface{}

	// Visit a parse tree produced by ATLParser#priority_2.
	VisitPriority_2(ctx *Priority_2Context) interface{}

	// Visit a parse tree produced by ATLParser#priority_3.
	VisitPriority_3(ctx *Priority_3Context) interface{}

	// Visit a parse tree produced by ATLParser#priority_4.
	VisitPriority_4(ctx *Priority_4Context) interface{}

	// Visit a parse tree produced by ATLParser#priority_5.
	VisitPriority_5(ctx *Priority_5Context) interface{}

	// Visit a parse tree produced by ATLParser#matchedRule_abstractContents.
	VisitMatchedRule_abstractContents(ctx *MatchedRule_abstractContentsContext) interface{}

	// Visit a parse tree produced by ATLParser#oclType_abstractContents.
	VisitOclType_abstractContents(ctx *OclType_abstractContentsContext) interface{}

	// Visit a parse tree produced by ATLParser#oclAnyType_abstractContents.
	VisitOclAnyType_abstractContents(ctx *OclAnyType_abstractContentsContext) interface{}

	// Visit a parse tree produced by ATLParser#collectionType_abstractContents.
	VisitCollectionType_abstractContents(ctx *CollectionType_abstractContentsContext) interface{}

	// Visit a parse tree produced by ATLParser#primary_oclExpression.
	VisitPrimary_oclExpression(ctx *Primary_oclExpressionContext) interface{}
}
