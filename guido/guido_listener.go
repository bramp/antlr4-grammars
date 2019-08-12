// Code generated from guido.g4 by ANTLR 4.7.2. DO NOT EDIT.

package guido // guido
import "github.com/antlr/antlr4/runtime/Go/antlr"

// guidoListener is a complete listener for a parse tree produced by guidoParser.
type guidoListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterSequencelist is called when entering the sequencelist production.
	EnterSequencelist(c *SequencelistContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterTagname is called when entering the tagname production.
	EnterTagname(c *TagnameContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterKvpair is called when entering the kvpair production.
	EnterKvpair(c *KvpairContext)

	// EnterNotes is called when entering the notes production.
	EnterNotes(c *NotesContext)

	// EnterNote is called when entering the note production.
	EnterNote(c *NoteContext)

	// EnterChord is called when entering the chord production.
	EnterChord(c *ChordContext)

	// EnterNotename is called when entering the notename production.
	EnterNotename(c *NotenameContext)

	// EnterAccidental is called when entering the accidental production.
	EnterAccidental(c *AccidentalContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterOctave is called when entering the octave production.
	EnterOctave(c *OctaveContext)

	// EnterFraction is called when entering the fraction production.
	EnterFraction(c *FractionContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// EnterDotting is called when entering the dotting production.
	EnterDotting(c *DottingContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterTempo is called when entering the tempo production.
	EnterTempo(c *TempoContext)

	// EnterClef is called when entering the clef production.
	EnterClef(c *ClefContext)

	// EnterMeter is called when entering the meter production.
	EnterMeter(c *MeterContext)

	// EnterSlur is called when entering the slur production.
	EnterSlur(c *SlurContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterBarformat is called when entering the barformat production.
	EnterBarformat(c *BarformatContext)

	// EnterStaff is called when entering the staff production.
	EnterStaff(c *StaffContext)

	// EnterRepeatEnd is called when entering the repeatEnd production.
	EnterRepeatEnd(c *RepeatEndContext)

	// EnterT is called when entering the t production.
	EnterT(c *TContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitSequencelist is called when exiting the sequencelist production.
	ExitSequencelist(c *SequencelistContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitTagname is called when exiting the tagname production.
	ExitTagname(c *TagnameContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitKvpair is called when exiting the kvpair production.
	ExitKvpair(c *KvpairContext)

	// ExitNotes is called when exiting the notes production.
	ExitNotes(c *NotesContext)

	// ExitNote is called when exiting the note production.
	ExitNote(c *NoteContext)

	// ExitChord is called when exiting the chord production.
	ExitChord(c *ChordContext)

	// ExitNotename is called when exiting the notename production.
	ExitNotename(c *NotenameContext)

	// ExitAccidental is called when exiting the accidental production.
	ExitAccidental(c *AccidentalContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitOctave is called when exiting the octave production.
	ExitOctave(c *OctaveContext)

	// ExitFraction is called when exiting the fraction production.
	ExitFraction(c *FractionContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)

	// ExitDotting is called when exiting the dotting production.
	ExitDotting(c *DottingContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitTempo is called when exiting the tempo production.
	ExitTempo(c *TempoContext)

	// ExitClef is called when exiting the clef production.
	ExitClef(c *ClefContext)

	// ExitMeter is called when exiting the meter production.
	ExitMeter(c *MeterContext)

	// ExitSlur is called when exiting the slur production.
	ExitSlur(c *SlurContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitBarformat is called when exiting the barformat production.
	ExitBarformat(c *BarformatContext)

	// ExitStaff is called when exiting the staff production.
	ExitStaff(c *StaffContext)

	// ExitRepeatEnd is called when exiting the repeatEnd production.
	ExitRepeatEnd(c *RepeatEndContext)

	// ExitT is called when exiting the t production.
	ExitT(c *TContext)
}
