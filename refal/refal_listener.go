// Code generated from refal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package refal // refal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// refalListener is a complete listener for a parse tree produced by refalParser.
type refalListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterF_definition is called when entering the f_definition production.
	EnterF_definition(c *F_definitionContext)

	// EnterExternal_decl is called when entering the external_decl production.
	EnterExternal_decl(c *External_declContext)

	// EnterF_name_list is called when entering the f_name_list production.
	EnterF_name_list(c *F_name_listContext)

	// EnterF_name is called when entering the f_name production.
	EnterF_name(c *F_nameContext)

	// EnterBlock_ is called when entering the block_ production.
	EnterBlock_(c *Block_Context)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterLeft_side is called when entering the left_side production.
	EnterLeft_side(c *Left_sideContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterArg_ is called when entering the arg_ production.
	EnterArg_(c *Arg_Context)

	// EnterRight_side is called when entering the right_side production.
	EnterRight_side(c *Right_sideContext)

	// EnterExpression_ is called when entering the expression_ production.
	EnterExpression_(c *Expression_Context)

	// EnterTerm_ is called when entering the term_ production.
	EnterTerm_(c *Term_Context)

	// EnterBlock_ending is called when entering the block_ending production.
	EnterBlock_ending(c *Block_endingContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSvar is called when entering the svar production.
	EnterSvar(c *SvarContext)

	// EnterTvar is called when entering the tvar production.
	EnterTvar(c *TvarContext)

	// EnterEvar is called when entering the evar production.
	EnterEvar(c *EvarContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitF_definition is called when exiting the f_definition production.
	ExitF_definition(c *F_definitionContext)

	// ExitExternal_decl is called when exiting the external_decl production.
	ExitExternal_decl(c *External_declContext)

	// ExitF_name_list is called when exiting the f_name_list production.
	ExitF_name_list(c *F_name_listContext)

	// ExitF_name is called when exiting the f_name production.
	ExitF_name(c *F_nameContext)

	// ExitBlock_ is called when exiting the block_ production.
	ExitBlock_(c *Block_Context)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitLeft_side is called when exiting the left_side production.
	ExitLeft_side(c *Left_sideContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitArg_ is called when exiting the arg_ production.
	ExitArg_(c *Arg_Context)

	// ExitRight_side is called when exiting the right_side production.
	ExitRight_side(c *Right_sideContext)

	// ExitExpression_ is called when exiting the expression_ production.
	ExitExpression_(c *Expression_Context)

	// ExitTerm_ is called when exiting the term_ production.
	ExitTerm_(c *Term_Context)

	// ExitBlock_ending is called when exiting the block_ending production.
	ExitBlock_ending(c *Block_endingContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSvar is called when exiting the svar production.
	ExitSvar(c *SvarContext)

	// ExitTvar is called when exiting the tvar production.
	ExitTvar(c *TvarContext)

	// ExitEvar is called when exiting the evar production.
	ExitEvar(c *EvarContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)
}
