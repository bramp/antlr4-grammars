// Code generated from fen.g4 by ANTLR 4.7.2. DO NOT EDIT.

package fen // fen
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefenListener is a complete listener for a parse tree produced by fenParser.
type BasefenListener struct{}

var _ fenListener = &BasefenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFen is called when production fen is entered.
func (s *BasefenListener) EnterFen(ctx *FenContext) {}

// ExitFen is called when production fen is exited.
func (s *BasefenListener) ExitFen(ctx *FenContext) {}

// EnterColor is called when production color is entered.
func (s *BasefenListener) EnterColor(ctx *ColorContext) {}

// ExitColor is called when production color is exited.
func (s *BasefenListener) ExitColor(ctx *ColorContext) {}

// EnterCastling is called when production castling is entered.
func (s *BasefenListener) EnterCastling(ctx *CastlingContext) {}

// ExitCastling is called when production castling is exited.
func (s *BasefenListener) ExitCastling(ctx *CastlingContext) {}

// EnterEnpassant is called when production enpassant is entered.
func (s *BasefenListener) EnterEnpassant(ctx *EnpassantContext) {}

// ExitEnpassant is called when production enpassant is exited.
func (s *BasefenListener) ExitEnpassant(ctx *EnpassantContext) {}

// EnterPosition is called when production position is entered.
func (s *BasefenListener) EnterPosition(ctx *PositionContext) {}

// ExitPosition is called when production position is exited.
func (s *BasefenListener) ExitPosition(ctx *PositionContext) {}

// EnterHalfmoveclock is called when production halfmoveclock is entered.
func (s *BasefenListener) EnterHalfmoveclock(ctx *HalfmoveclockContext) {}

// ExitHalfmoveclock is called when production halfmoveclock is exited.
func (s *BasefenListener) ExitHalfmoveclock(ctx *HalfmoveclockContext) {}

// EnterFullmoveclock is called when production fullmoveclock is entered.
func (s *BasefenListener) EnterFullmoveclock(ctx *FullmoveclockContext) {}

// ExitFullmoveclock is called when production fullmoveclock is exited.
func (s *BasefenListener) ExitFullmoveclock(ctx *FullmoveclockContext) {}

// EnterPlacement is called when production placement is entered.
func (s *BasefenListener) EnterPlacement(ctx *PlacementContext) {}

// ExitPlacement is called when production placement is exited.
func (s *BasefenListener) ExitPlacement(ctx *PlacementContext) {}

// EnterRank is called when production rank is entered.
func (s *BasefenListener) EnterRank(ctx *RankContext) {}

// ExitRank is called when production rank is exited.
func (s *BasefenListener) ExitRank(ctx *RankContext) {}

// EnterPiece is called when production piece is entered.
func (s *BasefenListener) EnterPiece(ctx *PieceContext) {}

// ExitPiece is called when production piece is exited.
func (s *BasefenListener) ExitPiece(ctx *PieceContext) {}
