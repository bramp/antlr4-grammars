// Code generated from Erlang.g4 by ANTLR 4.9.3. DO NOT EDIT.

package erlang // Erlang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseErlangListener is a complete listener for a parse tree produced by ErlangParser.
type BaseErlangListener struct{}

var _ ErlangListener = &BaseErlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseErlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseErlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseErlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseErlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterForms is called when production forms is entered.
func (s *BaseErlangListener) EnterForms(ctx *FormsContext) {}

// ExitForms is called when production forms is exited.
func (s *BaseErlangListener) ExitForms(ctx *FormsContext) {}

// EnterForm is called when production form is entered.
func (s *BaseErlangListener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseErlangListener) ExitForm(ctx *FormContext) {}

// EnterTokAtom is called when production tokAtom is entered.
func (s *BaseErlangListener) EnterTokAtom(ctx *TokAtomContext) {}

// ExitTokAtom is called when production tokAtom is exited.
func (s *BaseErlangListener) ExitTokAtom(ctx *TokAtomContext) {}

// EnterTokVar is called when production tokVar is entered.
func (s *BaseErlangListener) EnterTokVar(ctx *TokVarContext) {}

// ExitTokVar is called when production tokVar is exited.
func (s *BaseErlangListener) ExitTokVar(ctx *TokVarContext) {}

// EnterTokFloat is called when production tokFloat is entered.
func (s *BaseErlangListener) EnterTokFloat(ctx *TokFloatContext) {}

// ExitTokFloat is called when production tokFloat is exited.
func (s *BaseErlangListener) ExitTokFloat(ctx *TokFloatContext) {}

// EnterTokInteger is called when production tokInteger is entered.
func (s *BaseErlangListener) EnterTokInteger(ctx *TokIntegerContext) {}

// ExitTokInteger is called when production tokInteger is exited.
func (s *BaseErlangListener) ExitTokInteger(ctx *TokIntegerContext) {}

// EnterTokChar is called when production tokChar is entered.
func (s *BaseErlangListener) EnterTokChar(ctx *TokCharContext) {}

// ExitTokChar is called when production tokChar is exited.
func (s *BaseErlangListener) ExitTokChar(ctx *TokCharContext) {}

// EnterTokString is called when production tokString is entered.
func (s *BaseErlangListener) EnterTokString(ctx *TokStringContext) {}

// ExitTokString is called when production tokString is exited.
func (s *BaseErlangListener) ExitTokString(ctx *TokStringContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseErlangListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseErlangListener) ExitAttribute(ctx *AttributeContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseErlangListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseErlangListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterSpecFun is called when production specFun is entered.
func (s *BaseErlangListener) EnterSpecFun(ctx *SpecFunContext) {}

// ExitSpecFun is called when production specFun is exited.
func (s *BaseErlangListener) ExitSpecFun(ctx *SpecFunContext) {}

// EnterTypedAttrVal is called when production typedAttrVal is entered.
func (s *BaseErlangListener) EnterTypedAttrVal(ctx *TypedAttrValContext) {}

// ExitTypedAttrVal is called when production typedAttrVal is exited.
func (s *BaseErlangListener) ExitTypedAttrVal(ctx *TypedAttrValContext) {}

// EnterTypedRecordFields is called when production typedRecordFields is entered.
func (s *BaseErlangListener) EnterTypedRecordFields(ctx *TypedRecordFieldsContext) {}

// ExitTypedRecordFields is called when production typedRecordFields is exited.
func (s *BaseErlangListener) ExitTypedRecordFields(ctx *TypedRecordFieldsContext) {}

// EnterTypedExprs is called when production typedExprs is entered.
func (s *BaseErlangListener) EnterTypedExprs(ctx *TypedExprsContext) {}

// ExitTypedExprs is called when production typedExprs is exited.
func (s *BaseErlangListener) ExitTypedExprs(ctx *TypedExprsContext) {}

// EnterTypedExpr is called when production typedExpr is entered.
func (s *BaseErlangListener) EnterTypedExpr(ctx *TypedExprContext) {}

// ExitTypedExpr is called when production typedExpr is exited.
func (s *BaseErlangListener) ExitTypedExpr(ctx *TypedExprContext) {}

// EnterTypeSigs is called when production typeSigs is entered.
func (s *BaseErlangListener) EnterTypeSigs(ctx *TypeSigsContext) {}

// ExitTypeSigs is called when production typeSigs is exited.
func (s *BaseErlangListener) ExitTypeSigs(ctx *TypeSigsContext) {}

// EnterTypeSig is called when production typeSig is entered.
func (s *BaseErlangListener) EnterTypeSig(ctx *TypeSigContext) {}

// ExitTypeSig is called when production typeSig is exited.
func (s *BaseErlangListener) ExitTypeSig(ctx *TypeSigContext) {}

// EnterTypeGuards is called when production typeGuards is entered.
func (s *BaseErlangListener) EnterTypeGuards(ctx *TypeGuardsContext) {}

// ExitTypeGuards is called when production typeGuards is exited.
func (s *BaseErlangListener) ExitTypeGuards(ctx *TypeGuardsContext) {}

// EnterTypeGuard is called when production typeGuard is entered.
func (s *BaseErlangListener) EnterTypeGuard(ctx *TypeGuardContext) {}

// ExitTypeGuard is called when production typeGuard is exited.
func (s *BaseErlangListener) ExitTypeGuard(ctx *TypeGuardContext) {}

// EnterTopTypes is called when production topTypes is entered.
func (s *BaseErlangListener) EnterTopTypes(ctx *TopTypesContext) {}

// ExitTopTypes is called when production topTypes is exited.
func (s *BaseErlangListener) ExitTopTypes(ctx *TopTypesContext) {}

// EnterTopType is called when production topType is entered.
func (s *BaseErlangListener) EnterTopType(ctx *TopTypeContext) {}

// ExitTopType is called when production topType is exited.
func (s *BaseErlangListener) ExitTopType(ctx *TopTypeContext) {}

// EnterTopType100 is called when production topType100 is entered.
func (s *BaseErlangListener) EnterTopType100(ctx *TopType100Context) {}

// ExitTopType100 is called when production topType100 is exited.
func (s *BaseErlangListener) ExitTopType100(ctx *TopType100Context) {}

// EnterType200 is called when production type200 is entered.
func (s *BaseErlangListener) EnterType200(ctx *Type200Context) {}

// ExitType200 is called when production type200 is exited.
func (s *BaseErlangListener) ExitType200(ctx *Type200Context) {}

// EnterType300 is called when production type300 is entered.
func (s *BaseErlangListener) EnterType300(ctx *Type300Context) {}

// ExitType300 is called when production type300 is exited.
func (s *BaseErlangListener) ExitType300(ctx *Type300Context) {}

// EnterType400 is called when production type400 is entered.
func (s *BaseErlangListener) EnterType400(ctx *Type400Context) {}

// ExitType400 is called when production type400 is exited.
func (s *BaseErlangListener) ExitType400(ctx *Type400Context) {}

// EnterType500 is called when production type500 is entered.
func (s *BaseErlangListener) EnterType500(ctx *Type500Context) {}

// ExitType500 is called when production type500 is exited.
func (s *BaseErlangListener) ExitType500(ctx *Type500Context) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseErlangListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseErlangListener) ExitType_(ctx *Type_Context) {}

// EnterFunType100 is called when production funType100 is entered.
func (s *BaseErlangListener) EnterFunType100(ctx *FunType100Context) {}

// ExitFunType100 is called when production funType100 is exited.
func (s *BaseErlangListener) ExitFunType100(ctx *FunType100Context) {}

// EnterFunType is called when production funType is entered.
func (s *BaseErlangListener) EnterFunType(ctx *FunTypeContext) {}

// ExitFunType is called when production funType is exited.
func (s *BaseErlangListener) ExitFunType(ctx *FunTypeContext) {}

// EnterMapPairTypes is called when production mapPairTypes is entered.
func (s *BaseErlangListener) EnterMapPairTypes(ctx *MapPairTypesContext) {}

// ExitMapPairTypes is called when production mapPairTypes is exited.
func (s *BaseErlangListener) ExitMapPairTypes(ctx *MapPairTypesContext) {}

// EnterMapPairType is called when production mapPairType is entered.
func (s *BaseErlangListener) EnterMapPairType(ctx *MapPairTypeContext) {}

// ExitMapPairType is called when production mapPairType is exited.
func (s *BaseErlangListener) ExitMapPairType(ctx *MapPairTypeContext) {}

// EnterFieldTypes is called when production fieldTypes is entered.
func (s *BaseErlangListener) EnterFieldTypes(ctx *FieldTypesContext) {}

// ExitFieldTypes is called when production fieldTypes is exited.
func (s *BaseErlangListener) ExitFieldTypes(ctx *FieldTypesContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseErlangListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseErlangListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterBinaryType is called when production binaryType is entered.
func (s *BaseErlangListener) EnterBinaryType(ctx *BinaryTypeContext) {}

// ExitBinaryType is called when production binaryType is exited.
func (s *BaseErlangListener) ExitBinaryType(ctx *BinaryTypeContext) {}

// EnterBinBaseType is called when production binBaseType is entered.
func (s *BaseErlangListener) EnterBinBaseType(ctx *BinBaseTypeContext) {}

// ExitBinBaseType is called when production binBaseType is exited.
func (s *BaseErlangListener) ExitBinBaseType(ctx *BinBaseTypeContext) {}

// EnterBinUnitType is called when production binUnitType is entered.
func (s *BaseErlangListener) EnterBinUnitType(ctx *BinUnitTypeContext) {}

// ExitBinUnitType is called when production binUnitType is exited.
func (s *BaseErlangListener) ExitBinUnitType(ctx *BinUnitTypeContext) {}

// EnterAttrVal is called when production attrVal is entered.
func (s *BaseErlangListener) EnterAttrVal(ctx *AttrValContext) {}

// ExitAttrVal is called when production attrVal is exited.
func (s *BaseErlangListener) ExitAttrVal(ctx *AttrValContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaseErlangListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaseErlangListener) ExitFunction_(ctx *Function_Context) {}

// EnterFunctionClause is called when production functionClause is entered.
func (s *BaseErlangListener) EnterFunctionClause(ctx *FunctionClauseContext) {}

// ExitFunctionClause is called when production functionClause is exited.
func (s *BaseErlangListener) ExitFunctionClause(ctx *FunctionClauseContext) {}

// EnterClauseArgs is called when production clauseArgs is entered.
func (s *BaseErlangListener) EnterClauseArgs(ctx *ClauseArgsContext) {}

// ExitClauseArgs is called when production clauseArgs is exited.
func (s *BaseErlangListener) ExitClauseArgs(ctx *ClauseArgsContext) {}

// EnterClauseGuard is called when production clauseGuard is entered.
func (s *BaseErlangListener) EnterClauseGuard(ctx *ClauseGuardContext) {}

// ExitClauseGuard is called when production clauseGuard is exited.
func (s *BaseErlangListener) ExitClauseGuard(ctx *ClauseGuardContext) {}

// EnterClauseBody is called when production clauseBody is entered.
func (s *BaseErlangListener) EnterClauseBody(ctx *ClauseBodyContext) {}

// ExitClauseBody is called when production clauseBody is exited.
func (s *BaseErlangListener) ExitClauseBody(ctx *ClauseBodyContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseErlangListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseErlangListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr100 is called when production expr100 is entered.
func (s *BaseErlangListener) EnterExpr100(ctx *Expr100Context) {}

// ExitExpr100 is called when production expr100 is exited.
func (s *BaseErlangListener) ExitExpr100(ctx *Expr100Context) {}

// EnterExpr150 is called when production expr150 is entered.
func (s *BaseErlangListener) EnterExpr150(ctx *Expr150Context) {}

// ExitExpr150 is called when production expr150 is exited.
func (s *BaseErlangListener) ExitExpr150(ctx *Expr150Context) {}

// EnterExpr160 is called when production expr160 is entered.
func (s *BaseErlangListener) EnterExpr160(ctx *Expr160Context) {}

// ExitExpr160 is called when production expr160 is exited.
func (s *BaseErlangListener) ExitExpr160(ctx *Expr160Context) {}

// EnterExpr200 is called when production expr200 is entered.
func (s *BaseErlangListener) EnterExpr200(ctx *Expr200Context) {}

// ExitExpr200 is called when production expr200 is exited.
func (s *BaseErlangListener) ExitExpr200(ctx *Expr200Context) {}

// EnterExpr300 is called when production expr300 is entered.
func (s *BaseErlangListener) EnterExpr300(ctx *Expr300Context) {}

// ExitExpr300 is called when production expr300 is exited.
func (s *BaseErlangListener) ExitExpr300(ctx *Expr300Context) {}

// EnterExpr400 is called when production expr400 is entered.
func (s *BaseErlangListener) EnterExpr400(ctx *Expr400Context) {}

// ExitExpr400 is called when production expr400 is exited.
func (s *BaseErlangListener) ExitExpr400(ctx *Expr400Context) {}

// EnterExpr500 is called when production expr500 is entered.
func (s *BaseErlangListener) EnterExpr500(ctx *Expr500Context) {}

// ExitExpr500 is called when production expr500 is exited.
func (s *BaseErlangListener) ExitExpr500(ctx *Expr500Context) {}

// EnterExpr600 is called when production expr600 is entered.
func (s *BaseErlangListener) EnterExpr600(ctx *Expr600Context) {}

// ExitExpr600 is called when production expr600 is exited.
func (s *BaseErlangListener) ExitExpr600(ctx *Expr600Context) {}

// EnterExpr650 is called when production expr650 is entered.
func (s *BaseErlangListener) EnterExpr650(ctx *Expr650Context) {}

// ExitExpr650 is called when production expr650 is exited.
func (s *BaseErlangListener) ExitExpr650(ctx *Expr650Context) {}

// EnterExpr700 is called when production expr700 is entered.
func (s *BaseErlangListener) EnterExpr700(ctx *Expr700Context) {}

// ExitExpr700 is called when production expr700 is exited.
func (s *BaseErlangListener) ExitExpr700(ctx *Expr700Context) {}

// EnterExpr800 is called when production expr800 is entered.
func (s *BaseErlangListener) EnterExpr800(ctx *Expr800Context) {}

// ExitExpr800 is called when production expr800 is exited.
func (s *BaseErlangListener) ExitExpr800(ctx *Expr800Context) {}

// EnterExprMax is called when production exprMax is entered.
func (s *BaseErlangListener) EnterExprMax(ctx *ExprMaxContext) {}

// ExitExprMax is called when production exprMax is exited.
func (s *BaseErlangListener) ExitExprMax(ctx *ExprMaxContext) {}

// EnterPatExpr is called when production patExpr is entered.
func (s *BaseErlangListener) EnterPatExpr(ctx *PatExprContext) {}

// ExitPatExpr is called when production patExpr is exited.
func (s *BaseErlangListener) ExitPatExpr(ctx *PatExprContext) {}

// EnterPatExpr200 is called when production patExpr200 is entered.
func (s *BaseErlangListener) EnterPatExpr200(ctx *PatExpr200Context) {}

// ExitPatExpr200 is called when production patExpr200 is exited.
func (s *BaseErlangListener) ExitPatExpr200(ctx *PatExpr200Context) {}

// EnterPatExpr300 is called when production patExpr300 is entered.
func (s *BaseErlangListener) EnterPatExpr300(ctx *PatExpr300Context) {}

// ExitPatExpr300 is called when production patExpr300 is exited.
func (s *BaseErlangListener) ExitPatExpr300(ctx *PatExpr300Context) {}

// EnterPatExpr400 is called when production patExpr400 is entered.
func (s *BaseErlangListener) EnterPatExpr400(ctx *PatExpr400Context) {}

// ExitPatExpr400 is called when production patExpr400 is exited.
func (s *BaseErlangListener) ExitPatExpr400(ctx *PatExpr400Context) {}

// EnterPatExpr500 is called when production patExpr500 is entered.
func (s *BaseErlangListener) EnterPatExpr500(ctx *PatExpr500Context) {}

// ExitPatExpr500 is called when production patExpr500 is exited.
func (s *BaseErlangListener) ExitPatExpr500(ctx *PatExpr500Context) {}

// EnterPatExpr600 is called when production patExpr600 is entered.
func (s *BaseErlangListener) EnterPatExpr600(ctx *PatExpr600Context) {}

// ExitPatExpr600 is called when production patExpr600 is exited.
func (s *BaseErlangListener) ExitPatExpr600(ctx *PatExpr600Context) {}

// EnterPatExpr650 is called when production patExpr650 is entered.
func (s *BaseErlangListener) EnterPatExpr650(ctx *PatExpr650Context) {}

// ExitPatExpr650 is called when production patExpr650 is exited.
func (s *BaseErlangListener) ExitPatExpr650(ctx *PatExpr650Context) {}

// EnterPatExpr700 is called when production patExpr700 is entered.
func (s *BaseErlangListener) EnterPatExpr700(ctx *PatExpr700Context) {}

// ExitPatExpr700 is called when production patExpr700 is exited.
func (s *BaseErlangListener) ExitPatExpr700(ctx *PatExpr700Context) {}

// EnterPatExpr800 is called when production patExpr800 is entered.
func (s *BaseErlangListener) EnterPatExpr800(ctx *PatExpr800Context) {}

// ExitPatExpr800 is called when production patExpr800 is exited.
func (s *BaseErlangListener) ExitPatExpr800(ctx *PatExpr800Context) {}

// EnterPatExprMax is called when production patExprMax is entered.
func (s *BaseErlangListener) EnterPatExprMax(ctx *PatExprMaxContext) {}

// ExitPatExprMax is called when production patExprMax is exited.
func (s *BaseErlangListener) ExitPatExprMax(ctx *PatExprMaxContext) {}

// EnterMapPatExpr is called when production mapPatExpr is entered.
func (s *BaseErlangListener) EnterMapPatExpr(ctx *MapPatExprContext) {}

// ExitMapPatExpr is called when production mapPatExpr is exited.
func (s *BaseErlangListener) ExitMapPatExpr(ctx *MapPatExprContext) {}

// EnterRecordPatExpr is called when production recordPatExpr is entered.
func (s *BaseErlangListener) EnterRecordPatExpr(ctx *RecordPatExprContext) {}

// ExitRecordPatExpr is called when production recordPatExpr is exited.
func (s *BaseErlangListener) ExitRecordPatExpr(ctx *RecordPatExprContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseErlangListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseErlangListener) ExitList_(ctx *List_Context) {}

// EnterTail is called when production tail is entered.
func (s *BaseErlangListener) EnterTail(ctx *TailContext) {}

// ExitTail is called when production tail is exited.
func (s *BaseErlangListener) ExitTail(ctx *TailContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseErlangListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseErlangListener) ExitBinary(ctx *BinaryContext) {}

// EnterBinElements is called when production binElements is entered.
func (s *BaseErlangListener) EnterBinElements(ctx *BinElementsContext) {}

// ExitBinElements is called when production binElements is exited.
func (s *BaseErlangListener) ExitBinElements(ctx *BinElementsContext) {}

// EnterBinElement is called when production binElement is entered.
func (s *BaseErlangListener) EnterBinElement(ctx *BinElementContext) {}

// ExitBinElement is called when production binElement is exited.
func (s *BaseErlangListener) ExitBinElement(ctx *BinElementContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseErlangListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseErlangListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterOptBitSizeExpr is called when production optBitSizeExpr is entered.
func (s *BaseErlangListener) EnterOptBitSizeExpr(ctx *OptBitSizeExprContext) {}

// ExitOptBitSizeExpr is called when production optBitSizeExpr is exited.
func (s *BaseErlangListener) ExitOptBitSizeExpr(ctx *OptBitSizeExprContext) {}

// EnterOptBitTypeList is called when production optBitTypeList is entered.
func (s *BaseErlangListener) EnterOptBitTypeList(ctx *OptBitTypeListContext) {}

// ExitOptBitTypeList is called when production optBitTypeList is exited.
func (s *BaseErlangListener) ExitOptBitTypeList(ctx *OptBitTypeListContext) {}

// EnterBitTypeList is called when production bitTypeList is entered.
func (s *BaseErlangListener) EnterBitTypeList(ctx *BitTypeListContext) {}

// ExitBitTypeList is called when production bitTypeList is exited.
func (s *BaseErlangListener) ExitBitTypeList(ctx *BitTypeListContext) {}

// EnterBitType is called when production bitType is entered.
func (s *BaseErlangListener) EnterBitType(ctx *BitTypeContext) {}

// ExitBitType is called when production bitType is exited.
func (s *BaseErlangListener) ExitBitType(ctx *BitTypeContext) {}

// EnterBitSizeExpr is called when production bitSizeExpr is entered.
func (s *BaseErlangListener) EnterBitSizeExpr(ctx *BitSizeExprContext) {}

// ExitBitSizeExpr is called when production bitSizeExpr is exited.
func (s *BaseErlangListener) ExitBitSizeExpr(ctx *BitSizeExprContext) {}

// EnterListComprehension is called when production listComprehension is entered.
func (s *BaseErlangListener) EnterListComprehension(ctx *ListComprehensionContext) {}

// ExitListComprehension is called when production listComprehension is exited.
func (s *BaseErlangListener) ExitListComprehension(ctx *ListComprehensionContext) {}

// EnterBinaryComprehension is called when production binaryComprehension is entered.
func (s *BaseErlangListener) EnterBinaryComprehension(ctx *BinaryComprehensionContext) {}

// ExitBinaryComprehension is called when production binaryComprehension is exited.
func (s *BaseErlangListener) ExitBinaryComprehension(ctx *BinaryComprehensionContext) {}

// EnterLcExprs is called when production lcExprs is entered.
func (s *BaseErlangListener) EnterLcExprs(ctx *LcExprsContext) {}

// ExitLcExprs is called when production lcExprs is exited.
func (s *BaseErlangListener) ExitLcExprs(ctx *LcExprsContext) {}

// EnterLcExpr is called when production lcExpr is entered.
func (s *BaseErlangListener) EnterLcExpr(ctx *LcExprContext) {}

// ExitLcExpr is called when production lcExpr is exited.
func (s *BaseErlangListener) ExitLcExpr(ctx *LcExprContext) {}

// EnterTuple_ is called when production tuple_ is entered.
func (s *BaseErlangListener) EnterTuple_(ctx *Tuple_Context) {}

// ExitTuple_ is called when production tuple_ is exited.
func (s *BaseErlangListener) ExitTuple_(ctx *Tuple_Context) {}

// EnterMapExpr is called when production mapExpr is entered.
func (s *BaseErlangListener) EnterMapExpr(ctx *MapExprContext) {}

// ExitMapExpr is called when production mapExpr is exited.
func (s *BaseErlangListener) ExitMapExpr(ctx *MapExprContext) {}

// EnterMapTuple is called when production mapTuple is entered.
func (s *BaseErlangListener) EnterMapTuple(ctx *MapTupleContext) {}

// ExitMapTuple is called when production mapTuple is exited.
func (s *BaseErlangListener) ExitMapTuple(ctx *MapTupleContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseErlangListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseErlangListener) ExitMapField(ctx *MapFieldContext) {}

// EnterMapFieldAssoc is called when production mapFieldAssoc is entered.
func (s *BaseErlangListener) EnterMapFieldAssoc(ctx *MapFieldAssocContext) {}

// ExitMapFieldAssoc is called when production mapFieldAssoc is exited.
func (s *BaseErlangListener) ExitMapFieldAssoc(ctx *MapFieldAssocContext) {}

// EnterMapFieldExact is called when production mapFieldExact is entered.
func (s *BaseErlangListener) EnterMapFieldExact(ctx *MapFieldExactContext) {}

// ExitMapFieldExact is called when production mapFieldExact is exited.
func (s *BaseErlangListener) ExitMapFieldExact(ctx *MapFieldExactContext) {}

// EnterMapKey is called when production mapKey is entered.
func (s *BaseErlangListener) EnterMapKey(ctx *MapKeyContext) {}

// ExitMapKey is called when production mapKey is exited.
func (s *BaseErlangListener) ExitMapKey(ctx *MapKeyContext) {}

// EnterRecordExpr is called when production recordExpr is entered.
func (s *BaseErlangListener) EnterRecordExpr(ctx *RecordExprContext) {}

// ExitRecordExpr is called when production recordExpr is exited.
func (s *BaseErlangListener) ExitRecordExpr(ctx *RecordExprContext) {}

// EnterRecordTuple is called when production recordTuple is entered.
func (s *BaseErlangListener) EnterRecordTuple(ctx *RecordTupleContext) {}

// ExitRecordTuple is called when production recordTuple is exited.
func (s *BaseErlangListener) ExitRecordTuple(ctx *RecordTupleContext) {}

// EnterRecordFields is called when production recordFields is entered.
func (s *BaseErlangListener) EnterRecordFields(ctx *RecordFieldsContext) {}

// ExitRecordFields is called when production recordFields is exited.
func (s *BaseErlangListener) ExitRecordFields(ctx *RecordFieldsContext) {}

// EnterRecordField is called when production recordField is entered.
func (s *BaseErlangListener) EnterRecordField(ctx *RecordFieldContext) {}

// ExitRecordField is called when production recordField is exited.
func (s *BaseErlangListener) ExitRecordField(ctx *RecordFieldContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseErlangListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseErlangListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseErlangListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseErlangListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterIfClauses is called when production ifClauses is entered.
func (s *BaseErlangListener) EnterIfClauses(ctx *IfClausesContext) {}

// ExitIfClauses is called when production ifClauses is exited.
func (s *BaseErlangListener) ExitIfClauses(ctx *IfClausesContext) {}

// EnterIfClause is called when production ifClause is entered.
func (s *BaseErlangListener) EnterIfClause(ctx *IfClauseContext) {}

// ExitIfClause is called when production ifClause is exited.
func (s *BaseErlangListener) ExitIfClause(ctx *IfClauseContext) {}

// EnterCaseExpr is called when production caseExpr is entered.
func (s *BaseErlangListener) EnterCaseExpr(ctx *CaseExprContext) {}

// ExitCaseExpr is called when production caseExpr is exited.
func (s *BaseErlangListener) ExitCaseExpr(ctx *CaseExprContext) {}

// EnterCrClauses is called when production crClauses is entered.
func (s *BaseErlangListener) EnterCrClauses(ctx *CrClausesContext) {}

// ExitCrClauses is called when production crClauses is exited.
func (s *BaseErlangListener) ExitCrClauses(ctx *CrClausesContext) {}

// EnterCrClause is called when production crClause is entered.
func (s *BaseErlangListener) EnterCrClause(ctx *CrClauseContext) {}

// ExitCrClause is called when production crClause is exited.
func (s *BaseErlangListener) ExitCrClause(ctx *CrClauseContext) {}

// EnterReceiveExpr is called when production receiveExpr is entered.
func (s *BaseErlangListener) EnterReceiveExpr(ctx *ReceiveExprContext) {}

// ExitReceiveExpr is called when production receiveExpr is exited.
func (s *BaseErlangListener) ExitReceiveExpr(ctx *ReceiveExprContext) {}

// EnterFunExpr is called when production funExpr is entered.
func (s *BaseErlangListener) EnterFunExpr(ctx *FunExprContext) {}

// ExitFunExpr is called when production funExpr is exited.
func (s *BaseErlangListener) ExitFunExpr(ctx *FunExprContext) {}

// EnterAtomOrVar is called when production atomOrVar is entered.
func (s *BaseErlangListener) EnterAtomOrVar(ctx *AtomOrVarContext) {}

// ExitAtomOrVar is called when production atomOrVar is exited.
func (s *BaseErlangListener) ExitAtomOrVar(ctx *AtomOrVarContext) {}

// EnterIntegerOrVar is called when production integerOrVar is entered.
func (s *BaseErlangListener) EnterIntegerOrVar(ctx *IntegerOrVarContext) {}

// ExitIntegerOrVar is called when production integerOrVar is exited.
func (s *BaseErlangListener) ExitIntegerOrVar(ctx *IntegerOrVarContext) {}

// EnterFunClauses is called when production funClauses is entered.
func (s *BaseErlangListener) EnterFunClauses(ctx *FunClausesContext) {}

// ExitFunClauses is called when production funClauses is exited.
func (s *BaseErlangListener) ExitFunClauses(ctx *FunClausesContext) {}

// EnterFunClause is called when production funClause is entered.
func (s *BaseErlangListener) EnterFunClause(ctx *FunClauseContext) {}

// ExitFunClause is called when production funClause is exited.
func (s *BaseErlangListener) ExitFunClause(ctx *FunClauseContext) {}

// EnterTryExpr is called when production tryExpr is entered.
func (s *BaseErlangListener) EnterTryExpr(ctx *TryExprContext) {}

// ExitTryExpr is called when production tryExpr is exited.
func (s *BaseErlangListener) ExitTryExpr(ctx *TryExprContext) {}

// EnterTryCatch is called when production tryCatch is entered.
func (s *BaseErlangListener) EnterTryCatch(ctx *TryCatchContext) {}

// ExitTryCatch is called when production tryCatch is exited.
func (s *BaseErlangListener) ExitTryCatch(ctx *TryCatchContext) {}

// EnterTryClauses is called when production tryClauses is entered.
func (s *BaseErlangListener) EnterTryClauses(ctx *TryClausesContext) {}

// ExitTryClauses is called when production tryClauses is exited.
func (s *BaseErlangListener) ExitTryClauses(ctx *TryClausesContext) {}

// EnterTryClause is called when production tryClause is entered.
func (s *BaseErlangListener) EnterTryClause(ctx *TryClauseContext) {}

// ExitTryClause is called when production tryClause is exited.
func (s *BaseErlangListener) ExitTryClause(ctx *TryClauseContext) {}

// EnterTryOptStackTrace is called when production tryOptStackTrace is entered.
func (s *BaseErlangListener) EnterTryOptStackTrace(ctx *TryOptStackTraceContext) {}

// ExitTryOptStackTrace is called when production tryOptStackTrace is exited.
func (s *BaseErlangListener) ExitTryOptStackTrace(ctx *TryOptStackTraceContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseErlangListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseErlangListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterPatArgumentList is called when production patArgumentList is entered.
func (s *BaseErlangListener) EnterPatArgumentList(ctx *PatArgumentListContext) {}

// ExitPatArgumentList is called when production patArgumentList is exited.
func (s *BaseErlangListener) ExitPatArgumentList(ctx *PatArgumentListContext) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseErlangListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseErlangListener) ExitExprs(ctx *ExprsContext) {}

// EnterPatExprs is called when production patExprs is entered.
func (s *BaseErlangListener) EnterPatExprs(ctx *PatExprsContext) {}

// ExitPatExprs is called when production patExprs is exited.
func (s *BaseErlangListener) ExitPatExprs(ctx *PatExprsContext) {}

// EnterGuard_ is called when production guard_ is entered.
func (s *BaseErlangListener) EnterGuard_(ctx *Guard_Context) {}

// ExitGuard_ is called when production guard_ is exited.
func (s *BaseErlangListener) ExitGuard_(ctx *Guard_Context) {}

// EnterAtomic is called when production atomic is entered.
func (s *BaseErlangListener) EnterAtomic(ctx *AtomicContext) {}

// ExitAtomic is called when production atomic is exited.
func (s *BaseErlangListener) ExitAtomic(ctx *AtomicContext) {}

// EnterPrefixOp is called when production prefixOp is entered.
func (s *BaseErlangListener) EnterPrefixOp(ctx *PrefixOpContext) {}

// ExitPrefixOp is called when production prefixOp is exited.
func (s *BaseErlangListener) ExitPrefixOp(ctx *PrefixOpContext) {}

// EnterMultOp is called when production multOp is entered.
func (s *BaseErlangListener) EnterMultOp(ctx *MultOpContext) {}

// ExitMultOp is called when production multOp is exited.
func (s *BaseErlangListener) ExitMultOp(ctx *MultOpContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BaseErlangListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BaseErlangListener) ExitAddOp(ctx *AddOpContext) {}

// EnterListOp is called when production listOp is entered.
func (s *BaseErlangListener) EnterListOp(ctx *ListOpContext) {}

// ExitListOp is called when production listOp is exited.
func (s *BaseErlangListener) ExitListOp(ctx *ListOpContext) {}

// EnterCompOp is called when production compOp is entered.
func (s *BaseErlangListener) EnterCompOp(ctx *CompOpContext) {}

// ExitCompOp is called when production compOp is exited.
func (s *BaseErlangListener) ExitCompOp(ctx *CompOpContext) {}
