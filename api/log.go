// Copyright 2015 Benjamin Campbell <benji@benjica.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"log"
	"net/http"
)

type logWriter struct {
	writer     http.ResponseWriter
	statusCode int
}

func newLogWriter(w http.ResponseWriter) *logWriter {
	return &logWriter{writer: w}
}

func (w *logWriter) Header() http.Header { return w.writer.Header() }
func (w *logWriter) Write(bytes []byte) (int, error) {
	return w.writer.Write(bytes)
}
func (w *logWriter) WriteHeader(code int) {
	w.statusCode = code
	w.writer.WriteHeader(code)
}

// Log wraps an api call in a logging handler
func handleLog(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logWriter := newLogWriter(w)
		f(logWriter, r)
		log.Printf("%s\t%d %s", r.Method, logWriter.statusCode, r.URL.Path)
	}
}
