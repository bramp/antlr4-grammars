// Code generated from Thrift.g4 by ANTLR 4.9.3. DO NOT EDIT.

package thrift // Thrift
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ThriftListener is a complete listener for a parse tree produced by ThriftParser.
type ThriftListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterInclude_ is called when entering the include_ production.
	EnterInclude_(c *Include_Context)

	// EnterNamespace_ is called when entering the namespace_ production.
	EnterNamespace_(c *Namespace_Context)

	// EnterCpp_include is called when entering the cpp_include production.
	EnterCpp_include(c *Cpp_includeContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterConst_rule is called when entering the const_rule production.
	EnterConst_rule(c *Const_ruleContext)

	// EnterTypedef_ is called when entering the typedef_ production.
	EnterTypedef_(c *Typedef_Context)

	// EnterEnum_rule is called when entering the enum_rule production.
	EnterEnum_rule(c *Enum_ruleContext)

	// EnterEnum_field is called when entering the enum_field production.
	EnterEnum_field(c *Enum_fieldContext)

	// EnterSenum is called when entering the senum production.
	EnterSenum(c *SenumContext)

	// EnterStruct_ is called when entering the struct_ production.
	EnterStruct_(c *Struct_Context)

	// EnterUnion_ is called when entering the union_ production.
	EnterUnion_(c *Union_Context)

	// EnterException is called when entering the exception production.
	EnterException(c *ExceptionContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterField_id is called when entering the field_id production.
	EnterField_id(c *Field_idContext)

	// EnterField_req is called when entering the field_req production.
	EnterField_req(c *Field_reqContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterOneway is called when entering the oneway production.
	EnterOneway(c *OnewayContext)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterThrows_list is called when entering the throws_list production.
	EnterThrows_list(c *Throws_listContext)

	// EnterType_annotations is called when entering the type_annotations production.
	EnterType_annotations(c *Type_annotationsContext)

	// EnterType_annotation is called when entering the type_annotation production.
	EnterType_annotation(c *Type_annotationContext)

	// EnterAnnotation_value is called when entering the annotation_value production.
	EnterAnnotation_value(c *Annotation_valueContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// EnterBase_type is called when entering the base_type production.
	EnterBase_type(c *Base_typeContext)

	// EnterContainer_type is called when entering the container_type production.
	EnterContainer_type(c *Container_typeContext)

	// EnterMap_type is called when entering the map_type production.
	EnterMap_type(c *Map_typeContext)

	// EnterSet_type is called when entering the set_type production.
	EnterSet_type(c *Set_typeContext)

	// EnterList_type is called when entering the list_type production.
	EnterList_type(c *List_typeContext)

	// EnterCpp_type is called when entering the cpp_type production.
	EnterCpp_type(c *Cpp_typeContext)

	// EnterConst_value is called when entering the const_value production.
	EnterConst_value(c *Const_valueContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterConst_list is called when entering the const_list production.
	EnterConst_list(c *Const_listContext)

	// EnterConst_map_entry is called when entering the const_map_entry production.
	EnterConst_map_entry(c *Const_map_entryContext)

	// EnterConst_map is called when entering the const_map production.
	EnterConst_map(c *Const_mapContext)

	// EnterList_separator is called when entering the list_separator production.
	EnterList_separator(c *List_separatorContext)

	// EnterReal_base_type is called when entering the real_base_type production.
	EnterReal_base_type(c *Real_base_typeContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitInclude_ is called when exiting the include_ production.
	ExitInclude_(c *Include_Context)

	// ExitNamespace_ is called when exiting the namespace_ production.
	ExitNamespace_(c *Namespace_Context)

	// ExitCpp_include is called when exiting the cpp_include production.
	ExitCpp_include(c *Cpp_includeContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitConst_rule is called when exiting the const_rule production.
	ExitConst_rule(c *Const_ruleContext)

	// ExitTypedef_ is called when exiting the typedef_ production.
	ExitTypedef_(c *Typedef_Context)

	// ExitEnum_rule is called when exiting the enum_rule production.
	ExitEnum_rule(c *Enum_ruleContext)

	// ExitEnum_field is called when exiting the enum_field production.
	ExitEnum_field(c *Enum_fieldContext)

	// ExitSenum is called when exiting the senum production.
	ExitSenum(c *SenumContext)

	// ExitStruct_ is called when exiting the struct_ production.
	ExitStruct_(c *Struct_Context)

	// ExitUnion_ is called when exiting the union_ production.
	ExitUnion_(c *Union_Context)

	// ExitException is called when exiting the exception production.
	ExitException(c *ExceptionContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitField_id is called when exiting the field_id production.
	ExitField_id(c *Field_idContext)

	// ExitField_req is called when exiting the field_req production.
	ExitField_req(c *Field_reqContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitOneway is called when exiting the oneway production.
	ExitOneway(c *OnewayContext)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitThrows_list is called when exiting the throws_list production.
	ExitThrows_list(c *Throws_listContext)

	// ExitType_annotations is called when exiting the type_annotations production.
	ExitType_annotations(c *Type_annotationsContext)

	// ExitType_annotation is called when exiting the type_annotation production.
	ExitType_annotation(c *Type_annotationContext)

	// ExitAnnotation_value is called when exiting the annotation_value production.
	ExitAnnotation_value(c *Annotation_valueContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)

	// ExitBase_type is called when exiting the base_type production.
	ExitBase_type(c *Base_typeContext)

	// ExitContainer_type is called when exiting the container_type production.
	ExitContainer_type(c *Container_typeContext)

	// ExitMap_type is called when exiting the map_type production.
	ExitMap_type(c *Map_typeContext)

	// ExitSet_type is called when exiting the set_type production.
	ExitSet_type(c *Set_typeContext)

	// ExitList_type is called when exiting the list_type production.
	ExitList_type(c *List_typeContext)

	// ExitCpp_type is called when exiting the cpp_type production.
	ExitCpp_type(c *Cpp_typeContext)

	// ExitConst_value is called when exiting the const_value production.
	ExitConst_value(c *Const_valueContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitConst_list is called when exiting the const_list production.
	ExitConst_list(c *Const_listContext)

	// ExitConst_map_entry is called when exiting the const_map_entry production.
	ExitConst_map_entry(c *Const_map_entryContext)

	// ExitConst_map is called when exiting the const_map production.
	ExitConst_map(c *Const_mapContext)

	// ExitList_separator is called when exiting the list_separator production.
	ExitList_separator(c *List_separatorContext)

	// ExitReal_base_type is called when exiting the real_base_type production.
	ExitReal_base_type(c *Real_base_typeContext)
}
