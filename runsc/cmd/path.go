// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"path/filepath"
)

// absPath turns the given path into an absolute path (if it is not already
// absolute) by prepending the base path.
func absPath(base, rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(base, rel)
}

// getwdOrDie returns the current working directory and dies if it cannot.
func getwdOrDie() string {
	wd, err := os.Getwd()
	if err != nil {
		Fatalf("error getting current working directory: %v", err)
	}
	return wd
}
