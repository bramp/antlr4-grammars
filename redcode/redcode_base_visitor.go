// Generated from redcode.g4 by ANTLR 4.7.

package redcode // redcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseredcodeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseredcodeVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitOpcode(ctx *OpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitModifier(ctx *ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitMmode(ctx *MmodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseredcodeVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
