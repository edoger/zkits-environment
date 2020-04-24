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
	"sync/atomic"
	"testing"
)

func doTestEnvironment(f func()) {
	// Oh, we still need to restore the scene!
	defer func() {
		current = Development
		atomic.StoreInt32(&locked, 0)
		supported = []Env{Development, Testing, Prerelease, Production}
	}()

	f()
}

func TestEnvironment(t *testing.T) {
	// Copy from supported.
	list := []Env{Development, Testing, Prerelease, Production}

	doTestEnvironment(func() {
		// The default runtime environment is Development!
		if got := Get(); got != Development {
			t.Fatal(got)
		}

		for _, env := range list {
			if err := Set(env); err != nil {
				t.Fatal(err)
			}
			if got := Get(); got != env {
				t.Fatal(got)
			}
		}

		if err := Set("foo"); err == nil {
			t.Fatal("No error")
		} else {
			if err != ErrInvalidEnv {
				t.Fatal(err)
			}
		}

		Register("foo")

		if err := Set("foo"); err != nil {
			t.Fatal(err)
		}

		if Locked() {
			t.Fatal("Locked")
		}

		Lock()

		if !Locked() {
			t.Fatal("Not Locked")
		}
		for _, env := range list {
			if err := Set(env); err == nil {
				t.Fatal("No error")
			} else {
				if err != ErrLocked {
					t.Fatal(err)
				}
			}
		}
		if err := Set("foo"); err == nil {
			t.Fatal("No error")
		} else {
			if err != ErrLocked {
				t.Fatal(err)
			}
		}
	})
}

func TestSetAndLock(t *testing.T) {
	doTestEnvironment(func() {
		if Locked() {
			t.Fatal("Locked")
		}

		// The default runtime environment is Development!
		if got := Get(); got != Development {
			t.Fatal(got)
		}

		if err := SetAndLock(Testing); err != nil {
			t.Fatal(err)
		}
		if !Locked() {
			t.Fatal("Not Locked")
		}
		if got := Get(); got != Testing {
			t.Fatal(got)
		}

		if err := SetAndLock(Production); err == nil {
			t.Fatal("No error")
		} else {
			if err != ErrLocked {
				t.Fatal(err)
			}
		}
	})
}
