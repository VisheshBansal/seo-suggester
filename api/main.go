package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
}

func main() {
	mode := viper.GetString("mode")

	// Initializing handlers
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `"Ok! Working!"`)
		w.WriteHeader(http.StatusOK)
		return
	})
	// HTTP(s) binding
	serverprefix := "server_" + mode
	host := viper.GetString(serverprefix + ".host")
	port := os.Getenv("PORT")

	if port == "" {
		port = viper.GetString(serverprefix + ".port")
	}

	conn := host + ":" + port

	srv := &http.Server{
		Addr:    conn,
		Handler: r,
	}

	log.Printf("Starting in %s mode", mode)
	log.Printf("Server running on %s", conn)
	log.Fatal(srv.ListenAndServe())
}
