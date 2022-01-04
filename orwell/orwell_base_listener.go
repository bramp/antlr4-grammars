// Code generated from orwell.g4 by ANTLR 4.9.3. DO NOT EDIT.

package orwell // orwell
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseorwellListener is a complete listener for a parse tree produced by orwellParser.
type BaseorwellListener struct{}

var _ orwellListener = &BaseorwellListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseorwellListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseorwellListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseorwellListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseorwellListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseorwellListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseorwellListener) ExitProgram(ctx *ProgramContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseorwellListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseorwellListener) ExitDecl(ctx *DeclContext) {}

// EnterSyndecl is called when production syndecl is entered.
func (s *BaseorwellListener) EnterSyndecl(ctx *SyndeclContext) {}

// ExitSyndecl is called when production syndecl is exited.
func (s *BaseorwellListener) ExitSyndecl(ctx *SyndeclContext) {}

// EnterCondecl is called when production condecl is entered.
func (s *BaseorwellListener) EnterCondecl(ctx *CondeclContext) {}

// ExitCondecl is called when production condecl is exited.
func (s *BaseorwellListener) ExitCondecl(ctx *CondeclContext) {}

// EnterTypedecl is called when production typedecl is entered.
func (s *BaseorwellListener) EnterTypedecl(ctx *TypedeclContext) {}

// ExitTypedecl is called when production typedecl is exited.
func (s *BaseorwellListener) ExitTypedecl(ctx *TypedeclContext) {}

// EnterName is called when production name is entered.
func (s *BaseorwellListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseorwellListener) ExitName(ctx *NameContext) {}

// EnterTylhs is called when production tylhs is entered.
func (s *BaseorwellListener) EnterTylhs(ctx *TylhsContext) {}

// ExitTylhs is called when production tylhs is exited.
func (s *BaseorwellListener) ExitTylhs(ctx *TylhsContext) {}

// EnterTylhs1 is called when production tylhs1 is entered.
func (s *BaseorwellListener) EnterTylhs1(ctx *Tylhs1Context) {}

// ExitTylhs1 is called when production tylhs1 is exited.
func (s *BaseorwellListener) ExitTylhs1(ctx *Tylhs1Context) {}

// EnterTylhsprimary is called when production tylhsprimary is entered.
func (s *BaseorwellListener) EnterTylhsprimary(ctx *TylhsprimaryContext) {}

// ExitTylhsprimary is called when production tylhsprimary is exited.
func (s *BaseorwellListener) ExitTylhsprimary(ctx *TylhsprimaryContext) {}

// EnterTylhssection is called when production tylhssection is entered.
func (s *BaseorwellListener) EnterTylhssection(ctx *TylhssectionContext) {}

// ExitTylhssection is called when production tylhssection is exited.
func (s *BaseorwellListener) ExitTylhssection(ctx *TylhssectionContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseorwellListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseorwellListener) ExitType_(ctx *Type_Context) {}

// EnterTyterm1 is called when production tyterm1 is entered.
func (s *BaseorwellListener) EnterTyterm1(ctx *Tyterm1Context) {}

// ExitTyterm1 is called when production tyterm1 is exited.
func (s *BaseorwellListener) ExitTyterm1(ctx *Tyterm1Context) {}

// EnterTyterm2 is called when production tyterm2 is entered.
func (s *BaseorwellListener) EnterTyterm2(ctx *Tyterm2Context) {}

// ExitTyterm2 is called when production tyterm2 is exited.
func (s *BaseorwellListener) ExitTyterm2(ctx *Tyterm2Context) {}

// EnterTyprimaryname is called when production typrimaryname is entered.
func (s *BaseorwellListener) EnterTyprimaryname(ctx *TyprimarynameContext) {}

// ExitTyprimaryname is called when production typrimaryname is exited.
func (s *BaseorwellListener) ExitTyprimaryname(ctx *TyprimarynameContext) {}

// EnterTyprimary is called when production typrimary is entered.
func (s *BaseorwellListener) EnterTyprimary(ctx *TyprimaryContext) {}

// ExitTyprimary is called when production typrimary is exited.
func (s *BaseorwellListener) ExitTyprimary(ctx *TyprimaryContext) {}

// EnterTysection is called when production tysection is entered.
func (s *BaseorwellListener) EnterTysection(ctx *TysectionContext) {}

// ExitTysection is called when production tysection is exited.
func (s *BaseorwellListener) ExitTysection(ctx *TysectionContext) {}

// EnterTylist is called when production tylist is entered.
func (s *BaseorwellListener) EnterTylist(ctx *TylistContext) {}

// ExitTylist is called when production tylist is exited.
func (s *BaseorwellListener) ExitTylist(ctx *TylistContext) {}

// EnterTytuple is called when production tytuple is entered.
func (s *BaseorwellListener) EnterTytuple(ctx *TytupleContext) {}

// ExitTytuple is called when production tytuple is exited.
func (s *BaseorwellListener) ExitTytuple(ctx *TytupleContext) {}

// EnterConstruct is called when production construct is entered.
func (s *BaseorwellListener) EnterConstruct(ctx *ConstructContext) {}

// ExitConstruct is called when production construct is exited.
func (s *BaseorwellListener) ExitConstruct(ctx *ConstructContext) {}

// EnterOpdecl is called when production opdecl is entered.
func (s *BaseorwellListener) EnterOpdecl(ctx *OpdeclContext) {}

// ExitOpdecl is called when production opdecl is exited.
func (s *BaseorwellListener) ExitOpdecl(ctx *OpdeclContext) {}

// EnterOpkind is called when production opkind is entered.
func (s *BaseorwellListener) EnterOpkind(ctx *OpkindContext) {}

// ExitOpkind is called when production opkind is exited.
func (s *BaseorwellListener) ExitOpkind(ctx *OpkindContext) {}

// EnterAssoc is called when production assoc is entered.
func (s *BaseorwellListener) EnterAssoc(ctx *AssocContext) {}

// ExitAssoc is called when production assoc is exited.
func (s *BaseorwellListener) ExitAssoc(ctx *AssocContext) {}

// EnterDef_ is called when production def_ is entered.
func (s *BaseorwellListener) EnterDef_(ctx *Def_Context) {}

// ExitDef_ is called when production def_ is exited.
func (s *BaseorwellListener) ExitDef_(ctx *Def_Context) {}

// EnterRhs is called when production rhs is entered.
func (s *BaseorwellListener) EnterRhs(ctx *RhsContext) {}

// ExitRhs is called when production rhs is exited.
func (s *BaseorwellListener) ExitRhs(ctx *RhsContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseorwellListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseorwellListener) ExitConditional(ctx *ConditionalContext) {}

// EnterIfpart is called when production ifpart is entered.
func (s *BaseorwellListener) EnterIfpart(ctx *IfpartContext) {}

// ExitIfpart is called when production ifpart is exited.
func (s *BaseorwellListener) ExitIfpart(ctx *IfpartContext) {}

// EnterOtherpart is called when production otherpart is entered.
func (s *BaseorwellListener) EnterOtherpart(ctx *OtherpartContext) {}

// ExitOtherpart is called when production otherpart is exited.
func (s *BaseorwellListener) ExitOtherpart(ctx *OtherpartContext) {}

// EnterWherepart is called when production wherepart is entered.
func (s *BaseorwellListener) EnterWherepart(ctx *WherepartContext) {}

// ExitWherepart is called when production wherepart is exited.
func (s *BaseorwellListener) ExitWherepart(ctx *WherepartContext) {}

// EnterPat is called when production pat is entered.
func (s *BaseorwellListener) EnterPat(ctx *PatContext) {}

// ExitPat is called when production pat is exited.
func (s *BaseorwellListener) ExitPat(ctx *PatContext) {}

// EnterPat1 is called when production pat1 is entered.
func (s *BaseorwellListener) EnterPat1(ctx *Pat1Context) {}

// ExitPat1 is called when production pat1 is exited.
func (s *BaseorwellListener) ExitPat1(ctx *Pat1Context) {}

// EnterPat2 is called when production pat2 is entered.
func (s *BaseorwellListener) EnterPat2(ctx *Pat2Context) {}

// ExitPat2 is called when production pat2 is exited.
func (s *BaseorwellListener) ExitPat2(ctx *Pat2Context) {}

// EnterPatprimaryname is called when production patprimaryname is entered.
func (s *BaseorwellListener) EnterPatprimaryname(ctx *PatprimarynameContext) {}

// ExitPatprimaryname is called when production patprimaryname is exited.
func (s *BaseorwellListener) ExitPatprimaryname(ctx *PatprimarynameContext) {}

// EnterPatprimary is called when production patprimary is entered.
func (s *BaseorwellListener) EnterPatprimary(ctx *PatprimaryContext) {}

// ExitPatprimary is called when production patprimary is exited.
func (s *BaseorwellListener) ExitPatprimary(ctx *PatprimaryContext) {}

// EnterPatsection is called when production patsection is entered.
func (s *BaseorwellListener) EnterPatsection(ctx *PatsectionContext) {}

// ExitPatsection is called when production patsection is exited.
func (s *BaseorwellListener) ExitPatsection(ctx *PatsectionContext) {}

// EnterPattuple is called when production pattuple is entered.
func (s *BaseorwellListener) EnterPattuple(ctx *PattupleContext) {}

// ExitPattuple is called when production pattuple is exited.
func (s *BaseorwellListener) ExitPattuple(ctx *PattupleContext) {}

// EnterPatlist is called when production patlist is entered.
func (s *BaseorwellListener) EnterPatlist(ctx *PatlistContext) {}

// ExitPatlist is called when production patlist is exited.
func (s *BaseorwellListener) ExitPatlist(ctx *PatlistContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseorwellListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseorwellListener) ExitTerm(ctx *TermContext) {}

// EnterTerm1 is called when production term1 is entered.
func (s *BaseorwellListener) EnterTerm1(ctx *Term1Context) {}

// ExitTerm1 is called when production term1 is exited.
func (s *BaseorwellListener) ExitTerm1(ctx *Term1Context) {}

// EnterTerm2 is called when production term2 is entered.
func (s *BaseorwellListener) EnterTerm2(ctx *Term2Context) {}

// ExitTerm2 is called when production term2 is exited.
func (s *BaseorwellListener) ExitTerm2(ctx *Term2Context) {}

// EnterPrimaryname is called when production primaryname is entered.
func (s *BaseorwellListener) EnterPrimaryname(ctx *PrimarynameContext) {}

// ExitPrimaryname is called when production primaryname is exited.
func (s *BaseorwellListener) ExitPrimaryname(ctx *PrimarynameContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseorwellListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseorwellListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterSection is called when production section is entered.
func (s *BaseorwellListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseorwellListener) ExitSection(ctx *SectionContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseorwellListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseorwellListener) ExitList_(ctx *List_Context) {}

// EnterTuple_ is called when production tuple_ is entered.
func (s *BaseorwellListener) EnterTuple_(ctx *Tuple_Context) {}

// ExitTuple_ is called when production tuple_ is exited.
func (s *BaseorwellListener) ExitTuple_(ctx *Tuple_Context) {}

// EnterListform is called when production listform is entered.
func (s *BaseorwellListener) EnterListform(ctx *ListformContext) {}

// ExitListform is called when production listform is exited.
func (s *BaseorwellListener) ExitListform(ctx *ListformContext) {}

// EnterUpto is called when production upto is entered.
func (s *BaseorwellListener) EnterUpto(ctx *UptoContext) {}

// ExitUpto is called when production upto is exited.
func (s *BaseorwellListener) ExitUpto(ctx *UptoContext) {}

// EnterComp is called when production comp is entered.
func (s *BaseorwellListener) EnterComp(ctx *CompContext) {}

// ExitComp is called when production comp is exited.
func (s *BaseorwellListener) ExitComp(ctx *CompContext) {}

// EnterQualifier is called when production qualifier is entered.
func (s *BaseorwellListener) EnterQualifier(ctx *QualifierContext) {}

// ExitQualifier is called when production qualifier is exited.
func (s *BaseorwellListener) ExitQualifier(ctx *QualifierContext) {}

// EnterFliteral is called when production fliteral is entered.
func (s *BaseorwellListener) EnterFliteral(ctx *FliteralContext) {}

// ExitFliteral is called when production fliteral is exited.
func (s *BaseorwellListener) ExitFliteral(ctx *FliteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseorwellListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseorwellListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInfix is called when production infix is entered.
func (s *BaseorwellListener) EnterInfix(ctx *InfixContext) {}

// ExitInfix is called when production infix is exited.
func (s *BaseorwellListener) ExitInfix(ctx *InfixContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BaseorwellListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BaseorwellListener) ExitPrefix(ctx *PrefixContext) {}

// EnterTyname is called when production tyname is entered.
func (s *BaseorwellListener) EnterTyname(ctx *TynameContext) {}

// ExitTyname is called when production tyname is exited.
func (s *BaseorwellListener) ExitTyname(ctx *TynameContext) {}

// EnterTyvar is called when production tyvar is entered.
func (s *BaseorwellListener) EnterTyvar(ctx *TyvarContext) {}

// ExitTyvar is called when production tyvar is exited.
func (s *BaseorwellListener) ExitTyvar(ctx *TyvarContext) {}

// EnterCon is called when production con is entered.
func (s *BaseorwellListener) EnterCon(ctx *ConContext) {}

// ExitCon is called when production con is exited.
func (s *BaseorwellListener) ExitCon(ctx *ConContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseorwellListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseorwellListener) ExitVar_(ctx *Var_Context) {}
