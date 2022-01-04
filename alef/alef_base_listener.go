// Code generated from alef.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alef // alef
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasealefListener is a complete listener for a parse tree produced by alefParser.
type BasealefListener struct{}

var _ alefListener = &BasealefListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasealefListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasealefListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasealefListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasealefListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasealefListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasealefListener) ExitProgram(ctx *ProgramContext) {}

// EnterDecllist is called when production decllist is entered.
func (s *BasealefListener) EnterDecllist(ctx *DecllistContext) {}

// ExitDecllist is called when production decllist is exited.
func (s *BasealefListener) ExitDecllist(ctx *DecllistContext) {}

// EnterDecl is called when production decl is entered.
func (s *BasealefListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BasealefListener) ExitDecl(ctx *DeclContext) {}

// EnterZargs is called when production zargs is entered.
func (s *BasealefListener) EnterZargs(ctx *ZargsContext) {}

// ExitZargs is called when production zargs is exited.
func (s *BasealefListener) ExitZargs(ctx *ZargsContext) {}

// EnterZtname is called when production ztname is entered.
func (s *BasealefListener) EnterZtname(ctx *ZtnameContext) {}

// ExitZtname is called when production ztname is exited.
func (s *BasealefListener) ExitZtname(ctx *ZtnameContext) {}

// EnterAdtfunc is called when production adtfunc is entered.
func (s *BasealefListener) EnterAdtfunc(ctx *AdtfuncContext) {}

// ExitAdtfunc is called when production adtfunc is exited.
func (s *BasealefListener) ExitAdtfunc(ctx *AdtfuncContext) {}

// EnterTypespec is called when production typespec is entered.
func (s *BasealefListener) EnterTypespec(ctx *TypespecContext) {}

// ExitTypespec is called when production typespec is exited.
func (s *BasealefListener) ExitTypespec(ctx *TypespecContext) {}

// EnterZtag is called when production ztag is entered.
func (s *BasealefListener) EnterZtag(ctx *ZtagContext) {}

// ExitZtag is called when production ztag is exited.
func (s *BasealefListener) ExitZtag(ctx *ZtagContext) {}

// EnterZpolytype is called when production zpolytype is entered.
func (s *BasealefListener) EnterZpolytype(ctx *ZpolytypeContext) {}

// ExitZpolytype is called when production zpolytype is exited.
func (s *BasealefListener) ExitZpolytype(ctx *ZpolytypeContext) {}

// EnterPolytype is called when production polytype is entered.
func (s *BasealefListener) EnterPolytype(ctx *PolytypeContext) {}

// ExitPolytype is called when production polytype is exited.
func (s *BasealefListener) ExitPolytype(ctx *PolytypeContext) {}

// EnterSetlist is called when production setlist is entered.
func (s *BasealefListener) EnterSetlist(ctx *SetlistContext) {}

// ExitSetlist is called when production setlist is exited.
func (s *BasealefListener) ExitSetlist(ctx *SetlistContext) {}

// EnterSname is called when production sname is entered.
func (s *BasealefListener) EnterSname(ctx *SnameContext) {}

// ExitSname is called when production sname is exited.
func (s *BasealefListener) ExitSname(ctx *SnameContext) {}

// EnterName is called when production name is entered.
func (s *BasealefListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasealefListener) ExitName(ctx *NameContext) {}

// EnterMemberlist is called when production memberlist is entered.
func (s *BasealefListener) EnterMemberlist(ctx *MemberlistContext) {}

// ExitMemberlist is called when production memberlist is exited.
func (s *BasealefListener) ExitMemberlist(ctx *MemberlistContext) {}

// EnterVardecllist is called when production vardecllist is entered.
func (s *BasealefListener) EnterVardecllist(ctx *VardecllistContext) {}

// ExitVardecllist is called when production vardecllist is exited.
func (s *BasealefListener) ExitVardecllist(ctx *VardecllistContext) {}

// EnterIvardecl is called when production ivardecl is entered.
func (s *BasealefListener) EnterIvardecl(ctx *IvardeclContext) {}

// ExitIvardecl is called when production ivardecl is exited.
func (s *BasealefListener) ExitIvardecl(ctx *IvardeclContext) {}

// EnterZinit is called when production zinit is entered.
func (s *BasealefListener) EnterZinit(ctx *ZinitContext) {}

// ExitZinit is called when production zinit is exited.
func (s *BasealefListener) ExitZinit(ctx *ZinitContext) {}

// EnterZelist is called when production zelist is entered.
func (s *BasealefListener) EnterZelist(ctx *ZelistContext) {}

// ExitZelist is called when production zelist is exited.
func (s *BasealefListener) ExitZelist(ctx *ZelistContext) {}

// EnterVardecl is called when production vardecl is entered.
func (s *BasealefListener) EnterVardecl(ctx *VardeclContext) {}

// ExitVardecl is called when production vardecl is exited.
func (s *BasealefListener) ExitVardecl(ctx *VardeclContext) {}

// EnterArrayspec is called when production arrayspec is entered.
func (s *BasealefListener) EnterArrayspec(ctx *ArrayspecContext) {}

// ExitArrayspec is called when production arrayspec is exited.
func (s *BasealefListener) ExitArrayspec(ctx *ArrayspecContext) {}

// EnterIndsp is called when production indsp is entered.
func (s *BasealefListener) EnterIndsp(ctx *IndspContext) {}

// ExitIndsp is called when production indsp is exited.
func (s *BasealefListener) ExitIndsp(ctx *IndspContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BasealefListener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BasealefListener) ExitArglist(ctx *ArglistContext) {}

// EnterArglistp is called when production arglistp is entered.
func (s *BasealefListener) EnterArglistp(ctx *ArglistpContext) {}

// ExitArglistp is called when production arglistp is exited.
func (s *BasealefListener) ExitArglistp(ctx *ArglistpContext) {}

// EnterArg is called when production arg is entered.
func (s *BasealefListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BasealefListener) ExitArg(ctx *ArgContext) {}

// EnterTuplearg is called when production tuplearg is entered.
func (s *BasealefListener) EnterTuplearg(ctx *TupleargContext) {}

// ExitTuplearg is called when production tuplearg is exited.
func (s *BasealefListener) ExitTuplearg(ctx *TupleargContext) {}

// EnterAutolist is called when production autolist is entered.
func (s *BasealefListener) EnterAutolist(ctx *AutolistContext) {}

// ExitAutolist is called when production autolist is exited.
func (s *BasealefListener) ExitAutolist(ctx *AutolistContext) {}

// EnterAutodecl is called when production autodecl is entered.
func (s *BasealefListener) EnterAutodecl(ctx *AutodeclContext) {}

// ExitAutodecl is called when production autodecl is exited.
func (s *BasealefListener) ExitAutodecl(ctx *AutodeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BasealefListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasealefListener) ExitBlock(ctx *BlockContext) {}

// EnterSlist is called when production slist is entered.
func (s *BasealefListener) EnterSlist(ctx *SlistContext) {}

// ExitSlist is called when production slist is exited.
func (s *BasealefListener) ExitSlist(ctx *SlistContext) {}

// EnterTbody is called when production tbody is entered.
func (s *BasealefListener) EnterTbody(ctx *TbodyContext) {}

// ExitTbody is called when production tbody is exited.
func (s *BasealefListener) ExitTbody(ctx *TbodyContext) {}

// EnterCtlist is called when production ctlist is entered.
func (s *BasealefListener) EnterCtlist(ctx *CtlistContext) {}

// ExitCtlist is called when production ctlist is exited.
func (s *BasealefListener) ExitCtlist(ctx *CtlistContext) {}

// EnterTcase is called when production tcase is entered.
func (s *BasealefListener) EnterTcase(ctx *TcaseContext) {}

// ExitTcase is called when production tcase is exited.
func (s *BasealefListener) ExitTcase(ctx *TcaseContext) {}

// EnterCbody is called when production cbody is entered.
func (s *BasealefListener) EnterCbody(ctx *CbodyContext) {}

// ExitCbody is called when production cbody is exited.
func (s *BasealefListener) ExitCbody(ctx *CbodyContext) {}

// EnterClist is called when production clist is entered.
func (s *BasealefListener) EnterClist(ctx *ClistContext) {}

// ExitClist is called when production clist is exited.
func (s *BasealefListener) ExitClist(ctx *ClistContext) {}

// EnterCase_ is called when production case_ is entered.
func (s *BasealefListener) EnterCase_(ctx *Case_Context) {}

// ExitCase_ is called when production case_ is exited.
func (s *BasealefListener) ExitCase_(ctx *Case_Context) {}

// EnterRbody is called when production rbody is entered.
func (s *BasealefListener) EnterRbody(ctx *RbodyContext) {}

// ExitRbody is called when production rbody is exited.
func (s *BasealefListener) ExitRbody(ctx *RbodyContext) {}

// EnterZlab is called when production zlab is entered.
func (s *BasealefListener) EnterZlab(ctx *ZlabContext) {}

// ExitZlab is called when production zlab is exited.
func (s *BasealefListener) ExitZlab(ctx *ZlabContext) {}

// EnterStmnt is called when production stmnt is entered.
func (s *BasealefListener) EnterStmnt(ctx *StmntContext) {}

// ExitStmnt is called when production stmnt is exited.
func (s *BasealefListener) ExitStmnt(ctx *StmntContext) {}

// EnterInfo is called when production info is entered.
func (s *BasealefListener) EnterInfo(ctx *InfoContext) {}

// ExitInfo is called when production info is exited.
func (s *BasealefListener) ExitInfo(ctx *InfoContext) {}

// EnterNlstmnt is called when production nlstmnt is entered.
func (s *BasealefListener) EnterNlstmnt(ctx *NlstmntContext) {}

// ExitNlstmnt is called when production nlstmnt is exited.
func (s *BasealefListener) ExitNlstmnt(ctx *NlstmntContext) {}

// EnterZconst is called when production zconst is entered.
func (s *BasealefListener) EnterZconst(ctx *ZconstContext) {}

// ExitZconst is called when production zconst is exited.
func (s *BasealefListener) ExitZconst(ctx *ZconstContext) {}

// EnterZexpr is called when production zexpr is entered.
func (s *BasealefListener) EnterZexpr(ctx *ZexprContext) {}

// ExitZexpr is called when production zexpr is exited.
func (s *BasealefListener) ExitZexpr(ctx *ZexprContext) {}

// EnterExpr_ is called when production expr_ is entered.
func (s *BasealefListener) EnterExpr_(ctx *Expr_Context) {}

// ExitExpr_ is called when production expr_ is exited.
func (s *BasealefListener) ExitExpr_(ctx *Expr_Context) {}

// EnterCastexpr is called when production castexpr is entered.
func (s *BasealefListener) EnterCastexpr(ctx *CastexprContext) {}

// ExitCastexpr is called when production castexpr is exited.
func (s *BasealefListener) ExitCastexpr(ctx *CastexprContext) {}

// EnterTypecast is called when production typecast is entered.
func (s *BasealefListener) EnterTypecast(ctx *TypecastContext) {}

// ExitTypecast is called when production typecast is exited.
func (s *BasealefListener) ExitTypecast(ctx *TypecastContext) {}

// EnterMonexpr is called when production monexpr is entered.
func (s *BasealefListener) EnterMonexpr(ctx *MonexprContext) {}

// ExitMonexpr is called when production monexpr is exited.
func (s *BasealefListener) ExitMonexpr(ctx *MonexprContext) {}

// EnterZtelist is called when production ztelist is entered.
func (s *BasealefListener) EnterZtelist(ctx *ZtelistContext) {}

// ExitZtelist is called when production ztelist is exited.
func (s *BasealefListener) ExitZtelist(ctx *ZtelistContext) {}

// EnterTelist is called when production telist is entered.
func (s *BasealefListener) EnterTelist(ctx *TelistContext) {}

// ExitTelist is called when production telist is exited.
func (s *BasealefListener) ExitTelist(ctx *TelistContext) {}

// EnterTcomp is called when production tcomp is entered.
func (s *BasealefListener) EnterTcomp(ctx *TcompContext) {}

// ExitTcomp is called when production tcomp is exited.
func (s *BasealefListener) ExitTcomp(ctx *TcompContext) {}

// EnterTerm_ is called when production term_ is entered.
func (s *BasealefListener) EnterTerm_(ctx *Term_Context) {}

// ExitTerm_ is called when production term_ is exited.
func (s *BasealefListener) ExitTerm_(ctx *Term_Context) {}

// EnterStag is called when production stag is entered.
func (s *BasealefListener) EnterStag(ctx *StagContext) {}

// ExitStag is called when production stag is exited.
func (s *BasealefListener) ExitStag(ctx *StagContext) {}

// EnterZarlist is called when production zarlist is entered.
func (s *BasealefListener) EnterZarlist(ctx *ZarlistContext) {}

// ExitZarlist is called when production zarlist is exited.
func (s *BasealefListener) ExitZarlist(ctx *ZarlistContext) {}

// EnterElist is called when production elist is entered.
func (s *BasealefListener) EnterElist(ctx *ElistContext) {}

// ExitElist is called when production elist is exited.
func (s *BasealefListener) ExitElist(ctx *ElistContext) {}

// EnterTlist is called when production tlist is entered.
func (s *BasealefListener) EnterTlist(ctx *TlistContext) {}

// ExitTlist is called when production tlist is exited.
func (s *BasealefListener) ExitTlist(ctx *TlistContext) {}

// EnterTname is called when production tname is entered.
func (s *BasealefListener) EnterTname(ctx *TnameContext) {}

// ExitTname is called when production tname is exited.
func (s *BasealefListener) ExitTname(ctx *TnameContext) {}

// EnterVariant is called when production variant is entered.
func (s *BasealefListener) EnterVariant(ctx *VariantContext) {}

// ExitVariant is called when production variant is exited.
func (s *BasealefListener) ExitVariant(ctx *VariantContext) {}

// EnterXtname is called when production xtname is entered.
func (s *BasealefListener) EnterXtname(ctx *XtnameContext) {}

// ExitXtname is called when production xtname is exited.
func (s *BasealefListener) ExitXtname(ctx *XtnameContext) {}

// EnterBufdim is called when production bufdim is entered.
func (s *BasealefListener) EnterBufdim(ctx *BufdimContext) {}

// ExitBufdim is called when production bufdim is exited.
func (s *BasealefListener) ExitBufdim(ctx *BufdimContext) {}

// EnterSclass is called when production sclass is entered.
func (s *BasealefListener) EnterSclass(ctx *SclassContext) {}

// ExitSclass is called when production sclass is exited.
func (s *BasealefListener) ExitSclass(ctx *SclassContext) {}

// EnterTypename is called when production typename is entered.
func (s *BasealefListener) EnterTypename(ctx *TypenameContext) {}

// ExitTypename is called when production typename is exited.
func (s *BasealefListener) ExitTypename(ctx *TypenameContext) {}

// EnterEnum_member is called when production enum_member is entered.
func (s *BasealefListener) EnterEnum_member(ctx *Enum_memberContext) {}

// ExitEnum_member is called when production enum_member is exited.
func (s *BasealefListener) ExitEnum_member(ctx *Enum_memberContext) {}
