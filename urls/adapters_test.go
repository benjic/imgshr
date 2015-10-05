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
	"testing"

	"github.com/benjic/shrturl/faststore"
)

var mockURLModels = []faststore.URLModel{
	faststore.URLModel{Slug: "abc", URL: "http://url"},
	faststore.URLModel{Slug: "def", URL: "http://url"},
}

type mockFastStorer struct {
	allUrlsCalled bool
	getURLCalled  bool
	addURLCalled  bool
	missingItem   bool
	passedModel   faststore.URLModel
	id            string
}

func (s *mockFastStorer) AddURL(model faststore.URLModel) {
	s.addURLCalled = true
	s.passedModel = model
}
func (s *mockFastStorer) GetURL(id string) (faststore.URLModel, error) {
	s.getURLCalled = true
	if s.missingItem {
		return faststore.URLModel{}, errors.New("Mock missing item error")
	}
	return mockURLModels[0], nil
}
func (s *mockFastStorer) AllURLs() ([]faststore.URLModel, error) {
	s.allUrlsCalled = true
	return mockURLModels, nil
}
func createMockFastStorer() *mockFastStorer {
	return &mockFastStorer{}
}

func TestURLFastStorerAdapterAdd(t *testing.T) {
	mockStorer := createMockFastStorer()
	adapter := newURLFastStorerAdapter(mockStorer)

	adapter.add(urlModel{ID: "xyz", URL: "http://url"})

	if !mockStorer.addURLCalled {
		t.Errorf("Expected adapter to call AddURL but did not")
	}

}

func TestURLFastStorerAdapterList(t *testing.T) {
	mockStorer := createMockFastStorer()
	adapter := newURLFastStorerAdapter(mockStorer)

	for i, model := range adapter.list() {
		if model.ID != modelID(mockURLModels[i].Slug) {
			t.Errorf("Expect adapter to %d model ID. Got %s want %s", i, model.ID, mockURLModels[i].Slug)
		}
		if model.URL != mockURLModels[i].URL {
			t.Errorf("Expect adapter to %d model URL. Got %s want %s", i, model.URL, mockURLModels[i].URL)
		}
	}

	if !mockStorer.allUrlsCalled {
		t.Errorf("Expected adapter to call AllURLs function but it didn't")
	}
}

func TestURLFastStorerAdapterItem(t *testing.T) {
	mockStorer := createMockFastStorer()
	adapter := newURLFastStorerAdapter(mockStorer)

	model, err := adapter.find(modelID("abc"))

	if !mockStorer.getURLCalled {
		t.Errorf("Expected adapter to call underlying GetURL function but did not")
	}

	if err != nil {
		t.Errorf("Did not expect error from adapter find, %s", err)
	}

	if model.ID != modelID(mockURLModels[0].Slug) {
		t.Errorf("Expect adapter to  model ID. Got %s want %s", model.ID, mockURLModels[0].Slug)
	}

	if model.URL != mockURLModels[0].URL {
		t.Errorf("Expect adapter to model URL. Got %s want %s", model.URL, mockURLModels[0].URL)
	}
}

func TestURLFastStorerAdapterMissingItem(t *testing.T) {
	mockStorer := createMockFastStorer()
	mockStorer.missingItem = true

	adapter := newURLFastStorerAdapter(mockStorer)

	model, err := adapter.find(modelID("bogus"))

	if !mockStorer.getURLCalled {
		t.Errorf("Expected adapter to call underlying GetURL function but did not")
	}

	if err != errorModelNotFound {
		t.Errorf("Expected adapter to return proper error when no model found; got %s %s", err, model)
	}
}
