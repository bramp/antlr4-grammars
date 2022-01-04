// Code generated from alloy.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alloy // alloy
import "github.com/antlr/antlr4/runtime/Go/antlr"

// alloyListener is a complete listener for a parse tree produced by alloyParser.
type alloyListener interface {
	antlr.ParseTreeListener

	// EnterAlloyModule is called when entering the alloyModule production.
	EnterAlloyModule(c *AlloyModuleContext)

	// EnterModuleDecl is called when entering the moduleDecl production.
	EnterModuleDecl(c *ModuleDeclContext)

	// EnterImport_ is called when entering the import_ production.
	EnterImport_(c *Import_Context)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterSigDecl is called when entering the sigDecl production.
	EnterSigDecl(c *SigDeclContext)

	// EnterSigExt is called when entering the sigExt production.
	EnterSigExt(c *SigExtContext)

	// EnterMult is called when entering the mult production.
	EnterMult(c *MultContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterFactDecl is called when entering the factDecl production.
	EnterFactDecl(c *FactDeclContext)

	// EnterPredDecl is called when entering the predDecl production.
	EnterPredDecl(c *PredDeclContext)

	// EnterFunDecl is called when entering the funDecl production.
	EnterFunDecl(c *FunDeclContext)

	// EnterParaDecls is called when entering the paraDecls production.
	EnterParaDecls(c *ParaDeclsContext)

	// EnterAssertDecl is called when entering the assertDecl production.
	EnterAssertDecl(c *AssertDeclContext)

	// EnterCmdDecl is called when entering the cmdDecl production.
	EnterCmdDecl(c *CmdDeclContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterTypescope is called when entering the typescope production.
	EnterTypescope(c *TypescopeContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterConst_ is called when entering the const_ production.
	EnterConst_(c *Const_Context)

	// EnterUnOp is called when entering the unOp production.
	EnterUnOp(c *UnOpContext)

	// EnterBinOp is called when entering the binOp production.
	EnterBinOp(c *BinOpContext)

	// EnterArrowOp is called when entering the arrowOp production.
	EnterArrowOp(c *ArrowOpContext)

	// EnterCompareOp is called when entering the compareOp production.
	EnterCompareOp(c *CompareOpContext)

	// EnterLetDecl is called when entering the letDecl production.
	EnterLetDecl(c *LetDeclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockOrBar is called when entering the blockOrBar production.
	EnterBlockOrBar(c *BlockOrBarContext)

	// EnterQuant is called when entering the quant production.
	EnterQuant(c *QuantContext)

	// EnterQualName is called when entering the qualName production.
	EnterQualName(c *QualNameContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitAlloyModule is called when exiting the alloyModule production.
	ExitAlloyModule(c *AlloyModuleContext)

	// ExitModuleDecl is called when exiting the moduleDecl production.
	ExitModuleDecl(c *ModuleDeclContext)

	// ExitImport_ is called when exiting the import_ production.
	ExitImport_(c *Import_Context)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitSigDecl is called when exiting the sigDecl production.
	ExitSigDecl(c *SigDeclContext)

	// ExitSigExt is called when exiting the sigExt production.
	ExitSigExt(c *SigExtContext)

	// ExitMult is called when exiting the mult production.
	ExitMult(c *MultContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitFactDecl is called when exiting the factDecl production.
	ExitFactDecl(c *FactDeclContext)

	// ExitPredDecl is called when exiting the predDecl production.
	ExitPredDecl(c *PredDeclContext)

	// ExitFunDecl is called when exiting the funDecl production.
	ExitFunDecl(c *FunDeclContext)

	// ExitParaDecls is called when exiting the paraDecls production.
	ExitParaDecls(c *ParaDeclsContext)

	// ExitAssertDecl is called when exiting the assertDecl production.
	ExitAssertDecl(c *AssertDeclContext)

	// ExitCmdDecl is called when exiting the cmdDecl production.
	ExitCmdDecl(c *CmdDeclContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitTypescope is called when exiting the typescope production.
	ExitTypescope(c *TypescopeContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitConst_ is called when exiting the const_ production.
	ExitConst_(c *Const_Context)

	// ExitUnOp is called when exiting the unOp production.
	ExitUnOp(c *UnOpContext)

	// ExitBinOp is called when exiting the binOp production.
	ExitBinOp(c *BinOpContext)

	// ExitArrowOp is called when exiting the arrowOp production.
	ExitArrowOp(c *ArrowOpContext)

	// ExitCompareOp is called when exiting the compareOp production.
	ExitCompareOp(c *CompareOpContext)

	// ExitLetDecl is called when exiting the letDecl production.
	ExitLetDecl(c *LetDeclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockOrBar is called when exiting the blockOrBar production.
	ExitBlockOrBar(c *BlockOrBarContext)

	// ExitQuant is called when exiting the quant production.
	ExitQuant(c *QuantContext)

	// ExitQualName is called when exiting the qualName production.
	ExitQualName(c *QualNameContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
