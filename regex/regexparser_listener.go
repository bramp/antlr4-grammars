// Code generated from regexParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package regex // regexParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// regexParserListener is a complete listener for a parse tree produced by regexParser.
type regexParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterRegExp is called when entering the regExp production.
	EnterRegExp(c *RegExpContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// EnterPiece is called when entering the piece production.
	EnterPiece(c *PieceContext)

	// EnterQuantifier is called when entering the quantifier production.
	EnterQuantifier(c *QuantifierContext)

	// EnterQuantity is called when entering the quantity production.
	EnterQuantity(c *QuantityContext)

	// EnterQuantRange is called when entering the quantRange production.
	EnterQuantRange(c *QuantRangeContext)

	// EnterQuantMin is called when entering the quantMin production.
	EnterQuantMin(c *QuantMinContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterCharClass is called when entering the charClass production.
	EnterCharClass(c *CharClassContext)

	// EnterCharClassExpr is called when entering the charClassExpr production.
	EnterCharClassExpr(c *CharClassExprContext)

	// EnterCharGroup is called when entering the charGroup production.
	EnterCharGroup(c *CharGroupContext)

	// EnterPosCharGroup is called when entering the posCharGroup production.
	EnterPosCharGroup(c *PosCharGroupContext)

	// EnterCharRange is called when entering the charRange production.
	EnterCharRange(c *CharRangeContext)

	// EnterSeRange is called when entering the seRange production.
	EnterSeRange(c *SeRangeContext)

	// EnterCharOrEsc is called when entering the charOrEsc production.
	EnterCharOrEsc(c *CharOrEscContext)

	// EnterCharClassEsc is called when entering the charClassEsc production.
	EnterCharClassEsc(c *CharClassEscContext)

	// EnterCatEsc is called when entering the catEsc production.
	EnterCatEsc(c *CatEscContext)

	// EnterComplEsc is called when entering the complEsc production.
	EnterComplEsc(c *ComplEscContext)

	// EnterCharProp is called when entering the charProp production.
	EnterCharProp(c *CharPropContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitRegExp is called when exiting the regExp production.
	ExitRegExp(c *RegExpContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)

	// ExitPiece is called when exiting the piece production.
	ExitPiece(c *PieceContext)

	// ExitQuantifier is called when exiting the quantifier production.
	ExitQuantifier(c *QuantifierContext)

	// ExitQuantity is called when exiting the quantity production.
	ExitQuantity(c *QuantityContext)

	// ExitQuantRange is called when exiting the quantRange production.
	ExitQuantRange(c *QuantRangeContext)

	// ExitQuantMin is called when exiting the quantMin production.
	ExitQuantMin(c *QuantMinContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitCharClass is called when exiting the charClass production.
	ExitCharClass(c *CharClassContext)

	// ExitCharClassExpr is called when exiting the charClassExpr production.
	ExitCharClassExpr(c *CharClassExprContext)

	// ExitCharGroup is called when exiting the charGroup production.
	ExitCharGroup(c *CharGroupContext)

	// ExitPosCharGroup is called when exiting the posCharGroup production.
	ExitPosCharGroup(c *PosCharGroupContext)

	// ExitCharRange is called when exiting the charRange production.
	ExitCharRange(c *CharRangeContext)

	// ExitSeRange is called when exiting the seRange production.
	ExitSeRange(c *SeRangeContext)

	// ExitCharOrEsc is called when exiting the charOrEsc production.
	ExitCharOrEsc(c *CharOrEscContext)

	// ExitCharClassEsc is called when exiting the charClassEsc production.
	ExitCharClassEsc(c *CharClassEscContext)

	// ExitCatEsc is called when exiting the catEsc production.
	ExitCatEsc(c *CatEscContext)

	// ExitComplEsc is called when exiting the complEsc production.
	ExitComplEsc(c *ComplEscContext)

	// ExitCharProp is called when exiting the charProp production.
	ExitCharProp(c *CharPropContext)
}
