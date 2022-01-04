// Code generated from lcc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lcc // lcc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lccListener is a complete listener for a parse tree produced by lccParser.
type lccListener interface {
	antlr.ParseTreeListener

	// EnterLcc is called when entering the lcc production.
	EnterLcc(c *LccContext)

	// EnterTopic is called when entering the topic production.
	EnterTopic(c *TopicContext)

	// EnterSubtopic is called when entering the subtopic production.
	EnterSubtopic(c *SubtopicContext)

	// EnterSubclasses is called when entering the subclasses production.
	EnterSubclasses(c *SubclassesContext)

	// EnterSubclass is called when entering the subclass production.
	EnterSubclass(c *SubclassContext)

	// EnterCutters is called when entering the cutters production.
	EnterCutters(c *CuttersContext)

	// EnterCutter is called when entering the cutter production.
	EnterCutter(c *CutterContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// ExitLcc is called when exiting the lcc production.
	ExitLcc(c *LccContext)

	// ExitTopic is called when exiting the topic production.
	ExitTopic(c *TopicContext)

	// ExitSubtopic is called when exiting the subtopic production.
	ExitSubtopic(c *SubtopicContext)

	// ExitSubclasses is called when exiting the subclasses production.
	ExitSubclasses(c *SubclassesContext)

	// ExitSubclass is called when exiting the subclass production.
	ExitSubclass(c *SubclassContext)

	// ExitCutters is called when exiting the cutters production.
	ExitCutters(c *CuttersContext)

	// ExitCutter is called when exiting the cutter production.
	ExitCutter(c *CutterContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)
}
