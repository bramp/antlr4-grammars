// Generated from WavefrontOBJ.g4 by ANTLR 4.7.

package wavefront // WavefrontOBJ
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseWavefrontOBJVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseWavefrontOBJVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitFace(ctx *FaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitVertex(ctx *VertexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitMtllib(ctx *MtllibContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitUse_material(ctx *Use_materialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseWavefrontOBJVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}
