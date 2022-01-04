// Code generated from limbo.g4 by ANTLR 4.9.3. DO NOT EDIT.

package limbo // limbo
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselimboListener is a complete listener for a parse tree produced by limboParser.
type BaselimboListener struct{}

var _ limboListener = &BaselimboListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselimboListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselimboListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselimboListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselimboListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselimboListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaselimboListener) ExitProgram(ctx *ProgramContext) {}

// EnterTop_declaration_sequence is called when production top_declaration_sequence is entered.
func (s *BaselimboListener) EnterTop_declaration_sequence(ctx *Top_declaration_sequenceContext) {}

// ExitTop_declaration_sequence is called when production top_declaration_sequence is exited.
func (s *BaselimboListener) ExitTop_declaration_sequence(ctx *Top_declaration_sequenceContext) {}

// EnterTop_declaration is called when production top_declaration is entered.
func (s *BaselimboListener) EnterTop_declaration(ctx *Top_declarationContext) {}

// ExitTop_declaration is called when production top_declaration is exited.
func (s *BaselimboListener) ExitTop_declaration(ctx *Top_declarationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaselimboListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaselimboListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *BaselimboListener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *BaselimboListener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaselimboListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaselimboListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaselimboListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaselimboListener) ExitType_(ctx *Type_Context) {}

// EnterData_type is called when production data_type is entered.
func (s *BaselimboListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaselimboListener) ExitData_type(ctx *Data_typeContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaselimboListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaselimboListener) ExitString_(ctx *String_Context) {}

// EnterReal_ is called when production real_ is entered.
func (s *BaselimboListener) EnterReal_(ctx *Real_Context) {}

// ExitReal_ is called when production real_ is exited.
func (s *BaselimboListener) ExitReal_(ctx *Real_Context) {}

// EnterByte_ is called when production byte_ is entered.
func (s *BaselimboListener) EnterByte_(ctx *Byte_Context) {}

// ExitByte_ is called when production byte_ is exited.
func (s *BaselimboListener) ExitByte_(ctx *Byte_Context) {}

// EnterInt_ is called when production int_ is entered.
func (s *BaselimboListener) EnterInt_(ctx *Int_Context) {}

// ExitInt_ is called when production int_ is exited.
func (s *BaselimboListener) ExitInt_(ctx *Int_Context) {}

// EnterBig is called when production big is entered.
func (s *BaselimboListener) EnterBig(ctx *BigContext) {}

// ExitBig is called when production big is exited.
func (s *BaselimboListener) ExitBig(ctx *BigContext) {}

// EnterTuple_type is called when production tuple_type is entered.
func (s *BaselimboListener) EnterTuple_type(ctx *Tuple_typeContext) {}

// ExitTuple_type is called when production tuple_type is exited.
func (s *BaselimboListener) ExitTuple_type(ctx *Tuple_typeContext) {}

// EnterData_type_list is called when production data_type_list is entered.
func (s *BaselimboListener) EnterData_type_list(ctx *Data_type_listContext) {}

// ExitData_type_list is called when production data_type_list is exited.
func (s *BaselimboListener) ExitData_type_list(ctx *Data_type_listContext) {}

// EnterAdt_type is called when production adt_type is entered.
func (s *BaselimboListener) EnterAdt_type(ctx *Adt_typeContext) {}

// ExitAdt_type is called when production adt_type is exited.
func (s *BaselimboListener) ExitAdt_type(ctx *Adt_typeContext) {}

// EnterModule_type is called when production module_type is entered.
func (s *BaselimboListener) EnterModule_type(ctx *Module_typeContext) {}

// ExitModule_type is called when production module_type is exited.
func (s *BaselimboListener) ExitModule_type(ctx *Module_typeContext) {}

// EnterModule_qualified_type is called when production module_qualified_type is entered.
func (s *BaselimboListener) EnterModule_qualified_type(ctx *Module_qualified_typeContext) {}

// ExitModule_qualified_type is called when production module_qualified_type is exited.
func (s *BaselimboListener) ExitModule_qualified_type(ctx *Module_qualified_typeContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaselimboListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaselimboListener) ExitType_name(ctx *Type_nameContext) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BaselimboListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BaselimboListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterFunction_arg_ret is called when production function_arg_ret is entered.
func (s *BaselimboListener) EnterFunction_arg_ret(ctx *Function_arg_retContext) {}

// ExitFunction_arg_ret is called when production function_arg_ret is exited.
func (s *BaselimboListener) ExitFunction_arg_ret(ctx *Function_arg_retContext) {}

// EnterFormal_arg_list is called when production formal_arg_list is entered.
func (s *BaselimboListener) EnterFormal_arg_list(ctx *Formal_arg_listContext) {}

// ExitFormal_arg_list is called when production formal_arg_list is exited.
func (s *BaselimboListener) ExitFormal_arg_list(ctx *Formal_arg_listContext) {}

// EnterFormal_arg is called when production formal_arg is entered.
func (s *BaselimboListener) EnterFormal_arg(ctx *Formal_argContext) {}

// ExitFormal_arg is called when production formal_arg is exited.
func (s *BaselimboListener) ExitFormal_arg(ctx *Formal_argContext) {}

// EnterNil_or_D_list is called when production nil_or_D_list is entered.
func (s *BaselimboListener) EnterNil_or_D_list(ctx *Nil_or_D_listContext) {}

// ExitNil_or_D_list is called when production nil_or_D_list is exited.
func (s *BaselimboListener) ExitNil_or_D_list(ctx *Nil_or_D_listContext) {}

// EnterNil_or_D is called when production nil_or_D is entered.
func (s *BaselimboListener) EnterNil_or_D(ctx *Nil_or_DContext) {}

// ExitNil_or_D is called when production nil_or_D is exited.
func (s *BaselimboListener) ExitNil_or_D(ctx *Nil_or_DContext) {}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *BaselimboListener) EnterModule_declaration(ctx *Module_declarationContext) {}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *BaselimboListener) ExitModule_declaration(ctx *Module_declarationContext) {}

// EnterMod_member_list is called when production mod_member_list is entered.
func (s *BaselimboListener) EnterMod_member_list(ctx *Mod_member_listContext) {}

// ExitMod_member_list is called when production mod_member_list is exited.
func (s *BaselimboListener) ExitMod_member_list(ctx *Mod_member_listContext) {}

// EnterMod_member is called when production mod_member is entered.
func (s *BaselimboListener) EnterMod_member(ctx *Mod_memberContext) {}

// ExitMod_member is called when production mod_member is exited.
func (s *BaselimboListener) ExitMod_member(ctx *Mod_memberContext) {}

// EnterAdt_declaration is called when production adt_declaration is entered.
func (s *BaselimboListener) EnterAdt_declaration(ctx *Adt_declarationContext) {}

// ExitAdt_declaration is called when production adt_declaration is exited.
func (s *BaselimboListener) ExitAdt_declaration(ctx *Adt_declarationContext) {}

// EnterAdt_member_list is called when production adt_member_list is entered.
func (s *BaselimboListener) EnterAdt_member_list(ctx *Adt_member_listContext) {}

// ExitAdt_member_list is called when production adt_member_list is exited.
func (s *BaselimboListener) ExitAdt_member_list(ctx *Adt_member_listContext) {}

// EnterAdt_member is called when production adt_member is entered.
func (s *BaselimboListener) EnterAdt_member(ctx *Adt_memberContext) {}

// ExitAdt_member is called when production adt_member is exited.
func (s *BaselimboListener) ExitAdt_member(ctx *Adt_memberContext) {}

// EnterPick_member_list is called when production pick_member_list is entered.
func (s *BaselimboListener) EnterPick_member_list(ctx *Pick_member_listContext) {}

// ExitPick_member_list is called when production pick_member_list is exited.
func (s *BaselimboListener) ExitPick_member_list(ctx *Pick_member_listContext) {}

// EnterPick_tag_list is called when production pick_tag_list is entered.
func (s *BaselimboListener) EnterPick_tag_list(ctx *Pick_tag_listContext) {}

// ExitPick_tag_list is called when production pick_tag_list is exited.
func (s *BaselimboListener) ExitPick_tag_list(ctx *Pick_tag_listContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaselimboListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaselimboListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterFunction_name_part is called when production function_name_part is entered.
func (s *BaselimboListener) EnterFunction_name_part(ctx *Function_name_partContext) {}

// ExitFunction_name_part is called when production function_name_part is exited.
func (s *BaselimboListener) ExitFunction_name_part(ctx *Function_name_partContext) {}

// EnterStatements_ is called when production statements_ is entered.
func (s *BaselimboListener) EnterStatements_(ctx *Statements_Context) {}

// ExitStatements_ is called when production statements_ is exited.
func (s *BaselimboListener) ExitStatements_(ctx *Statements_Context) {}

// EnterStatement is called when production statement is entered.
func (s *BaselimboListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaselimboListener) ExitStatement(ctx *StatementContext) {}

// EnterLabel is called when production label is entered.
func (s *BaselimboListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaselimboListener) ExitLabel(ctx *LabelContext) {}

// EnterQual_statement_sequence is called when production qual_statement_sequence is entered.
func (s *BaselimboListener) EnterQual_statement_sequence(ctx *Qual_statement_sequenceContext) {}

// ExitQual_statement_sequence is called when production qual_statement_sequence is exited.
func (s *BaselimboListener) ExitQual_statement_sequence(ctx *Qual_statement_sequenceContext) {}

// EnterQual_list is called when production qual_list is entered.
func (s *BaselimboListener) EnterQual_list(ctx *Qual_listContext) {}

// ExitQual_list is called when production qual_list is exited.
func (s *BaselimboListener) ExitQual_list(ctx *Qual_listContext) {}

// EnterQualifier is called when production qualifier is entered.
func (s *BaselimboListener) EnterQualifier(ctx *QualifierContext) {}

// ExitQualifier is called when production qualifier is exited.
func (s *BaselimboListener) ExitQualifier(ctx *QualifierContext) {}

// EnterPqual_statement_sequence is called when production pqual_statement_sequence is entered.
func (s *BaselimboListener) EnterPqual_statement_sequence(ctx *Pqual_statement_sequenceContext) {}

// ExitPqual_statement_sequence is called when production pqual_statement_sequence is exited.
func (s *BaselimboListener) ExitPqual_statement_sequence(ctx *Pqual_statement_sequenceContext) {}

// EnterPqual_list is called when production pqual_list is entered.
func (s *BaselimboListener) EnterPqual_list(ctx *Pqual_listContext) {}

// ExitPqual_list is called when production pqual_list is exited.
func (s *BaselimboListener) ExitPqual_list(ctx *Pqual_listContext) {}

// EnterPqualifier is called when production pqualifier is entered.
func (s *BaselimboListener) EnterPqualifier(ctx *PqualifierContext) {}

// ExitPqualifier is called when production pqualifier is exited.
func (s *BaselimboListener) ExitPqualifier(ctx *PqualifierContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaselimboListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaselimboListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBinary_expression is called when production binary_expression is entered.
func (s *BaselimboListener) EnterBinary_expression(ctx *Binary_expressionContext) {}

// ExitBinary_expression is called when production binary_expression is exited.
func (s *BaselimboListener) ExitBinary_expression(ctx *Binary_expressionContext) {}

// EnterLvalue_expression is called when production lvalue_expression is entered.
func (s *BaselimboListener) EnterLvalue_expression(ctx *Lvalue_expressionContext) {}

// ExitLvalue_expression is called when production lvalue_expression is exited.
func (s *BaselimboListener) ExitLvalue_expression(ctx *Lvalue_expressionContext) {}

// EnterLvalue_expression_list is called when production lvalue_expression_list is entered.
func (s *BaselimboListener) EnterLvalue_expression_list(ctx *Lvalue_expression_listContext) {}

// ExitLvalue_expression_list is called when production lvalue_expression_list is exited.
func (s *BaselimboListener) ExitLvalue_expression_list(ctx *Lvalue_expression_listContext) {}

// EnterInit_list is called when production init_list is entered.
func (s *BaselimboListener) EnterInit_list(ctx *Init_listContext) {}

// ExitInit_list is called when production init_list is exited.
func (s *BaselimboListener) ExitInit_list(ctx *Init_listContext) {}

// EnterTerm is called when production term is entered.
func (s *BaselimboListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaselimboListener) ExitTerm(ctx *TermContext) {}

// EnterReal_constant is called when production real_constant is entered.
func (s *BaselimboListener) EnterReal_constant(ctx *Real_constantContext) {}

// ExitReal_constant is called when production real_constant is exited.
func (s *BaselimboListener) ExitReal_constant(ctx *Real_constantContext) {}

// EnterString_constant is called when production string_constant is entered.
func (s *BaselimboListener) EnterString_constant(ctx *String_constantContext) {}

// ExitString_constant is called when production string_constant is exited.
func (s *BaselimboListener) ExitString_constant(ctx *String_constantContext) {}

// EnterElement is called when production element is entered.
func (s *BaselimboListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaselimboListener) ExitElement(ctx *ElementContext) {}

// EnterSend_expression is called when production send_expression is entered.
func (s *BaselimboListener) EnterSend_expression(ctx *Send_expressionContext) {}

// ExitSend_expression is called when production send_expression is exited.
func (s *BaselimboListener) ExitSend_expression(ctx *Send_expressionContext) {}

// EnterDeclare_expression is called when production declare_expression is entered.
func (s *BaselimboListener) EnterDeclare_expression(ctx *Declare_expressionContext) {}

// ExitDeclare_expression is called when production declare_expression is exited.
func (s *BaselimboListener) ExitDeclare_expression(ctx *Declare_expressionContext) {}

// EnterLoad_expression is called when production load_expression is entered.
func (s *BaselimboListener) EnterLoad_expression(ctx *Load_expressionContext) {}

// ExitLoad_expression is called when production load_expression is exited.
func (s *BaselimboListener) ExitLoad_expression(ctx *Load_expressionContext) {}

// EnterMonadic_expression is called when production monadic_expression is entered.
func (s *BaselimboListener) EnterMonadic_expression(ctx *Monadic_expressionContext) {}

// ExitMonadic_expression is called when production monadic_expression is exited.
func (s *BaselimboListener) ExitMonadic_expression(ctx *Monadic_expressionContext) {}
