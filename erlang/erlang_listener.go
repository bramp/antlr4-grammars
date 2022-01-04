// Code generated from Erlang.g4 by ANTLR 4.9.3. DO NOT EDIT.

package erlang // Erlang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ErlangListener is a complete listener for a parse tree produced by ErlangParser.
type ErlangListener interface {
	antlr.ParseTreeListener

	// EnterForms is called when entering the forms production.
	EnterForms(c *FormsContext)

	// EnterForm is called when entering the form production.
	EnterForm(c *FormContext)

	// EnterTokAtom is called when entering the tokAtom production.
	EnterTokAtom(c *TokAtomContext)

	// EnterTokVar is called when entering the tokVar production.
	EnterTokVar(c *TokVarContext)

	// EnterTokFloat is called when entering the tokFloat production.
	EnterTokFloat(c *TokFloatContext)

	// EnterTokInteger is called when entering the tokInteger production.
	EnterTokInteger(c *TokIntegerContext)

	// EnterTokChar is called when entering the tokChar production.
	EnterTokChar(c *TokCharContext)

	// EnterTokString is called when entering the tokString production.
	EnterTokString(c *TokStringContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterSpecFun is called when entering the specFun production.
	EnterSpecFun(c *SpecFunContext)

	// EnterTypedAttrVal is called when entering the typedAttrVal production.
	EnterTypedAttrVal(c *TypedAttrValContext)

	// EnterTypedRecordFields is called when entering the typedRecordFields production.
	EnterTypedRecordFields(c *TypedRecordFieldsContext)

	// EnterTypedExprs is called when entering the typedExprs production.
	EnterTypedExprs(c *TypedExprsContext)

	// EnterTypedExpr is called when entering the typedExpr production.
	EnterTypedExpr(c *TypedExprContext)

	// EnterTypeSigs is called when entering the typeSigs production.
	EnterTypeSigs(c *TypeSigsContext)

	// EnterTypeSig is called when entering the typeSig production.
	EnterTypeSig(c *TypeSigContext)

	// EnterTypeGuards is called when entering the typeGuards production.
	EnterTypeGuards(c *TypeGuardsContext)

	// EnterTypeGuard is called when entering the typeGuard production.
	EnterTypeGuard(c *TypeGuardContext)

	// EnterTopTypes is called when entering the topTypes production.
	EnterTopTypes(c *TopTypesContext)

	// EnterTopType is called when entering the topType production.
	EnterTopType(c *TopTypeContext)

	// EnterTopType100 is called when entering the topType100 production.
	EnterTopType100(c *TopType100Context)

	// EnterType200 is called when entering the type200 production.
	EnterType200(c *Type200Context)

	// EnterType300 is called when entering the type300 production.
	EnterType300(c *Type300Context)

	// EnterType400 is called when entering the type400 production.
	EnterType400(c *Type400Context)

	// EnterType500 is called when entering the type500 production.
	EnterType500(c *Type500Context)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFunType100 is called when entering the funType100 production.
	EnterFunType100(c *FunType100Context)

	// EnterFunType is called when entering the funType production.
	EnterFunType(c *FunTypeContext)

	// EnterMapPairTypes is called when entering the mapPairTypes production.
	EnterMapPairTypes(c *MapPairTypesContext)

	// EnterMapPairType is called when entering the mapPairType production.
	EnterMapPairType(c *MapPairTypeContext)

	// EnterFieldTypes is called when entering the fieldTypes production.
	EnterFieldTypes(c *FieldTypesContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterBinaryType is called when entering the binaryType production.
	EnterBinaryType(c *BinaryTypeContext)

	// EnterBinBaseType is called when entering the binBaseType production.
	EnterBinBaseType(c *BinBaseTypeContext)

	// EnterBinUnitType is called when entering the binUnitType production.
	EnterBinUnitType(c *BinUnitTypeContext)

	// EnterAttrVal is called when entering the attrVal production.
	EnterAttrVal(c *AttrValContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterFunctionClause is called when entering the functionClause production.
	EnterFunctionClause(c *FunctionClauseContext)

	// EnterClauseArgs is called when entering the clauseArgs production.
	EnterClauseArgs(c *ClauseArgsContext)

	// EnterClauseGuard is called when entering the clauseGuard production.
	EnterClauseGuard(c *ClauseGuardContext)

	// EnterClauseBody is called when entering the clauseBody production.
	EnterClauseBody(c *ClauseBodyContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr100 is called when entering the expr100 production.
	EnterExpr100(c *Expr100Context)

	// EnterExpr150 is called when entering the expr150 production.
	EnterExpr150(c *Expr150Context)

	// EnterExpr160 is called when entering the expr160 production.
	EnterExpr160(c *Expr160Context)

	// EnterExpr200 is called when entering the expr200 production.
	EnterExpr200(c *Expr200Context)

	// EnterExpr300 is called when entering the expr300 production.
	EnterExpr300(c *Expr300Context)

	// EnterExpr400 is called when entering the expr400 production.
	EnterExpr400(c *Expr400Context)

	// EnterExpr500 is called when entering the expr500 production.
	EnterExpr500(c *Expr500Context)

	// EnterExpr600 is called when entering the expr600 production.
	EnterExpr600(c *Expr600Context)

	// EnterExpr650 is called when entering the expr650 production.
	EnterExpr650(c *Expr650Context)

	// EnterExpr700 is called when entering the expr700 production.
	EnterExpr700(c *Expr700Context)

	// EnterExpr800 is called when entering the expr800 production.
	EnterExpr800(c *Expr800Context)

	// EnterExprMax is called when entering the exprMax production.
	EnterExprMax(c *ExprMaxContext)

	// EnterPatExpr is called when entering the patExpr production.
	EnterPatExpr(c *PatExprContext)

	// EnterPatExpr200 is called when entering the patExpr200 production.
	EnterPatExpr200(c *PatExpr200Context)

	// EnterPatExpr300 is called when entering the patExpr300 production.
	EnterPatExpr300(c *PatExpr300Context)

	// EnterPatExpr400 is called when entering the patExpr400 production.
	EnterPatExpr400(c *PatExpr400Context)

	// EnterPatExpr500 is called when entering the patExpr500 production.
	EnterPatExpr500(c *PatExpr500Context)

	// EnterPatExpr600 is called when entering the patExpr600 production.
	EnterPatExpr600(c *PatExpr600Context)

	// EnterPatExpr650 is called when entering the patExpr650 production.
	EnterPatExpr650(c *PatExpr650Context)

	// EnterPatExpr700 is called when entering the patExpr700 production.
	EnterPatExpr700(c *PatExpr700Context)

	// EnterPatExpr800 is called when entering the patExpr800 production.
	EnterPatExpr800(c *PatExpr800Context)

	// EnterPatExprMax is called when entering the patExprMax production.
	EnterPatExprMax(c *PatExprMaxContext)

	// EnterMapPatExpr is called when entering the mapPatExpr production.
	EnterMapPatExpr(c *MapPatExprContext)

	// EnterRecordPatExpr is called when entering the recordPatExpr production.
	EnterRecordPatExpr(c *RecordPatExprContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterTail is called when entering the tail production.
	EnterTail(c *TailContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterBinElements is called when entering the binElements production.
	EnterBinElements(c *BinElementsContext)

	// EnterBinElement is called when entering the binElement production.
	EnterBinElement(c *BinElementContext)

	// EnterBitExpr is called when entering the bitExpr production.
	EnterBitExpr(c *BitExprContext)

	// EnterOptBitSizeExpr is called when entering the optBitSizeExpr production.
	EnterOptBitSizeExpr(c *OptBitSizeExprContext)

	// EnterOptBitTypeList is called when entering the optBitTypeList production.
	EnterOptBitTypeList(c *OptBitTypeListContext)

	// EnterBitTypeList is called when entering the bitTypeList production.
	EnterBitTypeList(c *BitTypeListContext)

	// EnterBitType is called when entering the bitType production.
	EnterBitType(c *BitTypeContext)

	// EnterBitSizeExpr is called when entering the bitSizeExpr production.
	EnterBitSizeExpr(c *BitSizeExprContext)

	// EnterListComprehension is called when entering the listComprehension production.
	EnterListComprehension(c *ListComprehensionContext)

	// EnterBinaryComprehension is called when entering the binaryComprehension production.
	EnterBinaryComprehension(c *BinaryComprehensionContext)

	// EnterLcExprs is called when entering the lcExprs production.
	EnterLcExprs(c *LcExprsContext)

	// EnterLcExpr is called when entering the lcExpr production.
	EnterLcExpr(c *LcExprContext)

	// EnterTuple_ is called when entering the tuple_ production.
	EnterTuple_(c *Tuple_Context)

	// EnterMapExpr is called when entering the mapExpr production.
	EnterMapExpr(c *MapExprContext)

	// EnterMapTuple is called when entering the mapTuple production.
	EnterMapTuple(c *MapTupleContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterMapFieldAssoc is called when entering the mapFieldAssoc production.
	EnterMapFieldAssoc(c *MapFieldAssocContext)

	// EnterMapFieldExact is called when entering the mapFieldExact production.
	EnterMapFieldExact(c *MapFieldExactContext)

	// EnterMapKey is called when entering the mapKey production.
	EnterMapKey(c *MapKeyContext)

	// EnterRecordExpr is called when entering the recordExpr production.
	EnterRecordExpr(c *RecordExprContext)

	// EnterRecordTuple is called when entering the recordTuple production.
	EnterRecordTuple(c *RecordTupleContext)

	// EnterRecordFields is called when entering the recordFields production.
	EnterRecordFields(c *RecordFieldsContext)

	// EnterRecordField is called when entering the recordField production.
	EnterRecordField(c *RecordFieldContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterIfClauses is called when entering the ifClauses production.
	EnterIfClauses(c *IfClausesContext)

	// EnterIfClause is called when entering the ifClause production.
	EnterIfClause(c *IfClauseContext)

	// EnterCaseExpr is called when entering the caseExpr production.
	EnterCaseExpr(c *CaseExprContext)

	// EnterCrClauses is called when entering the crClauses production.
	EnterCrClauses(c *CrClausesContext)

	// EnterCrClause is called when entering the crClause production.
	EnterCrClause(c *CrClauseContext)

	// EnterReceiveExpr is called when entering the receiveExpr production.
	EnterReceiveExpr(c *ReceiveExprContext)

	// EnterFunExpr is called when entering the funExpr production.
	EnterFunExpr(c *FunExprContext)

	// EnterAtomOrVar is called when entering the atomOrVar production.
	EnterAtomOrVar(c *AtomOrVarContext)

	// EnterIntegerOrVar is called when entering the integerOrVar production.
	EnterIntegerOrVar(c *IntegerOrVarContext)

	// EnterFunClauses is called when entering the funClauses production.
	EnterFunClauses(c *FunClausesContext)

	// EnterFunClause is called when entering the funClause production.
	EnterFunClause(c *FunClauseContext)

	// EnterTryExpr is called when entering the tryExpr production.
	EnterTryExpr(c *TryExprContext)

	// EnterTryCatch is called when entering the tryCatch production.
	EnterTryCatch(c *TryCatchContext)

	// EnterTryClauses is called when entering the tryClauses production.
	EnterTryClauses(c *TryClausesContext)

	// EnterTryClause is called when entering the tryClause production.
	EnterTryClause(c *TryClauseContext)

	// EnterTryOptStackTrace is called when entering the tryOptStackTrace production.
	EnterTryOptStackTrace(c *TryOptStackTraceContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterPatArgumentList is called when entering the patArgumentList production.
	EnterPatArgumentList(c *PatArgumentListContext)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// EnterPatExprs is called when entering the patExprs production.
	EnterPatExprs(c *PatExprsContext)

	// EnterGuard_ is called when entering the guard_ production.
	EnterGuard_(c *Guard_Context)

	// EnterAtomic is called when entering the atomic production.
	EnterAtomic(c *AtomicContext)

	// EnterPrefixOp is called when entering the prefixOp production.
	EnterPrefixOp(c *PrefixOpContext)

	// EnterMultOp is called when entering the multOp production.
	EnterMultOp(c *MultOpContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterListOp is called when entering the listOp production.
	EnterListOp(c *ListOpContext)

	// EnterCompOp is called when entering the compOp production.
	EnterCompOp(c *CompOpContext)

	// ExitForms is called when exiting the forms production.
	ExitForms(c *FormsContext)

	// ExitForm is called when exiting the form production.
	ExitForm(c *FormContext)

	// ExitTokAtom is called when exiting the tokAtom production.
	ExitTokAtom(c *TokAtomContext)

	// ExitTokVar is called when exiting the tokVar production.
	ExitTokVar(c *TokVarContext)

	// ExitTokFloat is called when exiting the tokFloat production.
	ExitTokFloat(c *TokFloatContext)

	// ExitTokInteger is called when exiting the tokInteger production.
	ExitTokInteger(c *TokIntegerContext)

	// ExitTokChar is called when exiting the tokChar production.
	ExitTokChar(c *TokCharContext)

	// ExitTokString is called when exiting the tokString production.
	ExitTokString(c *TokStringContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitSpecFun is called when exiting the specFun production.
	ExitSpecFun(c *SpecFunContext)

	// ExitTypedAttrVal is called when exiting the typedAttrVal production.
	ExitTypedAttrVal(c *TypedAttrValContext)

	// ExitTypedRecordFields is called when exiting the typedRecordFields production.
	ExitTypedRecordFields(c *TypedRecordFieldsContext)

	// ExitTypedExprs is called when exiting the typedExprs production.
	ExitTypedExprs(c *TypedExprsContext)

	// ExitTypedExpr is called when exiting the typedExpr production.
	ExitTypedExpr(c *TypedExprContext)

	// ExitTypeSigs is called when exiting the typeSigs production.
	ExitTypeSigs(c *TypeSigsContext)

	// ExitTypeSig is called when exiting the typeSig production.
	ExitTypeSig(c *TypeSigContext)

	// ExitTypeGuards is called when exiting the typeGuards production.
	ExitTypeGuards(c *TypeGuardsContext)

	// ExitTypeGuard is called when exiting the typeGuard production.
	ExitTypeGuard(c *TypeGuardContext)

	// ExitTopTypes is called when exiting the topTypes production.
	ExitTopTypes(c *TopTypesContext)

	// ExitTopType is called when exiting the topType production.
	ExitTopType(c *TopTypeContext)

	// ExitTopType100 is called when exiting the topType100 production.
	ExitTopType100(c *TopType100Context)

	// ExitType200 is called when exiting the type200 production.
	ExitType200(c *Type200Context)

	// ExitType300 is called when exiting the type300 production.
	ExitType300(c *Type300Context)

	// ExitType400 is called when exiting the type400 production.
	ExitType400(c *Type400Context)

	// ExitType500 is called when exiting the type500 production.
	ExitType500(c *Type500Context)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFunType100 is called when exiting the funType100 production.
	ExitFunType100(c *FunType100Context)

	// ExitFunType is called when exiting the funType production.
	ExitFunType(c *FunTypeContext)

	// ExitMapPairTypes is called when exiting the mapPairTypes production.
	ExitMapPairTypes(c *MapPairTypesContext)

	// ExitMapPairType is called when exiting the mapPairType production.
	ExitMapPairType(c *MapPairTypeContext)

	// ExitFieldTypes is called when exiting the fieldTypes production.
	ExitFieldTypes(c *FieldTypesContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitBinaryType is called when exiting the binaryType production.
	ExitBinaryType(c *BinaryTypeContext)

	// ExitBinBaseType is called when exiting the binBaseType production.
	ExitBinBaseType(c *BinBaseTypeContext)

	// ExitBinUnitType is called when exiting the binUnitType production.
	ExitBinUnitType(c *BinUnitTypeContext)

	// ExitAttrVal is called when exiting the attrVal production.
	ExitAttrVal(c *AttrValContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitFunctionClause is called when exiting the functionClause production.
	ExitFunctionClause(c *FunctionClauseContext)

	// ExitClauseArgs is called when exiting the clauseArgs production.
	ExitClauseArgs(c *ClauseArgsContext)

	// ExitClauseGuard is called when exiting the clauseGuard production.
	ExitClauseGuard(c *ClauseGuardContext)

	// ExitClauseBody is called when exiting the clauseBody production.
	ExitClauseBody(c *ClauseBodyContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr100 is called when exiting the expr100 production.
	ExitExpr100(c *Expr100Context)

	// ExitExpr150 is called when exiting the expr150 production.
	ExitExpr150(c *Expr150Context)

	// ExitExpr160 is called when exiting the expr160 production.
	ExitExpr160(c *Expr160Context)

	// ExitExpr200 is called when exiting the expr200 production.
	ExitExpr200(c *Expr200Context)

	// ExitExpr300 is called when exiting the expr300 production.
	ExitExpr300(c *Expr300Context)

	// ExitExpr400 is called when exiting the expr400 production.
	ExitExpr400(c *Expr400Context)

	// ExitExpr500 is called when exiting the expr500 production.
	ExitExpr500(c *Expr500Context)

	// ExitExpr600 is called when exiting the expr600 production.
	ExitExpr600(c *Expr600Context)

	// ExitExpr650 is called when exiting the expr650 production.
	ExitExpr650(c *Expr650Context)

	// ExitExpr700 is called when exiting the expr700 production.
	ExitExpr700(c *Expr700Context)

	// ExitExpr800 is called when exiting the expr800 production.
	ExitExpr800(c *Expr800Context)

	// ExitExprMax is called when exiting the exprMax production.
	ExitExprMax(c *ExprMaxContext)

	// ExitPatExpr is called when exiting the patExpr production.
	ExitPatExpr(c *PatExprContext)

	// ExitPatExpr200 is called when exiting the patExpr200 production.
	ExitPatExpr200(c *PatExpr200Context)

	// ExitPatExpr300 is called when exiting the patExpr300 production.
	ExitPatExpr300(c *PatExpr300Context)

	// ExitPatExpr400 is called when exiting the patExpr400 production.
	ExitPatExpr400(c *PatExpr400Context)

	// ExitPatExpr500 is called when exiting the patExpr500 production.
	ExitPatExpr500(c *PatExpr500Context)

	// ExitPatExpr600 is called when exiting the patExpr600 production.
	ExitPatExpr600(c *PatExpr600Context)

	// ExitPatExpr650 is called when exiting the patExpr650 production.
	ExitPatExpr650(c *PatExpr650Context)

	// ExitPatExpr700 is called when exiting the patExpr700 production.
	ExitPatExpr700(c *PatExpr700Context)

	// ExitPatExpr800 is called when exiting the patExpr800 production.
	ExitPatExpr800(c *PatExpr800Context)

	// ExitPatExprMax is called when exiting the patExprMax production.
	ExitPatExprMax(c *PatExprMaxContext)

	// ExitMapPatExpr is called when exiting the mapPatExpr production.
	ExitMapPatExpr(c *MapPatExprContext)

	// ExitRecordPatExpr is called when exiting the recordPatExpr production.
	ExitRecordPatExpr(c *RecordPatExprContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitTail is called when exiting the tail production.
	ExitTail(c *TailContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitBinElements is called when exiting the binElements production.
	ExitBinElements(c *BinElementsContext)

	// ExitBinElement is called when exiting the binElement production.
	ExitBinElement(c *BinElementContext)

	// ExitBitExpr is called when exiting the bitExpr production.
	ExitBitExpr(c *BitExprContext)

	// ExitOptBitSizeExpr is called when exiting the optBitSizeExpr production.
	ExitOptBitSizeExpr(c *OptBitSizeExprContext)

	// ExitOptBitTypeList is called when exiting the optBitTypeList production.
	ExitOptBitTypeList(c *OptBitTypeListContext)

	// ExitBitTypeList is called when exiting the bitTypeList production.
	ExitBitTypeList(c *BitTypeListContext)

	// ExitBitType is called when exiting the bitType production.
	ExitBitType(c *BitTypeContext)

	// ExitBitSizeExpr is called when exiting the bitSizeExpr production.
	ExitBitSizeExpr(c *BitSizeExprContext)

	// ExitListComprehension is called when exiting the listComprehension production.
	ExitListComprehension(c *ListComprehensionContext)

	// ExitBinaryComprehension is called when exiting the binaryComprehension production.
	ExitBinaryComprehension(c *BinaryComprehensionContext)

	// ExitLcExprs is called when exiting the lcExprs production.
	ExitLcExprs(c *LcExprsContext)

	// ExitLcExpr is called when exiting the lcExpr production.
	ExitLcExpr(c *LcExprContext)

	// ExitTuple_ is called when exiting the tuple_ production.
	ExitTuple_(c *Tuple_Context)

	// ExitMapExpr is called when exiting the mapExpr production.
	ExitMapExpr(c *MapExprContext)

	// ExitMapTuple is called when exiting the mapTuple production.
	ExitMapTuple(c *MapTupleContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitMapFieldAssoc is called when exiting the mapFieldAssoc production.
	ExitMapFieldAssoc(c *MapFieldAssocContext)

	// ExitMapFieldExact is called when exiting the mapFieldExact production.
	ExitMapFieldExact(c *MapFieldExactContext)

	// ExitMapKey is called when exiting the mapKey production.
	ExitMapKey(c *MapKeyContext)

	// ExitRecordExpr is called when exiting the recordExpr production.
	ExitRecordExpr(c *RecordExprContext)

	// ExitRecordTuple is called when exiting the recordTuple production.
	ExitRecordTuple(c *RecordTupleContext)

	// ExitRecordFields is called when exiting the recordFields production.
	ExitRecordFields(c *RecordFieldsContext)

	// ExitRecordField is called when exiting the recordField production.
	ExitRecordField(c *RecordFieldContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitIfClauses is called when exiting the ifClauses production.
	ExitIfClauses(c *IfClausesContext)

	// ExitIfClause is called when exiting the ifClause production.
	ExitIfClause(c *IfClauseContext)

	// ExitCaseExpr is called when exiting the caseExpr production.
	ExitCaseExpr(c *CaseExprContext)

	// ExitCrClauses is called when exiting the crClauses production.
	ExitCrClauses(c *CrClausesContext)

	// ExitCrClause is called when exiting the crClause production.
	ExitCrClause(c *CrClauseContext)

	// ExitReceiveExpr is called when exiting the receiveExpr production.
	ExitReceiveExpr(c *ReceiveExprContext)

	// ExitFunExpr is called when exiting the funExpr production.
	ExitFunExpr(c *FunExprContext)

	// ExitAtomOrVar is called when exiting the atomOrVar production.
	ExitAtomOrVar(c *AtomOrVarContext)

	// ExitIntegerOrVar is called when exiting the integerOrVar production.
	ExitIntegerOrVar(c *IntegerOrVarContext)

	// ExitFunClauses is called when exiting the funClauses production.
	ExitFunClauses(c *FunClausesContext)

	// ExitFunClause is called when exiting the funClause production.
	ExitFunClause(c *FunClauseContext)

	// ExitTryExpr is called when exiting the tryExpr production.
	ExitTryExpr(c *TryExprContext)

	// ExitTryCatch is called when exiting the tryCatch production.
	ExitTryCatch(c *TryCatchContext)

	// ExitTryClauses is called when exiting the tryClauses production.
	ExitTryClauses(c *TryClausesContext)

	// ExitTryClause is called when exiting the tryClause production.
	ExitTryClause(c *TryClauseContext)

	// ExitTryOptStackTrace is called when exiting the tryOptStackTrace production.
	ExitTryOptStackTrace(c *TryOptStackTraceContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitPatArgumentList is called when exiting the patArgumentList production.
	ExitPatArgumentList(c *PatArgumentListContext)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)

	// ExitPatExprs is called when exiting the patExprs production.
	ExitPatExprs(c *PatExprsContext)

	// ExitGuard_ is called when exiting the guard_ production.
	ExitGuard_(c *Guard_Context)

	// ExitAtomic is called when exiting the atomic production.
	ExitAtomic(c *AtomicContext)

	// ExitPrefixOp is called when exiting the prefixOp production.
	ExitPrefixOp(c *PrefixOpContext)

	// ExitMultOp is called when exiting the multOp production.
	ExitMultOp(c *MultOpContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitListOp is called when exiting the listOp production.
	ExitListOp(c *ListOpContext)

	// ExitCompOp is called when exiting the compOp production.
	ExitCompOp(c *CompOpContext)
}
