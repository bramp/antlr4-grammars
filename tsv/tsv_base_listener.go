// Code generated from tsv.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tsv // tsv
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetsvListener is a complete listener for a parse tree produced by tsvParser.
type BasetsvListener struct{}

var _ tsvListener = &BasetsvListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetsvListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetsvListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetsvListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetsvListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTsvFile is called when production tsvFile is entered.
func (s *BasetsvListener) EnterTsvFile(ctx *TsvFileContext) {}

// ExitTsvFile is called when production tsvFile is exited.
func (s *BasetsvListener) ExitTsvFile(ctx *TsvFileContext) {}

// EnterHdr is called when production hdr is entered.
func (s *BasetsvListener) EnterHdr(ctx *HdrContext) {}

// ExitHdr is called when production hdr is exited.
func (s *BasetsvListener) ExitHdr(ctx *HdrContext) {}

// EnterRow is called when production row is entered.
func (s *BasetsvListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BasetsvListener) ExitRow(ctx *RowContext) {}

// EnterField is called when production field is entered.
func (s *BasetsvListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasetsvListener) ExitField(ctx *FieldContext) {}
