// Code generated from FlatBuffers.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flatbuffers // FlatBuffers
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FlatBuffersListener is a complete listener for a parse tree produced by FlatBuffersParser.
type FlatBuffersListener interface {
	antlr.ParseTreeListener

	// EnterSchema is called when entering the schema production.
	EnterSchema(c *SchemaContext)

	// EnterInclude_ is called when entering the include_ production.
	EnterInclude_(c *Include_Context)

	// EnterNamespace_decl is called when entering the namespace_decl production.
	EnterNamespace_decl(c *Namespace_declContext)

	// EnterAttribute_decl is called when entering the attribute_decl production.
	EnterAttribute_decl(c *Attribute_declContext)

	// EnterType_decl is called when entering the type_decl production.
	EnterType_decl(c *Type_declContext)

	// EnterEnum_decl is called when entering the enum_decl production.
	EnterEnum_decl(c *Enum_declContext)

	// EnterUnion_decl is called when entering the union_decl production.
	EnterUnion_decl(c *Union_declContext)

	// EnterRoot_decl is called when entering the root_decl production.
	EnterRoot_decl(c *Root_declContext)

	// EnterField_decl is called when entering the field_decl production.
	EnterField_decl(c *Field_declContext)

	// EnterRpc_decl is called when entering the rpc_decl production.
	EnterRpc_decl(c *Rpc_declContext)

	// EnterRpc_method is called when entering the rpc_method production.
	EnterRpc_method(c *Rpc_methodContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterEnumval_decl is called when entering the enumval_decl production.
	EnterEnumval_decl(c *Enumval_declContext)

	// EnterCommasep_enumval_decl is called when entering the commasep_enumval_decl production.
	EnterCommasep_enumval_decl(c *Commasep_enumval_declContext)

	// EnterUnionval_with_opt_alias is called when entering the unionval_with_opt_alias production.
	EnterUnionval_with_opt_alias(c *Unionval_with_opt_aliasContext)

	// EnterCommasep_unionval_with_opt_alias is called when entering the commasep_unionval_with_opt_alias production.
	EnterCommasep_unionval_with_opt_alias(c *Commasep_unionval_with_opt_aliasContext)

	// EnterIdent_with_opt_single_value is called when entering the ident_with_opt_single_value production.
	EnterIdent_with_opt_single_value(c *Ident_with_opt_single_valueContext)

	// EnterCommasep_ident_with_opt_single_value is called when entering the commasep_ident_with_opt_single_value production.
	EnterCommasep_ident_with_opt_single_value(c *Commasep_ident_with_opt_single_valueContext)

	// EnterMetadata is called when entering the metadata production.
	EnterMetadata(c *MetadataContext)

	// EnterScalar is called when entering the scalar production.
	EnterScalar(c *ScalarContext)

	// EnterObject_ is called when entering the object_ production.
	EnterObject_(c *Object_Context)

	// EnterIdent_with_value is called when entering the ident_with_value production.
	EnterIdent_with_value(c *Ident_with_valueContext)

	// EnterCommasep_ident_with_value is called when entering the commasep_ident_with_value production.
	EnterCommasep_ident_with_value(c *Commasep_ident_with_valueContext)

	// EnterSingle_value is called when entering the single_value production.
	EnterSingle_value(c *Single_valueContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterCommasep_value is called when entering the commasep_value production.
	EnterCommasep_value(c *Commasep_valueContext)

	// EnterFile_extension_decl is called when entering the file_extension_decl production.
	EnterFile_extension_decl(c *File_extension_declContext)

	// EnterFile_identifier_decl is called when entering the file_identifier_decl production.
	EnterFile_identifier_decl(c *File_identifier_declContext)

	// EnterNs_ident is called when entering the ns_ident production.
	EnterNs_ident(c *Ns_identContext)

	// EnterInteger_const is called when entering the integer_const production.
	EnterInteger_const(c *Integer_constContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitSchema is called when exiting the schema production.
	ExitSchema(c *SchemaContext)

	// ExitInclude_ is called when exiting the include_ production.
	ExitInclude_(c *Include_Context)

	// ExitNamespace_decl is called when exiting the namespace_decl production.
	ExitNamespace_decl(c *Namespace_declContext)

	// ExitAttribute_decl is called when exiting the attribute_decl production.
	ExitAttribute_decl(c *Attribute_declContext)

	// ExitType_decl is called when exiting the type_decl production.
	ExitType_decl(c *Type_declContext)

	// ExitEnum_decl is called when exiting the enum_decl production.
	ExitEnum_decl(c *Enum_declContext)

	// ExitUnion_decl is called when exiting the union_decl production.
	ExitUnion_decl(c *Union_declContext)

	// ExitRoot_decl is called when exiting the root_decl production.
	ExitRoot_decl(c *Root_declContext)

	// ExitField_decl is called when exiting the field_decl production.
	ExitField_decl(c *Field_declContext)

	// ExitRpc_decl is called when exiting the rpc_decl production.
	ExitRpc_decl(c *Rpc_declContext)

	// ExitRpc_method is called when exiting the rpc_method production.
	ExitRpc_method(c *Rpc_methodContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitEnumval_decl is called when exiting the enumval_decl production.
	ExitEnumval_decl(c *Enumval_declContext)

	// ExitCommasep_enumval_decl is called when exiting the commasep_enumval_decl production.
	ExitCommasep_enumval_decl(c *Commasep_enumval_declContext)

	// ExitUnionval_with_opt_alias is called when exiting the unionval_with_opt_alias production.
	ExitUnionval_with_opt_alias(c *Unionval_with_opt_aliasContext)

	// ExitCommasep_unionval_with_opt_alias is called when exiting the commasep_unionval_with_opt_alias production.
	ExitCommasep_unionval_with_opt_alias(c *Commasep_unionval_with_opt_aliasContext)

	// ExitIdent_with_opt_single_value is called when exiting the ident_with_opt_single_value production.
	ExitIdent_with_opt_single_value(c *Ident_with_opt_single_valueContext)

	// ExitCommasep_ident_with_opt_single_value is called when exiting the commasep_ident_with_opt_single_value production.
	ExitCommasep_ident_with_opt_single_value(c *Commasep_ident_with_opt_single_valueContext)

	// ExitMetadata is called when exiting the metadata production.
	ExitMetadata(c *MetadataContext)

	// ExitScalar is called when exiting the scalar production.
	ExitScalar(c *ScalarContext)

	// ExitObject_ is called when exiting the object_ production.
	ExitObject_(c *Object_Context)

	// ExitIdent_with_value is called when exiting the ident_with_value production.
	ExitIdent_with_value(c *Ident_with_valueContext)

	// ExitCommasep_ident_with_value is called when exiting the commasep_ident_with_value production.
	ExitCommasep_ident_with_value(c *Commasep_ident_with_valueContext)

	// ExitSingle_value is called when exiting the single_value production.
	ExitSingle_value(c *Single_valueContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitCommasep_value is called when exiting the commasep_value production.
	ExitCommasep_value(c *Commasep_valueContext)

	// ExitFile_extension_decl is called when exiting the file_extension_decl production.
	ExitFile_extension_decl(c *File_extension_declContext)

	// ExitFile_identifier_decl is called when exiting the file_identifier_decl production.
	ExitFile_identifier_decl(c *File_identifier_declContext)

	// ExitNs_ident is called when exiting the ns_ident production.
	ExitNs_ident(c *Ns_identContext)

	// ExitInteger_const is called when exiting the integer_const production.
	ExitInteger_const(c *Integer_constContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
