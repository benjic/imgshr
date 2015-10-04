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

import "testing"

func TestMemoryStoreList(t *testing.T) {
	cases := []struct {
		models []urlModel
	}{
		{[]urlModel{
			urlModel{"abc", "http://url"},
			urlModel{"def", "http://url"},
		}},
	}

	for _, c := range cases {
		store := &memoryStore{c.models}

		models := store.list()

		if len(models) != len(c.models) {
			t.Errorf("Expected list to be length %d; got %d", len(c.models), len(models))
		}
	}
}
func TestMemoryStoreFind(t *testing.T) {
	cases := []struct {
		models []urlModel
		model  urlModel
		err    error
	}{
		{
			[]urlModel{urlModel{"abc", "http://url"}},
			urlModel{"abc", "http://url"},
			nil,
		},
		{
			[]urlModel{},
			urlModel{"abc", "http://url"},
			errorModelNotFound,
		},
	}

	for _, c := range cases {
		store := &memoryStore{c.models}
		model, err := store.find(c.model.ID)

		if c.err != nil {
			if err != c.err {
				t.Errorf("Expeceted error %s; got %s", c.err, err)
			}
		} else {
			if model.ID != c.model.ID {
				t.Errorf("Expected model ID %s; got %s", c.model.ID, model.ID)
			}
			if model.URL != c.model.URL {
				t.Errorf("Expected model URL %s; got %s", c.model.URL, model.URL)
			}

		}
	}
}
func TestMemoryStoreAdd(t *testing.T) {
	store := &memoryStore{make([]urlModel, 0)}
	store.add(urlModel{"abc", "http://url"})

	if len(store.models) != 1 {
		t.Errorf("Expected underlying store to increase by 1; got %d", len(store.models))
	}

	if store.models[0].ID != "abc" {
		t.Errorf("Expected underlying store to contain given model with id \"abc\"; got %s", store.models[0].ID)
	}
}
