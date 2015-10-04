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

// Package api provides a backing api for the shrturl application
package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benjic/shrturl/urls"
	"github.com/gorilla/mux"
)

// Version identifies the api version
const Version = "v1"

// ShrtURLAPI is a handler for shrturl API calls
type ShrtURLAPI struct{ router *mux.Router }

// Register attaches a api to the given router.
func Register(router *mux.Router) (api *ShrtURLAPI) {
	api = &ShrtURLAPI{}
	api.router = router.PathPrefix(fmt.Sprintf("/%s", Version)).Subrouter()

	urls.Register(api.router, api.handleFunc)

	return
}

func (api *ShrtURLAPI) handleFunc(f http.HandlerFunc) http.HandlerFunc {
	return handleLog(handleError(f))
}

func handleError(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {

				log.Printf("ERROR: %s", r)
				http.Error(w, "Server Error", http.StatusInternalServerError)
			}
		}()

		f(w, r)
	}
}
