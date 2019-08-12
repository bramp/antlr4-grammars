// Code generated from WavefrontOBJ.g4 by ANTLR 4.7.2. DO NOT EDIT.

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

// EnterLine is called when production line is entered.
func (s *BaseWavefrontOBJListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseWavefrontOBJListener) ExitLine(ctx *LineContext) {}

// EnterFace is called when production face is entered.
func (s *BaseWavefrontOBJListener) EnterFace(ctx *FaceContext) {}

// ExitFace is called when production face is exited.
func (s *BaseWavefrontOBJListener) ExitFace(ctx *FaceContext) {}

// EnterVertex is called when production vertex is entered.
func (s *BaseWavefrontOBJListener) EnterVertex(ctx *VertexContext) {}

// ExitVertex is called when production vertex is exited.
func (s *BaseWavefrontOBJListener) ExitVertex(ctx *VertexContext) {}

// EnterMtllib is called when production mtllib is entered.
func (s *BaseWavefrontOBJListener) EnterMtllib(ctx *MtllibContext) {}

// ExitMtllib is called when production mtllib is exited.
func (s *BaseWavefrontOBJListener) ExitMtllib(ctx *MtllibContext) {}

// EnterObject is called when production object is entered.
func (s *BaseWavefrontOBJListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseWavefrontOBJListener) ExitObject(ctx *ObjectContext) {}

// EnterUse_material is called when production use_material is entered.
func (s *BaseWavefrontOBJListener) EnterUse_material(ctx *Use_materialContext) {}

// ExitUse_material is called when production use_material is exited.
func (s *BaseWavefrontOBJListener) ExitUse_material(ctx *Use_materialContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseWavefrontOBJListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseWavefrontOBJListener) ExitGroup(ctx *GroupContext) {}
