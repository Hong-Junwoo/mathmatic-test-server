package memory

import (
	"context"

	"mathmatic-test-server/internal/core/domain"
	"mathmatic-test-server/internal/ports/outbound"
)

// ProblemRepository 는 메모리에 문제를 보관하는 driven 어댑터입니다.
// 나중에 DB 어댑터로 교체해도 outbound.ProblemRepository 만 지키면 됩니다.
// CD Test
type ProblemRepository struct {
	problems map[string]domain.Problem
}

var _ outbound.ProblemRepository = (*ProblemRepository)(nil)

func NewProblemRepository(seed ...domain.Problem) *ProblemRepository {
	m := make(map[string]domain.Problem, len(seed))
	for _, p := range seed {
		m[p.ID] = p
	}
	return &ProblemRepository{problems: m}
}

func (r *ProblemRepository) FindByID(_ context.Context, id string) (domain.Problem, error) {
	p, ok := r.problems[id]
	if !ok {
		return domain.Problem{}, domain.ErrProblemNotFound
	}
	return p, nil
}
