package main

import (
		"fmt"
		"net/http"
		"os"
		"strings"
		"log"
)

func ekstrakEnv() (map[string]string, error) {
	environS := os.Environ()
	environ := make(map[string]string)
	for _, val := range environS {
		split := strings.SplitN(val, "=", 2)
		if len(split) != 2 {
			return environ, fmt.Errorf("Some weird env vars")
		}
		environ[split[0]] = split[1]
	}
	for key := range environ {
		if !(strings.HasSuffix(key, "_SERVICE_HOST") ||
			strings.HasSuffix(key, "_SERVICE_PORT")) {
			delete(environ, key)
		}
	}
	return environ, nil
}


func lihatStatus(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Environment Variabel: ", os.Environ(), "\n")
	
	ekstraksi, err := ekstrakEnv()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(resp, "Ekstraksi Environment: ", ekstraksi, "\n")

}

func main() {
	http.HandleFunc("/", lihatStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}