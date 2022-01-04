// Code generated from FlatBuffers.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flatbuffers // FlatBuffers
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFlatBuffersListener is a complete listener for a parse tree produced by FlatBuffersParser.
type BaseFlatBuffersListener struct{}

var _ FlatBuffersListener = &BaseFlatBuffersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlatBuffersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlatBuffersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlatBuffersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlatBuffersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseFlatBuffersListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseFlatBuffersListener) ExitSchema(ctx *SchemaContext) {}

// EnterInclude_ is called when production include_ is entered.
func (s *BaseFlatBuffersListener) EnterInclude_(ctx *Include_Context) {}

// ExitInclude_ is called when production include_ is exited.
func (s *BaseFlatBuffersListener) ExitInclude_(ctx *Include_Context) {}

// EnterNamespace_decl is called when production namespace_decl is entered.
func (s *BaseFlatBuffersListener) EnterNamespace_decl(ctx *Namespace_declContext) {}

// ExitNamespace_decl is called when production namespace_decl is exited.
func (s *BaseFlatBuffersListener) ExitNamespace_decl(ctx *Namespace_declContext) {}

// EnterAttribute_decl is called when production attribute_decl is entered.
func (s *BaseFlatBuffersListener) EnterAttribute_decl(ctx *Attribute_declContext) {}

// ExitAttribute_decl is called when production attribute_decl is exited.
func (s *BaseFlatBuffersListener) ExitAttribute_decl(ctx *Attribute_declContext) {}

// EnterType_decl is called when production type_decl is entered.
func (s *BaseFlatBuffersListener) EnterType_decl(ctx *Type_declContext) {}

// ExitType_decl is called when production type_decl is exited.
func (s *BaseFlatBuffersListener) ExitType_decl(ctx *Type_declContext) {}

// EnterEnum_decl is called when production enum_decl is entered.
func (s *BaseFlatBuffersListener) EnterEnum_decl(ctx *Enum_declContext) {}

// ExitEnum_decl is called when production enum_decl is exited.
func (s *BaseFlatBuffersListener) ExitEnum_decl(ctx *Enum_declContext) {}

// EnterUnion_decl is called when production union_decl is entered.
func (s *BaseFlatBuffersListener) EnterUnion_decl(ctx *Union_declContext) {}

// ExitUnion_decl is called when production union_decl is exited.
func (s *BaseFlatBuffersListener) ExitUnion_decl(ctx *Union_declContext) {}

// EnterRoot_decl is called when production root_decl is entered.
func (s *BaseFlatBuffersListener) EnterRoot_decl(ctx *Root_declContext) {}

// ExitRoot_decl is called when production root_decl is exited.
func (s *BaseFlatBuffersListener) ExitRoot_decl(ctx *Root_declContext) {}

// EnterField_decl is called when production field_decl is entered.
func (s *BaseFlatBuffersListener) EnterField_decl(ctx *Field_declContext) {}

// ExitField_decl is called when production field_decl is exited.
func (s *BaseFlatBuffersListener) ExitField_decl(ctx *Field_declContext) {}

// EnterRpc_decl is called when production rpc_decl is entered.
func (s *BaseFlatBuffersListener) EnterRpc_decl(ctx *Rpc_declContext) {}

// ExitRpc_decl is called when production rpc_decl is exited.
func (s *BaseFlatBuffersListener) ExitRpc_decl(ctx *Rpc_declContext) {}

// EnterRpc_method is called when production rpc_method is entered.
func (s *BaseFlatBuffersListener) EnterRpc_method(ctx *Rpc_methodContext) {}

// ExitRpc_method is called when production rpc_method is exited.
func (s *BaseFlatBuffersListener) ExitRpc_method(ctx *Rpc_methodContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseFlatBuffersListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseFlatBuffersListener) ExitType_(ctx *Type_Context) {}

// EnterEnumval_decl is called when production enumval_decl is entered.
func (s *BaseFlatBuffersListener) EnterEnumval_decl(ctx *Enumval_declContext) {}

// ExitEnumval_decl is called when production enumval_decl is exited.
func (s *BaseFlatBuffersListener) ExitEnumval_decl(ctx *Enumval_declContext) {}

// EnterCommasep_enumval_decl is called when production commasep_enumval_decl is entered.
func (s *BaseFlatBuffersListener) EnterCommasep_enumval_decl(ctx *Commasep_enumval_declContext) {}

// ExitCommasep_enumval_decl is called when production commasep_enumval_decl is exited.
func (s *BaseFlatBuffersListener) ExitCommasep_enumval_decl(ctx *Commasep_enumval_declContext) {}

// EnterUnionval_with_opt_alias is called when production unionval_with_opt_alias is entered.
func (s *BaseFlatBuffersListener) EnterUnionval_with_opt_alias(ctx *Unionval_with_opt_aliasContext) {}

// ExitUnionval_with_opt_alias is called when production unionval_with_opt_alias is exited.
func (s *BaseFlatBuffersListener) ExitUnionval_with_opt_alias(ctx *Unionval_with_opt_aliasContext) {}

// EnterCommasep_unionval_with_opt_alias is called when production commasep_unionval_with_opt_alias is entered.
func (s *BaseFlatBuffersListener) EnterCommasep_unionval_with_opt_alias(ctx *Commasep_unionval_with_opt_aliasContext) {
}

// ExitCommasep_unionval_with_opt_alias is called when production commasep_unionval_with_opt_alias is exited.
func (s *BaseFlatBuffersListener) ExitCommasep_unionval_with_opt_alias(ctx *Commasep_unionval_with_opt_aliasContext) {
}

// EnterIdent_with_opt_single_value is called when production ident_with_opt_single_value is entered.
func (s *BaseFlatBuffersListener) EnterIdent_with_opt_single_value(ctx *Ident_with_opt_single_valueContext) {
}

// ExitIdent_with_opt_single_value is called when production ident_with_opt_single_value is exited.
func (s *BaseFlatBuffersListener) ExitIdent_with_opt_single_value(ctx *Ident_with_opt_single_valueContext) {
}

// EnterCommasep_ident_with_opt_single_value is called when production commasep_ident_with_opt_single_value is entered.
func (s *BaseFlatBuffersListener) EnterCommasep_ident_with_opt_single_value(ctx *Commasep_ident_with_opt_single_valueContext) {
}

// ExitCommasep_ident_with_opt_single_value is called when production commasep_ident_with_opt_single_value is exited.
func (s *BaseFlatBuffersListener) ExitCommasep_ident_with_opt_single_value(ctx *Commasep_ident_with_opt_single_valueContext) {
}

// EnterMetadata is called when production metadata is entered.
func (s *BaseFlatBuffersListener) EnterMetadata(ctx *MetadataContext) {}

// ExitMetadata is called when production metadata is exited.
func (s *BaseFlatBuffersListener) ExitMetadata(ctx *MetadataContext) {}

// EnterScalar is called when production scalar is entered.
func (s *BaseFlatBuffersListener) EnterScalar(ctx *ScalarContext) {}

// ExitScalar is called when production scalar is exited.
func (s *BaseFlatBuffersListener) ExitScalar(ctx *ScalarContext) {}

// EnterObject_ is called when production object_ is entered.
func (s *BaseFlatBuffersListener) EnterObject_(ctx *Object_Context) {}

// ExitObject_ is called when production object_ is exited.
func (s *BaseFlatBuffersListener) ExitObject_(ctx *Object_Context) {}

// EnterIdent_with_value is called when production ident_with_value is entered.
func (s *BaseFlatBuffersListener) EnterIdent_with_value(ctx *Ident_with_valueContext) {}

// ExitIdent_with_value is called when production ident_with_value is exited.
func (s *BaseFlatBuffersListener) ExitIdent_with_value(ctx *Ident_with_valueContext) {}

// EnterCommasep_ident_with_value is called when production commasep_ident_with_value is entered.
func (s *BaseFlatBuffersListener) EnterCommasep_ident_with_value(ctx *Commasep_ident_with_valueContext) {
}

// ExitCommasep_ident_with_value is called when production commasep_ident_with_value is exited.
func (s *BaseFlatBuffersListener) ExitCommasep_ident_with_value(ctx *Commasep_ident_with_valueContext) {
}

// EnterSingle_value is called when production single_value is entered.
func (s *BaseFlatBuffersListener) EnterSingle_value(ctx *Single_valueContext) {}

// ExitSingle_value is called when production single_value is exited.
func (s *BaseFlatBuffersListener) ExitSingle_value(ctx *Single_valueContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFlatBuffersListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFlatBuffersListener) ExitValue(ctx *ValueContext) {}

// EnterCommasep_value is called when production commasep_value is entered.
func (s *BaseFlatBuffersListener) EnterCommasep_value(ctx *Commasep_valueContext) {}

// ExitCommasep_value is called when production commasep_value is exited.
func (s *BaseFlatBuffersListener) ExitCommasep_value(ctx *Commasep_valueContext) {}

// EnterFile_extension_decl is called when production file_extension_decl is entered.
func (s *BaseFlatBuffersListener) EnterFile_extension_decl(ctx *File_extension_declContext) {}

// ExitFile_extension_decl is called when production file_extension_decl is exited.
func (s *BaseFlatBuffersListener) ExitFile_extension_decl(ctx *File_extension_declContext) {}

// EnterFile_identifier_decl is called when production file_identifier_decl is entered.
func (s *BaseFlatBuffersListener) EnterFile_identifier_decl(ctx *File_identifier_declContext) {}

// ExitFile_identifier_decl is called when production file_identifier_decl is exited.
func (s *BaseFlatBuffersListener) ExitFile_identifier_decl(ctx *File_identifier_declContext) {}

// EnterNs_ident is called when production ns_ident is entered.
func (s *BaseFlatBuffersListener) EnterNs_ident(ctx *Ns_identContext) {}

// ExitNs_ident is called when production ns_ident is exited.
func (s *BaseFlatBuffersListener) ExitNs_ident(ctx *Ns_identContext) {}

// EnterInteger_const is called when production integer_const is entered.
func (s *BaseFlatBuffersListener) EnterInteger_const(ctx *Integer_constContext) {}

// ExitInteger_const is called when production integer_const is exited.
func (s *BaseFlatBuffersListener) ExitInteger_const(ctx *Integer_constContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseFlatBuffersListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseFlatBuffersListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseFlatBuffersListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseFlatBuffersListener) ExitKeywords(ctx *KeywordsContext) {}
