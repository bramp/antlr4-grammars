// Generated from StackTrace.g4 by ANTLR 4.7.

package stacktrace // StackTrace
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStackTraceListener is a complete listener for a parse tree produced by StackTraceParser.
type BaseStackTraceListener struct{}

var _ StackTraceListener = &BaseStackTraceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStackTraceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStackTraceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStackTraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStackTraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BaseStackTraceListener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BaseStackTraceListener) ExitStartRule(ctx *StartRuleContext) {}

// EnterStackTrace is called when production stackTrace is entered.
func (s *BaseStackTraceListener) EnterStackTrace(ctx *StackTraceContext) {}

// ExitStackTrace is called when production stackTrace is exited.
func (s *BaseStackTraceListener) ExitStackTrace(ctx *StackTraceContext) {}

// EnterStackTraceLine is called when production stackTraceLine is entered.
func (s *BaseStackTraceListener) EnterStackTraceLine(ctx *StackTraceLineContext) {}

// ExitStackTraceLine is called when production stackTraceLine is exited.
func (s *BaseStackTraceListener) ExitStackTraceLine(ctx *StackTraceLineContext) {}

// EnterAtLine is called when production atLine is entered.
func (s *BaseStackTraceListener) EnterAtLine(ctx *AtLineContext) {}

// ExitAtLine is called when production atLine is exited.
func (s *BaseStackTraceListener) ExitAtLine(ctx *AtLineContext) {}

// EnterCausedByLine is called when production causedByLine is entered.
func (s *BaseStackTraceListener) EnterCausedByLine(ctx *CausedByLineContext) {}

// ExitCausedByLine is called when production causedByLine is exited.
func (s *BaseStackTraceListener) ExitCausedByLine(ctx *CausedByLineContext) {}

// EnterEllipsisLine is called when production ellipsisLine is entered.
func (s *BaseStackTraceListener) EnterEllipsisLine(ctx *EllipsisLineContext) {}

// ExitEllipsisLine is called when production ellipsisLine is exited.
func (s *BaseStackTraceListener) ExitEllipsisLine(ctx *EllipsisLineContext) {}

// EnterMessageLine is called when production messageLine is entered.
func (s *BaseStackTraceListener) EnterMessageLine(ctx *MessageLineContext) {}

// ExitMessageLine is called when production messageLine is exited.
func (s *BaseStackTraceListener) ExitMessageLine(ctx *MessageLineContext) {}

// EnterQualifiedClass is called when production qualifiedClass is entered.
func (s *BaseStackTraceListener) EnterQualifiedClass(ctx *QualifiedClassContext) {}

// ExitQualifiedClass is called when production qualifiedClass is exited.
func (s *BaseStackTraceListener) ExitQualifiedClass(ctx *QualifiedClassContext) {}

// EnterInnerClassName is called when production innerClassName is entered.
func (s *BaseStackTraceListener) EnterInnerClassName(ctx *InnerClassNameContext) {}

// ExitInnerClassName is called when production innerClassName is exited.
func (s *BaseStackTraceListener) ExitInnerClassName(ctx *InnerClassNameContext) {}

// EnterClassFile is called when production classFile is entered.
func (s *BaseStackTraceListener) EnterClassFile(ctx *ClassFileContext) {}

// ExitClassFile is called when production classFile is exited.
func (s *BaseStackTraceListener) ExitClassFile(ctx *ClassFileContext) {}

// EnterQualifiedMethod is called when production qualifiedMethod is entered.
func (s *BaseStackTraceListener) EnterQualifiedMethod(ctx *QualifiedMethodContext) {}

// ExitQualifiedMethod is called when production qualifiedMethod is exited.
func (s *BaseStackTraceListener) ExitQualifiedMethod(ctx *QualifiedMethodContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BaseStackTraceListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BaseStackTraceListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseStackTraceListener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseStackTraceListener) ExitMethodName(ctx *MethodNameContext) {}

// EnterPackagePath is called when production packagePath is entered.
func (s *BaseStackTraceListener) EnterPackagePath(ctx *PackagePathContext) {}

// ExitPackagePath is called when production packagePath is exited.
func (s *BaseStackTraceListener) ExitPackagePath(ctx *PackagePathContext) {}

// EnterClassName is called when production className is entered.
func (s *BaseStackTraceListener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *BaseStackTraceListener) ExitClassName(ctx *ClassNameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseStackTraceListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseStackTraceListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseStackTraceListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseStackTraceListener) ExitMessage(ctx *MessageContext) {}
