// Generated from morsecode.g4 by ANTLR 4.7.

package morsecode // morsecode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by morsecodeParser.
type morsecodeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by morsecodeParser#morsecode.
	VisitMorsecode(ctx *MorsecodeContext) interface{}

	// Visit a parse tree produced by morsecodeParser#letter.
	VisitLetter(ctx *LetterContext) interface{}

	// Visit a parse tree produced by morsecodeParser#a.
	VisitA(ctx *AContext) interface{}

	// Visit a parse tree produced by morsecodeParser#b.
	VisitB(ctx *BContext) interface{}

	// Visit a parse tree produced by morsecodeParser#c.
	VisitC(ctx *CContext) interface{}

	// Visit a parse tree produced by morsecodeParser#d.
	VisitD(ctx *DContext) interface{}

	// Visit a parse tree produced by morsecodeParser#e.
	VisitE(ctx *EContext) interface{}

	// Visit a parse tree produced by morsecodeParser#f.
	VisitF(ctx *FContext) interface{}

	// Visit a parse tree produced by morsecodeParser#g.
	VisitG(ctx *GContext) interface{}

	// Visit a parse tree produced by morsecodeParser#h.
	VisitH(ctx *HContext) interface{}

	// Visit a parse tree produced by morsecodeParser#i.
	VisitI(ctx *IContext) interface{}

	// Visit a parse tree produced by morsecodeParser#j.
	VisitJ(ctx *JContext) interface{}

	// Visit a parse tree produced by morsecodeParser#k.
	VisitK(ctx *KContext) interface{}

	// Visit a parse tree produced by morsecodeParser#l.
	VisitL(ctx *LContext) interface{}

	// Visit a parse tree produced by morsecodeParser#m.
	VisitM(ctx *MContext) interface{}

	// Visit a parse tree produced by morsecodeParser#n.
	VisitN(ctx *NContext) interface{}

	// Visit a parse tree produced by morsecodeParser#o.
	VisitO(ctx *OContext) interface{}

	// Visit a parse tree produced by morsecodeParser#p.
	VisitP(ctx *PContext) interface{}

	// Visit a parse tree produced by morsecodeParser#q.
	VisitQ(ctx *QContext) interface{}

	// Visit a parse tree produced by morsecodeParser#r.
	VisitR(ctx *RContext) interface{}

	// Visit a parse tree produced by morsecodeParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by morsecodeParser#t.
	VisitT(ctx *TContext) interface{}

	// Visit a parse tree produced by morsecodeParser#u.
	VisitU(ctx *UContext) interface{}

	// Visit a parse tree produced by morsecodeParser#v.
	VisitV(ctx *VContext) interface{}

	// Visit a parse tree produced by morsecodeParser#w.
	VisitW(ctx *WContext) interface{}

	// Visit a parse tree produced by morsecodeParser#x.
	VisitX(ctx *XContext) interface{}

	// Visit a parse tree produced by morsecodeParser#y.
	VisitY(ctx *YContext) interface{}

	// Visit a parse tree produced by morsecodeParser#z.
	VisitZ(ctx *ZContext) interface{}

	// Visit a parse tree produced by morsecodeParser#one.
	VisitOne(ctx *OneContext) interface{}

	// Visit a parse tree produced by morsecodeParser#two.
	VisitTwo(ctx *TwoContext) interface{}

	// Visit a parse tree produced by morsecodeParser#three.
	VisitThree(ctx *ThreeContext) interface{}

	// Visit a parse tree produced by morsecodeParser#four.
	VisitFour(ctx *FourContext) interface{}

	// Visit a parse tree produced by morsecodeParser#five.
	VisitFive(ctx *FiveContext) interface{}

	// Visit a parse tree produced by morsecodeParser#six.
	VisitSix(ctx *SixContext) interface{}

	// Visit a parse tree produced by morsecodeParser#seven.
	VisitSeven(ctx *SevenContext) interface{}

	// Visit a parse tree produced by morsecodeParser#eight.
	VisitEight(ctx *EightContext) interface{}

	// Visit a parse tree produced by morsecodeParser#nine.
	VisitNine(ctx *NineContext) interface{}

	// Visit a parse tree produced by morsecodeParser#zero.
	VisitZero(ctx *ZeroContext) interface{}
}
