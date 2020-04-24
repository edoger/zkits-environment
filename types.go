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

// Env type defines the runtime environment.
type Env string

// String method returns the current runtime environment string.
func (e Env) String() string { return string(e) }

// Is method returns whether the given runtime environment is equal to the
// current runtime environment.
func (e Env) Is(env Env) bool { return e == env }

// In method returns whether the current runtime environment is in the given
// runtime environment list.
func (e Env) In(envs []Env) bool {
	for i, j := 0, len(envs); i < j; i++ {
		if e.Is(envs[i]) {
			return true
		}
	}
	return false
}
