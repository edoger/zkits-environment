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
	"sync/atomic"
)

var (
	// Mutex when operating on the current runtime environment.
	mutex sync.RWMutex

	// The current environment.
	current = Development

	// Is the current runtime environment locked?
	locked = int32(0)

	// List of supported environments.
	supported = []Env{Development, Testing, Prerelease, Production}
)

// Get returns the current runtime environment.
func Get() Env {
	mutex.RLock()
	defer mutex.RUnlock()

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

// Lock locks the current runtime environment.
// After locking, the current runtime environment cannot be changed.
func Lock() {
	mutex.Lock()
	defer mutex.Unlock()

	atomic.StoreInt32(&locked, 1)
}

// Locked returns whether the current runtime environment is locked.
func Locked() bool {
	return atomic.LoadInt32(&locked) == 1
}

// Set sets the current runtime environment.
// If the given runtime environment is not supported, ErrInvalidEnv error is returned.
// If the current runtime environment is locked, ErrLocked error is returned.
func Set(env Env) error {
	mutex.Lock()
	defer mutex.Unlock()

	return doSet(env)
}

func doSet(env Env) error {
	if Locked() {
		return ErrLocked
	}
	if !env.In(supported) {
		return ErrInvalidEnv
	}

	current = env
	return nil
}

// SetAndLock sets and locks the current runtime environment.
// If the runtime environment settings fail, they are not locked.
func SetAndLock(env Env) error {
	mutex.Lock()
	defer mutex.Unlock()

	if err := doSet(env); err != nil {
		return err
	}

	atomic.StoreInt32(&locked, 1)
	return nil
}
