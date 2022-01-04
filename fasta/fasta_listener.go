// Code generated from fasta.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fasta // fasta
import "github.com/antlr/antlr4/runtime/Go/antlr"

// fastaListener is a complete listener for a parse tree produced by fastaParser.
type fastaListener interface {
	antlr.ParseTreeListener

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterSequencelines is called when entering the sequencelines production.
	EnterSequencelines(c *SequencelinesContext)

	// EnterDescriptionline is called when entering the descriptionline production.
	EnterDescriptionline(c *DescriptionlineContext)

	// EnterCommentline is called when entering the commentline production.
	EnterCommentline(c *CommentlineContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitSequencelines is called when exiting the sequencelines production.
	ExitSequencelines(c *SequencelinesContext)

	// ExitDescriptionline is called when exiting the descriptionline production.
	ExitDescriptionline(c *DescriptionlineContext)

	// ExitCommentline is called when exiting the commentline production.
	ExitCommentline(c *CommentlineContext)
}
