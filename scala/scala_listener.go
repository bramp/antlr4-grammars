// Code generated from Scala.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scala // Scala
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScalaListener is a complete listener for a parse tree produced by ScalaParser.
type ScalaListener interface {
	antlr.ParseTreeListener

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterQualId is called when entering the qualId production.
	EnterQualId(c *QualIdContext)

	// EnterIds is called when entering the ids production.
	EnterIds(c *IdsContext)

	// EnterStableId is called when entering the stableId production.
	EnterStableId(c *StableIdContext)

	// EnterClassQualifier is called when entering the classQualifier production.
	EnterClassQualifier(c *ClassQualifierContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFunctionArgTypes is called when entering the functionArgTypes production.
	EnterFunctionArgTypes(c *FunctionArgTypesContext)

	// EnterExistentialClause is called when entering the existentialClause production.
	EnterExistentialClause(c *ExistentialClauseContext)

	// EnterExistentialDcl is called when entering the existentialDcl production.
	EnterExistentialDcl(c *ExistentialDclContext)

	// EnterInfixType is called when entering the infixType production.
	EnterInfixType(c *InfixTypeContext)

	// EnterCompoundType is called when entering the compoundType production.
	EnterCompoundType(c *CompoundTypeContext)

	// EnterAnnotType is called when entering the annotType production.
	EnterAnnotType(c *AnnotTypeContext)

	// EnterSimpleType is called when entering the simpleType production.
	EnterSimpleType(c *SimpleTypeContext)

	// EnterTypeArgs is called when entering the typeArgs production.
	EnterTypeArgs(c *TypeArgsContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterRefinement is called when entering the refinement production.
	EnterRefinement(c *RefinementContext)

	// EnterRefineStat is called when entering the refineStat production.
	EnterRefineStat(c *RefineStatContext)

	// EnterTypePat is called when entering the typePat production.
	EnterTypePat(c *TypePatContext)

	// EnterAscription is called when entering the ascription production.
	EnterAscription(c *AscriptionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr1 is called when entering the expr1 production.
	EnterExpr1(c *Expr1Context)

	// EnterPrefixDef is called when entering the prefixDef production.
	EnterPrefixDef(c *PrefixDefContext)

	// EnterPostfixExpr is called when entering the postfixExpr production.
	EnterPostfixExpr(c *PostfixExprContext)

	// EnterInfixExpr is called when entering the infixExpr production.
	EnterInfixExpr(c *InfixExprContext)

	// EnterPrefixExpr is called when entering the prefixExpr production.
	EnterPrefixExpr(c *PrefixExprContext)

	// EnterSimpleExpr is called when entering the simpleExpr production.
	EnterSimpleExpr(c *SimpleExprContext)

	// EnterSimpleExpr1 is called when entering the simpleExpr1 production.
	EnterSimpleExpr1(c *SimpleExpr1Context)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// EnterArgumentExprs is called when entering the argumentExprs production.
	EnterArgumentExprs(c *ArgumentExprsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterBlockExpr is called when entering the blockExpr production.
	EnterBlockExpr(c *BlockExprContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStat is called when entering the blockStat production.
	EnterBlockStat(c *BlockStatContext)

	// EnterResultExpr is called when entering the resultExpr production.
	EnterResultExpr(c *ResultExprContext)

	// EnterEnumerators is called when entering the enumerators production.
	EnterEnumerators(c *EnumeratorsContext)

	// EnterGenerator is called when entering the generator production.
	EnterGenerator(c *GeneratorContext)

	// EnterCaseClauses is called when entering the caseClauses production.
	EnterCaseClauses(c *CaseClausesContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterGuard_ is called when entering the guard_ production.
	EnterGuard_(c *Guard_Context)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterPattern1 is called when entering the pattern1 production.
	EnterPattern1(c *Pattern1Context)

	// EnterPattern2 is called when entering the pattern2 production.
	EnterPattern2(c *Pattern2Context)

	// EnterPattern3 is called when entering the pattern3 production.
	EnterPattern3(c *Pattern3Context)

	// EnterSimplePattern is called when entering the simplePattern production.
	EnterSimplePattern(c *SimplePatternContext)

	// EnterPatterns is called when entering the patterns production.
	EnterPatterns(c *PatternsContext)

	// EnterTypeParamClause is called when entering the typeParamClause production.
	EnterTypeParamClause(c *TypeParamClauseContext)

	// EnterFunTypeParamClause is called when entering the funTypeParamClause production.
	EnterFunTypeParamClause(c *FunTypeParamClauseContext)

	// EnterVariantTypeParam is called when entering the variantTypeParam production.
	EnterVariantTypeParam(c *VariantTypeParamContext)

	// EnterTypeParam is called when entering the typeParam production.
	EnterTypeParam(c *TypeParamContext)

	// EnterParamClauses is called when entering the paramClauses production.
	EnterParamClauses(c *ParamClausesContext)

	// EnterParamClause is called when entering the paramClause production.
	EnterParamClause(c *ParamClauseContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParamType is called when entering the paramType production.
	EnterParamType(c *ParamTypeContext)

	// EnterClassParamClauses is called when entering the classParamClauses production.
	EnterClassParamClauses(c *ClassParamClausesContext)

	// EnterClassParamClause is called when entering the classParamClause production.
	EnterClassParamClause(c *ClassParamClauseContext)

	// EnterClassParams is called when entering the classParams production.
	EnterClassParams(c *ClassParamsContext)

	// EnterClassParam is called when entering the classParam production.
	EnterClassParam(c *ClassParamContext)

	// EnterBindings is called when entering the bindings production.
	EnterBindings(c *BindingsContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterLocalModifier is called when entering the localModifier production.
	EnterLocalModifier(c *LocalModifierContext)

	// EnterAccessModifier is called when entering the accessModifier production.
	EnterAccessModifier(c *AccessModifierContext)

	// EnterAccessQualifier is called when entering the accessQualifier production.
	EnterAccessQualifier(c *AccessQualifierContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterConstrAnnotation is called when entering the constrAnnotation production.
	EnterConstrAnnotation(c *ConstrAnnotationContext)

	// EnterTemplateBody is called when entering the templateBody production.
	EnterTemplateBody(c *TemplateBodyContext)

	// EnterTemplateStat is called when entering the templateStat production.
	EnterTemplateStat(c *TemplateStatContext)

	// EnterSelfType is called when entering the selfType production.
	EnterSelfType(c *SelfTypeContext)

	// EnterImport_ is called when entering the import_ production.
	EnterImport_(c *Import_Context)

	// EnterImportExpr is called when entering the importExpr production.
	EnterImportExpr(c *ImportExprContext)

	// EnterImportSelectors is called when entering the importSelectors production.
	EnterImportSelectors(c *ImportSelectorsContext)

	// EnterImportSelector is called when entering the importSelector production.
	EnterImportSelector(c *ImportSelectorContext)

	// EnterDcl is called when entering the dcl production.
	EnterDcl(c *DclContext)

	// EnterValDcl is called when entering the valDcl production.
	EnterValDcl(c *ValDclContext)

	// EnterVarDcl is called when entering the varDcl production.
	EnterVarDcl(c *VarDclContext)

	// EnterFunDcl is called when entering the funDcl production.
	EnterFunDcl(c *FunDclContext)

	// EnterFunSig is called when entering the funSig production.
	EnterFunSig(c *FunSigContext)

	// EnterTypeDcl is called when entering the typeDcl production.
	EnterTypeDcl(c *TypeDclContext)

	// EnterPatVarDef is called when entering the patVarDef production.
	EnterPatVarDef(c *PatVarDefContext)

	// EnterDef_ is called when entering the def_ production.
	EnterDef_(c *Def_Context)

	// EnterPatDef is called when entering the patDef production.
	EnterPatDef(c *PatDefContext)

	// EnterVarDef is called when entering the varDef production.
	EnterVarDef(c *VarDefContext)

	// EnterFunDef is called when entering the funDef production.
	EnterFunDef(c *FunDefContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterTmplDef is called when entering the tmplDef production.
	EnterTmplDef(c *TmplDefContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterTraitDef is called when entering the traitDef production.
	EnterTraitDef(c *TraitDefContext)

	// EnterObjectDef is called when entering the objectDef production.
	EnterObjectDef(c *ObjectDefContext)

	// EnterClassTemplateOpt is called when entering the classTemplateOpt production.
	EnterClassTemplateOpt(c *ClassTemplateOptContext)

	// EnterTraitTemplateOpt is called when entering the traitTemplateOpt production.
	EnterTraitTemplateOpt(c *TraitTemplateOptContext)

	// EnterClassTemplate is called when entering the classTemplate production.
	EnterClassTemplate(c *ClassTemplateContext)

	// EnterTraitTemplate is called when entering the traitTemplate production.
	EnterTraitTemplate(c *TraitTemplateContext)

	// EnterClassParents is called when entering the classParents production.
	EnterClassParents(c *ClassParentsContext)

	// EnterTraitParents is called when entering the traitParents production.
	EnterTraitParents(c *TraitParentsContext)

	// EnterConstr is called when entering the constr production.
	EnterConstr(c *ConstrContext)

	// EnterEarlyDefs is called when entering the earlyDefs production.
	EnterEarlyDefs(c *EarlyDefsContext)

	// EnterEarlyDef is called when entering the earlyDef production.
	EnterEarlyDef(c *EarlyDefContext)

	// EnterConstrExpr is called when entering the constrExpr production.
	EnterConstrExpr(c *ConstrExprContext)

	// EnterConstrBlock is called when entering the constrBlock production.
	EnterConstrBlock(c *ConstrBlockContext)

	// EnterSelfInvocation is called when entering the selfInvocation production.
	EnterSelfInvocation(c *SelfInvocationContext)

	// EnterTopStatSeq is called when entering the topStatSeq production.
	EnterTopStatSeq(c *TopStatSeqContext)

	// EnterTopStat is called when entering the topStat production.
	EnterTopStat(c *TopStatContext)

	// EnterPackaging is called when entering the packaging production.
	EnterPackaging(c *PackagingContext)

	// EnterPackageObject is called when entering the packageObject production.
	EnterPackageObject(c *PackageObjectContext)

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitQualId is called when exiting the qualId production.
	ExitQualId(c *QualIdContext)

	// ExitIds is called when exiting the ids production.
	ExitIds(c *IdsContext)

	// ExitStableId is called when exiting the stableId production.
	ExitStableId(c *StableIdContext)

	// ExitClassQualifier is called when exiting the classQualifier production.
	ExitClassQualifier(c *ClassQualifierContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFunctionArgTypes is called when exiting the functionArgTypes production.
	ExitFunctionArgTypes(c *FunctionArgTypesContext)

	// ExitExistentialClause is called when exiting the existentialClause production.
	ExitExistentialClause(c *ExistentialClauseContext)

	// ExitExistentialDcl is called when exiting the existentialDcl production.
	ExitExistentialDcl(c *ExistentialDclContext)

	// ExitInfixType is called when exiting the infixType production.
	ExitInfixType(c *InfixTypeContext)

	// ExitCompoundType is called when exiting the compoundType production.
	ExitCompoundType(c *CompoundTypeContext)

	// ExitAnnotType is called when exiting the annotType production.
	ExitAnnotType(c *AnnotTypeContext)

	// ExitSimpleType is called when exiting the simpleType production.
	ExitSimpleType(c *SimpleTypeContext)

	// ExitTypeArgs is called when exiting the typeArgs production.
	ExitTypeArgs(c *TypeArgsContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitRefinement is called when exiting the refinement production.
	ExitRefinement(c *RefinementContext)

	// ExitRefineStat is called when exiting the refineStat production.
	ExitRefineStat(c *RefineStatContext)

	// ExitTypePat is called when exiting the typePat production.
	ExitTypePat(c *TypePatContext)

	// ExitAscription is called when exiting the ascription production.
	ExitAscription(c *AscriptionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr1 is called when exiting the expr1 production.
	ExitExpr1(c *Expr1Context)

	// ExitPrefixDef is called when exiting the prefixDef production.
	ExitPrefixDef(c *PrefixDefContext)

	// ExitPostfixExpr is called when exiting the postfixExpr production.
	ExitPostfixExpr(c *PostfixExprContext)

	// ExitInfixExpr is called when exiting the infixExpr production.
	ExitInfixExpr(c *InfixExprContext)

	// ExitPrefixExpr is called when exiting the prefixExpr production.
	ExitPrefixExpr(c *PrefixExprContext)

	// ExitSimpleExpr is called when exiting the simpleExpr production.
	ExitSimpleExpr(c *SimpleExprContext)

	// ExitSimpleExpr1 is called when exiting the simpleExpr1 production.
	ExitSimpleExpr1(c *SimpleExpr1Context)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)

	// ExitArgumentExprs is called when exiting the argumentExprs production.
	ExitArgumentExprs(c *ArgumentExprsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitBlockExpr is called when exiting the blockExpr production.
	ExitBlockExpr(c *BlockExprContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStat is called when exiting the blockStat production.
	ExitBlockStat(c *BlockStatContext)

	// ExitResultExpr is called when exiting the resultExpr production.
	ExitResultExpr(c *ResultExprContext)

	// ExitEnumerators is called when exiting the enumerators production.
	ExitEnumerators(c *EnumeratorsContext)

	// ExitGenerator is called when exiting the generator production.
	ExitGenerator(c *GeneratorContext)

	// ExitCaseClauses is called when exiting the caseClauses production.
	ExitCaseClauses(c *CaseClausesContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitGuard_ is called when exiting the guard_ production.
	ExitGuard_(c *Guard_Context)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitPattern1 is called when exiting the pattern1 production.
	ExitPattern1(c *Pattern1Context)

	// ExitPattern2 is called when exiting the pattern2 production.
	ExitPattern2(c *Pattern2Context)

	// ExitPattern3 is called when exiting the pattern3 production.
	ExitPattern3(c *Pattern3Context)

	// ExitSimplePattern is called when exiting the simplePattern production.
	ExitSimplePattern(c *SimplePatternContext)

	// ExitPatterns is called when exiting the patterns production.
	ExitPatterns(c *PatternsContext)

	// ExitTypeParamClause is called when exiting the typeParamClause production.
	ExitTypeParamClause(c *TypeParamClauseContext)

	// ExitFunTypeParamClause is called when exiting the funTypeParamClause production.
	ExitFunTypeParamClause(c *FunTypeParamClauseContext)

	// ExitVariantTypeParam is called when exiting the variantTypeParam production.
	ExitVariantTypeParam(c *VariantTypeParamContext)

	// ExitTypeParam is called when exiting the typeParam production.
	ExitTypeParam(c *TypeParamContext)

	// ExitParamClauses is called when exiting the paramClauses production.
	ExitParamClauses(c *ParamClausesContext)

	// ExitParamClause is called when exiting the paramClause production.
	ExitParamClause(c *ParamClauseContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParamType is called when exiting the paramType production.
	ExitParamType(c *ParamTypeContext)

	// ExitClassParamClauses is called when exiting the classParamClauses production.
	ExitClassParamClauses(c *ClassParamClausesContext)

	// ExitClassParamClause is called when exiting the classParamClause production.
	ExitClassParamClause(c *ClassParamClauseContext)

	// ExitClassParams is called when exiting the classParams production.
	ExitClassParams(c *ClassParamsContext)

	// ExitClassParam is called when exiting the classParam production.
	ExitClassParam(c *ClassParamContext)

	// ExitBindings is called when exiting the bindings production.
	ExitBindings(c *BindingsContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitLocalModifier is called when exiting the localModifier production.
	ExitLocalModifier(c *LocalModifierContext)

	// ExitAccessModifier is called when exiting the accessModifier production.
	ExitAccessModifier(c *AccessModifierContext)

	// ExitAccessQualifier is called when exiting the accessQualifier production.
	ExitAccessQualifier(c *AccessQualifierContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitConstrAnnotation is called when exiting the constrAnnotation production.
	ExitConstrAnnotation(c *ConstrAnnotationContext)

	// ExitTemplateBody is called when exiting the templateBody production.
	ExitTemplateBody(c *TemplateBodyContext)

	// ExitTemplateStat is called when exiting the templateStat production.
	ExitTemplateStat(c *TemplateStatContext)

	// ExitSelfType is called when exiting the selfType production.
	ExitSelfType(c *SelfTypeContext)

	// ExitImport_ is called when exiting the import_ production.
	ExitImport_(c *Import_Context)

	// ExitImportExpr is called when exiting the importExpr production.
	ExitImportExpr(c *ImportExprContext)

	// ExitImportSelectors is called when exiting the importSelectors production.
	ExitImportSelectors(c *ImportSelectorsContext)

	// ExitImportSelector is called when exiting the importSelector production.
	ExitImportSelector(c *ImportSelectorContext)

	// ExitDcl is called when exiting the dcl production.
	ExitDcl(c *DclContext)

	// ExitValDcl is called when exiting the valDcl production.
	ExitValDcl(c *ValDclContext)

	// ExitVarDcl is called when exiting the varDcl production.
	ExitVarDcl(c *VarDclContext)

	// ExitFunDcl is called when exiting the funDcl production.
	ExitFunDcl(c *FunDclContext)

	// ExitFunSig is called when exiting the funSig production.
	ExitFunSig(c *FunSigContext)

	// ExitTypeDcl is called when exiting the typeDcl production.
	ExitTypeDcl(c *TypeDclContext)

	// ExitPatVarDef is called when exiting the patVarDef production.
	ExitPatVarDef(c *PatVarDefContext)

	// ExitDef_ is called when exiting the def_ production.
	ExitDef_(c *Def_Context)

	// ExitPatDef is called when exiting the patDef production.
	ExitPatDef(c *PatDefContext)

	// ExitVarDef is called when exiting the varDef production.
	ExitVarDef(c *VarDefContext)

	// ExitFunDef is called when exiting the funDef production.
	ExitFunDef(c *FunDefContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitTmplDef is called when exiting the tmplDef production.
	ExitTmplDef(c *TmplDefContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitTraitDef is called when exiting the traitDef production.
	ExitTraitDef(c *TraitDefContext)

	// ExitObjectDef is called when exiting the objectDef production.
	ExitObjectDef(c *ObjectDefContext)

	// ExitClassTemplateOpt is called when exiting the classTemplateOpt production.
	ExitClassTemplateOpt(c *ClassTemplateOptContext)

	// ExitTraitTemplateOpt is called when exiting the traitTemplateOpt production.
	ExitTraitTemplateOpt(c *TraitTemplateOptContext)

	// ExitClassTemplate is called when exiting the classTemplate production.
	ExitClassTemplate(c *ClassTemplateContext)

	// ExitTraitTemplate is called when exiting the traitTemplate production.
	ExitTraitTemplate(c *TraitTemplateContext)

	// ExitClassParents is called when exiting the classParents production.
	ExitClassParents(c *ClassParentsContext)

	// ExitTraitParents is called when exiting the traitParents production.
	ExitTraitParents(c *TraitParentsContext)

	// ExitConstr is called when exiting the constr production.
	ExitConstr(c *ConstrContext)

	// ExitEarlyDefs is called when exiting the earlyDefs production.
	ExitEarlyDefs(c *EarlyDefsContext)

	// ExitEarlyDef is called when exiting the earlyDef production.
	ExitEarlyDef(c *EarlyDefContext)

	// ExitConstrExpr is called when exiting the constrExpr production.
	ExitConstrExpr(c *ConstrExprContext)

	// ExitConstrBlock is called when exiting the constrBlock production.
	ExitConstrBlock(c *ConstrBlockContext)

	// ExitSelfInvocation is called when exiting the selfInvocation production.
	ExitSelfInvocation(c *SelfInvocationContext)

	// ExitTopStatSeq is called when exiting the topStatSeq production.
	ExitTopStatSeq(c *TopStatSeqContext)

	// ExitTopStat is called when exiting the topStat production.
	ExitTopStat(c *TopStatContext)

	// ExitPackaging is called when exiting the packaging production.
	ExitPackaging(c *PackagingContext)

	// ExitPackageObject is called when exiting the packageObject production.
	ExitPackageObject(c *PackageObjectContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)
}
