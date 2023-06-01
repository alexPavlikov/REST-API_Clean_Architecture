package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var appErr *AppError
		err := h(w, r)
		if errors.As(err, &appErr) {
			if errors.Is(appErr, ErrorNotFound) {
				w.WriteHeader(http.StatusNotFound)
				w.Write(ErrorNotFound.Marhsal())
				return
			} else if errors.Is(err, NoAuthError) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(NoAuthError.Marhsal())
				return
			}

			err = err.(*AppError)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(appErr.Marhsal())
		}
		// w.WriteHeader(http.StatusTeapot)
		// w.Write(systemError(err).Marhsal())
	}
}
