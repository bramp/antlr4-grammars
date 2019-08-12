// Code generated from WavefrontOBJ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package wavefrontobj // WavefrontOBJ
import "github.com/antlr/antlr4/runtime/Go/antlr"

// WavefrontOBJListener is a complete listener for a parse tree produced by WavefrontOBJParser.
type WavefrontOBJListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterFace is called when entering the face production.
	EnterFace(c *FaceContext)

	// EnterVertex is called when entering the vertex production.
	EnterVertex(c *VertexContext)

	// EnterMtllib is called when entering the mtllib production.
	EnterMtllib(c *MtllibContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterUse_material is called when entering the use_material production.
	EnterUse_material(c *Use_materialContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitFace is called when exiting the face production.
	ExitFace(c *FaceContext)

	// ExitVertex is called when exiting the vertex production.
	ExitVertex(c *VertexContext)

	// ExitMtllib is called when exiting the mtllib production.
	ExitMtllib(c *MtllibContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitUse_material is called when exiting the use_material production.
	ExitUse_material(c *Use_materialContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)
}
