// Code generated from Pddl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pddl // Pddl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePddlListener is a complete listener for a parse tree produced by PddlParser.
type BasePddlListener struct{}

var _ PddlListener = &BasePddlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePddlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePddlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePddlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePddlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPddlDoc is called when production pddlDoc is entered.
func (s *BasePddlListener) EnterPddlDoc(ctx *PddlDocContext) {}

// ExitPddlDoc is called when production pddlDoc is exited.
func (s *BasePddlListener) ExitPddlDoc(ctx *PddlDocContext) {}

// EnterDomain is called when production domain is entered.
func (s *BasePddlListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BasePddlListener) ExitDomain(ctx *DomainContext) {}

// EnterDomainName is called when production domainName is entered.
func (s *BasePddlListener) EnterDomainName(ctx *DomainNameContext) {}

// ExitDomainName is called when production domainName is exited.
func (s *BasePddlListener) ExitDomainName(ctx *DomainNameContext) {}

// EnterRequireDef is called when production requireDef is entered.
func (s *BasePddlListener) EnterRequireDef(ctx *RequireDefContext) {}

// ExitRequireDef is called when production requireDef is exited.
func (s *BasePddlListener) ExitRequireDef(ctx *RequireDefContext) {}

// EnterTypesDef is called when production typesDef is entered.
func (s *BasePddlListener) EnterTypesDef(ctx *TypesDefContext) {}

// ExitTypesDef is called when production typesDef is exited.
func (s *BasePddlListener) ExitTypesDef(ctx *TypesDefContext) {}

// EnterTypedNameList is called when production typedNameList is entered.
func (s *BasePddlListener) EnterTypedNameList(ctx *TypedNameListContext) {}

// ExitTypedNameList is called when production typedNameList is exited.
func (s *BasePddlListener) ExitTypedNameList(ctx *TypedNameListContext) {}

// EnterSingleTypeNameList is called when production singleTypeNameList is entered.
func (s *BasePddlListener) EnterSingleTypeNameList(ctx *SingleTypeNameListContext) {}

// ExitSingleTypeNameList is called when production singleTypeNameList is exited.
func (s *BasePddlListener) ExitSingleTypeNameList(ctx *SingleTypeNameListContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasePddlListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasePddlListener) ExitType_(ctx *Type_Context) {}

// EnterPrimType is called when production primType is entered.
func (s *BasePddlListener) EnterPrimType(ctx *PrimTypeContext) {}

// ExitPrimType is called when production primType is exited.
func (s *BasePddlListener) ExitPrimType(ctx *PrimTypeContext) {}

// EnterFunctionsDef is called when production functionsDef is entered.
func (s *BasePddlListener) EnterFunctionsDef(ctx *FunctionsDefContext) {}

// ExitFunctionsDef is called when production functionsDef is exited.
func (s *BasePddlListener) ExitFunctionsDef(ctx *FunctionsDefContext) {}

// EnterFunctionList is called when production functionList is entered.
func (s *BasePddlListener) EnterFunctionList(ctx *FunctionListContext) {}

// ExitFunctionList is called when production functionList is exited.
func (s *BasePddlListener) ExitFunctionList(ctx *FunctionListContext) {}

// EnterAtomicFunctionSkeleton is called when production atomicFunctionSkeleton is entered.
func (s *BasePddlListener) EnterAtomicFunctionSkeleton(ctx *AtomicFunctionSkeletonContext) {}

// ExitAtomicFunctionSkeleton is called when production atomicFunctionSkeleton is exited.
func (s *BasePddlListener) ExitAtomicFunctionSkeleton(ctx *AtomicFunctionSkeletonContext) {}

// EnterFunctionSymbol is called when production functionSymbol is entered.
func (s *BasePddlListener) EnterFunctionSymbol(ctx *FunctionSymbolContext) {}

// ExitFunctionSymbol is called when production functionSymbol is exited.
func (s *BasePddlListener) ExitFunctionSymbol(ctx *FunctionSymbolContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BasePddlListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BasePddlListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterConstantsDef is called when production constantsDef is entered.
func (s *BasePddlListener) EnterConstantsDef(ctx *ConstantsDefContext) {}

// ExitConstantsDef is called when production constantsDef is exited.
func (s *BasePddlListener) ExitConstantsDef(ctx *ConstantsDefContext) {}

// EnterPredicatesDef is called when production predicatesDef is entered.
func (s *BasePddlListener) EnterPredicatesDef(ctx *PredicatesDefContext) {}

// ExitPredicatesDef is called when production predicatesDef is exited.
func (s *BasePddlListener) ExitPredicatesDef(ctx *PredicatesDefContext) {}

// EnterAtomicFormulaSkeleton is called when production atomicFormulaSkeleton is entered.
func (s *BasePddlListener) EnterAtomicFormulaSkeleton(ctx *AtomicFormulaSkeletonContext) {}

// ExitAtomicFormulaSkeleton is called when production atomicFormulaSkeleton is exited.
func (s *BasePddlListener) ExitAtomicFormulaSkeleton(ctx *AtomicFormulaSkeletonContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasePddlListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasePddlListener) ExitPredicate(ctx *PredicateContext) {}

// EnterTypedVariableList is called when production typedVariableList is entered.
func (s *BasePddlListener) EnterTypedVariableList(ctx *TypedVariableListContext) {}

// ExitTypedVariableList is called when production typedVariableList is exited.
func (s *BasePddlListener) ExitTypedVariableList(ctx *TypedVariableListContext) {}

// EnterSingleTypeVarList is called when production singleTypeVarList is entered.
func (s *BasePddlListener) EnterSingleTypeVarList(ctx *SingleTypeVarListContext) {}

// ExitSingleTypeVarList is called when production singleTypeVarList is exited.
func (s *BasePddlListener) ExitSingleTypeVarList(ctx *SingleTypeVarListContext) {}

// EnterConstraints is called when production constraints is entered.
func (s *BasePddlListener) EnterConstraints(ctx *ConstraintsContext) {}

// ExitConstraints is called when production constraints is exited.
func (s *BasePddlListener) ExitConstraints(ctx *ConstraintsContext) {}

// EnterStructureDef is called when production structureDef is entered.
func (s *BasePddlListener) EnterStructureDef(ctx *StructureDefContext) {}

// ExitStructureDef is called when production structureDef is exited.
func (s *BasePddlListener) ExitStructureDef(ctx *StructureDefContext) {}

// EnterActionDef is called when production actionDef is entered.
func (s *BasePddlListener) EnterActionDef(ctx *ActionDefContext) {}

// ExitActionDef is called when production actionDef is exited.
func (s *BasePddlListener) ExitActionDef(ctx *ActionDefContext) {}

// EnterActionSymbol is called when production actionSymbol is entered.
func (s *BasePddlListener) EnterActionSymbol(ctx *ActionSymbolContext) {}

// ExitActionSymbol is called when production actionSymbol is exited.
func (s *BasePddlListener) ExitActionSymbol(ctx *ActionSymbolContext) {}

// EnterActionDefBody is called when production actionDefBody is entered.
func (s *BasePddlListener) EnterActionDefBody(ctx *ActionDefBodyContext) {}

// ExitActionDefBody is called when production actionDefBody is exited.
func (s *BasePddlListener) ExitActionDefBody(ctx *ActionDefBodyContext) {}

// EnterGoalDesc is called when production goalDesc is entered.
func (s *BasePddlListener) EnterGoalDesc(ctx *GoalDescContext) {}

// ExitGoalDesc is called when production goalDesc is exited.
func (s *BasePddlListener) ExitGoalDesc(ctx *GoalDescContext) {}

// EnterFComp is called when production fComp is entered.
func (s *BasePddlListener) EnterFComp(ctx *FCompContext) {}

// ExitFComp is called when production fComp is exited.
func (s *BasePddlListener) ExitFComp(ctx *FCompContext) {}

// EnterAtomicTermFormula is called when production atomicTermFormula is entered.
func (s *BasePddlListener) EnterAtomicTermFormula(ctx *AtomicTermFormulaContext) {}

// ExitAtomicTermFormula is called when production atomicTermFormula is exited.
func (s *BasePddlListener) ExitAtomicTermFormula(ctx *AtomicTermFormulaContext) {}

// EnterTerm is called when production term is entered.
func (s *BasePddlListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasePddlListener) ExitTerm(ctx *TermContext) {}

// EnterDurativeActionDef is called when production durativeActionDef is entered.
func (s *BasePddlListener) EnterDurativeActionDef(ctx *DurativeActionDefContext) {}

// ExitDurativeActionDef is called when production durativeActionDef is exited.
func (s *BasePddlListener) ExitDurativeActionDef(ctx *DurativeActionDefContext) {}

// EnterDaDefBody is called when production daDefBody is entered.
func (s *BasePddlListener) EnterDaDefBody(ctx *DaDefBodyContext) {}

// ExitDaDefBody is called when production daDefBody is exited.
func (s *BasePddlListener) ExitDaDefBody(ctx *DaDefBodyContext) {}

// EnterDaGD is called when production daGD is entered.
func (s *BasePddlListener) EnterDaGD(ctx *DaGDContext) {}

// ExitDaGD is called when production daGD is exited.
func (s *BasePddlListener) ExitDaGD(ctx *DaGDContext) {}

// EnterPrefTimedGD is called when production prefTimedGD is entered.
func (s *BasePddlListener) EnterPrefTimedGD(ctx *PrefTimedGDContext) {}

// ExitPrefTimedGD is called when production prefTimedGD is exited.
func (s *BasePddlListener) ExitPrefTimedGD(ctx *PrefTimedGDContext) {}

// EnterTimedGD is called when production timedGD is entered.
func (s *BasePddlListener) EnterTimedGD(ctx *TimedGDContext) {}

// ExitTimedGD is called when production timedGD is exited.
func (s *BasePddlListener) ExitTimedGD(ctx *TimedGDContext) {}

// EnterTimeSpecifier is called when production timeSpecifier is entered.
func (s *BasePddlListener) EnterTimeSpecifier(ctx *TimeSpecifierContext) {}

// ExitTimeSpecifier is called when production timeSpecifier is exited.
func (s *BasePddlListener) ExitTimeSpecifier(ctx *TimeSpecifierContext) {}

// EnterInterval is called when production interval is entered.
func (s *BasePddlListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BasePddlListener) ExitInterval(ctx *IntervalContext) {}

// EnterDerivedDef is called when production derivedDef is entered.
func (s *BasePddlListener) EnterDerivedDef(ctx *DerivedDefContext) {}

// ExitDerivedDef is called when production derivedDef is exited.
func (s *BasePddlListener) ExitDerivedDef(ctx *DerivedDefContext) {}

// EnterFExp is called when production fExp is entered.
func (s *BasePddlListener) EnterFExp(ctx *FExpContext) {}

// ExitFExp is called when production fExp is exited.
func (s *BasePddlListener) ExitFExp(ctx *FExpContext) {}

// EnterFExp2 is called when production fExp2 is entered.
func (s *BasePddlListener) EnterFExp2(ctx *FExp2Context) {}

// ExitFExp2 is called when production fExp2 is exited.
func (s *BasePddlListener) ExitFExp2(ctx *FExp2Context) {}

// EnterFHead is called when production fHead is entered.
func (s *BasePddlListener) EnterFHead(ctx *FHeadContext) {}

// ExitFHead is called when production fHead is exited.
func (s *BasePddlListener) ExitFHead(ctx *FHeadContext) {}

// EnterEffect is called when production effect is entered.
func (s *BasePddlListener) EnterEffect(ctx *EffectContext) {}

// ExitEffect is called when production effect is exited.
func (s *BasePddlListener) ExitEffect(ctx *EffectContext) {}

// EnterCEffect is called when production cEffect is entered.
func (s *BasePddlListener) EnterCEffect(ctx *CEffectContext) {}

// ExitCEffect is called when production cEffect is exited.
func (s *BasePddlListener) ExitCEffect(ctx *CEffectContext) {}

// EnterPEffect is called when production pEffect is entered.
func (s *BasePddlListener) EnterPEffect(ctx *PEffectContext) {}

// ExitPEffect is called when production pEffect is exited.
func (s *BasePddlListener) ExitPEffect(ctx *PEffectContext) {}

// EnterCondEffect is called when production condEffect is entered.
func (s *BasePddlListener) EnterCondEffect(ctx *CondEffectContext) {}

// ExitCondEffect is called when production condEffect is exited.
func (s *BasePddlListener) ExitCondEffect(ctx *CondEffectContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BasePddlListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BasePddlListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterBinaryComp is called when production binaryComp is entered.
func (s *BasePddlListener) EnterBinaryComp(ctx *BinaryCompContext) {}

// ExitBinaryComp is called when production binaryComp is exited.
func (s *BasePddlListener) ExitBinaryComp(ctx *BinaryCompContext) {}

// EnterAssignOp is called when production assignOp is entered.
func (s *BasePddlListener) EnterAssignOp(ctx *AssignOpContext) {}

// ExitAssignOp is called when production assignOp is exited.
func (s *BasePddlListener) ExitAssignOp(ctx *AssignOpContext) {}

// EnterDurationConstraint is called when production durationConstraint is entered.
func (s *BasePddlListener) EnterDurationConstraint(ctx *DurationConstraintContext) {}

// ExitDurationConstraint is called when production durationConstraint is exited.
func (s *BasePddlListener) ExitDurationConstraint(ctx *DurationConstraintContext) {}

// EnterSimpleDurationConstraint is called when production simpleDurationConstraint is entered.
func (s *BasePddlListener) EnterSimpleDurationConstraint(ctx *SimpleDurationConstraintContext) {}

// ExitSimpleDurationConstraint is called when production simpleDurationConstraint is exited.
func (s *BasePddlListener) ExitSimpleDurationConstraint(ctx *SimpleDurationConstraintContext) {}

// EnterDurOp is called when production durOp is entered.
func (s *BasePddlListener) EnterDurOp(ctx *DurOpContext) {}

// ExitDurOp is called when production durOp is exited.
func (s *BasePddlListener) ExitDurOp(ctx *DurOpContext) {}

// EnterDurValue is called when production durValue is entered.
func (s *BasePddlListener) EnterDurValue(ctx *DurValueContext) {}

// ExitDurValue is called when production durValue is exited.
func (s *BasePddlListener) ExitDurValue(ctx *DurValueContext) {}

// EnterDaEffect is called when production daEffect is entered.
func (s *BasePddlListener) EnterDaEffect(ctx *DaEffectContext) {}

// ExitDaEffect is called when production daEffect is exited.
func (s *BasePddlListener) ExitDaEffect(ctx *DaEffectContext) {}

// EnterTimedEffect is called when production timedEffect is entered.
func (s *BasePddlListener) EnterTimedEffect(ctx *TimedEffectContext) {}

// ExitTimedEffect is called when production timedEffect is exited.
func (s *BasePddlListener) ExitTimedEffect(ctx *TimedEffectContext) {}

// EnterFAssignDA is called when production fAssignDA is entered.
func (s *BasePddlListener) EnterFAssignDA(ctx *FAssignDAContext) {}

// ExitFAssignDA is called when production fAssignDA is exited.
func (s *BasePddlListener) ExitFAssignDA(ctx *FAssignDAContext) {}

// EnterFExpDA is called when production fExpDA is entered.
func (s *BasePddlListener) EnterFExpDA(ctx *FExpDAContext) {}

// ExitFExpDA is called when production fExpDA is exited.
func (s *BasePddlListener) ExitFExpDA(ctx *FExpDAContext) {}

// EnterProblem is called when production problem is entered.
func (s *BasePddlListener) EnterProblem(ctx *ProblemContext) {}

// ExitProblem is called when production problem is exited.
func (s *BasePddlListener) ExitProblem(ctx *ProblemContext) {}

// EnterProblemDecl is called when production problemDecl is entered.
func (s *BasePddlListener) EnterProblemDecl(ctx *ProblemDeclContext) {}

// ExitProblemDecl is called when production problemDecl is exited.
func (s *BasePddlListener) ExitProblemDecl(ctx *ProblemDeclContext) {}

// EnterProblemDomain is called when production problemDomain is entered.
func (s *BasePddlListener) EnterProblemDomain(ctx *ProblemDomainContext) {}

// ExitProblemDomain is called when production problemDomain is exited.
func (s *BasePddlListener) ExitProblemDomain(ctx *ProblemDomainContext) {}

// EnterObjectDecl is called when production objectDecl is entered.
func (s *BasePddlListener) EnterObjectDecl(ctx *ObjectDeclContext) {}

// ExitObjectDecl is called when production objectDecl is exited.
func (s *BasePddlListener) ExitObjectDecl(ctx *ObjectDeclContext) {}

// EnterInit_ is called when production init_ is entered.
func (s *BasePddlListener) EnterInit_(ctx *Init_Context) {}

// ExitInit_ is called when production init_ is exited.
func (s *BasePddlListener) ExitInit_(ctx *Init_Context) {}

// EnterInitEl is called when production initEl is entered.
func (s *BasePddlListener) EnterInitEl(ctx *InitElContext) {}

// ExitInitEl is called when production initEl is exited.
func (s *BasePddlListener) ExitInitEl(ctx *InitElContext) {}

// EnterNameLiteral is called when production nameLiteral is entered.
func (s *BasePddlListener) EnterNameLiteral(ctx *NameLiteralContext) {}

// ExitNameLiteral is called when production nameLiteral is exited.
func (s *BasePddlListener) ExitNameLiteral(ctx *NameLiteralContext) {}

// EnterAtomicNameFormula is called when production atomicNameFormula is entered.
func (s *BasePddlListener) EnterAtomicNameFormula(ctx *AtomicNameFormulaContext) {}

// ExitAtomicNameFormula is called when production atomicNameFormula is exited.
func (s *BasePddlListener) ExitAtomicNameFormula(ctx *AtomicNameFormulaContext) {}

// EnterGoal is called when production goal is entered.
func (s *BasePddlListener) EnterGoal(ctx *GoalContext) {}

// ExitGoal is called when production goal is exited.
func (s *BasePddlListener) ExitGoal(ctx *GoalContext) {}

// EnterProbConstraints is called when production probConstraints is entered.
func (s *BasePddlListener) EnterProbConstraints(ctx *ProbConstraintsContext) {}

// ExitProbConstraints is called when production probConstraints is exited.
func (s *BasePddlListener) ExitProbConstraints(ctx *ProbConstraintsContext) {}

// EnterPrefConGD is called when production prefConGD is entered.
func (s *BasePddlListener) EnterPrefConGD(ctx *PrefConGDContext) {}

// ExitPrefConGD is called when production prefConGD is exited.
func (s *BasePddlListener) ExitPrefConGD(ctx *PrefConGDContext) {}

// EnterMetricSpec is called when production metricSpec is entered.
func (s *BasePddlListener) EnterMetricSpec(ctx *MetricSpecContext) {}

// ExitMetricSpec is called when production metricSpec is exited.
func (s *BasePddlListener) ExitMetricSpec(ctx *MetricSpecContext) {}

// EnterOptimization is called when production optimization is entered.
func (s *BasePddlListener) EnterOptimization(ctx *OptimizationContext) {}

// ExitOptimization is called when production optimization is exited.
func (s *BasePddlListener) ExitOptimization(ctx *OptimizationContext) {}

// EnterMetricFExp is called when production metricFExp is entered.
func (s *BasePddlListener) EnterMetricFExp(ctx *MetricFExpContext) {}

// ExitMetricFExp is called when production metricFExp is exited.
func (s *BasePddlListener) ExitMetricFExp(ctx *MetricFExpContext) {}

// EnterConGD is called when production conGD is entered.
func (s *BasePddlListener) EnterConGD(ctx *ConGDContext) {}

// ExitConGD is called when production conGD is exited.
func (s *BasePddlListener) ExitConGD(ctx *ConGDContext) {}
