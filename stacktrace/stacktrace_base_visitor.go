// Generated from StackTrace.g4 by ANTLR 4.7.

package stacktrace // StackTrace
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseStackTraceVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStackTraceVisitor) VisitStartRule(ctx *StartRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitStackTrace(ctx *StackTraceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitStackTraceLine(ctx *StackTraceLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitAtLine(ctx *AtLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitCausedByLine(ctx *CausedByLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitEllipsisLine(ctx *EllipsisLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitMessageLine(ctx *MessageLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitQualifiedClass(ctx *QualifiedClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitInnerClassName(ctx *InnerClassNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitClassFile(ctx *ClassFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitQualifiedMethod(ctx *QualifiedMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitConstructor(ctx *ConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitMethodName(ctx *MethodNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitPackagePath(ctx *PackagePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitClassName(ctx *ClassNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStackTraceVisitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}
