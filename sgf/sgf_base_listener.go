// Code generated from sgf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sgf // sgf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesgfListener is a complete listener for a parse tree produced by sgfParser.
type BasesgfListener struct{}

var _ sgfListener = &BasesgfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesgfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesgfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesgfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesgfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCollection is called when production collection is entered.
func (s *BasesgfListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BasesgfListener) ExitCollection(ctx *CollectionContext) {}

// EnterGameTree is called when production gameTree is entered.
func (s *BasesgfListener) EnterGameTree(ctx *GameTreeContext) {}

// ExitGameTree is called when production gameTree is exited.
func (s *BasesgfListener) ExitGameTree(ctx *GameTreeContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BasesgfListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BasesgfListener) ExitSequence(ctx *SequenceContext) {}

// EnterNode is called when production node is entered.
func (s *BasesgfListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BasesgfListener) ExitNode(ctx *NodeContext) {}

// EnterProperty_ is called when production property_ is entered.
func (s *BasesgfListener) EnterProperty_(ctx *Property_Context) {}

// ExitProperty_ is called when production property_ is exited.
func (s *BasesgfListener) ExitProperty_(ctx *Property_Context) {}

// EnterMove is called when production move is entered.
func (s *BasesgfListener) EnterMove(ctx *MoveContext) {}

// ExitMove is called when production move is exited.
func (s *BasesgfListener) ExitMove(ctx *MoveContext) {}

// EnterSetup is called when production setup is entered.
func (s *BasesgfListener) EnterSetup(ctx *SetupContext) {}

// ExitSetup is called when production setup is exited.
func (s *BasesgfListener) ExitSetup(ctx *SetupContext) {}

// EnterNodeAnnotation is called when production nodeAnnotation is entered.
func (s *BasesgfListener) EnterNodeAnnotation(ctx *NodeAnnotationContext) {}

// ExitNodeAnnotation is called when production nodeAnnotation is exited.
func (s *BasesgfListener) ExitNodeAnnotation(ctx *NodeAnnotationContext) {}

// EnterMoveAnnotation is called when production moveAnnotation is entered.
func (s *BasesgfListener) EnterMoveAnnotation(ctx *MoveAnnotationContext) {}

// ExitMoveAnnotation is called when production moveAnnotation is exited.
func (s *BasesgfListener) ExitMoveAnnotation(ctx *MoveAnnotationContext) {}

// EnterMarkup is called when production markup is entered.
func (s *BasesgfListener) EnterMarkup(ctx *MarkupContext) {}

// ExitMarkup is called when production markup is exited.
func (s *BasesgfListener) ExitMarkup(ctx *MarkupContext) {}

// EnterRoot is called when production root is entered.
func (s *BasesgfListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasesgfListener) ExitRoot(ctx *RootContext) {}

// EnterGameInfo is called when production gameInfo is entered.
func (s *BasesgfListener) EnterGameInfo(ctx *GameInfoContext) {}

// ExitGameInfo is called when production gameInfo is exited.
func (s *BasesgfListener) ExitGameInfo(ctx *GameInfoContext) {}

// EnterTiming is called when production timing is entered.
func (s *BasesgfListener) EnterTiming(ctx *TimingContext) {}

// ExitTiming is called when production timing is exited.
func (s *BasesgfListener) ExitTiming(ctx *TimingContext) {}

// EnterMisc is called when production misc is entered.
func (s *BasesgfListener) EnterMisc(ctx *MiscContext) {}

// ExitMisc is called when production misc is exited.
func (s *BasesgfListener) ExitMisc(ctx *MiscContext) {}

// EnterLoa is called when production loa is entered.
func (s *BasesgfListener) EnterLoa(ctx *LoaContext) {}

// ExitLoa is called when production loa is exited.
func (s *BasesgfListener) ExitLoa(ctx *LoaContext) {}

// EnterGo_ is called when production go_ is entered.
func (s *BasesgfListener) EnterGo_(ctx *Go_Context) {}

// ExitGo_ is called when production go_ is exited.
func (s *BasesgfListener) ExitGo_(ctx *Go_Context) {}

// EnterPrivateProp is called when production privateProp is entered.
func (s *BasesgfListener) EnterPrivateProp(ctx *PrivatePropContext) {}

// ExitPrivateProp is called when production privateProp is exited.
func (s *BasesgfListener) ExitPrivateProp(ctx *PrivatePropContext) {}
