// Copyright 2016 Google Inc.
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

package adapter

import (
	"google.golang.org/genproto/googleapis/rpc/status"
)

type (
	// DenialsAspect always fail with an error.
	DenialsAspect interface {
		Aspect

		// Deny always returns an error
		Deny() status.Status
	}

	// DenialsBuilder builds instances of the DenyChecker aspect.
	DenialsBuilder interface {
		Builder

		// NewDenyChecker returns a new instance of the DenyChecker aspect.
		NewDenyChecker(env Env, c AspectConfig) (DenialsAspect, error)
	}
)
