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
	"math/rand"
	"time"
)

type memoryStore struct {
	models []urlModel
}

func (s *memoryStore) list() []urlModel {
	s.delay()
	return s.models
}

func (s *memoryStore) find(id modelID) (urlModel, error) {
	s.delay()
	for _, model := range s.models {
		time.Sleep(10 * time.Millisecond)
		if model.ID == id {
			return model, nil
		}
	}

	return urlModel{}, errorModelNotFound
}

func (s *memoryStore) add(model urlModel) {
	s.delay()
	s.models = append(s.models, model)
}

func (s *memoryStore) delay() {
	amount := time.Duration(((rand.Int() % 10) * len(s.models))) * time.Millisecond
	time.Sleep(amount)
}
