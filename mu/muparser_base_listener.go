// Code generated from MuParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mu // MuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMuParserListener is a complete listener for a parse tree produced by MuParserParser.
type BaseMuParserListener struct{}

var _ MuParserListener = &BaseMuParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMuParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMuParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMuParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMuParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgExpr is called when production progExpr is entered.
func (s *BaseMuParserListener) EnterProgExpr(ctx *ProgExprContext) {}

// ExitProgExpr is called when production progExpr is exited.
func (s *BaseMuParserListener) ExitProgExpr(ctx *ProgExprContext) {}

// EnterFunctionMultiExpr is called when production functionMultiExpr is entered.
func (s *BaseMuParserListener) EnterFunctionMultiExpr(ctx *FunctionMultiExprContext) {}

// ExitFunctionMultiExpr is called when production functionMultiExpr is exited.
func (s *BaseMuParserListener) ExitFunctionMultiExpr(ctx *FunctionMultiExprContext) {}

// EnterAddSubExpr is called when production addSubExpr is entered.
func (s *BaseMuParserListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production addSubExpr is exited.
func (s *BaseMuParserListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseMuParserListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseMuParserListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseMuParserListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseMuParserListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BaseMuParserListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BaseMuParserListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterFunctionExpr is called when production functionExpr is entered.
func (s *BaseMuParserListener) EnterFunctionExpr(ctx *FunctionExprContext) {}

// ExitFunctionExpr is called when production functionExpr is exited.
func (s *BaseMuParserListener) ExitFunctionExpr(ctx *FunctionExprContext) {}

// EnterUnaryMinusExpr is called when production unaryMinusExpr is entered.
func (s *BaseMuParserListener) EnterUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// ExitUnaryMinusExpr is called when production unaryMinusExpr is exited.
func (s *BaseMuParserListener) ExitUnaryMinusExpr(ctx *UnaryMinusExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseMuParserListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseMuParserListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseMuParserListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseMuParserListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterMulDivExpr is called when production mulDivExpr is entered.
func (s *BaseMuParserListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production mulDivExpr is exited.
func (s *BaseMuParserListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BaseMuParserListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BaseMuParserListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseMuParserListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseMuParserListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterIteExpr is called when production iteExpr is entered.
func (s *BaseMuParserListener) EnterIteExpr(ctx *IteExprContext) {}

// ExitIteExpr is called when production iteExpr is exited.
func (s *BaseMuParserListener) ExitIteExpr(ctx *IteExprContext) {}

// EnterParExpr is called when production parExpr is entered.
func (s *BaseMuParserListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production parExpr is exited.
func (s *BaseMuParserListener) ExitParExpr(ctx *ParExprContext) {}

// EnterNumberAtom is called when production numberAtom is entered.
func (s *BaseMuParserListener) EnterNumberAtom(ctx *NumberAtomContext) {}

// ExitNumberAtom is called when production numberAtom is exited.
func (s *BaseMuParserListener) ExitNumberAtom(ctx *NumberAtomContext) {}

// EnterBooleanAtom is called when production booleanAtom is entered.
func (s *BaseMuParserListener) EnterBooleanAtom(ctx *BooleanAtomContext) {}

// ExitBooleanAtom is called when production booleanAtom is exited.
func (s *BaseMuParserListener) ExitBooleanAtom(ctx *BooleanAtomContext) {}

// EnterPredefinedConstantAtom is called when production predefinedConstantAtom is entered.
func (s *BaseMuParserListener) EnterPredefinedConstantAtom(ctx *PredefinedConstantAtomContext) {}

// ExitPredefinedConstantAtom is called when production predefinedConstantAtom is exited.
func (s *BaseMuParserListener) ExitPredefinedConstantAtom(ctx *PredefinedConstantAtomContext) {}

// EnterIdAtom is called when production idAtom is entered.
func (s *BaseMuParserListener) EnterIdAtom(ctx *IdAtomContext) {}

// ExitIdAtom is called when production idAtom is exited.
func (s *BaseMuParserListener) ExitIdAtom(ctx *IdAtomContext) {}
