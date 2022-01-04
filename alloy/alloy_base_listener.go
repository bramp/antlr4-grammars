// Code generated from alloy.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alloy // alloy
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasealloyListener is a complete listener for a parse tree produced by alloyParser.
type BasealloyListener struct{}

var _ alloyListener = &BasealloyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasealloyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasealloyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasealloyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasealloyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAlloyModule is called when production alloyModule is entered.
func (s *BasealloyListener) EnterAlloyModule(ctx *AlloyModuleContext) {}

// ExitAlloyModule is called when production alloyModule is exited.
func (s *BasealloyListener) ExitAlloyModule(ctx *AlloyModuleContext) {}

// EnterModuleDecl is called when production moduleDecl is entered.
func (s *BasealloyListener) EnterModuleDecl(ctx *ModuleDeclContext) {}

// ExitModuleDecl is called when production moduleDecl is exited.
func (s *BasealloyListener) ExitModuleDecl(ctx *ModuleDeclContext) {}

// EnterImport_ is called when production import_ is entered.
func (s *BasealloyListener) EnterImport_(ctx *Import_Context) {}

// ExitImport_ is called when production import_ is exited.
func (s *BasealloyListener) ExitImport_(ctx *Import_Context) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BasealloyListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BasealloyListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterSigDecl is called when production sigDecl is entered.
func (s *BasealloyListener) EnterSigDecl(ctx *SigDeclContext) {}

// ExitSigDecl is called when production sigDecl is exited.
func (s *BasealloyListener) ExitSigDecl(ctx *SigDeclContext) {}

// EnterSigExt is called when production sigExt is entered.
func (s *BasealloyListener) EnterSigExt(ctx *SigExtContext) {}

// ExitSigExt is called when production sigExt is exited.
func (s *BasealloyListener) ExitSigExt(ctx *SigExtContext) {}

// EnterMult is called when production mult is entered.
func (s *BasealloyListener) EnterMult(ctx *MultContext) {}

// ExitMult is called when production mult is exited.
func (s *BasealloyListener) ExitMult(ctx *MultContext) {}

// EnterDecl is called when production decl is entered.
func (s *BasealloyListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BasealloyListener) ExitDecl(ctx *DeclContext) {}

// EnterFactDecl is called when production factDecl is entered.
func (s *BasealloyListener) EnterFactDecl(ctx *FactDeclContext) {}

// ExitFactDecl is called when production factDecl is exited.
func (s *BasealloyListener) ExitFactDecl(ctx *FactDeclContext) {}

// EnterPredDecl is called when production predDecl is entered.
func (s *BasealloyListener) EnterPredDecl(ctx *PredDeclContext) {}

// ExitPredDecl is called when production predDecl is exited.
func (s *BasealloyListener) ExitPredDecl(ctx *PredDeclContext) {}

// EnterFunDecl is called when production funDecl is entered.
func (s *BasealloyListener) EnterFunDecl(ctx *FunDeclContext) {}

// ExitFunDecl is called when production funDecl is exited.
func (s *BasealloyListener) ExitFunDecl(ctx *FunDeclContext) {}

// EnterParaDecls is called when production paraDecls is entered.
func (s *BasealloyListener) EnterParaDecls(ctx *ParaDeclsContext) {}

// ExitParaDecls is called when production paraDecls is exited.
func (s *BasealloyListener) ExitParaDecls(ctx *ParaDeclsContext) {}

// EnterAssertDecl is called when production assertDecl is entered.
func (s *BasealloyListener) EnterAssertDecl(ctx *AssertDeclContext) {}

// ExitAssertDecl is called when production assertDecl is exited.
func (s *BasealloyListener) ExitAssertDecl(ctx *AssertDeclContext) {}

// EnterCmdDecl is called when production cmdDecl is entered.
func (s *BasealloyListener) EnterCmdDecl(ctx *CmdDeclContext) {}

// ExitCmdDecl is called when production cmdDecl is exited.
func (s *BasealloyListener) ExitCmdDecl(ctx *CmdDeclContext) {}

// EnterScope is called when production scope is entered.
func (s *BasealloyListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BasealloyListener) ExitScope(ctx *ScopeContext) {}

// EnterTypescope is called when production typescope is entered.
func (s *BasealloyListener) EnterTypescope(ctx *TypescopeContext) {}

// ExitTypescope is called when production typescope is exited.
func (s *BasealloyListener) ExitTypescope(ctx *TypescopeContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasealloyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasealloyListener) ExitExpr(ctx *ExprContext) {}

// EnterConst_ is called when production const_ is entered.
func (s *BasealloyListener) EnterConst_(ctx *Const_Context) {}

// ExitConst_ is called when production const_ is exited.
func (s *BasealloyListener) ExitConst_(ctx *Const_Context) {}

// EnterUnOp is called when production unOp is entered.
func (s *BasealloyListener) EnterUnOp(ctx *UnOpContext) {}

// ExitUnOp is called when production unOp is exited.
func (s *BasealloyListener) ExitUnOp(ctx *UnOpContext) {}

// EnterBinOp is called when production binOp is entered.
func (s *BasealloyListener) EnterBinOp(ctx *BinOpContext) {}

// ExitBinOp is called when production binOp is exited.
func (s *BasealloyListener) ExitBinOp(ctx *BinOpContext) {}

// EnterArrowOp is called when production arrowOp is entered.
func (s *BasealloyListener) EnterArrowOp(ctx *ArrowOpContext) {}

// ExitArrowOp is called when production arrowOp is exited.
func (s *BasealloyListener) ExitArrowOp(ctx *ArrowOpContext) {}

// EnterCompareOp is called when production compareOp is entered.
func (s *BasealloyListener) EnterCompareOp(ctx *CompareOpContext) {}

// ExitCompareOp is called when production compareOp is exited.
func (s *BasealloyListener) ExitCompareOp(ctx *CompareOpContext) {}

// EnterLetDecl is called when production letDecl is entered.
func (s *BasealloyListener) EnterLetDecl(ctx *LetDeclContext) {}

// ExitLetDecl is called when production letDecl is exited.
func (s *BasealloyListener) ExitLetDecl(ctx *LetDeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BasealloyListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasealloyListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockOrBar is called when production blockOrBar is entered.
func (s *BasealloyListener) EnterBlockOrBar(ctx *BlockOrBarContext) {}

// ExitBlockOrBar is called when production blockOrBar is exited.
func (s *BasealloyListener) ExitBlockOrBar(ctx *BlockOrBarContext) {}

// EnterQuant is called when production quant is entered.
func (s *BasealloyListener) EnterQuant(ctx *QuantContext) {}

// ExitQuant is called when production quant is exited.
func (s *BasealloyListener) ExitQuant(ctx *QuantContext) {}

// EnterQualName is called when production qualName is entered.
func (s *BasealloyListener) EnterQualName(ctx *QualNameContext) {}

// ExitQualName is called when production qualName is exited.
func (s *BasealloyListener) ExitQualName(ctx *QualNameContext) {}

// EnterName is called when production name is entered.
func (s *BasealloyListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasealloyListener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *BasealloyListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasealloyListener) ExitNumber(ctx *NumberContext) {}
