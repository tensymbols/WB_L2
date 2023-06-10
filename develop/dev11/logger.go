package main

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handler: handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
