// Code generated from bcpl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bcpl // bcpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebcplListener is a complete listener for a parse tree produced by bcplParser.
type BasebcplListener struct{}

var _ bcplListener = &BasebcplListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebcplListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebcplListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebcplListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebcplListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLetter is called when production letter is entered.
func (s *BasebcplListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BasebcplListener) ExitLetter(ctx *LetterContext) {}

// EnterOctal_digit is called when production octal_digit is entered.
func (s *BasebcplListener) EnterOctal_digit(ctx *Octal_digitContext) {}

// ExitOctal_digit is called when production octal_digit is exited.
func (s *BasebcplListener) ExitOctal_digit(ctx *Octal_digitContext) {}

// EnterHex_digit is called when production hex_digit is entered.
func (s *BasebcplListener) EnterHex_digit(ctx *Hex_digitContext) {}

// ExitHex_digit is called when production hex_digit is exited.
func (s *BasebcplListener) ExitHex_digit(ctx *Hex_digitContext) {}

// EnterDigit is called when production digit is entered.
func (s *BasebcplListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BasebcplListener) ExitDigit(ctx *DigitContext) {}

// EnterString_constant is called when production string_constant is entered.
func (s *BasebcplListener) EnterString_constant(ctx *String_constantContext) {}

// ExitString_constant is called when production string_constant is exited.
func (s *BasebcplListener) ExitString_constant(ctx *String_constantContext) {}

// EnterCharacter_constant is called when production character_constant is entered.
func (s *BasebcplListener) EnterCharacter_constant(ctx *Character_constantContext) {}

// ExitCharacter_constant is called when production character_constant is exited.
func (s *BasebcplListener) ExitCharacter_constant(ctx *Character_constantContext) {}

// EnterOctal_number is called when production octal_number is entered.
func (s *BasebcplListener) EnterOctal_number(ctx *Octal_numberContext) {}

// ExitOctal_number is called when production octal_number is exited.
func (s *BasebcplListener) ExitOctal_number(ctx *Octal_numberContext) {}

// EnterHex_number is called when production hex_number is entered.
func (s *BasebcplListener) EnterHex_number(ctx *Hex_numberContext) {}

// ExitHex_number is called when production hex_number is exited.
func (s *BasebcplListener) ExitHex_number(ctx *Hex_numberContext) {}

// EnterNumber is called when production number is entered.
func (s *BasebcplListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasebcplListener) ExitNumber(ctx *NumberContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasebcplListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasebcplListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterOne_character is called when production one_character is entered.
func (s *BasebcplListener) EnterOne_character(ctx *One_characterContext) {}

// ExitOne_character is called when production one_character is exited.
func (s *BasebcplListener) ExitOne_character(ctx *One_characterContext) {}

// EnterName is called when production name is entered.
func (s *BasebcplListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasebcplListener) ExitName(ctx *NameContext) {}

// EnterAddress_op is called when production address_op is entered.
func (s *BasebcplListener) EnterAddress_op(ctx *Address_opContext) {}

// ExitAddress_op is called when production address_op is exited.
func (s *BasebcplListener) ExitAddress_op(ctx *Address_opContext) {}

// EnterMult_op is called when production mult_op is entered.
func (s *BasebcplListener) EnterMult_op(ctx *Mult_opContext) {}

// ExitMult_op is called when production mult_op is exited.
func (s *BasebcplListener) ExitMult_op(ctx *Mult_opContext) {}

// EnterAdd_op is called when production add_op is entered.
func (s *BasebcplListener) EnterAdd_op(ctx *Add_opContext) {}

// ExitAdd_op is called when production add_op is exited.
func (s *BasebcplListener) ExitAdd_op(ctx *Add_opContext) {}

// EnterRel_op is called when production rel_op is entered.
func (s *BasebcplListener) EnterRel_op(ctx *Rel_opContext) {}

// ExitRel_op is called when production rel_op is exited.
func (s *BasebcplListener) ExitRel_op(ctx *Rel_opContext) {}

// EnterShift_op is called when production shift_op is entered.
func (s *BasebcplListener) EnterShift_op(ctx *Shift_opContext) {}

// ExitShift_op is called when production shift_op is exited.
func (s *BasebcplListener) ExitShift_op(ctx *Shift_opContext) {}

// EnterAnd_op is called when production and_op is entered.
func (s *BasebcplListener) EnterAnd_op(ctx *And_opContext) {}

// ExitAnd_op is called when production and_op is exited.
func (s *BasebcplListener) ExitAnd_op(ctx *And_opContext) {}

// EnterOr_op is called when production or_op is entered.
func (s *BasebcplListener) EnterOr_op(ctx *Or_opContext) {}

// ExitOr_op is called when production or_op is exited.
func (s *BasebcplListener) ExitOr_op(ctx *Or_opContext) {}

// EnterEqv_op is called when production eqv_op is entered.
func (s *BasebcplListener) EnterEqv_op(ctx *Eqv_opContext) {}

// ExitEqv_op is called when production eqv_op is exited.
func (s *BasebcplListener) ExitEqv_op(ctx *Eqv_opContext) {}

// EnterNot_op is called when production not_op is entered.
func (s *BasebcplListener) EnterNot_op(ctx *Not_opContext) {}

// ExitNot_op is called when production not_op is exited.
func (s *BasebcplListener) ExitNot_op(ctx *Not_opContext) {}

// EnterElement is called when production element is entered.
func (s *BasebcplListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasebcplListener) ExitElement(ctx *ElementContext) {}

// EnterPrimary_E is called when production primary_E is entered.
func (s *BasebcplListener) EnterPrimary_E(ctx *Primary_EContext) {}

// ExitPrimary_E is called when production primary_E is exited.
func (s *BasebcplListener) ExitPrimary_E(ctx *Primary_EContext) {}

// EnterVector_E is called when production vector_E is entered.
func (s *BasebcplListener) EnterVector_E(ctx *Vector_EContext) {}

// ExitVector_E is called when production vector_E is exited.
func (s *BasebcplListener) ExitVector_E(ctx *Vector_EContext) {}

// EnterAddress_E is called when production address_E is entered.
func (s *BasebcplListener) EnterAddress_E(ctx *Address_EContext) {}

// ExitAddress_E is called when production address_E is exited.
func (s *BasebcplListener) ExitAddress_E(ctx *Address_EContext) {}

// EnterMult_E is called when production mult_E is entered.
func (s *BasebcplListener) EnterMult_E(ctx *Mult_EContext) {}

// ExitMult_E is called when production mult_E is exited.
func (s *BasebcplListener) ExitMult_E(ctx *Mult_EContext) {}

// EnterAdd_E is called when production add_E is entered.
func (s *BasebcplListener) EnterAdd_E(ctx *Add_EContext) {}

// ExitAdd_E is called when production add_E is exited.
func (s *BasebcplListener) ExitAdd_E(ctx *Add_EContext) {}

// EnterRel_E is called when production rel_E is entered.
func (s *BasebcplListener) EnterRel_E(ctx *Rel_EContext) {}

// ExitRel_E is called when production rel_E is exited.
func (s *BasebcplListener) ExitRel_E(ctx *Rel_EContext) {}

// EnterShift_E is called when production shift_E is entered.
func (s *BasebcplListener) EnterShift_E(ctx *Shift_EContext) {}

// ExitShift_E is called when production shift_E is exited.
func (s *BasebcplListener) ExitShift_E(ctx *Shift_EContext) {}

// EnterNot_E is called when production not_E is entered.
func (s *BasebcplListener) EnterNot_E(ctx *Not_EContext) {}

// ExitNot_E is called when production not_E is exited.
func (s *BasebcplListener) ExitNot_E(ctx *Not_EContext) {}

// EnterAnd_E is called when production and_E is entered.
func (s *BasebcplListener) EnterAnd_E(ctx *And_EContext) {}

// ExitAnd_E is called when production and_E is exited.
func (s *BasebcplListener) ExitAnd_E(ctx *And_EContext) {}

// EnterOr_E is called when production or_E is entered.
func (s *BasebcplListener) EnterOr_E(ctx *Or_EContext) {}

// ExitOr_E is called when production or_E is exited.
func (s *BasebcplListener) ExitOr_E(ctx *Or_EContext) {}

// EnterEqv_E is called when production eqv_E is entered.
func (s *BasebcplListener) EnterEqv_E(ctx *Eqv_EContext) {}

// ExitEqv_E is called when production eqv_E is exited.
func (s *BasebcplListener) ExitEqv_E(ctx *Eqv_EContext) {}

// EnterConditional_E is called when production conditional_E is entered.
func (s *BasebcplListener) EnterConditional_E(ctx *Conditional_EContext) {}

// ExitConditional_E is called when production conditional_E is exited.
func (s *BasebcplListener) ExitConditional_E(ctx *Conditional_EContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasebcplListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasebcplListener) ExitExpression(ctx *ExpressionContext) {}

// EnterC_element is called when production c_element is entered.
func (s *BasebcplListener) EnterC_element(ctx *C_elementContext) {}

// ExitC_element is called when production c_element is exited.
func (s *BasebcplListener) ExitC_element(ctx *C_elementContext) {}

// EnterC_mult_E is called when production c_mult_E is entered.
func (s *BasebcplListener) EnterC_mult_E(ctx *C_mult_EContext) {}

// ExitC_mult_E is called when production c_mult_E is exited.
func (s *BasebcplListener) ExitC_mult_E(ctx *C_mult_EContext) {}

// EnterC_add_E is called when production c_add_E is entered.
func (s *BasebcplListener) EnterC_add_E(ctx *C_add_EContext) {}

// ExitC_add_E is called when production c_add_E is exited.
func (s *BasebcplListener) ExitC_add_E(ctx *C_add_EContext) {}

// EnterC_shift_E is called when production c_shift_E is entered.
func (s *BasebcplListener) EnterC_shift_E(ctx *C_shift_EContext) {}

// ExitC_shift_E is called when production c_shift_E is exited.
func (s *BasebcplListener) ExitC_shift_E(ctx *C_shift_EContext) {}

// EnterC_and_E is called when production c_and_E is entered.
func (s *BasebcplListener) EnterC_and_E(ctx *C_and_EContext) {}

// ExitC_and_E is called when production c_and_E is exited.
func (s *BasebcplListener) ExitC_and_E(ctx *C_and_EContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BasebcplListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BasebcplListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasebcplListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasebcplListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BasebcplListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BasebcplListener) ExitName_list(ctx *Name_listContext) {}

// EnterManifest_item is called when production manifest_item is entered.
func (s *BasebcplListener) EnterManifest_item(ctx *Manifest_itemContext) {}

// ExitManifest_item is called when production manifest_item is exited.
func (s *BasebcplListener) ExitManifest_item(ctx *Manifest_itemContext) {}

// EnterManifest_list is called when production manifest_list is entered.
func (s *BasebcplListener) EnterManifest_list(ctx *Manifest_listContext) {}

// ExitManifest_list is called when production manifest_list is exited.
func (s *BasebcplListener) ExitManifest_list(ctx *Manifest_listContext) {}

// EnterManifest_declaration is called when production manifest_declaration is entered.
func (s *BasebcplListener) EnterManifest_declaration(ctx *Manifest_declarationContext) {}

// ExitManifest_declaration is called when production manifest_declaration is exited.
func (s *BasebcplListener) ExitManifest_declaration(ctx *Manifest_declarationContext) {}

// EnterStatic_declaration is called when production static_declaration is entered.
func (s *BasebcplListener) EnterStatic_declaration(ctx *Static_declarationContext) {}

// ExitStatic_declaration is called when production static_declaration is exited.
func (s *BasebcplListener) ExitStatic_declaration(ctx *Static_declarationContext) {}

// EnterGlobal_item is called when production global_item is entered.
func (s *BasebcplListener) EnterGlobal_item(ctx *Global_itemContext) {}

// ExitGlobal_item is called when production global_item is exited.
func (s *BasebcplListener) ExitGlobal_item(ctx *Global_itemContext) {}

// EnterGlobal_list is called when production global_list is entered.
func (s *BasebcplListener) EnterGlobal_list(ctx *Global_listContext) {}

// ExitGlobal_list is called when production global_list is exited.
func (s *BasebcplListener) ExitGlobal_list(ctx *Global_listContext) {}

// EnterGlobal_declaration is called when production global_declaration is entered.
func (s *BasebcplListener) EnterGlobal_declaration(ctx *Global_declarationContext) {}

// ExitGlobal_declaration is called when production global_declaration is exited.
func (s *BasebcplListener) ExitGlobal_declaration(ctx *Global_declarationContext) {}

// EnterSimple_definition is called when production simple_definition is entered.
func (s *BasebcplListener) EnterSimple_definition(ctx *Simple_definitionContext) {}

// ExitSimple_definition is called when production simple_definition is exited.
func (s *BasebcplListener) ExitSimple_definition(ctx *Simple_definitionContext) {}

// EnterVector_definition is called when production vector_definition is entered.
func (s *BasebcplListener) EnterVector_definition(ctx *Vector_definitionContext) {}

// ExitVector_definition is called when production vector_definition is exited.
func (s *BasebcplListener) ExitVector_definition(ctx *Vector_definitionContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BasebcplListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BasebcplListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterRoutine_definition is called when production routine_definition is entered.
func (s *BasebcplListener) EnterRoutine_definition(ctx *Routine_definitionContext) {}

// ExitRoutine_definition is called when production routine_definition is exited.
func (s *BasebcplListener) ExitRoutine_definition(ctx *Routine_definitionContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BasebcplListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BasebcplListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterSimultaneous_declaration is called when production simultaneous_declaration is entered.
func (s *BasebcplListener) EnterSimultaneous_declaration(ctx *Simultaneous_declarationContext) {}

// ExitSimultaneous_declaration is called when production simultaneous_declaration is exited.
func (s *BasebcplListener) ExitSimultaneous_declaration(ctx *Simultaneous_declarationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasebcplListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasebcplListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterLhse is called when production lhse is entered.
func (s *BasebcplListener) EnterLhse(ctx *LhseContext) {}

// ExitLhse is called when production lhse is exited.
func (s *BasebcplListener) ExitLhse(ctx *LhseContext) {}

// EnterLeft_hand_side_list is called when production left_hand_side_list is entered.
func (s *BasebcplListener) EnterLeft_hand_side_list(ctx *Left_hand_side_listContext) {}

// ExitLeft_hand_side_list is called when production left_hand_side_list is exited.
func (s *BasebcplListener) ExitLeft_hand_side_list(ctx *Left_hand_side_listContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasebcplListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasebcplListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSimple_command is called when production simple_command is entered.
func (s *BasebcplListener) EnterSimple_command(ctx *Simple_commandContext) {}

// ExitSimple_command is called when production simple_command is exited.
func (s *BasebcplListener) ExitSimple_command(ctx *Simple_commandContext) {}

// EnterGoto_command is called when production goto_command is entered.
func (s *BasebcplListener) EnterGoto_command(ctx *Goto_commandContext) {}

// ExitGoto_command is called when production goto_command is exited.
func (s *BasebcplListener) ExitGoto_command(ctx *Goto_commandContext) {}

// EnterRoutine_command is called when production routine_command is entered.
func (s *BasebcplListener) EnterRoutine_command(ctx *Routine_commandContext) {}

// ExitRoutine_command is called when production routine_command is exited.
func (s *BasebcplListener) ExitRoutine_command(ctx *Routine_commandContext) {}

// EnterResultis_command is called when production resultis_command is entered.
func (s *BasebcplListener) EnterResultis_command(ctx *Resultis_commandContext) {}

// ExitResultis_command is called when production resultis_command is exited.
func (s *BasebcplListener) ExitResultis_command(ctx *Resultis_commandContext) {}

// EnterSwitchon_command is called when production switchon_command is entered.
func (s *BasebcplListener) EnterSwitchon_command(ctx *Switchon_commandContext) {}

// ExitSwitchon_command is called when production switchon_command is exited.
func (s *BasebcplListener) ExitSwitchon_command(ctx *Switchon_commandContext) {}

// EnterRepeatable_command is called when production repeatable_command is entered.
func (s *BasebcplListener) EnterRepeatable_command(ctx *Repeatable_commandContext) {}

// ExitRepeatable_command is called when production repeatable_command is exited.
func (s *BasebcplListener) ExitRepeatable_command(ctx *Repeatable_commandContext) {}

// EnterRepeated_command is called when production repeated_command is entered.
func (s *BasebcplListener) EnterRepeated_command(ctx *Repeated_commandContext) {}

// ExitRepeated_command is called when production repeated_command is exited.
func (s *BasebcplListener) ExitRepeated_command(ctx *Repeated_commandContext) {}

// EnterUntil_command is called when production until_command is entered.
func (s *BasebcplListener) EnterUntil_command(ctx *Until_commandContext) {}

// ExitUntil_command is called when production until_command is exited.
func (s *BasebcplListener) ExitUntil_command(ctx *Until_commandContext) {}

// EnterWhile_command is called when production while_command is entered.
func (s *BasebcplListener) EnterWhile_command(ctx *While_commandContext) {}

// ExitWhile_command is called when production while_command is exited.
func (s *BasebcplListener) ExitWhile_command(ctx *While_commandContext) {}

// EnterFor_command is called when production for_command is entered.
func (s *BasebcplListener) EnterFor_command(ctx *For_commandContext) {}

// ExitFor_command is called when production for_command is exited.
func (s *BasebcplListener) ExitFor_command(ctx *For_commandContext) {}

// EnterRepetitive_command is called when production repetitive_command is entered.
func (s *BasebcplListener) EnterRepetitive_command(ctx *Repetitive_commandContext) {}

// ExitRepetitive_command is called when production repetitive_command is exited.
func (s *BasebcplListener) ExitRepetitive_command(ctx *Repetitive_commandContext) {}

// EnterTest_command is called when production test_command is entered.
func (s *BasebcplListener) EnterTest_command(ctx *Test_commandContext) {}

// ExitTest_command is called when production test_command is exited.
func (s *BasebcplListener) ExitTest_command(ctx *Test_commandContext) {}

// EnterIf_command is called when production if_command is entered.
func (s *BasebcplListener) EnterIf_command(ctx *If_commandContext) {}

// ExitIf_command is called when production if_command is exited.
func (s *BasebcplListener) ExitIf_command(ctx *If_commandContext) {}

// EnterUnless_command is called when production unless_command is entered.
func (s *BasebcplListener) EnterUnless_command(ctx *Unless_commandContext) {}

// ExitUnless_command is called when production unless_command is exited.
func (s *BasebcplListener) ExitUnless_command(ctx *Unless_commandContext) {}

// EnterUnlabelled_command is called when production unlabelled_command is entered.
func (s *BasebcplListener) EnterUnlabelled_command(ctx *Unlabelled_commandContext) {}

// ExitUnlabelled_command is called when production unlabelled_command is exited.
func (s *BasebcplListener) ExitUnlabelled_command(ctx *Unlabelled_commandContext) {}

// EnterLabel_prefix is called when production label_prefix is entered.
func (s *BasebcplListener) EnterLabel_prefix(ctx *Label_prefixContext) {}

// ExitLabel_prefix is called when production label_prefix is exited.
func (s *BasebcplListener) ExitLabel_prefix(ctx *Label_prefixContext) {}

// EnterCase_prefix is called when production case_prefix is entered.
func (s *BasebcplListener) EnterCase_prefix(ctx *Case_prefixContext) {}

// ExitCase_prefix is called when production case_prefix is exited.
func (s *BasebcplListener) ExitCase_prefix(ctx *Case_prefixContext) {}

// EnterDefault_prefix is called when production default_prefix is entered.
func (s *BasebcplListener) EnterDefault_prefix(ctx *Default_prefixContext) {}

// ExitDefault_prefix is called when production default_prefix is exited.
func (s *BasebcplListener) ExitDefault_prefix(ctx *Default_prefixContext) {}

// EnterPrefix_ is called when production prefix_ is entered.
func (s *BasebcplListener) EnterPrefix_(ctx *Prefix_Context) {}

// ExitPrefix_ is called when production prefix_ is exited.
func (s *BasebcplListener) ExitPrefix_(ctx *Prefix_Context) {}

// EnterCommand is called when production command is entered.
func (s *BasebcplListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasebcplListener) ExitCommand(ctx *CommandContext) {}

// EnterCommand_list is called when production command_list is entered.
func (s *BasebcplListener) EnterCommand_list(ctx *Command_listContext) {}

// ExitCommand_list is called when production command_list is exited.
func (s *BasebcplListener) ExitCommand_list(ctx *Command_listContext) {}

// EnterDeclaration_part is called when production declaration_part is entered.
func (s *BasebcplListener) EnterDeclaration_part(ctx *Declaration_partContext) {}

// ExitDeclaration_part is called when production declaration_part is exited.
func (s *BasebcplListener) ExitDeclaration_part(ctx *Declaration_partContext) {}

// EnterBlock is called when production block is entered.
func (s *BasebcplListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasebcplListener) ExitBlock(ctx *BlockContext) {}

// EnterCompound_command is called when production compound_command is entered.
func (s *BasebcplListener) EnterCompound_command(ctx *Compound_commandContext) {}

// ExitCompound_command is called when production compound_command is exited.
func (s *BasebcplListener) ExitCompound_command(ctx *Compound_commandContext) {}

// EnterProgram is called when production program is entered.
func (s *BasebcplListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasebcplListener) ExitProgram(ctx *ProgramContext) {}

// EnterDirective is called when production directive is entered.
func (s *BasebcplListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BasebcplListener) ExitDirective(ctx *DirectiveContext) {}
