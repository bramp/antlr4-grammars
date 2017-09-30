// Generated from PCRE.g4 by ANTLR 4.7.

package pcre // PCRE
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePCREVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePCREVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAlternation(ctx *AlternationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitQuantifier(ctx *QuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitQuantifier_type(ctx *Quantifier_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_class(ctx *Character_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBackreference(ctx *BackreferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBackreference_or_octal(ctx *Backreference_or_octalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCapture(ctx *CaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNon_capture(ctx *Non_captureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOption_flags(ctx *Option_flagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOption_flag(ctx *Option_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLook_around(ctx *Look_aroundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitSubroutine_reference(ctx *Subroutine_referenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitConditional(ctx *ConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBacktrack_control(ctx *Backtrack_controlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNewline_convention(ctx *Newline_conventionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCallout(ctx *CalloutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCc_atom(ctx *Cc_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitShared_atom(ctx *Shared_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCc_literal(ctx *Cc_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitShared_literal(ctx *Shared_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOctal_char(ctx *Octal_charContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOctal_digit(ctx *Octal_digitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitDigits(ctx *DigitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAlpha_nums(ctx *Alpha_numsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNon_close_parens(ctx *Non_close_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNon_close_paren(ctx *Non_close_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLetter(ctx *LetterContext) interface{} {
	return v.VisitChildren(ctx)
}
