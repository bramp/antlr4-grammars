// Generated from CLIF.g4 by ANTLR 4.7.

package clif // CLIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CLIFParser.
type CLIFVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CLIFParser#termseq.
	VisitTermseq(ctx *TermseqContext) interface{}

	// Visit a parse tree produced by CLIFParser#interpretedname.
	VisitInterpretedname(ctx *InterpretednameContext) interface{}

	// Visit a parse tree produced by CLIFParser#interpretablename.
	VisitInterpretablename(ctx *InterpretablenameContext) interface{}

	// Visit a parse tree produced by CLIFParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by CLIFParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by CLIFParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by CLIFParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by CLIFParser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by CLIFParser#atomsent.
	VisitAtomsent(ctx *AtomsentContext) interface{}

	// Visit a parse tree produced by CLIFParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by CLIFParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by CLIFParser#boolsent.
	VisitBoolsent(ctx *BoolsentContext) interface{}

	// Visit a parse tree produced by CLIFParser#quantsent.
	VisitQuantsent(ctx *QuantsentContext) interface{}

	// Visit a parse tree produced by CLIFParser#boundlist.
	VisitBoundlist(ctx *BoundlistContext) interface{}

	// Visit a parse tree produced by CLIFParser#commentsent.
	VisitCommentsent(ctx *CommentsentContext) interface{}

	// Visit a parse tree produced by CLIFParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by CLIFParser#phrase.
	VisitPhrase(ctx *PhraseContext) interface{}

	// Visit a parse tree produced by CLIFParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by CLIFParser#cltext.
	VisitCltext(ctx *CltextContext) interface{}

	// Visit a parse tree produced by CLIFParser#namedtext.
	VisitNamedtext(ctx *NamedtextContext) interface{}
}
