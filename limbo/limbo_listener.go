// Code generated from limbo.g4 by ANTLR 4.9.3. DO NOT EDIT.

package limbo // limbo
import "github.com/antlr/antlr4/runtime/Go/antlr"

// limboListener is a complete listener for a parse tree produced by limboParser.
type limboListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTop_declaration_sequence is called when entering the top_declaration_sequence production.
	EnterTop_declaration_sequence(c *Top_declaration_sequenceContext)

	// EnterTop_declaration is called when entering the top_declaration production.
	EnterTop_declaration(c *Top_declarationContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterReal_ is called when entering the real_ production.
	EnterReal_(c *Real_Context)

	// EnterByte_ is called when entering the byte_ production.
	EnterByte_(c *Byte_Context)

	// EnterInt_ is called when entering the int_ production.
	EnterInt_(c *Int_Context)

	// EnterBig is called when entering the big production.
	EnterBig(c *BigContext)

	// EnterTuple_type is called when entering the tuple_type production.
	EnterTuple_type(c *Tuple_typeContext)

	// EnterData_type_list is called when entering the data_type_list production.
	EnterData_type_list(c *Data_type_listContext)

	// EnterAdt_type is called when entering the adt_type production.
	EnterAdt_type(c *Adt_typeContext)

	// EnterModule_type is called when entering the module_type production.
	EnterModule_type(c *Module_typeContext)

	// EnterModule_qualified_type is called when entering the module_qualified_type production.
	EnterModule_qualified_type(c *Module_qualified_typeContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterFunction_arg_ret is called when entering the function_arg_ret production.
	EnterFunction_arg_ret(c *Function_arg_retContext)

	// EnterFormal_arg_list is called when entering the formal_arg_list production.
	EnterFormal_arg_list(c *Formal_arg_listContext)

	// EnterFormal_arg is called when entering the formal_arg production.
	EnterFormal_arg(c *Formal_argContext)

	// EnterNil_or_D_list is called when entering the nil_or_D_list production.
	EnterNil_or_D_list(c *Nil_or_D_listContext)

	// EnterNil_or_D is called when entering the nil_or_D production.
	EnterNil_or_D(c *Nil_or_DContext)

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterMod_member_list is called when entering the mod_member_list production.
	EnterMod_member_list(c *Mod_member_listContext)

	// EnterMod_member is called when entering the mod_member production.
	EnterMod_member(c *Mod_memberContext)

	// EnterAdt_declaration is called when entering the adt_declaration production.
	EnterAdt_declaration(c *Adt_declarationContext)

	// EnterAdt_member_list is called when entering the adt_member_list production.
	EnterAdt_member_list(c *Adt_member_listContext)

	// EnterAdt_member is called when entering the adt_member production.
	EnterAdt_member(c *Adt_memberContext)

	// EnterPick_member_list is called when entering the pick_member_list production.
	EnterPick_member_list(c *Pick_member_listContext)

	// EnterPick_tag_list is called when entering the pick_tag_list production.
	EnterPick_tag_list(c *Pick_tag_listContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterFunction_name_part is called when entering the function_name_part production.
	EnterFunction_name_part(c *Function_name_partContext)

	// EnterStatements_ is called when entering the statements_ production.
	EnterStatements_(c *Statements_Context)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterQual_statement_sequence is called when entering the qual_statement_sequence production.
	EnterQual_statement_sequence(c *Qual_statement_sequenceContext)

	// EnterQual_list is called when entering the qual_list production.
	EnterQual_list(c *Qual_listContext)

	// EnterQualifier is called when entering the qualifier production.
	EnterQualifier(c *QualifierContext)

	// EnterPqual_statement_sequence is called when entering the pqual_statement_sequence production.
	EnterPqual_statement_sequence(c *Pqual_statement_sequenceContext)

	// EnterPqual_list is called when entering the pqual_list production.
	EnterPqual_list(c *Pqual_listContext)

	// EnterPqualifier is called when entering the pqualifier production.
	EnterPqualifier(c *PqualifierContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBinary_expression is called when entering the binary_expression production.
	EnterBinary_expression(c *Binary_expressionContext)

	// EnterLvalue_expression is called when entering the lvalue_expression production.
	EnterLvalue_expression(c *Lvalue_expressionContext)

	// EnterLvalue_expression_list is called when entering the lvalue_expression_list production.
	EnterLvalue_expression_list(c *Lvalue_expression_listContext)

	// EnterInit_list is called when entering the init_list production.
	EnterInit_list(c *Init_listContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterReal_constant is called when entering the real_constant production.
	EnterReal_constant(c *Real_constantContext)

	// EnterString_constant is called when entering the string_constant production.
	EnterString_constant(c *String_constantContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterSend_expression is called when entering the send_expression production.
	EnterSend_expression(c *Send_expressionContext)

	// EnterDeclare_expression is called when entering the declare_expression production.
	EnterDeclare_expression(c *Declare_expressionContext)

	// EnterLoad_expression is called when entering the load_expression production.
	EnterLoad_expression(c *Load_expressionContext)

	// EnterMonadic_expression is called when entering the monadic_expression production.
	EnterMonadic_expression(c *Monadic_expressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTop_declaration_sequence is called when exiting the top_declaration_sequence production.
	ExitTop_declaration_sequence(c *Top_declaration_sequenceContext)

	// ExitTop_declaration is called when exiting the top_declaration production.
	ExitTop_declaration(c *Top_declarationContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitReal_ is called when exiting the real_ production.
	ExitReal_(c *Real_Context)

	// ExitByte_ is called when exiting the byte_ production.
	ExitByte_(c *Byte_Context)

	// ExitInt_ is called when exiting the int_ production.
	ExitInt_(c *Int_Context)

	// ExitBig is called when exiting the big production.
	ExitBig(c *BigContext)

	// ExitTuple_type is called when exiting the tuple_type production.
	ExitTuple_type(c *Tuple_typeContext)

	// ExitData_type_list is called when exiting the data_type_list production.
	ExitData_type_list(c *Data_type_listContext)

	// ExitAdt_type is called when exiting the adt_type production.
	ExitAdt_type(c *Adt_typeContext)

	// ExitModule_type is called when exiting the module_type production.
	ExitModule_type(c *Module_typeContext)

	// ExitModule_qualified_type is called when exiting the module_qualified_type production.
	ExitModule_qualified_type(c *Module_qualified_typeContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitFunction_arg_ret is called when exiting the function_arg_ret production.
	ExitFunction_arg_ret(c *Function_arg_retContext)

	// ExitFormal_arg_list is called when exiting the formal_arg_list production.
	ExitFormal_arg_list(c *Formal_arg_listContext)

	// ExitFormal_arg is called when exiting the formal_arg production.
	ExitFormal_arg(c *Formal_argContext)

	// ExitNil_or_D_list is called when exiting the nil_or_D_list production.
	ExitNil_or_D_list(c *Nil_or_D_listContext)

	// ExitNil_or_D is called when exiting the nil_or_D production.
	ExitNil_or_D(c *Nil_or_DContext)

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitMod_member_list is called when exiting the mod_member_list production.
	ExitMod_member_list(c *Mod_member_listContext)

	// ExitMod_member is called when exiting the mod_member production.
	ExitMod_member(c *Mod_memberContext)

	// ExitAdt_declaration is called when exiting the adt_declaration production.
	ExitAdt_declaration(c *Adt_declarationContext)

	// ExitAdt_member_list is called when exiting the adt_member_list production.
	ExitAdt_member_list(c *Adt_member_listContext)

	// ExitAdt_member is called when exiting the adt_member production.
	ExitAdt_member(c *Adt_memberContext)

	// ExitPick_member_list is called when exiting the pick_member_list production.
	ExitPick_member_list(c *Pick_member_listContext)

	// ExitPick_tag_list is called when exiting the pick_tag_list production.
	ExitPick_tag_list(c *Pick_tag_listContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitFunction_name_part is called when exiting the function_name_part production.
	ExitFunction_name_part(c *Function_name_partContext)

	// ExitStatements_ is called when exiting the statements_ production.
	ExitStatements_(c *Statements_Context)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitQual_statement_sequence is called when exiting the qual_statement_sequence production.
	ExitQual_statement_sequence(c *Qual_statement_sequenceContext)

	// ExitQual_list is called when exiting the qual_list production.
	ExitQual_list(c *Qual_listContext)

	// ExitQualifier is called when exiting the qualifier production.
	ExitQualifier(c *QualifierContext)

	// ExitPqual_statement_sequence is called when exiting the pqual_statement_sequence production.
	ExitPqual_statement_sequence(c *Pqual_statement_sequenceContext)

	// ExitPqual_list is called when exiting the pqual_list production.
	ExitPqual_list(c *Pqual_listContext)

	// ExitPqualifier is called when exiting the pqualifier production.
	ExitPqualifier(c *PqualifierContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBinary_expression is called when exiting the binary_expression production.
	ExitBinary_expression(c *Binary_expressionContext)

	// ExitLvalue_expression is called when exiting the lvalue_expression production.
	ExitLvalue_expression(c *Lvalue_expressionContext)

	// ExitLvalue_expression_list is called when exiting the lvalue_expression_list production.
	ExitLvalue_expression_list(c *Lvalue_expression_listContext)

	// ExitInit_list is called when exiting the init_list production.
	ExitInit_list(c *Init_listContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitReal_constant is called when exiting the real_constant production.
	ExitReal_constant(c *Real_constantContext)

	// ExitString_constant is called when exiting the string_constant production.
	ExitString_constant(c *String_constantContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitSend_expression is called when exiting the send_expression production.
	ExitSend_expression(c *Send_expressionContext)

	// ExitDeclare_expression is called when exiting the declare_expression production.
	ExitDeclare_expression(c *Declare_expressionContext)

	// ExitLoad_expression is called when exiting the load_expression production.
	ExitLoad_expression(c *Load_expressionContext)

	// ExitMonadic_expression is called when exiting the monadic_expression production.
	ExitMonadic_expression(c *Monadic_expressionContext)
}
