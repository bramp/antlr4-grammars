// Code generated from metamath.g4 by ANTLR 4.9.3. DO NOT EDIT.

package metamath // metamath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// metamathListener is a complete listener for a parse tree produced by metamathParser.
type metamathListener interface {
	antlr.ParseTreeListener

	// EnterDatabase is called when entering the database production.
	EnterDatabase(c *DatabaseContext)

	// EnterOutermostscopestmt is called when entering the outermostscopestmt production.
	EnterOutermostscopestmt(c *OutermostscopestmtContext)

	// EnterIncludestmt is called when entering the includestmt production.
	EnterIncludestmt(c *IncludestmtContext)

	// EnterConstantstmt is called when entering the constantstmt production.
	EnterConstantstmt(c *ConstantstmtContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterVariablestmt is called when entering the variablestmt production.
	EnterVariablestmt(c *VariablestmtContext)

	// EnterDisjointstmt is called when entering the disjointstmt production.
	EnterDisjointstmt(c *DisjointstmtContext)

	// EnterHypothesisstmt is called when entering the hypothesisstmt production.
	EnterHypothesisstmt(c *HypothesisstmtContext)

	// EnterFloatingstmt is called when entering the floatingstmt production.
	EnterFloatingstmt(c *FloatingstmtContext)

	// EnterEssentialstmt is called when entering the essentialstmt production.
	EnterEssentialstmt(c *EssentialstmtContext)

	// EnterAssertstmt is called when entering the assertstmt production.
	EnterAssertstmt(c *AssertstmtContext)

	// EnterAxiomstmt is called when entering the axiomstmt production.
	EnterAxiomstmt(c *AxiomstmtContext)

	// EnterProvablestmt is called when entering the provablestmt production.
	EnterProvablestmt(c *ProvablestmtContext)

	// EnterProof is called when entering the proof production.
	EnterProof(c *ProofContext)

	// EnterUncompressedproof is called when entering the uncompressedproof production.
	EnterUncompressedproof(c *UncompressedproofContext)

	// EnterCompressedproof is called when entering the compressedproof production.
	EnterCompressedproof(c *CompressedproofContext)

	// EnterTypecode is called when entering the typecode production.
	EnterTypecode(c *TypecodeContext)

	// EnterMathsymbol is called when entering the mathsymbol production.
	EnterMathsymbol(c *MathsymbolContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// ExitDatabase is called when exiting the database production.
	ExitDatabase(c *DatabaseContext)

	// ExitOutermostscopestmt is called when exiting the outermostscopestmt production.
	ExitOutermostscopestmt(c *OutermostscopestmtContext)

	// ExitIncludestmt is called when exiting the includestmt production.
	ExitIncludestmt(c *IncludestmtContext)

	// ExitConstantstmt is called when exiting the constantstmt production.
	ExitConstantstmt(c *ConstantstmtContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitVariablestmt is called when exiting the variablestmt production.
	ExitVariablestmt(c *VariablestmtContext)

	// ExitDisjointstmt is called when exiting the disjointstmt production.
	ExitDisjointstmt(c *DisjointstmtContext)

	// ExitHypothesisstmt is called when exiting the hypothesisstmt production.
	ExitHypothesisstmt(c *HypothesisstmtContext)

	// ExitFloatingstmt is called when exiting the floatingstmt production.
	ExitFloatingstmt(c *FloatingstmtContext)

	// ExitEssentialstmt is called when exiting the essentialstmt production.
	ExitEssentialstmt(c *EssentialstmtContext)

	// ExitAssertstmt is called when exiting the assertstmt production.
	ExitAssertstmt(c *AssertstmtContext)

	// ExitAxiomstmt is called when exiting the axiomstmt production.
	ExitAxiomstmt(c *AxiomstmtContext)

	// ExitProvablestmt is called when exiting the provablestmt production.
	ExitProvablestmt(c *ProvablestmtContext)

	// ExitProof is called when exiting the proof production.
	ExitProof(c *ProofContext)

	// ExitUncompressedproof is called when exiting the uncompressedproof production.
	ExitUncompressedproof(c *UncompressedproofContext)

	// ExitCompressedproof is called when exiting the compressedproof production.
	ExitCompressedproof(c *CompressedproofContext)

	// ExitTypecode is called when exiting the typecode production.
	ExitTypecode(c *TypecodeContext)

	// ExitMathsymbol is called when exiting the mathsymbol production.
	ExitMathsymbol(c *MathsymbolContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)
}
