package application

import (
	"context"

	"mathmatic-test-server/internal/core/domain"
	"mathmatic-test-server/internal/ports/inbound"
	"mathmatic-test-server/internal/ports/outbound"
)

// ProblemService 는 inbound.ProblemUseCase 를 구현하는 유스케이스 계층입니다.
type ProblemService struct {
	repo outbound.ProblemRepository
}

var _ inbound.ProblemUseCase = (*ProblemService)(nil)

func NewProblemService(repo outbound.ProblemRepository) *ProblemService {
	return &ProblemService{repo: repo}
}

func (s *ProblemService) GetProblem(ctx context.Context, id string) (domain.Problem, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProblemService) CheckAnswer(ctx context.Context, id string, answer int) (bool, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return false, err
	}
	return p.IsCorrect(answer), nil
}
