// Copyright 2017 Google Inc.
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

package adapterManager

import (
	"fmt"
	"testing"

	"istio.io/mixer/pkg/adapter"
	"istio.io/mixer/pkg/aspect"
)

type testBuilder struct {
	name string
}

func (t testBuilder) Name() string                                              { return t.name }
func (testBuilder) Close() error                                                { return nil }
func (testBuilder) Description() string                                         { return "mock builder for testing" }
func (testBuilder) DefaultConfig() adapter.AspectConfig                         { return nil }
func (testBuilder) ValidateConfig(c adapter.AspectConfig) *adapter.ConfigErrors { return nil }

type denyBuilder struct{ testBuilder }

func (denyBuilder) NewDenyChecker(env adapter.Env, cfg adapter.AspectConfig) (adapter.DenialsAspect, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestRegisterDenyChecker(t *testing.T) {
	reg := newRegistry(nil)
	builder := denyBuilder{testBuilder{name: "foo"}}

	reg.RegisterDenialsBuilder(builder)

	impl, ok := reg.FindBuilder(aspect.DenyKind, builder.Name())
	if !ok {
		t.Errorf("No builder by impl with name %s, expected builder: %v", builder.Name(), builder)
	}

	if deny, ok := impl.(denyBuilder); !ok || deny != builder {
		t.Errorf("reg.ByImpl(%s) expected builder '%v', actual '%v'", builder.Name(), builder, impl)
	}
}

type listBuilder struct{ testBuilder }

func (listBuilder) NewListChecker(env adapter.Env, cfg adapter.AspectConfig) (adapter.ListsAdapter, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestRegisterListChecker(t *testing.T) {
	reg := newRegistry(nil)
	builder := listBuilder{testBuilder{name: "foo"}}

	reg.RegisterListsBuilder(builder)

	impl, ok := reg.FindBuilder(aspect.ListKind, builder.Name())
	if !ok {
		t.Errorf("No builder by impl with name %s, expected builder: %v", builder.Name(), builder)
	}

	if deny, ok := impl.(listBuilder); !ok || deny != builder {
		t.Errorf("reg.ByImpl(%s) expected builder '%v', actual '%v'", builder.Name(), builder, impl)
	}
}

type loggerBuilder struct{ testBuilder }

func (loggerBuilder) NewLogger(env adapter.Env, cfg adapter.AspectConfig) (adapter.ApplicationLogsAspect, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestRegisterLogger(t *testing.T) {
	reg := newRegistry(nil)
	builder := loggerBuilder{testBuilder{name: "foo"}}

	reg.RegisterApplicationLogsBuilder(builder)

	impl, ok := reg.FindBuilder(aspect.LogKind, builder.Name())
	if !ok {
		t.Errorf("No builder by impl with name %s, expected builder: %v", builder.Name(), builder)
	}

	if deny, ok := impl.(loggerBuilder); !ok || deny != builder {
		t.Errorf("reg.ByImpl(%s) expected builder '%v', actual '%v'", builder.Name(), builder, impl)
	}
}

type accessLoggerBuilder struct{ testBuilder }

func (accessLoggerBuilder) NewAccessLogger(env adapter.Env, cfg adapter.AspectConfig) (adapter.AccessLogsAspect, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestRegistry_RegisterAccessLogger(t *testing.T) {
	reg := newRegistry(nil)
	builder := accessLoggerBuilder{testBuilder{name: "foo"}}

	reg.RegisterAccessLogsBuilder(builder)

	impl, ok := reg.FindBuilder(aspect.AccessLogKind, builder.Name())
	if !ok {
		t.Errorf("No builder by impl with name %s, expected builder: %v", builder.Name(), builder)
	}

	if deny, ok := impl.(accessLoggerBuilder); !ok || deny != builder {
		t.Errorf("reg.ByImpl(%s) expected builder '%v', actual '%v'", builder.Name(), builder, impl)
	}
}

type quotaBuilder struct{ testBuilder }

func (quotaBuilder) NewQuota(env adapter.Env, cfg adapter.AspectConfig, d map[string]*adapter.QuotaDefinition) (adapter.QuotasAspect, error) {
	return nil, fmt.Errorf("not implemented")
}

func TestRegisterQuota(t *testing.T) {
	reg := newRegistry(nil)
	builder := quotaBuilder{testBuilder{name: "foo"}}

	reg.RegisterQuotasBuilder(builder)
	impl, ok := reg.FindBuilder(aspect.QuotaKind, builder.Name())
	if !ok {
		t.Errorf("No builder by impl with name %s, expected builder: %v", builder.Name(), builder)
	}

	if deny, ok := impl.(quotaBuilder); !ok || deny != builder {
		t.Errorf("reg.ByImpl(%s) expected builder '%v', actual '%v'", builder.Name(), builder, impl)
	}
}

func TestCollision(t *testing.T) {
	reg := newRegistry(nil)
	name := "some name that they both have"

	a1 := listBuilder{testBuilder{name}}
	reg.RegisterListsBuilder(a1)

	if a, ok := reg.FindBuilder(aspect.ListKind, name); !ok || a != a1 {
		t.Errorf("Failed to get first adapter by impl name; expected: '%v', actual: '%v'", a1, a)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected to recover from panic registering duplicate builder, but recover was nil.")
		}
	}()

	a2 := listBuilder{testBuilder{name}}
	reg.RegisterListsBuilder(a2)
	t.Error("Should not reach this statement due to panic.")
}
