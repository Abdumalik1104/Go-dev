package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var jwtKey = []byte("your_secret_key")
var validate = validator.New()

func main() {
	r := mux.NewRouter()

	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))
	r.Use(csrfMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Secure Web App"))
	}).Methods("GET")

	r.HandleFunc("/get-csrf-token", func(w http.ResponseWriter, r *http.Request) {
		token := csrf.Token(r)
		w.Write([]byte(token))
	}).Methods("GET")

	r.HandleFunc("/login", loginHandler).Methods("POST")

	r.Handle("/protected", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have accessed a protected route!"))
	}))).Methods("GET")

	r.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    ":8443",
		Handler: r,
	}
	log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}

type User struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func generateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	return token.SignedString(jwtKey)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})

}
