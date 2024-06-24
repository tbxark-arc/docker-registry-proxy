package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	registryHost = "registry-1.docker.io"
	serveAddress = "localhost:8989"
	BuildVersion = "dev"
)

func main() {
	flag.StringVar(&registryHost, "registry", registryHost, "Docker Registry Host")
	flag.StringVar(&serveAddress, "address", serveAddress, "Serve Address")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(serveAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	originalHost := r.Host
	if strings.HasPrefix(path, "/v2/") {
		registryURL := fmt.Sprintf("https://%s%s", registryHost, path)
		req, err := http.NewRequest(r.Method, registryURL, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for name, values := range r.Header {
			for _, value := range values {
				req.Header.Add(name, value)
			}
		}
		req.Header.Set("Host", registryHost)

		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}
		w.Header().Set("Access-Control-Allow-Origin", originalHost)
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Printf("Error copying response body: %v", err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Docker Registry Proxy"))
	}
}
