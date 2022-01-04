// Code generated from fen.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fen // fen
import "github.com/antlr/antlr4/runtime/Go/antlr"

// fenListener is a complete listener for a parse tree produced by fenParser.
type fenListener interface {
	antlr.ParseTreeListener

	// EnterFen is called when entering the fen production.
	EnterFen(c *FenContext)

	// EnterColor is called when entering the color production.
	EnterColor(c *ColorContext)

	// EnterCastling is called when entering the castling production.
	EnterCastling(c *CastlingContext)

	// EnterEnpassant is called when entering the enpassant production.
	EnterEnpassant(c *EnpassantContext)

	// EnterPosition is called when entering the position production.
	EnterPosition(c *PositionContext)

	// EnterHalfmoveclock is called when entering the halfmoveclock production.
	EnterHalfmoveclock(c *HalfmoveclockContext)

	// EnterFullmoveclock is called when entering the fullmoveclock production.
	EnterFullmoveclock(c *FullmoveclockContext)

	// EnterPlacement is called when entering the placement production.
	EnterPlacement(c *PlacementContext)

	// EnterRank is called when entering the rank production.
	EnterRank(c *RankContext)

	// EnterPiece is called when entering the piece production.
	EnterPiece(c *PieceContext)

	// ExitFen is called when exiting the fen production.
	ExitFen(c *FenContext)

	// ExitColor is called when exiting the color production.
	ExitColor(c *ColorContext)

	// ExitCastling is called when exiting the castling production.
	ExitCastling(c *CastlingContext)

	// ExitEnpassant is called when exiting the enpassant production.
	ExitEnpassant(c *EnpassantContext)

	// ExitPosition is called when exiting the position production.
	ExitPosition(c *PositionContext)

	// ExitHalfmoveclock is called when exiting the halfmoveclock production.
	ExitHalfmoveclock(c *HalfmoveclockContext)

	// ExitFullmoveclock is called when exiting the fullmoveclock production.
	ExitFullmoveclock(c *FullmoveclockContext)

	// ExitPlacement is called when exiting the placement production.
	ExitPlacement(c *PlacementContext)

	// ExitRank is called when exiting the rank production.
	ExitRank(c *RankContext)

	// ExitPiece is called when exiting the piece production.
	ExitPiece(c *PieceContext)
}
