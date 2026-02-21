package handler

import (
	"net/http"

	"github.com/patakuti/markdown-proxy/internal/config"
)

type RemoteHandler struct {
	cfg *config.Config
}

func NewRemoteHandler(cfg *config.Config) *RemoteHandler {
	return &RemoteHandler{cfg: cfg}
}

func (h *RemoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Placeholder - will be implemented in Phase 5
	http.Error(w, "Remote handler not yet implemented", http.StatusNotImplemented)
}
