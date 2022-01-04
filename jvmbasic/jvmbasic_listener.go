// Code generated from jvmBasic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package jvmbasic // jvmBasic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// jvmBasicListener is a complete listener for a parse tree produced by jvmBasicParser.
type jvmBasicListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterAmperoper is called when entering the amperoper production.
	EnterAmperoper(c *AmperoperContext)

	// EnterLinenumber is called when entering the linenumber production.
	EnterLinenumber(c *LinenumberContext)

	// EnterAmprstmt is called when entering the amprstmt production.
	EnterAmprstmt(c *AmprstmtContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVardecl is called when entering the vardecl production.
	EnterVardecl(c *VardeclContext)

	// EnterPrintstmt1 is called when entering the printstmt1 production.
	EnterPrintstmt1(c *Printstmt1Context)

	// EnterPrintlist is called when entering the printlist production.
	EnterPrintlist(c *PrintlistContext)

	// EnterGetstmt is called when entering the getstmt production.
	EnterGetstmt(c *GetstmtContext)

	// EnterLetstmt is called when entering the letstmt production.
	EnterLetstmt(c *LetstmtContext)

	// EnterVariableassignment is called when entering the variableassignment production.
	EnterVariableassignment(c *VariableassignmentContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// EnterNeq is called when entering the neq production.
	EnterNeq(c *NeqContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterForstmt1 is called when entering the forstmt1 production.
	EnterForstmt1(c *Forstmt1Context)

	// EnterForstmt2 is called when entering the forstmt2 production.
	EnterForstmt2(c *Forstmt2Context)

	// EnterNextstmt is called when entering the nextstmt production.
	EnterNextstmt(c *NextstmtContext)

	// EnterInputstmt is called when entering the inputstmt production.
	EnterInputstmt(c *InputstmtContext)

	// EnterReadstmt is called when entering the readstmt production.
	EnterReadstmt(c *ReadstmtContext)

	// EnterDimstmt is called when entering the dimstmt production.
	EnterDimstmt(c *DimstmtContext)

	// EnterGotostmt is called when entering the gotostmt production.
	EnterGotostmt(c *GotostmtContext)

	// EnterGosubstmt is called when entering the gosubstmt production.
	EnterGosubstmt(c *GosubstmtContext)

	// EnterPokestmt is called when entering the pokestmt production.
	EnterPokestmt(c *PokestmtContext)

	// EnterCallstmt is called when entering the callstmt production.
	EnterCallstmt(c *CallstmtContext)

	// EnterHplotstmt is called when entering the hplotstmt production.
	EnterHplotstmt(c *HplotstmtContext)

	// EnterVplotstmt is called when entering the vplotstmt production.
	EnterVplotstmt(c *VplotstmtContext)

	// EnterPlotstmt is called when entering the plotstmt production.
	EnterPlotstmt(c *PlotstmtContext)

	// EnterOngotostmt is called when entering the ongotostmt production.
	EnterOngotostmt(c *OngotostmtContext)

	// EnterOngosubstmt is called when entering the ongosubstmt production.
	EnterOngosubstmt(c *OngosubstmtContext)

	// EnterVtabstmnt is called when entering the vtabstmnt production.
	EnterVtabstmnt(c *VtabstmntContext)

	// EnterHtabstmnt is called when entering the htabstmnt production.
	EnterHtabstmnt(c *HtabstmntContext)

	// EnterHimemstmt is called when entering the himemstmt production.
	EnterHimemstmt(c *HimemstmtContext)

	// EnterLomemstmt is called when entering the lomemstmt production.
	EnterLomemstmt(c *LomemstmtContext)

	// EnterDatastmt is called when entering the datastmt production.
	EnterDatastmt(c *DatastmtContext)

	// EnterDatum is called when entering the datum production.
	EnterDatum(c *DatumContext)

	// EnterWaitstmt is called when entering the waitstmt production.
	EnterWaitstmt(c *WaitstmtContext)

	// EnterXdrawstmt is called when entering the xdrawstmt production.
	EnterXdrawstmt(c *XdrawstmtContext)

	// EnterDrawstmt is called when entering the drawstmt production.
	EnterDrawstmt(c *DrawstmtContext)

	// EnterDefstmt is called when entering the defstmt production.
	EnterDefstmt(c *DefstmtContext)

	// EnterTabstmt is called when entering the tabstmt production.
	EnterTabstmt(c *TabstmtContext)

	// EnterSpeedstmt is called when entering the speedstmt production.
	EnterSpeedstmt(c *SpeedstmtContext)

	// EnterRotstmt is called when entering the rotstmt production.
	EnterRotstmt(c *RotstmtContext)

	// EnterScalestmt is called when entering the scalestmt production.
	EnterScalestmt(c *ScalestmtContext)

	// EnterColorstmt is called when entering the colorstmt production.
	EnterColorstmt(c *ColorstmtContext)

	// EnterHcolorstmt is called when entering the hcolorstmt production.
	EnterHcolorstmt(c *HcolorstmtContext)

	// EnterHlinstmt is called when entering the hlinstmt production.
	EnterHlinstmt(c *HlinstmtContext)

	// EnterVlinstmt is called when entering the vlinstmt production.
	EnterVlinstmt(c *VlinstmtContext)

	// EnterOnerrstmt is called when entering the onerrstmt production.
	EnterOnerrstmt(c *OnerrstmtContext)

	// EnterPrstmt is called when entering the prstmt production.
	EnterPrstmt(c *PrstmtContext)

	// EnterInstmt is called when entering the instmt production.
	EnterInstmt(c *InstmtContext)

	// EnterStorestmt is called when entering the storestmt production.
	EnterStorestmt(c *StorestmtContext)

	// EnterRecallstmt is called when entering the recallstmt production.
	EnterRecallstmt(c *RecallstmtContext)

	// EnterListstmt is called when entering the liststmt production.
	EnterListstmt(c *ListstmtContext)

	// EnterPopstmt is called when entering the popstmt production.
	EnterPopstmt(c *PopstmtContext)

	// EnterAmptstmt is called when entering the amptstmt production.
	EnterAmptstmt(c *AmptstmtContext)

	// EnterIncludestmt is called when entering the includestmt production.
	EnterIncludestmt(c *IncludestmtContext)

	// EnterEndstmt is called when entering the endstmt production.
	EnterEndstmt(c *EndstmtContext)

	// EnterReturnstmt is called when entering the returnstmt production.
	EnterReturnstmt(c *ReturnstmtContext)

	// EnterRestorestmt is called when entering the restorestmt production.
	EnterRestorestmt(c *RestorestmtContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterFunc_ is called when entering the func_ production.
	EnterFunc_(c *Func_Context)

	// EnterSignExpression is called when entering the signExpression production.
	EnterSignExpression(c *SignExpressionContext)

	// EnterExponentExpression is called when entering the exponentExpression production.
	EnterExponentExpression(c *ExponentExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterAddingExpression is called when entering the addingExpression production.
	EnterAddingExpression(c *AddingExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterVarname is called when entering the varname production.
	EnterVarname(c *VarnameContext)

	// EnterVarsuffix is called when entering the varsuffix production.
	EnterVarsuffix(c *VarsuffixContext)

	// EnterVarlist is called when entering the varlist production.
	EnterVarlist(c *VarlistContext)

	// EnterExprlist is called when entering the exprlist production.
	EnterExprlist(c *ExprlistContext)

	// EnterSqrfunc is called when entering the sqrfunc production.
	EnterSqrfunc(c *SqrfuncContext)

	// EnterChrfunc is called when entering the chrfunc production.
	EnterChrfunc(c *ChrfuncContext)

	// EnterLenfunc is called when entering the lenfunc production.
	EnterLenfunc(c *LenfuncContext)

	// EnterAscfunc is called when entering the ascfunc production.
	EnterAscfunc(c *AscfuncContext)

	// EnterMidfunc is called when entering the midfunc production.
	EnterMidfunc(c *MidfuncContext)

	// EnterPdlfunc is called when entering the pdlfunc production.
	EnterPdlfunc(c *PdlfuncContext)

	// EnterPeekfunc is called when entering the peekfunc production.
	EnterPeekfunc(c *PeekfuncContext)

	// EnterIntfunc is called when entering the intfunc production.
	EnterIntfunc(c *IntfuncContext)

	// EnterSpcfunc is called when entering the spcfunc production.
	EnterSpcfunc(c *SpcfuncContext)

	// EnterFrefunc is called when entering the frefunc production.
	EnterFrefunc(c *FrefuncContext)

	// EnterPosfunc is called when entering the posfunc production.
	EnterPosfunc(c *PosfuncContext)

	// EnterUsrfunc is called when entering the usrfunc production.
	EnterUsrfunc(c *UsrfuncContext)

	// EnterLeftfunc is called when entering the leftfunc production.
	EnterLeftfunc(c *LeftfuncContext)

	// EnterRightfunc is called when entering the rightfunc production.
	EnterRightfunc(c *RightfuncContext)

	// EnterStrfunc is called when entering the strfunc production.
	EnterStrfunc(c *StrfuncContext)

	// EnterFnfunc is called when entering the fnfunc production.
	EnterFnfunc(c *FnfuncContext)

	// EnterValfunc is called when entering the valfunc production.
	EnterValfunc(c *ValfuncContext)

	// EnterScrnfunc is called when entering the scrnfunc production.
	EnterScrnfunc(c *ScrnfuncContext)

	// EnterSinfunc is called when entering the sinfunc production.
	EnterSinfunc(c *SinfuncContext)

	// EnterCosfunc is called when entering the cosfunc production.
	EnterCosfunc(c *CosfuncContext)

	// EnterTanfunc is called when entering the tanfunc production.
	EnterTanfunc(c *TanfuncContext)

	// EnterAtnfunc is called when entering the atnfunc production.
	EnterAtnfunc(c *AtnfuncContext)

	// EnterRndfunc is called when entering the rndfunc production.
	EnterRndfunc(c *RndfuncContext)

	// EnterSgnfunc is called when entering the sgnfunc production.
	EnterSgnfunc(c *SgnfuncContext)

	// EnterExpfunc is called when entering the expfunc production.
	EnterExpfunc(c *ExpfuncContext)

	// EnterLogfunc is called when entering the logfunc production.
	EnterLogfunc(c *LogfuncContext)

	// EnterAbsfunc is called when entering the absfunc production.
	EnterAbsfunc(c *AbsfuncContext)

	// EnterTabfunc is called when entering the tabfunc production.
	EnterTabfunc(c *TabfuncContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitAmperoper is called when exiting the amperoper production.
	ExitAmperoper(c *AmperoperContext)

	// ExitLinenumber is called when exiting the linenumber production.
	ExitLinenumber(c *LinenumberContext)

	// ExitAmprstmt is called when exiting the amprstmt production.
	ExitAmprstmt(c *AmprstmtContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVardecl is called when exiting the vardecl production.
	ExitVardecl(c *VardeclContext)

	// ExitPrintstmt1 is called when exiting the printstmt1 production.
	ExitPrintstmt1(c *Printstmt1Context)

	// ExitPrintlist is called when exiting the printlist production.
	ExitPrintlist(c *PrintlistContext)

	// ExitGetstmt is called when exiting the getstmt production.
	ExitGetstmt(c *GetstmtContext)

	// ExitLetstmt is called when exiting the letstmt production.
	ExitLetstmt(c *LetstmtContext)

	// ExitVariableassignment is called when exiting the variableassignment production.
	ExitVariableassignment(c *VariableassignmentContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)

	// ExitNeq is called when exiting the neq production.
	ExitNeq(c *NeqContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitForstmt1 is called when exiting the forstmt1 production.
	ExitForstmt1(c *Forstmt1Context)

	// ExitForstmt2 is called when exiting the forstmt2 production.
	ExitForstmt2(c *Forstmt2Context)

	// ExitNextstmt is called when exiting the nextstmt production.
	ExitNextstmt(c *NextstmtContext)

	// ExitInputstmt is called when exiting the inputstmt production.
	ExitInputstmt(c *InputstmtContext)

	// ExitReadstmt is called when exiting the readstmt production.
	ExitReadstmt(c *ReadstmtContext)

	// ExitDimstmt is called when exiting the dimstmt production.
	ExitDimstmt(c *DimstmtContext)

	// ExitGotostmt is called when exiting the gotostmt production.
	ExitGotostmt(c *GotostmtContext)

	// ExitGosubstmt is called when exiting the gosubstmt production.
	ExitGosubstmt(c *GosubstmtContext)

	// ExitPokestmt is called when exiting the pokestmt production.
	ExitPokestmt(c *PokestmtContext)

	// ExitCallstmt is called when exiting the callstmt production.
	ExitCallstmt(c *CallstmtContext)

	// ExitHplotstmt is called when exiting the hplotstmt production.
	ExitHplotstmt(c *HplotstmtContext)

	// ExitVplotstmt is called when exiting the vplotstmt production.
	ExitVplotstmt(c *VplotstmtContext)

	// ExitPlotstmt is called when exiting the plotstmt production.
	ExitPlotstmt(c *PlotstmtContext)

	// ExitOngotostmt is called when exiting the ongotostmt production.
	ExitOngotostmt(c *OngotostmtContext)

	// ExitOngosubstmt is called when exiting the ongosubstmt production.
	ExitOngosubstmt(c *OngosubstmtContext)

	// ExitVtabstmnt is called when exiting the vtabstmnt production.
	ExitVtabstmnt(c *VtabstmntContext)

	// ExitHtabstmnt is called when exiting the htabstmnt production.
	ExitHtabstmnt(c *HtabstmntContext)

	// ExitHimemstmt is called when exiting the himemstmt production.
	ExitHimemstmt(c *HimemstmtContext)

	// ExitLomemstmt is called when exiting the lomemstmt production.
	ExitLomemstmt(c *LomemstmtContext)

	// ExitDatastmt is called when exiting the datastmt production.
	ExitDatastmt(c *DatastmtContext)

	// ExitDatum is called when exiting the datum production.
	ExitDatum(c *DatumContext)

	// ExitWaitstmt is called when exiting the waitstmt production.
	ExitWaitstmt(c *WaitstmtContext)

	// ExitXdrawstmt is called when exiting the xdrawstmt production.
	ExitXdrawstmt(c *XdrawstmtContext)

	// ExitDrawstmt is called when exiting the drawstmt production.
	ExitDrawstmt(c *DrawstmtContext)

	// ExitDefstmt is called when exiting the defstmt production.
	ExitDefstmt(c *DefstmtContext)

	// ExitTabstmt is called when exiting the tabstmt production.
	ExitTabstmt(c *TabstmtContext)

	// ExitSpeedstmt is called when exiting the speedstmt production.
	ExitSpeedstmt(c *SpeedstmtContext)

	// ExitRotstmt is called when exiting the rotstmt production.
	ExitRotstmt(c *RotstmtContext)

	// ExitScalestmt is called when exiting the scalestmt production.
	ExitScalestmt(c *ScalestmtContext)

	// ExitColorstmt is called when exiting the colorstmt production.
	ExitColorstmt(c *ColorstmtContext)

	// ExitHcolorstmt is called when exiting the hcolorstmt production.
	ExitHcolorstmt(c *HcolorstmtContext)

	// ExitHlinstmt is called when exiting the hlinstmt production.
	ExitHlinstmt(c *HlinstmtContext)

	// ExitVlinstmt is called when exiting the vlinstmt production.
	ExitVlinstmt(c *VlinstmtContext)

	// ExitOnerrstmt is called when exiting the onerrstmt production.
	ExitOnerrstmt(c *OnerrstmtContext)

	// ExitPrstmt is called when exiting the prstmt production.
	ExitPrstmt(c *PrstmtContext)

	// ExitInstmt is called when exiting the instmt production.
	ExitInstmt(c *InstmtContext)

	// ExitStorestmt is called when exiting the storestmt production.
	ExitStorestmt(c *StorestmtContext)

	// ExitRecallstmt is called when exiting the recallstmt production.
	ExitRecallstmt(c *RecallstmtContext)

	// ExitListstmt is called when exiting the liststmt production.
	ExitListstmt(c *ListstmtContext)

	// ExitPopstmt is called when exiting the popstmt production.
	ExitPopstmt(c *PopstmtContext)

	// ExitAmptstmt is called when exiting the amptstmt production.
	ExitAmptstmt(c *AmptstmtContext)

	// ExitIncludestmt is called when exiting the includestmt production.
	ExitIncludestmt(c *IncludestmtContext)

	// ExitEndstmt is called when exiting the endstmt production.
	ExitEndstmt(c *EndstmtContext)

	// ExitReturnstmt is called when exiting the returnstmt production.
	ExitReturnstmt(c *ReturnstmtContext)

	// ExitRestorestmt is called when exiting the restorestmt production.
	ExitRestorestmt(c *RestorestmtContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitFunc_ is called when exiting the func_ production.
	ExitFunc_(c *Func_Context)

	// ExitSignExpression is called when exiting the signExpression production.
	ExitSignExpression(c *SignExpressionContext)

	// ExitExponentExpression is called when exiting the exponentExpression production.
	ExitExponentExpression(c *ExponentExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitAddingExpression is called when exiting the addingExpression production.
	ExitAddingExpression(c *AddingExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitVarname is called when exiting the varname production.
	ExitVarname(c *VarnameContext)

	// ExitVarsuffix is called when exiting the varsuffix production.
	ExitVarsuffix(c *VarsuffixContext)

	// ExitVarlist is called when exiting the varlist production.
	ExitVarlist(c *VarlistContext)

	// ExitExprlist is called when exiting the exprlist production.
	ExitExprlist(c *ExprlistContext)

	// ExitSqrfunc is called when exiting the sqrfunc production.
	ExitSqrfunc(c *SqrfuncContext)

	// ExitChrfunc is called when exiting the chrfunc production.
	ExitChrfunc(c *ChrfuncContext)

	// ExitLenfunc is called when exiting the lenfunc production.
	ExitLenfunc(c *LenfuncContext)

	// ExitAscfunc is called when exiting the ascfunc production.
	ExitAscfunc(c *AscfuncContext)

	// ExitMidfunc is called when exiting the midfunc production.
	ExitMidfunc(c *MidfuncContext)

	// ExitPdlfunc is called when exiting the pdlfunc production.
	ExitPdlfunc(c *PdlfuncContext)

	// ExitPeekfunc is called when exiting the peekfunc production.
	ExitPeekfunc(c *PeekfuncContext)

	// ExitIntfunc is called when exiting the intfunc production.
	ExitIntfunc(c *IntfuncContext)

	// ExitSpcfunc is called when exiting the spcfunc production.
	ExitSpcfunc(c *SpcfuncContext)

	// ExitFrefunc is called when exiting the frefunc production.
	ExitFrefunc(c *FrefuncContext)

	// ExitPosfunc is called when exiting the posfunc production.
	ExitPosfunc(c *PosfuncContext)

	// ExitUsrfunc is called when exiting the usrfunc production.
	ExitUsrfunc(c *UsrfuncContext)

	// ExitLeftfunc is called when exiting the leftfunc production.
	ExitLeftfunc(c *LeftfuncContext)

	// ExitRightfunc is called when exiting the rightfunc production.
	ExitRightfunc(c *RightfuncContext)

	// ExitStrfunc is called when exiting the strfunc production.
	ExitStrfunc(c *StrfuncContext)

	// ExitFnfunc is called when exiting the fnfunc production.
	ExitFnfunc(c *FnfuncContext)

	// ExitValfunc is called when exiting the valfunc production.
	ExitValfunc(c *ValfuncContext)

	// ExitScrnfunc is called when exiting the scrnfunc production.
	ExitScrnfunc(c *ScrnfuncContext)

	// ExitSinfunc is called when exiting the sinfunc production.
	ExitSinfunc(c *SinfuncContext)

	// ExitCosfunc is called when exiting the cosfunc production.
	ExitCosfunc(c *CosfuncContext)

	// ExitTanfunc is called when exiting the tanfunc production.
	ExitTanfunc(c *TanfuncContext)

	// ExitAtnfunc is called when exiting the atnfunc production.
	ExitAtnfunc(c *AtnfuncContext)

	// ExitRndfunc is called when exiting the rndfunc production.
	ExitRndfunc(c *RndfuncContext)

	// ExitSgnfunc is called when exiting the sgnfunc production.
	ExitSgnfunc(c *SgnfuncContext)

	// ExitExpfunc is called when exiting the expfunc production.
	ExitExpfunc(c *ExpfuncContext)

	// ExitLogfunc is called when exiting the logfunc production.
	ExitLogfunc(c *LogfuncContext)

	// ExitAbsfunc is called when exiting the absfunc production.
	ExitAbsfunc(c *AbsfuncContext)

	// ExitTabfunc is called when exiting the tabfunc production.
	ExitTabfunc(c *TabfuncContext)
}
