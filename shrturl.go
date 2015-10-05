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
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/benjic/shrturl/api"
	"github.com/gorilla/mux"
)

func main() {
	// Register handlers
	r := mux.NewRouter()
	_, err := api.Register(r)

	// Determine if api is functional
	if err != nil {
		log.Fatalf("Unable to standup api service: %s", err)
		os.Exit(1)
	}

	// Start http server
	log.Print("shrturl service now running localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
