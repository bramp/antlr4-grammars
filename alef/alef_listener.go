// Code generated from alef.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alef // alef
import "github.com/antlr/antlr4/runtime/Go/antlr"

// alefListener is a complete listener for a parse tree produced by alefParser.
type alefListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDecllist is called when entering the decllist production.
	EnterDecllist(c *DecllistContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterZargs is called when entering the zargs production.
	EnterZargs(c *ZargsContext)

	// EnterZtname is called when entering the ztname production.
	EnterZtname(c *ZtnameContext)

	// EnterAdtfunc is called when entering the adtfunc production.
	EnterAdtfunc(c *AdtfuncContext)

	// EnterTypespec is called when entering the typespec production.
	EnterTypespec(c *TypespecContext)

	// EnterZtag is called when entering the ztag production.
	EnterZtag(c *ZtagContext)

	// EnterZpolytype is called when entering the zpolytype production.
	EnterZpolytype(c *ZpolytypeContext)

	// EnterPolytype is called when entering the polytype production.
	EnterPolytype(c *PolytypeContext)

	// EnterSetlist is called when entering the setlist production.
	EnterSetlist(c *SetlistContext)

	// EnterSname is called when entering the sname production.
	EnterSname(c *SnameContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterMemberlist is called when entering the memberlist production.
	EnterMemberlist(c *MemberlistContext)

	// EnterVardecllist is called when entering the vardecllist production.
	EnterVardecllist(c *VardecllistContext)

	// EnterIvardecl is called when entering the ivardecl production.
	EnterIvardecl(c *IvardeclContext)

	// EnterZinit is called when entering the zinit production.
	EnterZinit(c *ZinitContext)

	// EnterZelist is called when entering the zelist production.
	EnterZelist(c *ZelistContext)

	// EnterVardecl is called when entering the vardecl production.
	EnterVardecl(c *VardeclContext)

	// EnterArrayspec is called when entering the arrayspec production.
	EnterArrayspec(c *ArrayspecContext)

	// EnterIndsp is called when entering the indsp production.
	EnterIndsp(c *IndspContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterArglistp is called when entering the arglistp production.
	EnterArglistp(c *ArglistpContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterTuplearg is called when entering the tuplearg production.
	EnterTuplearg(c *TupleargContext)

	// EnterAutolist is called when entering the autolist production.
	EnterAutolist(c *AutolistContext)

	// EnterAutodecl is called when entering the autodecl production.
	EnterAutodecl(c *AutodeclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSlist is called when entering the slist production.
	EnterSlist(c *SlistContext)

	// EnterTbody is called when entering the tbody production.
	EnterTbody(c *TbodyContext)

	// EnterCtlist is called when entering the ctlist production.
	EnterCtlist(c *CtlistContext)

	// EnterTcase is called when entering the tcase production.
	EnterTcase(c *TcaseContext)

	// EnterCbody is called when entering the cbody production.
	EnterCbody(c *CbodyContext)

	// EnterClist is called when entering the clist production.
	EnterClist(c *ClistContext)

	// EnterCase_ is called when entering the case_ production.
	EnterCase_(c *Case_Context)

	// EnterRbody is called when entering the rbody production.
	EnterRbody(c *RbodyContext)

	// EnterZlab is called when entering the zlab production.
	EnterZlab(c *ZlabContext)

	// EnterStmnt is called when entering the stmnt production.
	EnterStmnt(c *StmntContext)

	// EnterInfo is called when entering the info production.
	EnterInfo(c *InfoContext)

	// EnterNlstmnt is called when entering the nlstmnt production.
	EnterNlstmnt(c *NlstmntContext)

	// EnterZconst is called when entering the zconst production.
	EnterZconst(c *ZconstContext)

	// EnterZexpr is called when entering the zexpr production.
	EnterZexpr(c *ZexprContext)

	// EnterExpr_ is called when entering the expr_ production.
	EnterExpr_(c *Expr_Context)

	// EnterCastexpr is called when entering the castexpr production.
	EnterCastexpr(c *CastexprContext)

	// EnterTypecast is called when entering the typecast production.
	EnterTypecast(c *TypecastContext)

	// EnterMonexpr is called when entering the monexpr production.
	EnterMonexpr(c *MonexprContext)

	// EnterZtelist is called when entering the ztelist production.
	EnterZtelist(c *ZtelistContext)

	// EnterTelist is called when entering the telist production.
	EnterTelist(c *TelistContext)

	// EnterTcomp is called when entering the tcomp production.
	EnterTcomp(c *TcompContext)

	// EnterTerm_ is called when entering the term_ production.
	EnterTerm_(c *Term_Context)

	// EnterStag is called when entering the stag production.
	EnterStag(c *StagContext)

	// EnterZarlist is called when entering the zarlist production.
	EnterZarlist(c *ZarlistContext)

	// EnterElist is called when entering the elist production.
	EnterElist(c *ElistContext)

	// EnterTlist is called when entering the tlist production.
	EnterTlist(c *TlistContext)

	// EnterTname is called when entering the tname production.
	EnterTname(c *TnameContext)

	// EnterVariant is called when entering the variant production.
	EnterVariant(c *VariantContext)

	// EnterXtname is called when entering the xtname production.
	EnterXtname(c *XtnameContext)

	// EnterBufdim is called when entering the bufdim production.
	EnterBufdim(c *BufdimContext)

	// EnterSclass is called when entering the sclass production.
	EnterSclass(c *SclassContext)

	// EnterTypename is called when entering the typename production.
	EnterTypename(c *TypenameContext)

	// EnterEnum_member is called when entering the enum_member production.
	EnterEnum_member(c *Enum_memberContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDecllist is called when exiting the decllist production.
	ExitDecllist(c *DecllistContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitZargs is called when exiting the zargs production.
	ExitZargs(c *ZargsContext)

	// ExitZtname is called when exiting the ztname production.
	ExitZtname(c *ZtnameContext)

	// ExitAdtfunc is called when exiting the adtfunc production.
	ExitAdtfunc(c *AdtfuncContext)

	// ExitTypespec is called when exiting the typespec production.
	ExitTypespec(c *TypespecContext)

	// ExitZtag is called when exiting the ztag production.
	ExitZtag(c *ZtagContext)

	// ExitZpolytype is called when exiting the zpolytype production.
	ExitZpolytype(c *ZpolytypeContext)

	// ExitPolytype is called when exiting the polytype production.
	ExitPolytype(c *PolytypeContext)

	// ExitSetlist is called when exiting the setlist production.
	ExitSetlist(c *SetlistContext)

	// ExitSname is called when exiting the sname production.
	ExitSname(c *SnameContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitMemberlist is called when exiting the memberlist production.
	ExitMemberlist(c *MemberlistContext)

	// ExitVardecllist is called when exiting the vardecllist production.
	ExitVardecllist(c *VardecllistContext)

	// ExitIvardecl is called when exiting the ivardecl production.
	ExitIvardecl(c *IvardeclContext)

	// ExitZinit is called when exiting the zinit production.
	ExitZinit(c *ZinitContext)

	// ExitZelist is called when exiting the zelist production.
	ExitZelist(c *ZelistContext)

	// ExitVardecl is called when exiting the vardecl production.
	ExitVardecl(c *VardeclContext)

	// ExitArrayspec is called when exiting the arrayspec production.
	ExitArrayspec(c *ArrayspecContext)

	// ExitIndsp is called when exiting the indsp production.
	ExitIndsp(c *IndspContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitArglistp is called when exiting the arglistp production.
	ExitArglistp(c *ArglistpContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitTuplearg is called when exiting the tuplearg production.
	ExitTuplearg(c *TupleargContext)

	// ExitAutolist is called when exiting the autolist production.
	ExitAutolist(c *AutolistContext)

	// ExitAutodecl is called when exiting the autodecl production.
	ExitAutodecl(c *AutodeclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSlist is called when exiting the slist production.
	ExitSlist(c *SlistContext)

	// ExitTbody is called when exiting the tbody production.
	ExitTbody(c *TbodyContext)

	// ExitCtlist is called when exiting the ctlist production.
	ExitCtlist(c *CtlistContext)

	// ExitTcase is called when exiting the tcase production.
	ExitTcase(c *TcaseContext)

	// ExitCbody is called when exiting the cbody production.
	ExitCbody(c *CbodyContext)

	// ExitClist is called when exiting the clist production.
	ExitClist(c *ClistContext)

	// ExitCase_ is called when exiting the case_ production.
	ExitCase_(c *Case_Context)

	// ExitRbody is called when exiting the rbody production.
	ExitRbody(c *RbodyContext)

	// ExitZlab is called when exiting the zlab production.
	ExitZlab(c *ZlabContext)

	// ExitStmnt is called when exiting the stmnt production.
	ExitStmnt(c *StmntContext)

	// ExitInfo is called when exiting the info production.
	ExitInfo(c *InfoContext)

	// ExitNlstmnt is called when exiting the nlstmnt production.
	ExitNlstmnt(c *NlstmntContext)

	// ExitZconst is called when exiting the zconst production.
	ExitZconst(c *ZconstContext)

	// ExitZexpr is called when exiting the zexpr production.
	ExitZexpr(c *ZexprContext)

	// ExitExpr_ is called when exiting the expr_ production.
	ExitExpr_(c *Expr_Context)

	// ExitCastexpr is called when exiting the castexpr production.
	ExitCastexpr(c *CastexprContext)

	// ExitTypecast is called when exiting the typecast production.
	ExitTypecast(c *TypecastContext)

	// ExitMonexpr is called when exiting the monexpr production.
	ExitMonexpr(c *MonexprContext)

	// ExitZtelist is called when exiting the ztelist production.
	ExitZtelist(c *ZtelistContext)

	// ExitTelist is called when exiting the telist production.
	ExitTelist(c *TelistContext)

	// ExitTcomp is called when exiting the tcomp production.
	ExitTcomp(c *TcompContext)

	// ExitTerm_ is called when exiting the term_ production.
	ExitTerm_(c *Term_Context)

	// ExitStag is called when exiting the stag production.
	ExitStag(c *StagContext)

	// ExitZarlist is called when exiting the zarlist production.
	ExitZarlist(c *ZarlistContext)

	// ExitElist is called when exiting the elist production.
	ExitElist(c *ElistContext)

	// ExitTlist is called when exiting the tlist production.
	ExitTlist(c *TlistContext)

	// ExitTname is called when exiting the tname production.
	ExitTname(c *TnameContext)

	// ExitVariant is called when exiting the variant production.
	ExitVariant(c *VariantContext)

	// ExitXtname is called when exiting the xtname production.
	ExitXtname(c *XtnameContext)

	// ExitBufdim is called when exiting the bufdim production.
	ExitBufdim(c *BufdimContext)

	// ExitSclass is called when exiting the sclass production.
	ExitSclass(c *SclassContext)

	// ExitTypename is called when exiting the typename production.
	ExitTypename(c *TypenameContext)

	// ExitEnum_member is called when exiting the enum_member production.
	ExitEnum_member(c *Enum_memberContext)
}
