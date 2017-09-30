// Generated from fasta.g4 by ANTLR 4.7.

package fasta // fasta
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by fastaParser.
type fastaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by fastaParser#sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by fastaParser#section.
	VisitSection(ctx *SectionContext) interface{}

	// Visit a parse tree produced by fastaParser#sequencelines.
	VisitSequencelines(ctx *SequencelinesContext) interface{}

	// Visit a parse tree produced by fastaParser#descriptionline.
	VisitDescriptionline(ctx *DescriptionlineContext) interface{}

	// Visit a parse tree produced by fastaParser#commentline.
	VisitCommentline(ctx *CommentlineContext) interface{}
}
