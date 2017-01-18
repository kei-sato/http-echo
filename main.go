package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
    "strconv"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	s := []string{
		fmt.Sprintf("%s %s %s", req.Method, req.URL.RequestURI(), req.Proto),
		fmt.Sprintf("Host: %s", req.Host),
		fmt.Sprintf("Content-Length: %d", req.ContentLength),
	}

	for k, v := range req.Header {
		s = append(s, fmt.Sprintf("%s: %s", k, strings.Join(v, ", ")))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	body := buf.String()
	if body != "" {
		s = append(s, "\n"+body)
	}

	
	headers := os.Getenv("RESPONSE_HEADERS")
	if headers != "" {
        var ss []string

        ss = strings.Split(headers, "\r\n")
        for _, pair := range ss {
            z := strings.Split(pair, ":")
            rw.Header().Set(z[0], z[1])
        }
	}
	
	status := os.Getenv("STATUS_CODE")
	if status != "" {
        statusCode, _ := strconv.Atoi(status)
        rw.WriteHeader(statusCode)
	}

	fmt.Fprint(rw, strings.Join(s, "\n"))
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

    if port == "443" {
        http.ListenAndServeTLS(":"+port, "cert.pem", "cert.key", nil)
    } else {
        http.ListenAndServe(":"+port, nil)
    }
}
