// Generated from mdx.g4 by ANTLR 4.7.

package mdx // mdx
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemdxListener is a complete listener for a parse tree produced by mdxParser.
type BasemdxListener struct{}

var _ mdxListener = &BasemdxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemdxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemdxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemdxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemdxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMdx_statement is called when production mdx_statement is entered.
func (s *BasemdxListener) EnterMdx_statement(ctx *Mdx_statementContext) {}

// ExitMdx_statement is called when production mdx_statement is exited.
func (s *BasemdxListener) ExitMdx_statement(ctx *Mdx_statementContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BasemdxListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BasemdxListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterFormula_specification is called when production formula_specification is entered.
func (s *BasemdxListener) EnterFormula_specification(ctx *Formula_specificationContext) {}

// ExitFormula_specification is called when production formula_specification is exited.
func (s *BasemdxListener) ExitFormula_specification(ctx *Formula_specificationContext) {}

// EnterSingle_formula_specification is called when production single_formula_specification is entered.
func (s *BasemdxListener) EnterSingle_formula_specification(ctx *Single_formula_specificationContext) {
}

// ExitSingle_formula_specification is called when production single_formula_specification is exited.
func (s *BasemdxListener) ExitSingle_formula_specification(ctx *Single_formula_specificationContext) {}

// EnterSet_specification is called when production set_specification is entered.
func (s *BasemdxListener) EnterSet_specification(ctx *Set_specificationContext) {}

// ExitSet_specification is called when production set_specification is exited.
func (s *BasemdxListener) ExitSet_specification(ctx *Set_specificationContext) {}

// EnterMember_specification is called when production member_specification is entered.
func (s *BasemdxListener) EnterMember_specification(ctx *Member_specificationContext) {}

// ExitMember_specification is called when production member_specification is exited.
func (s *BasemdxListener) ExitMember_specification(ctx *Member_specificationContext) {}

// EnterAxis_specification_list is called when production axis_specification_list is entered.
func (s *BasemdxListener) EnterAxis_specification_list(ctx *Axis_specification_listContext) {}

// ExitAxis_specification_list is called when production axis_specification_list is exited.
func (s *BasemdxListener) ExitAxis_specification_list(ctx *Axis_specification_listContext) {}

// EnterMember_property_def_list is called when production member_property_def_list is entered.
func (s *BasemdxListener) EnterMember_property_def_list(ctx *Member_property_def_listContext) {}

// ExitMember_property_def_list is called when production member_property_def_list is exited.
func (s *BasemdxListener) ExitMember_property_def_list(ctx *Member_property_def_listContext) {}

// EnterMember_name is called when production member_name is entered.
func (s *BasemdxListener) EnterMember_name(ctx *Member_nameContext) {}

// ExitMember_name is called when production member_name is exited.
func (s *BasemdxListener) ExitMember_name(ctx *Member_nameContext) {}

// EnterMember_property_definition is called when production member_property_definition is entered.
func (s *BasemdxListener) EnterMember_property_definition(ctx *Member_property_definitionContext) {}

// ExitMember_property_definition is called when production member_property_definition is exited.
func (s *BasemdxListener) ExitMember_property_definition(ctx *Member_property_definitionContext) {}

// EnterSet_name is called when production set_name is entered.
func (s *BasemdxListener) EnterSet_name(ctx *Set_nameContext) {}

// ExitSet_name is called when production set_name is exited.
func (s *BasemdxListener) ExitSet_name(ctx *Set_nameContext) {}

// EnterCompound_id is called when production compound_id is entered.
func (s *BasemdxListener) EnterCompound_id(ctx *Compound_idContext) {}

// ExitCompound_id is called when production compound_id is exited.
func (s *BasemdxListener) ExitCompound_id(ctx *Compound_idContext) {}

// EnterAxis_specification is called when production axis_specification is entered.
func (s *BasemdxListener) EnterAxis_specification(ctx *Axis_specificationContext) {}

// ExitAxis_specification is called when production axis_specification is exited.
func (s *BasemdxListener) ExitAxis_specification(ctx *Axis_specificationContext) {}

// EnterAxis_name is called when production axis_name is entered.
func (s *BasemdxListener) EnterAxis_name(ctx *Axis_nameContext) {}

// ExitAxis_name is called when production axis_name is exited.
func (s *BasemdxListener) ExitAxis_name(ctx *Axis_nameContext) {}

// EnterDim_props is called when production dim_props is entered.
func (s *BasemdxListener) EnterDim_props(ctx *Dim_propsContext) {}

// ExitDim_props is called when production dim_props is exited.
func (s *BasemdxListener) ExitDim_props(ctx *Dim_propsContext) {}

// EnterProperty_list is called when production property_list is entered.
func (s *BasemdxListener) EnterProperty_list(ctx *Property_listContext) {}

// ExitProperty_list is called when production property_list is exited.
func (s *BasemdxListener) ExitProperty_list(ctx *Property_listContext) {}

// EnterProperty is called when production property is entered.
func (s *BasemdxListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BasemdxListener) ExitProperty(ctx *PropertyContext) {}

// EnterCube_specification is called when production cube_specification is entered.
func (s *BasemdxListener) EnterCube_specification(ctx *Cube_specificationContext) {}

// ExitCube_specification is called when production cube_specification is exited.
func (s *BasemdxListener) ExitCube_specification(ctx *Cube_specificationContext) {}

// EnterCube_name is called when production cube_name is entered.
func (s *BasemdxListener) EnterCube_name(ctx *Cube_nameContext) {}

// ExitCube_name is called when production cube_name is exited.
func (s *BasemdxListener) ExitCube_name(ctx *Cube_nameContext) {}

// EnterSlicer_specification is called when production slicer_specification is entered.
func (s *BasemdxListener) EnterSlicer_specification(ctx *Slicer_specificationContext) {}

// ExitSlicer_specification is called when production slicer_specification is exited.
func (s *BasemdxListener) ExitSlicer_specification(ctx *Slicer_specificationContext) {}

// EnterCell_props is called when production cell_props is entered.
func (s *BasemdxListener) EnterCell_props(ctx *Cell_propsContext) {}

// ExitCell_props is called when production cell_props is exited.
func (s *BasemdxListener) ExitCell_props(ctx *Cell_propsContext) {}

// EnterCell_property_list is called when production cell_property_list is entered.
func (s *BasemdxListener) EnterCell_property_list(ctx *Cell_property_listContext) {}

// ExitCell_property_list is called when production cell_property_list is exited.
func (s *BasemdxListener) ExitCell_property_list(ctx *Cell_property_listContext) {}

// EnterCell_property is called when production cell_property is entered.
func (s *BasemdxListener) EnterCell_property(ctx *Cell_propertyContext) {}

// ExitCell_property is called when production cell_property is exited.
func (s *BasemdxListener) ExitCell_property(ctx *Cell_propertyContext) {}

// EnterMandatory_cell_property is called when production mandatory_cell_property is entered.
func (s *BasemdxListener) EnterMandatory_cell_property(ctx *Mandatory_cell_propertyContext) {}

// ExitMandatory_cell_property is called when production mandatory_cell_property is exited.
func (s *BasemdxListener) ExitMandatory_cell_property(ctx *Mandatory_cell_propertyContext) {}

// EnterProvider_specific_cell_property is called when production provider_specific_cell_property is entered.
func (s *BasemdxListener) EnterProvider_specific_cell_property(ctx *Provider_specific_cell_propertyContext) {
}

// ExitProvider_specific_cell_property is called when production provider_specific_cell_property is exited.
func (s *BasemdxListener) ExitProvider_specific_cell_property(ctx *Provider_specific_cell_propertyContext) {
}

// EnterExpression is called when production expression is entered.
func (s *BasemdxListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasemdxListener) ExitExpression(ctx *ExpressionContext) {}

// EnterValue_expression is called when production value_expression is entered.
func (s *BasemdxListener) EnterValue_expression(ctx *Value_expressionContext) {}

// ExitValue_expression is called when production value_expression is exited.
func (s *BasemdxListener) ExitValue_expression(ctx *Value_expressionContext) {}

// EnterValue_xor_expression is called when production value_xor_expression is entered.
func (s *BasemdxListener) EnterValue_xor_expression(ctx *Value_xor_expressionContext) {}

// ExitValue_xor_expression is called when production value_xor_expression is exited.
func (s *BasemdxListener) ExitValue_xor_expression(ctx *Value_xor_expressionContext) {}

// EnterValue_or_expression is called when production value_or_expression is entered.
func (s *BasemdxListener) EnterValue_or_expression(ctx *Value_or_expressionContext) {}

// ExitValue_or_expression is called when production value_or_expression is exited.
func (s *BasemdxListener) ExitValue_or_expression(ctx *Value_or_expressionContext) {}

// EnterTerm5 is called when production term5 is entered.
func (s *BasemdxListener) EnterTerm5(ctx *Term5Context) {}

// ExitTerm5 is called when production term5 is exited.
func (s *BasemdxListener) ExitTerm5(ctx *Term5Context) {}

// EnterTerm4 is called when production term4 is entered.
func (s *BasemdxListener) EnterTerm4(ctx *Term4Context) {}

// ExitTerm4 is called when production term4 is exited.
func (s *BasemdxListener) ExitTerm4(ctx *Term4Context) {}

// EnterTerm3 is called when production term3 is entered.
func (s *BasemdxListener) EnterTerm3(ctx *Term3Context) {}

// ExitTerm3 is called when production term3 is exited.
func (s *BasemdxListener) ExitTerm3(ctx *Term3Context) {}

// EnterTerm2 is called when production term2 is entered.
func (s *BasemdxListener) EnterTerm2(ctx *Term2Context) {}

// ExitTerm2 is called when production term2 is exited.
func (s *BasemdxListener) ExitTerm2(ctx *Term2Context) {}

// EnterTerm is called when production term is entered.
func (s *BasemdxListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasemdxListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasemdxListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasemdxListener) ExitFactor(ctx *FactorContext) {}

// EnterFunction is called when production function is entered.
func (s *BasemdxListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasemdxListener) ExitFunction(ctx *FunctionContext) {}

// EnterValue_expression_primary is called when production value_expression_primary is entered.
func (s *BasemdxListener) EnterValue_expression_primary(ctx *Value_expression_primaryContext) {}

// ExitValue_expression_primary is called when production value_expression_primary is exited.
func (s *BasemdxListener) ExitValue_expression_primary(ctx *Value_expression_primaryContext) {}

// EnterValue_expression_primary0 is called when production value_expression_primary0 is entered.
func (s *BasemdxListener) EnterValue_expression_primary0(ctx *Value_expression_primary0Context) {}

// ExitValue_expression_primary0 is called when production value_expression_primary0 is exited.
func (s *BasemdxListener) ExitValue_expression_primary0(ctx *Value_expression_primary0Context) {}

// EnterExp_list is called when production exp_list is entered.
func (s *BasemdxListener) EnterExp_list(ctx *Exp_listContext) {}

// ExitExp_list is called when production exp_list is exited.
func (s *BasemdxListener) ExitExp_list(ctx *Exp_listContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BasemdxListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BasemdxListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterWhen_list is called when production when_list is entered.
func (s *BasemdxListener) EnterWhen_list(ctx *When_listContext) {}

// ExitWhen_list is called when production when_list is exited.
func (s *BasemdxListener) ExitWhen_list(ctx *When_listContext) {}

// EnterWhen_clause is called when production when_clause is entered.
func (s *BasemdxListener) EnterWhen_clause(ctx *When_clauseContext) {}

// ExitWhen_clause is called when production when_clause is exited.
func (s *BasemdxListener) ExitWhen_clause(ctx *When_clauseContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BasemdxListener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BasemdxListener) ExitComp_op(ctx *Comp_opContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasemdxListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasemdxListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterUnquoted_identifier is called when production unquoted_identifier is entered.
func (s *BasemdxListener) EnterUnquoted_identifier(ctx *Unquoted_identifierContext) {}

// ExitUnquoted_identifier is called when production unquoted_identifier is exited.
func (s *BasemdxListener) ExitUnquoted_identifier(ctx *Unquoted_identifierContext) {}

// EnterAmp_quoted_identifier is called when production amp_quoted_identifier is entered.
func (s *BasemdxListener) EnterAmp_quoted_identifier(ctx *Amp_quoted_identifierContext) {}

// ExitAmp_quoted_identifier is called when production amp_quoted_identifier is exited.
func (s *BasemdxListener) ExitAmp_quoted_identifier(ctx *Amp_quoted_identifierContext) {}

// EnterQuoted_identifier is called when production quoted_identifier is entered.
func (s *BasemdxListener) EnterQuoted_identifier(ctx *Quoted_identifierContext) {}

// ExitQuoted_identifier is called when production quoted_identifier is exited.
func (s *BasemdxListener) ExitQuoted_identifier(ctx *Quoted_identifierContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BasemdxListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BasemdxListener) ExitKeyword(ctx *KeywordContext) {}
