// Code generated from orwell.g4 by ANTLR 4.9.3. DO NOT EDIT.

package orwell // orwell
import "github.com/antlr/antlr4/runtime/Go/antlr"

// orwellListener is a complete listener for a parse tree produced by orwellParser.
type orwellListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterSyndecl is called when entering the syndecl production.
	EnterSyndecl(c *SyndeclContext)

	// EnterCondecl is called when entering the condecl production.
	EnterCondecl(c *CondeclContext)

	// EnterTypedecl is called when entering the typedecl production.
	EnterTypedecl(c *TypedeclContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTylhs is called when entering the tylhs production.
	EnterTylhs(c *TylhsContext)

	// EnterTylhs1 is called when entering the tylhs1 production.
	EnterTylhs1(c *Tylhs1Context)

	// EnterTylhsprimary is called when entering the tylhsprimary production.
	EnterTylhsprimary(c *TylhsprimaryContext)

	// EnterTylhssection is called when entering the tylhssection production.
	EnterTylhssection(c *TylhssectionContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTyterm1 is called when entering the tyterm1 production.
	EnterTyterm1(c *Tyterm1Context)

	// EnterTyterm2 is called when entering the tyterm2 production.
	EnterTyterm2(c *Tyterm2Context)

	// EnterTyprimaryname is called when entering the typrimaryname production.
	EnterTyprimaryname(c *TyprimarynameContext)

	// EnterTyprimary is called when entering the typrimary production.
	EnterTyprimary(c *TyprimaryContext)

	// EnterTysection is called when entering the tysection production.
	EnterTysection(c *TysectionContext)

	// EnterTylist is called when entering the tylist production.
	EnterTylist(c *TylistContext)

	// EnterTytuple is called when entering the tytuple production.
	EnterTytuple(c *TytupleContext)

	// EnterConstruct is called when entering the construct production.
	EnterConstruct(c *ConstructContext)

	// EnterOpdecl is called when entering the opdecl production.
	EnterOpdecl(c *OpdeclContext)

	// EnterOpkind is called when entering the opkind production.
	EnterOpkind(c *OpkindContext)

	// EnterAssoc is called when entering the assoc production.
	EnterAssoc(c *AssocContext)

	// EnterDef_ is called when entering the def_ production.
	EnterDef_(c *Def_Context)

	// EnterRhs is called when entering the rhs production.
	EnterRhs(c *RhsContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterIfpart is called when entering the ifpart production.
	EnterIfpart(c *IfpartContext)

	// EnterOtherpart is called when entering the otherpart production.
	EnterOtherpart(c *OtherpartContext)

	// EnterWherepart is called when entering the wherepart production.
	EnterWherepart(c *WherepartContext)

	// EnterPat is called when entering the pat production.
	EnterPat(c *PatContext)

	// EnterPat1 is called when entering the pat1 production.
	EnterPat1(c *Pat1Context)

	// EnterPat2 is called when entering the pat2 production.
	EnterPat2(c *Pat2Context)

	// EnterPatprimaryname is called when entering the patprimaryname production.
	EnterPatprimaryname(c *PatprimarynameContext)

	// EnterPatprimary is called when entering the patprimary production.
	EnterPatprimary(c *PatprimaryContext)

	// EnterPatsection is called when entering the patsection production.
	EnterPatsection(c *PatsectionContext)

	// EnterPattuple is called when entering the pattuple production.
	EnterPattuple(c *PattupleContext)

	// EnterPatlist is called when entering the patlist production.
	EnterPatlist(c *PatlistContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTerm1 is called when entering the term1 production.
	EnterTerm1(c *Term1Context)

	// EnterTerm2 is called when entering the term2 production.
	EnterTerm2(c *Term2Context)

	// EnterPrimaryname is called when entering the primaryname production.
	EnterPrimaryname(c *PrimarynameContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterTuple_ is called when entering the tuple_ production.
	EnterTuple_(c *Tuple_Context)

	// EnterListform is called when entering the listform production.
	EnterListform(c *ListformContext)

	// EnterUpto is called when entering the upto production.
	EnterUpto(c *UptoContext)

	// EnterComp is called when entering the comp production.
	EnterComp(c *CompContext)

	// EnterQualifier is called when entering the qualifier production.
	EnterQualifier(c *QualifierContext)

	// EnterFliteral is called when entering the fliteral production.
	EnterFliteral(c *FliteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInfix is called when entering the infix production.
	EnterInfix(c *InfixContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterTyname is called when entering the tyname production.
	EnterTyname(c *TynameContext)

	// EnterTyvar is called when entering the tyvar production.
	EnterTyvar(c *TyvarContext)

	// EnterCon is called when entering the con production.
	EnterCon(c *ConContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitSyndecl is called when exiting the syndecl production.
	ExitSyndecl(c *SyndeclContext)

	// ExitCondecl is called when exiting the condecl production.
	ExitCondecl(c *CondeclContext)

	// ExitTypedecl is called when exiting the typedecl production.
	ExitTypedecl(c *TypedeclContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTylhs is called when exiting the tylhs production.
	ExitTylhs(c *TylhsContext)

	// ExitTylhs1 is called when exiting the tylhs1 production.
	ExitTylhs1(c *Tylhs1Context)

	// ExitTylhsprimary is called when exiting the tylhsprimary production.
	ExitTylhsprimary(c *TylhsprimaryContext)

	// ExitTylhssection is called when exiting the tylhssection production.
	ExitTylhssection(c *TylhssectionContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTyterm1 is called when exiting the tyterm1 production.
	ExitTyterm1(c *Tyterm1Context)

	// ExitTyterm2 is called when exiting the tyterm2 production.
	ExitTyterm2(c *Tyterm2Context)

	// ExitTyprimaryname is called when exiting the typrimaryname production.
	ExitTyprimaryname(c *TyprimarynameContext)

	// ExitTyprimary is called when exiting the typrimary production.
	ExitTyprimary(c *TyprimaryContext)

	// ExitTysection is called when exiting the tysection production.
	ExitTysection(c *TysectionContext)

	// ExitTylist is called when exiting the tylist production.
	ExitTylist(c *TylistContext)

	// ExitTytuple is called when exiting the tytuple production.
	ExitTytuple(c *TytupleContext)

	// ExitConstruct is called when exiting the construct production.
	ExitConstruct(c *ConstructContext)

	// ExitOpdecl is called when exiting the opdecl production.
	ExitOpdecl(c *OpdeclContext)

	// ExitOpkind is called when exiting the opkind production.
	ExitOpkind(c *OpkindContext)

	// ExitAssoc is called when exiting the assoc production.
	ExitAssoc(c *AssocContext)

	// ExitDef_ is called when exiting the def_ production.
	ExitDef_(c *Def_Context)

	// ExitRhs is called when exiting the rhs production.
	ExitRhs(c *RhsContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitIfpart is called when exiting the ifpart production.
	ExitIfpart(c *IfpartContext)

	// ExitOtherpart is called when exiting the otherpart production.
	ExitOtherpart(c *OtherpartContext)

	// ExitWherepart is called when exiting the wherepart production.
	ExitWherepart(c *WherepartContext)

	// ExitPat is called when exiting the pat production.
	ExitPat(c *PatContext)

	// ExitPat1 is called when exiting the pat1 production.
	ExitPat1(c *Pat1Context)

	// ExitPat2 is called when exiting the pat2 production.
	ExitPat2(c *Pat2Context)

	// ExitPatprimaryname is called when exiting the patprimaryname production.
	ExitPatprimaryname(c *PatprimarynameContext)

	// ExitPatprimary is called when exiting the patprimary production.
	ExitPatprimary(c *PatprimaryContext)

	// ExitPatsection is called when exiting the patsection production.
	ExitPatsection(c *PatsectionContext)

	// ExitPattuple is called when exiting the pattuple production.
	ExitPattuple(c *PattupleContext)

	// ExitPatlist is called when exiting the patlist production.
	ExitPatlist(c *PatlistContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTerm1 is called when exiting the term1 production.
	ExitTerm1(c *Term1Context)

	// ExitTerm2 is called when exiting the term2 production.
	ExitTerm2(c *Term2Context)

	// ExitPrimaryname is called when exiting the primaryname production.
	ExitPrimaryname(c *PrimarynameContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitTuple_ is called when exiting the tuple_ production.
	ExitTuple_(c *Tuple_Context)

	// ExitListform is called when exiting the listform production.
	ExitListform(c *ListformContext)

	// ExitUpto is called when exiting the upto production.
	ExitUpto(c *UptoContext)

	// ExitComp is called when exiting the comp production.
	ExitComp(c *CompContext)

	// ExitQualifier is called when exiting the qualifier production.
	ExitQualifier(c *QualifierContext)

	// ExitFliteral is called when exiting the fliteral production.
	ExitFliteral(c *FliteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInfix is called when exiting the infix production.
	ExitInfix(c *InfixContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitTyname is called when exiting the tyname production.
	ExitTyname(c *TynameContext)

	// ExitTyvar is called when exiting the tyvar production.
	ExitTyvar(c *TyvarContext)

	// ExitCon is called when exiting the con production.
	ExitCon(c *ConContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)
}
