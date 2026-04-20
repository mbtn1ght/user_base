package router

import (
	"path"
	"strings"
	"user_base/internal/controller/http/v1"

	"user_base/internal/usecase"
	"user_base/pkg/logger"

	"github.com/go-chi/chi/v5"
)

func New(basePath string, uc *usecase.UseCase) *chi.Mux {
	r := chi.NewRouter()
	r.Use(logger.LoggingMiddleware)

	controller := v1.New(uc)
	normalizedBasePath := normalizeBasePath(basePath)

	mountProfileRoutes(r, controller)

	if normalizedBasePath != "/" {
		r.Route(normalizedBasePath, func(router chi.Router) {
			mountProfileRoutes(router, controller)
		})
	}

	return r
}

func mountProfileRoutes(router chi.Router, controller *v1.Handlers) {
	router.Post("/profile", controller.CreateProfile)
	router.Get("/profile/{profileID}", controller.GetProfile)
	router.Put("/profile/{profileID}", controller.UpdateProfile)
	router.Put("/profile", controller.UpdateProfile)
	router.Delete("/profile/{profileID}", controller.DeleteProfile)
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
