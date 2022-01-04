// Code generated from tcp.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tcp // tcp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tcpListener is a complete listener for a parse tree produced by tcpParser.
type tcpListener interface {
	antlr.ParseTreeListener

	// EnterSegmentheader is called when entering the segmentheader production.
	EnterSegmentheader(c *SegmentheaderContext)

	// EnterSourceport is called when entering the sourceport production.
	EnterSourceport(c *SourceportContext)

	// EnterDestinationport is called when entering the destinationport production.
	EnterDestinationport(c *DestinationportContext)

	// EnterSequencenumber is called when entering the sequencenumber production.
	EnterSequencenumber(c *SequencenumberContext)

	// EnterAcknumber is called when entering the acknumber production.
	EnterAcknumber(c *AcknumberContext)

	// EnterFlags is called when entering the flags production.
	EnterFlags(c *FlagsContext)

	// EnterWindowsize is called when entering the windowsize production.
	EnterWindowsize(c *WindowsizeContext)

	// EnterChecksum is called when entering the checksum production.
	EnterChecksum(c *ChecksumContext)

	// EnterUrgent is called when entering the urgent production.
	EnterUrgent(c *UrgentContext)

	// EnterDword_ is called when entering the dword_ production.
	EnterDword_(c *Dword_Context)

	// EnterWord_ is called when entering the word_ production.
	EnterWord_(c *Word_Context)

	// EnterByte_ is called when entering the byte_ production.
	EnterByte_(c *Byte_Context)

	// ExitSegmentheader is called when exiting the segmentheader production.
	ExitSegmentheader(c *SegmentheaderContext)

	// ExitSourceport is called when exiting the sourceport production.
	ExitSourceport(c *SourceportContext)

	// ExitDestinationport is called when exiting the destinationport production.
	ExitDestinationport(c *DestinationportContext)

	// ExitSequencenumber is called when exiting the sequencenumber production.
	ExitSequencenumber(c *SequencenumberContext)

	// ExitAcknumber is called when exiting the acknumber production.
	ExitAcknumber(c *AcknumberContext)

	// ExitFlags is called when exiting the flags production.
	ExitFlags(c *FlagsContext)

	// ExitWindowsize is called when exiting the windowsize production.
	ExitWindowsize(c *WindowsizeContext)

	// ExitChecksum is called when exiting the checksum production.
	ExitChecksum(c *ChecksumContext)

	// ExitUrgent is called when exiting the urgent production.
	ExitUrgent(c *UrgentContext)

	// ExitDword_ is called when exiting the dword_ production.
	ExitDword_(c *Dword_Context)

	// ExitWord_ is called when exiting the word_ production.
	ExitWord_(c *Word_Context)

	// ExitByte_ is called when exiting the byte_ production.
	ExitByte_(c *Byte_Context)
}
