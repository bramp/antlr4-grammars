// Generated from memcached_protocol.g4 by ANTLR 4.7.

package memcached_protocol // memcached_protocol
import "github.com/antlr/antlr4/runtime/Go/antlr"

type Basememcached_protocolVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Basememcached_protocolVisitor) VisitCommand_line(ctx *Command_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStorage_command(ctx *Storage_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStorage_command_name(ctx *Storage_command_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitRetrieval_command(ctx *Retrieval_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitDelete_command(ctx *Delete_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitIncrement_command(ctx *Increment_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitDecrement_command(ctx *Decrement_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStatistics_command(ctx *Statistics_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStatistics_option(ctx *Statistics_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitFlush_command(ctx *Flush_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitVersion_command(ctx *Version_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitVerbosity_command(ctx *Verbosity_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitQuit_command(ctx *Quit_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStorage_response(ctx *Storage_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitRetrieval_response(ctx *Retrieval_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitDeletion_response(ctx *Deletion_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitIncr_or_decr_response(ctx *Incr_or_decr_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStatistics_response(ctx *Statistics_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitError_response(ctx *Error_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitGeneral_statistic(ctx *General_statisticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitSize_statistic(ctx *Size_statisticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitGeneral_error(ctx *General_errorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitClient_error_message(ctx *Client_error_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitServer_error_message(ctx *Server_error_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitEnd(ctx *EndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitNoreply(ctx *NoreplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitFlags(ctx *FlagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitExptime(ctx *ExptimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitBytes(ctx *BytesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitCas_unique(ctx *Cas_uniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitTime(ctx *TimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitDelay(ctx *DelayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitVerbosity_level(ctx *Verbosity_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStatistic_name(ctx *Statistic_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitStatistic_value(ctx *Statistic_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitSize(ctx *SizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basememcached_protocolVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}
