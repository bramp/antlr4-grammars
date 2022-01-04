// Code generated from tcp.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tcp // tcp
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetcpListener is a complete listener for a parse tree produced by tcpParser.
type BasetcpListener struct{}

var _ tcpListener = &BasetcpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetcpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetcpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetcpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetcpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSegmentheader is called when production segmentheader is entered.
func (s *BasetcpListener) EnterSegmentheader(ctx *SegmentheaderContext) {}

// ExitSegmentheader is called when production segmentheader is exited.
func (s *BasetcpListener) ExitSegmentheader(ctx *SegmentheaderContext) {}

// EnterSourceport is called when production sourceport is entered.
func (s *BasetcpListener) EnterSourceport(ctx *SourceportContext) {}

// ExitSourceport is called when production sourceport is exited.
func (s *BasetcpListener) ExitSourceport(ctx *SourceportContext) {}

// EnterDestinationport is called when production destinationport is entered.
func (s *BasetcpListener) EnterDestinationport(ctx *DestinationportContext) {}

// ExitDestinationport is called when production destinationport is exited.
func (s *BasetcpListener) ExitDestinationport(ctx *DestinationportContext) {}

// EnterSequencenumber is called when production sequencenumber is entered.
func (s *BasetcpListener) EnterSequencenumber(ctx *SequencenumberContext) {}

// ExitSequencenumber is called when production sequencenumber is exited.
func (s *BasetcpListener) ExitSequencenumber(ctx *SequencenumberContext) {}

// EnterAcknumber is called when production acknumber is entered.
func (s *BasetcpListener) EnterAcknumber(ctx *AcknumberContext) {}

// ExitAcknumber is called when production acknumber is exited.
func (s *BasetcpListener) ExitAcknumber(ctx *AcknumberContext) {}

// EnterFlags is called when production flags is entered.
func (s *BasetcpListener) EnterFlags(ctx *FlagsContext) {}

// ExitFlags is called when production flags is exited.
func (s *BasetcpListener) ExitFlags(ctx *FlagsContext) {}

// EnterWindowsize is called when production windowsize is entered.
func (s *BasetcpListener) EnterWindowsize(ctx *WindowsizeContext) {}

// ExitWindowsize is called when production windowsize is exited.
func (s *BasetcpListener) ExitWindowsize(ctx *WindowsizeContext) {}

// EnterChecksum is called when production checksum is entered.
func (s *BasetcpListener) EnterChecksum(ctx *ChecksumContext) {}

// ExitChecksum is called when production checksum is exited.
func (s *BasetcpListener) ExitChecksum(ctx *ChecksumContext) {}

// EnterUrgent is called when production urgent is entered.
func (s *BasetcpListener) EnterUrgent(ctx *UrgentContext) {}

// ExitUrgent is called when production urgent is exited.
func (s *BasetcpListener) ExitUrgent(ctx *UrgentContext) {}

// EnterDword_ is called when production dword_ is entered.
func (s *BasetcpListener) EnterDword_(ctx *Dword_Context) {}

// ExitDword_ is called when production dword_ is exited.
func (s *BasetcpListener) ExitDword_(ctx *Dword_Context) {}

// EnterWord_ is called when production word_ is entered.
func (s *BasetcpListener) EnterWord_(ctx *Word_Context) {}

// ExitWord_ is called when production word_ is exited.
func (s *BasetcpListener) ExitWord_(ctx *Word_Context) {}

// EnterByte_ is called when production byte_ is entered.
func (s *BasetcpListener) EnterByte_(ctx *Byte_Context) {}

// ExitByte_ is called when production byte_ is exited.
func (s *BasetcpListener) ExitByte_(ctx *Byte_Context) {}
