// Code generated from prolog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package prolog // prolog
import "github.com/antlr/antlr4/runtime/Go/antlr"

// prologListener is a complete listener for a parse tree produced by prologParser.
type prologListener interface {
	antlr.ParseTreeListener

	// EnterP_text is called when entering the p_text production.
	EnterP_text(c *P_textContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterTermlist is called when entering the termlist production.
	EnterTermlist(c *TermlistContext)

	// EnterAtom_term is called when entering the atom_term production.
	EnterAtom_term(c *Atom_termContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterBraced_term is called when entering the braced_term production.
	EnterBraced_term(c *Braced_termContext)

	// EnterList_term is called when entering the list_term production.
	EnterList_term(c *List_termContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterCompound_term is called when entering the compound_term production.
	EnterCompound_term(c *Compound_termContext)

	// EnterInteger_term is called when entering the integer_term production.
	EnterInteger_term(c *Integer_termContext)

	// EnterCurly_bracketed_term is called when entering the curly_bracketed_term production.
	EnterCurly_bracketed_term(c *Curly_bracketed_termContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterEmpty_list is called when entering the empty_list production.
	EnterEmpty_list(c *Empty_listContext)

	// EnterEmpty_braces is called when entering the empty_braces production.
	EnterEmpty_braces(c *Empty_bracesContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterGraphic is called when entering the graphic production.
	EnterGraphic(c *GraphicContext)

	// EnterQuoted_string is called when entering the quoted_string production.
	EnterQuoted_string(c *Quoted_stringContext)

	// EnterDq_string is called when entering the dq_string production.
	EnterDq_string(c *Dq_stringContext)

	// EnterBackq_string is called when entering the backq_string production.
	EnterBackq_string(c *Backq_stringContext)

	// EnterSemicolon is called when entering the semicolon production.
	EnterSemicolon(c *SemicolonContext)

	// EnterCut is called when entering the cut production.
	EnterCut(c *CutContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// ExitP_text is called when exiting the p_text production.
	ExitP_text(c *P_textContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitTermlist is called when exiting the termlist production.
	ExitTermlist(c *TermlistContext)

	// ExitAtom_term is called when exiting the atom_term production.
	ExitAtom_term(c *Atom_termContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitBraced_term is called when exiting the braced_term production.
	ExitBraced_term(c *Braced_termContext)

	// ExitList_term is called when exiting the list_term production.
	ExitList_term(c *List_termContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitCompound_term is called when exiting the compound_term production.
	ExitCompound_term(c *Compound_termContext)

	// ExitInteger_term is called when exiting the integer_term production.
	ExitInteger_term(c *Integer_termContext)

	// ExitCurly_bracketed_term is called when exiting the curly_bracketed_term production.
	ExitCurly_bracketed_term(c *Curly_bracketed_termContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitEmpty_list is called when exiting the empty_list production.
	ExitEmpty_list(c *Empty_listContext)

	// ExitEmpty_braces is called when exiting the empty_braces production.
	ExitEmpty_braces(c *Empty_bracesContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitGraphic is called when exiting the graphic production.
	ExitGraphic(c *GraphicContext)

	// ExitQuoted_string is called when exiting the quoted_string production.
	ExitQuoted_string(c *Quoted_stringContext)

	// ExitDq_string is called when exiting the dq_string production.
	ExitDq_string(c *Dq_stringContext)

	// ExitBackq_string is called when exiting the backq_string production.
	ExitBackq_string(c *Backq_stringContext)

	// ExitSemicolon is called when exiting the semicolon production.
	ExitSemicolon(c *SemicolonContext)

	// ExitCut is called when exiting the cut production.
	ExitCut(c *CutContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)
}
