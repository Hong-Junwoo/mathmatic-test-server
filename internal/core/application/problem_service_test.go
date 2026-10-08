package application

import (
	"context"
	"errors"
	"testing"

	"mathmatic-test-server/internal/core/domain"
)

// fakeRepo 는 아웃바운드 포트를 흉내 내는 테스트용 구현입니다.
type fakeRepo struct {
	problems map[string]domain.Problem
}

func (f fakeRepo) FindByID(_ context.Context, id string) (domain.Problem, error) {
	p, ok := f.problems[id]
	if !ok {
		return domain.Problem{}, domain.ErrProblemNotFound
	}
	return p, nil
}

func newTestService() *ProblemService {
	return NewProblemService(fakeRepo{problems: map[string]domain.Problem{
		"1": {ID: "1", Question: "1 + 1 = ?", Answer: 2},
	}})
}

func TestCheckAnswer(t *testing.T) {
	svc := newTestService()

	ok, err := svc.CheckAnswer(context.Background(), "1", 2)
	if err != nil || !ok {
		t.Fatalf("정답이어야 함: ok=%v err=%v", ok, err)
	}

	ok, err = svc.CheckAnswer(context.Background(), "1", 3)
	if err != nil || ok {
		t.Fatalf("오답이어야 함: ok=%v err=%v", ok, err)
	}
}

func TestCheckAnswerNotFound(t *testing.T) {
	_, err := newTestService().CheckAnswer(context.Background(), "999", 1)
	if !errors.Is(err, domain.ErrProblemNotFound) {
		t.Fatalf("ErrProblemNotFound 기대, 실제: %v", err)
	}
}
