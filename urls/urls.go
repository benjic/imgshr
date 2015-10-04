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

package urls

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type urlHandler struct {
	store store
}

func newURLHandler() *urlHandler {
	return &urlHandler{&memoryStore{models: make([]urlModel, 0)}}
}

// Register adds the urls resource to the given api instance
func Register(r *mux.Router, register func(http.HandlerFunc) http.HandlerFunc) (err error) {

	urls := newURLHandler()
	r.Methods("GET").Subrouter().HandleFunc("/urls", register(urls.list()))
	r.Methods("GET").Subrouter().HandleFunc("/urls/{id:[a-zA-Z0-9]+}", register(urls.item()))
	r.Methods("POST").Subrouter().HandleFunc("/urls", register(urls.add()))

	return
}

func (h *urlHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(h.store.list())
	}
}

func (h *urlHandler) item() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := modelID(mux.Vars(r)["id"])
		model, err := h.store.find(id)

		if err == errorModelNotFound {
			http.Error(w, "Resource Not Found", 404)
			return
		}

		if strings.Contains(r.Header.Get("Accept"), "text/html") {
			http.Redirect(w, r, model.URL, http.StatusTemporaryRedirect)
		} else {
			json.NewEncoder(w).Encode(model)
		}
	}
}

func (h *urlHandler) add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		model := &urlModel{}

		err := json.NewDecoder(r.Body).Decode(model)
		model.ID = createModelID()

		h.store.add(*model)
		if err != nil {
			fmt.Printf("Error: %s", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
	}
}
