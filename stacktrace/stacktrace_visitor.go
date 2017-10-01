// Generated from StackTrace.g4 by ANTLR 4.7.

package stacktrace // StackTrace
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by StackTraceParser.
type StackTraceVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by StackTraceParser#startRule.
	VisitStartRule(ctx *StartRuleContext) interface{}

	// Visit a parse tree produced by StackTraceParser#stackTrace.
	VisitStackTrace(ctx *StackTraceContext) interface{}

	// Visit a parse tree produced by StackTraceParser#stackTraceLine.
	VisitStackTraceLine(ctx *StackTraceLineContext) interface{}

	// Visit a parse tree produced by StackTraceParser#atLine.
	VisitAtLine(ctx *AtLineContext) interface{}

	// Visit a parse tree produced by StackTraceParser#causedByLine.
	VisitCausedByLine(ctx *CausedByLineContext) interface{}

	// Visit a parse tree produced by StackTraceParser#ellipsisLine.
	VisitEllipsisLine(ctx *EllipsisLineContext) interface{}

	// Visit a parse tree produced by StackTraceParser#messageLine.
	VisitMessageLine(ctx *MessageLineContext) interface{}

	// Visit a parse tree produced by StackTraceParser#qualifiedClass.
	VisitQualifiedClass(ctx *QualifiedClassContext) interface{}

	// Visit a parse tree produced by StackTraceParser#innerClassName.
	VisitInnerClassName(ctx *InnerClassNameContext) interface{}

	// Visit a parse tree produced by StackTraceParser#classFile.
	VisitClassFile(ctx *ClassFileContext) interface{}

	// Visit a parse tree produced by StackTraceParser#qualifiedMethod.
	VisitQualifiedMethod(ctx *QualifiedMethodContext) interface{}

	// Visit a parse tree produced by StackTraceParser#constructor.
	VisitConstructor(ctx *ConstructorContext) interface{}

	// Visit a parse tree produced by StackTraceParser#methodName.
	VisitMethodName(ctx *MethodNameContext) interface{}

	// Visit a parse tree produced by StackTraceParser#packagePath.
	VisitPackagePath(ctx *PackagePathContext) interface{}

	// Visit a parse tree produced by StackTraceParser#className.
	VisitClassName(ctx *ClassNameContext) interface{}

	// Visit a parse tree produced by StackTraceParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by StackTraceParser#message.
	VisitMessage(ctx *MessageContext) interface{}
}
