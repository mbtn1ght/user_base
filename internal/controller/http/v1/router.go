package v1

import (
	"path"
	"strings"

	"user_base/internal/controller/http"
	"user_base/internal/usecase"
	"user_base/pkg/logger"

	"github.com/go-chi/chi/v5"
)

func InitializeRouter(basePath string, uc *usecase.UseCase) *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.LoggingMiddleware)

	controller := http.New(uc)
	normalizedBasePath := normalizeBasePath(basePath)

	mountProfileRoutes(r, controller)

	if normalizedBasePath != "/" {
		r.Route(normalizedBasePath, func(router chi.Router) {
			mountProfileRoutes(router, controller)
		})
	}

	return r
}

func mountProfileRoutes(router chi.Router, controller *http.Handlers) {
	router.Post("/profile", controller.CreateProfile)
}

func normalizeBasePath(basePath string) string {
	trimmed := strings.TrimSpace(basePath)
	if trimmed == "" {
		return "/"
	}

	if !strings.HasPrefix(trimmed, "/") {
		trimmed = "/" + trimmed
	}

	cleaned := path.Clean(trimmed)
	if cleaned == "." {
		return "/"
	}

	return cleaned
}
