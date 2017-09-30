// Generated from lcc.g4 by ANTLR 4.7.

package lcc // lcc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselccListener is a complete listener for a parse tree produced by lccParser.
type BaselccListener struct{}

var _ lccListener = &BaselccListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselccListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselccListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselccListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselccListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLcc is called when production lcc is entered.
func (s *BaselccListener) EnterLcc(ctx *LccContext) {}

// ExitLcc is called when production lcc is exited.
func (s *BaselccListener) ExitLcc(ctx *LccContext) {}

// EnterTopic is called when production topic is entered.
func (s *BaselccListener) EnterTopic(ctx *TopicContext) {}

// ExitTopic is called when production topic is exited.
func (s *BaselccListener) ExitTopic(ctx *TopicContext) {}

// EnterSubtopic is called when production subtopic is entered.
func (s *BaselccListener) EnterSubtopic(ctx *SubtopicContext) {}

// ExitSubtopic is called when production subtopic is exited.
func (s *BaselccListener) ExitSubtopic(ctx *SubtopicContext) {}

// EnterSubclasses is called when production subclasses is entered.
func (s *BaselccListener) EnterSubclasses(ctx *SubclassesContext) {}

// ExitSubclasses is called when production subclasses is exited.
func (s *BaselccListener) ExitSubclasses(ctx *SubclassesContext) {}

// EnterSubclass is called when production subclass is entered.
func (s *BaselccListener) EnterSubclass(ctx *SubclassContext) {}

// ExitSubclass is called when production subclass is exited.
func (s *BaselccListener) ExitSubclass(ctx *SubclassContext) {}

// EnterCutters is called when production cutters is entered.
func (s *BaselccListener) EnterCutters(ctx *CuttersContext) {}

// ExitCutters is called when production cutters is exited.
func (s *BaselccListener) ExitCutters(ctx *CuttersContext) {}

// EnterCutter is called when production cutter is entered.
func (s *BaselccListener) EnterCutter(ctx *CutterContext) {}

// ExitCutter is called when production cutter is exited.
func (s *BaselccListener) ExitCutter(ctx *CutterContext) {}

// EnterDate is called when production date is entered.
func (s *BaselccListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaselccListener) ExitDate(ctx *DateContext) {}
