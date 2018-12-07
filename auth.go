package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// implements and wraps http.Handler interface
type authHandler struct {
	// this type extends http.Handler
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if there is an auth cookie
	_, err := r.Cookie("auth")

	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// simply creates authHandler that wraps any other http.Handler
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// dynamic path-matching mechanism
// loginHandler handles the third-party login process.
// format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
    if len(segs) < 3 {
        http.Error(w, "Unsupported url", http.StatusBadRequest)
        return
    }
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
