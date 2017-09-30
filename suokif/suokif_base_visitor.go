// Generated from SUOKIF.g4 by ANTLR 4.7.

package suokif // SUOKIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSUOKIFVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSUOKIFVisitor) VisitTop_level(ctx *Top_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitFunterm(ctx *FuntermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitSentence(ctx *SentenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitRelsent(ctx *RelsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitLogsent(ctx *LogsentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSUOKIFVisitor) VisitQuantsent(ctx *QuantsentContext) interface{} {
	return v.VisitChildren(ctx)
}
