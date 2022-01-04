// Code generated from guido.g4 by ANTLR 4.9.3. DO NOT EDIT.

package guido // guido
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseguidoListener is a complete listener for a parse tree produced by guidoParser.
type BaseguidoListener struct{}

var _ guidoListener = &BaseguidoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseguidoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseguidoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseguidoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseguidoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseguidoListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseguidoListener) ExitProg(ctx *ProgContext) {}

// EnterSegment is called when production segment is entered.
func (s *BaseguidoListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BaseguidoListener) ExitSegment(ctx *SegmentContext) {}

// EnterSequencelist is called when production sequencelist is entered.
func (s *BaseguidoListener) EnterSequencelist(ctx *SequencelistContext) {}

// ExitSequencelist is called when production sequencelist is exited.
func (s *BaseguidoListener) ExitSequencelist(ctx *SequencelistContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BaseguidoListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BaseguidoListener) ExitSequence(ctx *SequenceContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseguidoListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseguidoListener) ExitTag(ctx *TagContext) {}

// EnterTagname is called when production tagname is entered.
func (s *BaseguidoListener) EnterTagname(ctx *TagnameContext) {}

// ExitTagname is called when production tagname is exited.
func (s *BaseguidoListener) ExitTagname(ctx *TagnameContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseguidoListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseguidoListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseguidoListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseguidoListener) ExitParameter(ctx *ParameterContext) {}

// EnterKvpair is called when production kvpair is entered.
func (s *BaseguidoListener) EnterKvpair(ctx *KvpairContext) {}

// ExitKvpair is called when production kvpair is exited.
func (s *BaseguidoListener) ExitKvpair(ctx *KvpairContext) {}

// EnterNotes is called when production notes is entered.
func (s *BaseguidoListener) EnterNotes(ctx *NotesContext) {}

// ExitNotes is called when production notes is exited.
func (s *BaseguidoListener) ExitNotes(ctx *NotesContext) {}

// EnterNote is called when production note is entered.
func (s *BaseguidoListener) EnterNote(ctx *NoteContext) {}

// ExitNote is called when production note is exited.
func (s *BaseguidoListener) ExitNote(ctx *NoteContext) {}

// EnterChord is called when production chord is entered.
func (s *BaseguidoListener) EnterChord(ctx *ChordContext) {}

// ExitChord is called when production chord is exited.
func (s *BaseguidoListener) ExitChord(ctx *ChordContext) {}

// EnterNotename is called when production notename is entered.
func (s *BaseguidoListener) EnterNotename(ctx *NotenameContext) {}

// ExitNotename is called when production notename is exited.
func (s *BaseguidoListener) ExitNotename(ctx *NotenameContext) {}

// EnterAccidental is called when production accidental is entered.
func (s *BaseguidoListener) EnterAccidental(ctx *AccidentalContext) {}

// ExitAccidental is called when production accidental is exited.
func (s *BaseguidoListener) ExitAccidental(ctx *AccidentalContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseguidoListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseguidoListener) ExitNumber(ctx *NumberContext) {}

// EnterOctave is called when production octave is entered.
func (s *BaseguidoListener) EnterOctave(ctx *OctaveContext) {}

// ExitOctave is called when production octave is exited.
func (s *BaseguidoListener) ExitOctave(ctx *OctaveContext) {}

// EnterFraction is called when production fraction is entered.
func (s *BaseguidoListener) EnterFraction(ctx *FractionContext) {}

// ExitFraction is called when production fraction is exited.
func (s *BaseguidoListener) ExitFraction(ctx *FractionContext) {}

// EnterDuration is called when production duration is entered.
func (s *BaseguidoListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BaseguidoListener) ExitDuration(ctx *DurationContext) {}

// EnterDotting is called when production dotting is entered.
func (s *BaseguidoListener) EnterDotting(ctx *DottingContext) {}

// ExitDotting is called when production dotting is exited.
func (s *BaseguidoListener) ExitDotting(ctx *DottingContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseguidoListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseguidoListener) ExitTitle(ctx *TitleContext) {}

// EnterTempo is called when production tempo is entered.
func (s *BaseguidoListener) EnterTempo(ctx *TempoContext) {}

// ExitTempo is called when production tempo is exited.
func (s *BaseguidoListener) ExitTempo(ctx *TempoContext) {}

// EnterClef is called when production clef is entered.
func (s *BaseguidoListener) EnterClef(ctx *ClefContext) {}

// ExitClef is called when production clef is exited.
func (s *BaseguidoListener) ExitClef(ctx *ClefContext) {}

// EnterMeter is called when production meter is entered.
func (s *BaseguidoListener) EnterMeter(ctx *MeterContext) {}

// ExitMeter is called when production meter is exited.
func (s *BaseguidoListener) ExitMeter(ctx *MeterContext) {}

// EnterSlur is called when production slur is entered.
func (s *BaseguidoListener) EnterSlur(ctx *SlurContext) {}

// ExitSlur is called when production slur is exited.
func (s *BaseguidoListener) ExitSlur(ctx *SlurContext) {}

// EnterKey is called when production key is entered.
func (s *BaseguidoListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseguidoListener) ExitKey(ctx *KeyContext) {}

// EnterBarformat is called when production barformat is entered.
func (s *BaseguidoListener) EnterBarformat(ctx *BarformatContext) {}

// ExitBarformat is called when production barformat is exited.
func (s *BaseguidoListener) ExitBarformat(ctx *BarformatContext) {}

// EnterStaff is called when production staff is entered.
func (s *BaseguidoListener) EnterStaff(ctx *StaffContext) {}

// ExitStaff is called when production staff is exited.
func (s *BaseguidoListener) ExitStaff(ctx *StaffContext) {}

// EnterRepeatEnd is called when production repeatEnd is entered.
func (s *BaseguidoListener) EnterRepeatEnd(ctx *RepeatEndContext) {}

// ExitRepeatEnd is called when production repeatEnd is exited.
func (s *BaseguidoListener) ExitRepeatEnd(ctx *RepeatEndContext) {}

// EnterT is called when production t is entered.
func (s *BaseguidoListener) EnterT(ctx *TContext) {}

// ExitT is called when production t is exited.
func (s *BaseguidoListener) ExitT(ctx *TContext) {}
