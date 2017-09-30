// Generated from clf.g4 by ANTLR 4.7.

package clf // clf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseclfListener is a complete listener for a parse tree produced by clfParser.
type BaseclfListener struct{}

var _ clfListener = &BaseclfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseclfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseclfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseclfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseclfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLog is called when production log is entered.
func (s *BaseclfListener) EnterLog(ctx *LogContext) {}

// ExitLog is called when production log is exited.
func (s *BaseclfListener) ExitLog(ctx *LogContext) {}

// EnterLine is called when production line is entered.
func (s *BaseclfListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseclfListener) ExitLine(ctx *LineContext) {}

// EnterHost is called when production host is entered.
func (s *BaseclfListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseclfListener) ExitHost(ctx *HostContext) {}

// EnterLogname is called when production logname is entered.
func (s *BaseclfListener) EnterLogname(ctx *LognameContext) {}

// ExitLogname is called when production logname is exited.
func (s *BaseclfListener) ExitLogname(ctx *LognameContext) {}

// EnterUsername is called when production username is entered.
func (s *BaseclfListener) EnterUsername(ctx *UsernameContext) {}

// ExitUsername is called when production username is exited.
func (s *BaseclfListener) ExitUsername(ctx *UsernameContext) {}

// EnterDatetimetz is called when production datetimetz is entered.
func (s *BaseclfListener) EnterDatetimetz(ctx *DatetimetzContext) {}

// ExitDatetimetz is called when production datetimetz is exited.
func (s *BaseclfListener) ExitDatetimetz(ctx *DatetimetzContext) {}

// EnterReferer is called when production referer is entered.
func (s *BaseclfListener) EnterReferer(ctx *RefererContext) {}

// ExitReferer is called when production referer is exited.
func (s *BaseclfListener) ExitReferer(ctx *RefererContext) {}

// EnterRequest is called when production request is entered.
func (s *BaseclfListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BaseclfListener) ExitRequest(ctx *RequestContext) {}

// EnterUseragent is called when production useragent is entered.
func (s *BaseclfListener) EnterUseragent(ctx *UseragentContext) {}

// ExitUseragent is called when production useragent is exited.
func (s *BaseclfListener) ExitUseragent(ctx *UseragentContext) {}

// EnterStatuscode is called when production statuscode is entered.
func (s *BaseclfListener) EnterStatuscode(ctx *StatuscodeContext) {}

// ExitStatuscode is called when production statuscode is exited.
func (s *BaseclfListener) ExitStatuscode(ctx *StatuscodeContext) {}

// EnterBytes is called when production bytes is entered.
func (s *BaseclfListener) EnterBytes(ctx *BytesContext) {}

// ExitBytes is called when production bytes is exited.
func (s *BaseclfListener) ExitBytes(ctx *BytesContext) {}
