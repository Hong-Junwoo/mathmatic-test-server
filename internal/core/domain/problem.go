package domain

import "errors"

// ErrProblemNotFound 는 요청한 문제가 존재하지 않을 때 반환됩니다.
var ErrProblemNotFound = errors.New("problem not found")

// Problem 은 수학 문제 하나를 나타내는 도메인 엔티티입니다.
type Problem struct {
	ID       string
	Question string
	Answer   int
}

// IsCorrect 는 제출된 답이 정답인지 판단합니다.
func (p Problem) IsCorrect(answer int) bool {
	return p.Answer == answer
}
