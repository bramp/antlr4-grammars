// Code generated from PowerBuilderDWParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerbuilderdw // PowerBuilderDWParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePowerBuilderDWParserListener is a complete listener for a parse tree produced by PowerBuilderDWParser.
type BasePowerBuilderDWParserListener struct{}

var _ PowerBuilderDWParserListener = &BasePowerBuilderDWParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePowerBuilderDWParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePowerBuilderDWParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePowerBuilderDWParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePowerBuilderDWParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart_rule is called when production start_rule is entered.
func (s *BasePowerBuilderDWParserListener) EnterStart_rule(ctx *Start_ruleContext) {}

// ExitStart_rule is called when production start_rule is exited.
func (s *BasePowerBuilderDWParserListener) ExitStart_rule(ctx *Start_ruleContext) {}

// EnterHeader_rule is called when production header_rule is entered.
func (s *BasePowerBuilderDWParserListener) EnterHeader_rule(ctx *Header_ruleContext) {}

// ExitHeader_rule is called when production header_rule is exited.
func (s *BasePowerBuilderDWParserListener) ExitHeader_rule(ctx *Header_ruleContext) {}

// EnterDatawindow_rule is called when production datawindow_rule is entered.
func (s *BasePowerBuilderDWParserListener) EnterDatawindow_rule(ctx *Datawindow_ruleContext) {}

// ExitDatawindow_rule is called when production datawindow_rule is exited.
func (s *BasePowerBuilderDWParserListener) ExitDatawindow_rule(ctx *Datawindow_ruleContext) {}

// EnterDatawindow_property is called when production datawindow_property is entered.
func (s *BasePowerBuilderDWParserListener) EnterDatawindow_property(ctx *Datawindow_propertyContext) {
}

// ExitDatawindow_property is called when production datawindow_property is exited.
func (s *BasePowerBuilderDWParserListener) ExitDatawindow_property(ctx *Datawindow_propertyContext) {}

// EnterTable_attribute is called when production table_attribute is entered.
func (s *BasePowerBuilderDWParserListener) EnterTable_attribute(ctx *Table_attributeContext) {}

// ExitTable_attribute is called when production table_attribute is exited.
func (s *BasePowerBuilderDWParserListener) ExitTable_attribute(ctx *Table_attributeContext) {}

// EnterColumn_attribute is called when production column_attribute is entered.
func (s *BasePowerBuilderDWParserListener) EnterColumn_attribute(ctx *Column_attributeContext) {}

// ExitColumn_attribute is called when production column_attribute is exited.
func (s *BasePowerBuilderDWParserListener) ExitColumn_attribute(ctx *Column_attributeContext) {}

// EnterRetrieve_attribute is called when production retrieve_attribute is entered.
func (s *BasePowerBuilderDWParserListener) EnterRetrieve_attribute(ctx *Retrieve_attributeContext) {}

// ExitRetrieve_attribute is called when production retrieve_attribute is exited.
func (s *BasePowerBuilderDWParserListener) ExitRetrieve_attribute(ctx *Retrieve_attributeContext) {}

// EnterDatawindow_property_attribute_sub is called when production datawindow_property_attribute_sub is entered.
func (s *BasePowerBuilderDWParserListener) EnterDatawindow_property_attribute_sub(ctx *Datawindow_property_attribute_subContext) {
}

// ExitDatawindow_property_attribute_sub is called when production datawindow_property_attribute_sub is exited.
func (s *BasePowerBuilderDWParserListener) ExitDatawindow_property_attribute_sub(ctx *Datawindow_property_attribute_subContext) {
}

// EnterAttribute_name is called when production attribute_name is entered.
func (s *BasePowerBuilderDWParserListener) EnterAttribute_name(ctx *Attribute_nameContext) {}

// ExitAttribute_name is called when production attribute_name is exited.
func (s *BasePowerBuilderDWParserListener) ExitAttribute_name(ctx *Attribute_nameContext) {}

// EnterIdentifier_name is called when production identifier_name is entered.
func (s *BasePowerBuilderDWParserListener) EnterIdentifier_name(ctx *Identifier_nameContext) {}

// ExitIdentifier_name is called when production identifier_name is exited.
func (s *BasePowerBuilderDWParserListener) ExitIdentifier_name(ctx *Identifier_nameContext) {}

// EnterAttribute_value is called when production attribute_value is entered.
func (s *BasePowerBuilderDWParserListener) EnterAttribute_value(ctx *Attribute_valueContext) {}

// ExitAttribute_value is called when production attribute_value is exited.
func (s *BasePowerBuilderDWParserListener) ExitAttribute_value(ctx *Attribute_valueContext) {}

// EnterNumeric_atom is called when production numeric_atom is entered.
func (s *BasePowerBuilderDWParserListener) EnterNumeric_atom(ctx *Numeric_atomContext) {}

// ExitNumeric_atom is called when production numeric_atom is exited.
func (s *BasePowerBuilderDWParserListener) ExitNumeric_atom(ctx *Numeric_atomContext) {}

// EnterDataTypeSub is called when production dataTypeSub is entered.
func (s *BasePowerBuilderDWParserListener) EnterDataTypeSub(ctx *DataTypeSubContext) {}

// ExitDataTypeSub is called when production dataTypeSub is exited.
func (s *BasePowerBuilderDWParserListener) ExitDataTypeSub(ctx *DataTypeSubContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePowerBuilderDWParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePowerBuilderDWParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArray_decl_sub is called when production array_decl_sub is entered.
func (s *BasePowerBuilderDWParserListener) EnterArray_decl_sub(ctx *Array_decl_subContext) {}

// ExitArray_decl_sub is called when production array_decl_sub is exited.
func (s *BasePowerBuilderDWParserListener) ExitArray_decl_sub(ctx *Array_decl_subContext) {}

// EnterClose_call_sub is called when production close_call_sub is entered.
func (s *BasePowerBuilderDWParserListener) EnterClose_call_sub(ctx *Close_call_subContext) {}

// ExitClose_call_sub is called when production close_call_sub is exited.
func (s *BasePowerBuilderDWParserListener) ExitClose_call_sub(ctx *Close_call_subContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasePowerBuilderDWParserListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasePowerBuilderDWParserListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterAtom_sub_call1 is called when production atom_sub_call1 is entered.
func (s *BasePowerBuilderDWParserListener) EnterAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// ExitAtom_sub_call1 is called when production atom_sub_call1 is exited.
func (s *BasePowerBuilderDWParserListener) ExitAtom_sub_call1(ctx *Atom_sub_call1Context) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePowerBuilderDWParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePowerBuilderDWParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifier_name_ex is called when production identifier_name_ex is entered.
func (s *BasePowerBuilderDWParserListener) EnterIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// ExitIdentifier_name_ex is called when production identifier_name_ex is exited.
func (s *BasePowerBuilderDWParserListener) ExitIdentifier_name_ex(ctx *Identifier_name_exContext) {}

// EnterAtom_sub_member1 is called when production atom_sub_member1 is entered.
func (s *BasePowerBuilderDWParserListener) EnterAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// ExitAtom_sub_member1 is called when production atom_sub_member1 is exited.
func (s *BasePowerBuilderDWParserListener) ExitAtom_sub_member1(ctx *Atom_sub_member1Context) {}

// EnterBoolean_atom is called when production boolean_atom is entered.
func (s *BasePowerBuilderDWParserListener) EnterBoolean_atom(ctx *Boolean_atomContext) {}

// ExitBoolean_atom is called when production boolean_atom is exited.
func (s *BasePowerBuilderDWParserListener) ExitBoolean_atom(ctx *Boolean_atomContext) {}
