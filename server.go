package main

import (
	"crypto/sha512"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"time"
)

func startServer(config Config) {
	server := &http.Server{
		Addr:         config.listenAddress,
		Handler:      handleRequest(config),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on " + config.listenAddress)
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatalln("Unable to start server!", err2)
	}
}

func handleRequest(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/activate" && r.Method == "POST" {
			username, password, ok := r.BasicAuth()

			if ok && checkCredentials(config, username, password) {
				response, err := handleShutdown(config)
				if err != nil {
					log.Println("Error shutting down!", err)
					w.WriteHeader(500)
					_, _ = fmt.Fprint(w, err)
					return
				}
				_, _ = fmt.Fprint(w, response)
				return
			} else {
				w.Header().Set("WWW-Authenticate", `Basic realm="shutdownd", charset="UTF-8"`)
				http.Error(w, "Credentials Required!", http.StatusUnauthorized)
				return
			}
		} else {
			log.Println("Not found: '" + r.RequestURI)
			w.WriteHeader(404)
		}
	}
}

func checkCredentials(config Config, username string, password string) bool {
	usernameHash := sha512.Sum512([]byte(username))
	passwordHash := sha512.Sum512([]byte(password))
	expectedUsernameHash := sha512.Sum512([]byte(config.authUsername))
	expectedPasswordHash := sha512.Sum512([]byte(config.authPassword))

	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1
	return usernameMatch && passwordMatch
}
