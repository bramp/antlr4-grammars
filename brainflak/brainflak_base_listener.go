// Code generated from brainflak.g4 by ANTLR 4.9.3. DO NOT EDIT.

package brainflak // brainflak
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebrainflakListener is a complete listener for a parse tree produced by brainflakParser.
type BasebrainflakListener struct{}

var _ brainflakListener = &BasebrainflakListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebrainflakListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebrainflakListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebrainflakListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebrainflakListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile_ is called when production file_ is entered.
func (s *BasebrainflakListener) EnterFile_(ctx *File_Context) {}

// ExitFile_ is called when production file_ is exited.
func (s *BasebrainflakListener) ExitFile_(ctx *File_Context) {}

// EnterStatement is called when production statement is entered.
func (s *BasebrainflakListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasebrainflakListener) ExitStatement(ctx *StatementContext) {}

// EnterParenstmt is called when production parenstmt is entered.
func (s *BasebrainflakListener) EnterParenstmt(ctx *ParenstmtContext) {}

// ExitParenstmt is called when production parenstmt is exited.
func (s *BasebrainflakListener) ExitParenstmt(ctx *ParenstmtContext) {}

// EnterBracestmt is called when production bracestmt is entered.
func (s *BasebrainflakListener) EnterBracestmt(ctx *BracestmtContext) {}

// ExitBracestmt is called when production bracestmt is exited.
func (s *BasebrainflakListener) ExitBracestmt(ctx *BracestmtContext) {}

// EnterBracketstmt is called when production bracketstmt is entered.
func (s *BasebrainflakListener) EnterBracketstmt(ctx *BracketstmtContext) {}

// ExitBracketstmt is called when production bracketstmt is exited.
func (s *BasebrainflakListener) ExitBracketstmt(ctx *BracketstmtContext) {}

// EnterGtltstmt is called when production gtltstmt is entered.
func (s *BasebrainflakListener) EnterGtltstmt(ctx *GtltstmtContext) {}

// ExitGtltstmt is called when production gtltstmt is exited.
func (s *BasebrainflakListener) ExitGtltstmt(ctx *GtltstmtContext) {}
