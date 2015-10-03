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
	"fmt"
	"net/http"
)

type urlHandler struct{}

// Register adds the urls resource to the given api instance
func Register(register func(string, http.HandlerFunc)) (err error) {

	urls := &urlHandler{}
	register("/urls", urls.list())

	return
}

func (h *urlHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("You've made it!")
	}
}
