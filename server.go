package main

import (
	"crypto/sha512"
	"crypto/subtle"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func httpServer(config Config) http.Server {
	return http.Server{
		Addr:         config.ListenAddress,
		Handler:      handleRequest(config),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func tlsServer(config Config) http.Server {
	return http.Server{
		Addr:         config.ListenAddress,
		Handler:      handleRequest(config),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}
}

func startServer(config Config) {
	var server http.Server
	if config.UseTls {
		server = tlsServer(config)
	} else {
		server = httpServer(config)
	}
	log.Println("Starting server on " + config.ListenAddress)
	if config.UseTls {
		err := server.ListenAndServeTLS(config.TlsCertificateFile, config.TlsCertificateKey)
		if err != nil {
			log.Fatalln("Unable to start server!", err)
		}
	} else {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln("Unable to start server!", err)
		}
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
	expectedUsernameHash := sha512.Sum512([]byte(config.AuthUsername))
	expectedPasswordHash := sha512.Sum512([]byte(config.AuthPassword))

	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1
	return usernameMatch && passwordMatch
}
