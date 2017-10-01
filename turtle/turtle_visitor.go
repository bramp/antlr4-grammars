// Generated from TURTLE.g4 by ANTLR 4.7.

package turtle // TURTLE
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TURTLEParser.
type TURTLEVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TURTLEParser#turtleDoc.
	VisitTurtleDoc(ctx *TurtleDocContext) interface{}

	// Visit a parse tree produced by TURTLEParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by TURTLEParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by TURTLEParser#prefixID.
	VisitPrefixID(ctx *PrefixIDContext) interface{}

	// Visit a parse tree produced by TURTLEParser#base.
	VisitBase(ctx *BaseContext) interface{}

	// Visit a parse tree produced by TURTLEParser#sparqlBase.
	VisitSparqlBase(ctx *SparqlBaseContext) interface{}

	// Visit a parse tree produced by TURTLEParser#sparqlPrefix.
	VisitSparqlPrefix(ctx *SparqlPrefixContext) interface{}

	// Visit a parse tree produced by TURTLEParser#triples.
	VisitTriples(ctx *TriplesContext) interface{}

	// Visit a parse tree produced by TURTLEParser#predicateObjectList.
	VisitPredicateObjectList(ctx *PredicateObjectListContext) interface{}

	// Visit a parse tree produced by TURTLEParser#objectList.
	VisitObjectList(ctx *ObjectListContext) interface{}

	// Visit a parse tree produced by TURTLEParser#verb.
	VisitVerb(ctx *VerbContext) interface{}

	// Visit a parse tree produced by TURTLEParser#subject.
	VisitSubject(ctx *SubjectContext) interface{}

	// Visit a parse tree produced by TURTLEParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by TURTLEParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by TURTLEParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TURTLEParser#blankNodePropertyList.
	VisitBlankNodePropertyList(ctx *BlankNodePropertyListContext) interface{}

	// Visit a parse tree produced by TURTLEParser#collection.
	VisitCollection(ctx *CollectionContext) interface{}

	// Visit a parse tree produced by TURTLEParser#rdfLiteral.
	VisitRdfLiteral(ctx *RdfLiteralContext) interface{}

	// Visit a parse tree produced by TURTLEParser#iri.
	VisitIri(ctx *IriContext) interface{}
}
