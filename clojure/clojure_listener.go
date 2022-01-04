// Code generated from Clojure.g4 by ANTLR 4.9.3. DO NOT EDIT.

package clojure // Clojure
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClojureListener is a complete listener for a parse tree produced by ClojureParser.
type ClojureListener interface {
	antlr.ParseTreeListener

	// EnterFile_ is called when entering the file_ production.
	EnterFile_(c *File_Context)

	// EnterForm is called when entering the form production.
	EnterForm(c *FormContext)

	// EnterForms is called when entering the forms production.
	EnterForms(c *FormsContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterMap_ is called when entering the map_ production.
	EnterMap_(c *Map_Context)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// EnterReader_macro is called when entering the reader_macro production.
	EnterReader_macro(c *Reader_macroContext)

	// EnterQuote is called when entering the quote production.
	EnterQuote(c *QuoteContext)

	// EnterBacktick is called when entering the backtick production.
	EnterBacktick(c *BacktickContext)

	// EnterUnquote is called when entering the unquote production.
	EnterUnquote(c *UnquoteContext)

	// EnterUnquote_splicing is called when entering the unquote_splicing production.
	EnterUnquote_splicing(c *Unquote_splicingContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterDeref is called when entering the deref production.
	EnterDeref(c *DerefContext)

	// EnterGensym is called when entering the gensym production.
	EnterGensym(c *GensymContext)

	// EnterLambda_ is called when entering the lambda_ production.
	EnterLambda_(c *Lambda_Context)

	// EnterMeta_data is called when entering the meta_data production.
	EnterMeta_data(c *Meta_dataContext)

	// EnterVar_quote is called when entering the var_quote production.
	EnterVar_quote(c *Var_quoteContext)

	// EnterHost_expr is called when entering the host_expr production.
	EnterHost_expr(c *Host_exprContext)

	// EnterDiscard is called when entering the discard production.
	EnterDiscard(c *DiscardContext)

	// EnterDispatch is called when entering the dispatch production.
	EnterDispatch(c *DispatchContext)

	// EnterRegex is called when entering the regex production.
	EnterRegex(c *RegexContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterHex_ is called when entering the hex_ production.
	EnterHex_(c *Hex_Context)

	// EnterBin_ is called when entering the bin_ production.
	EnterBin_(c *Bin_Context)

	// EnterBign is called when entering the bign production.
	EnterBign(c *BignContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterCharacter is called when entering the character production.
	EnterCharacter(c *CharacterContext)

	// EnterNamed_char is called when entering the named_char production.
	EnterNamed_char(c *Named_charContext)

	// EnterAny_char is called when entering the any_char production.
	EnterAny_char(c *Any_charContext)

	// EnterU_hex_quad is called when entering the u_hex_quad production.
	EnterU_hex_quad(c *U_hex_quadContext)

	// EnterNil_ is called when entering the nil_ production.
	EnterNil_(c *Nil_Context)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterSimple_keyword is called when entering the simple_keyword production.
	EnterSimple_keyword(c *Simple_keywordContext)

	// EnterMacro_keyword is called when entering the macro_keyword production.
	EnterMacro_keyword(c *Macro_keywordContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterSimple_sym is called when entering the simple_sym production.
	EnterSimple_sym(c *Simple_symContext)

	// EnterNs_symbol is called when entering the ns_symbol production.
	EnterNs_symbol(c *Ns_symbolContext)

	// EnterParam_name is called when entering the param_name production.
	EnterParam_name(c *Param_nameContext)

	// ExitFile_ is called when exiting the file_ production.
	ExitFile_(c *File_Context)

	// ExitForm is called when exiting the form production.
	ExitForm(c *FormContext)

	// ExitForms is called when exiting the forms production.
	ExitForms(c *FormsContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitMap_ is called when exiting the map_ production.
	ExitMap_(c *Map_Context)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)

	// ExitReader_macro is called when exiting the reader_macro production.
	ExitReader_macro(c *Reader_macroContext)

	// ExitQuote is called when exiting the quote production.
	ExitQuote(c *QuoteContext)

	// ExitBacktick is called when exiting the backtick production.
	ExitBacktick(c *BacktickContext)

	// ExitUnquote is called when exiting the unquote production.
	ExitUnquote(c *UnquoteContext)

	// ExitUnquote_splicing is called when exiting the unquote_splicing production.
	ExitUnquote_splicing(c *Unquote_splicingContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitDeref is called when exiting the deref production.
	ExitDeref(c *DerefContext)

	// ExitGensym is called when exiting the gensym production.
	ExitGensym(c *GensymContext)

	// ExitLambda_ is called when exiting the lambda_ production.
	ExitLambda_(c *Lambda_Context)

	// ExitMeta_data is called when exiting the meta_data production.
	ExitMeta_data(c *Meta_dataContext)

	// ExitVar_quote is called when exiting the var_quote production.
	ExitVar_quote(c *Var_quoteContext)

	// ExitHost_expr is called when exiting the host_expr production.
	ExitHost_expr(c *Host_exprContext)

	// ExitDiscard is called when exiting the discard production.
	ExitDiscard(c *DiscardContext)

	// ExitDispatch is called when exiting the dispatch production.
	ExitDispatch(c *DispatchContext)

	// ExitRegex is called when exiting the regex production.
	ExitRegex(c *RegexContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitHex_ is called when exiting the hex_ production.
	ExitHex_(c *Hex_Context)

	// ExitBin_ is called when exiting the bin_ production.
	ExitBin_(c *Bin_Context)

	// ExitBign is called when exiting the bign production.
	ExitBign(c *BignContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitCharacter is called when exiting the character production.
	ExitCharacter(c *CharacterContext)

	// ExitNamed_char is called when exiting the named_char production.
	ExitNamed_char(c *Named_charContext)

	// ExitAny_char is called when exiting the any_char production.
	ExitAny_char(c *Any_charContext)

	// ExitU_hex_quad is called when exiting the u_hex_quad production.
	ExitU_hex_quad(c *U_hex_quadContext)

	// ExitNil_ is called when exiting the nil_ production.
	ExitNil_(c *Nil_Context)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitSimple_keyword is called when exiting the simple_keyword production.
	ExitSimple_keyword(c *Simple_keywordContext)

	// ExitMacro_keyword is called when exiting the macro_keyword production.
	ExitMacro_keyword(c *Macro_keywordContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitSimple_sym is called when exiting the simple_sym production.
	ExitSimple_sym(c *Simple_symContext)

	// ExitNs_symbol is called when exiting the ns_symbol production.
	ExitNs_symbol(c *Ns_symbolContext)

	// ExitParam_name is called when exiting the param_name production.
	ExitParam_name(c *Param_nameContext)
}
