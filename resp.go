package httpgo

import (
	"encoding/xml"
	"github.com/pkg6/httpgo/errors"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	// BusinessCodeOK represents the business code for success.
	BusinessCodeOK = 0
	// BusinessMsgOK represents the business message for success.
	BusinessMsgOK = "ok"
	// BusinessCodeError represents the business code for error.
	BusinessCodeError = -1

	ResponseContentTypeJSON = "application/json"
	ResponseContentTypeXML  = "application/xml"
)

const (
	xmlVersion  = "1.0"
	xmlEncoding = "UTF-8"
)

type BaseResponseErr interface {
	error
	Code() int
}

// OKResponse writes HTTP 200 OK into w.
func OKResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	// Code represents the business code, not the http status code.
	Code int `json:"code" xml:"code"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgSuccess.
	Msg string `json:"msg" xml:"msg"`
	// Data represents the business data.
	Data T `json:"data,omitempty" xml:"data,omitempty"`
}

type baseXmlResponse[T any] struct {
	XMLName  xml.Name `xml:"xml"`
	Version  string   `xml:"version,attr"`
	Encoding string   `xml:"encoding,attr"`
	BaseResponse[T]
}

func wrapXmlBaseResponse(v any) baseXmlResponse[any] {
	base := wrapBaseResponse(v)
	return baseXmlResponse[any]{
		Version:      xmlVersion,
		Encoding:     xmlEncoding,
		BaseResponse: base,
	}
}

func wrapBaseResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]
	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case BaseResponseErr:
		resp.Code = data.Code()
		resp.Msg = data.Error()
	case *status.Status:
		resp.Code = int(data.Code())
		resp.Msg = data.Message()
	case error:
		resp.Code = BusinessCodeError
		resp.Msg = data.Error()
	default:
		resp.Code = BusinessCodeOK
		resp.Msg = BusinessMsgOK
		resp.Data = v
	}
	return resp
}
