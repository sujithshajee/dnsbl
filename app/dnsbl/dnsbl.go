package dnsbl

import "context"

type ResponseCode string

const (
	ResponseCode_SpamhausSBL_Data         = "127.0.0.2"
	ResponseCode_SpamhausSBLCSS_Data      = "127.0.0.3"
	ResponseCode_SpamhausCBL_Data         = "127.0.0.4"
	ResponseCode_SpamhausDROP_Data        = "127.0.0.10"
	ResponseCode_ISP_Maintained           = "127.0.0.10"
	ResponseCode_Spamhaus_Maintained      = "127.0.0.11"
	ResponseCode_Typing_Error             = "127.255.255.252"
	ResponseCode_Query_PublicOpenResolver = "127.255.255.254"
	ResponseCode_Excessive_queries        = "127.255.255.255"
	ResponseCode_Unknown                  = "0.0.0.0"
)

type Response struct {
	ResponseCodes []ResponseCode
}

func freshResponse() *Response {
	return &Response{
		ResponseCodes: []ResponseCode{},
	}
}

type DNSBL interface {
	Query(ctx context.Context, ip string) (*Response, error)
}
