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
    'id': 'abc',
    'url': 'http://test.com/blah'
  },
  {
    'id': 'def',
    'url': 'ssh://github.com:taco/repo.git'
  }
]`

func createMockFileStore(input string) *fileStore {
	strReader := bytes.NewReader([]byte(input))
	mockFile := bufio.NewReadWriter(bufio.NewReader(strReader), nil)

	return &fileStore{file: mockFile}
}

func TestFileStoreListsURLs(t *testing.T) {}
func TestFileStoreFindURL(t *testing.T)   {}
func TestFileStoreAddURL(t *testing.T)    {}
