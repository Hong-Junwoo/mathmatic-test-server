package di

import (
	"net/http"

	"mathmatic-test-server/internal/adapters/driven/memory"
	"mathmatic-test-server/internal/adapters/driving/httpapi"
	"mathmatic-test-server/internal/core/application"
	"mathmatic-test-server/internal/core/domain"
)

// NewHTTPHandler 는 모든 의존성을 조립(wiring)해 HTTP 핸들러를 반환합니다.
// driven 어댑터 -> 유스케이스 -> driving 어댑터 순으로 연결합니다.
func NewHTTPHandler() http.Handler {
	repo := memory.NewProblemRepository(
		domain.Problem{ID: "1", Question: "1 + 1 = ?", Answer: 2},
		domain.Problem{ID: "2", Question: "7 * 8 = ?", Answer: 56},
	)
	svc := application.NewProblemService(repo)
	return httpapi.NewHandler(svc)
}
