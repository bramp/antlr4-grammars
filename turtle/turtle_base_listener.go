// Generated from TURTLE.g4 by ANTLR 4.7.

package turtle // TURTLE
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTURTLEListener is a complete listener for a parse tree produced by TURTLEParser.
type BaseTURTLEListener struct{}

var _ TURTLEListener = &BaseTURTLEListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTURTLEListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTURTLEListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTURTLEListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTURTLEListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTurtleDoc is called when production turtleDoc is entered.
func (s *BaseTURTLEListener) EnterTurtleDoc(ctx *TurtleDocContext) {}

// ExitTurtleDoc is called when production turtleDoc is exited.
func (s *BaseTURTLEListener) ExitTurtleDoc(ctx *TurtleDocContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTURTLEListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTURTLEListener) ExitStatement(ctx *StatementContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseTURTLEListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseTURTLEListener) ExitDirective(ctx *DirectiveContext) {}

// EnterPrefixID is called when production prefixID is entered.
func (s *BaseTURTLEListener) EnterPrefixID(ctx *PrefixIDContext) {}

// ExitPrefixID is called when production prefixID is exited.
func (s *BaseTURTLEListener) ExitPrefixID(ctx *PrefixIDContext) {}

// EnterBase is called when production base is entered.
func (s *BaseTURTLEListener) EnterBase(ctx *BaseContext) {}

// ExitBase is called when production base is exited.
func (s *BaseTURTLEListener) ExitBase(ctx *BaseContext) {}

// EnterSparqlBase is called when production sparqlBase is entered.
func (s *BaseTURTLEListener) EnterSparqlBase(ctx *SparqlBaseContext) {}

// ExitSparqlBase is called when production sparqlBase is exited.
func (s *BaseTURTLEListener) ExitSparqlBase(ctx *SparqlBaseContext) {}

// EnterSparqlPrefix is called when production sparqlPrefix is entered.
func (s *BaseTURTLEListener) EnterSparqlPrefix(ctx *SparqlPrefixContext) {}

// ExitSparqlPrefix is called when production sparqlPrefix is exited.
func (s *BaseTURTLEListener) ExitSparqlPrefix(ctx *SparqlPrefixContext) {}

// EnterTriples is called when production triples is entered.
func (s *BaseTURTLEListener) EnterTriples(ctx *TriplesContext) {}

// ExitTriples is called when production triples is exited.
func (s *BaseTURTLEListener) ExitTriples(ctx *TriplesContext) {}

// EnterPredicateObjectList is called when production predicateObjectList is entered.
func (s *BaseTURTLEListener) EnterPredicateObjectList(ctx *PredicateObjectListContext) {}

// ExitPredicateObjectList is called when production predicateObjectList is exited.
func (s *BaseTURTLEListener) ExitPredicateObjectList(ctx *PredicateObjectListContext) {}

// EnterObjectList is called when production objectList is entered.
func (s *BaseTURTLEListener) EnterObjectList(ctx *ObjectListContext) {}

// ExitObjectList is called when production objectList is exited.
func (s *BaseTURTLEListener) ExitObjectList(ctx *ObjectListContext) {}

// EnterVerb is called when production verb is entered.
func (s *BaseTURTLEListener) EnterVerb(ctx *VerbContext) {}

// ExitVerb is called when production verb is exited.
func (s *BaseTURTLEListener) ExitVerb(ctx *VerbContext) {}

// EnterSubject is called when production subject is entered.
func (s *BaseTURTLEListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BaseTURTLEListener) ExitSubject(ctx *SubjectContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseTURTLEListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseTURTLEListener) ExitPredicate(ctx *PredicateContext) {}

// EnterObject is called when production object is entered.
func (s *BaseTURTLEListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseTURTLEListener) ExitObject(ctx *ObjectContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTURTLEListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTURTLEListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBlankNodePropertyList is called when production blankNodePropertyList is entered.
func (s *BaseTURTLEListener) EnterBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// ExitBlankNodePropertyList is called when production blankNodePropertyList is exited.
func (s *BaseTURTLEListener) ExitBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseTURTLEListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseTURTLEListener) ExitCollection(ctx *CollectionContext) {}

// EnterRdfLiteral is called when production rdfLiteral is entered.
func (s *BaseTURTLEListener) EnterRdfLiteral(ctx *RdfLiteralContext) {}

// ExitRdfLiteral is called when production rdfLiteral is exited.
func (s *BaseTURTLEListener) ExitRdfLiteral(ctx *RdfLiteralContext) {}

// EnterIri is called when production iri is entered.
func (s *BaseTURTLEListener) EnterIri(ctx *IriContext) {}

// ExitIri is called when production iri is exited.
func (s *BaseTURTLEListener) ExitIri(ctx *IriContext) {}
