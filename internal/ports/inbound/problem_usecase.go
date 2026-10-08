package inbound

import (
	"context"

	"mathmatic-test-server/internal/core/domain"
)

// ProblemUseCase 는 애플리케이션이 외부(HTTP 등)에 제공하는 기능입니다. (인바운드 포트)
type ProblemUseCase interface {
	GetProblem(ctx context.Context, id string) (domain.Problem, error)
	CheckAnswer(ctx context.Context, id string, answer int) (bool, error)
}
