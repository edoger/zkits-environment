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

const (
	// Development is the development environment, which is also the default
	// runtime environment.
	Development Env = "development"
	// Testing is a test environment, usually used for initial quality acceptance.
	Testing Env = "testing"
	// Prerelease is a pre release environment, usually used for grayscale
	// testing or quality acceptance.
	Prerelease Env = "prerelease"
	// Production is the production environment and the final deployment
	// environment of the application.
	Production Env = "production"
)

// The global default runtime environment manager.
var defaultManager = New()

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

// Get returns the current runtime environment.
func Get() Env { return defaultManager.Get() }

// Register registers a custom runtime environment.
// If you want to add a custom environment, this method must be called
// before the Set() method.
func Register(env Env) { defaultManager.Register(env) }

// Lock locks the current runtime environment.
// After locking, the current runtime environment cannot be changed.
func Lock() { defaultManager.Lock() }

// Locked returns whether the current runtime environment is locked.
func Locked() bool { return defaultManager.Locked() }

// Set sets the current runtime environment.
// If the given runtime environment is not supported, ErrInvalidEnv error is returned.
// If the current runtime environment is locked, ErrLocked error is returned.
func Set(env Env) error { return defaultManager.Set(env) }

// SetAndLock sets and locks the current runtime environment.
// If the runtime environment settings fail, they are not locked.
func SetAndLock(env Env) error { return defaultManager.SetAndLock(env) }
