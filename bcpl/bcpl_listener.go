// Code generated from bcpl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcpl // bcpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bcplListener is a complete listener for a parse tree produced by bcplParser.
type bcplListener interface {
	antlr.ParseTreeListener

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterOctal_digit is called when entering the octal_digit production.
	EnterOctal_digit(c *Octal_digitContext)

	// EnterHex_digit is called when entering the hex_digit production.
	EnterHex_digit(c *Hex_digitContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterString_constant is called when entering the string_constant production.
	EnterString_constant(c *String_constantContext)

	// EnterCharacter_constant is called when entering the character_constant production.
	EnterCharacter_constant(c *Character_constantContext)

	// EnterOctal_number is called when entering the octal_number production.
	EnterOctal_number(c *Octal_numberContext)

	// EnterHex_number is called when entering the hex_number production.
	EnterHex_number(c *Hex_numberContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterOne_character is called when entering the one_character production.
	EnterOne_character(c *One_characterContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterAddress_op is called when entering the address_op production.
	EnterAddress_op(c *Address_opContext)

	// EnterMult_op is called when entering the mult_op production.
	EnterMult_op(c *Mult_opContext)

	// EnterAdd_op is called when entering the add_op production.
	EnterAdd_op(c *Add_opContext)

	// EnterRel_op is called when entering the rel_op production.
	EnterRel_op(c *Rel_opContext)

	// EnterShift_op is called when entering the shift_op production.
	EnterShift_op(c *Shift_opContext)

	// EnterAnd_op is called when entering the and_op production.
	EnterAnd_op(c *And_opContext)

	// EnterOr_op is called when entering the or_op production.
	EnterOr_op(c *Or_opContext)

	// EnterEqv_op is called when entering the eqv_op production.
	EnterEqv_op(c *Eqv_opContext)

	// EnterNot_op is called when entering the not_op production.
	EnterNot_op(c *Not_opContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterPrimary_E is called when entering the primary_E production.
	EnterPrimary_E(c *Primary_EContext)

	// EnterVector_E is called when entering the vector_E production.
	EnterVector_E(c *Vector_EContext)

	// EnterAddress_E is called when entering the address_E production.
	EnterAddress_E(c *Address_EContext)

	// EnterMult_E is called when entering the mult_E production.
	EnterMult_E(c *Mult_EContext)

	// EnterAdd_E is called when entering the add_E production.
	EnterAdd_E(c *Add_EContext)

	// EnterRel_E is called when entering the rel_E production.
	EnterRel_E(c *Rel_EContext)

	// EnterShift_E is called when entering the shift_E production.
	EnterShift_E(c *Shift_EContext)

	// EnterNot_E is called when entering the not_E production.
	EnterNot_E(c *Not_EContext)

	// EnterAnd_E is called when entering the and_E production.
	EnterAnd_E(c *And_EContext)

	// EnterOr_E is called when entering the or_E production.
	EnterOr_E(c *Or_EContext)

	// EnterEqv_E is called when entering the eqv_E production.
	EnterEqv_E(c *Eqv_EContext)

	// EnterConditional_E is called when entering the conditional_E production.
	EnterConditional_E(c *Conditional_EContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterC_element is called when entering the c_element production.
	EnterC_element(c *C_elementContext)

	// EnterC_mult_E is called when entering the c_mult_E production.
	EnterC_mult_E(c *C_mult_EContext)

	// EnterC_add_E is called when entering the c_add_E production.
	EnterC_add_E(c *C_add_EContext)

	// EnterC_shift_E is called when entering the c_shift_E production.
	EnterC_shift_E(c *C_shift_EContext)

	// EnterC_and_E is called when entering the c_and_E production.
	EnterC_and_E(c *C_and_EContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterManifest_item is called when entering the manifest_item production.
	EnterManifest_item(c *Manifest_itemContext)

	// EnterManifest_list is called when entering the manifest_list production.
	EnterManifest_list(c *Manifest_listContext)

	// EnterManifest_declaration is called when entering the manifest_declaration production.
	EnterManifest_declaration(c *Manifest_declarationContext)

	// EnterStatic_declaration is called when entering the static_declaration production.
	EnterStatic_declaration(c *Static_declarationContext)

	// EnterGlobal_item is called when entering the global_item production.
	EnterGlobal_item(c *Global_itemContext)

	// EnterGlobal_list is called when entering the global_list production.
	EnterGlobal_list(c *Global_listContext)

	// EnterGlobal_declaration is called when entering the global_declaration production.
	EnterGlobal_declaration(c *Global_declarationContext)

	// EnterSimple_definition is called when entering the simple_definition production.
	EnterSimple_definition(c *Simple_definitionContext)

	// EnterVector_definition is called when entering the vector_definition production.
	EnterVector_definition(c *Vector_definitionContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterRoutine_definition is called when entering the routine_definition production.
	EnterRoutine_definition(c *Routine_definitionContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterSimultaneous_declaration is called when entering the simultaneous_declaration production.
	EnterSimultaneous_declaration(c *Simultaneous_declarationContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterLhse is called when entering the lhse production.
	EnterLhse(c *LhseContext)

	// EnterLeft_hand_side_list is called when entering the left_hand_side_list production.
	EnterLeft_hand_side_list(c *Left_hand_side_listContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSimple_command is called when entering the simple_command production.
	EnterSimple_command(c *Simple_commandContext)

	// EnterGoto_command is called when entering the goto_command production.
	EnterGoto_command(c *Goto_commandContext)

	// EnterRoutine_command is called when entering the routine_command production.
	EnterRoutine_command(c *Routine_commandContext)

	// EnterResultis_command is called when entering the resultis_command production.
	EnterResultis_command(c *Resultis_commandContext)

	// EnterSwitchon_command is called when entering the switchon_command production.
	EnterSwitchon_command(c *Switchon_commandContext)

	// EnterRepeatable_command is called when entering the repeatable_command production.
	EnterRepeatable_command(c *Repeatable_commandContext)

	// EnterRepeated_command is called when entering the repeated_command production.
	EnterRepeated_command(c *Repeated_commandContext)

	// EnterUntil_command is called when entering the until_command production.
	EnterUntil_command(c *Until_commandContext)

	// EnterWhile_command is called when entering the while_command production.
	EnterWhile_command(c *While_commandContext)

	// EnterFor_command is called when entering the for_command production.
	EnterFor_command(c *For_commandContext)

	// EnterRepetitive_command is called when entering the repetitive_command production.
	EnterRepetitive_command(c *Repetitive_commandContext)

	// EnterTest_command is called when entering the test_command production.
	EnterTest_command(c *Test_commandContext)

	// EnterIf_command is called when entering the if_command production.
	EnterIf_command(c *If_commandContext)

	// EnterUnless_command is called when entering the unless_command production.
	EnterUnless_command(c *Unless_commandContext)

	// EnterUnlabelled_command is called when entering the unlabelled_command production.
	EnterUnlabelled_command(c *Unlabelled_commandContext)

	// EnterLabel_prefix is called when entering the label_prefix production.
	EnterLabel_prefix(c *Label_prefixContext)

	// EnterCase_prefix is called when entering the case_prefix production.
	EnterCase_prefix(c *Case_prefixContext)

	// EnterDefault_prefix is called when entering the default_prefix production.
	EnterDefault_prefix(c *Default_prefixContext)

	// EnterPrefix_ is called when entering the prefix_ production.
	EnterPrefix_(c *Prefix_Context)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterCommand_list is called when entering the command_list production.
	EnterCommand_list(c *Command_listContext)

	// EnterDeclaration_part is called when entering the declaration_part production.
	EnterDeclaration_part(c *Declaration_partContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterCompound_command is called when entering the compound_command production.
	EnterCompound_command(c *Compound_commandContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitOctal_digit is called when exiting the octal_digit production.
	ExitOctal_digit(c *Octal_digitContext)

	// ExitHex_digit is called when exiting the hex_digit production.
	ExitHex_digit(c *Hex_digitContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitString_constant is called when exiting the string_constant production.
	ExitString_constant(c *String_constantContext)

	// ExitCharacter_constant is called when exiting the character_constant production.
	ExitCharacter_constant(c *Character_constantContext)

	// ExitOctal_number is called when exiting the octal_number production.
	ExitOctal_number(c *Octal_numberContext)

	// ExitHex_number is called when exiting the hex_number production.
	ExitHex_number(c *Hex_numberContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitOne_character is called when exiting the one_character production.
	ExitOne_character(c *One_characterContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitAddress_op is called when exiting the address_op production.
	ExitAddress_op(c *Address_opContext)

	// ExitMult_op is called when exiting the mult_op production.
	ExitMult_op(c *Mult_opContext)

	// ExitAdd_op is called when exiting the add_op production.
	ExitAdd_op(c *Add_opContext)

	// ExitRel_op is called when exiting the rel_op production.
	ExitRel_op(c *Rel_opContext)

	// ExitShift_op is called when exiting the shift_op production.
	ExitShift_op(c *Shift_opContext)

	// ExitAnd_op is called when exiting the and_op production.
	ExitAnd_op(c *And_opContext)

	// ExitOr_op is called when exiting the or_op production.
	ExitOr_op(c *Or_opContext)

	// ExitEqv_op is called when exiting the eqv_op production.
	ExitEqv_op(c *Eqv_opContext)

	// ExitNot_op is called when exiting the not_op production.
	ExitNot_op(c *Not_opContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitPrimary_E is called when exiting the primary_E production.
	ExitPrimary_E(c *Primary_EContext)

	// ExitVector_E is called when exiting the vector_E production.
	ExitVector_E(c *Vector_EContext)

	// ExitAddress_E is called when exiting the address_E production.
	ExitAddress_E(c *Address_EContext)

	// ExitMult_E is called when exiting the mult_E production.
	ExitMult_E(c *Mult_EContext)

	// ExitAdd_E is called when exiting the add_E production.
	ExitAdd_E(c *Add_EContext)

	// ExitRel_E is called when exiting the rel_E production.
	ExitRel_E(c *Rel_EContext)

	// ExitShift_E is called when exiting the shift_E production.
	ExitShift_E(c *Shift_EContext)

	// ExitNot_E is called when exiting the not_E production.
	ExitNot_E(c *Not_EContext)

	// ExitAnd_E is called when exiting the and_E production.
	ExitAnd_E(c *And_EContext)

	// ExitOr_E is called when exiting the or_E production.
	ExitOr_E(c *Or_EContext)

	// ExitEqv_E is called when exiting the eqv_E production.
	ExitEqv_E(c *Eqv_EContext)

	// ExitConditional_E is called when exiting the conditional_E production.
	ExitConditional_E(c *Conditional_EContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitC_element is called when exiting the c_element production.
	ExitC_element(c *C_elementContext)

	// ExitC_mult_E is called when exiting the c_mult_E production.
	ExitC_mult_E(c *C_mult_EContext)

	// ExitC_add_E is called when exiting the c_add_E production.
	ExitC_add_E(c *C_add_EContext)

	// ExitC_shift_E is called when exiting the c_shift_E production.
	ExitC_shift_E(c *C_shift_EContext)

	// ExitC_and_E is called when exiting the c_and_E production.
	ExitC_and_E(c *C_and_EContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitManifest_item is called when exiting the manifest_item production.
	ExitManifest_item(c *Manifest_itemContext)

	// ExitManifest_list is called when exiting the manifest_list production.
	ExitManifest_list(c *Manifest_listContext)

	// ExitManifest_declaration is called when exiting the manifest_declaration production.
	ExitManifest_declaration(c *Manifest_declarationContext)

	// ExitStatic_declaration is called when exiting the static_declaration production.
	ExitStatic_declaration(c *Static_declarationContext)

	// ExitGlobal_item is called when exiting the global_item production.
	ExitGlobal_item(c *Global_itemContext)

	// ExitGlobal_list is called when exiting the global_list production.
	ExitGlobal_list(c *Global_listContext)

	// ExitGlobal_declaration is called when exiting the global_declaration production.
	ExitGlobal_declaration(c *Global_declarationContext)

	// ExitSimple_definition is called when exiting the simple_definition production.
	ExitSimple_definition(c *Simple_definitionContext)

	// ExitVector_definition is called when exiting the vector_definition production.
	ExitVector_definition(c *Vector_definitionContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitRoutine_definition is called when exiting the routine_definition production.
	ExitRoutine_definition(c *Routine_definitionContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitSimultaneous_declaration is called when exiting the simultaneous_declaration production.
	ExitSimultaneous_declaration(c *Simultaneous_declarationContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitLhse is called when exiting the lhse production.
	ExitLhse(c *LhseContext)

	// ExitLeft_hand_side_list is called when exiting the left_hand_side_list production.
	ExitLeft_hand_side_list(c *Left_hand_side_listContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSimple_command is called when exiting the simple_command production.
	ExitSimple_command(c *Simple_commandContext)

	// ExitGoto_command is called when exiting the goto_command production.
	ExitGoto_command(c *Goto_commandContext)

	// ExitRoutine_command is called when exiting the routine_command production.
	ExitRoutine_command(c *Routine_commandContext)

	// ExitResultis_command is called when exiting the resultis_command production.
	ExitResultis_command(c *Resultis_commandContext)

	// ExitSwitchon_command is called when exiting the switchon_command production.
	ExitSwitchon_command(c *Switchon_commandContext)

	// ExitRepeatable_command is called when exiting the repeatable_command production.
	ExitRepeatable_command(c *Repeatable_commandContext)

	// ExitRepeated_command is called when exiting the repeated_command production.
	ExitRepeated_command(c *Repeated_commandContext)

	// ExitUntil_command is called when exiting the until_command production.
	ExitUntil_command(c *Until_commandContext)

	// ExitWhile_command is called when exiting the while_command production.
	ExitWhile_command(c *While_commandContext)

	// ExitFor_command is called when exiting the for_command production.
	ExitFor_command(c *For_commandContext)

	// ExitRepetitive_command is called when exiting the repetitive_command production.
	ExitRepetitive_command(c *Repetitive_commandContext)

	// ExitTest_command is called when exiting the test_command production.
	ExitTest_command(c *Test_commandContext)

	// ExitIf_command is called when exiting the if_command production.
	ExitIf_command(c *If_commandContext)

	// ExitUnless_command is called when exiting the unless_command production.
	ExitUnless_command(c *Unless_commandContext)

	// ExitUnlabelled_command is called when exiting the unlabelled_command production.
	ExitUnlabelled_command(c *Unlabelled_commandContext)

	// ExitLabel_prefix is called when exiting the label_prefix production.
	ExitLabel_prefix(c *Label_prefixContext)

	// ExitCase_prefix is called when exiting the case_prefix production.
	ExitCase_prefix(c *Case_prefixContext)

	// ExitDefault_prefix is called when exiting the default_prefix production.
	ExitDefault_prefix(c *Default_prefixContext)

	// ExitPrefix_ is called when exiting the prefix_ production.
	ExitPrefix_(c *Prefix_Context)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitCommand_list is called when exiting the command_list production.
	ExitCommand_list(c *Command_listContext)

	// ExitDeclaration_part is called when exiting the declaration_part production.
	ExitDeclaration_part(c *Declaration_partContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitCompound_command is called when exiting the compound_command production.
	ExitCompound_command(c *Compound_commandContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)
}
