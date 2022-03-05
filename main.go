package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func handler(w http.ResponseWriter, req *http.Request) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err == nil {
		fmt.Fprintf(w, "%s\n", ip)
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {
	port := getEnv("PORT", "3000")

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
