// Generated from MuParser.g4 by ANTLR 4.7.

package muparser // MuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MuParserListener is a complete listener for a parse tree produced by MuParserParser.
type MuParserListener interface {
	antlr.ParseTreeListener

	// EnterProgExpr is called when entering the progExpr production.
	EnterProgExpr(c *ProgExprContext)

	// EnterFunctionMultiExpr is called when entering the functionMultiExpr production.
	EnterFunctionMultiExpr(c *FunctionMultiExprContext)

	// EnterAddSubExpr is called when entering the addSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterFunctionExpr is called when entering the functionExpr production.
	EnterFunctionExpr(c *FunctionExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterMulDivExpr is called when entering the mulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterIteExpr is called when entering the iteExpr production.
	EnterIteExpr(c *IteExprContext)

	// EnterParExpr is called when entering the parExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterNumberAtom is called when entering the numberAtom production.
	EnterNumberAtom(c *NumberAtomContext)

	// EnterBooleanAtom is called when entering the booleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterPredefinedConstantAtom is called when entering the predefinedConstantAtom production.
	EnterPredefinedConstantAtom(c *PredefinedConstantAtomContext)

	// EnterIdAtom is called when entering the idAtom production.
	EnterIdAtom(c *IdAtomContext)

	// ExitProgExpr is called when exiting the progExpr production.
	ExitProgExpr(c *ProgExprContext)

	// ExitFunctionMultiExpr is called when exiting the functionMultiExpr production.
	ExitFunctionMultiExpr(c *FunctionMultiExprContext)

	// ExitAddSubExpr is called when exiting the addSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitFunctionExpr is called when exiting the functionExpr production.
	ExitFunctionExpr(c *FunctionExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitMulDivExpr is called when exiting the mulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitIteExpr is called when exiting the iteExpr production.
	ExitIteExpr(c *IteExprContext)

	// ExitParExpr is called when exiting the parExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitNumberAtom is called when exiting the numberAtom production.
	ExitNumberAtom(c *NumberAtomContext)

	// ExitBooleanAtom is called when exiting the booleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitPredefinedConstantAtom is called when exiting the predefinedConstantAtom production.
	ExitPredefinedConstantAtom(c *PredefinedConstantAtomContext)

	// ExitIdAtom is called when exiting the idAtom production.
	ExitIdAtom(c *IdAtomContext)
}
