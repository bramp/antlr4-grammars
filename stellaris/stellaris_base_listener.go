// Code generated from stellaris.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stellaris // stellaris
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasestellarisListener is a complete listener for a parse tree produced by stellarisParser.
type BasestellarisListener struct{}

var _ stellarisListener = &BasestellarisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasestellarisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasestellarisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasestellarisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasestellarisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterContent is called when production content is entered.
func (s *BasestellarisListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BasestellarisListener) ExitContent(ctx *ContentContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasestellarisListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasestellarisListener) ExitExpr(ctx *ExprContext) {}

// EnterKeyval is called when production keyval is entered.
func (s *BasestellarisListener) EnterKeyval(ctx *KeyvalContext) {}

// ExitKeyval is called when production keyval is exited.
func (s *BasestellarisListener) ExitKeyval(ctx *KeyvalContext) {}

// EnterKey is called when production key is entered.
func (s *BasestellarisListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasestellarisListener) ExitKey(ctx *KeyContext) {}

// EnterVal is called when production val is entered.
func (s *BasestellarisListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BasestellarisListener) ExitVal(ctx *ValContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BasestellarisListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BasestellarisListener) ExitAttrib(ctx *AttribContext) {}

// EnterAccessor is called when production accessor is entered.
func (s *BasestellarisListener) EnterAccessor(ctx *AccessorContext) {}

// ExitAccessor is called when production accessor is exited.
func (s *BasestellarisListener) ExitAccessor(ctx *AccessorContext) {}

// EnterGroup is called when production group is entered.
func (s *BasestellarisListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BasestellarisListener) ExitGroup(ctx *GroupContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BasestellarisListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BasestellarisListener) ExitId_(ctx *Id_Context) {}
