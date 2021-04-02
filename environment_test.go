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
	"strings"
	"testing"
)

func TestSetDefaultManager(t *testing.T) {
	defer func() { defaultManager = New() }()
	SetDefaultManager(NewEmpty())
}

func TestGetDefaultManager(t *testing.T) {
	if GetDefaultManager() == nil {
		t.Fatal("GetDefaultManager() return nil")
	}
}

func TestEnv_String(t *testing.T) {
	items := []struct {
		Given Env
		Want  string
	}{
		{Development, "development"},
		{Testing, "testing"},
		{Prerelease, "prerelease"},
		{Production, "production"},
		{Env("test"), "test"},
		{}, // For zero value.
	}

	for _, item := range items {
		if got := item.Given.String(); got != item.Want {
			t.Fatalf("Env.String(): want %s, got %s", item.Want, got)
		}
	}
}

func TestEnv_Is(t *testing.T) {
	env := Testing
	items := []struct {
		Given Env
		Want  bool
	}{
		{Development, false},
		{Testing, true},
		{Prerelease, false},
		{Production, false},
		{}, // For zero value.
	}

	for _, item := range items {
		if got := env.Is(item.Given); got != item.Want {
			t.Fatalf("Env.Is(): want %v, got %v", item.Want, got)
		}
	}
}

func TestEnv_In(t *testing.T) {
	env := Testing
	items := []struct {
		Given []Env
		Want  bool
	}{
		{[]Env{Development}, false},
		{[]Env{Testing}, true},
		{[]Env{Prerelease}, false},
		{[]Env{Production}, false},
		{[]Env{Development, Testing}, true},
		{[]Env{Development, Prerelease, Production}, false},
		{}, // For zero value.
	}

	for _, item := range items {
		if got := env.In(item.Given); got != item.Want {
			t.Fatalf("Env.In(): want %v, got %v", item.Want, got)
		}
	}
}

func TestDefaultManager(t *testing.T) {
	do := func(fs ...func()) {
		defer func() { defaultManager = New() }()

		for _, f := range fs {
			defaultManager = New()
			f()
			if defaultManager == nil {
				t.Fatal("Default manager is nil")
			}
		}
	}

	do(func() {
		if got := Get(); got != Development {
			t.Fatalf("Get(): %v", got)
		}
	}, func() {
		if got := Is(Development); got != true {
			t.Fatalf("Is(): %v", got)
		}
		if got := Is(Production); got != false {
			t.Fatalf("Is(): %v", got)
		}
	}, func() {
		if got := In([]Env{Development}); got != true {
			t.Fatalf("In(): %v", got)
		}
		if got := In([]Env{Production, Development}); got != true {
			t.Fatalf("In(): %v", got)
		}
		if got := In([]Env{Production}); got != false {
			t.Fatalf("In(): %v", got)
		}
		if got := In([]Env{Production, Testing}); got != false {
			t.Fatalf("In(): %v", got)
		}
	})

	do(func() {
		if got := Locked(); got != false {
			t.Fatalf("Locked(): %v", got)
		}

		Lock()

		if got := Locked(); got != true {
			t.Fatalf("Locked(): %v", got)
		}
	})

	do(func() {
		items := []Env{Development, Testing, Prerelease, Production}

		for _, env := range items {
			if err := Set(env); err != nil {
				t.Fatalf("Set(): %s", err)
			}
		}
	}, func() {
		if err := Set("unknown"); err == nil {
			t.Fatal("Set(): no error")
		} else {
			if err != ErrInvalidEnv {
				t.Fatalf("Set(): %s", err)
			}
		}

		if got := Registered("unknown"); got != false {
			t.Fatalf("Registered(): %v", got)
		}
		Register("unknown")
		if got := Registered("unknown"); got != true {
			t.Fatalf("Registered(): %v", got)
		}

		if err := Set("unknown"); err != nil {
			t.Fatalf("Set(): %s", err)
		}
	}, func() {
		Lock()

		if err := Set(Testing); err == nil {
			t.Fatal("Set(): no error")
		} else {
			if err != ErrLocked {
				t.Fatalf("Set(): %s", err)
			}
		}
	}, func() {
		if err := SetAndLock(Testing); err != nil {
			t.Fatalf("SetAndLock(): %s", err)
		}

		if got := Locked(); got != true {
			t.Fatalf("Locked(): %v", got)
		}

		if err := SetAndLock(Production); err == nil {
			t.Fatal("SetAndLock(): no error")
		} else {
			if err != ErrLocked {
				t.Fatalf("SetAndLock(): %s", err)
			}
		}
	})

	do(func() {
		var n int
		Listen(func(after, before Env) {
			if after != Testing || before != Development {
				t.Fatalf("Listen(): after %v, before %v", after, before)
			}
			n++
		})

		if err := Set(Testing); err != nil {
			t.Fatalf("Listen(): %s", err)
		}
		if n != 1 {
			t.Fatalf("Listen(): %d", n)
		}
	}, func() {
		var ss []string

		Listen(func(after, before Env) { ss = append(ss, "TEST") })
		if UnListen() == nil {
			t.Fatal("UnListen(): nil")
		}

		Listen(func(after, before Env) { ss = append(ss, "A") })
		Listen(func(after, before Env) { ss = append(ss, "B") })
		Listen(func(after, before Env) { ss = append(ss, "C") })

		if err := Set(Testing); err != nil {
			t.Fatalf("Set(): %s", err)
		}
		if got := strings.Join(ss, "-"); got != "A-B-C" {
			t.Fatalf("Listen(): %s", got)
		}

		if UnListen() == nil {
			t.Fatal("UnListen(): nil")
		}

		if err := Set(Prerelease); err != nil {
			t.Fatalf("Set(): %s", err)
		}
		if got := strings.Join(ss, "-"); got != "A-B-C-A-B" {
			t.Fatalf("Listen(): %s", got)
		}

		if got := len(UnListenAll()); got != 2 {
			t.Fatalf("UnListenAll(): %d", got)
		}

		if err := Set(Production); err != nil {
			t.Fatalf("Set(): %s", err)
		}
		if got := strings.Join(ss, "-"); got != "A-B-C-A-B" {
			t.Fatalf("Listen(): %s", got)
		}

		if UnListen() != nil {
			t.Fatal("UnListen(): not nil")
		}
		if got := len(UnListenAll()); got != 0 {
			t.Fatalf("UnListenAll(): %d", got)
		}
	})
}
