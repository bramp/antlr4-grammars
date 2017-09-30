// Generated from clf.g4 by ANTLR 4.7.

package clf // clf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// clfListener is a complete listener for a parse tree produced by clfParser.
type clfListener interface {
	antlr.ParseTreeListener

	// EnterLog is called when entering the log production.
	EnterLog(c *LogContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterLogname is called when entering the logname production.
	EnterLogname(c *LognameContext)

	// EnterUsername is called when entering the username production.
	EnterUsername(c *UsernameContext)

	// EnterDatetimetz is called when entering the datetimetz production.
	EnterDatetimetz(c *DatetimetzContext)

	// EnterReferer is called when entering the referer production.
	EnterReferer(c *RefererContext)

	// EnterRequest is called when entering the request production.
	EnterRequest(c *RequestContext)

	// EnterUseragent is called when entering the useragent production.
	EnterUseragent(c *UseragentContext)

	// EnterStatuscode is called when entering the statuscode production.
	EnterStatuscode(c *StatuscodeContext)

	// EnterBytes is called when entering the bytes production.
	EnterBytes(c *BytesContext)

	// ExitLog is called when exiting the log production.
	ExitLog(c *LogContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitLogname is called when exiting the logname production.
	ExitLogname(c *LognameContext)

	// ExitUsername is called when exiting the username production.
	ExitUsername(c *UsernameContext)

	// ExitDatetimetz is called when exiting the datetimetz production.
	ExitDatetimetz(c *DatetimetzContext)

	// ExitReferer is called when exiting the referer production.
	ExitReferer(c *RefererContext)

	// ExitRequest is called when exiting the request production.
	ExitRequest(c *RequestContext)

	// ExitUseragent is called when exiting the useragent production.
	ExitUseragent(c *UseragentContext)

	// ExitStatuscode is called when exiting the statuscode production.
	ExitStatuscode(c *StatuscodeContext)

	// ExitBytes is called when exiting the bytes production.
	ExitBytes(c *BytesContext)
}
