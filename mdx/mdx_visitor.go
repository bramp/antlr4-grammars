// Generated from mdx.g4 by ANTLR 4.7.

package mdx // mdx
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by mdxParser.
type mdxVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mdxParser#mdx_statement.
	VisitMdx_statement(ctx *Mdx_statementContext) interface{}

	// Visit a parse tree produced by mdxParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by mdxParser#formula_specification.
	VisitFormula_specification(ctx *Formula_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#single_formula_specification.
	VisitSingle_formula_specification(ctx *Single_formula_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#set_specification.
	VisitSet_specification(ctx *Set_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#member_specification.
	VisitMember_specification(ctx *Member_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#axis_specification_list.
	VisitAxis_specification_list(ctx *Axis_specification_listContext) interface{}

	// Visit a parse tree produced by mdxParser#member_property_def_list.
	VisitMember_property_def_list(ctx *Member_property_def_listContext) interface{}

	// Visit a parse tree produced by mdxParser#member_name.
	VisitMember_name(ctx *Member_nameContext) interface{}

	// Visit a parse tree produced by mdxParser#member_property_definition.
	VisitMember_property_definition(ctx *Member_property_definitionContext) interface{}

	// Visit a parse tree produced by mdxParser#set_name.
	VisitSet_name(ctx *Set_nameContext) interface{}

	// Visit a parse tree produced by mdxParser#compound_id.
	VisitCompound_id(ctx *Compound_idContext) interface{}

	// Visit a parse tree produced by mdxParser#axis_specification.
	VisitAxis_specification(ctx *Axis_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#axis_name.
	VisitAxis_name(ctx *Axis_nameContext) interface{}

	// Visit a parse tree produced by mdxParser#dim_props.
	VisitDim_props(ctx *Dim_propsContext) interface{}

	// Visit a parse tree produced by mdxParser#property_list.
	VisitProperty_list(ctx *Property_listContext) interface{}

	// Visit a parse tree produced by mdxParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by mdxParser#cube_specification.
	VisitCube_specification(ctx *Cube_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#cube_name.
	VisitCube_name(ctx *Cube_nameContext) interface{}

	// Visit a parse tree produced by mdxParser#slicer_specification.
	VisitSlicer_specification(ctx *Slicer_specificationContext) interface{}

	// Visit a parse tree produced by mdxParser#cell_props.
	VisitCell_props(ctx *Cell_propsContext) interface{}

	// Visit a parse tree produced by mdxParser#cell_property_list.
	VisitCell_property_list(ctx *Cell_property_listContext) interface{}

	// Visit a parse tree produced by mdxParser#cell_property.
	VisitCell_property(ctx *Cell_propertyContext) interface{}

	// Visit a parse tree produced by mdxParser#mandatory_cell_property.
	VisitMandatory_cell_property(ctx *Mandatory_cell_propertyContext) interface{}

	// Visit a parse tree produced by mdxParser#provider_specific_cell_property.
	VisitProvider_specific_cell_property(ctx *Provider_specific_cell_propertyContext) interface{}

	// Visit a parse tree produced by mdxParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by mdxParser#value_expression.
	VisitValue_expression(ctx *Value_expressionContext) interface{}

	// Visit a parse tree produced by mdxParser#value_xor_expression.
	VisitValue_xor_expression(ctx *Value_xor_expressionContext) interface{}

	// Visit a parse tree produced by mdxParser#value_or_expression.
	VisitValue_or_expression(ctx *Value_or_expressionContext) interface{}

	// Visit a parse tree produced by mdxParser#term5.
	VisitTerm5(ctx *Term5Context) interface{}

	// Visit a parse tree produced by mdxParser#term4.
	VisitTerm4(ctx *Term4Context) interface{}

	// Visit a parse tree produced by mdxParser#term3.
	VisitTerm3(ctx *Term3Context) interface{}

	// Visit a parse tree produced by mdxParser#term2.
	VisitTerm2(ctx *Term2Context) interface{}

	// Visit a parse tree produced by mdxParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by mdxParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by mdxParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by mdxParser#value_expression_primary.
	VisitValue_expression_primary(ctx *Value_expression_primaryContext) interface{}

	// Visit a parse tree produced by mdxParser#value_expression_primary0.
	VisitValue_expression_primary0(ctx *Value_expression_primary0Context) interface{}

	// Visit a parse tree produced by mdxParser#exp_list.
	VisitExp_list(ctx *Exp_listContext) interface{}

	// Visit a parse tree produced by mdxParser#case_expression.
	VisitCase_expression(ctx *Case_expressionContext) interface{}

	// Visit a parse tree produced by mdxParser#when_list.
	VisitWhen_list(ctx *When_listContext) interface{}

	// Visit a parse tree produced by mdxParser#when_clause.
	VisitWhen_clause(ctx *When_clauseContext) interface{}

	// Visit a parse tree produced by mdxParser#comp_op.
	VisitComp_op(ctx *Comp_opContext) interface{}

	// Visit a parse tree produced by mdxParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by mdxParser#unquoted_identifier.
	VisitUnquoted_identifier(ctx *Unquoted_identifierContext) interface{}

	// Visit a parse tree produced by mdxParser#amp_quoted_identifier.
	VisitAmp_quoted_identifier(ctx *Amp_quoted_identifierContext) interface{}

	// Visit a parse tree produced by mdxParser#quoted_identifier.
	VisitQuoted_identifier(ctx *Quoted_identifierContext) interface{}

	// Visit a parse tree produced by mdxParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}
}
