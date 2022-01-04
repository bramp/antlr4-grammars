// Code generated from PCRE.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pcre // PCRE
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PCREListener is a complete listener for a parse tree produced by PCREParser.
type PCREListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterAlternation is called when entering the alternation production.
	EnterAlternation(c *AlternationContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterQuantifier is called when entering the quantifier production.
	EnterQuantifier(c *QuantifierContext)

	// EnterQuantifier_type is called when entering the quantifier_type production.
	EnterQuantifier_type(c *Quantifier_typeContext)

	// EnterCharacter_class is called when entering the character_class production.
	EnterCharacter_class(c *Character_classContext)

	// EnterBackreference is called when entering the backreference production.
	EnterBackreference(c *BackreferenceContext)

	// EnterBackreference_or_octal is called when entering the backreference_or_octal production.
	EnterBackreference_or_octal(c *Backreference_or_octalContext)

	// EnterCapture is called when entering the capture production.
	EnterCapture(c *CaptureContext)

	// EnterNon_capture is called when entering the non_capture production.
	EnterNon_capture(c *Non_captureContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOption_flags is called when entering the option_flags production.
	EnterOption_flags(c *Option_flagsContext)

	// EnterOption_flag is called when entering the option_flag production.
	EnterOption_flag(c *Option_flagContext)

	// EnterLook_around is called when entering the look_around production.
	EnterLook_around(c *Look_aroundContext)

	// EnterSubroutine_reference is called when entering the subroutine_reference production.
	EnterSubroutine_reference(c *Subroutine_referenceContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterBacktrack_control is called when entering the backtrack_control production.
	EnterBacktrack_control(c *Backtrack_controlContext)

	// EnterNewline_convention is called when entering the newline_convention production.
	EnterNewline_convention(c *Newline_conventionContext)

	// EnterCallout is called when entering the callout production.
	EnterCallout(c *CalloutContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterCc_atom is called when entering the cc_atom production.
	EnterCc_atom(c *Cc_atomContext)

	// EnterShared_atom is called when entering the shared_atom production.
	EnterShared_atom(c *Shared_atomContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterCc_literal is called when entering the cc_literal production.
	EnterCc_literal(c *Cc_literalContext)

	// EnterShared_literal is called when entering the shared_literal production.
	EnterShared_literal(c *Shared_literalContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterOctal_char is called when entering the octal_char production.
	EnterOctal_char(c *Octal_charContext)

	// EnterOctal_digit is called when entering the octal_digit production.
	EnterOctal_digit(c *Octal_digitContext)

	// EnterDigits is called when entering the digits production.
	EnterDigits(c *DigitsContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterAlpha_nums is called when entering the alpha_nums production.
	EnterAlpha_nums(c *Alpha_numsContext)

	// EnterNon_close_parens is called when entering the non_close_parens production.
	EnterNon_close_parens(c *Non_close_parensContext)

	// EnterNon_close_paren is called when entering the non_close_paren production.
	EnterNon_close_paren(c *Non_close_parenContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitAlternation is called when exiting the alternation production.
	ExitAlternation(c *AlternationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitQuantifier is called when exiting the quantifier production.
	ExitQuantifier(c *QuantifierContext)

	// ExitQuantifier_type is called when exiting the quantifier_type production.
	ExitQuantifier_type(c *Quantifier_typeContext)

	// ExitCharacter_class is called when exiting the character_class production.
	ExitCharacter_class(c *Character_classContext)

	// ExitBackreference is called when exiting the backreference production.
	ExitBackreference(c *BackreferenceContext)

	// ExitBackreference_or_octal is called when exiting the backreference_or_octal production.
	ExitBackreference_or_octal(c *Backreference_or_octalContext)

	// ExitCapture is called when exiting the capture production.
	ExitCapture(c *CaptureContext)

	// ExitNon_capture is called when exiting the non_capture production.
	ExitNon_capture(c *Non_captureContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOption_flags is called when exiting the option_flags production.
	ExitOption_flags(c *Option_flagsContext)

	// ExitOption_flag is called when exiting the option_flag production.
	ExitOption_flag(c *Option_flagContext)

	// ExitLook_around is called when exiting the look_around production.
	ExitLook_around(c *Look_aroundContext)

	// ExitSubroutine_reference is called when exiting the subroutine_reference production.
	ExitSubroutine_reference(c *Subroutine_referenceContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitBacktrack_control is called when exiting the backtrack_control production.
	ExitBacktrack_control(c *Backtrack_controlContext)

	// ExitNewline_convention is called when exiting the newline_convention production.
	ExitNewline_convention(c *Newline_conventionContext)

	// ExitCallout is called when exiting the callout production.
	ExitCallout(c *CalloutContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitCc_atom is called when exiting the cc_atom production.
	ExitCc_atom(c *Cc_atomContext)

	// ExitShared_atom is called when exiting the shared_atom production.
	ExitShared_atom(c *Shared_atomContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitCc_literal is called when exiting the cc_literal production.
	ExitCc_literal(c *Cc_literalContext)

	// ExitShared_literal is called when exiting the shared_literal production.
	ExitShared_literal(c *Shared_literalContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitOctal_char is called when exiting the octal_char production.
	ExitOctal_char(c *Octal_charContext)

	// ExitOctal_digit is called when exiting the octal_digit production.
	ExitOctal_digit(c *Octal_digitContext)

	// ExitDigits is called when exiting the digits production.
	ExitDigits(c *DigitsContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitAlpha_nums is called when exiting the alpha_nums production.
	ExitAlpha_nums(c *Alpha_numsContext)

	// ExitNon_close_parens is called when exiting the non_close_parens production.
	ExitNon_close_parens(c *Non_close_parensContext)

	// ExitNon_close_paren is called when exiting the non_close_paren production.
	ExitNon_close_paren(c *Non_close_parenContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)
}
