package main

import (
	"fmt"
	"net/http"

	"github.com/kosha/twilio-connector/pkg/app"
	"github.com/kosha/twilio-connector/pkg/logger"
)
var (
	log  = logger.New("app", "twilio-connector")
	port = 8015
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {

	a := app.App{}
	a.Initialize(log)
	
	log.Infof("Running twilio-connector on port %d", port)
	a.Run(fmt.Sprintf(":%d", port))
}
