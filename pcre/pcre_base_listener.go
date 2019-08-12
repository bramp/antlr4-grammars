// Code generated from PCRE.g4 by ANTLR 4.7.2. DO NOT EDIT.

package pcre // PCRE
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePCREListener is a complete listener for a parse tree produced by PCREParser.
type BasePCREListener struct{}

var _ PCREListener = &BasePCREListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePCREListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePCREListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePCREListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePCREListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BasePCREListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BasePCREListener) ExitParse(ctx *ParseContext) {}

// EnterAlternation is called when production alternation is entered.
func (s *BasePCREListener) EnterAlternation(ctx *AlternationContext) {}

// ExitAlternation is called when production alternation is exited.
func (s *BasePCREListener) ExitAlternation(ctx *AlternationContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePCREListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePCREListener) ExitExpr(ctx *ExprContext) {}

// EnterElement is called when production element is entered.
func (s *BasePCREListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasePCREListener) ExitElement(ctx *ElementContext) {}

// EnterQuantifier is called when production quantifier is entered.
func (s *BasePCREListener) EnterQuantifier(ctx *QuantifierContext) {}

// ExitQuantifier is called when production quantifier is exited.
func (s *BasePCREListener) ExitQuantifier(ctx *QuantifierContext) {}

// EnterQuantifier_type is called when production quantifier_type is entered.
func (s *BasePCREListener) EnterQuantifier_type(ctx *Quantifier_typeContext) {}

// ExitQuantifier_type is called when production quantifier_type is exited.
func (s *BasePCREListener) ExitQuantifier_type(ctx *Quantifier_typeContext) {}

// EnterCharacter_class is called when production character_class is entered.
func (s *BasePCREListener) EnterCharacter_class(ctx *Character_classContext) {}

// ExitCharacter_class is called when production character_class is exited.
func (s *BasePCREListener) ExitCharacter_class(ctx *Character_classContext) {}

// EnterBackreference is called when production backreference is entered.
func (s *BasePCREListener) EnterBackreference(ctx *BackreferenceContext) {}

// ExitBackreference is called when production backreference is exited.
func (s *BasePCREListener) ExitBackreference(ctx *BackreferenceContext) {}

// EnterBackreference_or_octal is called when production backreference_or_octal is entered.
func (s *BasePCREListener) EnterBackreference_or_octal(ctx *Backreference_or_octalContext) {}

// ExitBackreference_or_octal is called when production backreference_or_octal is exited.
func (s *BasePCREListener) ExitBackreference_or_octal(ctx *Backreference_or_octalContext) {}

// EnterCapture is called when production capture is entered.
func (s *BasePCREListener) EnterCapture(ctx *CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *BasePCREListener) ExitCapture(ctx *CaptureContext) {}

// EnterNon_capture is called when production non_capture is entered.
func (s *BasePCREListener) EnterNon_capture(ctx *Non_captureContext) {}

// ExitNon_capture is called when production non_capture is exited.
func (s *BasePCREListener) ExitNon_capture(ctx *Non_captureContext) {}

// EnterComment is called when production comment is entered.
func (s *BasePCREListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasePCREListener) ExitComment(ctx *CommentContext) {}

// EnterOption is called when production option is entered.
func (s *BasePCREListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BasePCREListener) ExitOption(ctx *OptionContext) {}

// EnterOption_flags is called when production option_flags is entered.
func (s *BasePCREListener) EnterOption_flags(ctx *Option_flagsContext) {}

// ExitOption_flags is called when production option_flags is exited.
func (s *BasePCREListener) ExitOption_flags(ctx *Option_flagsContext) {}

// EnterOption_flag is called when production option_flag is entered.
func (s *BasePCREListener) EnterOption_flag(ctx *Option_flagContext) {}

// ExitOption_flag is called when production option_flag is exited.
func (s *BasePCREListener) ExitOption_flag(ctx *Option_flagContext) {}

// EnterLook_around is called when production look_around is entered.
func (s *BasePCREListener) EnterLook_around(ctx *Look_aroundContext) {}

// ExitLook_around is called when production look_around is exited.
func (s *BasePCREListener) ExitLook_around(ctx *Look_aroundContext) {}

// EnterSubroutine_reference is called when production subroutine_reference is entered.
func (s *BasePCREListener) EnterSubroutine_reference(ctx *Subroutine_referenceContext) {}

// ExitSubroutine_reference is called when production subroutine_reference is exited.
func (s *BasePCREListener) ExitSubroutine_reference(ctx *Subroutine_referenceContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BasePCREListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BasePCREListener) ExitConditional(ctx *ConditionalContext) {}

// EnterBacktrack_control is called when production backtrack_control is entered.
func (s *BasePCREListener) EnterBacktrack_control(ctx *Backtrack_controlContext) {}

// ExitBacktrack_control is called when production backtrack_control is exited.
func (s *BasePCREListener) ExitBacktrack_control(ctx *Backtrack_controlContext) {}

// EnterNewline_convention is called when production newline_convention is entered.
func (s *BasePCREListener) EnterNewline_convention(ctx *Newline_conventionContext) {}

// ExitNewline_convention is called when production newline_convention is exited.
func (s *BasePCREListener) ExitNewline_convention(ctx *Newline_conventionContext) {}

// EnterCallout is called when production callout is entered.
func (s *BasePCREListener) EnterCallout(ctx *CalloutContext) {}

// ExitCallout is called when production callout is exited.
func (s *BasePCREListener) ExitCallout(ctx *CalloutContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePCREListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePCREListener) ExitAtom(ctx *AtomContext) {}

// EnterCc_atom is called when production cc_atom is entered.
func (s *BasePCREListener) EnterCc_atom(ctx *Cc_atomContext) {}

// ExitCc_atom is called when production cc_atom is exited.
func (s *BasePCREListener) ExitCc_atom(ctx *Cc_atomContext) {}

// EnterShared_atom is called when production shared_atom is entered.
func (s *BasePCREListener) EnterShared_atom(ctx *Shared_atomContext) {}

// ExitShared_atom is called when production shared_atom is exited.
func (s *BasePCREListener) ExitShared_atom(ctx *Shared_atomContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePCREListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePCREListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCc_literal is called when production cc_literal is entered.
func (s *BasePCREListener) EnterCc_literal(ctx *Cc_literalContext) {}

// ExitCc_literal is called when production cc_literal is exited.
func (s *BasePCREListener) ExitCc_literal(ctx *Cc_literalContext) {}

// EnterShared_literal is called when production shared_literal is entered.
func (s *BasePCREListener) EnterShared_literal(ctx *Shared_literalContext) {}

// ExitShared_literal is called when production shared_literal is exited.
func (s *BasePCREListener) ExitShared_literal(ctx *Shared_literalContext) {}

// EnterNumber is called when production number is entered.
func (s *BasePCREListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasePCREListener) ExitNumber(ctx *NumberContext) {}

// EnterOctal_char is called when production octal_char is entered.
func (s *BasePCREListener) EnterOctal_char(ctx *Octal_charContext) {}

// ExitOctal_char is called when production octal_char is exited.
func (s *BasePCREListener) ExitOctal_char(ctx *Octal_charContext) {}

// EnterOctal_digit is called when production octal_digit is entered.
func (s *BasePCREListener) EnterOctal_digit(ctx *Octal_digitContext) {}

// ExitOctal_digit is called when production octal_digit is exited.
func (s *BasePCREListener) ExitOctal_digit(ctx *Octal_digitContext) {}

// EnterDigits is called when production digits is entered.
func (s *BasePCREListener) EnterDigits(ctx *DigitsContext) {}

// ExitDigits is called when production digits is exited.
func (s *BasePCREListener) ExitDigits(ctx *DigitsContext) {}

// EnterDigit is called when production digit is entered.
func (s *BasePCREListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BasePCREListener) ExitDigit(ctx *DigitContext) {}

// EnterName is called when production name is entered.
func (s *BasePCREListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasePCREListener) ExitName(ctx *NameContext) {}

// EnterAlpha_nums is called when production alpha_nums is entered.
func (s *BasePCREListener) EnterAlpha_nums(ctx *Alpha_numsContext) {}

// ExitAlpha_nums is called when production alpha_nums is exited.
func (s *BasePCREListener) ExitAlpha_nums(ctx *Alpha_numsContext) {}

// EnterNon_close_parens is called when production non_close_parens is entered.
func (s *BasePCREListener) EnterNon_close_parens(ctx *Non_close_parensContext) {}

// ExitNon_close_parens is called when production non_close_parens is exited.
func (s *BasePCREListener) ExitNon_close_parens(ctx *Non_close_parensContext) {}

// EnterNon_close_paren is called when production non_close_paren is entered.
func (s *BasePCREListener) EnterNon_close_paren(ctx *Non_close_parenContext) {}

// ExitNon_close_paren is called when production non_close_paren is exited.
func (s *BasePCREListener) ExitNon_close_paren(ctx *Non_close_parenContext) {}

// EnterLetter is called when production letter is entered.
func (s *BasePCREListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BasePCREListener) ExitLetter(ctx *LetterContext) {}
