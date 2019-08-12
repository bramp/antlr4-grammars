// Code generated from CLIF.g4 by ANTLR 4.7.2. DO NOT EDIT.

package clif // CLIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CLIFListener is a complete listener for a parse tree produced by CLIFParser.
type CLIFListener interface {
	antlr.ParseTreeListener

	// EnterTermseq is called when entering the termseq production.
	EnterTermseq(c *TermseqContext)

	// EnterInterpretedname is called when entering the interpretedname production.
	EnterInterpretedname(c *InterpretednameContext)

	// EnterInterpretablename is called when entering the interpretablename production.
	EnterInterpretablename(c *InterpretablenameContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterAtomsent is called when entering the atomsent production.
	EnterAtomsent(c *AtomsentContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterBoolsent is called when entering the boolsent production.
	EnterBoolsent(c *BoolsentContext)

	// EnterQuantsent is called when entering the quantsent production.
	EnterQuantsent(c *QuantsentContext)

	// EnterBoundlist is called when entering the boundlist production.
	EnterBoundlist(c *BoundlistContext)

	// EnterCommentsent is called when entering the commentsent production.
	EnterCommentsent(c *CommentsentContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterPhrase is called when entering the phrase production.
	EnterPhrase(c *PhraseContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterCltext is called when entering the cltext production.
	EnterCltext(c *CltextContext)

	// EnterNamedtext is called when entering the namedtext production.
	EnterNamedtext(c *NamedtextContext)

	// ExitTermseq is called when exiting the termseq production.
	ExitTermseq(c *TermseqContext)

	// ExitInterpretedname is called when exiting the interpretedname production.
	ExitInterpretedname(c *InterpretednameContext)

	// ExitInterpretablename is called when exiting the interpretablename production.
	ExitInterpretablename(c *InterpretablenameContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitAtomsent is called when exiting the atomsent production.
	ExitAtomsent(c *AtomsentContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitBoolsent is called when exiting the boolsent production.
	ExitBoolsent(c *BoolsentContext)

	// ExitQuantsent is called when exiting the quantsent production.
	ExitQuantsent(c *QuantsentContext)

	// ExitBoundlist is called when exiting the boundlist production.
	ExitBoundlist(c *BoundlistContext)

	// ExitCommentsent is called when exiting the commentsent production.
	ExitCommentsent(c *CommentsentContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitPhrase is called when exiting the phrase production.
	ExitPhrase(c *PhraseContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitCltext is called when exiting the cltext production.
	ExitCltext(c *CltextContext)

	// ExitNamedtext is called when exiting the namedtext production.
	ExitNamedtext(c *NamedtextContext)
}
