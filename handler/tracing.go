package handler

import (
	"net/http"
	"fmt"
	"io"
)

func Tracing(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range req.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
