// Code generated from prolog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prolog // prolog
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseprologListener is a complete listener for a parse tree produced by prologParser.
type BaseprologListener struct{}

var _ prologListener = &BaseprologListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseprologListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseprologListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseprologListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseprologListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterP_text is called when production p_text is entered.
func (s *BaseprologListener) EnterP_text(ctx *P_textContext) {}

// ExitP_text is called when production p_text is exited.
func (s *BaseprologListener) ExitP_text(ctx *P_textContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseprologListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseprologListener) ExitDirective(ctx *DirectiveContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseprologListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseprologListener) ExitClause(ctx *ClauseContext) {}

// EnterTermlist is called when production termlist is entered.
func (s *BaseprologListener) EnterTermlist(ctx *TermlistContext) {}

// ExitTermlist is called when production termlist is exited.
func (s *BaseprologListener) ExitTermlist(ctx *TermlistContext) {}

// EnterAtom_term is called when production atom_term is entered.
func (s *BaseprologListener) EnterAtom_term(ctx *Atom_termContext) {}

// ExitAtom_term is called when production atom_term is exited.
func (s *BaseprologListener) ExitAtom_term(ctx *Atom_termContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseprologListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseprologListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseprologListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseprologListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterBraced_term is called when production braced_term is entered.
func (s *BaseprologListener) EnterBraced_term(ctx *Braced_termContext) {}

// ExitBraced_term is called when production braced_term is exited.
func (s *BaseprologListener) ExitBraced_term(ctx *Braced_termContext) {}

// EnterList_term is called when production list_term is entered.
func (s *BaseprologListener) EnterList_term(ctx *List_termContext) {}

// ExitList_term is called when production list_term is exited.
func (s *BaseprologListener) ExitList_term(ctx *List_termContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseprologListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseprologListener) ExitVariable(ctx *VariableContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseprologListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseprologListener) ExitFloat(ctx *FloatContext) {}

// EnterCompound_term is called when production compound_term is entered.
func (s *BaseprologListener) EnterCompound_term(ctx *Compound_termContext) {}

// ExitCompound_term is called when production compound_term is exited.
func (s *BaseprologListener) ExitCompound_term(ctx *Compound_termContext) {}

// EnterInteger_term is called when production integer_term is entered.
func (s *BaseprologListener) EnterInteger_term(ctx *Integer_termContext) {}

// ExitInteger_term is called when production integer_term is exited.
func (s *BaseprologListener) ExitInteger_term(ctx *Integer_termContext) {}

// EnterCurly_bracketed_term is called when production curly_bracketed_term is entered.
func (s *BaseprologListener) EnterCurly_bracketed_term(ctx *Curly_bracketed_termContext) {}

// ExitCurly_bracketed_term is called when production curly_bracketed_term is exited.
func (s *BaseprologListener) ExitCurly_bracketed_term(ctx *Curly_bracketed_termContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseprologListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseprologListener) ExitOperator(ctx *OperatorContext) {}

// EnterEmpty_list is called when production empty_list is entered.
func (s *BaseprologListener) EnterEmpty_list(ctx *Empty_listContext) {}

// ExitEmpty_list is called when production empty_list is exited.
func (s *BaseprologListener) ExitEmpty_list(ctx *Empty_listContext) {}

// EnterEmpty_braces is called when production empty_braces is entered.
func (s *BaseprologListener) EnterEmpty_braces(ctx *Empty_bracesContext) {}

// ExitEmpty_braces is called when production empty_braces is exited.
func (s *BaseprologListener) ExitEmpty_braces(ctx *Empty_bracesContext) {}

// EnterName is called when production name is entered.
func (s *BaseprologListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseprologListener) ExitName(ctx *NameContext) {}

// EnterGraphic is called when production graphic is entered.
func (s *BaseprologListener) EnterGraphic(ctx *GraphicContext) {}

// ExitGraphic is called when production graphic is exited.
func (s *BaseprologListener) ExitGraphic(ctx *GraphicContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BaseprologListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BaseprologListener) ExitQuoted_string(ctx *Quoted_stringContext) {}

// EnterDq_string is called when production dq_string is entered.
func (s *BaseprologListener) EnterDq_string(ctx *Dq_stringContext) {}

// ExitDq_string is called when production dq_string is exited.
func (s *BaseprologListener) ExitDq_string(ctx *Dq_stringContext) {}

// EnterBackq_string is called when production backq_string is entered.
func (s *BaseprologListener) EnterBackq_string(ctx *Backq_stringContext) {}

// ExitBackq_string is called when production backq_string is exited.
func (s *BaseprologListener) ExitBackq_string(ctx *Backq_stringContext) {}

// EnterSemicolon is called when production semicolon is entered.
func (s *BaseprologListener) EnterSemicolon(ctx *SemicolonContext) {}

// ExitSemicolon is called when production semicolon is exited.
func (s *BaseprologListener) ExitSemicolon(ctx *SemicolonContext) {}

// EnterCut is called when production cut is entered.
func (s *BaseprologListener) EnterCut(ctx *CutContext) {}

// ExitCut is called when production cut is exited.
func (s *BaseprologListener) ExitCut(ctx *CutContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseprologListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseprologListener) ExitInteger(ctx *IntegerContext) {}
