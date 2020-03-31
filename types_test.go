// Copyright 2020 The ZKits Project Authors.
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

package environment

import (
	"testing"
)

func TestEnv_String(t *testing.T) {
	var env Env
	if got := env.String(); got != "" {
		t.Fatal(got)
	}

	env = "test"
	if got := env.String(); got != "test" {
		t.Fatal(got)
	}
}

func TestEnv_Is(t *testing.T) {
	var env Env
	if !env.Is("") {
		t.Fatal("FALSE")
	}

	env = "test"
	if !env.Is("test") {
		t.Fatal("FALSE")
	}
}

func TestEnv_In(t *testing.T) {
	var env Env
	if env.In(nil) {
		t.Fatal("TRUE")
	}

	env = "test"
	if !env.In([]Env{"test"}) {
		t.Fatal("FALSE")
	}
}
