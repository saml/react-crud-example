package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func getEnvInt(envVar string, defaultValue int) int {
	i, err := strconv.ParseInt(os.Getenv(envVar), 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}

func getEnvString(envVar string, defaultValue string) string {
	s := os.Getenv(envVar)
	if s == "" {
		return defaultValue
	}
	return s
}

func main() {
	pwd, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	port := flag.Int("port", getEnvInt("PORT", 8080), "Port to listen to.")
	docRoot := flag.String("htdocs", getEnvString("HTDOCS", pwd), "Doc root")
	flag.Parse()

	bind := fmt.Sprintf(":%d", *port)
	log.Printf("Listening to http://localhost%s", bind)
	http.ListenAndServe(bind, http.FileServer(http.Dir(*docRoot)))
}
