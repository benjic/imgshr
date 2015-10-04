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
	"errors"
	"math/rand"
)

var errorModelNotFound = errors.New("Model with given ID not found")

type modelID string

type urlModel struct {
	ID  modelID
	URL string
}

type store interface {
	list() []urlModel
	find(modelID) (urlModel, error)
	add(urlModel)
}

func createModelID() modelID {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	newID := make([]byte, 8)
	for i := range newID {
		newID[i] = chars[rand.Int63()%int64(len(chars))]
	}

	return modelID(newID)
}
