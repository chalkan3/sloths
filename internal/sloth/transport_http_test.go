package sloth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestHTTP(t *testing.T) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)

	svc := NewService(NewRepository(nil), nil, nil)
	r := NewHTTPServer(svc, logger)
	srv := httptest.NewServer(r)

	for _, testCase := range []struct {
		method, url, body string
		want              int
	}{
		{"POST", srv.URL + "/v1/sloth", `{"sloth": {"name": "maria", "family": "guicas"}}`, http.StatusOK},
		{"PUT", srv.URL + "/v1/sloth/2", `{"sloth": {"name": "lady", "family": "guicas"}}`, http.StatusOK},
		{"GET", srv.URL + "/v1/sloth/2", ``, http.StatusOK},
		{"GET", srv.URL + "/v1/sloth", ``, http.StatusOK},
	} {
		req, _ := http.NewRequest(testCase.method, testCase.url, strings.NewReader(testCase.body))
		resp, _ := http.DefaultClient.Do(req)
		if testCase.want != resp.StatusCode {
			t.Errorf("%s %s %s: want %d have %d", testCase.method, testCase.url, testCase.body, testCase.want, resp.StatusCode)
		}

	}
}
