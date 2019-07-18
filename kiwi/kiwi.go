/*
 * kiwi.
 *
 * postgreSQL protocol interaction library.
 */

package kiwi

var kiwi_fe_type = map[string]byte {
	"KIWI_FE_TERMINATE" :         byte('X'),
	"KIWI_FE_PASSWORD_MESSAGE" :  byte('p'),
	"KIWI_FE_QUERY" :             byte('Q'),
	"KIWI_FE_FUNCTION_CALL" :     byte('F'),
	"KIWI_FE_PARSE" :             byte('P'),
	"KIWI_FE_BIND" :              byte('B'),
	"KIWI_FE_DESCRIBE" :          byte('D'),
	"KIWI_FE_EXECUTE" :           byte('E'),
	"KIWI_FE_SYNC" :              byte('S'),
	"KIWI_FE_CLOSE" :             byte('C'),
	"KIWI_FE_COPY_DATA" :         byte('d'),
	"KIWI_FE_COPY_DONE" :         byte('c'),
	"KIWI_FE_COPY_FAIL" :         byte('f'),
}

var kiwi_fe_type_desc = map[string]string {
	"KIWI_FE_TERMINATE" :         "Terminate",
	"KIWI_FE_PASSWORD_MESSAGE" :  "PasswordMessage",
	"KIWI_FE_QUERY" :             "Query",
	"KIWI_FE_FUNCTION_CALL" :     "FunctionCall",
	"KIWI_FE_PARSE" :             "Parse",
	"KIWI_FE_BIND" :              "Bind",
	"KIWI_FE_DESCRIBE" :          "Describe",
	"KIWI_FE_EXECUTE" :           "Execute",
	"KIWI_FE_SYNC" :              "Sync",
	"KIWI_FE_CLOSE" :             "Close",
	"KIWI_FE_COPY_DATA" :         "CopyData",
	"KIWI_FE_COPY_DONE" :         "CopyDone",
	"KIWI_FE_COPY_FAIL" :         "CopyFail",
}


var kiwi_be_type = map[string]byte{
	"KIWI_BE_AUTHENTICATION"             : byte('R'),
	"KIWI_BE_BACKEND_KEY_DATA"           : byte('K'),
	"KIWI_BE_PARSE_COMPLETE"             : byte('1'),
	"KIWI_BE_BIND_COMPLETE"              : byte('2'),
	"KIWI_BE_CLOSE_COMPLETE"             : byte('3'),
	"KIWI_BE_COMMAND_COMPLETE"           : byte('C'),
	"KIWI_BE_COPY_IN_RESPONSE"           : byte('G'),
	"KIWI_BE_COPY_OUT_RESPONSE"          : byte('H'),
	"KIWI_BE_COPY_BOTH_RESPONSE"         : byte('W'),
	"KIWI_BE_COPY_DATA"                  : byte('d'),
	"KIWI_BE_COPY_DONE"                  : byte('c'),
	"KIWI_BE_COPY_FAIL"                  : byte('f'),
	"KIWI_BE_DATA_ROW"                   : byte('D'),
	"KIWI_BE_EMPTY_QUERY_RESPONSE"       : byte('I'),
	"KIWI_BE_ERROR_RESPONSE"             : byte('E'),
	"KIWI_BE_FUNCTION_CALL_RESPONSE"     : byte('V'),
	"KIWI_BE_NEGOTIATE_PROTOCOL_VERSION" : byte('v'),
	"KIWI_BE_NO_DATA"                    : byte('n'),
	"KIWI_BE_NOTICE_RESPONSE"            : byte('N'),
	"KIWI_BE_NOTIFICATION_RESPONSE"      : byte('A'),
	"KIWI_BE_PARAMETER_DESCRIPTION"      : byte('t'),
	"KIWI_BE_PARAMETER_STATUS"           : byte('S'),
	"KIWI_BE_PORTAL_SUSPENDED"           : byte('s'),
	"KIWI_BE_READY_FOR_QUERY"            : byte('Z'),
	"KIWI_BE_ROW_DESCRIPTION"            : byte('T'),
}

var kiwi_be_type_desc = map[string]string{
	"KIWI_BE_AUTHENTICATION"             : "Authentication",
	"KIWI_BE_BACKEND_KEY_DATA"           : "BackendKeyData",
	"KIWI_BE_PARSE_COMPLETE"             : "ParseComplete",
	"KIWI_BE_BIND_COMPLETE"              : "BindComplete",
	"KIWI_BE_CLOSE_COMPLETE"             : "CloseComplete",
	"KIWI_BE_COMMAND_COMPLETE"           : "CommandComplete",
	"KIWI_BE_COPY_IN_RESPONSE"           : "CopyInResponse",
	"KIWI_BE_COPY_OUT_RESPONSE"          : "CopyOutResponse",
	"KIWI_BE_COPY_BOTH_RESPONSE"         : "CopyBothResponse",
	"KIWI_BE_COPY_DATA"                  : "CopyData",
	"KIWI_BE_COPY_DONE"                  : "CopyDone",
	"KIWI_BE_COPY_FAIL"                  : "CopyFail",
	"KIWI_BE_DATA_ROW"                   : "DataRow",
	"KIWI_BE_EMPTY_QUERY_RESPONSE"       : "EmptyQueryResponse",
	"KIWI_BE_ERROR_RESPONSE"             : "ErrorResponse",
	"KIWI_BE_FUNCTION_CALL_RESPONSE"     : "FunctionCallResponse",
	"KIWI_BE_NEGOTIATE_PROTOCOL_VERSION" : "NegotiateProtocolVersion",
	"KIWI_BE_NO_DATA"                    : "NoData",
	"KIWI_BE_NOTICE_RESPONSE"            : "NoticeResponse",
	"KIWI_BE_NOTIFICATION_RESPONSE"      : "NotificationResponse",
	"KIWI_BE_PARAMETER_DESCRIPTION"      : "ParameterDescription",
	"KIWI_BE_PARAMETER_STATUS"           : "ParameterStatus",
	"KIWI_BE_PORTAL_SUSPENDED"           : "PortalSuspended",
	"KIWI_BE_READY_FOR_QUERY"            : "ReadyForQuery",
	"KIWI_BE_ROW_DESCRIPTION"            : "RowDescription",
}

type kiwi_header struct {
	h_type uint8
    len uint32
}

func mapkey(m map[string]byte, value byte) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	if len(key) == 0 {
		key = "Unknown"
	}
	return
}

