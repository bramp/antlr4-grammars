// Code generated from StackTrace.g4 by ANTLR 4.7.2. DO NOT EDIT.

package stacktrace // StackTrace
import "github.com/antlr/antlr4/runtime/Go/antlr"

// StackTraceListener is a complete listener for a parse tree produced by StackTraceParser.
type StackTraceListener interface {
	antlr.ParseTreeListener

	// EnterStartRule is called when entering the startRule production.
	EnterStartRule(c *StartRuleContext)

	// EnterStackTrace is called when entering the stackTrace production.
	EnterStackTrace(c *StackTraceContext)

	// EnterStackTraceLine is called when entering the stackTraceLine production.
	EnterStackTraceLine(c *StackTraceLineContext)

	// EnterAtLine is called when entering the atLine production.
	EnterAtLine(c *AtLineContext)

	// EnterCausedByLine is called when entering the causedByLine production.
	EnterCausedByLine(c *CausedByLineContext)

	// EnterEllipsisLine is called when entering the ellipsisLine production.
	EnterEllipsisLine(c *EllipsisLineContext)

	// EnterMessageLine is called when entering the messageLine production.
	EnterMessageLine(c *MessageLineContext)

	// EnterQualifiedClass is called when entering the qualifiedClass production.
	EnterQualifiedClass(c *QualifiedClassContext)

	// EnterInnerClassName is called when entering the innerClassName production.
	EnterInnerClassName(c *InnerClassNameContext)

	// EnterClassFile is called when entering the classFile production.
	EnterClassFile(c *ClassFileContext)

	// EnterQualifiedMethod is called when entering the qualifiedMethod production.
	EnterQualifiedMethod(c *QualifiedMethodContext)

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterMethodName is called when entering the methodName production.
	EnterMethodName(c *MethodNameContext)

	// EnterPackagePath is called when entering the packagePath production.
	EnterPackagePath(c *PackagePathContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// ExitStartRule is called when exiting the startRule production.
	ExitStartRule(c *StartRuleContext)

	// ExitStackTrace is called when exiting the stackTrace production.
	ExitStackTrace(c *StackTraceContext)

	// ExitStackTraceLine is called when exiting the stackTraceLine production.
	ExitStackTraceLine(c *StackTraceLineContext)

	// ExitAtLine is called when exiting the atLine production.
	ExitAtLine(c *AtLineContext)

	// ExitCausedByLine is called when exiting the causedByLine production.
	ExitCausedByLine(c *CausedByLineContext)

	// ExitEllipsisLine is called when exiting the ellipsisLine production.
	ExitEllipsisLine(c *EllipsisLineContext)

	// ExitMessageLine is called when exiting the messageLine production.
	ExitMessageLine(c *MessageLineContext)

	// ExitQualifiedClass is called when exiting the qualifiedClass production.
	ExitQualifiedClass(c *QualifiedClassContext)

	// ExitInnerClassName is called when exiting the innerClassName production.
	ExitInnerClassName(c *InnerClassNameContext)

	// ExitClassFile is called when exiting the classFile production.
	ExitClassFile(c *ClassFileContext)

	// ExitQualifiedMethod is called when exiting the qualifiedMethod production.
	ExitQualifiedMethod(c *QualifiedMethodContext)

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitMethodName is called when exiting the methodName production.
	ExitMethodName(c *MethodNameContext)

	// ExitPackagePath is called when exiting the packagePath production.
	ExitPackagePath(c *PackagePathContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)
}
