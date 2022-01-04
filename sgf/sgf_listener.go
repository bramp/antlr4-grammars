// Code generated from sgf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sgf // sgf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// sgfListener is a complete listener for a parse tree produced by sgfParser.
type sgfListener interface {
	antlr.ParseTreeListener

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterGameTree is called when entering the gameTree production.
	EnterGameTree(c *GameTreeContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterProperty_ is called when entering the property_ production.
	EnterProperty_(c *Property_Context)

	// EnterMove is called when entering the move production.
	EnterMove(c *MoveContext)

	// EnterSetup is called when entering the setup production.
	EnterSetup(c *SetupContext)

	// EnterNodeAnnotation is called when entering the nodeAnnotation production.
	EnterNodeAnnotation(c *NodeAnnotationContext)

	// EnterMoveAnnotation is called when entering the moveAnnotation production.
	EnterMoveAnnotation(c *MoveAnnotationContext)

	// EnterMarkup is called when entering the markup production.
	EnterMarkup(c *MarkupContext)

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterGameInfo is called when entering the gameInfo production.
	EnterGameInfo(c *GameInfoContext)

	// EnterTiming is called when entering the timing production.
	EnterTiming(c *TimingContext)

	// EnterMisc is called when entering the misc production.
	EnterMisc(c *MiscContext)

	// EnterLoa is called when entering the loa production.
	EnterLoa(c *LoaContext)

	// EnterGo_ is called when entering the go_ production.
	EnterGo_(c *Go_Context)

	// EnterPrivateProp is called when entering the privateProp production.
	EnterPrivateProp(c *PrivatePropContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitGameTree is called when exiting the gameTree production.
	ExitGameTree(c *GameTreeContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitProperty_ is called when exiting the property_ production.
	ExitProperty_(c *Property_Context)

	// ExitMove is called when exiting the move production.
	ExitMove(c *MoveContext)

	// ExitSetup is called when exiting the setup production.
	ExitSetup(c *SetupContext)

	// ExitNodeAnnotation is called when exiting the nodeAnnotation production.
	ExitNodeAnnotation(c *NodeAnnotationContext)

	// ExitMoveAnnotation is called when exiting the moveAnnotation production.
	ExitMoveAnnotation(c *MoveAnnotationContext)

	// ExitMarkup is called when exiting the markup production.
	ExitMarkup(c *MarkupContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitGameInfo is called when exiting the gameInfo production.
	ExitGameInfo(c *GameInfoContext)

	// ExitTiming is called when exiting the timing production.
	ExitTiming(c *TimingContext)

	// ExitMisc is called when exiting the misc production.
	ExitMisc(c *MiscContext)

	// ExitLoa is called when exiting the loa production.
	ExitLoa(c *LoaContext)

	// ExitGo_ is called when exiting the go_ production.
	ExitGo_(c *Go_Context)

	// ExitPrivateProp is called when exiting the privateProp production.
	ExitPrivateProp(c *PrivatePropContext)
}
