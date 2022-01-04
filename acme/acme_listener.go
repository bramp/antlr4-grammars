// Code generated from acme.g4 by ANTLR 4.9.3. DO NOT EDIT.

package acme // acme
import "github.com/antlr/antlr4/runtime/Go/antlr"

// acmeListener is a complete listener for a parse tree produced by acmeParser.
type acmeListener interface {
	antlr.ParseTreeListener

	// EnterAcmeCompUnit is called when entering the acmeCompUnit production.
	EnterAcmeCompUnit(c *AcmeCompUnitContext)

	// EnterAcmeImportDeclaration is called when entering the acmeImportDeclaration production.
	EnterAcmeImportDeclaration(c *AcmeImportDeclarationContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterCodeLiteral is called when entering the codeLiteral production.
	EnterCodeLiteral(c *CodeLiteralContext)

	// EnterAcmeFamilyDeclaration is called when entering the acmeFamilyDeclaration production.
	EnterAcmeFamilyDeclaration(c *AcmeFamilyDeclarationContext)

	// EnterAcmeFamilyBody is called when entering the acmeFamilyBody production.
	EnterAcmeFamilyBody(c *AcmeFamilyBodyContext)

	// EnterAcmeSystemDeclaration is called when entering the acmeSystemDeclaration production.
	EnterAcmeSystemDeclaration(c *AcmeSystemDeclarationContext)

	// EnterAcmeSystemBody is called when entering the acmeSystemBody production.
	EnterAcmeSystemBody(c *AcmeSystemBodyContext)

	// EnterAcmeDesignDeclaration is called when entering the acmeDesignDeclaration production.
	EnterAcmeDesignDeclaration(c *AcmeDesignDeclarationContext)

	// EnterAcmeComponentTypeRef is called when entering the acmeComponentTypeRef production.
	EnterAcmeComponentTypeRef(c *AcmeComponentTypeRefContext)

	// EnterAcmeComponentInstantiatedTypeRef is called when entering the acmeComponentInstantiatedTypeRef production.
	EnterAcmeComponentInstantiatedTypeRef(c *AcmeComponentInstantiatedTypeRefContext)

	// EnterAcmeConnectorTypeRef is called when entering the acmeConnectorTypeRef production.
	EnterAcmeConnectorTypeRef(c *AcmeConnectorTypeRefContext)

	// EnterAcmeConnectorInstantiatedTypeRef is called when entering the acmeConnectorInstantiatedTypeRef production.
	EnterAcmeConnectorInstantiatedTypeRef(c *AcmeConnectorInstantiatedTypeRefContext)

	// EnterAcmePortTypeRef is called when entering the acmePortTypeRef production.
	EnterAcmePortTypeRef(c *AcmePortTypeRefContext)

	// EnterAcmePortInstantiatedTypeRef is called when entering the acmePortInstantiatedTypeRef production.
	EnterAcmePortInstantiatedTypeRef(c *AcmePortInstantiatedTypeRefContext)

	// EnterAcmeGroupTypeRef is called when entering the acmeGroupTypeRef production.
	EnterAcmeGroupTypeRef(c *AcmeGroupTypeRefContext)

	// EnterAcmeGroupInstantiatedTypeRef is called when entering the acmeGroupInstantiatedTypeRef production.
	EnterAcmeGroupInstantiatedTypeRef(c *AcmeGroupInstantiatedTypeRefContext)

	// EnterAcmeRoleTypeRef is called when entering the acmeRoleTypeRef production.
	EnterAcmeRoleTypeRef(c *AcmeRoleTypeRefContext)

	// EnterAcmeRoleInstantiatedTypeRef is called when entering the acmeRoleInstantiatedTypeRef production.
	EnterAcmeRoleInstantiatedTypeRef(c *AcmeRoleInstantiatedTypeRefContext)

	// EnterAcmeViewTypeRef is called when entering the acmeViewTypeRef production.
	EnterAcmeViewTypeRef(c *AcmeViewTypeRefContext)

	// EnterAcmeViewInstantiatedTypeRef is called when entering the acmeViewInstantiatedTypeRef production.
	EnterAcmeViewInstantiatedTypeRef(c *AcmeViewInstantiatedTypeRefContext)

	// EnterAcmeFamilyRef is called when entering the acmeFamilyRef production.
	EnterAcmeFamilyRef(c *AcmeFamilyRefContext)

	// EnterAcmeFamilyInstantiationRef is called when entering the acmeFamilyInstantiationRef production.
	EnterAcmeFamilyInstantiationRef(c *AcmeFamilyInstantiationRefContext)

	// EnterAcmeElementTypeRef is called when entering the acmeElementTypeRef production.
	EnterAcmeElementTypeRef(c *AcmeElementTypeRefContext)

	// EnterAcmePropertyTypeDeclarationRef is called when entering the acmePropertyTypeDeclarationRef production.
	EnterAcmePropertyTypeDeclarationRef(c *AcmePropertyTypeDeclarationRefContext)

	// EnterAcmeInstanceRef is called when entering the acmeInstanceRef production.
	EnterAcmeInstanceRef(c *AcmeInstanceRefContext)

	// EnterAcmeGenericElementTypeDeclaration is called when entering the acmeGenericElementTypeDeclaration production.
	EnterAcmeGenericElementTypeDeclaration(c *AcmeGenericElementTypeDeclarationContext)

	// EnterAcmeGenericElementBody is called when entering the acmeGenericElementBody production.
	EnterAcmeGenericElementBody(c *AcmeGenericElementBodyContext)

	// EnterAcmeGroupTypeDeclaration is called when entering the acmeGroupTypeDeclaration production.
	EnterAcmeGroupTypeDeclaration(c *AcmeGroupTypeDeclarationContext)

	// EnterAcmeGroupDeclaration is called when entering the acmeGroupDeclaration production.
	EnterAcmeGroupDeclaration(c *AcmeGroupDeclarationContext)

	// EnterAcmeGroupBody is called when entering the acmeGroupBody production.
	EnterAcmeGroupBody(c *AcmeGroupBodyContext)

	// EnterAcmeMembersBlock is called when entering the acmeMembersBlock production.
	EnterAcmeMembersBlock(c *AcmeMembersBlockContext)

	// EnterAcmePortTypeDeclaration is called when entering the acmePortTypeDeclaration production.
	EnterAcmePortTypeDeclaration(c *AcmePortTypeDeclarationContext)

	// EnterAcmePortDeclaration is called when entering the acmePortDeclaration production.
	EnterAcmePortDeclaration(c *AcmePortDeclarationContext)

	// EnterAcmePortBody is called when entering the acmePortBody production.
	EnterAcmePortBody(c *AcmePortBodyContext)

	// EnterAcmeRoleTypeDeclaration is called when entering the acmeRoleTypeDeclaration production.
	EnterAcmeRoleTypeDeclaration(c *AcmeRoleTypeDeclarationContext)

	// EnterAcmeRoleDeclaration is called when entering the acmeRoleDeclaration production.
	EnterAcmeRoleDeclaration(c *AcmeRoleDeclarationContext)

	// EnterAcmeRoleBody is called when entering the acmeRoleBody production.
	EnterAcmeRoleBody(c *AcmeRoleBodyContext)

	// EnterAcmeComponentTypeDeclaration is called when entering the acmeComponentTypeDeclaration production.
	EnterAcmeComponentTypeDeclaration(c *AcmeComponentTypeDeclarationContext)

	// EnterAcmeComponentDeclaration is called when entering the acmeComponentDeclaration production.
	EnterAcmeComponentDeclaration(c *AcmeComponentDeclarationContext)

	// EnterAcmeComponentBody is called when entering the acmeComponentBody production.
	EnterAcmeComponentBody(c *AcmeComponentBodyContext)

	// EnterAcmeConnectorTypeDeclaration is called when entering the acmeConnectorTypeDeclaration production.
	EnterAcmeConnectorTypeDeclaration(c *AcmeConnectorTypeDeclarationContext)

	// EnterAcmeConnectorDeclaration is called when entering the acmeConnectorDeclaration production.
	EnterAcmeConnectorDeclaration(c *AcmeConnectorDeclarationContext)

	// EnterAcmeConnectorBody is called when entering the acmeConnectorBody production.
	EnterAcmeConnectorBody(c *AcmeConnectorBodyContext)

	// EnterAcmeRepresentationDeclaration is called when entering the acmeRepresentationDeclaration production.
	EnterAcmeRepresentationDeclaration(c *AcmeRepresentationDeclarationContext)

	// EnterAcmeBindingsMapDeclaration is called when entering the acmeBindingsMapDeclaration production.
	EnterAcmeBindingsMapDeclaration(c *AcmeBindingsMapDeclarationContext)

	// EnterAcmeBindingDeclaration is called when entering the acmeBindingDeclaration production.
	EnterAcmeBindingDeclaration(c *AcmeBindingDeclarationContext)

	// EnterAcmeAttachmentDeclaration is called when entering the acmeAttachmentDeclaration production.
	EnterAcmeAttachmentDeclaration(c *AcmeAttachmentDeclarationContext)

	// EnterAcmePropertyDeclaration is called when entering the acmePropertyDeclaration production.
	EnterAcmePropertyDeclaration(c *AcmePropertyDeclarationContext)

	// EnterAcmePropertyValueDeclaration is called when entering the acmePropertyValueDeclaration production.
	EnterAcmePropertyValueDeclaration(c *AcmePropertyValueDeclarationContext)

	// EnterEnumidentifier is called when entering the enumidentifier production.
	EnterEnumidentifier(c *EnumidentifierContext)

	// EnterAcmePropertyElement is called when entering the acmePropertyElement production.
	EnterAcmePropertyElement(c *AcmePropertyElementContext)

	// EnterAcmePropertyCompoundElement is called when entering the acmePropertyCompoundElement production.
	EnterAcmePropertyCompoundElement(c *AcmePropertyCompoundElementContext)

	// EnterAcmePropertySet is called when entering the acmePropertySet production.
	EnterAcmePropertySet(c *AcmePropertySetContext)

	// EnterAcmePropertyRecordEntry is called when entering the acmePropertyRecordEntry production.
	EnterAcmePropertyRecordEntry(c *AcmePropertyRecordEntryContext)

	// EnterAcmePropertyRecord is called when entering the acmePropertyRecord production.
	EnterAcmePropertyRecord(c *AcmePropertyRecordContext)

	// EnterAcmePropertySequence is called when entering the acmePropertySequence production.
	EnterAcmePropertySequence(c *AcmePropertySequenceContext)

	// EnterAcmePropertyTypeRecord is called when entering the acmePropertyTypeRecord production.
	EnterAcmePropertyTypeRecord(c *AcmePropertyTypeRecordContext)

	// EnterAcmePropertyTypeSet is called when entering the acmePropertyTypeSet production.
	EnterAcmePropertyTypeSet(c *AcmePropertyTypeSetContext)

	// EnterAcmePropertyTypeSequence is called when entering the acmePropertyTypeSequence production.
	EnterAcmePropertyTypeSequence(c *AcmePropertyTypeSequenceContext)

	// EnterAcmePropertyTypeEnum is called when entering the acmePropertyTypeEnum production.
	EnterAcmePropertyTypeEnum(c *AcmePropertyTypeEnumContext)

	// EnterAcmePropertyRecordFieldDescription is called when entering the acmePropertyRecordFieldDescription production.
	EnterAcmePropertyRecordFieldDescription(c *AcmePropertyRecordFieldDescriptionContext)

	// EnterAcmePropertyTypeRef is called when entering the acmePropertyTypeRef production.
	EnterAcmePropertyTypeRef(c *AcmePropertyTypeRefContext)

	// EnterAcmePropertyTypeStructure is called when entering the acmePropertyTypeStructure production.
	EnterAcmePropertyTypeStructure(c *AcmePropertyTypeStructureContext)

	// EnterAcmePropertyTypeDeclaration is called when entering the acmePropertyTypeDeclaration production.
	EnterAcmePropertyTypeDeclaration(c *AcmePropertyTypeDeclarationContext)

	// EnterAcmePropertyBlockEntry is called when entering the acmePropertyBlockEntry production.
	EnterAcmePropertyBlockEntry(c *AcmePropertyBlockEntryContext)

	// EnterAcmePropertyBlock is called when entering the acmePropertyBlock production.
	EnterAcmePropertyBlock(c *AcmePropertyBlockContext)

	// EnterAcmePropertyTypeInt is called when entering the acmePropertyTypeInt production.
	EnterAcmePropertyTypeInt(c *AcmePropertyTypeIntContext)

	// EnterAcmePropertyTypeAny is called when entering the acmePropertyTypeAny production.
	EnterAcmePropertyTypeAny(c *AcmePropertyTypeAnyContext)

	// EnterAcmePropertyTypeFloat is called when entering the acmePropertyTypeFloat production.
	EnterAcmePropertyTypeFloat(c *AcmePropertyTypeFloatContext)

	// EnterAcmePropertyTypeDouble is called when entering the acmePropertyTypeDouble production.
	EnterAcmePropertyTypeDouble(c *AcmePropertyTypeDoubleContext)

	// EnterAcmePropertyTypeString is called when entering the acmePropertyTypeString production.
	EnterAcmePropertyTypeString(c *AcmePropertyTypeStringContext)

	// EnterAcmePropertyTypeBoolean is called when entering the acmePropertyTypeBoolean production.
	EnterAcmePropertyTypeBoolean(c *AcmePropertyTypeBooleanContext)

	// EnterAcmeViewDeclaration is called when entering the acmeViewDeclaration production.
	EnterAcmeViewDeclaration(c *AcmeViewDeclarationContext)

	// EnterAcmeViewTypeDeclaration is called when entering the acmeViewTypeDeclaration production.
	EnterAcmeViewTypeDeclaration(c *AcmeViewTypeDeclarationContext)

	// EnterAcmeViewBody is called when entering the acmeViewBody production.
	EnterAcmeViewBody(c *AcmeViewBodyContext)

	// EnterDesignRule is called when entering the designRule production.
	EnterDesignRule(c *DesignRuleContext)

	// EnterAcmeDesignAnalysisDeclaration is called when entering the acmeDesignAnalysisDeclaration production.
	EnterAcmeDesignAnalysisDeclaration(c *AcmeDesignAnalysisDeclarationContext)

	// EnterFormalParam is called when entering the formalParam production.
	EnterFormalParam(c *FormalParamContext)

	// EnterTerminatedDesignRuleExpression is called when entering the terminatedDesignRuleExpression production.
	EnterTerminatedDesignRuleExpression(c *TerminatedDesignRuleExpressionContext)

	// EnterDesignRuleExpression is called when entering the designRuleExpression production.
	EnterDesignRuleExpression(c *DesignRuleExpressionContext)

	// EnterASTDRImpliesExpression is called when entering the aSTDRImpliesExpression production.
	EnterASTDRImpliesExpression(c *ASTDRImpliesExpressionContext)

	// EnterDRIffExpression is called when entering the dRIffExpression production.
	EnterDRIffExpression(c *DRIffExpressionContext)

	// EnterDRAndExpression is called when entering the dRAndExpression production.
	EnterDRAndExpression(c *DRAndExpressionContext)

	// EnterDRNegateExpression is called when entering the dRNegateExpression production.
	EnterDRNegateExpression(c *DRNegateExpressionContext)

	// EnterDREqualityExpression is called when entering the dREqualityExpression production.
	EnterDREqualityExpression(c *DREqualityExpressionContext)

	// EnterDRRelationalExpression is called when entering the dRRelationalExpression production.
	EnterDRRelationalExpression(c *DRRelationalExpressionContext)

	// EnterDRAdditiveExpression is called when entering the dRAdditiveExpression production.
	EnterDRAdditiveExpression(c *DRAdditiveExpressionContext)

	// EnterDRMultiplicativeExpression is called when entering the dRMultiplicativeExpression production.
	EnterDRMultiplicativeExpression(c *DRMultiplicativeExpressionContext)

	// EnterDRNegativeExpression is called when entering the dRNegativeExpression production.
	EnterDRNegativeExpression(c *DRNegativeExpressionContext)

	// EnterPrimitiveExpression is called when entering the primitiveExpression production.
	EnterPrimitiveExpression(c *PrimitiveExpressionContext)

	// EnterParentheticalExpression is called when entering the parentheticalExpression production.
	EnterParentheticalExpression(c *ParentheticalExpressionContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterSetReference is called when entering the setReference production.
	EnterSetReference(c *SetReferenceContext)

	// EnterActualParams is called when entering the actualParams production.
	EnterActualParams(c *ActualParamsContext)

	// EnterLiteralConstant is called when entering the literalConstant production.
	EnterLiteralConstant(c *LiteralConstantContext)

	// EnterQuantifiedExpression is called when entering the quantifiedExpression production.
	EnterQuantifiedExpression(c *QuantifiedExpressionContext)

	// EnterDistinctVariableSetDeclaration is called when entering the distinctVariableSetDeclaration production.
	EnterDistinctVariableSetDeclaration(c *DistinctVariableSetDeclarationContext)

	// EnterVariableSetDeclaration is called when entering the variableSetDeclaration production.
	EnterVariableSetDeclaration(c *VariableSetDeclarationContext)

	// EnterSequenceExpression is called when entering the sequenceExpression production.
	EnterSequenceExpression(c *SequenceExpressionContext)

	// EnterSetExpression is called when entering the setExpression production.
	EnterSetExpression(c *SetExpressionContext)

	// EnterPathExpression is called when entering the pathExpression production.
	EnterPathExpression(c *PathExpressionContext)

	// EnterPathExpressionContinuation is called when entering the pathExpressionContinuation production.
	EnterPathExpressionContinuation(c *PathExpressionContinuationContext)

	// EnterLiteralSet is called when entering the literalSet production.
	EnterLiteralSet(c *LiteralSetContext)

	// EnterLiteralSequence is called when entering the literalSequence production.
	EnterLiteralSequence(c *LiteralSequenceContext)

	// EnterLiteralRecordEntry is called when entering the literalRecordEntry production.
	EnterLiteralRecordEntry(c *LiteralRecordEntryContext)

	// EnterLiteralRecord is called when entering the literalRecord production.
	EnterLiteralRecord(c *LiteralRecordContext)

	// EnterSetConstructor is called when entering the setConstructor production.
	EnterSetConstructor(c *SetConstructorContext)

	// EnterAcmeTypeRef is called when entering the acmeTypeRef production.
	EnterAcmeTypeRef(c *AcmeTypeRefContext)

	// ExitAcmeCompUnit is called when exiting the acmeCompUnit production.
	ExitAcmeCompUnit(c *AcmeCompUnitContext)

	// ExitAcmeImportDeclaration is called when exiting the acmeImportDeclaration production.
	ExitAcmeImportDeclaration(c *AcmeImportDeclarationContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitCodeLiteral is called when exiting the codeLiteral production.
	ExitCodeLiteral(c *CodeLiteralContext)

	// ExitAcmeFamilyDeclaration is called when exiting the acmeFamilyDeclaration production.
	ExitAcmeFamilyDeclaration(c *AcmeFamilyDeclarationContext)

	// ExitAcmeFamilyBody is called when exiting the acmeFamilyBody production.
	ExitAcmeFamilyBody(c *AcmeFamilyBodyContext)

	// ExitAcmeSystemDeclaration is called when exiting the acmeSystemDeclaration production.
	ExitAcmeSystemDeclaration(c *AcmeSystemDeclarationContext)

	// ExitAcmeSystemBody is called when exiting the acmeSystemBody production.
	ExitAcmeSystemBody(c *AcmeSystemBodyContext)

	// ExitAcmeDesignDeclaration is called when exiting the acmeDesignDeclaration production.
	ExitAcmeDesignDeclaration(c *AcmeDesignDeclarationContext)

	// ExitAcmeComponentTypeRef is called when exiting the acmeComponentTypeRef production.
	ExitAcmeComponentTypeRef(c *AcmeComponentTypeRefContext)

	// ExitAcmeComponentInstantiatedTypeRef is called when exiting the acmeComponentInstantiatedTypeRef production.
	ExitAcmeComponentInstantiatedTypeRef(c *AcmeComponentInstantiatedTypeRefContext)

	// ExitAcmeConnectorTypeRef is called when exiting the acmeConnectorTypeRef production.
	ExitAcmeConnectorTypeRef(c *AcmeConnectorTypeRefContext)

	// ExitAcmeConnectorInstantiatedTypeRef is called when exiting the acmeConnectorInstantiatedTypeRef production.
	ExitAcmeConnectorInstantiatedTypeRef(c *AcmeConnectorInstantiatedTypeRefContext)

	// ExitAcmePortTypeRef is called when exiting the acmePortTypeRef production.
	ExitAcmePortTypeRef(c *AcmePortTypeRefContext)

	// ExitAcmePortInstantiatedTypeRef is called when exiting the acmePortInstantiatedTypeRef production.
	ExitAcmePortInstantiatedTypeRef(c *AcmePortInstantiatedTypeRefContext)

	// ExitAcmeGroupTypeRef is called when exiting the acmeGroupTypeRef production.
	ExitAcmeGroupTypeRef(c *AcmeGroupTypeRefContext)

	// ExitAcmeGroupInstantiatedTypeRef is called when exiting the acmeGroupInstantiatedTypeRef production.
	ExitAcmeGroupInstantiatedTypeRef(c *AcmeGroupInstantiatedTypeRefContext)

	// ExitAcmeRoleTypeRef is called when exiting the acmeRoleTypeRef production.
	ExitAcmeRoleTypeRef(c *AcmeRoleTypeRefContext)

	// ExitAcmeRoleInstantiatedTypeRef is called when exiting the acmeRoleInstantiatedTypeRef production.
	ExitAcmeRoleInstantiatedTypeRef(c *AcmeRoleInstantiatedTypeRefContext)

	// ExitAcmeViewTypeRef is called when exiting the acmeViewTypeRef production.
	ExitAcmeViewTypeRef(c *AcmeViewTypeRefContext)

	// ExitAcmeViewInstantiatedTypeRef is called when exiting the acmeViewInstantiatedTypeRef production.
	ExitAcmeViewInstantiatedTypeRef(c *AcmeViewInstantiatedTypeRefContext)

	// ExitAcmeFamilyRef is called when exiting the acmeFamilyRef production.
	ExitAcmeFamilyRef(c *AcmeFamilyRefContext)

	// ExitAcmeFamilyInstantiationRef is called when exiting the acmeFamilyInstantiationRef production.
	ExitAcmeFamilyInstantiationRef(c *AcmeFamilyInstantiationRefContext)

	// ExitAcmeElementTypeRef is called when exiting the acmeElementTypeRef production.
	ExitAcmeElementTypeRef(c *AcmeElementTypeRefContext)

	// ExitAcmePropertyTypeDeclarationRef is called when exiting the acmePropertyTypeDeclarationRef production.
	ExitAcmePropertyTypeDeclarationRef(c *AcmePropertyTypeDeclarationRefContext)

	// ExitAcmeInstanceRef is called when exiting the acmeInstanceRef production.
	ExitAcmeInstanceRef(c *AcmeInstanceRefContext)

	// ExitAcmeGenericElementTypeDeclaration is called when exiting the acmeGenericElementTypeDeclaration production.
	ExitAcmeGenericElementTypeDeclaration(c *AcmeGenericElementTypeDeclarationContext)

	// ExitAcmeGenericElementBody is called when exiting the acmeGenericElementBody production.
	ExitAcmeGenericElementBody(c *AcmeGenericElementBodyContext)

	// ExitAcmeGroupTypeDeclaration is called when exiting the acmeGroupTypeDeclaration production.
	ExitAcmeGroupTypeDeclaration(c *AcmeGroupTypeDeclarationContext)

	// ExitAcmeGroupDeclaration is called when exiting the acmeGroupDeclaration production.
	ExitAcmeGroupDeclaration(c *AcmeGroupDeclarationContext)

	// ExitAcmeGroupBody is called when exiting the acmeGroupBody production.
	ExitAcmeGroupBody(c *AcmeGroupBodyContext)

	// ExitAcmeMembersBlock is called when exiting the acmeMembersBlock production.
	ExitAcmeMembersBlock(c *AcmeMembersBlockContext)

	// ExitAcmePortTypeDeclaration is called when exiting the acmePortTypeDeclaration production.
	ExitAcmePortTypeDeclaration(c *AcmePortTypeDeclarationContext)

	// ExitAcmePortDeclaration is called when exiting the acmePortDeclaration production.
	ExitAcmePortDeclaration(c *AcmePortDeclarationContext)

	// ExitAcmePortBody is called when exiting the acmePortBody production.
	ExitAcmePortBody(c *AcmePortBodyContext)

	// ExitAcmeRoleTypeDeclaration is called when exiting the acmeRoleTypeDeclaration production.
	ExitAcmeRoleTypeDeclaration(c *AcmeRoleTypeDeclarationContext)

	// ExitAcmeRoleDeclaration is called when exiting the acmeRoleDeclaration production.
	ExitAcmeRoleDeclaration(c *AcmeRoleDeclarationContext)

	// ExitAcmeRoleBody is called when exiting the acmeRoleBody production.
	ExitAcmeRoleBody(c *AcmeRoleBodyContext)

	// ExitAcmeComponentTypeDeclaration is called when exiting the acmeComponentTypeDeclaration production.
	ExitAcmeComponentTypeDeclaration(c *AcmeComponentTypeDeclarationContext)

	// ExitAcmeComponentDeclaration is called when exiting the acmeComponentDeclaration production.
	ExitAcmeComponentDeclaration(c *AcmeComponentDeclarationContext)

	// ExitAcmeComponentBody is called when exiting the acmeComponentBody production.
	ExitAcmeComponentBody(c *AcmeComponentBodyContext)

	// ExitAcmeConnectorTypeDeclaration is called when exiting the acmeConnectorTypeDeclaration production.
	ExitAcmeConnectorTypeDeclaration(c *AcmeConnectorTypeDeclarationContext)

	// ExitAcmeConnectorDeclaration is called when exiting the acmeConnectorDeclaration production.
	ExitAcmeConnectorDeclaration(c *AcmeConnectorDeclarationContext)

	// ExitAcmeConnectorBody is called when exiting the acmeConnectorBody production.
	ExitAcmeConnectorBody(c *AcmeConnectorBodyContext)

	// ExitAcmeRepresentationDeclaration is called when exiting the acmeRepresentationDeclaration production.
	ExitAcmeRepresentationDeclaration(c *AcmeRepresentationDeclarationContext)

	// ExitAcmeBindingsMapDeclaration is called when exiting the acmeBindingsMapDeclaration production.
	ExitAcmeBindingsMapDeclaration(c *AcmeBindingsMapDeclarationContext)

	// ExitAcmeBindingDeclaration is called when exiting the acmeBindingDeclaration production.
	ExitAcmeBindingDeclaration(c *AcmeBindingDeclarationContext)

	// ExitAcmeAttachmentDeclaration is called when exiting the acmeAttachmentDeclaration production.
	ExitAcmeAttachmentDeclaration(c *AcmeAttachmentDeclarationContext)

	// ExitAcmePropertyDeclaration is called when exiting the acmePropertyDeclaration production.
	ExitAcmePropertyDeclaration(c *AcmePropertyDeclarationContext)

	// ExitAcmePropertyValueDeclaration is called when exiting the acmePropertyValueDeclaration production.
	ExitAcmePropertyValueDeclaration(c *AcmePropertyValueDeclarationContext)

	// ExitEnumidentifier is called when exiting the enumidentifier production.
	ExitEnumidentifier(c *EnumidentifierContext)

	// ExitAcmePropertyElement is called when exiting the acmePropertyElement production.
	ExitAcmePropertyElement(c *AcmePropertyElementContext)

	// ExitAcmePropertyCompoundElement is called when exiting the acmePropertyCompoundElement production.
	ExitAcmePropertyCompoundElement(c *AcmePropertyCompoundElementContext)

	// ExitAcmePropertySet is called when exiting the acmePropertySet production.
	ExitAcmePropertySet(c *AcmePropertySetContext)

	// ExitAcmePropertyRecordEntry is called when exiting the acmePropertyRecordEntry production.
	ExitAcmePropertyRecordEntry(c *AcmePropertyRecordEntryContext)

	// ExitAcmePropertyRecord is called when exiting the acmePropertyRecord production.
	ExitAcmePropertyRecord(c *AcmePropertyRecordContext)

	// ExitAcmePropertySequence is called when exiting the acmePropertySequence production.
	ExitAcmePropertySequence(c *AcmePropertySequenceContext)

	// ExitAcmePropertyTypeRecord is called when exiting the acmePropertyTypeRecord production.
	ExitAcmePropertyTypeRecord(c *AcmePropertyTypeRecordContext)

	// ExitAcmePropertyTypeSet is called when exiting the acmePropertyTypeSet production.
	ExitAcmePropertyTypeSet(c *AcmePropertyTypeSetContext)

	// ExitAcmePropertyTypeSequence is called when exiting the acmePropertyTypeSequence production.
	ExitAcmePropertyTypeSequence(c *AcmePropertyTypeSequenceContext)

	// ExitAcmePropertyTypeEnum is called when exiting the acmePropertyTypeEnum production.
	ExitAcmePropertyTypeEnum(c *AcmePropertyTypeEnumContext)

	// ExitAcmePropertyRecordFieldDescription is called when exiting the acmePropertyRecordFieldDescription production.
	ExitAcmePropertyRecordFieldDescription(c *AcmePropertyRecordFieldDescriptionContext)

	// ExitAcmePropertyTypeRef is called when exiting the acmePropertyTypeRef production.
	ExitAcmePropertyTypeRef(c *AcmePropertyTypeRefContext)

	// ExitAcmePropertyTypeStructure is called when exiting the acmePropertyTypeStructure production.
	ExitAcmePropertyTypeStructure(c *AcmePropertyTypeStructureContext)

	// ExitAcmePropertyTypeDeclaration is called when exiting the acmePropertyTypeDeclaration production.
	ExitAcmePropertyTypeDeclaration(c *AcmePropertyTypeDeclarationContext)

	// ExitAcmePropertyBlockEntry is called when exiting the acmePropertyBlockEntry production.
	ExitAcmePropertyBlockEntry(c *AcmePropertyBlockEntryContext)

	// ExitAcmePropertyBlock is called when exiting the acmePropertyBlock production.
	ExitAcmePropertyBlock(c *AcmePropertyBlockContext)

	// ExitAcmePropertyTypeInt is called when exiting the acmePropertyTypeInt production.
	ExitAcmePropertyTypeInt(c *AcmePropertyTypeIntContext)

	// ExitAcmePropertyTypeAny is called when exiting the acmePropertyTypeAny production.
	ExitAcmePropertyTypeAny(c *AcmePropertyTypeAnyContext)

	// ExitAcmePropertyTypeFloat is called when exiting the acmePropertyTypeFloat production.
	ExitAcmePropertyTypeFloat(c *AcmePropertyTypeFloatContext)

	// ExitAcmePropertyTypeDouble is called when exiting the acmePropertyTypeDouble production.
	ExitAcmePropertyTypeDouble(c *AcmePropertyTypeDoubleContext)

	// ExitAcmePropertyTypeString is called when exiting the acmePropertyTypeString production.
	ExitAcmePropertyTypeString(c *AcmePropertyTypeStringContext)

	// ExitAcmePropertyTypeBoolean is called when exiting the acmePropertyTypeBoolean production.
	ExitAcmePropertyTypeBoolean(c *AcmePropertyTypeBooleanContext)

	// ExitAcmeViewDeclaration is called when exiting the acmeViewDeclaration production.
	ExitAcmeViewDeclaration(c *AcmeViewDeclarationContext)

	// ExitAcmeViewTypeDeclaration is called when exiting the acmeViewTypeDeclaration production.
	ExitAcmeViewTypeDeclaration(c *AcmeViewTypeDeclarationContext)

	// ExitAcmeViewBody is called when exiting the acmeViewBody production.
	ExitAcmeViewBody(c *AcmeViewBodyContext)

	// ExitDesignRule is called when exiting the designRule production.
	ExitDesignRule(c *DesignRuleContext)

	// ExitAcmeDesignAnalysisDeclaration is called when exiting the acmeDesignAnalysisDeclaration production.
	ExitAcmeDesignAnalysisDeclaration(c *AcmeDesignAnalysisDeclarationContext)

	// ExitFormalParam is called when exiting the formalParam production.
	ExitFormalParam(c *FormalParamContext)

	// ExitTerminatedDesignRuleExpression is called when exiting the terminatedDesignRuleExpression production.
	ExitTerminatedDesignRuleExpression(c *TerminatedDesignRuleExpressionContext)

	// ExitDesignRuleExpression is called when exiting the designRuleExpression production.
	ExitDesignRuleExpression(c *DesignRuleExpressionContext)

	// ExitASTDRImpliesExpression is called when exiting the aSTDRImpliesExpression production.
	ExitASTDRImpliesExpression(c *ASTDRImpliesExpressionContext)

	// ExitDRIffExpression is called when exiting the dRIffExpression production.
	ExitDRIffExpression(c *DRIffExpressionContext)

	// ExitDRAndExpression is called when exiting the dRAndExpression production.
	ExitDRAndExpression(c *DRAndExpressionContext)

	// ExitDRNegateExpression is called when exiting the dRNegateExpression production.
	ExitDRNegateExpression(c *DRNegateExpressionContext)

	// ExitDREqualityExpression is called when exiting the dREqualityExpression production.
	ExitDREqualityExpression(c *DREqualityExpressionContext)

	// ExitDRRelationalExpression is called when exiting the dRRelationalExpression production.
	ExitDRRelationalExpression(c *DRRelationalExpressionContext)

	// ExitDRAdditiveExpression is called when exiting the dRAdditiveExpression production.
	ExitDRAdditiveExpression(c *DRAdditiveExpressionContext)

	// ExitDRMultiplicativeExpression is called when exiting the dRMultiplicativeExpression production.
	ExitDRMultiplicativeExpression(c *DRMultiplicativeExpressionContext)

	// ExitDRNegativeExpression is called when exiting the dRNegativeExpression production.
	ExitDRNegativeExpression(c *DRNegativeExpressionContext)

	// ExitPrimitiveExpression is called when exiting the primitiveExpression production.
	ExitPrimitiveExpression(c *PrimitiveExpressionContext)

	// ExitParentheticalExpression is called when exiting the parentheticalExpression production.
	ExitParentheticalExpression(c *ParentheticalExpressionContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitSetReference is called when exiting the setReference production.
	ExitSetReference(c *SetReferenceContext)

	// ExitActualParams is called when exiting the actualParams production.
	ExitActualParams(c *ActualParamsContext)

	// ExitLiteralConstant is called when exiting the literalConstant production.
	ExitLiteralConstant(c *LiteralConstantContext)

	// ExitQuantifiedExpression is called when exiting the quantifiedExpression production.
	ExitQuantifiedExpression(c *QuantifiedExpressionContext)

	// ExitDistinctVariableSetDeclaration is called when exiting the distinctVariableSetDeclaration production.
	ExitDistinctVariableSetDeclaration(c *DistinctVariableSetDeclarationContext)

	// ExitVariableSetDeclaration is called when exiting the variableSetDeclaration production.
	ExitVariableSetDeclaration(c *VariableSetDeclarationContext)

	// ExitSequenceExpression is called when exiting the sequenceExpression production.
	ExitSequenceExpression(c *SequenceExpressionContext)

	// ExitSetExpression is called when exiting the setExpression production.
	ExitSetExpression(c *SetExpressionContext)

	// ExitPathExpression is called when exiting the pathExpression production.
	ExitPathExpression(c *PathExpressionContext)

	// ExitPathExpressionContinuation is called when exiting the pathExpressionContinuation production.
	ExitPathExpressionContinuation(c *PathExpressionContinuationContext)

	// ExitLiteralSet is called when exiting the literalSet production.
	ExitLiteralSet(c *LiteralSetContext)

	// ExitLiteralSequence is called when exiting the literalSequence production.
	ExitLiteralSequence(c *LiteralSequenceContext)

	// ExitLiteralRecordEntry is called when exiting the literalRecordEntry production.
	ExitLiteralRecordEntry(c *LiteralRecordEntryContext)

	// ExitLiteralRecord is called when exiting the literalRecord production.
	ExitLiteralRecord(c *LiteralRecordContext)

	// ExitSetConstructor is called when exiting the setConstructor production.
	ExitSetConstructor(c *SetConstructorContext)

	// ExitAcmeTypeRef is called when exiting the acmeTypeRef production.
	ExitAcmeTypeRef(c *AcmeTypeRefContext)
}
