// Code generated from properties.g4 by ANTLR 4.9.3. DO NOT EDIT.

package properties // properties
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepropertiesListener is a complete listener for a parse tree produced by propertiesParser.
type BasepropertiesListener struct{}

var _ propertiesListener = &BasepropertiesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepropertiesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepropertiesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepropertiesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepropertiesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPropertiesFile is called when production propertiesFile is entered.
func (s *BasepropertiesListener) EnterPropertiesFile(ctx *PropertiesFileContext) {}

// ExitPropertiesFile is called when production propertiesFile is exited.
func (s *BasepropertiesListener) ExitPropertiesFile(ctx *PropertiesFileContext) {}

// EnterRow is called when production row is entered.
func (s *BasepropertiesListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BasepropertiesListener) ExitRow(ctx *RowContext) {}

// EnterDecl is called when production decl is entered.
func (s *BasepropertiesListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BasepropertiesListener) ExitDecl(ctx *DeclContext) {}

// EnterKey is called when production key is entered.
func (s *BasepropertiesListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasepropertiesListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BasepropertiesListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasepropertiesListener) ExitValue(ctx *ValueContext) {}

// EnterComment is called when production comment is entered.
func (s *BasepropertiesListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasepropertiesListener) ExitComment(ctx *CommentContext) {}
