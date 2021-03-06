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

// Package test provides utilities for testing the //pkg/aspect code.
package test

import (
	"errors"

	"istio.io/mixer/pkg/adapter"
)

// Logger is a test struct that implements logger and accessLogger aspects.
type Logger struct {
	adapter.AccessLogsBuilder
	adapter.ApplicationLogsBuilder

	DefaultCfg     adapter.AspectConfig
	EntryCount     int
	Logs           []adapter.LogEntry
	AccessLogs     []adapter.LogEntry
	ErrOnNewAspect bool
	ErrOnLog       bool
}

// NewLogger returns a new instance of the Logger aspect.
func (t *Logger) NewLogger(e adapter.Env, m adapter.AspectConfig) (adapter.ApplicationLogsAspect, error) {
	if t.ErrOnNewAspect {
		return nil, errors.New("new aspect error")
	}
	return t, nil
}

// NewAccessLogger returns a new instance of the accessLogger aspect.
func (t *Logger) NewAccessLogger(e adapter.Env, m adapter.AspectConfig) (adapter.AccessLogsAspect, error) {
	if t.ErrOnNewAspect {
		return nil, errors.New("new aspect error")
	}
	return t, nil
}

// Name returns the official name of this builder.
func (t *Logger) Name() string { return "testLogger" }

// Description returns a user-friendly description of this builder.
func (t *Logger) Description() string { return "A test logger" }

// DefaultConfig returns a default configuration struct for this adapter.
func (t *Logger) DefaultConfig() adapter.AspectConfig { return t.DefaultCfg }

// ValidateConfig determines whether the given configuration meets all correctness requirements.
func (t *Logger) ValidateConfig(c adapter.AspectConfig) (ce *adapter.ConfigErrors) { return nil }

// Log simulates processing a batch of log entries.
func (t *Logger) Log(l []adapter.LogEntry) error {
	if t.ErrOnLog {
		return errors.New("log error")
	}
	t.EntryCount++
	t.Logs = append(t.Logs, l...)
	return nil
}

// LogAccess simulates processing a batch of access log entries.
func (t *Logger) LogAccess(l []adapter.LogEntry) error {
	if t.ErrOnLog {
		return errors.New("log access error")
	}
	t.EntryCount++
	t.AccessLogs = append(t.AccessLogs, l...)
	return nil
}

// Close does nothing at the moment.
func (t *Logger) Close() error { return nil }

// NewLogEntry creates an adapter.LogEntry instance.
func NewLogEntry(n string, l map[string]interface{}, ts string, s adapter.Severity, tp string, sp map[string]interface{}) adapter.LogEntry {
	return adapter.LogEntry{
		LogName:       n,
		Labels:        l,
		Timestamp:     ts,
		Severity:      s,
		TextPayload:   tp,
		StructPayload: sp,
	}
}
