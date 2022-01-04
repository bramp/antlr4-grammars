// Code generated from alpaca.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alpaca // alpaca
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasealpacaListener is a complete listener for a parse tree produced by alpacaParser.
type BasealpacaListener struct{}

var _ alpacaListener = &BasealpacaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasealpacaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasealpacaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasealpacaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasealpacaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasealpacaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasealpacaListener) ExitProg(ctx *ProgContext) {}

// EnterDefns is called when production defns is entered.
func (s *BasealpacaListener) EnterDefns(ctx *DefnsContext) {}

// ExitDefns is called when production defns is exited.
func (s *BasealpacaListener) ExitDefns(ctx *DefnsContext) {}

// EnterDefn is called when production defn is entered.
func (s *BasealpacaListener) EnterDefn(ctx *DefnContext) {}

// ExitDefn is called when production defn is exited.
func (s *BasealpacaListener) ExitDefn(ctx *DefnContext) {}

// EnterStateDefn is called when production stateDefn is entered.
func (s *BasealpacaListener) EnterStateDefn(ctx *StateDefnContext) {}

// ExitStateDefn is called when production stateDefn is exited.
func (s *BasealpacaListener) ExitStateDefn(ctx *StateDefnContext) {}

// EnterClassDefn is called when production classDefn is entered.
func (s *BasealpacaListener) EnterClassDefn(ctx *ClassDefnContext) {}

// ExitClassDefn is called when production classDefn is exited.
func (s *BasealpacaListener) ExitClassDefn(ctx *ClassDefnContext) {}

// EnterNbhdDefn is called when production nbhdDefn is entered.
func (s *BasealpacaListener) EnterNbhdDefn(ctx *NbhdDefnContext) {}

// ExitNbhdDefn is called when production nbhdDefn is exited.
func (s *BasealpacaListener) ExitNbhdDefn(ctx *NbhdDefnContext) {}

// EnterClassID is called when production classID is entered.
func (s *BasealpacaListener) EnterClassID(ctx *ClassIDContext) {}

// ExitClassID is called when production classID is exited.
func (s *BasealpacaListener) ExitClassID(ctx *ClassIDContext) {}

// EnterStateID is called when production stateID is entered.
func (s *BasealpacaListener) EnterStateID(ctx *StateIDContext) {}

// ExitStateID is called when production stateID is exited.
func (s *BasealpacaListener) ExitStateID(ctx *StateIDContext) {}

// EnterNbhdID is called when production nbhdID is entered.
func (s *BasealpacaListener) EnterNbhdID(ctx *NbhdIDContext) {}

// ExitNbhdID is called when production nbhdID is exited.
func (s *BasealpacaListener) ExitNbhdID(ctx *NbhdIDContext) {}

// EnterMembershipDecl is called when production membershipDecl is entered.
func (s *BasealpacaListener) EnterMembershipDecl(ctx *MembershipDeclContext) {}

// ExitMembershipDecl is called when production membershipDecl is exited.
func (s *BasealpacaListener) ExitMembershipDecl(ctx *MembershipDeclContext) {}

// EnterClassRef is called when production classRef is entered.
func (s *BasealpacaListener) EnterClassRef(ctx *ClassRefContext) {}

// ExitClassRef is called when production classRef is exited.
func (s *BasealpacaListener) ExitClassRef(ctx *ClassRefContext) {}

// EnterRules is called when production rules is entered.
func (s *BasealpacaListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BasealpacaListener) ExitRules(ctx *RulesContext) {}

// EnterRule_ is called when production rule_ is entered.
func (s *BasealpacaListener) EnterRule_(ctx *Rule_Context) {}

// ExitRule_ is called when production rule_ is exited.
func (s *BasealpacaListener) ExitRule_(ctx *Rule_Context) {}

// EnterStateRef is called when production stateRef is entered.
func (s *BasealpacaListener) EnterStateRef(ctx *StateRefContext) {}

// ExitStateRef is called when production stateRef is exited.
func (s *BasealpacaListener) ExitStateRef(ctx *StateRefContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasealpacaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasealpacaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BasealpacaListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasealpacaListener) ExitTerm(ctx *TermContext) {}

// EnterRelationalPred is called when production relationalPred is entered.
func (s *BasealpacaListener) EnterRelationalPred(ctx *RelationalPredContext) {}

// ExitRelationalPred is called when production relationalPred is exited.
func (s *BasealpacaListener) ExitRelationalPred(ctx *RelationalPredContext) {}

// EnterAdjacencyPred is called when production adjacencyPred is entered.
func (s *BasealpacaListener) EnterAdjacencyPred(ctx *AdjacencyPredContext) {}

// ExitAdjacencyPred is called when production adjacencyPred is exited.
func (s *BasealpacaListener) ExitAdjacencyPred(ctx *AdjacencyPredContext) {}

// EnterBoolPrimitive is called when production boolPrimitive is entered.
func (s *BasealpacaListener) EnterBoolPrimitive(ctx *BoolPrimitiveContext) {}

// ExitBoolPrimitive is called when production boolPrimitive is exited.
func (s *BasealpacaListener) ExitBoolPrimitive(ctx *BoolPrimitiveContext) {}

// EnterNeigbourhood is called when production neigbourhood is entered.
func (s *BasealpacaListener) EnterNeigbourhood(ctx *NeigbourhoodContext) {}

// ExitNeigbourhood is called when production neigbourhood is exited.
func (s *BasealpacaListener) ExitNeigbourhood(ctx *NeigbourhoodContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasealpacaListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasealpacaListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterNaturalnumber is called when production naturalnumber is entered.
func (s *BasealpacaListener) EnterNaturalnumber(ctx *NaturalnumberContext) {}

// ExitNaturalnumber is called when production naturalnumber is exited.
func (s *BasealpacaListener) ExitNaturalnumber(ctx *NaturalnumberContext) {}

// EnterArrowchain is called when production arrowchain is entered.
func (s *BasealpacaListener) EnterArrowchain(ctx *ArrowchainContext) {}

// ExitArrowchain is called when production arrowchain is exited.
func (s *BasealpacaListener) ExitArrowchain(ctx *ArrowchainContext) {}
