// Generated from p.g4 by ANTLR 4.7.

package p // p
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasepVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasepVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepVisitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepVisitor) VisitIterate(ctx *IterateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
