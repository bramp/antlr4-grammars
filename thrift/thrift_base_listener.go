// Code generated from Thrift.g4 by ANTLR 4.9.3. DO NOT EDIT.

package thrift // Thrift
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseThriftListener is a complete listener for a parse tree produced by ThriftParser.
type BaseThriftListener struct{}

var _ ThriftListener = &BaseThriftListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseThriftListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseThriftListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseThriftListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseThriftListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseThriftListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseThriftListener) ExitDocument(ctx *DocumentContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseThriftListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseThriftListener) ExitHeader(ctx *HeaderContext) {}

// EnterInclude_ is called when production include_ is entered.
func (s *BaseThriftListener) EnterInclude_(ctx *Include_Context) {}

// ExitInclude_ is called when production include_ is exited.
func (s *BaseThriftListener) ExitInclude_(ctx *Include_Context) {}

// EnterNamespace_ is called when production namespace_ is entered.
func (s *BaseThriftListener) EnterNamespace_(ctx *Namespace_Context) {}

// ExitNamespace_ is called when production namespace_ is exited.
func (s *BaseThriftListener) ExitNamespace_(ctx *Namespace_Context) {}

// EnterCpp_include is called when production cpp_include is entered.
func (s *BaseThriftListener) EnterCpp_include(ctx *Cpp_includeContext) {}

// ExitCpp_include is called when production cpp_include is exited.
func (s *BaseThriftListener) ExitCpp_include(ctx *Cpp_includeContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseThriftListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseThriftListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterConst_rule is called when production const_rule is entered.
func (s *BaseThriftListener) EnterConst_rule(ctx *Const_ruleContext) {}

// ExitConst_rule is called when production const_rule is exited.
func (s *BaseThriftListener) ExitConst_rule(ctx *Const_ruleContext) {}

// EnterTypedef_ is called when production typedef_ is entered.
func (s *BaseThriftListener) EnterTypedef_(ctx *Typedef_Context) {}

// ExitTypedef_ is called when production typedef_ is exited.
func (s *BaseThriftListener) ExitTypedef_(ctx *Typedef_Context) {}

// EnterEnum_rule is called when production enum_rule is entered.
func (s *BaseThriftListener) EnterEnum_rule(ctx *Enum_ruleContext) {}

// ExitEnum_rule is called when production enum_rule is exited.
func (s *BaseThriftListener) ExitEnum_rule(ctx *Enum_ruleContext) {}

// EnterEnum_field is called when production enum_field is entered.
func (s *BaseThriftListener) EnterEnum_field(ctx *Enum_fieldContext) {}

// ExitEnum_field is called when production enum_field is exited.
func (s *BaseThriftListener) ExitEnum_field(ctx *Enum_fieldContext) {}

// EnterSenum is called when production senum is entered.
func (s *BaseThriftListener) EnterSenum(ctx *SenumContext) {}

// ExitSenum is called when production senum is exited.
func (s *BaseThriftListener) ExitSenum(ctx *SenumContext) {}

// EnterStruct_ is called when production struct_ is entered.
func (s *BaseThriftListener) EnterStruct_(ctx *Struct_Context) {}

// ExitStruct_ is called when production struct_ is exited.
func (s *BaseThriftListener) ExitStruct_(ctx *Struct_Context) {}

// EnterUnion_ is called when production union_ is entered.
func (s *BaseThriftListener) EnterUnion_(ctx *Union_Context) {}

// ExitUnion_ is called when production union_ is exited.
func (s *BaseThriftListener) ExitUnion_(ctx *Union_Context) {}

// EnterException is called when production exception is entered.
func (s *BaseThriftListener) EnterException(ctx *ExceptionContext) {}

// ExitException is called when production exception is exited.
func (s *BaseThriftListener) ExitException(ctx *ExceptionContext) {}

// EnterService is called when production service is entered.
func (s *BaseThriftListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseThriftListener) ExitService(ctx *ServiceContext) {}

// EnterField is called when production field is entered.
func (s *BaseThriftListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseThriftListener) ExitField(ctx *FieldContext) {}

// EnterField_id is called when production field_id is entered.
func (s *BaseThriftListener) EnterField_id(ctx *Field_idContext) {}

// ExitField_id is called when production field_id is exited.
func (s *BaseThriftListener) ExitField_id(ctx *Field_idContext) {}

// EnterField_req is called when production field_req is entered.
func (s *BaseThriftListener) EnterField_req(ctx *Field_reqContext) {}

// ExitField_req is called when production field_req is exited.
func (s *BaseThriftListener) ExitField_req(ctx *Field_reqContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaseThriftListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaseThriftListener) ExitFunction_(ctx *Function_Context) {}

// EnterOneway is called when production oneway is entered.
func (s *BaseThriftListener) EnterOneway(ctx *OnewayContext) {}

// ExitOneway is called when production oneway is exited.
func (s *BaseThriftListener) ExitOneway(ctx *OnewayContext) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BaseThriftListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BaseThriftListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterThrows_list is called when production throws_list is entered.
func (s *BaseThriftListener) EnterThrows_list(ctx *Throws_listContext) {}

// ExitThrows_list is called when production throws_list is exited.
func (s *BaseThriftListener) ExitThrows_list(ctx *Throws_listContext) {}

// EnterType_annotations is called when production type_annotations is entered.
func (s *BaseThriftListener) EnterType_annotations(ctx *Type_annotationsContext) {}

// ExitType_annotations is called when production type_annotations is exited.
func (s *BaseThriftListener) ExitType_annotations(ctx *Type_annotationsContext) {}

// EnterType_annotation is called when production type_annotation is entered.
func (s *BaseThriftListener) EnterType_annotation(ctx *Type_annotationContext) {}

// ExitType_annotation is called when production type_annotation is exited.
func (s *BaseThriftListener) ExitType_annotation(ctx *Type_annotationContext) {}

// EnterAnnotation_value is called when production annotation_value is entered.
func (s *BaseThriftListener) EnterAnnotation_value(ctx *Annotation_valueContext) {}

// ExitAnnotation_value is called when production annotation_value is exited.
func (s *BaseThriftListener) ExitAnnotation_value(ctx *Annotation_valueContext) {}

// EnterField_type is called when production field_type is entered.
func (s *BaseThriftListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BaseThriftListener) ExitField_type(ctx *Field_typeContext) {}

// EnterBase_type is called when production base_type is entered.
func (s *BaseThriftListener) EnterBase_type(ctx *Base_typeContext) {}

// ExitBase_type is called when production base_type is exited.
func (s *BaseThriftListener) ExitBase_type(ctx *Base_typeContext) {}

// EnterContainer_type is called when production container_type is entered.
func (s *BaseThriftListener) EnterContainer_type(ctx *Container_typeContext) {}

// ExitContainer_type is called when production container_type is exited.
func (s *BaseThriftListener) ExitContainer_type(ctx *Container_typeContext) {}

// EnterMap_type is called when production map_type is entered.
func (s *BaseThriftListener) EnterMap_type(ctx *Map_typeContext) {}

// ExitMap_type is called when production map_type is exited.
func (s *BaseThriftListener) ExitMap_type(ctx *Map_typeContext) {}

// EnterSet_type is called when production set_type is entered.
func (s *BaseThriftListener) EnterSet_type(ctx *Set_typeContext) {}

// ExitSet_type is called when production set_type is exited.
func (s *BaseThriftListener) ExitSet_type(ctx *Set_typeContext) {}

// EnterList_type is called when production list_type is entered.
func (s *BaseThriftListener) EnterList_type(ctx *List_typeContext) {}

// ExitList_type is called when production list_type is exited.
func (s *BaseThriftListener) ExitList_type(ctx *List_typeContext) {}

// EnterCpp_type is called when production cpp_type is entered.
func (s *BaseThriftListener) EnterCpp_type(ctx *Cpp_typeContext) {}

// ExitCpp_type is called when production cpp_type is exited.
func (s *BaseThriftListener) ExitCpp_type(ctx *Cpp_typeContext) {}

// EnterConst_value is called when production const_value is entered.
func (s *BaseThriftListener) EnterConst_value(ctx *Const_valueContext) {}

// ExitConst_value is called when production const_value is exited.
func (s *BaseThriftListener) ExitConst_value(ctx *Const_valueContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseThriftListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseThriftListener) ExitInteger(ctx *IntegerContext) {}

// EnterConst_list is called when production const_list is entered.
func (s *BaseThriftListener) EnterConst_list(ctx *Const_listContext) {}

// ExitConst_list is called when production const_list is exited.
func (s *BaseThriftListener) ExitConst_list(ctx *Const_listContext) {}

// EnterConst_map_entry is called when production const_map_entry is entered.
func (s *BaseThriftListener) EnterConst_map_entry(ctx *Const_map_entryContext) {}

// ExitConst_map_entry is called when production const_map_entry is exited.
func (s *BaseThriftListener) ExitConst_map_entry(ctx *Const_map_entryContext) {}

// EnterConst_map is called when production const_map is entered.
func (s *BaseThriftListener) EnterConst_map(ctx *Const_mapContext) {}

// ExitConst_map is called when production const_map is exited.
func (s *BaseThriftListener) ExitConst_map(ctx *Const_mapContext) {}

// EnterList_separator is called when production list_separator is entered.
func (s *BaseThriftListener) EnterList_separator(ctx *List_separatorContext) {}

// ExitList_separator is called when production list_separator is exited.
func (s *BaseThriftListener) ExitList_separator(ctx *List_separatorContext) {}

// EnterReal_base_type is called when production real_base_type is entered.
func (s *BaseThriftListener) EnterReal_base_type(ctx *Real_base_typeContext) {}

// ExitReal_base_type is called when production real_base_type is exited.
func (s *BaseThriftListener) ExitReal_base_type(ctx *Real_base_typeContext) {}
