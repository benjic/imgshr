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
	"bytes"
	"testing"
)

const testUrls string = `[
  {
    "ID": "abc",
    "url": "http://test.com/blah"
  },
  {
    "ID": "def",
    "url": "ssh://github.com:taco/repo.git"
  }
]`

func createMockFileStore(input string) (*fileStore, *bytes.Buffer) {
	strReader := bytes.NewReader([]byte(input))
	bufWriter := bytes.NewBuffer(nil)
	mockFile := bufio.NewReadWriter(bufio.NewReader(strReader), bufio.NewWriter(bufWriter))

	return &fileStore{file: mockFile}, bufWriter
}

func TestFileStoreListsURLs(t *testing.T) {
	cases := []struct {
		source string
		models []URLModel
	}{
		{
			testUrls,
			[]URLModel{
				URLModel{"abc", "http://test.com/blah"},
				URLModel{"def", "ssh://github.com:taco/repo.git"},
			},
		},
	}

	for _, c := range cases {
		store, _ := createMockFileStore(c.source)
		urls := store.list()

		if len(urls) != len(c.models) {
			t.Errorf("Expected filestore to return %d items; got: %d", len(c.models), len(urls))
		}

		for i, url := range urls {
			if url.ID != c.models[i].ID {
				t.Errorf("Expected model ID to be %s; got: %s", c.models[i].ID, url.ID)
			}

			if url.URL != c.models[i].URL {
				t.Errorf("Expected model url to be %s; got: %s", c.models[i].URL, url.URL)
			}
		}
	}
}

func TestFileStoreFindURL(t *testing.T) {
	cases := []struct {
		source string
		model  URLModel
		err    error
	}{
		{
			testUrls,
			URLModel{"def", "ssh://github.com:taco/repo.git"},
			nil,
		},
		{
			testUrls,
			URLModel{"xyz", "http://www.tacobell.com"},
			ErrorModelNotFound,
		},
	}

	for _, c := range cases {
		store, _ := createMockFileStore(c.source)
		url, err := store.find(c.model.ID)
		if c.err != nil {
			if err != c.err {
				t.Errorf("Expected to get error %s; got: %s", c.err, err)
			}
		} else {

			if url.ID != c.model.ID {
				t.Errorf("Expected model ID to be %s; got: %s", c.model.ID, url.ID)
			}

			if url.URL != c.model.URL {
				t.Errorf("Expected model ID to be %s; got: %s", c.model.URL, url.URL)
			}
		}
	}
}

func TestFileStoreAddURL(t *testing.T) {
	cases := []struct {
		input, output string
		model         URLModel
		err           error
	}{
		{
			"[]",
			"[{\"ID\":\"def\",\"URL\":\"ssh://github.com:taco/repo.git\"}]\n",
			URLModel{"def", "ssh://github.com:taco/repo.git"},
			nil,
		},
	}

	for _, c := range cases {

		store, buf := createMockFileStore(c.input)

		err := store.add(c.model)
		output := buf.String()

		if err != c.err {
			t.Errorf("Expected to get error %s; got: %s", c.err, err)
		}

		if output != c.output {
			t.Errorf("Expected to write output:\nExpected: '%s'\n     got: '%s'\n", c.output, output)
		}
	}
}
