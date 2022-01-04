// Code generated from WebIDL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package webidl // WebIDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWebIDLListener is a complete listener for a parse tree produced by WebIDLParser.
type BaseWebIDLListener struct{}

var _ WebIDLListener = &BaseWebIDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWebIDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWebIDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWebIDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWebIDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWebIDL is called when production webIDL is entered.
func (s *BaseWebIDLListener) EnterWebIDL(ctx *WebIDLContext) {}

// ExitWebIDL is called when production webIDL is exited.
func (s *BaseWebIDLListener) ExitWebIDL(ctx *WebIDLContext) {}

// EnterDefinitions is called when production definitions is entered.
func (s *BaseWebIDLListener) EnterDefinitions(ctx *DefinitionsContext) {}

// ExitDefinitions is called when production definitions is exited.
func (s *BaseWebIDLListener) ExitDefinitions(ctx *DefinitionsContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseWebIDLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseWebIDLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterArgumentNameKeyword is called when production argumentNameKeyword is entered.
func (s *BaseWebIDLListener) EnterArgumentNameKeyword(ctx *ArgumentNameKeywordContext) {}

// ExitArgumentNameKeyword is called when production argumentNameKeyword is exited.
func (s *BaseWebIDLListener) ExitArgumentNameKeyword(ctx *ArgumentNameKeywordContext) {}

// EnterCallbackOrInterfaceOrMixin is called when production callbackOrInterfaceOrMixin is entered.
func (s *BaseWebIDLListener) EnterCallbackOrInterfaceOrMixin(ctx *CallbackOrInterfaceOrMixinContext) {
}

// ExitCallbackOrInterfaceOrMixin is called when production callbackOrInterfaceOrMixin is exited.
func (s *BaseWebIDLListener) ExitCallbackOrInterfaceOrMixin(ctx *CallbackOrInterfaceOrMixinContext) {}

// EnterInterfaceOrMixin is called when production interfaceOrMixin is entered.
func (s *BaseWebIDLListener) EnterInterfaceOrMixin(ctx *InterfaceOrMixinContext) {}

// ExitInterfaceOrMixin is called when production interfaceOrMixin is exited.
func (s *BaseWebIDLListener) ExitInterfaceOrMixin(ctx *InterfaceOrMixinContext) {}

// EnterInterfaceRest is called when production interfaceRest is entered.
func (s *BaseWebIDLListener) EnterInterfaceRest(ctx *InterfaceRestContext) {}

// ExitInterfaceRest is called when production interfaceRest is exited.
func (s *BaseWebIDLListener) ExitInterfaceRest(ctx *InterfaceRestContext) {}

// EnterPartial is called when production partial is entered.
func (s *BaseWebIDLListener) EnterPartial(ctx *PartialContext) {}

// ExitPartial is called when production partial is exited.
func (s *BaseWebIDLListener) ExitPartial(ctx *PartialContext) {}

// EnterPartialDefinition is called when production partialDefinition is entered.
func (s *BaseWebIDLListener) EnterPartialDefinition(ctx *PartialDefinitionContext) {}

// ExitPartialDefinition is called when production partialDefinition is exited.
func (s *BaseWebIDLListener) ExitPartialDefinition(ctx *PartialDefinitionContext) {}

// EnterPartialInterfaceOrPartialMixin is called when production partialInterfaceOrPartialMixin is entered.
func (s *BaseWebIDLListener) EnterPartialInterfaceOrPartialMixin(ctx *PartialInterfaceOrPartialMixinContext) {
}

// ExitPartialInterfaceOrPartialMixin is called when production partialInterfaceOrPartialMixin is exited.
func (s *BaseWebIDLListener) ExitPartialInterfaceOrPartialMixin(ctx *PartialInterfaceOrPartialMixinContext) {
}

// EnterPartialInterfaceRest is called when production partialInterfaceRest is entered.
func (s *BaseWebIDLListener) EnterPartialInterfaceRest(ctx *PartialInterfaceRestContext) {}

// ExitPartialInterfaceRest is called when production partialInterfaceRest is exited.
func (s *BaseWebIDLListener) ExitPartialInterfaceRest(ctx *PartialInterfaceRestContext) {}

// EnterInterfaceMembers is called when production interfaceMembers is entered.
func (s *BaseWebIDLListener) EnterInterfaceMembers(ctx *InterfaceMembersContext) {}

// ExitInterfaceMembers is called when production interfaceMembers is exited.
func (s *BaseWebIDLListener) ExitInterfaceMembers(ctx *InterfaceMembersContext) {}

// EnterInterfaceMember is called when production interfaceMember is entered.
func (s *BaseWebIDLListener) EnterInterfaceMember(ctx *InterfaceMemberContext) {}

// ExitInterfaceMember is called when production interfaceMember is exited.
func (s *BaseWebIDLListener) ExitInterfaceMember(ctx *InterfaceMemberContext) {}

// EnterPartialInterfaceMembers is called when production partialInterfaceMembers is entered.
func (s *BaseWebIDLListener) EnterPartialInterfaceMembers(ctx *PartialInterfaceMembersContext) {}

// ExitPartialInterfaceMembers is called when production partialInterfaceMembers is exited.
func (s *BaseWebIDLListener) ExitPartialInterfaceMembers(ctx *PartialInterfaceMembersContext) {}

// EnterPartialInterfaceMember is called when production partialInterfaceMember is entered.
func (s *BaseWebIDLListener) EnterPartialInterfaceMember(ctx *PartialInterfaceMemberContext) {}

// ExitPartialInterfaceMember is called when production partialInterfaceMember is exited.
func (s *BaseWebIDLListener) ExitPartialInterfaceMember(ctx *PartialInterfaceMemberContext) {}

// EnterInheritance is called when production inheritance is entered.
func (s *BaseWebIDLListener) EnterInheritance(ctx *InheritanceContext) {}

// ExitInheritance is called when production inheritance is exited.
func (s *BaseWebIDLListener) ExitInheritance(ctx *InheritanceContext) {}

// EnterMixinRest is called when production mixinRest is entered.
func (s *BaseWebIDLListener) EnterMixinRest(ctx *MixinRestContext) {}

// ExitMixinRest is called when production mixinRest is exited.
func (s *BaseWebIDLListener) ExitMixinRest(ctx *MixinRestContext) {}

// EnterMixinMembers is called when production mixinMembers is entered.
func (s *BaseWebIDLListener) EnterMixinMembers(ctx *MixinMembersContext) {}

// ExitMixinMembers is called when production mixinMembers is exited.
func (s *BaseWebIDLListener) ExitMixinMembers(ctx *MixinMembersContext) {}

// EnterMixinMember is called when production mixinMember is entered.
func (s *BaseWebIDLListener) EnterMixinMember(ctx *MixinMemberContext) {}

// ExitMixinMember is called when production mixinMember is exited.
func (s *BaseWebIDLListener) ExitMixinMember(ctx *MixinMemberContext) {}

// EnterIncludesStatement is called when production includesStatement is entered.
func (s *BaseWebIDLListener) EnterIncludesStatement(ctx *IncludesStatementContext) {}

// ExitIncludesStatement is called when production includesStatement is exited.
func (s *BaseWebIDLListener) ExitIncludesStatement(ctx *IncludesStatementContext) {}

// EnterCallbackRestOrInterface is called when production callbackRestOrInterface is entered.
func (s *BaseWebIDLListener) EnterCallbackRestOrInterface(ctx *CallbackRestOrInterfaceContext) {}

// ExitCallbackRestOrInterface is called when production callbackRestOrInterface is exited.
func (s *BaseWebIDLListener) ExitCallbackRestOrInterface(ctx *CallbackRestOrInterfaceContext) {}

// EnterCallbackInterfaceMembers is called when production callbackInterfaceMembers is entered.
func (s *BaseWebIDLListener) EnterCallbackInterfaceMembers(ctx *CallbackInterfaceMembersContext) {}

// ExitCallbackInterfaceMembers is called when production callbackInterfaceMembers is exited.
func (s *BaseWebIDLListener) ExitCallbackInterfaceMembers(ctx *CallbackInterfaceMembersContext) {}

// EnterCallbackInterfaceMember is called when production callbackInterfaceMember is entered.
func (s *BaseWebIDLListener) EnterCallbackInterfaceMember(ctx *CallbackInterfaceMemberContext) {}

// ExitCallbackInterfaceMember is called when production callbackInterfaceMember is exited.
func (s *BaseWebIDLListener) ExitCallbackInterfaceMember(ctx *CallbackInterfaceMemberContext) {}

// EnterConst_ is called when production const_ is entered.
func (s *BaseWebIDLListener) EnterConst_(ctx *Const_Context) {}

// ExitConst_ is called when production const_ is exited.
func (s *BaseWebIDLListener) ExitConst_(ctx *Const_Context) {}

// EnterConstValue is called when production constValue is entered.
func (s *BaseWebIDLListener) EnterConstValue(ctx *ConstValueContext) {}

// ExitConstValue is called when production constValue is exited.
func (s *BaseWebIDLListener) ExitConstValue(ctx *ConstValueContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseWebIDLListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseWebIDLListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseWebIDLListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseWebIDLListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterConstType is called when production constType is entered.
func (s *BaseWebIDLListener) EnterConstType(ctx *ConstTypeContext) {}

// ExitConstType is called when production constType is exited.
func (s *BaseWebIDLListener) ExitConstType(ctx *ConstTypeContext) {}

// EnterReadonlyMember is called when production readonlyMember is entered.
func (s *BaseWebIDLListener) EnterReadonlyMember(ctx *ReadonlyMemberContext) {}

// ExitReadonlyMember is called when production readonlyMember is exited.
func (s *BaseWebIDLListener) ExitReadonlyMember(ctx *ReadonlyMemberContext) {}

// EnterReadonlyMemberRest is called when production readonlyMemberRest is entered.
func (s *BaseWebIDLListener) EnterReadonlyMemberRest(ctx *ReadonlyMemberRestContext) {}

// ExitReadonlyMemberRest is called when production readonlyMemberRest is exited.
func (s *BaseWebIDLListener) ExitReadonlyMemberRest(ctx *ReadonlyMemberRestContext) {}

// EnterReadWriteAttribute is called when production readWriteAttribute is entered.
func (s *BaseWebIDLListener) EnterReadWriteAttribute(ctx *ReadWriteAttributeContext) {}

// ExitReadWriteAttribute is called when production readWriteAttribute is exited.
func (s *BaseWebIDLListener) ExitReadWriteAttribute(ctx *ReadWriteAttributeContext) {}

// EnterInheritAttribute is called when production inheritAttribute is entered.
func (s *BaseWebIDLListener) EnterInheritAttribute(ctx *InheritAttributeContext) {}

// ExitInheritAttribute is called when production inheritAttribute is exited.
func (s *BaseWebIDLListener) ExitInheritAttribute(ctx *InheritAttributeContext) {}

// EnterAttributeRest is called when production attributeRest is entered.
func (s *BaseWebIDLListener) EnterAttributeRest(ctx *AttributeRestContext) {}

// ExitAttributeRest is called when production attributeRest is exited.
func (s *BaseWebIDLListener) ExitAttributeRest(ctx *AttributeRestContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseWebIDLListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseWebIDLListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterAttributeNameKeyword is called when production attributeNameKeyword is entered.
func (s *BaseWebIDLListener) EnterAttributeNameKeyword(ctx *AttributeNameKeywordContext) {}

// ExitAttributeNameKeyword is called when production attributeNameKeyword is exited.
func (s *BaseWebIDLListener) ExitAttributeNameKeyword(ctx *AttributeNameKeywordContext) {}

// EnterOptionalReadOnly is called when production optionalReadOnly is entered.
func (s *BaseWebIDLListener) EnterOptionalReadOnly(ctx *OptionalReadOnlyContext) {}

// ExitOptionalReadOnly is called when production optionalReadOnly is exited.
func (s *BaseWebIDLListener) ExitOptionalReadOnly(ctx *OptionalReadOnlyContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseWebIDLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseWebIDLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseWebIDLListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseWebIDLListener) ExitOperation(ctx *OperationContext) {}

// EnterRegularOperation is called when production regularOperation is entered.
func (s *BaseWebIDLListener) EnterRegularOperation(ctx *RegularOperationContext) {}

// ExitRegularOperation is called when production regularOperation is exited.
func (s *BaseWebIDLListener) ExitRegularOperation(ctx *RegularOperationContext) {}

// EnterSpecialOperation is called when production specialOperation is entered.
func (s *BaseWebIDLListener) EnterSpecialOperation(ctx *SpecialOperationContext) {}

// ExitSpecialOperation is called when production specialOperation is exited.
func (s *BaseWebIDLListener) ExitSpecialOperation(ctx *SpecialOperationContext) {}

// EnterSpecial is called when production special is entered.
func (s *BaseWebIDLListener) EnterSpecial(ctx *SpecialContext) {}

// ExitSpecial is called when production special is exited.
func (s *BaseWebIDLListener) ExitSpecial(ctx *SpecialContext) {}

// EnterOperationRest is called when production operationRest is entered.
func (s *BaseWebIDLListener) EnterOperationRest(ctx *OperationRestContext) {}

// ExitOperationRest is called when production operationRest is exited.
func (s *BaseWebIDLListener) ExitOperationRest(ctx *OperationRestContext) {}

// EnterOptionalOperationName is called when production optionalOperationName is entered.
func (s *BaseWebIDLListener) EnterOptionalOperationName(ctx *OptionalOperationNameContext) {}

// ExitOptionalOperationName is called when production optionalOperationName is exited.
func (s *BaseWebIDLListener) ExitOptionalOperationName(ctx *OptionalOperationNameContext) {}

// EnterOperationName is called when production operationName is entered.
func (s *BaseWebIDLListener) EnterOperationName(ctx *OperationNameContext) {}

// ExitOperationName is called when production operationName is exited.
func (s *BaseWebIDLListener) ExitOperationName(ctx *OperationNameContext) {}

// EnterOperationNameKeyword is called when production operationNameKeyword is entered.
func (s *BaseWebIDLListener) EnterOperationNameKeyword(ctx *OperationNameKeywordContext) {}

// ExitOperationNameKeyword is called when production operationNameKeyword is exited.
func (s *BaseWebIDLListener) ExitOperationNameKeyword(ctx *OperationNameKeywordContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseWebIDLListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseWebIDLListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseWebIDLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseWebIDLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseWebIDLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseWebIDLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterArgumentRest is called when production argumentRest is entered.
func (s *BaseWebIDLListener) EnterArgumentRest(ctx *ArgumentRestContext) {}

// ExitArgumentRest is called when production argumentRest is exited.
func (s *BaseWebIDLListener) ExitArgumentRest(ctx *ArgumentRestContext) {}

// EnterArgumentName is called when production argumentName is entered.
func (s *BaseWebIDLListener) EnterArgumentName(ctx *ArgumentNameContext) {}

// ExitArgumentName is called when production argumentName is exited.
func (s *BaseWebIDLListener) ExitArgumentName(ctx *ArgumentNameContext) {}

// EnterEllipsis is called when production ellipsis is entered.
func (s *BaseWebIDLListener) EnterEllipsis(ctx *EllipsisContext) {}

// ExitEllipsis is called when production ellipsis is exited.
func (s *BaseWebIDLListener) ExitEllipsis(ctx *EllipsisContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BaseWebIDLListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BaseWebIDLListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterStringifier is called when production stringifier is entered.
func (s *BaseWebIDLListener) EnterStringifier(ctx *StringifierContext) {}

// ExitStringifier is called when production stringifier is exited.
func (s *BaseWebIDLListener) ExitStringifier(ctx *StringifierContext) {}

// EnterStringifierRest is called when production stringifierRest is entered.
func (s *BaseWebIDLListener) EnterStringifierRest(ctx *StringifierRestContext) {}

// ExitStringifierRest is called when production stringifierRest is exited.
func (s *BaseWebIDLListener) ExitStringifierRest(ctx *StringifierRestContext) {}

// EnterStaticMember is called when production staticMember is entered.
func (s *BaseWebIDLListener) EnterStaticMember(ctx *StaticMemberContext) {}

// ExitStaticMember is called when production staticMember is exited.
func (s *BaseWebIDLListener) ExitStaticMember(ctx *StaticMemberContext) {}

// EnterStaticMemberRest is called when production staticMemberRest is entered.
func (s *BaseWebIDLListener) EnterStaticMemberRest(ctx *StaticMemberRestContext) {}

// ExitStaticMemberRest is called when production staticMemberRest is exited.
func (s *BaseWebIDLListener) ExitStaticMemberRest(ctx *StaticMemberRestContext) {}

// EnterIterable is called when production iterable is entered.
func (s *BaseWebIDLListener) EnterIterable(ctx *IterableContext) {}

// ExitIterable is called when production iterable is exited.
func (s *BaseWebIDLListener) ExitIterable(ctx *IterableContext) {}

// EnterOptionalType is called when production optionalType is entered.
func (s *BaseWebIDLListener) EnterOptionalType(ctx *OptionalTypeContext) {}

// ExitOptionalType is called when production optionalType is exited.
func (s *BaseWebIDLListener) ExitOptionalType(ctx *OptionalTypeContext) {}

// EnterAsyncIterable is called when production asyncIterable is entered.
func (s *BaseWebIDLListener) EnterAsyncIterable(ctx *AsyncIterableContext) {}

// ExitAsyncIterable is called when production asyncIterable is exited.
func (s *BaseWebIDLListener) ExitAsyncIterable(ctx *AsyncIterableContext) {}

// EnterOptionalArgumentList is called when production optionalArgumentList is entered.
func (s *BaseWebIDLListener) EnterOptionalArgumentList(ctx *OptionalArgumentListContext) {}

// ExitOptionalArgumentList is called when production optionalArgumentList is exited.
func (s *BaseWebIDLListener) ExitOptionalArgumentList(ctx *OptionalArgumentListContext) {}

// EnterReadWriteMaplike is called when production readWriteMaplike is entered.
func (s *BaseWebIDLListener) EnterReadWriteMaplike(ctx *ReadWriteMaplikeContext) {}

// ExitReadWriteMaplike is called when production readWriteMaplike is exited.
func (s *BaseWebIDLListener) ExitReadWriteMaplike(ctx *ReadWriteMaplikeContext) {}

// EnterMaplikeRest is called when production maplikeRest is entered.
func (s *BaseWebIDLListener) EnterMaplikeRest(ctx *MaplikeRestContext) {}

// ExitMaplikeRest is called when production maplikeRest is exited.
func (s *BaseWebIDLListener) ExitMaplikeRest(ctx *MaplikeRestContext) {}

// EnterReadWriteSetlike is called when production readWriteSetlike is entered.
func (s *BaseWebIDLListener) EnterReadWriteSetlike(ctx *ReadWriteSetlikeContext) {}

// ExitReadWriteSetlike is called when production readWriteSetlike is exited.
func (s *BaseWebIDLListener) ExitReadWriteSetlike(ctx *ReadWriteSetlikeContext) {}

// EnterSetlikeRest is called when production setlikeRest is entered.
func (s *BaseWebIDLListener) EnterSetlikeRest(ctx *SetlikeRestContext) {}

// ExitSetlikeRest is called when production setlikeRest is exited.
func (s *BaseWebIDLListener) ExitSetlikeRest(ctx *SetlikeRestContext) {}

// EnterNamespace_ is called when production namespace_ is entered.
func (s *BaseWebIDLListener) EnterNamespace_(ctx *Namespace_Context) {}

// ExitNamespace_ is called when production namespace_ is exited.
func (s *BaseWebIDLListener) ExitNamespace_(ctx *Namespace_Context) {}

// EnterNamespaceMembers is called when production namespaceMembers is entered.
func (s *BaseWebIDLListener) EnterNamespaceMembers(ctx *NamespaceMembersContext) {}

// ExitNamespaceMembers is called when production namespaceMembers is exited.
func (s *BaseWebIDLListener) ExitNamespaceMembers(ctx *NamespaceMembersContext) {}

// EnterNamespaceMember is called when production namespaceMember is entered.
func (s *BaseWebIDLListener) EnterNamespaceMember(ctx *NamespaceMemberContext) {}

// ExitNamespaceMember is called when production namespaceMember is exited.
func (s *BaseWebIDLListener) ExitNamespaceMember(ctx *NamespaceMemberContext) {}

// EnterDictionary is called when production dictionary is entered.
func (s *BaseWebIDLListener) EnterDictionary(ctx *DictionaryContext) {}

// ExitDictionary is called when production dictionary is exited.
func (s *BaseWebIDLListener) ExitDictionary(ctx *DictionaryContext) {}

// EnterDictionaryMembers is called when production dictionaryMembers is entered.
func (s *BaseWebIDLListener) EnterDictionaryMembers(ctx *DictionaryMembersContext) {}

// ExitDictionaryMembers is called when production dictionaryMembers is exited.
func (s *BaseWebIDLListener) ExitDictionaryMembers(ctx *DictionaryMembersContext) {}

// EnterDictionaryMember is called when production dictionaryMember is entered.
func (s *BaseWebIDLListener) EnterDictionaryMember(ctx *DictionaryMemberContext) {}

// ExitDictionaryMember is called when production dictionaryMember is exited.
func (s *BaseWebIDLListener) ExitDictionaryMember(ctx *DictionaryMemberContext) {}

// EnterDictionaryMemberRest is called when production dictionaryMemberRest is entered.
func (s *BaseWebIDLListener) EnterDictionaryMemberRest(ctx *DictionaryMemberRestContext) {}

// ExitDictionaryMemberRest is called when production dictionaryMemberRest is exited.
func (s *BaseWebIDLListener) ExitDictionaryMemberRest(ctx *DictionaryMemberRestContext) {}

// EnterPartialDictionary is called when production partialDictionary is entered.
func (s *BaseWebIDLListener) EnterPartialDictionary(ctx *PartialDictionaryContext) {}

// ExitPartialDictionary is called when production partialDictionary is exited.
func (s *BaseWebIDLListener) ExitPartialDictionary(ctx *PartialDictionaryContext) {}

// EnterDefault_ is called when production default_ is entered.
func (s *BaseWebIDLListener) EnterDefault_(ctx *Default_Context) {}

// ExitDefault_ is called when production default_ is exited.
func (s *BaseWebIDLListener) ExitDefault_(ctx *Default_Context) {}

// EnterEnum_ is called when production enum_ is entered.
func (s *BaseWebIDLListener) EnterEnum_(ctx *Enum_Context) {}

// ExitEnum_ is called when production enum_ is exited.
func (s *BaseWebIDLListener) ExitEnum_(ctx *Enum_Context) {}

// EnterEnumValueList is called when production enumValueList is entered.
func (s *BaseWebIDLListener) EnterEnumValueList(ctx *EnumValueListContext) {}

// ExitEnumValueList is called when production enumValueList is exited.
func (s *BaseWebIDLListener) ExitEnumValueList(ctx *EnumValueListContext) {}

// EnterEnumValueListComma is called when production enumValueListComma is entered.
func (s *BaseWebIDLListener) EnterEnumValueListComma(ctx *EnumValueListCommaContext) {}

// ExitEnumValueListComma is called when production enumValueListComma is exited.
func (s *BaseWebIDLListener) ExitEnumValueListComma(ctx *EnumValueListCommaContext) {}

// EnterEnumValueListString is called when production enumValueListString is entered.
func (s *BaseWebIDLListener) EnterEnumValueListString(ctx *EnumValueListStringContext) {}

// ExitEnumValueListString is called when production enumValueListString is exited.
func (s *BaseWebIDLListener) ExitEnumValueListString(ctx *EnumValueListStringContext) {}

// EnterCallbackRest is called when production callbackRest is entered.
func (s *BaseWebIDLListener) EnterCallbackRest(ctx *CallbackRestContext) {}

// ExitCallbackRest is called when production callbackRest is exited.
func (s *BaseWebIDLListener) ExitCallbackRest(ctx *CallbackRestContext) {}

// EnterTypedef_ is called when production typedef_ is entered.
func (s *BaseWebIDLListener) EnterTypedef_(ctx *Typedef_Context) {}

// ExitTypedef_ is called when production typedef_ is exited.
func (s *BaseWebIDLListener) ExitTypedef_(ctx *Typedef_Context) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseWebIDLListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseWebIDLListener) ExitType_(ctx *Type_Context) {}

// EnterTypeWithExtendedAttributes is called when production typeWithExtendedAttributes is entered.
func (s *BaseWebIDLListener) EnterTypeWithExtendedAttributes(ctx *TypeWithExtendedAttributesContext) {
}

// ExitTypeWithExtendedAttributes is called when production typeWithExtendedAttributes is exited.
func (s *BaseWebIDLListener) ExitTypeWithExtendedAttributes(ctx *TypeWithExtendedAttributesContext) {}

// EnterSingleType is called when production singleType is entered.
func (s *BaseWebIDLListener) EnterSingleType(ctx *SingleTypeContext) {}

// ExitSingleType is called when production singleType is exited.
func (s *BaseWebIDLListener) ExitSingleType(ctx *SingleTypeContext) {}

// EnterUnionType is called when production unionType is entered.
func (s *BaseWebIDLListener) EnterUnionType(ctx *UnionTypeContext) {}

// ExitUnionType is called when production unionType is exited.
func (s *BaseWebIDLListener) ExitUnionType(ctx *UnionTypeContext) {}

// EnterUnionMemberType is called when production unionMemberType is entered.
func (s *BaseWebIDLListener) EnterUnionMemberType(ctx *UnionMemberTypeContext) {}

// ExitUnionMemberType is called when production unionMemberType is exited.
func (s *BaseWebIDLListener) ExitUnionMemberType(ctx *UnionMemberTypeContext) {}

// EnterUnionMemberTypes is called when production unionMemberTypes is entered.
func (s *BaseWebIDLListener) EnterUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// ExitUnionMemberTypes is called when production unionMemberTypes is exited.
func (s *BaseWebIDLListener) ExitUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// EnterDistinguishableType is called when production distinguishableType is entered.
func (s *BaseWebIDLListener) EnterDistinguishableType(ctx *DistinguishableTypeContext) {}

// ExitDistinguishableType is called when production distinguishableType is exited.
func (s *BaseWebIDLListener) ExitDistinguishableType(ctx *DistinguishableTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseWebIDLListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseWebIDLListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterUnrestrictedFloatType is called when production unrestrictedFloatType is entered.
func (s *BaseWebIDLListener) EnterUnrestrictedFloatType(ctx *UnrestrictedFloatTypeContext) {}

// ExitUnrestrictedFloatType is called when production unrestrictedFloatType is exited.
func (s *BaseWebIDLListener) ExitUnrestrictedFloatType(ctx *UnrestrictedFloatTypeContext) {}

// EnterFloatType is called when production floatType is entered.
func (s *BaseWebIDLListener) EnterFloatType(ctx *FloatTypeContext) {}

// ExitFloatType is called when production floatType is exited.
func (s *BaseWebIDLListener) ExitFloatType(ctx *FloatTypeContext) {}

// EnterUnsignedIntegerType is called when production unsignedIntegerType is entered.
func (s *BaseWebIDLListener) EnterUnsignedIntegerType(ctx *UnsignedIntegerTypeContext) {}

// ExitUnsignedIntegerType is called when production unsignedIntegerType is exited.
func (s *BaseWebIDLListener) ExitUnsignedIntegerType(ctx *UnsignedIntegerTypeContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseWebIDLListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseWebIDLListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterOptionalLong is called when production optionalLong is entered.
func (s *BaseWebIDLListener) EnterOptionalLong(ctx *OptionalLongContext) {}

// ExitOptionalLong is called when production optionalLong is exited.
func (s *BaseWebIDLListener) ExitOptionalLong(ctx *OptionalLongContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseWebIDLListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseWebIDLListener) ExitStringType(ctx *StringTypeContext) {}

// EnterPromiseType is called when production promiseType is entered.
func (s *BaseWebIDLListener) EnterPromiseType(ctx *PromiseTypeContext) {}

// ExitPromiseType is called when production promiseType is exited.
func (s *BaseWebIDLListener) ExitPromiseType(ctx *PromiseTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BaseWebIDLListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BaseWebIDLListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterNull_ is called when production null_ is entered.
func (s *BaseWebIDLListener) EnterNull_(ctx *Null_Context) {}

// ExitNull_ is called when production null_ is exited.
func (s *BaseWebIDLListener) ExitNull_(ctx *Null_Context) {}

// EnterBufferRelatedType is called when production bufferRelatedType is entered.
func (s *BaseWebIDLListener) EnterBufferRelatedType(ctx *BufferRelatedTypeContext) {}

// ExitBufferRelatedType is called when production bufferRelatedType is exited.
func (s *BaseWebIDLListener) ExitBufferRelatedType(ctx *BufferRelatedTypeContext) {}

// EnterExtendedAttributeList is called when production extendedAttributeList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeList(ctx *ExtendedAttributeListContext) {}

// ExitExtendedAttributeList is called when production extendedAttributeList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeList(ctx *ExtendedAttributeListContext) {}

// EnterExtendedAttributes is called when production extendedAttributes is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributes(ctx *ExtendedAttributesContext) {}

// ExitExtendedAttributes is called when production extendedAttributes is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributes(ctx *ExtendedAttributesContext) {}

// EnterExtendedAttribute is called when production extendedAttribute is entered.
func (s *BaseWebIDLListener) EnterExtendedAttribute(ctx *ExtendedAttributeContext) {}

// ExitExtendedAttribute is called when production extendedAttribute is exited.
func (s *BaseWebIDLListener) ExitExtendedAttribute(ctx *ExtendedAttributeContext) {}

// EnterOther is called when production other is entered.
func (s *BaseWebIDLListener) EnterOther(ctx *OtherContext) {}

// ExitOther is called when production other is exited.
func (s *BaseWebIDLListener) ExitOther(ctx *OtherContext) {}

// EnterOtherOrComma is called when production otherOrComma is entered.
func (s *BaseWebIDLListener) EnterOtherOrComma(ctx *OtherOrCommaContext) {}

// ExitOtherOrComma is called when production otherOrComma is exited.
func (s *BaseWebIDLListener) ExitOtherOrComma(ctx *OtherOrCommaContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseWebIDLListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseWebIDLListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BaseWebIDLListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BaseWebIDLListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterExtendedAttributeNoArgs is called when production extendedAttributeNoArgs is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeNoArgs(ctx *ExtendedAttributeNoArgsContext) {}

// ExitExtendedAttributeNoArgs is called when production extendedAttributeNoArgs is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeNoArgs(ctx *ExtendedAttributeNoArgsContext) {}

// EnterExtendedAttributeArgList is called when production extendedAttributeArgList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeArgList(ctx *ExtendedAttributeArgListContext) {}

// ExitExtendedAttributeArgList is called when production extendedAttributeArgList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeArgList(ctx *ExtendedAttributeArgListContext) {}

// EnterExtendedAttributeIdent is called when production extendedAttributeIdent is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeIdent(ctx *ExtendedAttributeIdentContext) {}

// ExitExtendedAttributeIdent is called when production extendedAttributeIdent is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeIdent(ctx *ExtendedAttributeIdentContext) {}

// EnterExtendedAttributeIdentList is called when production extendedAttributeIdentList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeIdentList(ctx *ExtendedAttributeIdentListContext) {
}

// ExitExtendedAttributeIdentList is called when production extendedAttributeIdentList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeIdentList(ctx *ExtendedAttributeIdentListContext) {}

// EnterExtendedAttributeNamedArgList is called when production extendedAttributeNamedArgList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeNamedArgList(ctx *ExtendedAttributeNamedArgListContext) {
}

// ExitExtendedAttributeNamedArgList is called when production extendedAttributeNamedArgList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeNamedArgList(ctx *ExtendedAttributeNamedArgListContext) {
}

// EnterExtendedAttributeString is called when production extendedAttributeString is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeString(ctx *ExtendedAttributeStringContext) {}

// ExitExtendedAttributeString is called when production extendedAttributeString is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeString(ctx *ExtendedAttributeStringContext) {}

// EnterExtendedAttributeStringList is called when production extendedAttributeStringList is entered.
func (s *BaseWebIDLListener) EnterExtendedAttributeStringList(ctx *ExtendedAttributeStringListContext) {
}

// ExitExtendedAttributeStringList is called when production extendedAttributeStringList is exited.
func (s *BaseWebIDLListener) ExitExtendedAttributeStringList(ctx *ExtendedAttributeStringListContext) {
}

// EnterStringList is called when production stringList is entered.
func (s *BaseWebIDLListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseWebIDLListener) ExitStringList(ctx *StringListContext) {}

// EnterStrings is called when production strings is entered.
func (s *BaseWebIDLListener) EnterStrings(ctx *StringsContext) {}

// ExitStrings is called when production strings is exited.
func (s *BaseWebIDLListener) ExitStrings(ctx *StringsContext) {}
