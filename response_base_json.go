package httpgo

import (
	"errors"
	"fmt"
	"github.com/pkg6/httpgo/json"
	"net/http"
)

// OKJSONResponse writes v into w with http.StatusOK.
func OKJSONResponse(w http.ResponseWriter, v any) error {
	return JSONResponse(w, http.StatusOK, v)
}

func JSONResponse(w http.ResponseWriter, code int, v any) error {
	return JSONResponseWithCode(w, code, wrapBaseResponse(v))
}

// JSONResponseWithCode writes v as json string into w with code.
func JSONResponseWithCode(w http.ResponseWriter, code int, v any) error {
	return doWriteJson(w, code, v)
}

func doWriteJson(w http.ResponseWriter, code int, v any) error {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal json failed, error: %w", err)
	}
	w.Header().Set("Content-Type", ResponseContentTypeJSON)
	w.WriteHeader(code)
	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if !errors.Is(err, http.ErrHandlerTimeout) {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
	return nil
}
