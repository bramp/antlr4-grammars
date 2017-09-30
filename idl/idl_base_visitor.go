// Generated from IDL.g4 by ANTLR 4.7.

package idl // IDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseIDLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseIDLVisitor) VisitSpecification(ctx *SpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_or_forward_decl(ctx *Interface_or_forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_decl(ctx *Interface_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitForward_decl(ctx *Forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_header(ctx *Interface_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_body(ctx *Interface_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitExport(ctx *ExportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_inheritance_spec(ctx *Interface_inheritance_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_name(ctx *Interface_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitScoped_name(ctx *Scoped_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_forward_decl(ctx *Value_forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_box_decl(ctx *Value_box_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_abs_decl(ctx *Value_abs_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_decl(ctx *Value_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_header(ctx *Value_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_inheritance_spec(ctx *Value_inheritance_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_name(ctx *Value_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_element(ctx *Value_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitState_member(ctx *State_memberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInit_decl(ctx *Init_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInit_param_decls(ctx *Init_param_declsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInit_param_decl(ctx *Init_param_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInit_param_attribute(ctx *Init_param_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConst_decl(ctx *Const_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConst_type(ctx *Const_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConst_exp(ctx *Const_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitOr_expr(ctx *Or_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitXor_expr(ctx *Xor_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAnd_expr(ctx *And_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitShift_expr(ctx *Shift_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAdd_expr(ctx *Add_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitMult_expr(ctx *Mult_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnary_expr(ctx *Unary_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitPrimary_expr(ctx *Primary_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitPositive_int_const(ctx *Positive_int_constContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitType_decl(ctx *Type_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitType_declarator(ctx *Type_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitType_spec(ctx *Type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSimple_type_spec(ctx *Simple_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitBase_type_spec(ctx *Base_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitTemplate_type_spec(ctx *Template_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConstr_type_spec(ctx *Constr_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitDeclarators(ctx *DeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitDeclarator(ctx *DeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSimple_declarator(ctx *Simple_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComplex_declarator(ctx *Complex_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFloating_pt_type(ctx *Floating_pt_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInteger_type(ctx *Integer_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSigned_int(ctx *Signed_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSigned_short_int(ctx *Signed_short_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSigned_long_int(ctx *Signed_long_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSigned_longlong_int(ctx *Signed_longlong_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnsigned_int(ctx *Unsigned_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnsigned_short_int(ctx *Unsigned_short_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnsigned_long_int(ctx *Unsigned_long_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnsigned_longlong_int(ctx *Unsigned_longlong_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitChar_type(ctx *Char_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitWide_char_type(ctx *Wide_char_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitBoolean_type(ctx *Boolean_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitOctet_type(ctx *Octet_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAny_type(ctx *Any_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitObject_type(ctx *Object_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitStruct_type(ctx *Struct_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitMember_list(ctx *Member_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUnion_type(ctx *Union_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSwitch_type_spec(ctx *Switch_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSwitch_body(ctx *Switch_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitCase_stmt(ctx *Case_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitCase_label(ctx *Case_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitElement_spec(ctx *Element_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEnum_type(ctx *Enum_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEnumerator(ctx *EnumeratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSequence_type(ctx *Sequence_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitString_type(ctx *String_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitWide_string_type(ctx *Wide_string_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitArray_declarator(ctx *Array_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFixed_array_size(ctx *Fixed_array_sizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAttr_decl(ctx *Attr_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitExcept_decl(ctx *Except_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitOp_decl(ctx *Op_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitOp_attribute(ctx *Op_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitOp_type_spec(ctx *Op_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitParameter_decls(ctx *Parameter_declsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitParam_decl(ctx *Param_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitParam_attribute(ctx *Param_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitRaises_expr(ctx *Raises_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitContext_expr(ctx *Context_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitParam_type_spec(ctx *Param_type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFixed_pt_type(ctx *Fixed_pt_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFixed_pt_const_type(ctx *Fixed_pt_const_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitValue_base_type(ctx *Value_base_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConstr_forward_decl(ctx *Constr_forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitImport_decl(ctx *Import_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitImported_scope(ctx *Imported_scopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitType_id_decl(ctx *Type_id_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitType_prefix_decl(ctx *Type_prefix_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitReadonly_attr_spec(ctx *Readonly_attr_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitReadonly_attr_declarator(ctx *Readonly_attr_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAttr_spec(ctx *Attr_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAttr_declarator(ctx *Attr_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitAttr_raises_expr(ctx *Attr_raises_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitGet_excep_expr(ctx *Get_excep_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSet_excep_expr(ctx *Set_excep_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitException_list(ctx *Exception_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent(ctx *ComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_forward_decl(ctx *Component_forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_decl(ctx *Component_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_header(ctx *Component_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitSupported_interface_spec(ctx *Supported_interface_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_inheritance_spec(ctx *Component_inheritance_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_body(ctx *Component_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitComponent_export(ctx *Component_exportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitProvides_decl(ctx *Provides_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitInterface_type(ctx *Interface_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitUses_decl(ctx *Uses_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEmits_decl(ctx *Emits_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitPublishes_decl(ctx *Publishes_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitConsumes_decl(ctx *Consumes_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitHome_decl(ctx *Home_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitHome_header(ctx *Home_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitHome_inheritance_spec(ctx *Home_inheritance_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitPrimary_key_spec(ctx *Primary_key_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitHome_body(ctx *Home_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitHome_export(ctx *Home_exportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFactory_decl(ctx *Factory_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitFinder_decl(ctx *Finder_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEvent_forward_decl(ctx *Event_forward_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEvent_abs_decl(ctx *Event_abs_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEvent_decl(ctx *Event_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseIDLVisitor) VisitEvent_header(ctx *Event_headerContext) interface{} {
	return v.VisitChildren(ctx)
}
