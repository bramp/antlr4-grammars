// Code generated from CapnProto.g4 by ANTLR 4.9.3. DO NOT EDIT.

package capnproto // CapnProto
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CapnProtoListener is a complete listener for a parse tree produced by CapnProtoParser.
type CapnProtoListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterFile_identifier is called when entering the file_identifier production.
	EnterFile_identifier(c *File_identifierContext)

	// EnterUsing_import is called when entering the using_import production.
	EnterUsing_import(c *Using_importContext)

	// EnterNamespace_ is called when entering the namespace_ production.
	EnterNamespace_(c *Namespace_Context)

	// EnterDocument_content is called when entering the document_content production.
	EnterDocument_content(c *Document_contentContext)

	// EnterStruct_def is called when entering the struct_def production.
	EnterStruct_def(c *Struct_defContext)

	// EnterStruct_content is called when entering the struct_content production.
	EnterStruct_content(c *Struct_contentContext)

	// EnterInterface_def is called when entering the interface_def production.
	EnterInterface_def(c *Interface_defContext)

	// EnterInterface_content is called when entering the interface_content production.
	EnterInterface_content(c *Interface_contentContext)

	// EnterField_def is called when entering the field_def production.
	EnterField_def(c *Field_defContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterInner_type is called when entering the inner_type production.
	EnterInner_type(c *Inner_typeContext)

	// EnterEnum_def is called when entering the enum_def production.
	EnterEnum_def(c *Enum_defContext)

	// EnterAnnotation_reference is called when entering the annotation_reference production.
	EnterAnnotation_reference(c *Annotation_referenceContext)

	// EnterEnum_content is called when entering the enum_content production.
	EnterEnum_content(c *Enum_contentContext)

	// EnterNamed_union_def is called when entering the named_union_def production.
	EnterNamed_union_def(c *Named_union_defContext)

	// EnterUnnamed_union_def is called when entering the unnamed_union_def production.
	EnterUnnamed_union_def(c *Unnamed_union_defContext)

	// EnterUnion_content is called when entering the union_content production.
	EnterUnion_content(c *Union_contentContext)

	// EnterGroup_def is called when entering the group_def production.
	EnterGroup_def(c *Group_defContext)

	// EnterGroup_content is called when entering the group_content production.
	EnterGroup_content(c *Group_contentContext)

	// EnterFunction_def is called when entering the function_def production.
	EnterFunction_def(c *Function_defContext)

	// EnterGeneric_type_parameters is called when entering the generic_type_parameters production.
	EnterGeneric_type_parameters(c *Generic_type_parametersContext)

	// EnterFunction_parameters is called when entering the function_parameters production.
	EnterFunction_parameters(c *Function_parametersContext)

	// EnterAnnotation_def is called when entering the annotation_def production.
	EnterAnnotation_def(c *Annotation_defContext)

	// EnterAnnotation_parameters is called when entering the annotation_parameters production.
	EnterAnnotation_parameters(c *Annotation_parametersContext)

	// EnterConst_def is called when entering the const_def production.
	EnterConst_def(c *Const_defContext)

	// EnterConst_value is called when entering the const_value production.
	EnterConst_value(c *Const_valueContext)

	// EnterLiteral_union is called when entering the literal_union production.
	EnterLiteral_union(c *Literal_unionContext)

	// EnterLiteral_list is called when entering the literal_list production.
	EnterLiteral_list(c *Literal_listContext)

	// EnterLiteral_bytes is called when entering the literal_bytes production.
	EnterLiteral_bytes(c *Literal_bytesContext)

	// EnterUnion_mapping is called when entering the union_mapping production.
	EnterUnion_mapping(c *Union_mappingContext)

	// EnterInner_using is called when entering the inner_using production.
	EnterInner_using(c *Inner_usingContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitFile_identifier is called when exiting the file_identifier production.
	ExitFile_identifier(c *File_identifierContext)

	// ExitUsing_import is called when exiting the using_import production.
	ExitUsing_import(c *Using_importContext)

	// ExitNamespace_ is called when exiting the namespace_ production.
	ExitNamespace_(c *Namespace_Context)

	// ExitDocument_content is called when exiting the document_content production.
	ExitDocument_content(c *Document_contentContext)

	// ExitStruct_def is called when exiting the struct_def production.
	ExitStruct_def(c *Struct_defContext)

	// ExitStruct_content is called when exiting the struct_content production.
	ExitStruct_content(c *Struct_contentContext)

	// ExitInterface_def is called when exiting the interface_def production.
	ExitInterface_def(c *Interface_defContext)

	// ExitInterface_content is called when exiting the interface_content production.
	ExitInterface_content(c *Interface_contentContext)

	// ExitField_def is called when exiting the field_def production.
	ExitField_def(c *Field_defContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitInner_type is called when exiting the inner_type production.
	ExitInner_type(c *Inner_typeContext)

	// ExitEnum_def is called when exiting the enum_def production.
	ExitEnum_def(c *Enum_defContext)

	// ExitAnnotation_reference is called when exiting the annotation_reference production.
	ExitAnnotation_reference(c *Annotation_referenceContext)

	// ExitEnum_content is called when exiting the enum_content production.
	ExitEnum_content(c *Enum_contentContext)

	// ExitNamed_union_def is called when exiting the named_union_def production.
	ExitNamed_union_def(c *Named_union_defContext)

	// ExitUnnamed_union_def is called when exiting the unnamed_union_def production.
	ExitUnnamed_union_def(c *Unnamed_union_defContext)

	// ExitUnion_content is called when exiting the union_content production.
	ExitUnion_content(c *Union_contentContext)

	// ExitGroup_def is called when exiting the group_def production.
	ExitGroup_def(c *Group_defContext)

	// ExitGroup_content is called when exiting the group_content production.
	ExitGroup_content(c *Group_contentContext)

	// ExitFunction_def is called when exiting the function_def production.
	ExitFunction_def(c *Function_defContext)

	// ExitGeneric_type_parameters is called when exiting the generic_type_parameters production.
	ExitGeneric_type_parameters(c *Generic_type_parametersContext)

	// ExitFunction_parameters is called when exiting the function_parameters production.
	ExitFunction_parameters(c *Function_parametersContext)

	// ExitAnnotation_def is called when exiting the annotation_def production.
	ExitAnnotation_def(c *Annotation_defContext)

	// ExitAnnotation_parameters is called when exiting the annotation_parameters production.
	ExitAnnotation_parameters(c *Annotation_parametersContext)

	// ExitConst_def is called when exiting the const_def production.
	ExitConst_def(c *Const_defContext)

	// ExitConst_value is called when exiting the const_value production.
	ExitConst_value(c *Const_valueContext)

	// ExitLiteral_union is called when exiting the literal_union production.
	ExitLiteral_union(c *Literal_unionContext)

	// ExitLiteral_list is called when exiting the literal_list production.
	ExitLiteral_list(c *Literal_listContext)

	// ExitLiteral_bytes is called when exiting the literal_bytes production.
	ExitLiteral_bytes(c *Literal_bytesContext)

	// ExitUnion_mapping is called when exiting the union_mapping production.
	ExitUnion_mapping(c *Union_mappingContext)

	// ExitInner_using is called when exiting the inner_using production.
	ExitInner_using(c *Inner_usingContext)
}
