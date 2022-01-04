// Code generated from WavefrontOBJ.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wavefrontobj // WavefrontOBJ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWavefrontOBJListener is a complete listener for a parse tree produced by WavefrontOBJParser.
type BaseWavefrontOBJListener struct{}

var _ WavefrontOBJListener = &BaseWavefrontOBJListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWavefrontOBJListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWavefrontOBJListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWavefrontOBJListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWavefrontOBJListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseWavefrontOBJListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseWavefrontOBJListener) ExitStart(ctx *StartContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseWavefrontOBJListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseWavefrontOBJListener) ExitStatement(ctx *StatementContext) {}

// EnterCall is called when production call is entered.
func (s *BaseWavefrontOBJListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseWavefrontOBJListener) ExitCall(ctx *CallContext) {}

// EnterCsh is called when production csh is entered.
func (s *BaseWavefrontOBJListener) EnterCsh(ctx *CshContext) {}

// ExitCsh is called when production csh is exited.
func (s *BaseWavefrontOBJListener) ExitCsh(ctx *CshContext) {}

// EnterVertex is called when production vertex is entered.
func (s *BaseWavefrontOBJListener) EnterVertex(ctx *VertexContext) {}

// ExitVertex is called when production vertex is exited.
func (s *BaseWavefrontOBJListener) ExitVertex(ctx *VertexContext) {}

// EnterVertex_parameter is called when production vertex_parameter is entered.
func (s *BaseWavefrontOBJListener) EnterVertex_parameter(ctx *Vertex_parameterContext) {}

// ExitVertex_parameter is called when production vertex_parameter is exited.
func (s *BaseWavefrontOBJListener) ExitVertex_parameter(ctx *Vertex_parameterContext) {}

// EnterVertex_normal is called when production vertex_normal is entered.
func (s *BaseWavefrontOBJListener) EnterVertex_normal(ctx *Vertex_normalContext) {}

// ExitVertex_normal is called when production vertex_normal is exited.
func (s *BaseWavefrontOBJListener) ExitVertex_normal(ctx *Vertex_normalContext) {}

// EnterVertex_texture is called when production vertex_texture is entered.
func (s *BaseWavefrontOBJListener) EnterVertex_texture(ctx *Vertex_textureContext) {}

// ExitVertex_texture is called when production vertex_texture is exited.
func (s *BaseWavefrontOBJListener) ExitVertex_texture(ctx *Vertex_textureContext) {}

// EnterCurve_surface_type is called when production curve_surface_type is entered.
func (s *BaseWavefrontOBJListener) EnterCurve_surface_type(ctx *Curve_surface_typeContext) {}

// ExitCurve_surface_type is called when production curve_surface_type is exited.
func (s *BaseWavefrontOBJListener) ExitCurve_surface_type(ctx *Curve_surface_typeContext) {}

// EnterDegree is called when production degree is entered.
func (s *BaseWavefrontOBJListener) EnterDegree(ctx *DegreeContext) {}

// ExitDegree is called when production degree is exited.
func (s *BaseWavefrontOBJListener) ExitDegree(ctx *DegreeContext) {}

// EnterBasis_matrix is called when production basis_matrix is entered.
func (s *BaseWavefrontOBJListener) EnterBasis_matrix(ctx *Basis_matrixContext) {}

// ExitBasis_matrix is called when production basis_matrix is exited.
func (s *BaseWavefrontOBJListener) ExitBasis_matrix(ctx *Basis_matrixContext) {}

// EnterStep is called when production step is entered.
func (s *BaseWavefrontOBJListener) EnterStep(ctx *StepContext) {}

// ExitStep is called when production step is exited.
func (s *BaseWavefrontOBJListener) ExitStep(ctx *StepContext) {}

// EnterPoints is called when production points is entered.
func (s *BaseWavefrontOBJListener) EnterPoints(ctx *PointsContext) {}

// ExitPoints is called when production points is exited.
func (s *BaseWavefrontOBJListener) ExitPoints(ctx *PointsContext) {}

// EnterLines is called when production lines is entered.
func (s *BaseWavefrontOBJListener) EnterLines(ctx *LinesContext) {}

// ExitLines is called when production lines is exited.
func (s *BaseWavefrontOBJListener) ExitLines(ctx *LinesContext) {}

// EnterFaces is called when production faces is entered.
func (s *BaseWavefrontOBJListener) EnterFaces(ctx *FacesContext) {}

// ExitFaces is called when production faces is exited.
func (s *BaseWavefrontOBJListener) ExitFaces(ctx *FacesContext) {}

// EnterFree_form_surface is called when production free_form_surface is entered.
func (s *BaseWavefrontOBJListener) EnterFree_form_surface(ctx *Free_form_surfaceContext) {}

// ExitFree_form_surface is called when production free_form_surface is exited.
func (s *BaseWavefrontOBJListener) ExitFree_form_surface(ctx *Free_form_surfaceContext) {}

// EnterCurve is called when production curve is entered.
func (s *BaseWavefrontOBJListener) EnterCurve(ctx *CurveContext) {}

// ExitCurve is called when production curve is exited.
func (s *BaseWavefrontOBJListener) ExitCurve(ctx *CurveContext) {}

// EnterCurve2d is called when production curve2d is entered.
func (s *BaseWavefrontOBJListener) EnterCurve2d(ctx *Curve2dContext) {}

// ExitCurve2d is called when production curve2d is exited.
func (s *BaseWavefrontOBJListener) ExitCurve2d(ctx *Curve2dContext) {}

// EnterSurface is called when production surface is entered.
func (s *BaseWavefrontOBJListener) EnterSurface(ctx *SurfaceContext) {}

// ExitSurface is called when production surface is exited.
func (s *BaseWavefrontOBJListener) ExitSurface(ctx *SurfaceContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseWavefrontOBJListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseWavefrontOBJListener) ExitParameter(ctx *ParameterContext) {}

// EnterOuter_trimming_loop is called when production outer_trimming_loop is entered.
func (s *BaseWavefrontOBJListener) EnterOuter_trimming_loop(ctx *Outer_trimming_loopContext) {}

// ExitOuter_trimming_loop is called when production outer_trimming_loop is exited.
func (s *BaseWavefrontOBJListener) ExitOuter_trimming_loop(ctx *Outer_trimming_loopContext) {}

// EnterInner_trimming_loop is called when production inner_trimming_loop is entered.
func (s *BaseWavefrontOBJListener) EnterInner_trimming_loop(ctx *Inner_trimming_loopContext) {}

// ExitInner_trimming_loop is called when production inner_trimming_loop is exited.
func (s *BaseWavefrontOBJListener) ExitInner_trimming_loop(ctx *Inner_trimming_loopContext) {}

// EnterSpecial_curve is called when production special_curve is entered.
func (s *BaseWavefrontOBJListener) EnterSpecial_curve(ctx *Special_curveContext) {}

// ExitSpecial_curve is called when production special_curve is exited.
func (s *BaseWavefrontOBJListener) ExitSpecial_curve(ctx *Special_curveContext) {}

// EnterSpecial_point is called when production special_point is entered.
func (s *BaseWavefrontOBJListener) EnterSpecial_point(ctx *Special_pointContext) {}

// ExitSpecial_point is called when production special_point is exited.
func (s *BaseWavefrontOBJListener) ExitSpecial_point(ctx *Special_pointContext) {}

// EnterEnd is called when production end is entered.
func (s *BaseWavefrontOBJListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *BaseWavefrontOBJListener) ExitEnd(ctx *EndContext) {}

// EnterConnectivity is called when production connectivity is entered.
func (s *BaseWavefrontOBJListener) EnterConnectivity(ctx *ConnectivityContext) {}

// ExitConnectivity is called when production connectivity is exited.
func (s *BaseWavefrontOBJListener) ExitConnectivity(ctx *ConnectivityContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseWavefrontOBJListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseWavefrontOBJListener) ExitGroup(ctx *GroupContext) {}

// EnterSmoothing_group is called when production smoothing_group is entered.
func (s *BaseWavefrontOBJListener) EnterSmoothing_group(ctx *Smoothing_groupContext) {}

// ExitSmoothing_group is called when production smoothing_group is exited.
func (s *BaseWavefrontOBJListener) ExitSmoothing_group(ctx *Smoothing_groupContext) {}

// EnterMerging_group is called when production merging_group is entered.
func (s *BaseWavefrontOBJListener) EnterMerging_group(ctx *Merging_groupContext) {}

// ExitMerging_group is called when production merging_group is exited.
func (s *BaseWavefrontOBJListener) ExitMerging_group(ctx *Merging_groupContext) {}

// EnterObject_name is called when production object_name is entered.
func (s *BaseWavefrontOBJListener) EnterObject_name(ctx *Object_nameContext) {}

// ExitObject_name is called when production object_name is exited.
func (s *BaseWavefrontOBJListener) ExitObject_name(ctx *Object_nameContext) {}

// EnterBevel is called when production bevel is entered.
func (s *BaseWavefrontOBJListener) EnterBevel(ctx *BevelContext) {}

// ExitBevel is called when production bevel is exited.
func (s *BaseWavefrontOBJListener) ExitBevel(ctx *BevelContext) {}

// EnterColor_interpolation is called when production color_interpolation is entered.
func (s *BaseWavefrontOBJListener) EnterColor_interpolation(ctx *Color_interpolationContext) {}

// ExitColor_interpolation is called when production color_interpolation is exited.
func (s *BaseWavefrontOBJListener) ExitColor_interpolation(ctx *Color_interpolationContext) {}

// EnterDissolve_interpolation is called when production dissolve_interpolation is entered.
func (s *BaseWavefrontOBJListener) EnterDissolve_interpolation(ctx *Dissolve_interpolationContext) {}

// ExitDissolve_interpolation is called when production dissolve_interpolation is exited.
func (s *BaseWavefrontOBJListener) ExitDissolve_interpolation(ctx *Dissolve_interpolationContext) {}

// EnterLevel_of_detail is called when production level_of_detail is entered.
func (s *BaseWavefrontOBJListener) EnterLevel_of_detail(ctx *Level_of_detailContext) {}

// ExitLevel_of_detail is called when production level_of_detail is exited.
func (s *BaseWavefrontOBJListener) ExitLevel_of_detail(ctx *Level_of_detailContext) {}

// EnterMap_library is called when production map_library is entered.
func (s *BaseWavefrontOBJListener) EnterMap_library(ctx *Map_libraryContext) {}

// ExitMap_library is called when production map_library is exited.
func (s *BaseWavefrontOBJListener) ExitMap_library(ctx *Map_libraryContext) {}

// EnterUse_map is called when production use_map is entered.
func (s *BaseWavefrontOBJListener) EnterUse_map(ctx *Use_mapContext) {}

// ExitUse_map is called when production use_map is exited.
func (s *BaseWavefrontOBJListener) ExitUse_map(ctx *Use_mapContext) {}

// EnterMaterial_library is called when production material_library is entered.
func (s *BaseWavefrontOBJListener) EnterMaterial_library(ctx *Material_libraryContext) {}

// ExitMaterial_library is called when production material_library is exited.
func (s *BaseWavefrontOBJListener) ExitMaterial_library(ctx *Material_libraryContext) {}

// EnterUse_material is called when production use_material is entered.
func (s *BaseWavefrontOBJListener) EnterUse_material(ctx *Use_materialContext) {}

// ExitUse_material is called when production use_material is exited.
func (s *BaseWavefrontOBJListener) ExitUse_material(ctx *Use_materialContext) {}

// EnterShadow_object is called when production shadow_object is entered.
func (s *BaseWavefrontOBJListener) EnterShadow_object(ctx *Shadow_objectContext) {}

// ExitShadow_object is called when production shadow_object is exited.
func (s *BaseWavefrontOBJListener) ExitShadow_object(ctx *Shadow_objectContext) {}

// EnterTrace_object is called when production trace_object is entered.
func (s *BaseWavefrontOBJListener) EnterTrace_object(ctx *Trace_objectContext) {}

// ExitTrace_object is called when production trace_object is exited.
func (s *BaseWavefrontOBJListener) ExitTrace_object(ctx *Trace_objectContext) {}

// EnterCurve_approximation_technique is called when production curve_approximation_technique is entered.
func (s *BaseWavefrontOBJListener) EnterCurve_approximation_technique(ctx *Curve_approximation_techniqueContext) {
}

// ExitCurve_approximation_technique is called when production curve_approximation_technique is exited.
func (s *BaseWavefrontOBJListener) ExitCurve_approximation_technique(ctx *Curve_approximation_techniqueContext) {
}

// EnterSurface_approximation_technique is called when production surface_approximation_technique is entered.
func (s *BaseWavefrontOBJListener) EnterSurface_approximation_technique(ctx *Surface_approximation_techniqueContext) {
}

// ExitSurface_approximation_technique is called when production surface_approximation_technique is exited.
func (s *BaseWavefrontOBJListener) ExitSurface_approximation_technique(ctx *Surface_approximation_techniqueContext) {
}

// EnterDecimal is called when production decimal is entered.
func (s *BaseWavefrontOBJListener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production decimal is exited.
func (s *BaseWavefrontOBJListener) ExitDecimal(ctx *DecimalContext) {}
