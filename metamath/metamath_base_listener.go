// Code generated from metamath.g4 by ANTLR 4.9.3. DO NOT EDIT.

package metamath // metamath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemetamathListener is a complete listener for a parse tree produced by metamathParser.
type BasemetamathListener struct{}

var _ metamathListener = &BasemetamathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemetamathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemetamathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemetamathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemetamathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDatabase is called when production database is entered.
func (s *BasemetamathListener) EnterDatabase(ctx *DatabaseContext) {}

// ExitDatabase is called when production database is exited.
func (s *BasemetamathListener) ExitDatabase(ctx *DatabaseContext) {}

// EnterOutermostscopestmt is called when production outermostscopestmt is entered.
func (s *BasemetamathListener) EnterOutermostscopestmt(ctx *OutermostscopestmtContext) {}

// ExitOutermostscopestmt is called when production outermostscopestmt is exited.
func (s *BasemetamathListener) ExitOutermostscopestmt(ctx *OutermostscopestmtContext) {}

// EnterIncludestmt is called when production includestmt is entered.
func (s *BasemetamathListener) EnterIncludestmt(ctx *IncludestmtContext) {}

// ExitIncludestmt is called when production includestmt is exited.
func (s *BasemetamathListener) ExitIncludestmt(ctx *IncludestmtContext) {}

// EnterConstantstmt is called when production constantstmt is entered.
func (s *BasemetamathListener) EnterConstantstmt(ctx *ConstantstmtContext) {}

// ExitConstantstmt is called when production constantstmt is exited.
func (s *BasemetamathListener) ExitConstantstmt(ctx *ConstantstmtContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasemetamathListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasemetamathListener) ExitStmt(ctx *StmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BasemetamathListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasemetamathListener) ExitBlock(ctx *BlockContext) {}

// EnterVariablestmt is called when production variablestmt is entered.
func (s *BasemetamathListener) EnterVariablestmt(ctx *VariablestmtContext) {}

// ExitVariablestmt is called when production variablestmt is exited.
func (s *BasemetamathListener) ExitVariablestmt(ctx *VariablestmtContext) {}

// EnterDisjointstmt is called when production disjointstmt is entered.
func (s *BasemetamathListener) EnterDisjointstmt(ctx *DisjointstmtContext) {}

// ExitDisjointstmt is called when production disjointstmt is exited.
func (s *BasemetamathListener) ExitDisjointstmt(ctx *DisjointstmtContext) {}

// EnterHypothesisstmt is called when production hypothesisstmt is entered.
func (s *BasemetamathListener) EnterHypothesisstmt(ctx *HypothesisstmtContext) {}

// ExitHypothesisstmt is called when production hypothesisstmt is exited.
func (s *BasemetamathListener) ExitHypothesisstmt(ctx *HypothesisstmtContext) {}

// EnterFloatingstmt is called when production floatingstmt is entered.
func (s *BasemetamathListener) EnterFloatingstmt(ctx *FloatingstmtContext) {}

// ExitFloatingstmt is called when production floatingstmt is exited.
func (s *BasemetamathListener) ExitFloatingstmt(ctx *FloatingstmtContext) {}

// EnterEssentialstmt is called when production essentialstmt is entered.
func (s *BasemetamathListener) EnterEssentialstmt(ctx *EssentialstmtContext) {}

// ExitEssentialstmt is called when production essentialstmt is exited.
func (s *BasemetamathListener) ExitEssentialstmt(ctx *EssentialstmtContext) {}

// EnterAssertstmt is called when production assertstmt is entered.
func (s *BasemetamathListener) EnterAssertstmt(ctx *AssertstmtContext) {}

// ExitAssertstmt is called when production assertstmt is exited.
func (s *BasemetamathListener) ExitAssertstmt(ctx *AssertstmtContext) {}

// EnterAxiomstmt is called when production axiomstmt is entered.
func (s *BasemetamathListener) EnterAxiomstmt(ctx *AxiomstmtContext) {}

// ExitAxiomstmt is called when production axiomstmt is exited.
func (s *BasemetamathListener) ExitAxiomstmt(ctx *AxiomstmtContext) {}

// EnterProvablestmt is called when production provablestmt is entered.
func (s *BasemetamathListener) EnterProvablestmt(ctx *ProvablestmtContext) {}

// ExitProvablestmt is called when production provablestmt is exited.
func (s *BasemetamathListener) ExitProvablestmt(ctx *ProvablestmtContext) {}

// EnterProof is called when production proof is entered.
func (s *BasemetamathListener) EnterProof(ctx *ProofContext) {}

// ExitProof is called when production proof is exited.
func (s *BasemetamathListener) ExitProof(ctx *ProofContext) {}

// EnterUncompressedproof is called when production uncompressedproof is entered.
func (s *BasemetamathListener) EnterUncompressedproof(ctx *UncompressedproofContext) {}

// ExitUncompressedproof is called when production uncompressedproof is exited.
func (s *BasemetamathListener) ExitUncompressedproof(ctx *UncompressedproofContext) {}

// EnterCompressedproof is called when production compressedproof is entered.
func (s *BasemetamathListener) EnterCompressedproof(ctx *CompressedproofContext) {}

// ExitCompressedproof is called when production compressedproof is exited.
func (s *BasemetamathListener) ExitCompressedproof(ctx *CompressedproofContext) {}

// EnterTypecode is called when production typecode is entered.
func (s *BasemetamathListener) EnterTypecode(ctx *TypecodeContext) {}

// ExitTypecode is called when production typecode is exited.
func (s *BasemetamathListener) ExitTypecode(ctx *TypecodeContext) {}

// EnterMathsymbol is called when production mathsymbol is entered.
func (s *BasemetamathListener) EnterMathsymbol(ctx *MathsymbolContext) {}

// ExitMathsymbol is called when production mathsymbol is exited.
func (s *BasemetamathListener) ExitMathsymbol(ctx *MathsymbolContext) {}

// EnterFilename is called when production filename is entered.
func (s *BasemetamathListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BasemetamathListener) ExitFilename(ctx *FilenameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasemetamathListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasemetamathListener) ExitConstant(ctx *ConstantContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasemetamathListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasemetamathListener) ExitVariable(ctx *VariableContext) {}
