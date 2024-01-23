package httpgo

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// OKXMLResponse writes v into w with http.StatusOK.
func OKXMLResponse(w http.ResponseWriter, v any) error {
	return XMLResponse(w, http.StatusOK, v)
}

func XMLResponse(w http.ResponseWriter, code int, v any) error {
	return XMLResponseWithCode(w, code, wrapXmlBaseResponse(v))
}

// XMLResponseWithCode writes v as xml string into w with code.
func XMLResponseWithCode(w http.ResponseWriter, code int, v any) error {
	return doWriteXml(w, code, v)
}

func doWriteXml(w http.ResponseWriter, code int, v any) error {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal xml failed, error: %w", err)
	}
	w.Header().Set("Content-Type", ResponseContentTypeXML)
	w.WriteHeader(code)
	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
	return nil
}
