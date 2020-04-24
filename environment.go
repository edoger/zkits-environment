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
	"sync"
)

var (
	// Mutex when operating on the current runtime environment.
	mutex sync.Mutex

	// The current environment.
	current = Development

	// List of supported environments.
	supported = []Env{Development, Testing, Prerelease, Production}
)

// Get returns the current runtime environment.
func Get() Env {
	return current
}

// Register registers a custom runtime environment.
// If you want to add a custom environment, this method must be called
// before the Set() method.
func Register(env Env) {
	mutex.Lock()
	defer mutex.Unlock()

	if !env.In(supported) {
		supported = append(supported, env)
	}
}

// Set sets the current runtime environment.
// If the given environment is not supported, an ErrInvalidEnv error is returned.
func Set(env Env) error {
	mutex.Lock()
	defer mutex.Unlock()

	if !env.In(supported) {
		return ErrInvalidEnv
	}

	current = env
	return nil
}
