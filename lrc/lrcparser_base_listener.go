// Code generated from lrcParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lrc // lrcParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselrcParserListener is a complete listener for a parse tree produced by lrcParser.
type BaselrcParserListener struct{}

var _ lrcParserListener = &BaselrcParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselrcParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselrcParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselrcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselrcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLrc is called when production lrc is entered.
func (s *BaselrcParserListener) EnterLrc(ctx *LrcContext) {}

// ExitLrc is called when production lrc is exited.
func (s *BaselrcParserListener) ExitLrc(ctx *LrcContext) {}

// EnterLine is called when production line is entered.
func (s *BaselrcParserListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaselrcParserListener) ExitLine(ctx *LineContext) {}
