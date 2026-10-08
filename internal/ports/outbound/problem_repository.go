package outbound

import (
	"context"

	"mathmatic-test-server/internal/core/domain"
)

// ProblemRepository 는 애플리케이션이 외부(DB 등)에 요구하는 기능입니다. (아웃바운드 포트)
type ProblemRepository interface {
	FindByID(ctx context.Context, id string) (domain.Problem, error)
}
