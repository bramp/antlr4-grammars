// Code generated from beep.g4 by ANTLR 4.9.3. DO NOT EDIT.

package beep // beep
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebeepListener is a complete listener for a parse tree produced by beepParser.
type BasebeepListener struct{}

var _ beepListener = &BasebeepListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebeepListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebeepListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebeepListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebeepListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFrame is called when production frame is entered.
func (s *BasebeepListener) EnterFrame(ctx *FrameContext) {}

// ExitFrame is called when production frame is exited.
func (s *BasebeepListener) ExitFrame(ctx *FrameContext) {}

// EnterData is called when production data is entered.
func (s *BasebeepListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BasebeepListener) ExitData(ctx *DataContext) {}

// EnterHeader is called when production header is entered.
func (s *BasebeepListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BasebeepListener) ExitHeader(ctx *HeaderContext) {}

// EnterMsg is called when production msg is entered.
func (s *BasebeepListener) EnterMsg(ctx *MsgContext) {}

// ExitMsg is called when production msg is exited.
func (s *BasebeepListener) ExitMsg(ctx *MsgContext) {}

// EnterRpy is called when production rpy is entered.
func (s *BasebeepListener) EnterRpy(ctx *RpyContext) {}

// ExitRpy is called when production rpy is exited.
func (s *BasebeepListener) ExitRpy(ctx *RpyContext) {}

// EnterAns is called when production ans is entered.
func (s *BasebeepListener) EnterAns(ctx *AnsContext) {}

// ExitAns is called when production ans is exited.
func (s *BasebeepListener) ExitAns(ctx *AnsContext) {}

// EnterErr is called when production err is entered.
func (s *BasebeepListener) EnterErr(ctx *ErrContext) {}

// ExitErr is called when production err is exited.
func (s *BasebeepListener) ExitErr(ctx *ErrContext) {}

// EnterNul is called when production nul is entered.
func (s *BasebeepListener) EnterNul(ctx *NulContext) {}

// ExitNul is called when production nul is exited.
func (s *BasebeepListener) ExitNul(ctx *NulContext) {}

// EnterCommon is called when production common is entered.
func (s *BasebeepListener) EnterCommon(ctx *CommonContext) {}

// ExitCommon is called when production common is exited.
func (s *BasebeepListener) ExitCommon(ctx *CommonContext) {}

// EnterChannel is called when production channel is entered.
func (s *BasebeepListener) EnterChannel(ctx *ChannelContext) {}

// ExitChannel is called when production channel is exited.
func (s *BasebeepListener) ExitChannel(ctx *ChannelContext) {}

// EnterMsgno is called when production msgno is entered.
func (s *BasebeepListener) EnterMsgno(ctx *MsgnoContext) {}

// ExitMsgno is called when production msgno is exited.
func (s *BasebeepListener) ExitMsgno(ctx *MsgnoContext) {}

// EnterMore is called when production more is entered.
func (s *BasebeepListener) EnterMore(ctx *MoreContext) {}

// ExitMore is called when production more is exited.
func (s *BasebeepListener) ExitMore(ctx *MoreContext) {}

// EnterSeqno is called when production seqno is entered.
func (s *BasebeepListener) EnterSeqno(ctx *SeqnoContext) {}

// ExitSeqno is called when production seqno is exited.
func (s *BasebeepListener) ExitSeqno(ctx *SeqnoContext) {}

// EnterSize is called when production size is entered.
func (s *BasebeepListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BasebeepListener) ExitSize(ctx *SizeContext) {}

// EnterAnsno is called when production ansno is entered.
func (s *BasebeepListener) EnterAnsno(ctx *AnsnoContext) {}

// ExitAnsno is called when production ansno is exited.
func (s *BasebeepListener) ExitAnsno(ctx *AnsnoContext) {}

// EnterPayload_trailer is called when production payload_trailer is entered.
func (s *BasebeepListener) EnterPayload_trailer(ctx *Payload_trailerContext) {}

// ExitPayload_trailer is called when production payload_trailer is exited.
func (s *BasebeepListener) ExitPayload_trailer(ctx *Payload_trailerContext) {}
