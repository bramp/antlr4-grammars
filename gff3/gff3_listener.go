// Code generated from gff3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gff3 // gff3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// gff3Listener is a complete listener for a parse tree produced by gff3Parser.
type gff3Listener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterDataline is called when entering the dataline production.
	EnterDataline(c *DatalineContext)

	// EnterAttributes is called when entering the attributes production.
	EnterAttributes(c *AttributesContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterSeqid is called when entering the seqid production.
	EnterSeqid(c *SeqidContext)

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterStrand is called when entering the strand production.
	EnterStrand(c *StrandContext)

	// EnterScore is called when entering the score production.
	EnterScore(c *ScoreContext)

	// EnterPhase is called when entering the phase production.
	EnterPhase(c *PhaseContext)

	// EnterCommentline is called when entering the commentline production.
	EnterCommentline(c *CommentlineContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitDataline is called when exiting the dataline production.
	ExitDataline(c *DatalineContext)

	// ExitAttributes is called when exiting the attributes production.
	ExitAttributes(c *AttributesContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitSeqid is called when exiting the seqid production.
	ExitSeqid(c *SeqidContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitStrand is called when exiting the strand production.
	ExitStrand(c *StrandContext)

	// ExitScore is called when exiting the score production.
	ExitScore(c *ScoreContext)

	// ExitPhase is called when exiting the phase production.
	ExitPhase(c *PhaseContext)

	// ExitCommentline is called when exiting the commentline production.
	ExitCommentline(c *CommentlineContext)
}
