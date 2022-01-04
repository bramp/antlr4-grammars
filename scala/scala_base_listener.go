// Code generated from Scala.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scala // Scala
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseScalaListener is a complete listener for a parse tree produced by ScalaParser.
type BaseScalaListener struct{}

var _ ScalaListener = &BaseScalaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseScalaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseScalaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseScalaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseScalaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseScalaListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseScalaListener) ExitLiteral(ctx *LiteralContext) {}

// EnterQualId is called when production qualId is entered.
func (s *BaseScalaListener) EnterQualId(ctx *QualIdContext) {}

// ExitQualId is called when production qualId is exited.
func (s *BaseScalaListener) ExitQualId(ctx *QualIdContext) {}

// EnterIds is called when production ids is entered.
func (s *BaseScalaListener) EnterIds(ctx *IdsContext) {}

// ExitIds is called when production ids is exited.
func (s *BaseScalaListener) ExitIds(ctx *IdsContext) {}

// EnterStableId is called when production stableId is entered.
func (s *BaseScalaListener) EnterStableId(ctx *StableIdContext) {}

// ExitStableId is called when production stableId is exited.
func (s *BaseScalaListener) ExitStableId(ctx *StableIdContext) {}

// EnterClassQualifier is called when production classQualifier is entered.
func (s *BaseScalaListener) EnterClassQualifier(ctx *ClassQualifierContext) {}

// ExitClassQualifier is called when production classQualifier is exited.
func (s *BaseScalaListener) ExitClassQualifier(ctx *ClassQualifierContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseScalaListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseScalaListener) ExitType_(ctx *Type_Context) {}

// EnterFunctionArgTypes is called when production functionArgTypes is entered.
func (s *BaseScalaListener) EnterFunctionArgTypes(ctx *FunctionArgTypesContext) {}

// ExitFunctionArgTypes is called when production functionArgTypes is exited.
func (s *BaseScalaListener) ExitFunctionArgTypes(ctx *FunctionArgTypesContext) {}

// EnterExistentialClause is called when production existentialClause is entered.
func (s *BaseScalaListener) EnterExistentialClause(ctx *ExistentialClauseContext) {}

// ExitExistentialClause is called when production existentialClause is exited.
func (s *BaseScalaListener) ExitExistentialClause(ctx *ExistentialClauseContext) {}

// EnterExistentialDcl is called when production existentialDcl is entered.
func (s *BaseScalaListener) EnterExistentialDcl(ctx *ExistentialDclContext) {}

// ExitExistentialDcl is called when production existentialDcl is exited.
func (s *BaseScalaListener) ExitExistentialDcl(ctx *ExistentialDclContext) {}

// EnterInfixType is called when production infixType is entered.
func (s *BaseScalaListener) EnterInfixType(ctx *InfixTypeContext) {}

// ExitInfixType is called when production infixType is exited.
func (s *BaseScalaListener) ExitInfixType(ctx *InfixTypeContext) {}

// EnterCompoundType is called when production compoundType is entered.
func (s *BaseScalaListener) EnterCompoundType(ctx *CompoundTypeContext) {}

// ExitCompoundType is called when production compoundType is exited.
func (s *BaseScalaListener) ExitCompoundType(ctx *CompoundTypeContext) {}

// EnterAnnotType is called when production annotType is entered.
func (s *BaseScalaListener) EnterAnnotType(ctx *AnnotTypeContext) {}

// ExitAnnotType is called when production annotType is exited.
func (s *BaseScalaListener) ExitAnnotType(ctx *AnnotTypeContext) {}

// EnterSimpleType is called when production simpleType is entered.
func (s *BaseScalaListener) EnterSimpleType(ctx *SimpleTypeContext) {}

// ExitSimpleType is called when production simpleType is exited.
func (s *BaseScalaListener) ExitSimpleType(ctx *SimpleTypeContext) {}

// EnterTypeArgs is called when production typeArgs is entered.
func (s *BaseScalaListener) EnterTypeArgs(ctx *TypeArgsContext) {}

// ExitTypeArgs is called when production typeArgs is exited.
func (s *BaseScalaListener) ExitTypeArgs(ctx *TypeArgsContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseScalaListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseScalaListener) ExitTypes(ctx *TypesContext) {}

// EnterRefinement is called when production refinement is entered.
func (s *BaseScalaListener) EnterRefinement(ctx *RefinementContext) {}

// ExitRefinement is called when production refinement is exited.
func (s *BaseScalaListener) ExitRefinement(ctx *RefinementContext) {}

// EnterRefineStat is called when production refineStat is entered.
func (s *BaseScalaListener) EnterRefineStat(ctx *RefineStatContext) {}

// ExitRefineStat is called when production refineStat is exited.
func (s *BaseScalaListener) ExitRefineStat(ctx *RefineStatContext) {}

// EnterTypePat is called when production typePat is entered.
func (s *BaseScalaListener) EnterTypePat(ctx *TypePatContext) {}

// ExitTypePat is called when production typePat is exited.
func (s *BaseScalaListener) ExitTypePat(ctx *TypePatContext) {}

// EnterAscription is called when production ascription is entered.
func (s *BaseScalaListener) EnterAscription(ctx *AscriptionContext) {}

// ExitAscription is called when production ascription is exited.
func (s *BaseScalaListener) ExitAscription(ctx *AscriptionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseScalaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseScalaListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr1 is called when production expr1 is entered.
func (s *BaseScalaListener) EnterExpr1(ctx *Expr1Context) {}

// ExitExpr1 is called when production expr1 is exited.
func (s *BaseScalaListener) ExitExpr1(ctx *Expr1Context) {}

// EnterPrefixDef is called when production prefixDef is entered.
func (s *BaseScalaListener) EnterPrefixDef(ctx *PrefixDefContext) {}

// ExitPrefixDef is called when production prefixDef is exited.
func (s *BaseScalaListener) ExitPrefixDef(ctx *PrefixDefContext) {}

// EnterPostfixExpr is called when production postfixExpr is entered.
func (s *BaseScalaListener) EnterPostfixExpr(ctx *PostfixExprContext) {}

// ExitPostfixExpr is called when production postfixExpr is exited.
func (s *BaseScalaListener) ExitPostfixExpr(ctx *PostfixExprContext) {}

// EnterInfixExpr is called when production infixExpr is entered.
func (s *BaseScalaListener) EnterInfixExpr(ctx *InfixExprContext) {}

// ExitInfixExpr is called when production infixExpr is exited.
func (s *BaseScalaListener) ExitInfixExpr(ctx *InfixExprContext) {}

// EnterPrefixExpr is called when production prefixExpr is entered.
func (s *BaseScalaListener) EnterPrefixExpr(ctx *PrefixExprContext) {}

// ExitPrefixExpr is called when production prefixExpr is exited.
func (s *BaseScalaListener) ExitPrefixExpr(ctx *PrefixExprContext) {}

// EnterSimpleExpr is called when production simpleExpr is entered.
func (s *BaseScalaListener) EnterSimpleExpr(ctx *SimpleExprContext) {}

// ExitSimpleExpr is called when production simpleExpr is exited.
func (s *BaseScalaListener) ExitSimpleExpr(ctx *SimpleExprContext) {}

// EnterSimpleExpr1 is called when production simpleExpr1 is entered.
func (s *BaseScalaListener) EnterSimpleExpr1(ctx *SimpleExpr1Context) {}

// ExitSimpleExpr1 is called when production simpleExpr1 is exited.
func (s *BaseScalaListener) ExitSimpleExpr1(ctx *SimpleExpr1Context) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseScalaListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseScalaListener) ExitExprs(ctx *ExprsContext) {}

// EnterArgumentExprs is called when production argumentExprs is entered.
func (s *BaseScalaListener) EnterArgumentExprs(ctx *ArgumentExprsContext) {}

// ExitArgumentExprs is called when production argumentExprs is exited.
func (s *BaseScalaListener) ExitArgumentExprs(ctx *ArgumentExprsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseScalaListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseScalaListener) ExitArgs(ctx *ArgsContext) {}

// EnterBlockExpr is called when production blockExpr is entered.
func (s *BaseScalaListener) EnterBlockExpr(ctx *BlockExprContext) {}

// ExitBlockExpr is called when production blockExpr is exited.
func (s *BaseScalaListener) ExitBlockExpr(ctx *BlockExprContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseScalaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseScalaListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStat is called when production blockStat is entered.
func (s *BaseScalaListener) EnterBlockStat(ctx *BlockStatContext) {}

// ExitBlockStat is called when production blockStat is exited.
func (s *BaseScalaListener) ExitBlockStat(ctx *BlockStatContext) {}

// EnterResultExpr is called when production resultExpr is entered.
func (s *BaseScalaListener) EnterResultExpr(ctx *ResultExprContext) {}

// ExitResultExpr is called when production resultExpr is exited.
func (s *BaseScalaListener) ExitResultExpr(ctx *ResultExprContext) {}

// EnterEnumerators is called when production enumerators is entered.
func (s *BaseScalaListener) EnterEnumerators(ctx *EnumeratorsContext) {}

// ExitEnumerators is called when production enumerators is exited.
func (s *BaseScalaListener) ExitEnumerators(ctx *EnumeratorsContext) {}

// EnterGenerator is called when production generator is entered.
func (s *BaseScalaListener) EnterGenerator(ctx *GeneratorContext) {}

// ExitGenerator is called when production generator is exited.
func (s *BaseScalaListener) ExitGenerator(ctx *GeneratorContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseScalaListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseScalaListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseScalaListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseScalaListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterGuard_ is called when production guard_ is entered.
func (s *BaseScalaListener) EnterGuard_(ctx *Guard_Context) {}

// ExitGuard_ is called when production guard_ is exited.
func (s *BaseScalaListener) ExitGuard_(ctx *Guard_Context) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseScalaListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseScalaListener) ExitPattern(ctx *PatternContext) {}

// EnterPattern1 is called when production pattern1 is entered.
func (s *BaseScalaListener) EnterPattern1(ctx *Pattern1Context) {}

// ExitPattern1 is called when production pattern1 is exited.
func (s *BaseScalaListener) ExitPattern1(ctx *Pattern1Context) {}

// EnterPattern2 is called when production pattern2 is entered.
func (s *BaseScalaListener) EnterPattern2(ctx *Pattern2Context) {}

// ExitPattern2 is called when production pattern2 is exited.
func (s *BaseScalaListener) ExitPattern2(ctx *Pattern2Context) {}

// EnterPattern3 is called when production pattern3 is entered.
func (s *BaseScalaListener) EnterPattern3(ctx *Pattern3Context) {}

// ExitPattern3 is called when production pattern3 is exited.
func (s *BaseScalaListener) ExitPattern3(ctx *Pattern3Context) {}

// EnterSimplePattern is called when production simplePattern is entered.
func (s *BaseScalaListener) EnterSimplePattern(ctx *SimplePatternContext) {}

// ExitSimplePattern is called when production simplePattern is exited.
func (s *BaseScalaListener) ExitSimplePattern(ctx *SimplePatternContext) {}

// EnterPatterns is called when production patterns is entered.
func (s *BaseScalaListener) EnterPatterns(ctx *PatternsContext) {}

// ExitPatterns is called when production patterns is exited.
func (s *BaseScalaListener) ExitPatterns(ctx *PatternsContext) {}

// EnterTypeParamClause is called when production typeParamClause is entered.
func (s *BaseScalaListener) EnterTypeParamClause(ctx *TypeParamClauseContext) {}

// ExitTypeParamClause is called when production typeParamClause is exited.
func (s *BaseScalaListener) ExitTypeParamClause(ctx *TypeParamClauseContext) {}

// EnterFunTypeParamClause is called when production funTypeParamClause is entered.
func (s *BaseScalaListener) EnterFunTypeParamClause(ctx *FunTypeParamClauseContext) {}

// ExitFunTypeParamClause is called when production funTypeParamClause is exited.
func (s *BaseScalaListener) ExitFunTypeParamClause(ctx *FunTypeParamClauseContext) {}

// EnterVariantTypeParam is called when production variantTypeParam is entered.
func (s *BaseScalaListener) EnterVariantTypeParam(ctx *VariantTypeParamContext) {}

// ExitVariantTypeParam is called when production variantTypeParam is exited.
func (s *BaseScalaListener) ExitVariantTypeParam(ctx *VariantTypeParamContext) {}

// EnterTypeParam is called when production typeParam is entered.
func (s *BaseScalaListener) EnterTypeParam(ctx *TypeParamContext) {}

// ExitTypeParam is called when production typeParam is exited.
func (s *BaseScalaListener) ExitTypeParam(ctx *TypeParamContext) {}

// EnterParamClauses is called when production paramClauses is entered.
func (s *BaseScalaListener) EnterParamClauses(ctx *ParamClausesContext) {}

// ExitParamClauses is called when production paramClauses is exited.
func (s *BaseScalaListener) ExitParamClauses(ctx *ParamClausesContext) {}

// EnterParamClause is called when production paramClause is entered.
func (s *BaseScalaListener) EnterParamClause(ctx *ParamClauseContext) {}

// ExitParamClause is called when production paramClause is exited.
func (s *BaseScalaListener) ExitParamClause(ctx *ParamClauseContext) {}

// EnterParams is called when production params is entered.
func (s *BaseScalaListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseScalaListener) ExitParams(ctx *ParamsContext) {}

// EnterParam is called when production param is entered.
func (s *BaseScalaListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseScalaListener) ExitParam(ctx *ParamContext) {}

// EnterParamType is called when production paramType is entered.
func (s *BaseScalaListener) EnterParamType(ctx *ParamTypeContext) {}

// ExitParamType is called when production paramType is exited.
func (s *BaseScalaListener) ExitParamType(ctx *ParamTypeContext) {}

// EnterClassParamClauses is called when production classParamClauses is entered.
func (s *BaseScalaListener) EnterClassParamClauses(ctx *ClassParamClausesContext) {}

// ExitClassParamClauses is called when production classParamClauses is exited.
func (s *BaseScalaListener) ExitClassParamClauses(ctx *ClassParamClausesContext) {}

// EnterClassParamClause is called when production classParamClause is entered.
func (s *BaseScalaListener) EnterClassParamClause(ctx *ClassParamClauseContext) {}

// ExitClassParamClause is called when production classParamClause is exited.
func (s *BaseScalaListener) ExitClassParamClause(ctx *ClassParamClauseContext) {}

// EnterClassParams is called when production classParams is entered.
func (s *BaseScalaListener) EnterClassParams(ctx *ClassParamsContext) {}

// ExitClassParams is called when production classParams is exited.
func (s *BaseScalaListener) ExitClassParams(ctx *ClassParamsContext) {}

// EnterClassParam is called when production classParam is entered.
func (s *BaseScalaListener) EnterClassParam(ctx *ClassParamContext) {}

// ExitClassParam is called when production classParam is exited.
func (s *BaseScalaListener) ExitClassParam(ctx *ClassParamContext) {}

// EnterBindings is called when production bindings is entered.
func (s *BaseScalaListener) EnterBindings(ctx *BindingsContext) {}

// ExitBindings is called when production bindings is exited.
func (s *BaseScalaListener) ExitBindings(ctx *BindingsContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseScalaListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseScalaListener) ExitBinding(ctx *BindingContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseScalaListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseScalaListener) ExitModifier(ctx *ModifierContext) {}

// EnterLocalModifier is called when production localModifier is entered.
func (s *BaseScalaListener) EnterLocalModifier(ctx *LocalModifierContext) {}

// ExitLocalModifier is called when production localModifier is exited.
func (s *BaseScalaListener) ExitLocalModifier(ctx *LocalModifierContext) {}

// EnterAccessModifier is called when production accessModifier is entered.
func (s *BaseScalaListener) EnterAccessModifier(ctx *AccessModifierContext) {}

// ExitAccessModifier is called when production accessModifier is exited.
func (s *BaseScalaListener) ExitAccessModifier(ctx *AccessModifierContext) {}

// EnterAccessQualifier is called when production accessQualifier is entered.
func (s *BaseScalaListener) EnterAccessQualifier(ctx *AccessQualifierContext) {}

// ExitAccessQualifier is called when production accessQualifier is exited.
func (s *BaseScalaListener) ExitAccessQualifier(ctx *AccessQualifierContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseScalaListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseScalaListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterConstrAnnotation is called when production constrAnnotation is entered.
func (s *BaseScalaListener) EnterConstrAnnotation(ctx *ConstrAnnotationContext) {}

// ExitConstrAnnotation is called when production constrAnnotation is exited.
func (s *BaseScalaListener) ExitConstrAnnotation(ctx *ConstrAnnotationContext) {}

// EnterTemplateBody is called when production templateBody is entered.
func (s *BaseScalaListener) EnterTemplateBody(ctx *TemplateBodyContext) {}

// ExitTemplateBody is called when production templateBody is exited.
func (s *BaseScalaListener) ExitTemplateBody(ctx *TemplateBodyContext) {}

// EnterTemplateStat is called when production templateStat is entered.
func (s *BaseScalaListener) EnterTemplateStat(ctx *TemplateStatContext) {}

// ExitTemplateStat is called when production templateStat is exited.
func (s *BaseScalaListener) ExitTemplateStat(ctx *TemplateStatContext) {}

// EnterSelfType is called when production selfType is entered.
func (s *BaseScalaListener) EnterSelfType(ctx *SelfTypeContext) {}

// ExitSelfType is called when production selfType is exited.
func (s *BaseScalaListener) ExitSelfType(ctx *SelfTypeContext) {}

// EnterImport_ is called when production import_ is entered.
func (s *BaseScalaListener) EnterImport_(ctx *Import_Context) {}

// ExitImport_ is called when production import_ is exited.
func (s *BaseScalaListener) ExitImport_(ctx *Import_Context) {}

// EnterImportExpr is called when production importExpr is entered.
func (s *BaseScalaListener) EnterImportExpr(ctx *ImportExprContext) {}

// ExitImportExpr is called when production importExpr is exited.
func (s *BaseScalaListener) ExitImportExpr(ctx *ImportExprContext) {}

// EnterImportSelectors is called when production importSelectors is entered.
func (s *BaseScalaListener) EnterImportSelectors(ctx *ImportSelectorsContext) {}

// ExitImportSelectors is called when production importSelectors is exited.
func (s *BaseScalaListener) ExitImportSelectors(ctx *ImportSelectorsContext) {}

// EnterImportSelector is called when production importSelector is entered.
func (s *BaseScalaListener) EnterImportSelector(ctx *ImportSelectorContext) {}

// ExitImportSelector is called when production importSelector is exited.
func (s *BaseScalaListener) ExitImportSelector(ctx *ImportSelectorContext) {}

// EnterDcl is called when production dcl is entered.
func (s *BaseScalaListener) EnterDcl(ctx *DclContext) {}

// ExitDcl is called when production dcl is exited.
func (s *BaseScalaListener) ExitDcl(ctx *DclContext) {}

// EnterValDcl is called when production valDcl is entered.
func (s *BaseScalaListener) EnterValDcl(ctx *ValDclContext) {}

// ExitValDcl is called when production valDcl is exited.
func (s *BaseScalaListener) ExitValDcl(ctx *ValDclContext) {}

// EnterVarDcl is called when production varDcl is entered.
func (s *BaseScalaListener) EnterVarDcl(ctx *VarDclContext) {}

// ExitVarDcl is called when production varDcl is exited.
func (s *BaseScalaListener) ExitVarDcl(ctx *VarDclContext) {}

// EnterFunDcl is called when production funDcl is entered.
func (s *BaseScalaListener) EnterFunDcl(ctx *FunDclContext) {}

// ExitFunDcl is called when production funDcl is exited.
func (s *BaseScalaListener) ExitFunDcl(ctx *FunDclContext) {}

// EnterFunSig is called when production funSig is entered.
func (s *BaseScalaListener) EnterFunSig(ctx *FunSigContext) {}

// ExitFunSig is called when production funSig is exited.
func (s *BaseScalaListener) ExitFunSig(ctx *FunSigContext) {}

// EnterTypeDcl is called when production typeDcl is entered.
func (s *BaseScalaListener) EnterTypeDcl(ctx *TypeDclContext) {}

// ExitTypeDcl is called when production typeDcl is exited.
func (s *BaseScalaListener) ExitTypeDcl(ctx *TypeDclContext) {}

// EnterPatVarDef is called when production patVarDef is entered.
func (s *BaseScalaListener) EnterPatVarDef(ctx *PatVarDefContext) {}

// ExitPatVarDef is called when production patVarDef is exited.
func (s *BaseScalaListener) ExitPatVarDef(ctx *PatVarDefContext) {}

// EnterDef_ is called when production def_ is entered.
func (s *BaseScalaListener) EnterDef_(ctx *Def_Context) {}

// ExitDef_ is called when production def_ is exited.
func (s *BaseScalaListener) ExitDef_(ctx *Def_Context) {}

// EnterPatDef is called when production patDef is entered.
func (s *BaseScalaListener) EnterPatDef(ctx *PatDefContext) {}

// ExitPatDef is called when production patDef is exited.
func (s *BaseScalaListener) ExitPatDef(ctx *PatDefContext) {}

// EnterVarDef is called when production varDef is entered.
func (s *BaseScalaListener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production varDef is exited.
func (s *BaseScalaListener) ExitVarDef(ctx *VarDefContext) {}

// EnterFunDef is called when production funDef is entered.
func (s *BaseScalaListener) EnterFunDef(ctx *FunDefContext) {}

// ExitFunDef is called when production funDef is exited.
func (s *BaseScalaListener) ExitFunDef(ctx *FunDefContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseScalaListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseScalaListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterTmplDef is called when production tmplDef is entered.
func (s *BaseScalaListener) EnterTmplDef(ctx *TmplDefContext) {}

// ExitTmplDef is called when production tmplDef is exited.
func (s *BaseScalaListener) ExitTmplDef(ctx *TmplDefContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseScalaListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseScalaListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterTraitDef is called when production traitDef is entered.
func (s *BaseScalaListener) EnterTraitDef(ctx *TraitDefContext) {}

// ExitTraitDef is called when production traitDef is exited.
func (s *BaseScalaListener) ExitTraitDef(ctx *TraitDefContext) {}

// EnterObjectDef is called when production objectDef is entered.
func (s *BaseScalaListener) EnterObjectDef(ctx *ObjectDefContext) {}

// ExitObjectDef is called when production objectDef is exited.
func (s *BaseScalaListener) ExitObjectDef(ctx *ObjectDefContext) {}

// EnterClassTemplateOpt is called when production classTemplateOpt is entered.
func (s *BaseScalaListener) EnterClassTemplateOpt(ctx *ClassTemplateOptContext) {}

// ExitClassTemplateOpt is called when production classTemplateOpt is exited.
func (s *BaseScalaListener) ExitClassTemplateOpt(ctx *ClassTemplateOptContext) {}

// EnterTraitTemplateOpt is called when production traitTemplateOpt is entered.
func (s *BaseScalaListener) EnterTraitTemplateOpt(ctx *TraitTemplateOptContext) {}

// ExitTraitTemplateOpt is called when production traitTemplateOpt is exited.
func (s *BaseScalaListener) ExitTraitTemplateOpt(ctx *TraitTemplateOptContext) {}

// EnterClassTemplate is called when production classTemplate is entered.
func (s *BaseScalaListener) EnterClassTemplate(ctx *ClassTemplateContext) {}

// ExitClassTemplate is called when production classTemplate is exited.
func (s *BaseScalaListener) ExitClassTemplate(ctx *ClassTemplateContext) {}

// EnterTraitTemplate is called when production traitTemplate is entered.
func (s *BaseScalaListener) EnterTraitTemplate(ctx *TraitTemplateContext) {}

// ExitTraitTemplate is called when production traitTemplate is exited.
func (s *BaseScalaListener) ExitTraitTemplate(ctx *TraitTemplateContext) {}

// EnterClassParents is called when production classParents is entered.
func (s *BaseScalaListener) EnterClassParents(ctx *ClassParentsContext) {}

// ExitClassParents is called when production classParents is exited.
func (s *BaseScalaListener) ExitClassParents(ctx *ClassParentsContext) {}

// EnterTraitParents is called when production traitParents is entered.
func (s *BaseScalaListener) EnterTraitParents(ctx *TraitParentsContext) {}

// ExitTraitParents is called when production traitParents is exited.
func (s *BaseScalaListener) ExitTraitParents(ctx *TraitParentsContext) {}

// EnterConstr is called when production constr is entered.
func (s *BaseScalaListener) EnterConstr(ctx *ConstrContext) {}

// ExitConstr is called when production constr is exited.
func (s *BaseScalaListener) ExitConstr(ctx *ConstrContext) {}

// EnterEarlyDefs is called when production earlyDefs is entered.
func (s *BaseScalaListener) EnterEarlyDefs(ctx *EarlyDefsContext) {}

// ExitEarlyDefs is called when production earlyDefs is exited.
func (s *BaseScalaListener) ExitEarlyDefs(ctx *EarlyDefsContext) {}

// EnterEarlyDef is called when production earlyDef is entered.
func (s *BaseScalaListener) EnterEarlyDef(ctx *EarlyDefContext) {}

// ExitEarlyDef is called when production earlyDef is exited.
func (s *BaseScalaListener) ExitEarlyDef(ctx *EarlyDefContext) {}

// EnterConstrExpr is called when production constrExpr is entered.
func (s *BaseScalaListener) EnterConstrExpr(ctx *ConstrExprContext) {}

// ExitConstrExpr is called when production constrExpr is exited.
func (s *BaseScalaListener) ExitConstrExpr(ctx *ConstrExprContext) {}

// EnterConstrBlock is called when production constrBlock is entered.
func (s *BaseScalaListener) EnterConstrBlock(ctx *ConstrBlockContext) {}

// ExitConstrBlock is called when production constrBlock is exited.
func (s *BaseScalaListener) ExitConstrBlock(ctx *ConstrBlockContext) {}

// EnterSelfInvocation is called when production selfInvocation is entered.
func (s *BaseScalaListener) EnterSelfInvocation(ctx *SelfInvocationContext) {}

// ExitSelfInvocation is called when production selfInvocation is exited.
func (s *BaseScalaListener) ExitSelfInvocation(ctx *SelfInvocationContext) {}

// EnterTopStatSeq is called when production topStatSeq is entered.
func (s *BaseScalaListener) EnterTopStatSeq(ctx *TopStatSeqContext) {}

// ExitTopStatSeq is called when production topStatSeq is exited.
func (s *BaseScalaListener) ExitTopStatSeq(ctx *TopStatSeqContext) {}

// EnterTopStat is called when production topStat is entered.
func (s *BaseScalaListener) EnterTopStat(ctx *TopStatContext) {}

// ExitTopStat is called when production topStat is exited.
func (s *BaseScalaListener) ExitTopStat(ctx *TopStatContext) {}

// EnterPackaging is called when production packaging is entered.
func (s *BaseScalaListener) EnterPackaging(ctx *PackagingContext) {}

// ExitPackaging is called when production packaging is exited.
func (s *BaseScalaListener) ExitPackaging(ctx *PackagingContext) {}

// EnterPackageObject is called when production packageObject is entered.
func (s *BaseScalaListener) EnterPackageObject(ctx *PackageObjectContext) {}

// ExitPackageObject is called when production packageObject is exited.
func (s *BaseScalaListener) ExitPackageObject(ctx *PackageObjectContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseScalaListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseScalaListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}
