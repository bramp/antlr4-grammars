// Generated from IDL.g4 by ANTLR 4.7.

package idl // IDL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIDLListener is a complete listener for a parse tree produced by IDLParser.
type BaseIDLListener struct{}

var _ IDLListener = &BaseIDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSpecification is called when production specification is entered.
func (s *BaseIDLListener) EnterSpecification(ctx *SpecificationContext) {}

// ExitSpecification is called when production specification is exited.
func (s *BaseIDLListener) ExitSpecification(ctx *SpecificationContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseIDLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseIDLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterModule is called when production module is entered.
func (s *BaseIDLListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseIDLListener) ExitModule(ctx *ModuleContext) {}

// EnterInterface_or_forward_decl is called when production interface_or_forward_decl is entered.
func (s *BaseIDLListener) EnterInterface_or_forward_decl(ctx *Interface_or_forward_declContext) {}

// ExitInterface_or_forward_decl is called when production interface_or_forward_decl is exited.
func (s *BaseIDLListener) ExitInterface_or_forward_decl(ctx *Interface_or_forward_declContext) {}

// EnterInterface_decl is called when production interface_decl is entered.
func (s *BaseIDLListener) EnterInterface_decl(ctx *Interface_declContext) {}

// ExitInterface_decl is called when production interface_decl is exited.
func (s *BaseIDLListener) ExitInterface_decl(ctx *Interface_declContext) {}

// EnterForward_decl is called when production forward_decl is entered.
func (s *BaseIDLListener) EnterForward_decl(ctx *Forward_declContext) {}

// ExitForward_decl is called when production forward_decl is exited.
func (s *BaseIDLListener) ExitForward_decl(ctx *Forward_declContext) {}

// EnterInterface_header is called when production interface_header is entered.
func (s *BaseIDLListener) EnterInterface_header(ctx *Interface_headerContext) {}

// ExitInterface_header is called when production interface_header is exited.
func (s *BaseIDLListener) ExitInterface_header(ctx *Interface_headerContext) {}

// EnterInterface_body is called when production interface_body is entered.
func (s *BaseIDLListener) EnterInterface_body(ctx *Interface_bodyContext) {}

// ExitInterface_body is called when production interface_body is exited.
func (s *BaseIDLListener) ExitInterface_body(ctx *Interface_bodyContext) {}

// EnterExport is called when production export is entered.
func (s *BaseIDLListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseIDLListener) ExitExport(ctx *ExportContext) {}

// EnterInterface_inheritance_spec is called when production interface_inheritance_spec is entered.
func (s *BaseIDLListener) EnterInterface_inheritance_spec(ctx *Interface_inheritance_specContext) {}

// ExitInterface_inheritance_spec is called when production interface_inheritance_spec is exited.
func (s *BaseIDLListener) ExitInterface_inheritance_spec(ctx *Interface_inheritance_specContext) {}

// EnterInterface_name is called when production interface_name is entered.
func (s *BaseIDLListener) EnterInterface_name(ctx *Interface_nameContext) {}

// ExitInterface_name is called when production interface_name is exited.
func (s *BaseIDLListener) ExitInterface_name(ctx *Interface_nameContext) {}

// EnterScoped_name is called when production scoped_name is entered.
func (s *BaseIDLListener) EnterScoped_name(ctx *Scoped_nameContext) {}

// ExitScoped_name is called when production scoped_name is exited.
func (s *BaseIDLListener) ExitScoped_name(ctx *Scoped_nameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseIDLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseIDLListener) ExitValue(ctx *ValueContext) {}

// EnterValue_forward_decl is called when production value_forward_decl is entered.
func (s *BaseIDLListener) EnterValue_forward_decl(ctx *Value_forward_declContext) {}

// ExitValue_forward_decl is called when production value_forward_decl is exited.
func (s *BaseIDLListener) ExitValue_forward_decl(ctx *Value_forward_declContext) {}

// EnterValue_box_decl is called when production value_box_decl is entered.
func (s *BaseIDLListener) EnterValue_box_decl(ctx *Value_box_declContext) {}

// ExitValue_box_decl is called when production value_box_decl is exited.
func (s *BaseIDLListener) ExitValue_box_decl(ctx *Value_box_declContext) {}

// EnterValue_abs_decl is called when production value_abs_decl is entered.
func (s *BaseIDLListener) EnterValue_abs_decl(ctx *Value_abs_declContext) {}

// ExitValue_abs_decl is called when production value_abs_decl is exited.
func (s *BaseIDLListener) ExitValue_abs_decl(ctx *Value_abs_declContext) {}

// EnterValue_decl is called when production value_decl is entered.
func (s *BaseIDLListener) EnterValue_decl(ctx *Value_declContext) {}

// ExitValue_decl is called when production value_decl is exited.
func (s *BaseIDLListener) ExitValue_decl(ctx *Value_declContext) {}

// EnterValue_header is called when production value_header is entered.
func (s *BaseIDLListener) EnterValue_header(ctx *Value_headerContext) {}

// ExitValue_header is called when production value_header is exited.
func (s *BaseIDLListener) ExitValue_header(ctx *Value_headerContext) {}

// EnterValue_inheritance_spec is called when production value_inheritance_spec is entered.
func (s *BaseIDLListener) EnterValue_inheritance_spec(ctx *Value_inheritance_specContext) {}

// ExitValue_inheritance_spec is called when production value_inheritance_spec is exited.
func (s *BaseIDLListener) ExitValue_inheritance_spec(ctx *Value_inheritance_specContext) {}

// EnterValue_name is called when production value_name is entered.
func (s *BaseIDLListener) EnterValue_name(ctx *Value_nameContext) {}

// ExitValue_name is called when production value_name is exited.
func (s *BaseIDLListener) ExitValue_name(ctx *Value_nameContext) {}

// EnterValue_element is called when production value_element is entered.
func (s *BaseIDLListener) EnterValue_element(ctx *Value_elementContext) {}

// ExitValue_element is called when production value_element is exited.
func (s *BaseIDLListener) ExitValue_element(ctx *Value_elementContext) {}

// EnterState_member is called when production state_member is entered.
func (s *BaseIDLListener) EnterState_member(ctx *State_memberContext) {}

// ExitState_member is called when production state_member is exited.
func (s *BaseIDLListener) ExitState_member(ctx *State_memberContext) {}

// EnterInit_decl is called when production init_decl is entered.
func (s *BaseIDLListener) EnterInit_decl(ctx *Init_declContext) {}

// ExitInit_decl is called when production init_decl is exited.
func (s *BaseIDLListener) ExitInit_decl(ctx *Init_declContext) {}

// EnterInit_param_decls is called when production init_param_decls is entered.
func (s *BaseIDLListener) EnterInit_param_decls(ctx *Init_param_declsContext) {}

// ExitInit_param_decls is called when production init_param_decls is exited.
func (s *BaseIDLListener) ExitInit_param_decls(ctx *Init_param_declsContext) {}

// EnterInit_param_decl is called when production init_param_decl is entered.
func (s *BaseIDLListener) EnterInit_param_decl(ctx *Init_param_declContext) {}

// ExitInit_param_decl is called when production init_param_decl is exited.
func (s *BaseIDLListener) ExitInit_param_decl(ctx *Init_param_declContext) {}

// EnterInit_param_attribute is called when production init_param_attribute is entered.
func (s *BaseIDLListener) EnterInit_param_attribute(ctx *Init_param_attributeContext) {}

// ExitInit_param_attribute is called when production init_param_attribute is exited.
func (s *BaseIDLListener) ExitInit_param_attribute(ctx *Init_param_attributeContext) {}

// EnterConst_decl is called when production const_decl is entered.
func (s *BaseIDLListener) EnterConst_decl(ctx *Const_declContext) {}

// ExitConst_decl is called when production const_decl is exited.
func (s *BaseIDLListener) ExitConst_decl(ctx *Const_declContext) {}

// EnterConst_type is called when production const_type is entered.
func (s *BaseIDLListener) EnterConst_type(ctx *Const_typeContext) {}

// ExitConst_type is called when production const_type is exited.
func (s *BaseIDLListener) ExitConst_type(ctx *Const_typeContext) {}

// EnterConst_exp is called when production const_exp is entered.
func (s *BaseIDLListener) EnterConst_exp(ctx *Const_expContext) {}

// ExitConst_exp is called when production const_exp is exited.
func (s *BaseIDLListener) ExitConst_exp(ctx *Const_expContext) {}

// EnterOr_expr is called when production or_expr is entered.
func (s *BaseIDLListener) EnterOr_expr(ctx *Or_exprContext) {}

// ExitOr_expr is called when production or_expr is exited.
func (s *BaseIDLListener) ExitOr_expr(ctx *Or_exprContext) {}

// EnterXor_expr is called when production xor_expr is entered.
func (s *BaseIDLListener) EnterXor_expr(ctx *Xor_exprContext) {}

// ExitXor_expr is called when production xor_expr is exited.
func (s *BaseIDLListener) ExitXor_expr(ctx *Xor_exprContext) {}

// EnterAnd_expr is called when production and_expr is entered.
func (s *BaseIDLListener) EnterAnd_expr(ctx *And_exprContext) {}

// ExitAnd_expr is called when production and_expr is exited.
func (s *BaseIDLListener) ExitAnd_expr(ctx *And_exprContext) {}

// EnterShift_expr is called when production shift_expr is entered.
func (s *BaseIDLListener) EnterShift_expr(ctx *Shift_exprContext) {}

// ExitShift_expr is called when production shift_expr is exited.
func (s *BaseIDLListener) ExitShift_expr(ctx *Shift_exprContext) {}

// EnterAdd_expr is called when production add_expr is entered.
func (s *BaseIDLListener) EnterAdd_expr(ctx *Add_exprContext) {}

// ExitAdd_expr is called when production add_expr is exited.
func (s *BaseIDLListener) ExitAdd_expr(ctx *Add_exprContext) {}

// EnterMult_expr is called when production mult_expr is entered.
func (s *BaseIDLListener) EnterMult_expr(ctx *Mult_exprContext) {}

// ExitMult_expr is called when production mult_expr is exited.
func (s *BaseIDLListener) ExitMult_expr(ctx *Mult_exprContext) {}

// EnterUnary_expr is called when production unary_expr is entered.
func (s *BaseIDLListener) EnterUnary_expr(ctx *Unary_exprContext) {}

// ExitUnary_expr is called when production unary_expr is exited.
func (s *BaseIDLListener) ExitUnary_expr(ctx *Unary_exprContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseIDLListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseIDLListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterPrimary_expr is called when production primary_expr is entered.
func (s *BaseIDLListener) EnterPrimary_expr(ctx *Primary_exprContext) {}

// ExitPrimary_expr is called when production primary_expr is exited.
func (s *BaseIDLListener) ExitPrimary_expr(ctx *Primary_exprContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseIDLListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseIDLListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPositive_int_const is called when production positive_int_const is entered.
func (s *BaseIDLListener) EnterPositive_int_const(ctx *Positive_int_constContext) {}

// ExitPositive_int_const is called when production positive_int_const is exited.
func (s *BaseIDLListener) ExitPositive_int_const(ctx *Positive_int_constContext) {}

// EnterType_decl is called when production type_decl is entered.
func (s *BaseIDLListener) EnterType_decl(ctx *Type_declContext) {}

// ExitType_decl is called when production type_decl is exited.
func (s *BaseIDLListener) ExitType_decl(ctx *Type_declContext) {}

// EnterType_declarator is called when production type_declarator is entered.
func (s *BaseIDLListener) EnterType_declarator(ctx *Type_declaratorContext) {}

// ExitType_declarator is called when production type_declarator is exited.
func (s *BaseIDLListener) ExitType_declarator(ctx *Type_declaratorContext) {}

// EnterType_spec is called when production type_spec is entered.
func (s *BaseIDLListener) EnterType_spec(ctx *Type_specContext) {}

// ExitType_spec is called when production type_spec is exited.
func (s *BaseIDLListener) ExitType_spec(ctx *Type_specContext) {}

// EnterSimple_type_spec is called when production simple_type_spec is entered.
func (s *BaseIDLListener) EnterSimple_type_spec(ctx *Simple_type_specContext) {}

// ExitSimple_type_spec is called when production simple_type_spec is exited.
func (s *BaseIDLListener) ExitSimple_type_spec(ctx *Simple_type_specContext) {}

// EnterBase_type_spec is called when production base_type_spec is entered.
func (s *BaseIDLListener) EnterBase_type_spec(ctx *Base_type_specContext) {}

// ExitBase_type_spec is called when production base_type_spec is exited.
func (s *BaseIDLListener) ExitBase_type_spec(ctx *Base_type_specContext) {}

// EnterTemplate_type_spec is called when production template_type_spec is entered.
func (s *BaseIDLListener) EnterTemplate_type_spec(ctx *Template_type_specContext) {}

// ExitTemplate_type_spec is called when production template_type_spec is exited.
func (s *BaseIDLListener) ExitTemplate_type_spec(ctx *Template_type_specContext) {}

// EnterConstr_type_spec is called when production constr_type_spec is entered.
func (s *BaseIDLListener) EnterConstr_type_spec(ctx *Constr_type_specContext) {}

// ExitConstr_type_spec is called when production constr_type_spec is exited.
func (s *BaseIDLListener) ExitConstr_type_spec(ctx *Constr_type_specContext) {}

// EnterDeclarators is called when production declarators is entered.
func (s *BaseIDLListener) EnterDeclarators(ctx *DeclaratorsContext) {}

// ExitDeclarators is called when production declarators is exited.
func (s *BaseIDLListener) ExitDeclarators(ctx *DeclaratorsContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseIDLListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseIDLListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterSimple_declarator is called when production simple_declarator is entered.
func (s *BaseIDLListener) EnterSimple_declarator(ctx *Simple_declaratorContext) {}

// ExitSimple_declarator is called when production simple_declarator is exited.
func (s *BaseIDLListener) ExitSimple_declarator(ctx *Simple_declaratorContext) {}

// EnterComplex_declarator is called when production complex_declarator is entered.
func (s *BaseIDLListener) EnterComplex_declarator(ctx *Complex_declaratorContext) {}

// ExitComplex_declarator is called when production complex_declarator is exited.
func (s *BaseIDLListener) ExitComplex_declarator(ctx *Complex_declaratorContext) {}

// EnterFloating_pt_type is called when production floating_pt_type is entered.
func (s *BaseIDLListener) EnterFloating_pt_type(ctx *Floating_pt_typeContext) {}

// ExitFloating_pt_type is called when production floating_pt_type is exited.
func (s *BaseIDLListener) ExitFloating_pt_type(ctx *Floating_pt_typeContext) {}

// EnterInteger_type is called when production integer_type is entered.
func (s *BaseIDLListener) EnterInteger_type(ctx *Integer_typeContext) {}

// ExitInteger_type is called when production integer_type is exited.
func (s *BaseIDLListener) ExitInteger_type(ctx *Integer_typeContext) {}

// EnterSigned_int is called when production signed_int is entered.
func (s *BaseIDLListener) EnterSigned_int(ctx *Signed_intContext) {}

// ExitSigned_int is called when production signed_int is exited.
func (s *BaseIDLListener) ExitSigned_int(ctx *Signed_intContext) {}

// EnterSigned_short_int is called when production signed_short_int is entered.
func (s *BaseIDLListener) EnterSigned_short_int(ctx *Signed_short_intContext) {}

// ExitSigned_short_int is called when production signed_short_int is exited.
func (s *BaseIDLListener) ExitSigned_short_int(ctx *Signed_short_intContext) {}

// EnterSigned_long_int is called when production signed_long_int is entered.
func (s *BaseIDLListener) EnterSigned_long_int(ctx *Signed_long_intContext) {}

// ExitSigned_long_int is called when production signed_long_int is exited.
func (s *BaseIDLListener) ExitSigned_long_int(ctx *Signed_long_intContext) {}

// EnterSigned_longlong_int is called when production signed_longlong_int is entered.
func (s *BaseIDLListener) EnterSigned_longlong_int(ctx *Signed_longlong_intContext) {}

// ExitSigned_longlong_int is called when production signed_longlong_int is exited.
func (s *BaseIDLListener) ExitSigned_longlong_int(ctx *Signed_longlong_intContext) {}

// EnterUnsigned_int is called when production unsigned_int is entered.
func (s *BaseIDLListener) EnterUnsigned_int(ctx *Unsigned_intContext) {}

// ExitUnsigned_int is called when production unsigned_int is exited.
func (s *BaseIDLListener) ExitUnsigned_int(ctx *Unsigned_intContext) {}

// EnterUnsigned_short_int is called when production unsigned_short_int is entered.
func (s *BaseIDLListener) EnterUnsigned_short_int(ctx *Unsigned_short_intContext) {}

// ExitUnsigned_short_int is called when production unsigned_short_int is exited.
func (s *BaseIDLListener) ExitUnsigned_short_int(ctx *Unsigned_short_intContext) {}

// EnterUnsigned_long_int is called when production unsigned_long_int is entered.
func (s *BaseIDLListener) EnterUnsigned_long_int(ctx *Unsigned_long_intContext) {}

// ExitUnsigned_long_int is called when production unsigned_long_int is exited.
func (s *BaseIDLListener) ExitUnsigned_long_int(ctx *Unsigned_long_intContext) {}

// EnterUnsigned_longlong_int is called when production unsigned_longlong_int is entered.
func (s *BaseIDLListener) EnterUnsigned_longlong_int(ctx *Unsigned_longlong_intContext) {}

// ExitUnsigned_longlong_int is called when production unsigned_longlong_int is exited.
func (s *BaseIDLListener) ExitUnsigned_longlong_int(ctx *Unsigned_longlong_intContext) {}

// EnterChar_type is called when production char_type is entered.
func (s *BaseIDLListener) EnterChar_type(ctx *Char_typeContext) {}

// ExitChar_type is called when production char_type is exited.
func (s *BaseIDLListener) ExitChar_type(ctx *Char_typeContext) {}

// EnterWide_char_type is called when production wide_char_type is entered.
func (s *BaseIDLListener) EnterWide_char_type(ctx *Wide_char_typeContext) {}

// ExitWide_char_type is called when production wide_char_type is exited.
func (s *BaseIDLListener) ExitWide_char_type(ctx *Wide_char_typeContext) {}

// EnterBoolean_type is called when production boolean_type is entered.
func (s *BaseIDLListener) EnterBoolean_type(ctx *Boolean_typeContext) {}

// ExitBoolean_type is called when production boolean_type is exited.
func (s *BaseIDLListener) ExitBoolean_type(ctx *Boolean_typeContext) {}

// EnterOctet_type is called when production octet_type is entered.
func (s *BaseIDLListener) EnterOctet_type(ctx *Octet_typeContext) {}

// ExitOctet_type is called when production octet_type is exited.
func (s *BaseIDLListener) ExitOctet_type(ctx *Octet_typeContext) {}

// EnterAny_type is called when production any_type is entered.
func (s *BaseIDLListener) EnterAny_type(ctx *Any_typeContext) {}

// ExitAny_type is called when production any_type is exited.
func (s *BaseIDLListener) ExitAny_type(ctx *Any_typeContext) {}

// EnterObject_type is called when production object_type is entered.
func (s *BaseIDLListener) EnterObject_type(ctx *Object_typeContext) {}

// ExitObject_type is called when production object_type is exited.
func (s *BaseIDLListener) ExitObject_type(ctx *Object_typeContext) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseIDLListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseIDLListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterMember_list is called when production member_list is entered.
func (s *BaseIDLListener) EnterMember_list(ctx *Member_listContext) {}

// ExitMember_list is called when production member_list is exited.
func (s *BaseIDLListener) ExitMember_list(ctx *Member_listContext) {}

// EnterMember is called when production member is entered.
func (s *BaseIDLListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseIDLListener) ExitMember(ctx *MemberContext) {}

// EnterUnion_type is called when production union_type is entered.
func (s *BaseIDLListener) EnterUnion_type(ctx *Union_typeContext) {}

// ExitUnion_type is called when production union_type is exited.
func (s *BaseIDLListener) ExitUnion_type(ctx *Union_typeContext) {}

// EnterSwitch_type_spec is called when production switch_type_spec is entered.
func (s *BaseIDLListener) EnterSwitch_type_spec(ctx *Switch_type_specContext) {}

// ExitSwitch_type_spec is called when production switch_type_spec is exited.
func (s *BaseIDLListener) ExitSwitch_type_spec(ctx *Switch_type_specContext) {}

// EnterSwitch_body is called when production switch_body is entered.
func (s *BaseIDLListener) EnterSwitch_body(ctx *Switch_bodyContext) {}

// ExitSwitch_body is called when production switch_body is exited.
func (s *BaseIDLListener) ExitSwitch_body(ctx *Switch_bodyContext) {}

// EnterCase_stmt is called when production case_stmt is entered.
func (s *BaseIDLListener) EnterCase_stmt(ctx *Case_stmtContext) {}

// ExitCase_stmt is called when production case_stmt is exited.
func (s *BaseIDLListener) ExitCase_stmt(ctx *Case_stmtContext) {}

// EnterCase_label is called when production case_label is entered.
func (s *BaseIDLListener) EnterCase_label(ctx *Case_labelContext) {}

// ExitCase_label is called when production case_label is exited.
func (s *BaseIDLListener) ExitCase_label(ctx *Case_labelContext) {}

// EnterElement_spec is called when production element_spec is entered.
func (s *BaseIDLListener) EnterElement_spec(ctx *Element_specContext) {}

// ExitElement_spec is called when production element_spec is exited.
func (s *BaseIDLListener) ExitElement_spec(ctx *Element_specContext) {}

// EnterEnum_type is called when production enum_type is entered.
func (s *BaseIDLListener) EnterEnum_type(ctx *Enum_typeContext) {}

// ExitEnum_type is called when production enum_type is exited.
func (s *BaseIDLListener) ExitEnum_type(ctx *Enum_typeContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseIDLListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseIDLListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterSequence_type is called when production sequence_type is entered.
func (s *BaseIDLListener) EnterSequence_type(ctx *Sequence_typeContext) {}

// ExitSequence_type is called when production sequence_type is exited.
func (s *BaseIDLListener) ExitSequence_type(ctx *Sequence_typeContext) {}

// EnterString_type is called when production string_type is entered.
func (s *BaseIDLListener) EnterString_type(ctx *String_typeContext) {}

// ExitString_type is called when production string_type is exited.
func (s *BaseIDLListener) ExitString_type(ctx *String_typeContext) {}

// EnterWide_string_type is called when production wide_string_type is entered.
func (s *BaseIDLListener) EnterWide_string_type(ctx *Wide_string_typeContext) {}

// ExitWide_string_type is called when production wide_string_type is exited.
func (s *BaseIDLListener) ExitWide_string_type(ctx *Wide_string_typeContext) {}

// EnterArray_declarator is called when production array_declarator is entered.
func (s *BaseIDLListener) EnterArray_declarator(ctx *Array_declaratorContext) {}

// ExitArray_declarator is called when production array_declarator is exited.
func (s *BaseIDLListener) ExitArray_declarator(ctx *Array_declaratorContext) {}

// EnterFixed_array_size is called when production fixed_array_size is entered.
func (s *BaseIDLListener) EnterFixed_array_size(ctx *Fixed_array_sizeContext) {}

// ExitFixed_array_size is called when production fixed_array_size is exited.
func (s *BaseIDLListener) ExitFixed_array_size(ctx *Fixed_array_sizeContext) {}

// EnterAttr_decl is called when production attr_decl is entered.
func (s *BaseIDLListener) EnterAttr_decl(ctx *Attr_declContext) {}

// ExitAttr_decl is called when production attr_decl is exited.
func (s *BaseIDLListener) ExitAttr_decl(ctx *Attr_declContext) {}

// EnterExcept_decl is called when production except_decl is entered.
func (s *BaseIDLListener) EnterExcept_decl(ctx *Except_declContext) {}

// ExitExcept_decl is called when production except_decl is exited.
func (s *BaseIDLListener) ExitExcept_decl(ctx *Except_declContext) {}

// EnterOp_decl is called when production op_decl is entered.
func (s *BaseIDLListener) EnterOp_decl(ctx *Op_declContext) {}

// ExitOp_decl is called when production op_decl is exited.
func (s *BaseIDLListener) ExitOp_decl(ctx *Op_declContext) {}

// EnterOp_attribute is called when production op_attribute is entered.
func (s *BaseIDLListener) EnterOp_attribute(ctx *Op_attributeContext) {}

// ExitOp_attribute is called when production op_attribute is exited.
func (s *BaseIDLListener) ExitOp_attribute(ctx *Op_attributeContext) {}

// EnterOp_type_spec is called when production op_type_spec is entered.
func (s *BaseIDLListener) EnterOp_type_spec(ctx *Op_type_specContext) {}

// ExitOp_type_spec is called when production op_type_spec is exited.
func (s *BaseIDLListener) ExitOp_type_spec(ctx *Op_type_specContext) {}

// EnterParameter_decls is called when production parameter_decls is entered.
func (s *BaseIDLListener) EnterParameter_decls(ctx *Parameter_declsContext) {}

// ExitParameter_decls is called when production parameter_decls is exited.
func (s *BaseIDLListener) ExitParameter_decls(ctx *Parameter_declsContext) {}

// EnterParam_decl is called when production param_decl is entered.
func (s *BaseIDLListener) EnterParam_decl(ctx *Param_declContext) {}

// ExitParam_decl is called when production param_decl is exited.
func (s *BaseIDLListener) ExitParam_decl(ctx *Param_declContext) {}

// EnterParam_attribute is called when production param_attribute is entered.
func (s *BaseIDLListener) EnterParam_attribute(ctx *Param_attributeContext) {}

// ExitParam_attribute is called when production param_attribute is exited.
func (s *BaseIDLListener) ExitParam_attribute(ctx *Param_attributeContext) {}

// EnterRaises_expr is called when production raises_expr is entered.
func (s *BaseIDLListener) EnterRaises_expr(ctx *Raises_exprContext) {}

// ExitRaises_expr is called when production raises_expr is exited.
func (s *BaseIDLListener) ExitRaises_expr(ctx *Raises_exprContext) {}

// EnterContext_expr is called when production context_expr is entered.
func (s *BaseIDLListener) EnterContext_expr(ctx *Context_exprContext) {}

// ExitContext_expr is called when production context_expr is exited.
func (s *BaseIDLListener) ExitContext_expr(ctx *Context_exprContext) {}

// EnterParam_type_spec is called when production param_type_spec is entered.
func (s *BaseIDLListener) EnterParam_type_spec(ctx *Param_type_specContext) {}

// ExitParam_type_spec is called when production param_type_spec is exited.
func (s *BaseIDLListener) ExitParam_type_spec(ctx *Param_type_specContext) {}

// EnterFixed_pt_type is called when production fixed_pt_type is entered.
func (s *BaseIDLListener) EnterFixed_pt_type(ctx *Fixed_pt_typeContext) {}

// ExitFixed_pt_type is called when production fixed_pt_type is exited.
func (s *BaseIDLListener) ExitFixed_pt_type(ctx *Fixed_pt_typeContext) {}

// EnterFixed_pt_const_type is called when production fixed_pt_const_type is entered.
func (s *BaseIDLListener) EnterFixed_pt_const_type(ctx *Fixed_pt_const_typeContext) {}

// ExitFixed_pt_const_type is called when production fixed_pt_const_type is exited.
func (s *BaseIDLListener) ExitFixed_pt_const_type(ctx *Fixed_pt_const_typeContext) {}

// EnterValue_base_type is called when production value_base_type is entered.
func (s *BaseIDLListener) EnterValue_base_type(ctx *Value_base_typeContext) {}

// ExitValue_base_type is called when production value_base_type is exited.
func (s *BaseIDLListener) ExitValue_base_type(ctx *Value_base_typeContext) {}

// EnterConstr_forward_decl is called when production constr_forward_decl is entered.
func (s *BaseIDLListener) EnterConstr_forward_decl(ctx *Constr_forward_declContext) {}

// ExitConstr_forward_decl is called when production constr_forward_decl is exited.
func (s *BaseIDLListener) ExitConstr_forward_decl(ctx *Constr_forward_declContext) {}

// EnterImport_decl is called when production import_decl is entered.
func (s *BaseIDLListener) EnterImport_decl(ctx *Import_declContext) {}

// ExitImport_decl is called when production import_decl is exited.
func (s *BaseIDLListener) ExitImport_decl(ctx *Import_declContext) {}

// EnterImported_scope is called when production imported_scope is entered.
func (s *BaseIDLListener) EnterImported_scope(ctx *Imported_scopeContext) {}

// ExitImported_scope is called when production imported_scope is exited.
func (s *BaseIDLListener) ExitImported_scope(ctx *Imported_scopeContext) {}

// EnterType_id_decl is called when production type_id_decl is entered.
func (s *BaseIDLListener) EnterType_id_decl(ctx *Type_id_declContext) {}

// ExitType_id_decl is called when production type_id_decl is exited.
func (s *BaseIDLListener) ExitType_id_decl(ctx *Type_id_declContext) {}

// EnterType_prefix_decl is called when production type_prefix_decl is entered.
func (s *BaseIDLListener) EnterType_prefix_decl(ctx *Type_prefix_declContext) {}

// ExitType_prefix_decl is called when production type_prefix_decl is exited.
func (s *BaseIDLListener) ExitType_prefix_decl(ctx *Type_prefix_declContext) {}

// EnterReadonly_attr_spec is called when production readonly_attr_spec is entered.
func (s *BaseIDLListener) EnterReadonly_attr_spec(ctx *Readonly_attr_specContext) {}

// ExitReadonly_attr_spec is called when production readonly_attr_spec is exited.
func (s *BaseIDLListener) ExitReadonly_attr_spec(ctx *Readonly_attr_specContext) {}

// EnterReadonly_attr_declarator is called when production readonly_attr_declarator is entered.
func (s *BaseIDLListener) EnterReadonly_attr_declarator(ctx *Readonly_attr_declaratorContext) {}

// ExitReadonly_attr_declarator is called when production readonly_attr_declarator is exited.
func (s *BaseIDLListener) ExitReadonly_attr_declarator(ctx *Readonly_attr_declaratorContext) {}

// EnterAttr_spec is called when production attr_spec is entered.
func (s *BaseIDLListener) EnterAttr_spec(ctx *Attr_specContext) {}

// ExitAttr_spec is called when production attr_spec is exited.
func (s *BaseIDLListener) ExitAttr_spec(ctx *Attr_specContext) {}

// EnterAttr_declarator is called when production attr_declarator is entered.
func (s *BaseIDLListener) EnterAttr_declarator(ctx *Attr_declaratorContext) {}

// ExitAttr_declarator is called when production attr_declarator is exited.
func (s *BaseIDLListener) ExitAttr_declarator(ctx *Attr_declaratorContext) {}

// EnterAttr_raises_expr is called when production attr_raises_expr is entered.
func (s *BaseIDLListener) EnterAttr_raises_expr(ctx *Attr_raises_exprContext) {}

// ExitAttr_raises_expr is called when production attr_raises_expr is exited.
func (s *BaseIDLListener) ExitAttr_raises_expr(ctx *Attr_raises_exprContext) {}

// EnterGet_excep_expr is called when production get_excep_expr is entered.
func (s *BaseIDLListener) EnterGet_excep_expr(ctx *Get_excep_exprContext) {}

// ExitGet_excep_expr is called when production get_excep_expr is exited.
func (s *BaseIDLListener) ExitGet_excep_expr(ctx *Get_excep_exprContext) {}

// EnterSet_excep_expr is called when production set_excep_expr is entered.
func (s *BaseIDLListener) EnterSet_excep_expr(ctx *Set_excep_exprContext) {}

// ExitSet_excep_expr is called when production set_excep_expr is exited.
func (s *BaseIDLListener) ExitSet_excep_expr(ctx *Set_excep_exprContext) {}

// EnterException_list is called when production exception_list is entered.
func (s *BaseIDLListener) EnterException_list(ctx *Exception_listContext) {}

// ExitException_list is called when production exception_list is exited.
func (s *BaseIDLListener) ExitException_list(ctx *Exception_listContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseIDLListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseIDLListener) ExitComponent(ctx *ComponentContext) {}

// EnterComponent_forward_decl is called when production component_forward_decl is entered.
func (s *BaseIDLListener) EnterComponent_forward_decl(ctx *Component_forward_declContext) {}

// ExitComponent_forward_decl is called when production component_forward_decl is exited.
func (s *BaseIDLListener) ExitComponent_forward_decl(ctx *Component_forward_declContext) {}

// EnterComponent_decl is called when production component_decl is entered.
func (s *BaseIDLListener) EnterComponent_decl(ctx *Component_declContext) {}

// ExitComponent_decl is called when production component_decl is exited.
func (s *BaseIDLListener) ExitComponent_decl(ctx *Component_declContext) {}

// EnterComponent_header is called when production component_header is entered.
func (s *BaseIDLListener) EnterComponent_header(ctx *Component_headerContext) {}

// ExitComponent_header is called when production component_header is exited.
func (s *BaseIDLListener) ExitComponent_header(ctx *Component_headerContext) {}

// EnterSupported_interface_spec is called when production supported_interface_spec is entered.
func (s *BaseIDLListener) EnterSupported_interface_spec(ctx *Supported_interface_specContext) {}

// ExitSupported_interface_spec is called when production supported_interface_spec is exited.
func (s *BaseIDLListener) ExitSupported_interface_spec(ctx *Supported_interface_specContext) {}

// EnterComponent_inheritance_spec is called when production component_inheritance_spec is entered.
func (s *BaseIDLListener) EnterComponent_inheritance_spec(ctx *Component_inheritance_specContext) {}

// ExitComponent_inheritance_spec is called when production component_inheritance_spec is exited.
func (s *BaseIDLListener) ExitComponent_inheritance_spec(ctx *Component_inheritance_specContext) {}

// EnterComponent_body is called when production component_body is entered.
func (s *BaseIDLListener) EnterComponent_body(ctx *Component_bodyContext) {}

// ExitComponent_body is called when production component_body is exited.
func (s *BaseIDLListener) ExitComponent_body(ctx *Component_bodyContext) {}

// EnterComponent_export is called when production component_export is entered.
func (s *BaseIDLListener) EnterComponent_export(ctx *Component_exportContext) {}

// ExitComponent_export is called when production component_export is exited.
func (s *BaseIDLListener) ExitComponent_export(ctx *Component_exportContext) {}

// EnterProvides_decl is called when production provides_decl is entered.
func (s *BaseIDLListener) EnterProvides_decl(ctx *Provides_declContext) {}

// ExitProvides_decl is called when production provides_decl is exited.
func (s *BaseIDLListener) ExitProvides_decl(ctx *Provides_declContext) {}

// EnterInterface_type is called when production interface_type is entered.
func (s *BaseIDLListener) EnterInterface_type(ctx *Interface_typeContext) {}

// ExitInterface_type is called when production interface_type is exited.
func (s *BaseIDLListener) ExitInterface_type(ctx *Interface_typeContext) {}

// EnterUses_decl is called when production uses_decl is entered.
func (s *BaseIDLListener) EnterUses_decl(ctx *Uses_declContext) {}

// ExitUses_decl is called when production uses_decl is exited.
func (s *BaseIDLListener) ExitUses_decl(ctx *Uses_declContext) {}

// EnterEmits_decl is called when production emits_decl is entered.
func (s *BaseIDLListener) EnterEmits_decl(ctx *Emits_declContext) {}

// ExitEmits_decl is called when production emits_decl is exited.
func (s *BaseIDLListener) ExitEmits_decl(ctx *Emits_declContext) {}

// EnterPublishes_decl is called when production publishes_decl is entered.
func (s *BaseIDLListener) EnterPublishes_decl(ctx *Publishes_declContext) {}

// ExitPublishes_decl is called when production publishes_decl is exited.
func (s *BaseIDLListener) ExitPublishes_decl(ctx *Publishes_declContext) {}

// EnterConsumes_decl is called when production consumes_decl is entered.
func (s *BaseIDLListener) EnterConsumes_decl(ctx *Consumes_declContext) {}

// ExitConsumes_decl is called when production consumes_decl is exited.
func (s *BaseIDLListener) ExitConsumes_decl(ctx *Consumes_declContext) {}

// EnterHome_decl is called when production home_decl is entered.
func (s *BaseIDLListener) EnterHome_decl(ctx *Home_declContext) {}

// ExitHome_decl is called when production home_decl is exited.
func (s *BaseIDLListener) ExitHome_decl(ctx *Home_declContext) {}

// EnterHome_header is called when production home_header is entered.
func (s *BaseIDLListener) EnterHome_header(ctx *Home_headerContext) {}

// ExitHome_header is called when production home_header is exited.
func (s *BaseIDLListener) ExitHome_header(ctx *Home_headerContext) {}

// EnterHome_inheritance_spec is called when production home_inheritance_spec is entered.
func (s *BaseIDLListener) EnterHome_inheritance_spec(ctx *Home_inheritance_specContext) {}

// ExitHome_inheritance_spec is called when production home_inheritance_spec is exited.
func (s *BaseIDLListener) ExitHome_inheritance_spec(ctx *Home_inheritance_specContext) {}

// EnterPrimary_key_spec is called when production primary_key_spec is entered.
func (s *BaseIDLListener) EnterPrimary_key_spec(ctx *Primary_key_specContext) {}

// ExitPrimary_key_spec is called when production primary_key_spec is exited.
func (s *BaseIDLListener) ExitPrimary_key_spec(ctx *Primary_key_specContext) {}

// EnterHome_body is called when production home_body is entered.
func (s *BaseIDLListener) EnterHome_body(ctx *Home_bodyContext) {}

// ExitHome_body is called when production home_body is exited.
func (s *BaseIDLListener) ExitHome_body(ctx *Home_bodyContext) {}

// EnterHome_export is called when production home_export is entered.
func (s *BaseIDLListener) EnterHome_export(ctx *Home_exportContext) {}

// ExitHome_export is called when production home_export is exited.
func (s *BaseIDLListener) ExitHome_export(ctx *Home_exportContext) {}

// EnterFactory_decl is called when production factory_decl is entered.
func (s *BaseIDLListener) EnterFactory_decl(ctx *Factory_declContext) {}

// ExitFactory_decl is called when production factory_decl is exited.
func (s *BaseIDLListener) ExitFactory_decl(ctx *Factory_declContext) {}

// EnterFinder_decl is called when production finder_decl is entered.
func (s *BaseIDLListener) EnterFinder_decl(ctx *Finder_declContext) {}

// ExitFinder_decl is called when production finder_decl is exited.
func (s *BaseIDLListener) ExitFinder_decl(ctx *Finder_declContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseIDLListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseIDLListener) ExitEvent(ctx *EventContext) {}

// EnterEvent_forward_decl is called when production event_forward_decl is entered.
func (s *BaseIDLListener) EnterEvent_forward_decl(ctx *Event_forward_declContext) {}

// ExitEvent_forward_decl is called when production event_forward_decl is exited.
func (s *BaseIDLListener) ExitEvent_forward_decl(ctx *Event_forward_declContext) {}

// EnterEvent_abs_decl is called when production event_abs_decl is entered.
func (s *BaseIDLListener) EnterEvent_abs_decl(ctx *Event_abs_declContext) {}

// ExitEvent_abs_decl is called when production event_abs_decl is exited.
func (s *BaseIDLListener) ExitEvent_abs_decl(ctx *Event_abs_declContext) {}

// EnterEvent_decl is called when production event_decl is entered.
func (s *BaseIDLListener) EnterEvent_decl(ctx *Event_declContext) {}

// ExitEvent_decl is called when production event_decl is exited.
func (s *BaseIDLListener) ExitEvent_decl(ctx *Event_declContext) {}

// EnterEvent_header is called when production event_header is entered.
func (s *BaseIDLListener) EnterEvent_header(ctx *Event_headerContext) {}

// ExitEvent_header is called when production event_header is exited.
func (s *BaseIDLListener) ExitEvent_header(ctx *Event_headerContext) {}
