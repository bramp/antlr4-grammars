// Generated from ATL.g4 by ANTLR 4.7.

package atl // ATL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseATLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseATLVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTargetModelPattern(ctx *TargetModelPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSourceModelPattern(ctx *SourceModelPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTransformationMode(ctx *TransformationModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitLibrary(ctx *LibraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitLibraryRef(ctx *LibraryRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitModuleElement(ctx *ModuleElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitHelper(ctx *HelperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclFeatureDefinition(ctx *OclFeatureDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclContextDefinition(ctx *OclContextDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclFeature(ctx *OclFeatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOperation(ctx *OperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitArule(ctx *AruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitMatchedRule(ctx *MatchedRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitLazyMatchedRule(ctx *LazyMatchedRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitRuleVariableDeclaration(ctx *RuleVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitCalledRule(ctx *CalledRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitInPattern(ctx *InPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitInPatternElement(ctx *InPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSimpleInPatternElement(ctx *SimpleInPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOutPattern(ctx *OutPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOutPatternElement(ctx *OutPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSimpleOutPatternElement(ctx *SimpleOutPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitForEachOutPatternElement(ctx *ForEachOutPatternElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBinding(ctx *BindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitActionBlock(ctx *ActionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBindingStat(ctx *BindingStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitExpressionStat(ctx *ExpressionStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitForStat(ctx *ForStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclModel(ctx *OclModelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclModelElement(ctx *OclModelElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclExpression(ctx *OclExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIteratorExp(ctx *IteratorExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIterateExp(ctx *IterateExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitCollectionOperationCallExp(ctx *CollectionOperationCallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOperationCallExp(ctx *OperationCallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitNavigationOrAttributeCallExp(ctx *NavigationOrAttributeCallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIterator(ctx *IteratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclUndefinedExp(ctx *OclUndefinedExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPrimitiveExp(ctx *PrimitiveExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitNumericExp(ctx *NumericExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBooleanExp(ctx *BooleanExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIntegerExp(ctx *IntegerExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitRealExp(ctx *RealExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitStringExp(ctx *StringExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIfExp(ctx *IfExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitVariableExp(ctx *VariableExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSuperExp(ctx *SuperExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitLetExp(ctx *LetExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitEnumLiteralExp(ctx *EnumLiteralExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitCollectionExp(ctx *CollectionExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBagExp(ctx *BagExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSetExp(ctx *SetExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOrderedSetExp(ctx *OrderedSetExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSequenceExp(ctx *SequenceExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitMapExp(ctx *MapExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitMapElement(ctx *MapElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTupleExp(ctx *TupleExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTuplePart(ctx *TuplePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclType(ctx *OclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclAnyType(ctx *OclAnyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitTupleTypeAttribute(ctx *TupleTypeAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPrimitive(ctx *PrimitiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitNumericType(ctx *NumericTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitIntegerType(ctx *IntegerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitRealType(ctx *RealTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBooleanType(ctx *BooleanTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitCollectionType(ctx *CollectionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitBagType(ctx *BagTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSetType(ctx *SetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOrderedSetType(ctx *OrderedSetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitSequenceType(ctx *SequenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_0(ctx *Priority_0Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_1(ctx *Priority_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_2(ctx *Priority_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_3(ctx *Priority_3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_4(ctx *Priority_4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPriority_5(ctx *Priority_5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitMatchedRule_abstractContents(ctx *MatchedRule_abstractContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclType_abstractContents(ctx *OclType_abstractContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitOclAnyType_abstractContents(ctx *OclAnyType_abstractContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitCollectionType_abstractContents(ctx *CollectionType_abstractContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseATLVisitor) VisitPrimary_oclExpression(ctx *Primary_oclExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
