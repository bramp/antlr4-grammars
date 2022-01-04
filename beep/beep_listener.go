// Code generated from beep.g4 by ANTLR 4.9.3. DO NOT EDIT.

package beep // beep
import "github.com/antlr/antlr4/runtime/Go/antlr"

// beepListener is a complete listener for a parse tree produced by beepParser.
type beepListener interface {
	antlr.ParseTreeListener

	// EnterFrame is called when entering the frame production.
	EnterFrame(c *FrameContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterMsg is called when entering the msg production.
	EnterMsg(c *MsgContext)

	// EnterRpy is called when entering the rpy production.
	EnterRpy(c *RpyContext)

	// EnterAns is called when entering the ans production.
	EnterAns(c *AnsContext)

	// EnterErr is called when entering the err production.
	EnterErr(c *ErrContext)

	// EnterNul is called when entering the nul production.
	EnterNul(c *NulContext)

	// EnterCommon is called when entering the common production.
	EnterCommon(c *CommonContext)

	// EnterChannel is called when entering the channel production.
	EnterChannel(c *ChannelContext)

	// EnterMsgno is called when entering the msgno production.
	EnterMsgno(c *MsgnoContext)

	// EnterMore is called when entering the more production.
	EnterMore(c *MoreContext)

	// EnterSeqno is called when entering the seqno production.
	EnterSeqno(c *SeqnoContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterAnsno is called when entering the ansno production.
	EnterAnsno(c *AnsnoContext)

	// EnterPayload_trailer is called when entering the payload_trailer production.
	EnterPayload_trailer(c *Payload_trailerContext)

	// ExitFrame is called when exiting the frame production.
	ExitFrame(c *FrameContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitMsg is called when exiting the msg production.
	ExitMsg(c *MsgContext)

	// ExitRpy is called when exiting the rpy production.
	ExitRpy(c *RpyContext)

	// ExitAns is called when exiting the ans production.
	ExitAns(c *AnsContext)

	// ExitErr is called when exiting the err production.
	ExitErr(c *ErrContext)

	// ExitNul is called when exiting the nul production.
	ExitNul(c *NulContext)

	// ExitCommon is called when exiting the common production.
	ExitCommon(c *CommonContext)

	// ExitChannel is called when exiting the channel production.
	ExitChannel(c *ChannelContext)

	// ExitMsgno is called when exiting the msgno production.
	ExitMsgno(c *MsgnoContext)

	// ExitMore is called when exiting the more production.
	ExitMore(c *MoreContext)

	// ExitSeqno is called when exiting the seqno production.
	ExitSeqno(c *SeqnoContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitAnsno is called when exiting the ansno production.
	ExitAnsno(c *AnsnoContext)

	// ExitPayload_trailer is called when exiting the payload_trailer production.
	ExitPayload_trailer(c *Payload_trailerContext)
}
