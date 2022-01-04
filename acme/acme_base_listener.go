// Code generated from acme.g4 by ANTLR 4.9.3. DO NOT EDIT.

package acme // acme
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseacmeListener is a complete listener for a parse tree produced by acmeParser.
type BaseacmeListener struct{}

var _ acmeListener = &BaseacmeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseacmeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseacmeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseacmeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseacmeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAcmeCompUnit is called when production acmeCompUnit is entered.
func (s *BaseacmeListener) EnterAcmeCompUnit(ctx *AcmeCompUnitContext) {}

// ExitAcmeCompUnit is called when production acmeCompUnit is exited.
func (s *BaseacmeListener) ExitAcmeCompUnit(ctx *AcmeCompUnitContext) {}

// EnterAcmeImportDeclaration is called when production acmeImportDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeImportDeclaration(ctx *AcmeImportDeclarationContext) {}

// ExitAcmeImportDeclaration is called when production acmeImportDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeImportDeclaration(ctx *AcmeImportDeclarationContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseacmeListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseacmeListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseacmeListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseacmeListener) ExitFilename(ctx *FilenameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseacmeListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseacmeListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterCodeLiteral is called when production codeLiteral is entered.
func (s *BaseacmeListener) EnterCodeLiteral(ctx *CodeLiteralContext) {}

// ExitCodeLiteral is called when production codeLiteral is exited.
func (s *BaseacmeListener) ExitCodeLiteral(ctx *CodeLiteralContext) {}

// EnterAcmeFamilyDeclaration is called when production acmeFamilyDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeFamilyDeclaration(ctx *AcmeFamilyDeclarationContext) {}

// ExitAcmeFamilyDeclaration is called when production acmeFamilyDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeFamilyDeclaration(ctx *AcmeFamilyDeclarationContext) {}

// EnterAcmeFamilyBody is called when production acmeFamilyBody is entered.
func (s *BaseacmeListener) EnterAcmeFamilyBody(ctx *AcmeFamilyBodyContext) {}

// ExitAcmeFamilyBody is called when production acmeFamilyBody is exited.
func (s *BaseacmeListener) ExitAcmeFamilyBody(ctx *AcmeFamilyBodyContext) {}

// EnterAcmeSystemDeclaration is called when production acmeSystemDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeSystemDeclaration(ctx *AcmeSystemDeclarationContext) {}

// ExitAcmeSystemDeclaration is called when production acmeSystemDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeSystemDeclaration(ctx *AcmeSystemDeclarationContext) {}

// EnterAcmeSystemBody is called when production acmeSystemBody is entered.
func (s *BaseacmeListener) EnterAcmeSystemBody(ctx *AcmeSystemBodyContext) {}

// ExitAcmeSystemBody is called when production acmeSystemBody is exited.
func (s *BaseacmeListener) ExitAcmeSystemBody(ctx *AcmeSystemBodyContext) {}

// EnterAcmeDesignDeclaration is called when production acmeDesignDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeDesignDeclaration(ctx *AcmeDesignDeclarationContext) {}

// ExitAcmeDesignDeclaration is called when production acmeDesignDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeDesignDeclaration(ctx *AcmeDesignDeclarationContext) {}

// EnterAcmeComponentTypeRef is called when production acmeComponentTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeComponentTypeRef(ctx *AcmeComponentTypeRefContext) {}

// ExitAcmeComponentTypeRef is called when production acmeComponentTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeComponentTypeRef(ctx *AcmeComponentTypeRefContext) {}

// EnterAcmeComponentInstantiatedTypeRef is called when production acmeComponentInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeComponentInstantiatedTypeRef(ctx *AcmeComponentInstantiatedTypeRefContext) {
}

// ExitAcmeComponentInstantiatedTypeRef is called when production acmeComponentInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeComponentInstantiatedTypeRef(ctx *AcmeComponentInstantiatedTypeRefContext) {
}

// EnterAcmeConnectorTypeRef is called when production acmeConnectorTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeConnectorTypeRef(ctx *AcmeConnectorTypeRefContext) {}

// ExitAcmeConnectorTypeRef is called when production acmeConnectorTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeConnectorTypeRef(ctx *AcmeConnectorTypeRefContext) {}

// EnterAcmeConnectorInstantiatedTypeRef is called when production acmeConnectorInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeConnectorInstantiatedTypeRef(ctx *AcmeConnectorInstantiatedTypeRefContext) {
}

// ExitAcmeConnectorInstantiatedTypeRef is called when production acmeConnectorInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeConnectorInstantiatedTypeRef(ctx *AcmeConnectorInstantiatedTypeRefContext) {
}

// EnterAcmePortTypeRef is called when production acmePortTypeRef is entered.
func (s *BaseacmeListener) EnterAcmePortTypeRef(ctx *AcmePortTypeRefContext) {}

// ExitAcmePortTypeRef is called when production acmePortTypeRef is exited.
func (s *BaseacmeListener) ExitAcmePortTypeRef(ctx *AcmePortTypeRefContext) {}

// EnterAcmePortInstantiatedTypeRef is called when production acmePortInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmePortInstantiatedTypeRef(ctx *AcmePortInstantiatedTypeRefContext) {
}

// ExitAcmePortInstantiatedTypeRef is called when production acmePortInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmePortInstantiatedTypeRef(ctx *AcmePortInstantiatedTypeRefContext) {}

// EnterAcmeGroupTypeRef is called when production acmeGroupTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeGroupTypeRef(ctx *AcmeGroupTypeRefContext) {}

// ExitAcmeGroupTypeRef is called when production acmeGroupTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeGroupTypeRef(ctx *AcmeGroupTypeRefContext) {}

// EnterAcmeGroupInstantiatedTypeRef is called when production acmeGroupInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeGroupInstantiatedTypeRef(ctx *AcmeGroupInstantiatedTypeRefContext) {
}

// ExitAcmeGroupInstantiatedTypeRef is called when production acmeGroupInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeGroupInstantiatedTypeRef(ctx *AcmeGroupInstantiatedTypeRefContext) {
}

// EnterAcmeRoleTypeRef is called when production acmeRoleTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeRoleTypeRef(ctx *AcmeRoleTypeRefContext) {}

// ExitAcmeRoleTypeRef is called when production acmeRoleTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeRoleTypeRef(ctx *AcmeRoleTypeRefContext) {}

// EnterAcmeRoleInstantiatedTypeRef is called when production acmeRoleInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeRoleInstantiatedTypeRef(ctx *AcmeRoleInstantiatedTypeRefContext) {
}

// ExitAcmeRoleInstantiatedTypeRef is called when production acmeRoleInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeRoleInstantiatedTypeRef(ctx *AcmeRoleInstantiatedTypeRefContext) {}

// EnterAcmeViewTypeRef is called when production acmeViewTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeViewTypeRef(ctx *AcmeViewTypeRefContext) {}

// ExitAcmeViewTypeRef is called when production acmeViewTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeViewTypeRef(ctx *AcmeViewTypeRefContext) {}

// EnterAcmeViewInstantiatedTypeRef is called when production acmeViewInstantiatedTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeViewInstantiatedTypeRef(ctx *AcmeViewInstantiatedTypeRefContext) {
}

// ExitAcmeViewInstantiatedTypeRef is called when production acmeViewInstantiatedTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeViewInstantiatedTypeRef(ctx *AcmeViewInstantiatedTypeRefContext) {}

// EnterAcmeFamilyRef is called when production acmeFamilyRef is entered.
func (s *BaseacmeListener) EnterAcmeFamilyRef(ctx *AcmeFamilyRefContext) {}

// ExitAcmeFamilyRef is called when production acmeFamilyRef is exited.
func (s *BaseacmeListener) ExitAcmeFamilyRef(ctx *AcmeFamilyRefContext) {}

// EnterAcmeFamilyInstantiationRef is called when production acmeFamilyInstantiationRef is entered.
func (s *BaseacmeListener) EnterAcmeFamilyInstantiationRef(ctx *AcmeFamilyInstantiationRefContext) {}

// ExitAcmeFamilyInstantiationRef is called when production acmeFamilyInstantiationRef is exited.
func (s *BaseacmeListener) ExitAcmeFamilyInstantiationRef(ctx *AcmeFamilyInstantiationRefContext) {}

// EnterAcmeElementTypeRef is called when production acmeElementTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeElementTypeRef(ctx *AcmeElementTypeRefContext) {}

// ExitAcmeElementTypeRef is called when production acmeElementTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeElementTypeRef(ctx *AcmeElementTypeRefContext) {}

// EnterAcmePropertyTypeDeclarationRef is called when production acmePropertyTypeDeclarationRef is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeDeclarationRef(ctx *AcmePropertyTypeDeclarationRefContext) {
}

// ExitAcmePropertyTypeDeclarationRef is called when production acmePropertyTypeDeclarationRef is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeDeclarationRef(ctx *AcmePropertyTypeDeclarationRefContext) {
}

// EnterAcmeInstanceRef is called when production acmeInstanceRef is entered.
func (s *BaseacmeListener) EnterAcmeInstanceRef(ctx *AcmeInstanceRefContext) {}

// ExitAcmeInstanceRef is called when production acmeInstanceRef is exited.
func (s *BaseacmeListener) ExitAcmeInstanceRef(ctx *AcmeInstanceRefContext) {}

// EnterAcmeGenericElementTypeDeclaration is called when production acmeGenericElementTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeGenericElementTypeDeclaration(ctx *AcmeGenericElementTypeDeclarationContext) {
}

// ExitAcmeGenericElementTypeDeclaration is called when production acmeGenericElementTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeGenericElementTypeDeclaration(ctx *AcmeGenericElementTypeDeclarationContext) {
}

// EnterAcmeGenericElementBody is called when production acmeGenericElementBody is entered.
func (s *BaseacmeListener) EnterAcmeGenericElementBody(ctx *AcmeGenericElementBodyContext) {}

// ExitAcmeGenericElementBody is called when production acmeGenericElementBody is exited.
func (s *BaseacmeListener) ExitAcmeGenericElementBody(ctx *AcmeGenericElementBodyContext) {}

// EnterAcmeGroupTypeDeclaration is called when production acmeGroupTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeGroupTypeDeclaration(ctx *AcmeGroupTypeDeclarationContext) {}

// ExitAcmeGroupTypeDeclaration is called when production acmeGroupTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeGroupTypeDeclaration(ctx *AcmeGroupTypeDeclarationContext) {}

// EnterAcmeGroupDeclaration is called when production acmeGroupDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeGroupDeclaration(ctx *AcmeGroupDeclarationContext) {}

// ExitAcmeGroupDeclaration is called when production acmeGroupDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeGroupDeclaration(ctx *AcmeGroupDeclarationContext) {}

// EnterAcmeGroupBody is called when production acmeGroupBody is entered.
func (s *BaseacmeListener) EnterAcmeGroupBody(ctx *AcmeGroupBodyContext) {}

// ExitAcmeGroupBody is called when production acmeGroupBody is exited.
func (s *BaseacmeListener) ExitAcmeGroupBody(ctx *AcmeGroupBodyContext) {}

// EnterAcmeMembersBlock is called when production acmeMembersBlock is entered.
func (s *BaseacmeListener) EnterAcmeMembersBlock(ctx *AcmeMembersBlockContext) {}

// ExitAcmeMembersBlock is called when production acmeMembersBlock is exited.
func (s *BaseacmeListener) ExitAcmeMembersBlock(ctx *AcmeMembersBlockContext) {}

// EnterAcmePortTypeDeclaration is called when production acmePortTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmePortTypeDeclaration(ctx *AcmePortTypeDeclarationContext) {}

// ExitAcmePortTypeDeclaration is called when production acmePortTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmePortTypeDeclaration(ctx *AcmePortTypeDeclarationContext) {}

// EnterAcmePortDeclaration is called when production acmePortDeclaration is entered.
func (s *BaseacmeListener) EnterAcmePortDeclaration(ctx *AcmePortDeclarationContext) {}

// ExitAcmePortDeclaration is called when production acmePortDeclaration is exited.
func (s *BaseacmeListener) ExitAcmePortDeclaration(ctx *AcmePortDeclarationContext) {}

// EnterAcmePortBody is called when production acmePortBody is entered.
func (s *BaseacmeListener) EnterAcmePortBody(ctx *AcmePortBodyContext) {}

// ExitAcmePortBody is called when production acmePortBody is exited.
func (s *BaseacmeListener) ExitAcmePortBody(ctx *AcmePortBodyContext) {}

// EnterAcmeRoleTypeDeclaration is called when production acmeRoleTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeRoleTypeDeclaration(ctx *AcmeRoleTypeDeclarationContext) {}

// ExitAcmeRoleTypeDeclaration is called when production acmeRoleTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeRoleTypeDeclaration(ctx *AcmeRoleTypeDeclarationContext) {}

// EnterAcmeRoleDeclaration is called when production acmeRoleDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeRoleDeclaration(ctx *AcmeRoleDeclarationContext) {}

// ExitAcmeRoleDeclaration is called when production acmeRoleDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeRoleDeclaration(ctx *AcmeRoleDeclarationContext) {}

// EnterAcmeRoleBody is called when production acmeRoleBody is entered.
func (s *BaseacmeListener) EnterAcmeRoleBody(ctx *AcmeRoleBodyContext) {}

// ExitAcmeRoleBody is called when production acmeRoleBody is exited.
func (s *BaseacmeListener) ExitAcmeRoleBody(ctx *AcmeRoleBodyContext) {}

// EnterAcmeComponentTypeDeclaration is called when production acmeComponentTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeComponentTypeDeclaration(ctx *AcmeComponentTypeDeclarationContext) {
}

// ExitAcmeComponentTypeDeclaration is called when production acmeComponentTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeComponentTypeDeclaration(ctx *AcmeComponentTypeDeclarationContext) {
}

// EnterAcmeComponentDeclaration is called when production acmeComponentDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeComponentDeclaration(ctx *AcmeComponentDeclarationContext) {}

// ExitAcmeComponentDeclaration is called when production acmeComponentDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeComponentDeclaration(ctx *AcmeComponentDeclarationContext) {}

// EnterAcmeComponentBody is called when production acmeComponentBody is entered.
func (s *BaseacmeListener) EnterAcmeComponentBody(ctx *AcmeComponentBodyContext) {}

// ExitAcmeComponentBody is called when production acmeComponentBody is exited.
func (s *BaseacmeListener) ExitAcmeComponentBody(ctx *AcmeComponentBodyContext) {}

// EnterAcmeConnectorTypeDeclaration is called when production acmeConnectorTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeConnectorTypeDeclaration(ctx *AcmeConnectorTypeDeclarationContext) {
}

// ExitAcmeConnectorTypeDeclaration is called when production acmeConnectorTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeConnectorTypeDeclaration(ctx *AcmeConnectorTypeDeclarationContext) {
}

// EnterAcmeConnectorDeclaration is called when production acmeConnectorDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeConnectorDeclaration(ctx *AcmeConnectorDeclarationContext) {}

// ExitAcmeConnectorDeclaration is called when production acmeConnectorDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeConnectorDeclaration(ctx *AcmeConnectorDeclarationContext) {}

// EnterAcmeConnectorBody is called when production acmeConnectorBody is entered.
func (s *BaseacmeListener) EnterAcmeConnectorBody(ctx *AcmeConnectorBodyContext) {}

// ExitAcmeConnectorBody is called when production acmeConnectorBody is exited.
func (s *BaseacmeListener) ExitAcmeConnectorBody(ctx *AcmeConnectorBodyContext) {}

// EnterAcmeRepresentationDeclaration is called when production acmeRepresentationDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeRepresentationDeclaration(ctx *AcmeRepresentationDeclarationContext) {
}

// ExitAcmeRepresentationDeclaration is called when production acmeRepresentationDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeRepresentationDeclaration(ctx *AcmeRepresentationDeclarationContext) {
}

// EnterAcmeBindingsMapDeclaration is called when production acmeBindingsMapDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeBindingsMapDeclaration(ctx *AcmeBindingsMapDeclarationContext) {}

// ExitAcmeBindingsMapDeclaration is called when production acmeBindingsMapDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeBindingsMapDeclaration(ctx *AcmeBindingsMapDeclarationContext) {}

// EnterAcmeBindingDeclaration is called when production acmeBindingDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeBindingDeclaration(ctx *AcmeBindingDeclarationContext) {}

// ExitAcmeBindingDeclaration is called when production acmeBindingDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeBindingDeclaration(ctx *AcmeBindingDeclarationContext) {}

// EnterAcmeAttachmentDeclaration is called when production acmeAttachmentDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeAttachmentDeclaration(ctx *AcmeAttachmentDeclarationContext) {}

// ExitAcmeAttachmentDeclaration is called when production acmeAttachmentDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeAttachmentDeclaration(ctx *AcmeAttachmentDeclarationContext) {}

// EnterAcmePropertyDeclaration is called when production acmePropertyDeclaration is entered.
func (s *BaseacmeListener) EnterAcmePropertyDeclaration(ctx *AcmePropertyDeclarationContext) {}

// ExitAcmePropertyDeclaration is called when production acmePropertyDeclaration is exited.
func (s *BaseacmeListener) ExitAcmePropertyDeclaration(ctx *AcmePropertyDeclarationContext) {}

// EnterAcmePropertyValueDeclaration is called when production acmePropertyValueDeclaration is entered.
func (s *BaseacmeListener) EnterAcmePropertyValueDeclaration(ctx *AcmePropertyValueDeclarationContext) {
}

// ExitAcmePropertyValueDeclaration is called when production acmePropertyValueDeclaration is exited.
func (s *BaseacmeListener) ExitAcmePropertyValueDeclaration(ctx *AcmePropertyValueDeclarationContext) {
}

// EnterEnumidentifier is called when production enumidentifier is entered.
func (s *BaseacmeListener) EnterEnumidentifier(ctx *EnumidentifierContext) {}

// ExitEnumidentifier is called when production enumidentifier is exited.
func (s *BaseacmeListener) ExitEnumidentifier(ctx *EnumidentifierContext) {}

// EnterAcmePropertyElement is called when production acmePropertyElement is entered.
func (s *BaseacmeListener) EnterAcmePropertyElement(ctx *AcmePropertyElementContext) {}

// ExitAcmePropertyElement is called when production acmePropertyElement is exited.
func (s *BaseacmeListener) ExitAcmePropertyElement(ctx *AcmePropertyElementContext) {}

// EnterAcmePropertyCompoundElement is called when production acmePropertyCompoundElement is entered.
func (s *BaseacmeListener) EnterAcmePropertyCompoundElement(ctx *AcmePropertyCompoundElementContext) {
}

// ExitAcmePropertyCompoundElement is called when production acmePropertyCompoundElement is exited.
func (s *BaseacmeListener) ExitAcmePropertyCompoundElement(ctx *AcmePropertyCompoundElementContext) {}

// EnterAcmePropertySet is called when production acmePropertySet is entered.
func (s *BaseacmeListener) EnterAcmePropertySet(ctx *AcmePropertySetContext) {}

// ExitAcmePropertySet is called when production acmePropertySet is exited.
func (s *BaseacmeListener) ExitAcmePropertySet(ctx *AcmePropertySetContext) {}

// EnterAcmePropertyRecordEntry is called when production acmePropertyRecordEntry is entered.
func (s *BaseacmeListener) EnterAcmePropertyRecordEntry(ctx *AcmePropertyRecordEntryContext) {}

// ExitAcmePropertyRecordEntry is called when production acmePropertyRecordEntry is exited.
func (s *BaseacmeListener) ExitAcmePropertyRecordEntry(ctx *AcmePropertyRecordEntryContext) {}

// EnterAcmePropertyRecord is called when production acmePropertyRecord is entered.
func (s *BaseacmeListener) EnterAcmePropertyRecord(ctx *AcmePropertyRecordContext) {}

// ExitAcmePropertyRecord is called when production acmePropertyRecord is exited.
func (s *BaseacmeListener) ExitAcmePropertyRecord(ctx *AcmePropertyRecordContext) {}

// EnterAcmePropertySequence is called when production acmePropertySequence is entered.
func (s *BaseacmeListener) EnterAcmePropertySequence(ctx *AcmePropertySequenceContext) {}

// ExitAcmePropertySequence is called when production acmePropertySequence is exited.
func (s *BaseacmeListener) ExitAcmePropertySequence(ctx *AcmePropertySequenceContext) {}

// EnterAcmePropertyTypeRecord is called when production acmePropertyTypeRecord is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeRecord(ctx *AcmePropertyTypeRecordContext) {}

// ExitAcmePropertyTypeRecord is called when production acmePropertyTypeRecord is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeRecord(ctx *AcmePropertyTypeRecordContext) {}

// EnterAcmePropertyTypeSet is called when production acmePropertyTypeSet is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeSet(ctx *AcmePropertyTypeSetContext) {}

// ExitAcmePropertyTypeSet is called when production acmePropertyTypeSet is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeSet(ctx *AcmePropertyTypeSetContext) {}

// EnterAcmePropertyTypeSequence is called when production acmePropertyTypeSequence is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeSequence(ctx *AcmePropertyTypeSequenceContext) {}

// ExitAcmePropertyTypeSequence is called when production acmePropertyTypeSequence is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeSequence(ctx *AcmePropertyTypeSequenceContext) {}

// EnterAcmePropertyTypeEnum is called when production acmePropertyTypeEnum is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeEnum(ctx *AcmePropertyTypeEnumContext) {}

// ExitAcmePropertyTypeEnum is called when production acmePropertyTypeEnum is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeEnum(ctx *AcmePropertyTypeEnumContext) {}

// EnterAcmePropertyRecordFieldDescription is called when production acmePropertyRecordFieldDescription is entered.
func (s *BaseacmeListener) EnterAcmePropertyRecordFieldDescription(ctx *AcmePropertyRecordFieldDescriptionContext) {
}

// ExitAcmePropertyRecordFieldDescription is called when production acmePropertyRecordFieldDescription is exited.
func (s *BaseacmeListener) ExitAcmePropertyRecordFieldDescription(ctx *AcmePropertyRecordFieldDescriptionContext) {
}

// EnterAcmePropertyTypeRef is called when production acmePropertyTypeRef is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeRef(ctx *AcmePropertyTypeRefContext) {}

// ExitAcmePropertyTypeRef is called when production acmePropertyTypeRef is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeRef(ctx *AcmePropertyTypeRefContext) {}

// EnterAcmePropertyTypeStructure is called when production acmePropertyTypeStructure is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeStructure(ctx *AcmePropertyTypeStructureContext) {}

// ExitAcmePropertyTypeStructure is called when production acmePropertyTypeStructure is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeStructure(ctx *AcmePropertyTypeStructureContext) {}

// EnterAcmePropertyTypeDeclaration is called when production acmePropertyTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeDeclaration(ctx *AcmePropertyTypeDeclarationContext) {
}

// ExitAcmePropertyTypeDeclaration is called when production acmePropertyTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeDeclaration(ctx *AcmePropertyTypeDeclarationContext) {}

// EnterAcmePropertyBlockEntry is called when production acmePropertyBlockEntry is entered.
func (s *BaseacmeListener) EnterAcmePropertyBlockEntry(ctx *AcmePropertyBlockEntryContext) {}

// ExitAcmePropertyBlockEntry is called when production acmePropertyBlockEntry is exited.
func (s *BaseacmeListener) ExitAcmePropertyBlockEntry(ctx *AcmePropertyBlockEntryContext) {}

// EnterAcmePropertyBlock is called when production acmePropertyBlock is entered.
func (s *BaseacmeListener) EnterAcmePropertyBlock(ctx *AcmePropertyBlockContext) {}

// ExitAcmePropertyBlock is called when production acmePropertyBlock is exited.
func (s *BaseacmeListener) ExitAcmePropertyBlock(ctx *AcmePropertyBlockContext) {}

// EnterAcmePropertyTypeInt is called when production acmePropertyTypeInt is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeInt(ctx *AcmePropertyTypeIntContext) {}

// ExitAcmePropertyTypeInt is called when production acmePropertyTypeInt is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeInt(ctx *AcmePropertyTypeIntContext) {}

// EnterAcmePropertyTypeAny is called when production acmePropertyTypeAny is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeAny(ctx *AcmePropertyTypeAnyContext) {}

// ExitAcmePropertyTypeAny is called when production acmePropertyTypeAny is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeAny(ctx *AcmePropertyTypeAnyContext) {}

// EnterAcmePropertyTypeFloat is called when production acmePropertyTypeFloat is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeFloat(ctx *AcmePropertyTypeFloatContext) {}

// ExitAcmePropertyTypeFloat is called when production acmePropertyTypeFloat is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeFloat(ctx *AcmePropertyTypeFloatContext) {}

// EnterAcmePropertyTypeDouble is called when production acmePropertyTypeDouble is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeDouble(ctx *AcmePropertyTypeDoubleContext) {}

// ExitAcmePropertyTypeDouble is called when production acmePropertyTypeDouble is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeDouble(ctx *AcmePropertyTypeDoubleContext) {}

// EnterAcmePropertyTypeString is called when production acmePropertyTypeString is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeString(ctx *AcmePropertyTypeStringContext) {}

// ExitAcmePropertyTypeString is called when production acmePropertyTypeString is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeString(ctx *AcmePropertyTypeStringContext) {}

// EnterAcmePropertyTypeBoolean is called when production acmePropertyTypeBoolean is entered.
func (s *BaseacmeListener) EnterAcmePropertyTypeBoolean(ctx *AcmePropertyTypeBooleanContext) {}

// ExitAcmePropertyTypeBoolean is called when production acmePropertyTypeBoolean is exited.
func (s *BaseacmeListener) ExitAcmePropertyTypeBoolean(ctx *AcmePropertyTypeBooleanContext) {}

// EnterAcmeViewDeclaration is called when production acmeViewDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeViewDeclaration(ctx *AcmeViewDeclarationContext) {}

// ExitAcmeViewDeclaration is called when production acmeViewDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeViewDeclaration(ctx *AcmeViewDeclarationContext) {}

// EnterAcmeViewTypeDeclaration is called when production acmeViewTypeDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeViewTypeDeclaration(ctx *AcmeViewTypeDeclarationContext) {}

// ExitAcmeViewTypeDeclaration is called when production acmeViewTypeDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeViewTypeDeclaration(ctx *AcmeViewTypeDeclarationContext) {}

// EnterAcmeViewBody is called when production acmeViewBody is entered.
func (s *BaseacmeListener) EnterAcmeViewBody(ctx *AcmeViewBodyContext) {}

// ExitAcmeViewBody is called when production acmeViewBody is exited.
func (s *BaseacmeListener) ExitAcmeViewBody(ctx *AcmeViewBodyContext) {}

// EnterDesignRule is called when production designRule is entered.
func (s *BaseacmeListener) EnterDesignRule(ctx *DesignRuleContext) {}

// ExitDesignRule is called when production designRule is exited.
func (s *BaseacmeListener) ExitDesignRule(ctx *DesignRuleContext) {}

// EnterAcmeDesignAnalysisDeclaration is called when production acmeDesignAnalysisDeclaration is entered.
func (s *BaseacmeListener) EnterAcmeDesignAnalysisDeclaration(ctx *AcmeDesignAnalysisDeclarationContext) {
}

// ExitAcmeDesignAnalysisDeclaration is called when production acmeDesignAnalysisDeclaration is exited.
func (s *BaseacmeListener) ExitAcmeDesignAnalysisDeclaration(ctx *AcmeDesignAnalysisDeclarationContext) {
}

// EnterFormalParam is called when production formalParam is entered.
func (s *BaseacmeListener) EnterFormalParam(ctx *FormalParamContext) {}

// ExitFormalParam is called when production formalParam is exited.
func (s *BaseacmeListener) ExitFormalParam(ctx *FormalParamContext) {}

// EnterTerminatedDesignRuleExpression is called when production terminatedDesignRuleExpression is entered.
func (s *BaseacmeListener) EnterTerminatedDesignRuleExpression(ctx *TerminatedDesignRuleExpressionContext) {
}

// ExitTerminatedDesignRuleExpression is called when production terminatedDesignRuleExpression is exited.
func (s *BaseacmeListener) ExitTerminatedDesignRuleExpression(ctx *TerminatedDesignRuleExpressionContext) {
}

// EnterDesignRuleExpression is called when production designRuleExpression is entered.
func (s *BaseacmeListener) EnterDesignRuleExpression(ctx *DesignRuleExpressionContext) {}

// ExitDesignRuleExpression is called when production designRuleExpression is exited.
func (s *BaseacmeListener) ExitDesignRuleExpression(ctx *DesignRuleExpressionContext) {}

// EnterASTDRImpliesExpression is called when production aSTDRImpliesExpression is entered.
func (s *BaseacmeListener) EnterASTDRImpliesExpression(ctx *ASTDRImpliesExpressionContext) {}

// ExitASTDRImpliesExpression is called when production aSTDRImpliesExpression is exited.
func (s *BaseacmeListener) ExitASTDRImpliesExpression(ctx *ASTDRImpliesExpressionContext) {}

// EnterDRIffExpression is called when production dRIffExpression is entered.
func (s *BaseacmeListener) EnterDRIffExpression(ctx *DRIffExpressionContext) {}

// ExitDRIffExpression is called when production dRIffExpression is exited.
func (s *BaseacmeListener) ExitDRIffExpression(ctx *DRIffExpressionContext) {}

// EnterDRAndExpression is called when production dRAndExpression is entered.
func (s *BaseacmeListener) EnterDRAndExpression(ctx *DRAndExpressionContext) {}

// ExitDRAndExpression is called when production dRAndExpression is exited.
func (s *BaseacmeListener) ExitDRAndExpression(ctx *DRAndExpressionContext) {}

// EnterDRNegateExpression is called when production dRNegateExpression is entered.
func (s *BaseacmeListener) EnterDRNegateExpression(ctx *DRNegateExpressionContext) {}

// ExitDRNegateExpression is called when production dRNegateExpression is exited.
func (s *BaseacmeListener) ExitDRNegateExpression(ctx *DRNegateExpressionContext) {}

// EnterDREqualityExpression is called when production dREqualityExpression is entered.
func (s *BaseacmeListener) EnterDREqualityExpression(ctx *DREqualityExpressionContext) {}

// ExitDREqualityExpression is called when production dREqualityExpression is exited.
func (s *BaseacmeListener) ExitDREqualityExpression(ctx *DREqualityExpressionContext) {}

// EnterDRRelationalExpression is called when production dRRelationalExpression is entered.
func (s *BaseacmeListener) EnterDRRelationalExpression(ctx *DRRelationalExpressionContext) {}

// ExitDRRelationalExpression is called when production dRRelationalExpression is exited.
func (s *BaseacmeListener) ExitDRRelationalExpression(ctx *DRRelationalExpressionContext) {}

// EnterDRAdditiveExpression is called when production dRAdditiveExpression is entered.
func (s *BaseacmeListener) EnterDRAdditiveExpression(ctx *DRAdditiveExpressionContext) {}

// ExitDRAdditiveExpression is called when production dRAdditiveExpression is exited.
func (s *BaseacmeListener) ExitDRAdditiveExpression(ctx *DRAdditiveExpressionContext) {}

// EnterDRMultiplicativeExpression is called when production dRMultiplicativeExpression is entered.
func (s *BaseacmeListener) EnterDRMultiplicativeExpression(ctx *DRMultiplicativeExpressionContext) {}

// ExitDRMultiplicativeExpression is called when production dRMultiplicativeExpression is exited.
func (s *BaseacmeListener) ExitDRMultiplicativeExpression(ctx *DRMultiplicativeExpressionContext) {}

// EnterDRNegativeExpression is called when production dRNegativeExpression is entered.
func (s *BaseacmeListener) EnterDRNegativeExpression(ctx *DRNegativeExpressionContext) {}

// ExitDRNegativeExpression is called when production dRNegativeExpression is exited.
func (s *BaseacmeListener) ExitDRNegativeExpression(ctx *DRNegativeExpressionContext) {}

// EnterPrimitiveExpression is called when production primitiveExpression is entered.
func (s *BaseacmeListener) EnterPrimitiveExpression(ctx *PrimitiveExpressionContext) {}

// ExitPrimitiveExpression is called when production primitiveExpression is exited.
func (s *BaseacmeListener) ExitPrimitiveExpression(ctx *PrimitiveExpressionContext) {}

// EnterParentheticalExpression is called when production parentheticalExpression is entered.
func (s *BaseacmeListener) EnterParentheticalExpression(ctx *ParentheticalExpressionContext) {}

// ExitParentheticalExpression is called when production parentheticalExpression is exited.
func (s *BaseacmeListener) ExitParentheticalExpression(ctx *ParentheticalExpressionContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseacmeListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseacmeListener) ExitReference(ctx *ReferenceContext) {}

// EnterSetReference is called when production setReference is entered.
func (s *BaseacmeListener) EnterSetReference(ctx *SetReferenceContext) {}

// ExitSetReference is called when production setReference is exited.
func (s *BaseacmeListener) ExitSetReference(ctx *SetReferenceContext) {}

// EnterActualParams is called when production actualParams is entered.
func (s *BaseacmeListener) EnterActualParams(ctx *ActualParamsContext) {}

// ExitActualParams is called when production actualParams is exited.
func (s *BaseacmeListener) ExitActualParams(ctx *ActualParamsContext) {}

// EnterLiteralConstant is called when production literalConstant is entered.
func (s *BaseacmeListener) EnterLiteralConstant(ctx *LiteralConstantContext) {}

// ExitLiteralConstant is called when production literalConstant is exited.
func (s *BaseacmeListener) ExitLiteralConstant(ctx *LiteralConstantContext) {}

// EnterQuantifiedExpression is called when production quantifiedExpression is entered.
func (s *BaseacmeListener) EnterQuantifiedExpression(ctx *QuantifiedExpressionContext) {}

// ExitQuantifiedExpression is called when production quantifiedExpression is exited.
func (s *BaseacmeListener) ExitQuantifiedExpression(ctx *QuantifiedExpressionContext) {}

// EnterDistinctVariableSetDeclaration is called when production distinctVariableSetDeclaration is entered.
func (s *BaseacmeListener) EnterDistinctVariableSetDeclaration(ctx *DistinctVariableSetDeclarationContext) {
}

// ExitDistinctVariableSetDeclaration is called when production distinctVariableSetDeclaration is exited.
func (s *BaseacmeListener) ExitDistinctVariableSetDeclaration(ctx *DistinctVariableSetDeclarationContext) {
}

// EnterVariableSetDeclaration is called when production variableSetDeclaration is entered.
func (s *BaseacmeListener) EnterVariableSetDeclaration(ctx *VariableSetDeclarationContext) {}

// ExitVariableSetDeclaration is called when production variableSetDeclaration is exited.
func (s *BaseacmeListener) ExitVariableSetDeclaration(ctx *VariableSetDeclarationContext) {}

// EnterSequenceExpression is called when production sequenceExpression is entered.
func (s *BaseacmeListener) EnterSequenceExpression(ctx *SequenceExpressionContext) {}

// ExitSequenceExpression is called when production sequenceExpression is exited.
func (s *BaseacmeListener) ExitSequenceExpression(ctx *SequenceExpressionContext) {}

// EnterSetExpression is called when production setExpression is entered.
func (s *BaseacmeListener) EnterSetExpression(ctx *SetExpressionContext) {}

// ExitSetExpression is called when production setExpression is exited.
func (s *BaseacmeListener) ExitSetExpression(ctx *SetExpressionContext) {}

// EnterPathExpression is called when production pathExpression is entered.
func (s *BaseacmeListener) EnterPathExpression(ctx *PathExpressionContext) {}

// ExitPathExpression is called when production pathExpression is exited.
func (s *BaseacmeListener) ExitPathExpression(ctx *PathExpressionContext) {}

// EnterPathExpressionContinuation is called when production pathExpressionContinuation is entered.
func (s *BaseacmeListener) EnterPathExpressionContinuation(ctx *PathExpressionContinuationContext) {}

// ExitPathExpressionContinuation is called when production pathExpressionContinuation is exited.
func (s *BaseacmeListener) ExitPathExpressionContinuation(ctx *PathExpressionContinuationContext) {}

// EnterLiteralSet is called when production literalSet is entered.
func (s *BaseacmeListener) EnterLiteralSet(ctx *LiteralSetContext) {}

// ExitLiteralSet is called when production literalSet is exited.
func (s *BaseacmeListener) ExitLiteralSet(ctx *LiteralSetContext) {}

// EnterLiteralSequence is called when production literalSequence is entered.
func (s *BaseacmeListener) EnterLiteralSequence(ctx *LiteralSequenceContext) {}

// ExitLiteralSequence is called when production literalSequence is exited.
func (s *BaseacmeListener) ExitLiteralSequence(ctx *LiteralSequenceContext) {}

// EnterLiteralRecordEntry is called when production literalRecordEntry is entered.
func (s *BaseacmeListener) EnterLiteralRecordEntry(ctx *LiteralRecordEntryContext) {}

// ExitLiteralRecordEntry is called when production literalRecordEntry is exited.
func (s *BaseacmeListener) ExitLiteralRecordEntry(ctx *LiteralRecordEntryContext) {}

// EnterLiteralRecord is called when production literalRecord is entered.
func (s *BaseacmeListener) EnterLiteralRecord(ctx *LiteralRecordContext) {}

// ExitLiteralRecord is called when production literalRecord is exited.
func (s *BaseacmeListener) ExitLiteralRecord(ctx *LiteralRecordContext) {}

// EnterSetConstructor is called when production setConstructor is entered.
func (s *BaseacmeListener) EnterSetConstructor(ctx *SetConstructorContext) {}

// ExitSetConstructor is called when production setConstructor is exited.
func (s *BaseacmeListener) ExitSetConstructor(ctx *SetConstructorContext) {}

// EnterAcmeTypeRef is called when production acmeTypeRef is entered.
func (s *BaseacmeListener) EnterAcmeTypeRef(ctx *AcmeTypeRefContext) {}

// ExitAcmeTypeRef is called when production acmeTypeRef is exited.
func (s *BaseacmeListener) ExitAcmeTypeRef(ctx *AcmeTypeRefContext) {}
