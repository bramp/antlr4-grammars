// Generated from robotwar.g4 by ANTLR 4.7.

package robotwars // robotwar
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaserobotwarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaserobotwarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitAccumstatement(ctx *AccumstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitAccumexpression(ctx *AccumexpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitGosubstatement(ctx *GosubstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitGotostatement(ctx *GotostatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitTostatement(ctx *TostatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitEndsubstatement(ctx *EndsubstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitIfstatement(ctx *IfstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitOperation(ctx *OperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitRegister(ctx *RegisterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaserobotwarVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
