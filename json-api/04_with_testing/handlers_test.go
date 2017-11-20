package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestTruth(t *testing.T) {
	if true != false {
		t.Error("expected true to be true")
	}
}

var (
	h  *handler
	r  *mux.Router
	ts *httptest.Server
)

func TestMain(m *testing.M) {
	h = newHandler(testProverbs())
	r = newRouter(h)
	ts = httptest.NewServer(r)
	os.Exit(m.Run())
}

func TestNoRoot(t *testing.T) {
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	want := http.StatusNotFound
	got := res.StatusCode
	if want != got {
		t.Errorf("expected status code %d but got %d", want, got)
	}
}

func TestGetProverbs(t *testing.T) {
	endpoint := strings.Join([]string{ts.URL, "proverbs"}, "/")

	res, err := http.Get(endpoint)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}

	var result []Proverb
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Error(err)
	}

	want := len(h.proverbs)
	got := len(result)
	if want != got {
		t.Errorf("expected %d proverbs but got %d", want, got)
	}
}

func TestCreateProverb(t *testing.T) {
	endpoint := strings.Join([]string{ts.URL, "proverbs"}, "/")

	testCases := []struct {
		scenario           string
		payload            string
		expectedStatusCode int
	}{
		{
			"valid input",
			`{"text":"Warriors should suffer their pain silently.", "Philosopher": "Erin Hunter"}`,
			http.StatusCreated,
		},
		{
			"invalid input",
			`{"text":"Warriors should suffer their pain silently.", "Author": "Erin Hunter"}`,
			http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			body := strings.NewReader(tc.payload)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			if err != nil {
				t.Error(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Error(err)
			}

			if res.StatusCode != tc.expectedStatusCode {
				t.Errorf("expected status code %d but got %d", tc.expectedStatusCode, res.StatusCode)
			}
		})
	}
}

func testProverbs() []Proverb {
	return []Proverb{
		Proverb{ID: 1, Text: "Don't panic.", Philosopher: "Rob Pike"},
		Proverb{ID: 2, Text: "Concurrency is not parallelism.", Philosopher: "Rob Pike"},
		Proverb{ID: 3, Text: "Documentation is for users.", Philosopher: "Rob Pike"},
		Proverb{ID: 4, Text: "The bigger the interface, the weaker the abstraction.", Philosopher: "Rob Pike"},
		Proverb{ID: 5, Text: "Make the zero value useful.", Philosopher: "Rob Pike"},
	}
}
