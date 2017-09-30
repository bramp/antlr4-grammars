// Generated from WavefrontOBJ.g4 by ANTLR 4.7.

package wavefront // WavefrontOBJ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by WavefrontOBJParser.
type WavefrontOBJVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by WavefrontOBJParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#face.
	VisitFace(ctx *FaceContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#vertex.
	VisitVertex(ctx *VertexContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#mtllib.
	VisitMtllib(ctx *MtllibContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#use_material.
	VisitUse_material(ctx *Use_materialContext) interface{}

	// Visit a parse tree produced by WavefrontOBJParser#group.
	VisitGroup(ctx *GroupContext) interface{}
}
