package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"mathmatic-test-server/internal/core/domain"
	"mathmatic-test-server/internal/ports/inbound"

	_ "mathmatic-test-server/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Handler 는 HTTP 요청을 유스케이스 호출로 바꿔주는 driving 어댑터입니다.
type Handler struct {
	useCase inbound.ProblemUseCase
}

// NewHandler 는 라우팅이 설정된 http.Handler 를 반환합니다.
func NewHandler(uc inbound.ProblemUseCase) http.Handler {
	h := &Handler{useCase: uc}

	api := http.NewServeMux()
	api.HandleFunc("GET /healthz", h.healthz)
	api.HandleFunc("GET /problems/{id}", h.getProblem)
	api.HandleFunc("POST /problems/{id}/answer", h.checkAnswer)

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))
	mux.HandleFunc("GET /api/", httpSwagger.WrapHandler)
	return mux
}

func (h *Handler) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) getProblem(w http.ResponseWriter, r *http.Request) {
	p, err := h.useCase.GetProblem(r.Context(), r.PathValue("id"))
	if err != nil {
		writeError(w, err)
		return
	}
	// 정답은 응답에 포함하지 않습니다.
	writeJSON(w, http.StatusOK, map[string]string{"id": p.ID, "question": p.Question})
}

func (h *Handler) checkAnswer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Answer int `json:"answer"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "잘못된 요청 본문입니다"})
		return
	}

	correct, err := h.useCase.CheckAnswer(r.Context(), r.PathValue("id"), req.Answer)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"correct": correct})
}

func writeError(w http.ResponseWriter, err error) {
	if errors.Is(err, domain.ErrProblemNotFound) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "문제를 찾을 수 없습니다"})
		return
	}
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "서버 오류"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
