package main

import "net/http"
import "strings"
import "fmt"
import "io/ioutil"
import "strconv"

func handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.URL.Path, ".wasm") {
			fmt.Println("serving wasm", w.Header().Get("Content-Type"))
			data, _ := ioutil.ReadFile(root + req.URL.Path)
			w.Header().Set("Content-Type", "application/wasm")
			w.Header().Set("Content-Length", strconv.FormatInt(int64(len(data)), 10))
			w.Write(data)
		} else {
			h.ServeHTTP(w, req)
		}
	})
}

const root = "/Users/eginez/repos/goland/src/github.com/eginez/ProofOfWork"

func main() {
	fs := http.FileServer(http.Dir(root))
	fmt.Println("Running server")
	http.ListenAndServe(":8000", handler(fs))
}
