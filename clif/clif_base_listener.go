// Code generated from CLIF.g4 by ANTLR 4.9.3. DO NOT EDIT.

package clif // CLIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCLIFListener is a complete listener for a parse tree produced by CLIFParser.
type BaseCLIFListener struct{}

var _ CLIFListener = &BaseCLIFListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCLIFListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCLIFListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCLIFListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCLIFListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTermseq is called when production termseq is entered.
func (s *BaseCLIFListener) EnterTermseq(ctx *TermseqContext) {}

// ExitTermseq is called when production termseq is exited.
func (s *BaseCLIFListener) ExitTermseq(ctx *TermseqContext) {}

// EnterInterpretedname is called when production interpretedname is entered.
func (s *BaseCLIFListener) EnterInterpretedname(ctx *InterpretednameContext) {}

// ExitInterpretedname is called when production interpretedname is exited.
func (s *BaseCLIFListener) ExitInterpretedname(ctx *InterpretednameContext) {}

// EnterInterpretablename is called when production interpretablename is entered.
func (s *BaseCLIFListener) EnterInterpretablename(ctx *InterpretablenameContext) {}

// ExitInterpretablename is called when production interpretablename is exited.
func (s *BaseCLIFListener) ExitInterpretablename(ctx *InterpretablenameContext) {}

// EnterName is called when production name is entered.
func (s *BaseCLIFListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseCLIFListener) ExitName(ctx *NameContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseCLIFListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseCLIFListener) ExitTerm(ctx *TermContext) {}

// EnterOperator_ is called when production operator_ is entered.
func (s *BaseCLIFListener) EnterOperator_(ctx *Operator_Context) {}

// ExitOperator_ is called when production operator_ is exited.
func (s *BaseCLIFListener) ExitOperator_(ctx *Operator_Context) {}

// EnterEquation is called when production equation is entered.
func (s *BaseCLIFListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BaseCLIFListener) ExitEquation(ctx *EquationContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseCLIFListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseCLIFListener) ExitSentence(ctx *SentenceContext) {}

// EnterAtomsent is called when production atomsent is entered.
func (s *BaseCLIFListener) EnterAtomsent(ctx *AtomsentContext) {}

// ExitAtomsent is called when production atomsent is exited.
func (s *BaseCLIFListener) ExitAtomsent(ctx *AtomsentContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseCLIFListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseCLIFListener) ExitAtom(ctx *AtomContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseCLIFListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseCLIFListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBoolsent is called when production boolsent is entered.
func (s *BaseCLIFListener) EnterBoolsent(ctx *BoolsentContext) {}

// ExitBoolsent is called when production boolsent is exited.
func (s *BaseCLIFListener) ExitBoolsent(ctx *BoolsentContext) {}

// EnterQuantsent is called when production quantsent is entered.
func (s *BaseCLIFListener) EnterQuantsent(ctx *QuantsentContext) {}

// ExitQuantsent is called when production quantsent is exited.
func (s *BaseCLIFListener) ExitQuantsent(ctx *QuantsentContext) {}

// EnterBoundlist is called when production boundlist is entered.
func (s *BaseCLIFListener) EnterBoundlist(ctx *BoundlistContext) {}

// ExitBoundlist is called when production boundlist is exited.
func (s *BaseCLIFListener) ExitBoundlist(ctx *BoundlistContext) {}

// EnterCommentsent is called when production commentsent is entered.
func (s *BaseCLIFListener) EnterCommentsent(ctx *CommentsentContext) {}

// ExitCommentsent is called when production commentsent is exited.
func (s *BaseCLIFListener) ExitCommentsent(ctx *CommentsentContext) {}

// EnterModule is called when production module is entered.
func (s *BaseCLIFListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseCLIFListener) ExitModule(ctx *ModuleContext) {}

// EnterPhrase is called when production phrase is entered.
func (s *BaseCLIFListener) EnterPhrase(ctx *PhraseContext) {}

// ExitPhrase is called when production phrase is exited.
func (s *BaseCLIFListener) ExitPhrase(ctx *PhraseContext) {}

// EnterText is called when production text is entered.
func (s *BaseCLIFListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseCLIFListener) ExitText(ctx *TextContext) {}

// EnterCltext is called when production cltext is entered.
func (s *BaseCLIFListener) EnterCltext(ctx *CltextContext) {}

// ExitCltext is called when production cltext is exited.
func (s *BaseCLIFListener) ExitCltext(ctx *CltextContext) {}

// EnterNamedtext is called when production namedtext is entered.
func (s *BaseCLIFListener) EnterNamedtext(ctx *NamedtextContext) {}

// ExitNamedtext is called when production namedtext is exited.
func (s *BaseCLIFListener) ExitNamedtext(ctx *NamedtextContext) {}
