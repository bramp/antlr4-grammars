// Generated from CLIF.g4 by ANTLR 4.7.

package clif // CLIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCLIFVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCLIFVisitor) VisitTermseq(ctx *TermseqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitInterpretedname(ctx *InterpretednameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitInterpretablename(ctx *InterpretablenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitSentence(ctx *SentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitAtomsent(ctx *AtomsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitBoolsent(ctx *BoolsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitQuantsent(ctx *QuantsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitBoundlist(ctx *BoundlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitCommentsent(ctx *CommentsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitPhrase(ctx *PhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitCltext(ctx *CltextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCLIFVisitor) VisitNamedtext(ctx *NamedtextContext) interface{} {
	return v.VisitChildren(ctx)
}
