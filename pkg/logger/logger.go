package logger

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(c Config) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	level, err := zerolog.ParseLevel(c.Level)
	if err != nil {
		zerolog.SetGlobalLevel(level)
	}

	log.Logger = log.With().
		Logger()

	if c.PrettyConsole {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	}

	log.Info().Msg("logger initialized")
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.status = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

type Config struct {
	AppName       string `envconfig:"APP_NAME"    required:"true"`
	AppVersion    string `envconfig:"APP_VERSION" required:"true"`
	Level         string `default:"error"         envconfig:"LOGGER_LEVEL"`
	PrettyConsole bool   `default:"false"         envconfig:"LOGGER_PRETTY_CONSOLE"`
}

// --- middleware для логирования ---
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(lrw, r)

		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", lrw.status).
			Msg("handled request")

	})
}
