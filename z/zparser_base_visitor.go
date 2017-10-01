// Generated from ZParser.g4 by ANTLR 4.7.

package z // ZParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseZParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZParserVisitor) VisitSpecification(ctx *SpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInheritingSection(ctx *InheritingSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitBaseSection(ctx *BaseSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGivenTypesParagraph(ctx *GivenTypesParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericAxiomaticDescriptionParagraph(ctx *GenericAxiomaticDescriptionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericSchemaDefinitionParagraph(ctx *GenericSchemaDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitHorizontalDefinitionParagraph(ctx *HorizontalDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericHorizontalDefinitionParagraph(ctx *GenericHorizontalDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericOperatorDefinitionParagraph(ctx *GenericOperatorDefinitionParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitFreeTypesParagraph(ctx *FreeTypesParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitConjectureParagraph(ctx *ConjectureParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericConjectureParagraph(ctx *GenericConjectureParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitFreetype(ctx *FreetypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitBranch(ctx *BranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitFormals(ctx *FormalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitExistentialQuantificationPredicate(ctx *ExistentialQuantificationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitConjunctionPredicate(ctx *ConjunctionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitEquivalencePredicate(ctx *EquivalencePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNewlineConjunctionPredicate(ctx *NewlineConjunctionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitImplicationPredicate(ctx *ImplicationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitRelationOperatorApplicationPredicate(ctx *RelationOperatorApplicationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitUniversalQuantificationPredicate(ctx *UniversalQuantificationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitTruthPredicate(ctx *TruthPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitFalsityPredicate(ctx *FalsityPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitParenthesizedPredicate(ctx *ParenthesizedPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDisjunctionPredicate(ctx *DisjunctionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSemicolonConjunctionPredicate(ctx *SemicolonConjunctionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaPredicatePredicate(ctx *SchemaPredicatePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitUniqueExistentialQuantificationPredicate(ctx *UniqueExistentialQuantificationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNegationPredicate(ctx *NegationPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSetComprehensionExpression(ctx *SetComprehensionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaEquivalenceExpression(ctx *SchemaEquivalenceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixApplicationExpression(ctx *NofixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaConstructionExpression(ctx *SchemaConstructionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericPostfixApplicationExpression(ctx *GenericPostfixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericPrefixApplicationExpression(ctx *GenericPrefixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitCharacteristicDefiniteDescriptionExpression(ctx *CharacteristicDefiniteDescriptionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaUniversalQuantificationExpression(ctx *SchemaUniversalQuantificationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSetExtensionExpression(ctx *SetExtensionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitFunctionConstructionExpression(ctx *FunctionConstructionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaNegationExpression(ctx *SchemaNegationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitTupleExtensionExpression(ctx *TupleExtensionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaCompositionExpression(ctx *SchemaCompositionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPowersetExpression(ctx *PowersetExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaDisjunctionExpression(ctx *SchemaDisjunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSubstitutionExpressionExpression(ctx *SubstitutionExpressionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaHidingExpression(ctx *SchemaHidingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixApplicationExpression(ctx *PostfixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitCartesianProductExpression(ctx *CartesianProductExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitBindingSelectionExpression(ctx *BindingSelectionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericInstantiationExpression(ctx *GenericInstantiationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaExistentialQuantificationExpression(ctx *SchemaExistentialQuantificationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNumberLiteralExpression(ctx *NumberLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixLeftApplicationExpression(ctx *InfixLeftApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitApplicationExpression(ctx *ApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaRenamingExpression(ctx *SchemaRenamingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaConjunctionExpression(ctx *SchemaConjunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaProjectionExpression(ctx *SchemaProjectionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixRightApplicationExpression(ctx *InfixRightApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitCharacteristicSetComprehensionExpression(ctx *CharacteristicSetComprehensionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaPreconditionExpression(ctx *SchemaPreconditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixApplicationExpression(ctx *PrefixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitBindingConstructionExpression(ctx *BindingConstructionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaPipingExpression(ctx *SchemaPipingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaImplicationExpression(ctx *SchemaImplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitBindingExtensionExpression(ctx *BindingExtensionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaDecorationExpression(ctx *SchemaDecorationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaUniqueExistentialQuantificationExpression(ctx *SchemaUniqueExistentialQuantificationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDefiniteDescriptionExpression(ctx *DefiniteDescriptionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitReferenceExpression(ctx *ReferenceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenericInfixApplicationExpression(ctx *GenericInfixApplicationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitTupleSelectionExpression(ctx *TupleSelectionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitSchemaText(ctx *SchemaTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDeclPart(ctx *DeclPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDeclNameExpression(ctx *DeclNameExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitOperatorTemplate(ctx *OperatorTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitCategoryTemplate(ctx *CategoryTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrec(ctx *PrecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitAssoc(ctx *AssocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixTemplate(ctx *PrefixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixTemplate(ctx *PostfixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixTemplate(ctx *InfixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixTemplate(ctx *NofixTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitDeclName(ctx *DeclNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitRefName(ctx *RefNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitOpName(ctx *OpNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixName(ctx *PrefixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixName(ctx *PostfixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixName(ctx *InfixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixName(ctx *NofixNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitGenName(ctx *GenNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixGenName(ctx *PrefixGenNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixGenName(ctx *PostfixGenNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixGenName(ctx *InfixGenNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixGenName(ctx *NofixGenNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixRel(ctx *PrefixRelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixRel(ctx *PostfixRelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixRel(ctx *InfixRelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixRel(ctx *NofixRelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitApplication(ctx *ApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPrefixApp(ctx *PrefixAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitPostfixApp(ctx *PostfixAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitInfixApp(ctx *InfixAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitNofixApp(ctx *NofixAppContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitExpSep(ctx *ExpSepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}
