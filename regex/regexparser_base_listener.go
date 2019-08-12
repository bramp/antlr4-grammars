// Code generated from regexParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package regex // regexParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseregexParserListener is a complete listener for a parse tree produced by regexParser.
type BaseregexParserListener struct{}

var _ regexParserListener = &BaseregexParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseregexParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseregexParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseregexParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseregexParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseregexParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseregexParserListener) ExitRoot(ctx *RootContext) {}

// EnterRegExp is called when production regExp is entered.
func (s *BaseregexParserListener) EnterRegExp(ctx *RegExpContext) {}

// ExitRegExp is called when production regExp is exited.
func (s *BaseregexParserListener) ExitRegExp(ctx *RegExpContext) {}

// EnterBranch is called when production branch is entered.
func (s *BaseregexParserListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BaseregexParserListener) ExitBranch(ctx *BranchContext) {}

// EnterPiece is called when production piece is entered.
func (s *BaseregexParserListener) EnterPiece(ctx *PieceContext) {}

// ExitPiece is called when production piece is exited.
func (s *BaseregexParserListener) ExitPiece(ctx *PieceContext) {}

// EnterQuantifier is called when production quantifier is entered.
func (s *BaseregexParserListener) EnterQuantifier(ctx *QuantifierContext) {}

// ExitQuantifier is called when production quantifier is exited.
func (s *BaseregexParserListener) ExitQuantifier(ctx *QuantifierContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BaseregexParserListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BaseregexParserListener) ExitQuantity(ctx *QuantityContext) {}

// EnterQuantRange is called when production quantRange is entered.
func (s *BaseregexParserListener) EnterQuantRange(ctx *QuantRangeContext) {}

// ExitQuantRange is called when production quantRange is exited.
func (s *BaseregexParserListener) ExitQuantRange(ctx *QuantRangeContext) {}

// EnterQuantMin is called when production quantMin is entered.
func (s *BaseregexParserListener) EnterQuantMin(ctx *QuantMinContext) {}

// ExitQuantMin is called when production quantMin is exited.
func (s *BaseregexParserListener) ExitQuantMin(ctx *QuantMinContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseregexParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseregexParserListener) ExitAtom(ctx *AtomContext) {}

// EnterCharClass is called when production charClass is entered.
func (s *BaseregexParserListener) EnterCharClass(ctx *CharClassContext) {}

// ExitCharClass is called when production charClass is exited.
func (s *BaseregexParserListener) ExitCharClass(ctx *CharClassContext) {}

// EnterCharClassExpr is called when production charClassExpr is entered.
func (s *BaseregexParserListener) EnterCharClassExpr(ctx *CharClassExprContext) {}

// ExitCharClassExpr is called when production charClassExpr is exited.
func (s *BaseregexParserListener) ExitCharClassExpr(ctx *CharClassExprContext) {}

// EnterCharGroup is called when production charGroup is entered.
func (s *BaseregexParserListener) EnterCharGroup(ctx *CharGroupContext) {}

// ExitCharGroup is called when production charGroup is exited.
func (s *BaseregexParserListener) ExitCharGroup(ctx *CharGroupContext) {}

// EnterPosCharGroup is called when production posCharGroup is entered.
func (s *BaseregexParserListener) EnterPosCharGroup(ctx *PosCharGroupContext) {}

// ExitPosCharGroup is called when production posCharGroup is exited.
func (s *BaseregexParserListener) ExitPosCharGroup(ctx *PosCharGroupContext) {}

// EnterCharRange is called when production charRange is entered.
func (s *BaseregexParserListener) EnterCharRange(ctx *CharRangeContext) {}

// ExitCharRange is called when production charRange is exited.
func (s *BaseregexParserListener) ExitCharRange(ctx *CharRangeContext) {}

// EnterSeRange is called when production seRange is entered.
func (s *BaseregexParserListener) EnterSeRange(ctx *SeRangeContext) {}

// ExitSeRange is called when production seRange is exited.
func (s *BaseregexParserListener) ExitSeRange(ctx *SeRangeContext) {}

// EnterCharOrEsc is called when production charOrEsc is entered.
func (s *BaseregexParserListener) EnterCharOrEsc(ctx *CharOrEscContext) {}

// ExitCharOrEsc is called when production charOrEsc is exited.
func (s *BaseregexParserListener) ExitCharOrEsc(ctx *CharOrEscContext) {}

// EnterCharClassEsc is called when production charClassEsc is entered.
func (s *BaseregexParserListener) EnterCharClassEsc(ctx *CharClassEscContext) {}

// ExitCharClassEsc is called when production charClassEsc is exited.
func (s *BaseregexParserListener) ExitCharClassEsc(ctx *CharClassEscContext) {}

// EnterCatEsc is called when production catEsc is entered.
func (s *BaseregexParserListener) EnterCatEsc(ctx *CatEscContext) {}

// ExitCatEsc is called when production catEsc is exited.
func (s *BaseregexParserListener) ExitCatEsc(ctx *CatEscContext) {}

// EnterComplEsc is called when production complEsc is entered.
func (s *BaseregexParserListener) EnterComplEsc(ctx *ComplEscContext) {}

// ExitComplEsc is called when production complEsc is exited.
func (s *BaseregexParserListener) ExitComplEsc(ctx *ComplEscContext) {}

// EnterCharProp is called when production charProp is entered.
func (s *BaseregexParserListener) EnterCharProp(ctx *CharPropContext) {}

// ExitCharProp is called when production charProp is exited.
func (s *BaseregexParserListener) ExitCharProp(ctx *CharPropContext) {}
