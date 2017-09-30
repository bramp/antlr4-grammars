// Generated from mdx.g4 by ANTLR 4.7.

package mdx // mdx
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasemdxVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemdxVisitor) VisitMdx_statement(ctx *Mdx_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitFormula_specification(ctx *Formula_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitSingle_formula_specification(ctx *Single_formula_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitSet_specification(ctx *Set_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitMember_specification(ctx *Member_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitAxis_specification_list(ctx *Axis_specification_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitMember_property_def_list(ctx *Member_property_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitMember_name(ctx *Member_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitMember_property_definition(ctx *Member_property_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitSet_name(ctx *Set_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCompound_id(ctx *Compound_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitAxis_specification(ctx *Axis_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitAxis_name(ctx *Axis_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitDim_props(ctx *Dim_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitProperty_list(ctx *Property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCube_specification(ctx *Cube_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCube_name(ctx *Cube_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitSlicer_specification(ctx *Slicer_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCell_props(ctx *Cell_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCell_property_list(ctx *Cell_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCell_property(ctx *Cell_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitMandatory_cell_property(ctx *Mandatory_cell_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitProvider_specific_cell_property(ctx *Provider_specific_cell_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitValue_expression(ctx *Value_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitValue_xor_expression(ctx *Value_xor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitValue_or_expression(ctx *Value_or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitTerm5(ctx *Term5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitTerm4(ctx *Term4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitTerm3(ctx *Term3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitTerm2(ctx *Term2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitValue_expression_primary(ctx *Value_expression_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitValue_expression_primary0(ctx *Value_expression_primary0Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitExp_list(ctx *Exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitCase_expression(ctx *Case_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitWhen_list(ctx *When_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitWhen_clause(ctx *When_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitComp_op(ctx *Comp_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitUnquoted_identifier(ctx *Unquoted_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitAmp_quoted_identifier(ctx *Amp_quoted_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitQuoted_identifier(ctx *Quoted_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemdxVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
