package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mathmatic-test-server/internal/adapters/driven/memory"
	"mathmatic-test-server/internal/core/application"
	"mathmatic-test-server/internal/core/domain"
)

func newTestHandler() http.Handler {
	repo := memory.NewProblemRepository(domain.Problem{ID: "1", Question: "1 + 1 = ?", Answer: 2})
	return NewHandler(application.NewProblemService(repo))
}

func do(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestHealthz(t *testing.T) {
	rec := do(newTestHandler(), http.MethodGet, "/healthz", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("상태 코드 = %d", rec.Code)
	}
}

func TestGetProblem(t *testing.T) {
	h := newTestHandler()

	if rec := do(h, http.MethodGet, "/problems/1", ""); rec.Code != http.StatusOK {
		t.Fatalf("상태 코드 = %d", rec.Code)
	}
	if rec := do(h, http.MethodGet, "/problems/999", ""); rec.Code != http.StatusNotFound {
		t.Fatalf("상태 코드 = %d, 404 기대", rec.Code)
	}
}

func TestCheckAnswer(t *testing.T) {
	h := newTestHandler()

	rec := do(h, http.MethodPost, "/problems/1/answer", `{"answer":2}`)
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), `"correct":true`) {
		t.Fatalf("정답 응답이 아님: %d %s", rec.Code, rec.Body.String())
	}

	rec = do(h, http.MethodPost, "/problems/1/answer", `not json`)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("상태 코드 = %d, 400 기대", rec.Code)
	}
}
