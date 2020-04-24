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
	"errors"
)

// ErrInvalidEnv represents that the given runtime environment is not
// registered or supported.
var ErrInvalidEnv = errors.New("invalid runtime environment")

// ErrLocked indicates that the current runtime environment is locked
// and cannot be changed.
var ErrLocked = errors.New("locked runtime environment")
