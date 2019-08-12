// Code generated from PROV_N.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prov_n // PROV_N
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PROV_NListener is a complete listener for a parse tree produced by PROV_NParser.
type PROV_NListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterNamespaceDeclarations is called when entering the namespaceDeclarations production.
	EnterNamespaceDeclarations(c *NamespaceDeclarationsContext)

	// EnterDefaultNamespaceDeclaration is called when entering the defaultNamespaceDeclaration production.
	EnterDefaultNamespaceDeclaration(c *DefaultNamespaceDeclarationContext)

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterBundle is called when entering the bundle production.
	EnterBundle(c *BundleContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEntityExpression is called when entering the entityExpression production.
	EnterEntityExpression(c *EntityExpressionContext)

	// EnterOptionalAttributeValuePairs is called when entering the optionalAttributeValuePairs production.
	EnterOptionalAttributeValuePairs(c *OptionalAttributeValuePairsContext)

	// EnterAttributeValuePairs is called when entering the attributeValuePairs production.
	EnterAttributeValuePairs(c *AttributeValuePairsContext)

	// EnterAttributeValuePair is called when entering the attributeValuePair production.
	EnterAttributeValuePair(c *AttributeValuePairContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterTypedLiteral is called when entering the typedLiteral production.
	EnterTypedLiteral(c *TypedLiteralContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterConvenienceNotation is called when entering the convenienceNotation production.
	EnterConvenienceNotation(c *ConvenienceNotationContext)

	// EnterActivityExpression is called when entering the activityExpression production.
	EnterActivityExpression(c *ActivityExpressionContext)

	// EnterTimeOrMarker is called when entering the timeOrMarker production.
	EnterTimeOrMarker(c *TimeOrMarkerContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterGenerationExpression is called when entering the generationExpression production.
	EnterGenerationExpression(c *GenerationExpressionContext)

	// EnterOptionalIdentifier is called when entering the optionalIdentifier production.
	EnterOptionalIdentifier(c *OptionalIdentifierContext)

	// EnterIdentifierOrMarker is called when entering the identifierOrMarker production.
	EnterIdentifierOrMarker(c *IdentifierOrMarkerContext)

	// EnterEIdentifier is called when entering the eIdentifier production.
	EnterEIdentifier(c *EIdentifierContext)

	// EnterEIdentifierOrMarker is called when entering the eIdentifierOrMarker production.
	EnterEIdentifierOrMarker(c *EIdentifierOrMarkerContext)

	// EnterAIdentifierOrMarker is called when entering the aIdentifierOrMarker production.
	EnterAIdentifierOrMarker(c *AIdentifierOrMarkerContext)

	// EnterAIdentifier is called when entering the aIdentifier production.
	EnterAIdentifier(c *AIdentifierContext)

	// EnterAgIdentifierOrMarker is called when entering the agIdentifierOrMarker production.
	EnterAgIdentifierOrMarker(c *AgIdentifierOrMarkerContext)

	// EnterAgIdentifier is called when entering the agIdentifier production.
	EnterAgIdentifier(c *AgIdentifierContext)

	// EnterCIdentifier is called when entering the cIdentifier production.
	EnterCIdentifier(c *CIdentifierContext)

	// EnterGIdentifier is called when entering the gIdentifier production.
	EnterGIdentifier(c *GIdentifierContext)

	// EnterGIdentifierOrMarker is called when entering the gIdentifierOrMarker production.
	EnterGIdentifierOrMarker(c *GIdentifierOrMarkerContext)

	// EnterUIdentifier is called when entering the uIdentifier production.
	EnterUIdentifier(c *UIdentifierContext)

	// EnterUIdentifierOrMarker is called when entering the uIdentifierOrMarker production.
	EnterUIdentifierOrMarker(c *UIdentifierOrMarkerContext)

	// EnterUsageExpression is called when entering the usageExpression production.
	EnterUsageExpression(c *UsageExpressionContext)

	// EnterStartExpression is called when entering the startExpression production.
	EnterStartExpression(c *StartExpressionContext)

	// EnterEndExpression is called when entering the endExpression production.
	EnterEndExpression(c *EndExpressionContext)

	// EnterInvalidationExpression is called when entering the invalidationExpression production.
	EnterInvalidationExpression(c *InvalidationExpressionContext)

	// EnterCommunicationExpression is called when entering the communicationExpression production.
	EnterCommunicationExpression(c *CommunicationExpressionContext)

	// EnterAgentExpression is called when entering the agentExpression production.
	EnterAgentExpression(c *AgentExpressionContext)

	// EnterAssociationExpression is called when entering the associationExpression production.
	EnterAssociationExpression(c *AssociationExpressionContext)

	// EnterAttributionExpression is called when entering the attributionExpression production.
	EnterAttributionExpression(c *AttributionExpressionContext)

	// EnterDelegationExpression is called when entering the delegationExpression production.
	EnterDelegationExpression(c *DelegationExpressionContext)

	// EnterDerivationExpression is called when entering the derivationExpression production.
	EnterDerivationExpression(c *DerivationExpressionContext)

	// EnterInfluenceExpression is called when entering the influenceExpression production.
	EnterInfluenceExpression(c *InfluenceExpressionContext)

	// EnterAlternateExpression is called when entering the alternateExpression production.
	EnterAlternateExpression(c *AlternateExpressionContext)

	// EnterSpecializationExpression is called when entering the specializationExpression production.
	EnterSpecializationExpression(c *SpecializationExpressionContext)

	// EnterMembershipExpression is called when entering the membershipExpression production.
	EnterMembershipExpression(c *MembershipExpressionContext)

	// EnterExtensibilityExpression is called when entering the extensibilityExpression production.
	EnterExtensibilityExpression(c *ExtensibilityExpressionContext)

	// EnterExtensibilityArgument is called when entering the extensibilityArgument production.
	EnterExtensibilityArgument(c *ExtensibilityArgumentContext)

	// EnterExtensibilityTuple is called when entering the extensibilityTuple production.
	EnterExtensibilityTuple(c *ExtensibilityTupleContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitNamespaceDeclarations is called when exiting the namespaceDeclarations production.
	ExitNamespaceDeclarations(c *NamespaceDeclarationsContext)

	// ExitDefaultNamespaceDeclaration is called when exiting the defaultNamespaceDeclaration production.
	ExitDefaultNamespaceDeclaration(c *DefaultNamespaceDeclarationContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitBundle is called when exiting the bundle production.
	ExitBundle(c *BundleContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEntityExpression is called when exiting the entityExpression production.
	ExitEntityExpression(c *EntityExpressionContext)

	// ExitOptionalAttributeValuePairs is called when exiting the optionalAttributeValuePairs production.
	ExitOptionalAttributeValuePairs(c *OptionalAttributeValuePairsContext)

	// ExitAttributeValuePairs is called when exiting the attributeValuePairs production.
	ExitAttributeValuePairs(c *AttributeValuePairsContext)

	// ExitAttributeValuePair is called when exiting the attributeValuePair production.
	ExitAttributeValuePair(c *AttributeValuePairContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitTypedLiteral is called when exiting the typedLiteral production.
	ExitTypedLiteral(c *TypedLiteralContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitConvenienceNotation is called when exiting the convenienceNotation production.
	ExitConvenienceNotation(c *ConvenienceNotationContext)

	// ExitActivityExpression is called when exiting the activityExpression production.
	ExitActivityExpression(c *ActivityExpressionContext)

	// ExitTimeOrMarker is called when exiting the timeOrMarker production.
	ExitTimeOrMarker(c *TimeOrMarkerContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitGenerationExpression is called when exiting the generationExpression production.
	ExitGenerationExpression(c *GenerationExpressionContext)

	// ExitOptionalIdentifier is called when exiting the optionalIdentifier production.
	ExitOptionalIdentifier(c *OptionalIdentifierContext)

	// ExitIdentifierOrMarker is called when exiting the identifierOrMarker production.
	ExitIdentifierOrMarker(c *IdentifierOrMarkerContext)

	// ExitEIdentifier is called when exiting the eIdentifier production.
	ExitEIdentifier(c *EIdentifierContext)

	// ExitEIdentifierOrMarker is called when exiting the eIdentifierOrMarker production.
	ExitEIdentifierOrMarker(c *EIdentifierOrMarkerContext)

	// ExitAIdentifierOrMarker is called when exiting the aIdentifierOrMarker production.
	ExitAIdentifierOrMarker(c *AIdentifierOrMarkerContext)

	// ExitAIdentifier is called when exiting the aIdentifier production.
	ExitAIdentifier(c *AIdentifierContext)

	// ExitAgIdentifierOrMarker is called when exiting the agIdentifierOrMarker production.
	ExitAgIdentifierOrMarker(c *AgIdentifierOrMarkerContext)

	// ExitAgIdentifier is called when exiting the agIdentifier production.
	ExitAgIdentifier(c *AgIdentifierContext)

	// ExitCIdentifier is called when exiting the cIdentifier production.
	ExitCIdentifier(c *CIdentifierContext)

	// ExitGIdentifier is called when exiting the gIdentifier production.
	ExitGIdentifier(c *GIdentifierContext)

	// ExitGIdentifierOrMarker is called when exiting the gIdentifierOrMarker production.
	ExitGIdentifierOrMarker(c *GIdentifierOrMarkerContext)

	// ExitUIdentifier is called when exiting the uIdentifier production.
	ExitUIdentifier(c *UIdentifierContext)

	// ExitUIdentifierOrMarker is called when exiting the uIdentifierOrMarker production.
	ExitUIdentifierOrMarker(c *UIdentifierOrMarkerContext)

	// ExitUsageExpression is called when exiting the usageExpression production.
	ExitUsageExpression(c *UsageExpressionContext)

	// ExitStartExpression is called when exiting the startExpression production.
	ExitStartExpression(c *StartExpressionContext)

	// ExitEndExpression is called when exiting the endExpression production.
	ExitEndExpression(c *EndExpressionContext)

	// ExitInvalidationExpression is called when exiting the invalidationExpression production.
	ExitInvalidationExpression(c *InvalidationExpressionContext)

	// ExitCommunicationExpression is called when exiting the communicationExpression production.
	ExitCommunicationExpression(c *CommunicationExpressionContext)

	// ExitAgentExpression is called when exiting the agentExpression production.
	ExitAgentExpression(c *AgentExpressionContext)

	// ExitAssociationExpression is called when exiting the associationExpression production.
	ExitAssociationExpression(c *AssociationExpressionContext)

	// ExitAttributionExpression is called when exiting the attributionExpression production.
	ExitAttributionExpression(c *AttributionExpressionContext)

	// ExitDelegationExpression is called when exiting the delegationExpression production.
	ExitDelegationExpression(c *DelegationExpressionContext)

	// ExitDerivationExpression is called when exiting the derivationExpression production.
	ExitDerivationExpression(c *DerivationExpressionContext)

	// ExitInfluenceExpression is called when exiting the influenceExpression production.
	ExitInfluenceExpression(c *InfluenceExpressionContext)

	// ExitAlternateExpression is called when exiting the alternateExpression production.
	ExitAlternateExpression(c *AlternateExpressionContext)

	// ExitSpecializationExpression is called when exiting the specializationExpression production.
	ExitSpecializationExpression(c *SpecializationExpressionContext)

	// ExitMembershipExpression is called when exiting the membershipExpression production.
	ExitMembershipExpression(c *MembershipExpressionContext)

	// ExitExtensibilityExpression is called when exiting the extensibilityExpression production.
	ExitExtensibilityExpression(c *ExtensibilityExpressionContext)

	// ExitExtensibilityArgument is called when exiting the extensibilityArgument production.
	ExitExtensibilityArgument(c *ExtensibilityArgumentContext)

	// ExitExtensibilityTuple is called when exiting the extensibilityTuple production.
	ExitExtensibilityTuple(c *ExtensibilityTupleContext)
}
