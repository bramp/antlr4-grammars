// Generated from IDL.g4 by ANTLR 4.7.

package idl // IDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by IDLParser.
type IDLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by IDLParser#specification.
	VisitSpecification(ctx *SpecificationContext) interface{}

	// Visit a parse tree produced by IDLParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by IDLParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_or_forward_decl.
	VisitInterface_or_forward_decl(ctx *Interface_or_forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_decl.
	VisitInterface_decl(ctx *Interface_declContext) interface{}

	// Visit a parse tree produced by IDLParser#forward_decl.
	VisitForward_decl(ctx *Forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_header.
	VisitInterface_header(ctx *Interface_headerContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_body.
	VisitInterface_body(ctx *Interface_bodyContext) interface{}

	// Visit a parse tree produced by IDLParser#export.
	VisitExport(ctx *ExportContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_inheritance_spec.
	VisitInterface_inheritance_spec(ctx *Interface_inheritance_specContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_name.
	VisitInterface_name(ctx *Interface_nameContext) interface{}

	// Visit a parse tree produced by IDLParser#scoped_name.
	VisitScoped_name(ctx *Scoped_nameContext) interface{}

	// Visit a parse tree produced by IDLParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by IDLParser#value_forward_decl.
	VisitValue_forward_decl(ctx *Value_forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#value_box_decl.
	VisitValue_box_decl(ctx *Value_box_declContext) interface{}

	// Visit a parse tree produced by IDLParser#value_abs_decl.
	VisitValue_abs_decl(ctx *Value_abs_declContext) interface{}

	// Visit a parse tree produced by IDLParser#value_decl.
	VisitValue_decl(ctx *Value_declContext) interface{}

	// Visit a parse tree produced by IDLParser#value_header.
	VisitValue_header(ctx *Value_headerContext) interface{}

	// Visit a parse tree produced by IDLParser#value_inheritance_spec.
	VisitValue_inheritance_spec(ctx *Value_inheritance_specContext) interface{}

	// Visit a parse tree produced by IDLParser#value_name.
	VisitValue_name(ctx *Value_nameContext) interface{}

	// Visit a parse tree produced by IDLParser#value_element.
	VisitValue_element(ctx *Value_elementContext) interface{}

	// Visit a parse tree produced by IDLParser#state_member.
	VisitState_member(ctx *State_memberContext) interface{}

	// Visit a parse tree produced by IDLParser#init_decl.
	VisitInit_decl(ctx *Init_declContext) interface{}

	// Visit a parse tree produced by IDLParser#init_param_decls.
	VisitInit_param_decls(ctx *Init_param_declsContext) interface{}

	// Visit a parse tree produced by IDLParser#init_param_decl.
	VisitInit_param_decl(ctx *Init_param_declContext) interface{}

	// Visit a parse tree produced by IDLParser#init_param_attribute.
	VisitInit_param_attribute(ctx *Init_param_attributeContext) interface{}

	// Visit a parse tree produced by IDLParser#const_decl.
	VisitConst_decl(ctx *Const_declContext) interface{}

	// Visit a parse tree produced by IDLParser#const_type.
	VisitConst_type(ctx *Const_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#const_exp.
	VisitConst_exp(ctx *Const_expContext) interface{}

	// Visit a parse tree produced by IDLParser#or_expr.
	VisitOr_expr(ctx *Or_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#xor_expr.
	VisitXor_expr(ctx *Xor_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#and_expr.
	VisitAnd_expr(ctx *And_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#shift_expr.
	VisitShift_expr(ctx *Shift_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#add_expr.
	VisitAdd_expr(ctx *Add_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#mult_expr.
	VisitMult_expr(ctx *Mult_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#unary_expr.
	VisitUnary_expr(ctx *Unary_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by IDLParser#primary_expr.
	VisitPrimary_expr(ctx *Primary_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by IDLParser#positive_int_const.
	VisitPositive_int_const(ctx *Positive_int_constContext) interface{}

	// Visit a parse tree produced by IDLParser#type_decl.
	VisitType_decl(ctx *Type_declContext) interface{}

	// Visit a parse tree produced by IDLParser#type_declarator.
	VisitType_declarator(ctx *Type_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#type_spec.
	VisitType_spec(ctx *Type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#simple_type_spec.
	VisitSimple_type_spec(ctx *Simple_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#base_type_spec.
	VisitBase_type_spec(ctx *Base_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#template_type_spec.
	VisitTemplate_type_spec(ctx *Template_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#constr_type_spec.
	VisitConstr_type_spec(ctx *Constr_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#declarators.
	VisitDeclarators(ctx *DeclaratorsContext) interface{}

	// Visit a parse tree produced by IDLParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#simple_declarator.
	VisitSimple_declarator(ctx *Simple_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#complex_declarator.
	VisitComplex_declarator(ctx *Complex_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#floating_pt_type.
	VisitFloating_pt_type(ctx *Floating_pt_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#integer_type.
	VisitInteger_type(ctx *Integer_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#signed_int.
	VisitSigned_int(ctx *Signed_intContext) interface{}

	// Visit a parse tree produced by IDLParser#signed_short_int.
	VisitSigned_short_int(ctx *Signed_short_intContext) interface{}

	// Visit a parse tree produced by IDLParser#signed_long_int.
	VisitSigned_long_int(ctx *Signed_long_intContext) interface{}

	// Visit a parse tree produced by IDLParser#signed_longlong_int.
	VisitSigned_longlong_int(ctx *Signed_longlong_intContext) interface{}

	// Visit a parse tree produced by IDLParser#unsigned_int.
	VisitUnsigned_int(ctx *Unsigned_intContext) interface{}

	// Visit a parse tree produced by IDLParser#unsigned_short_int.
	VisitUnsigned_short_int(ctx *Unsigned_short_intContext) interface{}

	// Visit a parse tree produced by IDLParser#unsigned_long_int.
	VisitUnsigned_long_int(ctx *Unsigned_long_intContext) interface{}

	// Visit a parse tree produced by IDLParser#unsigned_longlong_int.
	VisitUnsigned_longlong_int(ctx *Unsigned_longlong_intContext) interface{}

	// Visit a parse tree produced by IDLParser#char_type.
	VisitChar_type(ctx *Char_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#wide_char_type.
	VisitWide_char_type(ctx *Wide_char_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#boolean_type.
	VisitBoolean_type(ctx *Boolean_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#octet_type.
	VisitOctet_type(ctx *Octet_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#any_type.
	VisitAny_type(ctx *Any_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#object_type.
	VisitObject_type(ctx *Object_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#struct_type.
	VisitStruct_type(ctx *Struct_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#member_list.
	VisitMember_list(ctx *Member_listContext) interface{}

	// Visit a parse tree produced by IDLParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by IDLParser#union_type.
	VisitUnion_type(ctx *Union_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#switch_type_spec.
	VisitSwitch_type_spec(ctx *Switch_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#switch_body.
	VisitSwitch_body(ctx *Switch_bodyContext) interface{}

	// Visit a parse tree produced by IDLParser#case_stmt.
	VisitCase_stmt(ctx *Case_stmtContext) interface{}

	// Visit a parse tree produced by IDLParser#case_label.
	VisitCase_label(ctx *Case_labelContext) interface{}

	// Visit a parse tree produced by IDLParser#element_spec.
	VisitElement_spec(ctx *Element_specContext) interface{}

	// Visit a parse tree produced by IDLParser#enum_type.
	VisitEnum_type(ctx *Enum_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by IDLParser#sequence_type.
	VisitSequence_type(ctx *Sequence_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#string_type.
	VisitString_type(ctx *String_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#wide_string_type.
	VisitWide_string_type(ctx *Wide_string_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#array_declarator.
	VisitArray_declarator(ctx *Array_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#fixed_array_size.
	VisitFixed_array_size(ctx *Fixed_array_sizeContext) interface{}

	// Visit a parse tree produced by IDLParser#attr_decl.
	VisitAttr_decl(ctx *Attr_declContext) interface{}

	// Visit a parse tree produced by IDLParser#except_decl.
	VisitExcept_decl(ctx *Except_declContext) interface{}

	// Visit a parse tree produced by IDLParser#op_decl.
	VisitOp_decl(ctx *Op_declContext) interface{}

	// Visit a parse tree produced by IDLParser#op_attribute.
	VisitOp_attribute(ctx *Op_attributeContext) interface{}

	// Visit a parse tree produced by IDLParser#op_type_spec.
	VisitOp_type_spec(ctx *Op_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#parameter_decls.
	VisitParameter_decls(ctx *Parameter_declsContext) interface{}

	// Visit a parse tree produced by IDLParser#param_decl.
	VisitParam_decl(ctx *Param_declContext) interface{}

	// Visit a parse tree produced by IDLParser#param_attribute.
	VisitParam_attribute(ctx *Param_attributeContext) interface{}

	// Visit a parse tree produced by IDLParser#raises_expr.
	VisitRaises_expr(ctx *Raises_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#context_expr.
	VisitContext_expr(ctx *Context_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#param_type_spec.
	VisitParam_type_spec(ctx *Param_type_specContext) interface{}

	// Visit a parse tree produced by IDLParser#fixed_pt_type.
	VisitFixed_pt_type(ctx *Fixed_pt_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#fixed_pt_const_type.
	VisitFixed_pt_const_type(ctx *Fixed_pt_const_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#value_base_type.
	VisitValue_base_type(ctx *Value_base_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#constr_forward_decl.
	VisitConstr_forward_decl(ctx *Constr_forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#import_decl.
	VisitImport_decl(ctx *Import_declContext) interface{}

	// Visit a parse tree produced by IDLParser#imported_scope.
	VisitImported_scope(ctx *Imported_scopeContext) interface{}

	// Visit a parse tree produced by IDLParser#type_id_decl.
	VisitType_id_decl(ctx *Type_id_declContext) interface{}

	// Visit a parse tree produced by IDLParser#type_prefix_decl.
	VisitType_prefix_decl(ctx *Type_prefix_declContext) interface{}

	// Visit a parse tree produced by IDLParser#readonly_attr_spec.
	VisitReadonly_attr_spec(ctx *Readonly_attr_specContext) interface{}

	// Visit a parse tree produced by IDLParser#readonly_attr_declarator.
	VisitReadonly_attr_declarator(ctx *Readonly_attr_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#attr_spec.
	VisitAttr_spec(ctx *Attr_specContext) interface{}

	// Visit a parse tree produced by IDLParser#attr_declarator.
	VisitAttr_declarator(ctx *Attr_declaratorContext) interface{}

	// Visit a parse tree produced by IDLParser#attr_raises_expr.
	VisitAttr_raises_expr(ctx *Attr_raises_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#get_excep_expr.
	VisitGet_excep_expr(ctx *Get_excep_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#set_excep_expr.
	VisitSet_excep_expr(ctx *Set_excep_exprContext) interface{}

	// Visit a parse tree produced by IDLParser#exception_list.
	VisitException_list(ctx *Exception_listContext) interface{}

	// Visit a parse tree produced by IDLParser#component.
	VisitComponent(ctx *ComponentContext) interface{}

	// Visit a parse tree produced by IDLParser#component_forward_decl.
	VisitComponent_forward_decl(ctx *Component_forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#component_decl.
	VisitComponent_decl(ctx *Component_declContext) interface{}

	// Visit a parse tree produced by IDLParser#component_header.
	VisitComponent_header(ctx *Component_headerContext) interface{}

	// Visit a parse tree produced by IDLParser#supported_interface_spec.
	VisitSupported_interface_spec(ctx *Supported_interface_specContext) interface{}

	// Visit a parse tree produced by IDLParser#component_inheritance_spec.
	VisitComponent_inheritance_spec(ctx *Component_inheritance_specContext) interface{}

	// Visit a parse tree produced by IDLParser#component_body.
	VisitComponent_body(ctx *Component_bodyContext) interface{}

	// Visit a parse tree produced by IDLParser#component_export.
	VisitComponent_export(ctx *Component_exportContext) interface{}

	// Visit a parse tree produced by IDLParser#provides_decl.
	VisitProvides_decl(ctx *Provides_declContext) interface{}

	// Visit a parse tree produced by IDLParser#interface_type.
	VisitInterface_type(ctx *Interface_typeContext) interface{}

	// Visit a parse tree produced by IDLParser#uses_decl.
	VisitUses_decl(ctx *Uses_declContext) interface{}

	// Visit a parse tree produced by IDLParser#emits_decl.
	VisitEmits_decl(ctx *Emits_declContext) interface{}

	// Visit a parse tree produced by IDLParser#publishes_decl.
	VisitPublishes_decl(ctx *Publishes_declContext) interface{}

	// Visit a parse tree produced by IDLParser#consumes_decl.
	VisitConsumes_decl(ctx *Consumes_declContext) interface{}

	// Visit a parse tree produced by IDLParser#home_decl.
	VisitHome_decl(ctx *Home_declContext) interface{}

	// Visit a parse tree produced by IDLParser#home_header.
	VisitHome_header(ctx *Home_headerContext) interface{}

	// Visit a parse tree produced by IDLParser#home_inheritance_spec.
	VisitHome_inheritance_spec(ctx *Home_inheritance_specContext) interface{}

	// Visit a parse tree produced by IDLParser#primary_key_spec.
	VisitPrimary_key_spec(ctx *Primary_key_specContext) interface{}

	// Visit a parse tree produced by IDLParser#home_body.
	VisitHome_body(ctx *Home_bodyContext) interface{}

	// Visit a parse tree produced by IDLParser#home_export.
	VisitHome_export(ctx *Home_exportContext) interface{}

	// Visit a parse tree produced by IDLParser#factory_decl.
	VisitFactory_decl(ctx *Factory_declContext) interface{}

	// Visit a parse tree produced by IDLParser#finder_decl.
	VisitFinder_decl(ctx *Finder_declContext) interface{}

	// Visit a parse tree produced by IDLParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by IDLParser#event_forward_decl.
	VisitEvent_forward_decl(ctx *Event_forward_declContext) interface{}

	// Visit a parse tree produced by IDLParser#event_abs_decl.
	VisitEvent_abs_decl(ctx *Event_abs_declContext) interface{}

	// Visit a parse tree produced by IDLParser#event_decl.
	VisitEvent_decl(ctx *Event_declContext) interface{}

	// Visit a parse tree produced by IDLParser#event_header.
	VisitEvent_header(ctx *Event_headerContext) interface{}
}
