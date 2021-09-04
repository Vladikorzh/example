package http

import (
	"example/internal/metrics"
	"example/internal/metrics/encoding"
	"net/http"
)

type Handler struct {
	provider metrics.Provider
	encoder  encoding.Encoder
}

func NewHandler(provider metrics.Provider, encoder encoding.Encoder) http.Handler {
	return &Handler{
		provider: provider,
		encoder:  encoder,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetMetrics(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	m, err := h.provider.GetMetrics()
	if err != nil {
		InternalError(w, err)
		return
	}

	err = h.encoder.Encode(w, m)
	if err != nil {
		InternalError(w, err)
	}
}

func InternalError(w http.ResponseWriter, err error)  {
	http.Error(w, "internal error: " + err.Error(), http.StatusInternalServerError)
}