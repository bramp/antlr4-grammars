// Generated from TURTLE.g4 by ANTLR 4.7.

package turtle // TURTLE
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTURTLEVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTURTLEVisitor) VisitTurtleDoc(ctx *TurtleDocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitPrefixID(ctx *PrefixIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitBase(ctx *BaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitSparqlBase(ctx *SparqlBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitSparqlPrefix(ctx *SparqlPrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitTriples(ctx *TriplesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitPredicateObjectList(ctx *PredicateObjectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitObjectList(ctx *ObjectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitVerb(ctx *VerbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitSubject(ctx *SubjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitBlankNodePropertyList(ctx *BlankNodePropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitCollection(ctx *CollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitRdfLiteral(ctx *RdfLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTURTLEVisitor) VisitIri(ctx *IriContext) interface{} {
	return v.VisitChildren(ctx)
}
