// Code generated from PROV_N.g4 by ANTLR 4.9.3. DO NOT EDIT.

package prov_n // PROV_N
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePROV_NListener is a complete listener for a parse tree produced by PROV_NParser.
type BasePROV_NListener struct{}

var _ PROV_NListener = &BasePROV_NListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePROV_NListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePROV_NListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePROV_NListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePROV_NListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BasePROV_NListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BasePROV_NListener) ExitDocument(ctx *DocumentContext) {}

// EnterNamespaceDeclarations is called when production namespaceDeclarations is entered.
func (s *BasePROV_NListener) EnterNamespaceDeclarations(ctx *NamespaceDeclarationsContext) {}

// ExitNamespaceDeclarations is called when production namespaceDeclarations is exited.
func (s *BasePROV_NListener) ExitNamespaceDeclarations(ctx *NamespaceDeclarationsContext) {}

// EnterDefaultNamespaceDeclaration is called when production defaultNamespaceDeclaration is entered.
func (s *BasePROV_NListener) EnterDefaultNamespaceDeclaration(ctx *DefaultNamespaceDeclarationContext) {
}

// ExitDefaultNamespaceDeclaration is called when production defaultNamespaceDeclaration is exited.
func (s *BasePROV_NListener) ExitDefaultNamespaceDeclaration(ctx *DefaultNamespaceDeclarationContext) {
}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BasePROV_NListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BasePROV_NListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// EnterNamespace_ is called when production namespace_ is entered.
func (s *BasePROV_NListener) EnterNamespace_(ctx *Namespace_Context) {}

// ExitNamespace_ is called when production namespace_ is exited.
func (s *BasePROV_NListener) ExitNamespace_(ctx *Namespace_Context) {}

// EnterBundle is called when production bundle is entered.
func (s *BasePROV_NListener) EnterBundle(ctx *BundleContext) {}

// ExitBundle is called when production bundle is exited.
func (s *BasePROV_NListener) ExitBundle(ctx *BundleContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePROV_NListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePROV_NListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePROV_NListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePROV_NListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEntityExpression is called when production entityExpression is entered.
func (s *BasePROV_NListener) EnterEntityExpression(ctx *EntityExpressionContext) {}

// ExitEntityExpression is called when production entityExpression is exited.
func (s *BasePROV_NListener) ExitEntityExpression(ctx *EntityExpressionContext) {}

// EnterOptionalAttributeValuePairs is called when production optionalAttributeValuePairs is entered.
func (s *BasePROV_NListener) EnterOptionalAttributeValuePairs(ctx *OptionalAttributeValuePairsContext) {
}

// ExitOptionalAttributeValuePairs is called when production optionalAttributeValuePairs is exited.
func (s *BasePROV_NListener) ExitOptionalAttributeValuePairs(ctx *OptionalAttributeValuePairsContext) {
}

// EnterAttributeValuePairs is called when production attributeValuePairs is entered.
func (s *BasePROV_NListener) EnterAttributeValuePairs(ctx *AttributeValuePairsContext) {}

// ExitAttributeValuePairs is called when production attributeValuePairs is exited.
func (s *BasePROV_NListener) ExitAttributeValuePairs(ctx *AttributeValuePairsContext) {}

// EnterAttributeValuePair is called when production attributeValuePair is entered.
func (s *BasePROV_NListener) EnterAttributeValuePair(ctx *AttributeValuePairContext) {}

// ExitAttributeValuePair is called when production attributeValuePair is exited.
func (s *BasePROV_NListener) ExitAttributeValuePair(ctx *AttributeValuePairContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasePROV_NListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasePROV_NListener) ExitAttribute(ctx *AttributeContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePROV_NListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePROV_NListener) ExitLiteral(ctx *LiteralContext) {}

// EnterTypedLiteral is called when production typedLiteral is entered.
func (s *BasePROV_NListener) EnterTypedLiteral(ctx *TypedLiteralContext) {}

// ExitTypedLiteral is called when production typedLiteral is exited.
func (s *BasePROV_NListener) ExitTypedLiteral(ctx *TypedLiteralContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BasePROV_NListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BasePROV_NListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterConvenienceNotation is called when production convenienceNotation is entered.
func (s *BasePROV_NListener) EnterConvenienceNotation(ctx *ConvenienceNotationContext) {}

// ExitConvenienceNotation is called when production convenienceNotation is exited.
func (s *BasePROV_NListener) ExitConvenienceNotation(ctx *ConvenienceNotationContext) {}

// EnterActivityExpression is called when production activityExpression is entered.
func (s *BasePROV_NListener) EnterActivityExpression(ctx *ActivityExpressionContext) {}

// ExitActivityExpression is called when production activityExpression is exited.
func (s *BasePROV_NListener) ExitActivityExpression(ctx *ActivityExpressionContext) {}

// EnterTimeOrMarker is called when production timeOrMarker is entered.
func (s *BasePROV_NListener) EnterTimeOrMarker(ctx *TimeOrMarkerContext) {}

// ExitTimeOrMarker is called when production timeOrMarker is exited.
func (s *BasePROV_NListener) ExitTimeOrMarker(ctx *TimeOrMarkerContext) {}

// EnterTime is called when production time is entered.
func (s *BasePROV_NListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BasePROV_NListener) ExitTime(ctx *TimeContext) {}

// EnterGenerationExpression is called when production generationExpression is entered.
func (s *BasePROV_NListener) EnterGenerationExpression(ctx *GenerationExpressionContext) {}

// ExitGenerationExpression is called when production generationExpression is exited.
func (s *BasePROV_NListener) ExitGenerationExpression(ctx *GenerationExpressionContext) {}

// EnterOptionalIdentifier is called when production optionalIdentifier is entered.
func (s *BasePROV_NListener) EnterOptionalIdentifier(ctx *OptionalIdentifierContext) {}

// ExitOptionalIdentifier is called when production optionalIdentifier is exited.
func (s *BasePROV_NListener) ExitOptionalIdentifier(ctx *OptionalIdentifierContext) {}

// EnterIdentifierOrMarker is called when production identifierOrMarker is entered.
func (s *BasePROV_NListener) EnterIdentifierOrMarker(ctx *IdentifierOrMarkerContext) {}

// ExitIdentifierOrMarker is called when production identifierOrMarker is exited.
func (s *BasePROV_NListener) ExitIdentifierOrMarker(ctx *IdentifierOrMarkerContext) {}

// EnterEIdentifier is called when production eIdentifier is entered.
func (s *BasePROV_NListener) EnterEIdentifier(ctx *EIdentifierContext) {}

// ExitEIdentifier is called when production eIdentifier is exited.
func (s *BasePROV_NListener) ExitEIdentifier(ctx *EIdentifierContext) {}

// EnterEIdentifierOrMarker is called when production eIdentifierOrMarker is entered.
func (s *BasePROV_NListener) EnterEIdentifierOrMarker(ctx *EIdentifierOrMarkerContext) {}

// ExitEIdentifierOrMarker is called when production eIdentifierOrMarker is exited.
func (s *BasePROV_NListener) ExitEIdentifierOrMarker(ctx *EIdentifierOrMarkerContext) {}

// EnterAIdentifierOrMarker is called when production aIdentifierOrMarker is entered.
func (s *BasePROV_NListener) EnterAIdentifierOrMarker(ctx *AIdentifierOrMarkerContext) {}

// ExitAIdentifierOrMarker is called when production aIdentifierOrMarker is exited.
func (s *BasePROV_NListener) ExitAIdentifierOrMarker(ctx *AIdentifierOrMarkerContext) {}

// EnterAIdentifier is called when production aIdentifier is entered.
func (s *BasePROV_NListener) EnterAIdentifier(ctx *AIdentifierContext) {}

// ExitAIdentifier is called when production aIdentifier is exited.
func (s *BasePROV_NListener) ExitAIdentifier(ctx *AIdentifierContext) {}

// EnterAgIdentifierOrMarker is called when production agIdentifierOrMarker is entered.
func (s *BasePROV_NListener) EnterAgIdentifierOrMarker(ctx *AgIdentifierOrMarkerContext) {}

// ExitAgIdentifierOrMarker is called when production agIdentifierOrMarker is exited.
func (s *BasePROV_NListener) ExitAgIdentifierOrMarker(ctx *AgIdentifierOrMarkerContext) {}

// EnterAgIdentifier is called when production agIdentifier is entered.
func (s *BasePROV_NListener) EnterAgIdentifier(ctx *AgIdentifierContext) {}

// ExitAgIdentifier is called when production agIdentifier is exited.
func (s *BasePROV_NListener) ExitAgIdentifier(ctx *AgIdentifierContext) {}

// EnterCIdentifier is called when production cIdentifier is entered.
func (s *BasePROV_NListener) EnterCIdentifier(ctx *CIdentifierContext) {}

// ExitCIdentifier is called when production cIdentifier is exited.
func (s *BasePROV_NListener) ExitCIdentifier(ctx *CIdentifierContext) {}

// EnterGIdentifier is called when production gIdentifier is entered.
func (s *BasePROV_NListener) EnterGIdentifier(ctx *GIdentifierContext) {}

// ExitGIdentifier is called when production gIdentifier is exited.
func (s *BasePROV_NListener) ExitGIdentifier(ctx *GIdentifierContext) {}

// EnterGIdentifierOrMarker is called when production gIdentifierOrMarker is entered.
func (s *BasePROV_NListener) EnterGIdentifierOrMarker(ctx *GIdentifierOrMarkerContext) {}

// ExitGIdentifierOrMarker is called when production gIdentifierOrMarker is exited.
func (s *BasePROV_NListener) ExitGIdentifierOrMarker(ctx *GIdentifierOrMarkerContext) {}

// EnterUIdentifier is called when production uIdentifier is entered.
func (s *BasePROV_NListener) EnterUIdentifier(ctx *UIdentifierContext) {}

// ExitUIdentifier is called when production uIdentifier is exited.
func (s *BasePROV_NListener) ExitUIdentifier(ctx *UIdentifierContext) {}

// EnterUIdentifierOrMarker is called when production uIdentifierOrMarker is entered.
func (s *BasePROV_NListener) EnterUIdentifierOrMarker(ctx *UIdentifierOrMarkerContext) {}

// ExitUIdentifierOrMarker is called when production uIdentifierOrMarker is exited.
func (s *BasePROV_NListener) ExitUIdentifierOrMarker(ctx *UIdentifierOrMarkerContext) {}

// EnterUsageExpression is called when production usageExpression is entered.
func (s *BasePROV_NListener) EnterUsageExpression(ctx *UsageExpressionContext) {}

// ExitUsageExpression is called when production usageExpression is exited.
func (s *BasePROV_NListener) ExitUsageExpression(ctx *UsageExpressionContext) {}

// EnterStartExpression is called when production startExpression is entered.
func (s *BasePROV_NListener) EnterStartExpression(ctx *StartExpressionContext) {}

// ExitStartExpression is called when production startExpression is exited.
func (s *BasePROV_NListener) ExitStartExpression(ctx *StartExpressionContext) {}

// EnterEndExpression is called when production endExpression is entered.
func (s *BasePROV_NListener) EnterEndExpression(ctx *EndExpressionContext) {}

// ExitEndExpression is called when production endExpression is exited.
func (s *BasePROV_NListener) ExitEndExpression(ctx *EndExpressionContext) {}

// EnterInvalidationExpression is called when production invalidationExpression is entered.
func (s *BasePROV_NListener) EnterInvalidationExpression(ctx *InvalidationExpressionContext) {}

// ExitInvalidationExpression is called when production invalidationExpression is exited.
func (s *BasePROV_NListener) ExitInvalidationExpression(ctx *InvalidationExpressionContext) {}

// EnterCommunicationExpression is called when production communicationExpression is entered.
func (s *BasePROV_NListener) EnterCommunicationExpression(ctx *CommunicationExpressionContext) {}

// ExitCommunicationExpression is called when production communicationExpression is exited.
func (s *BasePROV_NListener) ExitCommunicationExpression(ctx *CommunicationExpressionContext) {}

// EnterAgentExpression is called when production agentExpression is entered.
func (s *BasePROV_NListener) EnterAgentExpression(ctx *AgentExpressionContext) {}

// ExitAgentExpression is called when production agentExpression is exited.
func (s *BasePROV_NListener) ExitAgentExpression(ctx *AgentExpressionContext) {}

// EnterAssociationExpression is called when production associationExpression is entered.
func (s *BasePROV_NListener) EnterAssociationExpression(ctx *AssociationExpressionContext) {}

// ExitAssociationExpression is called when production associationExpression is exited.
func (s *BasePROV_NListener) ExitAssociationExpression(ctx *AssociationExpressionContext) {}

// EnterAttributionExpression is called when production attributionExpression is entered.
func (s *BasePROV_NListener) EnterAttributionExpression(ctx *AttributionExpressionContext) {}

// ExitAttributionExpression is called when production attributionExpression is exited.
func (s *BasePROV_NListener) ExitAttributionExpression(ctx *AttributionExpressionContext) {}

// EnterDelegationExpression is called when production delegationExpression is entered.
func (s *BasePROV_NListener) EnterDelegationExpression(ctx *DelegationExpressionContext) {}

// ExitDelegationExpression is called when production delegationExpression is exited.
func (s *BasePROV_NListener) ExitDelegationExpression(ctx *DelegationExpressionContext) {}

// EnterDerivationExpression is called when production derivationExpression is entered.
func (s *BasePROV_NListener) EnterDerivationExpression(ctx *DerivationExpressionContext) {}

// ExitDerivationExpression is called when production derivationExpression is exited.
func (s *BasePROV_NListener) ExitDerivationExpression(ctx *DerivationExpressionContext) {}

// EnterInfluenceExpression is called when production influenceExpression is entered.
func (s *BasePROV_NListener) EnterInfluenceExpression(ctx *InfluenceExpressionContext) {}

// ExitInfluenceExpression is called when production influenceExpression is exited.
func (s *BasePROV_NListener) ExitInfluenceExpression(ctx *InfluenceExpressionContext) {}

// EnterAlternateExpression is called when production alternateExpression is entered.
func (s *BasePROV_NListener) EnterAlternateExpression(ctx *AlternateExpressionContext) {}

// ExitAlternateExpression is called when production alternateExpression is exited.
func (s *BasePROV_NListener) ExitAlternateExpression(ctx *AlternateExpressionContext) {}

// EnterSpecializationExpression is called when production specializationExpression is entered.
func (s *BasePROV_NListener) EnterSpecializationExpression(ctx *SpecializationExpressionContext) {}

// ExitSpecializationExpression is called when production specializationExpression is exited.
func (s *BasePROV_NListener) ExitSpecializationExpression(ctx *SpecializationExpressionContext) {}

// EnterMembershipExpression is called when production membershipExpression is entered.
func (s *BasePROV_NListener) EnterMembershipExpression(ctx *MembershipExpressionContext) {}

// ExitMembershipExpression is called when production membershipExpression is exited.
func (s *BasePROV_NListener) ExitMembershipExpression(ctx *MembershipExpressionContext) {}

// EnterExtensibilityExpression is called when production extensibilityExpression is entered.
func (s *BasePROV_NListener) EnterExtensibilityExpression(ctx *ExtensibilityExpressionContext) {}

// ExitExtensibilityExpression is called when production extensibilityExpression is exited.
func (s *BasePROV_NListener) ExitExtensibilityExpression(ctx *ExtensibilityExpressionContext) {}

// EnterExtensibilityArgument is called when production extensibilityArgument is entered.
func (s *BasePROV_NListener) EnterExtensibilityArgument(ctx *ExtensibilityArgumentContext) {}

// ExitExtensibilityArgument is called when production extensibilityArgument is exited.
func (s *BasePROV_NListener) ExitExtensibilityArgument(ctx *ExtensibilityArgumentContext) {}

// EnterExtensibilityTuple is called when production extensibilityTuple is entered.
func (s *BasePROV_NListener) EnterExtensibilityTuple(ctx *ExtensibilityTupleContext) {}

// ExitExtensibilityTuple is called when production extensibilityTuple is exited.
func (s *BasePROV_NListener) ExitExtensibilityTuple(ctx *ExtensibilityTupleContext) {}
