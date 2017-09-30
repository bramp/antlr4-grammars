// Generated from agc.g4 by ANTLR 4.7.

package agc // agc
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseagcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseagcVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitBlank_line(ctx *Blank_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitComment_line(ctx *Comment_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitInstruction_line(ctx *Instruction_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitErase_line(ctx *Erase_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitAssignment_line(ctx *Assignment_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitOpcodes(ctx *OpcodesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitWs(ctx *WsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitEol(ctx *EolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitInte(ctx *InteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitDecimal(ctx *DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitRegister(ctx *RegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitOpcode(ctx *OpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitAxt_opcode(ctx *Axt_opcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitPseudo_opcode(ctx *Pseudo_opcodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagcVisitor) VisitStandard_opcode(ctx *Standard_opcodeContext) interface{} {
	return v.VisitChildren(ctx)
}
