// Generated from TURTLE.g4 by ANTLR 4.7.

package turtle // TURTLE
import "github.com/antlr/antlr4/runtime/Go/antlr"

// TURTLEListener is a complete listener for a parse tree produced by TURTLEParser.
type TURTLEListener interface {
	antlr.ParseTreeListener

	// EnterTurtleDoc is called when entering the turtleDoc production.
	EnterTurtleDoc(c *TurtleDocContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterPrefixID is called when entering the prefixID production.
	EnterPrefixID(c *PrefixIDContext)

	// EnterBase is called when entering the base production.
	EnterBase(c *BaseContext)

	// EnterSparqlBase is called when entering the sparqlBase production.
	EnterSparqlBase(c *SparqlBaseContext)

	// EnterSparqlPrefix is called when entering the sparqlPrefix production.
	EnterSparqlPrefix(c *SparqlPrefixContext)

	// EnterTriples is called when entering the triples production.
	EnterTriples(c *TriplesContext)

	// EnterPredicateObjectList is called when entering the predicateObjectList production.
	EnterPredicateObjectList(c *PredicateObjectListContext)

	// EnterObjectList is called when entering the objectList production.
	EnterObjectList(c *ObjectListContext)

	// EnterVerb is called when entering the verb production.
	EnterVerb(c *VerbContext)

	// EnterSubject is called when entering the subject production.
	EnterSubject(c *SubjectContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBlankNodePropertyList is called when entering the blankNodePropertyList production.
	EnterBlankNodePropertyList(c *BlankNodePropertyListContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterRdfLiteral is called when entering the rdfLiteral production.
	EnterRdfLiteral(c *RdfLiteralContext)

	// EnterIri is called when entering the iri production.
	EnterIri(c *IriContext)

	// ExitTurtleDoc is called when exiting the turtleDoc production.
	ExitTurtleDoc(c *TurtleDocContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitPrefixID is called when exiting the prefixID production.
	ExitPrefixID(c *PrefixIDContext)

	// ExitBase is called when exiting the base production.
	ExitBase(c *BaseContext)

	// ExitSparqlBase is called when exiting the sparqlBase production.
	ExitSparqlBase(c *SparqlBaseContext)

	// ExitSparqlPrefix is called when exiting the sparqlPrefix production.
	ExitSparqlPrefix(c *SparqlPrefixContext)

	// ExitTriples is called when exiting the triples production.
	ExitTriples(c *TriplesContext)

	// ExitPredicateObjectList is called when exiting the predicateObjectList production.
	ExitPredicateObjectList(c *PredicateObjectListContext)

	// ExitObjectList is called when exiting the objectList production.
	ExitObjectList(c *ObjectListContext)

	// ExitVerb is called when exiting the verb production.
	ExitVerb(c *VerbContext)

	// ExitSubject is called when exiting the subject production.
	ExitSubject(c *SubjectContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBlankNodePropertyList is called when exiting the blankNodePropertyList production.
	ExitBlankNodePropertyList(c *BlankNodePropertyListContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitRdfLiteral is called when exiting the rdfLiteral production.
	ExitRdfLiteral(c *RdfLiteralContext)

	// ExitIri is called when exiting the iri production.
	ExitIri(c *IriContext)
}
