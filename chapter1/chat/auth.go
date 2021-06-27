package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if _,err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location","/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}else if err != nil {
		panic(err.Error())
	}else {
		h.next.ServeHTTP(w,r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next:handler}
}

// loginHandlerはサードパーティーへのログイン処理を受け持ちます。
// パスの形式: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path,"/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		log.Println("TODO: Login Method",provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "action %s is not supported",action)
	}
}