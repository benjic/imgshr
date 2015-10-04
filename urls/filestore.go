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
	"bufio"
	"encoding/json"
	"errors"
)

// ErrorModelNotFound indicates that a model was not found
var ErrorModelNotFound = errors.New("Model was not found")

// A URLModel represents a url resource
type URLModel struct {
	ID  string
	URL string
}

type fileStore struct{ file *bufio.ReadWriter }

func (s *fileStore) list() (models []URLModel) {
	dec := json.NewDecoder(s.file)
	dec.Decode(&models)

	return
}

func (s *fileStore) find(id string) (URLModel, error) {
	var models []URLModel

	dec := json.NewDecoder(s.file)
	dec.Decode(&models)

	for _, model := range models {
		if model.ID == id {
			return model, nil
		}
	}
	return URLModel{}, ErrorModelNotFound
}

func (s *fileStore) add(model URLModel) error {
	var models []URLModel

	dec := json.NewDecoder(s.file)
	dec.Decode(&models)

	models = append(models, model)

	enc := json.NewEncoder(s.file)
	enc.Encode(models)

	s.file.Flush()

	return nil
}
