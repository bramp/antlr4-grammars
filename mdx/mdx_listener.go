// Generated from mdx.g4 by ANTLR 4.7.

package mdx // mdx
import "github.com/antlr/antlr4/runtime/Go/antlr"

// mdxListener is a complete listener for a parse tree produced by mdxParser.
type mdxListener interface {
	antlr.ParseTreeListener

	// EnterMdx_statement is called when entering the mdx_statement production.
	EnterMdx_statement(c *Mdx_statementContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterFormula_specification is called when entering the formula_specification production.
	EnterFormula_specification(c *Formula_specificationContext)

	// EnterSingle_formula_specification is called when entering the single_formula_specification production.
	EnterSingle_formula_specification(c *Single_formula_specificationContext)

	// EnterSet_specification is called when entering the set_specification production.
	EnterSet_specification(c *Set_specificationContext)

	// EnterMember_specification is called when entering the member_specification production.
	EnterMember_specification(c *Member_specificationContext)

	// EnterAxis_specification_list is called when entering the axis_specification_list production.
	EnterAxis_specification_list(c *Axis_specification_listContext)

	// EnterMember_property_def_list is called when entering the member_property_def_list production.
	EnterMember_property_def_list(c *Member_property_def_listContext)

	// EnterMember_name is called when entering the member_name production.
	EnterMember_name(c *Member_nameContext)

	// EnterMember_property_definition is called when entering the member_property_definition production.
	EnterMember_property_definition(c *Member_property_definitionContext)

	// EnterSet_name is called when entering the set_name production.
	EnterSet_name(c *Set_nameContext)

	// EnterCompound_id is called when entering the compound_id production.
	EnterCompound_id(c *Compound_idContext)

	// EnterAxis_specification is called when entering the axis_specification production.
	EnterAxis_specification(c *Axis_specificationContext)

	// EnterAxis_name is called when entering the axis_name production.
	EnterAxis_name(c *Axis_nameContext)

	// EnterDim_props is called when entering the dim_props production.
	EnterDim_props(c *Dim_propsContext)

	// EnterProperty_list is called when entering the property_list production.
	EnterProperty_list(c *Property_listContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterCube_specification is called when entering the cube_specification production.
	EnterCube_specification(c *Cube_specificationContext)

	// EnterCube_name is called when entering the cube_name production.
	EnterCube_name(c *Cube_nameContext)

	// EnterSlicer_specification is called when entering the slicer_specification production.
	EnterSlicer_specification(c *Slicer_specificationContext)

	// EnterCell_props is called when entering the cell_props production.
	EnterCell_props(c *Cell_propsContext)

	// EnterCell_property_list is called when entering the cell_property_list production.
	EnterCell_property_list(c *Cell_property_listContext)

	// EnterCell_property is called when entering the cell_property production.
	EnterCell_property(c *Cell_propertyContext)

	// EnterMandatory_cell_property is called when entering the mandatory_cell_property production.
	EnterMandatory_cell_property(c *Mandatory_cell_propertyContext)

	// EnterProvider_specific_cell_property is called when entering the provider_specific_cell_property production.
	EnterProvider_specific_cell_property(c *Provider_specific_cell_propertyContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterValue_expression is called when entering the value_expression production.
	EnterValue_expression(c *Value_expressionContext)

	// EnterValue_xor_expression is called when entering the value_xor_expression production.
	EnterValue_xor_expression(c *Value_xor_expressionContext)

	// EnterValue_or_expression is called when entering the value_or_expression production.
	EnterValue_or_expression(c *Value_or_expressionContext)

	// EnterTerm5 is called when entering the term5 production.
	EnterTerm5(c *Term5Context)

	// EnterTerm4 is called when entering the term4 production.
	EnterTerm4(c *Term4Context)

	// EnterTerm3 is called when entering the term3 production.
	EnterTerm3(c *Term3Context)

	// EnterTerm2 is called when entering the term2 production.
	EnterTerm2(c *Term2Context)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterValue_expression_primary is called when entering the value_expression_primary production.
	EnterValue_expression_primary(c *Value_expression_primaryContext)

	// EnterValue_expression_primary0 is called when entering the value_expression_primary0 production.
	EnterValue_expression_primary0(c *Value_expression_primary0Context)

	// EnterExp_list is called when entering the exp_list production.
	EnterExp_list(c *Exp_listContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterWhen_list is called when entering the when_list production.
	EnterWhen_list(c *When_listContext)

	// EnterWhen_clause is called when entering the when_clause production.
	EnterWhen_clause(c *When_clauseContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterUnquoted_identifier is called when entering the unquoted_identifier production.
	EnterUnquoted_identifier(c *Unquoted_identifierContext)

	// EnterAmp_quoted_identifier is called when entering the amp_quoted_identifier production.
	EnterAmp_quoted_identifier(c *Amp_quoted_identifierContext)

	// EnterQuoted_identifier is called when entering the quoted_identifier production.
	EnterQuoted_identifier(c *Quoted_identifierContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitMdx_statement is called when exiting the mdx_statement production.
	ExitMdx_statement(c *Mdx_statementContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitFormula_specification is called when exiting the formula_specification production.
	ExitFormula_specification(c *Formula_specificationContext)

	// ExitSingle_formula_specification is called when exiting the single_formula_specification production.
	ExitSingle_formula_specification(c *Single_formula_specificationContext)

	// ExitSet_specification is called when exiting the set_specification production.
	ExitSet_specification(c *Set_specificationContext)

	// ExitMember_specification is called when exiting the member_specification production.
	ExitMember_specification(c *Member_specificationContext)

	// ExitAxis_specification_list is called when exiting the axis_specification_list production.
	ExitAxis_specification_list(c *Axis_specification_listContext)

	// ExitMember_property_def_list is called when exiting the member_property_def_list production.
	ExitMember_property_def_list(c *Member_property_def_listContext)

	// ExitMember_name is called when exiting the member_name production.
	ExitMember_name(c *Member_nameContext)

	// ExitMember_property_definition is called when exiting the member_property_definition production.
	ExitMember_property_definition(c *Member_property_definitionContext)

	// ExitSet_name is called when exiting the set_name production.
	ExitSet_name(c *Set_nameContext)

	// ExitCompound_id is called when exiting the compound_id production.
	ExitCompound_id(c *Compound_idContext)

	// ExitAxis_specification is called when exiting the axis_specification production.
	ExitAxis_specification(c *Axis_specificationContext)

	// ExitAxis_name is called when exiting the axis_name production.
	ExitAxis_name(c *Axis_nameContext)

	// ExitDim_props is called when exiting the dim_props production.
	ExitDim_props(c *Dim_propsContext)

	// ExitProperty_list is called when exiting the property_list production.
	ExitProperty_list(c *Property_listContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitCube_specification is called when exiting the cube_specification production.
	ExitCube_specification(c *Cube_specificationContext)

	// ExitCube_name is called when exiting the cube_name production.
	ExitCube_name(c *Cube_nameContext)

	// ExitSlicer_specification is called when exiting the slicer_specification production.
	ExitSlicer_specification(c *Slicer_specificationContext)

	// ExitCell_props is called when exiting the cell_props production.
	ExitCell_props(c *Cell_propsContext)

	// ExitCell_property_list is called when exiting the cell_property_list production.
	ExitCell_property_list(c *Cell_property_listContext)

	// ExitCell_property is called when exiting the cell_property production.
	ExitCell_property(c *Cell_propertyContext)

	// ExitMandatory_cell_property is called when exiting the mandatory_cell_property production.
	ExitMandatory_cell_property(c *Mandatory_cell_propertyContext)

	// ExitProvider_specific_cell_property is called when exiting the provider_specific_cell_property production.
	ExitProvider_specific_cell_property(c *Provider_specific_cell_propertyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitValue_expression is called when exiting the value_expression production.
	ExitValue_expression(c *Value_expressionContext)

	// ExitValue_xor_expression is called when exiting the value_xor_expression production.
	ExitValue_xor_expression(c *Value_xor_expressionContext)

	// ExitValue_or_expression is called when exiting the value_or_expression production.
	ExitValue_or_expression(c *Value_or_expressionContext)

	// ExitTerm5 is called when exiting the term5 production.
	ExitTerm5(c *Term5Context)

	// ExitTerm4 is called when exiting the term4 production.
	ExitTerm4(c *Term4Context)

	// ExitTerm3 is called when exiting the term3 production.
	ExitTerm3(c *Term3Context)

	// ExitTerm2 is called when exiting the term2 production.
	ExitTerm2(c *Term2Context)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitValue_expression_primary is called when exiting the value_expression_primary production.
	ExitValue_expression_primary(c *Value_expression_primaryContext)

	// ExitValue_expression_primary0 is called when exiting the value_expression_primary0 production.
	ExitValue_expression_primary0(c *Value_expression_primary0Context)

	// ExitExp_list is called when exiting the exp_list production.
	ExitExp_list(c *Exp_listContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitWhen_list is called when exiting the when_list production.
	ExitWhen_list(c *When_listContext)

	// ExitWhen_clause is called when exiting the when_clause production.
	ExitWhen_clause(c *When_clauseContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitUnquoted_identifier is called when exiting the unquoted_identifier production.
	ExitUnquoted_identifier(c *Unquoted_identifierContext)

	// ExitAmp_quoted_identifier is called when exiting the amp_quoted_identifier production.
	ExitAmp_quoted_identifier(c *Amp_quoted_identifierContext)

	// ExitQuoted_identifier is called when exiting the quoted_identifier production.
	ExitQuoted_identifier(c *Quoted_identifierContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}
