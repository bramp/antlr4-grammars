// Code generated from gedcom.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gedcom // gedcom
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegedcomListener is a complete listener for a parse tree produced by gedcomParser.
type BasegedcomListener struct{}

var _ gedcomListener = &BasegedcomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegedcomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegedcomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegedcomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegedcomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGedcom is called when production gedcom is entered.
func (s *BasegedcomListener) EnterGedcom(ctx *GedcomContext) {}

// ExitGedcom is called when production gedcom is exited.
func (s *BasegedcomListener) ExitGedcom(ctx *GedcomContext) {}

// EnterLine is called when production line is entered.
func (s *BasegedcomListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasegedcomListener) ExitLine(ctx *LineContext) {}

// EnterLevel is called when production level is entered.
func (s *BasegedcomListener) EnterLevel(ctx *LevelContext) {}

// ExitLevel is called when production level is exited.
func (s *BasegedcomListener) ExitLevel(ctx *LevelContext) {}

// EnterOpt_xref_id is called when production opt_xref_id is entered.
func (s *BasegedcomListener) EnterOpt_xref_id(ctx *Opt_xref_idContext) {}

// ExitOpt_xref_id is called when production opt_xref_id is exited.
func (s *BasegedcomListener) ExitOpt_xref_id(ctx *Opt_xref_idContext) {}

// EnterTag is called when production tag is entered.
func (s *BasegedcomListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasegedcomListener) ExitTag(ctx *TagContext) {}

// EnterLine_value is called when production line_value is entered.
func (s *BasegedcomListener) EnterLine_value(ctx *Line_valueContext) {}

// ExitLine_value is called when production line_value is exited.
func (s *BasegedcomListener) ExitLine_value(ctx *Line_valueContext) {}

// EnterLine_item is called when production line_item is entered.
func (s *BasegedcomListener) EnterLine_item(ctx *Line_itemContext) {}

// ExitLine_item is called when production line_item is exited.
func (s *BasegedcomListener) ExitLine_item(ctx *Line_itemContext) {}

// EnterEscape is called when production escape is entered.
func (s *BasegedcomListener) EnterEscape(ctx *EscapeContext) {}

// ExitEscape is called when production escape is exited.
func (s *BasegedcomListener) ExitEscape(ctx *EscapeContext) {}

// EnterNon_at is called when production non_at is entered.
func (s *BasegedcomListener) EnterNon_at(ctx *Non_atContext) {}

// ExitNon_at is called when production non_at is exited.
func (s *BasegedcomListener) ExitNon_at(ctx *Non_atContext) {}

// EnterEscape_text is called when production escape_text is entered.
func (s *BasegedcomListener) EnterEscape_text(ctx *Escape_textContext) {}

// ExitEscape_text is called when production escape_text is exited.
func (s *BasegedcomListener) ExitEscape_text(ctx *Escape_textContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BasegedcomListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BasegedcomListener) ExitPointer(ctx *PointerContext) {}

// EnterPointer_string is called when production pointer_string is entered.
func (s *BasegedcomListener) EnterPointer_string(ctx *Pointer_stringContext) {}

// ExitPointer_string is called when production pointer_string is exited.
func (s *BasegedcomListener) ExitPointer_string(ctx *Pointer_stringContext) {}

// EnterPointer_char is called when production pointer_char is entered.
func (s *BasegedcomListener) EnterPointer_char(ctx *Pointer_charContext) {}

// ExitPointer_char is called when production pointer_char is exited.
func (s *BasegedcomListener) ExitPointer_char(ctx *Pointer_charContext) {}

// EnterAlphanum is called when production alphanum is entered.
func (s *BasegedcomListener) EnterAlphanum(ctx *AlphanumContext) {}

// ExitAlphanum is called when production alphanum is exited.
func (s *BasegedcomListener) ExitAlphanum(ctx *AlphanumContext) {}

// EnterAnychar is called when production anychar is entered.
func (s *BasegedcomListener) EnterAnychar(ctx *AnycharContext) {}

// ExitAnychar is called when production anychar is exited.
func (s *BasegedcomListener) ExitAnychar(ctx *AnycharContext) {}

// EnterOtherchar is called when production otherchar is entered.
func (s *BasegedcomListener) EnterOtherchar(ctx *OthercharContext) {}

// ExitOtherchar is called when production otherchar is exited.
func (s *BasegedcomListener) ExitOtherchar(ctx *OthercharContext) {}
