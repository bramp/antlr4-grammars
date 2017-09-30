// Generated from fasta.g4 by ANTLR 4.7.

package fasta // fasta
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasefastaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasefastaVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefastaVisitor) VisitSection(ctx *SectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefastaVisitor) VisitSequencelines(ctx *SequencelinesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefastaVisitor) VisitDescriptionline(ctx *DescriptionlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasefastaVisitor) VisitCommentline(ctx *CommentlineContext) interface{} {
	return v.VisitChildren(ctx)
}
