// Code generated from Pddl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pddl // Pddl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PddlListener is a complete listener for a parse tree produced by PddlParser.
type PddlListener interface {
	antlr.ParseTreeListener

	// EnterPddlDoc is called when entering the pddlDoc production.
	EnterPddlDoc(c *PddlDocContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterDomainName is called when entering the domainName production.
	EnterDomainName(c *DomainNameContext)

	// EnterRequireDef is called when entering the requireDef production.
	EnterRequireDef(c *RequireDefContext)

	// EnterTypesDef is called when entering the typesDef production.
	EnterTypesDef(c *TypesDefContext)

	// EnterTypedNameList is called when entering the typedNameList production.
	EnterTypedNameList(c *TypedNameListContext)

	// EnterSingleTypeNameList is called when entering the singleTypeNameList production.
	EnterSingleTypeNameList(c *SingleTypeNameListContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterPrimType is called when entering the primType production.
	EnterPrimType(c *PrimTypeContext)

	// EnterFunctionsDef is called when entering the functionsDef production.
	EnterFunctionsDef(c *FunctionsDefContext)

	// EnterFunctionList is called when entering the functionList production.
	EnterFunctionList(c *FunctionListContext)

	// EnterAtomicFunctionSkeleton is called when entering the atomicFunctionSkeleton production.
	EnterAtomicFunctionSkeleton(c *AtomicFunctionSkeletonContext)

	// EnterFunctionSymbol is called when entering the functionSymbol production.
	EnterFunctionSymbol(c *FunctionSymbolContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterConstantsDef is called when entering the constantsDef production.
	EnterConstantsDef(c *ConstantsDefContext)

	// EnterPredicatesDef is called when entering the predicatesDef production.
	EnterPredicatesDef(c *PredicatesDefContext)

	// EnterAtomicFormulaSkeleton is called when entering the atomicFormulaSkeleton production.
	EnterAtomicFormulaSkeleton(c *AtomicFormulaSkeletonContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterTypedVariableList is called when entering the typedVariableList production.
	EnterTypedVariableList(c *TypedVariableListContext)

	// EnterSingleTypeVarList is called when entering the singleTypeVarList production.
	EnterSingleTypeVarList(c *SingleTypeVarListContext)

	// EnterConstraints is called when entering the constraints production.
	EnterConstraints(c *ConstraintsContext)

	// EnterStructureDef is called when entering the structureDef production.
	EnterStructureDef(c *StructureDefContext)

	// EnterActionDef is called when entering the actionDef production.
	EnterActionDef(c *ActionDefContext)

	// EnterActionSymbol is called when entering the actionSymbol production.
	EnterActionSymbol(c *ActionSymbolContext)

	// EnterActionDefBody is called when entering the actionDefBody production.
	EnterActionDefBody(c *ActionDefBodyContext)

	// EnterGoalDesc is called when entering the goalDesc production.
	EnterGoalDesc(c *GoalDescContext)

	// EnterFComp is called when entering the fComp production.
	EnterFComp(c *FCompContext)

	// EnterAtomicTermFormula is called when entering the atomicTermFormula production.
	EnterAtomicTermFormula(c *AtomicTermFormulaContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterDurativeActionDef is called when entering the durativeActionDef production.
	EnterDurativeActionDef(c *DurativeActionDefContext)

	// EnterDaDefBody is called when entering the daDefBody production.
	EnterDaDefBody(c *DaDefBodyContext)

	// EnterDaGD is called when entering the daGD production.
	EnterDaGD(c *DaGDContext)

	// EnterPrefTimedGD is called when entering the prefTimedGD production.
	EnterPrefTimedGD(c *PrefTimedGDContext)

	// EnterTimedGD is called when entering the timedGD production.
	EnterTimedGD(c *TimedGDContext)

	// EnterTimeSpecifier is called when entering the timeSpecifier production.
	EnterTimeSpecifier(c *TimeSpecifierContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterDerivedDef is called when entering the derivedDef production.
	EnterDerivedDef(c *DerivedDefContext)

	// EnterFExp is called when entering the fExp production.
	EnterFExp(c *FExpContext)

	// EnterFExp2 is called when entering the fExp2 production.
	EnterFExp2(c *FExp2Context)

	// EnterFHead is called when entering the fHead production.
	EnterFHead(c *FHeadContext)

	// EnterEffect is called when entering the effect production.
	EnterEffect(c *EffectContext)

	// EnterCEffect is called when entering the cEffect production.
	EnterCEffect(c *CEffectContext)

	// EnterPEffect is called when entering the pEffect production.
	EnterPEffect(c *PEffectContext)

	// EnterCondEffect is called when entering the condEffect production.
	EnterCondEffect(c *CondEffectContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterBinaryComp is called when entering the binaryComp production.
	EnterBinaryComp(c *BinaryCompContext)

	// EnterAssignOp is called when entering the assignOp production.
	EnterAssignOp(c *AssignOpContext)

	// EnterDurationConstraint is called when entering the durationConstraint production.
	EnterDurationConstraint(c *DurationConstraintContext)

	// EnterSimpleDurationConstraint is called when entering the simpleDurationConstraint production.
	EnterSimpleDurationConstraint(c *SimpleDurationConstraintContext)

	// EnterDurOp is called when entering the durOp production.
	EnterDurOp(c *DurOpContext)

	// EnterDurValue is called when entering the durValue production.
	EnterDurValue(c *DurValueContext)

	// EnterDaEffect is called when entering the daEffect production.
	EnterDaEffect(c *DaEffectContext)

	// EnterTimedEffect is called when entering the timedEffect production.
	EnterTimedEffect(c *TimedEffectContext)

	// EnterFAssignDA is called when entering the fAssignDA production.
	EnterFAssignDA(c *FAssignDAContext)

	// EnterFExpDA is called when entering the fExpDA production.
	EnterFExpDA(c *FExpDAContext)

	// EnterProblem is called when entering the problem production.
	EnterProblem(c *ProblemContext)

	// EnterProblemDecl is called when entering the problemDecl production.
	EnterProblemDecl(c *ProblemDeclContext)

	// EnterProblemDomain is called when entering the problemDomain production.
	EnterProblemDomain(c *ProblemDomainContext)

	// EnterObjectDecl is called when entering the objectDecl production.
	EnterObjectDecl(c *ObjectDeclContext)

	// EnterInit_ is called when entering the init_ production.
	EnterInit_(c *Init_Context)

	// EnterInitEl is called when entering the initEl production.
	EnterInitEl(c *InitElContext)

	// EnterNameLiteral is called when entering the nameLiteral production.
	EnterNameLiteral(c *NameLiteralContext)

	// EnterAtomicNameFormula is called when entering the atomicNameFormula production.
	EnterAtomicNameFormula(c *AtomicNameFormulaContext)

	// EnterGoal is called when entering the goal production.
	EnterGoal(c *GoalContext)

	// EnterProbConstraints is called when entering the probConstraints production.
	EnterProbConstraints(c *ProbConstraintsContext)

	// EnterPrefConGD is called when entering the prefConGD production.
	EnterPrefConGD(c *PrefConGDContext)

	// EnterMetricSpec is called when entering the metricSpec production.
	EnterMetricSpec(c *MetricSpecContext)

	// EnterOptimization is called when entering the optimization production.
	EnterOptimization(c *OptimizationContext)

	// EnterMetricFExp is called when entering the metricFExp production.
	EnterMetricFExp(c *MetricFExpContext)

	// EnterConGD is called when entering the conGD production.
	EnterConGD(c *ConGDContext)

	// ExitPddlDoc is called when exiting the pddlDoc production.
	ExitPddlDoc(c *PddlDocContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitDomainName is called when exiting the domainName production.
	ExitDomainName(c *DomainNameContext)

	// ExitRequireDef is called when exiting the requireDef production.
	ExitRequireDef(c *RequireDefContext)

	// ExitTypesDef is called when exiting the typesDef production.
	ExitTypesDef(c *TypesDefContext)

	// ExitTypedNameList is called when exiting the typedNameList production.
	ExitTypedNameList(c *TypedNameListContext)

	// ExitSingleTypeNameList is called when exiting the singleTypeNameList production.
	ExitSingleTypeNameList(c *SingleTypeNameListContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitPrimType is called when exiting the primType production.
	ExitPrimType(c *PrimTypeContext)

	// ExitFunctionsDef is called when exiting the functionsDef production.
	ExitFunctionsDef(c *FunctionsDefContext)

	// ExitFunctionList is called when exiting the functionList production.
	ExitFunctionList(c *FunctionListContext)

	// ExitAtomicFunctionSkeleton is called when exiting the atomicFunctionSkeleton production.
	ExitAtomicFunctionSkeleton(c *AtomicFunctionSkeletonContext)

	// ExitFunctionSymbol is called when exiting the functionSymbol production.
	ExitFunctionSymbol(c *FunctionSymbolContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitConstantsDef is called when exiting the constantsDef production.
	ExitConstantsDef(c *ConstantsDefContext)

	// ExitPredicatesDef is called when exiting the predicatesDef production.
	ExitPredicatesDef(c *PredicatesDefContext)

	// ExitAtomicFormulaSkeleton is called when exiting the atomicFormulaSkeleton production.
	ExitAtomicFormulaSkeleton(c *AtomicFormulaSkeletonContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitTypedVariableList is called when exiting the typedVariableList production.
	ExitTypedVariableList(c *TypedVariableListContext)

	// ExitSingleTypeVarList is called when exiting the singleTypeVarList production.
	ExitSingleTypeVarList(c *SingleTypeVarListContext)

	// ExitConstraints is called when exiting the constraints production.
	ExitConstraints(c *ConstraintsContext)

	// ExitStructureDef is called when exiting the structureDef production.
	ExitStructureDef(c *StructureDefContext)

	// ExitActionDef is called when exiting the actionDef production.
	ExitActionDef(c *ActionDefContext)

	// ExitActionSymbol is called when exiting the actionSymbol production.
	ExitActionSymbol(c *ActionSymbolContext)

	// ExitActionDefBody is called when exiting the actionDefBody production.
	ExitActionDefBody(c *ActionDefBodyContext)

	// ExitGoalDesc is called when exiting the goalDesc production.
	ExitGoalDesc(c *GoalDescContext)

	// ExitFComp is called when exiting the fComp production.
	ExitFComp(c *FCompContext)

	// ExitAtomicTermFormula is called when exiting the atomicTermFormula production.
	ExitAtomicTermFormula(c *AtomicTermFormulaContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitDurativeActionDef is called when exiting the durativeActionDef production.
	ExitDurativeActionDef(c *DurativeActionDefContext)

	// ExitDaDefBody is called when exiting the daDefBody production.
	ExitDaDefBody(c *DaDefBodyContext)

	// ExitDaGD is called when exiting the daGD production.
	ExitDaGD(c *DaGDContext)

	// ExitPrefTimedGD is called when exiting the prefTimedGD production.
	ExitPrefTimedGD(c *PrefTimedGDContext)

	// ExitTimedGD is called when exiting the timedGD production.
	ExitTimedGD(c *TimedGDContext)

	// ExitTimeSpecifier is called when exiting the timeSpecifier production.
	ExitTimeSpecifier(c *TimeSpecifierContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitDerivedDef is called when exiting the derivedDef production.
	ExitDerivedDef(c *DerivedDefContext)

	// ExitFExp is called when exiting the fExp production.
	ExitFExp(c *FExpContext)

	// ExitFExp2 is called when exiting the fExp2 production.
	ExitFExp2(c *FExp2Context)

	// ExitFHead is called when exiting the fHead production.
	ExitFHead(c *FHeadContext)

	// ExitEffect is called when exiting the effect production.
	ExitEffect(c *EffectContext)

	// ExitCEffect is called when exiting the cEffect production.
	ExitCEffect(c *CEffectContext)

	// ExitPEffect is called when exiting the pEffect production.
	ExitPEffect(c *PEffectContext)

	// ExitCondEffect is called when exiting the condEffect production.
	ExitCondEffect(c *CondEffectContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitBinaryComp is called when exiting the binaryComp production.
	ExitBinaryComp(c *BinaryCompContext)

	// ExitAssignOp is called when exiting the assignOp production.
	ExitAssignOp(c *AssignOpContext)

	// ExitDurationConstraint is called when exiting the durationConstraint production.
	ExitDurationConstraint(c *DurationConstraintContext)

	// ExitSimpleDurationConstraint is called when exiting the simpleDurationConstraint production.
	ExitSimpleDurationConstraint(c *SimpleDurationConstraintContext)

	// ExitDurOp is called when exiting the durOp production.
	ExitDurOp(c *DurOpContext)

	// ExitDurValue is called when exiting the durValue production.
	ExitDurValue(c *DurValueContext)

	// ExitDaEffect is called when exiting the daEffect production.
	ExitDaEffect(c *DaEffectContext)

	// ExitTimedEffect is called when exiting the timedEffect production.
	ExitTimedEffect(c *TimedEffectContext)

	// ExitFAssignDA is called when exiting the fAssignDA production.
	ExitFAssignDA(c *FAssignDAContext)

	// ExitFExpDA is called when exiting the fExpDA production.
	ExitFExpDA(c *FExpDAContext)

	// ExitProblem is called when exiting the problem production.
	ExitProblem(c *ProblemContext)

	// ExitProblemDecl is called when exiting the problemDecl production.
	ExitProblemDecl(c *ProblemDeclContext)

	// ExitProblemDomain is called when exiting the problemDomain production.
	ExitProblemDomain(c *ProblemDomainContext)

	// ExitObjectDecl is called when exiting the objectDecl production.
	ExitObjectDecl(c *ObjectDeclContext)

	// ExitInit_ is called when exiting the init_ production.
	ExitInit_(c *Init_Context)

	// ExitInitEl is called when exiting the initEl production.
	ExitInitEl(c *InitElContext)

	// ExitNameLiteral is called when exiting the nameLiteral production.
	ExitNameLiteral(c *NameLiteralContext)

	// ExitAtomicNameFormula is called when exiting the atomicNameFormula production.
	ExitAtomicNameFormula(c *AtomicNameFormulaContext)

	// ExitGoal is called when exiting the goal production.
	ExitGoal(c *GoalContext)

	// ExitProbConstraints is called when exiting the probConstraints production.
	ExitProbConstraints(c *ProbConstraintsContext)

	// ExitPrefConGD is called when exiting the prefConGD production.
	ExitPrefConGD(c *PrefConGDContext)

	// ExitMetricSpec is called when exiting the metricSpec production.
	ExitMetricSpec(c *MetricSpecContext)

	// ExitOptimization is called when exiting the optimization production.
	ExitOptimization(c *OptimizationContext)

	// ExitMetricFExp is called when exiting the metricFExp production.
	ExitMetricFExp(c *MetricFExpContext)

	// ExitConGD is called when exiting the conGD production.
	ExitConGD(c *ConGDContext)
}
