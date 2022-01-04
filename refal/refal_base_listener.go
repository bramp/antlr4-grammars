// Code generated from refal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package refal // refal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaserefalListener is a complete listener for a parse tree produced by refalParser.
type BaserefalListener struct{}

var _ refalListener = &BaserefalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaserefalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaserefalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaserefalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaserefalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaserefalListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaserefalListener) ExitProgram(ctx *ProgramContext) {}

// EnterF_definition is called when production f_definition is entered.
func (s *BaserefalListener) EnterF_definition(ctx *F_definitionContext) {}

// ExitF_definition is called when production f_definition is exited.
func (s *BaserefalListener) ExitF_definition(ctx *F_definitionContext) {}

// EnterExternal_decl is called when production external_decl is entered.
func (s *BaserefalListener) EnterExternal_decl(ctx *External_declContext) {}

// ExitExternal_decl is called when production external_decl is exited.
func (s *BaserefalListener) ExitExternal_decl(ctx *External_declContext) {}

// EnterF_name_list is called when production f_name_list is entered.
func (s *BaserefalListener) EnterF_name_list(ctx *F_name_listContext) {}

// ExitF_name_list is called when production f_name_list is exited.
func (s *BaserefalListener) ExitF_name_list(ctx *F_name_listContext) {}

// EnterF_name is called when production f_name is entered.
func (s *BaserefalListener) EnterF_name(ctx *F_nameContext) {}

// ExitF_name is called when production f_name is exited.
func (s *BaserefalListener) ExitF_name(ctx *F_nameContext) {}

// EnterBlock_ is called when production block_ is entered.
func (s *BaserefalListener) EnterBlock_(ctx *Block_Context) {}

// ExitBlock_ is called when production block_ is exited.
func (s *BaserefalListener) ExitBlock_(ctx *Block_Context) {}

// EnterSentence is called when production sentence is entered.
func (s *BaserefalListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaserefalListener) ExitSentence(ctx *SentenceContext) {}

// EnterLeft_side is called when production left_side is entered.
func (s *BaserefalListener) EnterLeft_side(ctx *Left_sideContext) {}

// ExitLeft_side is called when production left_side is exited.
func (s *BaserefalListener) ExitLeft_side(ctx *Left_sideContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaserefalListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaserefalListener) ExitConditions(ctx *ConditionsContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaserefalListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaserefalListener) ExitPattern(ctx *PatternContext) {}

// EnterArg_ is called when production arg_ is entered.
func (s *BaserefalListener) EnterArg_(ctx *Arg_Context) {}

// ExitArg_ is called when production arg_ is exited.
func (s *BaserefalListener) ExitArg_(ctx *Arg_Context) {}

// EnterRight_side is called when production right_side is entered.
func (s *BaserefalListener) EnterRight_side(ctx *Right_sideContext) {}

// ExitRight_side is called when production right_side is exited.
func (s *BaserefalListener) ExitRight_side(ctx *Right_sideContext) {}

// EnterExpression_ is called when production expression_ is entered.
func (s *BaserefalListener) EnterExpression_(ctx *Expression_Context) {}

// ExitExpression_ is called when production expression_ is exited.
func (s *BaserefalListener) ExitExpression_(ctx *Expression_Context) {}

// EnterTerm_ is called when production term_ is entered.
func (s *BaserefalListener) EnterTerm_(ctx *Term_Context) {}

// ExitTerm_ is called when production term_ is exited.
func (s *BaserefalListener) ExitTerm_(ctx *Term_Context) {}

// EnterBlock_ending is called when production block_ending is entered.
func (s *BaserefalListener) EnterBlock_ending(ctx *Block_endingContext) {}

// ExitBlock_ending is called when production block_ending is exited.
func (s *BaserefalListener) ExitBlock_ending(ctx *Block_endingContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaserefalListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaserefalListener) ExitSymbol(ctx *SymbolContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaserefalListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaserefalListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaserefalListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaserefalListener) ExitVariable(ctx *VariableContext) {}

// EnterSvar is called when production svar is entered.
func (s *BaserefalListener) EnterSvar(ctx *SvarContext) {}

// ExitSvar is called when production svar is exited.
func (s *BaserefalListener) ExitSvar(ctx *SvarContext) {}

// EnterTvar is called when production tvar is entered.
func (s *BaserefalListener) EnterTvar(ctx *TvarContext) {}

// ExitTvar is called when production tvar is exited.
func (s *BaserefalListener) ExitTvar(ctx *TvarContext) {}

// EnterEvar is called when production evar is entered.
func (s *BaserefalListener) EnterEvar(ctx *EvarContext) {}

// ExitEvar is called when production evar is exited.
func (s *BaserefalListener) ExitEvar(ctx *EvarContext) {}

// EnterIndex is called when production index is entered.
func (s *BaserefalListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaserefalListener) ExitIndex(ctx *IndexContext) {}
