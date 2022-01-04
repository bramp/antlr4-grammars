// Code generated from http.g4 by ANTLR 4.9.3. DO NOT EDIT.

package http // http
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasehttpListener is a complete listener for a parse tree produced by httpParser.
type BasehttpListener struct{}

var _ httpListener = &BasehttpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasehttpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasehttpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasehttpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasehttpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHttp_message is called when production http_message is entered.
func (s *BasehttpListener) EnterHttp_message(ctx *Http_messageContext) {}

// ExitHttp_message is called when production http_message is exited.
func (s *BasehttpListener) ExitHttp_message(ctx *Http_messageContext) {}

// EnterStart_line is called when production start_line is entered.
func (s *BasehttpListener) EnterStart_line(ctx *Start_lineContext) {}

// ExitStart_line is called when production start_line is exited.
func (s *BasehttpListener) ExitStart_line(ctx *Start_lineContext) {}

// EnterRequest_line is called when production request_line is entered.
func (s *BasehttpListener) EnterRequest_line(ctx *Request_lineContext) {}

// ExitRequest_line is called when production request_line is exited.
func (s *BasehttpListener) ExitRequest_line(ctx *Request_lineContext) {}

// EnterMethod is called when production method is entered.
func (s *BasehttpListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BasehttpListener) ExitMethod(ctx *MethodContext) {}

// EnterRequest_target is called when production request_target is entered.
func (s *BasehttpListener) EnterRequest_target(ctx *Request_targetContext) {}

// ExitRequest_target is called when production request_target is exited.
func (s *BasehttpListener) ExitRequest_target(ctx *Request_targetContext) {}

// EnterOrigin_form is called when production origin_form is entered.
func (s *BasehttpListener) EnterOrigin_form(ctx *Origin_formContext) {}

// ExitOrigin_form is called when production origin_form is exited.
func (s *BasehttpListener) ExitOrigin_form(ctx *Origin_formContext) {}

// EnterAbsolute_path is called when production absolute_path is entered.
func (s *BasehttpListener) EnterAbsolute_path(ctx *Absolute_pathContext) {}

// ExitAbsolute_path is called when production absolute_path is exited.
func (s *BasehttpListener) ExitAbsolute_path(ctx *Absolute_pathContext) {}

// EnterSegment is called when production segment is entered.
func (s *BasehttpListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BasehttpListener) ExitSegment(ctx *SegmentContext) {}

// EnterQuery is called when production query is entered.
func (s *BasehttpListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasehttpListener) ExitQuery(ctx *QueryContext) {}

// EnterHttp_version is called when production http_version is entered.
func (s *BasehttpListener) EnterHttp_version(ctx *Http_versionContext) {}

// ExitHttp_version is called when production http_version is exited.
func (s *BasehttpListener) ExitHttp_version(ctx *Http_versionContext) {}

// EnterHttp_name is called when production http_name is entered.
func (s *BasehttpListener) EnterHttp_name(ctx *Http_nameContext) {}

// ExitHttp_name is called when production http_name is exited.
func (s *BasehttpListener) ExitHttp_name(ctx *Http_nameContext) {}

// EnterHeader_field is called when production header_field is entered.
func (s *BasehttpListener) EnterHeader_field(ctx *Header_fieldContext) {}

// ExitHeader_field is called when production header_field is exited.
func (s *BasehttpListener) ExitHeader_field(ctx *Header_fieldContext) {}

// EnterField_name is called when production field_name is entered.
func (s *BasehttpListener) EnterField_name(ctx *Field_nameContext) {}

// ExitField_name is called when production field_name is exited.
func (s *BasehttpListener) ExitField_name(ctx *Field_nameContext) {}

// EnterToken is called when production token is entered.
func (s *BasehttpListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BasehttpListener) ExitToken(ctx *TokenContext) {}

// EnterField_value is called when production field_value is entered.
func (s *BasehttpListener) EnterField_value(ctx *Field_valueContext) {}

// ExitField_value is called when production field_value is exited.
func (s *BasehttpListener) ExitField_value(ctx *Field_valueContext) {}

// EnterField_content is called when production field_content is entered.
func (s *BasehttpListener) EnterField_content(ctx *Field_contentContext) {}

// ExitField_content is called when production field_content is exited.
func (s *BasehttpListener) ExitField_content(ctx *Field_contentContext) {}

// EnterField_vchar is called when production field_vchar is entered.
func (s *BasehttpListener) EnterField_vchar(ctx *Field_vcharContext) {}

// ExitField_vchar is called when production field_vchar is exited.
func (s *BasehttpListener) ExitField_vchar(ctx *Field_vcharContext) {}

// EnterObs_text is called when production obs_text is entered.
func (s *BasehttpListener) EnterObs_text(ctx *Obs_textContext) {}

// ExitObs_text is called when production obs_text is exited.
func (s *BasehttpListener) ExitObs_text(ctx *Obs_textContext) {}

// EnterObs_fold is called when production obs_fold is entered.
func (s *BasehttpListener) EnterObs_fold(ctx *Obs_foldContext) {}

// ExitObs_fold is called when production obs_fold is exited.
func (s *BasehttpListener) ExitObs_fold(ctx *Obs_foldContext) {}

// EnterPchar is called when production pchar is entered.
func (s *BasehttpListener) EnterPchar(ctx *PcharContext) {}

// ExitPchar is called when production pchar is exited.
func (s *BasehttpListener) ExitPchar(ctx *PcharContext) {}

// EnterUnreserved is called when production unreserved is entered.
func (s *BasehttpListener) EnterUnreserved(ctx *UnreservedContext) {}

// ExitUnreserved is called when production unreserved is exited.
func (s *BasehttpListener) ExitUnreserved(ctx *UnreservedContext) {}

// EnterSub_delims is called when production sub_delims is entered.
func (s *BasehttpListener) EnterSub_delims(ctx *Sub_delimsContext) {}

// ExitSub_delims is called when production sub_delims is exited.
func (s *BasehttpListener) ExitSub_delims(ctx *Sub_delimsContext) {}

// EnterTchar is called when production tchar is entered.
func (s *BasehttpListener) EnterTchar(ctx *TcharContext) {}

// ExitTchar is called when production tchar is exited.
func (s *BasehttpListener) ExitTchar(ctx *TcharContext) {}

// EnterVCHAR is called when production vCHAR is entered.
func (s *BasehttpListener) EnterVCHAR(ctx *VCHARContext) {}

// ExitVCHAR is called when production vCHAR is exited.
func (s *BasehttpListener) ExitVCHAR(ctx *VCHARContext) {}
