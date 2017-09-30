// Generated from morsecode.g4 by ANTLR 4.7.

package morsecode // morsecode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// morsecodeListener is a complete listener for a parse tree produced by morsecodeParser.
type morsecodeListener interface {
	antlr.ParseTreeListener

	// EnterMorsecode is called when entering the morsecode production.
	EnterMorsecode(c *MorsecodeContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterA is called when entering the a production.
	EnterA(c *AContext)

	// EnterB is called when entering the b production.
	EnterB(c *BContext)

	// EnterC is called when entering the c production.
	EnterC(c *CContext)

	// EnterD is called when entering the d production.
	EnterD(c *DContext)

	// EnterE is called when entering the e production.
	EnterE(c *EContext)

	// EnterF is called when entering the f production.
	EnterF(c *FContext)

	// EnterG is called when entering the g production.
	EnterG(c *GContext)

	// EnterH is called when entering the h production.
	EnterH(c *HContext)

	// EnterI is called when entering the i production.
	EnterI(c *IContext)

	// EnterJ is called when entering the j production.
	EnterJ(c *JContext)

	// EnterK is called when entering the k production.
	EnterK(c *KContext)

	// EnterL is called when entering the l production.
	EnterL(c *LContext)

	// EnterM is called when entering the m production.
	EnterM(c *MContext)

	// EnterN is called when entering the n production.
	EnterN(c *NContext)

	// EnterO is called when entering the o production.
	EnterO(c *OContext)

	// EnterP is called when entering the p production.
	EnterP(c *PContext)

	// EnterQ is called when entering the q production.
	EnterQ(c *QContext)

	// EnterR is called when entering the r production.
	EnterR(c *RContext)

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterT is called when entering the t production.
	EnterT(c *TContext)

	// EnterU is called when entering the u production.
	EnterU(c *UContext)

	// EnterV is called when entering the v production.
	EnterV(c *VContext)

	// EnterW is called when entering the w production.
	EnterW(c *WContext)

	// EnterX is called when entering the x production.
	EnterX(c *XContext)

	// EnterY is called when entering the y production.
	EnterY(c *YContext)

	// EnterZ is called when entering the z production.
	EnterZ(c *ZContext)

	// EnterOne is called when entering the one production.
	EnterOne(c *OneContext)

	// EnterTwo is called when entering the two production.
	EnterTwo(c *TwoContext)

	// EnterThree is called when entering the three production.
	EnterThree(c *ThreeContext)

	// EnterFour is called when entering the four production.
	EnterFour(c *FourContext)

	// EnterFive is called when entering the five production.
	EnterFive(c *FiveContext)

	// EnterSix is called when entering the six production.
	EnterSix(c *SixContext)

	// EnterSeven is called when entering the seven production.
	EnterSeven(c *SevenContext)

	// EnterEight is called when entering the eight production.
	EnterEight(c *EightContext)

	// EnterNine is called when entering the nine production.
	EnterNine(c *NineContext)

	// EnterZero is called when entering the zero production.
	EnterZero(c *ZeroContext)

	// ExitMorsecode is called when exiting the morsecode production.
	ExitMorsecode(c *MorsecodeContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitA is called when exiting the a production.
	ExitA(c *AContext)

	// ExitB is called when exiting the b production.
	ExitB(c *BContext)

	// ExitC is called when exiting the c production.
	ExitC(c *CContext)

	// ExitD is called when exiting the d production.
	ExitD(c *DContext)

	// ExitE is called when exiting the e production.
	ExitE(c *EContext)

	// ExitF is called when exiting the f production.
	ExitF(c *FContext)

	// ExitG is called when exiting the g production.
	ExitG(c *GContext)

	// ExitH is called when exiting the h production.
	ExitH(c *HContext)

	// ExitI is called when exiting the i production.
	ExitI(c *IContext)

	// ExitJ is called when exiting the j production.
	ExitJ(c *JContext)

	// ExitK is called when exiting the k production.
	ExitK(c *KContext)

	// ExitL is called when exiting the l production.
	ExitL(c *LContext)

	// ExitM is called when exiting the m production.
	ExitM(c *MContext)

	// ExitN is called when exiting the n production.
	ExitN(c *NContext)

	// ExitO is called when exiting the o production.
	ExitO(c *OContext)

	// ExitP is called when exiting the p production.
	ExitP(c *PContext)

	// ExitQ is called when exiting the q production.
	ExitQ(c *QContext)

	// ExitR is called when exiting the r production.
	ExitR(c *RContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitT is called when exiting the t production.
	ExitT(c *TContext)

	// ExitU is called when exiting the u production.
	ExitU(c *UContext)

	// ExitV is called when exiting the v production.
	ExitV(c *VContext)

	// ExitW is called when exiting the w production.
	ExitW(c *WContext)

	// ExitX is called when exiting the x production.
	ExitX(c *XContext)

	// ExitY is called when exiting the y production.
	ExitY(c *YContext)

	// ExitZ is called when exiting the z production.
	ExitZ(c *ZContext)

	// ExitOne is called when exiting the one production.
	ExitOne(c *OneContext)

	// ExitTwo is called when exiting the two production.
	ExitTwo(c *TwoContext)

	// ExitThree is called when exiting the three production.
	ExitThree(c *ThreeContext)

	// ExitFour is called when exiting the four production.
	ExitFour(c *FourContext)

	// ExitFive is called when exiting the five production.
	ExitFive(c *FiveContext)

	// ExitSix is called when exiting the six production.
	ExitSix(c *SixContext)

	// ExitSeven is called when exiting the seven production.
	ExitSeven(c *SevenContext)

	// ExitEight is called when exiting the eight production.
	ExitEight(c *EightContext)

	// ExitNine is called when exiting the nine production.
	ExitNine(c *NineContext)

	// ExitZero is called when exiting the zero production.
	ExitZero(c *ZeroContext)
}
