// Code generated from WebIDL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package webidl // WebIDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// WebIDLListener is a complete listener for a parse tree produced by WebIDLParser.
type WebIDLListener interface {
	antlr.ParseTreeListener

	// EnterWebIDL is called when entering the webIDL production.
	EnterWebIDL(c *WebIDLContext)

	// EnterDefinitions is called when entering the definitions production.
	EnterDefinitions(c *DefinitionsContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterArgumentNameKeyword is called when entering the argumentNameKeyword production.
	EnterArgumentNameKeyword(c *ArgumentNameKeywordContext)

	// EnterCallbackOrInterfaceOrMixin is called when entering the callbackOrInterfaceOrMixin production.
	EnterCallbackOrInterfaceOrMixin(c *CallbackOrInterfaceOrMixinContext)

	// EnterInterfaceOrMixin is called when entering the interfaceOrMixin production.
	EnterInterfaceOrMixin(c *InterfaceOrMixinContext)

	// EnterInterfaceRest is called when entering the interfaceRest production.
	EnterInterfaceRest(c *InterfaceRestContext)

	// EnterPartial is called when entering the partial production.
	EnterPartial(c *PartialContext)

	// EnterPartialDefinition is called when entering the partialDefinition production.
	EnterPartialDefinition(c *PartialDefinitionContext)

	// EnterPartialInterfaceOrPartialMixin is called when entering the partialInterfaceOrPartialMixin production.
	EnterPartialInterfaceOrPartialMixin(c *PartialInterfaceOrPartialMixinContext)

	// EnterPartialInterfaceRest is called when entering the partialInterfaceRest production.
	EnterPartialInterfaceRest(c *PartialInterfaceRestContext)

	// EnterInterfaceMembers is called when entering the interfaceMembers production.
	EnterInterfaceMembers(c *InterfaceMembersContext)

	// EnterInterfaceMember is called when entering the interfaceMember production.
	EnterInterfaceMember(c *InterfaceMemberContext)

	// EnterPartialInterfaceMembers is called when entering the partialInterfaceMembers production.
	EnterPartialInterfaceMembers(c *PartialInterfaceMembersContext)

	// EnterPartialInterfaceMember is called when entering the partialInterfaceMember production.
	EnterPartialInterfaceMember(c *PartialInterfaceMemberContext)

	// EnterInheritance is called when entering the inheritance production.
	EnterInheritance(c *InheritanceContext)

	// EnterMixinRest is called when entering the mixinRest production.
	EnterMixinRest(c *MixinRestContext)

	// EnterMixinMembers is called when entering the mixinMembers production.
	EnterMixinMembers(c *MixinMembersContext)

	// EnterMixinMember is called when entering the mixinMember production.
	EnterMixinMember(c *MixinMemberContext)

	// EnterIncludesStatement is called when entering the includesStatement production.
	EnterIncludesStatement(c *IncludesStatementContext)

	// EnterCallbackRestOrInterface is called when entering the callbackRestOrInterface production.
	EnterCallbackRestOrInterface(c *CallbackRestOrInterfaceContext)

	// EnterCallbackInterfaceMembers is called when entering the callbackInterfaceMembers production.
	EnterCallbackInterfaceMembers(c *CallbackInterfaceMembersContext)

	// EnterCallbackInterfaceMember is called when entering the callbackInterfaceMember production.
	EnterCallbackInterfaceMember(c *CallbackInterfaceMemberContext)

	// EnterConst_ is called when entering the const_ production.
	EnterConst_(c *Const_Context)

	// EnterConstValue is called when entering the constValue production.
	EnterConstValue(c *ConstValueContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterConstType is called when entering the constType production.
	EnterConstType(c *ConstTypeContext)

	// EnterReadonlyMember is called when entering the readonlyMember production.
	EnterReadonlyMember(c *ReadonlyMemberContext)

	// EnterReadonlyMemberRest is called when entering the readonlyMemberRest production.
	EnterReadonlyMemberRest(c *ReadonlyMemberRestContext)

	// EnterReadWriteAttribute is called when entering the readWriteAttribute production.
	EnterReadWriteAttribute(c *ReadWriteAttributeContext)

	// EnterInheritAttribute is called when entering the inheritAttribute production.
	EnterInheritAttribute(c *InheritAttributeContext)

	// EnterAttributeRest is called when entering the attributeRest production.
	EnterAttributeRest(c *AttributeRestContext)

	// EnterAttributeName is called when entering the attributeName production.
	EnterAttributeName(c *AttributeNameContext)

	// EnterAttributeNameKeyword is called when entering the attributeNameKeyword production.
	EnterAttributeNameKeyword(c *AttributeNameKeywordContext)

	// EnterOptionalReadOnly is called when entering the optionalReadOnly production.
	EnterOptionalReadOnly(c *OptionalReadOnlyContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterRegularOperation is called when entering the regularOperation production.
	EnterRegularOperation(c *RegularOperationContext)

	// EnterSpecialOperation is called when entering the specialOperation production.
	EnterSpecialOperation(c *SpecialOperationContext)

	// EnterSpecial is called when entering the special production.
	EnterSpecial(c *SpecialContext)

	// EnterOperationRest is called when entering the operationRest production.
	EnterOperationRest(c *OperationRestContext)

	// EnterOptionalOperationName is called when entering the optionalOperationName production.
	EnterOptionalOperationName(c *OptionalOperationNameContext)

	// EnterOperationName is called when entering the operationName production.
	EnterOperationName(c *OperationNameContext)

	// EnterOperationNameKeyword is called when entering the operationNameKeyword production.
	EnterOperationNameKeyword(c *OperationNameKeywordContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterArgumentRest is called when entering the argumentRest production.
	EnterArgumentRest(c *ArgumentRestContext)

	// EnterArgumentName is called when entering the argumentName production.
	EnterArgumentName(c *ArgumentNameContext)

	// EnterEllipsis is called when entering the ellipsis production.
	EnterEllipsis(c *EllipsisContext)

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterStringifier is called when entering the stringifier production.
	EnterStringifier(c *StringifierContext)

	// EnterStringifierRest is called when entering the stringifierRest production.
	EnterStringifierRest(c *StringifierRestContext)

	// EnterStaticMember is called when entering the staticMember production.
	EnterStaticMember(c *StaticMemberContext)

	// EnterStaticMemberRest is called when entering the staticMemberRest production.
	EnterStaticMemberRest(c *StaticMemberRestContext)

	// EnterIterable is called when entering the iterable production.
	EnterIterable(c *IterableContext)

	// EnterOptionalType is called when entering the optionalType production.
	EnterOptionalType(c *OptionalTypeContext)

	// EnterAsyncIterable is called when entering the asyncIterable production.
	EnterAsyncIterable(c *AsyncIterableContext)

	// EnterOptionalArgumentList is called when entering the optionalArgumentList production.
	EnterOptionalArgumentList(c *OptionalArgumentListContext)

	// EnterReadWriteMaplike is called when entering the readWriteMaplike production.
	EnterReadWriteMaplike(c *ReadWriteMaplikeContext)

	// EnterMaplikeRest is called when entering the maplikeRest production.
	EnterMaplikeRest(c *MaplikeRestContext)

	// EnterReadWriteSetlike is called when entering the readWriteSetlike production.
	EnterReadWriteSetlike(c *ReadWriteSetlikeContext)

	// EnterSetlikeRest is called when entering the setlikeRest production.
	EnterSetlikeRest(c *SetlikeRestContext)

	// EnterNamespace_ is called when entering the namespace_ production.
	EnterNamespace_(c *Namespace_Context)

	// EnterNamespaceMembers is called when entering the namespaceMembers production.
	EnterNamespaceMembers(c *NamespaceMembersContext)

	// EnterNamespaceMember is called when entering the namespaceMember production.
	EnterNamespaceMember(c *NamespaceMemberContext)

	// EnterDictionary is called when entering the dictionary production.
	EnterDictionary(c *DictionaryContext)

	// EnterDictionaryMembers is called when entering the dictionaryMembers production.
	EnterDictionaryMembers(c *DictionaryMembersContext)

	// EnterDictionaryMember is called when entering the dictionaryMember production.
	EnterDictionaryMember(c *DictionaryMemberContext)

	// EnterDictionaryMemberRest is called when entering the dictionaryMemberRest production.
	EnterDictionaryMemberRest(c *DictionaryMemberRestContext)

	// EnterPartialDictionary is called when entering the partialDictionary production.
	EnterPartialDictionary(c *PartialDictionaryContext)

	// EnterDefault_ is called when entering the default_ production.
	EnterDefault_(c *Default_Context)

	// EnterEnum_ is called when entering the enum_ production.
	EnterEnum_(c *Enum_Context)

	// EnterEnumValueList is called when entering the enumValueList production.
	EnterEnumValueList(c *EnumValueListContext)

	// EnterEnumValueListComma is called when entering the enumValueListComma production.
	EnterEnumValueListComma(c *EnumValueListCommaContext)

	// EnterEnumValueListString is called when entering the enumValueListString production.
	EnterEnumValueListString(c *EnumValueListStringContext)

	// EnterCallbackRest is called when entering the callbackRest production.
	EnterCallbackRest(c *CallbackRestContext)

	// EnterTypedef_ is called when entering the typedef_ production.
	EnterTypedef_(c *Typedef_Context)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeWithExtendedAttributes is called when entering the typeWithExtendedAttributes production.
	EnterTypeWithExtendedAttributes(c *TypeWithExtendedAttributesContext)

	// EnterSingleType is called when entering the singleType production.
	EnterSingleType(c *SingleTypeContext)

	// EnterUnionType is called when entering the unionType production.
	EnterUnionType(c *UnionTypeContext)

	// EnterUnionMemberType is called when entering the unionMemberType production.
	EnterUnionMemberType(c *UnionMemberTypeContext)

	// EnterUnionMemberTypes is called when entering the unionMemberTypes production.
	EnterUnionMemberTypes(c *UnionMemberTypesContext)

	// EnterDistinguishableType is called when entering the distinguishableType production.
	EnterDistinguishableType(c *DistinguishableTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterUnrestrictedFloatType is called when entering the unrestrictedFloatType production.
	EnterUnrestrictedFloatType(c *UnrestrictedFloatTypeContext)

	// EnterFloatType is called when entering the floatType production.
	EnterFloatType(c *FloatTypeContext)

	// EnterUnsignedIntegerType is called when entering the unsignedIntegerType production.
	EnterUnsignedIntegerType(c *UnsignedIntegerTypeContext)

	// EnterIntegerType is called when entering the integerType production.
	EnterIntegerType(c *IntegerTypeContext)

	// EnterOptionalLong is called when entering the optionalLong production.
	EnterOptionalLong(c *OptionalLongContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterPromiseType is called when entering the promiseType production.
	EnterPromiseType(c *PromiseTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterNull_ is called when entering the null_ production.
	EnterNull_(c *Null_Context)

	// EnterBufferRelatedType is called when entering the bufferRelatedType production.
	EnterBufferRelatedType(c *BufferRelatedTypeContext)

	// EnterExtendedAttributeList is called when entering the extendedAttributeList production.
	EnterExtendedAttributeList(c *ExtendedAttributeListContext)

	// EnterExtendedAttributes is called when entering the extendedAttributes production.
	EnterExtendedAttributes(c *ExtendedAttributesContext)

	// EnterExtendedAttribute is called when entering the extendedAttribute production.
	EnterExtendedAttribute(c *ExtendedAttributeContext)

	// EnterOther is called when entering the other production.
	EnterOther(c *OtherContext)

	// EnterOtherOrComma is called when entering the otherOrComma production.
	EnterOtherOrComma(c *OtherOrCommaContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIdentifiers is called when entering the identifiers production.
	EnterIdentifiers(c *IdentifiersContext)

	// EnterExtendedAttributeNoArgs is called when entering the extendedAttributeNoArgs production.
	EnterExtendedAttributeNoArgs(c *ExtendedAttributeNoArgsContext)

	// EnterExtendedAttributeArgList is called when entering the extendedAttributeArgList production.
	EnterExtendedAttributeArgList(c *ExtendedAttributeArgListContext)

	// EnterExtendedAttributeIdent is called when entering the extendedAttributeIdent production.
	EnterExtendedAttributeIdent(c *ExtendedAttributeIdentContext)

	// EnterExtendedAttributeIdentList is called when entering the extendedAttributeIdentList production.
	EnterExtendedAttributeIdentList(c *ExtendedAttributeIdentListContext)

	// EnterExtendedAttributeNamedArgList is called when entering the extendedAttributeNamedArgList production.
	EnterExtendedAttributeNamedArgList(c *ExtendedAttributeNamedArgListContext)

	// EnterExtendedAttributeString is called when entering the extendedAttributeString production.
	EnterExtendedAttributeString(c *ExtendedAttributeStringContext)

	// EnterExtendedAttributeStringList is called when entering the extendedAttributeStringList production.
	EnterExtendedAttributeStringList(c *ExtendedAttributeStringListContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterStrings is called when entering the strings production.
	EnterStrings(c *StringsContext)

	// ExitWebIDL is called when exiting the webIDL production.
	ExitWebIDL(c *WebIDLContext)

	// ExitDefinitions is called when exiting the definitions production.
	ExitDefinitions(c *DefinitionsContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitArgumentNameKeyword is called when exiting the argumentNameKeyword production.
	ExitArgumentNameKeyword(c *ArgumentNameKeywordContext)

	// ExitCallbackOrInterfaceOrMixin is called when exiting the callbackOrInterfaceOrMixin production.
	ExitCallbackOrInterfaceOrMixin(c *CallbackOrInterfaceOrMixinContext)

	// ExitInterfaceOrMixin is called when exiting the interfaceOrMixin production.
	ExitInterfaceOrMixin(c *InterfaceOrMixinContext)

	// ExitInterfaceRest is called when exiting the interfaceRest production.
	ExitInterfaceRest(c *InterfaceRestContext)

	// ExitPartial is called when exiting the partial production.
	ExitPartial(c *PartialContext)

	// ExitPartialDefinition is called when exiting the partialDefinition production.
	ExitPartialDefinition(c *PartialDefinitionContext)

	// ExitPartialInterfaceOrPartialMixin is called when exiting the partialInterfaceOrPartialMixin production.
	ExitPartialInterfaceOrPartialMixin(c *PartialInterfaceOrPartialMixinContext)

	// ExitPartialInterfaceRest is called when exiting the partialInterfaceRest production.
	ExitPartialInterfaceRest(c *PartialInterfaceRestContext)

	// ExitInterfaceMembers is called when exiting the interfaceMembers production.
	ExitInterfaceMembers(c *InterfaceMembersContext)

	// ExitInterfaceMember is called when exiting the interfaceMember production.
	ExitInterfaceMember(c *InterfaceMemberContext)

	// ExitPartialInterfaceMembers is called when exiting the partialInterfaceMembers production.
	ExitPartialInterfaceMembers(c *PartialInterfaceMembersContext)

	// ExitPartialInterfaceMember is called when exiting the partialInterfaceMember production.
	ExitPartialInterfaceMember(c *PartialInterfaceMemberContext)

	// ExitInheritance is called when exiting the inheritance production.
	ExitInheritance(c *InheritanceContext)

	// ExitMixinRest is called when exiting the mixinRest production.
	ExitMixinRest(c *MixinRestContext)

	// ExitMixinMembers is called when exiting the mixinMembers production.
	ExitMixinMembers(c *MixinMembersContext)

	// ExitMixinMember is called when exiting the mixinMember production.
	ExitMixinMember(c *MixinMemberContext)

	// ExitIncludesStatement is called when exiting the includesStatement production.
	ExitIncludesStatement(c *IncludesStatementContext)

	// ExitCallbackRestOrInterface is called when exiting the callbackRestOrInterface production.
	ExitCallbackRestOrInterface(c *CallbackRestOrInterfaceContext)

	// ExitCallbackInterfaceMembers is called when exiting the callbackInterfaceMembers production.
	ExitCallbackInterfaceMembers(c *CallbackInterfaceMembersContext)

	// ExitCallbackInterfaceMember is called when exiting the callbackInterfaceMember production.
	ExitCallbackInterfaceMember(c *CallbackInterfaceMemberContext)

	// ExitConst_ is called when exiting the const_ production.
	ExitConst_(c *Const_Context)

	// ExitConstValue is called when exiting the constValue production.
	ExitConstValue(c *ConstValueContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitConstType is called when exiting the constType production.
	ExitConstType(c *ConstTypeContext)

	// ExitReadonlyMember is called when exiting the readonlyMember production.
	ExitReadonlyMember(c *ReadonlyMemberContext)

	// ExitReadonlyMemberRest is called when exiting the readonlyMemberRest production.
	ExitReadonlyMemberRest(c *ReadonlyMemberRestContext)

	// ExitReadWriteAttribute is called when exiting the readWriteAttribute production.
	ExitReadWriteAttribute(c *ReadWriteAttributeContext)

	// ExitInheritAttribute is called when exiting the inheritAttribute production.
	ExitInheritAttribute(c *InheritAttributeContext)

	// ExitAttributeRest is called when exiting the attributeRest production.
	ExitAttributeRest(c *AttributeRestContext)

	// ExitAttributeName is called when exiting the attributeName production.
	ExitAttributeName(c *AttributeNameContext)

	// ExitAttributeNameKeyword is called when exiting the attributeNameKeyword production.
	ExitAttributeNameKeyword(c *AttributeNameKeywordContext)

	// ExitOptionalReadOnly is called when exiting the optionalReadOnly production.
	ExitOptionalReadOnly(c *OptionalReadOnlyContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitRegularOperation is called when exiting the regularOperation production.
	ExitRegularOperation(c *RegularOperationContext)

	// ExitSpecialOperation is called when exiting the specialOperation production.
	ExitSpecialOperation(c *SpecialOperationContext)

	// ExitSpecial is called when exiting the special production.
	ExitSpecial(c *SpecialContext)

	// ExitOperationRest is called when exiting the operationRest production.
	ExitOperationRest(c *OperationRestContext)

	// ExitOptionalOperationName is called when exiting the optionalOperationName production.
	ExitOptionalOperationName(c *OptionalOperationNameContext)

	// ExitOperationName is called when exiting the operationName production.
	ExitOperationName(c *OperationNameContext)

	// ExitOperationNameKeyword is called when exiting the operationNameKeyword production.
	ExitOperationNameKeyword(c *OperationNameKeywordContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitArgumentRest is called when exiting the argumentRest production.
	ExitArgumentRest(c *ArgumentRestContext)

	// ExitArgumentName is called when exiting the argumentName production.
	ExitArgumentName(c *ArgumentNameContext)

	// ExitEllipsis is called when exiting the ellipsis production.
	ExitEllipsis(c *EllipsisContext)

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitStringifier is called when exiting the stringifier production.
	ExitStringifier(c *StringifierContext)

	// ExitStringifierRest is called when exiting the stringifierRest production.
	ExitStringifierRest(c *StringifierRestContext)

	// ExitStaticMember is called when exiting the staticMember production.
	ExitStaticMember(c *StaticMemberContext)

	// ExitStaticMemberRest is called when exiting the staticMemberRest production.
	ExitStaticMemberRest(c *StaticMemberRestContext)

	// ExitIterable is called when exiting the iterable production.
	ExitIterable(c *IterableContext)

	// ExitOptionalType is called when exiting the optionalType production.
	ExitOptionalType(c *OptionalTypeContext)

	// ExitAsyncIterable is called when exiting the asyncIterable production.
	ExitAsyncIterable(c *AsyncIterableContext)

	// ExitOptionalArgumentList is called when exiting the optionalArgumentList production.
	ExitOptionalArgumentList(c *OptionalArgumentListContext)

	// ExitReadWriteMaplike is called when exiting the readWriteMaplike production.
	ExitReadWriteMaplike(c *ReadWriteMaplikeContext)

	// ExitMaplikeRest is called when exiting the maplikeRest production.
	ExitMaplikeRest(c *MaplikeRestContext)

	// ExitReadWriteSetlike is called when exiting the readWriteSetlike production.
	ExitReadWriteSetlike(c *ReadWriteSetlikeContext)

	// ExitSetlikeRest is called when exiting the setlikeRest production.
	ExitSetlikeRest(c *SetlikeRestContext)

	// ExitNamespace_ is called when exiting the namespace_ production.
	ExitNamespace_(c *Namespace_Context)

	// ExitNamespaceMembers is called when exiting the namespaceMembers production.
	ExitNamespaceMembers(c *NamespaceMembersContext)

	// ExitNamespaceMember is called when exiting the namespaceMember production.
	ExitNamespaceMember(c *NamespaceMemberContext)

	// ExitDictionary is called when exiting the dictionary production.
	ExitDictionary(c *DictionaryContext)

	// ExitDictionaryMembers is called when exiting the dictionaryMembers production.
	ExitDictionaryMembers(c *DictionaryMembersContext)

	// ExitDictionaryMember is called when exiting the dictionaryMember production.
	ExitDictionaryMember(c *DictionaryMemberContext)

	// ExitDictionaryMemberRest is called when exiting the dictionaryMemberRest production.
	ExitDictionaryMemberRest(c *DictionaryMemberRestContext)

	// ExitPartialDictionary is called when exiting the partialDictionary production.
	ExitPartialDictionary(c *PartialDictionaryContext)

	// ExitDefault_ is called when exiting the default_ production.
	ExitDefault_(c *Default_Context)

	// ExitEnum_ is called when exiting the enum_ production.
	ExitEnum_(c *Enum_Context)

	// ExitEnumValueList is called when exiting the enumValueList production.
	ExitEnumValueList(c *EnumValueListContext)

	// ExitEnumValueListComma is called when exiting the enumValueListComma production.
	ExitEnumValueListComma(c *EnumValueListCommaContext)

	// ExitEnumValueListString is called when exiting the enumValueListString production.
	ExitEnumValueListString(c *EnumValueListStringContext)

	// ExitCallbackRest is called when exiting the callbackRest production.
	ExitCallbackRest(c *CallbackRestContext)

	// ExitTypedef_ is called when exiting the typedef_ production.
	ExitTypedef_(c *Typedef_Context)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeWithExtendedAttributes is called when exiting the typeWithExtendedAttributes production.
	ExitTypeWithExtendedAttributes(c *TypeWithExtendedAttributesContext)

	// ExitSingleType is called when exiting the singleType production.
	ExitSingleType(c *SingleTypeContext)

	// ExitUnionType is called when exiting the unionType production.
	ExitUnionType(c *UnionTypeContext)

	// ExitUnionMemberType is called when exiting the unionMemberType production.
	ExitUnionMemberType(c *UnionMemberTypeContext)

	// ExitUnionMemberTypes is called when exiting the unionMemberTypes production.
	ExitUnionMemberTypes(c *UnionMemberTypesContext)

	// ExitDistinguishableType is called when exiting the distinguishableType production.
	ExitDistinguishableType(c *DistinguishableTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitUnrestrictedFloatType is called when exiting the unrestrictedFloatType production.
	ExitUnrestrictedFloatType(c *UnrestrictedFloatTypeContext)

	// ExitFloatType is called when exiting the floatType production.
	ExitFloatType(c *FloatTypeContext)

	// ExitUnsignedIntegerType is called when exiting the unsignedIntegerType production.
	ExitUnsignedIntegerType(c *UnsignedIntegerTypeContext)

	// ExitIntegerType is called when exiting the integerType production.
	ExitIntegerType(c *IntegerTypeContext)

	// ExitOptionalLong is called when exiting the optionalLong production.
	ExitOptionalLong(c *OptionalLongContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitPromiseType is called when exiting the promiseType production.
	ExitPromiseType(c *PromiseTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitNull_ is called when exiting the null_ production.
	ExitNull_(c *Null_Context)

	// ExitBufferRelatedType is called when exiting the bufferRelatedType production.
	ExitBufferRelatedType(c *BufferRelatedTypeContext)

	// ExitExtendedAttributeList is called when exiting the extendedAttributeList production.
	ExitExtendedAttributeList(c *ExtendedAttributeListContext)

	// ExitExtendedAttributes is called when exiting the extendedAttributes production.
	ExitExtendedAttributes(c *ExtendedAttributesContext)

	// ExitExtendedAttribute is called when exiting the extendedAttribute production.
	ExitExtendedAttribute(c *ExtendedAttributeContext)

	// ExitOther is called when exiting the other production.
	ExitOther(c *OtherContext)

	// ExitOtherOrComma is called when exiting the otherOrComma production.
	ExitOtherOrComma(c *OtherOrCommaContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIdentifiers is called when exiting the identifiers production.
	ExitIdentifiers(c *IdentifiersContext)

	// ExitExtendedAttributeNoArgs is called when exiting the extendedAttributeNoArgs production.
	ExitExtendedAttributeNoArgs(c *ExtendedAttributeNoArgsContext)

	// ExitExtendedAttributeArgList is called when exiting the extendedAttributeArgList production.
	ExitExtendedAttributeArgList(c *ExtendedAttributeArgListContext)

	// ExitExtendedAttributeIdent is called when exiting the extendedAttributeIdent production.
	ExitExtendedAttributeIdent(c *ExtendedAttributeIdentContext)

	// ExitExtendedAttributeIdentList is called when exiting the extendedAttributeIdentList production.
	ExitExtendedAttributeIdentList(c *ExtendedAttributeIdentListContext)

	// ExitExtendedAttributeNamedArgList is called when exiting the extendedAttributeNamedArgList production.
	ExitExtendedAttributeNamedArgList(c *ExtendedAttributeNamedArgListContext)

	// ExitExtendedAttributeString is called when exiting the extendedAttributeString production.
	ExitExtendedAttributeString(c *ExtendedAttributeStringContext)

	// ExitExtendedAttributeStringList is called when exiting the extendedAttributeStringList production.
	ExitExtendedAttributeStringList(c *ExtendedAttributeStringListContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitStrings is called when exiting the strings production.
	ExitStrings(c *StringsContext)
}
