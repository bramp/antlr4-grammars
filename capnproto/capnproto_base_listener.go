// Code generated from CapnProto.g4 by ANTLR 4.9.3. DO NOT EDIT.

package capnproto // CapnProto
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCapnProtoListener is a complete listener for a parse tree produced by CapnProtoParser.
type BaseCapnProtoListener struct{}

var _ CapnProtoListener = &BaseCapnProtoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCapnProtoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCapnProtoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCapnProtoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCapnProtoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseCapnProtoListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseCapnProtoListener) ExitDocument(ctx *DocumentContext) {}

// EnterFile_identifier is called when production file_identifier is entered.
func (s *BaseCapnProtoListener) EnterFile_identifier(ctx *File_identifierContext) {}

// ExitFile_identifier is called when production file_identifier is exited.
func (s *BaseCapnProtoListener) ExitFile_identifier(ctx *File_identifierContext) {}

// EnterUsing_import is called when production using_import is entered.
func (s *BaseCapnProtoListener) EnterUsing_import(ctx *Using_importContext) {}

// ExitUsing_import is called when production using_import is exited.
func (s *BaseCapnProtoListener) ExitUsing_import(ctx *Using_importContext) {}

// EnterNamespace_ is called when production namespace_ is entered.
func (s *BaseCapnProtoListener) EnterNamespace_(ctx *Namespace_Context) {}

// ExitNamespace_ is called when production namespace_ is exited.
func (s *BaseCapnProtoListener) ExitNamespace_(ctx *Namespace_Context) {}

// EnterDocument_content is called when production document_content is entered.
func (s *BaseCapnProtoListener) EnterDocument_content(ctx *Document_contentContext) {}

// ExitDocument_content is called when production document_content is exited.
func (s *BaseCapnProtoListener) ExitDocument_content(ctx *Document_contentContext) {}

// EnterStruct_def is called when production struct_def is entered.
func (s *BaseCapnProtoListener) EnterStruct_def(ctx *Struct_defContext) {}

// ExitStruct_def is called when production struct_def is exited.
func (s *BaseCapnProtoListener) ExitStruct_def(ctx *Struct_defContext) {}

// EnterStruct_content is called when production struct_content is entered.
func (s *BaseCapnProtoListener) EnterStruct_content(ctx *Struct_contentContext) {}

// ExitStruct_content is called when production struct_content is exited.
func (s *BaseCapnProtoListener) ExitStruct_content(ctx *Struct_contentContext) {}

// EnterInterface_def is called when production interface_def is entered.
func (s *BaseCapnProtoListener) EnterInterface_def(ctx *Interface_defContext) {}

// ExitInterface_def is called when production interface_def is exited.
func (s *BaseCapnProtoListener) ExitInterface_def(ctx *Interface_defContext) {}

// EnterInterface_content is called when production interface_content is entered.
func (s *BaseCapnProtoListener) EnterInterface_content(ctx *Interface_contentContext) {}

// ExitInterface_content is called when production interface_content is exited.
func (s *BaseCapnProtoListener) ExitInterface_content(ctx *Interface_contentContext) {}

// EnterField_def is called when production field_def is entered.
func (s *BaseCapnProtoListener) EnterField_def(ctx *Field_defContext) {}

// ExitField_def is called when production field_def is exited.
func (s *BaseCapnProtoListener) ExitField_def(ctx *Field_defContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseCapnProtoListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseCapnProtoListener) ExitType_(ctx *Type_Context) {}

// EnterInner_type is called when production inner_type is entered.
func (s *BaseCapnProtoListener) EnterInner_type(ctx *Inner_typeContext) {}

// ExitInner_type is called when production inner_type is exited.
func (s *BaseCapnProtoListener) ExitInner_type(ctx *Inner_typeContext) {}

// EnterEnum_def is called when production enum_def is entered.
func (s *BaseCapnProtoListener) EnterEnum_def(ctx *Enum_defContext) {}

// ExitEnum_def is called when production enum_def is exited.
func (s *BaseCapnProtoListener) ExitEnum_def(ctx *Enum_defContext) {}

// EnterAnnotation_reference is called when production annotation_reference is entered.
func (s *BaseCapnProtoListener) EnterAnnotation_reference(ctx *Annotation_referenceContext) {}

// ExitAnnotation_reference is called when production annotation_reference is exited.
func (s *BaseCapnProtoListener) ExitAnnotation_reference(ctx *Annotation_referenceContext) {}

// EnterEnum_content is called when production enum_content is entered.
func (s *BaseCapnProtoListener) EnterEnum_content(ctx *Enum_contentContext) {}

// ExitEnum_content is called when production enum_content is exited.
func (s *BaseCapnProtoListener) ExitEnum_content(ctx *Enum_contentContext) {}

// EnterNamed_union_def is called when production named_union_def is entered.
func (s *BaseCapnProtoListener) EnterNamed_union_def(ctx *Named_union_defContext) {}

// ExitNamed_union_def is called when production named_union_def is exited.
func (s *BaseCapnProtoListener) ExitNamed_union_def(ctx *Named_union_defContext) {}

// EnterUnnamed_union_def is called when production unnamed_union_def is entered.
func (s *BaseCapnProtoListener) EnterUnnamed_union_def(ctx *Unnamed_union_defContext) {}

// ExitUnnamed_union_def is called when production unnamed_union_def is exited.
func (s *BaseCapnProtoListener) ExitUnnamed_union_def(ctx *Unnamed_union_defContext) {}

// EnterUnion_content is called when production union_content is entered.
func (s *BaseCapnProtoListener) EnterUnion_content(ctx *Union_contentContext) {}

// ExitUnion_content is called when production union_content is exited.
func (s *BaseCapnProtoListener) ExitUnion_content(ctx *Union_contentContext) {}

// EnterGroup_def is called when production group_def is entered.
func (s *BaseCapnProtoListener) EnterGroup_def(ctx *Group_defContext) {}

// ExitGroup_def is called when production group_def is exited.
func (s *BaseCapnProtoListener) ExitGroup_def(ctx *Group_defContext) {}

// EnterGroup_content is called when production group_content is entered.
func (s *BaseCapnProtoListener) EnterGroup_content(ctx *Group_contentContext) {}

// ExitGroup_content is called when production group_content is exited.
func (s *BaseCapnProtoListener) ExitGroup_content(ctx *Group_contentContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BaseCapnProtoListener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BaseCapnProtoListener) ExitFunction_def(ctx *Function_defContext) {}

// EnterGeneric_type_parameters is called when production generic_type_parameters is entered.
func (s *BaseCapnProtoListener) EnterGeneric_type_parameters(ctx *Generic_type_parametersContext) {}

// ExitGeneric_type_parameters is called when production generic_type_parameters is exited.
func (s *BaseCapnProtoListener) ExitGeneric_type_parameters(ctx *Generic_type_parametersContext) {}

// EnterFunction_parameters is called when production function_parameters is entered.
func (s *BaseCapnProtoListener) EnterFunction_parameters(ctx *Function_parametersContext) {}

// ExitFunction_parameters is called when production function_parameters is exited.
func (s *BaseCapnProtoListener) ExitFunction_parameters(ctx *Function_parametersContext) {}

// EnterAnnotation_def is called when production annotation_def is entered.
func (s *BaseCapnProtoListener) EnterAnnotation_def(ctx *Annotation_defContext) {}

// ExitAnnotation_def is called when production annotation_def is exited.
func (s *BaseCapnProtoListener) ExitAnnotation_def(ctx *Annotation_defContext) {}

// EnterAnnotation_parameters is called when production annotation_parameters is entered.
func (s *BaseCapnProtoListener) EnterAnnotation_parameters(ctx *Annotation_parametersContext) {}

// ExitAnnotation_parameters is called when production annotation_parameters is exited.
func (s *BaseCapnProtoListener) ExitAnnotation_parameters(ctx *Annotation_parametersContext) {}

// EnterConst_def is called when production const_def is entered.
func (s *BaseCapnProtoListener) EnterConst_def(ctx *Const_defContext) {}

// ExitConst_def is called when production const_def is exited.
func (s *BaseCapnProtoListener) ExitConst_def(ctx *Const_defContext) {}

// EnterConst_value is called when production const_value is entered.
func (s *BaseCapnProtoListener) EnterConst_value(ctx *Const_valueContext) {}

// ExitConst_value is called when production const_value is exited.
func (s *BaseCapnProtoListener) ExitConst_value(ctx *Const_valueContext) {}

// EnterLiteral_union is called when production literal_union is entered.
func (s *BaseCapnProtoListener) EnterLiteral_union(ctx *Literal_unionContext) {}

// ExitLiteral_union is called when production literal_union is exited.
func (s *BaseCapnProtoListener) ExitLiteral_union(ctx *Literal_unionContext) {}

// EnterLiteral_list is called when production literal_list is entered.
func (s *BaseCapnProtoListener) EnterLiteral_list(ctx *Literal_listContext) {}

// ExitLiteral_list is called when production literal_list is exited.
func (s *BaseCapnProtoListener) ExitLiteral_list(ctx *Literal_listContext) {}

// EnterLiteral_bytes is called when production literal_bytes is entered.
func (s *BaseCapnProtoListener) EnterLiteral_bytes(ctx *Literal_bytesContext) {}

// ExitLiteral_bytes is called when production literal_bytes is exited.
func (s *BaseCapnProtoListener) ExitLiteral_bytes(ctx *Literal_bytesContext) {}

// EnterUnion_mapping is called when production union_mapping is entered.
func (s *BaseCapnProtoListener) EnterUnion_mapping(ctx *Union_mappingContext) {}

// ExitUnion_mapping is called when production union_mapping is exited.
func (s *BaseCapnProtoListener) ExitUnion_mapping(ctx *Union_mappingContext) {}

// EnterInner_using is called when production inner_using is entered.
func (s *BaseCapnProtoListener) EnterInner_using(ctx *Inner_usingContext) {}

// ExitInner_using is called when production inner_using is exited.
func (s *BaseCapnProtoListener) ExitInner_using(ctx *Inner_usingContext) {}
