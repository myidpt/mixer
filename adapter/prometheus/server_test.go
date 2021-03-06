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

package prometheus

import (
	"fmt"
	"net/http"
	"testing"

	"istio.io/mixer/pkg/adapter"
)

type testLogger struct {
	adapter.Logger
}

func (t testLogger) Errorf(format string, args ...interface{}) error { return nil }

func TestServer(t *testing.T) {
	testAddr := "127.0.0.1:9992"
	s := newServer(testAddr)
	if err := s.Start(testLogger{}); err != nil {
		t.Fatalf("Start() failed unexpectedly: %v", err)
	}

	testURL := fmt.Sprintf("http://%s%s", testAddr, metricsPath)
	// verify a response is returned from "/metrics"
	resp, err := http.Get(testURL)
	if err != nil {
		t.Fatalf("Failed to retrieve '%s' path: %v", metricsPath, err)
	}

	defer func() {
		err := resp.Body.Close()
		t.Logf("Error closing response body: %v", err)
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("http.GET => %v, wanted '%v'", resp.StatusCode, http.StatusOK)
	}

	if err := s.Close(); err != nil {
		t.Errorf("Failed to close server properly: %v", err)
	}

	if resp, err := http.Get(testURL); err == nil {
		t.Errorf("http.GET should have failed for '%s'; got %v", metricsPath, resp)
	}
}
