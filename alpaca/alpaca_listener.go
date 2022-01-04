// Code generated from alpaca.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alpaca // alpaca
import "github.com/antlr/antlr4/runtime/Go/antlr"

// alpacaListener is a complete listener for a parse tree produced by alpacaParser.
type alpacaListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDefns is called when entering the defns production.
	EnterDefns(c *DefnsContext)

	// EnterDefn is called when entering the defn production.
	EnterDefn(c *DefnContext)

	// EnterStateDefn is called when entering the stateDefn production.
	EnterStateDefn(c *StateDefnContext)

	// EnterClassDefn is called when entering the classDefn production.
	EnterClassDefn(c *ClassDefnContext)

	// EnterNbhdDefn is called when entering the nbhdDefn production.
	EnterNbhdDefn(c *NbhdDefnContext)

	// EnterClassID is called when entering the classID production.
	EnterClassID(c *ClassIDContext)

	// EnterStateID is called when entering the stateID production.
	EnterStateID(c *StateIDContext)

	// EnterNbhdID is called when entering the nbhdID production.
	EnterNbhdID(c *NbhdIDContext)

	// EnterMembershipDecl is called when entering the membershipDecl production.
	EnterMembershipDecl(c *MembershipDeclContext)

	// EnterClassRef is called when entering the classRef production.
	EnterClassRef(c *ClassRefContext)

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterRule_ is called when entering the rule_ production.
	EnterRule_(c *Rule_Context)

	// EnterStateRef is called when entering the stateRef production.
	EnterStateRef(c *StateRefContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterRelationalPred is called when entering the relationalPred production.
	EnterRelationalPred(c *RelationalPredContext)

	// EnterAdjacencyPred is called when entering the adjacencyPred production.
	EnterAdjacencyPred(c *AdjacencyPredContext)

	// EnterBoolPrimitive is called when entering the boolPrimitive production.
	EnterBoolPrimitive(c *BoolPrimitiveContext)

	// EnterNeigbourhood is called when entering the neigbourhood production.
	EnterNeigbourhood(c *NeigbourhoodContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterNaturalnumber is called when entering the naturalnumber production.
	EnterNaturalnumber(c *NaturalnumberContext)

	// EnterArrowchain is called when entering the arrowchain production.
	EnterArrowchain(c *ArrowchainContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDefns is called when exiting the defns production.
	ExitDefns(c *DefnsContext)

	// ExitDefn is called when exiting the defn production.
	ExitDefn(c *DefnContext)

	// ExitStateDefn is called when exiting the stateDefn production.
	ExitStateDefn(c *StateDefnContext)

	// ExitClassDefn is called when exiting the classDefn production.
	ExitClassDefn(c *ClassDefnContext)

	// ExitNbhdDefn is called when exiting the nbhdDefn production.
	ExitNbhdDefn(c *NbhdDefnContext)

	// ExitClassID is called when exiting the classID production.
	ExitClassID(c *ClassIDContext)

	// ExitStateID is called when exiting the stateID production.
	ExitStateID(c *StateIDContext)

	// ExitNbhdID is called when exiting the nbhdID production.
	ExitNbhdID(c *NbhdIDContext)

	// ExitMembershipDecl is called when exiting the membershipDecl production.
	ExitMembershipDecl(c *MembershipDeclContext)

	// ExitClassRef is called when exiting the classRef production.
	ExitClassRef(c *ClassRefContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitRule_ is called when exiting the rule_ production.
	ExitRule_(c *Rule_Context)

	// ExitStateRef is called when exiting the stateRef production.
	ExitStateRef(c *StateRefContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitRelationalPred is called when exiting the relationalPred production.
	ExitRelationalPred(c *RelationalPredContext)

	// ExitAdjacencyPred is called when exiting the adjacencyPred production.
	ExitAdjacencyPred(c *AdjacencyPredContext)

	// ExitBoolPrimitive is called when exiting the boolPrimitive production.
	ExitBoolPrimitive(c *BoolPrimitiveContext)

	// ExitNeigbourhood is called when exiting the neigbourhood production.
	ExitNeigbourhood(c *NeigbourhoodContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitNaturalnumber is called when exiting the naturalnumber production.
	ExitNaturalnumber(c *NaturalnumberContext)

	// ExitArrowchain is called when exiting the arrowchain production.
	ExitArrowchain(c *ArrowchainContext)
}
