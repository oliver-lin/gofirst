package main

import (
	"fmt"
	"net/http"
	"os"
)

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.Form)
}
func main() {
	fmt.Println("net http包使用")
	serv := http.Server{
		Addr: ":8003",
	}
	http.HandleFunc("/process", process)
	serv.ListenAndServe()

}

var apiKey = os.Getenv("TWITTER_BEARER_TOKEN")

type Tweet = struct{ Text string }
