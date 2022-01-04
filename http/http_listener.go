// Code generated from http.g4 by ANTLR 4.9.3. DO NOT EDIT.

package http // http
import "github.com/antlr/antlr4/runtime/Go/antlr"

// httpListener is a complete listener for a parse tree produced by httpParser.
type httpListener interface {
	antlr.ParseTreeListener

	// EnterHttp_message is called when entering the http_message production.
	EnterHttp_message(c *Http_messageContext)

	// EnterStart_line is called when entering the start_line production.
	EnterStart_line(c *Start_lineContext)

	// EnterRequest_line is called when entering the request_line production.
	EnterRequest_line(c *Request_lineContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterRequest_target is called when entering the request_target production.
	EnterRequest_target(c *Request_targetContext)

	// EnterOrigin_form is called when entering the origin_form production.
	EnterOrigin_form(c *Origin_formContext)

	// EnterAbsolute_path is called when entering the absolute_path production.
	EnterAbsolute_path(c *Absolute_pathContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterHttp_version is called when entering the http_version production.
	EnterHttp_version(c *Http_versionContext)

	// EnterHttp_name is called when entering the http_name production.
	EnterHttp_name(c *Http_nameContext)

	// EnterHeader_field is called when entering the header_field production.
	EnterHeader_field(c *Header_fieldContext)

	// EnterField_name is called when entering the field_name production.
	EnterField_name(c *Field_nameContext)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterField_value is called when entering the field_value production.
	EnterField_value(c *Field_valueContext)

	// EnterField_content is called when entering the field_content production.
	EnterField_content(c *Field_contentContext)

	// EnterField_vchar is called when entering the field_vchar production.
	EnterField_vchar(c *Field_vcharContext)

	// EnterObs_text is called when entering the obs_text production.
	EnterObs_text(c *Obs_textContext)

	// EnterObs_fold is called when entering the obs_fold production.
	EnterObs_fold(c *Obs_foldContext)

	// EnterPchar is called when entering the pchar production.
	EnterPchar(c *PcharContext)

	// EnterUnreserved is called when entering the unreserved production.
	EnterUnreserved(c *UnreservedContext)

	// EnterSub_delims is called when entering the sub_delims production.
	EnterSub_delims(c *Sub_delimsContext)

	// EnterTchar is called when entering the tchar production.
	EnterTchar(c *TcharContext)

	// EnterVCHAR is called when entering the vCHAR production.
	EnterVCHAR(c *VCHARContext)

	// ExitHttp_message is called when exiting the http_message production.
	ExitHttp_message(c *Http_messageContext)

	// ExitStart_line is called when exiting the start_line production.
	ExitStart_line(c *Start_lineContext)

	// ExitRequest_line is called when exiting the request_line production.
	ExitRequest_line(c *Request_lineContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitRequest_target is called when exiting the request_target production.
	ExitRequest_target(c *Request_targetContext)

	// ExitOrigin_form is called when exiting the origin_form production.
	ExitOrigin_form(c *Origin_formContext)

	// ExitAbsolute_path is called when exiting the absolute_path production.
	ExitAbsolute_path(c *Absolute_pathContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitHttp_version is called when exiting the http_version production.
	ExitHttp_version(c *Http_versionContext)

	// ExitHttp_name is called when exiting the http_name production.
	ExitHttp_name(c *Http_nameContext)

	// ExitHeader_field is called when exiting the header_field production.
	ExitHeader_field(c *Header_fieldContext)

	// ExitField_name is called when exiting the field_name production.
	ExitField_name(c *Field_nameContext)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitField_value is called when exiting the field_value production.
	ExitField_value(c *Field_valueContext)

	// ExitField_content is called when exiting the field_content production.
	ExitField_content(c *Field_contentContext)

	// ExitField_vchar is called when exiting the field_vchar production.
	ExitField_vchar(c *Field_vcharContext)

	// ExitObs_text is called when exiting the obs_text production.
	ExitObs_text(c *Obs_textContext)

	// ExitObs_fold is called when exiting the obs_fold production.
	ExitObs_fold(c *Obs_foldContext)

	// ExitPchar is called when exiting the pchar production.
	ExitPchar(c *PcharContext)

	// ExitUnreserved is called when exiting the unreserved production.
	ExitUnreserved(c *UnreservedContext)

	// ExitSub_delims is called when exiting the sub_delims production.
	ExitSub_delims(c *Sub_delimsContext)

	// ExitTchar is called when exiting the tchar production.
	ExitTchar(c *TcharContext)

	// ExitVCHAR is called when exiting the vCHAR production.
	ExitVCHAR(c *VCHARContext)
}
