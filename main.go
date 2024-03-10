package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	filename = os.Getenv("LOGINEXPORTER_FILENAME")
	port     = os.Getenv("LOGINEXPORTER_PORT")
)

func loginreader() {

}

func main() {

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(port, nil)
}
