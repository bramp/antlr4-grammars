// Code generated from PowerBuilderDWParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerbuilderdw // PowerBuilderDWParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PowerBuilderDWParserListener is a complete listener for a parse tree produced by PowerBuilderDWParser.
type PowerBuilderDWParserListener interface {
	antlr.ParseTreeListener

	// EnterStart_rule is called when entering the start_rule production.
	EnterStart_rule(c *Start_ruleContext)

	// EnterHeader_rule is called when entering the header_rule production.
	EnterHeader_rule(c *Header_ruleContext)

	// EnterDatawindow_rule is called when entering the datawindow_rule production.
	EnterDatawindow_rule(c *Datawindow_ruleContext)

	// EnterDatawindow_property is called when entering the datawindow_property production.
	EnterDatawindow_property(c *Datawindow_propertyContext)

	// EnterTable_attribute is called when entering the table_attribute production.
	EnterTable_attribute(c *Table_attributeContext)

	// EnterColumn_attribute is called when entering the column_attribute production.
	EnterColumn_attribute(c *Column_attributeContext)

	// EnterRetrieve_attribute is called when entering the retrieve_attribute production.
	EnterRetrieve_attribute(c *Retrieve_attributeContext)

	// EnterDatawindow_property_attribute_sub is called when entering the datawindow_property_attribute_sub production.
	EnterDatawindow_property_attribute_sub(c *Datawindow_property_attribute_subContext)

	// EnterAttribute_name is called when entering the attribute_name production.
	EnterAttribute_name(c *Attribute_nameContext)

	// EnterIdentifier_name is called when entering the identifier_name production.
	EnterIdentifier_name(c *Identifier_nameContext)

	// EnterAttribute_value is called when entering the attribute_value production.
	EnterAttribute_value(c *Attribute_valueContext)

	// EnterNumeric_atom is called when entering the numeric_atom production.
	EnterNumeric_atom(c *Numeric_atomContext)

	// EnterDataTypeSub is called when entering the dataTypeSub production.
	EnterDataTypeSub(c *DataTypeSubContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArray_decl_sub is called when entering the array_decl_sub production.
	EnterArray_decl_sub(c *Array_decl_subContext)

	// EnterClose_call_sub is called when entering the close_call_sub production.
	EnterClose_call_sub(c *Close_call_subContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterAtom_sub_call1 is called when entering the atom_sub_call1 production.
	EnterAtom_sub_call1(c *Atom_sub_call1Context)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifier_name_ex is called when entering the identifier_name_ex production.
	EnterIdentifier_name_ex(c *Identifier_name_exContext)

	// EnterAtom_sub_member1 is called when entering the atom_sub_member1 production.
	EnterAtom_sub_member1(c *Atom_sub_member1Context)

	// EnterBoolean_atom is called when entering the boolean_atom production.
	EnterBoolean_atom(c *Boolean_atomContext)

	// ExitStart_rule is called when exiting the start_rule production.
	ExitStart_rule(c *Start_ruleContext)

	// ExitHeader_rule is called when exiting the header_rule production.
	ExitHeader_rule(c *Header_ruleContext)

	// ExitDatawindow_rule is called when exiting the datawindow_rule production.
	ExitDatawindow_rule(c *Datawindow_ruleContext)

	// ExitDatawindow_property is called when exiting the datawindow_property production.
	ExitDatawindow_property(c *Datawindow_propertyContext)

	// ExitTable_attribute is called when exiting the table_attribute production.
	ExitTable_attribute(c *Table_attributeContext)

	// ExitColumn_attribute is called when exiting the column_attribute production.
	ExitColumn_attribute(c *Column_attributeContext)

	// ExitRetrieve_attribute is called when exiting the retrieve_attribute production.
	ExitRetrieve_attribute(c *Retrieve_attributeContext)

	// ExitDatawindow_property_attribute_sub is called when exiting the datawindow_property_attribute_sub production.
	ExitDatawindow_property_attribute_sub(c *Datawindow_property_attribute_subContext)

	// ExitAttribute_name is called when exiting the attribute_name production.
	ExitAttribute_name(c *Attribute_nameContext)

	// ExitIdentifier_name is called when exiting the identifier_name production.
	ExitIdentifier_name(c *Identifier_nameContext)

	// ExitAttribute_value is called when exiting the attribute_value production.
	ExitAttribute_value(c *Attribute_valueContext)

	// ExitNumeric_atom is called when exiting the numeric_atom production.
	ExitNumeric_atom(c *Numeric_atomContext)

	// ExitDataTypeSub is called when exiting the dataTypeSub production.
	ExitDataTypeSub(c *DataTypeSubContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArray_decl_sub is called when exiting the array_decl_sub production.
	ExitArray_decl_sub(c *Array_decl_subContext)

	// ExitClose_call_sub is called when exiting the close_call_sub production.
	ExitClose_call_sub(c *Close_call_subContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitAtom_sub_call1 is called when exiting the atom_sub_call1 production.
	ExitAtom_sub_call1(c *Atom_sub_call1Context)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifier_name_ex is called when exiting the identifier_name_ex production.
	ExitIdentifier_name_ex(c *Identifier_name_exContext)

	// ExitAtom_sub_member1 is called when exiting the atom_sub_member1 production.
	ExitAtom_sub_member1(c *Atom_sub_member1Context)

	// ExitBoolean_atom is called when exiting the boolean_atom production.
	ExitBoolean_atom(c *Boolean_atomContext)
}
