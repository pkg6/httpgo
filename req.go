package httpgo

import (
	"net/http"
	"regexp"
	"strings"
)

var (
	CompileIsScript = []string{
		"curl",
		"wget",
		"collectd",
		"python",
		"urllib",
		"java",
		"jakarta",
		"httpclient",
		"phpcrawl",
		"libwww",
		"perl",
		"go-http",
		"okhttp",
		"lua-resty",
		"winhttp",
		"awesomium",
	}
)

func IsAjax(req *http.Request) bool {
	return req.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func IsScript(req *http.Request) bool {
	return regexp.
		MustCompile(strings.Join(CompileIsScript, "|")).
		MatchString(strings.ToLower(req.Header.Get("User-Agent")))
}

func IsSSL(req *http.Request) bool {
	return strings.EqualFold(req.URL.Scheme, "https") || req.TLS != nil
}
