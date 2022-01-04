// Code generated from WavefrontOBJ.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wavefrontobj // WavefrontOBJ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// WavefrontOBJListener is a complete listener for a parse tree produced by WavefrontOBJParser.
type WavefrontOBJListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterCsh is called when entering the csh production.
	EnterCsh(c *CshContext)

	// EnterVertex is called when entering the vertex production.
	EnterVertex(c *VertexContext)

	// EnterVertex_parameter is called when entering the vertex_parameter production.
	EnterVertex_parameter(c *Vertex_parameterContext)

	// EnterVertex_normal is called when entering the vertex_normal production.
	EnterVertex_normal(c *Vertex_normalContext)

	// EnterVertex_texture is called when entering the vertex_texture production.
	EnterVertex_texture(c *Vertex_textureContext)

	// EnterCurve_surface_type is called when entering the curve_surface_type production.
	EnterCurve_surface_type(c *Curve_surface_typeContext)

	// EnterDegree is called when entering the degree production.
	EnterDegree(c *DegreeContext)

	// EnterBasis_matrix is called when entering the basis_matrix production.
	EnterBasis_matrix(c *Basis_matrixContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterPoints is called when entering the points production.
	EnterPoints(c *PointsContext)

	// EnterLines is called when entering the lines production.
	EnterLines(c *LinesContext)

	// EnterFaces is called when entering the faces production.
	EnterFaces(c *FacesContext)

	// EnterFree_form_surface is called when entering the free_form_surface production.
	EnterFree_form_surface(c *Free_form_surfaceContext)

	// EnterCurve is called when entering the curve production.
	EnterCurve(c *CurveContext)

	// EnterCurve2d is called when entering the curve2d production.
	EnterCurve2d(c *Curve2dContext)

	// EnterSurface is called when entering the surface production.
	EnterSurface(c *SurfaceContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterOuter_trimming_loop is called when entering the outer_trimming_loop production.
	EnterOuter_trimming_loop(c *Outer_trimming_loopContext)

	// EnterInner_trimming_loop is called when entering the inner_trimming_loop production.
	EnterInner_trimming_loop(c *Inner_trimming_loopContext)

	// EnterSpecial_curve is called when entering the special_curve production.
	EnterSpecial_curve(c *Special_curveContext)

	// EnterSpecial_point is called when entering the special_point production.
	EnterSpecial_point(c *Special_pointContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterConnectivity is called when entering the connectivity production.
	EnterConnectivity(c *ConnectivityContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterSmoothing_group is called when entering the smoothing_group production.
	EnterSmoothing_group(c *Smoothing_groupContext)

	// EnterMerging_group is called when entering the merging_group production.
	EnterMerging_group(c *Merging_groupContext)

	// EnterObject_name is called when entering the object_name production.
	EnterObject_name(c *Object_nameContext)

	// EnterBevel is called when entering the bevel production.
	EnterBevel(c *BevelContext)

	// EnterColor_interpolation is called when entering the color_interpolation production.
	EnterColor_interpolation(c *Color_interpolationContext)

	// EnterDissolve_interpolation is called when entering the dissolve_interpolation production.
	EnterDissolve_interpolation(c *Dissolve_interpolationContext)

	// EnterLevel_of_detail is called when entering the level_of_detail production.
	EnterLevel_of_detail(c *Level_of_detailContext)

	// EnterMap_library is called when entering the map_library production.
	EnterMap_library(c *Map_libraryContext)

	// EnterUse_map is called when entering the use_map production.
	EnterUse_map(c *Use_mapContext)

	// EnterMaterial_library is called when entering the material_library production.
	EnterMaterial_library(c *Material_libraryContext)

	// EnterUse_material is called when entering the use_material production.
	EnterUse_material(c *Use_materialContext)

	// EnterShadow_object is called when entering the shadow_object production.
	EnterShadow_object(c *Shadow_objectContext)

	// EnterTrace_object is called when entering the trace_object production.
	EnterTrace_object(c *Trace_objectContext)

	// EnterCurve_approximation_technique is called when entering the curve_approximation_technique production.
	EnterCurve_approximation_technique(c *Curve_approximation_techniqueContext)

	// EnterSurface_approximation_technique is called when entering the surface_approximation_technique production.
	EnterSurface_approximation_technique(c *Surface_approximation_techniqueContext)

	// EnterDecimal is called when entering the decimal production.
	EnterDecimal(c *DecimalContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitCsh is called when exiting the csh production.
	ExitCsh(c *CshContext)

	// ExitVertex is called when exiting the vertex production.
	ExitVertex(c *VertexContext)

	// ExitVertex_parameter is called when exiting the vertex_parameter production.
	ExitVertex_parameter(c *Vertex_parameterContext)

	// ExitVertex_normal is called when exiting the vertex_normal production.
	ExitVertex_normal(c *Vertex_normalContext)

	// ExitVertex_texture is called when exiting the vertex_texture production.
	ExitVertex_texture(c *Vertex_textureContext)

	// ExitCurve_surface_type is called when exiting the curve_surface_type production.
	ExitCurve_surface_type(c *Curve_surface_typeContext)

	// ExitDegree is called when exiting the degree production.
	ExitDegree(c *DegreeContext)

	// ExitBasis_matrix is called when exiting the basis_matrix production.
	ExitBasis_matrix(c *Basis_matrixContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitPoints is called when exiting the points production.
	ExitPoints(c *PointsContext)

	// ExitLines is called when exiting the lines production.
	ExitLines(c *LinesContext)

	// ExitFaces is called when exiting the faces production.
	ExitFaces(c *FacesContext)

	// ExitFree_form_surface is called when exiting the free_form_surface production.
	ExitFree_form_surface(c *Free_form_surfaceContext)

	// ExitCurve is called when exiting the curve production.
	ExitCurve(c *CurveContext)

	// ExitCurve2d is called when exiting the curve2d production.
	ExitCurve2d(c *Curve2dContext)

	// ExitSurface is called when exiting the surface production.
	ExitSurface(c *SurfaceContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitOuter_trimming_loop is called when exiting the outer_trimming_loop production.
	ExitOuter_trimming_loop(c *Outer_trimming_loopContext)

	// ExitInner_trimming_loop is called when exiting the inner_trimming_loop production.
	ExitInner_trimming_loop(c *Inner_trimming_loopContext)

	// ExitSpecial_curve is called when exiting the special_curve production.
	ExitSpecial_curve(c *Special_curveContext)

	// ExitSpecial_point is called when exiting the special_point production.
	ExitSpecial_point(c *Special_pointContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitConnectivity is called when exiting the connectivity production.
	ExitConnectivity(c *ConnectivityContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitSmoothing_group is called when exiting the smoothing_group production.
	ExitSmoothing_group(c *Smoothing_groupContext)

	// ExitMerging_group is called when exiting the merging_group production.
	ExitMerging_group(c *Merging_groupContext)

	// ExitObject_name is called when exiting the object_name production.
	ExitObject_name(c *Object_nameContext)

	// ExitBevel is called when exiting the bevel production.
	ExitBevel(c *BevelContext)

	// ExitColor_interpolation is called when exiting the color_interpolation production.
	ExitColor_interpolation(c *Color_interpolationContext)

	// ExitDissolve_interpolation is called when exiting the dissolve_interpolation production.
	ExitDissolve_interpolation(c *Dissolve_interpolationContext)

	// ExitLevel_of_detail is called when exiting the level_of_detail production.
	ExitLevel_of_detail(c *Level_of_detailContext)

	// ExitMap_library is called when exiting the map_library production.
	ExitMap_library(c *Map_libraryContext)

	// ExitUse_map is called when exiting the use_map production.
	ExitUse_map(c *Use_mapContext)

	// ExitMaterial_library is called when exiting the material_library production.
	ExitMaterial_library(c *Material_libraryContext)

	// ExitUse_material is called when exiting the use_material production.
	ExitUse_material(c *Use_materialContext)

	// ExitShadow_object is called when exiting the shadow_object production.
	ExitShadow_object(c *Shadow_objectContext)

	// ExitTrace_object is called when exiting the trace_object production.
	ExitTrace_object(c *Trace_objectContext)

	// ExitCurve_approximation_technique is called when exiting the curve_approximation_technique production.
	ExitCurve_approximation_technique(c *Curve_approximation_techniqueContext)

	// ExitSurface_approximation_technique is called when exiting the surface_approximation_technique production.
	ExitSurface_approximation_technique(c *Surface_approximation_techniqueContext)

	// ExitDecimal is called when exiting the decimal production.
	ExitDecimal(c *DecimalContext)
}
