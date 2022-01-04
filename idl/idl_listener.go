// Code generated from IDL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package idl // IDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// IDLListener is a complete listener for a parse tree produced by IDLParser.
type IDLListener interface {
	antlr.ParseTreeListener

	// EnterSpecification is called when entering the specification production.
	EnterSpecification(c *SpecificationContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterInterface_or_forward_decl is called when entering the interface_or_forward_decl production.
	EnterInterface_or_forward_decl(c *Interface_or_forward_declContext)

	// EnterInterface_decl is called when entering the interface_decl production.
	EnterInterface_decl(c *Interface_declContext)

	// EnterForward_decl is called when entering the forward_decl production.
	EnterForward_decl(c *Forward_declContext)

	// EnterInterface_header is called when entering the interface_header production.
	EnterInterface_header(c *Interface_headerContext)

	// EnterInterface_body is called when entering the interface_body production.
	EnterInterface_body(c *Interface_bodyContext)

	// EnterExport_ is called when entering the export_ production.
	EnterExport_(c *Export_Context)

	// EnterInterface_inheritance_spec is called when entering the interface_inheritance_spec production.
	EnterInterface_inheritance_spec(c *Interface_inheritance_specContext)

	// EnterInterface_name is called when entering the interface_name production.
	EnterInterface_name(c *Interface_nameContext)

	// EnterA_scoped_name is called when entering the a_scoped_name production.
	EnterA_scoped_name(c *A_scoped_nameContext)

	// EnterScoped_name is called when entering the scoped_name production.
	EnterScoped_name(c *Scoped_nameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValue_forward_decl is called when entering the value_forward_decl production.
	EnterValue_forward_decl(c *Value_forward_declContext)

	// EnterValue_box_decl is called when entering the value_box_decl production.
	EnterValue_box_decl(c *Value_box_declContext)

	// EnterValue_abs_decl is called when entering the value_abs_decl production.
	EnterValue_abs_decl(c *Value_abs_declContext)

	// EnterValue_decl is called when entering the value_decl production.
	EnterValue_decl(c *Value_declContext)

	// EnterValue_header is called when entering the value_header production.
	EnterValue_header(c *Value_headerContext)

	// EnterValue_inheritance_spec is called when entering the value_inheritance_spec production.
	EnterValue_inheritance_spec(c *Value_inheritance_specContext)

	// EnterValue_name is called when entering the value_name production.
	EnterValue_name(c *Value_nameContext)

	// EnterValue_element is called when entering the value_element production.
	EnterValue_element(c *Value_elementContext)

	// EnterState_member is called when entering the state_member production.
	EnterState_member(c *State_memberContext)

	// EnterInit_decl is called when entering the init_decl production.
	EnterInit_decl(c *Init_declContext)

	// EnterInit_param_decls is called when entering the init_param_decls production.
	EnterInit_param_decls(c *Init_param_declsContext)

	// EnterInit_param_decl is called when entering the init_param_decl production.
	EnterInit_param_decl(c *Init_param_declContext)

	// EnterInit_param_attribute is called when entering the init_param_attribute production.
	EnterInit_param_attribute(c *Init_param_attributeContext)

	// EnterConst_decl is called when entering the const_decl production.
	EnterConst_decl(c *Const_declContext)

	// EnterConst_type is called when entering the const_type production.
	EnterConst_type(c *Const_typeContext)

	// EnterConst_exp is called when entering the const_exp production.
	EnterConst_exp(c *Const_expContext)

	// EnterOr_expr is called when entering the or_expr production.
	EnterOr_expr(c *Or_exprContext)

	// EnterXor_expr is called when entering the xor_expr production.
	EnterXor_expr(c *Xor_exprContext)

	// EnterAnd_expr is called when entering the and_expr production.
	EnterAnd_expr(c *And_exprContext)

	// EnterShift_expr is called when entering the shift_expr production.
	EnterShift_expr(c *Shift_exprContext)

	// EnterAdd_expr is called when entering the add_expr production.
	EnterAdd_expr(c *Add_exprContext)

	// EnterMult_expr is called when entering the mult_expr production.
	EnterMult_expr(c *Mult_exprContext)

	// EnterUnary_expr is called when entering the unary_expr production.
	EnterUnary_expr(c *Unary_exprContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterPrimary_expr is called when entering the primary_expr production.
	EnterPrimary_expr(c *Primary_exprContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPositive_int_const is called when entering the positive_int_const production.
	EnterPositive_int_const(c *Positive_int_constContext)

	// EnterType_decl is called when entering the type_decl production.
	EnterType_decl(c *Type_declContext)

	// EnterType_declarator is called when entering the type_declarator production.
	EnterType_declarator(c *Type_declaratorContext)

	// EnterType_spec is called when entering the type_spec production.
	EnterType_spec(c *Type_specContext)

	// EnterSimple_type_spec is called when entering the simple_type_spec production.
	EnterSimple_type_spec(c *Simple_type_specContext)

	// EnterBitfield_type_spec is called when entering the bitfield_type_spec production.
	EnterBitfield_type_spec(c *Bitfield_type_specContext)

	// EnterBase_type_spec is called when entering the base_type_spec production.
	EnterBase_type_spec(c *Base_type_specContext)

	// EnterTemplate_type_spec is called when entering the template_type_spec production.
	EnterTemplate_type_spec(c *Template_type_specContext)

	// EnterConstr_type_spec is called when entering the constr_type_spec production.
	EnterConstr_type_spec(c *Constr_type_specContext)

	// EnterSimple_declarators is called when entering the simple_declarators production.
	EnterSimple_declarators(c *Simple_declaratorsContext)

	// EnterDeclarators is called when entering the declarators production.
	EnterDeclarators(c *DeclaratorsContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterSimple_declarator is called when entering the simple_declarator production.
	EnterSimple_declarator(c *Simple_declaratorContext)

	// EnterComplex_declarator is called when entering the complex_declarator production.
	EnterComplex_declarator(c *Complex_declaratorContext)

	// EnterFloating_pt_type is called when entering the floating_pt_type production.
	EnterFloating_pt_type(c *Floating_pt_typeContext)

	// EnterInteger_type is called when entering the integer_type production.
	EnterInteger_type(c *Integer_typeContext)

	// EnterSigned_int is called when entering the signed_int production.
	EnterSigned_int(c *Signed_intContext)

	// EnterSigned_tiny_int is called when entering the signed_tiny_int production.
	EnterSigned_tiny_int(c *Signed_tiny_intContext)

	// EnterSigned_short_int is called when entering the signed_short_int production.
	EnterSigned_short_int(c *Signed_short_intContext)

	// EnterSigned_long_int is called when entering the signed_long_int production.
	EnterSigned_long_int(c *Signed_long_intContext)

	// EnterSigned_longlong_int is called when entering the signed_longlong_int production.
	EnterSigned_longlong_int(c *Signed_longlong_intContext)

	// EnterUnsigned_int is called when entering the unsigned_int production.
	EnterUnsigned_int(c *Unsigned_intContext)

	// EnterUnsigned_tiny_int is called when entering the unsigned_tiny_int production.
	EnterUnsigned_tiny_int(c *Unsigned_tiny_intContext)

	// EnterUnsigned_short_int is called when entering the unsigned_short_int production.
	EnterUnsigned_short_int(c *Unsigned_short_intContext)

	// EnterUnsigned_long_int is called when entering the unsigned_long_int production.
	EnterUnsigned_long_int(c *Unsigned_long_intContext)

	// EnterUnsigned_longlong_int is called when entering the unsigned_longlong_int production.
	EnterUnsigned_longlong_int(c *Unsigned_longlong_intContext)

	// EnterChar_type is called when entering the char_type production.
	EnterChar_type(c *Char_typeContext)

	// EnterWide_char_type is called when entering the wide_char_type production.
	EnterWide_char_type(c *Wide_char_typeContext)

	// EnterBoolean_type is called when entering the boolean_type production.
	EnterBoolean_type(c *Boolean_typeContext)

	// EnterOctet_type is called when entering the octet_type production.
	EnterOctet_type(c *Octet_typeContext)

	// EnterAny_type is called when entering the any_type production.
	EnterAny_type(c *Any_typeContext)

	// EnterObject_type is called when entering the object_type production.
	EnterObject_type(c *Object_typeContext)

	// EnterAnnotation_decl is called when entering the annotation_decl production.
	EnterAnnotation_decl(c *Annotation_declContext)

	// EnterAnnotation_def is called when entering the annotation_def production.
	EnterAnnotation_def(c *Annotation_defContext)

	// EnterAnnotation_header is called when entering the annotation_header production.
	EnterAnnotation_header(c *Annotation_headerContext)

	// EnterAnnotation_inheritance_spec is called when entering the annotation_inheritance_spec production.
	EnterAnnotation_inheritance_spec(c *Annotation_inheritance_specContext)

	// EnterAnnotation_body is called when entering the annotation_body production.
	EnterAnnotation_body(c *Annotation_bodyContext)

	// EnterAnnotation_member is called when entering the annotation_member production.
	EnterAnnotation_member(c *Annotation_memberContext)

	// EnterAnnotation_forward_dcl is called when entering the annotation_forward_dcl production.
	EnterAnnotation_forward_dcl(c *Annotation_forward_dclContext)

	// EnterBitset_type is called when entering the bitset_type production.
	EnterBitset_type(c *Bitset_typeContext)

	// EnterBitfield is called when entering the bitfield production.
	EnterBitfield(c *BitfieldContext)

	// EnterBitfield_spec is called when entering the bitfield_spec production.
	EnterBitfield_spec(c *Bitfield_specContext)

	// EnterBitmask_type is called when entering the bitmask_type production.
	EnterBitmask_type(c *Bitmask_typeContext)

	// EnterBit_values is called when entering the bit_values production.
	EnterBit_values(c *Bit_valuesContext)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterMember_list is called when entering the member_list production.
	EnterMember_list(c *Member_listContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterUnion_type is called when entering the union_type production.
	EnterUnion_type(c *Union_typeContext)

	// EnterSwitch_type_spec is called when entering the switch_type_spec production.
	EnterSwitch_type_spec(c *Switch_type_specContext)

	// EnterSwitch_body is called when entering the switch_body production.
	EnterSwitch_body(c *Switch_bodyContext)

	// EnterCase_stmt is called when entering the case_stmt production.
	EnterCase_stmt(c *Case_stmtContext)

	// EnterCase_label is called when entering the case_label production.
	EnterCase_label(c *Case_labelContext)

	// EnterElement_spec is called when entering the element_spec production.
	EnterElement_spec(c *Element_specContext)

	// EnterEnum_type is called when entering the enum_type production.
	EnterEnum_type(c *Enum_typeContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterSequence_type is called when entering the sequence_type production.
	EnterSequence_type(c *Sequence_typeContext)

	// EnterSet_type is called when entering the set_type production.
	EnterSet_type(c *Set_typeContext)

	// EnterMap_type is called when entering the map_type production.
	EnterMap_type(c *Map_typeContext)

	// EnterString_type is called when entering the string_type production.
	EnterString_type(c *String_typeContext)

	// EnterWide_string_type is called when entering the wide_string_type production.
	EnterWide_string_type(c *Wide_string_typeContext)

	// EnterArray_declarator is called when entering the array_declarator production.
	EnterArray_declarator(c *Array_declaratorContext)

	// EnterFixed_array_size is called when entering the fixed_array_size production.
	EnterFixed_array_size(c *Fixed_array_sizeContext)

	// EnterAttr_decl is called when entering the attr_decl production.
	EnterAttr_decl(c *Attr_declContext)

	// EnterExcept_decl is called when entering the except_decl production.
	EnterExcept_decl(c *Except_declContext)

	// EnterOp_decl is called when entering the op_decl production.
	EnterOp_decl(c *Op_declContext)

	// EnterOp_attribute is called when entering the op_attribute production.
	EnterOp_attribute(c *Op_attributeContext)

	// EnterOp_type_spec is called when entering the op_type_spec production.
	EnterOp_type_spec(c *Op_type_specContext)

	// EnterParameter_decls is called when entering the parameter_decls production.
	EnterParameter_decls(c *Parameter_declsContext)

	// EnterParam_decl is called when entering the param_decl production.
	EnterParam_decl(c *Param_declContext)

	// EnterParam_attribute is called when entering the param_attribute production.
	EnterParam_attribute(c *Param_attributeContext)

	// EnterRaises_expr is called when entering the raises_expr production.
	EnterRaises_expr(c *Raises_exprContext)

	// EnterContext_expr is called when entering the context_expr production.
	EnterContext_expr(c *Context_exprContext)

	// EnterParam_type_spec is called when entering the param_type_spec production.
	EnterParam_type_spec(c *Param_type_specContext)

	// EnterFixed_pt_type is called when entering the fixed_pt_type production.
	EnterFixed_pt_type(c *Fixed_pt_typeContext)

	// EnterFixed_pt_const_type is called when entering the fixed_pt_const_type production.
	EnterFixed_pt_const_type(c *Fixed_pt_const_typeContext)

	// EnterValue_base_type is called when entering the value_base_type production.
	EnterValue_base_type(c *Value_base_typeContext)

	// EnterConstr_forward_decl is called when entering the constr_forward_decl production.
	EnterConstr_forward_decl(c *Constr_forward_declContext)

	// EnterImport_decl is called when entering the import_decl production.
	EnterImport_decl(c *Import_declContext)

	// EnterImported_scope is called when entering the imported_scope production.
	EnterImported_scope(c *Imported_scopeContext)

	// EnterType_id_decl is called when entering the type_id_decl production.
	EnterType_id_decl(c *Type_id_declContext)

	// EnterType_prefix_decl is called when entering the type_prefix_decl production.
	EnterType_prefix_decl(c *Type_prefix_declContext)

	// EnterReadonly_attr_spec is called when entering the readonly_attr_spec production.
	EnterReadonly_attr_spec(c *Readonly_attr_specContext)

	// EnterReadonly_attr_declarator is called when entering the readonly_attr_declarator production.
	EnterReadonly_attr_declarator(c *Readonly_attr_declaratorContext)

	// EnterAttr_spec is called when entering the attr_spec production.
	EnterAttr_spec(c *Attr_specContext)

	// EnterAttr_declarator is called when entering the attr_declarator production.
	EnterAttr_declarator(c *Attr_declaratorContext)

	// EnterAttr_raises_expr is called when entering the attr_raises_expr production.
	EnterAttr_raises_expr(c *Attr_raises_exprContext)

	// EnterGet_excep_expr is called when entering the get_excep_expr production.
	EnterGet_excep_expr(c *Get_excep_exprContext)

	// EnterSet_excep_expr is called when entering the set_excep_expr production.
	EnterSet_excep_expr(c *Set_excep_exprContext)

	// EnterException_list is called when entering the exception_list production.
	EnterException_list(c *Exception_listContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterComponent_forward_decl is called when entering the component_forward_decl production.
	EnterComponent_forward_decl(c *Component_forward_declContext)

	// EnterComponent_decl is called when entering the component_decl production.
	EnterComponent_decl(c *Component_declContext)

	// EnterComponent_header is called when entering the component_header production.
	EnterComponent_header(c *Component_headerContext)

	// EnterSupported_interface_spec is called when entering the supported_interface_spec production.
	EnterSupported_interface_spec(c *Supported_interface_specContext)

	// EnterComponent_inheritance_spec is called when entering the component_inheritance_spec production.
	EnterComponent_inheritance_spec(c *Component_inheritance_specContext)

	// EnterComponent_body is called when entering the component_body production.
	EnterComponent_body(c *Component_bodyContext)

	// EnterComponent_export is called when entering the component_export production.
	EnterComponent_export(c *Component_exportContext)

	// EnterProvides_decl is called when entering the provides_decl production.
	EnterProvides_decl(c *Provides_declContext)

	// EnterInterface_type is called when entering the interface_type production.
	EnterInterface_type(c *Interface_typeContext)

	// EnterUses_decl is called when entering the uses_decl production.
	EnterUses_decl(c *Uses_declContext)

	// EnterEmits_decl is called when entering the emits_decl production.
	EnterEmits_decl(c *Emits_declContext)

	// EnterPublishes_decl is called when entering the publishes_decl production.
	EnterPublishes_decl(c *Publishes_declContext)

	// EnterConsumes_decl is called when entering the consumes_decl production.
	EnterConsumes_decl(c *Consumes_declContext)

	// EnterHome_decl is called when entering the home_decl production.
	EnterHome_decl(c *Home_declContext)

	// EnterHome_header is called when entering the home_header production.
	EnterHome_header(c *Home_headerContext)

	// EnterHome_inheritance_spec is called when entering the home_inheritance_spec production.
	EnterHome_inheritance_spec(c *Home_inheritance_specContext)

	// EnterPrimary_key_spec is called when entering the primary_key_spec production.
	EnterPrimary_key_spec(c *Primary_key_specContext)

	// EnterHome_body is called when entering the home_body production.
	EnterHome_body(c *Home_bodyContext)

	// EnterHome_export is called when entering the home_export production.
	EnterHome_export(c *Home_exportContext)

	// EnterFactory_decl is called when entering the factory_decl production.
	EnterFactory_decl(c *Factory_declContext)

	// EnterFinder_decl is called when entering the finder_decl production.
	EnterFinder_decl(c *Finder_declContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterEvent_forward_decl is called when entering the event_forward_decl production.
	EnterEvent_forward_decl(c *Event_forward_declContext)

	// EnterEvent_abs_decl is called when entering the event_abs_decl production.
	EnterEvent_abs_decl(c *Event_abs_declContext)

	// EnterEvent_decl is called when entering the event_decl production.
	EnterEvent_decl(c *Event_declContext)

	// EnterEvent_header is called when entering the event_header production.
	EnterEvent_header(c *Event_headerContext)

	// EnterAnnapps is called when entering the annapps production.
	EnterAnnapps(c *AnnappsContext)

	// EnterAnnotation_appl is called when entering the annotation_appl production.
	EnterAnnotation_appl(c *Annotation_applContext)

	// EnterAnnotation_appl_params is called when entering the annotation_appl_params production.
	EnterAnnotation_appl_params(c *Annotation_appl_paramsContext)

	// EnterAnnotation_appl_param is called when entering the annotation_appl_param production.
	EnterAnnotation_appl_param(c *Annotation_appl_paramContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitSpecification is called when exiting the specification production.
	ExitSpecification(c *SpecificationContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitInterface_or_forward_decl is called when exiting the interface_or_forward_decl production.
	ExitInterface_or_forward_decl(c *Interface_or_forward_declContext)

	// ExitInterface_decl is called when exiting the interface_decl production.
	ExitInterface_decl(c *Interface_declContext)

	// ExitForward_decl is called when exiting the forward_decl production.
	ExitForward_decl(c *Forward_declContext)

	// ExitInterface_header is called when exiting the interface_header production.
	ExitInterface_header(c *Interface_headerContext)

	// ExitInterface_body is called when exiting the interface_body production.
	ExitInterface_body(c *Interface_bodyContext)

	// ExitExport_ is called when exiting the export_ production.
	ExitExport_(c *Export_Context)

	// ExitInterface_inheritance_spec is called when exiting the interface_inheritance_spec production.
	ExitInterface_inheritance_spec(c *Interface_inheritance_specContext)

	// ExitInterface_name is called when exiting the interface_name production.
	ExitInterface_name(c *Interface_nameContext)

	// ExitA_scoped_name is called when exiting the a_scoped_name production.
	ExitA_scoped_name(c *A_scoped_nameContext)

	// ExitScoped_name is called when exiting the scoped_name production.
	ExitScoped_name(c *Scoped_nameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValue_forward_decl is called when exiting the value_forward_decl production.
	ExitValue_forward_decl(c *Value_forward_declContext)

	// ExitValue_box_decl is called when exiting the value_box_decl production.
	ExitValue_box_decl(c *Value_box_declContext)

	// ExitValue_abs_decl is called when exiting the value_abs_decl production.
	ExitValue_abs_decl(c *Value_abs_declContext)

	// ExitValue_decl is called when exiting the value_decl production.
	ExitValue_decl(c *Value_declContext)

	// ExitValue_header is called when exiting the value_header production.
	ExitValue_header(c *Value_headerContext)

	// ExitValue_inheritance_spec is called when exiting the value_inheritance_spec production.
	ExitValue_inheritance_spec(c *Value_inheritance_specContext)

	// ExitValue_name is called when exiting the value_name production.
	ExitValue_name(c *Value_nameContext)

	// ExitValue_element is called when exiting the value_element production.
	ExitValue_element(c *Value_elementContext)

	// ExitState_member is called when exiting the state_member production.
	ExitState_member(c *State_memberContext)

	// ExitInit_decl is called when exiting the init_decl production.
	ExitInit_decl(c *Init_declContext)

	// ExitInit_param_decls is called when exiting the init_param_decls production.
	ExitInit_param_decls(c *Init_param_declsContext)

	// ExitInit_param_decl is called when exiting the init_param_decl production.
	ExitInit_param_decl(c *Init_param_declContext)

	// ExitInit_param_attribute is called when exiting the init_param_attribute production.
	ExitInit_param_attribute(c *Init_param_attributeContext)

	// ExitConst_decl is called when exiting the const_decl production.
	ExitConst_decl(c *Const_declContext)

	// ExitConst_type is called when exiting the const_type production.
	ExitConst_type(c *Const_typeContext)

	// ExitConst_exp is called when exiting the const_exp production.
	ExitConst_exp(c *Const_expContext)

	// ExitOr_expr is called when exiting the or_expr production.
	ExitOr_expr(c *Or_exprContext)

	// ExitXor_expr is called when exiting the xor_expr production.
	ExitXor_expr(c *Xor_exprContext)

	// ExitAnd_expr is called when exiting the and_expr production.
	ExitAnd_expr(c *And_exprContext)

	// ExitShift_expr is called when exiting the shift_expr production.
	ExitShift_expr(c *Shift_exprContext)

	// ExitAdd_expr is called when exiting the add_expr production.
	ExitAdd_expr(c *Add_exprContext)

	// ExitMult_expr is called when exiting the mult_expr production.
	ExitMult_expr(c *Mult_exprContext)

	// ExitUnary_expr is called when exiting the unary_expr production.
	ExitUnary_expr(c *Unary_exprContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitPrimary_expr is called when exiting the primary_expr production.
	ExitPrimary_expr(c *Primary_exprContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPositive_int_const is called when exiting the positive_int_const production.
	ExitPositive_int_const(c *Positive_int_constContext)

	// ExitType_decl is called when exiting the type_decl production.
	ExitType_decl(c *Type_declContext)

	// ExitType_declarator is called when exiting the type_declarator production.
	ExitType_declarator(c *Type_declaratorContext)

	// ExitType_spec is called when exiting the type_spec production.
	ExitType_spec(c *Type_specContext)

	// ExitSimple_type_spec is called when exiting the simple_type_spec production.
	ExitSimple_type_spec(c *Simple_type_specContext)

	// ExitBitfield_type_spec is called when exiting the bitfield_type_spec production.
	ExitBitfield_type_spec(c *Bitfield_type_specContext)

	// ExitBase_type_spec is called when exiting the base_type_spec production.
	ExitBase_type_spec(c *Base_type_specContext)

	// ExitTemplate_type_spec is called when exiting the template_type_spec production.
	ExitTemplate_type_spec(c *Template_type_specContext)

	// ExitConstr_type_spec is called when exiting the constr_type_spec production.
	ExitConstr_type_spec(c *Constr_type_specContext)

	// ExitSimple_declarators is called when exiting the simple_declarators production.
	ExitSimple_declarators(c *Simple_declaratorsContext)

	// ExitDeclarators is called when exiting the declarators production.
	ExitDeclarators(c *DeclaratorsContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitSimple_declarator is called when exiting the simple_declarator production.
	ExitSimple_declarator(c *Simple_declaratorContext)

	// ExitComplex_declarator is called when exiting the complex_declarator production.
	ExitComplex_declarator(c *Complex_declaratorContext)

	// ExitFloating_pt_type is called when exiting the floating_pt_type production.
	ExitFloating_pt_type(c *Floating_pt_typeContext)

	// ExitInteger_type is called when exiting the integer_type production.
	ExitInteger_type(c *Integer_typeContext)

	// ExitSigned_int is called when exiting the signed_int production.
	ExitSigned_int(c *Signed_intContext)

	// ExitSigned_tiny_int is called when exiting the signed_tiny_int production.
	ExitSigned_tiny_int(c *Signed_tiny_intContext)

	// ExitSigned_short_int is called when exiting the signed_short_int production.
	ExitSigned_short_int(c *Signed_short_intContext)

	// ExitSigned_long_int is called when exiting the signed_long_int production.
	ExitSigned_long_int(c *Signed_long_intContext)

	// ExitSigned_longlong_int is called when exiting the signed_longlong_int production.
	ExitSigned_longlong_int(c *Signed_longlong_intContext)

	// ExitUnsigned_int is called when exiting the unsigned_int production.
	ExitUnsigned_int(c *Unsigned_intContext)

	// ExitUnsigned_tiny_int is called when exiting the unsigned_tiny_int production.
	ExitUnsigned_tiny_int(c *Unsigned_tiny_intContext)

	// ExitUnsigned_short_int is called when exiting the unsigned_short_int production.
	ExitUnsigned_short_int(c *Unsigned_short_intContext)

	// ExitUnsigned_long_int is called when exiting the unsigned_long_int production.
	ExitUnsigned_long_int(c *Unsigned_long_intContext)

	// ExitUnsigned_longlong_int is called when exiting the unsigned_longlong_int production.
	ExitUnsigned_longlong_int(c *Unsigned_longlong_intContext)

	// ExitChar_type is called when exiting the char_type production.
	ExitChar_type(c *Char_typeContext)

	// ExitWide_char_type is called when exiting the wide_char_type production.
	ExitWide_char_type(c *Wide_char_typeContext)

	// ExitBoolean_type is called when exiting the boolean_type production.
	ExitBoolean_type(c *Boolean_typeContext)

	// ExitOctet_type is called when exiting the octet_type production.
	ExitOctet_type(c *Octet_typeContext)

	// ExitAny_type is called when exiting the any_type production.
	ExitAny_type(c *Any_typeContext)

	// ExitObject_type is called when exiting the object_type production.
	ExitObject_type(c *Object_typeContext)

	// ExitAnnotation_decl is called when exiting the annotation_decl production.
	ExitAnnotation_decl(c *Annotation_declContext)

	// ExitAnnotation_def is called when exiting the annotation_def production.
	ExitAnnotation_def(c *Annotation_defContext)

	// ExitAnnotation_header is called when exiting the annotation_header production.
	ExitAnnotation_header(c *Annotation_headerContext)

	// ExitAnnotation_inheritance_spec is called when exiting the annotation_inheritance_spec production.
	ExitAnnotation_inheritance_spec(c *Annotation_inheritance_specContext)

	// ExitAnnotation_body is called when exiting the annotation_body production.
	ExitAnnotation_body(c *Annotation_bodyContext)

	// ExitAnnotation_member is called when exiting the annotation_member production.
	ExitAnnotation_member(c *Annotation_memberContext)

	// ExitAnnotation_forward_dcl is called when exiting the annotation_forward_dcl production.
	ExitAnnotation_forward_dcl(c *Annotation_forward_dclContext)

	// ExitBitset_type is called when exiting the bitset_type production.
	ExitBitset_type(c *Bitset_typeContext)

	// ExitBitfield is called when exiting the bitfield production.
	ExitBitfield(c *BitfieldContext)

	// ExitBitfield_spec is called when exiting the bitfield_spec production.
	ExitBitfield_spec(c *Bitfield_specContext)

	// ExitBitmask_type is called when exiting the bitmask_type production.
	ExitBitmask_type(c *Bitmask_typeContext)

	// ExitBit_values is called when exiting the bit_values production.
	ExitBit_values(c *Bit_valuesContext)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitMember_list is called when exiting the member_list production.
	ExitMember_list(c *Member_listContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitUnion_type is called when exiting the union_type production.
	ExitUnion_type(c *Union_typeContext)

	// ExitSwitch_type_spec is called when exiting the switch_type_spec production.
	ExitSwitch_type_spec(c *Switch_type_specContext)

	// ExitSwitch_body is called when exiting the switch_body production.
	ExitSwitch_body(c *Switch_bodyContext)

	// ExitCase_stmt is called when exiting the case_stmt production.
	ExitCase_stmt(c *Case_stmtContext)

	// ExitCase_label is called when exiting the case_label production.
	ExitCase_label(c *Case_labelContext)

	// ExitElement_spec is called when exiting the element_spec production.
	ExitElement_spec(c *Element_specContext)

	// ExitEnum_type is called when exiting the enum_type production.
	ExitEnum_type(c *Enum_typeContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitSequence_type is called when exiting the sequence_type production.
	ExitSequence_type(c *Sequence_typeContext)

	// ExitSet_type is called when exiting the set_type production.
	ExitSet_type(c *Set_typeContext)

	// ExitMap_type is called when exiting the map_type production.
	ExitMap_type(c *Map_typeContext)

	// ExitString_type is called when exiting the string_type production.
	ExitString_type(c *String_typeContext)

	// ExitWide_string_type is called when exiting the wide_string_type production.
	ExitWide_string_type(c *Wide_string_typeContext)

	// ExitArray_declarator is called when exiting the array_declarator production.
	ExitArray_declarator(c *Array_declaratorContext)

	// ExitFixed_array_size is called when exiting the fixed_array_size production.
	ExitFixed_array_size(c *Fixed_array_sizeContext)

	// ExitAttr_decl is called when exiting the attr_decl production.
	ExitAttr_decl(c *Attr_declContext)

	// ExitExcept_decl is called when exiting the except_decl production.
	ExitExcept_decl(c *Except_declContext)

	// ExitOp_decl is called when exiting the op_decl production.
	ExitOp_decl(c *Op_declContext)

	// ExitOp_attribute is called when exiting the op_attribute production.
	ExitOp_attribute(c *Op_attributeContext)

	// ExitOp_type_spec is called when exiting the op_type_spec production.
	ExitOp_type_spec(c *Op_type_specContext)

	// ExitParameter_decls is called when exiting the parameter_decls production.
	ExitParameter_decls(c *Parameter_declsContext)

	// ExitParam_decl is called when exiting the param_decl production.
	ExitParam_decl(c *Param_declContext)

	// ExitParam_attribute is called when exiting the param_attribute production.
	ExitParam_attribute(c *Param_attributeContext)

	// ExitRaises_expr is called when exiting the raises_expr production.
	ExitRaises_expr(c *Raises_exprContext)

	// ExitContext_expr is called when exiting the context_expr production.
	ExitContext_expr(c *Context_exprContext)

	// ExitParam_type_spec is called when exiting the param_type_spec production.
	ExitParam_type_spec(c *Param_type_specContext)

	// ExitFixed_pt_type is called when exiting the fixed_pt_type production.
	ExitFixed_pt_type(c *Fixed_pt_typeContext)

	// ExitFixed_pt_const_type is called when exiting the fixed_pt_const_type production.
	ExitFixed_pt_const_type(c *Fixed_pt_const_typeContext)

	// ExitValue_base_type is called when exiting the value_base_type production.
	ExitValue_base_type(c *Value_base_typeContext)

	// ExitConstr_forward_decl is called when exiting the constr_forward_decl production.
	ExitConstr_forward_decl(c *Constr_forward_declContext)

	// ExitImport_decl is called when exiting the import_decl production.
	ExitImport_decl(c *Import_declContext)

	// ExitImported_scope is called when exiting the imported_scope production.
	ExitImported_scope(c *Imported_scopeContext)

	// ExitType_id_decl is called when exiting the type_id_decl production.
	ExitType_id_decl(c *Type_id_declContext)

	// ExitType_prefix_decl is called when exiting the type_prefix_decl production.
	ExitType_prefix_decl(c *Type_prefix_declContext)

	// ExitReadonly_attr_spec is called when exiting the readonly_attr_spec production.
	ExitReadonly_attr_spec(c *Readonly_attr_specContext)

	// ExitReadonly_attr_declarator is called when exiting the readonly_attr_declarator production.
	ExitReadonly_attr_declarator(c *Readonly_attr_declaratorContext)

	// ExitAttr_spec is called when exiting the attr_spec production.
	ExitAttr_spec(c *Attr_specContext)

	// ExitAttr_declarator is called when exiting the attr_declarator production.
	ExitAttr_declarator(c *Attr_declaratorContext)

	// ExitAttr_raises_expr is called when exiting the attr_raises_expr production.
	ExitAttr_raises_expr(c *Attr_raises_exprContext)

	// ExitGet_excep_expr is called when exiting the get_excep_expr production.
	ExitGet_excep_expr(c *Get_excep_exprContext)

	// ExitSet_excep_expr is called when exiting the set_excep_expr production.
	ExitSet_excep_expr(c *Set_excep_exprContext)

	// ExitException_list is called when exiting the exception_list production.
	ExitException_list(c *Exception_listContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitComponent_forward_decl is called when exiting the component_forward_decl production.
	ExitComponent_forward_decl(c *Component_forward_declContext)

	// ExitComponent_decl is called when exiting the component_decl production.
	ExitComponent_decl(c *Component_declContext)

	// ExitComponent_header is called when exiting the component_header production.
	ExitComponent_header(c *Component_headerContext)

	// ExitSupported_interface_spec is called when exiting the supported_interface_spec production.
	ExitSupported_interface_spec(c *Supported_interface_specContext)

	// ExitComponent_inheritance_spec is called when exiting the component_inheritance_spec production.
	ExitComponent_inheritance_spec(c *Component_inheritance_specContext)

	// ExitComponent_body is called when exiting the component_body production.
	ExitComponent_body(c *Component_bodyContext)

	// ExitComponent_export is called when exiting the component_export production.
	ExitComponent_export(c *Component_exportContext)

	// ExitProvides_decl is called when exiting the provides_decl production.
	ExitProvides_decl(c *Provides_declContext)

	// ExitInterface_type is called when exiting the interface_type production.
	ExitInterface_type(c *Interface_typeContext)

	// ExitUses_decl is called when exiting the uses_decl production.
	ExitUses_decl(c *Uses_declContext)

	// ExitEmits_decl is called when exiting the emits_decl production.
	ExitEmits_decl(c *Emits_declContext)

	// ExitPublishes_decl is called when exiting the publishes_decl production.
	ExitPublishes_decl(c *Publishes_declContext)

	// ExitConsumes_decl is called when exiting the consumes_decl production.
	ExitConsumes_decl(c *Consumes_declContext)

	// ExitHome_decl is called when exiting the home_decl production.
	ExitHome_decl(c *Home_declContext)

	// ExitHome_header is called when exiting the home_header production.
	ExitHome_header(c *Home_headerContext)

	// ExitHome_inheritance_spec is called when exiting the home_inheritance_spec production.
	ExitHome_inheritance_spec(c *Home_inheritance_specContext)

	// ExitPrimary_key_spec is called when exiting the primary_key_spec production.
	ExitPrimary_key_spec(c *Primary_key_specContext)

	// ExitHome_body is called when exiting the home_body production.
	ExitHome_body(c *Home_bodyContext)

	// ExitHome_export is called when exiting the home_export production.
	ExitHome_export(c *Home_exportContext)

	// ExitFactory_decl is called when exiting the factory_decl production.
	ExitFactory_decl(c *Factory_declContext)

	// ExitFinder_decl is called when exiting the finder_decl production.
	ExitFinder_decl(c *Finder_declContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitEvent_forward_decl is called when exiting the event_forward_decl production.
	ExitEvent_forward_decl(c *Event_forward_declContext)

	// ExitEvent_abs_decl is called when exiting the event_abs_decl production.
	ExitEvent_abs_decl(c *Event_abs_declContext)

	// ExitEvent_decl is called when exiting the event_decl production.
	ExitEvent_decl(c *Event_declContext)

	// ExitEvent_header is called when exiting the event_header production.
	ExitEvent_header(c *Event_headerContext)

	// ExitAnnapps is called when exiting the annapps production.
	ExitAnnapps(c *AnnappsContext)

	// ExitAnnotation_appl is called when exiting the annotation_appl production.
	ExitAnnotation_appl(c *Annotation_applContext)

	// ExitAnnotation_appl_params is called when exiting the annotation_appl_params production.
	ExitAnnotation_appl_params(c *Annotation_appl_paramsContext)

	// ExitAnnotation_appl_param is called when exiting the annotation_appl_param production.
	ExitAnnotation_appl_param(c *Annotation_appl_paramContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
